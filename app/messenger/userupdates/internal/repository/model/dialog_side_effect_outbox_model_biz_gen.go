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

	"github.com/teamgram/marmota/pkg/stores/sqlx"
)

var _ *sql.Result
var _ = fmt.Sprintf
var _ = strings.Join
var _ = errors.Is
var _ *sqlx.DB
var _ *sqlx.Tx

type bizDialogSideEffectOutboxModel interface {
	Insert(ctx context.Context, data *DialogSideEffectOutbox) (lastInsertId, rowsAffected int64, err error)
	SelectOne(ctx context.Context, sideEffectId int64) (*DialogSideEffectOutbox, error)
	SelectPendingForUpdate(ctx context.Context, pendingStatus int32, failedRetryableStatus int32, now string, publishingStatus int32, limit int32) ([]DialogSideEffectOutbox, error)
	SelectPendingForUpdateWithCB(ctx context.Context, pendingStatus int32, failedRetryableStatus int32, now string, publishingStatus int32, limit int32, cb func(sz, i int, v *DialogSideEffectOutbox)) ([]DialogSideEffectOutbox, error)
	MarkBlockedIfOld(ctx context.Context, status int32, leaseUntil string, lastErrorCode string, sideEffectId int64, blockedBefore string) (rowsAffected int64, err error)
	SelectExistingSideEffect(ctx context.Context, kind string, sourceOperationId string) (*DialogSideEffectOutbox, error)
	MarkPublishing(ctx context.Context, status int32, leaseOwner string, leaseUntil string, sideEffectId int64) (rowsAffected int64, err error)
	MarkCompleted(ctx context.Context, status int32, leaseUntil string, sideEffectId int64) (rowsAffected int64, err error)
	MarkRetryableFailure(ctx context.Context, status int32, nextRetryAt string, leaseUntil string, lastErrorCode string, sideEffectId int64) (rowsAffected int64, err error)
	MarkBlocked(ctx context.Context, status int32, leaseUntil string, lastErrorCode string, sideEffectId int64) (rowsAffected int64, err error)
	SelectBySourceOperationKind(ctx context.Context, sourceOperationId string, kind string) ([]DialogSideEffectOutbox, error)
	SelectBySourceOperationKindWithCB(ctx context.Context, sourceOperationId string, kind string, cb func(sz, i int, v *DialogSideEffectOutbox)) ([]DialogSideEffectOutbox, error)
}

type DialogSideEffectOutboxTxModel interface {
	Insert(data *DialogSideEffectOutbox) (lastInsertId, rowsAffected int64, err error)
	SelectOne(sideEffectId int64) (*DialogSideEffectOutbox, error)
	SelectPendingForUpdate(pendingStatus int32, failedRetryableStatus int32, now string, publishingStatus int32, limit int32) ([]DialogSideEffectOutbox, error)
	MarkBlockedIfOld(status int32, leaseUntil string, lastErrorCode string, sideEffectId int64, blockedBefore string) (rowsAffected int64, err error)
	SelectExistingSideEffect(kind string, sourceOperationId string) (*DialogSideEffectOutbox, error)
	MarkPublishing(status int32, leaseOwner string, leaseUntil string, sideEffectId int64) (rowsAffected int64, err error)
	MarkCompleted(status int32, leaseUntil string, sideEffectId int64) (rowsAffected int64, err error)
	MarkRetryableFailure(status int32, nextRetryAt string, leaseUntil string, lastErrorCode string, sideEffectId int64) (rowsAffected int64, err error)
	MarkBlocked(status int32, leaseUntil string, lastErrorCode string, sideEffectId int64) (rowsAffected int64, err error)
	SelectBySourceOperationKind(sourceOperationId string, kind string) ([]DialogSideEffectOutbox, error)
}

type defaultDialogSideEffectOutboxTxModel struct {
	tx *sqlx.Tx
}

func NewDialogSideEffectOutboxTxModel(tx *sqlx.Tx) DialogSideEffectOutboxTxModel {
	return &defaultDialogSideEffectOutboxTxModel{tx: tx}
}

