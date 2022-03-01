/*
 * WARNING! All changes made in this file will be lost!
 *   Created from by 'dalgen'
 *
 * Copyright (c) 2021-present,  Teamgram Studio (https://teamgram.io).
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package mysql_dao

import (
	"context"
	"database/sql"

	"github.com/teamgram/marmota/pkg/stores/sqlx"
	"github.com/teamgram/teamgram-server/app/service/biz/updates/internal/dal/dataobject"

	"github.com/zeromicro/go-zero/core/logx"
)

var _ *sql.Result

type AuthSeqUpdatesDAO struct {
	db *sqlx.DB
}

func NewAuthSeqUpdatesDAO(db *sqlx.DB) *AuthSeqUpdatesDAO {
	return &AuthSeqUpdatesDAO{db}
}

// Insert
// insert into auth_seq_updates(auth_id, user_id, seq, update_type, update_data, date2) values (:auth_id, :user_id, :seq, :update_type, :update_data, :date2)
func (dao *AuthSeqUpdatesDAO) Insert(ctx context.Context, do *dataobject.AuthSeqUpdatesDO) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into auth_seq_updates(auth_id, user_id, seq, update_type, update_data, date2) values (:auth_id, :user_id, :seq, :update_type, :update_data, :date2)"
		r     sql.Result
	)

	r, err = dao.db.NamedExec(ctx, query, do)
	if err != nil {
		logx.WithContext(ctx).Errorf("namedExec in Insert(%v), error: %v", do, err)
		return
	}

	lastInsertId, err = r.LastInsertId()
	if err != nil {
		logx.WithContext(ctx).Errorf("lastInsertId in Insert(%v)_error: %v", do, err)
		return
	}
	rowsAffected, err = r.RowsAffected()
	if err != nil {
		logx.WithContext(ctx).Errorf("rowsAffected in Insert(%v)_error: %v", do, err)
	}

	return
}

// InsertTx
// insert into auth_seq_updates(auth_id, user_id, seq, update_type, update_data, date2) values (:auth_id, :user_id, :seq, :update_type, :update_data, :date2)
func (dao *AuthSeqUpdatesDAO) InsertTx(tx *sqlx.Tx, do *dataobject.AuthSeqUpdatesDO) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into auth_seq_updates(auth_id, user_id, seq, update_type, update_data, date2) values (:auth_id, :user_id, :seq, :update_type, :update_data, :date2)"
		r     sql.Result
	)

	r, err = tx.NamedExec(query, do)
	if err != nil {
		logx.WithContext(tx.Context()).Errorf("namedExec in Insert(%v), error: %v", do, err)
		return
	}

	lastInsertId, err = r.LastInsertId()
	if err != nil {
		logx.WithContext(tx.Context()).Errorf("lastInsertId in Insert(%v)_error: %v", do, err)
		return
	}
	rowsAffected, err = r.RowsAffected()
	if err != nil {
		logx.WithContext(tx.Context()).Errorf("rowsAffected in Insert(%v)_error: %v", do, err)
	}

	return
}

// SelectLastSeq
// select seq from auth_seq_updates where auth_id = :auth_id and user_id = :user_id order by seq desc limit 1
func (dao *AuthSeqUpdatesDAO) SelectLastSeq(ctx context.Context, auth_id int64, user_id int64) (rValue *dataobject.AuthSeqUpdatesDO, err error) {
	var (
		query = "select seq from auth_seq_updates where auth_id = ? and user_id = ? order by seq desc limit 1"
		rows  *sqlx.Rows
	)
	rows, err = dao.db.Query(ctx, query, auth_id, user_id)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectLastSeq(_), error: %v", err)
		return
	}

	defer rows.Close()

	do := &dataobject.AuthSeqUpdatesDO{}
	if rows.Next() {
		// TODO(@benqi): not use reflect
		err = rows.StructScan(do)
		if err != nil {
			logx.WithContext(ctx).Errorf("structScan in SelectLastSeq(_), error: %v", err)
			return
		} else {
			rValue = do
		}
	}

	return
}

// SelectByGtSeq
// select auth_id, user_id, seq, update_type, update_data, date2 from auth_seq_updates where auth_id = :auth_id and user_id = :user_id and seq > :seq order by seq asc
func (dao *AuthSeqUpdatesDAO) SelectByGtSeq(ctx context.Context, auth_id int64, user_id int64, seq int32) (rList []dataobject.AuthSeqUpdatesDO, err error) {
	var (
		query = "select auth_id, user_id, seq, update_type, update_data, date2 from auth_seq_updates where auth_id = ? and user_id = ? and seq > ? order by seq asc"
		rows  *sqlx.Rows
	)
	rows, err = dao.db.Query(ctx, query, auth_id, user_id, seq)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectByGtSeq(_), error: %v", err)
		return
	}

	defer rows.Close()

	var values []dataobject.AuthSeqUpdatesDO
	for rows.Next() {
		v := dataobject.AuthSeqUpdatesDO{}

		// TODO(@benqi): not use reflect
		err = rows.StructScan(&v)
		if err != nil {
			logx.WithContext(ctx).Errorf("structScan in SelectByGtSeq(_), error: %v", err)
			return
		}
		values = append(values, v)
	}
	rList = values

	return
}

// SelectByGtSeqWithCB
// select auth_id, user_id, seq, update_type, update_data, date2 from auth_seq_updates where auth_id = :auth_id and user_id = :user_id and seq > :seq order by seq asc
func (dao *AuthSeqUpdatesDAO) SelectByGtSeqWithCB(ctx context.Context, auth_id int64, user_id int64, seq int32, cb func(i int, v *dataobject.AuthSeqUpdatesDO)) (rList []dataobject.AuthSeqUpdatesDO, err error) {
	var (
		query = "select auth_id, user_id, seq, update_type, update_data, date2 from auth_seq_updates where auth_id = ? and user_id = ? and seq > ? order by seq asc"
		rows  *sqlx.Rows
	)
	rows, err = dao.db.Query(ctx, query, auth_id, user_id, seq)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectByGtSeq(_), error: %v", err)
		return
	}

	defer func() {
		rows.Close()
		if err == nil && cb != nil {
			for i := 0; i < len(rList); i++ {
				cb(i, &rList[i])
			}
		}
	}()

	var values []dataobject.AuthSeqUpdatesDO
	for rows.Next() {
		v := dataobject.AuthSeqUpdatesDO{}

		// TODO(@benqi): not use reflect
		err = rows.StructScan(&v)
		if err != nil {
			logx.WithContext(ctx).Errorf("structScan in SelectByGtSeq(_), error: %v", err)
			return
		}
		values = append(values, v)
	}
	rList = values

	return
}

// SelectByGtDate
// select auth_id, user_id, seq, update_type, update_data, date2 from auth_seq_updates where auth_id = :auth_id and user_id = :user_id and date2 > :date2 order by seq asc limit :limit
func (dao *AuthSeqUpdatesDAO) SelectByGtDate(ctx context.Context, auth_id int64, user_id int64, date2 int64, limit int32) (rList []dataobject.AuthSeqUpdatesDO, err error) {
	var (
		query = "select auth_id, user_id, seq, update_type, update_data, date2 from auth_seq_updates where auth_id = ? and user_id = ? and date2 > ? order by seq asc limit ?"
		rows  *sqlx.Rows
	)
	rows, err = dao.db.Query(ctx, query, auth_id, user_id, date2, limit)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectByGtDate(_), error: %v", err)
		return
	}

	defer rows.Close()

	var values []dataobject.AuthSeqUpdatesDO
	for rows.Next() {
		v := dataobject.AuthSeqUpdatesDO{}

		// TODO(@benqi): not use reflect
		err = rows.StructScan(&v)
		if err != nil {
			logx.WithContext(ctx).Errorf("structScan in SelectByGtDate(_), error: %v", err)
			return
		}
		values = append(values, v)
	}
	rList = values

	return
}

// SelectByGtDateWithCB
// select auth_id, user_id, seq, update_type, update_data, date2 from auth_seq_updates where auth_id = :auth_id and user_id = :user_id and date2 > :date2 order by seq asc limit :limit
func (dao *AuthSeqUpdatesDAO) SelectByGtDateWithCB(ctx context.Context, auth_id int64, user_id int64, date2 int64, limit int32, cb func(i int, v *dataobject.AuthSeqUpdatesDO)) (rList []dataobject.AuthSeqUpdatesDO, err error) {
	var (
		query = "select auth_id, user_id, seq, update_type, update_data, date2 from auth_seq_updates where auth_id = ? and user_id = ? and date2 > ? order by seq asc limit ?"
		rows  *sqlx.Rows
	)
	rows, err = dao.db.Query(ctx, query, auth_id, user_id, date2, limit)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectByGtDate(_), error: %v", err)
		return
	}

	defer func() {
		rows.Close()
		if err == nil && cb != nil {
			for i := 0; i < len(rList); i++ {
				cb(i, &rList[i])
			}
		}
	}()

	var values []dataobject.AuthSeqUpdatesDO
	for rows.Next() {
		v := dataobject.AuthSeqUpdatesDO{}

		// TODO(@benqi): not use reflect
		err = rows.StructScan(&v)
		if err != nil {
			logx.WithContext(ctx).Errorf("structScan in SelectByGtDate(_), error: %v", err)
			return
		}
		values = append(values, v)
	}
	rList = values

	return
}
