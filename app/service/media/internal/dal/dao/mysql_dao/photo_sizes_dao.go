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
	"github.com/teamgram/teamgram-server/app/service/media/internal/dal/dataobject"

	"github.com/zeromicro/go-zero/core/logx"
)

var _ *sql.Result
var _ = fmt.Sprintf
var _ = strings.Join
var _ = errors.Is

type PhotoSizesDAO struct {
	db *sqlx.DB
}

func NewPhotoSizesDAO(db *sqlx.DB) *PhotoSizesDAO {
	return &PhotoSizesDAO{
		db: db,
	}
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
func (dao *PhotoSizesDAO) SelectListByPhotoSizeId(ctx context.Context, photoSizeId int64) (rList []dataobject.PhotoSizesDO, err error) {
	var (
		query  = "select id, photo_size_id, size_type, width, height, file_size, file_path, cached_type, cached_bytes from photo_sizes where photo_size_id = ? order by id asc"
		values []dataobject.PhotoSizesDO
	)
	err = dao.db.QueryRowsPartial(ctx, &values, query, photoSizeId)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectListByPhotoSizeId(_), error: %v", err)
		return
	}

	rList = values

	return
}

// SelectListByPhotoSizeIdWithCB
// select id, photo_size_id, size_type, width, height, file_size, file_path, cached_type, cached_bytes from photo_sizes where photo_size_id = :photo_size_id order by id asc
func (dao *PhotoSizesDAO) SelectListByPhotoSizeIdWithCB(ctx context.Context, photoSizeId int64, cb func(sz, i int, v *dataobject.PhotoSizesDO)) (rList []dataobject.PhotoSizesDO, err error) {
	var (
		query  = "select id, photo_size_id, size_type, width, height, file_size, file_path, cached_type, cached_bytes from photo_sizes where photo_size_id = ? order by id asc"
		values []dataobject.PhotoSizesDO
	)
	err = dao.db.QueryRowsPartial(ctx, &values, query, photoSizeId)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectListByPhotoSizeId(_), error: %v", err)
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

// SelectListByPhotoSizeIdList
// select id, photo_size_id, size_type, width, height, file_size, file_path, cached_type, cached_bytes from photo_sizes where photo_size_id in (:idList) order by id asc
func (dao *PhotoSizesDAO) SelectListByPhotoSizeIdList(ctx context.Context, idList []int64) (rList []dataobject.PhotoSizesDO, err error) {
	var (
		query  = fmt.Sprintf("select id, photo_size_id, size_type, width, height, file_size, file_path, cached_type, cached_bytes from photo_sizes where photo_size_id in (%s) order by id asc", sqlx.InInt64List(idList))
		values []dataobject.PhotoSizesDO
	)
	if len(idList) == 0 {
		rList = []dataobject.PhotoSizesDO{}
		return
	}

	err = dao.db.QueryRowsPartial(ctx, &values, query)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectListByPhotoSizeIdList(_), error: %v", err)
		return
	}

	rList = values

	return
}

// SelectListByPhotoSizeIdListWithCB
// select id, photo_size_id, size_type, width, height, file_size, file_path, cached_type, cached_bytes from photo_sizes where photo_size_id in (:idList) order by id asc
func (dao *PhotoSizesDAO) SelectListByPhotoSizeIdListWithCB(ctx context.Context, idList []int64, cb func(sz, i int, v *dataobject.PhotoSizesDO)) (rList []dataobject.PhotoSizesDO, err error) {
	var (
		query  = fmt.Sprintf("select id, photo_size_id, size_type, width, height, file_size, file_path, cached_type, cached_bytes from photo_sizes where photo_size_id in (%s) order by id asc", sqlx.InInt64List(idList))
		values []dataobject.PhotoSizesDO
	)
	if len(idList) == 0 {
		rList = []dataobject.PhotoSizesDO{}
		return
	}

	err = dao.db.QueryRowsPartial(ctx, &values, query)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectListByPhotoSizeIdList(_), error: %v", err)
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
