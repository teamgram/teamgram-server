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

package repository

import (
	"context"

	"github.com/teamgram/teamgram-server/v2/app/bff/dialogs/internal/config"
	msgclient "github.com/teamgram/teamgram-server/v2/app/messenger/msg/client"
	syncclient "github.com/teamgram/teamgram-server/v2/app/messenger/sync/client"
	syncpb "github.com/teamgram/teamgram-server/v2/app/messenger/sync/sync"
	userupdatesclient "github.com/teamgram/teamgram-server/v2/app/messenger/userupdates/client"
	chatclient "github.com/teamgram/teamgram-server/v2/app/service/biz/chat/client"
	dialogclient "github.com/teamgram/teamgram-server/v2/app/service/biz/dialog/client"
	userclient "github.com/teamgram/teamgram-server/v2/app/service/biz/user/client"
	"github.com/teamgram/teamgram-server/v2/pkg/net/kitex"
	"github.com/teamgram/teamgram-server/v2/pkg/net/kitex/identity"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

// Repository is the dependency container for repository instances.
type Repository struct {
	ChatClient        chatclient.ChatClient
	DialogClient      dialogclient.DialogClient
	MsgClient         msgclient.MsgClient
	SyncClient        syncclient.SyncClient
	UserupdatesClient userupdatesclient.UserupdatesClient
	UserClient        userclient.UserClient
}

// NewRepository creates a new Repository.
func NewRepository(c config.Config) *Repository {
	r := &Repository{
		ChatClient:        chatclient.NewChatClient(chatclient.MustNewKitexClient(c.ChatClient)),
		DialogClient:      dialogclient.NewDialogClient(dialogclient.MustNewKitexClient(c.DialogClient)),
		MsgClient:         msgclient.NewMsgClient(msgclient.MustNewKitexClient(c.MsgClient)),
		UserupdatesClient: userupdatesclient.NewUserupdatesClient(userupdatesclient.MustNewKitexClient(c.UserupdatesClient)),
		UserClient:        userclient.NewUserClient(userclient.MustNewKitexClient(c.UserClient)),
	}
	if hasSyncClientConfig(c.SyncClient) {
		r.SyncClient = syncclient.NewSyncClient(syncclient.MustNewKitexClient(c.SyncClient))
	}
	return r
}

func hasSyncClientConfig(c kitex.RpcClientConf) bool {
	return c.DestService != "" && c.ServiceName != ""
}

// PushTypingUpdates sends a realtime-only typing update to the sync service.
func (r *Repository) PushTypingUpdates(ctx context.Context, userID int64, updates tg.UpdatesClazz) error {
	ctx = identity.WithCallerService(ctx, "bff.dialogs")
	_, err := r.SyncClient.SyncPushUpdates(ctx, &syncpb.TLSyncPushUpdates{
		UserId:  userID,
		Updates: updates,
	})
	return err
}
