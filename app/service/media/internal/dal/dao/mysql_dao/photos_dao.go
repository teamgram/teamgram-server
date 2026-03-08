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

	"github.com/teamgram/marmota/pkg/stores/sqlx"
	"github.com/teamgram/teamgram-server/app/service/media/internal/dal/dataobject"

	"github.com/zeromicro/go-zero/core/logx"
)

type PhotosDAO struct {
	db *sqlx.DB
}

func NewPhotosDAO(db *sqlx.DB) *PhotosDAO {
	return &PhotosDAO{
		db: db,
	}
}

// Insert
// insert into photos(photo_id, access_hash, has_stickers, dc_id, date2, has_video, input_file_name, ext) values (:photo_id, :access_hash, :has_stickers, :dc_id, :date2, :has_video, :input_file_name, :ext)
func (dao *PhotosDAO) Insert(ctx context.Context, do *dataobject.PhotosDO) (lastInsertId, rowsAffected int64, err error) {
	var (
		query string
		r     sql.Result
	)
	query = "insert into photos(photo_id, access_hash, has_stickers, dc_id, date2, has_video, input_file_name, ext) values (:photo_id, :access_hash, :has_stickers, :dc_id, :date2, :has_video, :input_file_name, :ext)"

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
// insert into photos(photo_id, access_hash, has_stickers, dc_id, date2, has_video, input_file_name, ext) values (:photo_id, :access_hash, :has_stickers, :dc_id, :date2, :has_video, :input_file_name, :ext)
func (dao *PhotosDAO) InsertTx(tx *sqlx.Tx, do *dataobject.PhotosDO) (lastInsertId, rowsAffected int64, err error) {
	var (
		query string
		r     sql.Result
	)
	query = "insert into photos(photo_id, access_hash, has_stickers, dc_id, date2, has_video, input_file_name, ext) values (:photo_id, :access_hash, :has_stickers, :dc_id, :date2, :has_video, :input_file_name, :ext)"

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

// SelectByPhotoId
// select id, photo_id, access_hash, has_stickers, dc_id, date2, has_video, input_file_name, ext from photos where photo_id = :photo_id limit 1
func (dao *PhotosDAO) SelectByPhotoId(ctx context.Context, photoId int64) (rValue *dataobject.PhotosDO, err error) {
	var (
		query string
		do    = &dataobject.PhotosDO{}
	)
	query = "select id, photo_id, access_hash, has_stickers, dc_id, date2, has_video, input_file_name, ext from photos where photo_id = ? limit 1"

	err = dao.db.QueryRowPartial(ctx, do, query, photoId)

	if err != nil {
		if !errors.Is(err, sqlx.ErrNotFound) {
			logx.WithContext(ctx).Errorf("queryx in SelectByPhotoId(_), error: %v", err)
			return
		} else {
			// not found not error, return nil, nil
			err = nil
		}
	} else {
		rValue = do
	}

	return
}
