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

const maxSendMessageV2BatchSize = 100

// MsgSendMessageV2
// msg.sendMessageV2 user_id:long auth_key_id:long peer_type:int peer_id:long message:Vector<OutboxMessage> = Updates;
func (c *MsgCore) MsgSendMessageV2(in *msg.TLMsgSendMessageV2) (*tg.Updates, error) {
	if in == nil || len(in.Message) == 0 {
		return nil, fmt.Errorf("%w: first slice requires at least one message", msg.ErrSendStateConflict)
	}
	if len(in.Message) > maxSendMessageV2BatchSize {
		return nil, fmt.Errorf("%w: max=%d got=%d", msg.ErrBatchTooLarge, maxSendMessageV2BatchSize, len(in.Message))
	}
	if in.PeerType != payload.PeerTypeUser {
		return nil, fmt.Errorf("%w: first slice only supports user peer", msg.ErrSendStateConflict)
	}
	normalizedBatch := make([]normalizedOutboxMessage, 0, len(in.Message))
	for _, outbox := range in.Message {
		if outbox == nil {
			return nil, fmt.Errorf("%w: missing outbox message", msg.ErrSendStateConflict)
		}
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
		normalizedBatch = append(normalizedBatch, normalized)
	}
	sideEffects := batchSideEffectsFromRequest(in)
	if len(normalizedBatch) > 1 {
		return c.sendMessageV2Batch(in, normalizedBatch, sideEffects)
	}

	normalized := normalizedBatch[0]
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
	if isSendStateSenderCommitted(sendState) {
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

	if in.UserId != in.PeerId && needsReceiverAck(sendState.Status) {
		receiverOp, err := buildReceiverOperationEnvelope(in, canonical, normalized)
		if err != nil {
			return nil, err
		}
		ack, err := c.dispatchBrokerDurableAck(receiverOp)
		if err != nil {
			_ = c.svcCtx.Repo.MarkRetryableFailure(c.ctx, repository.MarkRetryableFailureInput{SendStateID: canonical.SendStateID, LastErrorCode: "receiver_ack_failed", LastErrorMessage: "receiver durable ack failed"})
			return nil, fmt.Errorf("%w: %w", msg.ErrSenderSyncFailed, err)
		}
		c.Logger.Debugf(
			"msg.sendMessageV2 - receiver operation published: operation_id=%s topic=%s partition=%d offset=%d",
			receiverOp.OperationID,
			ack.Topic,
			ack.Partition,
			ack.Offset,
		)
	}
	if needsReceiverAck(sendState.Status) {
		if err := c.svcCtx.Repo.MarkReceiverOpsAcked(c.ctx, canonical.SendStateID, 0); err != nil {
			return nil, err
		}
	}
	if !isCompleted(sendState.Status) {
		if err := c.svcCtx.Repo.MarkCompleted(c.ctx, canonical.SendStateID); err != nil {
			return nil, err
		}
	}

	if requiresFullSentUpdates(normalized) {
		return fullSentMessageUpdates(in.UserId, in.PeerId, []repository.CanonicalMessageResult{*canonical}, []*userupdates.UserOperationResult{senderResult}, []normalizedOutboxMessage{normalized})
	}
	return shortSentMessage(canonical, senderResult)
}

func (c *MsgCore) sendMessageV2Batch(in *msg.TLMsgSendMessageV2, normalizedBatch []normalizedOutboxMessage, sideEffects batchSideEffects) (*tg.Updates, error) {
	batchInput := repository.CreateCanonicalBatchInput{
		SenderUserID: in.UserId,
		PeerType:     in.PeerType,
		PeerID:       in.PeerId,
		Items:        make([]repository.CreateCanonicalBatchItem, 0, len(normalizedBatch)),
	}
	for _, normalized := range normalizedBatch {
		_, requestHash, err := marshalSendRequestV3(normalized, sideEffects)
		if err != nil {
			return nil, err
		}
		canonicalPayloads, err := normalizedCanonicalPayloads(normalized)
		if err != nil {
			return nil, err
		}
		batchInput.Items = append(batchInput.Items, repository.CreateCanonicalBatchItem{
			ClientRandomID:               normalized.RandomID,
			RequestPayloadSchemaVersion:  payload.MessageOperationSchemaVersionV3,
			RequestPayloadHash:           requestHash,
			MessageText:                  normalized.MessageText,
			MessageDate:                  time.Now().UTC().Unix(),
			EntitiesPayloadSchemaVersion: canonicalPayloads.EntitiesPayloadSchemaVersion,
			EntitiesPayload:              canonicalPayloads.EntitiesPayload,
			MediaRefSchemaVersion:        canonicalPayloads.MediaRefSchemaVersion,
			MediaRefPayload:              canonicalPayloads.MediaRefPayload,
			MessageAttrsSchemaVersion:    canonicalPayloads.MessageAttrsSchemaVersion,
			MessageAttrsPayload:          canonicalPayloads.MessageAttrsPayload,
			ForwardRefSchemaVersion:      canonicalPayloads.ForwardRefSchemaVersion,
			ForwardRefPayload:            canonicalPayloads.ForwardRefPayload,
		})
	}

	canonicalBatch, err := c.svcCtx.Repo.CreateOrGetCanonicalBatchByClientRandom(c.ctx, batchInput)
	if err != nil {
		return nil, err
	}
	if canonicalBatch == nil || len(canonicalBatch.Items) != len(normalizedBatch) {
		return nil, msg.ErrMsgStorage
	}

	results := make([]*userupdates.UserOperationResult, len(normalizedBatch))
	envelopes := make([]OperationEnvelope, 0, len(normalizedBatch))
	envelopeIndexes := make([]int, 0, len(normalizedBatch))
	for i := range normalizedBatch {
		canonical := &canonicalBatch.Items[i]
		operationID := payload.SenderOperationID(canonical.CanonicalMessageID, in.UserId)
		if isSenderCommitted(canonical) {
			results[i] = senderResultFromCanonical(canonical, operationID, in.UserId)
			continue
		}
		envelope, _, err := c.buildSenderOperationEnvelope(in, canonical, operationID, normalizedBatch[i], sideEffects)
		if err != nil {
			return nil, err
		}
		envelopes = append(envelopes, envelope)
		envelopeIndexes = append(envelopeIndexes, i)
	}
	dispatchedResults, err := c.dispatchRequesterBatchSync(envelopes)
	if err != nil {
		return nil, err
	}
	if len(dispatchedResults) != len(envelopeIndexes) {
		return nil, msg.ErrSenderSyncFailed
	}
	for i, result := range dispatchedResults {
		index := envelopeIndexes[i]
		canonical := &canonicalBatch.Items[index]
		operationID := payload.SenderOperationID(canonical.CanonicalMessageID, in.UserId)
		if err := c.markSenderCommitted(canonical, operationID, result); err != nil {
			_ = c.svcCtx.Repo.MarkRetryableFailure(c.ctx, repository.MarkRetryableFailureInput{SendStateID: canonical.SendStateID, LastErrorCode: "sender_batch_commit_failed", LastErrorMessage: "sender batch commit failed"})
			return nil, fmt.Errorf("%w: %v", msg.ErrSenderSyncFailed, err)
		}
		results[index] = result
	}
	if err := c.dispatchBatchReceiverOps(in, canonicalBatch.Items, normalizedBatch); err != nil {
		return nil, err
	}
	for i := range canonicalBatch.Items {
		canonical := &canonicalBatch.Items[i]
		if needsReceiverAck(canonical.SendStateStatus) {
			if err := c.svcCtx.Repo.MarkReceiverOpsAcked(c.ctx, canonical.SendStateID, 0); err != nil {
				return nil, err
			}
		}
		if !isCompleted(canonical.SendStateStatus) {
			if err := c.svcCtx.Repo.MarkCompleted(c.ctx, canonical.SendStateID); err != nil {
				return nil, err
			}
		}
	}
	return fullSentMessageUpdates(in.UserId, in.PeerId, canonicalBatch.Items, results, normalizedBatch)
}

func (c *MsgCore) dispatchBatchReceiverOps(in *msg.TLMsgSendMessageV2, canonicals []repository.CanonicalMessageResult, normalizedBatch []normalizedOutboxMessage) error {
	if in.UserId == in.PeerId {
		return nil
	}
	envelopes := make([]OperationEnvelope, 0, len(canonicals))
	sendStateIDs := make([]int64, 0, len(canonicals))
	for i := range canonicals {
		canonical := &canonicals[i]
		if !needsReceiverAck(canonical.SendStateStatus) {
			continue
		}
		envelope, err := buildReceiverOperationEnvelope(in, canonical, normalizedBatch[i])
		if err != nil {
			return err
		}
		envelopes = append(envelopes, envelope)
		sendStateIDs = append(sendStateIDs, canonical.SendStateID)
	}
	for _, envelope := range envelopes {
		if _, err := c.dispatchBrokerDurableAck(envelope); err != nil {
			for _, sendStateID := range sendStateIDs {
				_ = c.svcCtx.Repo.MarkRetryableFailure(c.ctx, repository.MarkRetryableFailureInput{SendStateID: sendStateID, LastErrorCode: "receiver_batch_ack_failed", LastErrorMessage: "receiver batch durable ack failed"})
			}
			return fmt.Errorf("%w: %w", msg.ErrSenderSyncFailed, err)
		}
	}
	return nil
}

func isSenderCommitted(canonical *repository.CanonicalMessageResult) bool {
	if canonical == nil {
		return false
	}
	if canonical.SendStateStatus == repository.SendStateStatusSenderCommitted ||
		canonical.SendStateStatus == repository.SendStateStatusReceiverAcked ||
		canonical.SendStateStatus == repository.SendStateStatusCompleted {
		return true
	}
	return canonical.SendStateStatus == repository.SendStateStatusFailedRetryable &&
		canonical.SenderOperationID != "" &&
		len(canonical.SenderUpdatePayload) > 0
}

func isSendStateSenderCommitted(sendState *repository.SendState) bool {
	if sendState == nil {
		return false
	}
	if sendState.Status == repository.SendStateStatusSenderCommitted ||
		sendState.Status == repository.SendStateStatusReceiverAcked ||
		sendState.Status == repository.SendStateStatusCompleted {
		return true
	}
	return sendState.Status == repository.SendStateStatusFailedRetryable &&
		sendState.SenderOperationID != "" &&
		len(sendState.SenderUpdatePayload) > 0
}

func needsReceiverAck(status int32) bool {
	return status != repository.SendStateStatusReceiverAcked && status != repository.SendStateStatusCompleted
}

func isCompleted(status int32) bool {
	return status == repository.SendStateStatusCompleted
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

func senderResultFromCanonical(canonical *repository.CanonicalMessageResult, operationID string, senderUserID int64) *userupdates.UserOperationResult {
	if canonical == nil {
		return nil
	}
	if canonical.SenderOperationID != "" {
		operationID = canonical.SenderOperationID
	}
	return userupdates.MakeTLUserOperationResult(&userupdates.TLUserOperationResult{
		UserId:              senderUserID,
		OperationId:         operationID,
		Status:              1,
		Pts:                 canonical.SenderPTS,
		PtsCount:            canonical.SenderPTSCount,
		CurrentPts:          canonical.SenderPTS,
		ResponsePayload:     canonical.SenderUpdatePayload,
		ResponsePayloadHash: canonical.SenderUpdatePayloadHash,
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
	envelope, hashBytes, err := c.buildSenderOperationEnvelope(in, canonical, operationID, normalized, sideEffects)
	if err != nil {
		return nil, nil, err
	}
	result, err := c.dispatchRequesterSync(envelope, nil)
	if err != nil {
		return nil, nil, err
	}
	return result, hashBytes, nil
}

func (c *MsgCore) buildSenderOperationEnvelope(in *msg.TLMsgSendMessageV2, canonical *repository.CanonicalMessageResult, operationID string, normalized normalizedOutboxMessage, sideEffects batchSideEffects) (OperationEnvelope, []byte, error) {
	body, _, hashBytes, err := buildNormalizedMessageOperationPayload(normalized, in.PeerId, in.PeerId, true, canonical, sideEffects)
	if err != nil {
		return OperationEnvelope{}, nil, err
	}
	authKeyID := in.AuthKeyId
	return OperationEnvelope{
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
	}, hashBytes, nil
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

func fullSentMessageUpdates(senderUserID int64, peerID int64, canonicals []repository.CanonicalMessageResult, results []*userupdates.UserOperationResult, normalized []normalizedOutboxMessage) (*tg.Updates, error) {
	if len(canonicals) != len(results) || len(canonicals) != len(normalized) {
		return nil, msg.ErrSenderSyncFailed
	}
	updates := make([]tg.UpdateClazz, 0, len(canonicals))
	var updatesDate int32
	for i := range canonicals {
		response, err := operationResponseV2(results[i], "sender batch")
		if err != nil {
			return nil, err
		}
		userMessageID, err := operationResponseUserMessageID(response, "sender batch")
		if err != nil {
			return nil, err
		}
		pts, err := int64ToInt32(response.Pts, "pts")
		if err != nil {
			return nil, err
		}
		date, err := msgDateInt32FromUnixSeconds(canonicals[i].MessageDate, "batch message date")
		if err != nil {
			return nil, err
		}
		updatesDate = date
		updates = append(updates, tg.MakeTLUpdateNewMessage(&tg.TLUpdateNewMessage{
			Message: tg.MakeTLMessage(&tg.TLMessage{
				Out:         true,
				Silent:      normalized[i].Attrs.Silent,
				Noforwards:  normalized[i].Attrs.Noforwards,
				InvertMedia: normalized[i].Attrs.InvertMedia,
				Id:          userMessageID,
				FromId:      tg.MakePeerUser(senderUserID),
				PeerId:      tg.MakePeerUser(peerID),
				FwdFrom:     sentMessageForwardHeader(normalized[i].ForwardRef),
				ReplyTo:     sentMessageReplyHeader(normalized[i].ReplyToUserMessageID),
				Date:        date,
				Message:     normalized[i].MessageText,
				Media:       sentMessageMedia(normalized[i].MediaRef),
				Entities:    sentMessageEntities(normalized[i].Entities),
				GroupedId:   sentMessageGroupedID(normalized[i].attrsPtr()),
				TtlPeriod:   sentMessageTTLPeriod(normalized[i].MediaRef),
			}),
			Pts:      pts,
			PtsCount: response.PtsCount,
		}))
	}
	return tg.MakeTLUpdates(&tg.TLUpdates{
		Updates: updates,
		Users:   []tg.UserClazz{},
		Chats:   []tg.ChatClazz{},
		Date:    updatesDate,
		Seq:     0,
	}).ToUpdates(), nil
}

func requiresFullSentUpdates(normalized normalizedOutboxMessage) bool {
	if normalized.MediaRef != nil || normalized.ForwardRef != nil || len(normalized.Entities) > 0 {
		return true
	}
	return normalized.hasAttrs()
}

func sentMessageMedia(media *payload.MediaRefV1) tg.MessageMediaClazz {
	if media == nil {
		return nil
	}
	ttl := sentMessageTTLPeriod(media)
	switch media.Kind {
	case "photo":
		return tg.MakeTLMessageMediaPhoto(&tg.TLMessageMediaPhoto{
			Photo:      tg.MakeTLPhotoEmpty(&tg.TLPhotoEmpty{Id: media.ID}),
			TtlSeconds: ttl,
		})
	case "document":
		return tg.MakeTLMessageMediaDocument(&tg.TLMessageMediaDocument{
			Document:   tg.MakeTLDocumentEmpty(&tg.TLDocumentEmpty{Id: media.ID}),
			TtlSeconds: ttl,
		})
	default:
		return tg.MakeTLMessageMediaEmpty(&tg.TLMessageMediaEmpty{})
	}
}

func sentMessageTTLPeriod(media *payload.MediaRefV1) *int32 {
	if media == nil || media.TTLSeconds == 0 {
		return nil
	}
	ttl := media.TTLSeconds
	return &ttl
}

func sentMessageGroupedID(attrs *payload.MessageAttrsV1) *int64 {
	if attrs == nil || attrs.GroupedID == 0 {
		return nil
	}
	groupedID := attrs.GroupedID
	return &groupedID
}

func sentMessageForwardHeader(ref *payload.ForwardRefV1) tg.MessageFwdHeaderClazz {
	if ref == nil {
		return nil
	}
	date, err := msgDateInt32FromUnixSeconds(ref.Date, "forward date")
	if err != nil {
		return nil
	}
	var sourceMessageID *int32
	if ref.SourceMessageID > 0 {
		v, err := int64ToInt32(ref.SourceMessageID, "forward source message id")
		if err != nil {
			return nil
		}
		sourceMessageID = &v
	}
	var savedFromMessageID *int32
	if ref.SavedFromMessageID > 0 {
		v, err := int64ToInt32(ref.SavedFromMessageID, "forward saved message id")
		if err != nil {
			return nil
		}
		savedFromMessageID = &v
	}
	return tg.MakeTLMessageFwdHeader(&tg.TLMessageFwdHeader{
		FromId:         sentMessageForwardPeer(ref),
		FromName:       stringPtr(ref.FromName),
		Date:           date,
		ChannelPost:    sourceMessageID,
		SavedFromPeer:  sentMessagePeerFromOptional(ref.SavedFromPeerType, ref.SavedFromPeerID),
		SavedFromMsgId: savedFromMessageID,
	})
}

func sentMessageForwardPeer(ref *payload.ForwardRefV1) tg.PeerClazz {
	if ref.FromUserID > 0 {
		return tg.MakePeerUser(ref.FromUserID)
	}
	return sentMessagePeerFromOptional(ref.SourcePeerType, ref.SourcePeerID)
}

func sentMessagePeerFromOptional(peerType int32, peerID int64) tg.PeerClazz {
	if peerID == 0 {
		return nil
	}
	switch peerType {
	case payload.PeerTypeUser:
		return tg.MakePeerUser(peerID)
	case payload.PeerTypeChat:
		return tg.MakePeerChat(peerID)
	case payload.PeerTypeChannel:
		return tg.MakePeerChannel(peerID)
	default:
		return nil
	}
}

func sentMessageReplyHeader(userMessageID int64) tg.MessageReplyHeaderClazz {
	if userMessageID <= 0 {
		return nil
	}
	replyToMsgID, err := int64ToInt32(userMessageID, "reply user message id")
	if err != nil {
		return nil
	}
	return tg.MakeTLMessageReplyHeader(&tg.TLMessageReplyHeader{ReplyToMsgId: &replyToMsgID})
}

func sentMessageEntities(entities []payload.MessageEntityV1) []tg.MessageEntityClazz {
	if len(entities) == 0 {
		return nil
	}
	out := make([]tg.MessageEntityClazz, 0, len(entities))
	for _, entity := range entities {
		switch entity.Kind {
		case "mention":
			out = append(out, tg.MakeTLMessageEntityMention(&tg.TLMessageEntityMention{Offset: entity.Offset, Length: entity.Length}))
		case "hashtag":
			out = append(out, tg.MakeTLMessageEntityHashtag(&tg.TLMessageEntityHashtag{Offset: entity.Offset, Length: entity.Length}))
		case "bot_command":
			out = append(out, tg.MakeTLMessageEntityBotCommand(&tg.TLMessageEntityBotCommand{Offset: entity.Offset, Length: entity.Length}))
		case "url":
			out = append(out, tg.MakeTLMessageEntityUrl(&tg.TLMessageEntityUrl{Offset: entity.Offset, Length: entity.Length}))
		case "email":
			out = append(out, tg.MakeTLMessageEntityEmail(&tg.TLMessageEntityEmail{Offset: entity.Offset, Length: entity.Length}))
		case "bold":
			out = append(out, tg.MakeTLMessageEntityBold(&tg.TLMessageEntityBold{Offset: entity.Offset, Length: entity.Length}))
		case "italic":
			out = append(out, tg.MakeTLMessageEntityItalic(&tg.TLMessageEntityItalic{Offset: entity.Offset, Length: entity.Length}))
		case "code":
			out = append(out, tg.MakeTLMessageEntityCode(&tg.TLMessageEntityCode{Offset: entity.Offset, Length: entity.Length}))
		case "pre":
			out = append(out, tg.MakeTLMessageEntityPre(&tg.TLMessageEntityPre{Offset: entity.Offset, Length: entity.Length, Language: entity.URL}))
		case "text_url":
			out = append(out, tg.MakeTLMessageEntityTextUrl(&tg.TLMessageEntityTextUrl{Offset: entity.Offset, Length: entity.Length, Url: entity.URL}))
		case "phone":
			out = append(out, tg.MakeTLMessageEntityPhone(&tg.TLMessageEntityPhone{Offset: entity.Offset, Length: entity.Length}))
		case "cashtag":
			out = append(out, tg.MakeTLMessageEntityCashtag(&tg.TLMessageEntityCashtag{Offset: entity.Offset, Length: entity.Length}))
		case "underline":
			out = append(out, tg.MakeTLMessageEntityUnderline(&tg.TLMessageEntityUnderline{Offset: entity.Offset, Length: entity.Length}))
		case "strike":
			out = append(out, tg.MakeTLMessageEntityStrike(&tg.TLMessageEntityStrike{Offset: entity.Offset, Length: entity.Length}))
		case "bank_card":
			out = append(out, tg.MakeTLMessageEntityBankCard(&tg.TLMessageEntityBankCard{Offset: entity.Offset, Length: entity.Length}))
		case "spoiler":
			out = append(out, tg.MakeTLMessageEntitySpoiler(&tg.TLMessageEntitySpoiler{Offset: entity.Offset, Length: entity.Length}))
		case "blockquote":
			out = append(out, tg.MakeTLMessageEntityBlockquote(&tg.TLMessageEntityBlockquote{Offset: entity.Offset, Length: entity.Length}))
		}
	}
	return out
}

func stringPtr(v string) *string {
	if v == "" {
		return nil
	}
	return &v
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
