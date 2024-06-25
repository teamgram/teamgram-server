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

type UnregisteredContactsDAO struct {
	db *sqlx.DB
}

func NewUnregisteredContactsDAO(db *sqlx.DB) *UnregisteredContactsDAO {
	return &UnregisteredContactsDAO{
		db: db,
	}
}

// InsertOrUpdate
// insert into unregistered_contacts(phone, importer_user_id, import_first_name, import_last_name) values (:phone, :importer_user_id, :import_first_name, :import_last_name) on duplicate key update import_first_name = values(import_first_name), import_last_name = values(import_last_name)
func (dao *UnregisteredContactsDAO) InsertOrUpdate(ctx context.Context, do *dataobject.UnregisteredContactsDO) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into unregistered_contacts(phone, importer_user_id, import_first_name, import_last_name) values (:phone, :importer_user_id, :import_first_name, :import_last_name) on duplicate key update import_first_name = values(import_first_name), import_last_name = values(import_last_name)"
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
// insert into unregistered_contacts(phone, importer_user_id, import_first_name, import_last_name) values (:phone, :importer_user_id, :import_first_name, :import_last_name) on duplicate key update import_first_name = values(import_first_name), import_last_name = values(import_last_name)
func (dao *UnregisteredContactsDAO) InsertOrUpdateTx(tx *sqlx.Tx, do *dataobject.UnregisteredContactsDO) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into unregistered_contacts(phone, importer_user_id, import_first_name, import_last_name) values (:phone, :importer_user_id, :import_first_name, :import_last_name) on duplicate key update import_first_name = values(import_first_name), import_last_name = values(import_last_name)"
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

// SelectImportersByPhone
// select id, importer_user_id, phone, import_first_name, import_last_name from unregistered_contacts where phone = :phone and imported = 0
func (dao *UnregisteredContactsDAO) SelectImportersByPhone(ctx context.Context, phone string) (rList []dataobject.UnregisteredContactsDO, err error) {
	var (
		query  = "select id, importer_user_id, phone, import_first_name, import_last_name from unregistered_contacts where phone = ? and imported = 0"
		values []dataobject.UnregisteredContactsDO
	)
	err = dao.db.QueryRowsPartial(ctx, &values, query, phone)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectImportersByPhone(_), error: %v", err)
		return
	}

	rList = values

	return
}

// SelectImportersByPhoneWithCB
// select id, importer_user_id, phone, import_first_name, import_last_name from unregistered_contacts where phone = :phone and imported = 0
func (dao *UnregisteredContactsDAO) SelectImportersByPhoneWithCB(ctx context.Context, phone string, cb func(sz, i int, v *dataobject.UnregisteredContactsDO)) (rList []dataobject.UnregisteredContactsDO, err error) {
	var (
		query  = "select id, importer_user_id, phone, import_first_name, import_last_name from unregistered_contacts where phone = ? and imported = 0"
		values []dataobject.UnregisteredContactsDO
	)
	err = dao.db.QueryRowsPartial(ctx, &values, query, phone)

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
func (dao *UnregisteredContactsDAO) UpdateContactName(ctx context.Context, importFirstName string, importLastName string, id int64) (rowsAffected int64, err error) {
	var (
		query   = "update unregistered_contacts set import_first_name = ?, import_last_name = ? where id = ?"
		rResult sql.Result
	)

	rResult, err = dao.db.Exec(ctx, query, importFirstName, importLastName, id)

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
func (dao *UnregisteredContactsDAO) UpdateContactNameTx(tx *sqlx.Tx, importFirstName string, importLastName string, id int64) (rowsAffected int64, err error) {
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
func (dao *UnregisteredContactsDAO) DeleteContacts(ctx context.Context, idList []int64) (rowsAffected int64, err error) {
	var (
		query   = fmt.Sprintf("update unregistered_contacts set imported = 1 where id in (%s)", sqlx.InInt64List(idList))
		rResult sql.Result
	)

	if len(idList) == 0 {
		return
	}

	rResult, err = dao.db.Exec(ctx, query)

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
func (dao *UnregisteredContactsDAO) DeleteContactsTx(tx *sqlx.Tx, idList []int64) (rowsAffected int64, err error) {
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
func (dao *UnregisteredContactsDAO) DeleteImportersByPhone(ctx context.Context, phone string) (rowsAffected int64, err error) {
	var (
		query   = "update unregistered_contacts set imported = 1 where phone = ?"
		rResult sql.Result
	)

	rResult, err = dao.db.Exec(ctx, query, phone)

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
func (dao *UnregisteredContactsDAO) DeleteImportersByPhoneTx(tx *sqlx.Tx, phone string) (rowsAffected int64, err error) {
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
