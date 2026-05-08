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

	"github.com/teamgram/teamgram-server/v2/app/messenger/msg/internal/config"
	"github.com/teamgram/teamgram-server/v2/app/messenger/msg/internal/repository"
	receiverevent "github.com/teamgram/teamgram-server/v2/app/messenger/msg/internal/repository/event"
	userupdatesclient "github.com/teamgram/teamgram-server/v2/app/messenger/userupdates/client"
	"github.com/teamgram/teamgram-server/v2/app/messenger/userupdates/userupdates"
	"github.com/teamgram/teamgram-server/v2/pkg/net/kitex"
)

type MsgRepository interface {
	repository.MessageRepository
	repository.MessageSendStateRepository
}

type UserUpdatesClient interface {
	UserupdatesProcessUserOperation(ctx context.Context, in *userupdates.TLUserupdatesProcessUserOperation) (*userupdates.UserOperationResult, error)
	UserupdatesProcessUserOperationWithEffects(ctx context.Context, in *userupdates.TLUserupdatesProcessUserOperationWithEffects) (*userupdates.UserOperationResult, error)
	UserupdatesGetOperationResult(ctx context.Context, in *userupdates.TLUserupdatesGetOperationResult) (*userupdates.UserOperationResult, error)
}

type ServiceContext struct {
	Config            config.Config
	Repo              MsgRepository
	UserUpdates       UserUpdatesClient
	ReceiverPublisher repository.ReceiverOperationPublisher
}

func NewServiceContext(c config.Config) *ServiceContext {
	var updates UserUpdatesClient
	if hasRPCClientConfig(c.Userupdates) {
		updates = userupdatesclient.NewUserupdatesClient(userupdatesclient.MustNewKitexClient(c.Userupdates))
	}
	sc := &ServiceContext{
		Config:      c,
		Repo:        repository.NewRepository(c),
		UserUpdates: updates,
	}
	if c.ReceiverOperations != nil {
		publisher, err := receiverevent.NewKafkaReceiverOperationPublisher(c.ReceiverOperations)
		if err != nil {
			panic(err)
		}
		sc.ReceiverPublisher = publisher
	}
	return sc
}

func (s *ServiceContext) Close() error {
	if s == nil {
		return nil
	}
	var closeErr error
	if s.Repo != nil {
		if closer, ok := s.Repo.(interface{ Close() error }); ok {
			closeErr = closer.Close()
		}
	}
	if s.ReceiverPublisher != nil {
		if closer, ok := s.ReceiverPublisher.(interface{ Close() error }); ok {
			if err := closer.Close(); err != nil && closeErr == nil {
				closeErr = err
			}
		}
	}
	return closeErr
}

func hasRPCClientConfig(c kitex.RpcClientConf) bool {
	if c.DestService == "" {
		return false
	}
	return len(c.Endpoints) > 0 || c.Target != "" || c.HasEtcd()
}
