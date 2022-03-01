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
	"github.com/teamgram/teamgram-server/app/service/biz/banned/internal/dal/dataobject"

	"github.com/zeromicro/go-zero/core/logx"
)

var _ *sql.Result

type BannedDAO struct {
	db *sqlx.DB
}

func NewBannedDAO(db *sqlx.DB) *BannedDAO {
	return &BannedDAO{db}
}

// InsertOrUpdate
// insert into banned(phone, banned_time, expires, banned_reason, log, state) values (:phone, :banned_time, :expires, :banned_reason, :log, :state) on duplicate key update banned_time = values(banned_time), expires = values(expires), banned_reason = values(banned_reason), log = values(log), state = values(state)
// TODO(@benqi): sqlmap
func (dao *BannedDAO) InsertOrUpdate(ctx context.Context, do *dataobject.BannedDO) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into banned(phone, banned_time, expires, banned_reason, log, state) values (:phone, :banned_time, :expires, :banned_reason, :log, :state) on duplicate key update banned_time = values(banned_time), expires = values(expires), banned_reason = values(banned_reason), log = values(log), state = values(state)"
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
// insert into banned(phone, banned_time, expires, banned_reason, log, state) values (:phone, :banned_time, :expires, :banned_reason, :log, :state) on duplicate key update banned_time = values(banned_time), expires = values(expires), banned_reason = values(banned_reason), log = values(log), state = values(state)
// TODO(@benqi): sqlmap
func (dao *BannedDAO) InsertOrUpdateTx(tx *sqlx.Tx, do *dataobject.BannedDO) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into banned(phone, banned_time, expires, banned_reason, log, state) values (:phone, :banned_time, :expires, :banned_reason, :log, :state) on duplicate key update banned_time = values(banned_time), expires = values(expires), banned_reason = values(banned_reason), log = values(log), state = values(state)"
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

// CheckBannedByPhone
// select id from banned where phone = :phone and state > 0
// TODO(@benqi): sqlmap
func (dao *BannedDAO) CheckBannedByPhone(ctx context.Context, phone string) (rValue *dataobject.BannedDO, err error) {
	var (
		query = "select id from banned where phone = ? and state > 0"
		rows  *sqlx.Rows
	)
	rows, err = dao.db.Query(ctx, query, phone)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in CheckBannedByPhone(_), error: %v", err)
		return
	}

	defer rows.Close()

	do := &dataobject.BannedDO{}
	if rows.Next() {
		// TODO(@benqi): not use reflect
		err = rows.StructScan(do)
		if err != nil {
			logx.WithContext(ctx).Errorf("structScan in CheckBannedByPhone(_), error: %v", err)
			return
		} else {
			rValue = do
		}
	}

	return
}

// SelectPhoneList
// select phone from banned where state > 0 and phone in (:pList)
// TODO(@benqi): sqlmap
func (dao *BannedDAO) SelectPhoneList(ctx context.Context, pList []string) (rList []string, err error) {
	var (
		query = "select phone from banned where state > 0 and phone in (?)"
		a     []interface{}
	)
	if len(pList) == 0 {
		rList = []string{}
		return
	}

	query, a, err = sqlx.In(query, pList)
	if err != nil {
		// r sql.Result
		logx.WithContext(ctx).Errorf("sqlx.In in SelectPhoneList(_), error: %v", err)
		return
	}

	err = dao.db.Select(ctx, &rList, query, a...)

	if err != nil {
		logx.WithContext(ctx).Errorf("select in SelectPhoneList(_), error: %v", err)
	}

	return
}

// SelectPhoneListWithCB
// select phone from banned where state > 0 and phone in (:pList)
// TODO(@benqi): sqlmap
func (dao *BannedDAO) SelectPhoneListWithCB(ctx context.Context, pList []string, cb func(i int, v string)) (rList []string, err error) {
	var (
		query = "select phone from banned where state > 0 and phone in (?)"
		a     []interface{}
	)
	if len(pList) == 0 {
		rList = []string{}
		return
	}

	query, a, err = sqlx.In(query, pList)
	if err != nil {
		// r sql.Result
		logx.WithContext(ctx).Errorf("sqlx.In in SelectPhoneList(_), error: %v", err)
		return
	}

	err = dao.db.Select(ctx, &rList, query, a...)

	if err != nil {
		logx.WithContext(ctx).Errorf("select in SelectPhoneList(_), error: %v", err)
	}

	if cb != nil {
		for i := 0; i < len(rList); i++ {
			cb(i, rList[i])
		}
	}

	return
}

// Update
// update banned set expires = :expires, state = 0, log = :log, state = :state where phone = :phone
// TODO(@benqi): sqlmap
func (dao *BannedDAO) Update(ctx context.Context, expires int64, log string, state int32, phone string) (rowsAffected int64, err error) {
	var (
		query   = "update banned set expires = ?, state = 0, log = ?, state = ? where phone = ?"
		rResult sql.Result
	)
	rResult, err = dao.db.Exec(ctx, query, expires, log, state, phone)

	if err != nil {
		logx.WithContext(ctx).Errorf("exec in Update(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		logx.WithContext(ctx).Errorf("rowsAffected in Update(_), error: %v", err)
	}

	return
}

// update banned set expires = :expires, state = 0, log = :log, state = :state where phone = :phone
// UpdateTx
// TODO(@benqi): sqlmap
func (dao *BannedDAO) UpdateTx(tx *sqlx.Tx, expires int64, log string, state int32, phone string) (rowsAffected int64, err error) {
	var (
		query   = "update banned set expires = ?, state = 0, log = ?, state = ? where phone = ?"
		rResult sql.Result
	)
	rResult, err = tx.Exec(query, expires, log, state, phone)

	if err != nil {
		logx.WithContext(tx.Context()).Errorf("exec in Update(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		logx.WithContext(tx.Context()).Errorf("rowsAffected in Update(_), error: %v", err)
	}

	return
}
