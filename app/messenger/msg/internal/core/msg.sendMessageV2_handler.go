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
	"errors"
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
	outboxTLMessage, err := outboxTLMessage(outbox)
	if err != nil {
		return nil, err
	}
	messageText := outboxTLMessage.Message
	replyToCanonicalMessageID, err := c.replyToCanonicalMessageID(in, outboxTLMessage)
	if err != nil {
		return nil, err
	}
	var sourcePermAuthKeyID int64
	if in.SourcePermAuthKeyId != nil {
		sourcePermAuthKeyID = *in.SourcePermAuthKeyId
	}
	var clearDraftBeforeDate int32
	if in.ClearDraftBeforeDate != nil {
		clearDraftBeforeDate = *in.ClearDraftBeforeDate
	}
	_, requestHash, err := marshalSendRequest(in.UserId, in.PeerType, in.PeerId, outbox.RandomId, messageText, replyToCanonicalMessageID, in.ClearDraft, sourcePermAuthKeyID, clearDraftBeforeDate)
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

	messageDate := time.Now().UTC().Unix()
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
	if canonical.CreatedNew || sendState.Status < repository.SendStateStatusCanonical {
		if err := c.svcCtx.Repo.MarkCanonicalCreated(c.ctx, canonical.SendStateID, canonical.CanonicalMessageID, canonical.PeerSeq); err != nil {
			return nil, err
		}
	}

	senderOperationID := payload.SenderOperationID(canonical.CanonicalMessageID, in.UserId)
	var senderResult *userupdates.UserOperationResult
	if sendState.Status >= repository.SendStateStatusSenderCommitted {
		senderResult = senderResultFromSendState(sendState)
	} else {
		var senderHash []byte
		senderResult, senderHash, err = c.processSenderOperation(in, canonical, senderOperationID, messageText, replyToCanonicalMessageID)
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
	}

	if in.UserId != in.PeerId && sendState.Status < repository.SendStateStatusReceiverAcked {
		receiverOp, err := buildReceiverOperation(in, canonical, messageText, replyToCanonicalMessageID)
		if err != nil {
			return nil, err
		}
		if c.svcCtx.ReceiverPublisher == nil {
			return nil, msg.ErrReceiverBackpressure
		}
		ack, err := c.svcCtx.ReceiverPublisher.Publish(c.ctx, receiverOp)
		if err != nil {
			c.Logger.Errorf("msg.sendMessageV2 - receiver operation publish failed: operation_id=%s err=%v", receiverOp.OperationID, err)
			return nil, msg.ErrReceiverBackpressure
		}
		c.Logger.Debugf(
			"msg.sendMessageV2 - receiver operation published: operation_id=%s topic=%s partition=%d offset=%d",
			receiverOp.OperationID,
			ack.Topic,
			ack.Partition,
			ack.Offset,
		)
	}
	if sendState.Status < repository.SendStateStatusReceiverAcked {
		if err := c.svcCtx.Repo.MarkReceiverOpsAcked(c.ctx, canonical.SendStateID, 0); err != nil {
			return nil, err
		}
	}
	if sendState.Status < repository.SendStateStatusCompleted {
		if err := c.svcCtx.Repo.MarkCompleted(c.ctx, canonical.SendStateID); err != nil {
			return nil, err
		}
	}

	return shortSentMessage(canonical, senderResult)
}

func senderResultFromSendState(sendState *repository.SendState) *userupdates.UserOperationResult {
	if sendState == nil {
		return nil
	}
	return userupdates.MakeTLUserOperationResult(&userupdates.TLUserOperationResult{
		UserId:              sendState.SenderUserID,
		OperationId:         sendState.SenderOperationID,
		Status:              1,
		Pts:                 sendState.SenderPTS,
		PtsCount:            sendState.SenderPTSCount,
		CurrentPts:          sendState.SenderPTS,
		ResponsePayload:     sendState.SenderUpdatePayload,
		ResponsePayloadHash: sendState.SenderUpdatePayloadHash,
	})
}

type sendRequestPayloadV1 struct {
	SchemaVersion             int    `json:"schema_version"`
	SenderUserID              int64  `json:"sender_user_id"`
	PeerType                  int32  `json:"peer_type"`
	PeerID                    int64  `json:"peer_id"`
	ClientRandomID            int64  `json:"client_random_id"`
	MessageText               string `json:"message_text"`
	ReplyToCanonicalMessageID int64  `json:"reply_to_canonical_message_id,omitempty"`
	ClearDraft                bool   `json:"clear_draft,omitempty"`
	SourcePermAuthKeyID       int64  `json:"source_perm_auth_key_id,omitempty"`
	ClearDraftBeforeDate      int32  `json:"clear_draft_before_date,omitempty"`
}

func marshalSendRequest(senderUserID int64, peerType int32, peerID int64, randomID int64, text string, replyToCanonicalMessageID int64, clearDraft bool, sourcePermAuthKeyID int64, clearDraftBeforeDate int32) ([]byte, []byte, error) {
	body, err := json.Marshal(sendRequestPayloadV1{
		SchemaVersion:             payload.MessageOperationSchemaVersion,
		SenderUserID:              senderUserID,
		PeerType:                  peerType,
		PeerID:                    peerID,
		ClientRandomID:            randomID,
		MessageText:               text,
		ReplyToCanonicalMessageID: replyToCanonicalMessageID,
		ClearDraft:                clearDraft,
		SourcePermAuthKeyID:       sourcePermAuthKeyID,
		ClearDraftBeforeDate:      clearDraftBeforeDate,
	})
	if err != nil {
		return nil, nil, fmt.Errorf("%w: marshal send request: %v", msg.ErrMsgStorage, err)
	}
	return body, payload.HashBytes(body), nil
}

