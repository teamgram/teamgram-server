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
	dialogFilterTagsFieldNames          = builder.RawFieldNames(&DialogFilterTags{})
	dialogFilterTagsRows                = strings.Join(dialogFilterTagsFieldNames, ",")
	dialogFilterTagsRowsExpectAutoSet   = strings.Join(stringx.Remove(dialogFilterTagsFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	dialogFilterTagsRowsWithPlaceHolder = strings.Join(stringx.Remove(dialogFilterTagsFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"
)

type (
	dialogFilterTagsModel interface {
		Insert2(ctx context.Context, data *DialogFilterTags) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*DialogFilterTags, error)
		FindListByIdList(ctx context.Context, id ...int64) ([]DialogFilterTags, error)
		Update2(ctx context.Context, data *DialogFilterTags) error
		Delete2(ctx context.Context, id int64) error

		FindOneByUserId(ctx context.Context, userId int64) (*DialogFilterTags, error)
		FindListByUserIdList(ctx context.Context, userId ...int64) ([]DialogFilterTags, error)
	}

	defaultDialogFilterTagsModel struct {
		db *sqlx.DB
	}

	DialogFilterTags struct {
		Id      int64 `db:"id" json:"id"`
		UserId  int64 `db:"user_id" json:"user_id"`
		Enabled bool  `db:"enabled" json:"enabled"`
	}
)

func newDialogFilterTagsModel(db *sqlx.DB) *defaultDialogFilterTagsModel {
	return &defaultDialogFilterTagsModel{
		db: db,
	}
}

func (m *defaultDialogFilterTagsModel) Insert2(ctx context.Context, data *DialogFilterTags) (sql.Result, error) {
	tableName := "dialog_filter_tags"
	query := fmt.Sprintf("insert into `%s` (%s) values (?, ?)", tableName, dialogFilterTagsRowsExpectAutoSet)

	r, err := m.db.Exec(ctx, query, data.UserId, data.Enabled)
	if err != nil {
		return nil, fmt.Errorf("dialog_filter_tags.Insert2 exec: %w", err)
	}

	return r, nil
}

func (m *defaultDialogFilterTagsModel) Delete2(ctx context.Context, id int64) error {
	tableName := "dialog_filter_tags"
	query := fmt.Sprintf("delete from `%s` where `id` = ?", tableName)

	_, err := m.db.Exec(ctx, query, id)
	if err != nil {
		return fmt.Errorf("dialog_filter_tags.Delete2 exec: %w", err)
	}

	return nil
}

func (m *defaultDialogFilterTagsModel) FindOne(ctx context.Context, id int64) (*DialogFilterTags, error) {
	tableName := "dialog_filter_tags"
	query := fmt.Sprintf("select %s from %s where id = ? limit 1", dialogFilterTagsRows, tableName)
	var resp DialogFilterTags

	err := m.db.QueryRowPartial(ctx, &resp, query, id)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return nil, &NotFoundError{
				Resource: "dialog_filter_tags",
				Key:      fmt.Sprintf("id=%v", id),
				Cause:    err,
			}
		}
		return nil, fmt.Errorf("dialog_filter_tags.FindOne: %w", err)
	}

	return &resp, nil
}

func (m *defaultDialogFilterTagsModel) FindListByIdList(ctx context.Context, id ...int64) ([]DialogFilterTags, error) {
	if len(id) == 0 {
		return []DialogFilterTags{}, nil
	}
	tableName := "dialog_filter_tags"

	query := fmt.Sprintf("select %s from %s where id in (%s)", dialogFilterTagsRows, tableName, sqlx.InInt64List(id))

	var resp []DialogFilterTags
	err := m.db.QueryRowsPartial(ctx, &resp, query)
	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return []DialogFilterTags{}, nil
		}
		return nil, fmt.Errorf("dialog_filter_tags.FindListByIdList: %w", err)
	}

	return resp, nil
}

func (m *defaultDialogFilterTagsModel) Update2(ctx context.Context, data *DialogFilterTags) error {
	tableName := "dialog_filter_tags"
	query := fmt.Sprintf("update `%s` set %s where `id` = ?", tableName, dialogFilterTagsRowsWithPlaceHolder)

	_, err := m.db.Exec(ctx, query, data.UserId, data.Enabled, data.Id)
	if err != nil {
		return fmt.Errorf("dialog_filter_tags.Update2 exec: %w", err)
	}

	return nil
}

func (m *defaultDialogFilterTagsModel) FindOneByUserId(ctx context.Context, userId int64) (*DialogFilterTags, error) {
	tableName := "dialog_filter_tags"
	query := fmt.Sprintf("select %s from %s where user_id = ? limit 1", dialogFilterTagsRows, tableName)
	var resp DialogFilterTags

	err := m.db.QueryRowPartial(ctx, &resp, query, userId)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return nil, &NotFoundError{
				Resource: "dialog_filter_tags",
				Key:      fmt.Sprintf("user_id=%v", userId),
				Cause:    err,
			}
		}
		return nil, fmt.Errorf("dialog_filter_tags.FindOneByUserId: %w", err)
	}

	return &resp, nil
}

func (m *defaultDialogFilterTagsModel) FindListByUserIdList(ctx context.Context, userId ...int64) ([]DialogFilterTags, error) {
	if len(userId) == 0 {
		return []DialogFilterTags{}, nil
	}
	tableName := "dialog_filter_tags"

	query := fmt.Sprintf("select %s from %s where user_id in (%s)", dialogFilterTagsRows, tableName, sqlx.InInt64List(userId))

	var resp []DialogFilterTags
	err := m.db.QueryRowsPartial(ctx, &resp, query)
	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return []DialogFilterTags{}, nil
		}
		return nil, fmt.Errorf("dialog_filter_tags.FindListByUserIdList: %w", err)
	}

	return resp, nil
}
