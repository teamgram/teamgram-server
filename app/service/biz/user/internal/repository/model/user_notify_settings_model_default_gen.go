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
	user_notify_settingsFieldNames          = builder.RawFieldNames(&UserNotifySettings{})
	user_notify_settingsRows                = strings.Join(user_notify_settingsFieldNames, ",")
	user_notify_settingsRowsExpectAutoSet   = strings.Join(stringx.Remove(user_notify_settingsFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	user_notify_settingsRowsWithPlaceHolder = strings.Join(stringx.Remove(user_notify_settingsFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"

	cacheTUserNotifySettingsIdPrefix = "cache:t:user_notify_settings:id:"

	cacheUserNotifySettingsIdPrefix = "cache#UserNotifySettings#id"

	cacheUserNotifySettingsUserIdPeerTypePeerIdPrefix = "cache#UserId#PeerType#PeerId"
)

type (
	user_notify_settingsModel interface {
		Insert2(ctx context.Context, data *UserNotifySettings) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*UserNotifySettings, error)
		FindListByIdList(ctx context.Context, id ...int64) ([]UserNotifySettings, error)
		Update2(ctx context.Context, data *UserNotifySettings) error
		Delete2(ctx context.Context, id int64) error

		FindOneByUserIdPeerTypePeerId(ctx context.Context, userId int64, peerType int32, peerId int64) (*UserNotifySettings, error)
	}

	defaultUserNotifySettingsModel struct {
		db *sqlx.DB
	}

	UserNotifySettings struct {
		Id           int64  `db:"id" json:"id"`
		UserId       int64  `db:"user_id" json:"user_id"`
		PeerType     int32  `db:"peer_type" json:"peer_type"`
		PeerId       int64  `db:"peer_id" json:"peer_id"`
		ShowPreviews int32  `db:"show_previews" json:"show_previews"`
		Silent       int32  `db:"silent" json:"silent"`
		MuteUntil    int32  `db:"mute_until" json:"mute_until"`
		Sound        string `db:"sound" json:"sound"`
		Deleted      bool   `db:"deleted" json:"deleted"`
	}
)

func newUserNotifySettingsModel(db *sqlx.DB) *defaultUserNotifySettingsModel {
	return &defaultUserNotifySettingsModel{
		db: db,
	}
}

func (m *defaultUserNotifySettingsModel) Insert2(ctx context.Context, data *UserNotifySettings) (sql.Result, error) {
	query := fmt.Sprintf("insert into `user_notify_settings` (%s) values (?, ?, ?, ?, ?, ?, ?, ?)", user_notify_settingsRowsExpectAutoSet)
	return m.db.Exec(ctx, query, data.UserId, data.PeerType, data.PeerId, data.ShowPreviews, data.Silent, data.MuteUntil, data.Sound, data.Deleted)
}

func (m *defaultUserNotifySettingsModel) Delete2(ctx context.Context, id int64) error {
	query := "delete from `user_notify_settings` where `id` = ?"
	_, err := m.db.Exec(ctx, query, id)
	return err
}

func (m *defaultUserNotifySettingsModel) FindOne(ctx context.Context, id int64) (*UserNotifySettings, error) {
	query := fmt.Sprintf("select %s from user_notify_settings where id = ? limit 1", user_notify_settingsRows)
	var resp UserNotifySettings
	err := m.db.QueryRowPartial(ctx, &resp, query, id)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (m *defaultUserNotifySettingsModel) FindListByIdList(ctx context.Context, id ...int64) ([]UserNotifySettings, error) {
	if len(id) == 0 {
		return []UserNotifySettings{}, nil
	}

	query := fmt.Sprintf("select %s from user_notify_settings where id in (%s)", user_notify_settingsRows, sqlx.InInt64List(id))

	var resp []UserNotifySettings
	err := m.db.QueryRowsPartial(ctx, &resp, query)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (m *defaultUserNotifySettingsModel) Update2(ctx context.Context, data *UserNotifySettings) error {
	query := fmt.Sprintf("update `user_notify_settings` set %s where `id` = ?", user_notify_settingsRowsWithPlaceHolder)
	_, err := m.db.Exec(ctx, query, data.UserId, data.PeerType, data.PeerId, data.ShowPreviews, data.Silent, data.MuteUntil, data.Sound, data.Deleted, data.Id)
	return err
}

func (m *defaultUserNotifySettingsModel) FindOneByUserIdPeerTypePeerId(ctx context.Context, userId int64, peerType int32, peerId int64) (*UserNotifySettings, error) {
	query := fmt.Sprintf("select %s from user_notify_settings where user_id = ? AND peer_type = ? AND peer_id = ? limit 1", user_notify_settingsRows)
	var resp UserNotifySettings
	err := m.db.QueryRowPartial(ctx, &resp, query, userId, peerType, peerId)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (m *defaultUserNotifySettingsModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s#%v", cacheUserNotifySettingsIdPrefix, primary)
}

func (m *defaultUserNotifySettingsModel) queryPrimary(ctx context.Context, v interface{}, primary interface{}) error {
	query := fmt.Sprintf("select %s from user_notify_settings where id = ? limit 1", user_notify_settingsRows)
	return m.db.QueryRowPartial(ctx, v, query, primary)
}
