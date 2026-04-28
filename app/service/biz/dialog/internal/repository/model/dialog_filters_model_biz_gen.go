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
)

var _ *sql.Result
var _ = fmt.Sprintf
var _ = strings.Join
var _ = errors.Is
var _ *sqlx.DB

type (
	bizDialogFiltersModel interface {
		InsertOrUpdate(ctx context.Context, data *DialogFilters) (lastInsertId, rowsAffected int64, err error)
		InsertOrUpdateTx(tx *sqlx.Tx, data *DialogFilters) (lastInsertId, rowsAffected int64, err error)

		SelectBySlug(ctx context.Context, userId int64, slug string) (*DialogFilters, error)

		Select(ctx context.Context, userId int64, dialogFilterId int32) (*DialogFilters, error)

		SelectList(ctx context.Context, userId int64) ([]DialogFilters, error)
		SelectListWithCB(ctx context.Context, userId int64, cb func(sz, i int, v *DialogFilters)) ([]DialogFilters, error)

		UpdateOrder(ctx context.Context, orderValue int64, userId int64, dialogFilterId int32) (rowsAffected int64, err error)
		UpdateOrderTx(tx *sqlx.Tx, orderValue int64, userId int64, dialogFilterId int32) (rowsAffected int64, err error)

		Clear(ctx context.Context, userId int64, dialogFilterId int32) (rowsAffected int64, err error)
		ClearTx(tx *sqlx.Tx, userId int64, dialogFilterId int32) (rowsAffected int64, err error)
	}
)

// InsertOrUpdate
// insert into dialog_filters(user_id, dialog_filter_id, is_chatlist, joined_by_slug, slug, dialog_filter, order_value) values (:user_id, :dialog_filter_id, :is_chatlist, :joined_by_slug, :slug, :dialog_filter, :order_value) on duplicate key update is_chatlist = values(is_chatlist), dialog_filter = values(dialog_filter), joined_by_slug = values(joined_by_slug), slug = values(slug), order_value = values(order_value), deleted = 0
func (m *defaultDialogFiltersModel) InsertOrUpdate(ctx context.Context, data *DialogFilters) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into dialog_filters(user_id, dialog_filter_id, is_chatlist, joined_by_slug, slug, dialog_filter, order_value) values (:user_id, :dialog_filter_id, :is_chatlist, :joined_by_slug, :slug, :dialog_filter, :order_value) on duplicate key update is_chatlist = values(is_chatlist), dialog_filter = values(dialog_filter), joined_by_slug = values(joined_by_slug), slug = values(slug), order_value = values(order_value), deleted = 0"
		r     sql.Result
	)

	r, err = m.db.NamedExec(ctx, query, data)
	if err != nil {
		err = fmt.Errorf("dialog_filters.InsertOrUpdate named exec: %w", err)
		return
	}

	lastInsertId, err = r.LastInsertId()
	if err != nil {
		err = fmt.Errorf("dialog_filters.InsertOrUpdate last insert id: %w", err)
		return
	}
	rowsAffected, err = r.RowsAffected()
	if err != nil {
		err = fmt.Errorf("dialog_filters.InsertOrUpdate rows affected: %w", err)
	}

	return

}

// InsertOrUpdateTx
// insert into dialog_filters(user_id, dialog_filter_id, is_chatlist, joined_by_slug, slug, dialog_filter, order_value) values (:user_id, :dialog_filter_id, :is_chatlist, :joined_by_slug, :slug, :dialog_filter, :order_value) on duplicate key update is_chatlist = values(is_chatlist), dialog_filter = values(dialog_filter), joined_by_slug = values(joined_by_slug), slug = values(slug), order_value = values(order_value), deleted = 0
func (m *defaultDialogFiltersModel) InsertOrUpdateTx(tx *sqlx.Tx, data *DialogFilters) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into dialog_filters(user_id, dialog_filter_id, is_chatlist, joined_by_slug, slug, dialog_filter, order_value) values (:user_id, :dialog_filter_id, :is_chatlist, :joined_by_slug, :slug, :dialog_filter, :order_value) on duplicate key update is_chatlist = values(is_chatlist), dialog_filter = values(dialog_filter), joined_by_slug = values(joined_by_slug), slug = values(slug), order_value = values(order_value), deleted = 0"
		r     sql.Result
	)

	r, err = tx.NamedExec(query, data)
	if err != nil {
		err = fmt.Errorf("dialog_filters.InsertOrUpdateTx named exec: %w", err)
		return
	}

	lastInsertId, err = r.LastInsertId()
	if err != nil {
		err = fmt.Errorf("dialog_filters.InsertOrUpdateTx last insert id: %w", err)
		return
	}
	rowsAffected, err = r.RowsAffected()
	if err != nil {
		err = fmt.Errorf("dialog_filters.InsertOrUpdateTx rows affected: %w", err)
	}

	return
}

