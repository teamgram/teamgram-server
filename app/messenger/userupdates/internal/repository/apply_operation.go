package repository

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"math"
	"strings"
	"unicode"
	"unicode/utf8"

	"github.com/teamgram/marmota/pkg/stores/sqlx"
	"github.com/teamgram/teamgram-server/v2/app/messenger/userupdates/internal/repository/model"
	"github.com/teamgram/teamgram-server/v2/app/messenger/userupdates/payload"
	"github.com/teamgram/teamgram-server/v2/app/messenger/userupdates/userupdates"
)

func (r *Repository) ClaimPartitionOwner(ctx context.Context, partitionID int32) (int64, error) {
	_, err := r.requireDB()
	if err != nil {
		return 0, err
	}
	if _, _, err := r.models.UserupdatesPartitionFencesModel.InsertIgnore(ctx, &model.UserupdatesPartitionFences{
		PartitionId:     partitionID,
		OwnerEpoch:      0,
		OwnerInstanceId: "unassigned",
		LeaseId:         "",
	}); err != nil {
		return 0, storageError("insert partition fence", err)
	}

	fence, err := r.models.UserupdatesPartitionFencesModel.SelectByPartitionId(ctx, partitionID)
	if err != nil {
		return 0, storageError("select partition fence", err)
	}

	affected, err := r.models.UserupdatesPartitionFencesModel.CasAcquireOwner(ctx, r.OwnerInstance(), "", partitionID, fence.OwnerEpoch)
	if err != nil {
		return 0, storageError("claim partition owner", err)
	}
	if affected == 0 {
		return 0, userupdates.ErrOwnerFenceFailed
	}
	return fence.OwnerEpoch + 1, nil
}

