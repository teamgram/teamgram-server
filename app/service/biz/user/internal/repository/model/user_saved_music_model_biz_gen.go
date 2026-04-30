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
)

var _ *sql.Result
var _ = fmt.Sprintf
var _ = strings.Join
var _ = errors.Is
var _ *sqlx.DB
var _ *sqlx.Tx

type bizUserSavedMusicModel interface {
	InsertOrUpdate(ctx context.Context, data *UserSavedMusic) (lastInsertId, rowsAffected int64, err error)
	SelectList(ctx context.Context, userId int64) ([]UserSavedMusic, error)
	SelectListWithCB(ctx context.Context, userId int64, cb func(sz, i int, v *UserSavedMusic)) ([]UserSavedMusic, error)
	SelectListByIdList(ctx context.Context, userId int64, idList []int64) ([]UserSavedMusic, error)
	SelectListByIdListWithCB(ctx context.Context, userId int64, idList []int64, cb func(sz, i int, v *UserSavedMusic)) ([]UserSavedMusic, error)
	Delete(ctx context.Context, userId int64, savedMusicId int64) (rowsAffected int64, err error)
}

type UserSavedMusicTxModel interface {
	InsertOrUpdate(data *UserSavedMusic) (lastInsertId, rowsAffected int64, err error)
	SelectList(userId int64) ([]UserSavedMusic, error)
	SelectListByIdList(userId int64, idList []int64) ([]UserSavedMusic, error)
	Delete(userId int64, savedMusicId int64) (rowsAffected int64, err error)
}

type defaultUserSavedMusicTxModel struct {
	tx *sqlx.Tx
}

func NewUserSavedMusicTxModel(tx *sqlx.Tx) UserSavedMusicTxModel {
	return &defaultUserSavedMusicTxModel{tx: tx}
}

// InsertOrUpdate
// insert into user_saved_music(user_id, saved_music_id, order2) values (:user_id, :saved_music_id, :order2) on duplicate key update deleted = 0
func (m *defaultUserSavedMusicModel) InsertOrUpdate(ctx context.Context, data *UserSavedMusic) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into user_saved_music(user_id, saved_music_id, order2) values (:user_id, :saved_music_id, :order2) on duplicate key update deleted = 0"
		r     sql.Result
	)

	r, err = m.db.NamedExec(ctx, query, data)
	if err != nil {
		err = fmt.Errorf("user_saved_music.InsertOrUpdate named exec: %w", err)
		return
	}

	lastInsertId, err = r.LastInsertId()
	if err != nil {
		err = fmt.Errorf("user_saved_music.InsertOrUpdate last insert id: %w", err)
		return
	}
	rowsAffected, err = r.RowsAffected()
	if err != nil {
		err = fmt.Errorf("user_saved_music.InsertOrUpdate rows affected: %w", err)
	}

	return

}

// InsertOrUpdate
// insert into user_saved_music(user_id, saved_music_id, order2) values (:user_id, :saved_music_id, :order2) on duplicate key update deleted = 0
func (m *defaultUserSavedMusicTxModel) InsertOrUpdate(data *UserSavedMusic) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into user_saved_music(user_id, saved_music_id, order2) values (:user_id, :saved_music_id, :order2) on duplicate key update deleted = 0"
		r     sql.Result
	)

	r, err = m.tx.NamedExec(query, data)
	if err != nil {
		err = fmt.Errorf("user_saved_music.InsertOrUpdate named exec: %w", err)
		return
	}

	lastInsertId, err = r.LastInsertId()
	if err != nil {
		err = fmt.Errorf("user_saved_music.InsertOrUpdate last insert id: %w", err)
		return
	}
	rowsAffected, err = r.RowsAffected()
	if err != nil {
		err = fmt.Errorf("user_saved_music.InsertOrUpdate rows affected: %w", err)
	}

	return
}

// SelectList
// select id, user_id, saved_music_id from user_saved_music where user_id = :user_id and deleted = 0
func (m *defaultUserSavedMusicModel) SelectList(ctx context.Context, userId int64) (rList []UserSavedMusic, err error) {
	var (
		query  = "select id, user_id, saved_music_id from user_saved_music where user_id = ? and deleted = 0"
		values []UserSavedMusic
	)
	err = m.db.QueryRowsPartial(ctx, &values, query, userId)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			rList = []UserSavedMusic{}
			err = nil
			return
		}
		err = fmt.Errorf("user_saved_music.SelectList: %w", err)
		return
	}

	rList = values

	return
}

// SelectList
// select id, user_id, saved_music_id from user_saved_music where user_id = :user_id and deleted = 0
func (m *defaultUserSavedMusicTxModel) SelectList(userId int64) (rList []UserSavedMusic, err error) {
	var (
		query  = "select id, user_id, saved_music_id from user_saved_music where user_id = ? and deleted = 0"
		values []UserSavedMusic
	)
	err = m.tx.QueryRowsPartial(&values, query, userId)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			rList = []UserSavedMusic{}
			err = nil
			return
		}
		err = fmt.Errorf("user_saved_music.SelectList: %w", err)
		return
	}

	rList = values

	return
}

