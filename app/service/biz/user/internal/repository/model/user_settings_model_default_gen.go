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
	user_settingsFieldNames          = builder.RawFieldNames(&UserSettings{})
	user_settingsRows                = strings.Join(user_settingsFieldNames, ",")
	user_settingsRowsExpectAutoSet   = strings.Join(stringx.Remove(user_settingsFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	user_settingsRowsWithPlaceHolder = strings.Join(stringx.Remove(user_settingsFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"

	cacheTUserSettingsIdPrefix = "cache:t:user_settings:id:"

	cacheUserSettingsIdPrefix = "cache#UserSettings#id"

	cacheUserSettingsUserIdKey2Prefix = "cache#UserId#Key2"
)

type (
	user_settingsModel interface {
		Insert2(ctx context.Context, data *UserSettings) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*UserSettings, error)
		FindListByIdList(ctx context.Context, id ...int64) ([]UserSettings, error)
		Update2(ctx context.Context, data *UserSettings) error
		Delete2(ctx context.Context, id int64) error

		FindOneByUserIdKey2(ctx context.Context, userId int64, key2 string) (*UserSettings, error)
	}

	defaultUserSettingsModel struct {
		db *sqlx.DB
	}

	UserSettings struct {
		Id      int64  `db:"id" json:"id"`
		UserId  int64  `db:"user_id" json:"user_id"`
		Key2    string `db:"key2" json:"key2"`
		Value   string `db:"value" json:"value"`
		Deleted bool   `db:"deleted" json:"deleted"`
	}
)

func newUserSettingsModel(db *sqlx.DB) *defaultUserSettingsModel {
	return &defaultUserSettingsModel{
		db: db,
	}
}

func (m *defaultUserSettingsModel) Insert2(ctx context.Context, data *UserSettings) (sql.Result, error) {
	query := fmt.Sprintf("insert into `user_settings` (%s) values (?, ?, ?, ?)", user_settingsRowsExpectAutoSet)
	return m.db.Exec(ctx, query, data.UserId, data.Key2, data.Value, data.Deleted)
}

func (m *defaultUserSettingsModel) Delete2(ctx context.Context, id int64) error {
	query := "delete from `user_settings` where `id` = ?"
	_, err := m.db.Exec(ctx, query, id)
	return err
}

func (m *defaultUserSettingsModel) FindOne(ctx context.Context, id int64) (*UserSettings, error) {
	query := fmt.Sprintf("select %s from user_settings where id = ? limit 1", user_settingsRows)
	var resp UserSettings
	err := m.db.QueryRowPartial(ctx, &resp, query, id)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (m *defaultUserSettingsModel) FindListByIdList(ctx context.Context, id ...int64) ([]UserSettings, error) {
	if len(id) == 0 {
		return []UserSettings{}, nil
	}

	query := fmt.Sprintf("select %s from user_settings where id in (%s)", user_settingsRows, sqlx.InInt64List(id))

	var resp []UserSettings
	err := m.db.QueryRowsPartial(ctx, &resp, query)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (m *defaultUserSettingsModel) Update2(ctx context.Context, data *UserSettings) error {
	query := fmt.Sprintf("update `user_settings` set %s where `id` = ?", user_settingsRowsWithPlaceHolder)
	_, err := m.db.Exec(ctx, query, data.UserId, data.Key2, data.Value, data.Deleted, data.Id)
	return err
}

func (m *defaultUserSettingsModel) FindOneByUserIdKey2(ctx context.Context, userId int64, key2 string) (*UserSettings, error) {
	query := fmt.Sprintf("select %s from user_settings where user_id = ? AND key2 = ? limit 1", user_settingsRows)
	var resp UserSettings
	err := m.db.QueryRowPartial(ctx, &resp, query, userId, key2)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (m *defaultUserSettingsModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s#%v", cacheUserSettingsIdPrefix, primary)
}

func (m *defaultUserSettingsModel) queryPrimary(ctx context.Context, v interface{}, primary interface{}) error {
	query := fmt.Sprintf("select %s from user_settings where id = ? limit 1", user_settingsRows)
	return m.db.QueryRowPartial(ctx, v, query, primary)
}
