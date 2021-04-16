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
	"math"
	"math/rand"
	"time"

	"github.com/golang/glog"
	"github.com/nebula-chat/chatengine/messenger/biz_server/biz/base"
	"github.com/nebula-chat/chatengine/messenger/biz_server/biz/core"
	"github.com/nebula-chat/chatengine/messenger/biz_server/biz/dal/dataobject"
	"github.com/nebula-chat/chatengine/mtproto"
	"github.com/nebula-chat/chatengine/pkg/random2"
	base2 "github.com/nebula-chat/chatengine/pkg/util"
)

// type ParticipantType int
//const (
//	kChannelParticipant       	= 0
//	kChannelParticipantSelf   	= 1
//	kChannelParticipantCreator	= 2
//	kChannelParticipantAdmin   	= 3
//	kChannelParticipantBanned  	= 4
//)

type channelData struct {
	*dataobject.ChannelsDO
	Username string
}

type channelLogicData struct {
	channelData
	cacheParticipantsData []channelParticipantData
	dao                   *channelsDAO
	cb                    core.PhotoCallback
	cb2                   core.NotifySettingCallback
	cb3                   core.UsernameCallback
}

//func (m *ChannelModel) MakeChannelLogic(channelId int32) (channelData *channelLogicData) {
//	channelData = &channelLogicData{
//		channel: &dataobject.ChannelsDO{Id: channelId},
//		dao:     m.dao,
//		cb:      m.photoCallback,
//	}
//	return
//}

func (m *ChannelModel) NewChannelLogicById(channelId int32) (channelData2 *channelLogicData, err error) {
	channelDO := m.dao.ChannelsDAO.Select(channelId)
	if channelDO == nil {
		err = mtproto.NewRpcError2(mtproto.TLRpcErrorCodes_CHAT_ID_INVALID)
	} else {
		username := m.usernameCallback.GetChannelUsername(channelId)
		channelData2 = &channelLogicData{
			channelData: channelData{
				Username:   username,
				ChannelsDO: channelDO,
			},
			cacheParticipantsData: make([]channelParticipantData, 0, 2),
			dao:                   m.dao,
			cb:                    m.photoCallback,
			cb2:                   m.notifySettingCallback,
			cb3:                   m.usernameCallback,
		}
	}
	return
}

// chatInviteAlready#5a686d7c chat:Chat = ChatInvite;
// chatInvite#db74f558 flags:# channel:flags.0?true broadcast:flags.1?true public:flags.2?true megagroup:flags.3?true title:string photo:ChatPhoto participants_count:int participants:flags.4?Vector<User> = ChatInvite;
func (m *ChannelModel) NewChannelLogicByLink(link string) (channelData2 *channelLogicData, err error) {
	channelDO := m.dao.ChannelsDAO.SelectByLink(link)
	if channelDO == nil {
		err = mtproto.NewRpcError2(mtproto.TLRpcErrorCodes_INVITE_HASH_INVALID)
	} else {
		var username = ""
		if channelDO.Public == 1 {
			username = m.usernameCallback.GetChannelUsername(channelDO.Id)
		}

		channelData2 = &channelLogicData{
			channelData: channelData{
				Username:   username,
				ChannelsDO: channelDO,
			},
			cacheParticipantsData: make([]channelParticipantData, 0, 2),
			dao:                   m.dao,
			cb:                    m.photoCallback,
			cb2:                   m.notifySettingCallback,
			cb3:                   m.usernameCallback,
		}
	}
	return
}

func (m *ChannelModel) NewChannelLogicByCreateChannel(creatorId int32, broadcast, megagroup bool, title, about string, randomId int64) (*channelLogicData, error) {
	// 1. check broadcast and megagroup
	if broadcast && megagroup || !broadcast && !megagroup {
		glog.Error("broadcast == megagroup")
		err := mtproto.NewRpcError2(mtproto.TLRpcErrorCodes_BAD_REQUEST)
		return nil, err
	}

	// 2. check title
	if title == "" {
		glog.Error("title empty")
		err := mtproto.NewRpcError2(mtproto.TLRpcErrorCodes_NO_CHAT_TITLE)
		return nil, err
	}

	channelDO := &dataobject.ChannelsDO{
		CreatorUserId:    creatorId,
		AccessHash:       rand.Int63(),
		RandomId:         randomId,
		ParticipantCount: 1,
		Title:            title,
		About:            about,
		Link:             random2.RandomAlphanumeric(22),
		PhotoId:          0,
		Broadcast:        base2.BoolToInt8(broadcast),
		Megagroup:        base2.BoolToInt8(megagroup),
		Democracy:        0,
		Signatures:       0,
		Version:          1,
		Date:             int32(time.Now().Unix()),
	}

	// TODO(@benqi): check channelDO.Id
	channelDO.Id = int32(m.dao.ChannelsDAO.Insert(channelDO))

	// TODO(@benqi): if channel is existed.
	participant := &dataobject.ChannelParticipantsDO{
		ChannelId: channelDO.Id,
		UserId:    creatorId,
		IsCreator: 1,
		JoinedAt:  channelDO.Date,
		Date:      channelDO.Date,
	}
	participant.Id = m.dao.ChannelParticipantsDAO.Insert(participant)

	channelParticipantData2 := channelParticipantData{participant}
	channelParticipantData2.ChannelParticipantsDO = participant

	// username := m.usernameCallback.GetChannelUsername(m)
	channelData := &channelLogicData{
		channelData: channelData{
			Username:   "",
			ChannelsDO: channelDO,
		},
		cacheParticipantsData: []channelParticipantData{channelParticipantData2},
		dao:                   m.dao,
		cb:                    m.photoCallback,
		cb2:                   m.notifySettingCallback,
		cb3:                   m.usernameCallback,
	}

	return channelData, nil
}

