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

package chat

import (
	"fmt"
	base2 "github.com/nebula-chat/chatengine/pkg/util"
	"github.com/nebula-chat/chatengine/messenger/biz_server/biz/base"
	"github.com/nebula-chat/chatengine/messenger/biz_server/biz/core"
	"github.com/nebula-chat/chatengine/messenger/biz_server/biz/dal/dataobject"
	"github.com/nebula-chat/chatengine/mtproto"
	"math/rand"
	"time"
	"github.com/golang/glog"
)

const (
	kChatParticipant        = 0
	kChatParticipantCreator = 1
	kChatParticipantAdmin   = 2
)

const (
	// kChatParticipantState = 0			//
	kChatParticipantStateLeft = 1		//
	kChatParticipantStateKicked = 2		//
)

type chatLogicData struct {
	chat         *dataobject.ChatsDO
	participants []dataobject.ChatParticipantsDO
	dao          *chatsDAO
	cb           core.PhotoCallback
}

func makeChatParticipantByDO(do *dataobject.ChatParticipantsDO) (participant *mtproto.ChatParticipant) {
	participant = &mtproto.ChatParticipant{Data2: &mtproto.ChatParticipant_Data{
		UserId:    do.UserId,
		// InviterId: do.InviterUserId,
		// Date:      do.InvitedAt,
	}}

	switch do.ParticipantType {
	case kChatParticipant:
		participant.Constructor = mtproto.TLConstructor_CRC32_chatParticipant
		participant.Data2.Date = do.InvitedAt
	case kChatParticipantCreator:
		participant.Constructor = mtproto.TLConstructor_CRC32_chatParticipantCreator
	case kChatParticipantAdmin:
		participant.Constructor = mtproto.TLConstructor_CRC32_chatParticipantAdmin
		participant.Data2.Date = do.InvitedAt
	default:
		panic("chatParticipant type error.")
	}

	return
}

func (m *ChatModel) NewChatLogicById(chatId int32) (chatData *chatLogicData, err error) {
	chatDO := m.dao.ChatsDAO.Select(chatId)
	if chatDO == nil {
		err = mtproto.NewRpcError2(mtproto.TLRpcErrorCodes_CHAT_ID_INVALID)
	} else {
		chatData = &chatLogicData{
			chat: chatDO,
			dao:  m.dao,
			cb:   m.photoCallback,
		}
	}
	return
}

const (
	kCreateChatFlood = 10  // 10s
)

func (m *ChatModel) CreateChat(creatorId int32, userIdList []int32, title string) (chatData *chatLogicData, err error) {
	date := int32(time.Now().Unix())
	// Check FLOOD_WAIT_
	chatDO := m.dao.ChatsDAO.SelectLastCreator(creatorId)
	if chatDO != nil {
		if date - chatDO.Date < kCreateChatFlood {
			err = mtproto.NewFloodWaitX2(int(date - chatDO.Date))
			glog.Error("create error: ", err, ". lastCreate = ", chatDO.Date)
			return nil, err
		}
	}

	// TODO(@benqi): 事务
	chatData = &chatLogicData{
		chat: &dataobject.ChatsDO{
			CreatorUserId: creatorId,
			AccessHash:    rand.Int63(),
			// TODO(@benqi): use message_id is randomid
			// RandomId:         helper.NextSnowflakeId(),
			ParticipantCount: int32(1 + len(userIdList)),
			Title:            title,
			PhotoId:          0,
			Version:          1,
			Date:             int32(time.Now().Unix()),
		},
		participants: make([]dataobject.ChatParticipantsDO, 1+len(userIdList)),
		dao:          m.dao,
		cb:           m.photoCallback,
	}
	chatData.chat.Id = int32(m.dao.ChatsDAO.Insert(chatData.chat))

	chatData.participants = make([]dataobject.ChatParticipantsDO, 1+len(userIdList))
	chatData.participants[0].ChatId = chatData.chat.Id
	chatData.participants[0].UserId = creatorId
	chatData.participants[0].ParticipantType = kChatParticipantCreator
	m.dao.ChatParticipantsDAO.Insert(&chatData.participants[0])

	for i := 0; i < len(userIdList); i++ {
		chatData.participants[i+1].ChatId = chatData.chat.Id
		chatData.participants[i+1].UserId = userIdList[i]
		chatData.participants[i+1].ParticipantType = kChatParticipant
		chatData.participants[i+1].InviterUserId = creatorId
		chatData.participants[i+1].InvitedAt = chatData.chat.Date
		m.dao.ChatParticipantsDAO.Insert(&chatData.participants[i+1])
	}
	return
}

