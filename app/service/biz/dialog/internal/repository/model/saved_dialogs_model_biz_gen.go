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

type bizSavedDialogsModel interface {
	InsertOrUpdate(ctx context.Context, data *SavedDialogs) (lastInsertId, rowsAffected int64, err error)
	UpsertTopFromMessage(ctx context.Context, data *SavedDialogs) (lastInsertId, rowsAffected int64, err error)
	Select(ctx context.Context, userId int64, peerType int32, peerId int64) (*SavedDialogs, error)
	SelectPinnedDialogs(ctx context.Context, userId int64) ([]SavedDialogs, error)
	SelectPinnedDialogsWithCB(ctx context.Context, userId int64, cb func(sz, i int, v *SavedDialogs)) ([]SavedDialogs, error)
	SelectDialogs(ctx context.Context, userId int64, topMessageDate int64, limit int32) ([]SavedDialogs, error)
	SelectDialogsWithCB(ctx context.Context, userId int64, topMessageDate int64, limit int32, cb func(sz, i int, v *SavedDialogs)) ([]SavedDialogs, error)
	UpdateUserUnPinned(ctx context.Context, userId int64) (rowsAffected int64, err error)
	UpdateUserPeerPinned(ctx context.Context, pinned bool, pinOrder int64, userId int64, peerType int32, peerId int64) (rowsAffected int64, err error)
	SelectUnpinnedDialogs(ctx context.Context, userId int64, topMessageDate int64, limit int32) ([]SavedDialogs, error)
	SelectUnpinnedDialogsWithCB(ctx context.Context, userId int64, topMessageDate int64, limit int32, cb func(sz, i int, v *SavedDialogs)) ([]SavedDialogs, error)
	ClearDuplicatePinOrder(ctx context.Context, userId int64, pinOrder int64, peerType int32, peerId int64) (rowsAffected int64, err error)
}

type SavedDialogsTxModel interface {
	InsertOrUpdate(data *SavedDialogs) (lastInsertId, rowsAffected int64, err error)
	UpsertTopFromMessage(data *SavedDialogs) (lastInsertId, rowsAffected int64, err error)
	Select(userId int64, peerType int32, peerId int64) (*SavedDialogs, error)
	SelectPinnedDialogs(userId int64) ([]SavedDialogs, error)
	SelectDialogs(userId int64, topMessageDate int64, limit int32) ([]SavedDialogs, error)
	UpdateUserUnPinned(userId int64) (rowsAffected int64, err error)
	UpdateUserPeerPinned(pinned bool, pinOrder int64, userId int64, peerType int32, peerId int64) (rowsAffected int64, err error)
	SelectUnpinnedDialogs(userId int64, topMessageDate int64, limit int32) ([]SavedDialogs, error)
	ClearDuplicatePinOrder(userId int64, pinOrder int64, peerType int32, peerId int64) (rowsAffected int64, err error)
}

type defaultSavedDialogsTxModel struct {
	tx *sqlx.Tx
}

func NewSavedDialogsTxModel(tx *sqlx.Tx) SavedDialogsTxModel {
	return &defaultSavedDialogsTxModel{tx: tx}
}

// InsertOrUpdate
// insert into saved_dialogs(user_id, peer_type, peer_id, top_peer_seq, top_canonical_message_id, top_message_date, pinned, pin_order, deleted, saved_schema_version, saved_payload) values (:user_id, :peer_type, :peer_id, :top_peer_seq, :top_canonical_message_id, :top_message_date, :pinned, :pin_order, :deleted, :saved_schema_version, :saved_payload) on duplicate key update top_peer_seq = values(top_peer_seq), top_canonical_message_id = values(top_canonical_message_id), top_message_date = values(top_message_date), deleted = values(deleted), saved_schema_version = values(saved_schema_version), saved_payload = values(saved_payload)
func (m *defaultSavedDialogsModel) InsertOrUpdate(ctx context.Context, data *SavedDialogs) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into saved_dialogs(user_id, peer_type, peer_id, top_peer_seq, top_canonical_message_id, top_message_date, pinned, pin_order, deleted, saved_schema_version, saved_payload) values (:user_id, :peer_type, :peer_id, :top_peer_seq, :top_canonical_message_id, :top_message_date, :pinned, :pin_order, :deleted, :saved_schema_version, :saved_payload) on duplicate key update top_peer_seq = values(top_peer_seq), top_canonical_message_id = values(top_canonical_message_id), top_message_date = values(top_message_date), deleted = values(deleted), saved_schema_version = values(saved_schema_version), saved_payload = values(saved_payload)"
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

