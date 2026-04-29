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

type (
	bizPushTaskOutboxModel interface {
		Insert(ctx context.Context, data *PushTaskOutbox) (lastInsertId, rowsAffected int64, err error)
		InsertTx(tx *sqlx.Tx, data *PushTaskOutbox) (lastInsertId, rowsAffected int64, err error)

		SelectPending(ctx context.Context, status int32, nextRetryAt string, limit int32) ([]PushTaskOutbox, error)
		SelectPendingWithCB(ctx context.Context, status int32, nextRetryAt string, limit int32, cb func(sz, i int, v *PushTaskOutbox)) ([]PushTaskOutbox, error)

		MarkPublishing(ctx context.Context, status int32, taskId int64) (rowsAffected int64, err error)
		MarkPublishingTx(tx *sqlx.Tx, status int32, taskId int64) (rowsAffected int64, err error)

		MarkPublished(ctx context.Context, status int32, publishedTopic string, publishedPartition int32, publishedOffset int64, publishedAt string, taskId int64) (rowsAffected int64, err error)
		MarkPublishedTx(tx *sqlx.Tx, status int32, publishedTopic string, publishedPartition int32, publishedOffset int64, publishedAt string, taskId int64) (rowsAffected int64, err error)

		MarkPublishFailed(ctx context.Context, status int32, nextRetryAt string, lastErrorCode string, taskId int64) (rowsAffected int64, err error)
		MarkPublishFailedTx(tx *sqlx.Tx, status int32, nextRetryAt string, lastErrorCode string, taskId int64) (rowsAffected int64, err error)
	}
)

// Insert
// insert into push_task_outbox(task_id, user_id, pts, push_type, peer_type, peer_id, operation_id, push_partition_id, task_schema_version, task_codec, task_payload, `status`, publish_attempts, next_retry_at, published_topic, published_partition, published_offset, last_error_code, created_at, updated_at, published_at) values (:task_id, :user_id, :pts, :push_type, :peer_type, :peer_id, :operation_id, :push_partition_id, :task_schema_version, :task_codec, :task_payload, :status, :publish_attempts, :next_retry_at, :published_topic, :published_partition, :published_offset, :last_error_code, NOW(6), NOW(6), :published_at)
func (m *defaultPushTaskOutboxModel) Insert(ctx context.Context, data *PushTaskOutbox) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into push_task_outbox(task_id, user_id, pts, push_type, peer_type, peer_id, operation_id, push_partition_id, task_schema_version, task_codec, task_payload, `status`, publish_attempts, next_retry_at, published_topic, published_partition, published_offset, last_error_code, created_at, updated_at, published_at) values (:task_id, :user_id, :pts, :push_type, :peer_type, :peer_id, :operation_id, :push_partition_id, :task_schema_version, :task_codec, :task_payload, :status, :publish_attempts, :next_retry_at, :published_topic, :published_partition, :published_offset, :last_error_code, NOW(6), NOW(6), :published_at)"
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

// InsertTx
// insert into push_task_outbox(task_id, user_id, pts, push_type, peer_type, peer_id, operation_id, push_partition_id, task_schema_version, task_codec, task_payload, `status`, publish_attempts, next_retry_at, published_topic, published_partition, published_offset, last_error_code, created_at, updated_at, published_at) values (:task_id, :user_id, :pts, :push_type, :peer_type, :peer_id, :operation_id, :push_partition_id, :task_schema_version, :task_codec, :task_payload, :status, :publish_attempts, :next_retry_at, :published_topic, :published_partition, :published_offset, :last_error_code, NOW(6), NOW(6), :published_at)
func (m *defaultPushTaskOutboxModel) InsertTx(tx *sqlx.Tx, data *PushTaskOutbox) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into push_task_outbox(task_id, user_id, pts, push_type, peer_type, peer_id, operation_id, push_partition_id, task_schema_version, task_codec, task_payload, `status`, publish_attempts, next_retry_at, published_topic, published_partition, published_offset, last_error_code, created_at, updated_at, published_at) values (:task_id, :user_id, :pts, :push_type, :peer_type, :peer_id, :operation_id, :push_partition_id, :task_schema_version, :task_codec, :task_payload, :status, :publish_attempts, :next_retry_at, :published_topic, :published_partition, :published_offset, :last_error_code, NOW(6), NOW(6), :published_at)"
		r     sql.Result
	)

	r, err = tx.NamedExec(query, data)
	if err != nil {
		err = fmt.Errorf("push_task_outbox.InsertTx named exec: %w", err)
		return
	}

	lastInsertId, err = r.LastInsertId()
	if err != nil {
		err = fmt.Errorf("push_task_outbox.InsertTx last insert id: %w", err)
		return
	}
	rowsAffected, err = r.RowsAffected()
	if err != nil {
		err = fmt.Errorf("push_task_outbox.InsertTx rows affected: %w", err)
	}

	return
}

