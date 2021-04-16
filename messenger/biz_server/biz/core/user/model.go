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

package user

import (
	"github.com/golang/glog"
	"github.com/nebula-chat/chatengine/messenger/biz_server/biz/core"
	"github.com/nebula-chat/chatengine/messenger/biz_server/biz/dal/dao"
	"github.com/nebula-chat/chatengine/messenger/biz_server/biz/dal/dao/mysql_dao"
)

type usersDAO struct {
	*mysql_dao.UsersDAO
	*mysql_dao.UserPresencesDAO
	*mysql_dao.UserContactsDAO
	*mysql_dao.UserDialogsDAO
	*mysql_dao.UserPasswordsDAO
	*mysql_dao.UserBlocksDAO
	*mysql_dao.BotsDAO
	*mysql_dao.BotCommandsDAO
	*mysql_dao.CommonDAO
	*mysql_dao.UserPrivacysDAO
}

type UserModel struct {
	dao              *usersDAO
	contactCallback  core.ContactCallback
	photoCallback    core.PhotoCallback
	usernameCallback core.UsernameCallback
	// notifySettingCallback core.NotifySettingCallback
	accountCallback core.AccountCallback
}

func (m *UserModel) InstallModel() {
	m.dao.UsersDAO = dao.GetUsersDAO(dao.DB_MASTER)
	m.dao.UserPresencesDAO = dao.GetUserPresencesDAO(dao.DB_MASTER)
	m.dao.UserContactsDAO = dao.GetUserContactsDAO(dao.DB_MASTER)
	m.dao.UserDialogsDAO = dao.GetUserDialogsDAO(dao.DB_MASTER)
	m.dao.UserPasswordsDAO = dao.GetUserPasswordsDAO(dao.DB_MASTER)
	m.dao.UserBlocksDAO = dao.GetUserBlocksDAO(dao.DB_MASTER)
	m.dao.CommonDAO = dao.GetCommonDAO(dao.DB_MASTER)
	m.dao.BotsDAO = dao.GetBotsDAO(dao.DB_MASTER)
	m.dao.BotCommandsDAO = dao.GetBotCommandsDAO(dao.DB_MASTER)
	m.dao.UserPrivacysDAO = dao.GetUserPrivacysDAO(dao.DB_MASTER)
}

func (m *UserModel) RegisterCallback(cb interface{}) {

	switch cb.(type) {
	case core.ContactCallback:
		glog.Info("userModel - register core.ContactCallback")
		m.contactCallback = cb.(core.ContactCallback)
	case core.PhotoCallback:
		glog.Info("userModel - register core.PhotoCallback")
		m.photoCallback = cb.(core.PhotoCallback)
	case core.UsernameCallback:
		glog.Info("userModel - register core.UsernameCallback")
		m.usernameCallback = cb.(core.UsernameCallback)
	// case core.NotifySettingCallback:
	// 	glog.Info("userModel - register core.NotifySettingCallback")
	// 	m.notifySettingCallback = cb.(core.NotifySettingCallback)
	case core.AccountCallback:
		glog.Info("userModel - register core.AccountCallback")
		m.accountCallback = cb.(core.AccountCallback)
	}

}

func init() {
	core.RegisterCoreModel(&UserModel{dao: &usersDAO{}})
}
