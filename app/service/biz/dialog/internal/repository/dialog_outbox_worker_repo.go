package repository

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/teamgram/marmota/pkg/stores/sqlx"
	"github.com/teamgram/teamgram-server/v2/app/service/biz/dialog/internal/repository/model"
)

const (
	dialogAuthSeqOutboxSelectRows = "outbox_id,user_id,source_perm_auth_key_id,target_auth_policy,operation_id,event_type,peer_type,peer_id,payload_schema_version,payload,payload_hash,status,attempt_count,next_retry_at,lease_owner,lease_until,published_seq,published_date,last_error_kind,last_error_message"

	dialogPublicUpdateOutboxSelectRows = "outbox_id,source_user_id,source_perm_auth_key_id,target_user_id,target_auth_policy,operation_id,delivery_path,public_update_type,peer_type,peer_id,payload_schema_version,payload,payload_hash,status,attempt_count,next_retry_at,lease_owner,lease_until,published_pts,published_pts_count,published_seq,published_date,last_error_kind,last_error_message"
)

type OutboxRetryState struct {
	AttemptCount     int32
	NextRetryAt      time.Time
	LastErrorKind    string
	LastErrorMessage string
}

func (r *Repository) ClaimDialogAuthSeqOutbox(ctx context.Context, owner string, now time.Time, leaseUntil time.Time, limit int32) ([]model.DialogAuthSeqOutbox, error) {
	db, err := r.requireDB()
	if err != nil {
		return nil, err
	}
	if owner == "" || limit <= 0 {
		return []model.DialogAuthSeqOutbox{}, nil
	}

	nowText := mysqlTimestamp(now)
	leaseText := mysqlTimestamp(leaseUntil)
	query := fmt.Sprintf(`
UPDATE dialog_auth_seq_outbox
SET status = ?, lease_owner = ?, lease_until = ?
WHERE outbox_id IN (
  SELECT outbox_id FROM (
    SELECT outbox_id
    FROM dialog_auth_seq_outbox
    WHERE ((status IN (?, ?) AND next_retry_at <= ?) OR (status = ? AND lease_until <= ?))
    ORDER BY next_retry_at ASC, outbox_id ASC
    LIMIT ?
  ) AS claimable
)`)
	if _, err := db.Exec(ctx, query, OutboxStatusPublishing, owner, leaseText, OutboxStatusPending, OutboxStatusFailedRetryable, nowText, OutboxStatusPublishing, nowText, limit); err != nil {
		return nil, storageError("claim dialog auth seq outbox", err)
	}

	var rows []model.DialogAuthSeqOutbox
	selectQuery := fmt.Sprintf("SELECT %s FROM dialog_auth_seq_outbox WHERE status = ? AND lease_owner = ? AND lease_until = ? ORDER BY outbox_id ASC LIMIT ?", dialogAuthSeqOutboxSelectRows)
	if err := db.QueryRowsPartial(ctx, &rows, selectQuery, OutboxStatusPublishing, owner, leaseText, limit); err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return []model.DialogAuthSeqOutbox{}, nil
		}
		return nil, storageError("select claimed dialog auth seq outbox", err)
	}
	return rows, nil
}

func (r *Repository) MarkDialogAuthSeqOutboxPublished(ctx context.Context, outboxID int64, seq int64, date int32) error {
	db, err := r.requireDB()
	if err != nil {
		return err
	}
	query := `UPDATE dialog_auth_seq_outbox
SET status = ?, published_seq = ?, published_date = ?, lease_owner = '', lease_until = ?, last_error_kind = '', last_error_message = ''
WHERE outbox_id = ?`
	if _, err := db.Exec(ctx, query, OutboxStatusPublished, seq, date, mysqlZeroTime(), outboxID); err != nil {
		return storageError("mark dialog auth seq outbox published", err)
	}
	return nil
}

