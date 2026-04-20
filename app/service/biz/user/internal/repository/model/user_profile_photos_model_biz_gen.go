/*
 * WARNING! All changes made in this file will be lost!
 *   Created from by 'dalgen'
 *
 * Copyright (c) 2026 The Teamgram Authors.
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
	bizUserProfilePhotosModel interface {
		InsertOrUpdate(ctx context.Context, data *UserProfilePhotos) (lastInsertId, rowsAffected int64, err error)
		InsertOrUpdateTx(tx *sqlx.Tx, data *UserProfilePhotos) (lastInsertId, rowsAffected int64, err error)

		SelectList(ctx context.Context, userId int64) ([]int64, error)
		SelectListWithCB(ctx context.Context, userId int64, cb func(sz, i int, v int64)) ([]int64, error)

		SelectNext(ctx context.Context, userId int64, idList []int64) (int64, error)

		Delete(ctx context.Context, userId int64, idList []int64) (rowsAffected int64, err error)
		DeleteTx(tx *sqlx.Tx, userId int64, idList []int64) (rowsAffected int64, err error)
	}
)

// InsertOrUpdate
// insert into user_profile_photos(user_id, photo_id, date2, deleted) values (:user_id, :photo_id, :date2, 0) on duplicate key update date2 = values(date2), deleted = 0
func (m *defaultUserProfilePhotosModel) InsertOrUpdate(ctx context.Context, data *UserProfilePhotos) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into user_profile_photos(user_id, photo_id, date2, deleted) values (:user_id, :photo_id, :date2, 0) on duplicate key update date2 = values(date2), deleted = 0"
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
// insert into user_profile_photos(user_id, photo_id, date2, deleted) values (:user_id, :photo_id, :date2, 0) on duplicate key update date2 = values(date2), deleted = 0
func (m *defaultUserProfilePhotosModel) InsertOrUpdateTx(tx *sqlx.Tx, data *UserProfilePhotos) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into user_profile_photos(user_id, photo_id, date2, deleted) values (:user_id, :photo_id, :date2, 0) on duplicate key update date2 = values(date2), deleted = 0"
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

// SelectList
// select photo_id from user_profile_photos where user_id = :user_id and deleted = 0 order by date2 desc
func (m *defaultUserProfilePhotosModel) SelectList(ctx context.Context, userId int64) (rList []int64, err error) {
	var query = "select photo_id from user_profile_photos where user_id = ? and deleted = 0 order by date2 desc"
	err = m.db.QueryRowsPartial(ctx, &rList, query, userId)

	if err != nil {
		logx.WithContext(ctx).Errorf("select in SelectList(_), error: %v", err)
	}

	return
}

// SelectListWithCB
// select photo_id from user_profile_photos where user_id = :user_id and deleted = 0 order by date2 desc
func (m *defaultUserProfilePhotosModel) SelectListWithCB(ctx context.Context, userId int64, cb func(sz, i int, v int64)) (rList []int64, err error) {
	var query = "select photo_id from user_profile_photos where user_id = ? and deleted = 0 order by date2 desc"
	err = m.db.QueryRowsPartial(ctx, &rList, query, userId)

	if err != nil {
		logx.WithContext(ctx).Errorf("select in SelectList(_), error: %v", err)
	}

	if cb != nil {
		sz := len(rList)
		for i := 0; i < sz; i++ {
			cb(sz, i, rList[i])
		}
	}

	return
}

// SelectNext
// select photo_id from user_profile_photos where user_id = :user_id and photo_id not in (:id_list) and deleted = 0 order by date2 desc limit 1
func (m *defaultUserProfilePhotosModel) SelectNext(ctx context.Context, userId int64, idList []int64) (rValue int64, err error) {
	var (
		query = fmt.Sprintf("select photo_id from user_profile_photos where user_id = ? and photo_id not in (%s) and deleted = 0 order by date2 desc limit 1", sqlx.InInt64List(idList))
	)

	if len(idList) == 0 {
		return
	}

	err = m.db.QueryRowPartial(ctx, &rValue, query, userId)

	if err != nil {
		if !errors.Is(err, sqlx.ErrNotFound) {
			logx.WithContext(ctx).Errorf("get in SelectNext(_), error: %v", err)
			return
		} else {
			err = nil
		}
	}

	return
}

// Delete
// update user_profile_photos set deleted = 1, date2 = 0 where user_id = :user_id and photo_id in (:id_list)
func (m *defaultUserProfilePhotosModel) Delete(ctx context.Context, userId int64, idList []int64) (rowsAffected int64, err error) {
	var (
		query   = fmt.Sprintf("update user_profile_photos set deleted = 1, date2 = 0 where user_id = ? and photo_id in (%s)", sqlx.InInt64List(idList))
		rResult sql.Result
	)

	if len(idList) == 0 {
		return
	}

	rResult, err = m.db.Exec(ctx, query, userId)

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
func (m *defaultUserProfilePhotosModel) DeleteTx(tx *sqlx.Tx, userId int64, idList []int64) (rowsAffected int64, err error) {
	var (
		query   = fmt.Sprintf("update user_profile_photos set deleted = 1, date2 = 0 where user_id = ? and photo_id in (%s)", sqlx.InInt64List(idList))
		rResult sql.Result
	)

	if len(idList) == 0 {
		return
	}

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
