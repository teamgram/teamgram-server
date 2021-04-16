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
	"fmt"
	"time"

	"github.com/gogo/protobuf/proto"
	"github.com/golang/glog"
	"github.com/nebula-chat/chatengine/messenger/biz_server/biz/base"
	"github.com/nebula-chat/chatengine/messenger/biz_server/biz/core"
	"github.com/nebula-chat/chatengine/messenger/biz_server/biz/dal/dataobject"
	"github.com/nebula-chat/chatengine/messenger/sync/sync_client"
	"github.com/nebula-chat/chatengine/mtproto"
	base2 "github.com/nebula-chat/chatengine/pkg/util"
)

type OnBoxCallback func(int32, *MessageBox2)

type MessageBox2 struct {
	OwnerId        int32
	MessageId      int32
	MessageBoxType int8
	Mentioned      bool
	MediaUnread    bool
	ReplyToMsgId   int32
	*MessageData
}

type MessageBox2List []*MessageBox2

func hasMention(entities []*mtproto.MessageEntity, userId int32) bool {
	for _, e := range entities {
		switch e.GetConstructor() {
		case mtproto.TLConstructor_CRC32_messageEntityMentionName:
			if e.Data2.UserId_5 == userId {
				return true
			}
		case mtproto.TLConstructor_CRC32_messageEntityMention:
			if e.Data2.UserId_5 == userId {
				return true
			}
		}
	}
	return false
}

func MakeMessageOutBox(userId int32, mediaUnread bool, replyToMsgId int32, messageData *MessageData) *MessageBox2 {
	return &MessageBox2{
		OwnerId:        userId,
		MessageId:      int32(core.NextMessageBoxId(userId)),
		MessageBoxType: MESSAGE_BOX_TYPE_OUTGOING,
		MediaUnread:    mediaUnread,
		ReplyToMsgId:   replyToMsgId,
		MessageData:    messageData,
	}
}
func (m *MessageModel) makeChannelMessageBoxByDO(boxDO *dataobject.ChannelMessagesDO) *MessageBox2 {
	// TODO(@benqi): check boxDO and dataDO
	mBox := &MessageBox2{
		OwnerId:        boxDO.ChannelId,
		MessageId:      boxDO.ChannelMessageId,
		MessageBoxType: MESSAGE_BOX_TYPE_CHANNEL,
		// MediaUnread:    base2.Int8ToBool(boxDO.HasMediaUnread),
	}

	mData := &MessageData{
		DialogId:        int64(-boxDO.ChannelId),
		DialogMessageId: boxDO.ChannelMessageId,
		SenderUserId:    boxDO.SenderUserId,
		Peer:            &base.PeerUtil{PeerType: base.PEER_CHANNEL, PeerId: boxDO.ChannelId},
		RandomId:        boxDO.RandomId,
		EditMessage:     boxDO.EditMessage,
		EditDate:        boxDO.EditDate,
		Views:           boxDO.Views,
		dao:             m.dao,
	}

	mData.Message, _ = decodeMessage(int(boxDO.MessageType), []byte(boxDO.MessageData))
	mBox.MessageData = mData
	return mBox
}

func MakeChannelMessageBox(channelId int32, messageData *MessageData) *MessageBox2 {
	return &MessageBox2{
		OwnerId:        channelId,
		MessageId:      int32(messageData.DialogMessageId),
		MessageBoxType: MESSAGE_BOX_TYPE_CHANNEL,
		MediaUnread:    false,
		MessageData:    messageData,
	}
}
func MakeMessageInBox(userId int32, mediaUnread bool, replyToMsgId int32, messageData *MessageData) *MessageBox2 {
	mentioned := hasMention(messageData.Message.Data2.Entities, userId)
	return &MessageBox2{
		OwnerId:        userId,
		MessageId:      int32(core.NextMessageBoxId(userId)),
		MessageBoxType: MESSAGE_BOX_TYPE_INCOMING,
		Mentioned:      mentioned,
		MediaUnread:    mentioned || mediaUnread,
		ReplyToMsgId:   replyToMsgId,
		MessageData:    messageData,
	}
}

