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
	SelectMainPinned(ctx context.Context, userId int64) ([]DialogPreferences, error)
	SelectMainPinnedWithCB(ctx context.Context, userId int64, cb func(sz, i int, v *DialogPreferences)) ([]DialogPreferences, error)
	SelectFolderPinned(ctx context.Context, userId int64, folderId int32) ([]DialogPreferences, error)
	SelectFolderPinnedWithCB(ctx context.Context, userId int64, folderId int32, cb func(sz, i int, v *DialogPreferences)) ([]DialogPreferences, error)
	ClearMainPinned(ctx context.Context, userId int64) (rowsAffected int64, err error)
	ClearMainPinnedExcept(ctx context.Context, userId int64, idList []int64) (rowsAffected int64, err error)
	ClearFolderPinned(ctx context.Context, userId int64, folderId int32) (rowsAffected int64, err error)
	ClearFolderPinnedExcept(ctx context.Context, userId int64, folderId int32, idList []int64) (rowsAffected int64, err error)
}

type DialogPreferencesTxModel interface {
	InsertOrUpdate(data *DialogPreferences) (lastInsertId, rowsAffected int64, err error)
	UpsertMainPin(data *DialogPreferences) (lastInsertId, rowsAffected int64, err error)
	UpsertFolderPin(data *DialogPreferences) (lastInsertId, rowsAffected int64, err error)
	UpsertFolderMembership(data *DialogPreferences) (lastInsertId, rowsAffected int64, err error)
	SelectByUserPeer(userId int64, peerType int32, peerId int64) (*DialogPreferences, error)
	SelectMainPinned(userId int64) ([]DialogPreferences, error)
	SelectFolderPinned(userId int64, folderId int32) ([]DialogPreferences, error)
	ClearMainPinned(userId int64) (rowsAffected int64, err error)
	ClearMainPinnedExcept(userId int64, idList []int64) (rowsAffected int64, err error)
	ClearFolderPinned(userId int64, folderId int32) (rowsAffected int64, err error)
	ClearFolderPinnedExcept(userId int64, folderId int32, idList []int64) (rowsAffected int64, err error)
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

// SelectMainPinned
// select user_id, peer_type, peer_id, peer_dialog_id, folder_id, main_pinned_order, folder_pinned_order, preferences_version from dialog_preferences where user_id = :user_id and main_pinned_order > 0 order by main_pinned_order desc
func (m *defaultDialogPreferencesModel) SelectMainPinned(ctx context.Context, userId int64) (rList []DialogPreferences, err error) {
	var (
		query  = "select user_id, peer_type, peer_id, peer_dialog_id, folder_id, main_pinned_order, folder_pinned_order, preferences_version from dialog_preferences where user_id = ? and main_pinned_order > 0 order by main_pinned_order desc"
		values []DialogPreferences
	)
	err = m.db.QueryRowsPartial(ctx, &values, query, userId)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			rList = []DialogPreferences{}
			err = nil
			return
		}
		err = fmt.Errorf("dialog_preferences.SelectMainPinned: %w", err)
		return
	}

	rList = values

	return
}

// SelectMainPinned
// select user_id, peer_type, peer_id, peer_dialog_id, folder_id, main_pinned_order, folder_pinned_order, preferences_version from dialog_preferences where user_id = :user_id and main_pinned_order > 0 order by main_pinned_order desc
func (m *defaultDialogPreferencesTxModel) SelectMainPinned(userId int64) (rList []DialogPreferences, err error) {
	var (
		query  = "select user_id, peer_type, peer_id, peer_dialog_id, folder_id, main_pinned_order, folder_pinned_order, preferences_version from dialog_preferences where user_id = ? and main_pinned_order > 0 order by main_pinned_order desc"
		values []DialogPreferences
	)
	err = m.tx.QueryRowsPartial(&values, query, userId)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			rList = []DialogPreferences{}
			err = nil
			return
		}
		err = fmt.Errorf("dialog_preferences.SelectMainPinned: %w", err)
		return
	}

	rList = values

	return
}

// SelectMainPinnedWithCB
// select user_id, peer_type, peer_id, peer_dialog_id, folder_id, main_pinned_order, folder_pinned_order, preferences_version from dialog_preferences where user_id = :user_id and main_pinned_order > 0 order by main_pinned_order desc
func (m *defaultDialogPreferencesModel) SelectMainPinnedWithCB(ctx context.Context, userId int64, cb func(sz, i int, v *DialogPreferences)) (rList []DialogPreferences, err error) {
	var (
		query  = "select user_id, peer_type, peer_id, peer_dialog_id, folder_id, main_pinned_order, folder_pinned_order, preferences_version from dialog_preferences where user_id = ? and main_pinned_order > 0 order by main_pinned_order desc"
		values []DialogPreferences
	)
	err = m.db.QueryRowsPartial(ctx, &values, query, userId)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			rList = []DialogPreferences{}
			err = nil
			return
		}
		err = fmt.Errorf("dialog_preferences.SelectMainPinnedWithCB: %w", err)
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

