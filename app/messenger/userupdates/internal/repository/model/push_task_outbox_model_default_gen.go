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

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	pushTaskOutboxFieldNames          = builder.RawFieldNames(&PushTaskOutbox{})
	pushTaskOutboxRows                = strings.Join(pushTaskOutboxFieldNames, ",")
	pushTaskOutboxRowsExpectAutoSet   = strings.Join(stringx.Remove(pushTaskOutboxFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	pushTaskOutboxRowsWithPlaceHolder = strings.Join(stringx.Remove(pushTaskOutboxFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"
)

type (
	pushTaskOutboxModel interface {
		Insert2(ctx context.Context, data *PushTaskOutbox) (sql.Result, error)
		FindOne(ctx context.Context, taskId int64) (*PushTaskOutbox, error)
		FindListByTaskIdList(ctx context.Context, taskId ...int64) ([]PushTaskOutbox, error)
		Update2(ctx context.Context, data *PushTaskOutbox) error
		Delete2(ctx context.Context, taskId int64) error

		FindOneByUserIdPtsPushType(ctx context.Context, userId int64, pts int64, pushType int32) (*PushTaskOutbox, error)
	}

	defaultPushTaskOutboxModel struct {
		db *sqlx.DB
	}

	PushTaskOutbox struct {
		TaskId             int64        `db:"task_id" json:"task_id"`
		UserId             int64        `db:"user_id" json:"user_id"`
		Pts                int64        `db:"pts" json:"pts"`
		PushType           int32        `db:"push_type" json:"push_type"`
		PeerType           int32        `db:"peer_type" json:"peer_type"`
		PeerId             int64        `db:"peer_id" json:"peer_id"`
		OperationId        string       `db:"operation_id" json:"operation_id"`
		PushPartitionId    int32        `db:"push_partition_id" json:"push_partition_id"`
		TaskSchemaVersion  int32        `db:"task_schema_version" json:"task_schema_version"`
		TaskCodec          int32        `db:"task_codec" json:"task_codec"`
		TaskPayload        []byte       `db:"task_payload" json:"task_payload"`
		Status             int32        `db:"status" json:"status"`
		PublishAttempts    int32        `db:"publish_attempts" json:"publish_attempts"`
		AvailableAt        time.Time    `db:"available_at" json:"available_at"`
		NextRetryAt        sql.NullTime `db:"next_retry_at" json:"next_retry_at"`
		PublishedTopic     string       `db:"published_topic" json:"published_topic"`
		PublishedPartition int32        `db:"published_partition" json:"published_partition"`
		PublishedOffset    int64        `db:"published_offset" json:"published_offset"`
		LastErrorCode      string       `db:"last_error_code" json:"last_error_code"`
		PublishedAt        sql.NullTime `db:"published_at" json:"published_at"`
	}
)

func newPushTaskOutboxModel(db *sqlx.DB) *defaultPushTaskOutboxModel {
	return &defaultPushTaskOutboxModel{
		db: db,
	}
}

func (m *defaultPushTaskOutboxModel) Insert2(ctx context.Context, data *PushTaskOutbox) (sql.Result, error) {
	tableName := "push_task_outbox"
	query := fmt.Sprintf("insert into `%s` (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", tableName, pushTaskOutboxRowsExpectAutoSet)

	r, err := m.db.Exec(ctx, query, data.UserId, data.Pts, data.PushType, data.PeerType, data.PeerId, data.OperationId, data.PushPartitionId, data.TaskSchemaVersion, data.TaskCodec, data.TaskPayload, data.Status, data.PublishAttempts, data.AvailableAt, data.NextRetryAt, data.PublishedTopic, data.PublishedPartition, data.PublishedOffset, data.LastErrorCode, data.PublishedAt)
	if err != nil {
		return nil, fmt.Errorf("push_task_outbox.Insert2 exec: %w", err)
	}

	return r, nil
}

func (m *defaultPushTaskOutboxModel) Delete2(ctx context.Context, taskId int64) error {
	tableName := "push_task_outbox"
	query := fmt.Sprintf("delete from `%s` where `task_id` = ?", tableName)

	_, err := m.db.Exec(ctx, query, taskId)
	if err != nil {
		return fmt.Errorf("push_task_outbox.Delete2 exec: %w", err)
	}

	return nil
}

func (m *defaultPushTaskOutboxModel) FindOne(ctx context.Context, taskId int64) (*PushTaskOutbox, error) {
	tableName := "push_task_outbox"
	query := fmt.Sprintf("select %s from %s where task_id = ? limit 1", pushTaskOutboxRows, tableName)
	var resp PushTaskOutbox

	err := m.db.QueryRowPartial(ctx, &resp, query, taskId)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return nil, &NotFoundError{
				Resource: "push_task_outbox",
				Key:      fmt.Sprintf("task_id=%v", taskId),
				Cause:    err,
			}
		}
		return nil, fmt.Errorf("push_task_outbox.FindOne: %w", err)
	}

	return &resp, nil
}

func (m *defaultPushTaskOutboxModel) FindListByTaskIdList(ctx context.Context, taskId ...int64) ([]PushTaskOutbox, error) {
	if len(taskId) == 0 {
		return []PushTaskOutbox{}, nil
	}
	tableName := "push_task_outbox"

	query := fmt.Sprintf("select %s from %s where task_id in (%s)", pushTaskOutboxRows, tableName, sqlx.InInt64List(taskId))

	var resp []PushTaskOutbox
	err := m.db.QueryRowsPartial(ctx, &resp, query)
	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return []PushTaskOutbox{}, nil
		}
		return nil, fmt.Errorf("push_task_outbox.FindListByTaskIdList: %w", err)
	}

	return resp, nil
}

func (m *defaultPushTaskOutboxModel) Update2(ctx context.Context, data *PushTaskOutbox) error {
	tableName := "push_task_outbox"
	query := fmt.Sprintf("update `%s` set %s where `task_id` = ?", tableName, pushTaskOutboxRowsWithPlaceHolder)

	_, err := m.db.Exec(ctx, query, data.UserId, data.Pts, data.PushType, data.PeerType, data.PeerId, data.OperationId, data.PushPartitionId, data.TaskSchemaVersion, data.TaskCodec, data.TaskPayload, data.Status, data.PublishAttempts, data.AvailableAt, data.NextRetryAt, data.PublishedTopic, data.PublishedPartition, data.PublishedOffset, data.LastErrorCode, data.PublishedAt, data.TaskId)
	if err != nil {
		return fmt.Errorf("push_task_outbox.Update2 exec: %w", err)
	}

	return nil
}

func (m *defaultPushTaskOutboxModel) FindOneByUserIdPtsPushType(ctx context.Context, userId int64, pts int64, pushType int32) (*PushTaskOutbox, error) {
	tableName := "push_task_outbox"
	query := fmt.Sprintf("select %s from %s where user_id = ? AND pts = ? AND push_type = ? limit 1", pushTaskOutboxRows, tableName)
	var resp PushTaskOutbox

	err := m.db.QueryRowPartial(ctx, &resp, query, userId, pts, pushType)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return nil, &NotFoundError{
				Resource: "push_task_outbox",
				Key:      fmt.Sprintf("user_id=%v,pts=%v,push_type=%v", userId, pts, pushType),
				Cause:    err,
			}
		}
		return nil, fmt.Errorf("push_task_outbox.FindOneByUserIdPtsPushType: %w", err)
	}

	return &resp, nil
}
