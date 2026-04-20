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
	auth_key_infosFieldNames          = builder.RawFieldNames(&AuthKeyInfos{})
	auth_key_infosRows                = strings.Join(auth_key_infosFieldNames, ",")
	auth_key_infosRowsExpectAutoSet   = strings.Join(stringx.Remove(auth_key_infosFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	auth_key_infosRowsWithPlaceHolder = strings.Join(stringx.Remove(auth_key_infosFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"

	cacheTAuthKeyInfosIdPrefix = "cache:t:auth_key_infos:id:"

	cacheAuthKeyInfosIdPrefix = "cache#AuthKeyInfos#id"

	cacheAuthKeyInfosAuthKeyIdPrefix = "cache#AuthKeyId"
)

type (
	auth_key_infosModel interface {
		Insert2(ctx context.Context, data *AuthKeyInfos) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*AuthKeyInfos, error)
		FindListByIdList(ctx context.Context, id ...int64) ([]AuthKeyInfos, error)
		Update2(ctx context.Context, data *AuthKeyInfos) error
		Delete2(ctx context.Context, id int64) error

		FindOneByAuthKeyId(ctx context.Context, authKeyId int64) (*AuthKeyInfos, error)
		FindListByAuthKeyIdList(ctx context.Context, authKeyId ...int64) ([]AuthKeyInfos, error)
	}

	defaultAuthKeyInfosModel struct {
		db *sqlx.DB
	}

	AuthKeyInfos struct {
		Id                 int64 `db:"id" json:"id"`
		AuthKeyId          int64 `db:"auth_key_id" json:"auth_key_id"`
		AuthKeyType        int32 `db:"auth_key_type" json:"auth_key_type"`
		PermAuthKeyId      int64 `db:"perm_auth_key_id" json:"perm_auth_key_id"`
		TempAuthKeyId      int64 `db:"temp_auth_key_id" json:"temp_auth_key_id"`
		MediaTempAuthKeyId int64 `db:"media_temp_auth_key_id" json:"media_temp_auth_key_id"`
		Deleted            bool  `db:"deleted" json:"deleted"`
	}
)

func newAuthKeyInfosModel(db *sqlx.DB) *defaultAuthKeyInfosModel {
	return &defaultAuthKeyInfosModel{
		db: db,
	}
}

func (m *defaultAuthKeyInfosModel) Insert2(ctx context.Context, data *AuthKeyInfos) (sql.Result, error) {
	query := fmt.Sprintf("insert into `auth_key_infos` (%s) values (?, ?, ?, ?, ?, ?)", auth_key_infosRowsExpectAutoSet)
	return m.db.Exec(ctx, query, data.AuthKeyId, data.AuthKeyType, data.PermAuthKeyId, data.TempAuthKeyId, data.MediaTempAuthKeyId, data.Deleted)
}

func (m *defaultAuthKeyInfosModel) Delete2(ctx context.Context, id int64) error {
	query := "delete from `auth_key_infos` where `id` = ?"
	_, err := m.db.Exec(ctx, query, id)
	return err
}

func (m *defaultAuthKeyInfosModel) FindOne(ctx context.Context, id int64) (*AuthKeyInfos, error) {
	query := fmt.Sprintf("select %s from auth_key_infos where id = ? limit 1", auth_key_infosRows)
	var resp AuthKeyInfos
	err := m.db.QueryRowPartial(ctx, &resp, query, id)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (m *defaultAuthKeyInfosModel) FindListByIdList(ctx context.Context, id ...int64) ([]AuthKeyInfos, error) {
	if len(id) == 0 {
		return []AuthKeyInfos{}, nil
	}

	query := fmt.Sprintf("select %s from auth_key_infos where id in (%s)", auth_key_infosRows, sqlx.InInt64List(id))

	var resp []AuthKeyInfos
	err := m.db.QueryRowsPartial(ctx, &resp, query)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (m *defaultAuthKeyInfosModel) Update2(ctx context.Context, data *AuthKeyInfos) error {
	query := fmt.Sprintf("update `auth_key_infos` set %s where `id` = ?", auth_key_infosRowsWithPlaceHolder)
	_, err := m.db.Exec(ctx, query, data.AuthKeyId, data.AuthKeyType, data.PermAuthKeyId, data.TempAuthKeyId, data.MediaTempAuthKeyId, data.Deleted, data.Id)
	return err
}

func (m *defaultAuthKeyInfosModel) FindOneByAuthKeyId(ctx context.Context, authKeyId int64) (*AuthKeyInfos, error) {
	query := fmt.Sprintf("select %s from auth_key_infos where auth_key_id = ? limit 1", auth_key_infosRows)
	var resp AuthKeyInfos
	err := m.db.QueryRowPartial(ctx, &resp, query, authKeyId)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (m *defaultAuthKeyInfosModel) FindListByAuthKeyIdList(ctx context.Context, authKeyId ...int64) ([]AuthKeyInfos, error) {
	if len(authKeyId) == 0 {
		return []AuthKeyInfos{}, nil
	}

	query := fmt.Sprintf("select %s from auth_key_infos where auth_key_id in (%s)", auth_key_infosRows, sqlx.InInt64List(authKeyId))

	var resp []AuthKeyInfos
	err := m.db.QueryRowsPartial(ctx, &resp, query)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (m *defaultAuthKeyInfosModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s#%v", cacheAuthKeyInfosIdPrefix, primary)
}

func (m *defaultAuthKeyInfosModel) queryPrimary(ctx context.Context, v interface{}, primary interface{}) error {
	query := fmt.Sprintf("select %s from auth_key_infos where id = ? limit 1", auth_key_infosRows)
	return m.db.QueryRowPartial(ctx, v, query, primary)
}
