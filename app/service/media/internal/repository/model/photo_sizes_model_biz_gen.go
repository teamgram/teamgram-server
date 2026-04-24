/*
 * WARNING! All changes made in this file will be lost!
 *   Created from by 'dalgen'
 *
 * Copyright (c) 2026-present, The Teamgram Authors (https://teamgram.net).
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
)

var _ *sql.Result
var _ = fmt.Sprintf
var _ = strings.Join
var _ = errors.Is
var _ *sqlx.DB

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
		err = fmt.Errorf("photo_sizes.Insert named exec: %w", err)
		return
	}

	lastInsertId, err = r.LastInsertId()
	if err != nil {
		err = fmt.Errorf("photo_sizes.Insert last insert id: %w", err)
		return
	}
	rowsAffected, err = r.RowsAffected()
	if err != nil {
		err = fmt.Errorf("photo_sizes.Insert rows affected: %w", err)
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
		err = fmt.Errorf("photo_sizes.InsertTx named exec: %w", err)
		return
	}

	lastInsertId, err = r.LastInsertId()
	if err != nil {
		err = fmt.Errorf("photo_sizes.InsertTx last insert id: %w", err)
		return
	}
	rowsAffected, err = r.RowsAffected()
	if err != nil {
		err = fmt.Errorf("photo_sizes.InsertTx rows affected: %w", err)
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
		err = fmt.Errorf("photo_sizes.SelectListByPhotoSizeId: %w", err)
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
		err = fmt.Errorf("photo_sizes.SelectListByPhotoSizeIdWithCB: %w", err)
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
		err = fmt.Errorf("photo_sizes.SelectListByPhotoSizeIdList: %w", err)
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
		err = fmt.Errorf("photo_sizes.SelectListByPhotoSizeIdListWithCB: %w", err)
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
