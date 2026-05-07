package repository

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/teamgram/marmota/pkg/stores/sqlx"
	"github.com/teamgram/teamgram-server/v2/app/messenger/userupdates/internal/repository/model"
	"github.com/teamgram/teamgram-server/v2/app/messenger/userupdates/payload"
)

type savedDialogSideEffectPayloadV1 struct {
	SchemaVersion         int   `json:"schema_version"`
	SavedPeerType         int32 `json:"saved_peer_type"`
	SavedPeerID           int64 `json:"saved_peer_id"`
	TopPeerSeq            int64 `json:"top_peer_seq"`
	TopCanonicalMessageID int64 `json:"top_canonical_message_id"`
	MessageDate           int32 `json:"message_date"`
	Deleted               bool  `json:"deleted"`
	Top                   bool  `json:"top"`
}

type clearDraftSideEffectPayloadV1 struct {
	SchemaVersion      int   `json:"schema_version"`
	ClearBeforeDate    int32 `json:"clear_before_date"`
	SourceMessageDate  int32 `json:"source_message_date"`
	SourcePeerSeq      int64 `json:"source_peer_seq"`
	CanonicalMessageID int64 `json:"canonical_message_id"`
}

func (r *Repository) InsertDialogSideEffectTx(txModels *model.TxModels, row DialogSideEffect) error {
	if txModels == nil || txModels.DialogSideEffectOutboxModel == nil {
		return storageError("insert dialog side effect", errors.New("dialog side effect model is not configured"))
	}
	if row.PayloadSchemaVersion == 0 {
		row.PayloadSchemaVersion = 1
	}
	if len(row.PayloadHash) == 0 && len(row.Payload) != 0 {
		row.PayloadHash = payload.HashBytes(row.Payload)
	}
	if row.Status == 0 {
		row.Status = DialogSideEffectStatusPending
	}
	if row.NextRetryAt <= 0 {
		row.NextRetryAt = unixNow()
	}
	_, rowsAffected, err := txModels.DialogSideEffectOutboxModel.Insert(toDialogSideEffectModel(row))
	if err != nil {
		return storageError("insert dialog side effect", err)
	}
	if rowsAffected == 0 {
		return validateExistingDialogSideEffectTx(txModels, row)
	}
	return nil
}

func validateExistingDialogSideEffectTx(txModels *model.TxModels, row DialogSideEffect) error {
	existing, err := txModels.DialogSideEffectOutboxModel.SelectExistingSideEffect(row.Kind, row.SourceOperationID)
	if err != nil {
		return storageError("select existing dialog side effect", err)
	}
	if !bytes.Equal(existing.PayloadHash, row.PayloadHash) {
		return storageError("insert dialog side effect", fmt.Errorf("payload conflict for kind=%s source_operation_id=%s", row.Kind, row.SourceOperationID))
	}
	if existing.UserId != row.UserID || existing.PeerType != row.PeerType || existing.PeerId != row.PeerID {
		return storageError("insert dialog side effect", fmt.Errorf("peer conflict for kind=%s source_operation_id=%s", row.Kind, row.SourceOperationID))
	}
	return nil
}

