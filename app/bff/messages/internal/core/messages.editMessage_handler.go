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

// MessagesEditMessage
// messages.editMessage#dfd14005 flags:# no_webpage:flags.1?true invert_media:flags.16?true peer:InputPeer id:int message:flags.11?string media:flags.14?InputMedia reply_markup:flags.2?ReplyMarkup entities:flags.3?Vector<MessageEntity> schedule_date:flags.15?int quick_reply_shortcut_id:flags.17?int = Updates;
func (c *MessagesCore) MessagesEditMessage(in *tg.TLMessagesEditMessage) (*tg.Updates, error) {
	var userId int64
	if c.MD != nil {
		userId = c.MD.UserId
	}

	peer := tg.FromInputPeer2(userId, in.Peer)
	switch peer.PeerType {
	case tg.PEER_SELF, tg.PEER_USER, tg.PEER_CHAT:
	case tg.PEER_CHANNEL:
		return nil, tg.ErrEnterpriseIsBlocked
	default:
		return nil, tg.ErrPeerIdInvalid
	}

	// When MsgClient is wired, delegate to msg service.
	if c.svcCtx != nil && c.svcCtx.MsgClient != nil {
		var authKeyId int64
		if c.MD != nil {
			authKeyId = c.MD.AuthId
		}

		var editType int32
		if in.Media != nil {
			editType = 1 // edit media
		} else {
			editType = 0 // edit text
		}

		var message string
		if in.Message != nil {
			message = *in.Message
		}

		var entities []tg.MessageEntityClazz
		if len(in.Entities) > 0 {
			entities = in.Entities
		}

		outboxMsg := msg.MakeTLOutboxMessage(&msg.TLOutboxMessage{
			NoWebpage: in.NoWebpage,
			RandomId:  int64(in.Id),
			Message: tg.MakeTLMessage(&tg.TLMessage{
				Out:      true,
				Date:     int32(time.Now().Unix()),
				Message:  message,
				Entities: entities,
			}),
		})

		return c.svcCtx.MsgClient.MsgEditMessageV2(c.ctx, &msg.TLMsgEditMessageV2{
			UserId:     userId,
			AuthKeyId:  authKeyId,
			PeerType:   peer.PeerType,
			PeerId:     peer.PeerId,
			EditType:   editType,
			NewMessage: outboxMsg,
			DstMessage: &tg.TLMessageBox{
				MessageId: in.Id,
			},
		})
	}

	// Fallback placeholder when MsgClient is not available.
	var entities []tg.MessageEntityClazz
	if len(in.Entities) > 0 {
		entities = in.Entities
	}

	return tg.MakeTLUpdateShortSentMessage(&tg.TLUpdateShortSentMessage{
		Out:      true,
		Id:       in.Id,
		Pts:      1,
		PtsCount: 1,
		Date:     int32(time.Now().Unix()),
		Entities: entities,
	}).ToUpdates(), nil
}