func (r *Repository) ApplyUserOperation(ctx context.Context, in ApplyUserOperationInput) (*ApplyUserOperationResult, error) {
	db, err := r.requireDB()
	if err != nil {
		return nil, err
	}
	if !bytes.Equal(in.PayloadHash, payload.HashBytes(in.Payload)) {
		return nil, userupdates.ErrOperationPayloadConflict
	}
	var out *ApplyUserOperationResult
	err = db.Transact(ctx, func(tx *sqlx.Tx) error {
		result, err := r.applyUserOperationTx(ctx, tx, in)
		if err != nil {
			return err
		}
		out = result
		return nil
	})
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (r *Repository) applyUserOperationTx(ctx context.Context, tx *sqlx.Tx, in ApplyUserOperationInput) (*ApplyUserOperationResult, error) {
	txModels := r.models.WithTx(tx)
	fence, err := r.ensurePartitionOwnedTx(txModels, in.PartitionID)
	if err != nil {
		return nil, err
	}

	if _, _, err := txModels.UserPtsStateModel.InsertIgnore(&model.UserPtsState{
		UserId:       in.UserID,
		Pts:          0,
		PtsUpdatedAt: unixNow(),
		PartitionId:  in.PartitionID,
		OwnerEpoch:   fence.OwnerEpoch,
		RowVersion:   0,
	}); err != nil {
		return nil, storageError("init user pts state", err)
	}

	state, err := r.lockUserPTSState(txModels, in.UserID)
	if err != nil {
		return nil, err
	}
	if state.PartitionId != in.PartitionID {
		return nil, fmt.Errorf("%w: user %d partition %d != operation partition %d", userupdates.ErrNotOwner, in.UserID, state.PartitionId, in.PartitionID)
	}

	existing, found, err := selectOperationResult(txModels, in.UserID, in.OperationID)
	if err != nil {
		return nil, err
	}
	if found {
		if !bytes.Equal(existing.PayloadHash, in.PayloadHash) {
			return nil, userupdates.ErrOperationPayloadConflict
		}
		return &ApplyUserOperationResult{
			UserID:                in.UserID,
			OperationID:           in.OperationID,
			Pts:                   existing.Pts,
			PtsCount:              existing.PtsCount,
			ResponseSchemaVersion: existing.ResponseSchemaVersion,
			ResponsePayload:       existing.ResponsePayload,
			ResponseHash:          existing.ResponseHash,
			AlreadyApplied:        true,
		}, nil
	}

	op, err := decodeMessageOperation(in.Payload)
	if err != nil {
		return nil, err
	}
	if len(in.DependencyPts) != 0 || len(op.DependencyPts) != 0 {
		return nil, userupdates.ErrOperationTerminal
	}
	if op.ClearDraft && op.SourcePermAuthKeyID == 0 {
		return nil, fmt.Errorf("%w: clear draft side effect requires source permanent auth key", userupdates.ErrOperationTerminal)
	}
	if op.ReplyToCanonicalMessageID != 0 {
		row, err := txModels.UserMessageViewsModel.SelectByUserCanonical(in.UserID, op.ReplyToCanonicalMessageID)
		if err != nil {
			if errors.Is(err, model.ErrNotFound) {
				return nil, fmt.Errorf("%w: reply target canonical_message_id=%d not visible to user_id=%d", userupdates.ErrOperationTerminal, op.ReplyToCanonicalMessageID, in.UserID)
			}
			return nil, storageError("select reply target view", err)
		}
		op.ReplyToPeerSeq = row.PeerSeq
		op.ReplyToUserMessageID = row.UserMessageId
	}
	if err := resolvePublicIDsForOperation(txModels, in.UserID, &op.MessageOperationV1); err != nil {
		return nil, err
	}

	nextPTS := state.Pts + 1
	ptsCount := int32(1)
	eventPayload, eventPayloadHash, responsePayload, responsePayloadHash, err := buildEventAndResponse(in, op, nextPTS, ptsCount)
	if err != nil {
		return nil, err
	}
	switch op.OperationKind {
	case payload.OperationKindSendMessage:
		if err := insertUserMessageView(txModels, in, op, eventPayload); err != nil {
			return nil, err
		}
		if err := upsertUserDialog(txModels, in, op.MessageOperationV1, nextPTS, eventPayload); err != nil {
			return nil, err
		}
		if err := insertDialogSideEffects(ctx, txModels, r, in, op.MessageOperationV1); err != nil {
			return nil, err
		}
	case payload.OperationKindReadHistory:
		if err := applyReadHistory(txModels, in, op.MessageOperationV1, nextPTS); err != nil {
			return nil, err
		}
	case payload.OperationKindDeleteMessages:
		var updated payload.MessageOperationV1
		updated, err = applyDeleteMessages(txModels, in, op.MessageOperationV1, nextPTS)
		if err != nil {
			return nil, err
		}
		op.MessageOperationV1 = updated
		eventPayload, eventPayloadHash, responsePayload, responsePayloadHash, err = buildEventAndResponse(in, op, nextPTS, ptsCount)
		if err != nil {
			return nil, err
		}
	case payload.OperationKindDeleteHistory:
		if err := applyDeleteHistory(txModels, in, op.MessageOperationV1, nextPTS); err != nil {
			return nil, err
		}
	case payload.OperationKindEditMessage:
		if err := applyEditMessage(txModels, in, op.MessageOperationV1, eventPayload, nextPTS); err != nil {
			return nil, err
		}
	case payload.OperationKindUpdatePinnedMessage:
		if err := applyUpdatePinnedMessage(txModels, in, op.MessageOperationV1, nextPTS); err != nil {
			return nil, err
		}
	case payload.OperationKindMarkDialogUnread:
		if err := applyMarkDialogUnread(txModels, in, op.MessageOperationV1, nextPTS); err != nil {
			return nil, err
		}
	case payload.OperationKindScheduledMarker:
		if err := applyScheduledMarker(txModels, in, op.MessageOperationV1, nextPTS); err != nil {
			return nil, err
		}
	default:
		return nil, fmt.Errorf("%w: unsupported operation kind=%s", userupdates.ErrOperationTerminal, op.OperationKind)
	}
	if err := insertPTSEvent(txModels, in, op, nextPTS, ptsCount, eventPayload, eventPayloadHash); err != nil {
		return nil, err
	}
	if err := insertPushTask(ctx, txModels, r, in, op.MessageOperationV1, nextPTS, eventPayload); err != nil {
		return nil, err
	}
	if err := insertOperationResult(txModels, in, nextPTS, ptsCount, responsePayload, responsePayloadHash); err != nil {
		return nil, err
	}
	if err := r.insertAffectedOutboxesTx(ctx, txModels, in.AffectedOutboxes); err != nil {
		return nil, err
	}
	affected, err := txModels.UserPtsStateModel.UpdatePts(nextPTS, unixNow(), in.PartitionID, fence.OwnerEpoch, in.UserID)
	if err != nil {
		return nil, storageError("update user pts state", err)
	}
	if affected == 0 {
		return nil, userupdates.ErrPtsContinuityViolation
	}
	return &ApplyUserOperationResult{
		UserID:                in.UserID,
		OperationID:           in.OperationID,
		Pts:                   nextPTS,
		PtsCount:              ptsCount,
		ResponseSchemaVersion: payload.OperationResponseSchemaVersion,
		ResponsePayload:       responsePayload,
		ResponseHash:          responsePayloadHash,
	}, nil
}

type messageOperation struct {
	payload.MessageOperationV1
	EventSchemaVersion int
	MediaRef           *payload.MediaRefV1
	Attrs              *payload.MessageAttrsV1
	ForwardRef         *payload.ForwardRefV1
}

func decodeMessageOperation(body []byte) (messageOperation, error) {
	var envelope struct {
		SchemaVersion int    `json:"schema_version"`
		OperationKind string `json:"operation_kind"`
	}
	if err := json.Unmarshal(body, &envelope); err != nil {
		return messageOperation{}, fmt.Errorf("%w: decode message operation envelope: %v", userupdates.ErrOperationTerminal, err)
	}
	switch envelope.SchemaVersion {
	case payload.MessageOperationSchemaVersionV1, payload.MessageOperationSchemaVersion:
		var op payload.MessageOperationV1
		if err := json.Unmarshal(body, &op); err != nil {
			return messageOperation{}, fmt.Errorf("%w: decode message operation: %v", userupdates.ErrOperationTerminal, err)
		}
		return messageOperationFromV1(op), nil
	case payload.MessageOperationSchemaVersionV3:
		var op payload.MessageOperationV3
		if err := json.Unmarshal(body, &op); err != nil {
			return messageOperation{}, fmt.Errorf("%w: decode v3 message operation: %v", userupdates.ErrOperationTerminal, err)
		}
		return messageOperationFromV3(op), nil
	default:
		return messageOperation{}, fmt.Errorf("%w: unsupported operation schema=%d kind=%s", userupdates.ErrOperationTerminal, envelope.SchemaVersion, envelope.OperationKind)
	}
}

func messageOperationFromV1(op payload.MessageOperationV1) messageOperation {
	return messageOperation{
		MessageOperationV1: op,
		EventSchemaVersion: payload.MessageEventSchemaVersion,
	}
}

func messageOperationFromV3(op payload.MessageOperationV3) messageOperation {
	return messageOperation{
		MessageOperationV1: payload.MessageOperationV1{
			SchemaVersion:              op.SchemaVersion,
			OperationKind:              op.OperationKind,
			CanonicalMessageID:         op.CanonicalMessageID,
			PeerType:                   op.PeerType,
			PeerID:                     op.PeerID,
			PeerSeq:                    op.PeerSeq,
			FromUserID:                 op.FromUserID,
			ToUserID:                   op.ToUserID,
			Date:                       op.Date,
			EditDate:                   op.EditDate,
			EditVersion:                op.EditVersion,
			Out:                        op.Out,
			MessageText:                op.MessageText,
			Entities:                   op.Entities,
			UserMessageID:              op.UserMessageID,
			ReplyToCanonicalMessageID:  op.ReplyToCanonicalMessageID,
			ReplyToPeerSeq:             op.ReplyToPeerSeq,
			ReplyToUserMessageID:       op.ReplyToUserMessageID,
			DependencyPts:              op.DependencyPts,
			ClearDraft:                 op.ClearDraft,
			SourcePermAuthKeyID:        op.SourcePermAuthKeyID,
			ClearDraftBeforeDate:       op.ClearDraftBeforeDate,
			SavedDialogSideEffect:      op.SavedDialogSideEffect,
			ReadMaxUserMessageID:       op.ReadMaxUserMessageID,
			ReadInboxMaxPeerSeq:        op.ReadInboxMaxPeerSeq,
			ReadInboxMaxUserMessageID:  op.ReadInboxMaxUserMessageID,
			ReadOutboxMaxPeerSeq:       op.ReadOutboxMaxPeerSeq,
			ReadOutboxMaxUserMessageID: op.ReadOutboxMaxUserMessageID,
			DeletePeerSeqs:             op.DeletePeerSeqs,
			DeleteUserMessageIDs:       op.DeleteUserMessageIDs,
			DeleteMaxPeerSeq:           op.DeleteMaxPeerSeq,
			JustClear:                  op.JustClear,
			Revoke:                     op.Revoke,
			UnreadMark:                 op.UnreadMark,
			PinnedPeerSeq:              op.PinnedPeerSeq,
			PinnedUserMessageID:        op.PinnedUserMessageID,
			PinnedCanonicalMessageID:   op.PinnedCanonicalMessageID,
			HasScheduled:               op.HasScheduled,
		},
		EventSchemaVersion: payload.MessageEventSchemaVersionV3,
		MediaRef:           op.MediaRef,
		Attrs:              op.Attrs,
		ForwardRef:         op.ForwardRef,
	}
}

func resolvePublicIDsForOperation(txModels *model.TxModels, userID int64, op *payload.MessageOperationV1) error {
	if op == nil {
		return nil
	}
	if createsUserMessageView(op.OperationKind) {
		existing, found, err := existingUserMessageView(txModels, userID, op.CanonicalMessageID)
		if err != nil {
			return err
		}
		if found {
			if err := ensureExistingMessageViewMatchesOperation(existing, *op); err != nil {
				return err
			}
			op.UserMessageID = existing.UserMessageId
		}
		if op.UserMessageID == 0 {
			userMessageID, err := nextUserMessageID(txModels, userID)
			if err != nil {
				return err
			}
			op.UserMessageID = userMessageID
		}
	}
	if op.OperationKind == payload.OperationKindEditMessage && op.UserMessageID == 0 {
		existingID, found, err := existingUserMessageID(txModels, userID, op.CanonicalMessageID)
		if err != nil {
			return err
		}
		if !found || existingID <= 0 {
			return fmt.Errorf("%w: edit target canonical_message_id=%d not visible to user_id=%d", userupdates.ErrOperationTerminal, op.CanonicalMessageID, userID)
		}
		op.UserMessageID = existingID
	}
	if op.ReplyToUserMessageID == 0 && op.ReplyToPeerSeq > 0 {
		id, err := resolveUserMessageIDByPeerSeq(txModels, userID, op.PeerType, op.PeerID, op.ReplyToPeerSeq)
		if err != nil {
			return err
		}
		op.ReplyToUserMessageID = id
	}
	if op.ReadInboxMaxUserMessageID == 0 && op.ReadInboxMaxPeerSeq > 0 {
		id, err := resolveUserMessageIDByPeerSeq(txModels, userID, op.PeerType, op.PeerID, op.ReadInboxMaxPeerSeq)
		if err != nil {
			return err
		}
		op.ReadInboxMaxUserMessageID = id
	}
	if op.ReadOutboxMaxUserMessageID == 0 && op.ReadOutboxMaxPeerSeq > 0 {
		id, err := resolveUserMessageIDByPeerSeq(txModels, userID, op.PeerType, op.PeerID, op.ReadOutboxMaxPeerSeq)
		if err != nil {
			return err
		}
		op.ReadOutboxMaxUserMessageID = id
	}
	if op.ReadMaxUserMessageID == 0 {
		if op.Out {
			op.ReadMaxUserMessageID = op.ReadOutboxMaxUserMessageID
		} else {
			op.ReadMaxUserMessageID = op.ReadInboxMaxUserMessageID
		}
	}
	if op.PinnedUserMessageID == 0 {
		pinnedPeerSeq := op.PinnedPeerSeq
		if pinnedPeerSeq == 0 {
			pinnedPeerSeq = op.PeerSeq
		}
		id, err := resolveUserMessageIDByPeerSeq(txModels, userID, op.PeerType, op.PeerID, pinnedPeerSeq)
		if err != nil {
			return err
		}
		op.PinnedUserMessageID = id
	}
	return nil
}

func (r *Repository) lockPartitionFence(txModels *model.TxModels, partitionID int32) (*model.UserupdatesPartitionFences, error) {
	fence, err := txModels.UserupdatesPartitionFencesModel.SelectByPartitionId(partitionID)
	if err != nil {
		if errors.Is(err, model.ErrNotFound) {
			return nil, userupdates.ErrNotOwner
		}
		return nil, storageError("lock partition fence", err)
	}
	return fence, nil
}

func (r *Repository) ensurePartitionOwnedTx(txModels *model.TxModels, partitionID int32) (*model.UserupdatesPartitionFences, error) {
	fence, err := txModels.UserupdatesPartitionFencesModel.SelectByPartitionId(partitionID)
	if err != nil {
		if !errors.Is(err, model.ErrNotFound) {
			return nil, storageError("lock partition fence", err)
		}
		if _, _, err := txModels.UserupdatesPartitionFencesModel.InsertIgnore(&model.UserupdatesPartitionFences{
			PartitionId:     partitionID,
			OwnerEpoch:      0,
			OwnerInstanceId: "unassigned",
			LeaseId:         "",
		}); err != nil {
			return nil, storageError("insert partition fence", err)
		}
		fence, err = txModels.UserupdatesPartitionFencesModel.SelectByPartitionId(partitionID)
		if err != nil {
			return nil, storageError("lock partition fence", err)
		}
	}
	if fence.OwnerInstanceId == r.OwnerInstance() {
		return fence, nil
	}
	if fence.OwnerInstanceId != "" && fence.OwnerInstanceId != "unassigned" {
		return nil, userupdates.ErrNotOwner
	}
	affected, err := txModels.UserupdatesPartitionFencesModel.CasAcquireOwner(r.OwnerInstance(), "", partitionID, fence.OwnerEpoch)
	if err != nil {
		return nil, storageError("claim partition owner", err)
	}
	if affected == 0 {
		return nil, userupdates.ErrOwnerFenceFailed
	}
	fence.OwnerEpoch++
	fence.OwnerInstanceId = r.OwnerInstance()
	fence.LeaseId = ""
	return fence, nil
}

func (r *Repository) lockUserPTSState(txModels *model.TxModels, userID int64) (*model.UserPtsState, error) {
	state, err := txModels.UserPtsStateModel.SelectForUpdate(userID)
	if err != nil {
		return nil, storageError("lock user pts state", err)
	}
	return state, nil
}

func buildEventAndResponse(in ApplyUserOperationInput, op messageOperation, pts int64, ptsCount int32) ([]byte, []byte, []byte, []byte, error) {
	eventKind := payload.EventKindNewMessage
	if op.OperationKind != payload.OperationKindSendMessage {
		eventKind = op.OperationKind
	}
	eventPayload, err := marshalMessageEvent(in, op, eventKind)
	if err != nil {
		return nil, nil, nil, nil, err
	}
	response := payload.OperationResponseV2{
		SchemaVersion: payload.OperationResponseSchemaVersion,
		OperationID:   in.OperationID,
		Pts:           pts,
		PtsCount:      ptsCount,
		EventType:     eventKind,
		UserMessageID: op.UserMessageID,
	}
	responsePayload, err := json.Marshal(response)
	if err != nil {
		return nil, nil, nil, nil, storageError("marshal operation response", err)
	}
	return eventPayload, payload.HashBytes(eventPayload), responsePayload, payload.HashBytes(responsePayload), nil
}

func marshalMessageEvent(in ApplyUserOperationInput, op messageOperation, eventKind string) ([]byte, error) {
	if op.EventSchemaVersion == payload.MessageEventSchemaVersionV3 {
		event := payload.MessageEventV3{
			SchemaVersion:        payload.MessageEventSchemaVersionV3,
			EventKind:            eventKind,
			CanonicalMessageID:   op.CanonicalMessageID,
			PeerSeq:              op.PeerSeq,
			MessageID:            op.UserMessageID,
			PeerType:             op.PeerType,
			PeerID:               op.PeerID,
			FromUserID:           op.FromUserID,
			ToUserID:             op.ToUserID,
			Date:                 op.Date,
			EditDate:             op.EditDate,
			EditVersion:          op.EditVersion,
			Out:                  op.Out,
			MessageText:          op.MessageText,
			Entities:             op.Entities,
			ReplyToUserMessageID: op.ReplyToUserMessageID,
			ReadMaxUserMessageID: op.ReadMaxUserMessageID,
			DeleteUserMessageIDs: append([]int64(nil), op.DeleteUserMessageIDs...),
			PinnedUserMessageID:  op.PinnedUserMessageID,
			AuthKeyIdExclude:     in.AuthKeyIDExclude,
			MediaRef:             op.MediaRef,
			Attrs:                op.Attrs,
			ForwardRef:           op.ForwardRef,
		}
		body, err := json.Marshal(event)
		if err != nil {
			return nil, storageError("marshal v3 message event", err)
		}
		return body, nil
	}
	event := payload.MessageEventV2{
		SchemaVersion:        payload.MessageEventSchemaVersion,
		EventKind:            eventKind,
		CanonicalMessageID:   op.CanonicalMessageID,
		PeerSeq:              op.PeerSeq,
		MessageID:            op.UserMessageID,
		PeerType:             op.PeerType,
		PeerID:               op.PeerID,
		FromUserID:           op.FromUserID,
		ToUserID:             op.ToUserID,
		Date:                 op.Date,
		EditDate:             op.EditDate,
		EditVersion:          op.EditVersion,
		Out:                  op.Out,
		MessageText:          op.MessageText,
		Entities:             op.Entities,
		ReplyToUserMessageID: op.ReplyToUserMessageID,
		ReadMaxUserMessageID: op.ReadMaxUserMessageID,
		DeleteUserMessageIDs: append([]int64(nil), op.DeleteUserMessageIDs...),
		PinnedUserMessageID:  op.PinnedUserMessageID,
		AuthKeyIdExclude:     in.AuthKeyIDExclude,
	}
	body, err := json.Marshal(event)
	if err != nil {
		return nil, storageError("marshal message event", err)
	}
	return body, nil
}

func insertUserMessageView(txModels *model.TxModels, in ApplyUserOperationInput, op messageOperation, viewPayload []byte) error {
	existing, err := txModels.UserMessageViewsModel.SelectByUserCanonical(in.UserID, op.CanonicalMessageID)
	if err != nil && !errors.Is(err, model.ErrNotFound) {
		return storageError("select existing user message view before insert", err)
	}
	if existing != nil {
		if existing.UserMessageId != op.UserMessageID || !bytes.Equal(existing.ViewPayload, viewPayload) {
			return userupdates.ErrOperationPayloadConflict
		}
	}
	_, _, err = txModels.UserMessageViewsModel.InsertOrUpdate(&model.UserMessageViews{
		UserId:             in.UserID,
		PeerType:           op.PeerType,
		PeerId:             op.PeerID,
		PeerSeq:            op.PeerSeq,
		CanonicalMessageId: op.CanonicalMessageID,
		UserMessageId:      op.UserMessageID,
		FromUserId:         op.FromUserID,
		Outgoing:           op.Out,
		MessageKind:        messageKindForOperation(op),
		MessageStatus:      MessageStatusLive,
		EditVersion:        0,
		Date:               int64(op.Date),
		EditDate:           0,
		DeletedAt:          0,
		ViewSchemaVersion:  int32(op.EventSchemaVersion),
		ViewPayload:        viewPayload,
	})
	if err != nil {
		return storageError("insert user message view", err)
	}
	if err := insertHashTagsTx(txModels, in.UserID, op.MessageOperationV1); err != nil {
		return err
	}
	return nil
}

func messageKindForOperation(op messageOperation) int32 {
	if op.MediaRef != nil {
		return MessageKindMedia
	}
	return MessageKindText
}

func upsertUserDialog(txModels *model.TxModels, in ApplyUserOperationInput, op payload.MessageOperationV1, nextPTS int64, dialogPayload []byte) error {
	unread := int32(0)
	if !op.Out {
		unread = 1
	}
	now := unixNow()
	_, _, err := txModels.UserDialogsModel.InsertOrUpdateMessageEvent(&model.UserDialogs{
		UserId:                     in.UserID,
		PeerType:                   op.PeerType,
		PeerId:                     op.PeerID,
		TopPeerSeq:                 op.PeerSeq,
		TopUserMessageId:           op.UserMessageID,
		TopCanonicalMessageId:      op.CanonicalMessageID,
		TopMessageDate:             int64(op.Date),
		TopMessageStatus:           MessageStatusLive,
		ReadInboxMaxPeerSeq:        op.ReadInboxMaxPeerSeq,
		ReadInboxMaxUserMessageId:  op.ReadInboxMaxUserMessageID,
		ReadOutboxMaxPeerSeq:       op.ReadOutboxMaxPeerSeq,
		ReadOutboxMaxUserMessageId: op.ReadOutboxMaxUserMessageID,
		UnreadCount:                unread,
		UnreadMentionsCount:        0,
		UnreadReactionsCount:       0,
		UnreadMark:                 false,
		PinnedPeerSeq:              op.PinnedPeerSeq,
		PinnedUserMessageId:        op.PinnedUserMessageID,
		PinnedCanonicalMessageId:   0,
		HasScheduled:               false,
		AvailableMinPeerSeq:        0,
		AvailableMinUserMessageId:  0,
		Hidden:                     false,
		DeletedAt:                  0,
		LastPts:                    nextPTS,
		LastPtsAt:                  now,
		DialogSchemaVersion:        1,
		DialogPayload:              dialogPayload,
	})
	if err != nil {
		return storageError("upsert user dialog", err)
	}
	return nil
}

func applyReadHistory(txModels *model.TxModels, in ApplyUserOperationInput, op payload.MessageOperationV1, nextPTS int64) error {
	readInbox := op.ReadInboxMaxPeerSeq
	readOutbox := op.ReadOutboxMaxPeerSeq
	readInboxUserMessageID := op.ReadInboxMaxUserMessageID
	readOutboxUserMessageID := op.ReadOutboxMaxUserMessageID
	var row *model.UserDialogs
	if readInbox == 0 || readOutbox == 0 || op.Out {
		if current, err := txModels.UserDialogsModel.SelectByUserPeer(in.UserID, op.PeerType, op.PeerID); err == nil {
			row = current
			if readInbox == 0 {
				readInbox = row.ReadInboxMaxPeerSeq
				readInboxUserMessageID = row.ReadInboxMaxUserMessageId
			}
			if readOutbox == 0 {
				readOutbox = row.ReadOutboxMaxPeerSeq
				readOutboxUserMessageID = row.ReadOutboxMaxUserMessageId
			}
		} else if !errors.Is(err, model.ErrNotFound) {
			return storageError("select dialog before read history", err)
		}
	}
	if readInbox == 0 && readOutbox == 0 {
		if op.Out {
			readOutbox = op.PeerSeq
			readOutboxUserMessageID = op.ReadMaxUserMessageID
		} else {
			readInbox = op.PeerSeq
			readInboxUserMessageID = op.ReadMaxUserMessageID
		}
	}
	unreadCount := int32(0)
	unreadMentionsCount := int32(0)
	unreadReactionsCount := int32(0)
	unreadMark := false
	if op.Out && row != nil {
		unreadCount = row.UnreadCount
		unreadMentionsCount = row.UnreadMentionsCount
		unreadReactionsCount = row.UnreadReactionsCount
		unreadMark = row.UnreadMark
	}
	_, err := txModels.UserDialogsModel.UpdateReadState(unreadCount, unreadMentionsCount, unreadReactionsCount, unreadMark, readInbox, readInboxUserMessageID, readOutbox, readOutboxUserMessageID, nextPTS, unixNow(), in.UserID, op.PeerType, op.PeerID)
	if err != nil {
		return storageError("apply read history", err)
	}
	if op.Out && readOutbox > 0 {
		if err := insertOutboxReadDateTx(txModels, in.UserID, op.PeerType, op.PeerID, op.PeerID, readOutbox, op.Date); err != nil {
			return err
		}
	}
	return nil
}

func insertHashTagsTx(txModels *model.TxModels, userID int64, op payload.MessageOperationV1) error {
	tags := extractHashTags(op.MessageText)
	if len(tags) == 0 || op.PeerSeq <= 0 {
		return nil
	}
	if op.PeerSeq > math.MaxInt32 {
		return storageError("index hashtag", fmt.Errorf("peer seq out of int32 range: %d", op.PeerSeq))
	}
	for _, tag := range tags {
		if _, _, err := txModels.HashTagsModel.InsertOrUpdate(&model.HashTags{
			UserId:               userID,
			PeerType:             op.PeerType,
			PeerId:               op.PeerID,
			HashTag:              tag,
			HashTagMessageId:     int32(op.PeerSeq),
			HashTagUserMessageId: op.UserMessageID,
		}); err != nil {
			return storageError("index hashtag", err)
		}
	}
	return nil
}

func insertOutboxReadDateTx(txModels *model.TxModels, userID int64, peerType int32, peerID int64, readUserID int64, maxPeerSeq int64, date int32) error {
	if maxPeerSeq > math.MaxInt32 {
		return storageError("insert outbox read date", fmt.Errorf("peer seq out of int32 range: %d", maxPeerSeq))
	}
	if _, _, err := txModels.MessageReadOutboxModel.InsertOrUpdate(&model.MessageReadOutbox{
		UserId:            userID,
		PeerType:          peerType,
		PeerId:            peerID,
		ReadUserId:        readUserID,
		ReadOutboxMaxId:   int32(maxPeerSeq),
		ReadOutboxMaxDate: unixOrNow(int64(date)),
	}); err != nil {
		return storageError("insert outbox read date", err)
	}
	return nil
}

func extractHashTags(text string) []string {
	seen := map[string]struct{}{}
	var out []string
	for i := 0; i < len(text); {
		r, size := utf8.DecodeRuneInString(text[i:])
		if r != '#' {
			i += size
			continue
		}
		start := i + size
		end := start
		for end < len(text) {
			next, nextSize := utf8.DecodeRuneInString(text[end:])
			if !isHashTagRune(next) {
				break
			}
			end += nextSize
		}
		if end > start {
			tag := strings.ToLower(text[start:end])
			if _, ok := seen[tag]; !ok {
				seen[tag] = struct{}{}
				out = append(out, tag)
			}
		}
		i = end
	}
	return out
}

func isHashTagRune(r rune) bool {
	return r == '_' || unicode.IsLetter(r) || unicode.IsDigit(r)
}

func applyDeleteMessages(txModels *model.TxModels, in ApplyUserOperationInput, op payload.MessageOperationV1, nextPTS int64) (payload.MessageOperationV1, error) {
	deletedAt := unixNow()
	deletedUserMessageIDs := make([]int64, 0, len(op.DeletePeerSeqs))
	unreadDeletes := int32(0)
	var dialog *model.UserDialogs
	if row, err := txModels.UserDialogsModel.SelectByUserPeer(in.UserID, op.PeerType, op.PeerID); err == nil {
		dialog = row
	} else if !errors.Is(err, model.ErrNotFound) {
		return op, storageError("select dialog before delete messages", err)
	}
	for _, peerSeq := range op.DeletePeerSeqs {
		row, err := txModels.UserMessageViewsModel.SelectByUserPeerSeq(in.UserID, op.PeerType, op.PeerID, peerSeq)
		if err != nil {
			if errors.Is(err, model.ErrNotFound) {
				continue
			}
			return op, storageError("select message before delete", err)
		}
		if row.MessageStatus == MessageStatusDeleted {
			continue
		}
		deletedUserMessageIDs = append(deletedUserMessageIDs, row.UserMessageId)
		if dialog != nil && !row.Outgoing && row.PeerSeq > dialog.ReadInboxMaxPeerSeq {
			unreadDeletes++
		}
		row.MessageStatus = MessageStatusDeleted
		row.DeletedAt = deletedAt
		row.ViewPayload = []byte(`{"schema_version":1,"deleted":true}`)
		if _, _, err := txModels.UserMessageViewsModel.InsertOrUpdate(row); err != nil {
			return op, storageError("mark message deleted", err)
		}
	}
	op.DeleteUserMessageIDs = deletedUserMessageIDs
	if dialog != nil && unreadDeletes > 0 {
		unreadCount := dialog.UnreadCount - unreadDeletes
		if unreadCount < 0 {
			unreadCount = 0
		}
		if _, err := txModels.UserDialogsModel.UpdateUnreadAfterDelete(
			unreadCount,
			dialog.UnreadMentionsCount,
			dialog.UnreadReactionsCount,
			dialog.UnreadMark,
			nextPTS,
			unixNow(),
			in.UserID,
			op.PeerType,
			op.PeerID,
		); err != nil {
			return op, storageError("update unread count after delete messages", err)
		}
	}
	if err := recomputeDialogTop(txModels, in.UserID, op.PeerType, op.PeerID, nextPTS); err != nil {
		return op, err
	}
	return op, nil
}

func applyDeleteHistory(txModels *model.TxModels, in ApplyUserOperationInput, op payload.MessageOperationV1, nextPTS int64) error {
	maxPeerSeq := op.DeleteMaxPeerSeq
	if maxPeerSeq == 0 {
		maxPeerSeq = op.PeerSeq
	}
	if maxPeerSeq <= 0 {
		if row, err := txModels.UserDialogsModel.SelectByUserPeer(in.UserID, op.PeerType, op.PeerID); err == nil {
			maxPeerSeq = row.TopPeerSeq
		} else if !errors.Is(err, model.ErrNotFound) {
			return storageError("select dialog before delete history", err)
		}
	}
	updateAvailableMin := op.JustClear || maxPeerSeq > 0
	availableMinUserMessageID := int64(0)
	if updateAvailableMin {
		var err error
		availableMinUserMessageID, err = resolveUserMessageIDByPeerSeq(txModels, in.UserID, op.PeerType, op.PeerID, maxPeerSeq)
		if err != nil {
			return err
		}
	}
	if maxPeerSeq > 0 {
		if _, err := txModels.UserMessageViewsModel.MarkHistoryDeleted(MessageStatusDeleted, []byte(`{"schema_version":1,"deleted":true}`), in.UserID, op.PeerType, op.PeerID, maxPeerSeq); err != nil {
			return storageError("mark history deleted", err)
		}
	}
	if err := recomputeDialogTop(txModels, in.UserID, op.PeerType, op.PeerID, nextPTS); err != nil {
		return err
	}
	if updateAvailableMin {
		if _, err := txModels.UserDialogsModel.UpdateAvailableMinPeerSeq(maxPeerSeq, availableMinUserMessageID, nextPTS, unixNow(), in.UserID, op.PeerType, op.PeerID); err != nil {
			return storageError("update dialog available min peer seq", err)
		}
	}
	return nil
}

func applyEditMessage(txModels *model.TxModels, in ApplyUserOperationInput, op payload.MessageOperationV1, viewPayload []byte, nextPTS int64) error {
	row, err := txModels.UserMessageViewsModel.SelectByUserCanonical(in.UserID, op.CanonicalMessageID)
	if err != nil {
		if errors.Is(err, model.ErrNotFound) {
			return userupdates.ErrOperationTerminal
		}
		return storageError("select message before edit", err)
	}
	if row.MessageStatus != MessageStatusLive {
		return userupdates.ErrOperationTerminal
	}
	editVersion := op.EditVersion
	if editVersion == 0 {
		editVersion = row.EditVersion + 1
	}
	row.EditVersion = editVersion
	row.ViewSchemaVersion = int32(payload.MessageEventSchemaVersion)
	row.ViewPayload = viewPayload
	if _, _, err := txModels.UserMessageViewsModel.InsertOrUpdate(row); err != nil {
		return storageError("update message view after edit", err)
	}
	if op.EditDate > 0 {
		if _, err := txModels.UserMessageViewsModel.UpdateEditDateByUserCanonical(int64(op.EditDate), in.UserID, op.CanonicalMessageID); err != nil {
			return storageError("update message view edit date", err)
		}
	}
	if err := insertHashTagsTx(txModels, in.UserID, op); err != nil {
		return err
	}
	dialog, err := txModels.UserDialogsModel.SelectByUserPeer(in.UserID, op.PeerType, op.PeerID)
	if err != nil {
		if errors.Is(err, model.ErrNotFound) {
			return nil
		}
		return storageError("select dialog before edit", err)
	}
	if dialog.TopCanonicalMessageId != op.CanonicalMessageID {
		return nil
	}
	dialog.DialogPayload = viewPayload
	dialog.LastPts = nextPTS
	dialog.LastPtsAt = unixNow()
	if _, _, err := txModels.UserDialogsModel.InsertOrUpdateMessageEvent(dialog); err != nil {
		return storageError("update dialog payload after edit", err)
	}
	return nil
}

func recomputeDialogTop(txModels *model.TxModels, userID int64, peerType int32, peerID int64, nextPTS int64) error {
	top, err := txModels.UserMessageViewsModel.SelectTopLiveByUserPeer(userID, peerType, peerID, MessageStatusLive)
	if err != nil {
		if !errors.Is(err, sqlx.ErrNotFound) && !errors.Is(err, model.ErrNotFound) {
			return storageError("select top message after delete", err)
		}
		if _, execErr := txModels.UserDialogsModel.ClearDialogTopAfterDelete(MessageStatusDeleted, unixNow(), nextPTS, unixNow(), userID, peerType, peerID); execErr != nil {
			return storageError("clear dialog top after delete", execErr)
		}
		return nil
	}
	if _, err := txModels.UserDialogsModel.UpdateDialogTopAfterDelete(top.PeerSeq, top.UserMessageId, top.CanonicalMessageId, top.Date, top.MessageStatus, 0, nextPTS, unixNow(), userID, peerType, peerID); err != nil {
		return storageError("update dialog top after delete", err)
	}
	return nil
}

func applyUpdatePinnedMessage(txModels *model.TxModels, in ApplyUserOperationInput, op payload.MessageOperationV1, nextPTS int64) error {
	pinnedPeerSeq := op.PinnedPeerSeq
	if pinnedPeerSeq == 0 {
		pinnedPeerSeq = op.PeerSeq
	}
	pinnedCanonicalID := op.PinnedCanonicalMessageID
	if pinnedCanonicalID == 0 {
		pinnedCanonicalID = op.CanonicalMessageID
	}
	pinnedUserMessageID := op.PinnedUserMessageID
	if pinnedUserMessageID == 0 {
		resolvedID, err := resolveUserMessageIDByPeerSeq(txModels, in.UserID, op.PeerType, op.PeerID, pinnedPeerSeq)
		if err != nil {
			return err
		}
		pinnedUserMessageID = resolvedID
	}
	_, err := txModels.UserDialogsModel.UpdatePinnedMessage(pinnedPeerSeq, pinnedUserMessageID, pinnedCanonicalID, nextPTS, unixNow(), in.UserID, op.PeerType, op.PeerID)
	if err != nil {
		return storageError("apply update pinned message", err)
	}
	return nil
}

func applyMarkDialogUnread(txModels *model.TxModels, in ApplyUserOperationInput, op payload.MessageOperationV1, nextPTS int64) error {
	unreadMark := true
	if op.UnreadMark != nil {
		unreadMark = *op.UnreadMark
	}
	row, err := txModels.UserDialogsModel.SelectByUserPeer(in.UserID, op.PeerType, op.PeerID)
	if err != nil {
		if errors.Is(err, model.ErrNotFound) {
			return nil
		}
		return storageError("select dialog before mark unread", err)
	}
	_, err = txModels.UserDialogsModel.UpdateReadState(
		row.UnreadCount,
		row.UnreadMentionsCount,
		row.UnreadReactionsCount,
		unreadMark,
		row.ReadInboxMaxPeerSeq,
		row.ReadInboxMaxUserMessageId,
		row.ReadOutboxMaxPeerSeq,
		row.ReadOutboxMaxUserMessageId,
		nextPTS,
		unixNow(),
		in.UserID,
		op.PeerType,
		op.PeerID,
	)
	if err != nil {
		return storageError("apply mark dialog unread", err)
	}
	return nil
}

func applyScheduledMarker(txModels *model.TxModels, in ApplyUserOperationInput, op payload.MessageOperationV1, nextPTS int64) error {
	hasScheduled := false
	if op.HasScheduled != nil {
		hasScheduled = *op.HasScheduled
	}
	_, _, err := txModels.UserDialogsModel.InsertOrUpdateMessageEvent(&model.UserDialogs{
		UserId:                     in.UserID,
		PeerType:                   op.PeerType,
		PeerId:                     op.PeerID,
		TopPeerSeq:                 op.PeerSeq,
		TopUserMessageId:           op.UserMessageID,
		TopCanonicalMessageId:      op.CanonicalMessageID,
		TopMessageDate:             int64(op.Date),
		TopMessageStatus:           MessageStatusLive,
		ReadInboxMaxPeerSeq:        op.ReadInboxMaxPeerSeq,
		ReadInboxMaxUserMessageId:  op.ReadInboxMaxUserMessageID,
		ReadOutboxMaxPeerSeq:       op.ReadOutboxMaxPeerSeq,
		ReadOutboxMaxUserMessageId: op.ReadOutboxMaxUserMessageID,
		UnreadCount:                0,
		UnreadMentionsCount:        0,
		UnreadReactionsCount:       0,
		UnreadMark:                 false,
		PinnedPeerSeq:              op.PinnedPeerSeq,
		PinnedUserMessageId:        op.PinnedUserMessageID,
		PinnedCanonicalMessageId:   0,
		HasScheduled:               hasScheduled,
		AvailableMinPeerSeq:        0,
		AvailableMinUserMessageId:  0,
		Hidden:                     false,
		DeletedAt:                  0,
		LastPts:                    nextPTS,
		LastPtsAt:                  unixNow(),
		DialogSchemaVersion:        1,
		DialogPayload:              []byte(`{"schema_version":1}`),
	})
	if err != nil {
		return storageError("apply scheduled marker", err)
	}
	return nil
}

func insertPTSEvent(txModels *model.TxModels, in ApplyUserOperationInput, op messageOperation, pts int64, ptsCount int32, eventPayload []byte, eventPayloadHash []byte) error {
	eventType := EventTypeNewMessage
	switch op.OperationKind {
	case payload.OperationKindReadHistory:
		eventType = EventTypeReadHistory
	case payload.OperationKindDeleteMessages:
		eventType = EventTypeDeleteMessages
	case payload.OperationKindDeleteHistory:
		eventType = EventTypeDeleteHistory
	case payload.OperationKindEditMessage:
		eventType = EventTypeEditMessage
	case payload.OperationKindUpdatePinnedMessage:
		eventType = EventTypeUpdatePinnedMessage
	case payload.OperationKindMarkDialogUnread:
		eventType = EventTypeMarkDialogUnread
	case payload.OperationKindScheduledMarker:
		eventType = EventTypeScheduledMarker
	}
	_, _, err := txModels.UserPtsEventsModel.Insert(&model.UserPtsEvents{
		UserId:             in.UserID,
		Pts:                pts,
		PtsCount:           ptsCount,
		OperationId:        in.OperationID,
		OpType:             in.OpType,
		EventType:          eventType,
		PeerType:           op.PeerType,
		PeerId:             op.PeerID,
		CanonicalMessageId: op.CanonicalMessageID,
		PeerSeq:            op.PeerSeq,
		ActorUserId:        op.FromUserID,
		EventSchemaVersion: int32(op.EventSchemaVersion),
		EventCodec:         PayloadCodecJSON,
		EventPayload:       eventPayload,
		EventPayloadHash:   eventPayloadHash,
	})
	if err != nil {
		return fmt.Errorf("%w: insert pts event: %w", userupdates.ErrPtsContinuityViolation, err)
	}
	return nil
}

func insertPushTask(ctx context.Context, txModels *model.TxModels, r *Repository, in ApplyUserOperationInput, op payload.MessageOperationV1, pts int64, taskPayload []byte) error {
	taskID, err := r.idgen.NextID(ctx)
	if err != nil {
		return storageError("next push task id", err)
	}
	route := payload.RouteUser(in.UserID)
	_, _, err = txModels.PushTaskOutboxModel.Insert(&model.PushTaskOutbox{
		TaskId:             taskID,
		UserId:             in.UserID,
		Pts:                pts,
		PushType:           PushTypeUserUpdate,
		PeerType:           op.PeerType,
		PeerId:             op.PeerID,
		OperationId:        in.OperationID,
		PushPartitionId:    int32(route.PushPartitionID),
		TaskSchemaVersion:  1,
		TaskCodec:          PayloadCodecJSON,
		TaskPayload:        taskPayload,
		Status:             PushTaskStatusPending,
		PublishAttempts:    0,
		AvailableAt:        unixNow(),
		PublishedTopic:     "",
		PublishedPartition: 0,
		PublishedOffset:    0,
		LastErrorCode:      "",
	})
	if err != nil {
		return storageError("insert push task", err)
	}
	return nil
}

func insertDialogSideEffects(ctx context.Context, txModels *model.TxModels, r *Repository, in ApplyUserOperationInput, op payload.MessageOperationV1) error {
	if op.ClearDraft {
		clearBeforeDate := op.ClearDraftBeforeDate
		if clearBeforeDate == 0 {
			clearBeforeDate = op.Date
		}
		body, err := json.Marshal(clearDraftSideEffectPayloadV1{
			SchemaVersion:      1,
			ClearBeforeDate:    clearBeforeDate,
			SourceMessageDate:  op.Date,
			SourcePeerSeq:      op.PeerSeq,
			CanonicalMessageID: op.CanonicalMessageID,
		})
		if err != nil {
			return storageError("marshal clear draft side effect", err)
		}
		sideEffectID, err := r.idgen.NextID(ctx)
		if err != nil {
			return storageError("next dialog side effect id", err)
		}
		if err := r.InsertDialogSideEffectTx(txModels, DialogSideEffect{
			SideEffectID:             sideEffectID,
			Kind:                     DialogSideEffectKindClearDraftAfterSend,
			UserID:                   in.UserID,
			PeerType:                 op.PeerType,
			PeerID:                   op.PeerID,
			SourcePermAuthKeyID:      op.SourcePermAuthKeyID,
			SourceOperationID:        in.OperationID,
			SourceMessageDate:        int64(op.Date),
			SourcePeerSeq:            op.PeerSeq,
			SourceCanonicalMessageID: op.CanonicalMessageID,
			ClearBeforeDate:          int64(clearBeforeDate),
			PayloadSchemaVersion:     1,
			Payload:                  body,
			PayloadHash:              payload.HashBytes(body),
			Status:                   DialogSideEffectStatusPending,
			AttemptCount:             0,
			NextRetryAt:              unixNow(),
		}); err != nil {
			return err
		}
	}
	if shouldWriteSavedDialogSideEffect(in, op) {
		body, err := json.Marshal(savedDialogSideEffectPayloadV1{
			SchemaVersion:         1,
			SavedPeerType:         op.PeerType,
			SavedPeerID:           op.PeerID,
			TopPeerSeq:            op.PeerSeq,
			TopUserMessageID:      op.UserMessageID,
			TopCanonicalMessageID: op.CanonicalMessageID,
			MessageDate:           op.Date,
			Deleted:               false,
			Top:                   true,
		})
		if err != nil {
			return storageError("marshal saved dialog side effect", err)
		}
		sideEffectID, err := r.idgen.NextID(ctx)
		if err != nil {
			return storageError("next dialog side effect id", err)
		}
		if err := r.InsertDialogSideEffectTx(txModels, DialogSideEffect{
			SideEffectID:             sideEffectID,
			Kind:                     DialogSideEffectKindUpsertSavedDialogFromMessage,
			UserID:                   in.UserID,
			PeerType:                 op.PeerType,
			PeerID:                   op.PeerID,
			SourceOperationID:        in.OperationID,
			SourceMessageDate:        int64(op.Date),
			SourcePeerSeq:            op.PeerSeq,
			SourceCanonicalMessageID: op.CanonicalMessageID,
			PayloadSchemaVersion:     1,
			Payload:                  body,
			PayloadHash:              payload.HashBytes(body),
			Status:                   DialogSideEffectStatusPending,
			AttemptCount:             0,
			NextRetryAt:              unixNow(),
		}); err != nil {
			return err
		}
	}
	return nil
}

func shouldWriteSavedDialogSideEffect(in ApplyUserOperationInput, op payload.MessageOperationV1) bool {
	if op.SavedDialogSideEffect {
		return true
	}
	return op.PeerType == payload.PeerTypeUser && op.PeerID == in.UserID
}

func insertOperationResult(txModels *model.TxModels, in ApplyUserOperationInput, pts int64, ptsCount int32, responsePayload []byte, responseHash []byte) error {
	_, _, err := txModels.UserOperationResultsModel.Insert(&model.UserOperationResults{
		UserId:                in.UserID,
		OperationId:           in.OperationID,
		OpType:                in.OpType,
		Status:                OperationResultStatusCompleted,
		Pts:                   pts,
		PtsCount:              ptsCount,
		PayloadHash:           in.PayloadHash,
		ResponseSchemaVersion: payload.OperationResponseSchemaVersion,
		ResponseCodec:         PayloadCodecJSON,
		ResponsePayload:       responsePayload,
		ResponsePayloadHash:   responseHash,
		TerminalErrorCode:     "",
		CompletedAt:           unixNow(),
	})
	if err != nil {
		return storageError("insert operation result", err)
	}
	return nil
}