func (r *Repository) ClaimDialogSideEffects(ctx context.Context, now time.Time, limit int32) ([]DialogSideEffect, error) {
	db, err := r.requireDB()
	if err != nil {
		return nil, err
	}
	if limit <= 0 || limit > OutboxWorkerBatchSize {
		limit = OutboxWorkerBatchSize
	}
	var out []DialogSideEffect
	err = db.Transact(ctx, func(tx *sqlx.Tx) error {
		txModels := r.models.WithTx(tx)
		rows, err := txModels.DialogSideEffectOutboxModel.SelectPendingForUpdate(
			DialogSideEffectStatusPending,
			DialogSideEffectStatusFailedRetryable,
			now.UTC().Unix(),
			DialogSideEffectStatusPublishing,
			limit,
		)
		if err != nil {
			if errors.Is(err, sqlx.ErrNotFound) {
				out = []DialogSideEffect{}
				return nil
			}
			return storageError("claim dialog side effects", err)
		}
		leaseUntil := now.UTC().Add(5 * time.Minute).Unix()
		out = make([]DialogSideEffect, 0, len(rows))
		for i := range rows {
			row := rows[i]
			affected, err := txModels.DialogSideEffectOutboxModel.MarkPublishing(
				DialogSideEffectStatusPublishing,
				r.OwnerInstance(),
				leaseUntil,
				row.SideEffectId,
			)
			if err != nil {
				return storageError("mark dialog side effect publishing", err)
			}
			if affected == 0 {
				return storageError("mark dialog side effect publishing", sqlx.ErrNotFound)
			}
			row.Status = DialogSideEffectStatusPublishing
			row.AttemptCount++
			row.LeaseOwner = r.OwnerInstance()
			row.LeaseUntil = leaseUntil
			dto, err := fromDialogSideEffectModel(row)
			if err != nil {
				return err
			}
			out = append(out, dto)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (r *Repository) ClaimDialogSideEffectsByKind(ctx context.Context, kind string, now time.Time, limit int32) ([]DialogSideEffect, error) {
	db, err := r.requireDB()
	if err != nil {
		return nil, err
	}
	if limit <= 0 || limit > OutboxWorkerBatchSize {
		limit = OutboxWorkerBatchSize
	}
	var out []DialogSideEffect
	err = db.Transact(ctx, func(tx *sqlx.Tx) error {
		txModels := r.models.WithTx(tx)
		rows, err := txModels.DialogSideEffectOutboxModel.SelectPendingForUpdateByKind(
			kind,
			DialogSideEffectStatusPending,
			DialogSideEffectStatusFailedRetryable,
			now.UTC().Unix(),
			DialogSideEffectStatusPublishing,
			limit,
		)
		if err != nil {
			if errors.Is(err, sqlx.ErrNotFound) {
				out = []DialogSideEffect{}
				return nil
			}
			return storageError("claim dialog side effects by kind", err)
		}
		leaseUntil := now.UTC().Add(5 * time.Minute).Unix()
		out = make([]DialogSideEffect, 0, len(rows))
		for i := range rows {
			row := rows[i]
			affected, err := txModels.DialogSideEffectOutboxModel.MarkPublishing(
				DialogSideEffectStatusPublishing,
				r.OwnerInstance(),
				leaseUntil,
				row.SideEffectId,
			)
			if err != nil {
				return storageError("mark dialog side effect publishing", err)
			}
			if affected == 0 {
				return storageError("mark dialog side effect publishing", sqlx.ErrNotFound)
			}
			row.Status = DialogSideEffectStatusPublishing
			row.AttemptCount++
			row.LeaseOwner = r.OwnerInstance()
			row.LeaseUntil = leaseUntil
			dto, err := fromDialogSideEffectModel(row)
			if err != nil {
				return err
			}
			out = append(out, dto)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (r *Repository) MarkDialogSideEffectCompleted(ctx context.Context, sideEffectID int64) error {
	rows, err := r.models.DialogSideEffectOutboxModel.MarkCompleted(
		ctx,
		DialogSideEffectStatusCompleted,
		0,
		sideEffectID,
	)
	if err != nil {
		return storageError("mark dialog side effect completed", err)
	}
	if rows == 0 {
		return storageError("mark dialog side effect completed", sqlx.ErrNotFound)
	}
	return nil
}

func (r *Repository) MarkDialogSideEffectRetryableFailure(ctx context.Context, sideEffectID int64, errCode string, now time.Time) error {
	row, err := r.models.DialogSideEffectOutboxModel.SelectOne(ctx, sideEffectID)
	if err != nil {
		return storageError("select dialog side effect for retry", err)
	}
	if row.AttemptCount >= BlockedAfterAttempts-1 {
		return r.MarkDialogSideEffectBlocked(ctx, sideEffectID, errCode)
	}
	blockedBefore := now.UTC().Add(-time.Duration(BlockedAfterAgeSeconds) * time.Second).Unix()
	rows, err := r.models.DialogSideEffectOutboxModel.MarkBlockedIfOld(
		ctx,
		DialogSideEffectStatusBlocked,
		0,
		errCode,
		sideEffectID,
		blockedBefore,
	)
	if err != nil {
		return storageError("mark old dialog side effect blocked", err)
	}
	if rows > 0 {
		return nil
	}
	nextRetryAt := now.UTC().Unix() + int64(dialogSideEffectRetryDelay(sideEffectID, row.AttemptCount).Seconds())
	rows, err = r.models.DialogSideEffectOutboxModel.MarkRetryableFailure(
		ctx,
		DialogSideEffectStatusFailedRetryable,
		nextRetryAt,
		0,
		errCode,
		sideEffectID,
	)
	if err != nil {
		return storageError("mark dialog side effect retryable failure", err)
	}
	if rows == 0 {
		return storageError("mark dialog side effect retryable failure", sqlx.ErrNotFound)
	}
	return nil
}

func (r *Repository) MarkDialogSideEffectBlocked(ctx context.Context, sideEffectID int64, errCode string) error {
	rows, err := r.models.DialogSideEffectOutboxModel.MarkBlocked(
		ctx,
		DialogSideEffectStatusBlocked,
		0,
		errCode,
		sideEffectID,
	)
	if err != nil {
		return storageError("mark dialog side effect blocked", err)
	}
	if rows == 0 {
		return storageError("mark dialog side effect blocked", sqlx.ErrNotFound)
	}
	return nil
}

func (r *Repository) ResetDialogSideEffectOutboxBlocked(ctx context.Context, ids []int64) error {
	if _, err := r.requireDB(); err != nil {
		return err
	}
	for _, id := range ids {
		_, err := r.models.DialogSideEffectOutboxModel.ResetBlocked(
			ctx,
			DialogSideEffectStatusPending,
			unixNow(),
			0,
			DialogSideEffectStatusBlocked,
			id,
		)
		if err != nil {
			return storageError("reset dialog side effect outbox blocked", err)
		}
	}
	return nil
}

func dialogSideEffectRetryDelay(sideEffectID int64, attemptCount int32) time.Duration {
	if attemptCount <= 1 {
		return time.Duration(InitialRetryDelaySeconds) * time.Second
	}
	seconds := InitialRetryDelaySeconds
	for i := int32(1); i < attemptCount; i++ {
		seconds *= 2
		if seconds >= MaxRetryDelaySeconds {
			return time.Duration(MaxRetryDelaySeconds) * time.Second
		}
	}
	jitterRange := seconds / 4
	if jitterRange < 1 {
		jitterRange = 1
	}
	jitter := int(sideEffectID % int64(jitterRange+1))
	seconds += jitter
	if seconds > MaxRetryDelaySeconds {
		seconds = MaxRetryDelaySeconds
	}
	return time.Duration(seconds) * time.Second
}

func toDialogSideEffectModel(row DialogSideEffect) *model.DialogSideEffectOutbox {
	return &model.DialogSideEffectOutbox{
		SideEffectId:             row.SideEffectID,
		Kind:                     row.Kind,
		UserId:                   row.UserID,
		PeerType:                 row.PeerType,
		PeerId:                   row.PeerID,
		SourcePermAuthKeyId:      row.SourcePermAuthKeyID,
		SourceOperationId:        row.SourceOperationID,
		SourceMessageDate:        row.SourceMessageDate,
		SourcePeerSeq:            row.SourcePeerSeq,
		SourceCanonicalMessageId: row.SourceCanonicalMessageID,
		ClearBeforeDate:          unixOrZero(row.ClearBeforeDate),
		PayloadSchemaVersion:     row.PayloadSchemaVersion,
		Payload:                  row.Payload,
		PayloadHash:              row.PayloadHash,
		Status:                   row.Status,
		AttemptCount:             row.AttemptCount,
		NextRetryAt:              unixOrZero(row.NextRetryAt),
		LeaseOwner:               row.LeaseOwner,
		LeaseUntil:               unixOrZero(row.LeaseUntil),
		LastErrorCode:            row.LastErrorCode,
	}
}

func fromDialogSideEffectModel(row model.DialogSideEffectOutbox) (DialogSideEffect, error) {
	return DialogSideEffect{
		SideEffectID:             row.SideEffectId,
		Kind:                     row.Kind,
		UserID:                   row.UserId,
		PeerType:                 row.PeerType,
		PeerID:                   row.PeerId,
		SourcePermAuthKeyID:      row.SourcePermAuthKeyId,
		SourceOperationID:        row.SourceOperationId,
		SourceMessageDate:        row.SourceMessageDate,
		SourcePeerSeq:            row.SourcePeerSeq,
		SourceCanonicalMessageID: row.SourceCanonicalMessageId,
		ClearBeforeDate:          unixOrZero(row.ClearBeforeDate),
		PayloadSchemaVersion:     row.PayloadSchemaVersion,
		Payload:                  row.Payload,
		PayloadHash:              row.PayloadHash,
		Status:                   row.Status,
		AttemptCount:             row.AttemptCount,
		NextRetryAt:              row.NextRetryAt,
		LeaseOwner:               row.LeaseOwner,
		LeaseUntil:               unixOrZero(row.LeaseUntil),
		LastErrorCode:            row.LastErrorCode,
	}, nil
}
