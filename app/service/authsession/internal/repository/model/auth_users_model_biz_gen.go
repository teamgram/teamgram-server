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
	bizAuthUsersModel interface {
		InsertOrUpdates(ctx context.Context, data *AuthUsers) (lastInsertId, rowsAffected int64, err error)
		InsertOrUpdatesTx(tx *sqlx.Tx, data *AuthUsers) (lastInsertId, rowsAffected int64, err error)

		Select(ctx context.Context, authKeyId int64) (*AuthUsers, error)

		UpdateAndroidPushSessionId(ctx context.Context, androidPushSessionId int64, authKeyId int64, userId int64) (rowsAffected int64, err error)
		UpdateAndroidPushSessionIdTx(tx *sqlx.Tx, androidPushSessionId int64, authKeyId int64, userId int64) (rowsAffected int64, err error)

		SelectAuthKeyIds(ctx context.Context, userId int64) ([]AuthUsers, error)
		SelectAuthKeyIdsWithCB(ctx context.Context, userId int64, cb func(sz, i int, v *AuthUsers)) ([]AuthUsers, error)

		DeleteByHashList(ctx context.Context, idList []int64) (rowsAffected int64, err error)
		DeleteByHashListTx(tx *sqlx.Tx, idList []int64) (rowsAffected int64, err error)

		SelectListByUserId(ctx context.Context, userId int64) ([]AuthUsers, error)
		SelectListByUserIdWithCB(ctx context.Context, userId int64, cb func(sz, i int, v *AuthUsers)) ([]AuthUsers, error)

		Delete(ctx context.Context, authKeyId int64, userId int64) (rowsAffected int64, err error)
		DeleteTx(tx *sqlx.Tx, authKeyId int64, userId int64) (rowsAffected int64, err error)

		DeleteUser(ctx context.Context, userId int64) (rowsAffected int64, err error)
		DeleteUserTx(tx *sqlx.Tx, userId int64) (rowsAffected int64, err error)
	}
)

// InsertOrUpdates
// insert into auth_users(auth_key_id, user_id, hash, date_created, date_active) values (:auth_key_id, :user_id, :hash, :date_created, :date_active) on duplicate key update hash = values(hash), date_active = values(date_active), deleted = 0
func (m *defaultAuthUsersModel) InsertOrUpdates(ctx context.Context, data *AuthUsers) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into auth_users(auth_key_id, user_id, hash, date_created, date_active) values (:auth_key_id, :user_id, :hash, :date_created, :date_active) on duplicate key update hash = values(hash), date_active = values(date_active), deleted = 0"
		r     sql.Result
	)

	r, err = m.db.NamedExec(ctx, query, data)
	if err != nil {
		err = fmt.Errorf("auth_users.InsertOrUpdates named exec: %w", err)
		return
	}

	lastInsertId, err = r.LastInsertId()
	if err != nil {
		err = fmt.Errorf("auth_users.InsertOrUpdates last insert id: %w", err)
		return
	}
	rowsAffected, err = r.RowsAffected()
	if err != nil {
		err = fmt.Errorf("auth_users.InsertOrUpdates rows affected: %w", err)
	}

	return

}

// InsertOrUpdatesTx
// insert into auth_users(auth_key_id, user_id, hash, date_created, date_active) values (:auth_key_id, :user_id, :hash, :date_created, :date_active) on duplicate key update hash = values(hash), date_active = values(date_active), deleted = 0
func (m *defaultAuthUsersModel) InsertOrUpdatesTx(tx *sqlx.Tx, data *AuthUsers) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into auth_users(auth_key_id, user_id, hash, date_created, date_active) values (:auth_key_id, :user_id, :hash, :date_created, :date_active) on duplicate key update hash = values(hash), date_active = values(date_active), deleted = 0"
		r     sql.Result
	)

	r, err = tx.NamedExec(query, data)
	if err != nil {
		err = fmt.Errorf("auth_users.InsertOrUpdatesTx named exec: %w", err)
		return
	}

	lastInsertId, err = r.LastInsertId()
	if err != nil {
		err = fmt.Errorf("auth_users.InsertOrUpdatesTx last insert id: %w", err)
		return
	}
	rowsAffected, err = r.RowsAffected()
	if err != nil {
		err = fmt.Errorf("auth_users.InsertOrUpdatesTx rows affected: %w", err)
	}

	return
}

