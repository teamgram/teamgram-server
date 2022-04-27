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

type PopularContactsDAO struct {
	db *sqlx.DB
}

func NewPopularContactsDAO(db *sqlx.DB) *PopularContactsDAO {
	return &PopularContactsDAO{db}
}

// InsertOrUpdate
// insert into popular_contacts(phone, importers, deleted) values (:phone, :importers, 0) on duplicate key update importers = importers + 1
// TODO(@benqi): sqlmap
func (dao *PopularContactsDAO) InsertOrUpdate(ctx context.Context, do *dataobject.PopularContactsDO) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into popular_contacts(phone, importers, deleted) values (:phone, :importers, 0) on duplicate key update importers = importers + 1"
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
// insert into popular_contacts(phone, importers, deleted) values (:phone, :importers, 0) on duplicate key update importers = importers + 1
// TODO(@benqi): sqlmap
func (dao *PopularContactsDAO) InsertOrUpdateTx(tx *sqlx.Tx, do *dataobject.PopularContactsDO) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into popular_contacts(phone, importers, deleted) values (:phone, :importers, 0) on duplicate key update importers = importers + 1"
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

// IncreaseImporters
// update popular_contacts set importers = importers + 1 where phone = :phone
// TODO(@benqi): sqlmap
func (dao *PopularContactsDAO) IncreaseImporters(ctx context.Context, phone string) (rowsAffected int64, err error) {
	var (
		query   = "update popular_contacts set importers = importers + 1 where phone = ?"
		rResult sql.Result
	)
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

// update popular_contacts set importers = importers + 1 where phone = :phone
// IncreaseImportersTx
// TODO(@benqi): sqlmap
func (dao *PopularContactsDAO) IncreaseImportersTx(tx *sqlx.Tx, phone string) (rowsAffected int64, err error) {
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
// TODO(@benqi): sqlmap
func (dao *PopularContactsDAO) IncreaseImportersList(ctx context.Context, phoneList []string) (rowsAffected int64, err error) {
	var (
		query   = "update popular_contacts set importers = importers + 1 where phone in (?)"
		a       []interface{}
		rResult sql.Result
	)

	if len(phoneList) == 0 {
		return
	}

	query, a, err = sqlx.In(query, phoneList)
	if err != nil {
		// r sql.Result
		logx.WithContext(ctx).Errorf("sqlx.In in IncreaseImportersList(_), error: %v", err)
		return
	}
	rResult, err = dao.db.Exec(ctx, query, a...)

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

// update popular_contacts set importers = importers + 1 where phone in (:phoneList)
// IncreaseImportersListTx
// TODO(@benqi): sqlmap
func (dao *PopularContactsDAO) IncreaseImportersListTx(tx *sqlx.Tx, phoneList []string) (rowsAffected int64, err error) {
	var (
		query   = "update popular_contacts set importers = importers + 1 where phone in (?)"
		a       []interface{}
		rResult sql.Result
	)

	if len(phoneList) == 0 {
		return
	}

	query, a, err = sqlx.In(query, phoneList)
	if err != nil {
		// r sql.Result
		logx.WithContext(tx.Context()).Errorf("sqlx.In in IncreaseImportersList(_), error: %v", err)
		return
	}
	rResult, err = tx.Exec(query, a...)

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
// TODO(@benqi): sqlmap
func (dao *PopularContactsDAO) SelectImporters(ctx context.Context, phone string) (rValue *dataobject.PopularContactsDO, err error) {
	var (
		query = "select phone, importers from popular_contacts where phone = ?"
		do    = &dataobject.PopularContactsDO{}
	)
	err = dao.db.QueryRowPartial(ctx, do, query, phone)

	if err != nil {
		if err != sqlx.ErrNotFound {
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
// TODO(@benqi): sqlmap
func (dao *PopularContactsDAO) SelectImportersList(ctx context.Context, phoneList []string) (rList []dataobject.PopularContactsDO, err error) {
	var (
		query  = "select phone, importers from popular_contacts where phone in (?)"
		a      []interface{}
		values []dataobject.PopularContactsDO
	)
	if len(phoneList) == 0 {
		rList = []dataobject.PopularContactsDO{}
		return
	}

	query, a, err = sqlx.In(query, phoneList)
	if err != nil {
		// r sql.Result
		logx.WithContext(ctx).Errorf("sqlx.In in SelectImportersList(_), error: %v", err)
		return
	}
	err = dao.db.QueryRowsPartial(ctx, &values, query, a...)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectImportersList(_), error: %v", err)
		return
	}

	rList = values

	return
}

// SelectImportersListWithCB
// select phone, importers from popular_contacts where phone in (:phoneList)
// TODO(@benqi): sqlmap
func (dao *PopularContactsDAO) SelectImportersListWithCB(ctx context.Context, phoneList []string, cb func(i int, v *dataobject.PopularContactsDO)) (rList []dataobject.PopularContactsDO, err error) {
	var (
		query  = "select phone, importers from popular_contacts where phone in (?)"
		a      []interface{}
		values []dataobject.PopularContactsDO
	)
	if len(phoneList) == 0 {
		rList = []dataobject.PopularContactsDO{}
		return
	}

	query, a, err = sqlx.In(query, phoneList)
	if err != nil {
		// r sql.Result
		logx.WithContext(ctx).Errorf("sqlx.In in SelectImportersList(_), error: %v", err)
		return
	}
	err = dao.db.QueryRowsPartial(ctx, &values, query, a...)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectImportersList(_), error: %v", err)
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
