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
	"github.com/zeromicro/go-zero/core/logx"
)

var _ *sql.Result
var _ = fmt.Sprintf
var _ = strings.Join
var _ = errors.Is
var _ *sqlx.DB
var _ *logx.Logger

type (
	bizChatParticipantsModel interface {
		Insert(ctx context.Context, data *ChatParticipants) (lastInsertId, rowsAffected int64, err error)
		InsertTx(tx *sqlx.Tx, data *ChatParticipants) (lastInsertId, rowsAffected int64, err error)

		InsertBulk(ctx context.Context, doList []*ChatParticipants) (lastInsertId, rowsAffected int64, err error)
		InsertBulkTx(tx *sqlx.Tx, doList []*ChatParticipants) (lastInsertId, rowsAffected int64, err error)

		InsertOrUpdate(ctx context.Context, data *ChatParticipants) (lastInsertId, rowsAffected int64, err error)
		InsertOrUpdateTx(tx *sqlx.Tx, data *ChatParticipants) (lastInsertId, rowsAffected int64, err error)

		SelectList(ctx context.Context, chatId int64) ([]ChatParticipants, error)
		SelectListWithCB(ctx context.Context, chatId int64, cb func(sz, i int, v *ChatParticipants)) ([]ChatParticipants, error)

		SelectByParticipant(ctx context.Context, chatId int64, userId int64) (*ChatParticipants, error)

		Update(ctx context.Context, participantType int32, inviterUserId int64, invitedAt int64, id int64) (rowsAffected int64, err error)
		UpdateTx(tx *sqlx.Tx, participantType int32, inviterUserId int64, invitedAt int64, id int64) (rowsAffected int64, err error)

		UpdateKicked(ctx context.Context, kickedAt int64, id int64) (rowsAffected int64, err error)
		UpdateKickedTx(tx *sqlx.Tx, kickedAt int64, id int64) (rowsAffected int64, err error)

		UpdateLeft(ctx context.Context, leftAt int64, id int64) (rowsAffected int64, err error)
		UpdateLeftTx(tx *sqlx.Tx, leftAt int64, id int64) (rowsAffected int64, err error)

		UpdatePinnedMsgId(ctx context.Context, userId int64, chatId int64) (rowsAffected int64, err error)
		UpdatePinnedMsgIdTx(tx *sqlx.Tx, userId int64, chatId int64) (rowsAffected int64, err error)

		UpdateParticipantType(ctx context.Context, participantType int32, id int64) (rowsAffected int64, err error)
		UpdateParticipantTypeTx(tx *sqlx.Tx, participantType int32, id int64) (rowsAffected int64, err error)

		SaveDraft(ctx context.Context, userId int64, chatId int64) (rowsAffected int64, err error)
		SaveDraftTx(tx *sqlx.Tx, userId int64, chatId int64) (rowsAffected int64, err error)

		ClearDraft(ctx context.Context, userId int64, chatId int64) (rowsAffected int64, err error)
		ClearDraftTx(tx *sqlx.Tx, userId int64, chatId int64) (rowsAffected int64, err error)

		SelectDraftList(ctx context.Context, userId int64) ([]ChatParticipants, error)
		SelectDraftListWithCB(ctx context.Context, userId int64, cb func(sz, i int, v *ChatParticipants)) ([]ChatParticipants, error)

		UpdateOutboxDialog(ctx context.Context, cMap map[string]interface{}, userId int64, chatId int64) (rowsAffected int64, err error)
		UpdateOutboxDialogTx(tx *sqlx.Tx, cMap map[string]interface{}, userId int64, chatId int64) (rowsAffected int64, err error)

		UpdateUnreadByPeer(ctx context.Context, userId int64, chatId int64) (rowsAffected int64, err error)
		UpdateUnreadByPeerTx(tx *sqlx.Tx, userId int64, chatId int64) (rowsAffected int64, err error)

		UpdateReadOutboxMaxIdByPeer(ctx context.Context, userId int64, chatId int64) (rowsAffected int64, err error)
		UpdateReadOutboxMaxIdByPeerTx(tx *sqlx.Tx, userId int64, chatId int64) (rowsAffected int64, err error)

		SelectByOffsetId(ctx context.Context, userId int64, userId2 int32, limit int32) ([]ChatParticipants, error)
		SelectByOffsetIdWithCB(ctx context.Context, userId int64, userId2 int32, limit int32, cb func(sz, i int, v *ChatParticipants)) ([]ChatParticipants, error)

		SelectExcludePinnedByOffsetId(ctx context.Context, userId int64, userId2 int32, limit int32) ([]ChatParticipants, error)
		SelectExcludePinnedByOffsetIdWithCB(ctx context.Context, userId int64, userId2 int32, limit int32, cb func(sz, i int, v *ChatParticipants)) ([]ChatParticipants, error)

		SelectListByChatIdList(ctx context.Context, userId int64, idList []int32) ([]ChatParticipants, error)
		SelectListByChatIdListWithCB(ctx context.Context, userId int64, idList []int32, cb func(sz, i int, v *ChatParticipants)) ([]ChatParticipants, error)

		UpdatePinned(ctx context.Context, userId int64, chatId int64) (rowsAffected int64, err error)
		UpdatePinnedTx(tx *sqlx.Tx, userId int64, chatId int64) (rowsAffected int64, err error)

		SelectPinnedDialogs(ctx context.Context, userId int64) ([]ChatParticipants, error)
		SelectPinnedDialogsWithCB(ctx context.Context, userId int64, cb func(sz, i int, v *ChatParticipants)) ([]ChatParticipants, error)

		UpdateInboxDialog(ctx context.Context, cMap map[string]interface{}, userId int64, chatId int64) (rowsAffected int64, err error)
		UpdateInboxDialogTx(tx *sqlx.Tx, cMap map[string]interface{}, userId int64, chatId int64) (rowsAffected int64, err error)

		UpdateMarkDialogUnread(ctx context.Context, userId int64, chatId int64) (rowsAffected int64, err error)
		UpdateMarkDialogUnreadTx(tx *sqlx.Tx, userId int64, chatId int64) (rowsAffected int64, err error)

		SelectMarkDialogUnreadList(ctx context.Context, userId int64) ([]int64, error)
		SelectMarkDialogUnreadListWithCB(ctx context.Context, userId int64, cb func(sz, i int, v int64)) ([]int64, error)

		UpdateCustomMap(ctx context.Context, cMap map[string]interface{}, userId int64, chatId int64) (rowsAffected int64, err error)
		UpdateCustomMapTx(tx *sqlx.Tx, cMap map[string]interface{}, userId int64, chatId int64) (rowsAffected int64, err error)
	}
)

