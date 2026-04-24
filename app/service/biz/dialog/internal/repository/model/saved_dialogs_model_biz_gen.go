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

type (
	bizSavedDialogsModel interface {
		InsertOrUpdate(ctx context.Context, data *SavedDialogs) (lastInsertId, rowsAffected int64, err error)
		InsertOrUpdateTx(tx *sqlx.Tx, data *SavedDialogs) (lastInsertId, rowsAffected int64, err error)

		Select(ctx context.Context, userId int64, peerType int32, peerId int64) (*SavedDialogs, error)

		SelectPinnedDialogs(ctx context.Context, userId int64) ([]SavedDialogs, error)
		SelectPinnedDialogsWithCB(ctx context.Context, userId int64, cb func(sz, i int, v *SavedDialogs)) ([]SavedDialogs, error)

		SelectExcludePinnedDialogs(ctx context.Context, userId int64, topMessage int32, limit int32) ([]SavedDialogs, error)
		SelectExcludePinnedDialogsWithCB(ctx context.Context, userId int64, topMessage int32, limit int32, cb func(sz, i int, v *SavedDialogs)) ([]SavedDialogs, error)

		SelectDialogs(ctx context.Context, userId int64, topMessage int32, limit int32) ([]SavedDialogs, error)
		SelectDialogsWithCB(ctx context.Context, userId int64, topMessage int32, limit int32, cb func(sz, i int, v *SavedDialogs)) ([]SavedDialogs, error)

		UpdateUserUnPinned(ctx context.Context, userId int64) (rowsAffected int64, err error)
		UpdateUserUnPinnedTx(tx *sqlx.Tx, userId int64) (rowsAffected int64, err error)

		UpdateUserPeerPinned(ctx context.Context, pinned int64, userId int64, peerType int32, peerId int64) (rowsAffected int64, err error)
		UpdateUserPeerPinnedTx(tx *sqlx.Tx, pinned int64, userId int64, peerType int32, peerId int64) (rowsAffected int64, err error)
	}
)

// InsertOrUpdate
// insert into saved_dialogs(user_id, peer_type, peer_id, pinned, top_message) values (:user_id, :peer_type, :peer_id, 0, :top_message) on duplicate key update top_message = values(top_message), deleted = 0
func (m *defaultSavedDialogsModel) InsertOrUpdate(ctx context.Context, data *SavedDialogs) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into saved_dialogs(user_id, peer_type, peer_id, pinned, top_message) values (:user_id, :peer_type, :peer_id, 0, :top_message) on duplicate key update top_message = values(top_message), deleted = 0"
		r     sql.Result
	)

	r, err = m.db.NamedExec(ctx, query, data)
	if err != nil {
		err = fmt.Errorf("saved_dialogs.InsertOrUpdate named exec: %w", err)
		return
	}

	lastInsertId, err = r.LastInsertId()
	if err != nil {
		err = fmt.Errorf("saved_dialogs.InsertOrUpdate last insert id: %w", err)
		return
	}
	rowsAffected, err = r.RowsAffected()
	if err != nil {
		err = fmt.Errorf("saved_dialogs.InsertOrUpdate rows affected: %w", err)
	}

	return

}

// InsertOrUpdateTx
// insert into saved_dialogs(user_id, peer_type, peer_id, pinned, top_message) values (:user_id, :peer_type, :peer_id, 0, :top_message) on duplicate key update top_message = values(top_message), deleted = 0
func (m *defaultSavedDialogsModel) InsertOrUpdateTx(tx *sqlx.Tx, data *SavedDialogs) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into saved_dialogs(user_id, peer_type, peer_id, pinned, top_message) values (:user_id, :peer_type, :peer_id, 0, :top_message) on duplicate key update top_message = values(top_message), deleted = 0"
		r     sql.Result
	)

	r, err = tx.NamedExec(query, data)
	if err != nil {
		err = fmt.Errorf("saved_dialogs.InsertOrUpdateTx named exec: %w", err)
		return
	}

	lastInsertId, err = r.LastInsertId()
	if err != nil {
		err = fmt.Errorf("saved_dialogs.InsertOrUpdateTx last insert id: %w", err)
		return
	}
	rowsAffected, err = r.RowsAffected()
	if err != nil {
		err = fmt.Errorf("saved_dialogs.InsertOrUpdateTx rows affected: %w", err)
	}

	return
}