// Insert
// insert ignore into dialog_side_effect_outbox(side_effect_id, kind, user_id, peer_type, peer_id, source_perm_auth_key_id, source_operation_id, source_message_date, source_peer_seq, source_canonical_message_id, clear_before_date, payload_schema_version, payload, payload_hash, `status`, attempt_count, next_retry_at, lease_owner, lease_until, last_error_code) values (:side_effect_id, :kind, :user_id, :peer_type, :peer_id, :source_perm_auth_key_id, :source_operation_id, :source_message_date, :source_peer_seq, :source_canonical_message_id, :clear_before_date, :payload_schema_version, :payload, :payload_hash, :status, :attempt_count, :next_retry_at, :lease_owner, :lease_until, :last_error_code)
func (m *defaultDialogSideEffectOutboxModel) Insert(ctx context.Context, data *DialogSideEffectOutbox) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert ignore into dialog_side_effect_outbox(side_effect_id, kind, user_id, peer_type, peer_id, source_perm_auth_key_id, source_operation_id, source_message_date, source_peer_seq, source_canonical_message_id, clear_before_date, payload_schema_version, payload, payload_hash, `status`, attempt_count, next_retry_at, lease_owner, lease_until, last_error_code) values (:side_effect_id, :kind, :user_id, :peer_type, :peer_id, :source_perm_auth_key_id, :source_operation_id, :source_message_date, :source_peer_seq, :source_canonical_message_id, :clear_before_date, :payload_schema_version, :payload, :payload_hash, :status, :attempt_count, :next_retry_at, :lease_owner, :lease_until, :last_error_code)"
		r     sql.Result
	)

	r, err = m.db.NamedExec(ctx, query, data)
	if err != nil {
		err = fmt.Errorf("dialog_side_effect_outbox.Insert named exec: %w", err)
		return
	}

	lastInsertId, err = r.LastInsertId()
	if err != nil {
		err = fmt.Errorf("dialog_side_effect_outbox.Insert last insert id: %w", err)
		return
	}
	rowsAffected, err = r.RowsAffected()
	if err != nil {
		err = fmt.Errorf("dialog_side_effect_outbox.Insert rows affected: %w", err)
	}

	return

}

// Insert
// insert ignore into dialog_side_effect_outbox(side_effect_id, kind, user_id, peer_type, peer_id, source_perm_auth_key_id, source_operation_id, source_message_date, source_peer_seq, source_canonical_message_id, clear_before_date, payload_schema_version, payload, payload_hash, `status`, attempt_count, next_retry_at, lease_owner, lease_until, last_error_code) values (:side_effect_id, :kind, :user_id, :peer_type, :peer_id, :source_perm_auth_key_id, :source_operation_id, :source_message_date, :source_peer_seq, :source_canonical_message_id, :clear_before_date, :payload_schema_version, :payload, :payload_hash, :status, :attempt_count, :next_retry_at, :lease_owner, :lease_until, :last_error_code)
func (m *defaultDialogSideEffectOutboxTxModel) Insert(data *DialogSideEffectOutbox) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert ignore into dialog_side_effect_outbox(side_effect_id, kind, user_id, peer_type, peer_id, source_perm_auth_key_id, source_operation_id, source_message_date, source_peer_seq, source_canonical_message_id, clear_before_date, payload_schema_version, payload, payload_hash, `status`, attempt_count, next_retry_at, lease_owner, lease_until, last_error_code) values (:side_effect_id, :kind, :user_id, :peer_type, :peer_id, :source_perm_auth_key_id, :source_operation_id, :source_message_date, :source_peer_seq, :source_canonical_message_id, :clear_before_date, :payload_schema_version, :payload, :payload_hash, :status, :attempt_count, :next_retry_at, :lease_owner, :lease_until, :last_error_code)"
		r     sql.Result
	)

	r, err = m.tx.NamedExec(query, data)
	if err != nil {
		err = fmt.Errorf("dialog_side_effect_outbox.Insert named exec: %w", err)
		return
	}

	lastInsertId, err = r.LastInsertId()
	if err != nil {
		err = fmt.Errorf("dialog_side_effect_outbox.Insert last insert id: %w", err)
		return
	}
	rowsAffected, err = r.RowsAffected()
	if err != nil {
		err = fmt.Errorf("dialog_side_effect_outbox.Insert rows affected: %w", err)
	}

	return
}

