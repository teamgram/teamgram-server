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
	"github.com/teamgram/teamgram-server/app/messenger/sync/internal/dal/dataobject"

	"github.com/zeromicro/go-zero/core/logx"
)

var _ *sql.Result

type UserQtsUpdatesDAO struct {
	db *sqlx.DB
}

func NewUserQtsUpdatesDAO(db *sqlx.DB) *UserQtsUpdatesDAO {
	return &UserQtsUpdatesDAO{db}
}

// Insert
// insert into user_qts_updates(user_id, qts, update_type, update_data, date2) values (:user_id, :qts, :update_type, :update_data, :date2)
// TODO(@benqi): sqlmap
func (dao *UserQtsUpdatesDAO) Insert(ctx context.Context, do *dataobject.UserQtsUpdatesDO) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into user_qts_updates(user_id, qts, update_type, update_data, date2) values (:user_id, :qts, :update_type, :update_data, :date2)"
		r     sql.Result
	)

	r, err = dao.db.NamedExec(ctx, query, do)
	if err != nil {
		logx.WithContext(ctx).Errorf("namedExec in Insert(%v), error: %v", do, err)
		return
	}

	lastInsertId, err = r.LastInsertId()
	if err != nil {
		logx.WithContext(ctx).Errorf("lastInsertId in Insert(%v)_error: %v", do, err)
		return
	}
	rowsAffected, err = r.RowsAffected()
	if err != nil {
		logx.WithContext(ctx).Errorf("rowsAffected in Insert(%v)_error: %v", do, err)
	}

	return
}

// InsertTx
// insert into user_qts_updates(user_id, qts, update_type, update_data, date2) values (:user_id, :qts, :update_type, :update_data, :date2)
// TODO(@benqi): sqlmap
func (dao *UserQtsUpdatesDAO) InsertTx(tx *sqlx.Tx, do *dataobject.UserQtsUpdatesDO) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into user_qts_updates(user_id, qts, update_type, update_data, date2) values (:user_id, :qts, :update_type, :update_data, :date2)"
		r     sql.Result
	)

	r, err = tx.NamedExec(query, do)
	if err != nil {
		logx.WithContext(tx.Context()).Errorf("namedExec in Insert(%v), error: %v", do, err)
		return
	}

	lastInsertId, err = r.LastInsertId()
	if err != nil {
		logx.WithContext(tx.Context()).Errorf("lastInsertId in Insert(%v)_error: %v", do, err)
		return
	}
	rowsAffected, err = r.RowsAffected()
	if err != nil {
		logx.WithContext(tx.Context()).Errorf("rowsAffected in Insert(%v)_error: %v", do, err)
	}

	return
}

// SelectLastQts
// select qts from user_qts_updates where user_id = :user_id order by qts desc limit 1
// TODO(@benqi): sqlmap
func (dao *UserQtsUpdatesDAO) SelectLastQts(ctx context.Context, user_id int64) (rValue *dataobject.UserQtsUpdatesDO, err error) {
	var (
		query = "select qts from user_qts_updates where user_id = ? order by qts desc limit 1"
		rows  *sqlx.Rows
	)
	rows, err = dao.db.Query(ctx, query, user_id)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectLastQts(_), error: %v", err)
		return
	}

	defer rows.Close()

	do := &dataobject.UserQtsUpdatesDO{}
	if rows.Next() {
		// TODO(@benqi): not use reflect
		err = rows.StructScan(do)
		if err != nil {
			logx.WithContext(ctx).Errorf("structScan in SelectLastQts(_), error: %v", err)
			return
		} else {
			rValue = do
		}
	}

	return
}

// SelectByGtQts
// select user_id, qts, update_type, update_data, date2 from user_qts_updates where user_id = :user_id and qts > :qts order by qts asc
// TODO(@benqi): sqlmap
func (dao *UserQtsUpdatesDAO) SelectByGtQts(ctx context.Context, user_id int64, qts int32) (rList []dataobject.UserQtsUpdatesDO, err error) {
	var (
		query = "select user_id, qts, update_type, update_data, date2 from user_qts_updates where user_id = ? and qts > ? order by qts asc"
		rows  *sqlx.Rows
	)
	rows, err = dao.db.Query(ctx, query, user_id, qts)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectByGtQts(_), error: %v", err)
		return
	}

	defer rows.Close()

	var values []dataobject.UserQtsUpdatesDO
	for rows.Next() {
		v := dataobject.UserQtsUpdatesDO{}

		// TODO(@benqi): not use reflect
		err = rows.StructScan(&v)
		if err != nil {
			logx.WithContext(ctx).Errorf("structScan in SelectByGtQts(_), error: %v", err)
			return
		}
		values = append(values, v)
	}
	rList = values

	return
}

// SelectByGtQtsWithCB
// select user_id, qts, update_type, update_data, date2 from user_qts_updates where user_id = :user_id and qts > :qts order by qts asc
// TODO(@benqi): sqlmap
func (dao *UserQtsUpdatesDAO) SelectByGtQtsWithCB(ctx context.Context, user_id int64, qts int32, cb func(i int, v *dataobject.UserQtsUpdatesDO)) (rList []dataobject.UserQtsUpdatesDO, err error) {
	var (
		query = "select user_id, qts, update_type, update_data, date2 from user_qts_updates where user_id = ? and qts > ? order by qts asc"
		rows  *sqlx.Rows
	)
	rows, err = dao.db.Query(ctx, query, user_id, qts)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectByGtQts(_), error: %v", err)
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

	var values []dataobject.UserQtsUpdatesDO
	for rows.Next() {
		v := dataobject.UserQtsUpdatesDO{}

		// TODO(@benqi): not use reflect
		err = rows.StructScan(&v)
		if err != nil {
			logx.WithContext(ctx).Errorf("structScan in SelectByGtQts(_), error: %v", err)
			return
		}
		values = append(values, v)
	}
	rList = values

	return
}
