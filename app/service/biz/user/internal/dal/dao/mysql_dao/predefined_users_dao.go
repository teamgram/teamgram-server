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
	"errors"
	"fmt"
	"strings"

	"github.com/teamgram/marmota/pkg/stores/sqlx"
	"github.com/teamgram/teamgram-server/app/service/biz/user/internal/dal/dataobject"

	"github.com/zeromicro/go-zero/core/logx"
)

type PredefinedUsersDAO struct {
	db *sqlx.DB
}

func NewPredefinedUsersDAO(db *sqlx.DB) *PredefinedUsersDAO {
	return &PredefinedUsersDAO{
		db: db,
	}
}

// Insert
// insert into predefined_users(first_name, last_name, username, phone, code, verified) values (:first_name, :last_name, :username, :phone, :code, :verified)
func (dao *PredefinedUsersDAO) Insert(ctx context.Context, do *dataobject.PredefinedUsersDO) (lastInsertId, rowsAffected int64, err error) {
	var (
		query string
		r     sql.Result
	)
	query = "insert into predefined_users(first_name, last_name, username, phone, code, verified) values (:first_name, :last_name, :username, :phone, :code, :verified)"

	r, err = dao.db.NamedExec(ctx, query, do)
	if err != nil {
		logx.WithContext(ctx).Errorf("namedExec in Insert(%v), error: %v", do, err)
		return
	}

	lastInsertId, err = r.LastInsertId()
	if err != nil {
		logx.WithContext(ctx).Errorf("lastInsertId in Insert(%v), error: %v", do, err)
		return
	}
	rowsAffected, err = r.RowsAffected()
	if err != nil {
		logx.WithContext(ctx).Errorf("rowsAffected in Insert(%v), error: %v", do, err)
	}

	return
}

// InsertTx
// insert into predefined_users(first_name, last_name, username, phone, code, verified) values (:first_name, :last_name, :username, :phone, :code, :verified)
func (dao *PredefinedUsersDAO) InsertTx(tx *sqlx.Tx, do *dataobject.PredefinedUsersDO) (lastInsertId, rowsAffected int64, err error) {
	var (
		query string
		r     sql.Result
	)
	query = "insert into predefined_users(first_name, last_name, username, phone, code, verified) values (:first_name, :last_name, :username, :phone, :code, :verified)"

	r, err = tx.NamedExec(query, do)
	if err != nil {
		logx.WithContext(tx.Context()).Errorf("namedExec in Insert(%v), error: %v", do, err)
		return
	}

	lastInsertId, err = r.LastInsertId()
	if err != nil {
		logx.WithContext(tx.Context()).Errorf("lastInsertId in Insert(%v), error: %v", do, err)
		return
	}
	rowsAffected, err = r.RowsAffected()
	if err != nil {
		logx.WithContext(tx.Context()).Errorf("rowsAffected in Insert(%v), error: %v", do, err)
	}

	return
}

// SelectByPhone
// select id, phone, first_name, last_name, username, code, verified, registered_user_id from predefined_users where phone = :phone and deleted = 0 limit 1
func (dao *PredefinedUsersDAO) SelectByPhone(ctx context.Context, phone string) (rValue *dataobject.PredefinedUsersDO, err error) {
	var (
		query string
		do    = &dataobject.PredefinedUsersDO{}
	)
	query = "select id, phone, first_name, last_name, username, code, verified, registered_user_id from predefined_users where phone = ? and deleted = 0 limit 1"

	err = dao.db.QueryRowPartial(ctx, do, query, phone)

	if err != nil {
		if !errors.Is(err, sqlx.ErrNotFound) {
			logx.WithContext(ctx).Errorf("queryx in SelectByPhone(_), error: %v", err)
			return
		} else {
			// not found not error, return nil, nil
			err = nil
		}
	} else {
		rValue = do
	}

	return
}

