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

package message

import "github.com/nebula-chat/chatengine/mtproto"

func (m *MessageModel) GetChannelMessage(userId, channelId, id int32) (message *mtproto.Message) {
	do := m.dao.ChannelMessagesDAO.SelectByMessageId(channelId, id)
	if do == nil {
		return
	}
	boxDO := m.makeChannelMessageBoxByDO(do)
	return boxDO.ToMessage(userId)
}

func (m *MessageModel) GetChannelMessageList(userId, channelId int32, idList []int32) (messages []*mtproto.Message) {
	if len(idList) == 0 {
		messages = []*mtproto.Message{}
	} else {
		doList := m.dao.ChannelMessagesDAO.SelectByMessageIdList(channelId, idList)
		messages = make([]*mtproto.Message, 0, len(doList))
		for i := 0; i < len(doList); i++ {
			// TODO(@benqi): check data
			boxDO := m.makeChannelMessageBoxByDO(&doList[i])
			messages = append(messages, boxDO.ToMessage(userId))
		}
	}
	return
}

/*
import (
	"encoding/json"
	"fmt"
	"github.com/golang/glog"
	"github.com/nebulaim/telegramd/biz/base"
	"github.com/nebulaim/telegramd/biz/core"
	"github.com/nebulaim/telegramd/biz/dal/dataobject"
	"github.com/nebulaim/telegramd/proto/mtproto"
	"time"
)

type ChannelMessageBox struct {
	SenderUserId        int32
	ChannelId           int32
	ChannelMessageBoxId int32
	MessageId           int64
	RandomId            int64
	Message             *mtproto.Message
}

type ChannelBoxCreated func(int32)

func (m *MessageModel) CreateChannelMessageBoxByNew(fromId, channelId int32, clientRandomId int64, message2 *mtproto.Message, cb ChannelBoxCreated) (box *ChannelMessageBox) {
	now := int32(time.Now().Unix())
	boxId := int32(core.NextChannelMessageBoxId(channelId))
	messageDatasDO := &dataobject.MessageDatasDO{
		DialogId:     int64(-channelId),
		MessageId:    core.GetUUID(),
		SenderUserId: fromId,
		PeerType:     int8(base.PEER_CHANNEL),
		PeerId:       channelId,
		RandomId:     clientRandomId,
		Date:         now,
		Deleted:      0,
	}

	channelMessageBoxesDO := &dataobject.ChannelMessageBoxesDO{
		SenderUserId:        fromId,
		ChannelId:           channelId,
		ChannelMessageBoxId: boxId,
		MessageId:           messageDatasDO.MessageId,
		Date:                now,
	}

	switch message2.GetConstructor() {
	case mtproto.TLConstructor_CRC32_messageEmpty:
		messageDatasDO.MessageType = MESSAGE_TYPE_MESSAGE_EMPTY
	case mtproto.TLConstructor_CRC32_message:
		messageDatasDO.MessageType = MESSAGE_TYPE_MESSAGE
		message := message2.To_Message()

		// mentioned = message.GetMentioned()
		message.SetId(channelMessageBoxesDO.ChannelMessageBoxId)
	case mtproto.TLConstructor_CRC32_messageService:
		messageDatasDO.MessageType = MESSAGE_TYPE_MESSAGE_SERVICE
		message := message2.To_MessageService()

		// mentioned = message.GetMentioned()
		message.SetId(channelMessageBoxesDO.ChannelMessageBoxId)
	}

	messageData, _ := json.Marshal(message2)
	messageDatasDO.MessageData = string(messageData)

	//// TODO(@benqi): pocess clientRandomId dup
	m.dao.MessageDatasDAO.Insert(messageDatasDO)
	m.dao.ChannelMessageBoxesDAO.Insert(channelMessageBoxesDO)

	box = &ChannelMessageBox{
		SenderUserId:        fromId,
		ChannelId:           channelId,
		ChannelMessageBoxId: boxId,
		MessageId:           channelMessageBoxesDO.MessageId,
		RandomId:            clientRandomId,
		Message:             message2,
	}

	if cb != nil {
		cb(box.ChannelMessageBoxId)
	}

	return
}

func doToChannelMessage(do *dataobject.MessageDatasDO) (*mtproto.Message, error) {
	message := &mtproto.Message{
		Data2: &mtproto.Message_Data{},
	}

	switch do.MessageType {
	case MESSAGE_TYPE_MESSAGE_EMPTY:
		message.Constructor = mtproto.TLConstructor_CRC32_messageEmpty
		// message = message2
	case MESSAGE_TYPE_MESSAGE:
		// err := proto.Unmarshal(messageDO.MessageData, message)
		err := json.Unmarshal([]byte(do.MessageData), message)
		if err != nil {
			glog.Errorf("messageDOToMessage - Unmarshal message(%d)error: %v", do.Id, err)
			return nil, err
		} else {
			message.Constructor = mtproto.TLConstructor_CRC32_message
		}
	case MESSAGE_TYPE_MESSAGE_SERVICE:
		err := json.Unmarshal([]byte(do.MessageData), message)
		if err != nil {
			glog.Errorf("messageDOToMessage - Unmarshal message(%d)error: %v", do.Id, err)
			return nil, err
		} else {
			message.Constructor = mtproto.TLConstructor_CRC32_messageService
		}
	default:
		err := fmt.Errorf("messageDOToMessage - Invalid messageType, db's data error, message(%d)", do.Id)
		glog.Error(err)
		return nil, err
	}

	return message, nil
}

func (m *MessageModel) GetChannelMessage(channelId int32, id int32) (message *mtproto.Message) {
	do := m.dao.ChannelMessageBoxesDAO.SelectByMessageId(channelId, id)
	if do == nil {
		return
	}

	messageDO := m.dao.MessageDatasDAO.SelectByMessageId(do.MessageId)
	message, _ = doToChannelMessage(messageDO)
	return
}

func (m *MessageModel) GetChannelMessageList(channelId int32, idList []int32) (messages []*mtproto.Message) {
	if len(idList) == 0 {
		messages = []*mtproto.Message{}
	} else {
		doList := m.dao.ChannelMessageBoxesDAO.SelectByMessageIdList(channelId, idList)
		messages = make([]*mtproto.Message, 0, len(doList))
		for i := 0; i < len(doList); i++ {
			// TODO(@benqi): check data
			messageDO := m.dao.MessageDatasDAO.SelectByMessageId(doList[i].MessageId)
			if messageDO == nil {
				continue
			}
			m, _ := doToChannelMessage(messageDO)
			if m != nil {
				messages = append(messages, m)
			}
		}
	}
	return
}
*/
