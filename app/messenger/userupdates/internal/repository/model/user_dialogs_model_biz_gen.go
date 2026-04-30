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

type bizUserDialogsModel interface {
	InsertOrUpdateMessageEvent(ctx context.Context, data *UserDialogs) (lastInsertId, rowsAffected int64, err error)
	SelectByUserPeer(ctx context.Context, userId int64, peerType int32, peerId int64) (*UserDialogs, error)
	SelectByUser(ctx context.Context, userId int64, limit int32) ([]UserDialogs, error)
	SelectByUserWithCB(ctx context.Context, userId int64, limit int32, cb func(sz, i int, v *UserDialogs)) ([]UserDialogs, error)
	UpdateReadState(ctx context.Context, unreadCount int32, unreadMentionsCount int32, readInboxMaxPeerSeq int64, readOutboxMaxPeerSeq int64, userId int64, peerType int32, peerId int64) (rowsAffected int64, err error)
	UpdatePinned(ctx context.Context, pinned bool, userId int64, peerType int32, peerId int64) (rowsAffected int64, err error)
}

type UserDialogsTxModel interface {
	InsertOrUpdateMessageEvent(data *UserDialogs) (lastInsertId, rowsAffected int64, err error)
	SelectByUserPeer(userId int64, peerType int32, peerId int64) (*UserDialogs, error)
	SelectByUser(userId int64, limit int32) ([]UserDialogs, error)
	UpdateReadState(unreadCount int32, unreadMentionsCount int32, readInboxMaxPeerSeq int64, readOutboxMaxPeerSeq int64, userId int64, peerType int32, peerId int64) (rowsAffected int64, err error)
	UpdatePinned(pinned bool, userId int64, peerType int32, peerId int64) (rowsAffected int64, err error)
}

type defaultUserDialogsTxModel struct {
	tx *sqlx.Tx
}

func NewUserDialogsTxModel(tx *sqlx.Tx) UserDialogsTxModel {
	return &defaultUserDialogsTxModel{tx: tx}
}

// InsertOrUpdateMessageEvent
// insert into user_dialogs(user_id, peer_type, peer_id, top_peer_seq, top_canonical_message_id, top_message_date, unread_count, unread_mentions_count, read_inbox_max_peer_seq, read_outbox_max_peer_seq, pinned, folder_id, dialog_schema_version, dialog_payload) values (:user_id, :peer_type, :peer_id, :top_peer_seq, :top_canonical_message_id, :top_message_date, :unread_count, :unread_mentions_count, :read_inbox_max_peer_seq, :read_outbox_max_peer_seq, :pinned, :folder_id, :dialog_schema_version, :dialog_payload) on duplicate key update top_peer_seq = values(top_peer_seq), top_canonical_message_id = values(top_canonical_message_id), top_message_date = values(top_message_date), unread_count = unread_count + values(unread_count), unread_mentions_count = unread_mentions_count + values(unread_mentions_count), dialog_schema_version = values(dialog_schema_version), dialog_payload = values(dialog_payload)
func (m *defaultUserDialogsModel) InsertOrUpdateMessageEvent(ctx context.Context, data *UserDialogs) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into user_dialogs(user_id, peer_type, peer_id, top_peer_seq, top_canonical_message_id, top_message_date, unread_count, unread_mentions_count, read_inbox_max_peer_seq, read_outbox_max_peer_seq, pinned, folder_id, dialog_schema_version, dialog_payload) values (:user_id, :peer_type, :peer_id, :top_peer_seq, :top_canonical_message_id, :top_message_date, :unread_count, :unread_mentions_count, :read_inbox_max_peer_seq, :read_outbox_max_peer_seq, :pinned, :folder_id, :dialog_schema_version, :dialog_payload) on duplicate key update top_peer_seq = values(top_peer_seq), top_canonical_message_id = values(top_canonical_message_id), top_message_date = values(top_message_date), unread_count = unread_count + values(unread_count), unread_mentions_count = unread_mentions_count + values(unread_mentions_count), dialog_schema_version = values(dialog_schema_version), dialog_payload = values(dialog_payload)"
		r     sql.Result
	)

	r, err = m.db.NamedExec(ctx, query, data)
	if err != nil {
		err = fmt.Errorf("user_dialogs.InsertOrUpdateMessageEvent named exec: %w", err)
		return
	}

	lastInsertId, err = r.LastInsertId()
	if err != nil {
		err = fmt.Errorf("user_dialogs.InsertOrUpdateMessageEvent last insert id: %w", err)
		return
	}
	rowsAffected, err = r.RowsAffected()
	if err != nil {
		err = fmt.Errorf("user_dialogs.InsertOrUpdateMessageEvent rows affected: %w", err)
	}

	return

}