// SelectPredefinedUsersAll
// select id, phone, first_name, last_name, username, code, verified, registered_user_id from predefined_users where deleted = 0 order by username asc
func (dao *PredefinedUsersDAO) SelectPredefinedUsersAll(ctx context.Context) (rList []dataobject.PredefinedUsersDO, err error) {
	var (
		query  string
		values []dataobject.PredefinedUsersDO
	)
	query = "select id, phone, first_name, last_name, username, code, verified, registered_user_id from predefined_users where deleted = 0 order by username asc"

	err = dao.db.QueryRowsPartial(ctx, &values, query)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectPredefinedUsersAll(_), error: %v", err)
		return
	}

	rList = values

	return
}

// SelectPredefinedUsersAllWithCB
// select id, phone, first_name, last_name, username, code, verified, registered_user_id from predefined_users where deleted = 0 order by username asc
func (dao *PredefinedUsersDAO) SelectPredefinedUsersAllWithCB(ctx context.Context, cb func(sz, i int, v *dataobject.PredefinedUsersDO)) (rList []dataobject.PredefinedUsersDO, err error) {
	var (
		query  string
		values []dataobject.PredefinedUsersDO
	)
	query = "select id, phone, first_name, last_name, username, code, verified, registered_user_id from predefined_users where deleted = 0 order by username asc"

	err = dao.db.QueryRowsPartial(ctx, &values, query)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectPredefinedUsersAll(_), error: %v", err)
		return
	}

	rList = values

	if cb != nil {
		sz := len(rList)
		for i := range sz {
			cb(sz, i, &rList[i])
		}
	}

	return
}

// Delete
// update predefined_users set deleted = 0 where phone = :phone
func (dao *PredefinedUsersDAO) Delete(ctx context.Context, phone string) (rowsAffected int64, err error) {
	var (
		query   string
		rResult sql.Result
	)
	query = "update predefined_users set deleted = 0 where phone = ?"

	rResult, err = dao.db.Exec(ctx, query, phone)

	if err != nil {
		logx.WithContext(ctx).Errorf("exec in Delete(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		logx.WithContext(ctx).Errorf("rowsAffected in Delete(_), error: %v", err)
	}

	return
}

// DeleteTx
// update predefined_users set deleted = 0 where phone = :phone
func (dao *PredefinedUsersDAO) DeleteTx(tx *sqlx.Tx, phone string) (rowsAffected int64, err error) {
	var (
		query   string
		rResult sql.Result
	)
	query = "update predefined_users set deleted = 0 where phone = ?"

	rResult, err = tx.Exec(query, phone)

	if err != nil {
		logx.WithContext(tx.Context()).Errorf("exec in Delete(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		logx.WithContext(tx.Context()).Errorf("rowsAffected in Delete(_), error: %v", err)
	}

	return
}

// Update
// update predefined_users set %s where phone = :phone
func (dao *PredefinedUsersDAO) Update(ctx context.Context, cMap map[string]interface{}, phone string) (rowsAffected int64, err error) {
	names := make([]string, 0, len(cMap))
	aValues := make([]interface{}, 0, len(cMap))
	for k, v := range cMap {
		names = append(names, k+" = ?")
		aValues = append(aValues, v)
	}

	var (
		query   string
		rResult sql.Result
	)
	query = fmt.Sprintf("update predefined_users set %s where phone = ?", strings.Join(names, ", "))

	aValues = append(aValues, phone)

	rResult, err = dao.db.Exec(ctx, query, aValues...)

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

// UpdateTx
// update predefined_users set %s where phone = :phone
func (dao *PredefinedUsersDAO) UpdateTx(tx *sqlx.Tx, cMap map[string]interface{}, phone string) (rowsAffected int64, err error) {
	names := make([]string, 0, len(cMap))
	aValues := make([]interface{}, 0, len(cMap))
	for k, v := range cMap {
		names = append(names, k+" = ?")
		aValues = append(aValues, v)
	}

	var (
		query   string
		rResult sql.Result
	)
	query = fmt.Sprintf("update predefined_users set %s where phone = ?", strings.Join(names, ", "))

	aValues = append(aValues, phone)

	rResult, err = tx.Exec(query, aValues...)

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
