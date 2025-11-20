// Copyright 2025 Teamgram Authors
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
	"github.com/teamgram/marmota/pkg/stores/sqlc"
	"github.com/teamgram/marmota/pkg/stores/sqlx"
	"github.com/teamgram/teamgram-server/app/messenger/msg/msgtransfer/internal/config"
	sync_client "github.com/teamgram/teamgram-server/app/messenger/sync/client"
	dialog_client "github.com/teamgram/teamgram-server/app/service/biz/dialog/client"
	idgen_client "github.com/teamgram/teamgram-server/app/service/idgen/client"
)

type Dao struct {
	*Mysql
	sqlc.CachedConn
	idgen_client.IDGenClient2
	dialog_client.DialogClient
	SyncClient sync_client.SyncClient
}

func New(c config.Config) *Dao {
	db := sqlx.NewMySQL(&c.Mysql)
	return &Dao{
		Mysql:        NewMysqlDao(db, c.MessageSharding),
		CachedConn:   sqlc.NewConn(db, c.Cache),
		IDGenClient2: idgen_client.NewIDGenClient2(rpcx.GetCachedRpcClient(c.IdgenClient)),
		DialogClient: dialog_client.NewDialogClient(rpcx.GetCachedRpcClient(c.DialogClient)),
		SyncClient:   sync_client.NewSyncMqClient(kafka.GetCachedMQClient(c.SyncClient)),
	}
}
