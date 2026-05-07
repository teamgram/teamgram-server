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
	userPtsStateFieldNames          = builder.RawFieldNames(&UserPtsState{})
	userPtsStateRows                = strings.Join(userPtsStateFieldNames, ",")
	userPtsStateRowsExpectAutoSet   = strings.Join(stringx.Remove(userPtsStateFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	userPtsStateRowsWithPlaceHolder = strings.Join(stringx.Remove(userPtsStateFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"
)

type (
	userPtsStateModel interface {
		Insert2(ctx context.Context, data *UserPtsState) (sql.Result, error)
		FindOne(ctx context.Context, userId int64) (*UserPtsState, error)
		FindListByUserIdList(ctx context.Context, userId ...int64) ([]UserPtsState, error)
		Update2(ctx context.Context, data *UserPtsState) error
		Delete2(ctx context.Context, userId int64) error
	}

	defaultUserPtsStateModel struct {
		db *sqlx.DB
	}

	UserPtsState struct {
		UserId       int64 `db:"user_id" json:"user_id"`
		Pts          int64 `db:"pts" json:"pts"`
		PtsUpdatedAt int64 `db:"pts_updated_at" json:"pts_updated_at"`
		PartitionId  int32 `db:"partition_id" json:"partition_id"`
		OwnerEpoch   int64 `db:"owner_epoch" json:"owner_epoch"`
		RowVersion   int64 `db:"row_version" json:"row_version"`
	}
)

func newUserPtsStateModel(db *sqlx.DB) *defaultUserPtsStateModel {
	return &defaultUserPtsStateModel{
		db: db,
	}
}

func (m *defaultUserPtsStateModel) Insert2(ctx context.Context, data *UserPtsState) (sql.Result, error) {
	tableName := "user_pts_state"
	query := fmt.Sprintf("insert into `%s` (%s) values (?, ?, ?, ?, ?)", tableName, userPtsStateRowsExpectAutoSet)

	r, err := m.db.Exec(ctx, query, data.Pts, data.PtsUpdatedAt, data.PartitionId, data.OwnerEpoch, data.RowVersion)
	if err != nil {
		return nil, fmt.Errorf("user_pts_state.Insert2 exec: %w", err)
	}

	return r, nil
}

func (m *defaultUserPtsStateModel) Delete2(ctx context.Context, userId int64) error {
	tableName := "user_pts_state"
	query := fmt.Sprintf("delete from `%s` where `user_id` = ?", tableName)

	_, err := m.db.Exec(ctx, query, userId)
	if err != nil {
		return fmt.Errorf("user_pts_state.Delete2 exec: %w", err)
	}

	return nil
}

func (m *defaultUserPtsStateModel) FindOne(ctx context.Context, userId int64) (*UserPtsState, error) {
	tableName := "user_pts_state"
	query := fmt.Sprintf("select %s from %s where user_id = ? limit 1", userPtsStateRows, tableName)
	var resp UserPtsState

	err := m.db.QueryRowPartial(ctx, &resp, query, userId)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return nil, &NotFoundError{
				Resource: "user_pts_state",
				Key:      fmt.Sprintf("user_id=%v", userId),
				Cause:    err,
			}
		}
		return nil, fmt.Errorf("user_pts_state.FindOne: %w", err)
	}

	return &resp, nil
}

func (m *defaultUserPtsStateModel) FindListByUserIdList(ctx context.Context, userId ...int64) ([]UserPtsState, error) {
	if len(userId) == 0 {
		return []UserPtsState{}, nil
	}
	tableName := "user_pts_state"

	query := fmt.Sprintf("select %s from %s where user_id in (%s)", userPtsStateRows, tableName, sqlx.InInt64List(userId))

	var resp []UserPtsState
	err := m.db.QueryRowsPartial(ctx, &resp, query)
	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return []UserPtsState{}, nil
		}
		return nil, fmt.Errorf("user_pts_state.FindListByUserIdList: %w", err)
	}

	return resp, nil
}

func (m *defaultUserPtsStateModel) Update2(ctx context.Context, data *UserPtsState) error {
	tableName := "user_pts_state"
	query := fmt.Sprintf("update `%s` set %s where `user_id` = ?", tableName, userPtsStateRowsWithPlaceHolder)

	_, err := m.db.Exec(ctx, query, data.Pts, data.PtsUpdatedAt, data.PartitionId, data.OwnerEpoch, data.RowVersion, data.UserId)
	if err != nil {
		return fmt.Errorf("user_pts_state.Update2 exec: %w", err)
	}

	return nil
}
