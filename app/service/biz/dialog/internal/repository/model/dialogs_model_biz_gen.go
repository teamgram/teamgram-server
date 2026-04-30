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

type bizDialogsModel interface {
	InsertIgnore(ctx context.Context, data *Dialogs) (lastInsertId, rowsAffected int64, err error)
	InsertOrUpdate(ctx context.Context, data *Dialogs) (lastInsertId, rowsAffected int64, err error)
	InsertOrUpdateDialog(ctx context.Context, data *Dialogs) (lastInsertId, rowsAffected int64, err error)
	UpdateOutboxDialog(ctx context.Context, topMessage int32, date2 int64, userId int64, peerType int32, peerId int64) (rowsAffected int64, err error)
	UpdateInboxDialog(ctx context.Context, cMap map[string]interface{}, userId int64, peerType int32, peerId int64) (rowsAffected int64, err error)
	SelectPinnedDialogs(ctx context.Context, userId int64) ([]Dialogs, error)
	SelectPinnedDialogsWithCB(ctx context.Context, userId int64, cb func(sz, i int, v *Dialogs)) ([]Dialogs, error)
	SelectFolderPinnedDialogs(ctx context.Context, userId int64) ([]Dialogs, error)
	SelectFolderPinnedDialogsWithCB(ctx context.Context, userId int64, cb func(sz, i int, v *Dialogs)) ([]Dialogs, error)
	SelectPeerDialogList(ctx context.Context, userId int64, idList []int64) ([]Dialogs, error)
	SelectPeerDialogListWithCB(ctx context.Context, userId int64, idList []int64, cb func(sz, i int, v *Dialogs)) ([]Dialogs, error)
	SelectDialog(ctx context.Context, userId int64, peerType int32, peerId int64) (*Dialogs, error)
	SelectByPeerDialogId(ctx context.Context, userId int64, peerDialogId int64) (*Dialogs, error)
	SelectDialogs(ctx context.Context, userId int64, folderId int32) ([]Dialogs, error)
	SelectDialogsWithCB(ctx context.Context, userId int64, folderId int32, cb func(sz, i int, v *Dialogs)) ([]Dialogs, error)
	SelectExcludePinnedDialogs(ctx context.Context, userId int64) ([]Dialogs, error)
	SelectExcludePinnedDialogsWithCB(ctx context.Context, userId int64, cb func(sz, i int, v *Dialogs)) ([]Dialogs, error)
	SelectExcludeFolderPinnedDialogs(ctx context.Context, userId int64) ([]Dialogs, error)
	SelectExcludeFolderPinnedDialogsWithCB(ctx context.Context, userId int64, cb func(sz, i int, v *Dialogs)) ([]Dialogs, error)
	UpdateReadInboxMaxId(ctx context.Context, unreadCount int32, readInboxMaxId int32, userId int64, peerDialogId int64) (rowsAffected int64, err error)
	UpdateReadOutboxMaxId(ctx context.Context, readOutboxMaxId int32, userId int64, peerDialogId int64) (rowsAffected int64, err error)
	UpdateTopMessage(ctx context.Context, topMessage int32, userId int64, peerDialogId int64) (rowsAffected int64, err error)
	UpdatePinnedMsgId(ctx context.Context, pinnedMsgId int32, userId int64, peerDialogId int64) (rowsAffected int64, err error)
	Delete(ctx context.Context, userId int64, peerType int32, peerId int64) (rowsAffected int64, err error)
	SelectDialogsByGTReadInboxMaxId(ctx context.Context, peerType int32, peerId int64, readInboxMaxId int32, userId int64) ([]int64, error)
	SelectDialogsByGTReadInboxMaxIdWithCB(ctx context.Context, peerType int32, peerId int64, readInboxMaxId int32, userId int64, cb func(sz, i int, v int64)) ([]int64, error)
	UpdateCustomMap(ctx context.Context, cMap map[string]interface{}, userId int64, peerType int32, peerId int64) (rowsAffected int64, err error)
	SaveDraft(ctx context.Context, draftType int32, draftMessageData string, userId int64, peerType int32, peerId int64) (rowsAffected int64, err error)
	SelectAllDrafts(ctx context.Context, userId int64) ([]Dialogs, error)
	SelectAllDraftsWithCB(ctx context.Context, userId int64, cb func(sz, i int, v *Dialogs)) ([]Dialogs, error)
	ClearAllDrafts(ctx context.Context, userId int64) (rowsAffected int64, err error)
	UpdatePeerFolderId(ctx context.Context, folderId int32, userId int64, peerType int32, peerId int64) (rowsAffected int64, err error)
	UpdatePeerDialogListFolderId(ctx context.Context, folderId int32, userId int64, idList []int64) (rowsAffected int64, err error)
	UpdatePeerDialogListPinned(ctx context.Context, pinned int64, userId int64, idList []int64) (rowsAffected int64, err error)
	UpdateFolderPeerDialogListPinned(ctx context.Context, folderPinned int64, userId int64, idList []int64) (rowsAffected int64, err error)
	UpdateUnPinnedNotIdList(ctx context.Context, userId int64, idList []int64) (rowsAffected int64, err error)
	UpdateFolderUnPinnedNotIdList(ctx context.Context, userId int64, idList []int64) (rowsAffected int64, err error)
	SelectAllDialogs(ctx context.Context, userId int64) ([]Dialogs, error)
	SelectAllDialogsWithCB(ctx context.Context, userId int64, cb func(sz, i int, v *Dialogs)) ([]Dialogs, error)
	SelectDialogsByPeerType(ctx context.Context, userId int64, peerTypeList []int32) ([]Dialogs, error)
	SelectDialogsByPeerTypeWithCB(ctx context.Context, userId int64, peerTypeList []int32, cb func(sz, i int, v *Dialogs)) ([]Dialogs, error)
	UpdateUnreadCount(ctx context.Context, unreadCount int32, unreadMentionsCount int32, unreadReactionsCount int32, userId int64, peerType int32, peerId int64) (rowsAffected int64, err error)
}

type DialogsTxModel interface {
	InsertIgnore(data *Dialogs) (lastInsertId, rowsAffected int64, err error)
	InsertOrUpdate(data *Dialogs) (lastInsertId, rowsAffected int64, err error)
	InsertOrUpdateDialog(data *Dialogs) (lastInsertId, rowsAffected int64, err error)
	UpdateOutboxDialog(topMessage int32, date2 int64, userId int64, peerType int32, peerId int64) (rowsAffected int64, err error)
	UpdateInboxDialog(cMap map[string]interface{}, userId int64, peerType int32, peerId int64) (rowsAffected int64, err error)
	SelectPinnedDialogs(userId int64) ([]Dialogs, error)
	SelectFolderPinnedDialogs(userId int64) ([]Dialogs, error)
	SelectPeerDialogList(userId int64, idList []int64) ([]Dialogs, error)
	SelectDialog(userId int64, peerType int32, peerId int64) (*Dialogs, error)
	SelectByPeerDialogId(userId int64, peerDialogId int64) (*Dialogs, error)
	SelectDialogs(userId int64, folderId int32) ([]Dialogs, error)
	SelectExcludePinnedDialogs(userId int64) ([]Dialogs, error)
	SelectExcludeFolderPinnedDialogs(userId int64) ([]Dialogs, error)
	UpdateReadInboxMaxId(unreadCount int32, readInboxMaxId int32, userId int64, peerDialogId int64) (rowsAffected int64, err error)
	UpdateReadOutboxMaxId(readOutboxMaxId int32, userId int64, peerDialogId int64) (rowsAffected int64, err error)
	UpdateTopMessage(topMessage int32, userId int64, peerDialogId int64) (rowsAffected int64, err error)
	UpdatePinnedMsgId(pinnedMsgId int32, userId int64, peerDialogId int64) (rowsAffected int64, err error)
	Delete(userId int64, peerType int32, peerId int64) (rowsAffected int64, err error)
	SelectDialogsByGTReadInboxMaxId(peerType int32, peerId int64, readInboxMaxId int32, userId int64) ([]int64, error)
	UpdateCustomMap(cMap map[string]interface{}, userId int64, peerType int32, peerId int64) (rowsAffected int64, err error)
	SaveDraft(draftType int32, draftMessageData string, userId int64, peerType int32, peerId int64) (rowsAffected int64, err error)
	SelectAllDrafts(userId int64) ([]Dialogs, error)
	ClearAllDrafts(userId int64) (rowsAffected int64, err error)
	UpdatePeerFolderId(folderId int32, userId int64, peerType int32, peerId int64) (rowsAffected int64, err error)
	UpdatePeerDialogListFolderId(folderId int32, userId int64, idList []int64) (rowsAffected int64, err error)
	UpdatePeerDialogListPinned(pinned int64, userId int64, idList []int64) (rowsAffected int64, err error)
	UpdateFolderPeerDialogListPinned(folderPinned int64, userId int64, idList []int64) (rowsAffected int64, err error)
	UpdateUnPinnedNotIdList(userId int64, idList []int64) (rowsAffected int64, err error)
	UpdateFolderUnPinnedNotIdList(userId int64, idList []int64) (rowsAffected int64, err error)
	SelectAllDialogs(userId int64) ([]Dialogs, error)
	SelectDialogsByPeerType(userId int64, peerTypeList []int32) ([]Dialogs, error)
	UpdateUnreadCount(unreadCount int32, unreadMentionsCount int32, unreadReactionsCount int32, userId int64, peerType int32, peerId int64) (rowsAffected int64, err error)
}

type defaultDialogsTxModel struct {
	tx *sqlx.Tx
}

func NewDialogsTxModel(tx *sqlx.Tx) DialogsTxModel {
	return &defaultDialogsTxModel{tx: tx}
}

// InsertIgnore
// insert ignore into dialogs(user_id, peer_type, peer_id, peer_dialog_id, top_message, pinned_msg_id, read_inbox_max_id, read_outbox_max_id, unread_count, unread_mentions_count, unread_mark, draft_message_data, date2) values (:user_id, :peer_type, :peer_id, :peer_dialog_id, :top_message, :pinned_msg_id, :read_inbox_max_id, :read_outbox_max_id, :unread_count, :unread_mentions_count, :unread_mark, :draft_message_data, :date2)
func (m *defaultDialogsModel) InsertIgnore(ctx context.Context, data *Dialogs) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert ignore into dialogs(user_id, peer_type, peer_id, peer_dialog_id, top_message, pinned_msg_id, read_inbox_max_id, read_outbox_max_id, unread_count, unread_mentions_count, unread_mark, draft_message_data, date2) values (:user_id, :peer_type, :peer_id, :peer_dialog_id, :top_message, :pinned_msg_id, :read_inbox_max_id, :read_outbox_max_id, :unread_count, :unread_mentions_count, :unread_mark, :draft_message_data, :date2)"
		r     sql.Result
	)

	r, err = m.db.NamedExec(ctx, query, data)
	if err != nil {
		err = fmt.Errorf("dialogs.InsertIgnore named exec: %w", err)
		return
	}

	lastInsertId, err = r.LastInsertId()
	if err != nil {
		err = fmt.Errorf("dialogs.InsertIgnore last insert id: %w", err)
		return
	}
	rowsAffected, err = r.RowsAffected()
	if err != nil {
		err = fmt.Errorf("dialogs.InsertIgnore rows affected: %w", err)
	}

	return

}

// InsertIgnore
// insert ignore into dialogs(user_id, peer_type, peer_id, peer_dialog_id, top_message, pinned_msg_id, read_inbox_max_id, read_outbox_max_id, unread_count, unread_mentions_count, unread_mark, draft_message_data, date2) values (:user_id, :peer_type, :peer_id, :peer_dialog_id, :top_message, :pinned_msg_id, :read_inbox_max_id, :read_outbox_max_id, :unread_count, :unread_mentions_count, :unread_mark, :draft_message_data, :date2)
func (m *defaultDialogsTxModel) InsertIgnore(data *Dialogs) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert ignore into dialogs(user_id, peer_type, peer_id, peer_dialog_id, top_message, pinned_msg_id, read_inbox_max_id, read_outbox_max_id, unread_count, unread_mentions_count, unread_mark, draft_message_data, date2) values (:user_id, :peer_type, :peer_id, :peer_dialog_id, :top_message, :pinned_msg_id, :read_inbox_max_id, :read_outbox_max_id, :unread_count, :unread_mentions_count, :unread_mark, :draft_message_data, :date2)"
		r     sql.Result
	)

	r, err = m.tx.NamedExec(query, data)
	if err != nil {
		err = fmt.Errorf("dialogs.InsertIgnore named exec: %w", err)
		return
	}

	lastInsertId, err = r.LastInsertId()
	if err != nil {
		err = fmt.Errorf("dialogs.InsertIgnore last insert id: %w", err)
		return
	}
	rowsAffected, err = r.RowsAffected()
	if err != nil {
		err = fmt.Errorf("dialogs.InsertIgnore rows affected: %w", err)
	}

	return
}

