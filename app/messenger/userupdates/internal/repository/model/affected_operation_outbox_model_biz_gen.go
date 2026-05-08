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

type bizAffectedOperationOutboxModel interface {
	InsertIgnore(ctx context.Context, data *AffectedOperationOutbox) (lastInsertId, rowsAffected int64, err error)
	SelectByUserOperation(ctx context.Context, userId int64, operationId string) (*AffectedOperationOutbox, error)
	SelectPending(ctx context.Context, status int32, availableAt int64, limit int32) ([]AffectedOperationOutbox, error)
	SelectPendingWithCB(ctx context.Context, status int32, availableAt int64, limit int32, cb func(sz, i int, v *AffectedOperationOutbox)) ([]AffectedOperationOutbox, error)
	TryMarkProcessing(ctx context.Context, processingStatus int32, processingDeadline int64, outboxId int64, pendingStatus int32, now int64) (rowsAffected int64, err error)
	MarkCompleted(ctx context.Context, status int32, outboxId int64, processingStatus int32, processingDeadline int64) (rowsAffected int64, err error)
	MarkRetryable(ctx context.Context, status int32, availableAt int64, lastErrorCode string, lastErrorMessage string, outboxId int64, processingStatus int32, processingDeadline int64) (rowsAffected int64, err error)
	MarkFailedTerminal(ctx context.Context, status int32, lastErrorCode string, lastErrorMessage string, outboxId int64, processingStatus int32, processingDeadline int64) (rowsAffected int64, err error)
	ResetExpiredProcessing(ctx context.Context, pendingStatus int32, now int64, processingStatus int32, limit int32) (rowsAffected int64, err error)
	SelectTerminalBefore(ctx context.Context, completedStatus int32, failedTerminalStatus int32, beforeUpdatedAt string, limit int32) ([]AffectedOperationOutbox, error)
	SelectTerminalBeforeWithCB(ctx context.Context, completedStatus int32, failedTerminalStatus int32, beforeUpdatedAt string, limit int32, cb func(sz, i int, v *AffectedOperationOutbox)) ([]AffectedOperationOutbox, error)
}

type AffectedOperationOutboxTxModel interface {
	InsertIgnore(data *AffectedOperationOutbox) (lastInsertId, rowsAffected int64, err error)
	SelectByUserOperation(userId int64, operationId string) (*AffectedOperationOutbox, error)
	SelectPending(status int32, availableAt int64, limit int32) ([]AffectedOperationOutbox, error)
	TryMarkProcessing(processingStatus int32, processingDeadline int64, outboxId int64, pendingStatus int32, now int64) (rowsAffected int64, err error)
	MarkCompleted(status int32, outboxId int64, processingStatus int32, processingDeadline int64) (rowsAffected int64, err error)
	MarkRetryable(status int32, availableAt int64, lastErrorCode string, lastErrorMessage string, outboxId int64, processingStatus int32, processingDeadline int64) (rowsAffected int64, err error)
	MarkFailedTerminal(status int32, lastErrorCode string, lastErrorMessage string, outboxId int64, processingStatus int32, processingDeadline int64) (rowsAffected int64, err error)
	ResetExpiredProcessing(pendingStatus int32, now int64, processingStatus int32, limit int32) (rowsAffected int64, err error)
	SelectTerminalBefore(completedStatus int32, failedTerminalStatus int32, beforeUpdatedAt string, limit int32) ([]AffectedOperationOutbox, error)
}

type defaultAffectedOperationOutboxTxModel struct {
	tx *sqlx.Tx
}

func NewAffectedOperationOutboxTxModel(tx *sqlx.Tx) AffectedOperationOutboxTxModel {
	return &defaultAffectedOperationOutboxTxModel{tx: tx}
}

