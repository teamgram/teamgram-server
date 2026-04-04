// Copyright (c) 2024 The Teamgooo Authors. All rights reserved.
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

	sessionclient "github.com/teamgram/teamgram-server/v2/app/interface/session/client"
	"github.com/teamgram/teamgram-server/v2/app/interface/session/session"
	"github.com/teamgram/teamgram-server/v2/app/messenger/sync/internal/config"
	statusclient "github.com/teamgram/teamgram-server/v2/app/service/status/client"
	"github.com/teamgram/teamgram-server/v2/app/service/status/status"
	"github.com/teamgram/teamgram-server/v2/pkg/net/kitex"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

type SessionRpcResultClient interface {
	SessionPushUpdatesData(ctx context.Context, in *session.TLSessionPushUpdatesData) (*tg.Bool, error)
	SessionPushSessionUpdatesData(ctx context.Context, in *session.TLSessionPushSessionUpdatesData) (*tg.Bool, error)
	SessionPushRpcResultData(ctx context.Context, in *session.TLSessionPushRpcResultData) (*tg.Bool, error)
}

type StatusQueryClient interface {
	StatusGetUserOnlineSessions(ctx context.Context, in *status.TLStatusGetUserOnlineSessions) (*status.UserSessionEntryList, error)
}

type ServiceContext struct {
	Config        config.Config
	SessionClient SessionRpcResultClient
	StatusClient  StatusQueryClient
}

func NewServiceContext(c config.Config) *ServiceContext {
	ctx := &ServiceContext{
		Config: c,
	}
	if hasClient(c.SessionClient) {
		ctx.SessionClient = sessionclient.NewSessionClient(sessionclient.MustNewKitexClient(c.SessionClient))
	}
	if hasClient(c.StatusClient) {
		ctx.StatusClient = statusclient.NewStatusClient(statusclient.MustNewKitexClient(c.StatusClient))
	}
	return ctx
}

func hasClient(c kitex.RpcClientConf) bool {
	return c.DestService != "" || c.Target != "" || len(c.Endpoints) > 0 || c.HasEtcd()
}
