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
	predefinedUsersFieldNames          = builder.RawFieldNames(&PredefinedUsers{})
	predefinedUsersRows                = strings.Join(predefinedUsersFieldNames, ",")
	predefinedUsersRowsExpectAutoSet   = strings.Join(stringx.Remove(predefinedUsersFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	predefinedUsersRowsWithPlaceHolder = strings.Join(stringx.Remove(predefinedUsersFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"
)

type (
	predefinedUsersModel interface {
		Insert2(ctx context.Context, data *PredefinedUsers) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*PredefinedUsers, error)
		FindListByIdList(ctx context.Context, id ...int64) ([]PredefinedUsers, error)
		Update2(ctx context.Context, data *PredefinedUsers) error
		Delete2(ctx context.Context, id int64) error
	}

	defaultPredefinedUsersModel struct {
		db *sqlx.DB
	}

	PredefinedUsers struct {
		Id               int64  `db:"id" json:"id"`
		Phone            string `db:"phone" json:"phone"`
		FirstName        string `db:"first_name" json:"first_name"`
		LastName         string `db:"last_name" json:"last_name"`
		Username         string `db:"username" json:"username"`
		Code             string `db:"code" json:"code"`
		Verified         bool   `db:"verified" json:"verified"`
		RegisteredUserId int64  `db:"registered_user_id" json:"registered_user_id"`
		Deleted          bool   `db:"deleted" json:"deleted"`
	}
)

func newPredefinedUsersModel(db *sqlx.DB) *defaultPredefinedUsersModel {
	return &defaultPredefinedUsersModel{
		db: db,
	}
}

func (m *defaultPredefinedUsersModel) Insert2(ctx context.Context, data *PredefinedUsers) (sql.Result, error) {
	query := fmt.Sprintf("insert into `predefined_users` (%s) values (?, ?, ?, ?, ?, ?, ?, ?)", predefinedUsersRowsExpectAutoSet)

	r, err := m.db.Exec(ctx, query, data.Phone, data.FirstName, data.LastName, data.Username, data.Code, data.Verified, data.RegisteredUserId, data.Deleted)
	if err != nil {
		return nil, fmt.Errorf("predefined_users.Insert2 exec: %w", err)
	}

	return r, nil
}

func (m *defaultPredefinedUsersModel) Delete2(ctx context.Context, id int64) error {
	query := "delete from `predefined_users` where `id` = ?"

	_, err := m.db.Exec(ctx, query, id)
	if err != nil {
		return fmt.Errorf("predefined_users.Delete2 exec: %w", err)
	}

	return nil
}

func (m *defaultPredefinedUsersModel) FindOne(ctx context.Context, id int64) (*PredefinedUsers, error) {
	query := fmt.Sprintf("select %s from predefined_users where id = ? limit 1", predefinedUsersRows)
	var resp PredefinedUsers

	err := m.db.QueryRowPartial(ctx, &resp, query, id)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return nil, &NotFoundError{
				Resource: "predefined_users",
				Key:      fmt.Sprintf("id=%v", id),
				Cause:    err,
			}
		}
		return nil, fmt.Errorf("predefined_users.FindOne: %w", err)
	}

	return &resp, nil
}

func (m *defaultPredefinedUsersModel) FindListByIdList(ctx context.Context, id ...int64) ([]PredefinedUsers, error) {
	if len(id) == 0 {
		return []PredefinedUsers{}, nil
	}

	query := fmt.Sprintf("select %s from predefined_users where id in (%s)", predefinedUsersRows, sqlx.InInt64List(id))

	var resp []PredefinedUsers
	err := m.db.QueryRowsPartial(ctx, &resp, query)
	if err != nil {
		return nil, fmt.Errorf("predefined_users.FindListByIdList: %w", err)
	}

	return resp, nil
}

func (m *defaultPredefinedUsersModel) Update2(ctx context.Context, data *PredefinedUsers) error {
	query := fmt.Sprintf("update `predefined_users` set %s where `id` = ?", predefinedUsersRowsWithPlaceHolder)

	_, err := m.db.Exec(ctx, query, data.Phone, data.FirstName, data.LastName, data.Username, data.Code, data.Verified, data.RegisteredUserId, data.Deleted, data.Id)
	if err != nil {
		return fmt.Errorf("predefined_users.Update2 exec: %w", err)
	}

	return nil
}