func outboxTLMessage(outbox *msg.TLOutboxMessage) (*tg.TLMessage, error) {
	if outbox == nil || outbox.Message == nil {
		return nil, fmt.Errorf("%w: missing outbox message", msg.ErrSendStateConflict)
	}
	message, ok := outbox.Message.(*tg.TLMessage)
	if !ok {
		return nil, fmt.Errorf("%w: first slice only supports text message", msg.ErrSendStateConflict)
	}
	return message, nil
}

func (c *MsgCore) replyToCanonicalMessageID(in *msg.TLMsgSendMessageV2, message *tg.TLMessage) (int64, error) {
	if message == nil || message.ReplyTo == nil {
		return 0, nil
	}
	replyHeader, ok := message.ReplyTo.(*tg.TLMessageReplyHeader)
	if !ok || replyHeader.ReplyToMsgId == nil || *replyHeader.ReplyToMsgId <= 0 {
		return 0, msg.ErrReplyToInvalid
	}
	replyTo, err := c.svcCtx.Repo.GetCanonicalMessageByPeerSeq(c.ctx, in.UserId, in.PeerType, in.PeerId, int64(*replyHeader.ReplyToMsgId))
	if err != nil {
		if errors.Is(err, msg.ErrSendStateConflict) {
			return 0, msg.ErrReplyToInvalid
		}
		return 0, err
	}
	if replyTo == nil || replyTo.CanonicalMessageID == 0 {
		return 0, msg.ErrReplyToInvalid
	}
	return replyTo.CanonicalMessageID, nil
}

func (c *MsgCore) processSenderOperation(in *msg.TLMsgSendMessageV2, canonical *repository.CanonicalMessageResult, operationID string, text string, replyToCanonicalMessageID int64) (*userupdates.UserOperationResult, []byte, error) {
	var sourcePermAuthKeyID int64
	if in.SourcePermAuthKeyId != nil {
		sourcePermAuthKeyID = *in.SourcePermAuthKeyId
	}
	var clearDraftBeforeDate int32
	if in.ClearDraftBeforeDate != nil {
		clearDraftBeforeDate = *in.ClearDraftBeforeDate
	}
	body, _, hashBytes, err := buildMessageOperationPayload(in.UserId, in.PeerId, in.PeerId, true, canonical, text, replyToCanonicalMessageID, in.ClearDraft, sourcePermAuthKeyID, clearDraftBeforeDate)
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
			CanonicalDate:        int64Ptr(canonical.MessageDate),
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
		SenderUpdatePayloadHash:   result.ResponsePayloadHash,
	})
}

func buildReceiverOperation(in *msg.TLMsgSendMessageV2, canonical *repository.CanonicalMessageResult, text string, replyToCanonicalMessageID int64) (repository.ReceiverOperation, error) {
	operationID := payload.ReceiverOperationID(canonical.CanonicalMessageID, in.PeerId)
	body, hashHex, _, err := buildMessageOperationPayload(in.UserId, in.PeerId, in.UserId, false, canonical, text, replyToCanonicalMessageID, false, 0, 0)
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

func buildMessageOperationPayload(fromUserID int64, toUserID int64, peerID int64, out bool, canonical *repository.CanonicalMessageResult, text string, replyToCanonicalMessageID int64, clearDraft bool, sourcePermAuthKeyID int64, clearDraftBeforeDate int32) ([]byte, []byte, []byte, error) {
	date, err := msgDateInt32FromUnixSeconds(canonical.MessageDate, "message date")
	if err != nil {
		return nil, nil, nil, err
	}
	body, err := json.Marshal(payload.MessageOperationV1{
		SchemaVersion:             payload.MessageOperationSchemaVersion,
		OperationKind:             payload.OperationKindSendMessage,
		CanonicalMessageID:        canonical.CanonicalMessageID,
		PeerType:                  payload.PeerTypeUser,
		PeerID:                    peerID,
		PeerSeq:                   canonical.PeerSeq,
		FromUserID:                fromUserID,
		ToUserID:                  toUserID,
		Date:                      date,
		Out:                       out,
		MessageText:               text,
		ReplyToCanonicalMessageID: replyToCanonicalMessageID,
		ClearDraft:                clearDraft,
		SourcePermAuthKeyID:       sourcePermAuthKeyID,
		ClearDraftBeforeDate:      clearDraftBeforeDate,
	})
	if err != nil {
		return nil, nil, nil, fmt.Errorf("%w: marshal message operation from_user_id=%d peer_id=%d", msg.ErrMsgStorage, fromUserID, peerID)
	}
	hashBytes := payload.HashBytes(body)
	return body, hashBytes, hashBytes, nil
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
	date, err := msgDateInt32FromUnixSeconds(canonical.MessageDate, "short sent message date")
	if err != nil {
		return nil, err
	}
	return tg.MakeTLUpdateShortSentMessage(&tg.TLUpdateShortSentMessage{
		Out:      true,
		Id:       peerSeq,
		Pts:      pts,
		PtsCount: result.PtsCount,
		Date:     date,
	}).ToUpdates(), nil
}

func int64ToInt32(v int64, field string) (int32, error) {
	if v < math.MinInt32 || v > math.MaxInt32 {
		return 0, fmt.Errorf("%w: %s out of int32 range", msg.ErrSenderSyncFailed, field)
	}
	return int32(v), nil
}

func msgDateInt32FromUnixSeconds(seconds int64, field string) (int32, error) {
	date, err := tg.DateInt32FromUnixSeconds(seconds)
	if err != nil {
		return 0, fmt.Errorf("%w: convert %s: %v", msg.ErrMsgStorage, field, err)
	}
	return date, nil
}

func int64Ptr(v int64) *int64 {
	return &v
}
