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
	bizChatInviteParticipantsModel interface {
		Insert(ctx context.Context, data *ChatInviteParticipants) (lastInsertId, rowsAffected int64, err error)
		InsertTx(tx *sqlx.Tx, data *ChatInviteParticipants) (lastInsertId, rowsAffected int64, err error)

		SelectListByLink(ctx context.Context, link string, b int32) ([]ChatInviteParticipants, error)
		SelectListByLinkWithCB(ctx context.Context, link string, b int32, cb func(sz, i int, v *ChatInviteParticipants)) ([]ChatInviteParticipants, error)

		Delete(ctx context.Context, chatId int64, userId int64) (rowsAffected int64, err error)
		DeleteTx(tx *sqlx.Tx, chatId int64, userId int64) (rowsAffected int64, err error)

		SelectRecentRequestedList(ctx context.Context, chatId int64) ([]ChatInviteParticipants, error)
		SelectRecentRequestedListWithCB(ctx context.Context, chatId int64, cb func(sz, i int, v *ChatInviteParticipants)) ([]ChatInviteParticipants, error)

		UpdateChatId(ctx context.Context, chatId int64, link string) (rowsAffected int64, err error)
		UpdateChatIdTx(tx *sqlx.Tx, chatId int64, link string) (rowsAffected int64, err error)

		UpdateApprovedBy(ctx context.Context, approvedBy int64, chatId int64, userId int64) (rowsAffected int64, err error)
		UpdateApprovedByTx(tx *sqlx.Tx, approvedBy int64, chatId int64, userId int64) (rowsAffected int64, err error)
	}
)

// Insert
// insert into chat_invite_participants(chat_id, link, user_id, requested, approved_by, date2) values (:chat_id, :link, :user_id, :requested, :approved_by, :date2)
func (m *defaultChatInviteParticipantsModel) Insert(ctx context.Context, data *ChatInviteParticipants) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into chat_invite_participants(chat_id, link, user_id, requested, approved_by, date2) values (:chat_id, :link, :user_id, :requested, :approved_by, :date2)"
		r     sql.Result
	)

	r, err = m.db.NamedExec(ctx, query, data)
	if err != nil {
		err = fmt.Errorf("chat_invite_participants.Insert named exec: %w", err)
		return
	}

	lastInsertId, err = r.LastInsertId()
	if err != nil {
		err = fmt.Errorf("chat_invite_participants.Insert last insert id: %w", err)
		return
	}
	rowsAffected, err = r.RowsAffected()
	if err != nil {
		err = fmt.Errorf("chat_invite_participants.Insert rows affected: %w", err)
	}

	return

}

// InsertTx
// insert into chat_invite_participants(chat_id, link, user_id, requested, approved_by, date2) values (:chat_id, :link, :user_id, :requested, :approved_by, :date2)
func (m *defaultChatInviteParticipantsModel) InsertTx(tx *sqlx.Tx, data *ChatInviteParticipants) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into chat_invite_participants(chat_id, link, user_id, requested, approved_by, date2) values (:chat_id, :link, :user_id, :requested, :approved_by, :date2)"
		r     sql.Result
	)

	r, err = tx.NamedExec(query, data)
	if err != nil {
		err = fmt.Errorf("chat_invite_participants.InsertTx named exec: %w", err)
		return
	}

	lastInsertId, err = r.LastInsertId()
	if err != nil {
		err = fmt.Errorf("chat_invite_participants.InsertTx last insert id: %w", err)
		return
	}
	rowsAffected, err = r.RowsAffected()
	if err != nil {
		err = fmt.Errorf("chat_invite_participants.InsertTx rows affected: %w", err)
	}

	return
}

// SelectListByLink
// select id, chat_id, link, user_id, requested, approved_by, date2 from chat_invite_participants where link = :link and requested = :b
func (m *defaultChatInviteParticipantsModel) SelectListByLink(ctx context.Context, link string, b int32) (rList []ChatInviteParticipants, err error) {
	var (
		query  = "select id, chat_id, link, user_id, requested, approved_by, date2 from chat_invite_participants where link = ? and requested = ?"
		values []ChatInviteParticipants
	)
	err = m.db.QueryRowsPartial(ctx, &values, query, link, b)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			rList = []ChatInviteParticipants{}
			err = nil
			return
		}
		err = fmt.Errorf("chat_invite_participants.SelectListByLink: %w", err)
		return
	}

	rList = values

	return
}

// SelectListByLinkWithCB
// select id, chat_id, link, user_id, requested, approved_by, date2 from chat_invite_participants where link = :link and requested = :b
func (m *defaultChatInviteParticipantsModel) SelectListByLinkWithCB(ctx context.Context, link string, b int32, cb func(sz, i int, v *ChatInviteParticipants)) (rList []ChatInviteParticipants, err error) {
	var (
		query  = "select id, chat_id, link, user_id, requested, approved_by, date2 from chat_invite_participants where link = ? and requested = ?"
		values []ChatInviteParticipants
	)
	err = m.db.QueryRowsPartial(ctx, &values, query, link, b)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			rList = []ChatInviteParticipants{}
			err = nil
			return
		}
		err = fmt.Errorf("chat_invite_participants.SelectListByLinkWithCB: %w", err)
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

