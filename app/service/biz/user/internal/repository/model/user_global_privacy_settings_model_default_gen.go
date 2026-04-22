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
	userGlobalPrivacySettingsFieldNames          = builder.RawFieldNames(&UserGlobalPrivacySettings{})
	userGlobalPrivacySettingsRows                = strings.Join(userGlobalPrivacySettingsFieldNames, ",")
	userGlobalPrivacySettingsRowsExpectAutoSet   = strings.Join(stringx.Remove(userGlobalPrivacySettingsFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	userGlobalPrivacySettingsRowsWithPlaceHolder = strings.Join(stringx.Remove(userGlobalPrivacySettingsFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"

	cacheTUserGlobalPrivacySettingsIdPrefix = "cache:t:user_global_privacy_settings:id:"

	cacheUserGlobalPrivacySettingsIdPrefix = "cache#UserGlobalPrivacySettings#id"

	cacheUserGlobalPrivacySettingsUserIdPrefix = "cache#UserId"
)

type (
	userGlobalPrivacySettingsModel interface {
		Insert2(ctx context.Context, data *UserGlobalPrivacySettings) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*UserGlobalPrivacySettings, error)
		FindListByIdList(ctx context.Context, id ...int64) ([]UserGlobalPrivacySettings, error)
		Update2(ctx context.Context, data *UserGlobalPrivacySettings) error
		Delete2(ctx context.Context, id int64) error

		FindOneByUserId(ctx context.Context, userId int64) (*UserGlobalPrivacySettings, error)
		FindListByUserIdList(ctx context.Context, userId ...int64) ([]UserGlobalPrivacySettings, error)
	}

	defaultUserGlobalPrivacySettingsModel struct {
		db *sqlx.DB
	}

	UserGlobalPrivacySettings struct {
		Id                               int64 `db:"id" json:"id"`
		UserId                           int64 `db:"user_id" json:"user_id"`
		ArchiveAndMuteNewNoncontactPeers bool  `db:"archive_and_mute_new_noncontact_peers" json:"archive_and_mute_new_noncontact_peers"`
		KeepArchivedUnmuted              bool  `db:"keep_archived_unmuted" json:"keep_archived_unmuted"`
		KeepArchivedFolders              bool  `db:"keep_archived_folders" json:"keep_archived_folders"`
		HideReadMarks                    bool  `db:"hide_read_marks" json:"hide_read_marks"`
		NewNoncontactPeersRequirePremium bool  `db:"new_noncontact_peers_require_premium" json:"new_noncontact_peers_require_premium"`
	}
)

func newUserGlobalPrivacySettingsModel(db *sqlx.DB) *defaultUserGlobalPrivacySettingsModel {
	return &defaultUserGlobalPrivacySettingsModel{
		db: db,
	}
}

func (m *defaultUserGlobalPrivacySettingsModel) Insert2(ctx context.Context, data *UserGlobalPrivacySettings) (sql.Result, error) {
	query := fmt.Sprintf("insert into `user_global_privacy_settings` (%s) values (?, ?, ?, ?, ?, ?)", userGlobalPrivacySettingsRowsExpectAutoSet)
	return m.db.Exec(ctx, query, data.UserId, data.ArchiveAndMuteNewNoncontactPeers, data.KeepArchivedUnmuted, data.KeepArchivedFolders, data.HideReadMarks, data.NewNoncontactPeersRequirePremium)
}

func (m *defaultUserGlobalPrivacySettingsModel) Delete2(ctx context.Context, id int64) error {
	query := "delete from `user_global_privacy_settings` where `id` = ?"
	_, err := m.db.Exec(ctx, query, id)
	return err
}

func (m *defaultUserGlobalPrivacySettingsModel) FindOne(ctx context.Context, id int64) (*UserGlobalPrivacySettings, error) {
	query := fmt.Sprintf("select %s from user_global_privacy_settings where id = ? limit 1", userGlobalPrivacySettingsRows)
	var resp UserGlobalPrivacySettings
	err := m.db.QueryRowPartial(ctx, &resp, query, id)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (m *defaultUserGlobalPrivacySettingsModel) FindListByIdList(ctx context.Context, id ...int64) ([]UserGlobalPrivacySettings, error) {
	if len(id) == 0 {
		return []UserGlobalPrivacySettings{}, nil
	}

	query := fmt.Sprintf("select %s from user_global_privacy_settings where id in (%s)", userGlobalPrivacySettingsRows, sqlx.InInt64List(id))

	var resp []UserGlobalPrivacySettings
	err := m.db.QueryRowsPartial(ctx, &resp, query)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (m *defaultUserGlobalPrivacySettingsModel) Update2(ctx context.Context, data *UserGlobalPrivacySettings) error {
	query := fmt.Sprintf("update `user_global_privacy_settings` set %s where `id` = ?", userGlobalPrivacySettingsRowsWithPlaceHolder)
	_, err := m.db.Exec(ctx, query, data.UserId, data.ArchiveAndMuteNewNoncontactPeers, data.KeepArchivedUnmuted, data.KeepArchivedFolders, data.HideReadMarks, data.NewNoncontactPeersRequirePremium, data.Id)
	return err
}

func (m *defaultUserGlobalPrivacySettingsModel) FindOneByUserId(ctx context.Context, userId int64) (*UserGlobalPrivacySettings, error) {
	query := fmt.Sprintf("select %s from user_global_privacy_settings where user_id = ? limit 1", userGlobalPrivacySettingsRows)
	var resp UserGlobalPrivacySettings
	err := m.db.QueryRowPartial(ctx, &resp, query, userId)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (m *defaultUserGlobalPrivacySettingsModel) FindListByUserIdList(ctx context.Context, userId ...int64) ([]UserGlobalPrivacySettings, error) {
	if len(userId) == 0 {
		return []UserGlobalPrivacySettings{}, nil
	}

	query := fmt.Sprintf("select %s from user_global_privacy_settings where user_id in (%s)", userGlobalPrivacySettingsRows, sqlx.InInt64List(userId))

	var resp []UserGlobalPrivacySettings
	err := m.db.QueryRowsPartial(ctx, &resp, query)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (m *defaultUserGlobalPrivacySettingsModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s#%v", cacheUserGlobalPrivacySettingsIdPrefix, primary)
}

func (m *defaultUserGlobalPrivacySettingsModel) queryPrimary(ctx context.Context, v interface{}, primary interface{}) error {
	query := fmt.Sprintf("select %s from user_global_privacy_settings where id = ? limit 1", userGlobalPrivacySettingsRows)
	return m.db.QueryRowPartial(ctx, v, query, primary)
}
