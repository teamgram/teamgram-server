package repository

import (
	"context"
	"errors"

	"github.com/teamgram/teamgram-server/v2/app/messenger/userupdates/internal/repository/model"
	"github.com/teamgram/teamgram-server/v2/app/messenger/userupdates/userupdates"
)

func (r *Repository) GetOperationResult(ctx context.Context, userID int64, operationID string) (*OperationResult, error) {
	if _, err := r.requireDB(); err != nil {
		return nil, err
	}
	row, err := r.models.UserOperationResultsModel.SelectByOperation(ctx, userID, operationID)
	if err != nil {
		if errors.Is(err, model.ErrNotFound) {
			return nil, userupdates.ErrOperationTerminal
		}
		return nil, storageError("get operation result", err)
	}
	return operationResultFromModel(row), nil
}

func (r *Repository) GetState(ctx context.Context, userID int64, permAuthKeyID int64) (*UserState, error) {
	if _, err := r.requireDB(); err != nil {
		return nil, err
	}
	row, err := r.models.UserPtsStateModel.SelectByUserId(ctx, userID)
	if err != nil {
		if errors.Is(err, model.ErrNotFound) {
			state := &UserState{UserID: userID}
			if err := r.fillAuthSeqState(ctx, state); err != nil {
				return nil, err
			}
			return state, nil
		}
		return nil, storageError("get state", err)
	}
	state := &UserState{
		UserID:      row.UserId,
		Pts:         row.Pts,
		PartitionID: row.PartitionId,
		OwnerEpoch:  row.OwnerEpoch,
		RowVersion:  row.RowVersion,
	}
	if err := r.fillAuthSeqState(ctx, state); err != nil {
		return nil, err
	}
	return state, nil
}

func (r *Repository) fillAuthSeqState(ctx context.Context, state *UserState) error {
	authState, err := r.models.UserAuthSeqStateModel.SelectByUserId(ctx, state.UserID)
	if err != nil {
		if errors.Is(err, model.ErrNotFound) {
			return nil
		}
		return storageError("get auth seq state", err)
	}
	state.Seq = authState.Seq
	state.Date = authState.Date
	return nil
}

func (r *Repository) GetDifference(ctx context.Context, in GetDifferenceInput) (*GetDifferenceResult, error) {
	if _, err := r.requireDB(); err != nil {
		return nil, err
	}
	limit := in.Limit
	if limit <= 0 {
		limit = 100
	}
	rows, err := r.models.UserPtsEventsModel.SelectByGtPts(ctx, in.UserID, in.Pts, limit)
	if err != nil && !errors.Is(err, model.ErrNotFound) {
		return nil, storageError("get difference events", err)
	}
	state, err := r.GetState(ctx, in.UserID, in.PermAuthKeyID)
	if err != nil {
		return nil, err
	}
	events := make([]UserEvent, 0, len(rows))
	for _, row := range rows {
		events = append(events, userEventFromModel(row))
	}
	var authSeqEvents []AuthSeqEvent
	if in.Date != nil {
		authRows, err := r.models.UserAuthSeqEventsModel.SelectAfterDate(ctx, in.UserID, int32(*in.Date), limit)
		if err != nil && !errors.Is(err, model.ErrNotFound) {
			return nil, storageError("get auth seq events", err)
		}
		authSeqEvents = make([]AuthSeqEvent, 0, len(authRows))
		for _, row := range authRows {
			if row.TargetAuthPolicy == "not_source_perm_auth_key" && row.SourcePermAuthKeyId == in.PermAuthKeyID {
				continue
			}
			authSeqEvents = append(authSeqEvents, authSeqEventFromModel(row))
		}
	}
	return &GetDifferenceResult{State: *state, Events: events, AuthSeqEvents: authSeqEvents}, nil
}

func operationResultFromModel(r *model.UserOperationResults) *OperationResult {
	return &OperationResult{
		UserID:            r.UserId,
		OperationID:       r.OperationId,
		OpType:            r.OpType,
		Status:            r.Status,
		Pts:               r.Pts,
		PtsCount:          r.PtsCount,
		PayloadHash:       r.PayloadHash,
		ResponsePayload:   r.ResponsePayload,
		ResponseHash:      r.ResponsePayloadHash,
		TerminalErrorCode: r.TerminalErrorCode,
	}
}

func userEventFromModel(r model.UserPtsEvents) UserEvent {
	return UserEvent{
		UserID:             r.UserId,
		Pts:                r.Pts,
		PtsCount:           r.PtsCount,
		OperationID:        r.OperationId,
		OpType:             r.OpType,
		EventType:          r.EventType,
		PeerType:           r.PeerType,
		PeerID:             r.PeerId,
		CanonicalMessageID: r.CanonicalMessageId,
		PeerSeq:            r.PeerSeq,
		ActorUserID:        r.ActorUserId,
		EventSchemaVersion: r.EventSchemaVersion,
		EventCodec:         r.EventCodec,
		EventPayload:       r.EventPayload,
		EventPayloadHash:   r.EventPayloadHash,
	}
}

func selectOperationResult(txModels *model.TxModels, userID int64, operationID string) (*OperationResult, bool, error) {
	row, err := txModels.UserOperationResultsModel.SelectByOperation(userID, operationID)
	if err != nil {
		if errors.Is(err, model.ErrNotFound) {
			return nil, false, nil
		}
		return nil, false, storageError("select operation result", err)
	}
	return operationResultFromModel(row), true, nil
}
