/*
 * WARNING! All changes made in this file will be lost!
 *   Created from by 'dalgen'
 *
 * Copyright (c) 2021-present,  Teamgram Studio (https://teamgram.io).
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package mysql_dao

import (
	"context"
	"database/sql"

	"github.com/teamgram/marmota/pkg/stores/sqlx"
	"github.com/teamgram/teamgram-server/app/service/biz/dialog/internal/dal/dataobject"

	"github.com/zeromicro/go-zero/core/logx"
)

var _ *sql.Result

type DialogFiltersDAO struct {
	db *sqlx.DB
}

func NewDialogFiltersDAO(db *sqlx.DB) *DialogFiltersDAO {
	return &DialogFiltersDAO{db}
}

// InsertOrUpdate
// insert into dialog_filters(user_id, dialog_filter_id, dialog_filter, order_value) values (:user_id, :dialog_filter_id, :dialog_filter, :order_value) on duplicate key update dialog_filter = values(dialog_filter), order_value = values(order_value), deleted = 0
func (dao *DialogFiltersDAO) InsertOrUpdate(ctx context.Context, do *dataobject.DialogFiltersDO) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into dialog_filters(user_id, dialog_filter_id, dialog_filter, order_value) values (:user_id, :dialog_filter_id, :dialog_filter, :order_value) on duplicate key update dialog_filter = values(dialog_filter), order_value = values(order_value), deleted = 0"
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
// insert into dialog_filters(user_id, dialog_filter_id, dialog_filter, order_value) values (:user_id, :dialog_filter_id, :dialog_filter, :order_value) on duplicate key update dialog_filter = values(dialog_filter), order_value = values(order_value), deleted = 0
func (dao *DialogFiltersDAO) InsertOrUpdateTx(tx *sqlx.Tx, do *dataobject.DialogFiltersDO) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into dialog_filters(user_id, dialog_filter_id, dialog_filter, order_value) values (:user_id, :dialog_filter_id, :dialog_filter, :order_value) on duplicate key update dialog_filter = values(dialog_filter), order_value = values(order_value), deleted = 0"
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

// SelectList
// select user_id, dialog_filter_id, dialog_filter from dialog_filters where user_id = :user_id and deleted = 0 order by order_value desc
func (dao *DialogFiltersDAO) SelectList(ctx context.Context, user_id int64) (rList []dataobject.DialogFiltersDO, err error) {
	var (
		query = "select user_id, dialog_filter_id, dialog_filter from dialog_filters where user_id = ? and deleted = 0 order by order_value desc"
		rows  *sqlx.Rows
	)
	rows, err = dao.db.Query(ctx, query, user_id)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectList(_), error: %v", err)
		return
	}

	defer rows.Close()

	var values []dataobject.DialogFiltersDO
	for rows.Next() {
		v := dataobject.DialogFiltersDO{}

		// TODO(@benqi): not use reflect
		err = rows.StructScan(&v)
		if err != nil {
			logx.WithContext(ctx).Errorf("structScan in SelectList(_), error: %v", err)
			return
		}
		values = append(values, v)
	}
	rList = values

	return
}

// SelectListWithCB
// select user_id, dialog_filter_id, dialog_filter from dialog_filters where user_id = :user_id and deleted = 0 order by order_value desc
func (dao *DialogFiltersDAO) SelectListWithCB(ctx context.Context, user_id int64, cb func(i int, v *dataobject.DialogFiltersDO)) (rList []dataobject.DialogFiltersDO, err error) {
	var (
		query = "select user_id, dialog_filter_id, dialog_filter from dialog_filters where user_id = ? and deleted = 0 order by order_value desc"
		rows  *sqlx.Rows
	)
	rows, err = dao.db.Query(ctx, query, user_id)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectList(_), error: %v", err)
		return
	}

	defer func() {
		rows.Close()
		if err == nil && cb != nil {
			for i := 0; i < len(rList); i++ {
				cb(i, &rList[i])
			}
		}
	}()

	var values []dataobject.DialogFiltersDO
	for rows.Next() {
		v := dataobject.DialogFiltersDO{}

		// TODO(@benqi): not use reflect
		err = rows.StructScan(&v)
		if err != nil {
			logx.WithContext(ctx).Errorf("structScan in SelectList(_), error: %v", err)
			return
		}
		values = append(values, v)
	}
	rList = values

	return
}

// UpdateOrder
// update dialog_filters set order_value = :order_value where user_id = :user_id and dialog_filter_id = :dialog_filter_id
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