// InsertOrUpdate
// insert into saved_dialogs(user_id, peer_type, peer_id, top_peer_seq, top_canonical_message_id, top_message_date, pinned, pin_order, deleted, saved_schema_version, saved_payload) values (:user_id, :peer_type, :peer_id, :top_peer_seq, :top_canonical_message_id, :top_message_date, :pinned, :pin_order, :deleted, :saved_schema_version, :saved_payload) on duplicate key update top_peer_seq = values(top_peer_seq), top_canonical_message_id = values(top_canonical_message_id), top_message_date = values(top_message_date), deleted = values(deleted), saved_schema_version = values(saved_schema_version), saved_payload = values(saved_payload)
func (m *defaultSavedDialogsTxModel) InsertOrUpdate(data *SavedDialogs) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into saved_dialogs(user_id, peer_type, peer_id, top_peer_seq, top_canonical_message_id, top_message_date, pinned, pin_order, deleted, saved_schema_version, saved_payload) values (:user_id, :peer_type, :peer_id, :top_peer_seq, :top_canonical_message_id, :top_message_date, :pinned, :pin_order, :deleted, :saved_schema_version, :saved_payload) on duplicate key update top_peer_seq = values(top_peer_seq), top_canonical_message_id = values(top_canonical_message_id), top_message_date = values(top_message_date), deleted = values(deleted), saved_schema_version = values(saved_schema_version), saved_payload = values(saved_payload)"
		r     sql.Result
	)

	r, err = m.tx.NamedExec(query, data)
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

// UpsertTopFromMessage
// insert into saved_dialogs(user_id, peer_type, peer_id, top_peer_seq, top_canonical_message_id, top_message_date, pinned, pin_order, deleted, saved_schema_version, saved_payload) values (:user_id, :peer_type, :peer_id, :top_peer_seq, :top_canonical_message_id, :top_message_date, 0, 0, 0, :saved_schema_version, :saved_payload) on duplicate key update top_peer_seq = if(top_message_date <= values(top_message_date), values(top_peer_seq), top_peer_seq), top_canonical_message_id = if(top_message_date <= values(top_message_date), values(top_canonical_message_id), top_canonical_message_id), top_message_date = if(top_message_date <= values(top_message_date), values(top_message_date), top_message_date), deleted = 0, saved_schema_version = if(top_message_date <= values(top_message_date), values(saved_schema_version), saved_schema_version), saved_payload = if(top_message_date <= values(top_message_date), values(saved_payload), saved_payload)
func (m *defaultSavedDialogsModel) UpsertTopFromMessage(ctx context.Context, data *SavedDialogs) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into saved_dialogs(user_id, peer_type, peer_id, top_peer_seq, top_canonical_message_id, top_message_date, pinned, pin_order, deleted, saved_schema_version, saved_payload) values (:user_id, :peer_type, :peer_id, :top_peer_seq, :top_canonical_message_id, :top_message_date, 0, 0, 0, :saved_schema_version, :saved_payload) on duplicate key update top_peer_seq = if(top_message_date <= values(top_message_date), values(top_peer_seq), top_peer_seq), top_canonical_message_id = if(top_message_date <= values(top_message_date), values(top_canonical_message_id), top_canonical_message_id), top_message_date = if(top_message_date <= values(top_message_date), values(top_message_date), top_message_date), deleted = 0, saved_schema_version = if(top_message_date <= values(top_message_date), values(saved_schema_version), saved_schema_version), saved_payload = if(top_message_date <= values(top_message_date), values(saved_payload), saved_payload)"
		r     sql.Result
	)

	r, err = m.db.NamedExec(ctx, query, data)
	if err != nil {
		err = fmt.Errorf("saved_dialogs.UpsertTopFromMessage named exec: %w", err)
		return
	}

	lastInsertId, err = r.LastInsertId()
	if err != nil {
		err = fmt.Errorf("saved_dialogs.UpsertTopFromMessage last insert id: %w", err)
		return
	}
	rowsAffected, err = r.RowsAffected()
	if err != nil {
		err = fmt.Errorf("saved_dialogs.UpsertTopFromMessage rows affected: %w", err)
	}

	return

}

