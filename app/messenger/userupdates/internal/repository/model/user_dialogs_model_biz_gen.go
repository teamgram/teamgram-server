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
	SelectByUserPeers(ctx context.Context, userId int64, peerIdList []int64) ([]UserDialogs, error)
	SelectByUserPeersWithCB(ctx context.Context, userId int64, peerIdList []int64, cb func(sz, i int, v *UserDialogs)) ([]UserDialogs, error)
	SelectByUserCursor(ctx context.Context, userId int64, topMessageDate string, topPeerSeq int64, peerType int32, peerId int64, limit int32) ([]UserDialogs, error)
	SelectByUserCursorWithCB(ctx context.Context, userId int64, topMessageDate string, topPeerSeq int64, peerType int32, peerId int64, limit int32, cb func(sz, i int, v *UserDialogs)) ([]UserDialogs, error)
	UpdateReadState(ctx context.Context, unreadCount int32, unreadMentionsCount int32, unreadReactionsCount int32, unreadMark bool, readInboxMaxPeerSeq int64, readOutboxMaxPeerSeq int64, lastPts int64, lastPtsAt string, userId int64, peerType int32, peerId int64) (rowsAffected int64, err error)
	UpdatePinnedMessage(ctx context.Context, pinnedPeerSeq int64, pinnedCanonicalMessageId int64, lastPts int64, lastPtsAt string, userId int64, peerType int32, peerId int64) (rowsAffected int64, err error)
}

type UserDialogsTxModel interface {
	InsertOrUpdateMessageEvent(data *UserDialogs) (lastInsertId, rowsAffected int64, err error)
	SelectByUserPeer(userId int64, peerType int32, peerId int64) (*UserDialogs, error)
	SelectByUserPeers(userId int64, peerIdList []int64) ([]UserDialogs, error)
	SelectByUserCursor(userId int64, topMessageDate string, topPeerSeq int64, peerType int32, peerId int64, limit int32) ([]UserDialogs, error)
	UpdateReadState(unreadCount int32, unreadMentionsCount int32, unreadReactionsCount int32, unreadMark bool, readInboxMaxPeerSeq int64, readOutboxMaxPeerSeq int64, lastPts int64, lastPtsAt string, userId int64, peerType int32, peerId int64) (rowsAffected int64, err error)
	UpdatePinnedMessage(pinnedPeerSeq int64, pinnedCanonicalMessageId int64, lastPts int64, lastPtsAt string, userId int64, peerType int32, peerId int64) (rowsAffected int64, err error)
}

type defaultUserDialogsTxModel struct {
	tx *sqlx.Tx
}

func NewUserDialogsTxModel(tx *sqlx.Tx) UserDialogsTxModel {
	return &defaultUserDialogsTxModel{tx: tx}
}