// InsertOrUpdateMessageEvent
// insert into user_dialogs(user_id, peer_type, peer_id, top_peer_seq, top_canonical_message_id, top_message_date, unread_count, unread_mentions_count, read_inbox_max_peer_seq, read_outbox_max_peer_seq, pinned, folder_id, dialog_schema_version, dialog_payload) values (:user_id, :peer_type, :peer_id, :top_peer_seq, :top_canonical_message_id, :top_message_date, :unread_count, :unread_mentions_count, :read_inbox_max_peer_seq, :read_outbox_max_peer_seq, :pinned, :folder_id, :dialog_schema_version, :dialog_payload) on duplicate key update top_peer_seq = values(top_peer_seq), top_canonical_message_id = values(top_canonical_message_id), top_message_date = values(top_message_date), unread_count = unread_count + values(unread_count), unread_mentions_count = unread_mentions_count + values(unread_mentions_count), dialog_schema_version = values(dialog_schema_version), dialog_payload = values(dialog_payload)
func (m *defaultUserDialogsTxModel) InsertOrUpdateMessageEvent(data *UserDialogs) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into user_dialogs(user_id, peer_type, peer_id, top_peer_seq, top_canonical_message_id, top_message_date, unread_count, unread_mentions_count, read_inbox_max_peer_seq, read_outbox_max_peer_seq, pinned, folder_id, dialog_schema_version, dialog_payload) values (:user_id, :peer_type, :peer_id, :top_peer_seq, :top_canonical_message_id, :top_message_date, :unread_count, :unread_mentions_count, :read_inbox_max_peer_seq, :read_outbox_max_peer_seq, :pinned, :folder_id, :dialog_schema_version, :dialog_payload) on duplicate key update top_peer_seq = values(top_peer_seq), top_canonical_message_id = values(top_canonical_message_id), top_message_date = values(top_message_date), unread_count = unread_count + values(unread_count), unread_mentions_count = unread_mentions_count + values(unread_mentions_count), dialog_schema_version = values(dialog_schema_version), dialog_payload = values(dialog_payload)"
		r     sql.Result
	)

	r, err = m.tx.NamedExec(query, data)
	if err != nil {
		err = fmt.Errorf("user_dialogs.InsertOrUpdateMessageEvent named exec: %w", err)
		return
	}

	lastInsertId, err = r.LastInsertId()
	if err != nil {
		err = fmt.Errorf("user_dialogs.InsertOrUpdateMessageEvent last insert id: %w", err)
		return
	}
	rowsAffected, err = r.RowsAffected()
	if err != nil {
		err = fmt.Errorf("user_dialogs.InsertOrUpdateMessageEvent rows affected: %w", err)
	}

	return
}

// SelectByUserPeer
// select user_id, peer_type, peer_id, top_peer_seq, top_canonical_message_id, top_message_date, unread_count, unread_mentions_count, read_inbox_max_peer_seq, read_outbox_max_peer_seq, pinned, folder_id, dialog_schema_version, dialog_payload from user_dialogs where user_id = :user_id and peer_type = :peer_type and peer_id = :peer_id limit 1
func (m *defaultUserDialogsModel) SelectByUserPeer(ctx context.Context, userId int64, peerType int32, peerId int64) (rValue *UserDialogs, err error) {

	var (
		query = "select user_id, peer_type, peer_id, top_peer_seq, top_canonical_message_id, top_message_date, unread_count, unread_mentions_count, read_inbox_max_peer_seq, read_outbox_max_peer_seq, pinned, folder_id, dialog_schema_version, dialog_payload from user_dialogs where user_id = ? and peer_type = ? and peer_id = ? limit 1"
		do    = &UserDialogs{}
	)
	err = m.db.QueryRowPartial(ctx, do, query, userId, peerType, peerId)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return nil, &NotFoundError{
				Resource: "user_dialogs",
				Key:      fmt.Sprintf("user_id=%v,peer_type=%v,peer_id=%v", userId, peerType, peerId),
				Cause:    err,
			}
		}
		err = fmt.Errorf("user_dialogs.SelectByUserPeer: %w", err)
		return
	} else {
		rValue = do
	}

	return
}

