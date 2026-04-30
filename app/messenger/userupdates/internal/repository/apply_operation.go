package repository

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

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
		LeaseExpiresAt:  "",
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
	fence, err := r.lockPartitionFence(txModels, in.PartitionID)
	if err != nil {
		return nil, err
	}
	if fence.OwnerInstanceId != r.OwnerInstance() {
		return nil, userupdates.ErrNotOwner
	}

	if _, _, err := txModels.UserPtsStateModel.InsertIgnore(&model.UserPtsState{
		UserId:       in.UserID,
		Pts:          0,
		PtsUpdatedAt: mysqlNow(),
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
			UserID:          in.UserID,
			OperationID:     in.OperationID,
			Pts:             existing.Pts,
			PtsCount:        existing.PtsCount,
			ResponsePayload: existing.ResponsePayload,
			ResponseHash:    existing.ResponseHash,
			AlreadyApplied:  true,
		}, nil
	}

	var op payload.MessageOperationV1
	if err := json.Unmarshal(in.Payload, &op); err != nil {
		return nil, fmt.Errorf("%w: decode message operation: %v", userupdates.ErrOperationTerminal, err)
	}
	if op.SchemaVersion != payload.MessageOperationSchemaVersion || op.OperationKind != payload.OperationKindSendMessage {
		return nil, fmt.Errorf("%w: unsupported operation schema=%d kind=%s", userupdates.ErrOperationTerminal, op.SchemaVersion, op.OperationKind)
	}
	if len(in.DependencyPts) != 0 || len(op.DependencyPts) != 0 {
		return nil, userupdates.ErrOperationTerminal
	}

	nextPTS := state.Pts + 1
	ptsCount := int32(1)
	eventPayload, eventPayloadHash, responsePayload, responsePayloadHash, err := buildEventAndResponse(in, op, nextPTS, ptsCount)
	if err != nil {
		return nil, err
	}
	if err := insertUserMessageView(txModels, in, op, eventPayload); err != nil {
		return nil, err
	}
	if err := upsertUserDialog(txModels, in, op, eventPayload); err != nil {
		return nil, err
	}
	if err := insertPTSEvent(txModels, in, op, nextPTS, ptsCount, eventPayload, eventPayloadHash); err != nil {
		return nil, err
	}
	if err := insertPushTask(ctx, txModels, r, in, op, nextPTS, eventPayload); err != nil {
		return nil, err
	}
	if err := insertOperationResult(txModels, in, nextPTS, ptsCount, responsePayload, responsePayloadHash); err != nil {
		return nil, err
	}
	affected, err := txModels.UserPtsStateModel.UpdatePts(nextPTS, mysqlNow(), in.PartitionID, fence.OwnerEpoch, in.UserID)
	if err != nil {
		return nil, storageError("update user pts state", err)
	}
	if affected == 0 {
		return nil, userupdates.ErrPtsContinuityViolation
	}
	return &ApplyUserOperationResult{
		UserID:          in.UserID,
		OperationID:     in.OperationID,
		Pts:             nextPTS,
		PtsCount:        ptsCount,
		ResponsePayload: responsePayload,
		ResponseHash:    responsePayloadHash,
	}, nil
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

func (r *Repository) lockUserPTSState(txModels *model.TxModels, userID int64) (*model.UserPtsState, error) {
	state, err := txModels.UserPtsStateModel.SelectForUpdate(userID)
	if err != nil {
		return nil, storageError("lock user pts state", err)
	}
	return state, nil
}

func buildEventAndResponse(in ApplyUserOperationInput, op payload.MessageOperationV1, pts int64, ptsCount int32) ([]byte, []byte, []byte, []byte, error) {
	event := payload.MessageEventV1{
		SchemaVersion:      payload.MessageEventSchemaVersion,
		EventKind:          payload.EventKindNewMessage,
		CanonicalMessageID: op.CanonicalMessageID,
		MessageID:          op.PeerSeq,
		PeerType:           op.PeerType,
		PeerID:             op.PeerID,
		FromUserID:         op.FromUserID,
		ToUserID:           op.ToUserID,
		Date:               op.Date,
		Out:                op.Out,
		MessageText:        op.MessageText,
		Entities:           op.Entities,
	}
	eventPayload, err := json.Marshal(event)
	if err != nil {
		return nil, nil, nil, nil, storageError("marshal message event", err)
	}
	response := payload.OperationResponseV1{
		SchemaVersion: payload.OperationResponseSchemaVersion,
		OperationID:   in.OperationID,
		Pts:           pts,
		PtsCount:      ptsCount,
		EventType:     payload.EventKindNewMessage,
	}
	responsePayload, err := json.Marshal(response)
	if err != nil {
		return nil, nil, nil, nil, storageError("marshal operation response", err)
	}
	return eventPayload, payload.HashBytes(eventPayload), responsePayload, payload.HashBytes(responsePayload), nil
}