// InsertOrUpdateMessageEvent
// insert into user_dialogs(user_id, peer_type, peer_id, top_peer_seq, top_canonical_message_id, top_message_date, top_message_status, read_inbox_max_peer_seq, read_outbox_max_peer_seq, unread_count, unread_mentions_count, unread_reactions_count, unread_mark, pinned_peer_seq, pinned_canonical_message_id, has_scheduled, available_min_peer_seq, hidden, deleted_at, last_pts, last_pts_at, dialog_schema_version, dialog_payload) values (:user_id, :peer_type, :peer_id, :top_peer_seq, :top_canonical_message_id, :top_message_date, :top_message_status, :read_inbox_max_peer_seq, :read_outbox_max_peer_seq, :unread_count, :unread_mentions_count, :unread_reactions_count, :unread_mark, :pinned_peer_seq, :pinned_canonical_message_id, :has_scheduled, :available_min_peer_seq, :hidden, :deleted_at, :last_pts, :last_pts_at, :dialog_schema_version, :dialog_payload) on duplicate key update top_peer_seq = values(top_peer_seq), top_canonical_message_id = values(top_canonical_message_id), top_message_date = values(top_message_date), top_message_status = values(top_message_status), unread_count = unread_count + values(unread_count), unread_mentions_count = unread_mentions_count + values(unread_mentions_count), unread_reactions_count = unread_reactions_count + values(unread_reactions_count), unread_mark = values(unread_mark), hidden = values(hidden), deleted_at = values(deleted_at), last_pts = values(last_pts), last_pts_at = values(last_pts_at), dialog_schema_version = values(dialog_schema_version), dialog_payload = values(dialog_payload)
func (m *defaultUserDialogsModel) InsertOrUpdateMessageEvent(ctx context.Context, data *UserDialogs) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into user_dialogs(user_id, peer_type, peer_id, top_peer_seq, top_canonical_message_id, top_message_date, top_message_status, read_inbox_max_peer_seq, read_outbox_max_peer_seq, unread_count, unread_mentions_count, unread_reactions_count, unread_mark, pinned_peer_seq, pinned_canonical_message_id, has_scheduled, available_min_peer_seq, hidden, deleted_at, last_pts, last_pts_at, dialog_schema_version, dialog_payload) values (:user_id, :peer_type, :peer_id, :top_peer_seq, :top_canonical_message_id, :top_message_date, :top_message_status, :read_inbox_max_peer_seq, :read_outbox_max_peer_seq, :unread_count, :unread_mentions_count, :unread_reactions_count, :unread_mark, :pinned_peer_seq, :pinned_canonical_message_id, :has_scheduled, :available_min_peer_seq, :hidden, :deleted_at, :last_pts, :last_pts_at, :dialog_schema_version, :dialog_payload) on duplicate key update top_peer_seq = values(top_peer_seq), top_canonical_message_id = values(top_canonical_message_id), top_message_date = values(top_message_date), top_message_status = values(top_message_status), unread_count = unread_count + values(unread_count), unread_mentions_count = unread_mentions_count + values(unread_mentions_count), unread_reactions_count = unread_reactions_count + values(unread_reactions_count), unread_mark = values(unread_mark), hidden = values(hidden), deleted_at = values(deleted_at), last_pts = values(last_pts), last_pts_at = values(last_pts_at), dialog_schema_version = values(dialog_schema_version), dialog_payload = values(dialog_payload)"
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
// insert into user_dialogs(user_id, peer_type, peer_id, top_peer_seq, top_canonical_message_id, top_message_date, top_message_status, read_inbox_max_peer_seq, read_outbox_max_peer_seq, unread_count, unread_mentions_count, unread_reactions_count, unread_mark, pinned_peer_seq, pinned_canonical_message_id, has_scheduled, available_min_peer_seq, hidden, deleted_at, last_pts, last_pts_at, dialog_schema_version, dialog_payload) values (:user_id, :peer_type, :peer_id, :top_peer_seq, :top_canonical_message_id, :top_message_date, :top_message_status, :read_inbox_max_peer_seq, :read_outbox_max_peer_seq, :unread_count, :unread_mentions_count, :unread_reactions_count, :unread_mark, :pinned_peer_seq, :pinned_canonical_message_id, :has_scheduled, :available_min_peer_seq, :hidden, :deleted_at, :last_pts, :last_pts_at, :dialog_schema_version, :dialog_payload) on duplicate key update top_peer_seq = values(top_peer_seq), top_canonical_message_id = values(top_canonical_message_id), top_message_date = values(top_message_date), top_message_status = values(top_message_status), unread_count = unread_count + values(unread_count), unread_mentions_count = unread_mentions_count + values(unread_mentions_count), unread_reactions_count = unread_reactions_count + values(unread_reactions_count), unread_mark = values(unread_mark), hidden = values(hidden), deleted_at = values(deleted_at), last_pts = values(last_pts), last_pts_at = values(last_pts_at), dialog_schema_version = values(dialog_schema_version), dialog_payload = values(dialog_payload)
func (m *defaultUserDialogsTxModel) InsertOrUpdateMessageEvent(data *UserDialogs) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into user_dialogs(user_id, peer_type, peer_id, top_peer_seq, top_canonical_message_id, top_message_date, top_message_status, read_inbox_max_peer_seq, read_outbox_max_peer_seq, unread_count, unread_mentions_count, unread_reactions_count, unread_mark, pinned_peer_seq, pinned_canonical_message_id, has_scheduled, available_min_peer_seq, hidden, deleted_at, last_pts, last_pts_at, dialog_schema_version, dialog_payload) values (:user_id, :peer_type, :peer_id, :top_peer_seq, :top_canonical_message_id, :top_message_date, :top_message_status, :read_inbox_max_peer_seq, :read_outbox_max_peer_seq, :unread_count, :unread_mentions_count, :unread_reactions_count, :unread_mark, :pinned_peer_seq, :pinned_canonical_message_id, :has_scheduled, :available_min_peer_seq, :hidden, :deleted_at, :last_pts, :last_pts_at, :dialog_schema_version, :dialog_payload) on duplicate key update top_peer_seq = values(top_peer_seq), top_canonical_message_id = values(top_canonical_message_id), top_message_date = values(top_message_date), top_message_status = values(top_message_status), unread_count = unread_count + values(unread_count), unread_mentions_count = unread_mentions_count + values(unread_mentions_count), unread_reactions_count = unread_reactions_count + values(unread_reactions_count), unread_mark = values(unread_mark), hidden = values(hidden), deleted_at = values(deleted_at), last_pts = values(last_pts), last_pts_at = values(last_pts_at), dialog_schema_version = values(dialog_schema_version), dialog_payload = values(dialog_payload)"
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
// select user_id, peer_type, peer_id, top_peer_seq, top_canonical_message_id, top_message_date, top_message_status, read_inbox_max_peer_seq, read_outbox_max_peer_seq, unread_count, unread_mentions_count, unread_reactions_count, unread_mark, pinned_peer_seq, pinned_canonical_message_id, has_scheduled, available_min_peer_seq, hidden, deleted_at, last_pts, last_pts_at, dialog_schema_version, dialog_payload from user_dialogs where user_id = :user_id and peer_type = :peer_type and peer_id = :peer_id limit 1
func (m *defaultUserDialogsModel) SelectByUserPeer(ctx context.Context, userId int64, peerType int32, peerId int64) (rValue *UserDialogs, err error) {

	var (
		query = "select user_id, peer_type, peer_id, top_peer_seq, top_canonical_message_id, top_message_date, top_message_status, read_inbox_max_peer_seq, read_outbox_max_peer_seq, unread_count, unread_mentions_count, unread_reactions_count, unread_mark, pinned_peer_seq, pinned_canonical_message_id, has_scheduled, available_min_peer_seq, hidden, deleted_at, last_pts, last_pts_at, dialog_schema_version, dialog_payload from user_dialogs where user_id = ? and peer_type = ? and peer_id = ? limit 1"
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
// select user_id, peer_type, peer_id, top_peer_seq, top_canonical_message_id, top_message_date, top_message_status, read_inbox_max_peer_seq, read_outbox_max_peer_seq, unread_count, unread_mentions_count, unread_reactions_count, unread_mark, pinned_peer_seq, pinned_canonical_message_id, has_scheduled, available_min_peer_seq, hidden, deleted_at, last_pts, last_pts_at, dialog_schema_version, dialog_payload from user_dialogs where user_id = :user_id and peer_type = :peer_type and peer_id = :peer_id limit 1
func (m *defaultUserDialogsTxModel) SelectByUserPeer(userId int64, peerType int32, peerId int64) (rValue *UserDialogs, err error) {
	var (
		query = "select user_id, peer_type, peer_id, top_peer_seq, top_canonical_message_id, top_message_date, top_message_status, read_inbox_max_peer_seq, read_outbox_max_peer_seq, unread_count, unread_mentions_count, unread_reactions_count, unread_mark, pinned_peer_seq, pinned_canonical_message_id, has_scheduled, available_min_peer_seq, hidden, deleted_at, last_pts, last_pts_at, dialog_schema_version, dialog_payload from user_dialogs where user_id = ? and peer_type = ? and peer_id = ? limit 1"
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

// SelectByUserPeers
// select user_id, peer_type, peer_id, top_peer_seq, top_canonical_message_id, top_message_date, top_message_status, read_inbox_max_peer_seq, read_outbox_max_peer_seq, unread_count, unread_mentions_count, unread_reactions_count, unread_mark, pinned_peer_seq, pinned_canonical_message_id, has_scheduled, available_min_peer_seq, hidden, deleted_at, last_pts, last_pts_at, dialog_schema_version, dialog_payload from user_dialogs where user_id = :user_id and peer_id in (:peerIdList)
func (m *defaultUserDialogsModel) SelectByUserPeers(ctx context.Context, userId int64, peerIdList []int64) (rList []UserDialogs, err error) {
	var (
		query  = fmt.Sprintf("select user_id, peer_type, peer_id, top_peer_seq, top_canonical_message_id, top_message_date, top_message_status, read_inbox_max_peer_seq, read_outbox_max_peer_seq, unread_count, unread_mentions_count, unread_reactions_count, unread_mark, pinned_peer_seq, pinned_canonical_message_id, has_scheduled, available_min_peer_seq, hidden, deleted_at, last_pts, last_pts_at, dialog_schema_version, dialog_payload from user_dialogs where user_id = ? and peer_id in (%s)", sqlx.InInt64List(peerIdList))
		values []UserDialogs
	)
	if len(peerIdList) == 0 {
		rList = []UserDialogs{}
		return
	}

	err = m.db.QueryRowsPartial(ctx, &values, query, userId)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			rList = []UserDialogs{}
			err = nil
			return
		}
		err = fmt.Errorf("user_dialogs.SelectByUserPeers: %w", err)
		return
	}

	rList = values

	return
}

// SelectByUserPeers
// select user_id, peer_type, peer_id, top_peer_seq, top_canonical_message_id, top_message_date, top_message_status, read_inbox_max_peer_seq, read_outbox_max_peer_seq, unread_count, unread_mentions_count, unread_reactions_count, unread_mark, pinned_peer_seq, pinned_canonical_message_id, has_scheduled, available_min_peer_seq, hidden, deleted_at, last_pts, last_pts_at, dialog_schema_version, dialog_payload from user_dialogs where user_id = :user_id and peer_id in (:peerIdList)
func (m *defaultUserDialogsTxModel) SelectByUserPeers(userId int64, peerIdList []int64) (rList []UserDialogs, err error) {
	var (
		query  = fmt.Sprintf("select user_id, peer_type, peer_id, top_peer_seq, top_canonical_message_id, top_message_date, top_message_status, read_inbox_max_peer_seq, read_outbox_max_peer_seq, unread_count, unread_mentions_count, unread_reactions_count, unread_mark, pinned_peer_seq, pinned_canonical_message_id, has_scheduled, available_min_peer_seq, hidden, deleted_at, last_pts, last_pts_at, dialog_schema_version, dialog_payload from user_dialogs where user_id = ? and peer_id in (%s)", sqlx.InInt64List(peerIdList))
		values []UserDialogs
	)
	if len(peerIdList) == 0 {
		rList = []UserDialogs{}
		return
	}

	err = m.tx.QueryRowsPartial(&values, query, userId)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			rList = []UserDialogs{}
			err = nil
			return
		}
		err = fmt.Errorf("user_dialogs.SelectByUserPeers: %w", err)
		return
	}

	rList = values

	return
}

