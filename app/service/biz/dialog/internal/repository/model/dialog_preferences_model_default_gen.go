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
	dialogPreferencesFieldNames          = builder.RawFieldNames(&DialogPreferences{})
	dialogPreferencesRows                = strings.Join(dialogPreferencesFieldNames, ",")
	dialogPreferencesRowsExpectAutoSet   = strings.Join(stringx.Remove(dialogPreferencesFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	dialogPreferencesRowsWithPlaceHolder = strings.Join(stringx.Remove(dialogPreferencesFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"
)

type (
	dialogPreferencesModel interface {
		Insert2(ctx context.Context, data *DialogPreferences) (sql.Result, error)

		FindOneByUserIdPeerDialogId(ctx context.Context, userId int64, peerDialogId int64) (*DialogPreferences, error)
	}

	defaultDialogPreferencesModel struct {
		db *sqlx.DB
	}

	DialogPreferences struct {
		UserId             int64 `db:"user_id" json:"user_id"`
		PeerType           int32 `db:"peer_type" json:"peer_type"`
		PeerId             int64 `db:"peer_id" json:"peer_id"`
		PeerDialogId       int64 `db:"peer_dialog_id" json:"peer_dialog_id"`
		FolderId           int32 `db:"folder_id" json:"folder_id"`
		MainPinnedOrder    int64 `db:"main_pinned_order" json:"main_pinned_order"`
		FolderPinnedOrder  int64 `db:"folder_pinned_order" json:"folder_pinned_order"`
		PreferencesVersion int64 `db:"preferences_version" json:"preferences_version"`
	}
)

func newDialogPreferencesModel(db *sqlx.DB) *defaultDialogPreferencesModel {
	return &defaultDialogPreferencesModel{
		db: db,
	}
}

func (m *defaultDialogPreferencesModel) Insert2(ctx context.Context, data *DialogPreferences) (sql.Result, error) {
	tableName := "dialog_preferences"
	query := fmt.Sprintf("insert into `%s` (%s) values (?, ?, ?, ?, ?, ?, ?, ?)", tableName, dialogPreferencesRowsExpectAutoSet)

	r, err := m.db.Exec(ctx, query, data.UserId, data.PeerType, data.PeerId, data.PeerDialogId, data.FolderId, data.MainPinnedOrder, data.FolderPinnedOrder, data.PreferencesVersion)
	if err != nil {
		return nil, fmt.Errorf("dialog_preferences.Insert2 exec: %w", err)
	}

	return r, nil
}

func (m *defaultDialogPreferencesModel) FindOneByUserIdPeerDialogId(ctx context.Context, userId int64, peerDialogId int64) (*DialogPreferences, error) {
	tableName := "dialog_preferences"
	query := fmt.Sprintf("select %s from %s where user_id = ? AND peer_dialog_id = ? limit 1", dialogPreferencesRows, tableName)
	var resp DialogPreferences

	err := m.db.QueryRowPartial(ctx, &resp, query, userId, peerDialogId)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return nil, &NotFoundError{
				Resource: "dialog_preferences",
				Key:      fmt.Sprintf("user_id=%v,peer_dialog_id=%v", userId, peerDialogId),
				Cause:    err,
			}
		}
		return nil, fmt.Errorf("dialog_preferences.FindOneByUserIdPeerDialogId: %w", err)
	}

	return &resp, nil
}
