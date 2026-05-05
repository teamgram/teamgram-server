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
	dialogAuthSeqOutboxFieldNames          = builder.RawFieldNames(&DialogAuthSeqOutbox{})
	dialogAuthSeqOutboxRows                = strings.Join(dialogAuthSeqOutboxFieldNames, ",")
	dialogAuthSeqOutboxRowsExpectAutoSet   = strings.Join(stringx.Remove(dialogAuthSeqOutboxFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	dialogAuthSeqOutboxRowsWithPlaceHolder = strings.Join(stringx.Remove(dialogAuthSeqOutboxFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"
)

type (
	dialogAuthSeqOutboxModel interface {
		Insert2(ctx context.Context, data *DialogAuthSeqOutbox) (sql.Result, error)
		FindOne(ctx context.Context, outboxId int64) (*DialogAuthSeqOutbox, error)
		FindListByOutboxIdList(ctx context.Context, outboxId ...int64) ([]DialogAuthSeqOutbox, error)
		Update2(ctx context.Context, data *DialogAuthSeqOutbox) error
		Delete2(ctx context.Context, outboxId int64) error

		FindOneByUserIdOperationId(ctx context.Context, userId int64, operationId string) (*DialogAuthSeqOutbox, error)
	}

	defaultDialogAuthSeqOutboxModel struct {
		db *sqlx.DB
	}

	DialogAuthSeqOutbox struct {
		OutboxId             int64  `db:"outbox_id" json:"outbox_id"`
		UserId               int64  `db:"user_id" json:"user_id"`
		SourcePermAuthKeyId  int64  `db:"source_perm_auth_key_id" json:"source_perm_auth_key_id"`
		TargetAuthPolicy     string `db:"target_auth_policy" json:"target_auth_policy"`
		OperationId          string `db:"operation_id" json:"operation_id"`
		EventType            string `db:"event_type" json:"event_type"`
		PeerType             int32  `db:"peer_type" json:"peer_type"`
		PeerId               int64  `db:"peer_id" json:"peer_id"`
		PayloadSchemaVersion int32  `db:"payload_schema_version" json:"payload_schema_version"`
		Payload              []byte `db:"payload" json:"payload"`
		PayloadHash          []byte `db:"payload_hash" json:"payload_hash"`
		Status               int32  `db:"status" json:"status"`
		AttemptCount         int32  `db:"attempt_count" json:"attempt_count"`
		NextRetryAt          string `db:"next_retry_at" json:"next_retry_at"`
		LeaseOwner           string `db:"lease_owner" json:"lease_owner"`
		LeaseUntil           string `db:"lease_until" json:"lease_until"`
		PublishedSeq         int64  `db:"published_seq" json:"published_seq"`
		PublishedDate        int32  `db:"published_date" json:"published_date"`
		LastErrorKind        string `db:"last_error_kind" json:"last_error_kind"`
		LastErrorMessage     string `db:"last_error_message" json:"last_error_message"`
	}
)

func newDialogAuthSeqOutboxModel(db *sqlx.DB) *defaultDialogAuthSeqOutboxModel {
	return &defaultDialogAuthSeqOutboxModel{
		db: db,
	}
}

func (m *defaultDialogAuthSeqOutboxModel) Insert2(ctx context.Context, data *DialogAuthSeqOutbox) (sql.Result, error) {
	tableName := "dialog_auth_seq_outbox"
	query := fmt.Sprintf("insert into `%s` (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", tableName, dialogAuthSeqOutboxRowsExpectAutoSet)

	r, err := m.db.Exec(ctx, query, data.UserId, data.SourcePermAuthKeyId, data.TargetAuthPolicy, data.OperationId, data.EventType, data.PeerType, data.PeerId, data.PayloadSchemaVersion, data.Payload, data.PayloadHash, data.Status, data.AttemptCount, data.NextRetryAt, data.LeaseOwner, data.LeaseUntil, data.PublishedSeq, data.PublishedDate, data.LastErrorKind, data.LastErrorMessage)
	if err != nil {
		return nil, fmt.Errorf("dialog_auth_seq_outbox.Insert2 exec: %w", err)
	}

	return r, nil
}

func (m *defaultDialogAuthSeqOutboxModel) Delete2(ctx context.Context, outboxId int64) error {
	tableName := "dialog_auth_seq_outbox"
	query := fmt.Sprintf("delete from `%s` where `outbox_id` = ?", tableName)

	_, err := m.db.Exec(ctx, query, outboxId)
	if err != nil {
		return fmt.Errorf("dialog_auth_seq_outbox.Delete2 exec: %w", err)
	}

	return nil
}

func (m *defaultDialogAuthSeqOutboxModel) FindOne(ctx context.Context, outboxId int64) (*DialogAuthSeqOutbox, error) {
	tableName := "dialog_auth_seq_outbox"
	query := fmt.Sprintf("select %s from %s where outbox_id = ? limit 1", dialogAuthSeqOutboxRows, tableName)
	var resp DialogAuthSeqOutbox

	err := m.db.QueryRowPartial(ctx, &resp, query, outboxId)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return nil, &NotFoundError{
				Resource: "dialog_auth_seq_outbox",
				Key:      fmt.Sprintf("outbox_id=%v", outboxId),
				Cause:    err,
			}
		}
		return nil, fmt.Errorf("dialog_auth_seq_outbox.FindOne: %w", err)
	}

	return &resp, nil
}

func (m *defaultDialogAuthSeqOutboxModel) FindListByOutboxIdList(ctx context.Context, outboxId ...int64) ([]DialogAuthSeqOutbox, error) {
	if len(outboxId) == 0 {
		return []DialogAuthSeqOutbox{}, nil
	}
	tableName := "dialog_auth_seq_outbox"

	query := fmt.Sprintf("select %s from %s where outbox_id in (%s)", dialogAuthSeqOutboxRows, tableName, sqlx.InInt64List(outboxId))

	var resp []DialogAuthSeqOutbox
	err := m.db.QueryRowsPartial(ctx, &resp, query)
	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return []DialogAuthSeqOutbox{}, nil
		}
		return nil, fmt.Errorf("dialog_auth_seq_outbox.FindListByOutboxIdList: %w", err)
	}

	return resp, nil
}