func (r *Repository) MarkDialogAuthSeqOutboxRetryable(ctx context.Context, outboxID int64, retry OutboxRetryState) error {
	db, err := r.requireDB()
	if err != nil {
		return err
	}
	query := `UPDATE dialog_auth_seq_outbox
SET status = ?, attempt_count = ?, next_retry_at = ?, lease_owner = '', lease_until = ?, last_error_kind = ?, last_error_message = ?
WHERE outbox_id = ?`
	if _, err := db.Exec(ctx, query, OutboxStatusFailedRetryable, retry.AttemptCount, mysqlTimestamp(retry.NextRetryAt), mysqlZeroTime(), retry.LastErrorKind, truncateOutboxError(retry.LastErrorMessage), outboxID); err != nil {
		return storageError("mark dialog auth seq outbox retryable", err)
	}
	return nil
}

func (r *Repository) MarkDialogAuthSeqOutboxBlocked(ctx context.Context, outboxID int64, kind string, message string) error {
	db, err := r.requireDB()
	if err != nil {
		return err
	}
	query := `UPDATE dialog_auth_seq_outbox
SET status = ?, lease_owner = '', lease_until = ?, last_error_kind = ?, last_error_message = ?
WHERE outbox_id = ?`
	if _, err := db.Exec(ctx, query, OutboxStatusBlocked, mysqlZeroTime(), kind, truncateOutboxError(message), outboxID); err != nil {
		return storageError("mark dialog auth seq outbox blocked", err)
	}
	return nil
}

func (r *Repository) ResetDialogAuthSeqOutboxBlocked(ctx context.Context, ids []int64) error {
	db, err := r.requireDB()
	if err != nil {
		return err
	}
	for _, id := range ids {
		query := `UPDATE dialog_auth_seq_outbox
SET status = ?, attempt_count = 0, next_retry_at = ?, lease_owner = '', lease_until = ?, last_error_kind = '', last_error_message = ''
WHERE status = ? AND outbox_id = ?`
		if _, err := db.Exec(ctx, query, OutboxStatusPending, mysqlTimestamp(time.Now().UTC()), mysqlZeroTime(), OutboxStatusBlocked, id); err != nil {
			return storageError("reset dialog auth seq outbox blocked", err)
		}
	}
	return nil
}

func (r *Repository) ClaimDialogPublicUpdateOutbox(ctx context.Context, owner string, now time.Time, leaseUntil time.Time, limit int32) ([]model.DialogPublicUpdateOutbox, error) {
	db, err := r.requireDB()
	if err != nil {
		return nil, err
	}
	if owner == "" || limit <= 0 {
		return []model.DialogPublicUpdateOutbox{}, nil
	}

	nowText := mysqlTimestamp(now)
	leaseText := mysqlTimestamp(leaseUntil)
	query := fmt.Sprintf(`
UPDATE dialog_public_update_outbox
SET status = ?, lease_owner = ?, lease_until = ?
WHERE outbox_id IN (
  SELECT outbox_id FROM (
    SELECT outbox_id
    FROM dialog_public_update_outbox
    WHERE ((status IN (?, ?) AND next_retry_at <= ?) OR (status = ? AND lease_until <= ?))
    ORDER BY next_retry_at ASC, outbox_id ASC
    LIMIT ?
  ) AS claimable
)`)
	if _, err := db.Exec(ctx, query, OutboxStatusPublishing, owner, leaseText, OutboxStatusPending, OutboxStatusFailedRetryable, nowText, OutboxStatusPublishing, nowText, limit); err != nil {
		return nil, storageError("claim dialog public update outbox", err)
	}

	var rows []model.DialogPublicUpdateOutbox
	selectQuery := fmt.Sprintf("SELECT %s FROM dialog_public_update_outbox WHERE status = ? AND lease_owner = ? AND lease_until = ? ORDER BY outbox_id ASC LIMIT ?", dialogPublicUpdateOutboxSelectRows)
	if err := db.QueryRowsPartial(ctx, &rows, selectQuery, OutboxStatusPublishing, owner, leaseText, limit); err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return []model.DialogPublicUpdateOutbox{}, nil
		}
		return nil, storageError("select claimed dialog public update outbox", err)
	}
	return rows, nil
}

