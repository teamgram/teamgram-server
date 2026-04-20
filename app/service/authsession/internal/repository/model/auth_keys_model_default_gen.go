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
	"fmt"
	"strings"

	"github.com/teamgram/marmota/pkg/stores/sqlx"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	auth_keysFieldNames          = builder.RawFieldNames(&AuthKeys{})
	auth_keysRows                = strings.Join(auth_keysFieldNames, ",")
	auth_keysRowsExpectAutoSet   = strings.Join(stringx.Remove(auth_keysFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	auth_keysRowsWithPlaceHolder = strings.Join(stringx.Remove(auth_keysFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"

	cacheTAuthKeysIdPrefix = "cache:t:auth_keys:id:"

	cacheAuthKeysIdPrefix = "cache#AuthKeys#id"

	cacheAuthKeysAuthKeyIdPrefix = "cache#AuthKeyId"
)

type (
	auth_keysModel interface {
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
	}

	AuthKeys struct {
		Id        int64  `db:"id" json:"id"`
		AuthKeyId int64  `db:"auth_key_id" json:"auth_key_id"`
		Body      string `db:"body" json:"body"`
		Deleted   bool   `db:"deleted" json:"deleted"`
	}
)

func newAuthKeysModel(db *sqlx.DB) *defaultAuthKeysModel {
	return &defaultAuthKeysModel{
		db: db,
	}
}

func (m *defaultAuthKeysModel) Insert2(ctx context.Context, data *AuthKeys) (sql.Result, error) {
	query := fmt.Sprintf("insert into `auth_keys` (%s) values (?, ?, ?)", auth_keysRowsExpectAutoSet)
	return m.db.Exec(ctx, query, data.AuthKeyId, data.Body, data.Deleted)
}

func (m *defaultAuthKeysModel) Delete2(ctx context.Context, id int64) error {
	query := "delete from `auth_keys` where `id` = ?"
	_, err := m.db.Exec(ctx, query, id)
	return err
}

func (m *defaultAuthKeysModel) FindOne(ctx context.Context, id int64) (*AuthKeys, error) {
	query := fmt.Sprintf("select %s from auth_keys where id = ? limit 1", auth_keysRows)
	var resp AuthKeys
	err := m.db.QueryRowPartial(ctx, &resp, query, id)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (m *defaultAuthKeysModel) FindListByIdList(ctx context.Context, id ...int64) ([]AuthKeys, error) {
	if len(id) == 0 {
		return []AuthKeys{}, nil
	}

	query := fmt.Sprintf("select %s from auth_keys where id in (%s)", auth_keysRows, sqlx.InInt64List(id))

	var resp []AuthKeys
	err := m.db.QueryRowsPartial(ctx, &resp, query)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (m *defaultAuthKeysModel) Update2(ctx context.Context, data *AuthKeys) error {
	query := fmt.Sprintf("update `auth_keys` set %s where `id` = ?", auth_keysRowsWithPlaceHolder)
	_, err := m.db.Exec(ctx, query, data.AuthKeyId, data.Body, data.Deleted, data.Id)
	return err
}

func (m *defaultAuthKeysModel) FindOneByAuthKeyId(ctx context.Context, authKeyId int64) (*AuthKeys, error) {
	query := fmt.Sprintf("select %s from auth_keys where auth_key_id = ? limit 1", auth_keysRows)
	var resp AuthKeys
	err := m.db.QueryRowPartial(ctx, &resp, query, authKeyId)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (m *defaultAuthKeysModel) FindListByAuthKeyIdList(ctx context.Context, authKeyId ...int64) ([]AuthKeys, error) {
	if len(authKeyId) == 0 {
		return []AuthKeys{}, nil
	}

	query := fmt.Sprintf("select %s from auth_keys where auth_key_id in (%s)", auth_keysRows, sqlx.InInt64List(authKeyId))

	var resp []AuthKeys
	err := m.db.QueryRowsPartial(ctx, &resp, query)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (m *defaultAuthKeysModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s#%v", cacheAuthKeysIdPrefix, primary)
}

func (m *defaultAuthKeysModel) queryPrimary(ctx context.Context, v interface{}, primary interface{}) error {
	query := fmt.Sprintf("select %s from auth_keys where id = ? limit 1", auth_keysRows)
	return m.db.QueryRowPartial(ctx, v, query, primary)
}
