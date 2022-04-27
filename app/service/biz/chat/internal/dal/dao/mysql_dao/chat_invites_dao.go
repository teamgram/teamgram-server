/*
 * WARNING! All changes made in this file will be lost!
 *   Created from by 'dalgen'
 *
 * Copyright (c) 2022-present,  Teamgram Authors.
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package mysql_dao

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/teamgram/marmota/pkg/stores/sqlx"
	"github.com/teamgram/teamgram-server/app/service/biz/chat/internal/dal/dataobject"

	"github.com/zeromicro/go-zero/core/logx"
)

var _ *sql.Result

type ChatInvitesDAO struct {
	db *sqlx.DB
}

func NewChatInvitesDAO(db *sqlx.DB) *ChatInvitesDAO {
	return &ChatInvitesDAO{db}
}

// Insert
// insert into chat_invites(chat_id, admin_id, link, permanent, revoked, request_needed, start_date, expire_date, usage_limit, usage2, requested, title, date2) values (:chat_id, :admin_id, :link, :permanent, :revoked, :request_needed, :start_date, :expire_date, :usage_limit, :usage2, :requested, :title, :date2)
// TODO(@benqi): sqlmap
func (dao *ChatInvitesDAO) Insert(ctx context.Context, do *dataobject.ChatInvitesDO) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into chat_invites(chat_id, admin_id, link, permanent, revoked, request_needed, start_date, expire_date, usage_limit, usage2, requested, title, date2) values (:chat_id, :admin_id, :link, :permanent, :revoked, :request_needed, :start_date, :expire_date, :usage_limit, :usage2, :requested, :title, :date2)"
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
// insert into chat_invites(chat_id, admin_id, link, permanent, revoked, request_needed, start_date, expire_date, usage_limit, usage2, requested, title, date2) values (:chat_id, :admin_id, :link, :permanent, :revoked, :request_needed, :start_date, :expire_date, :usage_limit, :usage2, :requested, :title, :date2)
// TODO(@benqi): sqlmap
func (dao *ChatInvitesDAO) InsertTx(tx *sqlx.Tx, do *dataobject.ChatInvitesDO) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into chat_invites(chat_id, admin_id, link, permanent, revoked, request_needed, start_date, expire_date, usage_limit, usage2, requested, title, date2) values (:chat_id, :admin_id, :link, :permanent, :revoked, :request_needed, :start_date, :expire_date, :usage_limit, :usage2, :requested, :title, :date2)"
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

// SelectListByAdminId
// select id, chat_id, admin_id, link, permanent, revoked, request_needed, start_date, expire_date, usage_limit, usage2, requested, title, date2 from chat_invites where chat_id = :chat_id and admin_id = :admin_id
// TODO(@benqi): sqlmap
func (dao *ChatInvitesDAO) SelectListByAdminId(ctx context.Context, chat_id int64, admin_id int64) (rList []dataobject.ChatInvitesDO, err error) {
	var (
		query  = "select id, chat_id, admin_id, link, permanent, revoked, request_needed, start_date, expire_date, usage_limit, usage2, requested, title, date2 from chat_invites where chat_id = ? and admin_id = ?"
		values []dataobject.ChatInvitesDO
	)
	err = dao.db.QueryRowsPartial(ctx, &values, query, chat_id, admin_id)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectListByAdminId(_), error: %v", err)
		return
	}

	rList = values

	return
}

// SelectListByAdminIdWithCB
// select id, chat_id, admin_id, link, permanent, revoked, request_needed, start_date, expire_date, usage_limit, usage2, requested, title, date2 from chat_invites where chat_id = :chat_id and admin_id = :admin_id
// TODO(@benqi): sqlmap
func (dao *ChatInvitesDAO) SelectListByAdminIdWithCB(ctx context.Context, chat_id int64, admin_id int64, cb func(i int, v *dataobject.ChatInvitesDO)) (rList []dataobject.ChatInvitesDO, err error) {
	var (
		query  = "select id, chat_id, admin_id, link, permanent, revoked, request_needed, start_date, expire_date, usage_limit, usage2, requested, title, date2 from chat_invites where chat_id = ? and admin_id = ?"
		values []dataobject.ChatInvitesDO
	)
	err = dao.db.QueryRowsPartial(ctx, &values, query, chat_id, admin_id)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectListByAdminId(_), error: %v", err)
		return
	}

	rList = values

	if cb != nil {
		for i := 0; i < len(rList); i++ {
			cb(i, &rList[i])
		}
	}

	return
}

// SelectByLink
// select id, chat_id, admin_id, link, permanent, revoked, request_needed, start_date, expire_date, usage_limit, usage2, requested, title, date2 from chat_invites where link = :link
// TODO(@benqi): sqlmap
func (dao *ChatInvitesDAO) SelectByLink(ctx context.Context, link string) (rValue *dataobject.ChatInvitesDO, err error) {
	var (
		query = "select id, chat_id, admin_id, link, permanent, revoked, request_needed, start_date, expire_date, usage_limit, usage2, requested, title, date2 from chat_invites where link = ?"
		do    = &dataobject.ChatInvitesDO{}
	)
	err = dao.db.QueryRowPartial(ctx, do, query, link)

	if err != nil {
		if err != sqlx.ErrNotFound {
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

// SelectListByChatId
// select id, chat_id, admin_id, link, permanent, revoked, request_needed, start_date, expire_date, usage_limit, usage2, requested, title, date2 from chat_invites where chat_id = :chat_id
// TODO(@benqi): sqlmap
func (dao *ChatInvitesDAO) SelectListByChatId(ctx context.Context, chat_id int64) (rList []dataobject.ChatInvitesDO, err error) {
	var (
		query  = "select id, chat_id, admin_id, link, permanent, revoked, request_needed, start_date, expire_date, usage_limit, usage2, requested, title, date2 from chat_invites where chat_id = ?"
		values []dataobject.ChatInvitesDO
	)
	err = dao.db.QueryRowsPartial(ctx, &values, query, chat_id)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectListByChatId(_), error: %v", err)
		return
	}

	rList = values

	return
}

// SelectListByChatIdWithCB
// select id, chat_id, admin_id, link, permanent, revoked, request_needed, start_date, expire_date, usage_limit, usage2, requested, title, date2 from chat_invites where chat_id = :chat_id
// TODO(@benqi): sqlmap
func (dao *ChatInvitesDAO) SelectListByChatIdWithCB(ctx context.Context, chat_id int64, cb func(i int, v *dataobject.ChatInvitesDO)) (rList []dataobject.ChatInvitesDO, err error) {
	var (
		query  = "select id, chat_id, admin_id, link, permanent, revoked, request_needed, start_date, expire_date, usage_limit, usage2, requested, title, date2 from chat_invites where chat_id = ?"
		values []dataobject.ChatInvitesDO
	)
	err = dao.db.QueryRowsPartial(ctx, &values, query, chat_id)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectListByChatId(_), error: %v", err)
		return
	}

	rList = values

	if cb != nil {
		for i := 0; i < len(rList); i++ {
			cb(i, &rList[i])
		}
	}

	return
}

// Update
// update chat_invites set %s where chat_id = :chat_id and link = :link
// TODO(@benqi): sqlmap
func (dao *ChatInvitesDAO) Update(ctx context.Context, cMap map[string]interface{}, chat_id int64, link string) (rowsAffected int64, err error) {
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

	aValues = append(aValues, chat_id)
	aValues = append(aValues, link)

	rResult, err = dao.db.Exec(ctx, query, aValues...)

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
// TODO(@benqi): sqlmap
func (dao *ChatInvitesDAO) UpdateTx(tx *sqlx.Tx, cMap map[string]interface{}, chat_id int64, link string) (rowsAffected int64, err error) {
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

	aValues = append(aValues, chat_id)
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
// TODO(@benqi): sqlmap
func (dao *ChatInvitesDAO) DeleteByLink(ctx context.Context, chat_id int64, link string) (rowsAffected int64, err error) {
	var (
		query   = "delete from chat_invites where chat_id = ? and link = ?"
		rResult sql.Result
	)
	rResult, err = dao.db.Exec(ctx, query, chat_id, link)

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
// TODO(@benqi): sqlmap
func (dao *ChatInvitesDAO) DeleteByLinkTx(tx *sqlx.Tx, chat_id int64, link string) (rowsAffected int64, err error) {
	var (
		query   = "delete from chat_invites where chat_id = ? and link = ?"
		rResult sql.Result
	)
	rResult, err = tx.Exec(query, chat_id, link)

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
// TODO(@benqi): sqlmap
func (dao *ChatInvitesDAO) DeleteByRevoked(ctx context.Context, chat_id int64, admin_id int64) (rowsAffected int64, err error) {
	var (
		query   = "delete from chat_invites where chat_id = ? and admin_id = ? and revoked = 1"
		rResult sql.Result
	)
	rResult, err = dao.db.Exec(ctx, query, chat_id, admin_id)

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
// TODO(@benqi): sqlmap
func (dao *ChatInvitesDAO) DeleteByRevokedTx(tx *sqlx.Tx, chat_id int64, admin_id int64) (rowsAffected int64, err error) {
	var (
		query   = "delete from chat_invites where chat_id = ? and admin_id = ? and revoked = 1"
		rResult sql.Result
	)
	rResult, err = tx.Exec(query, chat_id, admin_id)

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
