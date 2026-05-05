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

type bizDialogPreferencesModel interface {
	InsertOrUpdate(ctx context.Context, data *DialogPreferences) (lastInsertId, rowsAffected int64, err error)
	UpsertMainPin(ctx context.Context, data *DialogPreferences) (lastInsertId, rowsAffected int64, err error)
	UpsertFolderPin(ctx context.Context, data *DialogPreferences) (lastInsertId, rowsAffected int64, err error)
	UpsertFolderMembership(ctx context.Context, data *DialogPreferences) (lastInsertId, rowsAffected int64, err error)
	SelectByUserPeer(ctx context.Context, userId int64, peerType int32, peerId int64) (*DialogPreferences, error)
}

type DialogPreferencesTxModel interface {
	InsertOrUpdate(data *DialogPreferences) (lastInsertId, rowsAffected int64, err error)
	UpsertMainPin(data *DialogPreferences) (lastInsertId, rowsAffected int64, err error)
	UpsertFolderPin(data *DialogPreferences) (lastInsertId, rowsAffected int64, err error)
	UpsertFolderMembership(data *DialogPreferences) (lastInsertId, rowsAffected int64, err error)
	SelectByUserPeer(userId int64, peerType int32, peerId int64) (*DialogPreferences, error)
}

type defaultDialogPreferencesTxModel struct {
	tx *sqlx.Tx
}

func NewDialogPreferencesTxModel(tx *sqlx.Tx) DialogPreferencesTxModel {
	return &defaultDialogPreferencesTxModel{tx: tx}
}

// InsertOrUpdate
// insert into dialog_preferences(user_id, peer_type, peer_id, peer_dialog_id, folder_id, main_pinned_order, folder_pinned_order, preferences_version) values (:user_id, :peer_type, :peer_id, :peer_dialog_id, :folder_id, :main_pinned_order, :folder_pinned_order, :preferences_version) on duplicate key update folder_id = values(folder_id), main_pinned_order = values(main_pinned_order), folder_pinned_order = values(folder_pinned_order), preferences_version = values(preferences_version)
func (m *defaultDialogPreferencesModel) InsertOrUpdate(ctx context.Context, data *DialogPreferences) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into dialog_preferences(user_id, peer_type, peer_id, peer_dialog_id, folder_id, main_pinned_order, folder_pinned_order, preferences_version) values (:user_id, :peer_type, :peer_id, :peer_dialog_id, :folder_id, :main_pinned_order, :folder_pinned_order, :preferences_version) on duplicate key update folder_id = values(folder_id), main_pinned_order = values(main_pinned_order), folder_pinned_order = values(folder_pinned_order), preferences_version = values(preferences_version)"
		r     sql.Result
	)

	r, err = m.db.NamedExec(ctx, query, data)
	if err != nil {
		err = fmt.Errorf("dialog_preferences.InsertOrUpdate named exec: %w", err)
		return
	}

	lastInsertId, err = r.LastInsertId()
	if err != nil {
		err = fmt.Errorf("dialog_preferences.InsertOrUpdate last insert id: %w", err)
		return
	}
	rowsAffected, err = r.RowsAffected()
	if err != nil {
		err = fmt.Errorf("dialog_preferences.InsertOrUpdate rows affected: %w", err)
	}

	return

}

// InsertOrUpdate
// insert into dialog_preferences(user_id, peer_type, peer_id, peer_dialog_id, folder_id, main_pinned_order, folder_pinned_order, preferences_version) values (:user_id, :peer_type, :peer_id, :peer_dialog_id, :folder_id, :main_pinned_order, :folder_pinned_order, :preferences_version) on duplicate key update folder_id = values(folder_id), main_pinned_order = values(main_pinned_order), folder_pinned_order = values(folder_pinned_order), preferences_version = values(preferences_version)
func (m *defaultDialogPreferencesTxModel) InsertOrUpdate(data *DialogPreferences) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into dialog_preferences(user_id, peer_type, peer_id, peer_dialog_id, folder_id, main_pinned_order, folder_pinned_order, preferences_version) values (:user_id, :peer_type, :peer_id, :peer_dialog_id, :folder_id, :main_pinned_order, :folder_pinned_order, :preferences_version) on duplicate key update folder_id = values(folder_id), main_pinned_order = values(main_pinned_order), folder_pinned_order = values(folder_pinned_order), preferences_version = values(preferences_version)"
		r     sql.Result
	)

	r, err = m.tx.NamedExec(query, data)
	if err != nil {
		err = fmt.Errorf("dialog_preferences.InsertOrUpdate named exec: %w", err)
		return
	}

	lastInsertId, err = r.LastInsertId()
	if err != nil {
		err = fmt.Errorf("dialog_preferences.InsertOrUpdate last insert id: %w", err)
		return
	}
	rowsAffected, err = r.RowsAffected()
	if err != nil {
		err = fmt.Errorf("dialog_preferences.InsertOrUpdate rows affected: %w", err)
	}

	return
}

