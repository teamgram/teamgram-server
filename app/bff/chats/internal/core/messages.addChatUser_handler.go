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
	"encoding/json"
	"math/rand"
	"time"

	msgpb "github.com/teamgram/teamgram-server/v2/app/messenger/msg/msg"
	"github.com/teamgram/teamgram-server/v2/app/messenger/userupdates/payload"
	chatpb "github.com/teamgram/teamgram-server/v2/app/service/biz/chat/chat"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

// MessagesAddChatUser
// messages.addChatUser#cbc6d107 chat_id:long user_id:InputUser fwd_limit:int = messages.InvitedUsers;
func (c *ChatsCore) MessagesAddChatUser(in *tg.TLMessagesAddChatUser) (*tg.MessagesInvitedUsers, error) {
	md := c.MD
	if md == nil || md.UserId <= 0 {
		return nil, tg.ErrUserIdInvalid
	}
	if md.PermAuthKeyId == 0 {
		return nil, tg.ErrAuthKeyPermEmpty
	}
	selfID := md.UserId
	user := tg.FromInputUser(selfID, in.UserId)
	if user.PeerType != tg.PEER_USER {
		return nil, tg.ErrUserIdInvalid
	}

	mutableChat, err := c.svcCtx.Repo.ChatClient.ChatAddChatUser(c.ctx, &chatpb.TLChatAddChatUser{
		ChatId:    in.ChatId,
		InviterId: selfID,
		UserId:    user.PeerId,
	})
	if err != nil {
		return nil, mapChatError(err)
	}

	participantsFact, err := chatParticipantsChangedFactFromMutableChatForActor(mutableChat, selfID, []int64{user.PeerId})
	if err != nil {
		c.Logger.Errorf("messages.addChatUser - malformed mutable chat: self_user_id=%d chat_id=%d invitee_user_id=%d err=%v", selfID, in.ChatId, user.PeerId, err)
		return nil, tg.ErrInternalServerError
	}
	if in.FwdLimit > 0 {
		participantsFact.FwdLimit = in.FwdLimit
	}
	attachFact, err := payload.WrapFact(payload.FactKindChatParticipantsChanged, participantsFact)
	if err != nil {
		c.Logger.Errorf("messages.addChatUser - wrap chat participants fact failed: self_user_id=%d chat_id=%d invitee_user_id=%d err=%v", selfID, in.ChatId, user.PeerId, err)
		return nil, tg.ErrInternalServerError
	}
	attachPayload, err := json.Marshal(attachFact)
	if err != nil {
		c.Logger.Errorf("messages.addChatUser - marshal chat participants fact failed: self_user_id=%d chat_id=%d invitee_user_id=%d err=%v", selfID, in.ChatId, user.PeerId, err)
		return nil, tg.ErrInternalServerError
	}

	updates, err := c.svcCtx.Repo.MsgClient.MsgSendMessage(c.ctx, &msgpb.TLMsgSendMessage{
		UserId:    selfID,
		AuthKeyId: md.PermAuthKeyId,
		AttachFacts: []msgpb.UpdateFactClazz{
			msgpb.MakeTLUpdateFact(&msgpb.TLUpdateFact{
				Kind:    payload.FactKindChatParticipantsChanged,
				Payload: attachPayload,
			}),
		},
		PeerType: payload.PeerTypeChat,
		PeerId:   mutableChat.Chat.Id,
		Message: []msgpb.OutboxMessageClazz{
			msgpb.MakeTLOutboxMessage(&msgpb.TLOutboxMessage{
				NoWebpage: true,
				RandomId:  normalizeCreateChatServiceMessageRandomID(rand.Int63()),
				Message: tg.MakeTLMessageService(&tg.TLMessageService{
					Out:    true,
					FromId: tg.MakePeerUser(selfID),
					PeerId: tg.MakePeerChat(mutableChat.Chat.Id),
					Date:   int32(time.Now().Unix()),
					Action: tg.MakeTLMessageActionChatAddUser(&tg.TLMessageActionChatAddUser{
						Users: []int64{user.PeerId},
					}),
				}),
			}),
		},
	})
	if err != nil {
		c.Logger.Errorf("messages.addChatUser - send add user service message failed: self_user_id=%d chat_id=%d invitee_user_id=%d err=%v", selfID, mutableChat.Chat.Id, user.PeerId, err)
		return nil, tg.ErrInternalServerError
	}
	if updates == nil {
		c.Logger.Errorf("messages.addChatUser - send add user service message returned nil updates: self_user_id=%d chat_id=%d invitee_user_id=%d", selfID, mutableChat.Chat.Id, user.PeerId)
		return nil, tg.ErrInternalServerError
	}

	return tg.MakeTLMessagesInvitedUsers(&tg.TLMessagesInvitedUsers{
		Updates:         updates.Clazz,
		MissingInvitees: []tg.MissingInviteeClazz{},
	}).ToMessagesInvitedUsers(), nil
}
