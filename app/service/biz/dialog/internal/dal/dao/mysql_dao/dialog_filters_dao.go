/*
 * WARNING! All changes made in this file will be lost!
 *   Created from by 'dalgen'
 *
 * Copyright (c) 2024-present,  Teamgram Authors.
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package mysql_dao

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/teamgram/marmota/pkg/stores/sqlx"
	"github.com/teamgram/teamgram-server/app/service/biz/dialog/internal/dal/dataobject"

	"github.com/zeromicro/go-zero/core/logx"
)

var _ *sql.Result
var _ = fmt.Sprintf
var _ = strings.Join

type DialogFiltersDAO struct {
	db *sqlx.DB
}

func NewDialogFiltersDAO(db *sqlx.DB) *DialogFiltersDAO {
	return &DialogFiltersDAO{
		db: db,
	}
}

// InsertOrUpdate
// insert into dialog_filters(user_id, dialog_filter_id, is_chatlist, joined_by_slug, slug, dialog_filter, order_value) values (:user_id, :dialog_filter_id, :is_chatlist, :joined_by_slug, :slug, :dialog_filter, :order_value) on duplicate key update is_chatlist = values(is_chatlist), dialog_filter = values(dialog_filter), joined_by_slug = values(joined_by_slug), slug = values(slug), order_value = values(order_value), deleted = 0
// TODO(@benqi): sqlmap
func (dao *DialogFiltersDAO) InsertOrUpdate(ctx context.Context, do *dataobject.DialogFiltersDO) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into dialog_filters(user_id, dialog_filter_id, is_chatlist, joined_by_slug, slug, dialog_filter, order_value) values (:user_id, :dialog_filter_id, :is_chatlist, :joined_by_slug, :slug, :dialog_filter, :order_value) on duplicate key update is_chatlist = values(is_chatlist), dialog_filter = values(dialog_filter), joined_by_slug = values(joined_by_slug), slug = values(slug), order_value = values(order_value), deleted = 0"
		r     sql.Result
	)

	r, err = dao.db.NamedExec(ctx, query, do)
	if err != nil {
		logx.WithContext(ctx).Errorf("namedExec in InsertOrUpdate(%v), error: %v", do, err)
		return
	}

	lastInsertId, err = r.LastInsertId()
	if err != nil {
		logx.WithContext(ctx).Errorf("lastInsertId in InsertOrUpdate(%v)_error: %v", do, err)
		return
	}
	rowsAffected, err = r.RowsAffected()
	if err != nil {
		logx.WithContext(ctx).Errorf("rowsAffected in InsertOrUpdate(%v)_error: %v", do, err)
	}

	return
}

// InsertOrUpdateTx
// insert into dialog_filters(user_id, dialog_filter_id, is_chatlist, joined_by_slug, slug, dialog_filter, order_value) values (:user_id, :dialog_filter_id, :is_chatlist, :joined_by_slug, :slug, :dialog_filter, :order_value) on duplicate key update is_chatlist = values(is_chatlist), dialog_filter = values(dialog_filter), joined_by_slug = values(joined_by_slug), slug = values(slug), order_value = values(order_value), deleted = 0
// TODO(@benqi): sqlmap
func (dao *DialogFiltersDAO) InsertOrUpdateTx(tx *sqlx.Tx, do *dataobject.DialogFiltersDO) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into dialog_filters(user_id, dialog_filter_id, is_chatlist, joined_by_slug, slug, dialog_filter, order_value) values (:user_id, :dialog_filter_id, :is_chatlist, :joined_by_slug, :slug, :dialog_filter, :order_value) on duplicate key update is_chatlist = values(is_chatlist), dialog_filter = values(dialog_filter), joined_by_slug = values(joined_by_slug), slug = values(slug), order_value = values(order_value), deleted = 0"
		r     sql.Result
	)

	r, err = tx.NamedExec(query, do)
	if err != nil {
		logx.WithContext(tx.Context()).Errorf("namedExec in InsertOrUpdate(%v), error: %v", do, err)
		return
	}

	lastInsertId, err = r.LastInsertId()
	if err != nil {
		logx.WithContext(tx.Context()).Errorf("lastInsertId in InsertOrUpdate(%v)_error: %v", do, err)
		return
	}
	rowsAffected, err = r.RowsAffected()
	if err != nil {
		logx.WithContext(tx.Context()).Errorf("rowsAffected in InsertOrUpdate(%v)_error: %v", do, err)
	}

	return
}

// SelectBySlug
// select id, user_id, dialog_filter_id, is_chatlist, joined_by_slug, slug, dialog_filter, order_value, from_suggested from dialog_filters where user_id = :user_id and slug = :slug and deleted = 0 order by order_value desc
// TODO(@benqi): sqlmap
func (dao *DialogFiltersDAO) SelectBySlug(ctx context.Context, user_id int64, slug string) (rValue *dataobject.DialogFiltersDO, err error) {
	var (
		query = "select id, user_id, dialog_filter_id, is_chatlist, joined_by_slug, slug, dialog_filter, order_value, from_suggested from dialog_filters where user_id = ? and slug = ? and deleted = 0 order by order_value desc"
		do    = &dataobject.DialogFiltersDO{}
	)
	err = dao.db.QueryRowPartial(ctx, do, query, user_id, slug)

	if err != nil {
		if err != sqlx.ErrNotFound {
			logx.WithContext(ctx).Errorf("queryx in SelectBySlug(_), error: %v", err)
			return
		} else {
			err = nil
		}
	} else {
		rValue = do
	}

	return
}

// Select
// select id, user_id, dialog_filter_id, is_chatlist, joined_by_slug, slug, dialog_filter, order_value, from_suggested from dialog_filters where user_id = :user_id and dialog_filter_id = :dialog_filter_id and deleted = 0 order by order_value desc
// TODO(@benqi): sqlmap
func (dao *DialogFiltersDAO) Select(ctx context.Context, user_id int64, dialog_filter_id int32) (rValue *dataobject.DialogFiltersDO, err error) {
	var (
		query = "select id, user_id, dialog_filter_id, is_chatlist, joined_by_slug, slug, dialog_filter, order_value, from_suggested from dialog_filters where user_id = ? and dialog_filter_id = ? and deleted = 0 order by order_value desc"
		do    = &dataobject.DialogFiltersDO{}
	)
	err = dao.db.QueryRowPartial(ctx, do, query, user_id, dialog_filter_id)

	if err != nil {
		if err != sqlx.ErrNotFound {
			logx.WithContext(ctx).Errorf("queryx in Select(_), error: %v", err)
			return
		} else {
			err = nil
		}
	} else {
		rValue = do
	}

	return
}

// SelectList
// select id, user_id, dialog_filter_id, is_chatlist, joined_by_slug, slug, dialog_filter, order_value, from_suggested from dialog_filters where user_id = :user_id and deleted = 0 order by order_value desc
// TODO(@benqi): sqlmap
func (dao *DialogFiltersDAO) SelectList(ctx context.Context, user_id int64) (rList []dataobject.DialogFiltersDO, err error) {
	var (
		query  = "select id, user_id, dialog_filter_id, is_chatlist, joined_by_slug, slug, dialog_filter, order_value, from_suggested from dialog_filters where user_id = ? and deleted = 0 order by order_value desc"
		values []dataobject.DialogFiltersDO
	)
	err = dao.db.QueryRowsPartial(ctx, &values, query, user_id)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectList(_), error: %v", err)
		return
	}

	rList = values

	return
}

// SelectListWithCB
// select id, user_id, dialog_filter_id, is_chatlist, joined_by_slug, slug, dialog_filter, order_value, from_suggested from dialog_filters where user_id = :user_id and deleted = 0 order by order_value desc
// TODO(@benqi): sqlmap
func (dao *DialogFiltersDAO) SelectListWithCB(ctx context.Context, user_id int64, cb func(i int, v *dataobject.DialogFiltersDO)) (rList []dataobject.DialogFiltersDO, err error) {
	var (
		query  = "select id, user_id, dialog_filter_id, is_chatlist, joined_by_slug, slug, dialog_filter, order_value, from_suggested from dialog_filters where user_id = ? and deleted = 0 order by order_value desc"
		values []dataobject.DialogFiltersDO
	)
	err = dao.db.QueryRowsPartial(ctx, &values, query, user_id)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectList(_), error: %v", err)
		return
	}

	rList = values

	if cb != nil {
		for i := 0; i < len(rList); i++ {
			cb(i, &rList[i])
		}
	}

	return
}

// UpdateOrder
// update dialog_filters set order_value = :order_value where user_id = :user_id and dialog_filter_id = :dialog_filter_id
// TODO(@benqi): sqlmap
func (dao *DialogFiltersDAO) UpdateOrder(ctx context.Context, order_value int64, user_id int64, dialog_filter_id int32) (rowsAffected int64, err error) {
	var (
		query   = "update dialog_filters set order_value = ? where user_id = ? and dialog_filter_id = ?"
		rResult sql.Result
	)
	rResult, err = dao.db.Exec(ctx, query, order_value, user_id, dialog_filter_id)

	if err != nil {
		logx.WithContext(ctx).Errorf("exec in UpdateOrder(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		logx.WithContext(ctx).Errorf("rowsAffected in UpdateOrder(_), error: %v", err)
	}

	return
}

// UpdateOrderTx
// update dialog_filters set order_value = :order_value where user_id = :user_id and dialog_filter_id = :dialog_filter_id
// TODO(@benqi): sqlmap
func (dao *DialogFiltersDAO) UpdateOrderTx(tx *sqlx.Tx, order_value int64, user_id int64, dialog_filter_id int32) (rowsAffected int64, err error) {
	var (
		query   = "update dialog_filters set order_value = ? where user_id = ? and dialog_filter_id = ?"
		rResult sql.Result
	)
	rResult, err = tx.Exec(query, order_value, user_id, dialog_filter_id)

	if err != nil {
		logx.WithContext(tx.Context()).Errorf("exec in UpdateOrder(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		logx.WithContext(tx.Context()).Errorf("rowsAffected in UpdateOrder(_), error: %v", err)
	}

	return
}

// Clear
// update dialog_filters set deleted = 1, dialog_filter = 'null', order_value = 0 where user_id = :user_id and dialog_filter_id = :dialog_filter_id
// TODO(@benqi): sqlmap
func (dao *DialogFiltersDAO) Clear(ctx context.Context, user_id int64, dialog_filter_id int32) (rowsAffected int64, err error) {
	var (
		query   = "update dialog_filters set deleted = 1, dialog_filter = 'null', order_value = 0 where user_id = ? and dialog_filter_id = ?"
		rResult sql.Result
	)
	rResult, err = dao.db.Exec(ctx, query, user_id, dialog_filter_id)

	if err != nil {
		logx.WithContext(ctx).Errorf("exec in Clear(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		logx.WithContext(ctx).Errorf("rowsAffected in Clear(_), error: %v", err)
	}

	return
}

// ClearTx
// update dialog_filters set deleted = 1, dialog_filter = 'null', order_value = 0 where user_id = :user_id and dialog_filter_id = :dialog_filter_id
// TODO(@benqi): sqlmap
func (dao *DialogFiltersDAO) ClearTx(tx *sqlx.Tx, user_id int64, dialog_filter_id int32) (rowsAffected int64, err error) {
	var (
		query   = "update dialog_filters set deleted = 1, dialog_filter = 'null', order_value = 0 where user_id = ? and dialog_filter_id = ?"
		rResult sql.Result
	)
	rResult, err = tx.Exec(query, user_id, dialog_filter_id)

	if err != nil {
		logx.WithContext(tx.Context()).Errorf("exec in Clear(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		logx.WithContext(tx.Context()).Errorf("rowsAffected in Clear(_), error: %v", err)
	}

	return
}