// UpsertTopFromMessage
// insert into saved_dialogs(user_id, peer_type, peer_id, top_peer_seq, top_canonical_message_id, top_message_date, pinned, pin_order, deleted, saved_schema_version, saved_payload) values (:user_id, :peer_type, :peer_id, :top_peer_seq, :top_canonical_message_id, :top_message_date, 0, 0, 0, :saved_schema_version, :saved_payload) on duplicate key update top_peer_seq = if(top_message_date <= values(top_message_date), values(top_peer_seq), top_peer_seq), top_canonical_message_id = if(top_message_date <= values(top_message_date), values(top_canonical_message_id), top_canonical_message_id), top_message_date = if(top_message_date <= values(top_message_date), values(top_message_date), top_message_date), deleted = 0, saved_schema_version = if(top_message_date <= values(top_message_date), values(saved_schema_version), saved_schema_version), saved_payload = if(top_message_date <= values(top_message_date), values(saved_payload), saved_payload)
func (m *defaultSavedDialogsTxModel) UpsertTopFromMessage(data *SavedDialogs) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into saved_dialogs(user_id, peer_type, peer_id, top_peer_seq, top_canonical_message_id, top_message_date, pinned, pin_order, deleted, saved_schema_version, saved_payload) values (:user_id, :peer_type, :peer_id, :top_peer_seq, :top_canonical_message_id, :top_message_date, 0, 0, 0, :saved_schema_version, :saved_payload) on duplicate key update top_peer_seq = if(top_message_date <= values(top_message_date), values(top_peer_seq), top_peer_seq), top_canonical_message_id = if(top_message_date <= values(top_message_date), values(top_canonical_message_id), top_canonical_message_id), top_message_date = if(top_message_date <= values(top_message_date), values(top_message_date), top_message_date), deleted = 0, saved_schema_version = if(top_message_date <= values(top_message_date), values(saved_schema_version), saved_schema_version), saved_payload = if(top_message_date <= values(top_message_date), values(saved_payload), saved_payload)"
		r     sql.Result
	)

	r, err = m.tx.NamedExec(query, data)
	if err != nil {
		err = fmt.Errorf("saved_dialogs.UpsertTopFromMessage named exec: %w", err)
		return
	}

	lastInsertId, err = r.LastInsertId()
	if err != nil {
		err = fmt.Errorf("saved_dialogs.UpsertTopFromMessage last insert id: %w", err)
		return
	}
	rowsAffected, err = r.RowsAffected()
	if err != nil {
		err = fmt.Errorf("saved_dialogs.UpsertTopFromMessage rows affected: %w", err)
	}

	return
}

// Select
// select user_id, peer_type, peer_id, top_peer_seq, top_canonical_message_id, top_message_date, pinned, pin_order, deleted, saved_schema_version, saved_payload from saved_dialogs where user_id = :user_id and peer_type = :peer_type and peer_id = :peer_id and deleted = 0 limit 1
func (m *defaultSavedDialogsModel) Select(ctx context.Context, userId int64, peerType int32, peerId int64) (rValue *SavedDialogs, err error) {

	var (
		query = "select user_id, peer_type, peer_id, top_peer_seq, top_canonical_message_id, top_message_date, pinned, pin_order, deleted, saved_schema_version, saved_payload from saved_dialogs where user_id = ? and peer_type = ? and peer_id = ? and deleted = 0 limit 1"
		do    = &SavedDialogs{}
	)
	err = m.db.QueryRowPartial(ctx, do, query, userId, peerType, peerId)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return nil, &NotFoundError{
				Resource: "saved_dialogs",
				Key:      fmt.Sprintf("user_id=%v,peer_type=%v,peer_id=%v", userId, peerType, peerId),
				Cause:    err,
			}
		}
		err = fmt.Errorf("saved_dialogs.Select: %w", err)
		return
	} else {
		rValue = do
	}

	return
}

