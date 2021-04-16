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
	"encoding/json"
	"fmt"

	"github.com/golang/glog"
	"github.com/nebula-chat/chatengine/messenger/biz_server/biz/base"
	"github.com/nebula-chat/chatengine/mtproto"
)

// messageDOToMessage
func encodeMessage(message *mtproto.Message) (messageType int, messageData []byte) {
	switch message.GetConstructor() {
	case mtproto.TLConstructor_CRC32_messageEmpty:
		messageType = MESSAGE_TYPE_MESSAGE_EMPTY
	case mtproto.TLConstructor_CRC32_message:
		messageType = MESSAGE_TYPE_MESSAGE
	case mtproto.TLConstructor_CRC32_messageService:
		messageType = MESSAGE_TYPE_MESSAGE_SERVICE
	default:
		//
	}

	messageData, _ = json.Marshal(message)
	return
}

func decodeMessage(messageType int, messageData []byte) (*mtproto.Message, error) {
	message := &mtproto.Message{
		Data2: &mtproto.Message_Data{},
	}

	switch messageType {
	case MESSAGE_TYPE_MESSAGE_EMPTY:
		message.Constructor = mtproto.TLConstructor_CRC32_messageEmpty
		// message = message2
	case MESSAGE_TYPE_MESSAGE:
		err := json.Unmarshal(messageData, message)
		if err != nil {
			glog.Errorf("messageDOToMessage - Unmarshal message(%s)error: %v", messageData, err)
			return nil, err
		} else {
			message.Constructor = mtproto.TLConstructor_CRC32_message
		}
	case MESSAGE_TYPE_MESSAGE_SERVICE:
		err := json.Unmarshal(messageData, message)
		if err != nil {
			glog.Errorf("messageDOToMessage - Unmarshal message(%s)error: %v", messageData, err)
			return nil, err
		} else {
			message.Constructor = mtproto.TLConstructor_CRC32_messageService
		}
	default:
		err := fmt.Errorf("messageDOToMessage - Invalid messageType, db's data error, message(%s)", messageData)
		glog.Error(err)
		return nil, err
	}

	return message, nil
}

func makeDialogId(fromId, peerType, peerId int32) (did int64) {
	switch peerType {
	case base.PEER_SELF:
		did = int64(fromId)<<32 | int64(fromId)
	case base.PEER_USER:
		if fromId <= peerId {
			did = int64(fromId)<<32 | int64(peerId)
		} else {
			did = int64(peerId)<<32 | int64(fromId)
		}
	case base.PEER_CHAT:
		did = int64(-peerId)
	case base.PEER_CHANNEL:
		did = int64(-peerId)
	}
	return
}
