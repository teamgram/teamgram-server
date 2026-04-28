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
	userNotifySettingsFieldNames          = builder.RawFieldNames(&UserNotifySettings{})
	userNotifySettingsRows                = strings.Join(userNotifySettingsFieldNames, ",")
	userNotifySettingsRowsExpectAutoSet   = strings.Join(stringx.Remove(userNotifySettingsFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	userNotifySettingsRowsWithPlaceHolder = strings.Join(stringx.Remove(userNotifySettingsFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"
)

type (
	userNotifySettingsModel interface {
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
	tableName := "user_notify_settings"
	query := fmt.Sprintf("insert into `%s` (%s) values (?, ?, ?, ?, ?, ?, ?, ?)", tableName, userNotifySettingsRowsExpectAutoSet)

	r, err := m.db.Exec(ctx, query, data.UserId, data.PeerType, data.PeerId, data.ShowPreviews, data.Silent, data.MuteUntil, data.Sound, data.Deleted)
	if err != nil {
		return nil, fmt.Errorf("user_notify_settings.Insert2 exec: %w", err)
	}

	return r, nil
}

func (m *defaultUserNotifySettingsModel) Delete2(ctx context.Context, id int64) error {
	tableName := "user_notify_settings"
	query := fmt.Sprintf("delete from `%s` where `id` = ?", tableName)

	_, err := m.db.Exec(ctx, query, id)
	if err != nil {
		return fmt.Errorf("user_notify_settings.Delete2 exec: %w", err)
	}

	return nil
}

func (m *defaultUserNotifySettingsModel) FindOne(ctx context.Context, id int64) (*UserNotifySettings, error) {
	tableName := "user_notify_settings"
	query := fmt.Sprintf("select %s from %s where id = ? limit 1", userNotifySettingsRows, tableName)
	var resp UserNotifySettings

	err := m.db.QueryRowPartial(ctx, &resp, query, id)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return nil, &NotFoundError{
				Resource: "user_notify_settings",
				Key:      fmt.Sprintf("id=%v", id),
				Cause:    err,
			}
		}
		return nil, fmt.Errorf("user_notify_settings.FindOne: %w", err)
	}

	return &resp, nil
}

func (m *defaultUserNotifySettingsModel) FindListByIdList(ctx context.Context, id ...int64) ([]UserNotifySettings, error) {
	if len(id) == 0 {
		return []UserNotifySettings{}, nil
	}
	tableName := "user_notify_settings"

	query := fmt.Sprintf("select %s from %s where id in (%s)", userNotifySettingsRows, tableName, sqlx.InInt64List(id))

	var resp []UserNotifySettings
	err := m.db.QueryRowsPartial(ctx, &resp, query)
	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return []UserNotifySettings{}, nil
		}
		return nil, fmt.Errorf("user_notify_settings.FindListByIdList: %w", err)
	}

	return resp, nil
}

func (m *defaultUserNotifySettingsModel) Update2(ctx context.Context, data *UserNotifySettings) error {
	tableName := "user_notify_settings"
	query := fmt.Sprintf("update `%s` set %s where `id` = ?", tableName, userNotifySettingsRowsWithPlaceHolder)

	_, err := m.db.Exec(ctx, query, data.UserId, data.PeerType, data.PeerId, data.ShowPreviews, data.Silent, data.MuteUntil, data.Sound, data.Deleted, data.Id)
	if err != nil {
		return fmt.Errorf("user_notify_settings.Update2 exec: %w", err)
	}

	return nil
}

func (m *defaultUserNotifySettingsModel) FindOneByUserIdPeerTypePeerId(ctx context.Context, userId int64, peerType int32, peerId int64) (*UserNotifySettings, error) {
	tableName := "user_notify_settings"
	query := fmt.Sprintf("select %s from %s where user_id = ? AND peer_type = ? AND peer_id = ? limit 1", userNotifySettingsRows, tableName)
	var resp UserNotifySettings

	err := m.db.QueryRowPartial(ctx, &resp, query, userId, peerType, peerId)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return nil, &NotFoundError{
				Resource: "user_notify_settings",
				Key:      fmt.Sprintf("user_id=%v,peer_type=%v,peer_id=%v", userId, peerType, peerId),
				Cause:    err,
			}
		}
		return nil, fmt.Errorf("user_notify_settings.FindOneByUserIdPeerTypePeerId: %w", err)
	}

	return &resp, nil
}
