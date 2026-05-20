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

	"github.com/teamgram/marmota/pkg/stores/sqlx"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	authSeqStateFieldNames          = builder.RawFieldNames(&AuthSeqState{})
	authSeqStateRows                = strings.Join(authSeqStateFieldNames, ",")
	authSeqStateRowsExpectAutoSet   = strings.Join(stringx.Remove(authSeqStateFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	authSeqStateRowsWithPlaceHolder = strings.Join(stringx.Remove(authSeqStateFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"
)

type (
	authSeqStateModel interface {
		Insert2(ctx context.Context, data *AuthSeqState) (sql.Result, error)
	}

	defaultAuthSeqStateModel struct {
		db *sqlx.DB
	}

	AuthSeqState struct {
		UserId        int64 `db:"user_id" json:"user_id"`
		PermAuthKeyId int64 `db:"perm_auth_key_id" json:"perm_auth_key_id"`
		Seq           int64 `db:"seq" json:"seq"`
		Date          int64 `db:"date" json:"date"`
		RowVersion    int64 `db:"row_version" json:"row_version"`
	}
)

func newAuthSeqStateModel(db *sqlx.DB) *defaultAuthSeqStateModel {
	return &defaultAuthSeqStateModel{
		db: db,
	}
}

func (m *defaultAuthSeqStateModel) Insert2(ctx context.Context, data *AuthSeqState) (sql.Result, error) {
	tableName := "auth_seq_state"
	query := fmt.Sprintf("insert into `%s` (%s) values (?, ?, ?, ?, ?)", tableName, authSeqStateRowsExpectAutoSet)

	r, err := m.db.Exec(ctx, query, data.UserId, data.PermAuthKeyId, data.Seq, data.Date, data.RowVersion)
	if err != nil {
		return nil, fmt.Errorf("auth_seq_state.Insert2 exec: %w", err)
	}

	return r, nil
}
