package repository

import (
	"context"
	"errors"
	"strings"
	"time"

	"github.com/teamgram/marmota/pkg/stores/sqlx"
	"github.com/teamgram/teamgram-server/v2/app/service/biz/dialog/internal/repository/model"
)

type OutboxRetryState struct {
	AttemptCount     int32
	NextRetryAt      time.Time
	LastErrorKind    string
	LastErrorMessage string
}

func (r *Repository) ClaimDialogAuthSeqOutbox(ctx context.Context, owner string, now time.Time, leaseUntil time.Time, limit int32) ([]model.DialogAuthSeqOutbox, error) {
	if owner == "" || limit <= 0 {
		return []model.DialogAuthSeqOutbox{}, nil
	}
	db, err := r.requireDB()
	if err != nil {
		return nil, err
	}
	claimedLeaseUntil := mysqlDateTimeForBind(leaseUntil)
	rows := []model.DialogAuthSeqOutbox{}
	err = db.Transact(ctx, func(tx *sqlx.Tx) error {
		txModels := r.model.WithTx(tx)
		candidates, err := txModels.DialogRepositoryQueries.SelectAuthSeqOutboxClaimCandidates(
			OutboxStatusPending,
			OutboxStatusFailedRetryable,
			mysqlDateTimeForBind(now),
			OutboxStatusPublishing,
			mysqlDateTimeForBind(now),
			limit,
		)
		if err != nil {
			if errors.Is(err, model.ErrNotFound) || errors.Is(err, sqlx.ErrNotFound) {
				return nil
			}
			return storageError("select claimed dialog auth seq outbox", err)
		}
		rows = make([]model.DialogAuthSeqOutbox, 0, len(candidates))
		for i := range candidates {
			row := candidates[i]
			if _, err := txModels.DialogAuthSeqOutboxModel.MarkPublishing(OutboxStatusPublishing, owner, claimedLeaseUntil, row.OutboxId); err != nil {
				return storageError("claim dialog auth seq outbox", err)
			}
			rows = append(rows, makeClaimedDialogAuthSeqOutbox(row, owner, claimedLeaseUntil))
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return rows, nil
}

func (r *Repository) MarkDialogAuthSeqOutboxPublished(ctx context.Context, outboxID int64, seq int64, date int32) error {
	_, err := r.requireDB()
	if err != nil {
		return err
	}
	if _, err := r.model.DialogAuthSeqOutboxModel.MarkPublished(ctx, OutboxStatusPublished, seq, date, mysqlZeroDateTime(), outboxID); err != nil {
		return storageError("mark dialog auth seq outbox published", err)
	}
	return nil
}

func (r *Repository) MarkDialogAuthSeqOutboxRetryable(ctx context.Context, outboxID int64, retry OutboxRetryState) error {
	_, err := r.requireDB()
	if err != nil {
		return err
	}
	if _, err := r.model.DialogAuthSeqOutboxModel.MarkRetryable(ctx, OutboxStatusFailedRetryable, retry.AttemptCount, mysqlDateTimeForBind(retry.NextRetryAt), mysqlZeroDateTime(), retry.LastErrorKind, truncateOutboxError(retry.LastErrorMessage), outboxID); err != nil {
		return storageError("mark dialog auth seq outbox retryable", err)
	}
	return nil
}

func (r *Repository) MarkDialogAuthSeqOutboxBlocked(ctx context.Context, outboxID int64, kind string, message string) error {
	_, err := r.requireDB()
	if err != nil {
		return err
	}
	if _, err := r.model.DialogAuthSeqOutboxModel.MarkBlocked(ctx, OutboxStatusBlocked, mysqlZeroDateTime(), kind, truncateOutboxError(message), outboxID); err != nil {
		return storageError("mark dialog auth seq outbox blocked", err)
	}
	return nil
}

func (r *Repository) ResetDialogAuthSeqOutboxBlocked(ctx context.Context, ids []int64) error {
	_, err := r.requireDB()
	if err != nil {
		return err
	}
	for _, id := range ids {
		if _, err := r.model.DialogAuthSeqOutboxModel.ResetBlocked(ctx, OutboxStatusPending, mysqlDateTimeForBind(time.Now().UTC()), mysqlZeroDateTime(), OutboxStatusBlocked, id); err != nil {
			return storageError("reset dialog auth seq outbox blocked", err)
		}
	}
	return nil
}

func (r *Repository) ClaimDialogPublicUpdateOutbox(ctx context.Context, owner string, now time.Time, leaseUntil time.Time, limit int32) ([]model.DialogPublicUpdateOutbox, error) {
	if owner == "" || limit <= 0 {
		return []model.DialogPublicUpdateOutbox{}, nil
	}
	db, err := r.requireDB()
	if err != nil {
		return nil, err
	}
	claimedLeaseUntil := mysqlDateTimeForBind(leaseUntil)
	rows := []model.DialogPublicUpdateOutbox{}
	err = db.Transact(ctx, func(tx *sqlx.Tx) error {
		txModels := r.model.WithTx(tx)
		candidates, err := txModels.DialogRepositoryQueries.SelectPublicUpdateOutboxClaimCandidates(
			OutboxStatusPending,
			OutboxStatusFailedRetryable,
			mysqlDateTimeForBind(now),
			OutboxStatusPublishing,
			mysqlDateTimeForBind(now),
			limit,
		)
		if err != nil {
			if errors.Is(err, model.ErrNotFound) || errors.Is(err, sqlx.ErrNotFound) {
				return nil
			}
			return storageError("select claimed dialog public update outbox", err)
		}
		rows = make([]model.DialogPublicUpdateOutbox, 0, len(candidates))
		for i := range candidates {
			row := candidates[i]
			if _, err := txModels.DialogPublicUpdateOutboxModel.MarkPublishing(OutboxStatusPublishing, owner, claimedLeaseUntil, row.OutboxId); err != nil {
				return storageError("claim dialog public update outbox", err)
			}
			rows = append(rows, makeClaimedDialogPublicUpdateOutbox(row, owner, claimedLeaseUntil))
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return rows, nil
}

func (r *Repository) MarkDialogPublicUpdateOutboxPublishedPTS(ctx context.Context, outboxID int64, pts int64, ptsCount int32) error {
	_, err := r.requireDB()
	if err != nil {
		return err
	}
	if _, err := r.model.DialogPublicUpdateOutboxModel.MarkPublishedPTS(ctx, OutboxStatusPublished, pts, ptsCount, mysqlZeroDateTime(), outboxID); err != nil {
		return storageError("mark dialog public update outbox pts published", err)
	}
	return nil
}

func (r *Repository) MarkDialogPublicUpdateOutboxPublishedAuthSeq(ctx context.Context, outboxID int64, seq int64, date int32) error {
	_, err := r.requireDB()
	if err != nil {
		return err
	}
	if _, err := r.model.DialogPublicUpdateOutboxModel.MarkPublishedAuthSeq(ctx, OutboxStatusPublished, seq, date, mysqlZeroDateTime(), outboxID); err != nil {
		return storageError("mark dialog public update outbox auth seq published", err)
	}
	return nil
}

func (r *Repository) MarkDialogPublicUpdateOutboxRetryable(ctx context.Context, outboxID int64, retry OutboxRetryState) error {
	_, err := r.requireDB()
	if err != nil {
		return err
	}
	if _, err := r.model.DialogPublicUpdateOutboxModel.MarkRetryable(ctx, OutboxStatusFailedRetryable, retry.AttemptCount, mysqlDateTimeForBind(retry.NextRetryAt), mysqlZeroDateTime(), retry.LastErrorKind, truncateOutboxError(retry.LastErrorMessage), outboxID); err != nil {
		return storageError("mark dialog public update outbox retryable", err)
	}
	return nil
}

func (r *Repository) MarkDialogPublicUpdateOutboxBlocked(ctx context.Context, outboxID int64, kind string, message string) error {
	_, err := r.requireDB()
	if err != nil {
		return err
	}
	if _, err := r.model.DialogPublicUpdateOutboxModel.MarkBlocked(ctx, OutboxStatusBlocked, mysqlZeroDateTime(), kind, truncateOutboxError(message), outboxID); err != nil {
		return storageError("mark dialog public update outbox blocked", err)
	}
	return nil
}

func (r *Repository) ResetDialogPublicUpdateOutboxBlocked(ctx context.Context, ids []int64) error {
	_, err := r.requireDB()
	if err != nil {
		return err
	}
	for _, id := range ids {
		if _, err := r.model.DialogPublicUpdateOutboxModel.ResetBlocked(ctx, OutboxStatusPending, mysqlDateTimeForBind(time.Now().UTC()), mysqlZeroDateTime(), OutboxStatusBlocked, id); err != nil {
			return storageError("reset dialog public update outbox blocked", err)
		}
	}
	return nil
}

func makeClaimedDialogAuthSeqOutbox(row model.DialogAuthSeqOutboxRow, owner string, leaseUntil time.Time) model.DialogAuthSeqOutbox {
	return model.DialogAuthSeqOutbox{
		OutboxId:             row.OutboxId,
		UserId:               row.UserId,
		SourcePermAuthKeyId:  row.SourcePermAuthKeyId,
		TargetAuthPolicy:     row.TargetAuthPolicy,
		OperationId:          row.OperationId,
		EventType:            row.EventType,
		PeerType:             row.PeerType,
		PeerId:               row.PeerId,
		PayloadSchemaVersion: row.PayloadSchemaVersion,
		Payload:              row.Payload,
		PayloadHash:          row.PayloadHash,
		Status:               OutboxStatusPublishing,
		AttemptCount:         row.AttemptCount,
		NextRetryAt:          row.NextRetryAt,
		LeaseOwner:           owner,
		LeaseUntil:           leaseUntil,
		PublishedSeq:         row.PublishedSeq,
		PublishedDate:        row.PublishedDate,
		LastErrorKind:        row.LastErrorKind,
		LastErrorMessage:     row.LastErrorMessage,
	}
}

func makeClaimedDialogPublicUpdateOutbox(row model.DialogPublicUpdateOutboxRow, owner string, leaseUntil time.Time) model.DialogPublicUpdateOutbox {
	return model.DialogPublicUpdateOutbox{
		OutboxId:             row.OutboxId,
		SourceUserId:         row.SourceUserId,
		SourcePermAuthKeyId:  row.SourcePermAuthKeyId,
		TargetUserId:         row.TargetUserId,
		TargetAuthPolicy:     row.TargetAuthPolicy,
		OperationId:          row.OperationId,
		DeliveryPath:         row.DeliveryPath,
		PublicUpdateType:     row.PublicUpdateType,
		PeerType:             row.PeerType,
		PeerId:               row.PeerId,
		PayloadSchemaVersion: row.PayloadSchemaVersion,
		Payload:              row.Payload,
		PayloadHash:          row.PayloadHash,
		Status:               OutboxStatusPublishing,
		AttemptCount:         row.AttemptCount,
		NextRetryAt:          row.NextRetryAt,
		LeaseOwner:           owner,
		LeaseUntil:           leaseUntil,
		PublishedPts:         row.PublishedPts,
		PublishedPtsCount:    row.PublishedPtsCount,
		PublishedSeq:         row.PublishedSeq,
		PublishedDate:        row.PublishedDate,
		LastErrorKind:        row.LastErrorKind,
		LastErrorMessage:     row.LastErrorMessage,
	}
}

func truncateOutboxError(s string) string {
	const max = 512
	s = strings.TrimSpace(s)
	if len(s) <= max {
		return s
	}
	return s[:max]
}
