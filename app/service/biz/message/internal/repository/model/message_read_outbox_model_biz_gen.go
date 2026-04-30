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

type bizMessageReadOutboxModel interface {
	InsertOrUpdate(ctx context.Context, data *MessageReadOutbox) (lastInsertId, rowsAffected int64, err error)
	SelectList(ctx context.Context, userId int64, readUserId int64, readOutboxMaxId int32) ([]MessageReadOutbox, error)
	SelectListWithCB(ctx context.Context, userId int64, readUserId int64, readOutboxMaxId int32, cb func(sz, i int, v *MessageReadOutbox)) ([]MessageReadOutbox, error)
}

type MessageReadOutboxTxModel interface {
	InsertOrUpdate(data *MessageReadOutbox) (lastInsertId, rowsAffected int64, err error)
	SelectList(userId int64, readUserId int64, readOutboxMaxId int32) ([]MessageReadOutbox, error)
}

type defaultMessageReadOutboxTxModel struct {
	tx *sqlx.Tx
}

func NewMessageReadOutboxTxModel(tx *sqlx.Tx) MessageReadOutboxTxModel {
	return &defaultMessageReadOutboxTxModel{tx: tx}
}

// InsertOrUpdate
// insert into message_read_outbox(user_id, peer_dialog_id, read_user_id, read_outbox_max_id, read_outbox_max_date) values (:user_id, :peer_dialog_id, :read_user_id, :read_outbox_max_id, :read_outbox_max_date) on duplicate key update read_outbox_max_date = values(read_outbox_max_date)
func (m *defaultMessageReadOutboxModel) InsertOrUpdate(ctx context.Context, data *MessageReadOutbox) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into message_read_outbox(user_id, peer_dialog_id, read_user_id, read_outbox_max_id, read_outbox_max_date) values (:user_id, :peer_dialog_id, :read_user_id, :read_outbox_max_id, :read_outbox_max_date) on duplicate key update read_outbox_max_date = values(read_outbox_max_date)"
		r     sql.Result
	)

	r, err = m.db.NamedExec(ctx, query, data)
	if err != nil {
		err = fmt.Errorf("message_read_outbox.InsertOrUpdate named exec: %w", err)
		return
	}

	lastInsertId, err = r.LastInsertId()
	if err != nil {
		err = fmt.Errorf("message_read_outbox.InsertOrUpdate last insert id: %w", err)
		return
	}
	rowsAffected, err = r.RowsAffected()
	if err != nil {
		err = fmt.Errorf("message_read_outbox.InsertOrUpdate rows affected: %w", err)
	}

	return

}

// InsertOrUpdate
// insert into message_read_outbox(user_id, peer_dialog_id, read_user_id, read_outbox_max_id, read_outbox_max_date) values (:user_id, :peer_dialog_id, :read_user_id, :read_outbox_max_id, :read_outbox_max_date) on duplicate key update read_outbox_max_date = values(read_outbox_max_date)
func (m *defaultMessageReadOutboxTxModel) InsertOrUpdate(data *MessageReadOutbox) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into message_read_outbox(user_id, peer_dialog_id, read_user_id, read_outbox_max_id, read_outbox_max_date) values (:user_id, :peer_dialog_id, :read_user_id, :read_outbox_max_id, :read_outbox_max_date) on duplicate key update read_outbox_max_date = values(read_outbox_max_date)"
		r     sql.Result
	)

	r, err = m.tx.NamedExec(query, data)
	if err != nil {
		err = fmt.Errorf("message_read_outbox.InsertOrUpdate named exec: %w", err)
		return
	}

	lastInsertId, err = r.LastInsertId()
	if err != nil {
		err = fmt.Errorf("message_read_outbox.InsertOrUpdate last insert id: %w", err)
		return
	}
	rowsAffected, err = r.RowsAffected()
	if err != nil {
		err = fmt.Errorf("message_read_outbox.InsertOrUpdate rows affected: %w", err)
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
		if errors.Is(err, sqlx.ErrNotFound) {
			rList = []MessageReadOutbox{}
			err = nil
			return
		}
		err = fmt.Errorf("message_read_outbox.SelectList: %w", err)
		return
	}

	rList = values

	return
}

// SelectList
// select id, user_id, peer_dialog_id, read_user_id, read_outbox_max_id, read_outbox_max_date from message_read_outbox where user_id = :user_id and read_user_id = :read_user_id and read_outbox_max_id >= :read_outbox_max_id order by read_outbox_max_id asc limit 1
func (m *defaultMessageReadOutboxTxModel) SelectList(userId int64, readUserId int64, readOutboxMaxId int32) (rList []MessageReadOutbox, err error) {
	var (
		query  = "select id, user_id, peer_dialog_id, read_user_id, read_outbox_max_id, read_outbox_max_date from message_read_outbox where user_id = ? and read_user_id = ? and read_outbox_max_id >= ? order by read_outbox_max_id asc limit 1"
		values []MessageReadOutbox
	)
	err = m.tx.QueryRowsPartial(&values, query, userId, readUserId, readOutboxMaxId)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			rList = []MessageReadOutbox{}
			err = nil
			return
		}
		err = fmt.Errorf("message_read_outbox.SelectList: %w", err)
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
		if errors.Is(err, sqlx.ErrNotFound) {
			rList = []MessageReadOutbox{}
			err = nil
			return
		}
		err = fmt.Errorf("message_read_outbox.SelectListWithCB: %w", err)
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