// SelectByUserPeer
// select user_id, peer_type, peer_id, top_peer_seq, top_canonical_message_id, top_message_date, unread_count, unread_mentions_count, read_inbox_max_peer_seq, read_outbox_max_peer_seq, pinned, folder_id, dialog_schema_version, dialog_payload from user_dialogs where user_id = :user_id and peer_type = :peer_type and peer_id = :peer_id limit 1
func (m *defaultUserDialogsTxModel) SelectByUserPeer(userId int64, peerType int32, peerId int64) (rValue *UserDialogs, err error) {
	var (
		query = "select user_id, peer_type, peer_id, top_peer_seq, top_canonical_message_id, top_message_date, unread_count, unread_mentions_count, read_inbox_max_peer_seq, read_outbox_max_peer_seq, pinned, folder_id, dialog_schema_version, dialog_payload from user_dialogs where user_id = ? and peer_type = ? and peer_id = ? limit 1"
		do    = &UserDialogs{}
	)
	err = m.tx.QueryRowPartial(do, query, userId, peerType, peerId)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return nil, &NotFoundError{
				Resource: "user_dialogs",
				Key:      fmt.Sprintf("user_id=%v,peer_type=%v,peer_id=%v", userId, peerType, peerId),
				Cause:    err,
			}
		}
		err = fmt.Errorf("user_dialogs.SelectByUserPeer: %w", err)
		return
	}
	rValue = do

	return
}

// SelectByUser
// select user_id, peer_type, peer_id, top_peer_seq, top_canonical_message_id, top_message_date, unread_count, unread_mentions_count, read_inbox_max_peer_seq, read_outbox_max_peer_seq, pinned, folder_id, dialog_schema_version, dialog_payload from user_dialogs where user_id = :user_id order by pinned desc, top_message_date desc limit :limit
func (m *defaultUserDialogsModel) SelectByUser(ctx context.Context, userId int64, limit int32) (rList []UserDialogs, err error) {
	var (
		query  = "select user_id, peer_type, peer_id, top_peer_seq, top_canonical_message_id, top_message_date, unread_count, unread_mentions_count, read_inbox_max_peer_seq, read_outbox_max_peer_seq, pinned, folder_id, dialog_schema_version, dialog_payload from user_dialogs where user_id = ? order by pinned desc, top_message_date desc limit ?"
		values []UserDialogs
	)
	err = m.db.QueryRowsPartial(ctx, &values, query, userId, limit)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			rList = []UserDialogs{}
			err = nil
			return
		}
		err = fmt.Errorf("user_dialogs.SelectByUser: %w", err)
		return
	}

	rList = values

	return
}

// SelectByUser
// select user_id, peer_type, peer_id, top_peer_seq, top_canonical_message_id, top_message_date, unread_count, unread_mentions_count, read_inbox_max_peer_seq, read_outbox_max_peer_seq, pinned, folder_id, dialog_schema_version, dialog_payload from user_dialogs where user_id = :user_id order by pinned desc, top_message_date desc limit :limit
func (m *defaultUserDialogsTxModel) SelectByUser(userId int64, limit int32) (rList []UserDialogs, err error) {
	var (
		query  = "select user_id, peer_type, peer_id, top_peer_seq, top_canonical_message_id, top_message_date, unread_count, unread_mentions_count, read_inbox_max_peer_seq, read_outbox_max_peer_seq, pinned, folder_id, dialog_schema_version, dialog_payload from user_dialogs where user_id = ? order by pinned desc, top_message_date desc limit ?"
		values []UserDialogs
	)
	err = m.tx.QueryRowsPartial(&values, query, userId, limit)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			rList = []UserDialogs{}
			err = nil
			return
		}
		err = fmt.Errorf("user_dialogs.SelectByUser: %w", err)
		return
	}

	rList = values

	return
}

// SelectByUserWithCB
// select user_id, peer_type, peer_id, top_peer_seq, top_canonical_message_id, top_message_date, unread_count, unread_mentions_count, read_inbox_max_peer_seq, read_outbox_max_peer_seq, pinned, folder_id, dialog_schema_version, dialog_payload from user_dialogs where user_id = :user_id order by pinned desc, top_message_date desc limit :limit
func (m *defaultUserDialogsModel) SelectByUserWithCB(ctx context.Context, userId int64, limit int32, cb func(sz, i int, v *UserDialogs)) (rList []UserDialogs, err error) {
	var (
		query  = "select user_id, peer_type, peer_id, top_peer_seq, top_canonical_message_id, top_message_date, unread_count, unread_mentions_count, read_inbox_max_peer_seq, read_outbox_max_peer_seq, pinned, folder_id, dialog_schema_version, dialog_payload from user_dialogs where user_id = ? order by pinned desc, top_message_date desc limit ?"
		values []UserDialogs
	)
	err = m.db.QueryRowsPartial(ctx, &values, query, userId, limit)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			rList = []UserDialogs{}
			err = nil
			return
		}
		err = fmt.Errorf("user_dialogs.SelectByUserWithCB: %w", err)
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