// SelectOne
// select side_effect_id, kind, user_id, peer_type, peer_id, source_perm_auth_key_id, source_operation_id, source_message_date, source_peer_seq, source_canonical_message_id, clear_before_date, payload_schema_version, payload, payload_hash, `status`, attempt_count, next_retry_at, lease_owner, lease_until, last_error_code from dialog_side_effect_outbox where side_effect_id = :side_effect_id limit 1
func (m *defaultDialogSideEffectOutboxModel) SelectOne(ctx context.Context, sideEffectId int64) (rValue *DialogSideEffectOutbox, err error) {

	var (
		query = "select side_effect_id, kind, user_id, peer_type, peer_id, source_perm_auth_key_id, source_operation_id, source_message_date, source_peer_seq, source_canonical_message_id, clear_before_date, payload_schema_version, payload, payload_hash, `status`, attempt_count, next_retry_at, lease_owner, lease_until, last_error_code from dialog_side_effect_outbox where side_effect_id = ? limit 1"
		do    = &DialogSideEffectOutbox{}
	)
	err = m.db.QueryRowPartial(ctx, do, query, sideEffectId)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return nil, &NotFoundError{
				Resource: "dialog_side_effect_outbox",
				Key:      fmt.Sprintf("side_effect_id=%v", sideEffectId),
				Cause:    err,
			}
		}
		err = fmt.Errorf("dialog_side_effect_outbox.SelectOne: %w", err)
		return
	} else {
		rValue = do
	}

	return
}

// SelectOne
// select side_effect_id, kind, user_id, peer_type, peer_id, source_perm_auth_key_id, source_operation_id, source_message_date, source_peer_seq, source_canonical_message_id, clear_before_date, payload_schema_version, payload, payload_hash, `status`, attempt_count, next_retry_at, lease_owner, lease_until, last_error_code from dialog_side_effect_outbox where side_effect_id = :side_effect_id limit 1
func (m *defaultDialogSideEffectOutboxTxModel) SelectOne(sideEffectId int64) (rValue *DialogSideEffectOutbox, err error) {
	var (
		query = "select side_effect_id, kind, user_id, peer_type, peer_id, source_perm_auth_key_id, source_operation_id, source_message_date, source_peer_seq, source_canonical_message_id, clear_before_date, payload_schema_version, payload, payload_hash, `status`, attempt_count, next_retry_at, lease_owner, lease_until, last_error_code from dialog_side_effect_outbox where side_effect_id = ? limit 1"
		do    = &DialogSideEffectOutbox{}
	)
	err = m.tx.QueryRowPartial(do, query, sideEffectId)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return nil, &NotFoundError{
				Resource: "dialog_side_effect_outbox",
				Key:      fmt.Sprintf("side_effect_id=%v", sideEffectId),
				Cause:    err,
			}
		}
		err = fmt.Errorf("dialog_side_effect_outbox.SelectOne: %w", err)
		return
	}
	rValue = do

	return
}

// SelectPendingForUpdate
// select side_effect_id, kind, user_id, peer_type, peer_id, source_perm_auth_key_id, source_operation_id, source_message_date, source_peer_seq, source_canonical_message_id, clear_before_date, payload_schema_version, payload, payload_hash, `status`, attempt_count, next_retry_at, lease_owner, lease_until, last_error_code from dialog_side_effect_outbox where ((`status` in (:pending_status, :failed_retryable_status) and next_retry_at <= :now) or (`status` = :publishing_status and lease_until <= :now)) order by next_retry_at asc, side_effect_id asc limit :limit for update
func (m *defaultDialogSideEffectOutboxModel) SelectPendingForUpdate(ctx context.Context, pendingStatus int32, failedRetryableStatus int32, now string, publishingStatus int32, limit int32) (rList []DialogSideEffectOutbox, err error) {
	var (
		query  = "select side_effect_id, kind, user_id, peer_type, peer_id, source_perm_auth_key_id, source_operation_id, source_message_date, source_peer_seq, source_canonical_message_id, clear_before_date, payload_schema_version, payload, payload_hash, `status`, attempt_count, next_retry_at, lease_owner, lease_until, last_error_code from dialog_side_effect_outbox where ((`status` in (?, ?) and next_retry_at <= ?) or (`status` = ? and lease_until <= ?)) order by next_retry_at asc, side_effect_id asc limit ? for update"
		values []DialogSideEffectOutbox
	)
	err = m.db.QueryRowsPartial(ctx, &values, query, pendingStatus, failedRetryableStatus, now, publishingStatus, now, limit)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			rList = []DialogSideEffectOutbox{}
			err = nil
			return
		}
		err = fmt.Errorf("dialog_side_effect_outbox.SelectPendingForUpdate: %w", err)
		return
	}

	rList = values

	return
}

