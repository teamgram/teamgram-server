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

type bizDialogVisualSettingsModel interface {
	Upsert(ctx context.Context, data *DialogVisualSettings) (lastInsertId, rowsAffected int64, err error)
	SelectByUserPeer(ctx context.Context, userId int64, peerType int32, peerId int64) (*DialogVisualSettings, error)
}

type DialogVisualSettingsTxModel interface {
	Upsert(data *DialogVisualSettings) (lastInsertId, rowsAffected int64, err error)
	SelectByUserPeer(userId int64, peerType int32, peerId int64) (*DialogVisualSettings, error)
}

type defaultDialogVisualSettingsTxModel struct {
	tx *sqlx.Tx
}

func NewDialogVisualSettingsTxModel(tx *sqlx.Tx) DialogVisualSettingsTxModel {
	return &defaultDialogVisualSettingsTxModel{tx: tx}
}

// Upsert
// insert into dialog_visual_settings(user_id, peer_type, peer_id, wallpaper_id, wallpaper_overridden, visual_version) values (:user_id, :peer_type, :peer_id, :wallpaper_id, :wallpaper_overridden, :visual_version) on duplicate key update wallpaper_id = values(wallpaper_id), wallpaper_overridden = values(wallpaper_overridden), visual_version = visual_version + 1
func (m *defaultDialogVisualSettingsModel) Upsert(ctx context.Context, data *DialogVisualSettings) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into dialog_visual_settings(user_id, peer_type, peer_id, wallpaper_id, wallpaper_overridden, visual_version) values (:user_id, :peer_type, :peer_id, :wallpaper_id, :wallpaper_overridden, :visual_version) on duplicate key update wallpaper_id = values(wallpaper_id), wallpaper_overridden = values(wallpaper_overridden), visual_version = visual_version + 1"
		r     sql.Result
	)

	r, err = m.db.NamedExec(ctx, query, data)
	if err != nil {
		err = fmt.Errorf("dialog_visual_settings.Upsert named exec: %w", err)
		return
	}

	lastInsertId, err = r.LastInsertId()
	if err != nil {
		err = fmt.Errorf("dialog_visual_settings.Upsert last insert id: %w", err)
		return
	}
	rowsAffected, err = r.RowsAffected()
	if err != nil {
		err = fmt.Errorf("dialog_visual_settings.Upsert rows affected: %w", err)
	}

	return

}

// Upsert
// insert into dialog_visual_settings(user_id, peer_type, peer_id, wallpaper_id, wallpaper_overridden, visual_version) values (:user_id, :peer_type, :peer_id, :wallpaper_id, :wallpaper_overridden, :visual_version) on duplicate key update wallpaper_id = values(wallpaper_id), wallpaper_overridden = values(wallpaper_overridden), visual_version = visual_version + 1
func (m *defaultDialogVisualSettingsTxModel) Upsert(data *DialogVisualSettings) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into dialog_visual_settings(user_id, peer_type, peer_id, wallpaper_id, wallpaper_overridden, visual_version) values (:user_id, :peer_type, :peer_id, :wallpaper_id, :wallpaper_overridden, :visual_version) on duplicate key update wallpaper_id = values(wallpaper_id), wallpaper_overridden = values(wallpaper_overridden), visual_version = visual_version + 1"
		r     sql.Result
	)

	r, err = m.tx.NamedExec(query, data)
	if err != nil {
		err = fmt.Errorf("dialog_visual_settings.Upsert named exec: %w", err)
		return
	}

	lastInsertId, err = r.LastInsertId()
	if err != nil {
		err = fmt.Errorf("dialog_visual_settings.Upsert last insert id: %w", err)
		return
	}
	rowsAffected, err = r.RowsAffected()
	if err != nil {
		err = fmt.Errorf("dialog_visual_settings.Upsert rows affected: %w", err)
	}

	return
}

// SelectByUserPeer
// select user_id, peer_type, peer_id, wallpaper_id, wallpaper_overridden, visual_version from dialog_visual_settings where user_id = :user_id and peer_type = :peer_type and peer_id = :peer_id limit 1
func (m *defaultDialogVisualSettingsModel) SelectByUserPeer(ctx context.Context, userId int64, peerType int32, peerId int64) (rValue *DialogVisualSettings, err error) {

	var (
		query = "select user_id, peer_type, peer_id, wallpaper_id, wallpaper_overridden, visual_version from dialog_visual_settings where user_id = ? and peer_type = ? and peer_id = ? limit 1"
		do    = &DialogVisualSettings{}
	)
	err = m.db.QueryRowPartial(ctx, do, query, userId, peerType, peerId)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return nil, &NotFoundError{
				Resource: "dialog_visual_settings",
				Key:      fmt.Sprintf("user_id=%v,peer_type=%v,peer_id=%v", userId, peerType, peerId),
				Cause:    err,
			}
		}
		err = fmt.Errorf("dialog_visual_settings.SelectByUserPeer: %w", err)
		return
	} else {
		rValue = do
	}

	return
}

// SelectByUserPeer
// select user_id, peer_type, peer_id, wallpaper_id, wallpaper_overridden, visual_version from dialog_visual_settings where user_id = :user_id and peer_type = :peer_type and peer_id = :peer_id limit 1
func (m *defaultDialogVisualSettingsTxModel) SelectByUserPeer(userId int64, peerType int32, peerId int64) (rValue *DialogVisualSettings, err error) {
	var (
		query = "select user_id, peer_type, peer_id, wallpaper_id, wallpaper_overridden, visual_version from dialog_visual_settings where user_id = ? and peer_type = ? and peer_id = ? limit 1"
		do    = &DialogVisualSettings{}
	)
	err = m.tx.QueryRowPartial(do, query, userId, peerType, peerId)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return nil, &NotFoundError{
				Resource: "dialog_visual_settings",
				Key:      fmt.Sprintf("user_id=%v,peer_type=%v,peer_id=%v", userId, peerType, peerId),
				Cause:    err,
			}
		}
		err = fmt.Errorf("dialog_visual_settings.SelectByUserPeer: %w", err)
		return
	}
	rValue = do

	return
}
