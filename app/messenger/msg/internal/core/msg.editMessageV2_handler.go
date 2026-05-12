// Copyright (c) 2026-present, The Teamgram Authors (https://teamgram.net).
//  All rights reserved.
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
	"fmt"
	"time"

	"github.com/teamgram/teamgram-server/v2/app/messenger/msg/internal/repository"
	"github.com/teamgram/teamgram-server/v2/app/messenger/msg/msg"
	"github.com/teamgram/teamgram-server/v2/app/messenger/userupdates/payload"
	"github.com/teamgram/teamgram-server/v2/app/messenger/userupdates/userupdates"
	chatpb "github.com/teamgram/teamgram-server/v2/app/service/biz/chat/chat"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

// MsgEditMessageV2
// msg.editMessageV2 user_id:long auth_key_id:long peer_type:int peer_id:long edit_type:int new_message:OutboxMessage dst_message:MessageBox = Updates;
func (c *MsgCore) MsgEditMessageV2(in *msg.TLMsgEditMessageV2) (*tg.Updates, error) {
	if in == nil || in.NewMessage == nil || in.DstMessage == nil {
		return nil, fmt.Errorf("%w: missing edit message input", msg.ErrSendStateConflict)
	}
	if in.PeerType != payload.PeerTypeUser && in.PeerType != payload.PeerTypeChat {
		return nil, fmt.Errorf("%w: unsupported edit peer type=%d", msg.ErrSendStateConflict, in.PeerType)
	}
	if in.PeerType == payload.PeerTypeChat {
		if err := c.checkChatAction(in.UserId, in.PeerId, chatpb.MessageActionEditOwnMessage, ""); err != nil {
			return nil, err
		}
	}
	newMessage, err := outboxTLMessage(in.NewMessage)
	if err != nil {
		return nil, err
	}
	if newMessage.Message == "" {
		return nil, msg.ErrSendStateConflict
	}
	if in.DstMessage.MessageId <= 0 {
		return nil, msg.ErrSendStateConflict
	}
	resolved, err := c.svcCtx.Repo.ResolveMessageID(c.ctx, in.UserId, in.PeerType, in.PeerId, int64(in.DstMessage.MessageId))
	if err != nil {
		return nil, err
	}
	if resolved == nil {
		return nil, msg.ErrSendStateConflict
	}
	peerSeq := resolved.PeerSeq

	editDate := time.Now().UTC().Unix()
	edited, err := c.svcCtx.Repo.EditCanonicalMessage(c.ctx, repository.EditCanonicalMessageInput{
		ActorUserID:     in.UserId,
		PeerType:        in.PeerType,
		PeerID:          in.PeerId,
		PeerSeq:         peerSeq,
		NewMessageText:  newMessage.Message,
		RequestEditDate: editDate,
	})
	if err != nil {
		return nil, err
	}
	if edited == nil {
		return nil, msg.ErrMsgStorage
	}

	effects, err := c.buildEditChatReceiverEffects(in, edited)
	if err != nil {
		return nil, err
	}
	senderResult, senderHash, err := c.processEditSenderOperation(in, edited, resolved.UserMessageID, effects)
	if err != nil {
		return nil, err
	}
	if in.PeerType == payload.PeerTypeUser && in.UserId != in.PeerId {
		receiverOp, err := buildEditReceiverOperationEnvelope(in, edited)
		if err != nil {
			return nil, err
		}
		if _, err := c.dispatchBrokerDurableAck(receiverOp); err != nil {
			return nil, err
		}
	}

	return shortEditMessage(edited, senderResult, senderHash, in.PeerType, in.PeerId)
}

