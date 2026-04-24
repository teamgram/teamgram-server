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
	bizChatParticipantsModel interface {
		Insert(ctx context.Context, data *ChatParticipants) (lastInsertId, rowsAffected int64, err error)
		InsertTx(tx *sqlx.Tx, data *ChatParticipants) (lastInsertId, rowsAffected int64, err error)

		InsertBulk(ctx context.Context, doList []*ChatParticipants) (lastInsertId, rowsAffected int64, err error)
		InsertBulkTx(tx *sqlx.Tx, doList []*ChatParticipants) (lastInsertId, rowsAffected int64, err error)

		InsertOrUpdate(ctx context.Context, data *ChatParticipants) (lastInsertId, rowsAffected int64, err error)
		InsertOrUpdateTx(tx *sqlx.Tx, data *ChatParticipants) (lastInsertId, rowsAffected int64, err error)

		SelectList(ctx context.Context, chatId int64) ([]ChatParticipants, error)
		SelectListWithCB(ctx context.Context, chatId int64, cb func(sz, i int, v *ChatParticipants)) ([]ChatParticipants, error)

		SelectChatParticipantIdList(ctx context.Context, chatId int64) ([]int64, error)
		SelectChatParticipantIdListWithCB(ctx context.Context, chatId int64, cb func(sz, i int, v int64)) ([]int64, error)

		SelectByParticipantId(ctx context.Context, chatId int64, userId int64) (*ChatParticipants, error)

		SelectListByParticipantIdList(ctx context.Context, chatId int64, idList []int64) ([]ChatParticipants, error)
		SelectListByParticipantIdListWithCB(ctx context.Context, chatId int64, idList []int64, cb func(sz, i int, v *ChatParticipants)) ([]ChatParticipants, error)

		Update(ctx context.Context, participantType int32, inviterUserId int64, invitedAt int64, isBot bool, id int64) (rowsAffected int64, err error)
		UpdateTx(tx *sqlx.Tx, participantType int32, inviterUserId int64, invitedAt int64, isBot bool, id int64) (rowsAffected int64, err error)

		UpdateKicked(ctx context.Context, kickedAt int64, id int64) (rowsAffected int64, err error)
		UpdateKickedTx(tx *sqlx.Tx, kickedAt int64, id int64) (rowsAffected int64, err error)

		UpdateLeft(ctx context.Context, leftAt int64, id int64) (rowsAffected int64, err error)
		UpdateLeftTx(tx *sqlx.Tx, leftAt int64, id int64) (rowsAffected int64, err error)

		UpdateParticipantType(ctx context.Context, participantType int32, id int64) (rowsAffected int64, err error)
		UpdateParticipantTypeTx(tx *sqlx.Tx, participantType int32, id int64) (rowsAffected int64, err error)

		SelectUsersChatIdList(ctx context.Context, idList []int64) ([]ChatParticipants, error)
		SelectUsersChatIdListWithCB(ctx context.Context, idList []int64, cb func(sz, i int, v *ChatParticipants)) ([]ChatParticipants, error)

		SelectMyAdminList(ctx context.Context, userId int64) ([]int64, error)
		SelectMyAdminListWithCB(ctx context.Context, userId int64, cb func(sz, i int, v int64)) ([]int64, error)

		SelectMyAllList(ctx context.Context, userId int64) ([]int64, error)
		SelectMyAllListWithCB(ctx context.Context, userId int64, cb func(sz, i int, v int64)) ([]int64, error)

		UpdateStateByChatId(ctx context.Context, state int32, chatId int64) (rowsAffected int64, err error)
		UpdateStateByChatIdTx(tx *sqlx.Tx, state int32, chatId int64) (rowsAffected int64, err error)

		UpdateLink(ctx context.Context, link string, chatId int64, userId int64) (rowsAffected int64, err error)
		UpdateLinkTx(tx *sqlx.Tx, link string, chatId int64, userId int64) (rowsAffected int64, err error)

		UpdateLinkUsage(ctx context.Context, usage2 int32, chatId int64, userId int64) (rowsAffected int64, err error)
		UpdateLinkUsageTx(tx *sqlx.Tx, usage2 int32, chatId int64, userId int64) (rowsAffected int64, err error)
	}
)