func (m *channelLogicData) IsDemocracy() bool {
	return m.Democracy == 1
}

func (m *channelLogicData) IsSignatures() bool {
	return m.Signatures == 1
}

func (m *channelLogicData) IsMegagroup() bool {
	return m.Megagroup == 1
}

func (m *channelLogicData) IsChannel() bool {
	return m.Broadcast == 1
}

func (m *channelLogicData) GetPhotoId() int64 {
	return m.PhotoId
}

func (m *channelLogicData) GetChannelId() int32 {
	return m.Id
}

func (m *channelLogicData) GetVersion() int32 {
	return m.Version
}

////// TODO(@benqi): 性能优化
//func (m *channelLogicData) checkUserIsAdministrator(userId int32) bool {
//	if userId == m.channel.CreatorUserId {
//		return true
//	}
//	participantsDO := m.dao.ChannelParticipantsDAO.SelectByUserId(m.channel.Id, userId)
//	return participantsDO.ParticipantType == kChannelParticipantAdmin
//}
//
//func (m *channelLogicData) updateCacheParticipants(participant *dataobject.ChannelParticipantsDO) {
//	for i := 0; i < len(m.cacheParticipants); i++ {
//		if participant.UserId == m.cacheParticipants[i].UserId {
//			m.cacheParticipants[i] = participant
//		}
//	}
//}

func (m *channelLogicData) checkOrLoadChannelParticipant(selfUserId int32) (found *channelParticipantData) {
	for i := 0; i < len(m.cacheParticipantsData); i++ {
		if m.cacheParticipantsData[i].UserId == selfUserId {
			found = &m.cacheParticipantsData[i]
			return
		}
	}

	participantDO := m.dao.ChannelParticipantsDAO.SelectByUserId(m.Id, selfUserId)
	if participantDO != nil {
		found = &channelParticipantData{participantDO}
		m.cacheParticipantsData = append(m.cacheParticipantsData, *found)
	}

	return
}

func (m *channelLogicData) checkOrLoadChannelParticipantList(idList []int32) (participantList []*channelParticipantData) {
	participantList = make([]*channelParticipantData, 0, len(idList))

	missedList := make([]int32, 0, len(idList))
	for _, id := range idList {
		found := false
		for i := 0; i < len(m.cacheParticipantsData); i++ {
			if m.cacheParticipantsData[i].UserId == id {
				participantList = append(participantList, &m.cacheParticipantsData[i])
				found = true
				break
			}
		}
		if !found {
			missedList = append(missedList, id)
		}
	}

	if len(missedList) == 0 {
		return
	}

	missedDOList := m.dao.ChannelParticipantsDAO.SelectByUserIdList(m.Id, missedList)
	for i := 0; i < len(missedList); i++ {
		m.cacheParticipantsData = append(m.cacheParticipantsData, channelParticipantData{&missedDOList[i]})
	}

	return
}

func (m *channelLogicData) reloadAllChannelParticipant() {
	doList := m.dao.ChannelParticipantsDAO.SelectByChannelId(m.Id)
	participantList := make([]channelParticipantData, 0, len(doList))
	for i := 0; i < len(doList); i++ {
		participantList = append(participantList, channelParticipantData{&doList[i]})
	}
	m.cacheParticipantsData = participantList
}

func (m *channelLogicData) MakeCreateChannelMessage(creatorId int32) *mtproto.Message {
	action := &mtproto.TLMessageActionChannelCreate{Data2: &mtproto.MessageAction_Data{
		Title: m.Title,
	}}
	return MakeChannelMessageService(creatorId, m.Id, action.To_MessageAction())
}

func (m *channelLogicData) MakeAddUserMessage(inviterId, channelUserId int32) *mtproto.Message {
	action := &mtproto.TLMessageActionChatAddUser{Data2: &mtproto.MessageAction_Data{
		Title: m.Title,
		Users: []int32{channelUserId},
	}}

	return MakeChannelMessageService(inviterId, m.Id, action.To_MessageAction())
}

func (m *channelLogicData) MakeDeleteUserMessage(operatorId, channelUserId int32) *mtproto.Message {
	action := &mtproto.TLMessageActionChatDeleteUser{Data2: &mtproto.MessageAction_Data{
		Title:  m.Title,
		UserId: channelUserId,
	}}

	return MakeChannelMessageService(operatorId, m.Id, action.To_MessageAction())
}

func (m *channelLogicData) MakeChannelEditTitleMessage(operatorId int32, title string) *mtproto.Message {
	action := &mtproto.TLMessageActionChatEditTitle{Data2: &mtproto.MessageAction_Data{
		Title: title,
	}}

	return MakeChannelMessageService(operatorId, m.Id, action.To_MessageAction())
}

func (m *channelLogicData) GetChannelParticipantList() []*mtproto.ChannelParticipant {
	m.reloadAllChannelParticipant()

	participantList := make([]*mtproto.ChannelParticipant, 0, len(m.cacheParticipantsData))
	for i := 0; i < len(m.cacheParticipantsData); i++ {
		participantList = append(participantList, m.cacheParticipantsData[i].ToChannelParticipant())
	}

	return participantList
}