func (c *MsgCore) processEditSenderOperation(in *msg.TLMsgEditMessageV2, edited *repository.EditMessageResult, userMessageID int64, effects []OperationEnvelope) (*userupdates.UserOperationResult, []byte, error) {
	body, hashBytes, err := buildEditMessageOperationPayload(in.UserId, in.PeerId, in.PeerType, in.PeerId, true, edited, userMessageID)
	if err != nil {
		return nil, nil, err
	}
	authKeyID := in.AuthKeyId
	result, err := c.dispatchRequesterSync(OperationEnvelope{
		UserID:               in.UserId,
		OperationID:          editMessageOperationID(edited.CanonicalMessageID, edited.EditVersion, in.UserId),
		OpType:               payload.OpTypeSendMessage,
		OperationKind:        payload.OperationKindEditMessage,
		ActorUserID:          in.UserId,
		AuthKeyID:            &authKeyID,
		AuthKeyIDExclude:     &authKeyID,
		PeerType:             in.PeerType,
		PeerID:               in.PeerId,
		CanonicalMessageID:   &edited.CanonicalMessageID,
		CanonicalPeerSeq:     &edited.PeerSeq,
		CanonicalDate:        int64Ptr(edited.MessageDate),
		PayloadSchemaVersion: payload.MessageOperationSchemaVersion,
		PayloadCodec:         payload.PayloadCodecJSON,
		PayloadHash:          hashBytes,
		Payload:              body,
		DeliveryPolicy:       DeliveryPolicyRequesterSync,
	}, effects)
	if err != nil {
		return nil, nil, err
	}
	return result, hashBytes, nil
}

func buildEditReceiverOperationEnvelope(in *msg.TLMsgEditMessageV2, edited *repository.EditMessageResult) (OperationEnvelope, error) {
	body, hashBytes, err := buildEditMessageOperationPayload(in.UserId, in.PeerId, in.PeerType, in.UserId, false, edited, 0)
	if err != nil {
		return OperationEnvelope{}, err
	}
	return OperationEnvelope{
		UserID:               in.PeerId,
		OperationID:          editMessageOperationID(edited.CanonicalMessageID, edited.EditVersion, in.PeerId),
		OpType:               payload.OpTypeSendMessage,
		OperationKind:        payload.OperationKindEditMessage,
		ActorUserID:          in.UserId,
		PeerType:             in.PeerType,
		PeerID:               in.UserId,
		CanonicalMessageID:   &edited.CanonicalMessageID,
		CanonicalPeerSeq:     &edited.PeerSeq,
		CanonicalDate:        int64Ptr(edited.MessageDate),
		PayloadSchemaVersion: payload.MessageOperationSchemaVersion,
		PayloadCodec:         payload.PayloadCodecJSON,
		Payload:              body,
		PayloadHash:          hashBytes,
		DeliveryPolicy:       DeliveryPolicyBrokerDurableAck,
	}, nil
}

func (c *MsgCore) buildEditChatReceiverEffects(in *msg.TLMsgEditMessageV2, edited *repository.EditMessageResult) ([]OperationEnvelope, error) {
	if in.PeerType != payload.PeerTypeChat {
		return nil, nil
	}
	receiverIDs, err := c.activeChatReceiverIDs(in.PeerId, in.UserId)
	if err != nil {
		return nil, err
	}
	effects := make([]OperationEnvelope, 0, len(receiverIDs))
	for _, receiverUserID := range receiverIDs {
		body, hashBytes, err := buildEditMessageOperationPayload(in.UserId, receiverUserID, payload.PeerTypeChat, in.PeerId, false, edited, 0)
		if err != nil {
			return nil, err
		}
		effects = append(effects, OperationEnvelope{
			UserID:               receiverUserID,
			OperationID:          editMessageOperationID(edited.CanonicalMessageID, edited.EditVersion, receiverUserID),
			OpType:               payload.OpTypeSendMessage,
			OperationKind:        payload.OperationKindEditMessage,
			ActorUserID:          in.UserId,
			PeerType:             payload.PeerTypeChat,
			PeerID:               in.PeerId,
			CanonicalMessageID:   &edited.CanonicalMessageID,
			CanonicalPeerSeq:     &edited.PeerSeq,
			CanonicalDate:        int64Ptr(edited.MessageDate),
			PayloadSchemaVersion: payload.MessageOperationSchemaVersion,
			PayloadCodec:         payload.PayloadCodecJSON,
			PayloadHash:          hashBytes,
			Payload:              body,
			DeliveryPolicy:       DeliveryPolicyDurableAsync,
		})
	}
	return effects, nil
}

