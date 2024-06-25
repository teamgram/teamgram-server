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
	"errors"
	"fmt"
	"strings"

	"github.com/teamgram/marmota/pkg/stores/sqlx"
	"github.com/teamgram/teamgram-server/app/service/biz/user/internal/dal/dataobject"

	"github.com/zeromicro/go-zero/core/logx"
)

var _ *sql.Result
var _ = fmt.Sprintf
var _ = strings.Join
var _ = errors.Is

type PhoneBooksDAO struct {
	db *sqlx.DB
}

func NewPhoneBooksDAO(db *sqlx.DB) *PhoneBooksDAO {
	return &PhoneBooksDAO{
		db: db,
	}
}

// InsertOrUpdate
// insert into phone_books(auth_key_id, client_id, phone, first_name, last_name) values (:auth_key_id, :client_id, :phone, :first_name, :last_name) on duplicate key update phone = values(phone), first_name = values(first_name), last_name = values(last_name)
func (dao *PhoneBooksDAO) InsertOrUpdate(ctx context.Context, do *dataobject.PhoneBooksDO) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into phone_books(auth_key_id, client_id, phone, first_name, last_name) values (:auth_key_id, :client_id, :phone, :first_name, :last_name) on duplicate key update phone = values(phone), first_name = values(first_name), last_name = values(last_name)"
		r     sql.Result
	)

	r, err = dao.db.NamedExec(ctx, query, do)
	if err != nil {
		logx.WithContext(ctx).Errorf("namedExec in InsertOrUpdate(%v), error: %v", do, err)
		return
	}

	lastInsertId, err = r.LastInsertId()
	if err != nil {
		logx.WithContext(ctx).Errorf("lastInsertId in InsertOrUpdate(%v)_error: %v", do, err)
		return
	}
	rowsAffected, err = r.RowsAffected()
	if err != nil {
		logx.WithContext(ctx).Errorf("rowsAffected in InsertOrUpdate(%v)_error: %v", do, err)
	}

	return
}

// InsertOrUpdateTx
// insert into phone_books(auth_key_id, client_id, phone, first_name, last_name) values (:auth_key_id, :client_id, :phone, :first_name, :last_name) on duplicate key update phone = values(phone), first_name = values(first_name), last_name = values(last_name)
func (dao *PhoneBooksDAO) InsertOrUpdateTx(tx *sqlx.Tx, do *dataobject.PhoneBooksDO) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into phone_books(auth_key_id, client_id, phone, first_name, last_name) values (:auth_key_id, :client_id, :phone, :first_name, :last_name) on duplicate key update phone = values(phone), first_name = values(first_name), last_name = values(last_name)"
		r     sql.Result
	)

	r, err = tx.NamedExec(query, do)
	if err != nil {
		logx.WithContext(tx.Context()).Errorf("namedExec in InsertOrUpdate(%v), error: %v", do, err)
		return
	}

	lastInsertId, err = r.LastInsertId()
	if err != nil {
		logx.WithContext(tx.Context()).Errorf("lastInsertId in InsertOrUpdate(%v)_error: %v", do, err)
		return
	}
	rowsAffected, err = r.RowsAffected()
	if err != nil {
		logx.WithContext(tx.Context()).Errorf("rowsAffected in InsertOrUpdate(%v)_error: %v", do, err)
	}

	return
}