// InsertIgnore
// insert ignore into affected_operation_outbox(outbox_id, user_id, requester_user_id, operation_id, op_type, operation_kind, peer_type, peer_id, payload_codec, payload_hash, payload, delivery_policy, `status`, retry_count, available_at, processing_deadline, last_error_code, last_error_message, bucket_id, partition_id, owner_token_payload) values (:outbox_id, :user_id, :requester_user_id, :operation_id, :op_type, :operation_kind, :peer_type, :peer_id, :payload_codec, :payload_hash, :payload, :delivery_policy, :status, :retry_count, :available_at, :processing_deadline, :last_error_code, :last_error_message, :bucket_id, :partition_id, :owner_token_payload)
func (m *defaultAffectedOperationOutboxModel) InsertIgnore(ctx context.Context, data *AffectedOperationOutbox) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert ignore into affected_operation_outbox(outbox_id, user_id, requester_user_id, operation_id, op_type, operation_kind, peer_type, peer_id, payload_codec, payload_hash, payload, delivery_policy, `status`, retry_count, available_at, processing_deadline, last_error_code, last_error_message, bucket_id, partition_id, owner_token_payload) values (:outbox_id, :user_id, :requester_user_id, :operation_id, :op_type, :operation_kind, :peer_type, :peer_id, :payload_codec, :payload_hash, :payload, :delivery_policy, :status, :retry_count, :available_at, :processing_deadline, :last_error_code, :last_error_message, :bucket_id, :partition_id, :owner_token_payload)"
		r     sql.Result
	)

	r, err = m.db.NamedExec(ctx, query, data)
	if err != nil {
		err = fmt.Errorf("affected_operation_outbox.InsertIgnore named exec: %w", err)
		return
	}

	lastInsertId, err = r.LastInsertId()
	if err != nil {
		err = fmt.Errorf("affected_operation_outbox.InsertIgnore last insert id: %w", err)
		return
	}
	rowsAffected, err = r.RowsAffected()
	if err != nil {
		err = fmt.Errorf("affected_operation_outbox.InsertIgnore rows affected: %w", err)
	}

	return

}

// InsertIgnore
// insert ignore into affected_operation_outbox(outbox_id, user_id, requester_user_id, operation_id, op_type, operation_kind, peer_type, peer_id, payload_codec, payload_hash, payload, delivery_policy, `status`, retry_count, available_at, processing_deadline, last_error_code, last_error_message, bucket_id, partition_id, owner_token_payload) values (:outbox_id, :user_id, :requester_user_id, :operation_id, :op_type, :operation_kind, :peer_type, :peer_id, :payload_codec, :payload_hash, :payload, :delivery_policy, :status, :retry_count, :available_at, :processing_deadline, :last_error_code, :last_error_message, :bucket_id, :partition_id, :owner_token_payload)
func (m *defaultAffectedOperationOutboxTxModel) InsertIgnore(data *AffectedOperationOutbox) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert ignore into affected_operation_outbox(outbox_id, user_id, requester_user_id, operation_id, op_type, operation_kind, peer_type, peer_id, payload_codec, payload_hash, payload, delivery_policy, `status`, retry_count, available_at, processing_deadline, last_error_code, last_error_message, bucket_id, partition_id, owner_token_payload) values (:outbox_id, :user_id, :requester_user_id, :operation_id, :op_type, :operation_kind, :peer_type, :peer_id, :payload_codec, :payload_hash, :payload, :delivery_policy, :status, :retry_count, :available_at, :processing_deadline, :last_error_code, :last_error_message, :bucket_id, :partition_id, :owner_token_payload)"
		r     sql.Result
	)

	r, err = m.tx.NamedExec(query, data)
	if err != nil {
		err = fmt.Errorf("affected_operation_outbox.InsertIgnore named exec: %w", err)
		return
	}

	lastInsertId, err = r.LastInsertId()
	if err != nil {
		err = fmt.Errorf("affected_operation_outbox.InsertIgnore last insert id: %w", err)
		return
	}
	rowsAffected, err = r.RowsAffected()
	if err != nil {
		err = fmt.Errorf("affected_operation_outbox.InsertIgnore rows affected: %w", err)
	}

	return
}

// SelectByUserOperation
// select outbox_id, user_id, requester_user_id, operation_id, op_type, operation_kind, peer_type, peer_id, payload_codec, payload_hash, payload, delivery_policy, `status`, retry_count, available_at, processing_deadline, last_error_code, last_error_message, bucket_id, partition_id, owner_token_payload from affected_operation_outbox where user_id = :user_id and operation_id = :operation_id limit 1
func (m *defaultAffectedOperationOutboxModel) SelectByUserOperation(ctx context.Context, userId int64, operationId string) (rValue *AffectedOperationOutbox, err error) {

	var (
		query = "select outbox_id, user_id, requester_user_id, operation_id, op_type, operation_kind, peer_type, peer_id, payload_codec, payload_hash, payload, delivery_policy, `status`, retry_count, available_at, processing_deadline, last_error_code, last_error_message, bucket_id, partition_id, owner_token_payload from affected_operation_outbox where user_id = ? and operation_id = ? limit 1"
		do    = &AffectedOperationOutbox{}
	)
	err = m.db.QueryRowPartial(ctx, do, query, userId, operationId)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return nil, &NotFoundError{
				Resource: "affected_operation_outbox",
				Key:      fmt.Sprintf("user_id=%v,operation_id=%v", userId, operationId),
				Cause:    err,
			}
		}
		err = fmt.Errorf("affected_operation_outbox.SelectByUserOperation: %w", err)
		return
	} else {
		rValue = do
	}

	return
}

