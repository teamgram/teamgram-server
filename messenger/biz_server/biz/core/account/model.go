// Copyright (c) 2018-present,  NebulaChat Studio (https://nebula.chat).
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

// Author: Benqi (wubenqi@gmail.com)

package account

import (
	"github.com/nebula-chat/chatengine/messenger/biz_server/biz/core"
	"github.com/nebula-chat/chatengine/messenger/biz_server/biz/dal/dao"
	"github.com/nebula-chat/chatengine/messenger/biz_server/biz/dal/dao/mysql_dao"
)

type accountsDAO struct {
	//*mysql_dao.CommonDAO
	*mysql_dao.AuthUsersDAO
	*mysql_dao.UsersDAO
	*mysql_dao.DevicesDAO
	*mysql_dao.UserNotifySettingsDAO
	*mysql_dao.UserPasswordsDAO
	*mysql_dao.UserPrivacysDAO
	*mysql_dao.ReportsDAO
	*mysql_dao.WallPapersDAO
	*mysql_dao.UsernameDAO
}

type AccountModel struct {
	dao *accountsDAO
}

func (m *AccountModel) InstallModel() {
	m.dao.AuthUsersDAO = dao.GetAuthUsersDAO(dao.DB_MASTER)
	m.dao.UsersDAO = dao.GetUsersDAO(dao.DB_MASTER)
	m.dao.DevicesDAO = dao.GetDevicesDAO(dao.DB_MASTER)
	m.dao.UserNotifySettingsDAO = dao.GetUserNotifySettingsDAO(dao.DB_MASTER)
	m.dao.UserPasswordsDAO = dao.GetUserPasswordsDAO(dao.DB_MASTER)
	m.dao.UserPrivacysDAO = dao.GetUserPrivacysDAO(dao.DB_MASTER)
	m.dao.ReportsDAO = dao.GetReportsDAO(dao.DB_MASTER)
	m.dao.WallPapersDAO = dao.GetWallPapersDAO(dao.DB_MASTER)
	m.dao.UsernameDAO = dao.GetUsernameDAO(dao.DB_MASTER)
}

func (m *AccountModel) RegisterCallback(cb interface{}) {
}

func init() {
	core.RegisterCoreModel(&AccountModel{dao: &accountsDAO{}})
}
