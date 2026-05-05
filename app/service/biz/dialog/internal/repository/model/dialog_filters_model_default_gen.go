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
	dialogFiltersFieldNames          = builder.RawFieldNames(&DialogFilters{})
	dialogFiltersRows                = strings.Join(dialogFiltersFieldNames, ",")
	dialogFiltersRowsExpectAutoSet   = strings.Join(stringx.Remove(dialogFiltersFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	dialogFiltersRowsWithPlaceHolder = strings.Join(stringx.Remove(dialogFiltersFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"
)

type (
	dialogFiltersModel interface {
		Insert2(ctx context.Context, data *DialogFilters) (sql.Result, error)

		FindOneByUserIdSlug(ctx context.Context, userId int64, slug string) (*DialogFilters, error)
	}

	defaultDialogFiltersModel struct {
		db *sqlx.DB
	}

	DialogFilters struct {
		UserId              int64  `db:"user_id" json:"user_id"`
		DialogFilterId      int32  `db:"dialog_filter_id" json:"dialog_filter_id"`
		Slug                string `db:"slug" json:"slug"`
		Title               string `db:"title" json:"title"`
		OrderValue          int64  `db:"order_value" json:"order_value"`
		Enabled             bool   `db:"enabled" json:"enabled"`
		Deleted             bool   `db:"deleted" json:"deleted"`
		FilterSchemaVersion int32  `db:"filter_schema_version" json:"filter_schema_version"`
		FilterPayload       []byte `db:"filter_payload" json:"filter_payload"`
	}
)

func newDialogFiltersModel(db *sqlx.DB) *defaultDialogFiltersModel {
	return &defaultDialogFiltersModel{
		db: db,
	}
}

func (m *defaultDialogFiltersModel) Insert2(ctx context.Context, data *DialogFilters) (sql.Result, error) {
	tableName := "dialog_filters"
	query := fmt.Sprintf("insert into `%s` (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?)", tableName, dialogFiltersRowsExpectAutoSet)

	r, err := m.db.Exec(ctx, query, data.UserId, data.DialogFilterId, data.Slug, data.Title, data.OrderValue, data.Enabled, data.Deleted, data.FilterSchemaVersion, data.FilterPayload)
	if err != nil {
		return nil, fmt.Errorf("dialog_filters.Insert2 exec: %w", err)
	}

	return r, nil
}

func (m *defaultDialogFiltersModel) FindOneByUserIdSlug(ctx context.Context, userId int64, slug string) (*DialogFilters, error) {
	tableName := "dialog_filters"
	query := fmt.Sprintf("select %s from %s where user_id = ? AND slug = ? limit 1", dialogFiltersRows, tableName)
	var resp DialogFilters

	err := m.db.QueryRowPartial(ctx, &resp, query, userId, slug)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return nil, &NotFoundError{
				Resource: "dialog_filters",
				Key:      fmt.Sprintf("user_id=%v,slug=%v", userId, slug),
				Cause:    err,
			}
		}
		return nil, fmt.Errorf("dialog_filters.FindOneByUserIdSlug: %w", err)
	}

	return &resp, nil
}
