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
	bizPhoneBooksModel interface {
		InsertOrUpdate(ctx context.Context, data *PhoneBooks) (lastInsertId, rowsAffected int64, err error)
		InsertOrUpdateTx(tx *sqlx.Tx, data *PhoneBooks) (lastInsertId, rowsAffected int64, err error)
	}
)

// InsertOrUpdate
// insert into phone_books(auth_key_id, client_id, phone, first_name, last_name) values (:auth_key_id, :client_id, :phone, :first_name, :last_name) on duplicate key update phone = values(phone), first_name = values(first_name), last_name = values(last_name)
func (m *defaultPhoneBooksModel) InsertOrUpdate(ctx context.Context, data *PhoneBooks) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into phone_books(auth_key_id, client_id, phone, first_name, last_name) values (:auth_key_id, :client_id, :phone, :first_name, :last_name) on duplicate key update phone = values(phone), first_name = values(first_name), last_name = values(last_name)"
		r     sql.Result
	)

	r, err = m.db.NamedExec(ctx, query, data)
	if err != nil {
		err = fmt.Errorf("phone_books.InsertOrUpdate named exec: %w", err)
		return
	}

	lastInsertId, err = r.LastInsertId()
	if err != nil {
		err = fmt.Errorf("phone_books.InsertOrUpdate last insert id: %w", err)
		return
	}
	rowsAffected, err = r.RowsAffected()
	if err != nil {
		err = fmt.Errorf("phone_books.InsertOrUpdate rows affected: %w", err)
	}

	return

}

// InsertOrUpdateTx
// insert into phone_books(auth_key_id, client_id, phone, first_name, last_name) values (:auth_key_id, :client_id, :phone, :first_name, :last_name) on duplicate key update phone = values(phone), first_name = values(first_name), last_name = values(last_name)
func (m *defaultPhoneBooksModel) InsertOrUpdateTx(tx *sqlx.Tx, data *PhoneBooks) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into phone_books(auth_key_id, client_id, phone, first_name, last_name) values (:auth_key_id, :client_id, :phone, :first_name, :last_name) on duplicate key update phone = values(phone), first_name = values(first_name), last_name = values(last_name)"
		r     sql.Result
	)

	r, err = tx.NamedExec(query, data)
	if err != nil {
		err = fmt.Errorf("phone_books.InsertOrUpdateTx named exec: %w", err)
		return
	}

	lastInsertId, err = r.LastInsertId()
	if err != nil {
		err = fmt.Errorf("phone_books.InsertOrUpdateTx last insert id: %w", err)
		return
	}
	rowsAffected, err = r.RowsAffected()
	if err != nil {
		err = fmt.Errorf("phone_books.InsertOrUpdateTx rows affected: %w", err)
	}

	return
}
