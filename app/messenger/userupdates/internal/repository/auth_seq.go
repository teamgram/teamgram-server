package repository

import (
	"bytes"
	"context"
	"errors"
	"time"

	"github.com/teamgram/marmota/pkg/stores/sqlx"
	"github.com/teamgram/teamgram-server/v2/app/messenger/userupdates/internal/cursor"
	"github.com/teamgram/teamgram-server/v2/app/messenger/userupdates/internal/repository/model"
	"github.com/teamgram/teamgram-server/v2/app/messenger/userupdates/payload"
	"github.com/teamgram/teamgram-server/v2/app/messenger/userupdates/userupdates"
)

func (r *Repository) AppendAuthSeqUpdate(ctx context.Context, in AuthSeqUpdateAppendInput) (*AuthSeqUpdateAppendResult, error) {
	db, err := r.requireDB()
	if err != nil {
		return nil, err
	}
	if len(in.TargetPermAuthKeyIDs) == 0 {
		return &AuthSeqUpdateAppendResult{UserID: in.UserID, OperationID: in.OperationID}, nil
	}
	if !bytes.Equal(in.PayloadHash, payload.HashBytes(in.TLBytes)) {
		return nil, userupdates.ErrOperationPayloadConflict
	}
	payloadID := AuthSeqPayloadID(in.PayloadHash)
	if payloadID == "" {
		return nil, userupdates.ErrOperationPayloadConflict
	}
	now := in.Now
	if now == 0 {
		now = time.Now().UTC().Unix()
	}
	var out *AuthSeqUpdateAppendResult
	err = db.Transact(ctx, func(tx *sqlx.Tx) error {
		result, err := r.appendAuthSeqUpdateTx(r.models.WithTx(tx), in, payloadID, now)
		if err != nil {
			return err
		}
		out = result
		return nil
	})
	return out, err
}

func (r *Repository) appendAuthSeqUpdateTx(txModels *model.TxModels, in AuthSeqUpdateAppendInput, payloadID string, now int64) (*AuthSeqUpdateAppendResult, error) {
	if _, _, err := txModels.AuthUpdatePayloadsModel.InsertIgnore(&model.AuthUpdatePayloads{
		PayloadId:   payloadID,
		UpdateType:  in.UpdateType,
		Codec:       AuthSeqCodecTLBinary,
		Layer:       in.Layer,
		TlBytes:     in.TLBytes,
		PayloadHash: in.PayloadHash,
		ExpireAt:    in.ExpireAt,
	}); err != nil {
		return nil, storageError("insert auth update payload", err)
	}
	result := &AuthSeqUpdateAppendResult{UserID: in.UserID, OperationID: in.OperationID}
	for _, authKeyID := range in.TargetPermAuthKeyIDs {
		existing, err := txModels.AuthSeqDeliveriesModel.SelectByOperation(in.UserID, authKeyID, in.OperationID)
		if err == nil {
			event, err := authSeqDeliveryEventFromRow(existing, in, payloadID)
			if err != nil {
				return nil, err
			}
			result.AlreadyApplied = true
			result.Deliveries = append(result.Deliveries, event)
			continue
		}
		if !errors.Is(err, model.ErrNotFound) {
			return nil, storageError("select auth seq delivery", err)
		}
		delivery, alreadyApplied, err := r.appendOneAuthSeqDeliveryTx(txModels, in, authKeyID, payloadID, now)
		if err != nil {
			return nil, err
		}
		if alreadyApplied {
			result.AlreadyApplied = true
		}
		result.Deliveries = append(result.Deliveries, delivery)
	}
	return result, nil
}

