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
	bizUserNotifySettingsModel interface {
		InsertOrUpdate(ctx context.Context, data *UserNotifySettings) (lastInsertId, rowsAffected int64, err error)
		InsertOrUpdateTx(tx *sqlx.Tx, data *UserNotifySettings) (lastInsertId, rowsAffected int64, err error)

		SelectAll(ctx context.Context, userId int64) ([]UserNotifySettings, error)
		SelectAllWithCB(ctx context.Context, userId int64, cb func(sz, i int, v *UserNotifySettings)) ([]UserNotifySettings, error)

		Select(ctx context.Context, userId int64, peerType int32, peerId int64) (*UserNotifySettings, error)

		DeleteAll(ctx context.Context, userId int64) (rowsAffected int64, err error)
		DeleteAllTx(tx *sqlx.Tx, userId int64) (rowsAffected int64, err error)
	}
)

// InsertOrUpdate
// insert into user_notify_settings(user_id, peer_type, peer_id, show_previews, silent, mute_until, sound) values (:user_id, :peer_type, :peer_id, :show_previews, :silent, :mute_until, :sound) on duplicate key update show_previews = values(show_previews), silent = values(silent), mute_until = values(mute_until), sound = values(sound), deleted = 0
func (m *defaultUserNotifySettingsModel) InsertOrUpdate(ctx context.Context, data *UserNotifySettings) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into user_notify_settings(user_id, peer_type, peer_id, show_previews, silent, mute_until, sound) values (:user_id, :peer_type, :peer_id, :show_previews, :silent, :mute_until, :sound) on duplicate key update show_previews = values(show_previews), silent = values(silent), mute_until = values(mute_until), sound = values(sound), deleted = 0"
		r     sql.Result
	)

	r, err = m.db.NamedExec(ctx, query, data)
	if err != nil {
		logx.WithContext(ctx).Errorf("namedExec in InsertOrUpdate(%v), error: %v", data, err)
		return
	}

	lastInsertId, err = r.LastInsertId()
	if err != nil {
		logx.WithContext(ctx).Errorf("lastInsertId in InsertOrUpdate(%v)_error: %v", data, err)
		return
	}
	rowsAffected, err = r.RowsAffected()
	if err != nil {
		logx.WithContext(ctx).Errorf("rowsAffected in InsertOrUpdate(%v)_error: %v", data, err)
	}

	return

}

// InsertOrUpdateTx
// insert into user_notify_settings(user_id, peer_type, peer_id, show_previews, silent, mute_until, sound) values (:user_id, :peer_type, :peer_id, :show_previews, :silent, :mute_until, :sound) on duplicate key update show_previews = values(show_previews), silent = values(silent), mute_until = values(mute_until), sound = values(sound), deleted = 0
func (m *defaultUserNotifySettingsModel) InsertOrUpdateTx(tx *sqlx.Tx, data *UserNotifySettings) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into user_notify_settings(user_id, peer_type, peer_id, show_previews, silent, mute_until, sound) values (:user_id, :peer_type, :peer_id, :show_previews, :silent, :mute_until, :sound) on duplicate key update show_previews = values(show_previews), silent = values(silent), mute_until = values(mute_until), sound = values(sound), deleted = 0"
		r     sql.Result
	)

	r, err = tx.NamedExec(query, data)
	if err != nil {
		logx.WithContext(tx.Context()).Errorf("namedExec in InsertOrUpdate(%v), error: %v", data, err)
		return
	}

	lastInsertId, err = r.LastInsertId()
	if err != nil {
		logx.WithContext(tx.Context()).Errorf("lastInsertId in InsertOrUpdate(%v)_error: %v", data, err)
		return
	}
	rowsAffected, err = r.RowsAffected()
	if err != nil {
		logx.WithContext(tx.Context()).Errorf("rowsAffected in InsertOrUpdate(%v)_error: %v", data, err)
	}

	return
}

// SelectAll
// select id, user_id, peer_type, peer_id, show_previews, silent, mute_until, sound from user_notify_settings where user_id = :user_id and deleted = 0
func (m *defaultUserNotifySettingsModel) SelectAll(ctx context.Context, userId int64) (rList []UserNotifySettings, err error) {
	var (
		query  = "select id, user_id, peer_type, peer_id, show_previews, silent, mute_until, sound from user_notify_settings where user_id = ? and deleted = 0"
		values []UserNotifySettings
	)
	err = m.db.QueryRowsPartial(ctx, &values, query, userId)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectAll(_), error: %v", err)
		return
	}

	rList = values

	return
}

// SelectAllWithCB
// select id, user_id, peer_type, peer_id, show_previews, silent, mute_until, sound from user_notify_settings where user_id = :user_id and deleted = 0
func (m *defaultUserNotifySettingsModel) SelectAllWithCB(ctx context.Context, userId int64, cb func(sz, i int, v *UserNotifySettings)) (rList []UserNotifySettings, err error) {
	var (
		query  = "select id, user_id, peer_type, peer_id, show_previews, silent, mute_until, sound from user_notify_settings where user_id = ? and deleted = 0"
		values []UserNotifySettings
	)
	err = m.db.QueryRowsPartial(ctx, &values, query, userId)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectAll(_), error: %v", err)
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

// Select
// select id, user_id, peer_type, peer_id, show_previews, silent, mute_until, sound from user_notify_settings where user_id = :user_id and peer_type = :peer_type and peer_id = :peer_id and deleted = 0
func (m *defaultUserNotifySettingsModel) Select(ctx context.Context, userId int64, peerType int32, peerId int64) (rValue *UserNotifySettings, err error) {

	var (
		query = "select id, user_id, peer_type, peer_id, show_previews, silent, mute_until, sound from user_notify_settings where user_id = ? and peer_type = ? and peer_id = ? and deleted = 0"
		do    = &UserNotifySettings{}
	)
	err = m.db.QueryRowPartial(ctx, do, query, userId, peerType, peerId)

	if err != nil {
		if !errors.Is(err, sqlx.ErrNotFound) {
			logx.WithContext(ctx).Errorf("queryx in Select(_), error: %v", err)
			return
		} else {
			err = nil
		}
	} else {
		rValue = do
	}

	return
}

// DeleteAll
// update user_notify_settings set deleted = 1 where user_id = :user_id
func (m *defaultUserNotifySettingsModel) DeleteAll(ctx context.Context, userId int64) (rowsAffected int64, err error) {

	var (
		query   = "update user_notify_settings set deleted = 1 where user_id = ?"
		rResult sql.Result
	)

	rResult, err = m.db.Exec(ctx, query, userId)

	if err != nil {
		logx.WithContext(ctx).Errorf("exec in DeleteAll(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		logx.WithContext(ctx).Errorf("rowsAffected in DeleteAll(_), error: %v", err)
	}

	return
}

// DeleteAllTx
// update user_notify_settings set deleted = 1 where user_id = :user_id
func (m *defaultUserNotifySettingsModel) DeleteAllTx(tx *sqlx.Tx, userId int64) (rowsAffected int64, err error) {
	var (
		query   = "update user_notify_settings set deleted = 1 where user_id = ?"
		rResult sql.Result
	)
	rResult, err = tx.Exec(query, userId)

	if err != nil {
		logx.WithContext(tx.Context()).Errorf("exec in DeleteAll(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		logx.WithContext(tx.Context()).Errorf("rowsAffected in DeleteAll(_), error: %v", err)
	}

	return
}