// InsertOrUpdate
// insert into dialogs(user_id, peer_type, peer_id, peer_dialog_id, top_message, pinned_msg_id, unread_count, unread_mentions_count, draft_message_data, date2) values (:user_id, :peer_type, :peer_id, :peer_dialog_id, :top_message, :pinned_msg_id, :unread_count, :unread_mentions_count, :draft_message_data, :date2) on duplicate key update top_message = values(top_message), unread_count = unread_count + values(unread_count), unread_mentions_count = unread_mentions_count + values(unread_mentions_count), date2 = values(date2)
func (m *defaultDialogsModel) InsertOrUpdate(ctx context.Context, data *Dialogs) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into dialogs(user_id, peer_type, peer_id, peer_dialog_id, top_message, pinned_msg_id, unread_count, unread_mentions_count, draft_message_data, date2) values (:user_id, :peer_type, :peer_id, :peer_dialog_id, :top_message, :pinned_msg_id, :unread_count, :unread_mentions_count, :draft_message_data, :date2) on duplicate key update top_message = values(top_message), unread_count = unread_count + values(unread_count), unread_mentions_count = unread_mentions_count + values(unread_mentions_count), date2 = values(date2)"
		r     sql.Result
	)

	r, err = m.db.NamedExec(ctx, query, data)
	if err != nil {
		err = fmt.Errorf("dialogs.InsertOrUpdate named exec: %w", err)
		return
	}

	lastInsertId, err = r.LastInsertId()
	if err != nil {
		err = fmt.Errorf("dialogs.InsertOrUpdate last insert id: %w", err)
		return
	}
	rowsAffected, err = r.RowsAffected()
	if err != nil {
		err = fmt.Errorf("dialogs.InsertOrUpdate rows affected: %w", err)
	}

	return

}

// InsertOrUpdate
// insert into dialogs(user_id, peer_type, peer_id, peer_dialog_id, top_message, pinned_msg_id, unread_count, unread_mentions_count, draft_message_data, date2) values (:user_id, :peer_type, :peer_id, :peer_dialog_id, :top_message, :pinned_msg_id, :unread_count, :unread_mentions_count, :draft_message_data, :date2) on duplicate key update top_message = values(top_message), unread_count = unread_count + values(unread_count), unread_mentions_count = unread_mentions_count + values(unread_mentions_count), date2 = values(date2)
func (m *defaultDialogsTxModel) InsertOrUpdate(data *Dialogs) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into dialogs(user_id, peer_type, peer_id, peer_dialog_id, top_message, pinned_msg_id, unread_count, unread_mentions_count, draft_message_data, date2) values (:user_id, :peer_type, :peer_id, :peer_dialog_id, :top_message, :pinned_msg_id, :unread_count, :unread_mentions_count, :draft_message_data, :date2) on duplicate key update top_message = values(top_message), unread_count = unread_count + values(unread_count), unread_mentions_count = unread_mentions_count + values(unread_mentions_count), date2 = values(date2)"
		r     sql.Result
	)

	r, err = m.tx.NamedExec(query, data)
	if err != nil {
		err = fmt.Errorf("dialogs.InsertOrUpdate named exec: %w", err)
		return
	}

	lastInsertId, err = r.LastInsertId()
	if err != nil {
		err = fmt.Errorf("dialogs.InsertOrUpdate last insert id: %w", err)
		return
	}
	rowsAffected, err = r.RowsAffected()
	if err != nil {
		err = fmt.Errorf("dialogs.InsertOrUpdate rows affected: %w", err)
	}

	return
}

// InsertOrUpdateDialog
// insert into dialogs(user_id, peer_type, peer_id, peer_dialog_id, top_message, pinned_msg_id, peer_dialog_id, read_inbox_max_id, draft_message_data, date2) values (:user_id, :peer_type, :peer_id, :peer_dialog_id, :top_message, :pinned_msg_id, :peer_dialog_id, :read_inbox_max_id, :draft_message_data, :date2) on duplicate key update top_message = values(top_message), read_inbox_max_id = values(read_inbox_max_id), draft_message_data = values(draft_message_data), date2 = values(date2), deleted = 0
func (m *defaultDialogsModel) InsertOrUpdateDialog(ctx context.Context, data *Dialogs) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into dialogs(user_id, peer_type, peer_id, peer_dialog_id, top_message, pinned_msg_id, peer_dialog_id, read_inbox_max_id, draft_message_data, date2) values (:user_id, :peer_type, :peer_id, :peer_dialog_id, :top_message, :pinned_msg_id, :peer_dialog_id, :read_inbox_max_id, :draft_message_data, :date2) on duplicate key update top_message = values(top_message), read_inbox_max_id = values(read_inbox_max_id), draft_message_data = values(draft_message_data), date2 = values(date2), deleted = 0"
		r     sql.Result
	)

	r, err = m.db.NamedExec(ctx, query, data)
	if err != nil {
		err = fmt.Errorf("dialogs.InsertOrUpdateDialog named exec: %w", err)
		return
	}

	lastInsertId, err = r.LastInsertId()
	if err != nil {
		err = fmt.Errorf("dialogs.InsertOrUpdateDialog last insert id: %w", err)
		return
	}
	rowsAffected, err = r.RowsAffected()
	if err != nil {
		err = fmt.Errorf("dialogs.InsertOrUpdateDialog rows affected: %w", err)
	}

	return

}

// InsertOrUpdateDialog
// insert into dialogs(user_id, peer_type, peer_id, peer_dialog_id, top_message, pinned_msg_id, peer_dialog_id, read_inbox_max_id, draft_message_data, date2) values (:user_id, :peer_type, :peer_id, :peer_dialog_id, :top_message, :pinned_msg_id, :peer_dialog_id, :read_inbox_max_id, :draft_message_data, :date2) on duplicate key update top_message = values(top_message), read_inbox_max_id = values(read_inbox_max_id), draft_message_data = values(draft_message_data), date2 = values(date2), deleted = 0
func (m *defaultDialogsTxModel) InsertOrUpdateDialog(data *Dialogs) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into dialogs(user_id, peer_type, peer_id, peer_dialog_id, top_message, pinned_msg_id, peer_dialog_id, read_inbox_max_id, draft_message_data, date2) values (:user_id, :peer_type, :peer_id, :peer_dialog_id, :top_message, :pinned_msg_id, :peer_dialog_id, :read_inbox_max_id, :draft_message_data, :date2) on duplicate key update top_message = values(top_message), read_inbox_max_id = values(read_inbox_max_id), draft_message_data = values(draft_message_data), date2 = values(date2), deleted = 0"
		r     sql.Result
	)

	r, err = m.tx.NamedExec(query, data)
	if err != nil {
		err = fmt.Errorf("dialogs.InsertOrUpdateDialog named exec: %w", err)
		return
	}

	lastInsertId, err = r.LastInsertId()
	if err != nil {
		err = fmt.Errorf("dialogs.InsertOrUpdateDialog last insert id: %w", err)
		return
	}
	rowsAffected, err = r.RowsAffected()
	if err != nil {
		err = fmt.Errorf("dialogs.InsertOrUpdateDialog rows affected: %w", err)
	}

	return
}

// UpdateOutboxDialog
// update dialogs set unread_count = 0, deleted = 0, top_message = :top_message, date2 = :date2, unread_mark = 0, draft_message_data = 'null' where user_id = :user_id and peer_type = :peer_type and peer_id = :peer_id
func (m *defaultDialogsModel) UpdateOutboxDialog(ctx context.Context, topMessage int32, date2 int64, userId int64, peerType int32, peerId int64) (rowsAffected int64, err error) {

	var (
		query   = "update dialogs set unread_count = 0, deleted = 0, top_message = ?, date2 = ?, unread_mark = 0, draft_message_data = 'null' where user_id = ? and peer_type = ? and peer_id = ?"
		rResult sql.Result
	)

	rResult, err = m.db.Exec(ctx, query, topMessage, date2, userId, peerType, peerId)

	if err != nil {
		err = fmt.Errorf("dialogs.UpdateOutboxDialog exec: %w", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		err = fmt.Errorf("dialogs.UpdateOutboxDialog rows affected: %w", err)
		return
	}

	return
}

// UpdateOutboxDialog
// update dialogs set unread_count = 0, deleted = 0, top_message = :top_message, date2 = :date2, unread_mark = 0, draft_message_data = 'null' where user_id = :user_id and peer_type = :peer_type and peer_id = :peer_id
func (m *defaultDialogsTxModel) UpdateOutboxDialog(topMessage int32, date2 int64, userId int64, peerType int32, peerId int64) (rowsAffected int64, err error) {
	var (
		query   = "update dialogs set unread_count = 0, deleted = 0, top_message = ?, date2 = ?, unread_mark = 0, draft_message_data = 'null' where user_id = ? and peer_type = ? and peer_id = ?"
		rResult sql.Result
	)
	rResult, err = m.tx.Exec(query, topMessage, date2, userId, peerType, peerId)

	if err != nil {
		err = fmt.Errorf("dialogs.UpdateOutboxDialog exec: %w", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		err = fmt.Errorf("dialogs.UpdateOutboxDialog rows affected: %w", err)
		return
	}

	return
}

// UpdateInboxDialog
// update dialogs set unread_count = unread_count + 1, deleted = 0, %s where user_id = :user_id and peer_type = :peer_type and peer_id = :peer_id
func (m *defaultDialogsModel) UpdateInboxDialog(ctx context.Context, cMap map[string]interface{}, userId int64, peerType int32, peerId int64) (rowsAffected int64, err error) {

	names := make([]string, 0, len(cMap))
	aValues := make([]interface{}, 0, len(cMap))
	for k, v := range cMap {
		names = append(names, k+" = ?")
		aValues = append(aValues, v)
	}

	var (
		query   = fmt.Sprintf("update dialogs set unread_count = unread_count + 1, deleted = 0, %s where user_id = ? and peer_type = ? and peer_id = ?", strings.Join(names, ", "))
		rResult sql.Result
	)

	aValues = append(aValues, userId)
	aValues = append(aValues, peerType)
	aValues = append(aValues, peerId)

	rResult, err = m.db.Exec(ctx, query, aValues...)

	if err != nil {
		err = fmt.Errorf("dialogs.UpdateInboxDialog exec: %w", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		err = fmt.Errorf("dialogs.UpdateInboxDialog rows affected: %w", err)
		return
	}

	return
}

// UpdateInboxDialog
// update dialogs set unread_count = unread_count + 1, deleted = 0, %s where user_id = :user_id and peer_type = :peer_type and peer_id = :peer_id
func (m *defaultDialogsTxModel) UpdateInboxDialog(cMap map[string]interface{}, userId int64, peerType int32, peerId int64) (rowsAffected int64, err error) {
	names := make([]string, 0, len(cMap))
	aValues := make([]interface{}, 0, len(cMap))
	for k, v := range cMap {
		names = append(names, k+" = ?")
		aValues = append(aValues, v)
	}

	var (
		query   = fmt.Sprintf("update dialogs set unread_count = unread_count + 1, deleted = 0, %s where user_id = ? and peer_type = ? and peer_id = ?", strings.Join(names, ", "))
		rResult sql.Result
	)

	aValues = append(aValues, userId)
	aValues = append(aValues, peerType)
	aValues = append(aValues, peerId)

	rResult, err = m.tx.Exec(query, aValues...)

	if err != nil {
		err = fmt.Errorf("dialogs.UpdateInboxDialog exec: %w", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		err = fmt.Errorf("dialogs.UpdateInboxDialog rows affected: %w", err)
		return
	}

	return
}

// SelectPinnedDialogs
// select id, user_id, peer_type, peer_id, peer_dialog_id, pinned, top_message, pinned_msg_id, read_inbox_max_id, read_outbox_max_id, unread_count, unread_mentions_count, unread_reactions_count, unread_mark, draft_type, draft_message_data, folder_id, folder_pinned, has_scheduled, ttl_period, theme_emoticon, wallpaper_id, wallpaper_overridden, date2 from dialogs where user_id = :user_id and folder_id = 0 and pinned > 0 and deleted = 0
func (m *defaultDialogsModel) SelectPinnedDialogs(ctx context.Context, userId int64) (rList []Dialogs, err error) {
	var (
		query  = "select id, user_id, peer_type, peer_id, peer_dialog_id, pinned, top_message, pinned_msg_id, read_inbox_max_id, read_outbox_max_id, unread_count, unread_mentions_count, unread_reactions_count, unread_mark, draft_type, draft_message_data, folder_id, folder_pinned, has_scheduled, ttl_period, theme_emoticon, wallpaper_id, wallpaper_overridden, date2 from dialogs where user_id = ? and folder_id = 0 and pinned > 0 and deleted = 0"
		values []Dialogs
	)
	err = m.db.QueryRowsPartial(ctx, &values, query, userId)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			rList = []Dialogs{}
			err = nil
			return
		}
		err = fmt.Errorf("dialogs.SelectPinnedDialogs: %w", err)
		return
	}

	rList = values

	return
}

// SelectPinnedDialogs
// select id, user_id, peer_type, peer_id, peer_dialog_id, pinned, top_message, pinned_msg_id, read_inbox_max_id, read_outbox_max_id, unread_count, unread_mentions_count, unread_reactions_count, unread_mark, draft_type, draft_message_data, folder_id, folder_pinned, has_scheduled, ttl_period, theme_emoticon, wallpaper_id, wallpaper_overridden, date2 from dialogs where user_id = :user_id and folder_id = 0 and pinned > 0 and deleted = 0
func (m *defaultDialogsTxModel) SelectPinnedDialogs(userId int64) (rList []Dialogs, err error) {
	var (
		query  = "select id, user_id, peer_type, peer_id, peer_dialog_id, pinned, top_message, pinned_msg_id, read_inbox_max_id, read_outbox_max_id, unread_count, unread_mentions_count, unread_reactions_count, unread_mark, draft_type, draft_message_data, folder_id, folder_pinned, has_scheduled, ttl_period, theme_emoticon, wallpaper_id, wallpaper_overridden, date2 from dialogs where user_id = ? and folder_id = 0 and pinned > 0 and deleted = 0"
		values []Dialogs
	)
	err = m.tx.QueryRowsPartial(&values, query, userId)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			rList = []Dialogs{}
			err = nil
			return
		}
		err = fmt.Errorf("dialogs.SelectPinnedDialogs: %w", err)
		return
	}

	rList = values

	return
}

