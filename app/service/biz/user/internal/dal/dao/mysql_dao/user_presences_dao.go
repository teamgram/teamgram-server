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
	"github.com/teamgram/teamgram-server/app/service/biz/user/internal/dal/dataobject"

	"github.com/zeromicro/go-zero/core/logx"
)

var _ *sql.Result

type UserPresencesDAO struct {
	db *sqlx.DB
}

func NewUserPresencesDAO(db *sqlx.DB) *UserPresencesDAO {
	return &UserPresencesDAO{db}
}

// Insert
// insert into user_presences(user_id, last_seen_at, expires) values (:user_id, :last_seen_at, :expires)
// TODO(@benqi): sqlmap
func (dao *UserPresencesDAO) Insert(ctx context.Context, do *dataobject.UserPresencesDO) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into user_presences(user_id, last_seen_at, expires) values (:user_id, :last_seen_at, :expires)"
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
// insert into user_presences(user_id, last_seen_at, expires) values (:user_id, :last_seen_at, :expires)
// TODO(@benqi): sqlmap
func (dao *UserPresencesDAO) InsertTx(tx *sqlx.Tx, do *dataobject.UserPresencesDO) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into user_presences(user_id, last_seen_at, expires) values (:user_id, :last_seen_at, :expires)"
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

// Select
// select id, user_id, last_seen_at, expires from user_presences where user_id = :user_id
// TODO(@benqi): sqlmap
func (dao *UserPresencesDAO) Select(ctx context.Context, user_id int64) (rValue *dataobject.UserPresencesDO, err error) {
	var (
		query = "select id, user_id, last_seen_at, expires from user_presences where user_id = ?"
		rows  *sqlx.Rows
	)
	rows, err = dao.db.Query(ctx, query, user_id)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in Select(_), error: %v", err)
		return
	}

	defer rows.Close()

	do := &dataobject.UserPresencesDO{}
	if rows.Next() {
		// TODO(@benqi): not use reflect
		err = rows.StructScan(do)
		if err != nil {
			logx.WithContext(ctx).Errorf("structScan in Select(_), error: %v", err)
			return
		} else {
			rValue = do
		}
	}

	return
}

// SelectList
// select id, user_id, last_seen_at, expires from user_presences where user_id in (:idList)
// TODO(@benqi): sqlmap
func (dao *UserPresencesDAO) SelectList(ctx context.Context, idList []int64) (rList []dataobject.UserPresencesDO, err error) {
	var (
		query = "select id, user_id, last_seen_at, expires from user_presences where user_id in (?)"
		a     []interface{}
		rows  *sqlx.Rows
	)
	if len(idList) == 0 {
		rList = []dataobject.UserPresencesDO{}
		return
	}

	query, a, err = sqlx.In(query, idList)
	if err != nil {
		// r sql.Result
		logx.WithContext(ctx).Errorf("sqlx.In in SelectList(_), error: %v", err)
		return
	}
	rows, err = dao.db.Query(ctx, query, a...)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectList(_), error: %v", err)
		return
	}

	defer rows.Close()

	var values []dataobject.UserPresencesDO
	for rows.Next() {
		v := dataobject.UserPresencesDO{}

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
// select id, user_id, last_seen_at, expires from user_presences where user_id in (:idList)
// TODO(@benqi): sqlmap
func (dao *UserPresencesDAO) SelectListWithCB(ctx context.Context, idList []int64, cb func(i int, v *dataobject.UserPresencesDO)) (rList []dataobject.UserPresencesDO, err error) {
	var (
		query = "select id, user_id, last_seen_at, expires from user_presences where user_id in (?)"
		a     []interface{}
		rows  *sqlx.Rows
	)
	if len(idList) == 0 {
		rList = []dataobject.UserPresencesDO{}
		return
	}

	query, a, err = sqlx.In(query, idList)
	if err != nil {
		// r sql.Result
		logx.WithContext(ctx).Errorf("sqlx.In in SelectList(_), error: %v", err)
		return
	}
	rows, err = dao.db.Query(ctx, query, a...)

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

	var values []dataobject.UserPresencesDO
	for rows.Next() {
		v := dataobject.UserPresencesDO{}

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

// UpdateLastSeenAt
// update user_presences set last_seen_at = :last_seen_at, expires = :expires where user_id = :user_id
// TODO(@benqi): sqlmap
func (dao *UserPresencesDAO) UpdateLastSeenAt(ctx context.Context, last_seen_at int64, expires int32, user_id int64) (rowsAffected int64, err error) {
	var (
		query   = "update user_presences set last_seen_at = ?, expires = ? where user_id = ?"
		rResult sql.Result
	)
	rResult, err = dao.db.Exec(ctx, query, last_seen_at, expires, user_id)

	if err != nil {
		logx.WithContext(ctx).Errorf("exec in UpdateLastSeenAt(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		logx.WithContext(ctx).Errorf("rowsAffected in UpdateLastSeenAt(_), error: %v", err)
	}

	return
}

// update user_presences set last_seen_at = :last_seen_at, expires = :expires where user_id = :user_id
// UpdateLastSeenAtTx
// TODO(@benqi): sqlmap
func (dao *UserPresencesDAO) UpdateLastSeenAtTx(tx *sqlx.Tx, last_seen_at int64, expires int32, user_id int64) (rowsAffected int64, err error) {
	var (
		query   = "update user_presences set last_seen_at = ?, expires = ? where user_id = ?"
		rResult sql.Result
	)
	rResult, err = tx.Exec(query, last_seen_at, expires, user_id)

	if err != nil {
		logx.WithContext(tx.Context()).Errorf("exec in UpdateLastSeenAt(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		logx.WithContext(tx.Context()).Errorf("rowsAffected in UpdateLastSeenAt(_), error: %v", err)
	}

	return
}
