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

type UserProfilePhotosDAO struct {
	db *sqlx.DB
}

func NewUserProfilePhotosDAO(db *sqlx.DB) *UserProfilePhotosDAO {
	return &UserProfilePhotosDAO{
		db: db,
	}
}

// InsertOrUpdate
// insert into user_profile_photos(user_id, photo_id, date2, deleted) values (:user_id, :photo_id, :date2, 0) on duplicate key update date2 = values(date2), deleted = 0
func (dao *UserProfilePhotosDAO) InsertOrUpdate(ctx context.Context, do *dataobject.UserProfilePhotosDO) (lastInsertId, rowsAffected int64, err error) {
	var (
		query string
		r     sql.Result
	)
	query = "insert into user_profile_photos(user_id, photo_id, date2, deleted) values (:user_id, :photo_id, :date2, 0) on duplicate key update date2 = values(date2), deleted = 0"

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
// insert into user_profile_photos(user_id, photo_id, date2, deleted) values (:user_id, :photo_id, :date2, 0) on duplicate key update date2 = values(date2), deleted = 0
func (dao *UserProfilePhotosDAO) InsertOrUpdateTx(tx *sqlx.Tx, do *dataobject.UserProfilePhotosDO) (lastInsertId, rowsAffected int64, err error) {
	var (
		query string
		r     sql.Result
	)
	query = "insert into user_profile_photos(user_id, photo_id, date2, deleted) values (:user_id, :photo_id, :date2, 0) on duplicate key update date2 = values(date2), deleted = 0"

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

// SelectList
// select photo_id from user_profile_photos where user_id = :user_id and deleted = 0 order by date2 desc
func (dao *UserProfilePhotosDAO) SelectList(ctx context.Context, userId int64) (rList []int64, err error) {
	var query string
	query = "select photo_id from user_profile_photos where user_id = ? and deleted = 0 order by date2 desc"

	err = dao.db.QueryRowsPartial(ctx, &rList, query, userId)

	if err != nil {
		logx.WithContext(ctx).Errorf("select in SelectList(_), error: %v", err)
	}

	return
}

// SelectListWithCB
// select photo_id from user_profile_photos where user_id = :user_id and deleted = 0 order by date2 desc
func (dao *UserProfilePhotosDAO) SelectListWithCB(ctx context.Context, userId int64, cb func(sz, i int, v int64)) (rList []int64, err error) {
	var query string
	query = "select photo_id from user_profile_photos where user_id = ? and deleted = 0 order by date2 desc"

	err = dao.db.QueryRowsPartial(ctx, &rList, query, userId)

	if err != nil {
		logx.WithContext(ctx).Errorf("select in SelectList(_), error: %v", err)
	}

	if cb != nil {
		sz := len(rList)
		for i := range sz {
			cb(sz, i, rList[i])
		}
	}

	return
}

// SelectNext
// select photo_id from user_profile_photos where user_id = :user_id and photo_id not in (:id_list) and deleted = 0 order by date2 desc limit 1
func (dao *UserProfilePhotosDAO) SelectNext(ctx context.Context, userId int64, idList []int64) (rValue int64, err error) {

	if len(idList) == 0 {
		return
	}
	var (
		query string
	)
	query = fmt.Sprintf("select photo_id from user_profile_photos where user_id = ? and photo_id not in (%s) and deleted = 0 order by date2 desc limit 1", sqlx.InInt64List(idList))

	err = dao.db.QueryRowPartial(ctx, &rValue, query, userId)

	if err != nil {
		if !errors.Is(err, sqlx.ErrNotFound) {
			logx.WithContext(ctx).Errorf("get in SelectNext(_), error: %v", err)
			return
		} else {
			// not found not error, return nil, nil
			err = nil
		}
	}

	return
}

// Delete
// update user_profile_photos set deleted = 1, date2 = 0 where user_id = :user_id and photo_id in (:id_list)
func (dao *UserProfilePhotosDAO) Delete(ctx context.Context, userId int64, idList []int64) (rowsAffected int64, err error) {

	if len(idList) == 0 {
		return
	}

	var (
		query   string
		rResult sql.Result
	)
	query = fmt.Sprintf("update user_profile_photos set deleted = 1, date2 = 0 where user_id = ? and photo_id in (%s)", sqlx.InInt64List(idList))

	rResult, err = dao.db.Exec(ctx, query, userId)

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
// update user_profile_photos set deleted = 1, date2 = 0 where user_id = :user_id and photo_id in (:id_list)
func (dao *UserProfilePhotosDAO) DeleteTx(tx *sqlx.Tx, userId int64, idList []int64) (rowsAffected int64, err error) {

	if len(idList) == 0 {
		return
	}
	var (
		query   string
		rResult sql.Result
	)
	query = fmt.Sprintf("update user_profile_photos set deleted = 1, date2 = 0 where user_id = ? and photo_id in (%s)", sqlx.InInt64List(idList))

	rResult, err = tx.Exec(query, userId)

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
