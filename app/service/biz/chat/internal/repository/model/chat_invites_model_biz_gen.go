/*
 * WARNING! All changes made in this file will be lost!
 *   Created from by 'dalgen'
 *
 * Copyright (c) 2026 The Teamgram Authors.
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
	bizChatInvitesModel interface {
		Insert(ctx context.Context, data *ChatInvites) (lastInsertId, rowsAffected int64, err error)
		InsertTx(tx *sqlx.Tx, data *ChatInvites) (lastInsertId, rowsAffected int64, err error)

		SelectListByAdminId(ctx context.Context, chatId int64, adminId int64) ([]ChatInvites, error)
		SelectListByAdminIdWithCB(ctx context.Context, chatId int64, adminId int64, cb func(sz, i int, v *ChatInvites)) ([]ChatInvites, error)

		SelectByLink(ctx context.Context, link string) (*ChatInvites, error)

		SelectAll(ctx context.Context) ([]ChatInvites, error)
		SelectAllWithCB(ctx context.Context, cb func(sz, i int, v *ChatInvites)) ([]ChatInvites, error)

		SelectListByChatId(ctx context.Context, chatId int64) ([]ChatInvites, error)
		SelectListByChatIdWithCB(ctx context.Context, chatId int64, cb func(sz, i int, v *ChatInvites)) ([]ChatInvites, error)

		Update(ctx context.Context, cMap map[string]interface{}, chatId int64, link string) (rowsAffected int64, err error)
		UpdateTx(tx *sqlx.Tx, cMap map[string]interface{}, chatId int64, link string) (rowsAffected int64, err error)

		DeleteByLink(ctx context.Context, chatId int64, link string) (rowsAffected int64, err error)
		DeleteByLinkTx(tx *sqlx.Tx, chatId int64, link string) (rowsAffected int64, err error)

		DeleteByRevoked(ctx context.Context, chatId int64, adminId int64) (rowsAffected int64, err error)
		DeleteByRevokedTx(tx *sqlx.Tx, chatId int64, adminId int64) (rowsAffected int64, err error)
	}
)

// Insert
// insert into chat_invites(chat_id, admin_id, link, permanent, revoked, request_needed, start_date, expire_date, usage_limit, usage2, requested, title, date2) values (:chat_id, :admin_id, :link, :permanent, :revoked, :request_needed, :start_date, :expire_date, :usage_limit, :usage2, :requested, :title, :date2)
func (m *defaultChatInvitesModel) Insert(ctx context.Context, data *ChatInvites) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into chat_invites(chat_id, admin_id, link, permanent, revoked, request_needed, start_date, expire_date, usage_limit, usage2, requested, title, date2) values (:chat_id, :admin_id, :link, :permanent, :revoked, :request_needed, :start_date, :expire_date, :usage_limit, :usage2, :requested, :title, :date2)"
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
// insert into chat_invites(chat_id, admin_id, link, permanent, revoked, request_needed, start_date, expire_date, usage_limit, usage2, requested, title, date2) values (:chat_id, :admin_id, :link, :permanent, :revoked, :request_needed, :start_date, :expire_date, :usage_limit, :usage2, :requested, :title, :date2)
func (m *defaultChatInvitesModel) InsertTx(tx *sqlx.Tx, data *ChatInvites) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into chat_invites(chat_id, admin_id, link, permanent, revoked, request_needed, start_date, expire_date, usage_limit, usage2, requested, title, date2) values (:chat_id, :admin_id, :link, :permanent, :revoked, :request_needed, :start_date, :expire_date, :usage_limit, :usage2, :requested, :title, :date2)"
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

// SelectListByAdminId
// select id, chat_id, admin_id, link, permanent, revoked, request_needed, start_date, expire_date, usage_limit, usage2, requested, title, date2 from chat_invites where chat_id = :chat_id and admin_id = :admin_id
func (m *defaultChatInvitesModel) SelectListByAdminId(ctx context.Context, chatId int64, adminId int64) (rList []ChatInvites, err error) {
	var (
		query  = "select id, chat_id, admin_id, link, permanent, revoked, request_needed, start_date, expire_date, usage_limit, usage2, requested, title, date2 from chat_invites where chat_id = ? and admin_id = ?"
		values []ChatInvites
	)
	err = m.db.QueryRowsPartial(ctx, &values, query, chatId, adminId)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectListByAdminId(_), error: %v", err)
		return
	}

	rList = values

	return
}

// SelectListByAdminIdWithCB
// select id, chat_id, admin_id, link, permanent, revoked, request_needed, start_date, expire_date, usage_limit, usage2, requested, title, date2 from chat_invites where chat_id = :chat_id and admin_id = :admin_id
func (m *defaultChatInvitesModel) SelectListByAdminIdWithCB(ctx context.Context, chatId int64, adminId int64, cb func(sz, i int, v *ChatInvites)) (rList []ChatInvites, err error) {
	var (
		query  = "select id, chat_id, admin_id, link, permanent, revoked, request_needed, start_date, expire_date, usage_limit, usage2, requested, title, date2 from chat_invites where chat_id = ? and admin_id = ?"
		values []ChatInvites
	)
	err = m.db.QueryRowsPartial(ctx, &values, query, chatId, adminId)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectListByAdminId(_), error: %v", err)
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

// SelectByLink
// select id, chat_id, admin_id, link, permanent, revoked, request_needed, start_date, expire_date, usage_limit, usage2, requested, title, date2 from chat_invites where link = :link
func (m *defaultChatInvitesModel) SelectByLink(ctx context.Context, link string) (rValue *ChatInvites, err error) {
	var (
		query = "select id, chat_id, admin_id, link, permanent, revoked, request_needed, start_date, expire_date, usage_limit, usage2, requested, title, date2 from chat_invites where link = ?"
		do    = &ChatInvites{}
	)
	err = m.db.QueryRowPartial(ctx, do, query, link)

	if err != nil {
		if !errors.Is(err, sqlx.ErrNotFound) {
			logx.WithContext(ctx).Errorf("queryx in SelectByLink(_), error: %v", err)
			return
		} else {
			err = nil
		}
	} else {
		rValue = do
	}

	return
}

// SelectAll
// select id, chat_id, admin_id, link, permanent, revoked, request_needed, start_date, expire_date, usage_limit, usage2, requested, title, date2 from chat_invites
func (m *defaultChatInvitesModel) SelectAll(ctx context.Context) (rList []ChatInvites, err error) {
	var (
		query  = "select id, chat_id, admin_id, link, permanent, revoked, request_needed, start_date, expire_date, usage_limit, usage2, requested, title, date2 from chat_invites"
		values []ChatInvites
	)
	err = m.db.QueryRowsPartial(ctx, &values, query)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectAll(_), error: %v", err)
		return
	}

	rList = values

	return
}

// SelectAllWithCB
// select id, chat_id, admin_id, link, permanent, revoked, request_needed, start_date, expire_date, usage_limit, usage2, requested, title, date2 from chat_invites
func (m *defaultChatInvitesModel) SelectAllWithCB(ctx context.Context, cb func(sz, i int, v *ChatInvites)) (rList []ChatInvites, err error) {
	var (
		query  = "select id, chat_id, admin_id, link, permanent, revoked, request_needed, start_date, expire_date, usage_limit, usage2, requested, title, date2 from chat_invites"
		values []ChatInvites
	)
	err = m.db.QueryRowsPartial(ctx, &values, query)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectAll(_), error: %v", err)
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

// SelectListByChatId
// select id, chat_id, admin_id, link, permanent, revoked, request_needed, start_date, expire_date, usage_limit, usage2, requested, title, date2 from chat_invites where chat_id = :chat_id
func (m *defaultChatInvitesModel) SelectListByChatId(ctx context.Context, chatId int64) (rList []ChatInvites, err error) {
	var (
		query  = "select id, chat_id, admin_id, link, permanent, revoked, request_needed, start_date, expire_date, usage_limit, usage2, requested, title, date2 from chat_invites where chat_id = ?"
		values []ChatInvites
	)
	err = m.db.QueryRowsPartial(ctx, &values, query, chatId)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectListByChatId(_), error: %v", err)
		return
	}

	rList = values

	return
}

// SelectListByChatIdWithCB
// select id, chat_id, admin_id, link, permanent, revoked, request_needed, start_date, expire_date, usage_limit, usage2, requested, title, date2 from chat_invites where chat_id = :chat_id
func (m *defaultChatInvitesModel) SelectListByChatIdWithCB(ctx context.Context, chatId int64, cb func(sz, i int, v *ChatInvites)) (rList []ChatInvites, err error) {
	var (
		query  = "select id, chat_id, admin_id, link, permanent, revoked, request_needed, start_date, expire_date, usage_limit, usage2, requested, title, date2 from chat_invites where chat_id = ?"
		values []ChatInvites
	)
	err = m.db.QueryRowsPartial(ctx, &values, query, chatId)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectListByChatId(_), error: %v", err)
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
// update chat_invites set %s where chat_id = :chat_id and link = :link
func (m *defaultChatInvitesModel) Update(ctx context.Context, cMap map[string]interface{}, chatId int64, link string) (rowsAffected int64, err error) {
	names := make([]string, 0, len(cMap))
	aValues := make([]interface{}, 0, len(cMap))
	for k, v := range cMap {
		names = append(names, k+" = ?")
		aValues = append(aValues, v)
	}

	var (
		query   = fmt.Sprintf("update chat_invites set %s where chat_id = ? and link = ?", strings.Join(names, ", "))
		rResult sql.Result
	)

	aValues = append(aValues, chatId)
	aValues = append(aValues, link)

	rResult, err = m.db.Exec(ctx, query, aValues...)

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
// update chat_invites set %s where chat_id = :chat_id and link = :link
func (m *defaultChatInvitesModel) UpdateTx(tx *sqlx.Tx, cMap map[string]interface{}, chatId int64, link string) (rowsAffected int64, err error) {
	names := make([]string, 0, len(cMap))
	aValues := make([]interface{}, 0, len(cMap))
	for k, v := range cMap {
		names = append(names, k+" = ?")
		aValues = append(aValues, v)
	}

	var (
		query   = fmt.Sprintf("update chat_invites set %s where chat_id = ? and link = ?", strings.Join(names, ", "))
		rResult sql.Result
	)

	aValues = append(aValues, chatId)
	aValues = append(aValues, link)

	rResult, err = tx.Exec(query, aValues...)

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

// DeleteByLink
// delete from chat_invites where chat_id = :chat_id and link = :link
func (m *defaultChatInvitesModel) DeleteByLink(ctx context.Context, chatId int64, link string) (rowsAffected int64, err error) {
	var (
		query   = "delete from chat_invites where chat_id = ? and link = ?"
		rResult sql.Result
	)
	rResult, err = m.db.Exec(ctx, query, chatId, link)

	if err != nil {
		logx.WithContext(ctx).Errorf("exec in DeleteByLink(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		logx.WithContext(ctx).Errorf("rowsAffected in DeleteByLink(_), error: %v", err)
	}

	return
}

// DeleteByLinkTx
// delete from chat_invites where chat_id = :chat_id and link = :link
func (m *defaultChatInvitesModel) DeleteByLinkTx(tx *sqlx.Tx, chatId int64, link string) (rowsAffected int64, err error) {
	var (
		query   = "delete from chat_invites where chat_id = ? and link = ?"
		rResult sql.Result
	)
	rResult, err = tx.Exec(query, chatId, link)

	if err != nil {
		logx.WithContext(tx.Context()).Errorf("exec in DeleteByLink(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		logx.WithContext(tx.Context()).Errorf("rowsAffected in DeleteByLink(_), error: %v", err)
	}

	return
}

// DeleteByRevoked
// delete from chat_invites where chat_id = :chat_id and admin_id = :admin_id and revoked = 1
func (m *defaultChatInvitesModel) DeleteByRevoked(ctx context.Context, chatId int64, adminId int64) (rowsAffected int64, err error) {
	var (
		query   = "delete from chat_invites where chat_id = ? and admin_id = ? and revoked = 1"
		rResult sql.Result
	)
	rResult, err = m.db.Exec(ctx, query, chatId, adminId)

	if err != nil {
		logx.WithContext(ctx).Errorf("exec in DeleteByRevoked(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		logx.WithContext(ctx).Errorf("rowsAffected in DeleteByRevoked(_), error: %v", err)
	}

	return
}

// DeleteByRevokedTx
// delete from chat_invites where chat_id = :chat_id and admin_id = :admin_id and revoked = 1
func (m *defaultChatInvitesModel) DeleteByRevokedTx(tx *sqlx.Tx, chatId int64, adminId int64) (rowsAffected int64, err error) {
	var (
		query   = "delete from chat_invites where chat_id = ? and admin_id = ? and revoked = 1"
		rResult sql.Result
	)
	rResult, err = tx.Exec(query, chatId, adminId)

	if err != nil {
		logx.WithContext(tx.Context()).Errorf("exec in DeleteByRevoked(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		logx.WithContext(tx.Context()).Errorf("rowsAffected in DeleteByRevoked(_), error: %v", err)
	}

	return
}