// Select
// select id, auth_key_id, user_id, hash, date_created, date_active, android_push_session_id from auth_users where auth_key_id = :auth_key_id and deleted = 0
func (m *defaultAuthUsersModel) Select(ctx context.Context, authKeyId int64) (rValue *AuthUsers, err error) {

	var (
		query = "select id, auth_key_id, user_id, hash, date_created, date_active, android_push_session_id from auth_users where auth_key_id = ? and deleted = 0"
		do    = &AuthUsers{}
	)
	err = m.db.QueryRowPartial(ctx, do, query, authKeyId)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return nil, &NotFoundError{
				Resource: "auth_users",
				Key:      fmt.Sprintf("auth_key_id=%v", authKeyId),
				Cause:    err,
			}
		}
		err = fmt.Errorf("auth_users.Select: %w", err)
		return
	} else {
		rValue = do
	}

	return
}

// UpdateAndroidPushSessionId
// update auth_users set android_push_session_id = :android_push_session_id where auth_key_id = :auth_key_id and user_id = :user_id
func (m *defaultAuthUsersModel) UpdateAndroidPushSessionId(ctx context.Context, androidPushSessionId int64, authKeyId int64, userId int64) (rowsAffected int64, err error) {

	var (
		query   = "update auth_users set android_push_session_id = ? where auth_key_id = ? and user_id = ?"
		rResult sql.Result
	)

	rResult, err = m.db.Exec(ctx, query, androidPushSessionId, authKeyId, userId)

	if err != nil {
		err = fmt.Errorf("auth_users.UpdateAndroidPushSessionId exec: %w", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		err = fmt.Errorf("auth_users.UpdateAndroidPushSessionId rows affected: %w", err)
		return
	}

	return
}

// UpdateAndroidPushSessionIdTx
// update auth_users set android_push_session_id = :android_push_session_id where auth_key_id = :auth_key_id and user_id = :user_id
func (m *defaultAuthUsersModel) UpdateAndroidPushSessionIdTx(tx *sqlx.Tx, androidPushSessionId int64, authKeyId int64, userId int64) (rowsAffected int64, err error) {
	var (
		query   = "update auth_users set android_push_session_id = ? where auth_key_id = ? and user_id = ?"
		rResult sql.Result
	)
	rResult, err = tx.Exec(query, androidPushSessionId, authKeyId, userId)

	if err != nil {
		err = fmt.Errorf("auth_users.UpdateAndroidPushSessionIdTx exec: %w", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		err = fmt.Errorf("auth_users.UpdateAndroidPushSessionIdTx rows affected: %w", err)
		return
	}

	return
}

// SelectAuthKeyIds
// select id, auth_key_id, user_id, hash, date_created, date_active, android_push_session_id from auth_users where user_id = :user_id and deleted = 0
func (m *defaultAuthUsersModel) SelectAuthKeyIds(ctx context.Context, userId int64) (rList []AuthUsers, err error) {
	var (
		query  = "select id, auth_key_id, user_id, hash, date_created, date_active, android_push_session_id from auth_users where user_id = ? and deleted = 0"
		values []AuthUsers
	)
	err = m.db.QueryRowsPartial(ctx, &values, query, userId)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			rList = []AuthUsers{}
			err = nil
			return
		}
		err = fmt.Errorf("auth_users.SelectAuthKeyIds: %w", err)
		return
	}

	rList = values

	return
}

// SelectAuthKeyIdsWithCB
// select id, auth_key_id, user_id, hash, date_created, date_active, android_push_session_id from auth_users where user_id = :user_id and deleted = 0
func (m *defaultAuthUsersModel) SelectAuthKeyIdsWithCB(ctx context.Context, userId int64, cb func(sz, i int, v *AuthUsers)) (rList []AuthUsers, err error) {
	var (
		query  = "select id, auth_key_id, user_id, hash, date_created, date_active, android_push_session_id from auth_users where user_id = ? and deleted = 0"
		values []AuthUsers
	)
	err = m.db.QueryRowsPartial(ctx, &values, query, userId)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			rList = []AuthUsers{}
			err = nil
			return
		}
		err = fmt.Errorf("auth_users.SelectAuthKeyIdsWithCB: %w", err)
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
// update auth_users set deleted = 1, date_created = 0, date_active = 0 where id in (:idList)
func (m *defaultAuthUsersModel) DeleteByHashList(ctx context.Context, idList []int64) (rowsAffected int64, err error) {

	var (
		query   = fmt.Sprintf("update auth_users set deleted = 1, date_created = 0, date_active = 0 where id in (%s)", sqlx.InInt64List(idList))
		rResult sql.Result
	)

	if len(idList) == 0 {
		return
	}

	rResult, err = m.db.Exec(ctx, query)

	if err != nil {
		err = fmt.Errorf("auth_users.DeleteByHashList exec: %w", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		err = fmt.Errorf("auth_users.DeleteByHashList rows affected: %w", err)
		return
	}

	return
}

// DeleteByHashListTx
// update auth_users set deleted = 1, date_created = 0, date_active = 0 where id in (:idList)
func (m *defaultAuthUsersModel) DeleteByHashListTx(tx *sqlx.Tx, idList []int64) (rowsAffected int64, err error) {
	var (
		query   = fmt.Sprintf("update auth_users set deleted = 1, date_created = 0, date_active = 0 where id in (%s)", sqlx.InInt64List(idList))
		rResult sql.Result
	)

	if len(idList) == 0 {
		return
	}

	rResult, err = tx.Exec(query)

	if err != nil {
		err = fmt.Errorf("auth_users.DeleteByHashListTx exec: %w", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		err = fmt.Errorf("auth_users.DeleteByHashListTx rows affected: %w", err)
		return
	}

	return
}

// SelectListByUserId
// select id, auth_key_id, user_id, hash, date_created, date_active, android_push_session_id from auth_users where user_id = :user_id and deleted = 0
func (m *defaultAuthUsersModel) SelectListByUserId(ctx context.Context, userId int64) (rList []AuthUsers, err error) {
	var (
		query  = "select id, auth_key_id, user_id, hash, date_created, date_active, android_push_session_id from auth_users where user_id = ? and deleted = 0"
		values []AuthUsers
	)
	err = m.db.QueryRowsPartial(ctx, &values, query, userId)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			rList = []AuthUsers{}
			err = nil
			return
		}
		err = fmt.Errorf("auth_users.SelectListByUserId: %w", err)
		return
	}

	rList = values

	return
}

