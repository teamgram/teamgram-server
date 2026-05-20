// Copyright (c) 2026-present, The Teamgram Authors (https://teamgram.net).
//  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Author: teamgramio (teamgram.io@gmail.com)

package svc

import (
	"context"
	"time"

	gatewayclient "github.com/teamgram/teamgram-server/v2/app/interface/gateway/client"
	"github.com/teamgram/teamgram-server/v2/app/messenger/userupdates/internal/config"
	"github.com/teamgram/teamgram-server/v2/app/messenger/userupdates/internal/repository"
	receiverevent "github.com/teamgram/teamgram-server/v2/app/messenger/userupdates/internal/repository/event"
	"github.com/teamgram/teamgram-server/v2/app/service/authsession/authsession"
	authsessionclient "github.com/teamgram/teamgram-server/v2/app/service/authsession/client"
	chatclient "github.com/teamgram/teamgram-server/v2/app/service/biz/chat/client"
	dialogclient "github.com/teamgram/teamgram-server/v2/app/service/biz/dialog/client"
	userclient "github.com/teamgram/teamgram-server/v2/app/service/biz/user/client"
	"github.com/teamgram/teamgram-server/v2/pkg/net/kitex"
)

type UserUpdatesRepository interface {
	ApplyUserOperation(ctx context.Context, in repository.ApplyUserOperationInput) (*repository.ApplyUserOperationResult, error)
	ApplyUserOperationBatch(ctx context.Context, inputs []repository.ApplyUserOperationInput) ([]repository.ApplyUserOperationResult, error)
	GetOperationResult(ctx context.Context, userID int64, operationID string) (*repository.OperationResult, error)
	GetState(ctx context.Context, userID int64, permAuthKeyID int64) (*repository.UserState, error)
	GetDifference(ctx context.Context, in repository.GetDifferenceInput) (*repository.GetDifferenceResult, error)
	AppendAuthSeqUpdate(ctx context.Context, in repository.AuthSeqUpdateAppendInput) (*repository.AuthSeqUpdateAppendResult, error)
	AppendDialogPtsSideEffect(ctx context.Context, in repository.DialogSideEffectAppendInput) (*repository.PtsAppendResult, error)
	ListDialogProjections(ctx context.Context, userID int64, cursor repository.DialogProjectionCursor, limit int32) ([]repository.DialogProjection, error)
	GetDialogProjectionsByPeers(ctx context.Context, userID int64, peers []repository.DialogProjectionPeer) (map[repository.DialogProjectionPeer]repository.DialogProjection, error)
	GetMessageViewsByPeerSeqs(ctx context.Context, userID int64, peers []repository.MessageViewPeerSeq) (map[repository.MessageViewPeerSeq]repository.MessageView, error)
	CountVisibleDialogs(ctx context.Context, userID int64) (int32, error)
}

type backgroundWorker interface {
	Run(ctx context.Context)
	Stop()
}

type waitableBackgroundWorker interface {
	Wait()
}

type PushOutboxNotifier interface {
	Wake()
}

type AuthsessionClient interface {
	AuthsessionGetPermAuthKeyIds(ctx context.Context, in *authsession.TLAuthsessionGetPermAuthKeyIds) (*authsession.VectorLong, error)
}

type ServiceContext struct {
	Config             config.Config
	Repo               UserUpdatesRepository
	AuthsessionClient  AuthsessionClient
	PushOutboxNotifier PushOutboxNotifier
	workers            []backgroundWorker
	closers            []interface{ Close() error }
	cancel             context.CancelFunc
}

