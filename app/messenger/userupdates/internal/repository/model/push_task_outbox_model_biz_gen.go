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

type bizPushTaskOutboxModel interface {
	Insert(ctx context.Context, data *PushTaskOutbox) (lastInsertId, rowsAffected int64, err error)
	SelectPending(ctx context.Context, status int32, availableAt string, limit int32) ([]PushTaskOutbox, error)
	SelectPendingWithCB(ctx context.Context, status int32, availableAt string, limit int32, cb func(sz, i int, v *PushTaskOutbox)) ([]PushTaskOutbox, error)
	SelectDueForPublish(ctx context.Context, pendingStatus int32, failedRetryableStatus int32, availableAt string, limit int32) ([]PushTaskOutbox, error)
	SelectDueForPublishWithCB(ctx context.Context, pendingStatus int32, failedRetryableStatus int32, availableAt string, limit int32, cb func(sz, i int, v *PushTaskOutbox)) ([]PushTaskOutbox, error)
	MarkPublishing(ctx context.Context, status int32, taskId int64) (rowsAffected int64, err error)
	TryMarkPublishingDue(ctx context.Context, status int32, leaseExpiresAt string, taskId int64, pendingStatus int32, failedRetryableStatus int32, now string) (rowsAffected int64, err error)
	MarkPublished(ctx context.Context, status int32, publishedTopic string, publishedPartition int32, publishedOffset int64, publishedAt string, taskId int64) (rowsAffected int64, err error)
	MarkPublishFailed(ctx context.Context, status int32, nextRetryAt string, availableAt string, lastErrorCode string, taskId int64) (rowsAffected int64, err error)
	ResetExpiredPublishing(ctx context.Context, pendingStatus int32, availableAt string, publishingStatus int32, expiredAt string, limit int32) (rowsAffected int64, err error)
}

type PushTaskOutboxTxModel interface {
	Insert(data *PushTaskOutbox) (lastInsertId, rowsAffected int64, err error)
	SelectPending(status int32, availableAt string, limit int32) ([]PushTaskOutbox, error)
	SelectDueForPublish(pendingStatus int32, failedRetryableStatus int32, availableAt string, limit int32) ([]PushTaskOutbox, error)
	MarkPublishing(status int32, taskId int64) (rowsAffected int64, err error)
	TryMarkPublishingDue(status int32, leaseExpiresAt string, taskId int64, pendingStatus int32, failedRetryableStatus int32, now string) (rowsAffected int64, err error)
	MarkPublished(status int32, publishedTopic string, publishedPartition int32, publishedOffset int64, publishedAt string, taskId int64) (rowsAffected int64, err error)
	MarkPublishFailed(status int32, nextRetryAt string, availableAt string, lastErrorCode string, taskId int64) (rowsAffected int64, err error)
	ResetExpiredPublishing(pendingStatus int32, availableAt string, publishingStatus int32, expiredAt string, limit int32) (rowsAffected int64, err error)
}

type defaultPushTaskOutboxTxModel struct {
	tx *sqlx.Tx
}

func NewPushTaskOutboxTxModel(tx *sqlx.Tx) PushTaskOutboxTxModel {
	return &defaultPushTaskOutboxTxModel{tx: tx}
}

