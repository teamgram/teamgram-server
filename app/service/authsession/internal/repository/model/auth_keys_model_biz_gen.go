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

		UpdatePermBinding(ctx context.Context, permAuthKeyId int64, authKeyId int64) (rowsAffected int64, err error)
		UpdatePermBindingTx(tx *sqlx.Tx, permAuthKeyId int64, authKeyId int64) (rowsAffected int64, err error)

		UpdateTempBinding(ctx context.Context, tempAuthKeyId int64, authKeyId int64) (rowsAffected int64, err error)
		UpdateTempBindingTx(tx *sqlx.Tx, tempAuthKeyId int64, authKeyId int64) (rowsAffected int64, err error)

		UpdateMediaTempBinding(ctx context.Context, mediaTempAuthKeyId int64, authKeyId int64) (rowsAffected int64, err error)
		UpdateMediaTempBindingTx(tx *sqlx.Tx, mediaTempAuthKeyId int64, authKeyId int64) (rowsAffected int64, err error)
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

// UpdatePermBinding
// update auth_keys set perm_auth_key_id = :perm_auth_key_id where auth_key_id = :auth_key_id
func (m *defaultAuthKeysModel) UpdatePermBinding(ctx context.Context, permAuthKeyId int64, authKeyId int64) (rowsAffected int64, err error) {
	var query = "update auth_keys set perm_auth_key_id = ? where auth_key_id = ?"
	oldData, err := m.FindOneByAuthKeyId(ctx, authKeyId)

	if err != nil {
		if errors.Is(err, ErrNotFound) {
			err = &NotFoundError{
				Resource: "auth_keys",
				Key:      fmt.Sprintf("auth_key_id=%v", authKeyId),
				Cause:    err,
			}

			return
		}
		err = fmt.Errorf("auth_keys.UpdatePermBinding find one: %w", err)
		return
	}

	var keys []string

	keys = m.cacheKeys(oldData)
	_, rowsAffected, err = m.Exec(ctx, func(ctx context.Context, conn *sqlx.DB) (int64, int64, error) {
		rResult, err := conn.Exec(ctx, query, permAuthKeyId, authKeyId)
		if err != nil {
			return 0, 0, fmt.Errorf("auth_keys.UpdatePermBinding exec: %w", err)
		}
		rowsAffected, err := rResult.RowsAffected()
		if err != nil {
			return 0, 0, fmt.Errorf("auth_keys.UpdatePermBinding rows affected: %w", err)
		}
		return 0, rowsAffected, nil
	}, keys...)

	if err == nil && rowsAffected == 0 {
		err = &NotFoundError{
			Resource: "auth_keys",
			Key:      fmt.Sprintf("auth_key_id=%v", authKeyId),
			Cause:    ErrNotFound,
		}
	}

	return
}