// SelectPinnedDialogsWithCB
// select id, user_id, peer_type, peer_id, peer_dialog_id, pinned, top_message, pinned_msg_id, read_inbox_max_id, read_outbox_max_id, unread_count, unread_mentions_count, unread_reactions_count, unread_mark, draft_type, draft_message_data, folder_id, folder_pinned, has_scheduled, ttl_period, theme_emoticon, wallpaper_id, wallpaper_overridden, date2 from dialogs where user_id = :user_id and folder_id = 0 and pinned > 0 and deleted = 0
func (m *defaultDialogsModel) SelectPinnedDialogsWithCB(ctx context.Context, userId int64, cb func(sz, i int, v *Dialogs)) (rList []Dialogs, err error) {
	var (
		query  = "select id, user_id, peer_type, peer_id, peer_dialog_id, pinned, top_message, pinned_msg_id, read_inbox_max_id, read_outbox_max_id, unread_count, unread_mentions_count, unread_reactions_count, unread_mark, draft_type, draft_message_data, folder_id, folder_pinned, has_scheduled, ttl_period, theme_emoticon, wallpaper_id, wallpaper_overridden, date2 from dialogs where user_id = ? and folder_id = 0 and pinned > 0 and deleted = 0"
		values []Dialogs
	)
	err = m.db.QueryRowsPartial(ctx, &values, query, userId)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			rList = []Dialogs{}
			err = nil
			return
		}
		err = fmt.Errorf("dialogs.SelectPinnedDialogsWithCB: %w", err)
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

// SelectFolderPinnedDialogs
// select id, user_id, peer_type, peer_id, peer_dialog_id, pinned, top_message, pinned_msg_id, read_inbox_max_id, read_outbox_max_id, unread_count, unread_mentions_count, unread_reactions_count, unread_mark, draft_type, draft_message_data, folder_id, folder_pinned, has_scheduled, ttl_period, theme_emoticon, wallpaper_id, wallpaper_overridden, date2 from dialogs where user_id = :user_id and folder_id = 1 and folder_pinned > 0 and deleted = 0
func (m *defaultDialogsModel) SelectFolderPinnedDialogs(ctx context.Context, userId int64) (rList []Dialogs, err error) {
	var (
		query  = "select id, user_id, peer_type, peer_id, peer_dialog_id, pinned, top_message, pinned_msg_id, read_inbox_max_id, read_outbox_max_id, unread_count, unread_mentions_count, unread_reactions_count, unread_mark, draft_type, draft_message_data, folder_id, folder_pinned, has_scheduled, ttl_period, theme_emoticon, wallpaper_id, wallpaper_overridden, date2 from dialogs where user_id = ? and folder_id = 1 and folder_pinned > 0 and deleted = 0"
		values []Dialogs
	)
	err = m.db.QueryRowsPartial(ctx, &values, query, userId)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			rList = []Dialogs{}
			err = nil
			return
		}
		err = fmt.Errorf("dialogs.SelectFolderPinnedDialogs: %w", err)
		return
	}

	rList = values

	return
}

// SelectFolderPinnedDialogs
// select id, user_id, peer_type, peer_id, peer_dialog_id, pinned, top_message, pinned_msg_id, read_inbox_max_id, read_outbox_max_id, unread_count, unread_mentions_count, unread_reactions_count, unread_mark, draft_type, draft_message_data, folder_id, folder_pinned, has_scheduled, ttl_period, theme_emoticon, wallpaper_id, wallpaper_overridden, date2 from dialogs where user_id = :user_id and folder_id = 1 and folder_pinned > 0 and deleted = 0
func (m *defaultDialogsTxModel) SelectFolderPinnedDialogs(userId int64) (rList []Dialogs, err error) {
	var (
		query  = "select id, user_id, peer_type, peer_id, peer_dialog_id, pinned, top_message, pinned_msg_id, read_inbox_max_id, read_outbox_max_id, unread_count, unread_mentions_count, unread_reactions_count, unread_mark, draft_type, draft_message_data, folder_id, folder_pinned, has_scheduled, ttl_period, theme_emoticon, wallpaper_id, wallpaper_overridden, date2 from dialogs where user_id = ? and folder_id = 1 and folder_pinned > 0 and deleted = 0"
		values []Dialogs
	)
	err = m.tx.QueryRowsPartial(&values, query, userId)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			rList = []Dialogs{}
			err = nil
			return
		}
		err = fmt.Errorf("dialogs.SelectFolderPinnedDialogs: %w", err)
		return
	}

	rList = values

	return
}

// SelectFolderPinnedDialogsWithCB
// select id, user_id, peer_type, peer_id, peer_dialog_id, pinned, top_message, pinned_msg_id, read_inbox_max_id, read_outbox_max_id, unread_count, unread_mentions_count, unread_reactions_count, unread_mark, draft_type, draft_message_data, folder_id, folder_pinned, has_scheduled, ttl_period, theme_emoticon, wallpaper_id, wallpaper_overridden, date2 from dialogs where user_id = :user_id and folder_id = 1 and folder_pinned > 0 and deleted = 0
func (m *defaultDialogsModel) SelectFolderPinnedDialogsWithCB(ctx context.Context, userId int64, cb func(sz, i int, v *Dialogs)) (rList []Dialogs, err error) {
	var (
		query  = "select id, user_id, peer_type, peer_id, peer_dialog_id, pinned, top_message, pinned_msg_id, read_inbox_max_id, read_outbox_max_id, unread_count, unread_mentions_count, unread_reactions_count, unread_mark, draft_type, draft_message_data, folder_id, folder_pinned, has_scheduled, ttl_period, theme_emoticon, wallpaper_id, wallpaper_overridden, date2 from dialogs where user_id = ? and folder_id = 1 and folder_pinned > 0 and deleted = 0"
		values []Dialogs
	)
	err = m.db.QueryRowsPartial(ctx, &values, query, userId)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			rList = []Dialogs{}
			err = nil
			return
		}
		err = fmt.Errorf("dialogs.SelectFolderPinnedDialogsWithCB: %w", err)
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

// SelectPeerDialogList
// select id, user_id, peer_type, peer_id, peer_dialog_id, pinned, top_message, pinned_msg_id, read_inbox_max_id, read_outbox_max_id, unread_count, unread_mentions_count, unread_reactions_count, unread_mark, draft_type, draft_message_data, folder_id, folder_pinned, has_scheduled, ttl_period, theme_emoticon, wallpaper_id, wallpaper_overridden, date2 from dialogs where user_id = :user_id and peer_dialog_id in (:idList) and deleted = 0
func (m *defaultDialogsModel) SelectPeerDialogList(ctx context.Context, userId int64, idList []int64) (rList []Dialogs, err error) {
	var (
		query  = fmt.Sprintf("select id, user_id, peer_type, peer_id, peer_dialog_id, pinned, top_message, pinned_msg_id, read_inbox_max_id, read_outbox_max_id, unread_count, unread_mentions_count, unread_reactions_count, unread_mark, draft_type, draft_message_data, folder_id, folder_pinned, has_scheduled, ttl_period, theme_emoticon, wallpaper_id, wallpaper_overridden, date2 from dialogs where user_id = ? and peer_dialog_id in (%s) and deleted = 0", sqlx.InInt64List(idList))
		values []Dialogs
	)
	if len(idList) == 0 {
		rList = []Dialogs{}
		return
	}

	err = m.db.QueryRowsPartial(ctx, &values, query, userId)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			rList = []Dialogs{}
			err = nil
			return
		}
		err = fmt.Errorf("dialogs.SelectPeerDialogList: %w", err)
		return
	}

	rList = values

	return
}

// SelectPeerDialogList
// select id, user_id, peer_type, peer_id, peer_dialog_id, pinned, top_message, pinned_msg_id, read_inbox_max_id, read_outbox_max_id, unread_count, unread_mentions_count, unread_reactions_count, unread_mark, draft_type, draft_message_data, folder_id, folder_pinned, has_scheduled, ttl_period, theme_emoticon, wallpaper_id, wallpaper_overridden, date2 from dialogs where user_id = :user_id and peer_dialog_id in (:idList) and deleted = 0
func (m *defaultDialogsTxModel) SelectPeerDialogList(userId int64, idList []int64) (rList []Dialogs, err error) {
	var (
		query  = fmt.Sprintf("select id, user_id, peer_type, peer_id, peer_dialog_id, pinned, top_message, pinned_msg_id, read_inbox_max_id, read_outbox_max_id, unread_count, unread_mentions_count, unread_reactions_count, unread_mark, draft_type, draft_message_data, folder_id, folder_pinned, has_scheduled, ttl_period, theme_emoticon, wallpaper_id, wallpaper_overridden, date2 from dialogs where user_id = ? and peer_dialog_id in (%s) and deleted = 0", sqlx.InInt64List(idList))
		values []Dialogs
	)
	if len(idList) == 0 {
		rList = []Dialogs{}
		return
	}

	err = m.tx.QueryRowsPartial(&values, query, userId)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			rList = []Dialogs{}
			err = nil
			return
		}
		err = fmt.Errorf("dialogs.SelectPeerDialogList: %w", err)
		return
	}

	rList = values

	return
}