// Select
// select user_id, peer_type, peer_id, top_peer_seq, top_canonical_message_id, top_message_date, pinned, pin_order, deleted, saved_schema_version, saved_payload from saved_dialogs where user_id = :user_id and peer_type = :peer_type and peer_id = :peer_id and deleted = 0 limit 1
func (m *defaultSavedDialogsTxModel) Select(userId int64, peerType int32, peerId int64) (rValue *SavedDialogs, err error) {
	var (
		query = "select user_id, peer_type, peer_id, top_peer_seq, top_canonical_message_id, top_message_date, pinned, pin_order, deleted, saved_schema_version, saved_payload from saved_dialogs where user_id = ? and peer_type = ? and peer_id = ? and deleted = 0 limit 1"
		do    = &SavedDialogs{}
	)
	err = m.tx.QueryRowPartial(do, query, userId, peerType, peerId)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return nil, &NotFoundError{
				Resource: "saved_dialogs",
				Key:      fmt.Sprintf("user_id=%v,peer_type=%v,peer_id=%v", userId, peerType, peerId),
				Cause:    err,
			}
		}
		err = fmt.Errorf("saved_dialogs.Select: %w", err)
		return
	}
	rValue = do

	return
}

// SelectPinnedDialogs
// select user_id, peer_type, peer_id, top_peer_seq, top_canonical_message_id, top_message_date, pinned, pin_order, deleted, saved_schema_version, saved_payload from saved_dialogs where user_id = :user_id and pinned = 1 and deleted = 0 order by pin_order asc
func (m *defaultSavedDialogsModel) SelectPinnedDialogs(ctx context.Context, userId int64) (rList []SavedDialogs, err error) {
	var (
		query  = "select user_id, peer_type, peer_id, top_peer_seq, top_canonical_message_id, top_message_date, pinned, pin_order, deleted, saved_schema_version, saved_payload from saved_dialogs where user_id = ? and pinned = 1 and deleted = 0 order by pin_order asc"
		values []SavedDialogs
	)
	err = m.db.QueryRowsPartial(ctx, &values, query, userId)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			rList = []SavedDialogs{}
			err = nil
			return
		}
		err = fmt.Errorf("saved_dialogs.SelectPinnedDialogs: %w", err)
		return
	}

	rList = values

	return
}

// SelectPinnedDialogs
// select user_id, peer_type, peer_id, top_peer_seq, top_canonical_message_id, top_message_date, pinned, pin_order, deleted, saved_schema_version, saved_payload from saved_dialogs where user_id = :user_id and pinned = 1 and deleted = 0 order by pin_order asc
func (m *defaultSavedDialogsTxModel) SelectPinnedDialogs(userId int64) (rList []SavedDialogs, err error) {
	var (
		query  = "select user_id, peer_type, peer_id, top_peer_seq, top_canonical_message_id, top_message_date, pinned, pin_order, deleted, saved_schema_version, saved_payload from saved_dialogs where user_id = ? and pinned = 1 and deleted = 0 order by pin_order asc"
		values []SavedDialogs
	)
	err = m.tx.QueryRowsPartial(&values, query, userId)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			rList = []SavedDialogs{}
			err = nil
			return
		}
		err = fmt.Errorf("saved_dialogs.SelectPinnedDialogs: %w", err)
		return
	}

	rList = values

	return
}

// SelectPinnedDialogsWithCB
// select user_id, peer_type, peer_id, top_peer_seq, top_canonical_message_id, top_message_date, pinned, pin_order, deleted, saved_schema_version, saved_payload from saved_dialogs where user_id = :user_id and pinned = 1 and deleted = 0 order by pin_order asc
func (m *defaultSavedDialogsModel) SelectPinnedDialogsWithCB(ctx context.Context, userId int64, cb func(sz, i int, v *SavedDialogs)) (rList []SavedDialogs, err error) {
	var (
		query  = "select user_id, peer_type, peer_id, top_peer_seq, top_canonical_message_id, top_message_date, pinned, pin_order, deleted, saved_schema_version, saved_payload from saved_dialogs where user_id = ? and pinned = 1 and deleted = 0 order by pin_order asc"
		values []SavedDialogs
	)
	err = m.db.QueryRowsPartial(ctx, &values, query, userId)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			rList = []SavedDialogs{}
			err = nil
			return
		}
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