func (this *chatLogicData) GetPhotoId() int64 {
	return this.chat.PhotoId
}

func (this *chatLogicData) GetChatId() int32 {
	return this.chat.Id
}

func (this *chatLogicData) GetVersion() int32 {
	return this.chat.Version
}

func (this *chatLogicData) checkOrLoadChatParticipantList() {
	if len(this.participants) == 0 {
		this.participants = this.dao.ChatParticipantsDAO.SelectList(this.chat.Id)
	}
}

func (this *chatLogicData) MakeMessageService(fromId int32, action *mtproto.MessageAction) *mtproto.Message {
	peer := &base.PeerUtil{
		PeerType: base.PEER_CHAT,
		PeerId:   this.chat.Id,
	}

	message := &mtproto.TLMessageService{Data2: &mtproto.Message_Data{
		Date:   this.chat.Date,
		FromId: fromId,
		ToId:   peer.ToPeer(),
		Action: action,
	}}
	return message.To_Message()
}

func (this *chatLogicData) MakeCreateChatMessage(creatorId int32) *mtproto.Message {
	idList := this.GetChatParticipantIdList()
	action := &mtproto.TLMessageActionChatCreate{Data2: &mtproto.MessageAction_Data{
		Title: this.chat.Title,
		Users: idList,
	}}

	return this.MakeMessageService(creatorId, action.To_MessageAction())
}

func (this *chatLogicData) MakeAddUserMessage(inviterId, chatUserId int32) *mtproto.Message {
	// idList := this.GetChatParticipantIdList()
	action := &mtproto.TLMessageActionChatAddUser{Data2: &mtproto.MessageAction_Data{
		Title: this.chat.Title,
		Users: []int32{chatUserId},
	}}

	return this.MakeMessageService(inviterId, action.To_MessageAction())
}

func (this *chatLogicData) MakeDeleteUserMessage(operatorId, chatUserId int32) *mtproto.Message {
	// idList := this.GetChatParticipantIdList()
	action := &mtproto.TLMessageActionChatDeleteUser{Data2: &mtproto.MessageAction_Data{
		Title:  this.chat.Title,
		UserId: chatUserId,
	}}

	return this.MakeMessageService(operatorId, action.To_MessageAction())
}

func (this *chatLogicData) MakeChatEditTitleMessage(operatorId int32, title string) *mtproto.Message {
	// idList := this.GetChatParticipantIdList()
	action := &mtproto.TLMessageActionChatEditTitle{Data2: &mtproto.MessageAction_Data{
		Title: title,
	}}

	return this.MakeMessageService(operatorId, action.To_MessageAction())
}

func (this *chatLogicData) GetChatParticipantList() []*mtproto.ChatParticipant {
	this.checkOrLoadChatParticipantList()

	participantList := make([]*mtproto.ChatParticipant, 0, len(this.participants))
	for i := 0; i < len(this.participants); i++ {
		if this.participants[i].State == 0 {
			participantList = append(participantList, makeChatParticipantByDO(&this.participants[i]))
		}
	}
	return participantList
}

func (this *chatLogicData) GetChatParticipantIdList() []int32 {
	this.checkOrLoadChatParticipantList()

	idList := make([]int32, 0, len(this.participants))
	for i := 0; i < len(this.participants); i++ {
		if this.participants[i].State == 0 {
			idList = append(idList, this.participants[i].UserId)
		}
	}
	return idList
}

func (this *chatLogicData) 	GetChatParticipants() *mtproto.TLChatParticipants {
	this.checkOrLoadChatParticipantList()

	return &mtproto.TLChatParticipants{Data2: &mtproto.ChatParticipants_Data{
		ChatId:       this.chat.Id,
		Participants: this.GetChatParticipantList(),
		Version:      this.chat.Version,
	}}
}

func (this *chatLogicData) AddChatUser(inviterId, userId int32) error {
	this.checkOrLoadChatParticipantList()

	// TODO(@benqi): check userId exisits
	var founded = -1
	for i := 0; i < len(this.participants); i++ {
		if userId == this.participants[i].UserId {
			if this.participants[i].State == 1 {
				founded = i
			} else {
				return fmt.Errorf("userId exisits")
			}
		}
	}

	var now = int32(time.Now().Unix())

	if founded != -1 {
		this.participants[founded].State = 0
		this.dao.ChatParticipantsDAO.Update(int8(kChatParticipant), inviterId, now, this.participants[founded].Id)
	} else {
		chatParticipant := &dataobject.ChatParticipantsDO{
			ChatId:          this.chat.Id,
			UserId:          userId,
			ParticipantType: kChatParticipant,
			InviterUserId:   inviterId,
			InvitedAt:       now,
			// JoinedAt:        now,
		}
		chatParticipant.Id = int32(this.dao.ChatParticipantsDAO.Insert(chatParticipant))
		this.participants = append(this.participants, *chatParticipant)
	}

	// update chat
	this.chat.ParticipantCount += 1
	this.chat.Version += 1
	this.chat.Date = now
	this.dao.ChatsDAO.UpdateParticipantCount(this.chat.ParticipantCount, this.chat.Id)

	return nil
}