// UpsertMainPin
// insert into dialog_preferences(user_id, peer_type, peer_id, peer_dialog_id, folder_id, main_pinned_order, folder_pinned_order, preferences_version) values (:user_id, :peer_type, :peer_id, :peer_dialog_id, :folder_id, :main_pinned_order, 0, :preferences_version) on duplicate key update main_pinned_order = values(main_pinned_order), preferences_version = values(preferences_version)
func (m *defaultDialogPreferencesModel) UpsertMainPin(ctx context.Context, data *DialogPreferences) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into dialog_preferences(user_id, peer_type, peer_id, peer_dialog_id, folder_id, main_pinned_order, folder_pinned_order, preferences_version) values (:user_id, :peer_type, :peer_id, :peer_dialog_id, :folder_id, :main_pinned_order, 0, :preferences_version) on duplicate key update main_pinned_order = values(main_pinned_order), preferences_version = values(preferences_version)"
		r     sql.Result
	)

	r, err = m.db.NamedExec(ctx, query, data)
	if err != nil {
		err = fmt.Errorf("dialog_preferences.UpsertMainPin named exec: %w", err)
		return
	}

	lastInsertId, err = r.LastInsertId()
	if err != nil {
		err = fmt.Errorf("dialog_preferences.UpsertMainPin last insert id: %w", err)
		return
	}
	rowsAffected, err = r.RowsAffected()
	if err != nil {
		err = fmt.Errorf("dialog_preferences.UpsertMainPin rows affected: %w", err)
	}

	return

}

// UpsertMainPin
// insert into dialog_preferences(user_id, peer_type, peer_id, peer_dialog_id, folder_id, main_pinned_order, folder_pinned_order, preferences_version) values (:user_id, :peer_type, :peer_id, :peer_dialog_id, :folder_id, :main_pinned_order, 0, :preferences_version) on duplicate key update main_pinned_order = values(main_pinned_order), preferences_version = values(preferences_version)
func (m *defaultDialogPreferencesTxModel) UpsertMainPin(data *DialogPreferences) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into dialog_preferences(user_id, peer_type, peer_id, peer_dialog_id, folder_id, main_pinned_order, folder_pinned_order, preferences_version) values (:user_id, :peer_type, :peer_id, :peer_dialog_id, :folder_id, :main_pinned_order, 0, :preferences_version) on duplicate key update main_pinned_order = values(main_pinned_order), preferences_version = values(preferences_version)"
		r     sql.Result
	)

	r, err = m.tx.NamedExec(query, data)
	if err != nil {
		err = fmt.Errorf("dialog_preferences.UpsertMainPin named exec: %w", err)
		return
	}

	lastInsertId, err = r.LastInsertId()
	if err != nil {
		err = fmt.Errorf("dialog_preferences.UpsertMainPin last insert id: %w", err)
		return
	}
	rowsAffected, err = r.RowsAffected()
	if err != nil {
		err = fmt.Errorf("dialog_preferences.UpsertMainPin rows affected: %w", err)
	}

	return
}