// SelectDialogs
// select user_id, peer_type, peer_id, top_peer_seq, top_canonical_message_id, top_message_date, pinned, pin_order, deleted, saved_schema_version, saved_payload from saved_dialogs where user_id = :user_id and deleted = 0 and top_message_date < :top_message_date order by top_message_date desc limit :limit
func (m *defaultSavedDialogsModel) SelectDialogs(ctx context.Context, userId int64, topMessageDate int64, limit int32) (rList []SavedDialogs, err error) {
	var (
		query  = "select user_id, peer_type, peer_id, top_peer_seq, top_canonical_message_id, top_message_date, pinned, pin_order, deleted, saved_schema_version, saved_payload from saved_dialogs where user_id = ? and deleted = 0 and top_message_date < ? order by top_message_date desc limit ?"
		values []SavedDialogs
	)
	err = m.db.QueryRowsPartial(ctx, &values, query, userId, topMessageDate, limit)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			rList = []SavedDialogs{}
			err = nil
			return
		}
		err = fmt.Errorf("saved_dialogs.SelectDialogs: %w", err)
		return
	}

	rList = values

	return
}

// SelectDialogs
// select user_id, peer_type, peer_id, top_peer_seq, top_canonical_message_id, top_message_date, pinned, pin_order, deleted, saved_schema_version, saved_payload from saved_dialogs where user_id = :user_id and deleted = 0 and top_message_date < :top_message_date order by top_message_date desc limit :limit
func (m *defaultSavedDialogsTxModel) SelectDialogs(userId int64, topMessageDate int64, limit int32) (rList []SavedDialogs, err error) {
	var (
		query  = "select user_id, peer_type, peer_id, top_peer_seq, top_canonical_message_id, top_message_date, pinned, pin_order, deleted, saved_schema_version, saved_payload from saved_dialogs where user_id = ? and deleted = 0 and top_message_date < ? order by top_message_date desc limit ?"
		values []SavedDialogs
	)
	err = m.tx.QueryRowsPartial(&values, query, userId, topMessageDate, limit)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			rList = []SavedDialogs{}
			err = nil
			return
		}
		err = fmt.Errorf("saved_dialogs.SelectDialogs: %w", err)
		return
	}

	rList = values

	return
}

// SelectDialogsWithCB
// select user_id, peer_type, peer_id, top_peer_seq, top_canonical_message_id, top_message_date, pinned, pin_order, deleted, saved_schema_version, saved_payload from saved_dialogs where user_id = :user_id and deleted = 0 and top_message_date < :top_message_date order by top_message_date desc limit :limit
func (m *defaultSavedDialogsModel) SelectDialogsWithCB(ctx context.Context, userId int64, topMessageDate int64, limit int32, cb func(sz, i int, v *SavedDialogs)) (rList []SavedDialogs, err error) {
	var (
		query  = "select user_id, peer_type, peer_id, top_peer_seq, top_canonical_message_id, top_message_date, pinned, pin_order, deleted, saved_schema_version, saved_payload from saved_dialogs where user_id = ? and deleted = 0 and top_message_date < ? order by top_message_date desc limit ?"
		values []SavedDialogs
	)
	err = m.db.QueryRowsPartial(ctx, &values, query, userId, topMessageDate, limit)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			rList = []SavedDialogs{}
			err = nil
			return
		}
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
// update saved_dialogs set pinned = 0, pin_order = 0 where user_id = :user_id and pinned = 1 and deleted = 0
func (m *defaultSavedDialogsModel) UpdateUserUnPinned(ctx context.Context, userId int64) (rowsAffected int64, err error) {

	var (
		query   = "update saved_dialogs set pinned = 0, pin_order = 0 where user_id = ? and pinned = 1 and deleted = 0"
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
		return
	}

	return
}

