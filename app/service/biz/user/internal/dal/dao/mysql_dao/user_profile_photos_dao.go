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

type UserProfilePhotosDAO struct {
	db *sqlx.DB
}

func NewUserProfilePhotosDAO(db *sqlx.DB) *UserProfilePhotosDAO {
	return &UserProfilePhotosDAO{db}
}

// InsertOrUpdate
// insert into user_profile_photos(user_id, photo_id, date2, deleted) values (:user_id, :photo_id, :date2, 0) on duplicate key update deleted = 0
// TODO(@benqi): sqlmap
func (dao *UserProfilePhotosDAO) InsertOrUpdate(ctx context.Context, do *dataobject.UserProfilePhotosDO) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into user_profile_photos(user_id, photo_id, date2, deleted) values (:user_id, :photo_id, :date2, 0) on duplicate key update deleted = 0"
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
// insert into user_profile_photos(user_id, photo_id, date2, deleted) values (:user_id, :photo_id, :date2, 0) on duplicate key update deleted = 0
// TODO(@benqi): sqlmap
func (dao *UserProfilePhotosDAO) InsertOrUpdateTx(tx *sqlx.Tx, do *dataobject.UserProfilePhotosDO) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into user_profile_photos(user_id, photo_id, date2, deleted) values (:user_id, :photo_id, :date2, 0) on duplicate key update deleted = 0"
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
// select photo_id from user_profile_photos where user_id = :user_id and deleted = 0 order by id asc
// TODO(@benqi): sqlmap
func (dao *UserProfilePhotosDAO) SelectList(ctx context.Context, user_id int64) (rList []int64, err error) {
	var query = "select photo_id from user_profile_photos where user_id = ? and deleted = 0 order by id asc"
	err = dao.db.QueryRowsPartial(ctx, &rList, query, user_id)

	if err != nil {
		logx.WithContext(ctx).Errorf("select in SelectList(_), error: %v", err)
	}

	return
}

// SelectListWithCB
// select photo_id from user_profile_photos where user_id = :user_id and deleted = 0 order by id asc
// TODO(@benqi): sqlmap
func (dao *UserProfilePhotosDAO) SelectListWithCB(ctx context.Context, user_id int64, cb func(i int, v int64)) (rList []int64, err error) {
	var query = "select photo_id from user_profile_photos where user_id = ? and deleted = 0 order by id asc"
	err = dao.db.QueryRowsPartial(ctx, &rList, query, user_id)

	if err != nil {
		logx.WithContext(ctx).Errorf("select in SelectList(_), error: %v", err)
	}

	if cb != nil {
		for i := 0; i < len(rList); i++ {
			cb(i, rList[i])
		}
	}

	return
}

// SelectNext
// select photo_id from user_profile_photos where user_id = :user_id and photo_id not in (:id_list) and deleted = 0 order by id asc limit 1
// TODO(@benqi): sqlmap
func (dao *UserProfilePhotosDAO) SelectNext(ctx context.Context, user_id int64, id_list []int64) (rValue int64, err error) {
	var (
		query = "select photo_id from user_profile_photos where user_id = ? and photo_id not in (?) and deleted = 0 order by id asc limit 1"
		a     []interface{}
	)

	if len(id_list) == 0 {
		return
	}

	query, a, err = sqlx.In(query, user_id, id_list)
	if err != nil {
		// r sql.Result
		logx.WithContext(ctx).Errorf("sqlx.In in SelectNext(_), error: %v", err)
		return
	}

	err = dao.db.QueryRowPartial(ctx, &rValue, query, a...)

	if err != nil {
		if err != sqlx.ErrNotFound {
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
// TODO(@benqi): sqlmap
func (dao *UserProfilePhotosDAO) Delete(ctx context.Context, user_id int64, id_list []int64) (rowsAffected int64, err error) {
	var (
		query   = "update user_profile_photos set deleted = 1, date2 = 0 where user_id = ? and photo_id in (?)"
		a       []interface{}
		rResult sql.Result
	)

	if len(id_list) == 0 {
		return
	}

	query, a, err = sqlx.In(query, user_id, id_list)
	if err != nil {
		// r sql.Result
		logx.WithContext(ctx).Errorf("sqlx.In in Delete(_), error: %v", err)
		return
	}
	rResult, err = dao.db.Exec(ctx, query, a...)

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

// update user_profile_photos set deleted = 1, date2 = 0 where user_id = :user_id and photo_id in (:id_list)
// DeleteTx
// TODO(@benqi): sqlmap
func (dao *UserProfilePhotosDAO) DeleteTx(tx *sqlx.Tx, user_id int64, id_list []int64) (rowsAffected int64, err error) {
	var (
		query   = "update user_profile_photos set deleted = 1, date2 = 0 where user_id = ? and photo_id in (?)"
		a       []interface{}
		rResult sql.Result
	)

	if len(id_list) == 0 {
		return
	}

	query, a, err = sqlx.In(query, user_id, id_list)
	if err != nil {
		// r sql.Result
		logx.WithContext(tx.Context()).Errorf("sqlx.In in Delete(_), error: %v", err)
		return
	}
	rResult, err = tx.Exec(query, a...)

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
