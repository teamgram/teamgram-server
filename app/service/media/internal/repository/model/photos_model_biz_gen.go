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
	"github.com/zeromicro/go-zero/core/logx"
)

var _ *sql.Result
var _ = fmt.Sprintf
var _ = strings.Join
var _ = errors.Is
var _ *sqlx.DB
var _ *logx.Logger

type (
	bizPhotosModel interface {
		Insert(ctx context.Context, data *Photos) (lastInsertId, rowsAffected int64, err error)
		InsertTx(tx *sqlx.Tx, data *Photos) (lastInsertId, rowsAffected int64, err error)

		SelectByPhotoId(ctx context.Context, photoId int64) (*Photos, error)
	}
)

// Insert
// insert into photos(photo_id, access_hash, has_stickers, dc_id, date2, has_video, input_file_name, ext) values (:photo_id, :access_hash, :has_stickers, :dc_id, :date2, :has_video, :input_file_name, :ext)
func (m *defaultPhotosModel) Insert(ctx context.Context, data *Photos) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into photos(photo_id, access_hash, has_stickers, dc_id, date2, has_video, input_file_name, ext) values (:photo_id, :access_hash, :has_stickers, :dc_id, :date2, :has_video, :input_file_name, :ext)"
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
// insert into photos(photo_id, access_hash, has_stickers, dc_id, date2, has_video, input_file_name, ext) values (:photo_id, :access_hash, :has_stickers, :dc_id, :date2, :has_video, :input_file_name, :ext)
func (m *defaultPhotosModel) InsertTx(tx *sqlx.Tx, data *Photos) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into photos(photo_id, access_hash, has_stickers, dc_id, date2, has_video, input_file_name, ext) values (:photo_id, :access_hash, :has_stickers, :dc_id, :date2, :has_video, :input_file_name, :ext)"
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

// SelectByPhotoId
// select id, photo_id, access_hash, has_stickers, dc_id, date2, has_video, input_file_name, ext from photos where photo_id = :photo_id limit 1
func (m *defaultPhotosModel) SelectByPhotoId(ctx context.Context, photoId int64) (rValue *Photos, err error) {

	var (
		query = "select id, photo_id, access_hash, has_stickers, dc_id, date2, has_video, input_file_name, ext from photos where photo_id = ? limit 1"
		do    = &Photos{}
	)
	err = m.db.QueryRowPartial(ctx, do, query, photoId)

	if err != nil {
		if !errors.Is(err, sqlx.ErrNotFound) {
			logx.WithContext(ctx).Errorf("queryx in SelectByPhotoId(_), error: %v", err)
			return
		} else {
			err = nil
		}
	} else {
		rValue = do
	}

	return
}
