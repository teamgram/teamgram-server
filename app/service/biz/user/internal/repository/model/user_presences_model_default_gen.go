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
	userPresencesFieldNames          = builder.RawFieldNames(&UserPresences{})
	userPresencesRows                = strings.Join(userPresencesFieldNames, ",")
	userPresencesRowsExpectAutoSet   = strings.Join(stringx.Remove(userPresencesFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	userPresencesRowsWithPlaceHolder = strings.Join(stringx.Remove(userPresencesFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"
)

type (
	userPresencesModel interface {
		Insert2(ctx context.Context, data *UserPresences) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*UserPresences, error)
		FindListByIdList(ctx context.Context, id ...int64) ([]UserPresences, error)
		Update2(ctx context.Context, data *UserPresences) error
		Delete2(ctx context.Context, id int64) error

		FindOneByUserId(ctx context.Context, userId int64) (*UserPresences, error)
		FindListByUserIdList(ctx context.Context, userId ...int64) ([]UserPresences, error)
	}

	defaultUserPresencesModel struct {
		db *sqlx.DB
	}

	UserPresences struct {
		Id         int64 `db:"id" json:"id"`
		UserId     int64 `db:"user_id" json:"user_id"`
		LastSeenAt int64 `db:"last_seen_at" json:"last_seen_at"`
		Expires    int32 `db:"expires" json:"expires"`
	}
)

func newUserPresencesModel(db *sqlx.DB) *defaultUserPresencesModel {
	return &defaultUserPresencesModel{
		db: db,
	}
}

func (m *defaultUserPresencesModel) Insert2(ctx context.Context, data *UserPresences) (sql.Result, error) {
	query := fmt.Sprintf("insert into `user_presences` (%s) values (?, ?, ?)", userPresencesRowsExpectAutoSet)

	r, err := m.db.Exec(ctx, query, data.UserId, data.LastSeenAt, data.Expires)
	if err != nil {
		return nil, fmt.Errorf("user_presences.Insert2 exec: %w", err)
	}

	return r, nil
}

func (m *defaultUserPresencesModel) Delete2(ctx context.Context, id int64) error {
	query := "delete from `user_presences` where `id` = ?"

	_, err := m.db.Exec(ctx, query, id)
	if err != nil {
		return fmt.Errorf("user_presences.Delete2 exec: %w", err)
	}

	return nil
}

func (m *defaultUserPresencesModel) FindOne(ctx context.Context, id int64) (*UserPresences, error) {
	query := fmt.Sprintf("select %s from user_presences where id = ? limit 1", userPresencesRows)
	var resp UserPresences

	err := m.db.QueryRowPartial(ctx, &resp, query, id)
	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return nil, nil
		}
		return nil, fmt.Errorf("user_presences.FindOne: %w", err)
	}

	return &resp, nil
}

func (m *defaultUserPresencesModel) FindListByIdList(ctx context.Context, id ...int64) ([]UserPresences, error) {
	if len(id) == 0 {
		return []UserPresences{}, nil
	}

	query := fmt.Sprintf("select %s from user_presences where id in (%s)", userPresencesRows, sqlx.InInt64List(id))

	var resp []UserPresences
	err := m.db.QueryRowsPartial(ctx, &resp, query)
	if err != nil {
		return nil, fmt.Errorf("user_presences.FindListByIdList: %w", err)
	}

	return resp, nil
}

func (m *defaultUserPresencesModel) Update2(ctx context.Context, data *UserPresences) error {
	query := fmt.Sprintf("update `user_presences` set %s where `id` = ?", userPresencesRowsWithPlaceHolder)

	_, err := m.db.Exec(ctx, query, data.UserId, data.LastSeenAt, data.Expires, data.Id)
	if err != nil {
		return fmt.Errorf("user_presences.Update2 exec: %w", err)
	}

	return nil
}

func (m *defaultUserPresencesModel) FindOneByUserId(ctx context.Context, userId int64) (*UserPresences, error) {
	query := fmt.Sprintf("select %s from user_presences where user_id = ? limit 1", userPresencesRows)
	var resp UserPresences

	err := m.db.QueryRowPartial(ctx, &resp, query, userId)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return nil, nil
		}
		return nil, fmt.Errorf("user_presences.FindOneByUserId: %w", err)
	}

	return &resp, nil
}

func (m *defaultUserPresencesModel) FindListByUserIdList(ctx context.Context, userId ...int64) ([]UserPresences, error) {
	if len(userId) == 0 {
		return []UserPresences{}, nil
	}

	query := fmt.Sprintf("select %s from user_presences where user_id in (%s)", userPresencesRows, sqlx.InInt64List(userId))

	var resp []UserPresences
	err := m.db.QueryRowsPartial(ctx, &resp, query)
	if err != nil {
		return nil, fmt.Errorf("user_presences.FindListByUserIdList: %w", err)
	}

	return resp, nil
}
