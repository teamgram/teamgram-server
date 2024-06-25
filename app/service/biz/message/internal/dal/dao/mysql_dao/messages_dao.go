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
	"strconv"
	"strings"

	"github.com/teamgram/marmota/pkg/stores/sqlx"
	"github.com/teamgram/teamgram-server/app/service/biz/message/internal/dal/dataobject"

	"github.com/zeromicro/go-zero/core/logx"
)

var _ *sql.Result
var _ = fmt.Sprintf
var _ = strings.Join
var _ = errors.Is

type MessagesDAO struct {
	db           *sqlx.DB
	ShardingSize int
}

func NewMessagesDAO(db *sqlx.DB, shardingSize int) *MessagesDAO {
	if shardingSize <= 1 {
		shardingSize = 0
	}
	return &MessagesDAO{
		db:           db,
		ShardingSize: shardingSize,
	}
}

func (dao *MessagesDAO) CalcTableName(id int64) string {
	if dao.ShardingSize == 0 {
		return "messages"
	} else {
		return "messages_" + strconv.FormatInt(id%int64(dao.ShardingSize), 10)
	}
}

// InsertOrReturnId
// insert into messages(user_id, user_message_box_id, dialog_id1, dialog_id2, dialog_message_id, sender_user_id, peer_type, peer_id, random_id, message_filter_type, message_data, message, mentioned, media_unread, pinned, saved_peer_type, saved_peer_id, date2, ttl_period) values (:user_id, :user_message_box_id, :dialog_id1, :dialog_id2, :dialog_message_id, :sender_user_id, :peer_type, :peer_id, :random_id, :message_filter_type, :message_data, :message, :mentioned, :media_unread, :pinned, :saved_peer_type, :saved_peer_id, :date2, :ttl_period) on duplicate key update id = last_insert_id(id)
func (dao *MessagesDAO) InsertOrReturnId(ctx context.Context, do *dataobject.MessagesDO) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into " + dao.CalcTableName(do.UserId) + "(user_id, user_message_box_id, dialog_id1, dialog_id2, dialog_message_id, sender_user_id, peer_type, peer_id, random_id, message_filter_type, message_data, message, mentioned, media_unread, pinned, saved_peer_type, saved_peer_id, date2, ttl_period) values (:user_id, :user_message_box_id, :dialog_id1, :dialog_id2, :dialog_message_id, :sender_user_id, :peer_type, :peer_id, :random_id, :message_filter_type, :message_data, :message, :mentioned, :media_unread, :pinned, :saved_peer_type, :saved_peer_id, :date2, :ttl_period) on duplicate key update id = last_insert_id(id)"
		r     sql.Result
	)

	r, err = dao.db.NamedExec(ctx, query, do)
	if err != nil {
		logx.WithContext(ctx).Errorf("namedExec in InsertOrReturnId(%v), error: %v", do, err)
		return
	}

	lastInsertId, err = r.LastInsertId()
	if err != nil {
		logx.WithContext(ctx).Errorf("lastInsertId in InsertOrReturnId(%v)_error: %v", do, err)
		return
	}
	rowsAffected, err = r.RowsAffected()
	if err != nil {
		logx.WithContext(ctx).Errorf("rowsAffected in InsertOrReturnId(%v)_error: %v", do, err)
	}

	return
}

// InsertOrReturnIdTx
// insert into messages(user_id, user_message_box_id, dialog_id1, dialog_id2, dialog_message_id, sender_user_id, peer_type, peer_id, random_id, message_filter_type, message_data, message, mentioned, media_unread, pinned, saved_peer_type, saved_peer_id, date2, ttl_period) values (:user_id, :user_message_box_id, :dialog_id1, :dialog_id2, :dialog_message_id, :sender_user_id, :peer_type, :peer_id, :random_id, :message_filter_type, :message_data, :message, :mentioned, :media_unread, :pinned, :saved_peer_type, :saved_peer_id, :date2, :ttl_period) on duplicate key update id = last_insert_id(id)
func (dao *MessagesDAO) InsertOrReturnIdTx(tx *sqlx.Tx, do *dataobject.MessagesDO) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into " + dao.CalcTableName(do.UserId) + "(user_id, user_message_box_id, dialog_id1, dialog_id2, dialog_message_id, sender_user_id, peer_type, peer_id, random_id, message_filter_type, message_data, message, mentioned, media_unread, pinned, saved_peer_type, saved_peer_id, date2, ttl_period) values (:user_id, :user_message_box_id, :dialog_id1, :dialog_id2, :dialog_message_id, :sender_user_id, :peer_type, :peer_id, :random_id, :message_filter_type, :message_data, :message, :mentioned, :media_unread, :pinned, :saved_peer_type, :saved_peer_id, :date2, :ttl_period) on duplicate key update id = last_insert_id(id)"
		r     sql.Result
	)

	r, err = tx.NamedExec(query, do)
	if err != nil {
		logx.WithContext(tx.Context()).Errorf("namedExec in InsertOrReturnId(%v), error: %v", do, err)
		return
	}

	lastInsertId, err = r.LastInsertId()
	if err != nil {
		logx.WithContext(tx.Context()).Errorf("lastInsertId in InsertOrReturnId(%v)_error: %v", do, err)
		return
	}
	rowsAffected, err = r.RowsAffected()
	if err != nil {
		logx.WithContext(tx.Context()).Errorf("rowsAffected in InsertOrReturnId(%v)_error: %v", do, err)
	}

	return
}

// SelectByRandomId
// select user_id, user_message_box_id, dialog_id1, dialog_id2, dialog_message_id, sender_user_id, peer_type, peer_id, random_id, message_filter_type, message_data, message, mentioned, media_unread, pinned, has_reaction, reaction, reaction_date, reaction_unread, saved_peer_type, saved_peer_id, date2, ttl_period from messages where sender_user_id = :sender_user_id and random_id = :random_id and deleted = 0 limit 1
func (dao *MessagesDAO) SelectByRandomId(ctx context.Context, senderUserId int64, randomId int64) (rValue *dataobject.MessagesDO, err error) {
	var (
		query = "select user_id, user_message_box_id, dialog_id1, dialog_id2, dialog_message_id, sender_user_id, peer_type, peer_id, random_id, message_filter_type, message_data, message, mentioned, media_unread, pinned, has_reaction, reaction, reaction_date, reaction_unread, saved_peer_type, saved_peer_id, date2, ttl_period from " + dao.CalcTableName(senderUserId) + " where sender_user_id = ? and random_id = ? and deleted = 0 limit 1"
		do    = &dataobject.MessagesDO{}
	)
	err = dao.db.QueryRowPartial(ctx, do, query, senderUserId, randomId)

	if err != nil {
		if !errors.Is(err, sqlx.ErrNotFound) {
			logx.WithContext(ctx).Errorf("queryx in SelectByRandomId(_), error: %v", err)
			return
		} else {
			err = nil
		}
	} else {
		rValue = do
	}

	return
}

// SelectByMessageIdList
// select user_id, user_message_box_id, dialog_id1, dialog_id2, dialog_message_id, sender_user_id, peer_type, peer_id, random_id, message_filter_type, message_data, message, mentioned, media_unread, pinned, has_reaction, reaction, reaction_date, reaction_unread, saved_peer_type, saved_peer_id, date2, ttl_period from messages where user_id = :user_id and deleted = 0 and user_message_box_id in (:idList) order by user_message_box_id desc
func (dao *MessagesDAO) SelectByMessageIdList(ctx context.Context, userId int64, idList []int32) (rList []dataobject.MessagesDO, err error) {
	var (
		query  = fmt.Sprintf("select user_id, user_message_box_id, dialog_id1, dialog_id2, dialog_message_id, sender_user_id, peer_type, peer_id, random_id, message_filter_type, message_data, message, mentioned, media_unread, pinned, has_reaction, reaction, reaction_date, reaction_unread, saved_peer_type, saved_peer_id, date2, ttl_period from "+dao.CalcTableName(userId)+" where user_id = ? and deleted = 0 and user_message_box_id in (%s) order by user_message_box_id desc", sqlx.InInt32List(idList))
		values []dataobject.MessagesDO
	)

	if len(idList) == 0 {
		rList = []dataobject.MessagesDO{}
		return
	}

	err = dao.db.QueryRowsPartial(ctx, &values, query, userId)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectByMessageIdList(_), error: %v", err)
		return
	}

	rList = values

	return
}

// SelectByMessageIdListWithCB
// select user_id, user_message_box_id, dialog_id1, dialog_id2, dialog_message_id, sender_user_id, peer_type, peer_id, random_id, message_filter_type, message_data, message, mentioned, media_unread, pinned, has_reaction, reaction, reaction_date, reaction_unread, saved_peer_type, saved_peer_id, date2, ttl_period from messages where user_id = :user_id and deleted = 0 and user_message_box_id in (:idList) order by user_message_box_id desc
func (dao *MessagesDAO) SelectByMessageIdListWithCB(ctx context.Context, userId int64, idList []int32, cb func(sz, i int, v *dataobject.MessagesDO)) (rList []dataobject.MessagesDO, err error) {
	var (
		query  = fmt.Sprintf("select user_id, user_message_box_id, dialog_id1, dialog_id2, dialog_message_id, sender_user_id, peer_type, peer_id, random_id, message_filter_type, message_data, message, mentioned, media_unread, pinned, has_reaction, reaction, reaction_date, reaction_unread, saved_peer_type, saved_peer_id, date2, ttl_period from "+dao.CalcTableName(userId)+" where user_id = ? and deleted = 0 and user_message_box_id in (%s) order by user_message_box_id desc", sqlx.InInt32List(idList))
		values []dataobject.MessagesDO
	)

	if len(idList) == 0 {
		rList = []dataobject.MessagesDO{}
		return
	}

	err = dao.db.QueryRowsPartial(ctx, &values, query, userId)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectByMessageIdList(_), error: %v", err)
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

// SelectByMessageId
// select user_id, user_message_box_id, dialog_id1, dialog_id2, dialog_message_id, sender_user_id, peer_type, peer_id, random_id, message_filter_type, message_data, message, mentioned, media_unread, pinned, has_reaction, reaction, reaction_date, reaction_unread, saved_peer_type, saved_peer_id, date2, ttl_period from messages where user_id = :user_id and user_message_box_id = :user_message_box_id and deleted = 0 limit 1
func (dao *MessagesDAO) SelectByMessageId(ctx context.Context, userId int64, userMessageBoxId int32) (rValue *dataobject.MessagesDO, err error) {
	var (
		query = "select user_id, user_message_box_id, dialog_id1, dialog_id2, dialog_message_id, sender_user_id, peer_type, peer_id, random_id, message_filter_type, message_data, message, mentioned, media_unread, pinned, has_reaction, reaction, reaction_date, reaction_unread, saved_peer_type, saved_peer_id, date2, ttl_period from " + dao.CalcTableName(userId) + " where user_id = ? and user_message_box_id = ? and deleted = 0 limit 1"
		do    = &dataobject.MessagesDO{}
	)
	err = dao.db.QueryRowPartial(ctx, do, query, userId, userMessageBoxId)

	if err != nil {
		if !errors.Is(err, sqlx.ErrNotFound) {
			logx.WithContext(ctx).Errorf("queryx in SelectByMessageId(_), error: %v", err)
			return
		} else {
			err = nil
		}
	} else {
		rValue = do
	}

	return
}

// SelectByMessageDataIdList
// select user_id, user_message_box_id, dialog_id1, dialog_id2, dialog_message_id, sender_user_id, peer_type, peer_id, random_id, message_filter_type, message_data, message, mentioned, media_unread, pinned, has_reaction, reaction, reaction_date, reaction_unread, saved_peer_type, saved_peer_id, date2, ttl_period from messages where deleted = 0 and dialog_message_id in (:idList) order by user_message_box_id desc
// TODO(@benqi): sqlmap
func (dao *MessagesDAO) SelectByMessageDataIdList(ctx context.Context, tableName string, idList []int64) (rList []dataobject.MessagesDO, err error) {
	var (
		query  = "select user_id, user_message_box_id, dialog_id1, dialog_id2, dialog_message_id, sender_user_id, peer_type, peer_id, random_id, message_filter_type, message_data, message, mentioned, media_unread, pinned, has_reaction, reaction, reaction_date, reaction_unread, saved_peer_type, saved_peer_id, date2, ttl_period from " + tableName + " where deleted = 0 and dialog_message_id in (" + sqlx.InInt64List(idList) + ") order by user_message_box_id desc"
		values []dataobject.MessagesDO
	)
	if len(idList) == 0 {
		rList = []dataobject.MessagesDO{}
		return
	}

	err = dao.db.QueryRowsPartial(ctx, &values, query)
	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectByMessageDataIdList(_), error: %v", err)
		return
	}

	rList = values

	return
}