// Insert
// insert into push_task_outbox(task_id, user_id, pts, push_type, peer_type, peer_id, operation_id, push_partition_id, task_schema_version, task_codec, task_payload, `status`, publish_attempts, available_at, published_topic, published_partition, published_offset, last_error_code) values (:task_id, :user_id, :pts, :push_type, :peer_type, :peer_id, :operation_id, :push_partition_id, :task_schema_version, :task_codec, :task_payload, :status, :publish_attempts, :available_at, :published_topic, :published_partition, :published_offset, :last_error_code)
func (m *defaultPushTaskOutboxModel) Insert(ctx context.Context, data *PushTaskOutbox) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into push_task_outbox(task_id, user_id, pts, push_type, peer_type, peer_id, operation_id, push_partition_id, task_schema_version, task_codec, task_payload, `status`, publish_attempts, available_at, published_topic, published_partition, published_offset, last_error_code) values (:task_id, :user_id, :pts, :push_type, :peer_type, :peer_id, :operation_id, :push_partition_id, :task_schema_version, :task_codec, :task_payload, :status, :publish_attempts, :available_at, :published_topic, :published_partition, :published_offset, :last_error_code)"
		r     sql.Result
	)

	r, err = m.db.NamedExec(ctx, query, data)
	if err != nil {
		err = fmt.Errorf("push_task_outbox.Insert named exec: %w", err)
		return
	}

	lastInsertId, err = r.LastInsertId()
	if err != nil {
		err = fmt.Errorf("push_task_outbox.Insert last insert id: %w", err)
		return
	}
	rowsAffected, err = r.RowsAffected()
	if err != nil {
		err = fmt.Errorf("push_task_outbox.Insert rows affected: %w", err)
	}

	return

}

// Insert
// insert into push_task_outbox(task_id, user_id, pts, push_type, peer_type, peer_id, operation_id, push_partition_id, task_schema_version, task_codec, task_payload, `status`, publish_attempts, available_at, published_topic, published_partition, published_offset, last_error_code) values (:task_id, :user_id, :pts, :push_type, :peer_type, :peer_id, :operation_id, :push_partition_id, :task_schema_version, :task_codec, :task_payload, :status, :publish_attempts, :available_at, :published_topic, :published_partition, :published_offset, :last_error_code)
func (m *defaultPushTaskOutboxTxModel) Insert(data *PushTaskOutbox) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into push_task_outbox(task_id, user_id, pts, push_type, peer_type, peer_id, operation_id, push_partition_id, task_schema_version, task_codec, task_payload, `status`, publish_attempts, available_at, published_topic, published_partition, published_offset, last_error_code) values (:task_id, :user_id, :pts, :push_type, :peer_type, :peer_id, :operation_id, :push_partition_id, :task_schema_version, :task_codec, :task_payload, :status, :publish_attempts, :available_at, :published_topic, :published_partition, :published_offset, :last_error_code)"
		r     sql.Result
	)

	r, err = m.tx.NamedExec(query, data)
	if err != nil {
		err = fmt.Errorf("push_task_outbox.Insert named exec: %w", err)
		return
	}

	lastInsertId, err = r.LastInsertId()
	if err != nil {
		err = fmt.Errorf("push_task_outbox.Insert last insert id: %w", err)
		return
	}
	rowsAffected, err = r.RowsAffected()
	if err != nil {
		err = fmt.Errorf("push_task_outbox.Insert rows affected: %w", err)
	}

	return
}

// SelectPending
// select task_id, user_id, pts, push_type, peer_type, peer_id, operation_id, push_partition_id, task_schema_version, task_codec, task_payload, `status`, publish_attempts, available_at, published_topic, published_partition, published_offset, last_error_code from push_task_outbox where `status` = :status and available_at <= :available_at order by available_at asc, task_id asc limit :limit
func (m *defaultPushTaskOutboxModel) SelectPending(ctx context.Context, status int32, availableAt string, limit int32) (rList []PushTaskOutbox, err error) {
	var (
		query  = "select task_id, user_id, pts, push_type, peer_type, peer_id, operation_id, push_partition_id, task_schema_version, task_codec, task_payload, `status`, publish_attempts, available_at, published_topic, published_partition, published_offset, last_error_code from push_task_outbox where `status` = ? and available_at <= ? order by available_at asc, task_id asc limit ?"
		values []PushTaskOutbox
	)
	err = m.db.QueryRowsPartial(ctx, &values, query, status, availableAt, limit)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			rList = []PushTaskOutbox{}
			err = nil
			return
		}
		err = fmt.Errorf("push_task_outbox.SelectPending: %w", err)
		return
	}

	rList = values

	return
}

