/*
 * WARNING! All changes made in this file will be lost!
 *   Created from by 'dalgen'
 *
 * Copyright (c) 2026 The Teamgram Authors.
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
	bizUnregisteredContactsModel interface {
		InsertOrUpdate(ctx context.Context, data *UnregisteredContacts) (lastInsertId, rowsAffected int64, err error)
		InsertOrUpdateTx(tx *sqlx.Tx, data *UnregisteredContacts) (lastInsertId, rowsAffected int64, err error)

		SelectImportersByPhone(ctx context.Context, phone string) ([]UnregisteredContacts, error)
		SelectImportersByPhoneWithCB(ctx context.Context, phone string, cb func(sz, i int, v *UnregisteredContacts)) ([]UnregisteredContacts, error)

		UpdateContactName(ctx context.Context, importFirstName string, importLastName string, id int64) (rowsAffected int64, err error)
		UpdateContactNameTx(tx *sqlx.Tx, importFirstName string, importLastName string, id int64) (rowsAffected int64, err error)

		DeleteContacts(ctx context.Context, idList []int64) (rowsAffected int64, err error)
		DeleteContactsTx(tx *sqlx.Tx, idList []int64) (rowsAffected int64, err error)

		DeleteImportersByPhone(ctx context.Context, phone string) (rowsAffected int64, err error)
		DeleteImportersByPhoneTx(tx *sqlx.Tx, phone string) (rowsAffected int64, err error)
	}
)

// InsertOrUpdate
// insert into unregistered_contacts(phone, importer_user_id, import_first_name, import_last_name) values (:phone, :importer_user_id, :import_first_name, :import_last_name) on duplicate key update import_first_name = values(import_first_name), import_last_name = values(import_last_name)
func (m *defaultUnregisteredContactsModel) InsertOrUpdate(ctx context.Context, data *UnregisteredContacts) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into unregistered_contacts(phone, importer_user_id, import_first_name, import_last_name) values (:phone, :importer_user_id, :import_first_name, :import_last_name) on duplicate key update import_first_name = values(import_first_name), import_last_name = values(import_last_name)"
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
// insert into unregistered_contacts(phone, importer_user_id, import_first_name, import_last_name) values (:phone, :importer_user_id, :import_first_name, :import_last_name) on duplicate key update import_first_name = values(import_first_name), import_last_name = values(import_last_name)
func (m *defaultUnregisteredContactsModel) InsertOrUpdateTx(tx *sqlx.Tx, data *UnregisteredContacts) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into unregistered_contacts(phone, importer_user_id, import_first_name, import_last_name) values (:phone, :importer_user_id, :import_first_name, :import_last_name) on duplicate key update import_first_name = values(import_first_name), import_last_name = values(import_last_name)"
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

// SelectImportersByPhone
// select id, importer_user_id, phone, import_first_name, import_last_name from unregistered_contacts where phone = :phone and imported = 0
func (m *defaultUnregisteredContactsModel) SelectImportersByPhone(ctx context.Context, phone string) (rList []UnregisteredContacts, err error) {
	var (
		query  = "select id, importer_user_id, phone, import_first_name, import_last_name from unregistered_contacts where phone = ? and imported = 0"
		values []UnregisteredContacts
	)
	err = m.db.QueryRowsPartial(ctx, &values, query, phone)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectImportersByPhone(_), error: %v", err)
		return
	}

	rList = values

	return
}

// SelectImportersByPhoneWithCB
// select id, importer_user_id, phone, import_first_name, import_last_name from unregistered_contacts where phone = :phone and imported = 0
func (m *defaultUnregisteredContactsModel) SelectImportersByPhoneWithCB(ctx context.Context, phone string, cb func(sz, i int, v *UnregisteredContacts)) (rList []UnregisteredContacts, err error) {
	var (
		query  = "select id, importer_user_id, phone, import_first_name, import_last_name from unregistered_contacts where phone = ? and imported = 0"
		values []UnregisteredContacts
	)
	err = m.db.QueryRowsPartial(ctx, &values, query, phone)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectImportersByPhone(_), error: %v", err)
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

// UpdateContactName
// update unregistered_contacts set import_first_name = :import_first_name, import_last_name = :import_last_name where id = :id
func (m *defaultUnregisteredContactsModel) UpdateContactName(ctx context.Context, importFirstName string, importLastName string, id int64) (rowsAffected int64, err error) {
	var (
		query   = "update unregistered_contacts set import_first_name = ?, import_last_name = ? where id = ?"
		rResult sql.Result
	)

	rResult, err = m.db.Exec(ctx, query, importFirstName, importLastName, id)

	if err != nil {
		logx.WithContext(ctx).Errorf("exec in UpdateContactName(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		logx.WithContext(ctx).Errorf("rowsAffected in UpdateContactName(_), error: %v", err)
	}

	return
}

// UpdateContactNameTx
// update unregistered_contacts set import_first_name = :import_first_name, import_last_name = :import_last_name where id = :id
func (m *defaultUnregisteredContactsModel) UpdateContactNameTx(tx *sqlx.Tx, importFirstName string, importLastName string, id int64) (rowsAffected int64, err error) {
	var (
		query   = "update unregistered_contacts set import_first_name = ?, import_last_name = ? where id = ?"
		rResult sql.Result
	)
	rResult, err = tx.Exec(query, importFirstName, importLastName, id)

	if err != nil {
		logx.WithContext(tx.Context()).Errorf("exec in UpdateContactName(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		logx.WithContext(tx.Context()).Errorf("rowsAffected in UpdateContactName(_), error: %v", err)
	}

	return
}

// DeleteContacts
// update unregistered_contacts set imported = 1 where id in (:id_list)
func (m *defaultUnregisteredContactsModel) DeleteContacts(ctx context.Context, idList []int64) (rowsAffected int64, err error) {
	var (
		query   = fmt.Sprintf("update unregistered_contacts set imported = 1 where id in (%s)", sqlx.InInt64List(idList))
		rResult sql.Result
	)

	if len(idList) == 0 {
		return
	}

	rResult, err = m.db.Exec(ctx, query)

	if err != nil {
		logx.WithContext(ctx).Errorf("exec in DeleteContacts(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		logx.WithContext(ctx).Errorf("rowsAffected in DeleteContacts(_), error: %v", err)
	}

	return
}

// DeleteContactsTx
// update unregistered_contacts set imported = 1 where id in (:id_list)
func (m *defaultUnregisteredContactsModel) DeleteContactsTx(tx *sqlx.Tx, idList []int64) (rowsAffected int64, err error) {
	var (
		query   = fmt.Sprintf("update unregistered_contacts set imported = 1 where id in (%s)", sqlx.InInt64List(idList))
		rResult sql.Result
	)

	if len(idList) == 0 {
		return
	}

	rResult, err = tx.Exec(query)

	if err != nil {
		logx.WithContext(tx.Context()).Errorf("exec in DeleteContacts(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		logx.WithContext(tx.Context()).Errorf("rowsAffected in DeleteContacts(_), error: %v", err)
	}

	return
}

// DeleteImportersByPhone
// update unregistered_contacts set imported = 1 where phone = :phone
func (m *defaultUnregisteredContactsModel) DeleteImportersByPhone(ctx context.Context, phone string) (rowsAffected int64, err error) {
	var (
		query   = "update unregistered_contacts set imported = 1 where phone = ?"
		rResult sql.Result
	)

	rResult, err = m.db.Exec(ctx, query, phone)

	if err != nil {
		logx.WithContext(ctx).Errorf("exec in DeleteImportersByPhone(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		logx.WithContext(ctx).Errorf("rowsAffected in DeleteImportersByPhone(_), error: %v", err)
	}

	return
}

// DeleteImportersByPhoneTx
// update unregistered_contacts set imported = 1 where phone = :phone
func (m *defaultUnregisteredContactsModel) DeleteImportersByPhoneTx(tx *sqlx.Tx, phone string) (rowsAffected int64, err error) {
	var (
		query   = "update unregistered_contacts set imported = 1 where phone = ?"
		rResult sql.Result
	)
	rResult, err = tx.Exec(query, phone)

	if err != nil {
		logx.WithContext(tx.Context()).Errorf("exec in DeleteImportersByPhone(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		logx.WithContext(tx.Context()).Errorf("rowsAffected in DeleteImportersByPhone(_), error: %v", err)
	}

	return
}