// Insert
// insert into chat_participants(chat_id, user_id, participant_type, inviter_user_id, invited_at, draft_message_data) values (:chat_id, :user_id, :participant_type, :inviter_user_id, :invited_at, ”)
func (m *defaultChatParticipantsModel) Insert(ctx context.Context, data *ChatParticipants) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into chat_participants(chat_id, user_id, participant_type, inviter_user_id, invited_at, draft_message_data) values (:chat_id, :user_id, :participant_type, :inviter_user_id, :invited_at, '')"
		r     sql.Result
	)

	r, err = m.db.NamedExec(ctx, query, data)
	if err != nil {
		logx.WithContext(ctx).Errorf("namedExec in Insert(%v), error: %v", data, err)
		return
	}

	lastInsertId, err = r.LastInsertId()
	if err != nil {
		logx.WithContext(ctx).Errorf("lastInsertId in Insert(%v)_error: %v", data, err)
		return
	}
	rowsAffected, err = r.RowsAffected()
	if err != nil {
		logx.WithContext(ctx).Errorf("rowsAffected in Insert(%v)_error: %v", data, err)
	}

	return
}

// InsertTx
// insert into chat_participants(chat_id, user_id, participant_type, inviter_user_id, invited_at, draft_message_data) values (:chat_id, :user_id, :participant_type, :inviter_user_id, :invited_at, ”)
func (m *defaultChatParticipantsModel) InsertTx(tx *sqlx.Tx, data *ChatParticipants) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into chat_participants(chat_id, user_id, participant_type, inviter_user_id, invited_at, draft_message_data) values (:chat_id, :user_id, :participant_type, :inviter_user_id, :invited_at, '')"
		r     sql.Result
	)

	r, err = tx.NamedExec(query, data)
	if err != nil {
		logx.WithContext(tx.Context()).Errorf("namedExec in Insert(%v), error: %v", data, err)
		return
	}

	lastInsertId, err = r.LastInsertId()
	if err != nil {
		logx.WithContext(tx.Context()).Errorf("lastInsertId in Insert(%v)_error: %v", data, err)
		return
	}
	rowsAffected, err = r.RowsAffected()
	if err != nil {
		logx.WithContext(tx.Context()).Errorf("rowsAffected in Insert(%v)_error: %v", data, err)
	}

	return
}

// InsertBulk
// insert into chat_participants(chat_id, user_id, participant_type, inviter_user_id, invited_at, draft_message_data) values (:chat_id, :user_id, :participant_type, :inviter_user_id, :invited_at, ”)
func (m *defaultChatParticipantsModel) InsertBulk(ctx context.Context, doList []*ChatParticipants) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into chat_participants(chat_id, user_id, participant_type, inviter_user_id, invited_at, draft_message_data) values (:chat_id, :user_id, :participant_type, :inviter_user_id, :invited_at, '')"
		r     sql.Result
	)

	if len(doList) == 0 {
		return
	}

	r, err = m.db.NamedExec(ctx, query, doList)
	if err != nil {
		logx.WithContext(ctx).Errorf("namedExec in InsertBulk(%v), error: %v", doList, err)
		return
	}

	lastInsertId, err = r.LastInsertId()
	if err != nil {
		logx.WithContext(ctx).Errorf("lastInsertId in InsertBulk(%v)_error: %v", doList, err)
		return
	}
	rowsAffected, err = r.RowsAffected()
	if err != nil {
		logx.WithContext(ctx).Errorf("rowsAffected in InsertBulk(%v)_error: %v", doList, err)
	}

	return
}

// InsertBulkTx
// insert into chat_participants(chat_id, user_id, participant_type, inviter_user_id, invited_at, draft_message_data) values (:chat_id, :user_id, :participant_type, :inviter_user_id, :invited_at, ”)
func (m *defaultChatParticipantsModel) InsertBulkTx(tx *sqlx.Tx, doList []*ChatParticipants) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into chat_participants(chat_id, user_id, participant_type, inviter_user_id, invited_at, draft_message_data) values (:chat_id, :user_id, :participant_type, :inviter_user_id, :invited_at, '')"
		r     sql.Result
	)

	if len(doList) == 0 {
		return
	}

	r, err = tx.NamedExec(query, doList)
	if err != nil {
		logx.WithContext(tx.Context()).Errorf("namedExec in InsertBulk(%v), error: %v", doList, err)
		return
	}

	lastInsertId, err = r.LastInsertId()
	if err != nil {
		logx.WithContext(tx.Context()).Errorf("lastInsertId in InsertBulk(%v)_error: %v", doList, err)
		return
	}
	rowsAffected, err = r.RowsAffected()
	if err != nil {
		logx.WithContext(tx.Context()).Errorf("rowsAffected in InsertBulk(%v)_error: %v", doList, err)
	}

	return
}

// InsertOrUpdate
// insert into chat_participants(chat_id, user_id, participant_type, inviter_user_id, invited_at, draft_message_data) values (:chat_id, :user_id, :participant_type, :inviter_user_id, :invited_at, ”) on duplicate key update participant_type = values(participant_type), inviter_user_id = values(inviter_user_id), invited_at = values(invited_at), state = 0, kicked_at = 0, left_at = 0
func (m *defaultChatParticipantsModel) InsertOrUpdate(ctx context.Context, data *ChatParticipants) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into chat_participants(chat_id, user_id, participant_type, inviter_user_id, invited_at, draft_message_data) values (:chat_id, :user_id, :participant_type, :inviter_user_id, :invited_at, '') on duplicate key update participant_type = values(participant_type), inviter_user_id = values(inviter_user_id), invited_at = values(invited_at), state = 0, kicked_at = 0, left_at = 0"
		r     sql.Result
	)

	r, err = m.db.NamedExec(ctx, query, data)
	if err != nil {
		logx.WithContext(ctx).Errorf("namedExec in InsertOrUpdate(%v), error: %v", data, err)
		return
	}

	lastInsertId, err = r.LastInsertId()
	if err != nil {
		logx.WithContext(ctx).Errorf("lastInsertId in InsertOrUpdate(%v)_error: %v", data, err)
		return
	}
	rowsAffected, err = r.RowsAffected()
	if err != nil {
		logx.WithContext(ctx).Errorf("rowsAffected in InsertOrUpdate(%v)_error: %v", data, err)
	}

	return
}

