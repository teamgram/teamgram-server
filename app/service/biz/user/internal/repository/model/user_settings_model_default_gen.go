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

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	userSettingsFieldNames          = builder.RawFieldNames(&UserSettings{})
	userSettingsRows                = strings.Join(userSettingsFieldNames, ",")
	userSettingsRowsExpectAutoSet   = strings.Join(stringx.Remove(userSettingsFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	userSettingsRowsWithPlaceHolder = strings.Join(stringx.Remove(userSettingsFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"
)

type (
	userSettingsModel interface {
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
	query := fmt.Sprintf("insert into `user_settings` (%s) values (?, ?, ?, ?)", userSettingsRowsExpectAutoSet)

	r, err := m.db.Exec(ctx, query, data.UserId, data.Key2, data.Value, data.Deleted)
	if err != nil {
		return nil, fmt.Errorf("user_settings.Insert2 exec: %w", err)
	}

	return r, nil
}

func (m *defaultUserSettingsModel) Delete2(ctx context.Context, id int64) error {
	query := "delete from `user_settings` where `id` = ?"

	_, err := m.db.Exec(ctx, query, id)
	if err != nil {
		return fmt.Errorf("user_settings.Delete2 exec: %w", err)
	}

	return nil
}

func (m *defaultUserSettingsModel) FindOne(ctx context.Context, id int64) (*UserSettings, error) {
	query := fmt.Sprintf("select %s from user_settings where id = ? limit 1", userSettingsRows)
	var resp UserSettings

	err := m.db.QueryRowPartial(ctx, &resp, query, id)
	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return nil, nil
		}
		return nil, fmt.Errorf("user_settings.FindOne: %w", err)
	}

	return &resp, nil
}

func (m *defaultUserSettingsModel) FindListByIdList(ctx context.Context, id ...int64) ([]UserSettings, error) {
	if len(id) == 0 {
		return []UserSettings{}, nil
	}

	query := fmt.Sprintf("select %s from user_settings where id in (%s)", userSettingsRows, sqlx.InInt64List(id))

	var resp []UserSettings
	err := m.db.QueryRowsPartial(ctx, &resp, query)
	if err != nil {
		return nil, fmt.Errorf("user_settings.FindListByIdList: %w", err)
	}

	return resp, nil
}

func (m *defaultUserSettingsModel) Update2(ctx context.Context, data *UserSettings) error {
	query := fmt.Sprintf("update `user_settings` set %s where `id` = ?", userSettingsRowsWithPlaceHolder)

	_, err := m.db.Exec(ctx, query, data.UserId, data.Key2, data.Value, data.Deleted, data.Id)
	if err != nil {
		return fmt.Errorf("user_settings.Update2 exec: %w", err)
	}

	return nil
}

func (m *defaultUserSettingsModel) FindOneByUserIdKey2(ctx context.Context, userId int64, key2 string) (*UserSettings, error) {
	query := fmt.Sprintf("select %s from user_settings where user_id = ? AND key2 = ? limit 1", userSettingsRows)
	var resp UserSettings

	err := m.db.QueryRowPartial(ctx, &resp, query, userId, key2)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return nil, nil
		}
		return nil, fmt.Errorf("user_settings.FindOneByUserIdKey2: %w", err)
	}

	return &resp, nil
}