// SelectPending
// select task_id, user_id, pts, push_type, peer_type, peer_id, operation_id, push_partition_id, task_schema_version, task_codec, task_payload, `status`, publish_attempts, available_at, published_topic, published_partition, published_offset, last_error_code from push_task_outbox where `status` = :status and available_at <= :available_at order by available_at asc, task_id asc limit :limit
func (m *defaultPushTaskOutboxTxModel) SelectPending(status int32, availableAt string, limit int32) (rList []PushTaskOutbox, err error) {
	var (
		query  = "select task_id, user_id, pts, push_type, peer_type, peer_id, operation_id, push_partition_id, task_schema_version, task_codec, task_payload, `status`, publish_attempts, available_at, published_topic, published_partition, published_offset, last_error_code from push_task_outbox where `status` = ? and available_at <= ? order by available_at asc, task_id asc limit ?"
		values []PushTaskOutbox
	)
	err = m.tx.QueryRowsPartial(&values, query, status, availableAt, limit)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			rList = []PushTaskOutbox{}
			err = nil
			return
		}
		err = fmt.Errorf("push_task_outbox.SelectPending: %w", err)
		return
	}

	rList = values

	return
}

// SelectPendingWithCB
// select task_id, user_id, pts, push_type, peer_type, peer_id, operation_id, push_partition_id, task_schema_version, task_codec, task_payload, `status`, publish_attempts, available_at, published_topic, published_partition, published_offset, last_error_code from push_task_outbox where `status` = :status and available_at <= :available_at order by available_at asc, task_id asc limit :limit
func (m *defaultPushTaskOutboxModel) SelectPendingWithCB(ctx context.Context, status int32, availableAt string, limit int32, cb func(sz, i int, v *PushTaskOutbox)) (rList []PushTaskOutbox, err error) {
	var (
		query  = "select task_id, user_id, pts, push_type, peer_type, peer_id, operation_id, push_partition_id, task_schema_version, task_codec, task_payload, `status`, publish_attempts, available_at, published_topic, published_partition, published_offset, last_error_code from push_task_outbox where `status` = ? and available_at <= ? order by available_at asc, task_id asc limit ?"
		values []PushTaskOutbox
	)
	err = m.db.QueryRowsPartial(ctx, &values, query, status, availableAt, limit)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			rList = []PushTaskOutbox{}
			err = nil
			return
		}
		err = fmt.Errorf("push_task_outbox.SelectPendingWithCB: %w", err)
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

// SelectDueForPublish
// select task_id, user_id, pts, push_type, peer_type, peer_id, operation_id, push_partition_id, task_schema_version, task_codec, task_payload, `status`, publish_attempts, available_at, published_topic, published_partition, published_offset, last_error_code from push_task_outbox where `status` in (:pending_status, :failed_retryable_status) and available_at <= :available_at order by available_at asc, task_id asc limit :limit
func (m *defaultPushTaskOutboxModel) SelectDueForPublish(ctx context.Context, pendingStatus int32, failedRetryableStatus int32, availableAt string, limit int32) (rList []PushTaskOutbox, err error) {
	var (
		query  = "select task_id, user_id, pts, push_type, peer_type, peer_id, operation_id, push_partition_id, task_schema_version, task_codec, task_payload, `status`, publish_attempts, available_at, published_topic, published_partition, published_offset, last_error_code from push_task_outbox where `status` in (?, ?) and available_at <= ? order by available_at asc, task_id asc limit ?"
		values []PushTaskOutbox
	)
	err = m.db.QueryRowsPartial(ctx, &values, query, pendingStatus, failedRetryableStatus, availableAt, limit)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			rList = []PushTaskOutbox{}
			err = nil
			return
		}
		err = fmt.Errorf("push_task_outbox.SelectDueForPublish: %w", err)
		return
	}

	rList = values

	return
}