// UpdateReadState
// update user_dialogs set unread_count = :unread_count, unread_mentions_count = :unread_mentions_count, read_inbox_max_peer_seq = :read_inbox_max_peer_seq, read_outbox_max_peer_seq = :read_outbox_max_peer_seq where user_id = :user_id and peer_type = :peer_type and peer_id = :peer_id
func (m *defaultUserDialogsModel) UpdateReadState(ctx context.Context, unreadCount int32, unreadMentionsCount int32, readInboxMaxPeerSeq int64, readOutboxMaxPeerSeq int64, userId int64, peerType int32, peerId int64) (rowsAffected int64, err error) {

	var (
		query   = "update user_dialogs set unread_count = ?, unread_mentions_count = ?, read_inbox_max_peer_seq = ?, read_outbox_max_peer_seq = ? where user_id = ? and peer_type = ? and peer_id = ?"
		rResult sql.Result
	)

	rResult, err = m.db.Exec(ctx, query, unreadCount, unreadMentionsCount, readInboxMaxPeerSeq, readOutboxMaxPeerSeq, userId, peerType, peerId)

	if err != nil {
		err = fmt.Errorf("user_dialogs.UpdateReadState exec: %w", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		err = fmt.Errorf("user_dialogs.UpdateReadState rows affected: %w", err)
		return
	}

	return
}

// UpdateReadState
// update user_dialogs set unread_count = :unread_count, unread_mentions_count = :unread_mentions_count, read_inbox_max_peer_seq = :read_inbox_max_peer_seq, read_outbox_max_peer_seq = :read_outbox_max_peer_seq where user_id = :user_id and peer_type = :peer_type and peer_id = :peer_id
func (m *defaultUserDialogsTxModel) UpdateReadState(unreadCount int32, unreadMentionsCount int32, readInboxMaxPeerSeq int64, readOutboxMaxPeerSeq int64, userId int64, peerType int32, peerId int64) (rowsAffected int64, err error) {
	var (
		query   = "update user_dialogs set unread_count = ?, unread_mentions_count = ?, read_inbox_max_peer_seq = ?, read_outbox_max_peer_seq = ? where user_id = ? and peer_type = ? and peer_id = ?"
		rResult sql.Result
	)
	rResult, err = m.tx.Exec(query, unreadCount, unreadMentionsCount, readInboxMaxPeerSeq, readOutboxMaxPeerSeq, userId, peerType, peerId)

	if err != nil {
		err = fmt.Errorf("user_dialogs.UpdateReadState exec: %w", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		err = fmt.Errorf("user_dialogs.UpdateReadState rows affected: %w", err)
		return
	}

	return
}

// UpdatePinned
// update user_dialogs set pinned = :pinned where user_id = :user_id and peer_type = :peer_type and peer_id = :peer_id
func (m *defaultUserDialogsModel) UpdatePinned(ctx context.Context, pinned bool, userId int64, peerType int32, peerId int64) (rowsAffected int64, err error) {

	var (
		query   = "update user_dialogs set pinned = ? where user_id = ? and peer_type = ? and peer_id = ?"
		rResult sql.Result
	)

	rResult, err = m.db.Exec(ctx, query, pinned, userId, peerType, peerId)

	if err != nil {
		err = fmt.Errorf("user_dialogs.UpdatePinned exec: %w", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		err = fmt.Errorf("user_dialogs.UpdatePinned rows affected: %w", err)
		return
	}

	return
}

// UpdatePinned
// update user_dialogs set pinned = :pinned where user_id = :user_id and peer_type = :peer_type and peer_id = :peer_id
func (m *defaultUserDialogsTxModel) UpdatePinned(pinned bool, userId int64, peerType int32, peerId int64) (rowsAffected int64, err error) {
	var (
		query   = "update user_dialogs set pinned = ? where user_id = ? and peer_type = ? and peer_id = ?"
		rResult sql.Result
	)
	rResult, err = m.tx.Exec(query, pinned, userId, peerType, peerId)

	if err != nil {
		err = fmt.Errorf("user_dialogs.UpdatePinned exec: %w", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		err = fmt.Errorf("user_dialogs.UpdatePinned rows affected: %w", err)
		return
	}

	return
}
