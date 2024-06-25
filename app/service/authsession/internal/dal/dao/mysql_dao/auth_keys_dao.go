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

type AuthKeysDAO struct {
	db *sqlx.DB
}

func NewAuthKeysDAO(db *sqlx.DB) *AuthKeysDAO {
	return &AuthKeysDAO{
		db: db,
	}
}

// Insert
// insert into auth_keys(auth_key_id, body) values (:auth_key_id, :body)
func (dao *AuthKeysDAO) Insert(ctx context.Context, do *dataobject.AuthKeysDO) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into auth_keys(auth_key_id, body) values (:auth_key_id, :body)"
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
// insert into auth_keys(auth_key_id, body) values (:auth_key_id, :body)
func (dao *AuthKeysDAO) InsertTx(tx *sqlx.Tx, do *dataobject.AuthKeysDO) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into auth_keys(auth_key_id, body) values (:auth_key_id, :body)"
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

// SelectByAuthKeyId
// select auth_key_id, body from auth_keys where auth_key_id = :auth_key_id
func (dao *AuthKeysDAO) SelectByAuthKeyId(ctx context.Context, authKeyId int64) (rValue *dataobject.AuthKeysDO, err error) {
	var (
		query = "select auth_key_id, body from auth_keys where auth_key_id = ?"
		do    = &dataobject.AuthKeysDO{}
	)
	err = dao.db.QueryRowPartial(ctx, do, query, authKeyId)

	if err != nil {
		if !errors.Is(err, sqlx.ErrNotFound) {
			logx.WithContext(ctx).Errorf("queryx in SelectByAuthKeyId(_), error: %v", err)
			return
		} else {
			err = nil
		}
	} else {
		rValue = do
	}

	return
}
