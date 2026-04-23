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
	photoSizesFieldNames          = builder.RawFieldNames(&PhotoSizes{})
	photoSizesRows                = strings.Join(photoSizesFieldNames, ",")
	photoSizesRowsExpectAutoSet   = strings.Join(stringx.Remove(photoSizesFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	photoSizesRowsWithPlaceHolder = strings.Join(stringx.Remove(photoSizesFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"
)

type (
	photoSizesModel interface {
		Insert2(ctx context.Context, data *PhotoSizes) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*PhotoSizes, error)
		FindListByIdList(ctx context.Context, id ...int64) ([]PhotoSizes, error)
		Update2(ctx context.Context, data *PhotoSizes) error
		Delete2(ctx context.Context, id int64) error

		FindOneByPhotoSizeIdSizeType(ctx context.Context, photoSizeId int64, sizeType string) (*PhotoSizes, error)
	}

	defaultPhotoSizesModel struct {
		db *sqlx.DB
	}

	PhotoSizes struct {
		Id            int64  `db:"id" json:"id"`
		PhotoSizeId   int64  `db:"photo_size_id" json:"photo_size_id"`
		SizeType      string `db:"size_type" json:"size_type"`
		VolumeId      int64  `db:"volume_id" json:"volume_id"`
		LocalId       int32  `db:"local_id" json:"local_id"`
		Secret        int64  `db:"secret" json:"secret"`
		Width         int32  `db:"width" json:"width"`
		Height        int32  `db:"height" json:"height"`
		FileSize      int32  `db:"file_size" json:"file_size"`
		FilePath      string `db:"file_path" json:"file_path"`
		HasStripped   bool   `db:"has_stripped" json:"has_stripped"`
		StrippedBytes string `db:"stripped_bytes" json:"stripped_bytes"`
		CachedType    int32  `db:"cached_type" json:"cached_type"`
		CachedBytes   string `db:"cached_bytes" json:"cached_bytes"`
	}
)

func newPhotoSizesModel(db *sqlx.DB) *defaultPhotoSizesModel {
	return &defaultPhotoSizesModel{
		db: db,
	}
}

func (m *defaultPhotoSizesModel) Insert2(ctx context.Context, data *PhotoSizes) (sql.Result, error) {
	query := fmt.Sprintf("insert into `photo_sizes` (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", photoSizesRowsExpectAutoSet)

	return m.db.Exec(ctx, query, data.PhotoSizeId, data.SizeType, data.VolumeId, data.LocalId, data.Secret, data.Width, data.Height, data.FileSize, data.FilePath, data.HasStripped, data.StrippedBytes, data.CachedType, data.CachedBytes)

}

func (m *defaultPhotoSizesModel) Delete2(ctx context.Context, id int64) error {
	query := "delete from `photo_sizes` where `id` = ?"

	_, err := m.db.Exec(ctx, query, id)
	return err
}

func (m *defaultPhotoSizesModel) FindOne(ctx context.Context, id int64) (*PhotoSizes, error) {
	query := fmt.Sprintf("select %s from photo_sizes where id = ? limit 1", photoSizesRows)
	var resp PhotoSizes

	err := m.db.QueryRowPartial(ctx, &resp, query, id)

	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (m *defaultPhotoSizesModel) FindListByIdList(ctx context.Context, id ...int64) ([]PhotoSizes, error) {
	if len(id) == 0 {
		return []PhotoSizes{}, nil
	}

	query := fmt.Sprintf("select %s from photo_sizes where id in (%s)", photoSizesRows, sqlx.InInt64List(id))

	var resp []PhotoSizes
	err := m.db.QueryRowsPartial(ctx, &resp, query)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (m *defaultPhotoSizesModel) Update2(ctx context.Context, data *PhotoSizes) error {
	query := fmt.Sprintf("update `photo_sizes` set %s where `id` = ?", photoSizesRowsWithPlaceHolder)

	_, err := m.db.Exec(ctx, query, data.PhotoSizeId, data.SizeType, data.VolumeId, data.LocalId, data.Secret, data.Width, data.Height, data.FileSize, data.FilePath, data.HasStripped, data.StrippedBytes, data.CachedType, data.CachedBytes, data.Id)
	return err
}

func (m *defaultPhotoSizesModel) FindOneByPhotoSizeIdSizeType(ctx context.Context, photoSizeId int64, sizeType string) (*PhotoSizes, error) {
	query := fmt.Sprintf("select %s from photo_sizes where photo_size_id = ? AND size_type = ? limit 1", photoSizesRows)
	var resp PhotoSizes

	err := m.db.QueryRowPartial(ctx, &resp, query, photoSizeId, sizeType)

	if err != nil {
		return nil, err
	}
	return &resp, nil
}
