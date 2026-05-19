// Copyright (c) 2026 The Teamgram Authors. All rights reserved.
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
	"math/rand"
	"time"

	msgpb "github.com/teamgram/teamgram-server/v2/app/messenger/msg/msg"
	"github.com/teamgram/teamgram-server/v2/app/messenger/userupdates/payload"
	chatpb "github.com/teamgram/teamgram-server/v2/app/service/biz/chat/chat"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

// MessagesEditChatTitle
// messages.editChatTitle#73783ffd chat_id:long title:string = Updates;
func (c *ChatsCore) MessagesEditChatTitle(in *tg.TLMessagesEditChatTitle) (*tg.Updates, error) {
	selfID := selfID(c.MD)
	mutableChat, err := c.svcCtx.Repo.ChatClient.ChatEditChatTitle(c.ctx, &chatpb.TLChatEditChatTitle{
		ChatId:     in.ChatId,
		EditUserId: selfID,
		Title:      in.Title,
	})
	if err != nil {
		return nil, mapChatError(err)
	}

	updates, err := c.svcCtx.Repo.MsgClient.MsgSendMessage(c.ctx, &msgpb.TLMsgSendMessage{
		UserId:    selfID,
		AuthKeyId: c.MD.PermAuthKeyId,
		PeerType:  payload.PeerTypeChat,
		PeerId:    mutableChat.Chat.Id,
		Message: []msgpb.OutboxMessageClazz{
			msgpb.MakeTLOutboxMessage(&msgpb.TLOutboxMessage{
				NoWebpage: true,
				RandomId:  normalizeCreateChatServiceMessageRandomID(rand.Int63()),
				Message: tg.MakeTLMessageService(&tg.TLMessageService{
					Out:    true,
					FromId: tg.MakePeerUser(selfID),
					PeerId: tg.MakePeerChat(mutableChat.Chat.Id),
					Date:   int32(time.Now().Unix()),
					Action: tg.MakeTLMessageActionChatEditTitle(&tg.TLMessageActionChatEditTitle{
						Title: in.Title,
					}),
				}),
			}),
		},
	})
	if err != nil {
		c.Logger.Errorf("messages.editChatTitle - send edit title service message failed: self_user_id=%d chat_id=%d err=%v", selfID, mutableChat.Chat.Id, err)
		return nil, tg.ErrInternalServerError
	}
	if updates == nil {
		c.Logger.Errorf("messages.editChatTitle - send edit title service message returned nil updates: self_user_id=%d chat_id=%d", selfID, mutableChat.Chat.Id)
		return nil, tg.ErrInternalServerError
	}

	return updates, nil
}
