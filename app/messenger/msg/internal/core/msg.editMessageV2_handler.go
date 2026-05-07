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
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

// MsgEditMessageV2
// msg.editMessageV2 user_id:long auth_key_id:long peer_type:int peer_id:long edit_type:int new_message:OutboxMessage dst_message:MessageBox = Updates;
func (c *MsgCore) MsgEditMessageV2(in *msg.TLMsgEditMessageV2) (*tg.Updates, error) {
	if in == nil || in.NewMessage == nil || in.DstMessage == nil {
		return nil, fmt.Errorf("%w: missing edit message input", msg.ErrSendStateConflict)
	}
	if in.PeerType != payload.PeerTypeUser {
		return nil, fmt.Errorf("%w: first slice only supports user peer", msg.ErrSendStateConflict)
	}
	newMessage, err := outboxTLMessage(in.NewMessage)
	if err != nil {
		return nil, err
	}
	if newMessage.Message == "" {
		return nil, msg.ErrSendStateConflict
	}
	peerSeq := int64(in.DstMessage.MessageId)
	if peerSeq <= 0 {
		return nil, msg.ErrSendStateConflict
	}

	editDate := int32(time.Now().Unix())
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

	senderResult, senderHash, err := c.processEditSenderOperation(in, edited)
	if err != nil {
		return nil, err
	}
	if in.UserId != in.PeerId {
		receiverOp, err := buildEditReceiverOperation(in, edited)
		if err != nil {
			return nil, err
		}
		if c.svcCtx.ReceiverPublisher == nil {
			return nil, msg.ErrReceiverBackpressure
		}
		if _, err := c.svcCtx.ReceiverPublisher.Publish(c.ctx, receiverOp); err != nil {
			c.Logger.Errorf("msg.editMessageV2 - receiver operation publish failed: operation_id=%s err=%v", receiverOp.OperationID, err)
			return nil, msg.ErrReceiverBackpressure
		}
	}

	return shortEditMessage(edited, senderResult, senderHash)
}

func (c *MsgCore) processEditSenderOperation(in *msg.TLMsgEditMessageV2, edited *repository.EditMessageResult) (*userupdates.UserOperationResult, []byte, error) {
	if c.svcCtx.UserUpdates == nil {
		return nil, nil, msg.ErrSenderSyncFailed
	}
	body, hashBytes, err := buildEditMessageOperationPayload(in.UserId, in.PeerId, in.PeerId, true, edited)
	if err != nil {
		return nil, nil, err
	}
	route := payload.RouteUser(in.UserId)
	authKeyID := in.AuthKeyId
	result, err := c.svcCtx.UserUpdates.UserupdatesProcessUserOperation(c.ctx, &userupdates.TLUserupdatesProcessUserOperation{
		Operation: userupdates.MakeTLUserOperation(&userupdates.TLUserOperation{
			UserId:               in.UserId,
			BucketId:             int32(route.BucketID),
			PartitionId:          int32(route.ReceiverPartitionID),
			OperationId:          editMessageOperationID(edited.CanonicalMessageID, edited.EditVersion, in.UserId),
			OpType:               payload.OpTypeSendMessage,
			ActorUserId:          in.UserId,
			AuthKeyId:            &authKeyID,
			AuthKeyIdExclude:     &authKeyID,
			PeerType:             in.PeerType,
			PeerId:               in.PeerId,
			CanonicalMessageId:   &edited.CanonicalMessageID,
			CanonicalPeerSeq:     &edited.PeerSeq,
			CanonicalDate:        int64Ptr(int64(edited.MessageDate)),
			PayloadSchemaVersion: payload.MessageOperationSchemaVersion,
			PayloadCodec:         payload.PayloadCodecJSON,
			PayloadHash:          hashBytes,
			Payload:              body,
		}),
	})
	if err != nil {
		return nil, nil, fmt.Errorf("%w: %v", msg.ErrSenderSyncFailed, err)
	}
	return result, hashBytes, nil
}

