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
	dialogSideEffectOutboxFieldNames          = builder.RawFieldNames(&DialogSideEffectOutbox{})
	dialogSideEffectOutboxRows                = strings.Join(dialogSideEffectOutboxFieldNames, ",")
	dialogSideEffectOutboxRowsExpectAutoSet   = strings.Join(stringx.Remove(dialogSideEffectOutboxFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	dialogSideEffectOutboxRowsWithPlaceHolder = strings.Join(stringx.Remove(dialogSideEffectOutboxFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"
)

type (
	dialogSideEffectOutboxModel interface {
		Insert2(ctx context.Context, data *DialogSideEffectOutbox) (sql.Result, error)
		FindOne(ctx context.Context, sideEffectId int64) (*DialogSideEffectOutbox, error)
		FindListBySideEffectIdList(ctx context.Context, sideEffectId ...int64) ([]DialogSideEffectOutbox, error)
		Update2(ctx context.Context, data *DialogSideEffectOutbox) error
		Delete2(ctx context.Context, sideEffectId int64) error

		FindOneByKindSourceOperationId(ctx context.Context, kind string, sourceOperationId string) (*DialogSideEffectOutbox, error)
	}

	defaultDialogSideEffectOutboxModel struct {
		db *sqlx.DB
	}

	DialogSideEffectOutbox struct {
		SideEffectId             int64  `db:"side_effect_id" json:"side_effect_id"`
		Kind                     string `db:"kind" json:"kind"`
		UserId                   int64  `db:"user_id" json:"user_id"`
		PeerType                 int32  `db:"peer_type" json:"peer_type"`
		PeerId                   int64  `db:"peer_id" json:"peer_id"`
		SourcePermAuthKeyId      int64  `db:"source_perm_auth_key_id" json:"source_perm_auth_key_id"`
		SourceOperationId        string `db:"source_operation_id" json:"source_operation_id"`
		SourceMessageDate        int64  `db:"source_message_date" json:"source_message_date"`
		SourcePeerSeq            int64  `db:"source_peer_seq" json:"source_peer_seq"`
		SourceCanonicalMessageId int64  `db:"source_canonical_message_id" json:"source_canonical_message_id"`
		ClearBeforeDate          int64  `db:"clear_before_date" json:"clear_before_date"`
		PayloadSchemaVersion     int32  `db:"payload_schema_version" json:"payload_schema_version"`
		Payload                  []byte `db:"payload" json:"payload"`
		PayloadHash              []byte `db:"payload_hash" json:"payload_hash"`
		Status                   int32  `db:"status" json:"status"`
		AttemptCount             int32  `db:"attempt_count" json:"attempt_count"`
		NextRetryAt              int64  `db:"next_retry_at" json:"next_retry_at"`
		LeaseOwner               string `db:"lease_owner" json:"lease_owner"`
		LeaseUntil               int64  `db:"lease_until" json:"lease_until"`
		LastErrorCode            string `db:"last_error_code" json:"last_error_code"`
	}
)

func newDialogSideEffectOutboxModel(db *sqlx.DB) *defaultDialogSideEffectOutboxModel {
	return &defaultDialogSideEffectOutboxModel{
		db: db,
	}
}

func (m *defaultDialogSideEffectOutboxModel) Insert2(ctx context.Context, data *DialogSideEffectOutbox) (sql.Result, error) {
	tableName := "dialog_side_effect_outbox"
	query := fmt.Sprintf("insert into `%s` (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", tableName, dialogSideEffectOutboxRowsExpectAutoSet)

	r, err := m.db.Exec(ctx, query, data.Kind, data.UserId, data.PeerType, data.PeerId, data.SourcePermAuthKeyId, data.SourceOperationId, data.SourceMessageDate, data.SourcePeerSeq, data.SourceCanonicalMessageId, data.ClearBeforeDate, data.PayloadSchemaVersion, data.Payload, data.PayloadHash, data.Status, data.AttemptCount, data.NextRetryAt, data.LeaseOwner, data.LeaseUntil, data.LastErrorCode)
	if err != nil {
		return nil, fmt.Errorf("dialog_side_effect_outbox.Insert2 exec: %w", err)
	}

	return r, nil
}

func (m *defaultDialogSideEffectOutboxModel) Delete2(ctx context.Context, sideEffectId int64) error {
	tableName := "dialog_side_effect_outbox"
	query := fmt.Sprintf("delete from `%s` where `side_effect_id` = ?", tableName)

	_, err := m.db.Exec(ctx, query, sideEffectId)
	if err != nil {
		return fmt.Errorf("dialog_side_effect_outbox.Delete2 exec: %w", err)
	}

	return nil
}

func (m *defaultDialogSideEffectOutboxModel) FindOne(ctx context.Context, sideEffectId int64) (*DialogSideEffectOutbox, error) {
	tableName := "dialog_side_effect_outbox"
	query := fmt.Sprintf("select %s from %s where side_effect_id = ? limit 1", dialogSideEffectOutboxRows, tableName)
	var resp DialogSideEffectOutbox

	err := m.db.QueryRowPartial(ctx, &resp, query, sideEffectId)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return nil, &NotFoundError{
				Resource: "dialog_side_effect_outbox",
				Key:      fmt.Sprintf("side_effect_id=%v", sideEffectId),
				Cause:    err,
			}
		}
		return nil, fmt.Errorf("dialog_side_effect_outbox.FindOne: %w", err)
	}

	return &resp, nil
}

func (m *defaultDialogSideEffectOutboxModel) FindListBySideEffectIdList(ctx context.Context, sideEffectId ...int64) ([]DialogSideEffectOutbox, error) {
	if len(sideEffectId) == 0 {
		return []DialogSideEffectOutbox{}, nil
	}
	tableName := "dialog_side_effect_outbox"

	query := fmt.Sprintf("select %s from %s where side_effect_id in (%s)", dialogSideEffectOutboxRows, tableName, sqlx.InInt64List(sideEffectId))

	var resp []DialogSideEffectOutbox
	err := m.db.QueryRowsPartial(ctx, &resp, query)
	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return []DialogSideEffectOutbox{}, nil
		}
		return nil, fmt.Errorf("dialog_side_effect_outbox.FindListBySideEffectIdList: %w", err)
	}

	return resp, nil
}

