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

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	videoSizesFieldNames          = builder.RawFieldNames(&VideoSizes{})
	videoSizesRows                = strings.Join(videoSizesFieldNames, ",")
	videoSizesRowsExpectAutoSet   = strings.Join(stringx.Remove(videoSizesFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	videoSizesRowsWithPlaceHolder = strings.Join(stringx.Remove(videoSizesFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"
)

type (
	videoSizesModel interface {
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
	tableName := "video_sizes"
	query := fmt.Sprintf("insert into `%s` (%s) values (?, ?, ?, ?, ?, ?, ?)", tableName, videoSizesRowsExpectAutoSet)

	r, err := m.db.Exec(ctx, query, data.VideoSizeId, data.SizeType, data.Width, data.Height, data.FileSize, data.VideoStartTs, data.FilePath)
	if err != nil {
		return nil, fmt.Errorf("video_sizes.Insert2 exec: %w", err)
	}

	return r, nil
}

func (m *defaultVideoSizesModel) Delete2(ctx context.Context, id int64) error {
	tableName := "video_sizes"
	query := fmt.Sprintf("delete from `%s` where `id` = ?", tableName)

	_, err := m.db.Exec(ctx, query, id)
	if err != nil {
		return fmt.Errorf("video_sizes.Delete2 exec: %w", err)
	}

	return nil
}

func (m *defaultVideoSizesModel) FindOne(ctx context.Context, id int64) (*VideoSizes, error) {
	tableName := "video_sizes"
	query := fmt.Sprintf("select %s from %s where id = ? limit 1", videoSizesRows, tableName)
	var resp VideoSizes

	err := m.db.QueryRowPartial(ctx, &resp, query, id)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return nil, &NotFoundError{
				Resource: "video_sizes",
				Key:      fmt.Sprintf("id=%v", id),
				Cause:    err,
			}
		}
		return nil, fmt.Errorf("video_sizes.FindOne: %w", err)
	}

	return &resp, nil
}

func (m *defaultVideoSizesModel) FindListByIdList(ctx context.Context, id ...int64) ([]VideoSizes, error) {
	if len(id) == 0 {
		return []VideoSizes{}, nil
	}
	tableName := "video_sizes"

	query := fmt.Sprintf("select %s from %s where id in (%s)", videoSizesRows, tableName, sqlx.InInt64List(id))

	var resp []VideoSizes
	err := m.db.QueryRowsPartial(ctx, &resp, query)
	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return []VideoSizes{}, nil
		}
		return nil, fmt.Errorf("video_sizes.FindListByIdList: %w", err)
	}

	return resp, nil
}

func (m *defaultVideoSizesModel) Update2(ctx context.Context, data *VideoSizes) error {
	tableName := "video_sizes"
	query := fmt.Sprintf("update `%s` set %s where `id` = ?", tableName, videoSizesRowsWithPlaceHolder)

	_, err := m.db.Exec(ctx, query, data.VideoSizeId, data.SizeType, data.Width, data.Height, data.FileSize, data.VideoStartTs, data.FilePath, data.Id)
	if err != nil {
		return fmt.Errorf("video_sizes.Update2 exec: %w", err)
	}

	return nil
}

func (m *defaultVideoSizesModel) FindOneByVideoSizeIdSizeType(ctx context.Context, videoSizeId int64, sizeType string) (*VideoSizes, error) {
	tableName := "video_sizes"
	query := fmt.Sprintf("select %s from %s where video_size_id = ? AND size_type = ? limit 1", videoSizesRows, tableName)
	var resp VideoSizes

	err := m.db.QueryRowPartial(ctx, &resp, query, videoSizeId, sizeType)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return nil, &NotFoundError{
				Resource: "video_sizes",
				Key:      fmt.Sprintf("video_size_id=%v,size_type=%v", videoSizeId, sizeType),
				Cause:    err,
			}
		}
		return nil, fmt.Errorf("video_sizes.FindOneByVideoSizeIdSizeType: %w", err)
	}

	return &resp, nil
}