// TODO(@benqi): hash
func (m *channelLogicData) GetChannelParticipantListRecent(offset, limit, hash int32) []*mtproto.ChannelParticipant {
	m.reloadAllChannelParticipant()

	participantList := make([]*mtproto.ChannelParticipant, 0, len(m.cacheParticipantsData))

	offsetIndex := int32(0)
	count := int32(0)

	for i := 0; i < len(m.cacheParticipantsData); i++ {
		if m.cacheParticipantsData[i].IsLeft() || m.cacheParticipantsData[i].IsKicked() {
			continue
		}

		if offsetIndex < offset {
			offsetIndex++
			continue
		}
		if count > limit {
			break
		}

		participantList = append(participantList, m.cacheParticipantsData[i].ToChannelParticipant())
	}

	return participantList
}

// TODO(@benqi): hash
func (m *channelLogicData) GetChannelParticipantListAdmins(offset, limit, hash int32) []*mtproto.ChannelParticipant {
	m.reloadAllChannelParticipant()

	participantList := make([]*mtproto.ChannelParticipant, 0, len(m.cacheParticipantsData))

	offsetIndex := int32(0)
	count := int32(0)

	for i := 0; i < len(m.cacheParticipantsData); i++ {
		if !m.cacheParticipantsData[i].IsAdmin() {
			continue
		}

		if offsetIndex < offset {
			offsetIndex++
			continue
		}
		if count > limit {
			break
		}

		participantList = append(participantList, m.cacheParticipantsData[i].ToChannelParticipant())
	}

	return participantList
}

// TODO(@benqi): hash
func (m *channelLogicData) GetChannelParticipantListKicked(q string, offset, limit, hash int32) []*mtproto.ChannelParticipant {
	m.reloadAllChannelParticipant()

	participantList := make([]*mtproto.ChannelParticipant, 0, len(m.cacheParticipantsData))

	offsetIndex := int32(0)
	count := int32(0)

	for i := 0; i < len(m.cacheParticipantsData); i++ {
		// TODO(@benqi): filter q

		if !m.cacheParticipantsData[i].IsKicked() {
			continue
		}

		if offsetIndex < offset {
			offsetIndex++
			continue
		}
		if count > limit {
			break
		}

		participantList = append(participantList, m.cacheParticipantsData[i].ToChannelParticipant())
	}

	return participantList
}

// TODO(@benqi): hash
func (m *channelLogicData) GetChannelParticipantListBots(offset, limit, hash int32) []*mtproto.ChannelParticipant {
	m.reloadAllChannelParticipant()

	participantList := make([]*mtproto.ChannelParticipant, 0, len(m.cacheParticipantsData))

	offsetIndex := int32(0)
	count := int32(0)

	for i := 0; i < len(m.cacheParticipantsData); i++ {
		if offsetIndex < offset {
			offsetIndex++
			continue
		}
		if count > limit {
			break
		}

		// participantList = append(participantList, m.cacheParticipantsData[i].ToChannelParticipant())
	}

	return participantList
}

// TODO(@benqi): include kicked??
func (m *channelLogicData) GetChannelParticipantListBanned(q string, offset, limit, hash int32) []*mtproto.ChannelParticipant {
	m.reloadAllChannelParticipant()

	participantList := make([]*mtproto.ChannelParticipant, 0, len(m.cacheParticipantsData))

	offsetIndex := int32(0)
	count := int32(0)

	for i := 0; i < len(m.cacheParticipantsData); i++ {
		// TODO(@benqi): filter q

		if !m.cacheParticipantsData[i].IsBanned() {
			continue
		}

		if offsetIndex < offset {
			offsetIndex++
			continue
		}
		if count > limit {
			break
		}

		participantList = append(participantList, m.cacheParticipantsData[i].ToChannelParticipant())
	}

	return participantList
}

func (m *channelLogicData) GetChannelParticipantListSearch(q string, offset, limit, hash int32) []*mtproto.ChannelParticipant {
	m.reloadAllChannelParticipant()

	participantList := make([]*mtproto.ChannelParticipant, 0, len(m.cacheParticipantsData))

	offsetIndex := int32(0)
	count := int32(0)

	for i := 0; i < len(m.cacheParticipantsData); i++ {
		// TODO(@benqi): filter q

		//if !m.cacheParticipantsData[i].IsBanned() {
		//	continue
		//}

		if offsetIndex < offset {
			offsetIndex++
			continue
		}
		if count > limit {
			break
		}

		// participantList = append(participantList, m.cacheParticipantsData[i].ToChannelParticipant())
	}

	return participantList
}

func (m *channelLogicData) GetChannelParticipantIdList(excludeUserId int32) []int32 {
	m.reloadAllChannelParticipant()

	idList := make([]int32, 0, len(m.cacheParticipantsData))
	for i := 0; i < len(m.cacheParticipantsData); i++ {
		if m.cacheParticipantsData[i].UserId != excludeUserId &&
			!m.cacheParticipantsData[i].IsKicked() &&
			!m.cacheParticipantsData[i].IsLeft() {

			idList = append(idList, m.cacheParticipantsData[i].UserId)
		}
	}

	return idList
}

//func (m *channelLogicData) GetChannelParticipants() *mtproto.TLChannelsChannelParticipants {
//	m.checkOrLoadChannelParticipantList()
//
//	return &mtproto.TLChannelsChannelParticipants{Data2: &mtproto.Channels_ChannelParticipants_Data{
//		// ChatId: this.channel.Id,
//		Participants: m.GetChannelParticipantList(),
//		// Version: this.channel.Version,
//	}}
// }
//

func (m *channelLogicData) GetChannelParticipant(userId int32) *mtproto.ChannelParticipant {
	participant := m.checkOrLoadChannelParticipant(userId)

	if participant == nil {
		return nil
	}
	return participant.ToChannelParticipant()
}

