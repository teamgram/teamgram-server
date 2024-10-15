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

type MessageReadOutboxDAO struct {
	db *sqlx.DB
}

func NewMessageReadOutboxDAO(db *sqlx.DB) *MessageReadOutboxDAO {
	return &MessageReadOutboxDAO{
		db: db,
	}
}

// InsertOrUpdate
// insert into message_read_outbox(user_id, peer_dialog_id, read_user_id, read_outbox_max_id, read_outbox_max_date) values (:user_id, :peer_dialog_id, :read_user_id, :read_outbox_max_id, :read_outbox_max_date) on duplicate key update read_outbox_max_date = values(read_outbox_max_date)
func (dao *MessageReadOutboxDAO) InsertOrUpdate(ctx context.Context, do *dataobject.MessageReadOutboxDO) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into message_read_outbox(user_id, peer_dialog_id, read_user_id, read_outbox_max_id, read_outbox_max_date) values (:user_id, :peer_dialog_id, :read_user_id, :read_outbox_max_id, :read_outbox_max_date) on duplicate key update read_outbox_max_date = values(read_outbox_max_date)"
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
// insert into message_read_outbox(user_id, peer_dialog_id, read_user_id, read_outbox_max_id, read_outbox_max_date) values (:user_id, :peer_dialog_id, :read_user_id, :read_outbox_max_id, :read_outbox_max_date) on duplicate key update read_outbox_max_date = values(read_outbox_max_date)
func (dao *MessageReadOutboxDAO) InsertOrUpdateTx(tx *sqlx.Tx, do *dataobject.MessageReadOutboxDO) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into message_read_outbox(user_id, peer_dialog_id, read_user_id, read_outbox_max_id, read_outbox_max_date) values (:user_id, :peer_dialog_id, :read_user_id, :read_outbox_max_id, :read_outbox_max_date) on duplicate key update read_outbox_max_date = values(read_outbox_max_date)"
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
// select id, user_id, peer_dialog_id, read_user_id, read_outbox_max_id, read_outbox_max_date from message_read_outbox where user_id = :user_id and read_user_id = :read_user_id and read_outbox_max_id >= :read_outbox_max_id order by read_outbox_max_id asc limit 1
func (dao *MessageReadOutboxDAO) SelectList(ctx context.Context, userId int64, readUserId int64, readOutboxMaxId int32) (rList []dataobject.MessageReadOutboxDO, err error) {
	var (
		query  = "select id, user_id, peer_dialog_id, read_user_id, read_outbox_max_id, read_outbox_max_date from message_read_outbox where user_id = ? and read_user_id = ? and read_outbox_max_id >= ? order by read_outbox_max_id asc limit 1"
		values []dataobject.MessageReadOutboxDO
	)
	err = dao.db.QueryRowsPartial(ctx, &values, query, userId, readUserId, readOutboxMaxId)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectList(_), error: %v", err)
		return
	}

	rList = values

	return
}

// SelectListWithCB
// select id, user_id, peer_dialog_id, read_user_id, read_outbox_max_id, read_outbox_max_date from message_read_outbox where user_id = :user_id and read_user_id = :read_user_id and read_outbox_max_id >= :read_outbox_max_id order by read_outbox_max_id asc limit 1
func (dao *MessageReadOutboxDAO) SelectListWithCB(ctx context.Context, userId int64, readUserId int64, readOutboxMaxId int32, cb func(sz, i int, v *dataobject.MessageReadOutboxDO)) (rList []dataobject.MessageReadOutboxDO, err error) {
	var (
		query  = "select id, user_id, peer_dialog_id, read_user_id, read_outbox_max_id, read_outbox_max_date from message_read_outbox where user_id = ? and read_user_id = ? and read_outbox_max_id >= ? order by read_outbox_max_id asc limit 1"
		values []dataobject.MessageReadOutboxDO
	)
	err = dao.db.QueryRowsPartial(ctx, &values, query, userId, readUserId, readOutboxMaxId)

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
