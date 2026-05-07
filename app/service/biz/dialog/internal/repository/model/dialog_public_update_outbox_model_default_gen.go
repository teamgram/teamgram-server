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
	dialogPublicUpdateOutboxFieldNames          = builder.RawFieldNames(&DialogPublicUpdateOutbox{})
	dialogPublicUpdateOutboxRows                = strings.Join(dialogPublicUpdateOutboxFieldNames, ",")
	dialogPublicUpdateOutboxRowsExpectAutoSet   = strings.Join(stringx.Remove(dialogPublicUpdateOutboxFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	dialogPublicUpdateOutboxRowsWithPlaceHolder = strings.Join(stringx.Remove(dialogPublicUpdateOutboxFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"
)

type (
	dialogPublicUpdateOutboxModel interface {
		Insert2(ctx context.Context, data *DialogPublicUpdateOutbox) (sql.Result, error)
		FindOne(ctx context.Context, outboxId int64) (*DialogPublicUpdateOutbox, error)
		FindListByOutboxIdList(ctx context.Context, outboxId ...int64) ([]DialogPublicUpdateOutbox, error)
		Update2(ctx context.Context, data *DialogPublicUpdateOutbox) error
		Delete2(ctx context.Context, outboxId int64) error

		FindOneByTargetUserIdOperationIdDeliveryPathPublicUpdateType(ctx context.Context, targetUserId int64, operationId string, deliveryPath string, publicUpdateType string) (*DialogPublicUpdateOutbox, error)
	}

	defaultDialogPublicUpdateOutboxModel struct {
		db *sqlx.DB
	}

	DialogPublicUpdateOutbox struct {
		OutboxId             int64  `db:"outbox_id" json:"outbox_id"`
		SourceUserId         int64  `db:"source_user_id" json:"source_user_id"`
		SourcePermAuthKeyId  int64  `db:"source_perm_auth_key_id" json:"source_perm_auth_key_id"`
		TargetUserId         int64  `db:"target_user_id" json:"target_user_id"`
		TargetAuthPolicy     string `db:"target_auth_policy" json:"target_auth_policy"`
		OperationId          string `db:"operation_id" json:"operation_id"`
		DeliveryPath         string `db:"delivery_path" json:"delivery_path"`
		PublicUpdateType     string `db:"public_update_type" json:"public_update_type"`
		PeerType             int32  `db:"peer_type" json:"peer_type"`
		PeerId               int64  `db:"peer_id" json:"peer_id"`
		PayloadSchemaVersion int32  `db:"payload_schema_version" json:"payload_schema_version"`
		Payload              []byte `db:"payload" json:"payload"`
		PayloadHash          []byte `db:"payload_hash" json:"payload_hash"`
		Status               int32  `db:"status" json:"status"`
		AttemptCount         int32  `db:"attempt_count" json:"attempt_count"`
		NextRetryAt          int64  `db:"next_retry_at" json:"next_retry_at"`
		LeaseOwner           string `db:"lease_owner" json:"lease_owner"`
		LeaseUntil           int64  `db:"lease_until" json:"lease_until"`
		PublishedPts         int64  `db:"published_pts" json:"published_pts"`
		PublishedPtsCount    int32  `db:"published_pts_count" json:"published_pts_count"`
		PublishedSeq         int64  `db:"published_seq" json:"published_seq"`
		PublishedDate        int32  `db:"published_date" json:"published_date"`
		LastErrorKind        string `db:"last_error_kind" json:"last_error_kind"`
		LastErrorMessage     string `db:"last_error_message" json:"last_error_message"`
	}
)

func newDialogPublicUpdateOutboxModel(db *sqlx.DB) *defaultDialogPublicUpdateOutboxModel {
	return &defaultDialogPublicUpdateOutboxModel{
		db: db,
	}
}

func (m *defaultDialogPublicUpdateOutboxModel) Insert2(ctx context.Context, data *DialogPublicUpdateOutbox) (sql.Result, error) {
	tableName := "dialog_public_update_outbox"
	query := fmt.Sprintf("insert into `%s` (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", tableName, dialogPublicUpdateOutboxRowsExpectAutoSet)

	r, err := m.db.Exec(ctx, query, data.SourceUserId, data.SourcePermAuthKeyId, data.TargetUserId, data.TargetAuthPolicy, data.OperationId, data.DeliveryPath, data.PublicUpdateType, data.PeerType, data.PeerId, data.PayloadSchemaVersion, data.Payload, data.PayloadHash, data.Status, data.AttemptCount, data.NextRetryAt, data.LeaseOwner, data.LeaseUntil, data.PublishedPts, data.PublishedPtsCount, data.PublishedSeq, data.PublishedDate, data.LastErrorKind, data.LastErrorMessage)
	if err != nil {
		return nil, fmt.Errorf("dialog_public_update_outbox.Insert2 exec: %w", err)
	}

	return r, nil
}

func (m *defaultDialogPublicUpdateOutboxModel) Delete2(ctx context.Context, outboxId int64) error {
	tableName := "dialog_public_update_outbox"
	query := fmt.Sprintf("delete from `%s` where `outbox_id` = ?", tableName)

	_, err := m.db.Exec(ctx, query, outboxId)
	if err != nil {
		return fmt.Errorf("dialog_public_update_outbox.Delete2 exec: %w", err)
	}

	return nil
}

func (m *defaultDialogPublicUpdateOutboxModel) FindOne(ctx context.Context, outboxId int64) (*DialogPublicUpdateOutbox, error) {
	tableName := "dialog_public_update_outbox"
	query := fmt.Sprintf("select %s from %s where outbox_id = ? limit 1", dialogPublicUpdateOutboxRows, tableName)
	var resp DialogPublicUpdateOutbox

	err := m.db.QueryRowPartial(ctx, &resp, query, outboxId)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return nil, &NotFoundError{
				Resource: "dialog_public_update_outbox",
				Key:      fmt.Sprintf("outbox_id=%v", outboxId),
				Cause:    err,
			}
		}
		return nil, fmt.Errorf("dialog_public_update_outbox.FindOne: %w", err)
	}

	return &resp, nil
}

func (m *defaultDialogPublicUpdateOutboxModel) FindListByOutboxIdList(ctx context.Context, outboxId ...int64) ([]DialogPublicUpdateOutbox, error) {
	if len(outboxId) == 0 {
		return []DialogPublicUpdateOutbox{}, nil
	}
	tableName := "dialog_public_update_outbox"

	query := fmt.Sprintf("select %s from %s where outbox_id in (%s)", dialogPublicUpdateOutboxRows, tableName, sqlx.InInt64List(outboxId))

	var resp []DialogPublicUpdateOutbox
	err := m.db.QueryRowsPartial(ctx, &resp, query)
	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return []DialogPublicUpdateOutbox{}, nil
		}
		return nil, fmt.Errorf("dialog_public_update_outbox.FindListByOutboxIdList: %w", err)
	}

	return resp, nil
}

