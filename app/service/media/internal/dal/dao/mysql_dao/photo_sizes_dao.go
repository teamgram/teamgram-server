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
	"github.com/teamgram/teamgram-server/app/service/media/internal/dal/dataobject"

	"github.com/zeromicro/go-zero/core/logx"
)

var _ *sql.Result

type PhotoSizesDAO struct {
	db *sqlx.DB
}

func NewPhotoSizesDAO(db *sqlx.DB) *PhotoSizesDAO {
	return &PhotoSizesDAO{db}
}

// Insert
// insert into photo_sizes(photo_size_id, size_type, width, height, file_size, file_path, cached_type, cached_bytes) values (:photo_size_id, :size_type, :width, :height, :file_size, :file_path, :cached_type, :cached_bytes)
func (dao *PhotoSizesDAO) Insert(ctx context.Context, do *dataobject.PhotoSizesDO) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into photo_sizes(photo_size_id, size_type, width, height, file_size, file_path, cached_type, cached_bytes) values (:photo_size_id, :size_type, :width, :height, :file_size, :file_path, :cached_type, :cached_bytes)"
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
// insert into photo_sizes(photo_size_id, size_type, width, height, file_size, file_path, cached_type, cached_bytes) values (:photo_size_id, :size_type, :width, :height, :file_size, :file_path, :cached_type, :cached_bytes)
func (dao *PhotoSizesDAO) InsertTx(tx *sqlx.Tx, do *dataobject.PhotoSizesDO) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into photo_sizes(photo_size_id, size_type, width, height, file_size, file_path, cached_type, cached_bytes) values (:photo_size_id, :size_type, :width, :height, :file_size, :file_path, :cached_type, :cached_bytes)"
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

// SelectListByPhotoSizeId
// select id, photo_size_id, size_type, width, height, file_size, file_path, cached_type, cached_bytes from photo_sizes where photo_size_id = :photo_size_id order by id asc
func (dao *PhotoSizesDAO) SelectListByPhotoSizeId(ctx context.Context, photo_size_id int64) (rList []dataobject.PhotoSizesDO, err error) {
	var (
		query = "select id, photo_size_id, size_type, width, height, file_size, file_path, cached_type, cached_bytes from photo_sizes where photo_size_id = ? order by id asc"
		rows  *sqlx.Rows
	)
	rows, err = dao.db.Query(ctx, query, photo_size_id)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectListByPhotoSizeId(_), error: %v", err)
		return
	}

	defer rows.Close()

	var values []dataobject.PhotoSizesDO
	for rows.Next() {
		v := dataobject.PhotoSizesDO{}

		// TODO(@benqi): not use reflect
		err = rows.StructScan(&v)
		if err != nil {
			logx.WithContext(ctx).Errorf("structScan in SelectListByPhotoSizeId(_), error: %v", err)
			return
		}
		values = append(values, v)
	}
	rList = values

	return
}

// SelectListByPhotoSizeIdWithCB
// select id, photo_size_id, size_type, width, height, file_size, file_path, cached_type, cached_bytes from photo_sizes where photo_size_id = :photo_size_id order by id asc
func (dao *PhotoSizesDAO) SelectListByPhotoSizeIdWithCB(ctx context.Context, photo_size_id int64, cb func(i int, v *dataobject.PhotoSizesDO)) (rList []dataobject.PhotoSizesDO, err error) {
	var (
		query = "select id, photo_size_id, size_type, width, height, file_size, file_path, cached_type, cached_bytes from photo_sizes where photo_size_id = ? order by id asc"
		rows  *sqlx.Rows
	)
	rows, err = dao.db.Query(ctx, query, photo_size_id)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectListByPhotoSizeId(_), error: %v", err)
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

	var values []dataobject.PhotoSizesDO
	for rows.Next() {
		v := dataobject.PhotoSizesDO{}

		// TODO(@benqi): not use reflect
		err = rows.StructScan(&v)
		if err != nil {
			logx.WithContext(ctx).Errorf("structScan in SelectListByPhotoSizeId(_), error: %v", err)
			return
		}
		values = append(values, v)
	}
	rList = values

	return
}

// SelectListByPhotoSizeIdList
// select id, photo_size_id, size_type, width, height, file_size, file_path, cached_type, cached_bytes from photo_sizes where photo_size_id in (:idList) order by id asc
func (dao *PhotoSizesDAO) SelectListByPhotoSizeIdList(ctx context.Context, idList []int64) (rList []dataobject.PhotoSizesDO, err error) {
	var (
		query = "select id, photo_size_id, size_type, width, height, file_size, file_path, cached_type, cached_bytes from photo_sizes where photo_size_id in (?) order by id asc"
		a     []interface{}
		rows  *sqlx.Rows
	)
	if len(idList) == 0 {
		rList = []dataobject.PhotoSizesDO{}
		return
	}

	query, a, err = sqlx.In(query, idList)
	if err != nil {
		// r sql.Result
		logx.WithContext(ctx).Errorf("sqlx.In in SelectListByPhotoSizeIdList(_), error: %v", err)
		return
	}
	rows, err = dao.db.Query(ctx, query, a...)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectListByPhotoSizeIdList(_), error: %v", err)
		return
	}

	defer rows.Close()

	var values []dataobject.PhotoSizesDO
	for rows.Next() {
		v := dataobject.PhotoSizesDO{}

		// TODO(@benqi): not use reflect
		err = rows.StructScan(&v)
		if err != nil {
			logx.WithContext(ctx).Errorf("structScan in SelectListByPhotoSizeIdList(_), error: %v", err)
			return
		}
		values = append(values, v)
	}
	rList = values

	return
}

// SelectListByPhotoSizeIdListWithCB
// select id, photo_size_id, size_type, width, height, file_size, file_path, cached_type, cached_bytes from photo_sizes where photo_size_id in (:idList) order by id asc
func (dao *PhotoSizesDAO) SelectListByPhotoSizeIdListWithCB(ctx context.Context, idList []int64, cb func(i int, v *dataobject.PhotoSizesDO)) (rList []dataobject.PhotoSizesDO, err error) {
	var (
		query = "select id, photo_size_id, size_type, width, height, file_size, file_path, cached_type, cached_bytes from photo_sizes where photo_size_id in (?) order by id asc"
		a     []interface{}
		rows  *sqlx.Rows
	)
	if len(idList) == 0 {
		rList = []dataobject.PhotoSizesDO{}
		return
	}

	query, a, err = sqlx.In(query, idList)
	if err != nil {
		// r sql.Result
		logx.WithContext(ctx).Errorf("sqlx.In in SelectListByPhotoSizeIdList(_), error: %v", err)
		return
	}
	rows, err = dao.db.Query(ctx, query, a...)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectListByPhotoSizeIdList(_), error: %v", err)
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

	var values []dataobject.PhotoSizesDO
	for rows.Next() {
		v := dataobject.PhotoSizesDO{}

		// TODO(@benqi): not use reflect
		err = rows.StructScan(&v)
		if err != nil {
			logx.WithContext(ctx).Errorf("structScan in SelectListByPhotoSizeIdList(_), error: %v", err)
			return
		}
		values = append(values, v)
	}
	rList = values

	return
}