func (r *Repository) appendOneAuthSeqDeliveryTx(txModels *model.TxModels, in AuthSeqUpdateAppendInput, authKeyID int64, payloadID string, now int64) (AuthSeqDeliveryEvent, bool, error) {
	if _, _, err := txModels.AuthSeqStateModel.InsertIgnore(&model.AuthSeqState{
		UserId:        in.UserID,
		PermAuthKeyId: authKeyID,
		Seq:           0,
		Date:          0,
		RowVersion:    0,
	}); err != nil {
		return AuthSeqDeliveryEvent{}, false, storageError("init auth seq state", err)
	}
	state, err := txModels.AuthSeqStateModel.SelectForUpdate(in.UserID, authKeyID)
	if err != nil {
		return AuthSeqDeliveryEvent{}, false, storageError("lock auth seq state", err)
	}
	nextSeq := state.Seq + 1
	nextDate64 := now
	if nextDate64 <= state.Date {
		nextDate64 = state.Date + 1
	}
	if nextDate64 > now+AuthSeqMaxFutureSkewSeconds {
		return AuthSeqDeliveryEvent{}, false, userupdates.ErrPtsContinuityViolation
	}
	nextDate, err := cursor.CheckedInt32(nextDate64, "auth seq date")
	if err != nil {
		return AuthSeqDeliveryEvent{}, false, storageError("auth seq date", err)
	}
	if _, rowsAffected, err := txModels.AuthSeqDeliveriesModel.InsertIgnore(&model.AuthSeqDeliveries{
		UserId:              in.UserID,
		PermAuthKeyId:       authKeyID,
		Seq:                 nextSeq,
		Date:                int64(nextDate),
		PayloadId:           payloadID,
		ReplayPolicy:        in.ReplayPolicy,
		SourcePermAuthKeyId: in.SourcePermAuthKeyID,
		VisibilityPolicy:    in.VisibilityPolicy,
		OperationId:         in.OperationID,
		ExpireAt:            in.ExpireAt,
	}); err != nil {
		return AuthSeqDeliveryEvent{}, false, storageError("insert auth seq delivery", err)
	} else if rowsAffected == 0 {
		existing, err := txModels.AuthSeqDeliveriesModel.SelectByOperation(in.UserID, authKeyID, in.OperationID)
		if err != nil {
			return AuthSeqDeliveryEvent{}, false, storageError("select ignored auth seq delivery", err)
		}
		event, err := authSeqDeliveryEventFromRow(existing, in, payloadID)
		if err != nil {
			return AuthSeqDeliveryEvent{}, false, err
		}
		return event, true, nil
	}
	if affected, err := txModels.AuthSeqStateModel.UpdateSeqDate(nextSeq, int64(nextDate), in.UserID, authKeyID); err != nil {
		return AuthSeqDeliveryEvent{}, false, storageError("update auth seq state", err)
	} else if affected == 0 {
		return AuthSeqDeliveryEvent{}, false, userupdates.ErrPtsContinuityViolation
	}
	return AuthSeqDeliveryEvent{
		UserID:              in.UserID,
		PermAuthKeyID:       authKeyID,
		Seq:                 nextSeq,
		Date:                nextDate,
		PayloadID:           payloadID,
		ReplayPolicy:        in.ReplayPolicy,
		OperationID:         in.OperationID,
		SourcePermAuthKeyID: in.SourcePermAuthKeyID,
		VisibilityPolicy:    in.VisibilityPolicy,
		TLBytes:             in.TLBytes,
		PayloadHash:         in.PayloadHash,
		Layer:               in.Layer,
	}, false, nil
}

func authSeqDeliveryEventFromRow(row *model.AuthSeqDeliveries, in AuthSeqUpdateAppendInput, payloadID string) (AuthSeqDeliveryEvent, error) {
	if row.PayloadId != payloadID {
		return AuthSeqDeliveryEvent{}, userupdates.ErrOperationPayloadConflict
	}
	date, err := cursor.CheckedInt32(row.Date, "auth seq date")
	if err != nil {
		return AuthSeqDeliveryEvent{}, storageError("auth seq date", err)
	}
	return AuthSeqDeliveryEvent{
		UserID:              row.UserId,
		PermAuthKeyID:       row.PermAuthKeyId,
		Seq:                 row.Seq,
		Date:                date,
		PayloadID:           row.PayloadId,
		ReplayPolicy:        row.ReplayPolicy,
		OperationID:         row.OperationId,
		SourcePermAuthKeyID: row.SourcePermAuthKeyId,
		VisibilityPolicy:    row.VisibilityPolicy,
		TLBytes:             in.TLBytes,
		PayloadHash:         in.PayloadHash,
		Layer:               in.Layer,
	}, nil
}

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
		if _, _, err := txModels.UserPtsStateModel.InsertIgnore(&model.UserPtsState{UserId: in.UserID, Pts: 0, PtsUpdatedAt: unixNow(), PartitionId: 0, OwnerEpoch: 0, RowVersion: 0}); err != nil {
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
		if affected, err := txModels.UserPtsStateModel.UpdatePts(nextPts, unixNow(), 0, 0, in.UserID); err != nil {
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