func (m *MessageBox2) Insert() (lastInsertId int64) {
	lastInsertId = -1
	switch m.MessageBoxType {
	case MESSAGE_BOX_TYPE_INCOMING, MESSAGE_BOX_TYPE_OUTGOING:
		boxDO := &dataobject.MessageBoxesDO{
			UserId:           m.OwnerId,
			UserMessageBoxId: m.MessageId,
			DialogId:         m.DialogId,
			DialogMessageId:  m.DialogMessageId,
			MessageDataId:    m.MessageDataId,
			MessageBoxType:   m.MessageBoxType,
			ReplyToMsgId:     m.ReplyToMsgId,
			Mentioned:        base2.BoolToInt8(m.Mentioned),
			MediaUnread:      base2.BoolToInt8(m.MediaUnread),
			Date2:            int32(time.Now().Unix()),
			Deleted:          0,
		}
		lastInsertId = m.dao.MessageBoxesDAO.Insert(boxDO)
		if lastInsertId == 0 {
			// m.dao.MentionsDAO.InsertIgnore()
			do := m.dao.MessageBoxesDAO.SelectByMessageDataId(m.MessageDataId)
			if do != nil {
				m.MessageId = do.UserMessageBoxId
			} else {
				lastInsertId = -1
			}
		}
	case MESSAGE_BOX_TYPE_CHANNEL:
		// channel_message inserted.
	}
	return
}

func (m *MessageBox2) ToMessage(toUserId int32) *mtproto.Message {
	message := proto.Clone(m.Message).(*mtproto.Message)

	// fix message_data
	if m.EditDate != 0 {
		message.Data2.Message = m.EditMessage
		message.Data2.EditDate = m.EditDate
	}

	if m.Views != 0 {
		message.Data2.Views = m.Views
	}

	// fix message_box
	switch m.MessageBoxType {
	case MESSAGE_BOX_TYPE_OUTGOING:
		message.Data2.Out = true
		// Outbox media_unread is false
		message.Data2.ReplyToMsgId = m.ReplyToMsgId
		message.Data2.MediaUnread = false
	case MESSAGE_BOX_TYPE_INCOMING:
		message.Data2.Out = false
		message.Data2.ReplyToMsgId = m.ReplyToMsgId
		message.Data2.MediaUnread = m.MediaUnread
	case MESSAGE_BOX_TYPE_CHANNEL:
		if m.SenderUserId == toUserId {
			message.Data2.Out = true
			// Outbox media_unread is false
			message.Data2.MediaUnread = false
		} else {
			message.Data2.Out = false
			if m.HasMediaUnread {
				mediaUnreadDO := m.dao.ChannelMediaUnreadDAO.SelectMediaUnread(toUserId, int32(-m.DialogId), m.DialogMessageId)
				if mediaUnreadDO == nil {
					message.Data2.MediaUnread = m.HasMediaUnread
				} else {
					message.Data2.MediaUnread = false
				}
			}
		}
	default:
		// TODO(@benqi): unknown error.
	}

	message.Data2.Id = m.MessageId
	message.Data2.Mentioned = m.Mentioned
	message.Data2.MediaUnread = m.MediaUnread

	return message
}

type MessageData struct {
	MessageDataId   int64
	DialogId        int64
	DialogMessageId int32
	SenderUserId    int32
	Peer            *base.PeerUtil
	RandomId        int64
	Message         *mtproto.Message
	HasMediaUnread  bool
	EditMessage     string
	EditDate        int32
	Views           int32
	dao             *messagesDAO
}

///////////////////////////////////////////////////////////////////////////////////////////////////////////
// 发一条消息
// 存储MessageData
// 创建收件箱
// push到发件箱
//
func (m *MessageModel) MakeMessageData(senderUserId int32, peer *base.PeerUtil, clientRandomId int64, hasMediaUnread bool, message *mtproto.Message) *MessageData {
	// TODO(@benqi): check inputPeerEmpty
	did := makeDialogId(senderUserId, peer.PeerType, peer.PeerId)
	return &MessageData{
		MessageDataId:   core.GetUUID(),
		DialogId:        did,
		DialogMessageId: int32(core.NextMessageDataId(did)),
		SenderUserId:    senderUserId,
		Peer:            peer,
		RandomId:        clientRandomId,
		Message:         message,
		HasMediaUnread:  hasMediaUnread,
		// Views:           message.GetData2().GetViews(),
		dao: m.dao,
	}
}