// Delete
// delete from chat_invite_participants where chat_id = :chat_id and user_id = :user_id
func (m *defaultChatInviteParticipantsModel) Delete(ctx context.Context, chatId int64, userId int64) (rowsAffected int64, err error) {

	var (
		query   = "delete from chat_invite_participants where chat_id = ? and user_id = ?"
		rResult sql.Result
	)
	rResult, err = m.db.Exec(ctx, query, chatId, userId)

	if err != nil {
		err = fmt.Errorf("chat_invite_participants.Delete exec: %w", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		err = fmt.Errorf("chat_invite_participants.Delete rows affected: %w", err)
		return
	}

	return
}

// DeleteTx
// delete from chat_invite_participants where chat_id = :chat_id and user_id = :user_id
func (m *defaultChatInviteParticipantsModel) DeleteTx(tx *sqlx.Tx, chatId int64, userId int64) (rowsAffected int64, err error) {
	var (
		query   = "delete from chat_invite_participants where chat_id = ? and user_id = ?"
		rResult sql.Result
	)
	rResult, err = tx.Exec(query, chatId, userId)

	if err != nil {
		err = fmt.Errorf("chat_invite_participants.DeleteTx exec: %w", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		err = fmt.Errorf("chat_invite_participants.DeleteTx rows affected: %w", err)
		return
	}

	return
}

// SelectRecentRequestedList
// select id, chat_id, link, user_id, requested, approved_by, date2 from chat_invite_participants where chat_id = :chat_id and requested = 1
func (m *defaultChatInviteParticipantsModel) SelectRecentRequestedList(ctx context.Context, chatId int64) (rList []ChatInviteParticipants, err error) {
	var (
		query  = "select id, chat_id, link, user_id, requested, approved_by, date2 from chat_invite_participants where chat_id = ? and requested = 1"
		values []ChatInviteParticipants
	)
	err = m.db.QueryRowsPartial(ctx, &values, query, chatId)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			rList = []ChatInviteParticipants{}
			err = nil
			return
		}
		err = fmt.Errorf("chat_invite_participants.SelectRecentRequestedList: %w", err)
		return
	}

	rList = values

	return
}

// SelectRecentRequestedListWithCB
// select id, chat_id, link, user_id, requested, approved_by, date2 from chat_invite_participants where chat_id = :chat_id and requested = 1
func (m *defaultChatInviteParticipantsModel) SelectRecentRequestedListWithCB(ctx context.Context, chatId int64, cb func(sz, i int, v *ChatInviteParticipants)) (rList []ChatInviteParticipants, err error) {
	var (
		query  = "select id, chat_id, link, user_id, requested, approved_by, date2 from chat_invite_participants where chat_id = ? and requested = 1"
		values []ChatInviteParticipants
	)
	err = m.db.QueryRowsPartial(ctx, &values, query, chatId)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			rList = []ChatInviteParticipants{}
			err = nil
			return
		}
		err = fmt.Errorf("chat_invite_participants.SelectRecentRequestedListWithCB: %w", err)
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

// UpdateChatId
// update chat_invite_participants set chat_id = :chat_id where link = :link
func (m *defaultChatInviteParticipantsModel) UpdateChatId(ctx context.Context, chatId int64, link string) (rowsAffected int64, err error) {

	var (
		query   = "update chat_invite_participants set chat_id = ? where link = ?"
		rResult sql.Result
	)

	rResult, err = m.db.Exec(ctx, query, chatId, link)

	if err != nil {
		err = fmt.Errorf("chat_invite_participants.UpdateChatId exec: %w", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		err = fmt.Errorf("chat_invite_participants.UpdateChatId rows affected: %w", err)
		return
	}

	return
}

// UpdateChatIdTx
// update chat_invite_participants set chat_id = :chat_id where link = :link
func (m *defaultChatInviteParticipantsModel) UpdateChatIdTx(tx *sqlx.Tx, chatId int64, link string) (rowsAffected int64, err error) {
	var (
		query   = "update chat_invite_participants set chat_id = ? where link = ?"
		rResult sql.Result
	)
	rResult, err = tx.Exec(query, chatId, link)

	if err != nil {
		err = fmt.Errorf("chat_invite_participants.UpdateChatIdTx exec: %w", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		err = fmt.Errorf("chat_invite_participants.UpdateChatIdTx rows affected: %w", err)
		return
	}

	return
}

// UpdateApprovedBy
// update chat_invite_participants set requested = 0, approved_by = :approved_by where chat_id = :chat_id and user_id = :user_id
func (m *defaultChatInviteParticipantsModel) UpdateApprovedBy(ctx context.Context, approvedBy int64, chatId int64, userId int64) (rowsAffected int64, err error) {

	var (
		query   = "update chat_invite_participants set requested = 0, approved_by = ? where chat_id = ? and user_id = ?"
		rResult sql.Result
	)

	rResult, err = m.db.Exec(ctx, query, approvedBy, chatId, userId)

	if err != nil {
		err = fmt.Errorf("chat_invite_participants.UpdateApprovedBy exec: %w", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		err = fmt.Errorf("chat_invite_participants.UpdateApprovedBy rows affected: %w", err)
		return
	}

	return
}

// UpdateApprovedByTx
// update chat_invite_participants set requested = 0, approved_by = :approved_by where chat_id = :chat_id and user_id = :user_id
func (m *defaultChatInviteParticipantsModel) UpdateApprovedByTx(tx *sqlx.Tx, approvedBy int64, chatId int64, userId int64) (rowsAffected int64, err error) {
	var (
		query   = "update chat_invite_participants set requested = 0, approved_by = ? where chat_id = ? and user_id = ?"
		rResult sql.Result
	)
	rResult, err = tx.Exec(query, approvedBy, chatId, userId)

	if err != nil {
		err = fmt.Errorf("chat_invite_participants.UpdateApprovedByTx exec: %w", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		err = fmt.Errorf("chat_invite_participants.UpdateApprovedByTx rows affected: %w", err)
		return
	}

	return
}