func NewServiceContext(c config.Config) *ServiceContext {
	repo := repository.NewRepository(c)
	var userProjectionClient repository.UserProjectionClient
	var chatProjectionClient repository.ChatProjectionClient
	if hasRPCClientConfig(c.BizServiceClient) {
		userProjectionClient = userclient.NewUserClient(userclient.MustNewKitexClient(c.BizServiceClient))
		bizChatClient := chatclient.NewChatClient(chatclient.MustNewKitexClient(c.BizServiceClient))
		chatProjectionClient = bizChatClient
		repo.SetPeerProjectionClients(userProjectionClient, chatProjectionClient)
	}
	sc := &ServiceContext{
		Config: c,
		Repo:   repo,
	}
	if hasRPCClientConfig(c.Authsession) {
		sc.AuthsessionClient = authsessionclient.NewAuthsessionClient(authsessionclient.MustNewKitexClient(c.Authsession))
	}
	if closer, ok := any(repo).(interface{ Close() error }); ok {
		sc.closers = append(sc.closers, closer)
	}

	if c.ReceiverOperations != nil {
		consumer, err := receiverevent.NewReceiverConsumer(c.ReceiverOperations, repo)
		if err != nil {
			panic(err)
		}
		sc.workers = append(sc.workers, consumer)
		sc.closers = append(sc.closers, consumer)
	}

	if c.PushOutboxWorker.Enabled && c.PushTasks != nil {
		publisher, err := receiverevent.NewPushTaskPublisher(c.PushTasks)
		if err != nil {
			panic(err)
		}
		worker := repository.NewPushOutboxWorker(repo, publisher, repository.PushOutboxWorkerOptions{
			Interval:          time.Duration(c.PushOutboxWorker.PollInterval) * time.Millisecond,
			BatchSize:         c.PushOutboxWorker.BatchSize,
			PublishingTimeout: time.Duration(c.PushOutboxWorker.PublishingTimeoutMs) * time.Millisecond,
		})
		sc.PushOutboxNotifier = worker
		sc.workers = append(sc.workers, worker)
		sc.closers = append(sc.closers, publisher)
	}

	if c.AffectedOutboxWorker.Enabled {
		sc.workers = append(sc.workers, repository.NewAffectedOutboxWorker(repo, repository.AffectedOutboxWorkerOptions{
			Interval:          time.Duration(c.AffectedOutboxWorker.PollIntervalMs) * time.Millisecond,
			BatchSize:         c.AffectedOutboxWorker.BatchSize,
			ProcessingTimeout: time.Duration(c.AffectedOutboxWorker.ProcessingTimeoutMs) * time.Millisecond,
		}))
	}

	if c.PushTaskConsumer != nil {
		authsessionClient := authsessionclient.NewAuthsessionClient(authsessionclient.MustNewKitexClient(c.Authsession))
		gatewayClient := gatewayclient.NewGatewayClient(gatewayclient.MustNewKitexClient(c.Gateway))
		if userProjectionClient == nil {
			userProjectionClient = userclient.NewUserClient(userclient.MustNewKitexClient(c.BizServiceClient))
		}
		if chatProjectionClient == nil {
			bizChatClient := chatclient.NewChatClient(chatclient.MustNewKitexClient(c.BizServiceClient))
			chatProjectionClient = bizChatClient
			repo.SetPeerProjectionClients(userProjectionClient, chatProjectionClient)
		}
		consumer, err := receiverevent.NewPushTaskConsumer(c.PushTaskConsumer, receiverevent.NewPushTaskDispatcher(authsessionClient, gatewayClient, userProjectionClient, chatProjectionClient))
		if err != nil {
			panic(err)
		}
		sc.workers = append(sc.workers, consumer)
		sc.closers = append(sc.closers, consumer)
	}

	if c.DialogSideEffects.Enabled && hasRPCClientConfig(c.BizServiceClient) {
		dialogClient := dialogclient.NewDialogClient(dialogclient.MustNewKitexClient(c.BizServiceClient))
		sc.workers = append(sc.workers, repository.NewDialogSideEffectWorker(repo, dialogClient, repository.DialogSideEffectWorkerOptions{
			Interval:  time.Duration(c.DialogSideEffects.PollIntervalMs) * time.Millisecond,
			BatchSize: c.DialogSideEffects.BatchSize,
		}))
	}

	return sc
}

func (s *ServiceContext) StartWorkers() {
	if s == nil || len(s.workers) == 0 {
		return
	}
	ctx, cancel := context.WithCancel(context.Background())
	s.cancel = cancel
	for _, worker := range s.workers {
		go worker.Run(ctx)
	}
}

func (s *ServiceContext) StopWorkers() {
	if s == nil {
		return
	}
	if s.cancel != nil {
		s.cancel()
	}
	for _, worker := range s.workers {
		worker.Stop()
	}
	for _, worker := range s.workers {
		if waitable, ok := worker.(waitableBackgroundWorker); ok {
			waitable.Wait()
		}
	}
}

func (s *ServiceContext) Close() error {
	if s == nil {
		return nil
	}
	s.StopWorkers()
	var firstErr error
	for _, closer := range s.closers {
		if closer == nil {
			continue
		}
		if err := closer.Close(); err != nil && firstErr == nil {
			firstErr = err
		}
	}
	return firstErr
}

func hasRPCClientConfig(c kitex.RpcClientConf) bool {
	if c.DestService == "" {
		return false
	}
	return len(c.Endpoints) > 0 || c.Target != "" || c.HasEtcd()
}
