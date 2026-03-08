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

	"github.com/teamgram/marmota/pkg/stores/sqlx"
	"github.com/teamgram/teamgram-server/app/service/biz/user/internal/dal/dataobject"

	"github.com/zeromicro/go-zero/core/logx"
)

type PopularContactsDAO struct {
	db *sqlx.DB
}

func NewPopularContactsDAO(db *sqlx.DB) *PopularContactsDAO {
	return &PopularContactsDAO{
		db: db,
	}
}

// InsertOrUpdate
// insert into popular_contacts(phone, importers, deleted) values (:phone, :importers, 0) on duplicate key update importers = importers + 1
func (dao *PopularContactsDAO) InsertOrUpdate(ctx context.Context, do *dataobject.PopularContactsDO) (lastInsertId, rowsAffected int64, err error) {
	var (
		query string
		r     sql.Result
	)
	query = "insert into popular_contacts(phone, importers, deleted) values (:phone, :importers, 0) on duplicate key update importers = importers + 1"

	r, err = dao.db.NamedExec(ctx, query, do)
	if err != nil {
		logx.WithContext(ctx).Errorf("namedExec in InsertOrUpdate(%v), error: %v", do, err)
		return
	}

	lastInsertId, err = r.LastInsertId()
	if err != nil {
		logx.WithContext(ctx).Errorf("lastInsertId in InsertOrUpdate(%v), error: %v", do, err)
		return
	}
	rowsAffected, err = r.RowsAffected()
	if err != nil {
		logx.WithContext(ctx).Errorf("rowsAffected in InsertOrUpdate(%v), error: %v", do, err)
	}

	return
}

// InsertOrUpdateTx
// insert into popular_contacts(phone, importers, deleted) values (:phone, :importers, 0) on duplicate key update importers = importers + 1
func (dao *PopularContactsDAO) InsertOrUpdateTx(tx *sqlx.Tx, do *dataobject.PopularContactsDO) (lastInsertId, rowsAffected int64, err error) {
	var (
		query string
		r     sql.Result
	)
	query = "insert into popular_contacts(phone, importers, deleted) values (:phone, :importers, 0) on duplicate key update importers = importers + 1"

	r, err = tx.NamedExec(query, do)
	if err != nil {
		logx.WithContext(tx.Context()).Errorf("namedExec in InsertOrUpdate(%v), error: %v", do, err)
		return
	}

	lastInsertId, err = r.LastInsertId()
	if err != nil {
		logx.WithContext(tx.Context()).Errorf("lastInsertId in InsertOrUpdate(%v), error: %v", do, err)
		return
	}
	rowsAffected, err = r.RowsAffected()
	if err != nil {
		logx.WithContext(tx.Context()).Errorf("rowsAffected in InsertOrUpdate(%v), error: %v", do, err)
	}

	return
}

// IncreaseImporters
// update popular_contacts set importers = importers + 1 where phone = :phone
func (dao *PopularContactsDAO) IncreaseImporters(ctx context.Context, phone string) (rowsAffected int64, err error) {
	var (
		query   string
		rResult sql.Result
	)
	query = "update popular_contacts set importers = importers + 1 where phone = ?"

	rResult, err = dao.db.Exec(ctx, query, phone)

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
func (dao *PopularContactsDAO) IncreaseImportersTx(tx *sqlx.Tx, phone string) (rowsAffected int64, err error) {
	var (
		query   string
		rResult sql.Result
	)
	query = "update popular_contacts set importers = importers + 1 where phone = ?"

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
func (dao *PopularContactsDAO) IncreaseImportersList(ctx context.Context, phoneList []string) (rowsAffected int64, err error) {
	if len(phoneList) == 0 {
		return
	}

	var (
		query   string
		rResult sql.Result
	)
	query = fmt.Sprintf("update popular_contacts set importers = importers + 1 where phone in (%s)", sqlx.InStringList(phoneList))

	rResult, err = dao.db.Exec(ctx, query)

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
func (dao *PopularContactsDAO) IncreaseImportersListTx(tx *sqlx.Tx, phoneList []string) (rowsAffected int64, err error) {
	if len(phoneList) == 0 {
		return
	}
	var (
		query   string
		rResult sql.Result
	)
	query = fmt.Sprintf("update popular_contacts set importers = importers + 1 where phone in (%s)", sqlx.InStringList(phoneList))

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
func (dao *PopularContactsDAO) SelectImporters(ctx context.Context, phone string) (rValue *dataobject.PopularContactsDO, err error) {
	var (
		query string
		do    = &dataobject.PopularContactsDO{}
	)
	query = "select phone, importers from popular_contacts where phone = ?"

	err = dao.db.QueryRowPartial(ctx, do, query, phone)

	if err != nil {
		if !errors.Is(err, sqlx.ErrNotFound) {
			logx.WithContext(ctx).Errorf("queryx in SelectImporters(_), error: %v", err)
			return
		} else {
			// not found not error, return nil, nil
			err = nil
		}
	} else {
		rValue = do
	}

	return
}

// SelectImportersList
// select phone, importers from popular_contacts where phone in (:phoneList)
func (dao *PopularContactsDAO) SelectImportersList(ctx context.Context, phoneList []string) (rList []dataobject.PopularContactsDO, err error) {
	if len(phoneList) == 0 {
		rList = []dataobject.PopularContactsDO{}
		return
	}

	var (
		query  string
		values []dataobject.PopularContactsDO
	)
	query = fmt.Sprintf("select phone, importers from popular_contacts where phone in (%s)", sqlx.InStringList(phoneList))

	err = dao.db.QueryRowsPartial(ctx, &values, query)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectImportersList(_), error: %v", err)
		return
	}

	rList = values

	return
}

// SelectImportersListWithCB
// select phone, importers from popular_contacts where phone in (:phoneList)
func (dao *PopularContactsDAO) SelectImportersListWithCB(ctx context.Context, phoneList []string, cb func(sz, i int, v *dataobject.PopularContactsDO)) (rList []dataobject.PopularContactsDO, err error) {
	if len(phoneList) == 0 {
		rList = []dataobject.PopularContactsDO{}
		return
	}

	var (
		query  string
		values []dataobject.PopularContactsDO
	)
	query = fmt.Sprintf("select phone, importers from popular_contacts where phone in (%s)", sqlx.InStringList(phoneList))

	err = dao.db.QueryRowsPartial(ctx, &values, query)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectImportersList(_), error: %v", err)
		return
	}

	rList = values

	if cb != nil {
		sz := len(rList)
		for i := range sz {
			cb(sz, i, &rList[i])
		}
	}

	return
}