// SelectByUserOperation
// select outbox_id, user_id, requester_user_id, operation_id, op_type, operation_kind, peer_type, peer_id, payload_codec, payload_hash, payload, delivery_policy, `status`, retry_count, available_at, processing_deadline, last_error_code, last_error_message, bucket_id, partition_id, owner_token_payload from affected_operation_outbox where user_id = :user_id and operation_id = :operation_id limit 1
func (m *defaultAffectedOperationOutboxTxModel) SelectByUserOperation(userId int64, operationId string) (rValue *AffectedOperationOutbox, err error) {
	var (
		query = "select outbox_id, user_id, requester_user_id, operation_id, op_type, operation_kind, peer_type, peer_id, payload_codec, payload_hash, payload, delivery_policy, `status`, retry_count, available_at, processing_deadline, last_error_code, last_error_message, bucket_id, partition_id, owner_token_payload from affected_operation_outbox where user_id = ? and operation_id = ? limit 1"
		do    = &AffectedOperationOutbox{}
	)
	err = m.tx.QueryRowPartial(do, query, userId, operationId)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return nil, &NotFoundError{
				Resource: "affected_operation_outbox",
				Key:      fmt.Sprintf("user_id=%v,operation_id=%v", userId, operationId),
				Cause:    err,
			}
		}
		err = fmt.Errorf("affected_operation_outbox.SelectByUserOperation: %w", err)
		return
	}
	rValue = do

	return
}

// SelectPending
// select outbox_id, user_id, requester_user_id, operation_id, op_type, operation_kind, peer_type, peer_id, payload_codec, payload_hash, payload, delivery_policy, `status`, retry_count, available_at, processing_deadline, last_error_code, last_error_message, bucket_id, partition_id, owner_token_payload from affected_operation_outbox where `status` = :status and available_at <= :available_at order by available_at asc, outbox_id asc limit :limit
func (m *defaultAffectedOperationOutboxModel) SelectPending(ctx context.Context, status int32, availableAt int64, limit int32) (rList []AffectedOperationOutbox, err error) {
	var (
		query  = "select outbox_id, user_id, requester_user_id, operation_id, op_type, operation_kind, peer_type, peer_id, payload_codec, payload_hash, payload, delivery_policy, `status`, retry_count, available_at, processing_deadline, last_error_code, last_error_message, bucket_id, partition_id, owner_token_payload from affected_operation_outbox where `status` = ? and available_at <= ? order by available_at asc, outbox_id asc limit ?"
		values []AffectedOperationOutbox
	)
	err = m.db.QueryRowsPartial(ctx, &values, query, status, availableAt, limit)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			rList = []AffectedOperationOutbox{}
			err = nil
			return
		}
		err = fmt.Errorf("affected_operation_outbox.SelectPending: %w", err)
		return
	}

	rList = values

	return
}

// SelectPending
// select outbox_id, user_id, requester_user_id, operation_id, op_type, operation_kind, peer_type, peer_id, payload_codec, payload_hash, payload, delivery_policy, `status`, retry_count, available_at, processing_deadline, last_error_code, last_error_message, bucket_id, partition_id, owner_token_payload from affected_operation_outbox where `status` = :status and available_at <= :available_at order by available_at asc, outbox_id asc limit :limit
func (m *defaultAffectedOperationOutboxTxModel) SelectPending(status int32, availableAt int64, limit int32) (rList []AffectedOperationOutbox, err error) {
	var (
		query  = "select outbox_id, user_id, requester_user_id, operation_id, op_type, operation_kind, peer_type, peer_id, payload_codec, payload_hash, payload, delivery_policy, `status`, retry_count, available_at, processing_deadline, last_error_code, last_error_message, bucket_id, partition_id, owner_token_payload from affected_operation_outbox where `status` = ? and available_at <= ? order by available_at asc, outbox_id asc limit ?"
		values []AffectedOperationOutbox
	)
	err = m.tx.QueryRowsPartial(&values, query, status, availableAt, limit)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			rList = []AffectedOperationOutbox{}
			err = nil
			return
		}
		err = fmt.Errorf("affected_operation_outbox.SelectPending: %w", err)
		return
	}

	rList = values

	return
}