func (this *chatLogicData) findChatParticipant(selfUserId int32) (int, *dataobject.ChatParticipantsDO) {
	for i := 0; i < len(this.participants); i++ {
		if this.participants[i].UserId == selfUserId {
			return i, &this.participants[i]
		}
	}
	return -1, nil
}

func (this *chatLogicData) ToChat(selfUserId int32) *mtproto.Chat {
	// TODO(@benqi): kicked:flags.1?true left:flags.2?true admins_enabled:flags.3?true admin:flags.4?true deactivated:flags.5?true

	var forbidden = false
	for i := 0; i < len(this.participants); i++ {
		if this.participants[i].UserId == selfUserId && this.participants[i].State == 1 {
			forbidden = true
			break
		}
	}

	if forbidden {
		chat := &mtproto.TLChatForbidden{Data2: &mtproto.Chat_Data{
			Id:    this.chat.Id,
			Title: this.chat.Title,
		}}
		return chat.To_Chat()
	} else {
		chat := &mtproto.TLChat{Data2: &mtproto.Chat_Data{
			Creator:       this.chat.CreatorUserId == selfUserId,
			Id:            this.chat.Id,
			Title:         this.chat.Title,
			AdminsEnabled: this.chat.AdminsEnabled == 1,
			// Photo:             mtproto.NewTLChatPhotoEmpty().To_ChatPhoto(),
			ParticipantsCount: this.chat.ParticipantCount,
			Date:              this.chat.Date,
			Version:           this.chat.Version,
		}}

		if this.chat.PhotoId == 0 {
			chat.SetPhoto(mtproto.NewTLChatPhotoEmpty().To_ChatPhoto())
		} else {
			// sizeList, _ := nbfs_client.GetPhotoSizeList(this.chat.PhotoId)
			chat.SetPhoto(this.cb.GetChatPhoto(this.chat.PhotoId))
		}
		return chat.To_Chat()
	}
}

func (this *chatLogicData) CheckDeleteChatUser(operatorId, deleteUserId int32) error {
	// operatorId is creatorUserId，allow delete all user_id
	// other delete me
	if operatorId != this.chat.CreatorUserId && operatorId != deleteUserId {
		return mtproto.NewRpcError2(mtproto.TLRpcErrorCodes_NO_EDIT_CHAT_PERMISSION)
	}

	this.checkOrLoadChatParticipantList()
	var found = -1
	for i := 0; i < len(this.participants); i++ {
		if deleteUserId == this.participants[i].UserId {
			if this.participants[i].State == 0 {
				found = i
			}
			break
		}
	}

	if found == -1 {
		return mtproto.NewRpcError2(mtproto.TLRpcErrorCodes_PARTICIPANT_NOT_EXISTS)
	}

	return nil
}

func (this *chatLogicData) DeleteChatUser(operatorId, deleteUserId int32) error {
	// operatorId is creatorUserId，allow delete all user_id
	// other delete me
	if operatorId != this.chat.CreatorUserId && operatorId != deleteUserId {
		return mtproto.NewRpcError2(mtproto.TLRpcErrorCodes_NO_EDIT_CHAT_PERMISSION)
	}

	this.checkOrLoadChatParticipantList()
	var found = -1
	for i := 0; i < len(this.participants); i++ {
		if deleteUserId == this.participants[i].UserId {
			if this.participants[i].State == 0 {
				found = i
			}
			break
		}
	}

	var now = int32(time.Now().Unix())
	if found == -1 {
		return mtproto.NewRpcError2(mtproto.TLRpcErrorCodes_PARTICIPANT_NOT_EXISTS)
	}

	this.participants[found].State = 1
	this.dao.ChatParticipantsDAO.UpdateKicked(now, this.participants[found].Id)

	// delete found.
	// this.participants = append(this.participants[:found], this.participants[found+1:]...)

	this.chat.ParticipantCount = int32(len(this.participants) - 1)
	this.chat.Version += 1
	this.chat.Date = now
	this.dao.ChatsDAO.UpdateParticipantCount(this.chat.ParticipantCount, this.chat.Id)

	return nil
}

