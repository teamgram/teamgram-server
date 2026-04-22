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
	userProfilePhotosFieldNames          = builder.RawFieldNames(&UserProfilePhotos{})
	userProfilePhotosRows                = strings.Join(userProfilePhotosFieldNames, ",")
	userProfilePhotosRowsExpectAutoSet   = strings.Join(stringx.Remove(userProfilePhotosFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	userProfilePhotosRowsWithPlaceHolder = strings.Join(stringx.Remove(userProfilePhotosFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"

	cacheTUserProfilePhotosIdPrefix = "cache:t:user_profile_photos:id:"

	cacheUserProfilePhotosIdPrefix = "cache#UserProfilePhotos#id"

	cacheUserProfilePhotosUserIdPhotoIdPrefix = "cache#UserId#PhotoId"
)

type (
	userProfilePhotosModel interface {
		Insert2(ctx context.Context, data *UserProfilePhotos) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*UserProfilePhotos, error)
		FindListByIdList(ctx context.Context, id ...int64) ([]UserProfilePhotos, error)
		Update2(ctx context.Context, data *UserProfilePhotos) error
		Delete2(ctx context.Context, id int64) error

		FindOneByUserIdPhotoId(ctx context.Context, userId int64, photoId int64) (*UserProfilePhotos, error)
	}

	defaultUserProfilePhotosModel struct {
		db *sqlx.DB
	}

	UserProfilePhotos struct {
		Id      int64 `db:"id" json:"id"`
		UserId  int64 `db:"user_id" json:"user_id"`
		PhotoId int64 `db:"photo_id" json:"photo_id"`
		Date2   int64 `db:"date2" json:"date2"`
		Deleted bool  `db:"deleted" json:"deleted"`
	}
)

func newUserProfilePhotosModel(db *sqlx.DB) *defaultUserProfilePhotosModel {
	return &defaultUserProfilePhotosModel{
		db: db,
	}
}

func (m *defaultUserProfilePhotosModel) Insert2(ctx context.Context, data *UserProfilePhotos) (sql.Result, error) {
	query := fmt.Sprintf("insert into `user_profile_photos` (%s) values (?, ?, ?, ?)", userProfilePhotosRowsExpectAutoSet)
	return m.db.Exec(ctx, query, data.UserId, data.PhotoId, data.Date2, data.Deleted)
}

func (m *defaultUserProfilePhotosModel) Delete2(ctx context.Context, id int64) error {
	query := "delete from `user_profile_photos` where `id` = ?"
	_, err := m.db.Exec(ctx, query, id)
	return err
}

func (m *defaultUserProfilePhotosModel) FindOne(ctx context.Context, id int64) (*UserProfilePhotos, error) {
	query := fmt.Sprintf("select %s from user_profile_photos where id = ? limit 1", userProfilePhotosRows)
	var resp UserProfilePhotos
	err := m.db.QueryRowPartial(ctx, &resp, query, id)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (m *defaultUserProfilePhotosModel) FindListByIdList(ctx context.Context, id ...int64) ([]UserProfilePhotos, error) {
	if len(id) == 0 {
		return []UserProfilePhotos{}, nil
	}

	query := fmt.Sprintf("select %s from user_profile_photos where id in (%s)", userProfilePhotosRows, sqlx.InInt64List(id))

	var resp []UserProfilePhotos
	err := m.db.QueryRowsPartial(ctx, &resp, query)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (m *defaultUserProfilePhotosModel) Update2(ctx context.Context, data *UserProfilePhotos) error {
	query := fmt.Sprintf("update `user_profile_photos` set %s where `id` = ?", userProfilePhotosRowsWithPlaceHolder)
	_, err := m.db.Exec(ctx, query, data.UserId, data.PhotoId, data.Date2, data.Deleted, data.Id)
	return err
}

func (m *defaultUserProfilePhotosModel) FindOneByUserIdPhotoId(ctx context.Context, userId int64, photoId int64) (*UserProfilePhotos, error) {
	query := fmt.Sprintf("select %s from user_profile_photos where user_id = ? AND photo_id = ? limit 1", userProfilePhotosRows)
	var resp UserProfilePhotos
	err := m.db.QueryRowPartial(ctx, &resp, query, userId, photoId)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (m *defaultUserProfilePhotosModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s#%v", cacheUserProfilePhotosIdPrefix, primary)
}

func (m *defaultUserProfilePhotosModel) queryPrimary(ctx context.Context, v interface{}, primary interface{}) error {
	query := fmt.Sprintf("select %s from user_profile_photos where id = ? limit 1", userProfilePhotosRows)
	return m.db.QueryRowPartial(ctx, v, query, primary)
}
