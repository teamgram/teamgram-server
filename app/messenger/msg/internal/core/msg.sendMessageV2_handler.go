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
	"encoding/hex"
	"encoding/json"
	"fmt"
	"math"
	"time"

	"github.com/teamgram/teamgram-server/v2/app/messenger/msg/internal/repository"
	"github.com/teamgram/teamgram-server/v2/app/messenger/msg/msg"
	"github.com/teamgram/teamgram-server/v2/app/messenger/userupdates/payload"
	"github.com/teamgram/teamgram-server/v2/app/messenger/userupdates/userupdates"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

// MsgSendMessageV2
// msg.sendMessageV2 user_id:long auth_key_id:long peer_type:int peer_id:long message:Vector<OutboxMessage> = Updates;
func (c *MsgCore) MsgSendMessageV2(in *msg.TLMsgSendMessageV2) (*tg.Updates, error) {
	if in == nil || len(in.Message) != 1 || in.Message[0] == nil {
		return nil, fmt.Errorf("%w: first slice requires exactly one message", msg.ErrSendStateConflict)
	}
	if in.PeerType != payload.PeerTypeUser {
		return nil, fmt.Errorf("%w: first slice only supports user peer", msg.ErrSendStateConflict)
	}
	outbox := in.Message[0]
	messageText, err := outboxMessageText(outbox)
	if err != nil {
		return nil, err
	}
	_, requestHash, err := marshalSendRequest(in.UserId, in.PeerType, in.PeerId, outbox.RandomId, messageText)
	if err != nil {
		return nil, err
	}

	sendState, err := c.svcCtx.Repo.CreateOrLoadSendState(c.ctx, repository.CreateSendStateInput{
		SenderUserID:                in.UserId,
		PeerType:                    in.PeerType,
		PeerID:                      in.PeerId,
		ClientRandomID:              outbox.RandomId,
		RequestPayloadSchemaVersion: payload.MessageOperationSchemaVersion,
		RequestPayloadHash:          requestHash,
	})
	if err != nil {
		return nil, err
	}

	messageDate := int32(time.Now().Unix())
	canonical, err := c.svcCtx.Repo.CreateOrGetByClientRandom(c.ctx, repository.CreateCanonicalMessageInput{
		SendStateID:        sendState.SendStateID,
		SenderUserID:       in.UserId,
		PeerType:           in.PeerType,
		PeerID:             in.PeerId,
		ClientRandomID:     outbox.RandomId,
		RequestPayloadHash: requestHash,
		MessageText:        messageText,
		MessageDate:        messageDate,
	})
	if err != nil {
		return nil, err
	}
	if err := c.svcCtx.Repo.MarkCanonicalCreated(c.ctx, canonical.SendStateID, canonical.CanonicalMessageID, canonical.PeerSeq); err != nil {
		return nil, err
	}

	senderOperationID := payload.SenderOperationID(canonical.CanonicalMessageID, in.UserId)
	senderResult, senderHash, err := c.processSenderOperation(in, canonical, senderOperationID, messageText)
	if err != nil {
		return nil, err
	}
	if err := c.markSenderCommitted(canonical, senderOperationID, senderResult); err != nil {
		c.Logger.Errorf("msg.sendMessageV2 - sender commit recovery: send_state_id=%d canonical_message_id=%d operation_id=%s err=%v", canonical.SendStateID, canonical.CanonicalMessageID, senderOperationID, err)
		recovered, recoverErr := c.recoverSenderOperation(in.UserId, senderOperationID, senderHash)
		if recoverErr != nil {
			_ = c.svcCtx.Repo.MarkRetryableFailure(c.ctx, repository.MarkRetryableFailureInput{SendStateID: canonical.SendStateID, LastErrorCode: "sender_commit_recovery_failed", LastErrorMessage: "sender commit recovery failed"})
			return nil, fmt.Errorf("%w: %v", msg.ErrSenderSyncFailed, recoverErr)
		}
		if err := c.markSenderCommitted(canonical, senderOperationID, recovered); err != nil {
			_ = c.svcCtx.Repo.MarkRetryableFailure(c.ctx, repository.MarkRetryableFailureInput{SendStateID: canonical.SendStateID, LastErrorCode: "sender_commit_failed", LastErrorMessage: "sender commit failed"})
			return nil, fmt.Errorf("%w: %v", msg.ErrSenderSyncFailed, err)
		}
		senderResult = recovered
	}

	receiverOp, err := buildReceiverOperation(in, canonical, messageText)
	if err != nil {
		return nil, err
	}
	if c.svcCtx.ReceiverPublisher == nil {
		return nil, msg.ErrReceiverBackpressure
	}
	if err := c.svcCtx.ReceiverPublisher.Publish(c.ctx, receiverOp); err != nil {
		return nil, fmt.Errorf("%w: %v", msg.ErrReceiverBackpressure, err)
	}
	if err := c.svcCtx.Repo.MarkReceiverOpsAcked(c.ctx, canonical.SendStateID, 0); err != nil {
		return nil, err
	}
	if err := c.svcCtx.Repo.MarkCompleted(c.ctx, canonical.SendStateID); err != nil {
		return nil, err
	}

	return shortSentMessage(canonical, senderResult)
}

