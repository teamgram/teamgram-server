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
	bizVideoSizesModel interface {
		Insert(ctx context.Context, data *VideoSizes) (lastInsertId, rowsAffected int64, err error)
		InsertTx(tx *sqlx.Tx, data *VideoSizes) (lastInsertId, rowsAffected int64, err error)

		SelectListByVideoSizeId(ctx context.Context, videoSizeId int64) ([]VideoSizes, error)
		SelectListByVideoSizeIdWithCB(ctx context.Context, videoSizeId int64, cb func(sz, i int, v *VideoSizes)) ([]VideoSizes, error)

		SelectListByVideoSizeIdList(ctx context.Context, idList []int64) ([]VideoSizes, error)
		SelectListByVideoSizeIdListWithCB(ctx context.Context, idList []int64, cb func(sz, i int, v *VideoSizes)) ([]VideoSizes, error)
	}
)

// Insert
// insert into video_sizes(video_size_id, size_type, width, height, file_size, video_start_ts, file_path) values (:video_size_id, :size_type, :width, :height, :file_size, :video_start_ts, :file_path)
func (m *defaultVideoSizesModel) Insert(ctx context.Context, data *VideoSizes) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into video_sizes(video_size_id, size_type, width, height, file_size, video_start_ts, file_path) values (:video_size_id, :size_type, :width, :height, :file_size, :video_start_ts, :file_path)"
		r     sql.Result
	)

	r, err = m.db.NamedExec(ctx, query, data)
	if err != nil {
		err = fmt.Errorf("video_sizes.Insert named exec: %w", err)
		return
	}

	lastInsertId, err = r.LastInsertId()
	if err != nil {
		err = fmt.Errorf("video_sizes.Insert last insert id: %w", err)
		return
	}
	rowsAffected, err = r.RowsAffected()
	if err != nil {
		err = fmt.Errorf("video_sizes.Insert rows affected: %w", err)
	}

	return

}

// InsertTx
// insert into video_sizes(video_size_id, size_type, width, height, file_size, video_start_ts, file_path) values (:video_size_id, :size_type, :width, :height, :file_size, :video_start_ts, :file_path)
func (m *defaultVideoSizesModel) InsertTx(tx *sqlx.Tx, data *VideoSizes) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into video_sizes(video_size_id, size_type, width, height, file_size, video_start_ts, file_path) values (:video_size_id, :size_type, :width, :height, :file_size, :video_start_ts, :file_path)"
		r     sql.Result
	)

	r, err = tx.NamedExec(query, data)
	if err != nil {
		err = fmt.Errorf("video_sizes.InsertTx named exec: %w", err)
		return
	}

	lastInsertId, err = r.LastInsertId()
	if err != nil {
		err = fmt.Errorf("video_sizes.InsertTx last insert id: %w", err)
		return
	}
	rowsAffected, err = r.RowsAffected()
	if err != nil {
		err = fmt.Errorf("video_sizes.InsertTx rows affected: %w", err)
	}

	return
}

// SelectListByVideoSizeId
// select id, video_size_id, size_type, width, height, file_size, video_start_ts, file_path from video_sizes where video_size_id = :video_size_id order by id asc
func (m *defaultVideoSizesModel) SelectListByVideoSizeId(ctx context.Context, videoSizeId int64) (rList []VideoSizes, err error) {
	var (
		query  = "select id, video_size_id, size_type, width, height, file_size, video_start_ts, file_path from video_sizes where video_size_id = ? order by id asc"
		values []VideoSizes
	)
	err = m.db.QueryRowsPartial(ctx, &values, query, videoSizeId)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			rList = []VideoSizes{}
			err = nil
			return
		}
		err = fmt.Errorf("video_sizes.SelectListByVideoSizeId: %w", err)
		return
	}

	rList = values

	return
}

// SelectListByVideoSizeIdWithCB
// select id, video_size_id, size_type, width, height, file_size, video_start_ts, file_path from video_sizes where video_size_id = :video_size_id order by id asc
func (m *defaultVideoSizesModel) SelectListByVideoSizeIdWithCB(ctx context.Context, videoSizeId int64, cb func(sz, i int, v *VideoSizes)) (rList []VideoSizes, err error) {
	var (
		query  = "select id, video_size_id, size_type, width, height, file_size, video_start_ts, file_path from video_sizes where video_size_id = ? order by id asc"
		values []VideoSizes
	)
	err = m.db.QueryRowsPartial(ctx, &values, query, videoSizeId)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			rList = []VideoSizes{}
			err = nil
			return
		}
		err = fmt.Errorf("video_sizes.SelectListByVideoSizeIdWithCB: %w", err)
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

// SelectListByVideoSizeIdList
// select id, video_size_id, size_type, width, height, file_size, video_start_ts, file_path from video_sizes where video_size_id in (:idList) order by id asc
func (m *defaultVideoSizesModel) SelectListByVideoSizeIdList(ctx context.Context, idList []int64) (rList []VideoSizes, err error) {
	var (
		query  = fmt.Sprintf("select id, video_size_id, size_type, width, height, file_size, video_start_ts, file_path from video_sizes where video_size_id in (%s) order by id asc", sqlx.InInt64List(idList))
		values []VideoSizes
	)
	if len(idList) == 0 {
		rList = []VideoSizes{}
		return
	}

	err = m.db.QueryRowsPartial(ctx, &values, query)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			rList = []VideoSizes{}
			err = nil
			return
		}
		err = fmt.Errorf("video_sizes.SelectListByVideoSizeIdList: %w", err)
		return
	}

	rList = values

	return
}

// SelectListByVideoSizeIdListWithCB
// select id, video_size_id, size_type, width, height, file_size, video_start_ts, file_path from video_sizes where video_size_id in (:idList) order by id asc
func (m *defaultVideoSizesModel) SelectListByVideoSizeIdListWithCB(ctx context.Context, idList []int64, cb func(sz, i int, v *VideoSizes)) (rList []VideoSizes, err error) {
	var (
		query  = fmt.Sprintf("select id, video_size_id, size_type, width, height, file_size, video_start_ts, file_path from video_sizes where video_size_id in (%s) order by id asc", sqlx.InInt64List(idList))
		values []VideoSizes
	)
	if len(idList) == 0 {
		rList = []VideoSizes{}
		return
	}

	err = m.db.QueryRowsPartial(ctx, &values, query)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			rList = []VideoSizes{}
			err = nil
			return
		}
		err = fmt.Errorf("video_sizes.SelectListByVideoSizeIdListWithCB: %w", err)
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