// Insert
// insert into chat_participants(chat_id, user_id, participant_type, link, inviter_user_id, invited_at, is_bot, date2) values (:chat_id, :user_id, :participant_type, :link, :inviter_user_id, :invited_at, :is_bot, :date2)
func (m *defaultChatParticipantsModel) Insert(ctx context.Context, data *ChatParticipants) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into chat_participants(chat_id, user_id, participant_type, link, inviter_user_id, invited_at, is_bot, date2) values (:chat_id, :user_id, :participant_type, :link, :inviter_user_id, :invited_at, :is_bot, :date2)"
		r     sql.Result
	)

	r, err = m.db.NamedExec(ctx, query, data)
	if err != nil {
		err = fmt.Errorf("chat_participants.Insert named exec: %w", err)
		return
	}

	lastInsertId, err = r.LastInsertId()
	if err != nil {
		err = fmt.Errorf("chat_participants.Insert last insert id: %w", err)
		return
	}
	rowsAffected, err = r.RowsAffected()
	if err != nil {
		err = fmt.Errorf("chat_participants.Insert rows affected: %w", err)
	}

	return

}

// InsertTx
// insert into chat_participants(chat_id, user_id, participant_type, link, inviter_user_id, invited_at, is_bot, date2) values (:chat_id, :user_id, :participant_type, :link, :inviter_user_id, :invited_at, :is_bot, :date2)
func (m *defaultChatParticipantsModel) InsertTx(tx *sqlx.Tx, data *ChatParticipants) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into chat_participants(chat_id, user_id, participant_type, link, inviter_user_id, invited_at, is_bot, date2) values (:chat_id, :user_id, :participant_type, :link, :inviter_user_id, :invited_at, :is_bot, :date2)"
		r     sql.Result
	)

	r, err = tx.NamedExec(query, data)
	if err != nil {
		err = fmt.Errorf("chat_participants.InsertTx named exec: %w", err)
		return
	}

	lastInsertId, err = r.LastInsertId()
	if err != nil {
		err = fmt.Errorf("chat_participants.InsertTx last insert id: %w", err)
		return
	}
	rowsAffected, err = r.RowsAffected()
	if err != nil {
		err = fmt.Errorf("chat_participants.InsertTx rows affected: %w", err)
	}

	return
}

// InsertBulk
// insert into chat_participants(chat_id, user_id, participant_type, link, inviter_user_id, invited_at, is_bot, date2) values (:chat_id, :user_id, :participant_type, :link, :inviter_user_id, :invited_at, :is_bot, :date2)
func (m *defaultChatParticipantsModel) InsertBulk(ctx context.Context, doList []*ChatParticipants) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into chat_participants(chat_id, user_id, participant_type, link, inviter_user_id, invited_at, is_bot, date2) values (:chat_id, :user_id, :participant_type, :link, :inviter_user_id, :invited_at, :is_bot, :date2)"
		r     sql.Result
	)

	if len(doList) == 0 {
		return
	}

	r, err = m.db.NamedExec(ctx, query, doList)
	if err != nil {
		err = fmt.Errorf("chat_participants.InsertBulk named exec: %w", err)
		return
	}

	lastInsertId, err = r.LastInsertId()
	if err != nil {
		err = fmt.Errorf("chat_participants.InsertBulk last insert id: %w", err)
		return
	}
	rowsAffected, err = r.RowsAffected()
	if err != nil {
		err = fmt.Errorf("chat_participants.InsertBulk rows affected: %w", err)
	}

	return
}

