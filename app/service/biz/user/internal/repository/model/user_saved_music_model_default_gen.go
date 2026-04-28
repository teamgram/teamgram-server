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
	userSavedMusicFieldNames          = builder.RawFieldNames(&UserSavedMusic{})
	userSavedMusicRows                = strings.Join(userSavedMusicFieldNames, ",")
	userSavedMusicRowsExpectAutoSet   = strings.Join(stringx.Remove(userSavedMusicFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	userSavedMusicRowsWithPlaceHolder = strings.Join(stringx.Remove(userSavedMusicFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"
)

type (
	userSavedMusicModel interface {
		Insert2(ctx context.Context, data *UserSavedMusic) (sql.Result, error)
		FindOne(ctx context.Context, id int32) (*UserSavedMusic, error)
		FindListByIdList(ctx context.Context, id ...int32) ([]UserSavedMusic, error)
		Update2(ctx context.Context, data *UserSavedMusic) error
		Delete2(ctx context.Context, id int32) error

		FindOneByUserIdSavedMusicId(ctx context.Context, userId int64, savedMusicId int64) (*UserSavedMusic, error)
	}

	defaultUserSavedMusicModel struct {
		db *sqlx.DB
	}

	UserSavedMusic struct {
		Id           int32 `db:"id" json:"id"`
		UserId       int64 `db:"user_id" json:"user_id"`
		SavedMusicId int64 `db:"saved_music_id" json:"saved_music_id"`
		Order2       int64 `db:"order2" json:"order2"`
		Deleted      bool  `db:"deleted" json:"deleted"`
	}
)

func newUserSavedMusicModel(db *sqlx.DB) *defaultUserSavedMusicModel {
	return &defaultUserSavedMusicModel{
		db: db,
	}
}

func (m *defaultUserSavedMusicModel) Insert2(ctx context.Context, data *UserSavedMusic) (sql.Result, error) {
	tableName := "user_saved_music"
	query := fmt.Sprintf("insert into `%s` (%s) values (?, ?, ?, ?)", tableName, userSavedMusicRowsExpectAutoSet)

	r, err := m.db.Exec(ctx, query, data.UserId, data.SavedMusicId, data.Order2, data.Deleted)
	if err != nil {
		return nil, fmt.Errorf("user_saved_music.Insert2 exec: %w", err)
	}

	return r, nil
}

func (m *defaultUserSavedMusicModel) Delete2(ctx context.Context, id int32) error {
	tableName := "user_saved_music"
	query := fmt.Sprintf("delete from `%s` where `id` = ?", tableName)

	_, err := m.db.Exec(ctx, query, id)
	if err != nil {
		return fmt.Errorf("user_saved_music.Delete2 exec: %w", err)
	}

	return nil
}

func (m *defaultUserSavedMusicModel) FindOne(ctx context.Context, id int32) (*UserSavedMusic, error) {
	tableName := "user_saved_music"
	query := fmt.Sprintf("select %s from %s where id = ? limit 1", userSavedMusicRows, tableName)
	var resp UserSavedMusic

	err := m.db.QueryRowPartial(ctx, &resp, query, id)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return nil, &NotFoundError{
				Resource: "user_saved_music",
				Key:      fmt.Sprintf("id=%v", id),
				Cause:    err,
			}
		}
		return nil, fmt.Errorf("user_saved_music.FindOne: %w", err)
	}

	return &resp, nil
}

func (m *defaultUserSavedMusicModel) FindListByIdList(ctx context.Context, id ...int32) ([]UserSavedMusic, error) {
	if len(id) == 0 {
		return []UserSavedMusic{}, nil
	}
	tableName := "user_saved_music"
	query := fmt.Sprintf("select %s from %s where id in (%s)", userSavedMusicRows, tableName, sqlx.InInt32List(id))

	var resp []UserSavedMusic
	err := m.db.QueryRowsPartial(ctx, &resp, query)
	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return []UserSavedMusic{}, nil
		}
		return nil, fmt.Errorf("user_saved_music.FindListByIdList: %w", err)
	}

	return resp, nil
}

func (m *defaultUserSavedMusicModel) Update2(ctx context.Context, data *UserSavedMusic) error {
	tableName := "user_saved_music"
	query := fmt.Sprintf("update `%s` set %s where `id` = ?", tableName, userSavedMusicRowsWithPlaceHolder)

	_, err := m.db.Exec(ctx, query, data.UserId, data.SavedMusicId, data.Order2, data.Deleted, data.Id)
	if err != nil {
		return fmt.Errorf("user_saved_music.Update2 exec: %w", err)
	}

	return nil
}

func (m *defaultUserSavedMusicModel) FindOneByUserIdSavedMusicId(ctx context.Context, userId int64, savedMusicId int64) (*UserSavedMusic, error) {
	tableName := "user_saved_music"
	query := fmt.Sprintf("select %s from %s where user_id = ? AND saved_music_id = ? limit 1", userSavedMusicRows, tableName)
	var resp UserSavedMusic

	err := m.db.QueryRowPartial(ctx, &resp, query, userId, savedMusicId)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return nil, &NotFoundError{
				Resource: "user_saved_music",
				Key:      fmt.Sprintf("user_id=%v,saved_music_id=%v", userId, savedMusicId),
				Cause:    err,
			}
		}
		return nil, fmt.Errorf("user_saved_music.FindOneByUserIdSavedMusicId: %w", err)
	}

	return &resp, nil
}
