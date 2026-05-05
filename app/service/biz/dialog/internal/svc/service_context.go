// Copyright (c) 2026 The Teamgram Authors. All rights reserved.
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

	userupdatesclient "github.com/teamgram/teamgram-server/v2/app/messenger/userupdates/client"
	"github.com/teamgram/teamgram-server/v2/app/service/biz/dialog/internal/config"
	"github.com/teamgram/teamgram-server/v2/app/service/biz/dialog/internal/repository"
	"github.com/teamgram/teamgram-server/v2/pkg/net/kitex"
)

type DialogRepository interface {
	ListDialogs(ctx context.Context, userID int64, excludePinned bool, folderID int32) ([]repository.DialogRecord, error)
	ListPinnedDialogs(ctx context.Context, userID int64, folderID int32) ([]repository.DialogRecord, error)
	CountDialogs(ctx context.Context, userID int64, excludePinned bool, folderID int32) (int32, error)
	GetDialogByPeer(ctx context.Context, userID int64, peerType int32, peerID int64) (*repository.DialogRecord, error)
	ListDialogsByPeerDialogIDs(ctx context.Context, userID int64, ids []int64) ([]repository.DialogRecord, error)
	BatchGetDialogExtras(ctx context.Context, userID int64, peers []repository.PeerRef) ([]repository.DialogExtrasRecord, error)
	ListDialogFilters(ctx context.Context, userID int64) ([]repository.DialogFilterRecord, error)
	GetDialogFilter(ctx context.Context, userID int64, filterID int32) (*repository.DialogFilterRecord, error)
	GetDialogFilterBySlug(ctx context.Context, userID int64, slug string) (*repository.DialogFilterRecord, error)
	SaveDialogFilter(ctx context.Context, in repository.SaveDialogFilterInput) (*repository.DialogFilterRecord, error)
	DeleteDialogFilter(ctx context.Context, in repository.DeleteDialogFilterInput) error
	ReorderDialogFilters(ctx context.Context, in repository.ReorderDialogFiltersInput) error
	SetPeerWallpaper(ctx context.Context, in repository.PeerWallpaperInput) error
	SetPrivatePeerPolicy(ctx context.Context, in repository.PrivatePeerPolicyInput) (*repository.PrivatePeerPolicyResult, error)
	SaveDraft(ctx context.Context, in repository.SaveDraftInput) (*repository.DraftMutationResult, error)
	ClearDraft(ctx context.Context, in repository.ClearDraftInput) (*repository.DraftMutationResult, error)
	ClearDraftAfterSend(ctx context.Context, in repository.ClearDraftAfterSendInput) (*repository.DraftMutationResult, error)
	ClearAllDrafts(ctx context.Context, in repository.ClearAllDraftsInput) ([]repository.DraftMutationResult, error)
	ListActiveDrafts(ctx context.Context, userID int64) ([]repository.DraftRecord, error)
	ToggleDialogPin(ctx context.Context, in repository.ToggleDialogPinInput) (*repository.PreferenceMutationResult, error)
	ReorderPinnedDialogs(ctx context.Context, in repository.ReorderPinnedDialogsInput) (*repository.PreferenceMutationResult, error)
	EditPeerFolders(ctx context.Context, in repository.EditPeerFoldersInput) (*repository.PreferenceMutationResult, error)
}

type backgroundWorker interface {
	Run(ctx context.Context)
	Stop()
}

type waitableBackgroundWorker interface {
	Wait()
}

type ServiceContext struct {
	Config  config.Config
	Repo    DialogRepository
	workers []backgroundWorker
	closers []interface{ Close() error }
	cancel  context.CancelFunc
}

func NewServiceContext(c config.Config) *ServiceContext {
	repo := repository.NewRepository(c)
	sc := &ServiceContext{
		Config: c,
		Repo:   repo,
	}
	if closer, ok := any(repo).(interface{ Close() error }); ok {
		sc.closers = append(sc.closers, closer)
	}
	if c.DialogOutboxWorkers.Enabled && hasRPCClientConfig(c.Userupdates) {
		updates := userupdatesclient.NewUserupdatesClient(userupdatesclient.MustNewKitexClient(c.Userupdates))
		options := repository.DialogOutboxWorkerOptions{
			BatchSize:    c.DialogOutboxWorkers.BatchSize,
			LeaseSeconds: c.DialogOutboxWorkers.LeaseSeconds,
			PollInterval: time.Duration(c.DialogOutboxWorkers.PollIntervalMs) * time.Millisecond,
		}
		sc.workers = append(sc.workers,
			repository.NewDialogAuthSeqOutboxWorker(repo, updates, options),
			repository.NewDialogPublicUpdateOutboxWorker(repo, updates, options),
		)
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
