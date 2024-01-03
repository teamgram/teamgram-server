/*
 * WARNING! All changes made in this file will be lost!
 *   Created from by 'dalgen'
 *
 * Copyright (c) 2023-present,  Teamgram Authors.
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package mysql_dao

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/teamgram/marmota/pkg/stores/sqlx"
	"github.com/teamgram/teamgram-server/app/service/biz/user/internal/dal/dataobject"

	"github.com/zeromicro/go-zero/core/logx"
)

var _ *sql.Result
var _ = fmt.Sprintf
var _ = strings.Join

type CloseFriendsDAO struct {
	db *sqlx.DB
}

func NewCloseFriendsDAO(db *sqlx.DB) *CloseFriendsDAO {
	return &CloseFriendsDAO{
		db: db,
	}
}

// InsertBulk
// insert into close_friends(user_id, close_friend_id, `date`) values (:user_id, :close_friend_id, :date)
// TODO(@benqi): sqlmap
func (dao *CloseFriendsDAO) InsertBulk(ctx context.Context, doList []*dataobject.CloseFriendsDO) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into close_friends(user_id, close_friend_id, `date`) values (:user_id, :close_friend_id, :date)"
		r     sql.Result
	)

	if len(doList) == 0 {
		return
	}

	r, err = dao.db.NamedExec(ctx, query, doList)
	if err != nil {
		logx.WithContext(ctx).Errorf("namedExec in InsertBulk(%v), error: %v", doList, err)
		return
	}

	lastInsertId, err = r.LastInsertId()
	if err != nil {
		logx.WithContext(ctx).Errorf("lastInsertId in InsertBulk(%v)_error: %v", doList, err)
		return
	}
	rowsAffected, err = r.RowsAffected()
	if err != nil {
		logx.WithContext(ctx).Errorf("rowsAffected in InsertBulk(%v)_error: %v", doList, err)
	}

	return
}

// InsertBulkTx
// insert into close_friends(user_id, close_friend_id, `date`) values (:user_id, :close_friend_id, :date)
// TODO(@benqi): sqlmap
func (dao *CloseFriendsDAO) InsertBulkTx(tx *sqlx.Tx, doList []*dataobject.CloseFriendsDO) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into close_friends(user_id, close_friend_id, `date`) values (:user_id, :close_friend_id, :date)"
		r     sql.Result
	)

	if len(doList) == 0 {
		return
	}

	r, err = tx.NamedExec(query, doList)
	if err != nil {
		logx.WithContext(tx.Context()).Errorf("namedExec in InsertBulk(%v), error: %v", doList, err)
		return
	}

	lastInsertId, err = r.LastInsertId()
	if err != nil {
		logx.WithContext(tx.Context()).Errorf("lastInsertId in InsertBulk(%v)_error: %v", doList, err)
		return
	}
	rowsAffected, err = r.RowsAffected()
	if err != nil {
		logx.WithContext(tx.Context()).Errorf("rowsAffected in InsertBulk(%v)_error: %v", doList, err)
	}

	return
}

// SelectList
// select user_id, close_friend_id from close_friends where user_id = :user_id order by id asc limit :limit
// TODO(@benqi): sqlmap
func (dao *CloseFriendsDAO) SelectList(ctx context.Context, user_id int64) (rList []dataobject.CloseFriendsDO, err error) {
	var (
		query  = "select user_id, close_friend_id from close_friends where user_id = ? order by id asc limit ?"
		values []dataobject.CloseFriendsDO
	)
	err = dao.db.QueryRowsPartial(ctx, &values, query, user_id)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectList(_), error: %v", err)
		return
	}

	rList = values

	return
}

// SelectListWithCB
// select user_id, close_friend_id from close_friends where user_id = :user_id order by id asc limit :limit
// TODO(@benqi): sqlmap
func (dao *CloseFriendsDAO) SelectListWithCB(ctx context.Context, user_id int64, cb func(i int, v *dataobject.CloseFriendsDO)) (rList []dataobject.CloseFriendsDO, err error) {
	var (
		query  = "select user_id, close_friend_id from close_friends where user_id = ? order by id asc limit ?"
		values []dataobject.CloseFriendsDO
	)
	err = dao.db.QueryRowsPartial(ctx, &values, query, user_id)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectList(_), error: %v", err)
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

// Delete
// delete from close_friends where user_id = :user_id
// TODO(@benqi): sqlmap
func (dao *CloseFriendsDAO) Delete(ctx context.Context, user_id int64) (rowsAffected int64, err error) {
	var (
		query   = "delete from close_friends where user_id = ?"
		rResult sql.Result
	)
	rResult, err = dao.db.Exec(ctx, query, user_id)

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

// DeleteTx
// delete from close_friends where user_id = :user_id
// TODO(@benqi): sqlmap
func (dao *CloseFriendsDAO) DeleteTx(tx *sqlx.Tx, user_id int64) (rowsAffected int64, err error) {
	var (
		query   = "delete from close_friends where user_id = ?"
		rResult sql.Result
	)
	rResult, err = tx.Exec(query, user_id)

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
