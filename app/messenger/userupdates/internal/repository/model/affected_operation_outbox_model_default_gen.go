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

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	affectedOperationOutboxFieldNames          = builder.RawFieldNames(&AffectedOperationOutbox{})
	affectedOperationOutboxRows                = strings.Join(affectedOperationOutboxFieldNames, ",")
	affectedOperationOutboxRowsExpectAutoSet   = strings.Join(stringx.Remove(affectedOperationOutboxFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	affectedOperationOutboxRowsWithPlaceHolder = strings.Join(stringx.Remove(affectedOperationOutboxFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"
)

type (
	affectedOperationOutboxModel interface {
		Insert2(ctx context.Context, data *AffectedOperationOutbox) (sql.Result, error)
		FindOne(ctx context.Context, outboxId int64) (*AffectedOperationOutbox, error)
		FindListByOutboxIdList(ctx context.Context, outboxId ...int64) ([]AffectedOperationOutbox, error)
		Update2(ctx context.Context, data *AffectedOperationOutbox) error
		Delete2(ctx context.Context, outboxId int64) error

		FindOneByUserIdOperationId(ctx context.Context, userId int64, operationId string) (*AffectedOperationOutbox, error)
	}

	defaultAffectedOperationOutboxModel struct {
		db *sqlx.DB
	}

	AffectedOperationOutbox struct {
		OutboxId           int64  `db:"outbox_id" json:"outbox_id"`
		UserId             int64  `db:"user_id" json:"user_id"`
		RequesterUserId    int64  `db:"requester_user_id" json:"requester_user_id"`
		OperationId        string `db:"operation_id" json:"operation_id"`
		OpType             int32  `db:"op_type" json:"op_type"`
		OperationKind      string `db:"operation_kind" json:"operation_kind"`
		PeerType           int32  `db:"peer_type" json:"peer_type"`
		PeerId             int64  `db:"peer_id" json:"peer_id"`
		PayloadCodec       int32  `db:"payload_codec" json:"payload_codec"`
		PayloadHash        []byte `db:"payload_hash" json:"payload_hash"`
		Payload            []byte `db:"payload" json:"payload"`
		DeliveryPolicy     int32  `db:"delivery_policy" json:"delivery_policy"`
		Status             int32  `db:"status" json:"status"`
		RetryCount         int32  `db:"retry_count" json:"retry_count"`
		AvailableAt        int64  `db:"available_at" json:"available_at"`
		ProcessingDeadline int64  `db:"processing_deadline" json:"processing_deadline"`
		LastErrorCode      string `db:"last_error_code" json:"last_error_code"`
		LastErrorMessage   string `db:"last_error_message" json:"last_error_message"`
		BucketId           int32  `db:"bucket_id" json:"bucket_id"`
		PartitionId        int32  `db:"partition_id" json:"partition_id"`
		OwnerTokenPayload  []byte `db:"owner_token_payload" json:"owner_token_payload"`
	}
)

func newAffectedOperationOutboxModel(db *sqlx.DB) *defaultAffectedOperationOutboxModel {
	return &defaultAffectedOperationOutboxModel{
		db: db,
	}
}

func (m *defaultAffectedOperationOutboxModel) Insert2(ctx context.Context, data *AffectedOperationOutbox) (sql.Result, error) {
	tableName := "affected_operation_outbox"
	query := fmt.Sprintf("insert into `%s` (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", tableName, affectedOperationOutboxRowsExpectAutoSet)

	r, err := m.db.Exec(ctx, query, data.UserId, data.RequesterUserId, data.OperationId, data.OpType, data.OperationKind, data.PeerType, data.PeerId, data.PayloadCodec, data.PayloadHash, data.Payload, data.DeliveryPolicy, data.Status, data.RetryCount, data.AvailableAt, data.ProcessingDeadline, data.LastErrorCode, data.LastErrorMessage, data.BucketId, data.PartitionId, data.OwnerTokenPayload)
	if err != nil {
		return nil, fmt.Errorf("affected_operation_outbox.Insert2 exec: %w", err)
	}

	return r, nil
}

func (m *defaultAffectedOperationOutboxModel) Delete2(ctx context.Context, outboxId int64) error {
	tableName := "affected_operation_outbox"
	query := fmt.Sprintf("delete from `%s` where `outbox_id` = ?", tableName)

	_, err := m.db.Exec(ctx, query, outboxId)
	if err != nil {
		return fmt.Errorf("affected_operation_outbox.Delete2 exec: %w", err)
	}

	return nil
}

func (m *defaultAffectedOperationOutboxModel) FindOne(ctx context.Context, outboxId int64) (*AffectedOperationOutbox, error) {
	tableName := "affected_operation_outbox"
	query := fmt.Sprintf("select %s from %s where outbox_id = ? limit 1", affectedOperationOutboxRows, tableName)
	var resp AffectedOperationOutbox

	err := m.db.QueryRowPartial(ctx, &resp, query, outboxId)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return nil, &NotFoundError{
				Resource: "affected_operation_outbox",
				Key:      fmt.Sprintf("outbox_id=%v", outboxId),
				Cause:    err,
			}
		}
		return nil, fmt.Errorf("affected_operation_outbox.FindOne: %w", err)
	}

	return &resp, nil
}

func (m *defaultAffectedOperationOutboxModel) FindListByOutboxIdList(ctx context.Context, outboxId ...int64) ([]AffectedOperationOutbox, error) {
	if len(outboxId) == 0 {
		return []AffectedOperationOutbox{}, nil
	}
	tableName := "affected_operation_outbox"

	query := fmt.Sprintf("select %s from %s where outbox_id in (%s)", affectedOperationOutboxRows, tableName, sqlx.InInt64List(outboxId))

	var resp []AffectedOperationOutbox
	err := m.db.QueryRowsPartial(ctx, &resp, query)
	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return []AffectedOperationOutbox{}, nil
		}
		return nil, fmt.Errorf("affected_operation_outbox.FindListByOutboxIdList: %w", err)
	}

	return resp, nil
}

