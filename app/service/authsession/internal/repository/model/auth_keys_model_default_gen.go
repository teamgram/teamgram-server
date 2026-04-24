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

	"github.com/teamgram/marmota/pkg/stores/cache"
	"github.com/teamgram/marmota/pkg/stores/sqlc"
	"github.com/teamgram/marmota/pkg/stores/sqlx"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	authKeysFieldNames          = builder.RawFieldNames(&AuthKeys{})
	authKeysRows                = strings.Join(authKeysFieldNames, ",")
	authKeysRowsExpectAutoSet   = strings.Join(stringx.Remove(authKeysFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	authKeysRowsWithPlaceHolder = strings.Join(stringx.Remove(authKeysFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"

	cacheTAuthKeysIdPrefix = "cache:t:auth_keys:id:"

	cacheAuthKeysIdPrefix = "cache#AuthKeys#id"

	cacheAuthKeysAuthKeyIdPrefix = "cache#AuthKeyId"
)

type (
	authKeysModel interface {
		Insert2(ctx context.Context, data *AuthKeys) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*AuthKeys, error)
		FindListByIdList(ctx context.Context, id ...int64) ([]AuthKeys, error)
		Update2(ctx context.Context, data *AuthKeys) error
		Delete2(ctx context.Context, id int64) error

		FindOneByAuthKeyId(ctx context.Context, authKeyId int64) (*AuthKeys, error)
		FindListByAuthKeyIdList(ctx context.Context, authKeyId ...int64) ([]AuthKeys, error)
	}

	defaultAuthKeysModel struct {
		db *sqlx.DB
		sqlc.CachedConn
	}

	cachedExecResult struct {
		lastInsertId int64
		rowsAffected int64
	}

	AuthKeys struct {
		Id                 int64  `db:"id" json:"id"`
		AuthKeyId          int64  `db:"auth_key_id" json:"auth_key_id"`
		Body               string `db:"body" json:"body"`
		AuthKeyType        int32  `db:"auth_key_type" json:"auth_key_type"`
		PermAuthKeyId      int64  `db:"perm_auth_key_id" json:"perm_auth_key_id"`
		TempAuthKeyId      int64  `db:"temp_auth_key_id" json:"temp_auth_key_id"`
		MediaTempAuthKeyId int64  `db:"media_temp_auth_key_id" json:"media_temp_auth_key_id"`
		Deleted            bool   `db:"deleted" json:"deleted"`
	}
)

func (r cachedExecResult) LastInsertId() (int64, error) {
	return r.lastInsertId, nil
}

func (r cachedExecResult) RowsAffected() (int64, error) {
	return r.rowsAffected, nil
}

func newAuthKeysModel(db *sqlx.DB, c cache.CacheConf) *defaultAuthKeysModel {
	return &defaultAuthKeysModel{
		db:         db,
		CachedConn: sqlc.NewConn(db, c),
	}
}

func (m *defaultAuthKeysModel) Insert2(ctx context.Context, data *AuthKeys) (sql.Result, error) {
	query := fmt.Sprintf("insert into `auth_keys` (%s) values (?, ?, ?, ?, ?, ?, ?)", authKeysRowsExpectAutoSet)

	keys := m.uniqueCacheKeys(data)
	lastInsertId, rowsAffected, err := m.Exec(ctx, func(ctx context.Context, conn *sqlx.DB) (int64, int64, error) {
		r, err := conn.Exec(ctx, query, data.AuthKeyId, data.Body, data.AuthKeyType, data.PermAuthKeyId, data.TempAuthKeyId, data.MediaTempAuthKeyId, data.Deleted)
		if err != nil {
			return 0, 0, fmt.Errorf("auth_keys.Insert2 exec: %w", err)
		}
		lastInsertId, err := r.LastInsertId()
		if err != nil {
			return 0, 0, fmt.Errorf("auth_keys.Insert2 last insert id: %w", err)
		}
		rowsAffected, err := r.RowsAffected()
		if err != nil {
			return 0, 0, fmt.Errorf("auth_keys.Insert2 rows affected: %w", err)
		}
		return lastInsertId, rowsAffected, nil
	}, keys...)
	if err != nil {
		return nil, err
	}

	return cachedExecResult{lastInsertId: lastInsertId, rowsAffected: rowsAffected}, nil
}

func (m *defaultAuthKeysModel) Delete2(ctx context.Context, id int64) error {
	query := "delete from `auth_keys` where `id` = ?"

	oldData, err := m.FindOne(ctx, id)
	if err != nil && !errors.Is(err, sqlx.ErrNotFound) {
		return fmt.Errorf("auth_keys.Delete2 find one: %w", err)
	}
	if oldData == nil {
		return nil
	}

	keys := []string{m.formatPrimary(id)}
	keys = append(keys, m.cacheKeys(oldData)...)
	_, _, err = m.Exec(ctx, func(ctx context.Context, conn *sqlx.DB) (int64, int64, error) {
		r, err := conn.Exec(ctx, query, id)
		if err != nil {
			return 0, 0, fmt.Errorf("auth_keys.Delete2 exec: %w", err)
		}
		rowsAffected, err := r.RowsAffected()
		if err != nil {
			return 0, 0, fmt.Errorf("auth_keys.Delete2 rows affected: %w", err)
		}
		return 0, rowsAffected, nil
	}, keys...)

	return err
}

func (m *defaultAuthKeysModel) FindOne(ctx context.Context, id int64) (*AuthKeys, error) {
	query := fmt.Sprintf("select %s from auth_keys where id = ? limit 1", authKeysRows)
	var resp AuthKeys

	cacheKey := m.formatPrimary(id)
	err := m.QueryRow(ctx, &resp, cacheKey, func(ctx context.Context, conn *sqlx.DB, v interface{}) error {
		return conn.QueryRowPartial(ctx, v, query, id)
	})
	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return nil, nil
		}
		return nil, fmt.Errorf("auth_keys.FindOne: %w", err)
	}

	return &resp, nil
}