// SelectDueForPublish
// select task_id, user_id, pts, push_type, peer_type, peer_id, operation_id, push_partition_id, task_schema_version, task_codec, task_payload, `status`, publish_attempts, available_at, published_topic, published_partition, published_offset, last_error_code from push_task_outbox where `status` in (:pending_status, :failed_retryable_status) and available_at <= :available_at order by available_at asc, task_id asc limit :limit
func (m *defaultPushTaskOutboxTxModel) SelectDueForPublish(pendingStatus int32, failedRetryableStatus int32, availableAt string, limit int32) (rList []PushTaskOutbox, err error) {
	var (
		query  = "select task_id, user_id, pts, push_type, peer_type, peer_id, operation_id, push_partition_id, task_schema_version, task_codec, task_payload, `status`, publish_attempts, available_at, published_topic, published_partition, published_offset, last_error_code from push_task_outbox where `status` in (?, ?) and available_at <= ? order by available_at asc, task_id asc limit ?"
		values []PushTaskOutbox
	)
	err = m.tx.QueryRowsPartial(&values, query, pendingStatus, failedRetryableStatus, availableAt, limit)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			rList = []PushTaskOutbox{}
			err = nil
			return
		}
		err = fmt.Errorf("push_task_outbox.SelectDueForPublish: %w", err)
		return
	}

	rList = values

	return
}

// SelectDueForPublishWithCB
// select task_id, user_id, pts, push_type, peer_type, peer_id, operation_id, push_partition_id, task_schema_version, task_codec, task_payload, `status`, publish_attempts, available_at, published_topic, published_partition, published_offset, last_error_code from push_task_outbox where `status` in (:pending_status, :failed_retryable_status) and available_at <= :available_at order by available_at asc, task_id asc limit :limit
func (m *defaultPushTaskOutboxModel) SelectDueForPublishWithCB(ctx context.Context, pendingStatus int32, failedRetryableStatus int32, availableAt string, limit int32, cb func(sz, i int, v *PushTaskOutbox)) (rList []PushTaskOutbox, err error) {
	var (
		query  = "select task_id, user_id, pts, push_type, peer_type, peer_id, operation_id, push_partition_id, task_schema_version, task_codec, task_payload, `status`, publish_attempts, available_at, published_topic, published_partition, published_offset, last_error_code from push_task_outbox where `status` in (?, ?) and available_at <= ? order by available_at asc, task_id asc limit ?"
		values []PushTaskOutbox
	)
	err = m.db.QueryRowsPartial(ctx, &values, query, pendingStatus, failedRetryableStatus, availableAt, limit)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			rList = []PushTaskOutbox{}
			err = nil
			return
		}
		err = fmt.Errorf("push_task_outbox.SelectDueForPublishWithCB: %w", err)
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