func (m *defaultDialogPublicUpdateOutboxModel) Update2(ctx context.Context, data *DialogPublicUpdateOutbox) error {
	tableName := "dialog_public_update_outbox"
	query := fmt.Sprintf("update `%s` set %s where `outbox_id` = ?", tableName, dialogPublicUpdateOutboxRowsWithPlaceHolder)

	_, err := m.db.Exec(ctx, query, data.SourceUserId, data.SourcePermAuthKeyId, data.TargetUserId, data.TargetAuthPolicy, data.OperationId, data.DeliveryPath, data.PublicUpdateType, data.PeerType, data.PeerId, data.PayloadSchemaVersion, data.Payload, data.PayloadHash, data.Status, data.AttemptCount, data.NextRetryAt, data.LeaseOwner, data.LeaseUntil, data.PublishedPts, data.PublishedPtsCount, data.PublishedSeq, data.PublishedDate, data.LastErrorKind, data.LastErrorMessage, data.OutboxId)
	if err != nil {
		return fmt.Errorf("dialog_public_update_outbox.Update2 exec: %w", err)
	}

	return nil
}

func (m *defaultDialogPublicUpdateOutboxModel) FindOneByTargetUserIdOperationIdDeliveryPathPublicUpdateType(ctx context.Context, targetUserId int64, operationId string, deliveryPath string, publicUpdateType string) (*DialogPublicUpdateOutbox, error) {
	tableName := "dialog_public_update_outbox"
	query := fmt.Sprintf("select %s from %s where target_user_id = ? AND operation_id = ? AND delivery_path = ? AND public_update_type = ? limit 1", dialogPublicUpdateOutboxRows, tableName)
	var resp DialogPublicUpdateOutbox

	err := m.db.QueryRowPartial(ctx, &resp, query, targetUserId, operationId, deliveryPath, publicUpdateType)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return nil, &NotFoundError{
				Resource: "dialog_public_update_outbox",
				Key:      fmt.Sprintf("target_user_id=%v,operation_id=%v,delivery_path=%v,public_update_type=%v", targetUserId, operationId, deliveryPath, publicUpdateType),
				Cause:    err,
			}
		}
		return nil, fmt.Errorf("dialog_public_update_outbox.FindOneByTargetUserIdOperationIdDeliveryPathPublicUpdateType: %w", err)
	}

	return &resp, nil
}
