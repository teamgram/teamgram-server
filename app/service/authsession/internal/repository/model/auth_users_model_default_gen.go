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
	authUsersFieldNames          = builder.RawFieldNames(&AuthUsers{})
	authUsersRows                = strings.Join(authUsersFieldNames, ",")
	authUsersRowsExpectAutoSet   = strings.Join(stringx.Remove(authUsersFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	authUsersRowsWithPlaceHolder = strings.Join(stringx.Remove(authUsersFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"
)

type (
	authUsersModel interface {
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
	query := fmt.Sprintf("insert into `auth_users` (%s) values (?, ?, ?, ?, ?, ?, ?, ?)", authUsersRowsExpectAutoSet)

	r, err := m.db.Exec(ctx, query, data.AuthKeyId, data.UserId, data.Hash, data.DateCreated, data.DateActive, data.State, data.AndroidPushSessionId, data.Deleted)
	if err != nil {
		return nil, fmt.Errorf("auth_users.Insert2 exec: %w", err)
	}

	return r, nil
}

func (m *defaultAuthUsersModel) Delete2(ctx context.Context, id int64) error {
	query := "delete from `auth_users` where `id` = ?"

	_, err := m.db.Exec(ctx, query, id)
	if err != nil {
		return fmt.Errorf("auth_users.Delete2 exec: %w", err)
	}

	return nil
}

func (m *defaultAuthUsersModel) FindOne(ctx context.Context, id int64) (*AuthUsers, error) {
	query := fmt.Sprintf("select %s from auth_users where id = ? limit 1", authUsersRows)
	var resp AuthUsers

	err := m.db.QueryRowPartial(ctx, &resp, query, id)
	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return nil, nil
		}
		return nil, fmt.Errorf("auth_users.FindOne: %w", err)
	}

	return &resp, nil
}

func (m *defaultAuthUsersModel) FindListByIdList(ctx context.Context, id ...int64) ([]AuthUsers, error) {
	if len(id) == 0 {
		return []AuthUsers{}, nil
	}

	query := fmt.Sprintf("select %s from auth_users where id in (%s)", authUsersRows, sqlx.InInt64List(id))

	var resp []AuthUsers
	err := m.db.QueryRowsPartial(ctx, &resp, query)
	if err != nil {
		return nil, fmt.Errorf("auth_users.FindListByIdList: %w", err)
	}

	return resp, nil
}

func (m *defaultAuthUsersModel) Update2(ctx context.Context, data *AuthUsers) error {
	query := fmt.Sprintf("update `auth_users` set %s where `id` = ?", authUsersRowsWithPlaceHolder)

	_, err := m.db.Exec(ctx, query, data.AuthKeyId, data.UserId, data.Hash, data.DateCreated, data.DateActive, data.State, data.AndroidPushSessionId, data.Deleted, data.Id)
	if err != nil {
		return fmt.Errorf("auth_users.Update2 exec: %w", err)
	}

	return nil
}

func (m *defaultAuthUsersModel) FindOneByAuthKeyIdUserId(ctx context.Context, authKeyId int64, userId int64) (*AuthUsers, error) {
	query := fmt.Sprintf("select %s from auth_users where auth_key_id = ? AND user_id = ? limit 1", authUsersRows)
	var resp AuthUsers

	err := m.db.QueryRowPartial(ctx, &resp, query, authKeyId, userId)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return nil, nil
		}
		return nil, fmt.Errorf("auth_users.FindOneByAuthKeyIdUserId: %w", err)
	}

	return &resp, nil
}