func (m *channelLogicData) GetChannelParticipantListByIdList(idList []int32) []*mtproto.ChannelParticipant {
	cacheList := m.checkOrLoadChannelParticipantList(idList)
	participantList := make([]*mtproto.ChannelParticipant, 0, len(cacheList))

	for i := 0; i < len(cacheList); i++ {
		participantList = append(participantList, cacheList[i].ToChannelParticipant())
	}

	return participantList
}

func (m *channelLogicData) InviteToChannel(inviterId, userId int32) error {
	if inviterId == userId {
		err := mtproto.NewRpcError2(mtproto.TLRpcErrorCodes_USER_ALREADY_PARTICIPANT)
		glog.Errorf("inviteToChannel error - %s: (%d invite %d)", err, inviterId, userId)
		return err
	}

	inviterParticipant := m.checkOrLoadChannelParticipant(inviterId)
	if inviterParticipant == nil {
		err := mtproto.NewRpcError2(mtproto.TLRpcErrorCodes_USER_NOT_PARTICIPANT)
		glog.Errorf("inviteToChannel error - %s: (%d invite %d)", err, inviterId, userId)
		return err
	}

	if !(inviterParticipant.IsCreator() || inviterParticipant.CanInviteUsers()) {
		err := mtproto.NewRpcError2(mtproto.TLRpcErrorCodes_NO_INVITE_CHANNEL_PERMISSION)
		glog.Errorf("inviteToChannel error - %s: (%d invite %d)", err, inviterId, userId)
		return err
	}

	invitedParticipant := m.checkOrLoadChannelParticipant(userId)
	if invitedParticipant != nil {
		if !(invitedParticipant.IsBanned() || invitedParticipant.IsLeft()) {
			err := mtproto.NewRpcError2(mtproto.TLRpcErrorCodes_USER_ALREADY_PARTICIPANT)
			glog.Errorf("inviteToChannel error - %s: (%d invite %d)", err, inviterId, userId)
			return err
		}
	} else {
		//
	}

	// TODO(@benqi): check participant too much.

	var now = int32(time.Now().Unix())
	channelParticipant := &dataobject.ChannelParticipantsDO{
		ChannelId:     m.Id,
		UserId:        userId,
		InviterUserId: inviterId,
		InvitedAt:     now,
		JoinedAt:      now,
	}
	channelParticipant.Id = m.dao.ChannelParticipantsDAO.InsertOrUpdate(channelParticipant)

	m.ParticipantCount += 1
	m.dao.ChannelsDAO.UpdateParticipantCount(m.ParticipantCount, now, m.Id)

	if invitedParticipant != nil {
		invitedParticipant.ChannelParticipantsDO = channelParticipant
	} else {
		m.cacheParticipantsData = append(m.cacheParticipantsData, channelParticipantData{channelParticipant})
	}

	return nil
}

func (m *channelLogicData) JoinChannel(joinId int32) error {
	joinParticipant := m.checkOrLoadChannelParticipant(joinId)
	if joinParticipant != nil {
		if !(joinParticipant.IsBanned() || joinParticipant.IsLeft()) {
			err := mtproto.NewRpcError2(mtproto.TLRpcErrorCodes_USER_ALREADY_PARTICIPANT)
			glog.Errorf("joinChannel error - %s: (join %d)", err, joinId)
			return err
		}
	} else {
		//
	}

	// TODO(@benqi): check participant too much.

	var now = int32(time.Now().Unix())
	channelParticipant := &dataobject.ChannelParticipantsDO{
		ChannelId:     m.Id,
		UserId:        joinId,
		InviterUserId: m.CreatorUserId,
		InvitedAt:     now,
		JoinedAt:      now,
	}
	channelParticipant.Id = m.dao.ChannelParticipantsDAO.InsertOrUpdate(channelParticipant)

	m.ParticipantCount += 1
	m.dao.ChannelsDAO.UpdateParticipantCount(m.ParticipantCount, now, m.Id)

	if joinParticipant != nil {
		joinParticipant.ChannelParticipantsDO = channelParticipant
	} else {
		m.cacheParticipantsData = append(m.cacheParticipantsData, channelParticipantData{channelParticipant})
	}

	return nil
}

func (m *channelLogicData) LeaveChannel(userId int32) error {
	leftParticipant := m.checkOrLoadChannelParticipant(userId)
	if leftParticipant == nil {
		err := mtproto.NewRpcError2(mtproto.TLRpcErrorCodes_USER_NOT_PARTICIPANT)
		glog.Errorf("leaveChannel error - %s: (%d)", err, userId)
		return err
	}

	if leftParticipant.IsLeft() {
		err := mtproto.NewRpcError2(mtproto.TLRpcErrorCodes_USER_LEFT_CHAT)
		glog.Errorf("leaveChannel error - %s: (%d)", err, userId)
		return err
	}

	if leftParticipant.IsBanned() {
		err := mtproto.NewRpcError2(mtproto.TLRpcErrorCodes_USER_BANNED_IN_CHANNEL)
		glog.Errorf("leaveChannel error - %s: (%d)", err, userId)
		return err
	}

	var now = int32(time.Now().Unix())
	leftParticipant.ChannelParticipantsDO.IsLeft = 1
	leftParticipant.ChannelParticipantsDO.LeftAt = now
	leftParticipant.ChannelParticipantsDO.Date = now

	m.dao.ChannelParticipantsDAO.UpdateLeave(now, m.Id, userId)

	m.ParticipantCount -= 1
	m.dao.ChannelsDAO.UpdateParticipantCount(m.ParticipantCount, now, m.Id)
	return nil
}

