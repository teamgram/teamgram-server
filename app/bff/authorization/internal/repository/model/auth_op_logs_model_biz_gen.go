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
)

var _ *sql.Result
var _ = fmt.Sprintf
var _ = strings.Join
var _ = errors.Is
var _ *sqlx.DB

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
		err = fmt.Errorf("auth_op_logs.Insert named exec: %w", err)
		return
	}

	lastInsertId, err = r.LastInsertId()
	if err != nil {
		err = fmt.Errorf("auth_op_logs.Insert last insert id: %w", err)
		return
	}
	rowsAffected, err = r.RowsAffected()
	if err != nil {
		err = fmt.Errorf("auth_op_logs.Insert rows affected: %w", err)
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
		err = fmt.Errorf("auth_op_logs.InsertTx named exec: %w", err)
		return
	}

	lastInsertId, err = r.LastInsertId()
	if err != nil {
		err = fmt.Errorf("auth_op_logs.InsertTx last insert id: %w", err)
		return
	}
	rowsAffected, err = r.RowsAffected()
	if err != nil {
		err = fmt.Errorf("auth_op_logs.InsertTx rows affected: %w", err)
	}

	return
}