// InsertOrUpdateTx
// insert into chat_participants(chat_id, user_id, participant_type, inviter_user_id, invited_at, draft_message_data) values (:chat_id, :user_id, :participant_type, :inviter_user_id, :invited_at, ”) on duplicate key update participant_type = values(participant_type), inviter_user_id = values(inviter_user_id), invited_at = values(invited_at), state = 0, kicked_at = 0, left_at = 0
func (m *defaultChatParticipantsModel) InsertOrUpdateTx(tx *sqlx.Tx, data *ChatParticipants) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into chat_participants(chat_id, user_id, participant_type, inviter_user_id, invited_at, draft_message_data) values (:chat_id, :user_id, :participant_type, :inviter_user_id, :invited_at, '') on duplicate key update participant_type = values(participant_type), inviter_user_id = values(inviter_user_id), invited_at = values(invited_at), state = 0, kicked_at = 0, left_at = 0"
		r     sql.Result
	)

	r, err = tx.NamedExec(query, data)
	if err != nil {
		logx.WithContext(tx.Context()).Errorf("namedExec in InsertOrUpdate(%v), error: %v", data, err)
		return
	}

	lastInsertId, err = r.LastInsertId()
	if err != nil {
		logx.WithContext(tx.Context()).Errorf("lastInsertId in InsertOrUpdate(%v)_error: %v", data, err)
		return
	}
	rowsAffected, err = r.RowsAffected()
	if err != nil {
		logx.WithContext(tx.Context()).Errorf("rowsAffected in InsertOrUpdate(%v)_error: %v", data, err)
	}

	return
}

// SelectList
// select id, user_id, chat_id, participant_type, is_pinned, top_message, pinned_msg_id, read_inbox_max_id, read_outbox_max_id, unread_count, unread_mentions_count, unread_mark, draft_type, draft_message_data, inviter_user_id, invited_at, state, date2 from chat_participants where chat_id = :chat_id
func (m *defaultChatParticipantsModel) SelectList(ctx context.Context, chatId int64) (rList []ChatParticipants, err error) {
	var (
		query  = "select id, user_id, chat_id, participant_type, is_pinned, top_message, pinned_msg_id, read_inbox_max_id, read_outbox_max_id, unread_count, unread_mentions_count, unread_mark, draft_type, draft_message_data, inviter_user_id, invited_at, state, date2 from chat_participants where chat_id = ?"
		values []ChatParticipants
	)
	err = m.db.QueryRowsPartial(ctx, &values, query, chatId)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectList(_), error: %v", err)
		return
	}

	rList = values

	return
}