// SelectByMessageDataIdListWithCB
// select user_id, user_message_box_id, dialog_id1, dialog_id2, dialog_message_id, sender_user_id, peer_type, peer_id, random_id, message_filter_type, message_data, message, mentioned, media_unread, pinned, has_reaction, reaction, reaction_date, reaction_unread, saved_peer_type, saved_peer_id, date2, ttl_period from messages where deleted = 0 and dialog_message_id in (:idList) order by user_message_box_id desc
// TODO(@benqi): sqlmap
func (dao *MessagesDAO) SelectByMessageDataIdListWithCB(ctx context.Context, tableName string, idList []int64, cb func(sz, i int, v *dataobject.MessagesDO)) (rList []dataobject.MessagesDO, err error) {
	var (
		query  = "select user_id, user_message_box_id, dialog_id1, dialog_id2, dialog_message_id, sender_user_id, peer_type, peer_id, random_id, message_filter_type, message_data, message, mentioned, media_unread, pinned, has_reaction, reaction, reaction_date, reaction_unread, saved_peer_type, saved_peer_id, date2, ttl_period from " + tableName + " where deleted = 0 and dialog_message_id in (" + sqlx.InInt64List(idList) + ") order by user_message_box_id desc"
		values []dataobject.MessagesDO
	)
	if len(idList) == 0 {
		rList = []dataobject.MessagesDO{}
		return
	}

	err = dao.db.QueryRowsPartial(ctx, &values, query)
	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectByMessageDataIdList(_), error: %v", err)
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

// SelectByMessageDataIdUserIdList
// select user_id, user_message_box_id, dialog_id1, dialog_id2, dialog_message_id, sender_user_id, peer_type, peer_id, random_id, message_filter_type, message_data, message, mentioned, media_unread, pinned, has_reaction, reaction, reaction_date, reaction_unread, saved_peer_type, saved_peer_id, date2, ttl_period from messages where dialog_message_id = :dialog_message_id and user_id in (:idList) and deleted = 0
func (dao *MessagesDAO) SelectByMessageDataIdUserIdList(ctx context.Context, tableName string, dialogMessageId int64, idList []int64) (rList []dataobject.MessagesDO, err error) {
	var (
		query  = fmt.Sprintf("select user_id, user_message_box_id, dialog_id1, dialog_id2, dialog_message_id, sender_user_id, peer_type, peer_id, random_id, message_filter_type, message_data, message, mentioned, media_unread, pinned, has_reaction, reaction, reaction_date, reaction_unread, saved_peer_type, saved_peer_id, date2, ttl_period from "+tableName+" where dialog_message_id = ? and user_id in (%s) and deleted = 0", sqlx.InInt64List(idList))
		values []dataobject.MessagesDO
	)

	if len(idList) == 0 {
		rList = []dataobject.MessagesDO{}
		return
	}

	err = dao.db.QueryRowsPartial(ctx, &values, query, dialogMessageId)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectByMessageDataIdUserIdList(_), error: %v", err)
		return
	}

	rList = values

	return
}

// SelectByMessageDataIdUserIdListWithCB
// select user_id, user_message_box_id, dialog_id1, dialog_id2, dialog_message_id, sender_user_id, peer_type, peer_id, random_id, message_filter_type, message_data, message, mentioned, media_unread, pinned, has_reaction, reaction, reaction_date, reaction_unread, saved_peer_type, saved_peer_id, date2, ttl_period from messages where dialog_message_id = :dialog_message_id and user_id in (:idList) and deleted = 0
func (dao *MessagesDAO) SelectByMessageDataIdUserIdListWithCB(ctx context.Context, tableName string, dialogMessageId int64, idList []int64, cb func(sz, i int, v *dataobject.MessagesDO)) (rList []dataobject.MessagesDO, err error) {
	var (
		query  = fmt.Sprintf("select user_id, user_message_box_id, dialog_id1, dialog_id2, dialog_message_id, sender_user_id, peer_type, peer_id, random_id, message_filter_type, message_data, message, mentioned, media_unread, pinned, has_reaction, reaction, reaction_date, reaction_unread, saved_peer_type, saved_peer_id, date2, ttl_period from "+tableName+" where dialog_message_id = ? and user_id in (%s) and deleted = 0", sqlx.InInt64List(idList))
		values []dataobject.MessagesDO
	)

	if len(idList) == 0 {
		rList = []dataobject.MessagesDO{}
		return
	}

	err = dao.db.QueryRowsPartial(ctx, &values, query, dialogMessageId)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectByMessageDataIdUserIdList(_), error: %v", err)
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

// SelectByMessageDataId
// select user_id, user_message_box_id, dialog_id1, dialog_id2, dialog_message_id, sender_user_id, peer_type, peer_id, random_id, message_filter_type, message_data, message, mentioned, media_unread, pinned, has_reaction, reaction, reaction_date, reaction_unread, saved_peer_type, saved_peer_id, date2, ttl_period from messages where user_id = :user_id and dialog_message_id = :dialog_message_id and deleted = 0 limit 1
func (dao *MessagesDAO) SelectByMessageDataId(ctx context.Context, userId int64, dialogMessageId int64) (rValue *dataobject.MessagesDO, err error) {
	var (
		query = "select user_id, user_message_box_id, dialog_id1, dialog_id2, dialog_message_id, sender_user_id, peer_type, peer_id, random_id, message_filter_type, message_data, message, mentioned, media_unread, pinned, has_reaction, reaction, reaction_date, reaction_unread, saved_peer_type, saved_peer_id, date2, ttl_period from " + dao.CalcTableName(userId) + " where user_id = ? and dialog_message_id = ? and deleted = 0 limit 1"
		do    = &dataobject.MessagesDO{}
	)
	err = dao.db.QueryRowPartial(ctx, do, query, userId, dialogMessageId)

	if err != nil {
		if !errors.Is(err, sqlx.ErrNotFound) {
			logx.WithContext(ctx).Errorf("queryx in SelectByMessageDataId(_), error: %v", err)
			return
		} else {
			err = nil
		}
	} else {
		rValue = do
	}

	return
}

// SelectBackwardByOffsetIdLimit
// select user_id, user_message_box_id, dialog_id1, dialog_id2, dialog_message_id, sender_user_id, peer_type, peer_id, random_id, message_filter_type, message_data, message, mentioned, media_unread, pinned, has_reaction, reaction, reaction_date, reaction_unread, saved_peer_type, saved_peer_id, date2, ttl_period from messages where user_id = :user_id and (dialog_id1 = :dialog_id1 and dialog_id2 = :dialog_id2) and user_message_box_id < :user_message_box_id and deleted = 0 order by user_message_box_id desc limit :limit
func (dao *MessagesDAO) SelectBackwardByOffsetIdLimit(ctx context.Context, userId int64, dialogId1 int64, dialogId2 int64, userMessageBoxId int32, limit int32) (rList []dataobject.MessagesDO, err error) {
	var (
		query  = "select user_id, user_message_box_id, dialog_id1, dialog_id2, dialog_message_id, sender_user_id, peer_type, peer_id, random_id, message_filter_type, message_data, message, mentioned, media_unread, pinned, has_reaction, reaction, reaction_date, reaction_unread, saved_peer_type, saved_peer_id, date2, ttl_period from " + dao.CalcTableName(userId) + " where user_id = ? and (dialog_id1 = ? and dialog_id2 = ?) and user_message_box_id < ? and deleted = 0 order by user_message_box_id desc limit ?"
		values []dataobject.MessagesDO
	)
	err = dao.db.QueryRowsPartial(ctx, &values, query, userId, dialogId1, dialogId2, userMessageBoxId, limit)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectBackwardByOffsetIdLimit(_), error: %v", err)
		return
	}

	rList = values

	return
}

// SelectBackwardByOffsetIdLimitWithCB
// select user_id, user_message_box_id, dialog_id1, dialog_id2, dialog_message_id, sender_user_id, peer_type, peer_id, random_id, message_filter_type, message_data, message, mentioned, media_unread, pinned, has_reaction, reaction, reaction_date, reaction_unread, saved_peer_type, saved_peer_id, date2, ttl_period from messages where user_id = :user_id and (dialog_id1 = :dialog_id1 and dialog_id2 = :dialog_id2) and user_message_box_id < :user_message_box_id and deleted = 0 order by user_message_box_id desc limit :limit
func (dao *MessagesDAO) SelectBackwardByOffsetIdLimitWithCB(ctx context.Context, userId int64, dialogId1 int64, dialogId2 int64, userMessageBoxId int32, limit int32, cb func(sz, i int, v *dataobject.MessagesDO)) (rList []dataobject.MessagesDO, err error) {
	var (
		query  = "select user_id, user_message_box_id, dialog_id1, dialog_id2, dialog_message_id, sender_user_id, peer_type, peer_id, random_id, message_filter_type, message_data, message, mentioned, media_unread, pinned, has_reaction, reaction, reaction_date, reaction_unread, saved_peer_type, saved_peer_id, date2, ttl_period from " + dao.CalcTableName(userId) + " where user_id = ? and (dialog_id1 = ? and dialog_id2 = ?) and user_message_box_id < ? and deleted = 0 order by user_message_box_id desc limit ?"
		values []dataobject.MessagesDO
	)
	err = dao.db.QueryRowsPartial(ctx, &values, query, userId, dialogId1, dialogId2, userMessageBoxId, limit)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectBackwardByOffsetIdLimit(_), error: %v", err)
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

// SelectForwardByOffsetIdLimit
// select user_id, user_message_box_id, dialog_id1, dialog_id2, dialog_message_id, sender_user_id, peer_type, peer_id, random_id, message_filter_type, message_data, message, mentioned, media_unread, pinned, has_reaction, reaction, reaction_date, reaction_unread, saved_peer_type, saved_peer_id, date2, ttl_period from messages where user_id = :user_id and (dialog_id1 = :dialog_id1 and dialog_id2 = :dialog_id2) and user_message_box_id >= :user_message_box_id and deleted = 0 order by user_message_box_id asc limit :limit
func (dao *MessagesDAO) SelectForwardByOffsetIdLimit(ctx context.Context, userId int64, dialogId1 int64, dialogId2 int64, userMessageBoxId int32, limit int32) (rList []dataobject.MessagesDO, err error) {
	var (
		query  = "select user_id, user_message_box_id, dialog_id1, dialog_id2, dialog_message_id, sender_user_id, peer_type, peer_id, random_id, message_filter_type, message_data, message, mentioned, media_unread, pinned, has_reaction, reaction, reaction_date, reaction_unread, saved_peer_type, saved_peer_id, date2, ttl_period from " + dao.CalcTableName(userId) + " where user_id = ? and (dialog_id1 = ? and dialog_id2 = ?) and user_message_box_id >= ? and deleted = 0 order by user_message_box_id asc limit ?"
		values []dataobject.MessagesDO
	)
	err = dao.db.QueryRowsPartial(ctx, &values, query, userId, dialogId1, dialogId2, userMessageBoxId, limit)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectForwardByOffsetIdLimit(_), error: %v", err)
		return
	}

	rList = values

	return
}

