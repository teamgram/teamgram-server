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
	"github.com/teamgram/teamgram-server/app/service/biz/user/internal/dal/dataobject"

	"github.com/zeromicro/go-zero/core/logx"
)

var _ *sql.Result

type UserPeerBlocksDAO struct {
	db *sqlx.DB
}

func NewUserPeerBlocksDAO(db *sqlx.DB) *UserPeerBlocksDAO {
	return &UserPeerBlocksDAO{db}
}

// InsertOrUpdate
// insert into user_peer_blocks(user_id, peer_type, peer_id, `date`) values (:user_id, :peer_type, :peer_id, :date) on duplicate key update `date` = values(`date`), deleted = 0
// TODO(@benqi): sqlmap
func (dao *UserPeerBlocksDAO) InsertOrUpdate(ctx context.Context, do *dataobject.UserPeerBlocksDO) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into user_peer_blocks(user_id, peer_type, peer_id, `date`) values (:user_id, :peer_type, :peer_id, :date) on duplicate key update `date` = values(`date`), deleted = 0"
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
// insert into user_peer_blocks(user_id, peer_type, peer_id, `date`) values (:user_id, :peer_type, :peer_id, :date) on duplicate key update `date` = values(`date`), deleted = 0
// TODO(@benqi): sqlmap
func (dao *UserPeerBlocksDAO) InsertOrUpdateTx(tx *sqlx.Tx, do *dataobject.UserPeerBlocksDO) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into user_peer_blocks(user_id, peer_type, peer_id, `date`) values (:user_id, :peer_type, :peer_id, :date) on duplicate key update `date` = values(`date`), deleted = 0"
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

// SelectList
// select user_id, peer_type, peer_id, `date` from user_peer_blocks where user_id = :user_id and deleted = 0 order by id asc limit :limit
// TODO(@benqi): sqlmap
func (dao *UserPeerBlocksDAO) SelectList(ctx context.Context, user_id int64, limit int32) (rList []dataobject.UserPeerBlocksDO, err error) {
	var (
		query = "select user_id, peer_type, peer_id, `date` from user_peer_blocks where user_id = ? and deleted = 0 order by id asc limit ?"
		rows  *sqlx.Rows
	)
	rows, err = dao.db.Query(ctx, query, user_id, limit)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectList(_), error: %v", err)
		return
	}

	defer rows.Close()

	var values []dataobject.UserPeerBlocksDO
	for rows.Next() {
		v := dataobject.UserPeerBlocksDO{}

		// TODO(@benqi): not use reflect
		err = rows.StructScan(&v)
		if err != nil {
			logx.WithContext(ctx).Errorf("structScan in SelectList(_), error: %v", err)
			return
		}
		values = append(values, v)
	}
	rList = values

	return
}

// SelectListWithCB
// select user_id, peer_type, peer_id, `date` from user_peer_blocks where user_id = :user_id and deleted = 0 order by id asc limit :limit
// TODO(@benqi): sqlmap
func (dao *UserPeerBlocksDAO) SelectListWithCB(ctx context.Context, user_id int64, limit int32, cb func(i int, v *dataobject.UserPeerBlocksDO)) (rList []dataobject.UserPeerBlocksDO, err error) {
	var (
		query = "select user_id, peer_type, peer_id, `date` from user_peer_blocks where user_id = ? and deleted = 0 order by id asc limit ?"
		rows  *sqlx.Rows
	)
	rows, err = dao.db.Query(ctx, query, user_id, limit)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectList(_), error: %v", err)
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

	var values []dataobject.UserPeerBlocksDO
	for rows.Next() {
		v := dataobject.UserPeerBlocksDO{}

		// TODO(@benqi): not use reflect
		err = rows.StructScan(&v)
		if err != nil {
			logx.WithContext(ctx).Errorf("structScan in SelectList(_), error: %v", err)
			return
		}
		values = append(values, v)
	}
	rList = values

	return
}

