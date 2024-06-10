// Copyright 2022 Teamgram Authors
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

package svc

import (
	kafka "github.com/teamgram/marmota/pkg/mq"
	"github.com/teamgram/marmota/pkg/net/rpcx"
	"github.com/teamgram/marmota/pkg/stores/sqlx"
	inbox_client "github.com/teamgram/teamgram-server/app/messenger/msg/inbox/client"
	"github.com/teamgram/teamgram-server/app/messenger/msg/internal/dao"
	"github.com/teamgram/teamgram-server/app/messenger/msg/msg/internal/config"
	"github.com/teamgram/teamgram-server/app/messenger/msg/msg/plugin"
	sync_client "github.com/teamgram/teamgram-server/app/messenger/sync/client"
	chat_client "github.com/teamgram/teamgram-server/app/service/biz/chat/client"
	dialog_client "github.com/teamgram/teamgram-server/app/service/biz/dialog/client"
	user_client "github.com/teamgram/teamgram-server/app/service/biz/user/client"
	username_client "github.com/teamgram/teamgram-server/app/service/biz/username/client"
	idgen_client "github.com/teamgram/teamgram-server/app/service/idgen/client"
	"github.com/teamgram/teamgram-server/pkg/deduplication"
	"github.com/zeromicro/go-zero/core/stores/kv"
	"github.com/zeromicro/go-zero/core/stores/redis"
)

type ServiceContext struct {
	Config config.Config
	plugin.MsgPlugin
	*dao.Dao
}

func NewServiceContext(c config.Config, plugin plugin.MsgPlugin) *ServiceContext {
	db := sqlx.NewMySQL(&c.Mysql)
	svcCtx := &ServiceContext{
		Config:    c,
		MsgPlugin: plugin,
		Dao: &dao.Dao{
			Mysql:              dao.NewMysqlDao(db, c.MessageSharding),
			IDGenClient2:       idgen_client.NewIDGenClient2(rpcx.GetCachedRpcClient(c.IdgenClient)),
			UserClient:         user_client.NewUserClient(rpcx.GetCachedRpcClient(c.UserClient)),
			InboxClient:        inbox_client.NewInboxMqClient(kafka.MustKafkaProducer(c.InboxClient)),
			ChatClient:         chat_client.NewChatClient(rpcx.GetCachedRpcClient(c.ChatClient)),
			SyncClient:         sync_client.NewSyncMqClient(kafka.GetCachedMQClient(c.SyncClient)),
			DialogClient:       dialog_client.NewDialogClient(rpcx.GetCachedRpcClient(c.DialogClient)),
			MessageDeDuplicate: deduplication.NewMessageDeDuplicate(kv.NewStore(c.KV)),
			Redis:              redis.MustNewRedis(c.Redis2),
			UsernameClient:     username_client.NewUsernameClient(rpcx.GetCachedRpcClient(c.UsernameClient)),
		},
	}

	if plugin == nil {
		svcCtx.MsgPlugin = svcCtx.Dao
	}

	return svcCtx
}
