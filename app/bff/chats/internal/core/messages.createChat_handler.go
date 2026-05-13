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
	"crypto/sha256"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"math/rand"
	"time"

	msgpb "github.com/teamgram/teamgram-server/v2/app/messenger/msg/msg"
	"github.com/teamgram/teamgram-server/v2/app/messenger/userupdates/payload"
	chatpb "github.com/teamgram/teamgram-server/v2/app/service/biz/chat/chat"
	"github.com/teamgram/teamgram-server/v2/pkg/net/kitex/metadata"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

// MessagesCreateChat
// messages.createChat#92ceddd4 flags:# users:Vector<InputUser> title:string ttl_period:flags.0?int = messages.InvitedUsers;
func (c *ChatsCore) MessagesCreateChat(in *tg.TLMessagesCreateChat) (*tg.MessagesInvitedUsers, error) {
	md := c.MD
	if md == nil || md.UserId <= 0 {
		return nil, tg.ErrUserIdInvalid
	}
	if md.PermAuthKeyId == 0 {
		return nil, tg.ErrAuthKeyPermEmpty
	}

	selfID := md.UserId
	userIDs := make([]int64, 0, len(in.Users))
	for _, inputUser := range in.Users {
		user := tg.FromInputUser(selfID, inputUser)
		if user.PeerType != tg.PEER_USER {
			return nil, tg.ErrUserIdInvalid
		}
		userIDs = append(userIDs, user.PeerId)
	}

	clientMsgID := metadataClientMsgID(md)
	var clientMsgIDPtr *int64
	if clientMsgID != 0 {
		clientMsgIDPtr = &clientMsgID
	}

	mutableChat, err := c.svcCtx.Repo.ChatClient.ChatCreateChat2(c.ctx, &chatpb.TLChatCreateChat2{
		CreatorId:   selfID,
		UserIdList:  userIDs,
		Title:       in.Title,
		TtlPeriod:   in.TtlPeriod,
		ClientMsgId: clientMsgIDPtr,
	})
	if err != nil {
		return nil, mapChatError(err)
	}
	if in.TtlPeriod != nil {
		// TODO: send the TTL service action once default history TTL is supported for newly created chats.
	}
	participantsFact, err := chatParticipantsChangedFactFromMutableChat(mutableChat, selfID, userIDs)
	if err != nil {
		c.Logger.Errorf("messages.createChat - malformed mutable chat: self_user_id=%d err=%v", selfID, err)
		return nil, tg.ErrInternalServerError
	}
	attachFact, err := payload.WrapFact(payload.FactKindChatParticipantsChanged, participantsFact)
	if err != nil {
		c.Logger.Errorf("messages.createChat - wrap chat participants fact failed: self_user_id=%d chat_id=%d err=%v", selfID, mutableChat.Chat.Id, err)
		return nil, tg.ErrInternalServerError
	}
	attachPayload, err := json.Marshal(attachFact)
	if err != nil {
		c.Logger.Errorf("messages.createChat - marshal chat participants fact failed: self_user_id=%d chat_id=%d err=%v", selfID, mutableChat.Chat.Id, err)
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
				RandomId:  createChatServiceMessageRandomID(selfID, clientMsgID),
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
	if updates == nil {
		c.Logger.Errorf("messages.createChat - send create service message returned nil updates: self_user_id=%d chat_id=%d", selfID, mutableChat.Chat.Id)
		return nil, tg.ErrInternalServerError
	}

	return tg.MakeTLMessagesInvitedUsers(&tg.TLMessagesInvitedUsers{
		Updates:         updates.Clazz,
		MissingInvitees: []tg.MissingInviteeClazz{},
	}).ToMessagesInvitedUsers(), nil
}

func chatParticipantsChangedFactFromMutableChat(chat *tg.MutableChat, actorUserID int64, expectedInviteeUserIDs []int64) (payload.ChatParticipantsChangedFactV1, error) {
	if chat == nil {
		return payload.ChatParticipantsChangedFactV1{}, fmt.Errorf("mutable chat is nil")
	}
	if chat.Chat == nil {
		return payload.ChatParticipantsChangedFactV1{}, fmt.Errorf("mutable chat has no chat data")
	}
	if len(chat.ChatParticipants) == 0 {
		return payload.ChatParticipantsChangedFactV1{}, fmt.Errorf("mutable chat has no participants: chat_id=%d", chat.Chat.Id)
	}
	version := chat.Chat.Version
	if version <= 0 {
		version = 1
	}

	expectedInvitees := make(map[int64]struct{}, len(expectedInviteeUserIDs))
	for _, userID := range expectedInviteeUserIDs {
		expectedInvitees[userID] = struct{}{}
	}
	seenParticipants := make(map[int64]struct{}, len(chat.ChatParticipants))
	actorFound := false
	participants := make([]payload.ChatParticipantFactV1, 0, len(chat.ChatParticipants))
	for i, participant := range chat.ChatParticipants {
		if participant == nil {
			return payload.ChatParticipantsChangedFactV1{}, fmt.Errorf("mutable chat participant[%d] is nil: chat_id=%d", i, chat.Chat.Id)
		}
		if participant.ChatId != 0 && participant.ChatId != chat.Chat.Id {
			return payload.ChatParticipantsChangedFactV1{}, fmt.Errorf("mutable chat participant[%d] chat_id=%d does not match chat_id=%d", i, participant.ChatId, chat.Chat.Id)
		}
		if participant.UserId <= 0 {
			return payload.ChatParticipantsChangedFactV1{}, fmt.Errorf("mutable chat participant[%d] has invalid user_id=%d: chat_id=%d", i, participant.UserId, chat.Chat.Id)
		}
		if _, ok := seenParticipants[participant.UserId]; ok {
			return payload.ChatParticipantsChangedFactV1{}, fmt.Errorf("mutable chat has duplicate participant user_id=%d: chat_id=%d", participant.UserId, chat.Chat.Id)
		}
		seenParticipants[participant.UserId] = struct{}{}
		if participant.State != chatpb.ChatMemberStateNormal {
			return payload.ChatParticipantsChangedFactV1{}, fmt.Errorf("mutable chat participant[%d] has non-normal state=%d: chat_id=%d user_id=%d", i, participant.State, chat.Chat.Id, participant.UserId)
		}
		role, err := chatParticipantFactRole(participant.ParticipantType)
		if err != nil {
			return payload.ChatParticipantsChangedFactV1{}, fmt.Errorf("mutable chat participant[%d] role: chat_id=%d user_id=%d: %w", i, chat.Chat.Id, participant.UserId, err)
		}
		if participant.UserId == actorUserID {
			actorFound = true
			if role != "creator" {
				return payload.ChatParticipantsChangedFactV1{}, fmt.Errorf("mutable chat actor user_id=%d has role=%s, want creator: chat_id=%d", actorUserID, role, chat.Chat.Id)
			}
		}
		delete(expectedInvitees, participant.UserId)
		date, err := tg.DateInt32FromUnixSeconds(participant.Date)
		if err != nil {
			return payload.ChatParticipantsChangedFactV1{}, fmt.Errorf("mutable chat participant[%d] date: chat_id=%d user_id=%d: %w", i, chat.Chat.Id, participant.UserId, err)
		}
		participants = append(participants, payload.ChatParticipantFactV1{
			UserID:        participant.UserId,
			Role:          role,
			InviterUserID: participant.InviterUserId,
			Date:          date,
		})
	}
	if !actorFound {
		return payload.ChatParticipantsChangedFactV1{}, fmt.Errorf("mutable chat is missing actor user_id=%d: chat_id=%d", actorUserID, chat.Chat.Id)
	}
	for userID := range expectedInvitees {
		return payload.ChatParticipantsChangedFactV1{}, fmt.Errorf("mutable chat is missing invitee user_id=%d: chat_id=%d", userID, chat.Chat.Id)
	}

	return payload.ChatParticipantsChangedFactV1{
		SchemaVersion: 1,
		ChatID:        chat.Chat.Id,
		ActorUserID:   actorUserID,
		Version:       version,
		Participants:  participants,
	}, nil
}

func chatParticipantFactRole(participantType int32) (string, error) {
	switch participantType {
	case chatpb.ChatMemberCreator:
		return "creator", nil
	case chatpb.ChatMemberAdmin:
		return "admin", nil
	case chatpb.ChatMemberNormal:
		return "member", nil
	default:
		return "", fmt.Errorf("unsupported participant_type=%d", participantType)
	}
}

func metadataClientMsgID(md *metadata.RpcMetadata) int64 {
	if md == nil {
		return 0
	}
	return md.ClientMsgId
}

func createChatServiceMessageRandomID(actorUserID, clientMsgID int64) int64 {
	if clientMsgID == 0 {
		return normalizeCreateChatServiceMessageRandomID(rand.Int63())
	}
	sum := sha256.Sum256([]byte(fmt.Sprintf("create_chat:%d:%d:service_message", actorUserID, clientMsgID)))
	return createChatRandomIDFromDigest(sum[:])
}

func normalizeCreateChatServiceMessageRandomID(randomID int64) int64 {
	if randomID == 0 {
		return 1
	}
	return randomID
}

func createChatRandomIDFromDigest(digest []byte) int64 {
	if len(digest) < 8 {
		return 1
	}
	randomID := int64(binary.BigEndian.Uint64(digest[:8]) & uint64(^uint64(0)>>1))
	if randomID == 0 {
		return 1
	}
	return randomID
}