// SelectPendingForUpdate
// select side_effect_id, kind, user_id, peer_type, peer_id, source_perm_auth_key_id, source_operation_id, source_message_date, source_peer_seq, source_canonical_message_id, clear_before_date, payload_schema_version, payload, payload_hash, `status`, attempt_count, next_retry_at, lease_owner, lease_until, last_error_code from dialog_side_effect_outbox where ((`status` in (:pending_status, :failed_retryable_status) and next_retry_at <= :now) or (`status` = :publishing_status and lease_until <= :now)) order by next_retry_at asc, side_effect_id asc limit :limit for update
func (m *defaultDialogSideEffectOutboxTxModel) SelectPendingForUpdate(pendingStatus int32, failedRetryableStatus int32, now string, publishingStatus int32, limit int32) (rList []DialogSideEffectOutbox, err error) {
	var (
		query  = "select side_effect_id, kind, user_id, peer_type, peer_id, source_perm_auth_key_id, source_operation_id, source_message_date, source_peer_seq, source_canonical_message_id, clear_before_date, payload_schema_version, payload, payload_hash, `status`, attempt_count, next_retry_at, lease_owner, lease_until, last_error_code from dialog_side_effect_outbox where ((`status` in (?, ?) and next_retry_at <= ?) or (`status` = ? and lease_until <= ?)) order by next_retry_at asc, side_effect_id asc limit ? for update"
		values []DialogSideEffectOutbox
	)
	err = m.tx.QueryRowsPartial(&values, query, pendingStatus, failedRetryableStatus, now, publishingStatus, now, limit)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			rList = []DialogSideEffectOutbox{}
			err = nil
			return
		}
		err = fmt.Errorf("dialog_side_effect_outbox.SelectPendingForUpdate: %w", err)
		return
	}

	rList = values

	return
}

// SelectPendingForUpdateWithCB
// select side_effect_id, kind, user_id, peer_type, peer_id, source_perm_auth_key_id, source_operation_id, source_message_date, source_peer_seq, source_canonical_message_id, clear_before_date, payload_schema_version, payload, payload_hash, `status`, attempt_count, next_retry_at, lease_owner, lease_until, last_error_code from dialog_side_effect_outbox where ((`status` in (:pending_status, :failed_retryable_status) and next_retry_at <= :now) or (`status` = :publishing_status and lease_until <= :now)) order by next_retry_at asc, side_effect_id asc limit :limit for update
func (m *defaultDialogSideEffectOutboxModel) SelectPendingForUpdateWithCB(ctx context.Context, pendingStatus int32, failedRetryableStatus int32, now string, publishingStatus int32, limit int32, cb func(sz, i int, v *DialogSideEffectOutbox)) (rList []DialogSideEffectOutbox, err error) {
	var (
		query  = "select side_effect_id, kind, user_id, peer_type, peer_id, source_perm_auth_key_id, source_operation_id, source_message_date, source_peer_seq, source_canonical_message_id, clear_before_date, payload_schema_version, payload, payload_hash, `status`, attempt_count, next_retry_at, lease_owner, lease_until, last_error_code from dialog_side_effect_outbox where ((`status` in (?, ?) and next_retry_at <= ?) or (`status` = ? and lease_until <= ?)) order by next_retry_at asc, side_effect_id asc limit ? for update"
		values []DialogSideEffectOutbox
	)
	err = m.db.QueryRowsPartial(ctx, &values, query, pendingStatus, failedRetryableStatus, now, publishingStatus, now, limit)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			rList = []DialogSideEffectOutbox{}
			err = nil
			return
		}
		err = fmt.Errorf("dialog_side_effect_outbox.SelectPendingForUpdateWithCB: %w", err)
		return
	}

	rList = values

	if cb != nil {
		sz := len(rList)
		for i := 0; i < sz; i++ {
			cb(sz, i, &rList[i])
		}
	}

	return
}