// UpsertFolderPin
// insert into dialog_preferences(user_id, peer_type, peer_id, peer_dialog_id, folder_id, main_pinned_order, folder_pinned_order, preferences_version) values (:user_id, :peer_type, :peer_id, :peer_dialog_id, :folder_id, 0, :folder_pinned_order, :preferences_version) on duplicate key update folder_id = values(folder_id), folder_pinned_order = values(folder_pinned_order), preferences_version = values(preferences_version)
func (m *defaultDialogPreferencesModel) UpsertFolderPin(ctx context.Context, data *DialogPreferences) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into dialog_preferences(user_id, peer_type, peer_id, peer_dialog_id, folder_id, main_pinned_order, folder_pinned_order, preferences_version) values (:user_id, :peer_type, :peer_id, :peer_dialog_id, :folder_id, 0, :folder_pinned_order, :preferences_version) on duplicate key update folder_id = values(folder_id), folder_pinned_order = values(folder_pinned_order), preferences_version = values(preferences_version)"
		r     sql.Result
	)

	r, err = m.db.NamedExec(ctx, query, data)
	if err != nil {
		err = fmt.Errorf("dialog_preferences.UpsertFolderPin named exec: %w", err)
		return
	}

	lastInsertId, err = r.LastInsertId()
	if err != nil {
		err = fmt.Errorf("dialog_preferences.UpsertFolderPin last insert id: %w", err)
		return
	}
	rowsAffected, err = r.RowsAffected()
	if err != nil {
		err = fmt.Errorf("dialog_preferences.UpsertFolderPin rows affected: %w", err)
	}

	return

}

// UpsertFolderPin
// insert into dialog_preferences(user_id, peer_type, peer_id, peer_dialog_id, folder_id, main_pinned_order, folder_pinned_order, preferences_version) values (:user_id, :peer_type, :peer_id, :peer_dialog_id, :folder_id, 0, :folder_pinned_order, :preferences_version) on duplicate key update folder_id = values(folder_id), folder_pinned_order = values(folder_pinned_order), preferences_version = values(preferences_version)
func (m *defaultDialogPreferencesTxModel) UpsertFolderPin(data *DialogPreferences) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into dialog_preferences(user_id, peer_type, peer_id, peer_dialog_id, folder_id, main_pinned_order, folder_pinned_order, preferences_version) values (:user_id, :peer_type, :peer_id, :peer_dialog_id, :folder_id, 0, :folder_pinned_order, :preferences_version) on duplicate key update folder_id = values(folder_id), folder_pinned_order = values(folder_pinned_order), preferences_version = values(preferences_version)"
		r     sql.Result
	)

	r, err = m.tx.NamedExec(query, data)
	if err != nil {
		err = fmt.Errorf("dialog_preferences.UpsertFolderPin named exec: %w", err)
		return
	}

	lastInsertId, err = r.LastInsertId()
	if err != nil {
		err = fmt.Errorf("dialog_preferences.UpsertFolderPin last insert id: %w", err)
		return
	}
	rowsAffected, err = r.RowsAffected()
	if err != nil {
		err = fmt.Errorf("dialog_preferences.UpsertFolderPin rows affected: %w", err)
	}

	return
}

// UpsertFolderMembership
// insert into dialog_preferences(user_id, peer_type, peer_id, peer_dialog_id, folder_id, main_pinned_order, folder_pinned_order, preferences_version) values (:user_id, :peer_type, :peer_id, :peer_dialog_id, :folder_id, 0, 0, :preferences_version) on duplicate key update folder_id = values(folder_id), preferences_version = values(preferences_version)
func (m *defaultDialogPreferencesModel) UpsertFolderMembership(ctx context.Context, data *DialogPreferences) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into dialog_preferences(user_id, peer_type, peer_id, peer_dialog_id, folder_id, main_pinned_order, folder_pinned_order, preferences_version) values (:user_id, :peer_type, :peer_id, :peer_dialog_id, :folder_id, 0, 0, :preferences_version) on duplicate key update folder_id = values(folder_id), preferences_version = values(preferences_version)"
		r     sql.Result
	)

	r, err = m.db.NamedExec(ctx, query, data)
	if err != nil {
		err = fmt.Errorf("dialog_preferences.UpsertFolderMembership named exec: %w", err)
		return
	}

	lastInsertId, err = r.LastInsertId()
	if err != nil {
		err = fmt.Errorf("dialog_preferences.UpsertFolderMembership last insert id: %w", err)
		return
	}
	rowsAffected, err = r.RowsAffected()
	if err != nil {
		err = fmt.Errorf("dialog_preferences.UpsertFolderMembership rows affected: %w", err)
	}

	return

}

