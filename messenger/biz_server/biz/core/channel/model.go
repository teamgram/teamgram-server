/*
 *  Copyright (c) 2018, https://github.com/nebulaim
 *  All rights reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package channel

import (
	"fmt"

	"github.com/golang/glog"
	"github.com/nebula-chat/chatengine/messenger/biz_server/biz/core"
	"github.com/nebula-chat/chatengine/messenger/biz_server/biz/dal/dao"
	"github.com/nebula-chat/chatengine/messenger/biz_server/biz/dal/dao/mysql_dao"
	"github.com/nebula-chat/chatengine/mtproto"
)

type channelsDAO struct {
	*mysql_dao.CommonDAO
	*mysql_dao.UsersDAO
	*mysql_dao.ChannelsDAO
	*mysql_dao.ChannelParticipantsDAO
	*mysql_dao.UsernameDAO
}

type ChannelModel struct {
	dao                   *channelsDAO
	photoCallback         core.PhotoCallback
	notifySettingCallback core.NotifySettingCallback
	dialogCallback        core.DialogCallback
	usernameCallback      core.UsernameCallback
}

func (m *ChannelModel) InstallModel() {
	m.dao.CommonDAO = dao.GetCommonDAO(dao.DB_MASTER)
	m.dao.UsersDAO = dao.GetUsersDAO(dao.DB_MASTER)
	m.dao.ChannelsDAO = dao.GetChannelsDAO(dao.DB_MASTER)
	m.dao.ChannelParticipantsDAO = dao.GetChannelParticipantsDAO(dao.DB_MASTER)
	m.dao.UsernameDAO = dao.GetUsernameDAO(dao.DB_MASTER)
}

func (m *ChannelModel) RegisterCallback(cb interface{}) {
	switch cb.(type) {
	case core.PhotoCallback:
		glog.Info("channelModel - register core.PhotoCallback")
		m.photoCallback = cb.(core.PhotoCallback)
	case core.NotifySettingCallback:
		glog.Info("channelModel - register core.NotifySettingCallback")
		m.notifySettingCallback = cb.(core.NotifySettingCallback)
	case core.DialogCallback:
		glog.Info("channelModel - register core.DialogCallback")
		m.dialogCallback = cb.(core.DialogCallback)
	case core.UsernameCallback:
		glog.Info("channelModel - register core.UsernameCallback")
		m.usernameCallback = cb.(core.UsernameCallback)
	}
}

// GetUsersBySelfAndIDList
func (m *ChannelModel) GetChannelListBySelfAndIDList(selfUserId int32, idList []int32) (chats []*mtproto.Chat) {
	if len(idList) == 0 {
		return []*mtproto.Chat{}
	}

	chats = make([]*mtproto.Chat, 0, len(idList))

	// TODO(@benqi): 性能优化，从DB里一次性取出所有的chatList
	for _, id := range idList {
		chatData, err := m.NewChannelLogicById(id)
		if err != nil {
			glog.Error("getChatListBySelfIDList - not find chat_id: ", id)
			chatEmpty := &mtproto.TLChatEmpty{Data2: &mtproto.Chat_Data{
				Id: id,
			}}
			chats = append(chats, chatEmpty.To_Chat())
		} else {
			chats = append(chats, chatData.ToChannel(selfUserId))
		}
	}

	return
}

/////////////////////////////////////////////////////////////////////////////////////////////
//func (m *ChannelModel) CheckUserName(channelId int32, userName string) bool {
//	do := m.dao.UsernameDAO.SelectByUsername(userName)
//	return do == nil
//}
//
//func (m *ChannelModel) UpdateUsername(channelId int32, username string) error {
//	usernameDO := m.dao.UsernameDAO.SelectByUsername(username)
//	if usernameDO == nil {
//		usernameDO = &dataobject.UsernameDO{
//			PeerType: 4,
//			PeerId:   channelId,
//			Username: username,
//		}
//		m.dao.UsernameDAO.Insert(usernameDO)
//	} else {
//		m.dao.UsernameDAO.UpdateChannelUsername(username, channelId)
//	}
//	return nil
//}

///////////////////////////////////////////////////////////////////////////////////////////
func (m *ChannelModel) GetChannelBySelfID(selfUserId, channelId int32) (chat *mtproto.Chat) {
	channelData, err := m.NewChannelLogicById(channelId)
	if err != nil {
		glog.Error("GetChannelBySelfID - not find chat_id: ", channelId)
		channelEmpty := &mtproto.TLChatEmpty{Data2: &mtproto.Chat_Data{
			Id: channelId,
		}}
		chat = channelEmpty.To_Chat()
	} else {
		chat = channelData.ToChannel(selfUserId)
	}

	return
}

func (m *ChannelModel) GetChannelParticipantIdList(channelId int32) []int32 {
	doList := m.dao.ChannelParticipantsDAO.SelectByChannelId(channelId)
	idList := make([]int32, 0, len(doList))
	for i := 0; i < len(doList); i++ {
		if doList[i].State == 0 {
			idList = append(idList, doList[i].UserId)
		}
	}
	return idList
}

func (m *ChannelModel) GetChannelParticipant(channelId, userId int32) *mtproto.ChannelParticipant {
	do := m.dao.ChannelParticipantsDAO.SelectByUserId(channelId, userId)
	if do == nil {
		err := fmt.Errorf("not find userId in (%d, %d)", channelId, userId)
		glog.Error(err)
		return nil
	}

	channelParticipantData := channelParticipantData{do}
	return channelParticipantData.ToChannelParticipant()
}

func (m *ChannelModel) GetTopMessageListByIdList(idList []int32) (topMessages map[int32]int32) {
	topMessages = make(map[int32]int32)

	if len(idList) > 0 {
		doList := m.dao.ChannelsDAO.SelectByIdList(idList)
		for i := 0; i < len(doList); i++ {
			topMessages[doList[i].Id] = doList[i].TopMessage
		}
	}

	return
}

func init() {
	core.RegisterCoreModel(&ChannelModel{dao: &channelsDAO{}})
}