// SelectPeerDialogListWithCB
// select id, user_id, peer_type, peer_id, peer_dialog_id, pinned, top_message, pinned_msg_id, read_inbox_max_id, read_outbox_max_id, unread_count, unread_mentions_count, unread_reactions_count, unread_mark, draft_type, draft_message_data, folder_id, folder_pinned, has_scheduled, ttl_period, theme_emoticon, wallpaper_id, wallpaper_overridden, date2 from dialogs where user_id = :user_id and peer_dialog_id in (:idList) and deleted = 0
func (m *defaultDialogsModel) SelectPeerDialogListWithCB(ctx context.Context, userId int64, idList []int64, cb func(sz, i int, v *Dialogs)) (rList []Dialogs, err error) {
	var (
		query  = fmt.Sprintf("select id, user_id, peer_type, peer_id, peer_dialog_id, pinned, top_message, pinned_msg_id, read_inbox_max_id, read_outbox_max_id, unread_count, unread_mentions_count, unread_reactions_count, unread_mark, draft_type, draft_message_data, folder_id, folder_pinned, has_scheduled, ttl_period, theme_emoticon, wallpaper_id, wallpaper_overridden, date2 from dialogs where user_id = ? and peer_dialog_id in (%s) and deleted = 0", sqlx.InInt64List(idList))
		values []Dialogs
	)
	if len(idList) == 0 {
		rList = []Dialogs{}
		return
	}

	err = m.db.QueryRowsPartial(ctx, &values, query, userId)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			rList = []Dialogs{}
			err = nil
			return
		}
		err = fmt.Errorf("dialogs.SelectPeerDialogListWithCB: %w", err)
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

// SelectDialog
// select id, user_id, peer_type, peer_id, peer_dialog_id, pinned, top_message, pinned_msg_id, read_inbox_max_id, read_outbox_max_id, unread_count, unread_mentions_count, unread_reactions_count, unread_mark, draft_type, draft_message_data, folder_id, folder_pinned, has_scheduled, ttl_period, theme_emoticon, wallpaper_id, wallpaper_overridden, date2 from dialogs where user_id = :user_id and peer_type = :peer_type and peer_id = :peer_id and deleted = 0
func (m *defaultDialogsModel) SelectDialog(ctx context.Context, userId int64, peerType int32, peerId int64) (rValue *Dialogs, err error) {

	var (
		query = "select id, user_id, peer_type, peer_id, peer_dialog_id, pinned, top_message, pinned_msg_id, read_inbox_max_id, read_outbox_max_id, unread_count, unread_mentions_count, unread_reactions_count, unread_mark, draft_type, draft_message_data, folder_id, folder_pinned, has_scheduled, ttl_period, theme_emoticon, wallpaper_id, wallpaper_overridden, date2 from dialogs where user_id = ? and peer_type = ? and peer_id = ? and deleted = 0"
		do    = &Dialogs{}
	)
	err = m.db.QueryRowPartial(ctx, do, query, userId, peerType, peerId)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return nil, &NotFoundError{
				Resource: "dialogs",
				Key:      fmt.Sprintf("user_id=%v,peer_type=%v,peer_id=%v", userId, peerType, peerId),
				Cause:    err,
			}
		}
		err = fmt.Errorf("dialogs.SelectDialog: %w", err)
		return
	} else {
		rValue = do
	}

	return
}

// SelectDialog
// select id, user_id, peer_type, peer_id, peer_dialog_id, pinned, top_message, pinned_msg_id, read_inbox_max_id, read_outbox_max_id, unread_count, unread_mentions_count, unread_reactions_count, unread_mark, draft_type, draft_message_data, folder_id, folder_pinned, has_scheduled, ttl_period, theme_emoticon, wallpaper_id, wallpaper_overridden, date2 from dialogs where user_id = :user_id and peer_type = :peer_type and peer_id = :peer_id and deleted = 0
func (m *defaultDialogsTxModel) SelectDialog(userId int64, peerType int32, peerId int64) (rValue *Dialogs, err error) {
	var (
		query = "select id, user_id, peer_type, peer_id, peer_dialog_id, pinned, top_message, pinned_msg_id, read_inbox_max_id, read_outbox_max_id, unread_count, unread_mentions_count, unread_reactions_count, unread_mark, draft_type, draft_message_data, folder_id, folder_pinned, has_scheduled, ttl_period, theme_emoticon, wallpaper_id, wallpaper_overridden, date2 from dialogs where user_id = ? and peer_type = ? and peer_id = ? and deleted = 0"
		do    = &Dialogs{}
	)
	err = m.tx.QueryRowPartial(do, query, userId, peerType, peerId)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return nil, &NotFoundError{
				Resource: "dialogs",
				Key:      fmt.Sprintf("user_id=%v,peer_type=%v,peer_id=%v", userId, peerType, peerId),
				Cause:    err,
			}
		}
		err = fmt.Errorf("dialogs.SelectDialog: %w", err)
		return
	}
	rValue = do

	return
}

// SelectByPeerDialogId
// select id, user_id, peer_type, peer_id, peer_dialog_id, pinned, top_message, pinned_msg_id, read_inbox_max_id, read_outbox_max_id, unread_count, unread_mentions_count, unread_reactions_count, unread_mark, draft_type, draft_message_data, folder_id, folder_pinned, has_scheduled, ttl_period, theme_emoticon, wallpaper_id, wallpaper_overridden, date2 from dialogs where user_id = :user_id and peer_dialog_id = :peer_dialog_id and deleted = 0
func (m *defaultDialogsModel) SelectByPeerDialogId(ctx context.Context, userId int64, peerDialogId int64) (rValue *Dialogs, err error) {

	var (
		query = "select id, user_id, peer_type, peer_id, peer_dialog_id, pinned, top_message, pinned_msg_id, read_inbox_max_id, read_outbox_max_id, unread_count, unread_mentions_count, unread_reactions_count, unread_mark, draft_type, draft_message_data, folder_id, folder_pinned, has_scheduled, ttl_period, theme_emoticon, wallpaper_id, wallpaper_overridden, date2 from dialogs where user_id = ? and peer_dialog_id = ? and deleted = 0"
		do    = &Dialogs{}
	)
	err = m.db.QueryRowPartial(ctx, do, query, userId, peerDialogId)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return nil, &NotFoundError{
				Resource: "dialogs",
				Key:      fmt.Sprintf("user_id=%v,peer_dialog_id=%v", userId, peerDialogId),
				Cause:    err,
			}
		}
		err = fmt.Errorf("dialogs.SelectByPeerDialogId: %w", err)
		return
	} else {
		rValue = do
	}

	return
}

// SelectByPeerDialogId
// select id, user_id, peer_type, peer_id, peer_dialog_id, pinned, top_message, pinned_msg_id, read_inbox_max_id, read_outbox_max_id, unread_count, unread_mentions_count, unread_reactions_count, unread_mark, draft_type, draft_message_data, folder_id, folder_pinned, has_scheduled, ttl_period, theme_emoticon, wallpaper_id, wallpaper_overridden, date2 from dialogs where user_id = :user_id and peer_dialog_id = :peer_dialog_id and deleted = 0
func (m *defaultDialogsTxModel) SelectByPeerDialogId(userId int64, peerDialogId int64) (rValue *Dialogs, err error) {
	var (
		query = "select id, user_id, peer_type, peer_id, peer_dialog_id, pinned, top_message, pinned_msg_id, read_inbox_max_id, read_outbox_max_id, unread_count, unread_mentions_count, unread_reactions_count, unread_mark, draft_type, draft_message_data, folder_id, folder_pinned, has_scheduled, ttl_period, theme_emoticon, wallpaper_id, wallpaper_overridden, date2 from dialogs where user_id = ? and peer_dialog_id = ? and deleted = 0"
		do    = &Dialogs{}
	)
	err = m.tx.QueryRowPartial(do, query, userId, peerDialogId)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return nil, &NotFoundError{
				Resource: "dialogs",
				Key:      fmt.Sprintf("user_id=%v,peer_dialog_id=%v", userId, peerDialogId),
				Cause:    err,
			}
		}
		err = fmt.Errorf("dialogs.SelectByPeerDialogId: %w", err)
		return
	}
	rValue = do

	return
}

// SelectDialogs
// select id, user_id, peer_type, peer_id, peer_dialog_id, pinned, top_message, pinned_msg_id, read_inbox_max_id, read_outbox_max_id, unread_count, unread_mentions_count, unread_reactions_count, unread_mark, draft_type, draft_message_data, folder_id, folder_pinned, has_scheduled, ttl_period, theme_emoticon, wallpaper_id, wallpaper_overridden, date2 from dialogs where user_id = :user_id and folder_id = :folder_id and deleted = 0
func (m *defaultDialogsModel) SelectDialogs(ctx context.Context, userId int64, folderId int32) (rList []Dialogs, err error) {
	var (
		query  = "select id, user_id, peer_type, peer_id, peer_dialog_id, pinned, top_message, pinned_msg_id, read_inbox_max_id, read_outbox_max_id, unread_count, unread_mentions_count, unread_reactions_count, unread_mark, draft_type, draft_message_data, folder_id, folder_pinned, has_scheduled, ttl_period, theme_emoticon, wallpaper_id, wallpaper_overridden, date2 from dialogs where user_id = ? and folder_id = ? and deleted = 0"
		values []Dialogs
	)
	err = m.db.QueryRowsPartial(ctx, &values, query, userId, folderId)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			rList = []Dialogs{}
			err = nil
			return
		}
		err = fmt.Errorf("dialogs.SelectDialogs: %w", err)
		return
	}

	rList = values

	return
}

// SelectDialogs
// select id, user_id, peer_type, peer_id, peer_dialog_id, pinned, top_message, pinned_msg_id, read_inbox_max_id, read_outbox_max_id, unread_count, unread_mentions_count, unread_reactions_count, unread_mark, draft_type, draft_message_data, folder_id, folder_pinned, has_scheduled, ttl_period, theme_emoticon, wallpaper_id, wallpaper_overridden, date2 from dialogs where user_id = :user_id and folder_id = :folder_id and deleted = 0
func (m *defaultDialogsTxModel) SelectDialogs(userId int64, folderId int32) (rList []Dialogs, err error) {
	var (
		query  = "select id, user_id, peer_type, peer_id, peer_dialog_id, pinned, top_message, pinned_msg_id, read_inbox_max_id, read_outbox_max_id, unread_count, unread_mentions_count, unread_reactions_count, unread_mark, draft_type, draft_message_data, folder_id, folder_pinned, has_scheduled, ttl_period, theme_emoticon, wallpaper_id, wallpaper_overridden, date2 from dialogs where user_id = ? and folder_id = ? and deleted = 0"
		values []Dialogs
	)
	err = m.tx.QueryRowsPartial(&values, query, userId, folderId)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			rList = []Dialogs{}
			err = nil
			return
		}
		err = fmt.Errorf("dialogs.SelectDialogs: %w", err)
		return
	}

	rList = values

	return
}

// SelectDialogsWithCB
// select id, user_id, peer_type, peer_id, peer_dialog_id, pinned, top_message, pinned_msg_id, read_inbox_max_id, read_outbox_max_id, unread_count, unread_mentions_count, unread_reactions_count, unread_mark, draft_type, draft_message_data, folder_id, folder_pinned, has_scheduled, ttl_period, theme_emoticon, wallpaper_id, wallpaper_overridden, date2 from dialogs where user_id = :user_id and folder_id = :folder_id and deleted = 0
func (m *defaultDialogsModel) SelectDialogsWithCB(ctx context.Context, userId int64, folderId int32, cb func(sz, i int, v *Dialogs)) (rList []Dialogs, err error) {
	var (
		query  = "select id, user_id, peer_type, peer_id, peer_dialog_id, pinned, top_message, pinned_msg_id, read_inbox_max_id, read_outbox_max_id, unread_count, unread_mentions_count, unread_reactions_count, unread_mark, draft_type, draft_message_data, folder_id, folder_pinned, has_scheduled, ttl_period, theme_emoticon, wallpaper_id, wallpaper_overridden, date2 from dialogs where user_id = ? and folder_id = ? and deleted = 0"
		values []Dialogs
	)
	err = m.db.QueryRowsPartial(ctx, &values, query, userId, folderId)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			rList = []Dialogs{}
			err = nil
			return
		}
		err = fmt.Errorf("dialogs.SelectDialogsWithCB: %w", err)
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
// select id, user_id, peer_type, peer_id, peer_dialog_id, pinned, top_message, pinned_msg_id, read_inbox_max_id, read_outbox_max_id, unread_count, unread_mentions_count, unread_reactions_count, unread_mark, draft_type, draft_message_data, folder_id, folder_pinned, has_scheduled, ttl_period, theme_emoticon, wallpaper_id, wallpaper_overridden, date2 from dialogs where user_id = :user_id and folder_id = 0 and pinned = 0 and deleted = 0
func (m *defaultDialogsModel) SelectExcludePinnedDialogs(ctx context.Context, userId int64) (rList []Dialogs, err error) {
	var (
		query  = "select id, user_id, peer_type, peer_id, peer_dialog_id, pinned, top_message, pinned_msg_id, read_inbox_max_id, read_outbox_max_id, unread_count, unread_mentions_count, unread_reactions_count, unread_mark, draft_type, draft_message_data, folder_id, folder_pinned, has_scheduled, ttl_period, theme_emoticon, wallpaper_id, wallpaper_overridden, date2 from dialogs where user_id = ? and folder_id = 0 and pinned = 0 and deleted = 0"
		values []Dialogs
	)
	err = m.db.QueryRowsPartial(ctx, &values, query, userId)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			rList = []Dialogs{}
			err = nil
			return
		}
		err = fmt.Errorf("dialogs.SelectExcludePinnedDialogs: %w", err)
		return
	}

	rList = values

	return
}