// SelectListByIdList
// select peer_id from user_peer_blocks where user_id = :user_id and peer_type = 2 and peer_id in (:idList) and deleted = 0
// TODO(@benqi): sqlmap
func (dao *UserPeerBlocksDAO) SelectListByIdList(ctx context.Context, user_id int64, idList []int64) (rList []int64, err error) {
	var (
		query = "select peer_id from user_peer_blocks where user_id = ? and peer_type = 2 and peer_id in (?) and deleted = 0"
		a     []interface{}
	)

	if len(idList) == 0 {
		rList = []int64{}
		return
	}

	query, a, err = sqlx.In(query, user_id, idList)
	if err != nil {
		// r sql.Result
		logx.WithContext(ctx).Errorf("sqlx.In in SelectListByIdList(_), error: %v", err)
		return
	}

	err = dao.db.Select(ctx, &rList, query, a...)

	if err != nil {
		logx.WithContext(ctx).Errorf("select in SelectListByIdList(_), error: %v", err)
	}

	return
}

// SelectListByIdListWithCB
// select peer_id from user_peer_blocks where user_id = :user_id and peer_type = 2 and peer_id in (:idList) and deleted = 0
// TODO(@benqi): sqlmap
func (dao *UserPeerBlocksDAO) SelectListByIdListWithCB(ctx context.Context, user_id int64, idList []int64, cb func(i int, v int64)) (rList []int64, err error) {
	var (
		query = "select peer_id from user_peer_blocks where user_id = ? and peer_type = 2 and peer_id in (?) and deleted = 0"
		a     []interface{}
	)

	if len(idList) == 0 {
		rList = []int64{}
		return
	}

	query, a, err = sqlx.In(query, user_id, idList)
	if err != nil {
		// r sql.Result
		logx.WithContext(ctx).Errorf("sqlx.In in SelectListByIdList(_), error: %v", err)
		return
	}

	err = dao.db.Select(ctx, &rList, query, a...)

	if err != nil {
		logx.WithContext(ctx).Errorf("select in SelectListByIdList(_), error: %v", err)
	}

	if cb != nil {
		for i := 0; i < len(rList); i++ {
			cb(i, rList[i])
		}
	}

	return
}

// Select
// select user_id, peer_type, peer_id, `date` from user_peer_blocks where user_id = :user_id and peer_type = :peer_type and peer_id = :peer_id and deleted = 0
// TODO(@benqi): sqlmap
func (dao *UserPeerBlocksDAO) Select(ctx context.Context, user_id int64, peer_type int32, peer_id int64) (rValue *dataobject.UserPeerBlocksDO, err error) {
	var (
		query = "select user_id, peer_type, peer_id, `date` from user_peer_blocks where user_id = ? and peer_type = ? and peer_id = ? and deleted = 0"
		rows  *sqlx.Rows
	)
	rows, err = dao.db.Query(ctx, query, user_id, peer_type, peer_id)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in Select(_), error: %v", err)
		return
	}

	defer rows.Close()

	do := &dataobject.UserPeerBlocksDO{}
	if rows.Next() {
		// TODO(@benqi): not use reflect
		err = rows.StructScan(do)
		if err != nil {
			logx.WithContext(ctx).Errorf("structScan in Select(_), error: %v", err)
			return
		} else {
			rValue = do
		}
	}

	return
}

// Delete
// update user_peer_blocks set deleted = 1, `date` = 0 where user_id = :user_id and peer_type = :peer_type and peer_id = :peer_id
// TODO(@benqi): sqlmap
func (dao *UserPeerBlocksDAO) Delete(ctx context.Context, user_id int64, peer_type int32, peer_id int64) (rowsAffected int64, err error) {
	var (
		query   = "update user_peer_blocks set deleted = 1, `date` = 0 where user_id = ? and peer_type = ? and peer_id = ?"
		rResult sql.Result
	)
	rResult, err = dao.db.Exec(ctx, query, user_id, peer_type, peer_id)

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

// update user_peer_blocks set deleted = 1, `date` = 0 where user_id = :user_id and peer_type = :peer_type and peer_id = :peer_id
// DeleteTx
// TODO(@benqi): sqlmap
func (dao *UserPeerBlocksDAO) DeleteTx(tx *sqlx.Tx, user_id int64, peer_type int32, peer_id int64) (rowsAffected int64, err error) {
	var (
		query   = "update user_peer_blocks set deleted = 1, `date` = 0 where user_id = ? and peer_type = ? and peer_id = ?"
		rResult sql.Result
	)
	rResult, err = tx.Exec(query, user_id, peer_type, peer_id)

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
