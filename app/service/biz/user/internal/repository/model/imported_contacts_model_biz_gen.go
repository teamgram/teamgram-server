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
var _ *sqlx.Tx

type bizImportedContactsModel interface {
	InsertOrUpdate(ctx context.Context, data *ImportedContacts) (lastInsertId, rowsAffected int64, err error)
	SelectList(ctx context.Context, userId int64) ([]ImportedContacts, error)
	SelectListWithCB(ctx context.Context, userId int64, cb func(sz, i int, v *ImportedContacts)) ([]ImportedContacts, error)
	SelectListByImportedList(ctx context.Context, userId int64, idList []int64) ([]ImportedContacts, error)
	SelectListByImportedListWithCB(ctx context.Context, userId int64, idList []int64, cb func(sz, i int, v *ImportedContacts)) ([]ImportedContacts, error)
	SelectAllList(ctx context.Context, userId int64) ([]ImportedContacts, error)
	SelectAllListWithCB(ctx context.Context, userId int64, cb func(sz, i int, v *ImportedContacts)) ([]ImportedContacts, error)
	Delete(ctx context.Context, userId int64, importedUserId int64) (rowsAffected int64, err error)
}

type ImportedContactsTxModel interface {
	InsertOrUpdate(data *ImportedContacts) (lastInsertId, rowsAffected int64, err error)
	SelectList(userId int64) ([]ImportedContacts, error)
	SelectListByImportedList(userId int64, idList []int64) ([]ImportedContacts, error)
	SelectAllList(userId int64) ([]ImportedContacts, error)
	Delete(userId int64, importedUserId int64) (rowsAffected int64, err error)
}

type defaultImportedContactsTxModel struct {
	tx *sqlx.Tx
}

func NewImportedContactsTxModel(tx *sqlx.Tx) ImportedContactsTxModel {
	return &defaultImportedContactsTxModel{tx: tx}
}

// InsertOrUpdate
// insert into imported_contacts(user_id, imported_user_id) values (:user_id, :imported_user_id) on duplicate key update deleted = 0
func (m *defaultImportedContactsModel) InsertOrUpdate(ctx context.Context, data *ImportedContacts) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into imported_contacts(user_id, imported_user_id) values (:user_id, :imported_user_id) on duplicate key update deleted = 0"
		r     sql.Result
	)

	r, err = m.db.NamedExec(ctx, query, data)
	if err != nil {
		err = fmt.Errorf("imported_contacts.InsertOrUpdate named exec: %w", err)
		return
	}

	lastInsertId, err = r.LastInsertId()
	if err != nil {
		err = fmt.Errorf("imported_contacts.InsertOrUpdate last insert id: %w", err)
		return
	}
	rowsAffected, err = r.RowsAffected()
	if err != nil {
		err = fmt.Errorf("imported_contacts.InsertOrUpdate rows affected: %w", err)
	}

	return

}

// InsertOrUpdate
// insert into imported_contacts(user_id, imported_user_id) values (:user_id, :imported_user_id) on duplicate key update deleted = 0
func (m *defaultImportedContactsTxModel) InsertOrUpdate(data *ImportedContacts) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into imported_contacts(user_id, imported_user_id) values (:user_id, :imported_user_id) on duplicate key update deleted = 0"
		r     sql.Result
	)

	r, err = m.tx.NamedExec(query, data)
	if err != nil {
		err = fmt.Errorf("imported_contacts.InsertOrUpdate named exec: %w", err)
		return
	}

	lastInsertId, err = r.LastInsertId()
	if err != nil {
		err = fmt.Errorf("imported_contacts.InsertOrUpdate last insert id: %w", err)
		return
	}
	rowsAffected, err = r.RowsAffected()
	if err != nil {
		err = fmt.Errorf("imported_contacts.InsertOrUpdate rows affected: %w", err)
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
		if errors.Is(err, sqlx.ErrNotFound) {
			rList = []ImportedContacts{}
			err = nil
			return
		}
		err = fmt.Errorf("imported_contacts.SelectList: %w", err)
		return
	}

	rList = values

	return
}