// SelectExcludePinnedDialogs
// select id, user_id, peer_type, peer_id, peer_dialog_id, pinned, top_message, pinned_msg_id, read_inbox_max_id, read_outbox_max_id, unread_count, unread_mentions_count, unread_reactions_count, unread_mark, draft_type, draft_message_data, folder_id, folder_pinned, has_scheduled, ttl_period, theme_emoticon, wallpaper_id, wallpaper_overridden, date2 from dialogs where user_id = :user_id and folder_id = 0 and pinned = 0 and deleted = 0
func (m *defaultDialogsTxModel) SelectExcludePinnedDialogs(userId int64) (rList []Dialogs, err error) {
	var (
		query  = "select id, user_id, peer_type, peer_id, peer_dialog_id, pinned, top_message, pinned_msg_id, read_inbox_max_id, read_outbox_max_id, unread_count, unread_mentions_count, unread_reactions_count, unread_mark, draft_type, draft_message_data, folder_id, folder_pinned, has_scheduled, ttl_period, theme_emoticon, wallpaper_id, wallpaper_overridden, date2 from dialogs where user_id = ? and folder_id = 0 and pinned = 0 and deleted = 0"
		values []Dialogs
	)
	err = m.tx.QueryRowsPartial(&values, query, userId)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			rList = []Dialogs{}
			err = nil
			return
		}
		err = fmt.Errorf("dialogs.SelectExcludePinnedDialogs: %w", err)
		return
	}

	rList = values

	return
}

// SelectExcludePinnedDialogsWithCB
// select id, user_id, peer_type, peer_id, peer_dialog_id, pinned, top_message, pinned_msg_id, read_inbox_max_id, read_outbox_max_id, unread_count, unread_mentions_count, unread_reactions_count, unread_mark, draft_type, draft_message_data, folder_id, folder_pinned, has_scheduled, ttl_period, theme_emoticon, wallpaper_id, wallpaper_overridden, date2 from dialogs where user_id = :user_id and folder_id = 0 and pinned = 0 and deleted = 0
func (m *defaultDialogsModel) SelectExcludePinnedDialogsWithCB(ctx context.Context, userId int64, cb func(sz, i int, v *Dialogs)) (rList []Dialogs, err error) {
	var (
		query  = "select id, user_id, peer_type, peer_id, peer_dialog_id, pinned, top_message, pinned_msg_id, read_inbox_max_id, read_outbox_max_id, unread_count, unread_mentions_count, unread_reactions_count, unread_mark, draft_type, draft_message_data, folder_id, folder_pinned, has_scheduled, ttl_period, theme_emoticon, wallpaper_id, wallpaper_overridden, date2 from dialogs where user_id = ? and folder_id = 0 and pinned = 0 and deleted = 0"
		values []Dialogs
	)
	err = m.db.QueryRowsPartial(ctx, &values, query, userId)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			rList = []Dialogs{}
			err = nil
			return
		}
		err = fmt.Errorf("dialogs.SelectExcludePinnedDialogsWithCB: %w", err)
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

// SelectExcludeFolderPinnedDialogs
// select id, user_id, peer_type, peer_id, peer_dialog_id, pinned, top_message, pinned_msg_id, read_inbox_max_id, read_outbox_max_id, unread_count, unread_mentions_count, unread_reactions_count, unread_mark, draft_type, draft_message_data, folder_id, folder_pinned, has_scheduled, ttl_period, theme_emoticon, wallpaper_id, wallpaper_overridden, date2 from dialogs where user_id = :user_id and folder_id = 1 and folder_pinned = 0 and deleted = 0
func (m *defaultDialogsModel) SelectExcludeFolderPinnedDialogs(ctx context.Context, userId int64) (rList []Dialogs, err error) {
	var (
		query  = "select id, user_id, peer_type, peer_id, peer_dialog_id, pinned, top_message, pinned_msg_id, read_inbox_max_id, read_outbox_max_id, unread_count, unread_mentions_count, unread_reactions_count, unread_mark, draft_type, draft_message_data, folder_id, folder_pinned, has_scheduled, ttl_period, theme_emoticon, wallpaper_id, wallpaper_overridden, date2 from dialogs where user_id = ? and folder_id = 1 and folder_pinned = 0 and deleted = 0"
		values []Dialogs
	)
	err = m.db.QueryRowsPartial(ctx, &values, query, userId)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			rList = []Dialogs{}
			err = nil
			return
		}
		err = fmt.Errorf("dialogs.SelectExcludeFolderPinnedDialogs: %w", err)
		return
	}

	rList = values

	return
}

// SelectExcludeFolderPinnedDialogs
// select id, user_id, peer_type, peer_id, peer_dialog_id, pinned, top_message, pinned_msg_id, read_inbox_max_id, read_outbox_max_id, unread_count, unread_mentions_count, unread_reactions_count, unread_mark, draft_type, draft_message_data, folder_id, folder_pinned, has_scheduled, ttl_period, theme_emoticon, wallpaper_id, wallpaper_overridden, date2 from dialogs where user_id = :user_id and folder_id = 1 and folder_pinned = 0 and deleted = 0
func (m *defaultDialogsTxModel) SelectExcludeFolderPinnedDialogs(userId int64) (rList []Dialogs, err error) {
	var (
		query  = "select id, user_id, peer_type, peer_id, peer_dialog_id, pinned, top_message, pinned_msg_id, read_inbox_max_id, read_outbox_max_id, unread_count, unread_mentions_count, unread_reactions_count, unread_mark, draft_type, draft_message_data, folder_id, folder_pinned, has_scheduled, ttl_period, theme_emoticon, wallpaper_id, wallpaper_overridden, date2 from dialogs where user_id = ? and folder_id = 1 and folder_pinned = 0 and deleted = 0"
		values []Dialogs
	)
	err = m.tx.QueryRowsPartial(&values, query, userId)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			rList = []Dialogs{}
			err = nil
			return
		}
		err = fmt.Errorf("dialogs.SelectExcludeFolderPinnedDialogs: %w", err)
		return
	}

	rList = values

	return
}

// SelectExcludeFolderPinnedDialogsWithCB
// select id, user_id, peer_type, peer_id, peer_dialog_id, pinned, top_message, pinned_msg_id, read_inbox_max_id, read_outbox_max_id, unread_count, unread_mentions_count, unread_reactions_count, unread_mark, draft_type, draft_message_data, folder_id, folder_pinned, has_scheduled, ttl_period, theme_emoticon, wallpaper_id, wallpaper_overridden, date2 from dialogs where user_id = :user_id and folder_id = 1 and folder_pinned = 0 and deleted = 0
func (m *defaultDialogsModel) SelectExcludeFolderPinnedDialogsWithCB(ctx context.Context, userId int64, cb func(sz, i int, v *Dialogs)) (rList []Dialogs, err error) {
	var (
		query  = "select id, user_id, peer_type, peer_id, peer_dialog_id, pinned, top_message, pinned_msg_id, read_inbox_max_id, read_outbox_max_id, unread_count, unread_mentions_count, unread_reactions_count, unread_mark, draft_type, draft_message_data, folder_id, folder_pinned, has_scheduled, ttl_period, theme_emoticon, wallpaper_id, wallpaper_overridden, date2 from dialogs where user_id = ? and folder_id = 1 and folder_pinned = 0 and deleted = 0"
		values []Dialogs
	)
	err = m.db.QueryRowsPartial(ctx, &values, query, userId)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			rList = []Dialogs{}
			err = nil
			return
		}
		err = fmt.Errorf("dialogs.SelectExcludeFolderPinnedDialogsWithCB: %w", err)
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