func (m *MessageData) Insert() (rowsAffected int64) {
	rowsAffected = -1
	mtype, mdata := encodeMessage(m.Message)
	switch m.Peer.PeerType {
	case base.PEER_USER, base.PEER_CHAT:
		messageDataDO := &dataobject.MessageDatasDO{
			MessageDataId:   m.MessageDataId,
			DialogId:        m.DialogId,
			DialogMessageId: int32(m.DialogMessageId),
			SenderUserId:    m.SenderUserId,
			PeerType:        int8(m.Peer.PeerType),
			PeerId:          m.Peer.PeerId,
			RandomId:        m.RandomId,
			MessageType:     int8(mtype),
			MessageData:     string(mdata),
			MediaUnread:     base2.BoolToInt8(m.HasMediaUnread),
			HasMediaUnread:  base2.BoolToInt8(m.HasMediaUnread),
			Date:            int32(time.Now().Unix()),
			Deleted:         0,
		}
		glog.Info(messageDataDO)
		// TODO(@benqi): random_id已经存在
		rowsAffected = m.dao.MessageDatasDAO.Insert(messageDataDO)
		if rowsAffected == 0 {
			do := m.dao.MessageDatasDAO.SelectMessageByRandomId(m.SenderUserId, m.RandomId)
			if do != nil {
				m.MessageDataId = do.MessageDataId
				m.DialogMessageId = do.DialogMessageId
			} else {
				rowsAffected = -1
			}
		}
	case base.PEER_CHANNEL:
		channelMessageDataDO := &dataobject.ChannelMessagesDO{
			ChannelId:        int32(-m.DialogId),
			ChannelMessageId: m.DialogMessageId,
			SenderUserId:     m.SenderUserId,
			RandomId:         m.RandomId,
			MessageDataId:    m.MessageDataId,
			MessageType:      int8(mtype),
			MessageData:      string(mdata),
			HasMediaUnread:   base2.BoolToInt8(m.HasMediaUnread),
			Views:            m.Views,
			Date:             int32(time.Now().Unix()),
			Deleted:          0,
		}
		// TODO(@benqi): random_id已经存在
		rowsAffected = m.dao.ChannelMessagesDAO.Insert(channelMessageDataDO)
		if rowsAffected == 0 {
			do := m.dao.ChannelMessagesDAO.SelectByRandomId(m.SenderUserId, m.RandomId)
			if do != nil {
				m.MessageDataId = do.MessageDataId
				m.DialogMessageId = do.ChannelMessageId
			} else {
				rowsAffected = -1
			}
		}
	}

	return
}

func (m *MessageData) SaveMessageData() bool {
	m.dao.MessageDatasDAO.UpdateEditMessage(m.EditMessage, m.EditDate, m.DialogId, m.DialogMessageId)
	return true
}

func (m *MessageModel) SendInternalMessageToOutbox(senderUserId int32,
	peer *base.PeerUtil,
	clientRandomId int64,
	message *mtproto.Message,
	cb OnBoxCallback) (*MessageBox2, error) {

	messageData := m.MakeMessageData(senderUserId, peer, clientRandomId, false, message)
	glog.Info("messageData: ", messageData)

	lastInsertId := messageData.Insert()
	if lastInsertId == -1 {
		// glog.Error(err)
		err := fmt.Errorf("insert error")
		return nil, err
	}

	outBoxReplyToMsgId := message.Data2.ReplyToMsgId
	outBox := MakeMessageOutBox(senderUserId, false, outBoxReplyToMsgId, messageData)
	outBox.Insert()
	if cb != nil {
		cb(senderUserId, outBox)
	}

	return outBox, nil
}

