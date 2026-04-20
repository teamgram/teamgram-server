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
	bizPhotoSizesModel interface {
		Insert(ctx context.Context, data *PhotoSizes) (lastInsertId, rowsAffected int64, err error)
		InsertTx(tx *sqlx.Tx, data *PhotoSizes) (lastInsertId, rowsAffected int64, err error)

		SelectListByPhotoSizeId(ctx context.Context, photoSizeId int64) ([]PhotoSizes, error)
		SelectListByPhotoSizeIdWithCB(ctx context.Context, photoSizeId int64, cb func(sz, i int, v *PhotoSizes)) ([]PhotoSizes, error)

		SelectListByPhotoSizeIdList(ctx context.Context, idList []int64) ([]PhotoSizes, error)
		SelectListByPhotoSizeIdListWithCB(ctx context.Context, idList []int64, cb func(sz, i int, v *PhotoSizes)) ([]PhotoSizes, error)
	}
)

// Insert
// insert into photo_sizes(photo_size_id, size_type, width, height, file_size, file_path, cached_type, cached_bytes) values (:photo_size_id, :size_type, :width, :height, :file_size, :file_path, :cached_type, :cached_bytes)
func (m *defaultPhotoSizesModel) Insert(ctx context.Context, data *PhotoSizes) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into photo_sizes(photo_size_id, size_type, width, height, file_size, file_path, cached_type, cached_bytes) values (:photo_size_id, :size_type, :width, :height, :file_size, :file_path, :cached_type, :cached_bytes)"
		r     sql.Result
	)

	r, err = m.db.NamedExec(ctx, query, data)
	if err != nil {
		logx.WithContext(ctx).Errorf("namedExec in Insert(%v), error: %v", data, err)
		return
	}

	lastInsertId, err = r.LastInsertId()
	if err != nil {
		logx.WithContext(ctx).Errorf("lastInsertId in Insert(%v)_error: %v", data, err)
		return
	}
	rowsAffected, err = r.RowsAffected()
	if err != nil {
		logx.WithContext(ctx).Errorf("rowsAffected in Insert(%v)_error: %v", data, err)
	}

	return
}

// InsertTx
// insert into photo_sizes(photo_size_id, size_type, width, height, file_size, file_path, cached_type, cached_bytes) values (:photo_size_id, :size_type, :width, :height, :file_size, :file_path, :cached_type, :cached_bytes)
func (m *defaultPhotoSizesModel) InsertTx(tx *sqlx.Tx, data *PhotoSizes) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into photo_sizes(photo_size_id, size_type, width, height, file_size, file_path, cached_type, cached_bytes) values (:photo_size_id, :size_type, :width, :height, :file_size, :file_path, :cached_type, :cached_bytes)"
		r     sql.Result
	)

	r, err = tx.NamedExec(query, data)
	if err != nil {
		logx.WithContext(tx.Context()).Errorf("namedExec in Insert(%v), error: %v", data, err)
		return
	}

	lastInsertId, err = r.LastInsertId()
	if err != nil {
		logx.WithContext(tx.Context()).Errorf("lastInsertId in Insert(%v)_error: %v", data, err)
		return
	}
	rowsAffected, err = r.RowsAffected()
	if err != nil {
		logx.WithContext(tx.Context()).Errorf("rowsAffected in Insert(%v)_error: %v", data, err)
	}

	return
}

// SelectListByPhotoSizeId
// select id, photo_size_id, size_type, width, height, file_size, file_path, cached_type, cached_bytes from photo_sizes where photo_size_id = :photo_size_id order by id asc
func (m *defaultPhotoSizesModel) SelectListByPhotoSizeId(ctx context.Context, photoSizeId int64) (rList []PhotoSizes, err error) {
	var (
		query  = "select id, photo_size_id, size_type, width, height, file_size, file_path, cached_type, cached_bytes from photo_sizes where photo_size_id = ? order by id asc"
		values []PhotoSizes
	)
	err = m.db.QueryRowsPartial(ctx, &values, query, photoSizeId)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectListByPhotoSizeId(_), error: %v", err)
		return
	}

	rList = values

	return
}

// SelectListByPhotoSizeIdWithCB
// select id, photo_size_id, size_type, width, height, file_size, file_path, cached_type, cached_bytes from photo_sizes where photo_size_id = :photo_size_id order by id asc
func (m *defaultPhotoSizesModel) SelectListByPhotoSizeIdWithCB(ctx context.Context, photoSizeId int64, cb func(sz, i int, v *PhotoSizes)) (rList []PhotoSizes, err error) {
	var (
		query  = "select id, photo_size_id, size_type, width, height, file_size, file_path, cached_type, cached_bytes from photo_sizes where photo_size_id = ? order by id asc"
		values []PhotoSizes
	)
	err = m.db.QueryRowsPartial(ctx, &values, query, photoSizeId)

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
func (m *defaultPhotoSizesModel) SelectListByPhotoSizeIdList(ctx context.Context, idList []int64) (rList []PhotoSizes, err error) {
	var (
		query  = fmt.Sprintf("select id, photo_size_id, size_type, width, height, file_size, file_path, cached_type, cached_bytes from photo_sizes where photo_size_id in (%s) order by id asc", sqlx.InInt64List(idList))
		values []PhotoSizes
	)
	if len(idList) == 0 {
		rList = []PhotoSizes{}
		return
	}

	err = m.db.QueryRowsPartial(ctx, &values, query)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectListByPhotoSizeIdList(_), error: %v", err)
		return
	}

	rList = values

	return
}

// SelectListByPhotoSizeIdListWithCB
// select id, photo_size_id, size_type, width, height, file_size, file_path, cached_type, cached_bytes from photo_sizes where photo_size_id in (:idList) order by id asc
func (m *defaultPhotoSizesModel) SelectListByPhotoSizeIdListWithCB(ctx context.Context, idList []int64, cb func(sz, i int, v *PhotoSizes)) (rList []PhotoSizes, err error) {
	var (
		query  = fmt.Sprintf("select id, photo_size_id, size_type, width, height, file_size, file_path, cached_type, cached_bytes from photo_sizes where photo_size_id in (%s) order by id asc", sqlx.InInt64List(idList))
		values []PhotoSizes
	)
	if len(idList) == 0 {
		rList = []PhotoSizes{}
		return
	}

	err = m.db.QueryRowsPartial(ctx, &values, query)

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
