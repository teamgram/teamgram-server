/*
 * WARNING! All changes made in this file will be lost!
 *   Created from by 'dalgen'
 *
 * Copyright (c) 2022-present,  Teamgram Authors.
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package mysql_dao

import (
	"context"
	"database/sql"

	"github.com/teamgram/marmota/pkg/stores/sqlx"
	"github.com/teamgram/teamgram-server/app/service/biz/user/internal/dal/dataobject"

	"github.com/zeromicro/go-zero/core/logx"
)

var _ *sql.Result

type UnregisteredContactsDAO struct {
	db *sqlx.DB
}

func NewUnregisteredContactsDAO(db *sqlx.DB) *UnregisteredContactsDAO {
	return &UnregisteredContactsDAO{db}
}

// InsertOrUpdate
// insert into unregistered_contacts(phone, importer_user_id, import_first_name, import_last_name) values (:phone, :importer_user_id, :import_first_name, :import_last_name) on duplicate key update import_first_name = values(import_first_name), import_last_name = values(import_last_name)
// TODO(@benqi): sqlmap
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
// TODO(@benqi): sqlmap
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
// TODO(@benqi): sqlmap
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
// TODO(@benqi): sqlmap
func (dao *UnregisteredContactsDAO) SelectImportersByPhoneWithCB(ctx context.Context, phone string, cb func(i int, v *dataobject.UnregisteredContactsDO)) (rList []dataobject.UnregisteredContactsDO, err error) {
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
		for i := 0; i < len(rList); i++ {
			cb(i, &rList[i])
		}
	}

	return
}

// UpdateContactName
// update unregistered_contacts set import_first_name = :import_first_name, import_last_name = :import_last_name where id = :id
// TODO(@benqi): sqlmap
func (dao *UnregisteredContactsDAO) UpdateContactName(ctx context.Context, import_first_name string, import_last_name string, id int64) (rowsAffected int64, err error) {
	var (
		query   = "update unregistered_contacts set import_first_name = ?, import_last_name = ? where id = ?"
		rResult sql.Result
	)
	rResult, err = dao.db.Exec(ctx, query, import_first_name, import_last_name, id)

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

// update unregistered_contacts set import_first_name = :import_first_name, import_last_name = :import_last_name where id = :id
// UpdateContactNameTx
// TODO(@benqi): sqlmap
func (dao *UnregisteredContactsDAO) UpdateContactNameTx(tx *sqlx.Tx, import_first_name string, import_last_name string, id int64) (rowsAffected int64, err error) {
	var (
		query   = "update unregistered_contacts set import_first_name = ?, import_last_name = ? where id = ?"
		rResult sql.Result
	)
	rResult, err = tx.Exec(query, import_first_name, import_last_name, id)

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
// TODO(@benqi): sqlmap
func (dao *UnregisteredContactsDAO) DeleteContacts(ctx context.Context, id_list []int64) (rowsAffected int64, err error) {
	var (
		query   = "update unregistered_contacts set imported = 1 where id in (?)"
		a       []interface{}
		rResult sql.Result
	)

	if len(id_list) == 0 {
		return
	}

	query, a, err = sqlx.In(query, id_list)
	if err != nil {
		// r sql.Result
		logx.WithContext(ctx).Errorf("sqlx.In in DeleteContacts(_), error: %v", err)
		return
	}
	rResult, err = dao.db.Exec(ctx, query, a...)

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

// update unregistered_contacts set imported = 1 where id in (:id_list)
// DeleteContactsTx
// TODO(@benqi): sqlmap
func (dao *UnregisteredContactsDAO) DeleteContactsTx(tx *sqlx.Tx, id_list []int64) (rowsAffected int64, err error) {
	var (
		query   = "update unregistered_contacts set imported = 1 where id in (?)"
		a       []interface{}
		rResult sql.Result
	)

	if len(id_list) == 0 {
		return
	}

	query, a, err = sqlx.In(query, id_list)
	if err != nil {
		// r sql.Result
		logx.WithContext(tx.Context()).Errorf("sqlx.In in DeleteContacts(_), error: %v", err)
		return
	}
	rResult, err = tx.Exec(query, a...)

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
// TODO(@benqi): sqlmap
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

// update unregistered_contacts set imported = 1 where phone = :phone
// DeleteImportersByPhoneTx
// TODO(@benqi): sqlmap
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