func (m *defaultDialogAuthSeqOutboxModel) Update2(ctx context.Context, data *DialogAuthSeqOutbox) error {
	tableName := "dialog_auth_seq_outbox"
	query := fmt.Sprintf("update `%s` set %s where `outbox_id` = ?", tableName, dialogAuthSeqOutboxRowsWithPlaceHolder)

	_, err := m.db.Exec(ctx, query, data.UserId, data.SourcePermAuthKeyId, data.TargetAuthPolicy, data.OperationId, data.EventType, data.PeerType, data.PeerId, data.PayloadSchemaVersion, data.Payload, data.PayloadHash, data.Status, data.AttemptCount, data.NextRetryAt, data.LeaseOwner, data.LeaseUntil, data.PublishedSeq, data.PublishedDate, data.LastErrorKind, data.LastErrorMessage, data.OutboxId)
	if err != nil {
		return fmt.Errorf("dialog_auth_seq_outbox.Update2 exec: %w", err)
	}

	return nil
}

func (m *defaultDialogAuthSeqOutboxModel) FindOneByUserIdOperationId(ctx context.Context, userId int64, operationId string) (*DialogAuthSeqOutbox, error) {
	tableName := "dialog_auth_seq_outbox"
	query := fmt.Sprintf("select %s from %s where user_id = ? AND operation_id = ? limit 1", dialogAuthSeqOutboxRows, tableName)
	var resp DialogAuthSeqOutbox

	err := m.db.QueryRowPartial(ctx, &resp, query, userId, operationId)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return nil, &NotFoundError{
				Resource: "dialog_auth_seq_outbox",
				Key:      fmt.Sprintf("user_id=%v,operation_id=%v", userId, operationId),
				Cause:    err,
			}
		}
		return nil, fmt.Errorf("dialog_auth_seq_outbox.FindOneByUserIdOperationId: %w", err)
	}

	return &resp, nil
}