//func (m *channelLogicData) CheckDeleteChannelUser(operatorId, deleteUserId int32) error {
//	// operatorId is creatorUserId，allow delete all user_id
//	// other delete me
//	if operatorId != m.channel.CreatorUserId && operatorId != deleteUserId {
//		return mtproto.NewRpcError2(mtproto.TLRpcErrorCodes_NO_EDIT_CHAT_PERMISSION)
//	}
//
//	m.checkOrLoadChannelParticipantList()
//	var found = -1
//	for i := 0; i < len(m.participants); i++ {
//		if deleteUserId == m.participants[i].UserId {
//			if m.participants[i].State == 0 {
//				found = i
//			}
//			break
//		}
//	}
//
//	if found == -1 {
//		return mtproto.NewRpcError2(mtproto.TLRpcErrorCodes_PARTICIPANT_NOT_EXISTS)
//	}
//
//	return nil
//}
//
//func (m *channelLogicData) DeleteChannelUser(operatorId, deleteUserId int32) error {
//	// operatorId is creatorUserId，allow delete all user_id
//	// other delete me
//	if operatorId != m.channel.CreatorUserId && operatorId != deleteUserId {
//		return mtproto.NewRpcError2(mtproto.TLRpcErrorCodes_NO_EDIT_CHAT_PERMISSION)
//	}
//
//	m.checkOrLoadChannelParticipantList()
//	var found = -1
//	for i := 0; i < len(m.participants); i++ {
//		if deleteUserId == m.participants[i].UserId {
//			if m.participants[i].State == 0 {
//				found = i
//			}
//			break
//		}
//	}
//
//	if found == -1 {
//		return mtproto.NewRpcError2(mtproto.TLRpcErrorCodes_PARTICIPANT_NOT_EXISTS)
//	}
//
//	m.participants[found].State = 1
//	m.dao.ChannelParticipantsDAO.DeleteChannelUser(m.participants[found].Id)
//
//	// delete found.
//	// this.participants = append(this.participants[:found], this.participants[found+1:]...)
//
//	var now = int32(time.Now().Unix())
//	m.channel.ParticipantCount = int32(len(m.participants) - 1)
//	m.channel.Version += 1
//	m.channel.Date = now
//	m.dao.ChannelsDAO.UpdateParticipantCount(m.channel.ParticipantCount, now, m.channel.Id)
//
//	return nil
//}
//
func (m *channelLogicData) EditTitle(editUserId int32, title string) error {
	if title == "" {
		err := mtproto.NewRpcError2(mtproto.TLRpcErrorCodes_NO_CHAT_TITLE)
		glog.Errorf("editChannelTitle error - %s: (%d - %s)", err, editUserId, title)
		return err
	}

	if m.Title == title {
		err := mtproto.NewRpcError2(mtproto.TLRpcErrorCodes_CHAT_NOT_MODIFIED)
		glog.Errorf("editChannelTitle error - %s: (%d - %s)", err, editUserId, title)
		return err
	}

	editParticipant := m.checkOrLoadChannelParticipant(editUserId)

	if editParticipant == nil {
		err := mtproto.NewRpcError2(mtproto.TLRpcErrorCodes_USER_NOT_PARTICIPANT)
		glog.Errorf("editChannelTitle error - %s: (%d - %s)", err, editUserId, title)
		return err
	}

	if !(editParticipant.IsCreator() || editParticipant.CanChangeInfo()) {
		err := mtproto.NewRpcError2(mtproto.TLRpcErrorCodes_NO_EDIT_CHAT_PERMISSION)
		glog.Errorf("editChannelTitle error - %s: (%d - %s)", err, editUserId, title)
		return err
	}

	m.Title = title
	m.Date = int32(time.Now().Unix())
	// m.channel.Version += 1

	m.dao.ChannelsDAO.UpdateTitle(title, m.Date, m.Id)
	return nil
}

func (m *channelLogicData) EditAbout(aboutUserId int32, about string) error {
	if m.About == about {
		err := mtproto.NewRpcError2(mtproto.TLRpcErrorCodes_CHAT_NOT_MODIFIED)
		glog.Errorf("editChannelAbout error - %s: (%d - %s)", err, aboutUserId, about)
		return err
	}

	editAboutParticipant := m.checkOrLoadChannelParticipant(aboutUserId)

	if editAboutParticipant == nil {
		err := mtproto.NewRpcError2(mtproto.TLRpcErrorCodes_USER_NOT_PARTICIPANT)
		glog.Errorf("editChannelAbout error - %s: (%d - %s)", err, aboutUserId, about)
		return err
	}

	if !(editAboutParticipant.IsCreator() || editAboutParticipant.CanChangeInfo()) {
		err := mtproto.NewRpcError2(mtproto.TLRpcErrorCodes_NO_EDIT_CHAT_PERMISSION)
		glog.Errorf("editChannelAbout error - %s: (%d - %s)", err, aboutUserId, about)
		return err
	}

	m.About = about
	m.Date = int32(time.Now().Unix())
	// m.channel.Version += 1

	m.dao.ChannelsDAO.UpdateAbout(about, m.Date, m.Id)
	return nil
}

func (m *channelLogicData) EditPhoto(editUserId int32, photoId int64) error {
	editParticipant := m.checkOrLoadChannelParticipant(editUserId)

	if editParticipant == nil {
		err := mtproto.NewRpcError2(mtproto.TLRpcErrorCodes_USER_NOT_PARTICIPANT)
		glog.Errorf("editChannelPhoto error - %s: (%d - %d)", err, editUserId, photoId)
		return err
	}

	if !(editParticipant.IsCreator() || editParticipant.CanChangeInfo()) {
		err := mtproto.NewRpcError2(mtproto.TLRpcErrorCodes_NO_EDIT_CHAT_PERMISSION)
		glog.Errorf("editChannelPhoto error - %s: (%d - %d)", err, editUserId, photoId)
		return err
	}

	m.PhotoId = photoId
	m.Date = int32(time.Now().Unix())
	// m.channel.Version += 1

	m.dao.ChannelsDAO.UpdatePhotoId(photoId, m.Date, m.Id)

	return nil
}

