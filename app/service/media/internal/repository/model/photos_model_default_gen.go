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
	"fmt"
	"strings"

	"github.com/teamgram/marmota/pkg/stores/sqlx"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	photosFieldNames          = builder.RawFieldNames(&Photos{})
	photosRows                = strings.Join(photosFieldNames, ",")
	photosRowsExpectAutoSet   = strings.Join(stringx.Remove(photosFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	photosRowsWithPlaceHolder = strings.Join(stringx.Remove(photosFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"
)

type (
	photosModel interface {
		Insert2(ctx context.Context, data *Photos) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*Photos, error)
		FindListByIdList(ctx context.Context, id ...int64) ([]Photos, error)
		Update2(ctx context.Context, data *Photos) error
		Delete2(ctx context.Context, id int64) error

		FindOneByPhotoId(ctx context.Context, photoId int64) (*Photos, error)
		FindListByPhotoIdList(ctx context.Context, photoId ...int64) ([]Photos, error)
	}

	defaultPhotosModel struct {
		db *sqlx.DB
	}

	Photos struct {
		Id            int64  `db:"id" json:"id"`
		PhotoId       int64  `db:"photo_id" json:"photo_id"`
		AccessHash    int64  `db:"access_hash" json:"access_hash"`
		HasStickers   bool   `db:"has_stickers" json:"has_stickers"`
		DcId          int32  `db:"dc_id" json:"dc_id"`
		Date2         int64  `db:"date2" json:"date2"`
		HasVideo      bool   `db:"has_video" json:"has_video"`
		SizeId        int64  `db:"size_id" json:"size_id"`
		VideoSizeId   int64  `db:"video_size_id" json:"video_size_id"`
		InputFileName string `db:"input_file_name" json:"input_file_name"`
		Ext           string `db:"ext" json:"ext"`
	}
)

func newPhotosModel(db *sqlx.DB) *defaultPhotosModel {
	return &defaultPhotosModel{
		db: db,
	}
}

func (m *defaultPhotosModel) Insert2(ctx context.Context, data *Photos) (sql.Result, error) {
	query := fmt.Sprintf("insert into `photos` (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", photosRowsExpectAutoSet)

	return m.db.Exec(ctx, query, data.PhotoId, data.AccessHash, data.HasStickers, data.DcId, data.Date2, data.HasVideo, data.SizeId, data.VideoSizeId, data.InputFileName, data.Ext)

}

func (m *defaultPhotosModel) Delete2(ctx context.Context, id int64) error {
	query := "delete from `photos` where `id` = ?"

	_, err := m.db.Exec(ctx, query, id)
	return err
}

func (m *defaultPhotosModel) FindOne(ctx context.Context, id int64) (*Photos, error) {
	query := fmt.Sprintf("select %s from photos where id = ? limit 1", photosRows)
	var resp Photos

	err := m.db.QueryRowPartial(ctx, &resp, query, id)

	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (m *defaultPhotosModel) FindListByIdList(ctx context.Context, id ...int64) ([]Photos, error) {
	if len(id) == 0 {
		return []Photos{}, nil
	}

	query := fmt.Sprintf("select %s from photos where id in (%s)", photosRows, sqlx.InInt64List(id))

	var resp []Photos
	err := m.db.QueryRowsPartial(ctx, &resp, query)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (m *defaultPhotosModel) Update2(ctx context.Context, data *Photos) error {
	query := fmt.Sprintf("update `photos` set %s where `id` = ?", photosRowsWithPlaceHolder)

	_, err := m.db.Exec(ctx, query, data.PhotoId, data.AccessHash, data.HasStickers, data.DcId, data.Date2, data.HasVideo, data.SizeId, data.VideoSizeId, data.InputFileName, data.Ext, data.Id)
	return err
}

func (m *defaultPhotosModel) FindOneByPhotoId(ctx context.Context, photoId int64) (*Photos, error) {
	query := fmt.Sprintf("select %s from photos where photo_id = ? limit 1", photosRows)
	var resp Photos

	err := m.db.QueryRowPartial(ctx, &resp, query, photoId)

	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (m *defaultPhotosModel) FindListByPhotoIdList(ctx context.Context, photoId ...int64) ([]Photos, error) {
	if len(photoId) == 0 {
		return []Photos{}, nil
	}

	query := fmt.Sprintf("select %s from photos where photo_id in (%s)", photosRows, sqlx.InInt64List(photoId))

	var resp []Photos
	err := m.db.QueryRowsPartial(ctx, &resp, query)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