func (m *defaultAuthKeysModel) FindListByIdList(ctx context.Context, id ...int64) ([]AuthKeys, error) {
	if len(id) == 0 {
		return []AuthKeys{}, nil
	}

	query := fmt.Sprintf("select %s from auth_keys where id in (%s)", authKeysRows, sqlx.InInt64List(id))

	var resp []AuthKeys
	err := m.db.QueryRowsPartial(ctx, &resp, query)
	if err != nil {
		return nil, fmt.Errorf("auth_keys.FindListByIdList: %w", err)
	}

	return resp, nil
}

func (m *defaultAuthKeysModel) Update2(ctx context.Context, data *AuthKeys) error {
	query := fmt.Sprintf("update `auth_keys` set %s where `id` = ?", authKeysRowsWithPlaceHolder)

	oldData, err := m.FindOne(ctx, data.Id)
	if err != nil && !errors.Is(err, sqlx.ErrNotFound) {
		return fmt.Errorf("auth_keys.Update2 find one: %w", err)
	}
	if oldData == nil {
		return nil
	}

	keys := m.cacheKeys(data)
	keys = append(keys, m.cacheKeys(oldData)...)
	_, _, err = m.Exec(ctx, func(ctx context.Context, conn *sqlx.DB) (int64, int64, error) {
		r, err := conn.Exec(ctx, query, data.AuthKeyId, data.Body, data.AuthKeyType, data.PermAuthKeyId, data.TempAuthKeyId, data.MediaTempAuthKeyId, data.Deleted, data.Id)
		if err != nil {
			return 0, 0, fmt.Errorf("auth_keys.Update2 exec: %w", err)
		}
		rowsAffected, err := r.RowsAffected()
		if err != nil {
			return 0, 0, fmt.Errorf("auth_keys.Update2 rows affected: %w", err)
		}
		return 0, rowsAffected, nil
	}, keys...)

	return err
}

func (m *defaultAuthKeysModel) FindOneByAuthKeyId(ctx context.Context, authKeyId int64) (*AuthKeys, error) {
	query := fmt.Sprintf("select %s from auth_keys where auth_key_id = ? limit 1", authKeysRows)
	var resp AuthKeys

	cacheAuthKeysAuthKeyIdKey := fmt.Sprintf("%s#%v", cacheAuthKeysAuthKeyIdPrefix, authKeyId)
	err := m.QueryRowIndex(ctx, &resp, cacheAuthKeysAuthKeyIdKey, m.formatPrimary, func(ctx context.Context, conn *sqlx.DB, v interface{}) (interface{}, error) {
		if err := conn.QueryRowPartial(ctx, &resp, query, authKeyId); err != nil {
			return nil, err
		}
		return resp.Id, nil
	}, m.queryPrimary)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return nil, nil
		}
		return nil, fmt.Errorf("auth_keys.FindOneByAuthKeyId: %w", err)
	}

	return &resp, nil
}

func (m *defaultAuthKeysModel) FindListByAuthKeyIdList(ctx context.Context, authKeyId ...int64) ([]AuthKeys, error) {
	if len(authKeyId) == 0 {
		return []AuthKeys{}, nil
	}

	query := fmt.Sprintf("select %s from auth_keys where auth_key_id in (%s)", authKeysRows, sqlx.InInt64List(authKeyId))

	var resp []AuthKeys
	err := m.db.QueryRowsPartial(ctx, &resp, query)
	if err != nil {
		return nil, fmt.Errorf("auth_keys.FindListByAuthKeyIdList: %w", err)
	}

	return resp, nil
}

func (m *defaultAuthKeysModel) cacheKeys(data *AuthKeys) []string {
	if data == nil {
		return nil
	}
	keys := []string{m.formatPrimary(data.Id)}
	keys = append(keys, m.uniqueCacheKeys(data)...)

	return keys
}

func (m *defaultAuthKeysModel) uniqueCacheKeys(data *AuthKeys) []string {
	if data == nil {
		return nil
	}
	var keys []string

	keys = append(keys, fmt.Sprintf("%s#%v", cacheAuthKeysAuthKeyIdPrefix, data.AuthKeyId))

	return keys
}

func (m *defaultAuthKeysModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s#%v", cacheAuthKeysIdPrefix, primary)
}

func (m *defaultAuthKeysModel) queryPrimary(ctx context.Context, conn *sqlx.DB, v interface{}, primary interface{}) error {
	query := fmt.Sprintf("select %s from auth_keys where id = ? limit 1", authKeysRows)

	return conn.QueryRowPartial(ctx, v, query, primary)
}
