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
	dialogFiltersFieldNames          = builder.RawFieldNames(&DialogFilters{})
	dialogFiltersRows                = strings.Join(dialogFiltersFieldNames, ",")
	dialogFiltersRowsExpectAutoSet   = strings.Join(stringx.Remove(dialogFiltersFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	dialogFiltersRowsWithPlaceHolder = strings.Join(stringx.Remove(dialogFiltersFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"
)

type (
	dialogFiltersModel interface {
		Insert2(ctx context.Context, data *DialogFilters) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*DialogFilters, error)
		FindListByIdList(ctx context.Context, id ...int64) ([]DialogFilters, error)
		Update2(ctx context.Context, data *DialogFilters) error
		Delete2(ctx context.Context, id int64) error

		FindOneByUserIdDialogFilterId(ctx context.Context, userId int64, dialogFilterId int32) (*DialogFilters, error)
	}

	defaultDialogFiltersModel struct {
		db *sqlx.DB
	}

	DialogFilters struct {
		Id             int64  `db:"id" json:"id"`
		UserId         int64  `db:"user_id" json:"user_id"`
		DialogFilterId int32  `db:"dialog_filter_id" json:"dialog_filter_id"`
		IsChatlist     bool   `db:"is_chatlist" json:"is_chatlist"`
		JoinedBySlug   bool   `db:"joined_by_slug" json:"joined_by_slug"`
		Slug           string `db:"slug" json:"slug"`
		HasMyInvites   int32  `db:"has_my_invites" json:"has_my_invites"`
		DialogFilter   string `db:"dialog_filter" json:"dialog_filter"`
		OrderValue     int64  `db:"order_value" json:"order_value"`
		FromSuggested  int32  `db:"from_suggested" json:"from_suggested"`
		Deleted        bool   `db:"deleted" json:"deleted"`
	}
)

func newDialogFiltersModel(db *sqlx.DB) *defaultDialogFiltersModel {
	return &defaultDialogFiltersModel{
		db: db,
	}
}

func (m *defaultDialogFiltersModel) Insert2(ctx context.Context, data *DialogFilters) (sql.Result, error) {
	query := fmt.Sprintf("insert into `dialog_filters` (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", dialogFiltersRowsExpectAutoSet)

	return m.db.Exec(ctx, query, data.UserId, data.DialogFilterId, data.IsChatlist, data.JoinedBySlug, data.Slug, data.HasMyInvites, data.DialogFilter, data.OrderValue, data.FromSuggested, data.Deleted)

}

func (m *defaultDialogFiltersModel) Delete2(ctx context.Context, id int64) error {
	query := "delete from `dialog_filters` where `id` = ?"

	_, err := m.db.Exec(ctx, query, id)
	return err
}

func (m *defaultDialogFiltersModel) FindOne(ctx context.Context, id int64) (*DialogFilters, error) {
	query := fmt.Sprintf("select %s from dialog_filters where id = ? limit 1", dialogFiltersRows)
	var resp DialogFilters

	err := m.db.QueryRowPartial(ctx, &resp, query, id)

	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (m *defaultDialogFiltersModel) FindListByIdList(ctx context.Context, id ...int64) ([]DialogFilters, error) {
	if len(id) == 0 {
		return []DialogFilters{}, nil
	}

	query := fmt.Sprintf("select %s from dialog_filters where id in (%s)", dialogFiltersRows, sqlx.InInt64List(id))

	var resp []DialogFilters
	err := m.db.QueryRowsPartial(ctx, &resp, query)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (m *defaultDialogFiltersModel) Update2(ctx context.Context, data *DialogFilters) error {
	query := fmt.Sprintf("update `dialog_filters` set %s where `id` = ?", dialogFiltersRowsWithPlaceHolder)

	_, err := m.db.Exec(ctx, query, data.UserId, data.DialogFilterId, data.IsChatlist, data.JoinedBySlug, data.Slug, data.HasMyInvites, data.DialogFilter, data.OrderValue, data.FromSuggested, data.Deleted, data.Id)
	return err
}

func (m *defaultDialogFiltersModel) FindOneByUserIdDialogFilterId(ctx context.Context, userId int64, dialogFilterId int32) (*DialogFilters, error) {
	query := fmt.Sprintf("select %s from dialog_filters where user_id = ? AND dialog_filter_id = ? limit 1", dialogFiltersRows)
	var resp DialogFilters

	err := m.db.QueryRowPartial(ctx, &resp, query, userId, dialogFilterId)

	if err != nil {
		return nil, err
	}
	return &resp, nil
}
