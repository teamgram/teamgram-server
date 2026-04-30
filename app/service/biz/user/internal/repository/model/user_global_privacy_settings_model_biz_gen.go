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

type bizUserGlobalPrivacySettingsModel interface {
	InsertOrUpdate(ctx context.Context, data *UserGlobalPrivacySettings) (lastInsertId, rowsAffected int64, err error)
	Select(ctx context.Context, userId int64) (*UserGlobalPrivacySettings, error)
}

type UserGlobalPrivacySettingsTxModel interface {
	InsertOrUpdate(data *UserGlobalPrivacySettings) (lastInsertId, rowsAffected int64, err error)
	Select(userId int64) (*UserGlobalPrivacySettings, error)
}

type defaultUserGlobalPrivacySettingsTxModel struct {
	tx *sqlx.Tx
}

func NewUserGlobalPrivacySettingsTxModel(tx *sqlx.Tx) UserGlobalPrivacySettingsTxModel {
	return &defaultUserGlobalPrivacySettingsTxModel{tx: tx}
}

// InsertOrUpdate
// insert into user_global_privacy_settings(user_id, archive_and_mute_new_noncontact_peers, keep_archived_unmuted, keep_archived_folders, hide_read_marks, new_noncontact_peers_require_premium) values (:user_id, :archive_and_mute_new_noncontact_peers, :keep_archived_unmuted, :keep_archived_folders, :hide_read_marks, :new_noncontact_peers_require_premium) on duplicate key update archive_and_mute_new_noncontact_peers = values(archive_and_mute_new_noncontact_peers), keep_archived_unmuted = values(keep_archived_unmuted), keep_archived_folders = values(keep_archived_folders), hide_read_marks = values(hide_read_marks), new_noncontact_peers_require_premium = values(new_noncontact_peers_require_premium)
func (m *defaultUserGlobalPrivacySettingsModel) InsertOrUpdate(ctx context.Context, data *UserGlobalPrivacySettings) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into user_global_privacy_settings(user_id, archive_and_mute_new_noncontact_peers, keep_archived_unmuted, keep_archived_folders, hide_read_marks, new_noncontact_peers_require_premium) values (:user_id, :archive_and_mute_new_noncontact_peers, :keep_archived_unmuted, :keep_archived_folders, :hide_read_marks, :new_noncontact_peers_require_premium) on duplicate key update archive_and_mute_new_noncontact_peers = values(archive_and_mute_new_noncontact_peers), keep_archived_unmuted = values(keep_archived_unmuted), keep_archived_folders = values(keep_archived_folders), hide_read_marks = values(hide_read_marks), new_noncontact_peers_require_premium = values(new_noncontact_peers_require_premium)"
		r     sql.Result
	)

	r, err = m.db.NamedExec(ctx, query, data)
	if err != nil {
		err = fmt.Errorf("user_global_privacy_settings.InsertOrUpdate named exec: %w", err)
		return
	}

	lastInsertId, err = r.LastInsertId()
	if err != nil {
		err = fmt.Errorf("user_global_privacy_settings.InsertOrUpdate last insert id: %w", err)
		return
	}
	rowsAffected, err = r.RowsAffected()
	if err != nil {
		err = fmt.Errorf("user_global_privacy_settings.InsertOrUpdate rows affected: %w", err)
	}

	return

}

// InsertOrUpdate
// insert into user_global_privacy_settings(user_id, archive_and_mute_new_noncontact_peers, keep_archived_unmuted, keep_archived_folders, hide_read_marks, new_noncontact_peers_require_premium) values (:user_id, :archive_and_mute_new_noncontact_peers, :keep_archived_unmuted, :keep_archived_folders, :hide_read_marks, :new_noncontact_peers_require_premium) on duplicate key update archive_and_mute_new_noncontact_peers = values(archive_and_mute_new_noncontact_peers), keep_archived_unmuted = values(keep_archived_unmuted), keep_archived_folders = values(keep_archived_folders), hide_read_marks = values(hide_read_marks), new_noncontact_peers_require_premium = values(new_noncontact_peers_require_premium)
func (m *defaultUserGlobalPrivacySettingsTxModel) InsertOrUpdate(data *UserGlobalPrivacySettings) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into user_global_privacy_settings(user_id, archive_and_mute_new_noncontact_peers, keep_archived_unmuted, keep_archived_folders, hide_read_marks, new_noncontact_peers_require_premium) values (:user_id, :archive_and_mute_new_noncontact_peers, :keep_archived_unmuted, :keep_archived_folders, :hide_read_marks, :new_noncontact_peers_require_premium) on duplicate key update archive_and_mute_new_noncontact_peers = values(archive_and_mute_new_noncontact_peers), keep_archived_unmuted = values(keep_archived_unmuted), keep_archived_folders = values(keep_archived_folders), hide_read_marks = values(hide_read_marks), new_noncontact_peers_require_premium = values(new_noncontact_peers_require_premium)"
		r     sql.Result
	)

	r, err = m.tx.NamedExec(query, data)
	if err != nil {
		err = fmt.Errorf("user_global_privacy_settings.InsertOrUpdate named exec: %w", err)
		return
	}

	lastInsertId, err = r.LastInsertId()
	if err != nil {
		err = fmt.Errorf("user_global_privacy_settings.InsertOrUpdate last insert id: %w", err)
		return
	}
	rowsAffected, err = r.RowsAffected()
	if err != nil {
		err = fmt.Errorf("user_global_privacy_settings.InsertOrUpdate rows affected: %w", err)
	}

	return
}

// Select
// select id, user_id, archive_and_mute_new_noncontact_peers, keep_archived_unmuted, keep_archived_folders, hide_read_marks, new_noncontact_peers_require_premium from user_global_privacy_settings where user_id = :user_id
func (m *defaultUserGlobalPrivacySettingsModel) Select(ctx context.Context, userId int64) (rValue *UserGlobalPrivacySettings, err error) {

	var (
		query = "select id, user_id, archive_and_mute_new_noncontact_peers, keep_archived_unmuted, keep_archived_folders, hide_read_marks, new_noncontact_peers_require_premium from user_global_privacy_settings where user_id = ?"
		do    = &UserGlobalPrivacySettings{}
	)
	err = m.db.QueryRowPartial(ctx, do, query, userId)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return nil, &NotFoundError{
				Resource: "user_global_privacy_settings",
				Key:      fmt.Sprintf("user_id=%v", userId),
				Cause:    err,
			}
		}
		err = fmt.Errorf("user_global_privacy_settings.Select: %w", err)
		return
	} else {
		rValue = do
	}

	return
}

// Select
// select id, user_id, archive_and_mute_new_noncontact_peers, keep_archived_unmuted, keep_archived_folders, hide_read_marks, new_noncontact_peers_require_premium from user_global_privacy_settings where user_id = :user_id
func (m *defaultUserGlobalPrivacySettingsTxModel) Select(userId int64) (rValue *UserGlobalPrivacySettings, err error) {
	var (
		query = "select id, user_id, archive_and_mute_new_noncontact_peers, keep_archived_unmuted, keep_archived_folders, hide_read_marks, new_noncontact_peers_require_premium from user_global_privacy_settings where user_id = ?"
		do    = &UserGlobalPrivacySettings{}
	)
	err = m.tx.QueryRowPartial(do, query, userId)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return nil, &NotFoundError{
				Resource: "user_global_privacy_settings",
				Key:      fmt.Sprintf("user_id=%v", userId),
				Cause:    err,
			}
		}
		err = fmt.Errorf("user_global_privacy_settings.Select: %w", err)
		return
	}
	rValue = do

	return
}
