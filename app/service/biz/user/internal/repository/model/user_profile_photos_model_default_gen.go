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
	userProfilePhotosFieldNames          = builder.RawFieldNames(&UserProfilePhotos{})
	userProfilePhotosRows                = strings.Join(userProfilePhotosFieldNames, ",")
	userProfilePhotosRowsExpectAutoSet   = strings.Join(stringx.Remove(userProfilePhotosFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	userProfilePhotosRowsWithPlaceHolder = strings.Join(stringx.Remove(userProfilePhotosFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"
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

	r, err := m.db.Exec(ctx, query, data.UserId, data.PhotoId, data.Date2, data.Deleted)
	if err != nil {
		return nil, fmt.Errorf("user_profile_photos.Insert2 exec: %w", err)
	}

	return r, nil
}

func (m *defaultUserProfilePhotosModel) Delete2(ctx context.Context, id int64) error {
	query := "delete from `user_profile_photos` where `id` = ?"

	_, err := m.db.Exec(ctx, query, id)
	if err != nil {
		return fmt.Errorf("user_profile_photos.Delete2 exec: %w", err)
	}

	return nil
}

func (m *defaultUserProfilePhotosModel) FindOne(ctx context.Context, id int64) (*UserProfilePhotos, error) {
	query := fmt.Sprintf("select %s from user_profile_photos where id = ? limit 1", userProfilePhotosRows)
	var resp UserProfilePhotos

	err := m.db.QueryRowPartial(ctx, &resp, query, id)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return nil, &NotFoundError{
				Resource: "user_profile_photos",
				Key:      fmt.Sprintf("id=%v", id),
				Cause:    err,
			}
		}
		return nil, fmt.Errorf("user_profile_photos.FindOne: %w", err)
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
		if errors.Is(err, sqlx.ErrNotFound) {
			return []UserProfilePhotos{}, nil
		}
		return nil, fmt.Errorf("user_profile_photos.FindListByIdList: %w", err)
	}

	return resp, nil
}

func (m *defaultUserProfilePhotosModel) Update2(ctx context.Context, data *UserProfilePhotos) error {
	query := fmt.Sprintf("update `user_profile_photos` set %s where `id` = ?", userProfilePhotosRowsWithPlaceHolder)

	_, err := m.db.Exec(ctx, query, data.UserId, data.PhotoId, data.Date2, data.Deleted, data.Id)
	if err != nil {
		return fmt.Errorf("user_profile_photos.Update2 exec: %w", err)
	}

	return nil
}

func (m *defaultUserProfilePhotosModel) FindOneByUserIdPhotoId(ctx context.Context, userId int64, photoId int64) (*UserProfilePhotos, error) {
	query := fmt.Sprintf("select %s from user_profile_photos where user_id = ? AND photo_id = ? limit 1", userProfilePhotosRows)
	var resp UserProfilePhotos

	err := m.db.QueryRowPartial(ctx, &resp, query, userId, photoId)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return nil, &NotFoundError{
				Resource: "user_profile_photos",
				Key:      fmt.Sprintf("user_id=%v,photo_id=%v", userId, photoId),
				Cause:    err,
			}
		}
		return nil, fmt.Errorf("user_profile_photos.FindOneByUserIdPhotoId: %w", err)
	}

	return &resp, nil
}
