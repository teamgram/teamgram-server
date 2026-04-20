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
	default_history_ttlFieldNames          = builder.RawFieldNames(&DefaultHistoryTtl{})
	default_history_ttlRows                = strings.Join(default_history_ttlFieldNames, ",")
	default_history_ttlRowsExpectAutoSet   = strings.Join(stringx.Remove(default_history_ttlFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	default_history_ttlRowsWithPlaceHolder = strings.Join(stringx.Remove(default_history_ttlFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"

	cacheTDefaultHistoryTtlIdPrefix = "cache:t:default_history_ttl:id:"

	cacheDefaultHistoryTtlIdPrefix = "cache#DefaultHistoryTtl#id"

	cacheDefaultHistoryTtlUserIdPrefix = "cache#UserId"
)

type (
	default_history_ttlModel interface {
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
	query := fmt.Sprintf("insert into `default_history_ttl` (%s) values (?, ?)", default_history_ttlRowsExpectAutoSet)
	return m.db.Exec(ctx, query, data.UserId, data.Period)
}

func (m *defaultDefaultHistoryTtlModel) Delete2(ctx context.Context, id int64) error {
	query := "delete from `default_history_ttl` where `id` = ?"
	_, err := m.db.Exec(ctx, query, id)
	return err
}

func (m *defaultDefaultHistoryTtlModel) FindOne(ctx context.Context, id int64) (*DefaultHistoryTtl, error) {
	query := fmt.Sprintf("select %s from default_history_ttl where id = ? limit 1", default_history_ttlRows)
	var resp DefaultHistoryTtl
	err := m.db.QueryRowPartial(ctx, &resp, query, id)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (m *defaultDefaultHistoryTtlModel) FindListByIdList(ctx context.Context, id ...int64) ([]DefaultHistoryTtl, error) {
	if len(id) == 0 {
		return []DefaultHistoryTtl{}, nil
	}

	query := fmt.Sprintf("select %s from default_history_ttl where id in (%s)", default_history_ttlRows, sqlx.InInt64List(id))

	var resp []DefaultHistoryTtl
	err := m.db.QueryRowsPartial(ctx, &resp, query)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (m *defaultDefaultHistoryTtlModel) Update2(ctx context.Context, data *DefaultHistoryTtl) error {
	query := fmt.Sprintf("update `default_history_ttl` set %s where `id` = ?", default_history_ttlRowsWithPlaceHolder)
	_, err := m.db.Exec(ctx, query, data.UserId, data.Period, data.Id)
	return err
}

func (m *defaultDefaultHistoryTtlModel) FindOneByUserId(ctx context.Context, userId int64) (*DefaultHistoryTtl, error) {
	query := fmt.Sprintf("select %s from default_history_ttl where user_id = ? limit 1", default_history_ttlRows)
	var resp DefaultHistoryTtl
	err := m.db.QueryRowPartial(ctx, &resp, query, userId)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (m *defaultDefaultHistoryTtlModel) FindListByUserIdList(ctx context.Context, userId ...int64) ([]DefaultHistoryTtl, error) {
	if len(userId) == 0 {
		return []DefaultHistoryTtl{}, nil
	}

	query := fmt.Sprintf("select %s from default_history_ttl where user_id in (%s)", default_history_ttlRows, sqlx.InInt64List(userId))

	var resp []DefaultHistoryTtl
	err := m.db.QueryRowsPartial(ctx, &resp, query)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (m *defaultDefaultHistoryTtlModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s#%v", cacheDefaultHistoryTtlIdPrefix, primary)
}

func (m *defaultDefaultHistoryTtlModel) queryPrimary(ctx context.Context, v interface{}, primary interface{}) error {
	query := fmt.Sprintf("select %s from default_history_ttl where id = ? limit 1", default_history_ttlRows)
	return m.db.QueryRowPartial(ctx, v, query, primary)
}