// SelectPendingWithCB
// select outbox_id, user_id, requester_user_id, operation_id, op_type, operation_kind, peer_type, peer_id, payload_codec, payload_hash, payload, delivery_policy, `status`, retry_count, available_at, processing_deadline, last_error_code, last_error_message, bucket_id, partition_id, owner_token_payload from affected_operation_outbox where `status` = :status and available_at <= :available_at order by available_at asc, outbox_id asc limit :limit
func (m *defaultAffectedOperationOutboxModel) SelectPendingWithCB(ctx context.Context, status int32, availableAt int64, limit int32, cb func(sz, i int, v *AffectedOperationOutbox)) (rList []AffectedOperationOutbox, err error) {
	var (
		query  = "select outbox_id, user_id, requester_user_id, operation_id, op_type, operation_kind, peer_type, peer_id, payload_codec, payload_hash, payload, delivery_policy, `status`, retry_count, available_at, processing_deadline, last_error_code, last_error_message, bucket_id, partition_id, owner_token_payload from affected_operation_outbox where `status` = ? and available_at <= ? order by available_at asc, outbox_id asc limit ?"
		values []AffectedOperationOutbox
	)
	err = m.db.QueryRowsPartial(ctx, &values, query, status, availableAt, limit)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			rList = []AffectedOperationOutbox{}
			err = nil
			return
		}
		err = fmt.Errorf("affected_operation_outbox.SelectPendingWithCB: %w", err)
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

