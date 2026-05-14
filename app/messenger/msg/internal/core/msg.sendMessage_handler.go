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
	"github.com/teamgram/teamgram-server/v2/pkg/proto/bin"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/iface"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

const maxSendMessageBatchSize = 100

// MsgSendMessage
// msg.sendMessage user_id:long auth_key_id:long peer_type:int peer_id:long message:Vector<OutboxMessage> = Updates;
func (c *MsgCore) MsgSendMessage(in *msg.TLMsgSendMessage) (*tg.Updates, error) {
	if in == nil || len(in.Message) == 0 {
		return nil, fmt.Errorf("%w: first slice requires at least one message", msg.ErrSendStateConflict)
	}
	if len(in.Message) > maxSendMessageBatchSize {
		return nil, fmt.Errorf("%w: max=%d got=%d", msg.ErrBatchTooLarge, maxSendMessageBatchSize, len(in.Message))
	}
	if in.PeerType != payload.PeerTypeUser && in.PeerType != payload.PeerTypeChat {
		return nil, fmt.Errorf("%w: unsupported peer type=%d", msg.ErrSendStateConflict, in.PeerType)
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
	if err := c.revalidateForwardSources(normalizedBatch); err != nil {
		return nil, err
	}
	sideEffects := batchSideEffectsFromRequest(in)
	attachFacts, err := attachFactsFromRequest(in)
	if err != nil {
		return nil, err
	}
	if len(normalizedBatch) > 1 {
		if in.PeerType == payload.PeerTypeChat {
			return c.sendMessageChatBatch(in, normalizedBatch, sideEffects, attachFacts)
		}
		return c.sendMessageBatch(in, normalizedBatch, sideEffects, attachFacts)
	}

	normalized := normalizedBatch[0]
	if in.PeerType == payload.PeerTypeChat {
		action, mediaKind := chatSendActionForNormalized(normalized)
		if err := c.checkChatAction(in.UserId, in.PeerId, action, mediaKind); err != nil {
			return nil, err
		}
	}
	_, requestHash, err := marshalSendRequestV3(normalized, sideEffects, attachFacts...)
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
		ServiceActionSchemaVersion:   canonicalPayloads.ServiceActionSchemaVersion,
		ServiceActionPayload:         canonicalPayloads.ServiceActionPayload,
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
		effects, err := c.buildSendMessageChatReceiverEffects(in, canonical, normalized, attachFacts)
		if err != nil {
			return nil, err
		}
		senderResult, senderHash, err = c.processSenderOperation(in, canonical, senderOperationID, normalized, sideEffects, effects, attachFacts)
		if err != nil {
			return nil, err
		}
		if err := c.markSenderCommitted(canonical, senderOperationID, senderResult); err != nil {
			c.Logger.Errorf("msg.sendMessage - sender commit recovery: send_state_id=%d canonical_message_id=%d operation_id=%s err=%v", canonical.SendStateID, canonical.CanonicalMessageID, senderOperationID, err)
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

	if in.PeerType == payload.PeerTypeUser && in.UserId != in.PeerId && needsReceiverAck(sendState.Status) {
		receiverOp, err := buildReceiverOperationEnvelope(in, canonical, normalized, attachFacts)
		if err != nil {
			return nil, err
		}
		ack, err := c.dispatchBrokerDurableAck(receiverOp)
		if err != nil {
			_ = c.svcCtx.Repo.MarkRetryableFailure(c.ctx, repository.MarkRetryableFailureInput{SendStateID: canonical.SendStateID, LastErrorCode: "receiver_ack_failed", LastErrorMessage: "receiver durable ack failed"})
			return nil, fmt.Errorf("%w: %w", msg.ErrSenderSyncFailed, err)
		}
		c.Logger.Debugf(
			"msg.sendMessage - receiver operation published: operation_id=%s topic=%s partition=%d offset=%d",
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

	return sentMessageUpdatesFromUserupdatesEnvelope(senderResult, "sender")
}

func (c *MsgCore) revalidateForwardSources(normalizedBatch []normalizedOutboxMessage) error {
	sources := make([]repository.ForwardSourceIdentity, 0, len(normalizedBatch))
	for _, normalized := range normalizedBatch {
		if normalized.ForwardRef == nil {
			continue
		}
		if normalized.ForwardSourceCanonicalID <= 0 || normalized.ForwardSourceUserMessageID <= 0 {
			return msg.ErrMsgIdInvalid
		}
		sources = append(sources, repository.ForwardSourceIdentity{
			UserID:             normalized.FromUserID,
			UserMessageID:      normalized.ForwardSourceUserMessageID,
			CanonicalMessageID: normalized.ForwardSourceCanonicalID,
		})
	}
	if len(sources) == 0 {
		return nil
	}
	return c.svcCtx.Repo.RevalidateForwardSources(c.ctx, sources)
}

func (c *MsgCore) sendMessageBatch(in *msg.TLMsgSendMessage, normalizedBatch []normalizedOutboxMessage, sideEffects batchSideEffects, attachFacts []payload.UpdateFactV1) (*tg.Updates, error) {
	batchInput, err := buildCanonicalBatchInput(in, normalizedBatch, sideEffects, attachFacts)
	if err != nil {
		return nil, err
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
		envelope, _, err := c.buildSenderOperationEnvelope(in, canonical, operationID, normalizedBatch[i], sideEffects, attachFacts)
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
	if err := c.dispatchBatchReceiverOps(in, canonicalBatch.Items, normalizedBatch, attachFacts); err != nil {
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
	return combineUserupdatesReplyEnvelopes(results, "sender batch")
}

func buildCanonicalBatchInput(in *msg.TLMsgSendMessage, normalizedBatch []normalizedOutboxMessage, sideEffects batchSideEffects, attachFacts []payload.UpdateFactV1) (repository.CreateCanonicalBatchInput, error) {
	batchInput := repository.CreateCanonicalBatchInput{
		SenderUserID: in.UserId,
		PeerType:     in.PeerType,
		PeerID:       in.PeerId,
		Items:        make([]repository.CreateCanonicalBatchItem, 0, len(normalizedBatch)),
	}
	for _, normalized := range normalizedBatch {
		_, requestHash, err := marshalSendRequestV3(normalized, sideEffects, attachFacts...)
		if err != nil {
			return repository.CreateCanonicalBatchInput{}, err
		}
		canonicalPayloads, err := normalizedCanonicalPayloads(normalized)
		if err != nil {
			return repository.CreateCanonicalBatchInput{}, err
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
			ServiceActionSchemaVersion:   canonicalPayloads.ServiceActionSchemaVersion,
			ServiceActionPayload:         canonicalPayloads.ServiceActionPayload,
		})
	}
	return batchInput, nil
}

func (c *MsgCore) sendMessageChatBatch(in *msg.TLMsgSendMessage, normalizedBatch []normalizedOutboxMessage, sideEffects batchSideEffects, attachFacts []payload.UpdateFactV1) (*tg.Updates, error) {
	if err := c.checkChatBatchActions(in.UserId, in.PeerId, normalizedBatch); err != nil {
		return nil, err
	}
	batchInput, err := buildCanonicalBatchInput(in, normalizedBatch, sideEffects, attachFacts)
	if err != nil {
		return nil, err
	}
	canonicalBatch, err := c.svcCtx.Repo.CreateOrGetCanonicalBatchByClientRandom(c.ctx, batchInput)
	if err != nil {
		return nil, err
	}
	if canonicalBatch == nil || len(canonicalBatch.Items) != len(normalizedBatch) {
		return nil, msg.ErrMsgStorage
	}

	results := make([]*userupdates.UserOperationResult, len(normalizedBatch))
	for i := range normalizedBatch {
		canonical := &canonicalBatch.Items[i]
		operationID := payload.SenderOperationID(canonical.CanonicalMessageID, in.UserId)
		if isSenderCommitted(canonical) {
			results[i] = senderResultFromCanonical(canonical, operationID, in.UserId)
			continue
		}
		effects, err := c.buildSendMessageChatReceiverEffects(in, canonical, normalizedBatch[i], attachFacts)
		if err != nil {
			return nil, err
		}
		result, _, err := c.processSenderOperation(in, canonical, operationID, normalizedBatch[i], sideEffects, effects, attachFacts)
		if err != nil {
			return nil, err
		}
		if err := c.markSenderCommitted(canonical, operationID, result); err != nil {
			_ = c.svcCtx.Repo.MarkRetryableFailure(c.ctx, repository.MarkRetryableFailureInput{SendStateID: canonical.SendStateID, LastErrorCode: "sender_chat_batch_commit_failed", LastErrorMessage: "sender chat batch commit failed"})
			return nil, fmt.Errorf("%w: %v", msg.ErrSenderSyncFailed, err)
		}
		results[i] = result
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
	return combineUserupdatesReplyEnvelopes(results, "sender chat batch")
}

func (c *MsgCore) dispatchBatchReceiverOps(in *msg.TLMsgSendMessage, canonicals []repository.CanonicalMessageResult, normalizedBatch []normalizedOutboxMessage, attachFacts []payload.UpdateFactV1) error {
	if in.UserId == in.PeerId {
		return nil
	}
	sendStateIDs := make([]int64, 0, len(canonicals))
	receiverCanonicals := make([]repository.CanonicalMessageResult, 0, len(canonicals))
	receiverMessages := make([]normalizedOutboxMessage, 0, len(canonicals))
	for i := range canonicals {
		canonical := &canonicals[i]
		if !needsReceiverAck(canonical.SendStateStatus) {
			continue
		}
		sendStateIDs = append(sendStateIDs, canonical.SendStateID)
		receiverCanonicals = append(receiverCanonicals, *canonical)
		receiverMessages = append(receiverMessages, normalizedBatch[i])
	}
	if len(receiverCanonicals) == 0 {
		return nil
	}
	envelope, err := buildReceiverBatchOperationEnvelope(in, receiverCanonicals, receiverMessages, attachFacts)
	if err != nil {
		return err
	}
	if _, err := c.dispatchBrokerDurableAck(envelope); err != nil {
		for _, sendStateID := range sendStateIDs {
			_ = c.svcCtx.Repo.MarkRetryableFailure(c.ctx, repository.MarkRetryableFailureInput{SendStateID: sendStateID, LastErrorCode: "receiver_batch_ack_failed", LastErrorMessage: "receiver batch durable ack failed"})
		}
		return fmt.Errorf("%w: %w", msg.ErrSenderSyncFailed, err)
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

func batchSideEffectsFromRequest(in *msg.TLMsgSendMessage) batchSideEffects {
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

func attachFactsFromRequest(in *msg.TLMsgSendMessage) ([]payload.UpdateFactV1, error) {
	if in == nil || len(in.AttachFacts) == 0 {
		return nil, nil
	}
	out := make([]payload.UpdateFactV1, 0, len(in.AttachFacts))
	for i, fact := range in.AttachFacts {
		if fact == nil {
			return nil, fmt.Errorf("%w: attach_facts[%d] is nil", msg.ErrSendStateConflict, i)
		}
		var envelope payload.UpdateFactV1
		if err := json.Unmarshal(fact.Payload, &envelope); err != nil {
			return nil, fmt.Errorf("%w: decode attach_facts[%d]: %v", msg.ErrSendStateConflict, i, err)
		}
		if fact.Kind != envelope.Kind {
			return nil, fmt.Errorf("%w: attach_facts[%d] kind mismatch tl=%q json=%q", msg.ErrSendStateConflict, i, fact.Kind, envelope.Kind)
		}
		if _, err := payload.DecodeUpdateFact(envelope); err != nil {
			return nil, fmt.Errorf("%w: validate attach_facts[%d]: %v", msg.ErrSendStateConflict, i, err)
		}
		out = append(out, envelope)
	}
	return out, nil
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

func (c *MsgCore) processSenderOperation(in *msg.TLMsgSendMessage, canonical *repository.CanonicalMessageResult, operationID string, normalized normalizedOutboxMessage, sideEffects batchSideEffects, effects []OperationEnvelope, attachFacts []payload.UpdateFactV1) (*userupdates.UserOperationResult, []byte, error) {
	envelope, hashBytes, err := c.buildSenderOperationEnvelope(in, canonical, operationID, normalized, sideEffects, attachFacts)
	if err != nil {
		return nil, nil, err
	}
	result, err := c.dispatchRequesterSync(envelope, effects)
	if err != nil {
		return nil, nil, err
	}
	return result, hashBytes, nil
}

func (c *MsgCore) buildSenderOperationEnvelope(in *msg.TLMsgSendMessage, canonical *repository.CanonicalMessageResult, operationID string, normalized normalizedOutboxMessage, sideEffects batchSideEffects, attachFacts []payload.UpdateFactV1) (OperationEnvelope, []byte, error) {
	body, _, hashBytes, err := buildMessageOperationV4Payload(normalized, in.PeerId, in.PeerId, canonical, sideEffects, attachFacts)
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
		PayloadSchemaVersion: payload.MessageOperationSchemaVersionV4,
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
		SenderUpdateSchemaVersion: payload.OperationResponseSchemaVersionV3,
		SenderUpdatePayload:       result.ResponsePayload,
		SenderUpdatePayloadHash:   result.ResponsePayloadHash,
	})
}

func buildReceiverOperationEnvelope(in *msg.TLMsgSendMessage, canonical *repository.CanonicalMessageResult, normalized normalizedOutboxMessage, attachFacts []payload.UpdateFactV1) (OperationEnvelope, error) {
	operationID := payload.ReceiverOperationID(canonical.CanonicalMessageID, in.PeerId)
	body, hashBytes, _, err := buildMessageOperationV4Payload(normalized, in.PeerId, in.UserId, canonical, batchSideEffects{}, attachFacts)
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
		PayloadSchemaVersion: payload.MessageOperationSchemaVersionV4,
		PayloadCodec:         payload.PayloadCodecJSON,
		Payload:              body,
		PayloadHash:          hashBytes,
		DeliveryPolicy:       DeliveryPolicyBrokerDurableAck,
	}, nil
}

func buildReceiverBatchOperationEnvelope(in *msg.TLMsgSendMessage, canonicals []repository.CanonicalMessageResult, normalizedBatch []normalizedOutboxMessage, attachFacts []payload.UpdateFactV1) (OperationEnvelope, error) {
	if len(canonicals) == 0 || len(canonicals) != len(normalizedBatch) {
		return OperationEnvelope{}, msg.ErrMsgStorage
	}
	messages := make([]payload.NewMessageFactV1, 0, len(canonicals))
	canonicalMessageIDs := make([]int64, 0, len(canonicals))
	var firstCanonical *repository.CanonicalMessageResult
	for i := range canonicals {
		canonical := &canonicals[i]
		if firstCanonical == nil {
			firstCanonical = canonical
		}
		messageFact, err := newMessageFactForOperation(normalizedBatch[i], in.PeerId, in.UserId, canonical, batchSideEffects{})
		if err != nil {
			return OperationEnvelope{}, err
		}
		messages = append(messages, messageFact)
		canonicalMessageIDs = append(canonicalMessageIDs, canonical.CanonicalMessageID)
	}
	body, hashBytes, err := buildMessageOperationBatchV1Payload(messages, attachFacts)
	if err != nil {
		return OperationEnvelope{}, err
	}
	operationID := payload.ReceiverBatchOperationID(in.PeerId, canonicalMessageIDs)
	return OperationEnvelope{
		UserID:               in.PeerId,
		OperationID:          operationID,
		OpType:               payload.OpTypeSendMessage,
		OperationKind:        payload.OperationKindSendMessageBatch,
		ActorUserID:          in.UserId,
		PeerType:             in.PeerType,
		PeerID:               in.UserId,
		CanonicalMessageID:   &firstCanonical.CanonicalMessageID,
		CanonicalPeerSeq:     &firstCanonical.PeerSeq,
		CanonicalDate:        int64Ptr(firstCanonical.MessageDate),
		PayloadSchemaVersion: payload.MessageOperationSchemaVersionBatchV1,
		PayloadCodec:         payload.PayloadCodecJSON,
		Payload:              body,
		PayloadHash:          hashBytes,
		DeliveryPolicy:       DeliveryPolicyBrokerDurableAck,
	}, nil
}

func (c *MsgCore) buildSendMessageChatReceiverEffects(in *msg.TLMsgSendMessage, canonical *repository.CanonicalMessageResult, normalized normalizedOutboxMessage, attachFacts []payload.UpdateFactV1) ([]OperationEnvelope, error) {
	if in.PeerType != payload.PeerTypeChat {
		return nil, nil
	}
	memberIDs, err := c.activeChatReceiverIDs(in.PeerId, in.UserId)
	if err != nil {
		return nil, err
	}
	effects := make([]OperationEnvelope, 0, len(memberIDs))
	for _, receiverUserID := range memberIDs {
		effect, err := buildChatReceiverEffectEnvelope(in, canonical, normalized, receiverUserID, attachFacts)
		if err != nil {
			return nil, err
		}
		effects = append(effects, effect)
	}
	return effects, nil
}

func buildChatReceiverEffectEnvelope(in *msg.TLMsgSendMessage, canonical *repository.CanonicalMessageResult, normalized normalizedOutboxMessage, receiverUserID int64, attachFacts []payload.UpdateFactV1) (OperationEnvelope, error) {
	operationID := payload.ReceiverOperationID(canonical.CanonicalMessageID, receiverUserID)
	body, hashBytes, _, err := buildMessageOperationV4Payload(normalized, receiverUserID, in.PeerId, canonical, batchSideEffects{}, attachFacts)
	if err != nil {
		return OperationEnvelope{}, err
	}
	return OperationEnvelope{
		UserID:               receiverUserID,
		OperationID:          operationID,
		OpType:               payload.OpTypeSendMessage,
		OperationKind:        payload.OperationKindSendMessage,
		ActorUserID:          in.UserId,
		PeerType:             payload.PeerTypeChat,
		PeerID:               in.PeerId,
		CanonicalMessageID:   &canonical.CanonicalMessageID,
		CanonicalPeerSeq:     &canonical.PeerSeq,
		CanonicalDate:        int64Ptr(canonical.MessageDate),
		PayloadSchemaVersion: payload.MessageOperationSchemaVersionV4,
		PayloadCodec:         payload.PayloadCodecJSON,
		Payload:              body,
		PayloadHash:          hashBytes,
		DeliveryPolicy:       DeliveryPolicyDurableAsync,
	}, nil
}

func sentMessageUpdatesFromUserupdatesEnvelope(result *userupdates.UserOperationResult, operation string) (*tg.Updates, error) {
	response, err := operationResponseV3(result, operation)
	if err != nil {
		return nil, err
	}
	if len(response.ReplyEnvelope) == 0 {
		return nil, fmt.Errorf("%w: %s operation missing reply envelope", msg.ErrSenderSyncFailed, operation)
	}
	if response.ReplyEnvelopeCodec != payload.ReplyEnvelopeCodecTLBinary {
		return nil, fmt.Errorf("%w: unsupported %s reply envelope codec=%d", msg.ErrSenderSyncFailed, operation, response.ReplyEnvelopeCodec)
	}
	if response.ReplyEnvelopeSchema != payload.ReplyEnvelopeSchemaV1 {
		return nil, fmt.Errorf("%w: unsupported %s reply envelope schema=%d", msg.ErrSenderSyncFailed, operation, response.ReplyEnvelopeSchema)
	}
	obj, err := iface.DecodeObject(bin.NewDecoder(response.ReplyEnvelope))
	if err != nil {
		return nil, fmt.Errorf("%w: decode %s reply envelope: %v", msg.ErrSenderSyncFailed, operation, err)
	}
	if updates, ok := obj.(*tg.Updates); ok && updates != nil {
		return updates, nil
	}
	if updatesObject, ok := obj.(interface{ ToUpdates() *tg.Updates }); ok {
		if updates := updatesObject.ToUpdates(); updates != nil {
			return updates, nil
		}
	}
	if obj == nil {
		return nil, fmt.Errorf("%w: %s reply envelope type=%T", msg.ErrSenderSyncFailed, operation, obj)
	}
	return nil, fmt.Errorf("%w: %s reply envelope type=%T", msg.ErrSenderSyncFailed, operation, obj)
}

func combineUserupdatesReplyEnvelopes(results []*userupdates.UserOperationResult, operation string) (*tg.Updates, error) {
	if len(results) == 0 {
		return nil, msg.ErrSenderSyncFailed
	}
	if len(results) == 1 {
		return sentMessageUpdatesFromUserupdatesEnvelope(results[0], operation)
	}
	combined := &tg.TLUpdates{
		Updates: make([]tg.UpdateClazz, 0, len(results)*2),
		Users:   []tg.UserClazz{},
		Chats:   []tg.ChatClazz{},
	}
	for _, result := range results {
		reply, err := sentMessageUpdatesFromUserupdatesEnvelope(result, operation)
		if err != nil {
			return nil, err
		}
		replyUpdates, ok := reply.ToUpdates()
		if !ok {
			return nil, fmt.Errorf("%w: %s reply envelope type=%s", msg.ErrSenderSyncFailed, operation, reply.ClazzName())
		}
		combined.Updates = append(combined.Updates, replyUpdates.Updates...)
		combined.Users = append(combined.Users, replyUpdates.Users...)
		combined.Chats = append(combined.Chats, replyUpdates.Chats...)
		combined.Date = replyUpdates.Date
		combined.Seq = replyUpdates.Seq
	}
	return tg.MakeTLUpdates(combined).ToUpdates(), nil
}

func sentMessageServiceAction(ref *payload.ServiceActionRefV1) (tg.MessageActionClazz, error) {
	if ref == nil {
		return nil, nil
	}
	switch ref.Kind {
	case payload.ServiceActionKindChatCreate:
		return tg.MakeTLMessageActionChatCreate(&tg.TLMessageActionChatCreate{
			Title: ref.Title,
			Users: append([]int64(nil), ref.Users...),
		}), nil
	default:
		return nil, fmt.Errorf("%w: unsupported service action kind=%s schema=%d", msg.ErrMsgStorage, ref.Kind, ref.SchemaVersion)
	}
}

func sentMessageMedia(media *payload.MediaRefV1) tg.MessageMediaClazz {
	if media == nil {
		return nil
	}
	ttl := sentMessageTTLPeriod(media)
	switch media.Kind {
	case "photo":
		return tg.MakeTLMessageMediaPhoto(&tg.TLMessageMediaPhoto{
			Photo:      sentMessagePhoto(media),
			TtlSeconds: ttl,
		})
	case "document":
		flags := documentMediaFlags(media)
		return tg.MakeTLMessageMediaDocument(&tg.TLMessageMediaDocument{
			Spoiler:        flags.Spoiler,
			Video:          flags.Video,
			Round:          flags.Round,
			Voice:          flags.Voice,
			Document:       sentMessageDocument(media),
			VideoCover:     sentPhotoRef(media.VideoCover),
			VideoTimestamp: cloneInt32Ptr(media.VideoTimestamp),
			TtlSeconds:     ttl,
		})
	case "contact":
		return sentMessageContact(media)
	default:
		return tg.MakeTLMessageMediaEmpty(&tg.TLMessageMediaEmpty{})
	}
}

func sentMessageContact(media *payload.MediaRefV1) tg.MessageMediaClazz {
	return tg.MakeTLMessageMediaContact(&tg.TLMessageMediaContact{
		PhoneNumber: media.PhoneNumber,
		FirstName:   media.FirstName,
		LastName:    media.LastName,
		Vcard:       media.Vcard,
		UserId:      media.UserID,
	})
}

func sentMessagePhoto(media *payload.MediaRefV1) tg.PhotoClazz {
	if media.Date == 0 && media.DcID == 0 && len(media.PhotoSizes) == 0 {
		return tg.MakeTLPhotoEmpty(&tg.TLPhotoEmpty{Id: media.ID})
	}
	return tg.MakeTLPhoto(&tg.TLPhoto{
		Id:            media.ID,
		AccessHash:    media.AccessHash,
		FileReference: append([]byte(nil), media.FileReference...),
		Date:          media.Date,
		Sizes:         sentPhotoSizes(media.PhotoSizes),
		DcId:          media.DcID,
	})
}

func sentMessageDocument(media *payload.MediaRefV1) tg.DocumentClazz {
	if media.Date == 0 && media.DcID == 0 && media.Size == 0 && len(media.DocumentAttributes) == 0 {
		return tg.MakeTLDocumentEmpty(&tg.TLDocumentEmpty{Id: media.ID})
	}
	return tg.MakeTLDocument(&tg.TLDocument{
		Id:            media.ID,
		AccessHash:    media.AccessHash,
		FileReference: append([]byte(nil), media.FileReference...),
		Date:          media.Date,
		MimeType:      media.MimeType,
		Size2:         media.Size,
		Thumbs:        sentPhotoSizes(media.DocumentThumbs),
		VideoThumbs:   sentVideoSizes(media.DocumentVideoThumbs),
		DcId:          media.DcID,
		Attributes:    sentDocumentAttributes(media.DocumentAttributes),
	})
}

func sentPhotoRef(photo *payload.PhotoRefV1) tg.PhotoClazz {
	if photo == nil {
		return nil
	}
	if photo.Date == 0 && photo.DcID == 0 && len(photo.Sizes) == 0 && len(photo.VideoSizes) == 0 {
		return tg.MakeTLPhotoEmpty(&tg.TLPhotoEmpty{Id: photo.ID})
	}
	return tg.MakeTLPhoto(&tg.TLPhoto{
		Id:            photo.ID,
		AccessHash:    photo.AccessHash,
		FileReference: append([]byte(nil), photo.FileReference...),
		Date:          photo.Date,
		Sizes:         sentPhotoSizes(photo.Sizes),
		VideoSizes:    sentVideoSizes(photo.VideoSizes),
		DcId:          photo.DcID,
	})
}

func sentPhotoSizes(sizes []payload.PhotoSizeRefV1) []tg.PhotoSizeClazz {
	if len(sizes) == 0 {
		return nil
	}
	out := make([]tg.PhotoSizeClazz, 0, len(sizes))
	for _, size := range sizes {
		switch size.Kind {
		case "empty":
			out = append(out, tg.MakeTLPhotoSizeEmpty(&tg.TLPhotoSizeEmpty{Type: size.Type}))
		case "size":
			out = append(out, tg.MakeTLPhotoSize(&tg.TLPhotoSize{Type: size.Type, W: size.W, H: size.H, Size2: size.Size}))
		case "cached":
			out = append(out, tg.MakeTLPhotoCachedSize(&tg.TLPhotoCachedSize{Type: size.Type, W: size.W, H: size.H, Bytes: append([]byte(nil), size.Bytes...)}))
		case "stripped":
			out = append(out, tg.MakeTLPhotoStrippedSize(&tg.TLPhotoStrippedSize{Type: size.Type, Bytes: append([]byte(nil), size.Bytes...)}))
		case "progressive":
			out = append(out, tg.MakeTLPhotoSizeProgressive(&tg.TLPhotoSizeProgressive{Type: size.Type, W: size.W, H: size.H, Sizes: append([]int32(nil), size.Sizes...)}))
		case "path":
			out = append(out, tg.MakeTLPhotoPathSize(&tg.TLPhotoPathSize{Type: size.Type, Bytes: append([]byte(nil), size.Bytes...)}))
		}
	}
	return out
}

func sentVideoSizes(sizes []payload.VideoSizeRefV1) []tg.VideoSizeClazz {
	if len(sizes) == 0 {
		return nil
	}
	out := make([]tg.VideoSizeClazz, 0, len(sizes))
	for _, size := range sizes {
		switch size.Kind {
		case "size":
			out = append(out, tg.MakeTLVideoSize(&tg.TLVideoSize{
				Type:         size.Type,
				W:            size.W,
				H:            size.H,
				Size2:        size.Size,
				VideoStartTs: cloneFloat64Ptr(size.VideoStartTs),
			}))
		case "emoji_markup":
			out = append(out, tg.MakeTLVideoSizeEmojiMarkup(&tg.TLVideoSizeEmojiMarkup{
				EmojiId:          size.EmojiID,
				BackgroundColors: append([]int32(nil), size.BackgroundColors...),
			}))
		case "sticker_markup":
			out = append(out, tg.MakeTLVideoSizeStickerMarkup(&tg.TLVideoSizeStickerMarkup{
				Stickerset:       stickerSetRef(size.StickerSet),
				StickerId:        size.StickerID,
				BackgroundColors: append([]int32(nil), size.BackgroundColors...),
			}))
		}
	}
	return out
}

func sentDocumentAttributes(attrs []payload.DocumentAttributeRefV1) []tg.DocumentAttributeClazz {
	if len(attrs) == 0 {
		return nil
	}
	out := make([]tg.DocumentAttributeClazz, 0, len(attrs))
	for _, attr := range attrs {
		switch attr.Kind {
		case "filename":
			out = append(out, tg.MakeTLDocumentAttributeFilename(&tg.TLDocumentAttributeFilename{FileName: attr.FileName}))
		case "image_size":
			out = append(out, tg.MakeTLDocumentAttributeImageSize(&tg.TLDocumentAttributeImageSize{W: attr.W, H: attr.H}))
		case "animated":
			out = append(out, tg.MakeTLDocumentAttributeAnimated(&tg.TLDocumentAttributeAnimated{}))
		case "video":
			out = append(out, tg.MakeTLDocumentAttributeVideo(&tg.TLDocumentAttributeVideo{
				RoundMessage:      attr.RoundMessage,
				SupportsStreaming: attr.SupportsStreaming,
				Nosound:           attr.NoSound,
				Duration:          attr.DurationFloat,
				W:                 attr.W,
				H:                 attr.H,
				PreloadPrefixSize: attr.PreloadPrefixSize,
				VideoStartTs:      attr.VideoStartTs,
				VideoCodec:        attr.VideoCodec,
			}))
		case "audio":
			out = append(out, tg.MakeTLDocumentAttributeAudio(&tg.TLDocumentAttributeAudio{
				Voice:     attr.Voice,
				Duration:  attr.Duration,
				Title:     attr.Title,
				Performer: attr.Performer,
				Waveform:  append([]byte(nil), attr.Waveform...),
			}))
		case "sticker":
			out = append(out, tg.MakeTLDocumentAttributeSticker(&tg.TLDocumentAttributeSticker{
				Mask:       attr.Mask,
				Alt:        attr.Alt,
				Stickerset: stickerSetRef(attr.StickerSet),
				MaskCoords: maskCoordsRef(attr.MaskCoords),
			}))
		case "custom_emoji":
			out = append(out, tg.MakeTLDocumentAttributeCustomEmoji(&tg.TLDocumentAttributeCustomEmoji{
				Free:       attr.Free,
				TextColor:  attr.TextColor,
				Alt:        attr.Alt,
				Stickerset: stickerSetRef(attr.StickerSet),
			}))
		case "has_stickers":
			out = append(out, tg.MakeTLDocumentAttributeHasStickers(&tg.TLDocumentAttributeHasStickers{}))
		}
	}
	return out
}

func documentMediaFlags(media *payload.MediaRefV1) payload.DocumentMediaFlagsV1 {
	if media == nil {
		return payload.DocumentMediaFlagsV1{}
	}
	if media.DocumentMediaFlags != nil {
		return *media.DocumentMediaFlags
	}
	var flags payload.DocumentMediaFlagsV1
	for _, attr := range media.DocumentAttributes {
		switch attr.Kind {
		case "audio":
			flags.Voice = flags.Voice || attr.Voice
		case "video":
			flags.Round = flags.Round || attr.RoundMessage
			if !isWebMStickerOrCustomEmoji(media) {
				flags.Video = true
			}
		}
	}
	return flags
}

func isWebMStickerOrCustomEmoji(media *payload.MediaRefV1) bool {
	if media == nil || media.MimeType != "video/webm" {
		return false
	}
	for _, attr := range media.DocumentAttributes {
		if attr.Kind == "sticker" || attr.Kind == "custom_emoji" {
			return true
		}
	}
	return false
}

func stickerSetRef(ref *payload.StickerSetRefV1) tg.InputStickerSetClazz {
	if ref == nil {
		return tg.MakeTLInputStickerSetEmpty(&tg.TLInputStickerSetEmpty{})
	}
	switch ref.Kind {
	case "", "empty":
		return tg.MakeTLInputStickerSetEmpty(&tg.TLInputStickerSetEmpty{})
	case "id":
		return tg.MakeTLInputStickerSetID(&tg.TLInputStickerSetID{Id: ref.ID, AccessHash: ref.AccessHash})
	case "short_name":
		return tg.MakeTLInputStickerSetShortName(&tg.TLInputStickerSetShortName{ShortName: ref.ShortName})
	case "animated_emoji":
		return tg.MakeTLInputStickerSetAnimatedEmoji(&tg.TLInputStickerSetAnimatedEmoji{})
	case "dice":
		return tg.MakeTLInputStickerSetDice(&tg.TLInputStickerSetDice{Emoticon: ref.Emoticon})
	case "animated_emoji_animations":
		return tg.MakeTLInputStickerSetAnimatedEmojiAnimations(&tg.TLInputStickerSetAnimatedEmojiAnimations{})
	case "premium_gifts":
		return tg.MakeTLInputStickerSetPremiumGifts(&tg.TLInputStickerSetPremiumGifts{})
	case "emoji_generic_animations":
		return tg.MakeTLInputStickerSetEmojiGenericAnimations(&tg.TLInputStickerSetEmojiGenericAnimations{})
	case "emoji_default_statuses":
		return tg.MakeTLInputStickerSetEmojiDefaultStatuses(&tg.TLInputStickerSetEmojiDefaultStatuses{})
	case "emoji_default_topic_icons":
		return tg.MakeTLInputStickerSetEmojiDefaultTopicIcons(&tg.TLInputStickerSetEmojiDefaultTopicIcons{})
	case "emoji_channel_default_statuses":
		return tg.MakeTLInputStickerSetEmojiChannelDefaultStatuses(&tg.TLInputStickerSetEmojiChannelDefaultStatuses{})
	case "ton_gifts":
		return tg.MakeTLInputStickerSetTonGifts(&tg.TLInputStickerSetTonGifts{})
	default:
		return nil
	}
}

func maskCoordsRef(ref *payload.MaskCoordsRefV1) tg.MaskCoordsClazz {
	if ref == nil {
		return nil
	}
	return tg.MakeTLMaskCoords(&tg.TLMaskCoords{N: ref.N, X: ref.X, Y: ref.Y, Zoom: ref.Zoom})
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
	if ref.SourcePeerType != payload.PeerTypeUser {
		if peer := sentMessagePeerFromOptional(ref.SourcePeerType, ref.SourcePeerID); peer != nil {
			return peer
		}
	}
	if ref.FromUserID > 0 {
		return tg.MakePeerUser(ref.FromUserID)
	}
	if peer := sentMessagePeerFromOptional(ref.SourcePeerType, ref.SourcePeerID); peer != nil {
		return peer
	}
	return nil
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
		case "mention_name":
			out = append(out, tg.MakeTLMessageEntityMentionName(&tg.TLMessageEntityMentionName{Offset: entity.Offset, Length: entity.Length, UserId: entity.UserID}))
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

func operationResponseV3(result *userupdates.UserOperationResult, operation string) (payload.OperationResponseV3, error) {
	if result == nil {
		return payload.OperationResponseV3{}, msg.ErrSenderSyncFailed
	}
	var response payload.OperationResponseV3
	if err := json.Unmarshal(result.ResponsePayload, &response); err != nil {
		return payload.OperationResponseV3{}, fmt.Errorf("%w: decode %s operation response: %v", msg.ErrSenderSyncFailed, operation, err)
	}
	if response.SchemaVersion != payload.OperationResponseSchemaVersionV3 {
		return payload.OperationResponseV3{}, fmt.Errorf("%w: unsupported %s operation response schema=%d", msg.ErrSenderSyncFailed, operation, response.SchemaVersion)
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
