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
	normalized, err := normalizeOutboxMessage(normalizeOutboxInput{
		Ctx:          c.ctx,
		SenderUserID: in.UserId,
		PeerType:     in.PeerType,
		PeerID:       in.PeerId,
		Outbox:       outbox,
		Repo:         c.svcCtx.Repo,
	})
	if err != nil {
		return nil, err
	}
	sideEffects := batchSideEffectsFromRequest(in)
	_, requestHash, err := marshalSendRequestV3(normalized, sideEffects)
	if err != nil {
		return nil, err
	}

	sendState, err := c.svcCtx.Repo.CreateOrLoadSendState(c.ctx, repository.CreateSendStateInput{
		SenderUserID:                in.UserId,
		PeerType:                    in.PeerType,
		PeerID:                      in.PeerId,
		ClientRandomID:              normalized.RandomID,
		RequestPayloadSchemaVersion: payload.MessageOperationSchemaVersionV3,
		RequestPayloadHash:          requestHash,
		MessageText:                 normalized.MessageText,
		ReplyToCanonicalMessageID:   normalized.ReplyToCanonicalMessageID,
	})
	if err != nil {
		return nil, err
	}

	canonicalPayloads, err := normalizedCanonicalPayloads(normalized)
	if err != nil {
		return nil, err
	}
	messageDate := time.Now().UTC().Unix()
	canonical, err := c.svcCtx.Repo.CreateOrGetByClientRandom(c.ctx, repository.CreateCanonicalMessageInput{
		SendStateID:                  sendState.SendStateID,
		SenderUserID:                 in.UserId,
		PeerType:                     in.PeerType,
		PeerID:                       in.PeerId,
		ClientRandomID:               normalized.RandomID,
		RequestPayloadHash:           requestHash,
		MessageText:                  normalized.MessageText,
		MessageDate:                  messageDate,
		EntitiesPayloadSchemaVersion: canonicalPayloads.EntitiesPayloadSchemaVersion,
		EntitiesPayload:              canonicalPayloads.EntitiesPayload,
		MediaRefSchemaVersion:        canonicalPayloads.MediaRefSchemaVersion,
		MediaRefPayload:              canonicalPayloads.MediaRefPayload,
		MessageAttrsSchemaVersion:    canonicalPayloads.MessageAttrsSchemaVersion,
		MessageAttrsPayload:          canonicalPayloads.MessageAttrsPayload,
		ForwardRefSchemaVersion:      canonicalPayloads.ForwardRefSchemaVersion,
		ForwardRefPayload:            canonicalPayloads.ForwardRefPayload,
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
		senderResult, senderHash, err = c.processSenderOperation(in, canonical, senderOperationID, normalized, sideEffects)
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
		receiverOp, err := buildReceiverOperationEnvelope(in, canonical, normalized)
		if err != nil {
			return nil, err
		}
		ack, err := c.dispatchBrokerDurableAck(receiverOp)
		if err != nil {
			return nil, err
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

func batchSideEffectsFromRequest(in *msg.TLMsgSendMessageV2) batchSideEffects {
	if in == nil {
		return batchSideEffects{}
	}
	var sourcePermAuthKeyID int64
	if in.SourcePermAuthKeyId != nil {
		sourcePermAuthKeyID = *in.SourcePermAuthKeyId
	}
	var clearDraftBeforeDate int32
	if in.ClearDraftBeforeDate != nil {
		clearDraftBeforeDate = *in.ClearDraftBeforeDate
	}
	return batchSideEffects{
		ClearDraft:           in.ClearDraft,
		SourcePermAuthKeyID:  sourcePermAuthKeyID,
		ClearDraftBeforeDate: clearDraftBeforeDate,
	}
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

type resolvedReplyToMessage struct {
	CanonicalMessageID int64
	UserMessageID      int64
}

func (c *MsgCore) processSenderOperation(in *msg.TLMsgSendMessageV2, canonical *repository.CanonicalMessageResult, operationID string, normalized normalizedOutboxMessage, sideEffects batchSideEffects) (*userupdates.UserOperationResult, []byte, error) {
	body, _, hashBytes, err := buildNormalizedMessageOperationPayload(normalized, in.PeerId, in.PeerId, true, canonical, sideEffects)
	if err != nil {
		return nil, nil, err
	}
	authKeyID := in.AuthKeyId
	result, err := c.dispatchRequesterSync(OperationEnvelope{
		UserID:               in.UserId,
		OperationID:          operationID,
		OpType:               payload.OpTypeSendMessage,
		OperationKind:        payload.OperationKindSendMessage,
		ActorUserID:          in.UserId,
		AuthKeyID:            &authKeyID,
		AuthKeyIDExclude:     &authKeyID,
		PeerType:             in.PeerType,
		PeerID:               in.PeerId,
		CanonicalMessageID:   &canonical.CanonicalMessageID,
		CanonicalPeerSeq:     &canonical.PeerSeq,
		CanonicalDate:        int64Ptr(canonical.MessageDate),
		PayloadSchemaVersion: payload.MessageOperationSchemaVersionV3,
		PayloadCodec:         payload.PayloadCodecJSON,
		PayloadHash:          hashBytes,
		Payload:              body,
		DeliveryPolicy:       DeliveryPolicyRequesterSync,
	}, nil)
	if err != nil {
		return nil, nil, err
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

func buildReceiverOperationEnvelope(in *msg.TLMsgSendMessageV2, canonical *repository.CanonicalMessageResult, normalized normalizedOutboxMessage) (OperationEnvelope, error) {
	operationID := payload.ReceiverOperationID(canonical.CanonicalMessageID, in.PeerId)
	body, hashBytes, _, err := buildNormalizedMessageOperationPayload(normalized, in.PeerId, in.UserId, false, canonical, batchSideEffects{})
	if err != nil {
		return OperationEnvelope{}, err
	}
	return OperationEnvelope{
		UserID:               in.PeerId,
		OperationID:          operationID,
		OpType:               payload.OpTypeSendMessage,
		OperationKind:        payload.OperationKindSendMessage,
		ActorUserID:          in.UserId,
		PeerType:             in.PeerType,
		PeerID:               in.UserId,
		CanonicalMessageID:   &canonical.CanonicalMessageID,
		CanonicalPeerSeq:     &canonical.PeerSeq,
		CanonicalDate:        int64Ptr(canonical.MessageDate),
		PayloadSchemaVersion: payload.MessageOperationSchemaVersionV3,
		PayloadCodec:         payload.PayloadCodecJSON,
		Payload:              body,
		PayloadHash:          hashBytes,
		DeliveryPolicy:       DeliveryPolicyBrokerDurableAck,
	}, nil
}

func shortSentMessage(canonical *repository.CanonicalMessageResult, result *userupdates.UserOperationResult) (*tg.Updates, error) {
	if canonical == nil || result == nil {
		return nil, msg.ErrSenderSyncFailed
	}
	response, err := operationResponseV2(result, "sender")
	if err != nil {
		return nil, err
	}
	userMessageID, err := operationResponseUserMessageID(response, "sender")
	if err != nil {
		return nil, err
	}
	pts, err := int64ToInt32(response.Pts, "pts")
	if err != nil {
		return nil, err
	}
	date, err := msgDateInt32FromUnixSeconds(canonical.MessageDate, "short sent message date")
	if err != nil {
		return nil, err
	}
	return tg.MakeTLUpdateShortSentMessage(&tg.TLUpdateShortSentMessage{
		Out:      true,
		Id:       userMessageID,
		Pts:      pts,
		PtsCount: response.PtsCount,
		Date:     date,
	}).ToUpdates(), nil
}

func operationResponseV2(result *userupdates.UserOperationResult, operation string) (payload.OperationResponseV2, error) {
	if result == nil {
		return payload.OperationResponseV2{}, msg.ErrSenderSyncFailed
	}
	var response payload.OperationResponseV2
	if err := json.Unmarshal(result.ResponsePayload, &response); err != nil {
		return payload.OperationResponseV2{}, fmt.Errorf("%w: decode %s operation response: %v", msg.ErrMsgStorage, operation, err)
	}
	if response.SchemaVersion != payload.OperationResponseSchemaVersion {
		return payload.OperationResponseV2{}, fmt.Errorf("%w: unsupported %s operation response schema=%d", msg.ErrMsgStorage, operation, response.SchemaVersion)
	}
	return response, nil
}

func operationResponseUserMessageID(response payload.OperationResponseV2, operation string) (int32, error) {
	if response.UserMessageID <= 0 {
		return 0, fmt.Errorf("%w: %s operation missing user_message_id", msg.ErrMsgStorage, operation)
	}
	return int64ToInt32(response.UserMessageID, "user message id")
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
