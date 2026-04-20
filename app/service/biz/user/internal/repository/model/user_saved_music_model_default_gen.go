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
	userSavedMusicFieldNames          = builder.RawFieldNames(&UserSavedMusic{})
	userSavedMusicRows                = strings.Join(userSavedMusicFieldNames, ",")
	userSavedMusicRowsExpectAutoSet   = strings.Join(stringx.Remove(userSavedMusicFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	userSavedMusicRowsWithPlaceHolder = strings.Join(stringx.Remove(userSavedMusicFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"

	cacheTUserSavedMusicIdPrefix = "cache:t:user_saved_music:id:"

	cacheUserSavedMusicIdPrefix = "cache#UserSavedMusic#id"

	cacheUserSavedMusicUserIdSavedMusicIdPrefix = "cache#UserId#SavedMusicId"
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
	query := fmt.Sprintf("insert into `user_saved_music` (%s) values (?, ?, ?, ?)", userSavedMusicRowsExpectAutoSet)
	return m.db.Exec(ctx, query, data.UserId, data.SavedMusicId, data.Order2, data.Deleted)
}

func (m *defaultUserSavedMusicModel) Delete2(ctx context.Context, id int32) error {
	query := "delete from `user_saved_music` where `id` = ?"
	_, err := m.db.Exec(ctx, query, id)
	return err
}

func (m *defaultUserSavedMusicModel) FindOne(ctx context.Context, id int32) (*UserSavedMusic, error) {
	query := fmt.Sprintf("select %s from user_saved_music where id = ? limit 1", userSavedMusicRows)
	var resp UserSavedMusic
	err := m.db.QueryRowPartial(ctx, &resp, query, id)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (m *defaultUserSavedMusicModel) FindListByIdList(ctx context.Context, id ...int32) ([]UserSavedMusic, error) {
	if len(id) == 0 {
		return []UserSavedMusic{}, nil
	}
	query := fmt.Sprintf("select %s from user_saved_music where id in (%s)", userSavedMusicRows, sqlx.InInt32List(id))

	var resp []UserSavedMusic
	err := m.db.QueryRowsPartial(ctx, &resp, query)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (m *defaultUserSavedMusicModel) Update2(ctx context.Context, data *UserSavedMusic) error {
	query := fmt.Sprintf("update `user_saved_music` set %s where `id` = ?", userSavedMusicRowsWithPlaceHolder)
	_, err := m.db.Exec(ctx, query, data.UserId, data.SavedMusicId, data.Order2, data.Deleted, data.Id)
	return err
}

func (m *defaultUserSavedMusicModel) FindOneByUserIdSavedMusicId(ctx context.Context, userId int64, savedMusicId int64) (*UserSavedMusic, error) {
	query := fmt.Sprintf("select %s from user_saved_music where user_id = ? AND saved_music_id = ? limit 1", userSavedMusicRows)
	var resp UserSavedMusic
	err := m.db.QueryRowPartial(ctx, &resp, query, userId, savedMusicId)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (m *defaultUserSavedMusicModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s#%v", cacheUserSavedMusicIdPrefix, primary)
}

func (m *defaultUserSavedMusicModel) queryPrimary(ctx context.Context, v interface{}, primary interface{}) error {
	query := fmt.Sprintf("select %s from user_saved_music where id = ? limit 1", userSavedMusicRows)
	return m.db.QueryRowPartial(ctx, v, query, primary)
}
