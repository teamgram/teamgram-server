/*
 * WARNING! All changes made in this file will be lost!
 *   Created from by 'dalgen'
 *
 * Copyright (c) 2026-present, The Teamgram Authors (https://teamgram.net).
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package model

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/teamgram/marmota/pkg/stores/sqlx"
)

var _ *sql.Result
var _ = fmt.Sprintf
var _ = strings.Join
var _ = errors.Is
var _ *sqlx.DB
var _ *sqlx.Tx

type bizDialogAuthSeqOutboxModel interface {
	Insert(ctx context.Context, data *DialogAuthSeqOutbox) (lastInsertId, rowsAffected int64, err error)
	InsertIgnore(ctx context.Context, data *DialogAuthSeqOutbox) (lastInsertId, rowsAffected int64, err error)
	SelectByUserOperation(ctx context.Context, userId int64, operationId string) (*DialogAuthSeqOutbox, error)
	MarkPublishing(ctx context.Context, status int32, leaseOwner string, leaseUntil time.Time, outboxId int64) (rowsAffected int64, err error)
	MarkPublished(ctx context.Context, status int32, publishedSeq int64, publishedDate int32, leaseUntil time.Time, outboxId int64) (rowsAffected int64, err error)
	MarkRetryable(ctx context.Context, status int32, attemptCount int32, nextRetryAt time.Time, leaseUntil time.Time, lastErrorKind string, lastErrorMessage string, outboxId int64) (rowsAffected int64, err error)
	MarkBlocked(ctx context.Context, status int32, leaseUntil time.Time, lastErrorKind string, lastErrorMessage string, outboxId int64) (rowsAffected int64, err error)
	ResetBlocked(ctx context.Context, status int32, nextRetryAt time.Time, leaseUntil time.Time, oldStatus int32, outboxId int64) (rowsAffected int64, err error)
}

type DialogAuthSeqOutboxTxModel interface {
	Insert(data *DialogAuthSeqOutbox) (lastInsertId, rowsAffected int64, err error)
	InsertIgnore(data *DialogAuthSeqOutbox) (lastInsertId, rowsAffected int64, err error)
	SelectByUserOperation(userId int64, operationId string) (*DialogAuthSeqOutbox, error)
	MarkPublishing(status int32, leaseOwner string, leaseUntil time.Time, outboxId int64) (rowsAffected int64, err error)
	MarkPublished(status int32, publishedSeq int64, publishedDate int32, leaseUntil time.Time, outboxId int64) (rowsAffected int64, err error)
	MarkRetryable(status int32, attemptCount int32, nextRetryAt time.Time, leaseUntil time.Time, lastErrorKind string, lastErrorMessage string, outboxId int64) (rowsAffected int64, err error)
	MarkBlocked(status int32, leaseUntil time.Time, lastErrorKind string, lastErrorMessage string, outboxId int64) (rowsAffected int64, err error)
	ResetBlocked(status int32, nextRetryAt time.Time, leaseUntil time.Time, oldStatus int32, outboxId int64) (rowsAffected int64, err error)
}

type defaultDialogAuthSeqOutboxTxModel struct {
	tx *sqlx.Tx
}

func NewDialogAuthSeqOutboxTxModel(tx *sqlx.Tx) DialogAuthSeqOutboxTxModel {
	return &defaultDialogAuthSeqOutboxTxModel{tx: tx}
}

// Insert
// insert into dialog_auth_seq_outbox(outbox_id, user_id, source_perm_auth_key_id, target_auth_policy, operation_id, event_type, peer_type, peer_id, payload_schema_version, payload, payload_hash, `status`, attempt_count, next_retry_at, lease_owner, lease_until, published_seq, published_date, last_error_kind, last_error_message) values (:outbox_id, :user_id, :source_perm_auth_key_id, :target_auth_policy, :operation_id, :event_type, :peer_type, :peer_id, :payload_schema_version, :payload, :payload_hash, :status, :attempt_count, :next_retry_at, :lease_owner, :lease_until, :published_seq, :published_date, :last_error_kind, :last_error_message)
func (m *defaultDialogAuthSeqOutboxModel) Insert(ctx context.Context, data *DialogAuthSeqOutbox) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into dialog_auth_seq_outbox(outbox_id, user_id, source_perm_auth_key_id, target_auth_policy, operation_id, event_type, peer_type, peer_id, payload_schema_version, payload, payload_hash, `status`, attempt_count, next_retry_at, lease_owner, lease_until, published_seq, published_date, last_error_kind, last_error_message) values (:outbox_id, :user_id, :source_perm_auth_key_id, :target_auth_policy, :operation_id, :event_type, :peer_type, :peer_id, :payload_schema_version, :payload, :payload_hash, :status, :attempt_count, :next_retry_at, :lease_owner, :lease_until, :published_seq, :published_date, :last_error_kind, :last_error_message)"
		r     sql.Result
	)

	r, err = m.db.NamedExec(ctx, query, data)
	if err != nil {
		err = fmt.Errorf("dialog_auth_seq_outbox.Insert named exec: %w", err)
		return
	}

	lastInsertId, err = r.LastInsertId()
	if err != nil {
		err = fmt.Errorf("dialog_auth_seq_outbox.Insert last insert id: %w", err)
		return
	}
	rowsAffected, err = r.RowsAffected()
	if err != nil {
		err = fmt.Errorf("dialog_auth_seq_outbox.Insert rows affected: %w", err)
	}

	return

}

// Insert
// insert into dialog_auth_seq_outbox(outbox_id, user_id, source_perm_auth_key_id, target_auth_policy, operation_id, event_type, peer_type, peer_id, payload_schema_version, payload, payload_hash, `status`, attempt_count, next_retry_at, lease_owner, lease_until, published_seq, published_date, last_error_kind, last_error_message) values (:outbox_id, :user_id, :source_perm_auth_key_id, :target_auth_policy, :operation_id, :event_type, :peer_type, :peer_id, :payload_schema_version, :payload, :payload_hash, :status, :attempt_count, :next_retry_at, :lease_owner, :lease_until, :published_seq, :published_date, :last_error_kind, :last_error_message)
func (m *defaultDialogAuthSeqOutboxTxModel) Insert(data *DialogAuthSeqOutbox) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into dialog_auth_seq_outbox(outbox_id, user_id, source_perm_auth_key_id, target_auth_policy, operation_id, event_type, peer_type, peer_id, payload_schema_version, payload, payload_hash, `status`, attempt_count, next_retry_at, lease_owner, lease_until, published_seq, published_date, last_error_kind, last_error_message) values (:outbox_id, :user_id, :source_perm_auth_key_id, :target_auth_policy, :operation_id, :event_type, :peer_type, :peer_id, :payload_schema_version, :payload, :payload_hash, :status, :attempt_count, :next_retry_at, :lease_owner, :lease_until, :published_seq, :published_date, :last_error_kind, :last_error_message)"
		r     sql.Result
	)

	r, err = m.tx.NamedExec(query, data)
	if err != nil {
		err = fmt.Errorf("dialog_auth_seq_outbox.Insert named exec: %w", err)
		return
	}

	lastInsertId, err = r.LastInsertId()
	if err != nil {
		err = fmt.Errorf("dialog_auth_seq_outbox.Insert last insert id: %w", err)
		return
	}
	rowsAffected, err = r.RowsAffected()
	if err != nil {
		err = fmt.Errorf("dialog_auth_seq_outbox.Insert rows affected: %w", err)
	}

	return
}

// InsertIgnore
// insert ignore into dialog_auth_seq_outbox(outbox_id, user_id, source_perm_auth_key_id, target_auth_policy, operation_id, event_type, peer_type, peer_id, payload_schema_version, payload, payload_hash, `status`, attempt_count, next_retry_at, lease_owner, lease_until, published_seq, published_date, last_error_kind, last_error_message) values (:outbox_id, :user_id, :source_perm_auth_key_id, :target_auth_policy, :operation_id, :event_type, :peer_type, :peer_id, :payload_schema_version, :payload, :payload_hash, :status, :attempt_count, :next_retry_at, :lease_owner, :lease_until, :published_seq, :published_date, :last_error_kind, :last_error_message)
func (m *defaultDialogAuthSeqOutboxModel) InsertIgnore(ctx context.Context, data *DialogAuthSeqOutbox) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert ignore into dialog_auth_seq_outbox(outbox_id, user_id, source_perm_auth_key_id, target_auth_policy, operation_id, event_type, peer_type, peer_id, payload_schema_version, payload, payload_hash, `status`, attempt_count, next_retry_at, lease_owner, lease_until, published_seq, published_date, last_error_kind, last_error_message) values (:outbox_id, :user_id, :source_perm_auth_key_id, :target_auth_policy, :operation_id, :event_type, :peer_type, :peer_id, :payload_schema_version, :payload, :payload_hash, :status, :attempt_count, :next_retry_at, :lease_owner, :lease_until, :published_seq, :published_date, :last_error_kind, :last_error_message)"
		r     sql.Result
	)

	r, err = m.db.NamedExec(ctx, query, data)
	if err != nil {
		err = fmt.Errorf("dialog_auth_seq_outbox.InsertIgnore named exec: %w", err)
		return
	}

	lastInsertId, err = r.LastInsertId()
	if err != nil {
		err = fmt.Errorf("dialog_auth_seq_outbox.InsertIgnore last insert id: %w", err)
		return
	}
	rowsAffected, err = r.RowsAffected()
	if err != nil {
		err = fmt.Errorf("dialog_auth_seq_outbox.InsertIgnore rows affected: %w", err)
	}

	return

}

// InsertIgnore
// insert ignore into dialog_auth_seq_outbox(outbox_id, user_id, source_perm_auth_key_id, target_auth_policy, operation_id, event_type, peer_type, peer_id, payload_schema_version, payload, payload_hash, `status`, attempt_count, next_retry_at, lease_owner, lease_until, published_seq, published_date, last_error_kind, last_error_message) values (:outbox_id, :user_id, :source_perm_auth_key_id, :target_auth_policy, :operation_id, :event_type, :peer_type, :peer_id, :payload_schema_version, :payload, :payload_hash, :status, :attempt_count, :next_retry_at, :lease_owner, :lease_until, :published_seq, :published_date, :last_error_kind, :last_error_message)
func (m *defaultDialogAuthSeqOutboxTxModel) InsertIgnore(data *DialogAuthSeqOutbox) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert ignore into dialog_auth_seq_outbox(outbox_id, user_id, source_perm_auth_key_id, target_auth_policy, operation_id, event_type, peer_type, peer_id, payload_schema_version, payload, payload_hash, `status`, attempt_count, next_retry_at, lease_owner, lease_until, published_seq, published_date, last_error_kind, last_error_message) values (:outbox_id, :user_id, :source_perm_auth_key_id, :target_auth_policy, :operation_id, :event_type, :peer_type, :peer_id, :payload_schema_version, :payload, :payload_hash, :status, :attempt_count, :next_retry_at, :lease_owner, :lease_until, :published_seq, :published_date, :last_error_kind, :last_error_message)"
		r     sql.Result
	)

	r, err = m.tx.NamedExec(query, data)
	if err != nil {
		err = fmt.Errorf("dialog_auth_seq_outbox.InsertIgnore named exec: %w", err)
		return
	}

	lastInsertId, err = r.LastInsertId()
	if err != nil {
		err = fmt.Errorf("dialog_auth_seq_outbox.InsertIgnore last insert id: %w", err)
		return
	}
	rowsAffected, err = r.RowsAffected()
	if err != nil {
		err = fmt.Errorf("dialog_auth_seq_outbox.InsertIgnore rows affected: %w", err)
	}

	return
}

// SelectByUserOperation
// select outbox_id, user_id, source_perm_auth_key_id, target_auth_policy, operation_id, event_type, peer_type, peer_id, payload_schema_version, payload, payload_hash, `status`, attempt_count, next_retry_at, lease_owner, lease_until, published_seq, published_date, last_error_kind, last_error_message from dialog_auth_seq_outbox where user_id = :user_id and operation_id = :operation_id limit 1
func (m *defaultDialogAuthSeqOutboxModel) SelectByUserOperation(ctx context.Context, userId int64, operationId string) (rValue *DialogAuthSeqOutbox, err error) {

	var (
		query = "select outbox_id, user_id, source_perm_auth_key_id, target_auth_policy, operation_id, event_type, peer_type, peer_id, payload_schema_version, payload, payload_hash, `status`, attempt_count, next_retry_at, lease_owner, lease_until, published_seq, published_date, last_error_kind, last_error_message from dialog_auth_seq_outbox where user_id = ? and operation_id = ? limit 1"
		do    = &DialogAuthSeqOutbox{}
	)
	err = m.db.QueryRowPartial(ctx, do, query, userId, operationId)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return nil, &NotFoundError{
				Resource: "dialog_auth_seq_outbox",
				Key:      fmt.Sprintf("user_id=%v,operation_id=%v", userId, operationId),
				Cause:    err,
			}
		}
		err = fmt.Errorf("dialog_auth_seq_outbox.SelectByUserOperation: %w", err)
		return
	} else {
		rValue = do
	}

	return
}

// SelectByUserOperation
// select outbox_id, user_id, source_perm_auth_key_id, target_auth_policy, operation_id, event_type, peer_type, peer_id, payload_schema_version, payload, payload_hash, `status`, attempt_count, next_retry_at, lease_owner, lease_until, published_seq, published_date, last_error_kind, last_error_message from dialog_auth_seq_outbox where user_id = :user_id and operation_id = :operation_id limit 1
func (m *defaultDialogAuthSeqOutboxTxModel) SelectByUserOperation(userId int64, operationId string) (rValue *DialogAuthSeqOutbox, err error) {
	var (
		query = "select outbox_id, user_id, source_perm_auth_key_id, target_auth_policy, operation_id, event_type, peer_type, peer_id, payload_schema_version, payload, payload_hash, `status`, attempt_count, next_retry_at, lease_owner, lease_until, published_seq, published_date, last_error_kind, last_error_message from dialog_auth_seq_outbox where user_id = ? and operation_id = ? limit 1"
		do    = &DialogAuthSeqOutbox{}
	)
	err = m.tx.QueryRowPartial(do, query, userId, operationId)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return nil, &NotFoundError{
				Resource: "dialog_auth_seq_outbox",
				Key:      fmt.Sprintf("user_id=%v,operation_id=%v", userId, operationId),
				Cause:    err,
			}
		}
		err = fmt.Errorf("dialog_auth_seq_outbox.SelectByUserOperation: %w", err)
		return
	}
	rValue = do

	return
}

// MarkPublishing
// update dialog_auth_seq_outbox set `status` = :status, lease_owner = :lease_owner, lease_until = :lease_until where outbox_id = :outbox_id
func (m *defaultDialogAuthSeqOutboxModel) MarkPublishing(ctx context.Context, status int32, leaseOwner string, leaseUntil time.Time, outboxId int64) (rowsAffected int64, err error) {

	var (
		query   = "update dialog_auth_seq_outbox set `status` = ?, lease_owner = ?, lease_until = ? where outbox_id = ?"
		rResult sql.Result
	)

	rResult, err = m.db.Exec(ctx, query, status, leaseOwner, leaseUntil, outboxId)

	if err != nil {
		err = fmt.Errorf("dialog_auth_seq_outbox.MarkPublishing exec: %w", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		err = fmt.Errorf("dialog_auth_seq_outbox.MarkPublishing rows affected: %w", err)
		return
	}

	return
}

// MarkPublishing
// update dialog_auth_seq_outbox set `status` = :status, lease_owner = :lease_owner, lease_until = :lease_until where outbox_id = :outbox_id
func (m *defaultDialogAuthSeqOutboxTxModel) MarkPublishing(status int32, leaseOwner string, leaseUntil time.Time, outboxId int64) (rowsAffected int64, err error) {
	var (
		query   = "update dialog_auth_seq_outbox set `status` = ?, lease_owner = ?, lease_until = ? where outbox_id = ?"
		rResult sql.Result
	)
	rResult, err = m.tx.Exec(query, status, leaseOwner, leaseUntil, outboxId)

	if err != nil {
		err = fmt.Errorf("dialog_auth_seq_outbox.MarkPublishing exec: %w", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		err = fmt.Errorf("dialog_auth_seq_outbox.MarkPublishing rows affected: %w", err)
		return
	}

	return
}

// MarkPublished
// update dialog_auth_seq_outbox set `status` = :status, published_seq = :published_seq, published_date = :published_date, lease_owner = ”, lease_until = :lease_until, last_error_kind = ”, last_error_message = ” where outbox_id = :outbox_id
func (m *defaultDialogAuthSeqOutboxModel) MarkPublished(ctx context.Context, status int32, publishedSeq int64, publishedDate int32, leaseUntil time.Time, outboxId int64) (rowsAffected int64, err error) {

	var (
		query   = "update dialog_auth_seq_outbox set `status` = ?, published_seq = ?, published_date = ?, lease_owner = '', lease_until = ?, last_error_kind = '', last_error_message = '' where outbox_id = ?"
		rResult sql.Result
	)

	rResult, err = m.db.Exec(ctx, query, status, publishedSeq, publishedDate, leaseUntil, outboxId)

	if err != nil {
		err = fmt.Errorf("dialog_auth_seq_outbox.MarkPublished exec: %w", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		err = fmt.Errorf("dialog_auth_seq_outbox.MarkPublished rows affected: %w", err)
		return
	}

	return
}

// MarkPublished
// update dialog_auth_seq_outbox set `status` = :status, published_seq = :published_seq, published_date = :published_date, lease_owner = ”, lease_until = :lease_until, last_error_kind = ”, last_error_message = ” where outbox_id = :outbox_id
func (m *defaultDialogAuthSeqOutboxTxModel) MarkPublished(status int32, publishedSeq int64, publishedDate int32, leaseUntil time.Time, outboxId int64) (rowsAffected int64, err error) {
	var (
		query   = "update dialog_auth_seq_outbox set `status` = ?, published_seq = ?, published_date = ?, lease_owner = '', lease_until = ?, last_error_kind = '', last_error_message = '' where outbox_id = ?"
		rResult sql.Result
	)
	rResult, err = m.tx.Exec(query, status, publishedSeq, publishedDate, leaseUntil, outboxId)

	if err != nil {
		err = fmt.Errorf("dialog_auth_seq_outbox.MarkPublished exec: %w", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		err = fmt.Errorf("dialog_auth_seq_outbox.MarkPublished rows affected: %w", err)
		return
	}

	return
}

// MarkRetryable
// update dialog_auth_seq_outbox set `status` = :status, attempt_count = :attempt_count, next_retry_at = :next_retry_at, lease_owner = ”, lease_until = :lease_until, last_error_kind = :last_error_kind, last_error_message = :last_error_message where outbox_id = :outbox_id
func (m *defaultDialogAuthSeqOutboxModel) MarkRetryable(ctx context.Context, status int32, attemptCount int32, nextRetryAt time.Time, leaseUntil time.Time, lastErrorKind string, lastErrorMessage string, outboxId int64) (rowsAffected int64, err error) {

	var (
		query   = "update dialog_auth_seq_outbox set `status` = ?, attempt_count = ?, next_retry_at = ?, lease_owner = '', lease_until = ?, last_error_kind = ?, last_error_message = ? where outbox_id = ?"
		rResult sql.Result
	)

	rResult, err = m.db.Exec(ctx, query, status, attemptCount, nextRetryAt, leaseUntil, lastErrorKind, lastErrorMessage, outboxId)

	if err != nil {
		err = fmt.Errorf("dialog_auth_seq_outbox.MarkRetryable exec: %w", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		err = fmt.Errorf("dialog_auth_seq_outbox.MarkRetryable rows affected: %w", err)
		return
	}

	return
}

// MarkRetryable
// update dialog_auth_seq_outbox set `status` = :status, attempt_count = :attempt_count, next_retry_at = :next_retry_at, lease_owner = ”, lease_until = :lease_until, last_error_kind = :last_error_kind, last_error_message = :last_error_message where outbox_id = :outbox_id
func (m *defaultDialogAuthSeqOutboxTxModel) MarkRetryable(status int32, attemptCount int32, nextRetryAt time.Time, leaseUntil time.Time, lastErrorKind string, lastErrorMessage string, outboxId int64) (rowsAffected int64, err error) {
	var (
		query   = "update dialog_auth_seq_outbox set `status` = ?, attempt_count = ?, next_retry_at = ?, lease_owner = '', lease_until = ?, last_error_kind = ?, last_error_message = ? where outbox_id = ?"
		rResult sql.Result
	)
	rResult, err = m.tx.Exec(query, status, attemptCount, nextRetryAt, leaseUntil, lastErrorKind, lastErrorMessage, outboxId)

	if err != nil {
		err = fmt.Errorf("dialog_auth_seq_outbox.MarkRetryable exec: %w", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		err = fmt.Errorf("dialog_auth_seq_outbox.MarkRetryable rows affected: %w", err)
		return
	}

	return
}

// MarkBlocked
// update dialog_auth_seq_outbox set `status` = :status, lease_owner = ”, lease_until = :lease_until, last_error_kind = :last_error_kind, last_error_message = :last_error_message where outbox_id = :outbox_id
func (m *defaultDialogAuthSeqOutboxModel) MarkBlocked(ctx context.Context, status int32, leaseUntil time.Time, lastErrorKind string, lastErrorMessage string, outboxId int64) (rowsAffected int64, err error) {

	var (
		query   = "update dialog_auth_seq_outbox set `status` = ?, lease_owner = '', lease_until = ?, last_error_kind = ?, last_error_message = ? where outbox_id = ?"
		rResult sql.Result
	)

	rResult, err = m.db.Exec(ctx, query, status, leaseUntil, lastErrorKind, lastErrorMessage, outboxId)

	if err != nil {
		err = fmt.Errorf("dialog_auth_seq_outbox.MarkBlocked exec: %w", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		err = fmt.Errorf("dialog_auth_seq_outbox.MarkBlocked rows affected: %w", err)
		return
	}

	return
}

// MarkBlocked
// update dialog_auth_seq_outbox set `status` = :status, lease_owner = ”, lease_until = :lease_until, last_error_kind = :last_error_kind, last_error_message = :last_error_message where outbox_id = :outbox_id
func (m *defaultDialogAuthSeqOutboxTxModel) MarkBlocked(status int32, leaseUntil time.Time, lastErrorKind string, lastErrorMessage string, outboxId int64) (rowsAffected int64, err error) {
	var (
		query   = "update dialog_auth_seq_outbox set `status` = ?, lease_owner = '', lease_until = ?, last_error_kind = ?, last_error_message = ? where outbox_id = ?"
		rResult sql.Result
	)
	rResult, err = m.tx.Exec(query, status, leaseUntil, lastErrorKind, lastErrorMessage, outboxId)

	if err != nil {
		err = fmt.Errorf("dialog_auth_seq_outbox.MarkBlocked exec: %w", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		err = fmt.Errorf("dialog_auth_seq_outbox.MarkBlocked rows affected: %w", err)
		return
	}

	return
}

// ResetBlocked
// update dialog_auth_seq_outbox set `status` = :status, attempt_count = 0, next_retry_at = :next_retry_at, lease_owner = ”, lease_until = :lease_until, last_error_kind = ”, last_error_message = ” where `status` = :old_status and outbox_id = :outbox_id
func (m *defaultDialogAuthSeqOutboxModel) ResetBlocked(ctx context.Context, status int32, nextRetryAt time.Time, leaseUntil time.Time, oldStatus int32, outboxId int64) (rowsAffected int64, err error) {

	var (
		query   = "update dialog_auth_seq_outbox set `status` = ?, attempt_count = 0, next_retry_at = ?, lease_owner = '', lease_until = ?, last_error_kind = '', last_error_message = '' where `status` = ? and outbox_id = ?"
		rResult sql.Result
	)

	rResult, err = m.db.Exec(ctx, query, status, nextRetryAt, leaseUntil, oldStatus, outboxId)

	if err != nil {
		err = fmt.Errorf("dialog_auth_seq_outbox.ResetBlocked exec: %w", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		err = fmt.Errorf("dialog_auth_seq_outbox.ResetBlocked rows affected: %w", err)
		return
	}

	return
}

// ResetBlocked
// update dialog_auth_seq_outbox set `status` = :status, attempt_count = 0, next_retry_at = :next_retry_at, lease_owner = ”, lease_until = :lease_until, last_error_kind = ”, last_error_message = ” where `status` = :old_status and outbox_id = :outbox_id
func (m *defaultDialogAuthSeqOutboxTxModel) ResetBlocked(status int32, nextRetryAt time.Time, leaseUntil time.Time, oldStatus int32, outboxId int64) (rowsAffected int64, err error) {
	var (
		query   = "update dialog_auth_seq_outbox set `status` = ?, attempt_count = 0, next_retry_at = ?, lease_owner = '', lease_until = ?, last_error_kind = '', last_error_message = '' where `status` = ? and outbox_id = ?"
		rResult sql.Result
	)
	rResult, err = m.tx.Exec(query, status, nextRetryAt, leaseUntil, oldStatus, outboxId)

	if err != nil {
		err = fmt.Errorf("dialog_auth_seq_outbox.ResetBlocked exec: %w", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		err = fmt.Errorf("dialog_auth_seq_outbox.ResetBlocked rows affected: %w", err)
		return
	}

	return
}
