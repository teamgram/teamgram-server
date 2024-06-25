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
	"github.com/teamgram/teamgram-server/app/service/authsession/internal/dal/dataobject"

	"github.com/zeromicro/go-zero/core/logx"
)

var _ *sql.Result
var _ = fmt.Sprintf
var _ = strings.Join
var _ = errors.Is

type AuthUsersDAO struct {
	db *sqlx.DB
}

func NewAuthUsersDAO(db *sqlx.DB) *AuthUsersDAO {
	return &AuthUsersDAO{
		db: db,
	}
}

// InsertOrUpdates
// insert into auth_users(auth_key_id, user_id, hash, date_created, date_actived) values (:auth_key_id, :user_id, :hash, :date_created, :date_actived) on duplicate key update hash = values(hash), date_actived = values(date_actived), deleted = 0
func (dao *AuthUsersDAO) InsertOrUpdates(ctx context.Context, do *dataobject.AuthUsersDO) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into auth_users(auth_key_id, user_id, hash, date_created, date_actived) values (:auth_key_id, :user_id, :hash, :date_created, :date_actived) on duplicate key update hash = values(hash), date_actived = values(date_actived), deleted = 0"
		r     sql.Result
	)

	r, err = dao.db.NamedExec(ctx, query, do)
	if err != nil {
		logx.WithContext(ctx).Errorf("namedExec in InsertOrUpdates(%v), error: %v", do, err)
		return
	}

	lastInsertId, err = r.LastInsertId()
	if err != nil {
		logx.WithContext(ctx).Errorf("lastInsertId in InsertOrUpdates(%v)_error: %v", do, err)
		return
	}
	rowsAffected, err = r.RowsAffected()
	if err != nil {
		logx.WithContext(ctx).Errorf("rowsAffected in InsertOrUpdates(%v)_error: %v", do, err)
	}

	return
}

// InsertOrUpdatesTx
// insert into auth_users(auth_key_id, user_id, hash, date_created, date_actived) values (:auth_key_id, :user_id, :hash, :date_created, :date_actived) on duplicate key update hash = values(hash), date_actived = values(date_actived), deleted = 0
func (dao *AuthUsersDAO) InsertOrUpdatesTx(tx *sqlx.Tx, do *dataobject.AuthUsersDO) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into auth_users(auth_key_id, user_id, hash, date_created, date_actived) values (:auth_key_id, :user_id, :hash, :date_created, :date_actived) on duplicate key update hash = values(hash), date_actived = values(date_actived), deleted = 0"
		r     sql.Result
	)

	r, err = tx.NamedExec(query, do)
	if err != nil {
		logx.WithContext(tx.Context()).Errorf("namedExec in InsertOrUpdates(%v), error: %v", do, err)
		return
	}

	lastInsertId, err = r.LastInsertId()
	if err != nil {
		logx.WithContext(tx.Context()).Errorf("lastInsertId in InsertOrUpdates(%v)_error: %v", do, err)
		return
	}
	rowsAffected, err = r.RowsAffected()
	if err != nil {
		logx.WithContext(tx.Context()).Errorf("rowsAffected in InsertOrUpdates(%v)_error: %v", do, err)
	}

	return
}