type sendRequestPayloadV1 struct {
	SchemaVersion  int    `json:"schema_version"`
	SenderUserID   int64  `json:"sender_user_id"`
	PeerType       int32  `json:"peer_type"`
	PeerID         int64  `json:"peer_id"`
	ClientRandomID int64  `json:"client_random_id"`
	MessageText    string `json:"message_text"`
}

func marshalSendRequest(senderUserID int64, peerType int32, peerID int64, randomID int64, text string) ([]byte, string, error) {
	body, err := json.Marshal(sendRequestPayloadV1{
		SchemaVersion:  payload.MessageOperationSchemaVersion,
		SenderUserID:   senderUserID,
		PeerType:       peerType,
		PeerID:         peerID,
		ClientRandomID: randomID,
		MessageText:    text,
	})
	if err != nil {
		return nil, "", fmt.Errorf("%w: marshal send request: %v", msg.ErrMsgStorage, err)
	}
	return body, payload.HashBytes(body), nil
}

func outboxMessageText(outbox *msg.TLOutboxMessage) (string, error) {
	if outbox == nil || outbox.Message == nil {
		return "", fmt.Errorf("%w: missing outbox message", msg.ErrSendStateConflict)
	}
	message, ok := outbox.Message.(*tg.TLMessage)
	if !ok {
		return "", fmt.Errorf("%w: first slice only supports text message", msg.ErrSendStateConflict)
	}
	return message.Message, nil
}

