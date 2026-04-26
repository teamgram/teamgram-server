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
	userPtsUpdatesFieldNames          = builder.RawFieldNames(&UserPtsUpdates{})
	userPtsUpdatesRows                = strings.Join(userPtsUpdatesFieldNames, ",")
	userPtsUpdatesRowsExpectAutoSet   = strings.Join(stringx.Remove(userPtsUpdatesFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	userPtsUpdatesRowsWithPlaceHolder = strings.Join(stringx.Remove(userPtsUpdatesFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"
)

type (
	userPtsUpdatesModel interface {
		Insert2(ctx context.Context, data *UserPtsUpdates) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*UserPtsUpdates, error)
		FindListByIdList(ctx context.Context, id ...int64) ([]UserPtsUpdates, error)
		Update2(ctx context.Context, data *UserPtsUpdates) error
		Delete2(ctx context.Context, id int64) error
	}

	defaultUserPtsUpdatesModel struct {
		db *sqlx.DB
	}

	UserPtsUpdates struct {
		Id         int64  `db:"id" json:"id"`
		UserId     int64  `db:"user_id" json:"user_id"`
		Pts        int32  `db:"pts" json:"pts"`
		PtsCount   int32  `db:"pts_count" json:"pts_count"`
		UpdateType int32  `db:"update_type" json:"update_type"`
		UpdateData string `db:"update_data" json:"update_data"`
		Date2      int64  `db:"date2" json:"date2"`
	}
)

func newUserPtsUpdatesModel(db *sqlx.DB) *defaultUserPtsUpdatesModel {
	return &defaultUserPtsUpdatesModel{
		db: db,
	}
}

func (m *defaultUserPtsUpdatesModel) Insert2(ctx context.Context, data *UserPtsUpdates) (sql.Result, error) {
	query := fmt.Sprintf("insert into `user_pts_updates` (%s) values (?, ?, ?, ?, ?, ?)", userPtsUpdatesRowsExpectAutoSet)

	r, err := m.db.Exec(ctx, query, data.UserId, data.Pts, data.PtsCount, data.UpdateType, data.UpdateData, data.Date2)
	if err != nil {
		return nil, fmt.Errorf("user_pts_updates.Insert2 exec: %w", err)
	}

	return r, nil
}

func (m *defaultUserPtsUpdatesModel) Delete2(ctx context.Context, id int64) error {
	query := "delete from `user_pts_updates` where `id` = ?"

	_, err := m.db.Exec(ctx, query, id)
	if err != nil {
		return fmt.Errorf("user_pts_updates.Delete2 exec: %w", err)
	}

	return nil
}

func (m *defaultUserPtsUpdatesModel) FindOne(ctx context.Context, id int64) (*UserPtsUpdates, error) {
	query := fmt.Sprintf("select %s from user_pts_updates where id = ? limit 1", userPtsUpdatesRows)
	var resp UserPtsUpdates

	err := m.db.QueryRowPartial(ctx, &resp, query, id)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return nil, &NotFoundError{
				Resource: "user_pts_updates",
				Key:      fmt.Sprintf("id=%v", id),
				Cause:    err,
			}
		}
		return nil, fmt.Errorf("user_pts_updates.FindOne: %w", err)
	}

	return &resp, nil
}

func (m *defaultUserPtsUpdatesModel) FindListByIdList(ctx context.Context, id ...int64) ([]UserPtsUpdates, error) {
	if len(id) == 0 {
		return []UserPtsUpdates{}, nil
	}

	query := fmt.Sprintf("select %s from user_pts_updates where id in (%s)", userPtsUpdatesRows, sqlx.InInt64List(id))

	var resp []UserPtsUpdates
	err := m.db.QueryRowsPartial(ctx, &resp, query)
	if err != nil {
		return nil, fmt.Errorf("user_pts_updates.FindListByIdList: %w", err)
	}

	return resp, nil
}

func (m *defaultUserPtsUpdatesModel) Update2(ctx context.Context, data *UserPtsUpdates) error {
	query := fmt.Sprintf("update `user_pts_updates` set %s where `id` = ?", userPtsUpdatesRowsWithPlaceHolder)

	_, err := m.db.Exec(ctx, query, data.UserId, data.Pts, data.PtsCount, data.UpdateType, data.UpdateData, data.Date2, data.Id)
	if err != nil {
		return fmt.Errorf("user_pts_updates.Update2 exec: %w", err)
	}

	return nil
}
