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

type UserSavedMusicDAO struct {
	db *sqlx.DB
}

func NewUserSavedMusicDAO(db *sqlx.DB) *UserSavedMusicDAO {
	return &UserSavedMusicDAO{
		db: db,
	}
}

// InsertOrUpdate
// insert into user_saved_music(user_id, saved_music_id, order2) values (:user_id, :saved_music_id, :order2) on duplicate key update deleted = 0
func (dao *UserSavedMusicDAO) InsertOrUpdate(ctx context.Context, do *dataobject.UserSavedMusicDO) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into user_saved_music(user_id, saved_music_id, order2) values (:user_id, :saved_music_id, :order2) on duplicate key update deleted = 0"
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
// insert into user_saved_music(user_id, saved_music_id, order2) values (:user_id, :saved_music_id, :order2) on duplicate key update deleted = 0
func (dao *UserSavedMusicDAO) InsertOrUpdateTx(tx *sqlx.Tx, do *dataobject.UserSavedMusicDO) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into user_saved_music(user_id, saved_music_id, order2) values (:user_id, :saved_music_id, :order2) on duplicate key update deleted = 0"
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
// select id, user_id, saved_music_id from user_saved_music where user_id = :user_id and deleted = 0
func (dao *UserSavedMusicDAO) SelectList(ctx context.Context, userId int64) (rList []dataobject.UserSavedMusicDO, err error) {
	var (
		query  = "select id, user_id, saved_music_id from user_saved_music where user_id = ? and deleted = 0"
		values []dataobject.UserSavedMusicDO
	)
	err = dao.db.QueryRowsPartial(ctx, &values, query, userId)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectList(_), error: %v", err)
		return
	}

	rList = values

	return
}

// SelectListWithCB
// select id, user_id, saved_music_id from user_saved_music where user_id = :user_id and deleted = 0
func (dao *UserSavedMusicDAO) SelectListWithCB(ctx context.Context, userId int64, cb func(sz, i int, v *dataobject.UserSavedMusicDO)) (rList []dataobject.UserSavedMusicDO, err error) {
	var (
		query  = "select id, user_id, saved_music_id from user_saved_music where user_id = ? and deleted = 0"
		values []dataobject.UserSavedMusicDO
	)
	err = dao.db.QueryRowsPartial(ctx, &values, query, userId)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectList(_), error: %v", err)
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

// SelectListByIdList
// select id, user_id, saved_music_id from user_saved_music where user_id = :user_id and deleted = 0 and saved_music_id in (:idList)
func (dao *UserSavedMusicDAO) SelectListByIdList(ctx context.Context, userId int64, idList []int64) (rList []dataobject.UserSavedMusicDO, err error) {
	var (
		query  = fmt.Sprintf("select id, user_id, saved_music_id from user_saved_music where user_id = ? and deleted = 0 and saved_music_id in (%s)", sqlx.InInt64List(idList))
		values []dataobject.UserSavedMusicDO
	)

	if len(idList) == 0 {
		rList = []dataobject.UserSavedMusicDO{}
		return
	}

	err = dao.db.QueryRowsPartial(ctx, &values, query, userId)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectListByIdList(_), error: %v", err)
		return
	}

	rList = values

	return
}

// SelectListByIdListWithCB
// select id, user_id, saved_music_id from user_saved_music where user_id = :user_id and deleted = 0 and saved_music_id in (:idList)
func (dao *UserSavedMusicDAO) SelectListByIdListWithCB(ctx context.Context, userId int64, idList []int64, cb func(sz, i int, v *dataobject.UserSavedMusicDO)) (rList []dataobject.UserSavedMusicDO, err error) {
	var (
		query  = fmt.Sprintf("select id, user_id, saved_music_id from user_saved_music where user_id = ? and deleted = 0 and saved_music_id in (%s)", sqlx.InInt64List(idList))
		values []dataobject.UserSavedMusicDO
	)

	if len(idList) == 0 {
		rList = []dataobject.UserSavedMusicDO{}
		return
	}

	err = dao.db.QueryRowsPartial(ctx, &values, query, userId)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectListByIdList(_), error: %v", err)
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

// Delete
// update user_saved_music set deleted = 1, order2 = 0 where user_id = :user_id and saved_music_id = :saved_music_id
func (dao *UserSavedMusicDAO) Delete(ctx context.Context, userId int64, savedMusicId int64) (rowsAffected int64, err error) {
	var (
		query   = "update user_saved_music set deleted = 1, order2 = 0 where user_id = ? and saved_music_id = ?"
		rResult sql.Result
	)

	rResult, err = dao.db.Exec(ctx, query, userId, savedMusicId)

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
// update user_saved_music set deleted = 1, order2 = 0 where user_id = :user_id and saved_music_id = :saved_music_id
func (dao *UserSavedMusicDAO) DeleteTx(tx *sqlx.Tx, userId int64, savedMusicId int64) (rowsAffected int64, err error) {
	var (
		query   = "update user_saved_music set deleted = 1, order2 = 0 where user_id = ? and saved_music_id = ?"
		rResult sql.Result
	)
	rResult, err = tx.Exec(query, userId, savedMusicId)

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
