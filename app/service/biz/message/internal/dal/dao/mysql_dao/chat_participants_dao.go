/*
 * WARNING! All changes made in this file will be lost!
 *   Created from by 'dalgen'
 *
 * Copyright (c) 2024-present,  Teamgram Authors.
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package mysql_dao

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strings"

	"github.com/teamgram/marmota/pkg/stores/sqlx"
	"github.com/teamgram/teamgram-server/app/service/biz/message/internal/dal/dataobject"

	"github.com/zeromicro/go-zero/core/logx"
)

var _ *sql.Result
var _ = fmt.Sprintf
var _ = strings.Join
var _ = errors.Is

type ChatParticipantsDAO struct {
	db *sqlx.DB
}

func NewChatParticipantsDAO(db *sqlx.DB) *ChatParticipantsDAO {
	return &ChatParticipantsDAO{
		db: db,
	}
}

// Insert
// insert into chat_participants(chat_id, user_id, participant_type, inviter_user_id, invited_at, draft_message_data) values (:chat_id, :user_id, :participant_type, :inviter_user_id, :invited_at, ”)
func (dao *ChatParticipantsDAO) Insert(ctx context.Context, do *dataobject.ChatParticipantsDO) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into chat_participants(chat_id, user_id, participant_type, inviter_user_id, invited_at, draft_message_data) values (:chat_id, :user_id, :participant_type, :inviter_user_id, :invited_at, '')"
		r     sql.Result
	)

	r, err = dao.db.NamedExec(ctx, query, do)
	if err != nil {
		logx.WithContext(ctx).Errorf("namedExec in Insert(%v), error: %v", do, err)
		return
	}

	lastInsertId, err = r.LastInsertId()
	if err != nil {
		logx.WithContext(ctx).Errorf("lastInsertId in Insert(%v)_error: %v", do, err)
		return
	}
	rowsAffected, err = r.RowsAffected()
	if err != nil {
		logx.WithContext(ctx).Errorf("rowsAffected in Insert(%v)_error: %v", do, err)
	}

	return
}

// InsertTx
// insert into chat_participants(chat_id, user_id, participant_type, inviter_user_id, invited_at, draft_message_data) values (:chat_id, :user_id, :participant_type, :inviter_user_id, :invited_at, ”)
func (dao *ChatParticipantsDAO) InsertTx(tx *sqlx.Tx, do *dataobject.ChatParticipantsDO) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into chat_participants(chat_id, user_id, participant_type, inviter_user_id, invited_at, draft_message_data) values (:chat_id, :user_id, :participant_type, :inviter_user_id, :invited_at, '')"
		r     sql.Result
	)

	r, err = tx.NamedExec(query, do)
	if err != nil {
		logx.WithContext(tx.Context()).Errorf("namedExec in Insert(%v), error: %v", do, err)
		return
	}

	lastInsertId, err = r.LastInsertId()
	if err != nil {
		logx.WithContext(tx.Context()).Errorf("lastInsertId in Insert(%v)_error: %v", do, err)
		return
	}
	rowsAffected, err = r.RowsAffected()
	if err != nil {
		logx.WithContext(tx.Context()).Errorf("rowsAffected in Insert(%v)_error: %v", do, err)
	}

	return
}

// InsertBulk
// insert into chat_participants(chat_id, user_id, participant_type, inviter_user_id, invited_at, draft_message_data) values (:chat_id, :user_id, :participant_type, :inviter_user_id, :invited_at, ”)
func (dao *ChatParticipantsDAO) InsertBulk(ctx context.Context, doList []*dataobject.ChatParticipantsDO) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into chat_participants(chat_id, user_id, participant_type, inviter_user_id, invited_at, draft_message_data) values (:chat_id, :user_id, :participant_type, :inviter_user_id, :invited_at, '')"
		r     sql.Result
	)

	if len(doList) == 0 {
		return
	}

	r, err = dao.db.NamedExec(ctx, query, doList)
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
func (dao *ChatParticipantsDAO) InsertBulkTx(tx *sqlx.Tx, doList []*dataobject.ChatParticipantsDO) (lastInsertId, rowsAffected int64, err error) {
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
func (dao *ChatParticipantsDAO) InsertOrUpdate(ctx context.Context, do *dataobject.ChatParticipantsDO) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into chat_participants(chat_id, user_id, participant_type, inviter_user_id, invited_at, draft_message_data) values (:chat_id, :user_id, :participant_type, :inviter_user_id, :invited_at, '') on duplicate key update participant_type = values(participant_type), inviter_user_id = values(inviter_user_id), invited_at = values(invited_at), state = 0, kicked_at = 0, left_at = 0"
		r     sql.Result
	)

	r, err = dao.db.NamedExec(ctx, query, do)
	if err != nil {
		logx.WithContext(ctx).Errorf("namedExec in InsertOrUpdate(%v), error: %v", do, err)
		return
	}

	lastInsertId, err = r.LastInsertId()
	if err != nil {
		logx.WithContext(ctx).Errorf("lastInsertId in InsertOrUpdate(%v)_error: %v", do, err)
		return
	}
	rowsAffected, err = r.RowsAffected()
	if err != nil {
		logx.WithContext(ctx).Errorf("rowsAffected in InsertOrUpdate(%v)_error: %v", do, err)
	}

	return
}

// InsertOrUpdateTx
// insert into chat_participants(chat_id, user_id, participant_type, inviter_user_id, invited_at, draft_message_data) values (:chat_id, :user_id, :participant_type, :inviter_user_id, :invited_at, ”) on duplicate key update participant_type = values(participant_type), inviter_user_id = values(inviter_user_id), invited_at = values(invited_at), state = 0, kicked_at = 0, left_at = 0
func (dao *ChatParticipantsDAO) InsertOrUpdateTx(tx *sqlx.Tx, do *dataobject.ChatParticipantsDO) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into chat_participants(chat_id, user_id, participant_type, inviter_user_id, invited_at, draft_message_data) values (:chat_id, :user_id, :participant_type, :inviter_user_id, :invited_at, '') on duplicate key update participant_type = values(participant_type), inviter_user_id = values(inviter_user_id), invited_at = values(invited_at), state = 0, kicked_at = 0, left_at = 0"
		r     sql.Result
	)

	r, err = tx.NamedExec(query, do)
	if err != nil {
		logx.WithContext(tx.Context()).Errorf("namedExec in InsertOrUpdate(%v), error: %v", do, err)
		return
	}

	lastInsertId, err = r.LastInsertId()
	if err != nil {
		logx.WithContext(tx.Context()).Errorf("lastInsertId in InsertOrUpdate(%v)_error: %v", do, err)
		return
	}
	rowsAffected, err = r.RowsAffected()
	if err != nil {
		logx.WithContext(tx.Context()).Errorf("rowsAffected in InsertOrUpdate(%v)_error: %v", do, err)
	}

	return
}

// SelectList
// select id, user_id, chat_id, participant_type, is_pinned, top_message, pinned_msg_id, read_inbox_max_id, read_outbox_max_id, unread_count, unread_mentions_count, unread_mark, draft_type, draft_message_data, inviter_user_id, invited_at, state, date2 from chat_participants where chat_id = :chat_id
func (dao *ChatParticipantsDAO) SelectList(ctx context.Context, chatId int64) (rList []dataobject.ChatParticipantsDO, err error) {
	var (
		query  = "select id, user_id, chat_id, participant_type, is_pinned, top_message, pinned_msg_id, read_inbox_max_id, read_outbox_max_id, unread_count, unread_mentions_count, unread_mark, draft_type, draft_message_data, inviter_user_id, invited_at, state, date2 from chat_participants where chat_id = ?"
		values []dataobject.ChatParticipantsDO
	)
	err = dao.db.QueryRowsPartial(ctx, &values, query, chatId)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectList(_), error: %v", err)
		return
	}

	rList = values

	return
}

// SelectListWithCB
// select id, user_id, chat_id, participant_type, is_pinned, top_message, pinned_msg_id, read_inbox_max_id, read_outbox_max_id, unread_count, unread_mentions_count, unread_mark, draft_type, draft_message_data, inviter_user_id, invited_at, state, date2 from chat_participants where chat_id = :chat_id
func (dao *ChatParticipantsDAO) SelectListWithCB(ctx context.Context, chatId int64, cb func(sz, i int, v *dataobject.ChatParticipantsDO)) (rList []dataobject.ChatParticipantsDO, err error) {
	var (
		query  = "select id, user_id, chat_id, participant_type, is_pinned, top_message, pinned_msg_id, read_inbox_max_id, read_outbox_max_id, unread_count, unread_mentions_count, unread_mark, draft_type, draft_message_data, inviter_user_id, invited_at, state, date2 from chat_participants where chat_id = ?"
		values []dataobject.ChatParticipantsDO
	)
	err = dao.db.QueryRowsPartial(ctx, &values, query, chatId)

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
func (dao *ChatParticipantsDAO) SelectByParticipant(ctx context.Context, chatId int64, userId int64) (rValue *dataobject.ChatParticipantsDO, err error) {
	var (
		query = "select id, user_id, chat_id, participant_type, is_pinned, top_message, pinned_msg_id, read_inbox_max_id, read_outbox_max_id, unread_count, unread_mentions_count, unread_mark, draft_type, draft_message_data, inviter_user_id, invited_at, state, date2 from chat_participants where chat_id = ? and user_id = ?"
		do    = &dataobject.ChatParticipantsDO{}
	)
	err = dao.db.QueryRowPartial(ctx, do, query, chatId, userId)

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
func (dao *ChatParticipantsDAO) Update(ctx context.Context, participantType int32, inviterUserId int64, invitedAt int64, id int64) (rowsAffected int64, err error) {
	var (
		query   = "update chat_participants set participant_type = ?, inviter_user_id = ?, invited_at = ?, state = 0, kicked_at = 0, left_at = 0 where id = ?"
		rResult sql.Result
	)

	rResult, err = dao.db.Exec(ctx, query, participantType, inviterUserId, invitedAt, id)

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
func (dao *ChatParticipantsDAO) UpdateTx(tx *sqlx.Tx, participantType int32, inviterUserId int64, invitedAt int64, id int64) (rowsAffected int64, err error) {
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
func (dao *ChatParticipantsDAO) UpdateKicked(ctx context.Context, kickedAt int64, id int64) (rowsAffected int64, err error) {
	var (
		query   = "update chat_participants set state = 2, kicked_at = ? where id = ?"
		rResult sql.Result
	)

	rResult, err = dao.db.Exec(ctx, query, kickedAt, id)

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
func (dao *ChatParticipantsDAO) UpdateKickedTx(tx *sqlx.Tx, kickedAt int64, id int64) (rowsAffected int64, err error) {
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
func (dao *ChatParticipantsDAO) UpdateLeft(ctx context.Context, leftAt int64, id int64) (rowsAffected int64, err error) {
	var (
		query   = "update chat_participants set state = 1, left_at = ? where id = ?"
		rResult sql.Result
	)

	rResult, err = dao.db.Exec(ctx, query, leftAt, id)

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
func (dao *ChatParticipantsDAO) UpdateLeftTx(tx *sqlx.Tx, leftAt int64, id int64) (rowsAffected int64, err error) {
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
func (dao *ChatParticipantsDAO) UpdatePinnedMsgId(ctx context.Context, userId int64, chatId int64) (rowsAffected int64, err error) {
	var (
		query   = "update chat_participants set pinned_msg_id = ? where user_id = ? and chat_id = ?"
		rResult sql.Result
	)

	rResult, err = dao.db.Exec(ctx, query, userId, chatId)

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
func (dao *ChatParticipantsDAO) UpdatePinnedMsgIdTx(tx *sqlx.Tx, userId int64, chatId int64) (rowsAffected int64, err error) {
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
func (dao *ChatParticipantsDAO) UpdateParticipantType(ctx context.Context, participantType int32, id int64) (rowsAffected int64, err error) {
	var (
		query   = "update chat_participants set participant_type = ? where id = ?"
		rResult sql.Result
	)

	rResult, err = dao.db.Exec(ctx, query, participantType, id)

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
func (dao *ChatParticipantsDAO) UpdateParticipantTypeTx(tx *sqlx.Tx, participantType int32, id int64) (rowsAffected int64, err error) {
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
func (dao *ChatParticipantsDAO) SaveDraft(ctx context.Context, userId int64, chatId int64) (rowsAffected int64, err error) {
	var (
		query   = "update chat_participants set draft_type = 2, draft_message_data = ? where user_id = ? and chat_id = ?"
		rResult sql.Result
	)

	rResult, err = dao.db.Exec(ctx, query, userId, chatId)

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
func (dao *ChatParticipantsDAO) SaveDraftTx(tx *sqlx.Tx, userId int64, chatId int64) (rowsAffected int64, err error) {
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
func (dao *ChatParticipantsDAO) ClearDraft(ctx context.Context, userId int64, chatId int64) (rowsAffected int64, err error) {
	var (
		query   = "update chat_participants set draft_type = 0, draft_message_data = '' where user_id = ? and chat_id = ?"
		rResult sql.Result
	)

	rResult, err = dao.db.Exec(ctx, query, userId, chatId)

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
func (dao *ChatParticipantsDAO) ClearDraftTx(tx *sqlx.Tx, userId int64, chatId int64) (rowsAffected int64, err error) {
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
func (dao *ChatParticipantsDAO) SelectDraftList(ctx context.Context, userId int64) (rList []dataobject.ChatParticipantsDO, err error) {
	var (
		query  = "select user_id, chat_id, draft_type, draft_message_data from chat_participants where user_id = ?"
		values []dataobject.ChatParticipantsDO
	)
	err = dao.db.QueryRowsPartial(ctx, &values, query, userId)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectDraftList(_), error: %v", err)
		return
	}

	rList = values

	return
}

// SelectDraftListWithCB
// select user_id, chat_id, draft_type, draft_message_data from chat_participants where user_id = :user_id
func (dao *ChatParticipantsDAO) SelectDraftListWithCB(ctx context.Context, userId int64, cb func(sz, i int, v *dataobject.ChatParticipantsDO)) (rList []dataobject.ChatParticipantsDO, err error) {
	var (
		query  = "select user_id, chat_id, draft_type, draft_message_data from chat_participants where user_id = ?"
		values []dataobject.ChatParticipantsDO
	)
	err = dao.db.QueryRowsPartial(ctx, &values, query, userId)

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
func (dao *ChatParticipantsDAO) UpdateOutboxDialog(ctx context.Context, cMap map[string]interface{}, userId int64, chatId int64) (rowsAffected int64, err error) {
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

	rResult, err = dao.db.Exec(ctx, query, aValues...)

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
func (dao *ChatParticipantsDAO) UpdateOutboxDialogTx(tx *sqlx.Tx, cMap map[string]interface{}, userId int64, chatId int64) (rowsAffected int64, err error) {
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
func (dao *ChatParticipantsDAO) UpdateUnreadByPeer(ctx context.Context, userId int64, chatId int64) (rowsAffected int64, err error) {
	var (
		query   = "update chat_participants set unread_count = 0, unread_mark = 0, read_inbox_max_id = ? where user_id = ? and chat_id = ?"
		rResult sql.Result
	)

	rResult, err = dao.db.Exec(ctx, query, userId, chatId)

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
func (dao *ChatParticipantsDAO) UpdateUnreadByPeerTx(tx *sqlx.Tx, userId int64, chatId int64) (rowsAffected int64, err error) {
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
func (dao *ChatParticipantsDAO) UpdateReadOutboxMaxIdByPeer(ctx context.Context, userId int64, chatId int64) (rowsAffected int64, err error) {
	var (
		query   = "update chat_participants set read_outbox_max_id = ? where user_id = ? and chat_id = ?"
		rResult sql.Result
	)

	rResult, err = dao.db.Exec(ctx, query, userId, chatId)

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
func (dao *ChatParticipantsDAO) UpdateReadOutboxMaxIdByPeerTx(tx *sqlx.Tx, userId int64, chatId int64) (rowsAffected int64, err error) {
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
func (dao *ChatParticipantsDAO) SelectByOffsetId(ctx context.Context, userId int64, userId2 int32, limit int32) (rList []dataobject.ChatParticipantsDO, err error) {
	var (
		query  = "select id, user_id, chat_id, participant_type, is_pinned, top_message, read_inbox_max_id, read_outbox_max_id, unread_count, unread_mentions_count, unread_mark, draft_type, draft_message_data, inviter_user_id, invited_at, state, date2 from chat_participants where user_id = ? and chat_id in (select id from chats where id in (select chat_id from chat_participants where user_id = ?) and deactivated = 0) and top_message < ? and (state = 0 or state = 2) order by top_message desc limit ?"
		values []dataobject.ChatParticipantsDO
	)
	err = dao.db.QueryRowsPartial(ctx, &values, query, userId, userId2, limit)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectByOffsetId(_), error: %v", err)
		return
	}

	rList = values

	return
}

// SelectByOffsetIdWithCB
// select id, user_id, chat_id, participant_type, is_pinned, top_message, read_inbox_max_id, read_outbox_max_id, unread_count, unread_mentions_count, unread_mark, draft_type, draft_message_data, inviter_user_id, invited_at, state, date2 from chat_participants where user_id = :user_id and chat_id in (select id from chats where id in (select chat_id from chat_participants where user_id = :userId2) and deactivated = 0) and top_message < :top_message and (state = 0 or state = 2) order by top_message desc limit :limit
func (dao *ChatParticipantsDAO) SelectByOffsetIdWithCB(ctx context.Context, userId int64, userId2 int32, limit int32, cb func(sz, i int, v *dataobject.ChatParticipantsDO)) (rList []dataobject.ChatParticipantsDO, err error) {
	var (
		query  = "select id, user_id, chat_id, participant_type, is_pinned, top_message, read_inbox_max_id, read_outbox_max_id, unread_count, unread_mentions_count, unread_mark, draft_type, draft_message_data, inviter_user_id, invited_at, state, date2 from chat_participants where user_id = ? and chat_id in (select id from chats where id in (select chat_id from chat_participants where user_id = ?) and deactivated = 0) and top_message < ? and (state = 0 or state = 2) order by top_message desc limit ?"
		values []dataobject.ChatParticipantsDO
	)
	err = dao.db.QueryRowsPartial(ctx, &values, query, userId, userId2, limit)

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
func (dao *ChatParticipantsDAO) SelectExcludePinnedByOffsetId(ctx context.Context, userId int64, userId2 int32, limit int32) (rList []dataobject.ChatParticipantsDO, err error) {
	var (
		query  = "select id, user_id, chat_id, participant_type, is_pinned, top_message, pinned_msg_id, read_inbox_max_id, read_outbox_max_id, unread_count, unread_mentions_count, unread_mark, draft_type, draft_message_data, inviter_user_id, invited_at, state, date2 from chat_participants where user_id = ? and is_pinned = 0 and chat_id in (select id from chats where id in (select chat_id from chat_participants where user_id = ?) and deactivated = 0) and top_message < ? and (state = 0 or state = 2) order by top_message desc limit ?"
		values []dataobject.ChatParticipantsDO
	)
	err = dao.db.QueryRowsPartial(ctx, &values, query, userId, userId2, limit)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectExcludePinnedByOffsetId(_), error: %v", err)
		return
	}

	rList = values

	return
}

// SelectExcludePinnedByOffsetIdWithCB
// select id, user_id, chat_id, participant_type, is_pinned, top_message, pinned_msg_id, read_inbox_max_id, read_outbox_max_id, unread_count, unread_mentions_count, unread_mark, draft_type, draft_message_data, inviter_user_id, invited_at, state, date2 from chat_participants where user_id = :user_id and is_pinned = 0 and chat_id in (select id from chats where id in (select chat_id from chat_participants where user_id = :userId2) and deactivated = 0) and top_message < :top_message and (state = 0 or state = 2) order by top_message desc limit :limit
func (dao *ChatParticipantsDAO) SelectExcludePinnedByOffsetIdWithCB(ctx context.Context, userId int64, userId2 int32, limit int32, cb func(sz, i int, v *dataobject.ChatParticipantsDO)) (rList []dataobject.ChatParticipantsDO, err error) {
	var (
		query  = "select id, user_id, chat_id, participant_type, is_pinned, top_message, pinned_msg_id, read_inbox_max_id, read_outbox_max_id, unread_count, unread_mentions_count, unread_mark, draft_type, draft_message_data, inviter_user_id, invited_at, state, date2 from chat_participants where user_id = ? and is_pinned = 0 and chat_id in (select id from chats where id in (select chat_id from chat_participants where user_id = ?) and deactivated = 0) and top_message < ? and (state = 0 or state = 2) order by top_message desc limit ?"
		values []dataobject.ChatParticipantsDO
	)
	err = dao.db.QueryRowsPartial(ctx, &values, query, userId, userId2, limit)

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
func (dao *ChatParticipantsDAO) SelectListByChatIdList(ctx context.Context, userId int64, idList []int32) (rList []dataobject.ChatParticipantsDO, err error) {
	var (
		query  = fmt.Sprintf("select id, user_id, chat_id, participant_type, is_pinned, top_message, pinned_msg_id, read_inbox_max_id, read_outbox_max_id, unread_count, unread_mentions_count, unread_mark, draft_type, draft_message_data, inviter_user_id, invited_at, date2 from chat_participants where user_id = ? and chat_id in (%s) order by top_message desc", sqlx.InInt32List(idList))
		values []dataobject.ChatParticipantsDO
	)

	if len(idList) == 0 {
		rList = []dataobject.ChatParticipantsDO{}
		return
	}

	err = dao.db.QueryRowsPartial(ctx, &values, query, userId)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectListByChatIdList(_), error: %v", err)
		return
	}

	rList = values

	return
}

// SelectListByChatIdListWithCB
// select id, user_id, chat_id, participant_type, is_pinned, top_message, pinned_msg_id, read_inbox_max_id, read_outbox_max_id, unread_count, unread_mentions_count, unread_mark, draft_type, draft_message_data, inviter_user_id, invited_at, date2 from chat_participants where user_id = :user_id and chat_id in (:idList) order by top_message desc
func (dao *ChatParticipantsDAO) SelectListByChatIdListWithCB(ctx context.Context, userId int64, idList []int32, cb func(sz, i int, v *dataobject.ChatParticipantsDO)) (rList []dataobject.ChatParticipantsDO, err error) {
	var (
		query  = fmt.Sprintf("select id, user_id, chat_id, participant_type, is_pinned, top_message, pinned_msg_id, read_inbox_max_id, read_outbox_max_id, unread_count, unread_mentions_count, unread_mark, draft_type, draft_message_data, inviter_user_id, invited_at, date2 from chat_participants where user_id = ? and chat_id in (%s) order by top_message desc", sqlx.InInt32List(idList))
		values []dataobject.ChatParticipantsDO
	)

	if len(idList) == 0 {
		rList = []dataobject.ChatParticipantsDO{}
		return
	}

	err = dao.db.QueryRowsPartial(ctx, &values, query, userId)

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
func (dao *ChatParticipantsDAO) UpdatePinned(ctx context.Context, userId int64, chatId int64) (rowsAffected int64, err error) {
	var (
		query   = "update chat_participants set is_pinned = ? where user_id = ? and chat_id = ?"
		rResult sql.Result
	)

	rResult, err = dao.db.Exec(ctx, query, userId, chatId)

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
func (dao *ChatParticipantsDAO) UpdatePinnedTx(tx *sqlx.Tx, userId int64, chatId int64) (rowsAffected int64, err error) {
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
func (dao *ChatParticipantsDAO) SelectPinnedDialogs(ctx context.Context, userId int64) (rList []dataobject.ChatParticipantsDO, err error) {
	var (
		query  = "select id, user_id, chat_id, participant_type, is_pinned, top_message, pinned_msg_id, read_inbox_max_id, read_outbox_max_id, unread_count, unread_mentions_count, unread_mark, draft_type, draft_message_data, inviter_user_id, invited_at, date2 from chat_participants where user_id = ? and is_pinned = 1 and state = 0"
		values []dataobject.ChatParticipantsDO
	)
	err = dao.db.QueryRowsPartial(ctx, &values, query, userId)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectPinnedDialogs(_), error: %v", err)
		return
	}

	rList = values

	return
}

// SelectPinnedDialogsWithCB
// select id, user_id, chat_id, participant_type, is_pinned, top_message, pinned_msg_id, read_inbox_max_id, read_outbox_max_id, unread_count, unread_mentions_count, unread_mark, draft_type, draft_message_data, inviter_user_id, invited_at, date2 from chat_participants where user_id = :user_id and is_pinned = 1 and state = 0
func (dao *ChatParticipantsDAO) SelectPinnedDialogsWithCB(ctx context.Context, userId int64, cb func(sz, i int, v *dataobject.ChatParticipantsDO)) (rList []dataobject.ChatParticipantsDO, err error) {
	var (
		query  = "select id, user_id, chat_id, participant_type, is_pinned, top_message, pinned_msg_id, read_inbox_max_id, read_outbox_max_id, unread_count, unread_mentions_count, unread_mark, draft_type, draft_message_data, inviter_user_id, invited_at, date2 from chat_participants where user_id = ? and is_pinned = 1 and state = 0"
		values []dataobject.ChatParticipantsDO
	)
	err = dao.db.QueryRowsPartial(ctx, &values, query, userId)

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
func (dao *ChatParticipantsDAO) UpdateInboxDialog(ctx context.Context, cMap map[string]interface{}, userId int64, chatId int64) (rowsAffected int64, err error) {
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

	rResult, err = dao.db.Exec(ctx, query, aValues...)

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
func (dao *ChatParticipantsDAO) UpdateInboxDialogTx(tx *sqlx.Tx, cMap map[string]interface{}, userId int64, chatId int64) (rowsAffected int64, err error) {
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
func (dao *ChatParticipantsDAO) UpdateMarkDialogUnread(ctx context.Context, userId int64, chatId int64) (rowsAffected int64, err error) {
	var (
		query   = "update chat_participants set unread_mark = ? where user_id = ? and chat_id = ?"
		rResult sql.Result
	)

	rResult, err = dao.db.Exec(ctx, query, userId, chatId)

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
func (dao *ChatParticipantsDAO) UpdateMarkDialogUnreadTx(tx *sqlx.Tx, userId int64, chatId int64) (rowsAffected int64, err error) {
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
func (dao *ChatParticipantsDAO) SelectMarkDialogUnreadList(ctx context.Context, userId int64) (rList []int64, err error) {
	var query = "select chat_id from chat_participants where user_id = ? and unread_mark = 1 and state = 0"
	err = dao.db.QueryRowsPartial(ctx, &rList, query, userId)

	if err != nil {
		logx.WithContext(ctx).Errorf("select in SelectMarkDialogUnreadList(_), error: %v", err)
	}

	return
}

// SelectMarkDialogUnreadListWithCB
// select chat_id from chat_participants where user_id = :user_id and unread_mark = 1 and state = 0
func (dao *ChatParticipantsDAO) SelectMarkDialogUnreadListWithCB(ctx context.Context, userId int64, cb func(sz, i int, v int64)) (rList []int64, err error) {
	var query = "select chat_id from chat_participants where user_id = ? and unread_mark = 1 and state = 0"
	err = dao.db.QueryRowsPartial(ctx, &rList, query, userId)

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
func (dao *ChatParticipantsDAO) UpdateCustomMap(ctx context.Context, cMap map[string]interface{}, userId int64, chatId int64) (rowsAffected int64, err error) {
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

	rResult, err = dao.db.Exec(ctx, query, aValues...)

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
func (dao *ChatParticipantsDAO) UpdateCustomMapTx(tx *sqlx.Tx, cMap map[string]interface{}, userId int64, chatId int64) (rowsAffected int64, err error) {
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