// SelectListWithCB
// select id, user_id, chat_id, participant_type, is_pinned, top_message, pinned_msg_id, read_inbox_max_id, read_outbox_max_id, unread_count, unread_mentions_count, unread_mark, draft_type, draft_message_data, inviter_user_id, invited_at, state, date2 from chat_participants where chat_id = :chat_id
func (m *defaultChatParticipantsModel) SelectListWithCB(ctx context.Context, chatId int64, cb func(sz, i int, v *ChatParticipants)) (rList []ChatParticipants, err error) {
	var (
		query  = "select id, user_id, chat_id, participant_type, is_pinned, top_message, pinned_msg_id, read_inbox_max_id, read_outbox_max_id, unread_count, unread_mentions_count, unread_mark, draft_type, draft_message_data, inviter_user_id, invited_at, state, date2 from chat_participants where chat_id = ?"
		values []ChatParticipants
	)
	err = m.db.QueryRowsPartial(ctx, &values, query, chatId)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectList(_), error: %v", err)
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

// SelectByParticipant
// select id, user_id, chat_id, participant_type, is_pinned, top_message, pinned_msg_id, read_inbox_max_id, read_outbox_max_id, unread_count, unread_mentions_count, unread_mark, draft_type, draft_message_data, inviter_user_id, invited_at, state, date2 from chat_participants where chat_id = :chat_id and user_id = :user_id
func (m *defaultChatParticipantsModel) SelectByParticipant(ctx context.Context, chatId int64, userId int64) (rValue *ChatParticipants, err error) {
	var (
		query = "select id, user_id, chat_id, participant_type, is_pinned, top_message, pinned_msg_id, read_inbox_max_id, read_outbox_max_id, unread_count, unread_mentions_count, unread_mark, draft_type, draft_message_data, inviter_user_id, invited_at, state, date2 from chat_participants where chat_id = ? and user_id = ?"
		do    = &ChatParticipants{}
	)
	err = m.db.QueryRowPartial(ctx, do, query, chatId, userId)

	if err != nil {
		if !errors.Is(err, sqlx.ErrNotFound) {
			logx.WithContext(ctx).Errorf("queryx in SelectByParticipant(_), error: %v", err)
			return
		} else {
			err = nil
		}
	} else {
		rValue = do
	}

	return
}

// Update
// update chat_participants set participant_type = :participant_type, inviter_user_id = :inviter_user_id, invited_at = :invited_at, state = 0, kicked_at = 0, left_at = 0 where id = :id
func (m *defaultChatParticipantsModel) Update(ctx context.Context, participantType int32, inviterUserId int64, invitedAt int64, id int64) (rowsAffected int64, err error) {
	var (
		query   = "update chat_participants set participant_type = ?, inviter_user_id = ?, invited_at = ?, state = 0, kicked_at = 0, left_at = 0 where id = ?"
		rResult sql.Result
	)

	rResult, err = m.db.Exec(ctx, query, participantType, inviterUserId, invitedAt, id)

	if err != nil {
		logx.WithContext(ctx).Errorf("exec in Update(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		logx.WithContext(ctx).Errorf("rowsAffected in Update(_), error: %v", err)
	}

	return
}

// UpdateTx
// update chat_participants set participant_type = :participant_type, inviter_user_id = :inviter_user_id, invited_at = :invited_at, state = 0, kicked_at = 0, left_at = 0 where id = :id
func (m *defaultChatParticipantsModel) UpdateTx(tx *sqlx.Tx, participantType int32, inviterUserId int64, invitedAt int64, id int64) (rowsAffected int64, err error) {
	var (
		query   = "update chat_participants set participant_type = ?, inviter_user_id = ?, invited_at = ?, state = 0, kicked_at = 0, left_at = 0 where id = ?"
		rResult sql.Result
	)
	rResult, err = tx.Exec(query, participantType, inviterUserId, invitedAt, id)

	if err != nil {
		logx.WithContext(tx.Context()).Errorf("exec in Update(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		logx.WithContext(tx.Context()).Errorf("rowsAffected in Update(_), error: %v", err)
	}

	return
}

// UpdateKicked
// update chat_participants set state = 2, kicked_at = :kicked_at where id = :id
func (m *defaultChatParticipantsModel) UpdateKicked(ctx context.Context, kickedAt int64, id int64) (rowsAffected int64, err error) {
	var (
		query   = "update chat_participants set state = 2, kicked_at = ? where id = ?"
		rResult sql.Result
	)

	rResult, err = m.db.Exec(ctx, query, kickedAt, id)

	if err != nil {
		logx.WithContext(ctx).Errorf("exec in UpdateKicked(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		logx.WithContext(ctx).Errorf("rowsAffected in UpdateKicked(_), error: %v", err)
	}

	return
}

// UpdateKickedTx
// update chat_participants set state = 2, kicked_at = :kicked_at where id = :id
func (m *defaultChatParticipantsModel) UpdateKickedTx(tx *sqlx.Tx, kickedAt int64, id int64) (rowsAffected int64, err error) {
	var (
		query   = "update chat_participants set state = 2, kicked_at = ? where id = ?"
		rResult sql.Result
	)
	rResult, err = tx.Exec(query, kickedAt, id)

	if err != nil {
		logx.WithContext(tx.Context()).Errorf("exec in UpdateKicked(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		logx.WithContext(tx.Context()).Errorf("rowsAffected in UpdateKicked(_), error: %v", err)
	}

	return
}

// UpdateLeft
// update chat_participants set state = 1, left_at = :left_at where id = :id
func (m *defaultChatParticipantsModel) UpdateLeft(ctx context.Context, leftAt int64, id int64) (rowsAffected int64, err error) {
	var (
		query   = "update chat_participants set state = 1, left_at = ? where id = ?"
		rResult sql.Result
	)

	rResult, err = m.db.Exec(ctx, query, leftAt, id)

	if err != nil {
		logx.WithContext(ctx).Errorf("exec in UpdateLeft(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		logx.WithContext(ctx).Errorf("rowsAffected in UpdateLeft(_), error: %v", err)
	}

	return
}

// UpdateLeftTx
// update chat_participants set state = 1, left_at = :left_at where id = :id
func (m *defaultChatParticipantsModel) UpdateLeftTx(tx *sqlx.Tx, leftAt int64, id int64) (rowsAffected int64, err error) {
	var (
		query   = "update chat_participants set state = 1, left_at = ? where id = ?"
		rResult sql.Result
	)
	rResult, err = tx.Exec(query, leftAt, id)

	if err != nil {
		logx.WithContext(tx.Context()).Errorf("exec in UpdateLeft(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		logx.WithContext(tx.Context()).Errorf("rowsAffected in UpdateLeft(_), error: %v", err)
	}

	return
}

// UpdatePinnedMsgId
// update chat_participants set pinned_msg_id = :pinned_msg_id where user_id = :user_id and chat_id = :chat_id
func (m *defaultChatParticipantsModel) UpdatePinnedMsgId(ctx context.Context, userId int64, chatId int64) (rowsAffected int64, err error) {
	var (
		query   = "update chat_participants set pinned_msg_id = ? where user_id = ? and chat_id = ?"
		rResult sql.Result
	)

	rResult, err = m.db.Exec(ctx, query, userId, chatId)

	if err != nil {
		logx.WithContext(ctx).Errorf("exec in UpdatePinnedMsgId(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		logx.WithContext(ctx).Errorf("rowsAffected in UpdatePinnedMsgId(_), error: %v", err)
	}

	return
}

// UpdatePinnedMsgIdTx
// update chat_participants set pinned_msg_id = :pinned_msg_id where user_id = :user_id and chat_id = :chat_id
func (m *defaultChatParticipantsModel) UpdatePinnedMsgIdTx(tx *sqlx.Tx, userId int64, chatId int64) (rowsAffected int64, err error) {
	var (
		query   = "update chat_participants set pinned_msg_id = ? where user_id = ? and chat_id = ?"
		rResult sql.Result
	)
	rResult, err = tx.Exec(query, userId, chatId)

	if err != nil {
		logx.WithContext(tx.Context()).Errorf("exec in UpdatePinnedMsgId(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		logx.WithContext(tx.Context()).Errorf("rowsAffected in UpdatePinnedMsgId(_), error: %v", err)
	}

	return
}

// UpdateParticipantType
// update chat_participants set participant_type = :participant_type where id = :id
func (m *defaultChatParticipantsModel) UpdateParticipantType(ctx context.Context, participantType int32, id int64) (rowsAffected int64, err error) {
	var (
		query   = "update chat_participants set participant_type = ? where id = ?"
		rResult sql.Result
	)

	rResult, err = m.db.Exec(ctx, query, participantType, id)

	if err != nil {
		logx.WithContext(ctx).Errorf("exec in UpdateParticipantType(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		logx.WithContext(ctx).Errorf("rowsAffected in UpdateParticipantType(_), error: %v", err)
	}

	return
}

// UpdateParticipantTypeTx
// update chat_participants set participant_type = :participant_type where id = :id
func (m *defaultChatParticipantsModel) UpdateParticipantTypeTx(tx *sqlx.Tx, participantType int32, id int64) (rowsAffected int64, err error) {
	var (
		query   = "update chat_participants set participant_type = ? where id = ?"
		rResult sql.Result
	)
	rResult, err = tx.Exec(query, participantType, id)

	if err != nil {
		logx.WithContext(tx.Context()).Errorf("exec in UpdateParticipantType(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		logx.WithContext(tx.Context()).Errorf("rowsAffected in UpdateParticipantType(_), error: %v", err)
	}

	return
}

// SaveDraft
// update chat_participants set draft_type = 2, draft_message_data = :draft_message_data where user_id = :user_id and chat_id = :chat_id
func (m *defaultChatParticipantsModel) SaveDraft(ctx context.Context, userId int64, chatId int64) (rowsAffected int64, err error) {
	var (
		query   = "update chat_participants set draft_type = 2, draft_message_data = ? where user_id = ? and chat_id = ?"
		rResult sql.Result
	)

	rResult, err = m.db.Exec(ctx, query, userId, chatId)

	if err != nil {
		logx.WithContext(ctx).Errorf("exec in SaveDraft(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		logx.WithContext(ctx).Errorf("rowsAffected in SaveDraft(_), error: %v", err)
	}

	return
}

// SaveDraftTx
// update chat_participants set draft_type = 2, draft_message_data = :draft_message_data where user_id = :user_id and chat_id = :chat_id
func (m *defaultChatParticipantsModel) SaveDraftTx(tx *sqlx.Tx, userId int64, chatId int64) (rowsAffected int64, err error) {
	var (
		query   = "update chat_participants set draft_type = 2, draft_message_data = ? where user_id = ? and chat_id = ?"
		rResult sql.Result
	)
	rResult, err = tx.Exec(query, userId, chatId)

	if err != nil {
		logx.WithContext(tx.Context()).Errorf("exec in SaveDraft(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		logx.WithContext(tx.Context()).Errorf("rowsAffected in SaveDraft(_), error: %v", err)
	}

	return
}

// ClearDraft
// update chat_participants set draft_type = 0, draft_message_data = ” where user_id = :user_id and chat_id = :chat_id
func (m *defaultChatParticipantsModel) ClearDraft(ctx context.Context, userId int64, chatId int64) (rowsAffected int64, err error) {
	var (
		query   = "update chat_participants set draft_type = 0, draft_message_data = '' where user_id = ? and chat_id = ?"
		rResult sql.Result
	)

	rResult, err = m.db.Exec(ctx, query, userId, chatId)

	if err != nil {
		logx.WithContext(ctx).Errorf("exec in ClearDraft(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		logx.WithContext(ctx).Errorf("rowsAffected in ClearDraft(_), error: %v", err)
	}

	return
}

// ClearDraftTx
// update chat_participants set draft_type = 0, draft_message_data = ” where user_id = :user_id and chat_id = :chat_id
func (m *defaultChatParticipantsModel) ClearDraftTx(tx *sqlx.Tx, userId int64, chatId int64) (rowsAffected int64, err error) {
	var (
		query   = "update chat_participants set draft_type = 0, draft_message_data = '' where user_id = ? and chat_id = ?"
		rResult sql.Result
	)
	rResult, err = tx.Exec(query, userId, chatId)

	if err != nil {
		logx.WithContext(tx.Context()).Errorf("exec in ClearDraft(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		logx.WithContext(tx.Context()).Errorf("rowsAffected in ClearDraft(_), error: %v", err)
	}

	return
}

// SelectDraftList
// select user_id, chat_id, draft_type, draft_message_data from chat_participants where user_id = :user_id
func (m *defaultChatParticipantsModel) SelectDraftList(ctx context.Context, userId int64) (rList []ChatParticipants, err error) {
	var (
		query  = "select user_id, chat_id, draft_type, draft_message_data from chat_participants where user_id = ?"
		values []ChatParticipants
	)
	err = m.db.QueryRowsPartial(ctx, &values, query, userId)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectDraftList(_), error: %v", err)
		return
	}

	rList = values

	return
}

// SelectDraftListWithCB
// select user_id, chat_id, draft_type, draft_message_data from chat_participants where user_id = :user_id
func (m *defaultChatParticipantsModel) SelectDraftListWithCB(ctx context.Context, userId int64, cb func(sz, i int, v *ChatParticipants)) (rList []ChatParticipants, err error) {
	var (
		query  = "select user_id, chat_id, draft_type, draft_message_data from chat_participants where user_id = ?"
		values []ChatParticipants
	)
	err = m.db.QueryRowsPartial(ctx, &values, query, userId)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectDraftList(_), error: %v", err)
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

// UpdateOutboxDialog
// update chat_participants set unread_count = 0, %s where user_id = :user_id and chat_id = :chat_id
func (m *defaultChatParticipantsModel) UpdateOutboxDialog(ctx context.Context, cMap map[string]interface{}, userId int64, chatId int64) (rowsAffected int64, err error) {
	names := make([]string, 0, len(cMap))
	aValues := make([]interface{}, 0, len(cMap))
	for k, v := range cMap {
		names = append(names, k+" = ?")
		aValues = append(aValues, v)
	}

	var (
		query   = fmt.Sprintf("update chat_participants set unread_count = 0, %s where user_id = ? and chat_id = ?", strings.Join(names, ", "))
		rResult sql.Result
	)

	aValues = append(aValues, userId)
	aValues = append(aValues, chatId)

	rResult, err = m.db.Exec(ctx, query, aValues...)

	if err != nil {
		logx.WithContext(ctx).Errorf("exec in UpdateOutboxDialog(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		logx.WithContext(ctx).Errorf("rowsAffected in UpdateOutboxDialog(_), error: %v", err)
	}

	return
}

// UpdateOutboxDialogTx
// update chat_participants set unread_count = 0, %s where user_id = :user_id and chat_id = :chat_id
func (m *defaultChatParticipantsModel) UpdateOutboxDialogTx(tx *sqlx.Tx, cMap map[string]interface{}, userId int64, chatId int64) (rowsAffected int64, err error) {
	names := make([]string, 0, len(cMap))
	aValues := make([]interface{}, 0, len(cMap))
	for k, v := range cMap {
		names = append(names, k+" = ?")
		aValues = append(aValues, v)
	}

	var (
		query   = fmt.Sprintf("update chat_participants set unread_count = 0, %s where user_id = ? and chat_id = ?", strings.Join(names, ", "))
		rResult sql.Result
	)

	aValues = append(aValues, userId)
	aValues = append(aValues, chatId)

	rResult, err = tx.Exec(query, aValues...)

	if err != nil {
		logx.WithContext(tx.Context()).Errorf("exec in UpdateOutboxDialog(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		logx.WithContext(tx.Context()).Errorf("rowsAffected in UpdateOutboxDialog(_), error: %v", err)
	}

	return
}

// UpdateUnreadByPeer
// update chat_participants set unread_count = 0, unread_mark = 0, read_inbox_max_id = :read_inbox_max_id where user_id = :user_id and chat_id = :chat_id
func (m *defaultChatParticipantsModel) UpdateUnreadByPeer(ctx context.Context, userId int64, chatId int64) (rowsAffected int64, err error) {
	var (
		query   = "update chat_participants set unread_count = 0, unread_mark = 0, read_inbox_max_id = ? where user_id = ? and chat_id = ?"
		rResult sql.Result
	)

	rResult, err = m.db.Exec(ctx, query, userId, chatId)

	if err != nil {
		logx.WithContext(ctx).Errorf("exec in UpdateUnreadByPeer(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		logx.WithContext(ctx).Errorf("rowsAffected in UpdateUnreadByPeer(_), error: %v", err)
	}

	return
}

// UpdateUnreadByPeerTx
// update chat_participants set unread_count = 0, unread_mark = 0, read_inbox_max_id = :read_inbox_max_id where user_id = :user_id and chat_id = :chat_id
func (m *defaultChatParticipantsModel) UpdateUnreadByPeerTx(tx *sqlx.Tx, userId int64, chatId int64) (rowsAffected int64, err error) {
	var (
		query   = "update chat_participants set unread_count = 0, unread_mark = 0, read_inbox_max_id = ? where user_id = ? and chat_id = ?"
		rResult sql.Result
	)
	rResult, err = tx.Exec(query, userId, chatId)

	if err != nil {
		logx.WithContext(tx.Context()).Errorf("exec in UpdateUnreadByPeer(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		logx.WithContext(tx.Context()).Errorf("rowsAffected in UpdateUnreadByPeer(_), error: %v", err)
	}

	return
}

// UpdateReadOutboxMaxIdByPeer
// update chat_participants set read_outbox_max_id = :read_outbox_max_id where user_id = :user_id and chat_id = :chat_id
func (m *defaultChatParticipantsModel) UpdateReadOutboxMaxIdByPeer(ctx context.Context, userId int64, chatId int64) (rowsAffected int64, err error) {
	var (
		query   = "update chat_participants set read_outbox_max_id = ? where user_id = ? and chat_id = ?"
		rResult sql.Result
	)

	rResult, err = m.db.Exec(ctx, query, userId, chatId)

	if err != nil {
		logx.WithContext(ctx).Errorf("exec in UpdateReadOutboxMaxIdByPeer(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		logx.WithContext(ctx).Errorf("rowsAffected in UpdateReadOutboxMaxIdByPeer(_), error: %v", err)
	}

	return
}

// UpdateReadOutboxMaxIdByPeerTx
// update chat_participants set read_outbox_max_id = :read_outbox_max_id where user_id = :user_id and chat_id = :chat_id
func (m *defaultChatParticipantsModel) UpdateReadOutboxMaxIdByPeerTx(tx *sqlx.Tx, userId int64, chatId int64) (rowsAffected int64, err error) {
	var (
		query   = "update chat_participants set read_outbox_max_id = ? where user_id = ? and chat_id = ?"
		rResult sql.Result
	)
	rResult, err = tx.Exec(query, userId, chatId)

	if err != nil {
		logx.WithContext(tx.Context()).Errorf("exec in UpdateReadOutboxMaxIdByPeer(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		logx.WithContext(tx.Context()).Errorf("rowsAffected in UpdateReadOutboxMaxIdByPeer(_), error: %v", err)
	}

	return
}

// SelectByOffsetId
// select id, user_id, chat_id, participant_type, is_pinned, top_message, read_inbox_max_id, read_outbox_max_id, unread_count, unread_mentions_count, unread_mark, draft_type, draft_message_data, inviter_user_id, invited_at, state, date2 from chat_participants where user_id = :user_id and chat_id in (select id from chats where id in (select chat_id from chat_participants where user_id = :userId2) and deactivated = 0) and top_message < :top_message and (state = 0 or state = 2) order by top_message desc limit :limit
func (m *defaultChatParticipantsModel) SelectByOffsetId(ctx context.Context, userId int64, userId2 int32, limit int32) (rList []ChatParticipants, err error) {
	var (
		query  = "select id, user_id, chat_id, participant_type, is_pinned, top_message, read_inbox_max_id, read_outbox_max_id, unread_count, unread_mentions_count, unread_mark, draft_type, draft_message_data, inviter_user_id, invited_at, state, date2 from chat_participants where user_id = ? and chat_id in (select id from chats where id in (select chat_id from chat_participants where user_id = ?) and deactivated = 0) and top_message < ? and (state = 0 or state = 2) order by top_message desc limit ?"
		values []ChatParticipants
	)
	err = m.db.QueryRowsPartial(ctx, &values, query, userId, userId2, limit)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectByOffsetId(_), error: %v", err)
		return
	}

	rList = values

	return
}

// SelectByOffsetIdWithCB
// select id, user_id, chat_id, participant_type, is_pinned, top_message, read_inbox_max_id, read_outbox_max_id, unread_count, unread_mentions_count, unread_mark, draft_type, draft_message_data, inviter_user_id, invited_at, state, date2 from chat_participants where user_id = :user_id and chat_id in (select id from chats where id in (select chat_id from chat_participants where user_id = :userId2) and deactivated = 0) and top_message < :top_message and (state = 0 or state = 2) order by top_message desc limit :limit
func (m *defaultChatParticipantsModel) SelectByOffsetIdWithCB(ctx context.Context, userId int64, userId2 int32, limit int32, cb func(sz, i int, v *ChatParticipants)) (rList []ChatParticipants, err error) {
	var (
		query  = "select id, user_id, chat_id, participant_type, is_pinned, top_message, read_inbox_max_id, read_outbox_max_id, unread_count, unread_mentions_count, unread_mark, draft_type, draft_message_data, inviter_user_id, invited_at, state, date2 from chat_participants where user_id = ? and chat_id in (select id from chats where id in (select chat_id from chat_participants where user_id = ?) and deactivated = 0) and top_message < ? and (state = 0 or state = 2) order by top_message desc limit ?"
		values []ChatParticipants
	)
	err = m.db.QueryRowsPartial(ctx, &values, query, userId, userId2, limit)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectByOffsetId(_), error: %v", err)
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

// SelectExcludePinnedByOffsetId
// select id, user_id, chat_id, participant_type, is_pinned, top_message, pinned_msg_id, read_inbox_max_id, read_outbox_max_id, unread_count, unread_mentions_count, unread_mark, draft_type, draft_message_data, inviter_user_id, invited_at, state, date2 from chat_participants where user_id = :user_id and is_pinned = 0 and chat_id in (select id from chats where id in (select chat_id from chat_participants where user_id = :userId2) and deactivated = 0) and top_message < :top_message and (state = 0 or state = 2) order by top_message desc limit :limit
func (m *defaultChatParticipantsModel) SelectExcludePinnedByOffsetId(ctx context.Context, userId int64, userId2 int32, limit int32) (rList []ChatParticipants, err error) {
	var (
		query  = "select id, user_id, chat_id, participant_type, is_pinned, top_message, pinned_msg_id, read_inbox_max_id, read_outbox_max_id, unread_count, unread_mentions_count, unread_mark, draft_type, draft_message_data, inviter_user_id, invited_at, state, date2 from chat_participants where user_id = ? and is_pinned = 0 and chat_id in (select id from chats where id in (select chat_id from chat_participants where user_id = ?) and deactivated = 0) and top_message < ? and (state = 0 or state = 2) order by top_message desc limit ?"
		values []ChatParticipants
	)
	err = m.db.QueryRowsPartial(ctx, &values, query, userId, userId2, limit)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectExcludePinnedByOffsetId(_), error: %v", err)
		return
	}

	rList = values

	return
}

// SelectExcludePinnedByOffsetIdWithCB
// select id, user_id, chat_id, participant_type, is_pinned, top_message, pinned_msg_id, read_inbox_max_id, read_outbox_max_id, unread_count, unread_mentions_count, unread_mark, draft_type, draft_message_data, inviter_user_id, invited_at, state, date2 from chat_participants where user_id = :user_id and is_pinned = 0 and chat_id in (select id from chats where id in (select chat_id from chat_participants where user_id = :userId2) and deactivated = 0) and top_message < :top_message and (state = 0 or state = 2) order by top_message desc limit :limit
func (m *defaultChatParticipantsModel) SelectExcludePinnedByOffsetIdWithCB(ctx context.Context, userId int64, userId2 int32, limit int32, cb func(sz, i int, v *ChatParticipants)) (rList []ChatParticipants, err error) {
	var (
		query  = "select id, user_id, chat_id, participant_type, is_pinned, top_message, pinned_msg_id, read_inbox_max_id, read_outbox_max_id, unread_count, unread_mentions_count, unread_mark, draft_type, draft_message_data, inviter_user_id, invited_at, state, date2 from chat_participants where user_id = ? and is_pinned = 0 and chat_id in (select id from chats where id in (select chat_id from chat_participants where user_id = ?) and deactivated = 0) and top_message < ? and (state = 0 or state = 2) order by top_message desc limit ?"
		values []ChatParticipants
	)
	err = m.db.QueryRowsPartial(ctx, &values, query, userId, userId2, limit)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectExcludePinnedByOffsetId(_), error: %v", err)
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

// SelectListByChatIdList
// select id, user_id, chat_id, participant_type, is_pinned, top_message, pinned_msg_id, read_inbox_max_id, read_outbox_max_id, unread_count, unread_mentions_count, unread_mark, draft_type, draft_message_data, inviter_user_id, invited_at, date2 from chat_participants where user_id = :user_id and chat_id in (:idList) order by top_message desc
func (m *defaultChatParticipantsModel) SelectListByChatIdList(ctx context.Context, userId int64, idList []int32) (rList []ChatParticipants, err error) {
	var (
		query  = fmt.Sprintf("select id, user_id, chat_id, participant_type, is_pinned, top_message, pinned_msg_id, read_inbox_max_id, read_outbox_max_id, unread_count, unread_mentions_count, unread_mark, draft_type, draft_message_data, inviter_user_id, invited_at, date2 from chat_participants where user_id = ? and chat_id in (%s) order by top_message desc", sqlx.InInt32List(idList))
		values []ChatParticipants
	)
	if len(idList) == 0 {
		rList = []ChatParticipants{}
		return
	}

	err = m.db.QueryRowsPartial(ctx, &values, query, userId)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectListByChatIdList(_), error: %v", err)
		return
	}

	rList = values

	return
}

// SelectListByChatIdListWithCB
// select id, user_id, chat_id, participant_type, is_pinned, top_message, pinned_msg_id, read_inbox_max_id, read_outbox_max_id, unread_count, unread_mentions_count, unread_mark, draft_type, draft_message_data, inviter_user_id, invited_at, date2 from chat_participants where user_id = :user_id and chat_id in (:idList) order by top_message desc
func (m *defaultChatParticipantsModel) SelectListByChatIdListWithCB(ctx context.Context, userId int64, idList []int32, cb func(sz, i int, v *ChatParticipants)) (rList []ChatParticipants, err error) {
	var (
		query  = fmt.Sprintf("select id, user_id, chat_id, participant_type, is_pinned, top_message, pinned_msg_id, read_inbox_max_id, read_outbox_max_id, unread_count, unread_mentions_count, unread_mark, draft_type, draft_message_data, inviter_user_id, invited_at, date2 from chat_participants where user_id = ? and chat_id in (%s) order by top_message desc", sqlx.InInt32List(idList))
		values []ChatParticipants
	)
	if len(idList) == 0 {
		rList = []ChatParticipants{}
		return
	}

	err = m.db.QueryRowsPartial(ctx, &values, query, userId)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectListByChatIdList(_), error: %v", err)
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

// UpdatePinned
// update chat_participants set is_pinned = :is_pinned where user_id = :user_id and chat_id = :chat_id
func (m *defaultChatParticipantsModel) UpdatePinned(ctx context.Context, userId int64, chatId int64) (rowsAffected int64, err error) {
	var (
		query   = "update chat_participants set is_pinned = ? where user_id = ? and chat_id = ?"
		rResult sql.Result
	)

	rResult, err = m.db.Exec(ctx, query, userId, chatId)

	if err != nil {
		logx.WithContext(ctx).Errorf("exec in UpdatePinned(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		logx.WithContext(ctx).Errorf("rowsAffected in UpdatePinned(_), error: %v", err)
	}

	return
}

// UpdatePinnedTx
// update chat_participants set is_pinned = :is_pinned where user_id = :user_id and chat_id = :chat_id
func (m *defaultChatParticipantsModel) UpdatePinnedTx(tx *sqlx.Tx, userId int64, chatId int64) (rowsAffected int64, err error) {
	var (
		query   = "update chat_participants set is_pinned = ? where user_id = ? and chat_id = ?"
		rResult sql.Result
	)
	rResult, err = tx.Exec(query, userId, chatId)

	if err != nil {
		logx.WithContext(tx.Context()).Errorf("exec in UpdatePinned(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		logx.WithContext(tx.Context()).Errorf("rowsAffected in UpdatePinned(_), error: %v", err)
	}

	return
}

// SelectPinnedDialogs
// select id, user_id, chat_id, participant_type, is_pinned, top_message, pinned_msg_id, read_inbox_max_id, read_outbox_max_id, unread_count, unread_mentions_count, unread_mark, draft_type, draft_message_data, inviter_user_id, invited_at, date2 from chat_participants where user_id = :user_id and is_pinned = 1 and state = 0
func (m *defaultChatParticipantsModel) SelectPinnedDialogs(ctx context.Context, userId int64) (rList []ChatParticipants, err error) {
	var (
		query  = "select id, user_id, chat_id, participant_type, is_pinned, top_message, pinned_msg_id, read_inbox_max_id, read_outbox_max_id, unread_count, unread_mentions_count, unread_mark, draft_type, draft_message_data, inviter_user_id, invited_at, date2 from chat_participants where user_id = ? and is_pinned = 1 and state = 0"
		values []ChatParticipants
	)
	err = m.db.QueryRowsPartial(ctx, &values, query, userId)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectPinnedDialogs(_), error: %v", err)
		return
	}

	rList = values

	return
}

// SelectPinnedDialogsWithCB
// select id, user_id, chat_id, participant_type, is_pinned, top_message, pinned_msg_id, read_inbox_max_id, read_outbox_max_id, unread_count, unread_mentions_count, unread_mark, draft_type, draft_message_data, inviter_user_id, invited_at, date2 from chat_participants where user_id = :user_id and is_pinned = 1 and state = 0
func (m *defaultChatParticipantsModel) SelectPinnedDialogsWithCB(ctx context.Context, userId int64, cb func(sz, i int, v *ChatParticipants)) (rList []ChatParticipants, err error) {
	var (
		query  = "select id, user_id, chat_id, participant_type, is_pinned, top_message, pinned_msg_id, read_inbox_max_id, read_outbox_max_id, unread_count, unread_mentions_count, unread_mark, draft_type, draft_message_data, inviter_user_id, invited_at, date2 from chat_participants where user_id = ? and is_pinned = 1 and state = 0"
		values []ChatParticipants
	)
	err = m.db.QueryRowsPartial(ctx, &values, query, userId)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectPinnedDialogs(_), error: %v", err)
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

// UpdateInboxDialog
// update chat_participants set unread_count = unread_count + 1, %s where user_id = :user_id and chat_id = :chat_id
func (m *defaultChatParticipantsModel) UpdateInboxDialog(ctx context.Context, cMap map[string]interface{}, userId int64, chatId int64) (rowsAffected int64, err error) {
	names := make([]string, 0, len(cMap))
	aValues := make([]interface{}, 0, len(cMap))
	for k, v := range cMap {
		names = append(names, k+" = ?")
		aValues = append(aValues, v)
	}

	var (
		query   = fmt.Sprintf("update chat_participants set unread_count = unread_count + 1, %s where user_id = ? and chat_id = ?", strings.Join(names, ", "))
		rResult sql.Result
	)

	aValues = append(aValues, userId)
	aValues = append(aValues, chatId)

	rResult, err = m.db.Exec(ctx, query, aValues...)

	if err != nil {
		logx.WithContext(ctx).Errorf("exec in UpdateInboxDialog(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		logx.WithContext(ctx).Errorf("rowsAffected in UpdateInboxDialog(_), error: %v", err)
	}

	return
}

// UpdateInboxDialogTx
// update chat_participants set unread_count = unread_count + 1, %s where user_id = :user_id and chat_id = :chat_id
func (m *defaultChatParticipantsModel) UpdateInboxDialogTx(tx *sqlx.Tx, cMap map[string]interface{}, userId int64, chatId int64) (rowsAffected int64, err error) {
	names := make([]string, 0, len(cMap))
	aValues := make([]interface{}, 0, len(cMap))
	for k, v := range cMap {
		names = append(names, k+" = ?")
		aValues = append(aValues, v)
	}

	var (
		query   = fmt.Sprintf("update chat_participants set unread_count = unread_count + 1, %s where user_id = ? and chat_id = ?", strings.Join(names, ", "))
		rResult sql.Result
	)

	aValues = append(aValues, userId)
	aValues = append(aValues, chatId)

	rResult, err = tx.Exec(query, aValues...)

	if err != nil {
		logx.WithContext(tx.Context()).Errorf("exec in UpdateInboxDialog(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		logx.WithContext(tx.Context()).Errorf("rowsAffected in UpdateInboxDialog(_), error: %v", err)
	}

	return
}

// UpdateMarkDialogUnread
// update chat_participants set unread_mark = :unread_mark where user_id = :user_id and chat_id = :chat_id
func (m *defaultChatParticipantsModel) UpdateMarkDialogUnread(ctx context.Context, userId int64, chatId int64) (rowsAffected int64, err error) {
	var (
		query   = "update chat_participants set unread_mark = ? where user_id = ? and chat_id = ?"
		rResult sql.Result
	)

	rResult, err = m.db.Exec(ctx, query, userId, chatId)

	if err != nil {
		logx.WithContext(ctx).Errorf("exec in UpdateMarkDialogUnread(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		logx.WithContext(ctx).Errorf("rowsAffected in UpdateMarkDialogUnread(_), error: %v", err)
	}

	return
}

// UpdateMarkDialogUnreadTx
// update chat_participants set unread_mark = :unread_mark where user_id = :user_id and chat_id = :chat_id
func (m *defaultChatParticipantsModel) UpdateMarkDialogUnreadTx(tx *sqlx.Tx, userId int64, chatId int64) (rowsAffected int64, err error) {
	var (
		query   = "update chat_participants set unread_mark = ? where user_id = ? and chat_id = ?"
		rResult sql.Result
	)
	rResult, err = tx.Exec(query, userId, chatId)

	if err != nil {
		logx.WithContext(tx.Context()).Errorf("exec in UpdateMarkDialogUnread(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		logx.WithContext(tx.Context()).Errorf("rowsAffected in UpdateMarkDialogUnread(_), error: %v", err)
	}

	return
}

// SelectMarkDialogUnreadList
// select chat_id from chat_participants where user_id = :user_id and unread_mark = 1 and state = 0
func (m *defaultChatParticipantsModel) SelectMarkDialogUnreadList(ctx context.Context, userId int64) (rList []int64, err error) {
	var query = "select chat_id from chat_participants where user_id = ? and unread_mark = 1 and state = 0"
	err = m.db.QueryRowsPartial(ctx, &rList, query, userId)

	if err != nil {
		logx.WithContext(ctx).Errorf("select in SelectMarkDialogUnreadList(_), error: %v", err)
	}

	return
}

// SelectMarkDialogUnreadListWithCB
// select chat_id from chat_participants where user_id = :user_id and unread_mark = 1 and state = 0
func (m *defaultChatParticipantsModel) SelectMarkDialogUnreadListWithCB(ctx context.Context, userId int64, cb func(sz, i int, v int64)) (rList []int64, err error) {
	var query = "select chat_id from chat_participants where user_id = ? and unread_mark = 1 and state = 0"
	err = m.db.QueryRowsPartial(ctx, &rList, query, userId)

	if err != nil {
		logx.WithContext(ctx).Errorf("select in SelectMarkDialogUnreadList(_), error: %v", err)
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
// update chat_participants set %s where user_id = :user_id and chat_id = :chat_id
func (m *defaultChatParticipantsModel) UpdateCustomMap(ctx context.Context, cMap map[string]interface{}, userId int64, chatId int64) (rowsAffected int64, err error) {
	names := make([]string, 0, len(cMap))
	aValues := make([]interface{}, 0, len(cMap))
	for k, v := range cMap {
		names = append(names, k+" = ?")
		aValues = append(aValues, v)
	}

	var (
		query   = fmt.Sprintf("update chat_participants set %s where user_id = ? and chat_id = ?", strings.Join(names, ", "))
		rResult sql.Result
	)

	aValues = append(aValues, userId)
	aValues = append(aValues, chatId)

	rResult, err = m.db.Exec(ctx, query, aValues...)

	if err != nil {
		logx.WithContext(ctx).Errorf("exec in UpdateCustomMap(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		logx.WithContext(ctx).Errorf("rowsAffected in UpdateCustomMap(_), error: %v", err)
	}

	return
}

// UpdateCustomMapTx
// update chat_participants set %s where user_id = :user_id and chat_id = :chat_id
func (m *defaultChatParticipantsModel) UpdateCustomMapTx(tx *sqlx.Tx, cMap map[string]interface{}, userId int64, chatId int64) (rowsAffected int64, err error) {
	names := make([]string, 0, len(cMap))
	aValues := make([]interface{}, 0, len(cMap))
	for k, v := range cMap {
		names = append(names, k+" = ?")
		aValues = append(aValues, v)
	}

	var (
		query   = fmt.Sprintf("update chat_participants set %s where user_id = ? and chat_id = ?", strings.Join(names, ", "))
		rResult sql.Result
	)

	aValues = append(aValues, userId)
	aValues = append(aValues, chatId)

	rResult, err = tx.Exec(query, aValues...)

	if err != nil {
		logx.WithContext(tx.Context()).Errorf("exec in UpdateCustomMap(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		logx.WithContext(tx.Context()).Errorf("rowsAffected in UpdateCustomMap(_), error: %v", err)
	}

	return
}
