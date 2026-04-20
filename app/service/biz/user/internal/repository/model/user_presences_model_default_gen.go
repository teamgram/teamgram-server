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
	user_presencesFieldNames          = builder.RawFieldNames(&UserPresences{})
	user_presencesRows                = strings.Join(user_presencesFieldNames, ",")
	user_presencesRowsExpectAutoSet   = strings.Join(stringx.Remove(user_presencesFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	user_presencesRowsWithPlaceHolder = strings.Join(stringx.Remove(user_presencesFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"

	cacheTUserPresencesIdPrefix = "cache:t:user_presences:id:"

	cacheUserPresencesIdPrefix = "cache#UserPresences#id"

	cacheUserPresencesUserIdPrefix = "cache#UserId"
)

type (
	user_presencesModel interface {
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
	query := fmt.Sprintf("insert into `user_presences` (%s) values (?, ?, ?)", user_presencesRowsExpectAutoSet)
	return m.db.Exec(ctx, query, data.UserId, data.LastSeenAt, data.Expires)
}

func (m *defaultUserPresencesModel) Delete2(ctx context.Context, id int64) error {
	query := "delete from `user_presences` where `id` = ?"
	_, err := m.db.Exec(ctx, query, id)
	return err
}

func (m *defaultUserPresencesModel) FindOne(ctx context.Context, id int64) (*UserPresences, error) {
	query := fmt.Sprintf("select %s from user_presences where id = ? limit 1", user_presencesRows)
	var resp UserPresences
	err := m.db.QueryRowPartial(ctx, &resp, query, id)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (m *defaultUserPresencesModel) FindListByIdList(ctx context.Context, id ...int64) ([]UserPresences, error) {
	if len(id) == 0 {
		return []UserPresences{}, nil
	}

	query := fmt.Sprintf("select %s from user_presences where id in (%s)", user_presencesRows, sqlx.InInt64List(id))

	var resp []UserPresences
	err := m.db.QueryRowsPartial(ctx, &resp, query)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (m *defaultUserPresencesModel) Update2(ctx context.Context, data *UserPresences) error {
	query := fmt.Sprintf("update `user_presences` set %s where `id` = ?", user_presencesRowsWithPlaceHolder)
	_, err := m.db.Exec(ctx, query, data.UserId, data.LastSeenAt, data.Expires, data.Id)
	return err
}

func (m *defaultUserPresencesModel) FindOneByUserId(ctx context.Context, userId int64) (*UserPresences, error) {
	query := fmt.Sprintf("select %s from user_presences where user_id = ? limit 1", user_presencesRows)
	var resp UserPresences
	err := m.db.QueryRowPartial(ctx, &resp, query, userId)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (m *defaultUserPresencesModel) FindListByUserIdList(ctx context.Context, userId ...int64) ([]UserPresences, error) {
	if len(userId) == 0 {
		return []UserPresences{}, nil
	}

	query := fmt.Sprintf("select %s from user_presences where user_id in (%s)", user_presencesRows, sqlx.InInt64List(userId))

	var resp []UserPresences
	err := m.db.QueryRowsPartial(ctx, &resp, query)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (m *defaultUserPresencesModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s#%v", cacheUserPresencesIdPrefix, primary)
}

func (m *defaultUserPresencesModel) queryPrimary(ctx context.Context, v interface{}, primary interface{}) error {
	query := fmt.Sprintf("select %s from user_presences where id = ? limit 1", user_presencesRows)
	return m.db.QueryRowPartial(ctx, v, query, primary)
}
