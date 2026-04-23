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
	bizPopularContactsModel interface {
		InsertOrUpdate(ctx context.Context, data *PopularContacts) (lastInsertId, rowsAffected int64, err error)
		InsertOrUpdateTx(tx *sqlx.Tx, data *PopularContacts) (lastInsertId, rowsAffected int64, err error)

		IncreaseImporters(ctx context.Context, phone string) (rowsAffected int64, err error)
		IncreaseImportersTx(tx *sqlx.Tx, phone string) (rowsAffected int64, err error)

		IncreaseImportersList(ctx context.Context, phoneList []string) (rowsAffected int64, err error)
		IncreaseImportersListTx(tx *sqlx.Tx, phoneList []string) (rowsAffected int64, err error)

		SelectImporters(ctx context.Context, phone string) (*PopularContacts, error)

		SelectImportersList(ctx context.Context, phoneList []string) ([]PopularContacts, error)
		SelectImportersListWithCB(ctx context.Context, phoneList []string, cb func(sz, i int, v *PopularContacts)) ([]PopularContacts, error)
	}
)

// InsertOrUpdate
// insert into popular_contacts(phone, importers, deleted) values (:phone, :importers, 0) on duplicate key update importers = importers + 1
func (m *defaultPopularContactsModel) InsertOrUpdate(ctx context.Context, data *PopularContacts) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into popular_contacts(phone, importers, deleted) values (:phone, :importers, 0) on duplicate key update importers = importers + 1"
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
// insert into popular_contacts(phone, importers, deleted) values (:phone, :importers, 0) on duplicate key update importers = importers + 1
func (m *defaultPopularContactsModel) InsertOrUpdateTx(tx *sqlx.Tx, data *PopularContacts) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into popular_contacts(phone, importers, deleted) values (:phone, :importers, 0) on duplicate key update importers = importers + 1"
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

// IncreaseImporters
// update popular_contacts set importers = importers + 1 where phone = :phone
func (m *defaultPopularContactsModel) IncreaseImporters(ctx context.Context, phone string) (rowsAffected int64, err error) {

	var (
		query   = "update popular_contacts set importers = importers + 1 where phone = ?"
		rResult sql.Result
	)

	rResult, err = m.db.Exec(ctx, query, phone)

	if err != nil {
		logx.WithContext(ctx).Errorf("exec in IncreaseImporters(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		logx.WithContext(ctx).Errorf("rowsAffected in IncreaseImporters(_), error: %v", err)
	}

	return
}

// IncreaseImportersTx
// update popular_contacts set importers = importers + 1 where phone = :phone
func (m *defaultPopularContactsModel) IncreaseImportersTx(tx *sqlx.Tx, phone string) (rowsAffected int64, err error) {
	var (
		query   = "update popular_contacts set importers = importers + 1 where phone = ?"
		rResult sql.Result
	)
	rResult, err = tx.Exec(query, phone)

	if err != nil {
		logx.WithContext(tx.Context()).Errorf("exec in IncreaseImporters(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		logx.WithContext(tx.Context()).Errorf("rowsAffected in IncreaseImporters(_), error: %v", err)
	}

	return
}

// IncreaseImportersList
// update popular_contacts set importers = importers + 1 where phone in (:phoneList)
func (m *defaultPopularContactsModel) IncreaseImportersList(ctx context.Context, phoneList []string) (rowsAffected int64, err error) {

	var (
		query   = fmt.Sprintf("update popular_contacts set importers = importers + 1 where phone in (%s)", sqlx.InStringList(phoneList))
		rResult sql.Result
	)

	if len(phoneList) == 0 {
		return
	}

	rResult, err = m.db.Exec(ctx, query)

	if err != nil {
		logx.WithContext(ctx).Errorf("exec in IncreaseImportersList(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		logx.WithContext(ctx).Errorf("rowsAffected in IncreaseImportersList(_), error: %v", err)
	}

	return
}

// IncreaseImportersListTx
// update popular_contacts set importers = importers + 1 where phone in (:phoneList)
func (m *defaultPopularContactsModel) IncreaseImportersListTx(tx *sqlx.Tx, phoneList []string) (rowsAffected int64, err error) {
	var (
		query   = fmt.Sprintf("update popular_contacts set importers = importers + 1 where phone in (%s)", sqlx.InStringList(phoneList))
		rResult sql.Result
	)

	if len(phoneList) == 0 {
		return
	}

	rResult, err = tx.Exec(query)

	if err != nil {
		logx.WithContext(tx.Context()).Errorf("exec in IncreaseImportersList(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		logx.WithContext(tx.Context()).Errorf("rowsAffected in IncreaseImportersList(_), error: %v", err)
	}

	return
}

// SelectImporters
// select phone, importers from popular_contacts where phone = :phone
func (m *defaultPopularContactsModel) SelectImporters(ctx context.Context, phone string) (rValue *PopularContacts, err error) {

	var (
		query = "select phone, importers from popular_contacts where phone = ?"
		do    = &PopularContacts{}
	)
	err = m.db.QueryRowPartial(ctx, do, query, phone)

	if err != nil {
		if !errors.Is(err, sqlx.ErrNotFound) {
			logx.WithContext(ctx).Errorf("queryx in SelectImporters(_), error: %v", err)
			return
		} else {
			err = nil
		}
	} else {
		rValue = do
	}

	return
}

// SelectImportersList
// select phone, importers from popular_contacts where phone in (:phoneList)
func (m *defaultPopularContactsModel) SelectImportersList(ctx context.Context, phoneList []string) (rList []PopularContacts, err error) {
	var (
		query  = fmt.Sprintf("select phone, importers from popular_contacts where phone in (%s)", sqlx.InStringList(phoneList))
		values []PopularContacts
	)
	if len(phoneList) == 0 {
		rList = []PopularContacts{}
		return
	}

	err = m.db.QueryRowsPartial(ctx, &values, query)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectImportersList(_), error: %v", err)
		return
	}

	rList = values

	return
}

// SelectImportersListWithCB
// select phone, importers from popular_contacts where phone in (:phoneList)
func (m *defaultPopularContactsModel) SelectImportersListWithCB(ctx context.Context, phoneList []string, cb func(sz, i int, v *PopularContacts)) (rList []PopularContacts, err error) {
	var (
		query  = fmt.Sprintf("select phone, importers from popular_contacts where phone in (%s)", sqlx.InStringList(phoneList))
		values []PopularContacts
	)
	if len(phoneList) == 0 {
		rList = []PopularContacts{}
		return
	}

	err = m.db.QueryRowsPartial(ctx, &values, query)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectImportersList(_), error: %v", err)
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