// Select
// select user_id, peer_type, peer_id, pinned, top_message from saved_dialogs where user_id = :user_id and peer_type = :peer_type and peer_id = :peer_id and deleted = 0
func (m *defaultSavedDialogsModel) Select(ctx context.Context, userId int64, peerType int32, peerId int64) (rValue *SavedDialogs, err error) {

	var (
		query = "select user_id, peer_type, peer_id, pinned, top_message from saved_dialogs where user_id = ? and peer_type = ? and peer_id = ? and deleted = 0"
		do    = &SavedDialogs{}
	)
	err = m.db.QueryRowPartial(ctx, do, query, userId, peerType, peerId)

	if err != nil {
		if !errors.Is(err, sqlx.ErrNotFound) {
			err = fmt.Errorf("saved_dialogs.Select: %w", err)
			return
		} else {
			err = nil
		}
	} else {
		rValue = do
	}

	return
}

// SelectPinnedDialogs
// select user_id, peer_type, peer_id, pinned, top_message from saved_dialogs where user_id = :user_id and pinned > 0 and deleted = 0 order by pinned desc
func (m *defaultSavedDialogsModel) SelectPinnedDialogs(ctx context.Context, userId int64) (rList []SavedDialogs, err error) {
	var (
		query  = "select user_id, peer_type, peer_id, pinned, top_message from saved_dialogs where user_id = ? and pinned > 0 and deleted = 0 order by pinned desc"
		values []SavedDialogs
	)
	err = m.db.QueryRowsPartial(ctx, &values, query, userId)

	if err != nil {
		err = fmt.Errorf("saved_dialogs.SelectPinnedDialogs: %w", err)
		return
	}

	rList = values

	return
}

