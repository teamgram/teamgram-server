/*
 * WARNING! All changes made in this file will be lost!
 *   Created from by 'dalgen'
 *
 * Copyright (c) 2021-present,  Teamgram Studio (https://teamgram.io).
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
// insert into chat_invite_participants(link, user_id, date2) values (:link, :user_id, :date2)
func (dao *ChatInviteParticipantsDAO) Insert(ctx context.Context, do *dataobject.ChatInviteParticipantsDO) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into chat_invite_participants(link, user_id, date2) values (:link, :user_id, :date2)"
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
// insert into chat_invite_participants(link, user_id, date2) values (:link, :user_id, :date2)
func (dao *ChatInviteParticipantsDAO) InsertTx(tx *sqlx.Tx, do *dataobject.ChatInviteParticipantsDO) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into chat_invite_participants(link, user_id, date2) values (:link, :user_id, :date2)"
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
// select id, link, user_id, date2 from chat_invite_participants where link = :link
func (dao *ChatInviteParticipantsDAO) SelectListByLink(ctx context.Context, link string) (rList []dataobject.ChatInviteParticipantsDO, err error) {
	var (
		query = "select id, link, user_id, date2 from chat_invite_participants where link = ?"
		rows  *sqlx.Rows
	)
	rows, err = dao.db.Query(ctx, query, link)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectListByLink(_), error: %v", err)
		return
	}

	defer rows.Close()

	var values []dataobject.ChatInviteParticipantsDO
	for rows.Next() {
		v := dataobject.ChatInviteParticipantsDO{}

		// TODO(@benqi): not use reflect
		err = rows.StructScan(&v)
		if err != nil {
			logx.WithContext(ctx).Errorf("structScan in SelectListByLink(_), error: %v", err)
			return
		}
		values = append(values, v)
	}
	rList = values

	return
}

// SelectListByLinkWithCB
// select id, link, user_id, date2 from chat_invite_participants where link = :link
func (dao *ChatInviteParticipantsDAO) SelectListByLinkWithCB(ctx context.Context, link string, cb func(i int, v *dataobject.ChatInviteParticipantsDO)) (rList []dataobject.ChatInviteParticipantsDO, err error) {
	var (
		query = "select id, link, user_id, date2 from chat_invite_participants where link = ?"
		rows  *sqlx.Rows
	)
	rows, err = dao.db.Query(ctx, query, link)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectListByLink(_), error: %v", err)
		return
	}

	defer func() {
		rows.Close()
		if err == nil && cb != nil {
			for i := 0; i < len(rList); i++ {
				cb(i, &rList[i])
			}
		}
	}()

	var values []dataobject.ChatInviteParticipantsDO
	for rows.Next() {
		v := dataobject.ChatInviteParticipantsDO{}

		// TODO(@benqi): not use reflect
		err = rows.StructScan(&v)
		if err != nil {
			logx.WithContext(ctx).Errorf("structScan in SelectListByLink(_), error: %v", err)
			return
		}
		values = append(values, v)
	}
	rList = values

	return
}