func buildEditReceiverOperation(in *msg.TLMsgEditMessageV2, edited *repository.EditMessageResult) (repository.ReceiverOperation, error) {
	body, hashBytes, err := buildEditMessageOperationPayload(in.UserId, in.PeerId, in.UserId, false, edited)
	if err != nil {
		return repository.ReceiverOperation{}, err
	}
	route := payload.RouteUser(in.PeerId)
	return repository.ReceiverOperation{
		UserID:       in.PeerId,
		BucketID:     int32(route.BucketID),
		PartitionID:  int32(route.ReceiverPartitionID),
		OperationID:  editMessageOperationID(edited.CanonicalMessageID, edited.EditVersion, in.PeerId),
		OpType:       payload.OpTypeSendMessage,
		PeerType:     in.PeerType,
		PeerID:       in.UserId,
		PayloadCodec: payload.PayloadCodecJSON,
		Payload:      body,
		PayloadHash:  hashBytes,
	}, nil
}

func buildEditMessageOperationPayload(fromUserID int64, toUserID int64, peerID int64, out bool, edited *repository.EditMessageResult) ([]byte, []byte, error) {
	body, err := json.Marshal(payload.MessageOperationV1{
		SchemaVersion:      payload.MessageOperationSchemaVersion,
		OperationKind:      payload.OperationKindEditMessage,
		CanonicalMessageID: edited.CanonicalMessageID,
		PeerType:           payload.PeerTypeUser,
		PeerID:             peerID,
		PeerSeq:            edited.PeerSeq,
		FromUserID:         fromUserID,
		ToUserID:           toUserID,
		Date:               edited.MessageDate,
		EditDate:           edited.EditDate,
		EditVersion:        edited.EditVersion,
		Out:                out,
		MessageText:        edited.MessageText,
	})
	if err != nil {
		return nil, nil, fmt.Errorf("%w: marshal edit message operation from_user_id=%d peer_id=%d", msg.ErrMsgStorage, fromUserID, peerID)
	}
	return body, payload.HashBytes(body), nil
}

func shortEditMessage(edited *repository.EditMessageResult, result *userupdates.UserOperationResult, _ []byte) (*tg.Updates, error) {
	if edited == nil || result == nil {
		return nil, msg.ErrSenderSyncFailed
	}
	peerSeq, err := int64ToInt32(edited.PeerSeq, "peer seq")
	if err != nil {
		return nil, err
	}
	pts, err := int64ToInt32(result.Pts, "pts")
	if err != nil {
		return nil, err
	}
	message := tg.MakeTLMessage(&tg.TLMessage{
		Out:      true,
		Id:       peerSeq,
		FromId:   tg.MakePeerUser(edited.FromUserID),
		PeerId:   tg.MakePeerUser(edited.PeerID),
		Date:     edited.MessageDate,
		Message:  edited.MessageText,
		EditDate: &edited.EditDate,
	})
	return tg.MakeTLUpdates(&tg.TLUpdates{
		Updates: []tg.UpdateClazz{tg.MakeTLUpdateEditMessage(&tg.TLUpdateEditMessage{
			Message:  message,
			Pts:      pts,
			PtsCount: result.PtsCount,
		})},
		Users: editMessageUsers(edited.FromUserID, edited.PeerID),
		Chats: []tg.ChatClazz{},
		Date:  edited.EditDate - 1,
		Seq:   0,
	}).ToUpdates(), nil
}

func editMessageOperationID(canonicalMessageID int64, editVersion int32, userID int64) string {
	return fmt.Sprintf("v1:msg:%d:edit:%d:%d", canonicalMessageID, editVersion, userID)
}

func editMessageUsers(fromUserID, peerID int64) []tg.UserClazz {
	if fromUserID == peerID {
		return []tg.UserClazz{tg.MakeTLUser(&tg.TLUser{Id: fromUserID})}
	}
	return []tg.UserClazz{
		tg.MakeTLUser(&tg.TLUser{Id: fromUserID}),
		tg.MakeTLUser(&tg.TLUser{Id: peerID}),
	}
}