///////////////////////////////////////////////////////////////////////////////////////////////////////////
// 发消息
func (m *MessageModel) SendInternalMessage(senderUserId int32,
	peer *base.PeerUtil,
	clientRandomId int64,
	hasMediaUnread bool,
	message *mtproto.Message,
	cb OnBoxCallback) error {

	messageData := m.MakeMessageData(senderUserId, peer, clientRandomId, hasMediaUnread, message)
	glog.Info("messageData: ", messageData)

	lastInsertId := messageData.Insert()
	if lastInsertId == -1 {
		// glog.Error(err)
		err := fmt.Errorf("insert error")
		return err
	}

	switch peer.PeerType {
	case base.PEER_USER, base.PEER_CHAT:
		outBoxReplyToMsgId := message.Data2.ReplyToMsgId
		outBox := MakeMessageOutBox(senderUserId, false, outBoxReplyToMsgId, messageData)
		outBox.Insert()
		if cb != nil {
			cb(senderUserId, outBox)
		}

		if peer.PeerType == base.PEER_USER {
			if senderUserId != peer.PeerId {
				// TODO(@benqi): get replyToMsgId
				inBoxReplyToMsgId := m.GetPeerMessageId(outBox.OwnerId, outBoxReplyToMsgId, peer.PeerId)
				inBox := MakeMessageInBox(peer.PeerId, hasMediaUnread, inBoxReplyToMsgId, messageData)
				inBox.Insert()
				if cb != nil {
					cb(peer.PeerId, inBox)
				}
			} else {
				// send self.
			}
		} else {
			doList := m.dao.ChatParticipantsDAO.SelectList(peer.PeerId)
			for i := 0; i < len(doList); i++ {
				glog.Info("chatParticipants - ", doList[i], "; senderUserId = ", senderUserId)
				if senderUserId == doList[i].UserId || doList[i].State != 0 {
					continue
				}
				inBoxReplyToMsgId := m.GetPeerMessageId(outBox.OwnerId, outBoxReplyToMsgId, doList[i].UserId)
				inBox := MakeMessageInBox(doList[i].UserId, hasMediaUnread, inBoxReplyToMsgId, messageData)
				inBox.Insert()
				if inBox.Mentioned {
					m.InsertUnreadMention(inBox.OwnerId, int8(base.PEER_CHAT), peer.PeerId, inBox.MessageId)
				}

				// inbox := this.makeInboxMessageDO(fromId, int(base.PEER_CHAT), peerId, do.UserId)
				glog.Info("insertChatMessageToInbox - ", inBox)
				if cb != nil {
					cb(peer.PeerId, inBox)
				}
			}
			// GetParatiants.
		}
	case base.PEER_CHANNEL:
		// 发给Channel
		channelBox := MakeChannelMessageBox(peer.PeerId, messageData)
		channelBox.Insert()
		if cb != nil {
			cb(senderUserId, channelBox)
		}
	default:
	}

	// TODO(@benqi): error
	return nil
}

func (m *MessageModel) makeMessageBoxByDO(boxDO *dataobject.MessageBoxesDO, dataDO *dataobject.MessageDatasDO) *MessageBox2 {
	// TODO(@benqi): check boxDO and dataDO
	mBox := &MessageBox2{
		OwnerId:        boxDO.UserId,
		MessageId:      boxDO.UserMessageBoxId,
		MessageBoxType: boxDO.MessageBoxType,
		Mentioned:      base2.Int8ToBool(boxDO.Mentioned),
		MediaUnread:    base2.Int8ToBool(boxDO.MediaUnread),
		ReplyToMsgId:   boxDO.ReplyToMsgId,
	}

	mData := &MessageData{
		MessageDataId:   boxDO.MessageDataId,
		DialogId:        boxDO.DialogId,
		DialogMessageId: boxDO.DialogMessageId,
		SenderUserId:    dataDO.SenderUserId,
		Peer:            &base.PeerUtil{PeerType: int32(dataDO.PeerType), PeerId: dataDO.PeerId},
		RandomId:        dataDO.RandomId,
		EditMessage:     dataDO.EditMessage,
		EditDate:        dataDO.EditDate,
		// Views:           dataDO.Views,
		dao: m.dao,
	}

	mData.Message, _ = decodeMessage(int(dataDO.MessageType), []byte(dataDO.MessageData))
	mBox.MessageData = mData

	return mBox
}