// SelectPending
// select task_id, user_id, pts, push_type, peer_type, peer_id, operation_id, push_partition_id, task_schema_version, task_codec, task_payload, `status`, publish_attempts, next_retry_at, published_topic, published_partition, published_offset, last_error_code, created_at, updated_at, published_at from push_task_outbox where `status` = :status and (next_retry_at is null or next_retry_at <= :next_retry_at) order by created_at asc limit :limit
func (m *defaultPushTaskOutboxModel) SelectPending(ctx context.Context, status int32, nextRetryAt string, limit int32) (rList []PushTaskOutbox, err error) {
	var (
		query  = "select task_id, user_id, pts, push_type, peer_type, peer_id, operation_id, push_partition_id, task_schema_version, task_codec, task_payload, `status`, publish_attempts, next_retry_at, published_topic, published_partition, published_offset, last_error_code, created_at, updated_at, published_at from push_task_outbox where `status` = ? and (next_retry_at is null or next_retry_at <= ?) order by created_at asc limit ?"
		values []PushTaskOutbox
	)
	err = m.db.QueryRowsPartial(ctx, &values, query, status, nextRetryAt, limit)

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
// select task_id, user_id, pts, push_type, peer_type, peer_id, operation_id, push_partition_id, task_schema_version, task_codec, task_payload, `status`, publish_attempts, next_retry_at, published_topic, published_partition, published_offset, last_error_code, created_at, updated_at, published_at from push_task_outbox where `status` = :status and (next_retry_at is null or next_retry_at <= :next_retry_at) order by created_at asc limit :limit
func (m *defaultPushTaskOutboxModel) SelectPendingWithCB(ctx context.Context, status int32, nextRetryAt string, limit int32, cb func(sz, i int, v *PushTaskOutbox)) (rList []PushTaskOutbox, err error) {
	var (
		query  = "select task_id, user_id, pts, push_type, peer_type, peer_id, operation_id, push_partition_id, task_schema_version, task_codec, task_payload, `status`, publish_attempts, next_retry_at, published_topic, published_partition, published_offset, last_error_code, created_at, updated_at, published_at from push_task_outbox where `status` = ? and (next_retry_at is null or next_retry_at <= ?) order by created_at asc limit ?"
		values []PushTaskOutbox
	)
	err = m.db.QueryRowsPartial(ctx, &values, query, status, nextRetryAt, limit)

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

// MarkPublishing
// update push_task_outbox set `status` = :status, publish_attempts = publish_attempts + 1, updated_at = NOW(6) where task_id = :task_id
func (m *defaultPushTaskOutboxModel) MarkPublishing(ctx context.Context, status int32, taskId int64) (rowsAffected int64, err error) {

	var (
		query   = "update push_task_outbox set `status` = ?, publish_attempts = publish_attempts + 1, updated_at = NOW(6) where task_id = ?"
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

// MarkPublishingTx
// update push_task_outbox set `status` = :status, publish_attempts = publish_attempts + 1, updated_at = NOW(6) where task_id = :task_id
func (m *defaultPushTaskOutboxModel) MarkPublishingTx(tx *sqlx.Tx, status int32, taskId int64) (rowsAffected int64, err error) {
	var (
		query   = "update push_task_outbox set `status` = ?, publish_attempts = publish_attempts + 1, updated_at = NOW(6) where task_id = ?"
		rResult sql.Result
	)
	rResult, err = tx.Exec(query, status, taskId)

	if err != nil {
		err = fmt.Errorf("push_task_outbox.MarkPublishingTx exec: %w", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		err = fmt.Errorf("push_task_outbox.MarkPublishingTx rows affected: %w", err)
		return
	}

	return
}

// MarkPublished
// update push_task_outbox set `status` = :status, published_topic = :published_topic, published_partition = :published_partition, published_offset = :published_offset, published_at = :published_at, updated_at = NOW(6) where task_id = :task_id
func (m *defaultPushTaskOutboxModel) MarkPublished(ctx context.Context, status int32, publishedTopic string, publishedPartition int32, publishedOffset int64, publishedAt string, taskId int64) (rowsAffected int64, err error) {

	var (
		query   = "update push_task_outbox set `status` = ?, published_topic = ?, published_partition = ?, published_offset = ?, published_at = ?, updated_at = NOW(6) where task_id = ?"
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

// MarkPublishedTx
// update push_task_outbox set `status` = :status, published_topic = :published_topic, published_partition = :published_partition, published_offset = :published_offset, published_at = :published_at, updated_at = NOW(6) where task_id = :task_id
func (m *defaultPushTaskOutboxModel) MarkPublishedTx(tx *sqlx.Tx, status int32, publishedTopic string, publishedPartition int32, publishedOffset int64, publishedAt string, taskId int64) (rowsAffected int64, err error) {
	var (
		query   = "update push_task_outbox set `status` = ?, published_topic = ?, published_partition = ?, published_offset = ?, published_at = ?, updated_at = NOW(6) where task_id = ?"
		rResult sql.Result
	)
	rResult, err = tx.Exec(query, status, publishedTopic, publishedPartition, publishedOffset, publishedAt, taskId)

	if err != nil {
		err = fmt.Errorf("push_task_outbox.MarkPublishedTx exec: %w", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		err = fmt.Errorf("push_task_outbox.MarkPublishedTx rows affected: %w", err)
		return
	}

	return
}

// MarkPublishFailed
// update push_task_outbox set `status` = :status, next_retry_at = :next_retry_at, last_error_code = :last_error_code, updated_at = NOW(6) where task_id = :task_id
func (m *defaultPushTaskOutboxModel) MarkPublishFailed(ctx context.Context, status int32, nextRetryAt string, lastErrorCode string, taskId int64) (rowsAffected int64, err error) {

	var (
		query   = "update push_task_outbox set `status` = ?, next_retry_at = ?, last_error_code = ?, updated_at = NOW(6) where task_id = ?"
		rResult sql.Result
	)

	rResult, err = m.db.Exec(ctx, query, status, nextRetryAt, lastErrorCode, taskId)

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

// MarkPublishFailedTx
// update push_task_outbox set `status` = :status, next_retry_at = :next_retry_at, last_error_code = :last_error_code, updated_at = NOW(6) where task_id = :task_id
func (m *defaultPushTaskOutboxModel) MarkPublishFailedTx(tx *sqlx.Tx, status int32, nextRetryAt string, lastErrorCode string, taskId int64) (rowsAffected int64, err error) {
	var (
		query   = "update push_task_outbox set `status` = ?, next_retry_at = ?, last_error_code = ?, updated_at = NOW(6) where task_id = ?"
		rResult sql.Result
	)
	rResult, err = tx.Exec(query, status, nextRetryAt, lastErrorCode, taskId)

	if err != nil {
		err = fmt.Errorf("push_task_outbox.MarkPublishFailedTx exec: %w", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		err = fmt.Errorf("push_task_outbox.MarkPublishFailedTx rows affected: %w", err)
		return
	}

	return
}