// UpdateReadInboxMaxId
// update dialogs set unread_count = :unread_count, unread_mark = 0, read_inbox_max_id = :read_inbox_max_id where user_id = :user_id and peer_dialog_id = :peer_dialog_id
func (m *defaultDialogsModel) UpdateReadInboxMaxId(ctx context.Context, unreadCount int32, readInboxMaxId int32, userId int64, peerDialogId int64) (rowsAffected int64, err error) {

	var (
		query   = "update dialogs set unread_count = ?, unread_mark = 0, read_inbox_max_id = ? where user_id = ? and peer_dialog_id = ?"
		rResult sql.Result
	)

	rResult, err = m.db.Exec(ctx, query, unreadCount, readInboxMaxId, userId, peerDialogId)

	if err != nil {
		err = fmt.Errorf("dialogs.UpdateReadInboxMaxId exec: %w", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		err = fmt.Errorf("dialogs.UpdateReadInboxMaxId rows affected: %w", err)
		return
	}

	return
}

// UpdateReadInboxMaxId
// update dialogs set unread_count = :unread_count, unread_mark = 0, read_inbox_max_id = :read_inbox_max_id where user_id = :user_id and peer_dialog_id = :peer_dialog_id
func (m *defaultDialogsTxModel) UpdateReadInboxMaxId(unreadCount int32, readInboxMaxId int32, userId int64, peerDialogId int64) (rowsAffected int64, err error) {
	var (
		query   = "update dialogs set unread_count = ?, unread_mark = 0, read_inbox_max_id = ? where user_id = ? and peer_dialog_id = ?"
		rResult sql.Result
	)
	rResult, err = m.tx.Exec(query, unreadCount, readInboxMaxId, userId, peerDialogId)

	if err != nil {
		err = fmt.Errorf("dialogs.UpdateReadInboxMaxId exec: %w", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		err = fmt.Errorf("dialogs.UpdateReadInboxMaxId rows affected: %w", err)
		return
	}

	return
}

// UpdateReadOutboxMaxId
// update dialogs set read_outbox_max_id = :read_outbox_max_id where user_id = :user_id and peer_dialog_id = :peer_dialog_id
func (m *defaultDialogsModel) UpdateReadOutboxMaxId(ctx context.Context, readOutboxMaxId int32, userId int64, peerDialogId int64) (rowsAffected int64, err error) {

	var (
		query   = "update dialogs set read_outbox_max_id = ? where user_id = ? and peer_dialog_id = ?"
		rResult sql.Result
	)

	rResult, err = m.db.Exec(ctx, query, readOutboxMaxId, userId, peerDialogId)

	if err != nil {
		err = fmt.Errorf("dialogs.UpdateReadOutboxMaxId exec: %w", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		err = fmt.Errorf("dialogs.UpdateReadOutboxMaxId rows affected: %w", err)
		return
	}

	return
}

// UpdateReadOutboxMaxId
// update dialogs set read_outbox_max_id = :read_outbox_max_id where user_id = :user_id and peer_dialog_id = :peer_dialog_id
func (m *defaultDialogsTxModel) UpdateReadOutboxMaxId(readOutboxMaxId int32, userId int64, peerDialogId int64) (rowsAffected int64, err error) {
	var (
		query   = "update dialogs set read_outbox_max_id = ? where user_id = ? and peer_dialog_id = ?"
		rResult sql.Result
	)
	rResult, err = m.tx.Exec(query, readOutboxMaxId, userId, peerDialogId)

	if err != nil {
		err = fmt.Errorf("dialogs.UpdateReadOutboxMaxId exec: %w", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		err = fmt.Errorf("dialogs.UpdateReadOutboxMaxId rows affected: %w", err)
		return
	}

	return
}

// UpdateTopMessage
// update dialogs set top_message = :top_message where user_id = :user_id and peer_dialog_id = :peer_dialog_id
func (m *defaultDialogsModel) UpdateTopMessage(ctx context.Context, topMessage int32, userId int64, peerDialogId int64) (rowsAffected int64, err error) {

	var (
		query   = "update dialogs set top_message = ? where user_id = ? and peer_dialog_id = ?"
		rResult sql.Result
	)

	rResult, err = m.db.Exec(ctx, query, topMessage, userId, peerDialogId)

	if err != nil {
		err = fmt.Errorf("dialogs.UpdateTopMessage exec: %w", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		err = fmt.Errorf("dialogs.UpdateTopMessage rows affected: %w", err)
		return
	}

	return
}

// UpdateTopMessage
// update dialogs set top_message = :top_message where user_id = :user_id and peer_dialog_id = :peer_dialog_id
func (m *defaultDialogsTxModel) UpdateTopMessage(topMessage int32, userId int64, peerDialogId int64) (rowsAffected int64, err error) {
	var (
		query   = "update dialogs set top_message = ? where user_id = ? and peer_dialog_id = ?"
		rResult sql.Result
	)
	rResult, err = m.tx.Exec(query, topMessage, userId, peerDialogId)

	if err != nil {
		err = fmt.Errorf("dialogs.UpdateTopMessage exec: %w", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		err = fmt.Errorf("dialogs.UpdateTopMessage rows affected: %w", err)
		return
	}

	return
}

// UpdatePinnedMsgId
// update dialogs set pinned_msg_id = :pinned_msg_id where user_id = :user_id and peer_dialog_id = :peer_dialog_id
func (m *defaultDialogsModel) UpdatePinnedMsgId(ctx context.Context, pinnedMsgId int32, userId int64, peerDialogId int64) (rowsAffected int64, err error) {

	var (
		query   = "update dialogs set pinned_msg_id = ? where user_id = ? and peer_dialog_id = ?"
		rResult sql.Result
	)

	rResult, err = m.db.Exec(ctx, query, pinnedMsgId, userId, peerDialogId)

	if err != nil {
		err = fmt.Errorf("dialogs.UpdatePinnedMsgId exec: %w", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		err = fmt.Errorf("dialogs.UpdatePinnedMsgId rows affected: %w", err)
		return
	}

	return
}

// UpdatePinnedMsgId
// update dialogs set pinned_msg_id = :pinned_msg_id where user_id = :user_id and peer_dialog_id = :peer_dialog_id
func (m *defaultDialogsTxModel) UpdatePinnedMsgId(pinnedMsgId int32, userId int64, peerDialogId int64) (rowsAffected int64, err error) {
	var (
		query   = "update dialogs set pinned_msg_id = ? where user_id = ? and peer_dialog_id = ?"
		rResult sql.Result
	)
	rResult, err = m.tx.Exec(query, pinnedMsgId, userId, peerDialogId)

	if err != nil {
		err = fmt.Errorf("dialogs.UpdatePinnedMsgId exec: %w", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		err = fmt.Errorf("dialogs.UpdatePinnedMsgId rows affected: %w", err)
		return
	}

	return
}

// Delete
// delete from dialogs where user_id = :user_id and peer_type = :peer_type and peer_id = :peer_id
func (m *defaultDialogsModel) Delete(ctx context.Context, userId int64, peerType int32, peerId int64) (rowsAffected int64, err error) {

	var (
		query   = "delete from dialogs where user_id = ? and peer_type = ? and peer_id = ?"
		rResult sql.Result
	)
	rResult, err = m.db.Exec(ctx, query, userId, peerType, peerId)

	if err != nil {
		err = fmt.Errorf("dialogs.Delete exec: %w", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		err = fmt.Errorf("dialogs.Delete rows affected: %w", err)
		return
	}

	return
}

// Delete
// delete from dialogs where user_id = :user_id and peer_type = :peer_type and peer_id = :peer_id
func (m *defaultDialogsTxModel) Delete(userId int64, peerType int32, peerId int64) (rowsAffected int64, err error) {
	var (
		query   = "delete from dialogs where user_id = ? and peer_type = ? and peer_id = ?"
		rResult sql.Result
	)
	rResult, err = m.tx.Exec(query, userId, peerType, peerId)

	if err != nil {
		err = fmt.Errorf("dialogs.Delete exec: %w", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		err = fmt.Errorf("dialogs.Delete rows affected: %w", err)
		return
	}

	return
}

// SelectDialogsByGTReadInboxMaxId
// select user_id from dialogs where peer_type = :peer_type and peer_id = :peer_id and read_inbox_max_id >= :read_inbox_max_id and user_id != :user_id
func (m *defaultDialogsModel) SelectDialogsByGTReadInboxMaxId(ctx context.Context, peerType int32, peerId int64, readInboxMaxId int32, userId int64) (rList []int64, err error) {
	var query = "select user_id from dialogs where peer_type = ? and peer_id = ? and read_inbox_max_id >= ? and user_id != ?"
	err = m.db.QueryRowsPartial(ctx, &rList, query, peerType, peerId, readInboxMaxId, userId)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			rList = []int64{}
			err = nil
			return
		}
		err = fmt.Errorf("dialogs.SelectDialogsByGTReadInboxMaxId: %w", err)
	}

	return
}

// SelectDialogsByGTReadInboxMaxId
// select user_id from dialogs where peer_type = :peer_type and peer_id = :peer_id and read_inbox_max_id >= :read_inbox_max_id and user_id != :user_id
func (m *defaultDialogsTxModel) SelectDialogsByGTReadInboxMaxId(peerType int32, peerId int64, readInboxMaxId int32, userId int64) (rList []int64, err error) {
	var query = "select user_id from dialogs where peer_type = ? and peer_id = ? and read_inbox_max_id >= ? and user_id != ?"
	err = m.tx.QueryRowsPartial(&rList, query, peerType, peerId, readInboxMaxId, userId)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			rList = []int64{}
			err = nil
			return
		}
		err = fmt.Errorf("dialogs.SelectDialogsByGTReadInboxMaxId: %w", err)
	}

	return
}

// SelectDialogsByGTReadInboxMaxIdWithCB
// select user_id from dialogs where peer_type = :peer_type and peer_id = :peer_id and read_inbox_max_id >= :read_inbox_max_id and user_id != :user_id
func (m *defaultDialogsModel) SelectDialogsByGTReadInboxMaxIdWithCB(ctx context.Context, peerType int32, peerId int64, readInboxMaxId int32, userId int64, cb func(sz, i int, v int64)) (rList []int64, err error) {
	var query = "select user_id from dialogs where peer_type = ? and peer_id = ? and read_inbox_max_id >= ? and user_id != ?"
	err = m.db.QueryRowsPartial(ctx, &rList, query, peerType, peerId, readInboxMaxId, userId)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			rList = []int64{}
			err = nil
			return
		}
		err = fmt.Errorf("dialogs.SelectDialogsByGTReadInboxMaxIdWithCB: %w", err)
		return
	}

	if cb != nil {
		sz := len(rList)
		for i := 0; i < sz; i++ {
			cb(sz, i, rList[i])
		}
	}

	return
}

// UpdateCustomMap
// update dialogs set %s where user_id = :user_id and peer_type = :peer_type and peer_id = :peer_id
func (m *defaultDialogsModel) UpdateCustomMap(ctx context.Context, cMap map[string]interface{}, userId int64, peerType int32, peerId int64) (rowsAffected int64, err error) {

	names := make([]string, 0, len(cMap))
	aValues := make([]interface{}, 0, len(cMap))
	for k, v := range cMap {
		names = append(names, k+" = ?")
		aValues = append(aValues, v)
	}

	var (
		query   = fmt.Sprintf("update dialogs set %s where user_id = ? and peer_type = ? and peer_id = ?", strings.Join(names, ", "))
		rResult sql.Result
	)

	aValues = append(aValues, userId)
	aValues = append(aValues, peerType)
	aValues = append(aValues, peerId)

	rResult, err = m.db.Exec(ctx, query, aValues...)

	if err != nil {
		err = fmt.Errorf("dialogs.UpdateCustomMap exec: %w", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		err = fmt.Errorf("dialogs.UpdateCustomMap rows affected: %w", err)
		return
	}

	return
}

// UpdateCustomMap
// update dialogs set %s where user_id = :user_id and peer_type = :peer_type and peer_id = :peer_id
func (m *defaultDialogsTxModel) UpdateCustomMap(cMap map[string]interface{}, userId int64, peerType int32, peerId int64) (rowsAffected int64, err error) {
	names := make([]string, 0, len(cMap))
	aValues := make([]interface{}, 0, len(cMap))
	for k, v := range cMap {
		names = append(names, k+" = ?")
		aValues = append(aValues, v)
	}

	var (
		query   = fmt.Sprintf("update dialogs set %s where user_id = ? and peer_type = ? and peer_id = ?", strings.Join(names, ", "))
		rResult sql.Result
	)

	aValues = append(aValues, userId)
	aValues = append(aValues, peerType)
	aValues = append(aValues, peerId)

	rResult, err = m.tx.Exec(query, aValues...)

	if err != nil {
		err = fmt.Errorf("dialogs.UpdateCustomMap exec: %w", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		err = fmt.Errorf("dialogs.UpdateCustomMap rows affected: %w", err)
		return
	}

	return
}

// SaveDraft
// update dialogs set draft_type = :draft_type, draft_message_data = :draft_message_data where user_id = :user_id and peer_type = :peer_type and peer_id = :peer_id
func (m *defaultDialogsModel) SaveDraft(ctx context.Context, draftType int32, draftMessageData string, userId int64, peerType int32, peerId int64) (rowsAffected int64, err error) {

	var (
		query   = "update dialogs set draft_type = ?, draft_message_data = ? where user_id = ? and peer_type = ? and peer_id = ?"
		rResult sql.Result
	)

	rResult, err = m.db.Exec(ctx, query, draftType, draftMessageData, userId, peerType, peerId)

	if err != nil {
		err = fmt.Errorf("dialogs.SaveDraft exec: %w", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		err = fmt.Errorf("dialogs.SaveDraft rows affected: %w", err)
		return
	}

	return
}

// SaveDraft
// update dialogs set draft_type = :draft_type, draft_message_data = :draft_message_data where user_id = :user_id and peer_type = :peer_type and peer_id = :peer_id
func (m *defaultDialogsTxModel) SaveDraft(draftType int32, draftMessageData string, userId int64, peerType int32, peerId int64) (rowsAffected int64, err error) {
	var (
		query   = "update dialogs set draft_type = ?, draft_message_data = ? where user_id = ? and peer_type = ? and peer_id = ?"
		rResult sql.Result
	)
	rResult, err = m.tx.Exec(query, draftType, draftMessageData, userId, peerType, peerId)

	if err != nil {
		err = fmt.Errorf("dialogs.SaveDraft exec: %w", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		err = fmt.Errorf("dialogs.SaveDraft rows affected: %w", err)
		return
	}

	return
}

// SelectAllDrafts
// select id, user_id, peer_type, peer_id, draft_message_data from dialogs where user_id = :user_id and draft_type > 0
func (m *defaultDialogsModel) SelectAllDrafts(ctx context.Context, userId int64) (rList []Dialogs, err error) {
	var (
		query  = "select id, user_id, peer_type, peer_id, draft_message_data from dialogs where user_id = ? and draft_type > 0"
		values []Dialogs
	)
	err = m.db.QueryRowsPartial(ctx, &values, query, userId)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			rList = []Dialogs{}
			err = nil
			return
		}
		err = fmt.Errorf("dialogs.SelectAllDrafts: %w", err)
		return
	}

	rList = values

	return
}

// SelectAllDrafts
// select id, user_id, peer_type, peer_id, draft_message_data from dialogs where user_id = :user_id and draft_type > 0
func (m *defaultDialogsTxModel) SelectAllDrafts(userId int64) (rList []Dialogs, err error) {
	var (
		query  = "select id, user_id, peer_type, peer_id, draft_message_data from dialogs where user_id = ? and draft_type > 0"
		values []Dialogs
	)
	err = m.tx.QueryRowsPartial(&values, query, userId)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			rList = []Dialogs{}
			err = nil
			return
		}
		err = fmt.Errorf("dialogs.SelectAllDrafts: %w", err)
		return
	}

	rList = values

	return
}

// SelectAllDraftsWithCB
// select id, user_id, peer_type, peer_id, draft_message_data from dialogs where user_id = :user_id and draft_type > 0
func (m *defaultDialogsModel) SelectAllDraftsWithCB(ctx context.Context, userId int64, cb func(sz, i int, v *Dialogs)) (rList []Dialogs, err error) {
	var (
		query  = "select id, user_id, peer_type, peer_id, draft_message_data from dialogs where user_id = ? and draft_type > 0"
		values []Dialogs
	)
	err = m.db.QueryRowsPartial(ctx, &values, query, userId)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			rList = []Dialogs{}
			err = nil
			return
		}
		err = fmt.Errorf("dialogs.SelectAllDraftsWithCB: %w", err)
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

// ClearAllDrafts
// update dialogs set draft_type = 0, draft_message_data = 'null' where user_id = :user_id and draft_type = 2
func (m *defaultDialogsModel) ClearAllDrafts(ctx context.Context, userId int64) (rowsAffected int64, err error) {

	var (
		query   = "update dialogs set draft_type = 0, draft_message_data = 'null' where user_id = ? and draft_type = 2"
		rResult sql.Result
	)

	rResult, err = m.db.Exec(ctx, query, userId)

	if err != nil {
		err = fmt.Errorf("dialogs.ClearAllDrafts exec: %w", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		err = fmt.Errorf("dialogs.ClearAllDrafts rows affected: %w", err)
		return
	}

	return
}

// ClearAllDrafts
// update dialogs set draft_type = 0, draft_message_data = 'null' where user_id = :user_id and draft_type = 2
func (m *defaultDialogsTxModel) ClearAllDrafts(userId int64) (rowsAffected int64, err error) {
	var (
		query   = "update dialogs set draft_type = 0, draft_message_data = 'null' where user_id = ? and draft_type = 2"
		rResult sql.Result
	)
	rResult, err = m.tx.Exec(query, userId)

	if err != nil {
		err = fmt.Errorf("dialogs.ClearAllDrafts exec: %w", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		err = fmt.Errorf("dialogs.ClearAllDrafts rows affected: %w", err)
		return
	}

	return
}

// UpdatePeerFolderId
// update dialogs set folder_id = :folder_id where user_id = :user_id and peer_type = :peer_type and peer_id = :peer_id
func (m *defaultDialogsModel) UpdatePeerFolderId(ctx context.Context, folderId int32, userId int64, peerType int32, peerId int64) (rowsAffected int64, err error) {

	var (
		query   = "update dialogs set folder_id = ? where user_id = ? and peer_type = ? and peer_id = ?"
		rResult sql.Result
	)

	rResult, err = m.db.Exec(ctx, query, folderId, userId, peerType, peerId)

	if err != nil {
		err = fmt.Errorf("dialogs.UpdatePeerFolderId exec: %w", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		err = fmt.Errorf("dialogs.UpdatePeerFolderId rows affected: %w", err)
		return
	}

	return
}

// UpdatePeerFolderId
// update dialogs set folder_id = :folder_id where user_id = :user_id and peer_type = :peer_type and peer_id = :peer_id
func (m *defaultDialogsTxModel) UpdatePeerFolderId(folderId int32, userId int64, peerType int32, peerId int64) (rowsAffected int64, err error) {
	var (
		query   = "update dialogs set folder_id = ? where user_id = ? and peer_type = ? and peer_id = ?"
		rResult sql.Result
	)
	rResult, err = m.tx.Exec(query, folderId, userId, peerType, peerId)

	if err != nil {
		err = fmt.Errorf("dialogs.UpdatePeerFolderId exec: %w", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		err = fmt.Errorf("dialogs.UpdatePeerFolderId rows affected: %w", err)
		return
	}

	return
}

// UpdatePeerDialogListFolderId
// update dialogs set folder_id = :folder_id where user_id = :user_id and peer_dialog_id in (:idList)
func (m *defaultDialogsModel) UpdatePeerDialogListFolderId(ctx context.Context, folderId int32, userId int64, idList []int64) (rowsAffected int64, err error) {

	var (
		query   = fmt.Sprintf("update dialogs set folder_id = ? where user_id = ? and peer_dialog_id in (%s)", sqlx.InInt64List(idList))
		rResult sql.Result
	)

	if len(idList) == 0 {
		return
	}

	rResult, err = m.db.Exec(ctx, query, folderId, userId)

	if err != nil {
		err = fmt.Errorf("dialogs.UpdatePeerDialogListFolderId exec: %w", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		err = fmt.Errorf("dialogs.UpdatePeerDialogListFolderId rows affected: %w", err)
		return
	}

	return
}

// UpdatePeerDialogListFolderId
// update dialogs set folder_id = :folder_id where user_id = :user_id and peer_dialog_id in (:idList)
func (m *defaultDialogsTxModel) UpdatePeerDialogListFolderId(folderId int32, userId int64, idList []int64) (rowsAffected int64, err error) {
	var (
		query   = fmt.Sprintf("update dialogs set folder_id = ? where user_id = ? and peer_dialog_id in (%s)", sqlx.InInt64List(idList))
		rResult sql.Result
	)

	if len(idList) == 0 {
		return
	}

	rResult, err = m.tx.Exec(query, folderId, userId)

	if err != nil {
		err = fmt.Errorf("dialogs.UpdatePeerDialogListFolderId exec: %w", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		err = fmt.Errorf("dialogs.UpdatePeerDialogListFolderId rows affected: %w", err)
		return
	}

	return
}

// UpdatePeerDialogListPinned
// update dialogs set pinned = :pinned where user_id = :user_id and folder_id = 0 and peer_dialog_id in (:idList)
func (m *defaultDialogsModel) UpdatePeerDialogListPinned(ctx context.Context, pinned int64, userId int64, idList []int64) (rowsAffected int64, err error) {

	var (
		query   = fmt.Sprintf("update dialogs set pinned = ? where user_id = ? and folder_id = 0 and peer_dialog_id in (%s)", sqlx.InInt64List(idList))
		rResult sql.Result
	)

	if len(idList) == 0 {
		return
	}

	rResult, err = m.db.Exec(ctx, query, pinned, userId)

	if err != nil {
		err = fmt.Errorf("dialogs.UpdatePeerDialogListPinned exec: %w", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		err = fmt.Errorf("dialogs.UpdatePeerDialogListPinned rows affected: %w", err)
		return
	}

	return
}

// UpdatePeerDialogListPinned
// update dialogs set pinned = :pinned where user_id = :user_id and folder_id = 0 and peer_dialog_id in (:idList)
func (m *defaultDialogsTxModel) UpdatePeerDialogListPinned(pinned int64, userId int64, idList []int64) (rowsAffected int64, err error) {
	var (
		query   = fmt.Sprintf("update dialogs set pinned = ? where user_id = ? and folder_id = 0 and peer_dialog_id in (%s)", sqlx.InInt64List(idList))
		rResult sql.Result
	)

	if len(idList) == 0 {
		return
	}

	rResult, err = m.tx.Exec(query, pinned, userId)

	if err != nil {
		err = fmt.Errorf("dialogs.UpdatePeerDialogListPinned exec: %w", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		err = fmt.Errorf("dialogs.UpdatePeerDialogListPinned rows affected: %w", err)
		return
	}

	return
}

// UpdateFolderPeerDialogListPinned
// update dialogs set folder_pinned = :folder_pinned where user_id = :user_id and folder_id = 1 and peer_dialog_id in (:idList)
func (m *defaultDialogsModel) UpdateFolderPeerDialogListPinned(ctx context.Context, folderPinned int64, userId int64, idList []int64) (rowsAffected int64, err error) {

	var (
		query   = fmt.Sprintf("update dialogs set folder_pinned = ? where user_id = ? and folder_id = 1 and peer_dialog_id in (%s)", sqlx.InInt64List(idList))
		rResult sql.Result
	)

	if len(idList) == 0 {
		return
	}

	rResult, err = m.db.Exec(ctx, query, folderPinned, userId)

	if err != nil {
		err = fmt.Errorf("dialogs.UpdateFolderPeerDialogListPinned exec: %w", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		err = fmt.Errorf("dialogs.UpdateFolderPeerDialogListPinned rows affected: %w", err)
		return
	}

	return
}

// UpdateFolderPeerDialogListPinned
// update dialogs set folder_pinned = :folder_pinned where user_id = :user_id and folder_id = 1 and peer_dialog_id in (:idList)
func (m *defaultDialogsTxModel) UpdateFolderPeerDialogListPinned(folderPinned int64, userId int64, idList []int64) (rowsAffected int64, err error) {
	var (
		query   = fmt.Sprintf("update dialogs set folder_pinned = ? where user_id = ? and folder_id = 1 and peer_dialog_id in (%s)", sqlx.InInt64List(idList))
		rResult sql.Result
	)

	if len(idList) == 0 {
		return
	}

	rResult, err = m.tx.Exec(query, folderPinned, userId)

	if err != nil {
		err = fmt.Errorf("dialogs.UpdateFolderPeerDialogListPinned exec: %w", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		err = fmt.Errorf("dialogs.UpdateFolderPeerDialogListPinned rows affected: %w", err)
		return
	}

	return
}

// UpdateUnPinnedNotIdList
// update dialogs set pinned = 0 where user_id = :user_id and folder_id = 0 and pinned > 0 and peer_dialog_id not in (:idList)
func (m *defaultDialogsModel) UpdateUnPinnedNotIdList(ctx context.Context, userId int64, idList []int64) (rowsAffected int64, err error) {

	var (
		query   = fmt.Sprintf("update dialogs set pinned = 0 where user_id = ? and folder_id = 0 and pinned > 0 and peer_dialog_id not in (%s)", sqlx.InInt64List(idList))
		rResult sql.Result
	)

	if len(idList) == 0 {
		return
	}

	rResult, err = m.db.Exec(ctx, query, userId)

	if err != nil {
		err = fmt.Errorf("dialogs.UpdateUnPinnedNotIdList exec: %w", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		err = fmt.Errorf("dialogs.UpdateUnPinnedNotIdList rows affected: %w", err)
		return
	}

	return
}

// UpdateUnPinnedNotIdList
// update dialogs set pinned = 0 where user_id = :user_id and folder_id = 0 and pinned > 0 and peer_dialog_id not in (:idList)
func (m *defaultDialogsTxModel) UpdateUnPinnedNotIdList(userId int64, idList []int64) (rowsAffected int64, err error) {
	var (
		query   = fmt.Sprintf("update dialogs set pinned = 0 where user_id = ? and folder_id = 0 and pinned > 0 and peer_dialog_id not in (%s)", sqlx.InInt64List(idList))
		rResult sql.Result
	)

	if len(idList) == 0 {
		return
	}

	rResult, err = m.tx.Exec(query, userId)

	if err != nil {
		err = fmt.Errorf("dialogs.UpdateUnPinnedNotIdList exec: %w", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		err = fmt.Errorf("dialogs.UpdateUnPinnedNotIdList rows affected: %w", err)
		return
	}

	return
}

// UpdateFolderUnPinnedNotIdList
// update dialogs set folder_pinned = 0 where user_id = :user_id and folder_id = 1 and folder_pinned > 0 and peer_dialog_id not in (:idList)
func (m *defaultDialogsModel) UpdateFolderUnPinnedNotIdList(ctx context.Context, userId int64, idList []int64) (rowsAffected int64, err error) {

	var (
		query   = fmt.Sprintf("update dialogs set folder_pinned = 0 where user_id = ? and folder_id = 1 and folder_pinned > 0 and peer_dialog_id not in (%s)", sqlx.InInt64List(idList))
		rResult sql.Result
	)

	if len(idList) == 0 {
		return
	}

	rResult, err = m.db.Exec(ctx, query, userId)

	if err != nil {
		err = fmt.Errorf("dialogs.UpdateFolderUnPinnedNotIdList exec: %w", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		err = fmt.Errorf("dialogs.UpdateFolderUnPinnedNotIdList rows affected: %w", err)
		return
	}

	return
}

// UpdateFolderUnPinnedNotIdList
// update dialogs set folder_pinned = 0 where user_id = :user_id and folder_id = 1 and folder_pinned > 0 and peer_dialog_id not in (:idList)
func (m *defaultDialogsTxModel) UpdateFolderUnPinnedNotIdList(userId int64, idList []int64) (rowsAffected int64, err error) {
	var (
		query   = fmt.Sprintf("update dialogs set folder_pinned = 0 where user_id = ? and folder_id = 1 and folder_pinned > 0 and peer_dialog_id not in (%s)", sqlx.InInt64List(idList))
		rResult sql.Result
	)

	if len(idList) == 0 {
		return
	}

	rResult, err = m.tx.Exec(query, userId)

	if err != nil {
		err = fmt.Errorf("dialogs.UpdateFolderUnPinnedNotIdList exec: %w", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		err = fmt.Errorf("dialogs.UpdateFolderUnPinnedNotIdList rows affected: %w", err)
		return
	}

	return
}

// SelectAllDialogs
// select id, user_id, peer_type, peer_id, peer_dialog_id, pinned, top_message, pinned_msg_id, read_inbox_max_id, read_outbox_max_id, unread_count, unread_mentions_count, unread_reactions_count, unread_mark, draft_type, draft_message_data, folder_id, folder_pinned, has_scheduled, ttl_period, theme_emoticon, wallpaper_id, wallpaper_overridden, date2 from dialogs where user_id = :user_id and deleted = 0
func (m *defaultDialogsModel) SelectAllDialogs(ctx context.Context, userId int64) (rList []Dialogs, err error) {
	var (
		query  = "select id, user_id, peer_type, peer_id, peer_dialog_id, pinned, top_message, pinned_msg_id, read_inbox_max_id, read_outbox_max_id, unread_count, unread_mentions_count, unread_reactions_count, unread_mark, draft_type, draft_message_data, folder_id, folder_pinned, has_scheduled, ttl_period, theme_emoticon, wallpaper_id, wallpaper_overridden, date2 from dialogs where user_id = ? and deleted = 0"
		values []Dialogs
	)
	err = m.db.QueryRowsPartial(ctx, &values, query, userId)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			rList = []Dialogs{}
			err = nil
			return
		}
		err = fmt.Errorf("dialogs.SelectAllDialogs: %w", err)
		return
	}

	rList = values

	return
}

// SelectAllDialogs
// select id, user_id, peer_type, peer_id, peer_dialog_id, pinned, top_message, pinned_msg_id, read_inbox_max_id, read_outbox_max_id, unread_count, unread_mentions_count, unread_reactions_count, unread_mark, draft_type, draft_message_data, folder_id, folder_pinned, has_scheduled, ttl_period, theme_emoticon, wallpaper_id, wallpaper_overridden, date2 from dialogs where user_id = :user_id and deleted = 0
func (m *defaultDialogsTxModel) SelectAllDialogs(userId int64) (rList []Dialogs, err error) {
	var (
		query  = "select id, user_id, peer_type, peer_id, peer_dialog_id, pinned, top_message, pinned_msg_id, read_inbox_max_id, read_outbox_max_id, unread_count, unread_mentions_count, unread_reactions_count, unread_mark, draft_type, draft_message_data, folder_id, folder_pinned, has_scheduled, ttl_period, theme_emoticon, wallpaper_id, wallpaper_overridden, date2 from dialogs where user_id = ? and deleted = 0"
		values []Dialogs
	)
	err = m.tx.QueryRowsPartial(&values, query, userId)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			rList = []Dialogs{}
			err = nil
			return
		}
		err = fmt.Errorf("dialogs.SelectAllDialogs: %w", err)
		return
	}

	rList = values

	return
}

// SelectAllDialogsWithCB
// select id, user_id, peer_type, peer_id, peer_dialog_id, pinned, top_message, pinned_msg_id, read_inbox_max_id, read_outbox_max_id, unread_count, unread_mentions_count, unread_reactions_count, unread_mark, draft_type, draft_message_data, folder_id, folder_pinned, has_scheduled, ttl_period, theme_emoticon, wallpaper_id, wallpaper_overridden, date2 from dialogs where user_id = :user_id and deleted = 0
func (m *defaultDialogsModel) SelectAllDialogsWithCB(ctx context.Context, userId int64, cb func(sz, i int, v *Dialogs)) (rList []Dialogs, err error) {
	var (
		query  = "select id, user_id, peer_type, peer_id, peer_dialog_id, pinned, top_message, pinned_msg_id, read_inbox_max_id, read_outbox_max_id, unread_count, unread_mentions_count, unread_reactions_count, unread_mark, draft_type, draft_message_data, folder_id, folder_pinned, has_scheduled, ttl_period, theme_emoticon, wallpaper_id, wallpaper_overridden, date2 from dialogs where user_id = ? and deleted = 0"
		values []Dialogs
	)
	err = m.db.QueryRowsPartial(ctx, &values, query, userId)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			rList = []Dialogs{}
			err = nil
			return
		}
		err = fmt.Errorf("dialogs.SelectAllDialogsWithCB: %w", err)
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

// SelectDialogsByPeerType
// select id, user_id, peer_type, peer_id, peer_dialog_id, pinned, top_message, pinned_msg_id, read_inbox_max_id, read_outbox_max_id, unread_count, unread_mentions_count, unread_reactions_count, unread_mark, draft_type, draft_message_data, folder_id, folder_pinned, has_scheduled, ttl_period, theme_emoticon, wallpaper_id, wallpaper_overridden, date2 from dialogs where user_id = :user_id and peer_type in (:peerTypeList) and deleted = 0
func (m *defaultDialogsModel) SelectDialogsByPeerType(ctx context.Context, userId int64, peerTypeList []int32) (rList []Dialogs, err error) {
	var (
		query  = fmt.Sprintf("select id, user_id, peer_type, peer_id, peer_dialog_id, pinned, top_message, pinned_msg_id, read_inbox_max_id, read_outbox_max_id, unread_count, unread_mentions_count, unread_reactions_count, unread_mark, draft_type, draft_message_data, folder_id, folder_pinned, has_scheduled, ttl_period, theme_emoticon, wallpaper_id, wallpaper_overridden, date2 from dialogs where user_id = ? and peer_type in (%s) and deleted = 0", sqlx.InInt32List(peerTypeList))
		values []Dialogs
	)
	if len(peerTypeList) == 0 {
		rList = []Dialogs{}
		return
	}

	err = m.db.QueryRowsPartial(ctx, &values, query, userId)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			rList = []Dialogs{}
			err = nil
			return
		}
		err = fmt.Errorf("dialogs.SelectDialogsByPeerType: %w", err)
		return
	}

	rList = values

	return
}

// SelectDialogsByPeerType
// select id, user_id, peer_type, peer_id, peer_dialog_id, pinned, top_message, pinned_msg_id, read_inbox_max_id, read_outbox_max_id, unread_count, unread_mentions_count, unread_reactions_count, unread_mark, draft_type, draft_message_data, folder_id, folder_pinned, has_scheduled, ttl_period, theme_emoticon, wallpaper_id, wallpaper_overridden, date2 from dialogs where user_id = :user_id and peer_type in (:peerTypeList) and deleted = 0
func (m *defaultDialogsTxModel) SelectDialogsByPeerType(userId int64, peerTypeList []int32) (rList []Dialogs, err error) {
	var (
		query  = fmt.Sprintf("select id, user_id, peer_type, peer_id, peer_dialog_id, pinned, top_message, pinned_msg_id, read_inbox_max_id, read_outbox_max_id, unread_count, unread_mentions_count, unread_reactions_count, unread_mark, draft_type, draft_message_data, folder_id, folder_pinned, has_scheduled, ttl_period, theme_emoticon, wallpaper_id, wallpaper_overridden, date2 from dialogs where user_id = ? and peer_type in (%s) and deleted = 0", sqlx.InInt32List(peerTypeList))
		values []Dialogs
	)
	if len(peerTypeList) == 0 {
		rList = []Dialogs{}
		return
	}

	err = m.tx.QueryRowsPartial(&values, query, userId)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			rList = []Dialogs{}
			err = nil
			return
		}
		err = fmt.Errorf("dialogs.SelectDialogsByPeerType: %w", err)
		return
	}

	rList = values

	return
}

// SelectDialogsByPeerTypeWithCB
// select id, user_id, peer_type, peer_id, peer_dialog_id, pinned, top_message, pinned_msg_id, read_inbox_max_id, read_outbox_max_id, unread_count, unread_mentions_count, unread_reactions_count, unread_mark, draft_type, draft_message_data, folder_id, folder_pinned, has_scheduled, ttl_period, theme_emoticon, wallpaper_id, wallpaper_overridden, date2 from dialogs where user_id = :user_id and peer_type in (:peerTypeList) and deleted = 0
func (m *defaultDialogsModel) SelectDialogsByPeerTypeWithCB(ctx context.Context, userId int64, peerTypeList []int32, cb func(sz, i int, v *Dialogs)) (rList []Dialogs, err error) {
	var (
		query  = fmt.Sprintf("select id, user_id, peer_type, peer_id, peer_dialog_id, pinned, top_message, pinned_msg_id, read_inbox_max_id, read_outbox_max_id, unread_count, unread_mentions_count, unread_reactions_count, unread_mark, draft_type, draft_message_data, folder_id, folder_pinned, has_scheduled, ttl_period, theme_emoticon, wallpaper_id, wallpaper_overridden, date2 from dialogs where user_id = ? and peer_type in (%s) and deleted = 0", sqlx.InInt32List(peerTypeList))
		values []Dialogs
	)
	if len(peerTypeList) == 0 {
		rList = []Dialogs{}
		return
	}

	err = m.db.QueryRowsPartial(ctx, &values, query, userId)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			rList = []Dialogs{}
			err = nil
			return
		}
		err = fmt.Errorf("dialogs.SelectDialogsByPeerTypeWithCB: %w", err)
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

// UpdateUnreadCount
// update dialogs set unread_count = unread_count + (:unreadCount), unread_mentions_count = unread_mentions_count + (:unreadMentionsCount), unread_reactions_count = unread_reactions_count + (:unreadReactionsCount) where user_id = :user_id and peer_type = :peer_type and peer_id = :peer_id
func (m *defaultDialogsModel) UpdateUnreadCount(ctx context.Context, unreadCount int32, unreadMentionsCount int32, unreadReactionsCount int32, userId int64, peerType int32, peerId int64) (rowsAffected int64, err error) {

	var (
		query   = "update dialogs set unread_count = unread_count + (?), unread_mentions_count = unread_mentions_count + (?), unread_reactions_count = unread_reactions_count + (?) where user_id = ? and peer_type = ? and peer_id = ?"
		rResult sql.Result
	)

	rResult, err = m.db.Exec(ctx, query, unreadCount, unreadMentionsCount, unreadReactionsCount, userId, peerType, peerId)

	if err != nil {
		err = fmt.Errorf("dialogs.UpdateUnreadCount exec: %w", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		err = fmt.Errorf("dialogs.UpdateUnreadCount rows affected: %w", err)
		return
	}

	return
}

// UpdateUnreadCount
// update dialogs set unread_count = unread_count + (:unreadCount), unread_mentions_count = unread_mentions_count + (:unreadMentionsCount), unread_reactions_count = unread_reactions_count + (:unreadReactionsCount) where user_id = :user_id and peer_type = :peer_type and peer_id = :peer_id
func (m *defaultDialogsTxModel) UpdateUnreadCount(unreadCount int32, unreadMentionsCount int32, unreadReactionsCount int32, userId int64, peerType int32, peerId int64) (rowsAffected int64, err error) {
	var (
		query   = "update dialogs set unread_count = unread_count + (?), unread_mentions_count = unread_mentions_count + (?), unread_reactions_count = unread_reactions_count + (?) where user_id = ? and peer_type = ? and peer_id = ?"
		rResult sql.Result
	)
	rResult, err = m.tx.Exec(query, unreadCount, unreadMentionsCount, unreadReactionsCount, userId, peerType, peerId)

	if err != nil {
		err = fmt.Errorf("dialogs.UpdateUnreadCount exec: %w", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		err = fmt.Errorf("dialogs.UpdateUnreadCount rows affected: %w", err)
		return
	}

	return
}