func (m *defaultAffectedOperationOutboxModel) Update2(ctx context.Context, data *AffectedOperationOutbox) error {
	tableName := "affected_operation_outbox"
	query := fmt.Sprintf("update `%s` set %s where `outbox_id` = ?", tableName, affectedOperationOutboxRowsWithPlaceHolder)

	_, err := m.db.Exec(ctx, query, data.UserId, data.RequesterUserId, data.OperationId, data.OpType, data.OperationKind, data.PeerType, data.PeerId, data.PayloadCodec, data.PayloadHash, data.Payload, data.DeliveryPolicy, data.Status, data.RetryCount, data.AvailableAt, data.ProcessingDeadline, data.LastErrorCode, data.LastErrorMessage, data.BucketId, data.PartitionId, data.OwnerTokenPayload, data.OutboxId)
	if err != nil {
		return fmt.Errorf("affected_operation_outbox.Update2 exec: %w", err)
	}

	return nil
}

func (m *defaultAffectedOperationOutboxModel) FindOneByUserIdOperationId(ctx context.Context, userId int64, operationId string) (*AffectedOperationOutbox, error) {
	tableName := "affected_operation_outbox"
	query := fmt.Sprintf("select %s from %s where user_id = ? AND operation_id = ? limit 1", affectedOperationOutboxRows, tableName)
	var resp AffectedOperationOutbox

	err := m.db.QueryRowPartial(ctx, &resp, query, userId, operationId)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return nil, &NotFoundError{
				Resource: "affected_operation_outbox",
				Key:      fmt.Sprintf("user_id=%v,operation_id=%v", userId, operationId),
				Cause:    err,
			}
		}
		return nil, fmt.Errorf("affected_operation_outbox.FindOneByUserIdOperationId: %w", err)
	}

	return &resp, nil
}