// SelectPinnedDialogsWithCB
// select user_id, peer_type, peer_id, pinned, top_message from saved_dialogs where user_id = :user_id and pinned > 0 and deleted = 0 order by pinned desc
func (m *defaultSavedDialogsModel) SelectPinnedDialogsWithCB(ctx context.Context, userId int64, cb func(sz, i int, v *SavedDialogs)) (rList []SavedDialogs, err error) {
	var (
		query  = "select user_id, peer_type, peer_id, pinned, top_message from saved_dialogs where user_id = ? and pinned > 0 and deleted = 0 order by pinned desc"
		values []SavedDialogs
	)
	err = m.db.QueryRowsPartial(ctx, &values, query, userId)

	if err != nil {
		err = fmt.Errorf("saved_dialogs.SelectPinnedDialogsWithCB: %w", err)
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

// SelectExcludePinnedDialogs
// select user_id, peer_type, peer_id, pinned, top_message from saved_dialogs where user_id = :user_id and pinned = 0 and top_message < :top_message and deleted = 0 order by top_message desc limit :limit
func (m *defaultSavedDialogsModel) SelectExcludePinnedDialogs(ctx context.Context, userId int64, topMessage int32, limit int32) (rList []SavedDialogs, err error) {
	var (
		query  = "select user_id, peer_type, peer_id, pinned, top_message from saved_dialogs where user_id = ? and pinned = 0 and top_message < ? and deleted = 0 order by top_message desc limit ?"
		values []SavedDialogs
	)
	err = m.db.QueryRowsPartial(ctx, &values, query, userId, topMessage, limit)

	if err != nil {
		err = fmt.Errorf("saved_dialogs.SelectExcludePinnedDialogs: %w", err)
		return
	}

	rList = values

	return
}

// SelectExcludePinnedDialogsWithCB
// select user_id, peer_type, peer_id, pinned, top_message from saved_dialogs where user_id = :user_id and pinned = 0 and top_message < :top_message and deleted = 0 order by top_message desc limit :limit
func (m *defaultSavedDialogsModel) SelectExcludePinnedDialogsWithCB(ctx context.Context, userId int64, topMessage int32, limit int32, cb func(sz, i int, v *SavedDialogs)) (rList []SavedDialogs, err error) {
	var (
		query  = "select user_id, peer_type, peer_id, pinned, top_message from saved_dialogs where user_id = ? and pinned = 0 and top_message < ? and deleted = 0 order by top_message desc limit ?"
		values []SavedDialogs
	)
	err = m.db.QueryRowsPartial(ctx, &values, query, userId, topMessage, limit)

	if err != nil {
		err = fmt.Errorf("saved_dialogs.SelectExcludePinnedDialogsWithCB: %w", err)
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

// SelectDialogs
// select user_id, peer_type, peer_id, pinned, top_message from saved_dialogs where user_id = :user_id and top_message < :top_message and deleted = 0 order by top_message desc limit :limit
func (m *defaultSavedDialogsModel) SelectDialogs(ctx context.Context, userId int64, topMessage int32, limit int32) (rList []SavedDialogs, err error) {
	var (
		query  = "select user_id, peer_type, peer_id, pinned, top_message from saved_dialogs where user_id = ? and top_message < ? and deleted = 0 order by top_message desc limit ?"
		values []SavedDialogs
	)
	err = m.db.QueryRowsPartial(ctx, &values, query, userId, topMessage, limit)

	if err != nil {
		err = fmt.Errorf("saved_dialogs.SelectDialogs: %w", err)
		return
	}

	rList = values

	return
}

// SelectDialogsWithCB
// select user_id, peer_type, peer_id, pinned, top_message from saved_dialogs where user_id = :user_id and top_message < :top_message and deleted = 0 order by top_message desc limit :limit
func (m *defaultSavedDialogsModel) SelectDialogsWithCB(ctx context.Context, userId int64, topMessage int32, limit int32, cb func(sz, i int, v *SavedDialogs)) (rList []SavedDialogs, err error) {
	var (
		query  = "select user_id, peer_type, peer_id, pinned, top_message from saved_dialogs where user_id = ? and top_message < ? and deleted = 0 order by top_message desc limit ?"
		values []SavedDialogs
	)
	err = m.db.QueryRowsPartial(ctx, &values, query, userId, topMessage, limit)

	if err != nil {
		err = fmt.Errorf("saved_dialogs.SelectDialogsWithCB: %w", err)
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

// UpdateUserUnPinned
// update saved_dialogs set pinned = 0 where user_id = :user_id and pinned > 0 and deleted = 0
func (m *defaultSavedDialogsModel) UpdateUserUnPinned(ctx context.Context, userId int64) (rowsAffected int64, err error) {

	var (
		query   = "update saved_dialogs set pinned = 0 where user_id = ? and pinned > 0 and deleted = 0"
		rResult sql.Result
	)

	rResult, err = m.db.Exec(ctx, query, userId)

	if err != nil {
		err = fmt.Errorf("saved_dialogs.UpdateUserUnPinned exec: %w", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		err = fmt.Errorf("saved_dialogs.UpdateUserUnPinned rows affected: %w", err)
	}

	return
}

// UpdateUserUnPinnedTx
// update saved_dialogs set pinned = 0 where user_id = :user_id and pinned > 0 and deleted = 0
func (m *defaultSavedDialogsModel) UpdateUserUnPinnedTx(tx *sqlx.Tx, userId int64) (rowsAffected int64, err error) {
	var (
		query   = "update saved_dialogs set pinned = 0 where user_id = ? and pinned > 0 and deleted = 0"
		rResult sql.Result
	)
	rResult, err = tx.Exec(query, userId)

	if err != nil {
		err = fmt.Errorf("saved_dialogs.UpdateUserUnPinnedTx exec: %w", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		err = fmt.Errorf("saved_dialogs.UpdateUserUnPinnedTx rows affected: %w", err)
	}

	return
}

// UpdateUserPeerPinned
// update saved_dialogs set pinned = :pinned where user_id = :user_id and peer_type = :peer_type and peer_id = :peer_id
func (m *defaultSavedDialogsModel) UpdateUserPeerPinned(ctx context.Context, pinned int64, userId int64, peerType int32, peerId int64) (rowsAffected int64, err error) {

	var (
		query   = "update saved_dialogs set pinned = ? where user_id = ? and peer_type = ? and peer_id = ?"
		rResult sql.Result
	)

	rResult, err = m.db.Exec(ctx, query, pinned, userId, peerType, peerId)

	if err != nil {
		err = fmt.Errorf("saved_dialogs.UpdateUserPeerPinned exec: %w", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		err = fmt.Errorf("saved_dialogs.UpdateUserPeerPinned rows affected: %w", err)
	}

	return
}

// UpdateUserPeerPinnedTx
// update saved_dialogs set pinned = :pinned where user_id = :user_id and peer_type = :peer_type and peer_id = :peer_id
func (m *defaultSavedDialogsModel) UpdateUserPeerPinnedTx(tx *sqlx.Tx, pinned int64, userId int64, peerType int32, peerId int64) (rowsAffected int64, err error) {
	var (
		query   = "update saved_dialogs set pinned = ? where user_id = ? and peer_type = ? and peer_id = ?"
		rResult sql.Result
	)
	rResult, err = tx.Exec(query, pinned, userId, peerType, peerId)

	if err != nil {
		err = fmt.Errorf("saved_dialogs.UpdateUserPeerPinnedTx exec: %w", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		err = fmt.Errorf("saved_dialogs.UpdateUserPeerPinnedTx rows affected: %w", err)
	}

	return
}