// SelectListWithCB
// select id, user_id, saved_music_id from user_saved_music where user_id = :user_id and deleted = 0
func (m *defaultUserSavedMusicModel) SelectListWithCB(ctx context.Context, userId int64, cb func(sz, i int, v *UserSavedMusic)) (rList []UserSavedMusic, err error) {
	var (
		query  = "select id, user_id, saved_music_id from user_saved_music where user_id = ? and deleted = 0"
		values []UserSavedMusic
	)
	err = m.db.QueryRowsPartial(ctx, &values, query, userId)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			rList = []UserSavedMusic{}
			err = nil
			return
		}
		err = fmt.Errorf("user_saved_music.SelectListWithCB: %w", err)
		return
	}

	rList = values

	if cb != nil {
		sz := len(rList)
		for i := 0; i < sz; i++ {
			cb(sz, i, &rList[i])
		}
	}

	return
}

// SelectListByIdList
// select id, user_id, saved_music_id from user_saved_music where user_id = :user_id and deleted = 0 and saved_music_id in (:idList)
func (m *defaultUserSavedMusicModel) SelectListByIdList(ctx context.Context, userId int64, idList []int64) (rList []UserSavedMusic, err error) {
	var (
		query  = fmt.Sprintf("select id, user_id, saved_music_id from user_saved_music where user_id = ? and deleted = 0 and saved_music_id in (%s)", sqlx.InInt64List(idList))
		values []UserSavedMusic
	)
	if len(idList) == 0 {
		rList = []UserSavedMusic{}
		return
	}

	err = m.db.QueryRowsPartial(ctx, &values, query, userId)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			rList = []UserSavedMusic{}
			err = nil
			return
		}
		err = fmt.Errorf("user_saved_music.SelectListByIdList: %w", err)
		return
	}

	rList = values

	return
}

// SelectListByIdList
// select id, user_id, saved_music_id from user_saved_music where user_id = :user_id and deleted = 0 and saved_music_id in (:idList)
func (m *defaultUserSavedMusicTxModel) SelectListByIdList(userId int64, idList []int64) (rList []UserSavedMusic, err error) {
	var (
		query  = fmt.Sprintf("select id, user_id, saved_music_id from user_saved_music where user_id = ? and deleted = 0 and saved_music_id in (%s)", sqlx.InInt64List(idList))
		values []UserSavedMusic
	)
	if len(idList) == 0 {
		rList = []UserSavedMusic{}
		return
	}

	err = m.tx.QueryRowsPartial(&values, query, userId)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			rList = []UserSavedMusic{}
			err = nil
			return
		}
		err = fmt.Errorf("user_saved_music.SelectListByIdList: %w", err)
		return
	}

	rList = values

	return
}

// SelectListByIdListWithCB
// select id, user_id, saved_music_id from user_saved_music where user_id = :user_id and deleted = 0 and saved_music_id in (:idList)
func (m *defaultUserSavedMusicModel) SelectListByIdListWithCB(ctx context.Context, userId int64, idList []int64, cb func(sz, i int, v *UserSavedMusic)) (rList []UserSavedMusic, err error) {
	var (
		query  = fmt.Sprintf("select id, user_id, saved_music_id from user_saved_music where user_id = ? and deleted = 0 and saved_music_id in (%s)", sqlx.InInt64List(idList))
		values []UserSavedMusic
	)
	if len(idList) == 0 {
		rList = []UserSavedMusic{}
		return
	}

	err = m.db.QueryRowsPartial(ctx, &values, query, userId)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			rList = []UserSavedMusic{}
			err = nil
			return
		}
		err = fmt.Errorf("user_saved_music.SelectListByIdListWithCB: %w", err)
		return
	}

	rList = values

	if cb != nil {
		sz := len(rList)
		for i := 0; i < sz; i++ {
			cb(sz, i, &rList[i])
		}
	}

	return
}

// Delete
// update user_saved_music set deleted = 1, order2 = 0 where user_id = :user_id and saved_music_id = :saved_music_id
func (m *defaultUserSavedMusicModel) Delete(ctx context.Context, userId int64, savedMusicId int64) (rowsAffected int64, err error) {

	var (
		query   = "update user_saved_music set deleted = 1, order2 = 0 where user_id = ? and saved_music_id = ?"
		rResult sql.Result
	)

	rResult, err = m.db.Exec(ctx, query, userId, savedMusicId)

	if err != nil {
		err = fmt.Errorf("user_saved_music.Delete exec: %w", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		err = fmt.Errorf("user_saved_music.Delete rows affected: %w", err)
		return
	}

	return
}

// Delete
// update user_saved_music set deleted = 1, order2 = 0 where user_id = :user_id and saved_music_id = :saved_music_id
func (m *defaultUserSavedMusicTxModel) Delete(userId int64, savedMusicId int64) (rowsAffected int64, err error) {
	var (
		query   = "update user_saved_music set deleted = 1, order2 = 0 where user_id = ? and saved_music_id = ?"
		rResult sql.Result
	)
	rResult, err = m.tx.Exec(query, userId, savedMusicId)

	if err != nil {
		err = fmt.Errorf("user_saved_music.Delete exec: %w", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		err = fmt.Errorf("user_saved_music.Delete rows affected: %w", err)
		return
	}

	return
}
