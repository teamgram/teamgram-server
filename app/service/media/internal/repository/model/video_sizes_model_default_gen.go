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
	"fmt"
	"strings"

	"github.com/teamgram/marmota/pkg/stores/sqlx"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	video_sizesFieldNames          = builder.RawFieldNames(&VideoSizes{})
	video_sizesRows                = strings.Join(video_sizesFieldNames, ",")
	video_sizesRowsExpectAutoSet   = strings.Join(stringx.Remove(video_sizesFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	video_sizesRowsWithPlaceHolder = strings.Join(stringx.Remove(video_sizesFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"

	cacheTVideoSizesIdPrefix = "cache:t:video_sizes:id:"

	cacheVideoSizesIdPrefix = "cache#VideoSizes#id"

	cacheVideoSizesVideoSizeIdSizeTypePrefix = "cache#VideoSizeId#SizeType"
)

type (
	video_sizesModel interface {
		Insert2(ctx context.Context, data *VideoSizes) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*VideoSizes, error)
		FindListByIdList(ctx context.Context, id ...int64) ([]VideoSizes, error)
		Update2(ctx context.Context, data *VideoSizes) error
		Delete2(ctx context.Context, id int64) error

		FindOneByVideoSizeIdSizeType(ctx context.Context, videoSizeId int64, sizeType string) (*VideoSizes, error)
	}

	defaultVideoSizesModel struct {
		db *sqlx.DB
	}

	VideoSizes struct {
		Id           int64   `db:"id" json:"id"`
		VideoSizeId  int64   `db:"video_size_id" json:"video_size_id"`
		SizeType     string  `db:"size_type" json:"size_type"`
		Width        int32   `db:"width" json:"width"`
		Height       int32   `db:"height" json:"height"`
		FileSize     int32   `db:"file_size" json:"file_size"`
		VideoStartTs float64 `db:"video_start_ts" json:"video_start_ts"`
		FilePath     string  `db:"file_path" json:"file_path"`
	}
)

func newVideoSizesModel(db *sqlx.DB) *defaultVideoSizesModel {
	return &defaultVideoSizesModel{
		db: db,
	}
}

func (m *defaultVideoSizesModel) Insert2(ctx context.Context, data *VideoSizes) (sql.Result, error) {
	query := fmt.Sprintf("insert into `video_sizes` (%s) values (?, ?, ?, ?, ?, ?, ?)", video_sizesRowsExpectAutoSet)
	return m.db.Exec(ctx, query, data.VideoSizeId, data.SizeType, data.Width, data.Height, data.FileSize, data.VideoStartTs, data.FilePath)
}

func (m *defaultVideoSizesModel) Delete2(ctx context.Context, id int64) error {
	query := "delete from `video_sizes` where `id` = ?"
	_, err := m.db.Exec(ctx, query, id)
	return err
}

func (m *defaultVideoSizesModel) FindOne(ctx context.Context, id int64) (*VideoSizes, error) {
	query := fmt.Sprintf("select %s from video_sizes where id = ? limit 1", video_sizesRows)
	var resp VideoSizes
	err := m.db.QueryRowPartial(ctx, &resp, query, id)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (m *defaultVideoSizesModel) FindListByIdList(ctx context.Context, id ...int64) ([]VideoSizes, error) {
	if len(id) == 0 {
		return []VideoSizes{}, nil
	}

	query := fmt.Sprintf("select %s from video_sizes where id in (%s)", video_sizesRows, sqlx.InInt64List(id))

	var resp []VideoSizes
	err := m.db.QueryRowsPartial(ctx, &resp, query)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (m *defaultVideoSizesModel) Update2(ctx context.Context, data *VideoSizes) error {
	query := fmt.Sprintf("update `video_sizes` set %s where `id` = ?", video_sizesRowsWithPlaceHolder)
	_, err := m.db.Exec(ctx, query, data.VideoSizeId, data.SizeType, data.Width, data.Height, data.FileSize, data.VideoStartTs, data.FilePath, data.Id)
	return err
}

func (m *defaultVideoSizesModel) FindOneByVideoSizeIdSizeType(ctx context.Context, videoSizeId int64, sizeType string) (*VideoSizes, error) {
	query := fmt.Sprintf("select %s from video_sizes where video_size_id = ? AND size_type = ? limit 1", video_sizesRows)
	var resp VideoSizes
	err := m.db.QueryRowPartial(ctx, &resp, query, videoSizeId, sizeType)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (m *defaultVideoSizesModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s#%v", cacheVideoSizesIdPrefix, primary)
}

func (m *defaultVideoSizesModel) queryPrimary(ctx context.Context, v interface{}, primary interface{}) error {
	query := fmt.Sprintf("select %s from video_sizes where id = ? limit 1", video_sizesRows)
	return m.db.QueryRowPartial(ctx, v, query, primary)
}