func (c *MsgCore) processSenderOperation(in *msg.TLMsgSendMessageV2, canonical *repository.CanonicalMessageResult, operationID string, text string) (*userupdates.UserOperationResult, []byte, error) {
	body, _, hashBytes, err := buildMessageOperationPayload(in.UserId, in.UserId, in.PeerId, in.PeerId, true, canonical, text)
	if err != nil {
		return nil, nil, err
	}
	if c.svcCtx.UserUpdates == nil {
		return nil, nil, msg.ErrSenderSyncFailed
	}
	route := payload.RouteUser(in.UserId)
	authKeyID := in.AuthKeyId
	result, err := c.svcCtx.UserUpdates.UserupdatesProcessUserOperation(c.ctx, &userupdates.TLUserupdatesProcessUserOperation{
		Operation: userupdates.MakeTLUserOperation(&userupdates.TLUserOperation{
			UserId:               in.UserId,
			BucketId:             int32(route.BucketID),
			PartitionId:          int32(route.ReceiverPartitionID),
			OperationId:          operationID,
			OpType:               payload.OpTypeSendMessage,
			OpSource:             0,
			ActorUserId:          in.UserId,
			AuthKeyId:            &authKeyID,
			AuthKeyIdExclude:     &authKeyID,
			PeerType:             in.PeerType,
			PeerId:               in.PeerId,
			CanonicalMessageId:   &canonical.CanonicalMessageID,
			CanonicalPeerSeq:     &canonical.PeerSeq,
			CanonicalDate:        int64Ptr(int64(canonical.MessageDate)),
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

func (c *MsgCore) recoverSenderOperation(userID int64, operationID string, payloadHash []byte) (*userupdates.UserOperationResult, error) {
	if c.svcCtx.UserUpdates == nil {
		return nil, msg.ErrSenderSyncFailed
	}
	return c.svcCtx.UserUpdates.UserupdatesGetOperationResult(c.ctx, &userupdates.TLUserupdatesGetOperationResult{
		UserId:      userID,
		OperationId: operationID,
		PayloadHash: payloadHash,
	})
}

func (c *MsgCore) markSenderCommitted(canonical *repository.CanonicalMessageResult, operationID string, result *userupdates.UserOperationResult) error {
	if result == nil {
		return msg.ErrSenderSyncFailed
	}
	return c.svcCtx.Repo.MarkSenderCommitted(c.ctx, repository.MarkSenderCommittedInput{
		SendStateID:               canonical.SendStateID,
		SenderOperationID:         operationID,
		SenderPTS:                 result.Pts,
		SenderPTSCount:            result.PtsCount,
		SenderUpdateSchemaVersion: payload.OperationResponseSchemaVersion,
		SenderUpdatePayload:       result.ResponsePayload,
		SenderUpdatePayloadHash:   hex.EncodeToString(result.ResponsePayloadHash),
	})
}

func buildReceiverOperation(in *msg.TLMsgSendMessageV2, canonical *repository.CanonicalMessageResult, text string) (repository.ReceiverOperation, error) {
	operationID := payload.ReceiverOperationID(canonical.CanonicalMessageID, in.PeerId)
	body, hashHex, _, err := buildMessageOperationPayload(in.PeerId, in.UserId, in.PeerId, in.UserId, false, canonical, text)
	if err != nil {
		return repository.ReceiverOperation{}, err
	}
	route := payload.RouteUser(in.PeerId)
	return repository.ReceiverOperation{
		UserID:       in.PeerId,
		BucketID:     int32(route.BucketID),
		PartitionID:  int32(route.ReceiverPartitionID),
		OperationID:  operationID,
		OpType:       payload.OpTypeSendMessage,
		PeerType:     in.PeerType,
		PeerID:       in.UserId,
		PayloadCodec: payload.PayloadCodecJSON,
		Payload:      body,
		PayloadHash:  hashHex,
	}, nil
}

func buildMessageOperationPayload(userID int64, fromUserID int64, toUserID int64, peerID int64, out bool, canonical *repository.CanonicalMessageResult, text string) ([]byte, string, []byte, error) {
	body, err := json.Marshal(payload.MessageOperationV1{
		SchemaVersion:      payload.MessageOperationSchemaVersion,
		OperationKind:      payload.OperationKindSendMessage,
		CanonicalMessageID: canonical.CanonicalMessageID,
		PeerType:           payload.PeerTypeUser,
		PeerID:             peerID,
		PeerSeq:            canonical.PeerSeq,
		FromUserID:         fromUserID,
		ToUserID:           toUserID,
		Date:               canonical.MessageDate,
		Out:                out,
		MessageText:        text,
	})
	if err != nil {
		return nil, "", nil, fmt.Errorf("%w: marshal message operation user_id=%d", msg.ErrMsgStorage, userID)
	}
	hashHex := payload.HashBytes(body)
	hashBytes, err := hex.DecodeString(hashHex)
	if err != nil {
		return nil, "", nil, fmt.Errorf("%w: decode operation hash", msg.ErrMsgStorage)
	}
	return body, hashHex, hashBytes, nil
}

func shortSentMessage(canonical *repository.CanonicalMessageResult, result *userupdates.UserOperationResult) (*tg.Updates, error) {
	if canonical == nil || result == nil {
		return nil, msg.ErrSenderSyncFailed
	}
	peerSeq, err := int64ToInt32(canonical.PeerSeq, "peer seq")
	if err != nil {
		return nil, err
	}
	pts, err := int64ToInt32(result.Pts, "pts")
	if err != nil {
		return nil, err
	}
	return tg.MakeTLUpdateShortSentMessage(&tg.TLUpdateShortSentMessage{
		Out:      true,
		Id:       peerSeq,
		Pts:      pts,
		PtsCount: result.PtsCount,
		Date:     canonical.MessageDate,
	}).ToUpdates(), nil
}

func int64ToInt32(v int64, field string) (int32, error) {
	if v < math.MinInt32 || v > math.MaxInt32 {
		return 0, fmt.Errorf("%w: %s out of int32 range", msg.ErrSenderSyncFailed, field)
	}
	return int32(v), nil
}

func int64Ptr(v int64) *int64 {
	return &v
}