// SelectForwardByOffsetIdLimitWithCB
// select user_id, user_message_box_id, dialog_id1, dialog_id2, dialog_message_id, sender_user_id, peer_type, peer_id, random_id, message_filter_type, message_data, message, mentioned, media_unread, pinned, has_reaction, reaction, reaction_date, reaction_unread, saved_peer_type, saved_peer_id, date2, ttl_period from messages where user_id = :user_id and (dialog_id1 = :dialog_id1 and dialog_id2 = :dialog_id2) and user_message_box_id >= :user_message_box_id and deleted = 0 order by user_message_box_id asc limit :limit
func (dao *MessagesDAO) SelectForwardByOffsetIdLimitWithCB(ctx context.Context, userId int64, dialogId1 int64, dialogId2 int64, userMessageBoxId int32, limit int32, cb func(sz, i int, v *dataobject.MessagesDO)) (rList []dataobject.MessagesDO, err error) {
	var (
		query  = "select user_id, user_message_box_id, dialog_id1, dialog_id2, dialog_message_id, sender_user_id, peer_type, peer_id, random_id, message_filter_type, message_data, message, mentioned, media_unread, pinned, has_reaction, reaction, reaction_date, reaction_unread, saved_peer_type, saved_peer_id, date2, ttl_period from " + dao.CalcTableName(userId) + " where user_id = ? and (dialog_id1 = ? and dialog_id2 = ?) and user_message_box_id >= ? and deleted = 0 order by user_message_box_id asc limit ?"
		values []dataobject.MessagesDO
	)
	err = dao.db.QueryRowsPartial(ctx, &values, query, userId, dialogId1, dialogId2, userMessageBoxId, limit)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectForwardByOffsetIdLimit(_), error: %v", err)
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

// SelectBackwardByOffsetDateLimit
// select user_id, user_message_box_id, dialog_id1, dialog_id2, dialog_message_id, sender_user_id, peer_type, peer_id, random_id, message_filter_type, message_data, message, mentioned, media_unread, pinned, has_reaction, reaction, reaction_date, reaction_unread, saved_peer_type, saved_peer_id, date2, ttl_period from messages where user_id = :user_id and (dialog_id1 = :dialog_id1 and dialog_id2 = :dialog_id2) and date2 < :date2 and deleted = 0 order by user_message_box_id desc limit :limit
func (dao *MessagesDAO) SelectBackwardByOffsetDateLimit(ctx context.Context, userId int64, dialogId1 int64, dialogId2 int64, date2 int64, limit int32) (rList []dataobject.MessagesDO, err error) {
	var (
		query  = "select user_id, user_message_box_id, dialog_id1, dialog_id2, dialog_message_id, sender_user_id, peer_type, peer_id, random_id, message_filter_type, message_data, message, mentioned, media_unread, pinned, has_reaction, reaction, reaction_date, reaction_unread, saved_peer_type, saved_peer_id, date2, ttl_period from " + dao.CalcTableName(userId) + " where user_id = ? and (dialog_id1 = ? and dialog_id2 = ?) and date2 < ? and deleted = 0 order by user_message_box_id desc limit ?"
		values []dataobject.MessagesDO
	)
	err = dao.db.QueryRowsPartial(ctx, &values, query, userId, dialogId1, dialogId2, date2, limit)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectBackwardByOffsetDateLimit(_), error: %v", err)
		return
	}

	rList = values

	return
}

// SelectBackwardByOffsetDateLimitWithCB
// select user_id, user_message_box_id, dialog_id1, dialog_id2, dialog_message_id, sender_user_id, peer_type, peer_id, random_id, message_filter_type, message_data, message, mentioned, media_unread, pinned, has_reaction, reaction, reaction_date, reaction_unread, saved_peer_type, saved_peer_id, date2, ttl_period from messages where user_id = :user_id and (dialog_id1 = :dialog_id1 and dialog_id2 = :dialog_id2) and date2 < :date2 and deleted = 0 order by user_message_box_id desc limit :limit
func (dao *MessagesDAO) SelectBackwardByOffsetDateLimitWithCB(ctx context.Context, userId int64, dialogId1 int64, dialogId2 int64, date2 int64, limit int32, cb func(sz, i int, v *dataobject.MessagesDO)) (rList []dataobject.MessagesDO, err error) {
	var (
		query  = "select user_id, user_message_box_id, dialog_id1, dialog_id2, dialog_message_id, sender_user_id, peer_type, peer_id, random_id, message_filter_type, message_data, message, mentioned, media_unread, pinned, has_reaction, reaction, reaction_date, reaction_unread, saved_peer_type, saved_peer_id, date2, ttl_period from " + dao.CalcTableName(userId) + " where user_id = ? and (dialog_id1 = ? and dialog_id2 = ?) and date2 < ? and deleted = 0 order by user_message_box_id desc limit ?"
		values []dataobject.MessagesDO
	)
	err = dao.db.QueryRowsPartial(ctx, &values, query, userId, dialogId1, dialogId2, date2, limit)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectBackwardByOffsetDateLimit(_), error: %v", err)
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

// SelectForwardByOffsetDateLimit
// select user_id, user_message_box_id, dialog_id1, dialog_id2, dialog_message_id, sender_user_id, peer_type, peer_id, random_id, message_filter_type, message_data, message, mentioned, media_unread, pinned, has_reaction, reaction, reaction_date, reaction_unread, saved_peer_type, saved_peer_id, date2, ttl_period from messages where user_id = :user_id and (dialog_id1 = :dialog_id1 and dialog_id2 = :dialog_id2) and date2 >= :date2 and deleted = 0 order by user_message_box_id asc limit :limit
func (dao *MessagesDAO) SelectForwardByOffsetDateLimit(ctx context.Context, userId int64, dialogId1 int64, dialogId2 int64, date2 int64, limit int32) (rList []dataobject.MessagesDO, err error) {
	var (
		query  = "select user_id, user_message_box_id, dialog_id1, dialog_id2, dialog_message_id, sender_user_id, peer_type, peer_id, random_id, message_filter_type, message_data, message, mentioned, media_unread, pinned, has_reaction, reaction, reaction_date, reaction_unread, saved_peer_type, saved_peer_id, date2, ttl_period from " + dao.CalcTableName(userId) + " where user_id = ? and (dialog_id1 = ? and dialog_id2 = ?) and date2 >= ? and deleted = 0 order by user_message_box_id asc limit ?"
		values []dataobject.MessagesDO
	)
	err = dao.db.QueryRowsPartial(ctx, &values, query, userId, dialogId1, dialogId2, date2, limit)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectForwardByOffsetDateLimit(_), error: %v", err)
		return
	}

	rList = values

	return
}

// SelectForwardByOffsetDateLimitWithCB
// select user_id, user_message_box_id, dialog_id1, dialog_id2, dialog_message_id, sender_user_id, peer_type, peer_id, random_id, message_filter_type, message_data, message, mentioned, media_unread, pinned, has_reaction, reaction, reaction_date, reaction_unread, saved_peer_type, saved_peer_id, date2, ttl_period from messages where user_id = :user_id and (dialog_id1 = :dialog_id1 and dialog_id2 = :dialog_id2) and date2 >= :date2 and deleted = 0 order by user_message_box_id asc limit :limit
func (dao *MessagesDAO) SelectForwardByOffsetDateLimitWithCB(ctx context.Context, userId int64, dialogId1 int64, dialogId2 int64, date2 int64, limit int32, cb func(sz, i int, v *dataobject.MessagesDO)) (rList []dataobject.MessagesDO, err error) {
	var (
		query  = "select user_id, user_message_box_id, dialog_id1, dialog_id2, dialog_message_id, sender_user_id, peer_type, peer_id, random_id, message_filter_type, message_data, message, mentioned, media_unread, pinned, has_reaction, reaction, reaction_date, reaction_unread, saved_peer_type, saved_peer_id, date2, ttl_period from " + dao.CalcTableName(userId) + " where user_id = ? and (dialog_id1 = ? and dialog_id2 = ?) and date2 >= ? and deleted = 0 order by user_message_box_id asc limit ?"
		values []dataobject.MessagesDO
	)
	err = dao.db.QueryRowsPartial(ctx, &values, query, userId, dialogId1, dialogId2, date2, limit)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectForwardByOffsetDateLimit(_), error: %v", err)
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

// SelectPeerUserMessageId
// select user_message_box_id, message_box_type from messages where user_id = :peerId and deleted = 0 and dialog_message_id = (select dialog_message_id from messages where user_id = :user_id and user_message_box_id = :user_message_box_id and deleted = 0 limit 1)
func (dao *MessagesDAO) SelectPeerUserMessageId(ctx context.Context, peerId int64, userId int64, userMessageBoxId int32) (rValue *dataobject.MessagesDO, err error) {
	var (
		query = "select user_message_box_id, message_box_type from " + dao.CalcTableName(peerId) + " where user_id = ? and deleted = 0 and dialog_message_id = (select dialog_message_id from " + dao.CalcTableName(userId) + " where user_id = ? and user_message_box_id = ? and deleted = 0 limit 1)"
		do    = &dataobject.MessagesDO{}
	)
	err = dao.db.QueryRowPartial(ctx, do, query, peerId, userId, userMessageBoxId)

	if err != nil {
		if !errors.Is(err, sqlx.ErrNotFound) {
			logx.WithContext(ctx).Errorf("queryx in SelectPeerUserMessageId(_), error: %v", err)
			return
		} else {
			err = nil
		}
	} else {
		rValue = do
	}

	return
}

// SelectPeerUserMessage
// select user_id, user_message_box_id, dialog_id1, dialog_id2, dialog_message_id, sender_user_id, peer_type, peer_id, random_id, message_filter_type, message_data, message, mentioned, media_unread, pinned, has_reaction, reaction, reaction_date, reaction_unread, saved_peer_type, saved_peer_id, date2, ttl_period from messages where user_id = :peerId and deleted = 0 and dialog_message_id = (select dialog_message_id from messages where user_id = :user_id and user_message_box_id = :user_message_box_id and deleted = 0 limit 1)
func (dao *MessagesDAO) SelectPeerUserMessage(ctx context.Context, peerId int64, userId int64, userMessageBoxId int32) (rValue *dataobject.MessagesDO, err error) {
	var (
		query = "select user_id, user_message_box_id, dialog_id1, dialog_id2, dialog_message_id, sender_user_id, peer_type, peer_id, random_id, message_filter_type, message_data, message, mentioned, media_unread, pinned, has_reaction, reaction, reaction_date, reaction_unread, saved_peer_type, saved_peer_id, date2, ttl_period from " + dao.CalcTableName(peerId) + " where user_id = ? and deleted = 0 and dialog_message_id = (select dialog_message_id from " + dao.CalcTableName(userId) + " where user_id = ? and user_message_box_id = ? and deleted = 0 limit 1)"
		do    = &dataobject.MessagesDO{}
	)
	err = dao.db.QueryRowPartial(ctx, do, query, peerId, userId, userMessageBoxId)

	if err != nil {
		if !errors.Is(err, sqlx.ErrNotFound) {
			logx.WithContext(ctx).Errorf("queryx in SelectPeerUserMessage(_), error: %v", err)
			return
		} else {
			err = nil
		}
	} else {
		rValue = do
	}

	return
}

// SelectDialogLastMessageId
// select user_message_box_id from messages where user_id = :user_id and dialog_id1 = :dialog_id1 and dialog_id2 = :dialog_id2 and deleted = 0 order by user_message_box_id desc limit 1
func (dao *MessagesDAO) SelectDialogLastMessageId(ctx context.Context, userId int64, dialogId1 int64, dialogId2 int64) (rValue int32, err error) {
	var query = "select user_message_box_id from " + dao.CalcTableName(userId) + " where user_id = ? and dialog_id1 = ? and dialog_id2 = ? and deleted = 0 order by user_message_box_id desc limit 1"
	err = dao.db.QueryRowPartial(ctx, &rValue, query, userId, dialogId1, dialogId2)

	if err != nil {
		if !errors.Is(err, sqlx.ErrNotFound) {
			logx.WithContext(ctx).Errorf("get in SelectDialogLastMessageId(_), error: %v", err)
			return
		} else {
			err = nil
		}
	}

	return
}

// SelectDialogLastMessageIdNotIdList
// select user_message_box_id from messages where user_id = :user_id and dialog_id1 = :dialog_id1 and dialog_id2 = :dialog_id2 and user_message_box_id not in (:idList) and deleted = 0 order by user_message_box_id desc limit 1
func (dao *MessagesDAO) SelectDialogLastMessageIdNotIdList(ctx context.Context, userId int64, dialogId1 int64, dialogId2 int64, idList []int32) (rValue int32, err error) {
	var (
		query = fmt.Sprintf("select user_message_box_id from "+dao.CalcTableName(userId)+" where user_id = ? and dialog_id1 = ? and dialog_id2 = ? and user_message_box_id not in (%s) and deleted = 0 order by user_message_box_id desc limit 1", sqlx.InInt32List(idList))
	)

	if len(idList) == 0 {
		return
	}

	err = dao.db.QueryRowPartial(ctx, &rValue, query, userId, dialogId1, dialogId2)

	if err != nil {
		if !errors.Is(err, sqlx.ErrNotFound) {
			logx.WithContext(ctx).Errorf("get in SelectDialogLastMessageIdNotIdList(_), error: %v", err)
			return
		} else {
			err = nil
		}
	}

	return
}

// SelectDialogsByMessageIdList
// select user_id, user_message_box_id, dialog_id1, dialog_id2, dialog_message_id, sender_user_id, peer_type, peer_id, random_id, message_filter_type, message_data, message, mentioned, media_unread, pinned, has_reaction, reaction, reaction_date, reaction_unread, saved_peer_type, saved_peer_id, date2, ttl_period from messages where user_id = :user_id and user_message_box_id in (:idList) and deleted = 0
func (dao *MessagesDAO) SelectDialogsByMessageIdList(ctx context.Context, userId int64, idList []int32) (rList []dataobject.MessagesDO, err error) {
	var (
		query  = fmt.Sprintf("select user_id, user_message_box_id, dialog_id1, dialog_id2, dialog_message_id, sender_user_id, peer_type, peer_id, random_id, message_filter_type, message_data, message, mentioned, media_unread, pinned, has_reaction, reaction, reaction_date, reaction_unread, saved_peer_type, saved_peer_id, date2, ttl_period from "+dao.CalcTableName(userId)+" where user_id = ? and user_message_box_id in (%s) and deleted = 0", sqlx.InInt32List(idList))
		values []dataobject.MessagesDO
	)

	if len(idList) == 0 {
		rList = []dataobject.MessagesDO{}
		return
	}

	err = dao.db.QueryRowsPartial(ctx, &values, query, userId)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectDialogsByMessageIdList(_), error: %v", err)
		return
	}

	rList = values

	return
}

