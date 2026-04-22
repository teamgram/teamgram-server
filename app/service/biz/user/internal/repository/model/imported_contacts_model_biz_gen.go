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
	bizImportedContactsModel interface {
		InsertOrUpdate(ctx context.Context, data *ImportedContacts) (lastInsertId, rowsAffected int64, err error)
		InsertOrUpdateTx(tx *sqlx.Tx, data *ImportedContacts) (lastInsertId, rowsAffected int64, err error)

		SelectList(ctx context.Context, userId int64) ([]ImportedContacts, error)
		SelectListWithCB(ctx context.Context, userId int64, cb func(sz, i int, v *ImportedContacts)) ([]ImportedContacts, error)

		SelectListByImportedList(ctx context.Context, userId int64, idList []int64) ([]ImportedContacts, error)
		SelectListByImportedListWithCB(ctx context.Context, userId int64, idList []int64, cb func(sz, i int, v *ImportedContacts)) ([]ImportedContacts, error)

		SelectAllList(ctx context.Context, userId int64) ([]ImportedContacts, error)
		SelectAllListWithCB(ctx context.Context, userId int64, cb func(sz, i int, v *ImportedContacts)) ([]ImportedContacts, error)

		Delete(ctx context.Context, userId int64, importedUserId int64) (rowsAffected int64, err error)
		DeleteTx(tx *sqlx.Tx, userId int64, importedUserId int64) (rowsAffected int64, err error)
	}
)

// InsertOrUpdate
// insert into imported_contacts(user_id, imported_user_id) values (:user_id, :imported_user_id) on duplicate key update deleted = 0
func (m *defaultImportedContactsModel) InsertOrUpdate(ctx context.Context, data *ImportedContacts) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into imported_contacts(user_id, imported_user_id) values (:user_id, :imported_user_id) on duplicate key update deleted = 0"
		r     sql.Result
	)

	r, err = m.db.NamedExec(ctx, query, data)
	if err != nil {
		logx.WithContext(ctx).Errorf("namedExec in InsertOrUpdate(%v), error: %v", data, err)
		return
	}

	lastInsertId, err = r.LastInsertId()
	if err != nil {
		logx.WithContext(ctx).Errorf("lastInsertId in InsertOrUpdate(%v)_error: %v", data, err)
		return
	}
	rowsAffected, err = r.RowsAffected()
	if err != nil {
		logx.WithContext(ctx).Errorf("rowsAffected in InsertOrUpdate(%v)_error: %v", data, err)
	}

	return
}

// InsertOrUpdateTx
// insert into imported_contacts(user_id, imported_user_id) values (:user_id, :imported_user_id) on duplicate key update deleted = 0
func (m *defaultImportedContactsModel) InsertOrUpdateTx(tx *sqlx.Tx, data *ImportedContacts) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into imported_contacts(user_id, imported_user_id) values (:user_id, :imported_user_id) on duplicate key update deleted = 0"
		r     sql.Result
	)

	r, err = tx.NamedExec(query, data)
	if err != nil {
		logx.WithContext(tx.Context()).Errorf("namedExec in InsertOrUpdate(%v), error: %v", data, err)
		return
	}

	lastInsertId, err = r.LastInsertId()
	if err != nil {
		logx.WithContext(tx.Context()).Errorf("lastInsertId in InsertOrUpdate(%v)_error: %v", data, err)
		return
	}
	rowsAffected, err = r.RowsAffected()
	if err != nil {
		logx.WithContext(tx.Context()).Errorf("rowsAffected in InsertOrUpdate(%v)_error: %v", data, err)
	}

	return
}

// SelectList
// select id, user_id, imported_user_id from imported_contacts where user_id = :user_id and deleted = 0
func (m *defaultImportedContactsModel) SelectList(ctx context.Context, userId int64) (rList []ImportedContacts, err error) {
	var (
		query  = "select id, user_id, imported_user_id from imported_contacts where user_id = ? and deleted = 0"
		values []ImportedContacts
	)
	err = m.db.QueryRowsPartial(ctx, &values, query, userId)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectList(_), error: %v", err)
		return
	}

	rList = values

	return
}

// SelectListWithCB
// select id, user_id, imported_user_id from imported_contacts where user_id = :user_id and deleted = 0
func (m *defaultImportedContactsModel) SelectListWithCB(ctx context.Context, userId int64, cb func(sz, i int, v *ImportedContacts)) (rList []ImportedContacts, err error) {
	var (
		query  = "select id, user_id, imported_user_id from imported_contacts where user_id = ? and deleted = 0"
		values []ImportedContacts
	)
	err = m.db.QueryRowsPartial(ctx, &values, query, userId)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectList(_), error: %v", err)
		return
	}

	rList = values

	if cb != nil {
		sz := len(rList)
		for i := 0; i < sz; i++ {
			cb(sz, i, &rList[i])
		}
	}

	return
}

