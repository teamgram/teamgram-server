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

	"github.com/teamgram/marmota/pkg/stores/sqlx"
	"github.com/teamgram/teamgram-server/app/service/biz/chat/internal/dal/dataobject"

	"github.com/zeromicro/go-zero/core/logx"
)

var _ *sql.Result

type ChatInviteParticipantsDAO struct {
	db *sqlx.DB
}

func NewChatInviteParticipantsDAO(db *sqlx.DB) *ChatInviteParticipantsDAO {
	return &ChatInviteParticipantsDAO{db}
}

// Insert
// insert into chat_invite_participants(chat_id, link, user_id, date2) values (:chat_id, :link, :user_id, :date2)
// TODO(@benqi): sqlmap
func (dao *ChatInviteParticipantsDAO) Insert(ctx context.Context, do *dataobject.ChatInviteParticipantsDO) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into chat_invite_participants(chat_id, link, user_id, date2) values (:chat_id, :link, :user_id, :date2)"
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
// insert into chat_invite_participants(chat_id, link, user_id, date2) values (:chat_id, :link, :user_id, :date2)
// TODO(@benqi): sqlmap
func (dao *ChatInviteParticipantsDAO) InsertTx(tx *sqlx.Tx, do *dataobject.ChatInviteParticipantsDO) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into chat_invite_participants(chat_id, link, user_id, date2) values (:chat_id, :link, :user_id, :date2)"
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

// SelectListByLink
// select id, chat_id, link, user_id, date2 from chat_invite_participants where link = :link
// TODO(@benqi): sqlmap
func (dao *ChatInviteParticipantsDAO) SelectListByLink(ctx context.Context, link string) (rList []dataobject.ChatInviteParticipantsDO, err error) {
	var (
		query  = "select id, chat_id, link, user_id, date2 from chat_invite_participants where link = ?"
		values []dataobject.ChatInviteParticipantsDO
	)
	err = dao.db.QueryRowsPartial(ctx, &values, query, link)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectListByLink(_), error: %v", err)
		return
	}

	rList = values

	return
}

// SelectListByLinkWithCB
// select id, chat_id, link, user_id, date2 from chat_invite_participants where link = :link
// TODO(@benqi): sqlmap
func (dao *ChatInviteParticipantsDAO) SelectListByLinkWithCB(ctx context.Context, link string, cb func(i int, v *dataobject.ChatInviteParticipantsDO)) (rList []dataobject.ChatInviteParticipantsDO, err error) {
	var (
		query  = "select id, chat_id, link, user_id, date2 from chat_invite_participants where link = ?"
		values []dataobject.ChatInviteParticipantsDO
	)
	err = dao.db.QueryRowsPartial(ctx, &values, query, link)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectListByLink(_), error: %v", err)
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

// Delete
// delete from chat_invite_participants where chat_id = :chat_id and user_id = :user_id
// TODO(@benqi): sqlmap
func (dao *ChatInviteParticipantsDAO) Delete(ctx context.Context, chat_id int64, user_id int64) (rowsAffected int64, err error) {
	var (
		query   = "delete from chat_invite_participants where chat_id = ? and user_id = ?"
		rResult sql.Result
	)
	rResult, err = dao.db.Exec(ctx, query, chat_id, user_id)

	if err != nil {
		logx.WithContext(ctx).Errorf("exec in Delete(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		logx.WithContext(ctx).Errorf("rowsAffected in Delete(_), error: %v", err)
	}

	return
}

// DeleteTx
// delete from chat_invite_participants where chat_id = :chat_id and user_id = :user_id
// TODO(@benqi): sqlmap
func (dao *ChatInviteParticipantsDAO) DeleteTx(tx *sqlx.Tx, chat_id int64, user_id int64) (rowsAffected int64, err error) {
	var (
		query   = "delete from chat_invite_participants where chat_id = ? and user_id = ?"
		rResult sql.Result
	)
	rResult, err = tx.Exec(query, chat_id, user_id)

	if err != nil {
		logx.WithContext(tx.Context()).Errorf("exec in Delete(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		logx.WithContext(tx.Context()).Errorf("rowsAffected in Delete(_), error: %v", err)
	}

	return
}
