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

	keys := m.uniqueCacheKeys(data)
	lastInsertId, rowsAffected, err = m.Exec(ctx, func(ctx context.Context, conn *sqlx.DB) (int64, int64, error) {
		r, err = conn.NamedExec(ctx, query, data)
		if err != nil {
			return 0, 0, fmt.Errorf("auth_keys.InsertIgnore named exec: %w", err)
		}
		lastInsertId, err = r.LastInsertId()
		if err != nil {
			return 0, 0, fmt.Errorf("auth_keys.InsertIgnore last insert id: %w", err)
		}
		rowsAffected, err = r.RowsAffected()
		if err != nil {
			return 0, 0, fmt.Errorf("auth_keys.InsertIgnore rows affected: %w", err)
		}
		return lastInsertId, rowsAffected, nil
	}, keys...)
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
		err = fmt.Errorf("auth_keys.InsertIgnoreTx named exec: %w", err)
		return
	}

	lastInsertId, err = r.LastInsertId()
	if err != nil {
		err = fmt.Errorf("auth_keys.InsertIgnoreTx last insert id: %w", err)
		return
	}
	rowsAffected, err = r.RowsAffected()
	if err != nil {
		err = fmt.Errorf("auth_keys.InsertIgnoreTx rows affected: %w", err)
	}

	return
}

// SelectByAuthKeyId
// select id, auth_key_id, body, auth_key_type, perm_auth_key_id, temp_auth_key_id, media_temp_auth_key_id from auth_keys where auth_key_id = :auth_key_id
func (m *defaultAuthKeysModel) SelectByAuthKeyId(ctx context.Context, authKeyId int64) (rValue *AuthKeys, err error) {

	return m.FindOneByAuthKeyId(ctx, authKeyId)
}

// UpdateCustomMap
// update auth_keys set %s where auth_key_id = :auth_key_id
func (m *defaultAuthKeysModel) UpdateCustomMap(ctx context.Context, cMap map[string]interface{}, authKeyId int64) (rowsAffected int64, err error) {

	oldData, err := m.FindOneByAuthKeyId(ctx, authKeyId)

	if err != nil && !errors.Is(err, sqlx.ErrNotFound) {
		return
	}
	if oldData == nil {
		return
	}

	var keys []string

	keys = m.cacheKeys(oldData)
	names := make([]string, 0, len(cMap))
	aValues := make([]interface{}, 0, len(cMap))
	for k, v := range cMap {
		names = append(names, k+" = ?")
		aValues = append(aValues, v)
	}

	var query = fmt.Sprintf("update auth_keys set %s where auth_key_id = ?", strings.Join(names, ", "))
	aValues = append(aValues, authKeyId)

	_, rowsAffected, err = m.Exec(ctx, func(ctx context.Context, conn *sqlx.DB) (int64, int64, error) {
		rResult, err := conn.Exec(ctx, query, aValues...)
		if err != nil {
			return 0, 0, fmt.Errorf("auth_keys.UpdateCustomMap exec: %w", err)
		}
		rowsAffected, err := rResult.RowsAffected()
		if err != nil {
			return 0, 0, fmt.Errorf("auth_keys.UpdateCustomMap rows affected: %w", err)
		}
		return 0, rowsAffected, nil
	}, keys...)
	return
}

// UpdateCustomMapTx
// update auth_keys set %s where auth_key_id = :auth_key_id
func (m *defaultAuthKeysModel) UpdateCustomMapTx(tx *sqlx.Tx, cMap map[string]interface{}, authKeyId int64) (rowsAffected int64, err error) {
	names := make([]string, 0, len(cMap))
	aValues := make([]interface{}, 0, len(cMap))
	for k, v := range cMap {
		names = append(names, k+" = ?")
		aValues = append(aValues, v)
	}

	var (
		query   = fmt.Sprintf("update auth_keys set %s where auth_key_id = ?", strings.Join(names, ", "))
		rResult sql.Result
	)

	aValues = append(aValues, authKeyId)

	rResult, err = tx.Exec(query, aValues...)

	if err != nil {
		err = fmt.Errorf("auth_keys.UpdateCustomMapTx exec: %w", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		err = fmt.Errorf("auth_keys.UpdateCustomMapTx rows affected: %w", err)
	}

	return
}