// SelectByUserPeersWithCB
// select user_id, peer_type, peer_id, top_peer_seq, top_canonical_message_id, top_message_date, top_message_status, read_inbox_max_peer_seq, read_outbox_max_peer_seq, unread_count, unread_mentions_count, unread_reactions_count, unread_mark, pinned_peer_seq, pinned_canonical_message_id, has_scheduled, available_min_peer_seq, hidden, deleted_at, last_pts, last_pts_at, dialog_schema_version, dialog_payload from user_dialogs where user_id = :user_id and peer_id in (:peerIdList)
func (m *defaultUserDialogsModel) SelectByUserPeersWithCB(ctx context.Context, userId int64, peerIdList []int64, cb func(sz, i int, v *UserDialogs)) (rList []UserDialogs, err error) {
	var (
		query  = fmt.Sprintf("select user_id, peer_type, peer_id, top_peer_seq, top_canonical_message_id, top_message_date, top_message_status, read_inbox_max_peer_seq, read_outbox_max_peer_seq, unread_count, unread_mentions_count, unread_reactions_count, unread_mark, pinned_peer_seq, pinned_canonical_message_id, has_scheduled, available_min_peer_seq, hidden, deleted_at, last_pts, last_pts_at, dialog_schema_version, dialog_payload from user_dialogs where user_id = ? and peer_id in (%s)", sqlx.InInt64List(peerIdList))
		values []UserDialogs
	)
	if len(peerIdList) == 0 {
		rList = []UserDialogs{}
		return
	}

	err = m.db.QueryRowsPartial(ctx, &values, query, userId)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			rList = []UserDialogs{}
			err = nil
			return
		}
		err = fmt.Errorf("user_dialogs.SelectByUserPeersWithCB: %w", err)
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