// SelectFolderPinned
// select user_id, peer_type, peer_id, peer_dialog_id, folder_id, main_pinned_order, folder_pinned_order, preferences_version from dialog_preferences where user_id = :user_id and folder_id = :folder_id and folder_pinned_order > 0 order by folder_pinned_order desc
func (m *defaultDialogPreferencesModel) SelectFolderPinned(ctx context.Context, userId int64, folderId int32) (rList []DialogPreferences, err error) {
	var (
		query  = "select user_id, peer_type, peer_id, peer_dialog_id, folder_id, main_pinned_order, folder_pinned_order, preferences_version from dialog_preferences where user_id = ? and folder_id = ? and folder_pinned_order > 0 order by folder_pinned_order desc"
		values []DialogPreferences
	)
	err = m.db.QueryRowsPartial(ctx, &values, query, userId, folderId)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			rList = []DialogPreferences{}
			err = nil
			return
		}
		err = fmt.Errorf("dialog_preferences.SelectFolderPinned: %w", err)
		return
	}

	rList = values

	return
}

// SelectFolderPinned
// select user_id, peer_type, peer_id, peer_dialog_id, folder_id, main_pinned_order, folder_pinned_order, preferences_version from dialog_preferences where user_id = :user_id and folder_id = :folder_id and folder_pinned_order > 0 order by folder_pinned_order desc
func (m *defaultDialogPreferencesTxModel) SelectFolderPinned(userId int64, folderId int32) (rList []DialogPreferences, err error) {
	var (
		query  = "select user_id, peer_type, peer_id, peer_dialog_id, folder_id, main_pinned_order, folder_pinned_order, preferences_version from dialog_preferences where user_id = ? and folder_id = ? and folder_pinned_order > 0 order by folder_pinned_order desc"
		values []DialogPreferences
	)
	err = m.tx.QueryRowsPartial(&values, query, userId, folderId)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			rList = []DialogPreferences{}
			err = nil
			return
		}
		err = fmt.Errorf("dialog_preferences.SelectFolderPinned: %w", err)
		return
	}

	rList = values

	return
}

// SelectFolderPinnedWithCB
// select user_id, peer_type, peer_id, peer_dialog_id, folder_id, main_pinned_order, folder_pinned_order, preferences_version from dialog_preferences where user_id = :user_id and folder_id = :folder_id and folder_pinned_order > 0 order by folder_pinned_order desc
func (m *defaultDialogPreferencesModel) SelectFolderPinnedWithCB(ctx context.Context, userId int64, folderId int32, cb func(sz, i int, v *DialogPreferences)) (rList []DialogPreferences, err error) {
	var (
		query  = "select user_id, peer_type, peer_id, peer_dialog_id, folder_id, main_pinned_order, folder_pinned_order, preferences_version from dialog_preferences where user_id = ? and folder_id = ? and folder_pinned_order > 0 order by folder_pinned_order desc"
		values []DialogPreferences
	)
	err = m.db.QueryRowsPartial(ctx, &values, query, userId, folderId)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			rList = []DialogPreferences{}
			err = nil
			return
		}
		err = fmt.Errorf("dialog_preferences.SelectFolderPinnedWithCB: %w", err)
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