// UpdateUserUnPinned
// update saved_dialogs set pinned = 0, pin_order = 0 where user_id = :user_id and pinned = 1 and deleted = 0
func (m *defaultSavedDialogsTxModel) UpdateUserUnPinned(userId int64) (rowsAffected int64, err error) {
	var (
		query   = "update saved_dialogs set pinned = 0, pin_order = 0 where user_id = ? and pinned = 1 and deleted = 0"
		rResult sql.Result
	)
	rResult, err = m.tx.Exec(query, userId)

	if err != nil {
		err = fmt.Errorf("saved_dialogs.UpdateUserUnPinned exec: %w", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		err = fmt.Errorf("saved_dialogs.UpdateUserUnPinned rows affected: %w", err)
		return
	}

	return
}

// UpdateUserPeerPinned
// update saved_dialogs set pinned = :pinned, pin_order = :pin_order where user_id = :user_id and peer_type = :peer_type and peer_id = :peer_id
func (m *defaultSavedDialogsModel) UpdateUserPeerPinned(ctx context.Context, pinned bool, pinOrder int64, userId int64, peerType int32, peerId int64) (rowsAffected int64, err error) {

	var (
		query   = "update saved_dialogs set pinned = ?, pin_order = ? where user_id = ? and peer_type = ? and peer_id = ?"
		rResult sql.Result
	)

	rResult, err = m.db.Exec(ctx, query, pinned, pinOrder, userId, peerType, peerId)

	if err != nil {
		err = fmt.Errorf("saved_dialogs.UpdateUserPeerPinned exec: %w", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		err = fmt.Errorf("saved_dialogs.UpdateUserPeerPinned rows affected: %w", err)
		return
	}

	return
}

// UpdateUserPeerPinned
// update saved_dialogs set pinned = :pinned, pin_order = :pin_order where user_id = :user_id and peer_type = :peer_type and peer_id = :peer_id
func (m *defaultSavedDialogsTxModel) UpdateUserPeerPinned(pinned bool, pinOrder int64, userId int64, peerType int32, peerId int64) (rowsAffected int64, err error) {
	var (
		query   = "update saved_dialogs set pinned = ?, pin_order = ? where user_id = ? and peer_type = ? and peer_id = ?"
		rResult sql.Result
	)
	rResult, err = m.tx.Exec(query, pinned, pinOrder, userId, peerType, peerId)

	if err != nil {
		err = fmt.Errorf("saved_dialogs.UpdateUserPeerPinned exec: %w", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		err = fmt.Errorf("saved_dialogs.UpdateUserPeerPinned rows affected: %w", err)
		return
	}

	return
}

// SelectUnpinnedDialogs
// select user_id, peer_type, peer_id, top_peer_seq, top_canonical_message_id, top_message_date, pinned, pin_order, deleted, saved_schema_version, saved_payload from saved_dialogs where user_id = :user_id and deleted = 0 and pinned = 0 and top_message_date < :top_message_date order by top_message_date desc limit :limit
func (m *defaultSavedDialogsModel) SelectUnpinnedDialogs(ctx context.Context, userId int64, topMessageDate int64, limit int32) (rList []SavedDialogs, err error) {
	var (
		query  = "select user_id, peer_type, peer_id, top_peer_seq, top_canonical_message_id, top_message_date, pinned, pin_order, deleted, saved_schema_version, saved_payload from saved_dialogs where user_id = ? and deleted = 0 and pinned = 0 and top_message_date < ? order by top_message_date desc limit ?"
		values []SavedDialogs
	)
	err = m.db.QueryRowsPartial(ctx, &values, query, userId, topMessageDate, limit)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			rList = []SavedDialogs{}
			err = nil
			return
		}
		err = fmt.Errorf("saved_dialogs.SelectUnpinnedDialogs: %w", err)
		return
	}

	rList = values

	return
}

// SelectUnpinnedDialogs
// select user_id, peer_type, peer_id, top_peer_seq, top_canonical_message_id, top_message_date, pinned, pin_order, deleted, saved_schema_version, saved_payload from saved_dialogs where user_id = :user_id and deleted = 0 and pinned = 0 and top_message_date < :top_message_date order by top_message_date desc limit :limit
func (m *defaultSavedDialogsTxModel) SelectUnpinnedDialogs(userId int64, topMessageDate int64, limit int32) (rList []SavedDialogs, err error) {
	var (
		query  = "select user_id, peer_type, peer_id, top_peer_seq, top_canonical_message_id, top_message_date, pinned, pin_order, deleted, saved_schema_version, saved_payload from saved_dialogs where user_id = ? and deleted = 0 and pinned = 0 and top_message_date < ? order by top_message_date desc limit ?"
		values []SavedDialogs
	)
	err = m.tx.QueryRowsPartial(&values, query, userId, topMessageDate, limit)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			rList = []SavedDialogs{}
			err = nil
			return
		}
		err = fmt.Errorf("saved_dialogs.SelectUnpinnedDialogs: %w", err)
		return
	}

	rList = values

	return
}

