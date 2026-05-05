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
	dialogVisualSettingsFieldNames          = builder.RawFieldNames(&DialogVisualSettings{})
	dialogVisualSettingsRows                = strings.Join(dialogVisualSettingsFieldNames, ",")
	dialogVisualSettingsRowsExpectAutoSet   = strings.Join(stringx.Remove(dialogVisualSettingsFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	dialogVisualSettingsRowsWithPlaceHolder = strings.Join(stringx.Remove(dialogVisualSettingsFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"
)

type (
	dialogVisualSettingsModel interface {
		Insert2(ctx context.Context, data *DialogVisualSettings) (sql.Result, error)
	}

	defaultDialogVisualSettingsModel struct {
		db *sqlx.DB
	}

	DialogVisualSettings struct {
		UserId              int64 `db:"user_id" json:"user_id"`
		PeerType            int32 `db:"peer_type" json:"peer_type"`
		PeerId              int64 `db:"peer_id" json:"peer_id"`
		WallpaperId         int64 `db:"wallpaper_id" json:"wallpaper_id"`
		WallpaperOverridden bool  `db:"wallpaper_overridden" json:"wallpaper_overridden"`
		VisualVersion       int64 `db:"visual_version" json:"visual_version"`
	}
)

func newDialogVisualSettingsModel(db *sqlx.DB) *defaultDialogVisualSettingsModel {
	return &defaultDialogVisualSettingsModel{
		db: db,
	}
}

func (m *defaultDialogVisualSettingsModel) Insert2(ctx context.Context, data *DialogVisualSettings) (sql.Result, error) {
	tableName := "dialog_visual_settings"
	query := fmt.Sprintf("insert into `%s` (%s) values (?, ?, ?, ?, ?, ?)", tableName, dialogVisualSettingsRowsExpectAutoSet)

	r, err := m.db.Exec(ctx, query, data.UserId, data.PeerType, data.PeerId, data.WallpaperId, data.WallpaperOverridden, data.VisualVersion)
	if err != nil {
		return nil, fmt.Errorf("dialog_visual_settings.Insert2 exec: %w", err)
	}

	return r, nil
}