// SelectDialogsByMessageIdListWithCB
// select user_id, user_message_box_id, dialog_id1, dialog_id2, dialog_message_id, sender_user_id, peer_type, peer_id, random_id, message_filter_type, message_data, message, mentioned, media_unread, pinned, has_reaction, reaction, reaction_date, reaction_unread, saved_peer_type, saved_peer_id, date2, ttl_period from messages where user_id = :user_id and user_message_box_id in (:idList) and deleted = 0
func (dao *MessagesDAO) SelectDialogsByMessageIdListWithCB(ctx context.Context, userId int64, idList []int32, cb func(sz, i int, v *dataobject.MessagesDO)) (rList []dataobject.MessagesDO, err error) {
	var (
		query  = fmt.Sprintf("select user_id, user_message_box_id, dialog_id1, dialog_id2, dialog_message_id, sender_user_id, peer_type, peer_id, random_id, message_filter_type, message_data, message, mentioned, media_unread, pinned, has_reaction, reaction, reaction_date, reaction_unread, saved_peer_type, saved_peer_id, date2, ttl_period from "+dao.CalcTableName(userId)+" where user_id = ? and user_message_box_id in (%s) and deleted = 0", sqlx.InInt32List(idList))
		values []dataobject.MessagesDO
	)

	if len(idList) == 0 {
		rList = []dataobject.MessagesDO{}
		return
	}

	err = dao.db.QueryRowsPartial(ctx, &values, query, userId)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectDialogsByMessageIdList(_), error: %v", err)
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

// SelectDialogLastMessageList
// select user_id, user_message_box_id, dialog_id1, dialog_id2, dialog_message_id, sender_user_id, peer_type, peer_id, random_id, message_filter_type, message_data, message, mentioned, media_unread, pinned, has_reaction, reaction, reaction_date, reaction_unread, saved_peer_type, saved_peer_id, date2, ttl_period from messages where user_id = :user_id and (dialog_id1 = :dialog_id1 and dialog_id2 = :dialog_id2) and deleted = 0 order by user_message_box_id desc limit :limit
func (dao *MessagesDAO) SelectDialogLastMessageList(ctx context.Context, userId int64, dialogId1 int64, dialogId2 int64, limit int32) (rList []dataobject.MessagesDO, err error) {
	var (
		query  = "select user_id, user_message_box_id, dialog_id1, dialog_id2, dialog_message_id, sender_user_id, peer_type, peer_id, random_id, message_filter_type, message_data, message, mentioned, media_unread, pinned, has_reaction, reaction, reaction_date, reaction_unread, saved_peer_type, saved_peer_id, date2, ttl_period from " + dao.CalcTableName(userId) + " where user_id = ? and (dialog_id1 = ? and dialog_id2 = ?) and deleted = 0 order by user_message_box_id desc limit ?"
		values []dataobject.MessagesDO
	)
	err = dao.db.QueryRowsPartial(ctx, &values, query, userId, dialogId1, dialogId2, limit)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectDialogLastMessageList(_), error: %v", err)
		return
	}

	rList = values

	return
}