// SelectByUserCursor
// select user_id, peer_type, peer_id, top_peer_seq, top_canonical_message_id, top_message_date, top_message_status, read_inbox_max_peer_seq, read_outbox_max_peer_seq, unread_count, unread_mentions_count, unread_reactions_count, unread_mark, pinned_peer_seq, pinned_canonical_message_id, has_scheduled, available_min_peer_seq, hidden, deleted_at, last_pts, last_pts_at, dialog_schema_version, dialog_payload from user_dialogs where user_id = :user_id and hidden = 0 and (:top_message_date = ” or top_message_date < :top_message_date or (top_message_date = :top_message_date and top_peer_seq < :top_peer_seq) or (top_message_date = :top_message_date and top_peer_seq = :top_peer_seq and peer_type > :peer_type) or (top_message_date = :top_message_date and top_peer_seq = :top_peer_seq and peer_type = :peer_type and peer_id > :peer_id)) order by top_message_date desc, top_peer_seq desc, peer_type asc, peer_id asc limit :limit
func (m *defaultUserDialogsModel) SelectByUserCursor(ctx context.Context, userId int64, topMessageDate string, topPeerSeq int64, peerType int32, peerId int64, limit int32) (rList []UserDialogs, err error) {
	var (
		query  = "select user_id, peer_type, peer_id, top_peer_seq, top_canonical_message_id, top_message_date, top_message_status, read_inbox_max_peer_seq, read_outbox_max_peer_seq, unread_count, unread_mentions_count, unread_reactions_count, unread_mark, pinned_peer_seq, pinned_canonical_message_id, has_scheduled, available_min_peer_seq, hidden, deleted_at, last_pts, last_pts_at, dialog_schema_version, dialog_payload from user_dialogs where user_id = ? and hidden = 0 and (? = '' or top_message_date < ? or (top_message_date = ? and top_peer_seq < ?) or (top_message_date = ? and top_peer_seq = ? and peer_type > ?) or (top_message_date = ? and top_peer_seq = ? and peer_type = ? and peer_id > ?)) order by top_message_date desc, top_peer_seq desc, peer_type asc, peer_id asc limit ?"
		values []UserDialogs
	)
	err = m.db.QueryRowsPartial(ctx, &values, query, userId, topMessageDate, topMessageDate, topMessageDate, topPeerSeq, topMessageDate, topPeerSeq, peerType, topMessageDate, topPeerSeq, peerType, peerId, limit)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			rList = []UserDialogs{}
			err = nil
			return
		}
		err = fmt.Errorf("user_dialogs.SelectByUserCursor: %w", err)
		return
	}

	rList = values

	return
}