func (m *channelLogicData) EditAdminRights(operatorId, editChannelAdminsId int32, adminRights *mtproto.ChannelAdminRights) error {
	// editChatAdminId not creator
	if editChannelAdminsId == operatorId {
		err := mtproto.NewRpcError2(mtproto.TLRpcErrorCodes_BAD_REQUEST)
		glog.Errorf("editChannelAdminRights error - %s: (%d - %d - %v)", err, operatorId, editChannelAdminsId, adminRights)
		return err
	}

	//
	operatorParticipant := m.checkOrLoadChannelParticipant(operatorId)
	if operatorParticipant == nil {
		err := mtproto.NewRpcError2(mtproto.TLRpcErrorCodes_USER_NOT_PARTICIPANT)
		glog.Errorf("editChannelAdminRights error - %s: (%d - %d - %v)", err, operatorId, editChannelAdminsId, adminRights)
		return err
	}

	if !(operatorParticipant.IsCreator() || operatorParticipant.CanAddAdmins()) {
		err := mtproto.NewRpcError2(mtproto.TLRpcErrorCodes_NO_EDIT_CHAT_PERMISSION)
		glog.Errorf("editChannelAdminRights error - %s: (%d - %d - %v)", err, operatorId, editChannelAdminsId, adminRights)
		return err
	}

	//
	editChannelAdminsParticipant := m.checkOrLoadChannelParticipant(editChannelAdminsId)
	if editChannelAdminsParticipant == nil {
		err := mtproto.NewRpcError2(mtproto.TLRpcErrorCodes_USER_NOT_PARTICIPANT)
		glog.Errorf("editChannelAdminRights error - %s: (%d - %d - %v)", err, operatorId, editChannelAdminsId, adminRights)
		return err
	}

	if editChannelAdminsParticipant.IsLeft() {
		err := mtproto.NewRpcError2(mtproto.TLRpcErrorCodes_USER_LEFT_CHAT)
		glog.Errorf("editChannelAdminRights error - %s: (%d - %d - %v)", err, operatorId, editChannelAdminsId, adminRights)
		return err
	}

	//if editChannelAdminsParticipant.IsKicked() {
	//	err := mtproto.NewRpcError2(mtproto.TLRpcErrorCodes_USER_KICKED)
	//	glog.Errorf("editChannelAdminRights error - %s: (%d - %d - %v)", err, operatorId, editChannelAdminsId, adminRights)
	//	return err
	//}

	if editChannelAdminsParticipant.IsCreator() {
		// creator??
		glog.Warning("editChannelAdminRights - edit creator: (%d - %d - %v)", operatorId, editChannelAdminsId, adminRights)
	} else {
		now := int32(time.Now().Unix())
		editChannelAdminsParticipant.AdminRights = FromChannelAdminRights(adminRights.To_ChannelAdminRights())
		editChannelAdminsParticipant.PromotedBy = operatorId
		editChannelAdminsParticipant.PromotedBy = now
		// editChannelAdminsParticipant.BannedUntilDate = 0
		// editChannelAdminsParticipant.BannedRights = 0
		editChannelAdminsParticipant.Date = now

		m.dao.ChannelParticipantsDAO.UpdateAdminRights(editChannelAdminsParticipant.AdminRights, m.Id, editChannelAdminsId)
	}
	return nil
}

func (m *channelLogicData) EditBanned(operatorId, bannedUserId int32, bannedRights *mtproto.ChannelBannedRights) error {
	// editChatAdminId not creator
	if bannedUserId == operatorId {
		err := mtproto.NewRpcError2(mtproto.TLRpcErrorCodes_BAD_REQUEST)
		glog.Errorf("editBanned error - %s: (%d - %d - %v)", err, operatorId, bannedUserId, bannedRights)
		return err
	}

	//
	operatorParticipant := m.checkOrLoadChannelParticipant(operatorId)
	if operatorParticipant == nil {
		err := mtproto.NewRpcError2(mtproto.TLRpcErrorCodes_USER_NOT_PARTICIPANT)
		glog.Errorf("editBanned error - %s: (%d - %d - %v)", err, operatorId, bannedUserId, bannedRights)
		return err
	}

	if !(operatorParticipant.IsCreator() || operatorParticipant.CanBanUsers()) {
		err := mtproto.NewRpcError2(mtproto.TLRpcErrorCodes_NO_EDIT_CHAT_PERMISSION)
		glog.Errorf("editBanned error - %s: (%d - %d - %v)", err, operatorId, bannedUserId, bannedRights)
		return err
	}

	//
	bannedParticipant := m.checkOrLoadChannelParticipant(bannedUserId)
	if bannedParticipant == nil {
		err := mtproto.NewRpcError2(mtproto.TLRpcErrorCodes_USER_NOT_PARTICIPANT)
		glog.Errorf("editBanned error - %s: (%d - %d - %v)", err, operatorId, bannedUserId, bannedRights)
		return err
	}

	if bannedParticipant.IsLeft() {
		err := mtproto.NewRpcError2(mtproto.TLRpcErrorCodes_USER_LEFT_CHAT)
		glog.Errorf("editBanned error - %s: (%d - %d - %v)", err, operatorId, bannedUserId, bannedRights)
		return err
	}

	now := int32(time.Now().Unix())

	// kicked
	r, d := FromChannelBannedRights(bannedRights.To_ChannelBannedRights())
	if r == 255 && d == math.MaxInt32 {
		if !bannedParticipant.IsKicked() {
			m.ParticipantCount -= 1
			bannedParticipant.ChannelParticipantsDO.IsKicked = 1
			m.dao.ChannelsDAO.UpdateParticipantCount(m.ParticipantCount, int32(time.Now().Unix()), m.Id)
		}
	}

	bannedParticipant.KickedBy = operatorId
	bannedParticipant.KickedAt = now

	m.dao.ChannelParticipantsDAO.UpdateBannedRights(r, d, m.Id, bannedUserId)
	return nil
}

