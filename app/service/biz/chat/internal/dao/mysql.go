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
	"github.com/teamgram/marmota/pkg/stores/sqlx"
	"github.com/teamgram/teamgram-server/app/service/biz/chat/internal/dal/dao/mysql_dao"
)

type Mysql struct {
	*sqlx.DB
	*mysql_dao.ChatInviteParticipantsDAO
	*mysql_dao.ChatInvitesDAO
	*mysql_dao.ChatParticipantsDAO
	*mysql_dao.ChatsDAO
	*sqlx.CommonDAO
}

func newMysqlDao(db *sqlx.DB) *Mysql {
	return &Mysql{
		DB:                        db,
		ChatInviteParticipantsDAO: mysql_dao.NewChatInviteParticipantsDAO(db),
		ChatInvitesDAO:            mysql_dao.NewChatInvitesDAO(db),
		ChatParticipantsDAO:       mysql_dao.NewChatParticipantsDAO(db),
		ChatsDAO:                  mysql_dao.NewChatsDAO(db),
		CommonDAO:                 sqlx.NewCommonDAO(db),
	}
}
