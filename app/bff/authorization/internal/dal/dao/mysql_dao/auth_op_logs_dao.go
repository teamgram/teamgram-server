/*
 * WARNING! All changes made in this file will be lost!
 *   Created from by 'dalgen'
 *
 * Copyright (c) 2024-present,  Teamgram Authors.
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package mysql_dao

import (
	"context"
	"database/sql"

	"github.com/teamgram/marmota/pkg/stores/sqlx"
	"github.com/teamgram/teamgram-server/app/bff/authorization/internal/dal/dataobject"

	"github.com/zeromicro/go-zero/core/logx"
)

type AuthOpLogsDAO struct {
	db *sqlx.DB
}

func NewAuthOpLogsDAO(db *sqlx.DB) *AuthOpLogsDAO {
	return &AuthOpLogsDAO{
		db: db,
	}
}

// Insert
// insert into auth_op_logs(auth_key_id, ip, op_type, log_text) values (:auth_key_id, :ip, :op_type, :log_text)
func (dao *AuthOpLogsDAO) Insert(ctx context.Context, do *dataobject.AuthOpLogsDO) (lastInsertId, rowsAffected int64, err error) {
	var (
		query string
		r     sql.Result
	)
	query = "insert into auth_op_logs(auth_key_id, ip, op_type, log_text) values (:auth_key_id, :ip, :op_type, :log_text)"

	r, err = dao.db.NamedExec(ctx, query, do)
	if err != nil {
		logx.WithContext(ctx).Errorf("namedExec in Insert(%v), error: %v", do, err)
		return
	}

	lastInsertId, err = r.LastInsertId()
	if err != nil {
		logx.WithContext(ctx).Errorf("lastInsertId in Insert(%v), error: %v", do, err)
		return
	}
	rowsAffected, err = r.RowsAffected()
	if err != nil {
		logx.WithContext(ctx).Errorf("rowsAffected in Insert(%v), error: %v", do, err)
	}

	return
}

// InsertTx
// insert into auth_op_logs(auth_key_id, ip, op_type, log_text) values (:auth_key_id, :ip, :op_type, :log_text)
func (dao *AuthOpLogsDAO) InsertTx(tx *sqlx.Tx, do *dataobject.AuthOpLogsDO) (lastInsertId, rowsAffected int64, err error) {
	var (
		query string
		r     sql.Result
	)
	query = "insert into auth_op_logs(auth_key_id, ip, op_type, log_text) values (:auth_key_id, :ip, :op_type, :log_text)"

	r, err = tx.NamedExec(query, do)
	if err != nil {
		logx.WithContext(tx.Context()).Errorf("namedExec in Insert(%v), error: %v", do, err)
		return
	}

	lastInsertId, err = r.LastInsertId()
	if err != nil {
		logx.WithContext(tx.Context()).Errorf("lastInsertId in Insert(%v), error: %v", do, err)
		return
	}
	rowsAffected, err = r.RowsAffected()
	if err != nil {
		logx.WithContext(tx.Context()).Errorf("rowsAffected in Insert(%v), error: %v", do, err)
	}

	return
}