// UpdatePermBindingTx
// update auth_keys set perm_auth_key_id = :perm_auth_key_id where auth_key_id = :auth_key_id
func (m *defaultAuthKeysModel) UpdatePermBindingTx(tx *sqlx.Tx, permAuthKeyId int64, authKeyId int64) (rowsAffected int64, err error) {
	var (
		query   = "update auth_keys set perm_auth_key_id = ? where auth_key_id = ?"
		rResult sql.Result
	)
	rResult, err = tx.Exec(query, permAuthKeyId, authKeyId)

	if err != nil {
		err = fmt.Errorf("auth_keys.UpdatePermBindingTx exec: %w", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		err = fmt.Errorf("auth_keys.UpdatePermBindingTx rows affected: %w", err)
		return
	}

	if rowsAffected == 0 {
		err = &NotFoundError{
			Resource: "auth_keys",
			Key:      fmt.Sprintf("auth_key_id=%v", authKeyId),
			Cause:    ErrNotFound,
		}
	}

	return
}

// UpdateTempBinding
// update auth_keys set temp_auth_key_id = :temp_auth_key_id where auth_key_id = :auth_key_id
func (m *defaultAuthKeysModel) UpdateTempBinding(ctx context.Context, tempAuthKeyId int64, authKeyId int64) (rowsAffected int64, err error) {
	var query = "update auth_keys set temp_auth_key_id = ? where auth_key_id = ?"
	oldData, err := m.FindOneByAuthKeyId(ctx, authKeyId)

	if err != nil {
		if errors.Is(err, ErrNotFound) {
			err = &NotFoundError{
				Resource: "auth_keys",
				Key:      fmt.Sprintf("auth_key_id=%v", authKeyId),
				Cause:    err,
			}

			return
		}
		err = fmt.Errorf("auth_keys.UpdateTempBinding find one: %w", err)
		return
	}

	var keys []string

	keys = m.cacheKeys(oldData)
	_, rowsAffected, err = m.Exec(ctx, func(ctx context.Context, conn *sqlx.DB) (int64, int64, error) {
		rResult, err := conn.Exec(ctx, query, tempAuthKeyId, authKeyId)
		if err != nil {
			return 0, 0, fmt.Errorf("auth_keys.UpdateTempBinding exec: %w", err)
		}
		rowsAffected, err := rResult.RowsAffected()
		if err != nil {
			return 0, 0, fmt.Errorf("auth_keys.UpdateTempBinding rows affected: %w", err)
		}
		return 0, rowsAffected, nil
	}, keys...)

	if err == nil && rowsAffected == 0 {
		err = &NotFoundError{
			Resource: "auth_keys",
			Key:      fmt.Sprintf("auth_key_id=%v", authKeyId),
			Cause:    ErrNotFound,
		}
	}

	return
}

// UpdateTempBindingTx
// update auth_keys set temp_auth_key_id = :temp_auth_key_id where auth_key_id = :auth_key_id
func (m *defaultAuthKeysModel) UpdateTempBindingTx(tx *sqlx.Tx, tempAuthKeyId int64, authKeyId int64) (rowsAffected int64, err error) {
	var (
		query   = "update auth_keys set temp_auth_key_id = ? where auth_key_id = ?"
		rResult sql.Result
	)
	rResult, err = tx.Exec(query, tempAuthKeyId, authKeyId)

	if err != nil {
		err = fmt.Errorf("auth_keys.UpdateTempBindingTx exec: %w", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		err = fmt.Errorf("auth_keys.UpdateTempBindingTx rows affected: %w", err)
		return
	}

	if rowsAffected == 0 {
		err = &NotFoundError{
			Resource: "auth_keys",
			Key:      fmt.Sprintf("auth_key_id=%v", authKeyId),
			Cause:    ErrNotFound,
		}
	}

	return
}

// UpdateMediaTempBinding
// update auth_keys set media_temp_auth_key_id = :media_temp_auth_key_id where auth_key_id = :auth_key_id
func (m *defaultAuthKeysModel) UpdateMediaTempBinding(ctx context.Context, mediaTempAuthKeyId int64, authKeyId int64) (rowsAffected int64, err error) {
	var query = "update auth_keys set media_temp_auth_key_id = ? where auth_key_id = ?"
	oldData, err := m.FindOneByAuthKeyId(ctx, authKeyId)

	if err != nil {
		if errors.Is(err, ErrNotFound) {
			err = &NotFoundError{
				Resource: "auth_keys",
				Key:      fmt.Sprintf("auth_key_id=%v", authKeyId),
				Cause:    err,
			}

			return
		}
		err = fmt.Errorf("auth_keys.UpdateMediaTempBinding find one: %w", err)
		return
	}

	var keys []string

	keys = m.cacheKeys(oldData)
	_, rowsAffected, err = m.Exec(ctx, func(ctx context.Context, conn *sqlx.DB) (int64, int64, error) {
		rResult, err := conn.Exec(ctx, query, mediaTempAuthKeyId, authKeyId)
		if err != nil {
			return 0, 0, fmt.Errorf("auth_keys.UpdateMediaTempBinding exec: %w", err)
		}
		rowsAffected, err := rResult.RowsAffected()
		if err != nil {
			return 0, 0, fmt.Errorf("auth_keys.UpdateMediaTempBinding rows affected: %w", err)
		}
		return 0, rowsAffected, nil
	}, keys...)

	if err == nil && rowsAffected == 0 {
		err = &NotFoundError{
			Resource: "auth_keys",
			Key:      fmt.Sprintf("auth_key_id=%v", authKeyId),
			Cause:    ErrNotFound,
		}
	}

	return
}

// UpdateMediaTempBindingTx
// update auth_keys set media_temp_auth_key_id = :media_temp_auth_key_id where auth_key_id = :auth_key_id
func (m *defaultAuthKeysModel) UpdateMediaTempBindingTx(tx *sqlx.Tx, mediaTempAuthKeyId int64, authKeyId int64) (rowsAffected int64, err error) {
	var (
		query   = "update auth_keys set media_temp_auth_key_id = ? where auth_key_id = ?"
		rResult sql.Result
	)
	rResult, err = tx.Exec(query, mediaTempAuthKeyId, authKeyId)

	if err != nil {
		err = fmt.Errorf("auth_keys.UpdateMediaTempBindingTx exec: %w", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		err = fmt.Errorf("auth_keys.UpdateMediaTempBindingTx rows affected: %w", err)
		return
	}

	if rowsAffected == 0 {
		err = &NotFoundError{
			Resource: "auth_keys",
			Key:      fmt.Sprintf("auth_key_id=%v", authKeyId),
			Cause:    ErrNotFound,
		}
	}

	return
}
