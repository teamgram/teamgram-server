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
	authOpLogsFieldNames          = builder.RawFieldNames(&AuthOpLogs{})
	authOpLogsRows                = strings.Join(authOpLogsFieldNames, ",")
	authOpLogsRowsExpectAutoSet   = strings.Join(stringx.Remove(authOpLogsFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	authOpLogsRowsWithPlaceHolder = strings.Join(stringx.Remove(authOpLogsFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"
)

type (
	authOpLogsModel interface {
		Insert2(ctx context.Context, data *AuthOpLogs) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*AuthOpLogs, error)
		FindListByIdList(ctx context.Context, id ...int64) ([]AuthOpLogs, error)
		Update2(ctx context.Context, data *AuthOpLogs) error
		Delete2(ctx context.Context, id int64) error
	}

	defaultAuthOpLogsModel struct {
		db *sqlx.DB
	}

	AuthOpLogs struct {
		Id        int64  `db:"id" json:"id"`
		AuthKeyId int64  `db:"auth_key_id" json:"auth_key_id"`
		Ip        string `db:"ip" json:"ip"`
		OpType    int32  `db:"op_type" json:"op_type"`
		LogText   string `db:"log_text" json:"log_text"`
	}
)

func newAuthOpLogsModel(db *sqlx.DB) *defaultAuthOpLogsModel {
	return &defaultAuthOpLogsModel{
		db: db,
	}
}

func (m *defaultAuthOpLogsModel) Insert2(ctx context.Context, data *AuthOpLogs) (sql.Result, error) {
	tableName := "auth_op_logs"
	query := fmt.Sprintf("insert into `%s` (%s) values (?, ?, ?, ?)", tableName, authOpLogsRowsExpectAutoSet)

	r, err := m.db.Exec(ctx, query, data.AuthKeyId, data.Ip, data.OpType, data.LogText)
	if err != nil {
		return nil, fmt.Errorf("auth_op_logs.Insert2 exec: %w", err)
	}

	return r, nil
}

func (m *defaultAuthOpLogsModel) Delete2(ctx context.Context, id int64) error {
	tableName := "auth_op_logs"
	query := fmt.Sprintf("delete from `%s` where `id` = ?", tableName)

	_, err := m.db.Exec(ctx, query, id)
	if err != nil {
		return fmt.Errorf("auth_op_logs.Delete2 exec: %w", err)
	}

	return nil
}

func (m *defaultAuthOpLogsModel) FindOne(ctx context.Context, id int64) (*AuthOpLogs, error) {
	tableName := "auth_op_logs"
	query := fmt.Sprintf("select %s from %s where id = ? limit 1", authOpLogsRows, tableName)
	var resp AuthOpLogs

	err := m.db.QueryRowPartial(ctx, &resp, query, id)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return nil, &NotFoundError{
				Resource: "auth_op_logs",
				Key:      fmt.Sprintf("id=%v", id),
				Cause:    err,
			}
		}
		return nil, fmt.Errorf("auth_op_logs.FindOne: %w", err)
	}

	return &resp, nil
}

func (m *defaultAuthOpLogsModel) FindListByIdList(ctx context.Context, id ...int64) ([]AuthOpLogs, error) {
	if len(id) == 0 {
		return []AuthOpLogs{}, nil
	}
	tableName := "auth_op_logs"

	query := fmt.Sprintf("select %s from %s where id in (%s)", authOpLogsRows, tableName, sqlx.InInt64List(id))

	var resp []AuthOpLogs
	err := m.db.QueryRowsPartial(ctx, &resp, query)
	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return []AuthOpLogs{}, nil
		}
		return nil, fmt.Errorf("auth_op_logs.FindListByIdList: %w", err)
	}

	return resp, nil
}

func (m *defaultAuthOpLogsModel) Update2(ctx context.Context, data *AuthOpLogs) error {
	tableName := "auth_op_logs"
	query := fmt.Sprintf("update `%s` set %s where `id` = ?", tableName, authOpLogsRowsWithPlaceHolder)

	_, err := m.db.Exec(ctx, query, data.AuthKeyId, data.Ip, data.OpType, data.LogText, data.Id)
	if err != nil {
		return fmt.Errorf("auth_op_logs.Update2 exec: %w", err)
	}

	return nil
}