// InsertBulkTx
// insert into chat_participants(chat_id, user_id, participant_type, link, inviter_user_id, invited_at, is_bot, date2) values (:chat_id, :user_id, :participant_type, :link, :inviter_user_id, :invited_at, :is_bot, :date2)
func (m *defaultChatParticipantsModel) InsertBulkTx(tx *sqlx.Tx, doList []*ChatParticipants) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into chat_participants(chat_id, user_id, participant_type, link, inviter_user_id, invited_at, is_bot, date2) values (:chat_id, :user_id, :participant_type, :link, :inviter_user_id, :invited_at, :is_bot, :date2)"
		r     sql.Result
	)

	if len(doList) == 0 {
		return
	}

	r, err = tx.NamedExec(query, doList)
	if err != nil {
		err = fmt.Errorf("chat_participants.InsertBulkTx named exec: %w", err)
		return
	}

	lastInsertId, err = r.LastInsertId()
	if err != nil {
		err = fmt.Errorf("chat_participants.InsertBulkTx last insert id: %w", err)
		return
	}
	rowsAffected, err = r.RowsAffected()
	if err != nil {
		err = fmt.Errorf("chat_participants.InsertBulkTx rows affected: %w", err)
	}

	return
}

// InsertOrUpdate
// insert into chat_participants(chat_id, user_id, participant_type, link, inviter_user_id, invited_at, is_bot, date2) values (:chat_id, :user_id, :participant_type, :link, :inviter_user_id, :invited_at, :is_bot, :date2) on duplicate key update participant_type = values(participant_type), inviter_user_id = values(inviter_user_id), link = values(link), invited_at = values(invited_at), is_bot = values(is_bot), state = 0, kicked_at = 0, left_at = 0, date2 = values(date2)
func (m *defaultChatParticipantsModel) InsertOrUpdate(ctx context.Context, data *ChatParticipants) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into chat_participants(chat_id, user_id, participant_type, link, inviter_user_id, invited_at, is_bot, date2) values (:chat_id, :user_id, :participant_type, :link, :inviter_user_id, :invited_at, :is_bot, :date2) on duplicate key update participant_type = values(participant_type), inviter_user_id = values(inviter_user_id), link = values(link), invited_at = values(invited_at), is_bot = values(is_bot), state = 0, kicked_at = 0, left_at = 0, date2 = values(date2)"
		r     sql.Result
	)

	r, err = m.db.NamedExec(ctx, query, data)
	if err != nil {
		err = fmt.Errorf("chat_participants.InsertOrUpdate named exec: %w", err)
		return
	}

	lastInsertId, err = r.LastInsertId()
	if err != nil {
		err = fmt.Errorf("chat_participants.InsertOrUpdate last insert id: %w", err)
		return
	}
	rowsAffected, err = r.RowsAffected()
	if err != nil {
		err = fmt.Errorf("chat_participants.InsertOrUpdate rows affected: %w", err)
	}

	return

}

// InsertOrUpdateTx
// insert into chat_participants(chat_id, user_id, participant_type, link, inviter_user_id, invited_at, is_bot, date2) values (:chat_id, :user_id, :participant_type, :link, :inviter_user_id, :invited_at, :is_bot, :date2) on duplicate key update participant_type = values(participant_type), inviter_user_id = values(inviter_user_id), link = values(link), invited_at = values(invited_at), is_bot = values(is_bot), state = 0, kicked_at = 0, left_at = 0, date2 = values(date2)
func (m *defaultChatParticipantsModel) InsertOrUpdateTx(tx *sqlx.Tx, data *ChatParticipants) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into chat_participants(chat_id, user_id, participant_type, link, inviter_user_id, invited_at, is_bot, date2) values (:chat_id, :user_id, :participant_type, :link, :inviter_user_id, :invited_at, :is_bot, :date2) on duplicate key update participant_type = values(participant_type), inviter_user_id = values(inviter_user_id), link = values(link), invited_at = values(invited_at), is_bot = values(is_bot), state = 0, kicked_at = 0, left_at = 0, date2 = values(date2)"
		r     sql.Result
	)

	r, err = tx.NamedExec(query, data)
	if err != nil {
		err = fmt.Errorf("chat_participants.InsertOrUpdateTx named exec: %w", err)
		return
	}

	lastInsertId, err = r.LastInsertId()
	if err != nil {
		err = fmt.Errorf("chat_participants.InsertOrUpdateTx last insert id: %w", err)
		return
	}
	rowsAffected, err = r.RowsAffected()
	if err != nil {
		err = fmt.Errorf("chat_participants.InsertOrUpdateTx rows affected: %w", err)
	}

	return
}