// UpsertFolderMembership
// insert into dialog_preferences(user_id, peer_type, peer_id, peer_dialog_id, folder_id, main_pinned_order, folder_pinned_order, preferences_version) values (:user_id, :peer_type, :peer_id, :peer_dialog_id, :folder_id, 0, 0, :preferences_version) on duplicate key update folder_id = values(folder_id), preferences_version = values(preferences_version)
func (m *defaultDialogPreferencesTxModel) UpsertFolderMembership(data *DialogPreferences) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into dialog_preferences(user_id, peer_type, peer_id, peer_dialog_id, folder_id, main_pinned_order, folder_pinned_order, preferences_version) values (:user_id, :peer_type, :peer_id, :peer_dialog_id, :folder_id, 0, 0, :preferences_version) on duplicate key update folder_id = values(folder_id), preferences_version = values(preferences_version)"
		r     sql.Result
	)

	r, err = m.tx.NamedExec(query, data)
	if err != nil {
		err = fmt.Errorf("dialog_preferences.UpsertFolderMembership named exec: %w", err)
		return
	}

	lastInsertId, err = r.LastInsertId()
	if err != nil {
		err = fmt.Errorf("dialog_preferences.UpsertFolderMembership last insert id: %w", err)
		return
	}
	rowsAffected, err = r.RowsAffected()
	if err != nil {
		err = fmt.Errorf("dialog_preferences.UpsertFolderMembership rows affected: %w", err)
	}

	return
}

// SelectByUserPeer
// select user_id, peer_type, peer_id, peer_dialog_id, folder_id, main_pinned_order, folder_pinned_order, preferences_version from dialog_preferences where user_id = :user_id and peer_type = :peer_type and peer_id = :peer_id limit 1
func (m *defaultDialogPreferencesModel) SelectByUserPeer(ctx context.Context, userId int64, peerType int32, peerId int64) (rValue *DialogPreferences, err error) {

	var (
		query = "select user_id, peer_type, peer_id, peer_dialog_id, folder_id, main_pinned_order, folder_pinned_order, preferences_version from dialog_preferences where user_id = ? and peer_type = ? and peer_id = ? limit 1"
		do    = &DialogPreferences{}
	)
	err = m.db.QueryRowPartial(ctx, do, query, userId, peerType, peerId)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return nil, &NotFoundError{
				Resource: "dialog_preferences",
				Key:      fmt.Sprintf("user_id=%v,peer_type=%v,peer_id=%v", userId, peerType, peerId),
				Cause:    err,
			}
		}
		err = fmt.Errorf("dialog_preferences.SelectByUserPeer: %w", err)
		return
	} else {
		rValue = do
	}

	return
}

// SelectByUserPeer
// select user_id, peer_type, peer_id, peer_dialog_id, folder_id, main_pinned_order, folder_pinned_order, preferences_version from dialog_preferences where user_id = :user_id and peer_type = :peer_type and peer_id = :peer_id limit 1
func (m *defaultDialogPreferencesTxModel) SelectByUserPeer(userId int64, peerType int32, peerId int64) (rValue *DialogPreferences, err error) {
	var (
		query = "select user_id, peer_type, peer_id, peer_dialog_id, folder_id, main_pinned_order, folder_pinned_order, preferences_version from dialog_preferences where user_id = ? and peer_type = ? and peer_id = ? limit 1"
		do    = &DialogPreferences{}
	)
	err = m.tx.QueryRowPartial(do, query, userId, peerType, peerId)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return nil, &NotFoundError{
				Resource: "dialog_preferences",
				Key:      fmt.Sprintf("user_id=%v,peer_type=%v,peer_id=%v", userId, peerType, peerId),
				Cause:    err,
			}
		}
		err = fmt.Errorf("dialog_preferences.SelectByUserPeer: %w", err)
		return
	}
	rValue = do

	return
}
