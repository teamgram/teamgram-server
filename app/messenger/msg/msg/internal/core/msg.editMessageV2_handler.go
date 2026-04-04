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

// MsgEditMessageV2
// msg.editMessageV2 user_id:long auth_key_id:long peer_type:int peer_id:long edit_type:int new_message:OutboxMessage dst_message:MessageBox = Updates;
func (c *MsgCore) MsgEditMessageV2(in *msg.TLMsgEditMessageV2) (*tg.Updates, error) {
	if in.NewMessage == nil || in.DstMessage == nil {
		return nil, tg.ErrInputRequestInvalid
	}

	switch in.PeerType {
	case tg.PEER_SELF, tg.PEER_USER, tg.PEER_CHAT:
	case tg.PEER_CHANNEL:
		return nil, tg.ErrEnterpriseIsBlocked
	default:
		return nil, tg.ErrPeerIdInvalid
	}

	outbox := in.NewMessage.ToOutboxMessage()
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

	messageID := in.DstMessage.MessageId
	if messageID <= 0 {
		messageID = c.nextMessageId(outbox.RandomId)
	}

	return tg.MakeTLUpdateShortSentMessage(&tg.TLUpdateShortSentMessage{
		Out:      true,
		Id:       messageID,
		Pts:      1,
		PtsCount: 1,
		Date:     date,
		Entities: entities,
	}).ToUpdates(), nil
}
