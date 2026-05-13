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

// MessagesCreateChat
// messages.createChat#92ceddd4 flags:# users:Vector<InputUser> title:string ttl_period:flags.0?int = messages.InvitedUsers;
func (c *ChatsCore) MessagesCreateChat(in *tg.TLMessagesCreateChat) (*tg.MessagesInvitedUsers, error) {
	selfID := selfID(c.MD)
	userIDs := make([]int64, 0, len(in.Users))
	for _, inputUser := range in.Users {
		user := tg.FromInputUser(selfID, inputUser)
		if user.PeerType != tg.PEER_USER {
			return nil, tg.ErrUserIdInvalid
		}
		userIDs = append(userIDs, user.PeerId)
	}

	mutableChat, err := c.svcCtx.Repo.ChatClient.ChatCreateChat2(c.ctx, &chatpb.TLChatCreateChat2{
		CreatorId:  selfID,
		UserIdList: userIDs,
		Title:      in.Title,
		TtlPeriod:  in.TtlPeriod,
	})
	if err != nil {
		return nil, mapChatError(err)
	}
	if in.TtlPeriod != nil {
		// TODO: send the TTL service action once default history TTL is supported for newly created chats.
	}
	updates, err := c.svcCtx.Repo.MsgClient.MsgSendMessageV2(c.ctx, &msgpb.TLMsgSendMessageV2{
		UserId:    selfID,
		AuthKeyId: c.MD.PermAuthKeyId,
		PeerType:  payload.PeerTypeChat,
		PeerId:    mutableChat.Chat.Id,
		Message: []msgpb.OutboxMessageClazz{
			msgpb.MakeTLOutboxMessage(&msgpb.TLOutboxMessage{
				NoWebpage: true,
				RandomId:  rand.Int63(),
				Message: tg.MakeTLMessageService(&tg.TLMessageService{
					Out:    true,
					FromId: tg.MakePeerUser(selfID),
					PeerId: tg.MakePeerChat(mutableChat.Chat.Id),
					Date:   int32(time.Now().Unix()),
					Action: tg.MakeTLMessageActionChatCreate(&tg.TLMessageActionChatCreate{
						Title: in.Title,
						Users: userIDs,
					}),
					TtlPeriod: in.TtlPeriod,
				}),
			}),
		},
	})
	if err != nil {
		c.Logger.Errorf("messages.createChat - send create service message failed: self_user_id=%d chat_id=%d err=%v", selfID, mutableChat.Chat.Id, err)
		return nil, tg.ErrInternalServerError
	}

	return tg.MakeTLMessagesInvitedUsers(&tg.TLMessagesInvitedUsers{
		Updates:         updates.Clazz,
		MissingInvitees: []tg.MissingInviteeClazz{},
	}).ToMessagesInvitedUsers(), nil
}