// ClearMainPinned
// update dialog_preferences set main_pinned_order = 0 where user_id = :user_id and main_pinned_order > 0
func (m *defaultDialogPreferencesModel) ClearMainPinned(ctx context.Context, userId int64) (rowsAffected int64, err error) {

	var (
		query   = "update dialog_preferences set main_pinned_order = 0 where user_id = ? and main_pinned_order > 0"
		rResult sql.Result
	)

	rResult, err = m.db.Exec(ctx, query, userId)

	if err != nil {
		err = fmt.Errorf("dialog_preferences.ClearMainPinned exec: %w", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		err = fmt.Errorf("dialog_preferences.ClearMainPinned rows affected: %w", err)
		return
	}

	return
}

// ClearMainPinned
// update dialog_preferences set main_pinned_order = 0 where user_id = :user_id and main_pinned_order > 0
func (m *defaultDialogPreferencesTxModel) ClearMainPinned(userId int64) (rowsAffected int64, err error) {
	var (
		query   = "update dialog_preferences set main_pinned_order = 0 where user_id = ? and main_pinned_order > 0"
		rResult sql.Result
	)
	rResult, err = m.tx.Exec(query, userId)

	if err != nil {
		err = fmt.Errorf("dialog_preferences.ClearMainPinned exec: %w", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		err = fmt.Errorf("dialog_preferences.ClearMainPinned rows affected: %w", err)
		return
	}

	return
}

// ClearMainPinnedExcept
// update dialog_preferences set main_pinned_order = 0 where user_id = :user_id and main_pinned_order > 0 and peer_dialog_id not in (:idList)
func (m *defaultDialogPreferencesModel) ClearMainPinnedExcept(ctx context.Context, userId int64, idList []int64) (rowsAffected int64, err error) {

	var (
		query   = fmt.Sprintf("update dialog_preferences set main_pinned_order = 0 where user_id = ? and main_pinned_order > 0 and peer_dialog_id not in (%s)", sqlx.InInt64List(idList))
		rResult sql.Result
	)

	if len(idList) == 0 {
		return
	}

	rResult, err = m.db.Exec(ctx, query, userId)

	if err != nil {
		err = fmt.Errorf("dialog_preferences.ClearMainPinnedExcept exec: %w", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		err = fmt.Errorf("dialog_preferences.ClearMainPinnedExcept rows affected: %w", err)
		return
	}

	return
}

// ClearMainPinnedExcept
// update dialog_preferences set main_pinned_order = 0 where user_id = :user_id and main_pinned_order > 0 and peer_dialog_id not in (:idList)
func (m *defaultDialogPreferencesTxModel) ClearMainPinnedExcept(userId int64, idList []int64) (rowsAffected int64, err error) {
	var (
		query   = fmt.Sprintf("update dialog_preferences set main_pinned_order = 0 where user_id = ? and main_pinned_order > 0 and peer_dialog_id not in (%s)", sqlx.InInt64List(idList))
		rResult sql.Result
	)

	if len(idList) == 0 {
		return
	}

	rResult, err = m.tx.Exec(query, userId)

	if err != nil {
		err = fmt.Errorf("dialog_preferences.ClearMainPinnedExcept exec: %w", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		err = fmt.Errorf("dialog_preferences.ClearMainPinnedExcept rows affected: %w", err)
		return
	}

	return
}

// ClearFolderPinned
// update dialog_preferences set folder_pinned_order = 0 where user_id = :user_id and folder_id = :folder_id and folder_pinned_order > 0
func (m *defaultDialogPreferencesModel) ClearFolderPinned(ctx context.Context, userId int64, folderId int32) (rowsAffected int64, err error) {

	var (
		query   = "update dialog_preferences set folder_pinned_order = 0 where user_id = ? and folder_id = ? and folder_pinned_order > 0"
		rResult sql.Result
	)

	rResult, err = m.db.Exec(ctx, query, userId, folderId)

	if err != nil {
		err = fmt.Errorf("dialog_preferences.ClearFolderPinned exec: %w", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		err = fmt.Errorf("dialog_preferences.ClearFolderPinned rows affected: %w", err)
		return
	}

	return
}

// ClearFolderPinned
// update dialog_preferences set folder_pinned_order = 0 where user_id = :user_id and folder_id = :folder_id and folder_pinned_order > 0
func (m *defaultDialogPreferencesTxModel) ClearFolderPinned(userId int64, folderId int32) (rowsAffected int64, err error) {
	var (
		query   = "update dialog_preferences set folder_pinned_order = 0 where user_id = ? and folder_id = ? and folder_pinned_order > 0"
		rResult sql.Result
	)
	rResult, err = m.tx.Exec(query, userId, folderId)

	if err != nil {
		err = fmt.Errorf("dialog_preferences.ClearFolderPinned exec: %w", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		err = fmt.Errorf("dialog_preferences.ClearFolderPinned rows affected: %w", err)
		return
	}

	return
}

// ClearFolderPinnedExcept
// update dialog_preferences set folder_pinned_order = 0 where user_id = :user_id and folder_id = :folder_id and folder_pinned_order > 0 and peer_dialog_id not in (:idList)
func (m *defaultDialogPreferencesModel) ClearFolderPinnedExcept(ctx context.Context, userId int64, folderId int32, idList []int64) (rowsAffected int64, err error) {

	var (
		query   = fmt.Sprintf("update dialog_preferences set folder_pinned_order = 0 where user_id = ? and folder_id = ? and folder_pinned_order > 0 and peer_dialog_id not in (%s)", sqlx.InInt64List(idList))
		rResult sql.Result
	)

	if len(idList) == 0 {
		return
	}

	rResult, err = m.db.Exec(ctx, query, userId, folderId)

	if err != nil {
		err = fmt.Errorf("dialog_preferences.ClearFolderPinnedExcept exec: %w", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		err = fmt.Errorf("dialog_preferences.ClearFolderPinnedExcept rows affected: %w", err)
		return
	}

	return
}

// ClearFolderPinnedExcept
// update dialog_preferences set folder_pinned_order = 0 where user_id = :user_id and folder_id = :folder_id and folder_pinned_order > 0 and peer_dialog_id not in (:idList)
func (m *defaultDialogPreferencesTxModel) ClearFolderPinnedExcept(userId int64, folderId int32, idList []int64) (rowsAffected int64, err error) {
	var (
		query   = fmt.Sprintf("update dialog_preferences set folder_pinned_order = 0 where user_id = ? and folder_id = ? and folder_pinned_order > 0 and peer_dialog_id not in (%s)", sqlx.InInt64List(idList))
		rResult sql.Result
	)

	if len(idList) == 0 {
		return
	}

	rResult, err = m.tx.Exec(query, userId, folderId)

	if err != nil {
		err = fmt.Errorf("dialog_preferences.ClearFolderPinnedExcept exec: %w", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		err = fmt.Errorf("dialog_preferences.ClearFolderPinnedExcept rows affected: %w", err)
		return
	}

	return
}