func buildEditMessageOperationPayload(fromUserID int64, toUserID int64, peerType int32, peerID int64, out bool, edited *repository.EditMessageResult, userMessageID int64) ([]byte, []byte, error) {
	date, err := msgDateInt32FromUnixSeconds(edited.MessageDate, "edit message date")
	if err != nil {
		return nil, nil, err
	}
	editDate, err := msgDateInt32FromUnixSeconds(edited.EditDate, "edit date")
	if err != nil {
		return nil, nil, err
	}
	body, err := json.Marshal(payload.MessageOperationV1{
		SchemaVersion:      payload.MessageOperationSchemaVersion,
		OperationKind:      payload.OperationKindEditMessage,
		CanonicalMessageID: edited.CanonicalMessageID,
		PeerType:           peerType,
		PeerID:             peerID,
		PeerSeq:            edited.PeerSeq,
		FromUserID:         fromUserID,
		ToUserID:           toUserID,
		Date:               date,
		EditDate:           editDate,
		EditVersion:        edited.EditVersion,
		Out:                out,
		MessageText:        edited.MessageText,
		UserMessageID:      userMessageID,
	})
	if err != nil {
		return nil, nil, fmt.Errorf("%w: marshal edit message operation from_user_id=%d peer_id=%d", msg.ErrMsgStorage, fromUserID, peerID)
	}
	return body, payload.HashBytes(body), nil
}

func shortEditMessage(edited *repository.EditMessageResult, result *userupdates.UserOperationResult, _ []byte, peerType int32, peerID int64) (*tg.Updates, error) {
	if edited == nil || result == nil {
		return nil, msg.ErrSenderSyncFailed
	}
	response, err := operationResponseV2(result, "edit")
	if err != nil {
		return nil, err
	}
	userMessageID, err := operationResponseUserMessageID(response, "edit")
	if err != nil {
		return nil, err
	}
	pts, err := int64ToInt32(result.Pts, "pts")
	if err != nil {
		return nil, err
	}
	date, err := msgDateInt32FromUnixSeconds(edited.MessageDate, "edit update message date")
	if err != nil {
		return nil, err
	}
	editDate, err := msgDateInt32FromUnixSeconds(edited.EditDate, "edit update edit date")
	if err != nil {
		return nil, err
	}
	updateDate, err := msgDateInt32FromUnixSeconds(edited.EditDate-1, "edit updates date")
	if err != nil {
		return nil, err
	}
	message := tg.MakeTLMessage(&tg.TLMessage{
		Out:      true,
		Id:       userMessageID,
		FromId:   tg.MakePeerUser(edited.FromUserID),
		PeerId:   sentMessagePeerFromOptional(peerType, peerID),
		Date:     date,
		Message:  edited.MessageText,
		EditDate: &editDate,
	})
	return tg.MakeTLUpdates(&tg.TLUpdates{
		Updates: []tg.UpdateClazz{tg.MakeTLUpdateEditMessage(&tg.TLUpdateEditMessage{
			Message:  message,
			Pts:      pts,
			PtsCount: result.PtsCount,
		})},
		Users: []tg.UserClazz{},
		Chats: []tg.ChatClazz{},
		Date:  updateDate,
		Seq:   0,
	}).ToUpdates(), nil
}

func editMessageOperationID(canonicalMessageID int64, editVersion int32, userID int64) string {
	return fmt.Sprintf("v2:msg:%d:edit:%d:%d", canonicalMessageID, editVersion, userID)
}