// SelectBySlug
// select id, user_id, dialog_filter_id, is_chatlist, joined_by_slug, slug, dialog_filter, order_value, from_suggested from dialog_filters where user_id = :user_id and slug = :slug and deleted = 0 order by order_value desc
func (m *defaultDialogFiltersModel) SelectBySlug(ctx context.Context, userId int64, slug string) (rValue *DialogFilters, err error) {

	var (
		query = "select id, user_id, dialog_filter_id, is_chatlist, joined_by_slug, slug, dialog_filter, order_value, from_suggested from dialog_filters where user_id = ? and slug = ? and deleted = 0 order by order_value desc"
		do    = &DialogFilters{}
	)
	err = m.db.QueryRowPartial(ctx, do, query, userId, slug)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return nil, &NotFoundError{
				Resource: "dialog_filters",
				Key:      fmt.Sprintf("user_id=%v,slug=%v", userId, slug),
				Cause:    err,
			}
		}
		err = fmt.Errorf("dialog_filters.SelectBySlug: %w", err)
		return
	} else {
		rValue = do
	}

	return
}

// Select
// select id, user_id, dialog_filter_id, is_chatlist, joined_by_slug, slug, dialog_filter, order_value, from_suggested from dialog_filters where user_id = :user_id and dialog_filter_id = :dialog_filter_id and deleted = 0 order by order_value desc
func (m *defaultDialogFiltersModel) Select(ctx context.Context, userId int64, dialogFilterId int32) (rValue *DialogFilters, err error) {

	var (
		query = "select id, user_id, dialog_filter_id, is_chatlist, joined_by_slug, slug, dialog_filter, order_value, from_suggested from dialog_filters where user_id = ? and dialog_filter_id = ? and deleted = 0 order by order_value desc"
		do    = &DialogFilters{}
	)
	err = m.db.QueryRowPartial(ctx, do, query, userId, dialogFilterId)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return nil, &NotFoundError{
				Resource: "dialog_filters",
				Key:      fmt.Sprintf("user_id=%v,dialog_filter_id=%v", userId, dialogFilterId),
				Cause:    err,
			}
		}
		err = fmt.Errorf("dialog_filters.Select: %w", err)
		return
	} else {
		rValue = do
	}

	return
}

// SelectList
// select id, user_id, dialog_filter_id, is_chatlist, joined_by_slug, slug, dialog_filter, order_value, from_suggested from dialog_filters where user_id = :user_id and deleted = 0 order by order_value desc
func (m *defaultDialogFiltersModel) SelectList(ctx context.Context, userId int64) (rList []DialogFilters, err error) {
	var (
		query  = "select id, user_id, dialog_filter_id, is_chatlist, joined_by_slug, slug, dialog_filter, order_value, from_suggested from dialog_filters where user_id = ? and deleted = 0 order by order_value desc"
		values []DialogFilters
	)
	err = m.db.QueryRowsPartial(ctx, &values, query, userId)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			rList = []DialogFilters{}
			err = nil
			return
		}
		err = fmt.Errorf("dialog_filters.SelectList: %w", err)
		return
	}

	rList = values

	return
}