func (m *MessageModel) GetMessageBox2(peerType, ownerId, messageId int32) (*MessageBox2, error) {
	var (
		mBox *MessageBox2
	)
	glog.Infoln(ownerId)
	glog.Infoln(messageId)
	switch peerType {
	case base.PEER_USER, base.PEER_CHAT:
		boxDO := m.dao.MessageBoxesDAO.SelectByMessageId(ownerId, messageId)
		if boxDO == nil {
			return nil, fmt.Errorf("2222222")
			fmt.Errorf(string(ownerId), string(messageId))
			glog.Infoln(ownerId, messageId)
		}

		dataDO := m.dao.MessageDatasDAO.SelectMessage(boxDO.DialogId, boxDO.DialogMessageId)
		if dataDO == nil {
			return nil, fmt.Errorf("33333333")
		}

		mBox = m.makeMessageBoxByDO(boxDO, dataDO)
	case base.PEER_CHANNEL:
		glog.Warning("blocked, License key from https://nebula.chat required to unlock enterprise features.")
	default:
		return nil, fmt.Errorf("11111111")
	}

	return mBox, nil
}

func (m *MessageModel) GetMessageBox3(peerType int32, messageId int64) (*MessageBox2, error) {
	var (
		mBox *MessageBox2
	)

	glog.Infoln(messageId)
	switch peerType {
	case base.PEER_USER, base.PEER_CHAT:
		boxDO := m.dao.MessageBoxesDAO.SelectByMessageDataId(messageId)
		if boxDO == nil {
			return nil, fmt.Errorf("2222222")
			fmt.Errorf(string(messageId))
			glog.Infoln(messageId)
		}

		dataDO := m.dao.MessageDatasDAO.SelectMessage(boxDO.DialogId, boxDO.DialogMessageId)
		if dataDO == nil {
			return nil, fmt.Errorf("33333333")
		}

		mBox = m.makeMessageBoxByDO(boxDO, dataDO)
	case base.PEER_CHANNEL:
		glog.Warning("blocked, License key from https://nebula.chat required to unlock enterprise features.")
	default:
		return nil, fmt.Errorf("11111111")
	}

	return mBox, nil
}

func (m *MessageModel) SendChannelMessage(sendUserId int32,
	peer *base.PeerUtil,
	randomId int64,
	outboxMessage *mtproto.Message,
	resultCB func(pts, ptsCount int32, channelBox *MessageBox2) *mtproto.Updates,
	syncNotMeCB func(pts, ptsCount int32, channelBox *MessageBox2) ([]int32, int64, *mtproto.Updates, error),
	pushCB func(userId, pts, ptsCount int32, channelBox *MessageBox2) (*mtproto.Updates, error)) (*mtproto.Updates, error) {

	// TODO(@benqi): rollback
	defer func() {

	}()

	var channelBox *MessageBox2
	err := m.SendInternalMessage(sendUserId, peer, randomId, false, outboxMessage, func(ownerId int32, box2 *MessageBox2) {
		channelBox = box2
	})

	if err != nil {
		glog.Error(err)
		return nil, err
	}

	pts := int32(core.NextChannelPtsId(channelBox.OwnerId))
	ptsCount := int32(1)

	idList, authKeyId, syncUpdates, err := syncNotMeCB(pts, ptsCount, channelBox)
	sync_client.GetSyncClient().SyncChannelUpdatesNotMe(channelBox.OwnerId, sendUserId, authKeyId, syncUpdates)

	for _, id := range idList {
		pushUpdates, _ := pushCB(id, pts, ptsCount, channelBox)
		sync_client.GetSyncClient().PushChannelUpdates(channelBox.OwnerId, id, pushUpdates)
	}

	return resultCB(pts, ptsCount, channelBox), nil
}

