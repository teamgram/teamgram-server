// Copyright (c) 2024 The Teamgooo Authors. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Author: teamgramio (teamgram.io@gmail.com)

package core

import (
	"time"

	"github.com/teamgram/teamgram-server/v2/app/messenger/msg/msg/msg"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

var _ *tg.Bool

// MsgSendMessageV2
// msg.sendMessageV2 user_id:long auth_key_id:long peer_type:int peer_id:long message:Vector<OutboxMessage> = Updates;
func (c *MsgCore) MsgSendMessageV2(in *msg.TLMsgSendMessageV2) (*tg.Updates, error) {
	if len(in.Message) == 0 {
		return nil, tg.ErrInputRequestInvalid
	}

	switch in.PeerType {
	case tg.PEER_SELF, tg.PEER_USER, tg.PEER_CHAT:
		outbox := in.Message[0].ToOutboxMessage()
		if outbox == nil {
			return nil, tg.ErrInputRequestInvalid
		}

		date := int32(time.Now().Unix())
		var entities []tg.MessageEntityClazz
		if outbox.Message != nil {
			if msg2, ok := outbox.Message.(*tg.TLMessage); ok {
				if msg2.Date != 0 {
					date = msg2.Date
				}
				entities = msg2.Entities
			}
		}

		return tg.MakeTLUpdateShortSentMessage(&tg.TLUpdateShortSentMessage{
			Out:      true,
			Id:       placeholderMessageID(outbox.RandomId),
			Pts:      1,
			PtsCount: 1,
			Date:     date,
			Entities: entities,
		}).ToUpdates(), nil
	case tg.PEER_CHANNEL:
		return nil, tg.ErrEnterpriseIsBlocked
	default:
		return nil, tg.ErrPeerIdInvalid
	}
}

func placeholderMessageID(randomID int64) int32 {
	if randomID < 0 {
		randomID = -randomID
	}
	id := int32(randomID % 0x7fffffff)
	if id == 0 {
		id = 1
	}
	return id
}