// MarkPublishing
// update push_task_outbox set `status` = :status, publish_attempts = publish_attempts + 1 where task_id = :task_id
func (m *defaultPushTaskOutboxModel) MarkPublishing(ctx context.Context, status int32, taskId int64) (rowsAffected int64, err error) {

	var (
		query   = "update push_task_outbox set `status` = ?, publish_attempts = publish_attempts + 1 where task_id = ?"
		rResult sql.Result
	)

	rResult, err = m.db.Exec(ctx, query, status, taskId)

	if err != nil {
		err = fmt.Errorf("push_task_outbox.MarkPublishing exec: %w", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		err = fmt.Errorf("push_task_outbox.MarkPublishing rows affected: %w", err)
		return
	}

	return
}

// MarkPublishing
// update push_task_outbox set `status` = :status, publish_attempts = publish_attempts + 1 where task_id = :task_id
func (m *defaultPushTaskOutboxTxModel) MarkPublishing(status int32, taskId int64) (rowsAffected int64, err error) {
	var (
		query   = "update push_task_outbox set `status` = ?, publish_attempts = publish_attempts + 1 where task_id = ?"
		rResult sql.Result
	)
	rResult, err = m.tx.Exec(query, status, taskId)

	if err != nil {
		err = fmt.Errorf("push_task_outbox.MarkPublishing exec: %w", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		err = fmt.Errorf("push_task_outbox.MarkPublishing rows affected: %w", err)
		return
	}

	return
}

// TryMarkPublishingDue
// update push_task_outbox set `status` = :status, publish_attempts = publish_attempts + 1, available_at = :lease_expires_at where task_id = :task_id and `status` in (:pending_status, :failed_retryable_status) and available_at <= :now
func (m *defaultPushTaskOutboxModel) TryMarkPublishingDue(ctx context.Context, status int32, leaseExpiresAt string, taskId int64, pendingStatus int32, failedRetryableStatus int32, now string) (rowsAffected int64, err error) {

	var (
		query   = "update push_task_outbox set `status` = ?, publish_attempts = publish_attempts + 1, available_at = ? where task_id = ? and `status` in (?, ?) and available_at <= ?"
		rResult sql.Result
	)

	rResult, err = m.db.Exec(ctx, query, status, leaseExpiresAt, taskId, pendingStatus, failedRetryableStatus, now)

	if err != nil {
		err = fmt.Errorf("push_task_outbox.TryMarkPublishingDue exec: %w", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		err = fmt.Errorf("push_task_outbox.TryMarkPublishingDue rows affected: %w", err)
		return
	}

	return
}

// TryMarkPublishingDue
// update push_task_outbox set `status` = :status, publish_attempts = publish_attempts + 1, available_at = :lease_expires_at where task_id = :task_id and `status` in (:pending_status, :failed_retryable_status) and available_at <= :now
func (m *defaultPushTaskOutboxTxModel) TryMarkPublishingDue(status int32, leaseExpiresAt string, taskId int64, pendingStatus int32, failedRetryableStatus int32, now string) (rowsAffected int64, err error) {
	var (
		query   = "update push_task_outbox set `status` = ?, publish_attempts = publish_attempts + 1, available_at = ? where task_id = ? and `status` in (?, ?) and available_at <= ?"
		rResult sql.Result
	)
	rResult, err = m.tx.Exec(query, status, leaseExpiresAt, taskId, pendingStatus, failedRetryableStatus, now)

	if err != nil {
		err = fmt.Errorf("push_task_outbox.TryMarkPublishingDue exec: %w", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		err = fmt.Errorf("push_task_outbox.TryMarkPublishingDue rows affected: %w", err)
		return
	}

	return
}

// MarkPublished
// update push_task_outbox set `status` = :status, published_topic = :published_topic, published_partition = :published_partition, published_offset = :published_offset, published_at = :published_at where task_id = :task_id
func (m *defaultPushTaskOutboxModel) MarkPublished(ctx context.Context, status int32, publishedTopic string, publishedPartition int32, publishedOffset int64, publishedAt string, taskId int64) (rowsAffected int64, err error) {

	var (
		query   = "update push_task_outbox set `status` = ?, published_topic = ?, published_partition = ?, published_offset = ?, published_at = ? where task_id = ?"
		rResult sql.Result
	)

	rResult, err = m.db.Exec(ctx, query, status, publishedTopic, publishedPartition, publishedOffset, publishedAt, taskId)

	if err != nil {
		err = fmt.Errorf("push_task_outbox.MarkPublished exec: %w", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		err = fmt.Errorf("push_task_outbox.MarkPublished rows affected: %w", err)
		return
	}

	return
}

// MarkPublished
// update push_task_outbox set `status` = :status, published_topic = :published_topic, published_partition = :published_partition, published_offset = :published_offset, published_at = :published_at where task_id = :task_id
func (m *defaultPushTaskOutboxTxModel) MarkPublished(status int32, publishedTopic string, publishedPartition int32, publishedOffset int64, publishedAt string, taskId int64) (rowsAffected int64, err error) {
	var (
		query   = "update push_task_outbox set `status` = ?, published_topic = ?, published_partition = ?, published_offset = ?, published_at = ? where task_id = ?"
		rResult sql.Result
	)
	rResult, err = m.tx.Exec(query, status, publishedTopic, publishedPartition, publishedOffset, publishedAt, taskId)

	if err != nil {
		err = fmt.Errorf("push_task_outbox.MarkPublished exec: %w", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		err = fmt.Errorf("push_task_outbox.MarkPublished rows affected: %w", err)
		return
	}

	return
}

// MarkPublishFailed
// update push_task_outbox set `status` = :status, next_retry_at = :next_retry_at, available_at = :available_at, last_error_code = :last_error_code where task_id = :task_id
func (m *defaultPushTaskOutboxModel) MarkPublishFailed(ctx context.Context, status int32, nextRetryAt string, availableAt string, lastErrorCode string, taskId int64) (rowsAffected int64, err error) {

	var (
		query   = "update push_task_outbox set `status` = ?, next_retry_at = ?, available_at = ?, last_error_code = ? where task_id = ?"
		rResult sql.Result
	)

	rResult, err = m.db.Exec(ctx, query, status, nextRetryAt, availableAt, lastErrorCode, taskId)

	if err != nil {
		err = fmt.Errorf("push_task_outbox.MarkPublishFailed exec: %w", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		err = fmt.Errorf("push_task_outbox.MarkPublishFailed rows affected: %w", err)
		return
	}

	return
}

// MarkPublishFailed
// update push_task_outbox set `status` = :status, next_retry_at = :next_retry_at, available_at = :available_at, last_error_code = :last_error_code where task_id = :task_id
func (m *defaultPushTaskOutboxTxModel) MarkPublishFailed(status int32, nextRetryAt string, availableAt string, lastErrorCode string, taskId int64) (rowsAffected int64, err error) {
	var (
		query   = "update push_task_outbox set `status` = ?, next_retry_at = ?, available_at = ?, last_error_code = ? where task_id = ?"
		rResult sql.Result
	)
	rResult, err = m.tx.Exec(query, status, nextRetryAt, availableAt, lastErrorCode, taskId)

	if err != nil {
		err = fmt.Errorf("push_task_outbox.MarkPublishFailed exec: %w", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		err = fmt.Errorf("push_task_outbox.MarkPublishFailed rows affected: %w", err)
		return
	}

	return
}

// ResetExpiredPublishing
// update push_task_outbox set `status` = :pending_status, available_at = :available_at where `status` = :publishing_status and available_at <= :expired_at limit :limit
func (m *defaultPushTaskOutboxModel) ResetExpiredPublishing(ctx context.Context, pendingStatus int32, availableAt string, publishingStatus int32, expiredAt string, limit int32) (rowsAffected int64, err error) {

	var (
		query   = "update push_task_outbox set `status` = ?, available_at = ? where `status` = ? and available_at <= ? limit ?"
		rResult sql.Result
	)

	rResult, err = m.db.Exec(ctx, query, pendingStatus, availableAt, publishingStatus, expiredAt, limit)

	if err != nil {
		err = fmt.Errorf("push_task_outbox.ResetExpiredPublishing exec: %w", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		err = fmt.Errorf("push_task_outbox.ResetExpiredPublishing rows affected: %w", err)
		return
	}

	return
}

// ResetExpiredPublishing
// update push_task_outbox set `status` = :pending_status, available_at = :available_at where `status` = :publishing_status and available_at <= :expired_at limit :limit
func (m *defaultPushTaskOutboxTxModel) ResetExpiredPublishing(pendingStatus int32, availableAt string, publishingStatus int32, expiredAt string, limit int32) (rowsAffected int64, err error) {
	var (
		query   = "update push_task_outbox set `status` = ?, available_at = ? where `status` = ? and available_at <= ? limit ?"
		rResult sql.Result
	)
	rResult, err = m.tx.Exec(query, pendingStatus, availableAt, publishingStatus, expiredAt, limit)

	if err != nil {
		err = fmt.Errorf("push_task_outbox.ResetExpiredPublishing exec: %w", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		err = fmt.Errorf("push_task_outbox.ResetExpiredPublishing rows affected: %w", err)
		return
	}

	return
}
