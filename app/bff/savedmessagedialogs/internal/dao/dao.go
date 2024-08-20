// Copyright 2024 Teamgram Authors
//  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Author: teamgramio (teamgram.io@gmail.com)
//

package dao

import (
	kafka "github.com/teamgram/marmota/pkg/mq"
	"github.com/teamgram/marmota/pkg/net/rpcx"
	"github.com/teamgram/teamgram-server/app/bff/savedmessagedialogs/internal/config"
	sync_client "github.com/teamgram/teamgram-server/app/messenger/sync/client"
	chat_client "github.com/teamgram/teamgram-server/app/service/biz/chat/client"
	dialog_client "github.com/teamgram/teamgram-server/app/service/biz/dialog/client"
	message_client "github.com/teamgram/teamgram-server/app/service/biz/message/client"
	updates_client "github.com/teamgram/teamgram-server/app/service/biz/updates/client"
	user_client "github.com/teamgram/teamgram-server/app/service/biz/user/client"
)

type Dao struct {
	dialog_client.DialogClient
	user_client.UserClient
	chat_client.ChatClient
	sync_client.SyncClient
	updates_client.UpdatesClient
	message_client.MessageClient
}

func New(c config.Config) *Dao {
	return &Dao{
		DialogClient:  dialog_client.NewDialogClient(rpcx.GetCachedRpcClient(c.DialogClient)),
		UpdatesClient: updates_client.NewUpdatesClient(rpcx.GetCachedRpcClient(c.UpdatesClient)),
		UserClient:    user_client.NewUserClient(rpcx.GetCachedRpcClient(c.UserClient)),
		SyncClient:    sync_client.NewSyncMqClient(kafka.MustKafkaProducer(c.SyncClient)),
		MessageClient: message_client.NewMessageClient(rpcx.GetCachedRpcClient(c.MessageClient)),
		ChatClient:    chat_client.NewChatClient(rpcx.GetCachedRpcClient(c.ChatClient)),
	}
}
