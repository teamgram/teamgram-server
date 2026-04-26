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
	defaultHistoryTtlFieldNames          = builder.RawFieldNames(&DefaultHistoryTtl{})
	defaultHistoryTtlRows                = strings.Join(defaultHistoryTtlFieldNames, ",")
	defaultHistoryTtlRowsExpectAutoSet   = strings.Join(stringx.Remove(defaultHistoryTtlFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	defaultHistoryTtlRowsWithPlaceHolder = strings.Join(stringx.Remove(defaultHistoryTtlFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"
)

type (
	defaultHistoryTtlModel interface {
		Insert2(ctx context.Context, data *DefaultHistoryTtl) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*DefaultHistoryTtl, error)
		FindListByIdList(ctx context.Context, id ...int64) ([]DefaultHistoryTtl, error)
		Update2(ctx context.Context, data *DefaultHistoryTtl) error
		Delete2(ctx context.Context, id int64) error

		FindOneByUserId(ctx context.Context, userId int64) (*DefaultHistoryTtl, error)
		FindListByUserIdList(ctx context.Context, userId ...int64) ([]DefaultHistoryTtl, error)
	}

	defaultDefaultHistoryTtlModel struct {
		db *sqlx.DB
	}

	DefaultHistoryTtl struct {
		Id     int64 `db:"id" json:"id"`
		UserId int64 `db:"user_id" json:"user_id"`
		Period int32 `db:"period" json:"period"`
	}
)

func newDefaultHistoryTtlModel(db *sqlx.DB) *defaultDefaultHistoryTtlModel {
	return &defaultDefaultHistoryTtlModel{
		db: db,
	}
}

func (m *defaultDefaultHistoryTtlModel) Insert2(ctx context.Context, data *DefaultHistoryTtl) (sql.Result, error) {
	query := fmt.Sprintf("insert into `default_history_ttl` (%s) values (?, ?)", defaultHistoryTtlRowsExpectAutoSet)

	r, err := m.db.Exec(ctx, query, data.UserId, data.Period)
	if err != nil {
		return nil, fmt.Errorf("default_history_ttl.Insert2 exec: %w", err)
	}

	return r, nil
}

func (m *defaultDefaultHistoryTtlModel) Delete2(ctx context.Context, id int64) error {
	query := "delete from `default_history_ttl` where `id` = ?"

	_, err := m.db.Exec(ctx, query, id)
	if err != nil {
		return fmt.Errorf("default_history_ttl.Delete2 exec: %w", err)
	}

	return nil
}

func (m *defaultDefaultHistoryTtlModel) FindOne(ctx context.Context, id int64) (*DefaultHistoryTtl, error) {
	query := fmt.Sprintf("select %s from default_history_ttl where id = ? limit 1", defaultHistoryTtlRows)
	var resp DefaultHistoryTtl

	err := m.db.QueryRowPartial(ctx, &resp, query, id)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return nil, &NotFoundError{
				Resource: "default_history_ttl",
				Key:      fmt.Sprintf("id=%v", id),
				Cause:    err,
			}
		}
		return nil, fmt.Errorf("default_history_ttl.FindOne: %w", err)
	}

	return &resp, nil
}

func (m *defaultDefaultHistoryTtlModel) FindListByIdList(ctx context.Context, id ...int64) ([]DefaultHistoryTtl, error) {
	if len(id) == 0 {
		return []DefaultHistoryTtl{}, nil
	}

	query := fmt.Sprintf("select %s from default_history_ttl where id in (%s)", defaultHistoryTtlRows, sqlx.InInt64List(id))

	var resp []DefaultHistoryTtl
	err := m.db.QueryRowsPartial(ctx, &resp, query)
	if err != nil {
		return nil, fmt.Errorf("default_history_ttl.FindListByIdList: %w", err)
	}

	return resp, nil
}

func (m *defaultDefaultHistoryTtlModel) Update2(ctx context.Context, data *DefaultHistoryTtl) error {
	query := fmt.Sprintf("update `default_history_ttl` set %s where `id` = ?", defaultHistoryTtlRowsWithPlaceHolder)

	_, err := m.db.Exec(ctx, query, data.UserId, data.Period, data.Id)
	if err != nil {
		return fmt.Errorf("default_history_ttl.Update2 exec: %w", err)
	}

	return nil
}

func (m *defaultDefaultHistoryTtlModel) FindOneByUserId(ctx context.Context, userId int64) (*DefaultHistoryTtl, error) {
	query := fmt.Sprintf("select %s from default_history_ttl where user_id = ? limit 1", defaultHistoryTtlRows)
	var resp DefaultHistoryTtl

	err := m.db.QueryRowPartial(ctx, &resp, query, userId)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return nil, &NotFoundError{
				Resource: "default_history_ttl",
				Key:      fmt.Sprintf("user_id=%v", userId),
				Cause:    err,
			}
		}
		return nil, fmt.Errorf("default_history_ttl.FindOneByUserId: %w", err)
	}

	return &resp, nil
}

func (m *defaultDefaultHistoryTtlModel) FindListByUserIdList(ctx context.Context, userId ...int64) ([]DefaultHistoryTtl, error) {
	if len(userId) == 0 {
		return []DefaultHistoryTtl{}, nil
	}

	query := fmt.Sprintf("select %s from default_history_ttl where user_id in (%s)", defaultHistoryTtlRows, sqlx.InInt64List(userId))

	var resp []DefaultHistoryTtl
	err := m.db.QueryRowsPartial(ctx, &resp, query)
	if err != nil {
		return nil, fmt.Errorf("default_history_ttl.FindListByUserIdList: %w", err)
	}

	return resp, nil
}