// SelectList
// select id, chat_id, user_id, participant_type, link, inviter_user_id, invited_at, kicked_at, left_at, is_bot, state, date2 from chat_participants where chat_id = :chat_id
func (m *defaultChatParticipantsModel) SelectList(ctx context.Context, chatId int64) (rList []ChatParticipants, err error) {
	var (
		query  = "select id, chat_id, user_id, participant_type, link, inviter_user_id, invited_at, kicked_at, left_at, is_bot, state, date2 from chat_participants where chat_id = ?"
		values []ChatParticipants
	)
	err = m.db.QueryRowsPartial(ctx, &values, query, chatId)

	if err != nil {
		err = fmt.Errorf("chat_participants.SelectList: %w", err)
		return
	}

	rList = values

	return
}

// SelectListWithCB
// select id, chat_id, user_id, participant_type, link, inviter_user_id, invited_at, kicked_at, left_at, is_bot, state, date2 from chat_participants where chat_id = :chat_id
func (m *defaultChatParticipantsModel) SelectListWithCB(ctx context.Context, chatId int64, cb func(sz, i int, v *ChatParticipants)) (rList []ChatParticipants, err error) {
	var (
		query  = "select id, chat_id, user_id, participant_type, link, inviter_user_id, invited_at, kicked_at, left_at, is_bot, state, date2 from chat_participants where chat_id = ?"
		values []ChatParticipants
	)
	err = m.db.QueryRowsPartial(ctx, &values, query, chatId)

	if err != nil {
		err = fmt.Errorf("chat_participants.SelectListWithCB: %w", err)
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

// SelectChatParticipantIdList
// select user_id from chat_participants where chat_id = :chat_id
func (m *defaultChatParticipantsModel) SelectChatParticipantIdList(ctx context.Context, chatId int64) (rList []int64, err error) {
	var query = "select user_id from chat_participants where chat_id = ?"
	err = m.db.QueryRowsPartial(ctx, &rList, query, chatId)

	if err != nil {
		err = fmt.Errorf("chat_participants.SelectChatParticipantIdList: %w", err)
	}

	return
}

// SelectChatParticipantIdListWithCB
// select user_id from chat_participants where chat_id = :chat_id
func (m *defaultChatParticipantsModel) SelectChatParticipantIdListWithCB(ctx context.Context, chatId int64, cb func(sz, i int, v int64)) (rList []int64, err error) {
	var query = "select user_id from chat_participants where chat_id = ?"
	err = m.db.QueryRowsPartial(ctx, &rList, query, chatId)

	if err != nil {
		err = fmt.Errorf("chat_participants.SelectChatParticipantIdListWithCB: %w", err)
	}

	if cb != nil {
		sz := len(rList)
		for i := 0; i < sz; i++ {
			cb(sz, i, rList[i])
		}
	}

	return
}

// SelectByParticipantId
// select id, chat_id, user_id, participant_type, link, inviter_user_id, invited_at, kicked_at, left_at, is_bot, state, date2 from chat_participants where chat_id = :chat_id and user_id = :user_id
func (m *defaultChatParticipantsModel) SelectByParticipantId(ctx context.Context, chatId int64, userId int64) (rValue *ChatParticipants, err error) {

	var (
		query = "select id, chat_id, user_id, participant_type, link, inviter_user_id, invited_at, kicked_at, left_at, is_bot, state, date2 from chat_participants where chat_id = ? and user_id = ?"
		do    = &ChatParticipants{}
	)
	err = m.db.QueryRowPartial(ctx, do, query, chatId, userId)

	if err != nil {
		if !errors.Is(err, sqlx.ErrNotFound) {
			err = fmt.Errorf("chat_participants.SelectByParticipantId: %w", err)
			return
		} else {
			err = nil
		}
	} else {
		rValue = do
	}

	return
}

// SelectListByParticipantIdList
// select id, chat_id, user_id, participant_type, link, inviter_user_id, invited_at, kicked_at, left_at, is_bot, state, date2 from chat_participants where chat_id = :chat_id and user_id in (:idList)
func (m *defaultChatParticipantsModel) SelectListByParticipantIdList(ctx context.Context, chatId int64, idList []int64) (rList []ChatParticipants, err error) {
	var (
		query  = fmt.Sprintf("select id, chat_id, user_id, participant_type, link, inviter_user_id, invited_at, kicked_at, left_at, is_bot, state, date2 from chat_participants where chat_id = ? and user_id in (%s)", sqlx.InInt64List(idList))
		values []ChatParticipants
	)
	if len(idList) == 0 {
		rList = []ChatParticipants{}
		return
	}

	err = m.db.QueryRowsPartial(ctx, &values, query, chatId)

	if err != nil {
		err = fmt.Errorf("chat_participants.SelectListByParticipantIdList: %w", err)
		return
	}

	rList = values

	return
}

// SelectListByParticipantIdListWithCB
// select id, chat_id, user_id, participant_type, link, inviter_user_id, invited_at, kicked_at, left_at, is_bot, state, date2 from chat_participants where chat_id = :chat_id and user_id in (:idList)
func (m *defaultChatParticipantsModel) SelectListByParticipantIdListWithCB(ctx context.Context, chatId int64, idList []int64, cb func(sz, i int, v *ChatParticipants)) (rList []ChatParticipants, err error) {
	var (
		query  = fmt.Sprintf("select id, chat_id, user_id, participant_type, link, inviter_user_id, invited_at, kicked_at, left_at, is_bot, state, date2 from chat_participants where chat_id = ? and user_id in (%s)", sqlx.InInt64List(idList))
		values []ChatParticipants
	)
	if len(idList) == 0 {
		rList = []ChatParticipants{}
		return
	}

	err = m.db.QueryRowsPartial(ctx, &values, query, chatId)

	if err != nil {
		err = fmt.Errorf("chat_participants.SelectListByParticipantIdListWithCB: %w", err)
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

// Update
// update chat_participants set participant_type = :participant_type, inviter_user_id = :inviter_user_id, invited_at = :invited_at, state = 0, kicked_at = 0, left_at = 0, is_bot = :is_bot where id = :id
func (m *defaultChatParticipantsModel) Update(ctx context.Context, participantType int32, inviterUserId int64, invitedAt int64, isBot bool, id int64) (rowsAffected int64, err error) {

	var (
		query   = "update chat_participants set participant_type = ?, inviter_user_id = ?, invited_at = ?, state = 0, kicked_at = 0, left_at = 0, is_bot = ? where id = ?"
		rResult sql.Result
	)

	rResult, err = m.db.Exec(ctx, query, participantType, inviterUserId, invitedAt, isBot, id)

	if err != nil {
		err = fmt.Errorf("chat_participants.Update exec: %w", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		err = fmt.Errorf("chat_participants.Update rows affected: %w", err)
	}

	return
}

// UpdateTx
// update chat_participants set participant_type = :participant_type, inviter_user_id = :inviter_user_id, invited_at = :invited_at, state = 0, kicked_at = 0, left_at = 0, is_bot = :is_bot where id = :id
func (m *defaultChatParticipantsModel) UpdateTx(tx *sqlx.Tx, participantType int32, inviterUserId int64, invitedAt int64, isBot bool, id int64) (rowsAffected int64, err error) {
	var (
		query   = "update chat_participants set participant_type = ?, inviter_user_id = ?, invited_at = ?, state = 0, kicked_at = 0, left_at = 0, is_bot = ? where id = ?"
		rResult sql.Result
	)
	rResult, err = tx.Exec(query, participantType, inviterUserId, invitedAt, isBot, id)

	if err != nil {
		err = fmt.Errorf("chat_participants.UpdateTx exec: %w", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		err = fmt.Errorf("chat_participants.UpdateTx rows affected: %w", err)
	}

	return
}

// UpdateKicked
// update chat_participants set kicked_at = :kicked_at, left_at = 0, state = 2 where id = :id
func (m *defaultChatParticipantsModel) UpdateKicked(ctx context.Context, kickedAt int64, id int64) (rowsAffected int64, err error) {

	var (
		query   = "update chat_participants set kicked_at = ?, left_at = 0, state = 2 where id = ?"
		rResult sql.Result
	)

	rResult, err = m.db.Exec(ctx, query, kickedAt, id)

	if err != nil {
		err = fmt.Errorf("chat_participants.UpdateKicked exec: %w", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		err = fmt.Errorf("chat_participants.UpdateKicked rows affected: %w", err)
	}

	return
}

// UpdateKickedTx
// update chat_participants set kicked_at = :kicked_at, left_at = 0, state = 2 where id = :id
func (m *defaultChatParticipantsModel) UpdateKickedTx(tx *sqlx.Tx, kickedAt int64, id int64) (rowsAffected int64, err error) {
	var (
		query   = "update chat_participants set kicked_at = ?, left_at = 0, state = 2 where id = ?"
		rResult sql.Result
	)
	rResult, err = tx.Exec(query, kickedAt, id)

	if err != nil {
		err = fmt.Errorf("chat_participants.UpdateKickedTx exec: %w", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		err = fmt.Errorf("chat_participants.UpdateKickedTx rows affected: %w", err)
	}

	return
}

// UpdateLeft
// update chat_participants set kicked_at = 0, left_at = :left_at, state = 1 where id = :id
func (m *defaultChatParticipantsModel) UpdateLeft(ctx context.Context, leftAt int64, id int64) (rowsAffected int64, err error) {

	var (
		query   = "update chat_participants set kicked_at = 0, left_at = ?, state = 1 where id = ?"
		rResult sql.Result
	)

	rResult, err = m.db.Exec(ctx, query, leftAt, id)

	if err != nil {
		err = fmt.Errorf("chat_participants.UpdateLeft exec: %w", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		err = fmt.Errorf("chat_participants.UpdateLeft rows affected: %w", err)
	}

	return
}

// UpdateLeftTx
// update chat_participants set kicked_at = 0, left_at = :left_at, state = 1 where id = :id
func (m *defaultChatParticipantsModel) UpdateLeftTx(tx *sqlx.Tx, leftAt int64, id int64) (rowsAffected int64, err error) {
	var (
		query   = "update chat_participants set kicked_at = 0, left_at = ?, state = 1 where id = ?"
		rResult sql.Result
	)
	rResult, err = tx.Exec(query, leftAt, id)

	if err != nil {
		err = fmt.Errorf("chat_participants.UpdateLeftTx exec: %w", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		err = fmt.Errorf("chat_participants.UpdateLeftTx rows affected: %w", err)
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
		err = fmt.Errorf("chat_participants.UpdateParticipantType exec: %w", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		err = fmt.Errorf("chat_participants.UpdateParticipantType rows affected: %w", err)
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
		err = fmt.Errorf("chat_participants.UpdateParticipantTypeTx exec: %w", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		err = fmt.Errorf("chat_participants.UpdateParticipantTypeTx rows affected: %w", err)
	}

	return
}

// SelectUsersChatIdList
// select chat_participants.chat_id as chat_id, chat_participants.user_id as user_id from chat_participants, chats where chat_participants.state = 0 and chat_participants.user_id in (:idList) and chats.id = chat_participants.chat_id and chats.deactivated = 0
func (m *defaultChatParticipantsModel) SelectUsersChatIdList(ctx context.Context, idList []int64) (rList []ChatParticipants, err error) {
	var (
		query  = fmt.Sprintf("select chat_participants.chat_id as chat_id, chat_participants.user_id as user_id from chat_participants, chats where chat_participants.state = 0 and chat_participants.user_id in (%s) and chats.id = chat_participants.chat_id and chats.deactivated = 0", sqlx.InInt64List(idList))
		values []ChatParticipants
	)
	if len(idList) == 0 {
		rList = []ChatParticipants{}
		return
	}

	err = m.db.QueryRowsPartial(ctx, &values, query)

	if err != nil {
		err = fmt.Errorf("chat_participants.SelectUsersChatIdList: %w", err)
		return
	}

	rList = values

	return
}

// SelectUsersChatIdListWithCB
// select chat_participants.chat_id as chat_id, chat_participants.user_id as user_id from chat_participants, chats where chat_participants.state = 0 and chat_participants.user_id in (:idList) and chats.id = chat_participants.chat_id and chats.deactivated = 0
func (m *defaultChatParticipantsModel) SelectUsersChatIdListWithCB(ctx context.Context, idList []int64, cb func(sz, i int, v *ChatParticipants)) (rList []ChatParticipants, err error) {
	var (
		query  = fmt.Sprintf("select chat_participants.chat_id as chat_id, chat_participants.user_id as user_id from chat_participants, chats where chat_participants.state = 0 and chat_participants.user_id in (%s) and chats.id = chat_participants.chat_id and chats.deactivated = 0", sqlx.InInt64List(idList))
		values []ChatParticipants
	)
	if len(idList) == 0 {
		rList = []ChatParticipants{}
		return
	}

	err = m.db.QueryRowsPartial(ctx, &values, query)

	if err != nil {
		err = fmt.Errorf("chat_participants.SelectUsersChatIdListWithCB: %w", err)
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

// SelectMyAdminList
// select chat_id from chat_participants where user_id = :user_id and participant_type = 1 and state = 0
func (m *defaultChatParticipantsModel) SelectMyAdminList(ctx context.Context, userId int64) (rList []int64, err error) {
	var query = "select chat_id from chat_participants where user_id = ? and participant_type = 1 and state = 0"
	err = m.db.QueryRowsPartial(ctx, &rList, query, userId)

	if err != nil {
		err = fmt.Errorf("chat_participants.SelectMyAdminList: %w", err)
	}

	return
}

// SelectMyAdminListWithCB
// select chat_id from chat_participants where user_id = :user_id and participant_type = 1 and state = 0
func (m *defaultChatParticipantsModel) SelectMyAdminListWithCB(ctx context.Context, userId int64, cb func(sz, i int, v int64)) (rList []int64, err error) {
	var query = "select chat_id from chat_participants where user_id = ? and participant_type = 1 and state = 0"
	err = m.db.QueryRowsPartial(ctx, &rList, query, userId)

	if err != nil {
		err = fmt.Errorf("chat_participants.SelectMyAdminListWithCB: %w", err)
	}

	if cb != nil {
		sz := len(rList)
		for i := 0; i < sz; i++ {
			cb(sz, i, rList[i])
		}
	}

	return
}

// SelectMyAllList
// select chat_id from chat_participants where user_id = :user_id and state = 0
func (m *defaultChatParticipantsModel) SelectMyAllList(ctx context.Context, userId int64) (rList []int64, err error) {
	var query = "select chat_id from chat_participants where user_id = ? and state = 0"
	err = m.db.QueryRowsPartial(ctx, &rList, query, userId)

	if err != nil {
		err = fmt.Errorf("chat_participants.SelectMyAllList: %w", err)
	}

	return
}

// SelectMyAllListWithCB
// select chat_id from chat_participants where user_id = :user_id and state = 0
func (m *defaultChatParticipantsModel) SelectMyAllListWithCB(ctx context.Context, userId int64, cb func(sz, i int, v int64)) (rList []int64, err error) {
	var query = "select chat_id from chat_participants where user_id = ? and state = 0"
	err = m.db.QueryRowsPartial(ctx, &rList, query, userId)

	if err != nil {
		err = fmt.Errorf("chat_participants.SelectMyAllListWithCB: %w", err)
	}

	if cb != nil {
		sz := len(rList)
		for i := 0; i < sz; i++ {
			cb(sz, i, rList[i])
		}
	}

	return
}

// UpdateStateByChatId
// update chat_participants set state = :state where chat_id = :chat_id and state = 0
func (m *defaultChatParticipantsModel) UpdateStateByChatId(ctx context.Context, state int32, chatId int64) (rowsAffected int64, err error) {

	var (
		query   = "update chat_participants set state = ? where chat_id = ? and state = 0"
		rResult sql.Result
	)

	rResult, err = m.db.Exec(ctx, query, state, chatId)

	if err != nil {
		err = fmt.Errorf("chat_participants.UpdateStateByChatId exec: %w", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		err = fmt.Errorf("chat_participants.UpdateStateByChatId rows affected: %w", err)
	}

	return
}

// UpdateStateByChatIdTx
// update chat_participants set state = :state where chat_id = :chat_id and state = 0
func (m *defaultChatParticipantsModel) UpdateStateByChatIdTx(tx *sqlx.Tx, state int32, chatId int64) (rowsAffected int64, err error) {
	var (
		query   = "update chat_participants set state = ? where chat_id = ? and state = 0"
		rResult sql.Result
	)
	rResult, err = tx.Exec(query, state, chatId)

	if err != nil {
		err = fmt.Errorf("chat_participants.UpdateStateByChatIdTx exec: %w", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		err = fmt.Errorf("chat_participants.UpdateStateByChatIdTx rows affected: %w", err)
	}

	return
}

// UpdateLink
// update chat_participants set link = :link where chat_id = :chat_id and user_id = :user_id
func (m *defaultChatParticipantsModel) UpdateLink(ctx context.Context, link string, chatId int64, userId int64) (rowsAffected int64, err error) {

	var (
		query   = "update chat_participants set link = ? where chat_id = ? and user_id = ?"
		rResult sql.Result
	)

	rResult, err = m.db.Exec(ctx, query, link, chatId, userId)

	if err != nil {
		err = fmt.Errorf("chat_participants.UpdateLink exec: %w", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		err = fmt.Errorf("chat_participants.UpdateLink rows affected: %w", err)
	}

	return
}

// UpdateLinkTx
// update chat_participants set link = :link where chat_id = :chat_id and user_id = :user_id
func (m *defaultChatParticipantsModel) UpdateLinkTx(tx *sqlx.Tx, link string, chatId int64, userId int64) (rowsAffected int64, err error) {
	var (
		query   = "update chat_participants set link = ? where chat_id = ? and user_id = ?"
		rResult sql.Result
	)
	rResult, err = tx.Exec(query, link, chatId, userId)

	if err != nil {
		err = fmt.Errorf("chat_participants.UpdateLinkTx exec: %w", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		err = fmt.Errorf("chat_participants.UpdateLinkTx rows affected: %w", err)
	}

	return
}

// UpdateLinkUsage
// update chat_participants set usage2 = :usage2 where chat_id = :chat_id and user_id = :user_id
func (m *defaultChatParticipantsModel) UpdateLinkUsage(ctx context.Context, usage2 int32, chatId int64, userId int64) (rowsAffected int64, err error) {

	var (
		query   = "update chat_participants set usage2 = ? where chat_id = ? and user_id = ?"
		rResult sql.Result
	)

	rResult, err = m.db.Exec(ctx, query, usage2, chatId, userId)

	if err != nil {
		err = fmt.Errorf("chat_participants.UpdateLinkUsage exec: %w", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		err = fmt.Errorf("chat_participants.UpdateLinkUsage rows affected: %w", err)
	}

	return
}

// UpdateLinkUsageTx
// update chat_participants set usage2 = :usage2 where chat_id = :chat_id and user_id = :user_id
func (m *defaultChatParticipantsModel) UpdateLinkUsageTx(tx *sqlx.Tx, usage2 int32, chatId int64, userId int64) (rowsAffected int64, err error) {
	var (
		query   = "update chat_participants set usage2 = ? where chat_id = ? and user_id = ?"
		rResult sql.Result
	)
	rResult, err = tx.Exec(query, usage2, chatId, userId)

	if err != nil {
		err = fmt.Errorf("chat_participants.UpdateLinkUsageTx exec: %w", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		err = fmt.Errorf("chat_participants.UpdateLinkUsageTx rows affected: %w", err)
	}

	return
}
