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
	auth_usersFieldNames          = builder.RawFieldNames(&AuthUsers{})
	auth_usersRows                = strings.Join(auth_usersFieldNames, ",")
	auth_usersRowsExpectAutoSet   = strings.Join(stringx.Remove(auth_usersFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	auth_usersRowsWithPlaceHolder = strings.Join(stringx.Remove(auth_usersFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"

	cacheTAuthUsersIdPrefix = "cache:t:auth_users:id:"

	cacheAuthUsersIdPrefix = "cache#AuthUsers#id"

	cacheAuthUsersAuthKeyIdUserIdPrefix = "cache#AuthKeyId#UserId"
)

type (
	auth_usersModel interface {
		Insert2(ctx context.Context, data *AuthUsers) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*AuthUsers, error)
		FindListByIdList(ctx context.Context, id ...int64) ([]AuthUsers, error)
		Update2(ctx context.Context, data *AuthUsers) error
		Delete2(ctx context.Context, id int64) error

		FindOneByAuthKeyIdUserId(ctx context.Context, authKeyId int64, userId int64) (*AuthUsers, error)
	}

	defaultAuthUsersModel struct {
		db *sqlx.DB
	}

	AuthUsers struct {
		Id                   int64 `db:"id" json:"id"`
		AuthKeyId            int64 `db:"auth_key_id" json:"auth_key_id"`
		UserId               int64 `db:"user_id" json:"user_id"`
		Hash                 int64 `db:"hash" json:"hash"`
		DateCreated          int64 `db:"date_created" json:"date_created"`
		DateActive           int64 `db:"date_active" json:"date_active"`
		State                int32 `db:"state" json:"state"`
		AndroidPushSessionId int64 `db:"android_push_session_id" json:"android_push_session_id"`
		Deleted              bool  `db:"deleted" json:"deleted"`
	}
)

func newAuthUsersModel(db *sqlx.DB) *defaultAuthUsersModel {
	return &defaultAuthUsersModel{
		db: db,
	}
}

func (m *defaultAuthUsersModel) Insert2(ctx context.Context, data *AuthUsers) (sql.Result, error) {
	query := fmt.Sprintf("insert into `auth_users` (%s) values (?, ?, ?, ?, ?, ?, ?, ?)", auth_usersRowsExpectAutoSet)
	return m.db.Exec(ctx, query, data.AuthKeyId, data.UserId, data.Hash, data.DateCreated, data.DateActive, data.State, data.AndroidPushSessionId, data.Deleted)
}

func (m *defaultAuthUsersModel) Delete2(ctx context.Context, id int64) error {
	query := "delete from `auth_users` where `id` = ?"
	_, err := m.db.Exec(ctx, query, id)
	return err
}

func (m *defaultAuthUsersModel) FindOne(ctx context.Context, id int64) (*AuthUsers, error) {
	query := fmt.Sprintf("select %s from auth_users where id = ? limit 1", auth_usersRows)
	var resp AuthUsers
	err := m.db.QueryRowPartial(ctx, &resp, query, id)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (m *defaultAuthUsersModel) FindListByIdList(ctx context.Context, id ...int64) ([]AuthUsers, error) {
	if len(id) == 0 {
		return []AuthUsers{}, nil
	}

	query := fmt.Sprintf("select %s from auth_users where id in (%s)", auth_usersRows, sqlx.InInt64List(id))

	var resp []AuthUsers
	err := m.db.QueryRowsPartial(ctx, &resp, query)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (m *defaultAuthUsersModel) Update2(ctx context.Context, data *AuthUsers) error {
	query := fmt.Sprintf("update `auth_users` set %s where `id` = ?", auth_usersRowsWithPlaceHolder)
	_, err := m.db.Exec(ctx, query, data.AuthKeyId, data.UserId, data.Hash, data.DateCreated, data.DateActive, data.State, data.AndroidPushSessionId, data.Deleted, data.Id)
	return err
}

func (m *defaultAuthUsersModel) FindOneByAuthKeyIdUserId(ctx context.Context, authKeyId int64, userId int64) (*AuthUsers, error) {
	query := fmt.Sprintf("select %s from auth_users where auth_key_id = ? AND user_id = ? limit 1", auth_usersRows)
	var resp AuthUsers
	err := m.db.QueryRowPartial(ctx, &resp, query, authKeyId, userId)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (m *defaultAuthUsersModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s#%v", cacheAuthUsersIdPrefix, primary)
}

func (m *defaultAuthUsersModel) queryPrimary(ctx context.Context, v interface{}, primary interface{}) error {
	query := fmt.Sprintf("select %s from auth_users where id = ? limit 1", auth_usersRows)
	return m.db.QueryRowPartial(ctx, v, query, primary)
}