// TryMarkProcessing
// update affected_operation_outbox set `status` = :processing_status, processing_deadline = :processing_deadline where outbox_id = :outbox_id and `status` = :pending_status and available_at <= :now
func (m *defaultAffectedOperationOutboxModel) TryMarkProcessing(ctx context.Context, processingStatus int32, processingDeadline int64, outboxId int64, pendingStatus int32, now int64) (rowsAffected int64, err error) {

	var (
		query   = "update affected_operation_outbox set `status` = ?, processing_deadline = ? where outbox_id = ? and `status` = ? and available_at <= ?"
		rResult sql.Result
	)

	rResult, err = m.db.Exec(ctx, query, processingStatus, processingDeadline, outboxId, pendingStatus, now)

	if err != nil {
		err = fmt.Errorf("affected_operation_outbox.TryMarkProcessing exec: %w", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		err = fmt.Errorf("affected_operation_outbox.TryMarkProcessing rows affected: %w", err)
		return
	}

	return
}

// TryMarkProcessing
// update affected_operation_outbox set `status` = :processing_status, processing_deadline = :processing_deadline where outbox_id = :outbox_id and `status` = :pending_status and available_at <= :now
func (m *defaultAffectedOperationOutboxTxModel) TryMarkProcessing(processingStatus int32, processingDeadline int64, outboxId int64, pendingStatus int32, now int64) (rowsAffected int64, err error) {
	var (
		query   = "update affected_operation_outbox set `status` = ?, processing_deadline = ? where outbox_id = ? and `status` = ? and available_at <= ?"
		rResult sql.Result
	)
	rResult, err = m.tx.Exec(query, processingStatus, processingDeadline, outboxId, pendingStatus, now)

	if err != nil {
		err = fmt.Errorf("affected_operation_outbox.TryMarkProcessing exec: %w", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		err = fmt.Errorf("affected_operation_outbox.TryMarkProcessing rows affected: %w", err)
		return
	}

	return
}

// MarkCompleted
// update affected_operation_outbox set `status` = :status, processing_deadline = 0, last_error_code = ”, last_error_message = ” where outbox_id = :outbox_id and `status` = :processing_status and processing_deadline = :processing_deadline
func (m *defaultAffectedOperationOutboxModel) MarkCompleted(ctx context.Context, status int32, outboxId int64, processingStatus int32, processingDeadline int64) (rowsAffected int64, err error) {

	var (
		query   = "update affected_operation_outbox set `status` = ?, processing_deadline = 0, last_error_code = '', last_error_message = '' where outbox_id = ? and `status` = ? and processing_deadline = ?"
		rResult sql.Result
	)

	rResult, err = m.db.Exec(ctx, query, status, outboxId, processingStatus, processingDeadline)

	if err != nil {
		err = fmt.Errorf("affected_operation_outbox.MarkCompleted exec: %w", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		err = fmt.Errorf("affected_operation_outbox.MarkCompleted rows affected: %w", err)
		return
	}

	return
}

// MarkCompleted
// update affected_operation_outbox set `status` = :status, processing_deadline = 0, last_error_code = ”, last_error_message = ” where outbox_id = :outbox_id and `status` = :processing_status and processing_deadline = :processing_deadline
func (m *defaultAffectedOperationOutboxTxModel) MarkCompleted(status int32, outboxId int64, processingStatus int32, processingDeadline int64) (rowsAffected int64, err error) {
	var (
		query   = "update affected_operation_outbox set `status` = ?, processing_deadline = 0, last_error_code = '', last_error_message = '' where outbox_id = ? and `status` = ? and processing_deadline = ?"
		rResult sql.Result
	)
	rResult, err = m.tx.Exec(query, status, outboxId, processingStatus, processingDeadline)

	if err != nil {
		err = fmt.Errorf("affected_operation_outbox.MarkCompleted exec: %w", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		err = fmt.Errorf("affected_operation_outbox.MarkCompleted rows affected: %w", err)
		return
	}

	return
}

// MarkRetryable
// update affected_operation_outbox set `status` = :status, retry_count = retry_count + 1, available_at = :available_at, processing_deadline = 0, last_error_code = :last_error_code, last_error_message = :last_error_message where outbox_id = :outbox_id and `status` = :processing_status and processing_deadline = :processing_deadline
func (m *defaultAffectedOperationOutboxModel) MarkRetryable(ctx context.Context, status int32, availableAt int64, lastErrorCode string, lastErrorMessage string, outboxId int64, processingStatus int32, processingDeadline int64) (rowsAffected int64, err error) {

	var (
		query   = "update affected_operation_outbox set `status` = ?, retry_count = retry_count + 1, available_at = ?, processing_deadline = 0, last_error_code = ?, last_error_message = ? where outbox_id = ? and `status` = ? and processing_deadline = ?"
		rResult sql.Result
	)

	rResult, err = m.db.Exec(ctx, query, status, availableAt, lastErrorCode, lastErrorMessage, outboxId, processingStatus, processingDeadline)

	if err != nil {
		err = fmt.Errorf("affected_operation_outbox.MarkRetryable exec: %w", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		err = fmt.Errorf("affected_operation_outbox.MarkRetryable rows affected: %w", err)
		return
	}

	return
}

// MarkRetryable
// update affected_operation_outbox set `status` = :status, retry_count = retry_count + 1, available_at = :available_at, processing_deadline = 0, last_error_code = :last_error_code, last_error_message = :last_error_message where outbox_id = :outbox_id and `status` = :processing_status and processing_deadline = :processing_deadline
func (m *defaultAffectedOperationOutboxTxModel) MarkRetryable(status int32, availableAt int64, lastErrorCode string, lastErrorMessage string, outboxId int64, processingStatus int32, processingDeadline int64) (rowsAffected int64, err error) {
	var (
		query   = "update affected_operation_outbox set `status` = ?, retry_count = retry_count + 1, available_at = ?, processing_deadline = 0, last_error_code = ?, last_error_message = ? where outbox_id = ? and `status` = ? and processing_deadline = ?"
		rResult sql.Result
	)
	rResult, err = m.tx.Exec(query, status, availableAt, lastErrorCode, lastErrorMessage, outboxId, processingStatus, processingDeadline)

	if err != nil {
		err = fmt.Errorf("affected_operation_outbox.MarkRetryable exec: %w", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		err = fmt.Errorf("affected_operation_outbox.MarkRetryable rows affected: %w", err)
		return
	}

	return
}

// MarkFailedTerminal
// update affected_operation_outbox set `status` = :status, processing_deadline = 0, last_error_code = :last_error_code, last_error_message = :last_error_message where outbox_id = :outbox_id and `status` = :processing_status and processing_deadline = :processing_deadline
func (m *defaultAffectedOperationOutboxModel) MarkFailedTerminal(ctx context.Context, status int32, lastErrorCode string, lastErrorMessage string, outboxId int64, processingStatus int32, processingDeadline int64) (rowsAffected int64, err error) {

	var (
		query   = "update affected_operation_outbox set `status` = ?, processing_deadline = 0, last_error_code = ?, last_error_message = ? where outbox_id = ? and `status` = ? and processing_deadline = ?"
		rResult sql.Result
	)

	rResult, err = m.db.Exec(ctx, query, status, lastErrorCode, lastErrorMessage, outboxId, processingStatus, processingDeadline)

	if err != nil {
		err = fmt.Errorf("affected_operation_outbox.MarkFailedTerminal exec: %w", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		err = fmt.Errorf("affected_operation_outbox.MarkFailedTerminal rows affected: %w", err)
		return
	}

	return
}

// MarkFailedTerminal
// update affected_operation_outbox set `status` = :status, processing_deadline = 0, last_error_code = :last_error_code, last_error_message = :last_error_message where outbox_id = :outbox_id and `status` = :processing_status and processing_deadline = :processing_deadline
func (m *defaultAffectedOperationOutboxTxModel) MarkFailedTerminal(status int32, lastErrorCode string, lastErrorMessage string, outboxId int64, processingStatus int32, processingDeadline int64) (rowsAffected int64, err error) {
	var (
		query   = "update affected_operation_outbox set `status` = ?, processing_deadline = 0, last_error_code = ?, last_error_message = ? where outbox_id = ? and `status` = ? and processing_deadline = ?"
		rResult sql.Result
	)
	rResult, err = m.tx.Exec(query, status, lastErrorCode, lastErrorMessage, outboxId, processingStatus, processingDeadline)

	if err != nil {
		err = fmt.Errorf("affected_operation_outbox.MarkFailedTerminal exec: %w", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		err = fmt.Errorf("affected_operation_outbox.MarkFailedTerminal rows affected: %w", err)
		return
	}

	return
}

// ResetExpiredProcessing
// update affected_operation_outbox set `status` = :pending_status, available_at = :now, processing_deadline = 0 where `status` = :processing_status and processing_deadline > 0 and processing_deadline < :now limit :limit
func (m *defaultAffectedOperationOutboxModel) ResetExpiredProcessing(ctx context.Context, pendingStatus int32, now int64, processingStatus int32, limit int32) (rowsAffected int64, err error) {

	var (
		query   = "update affected_operation_outbox set `status` = ?, available_at = ?, processing_deadline = 0 where `status` = ? and processing_deadline > 0 and processing_deadline < ? limit ?"
		rResult sql.Result
	)

	rResult, err = m.db.Exec(ctx, query, pendingStatus, now, processingStatus, now, limit)

	if err != nil {
		err = fmt.Errorf("affected_operation_outbox.ResetExpiredProcessing exec: %w", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		err = fmt.Errorf("affected_operation_outbox.ResetExpiredProcessing rows affected: %w", err)
		return
	}

	return
}

// ResetExpiredProcessing
// update affected_operation_outbox set `status` = :pending_status, available_at = :now, processing_deadline = 0 where `status` = :processing_status and processing_deadline > 0 and processing_deadline < :now limit :limit
func (m *defaultAffectedOperationOutboxTxModel) ResetExpiredProcessing(pendingStatus int32, now int64, processingStatus int32, limit int32) (rowsAffected int64, err error) {
	var (
		query   = "update affected_operation_outbox set `status` = ?, available_at = ?, processing_deadline = 0 where `status` = ? and processing_deadline > 0 and processing_deadline < ? limit ?"
		rResult sql.Result
	)
	rResult, err = m.tx.Exec(query, pendingStatus, now, processingStatus, now, limit)

	if err != nil {
		err = fmt.Errorf("affected_operation_outbox.ResetExpiredProcessing exec: %w", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		err = fmt.Errorf("affected_operation_outbox.ResetExpiredProcessing rows affected: %w", err)
		return
	}

	return
}

// SelectTerminalBefore
// select outbox_id, user_id, requester_user_id, operation_id, op_type, operation_kind, peer_type, peer_id, payload_codec, payload_hash, payload, delivery_policy, `status`, retry_count, available_at, processing_deadline, last_error_code, last_error_message, bucket_id, partition_id, owner_token_payload from affected_operation_outbox where `status` in (:completed_status, :failed_terminal_status) and updated_at < :before_updated_at order by updated_at asc, outbox_id asc limit :limit
func (m *defaultAffectedOperationOutboxModel) SelectTerminalBefore(ctx context.Context, completedStatus int32, failedTerminalStatus int32, beforeUpdatedAt string, limit int32) (rList []AffectedOperationOutbox, err error) {
	var (
		query  = "select outbox_id, user_id, requester_user_id, operation_id, op_type, operation_kind, peer_type, peer_id, payload_codec, payload_hash, payload, delivery_policy, `status`, retry_count, available_at, processing_deadline, last_error_code, last_error_message, bucket_id, partition_id, owner_token_payload from affected_operation_outbox where `status` in (?, ?) and updated_at < ? order by updated_at asc, outbox_id asc limit ?"
		values []AffectedOperationOutbox
	)
	err = m.db.QueryRowsPartial(ctx, &values, query, completedStatus, failedTerminalStatus, beforeUpdatedAt, limit)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			rList = []AffectedOperationOutbox{}
			err = nil
			return
		}
		err = fmt.Errorf("affected_operation_outbox.SelectTerminalBefore: %w", err)
		return
	}

	rList = values

	return
}

// SelectTerminalBefore
// select outbox_id, user_id, requester_user_id, operation_id, op_type, operation_kind, peer_type, peer_id, payload_codec, payload_hash, payload, delivery_policy, `status`, retry_count, available_at, processing_deadline, last_error_code, last_error_message, bucket_id, partition_id, owner_token_payload from affected_operation_outbox where `status` in (:completed_status, :failed_terminal_status) and updated_at < :before_updated_at order by updated_at asc, outbox_id asc limit :limit
func (m *defaultAffectedOperationOutboxTxModel) SelectTerminalBefore(completedStatus int32, failedTerminalStatus int32, beforeUpdatedAt string, limit int32) (rList []AffectedOperationOutbox, err error) {
	var (
		query  = "select outbox_id, user_id, requester_user_id, operation_id, op_type, operation_kind, peer_type, peer_id, payload_codec, payload_hash, payload, delivery_policy, `status`, retry_count, available_at, processing_deadline, last_error_code, last_error_message, bucket_id, partition_id, owner_token_payload from affected_operation_outbox where `status` in (?, ?) and updated_at < ? order by updated_at asc, outbox_id asc limit ?"
		values []AffectedOperationOutbox
	)
	err = m.tx.QueryRowsPartial(&values, query, completedStatus, failedTerminalStatus, beforeUpdatedAt, limit)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			rList = []AffectedOperationOutbox{}
			err = nil
			return
		}
		err = fmt.Errorf("affected_operation_outbox.SelectTerminalBefore: %w", err)
		return
	}

	rList = values

	return
}

// SelectTerminalBeforeWithCB
// select outbox_id, user_id, requester_user_id, operation_id, op_type, operation_kind, peer_type, peer_id, payload_codec, payload_hash, payload, delivery_policy, `status`, retry_count, available_at, processing_deadline, last_error_code, last_error_message, bucket_id, partition_id, owner_token_payload from affected_operation_outbox where `status` in (:completed_status, :failed_terminal_status) and updated_at < :before_updated_at order by updated_at asc, outbox_id asc limit :limit
func (m *defaultAffectedOperationOutboxModel) SelectTerminalBeforeWithCB(ctx context.Context, completedStatus int32, failedTerminalStatus int32, beforeUpdatedAt string, limit int32, cb func(sz, i int, v *AffectedOperationOutbox)) (rList []AffectedOperationOutbox, err error) {
	var (
		query  = "select outbox_id, user_id, requester_user_id, operation_id, op_type, operation_kind, peer_type, peer_id, payload_codec, payload_hash, payload, delivery_policy, `status`, retry_count, available_at, processing_deadline, last_error_code, last_error_message, bucket_id, partition_id, owner_token_payload from affected_operation_outbox where `status` in (?, ?) and updated_at < ? order by updated_at asc, outbox_id asc limit ?"
		values []AffectedOperationOutbox
	)
	err = m.db.QueryRowsPartial(ctx, &values, query, completedStatus, failedTerminalStatus, beforeUpdatedAt, limit)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			rList = []AffectedOperationOutbox{}
			err = nil
			return
		}
		err = fmt.Errorf("affected_operation_outbox.SelectTerminalBeforeWithCB: %w", err)
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
