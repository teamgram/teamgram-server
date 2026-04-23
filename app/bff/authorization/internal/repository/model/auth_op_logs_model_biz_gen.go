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
	"github.com/zeromicro/go-zero/core/logx"
)

var _ *sql.Result
var _ = fmt.Sprintf
var _ = strings.Join
var _ = errors.Is
var _ *sqlx.DB
var _ *logx.Logger

type (
	bizAuthOpLogsModel interface {
		Insert(ctx context.Context, data *AuthOpLogs) (lastInsertId, rowsAffected int64, err error)
		InsertTx(tx *sqlx.Tx, data *AuthOpLogs) (lastInsertId, rowsAffected int64, err error)
	}
)

// Insert
// insert into auth_op_logs(auth_key_id, ip, op_type, log_text) values (:auth_key_id, :ip, :op_type, :log_text)
func (m *defaultAuthOpLogsModel) Insert(ctx context.Context, data *AuthOpLogs) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into auth_op_logs(auth_key_id, ip, op_type, log_text) values (:auth_key_id, :ip, :op_type, :log_text)"
		r     sql.Result
	)

	r, err = m.db.NamedExec(ctx, query, data)
	if err != nil {
		logx.WithContext(ctx).Errorf("namedExec in Insert(%v), error: %v", data, err)
		return
	}

	lastInsertId, err = r.LastInsertId()
	if err != nil {
		logx.WithContext(ctx).Errorf("lastInsertId in Insert(%v)_error: %v", data, err)
		return
	}
	rowsAffected, err = r.RowsAffected()
	if err != nil {
		logx.WithContext(ctx).Errorf("rowsAffected in Insert(%v)_error: %v", data, err)
	}

	return

}

// InsertTx
// insert into auth_op_logs(auth_key_id, ip, op_type, log_text) values (:auth_key_id, :ip, :op_type, :log_text)
func (m *defaultAuthOpLogsModel) InsertTx(tx *sqlx.Tx, data *AuthOpLogs) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into auth_op_logs(auth_key_id, ip, op_type, log_text) values (:auth_key_id, :ip, :op_type, :log_text)"
		r     sql.Result
	)

	r, err = tx.NamedExec(query, data)
	if err != nil {
		logx.WithContext(tx.Context()).Errorf("namedExec in Insert(%v), error: %v", data, err)
		return
	}

	lastInsertId, err = r.LastInsertId()
	if err != nil {
		logx.WithContext(tx.Context()).Errorf("lastInsertId in Insert(%v)_error: %v", data, err)
		return
	}
	rowsAffected, err = r.RowsAffected()
	if err != nil {
		logx.WithContext(tx.Context()).Errorf("rowsAffected in Insert(%v)_error: %v", data, err)
	}

	return
}