// SelectUnpinnedDialogsWithCB
// select user_id, peer_type, peer_id, top_peer_seq, top_canonical_message_id, top_message_date, pinned, pin_order, deleted, saved_schema_version, saved_payload from saved_dialogs where user_id = :user_id and deleted = 0 and pinned = 0 and top_message_date < :top_message_date order by top_message_date desc limit :limit
func (m *defaultSavedDialogsModel) SelectUnpinnedDialogsWithCB(ctx context.Context, userId int64, topMessageDate int64, limit int32, cb func(sz, i int, v *SavedDialogs)) (rList []SavedDialogs, err error) {
	var (
		query  = "select user_id, peer_type, peer_id, top_peer_seq, top_canonical_message_id, top_message_date, pinned, pin_order, deleted, saved_schema_version, saved_payload from saved_dialogs where user_id = ? and deleted = 0 and pinned = 0 and top_message_date < ? order by top_message_date desc limit ?"
		values []SavedDialogs
	)
	err = m.db.QueryRowsPartial(ctx, &values, query, userId, topMessageDate, limit)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			rList = []SavedDialogs{}
			err = nil
			return
		}
		err = fmt.Errorf("saved_dialogs.SelectUnpinnedDialogsWithCB: %w", err)
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

// ClearDuplicatePinOrder
// update saved_dialogs set pinned = 0, pin_order = 0 where user_id = :user_id and pin_order = :pin_order and not (peer_type = :peer_type and peer_id = :peer_id) and deleted = 0
func (m *defaultSavedDialogsModel) ClearDuplicatePinOrder(ctx context.Context, userId int64, pinOrder int64, peerType int32, peerId int64) (rowsAffected int64, err error) {

	var (
		query   = "update saved_dialogs set pinned = 0, pin_order = 0 where user_id = ? and pin_order = ? and not (peer_type = ? and peer_id = ?) and deleted = 0"
		rResult sql.Result
	)

	rResult, err = m.db.Exec(ctx, query, userId, pinOrder, peerType, peerId)

	if err != nil {
		err = fmt.Errorf("saved_dialogs.ClearDuplicatePinOrder exec: %w", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		err = fmt.Errorf("saved_dialogs.ClearDuplicatePinOrder rows affected: %w", err)
		return
	}

	return
}

// ClearDuplicatePinOrder
// update saved_dialogs set pinned = 0, pin_order = 0 where user_id = :user_id and pin_order = :pin_order and not (peer_type = :peer_type and peer_id = :peer_id) and deleted = 0
func (m *defaultSavedDialogsTxModel) ClearDuplicatePinOrder(userId int64, pinOrder int64, peerType int32, peerId int64) (rowsAffected int64, err error) {
	var (
		query   = "update saved_dialogs set pinned = 0, pin_order = 0 where user_id = ? and pin_order = ? and not (peer_type = ? and peer_id = ?) and deleted = 0"
		rResult sql.Result
	)
	rResult, err = m.tx.Exec(query, userId, pinOrder, peerType, peerId)

	if err != nil {
		err = fmt.Errorf("saved_dialogs.ClearDuplicatePinOrder exec: %w", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		err = fmt.Errorf("saved_dialogs.ClearDuplicatePinOrder rows affected: %w", err)
		return
	}

	return
}
