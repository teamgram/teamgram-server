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
	dialogPreferenceVersionsFieldNames          = builder.RawFieldNames(&DialogPreferenceVersions{})
	dialogPreferenceVersionsRows                = strings.Join(dialogPreferenceVersionsFieldNames, ",")
	dialogPreferenceVersionsRowsExpectAutoSet   = strings.Join(stringx.Remove(dialogPreferenceVersionsFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	dialogPreferenceVersionsRowsWithPlaceHolder = strings.Join(stringx.Remove(dialogPreferenceVersionsFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"
)

type (
	dialogPreferenceVersionsModel interface {
		Insert2(ctx context.Context, data *DialogPreferenceVersions) (sql.Result, error)
	}

	defaultDialogPreferenceVersionsModel struct {
		db *sqlx.DB
	}

	DialogPreferenceVersions struct {
		UserId           int64  `db:"user_id" json:"user_id"`
		ScopeKind        string `db:"scope_kind" json:"scope_kind"`
		FolderId         int32  `db:"folder_id" json:"folder_id"`
		AggregateVersion int64  `db:"aggregate_version" json:"aggregate_version"`
	}
)

func newDialogPreferenceVersionsModel(db *sqlx.DB) *defaultDialogPreferenceVersionsModel {
	return &defaultDialogPreferenceVersionsModel{
		db: db,
	}
}

func (m *defaultDialogPreferenceVersionsModel) Insert2(ctx context.Context, data *DialogPreferenceVersions) (sql.Result, error) {
	tableName := "dialog_preference_versions"
	query := fmt.Sprintf("insert into `%s` (%s) values (?, ?, ?, ?)", tableName, dialogPreferenceVersionsRowsExpectAutoSet)

	r, err := m.db.Exec(ctx, query, data.UserId, data.ScopeKind, data.FolderId, data.AggregateVersion)
	if err != nil {
		return nil, fmt.Errorf("dialog_preference_versions.Insert2 exec: %w", err)
	}

	return r, nil
}
