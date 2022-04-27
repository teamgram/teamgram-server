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
	"github.com/teamgram/teamgram-server/app/messenger/sync/internal/dal/dataobject"

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
// TODO(@benqi): sqlmap
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
// TODO(@benqi): sqlmap
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
// TODO(@benqi): sqlmap
func (dao *AuthSeqUpdatesDAO) SelectLastSeq(ctx context.Context, auth_id int64, user_id int64) (rValue *dataobject.AuthSeqUpdatesDO, err error) {
	var (
		query = "select seq from auth_seq_updates where auth_id = ? and user_id = ? order by seq desc limit 1"
		do    = &dataobject.AuthSeqUpdatesDO{}
	)
	err = dao.db.QueryRowPartial(ctx, do, query, auth_id, user_id)

	if err != nil {
		if err != sqlx.ErrNotFound {
			logx.WithContext(ctx).Errorf("queryx in SelectLastSeq(_), error: %v", err)
			return
		} else {
			err = nil
		}
	} else {
		rValue = do
	}

	return
}

// SelectByGtSeq
// select auth_id, user_id, seq, update_type, update_data, date2 from auth_seq_updates where auth_id = :auth_id and user_id = :user_id and seq > :seq order by seq asc
// TODO(@benqi): sqlmap
func (dao *AuthSeqUpdatesDAO) SelectByGtSeq(ctx context.Context, auth_id int64, user_id int64, seq int32) (rList []dataobject.AuthSeqUpdatesDO, err error) {
	var (
		query  = "select auth_id, user_id, seq, update_type, update_data, date2 from auth_seq_updates where auth_id = ? and user_id = ? and seq > ? order by seq asc"
		values []dataobject.AuthSeqUpdatesDO
	)
	err = dao.db.QueryRowsPartial(ctx, &values, query, auth_id, user_id, seq)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectByGtSeq(_), error: %v", err)
		return
	}

	rList = values

	return
}

// SelectByGtSeqWithCB
// select auth_id, user_id, seq, update_type, update_data, date2 from auth_seq_updates where auth_id = :auth_id and user_id = :user_id and seq > :seq order by seq asc
// TODO(@benqi): sqlmap
func (dao *AuthSeqUpdatesDAO) SelectByGtSeqWithCB(ctx context.Context, auth_id int64, user_id int64, seq int32, cb func(i int, v *dataobject.AuthSeqUpdatesDO)) (rList []dataobject.AuthSeqUpdatesDO, err error) {
	var (
		query  = "select auth_id, user_id, seq, update_type, update_data, date2 from auth_seq_updates where auth_id = ? and user_id = ? and seq > ? order by seq asc"
		values []dataobject.AuthSeqUpdatesDO
	)
	err = dao.db.QueryRowsPartial(ctx, &values, query, auth_id, user_id, seq)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectByGtSeq(_), error: %v", err)
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

// SelectByGtDate
// select auth_id, user_id, seq, update_type, update_data, date2 from auth_seq_updates where auth_id = :auth_id and user_id = :user_id and date2 > :date2 order by seq asc
// TODO(@benqi): sqlmap
func (dao *AuthSeqUpdatesDAO) SelectByGtDate(ctx context.Context, auth_id int64, user_id int64, date2 int64) (rList []dataobject.AuthSeqUpdatesDO, err error) {
	var (
		query  = "select auth_id, user_id, seq, update_type, update_data, date2 from auth_seq_updates where auth_id = ? and user_id = ? and date2 > ? order by seq asc"
		values []dataobject.AuthSeqUpdatesDO
	)
	err = dao.db.QueryRowsPartial(ctx, &values, query, auth_id, user_id, date2)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectByGtDate(_), error: %v", err)
		return
	}

	rList = values

	return
}

// SelectByGtDateWithCB
// select auth_id, user_id, seq, update_type, update_data, date2 from auth_seq_updates where auth_id = :auth_id and user_id = :user_id and date2 > :date2 order by seq asc
// TODO(@benqi): sqlmap
func (dao *AuthSeqUpdatesDAO) SelectByGtDateWithCB(ctx context.Context, auth_id int64, user_id int64, date2 int64, cb func(i int, v *dataobject.AuthSeqUpdatesDO)) (rList []dataobject.AuthSeqUpdatesDO, err error) {
	var (
		query  = "select auth_id, user_id, seq, update_type, update_data, date2 from auth_seq_updates where auth_id = ? and user_id = ? and date2 > ? order by seq asc"
		values []dataobject.AuthSeqUpdatesDO
	)
	err = dao.db.QueryRowsPartial(ctx, &values, query, auth_id, user_id, date2)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectByGtDate(_), error: %v", err)
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