// SelectList
// select id, user_id, imported_user_id from imported_contacts where user_id = :user_id and deleted = 0
func (m *defaultImportedContactsTxModel) SelectList(userId int64) (rList []ImportedContacts, err error) {
	var (
		query  = "select id, user_id, imported_user_id from imported_contacts where user_id = ? and deleted = 0"
		values []ImportedContacts
	)
	err = m.tx.QueryRowsPartial(&values, query, userId)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			rList = []ImportedContacts{}
			err = nil
			return
		}
		err = fmt.Errorf("imported_contacts.SelectList: %w", err)
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
		if errors.Is(err, sqlx.ErrNotFound) {
			rList = []ImportedContacts{}
			err = nil
			return
		}
		err = fmt.Errorf("imported_contacts.SelectListWithCB: %w", err)
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
		if errors.Is(err, sqlx.ErrNotFound) {
			rList = []ImportedContacts{}
			err = nil
			return
		}
		err = fmt.Errorf("imported_contacts.SelectListByImportedList: %w", err)
		return
	}

	rList = values

	return
}

// SelectListByImportedList
// select id, user_id, imported_user_id from imported_contacts where user_id = :user_id and deleted = 0 and imported_user_id in (:idList)
func (m *defaultImportedContactsTxModel) SelectListByImportedList(userId int64, idList []int64) (rList []ImportedContacts, err error) {
	var (
		query  = fmt.Sprintf("select id, user_id, imported_user_id from imported_contacts where user_id = ? and deleted = 0 and imported_user_id in (%s)", sqlx.InInt64List(idList))
		values []ImportedContacts
	)
	if len(idList) == 0 {
		rList = []ImportedContacts{}
		return
	}

	err = m.tx.QueryRowsPartial(&values, query, userId)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			rList = []ImportedContacts{}
			err = nil
			return
		}
		err = fmt.Errorf("imported_contacts.SelectListByImportedList: %w", err)
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
		if errors.Is(err, sqlx.ErrNotFound) {
			rList = []ImportedContacts{}
			err = nil
			return
		}
		err = fmt.Errorf("imported_contacts.SelectListByImportedListWithCB: %w", err)
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
		if errors.Is(err, sqlx.ErrNotFound) {
			rList = []ImportedContacts{}
			err = nil
			return
		}
		err = fmt.Errorf("imported_contacts.SelectAllList: %w", err)
		return
	}

	rList = values

	return
}

// SelectAllList
// select id, user_id, imported_user_id from imported_contacts where user_id = :user_id
func (m *defaultImportedContactsTxModel) SelectAllList(userId int64) (rList []ImportedContacts, err error) {
	var (
		query  = "select id, user_id, imported_user_id from imported_contacts where user_id = ?"
		values []ImportedContacts
	)
	err = m.tx.QueryRowsPartial(&values, query, userId)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			rList = []ImportedContacts{}
			err = nil
			return
		}
		err = fmt.Errorf("imported_contacts.SelectAllList: %w", err)
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
		if errors.Is(err, sqlx.ErrNotFound) {
			rList = []ImportedContacts{}
			err = nil
			return
		}
		err = fmt.Errorf("imported_contacts.SelectAllListWithCB: %w", err)
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
		err = fmt.Errorf("imported_contacts.Delete exec: %w", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		err = fmt.Errorf("imported_contacts.Delete rows affected: %w", err)
		return
	}

	return
}

// Delete
// update imported_contacts set deleted = 1 where user_id = :user_id and imported_user_id = :imported_user_id
func (m *defaultImportedContactsTxModel) Delete(userId int64, importedUserId int64) (rowsAffected int64, err error) {
	var (
		query   = "update imported_contacts set deleted = 1 where user_id = ? and imported_user_id = ?"
		rResult sql.Result
	)
	rResult, err = m.tx.Exec(query, userId, importedUserId)

	if err != nil {
		err = fmt.Errorf("imported_contacts.Delete exec: %w", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		err = fmt.Errorf("imported_contacts.Delete rows affected: %w", err)
		return
	}

	return
}