// MarkBlockedIfOld
// update dialog_side_effect_outbox set `status` = :status, lease_owner = ”, lease_until = :lease_until, last_error_code = :last_error_code where side_effect_id = :side_effect_id and created_at <= :blocked_before
func (m *defaultDialogSideEffectOutboxModel) MarkBlockedIfOld(ctx context.Context, status int32, leaseUntil string, lastErrorCode string, sideEffectId int64, blockedBefore string) (rowsAffected int64, err error) {

	var (
		query   = "update dialog_side_effect_outbox set `status` = ?, lease_owner = '', lease_until = ?, last_error_code = ? where side_effect_id = ? and created_at <= ?"
		rResult sql.Result
	)

	rResult, err = m.db.Exec(ctx, query, status, leaseUntil, lastErrorCode, sideEffectId, blockedBefore)

	if err != nil {
		err = fmt.Errorf("dialog_side_effect_outbox.MarkBlockedIfOld exec: %w", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		err = fmt.Errorf("dialog_side_effect_outbox.MarkBlockedIfOld rows affected: %w", err)
		return
	}

	return
}

// MarkBlockedIfOld
// update dialog_side_effect_outbox set `status` = :status, lease_owner = ”, lease_until = :lease_until, last_error_code = :last_error_code where side_effect_id = :side_effect_id and created_at <= :blocked_before
func (m *defaultDialogSideEffectOutboxTxModel) MarkBlockedIfOld(status int32, leaseUntil string, lastErrorCode string, sideEffectId int64, blockedBefore string) (rowsAffected int64, err error) {
	var (
		query   = "update dialog_side_effect_outbox set `status` = ?, lease_owner = '', lease_until = ?, last_error_code = ? where side_effect_id = ? and created_at <= ?"
		rResult sql.Result
	)
	rResult, err = m.tx.Exec(query, status, leaseUntil, lastErrorCode, sideEffectId, blockedBefore)

	if err != nil {
		err = fmt.Errorf("dialog_side_effect_outbox.MarkBlockedIfOld exec: %w", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		err = fmt.Errorf("dialog_side_effect_outbox.MarkBlockedIfOld rows affected: %w", err)
		return
	}

	return
}

// SelectExistingSideEffect
// select side_effect_id, kind, user_id, peer_type, peer_id, source_perm_auth_key_id, source_operation_id, source_message_date, source_peer_seq, source_canonical_message_id, clear_before_date, payload_schema_version, payload, payload_hash, `status`, attempt_count, next_retry_at, lease_owner, lease_until, last_error_code from dialog_side_effect_outbox where kind = :kind and source_operation_id = :source_operation_id limit 1
func (m *defaultDialogSideEffectOutboxModel) SelectExistingSideEffect(ctx context.Context, kind string, sourceOperationId string) (rValue *DialogSideEffectOutbox, err error) {

	var (
		query = "select side_effect_id, kind, user_id, peer_type, peer_id, source_perm_auth_key_id, source_operation_id, source_message_date, source_peer_seq, source_canonical_message_id, clear_before_date, payload_schema_version, payload, payload_hash, `status`, attempt_count, next_retry_at, lease_owner, lease_until, last_error_code from dialog_side_effect_outbox where kind = ? and source_operation_id = ? limit 1"
		do    = &DialogSideEffectOutbox{}
	)
	err = m.db.QueryRowPartial(ctx, do, query, kind, sourceOperationId)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return nil, &NotFoundError{
				Resource: "dialog_side_effect_outbox",
				Key:      fmt.Sprintf("kind=%v,source_operation_id=%v", kind, sourceOperationId),
				Cause:    err,
			}
		}
		err = fmt.Errorf("dialog_side_effect_outbox.SelectExistingSideEffect: %w", err)
		return
	} else {
		rValue = do
	}

	return
}

// SelectExistingSideEffect
// select side_effect_id, kind, user_id, peer_type, peer_id, source_perm_auth_key_id, source_operation_id, source_message_date, source_peer_seq, source_canonical_message_id, clear_before_date, payload_schema_version, payload, payload_hash, `status`, attempt_count, next_retry_at, lease_owner, lease_until, last_error_code from dialog_side_effect_outbox where kind = :kind and source_operation_id = :source_operation_id limit 1
func (m *defaultDialogSideEffectOutboxTxModel) SelectExistingSideEffect(kind string, sourceOperationId string) (rValue *DialogSideEffectOutbox, err error) {
	var (
		query = "select side_effect_id, kind, user_id, peer_type, peer_id, source_perm_auth_key_id, source_operation_id, source_message_date, source_peer_seq, source_canonical_message_id, clear_before_date, payload_schema_version, payload, payload_hash, `status`, attempt_count, next_retry_at, lease_owner, lease_until, last_error_code from dialog_side_effect_outbox where kind = ? and source_operation_id = ? limit 1"
		do    = &DialogSideEffectOutbox{}
	)
	err = m.tx.QueryRowPartial(do, query, kind, sourceOperationId)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return nil, &NotFoundError{
				Resource: "dialog_side_effect_outbox",
				Key:      fmt.Sprintf("kind=%v,source_operation_id=%v", kind, sourceOperationId),
				Cause:    err,
			}
		}
		err = fmt.Errorf("dialog_side_effect_outbox.SelectExistingSideEffect: %w", err)
		return
	}
	rValue = do

	return
}

