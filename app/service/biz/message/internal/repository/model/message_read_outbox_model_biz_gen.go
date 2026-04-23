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
	bizMessageReadOutboxModel interface {
		InsertOrUpdate(ctx context.Context, data *MessageReadOutbox) (lastInsertId, rowsAffected int64, err error)
		InsertOrUpdateTx(tx *sqlx.Tx, data *MessageReadOutbox) (lastInsertId, rowsAffected int64, err error)

		SelectList(ctx context.Context, userId int64, readUserId int64, readOutboxMaxId int32) ([]MessageReadOutbox, error)
		SelectListWithCB(ctx context.Context, userId int64, readUserId int64, readOutboxMaxId int32, cb func(sz, i int, v *MessageReadOutbox)) ([]MessageReadOutbox, error)
	}
)

// InsertOrUpdate
// insert into message_read_outbox(user_id, peer_dialog_id, read_user_id, read_outbox_max_id, read_outbox_max_date) values (:user_id, :peer_dialog_id, :read_user_id, :read_outbox_max_id, :read_outbox_max_date) on duplicate key update read_outbox_max_date = values(read_outbox_max_date)
func (m *defaultMessageReadOutboxModel) InsertOrUpdate(ctx context.Context, data *MessageReadOutbox) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into message_read_outbox(user_id, peer_dialog_id, read_user_id, read_outbox_max_id, read_outbox_max_date) values (:user_id, :peer_dialog_id, :read_user_id, :read_outbox_max_id, :read_outbox_max_date) on duplicate key update read_outbox_max_date = values(read_outbox_max_date)"
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
// insert into message_read_outbox(user_id, peer_dialog_id, read_user_id, read_outbox_max_id, read_outbox_max_date) values (:user_id, :peer_dialog_id, :read_user_id, :read_outbox_max_id, :read_outbox_max_date) on duplicate key update read_outbox_max_date = values(read_outbox_max_date)
func (m *defaultMessageReadOutboxModel) InsertOrUpdateTx(tx *sqlx.Tx, data *MessageReadOutbox) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into message_read_outbox(user_id, peer_dialog_id, read_user_id, read_outbox_max_id, read_outbox_max_date) values (:user_id, :peer_dialog_id, :read_user_id, :read_outbox_max_id, :read_outbox_max_date) on duplicate key update read_outbox_max_date = values(read_outbox_max_date)"
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
// select id, user_id, peer_dialog_id, read_user_id, read_outbox_max_id, read_outbox_max_date from message_read_outbox where user_id = :user_id and read_user_id = :read_user_id and read_outbox_max_id >= :read_outbox_max_id order by read_outbox_max_id asc limit 1
func (m *defaultMessageReadOutboxModel) SelectList(ctx context.Context, userId int64, readUserId int64, readOutboxMaxId int32) (rList []MessageReadOutbox, err error) {
	var (
		query  = "select id, user_id, peer_dialog_id, read_user_id, read_outbox_max_id, read_outbox_max_date from message_read_outbox where user_id = ? and read_user_id = ? and read_outbox_max_id >= ? order by read_outbox_max_id asc limit 1"
		values []MessageReadOutbox
	)
	err = m.db.QueryRowsPartial(ctx, &values, query, userId, readUserId, readOutboxMaxId)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectList(_), error: %v", err)
		return
	}

	rList = values

	return
}

// SelectListWithCB
// select id, user_id, peer_dialog_id, read_user_id, read_outbox_max_id, read_outbox_max_date from message_read_outbox where user_id = :user_id and read_user_id = :read_user_id and read_outbox_max_id >= :read_outbox_max_id order by read_outbox_max_id asc limit 1
func (m *defaultMessageReadOutboxModel) SelectListWithCB(ctx context.Context, userId int64, readUserId int64, readOutboxMaxId int32, cb func(sz, i int, v *MessageReadOutbox)) (rList []MessageReadOutbox, err error) {
	var (
		query  = "select id, user_id, peer_dialog_id, read_user_id, read_outbox_max_id, read_outbox_max_date from message_read_outbox where user_id = ? and read_user_id = ? and read_outbox_max_id >= ? order by read_outbox_max_id asc limit 1"
		values []MessageReadOutbox
	)
	err = m.db.QueryRowsPartial(ctx, &values, query, userId, readUserId, readOutboxMaxId)

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