// SelectListWithCB
// select id, user_id, dialog_filter_id, is_chatlist, joined_by_slug, slug, dialog_filter, order_value, from_suggested from dialog_filters where user_id = :user_id and deleted = 0 order by order_value desc
func (m *defaultDialogFiltersModel) SelectListWithCB(ctx context.Context, userId int64, cb func(sz, i int, v *DialogFilters)) (rList []DialogFilters, err error) {
	var (
		query  = "select id, user_id, dialog_filter_id, is_chatlist, joined_by_slug, slug, dialog_filter, order_value, from_suggested from dialog_filters where user_id = ? and deleted = 0 order by order_value desc"
		values []DialogFilters
	)
	err = m.db.QueryRowsPartial(ctx, &values, query, userId)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			rList = []DialogFilters{}
			err = nil
			return
		}
		err = fmt.Errorf("dialog_filters.SelectListWithCB: %w", err)
		return
	}

	rList = values

	if cb != nil {
		sz := len(rList)
		for i := 0; i < sz; i++ {
			cb(sz, i, &rList[i])
		}
	}

	return
}

// UpdateOrder
// update dialog_filters set order_value = :order_value where user_id = :user_id and dialog_filter_id = :dialog_filter_id
func (m *defaultDialogFiltersModel) UpdateOrder(ctx context.Context, orderValue int64, userId int64, dialogFilterId int32) (rowsAffected int64, err error) {

	var (
		query   = "update dialog_filters set order_value = ? where user_id = ? and dialog_filter_id = ?"
		rResult sql.Result
	)

	rResult, err = m.db.Exec(ctx, query, orderValue, userId, dialogFilterId)

	if err != nil {
		err = fmt.Errorf("dialog_filters.UpdateOrder exec: %w", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		err = fmt.Errorf("dialog_filters.UpdateOrder rows affected: %w", err)
		return
	}

	return
}

// UpdateOrderTx
// update dialog_filters set order_value = :order_value where user_id = :user_id and dialog_filter_id = :dialog_filter_id
func (m *defaultDialogFiltersModel) UpdateOrderTx(tx *sqlx.Tx, orderValue int64, userId int64, dialogFilterId int32) (rowsAffected int64, err error) {
	var (
		query   = "update dialog_filters set order_value = ? where user_id = ? and dialog_filter_id = ?"
		rResult sql.Result
	)
	rResult, err = tx.Exec(query, orderValue, userId, dialogFilterId)

	if err != nil {
		err = fmt.Errorf("dialog_filters.UpdateOrderTx exec: %w", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		err = fmt.Errorf("dialog_filters.UpdateOrderTx rows affected: %w", err)
		return
	}

	return
}

// Clear
// update dialog_filters set deleted = 1, dialog_filter = 'null', order_value = 0 where user_id = :user_id and dialog_filter_id = :dialog_filter_id
func (m *defaultDialogFiltersModel) Clear(ctx context.Context, userId int64, dialogFilterId int32) (rowsAffected int64, err error) {

	var (
		query   = "update dialog_filters set deleted = 1, dialog_filter = 'null', order_value = 0 where user_id = ? and dialog_filter_id = ?"
		rResult sql.Result
	)

	rResult, err = m.db.Exec(ctx, query, userId, dialogFilterId)

	if err != nil {
		err = fmt.Errorf("dialog_filters.Clear exec: %w", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		err = fmt.Errorf("dialog_filters.Clear rows affected: %w", err)
		return
	}

	return
}

// ClearTx
// update dialog_filters set deleted = 1, dialog_filter = 'null', order_value = 0 where user_id = :user_id and dialog_filter_id = :dialog_filter_id
func (m *defaultDialogFiltersModel) ClearTx(tx *sqlx.Tx, userId int64, dialogFilterId int32) (rowsAffected int64, err error) {
	var (
		query   = "update dialog_filters set deleted = 1, dialog_filter = 'null', order_value = 0 where user_id = ? and dialog_filter_id = ?"
		rResult sql.Result
	)
	rResult, err = tx.Exec(query, userId, dialogFilterId)

	if err != nil {
		err = fmt.Errorf("dialog_filters.ClearTx exec: %w", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		err = fmt.Errorf("dialog_filters.ClearTx rows affected: %w", err)
		return
	}

	return
}
