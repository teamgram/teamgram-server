/*
 * WARNING! All changes made in this file will be lost!
 *   Created from by 'dalgen'
 *
 * Copyright (c) 2026 The Teamgram Authors.
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
	"github.com/zeromicro/go-zero/core/logx"
)

var _ *sql.Result
var _ = fmt.Sprintf
var _ = strings.Join
var _ = errors.Is
var _ *sqlx.DB
var _ *logx.Logger

type (
	bizAuthKeysModel interface {
		Insert(ctx context.Context, data *AuthKeys) (lastInsertId, rowsAffected int64, err error)
		InsertTx(tx *sqlx.Tx, data *AuthKeys) (lastInsertId, rowsAffected int64, err error)

		SelectByAuthKeyId(ctx context.Context, authKeyId int64) (*AuthKeys, error)
	}
)

// Insert
// insert into auth_keys(auth_key_id, body) values (:auth_key_id, :body)
func (m *defaultAuthKeysModel) Insert(ctx context.Context, data *AuthKeys) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into auth_keys(auth_key_id, body) values (:auth_key_id, :body)"
		r     sql.Result
	)

	r, err = m.db.NamedExec(ctx, query, data)
	if err != nil {
		logx.WithContext(ctx).Errorf("namedExec in Insert(%v), error: %v", data, err)
		return
	}

	lastInsertId, err = r.LastInsertId()
	if err != nil {
		logx.WithContext(ctx).Errorf("lastInsertId in Insert(%v)_error: %v", data, err)
		return
	}
	rowsAffected, err = r.RowsAffected()
	if err != nil {
		logx.WithContext(ctx).Errorf("rowsAffected in Insert(%v)_error: %v", data, err)
	}

	return
}

// InsertTx
// insert into auth_keys(auth_key_id, body) values (:auth_key_id, :body)
func (m *defaultAuthKeysModel) InsertTx(tx *sqlx.Tx, data *AuthKeys) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into auth_keys(auth_key_id, body) values (:auth_key_id, :body)"
		r     sql.Result
	)

	r, err = tx.NamedExec(query, data)
	if err != nil {
		logx.WithContext(tx.Context()).Errorf("namedExec in Insert(%v), error: %v", data, err)
		return
	}

	lastInsertId, err = r.LastInsertId()
	if err != nil {
		logx.WithContext(tx.Context()).Errorf("lastInsertId in Insert(%v)_error: %v", data, err)
		return
	}
	rowsAffected, err = r.RowsAffected()
	if err != nil {
		logx.WithContext(tx.Context()).Errorf("rowsAffected in Insert(%v)_error: %v", data, err)
	}

	return
}

// SelectByAuthKeyId
// select auth_key_id, body from auth_keys where auth_key_id = :auth_key_id
func (m *defaultAuthKeysModel) SelectByAuthKeyId(ctx context.Context, authKeyId int64) (rValue *AuthKeys, err error) {
	var (
		query = "select auth_key_id, body from auth_keys where auth_key_id = ?"
		do    = &AuthKeys{}
	)
	err = m.db.QueryRowPartial(ctx, do, query, authKeyId)

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