func (m *MessageModel) SendMessage(sendUserId int32,
	peer *base.PeerUtil,
	randomId int64,
	outboxMessage *mtproto.Message,
	resultCB func(pts, ptsCount int32, outBox *MessageBox2) (*mtproto.Updates, error),
	syncNotMeCB func(pts, ptsCount int32, outBox *MessageBox2) (int64, *mtproto.Updates, error),
	pushCB func(pts, ptsCount int32, inBox *MessageBox2) (*mtproto.Updates, error)) (*mtproto.Updates, error) {

	// TODO(@benqi): rollback
	defer func() {

	}()

	var boxList []*MessageBox2
	err := m.SendInternalMessage(sendUserId, peer, randomId, false, outboxMessage, func(ownerId int32, box2 *MessageBox2) {
		// glog.Info("SendInternalMessage - ", box2)
		switch box2.MessageBoxType {
		case MESSAGE_BOX_TYPE_OUTGOING:
			// 1. update user_dialog
			m.dialogCallback.InsertOrUpdateDialog(box2.OwnerId, box2.Peer.PeerType, box2.Peer.PeerId, box2.MessageId, false, false)
		case MESSAGE_BOX_TYPE_INCOMING:
			if box2.Peer.PeerType == base.PEER_USER {
				// 1. update user_dialog
				m.dialogCallback.InsertOrUpdateDialog(box2.OwnerId, box2.Peer.PeerType, box2.SenderUserId, box2.MessageId, false, true)
			} else {
				hasMentioned := hasMention(box2.Message.Data2.Entities, box2.OwnerId)
				m.dialogCallback.InsertOrUpdateDialog(box2.OwnerId, box2.Peer.PeerType, box2.Peer.PeerId, box2.MessageId, hasMentioned, true)
			}
		case MESSAGE_BOX_TYPE_CHANNEL:
		default:
		}
		boxList = append(boxList, box2)
	})

	if err != nil {
		glog.Error(err)
		return nil, err
	}

	if len(boxList) == 0 {
		err = fmt.Errorf("boxList empty")
		glog.Error(err)
		return nil, err
	}

	// 3. 发件箱
	var (
		pts      = int32(0)
		ptsCount = int32(0)
	)

	outBox := boxList[0]
	pts = int32(core.NextPtsId(outBox.OwnerId))
	ptsCount = 1

	// return
	reply, err := resultCB(pts, ptsCount, outBox)
	if err != nil {
		glog.Error(err)
		return nil, err
	}

	if syncNotMeCB != nil {
		authKeyId, syncNotMeUpdates, err := syncNotMeCB(pts, ptsCount, outBox)
		if err == nil {
			sync_client.GetSyncClient().SyncUpdatesNotMe(sendUserId, authKeyId, syncNotMeUpdates)
		}
	}

	// 4. 收件箱
	inBoxes := boxList[1:]
	for i := 0; i < len(inBoxes); i++ {
		switch peer.PeerType {
		case base.PEER_USER:
			pts = int32(core.NextPtsId(inBoxes[i].OwnerId))
			ptsCount = 1
			pushUpdates, err := pushCB(pts, ptsCount, inBoxes[i])
			if err != nil {
				glog.Error(err)
				return nil, err
			}
			sync_client.GetSyncClient().PushUpdates(inBoxes[i].OwnerId, pushUpdates)
		case base.PEER_CHAT:
			pts = int32(core.NextPtsId(inBoxes[i].OwnerId))
			ptsCount = 1
			pushUpdates, err := pushCB(pts, ptsCount, inBoxes[i])
			if err != nil {
				glog.Error(err)
				return nil, err
			}
			sync_client.GetSyncClient().PushUpdates(inBoxes[i].OwnerId, pushUpdates)
		case base.PEER_CHANNEL:
		default:
		}
	}

	return reply, nil
}

