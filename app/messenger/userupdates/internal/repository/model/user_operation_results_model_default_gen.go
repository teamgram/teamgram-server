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
	"fmt"
	"strings"
	"time"

	"github.com/teamgram/marmota/pkg/stores/sqlx"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	userOperationResultsFieldNames          = builder.RawFieldNames(&UserOperationResults{})
	userOperationResultsRows                = strings.Join(userOperationResultsFieldNames, ",")
	userOperationResultsRowsExpectAutoSet   = strings.Join(stringx.Remove(userOperationResultsFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	userOperationResultsRowsWithPlaceHolder = strings.Join(stringx.Remove(userOperationResultsFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"
)

type (
	userOperationResultsModel interface {
		Insert2(ctx context.Context, data *UserOperationResults) (sql.Result, error)
	}

	defaultUserOperationResultsModel struct {
		db *sqlx.DB
	}

	UserOperationResults struct {
		UserId                int64     `db:"user_id" json:"user_id"`
		OperationId           string    `db:"operation_id" json:"operation_id"`
		OpType                int32     `db:"op_type" json:"op_type"`
		Status                int32     `db:"status" json:"status"`
		Pts                   int64     `db:"pts" json:"pts"`
		PtsCount              int32     `db:"pts_count" json:"pts_count"`
		PayloadHash           []byte    `db:"payload_hash" json:"payload_hash"`
		ResponseSchemaVersion int32     `db:"response_schema_version" json:"response_schema_version"`
		ResponseCodec         int32     `db:"response_codec" json:"response_codec"`
		ResponsePayload       []byte    `db:"response_payload" json:"response_payload"`
		ResponsePayloadHash   []byte    `db:"response_payload_hash" json:"response_payload_hash"`
		TerminalErrorCode     string    `db:"terminal_error_code" json:"terminal_error_code"`
		CompletedAt           time.Time `db:"completed_at" json:"completed_at"`
	}
)

func newUserOperationResultsModel(db *sqlx.DB) *defaultUserOperationResultsModel {
	return &defaultUserOperationResultsModel{
		db: db,
	}
}

func (m *defaultUserOperationResultsModel) Insert2(ctx context.Context, data *UserOperationResults) (sql.Result, error) {
	tableName := "user_operation_results"
	query := fmt.Sprintf("insert into `%s` (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", tableName, userOperationResultsRowsExpectAutoSet)

	r, err := m.db.Exec(ctx, query, data.UserId, data.OperationId, data.OpType, data.Status, data.Pts, data.PtsCount, data.PayloadHash, data.ResponseSchemaVersion, data.ResponseCodec, data.ResponsePayload, data.ResponsePayloadHash, data.TerminalErrorCode, data.CompletedAt)
	if err != nil {
		return nil, fmt.Errorf("user_operation_results.Insert2 exec: %w", err)
	}

	return r, nil
}