func (m *channelLogicData) ExportedChatInvite() string {
	if m.Link == "" {
		// TODO(@benqi): 检查唯一性
		m.Link = random2.RandomAlphanumeric(22)
		/// m.Link = "https://nebula.im/joinchat/" + base64.StdEncoding.EncodeToString(crypto.GenerateNonce(16))
		m.dao.ChannelsDAO.UpdateLink(m.Link, int32(time.Now().Unix()), m.Id)
	}
	return "https://t.me/joinchat/" + m.Link
}

func (m *channelLogicData) ToggleSignatures(operatorId int32, enabled bool) error {
	editParticipant := m.checkOrLoadChannelParticipant(operatorId)

	if editParticipant == nil {
		err := mtproto.NewRpcError2(mtproto.TLRpcErrorCodes_USER_NOT_PARTICIPANT)
		glog.Errorf("toggleSignatures error - %s: (%d - %v)", err, operatorId, enabled)
		return err
	}

	if !(editParticipant.IsCreator() || editParticipant.CanChangeInfo()) {
		err := mtproto.NewRpcError2(mtproto.TLRpcErrorCodes_NO_EDIT_CHAT_PERMISSION)
		glog.Errorf("toggleSignatures error - %s: (%d - %v)", err, operatorId, enabled)
		return err
	}

	m.Signatures = base2.BoolToInt8(enabled)
	m.dao.ChannelsDAO.UpdateSignatures(m.Signatures, m.Date, m.Id)

	return nil
}

func (m *channelLogicData) ToggleInvites(operatorId int32, enabled bool) error {
	editParticipant := m.checkOrLoadChannelParticipant(operatorId)

	if editParticipant == nil {
		err := mtproto.NewRpcError2(mtproto.TLRpcErrorCodes_USER_NOT_PARTICIPANT)
		glog.Errorf("toggleSignatures error - %s: (%d - %v)", err, operatorId, enabled)
		return err
	}

	if !(editParticipant.IsCreator() || editParticipant.CanChangeInfo()) {
		err := mtproto.NewRpcError2(mtproto.TLRpcErrorCodes_NO_EDIT_CHAT_PERMISSION)
		glog.Errorf("toggleSignatures error - %s: (%d - %v)", err, operatorId, enabled)
		return err
	}

	m.Democracy = base2.BoolToInt8(enabled)
	m.dao.ChannelsDAO.UpdateDemocracy(m.Democracy, m.Date, m.Id)

	return nil
}

func (m *channelLogicData) UpdateUsername(operatorId int32, username string, cb func(int32, string) bool) error {
	editParticipant := m.checkOrLoadChannelParticipant(operatorId)

	if editParticipant == nil {
		err := mtproto.NewRpcError2(mtproto.TLRpcErrorCodes_USER_NOT_PARTICIPANT)
		glog.Errorf("updateUsername error - %s: (%d - %s)", err, operatorId, username)
		return err
	}

	if !(editParticipant.IsCreator() || editParticipant.CanChangeInfo()) {
		err := mtproto.NewRpcError2(mtproto.TLRpcErrorCodes_NO_EDIT_CHAT_PERMISSION)
		glog.Errorf("updateUsername error - %s: (%d - %s)", err, operatorId, username)
		return err
	}

	cb(m.Id, username)
	//if cb(m.Id, username) {
	//	//if username == "" {
	//	//	m.Democracy = 0
	//	//} else {
	//	//	m.Democracy = 1
	//	//}
	//	m.dao.ChannelsDAO.UpdateDemocracy(m.Democracy, m.Date, m.Id)
	//}
	return nil
}

func (m *channelLogicData) SetTopMessage(topMessage int32) (err error) {
	m.dao.ChannelsDAO.UpdateTopMessage(topMessage, m.Id)
	return
}

func (m *channelLogicData) ReadOutboxHistory(userId, maxId int32) bool {
	affected := m.dao.ChannelParticipantsDAO.UpdateReadInboxMaxId(maxId, m.Id, userId)
	return affected > 0
}

func (m *channelLogicData) ToChannel(selfUserId int32) *mtproto.Chat {
	selfParticipant := m.checkOrLoadChannelParticipant(selfUserId)
	if selfParticipant == nil || selfParticipant.IsKicked() {
		return m.ToChannelForbidden()
	}

	channel := &mtproto.TLChannel{Data2: &mtproto.Chat_Data{
		Creator:      selfParticipant.IsCreator(),
		Left:         selfParticipant.IsLeft(),
		Broadcast:    m.IsChannel(),
		Megagroup:    m.IsMegagroup(),
		Democracy:    m.IsDemocracy(),
		Signatures:   m.IsSignatures(),
		Id:           m.Id,
		AccessHash:   m.AccessHash,
		Title:        m.Title,
		Photo:        m.getChatPhoto(),
		Username:     m.Username,
		Date:         m.Date,
		Version:      m.Version,
		AdminRights:  selfParticipant.ToChannelAdminRights(),
		BannedRights: selfParticipant.ToChannelBannedRights(),
	}}

	// if m.IsDemocracy() {
	// channel.SetUsername(m.cb3.GetChannelUsername(m.Id))
	// }

	return channel.To_Chat()
}