// SelectByUserCursor
// select user_id, peer_type, peer_id, top_peer_seq, top_canonical_message_id, top_message_date, top_message_status, read_inbox_max_peer_seq, read_outbox_max_peer_seq, unread_count, unread_mentions_count, unread_reactions_count, unread_mark, pinned_peer_seq, pinned_canonical_message_id, has_scheduled, available_min_peer_seq, hidden, deleted_at, last_pts, last_pts_at, dialog_schema_version, dialog_payload from user_dialogs where user_id = :user_id and hidden = 0 and (:top_message_date = ” or top_message_date < :top_message_date or (top_message_date = :top_message_date and top_peer_seq < :top_peer_seq) or (top_message_date = :top_message_date and top_peer_seq = :top_peer_seq and peer_type > :peer_type) or (top_message_date = :top_message_date and top_peer_seq = :top_peer_seq and peer_type = :peer_type and peer_id > :peer_id)) order by top_message_date desc, top_peer_seq desc, peer_type asc, peer_id asc limit :limit
func (m *defaultUserDialogsTxModel) SelectByUserCursor(userId int64, topMessageDate string, topPeerSeq int64, peerType int32, peerId int64, limit int32) (rList []UserDialogs, err error) {
	var (
		query  = "select user_id, peer_type, peer_id, top_peer_seq, top_canonical_message_id, top_message_date, top_message_status, read_inbox_max_peer_seq, read_outbox_max_peer_seq, unread_count, unread_mentions_count, unread_reactions_count, unread_mark, pinned_peer_seq, pinned_canonical_message_id, has_scheduled, available_min_peer_seq, hidden, deleted_at, last_pts, last_pts_at, dialog_schema_version, dialog_payload from user_dialogs where user_id = ? and hidden = 0 and (? = '' or top_message_date < ? or (top_message_date = ? and top_peer_seq < ?) or (top_message_date = ? and top_peer_seq = ? and peer_type > ?) or (top_message_date = ? and top_peer_seq = ? and peer_type = ? and peer_id > ?)) order by top_message_date desc, top_peer_seq desc, peer_type asc, peer_id asc limit ?"
		values []UserDialogs
	)
	err = m.tx.QueryRowsPartial(&values, query, userId, topMessageDate, topMessageDate, topMessageDate, topPeerSeq, topMessageDate, topPeerSeq, peerType, topMessageDate, topPeerSeq, peerType, peerId, limit)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			rList = []UserDialogs{}
			err = nil
			return
		}
		err = fmt.Errorf("user_dialogs.SelectByUserCursor: %w", err)
		return
	}

	rList = values

	return
}

