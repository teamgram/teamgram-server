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

package dao

import (
	"context"

	"github.com/teamgram/marmota/pkg/stores/sqlx"
	"github.com/teamgram/teamgram-server/app/messenger/sync/internal/dal/dao/mysql_dao"
)

type Mysql struct {
	*sqlx.DB
	*mysql_dao.AuthSeqUpdatesDAO
	*mysql_dao.BotUpdatesDAO
	*mysql_dao.ChannelPtsUpdatesDAO
	*mysql_dao.UserPtsUpdatesDAO
	*mysql_dao.UserQtsUpdatesDAO
	*sqlx.CommonDAO
}

func newMysqlDao(db *sqlx.DB) *Mysql {
	return &Mysql{
		DB:                   db,
		AuthSeqUpdatesDAO:    mysql_dao.NewAuthSeqUpdatesDAO(db),
		BotUpdatesDAO:        mysql_dao.NewBotUpdatesDAO(db),
		ChannelPtsUpdatesDAO: mysql_dao.NewChannelPtsUpdatesDAO(db),
		UserPtsUpdatesDAO:    mysql_dao.NewUserPtsUpdatesDAO(db),
		UserQtsUpdatesDAO:    mysql_dao.NewUserQtsUpdatesDAO(db),
		CommonDAO:            sqlx.NewCommonDAO(db),
	}
}

func (d *Mysql) Close() error {
	return d.DB.Close()
}

func (d *Mysql) Ping(ctx context.Context) (err error) {
	return d.DB.Ping(ctx)
}