func (m *channelLogicData) ToChannelForbidden() *mtproto.Chat {
	channelForbidden := &mtproto.TLChannelForbidden{Data2: &mtproto.Chat_Data{
		Broadcast:  base2.Int8ToBool(m.Broadcast),
		Megagroup:  base2.Int8ToBool(m.Megagroup),
		Id:         m.Id,
		AccessHash: m.AccessHash,
		Title:      m.Title,
	}}
	return channelForbidden.To_Chat()
}

func (m *channelLogicData) ToChannelFull(selfUserId int32) *mtproto.ChatFull {
	peer := &base.PeerUtil{
		PeerType: base.PEER_CHANNEL,
		PeerId:   m.Id,
	}

	var notifySettings *mtproto.PeerNotifySettings

	// TODO(@benqi): chat notifySetting...
	if m.cb2 == nil {
		notifySettings = &mtproto.PeerNotifySettings{
			Constructor: mtproto.TLConstructor_CRC32_peerNotifySettings,
			Data2: &mtproto.PeerNotifySettings_Data{
				ShowPreviews: mtproto.ToBool(true),
				Silent:       mtproto.ToBool(false),
				MuteUntil:    0,
				Sound:        "default",
			},
		}
	} else {
		notifySettings = m.cb2.GetNotifySettings(selfUserId, peer)
	}

	channelFull := &mtproto.TLChannelFull{Data2: &mtproto.ChatFull_Data{
		// CanViewParticipants:
		Id:                m.Id,
		About:             m.About,
		ParticipantsCount: m.ParticipantCount,
		AdminsCount:       1, // TODO(@benqi): calc adminscount
		ChatPhoto:         m.getPhoto(),
		NotifySettings:    notifySettings,
		ExportedInvite:    mtproto.NewTLChatInviteEmpty().To_ExportedChatInvite(), // TODO(@benqi):
		BotInfo:           []*mtproto.BotInfo{},
	}}

	selfParticipant := m.checkOrLoadChannelParticipant(selfUserId)
	if selfParticipant.IsAdmin() {
		channelFull.SetCanViewParticipants(true)
		channelFull.SetCanSetUsername(true)
	}

	exportedInvite := &mtproto.TLChatInviteExported{Data2: &mtproto.ExportedChatInvite_Data{
		Link: "https://t.me/joinchat/" + m.Link,
	}}

	channelFull.SetExportedInvite(exportedInvite.To_ExportedChatInvite())
	return channelFull.To_ChatFull()
}

func (m *channelLogicData) ToChatInvite(userId int32, cb func([]int32) []*mtproto.User) *mtproto.ChatInvite {
	var chatInvite *mtproto.ChatInvite
	invitedParticipant := m.checkOrLoadChannelParticipant(userId)
	if invitedParticipant != nil {
		_chatInviteAlready := &mtproto.TLChatInviteAlready{Data2: &mtproto.ChatInvite_Data{
			Chat: m.ToChannel(userId),
		}}
		chatInvite = _chatInviteAlready.To_ChatInvite()
	} else {
		_chatInvite := &mtproto.TLChatInvite{Data2: &mtproto.ChatInvite_Data{
			Channel:           true,
			Broadcast:         m.Broadcast == 1,
			Public:            m.Public == 1,
			Megagroup:         m.IsMegagroup(),
			Title:             m.Title,
			Photo:             m.getChatPhoto(),
			ParticipantsCount: m.ParticipantCount,
		}}

		if cb != nil {
			participants := m.GetChannelParticipantListRecent(0, 5, 0)
			idList := []int32{m.CreatorUserId}
			for _, p := range participants {
				if p.GetData2().GetUserId() != m.CreatorUserId {
					idList = append(idList, p.GetData2().GetUserId())
				}
			}
			_chatInvite.SetParticipants(cb(idList))
		} else {
			_chatInvite.SetParticipants(make([]*mtproto.User, 0))
		}

		chatInvite = _chatInvite.To_ChatInvite()
	}
	return chatInvite
}

func (m *channelLogicData) getPhoto() *mtproto.Photo {
	// TODO(@benqi): nbfs_client
	var photo *mtproto.Photo

	// TODO(@benqi):
	if m.GetPhotoId() == 0 {
		photoEmpty := &mtproto.TLPhotoEmpty{Data2: &mtproto.Photo_Data{
			Id: 0,
		}}
		photo = photoEmpty.To_Photo()
	} else {
		// sizes, _ := nbfs_client.GetPhotoSizeList(channelData.channel.PhotoId)
		// photo2 := photo2.MakeUserProfilePhoto(photoId, sizes)
		//channelPhoto := &mtproto.TLPhoto{ Data2: &mtproto.Photo_Data{
		//	Id:          channelData.channel.PhotoId,
		//	HasStickers: false,
		//	AccessHash:  channelData.channel.PhotoId, // photo2.GetFileAccessHash(file.GetData2().GetId(), file.GetData2().GetParts()),
		//	Date:        int32(time.Now().Unix()),
		//	Sizes:       sizes,
		//}}
		photo = m.cb.GetPhoto(m.PhotoId)
		// channelPhoto.To_Photo()
	}
	return photo
}

func (m *channelLogicData) getChatPhoto() *mtproto.ChatPhoto {
	var chatPhoto *mtproto.ChatPhoto
	if m.PhotoId == 0 {
		chatPhoto = mtproto.NewTLChatPhotoEmpty().To_ChatPhoto()
	} else {
		chatPhoto = m.cb.GetChatPhoto(m.PhotoId)
	}
	return chatPhoto
}