// MarkPublishing
// update dialog_side_effect_outbox set `status` = :status, attempt_count = attempt_count + 1, lease_owner = :lease_owner, lease_until = :lease_until where side_effect_id = :side_effect_id
func (m *defaultDialogSideEffectOutboxModel) MarkPublishing(ctx context.Context, status int32, leaseOwner string, leaseUntil string, sideEffectId int64) (rowsAffected int64, err error) {

	var (
		query   = "update dialog_side_effect_outbox set `status` = ?, attempt_count = attempt_count + 1, lease_owner = ?, lease_until = ? where side_effect_id = ?"
		rResult sql.Result
	)

	rResult, err = m.db.Exec(ctx, query, status, leaseOwner, leaseUntil, sideEffectId)

	if err != nil {
		err = fmt.Errorf("dialog_side_effect_outbox.MarkPublishing exec: %w", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		err = fmt.Errorf("dialog_side_effect_outbox.MarkPublishing rows affected: %w", err)
		return
	}

	return
}

// MarkPublishing
// update dialog_side_effect_outbox set `status` = :status, attempt_count = attempt_count + 1, lease_owner = :lease_owner, lease_until = :lease_until where side_effect_id = :side_effect_id
func (m *defaultDialogSideEffectOutboxTxModel) MarkPublishing(status int32, leaseOwner string, leaseUntil string, sideEffectId int64) (rowsAffected int64, err error) {
	var (
		query   = "update dialog_side_effect_outbox set `status` = ?, attempt_count = attempt_count + 1, lease_owner = ?, lease_until = ? where side_effect_id = ?"
		rResult sql.Result
	)
	rResult, err = m.tx.Exec(query, status, leaseOwner, leaseUntil, sideEffectId)

	if err != nil {
		err = fmt.Errorf("dialog_side_effect_outbox.MarkPublishing exec: %w", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		err = fmt.Errorf("dialog_side_effect_outbox.MarkPublishing rows affected: %w", err)
		return
	}

	return
}

// MarkCompleted
// update dialog_side_effect_outbox set `status` = :status, lease_owner = ”, lease_until = :lease_until, last_error_code = ” where side_effect_id = :side_effect_id
func (m *defaultDialogSideEffectOutboxModel) MarkCompleted(ctx context.Context, status int32, leaseUntil string, sideEffectId int64) (rowsAffected int64, err error) {

	var (
		query   = "update dialog_side_effect_outbox set `status` = ?, lease_owner = '', lease_until = ?, last_error_code = '' where side_effect_id = ?"
		rResult sql.Result
	)

	rResult, err = m.db.Exec(ctx, query, status, leaseUntil, sideEffectId)

	if err != nil {
		err = fmt.Errorf("dialog_side_effect_outbox.MarkCompleted exec: %w", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		err = fmt.Errorf("dialog_side_effect_outbox.MarkCompleted rows affected: %w", err)
		return
	}

	return
}

// MarkCompleted
// update dialog_side_effect_outbox set `status` = :status, lease_owner = ”, lease_until = :lease_until, last_error_code = ” where side_effect_id = :side_effect_id
func (m *defaultDialogSideEffectOutboxTxModel) MarkCompleted(status int32, leaseUntil string, sideEffectId int64) (rowsAffected int64, err error) {
	var (
		query   = "update dialog_side_effect_outbox set `status` = ?, lease_owner = '', lease_until = ?, last_error_code = '' where side_effect_id = ?"
		rResult sql.Result
	)
	rResult, err = m.tx.Exec(query, status, leaseUntil, sideEffectId)

	if err != nil {
		err = fmt.Errorf("dialog_side_effect_outbox.MarkCompleted exec: %w", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		err = fmt.Errorf("dialog_side_effect_outbox.MarkCompleted rows affected: %w", err)
		return
	}

	return
}

// MarkRetryableFailure
// update dialog_side_effect_outbox set `status` = :status, next_retry_at = :next_retry_at, lease_owner = ”, lease_until = :lease_until, last_error_code = :last_error_code where side_effect_id = :side_effect_id
func (m *defaultDialogSideEffectOutboxModel) MarkRetryableFailure(ctx context.Context, status int32, nextRetryAt string, leaseUntil string, lastErrorCode string, sideEffectId int64) (rowsAffected int64, err error) {

	var (
		query   = "update dialog_side_effect_outbox set `status` = ?, next_retry_at = ?, lease_owner = '', lease_until = ?, last_error_code = ? where side_effect_id = ?"
		rResult sql.Result
	)

	rResult, err = m.db.Exec(ctx, query, status, nextRetryAt, leaseUntil, lastErrorCode, sideEffectId)

	if err != nil {
		err = fmt.Errorf("dialog_side_effect_outbox.MarkRetryableFailure exec: %w", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		err = fmt.Errorf("dialog_side_effect_outbox.MarkRetryableFailure rows affected: %w", err)
		return
	}

	return
}

// MarkRetryableFailure
// update dialog_side_effect_outbox set `status` = :status, next_retry_at = :next_retry_at, lease_owner = ”, lease_until = :lease_until, last_error_code = :last_error_code where side_effect_id = :side_effect_id
func (m *defaultDialogSideEffectOutboxTxModel) MarkRetryableFailure(status int32, nextRetryAt string, leaseUntil string, lastErrorCode string, sideEffectId int64) (rowsAffected int64, err error) {
	var (
		query   = "update dialog_side_effect_outbox set `status` = ?, next_retry_at = ?, lease_owner = '', lease_until = ?, last_error_code = ? where side_effect_id = ?"
		rResult sql.Result
	)
	rResult, err = m.tx.Exec(query, status, nextRetryAt, leaseUntil, lastErrorCode, sideEffectId)

	if err != nil {
		err = fmt.Errorf("dialog_side_effect_outbox.MarkRetryableFailure exec: %w", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		err = fmt.Errorf("dialog_side_effect_outbox.MarkRetryableFailure rows affected: %w", err)
		return
	}

	return
}

// MarkBlocked
// update dialog_side_effect_outbox set `status` = :status, lease_owner = ”, lease_until = :lease_until, last_error_code = :last_error_code where side_effect_id = :side_effect_id
func (m *defaultDialogSideEffectOutboxModel) MarkBlocked(ctx context.Context, status int32, leaseUntil string, lastErrorCode string, sideEffectId int64) (rowsAffected int64, err error) {

	var (
		query   = "update dialog_side_effect_outbox set `status` = ?, lease_owner = '', lease_until = ?, last_error_code = ? where side_effect_id = ?"
		rResult sql.Result
	)

	rResult, err = m.db.Exec(ctx, query, status, leaseUntil, lastErrorCode, sideEffectId)

	if err != nil {
		err = fmt.Errorf("dialog_side_effect_outbox.MarkBlocked exec: %w", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		err = fmt.Errorf("dialog_side_effect_outbox.MarkBlocked rows affected: %w", err)
		return
	}

	return
}

// MarkBlocked
// update dialog_side_effect_outbox set `status` = :status, lease_owner = ”, lease_until = :lease_until, last_error_code = :last_error_code where side_effect_id = :side_effect_id
func (m *defaultDialogSideEffectOutboxTxModel) MarkBlocked(status int32, leaseUntil string, lastErrorCode string, sideEffectId int64) (rowsAffected int64, err error) {
	var (
		query   = "update dialog_side_effect_outbox set `status` = ?, lease_owner = '', lease_until = ?, last_error_code = ? where side_effect_id = ?"
		rResult sql.Result
	)
	rResult, err = m.tx.Exec(query, status, leaseUntil, lastErrorCode, sideEffectId)

	if err != nil {
		err = fmt.Errorf("dialog_side_effect_outbox.MarkBlocked exec: %w", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		err = fmt.Errorf("dialog_side_effect_outbox.MarkBlocked rows affected: %w", err)
		return
	}

	return
}

// SelectBySourceOperationKind
// select side_effect_id, kind, user_id, peer_type, peer_id, source_perm_auth_key_id, source_operation_id, source_message_date, source_peer_seq, source_canonical_message_id, clear_before_date, payload_schema_version, payload, payload_hash, `status`, attempt_count, next_retry_at, lease_owner, lease_until, last_error_code from dialog_side_effect_outbox where source_operation_id = :source_operation_id and kind = :kind order by side_effect_id asc
func (m *defaultDialogSideEffectOutboxModel) SelectBySourceOperationKind(ctx context.Context, sourceOperationId string, kind string) (rList []DialogSideEffectOutbox, err error) {
	var (
		query  = "select side_effect_id, kind, user_id, peer_type, peer_id, source_perm_auth_key_id, source_operation_id, source_message_date, source_peer_seq, source_canonical_message_id, clear_before_date, payload_schema_version, payload, payload_hash, `status`, attempt_count, next_retry_at, lease_owner, lease_until, last_error_code from dialog_side_effect_outbox where source_operation_id = ? and kind = ? order by side_effect_id asc"
		values []DialogSideEffectOutbox
	)
	err = m.db.QueryRowsPartial(ctx, &values, query, sourceOperationId, kind)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			rList = []DialogSideEffectOutbox{}
			err = nil
			return
		}
		err = fmt.Errorf("dialog_side_effect_outbox.SelectBySourceOperationKind: %w", err)
		return
	}

	rList = values

	return
}

// SelectBySourceOperationKind
// select side_effect_id, kind, user_id, peer_type, peer_id, source_perm_auth_key_id, source_operation_id, source_message_date, source_peer_seq, source_canonical_message_id, clear_before_date, payload_schema_version, payload, payload_hash, `status`, attempt_count, next_retry_at, lease_owner, lease_until, last_error_code from dialog_side_effect_outbox where source_operation_id = :source_operation_id and kind = :kind order by side_effect_id asc
func (m *defaultDialogSideEffectOutboxTxModel) SelectBySourceOperationKind(sourceOperationId string, kind string) (rList []DialogSideEffectOutbox, err error) {
	var (
		query  = "select side_effect_id, kind, user_id, peer_type, peer_id, source_perm_auth_key_id, source_operation_id, source_message_date, source_peer_seq, source_canonical_message_id, clear_before_date, payload_schema_version, payload, payload_hash, `status`, attempt_count, next_retry_at, lease_owner, lease_until, last_error_code from dialog_side_effect_outbox where source_operation_id = ? and kind = ? order by side_effect_id asc"
		values []DialogSideEffectOutbox
	)
	err = m.tx.QueryRowsPartial(&values, query, sourceOperationId, kind)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			rList = []DialogSideEffectOutbox{}
			err = nil
			return
		}
		err = fmt.Errorf("dialog_side_effect_outbox.SelectBySourceOperationKind: %w", err)
		return
	}

	rList = values

	return
}

// SelectBySourceOperationKindWithCB
// select side_effect_id, kind, user_id, peer_type, peer_id, source_perm_auth_key_id, source_operation_id, source_message_date, source_peer_seq, source_canonical_message_id, clear_before_date, payload_schema_version, payload, payload_hash, `status`, attempt_count, next_retry_at, lease_owner, lease_until, last_error_code from dialog_side_effect_outbox where source_operation_id = :source_operation_id and kind = :kind order by side_effect_id asc
func (m *defaultDialogSideEffectOutboxModel) SelectBySourceOperationKindWithCB(ctx context.Context, sourceOperationId string, kind string, cb func(sz, i int, v *DialogSideEffectOutbox)) (rList []DialogSideEffectOutbox, err error) {
	var (
		query  = "select side_effect_id, kind, user_id, peer_type, peer_id, source_perm_auth_key_id, source_operation_id, source_message_date, source_peer_seq, source_canonical_message_id, clear_before_date, payload_schema_version, payload, payload_hash, `status`, attempt_count, next_retry_at, lease_owner, lease_until, last_error_code from dialog_side_effect_outbox where source_operation_id = ? and kind = ? order by side_effect_id asc"
		values []DialogSideEffectOutbox
	)
	err = m.db.QueryRowsPartial(ctx, &values, query, sourceOperationId, kind)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			rList = []DialogSideEffectOutbox{}
			err = nil
			return
		}
		err = fmt.Errorf("dialog_side_effect_outbox.SelectBySourceOperationKindWithCB: %w", err)
		return
	}

	rList = values

	if cb != nil {
		sz := len(rList)
		for i := 0; i < sz; i++ {
			cb(sz, i, &rList[i])
		}
	}

	return
}