func (this *chatLogicData) EditChatTitle(editUserId int32, title string) error {
	this.checkOrLoadChatParticipantList()

	_, participant := this.findChatParticipant(editUserId)

	if participant == nil || participant.State == 1 {
		return mtproto.NewRpcError2(mtproto.TLRpcErrorCodes_PARTICIPANT_NOT_EXISTS)
	}

	// check editUserId is creator or admin
	if this.chat.AdminsEnabled != 0 && participant.ParticipantType == kChatParticipant {
		return mtproto.NewRpcError2(mtproto.TLRpcErrorCodes_NO_EDIT_CHAT_PERMISSION)
	}

	if this.chat.Title == title {
		return mtproto.NewRpcError2(mtproto.TLRpcErrorCodes_CHAT_NOT_MODIFIED)
	}

	this.chat.Title = title
	// this.chat.Date = int32(time.Now().Unix())
	this.chat.Version += 1

	this.dao.ChatsDAO.UpdateTitle(title, this.chat.Id)
	return nil
}

func (this *chatLogicData) EditChatPhoto(editUserId int32, photoId int64) error {
	this.checkOrLoadChatParticipantList()

	_, participant := this.findChatParticipant(editUserId)

	if participant == nil || participant.State == 1 {
		return mtproto.NewRpcError2(mtproto.TLRpcErrorCodes_PARTICIPANT_NOT_EXISTS)
	}

	// check editUserId is creator or admin
	if this.chat.AdminsEnabled != 0 && participant.ParticipantType == kChatParticipant {
		return mtproto.NewRpcError2(mtproto.TLRpcErrorCodes_NO_EDIT_CHAT_PERMISSION)
	}

	if this.chat.PhotoId == photoId {
		return mtproto.NewRpcError2(mtproto.TLRpcErrorCodes_CHAT_NOT_MODIFIED)
	}

	this.chat.PhotoId = photoId
	// this.chat.Date = int32(time.Now().Unix())
	this.chat.Version += 1

	this.dao.ChatsDAO.UpdatePhotoId(photoId, this.chat.Id)
	return nil
}

func (this *chatLogicData) EditChatAdmin(operatorId, editChatAdminId int32, isAdmin bool) error {
	// operatorId is creator
	if operatorId != this.chat.CreatorUserId {
		return mtproto.NewRpcError2(mtproto.TLRpcErrorCodes_NO_EDIT_CHAT_PERMISSION)
	}

	// editChatAdminId not creator
	if editChatAdminId == operatorId {
		return mtproto.NewRpcError2(mtproto.TLRpcErrorCodes_BAD_REQUEST)
	}

	this.checkOrLoadChatParticipantList()

	// check exists
	_, participant := this.findChatParticipant(editChatAdminId)
	if participant == nil || participant.State == 1 {
		return mtproto.NewRpcError2(mtproto.TLRpcErrorCodes_PARTICIPANT_NOT_EXISTS)
	}

	if isAdmin && participant.ParticipantType == kChatParticipantAdmin || !isAdmin && participant.ParticipantType == kChatParticipant {
		return mtproto.NewRpcError2(mtproto.TLRpcErrorCodes_CHAT_NOT_MODIFIED)
	}

	if isAdmin {
		participant.ParticipantType = kChatParticipantAdmin
	} else {
		participant.ParticipantType = kChatParticipant
	}
	this.dao.ChatParticipantsDAO.UpdateParticipantType(participant.ParticipantType, participant.Id)

	// update version
	// this.chat.Date = int32(time.Now().Unix())
	this.chat.Version += 1
	this.dao.ChatsDAO.UpdateVersion(this.chat.Id)

	return nil
}

func (this *chatLogicData) ToggleChatAdmins(userId int32, adminsEnabled bool) error {
	// check is creator
	if userId != this.chat.CreatorUserId {
		return mtproto.NewRpcError2(mtproto.TLRpcErrorCodes_NO_EDIT_CHAT_PERMISSION)
	}

	var (
		chatAdminsEnabled = this.chat.AdminsEnabled == 1
	)

	// Check modified
	if chatAdminsEnabled == adminsEnabled {
		return mtproto.NewRpcError2(mtproto.TLRpcErrorCodes_CHAT_NOT_MODIFIED)
	}

	this.chat.AdminsEnabled = base2.BoolToInt8(adminsEnabled)
	// this.chat.Date = int32(time.Now().Unix())
	this.chat.Version += 1

	this.dao.ChatsDAO.UpdateAdminsEnabled(this.chat.AdminsEnabled, this.chat.Id)

	return nil
}