func insertUserMessageView(txModels *model.TxModels, in ApplyUserOperationInput, op payload.MessageOperationV1, viewPayload []byte) error {
	_, _, err := txModels.UserMessageViewsModel.InsertOrUpdate(&model.UserMessageViews{
		UserId:             in.UserID,
		PeerType:           op.PeerType,
		PeerId:             op.PeerID,
		PeerSeq:            op.PeerSeq,
		CanonicalMessageId: op.CanonicalMessageID,
		FromUserId:         op.FromUserID,
		Outgoing:           op.Out,
		MessageKind:        MessageKindText,
		MessageStatus:      MessageStatusLive,
		EditVersion:        0,
		Date:               mysqlDate(op.Date),
		EditDate:           "",
		DeletedAt:          "",
		ViewSchemaVersion:  1,
		ViewPayload:        viewPayload,
	})
	if err != nil {
		return storageError("insert user message view", err)
	}
	return nil
}

func upsertUserDialog(txModels *model.TxModels, in ApplyUserOperationInput, op payload.MessageOperationV1, dialogPayload []byte) error {
	unread := int32(0)
	if !op.Out {
		unread = 1
	}
	_, _, err := txModels.UserDialogsModel.InsertOrUpdateMessageEvent(&model.UserDialogs{
		UserId:                in.UserID,
		PeerType:              op.PeerType,
		PeerId:                op.PeerID,
		TopPeerSeq:            op.PeerSeq,
		TopCanonicalMessageId: op.CanonicalMessageID,
		TopMessageDate:        mysqlDate(op.Date),
		UnreadCount:           unread,
		UnreadMentionsCount:   0,
		ReadInboxMaxPeerSeq:   0,
		ReadOutboxMaxPeerSeq:  0,
		Pinned:                false,
		FolderId:              0,
		DialogSchemaVersion:   1,
		DialogPayload:         dialogPayload,
	})
	if err != nil {
		return storageError("upsert user dialog", err)
	}
	return nil
}

func insertPTSEvent(txModels *model.TxModels, in ApplyUserOperationInput, op payload.MessageOperationV1, pts int64, ptsCount int32, eventPayload []byte, eventPayloadHash []byte) error {
	_, _, err := txModels.UserPtsEventsModel.Insert(&model.UserPtsEvents{
		UserId:             in.UserID,
		Pts:                pts,
		PtsCount:           ptsCount,
		OperationId:        in.OperationID,
		OpType:             in.OpType,
		EventType:          EventTypeNewMessage,
		PeerType:           op.PeerType,
		PeerId:             op.PeerID,
		CanonicalMessageId: op.CanonicalMessageID,
		PeerSeq:            op.PeerSeq,
		ActorUserId:        op.FromUserID,
		EventSchemaVersion: payload.MessageEventSchemaVersion,
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
		AvailableAt:        mysqlNow(),
		NextRetryAt:        "",
		PublishedTopic:     "",
		PublishedPartition: 0,
		PublishedOffset:    0,
		LastErrorCode:      "",
		PublishedAt:        "",
	})
	if err != nil {
		return storageError("insert push task", err)
	}
	return nil
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
		CompletedAt:           mysqlNow(),
	})
	if err != nil {
		return storageError("insert operation result", err)
	}
	return nil
}

func mysqlDate(unix int32) string {
	return time.Unix(int64(unix), 0).UTC().Format("2006-01-02 15:04:05.000000")
}

func mysqlNow() string {
	return time.Now().UTC().Format("2006-01-02 15:04:05.000000")
}
