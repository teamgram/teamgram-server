package repository

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/teamgram/teamgram-server/v2/app/messenger/userupdates/internal/repository/model"
	"github.com/teamgram/teamgram-server/v2/app/messenger/userupdates/payload"
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
			if err := r.fillUnreadCount(ctx, state); err != nil {
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
	if err := r.fillUnreadCount(ctx, state); err != nil {
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

func (r *Repository) fillUnreadCount(ctx context.Context, state *UserState) error {
	row, err := r.models.UserupdatesQueries.SumUnreadDialogs(ctx, state.UserID)
	if err != nil {
		return storageError("sum unread dialogs", err)
	}
	state.UnreadCount = row.Count
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
		event, err := r.userEventFromModel(ctx, row)
		if err != nil {
			return nil, err
		}
		events = append(events, event)
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

func (r *Repository) userEventFromModel(ctx context.Context, row model.UserPtsEvents) (UserEvent, error) {
	event := UserEvent{
		UserID:             row.UserId,
		Pts:                row.Pts,
		PtsCount:           row.PtsCount,
		OperationID:        row.OperationId,
		OpType:             row.OpType,
		EventType:          row.EventType,
		PeerType:           row.PeerType,
		PeerID:             row.PeerId,
		CanonicalMessageID: row.CanonicalMessageId,
		PeerSeq:            row.PeerSeq,
		ActorUserID:        row.ActorUserId,
		EventSchemaVersion: row.EventSchemaVersion,
		EventCodec:         row.EventCodec,
		EventPayload:       row.EventPayload,
		EventPayloadHash:   row.EventPayloadHash,
	}
	if row.EventSchemaVersion != payload.MessageEventSchemaVersionV1 || row.EventCodec != PayloadCodecJSON {
		return event, nil
	}
	hydrated, err := r.hydrateLegacyMessageEvent(ctx, event)
	if err != nil {
		return UserEvent{}, err
	}
	return hydrated, nil
}

func (r *Repository) hydrateLegacyMessageEvent(ctx context.Context, event UserEvent) (UserEvent, error) {
	var old payload.MessageEventV1
	if err := json.Unmarshal(event.EventPayload, &old); err != nil {
		return UserEvent{}, storageError("decode legacy message event", err)
	}
	if old.SchemaVersion != payload.MessageEventSchemaVersionV1 {
		return UserEvent{}, storageError("decode legacy message event", userupdates.ErrOperationTerminal)
	}
	userMessageID, err := r.resolveUserMessageIDByPeerSeq(ctx, event.UserID, event.PeerType, event.PeerID, event.PeerSeq)
	if err != nil {
		return UserEvent{}, err
	}
	if userMessageID <= 0 {
		return UserEvent{}, storageError("hydrate legacy message event", userupdates.ErrOperationTerminal)
	}
	replyToUserMessageID := int64(0)
	if old.ReplyToPeerSeq > 0 {
		replyToUserMessageID, err = r.resolveUserMessageIDByPeerSeq(ctx, event.UserID, event.PeerType, event.PeerID, old.ReplyToPeerSeq)
		if err != nil {
			return UserEvent{}, err
		}
	}
	next := payload.MessageEventV2{
		SchemaVersion:        payload.MessageEventSchemaVersion,
		EventKind:            old.EventKind,
		CanonicalMessageID:   old.CanonicalMessageID,
		PeerSeq:              event.PeerSeq,
		MessageID:            userMessageID,
		PeerType:             old.PeerType,
		PeerID:               old.PeerID,
		FromUserID:           old.FromUserID,
		ToUserID:             old.ToUserID,
		Date:                 old.Date,
		EditDate:             old.EditDate,
		EditVersion:          old.EditVersion,
		Out:                  old.Out,
		MessageText:          old.MessageText,
		Entities:             old.Entities,
		ReplyToUserMessageID: replyToUserMessageID,
		ReadMaxUserMessageID: userMessageID,
		AuthKeyIdExclude:     old.AuthKeyIdExclude,
	}
	body, err := json.Marshal(next)
	if err != nil {
		return UserEvent{}, storageError("marshal hydrated legacy message event", err)
	}
	event.EventSchemaVersion = payload.MessageEventSchemaVersion
	event.EventPayload = body
	event.EventPayloadHash = payload.HashBytes(body)
	return event, nil
}

func (r *Repository) resolveUserMessageIDByPeerSeq(ctx context.Context, userID int64, peerType int32, peerID, peerSeq int64) (int64, error) {
	row, err := r.models.UserMessageViewsModel.SelectNearestLiveByUserPeerSeq(ctx, userID, peerType, peerID, peerSeq, MessageStatusLive)
	if err != nil {
		if errors.Is(err, model.ErrNotFound) {
			return 0, nil
		}
		return 0, storageError("resolve user message id by peer seq", err)
	}
	return row.UserMessageId, nil
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
