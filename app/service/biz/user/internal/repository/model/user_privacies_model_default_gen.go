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
	userPrivaciesFieldNames          = builder.RawFieldNames(&UserPrivacies{})
	userPrivaciesRows                = strings.Join(userPrivaciesFieldNames, ",")
	userPrivaciesRowsExpectAutoSet   = strings.Join(stringx.Remove(userPrivaciesFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	userPrivaciesRowsWithPlaceHolder = strings.Join(stringx.Remove(userPrivaciesFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"
)

type (
	userPrivaciesModel interface {
		Insert2(ctx context.Context, data *UserPrivacies) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*UserPrivacies, error)
		FindListByIdList(ctx context.Context, id ...int64) ([]UserPrivacies, error)
		Update2(ctx context.Context, data *UserPrivacies) error
		Delete2(ctx context.Context, id int64) error

		FindOneByUserIdKeyType(ctx context.Context, userId int64, keyType int32) (*UserPrivacies, error)
	}

	defaultUserPrivaciesModel struct {
		db *sqlx.DB
	}

	UserPrivacies struct {
		Id      int64  `db:"id" json:"id"`
		UserId  int64  `db:"user_id" json:"user_id"`
		KeyType int32  `db:"key_type" json:"key_type"`
		Rules   string `db:"rules" json:"rules"`
	}
)

func newUserPrivaciesModel(db *sqlx.DB) *defaultUserPrivaciesModel {
	return &defaultUserPrivaciesModel{
		db: db,
	}
}

func (m *defaultUserPrivaciesModel) Insert2(ctx context.Context, data *UserPrivacies) (sql.Result, error) {
	query := fmt.Sprintf("insert into `user_privacies` (%s) values (?, ?, ?)", userPrivaciesRowsExpectAutoSet)

	return m.db.Exec(ctx, query, data.UserId, data.KeyType, data.Rules)

}

func (m *defaultUserPrivaciesModel) Delete2(ctx context.Context, id int64) error {
	query := "delete from `user_privacies` where `id` = ?"

	_, err := m.db.Exec(ctx, query, id)
	return err
}

func (m *defaultUserPrivaciesModel) FindOne(ctx context.Context, id int64) (*UserPrivacies, error) {
	query := fmt.Sprintf("select %s from user_privacies where id = ? limit 1", userPrivaciesRows)
	var resp UserPrivacies

	err := m.db.QueryRowPartial(ctx, &resp, query, id)

	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (m *defaultUserPrivaciesModel) FindListByIdList(ctx context.Context, id ...int64) ([]UserPrivacies, error) {
	if len(id) == 0 {
		return []UserPrivacies{}, nil
	}

	query := fmt.Sprintf("select %s from user_privacies where id in (%s)", userPrivaciesRows, sqlx.InInt64List(id))

	var resp []UserPrivacies
	err := m.db.QueryRowsPartial(ctx, &resp, query)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (m *defaultUserPrivaciesModel) Update2(ctx context.Context, data *UserPrivacies) error {
	query := fmt.Sprintf("update `user_privacies` set %s where `id` = ?", userPrivaciesRowsWithPlaceHolder)

	_, err := m.db.Exec(ctx, query, data.UserId, data.KeyType, data.Rules, data.Id)
	return err
}

func (m *defaultUserPrivaciesModel) FindOneByUserIdKeyType(ctx context.Context, userId int64, keyType int32) (*UserPrivacies, error) {
	query := fmt.Sprintf("select %s from user_privacies where user_id = ? AND key_type = ? limit 1", userPrivaciesRows)
	var resp UserPrivacies

	err := m.db.QueryRowPartial(ctx, &resp, query, userId, keyType)

	if err != nil {
		return nil, err
	}
	return &resp, nil
}
