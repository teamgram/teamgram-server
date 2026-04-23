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
	"fmt"
	"strings"

	"github.com/teamgram/marmota/pkg/stores/sqlx"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	authKeysFieldNames          = builder.RawFieldNames(&AuthKeys{})
	authKeysRows                = strings.Join(authKeysFieldNames, ",")
	authKeysRowsExpectAutoSet   = strings.Join(stringx.Remove(authKeysFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	authKeysRowsWithPlaceHolder = strings.Join(stringx.Remove(authKeysFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"
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

func newAuthKeysModel(db *sqlx.DB) *defaultAuthKeysModel {
	return &defaultAuthKeysModel{
		db: db,
	}
}

func (m *defaultAuthKeysModel) Insert2(ctx context.Context, data *AuthKeys) (sql.Result, error) {
	query := fmt.Sprintf("insert into `auth_keys` (%s) values (?, ?, ?, ?, ?, ?, ?)", authKeysRowsExpectAutoSet)
	return m.db.Exec(ctx, query, data.AuthKeyId, data.Body, data.AuthKeyType, data.PermAuthKeyId, data.TempAuthKeyId, data.MediaTempAuthKeyId, data.Deleted)
}

func (m *defaultAuthKeysModel) Delete2(ctx context.Context, id int64) error {
	query := "delete from `auth_keys` where `id` = ?"
	_, err := m.db.Exec(ctx, query, id)
	return err
}

func (m *defaultAuthKeysModel) FindOne(ctx context.Context, id int64) (*AuthKeys, error) {
	query := fmt.Sprintf("select %s from auth_keys where id = ? limit 1", authKeysRows)
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

	query := fmt.Sprintf("select %s from auth_keys where id in (%s)", authKeysRows, sqlx.InInt64List(id))

	var resp []AuthKeys
	err := m.db.QueryRowsPartial(ctx, &resp, query)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (m *defaultAuthKeysModel) Update2(ctx context.Context, data *AuthKeys) error {
	query := fmt.Sprintf("update `auth_keys` set %s where `id` = ?", authKeysRowsWithPlaceHolder)
	_, err := m.db.Exec(ctx, query, data.AuthKeyId, data.Body, data.AuthKeyType, data.PermAuthKeyId, data.TempAuthKeyId, data.MediaTempAuthKeyId, data.Deleted, data.Id)
	return err
}

func (m *defaultAuthKeysModel) FindOneByAuthKeyId(ctx context.Context, authKeyId int64) (*AuthKeys, error) {
	query := fmt.Sprintf("select %s from auth_keys where auth_key_id = ? limit 1", authKeysRows)
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

	query := fmt.Sprintf("select %s from auth_keys where auth_key_id in (%s)", authKeysRows, sqlx.InInt64List(authKeyId))

	var resp []AuthKeys
	err := m.db.QueryRowsPartial(ctx, &resp, query)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