// SelectListByUserIdWithCB
// select id, auth_key_id, user_id, hash, date_created, date_active, android_push_session_id from auth_users where user_id = :user_id and deleted = 0
func (m *defaultAuthUsersModel) SelectListByUserIdWithCB(ctx context.Context, userId int64, cb func(sz, i int, v *AuthUsers)) (rList []AuthUsers, err error) {
	var (
		query  = "select id, auth_key_id, user_id, hash, date_created, date_active, android_push_session_id from auth_users where user_id = ? and deleted = 0"
		values []AuthUsers
	)
	err = m.db.QueryRowsPartial(ctx, &values, query, userId)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			rList = []AuthUsers{}
			err = nil
			return
		}
		err = fmt.Errorf("auth_users.SelectListByUserIdWithCB: %w", err)
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
// update auth_users set deleted = 1, date_active = 0 where auth_key_id = :auth_key_id and user_id = :user_id
func (m *defaultAuthUsersModel) Delete(ctx context.Context, authKeyId int64, userId int64) (rowsAffected int64, err error) {

	var (
		query   = "update auth_users set deleted = 1, date_active = 0 where auth_key_id = ? and user_id = ?"
		rResult sql.Result
	)

	rResult, err = m.db.Exec(ctx, query, authKeyId, userId)

	if err != nil {
		err = fmt.Errorf("auth_users.Delete exec: %w", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		err = fmt.Errorf("auth_users.Delete rows affected: %w", err)
		return
	}

	return
}

// DeleteTx
// update auth_users set deleted = 1, date_active = 0 where auth_key_id = :auth_key_id and user_id = :user_id
func (m *defaultAuthUsersModel) DeleteTx(tx *sqlx.Tx, authKeyId int64, userId int64) (rowsAffected int64, err error) {
	var (
		query   = "update auth_users set deleted = 1, date_active = 0 where auth_key_id = ? and user_id = ?"
		rResult sql.Result
	)
	rResult, err = tx.Exec(query, authKeyId, userId)

	if err != nil {
		err = fmt.Errorf("auth_users.DeleteTx exec: %w", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		err = fmt.Errorf("auth_users.DeleteTx rows affected: %w", err)
		return
	}

	return
}

// DeleteUser
// update auth_users set deleted = 1, date_active = 0 where user_id = :user_id
func (m *defaultAuthUsersModel) DeleteUser(ctx context.Context, userId int64) (rowsAffected int64, err error) {

	var (
		query   = "update auth_users set deleted = 1, date_active = 0 where user_id = ?"
		rResult sql.Result
	)

	rResult, err = m.db.Exec(ctx, query, userId)

	if err != nil {
		err = fmt.Errorf("auth_users.DeleteUser exec: %w", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		err = fmt.Errorf("auth_users.DeleteUser rows affected: %w", err)
		return
	}

	return
}

// DeleteUserTx
// update auth_users set deleted = 1, date_active = 0 where user_id = :user_id
func (m *defaultAuthUsersModel) DeleteUserTx(tx *sqlx.Tx, userId int64) (rowsAffected int64, err error) {
	var (
		query   = "update auth_users set deleted = 1, date_active = 0 where user_id = ?"
		rResult sql.Result
	)
	rResult, err = tx.Exec(query, userId)

	if err != nil {
		err = fmt.Errorf("auth_users.DeleteUserTx exec: %w", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		err = fmt.Errorf("auth_users.DeleteUserTx rows affected: %w", err)
		return
	}

	return
}