// SelectByUserCursorWithCB
// select user_id, peer_type, peer_id, top_peer_seq, top_canonical_message_id, top_message_date, top_message_status, read_inbox_max_peer_seq, read_outbox_max_peer_seq, unread_count, unread_mentions_count, unread_reactions_count, unread_mark, pinned_peer_seq, pinned_canonical_message_id, has_scheduled, available_min_peer_seq, hidden, deleted_at, last_pts, last_pts_at, dialog_schema_version, dialog_payload from user_dialogs where user_id = :user_id and hidden = 0 and (:top_message_date = ” or top_message_date < :top_message_date or (top_message_date = :top_message_date and top_peer_seq < :top_peer_seq) or (top_message_date = :top_message_date and top_peer_seq = :top_peer_seq and peer_type > :peer_type) or (top_message_date = :top_message_date and top_peer_seq = :top_peer_seq and peer_type = :peer_type and peer_id > :peer_id)) order by top_message_date desc, top_peer_seq desc, peer_type asc, peer_id asc limit :limit
func (m *defaultUserDialogsModel) SelectByUserCursorWithCB(ctx context.Context, userId int64, topMessageDate string, topPeerSeq int64, peerType int32, peerId int64, limit int32, cb func(sz, i int, v *UserDialogs)) (rList []UserDialogs, err error) {
	var (
		query  = "select user_id, peer_type, peer_id, top_peer_seq, top_canonical_message_id, top_message_date, top_message_status, read_inbox_max_peer_seq, read_outbox_max_peer_seq, unread_count, unread_mentions_count, unread_reactions_count, unread_mark, pinned_peer_seq, pinned_canonical_message_id, has_scheduled, available_min_peer_seq, hidden, deleted_at, last_pts, last_pts_at, dialog_schema_version, dialog_payload from user_dialogs where user_id = ? and hidden = 0 and (? = '' or top_message_date < ? or (top_message_date = ? and top_peer_seq < ?) or (top_message_date = ? and top_peer_seq = ? and peer_type > ?) or (top_message_date = ? and top_peer_seq = ? and peer_type = ? and peer_id > ?)) order by top_message_date desc, top_peer_seq desc, peer_type asc, peer_id asc limit ?"
		values []UserDialogs
	)
	err = m.db.QueryRowsPartial(ctx, &values, query, userId, topMessageDate, topMessageDate, topMessageDate, topPeerSeq, topMessageDate, topPeerSeq, peerType, topMessageDate, topPeerSeq, peerType, peerId, limit)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			rList = []UserDialogs{}
			err = nil
			return
		}
		err = fmt.Errorf("user_dialogs.SelectByUserCursorWithCB: %w", err)
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
// update user_dialogs set unread_count = :unread_count, unread_mentions_count = :unread_mentions_count, unread_reactions_count = :unread_reactions_count, unread_mark = :unread_mark, read_inbox_max_peer_seq = :read_inbox_max_peer_seq, read_outbox_max_peer_seq = :read_outbox_max_peer_seq, last_pts = :last_pts, last_pts_at = :last_pts_at where user_id = :user_id and peer_type = :peer_type and peer_id = :peer_id
func (m *defaultUserDialogsModel) UpdateReadState(ctx context.Context, unreadCount int32, unreadMentionsCount int32, unreadReactionsCount int32, unreadMark bool, readInboxMaxPeerSeq int64, readOutboxMaxPeerSeq int64, lastPts int64, lastPtsAt string, userId int64, peerType int32, peerId int64) (rowsAffected int64, err error) {

	var (
		query   = "update user_dialogs set unread_count = ?, unread_mentions_count = ?, unread_reactions_count = ?, unread_mark = ?, read_inbox_max_peer_seq = ?, read_outbox_max_peer_seq = ?, last_pts = ?, last_pts_at = ? where user_id = ? and peer_type = ? and peer_id = ?"
		rResult sql.Result
	)

	rResult, err = m.db.Exec(ctx, query, unreadCount, unreadMentionsCount, unreadReactionsCount, unreadMark, readInboxMaxPeerSeq, readOutboxMaxPeerSeq, lastPts, lastPtsAt, userId, peerType, peerId)

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
// update user_dialogs set unread_count = :unread_count, unread_mentions_count = :unread_mentions_count, unread_reactions_count = :unread_reactions_count, unread_mark = :unread_mark, read_inbox_max_peer_seq = :read_inbox_max_peer_seq, read_outbox_max_peer_seq = :read_outbox_max_peer_seq, last_pts = :last_pts, last_pts_at = :last_pts_at where user_id = :user_id and peer_type = :peer_type and peer_id = :peer_id
func (m *defaultUserDialogsTxModel) UpdateReadState(unreadCount int32, unreadMentionsCount int32, unreadReactionsCount int32, unreadMark bool, readInboxMaxPeerSeq int64, readOutboxMaxPeerSeq int64, lastPts int64, lastPtsAt string, userId int64, peerType int32, peerId int64) (rowsAffected int64, err error) {
	var (
		query   = "update user_dialogs set unread_count = ?, unread_mentions_count = ?, unread_reactions_count = ?, unread_mark = ?, read_inbox_max_peer_seq = ?, read_outbox_max_peer_seq = ?, last_pts = ?, last_pts_at = ? where user_id = ? and peer_type = ? and peer_id = ?"
		rResult sql.Result
	)
	rResult, err = m.tx.Exec(query, unreadCount, unreadMentionsCount, unreadReactionsCount, unreadMark, readInboxMaxPeerSeq, readOutboxMaxPeerSeq, lastPts, lastPtsAt, userId, peerType, peerId)

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

// UpdatePinnedMessage
// update user_dialogs set pinned_peer_seq = :pinned_peer_seq, pinned_canonical_message_id = :pinned_canonical_message_id, last_pts = :last_pts, last_pts_at = :last_pts_at where user_id = :user_id and peer_type = :peer_type and peer_id = :peer_id
func (m *defaultUserDialogsModel) UpdatePinnedMessage(ctx context.Context, pinnedPeerSeq int64, pinnedCanonicalMessageId int64, lastPts int64, lastPtsAt string, userId int64, peerType int32, peerId int64) (rowsAffected int64, err error) {

	var (
		query   = "update user_dialogs set pinned_peer_seq = ?, pinned_canonical_message_id = ?, last_pts = ?, last_pts_at = ? where user_id = ? and peer_type = ? and peer_id = ?"
		rResult sql.Result
	)

	rResult, err = m.db.Exec(ctx, query, pinnedPeerSeq, pinnedCanonicalMessageId, lastPts, lastPtsAt, userId, peerType, peerId)

	if err != nil {
		err = fmt.Errorf("user_dialogs.UpdatePinnedMessage exec: %w", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		err = fmt.Errorf("user_dialogs.UpdatePinnedMessage rows affected: %w", err)
		return
	}

	return
}

// UpdatePinnedMessage
// update user_dialogs set pinned_peer_seq = :pinned_peer_seq, pinned_canonical_message_id = :pinned_canonical_message_id, last_pts = :last_pts, last_pts_at = :last_pts_at where user_id = :user_id and peer_type = :peer_type and peer_id = :peer_id
func (m *defaultUserDialogsTxModel) UpdatePinnedMessage(pinnedPeerSeq int64, pinnedCanonicalMessageId int64, lastPts int64, lastPtsAt string, userId int64, peerType int32, peerId int64) (rowsAffected int64, err error) {
	var (
		query   = "update user_dialogs set pinned_peer_seq = ?, pinned_canonical_message_id = ?, last_pts = ?, last_pts_at = ? where user_id = ? and peer_type = ? and peer_id = ?"
		rResult sql.Result
	)
	rResult, err = m.tx.Exec(query, pinnedPeerSeq, pinnedCanonicalMessageId, lastPts, lastPtsAt, userId, peerType, peerId)

	if err != nil {
		err = fmt.Errorf("user_dialogs.UpdatePinnedMessage exec: %w", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		err = fmt.Errorf("user_dialogs.UpdatePinnedMessage rows affected: %w", err)
		return
	}

	return
}
