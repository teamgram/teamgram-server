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

package message

import (
	"github.com/golang/glog"
	"github.com/nebula-chat/chatengine/messenger/biz_server/biz/core"
	"github.com/nebula-chat/chatengine/messenger/biz_server/biz/dal/dao"
	"github.com/nebula-chat/chatengine/messenger/biz_server/biz/dal/dao/mysql_dao"
	"github.com/nebula-chat/chatengine/pkg/redis_client"
)

type messagesDAO struct {
	*mysql_dao.MessageDatasDAO
	*mysql_dao.ChannelMessageBoxesDAO
	*mysql_dao.ChatParticipantsDAO
	*mysql_dao.MessageBoxesDAO
	*mysql_dao.ChannelMediaUnreadDAO
	*mysql_dao.ChannelMessagesDAO
	*mysql_dao.UsernameDAO
	*mysql_dao.UnreadMentionsDAO
	*mysql_dao.CommonDAO
	*mysql_dao.UserDialogsDAO
	*mysql_dao.MessageReactDAO
	*mysql_dao.MessageReactDataDAO
}

type MessageModel struct {
	dao *messagesDAO
	*redis_client.RedisPool
	dialogCallback core.DialogCallback
}

func (m *MessageModel) InstallModel() {
	m.dao.MessageDatasDAO = dao.GetMessageDatasDAO(dao.DB_MASTER)
	m.dao.ChannelMessageBoxesDAO = dao.GetChannelMessageBoxesDAO(dao.DB_MASTER)
	m.dao.ChatParticipantsDAO = dao.GetChatParticipantsDAO(dao.DB_MASTER)
	m.dao.MessageBoxesDAO = dao.GetMessageBoxesDAO(dao.DB_MASTER)
	m.dao.ChannelMediaUnreadDAO = dao.GetChannelMediaUnreadDAO(dao.DB_MASTER)
	m.dao.ChannelMessagesDAO = dao.GetChannelMessagesDAO(dao.DB_MASTER)
	m.dao.UsernameDAO = dao.GetUsernameDAO(dao.DB_MASTER)
	m.dao.MessageReactDataDAO = dao.GetMessageReactDataDAO(dao.DB_MASTER)

	m.dao.UnreadMentionsDAO = dao.GetUnreadMentionsDAO(dao.DB_MASTER)
	m.dao.CommonDAO = dao.GetCommonDAO(dao.DB_MASTER)
	m.dao.UserDialogsDAO = dao.GetUserDialogsDAO(dao.DB_MASTER)
	m.RedisPool = redis_client.GetRedisClient("cache")
}

func (m *MessageModel) RegisterCallback(cb interface{}) {
	switch cb.(type) {
	case core.DialogCallback:
		glog.Info("messageModel - register core.DialogCallback")
		m.dialogCallback = cb.(core.DialogCallback)
	}
}

func init() {
	core.RegisterCoreModel(&MessageModel{dao: &messagesDAO{}})
}
