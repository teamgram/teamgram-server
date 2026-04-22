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
		InsertIgnore(ctx context.Context, data *AuthKeys) (lastInsertId, rowsAffected int64, err error)
		InsertIgnoreTx(tx *sqlx.Tx, data *AuthKeys) (lastInsertId, rowsAffected int64, err error)

		SelectByAuthKeyId(ctx context.Context, authKeyId int64) (*AuthKeys, error)

		UpdateCustomMap(ctx context.Context, cMap map[string]interface{}, authKeyId int64) (rowsAffected int64, err error)
		UpdateCustomMapTx(tx *sqlx.Tx, cMap map[string]interface{}, authKeyId int64) (rowsAffected int64, err error)
	}
)

// InsertIgnore
// insert ignore into auth_keys(auth_key_id, body, auth_key_type, perm_auth_key_id, temp_auth_key_id, media_temp_auth_key_id) values (:auth_key_id, :body, :auth_key_type, :perm_auth_key_id, :temp_auth_key_id, :media_temp_auth_key_id)
func (m *defaultAuthKeysModel) InsertIgnore(ctx context.Context, data *AuthKeys) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert ignore into auth_keys(auth_key_id, body, auth_key_type, perm_auth_key_id, temp_auth_key_id, media_temp_auth_key_id) values (:auth_key_id, :body, :auth_key_type, :perm_auth_key_id, :temp_auth_key_id, :media_temp_auth_key_id)"
		r     sql.Result
	)

	r, err = m.db.NamedExec(ctx, query, data)
	if err != nil {
		logx.WithContext(ctx).Errorf("namedExec in InsertIgnore(%v), error: %v", data, err)
		return
	}

	lastInsertId, err = r.LastInsertId()
	if err != nil {
		logx.WithContext(ctx).Errorf("lastInsertId in InsertIgnore(%v)_error: %v", data, err)
		return
	}
	rowsAffected, err = r.RowsAffected()
	if err != nil {
		logx.WithContext(ctx).Errorf("rowsAffected in InsertIgnore(%v)_error: %v", data, err)
	}

	return
}

// InsertIgnoreTx
// insert ignore into auth_keys(auth_key_id, body, auth_key_type, perm_auth_key_id, temp_auth_key_id, media_temp_auth_key_id) values (:auth_key_id, :body, :auth_key_type, :perm_auth_key_id, :temp_auth_key_id, :media_temp_auth_key_id)
func (m *defaultAuthKeysModel) InsertIgnoreTx(tx *sqlx.Tx, data *AuthKeys) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert ignore into auth_keys(auth_key_id, body, auth_key_type, perm_auth_key_id, temp_auth_key_id, media_temp_auth_key_id) values (:auth_key_id, :body, :auth_key_type, :perm_auth_key_id, :temp_auth_key_id, :media_temp_auth_key_id)"
		r     sql.Result
	)

	r, err = tx.NamedExec(query, data)
	if err != nil {
		logx.WithContext(tx.Context()).Errorf("namedExec in InsertIgnore(%v), error: %v", data, err)
		return
	}

	lastInsertId, err = r.LastInsertId()
	if err != nil {
		logx.WithContext(tx.Context()).Errorf("lastInsertId in InsertIgnore(%v)_error: %v", data, err)
		return
	}
	rowsAffected, err = r.RowsAffected()
	if err != nil {
		logx.WithContext(tx.Context()).Errorf("rowsAffected in InsertIgnore(%v)_error: %v", data, err)
	}

	return
}

// SelectByAuthKeyId
// select id, auth_key_id, body, auth_key_type, perm_auth_key_id, temp_auth_key_id, media_temp_auth_key_id from auth_keys where auth_key_id = :auth_key_id
func (m *defaultAuthKeysModel) SelectByAuthKeyId(ctx context.Context, authKeyId int64) (rValue *AuthKeys, err error) {
	var (
		query = "select id, auth_key_id, body, auth_key_type, perm_auth_key_id, temp_auth_key_id, media_temp_auth_key_id from auth_keys where auth_key_id = ?"
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

// UpdateCustomMap
// update auth_key_infos set %s where auth_key_id = :auth_key_id
func (m *defaultAuthKeysModel) UpdateCustomMap(ctx context.Context, cMap map[string]interface{}, authKeyId int64) (rowsAffected int64, err error) {
	names := make([]string, 0, len(cMap))
	aValues := make([]interface{}, 0, len(cMap))
	for k, v := range cMap {
		names = append(names, k+" = ?")
		aValues = append(aValues, v)
	}

	var (
		query   = fmt.Sprintf("update auth_key_infos set %s where auth_key_id = ?", strings.Join(names, ", "))
		rResult sql.Result
	)

	aValues = append(aValues, authKeyId)

	rResult, err = m.db.Exec(ctx, query, aValues...)

	if err != nil {
		logx.WithContext(ctx).Errorf("exec in UpdateCustomMap(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		logx.WithContext(ctx).Errorf("rowsAffected in UpdateCustomMap(_), error: %v", err)
	}

	return
}

// UpdateCustomMapTx
// update auth_key_infos set %s where auth_key_id = :auth_key_id
func (m *defaultAuthKeysModel) UpdateCustomMapTx(tx *sqlx.Tx, cMap map[string]interface{}, authKeyId int64) (rowsAffected int64, err error) {
	names := make([]string, 0, len(cMap))
	aValues := make([]interface{}, 0, len(cMap))
	for k, v := range cMap {
		names = append(names, k+" = ?")
		aValues = append(aValues, v)
	}

	var (
		query   = fmt.Sprintf("update auth_key_infos set %s where auth_key_id = ?", strings.Join(names, ", "))
		rResult sql.Result
	)

	aValues = append(aValues, authKeyId)

	rResult, err = tx.Exec(query, aValues...)

	if err != nil {
		logx.WithContext(tx.Context()).Errorf("exec in UpdateCustomMap(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		logx.WithContext(tx.Context()).Errorf("rowsAffected in UpdateCustomMap(_), error: %v", err)
	}

	return
}
