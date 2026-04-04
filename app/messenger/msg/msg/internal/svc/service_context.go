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

	inboxclient "github.com/teamgram/teamgram-server/v2/app/messenger/msg/inbox/client"
	"github.com/teamgram/teamgram-server/v2/app/messenger/msg/inbox/inbox"
	"github.com/teamgram/teamgram-server/v2/app/messenger/msg/msg/internal/config"
	idgenclient "github.com/teamgram/teamgram-server/v2/app/service/idgen/client"
	"github.com/teamgram/teamgram-server/v2/app/service/idgen/idgen"
	syncclient "github.com/teamgram/teamgram-server/v2/app/messenger/sync/client"
	"github.com/teamgram/teamgram-server/v2/app/messenger/sync/sync"
	"github.com/teamgram/teamgram-server/v2/pkg/net/kitex"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

type InboxPushClient interface {
	InboxSendUserMessageToInboxV2(ctx context.Context, in *inbox.TLInboxSendUserMessageToInboxV2) (*tg.Void, error)
}

type SyncPushClient interface {
	SyncUpdatesMe(ctx context.Context, in *sync.TLSyncUpdatesMe) (*tg.Void, error)
	SyncUpdatesNotMe(ctx context.Context, in *sync.TLSyncUpdatesNotMe) (*tg.Void, error)
	SyncPushUpdates(ctx context.Context, in *sync.TLSyncPushUpdates) (*tg.Void, error)
}

type IdgenClient interface {
	IdgenNextId(ctx context.Context, in *idgen.TLIdgenNextId) (*tg.Int64, error)
}

type ServiceContext struct {
	Config      config.Config
	InboxClient InboxPushClient
	SyncClient  SyncPushClient
	IdgenClient IdgenClient
}

func NewServiceContext(c config.Config) *ServiceContext {
	ctx := &ServiceContext{
		Config: c,
	}
	if hasClient(c.InboxClient) {
		ctx.InboxClient = inboxclient.NewInboxClient(inboxclient.MustNewKitexClient(c.InboxClient))
	}
	if hasClient(c.SyncClient) {
		ctx.SyncClient = syncclient.NewSyncClient(syncclient.MustNewKitexClient(c.SyncClient))
	}
	if hasClient(c.IdgenClient) {
		ctx.IdgenClient = idgenclient.NewIdgenClient(idgenclient.MustNewKitexClient(c.IdgenClient))
	}
	return ctx
}

func hasClient(c kitex.RpcClientConf) bool {
	return c.DestService != "" || c.Target != "" || len(c.Endpoints) > 0 || c.HasEtcd()
}