// Select
// select id, auth_key_id, user_id, hash, date_created, date_actived from auth_users where auth_key_id = :auth_key_id and deleted = 0
func (dao *AuthUsersDAO) Select(ctx context.Context, authKeyId int64) (rValue *dataobject.AuthUsersDO, err error) {
	var (
		query = "select id, auth_key_id, user_id, hash, date_created, date_actived from auth_users where auth_key_id = ? and deleted = 0"
		do    = &dataobject.AuthUsersDO{}
	)
	err = dao.db.QueryRowPartial(ctx, do, query, authKeyId)

	if err != nil {
		if !errors.Is(err, sqlx.ErrNotFound) {
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

// SelectAuthKeyIds
// select id, auth_key_id, user_id, hash, date_created, date_actived from auth_users where user_id = :user_id and deleted = 0
func (dao *AuthUsersDAO) SelectAuthKeyIds(ctx context.Context, userId int64) (rList []dataobject.AuthUsersDO, err error) {
	var (
		query  = "select id, auth_key_id, user_id, hash, date_created, date_actived from auth_users where user_id = ? and deleted = 0"
		values []dataobject.AuthUsersDO
	)
	err = dao.db.QueryRowsPartial(ctx, &values, query, userId)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectAuthKeyIds(_), error: %v", err)
		return
	}

	rList = values

	return
}

// SelectAuthKeyIdsWithCB
// select id, auth_key_id, user_id, hash, date_created, date_actived from auth_users where user_id = :user_id and deleted = 0
func (dao *AuthUsersDAO) SelectAuthKeyIdsWithCB(ctx context.Context, userId int64, cb func(sz, i int, v *dataobject.AuthUsersDO)) (rList []dataobject.AuthUsersDO, err error) {
	var (
		query  = "select id, auth_key_id, user_id, hash, date_created, date_actived from auth_users where user_id = ? and deleted = 0"
		values []dataobject.AuthUsersDO
	)
	err = dao.db.QueryRowsPartial(ctx, &values, query, userId)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectAuthKeyIds(_), error: %v", err)
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

// DeleteByHashList
// update auth_users set deleted = 1, date_created = 0, date_actived = 0 where id in (:idList)
func (dao *AuthUsersDAO) DeleteByHashList(ctx context.Context, idList []int64) (rowsAffected int64, err error) {
	var (
		query   = fmt.Sprintf("update auth_users set deleted = 1, date_created = 0, date_actived = 0 where id in (%s)", sqlx.InInt64List(idList))
		rResult sql.Result
	)

	if len(idList) == 0 {
		return
	}

	rResult, err = dao.db.Exec(ctx, query)

	if err != nil {
		logx.WithContext(ctx).Errorf("exec in DeleteByHashList(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		logx.WithContext(ctx).Errorf("rowsAffected in DeleteByHashList(_), error: %v", err)
	}

	return
}

// DeleteByHashListTx
// update auth_users set deleted = 1, date_created = 0, date_actived = 0 where id in (:idList)
func (dao *AuthUsersDAO) DeleteByHashListTx(tx *sqlx.Tx, idList []int64) (rowsAffected int64, err error) {
	var (
		query   = fmt.Sprintf("update auth_users set deleted = 1, date_created = 0, date_actived = 0 where id in (%s)", sqlx.InInt64List(idList))
		rResult sql.Result
	)

	if len(idList) == 0 {
		return
	}

	rResult, err = tx.Exec(query)

	if err != nil {
		logx.WithContext(tx.Context()).Errorf("exec in DeleteByHashList(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		logx.WithContext(tx.Context()).Errorf("rowsAffected in DeleteByHashList(_), error: %v", err)
	}

	return
}

// SelectListByUserId
// select id, auth_key_id, user_id, hash, date_created, date_actived from auth_users where user_id = :user_id and deleted = 0
func (dao *AuthUsersDAO) SelectListByUserId(ctx context.Context, userId int64) (rList []dataobject.AuthUsersDO, err error) {
	var (
		query  = "select id, auth_key_id, user_id, hash, date_created, date_actived from auth_users where user_id = ? and deleted = 0"
		values []dataobject.AuthUsersDO
	)
	err = dao.db.QueryRowsPartial(ctx, &values, query, userId)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectListByUserId(_), error: %v", err)
		return
	}

	rList = values

	return
}

// SelectListByUserIdWithCB
// select id, auth_key_id, user_id, hash, date_created, date_actived from auth_users where user_id = :user_id and deleted = 0
func (dao *AuthUsersDAO) SelectListByUserIdWithCB(ctx context.Context, userId int64, cb func(sz, i int, v *dataobject.AuthUsersDO)) (rList []dataobject.AuthUsersDO, err error) {
	var (
		query  = "select id, auth_key_id, user_id, hash, date_created, date_actived from auth_users where user_id = ? and deleted = 0"
		values []dataobject.AuthUsersDO
	)
	err = dao.db.QueryRowsPartial(ctx, &values, query, userId)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectListByUserId(_), error: %v", err)
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

// Delete
// update auth_users set deleted = 1, date_actived = 0 where auth_key_id = :auth_key_id and user_id = :user_id
func (dao *AuthUsersDAO) Delete(ctx context.Context, authKeyId int64, userId int64) (rowsAffected int64, err error) {
	var (
		query   = "update auth_users set deleted = 1, date_actived = 0 where auth_key_id = ? and user_id = ?"
		rResult sql.Result
	)

	rResult, err = dao.db.Exec(ctx, query, authKeyId, userId)

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
// update auth_users set deleted = 1, date_actived = 0 where auth_key_id = :auth_key_id and user_id = :user_id
func (dao *AuthUsersDAO) DeleteTx(tx *sqlx.Tx, authKeyId int64, userId int64) (rowsAffected int64, err error) {
	var (
		query   = "update auth_users set deleted = 1, date_actived = 0 where auth_key_id = ? and user_id = ?"
		rResult sql.Result
	)
	rResult, err = tx.Exec(query, authKeyId, userId)

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

// DeleteUser
// update auth_users set deleted = 1, date_actived = 0 where user_id = :user_id
func (dao *AuthUsersDAO) DeleteUser(ctx context.Context, userId int64) (rowsAffected int64, err error) {
	var (
		query   = "update auth_users set deleted = 1, date_actived = 0 where user_id = ?"
		rResult sql.Result
	)

	rResult, err = dao.db.Exec(ctx, query, userId)

	if err != nil {
		logx.WithContext(ctx).Errorf("exec in DeleteUser(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		logx.WithContext(ctx).Errorf("rowsAffected in DeleteUser(_), error: %v", err)
	}

	return
}

// DeleteUserTx
// update auth_users set deleted = 1, date_actived = 0 where user_id = :user_id
func (dao *AuthUsersDAO) DeleteUserTx(tx *sqlx.Tx, userId int64) (rowsAffected int64, err error) {
	var (
		query   = "update auth_users set deleted = 1, date_actived = 0 where user_id = ?"
		rResult sql.Result
	)
	rResult, err = tx.Exec(query, userId)

	if err != nil {
		logx.WithContext(tx.Context()).Errorf("exec in DeleteUser(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		logx.WithContext(tx.Context()).Errorf("rowsAffected in DeleteUser(_), error: %v", err)
	}

	return
}