func (m *defaultDialogSideEffectOutboxModel) Update2(ctx context.Context, data *DialogSideEffectOutbox) error {
	tableName := "dialog_side_effect_outbox"
	query := fmt.Sprintf("update `%s` set %s where `side_effect_id` = ?", tableName, dialogSideEffectOutboxRowsWithPlaceHolder)

	_, err := m.db.Exec(ctx, query, data.Kind, data.UserId, data.PeerType, data.PeerId, data.SourcePermAuthKeyId, data.SourceOperationId, data.SourceMessageDate, data.SourcePeerSeq, data.SourceCanonicalMessageId, data.ClearBeforeDate, data.PayloadSchemaVersion, data.Payload, data.PayloadHash, data.Status, data.AttemptCount, data.NextRetryAt, data.LeaseOwner, data.LeaseUntil, data.LastErrorCode, data.SideEffectId)
	if err != nil {
		return fmt.Errorf("dialog_side_effect_outbox.Update2 exec: %w", err)
	}

	return nil
}

func (m *defaultDialogSideEffectOutboxModel) FindOneByKindSourceOperationId(ctx context.Context, kind string, sourceOperationId string) (*DialogSideEffectOutbox, error) {
	tableName := "dialog_side_effect_outbox"
	query := fmt.Sprintf("select %s from %s where kind = ? AND source_operation_id = ? limit 1", dialogSideEffectOutboxRows, tableName)
	var resp DialogSideEffectOutbox

	err := m.db.QueryRowPartial(ctx, &resp, query, kind, sourceOperationId)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return nil, &NotFoundError{
				Resource: "dialog_side_effect_outbox",
				Key:      fmt.Sprintf("kind=%v,source_operation_id=%v", kind, sourceOperationId),
				Cause:    err,
			}
		}
		return nil, fmt.Errorf("dialog_side_effect_outbox.FindOneByKindSourceOperationId: %w", err)
	}

	return &resp, nil
}