// SelectListByImportedList
// select id, user_id, imported_user_id from imported_contacts where user_id = :user_id and deleted = 0 and imported_user_id in (:idList)
func (m *defaultImportedContactsModel) SelectListByImportedList(ctx context.Context, userId int64, idList []int64) (rList []ImportedContacts, err error) {
	var (
		query  = fmt.Sprintf("select id, user_id, imported_user_id from imported_contacts where user_id = ? and deleted = 0 and imported_user_id in (%s)", sqlx.InInt64List(idList))
		values []ImportedContacts
	)
	if len(idList) == 0 {
		rList = []ImportedContacts{}
		return
	}

	err = m.db.QueryRowsPartial(ctx, &values, query, userId)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectListByImportedList(_), error: %v", err)
		return
	}

	rList = values

	return
}

// SelectListByImportedListWithCB
// select id, user_id, imported_user_id from imported_contacts where user_id = :user_id and deleted = 0 and imported_user_id in (:idList)
func (m *defaultImportedContactsModel) SelectListByImportedListWithCB(ctx context.Context, userId int64, idList []int64, cb func(sz, i int, v *ImportedContacts)) (rList []ImportedContacts, err error) {
	var (
		query  = fmt.Sprintf("select id, user_id, imported_user_id from imported_contacts where user_id = ? and deleted = 0 and imported_user_id in (%s)", sqlx.InInt64List(idList))
		values []ImportedContacts
	)
	if len(idList) == 0 {
		rList = []ImportedContacts{}
		return
	}

	err = m.db.QueryRowsPartial(ctx, &values, query, userId)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectListByImportedList(_), error: %v", err)
		return
	}

	rList = values

	if cb != nil {
		sz := len(rList)
		for i := 0; i < sz; i++ {
			cb(sz, i, &rList[i])
		}
	}

	return
}

// SelectAllList
// select id, user_id, imported_user_id from imported_contacts where user_id = :user_id
func (m *defaultImportedContactsModel) SelectAllList(ctx context.Context, userId int64) (rList []ImportedContacts, err error) {
	var (
		query  = "select id, user_id, imported_user_id from imported_contacts where user_id = ?"
		values []ImportedContacts
	)
	err = m.db.QueryRowsPartial(ctx, &values, query, userId)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectAllList(_), error: %v", err)
		return
	}

	rList = values

	return
}

// SelectAllListWithCB
// select id, user_id, imported_user_id from imported_contacts where user_id = :user_id
func (m *defaultImportedContactsModel) SelectAllListWithCB(ctx context.Context, userId int64, cb func(sz, i int, v *ImportedContacts)) (rList []ImportedContacts, err error) {
	var (
		query  = "select id, user_id, imported_user_id from imported_contacts where user_id = ?"
		values []ImportedContacts
	)
	err = m.db.QueryRowsPartial(ctx, &values, query, userId)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectAllList(_), error: %v", err)
		return
	}

	rList = values

	if cb != nil {
		sz := len(rList)
		for i := 0; i < sz; i++ {
			cb(sz, i, &rList[i])
		}
	}

	return
}

// Delete
// update imported_contacts set deleted = 1 where user_id = :user_id and imported_user_id = :imported_user_id
func (m *defaultImportedContactsModel) Delete(ctx context.Context, userId int64, importedUserId int64) (rowsAffected int64, err error) {
	var (
		query   = "update imported_contacts set deleted = 1 where user_id = ? and imported_user_id = ?"
		rResult sql.Result
	)

	rResult, err = m.db.Exec(ctx, query, userId, importedUserId)

	if err != nil {
		logx.WithContext(ctx).Errorf("exec in Delete(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		logx.WithContext(ctx).Errorf("rowsAffected in Delete(_), error: %v", err)
	}

	return
}

// DeleteTx
// update imported_contacts set deleted = 1 where user_id = :user_id and imported_user_id = :imported_user_id
func (m *defaultImportedContactsModel) DeleteTx(tx *sqlx.Tx, userId int64, importedUserId int64) (rowsAffected int64, err error) {
	var (
		query   = "update imported_contacts set deleted = 1 where user_id = ? and imported_user_id = ?"
		rResult sql.Result
	)
	rResult, err = tx.Exec(query, userId, importedUserId)

	if err != nil {
		logx.WithContext(tx.Context()).Errorf("exec in Delete(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		logx.WithContext(tx.Context()).Errorf("rowsAffected in Delete(_), error: %v", err)
	}

	return
}
