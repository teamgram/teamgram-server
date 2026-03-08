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
	"fmt"

	"github.com/teamgram/marmota/pkg/stores/sqlx"
	"github.com/teamgram/teamgram-server/app/service/media/internal/dal/dataobject"

	"github.com/zeromicro/go-zero/core/logx"
)

type VideoSizesDAO struct {
	db *sqlx.DB
}

func NewVideoSizesDAO(db *sqlx.DB) *VideoSizesDAO {
	return &VideoSizesDAO{
		db: db,
	}
}

// Insert
// insert into video_sizes(video_size_id, size_type, width, height, file_size, video_start_ts, file_path) values (:video_size_id, :size_type, :width, :height, :file_size, :video_start_ts, :file_path)
func (dao *VideoSizesDAO) Insert(ctx context.Context, do *dataobject.VideoSizesDO) (lastInsertId, rowsAffected int64, err error) {
	var (
		query string
		r     sql.Result
	)
	query = "insert into video_sizes(video_size_id, size_type, width, height, file_size, video_start_ts, file_path) values (:video_size_id, :size_type, :width, :height, :file_size, :video_start_ts, :file_path)"

	r, err = dao.db.NamedExec(ctx, query, do)
	if err != nil {
		logx.WithContext(ctx).Errorf("namedExec in Insert(%v), error: %v", do, err)
		return
	}

	lastInsertId, err = r.LastInsertId()
	if err != nil {
		logx.WithContext(ctx).Errorf("lastInsertId in Insert(%v), error: %v", do, err)
		return
	}
	rowsAffected, err = r.RowsAffected()
	if err != nil {
		logx.WithContext(ctx).Errorf("rowsAffected in Insert(%v), error: %v", do, err)
	}

	return
}

// InsertTx
// insert into video_sizes(video_size_id, size_type, width, height, file_size, video_start_ts, file_path) values (:video_size_id, :size_type, :width, :height, :file_size, :video_start_ts, :file_path)
func (dao *VideoSizesDAO) InsertTx(tx *sqlx.Tx, do *dataobject.VideoSizesDO) (lastInsertId, rowsAffected int64, err error) {
	var (
		query string
		r     sql.Result
	)
	query = "insert into video_sizes(video_size_id, size_type, width, height, file_size, video_start_ts, file_path) values (:video_size_id, :size_type, :width, :height, :file_size, :video_start_ts, :file_path)"

	r, err = tx.NamedExec(query, do)
	if err != nil {
		logx.WithContext(tx.Context()).Errorf("namedExec in Insert(%v), error: %v", do, err)
		return
	}

	lastInsertId, err = r.LastInsertId()
	if err != nil {
		logx.WithContext(tx.Context()).Errorf("lastInsertId in Insert(%v), error: %v", do, err)
		return
	}
	rowsAffected, err = r.RowsAffected()
	if err != nil {
		logx.WithContext(tx.Context()).Errorf("rowsAffected in Insert(%v), error: %v", do, err)
	}

	return
}

// SelectListByVideoSizeId
// select id, video_size_id, size_type, width, height, file_size, video_start_ts, file_path from video_sizes where video_size_id = :video_size_id order by id asc
func (dao *VideoSizesDAO) SelectListByVideoSizeId(ctx context.Context, videoSizeId int64) (rList []dataobject.VideoSizesDO, err error) {
	var (
		query  string
		values []dataobject.VideoSizesDO
	)
	query = "select id, video_size_id, size_type, width, height, file_size, video_start_ts, file_path from video_sizes where video_size_id = ? order by id asc"

	err = dao.db.QueryRowsPartial(ctx, &values, query, videoSizeId)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectListByVideoSizeId(_), error: %v", err)
		return
	}

	rList = values

	return
}

// SelectListByVideoSizeIdWithCB
// select id, video_size_id, size_type, width, height, file_size, video_start_ts, file_path from video_sizes where video_size_id = :video_size_id order by id asc
func (dao *VideoSizesDAO) SelectListByVideoSizeIdWithCB(ctx context.Context, videoSizeId int64, cb func(sz, i int, v *dataobject.VideoSizesDO)) (rList []dataobject.VideoSizesDO, err error) {
	var (
		query  string
		values []dataobject.VideoSizesDO
	)
	query = "select id, video_size_id, size_type, width, height, file_size, video_start_ts, file_path from video_sizes where video_size_id = ? order by id asc"

	err = dao.db.QueryRowsPartial(ctx, &values, query, videoSizeId)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectListByVideoSizeId(_), error: %v", err)
		return
	}

	rList = values

	if cb != nil {
		sz := len(rList)
		for i := range sz {
			cb(sz, i, &rList[i])
		}
	}

	return
}

// SelectListByVideoSizeIdList
// select id, video_size_id, size_type, width, height, file_size, video_start_ts, file_path from video_sizes where video_size_id in (:idList) order by id asc
func (dao *VideoSizesDAO) SelectListByVideoSizeIdList(ctx context.Context, idList []int64) (rList []dataobject.VideoSizesDO, err error) {
	if len(idList) == 0 {
		rList = []dataobject.VideoSizesDO{}
		return
	}

	var (
		query  string
		values []dataobject.VideoSizesDO
	)
	query = fmt.Sprintf("select id, video_size_id, size_type, width, height, file_size, video_start_ts, file_path from video_sizes where video_size_id in (%s) order by id asc", sqlx.InInt64List(idList))

	err = dao.db.QueryRowsPartial(ctx, &values, query)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectListByVideoSizeIdList(_), error: %v", err)
		return
	}

	rList = values

	return
}

// SelectListByVideoSizeIdListWithCB
// select id, video_size_id, size_type, width, height, file_size, video_start_ts, file_path from video_sizes where video_size_id in (:idList) order by id asc
func (dao *VideoSizesDAO) SelectListByVideoSizeIdListWithCB(ctx context.Context, idList []int64, cb func(sz, i int, v *dataobject.VideoSizesDO)) (rList []dataobject.VideoSizesDO, err error) {
	if len(idList) == 0 {
		rList = []dataobject.VideoSizesDO{}
		return
	}

	var (
		query  string
		values []dataobject.VideoSizesDO
	)
	query = fmt.Sprintf("select id, video_size_id, size_type, width, height, file_size, video_start_ts, file_path from video_sizes where video_size_id in (%s) order by id asc", sqlx.InInt64List(idList))

	err = dao.db.QueryRowsPartial(ctx, &values, query)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectListByVideoSizeIdList(_), error: %v", err)
		return
	}

	rList = values

	if cb != nil {
		sz := len(rList)
		for i := range sz {
			cb(sz, i, &rList[i])
		}
	}

	return
}
