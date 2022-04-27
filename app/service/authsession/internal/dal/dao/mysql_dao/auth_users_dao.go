/*
 * WARNING! All changes made in this file will be lost!
 *   Created from by 'dalgen'
 *
 * Copyright (c) 2022-present,  Teamgram Authors.
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package mysql_dao

import (
	"context"
	"database/sql"

	"github.com/teamgram/marmota/pkg/stores/sqlx"
	"github.com/teamgram/teamgram-server/app/service/authsession/internal/dal/dataobject"

	"github.com/zeromicro/go-zero/core/logx"
)

var _ *sql.Result

type AuthUsersDAO struct {
	db *sqlx.DB
}

func NewAuthUsersDAO(db *sqlx.DB) *AuthUsersDAO {
	return &AuthUsersDAO{db}
}

// InsertOrUpdates
// insert into auth_users(auth_key_id, user_id, hash, date_created, date_actived) values (:auth_key_id, :user_id, :hash, :date_created, :date_actived) on duplicate key update hash = values(hash), date_actived = values(date_actived), deleted = 0
// TODO(@benqi): sqlmap
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
// TODO(@benqi): sqlmap
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
// TODO(@benqi): sqlmap
func (dao *AuthUsersDAO) Select(ctx context.Context, auth_key_id int64) (rValue *dataobject.AuthUsersDO, err error) {
	var (
		query = "select id, auth_key_id, user_id, hash, date_created, date_actived from auth_users where auth_key_id = ? and deleted = 0"
		do    = &dataobject.AuthUsersDO{}
	)
	err = dao.db.QueryRowPartial(ctx, do, query, auth_key_id)

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

// SelectAuthKeyIds
// select id, auth_key_id, user_id, hash from auth_users where user_id = :user_id and deleted = 0
// TODO(@benqi): sqlmap
func (dao *AuthUsersDAO) SelectAuthKeyIds(ctx context.Context, user_id int64) (rList []dataobject.AuthUsersDO, err error) {
	var (
		query  = "select id, auth_key_id, user_id, hash from auth_users where user_id = ? and deleted = 0"
		values []dataobject.AuthUsersDO
	)
	err = dao.db.QueryRowsPartial(ctx, &values, query, user_id)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectAuthKeyIds(_), error: %v", err)
		return
	}

	rList = values

	return
}

// SelectAuthKeyIdsWithCB
// select id, auth_key_id, user_id, hash from auth_users where user_id = :user_id and deleted = 0
// TODO(@benqi): sqlmap
func (dao *AuthUsersDAO) SelectAuthKeyIdsWithCB(ctx context.Context, user_id int64, cb func(i int, v *dataobject.AuthUsersDO)) (rList []dataobject.AuthUsersDO, err error) {
	var (
		query  = "select id, auth_key_id, user_id, hash from auth_users where user_id = ? and deleted = 0"
		values []dataobject.AuthUsersDO
	)
	err = dao.db.QueryRowsPartial(ctx, &values, query, user_id)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectAuthKeyIds(_), error: %v", err)
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

// DeleteByHashList
// update auth_users set deleted = 1, date_created = 0, date_actived = 0 where id in (:idList)
// TODO(@benqi): sqlmap
func (dao *AuthUsersDAO) DeleteByHashList(ctx context.Context, idList []int64) (rowsAffected int64, err error) {
	var (
		query   = "update auth_users set deleted = 1, date_created = 0, date_actived = 0 where id in (?)"
		a       []interface{}
		rResult sql.Result
	)

	if len(idList) == 0 {
		return
	}

	query, a, err = sqlx.In(query, idList)
	if err != nil {
		// r sql.Result
		logx.WithContext(ctx).Errorf("sqlx.In in DeleteByHashList(_), error: %v", err)
		return
	}
	rResult, err = dao.db.Exec(ctx, query, a...)

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

// update auth_users set deleted = 1, date_created = 0, date_actived = 0 where id in (:idList)
// DeleteByHashListTx
// TODO(@benqi): sqlmap
func (dao *AuthUsersDAO) DeleteByHashListTx(tx *sqlx.Tx, idList []int64) (rowsAffected int64, err error) {
	var (
		query   = "update auth_users set deleted = 1, date_created = 0, date_actived = 0 where id in (?)"
		a       []interface{}
		rResult sql.Result
	)

	if len(idList) == 0 {
		return
	}

	query, a, err = sqlx.In(query, idList)
	if err != nil {
		// r sql.Result
		logx.WithContext(tx.Context()).Errorf("sqlx.In in DeleteByHashList(_), error: %v", err)
		return
	}
	rResult, err = tx.Exec(query, a...)

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
// select id, auth_key_id, user_id, hash from auth_users where user_id = :user_id and deleted = 0
// TODO(@benqi): sqlmap
func (dao *AuthUsersDAO) SelectListByUserId(ctx context.Context, user_id int64) (rList []dataobject.AuthUsersDO, err error) {
	var (
		query  = "select id, auth_key_id, user_id, hash from auth_users where user_id = ? and deleted = 0"
		values []dataobject.AuthUsersDO
	)
	err = dao.db.QueryRowsPartial(ctx, &values, query, user_id)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectListByUserId(_), error: %v", err)
		return
	}

	rList = values

	return
}

// SelectListByUserIdWithCB
// select id, auth_key_id, user_id, hash from auth_users where user_id = :user_id and deleted = 0
// TODO(@benqi): sqlmap
func (dao *AuthUsersDAO) SelectListByUserIdWithCB(ctx context.Context, user_id int64, cb func(i int, v *dataobject.AuthUsersDO)) (rList []dataobject.AuthUsersDO, err error) {
	var (
		query  = "select id, auth_key_id, user_id, hash from auth_users where user_id = ? and deleted = 0"
		values []dataobject.AuthUsersDO
	)
	err = dao.db.QueryRowsPartial(ctx, &values, query, user_id)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectListByUserId(_), error: %v", err)
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

// Delete
// update auth_users set deleted = 1, date_actived = 0 where auth_key_id = :auth_key_id and user_id = :user_id
// TODO(@benqi): sqlmap
func (dao *AuthUsersDAO) Delete(ctx context.Context, auth_key_id int64, user_id int64) (rowsAffected int64, err error) {
	var (
		query   = "update auth_users set deleted = 1, date_actived = 0 where auth_key_id = ? and user_id = ?"
		rResult sql.Result
	)
	rResult, err = dao.db.Exec(ctx, query, auth_key_id, user_id)

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

// update auth_users set deleted = 1, date_actived = 0 where auth_key_id = :auth_key_id and user_id = :user_id
// DeleteTx
// TODO(@benqi): sqlmap
func (dao *AuthUsersDAO) DeleteTx(tx *sqlx.Tx, auth_key_id int64, user_id int64) (rowsAffected int64, err error) {
	var (
		query   = "update auth_users set deleted = 1, date_actived = 0 where auth_key_id = ? and user_id = ?"
		rResult sql.Result
	)
	rResult, err = tx.Exec(query, auth_key_id, user_id)

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
// TODO(@benqi): sqlmap
func (dao *AuthUsersDAO) DeleteUser(ctx context.Context, user_id int64) (rowsAffected int64, err error) {
	var (
		query   = "update auth_users set deleted = 1, date_actived = 0 where user_id = ?"
		rResult sql.Result
	)
	rResult, err = dao.db.Exec(ctx, query, user_id)

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

// update auth_users set deleted = 1, date_actived = 0 where user_id = :user_id
// DeleteUserTx
// TODO(@benqi): sqlmap
func (dao *AuthUsersDAO) DeleteUserTx(tx *sqlx.Tx, user_id int64) (rowsAffected int64, err error) {
	var (
		query   = "update auth_users set deleted = 1, date_actived = 0 where user_id = ?"
		rResult sql.Result
	)
	rResult, err = tx.Exec(query, user_id)

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