func (m *MessageModel) SendMultiMessage(sendUserId int32,
	peer *base.PeerUtil,
	randomIdList []int64,
	outboxMessages []*mtproto.Message,
	resultCB func(pts, ptsCount int32, outBoxList []*MessageBox2) (*mtproto.Updates, error),
	syncNotMeCB func(pts, ptsCount int32, outBoxList []*MessageBox2) (int64, *mtproto.Updates, error),
	pushCB func(userId, pts, ptsCount int32, inBoxList []*MessageBox2) (*mtproto.Updates, error)) (*mtproto.Updates, error) {

	// TODO(@benqi): rollback
	defer func() {

	}()

	outBoxList := make([]*MessageBox2, 0)
	boxListMap := map[int32][]*MessageBox2{}

	for i, outboxMessage := range outboxMessages {
		err := m.SendInternalMessage(sendUserId, peer, randomIdList[i], false, outboxMessage, func(ownerId int32, box2 *MessageBox2) {
			switch box2.MessageBoxType {
			case MESSAGE_BOX_TYPE_OUTGOING:
				glog.Info("SendMultiMessage - ", box2)
				// 1. update user_dialog
				outBoxList = append(outBoxList, box2)
				m.dialogCallback.InsertOrUpdateDialog(box2.OwnerId, box2.Peer.PeerType, box2.Peer.PeerId, box2.MessageId, false, false)
			case MESSAGE_BOX_TYPE_INCOMING:
				glog.Info("SendMultiMessage - ", box2)
				var (
					inBoxList []*MessageBox2
				)
				if _, ok := boxListMap[box2.OwnerId]; ok {
					inBoxList = boxListMap[box2.OwnerId]
				}
				inBoxList = append(inBoxList, box2)
				boxListMap[box2.OwnerId] = inBoxList

				// 1. update user_dialog
				if box2.Peer.PeerType == base.PEER_USER {
					// 1. update user_dialog
					m.dialogCallback.InsertOrUpdateDialog(box2.OwnerId, box2.Peer.PeerType, box2.SenderUserId, box2.MessageId, false, true)
				} else {
					hasMentioned := hasMention(box2.Message.Data2.Entities, box2.OwnerId)
					m.dialogCallback.InsertOrUpdateDialog(box2.OwnerId, box2.Peer.PeerType, box2.Peer.PeerId, box2.MessageId, hasMentioned, true)
				}
			case MESSAGE_BOX_TYPE_CHANNEL:
			default:
			}
			// boxList = append(boxList, box2)
		})

		if err != nil {
			glog.Error(err)
			return nil, err
		}
	}

	//if len(boxListMap) == 0 {
	//	err := fmt.Errorf("boxListMap empty")
	//	glog.Error(err)
	//	return nil, err
	//}

	//// 3. 发件箱
	var (
		pts      = int32(0)
		ptsCount = int32(0)
		// updates *mtproto.Updates
	)

	pts = int32(core.NextNPtsId(sendUserId, len(outBoxList)))
	ptsCount = int32(len(outBoxList))

	// return
	replyUpdates, err := resultCB(pts, ptsCount, outBoxList)
	if err != nil {
		glog.Error(err)
		return nil, err
	}

	authKeyId, syncNotMeUpdates, err := syncNotMeCB(pts, ptsCount, outBoxList)
	sync_client.GetSyncClient().SyncUpdatesNotMe(sendUserId, authKeyId, syncNotMeUpdates)

	for inBoxUserId, inBoxList := range boxListMap {
		switch peer.PeerType {
		case base.PEER_USER:
			pts = int32(core.NextNPtsId(inBoxUserId, len(inBoxList)))
			ptsCount = int32(len(inBoxList))
			pushUpdates, err := pushCB(inBoxUserId, pts, ptsCount, inBoxList)
			if err != nil {
				glog.Error(err)
				return nil, err
			}
			sync_client.GetSyncClient().PushUpdates(inBoxUserId, pushUpdates)
		case base.PEER_CHAT:
			pts = int32(core.NextNPtsId(inBoxUserId, len(inBoxList)))
			pushUpdates, err := pushCB(inBoxUserId, pts, ptsCount, inBoxList)
			if err != nil {
				glog.Error(err)
				return nil, err
			}
			// glog.Info("push - ", inBoxUserId, ", ", pushUpdates)
			sync_client.GetSyncClient().PushUpdates(inBoxUserId, pushUpdates)
		case base.PEER_CHANNEL:
		default:
		}
	}

	return replyUpdates, nil
}
