// Copyright 2022 Teamgram Authors
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
//
// Author: teamgramio (teamgram.io@gmail.com)
//

package core

import (
	"math/rand"

	"github.com/teamgram/proto/mtproto"
	msgpb "github.com/teamgram/teamgram-server/app/messenger/msg/msg/msg"
	chatpb "github.com/teamgram/teamgram-server/app/service/biz/chat/chat"
)

// MessagesEditChatTitle
// messages.editChatTitle#73783ffd chat_id:long title:string = Updates;
func (c *ChatsCore) MessagesEditChatTitle(in *mtproto.TLMessagesEditChatTitle) (*mtproto.Updates, error) {
	if in.Title == "" {
		err := mtproto.ErrChatTitleEmpty
		c.Logger.Errorf("messages.editChatTitle - error: ", err)
		return nil, err
	}

	chat, err := c.svcCtx.Dao.ChatClient.Client().ChatEditChatTitle(c.ctx, &chatpb.TLChatEditChatTitle{
		ChatId:     in.ChatId,
		EditUserId: c.MD.UserId,
		Title:      in.Title,
	})
	if err != nil {
		c.Logger.Errorf("messages.editChatTitle - error: ", err)
		return nil, err
	}

	replyUpdates, err := c.svcCtx.Dao.MsgClient.MsgSendMessageV2(
		c.ctx,
		&msgpb.TLMsgSendMessageV2{
			UserId:    c.MD.UserId,
			AuthKeyId: c.MD.PermAuthKeyId,
			PeerType:  mtproto.PEER_CHAT,
			PeerId:    in.ChatId,
			Message: []*msgpb.OutboxMessage{
				msgpb.MakeTLOutboxMessage(&msgpb.OutboxMessage{
					NoWebpage:    true,
					Background:   false,
					RandomId:     rand.Int63(),
					Message:      chat.MakeMessageService(c.MD.UserId, mtproto.MakeMessageActionChatEditTitle(in.Title)),
					ScheduleDate: nil,
				}).To_OutboxMessage(),
			},
		})
	if err != nil {
		c.Logger.Errorf("messages.editChatTitle - error: %v", err)
		return nil, err
	}

	return replyUpdates, nil
}