// SelectDialogLastMessageListWithCB
// select user_id, user_message_box_id, dialog_id1, dialog_id2, dialog_message_id, sender_user_id, peer_type, peer_id, random_id, message_filter_type, message_data, message, mentioned, media_unread, pinned, has_reaction, reaction, reaction_date, reaction_unread, saved_peer_type, saved_peer_id, date2, ttl_period from messages where user_id = :user_id and (dialog_id1 = :dialog_id1 and dialog_id2 = :dialog_id2) and deleted = 0 order by user_message_box_id desc limit :limit
func (dao *MessagesDAO) SelectDialogLastMessageListWithCB(ctx context.Context, userId int64, dialogId1 int64, dialogId2 int64, limit int32, cb func(sz, i int, v *dataobject.MessagesDO)) (rList []dataobject.MessagesDO, err error) {
	var (
		query  = "select user_id, user_message_box_id, dialog_id1, dialog_id2, dialog_message_id, sender_user_id, peer_type, peer_id, random_id, message_filter_type, message_data, message, mentioned, media_unread, pinned, has_reaction, reaction, reaction_date, reaction_unread, saved_peer_type, saved_peer_id, date2, ttl_period from " + dao.CalcTableName(userId) + " where user_id = ? and (dialog_id1 = ? and dialog_id2 = ?) and deleted = 0 order by user_message_box_id desc limit ?"
		values []dataobject.MessagesDO
	)
	err = dao.db.QueryRowsPartial(ctx, &values, query, userId, dialogId1, dialogId2, limit)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectDialogLastMessageList(_), error: %v", err)
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

// DeleteMessagesByMessageIdList
// update messages set deleted = 1 where user_id = :user_id and user_message_box_id in (:idList) and deleted = 0
func (dao *MessagesDAO) DeleteMessagesByMessageIdList(ctx context.Context, userId int64, idList []int32) (rowsAffected int64, err error) {
	var (
		query   = fmt.Sprintf("update "+dao.CalcTableName(userId)+" set deleted = 1 where user_id = ? and user_message_box_id in (%s) and deleted = 0", sqlx.InInt32List(idList))
		rResult sql.Result
	)

	if len(idList) == 0 {
		return
	}

	rResult, err = dao.db.Exec(ctx, query, userId)

	if err != nil {
		logx.WithContext(ctx).Errorf("exec in DeleteMessagesByMessageIdList(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		logx.WithContext(ctx).Errorf("rowsAffected in DeleteMessagesByMessageIdList(_), error: %v", err)
	}

	return
}

// DeleteMessagesByMessageIdListTx
// update messages set deleted = 1 where user_id = :user_id and user_message_box_id in (:idList) and deleted = 0
func (dao *MessagesDAO) DeleteMessagesByMessageIdListTx(tx *sqlx.Tx, userId int64, idList []int32) (rowsAffected int64, err error) {
	var (
		query   = fmt.Sprintf("update "+dao.CalcTableName(userId)+" set deleted = 1 where user_id = ? and user_message_box_id in (%s) and deleted = 0", sqlx.InInt32List(idList))
		rResult sql.Result
	)

	if len(idList) == 0 {
		return
	}

	rResult, err = tx.Exec(query, userId)

	if err != nil {
		logx.WithContext(tx.Context()).Errorf("exec in DeleteMessagesByMessageIdList(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		logx.WithContext(tx.Context()).Errorf("rowsAffected in DeleteMessagesByMessageIdList(_), error: %v", err)
	}

	return
}

// SelectDialogMessageIdList
// select user_id, user_message_box_id, dialog_id1, dialog_id2, dialog_message_id, sender_user_id, peer_type, peer_id, random_id, message_filter_type, message_data, message, mentioned, media_unread, pinned, has_reaction, reaction, reaction_date, reaction_unread, saved_peer_type, saved_peer_id, date2, ttl_period from messages where user_id = :user_id and (dialog_id1 = :dialog_id1 and dialog_id2 = :dialog_id2) and deleted = 0 order by user_message_box_id desc
func (dao *MessagesDAO) SelectDialogMessageIdList(ctx context.Context, userId int64, dialogId1 int64, dialogId2 int64) (rList []dataobject.MessagesDO, err error) {
	var (
		query  = "select user_id, user_message_box_id, dialog_id1, dialog_id2, dialog_message_id, sender_user_id, peer_type, peer_id, random_id, message_filter_type, message_data, message, mentioned, media_unread, pinned, has_reaction, reaction, reaction_date, reaction_unread, saved_peer_type, saved_peer_id, date2, ttl_period from " + dao.CalcTableName(userId) + " where user_id = ? and (dialog_id1 = ? and dialog_id2 = ?) and deleted = 0 order by user_message_box_id desc"
		values []dataobject.MessagesDO
	)
	err = dao.db.QueryRowsPartial(ctx, &values, query, userId, dialogId1, dialogId2)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectDialogMessageIdList(_), error: %v", err)
		return
	}

	rList = values

	return
}

// SelectDialogMessageIdListWithCB
// select user_id, user_message_box_id, dialog_id1, dialog_id2, dialog_message_id, sender_user_id, peer_type, peer_id, random_id, message_filter_type, message_data, message, mentioned, media_unread, pinned, has_reaction, reaction, reaction_date, reaction_unread, saved_peer_type, saved_peer_id, date2, ttl_period from messages where user_id = :user_id and (dialog_id1 = :dialog_id1 and dialog_id2 = :dialog_id2) and deleted = 0 order by user_message_box_id desc
func (dao *MessagesDAO) SelectDialogMessageIdListWithCB(ctx context.Context, userId int64, dialogId1 int64, dialogId2 int64, cb func(sz, i int, v *dataobject.MessagesDO)) (rList []dataobject.MessagesDO, err error) {
	var (
		query  = "select user_id, user_message_box_id, dialog_id1, dialog_id2, dialog_message_id, sender_user_id, peer_type, peer_id, random_id, message_filter_type, message_data, message, mentioned, media_unread, pinned, has_reaction, reaction, reaction_date, reaction_unread, saved_peer_type, saved_peer_id, date2, ttl_period from " + dao.CalcTableName(userId) + " where user_id = ? and (dialog_id1 = ? and dialog_id2 = ?) and deleted = 0 order by user_message_box_id desc"
		values []dataobject.MessagesDO
	)
	err = dao.db.QueryRowsPartial(ctx, &values, query, userId, dialogId1, dialogId2)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectDialogMessageIdList(_), error: %v", err)
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

// UpdateMediaUnread
// update messages set media_unread = 0 where user_id = :user_id and user_message_box_id = :user_message_box_id
func (dao *MessagesDAO) UpdateMediaUnread(ctx context.Context, userId int64, userMessageBoxId int32) (rowsAffected int64, err error) {
	var (
		query   = "update " + dao.CalcTableName(userId) + " set media_unread = 0 where user_id = ? and user_message_box_id = ?"
		rResult sql.Result
	)

	rResult, err = dao.db.Exec(ctx, query, userId, userMessageBoxId)

	if err != nil {
		logx.WithContext(ctx).Errorf("exec in UpdateMediaUnread(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		logx.WithContext(ctx).Errorf("rowsAffected in UpdateMediaUnread(_), error: %v", err)
	}

	return
}

// UpdateMediaUnreadTx
// update messages set media_unread = 0 where user_id = :user_id and user_message_box_id = :user_message_box_id
func (dao *MessagesDAO) UpdateMediaUnreadTx(tx *sqlx.Tx, userId int64, userMessageBoxId int32) (rowsAffected int64, err error) {
	var (
		query   = "update " + dao.CalcTableName(userId) + " set media_unread = 0 where user_id = ? and user_message_box_id = ?"
		rResult sql.Result
	)
	rResult, err = tx.Exec(query, userId, userMessageBoxId)

	if err != nil {
		logx.WithContext(tx.Context()).Errorf("exec in UpdateMediaUnread(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		logx.WithContext(tx.Context()).Errorf("rowsAffected in UpdateMediaUnread(_), error: %v", err)
	}

	return
}

// UpdateMentionedAndMediaUnread
// update messages set mentioned = 0, media_unread = 0 where user_id = :user_id and user_message_box_id = :user_message_box_id
func (dao *MessagesDAO) UpdateMentionedAndMediaUnread(ctx context.Context, userId int64, userMessageBoxId int32) (rowsAffected int64, err error) {
	var (
		query   = "update " + dao.CalcTableName(userId) + " set mentioned = 0, media_unread = 0 where user_id = ? and user_message_box_id = ?"
		rResult sql.Result
	)

	rResult, err = dao.db.Exec(ctx, query, userId, userMessageBoxId)

	if err != nil {
		logx.WithContext(ctx).Errorf("exec in UpdateMentionedAndMediaUnread(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		logx.WithContext(ctx).Errorf("rowsAffected in UpdateMentionedAndMediaUnread(_), error: %v", err)
	}

	return
}

// UpdateMentionedAndMediaUnreadTx
// update messages set mentioned = 0, media_unread = 0 where user_id = :user_id and user_message_box_id = :user_message_box_id
func (dao *MessagesDAO) UpdateMentionedAndMediaUnreadTx(tx *sqlx.Tx, userId int64, userMessageBoxId int32) (rowsAffected int64, err error) {
	var (
		query   = "update " + dao.CalcTableName(userId) + " set mentioned = 0, media_unread = 0 where user_id = ? and user_message_box_id = ?"
		rResult sql.Result
	)
	rResult, err = tx.Exec(query, userId, userMessageBoxId)

	if err != nil {
		logx.WithContext(tx.Context()).Errorf("exec in UpdateMentionedAndMediaUnread(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		logx.WithContext(tx.Context()).Errorf("rowsAffected in UpdateMentionedAndMediaUnread(_), error: %v", err)
	}

	return
}

// SelectByMediaType
// select user_id, user_message_box_id, dialog_id1, dialog_id2, dialog_message_id, sender_user_id, peer_type, peer_id, random_id, message_filter_type, message_data, message, mentioned, media_unread, pinned, has_reaction, reaction, reaction_date, reaction_unread, saved_peer_type, saved_peer_id, date2, ttl_period from messages where user_id = :user_id and (dialog_id1 = :dialog_id1 and dialog_id2 = :dialog_id2) and message_filter_type = :message_filter_type and user_message_box_id < :user_message_box_id and deleted = 0 order by user_message_box_id desc limit :limit
func (dao *MessagesDAO) SelectByMediaType(ctx context.Context, userId int64, dialogId1 int64, dialogId2 int64, messageFilterType int32, userMessageBoxId int32, limit int32) (rList []dataobject.MessagesDO, err error) {
	var (
		query  = "select user_id, user_message_box_id, dialog_id1, dialog_id2, dialog_message_id, sender_user_id, peer_type, peer_id, random_id, message_filter_type, message_data, message, mentioned, media_unread, pinned, has_reaction, reaction, reaction_date, reaction_unread, saved_peer_type, saved_peer_id, date2, ttl_period from " + dao.CalcTableName(userId) + " where user_id = ? and (dialog_id1 = ? and dialog_id2 = ?) and message_filter_type = ? and user_message_box_id < ? and deleted = 0 order by user_message_box_id desc limit ?"
		values []dataobject.MessagesDO
	)
	err = dao.db.QueryRowsPartial(ctx, &values, query, userId, dialogId1, dialogId2, messageFilterType, userMessageBoxId, limit)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectByMediaType(_), error: %v", err)
		return
	}

	rList = values

	return
}

// SelectByMediaTypeWithCB
// select user_id, user_message_box_id, dialog_id1, dialog_id2, dialog_message_id, sender_user_id, peer_type, peer_id, random_id, message_filter_type, message_data, message, mentioned, media_unread, pinned, has_reaction, reaction, reaction_date, reaction_unread, saved_peer_type, saved_peer_id, date2, ttl_period from messages where user_id = :user_id and (dialog_id1 = :dialog_id1 and dialog_id2 = :dialog_id2) and message_filter_type = :message_filter_type and user_message_box_id < :user_message_box_id and deleted = 0 order by user_message_box_id desc limit :limit
func (dao *MessagesDAO) SelectByMediaTypeWithCB(ctx context.Context, userId int64, dialogId1 int64, dialogId2 int64, messageFilterType int32, userMessageBoxId int32, limit int32, cb func(sz, i int, v *dataobject.MessagesDO)) (rList []dataobject.MessagesDO, err error) {
	var (
		query  = "select user_id, user_message_box_id, dialog_id1, dialog_id2, dialog_message_id, sender_user_id, peer_type, peer_id, random_id, message_filter_type, message_data, message, mentioned, media_unread, pinned, has_reaction, reaction, reaction_date, reaction_unread, saved_peer_type, saved_peer_id, date2, ttl_period from " + dao.CalcTableName(userId) + " where user_id = ? and (dialog_id1 = ? and dialog_id2 = ?) and message_filter_type = ? and user_message_box_id < ? and deleted = 0 order by user_message_box_id desc limit ?"
		values []dataobject.MessagesDO
	)
	err = dao.db.QueryRowsPartial(ctx, &values, query, userId, dialogId1, dialogId2, messageFilterType, userMessageBoxId, limit)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectByMediaType(_), error: %v", err)
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

// SelectPhoneCallList
// select user_id, user_message_box_id, dialog_id1, dialog_id2, dialog_message_id, sender_user_id, peer_type, peer_id, random_id, message_filter_type, message_data, message, mentioned, media_unread, pinned, has_reaction, reaction, reaction_date, reaction_unread, saved_peer_type, saved_peer_id, date2, ttl_period from messages where user_id = :user_id and message_filter_type = :message_filter_type and user_message_box_id < :user_message_box_id and deleted = 0 order by user_message_box_id desc limit :limit
func (dao *MessagesDAO) SelectPhoneCallList(ctx context.Context, userId int64, messageFilterType int32, userMessageBoxId int32, limit int32) (rList []dataobject.MessagesDO, err error) {
	var (
		query  = "select user_id, user_message_box_id, dialog_id1, dialog_id2, dialog_message_id, sender_user_id, peer_type, peer_id, random_id, message_filter_type, message_data, message, mentioned, media_unread, pinned, has_reaction, reaction, reaction_date, reaction_unread, saved_peer_type, saved_peer_id, date2, ttl_period from " + dao.CalcTableName(userId) + " where user_id = ? and message_filter_type = ? and user_message_box_id < ? and deleted = 0 order by user_message_box_id desc limit ?"
		values []dataobject.MessagesDO
	)
	err = dao.db.QueryRowsPartial(ctx, &values, query, userId, messageFilterType, userMessageBoxId, limit)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectPhoneCallList(_), error: %v", err)
		return
	}

	rList = values

	return
}

// SelectPhoneCallListWithCB
// select user_id, user_message_box_id, dialog_id1, dialog_id2, dialog_message_id, sender_user_id, peer_type, peer_id, random_id, message_filter_type, message_data, message, mentioned, media_unread, pinned, has_reaction, reaction, reaction_date, reaction_unread, saved_peer_type, saved_peer_id, date2, ttl_period from messages where user_id = :user_id and message_filter_type = :message_filter_type and user_message_box_id < :user_message_box_id and deleted = 0 order by user_message_box_id desc limit :limit
func (dao *MessagesDAO) SelectPhoneCallListWithCB(ctx context.Context, userId int64, messageFilterType int32, userMessageBoxId int32, limit int32, cb func(sz, i int, v *dataobject.MessagesDO)) (rList []dataobject.MessagesDO, err error) {
	var (
		query  = "select user_id, user_message_box_id, dialog_id1, dialog_id2, dialog_message_id, sender_user_id, peer_type, peer_id, random_id, message_filter_type, message_data, message, mentioned, media_unread, pinned, has_reaction, reaction, reaction_date, reaction_unread, saved_peer_type, saved_peer_id, date2, ttl_period from " + dao.CalcTableName(userId) + " where user_id = ? and message_filter_type = ? and user_message_box_id < ? and deleted = 0 order by user_message_box_id desc limit ?"
		values []dataobject.MessagesDO
	)
	err = dao.db.QueryRowsPartial(ctx, &values, query, userId, messageFilterType, userMessageBoxId, limit)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectPhoneCallList(_), error: %v", err)
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

// Search
// select user_id, user_message_box_id, dialog_id1, dialog_id2, dialog_message_id, sender_user_id, peer_type, peer_id, random_id, message_filter_type, message_data, message, mentioned, media_unread, pinned, has_reaction, reaction, reaction_date, reaction_unread, saved_peer_type, saved_peer_id, date2, ttl_period from messages where user_id = :user_id and (dialog_id1 = :dialog_id1 and dialog_id2 = :dialog_id2) and user_message_box_id < :user_message_box_id and deleted = 0 and message != ” and message like :q2 order by user_message_box_id desc limit :limit
func (dao *MessagesDAO) Search(ctx context.Context, userId int64, dialogId1 int64, dialogId2 int64, userMessageBoxId int32, q2 string, limit int32) (rList []dataobject.MessagesDO, err error) {
	var (
		query  = "select user_id, user_message_box_id, dialog_id1, dialog_id2, dialog_message_id, sender_user_id, peer_type, peer_id, random_id, message_filter_type, message_data, message, mentioned, media_unread, pinned, has_reaction, reaction, reaction_date, reaction_unread, saved_peer_type, saved_peer_id, date2, ttl_period from " + dao.CalcTableName(userId) + " where user_id = ? and (dialog_id1 = ? and dialog_id2 = ?) and user_message_box_id < ? and deleted = 0 and message != '' and message like ? order by user_message_box_id desc limit ?"
		values []dataobject.MessagesDO
	)
	err = dao.db.QueryRowsPartial(ctx, &values, query, userId, dialogId1, dialogId2, userMessageBoxId, q2, limit)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in Search(_), error: %v", err)
		return
	}

	rList = values

	return
}

// SearchWithCB
// select user_id, user_message_box_id, dialog_id1, dialog_id2, dialog_message_id, sender_user_id, peer_type, peer_id, random_id, message_filter_type, message_data, message, mentioned, media_unread, pinned, has_reaction, reaction, reaction_date, reaction_unread, saved_peer_type, saved_peer_id, date2, ttl_period from messages where user_id = :user_id and (dialog_id1 = :dialog_id1 and dialog_id2 = :dialog_id2) and user_message_box_id < :user_message_box_id and deleted = 0 and message != ” and message like :q2 order by user_message_box_id desc limit :limit
func (dao *MessagesDAO) SearchWithCB(ctx context.Context, userId int64, dialogId1 int64, dialogId2 int64, userMessageBoxId int32, q2 string, limit int32, cb func(sz, i int, v *dataobject.MessagesDO)) (rList []dataobject.MessagesDO, err error) {
	var (
		query  = "select user_id, user_message_box_id, dialog_id1, dialog_id2, dialog_message_id, sender_user_id, peer_type, peer_id, random_id, message_filter_type, message_data, message, mentioned, media_unread, pinned, has_reaction, reaction, reaction_date, reaction_unread, saved_peer_type, saved_peer_id, date2, ttl_period from " + dao.CalcTableName(userId) + " where user_id = ? and (dialog_id1 = ? and dialog_id2 = ?) and user_message_box_id < ? and deleted = 0 and message != '' and message like ? order by user_message_box_id desc limit ?"
		values []dataobject.MessagesDO
	)
	err = dao.db.QueryRowsPartial(ctx, &values, query, userId, dialogId1, dialogId2, userMessageBoxId, q2, limit)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in Search(_), error: %v", err)
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

// SearchGlobal
// select user_id, user_message_box_id, dialog_id1, dialog_id2, dialog_message_id, sender_user_id, peer_type, peer_id, random_id, message_filter_type, message_data, message, mentioned, media_unread, pinned, has_reaction, reaction, reaction_date, reaction_unread, saved_peer_type, saved_peer_id, date2, ttl_period from messages where user_id = :user_id and user_message_box_id < :user_message_box_id and deleted = 0 and message != ” and message like :q2 order by user_message_box_id desc limit :limit
func (dao *MessagesDAO) SearchGlobal(ctx context.Context, userId int64, userMessageBoxId int32, q2 string, limit int32) (rList []dataobject.MessagesDO, err error) {
	var (
		query  = "select user_id, user_message_box_id, dialog_id1, dialog_id2, dialog_message_id, sender_user_id, peer_type, peer_id, random_id, message_filter_type, message_data, message, mentioned, media_unread, pinned, has_reaction, reaction, reaction_date, reaction_unread, saved_peer_type, saved_peer_id, date2, ttl_period from " + dao.CalcTableName(userId) + " where user_id = ? and user_message_box_id < ? and deleted = 0 and message != '' and message like ? order by user_message_box_id desc limit ?"
		values []dataobject.MessagesDO
	)
	err = dao.db.QueryRowsPartial(ctx, &values, query, userId, userMessageBoxId, q2, limit)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SearchGlobal(_), error: %v", err)
		return
	}

	rList = values

	return
}

// SearchGlobalWithCB
// select user_id, user_message_box_id, dialog_id1, dialog_id2, dialog_message_id, sender_user_id, peer_type, peer_id, random_id, message_filter_type, message_data, message, mentioned, media_unread, pinned, has_reaction, reaction, reaction_date, reaction_unread, saved_peer_type, saved_peer_id, date2, ttl_period from messages where user_id = :user_id and user_message_box_id < :user_message_box_id and deleted = 0 and message != ” and message like :q2 order by user_message_box_id desc limit :limit
func (dao *MessagesDAO) SearchGlobalWithCB(ctx context.Context, userId int64, userMessageBoxId int32, q2 string, limit int32, cb func(sz, i int, v *dataobject.MessagesDO)) (rList []dataobject.MessagesDO, err error) {
	var (
		query  = "select user_id, user_message_box_id, dialog_id1, dialog_id2, dialog_message_id, sender_user_id, peer_type, peer_id, random_id, message_filter_type, message_data, message, mentioned, media_unread, pinned, has_reaction, reaction, reaction_date, reaction_unread, saved_peer_type, saved_peer_id, date2, ttl_period from " + dao.CalcTableName(userId) + " where user_id = ? and user_message_box_id < ? and deleted = 0 and message != '' and message like ? order by user_message_box_id desc limit ?"
		values []dataobject.MessagesDO
	)
	err = dao.db.QueryRowsPartial(ctx, &values, query, userId, userMessageBoxId, q2, limit)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SearchGlobal(_), error: %v", err)
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

// SelectBackwardUnreadMentionsByOffsetIdLimit
// select user_id, user_message_box_id, dialog_id1, dialog_id2, dialog_message_id, sender_user_id, peer_type, peer_id, random_id, message_filter_type, message_data, message, mentioned, media_unread, pinned, has_reaction, reaction, reaction_date, reaction_unread, saved_peer_type, saved_peer_id, date2, ttl_period from messages where user_id = :user_id and (dialog_id1 = :dialog_id1 and dialog_id2 = :dialog_id2) and user_message_box_id < :user_message_box_id and mentioned = 1 and media_unread = 1 and deleted = 0 order by user_message_box_id desc limit :limit
func (dao *MessagesDAO) SelectBackwardUnreadMentionsByOffsetIdLimit(ctx context.Context, userId int64, dialogId1 int64, dialogId2 int64, userMessageBoxId int32, limit int32) (rList []dataobject.MessagesDO, err error) {
	var (
		query  = "select user_id, user_message_box_id, dialog_id1, dialog_id2, dialog_message_id, sender_user_id, peer_type, peer_id, random_id, message_filter_type, message_data, message, mentioned, media_unread, pinned, has_reaction, reaction, reaction_date, reaction_unread, saved_peer_type, saved_peer_id, date2, ttl_period from " + dao.CalcTableName(userId) + " where user_id = ? and (dialog_id1 = ? and dialog_id2 = ?) and user_message_box_id < ? and mentioned = 1 and media_unread = 1 and deleted = 0 order by user_message_box_id desc limit ?"
		values []dataobject.MessagesDO
	)
	err = dao.db.QueryRowsPartial(ctx, &values, query, userId, dialogId1, dialogId2, userMessageBoxId, limit)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectBackwardUnreadMentionsByOffsetIdLimit(_), error: %v", err)
		return
	}

	rList = values

	return
}

// SelectBackwardUnreadMentionsByOffsetIdLimitWithCB
// select user_id, user_message_box_id, dialog_id1, dialog_id2, dialog_message_id, sender_user_id, peer_type, peer_id, random_id, message_filter_type, message_data, message, mentioned, media_unread, pinned, has_reaction, reaction, reaction_date, reaction_unread, saved_peer_type, saved_peer_id, date2, ttl_period from messages where user_id = :user_id and (dialog_id1 = :dialog_id1 and dialog_id2 = :dialog_id2) and user_message_box_id < :user_message_box_id and mentioned = 1 and media_unread = 1 and deleted = 0 order by user_message_box_id desc limit :limit
func (dao *MessagesDAO) SelectBackwardUnreadMentionsByOffsetIdLimitWithCB(ctx context.Context, userId int64, dialogId1 int64, dialogId2 int64, userMessageBoxId int32, limit int32, cb func(sz, i int, v *dataobject.MessagesDO)) (rList []dataobject.MessagesDO, err error) {
	var (
		query  = "select user_id, user_message_box_id, dialog_id1, dialog_id2, dialog_message_id, sender_user_id, peer_type, peer_id, random_id, message_filter_type, message_data, message, mentioned, media_unread, pinned, has_reaction, reaction, reaction_date, reaction_unread, saved_peer_type, saved_peer_id, date2, ttl_period from " + dao.CalcTableName(userId) + " where user_id = ? and (dialog_id1 = ? and dialog_id2 = ?) and user_message_box_id < ? and mentioned = 1 and media_unread = 1 and deleted = 0 order by user_message_box_id desc limit ?"
		values []dataobject.MessagesDO
	)
	err = dao.db.QueryRowsPartial(ctx, &values, query, userId, dialogId1, dialogId2, userMessageBoxId, limit)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectBackwardUnreadMentionsByOffsetIdLimit(_), error: %v", err)
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

// SelectForwardUnreadMentionsByOffsetIdLimit
// select user_id, user_message_box_id, dialog_id1, dialog_id2, dialog_message_id, sender_user_id, peer_type, peer_id, random_id, message_filter_type, message_data, message, mentioned, media_unread, pinned, has_reaction, reaction, reaction_date, reaction_unread, saved_peer_type, saved_peer_id, date2, ttl_period from messages where user_id = :user_id and (dialog_id1 = :dialog_id1 and dialog_id2 = :dialog_id2) and user_message_box_id >= :user_message_box_id and mentioned = 1 and media_unread = 1 and deleted = 0 order by user_message_box_id asc limit :limit
func (dao *MessagesDAO) SelectForwardUnreadMentionsByOffsetIdLimit(ctx context.Context, userId int64, dialogId1 int64, dialogId2 int64, userMessageBoxId int32, limit int32) (rList []dataobject.MessagesDO, err error) {
	var (
		query  = "select user_id, user_message_box_id, dialog_id1, dialog_id2, dialog_message_id, sender_user_id, peer_type, peer_id, random_id, message_filter_type, message_data, message, mentioned, media_unread, pinned, has_reaction, reaction, reaction_date, reaction_unread, saved_peer_type, saved_peer_id, date2, ttl_period from " + dao.CalcTableName(userId) + " where user_id = ? and (dialog_id1 = ? and dialog_id2 = ?) and user_message_box_id >= ? and mentioned = 1 and media_unread = 1 and deleted = 0 order by user_message_box_id asc limit ?"
		values []dataobject.MessagesDO
	)
	err = dao.db.QueryRowsPartial(ctx, &values, query, userId, dialogId1, dialogId2, userMessageBoxId, limit)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectForwardUnreadMentionsByOffsetIdLimit(_), error: %v", err)
		return
	}

	rList = values

	return
}

// SelectForwardUnreadMentionsByOffsetIdLimitWithCB
// select user_id, user_message_box_id, dialog_id1, dialog_id2, dialog_message_id, sender_user_id, peer_type, peer_id, random_id, message_filter_type, message_data, message, mentioned, media_unread, pinned, has_reaction, reaction, reaction_date, reaction_unread, saved_peer_type, saved_peer_id, date2, ttl_period from messages where user_id = :user_id and (dialog_id1 = :dialog_id1 and dialog_id2 = :dialog_id2) and user_message_box_id >= :user_message_box_id and mentioned = 1 and media_unread = 1 and deleted = 0 order by user_message_box_id asc limit :limit
func (dao *MessagesDAO) SelectForwardUnreadMentionsByOffsetIdLimitWithCB(ctx context.Context, userId int64, dialogId1 int64, dialogId2 int64, userMessageBoxId int32, limit int32, cb func(sz, i int, v *dataobject.MessagesDO)) (rList []dataobject.MessagesDO, err error) {
	var (
		query  = "select user_id, user_message_box_id, dialog_id1, dialog_id2, dialog_message_id, sender_user_id, peer_type, peer_id, random_id, message_filter_type, message_data, message, mentioned, media_unread, pinned, has_reaction, reaction, reaction_date, reaction_unread, saved_peer_type, saved_peer_id, date2, ttl_period from " + dao.CalcTableName(userId) + " where user_id = ? and (dialog_id1 = ? and dialog_id2 = ?) and user_message_box_id >= ? and mentioned = 1 and media_unread = 1 and deleted = 0 order by user_message_box_id asc limit ?"
		values []dataobject.MessagesDO
	)
	err = dao.db.QueryRowsPartial(ctx, &values, query, userId, dialogId1, dialogId2, userMessageBoxId, limit)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectForwardUnreadMentionsByOffsetIdLimit(_), error: %v", err)
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

// SelectPinnedList
// select user_id, user_message_box_id, dialog_id1, dialog_id2, dialog_message_id, sender_user_id, peer_type, peer_id, random_id, message_filter_type, message_data, message, mentioned, media_unread, pinned, has_reaction, reaction, reaction_date, reaction_unread, saved_peer_type, saved_peer_id, date2, ttl_period from messages where user_id = :user_id and (dialog_id1 = :dialog_id1 and dialog_id2 = :dialog_id2) and pinned = 1 and deleted = 0 order by user_message_box_id desc
func (dao *MessagesDAO) SelectPinnedList(ctx context.Context, userId int64, dialogId1 int64, dialogId2 int64) (rList []dataobject.MessagesDO, err error) {
	var (
		query  = "select user_id, user_message_box_id, dialog_id1, dialog_id2, dialog_message_id, sender_user_id, peer_type, peer_id, random_id, message_filter_type, message_data, message, mentioned, media_unread, pinned, has_reaction, reaction, reaction_date, reaction_unread, saved_peer_type, saved_peer_id, date2, ttl_period from " + dao.CalcTableName(userId) + " where user_id = ? and (dialog_id1 = ? and dialog_id2 = ?) and pinned = 1 and deleted = 0 order by user_message_box_id desc"
		values []dataobject.MessagesDO
	)
	err = dao.db.QueryRowsPartial(ctx, &values, query, userId, dialogId1, dialogId2)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectPinnedList(_), error: %v", err)
		return
	}

	rList = values

	return
}

// SelectPinnedListWithCB
// select user_id, user_message_box_id, dialog_id1, dialog_id2, dialog_message_id, sender_user_id, peer_type, peer_id, random_id, message_filter_type, message_data, message, mentioned, media_unread, pinned, has_reaction, reaction, reaction_date, reaction_unread, saved_peer_type, saved_peer_id, date2, ttl_period from messages where user_id = :user_id and (dialog_id1 = :dialog_id1 and dialog_id2 = :dialog_id2) and pinned = 1 and deleted = 0 order by user_message_box_id desc
func (dao *MessagesDAO) SelectPinnedListWithCB(ctx context.Context, userId int64, dialogId1 int64, dialogId2 int64, cb func(sz, i int, v *dataobject.MessagesDO)) (rList []dataobject.MessagesDO, err error) {
	var (
		query  = "select user_id, user_message_box_id, dialog_id1, dialog_id2, dialog_message_id, sender_user_id, peer_type, peer_id, random_id, message_filter_type, message_data, message, mentioned, media_unread, pinned, has_reaction, reaction, reaction_date, reaction_unread, saved_peer_type, saved_peer_id, date2, ttl_period from " + dao.CalcTableName(userId) + " where user_id = ? and (dialog_id1 = ? and dialog_id2 = ?) and pinned = 1 and deleted = 0 order by user_message_box_id desc"
		values []dataobject.MessagesDO
	)
	err = dao.db.QueryRowsPartial(ctx, &values, query, userId, dialogId1, dialogId2)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectPinnedList(_), error: %v", err)
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

// SelectLastTwoPinnedList
// select user_message_box_id from messages where user_id = :user_id and (dialog_id1 = :dialog_id1 and dialog_id2 = :dialog_id2) and pinned = 1 and deleted = 0 order by user_message_box_id desc limit 2
func (dao *MessagesDAO) SelectLastTwoPinnedList(ctx context.Context, userId int64, dialogId1 int64, dialogId2 int64) (rList []int32, err error) {
	var query = "select user_message_box_id from " + dao.CalcTableName(userId) + " where user_id = ? and (dialog_id1 = ? and dialog_id2 = ?) and pinned = 1 and deleted = 0 order by user_message_box_id desc limit 2"
	err = dao.db.QueryRowsPartial(ctx, &rList, query, userId, dialogId1, dialogId2)

	if err != nil {
		logx.WithContext(ctx).Errorf("select in SelectLastTwoPinnedList(_), error: %v", err)
	}

	return
}

// SelectLastTwoPinnedListWithCB
// select user_message_box_id from messages where user_id = :user_id and (dialog_id1 = :dialog_id1 and dialog_id2 = :dialog_id2) and pinned = 1 and deleted = 0 order by user_message_box_id desc limit 2
func (dao *MessagesDAO) SelectLastTwoPinnedListWithCB(ctx context.Context, userId int64, dialogId1 int64, dialogId2 int64, cb func(sz, i int, v int32)) (rList []int32, err error) {
	var query = "select user_message_box_id from " + dao.CalcTableName(userId) + " where user_id = ? and (dialog_id1 = ? and dialog_id2 = ?) and pinned = 1 and deleted = 0 order by user_message_box_id desc limit 2"
	err = dao.db.QueryRowsPartial(ctx, &rList, query, userId, dialogId1, dialogId2)

	if err != nil {
		logx.WithContext(ctx).Errorf("select in SelectLastTwoPinnedList(_), error: %v", err)
	}

	if cb != nil {
		sz := len(rList)
		for i := 0; i < sz; i++ {
			cb(sz, i, rList[i])
		}
	}

	return
}

// UpdatePinned
// update messages set pinned = :pinned where user_id = :user_id and user_message_box_id = :user_message_box_id
func (dao *MessagesDAO) UpdatePinned(ctx context.Context, pinned bool, userId int64, userMessageBoxId int32) (rowsAffected int64, err error) {
	var (
		query   = "update " + dao.CalcTableName(userId) + " set pinned = ? where user_id = ? and user_message_box_id = ?"
		rResult sql.Result
	)

	rResult, err = dao.db.Exec(ctx, query, pinned, userId, userMessageBoxId)

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
// update messages set pinned = :pinned where user_id = :user_id and user_message_box_id = :user_message_box_id
func (dao *MessagesDAO) UpdatePinnedTx(tx *sqlx.Tx, pinned bool, userId int64, userMessageBoxId int32) (rowsAffected int64, err error) {
	var (
		query   = "update " + dao.CalcTableName(userId) + " set pinned = ? where user_id = ? and user_message_box_id = ?"
		rResult sql.Result
	)
	rResult, err = tx.Exec(query, pinned, userId, userMessageBoxId)

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

// SelectPinnedMessageIdList
// select user_message_box_id from messages where user_id = :user_id and (dialog_id1 = :dialog_id1 and dialog_id2 = :dialog_id2) and pinned = 1 and deleted = 0 order by user_message_box_id desc
func (dao *MessagesDAO) SelectPinnedMessageIdList(ctx context.Context, userId int64, dialogId1 int64, dialogId2 int64) (rList []int32, err error) {
	var query = "select user_message_box_id from " + dao.CalcTableName(userId) + " where user_id = ? and (dialog_id1 = ? and dialog_id2 = ?) and pinned = 1 and deleted = 0 order by user_message_box_id desc"
	err = dao.db.QueryRowsPartial(ctx, &rList, query, userId, dialogId1, dialogId2)

	if err != nil {
		logx.WithContext(ctx).Errorf("select in SelectPinnedMessageIdList(_), error: %v", err)
	}

	return
}

// SelectPinnedMessageIdListWithCB
// select user_message_box_id from messages where user_id = :user_id and (dialog_id1 = :dialog_id1 and dialog_id2 = :dialog_id2) and pinned = 1 and deleted = 0 order by user_message_box_id desc
func (dao *MessagesDAO) SelectPinnedMessageIdListWithCB(ctx context.Context, userId int64, dialogId1 int64, dialogId2 int64, cb func(sz, i int, v int32)) (rList []int32, err error) {
	var query = "select user_message_box_id from " + dao.CalcTableName(userId) + " where user_id = ? and (dialog_id1 = ? and dialog_id2 = ?) and pinned = 1 and deleted = 0 order by user_message_box_id desc"
	err = dao.db.QueryRowsPartial(ctx, &rList, query, userId, dialogId1, dialogId2)

	if err != nil {
		logx.WithContext(ctx).Errorf("select in SelectPinnedMessageIdList(_), error: %v", err)
	}

	if cb != nil {
		sz := len(rList)
		for i := 0; i < sz; i++ {
			cb(sz, i, rList[i])
		}
	}

	return
}

// UpdateUnPinnedByIdList
// update messages set pinned = 0 where user_id = :user_id and user_message_box_id in (:idList)
func (dao *MessagesDAO) UpdateUnPinnedByIdList(ctx context.Context, userId int64, idList []int32) (rowsAffected int64, err error) {
	var (
		query   = fmt.Sprintf("update "+dao.CalcTableName(userId)+" set pinned = 0 where user_id = ? and user_message_box_id in (%s)", sqlx.InInt32List(idList))
		rResult sql.Result
	)

	if len(idList) == 0 {
		return
	}

	rResult, err = dao.db.Exec(ctx, query, userId)

	if err != nil {
		logx.WithContext(ctx).Errorf("exec in UpdateUnPinnedByIdList(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		logx.WithContext(ctx).Errorf("rowsAffected in UpdateUnPinnedByIdList(_), error: %v", err)
	}

	return
}

// UpdateUnPinnedByIdListTx
// update messages set pinned = 0 where user_id = :user_id and user_message_box_id in (:idList)
func (dao *MessagesDAO) UpdateUnPinnedByIdListTx(tx *sqlx.Tx, userId int64, idList []int32) (rowsAffected int64, err error) {
	var (
		query   = fmt.Sprintf("update "+dao.CalcTableName(userId)+" set pinned = 0 where user_id = ? and user_message_box_id in (%s)", sqlx.InInt32List(idList))
		rResult sql.Result
	)

	if len(idList) == 0 {
		return
	}

	rResult, err = tx.Exec(query, userId)

	if err != nil {
		logx.WithContext(tx.Context()).Errorf("exec in UpdateUnPinnedByIdList(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		logx.WithContext(tx.Context()).Errorf("rowsAffected in UpdateUnPinnedByIdList(_), error: %v", err)
	}

	return
}

// UpdateEditMessage
// update messages set message_data = :message_data, message = :message where user_id = :user_id and user_message_box_id = :user_message_box_id
func (dao *MessagesDAO) UpdateEditMessage(ctx context.Context, messageData string, message string, userId int64, userMessageBoxId int32) (rowsAffected int64, err error) {
	var (
		query   = "update " + dao.CalcTableName(userId) + " set message_data = ?, message = ? where user_id = ? and user_message_box_id = ?"
		rResult sql.Result
	)

	rResult, err = dao.db.Exec(ctx, query, messageData, message, userId, userMessageBoxId)

	if err != nil {
		logx.WithContext(ctx).Errorf("exec in UpdateEditMessage(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		logx.WithContext(ctx).Errorf("rowsAffected in UpdateEditMessage(_), error: %v", err)
	}

	return
}

// UpdateEditMessageTx
// update messages set message_data = :message_data, message = :message where user_id = :user_id and user_message_box_id = :user_message_box_id
func (dao *MessagesDAO) UpdateEditMessageTx(tx *sqlx.Tx, messageData string, message string, userId int64, userMessageBoxId int32) (rowsAffected int64, err error) {
	var (
		query   = "update " + dao.CalcTableName(userId) + " set message_data = ?, message = ? where user_id = ? and user_message_box_id = ?"
		rResult sql.Result
	)
	rResult, err = tx.Exec(query, messageData, message, userId, userMessageBoxId)

	if err != nil {
		logx.WithContext(tx.Context()).Errorf("exec in UpdateEditMessage(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		logx.WithContext(tx.Context()).Errorf("rowsAffected in UpdateEditMessage(_), error: %v", err)
	}

	return
}

// UpdateCustomMap
// update messages set %s where user_id = :user_id and user_message_box_id = :user_message_box_id
func (dao *MessagesDAO) UpdateCustomMap(ctx context.Context, cMap map[string]interface{}, userId int64, userMessageBoxId int32) (rowsAffected int64, err error) {
	names := make([]string, 0, len(cMap))
	aValues := make([]interface{}, 0, len(cMap))
	for k, v := range cMap {
		names = append(names, k+" = ?")
		aValues = append(aValues, v)
	}

	var (
		query   = fmt.Sprintf("update "+dao.CalcTableName(userId)+" set %s where user_id = ? and user_message_box_id = ?", strings.Join(names, ", "))
		rResult sql.Result
	)

	aValues = append(aValues, userId)
	aValues = append(aValues, userMessageBoxId)

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
// update messages set %s where user_id = :user_id and user_message_box_id = :user_message_box_id
func (dao *MessagesDAO) UpdateCustomMapTx(tx *sqlx.Tx, cMap map[string]interface{}, userId int64, userMessageBoxId int32) (rowsAffected int64, err error) {
	names := make([]string, 0, len(cMap))
	aValues := make([]interface{}, 0, len(cMap))
	for k, v := range cMap {
		names = append(names, k+" = ?")
		aValues = append(aValues, v)
	}

	var (
		query   = fmt.Sprintf("update "+dao.CalcTableName(userId)+" set %s where user_id = ? and user_message_box_id = ?", strings.Join(names, ", "))
		rResult sql.Result
	)

	aValues = append(aValues, userId)
	aValues = append(aValues, userMessageBoxId)

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

// SelectBackwardBySendUserIdOffsetIdLimit
// select user_id, user_message_box_id, dialog_id1, dialog_id2, dialog_message_id, sender_user_id, peer_type, peer_id, random_id, message_filter_type, message_data, message, mentioned, media_unread, pinned, has_reaction, reaction, reaction_date, reaction_unread, saved_peer_type, saved_peer_id, date2, ttl_period from messages where user_id = :user_id and (dialog_id1 = :dialog_id1 and dialog_id2 = :dialog_id2) and sender_user_id = :sender_user_id and user_message_box_id < :user_message_box_id and deleted = 0 order by user_message_box_id desc limit :limit
func (dao *MessagesDAO) SelectBackwardBySendUserIdOffsetIdLimit(ctx context.Context, userId int64, dialogId1 int64, dialogId2 int64, senderUserId int64, userMessageBoxId int32, limit int32) (rList []dataobject.MessagesDO, err error) {
	var (
		query  = "select user_id, user_message_box_id, dialog_id1, dialog_id2, dialog_message_id, sender_user_id, peer_type, peer_id, random_id, message_filter_type, message_data, message, mentioned, media_unread, pinned, has_reaction, reaction, reaction_date, reaction_unread, saved_peer_type, saved_peer_id, date2, ttl_period from " + dao.CalcTableName(userId) + " where user_id = ? and (dialog_id1 = ? and dialog_id2 = ?) and sender_user_id = ? and user_message_box_id < ? and deleted = 0 order by user_message_box_id desc limit ?"
		values []dataobject.MessagesDO
	)
	err = dao.db.QueryRowsPartial(ctx, &values, query, userId, dialogId1, dialogId2, senderUserId, userMessageBoxId, limit)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectBackwardBySendUserIdOffsetIdLimit(_), error: %v", err)
		return
	}

	rList = values

	return
}

// SelectBackwardBySendUserIdOffsetIdLimitWithCB
// select user_id, user_message_box_id, dialog_id1, dialog_id2, dialog_message_id, sender_user_id, peer_type, peer_id, random_id, message_filter_type, message_data, message, mentioned, media_unread, pinned, has_reaction, reaction, reaction_date, reaction_unread, saved_peer_type, saved_peer_id, date2, ttl_period from messages where user_id = :user_id and (dialog_id1 = :dialog_id1 and dialog_id2 = :dialog_id2) and sender_user_id = :sender_user_id and user_message_box_id < :user_message_box_id and deleted = 0 order by user_message_box_id desc limit :limit
func (dao *MessagesDAO) SelectBackwardBySendUserIdOffsetIdLimitWithCB(ctx context.Context, userId int64, dialogId1 int64, dialogId2 int64, senderUserId int64, userMessageBoxId int32, limit int32, cb func(sz, i int, v *dataobject.MessagesDO)) (rList []dataobject.MessagesDO, err error) {
	var (
		query  = "select user_id, user_message_box_id, dialog_id1, dialog_id2, dialog_message_id, sender_user_id, peer_type, peer_id, random_id, message_filter_type, message_data, message, mentioned, media_unread, pinned, has_reaction, reaction, reaction_date, reaction_unread, saved_peer_type, saved_peer_id, date2, ttl_period from " + dao.CalcTableName(userId) + " where user_id = ? and (dialog_id1 = ? and dialog_id2 = ?) and sender_user_id = ? and user_message_box_id < ? and deleted = 0 order by user_message_box_id desc limit ?"
		values []dataobject.MessagesDO
	)
	err = dao.db.QueryRowsPartial(ctx, &values, query, userId, dialogId1, dialogId2, senderUserId, userMessageBoxId, limit)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectBackwardBySendUserIdOffsetIdLimit(_), error: %v", err)
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

// SelectBackwardSavedByOffsetIdLimit
// select user_id, user_message_box_id, dialog_id1, dialog_id2, dialog_message_id, sender_user_id, peer_type, peer_id, random_id, message_filter_type, message_data, message, mentioned, media_unread, pinned, has_reaction, reaction, reaction_date, reaction_unread, saved_peer_type, saved_peer_id, date2, ttl_period from messages where user_id = :user_id and (saved_peer_type = :saved_peer_type and saved_peer_id = :saved_peer_id) and user_message_box_id < :user_message_box_id and deleted = 0 order by user_message_box_id desc limit :limit
func (dao *MessagesDAO) SelectBackwardSavedByOffsetIdLimit(ctx context.Context, userId int64, savedPeerType int32, savedPeerId int64, userMessageBoxId int32, limit int32) (rList []dataobject.MessagesDO, err error) {
	var (
		query  = "select user_id, user_message_box_id, dialog_id1, dialog_id2, dialog_message_id, sender_user_id, peer_type, peer_id, random_id, message_filter_type, message_data, message, mentioned, media_unread, pinned, has_reaction, reaction, reaction_date, reaction_unread, saved_peer_type, saved_peer_id, date2, ttl_period from " + dao.CalcTableName(userId) + " where user_id = ? and (saved_peer_type = ? and saved_peer_id = ?) and user_message_box_id < ? and deleted = 0 order by user_message_box_id desc limit ?"
		values []dataobject.MessagesDO
	)
	err = dao.db.QueryRowsPartial(ctx, &values, query, userId, savedPeerType, savedPeerId, userMessageBoxId, limit)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectBackwardSavedByOffsetIdLimit(_), error: %v", err)
		return
	}

	rList = values

	return
}

// SelectBackwardSavedByOffsetIdLimitWithCB
// select user_id, user_message_box_id, dialog_id1, dialog_id2, dialog_message_id, sender_user_id, peer_type, peer_id, random_id, message_filter_type, message_data, message, mentioned, media_unread, pinned, has_reaction, reaction, reaction_date, reaction_unread, saved_peer_type, saved_peer_id, date2, ttl_period from messages where user_id = :user_id and (saved_peer_type = :saved_peer_type and saved_peer_id = :saved_peer_id) and user_message_box_id < :user_message_box_id and deleted = 0 order by user_message_box_id desc limit :limit
func (dao *MessagesDAO) SelectBackwardSavedByOffsetIdLimitWithCB(ctx context.Context, userId int64, savedPeerType int32, savedPeerId int64, userMessageBoxId int32, limit int32, cb func(sz, i int, v *dataobject.MessagesDO)) (rList []dataobject.MessagesDO, err error) {
	var (
		query  = "select user_id, user_message_box_id, dialog_id1, dialog_id2, dialog_message_id, sender_user_id, peer_type, peer_id, random_id, message_filter_type, message_data, message, mentioned, media_unread, pinned, has_reaction, reaction, reaction_date, reaction_unread, saved_peer_type, saved_peer_id, date2, ttl_period from " + dao.CalcTableName(userId) + " where user_id = ? and (saved_peer_type = ? and saved_peer_id = ?) and user_message_box_id < ? and deleted = 0 order by user_message_box_id desc limit ?"
		values []dataobject.MessagesDO
	)
	err = dao.db.QueryRowsPartial(ctx, &values, query, userId, savedPeerType, savedPeerId, userMessageBoxId, limit)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectBackwardSavedByOffsetIdLimit(_), error: %v", err)
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

// SelectForwardSavedByOffsetIdLimit
// select user_id, user_message_box_id, dialog_id1, dialog_id2, dialog_message_id, sender_user_id, peer_type, peer_id, random_id, message_filter_type, message_data, message, mentioned, media_unread, pinned, has_reaction, reaction, reaction_date, reaction_unread, saved_peer_type, saved_peer_id, date2, ttl_period from messages where user_id = :user_id and (saved_peer_type = :saved_peer_type and saved_peer_id = :saved_peer_id) and user_message_box_id >= :user_message_box_id and deleted = 0 order by user_message_box_id asc limit :limit
func (dao *MessagesDAO) SelectForwardSavedByOffsetIdLimit(ctx context.Context, userId int64, savedPeerType int32, savedPeerId int64, userMessageBoxId int32, limit int32) (rList []dataobject.MessagesDO, err error) {
	var (
		query  = "select user_id, user_message_box_id, dialog_id1, dialog_id2, dialog_message_id, sender_user_id, peer_type, peer_id, random_id, message_filter_type, message_data, message, mentioned, media_unread, pinned, has_reaction, reaction, reaction_date, reaction_unread, saved_peer_type, saved_peer_id, date2, ttl_period from " + dao.CalcTableName(userId) + " where user_id = ? and (saved_peer_type = ? and saved_peer_id = ?) and user_message_box_id >= ? and deleted = 0 order by user_message_box_id asc limit ?"
		values []dataobject.MessagesDO
	)
	err = dao.db.QueryRowsPartial(ctx, &values, query, userId, savedPeerType, savedPeerId, userMessageBoxId, limit)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectForwardSavedByOffsetIdLimit(_), error: %v", err)
		return
	}

	rList = values

	return
}

// SelectForwardSavedByOffsetIdLimitWithCB
// select user_id, user_message_box_id, dialog_id1, dialog_id2, dialog_message_id, sender_user_id, peer_type, peer_id, random_id, message_filter_type, message_data, message, mentioned, media_unread, pinned, has_reaction, reaction, reaction_date, reaction_unread, saved_peer_type, saved_peer_id, date2, ttl_period from messages where user_id = :user_id and (saved_peer_type = :saved_peer_type and saved_peer_id = :saved_peer_id) and user_message_box_id >= :user_message_box_id and deleted = 0 order by user_message_box_id asc limit :limit
func (dao *MessagesDAO) SelectForwardSavedByOffsetIdLimitWithCB(ctx context.Context, userId int64, savedPeerType int32, savedPeerId int64, userMessageBoxId int32, limit int32, cb func(sz, i int, v *dataobject.MessagesDO)) (rList []dataobject.MessagesDO, err error) {
	var (
		query  = "select user_id, user_message_box_id, dialog_id1, dialog_id2, dialog_message_id, sender_user_id, peer_type, peer_id, random_id, message_filter_type, message_data, message, mentioned, media_unread, pinned, has_reaction, reaction, reaction_date, reaction_unread, saved_peer_type, saved_peer_id, date2, ttl_period from " + dao.CalcTableName(userId) + " where user_id = ? and (saved_peer_type = ? and saved_peer_id = ?) and user_message_box_id >= ? and deleted = 0 order by user_message_box_id asc limit ?"
		values []dataobject.MessagesDO
	)
	err = dao.db.QueryRowsPartial(ctx, &values, query, userId, savedPeerType, savedPeerId, userMessageBoxId, limit)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectForwardSavedByOffsetIdLimit(_), error: %v", err)
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

// SelectBackwardSavedByOffsetDateLimit
// select user_id, user_message_box_id, dialog_id1, dialog_id2, dialog_message_id, sender_user_id, peer_type, peer_id, random_id, message_filter_type, message_data, message, mentioned, media_unread, pinned, has_reaction, reaction, reaction_date, reaction_unread, saved_peer_type, saved_peer_id, date2, ttl_period from messages where user_id = :user_id and (saved_peer_type = :saved_peer_type and saved_peer_id = :saved_peer_id) and date2 < :date2 and deleted = 0 order by user_message_box_id desc limit :limit
func (dao *MessagesDAO) SelectBackwardSavedByOffsetDateLimit(ctx context.Context, userId int64, savedPeerType int32, savedPeerId int64, date2 int64, limit int32) (rList []dataobject.MessagesDO, err error) {
	var (
		query  = "select user_id, user_message_box_id, dialog_id1, dialog_id2, dialog_message_id, sender_user_id, peer_type, peer_id, random_id, message_filter_type, message_data, message, mentioned, media_unread, pinned, has_reaction, reaction, reaction_date, reaction_unread, saved_peer_type, saved_peer_id, date2, ttl_period from " + dao.CalcTableName(userId) + " where user_id = ? and (saved_peer_type = ? and saved_peer_id = ?) and date2 < ? and deleted = 0 order by user_message_box_id desc limit ?"
		values []dataobject.MessagesDO
	)
	err = dao.db.QueryRowsPartial(ctx, &values, query, userId, savedPeerType, savedPeerId, date2, limit)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectBackwardSavedByOffsetDateLimit(_), error: %v", err)
		return
	}

	rList = values

	return
}

// SelectBackwardSavedByOffsetDateLimitWithCB
// select user_id, user_message_box_id, dialog_id1, dialog_id2, dialog_message_id, sender_user_id, peer_type, peer_id, random_id, message_filter_type, message_data, message, mentioned, media_unread, pinned, has_reaction, reaction, reaction_date, reaction_unread, saved_peer_type, saved_peer_id, date2, ttl_period from messages where user_id = :user_id and (saved_peer_type = :saved_peer_type and saved_peer_id = :saved_peer_id) and date2 < :date2 and deleted = 0 order by user_message_box_id desc limit :limit
func (dao *MessagesDAO) SelectBackwardSavedByOffsetDateLimitWithCB(ctx context.Context, userId int64, savedPeerType int32, savedPeerId int64, date2 int64, limit int32, cb func(sz, i int, v *dataobject.MessagesDO)) (rList []dataobject.MessagesDO, err error) {
	var (
		query  = "select user_id, user_message_box_id, dialog_id1, dialog_id2, dialog_message_id, sender_user_id, peer_type, peer_id, random_id, message_filter_type, message_data, message, mentioned, media_unread, pinned, has_reaction, reaction, reaction_date, reaction_unread, saved_peer_type, saved_peer_id, date2, ttl_period from " + dao.CalcTableName(userId) + " where user_id = ? and (saved_peer_type = ? and saved_peer_id = ?) and date2 < ? and deleted = 0 order by user_message_box_id desc limit ?"
		values []dataobject.MessagesDO
	)
	err = dao.db.QueryRowsPartial(ctx, &values, query, userId, savedPeerType, savedPeerId, date2, limit)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectBackwardSavedByOffsetDateLimit(_), error: %v", err)
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

// SelectForwardSavedByOffsetDateLimit
// select user_id, user_message_box_id, dialog_id1, dialog_id2, dialog_message_id, sender_user_id, peer_type, peer_id, random_id, message_filter_type, message_data, message, mentioned, media_unread, pinned, has_reaction, reaction, reaction_date, reaction_unread, saved_peer_type, saved_peer_id, date2, ttl_period from messages where user_id = :user_id and (saved_peer_type = :saved_peer_type and saved_peer_id = :saved_peer_id) and date2 >= :date2 and deleted = 0 order by user_message_box_id asc limit :limit
func (dao *MessagesDAO) SelectForwardSavedByOffsetDateLimit(ctx context.Context, userId int64, savedPeerType int32, savedPeerId int64, date2 int64, limit int32) (rList []dataobject.MessagesDO, err error) {
	var (
		query  = "select user_id, user_message_box_id, dialog_id1, dialog_id2, dialog_message_id, sender_user_id, peer_type, peer_id, random_id, message_filter_type, message_data, message, mentioned, media_unread, pinned, has_reaction, reaction, reaction_date, reaction_unread, saved_peer_type, saved_peer_id, date2, ttl_period from " + dao.CalcTableName(userId) + " where user_id = ? and (saved_peer_type = ? and saved_peer_id = ?) and date2 >= ? and deleted = 0 order by user_message_box_id asc limit ?"
		values []dataobject.MessagesDO
	)
	err = dao.db.QueryRowsPartial(ctx, &values, query, userId, savedPeerType, savedPeerId, date2, limit)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectForwardSavedByOffsetDateLimit(_), error: %v", err)
		return
	}

	rList = values

	return
}

// SelectForwardSavedByOffsetDateLimitWithCB
// select user_id, user_message_box_id, dialog_id1, dialog_id2, dialog_message_id, sender_user_id, peer_type, peer_id, random_id, message_filter_type, message_data, message, mentioned, media_unread, pinned, has_reaction, reaction, reaction_date, reaction_unread, saved_peer_type, saved_peer_id, date2, ttl_period from messages where user_id = :user_id and (saved_peer_type = :saved_peer_type and saved_peer_id = :saved_peer_id) and date2 >= :date2 and deleted = 0 order by user_message_box_id asc limit :limit
func (dao *MessagesDAO) SelectForwardSavedByOffsetDateLimitWithCB(ctx context.Context, userId int64, savedPeerType int32, savedPeerId int64, date2 int64, limit int32, cb func(sz, i int, v *dataobject.MessagesDO)) (rList []dataobject.MessagesDO, err error) {
	var (
		query  = "select user_id, user_message_box_id, dialog_id1, dialog_id2, dialog_message_id, sender_user_id, peer_type, peer_id, random_id, message_filter_type, message_data, message, mentioned, media_unread, pinned, has_reaction, reaction, reaction_date, reaction_unread, saved_peer_type, saved_peer_id, date2, ttl_period from " + dao.CalcTableName(userId) + " where user_id = ? and (saved_peer_type = ? and saved_peer_id = ?) and date2 >= ? and deleted = 0 order by user_message_box_id asc limit ?"
		values []dataobject.MessagesDO
	)
	err = dao.db.QueryRowsPartial(ctx, &values, query, userId, savedPeerType, savedPeerId, date2, limit)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectForwardSavedByOffsetDateLimit(_), error: %v", err)
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
