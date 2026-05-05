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
	userAuthSeqStateFieldNames          = builder.RawFieldNames(&UserAuthSeqState{})
	userAuthSeqStateRows                = strings.Join(userAuthSeqStateFieldNames, ",")
	userAuthSeqStateRowsExpectAutoSet   = strings.Join(stringx.Remove(userAuthSeqStateFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	userAuthSeqStateRowsWithPlaceHolder = strings.Join(stringx.Remove(userAuthSeqStateFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"
)

type (
	userAuthSeqStateModel interface {
		Insert2(ctx context.Context, data *UserAuthSeqState) (sql.Result, error)
		FindOne(ctx context.Context, userId int64) (*UserAuthSeqState, error)
		FindListByUserIdList(ctx context.Context, userId ...int64) ([]UserAuthSeqState, error)
		Update2(ctx context.Context, data *UserAuthSeqState) error
		Delete2(ctx context.Context, userId int64) error
	}

	defaultUserAuthSeqStateModel struct {
		db *sqlx.DB
	}

	UserAuthSeqState struct {
		UserId     int64 `db:"user_id" json:"user_id"`
		Seq        int64 `db:"seq" json:"seq"`
		Date       int32 `db:"date" json:"date"`
		RowVersion int64 `db:"row_version" json:"row_version"`
	}
)

func newUserAuthSeqStateModel(db *sqlx.DB) *defaultUserAuthSeqStateModel {
	return &defaultUserAuthSeqStateModel{
		db: db,
	}
}

func (m *defaultUserAuthSeqStateModel) Insert2(ctx context.Context, data *UserAuthSeqState) (sql.Result, error) {
	tableName := "user_auth_seq_state"
	query := fmt.Sprintf("insert into `%s` (%s) values (?, ?, ?)", tableName, userAuthSeqStateRowsExpectAutoSet)

	r, err := m.db.Exec(ctx, query, data.Seq, data.Date, data.RowVersion)
	if err != nil {
		return nil, fmt.Errorf("user_auth_seq_state.Insert2 exec: %w", err)
	}

	return r, nil
}

func (m *defaultUserAuthSeqStateModel) Delete2(ctx context.Context, userId int64) error {
	tableName := "user_auth_seq_state"
	query := fmt.Sprintf("delete from `%s` where `user_id` = ?", tableName)

	_, err := m.db.Exec(ctx, query, userId)
	if err != nil {
		return fmt.Errorf("user_auth_seq_state.Delete2 exec: %w", err)
	}

	return nil
}

func (m *defaultUserAuthSeqStateModel) FindOne(ctx context.Context, userId int64) (*UserAuthSeqState, error) {
	tableName := "user_auth_seq_state"
	query := fmt.Sprintf("select %s from %s where user_id = ? limit 1", userAuthSeqStateRows, tableName)
	var resp UserAuthSeqState

	err := m.db.QueryRowPartial(ctx, &resp, query, userId)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return nil, &NotFoundError{
				Resource: "user_auth_seq_state",
				Key:      fmt.Sprintf("user_id=%v", userId),
				Cause:    err,
			}
		}
		return nil, fmt.Errorf("user_auth_seq_state.FindOne: %w", err)
	}

	return &resp, nil
}

func (m *defaultUserAuthSeqStateModel) FindListByUserIdList(ctx context.Context, userId ...int64) ([]UserAuthSeqState, error) {
	if len(userId) == 0 {
		return []UserAuthSeqState{}, nil
	}
	tableName := "user_auth_seq_state"

	query := fmt.Sprintf("select %s from %s where user_id in (%s)", userAuthSeqStateRows, tableName, sqlx.InInt64List(userId))

	var resp []UserAuthSeqState
	err := m.db.QueryRowsPartial(ctx, &resp, query)
	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return []UserAuthSeqState{}, nil
		}
		return nil, fmt.Errorf("user_auth_seq_state.FindListByUserIdList: %w", err)
	}

	return resp, nil
}

func (m *defaultUserAuthSeqStateModel) Update2(ctx context.Context, data *UserAuthSeqState) error {
	tableName := "user_auth_seq_state"
	query := fmt.Sprintf("update `%s` set %s where `user_id` = ?", tableName, userAuthSeqStateRowsWithPlaceHolder)

	_, err := m.db.Exec(ctx, query, data.Seq, data.Date, data.RowVersion, data.UserId)
	if err != nil {
		return fmt.Errorf("user_auth_seq_state.Update2 exec: %w", err)
	}

	return nil
}
