package repository

import (
	"bytes"
	"context"
	"errors"
	"time"

	"github.com/teamgram/marmota/pkg/stores/sqlx"
	"github.com/teamgram/teamgram-server/v2/app/messenger/userupdates/internal/repository/model"
	"github.com/teamgram/teamgram-server/v2/app/messenger/userupdates/payload"
	"github.com/teamgram/teamgram-server/v2/app/messenger/userupdates/userupdates"
)

func (r *Repository) AppendDialogAuthSeqSideEffect(ctx context.Context, in DialogSideEffectAppendInput) (*AuthSeqAppendResult, error) {
	db, err := r.requireDB()
	if err != nil {
		return nil, err
	}
	if !bytes.Equal(in.PayloadHash, payload.HashBytes(in.Payload)) {
		return nil, userupdates.ErrOperationPayloadConflict
	}
	var out *AuthSeqAppendResult
	err = db.Transact(ctx, func(tx *sqlx.Tx) error {
		result, err := r.appendDialogAuthSeqSideEffectTx(r.models.WithTx(tx), in)
		if err != nil {
			return err
		}
		out = result
		return nil
	})
	return out, err
}

func (r *Repository) appendDialogAuthSeqSideEffectTx(txModels *model.TxModels, in DialogSideEffectAppendInput) (*AuthSeqAppendResult, error) {
	existing, err := txModels.UserAuthSeqEventsModel.SelectByOperation(in.UserID, in.OperationID)
	if err == nil {
		if !bytes.Equal(existing.EventPayloadHash, in.PayloadHash) {
			return nil, userupdates.ErrOperationPayloadConflict
		}
		return &AuthSeqAppendResult{UserID: in.UserID, OperationID: in.OperationID, Seq: existing.Seq, Date: existing.Date, AlreadyApplied: true}, nil
	}
	if !errors.Is(err, model.ErrNotFound) {
		return nil, storageError("select auth seq event", err)
	}
	if _, _, err := txModels.UserAuthSeqStateModel.InsertIgnore(&model.UserAuthSeqState{UserId: in.UserID, Seq: 0, Date: 0, RowVersion: 0}); err != nil {
		return nil, storageError("init auth seq state", err)
	}
	state, err := txModels.UserAuthSeqStateModel.SelectForUpdate(in.UserID)
	if err != nil {
		return nil, storageError("lock auth seq state", err)
	}
	nextSeq := state.Seq + 1
	eventDate := int32(time.Now().UTC().Unix())
	if eventDate <= state.Date {
		eventDate = state.Date + 1
	}
	if _, _, err := txModels.UserAuthSeqEventsModel.Insert(&model.UserAuthSeqEvents{
		UserId:              in.UserID,
		Seq:                 nextSeq,
		Date:                eventDate,
		OperationId:         in.OperationID,
		SourcePermAuthKeyId: in.SourcePermAuthKeyID,
		TargetAuthPolicy:    in.TargetAuthPolicy,
		PublicUpdateType:    in.PublicUpdateType,
		PeerType:            in.PeerType,
		PeerId:              in.PeerID,
		EventSchemaVersion:  in.PayloadSchemaVersion,
		EventCodec:          PayloadCodecJSON,
		EventPayload:        in.Payload,
		EventPayloadHash:    in.PayloadHash,
	}); err != nil {
		return nil, storageError("insert auth seq event", err)
	}
	if affected, err := txModels.UserAuthSeqStateModel.UpdateSeqDate(nextSeq, eventDate, in.UserID); err != nil {
		return nil, storageError("update auth seq state", err)
	} else if affected == 0 {
		return nil, userupdates.ErrPtsContinuityViolation
	}
	return &AuthSeqAppendResult{UserID: in.UserID, OperationID: in.OperationID, Seq: nextSeq, Date: eventDate}, nil
}

func (r *Repository) AppendDialogPtsSideEffect(ctx context.Context, in DialogSideEffectAppendInput) (*PtsAppendResult, error) {
	db, err := r.requireDB()
	if err != nil {
		return nil, err
	}
	if !bytes.Equal(in.PayloadHash, payload.HashBytes(in.Payload)) {
		return nil, userupdates.ErrOperationPayloadConflict
	}
	var out *PtsAppendResult
	err = db.Transact(ctx, func(tx *sqlx.Tx) error {
		txModels := r.models.WithTx(tx)
		existing, err := txModels.UserPtsEventsModel.SelectByOperation(in.UserID, in.OperationID)
		if err == nil {
			if !bytes.Equal(existing.EventPayloadHash, in.PayloadHash) {
				return userupdates.ErrOperationPayloadConflict
			}
			out = &PtsAppendResult{UserID: in.UserID, OperationID: in.OperationID, Pts: existing.Pts, PtsCount: existing.PtsCount, AlreadyApplied: true}
			return nil
		}
		if !errors.Is(err, model.ErrNotFound) {
			return storageError("select pts side effect event", err)
		}
		if _, _, err := txModels.UserPtsStateModel.InsertIgnore(&model.UserPtsState{UserId: in.UserID, Pts: 0, PtsUpdatedAt: mysqlNow(), PartitionId: 0, OwnerEpoch: 0, RowVersion: 0}); err != nil {
			return storageError("init pts state", err)
		}
		state, err := txModels.UserPtsStateModel.SelectForUpdate(in.UserID)
		if err != nil {
			return storageError("lock pts state", err)
		}
		nextPts := state.Pts + 1
		ptsCount := int32(1)
		if _, _, err := txModels.UserPtsEventsModel.Insert(&model.UserPtsEvents{
			UserId:             in.UserID,
			Pts:                nextPts,
			PtsCount:           ptsCount,
			OperationId:        in.OperationID,
			OpType:             OpTypeSendMessage,
			EventType:          EventTypeDialogPublicUpdate,
			PeerType:           in.PeerType,
			PeerId:             in.PeerID,
			EventSchemaVersion: in.PayloadSchemaVersion,
			EventCodec:         PayloadCodecJSON,
			EventPayload:       in.Payload,
			EventPayloadHash:   in.PayloadHash,
		}); err != nil {
			return storageError("insert pts side effect event", err)
		}
		if affected, err := txModels.UserPtsStateModel.UpdatePts(nextPts, mysqlNow(), 0, 0, in.UserID); err != nil {
			return storageError("update pts state", err)
		} else if affected == 0 {
			return userupdates.ErrPtsContinuityViolation
		}
		out = &PtsAppendResult{UserID: in.UserID, OperationID: in.OperationID, Pts: nextPts, PtsCount: ptsCount}
		return nil
	})
	return out, err
}

func authSeqEventFromModel(r model.UserAuthSeqEvents) AuthSeqEvent {
	return AuthSeqEvent{
		UserID:              r.UserId,
		Seq:                 r.Seq,
		Date:                r.Date,
		OperationID:         r.OperationId,
		SourcePermAuthKeyID: r.SourcePermAuthKeyId,
		TargetAuthPolicy:    r.TargetAuthPolicy,
		PublicUpdateType:    r.PublicUpdateType,
		PeerType:            r.PeerType,
		PeerID:              r.PeerId,
		EventSchemaVersion:  r.EventSchemaVersion,
		EventCodec:          r.EventCodec,
		EventPayload:        r.EventPayload,
		EventPayloadHash:    r.EventPayloadHash,
	}
}