func (r *Repository) MarkDialogPublicUpdateOutboxPublishedPTS(ctx context.Context, outboxID int64, pts int64, ptsCount int32) error {
	db, err := r.requireDB()
	if err != nil {
		return err
	}
	query := `UPDATE dialog_public_update_outbox
SET status = ?, published_pts = ?, published_pts_count = ?, lease_owner = '', lease_until = ?, last_error_kind = '', last_error_message = ''
WHERE outbox_id = ?`
	if _, err := db.Exec(ctx, query, OutboxStatusPublished, pts, ptsCount, mysqlZeroTime(), outboxID); err != nil {
		return storageError("mark dialog public update outbox pts published", err)
	}
	return nil
}

func (r *Repository) MarkDialogPublicUpdateOutboxPublishedAuthSeq(ctx context.Context, outboxID int64, seq int64, date int32) error {
	db, err := r.requireDB()
	if err != nil {
		return err
	}
	query := `UPDATE dialog_public_update_outbox
SET status = ?, published_seq = ?, published_date = ?, lease_owner = '', lease_until = ?, last_error_kind = '', last_error_message = ''
WHERE outbox_id = ?`
	if _, err := db.Exec(ctx, query, OutboxStatusPublished, seq, date, mysqlZeroTime(), outboxID); err != nil {
		return storageError("mark dialog public update outbox auth seq published", err)
	}
	return nil
}

func (r *Repository) MarkDialogPublicUpdateOutboxRetryable(ctx context.Context, outboxID int64, retry OutboxRetryState) error {
	db, err := r.requireDB()
	if err != nil {
		return err
	}
	query := `UPDATE dialog_public_update_outbox
SET status = ?, attempt_count = ?, next_retry_at = ?, lease_owner = '', lease_until = ?, last_error_kind = ?, last_error_message = ?
WHERE outbox_id = ?`
	if _, err := db.Exec(ctx, query, OutboxStatusFailedRetryable, retry.AttemptCount, mysqlTimestamp(retry.NextRetryAt), mysqlZeroTime(), retry.LastErrorKind, truncateOutboxError(retry.LastErrorMessage), outboxID); err != nil {
		return storageError("mark dialog public update outbox retryable", err)
	}
	return nil
}

func (r *Repository) MarkDialogPublicUpdateOutboxBlocked(ctx context.Context, outboxID int64, kind string, message string) error {
	db, err := r.requireDB()
	if err != nil {
		return err
	}
	query := `UPDATE dialog_public_update_outbox
SET status = ?, lease_owner = '', lease_until = ?, last_error_kind = ?, last_error_message = ?
WHERE outbox_id = ?`
	if _, err := db.Exec(ctx, query, OutboxStatusBlocked, mysqlZeroTime(), kind, truncateOutboxError(message), outboxID); err != nil {
		return storageError("mark dialog public update outbox blocked", err)
	}
	return nil
}

func (r *Repository) ResetDialogPublicUpdateOutboxBlocked(ctx context.Context, ids []int64) error {
	db, err := r.requireDB()
	if err != nil {
		return err
	}
	for _, id := range ids {
		query := `UPDATE dialog_public_update_outbox
SET status = ?, attempt_count = 0, next_retry_at = ?, lease_owner = '', lease_until = ?, last_error_kind = '', last_error_message = ''
WHERE status = ? AND outbox_id = ?`
		if _, err := db.Exec(ctx, query, OutboxStatusPending, mysqlTimestamp(time.Now().UTC()), mysqlZeroTime(), OutboxStatusBlocked, id); err != nil {
			return storageError("reset dialog public update outbox blocked", err)
		}
	}
	return nil
}

func truncateOutboxError(s string) string {
	const max = 512
	s = strings.TrimSpace(s)
	if len(s) <= max {
		return s
	}
	return s[:max]
}
