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
	bizMessagesModel interface {
		InsertOrReturnId(ctx context.Context, data *Messages) (lastInsertId, rowsAffected int64, err error)
		InsertOrReturnIdTx(tx *sqlx.Tx, data *Messages) (lastInsertId, rowsAffected int64, err error)

		SelectByRandomId(ctx context.Context, senderUserId int64, randomId int64) (*Messages, error)

		SelectByMessageIdList(ctx context.Context, userId int64, idList []int32) ([]Messages, error)
		SelectByMessageIdListWithCB(ctx context.Context, userId int64, idList []int32, cb func(sz, i int, v *Messages)) ([]Messages, error)

		SelectByMessageId(ctx context.Context, userId int64, userMessageBoxId int32) (*Messages, error)

		SelectByMessageDataIdList(ctx context.Context, idList []int64) ([]Messages, error)
		SelectByMessageDataIdListWithCB(ctx context.Context, idList []int64, cb func(sz, i int, v *Messages)) ([]Messages, error)

		SelectByMessageDataId(ctx context.Context, userId int64, dialogMessageId int64) (*Messages, error)

		SelectByMessageDataIdUserIdList(ctx context.Context, dialogMessageId int64, idList []int64) ([]Messages, error)
		SelectByMessageDataIdUserIdListWithCB(ctx context.Context, dialogMessageId int64, idList []int64, cb func(sz, i int, v *Messages)) ([]Messages, error)

		SelectBackwardByOffsetIdLimit(ctx context.Context, userId int64, dialogId1 int64, dialogId2 int64, userMessageBoxId int32, limit int32) ([]Messages, error)
		SelectBackwardByOffsetIdLimitWithCB(ctx context.Context, userId int64, dialogId1 int64, dialogId2 int64, userMessageBoxId int32, limit int32, cb func(sz, i int, v *Messages)) ([]Messages, error)

		SelectForwardByOffsetIdLimit(ctx context.Context, userId int64, dialogId1 int64, dialogId2 int64, userMessageBoxId int32, limit int32) ([]Messages, error)
		SelectForwardByOffsetIdLimitWithCB(ctx context.Context, userId int64, dialogId1 int64, dialogId2 int64, userMessageBoxId int32, limit int32, cb func(sz, i int, v *Messages)) ([]Messages, error)

		SelectBackwardByOffsetDateLimit(ctx context.Context, userId int64, dialogId1 int64, dialogId2 int64, date2 int64, limit int32) ([]Messages, error)
		SelectBackwardByOffsetDateLimitWithCB(ctx context.Context, userId int64, dialogId1 int64, dialogId2 int64, date2 int64, limit int32, cb func(sz, i int, v *Messages)) ([]Messages, error)

		SelectForwardByOffsetDateLimit(ctx context.Context, userId int64, dialogId1 int64, dialogId2 int64, date2 int64, limit int32) ([]Messages, error)
		SelectForwardByOffsetDateLimitWithCB(ctx context.Context, userId int64, dialogId1 int64, dialogId2 int64, date2 int64, limit int32, cb func(sz, i int, v *Messages)) ([]Messages, error)

		SelectPeerUserMessageId(ctx context.Context, peerId int64, userId int64, userMessageBoxId int32) (*Messages, error)

		SelectPeerUserMessage(ctx context.Context, peerId int64, userId int64, userMessageBoxId int32) (*Messages, error)

		SelectDialogLastMessageId(ctx context.Context, userId int64, dialogId1 int64, dialogId2 int64) (int32, error)

		SelectDialogLastMessageIdNotIdList(ctx context.Context, userId int64, dialogId1 int64, dialogId2 int64, idList []int32) (int32, error)

		SelectDialogsByMessageIdList(ctx context.Context, userId int64, idList []int32) ([]Messages, error)
		SelectDialogsByMessageIdListWithCB(ctx context.Context, userId int64, idList []int32, cb func(sz, i int, v *Messages)) ([]Messages, error)

		SelectDialogLastMessageList(ctx context.Context, userId int64, dialogId1 int64, dialogId2 int64, limit int32) ([]Messages, error)
		SelectDialogLastMessageListWithCB(ctx context.Context, userId int64, dialogId1 int64, dialogId2 int64, limit int32, cb func(sz, i int, v *Messages)) ([]Messages, error)

		DeleteMessagesByMessageIdList(ctx context.Context, userId int64, idList []int32) (rowsAffected int64, err error)
		DeleteMessagesByMessageIdListTx(tx *sqlx.Tx, userId int64, idList []int32) (rowsAffected int64, err error)

		SelectDialogMessageIdList(ctx context.Context, userId int64, dialogId1 int64, dialogId2 int64) ([]Messages, error)
		SelectDialogMessageIdListWithCB(ctx context.Context, userId int64, dialogId1 int64, dialogId2 int64, cb func(sz, i int, v *Messages)) ([]Messages, error)

		UpdateMediaUnread(ctx context.Context, userId int64, userMessageBoxId int32) (rowsAffected int64, err error)
		UpdateMediaUnreadTx(tx *sqlx.Tx, userId int64, userMessageBoxId int32) (rowsAffected int64, err error)

		UpdateMentionedAndMediaUnread(ctx context.Context, userId int64, userMessageBoxId int32) (rowsAffected int64, err error)
		UpdateMentionedAndMediaUnreadTx(tx *sqlx.Tx, userId int64, userMessageBoxId int32) (rowsAffected int64, err error)

		SelectByMediaType(ctx context.Context, userId int64, dialogId1 int64, dialogId2 int64, messageFilterType int32, userMessageBoxId int32, limit int32) ([]Messages, error)
		SelectByMediaTypeWithCB(ctx context.Context, userId int64, dialogId1 int64, dialogId2 int64, messageFilterType int32, userMessageBoxId int32, limit int32, cb func(sz, i int, v *Messages)) ([]Messages, error)

		SelectPhoneCallList(ctx context.Context, userId int64, messageFilterType int32, userMessageBoxId int32, limit int32) ([]Messages, error)
		SelectPhoneCallListWithCB(ctx context.Context, userId int64, messageFilterType int32, userMessageBoxId int32, limit int32, cb func(sz, i int, v *Messages)) ([]Messages, error)

		Search(ctx context.Context, userId int64, dialogId1 int64, dialogId2 int64, userMessageBoxId int32, q2 string, limit int32) ([]Messages, error)
		SearchWithCB(ctx context.Context, userId int64, dialogId1 int64, dialogId2 int64, userMessageBoxId int32, q2 string, limit int32, cb func(sz, i int, v *Messages)) ([]Messages, error)

		SearchGlobal(ctx context.Context, userId int64, userMessageBoxId int32, q2 string, limit int32) ([]Messages, error)
		SearchGlobalWithCB(ctx context.Context, userId int64, userMessageBoxId int32, q2 string, limit int32, cb func(sz, i int, v *Messages)) ([]Messages, error)

		SelectBackwardUnreadMentionsByOffsetIdLimit(ctx context.Context, userId int64, dialogId1 int64, dialogId2 int64, userMessageBoxId int32, limit int32) ([]Messages, error)
		SelectBackwardUnreadMentionsByOffsetIdLimitWithCB(ctx context.Context, userId int64, dialogId1 int64, dialogId2 int64, userMessageBoxId int32, limit int32, cb func(sz, i int, v *Messages)) ([]Messages, error)

		SelectForwardUnreadMentionsByOffsetIdLimit(ctx context.Context, userId int64, dialogId1 int64, dialogId2 int64, userMessageBoxId int32, limit int32) ([]Messages, error)
		SelectForwardUnreadMentionsByOffsetIdLimitWithCB(ctx context.Context, userId int64, dialogId1 int64, dialogId2 int64, userMessageBoxId int32, limit int32, cb func(sz, i int, v *Messages)) ([]Messages, error)

		SelectPinnedList(ctx context.Context, userId int64, dialogId1 int64, dialogId2 int64) ([]Messages, error)
		SelectPinnedListWithCB(ctx context.Context, userId int64, dialogId1 int64, dialogId2 int64, cb func(sz, i int, v *Messages)) ([]Messages, error)

		SelectLastTwoPinnedList(ctx context.Context, userId int64, dialogId1 int64, dialogId2 int64) ([]int32, error)
		SelectLastTwoPinnedListWithCB(ctx context.Context, userId int64, dialogId1 int64, dialogId2 int64, cb func(sz, i int, v int32)) ([]int32, error)

		UpdatePinned(ctx context.Context, pinned bool, userId int64, userMessageBoxId int32) (rowsAffected int64, err error)
		UpdatePinnedTx(tx *sqlx.Tx, pinned bool, userId int64, userMessageBoxId int32) (rowsAffected int64, err error)

		SelectPinnedMessageIdList(ctx context.Context, userId int64, dialogId1 int64, dialogId2 int64) ([]int32, error)
		SelectPinnedMessageIdListWithCB(ctx context.Context, userId int64, dialogId1 int64, dialogId2 int64, cb func(sz, i int, v int32)) ([]int32, error)

		UpdateUnPinnedByIdList(ctx context.Context, userId int64, idList []int32) (rowsAffected int64, err error)
		UpdateUnPinnedByIdListTx(tx *sqlx.Tx, userId int64, idList []int32) (rowsAffected int64, err error)

		UpdateEditMessage(ctx context.Context, messageData string, message string, userId int64, userMessageBoxId int32) (rowsAffected int64, err error)
		UpdateEditMessageTx(tx *sqlx.Tx, messageData string, message string, userId int64, userMessageBoxId int32) (rowsAffected int64, err error)

		UpdateCustomMap(ctx context.Context, cMap map[string]interface{}, userId int64, userMessageBoxId int32) (rowsAffected int64, err error)
		UpdateCustomMapTx(tx *sqlx.Tx, cMap map[string]interface{}, userId int64, userMessageBoxId int32) (rowsAffected int64, err error)

		SelectBackwardBySendUserIdOffsetIdLimit(ctx context.Context, userId int64, dialogId1 int64, dialogId2 int64, senderUserId int64, userMessageBoxId int32, limit int32) ([]Messages, error)
		SelectBackwardBySendUserIdOffsetIdLimitWithCB(ctx context.Context, userId int64, dialogId1 int64, dialogId2 int64, senderUserId int64, userMessageBoxId int32, limit int32, cb func(sz, i int, v *Messages)) ([]Messages, error)

		SelectBackwardSavedByOffsetIdLimit(ctx context.Context, userId int64, savedPeerType int32, savedPeerId int64, userMessageBoxId int32, limit int32) ([]Messages, error)
		SelectBackwardSavedByOffsetIdLimitWithCB(ctx context.Context, userId int64, savedPeerType int32, savedPeerId int64, userMessageBoxId int32, limit int32, cb func(sz, i int, v *Messages)) ([]Messages, error)

		SelectForwardSavedByOffsetIdLimit(ctx context.Context, userId int64, savedPeerType int32, savedPeerId int64, userMessageBoxId int32, limit int32) ([]Messages, error)
		SelectForwardSavedByOffsetIdLimitWithCB(ctx context.Context, userId int64, savedPeerType int32, savedPeerId int64, userMessageBoxId int32, limit int32, cb func(sz, i int, v *Messages)) ([]Messages, error)

		SelectBackwardSavedByOffsetDateLimit(ctx context.Context, userId int64, savedPeerType int32, savedPeerId int64, date2 int64, limit int32) ([]Messages, error)
		SelectBackwardSavedByOffsetDateLimitWithCB(ctx context.Context, userId int64, savedPeerType int32, savedPeerId int64, date2 int64, limit int32, cb func(sz, i int, v *Messages)) ([]Messages, error)

		SelectForwardSavedByOffsetDateLimit(ctx context.Context, userId int64, savedPeerType int32, savedPeerId int64, date2 int64, limit int32) ([]Messages, error)
		SelectForwardSavedByOffsetDateLimitWithCB(ctx context.Context, userId int64, savedPeerType int32, savedPeerId int64, date2 int64, limit int32, cb func(sz, i int, v *Messages)) ([]Messages, error)
	}
)

// InsertOrReturnId
// insert into messages(user_id, user_message_box_id, dialog_id1, dialog_id2, dialog_message_id, sender_user_id, peer_type, peer_id, random_id, message_filter_type, message_data, message, mentioned, media_unread, pinned, saved_peer_type, saved_peer_id, date2, ttl_period) values (:user_id, :user_message_box_id, :dialog_id1, :dialog_id2, :dialog_message_id, :sender_user_id, :peer_type, :peer_id, :random_id, :message_filter_type, :message_data, :message, :mentioned, :media_unread, :pinned, :saved_peer_type, :saved_peer_id, :date2, :ttl_period) on duplicate key update id = last_insert_id(id)
func (m *defaultMessagesModel) InsertOrReturnId(ctx context.Context, data *Messages) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into messages(user_id, user_message_box_id, dialog_id1, dialog_id2, dialog_message_id, sender_user_id, peer_type, peer_id, random_id, message_filter_type, message_data, message, mentioned, media_unread, pinned, saved_peer_type, saved_peer_id, date2, ttl_period) values (:user_id, :user_message_box_id, :dialog_id1, :dialog_id2, :dialog_message_id, :sender_user_id, :peer_type, :peer_id, :random_id, :message_filter_type, :message_data, :message, :mentioned, :media_unread, :pinned, :saved_peer_type, :saved_peer_id, :date2, :ttl_period) on duplicate key update id = last_insert_id(id)"
		r     sql.Result
	)

	r, err = m.db.NamedExec(ctx, query, data)
	if err != nil {
		logx.WithContext(ctx).Errorf("namedExec in InsertOrReturnId(%v), error: %v", data, err)
		return
	}

	lastInsertId, err = r.LastInsertId()
	if err != nil {
		logx.WithContext(ctx).Errorf("lastInsertId in InsertOrReturnId(%v)_error: %v", data, err)
		return
	}
	rowsAffected, err = r.RowsAffected()
	if err != nil {
		logx.WithContext(ctx).Errorf("rowsAffected in InsertOrReturnId(%v)_error: %v", data, err)
	}

	return
}

// InsertOrReturnIdTx
// insert into messages(user_id, user_message_box_id, dialog_id1, dialog_id2, dialog_message_id, sender_user_id, peer_type, peer_id, random_id, message_filter_type, message_data, message, mentioned, media_unread, pinned, saved_peer_type, saved_peer_id, date2, ttl_period) values (:user_id, :user_message_box_id, :dialog_id1, :dialog_id2, :dialog_message_id, :sender_user_id, :peer_type, :peer_id, :random_id, :message_filter_type, :message_data, :message, :mentioned, :media_unread, :pinned, :saved_peer_type, :saved_peer_id, :date2, :ttl_period) on duplicate key update id = last_insert_id(id)
func (m *defaultMessagesModel) InsertOrReturnIdTx(tx *sqlx.Tx, data *Messages) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into messages(user_id, user_message_box_id, dialog_id1, dialog_id2, dialog_message_id, sender_user_id, peer_type, peer_id, random_id, message_filter_type, message_data, message, mentioned, media_unread, pinned, saved_peer_type, saved_peer_id, date2, ttl_period) values (:user_id, :user_message_box_id, :dialog_id1, :dialog_id2, :dialog_message_id, :sender_user_id, :peer_type, :peer_id, :random_id, :message_filter_type, :message_data, :message, :mentioned, :media_unread, :pinned, :saved_peer_type, :saved_peer_id, :date2, :ttl_period) on duplicate key update id = last_insert_id(id)"
		r     sql.Result
	)

	r, err = tx.NamedExec(query, data)
	if err != nil {
		logx.WithContext(tx.Context()).Errorf("namedExec in InsertOrReturnId(%v), error: %v", data, err)
		return
	}

	lastInsertId, err = r.LastInsertId()
	if err != nil {
		logx.WithContext(tx.Context()).Errorf("lastInsertId in InsertOrReturnId(%v)_error: %v", data, err)
		return
	}
	rowsAffected, err = r.RowsAffected()
	if err != nil {
		logx.WithContext(tx.Context()).Errorf("rowsAffected in InsertOrReturnId(%v)_error: %v", data, err)
	}

	return
}

// SelectByRandomId
// select user_id, user_message_box_id, dialog_id1, dialog_id2, dialog_message_id, sender_user_id, peer_type, peer_id, random_id, message_filter_type, message_data, message, mentioned, media_unread, pinned, has_reaction, reaction, reaction_date, reaction_unread, saved_peer_type, saved_peer_id, date2, ttl_period from messages where sender_user_id = :sender_user_id and random_id = :random_id and deleted = 0 limit 1
func (m *defaultMessagesModel) SelectByRandomId(ctx context.Context, senderUserId int64, randomId int64) (rValue *Messages, err error) {
	var (
		query = "select user_id, user_message_box_id, dialog_id1, dialog_id2, dialog_message_id, sender_user_id, peer_type, peer_id, random_id, message_filter_type, message_data, message, mentioned, media_unread, pinned, has_reaction, reaction, reaction_date, reaction_unread, saved_peer_type, saved_peer_id, date2, ttl_period from messages where sender_user_id = ? and random_id = ? and deleted = 0 limit 1"
		do    = &Messages{}
	)
	err = m.db.QueryRowPartial(ctx, do, query, senderUserId, randomId)

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
func (m *defaultMessagesModel) SelectByMessageIdList(ctx context.Context, userId int64, idList []int32) (rList []Messages, err error) {
	var (
		query  = fmt.Sprintf("select user_id, user_message_box_id, dialog_id1, dialog_id2, dialog_message_id, sender_user_id, peer_type, peer_id, random_id, message_filter_type, message_data, message, mentioned, media_unread, pinned, has_reaction, reaction, reaction_date, reaction_unread, saved_peer_type, saved_peer_id, date2, ttl_period from messages where user_id = ? and deleted = 0 and user_message_box_id in (%s) order by user_message_box_id desc", sqlx.InInt32List(idList))
		values []Messages
	)
	if len(idList) == 0 {
		rList = []Messages{}
		return
	}

	err = m.db.QueryRowsPartial(ctx, &values, query, userId)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectByMessageIdList(_), error: %v", err)
		return
	}

	rList = values

	return
}

// SelectByMessageIdListWithCB
// select user_id, user_message_box_id, dialog_id1, dialog_id2, dialog_message_id, sender_user_id, peer_type, peer_id, random_id, message_filter_type, message_data, message, mentioned, media_unread, pinned, has_reaction, reaction, reaction_date, reaction_unread, saved_peer_type, saved_peer_id, date2, ttl_period from messages where user_id = :user_id and deleted = 0 and user_message_box_id in (:idList) order by user_message_box_id desc
func (m *defaultMessagesModel) SelectByMessageIdListWithCB(ctx context.Context, userId int64, idList []int32, cb func(sz, i int, v *Messages)) (rList []Messages, err error) {
	var (
		query  = fmt.Sprintf("select user_id, user_message_box_id, dialog_id1, dialog_id2, dialog_message_id, sender_user_id, peer_type, peer_id, random_id, message_filter_type, message_data, message, mentioned, media_unread, pinned, has_reaction, reaction, reaction_date, reaction_unread, saved_peer_type, saved_peer_id, date2, ttl_period from messages where user_id = ? and deleted = 0 and user_message_box_id in (%s) order by user_message_box_id desc", sqlx.InInt32List(idList))
		values []Messages
	)
	if len(idList) == 0 {
		rList = []Messages{}
		return
	}

	err = m.db.QueryRowsPartial(ctx, &values, query, userId)

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
func (m *defaultMessagesModel) SelectByMessageId(ctx context.Context, userId int64, userMessageBoxId int32) (rValue *Messages, err error) {
	var (
		query = "select user_id, user_message_box_id, dialog_id1, dialog_id2, dialog_message_id, sender_user_id, peer_type, peer_id, random_id, message_filter_type, message_data, message, mentioned, media_unread, pinned, has_reaction, reaction, reaction_date, reaction_unread, saved_peer_type, saved_peer_id, date2, ttl_period from messages where user_id = ? and user_message_box_id = ? and deleted = 0 limit 1"
		do    = &Messages{}
	)
	err = m.db.QueryRowPartial(ctx, do, query, userId, userMessageBoxId)

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
func (m *defaultMessagesModel) SelectByMessageDataIdList(ctx context.Context, idList []int64) (rList []Messages, err error) {
	var (
		query  = fmt.Sprintf("select user_id, user_message_box_id, dialog_id1, dialog_id2, dialog_message_id, sender_user_id, peer_type, peer_id, random_id, message_filter_type, message_data, message, mentioned, media_unread, pinned, has_reaction, reaction, reaction_date, reaction_unread, saved_peer_type, saved_peer_id, date2, ttl_period from messages where deleted = 0 and dialog_message_id in (%s) order by user_message_box_id desc", sqlx.InInt64List(idList))
		values []Messages
	)
	if len(idList) == 0 {
		rList = []Messages{}
		return
	}

	err = m.db.QueryRowsPartial(ctx, &values, query)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectByMessageDataIdList(_), error: %v", err)
		return
	}

	rList = values

	return
}

// SelectByMessageDataIdListWithCB
// select user_id, user_message_box_id, dialog_id1, dialog_id2, dialog_message_id, sender_user_id, peer_type, peer_id, random_id, message_filter_type, message_data, message, mentioned, media_unread, pinned, has_reaction, reaction, reaction_date, reaction_unread, saved_peer_type, saved_peer_id, date2, ttl_period from messages where deleted = 0 and dialog_message_id in (:idList) order by user_message_box_id desc
func (m *defaultMessagesModel) SelectByMessageDataIdListWithCB(ctx context.Context, idList []int64, cb func(sz, i int, v *Messages)) (rList []Messages, err error) {
	var (
		query  = fmt.Sprintf("select user_id, user_message_box_id, dialog_id1, dialog_id2, dialog_message_id, sender_user_id, peer_type, peer_id, random_id, message_filter_type, message_data, message, mentioned, media_unread, pinned, has_reaction, reaction, reaction_date, reaction_unread, saved_peer_type, saved_peer_id, date2, ttl_period from messages where deleted = 0 and dialog_message_id in (%s) order by user_message_box_id desc", sqlx.InInt64List(idList))
		values []Messages
	)
	if len(idList) == 0 {
		rList = []Messages{}
		return
	}

	err = m.db.QueryRowsPartial(ctx, &values, query)

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

// SelectByMessageDataId
// select user_id, user_message_box_id, dialog_id1, dialog_id2, dialog_message_id, sender_user_id, peer_type, peer_id, random_id, message_filter_type, message_data, message, mentioned, media_unread, pinned, has_reaction, reaction, reaction_date, reaction_unread, saved_peer_type, saved_peer_id, date2, ttl_period from messages where user_id = :user_id and dialog_message_id = :dialog_message_id and deleted = 0 limit 1
func (m *defaultMessagesModel) SelectByMessageDataId(ctx context.Context, userId int64, dialogMessageId int64) (rValue *Messages, err error) {
	var (
		query = "select user_id, user_message_box_id, dialog_id1, dialog_id2, dialog_message_id, sender_user_id, peer_type, peer_id, random_id, message_filter_type, message_data, message, mentioned, media_unread, pinned, has_reaction, reaction, reaction_date, reaction_unread, saved_peer_type, saved_peer_id, date2, ttl_period from messages where user_id = ? and dialog_message_id = ? and deleted = 0 limit 1"
		do    = &Messages{}
	)
	err = m.db.QueryRowPartial(ctx, do, query, userId, dialogMessageId)

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

// SelectByMessageDataIdUserIdList
// select user_id, user_message_box_id, dialog_id1, dialog_id2, dialog_message_id, sender_user_id, peer_type, peer_id, random_id, message_filter_type, message_data, message, mentioned, media_unread, pinned, has_reaction, reaction, reaction_date, reaction_unread, saved_peer_type, saved_peer_id, date2, ttl_period from messages where dialog_message_id = :dialog_message_id and user_id in (:idList) and deleted = 0
func (m *defaultMessagesModel) SelectByMessageDataIdUserIdList(ctx context.Context, dialogMessageId int64, idList []int64) (rList []Messages, err error) {
	var (
		query  = fmt.Sprintf("select user_id, user_message_box_id, dialog_id1, dialog_id2, dialog_message_id, sender_user_id, peer_type, peer_id, random_id, message_filter_type, message_data, message, mentioned, media_unread, pinned, has_reaction, reaction, reaction_date, reaction_unread, saved_peer_type, saved_peer_id, date2, ttl_period from messages where dialog_message_id = ? and user_id in (%s) and deleted = 0", sqlx.InInt64List(idList))
		values []Messages
	)
	if len(idList) == 0 {
		rList = []Messages{}
		return
	}

	err = m.db.QueryRowsPartial(ctx, &values, query, dialogMessageId)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectByMessageDataIdUserIdList(_), error: %v", err)
		return
	}

	rList = values

	return
}

// SelectByMessageDataIdUserIdListWithCB
// select user_id, user_message_box_id, dialog_id1, dialog_id2, dialog_message_id, sender_user_id, peer_type, peer_id, random_id, message_filter_type, message_data, message, mentioned, media_unread, pinned, has_reaction, reaction, reaction_date, reaction_unread, saved_peer_type, saved_peer_id, date2, ttl_period from messages where dialog_message_id = :dialog_message_id and user_id in (:idList) and deleted = 0
func (m *defaultMessagesModel) SelectByMessageDataIdUserIdListWithCB(ctx context.Context, dialogMessageId int64, idList []int64, cb func(sz, i int, v *Messages)) (rList []Messages, err error) {
	var (
		query  = fmt.Sprintf("select user_id, user_message_box_id, dialog_id1, dialog_id2, dialog_message_id, sender_user_id, peer_type, peer_id, random_id, message_filter_type, message_data, message, mentioned, media_unread, pinned, has_reaction, reaction, reaction_date, reaction_unread, saved_peer_type, saved_peer_id, date2, ttl_period from messages where dialog_message_id = ? and user_id in (%s) and deleted = 0", sqlx.InInt64List(idList))
		values []Messages
	)
	if len(idList) == 0 {
		rList = []Messages{}
		return
	}

	err = m.db.QueryRowsPartial(ctx, &values, query, dialogMessageId)

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

// SelectBackwardByOffsetIdLimit
// select user_id, user_message_box_id, dialog_id1, dialog_id2, dialog_message_id, sender_user_id, peer_type, peer_id, random_id, message_filter_type, message_data, message, mentioned, media_unread, pinned, has_reaction, reaction, reaction_date, reaction_unread, saved_peer_type, saved_peer_id, date2, ttl_period from messages where user_id = :user_id and (dialog_id1 = :dialog_id1 and dialog_id2 = :dialog_id2) and user_message_box_id < :user_message_box_id and deleted = 0 order by user_message_box_id desc limit :limit
func (m *defaultMessagesModel) SelectBackwardByOffsetIdLimit(ctx context.Context, userId int64, dialogId1 int64, dialogId2 int64, userMessageBoxId int32, limit int32) (rList []Messages, err error) {
	var (
		query  = "select user_id, user_message_box_id, dialog_id1, dialog_id2, dialog_message_id, sender_user_id, peer_type, peer_id, random_id, message_filter_type, message_data, message, mentioned, media_unread, pinned, has_reaction, reaction, reaction_date, reaction_unread, saved_peer_type, saved_peer_id, date2, ttl_period from messages where user_id = ? and (dialog_id1 = ? and dialog_id2 = ?) and user_message_box_id < ? and deleted = 0 order by user_message_box_id desc limit ?"
		values []Messages
	)
	err = m.db.QueryRowsPartial(ctx, &values, query, userId, dialogId1, dialogId2, userMessageBoxId, limit)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectBackwardByOffsetIdLimit(_), error: %v", err)
		return
	}

	rList = values

	return
}

// SelectBackwardByOffsetIdLimitWithCB
// select user_id, user_message_box_id, dialog_id1, dialog_id2, dialog_message_id, sender_user_id, peer_type, peer_id, random_id, message_filter_type, message_data, message, mentioned, media_unread, pinned, has_reaction, reaction, reaction_date, reaction_unread, saved_peer_type, saved_peer_id, date2, ttl_period from messages where user_id = :user_id and (dialog_id1 = :dialog_id1 and dialog_id2 = :dialog_id2) and user_message_box_id < :user_message_box_id and deleted = 0 order by user_message_box_id desc limit :limit
func (m *defaultMessagesModel) SelectBackwardByOffsetIdLimitWithCB(ctx context.Context, userId int64, dialogId1 int64, dialogId2 int64, userMessageBoxId int32, limit int32, cb func(sz, i int, v *Messages)) (rList []Messages, err error) {
	var (
		query  = "select user_id, user_message_box_id, dialog_id1, dialog_id2, dialog_message_id, sender_user_id, peer_type, peer_id, random_id, message_filter_type, message_data, message, mentioned, media_unread, pinned, has_reaction, reaction, reaction_date, reaction_unread, saved_peer_type, saved_peer_id, date2, ttl_period from messages where user_id = ? and (dialog_id1 = ? and dialog_id2 = ?) and user_message_box_id < ? and deleted = 0 order by user_message_box_id desc limit ?"
		values []Messages
	)
	err = m.db.QueryRowsPartial(ctx, &values, query, userId, dialogId1, dialogId2, userMessageBoxId, limit)

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
func (m *defaultMessagesModel) SelectForwardByOffsetIdLimit(ctx context.Context, userId int64, dialogId1 int64, dialogId2 int64, userMessageBoxId int32, limit int32) (rList []Messages, err error) {
	var (
		query  = "select user_id, user_message_box_id, dialog_id1, dialog_id2, dialog_message_id, sender_user_id, peer_type, peer_id, random_id, message_filter_type, message_data, message, mentioned, media_unread, pinned, has_reaction, reaction, reaction_date, reaction_unread, saved_peer_type, saved_peer_id, date2, ttl_period from messages where user_id = ? and (dialog_id1 = ? and dialog_id2 = ?) and user_message_box_id >= ? and deleted = 0 order by user_message_box_id asc limit ?"
		values []Messages
	)
	err = m.db.QueryRowsPartial(ctx, &values, query, userId, dialogId1, dialogId2, userMessageBoxId, limit)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectForwardByOffsetIdLimit(_), error: %v", err)
		return
	}

	rList = values

	return
}

// SelectForwardByOffsetIdLimitWithCB
// select user_id, user_message_box_id, dialog_id1, dialog_id2, dialog_message_id, sender_user_id, peer_type, peer_id, random_id, message_filter_type, message_data, message, mentioned, media_unread, pinned, has_reaction, reaction, reaction_date, reaction_unread, saved_peer_type, saved_peer_id, date2, ttl_period from messages where user_id = :user_id and (dialog_id1 = :dialog_id1 and dialog_id2 = :dialog_id2) and user_message_box_id >= :user_message_box_id and deleted = 0 order by user_message_box_id asc limit :limit
func (m *defaultMessagesModel) SelectForwardByOffsetIdLimitWithCB(ctx context.Context, userId int64, dialogId1 int64, dialogId2 int64, userMessageBoxId int32, limit int32, cb func(sz, i int, v *Messages)) (rList []Messages, err error) {
	var (
		query  = "select user_id, user_message_box_id, dialog_id1, dialog_id2, dialog_message_id, sender_user_id, peer_type, peer_id, random_id, message_filter_type, message_data, message, mentioned, media_unread, pinned, has_reaction, reaction, reaction_date, reaction_unread, saved_peer_type, saved_peer_id, date2, ttl_period from messages where user_id = ? and (dialog_id1 = ? and dialog_id2 = ?) and user_message_box_id >= ? and deleted = 0 order by user_message_box_id asc limit ?"
		values []Messages
	)
	err = m.db.QueryRowsPartial(ctx, &values, query, userId, dialogId1, dialogId2, userMessageBoxId, limit)

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
func (m *defaultMessagesModel) SelectBackwardByOffsetDateLimit(ctx context.Context, userId int64, dialogId1 int64, dialogId2 int64, date2 int64, limit int32) (rList []Messages, err error) {
	var (
		query  = "select user_id, user_message_box_id, dialog_id1, dialog_id2, dialog_message_id, sender_user_id, peer_type, peer_id, random_id, message_filter_type, message_data, message, mentioned, media_unread, pinned, has_reaction, reaction, reaction_date, reaction_unread, saved_peer_type, saved_peer_id, date2, ttl_period from messages where user_id = ? and (dialog_id1 = ? and dialog_id2 = ?) and date2 < ? and deleted = 0 order by user_message_box_id desc limit ?"
		values []Messages
	)
	err = m.db.QueryRowsPartial(ctx, &values, query, userId, dialogId1, dialogId2, date2, limit)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectBackwardByOffsetDateLimit(_), error: %v", err)
		return
	}

	rList = values

	return
}

// SelectBackwardByOffsetDateLimitWithCB
// select user_id, user_message_box_id, dialog_id1, dialog_id2, dialog_message_id, sender_user_id, peer_type, peer_id, random_id, message_filter_type, message_data, message, mentioned, media_unread, pinned, has_reaction, reaction, reaction_date, reaction_unread, saved_peer_type, saved_peer_id, date2, ttl_period from messages where user_id = :user_id and (dialog_id1 = :dialog_id1 and dialog_id2 = :dialog_id2) and date2 < :date2 and deleted = 0 order by user_message_box_id desc limit :limit
func (m *defaultMessagesModel) SelectBackwardByOffsetDateLimitWithCB(ctx context.Context, userId int64, dialogId1 int64, dialogId2 int64, date2 int64, limit int32, cb func(sz, i int, v *Messages)) (rList []Messages, err error) {
	var (
		query  = "select user_id, user_message_box_id, dialog_id1, dialog_id2, dialog_message_id, sender_user_id, peer_type, peer_id, random_id, message_filter_type, message_data, message, mentioned, media_unread, pinned, has_reaction, reaction, reaction_date, reaction_unread, saved_peer_type, saved_peer_id, date2, ttl_period from messages where user_id = ? and (dialog_id1 = ? and dialog_id2 = ?) and date2 < ? and deleted = 0 order by user_message_box_id desc limit ?"
		values []Messages
	)
	err = m.db.QueryRowsPartial(ctx, &values, query, userId, dialogId1, dialogId2, date2, limit)

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
func (m *defaultMessagesModel) SelectForwardByOffsetDateLimit(ctx context.Context, userId int64, dialogId1 int64, dialogId2 int64, date2 int64, limit int32) (rList []Messages, err error) {
	var (
		query  = "select user_id, user_message_box_id, dialog_id1, dialog_id2, dialog_message_id, sender_user_id, peer_type, peer_id, random_id, message_filter_type, message_data, message, mentioned, media_unread, pinned, has_reaction, reaction, reaction_date, reaction_unread, saved_peer_type, saved_peer_id, date2, ttl_period from messages where user_id = ? and (dialog_id1 = ? and dialog_id2 = ?) and date2 >= ? and deleted = 0 order by user_message_box_id asc limit ?"
		values []Messages
	)
	err = m.db.QueryRowsPartial(ctx, &values, query, userId, dialogId1, dialogId2, date2, limit)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectForwardByOffsetDateLimit(_), error: %v", err)
		return
	}

	rList = values

	return
}

// SelectForwardByOffsetDateLimitWithCB
// select user_id, user_message_box_id, dialog_id1, dialog_id2, dialog_message_id, sender_user_id, peer_type, peer_id, random_id, message_filter_type, message_data, message, mentioned, media_unread, pinned, has_reaction, reaction, reaction_date, reaction_unread, saved_peer_type, saved_peer_id, date2, ttl_period from messages where user_id = :user_id and (dialog_id1 = :dialog_id1 and dialog_id2 = :dialog_id2) and date2 >= :date2 and deleted = 0 order by user_message_box_id asc limit :limit
func (m *defaultMessagesModel) SelectForwardByOffsetDateLimitWithCB(ctx context.Context, userId int64, dialogId1 int64, dialogId2 int64, date2 int64, limit int32, cb func(sz, i int, v *Messages)) (rList []Messages, err error) {
	var (
		query  = "select user_id, user_message_box_id, dialog_id1, dialog_id2, dialog_message_id, sender_user_id, peer_type, peer_id, random_id, message_filter_type, message_data, message, mentioned, media_unread, pinned, has_reaction, reaction, reaction_date, reaction_unread, saved_peer_type, saved_peer_id, date2, ttl_period from messages where user_id = ? and (dialog_id1 = ? and dialog_id2 = ?) and date2 >= ? and deleted = 0 order by user_message_box_id asc limit ?"
		values []Messages
	)
	err = m.db.QueryRowsPartial(ctx, &values, query, userId, dialogId1, dialogId2, date2, limit)

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
func (m *defaultMessagesModel) SelectPeerUserMessageId(ctx context.Context, peerId int64, userId int64, userMessageBoxId int32) (rValue *Messages, err error) {
	var (
		query = "select user_message_box_id, message_box_type from messages where user_id = ? and deleted = 0 and dialog_message_id = (select dialog_message_id from messages where user_id = ? and user_message_box_id = ? and deleted = 0 limit 1)"
		do    = &Messages{}
	)
	err = m.db.QueryRowPartial(ctx, do, query, peerId, userId, userMessageBoxId)

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
func (m *defaultMessagesModel) SelectPeerUserMessage(ctx context.Context, peerId int64, userId int64, userMessageBoxId int32) (rValue *Messages, err error) {
	var (
		query = "select user_id, user_message_box_id, dialog_id1, dialog_id2, dialog_message_id, sender_user_id, peer_type, peer_id, random_id, message_filter_type, message_data, message, mentioned, media_unread, pinned, has_reaction, reaction, reaction_date, reaction_unread, saved_peer_type, saved_peer_id, date2, ttl_period from messages where user_id = ? and deleted = 0 and dialog_message_id = (select dialog_message_id from messages where user_id = ? and user_message_box_id = ? and deleted = 0 limit 1)"
		do    = &Messages{}
	)
	err = m.db.QueryRowPartial(ctx, do, query, peerId, userId, userMessageBoxId)

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
func (m *defaultMessagesModel) SelectDialogLastMessageId(ctx context.Context, userId int64, dialogId1 int64, dialogId2 int64) (rValue int32, err error) {
	var query = "select user_message_box_id from messages where user_id = ? and dialog_id1 = ? and dialog_id2 = ? and deleted = 0 order by user_message_box_id desc limit 1"
	err = m.db.QueryRowPartial(ctx, &rValue, query, userId, dialogId1, dialogId2)

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
func (m *defaultMessagesModel) SelectDialogLastMessageIdNotIdList(ctx context.Context, userId int64, dialogId1 int64, dialogId2 int64, idList []int32) (rValue int32, err error) {
	var (
		query = fmt.Sprintf("select user_message_box_id from messages where user_id = ? and dialog_id1 = ? and dialog_id2 = ? and user_message_box_id not in (%s) and deleted = 0 order by user_message_box_id desc limit 1", sqlx.InInt32List(idList))
	)

	if len(idList) == 0 {
		return
	}

	err = m.db.QueryRowPartial(ctx, &rValue, query, userId, dialogId1, dialogId2)

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
func (m *defaultMessagesModel) SelectDialogsByMessageIdList(ctx context.Context, userId int64, idList []int32) (rList []Messages, err error) {
	var (
		query  = fmt.Sprintf("select user_id, user_message_box_id, dialog_id1, dialog_id2, dialog_message_id, sender_user_id, peer_type, peer_id, random_id, message_filter_type, message_data, message, mentioned, media_unread, pinned, has_reaction, reaction, reaction_date, reaction_unread, saved_peer_type, saved_peer_id, date2, ttl_period from messages where user_id = ? and user_message_box_id in (%s) and deleted = 0", sqlx.InInt32List(idList))
		values []Messages
	)
	if len(idList) == 0 {
		rList = []Messages{}
		return
	}

	err = m.db.QueryRowsPartial(ctx, &values, query, userId)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectDialogsByMessageIdList(_), error: %v", err)
		return
	}

	rList = values

	return
}

// SelectDialogsByMessageIdListWithCB
// select user_id, user_message_box_id, dialog_id1, dialog_id2, dialog_message_id, sender_user_id, peer_type, peer_id, random_id, message_filter_type, message_data, message, mentioned, media_unread, pinned, has_reaction, reaction, reaction_date, reaction_unread, saved_peer_type, saved_peer_id, date2, ttl_period from messages where user_id = :user_id and user_message_box_id in (:idList) and deleted = 0
func (m *defaultMessagesModel) SelectDialogsByMessageIdListWithCB(ctx context.Context, userId int64, idList []int32, cb func(sz, i int, v *Messages)) (rList []Messages, err error) {
	var (
		query  = fmt.Sprintf("select user_id, user_message_box_id, dialog_id1, dialog_id2, dialog_message_id, sender_user_id, peer_type, peer_id, random_id, message_filter_type, message_data, message, mentioned, media_unread, pinned, has_reaction, reaction, reaction_date, reaction_unread, saved_peer_type, saved_peer_id, date2, ttl_period from messages where user_id = ? and user_message_box_id in (%s) and deleted = 0", sqlx.InInt32List(idList))
		values []Messages
	)
	if len(idList) == 0 {
		rList = []Messages{}
		return
	}

	err = m.db.QueryRowsPartial(ctx, &values, query, userId)

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
func (m *defaultMessagesModel) SelectDialogLastMessageList(ctx context.Context, userId int64, dialogId1 int64, dialogId2 int64, limit int32) (rList []Messages, err error) {
	var (
		query  = "select user_id, user_message_box_id, dialog_id1, dialog_id2, dialog_message_id, sender_user_id, peer_type, peer_id, random_id, message_filter_type, message_data, message, mentioned, media_unread, pinned, has_reaction, reaction, reaction_date, reaction_unread, saved_peer_type, saved_peer_id, date2, ttl_period from messages where user_id = ? and (dialog_id1 = ? and dialog_id2 = ?) and deleted = 0 order by user_message_box_id desc limit ?"
		values []Messages
	)
	err = m.db.QueryRowsPartial(ctx, &values, query, userId, dialogId1, dialogId2, limit)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectDialogLastMessageList(_), error: %v", err)
		return
	}

	rList = values

	return
}

// SelectDialogLastMessageListWithCB
// select user_id, user_message_box_id, dialog_id1, dialog_id2, dialog_message_id, sender_user_id, peer_type, peer_id, random_id, message_filter_type, message_data, message, mentioned, media_unread, pinned, has_reaction, reaction, reaction_date, reaction_unread, saved_peer_type, saved_peer_id, date2, ttl_period from messages where user_id = :user_id and (dialog_id1 = :dialog_id1 and dialog_id2 = :dialog_id2) and deleted = 0 order by user_message_box_id desc limit :limit
func (m *defaultMessagesModel) SelectDialogLastMessageListWithCB(ctx context.Context, userId int64, dialogId1 int64, dialogId2 int64, limit int32, cb func(sz, i int, v *Messages)) (rList []Messages, err error) {
	var (
		query  = "select user_id, user_message_box_id, dialog_id1, dialog_id2, dialog_message_id, sender_user_id, peer_type, peer_id, random_id, message_filter_type, message_data, message, mentioned, media_unread, pinned, has_reaction, reaction, reaction_date, reaction_unread, saved_peer_type, saved_peer_id, date2, ttl_period from messages where user_id = ? and (dialog_id1 = ? and dialog_id2 = ?) and deleted = 0 order by user_message_box_id desc limit ?"
		values []Messages
	)
	err = m.db.QueryRowsPartial(ctx, &values, query, userId, dialogId1, dialogId2, limit)

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
func (m *defaultMessagesModel) DeleteMessagesByMessageIdList(ctx context.Context, userId int64, idList []int32) (rowsAffected int64, err error) {
	var (
		query   = fmt.Sprintf("update messages set deleted = 1 where user_id = ? and user_message_box_id in (%s) and deleted = 0", sqlx.InInt32List(idList))
		rResult sql.Result
	)

	if len(idList) == 0 {
		return
	}

	rResult, err = m.db.Exec(ctx, query, userId)

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
func (m *defaultMessagesModel) DeleteMessagesByMessageIdListTx(tx *sqlx.Tx, userId int64, idList []int32) (rowsAffected int64, err error) {
	var (
		query   = fmt.Sprintf("update messages set deleted = 1 where user_id = ? and user_message_box_id in (%s) and deleted = 0", sqlx.InInt32List(idList))
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
func (m *defaultMessagesModel) SelectDialogMessageIdList(ctx context.Context, userId int64, dialogId1 int64, dialogId2 int64) (rList []Messages, err error) {
	var (
		query  = "select user_id, user_message_box_id, dialog_id1, dialog_id2, dialog_message_id, sender_user_id, peer_type, peer_id, random_id, message_filter_type, message_data, message, mentioned, media_unread, pinned, has_reaction, reaction, reaction_date, reaction_unread, saved_peer_type, saved_peer_id, date2, ttl_period from messages where user_id = ? and (dialog_id1 = ? and dialog_id2 = ?) and deleted = 0 order by user_message_box_id desc"
		values []Messages
	)
	err = m.db.QueryRowsPartial(ctx, &values, query, userId, dialogId1, dialogId2)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectDialogMessageIdList(_), error: %v", err)
		return
	}

	rList = values

	return
}

// SelectDialogMessageIdListWithCB
// select user_id, user_message_box_id, dialog_id1, dialog_id2, dialog_message_id, sender_user_id, peer_type, peer_id, random_id, message_filter_type, message_data, message, mentioned, media_unread, pinned, has_reaction, reaction, reaction_date, reaction_unread, saved_peer_type, saved_peer_id, date2, ttl_period from messages where user_id = :user_id and (dialog_id1 = :dialog_id1 and dialog_id2 = :dialog_id2) and deleted = 0 order by user_message_box_id desc
func (m *defaultMessagesModel) SelectDialogMessageIdListWithCB(ctx context.Context, userId int64, dialogId1 int64, dialogId2 int64, cb func(sz, i int, v *Messages)) (rList []Messages, err error) {
	var (
		query  = "select user_id, user_message_box_id, dialog_id1, dialog_id2, dialog_message_id, sender_user_id, peer_type, peer_id, random_id, message_filter_type, message_data, message, mentioned, media_unread, pinned, has_reaction, reaction, reaction_date, reaction_unread, saved_peer_type, saved_peer_id, date2, ttl_period from messages where user_id = ? and (dialog_id1 = ? and dialog_id2 = ?) and deleted = 0 order by user_message_box_id desc"
		values []Messages
	)
	err = m.db.QueryRowsPartial(ctx, &values, query, userId, dialogId1, dialogId2)

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
func (m *defaultMessagesModel) UpdateMediaUnread(ctx context.Context, userId int64, userMessageBoxId int32) (rowsAffected int64, err error) {
	var (
		query   = "update messages set media_unread = 0 where user_id = ? and user_message_box_id = ?"
		rResult sql.Result
	)

	rResult, err = m.db.Exec(ctx, query, userId, userMessageBoxId)

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
func (m *defaultMessagesModel) UpdateMediaUnreadTx(tx *sqlx.Tx, userId int64, userMessageBoxId int32) (rowsAffected int64, err error) {
	var (
		query   = "update messages set media_unread = 0 where user_id = ? and user_message_box_id = ?"
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
func (m *defaultMessagesModel) UpdateMentionedAndMediaUnread(ctx context.Context, userId int64, userMessageBoxId int32) (rowsAffected int64, err error) {
	var (
		query   = "update messages set mentioned = 0, media_unread = 0 where user_id = ? and user_message_box_id = ?"
		rResult sql.Result
	)

	rResult, err = m.db.Exec(ctx, query, userId, userMessageBoxId)

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
func (m *defaultMessagesModel) UpdateMentionedAndMediaUnreadTx(tx *sqlx.Tx, userId int64, userMessageBoxId int32) (rowsAffected int64, err error) {
	var (
		query   = "update messages set mentioned = 0, media_unread = 0 where user_id = ? and user_message_box_id = ?"
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
func (m *defaultMessagesModel) SelectByMediaType(ctx context.Context, userId int64, dialogId1 int64, dialogId2 int64, messageFilterType int32, userMessageBoxId int32, limit int32) (rList []Messages, err error) {
	var (
		query  = "select user_id, user_message_box_id, dialog_id1, dialog_id2, dialog_message_id, sender_user_id, peer_type, peer_id, random_id, message_filter_type, message_data, message, mentioned, media_unread, pinned, has_reaction, reaction, reaction_date, reaction_unread, saved_peer_type, saved_peer_id, date2, ttl_period from messages where user_id = ? and (dialog_id1 = ? and dialog_id2 = ?) and message_filter_type = ? and user_message_box_id < ? and deleted = 0 order by user_message_box_id desc limit ?"
		values []Messages
	)
	err = m.db.QueryRowsPartial(ctx, &values, query, userId, dialogId1, dialogId2, messageFilterType, userMessageBoxId, limit)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectByMediaType(_), error: %v", err)
		return
	}

	rList = values

	return
}

// SelectByMediaTypeWithCB
// select user_id, user_message_box_id, dialog_id1, dialog_id2, dialog_message_id, sender_user_id, peer_type, peer_id, random_id, message_filter_type, message_data, message, mentioned, media_unread, pinned, has_reaction, reaction, reaction_date, reaction_unread, saved_peer_type, saved_peer_id, date2, ttl_period from messages where user_id = :user_id and (dialog_id1 = :dialog_id1 and dialog_id2 = :dialog_id2) and message_filter_type = :message_filter_type and user_message_box_id < :user_message_box_id and deleted = 0 order by user_message_box_id desc limit :limit
func (m *defaultMessagesModel) SelectByMediaTypeWithCB(ctx context.Context, userId int64, dialogId1 int64, dialogId2 int64, messageFilterType int32, userMessageBoxId int32, limit int32, cb func(sz, i int, v *Messages)) (rList []Messages, err error) {
	var (
		query  = "select user_id, user_message_box_id, dialog_id1, dialog_id2, dialog_message_id, sender_user_id, peer_type, peer_id, random_id, message_filter_type, message_data, message, mentioned, media_unread, pinned, has_reaction, reaction, reaction_date, reaction_unread, saved_peer_type, saved_peer_id, date2, ttl_period from messages where user_id = ? and (dialog_id1 = ? and dialog_id2 = ?) and message_filter_type = ? and user_message_box_id < ? and deleted = 0 order by user_message_box_id desc limit ?"
		values []Messages
	)
	err = m.db.QueryRowsPartial(ctx, &values, query, userId, dialogId1, dialogId2, messageFilterType, userMessageBoxId, limit)

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
func (m *defaultMessagesModel) SelectPhoneCallList(ctx context.Context, userId int64, messageFilterType int32, userMessageBoxId int32, limit int32) (rList []Messages, err error) {
	var (
		query  = "select user_id, user_message_box_id, dialog_id1, dialog_id2, dialog_message_id, sender_user_id, peer_type, peer_id, random_id, message_filter_type, message_data, message, mentioned, media_unread, pinned, has_reaction, reaction, reaction_date, reaction_unread, saved_peer_type, saved_peer_id, date2, ttl_period from messages where user_id = ? and message_filter_type = ? and user_message_box_id < ? and deleted = 0 order by user_message_box_id desc limit ?"
		values []Messages
	)
	err = m.db.QueryRowsPartial(ctx, &values, query, userId, messageFilterType, userMessageBoxId, limit)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectPhoneCallList(_), error: %v", err)
		return
	}

	rList = values

	return
}

// SelectPhoneCallListWithCB
// select user_id, user_message_box_id, dialog_id1, dialog_id2, dialog_message_id, sender_user_id, peer_type, peer_id, random_id, message_filter_type, message_data, message, mentioned, media_unread, pinned, has_reaction, reaction, reaction_date, reaction_unread, saved_peer_type, saved_peer_id, date2, ttl_period from messages where user_id = :user_id and message_filter_type = :message_filter_type and user_message_box_id < :user_message_box_id and deleted = 0 order by user_message_box_id desc limit :limit
func (m *defaultMessagesModel) SelectPhoneCallListWithCB(ctx context.Context, userId int64, messageFilterType int32, userMessageBoxId int32, limit int32, cb func(sz, i int, v *Messages)) (rList []Messages, err error) {
	var (
		query  = "select user_id, user_message_box_id, dialog_id1, dialog_id2, dialog_message_id, sender_user_id, peer_type, peer_id, random_id, message_filter_type, message_data, message, mentioned, media_unread, pinned, has_reaction, reaction, reaction_date, reaction_unread, saved_peer_type, saved_peer_id, date2, ttl_period from messages where user_id = ? and message_filter_type = ? and user_message_box_id < ? and deleted = 0 order by user_message_box_id desc limit ?"
		values []Messages
	)
	err = m.db.QueryRowsPartial(ctx, &values, query, userId, messageFilterType, userMessageBoxId, limit)

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
func (m *defaultMessagesModel) Search(ctx context.Context, userId int64, dialogId1 int64, dialogId2 int64, userMessageBoxId int32, q2 string, limit int32) (rList []Messages, err error) {
	var (
		query  = "select user_id, user_message_box_id, dialog_id1, dialog_id2, dialog_message_id, sender_user_id, peer_type, peer_id, random_id, message_filter_type, message_data, message, mentioned, media_unread, pinned, has_reaction, reaction, reaction_date, reaction_unread, saved_peer_type, saved_peer_id, date2, ttl_period from messages where user_id = ? and (dialog_id1 = ? and dialog_id2 = ?) and user_message_box_id < ? and deleted = 0 and message != '' and message like ? order by user_message_box_id desc limit ?"
		values []Messages
	)
	err = m.db.QueryRowsPartial(ctx, &values, query, userId, dialogId1, dialogId2, userMessageBoxId, q2, limit)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in Search(_), error: %v", err)
		return
	}

	rList = values

	return
}

// SearchWithCB
// select user_id, user_message_box_id, dialog_id1, dialog_id2, dialog_message_id, sender_user_id, peer_type, peer_id, random_id, message_filter_type, message_data, message, mentioned, media_unread, pinned, has_reaction, reaction, reaction_date, reaction_unread, saved_peer_type, saved_peer_id, date2, ttl_period from messages where user_id = :user_id and (dialog_id1 = :dialog_id1 and dialog_id2 = :dialog_id2) and user_message_box_id < :user_message_box_id and deleted = 0 and message != ” and message like :q2 order by user_message_box_id desc limit :limit
func (m *defaultMessagesModel) SearchWithCB(ctx context.Context, userId int64, dialogId1 int64, dialogId2 int64, userMessageBoxId int32, q2 string, limit int32, cb func(sz, i int, v *Messages)) (rList []Messages, err error) {
	var (
		query  = "select user_id, user_message_box_id, dialog_id1, dialog_id2, dialog_message_id, sender_user_id, peer_type, peer_id, random_id, message_filter_type, message_data, message, mentioned, media_unread, pinned, has_reaction, reaction, reaction_date, reaction_unread, saved_peer_type, saved_peer_id, date2, ttl_period from messages where user_id = ? and (dialog_id1 = ? and dialog_id2 = ?) and user_message_box_id < ? and deleted = 0 and message != '' and message like ? order by user_message_box_id desc limit ?"
		values []Messages
	)
	err = m.db.QueryRowsPartial(ctx, &values, query, userId, dialogId1, dialogId2, userMessageBoxId, q2, limit)

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
func (m *defaultMessagesModel) SearchGlobal(ctx context.Context, userId int64, userMessageBoxId int32, q2 string, limit int32) (rList []Messages, err error) {
	var (
		query  = "select user_id, user_message_box_id, dialog_id1, dialog_id2, dialog_message_id, sender_user_id, peer_type, peer_id, random_id, message_filter_type, message_data, message, mentioned, media_unread, pinned, has_reaction, reaction, reaction_date, reaction_unread, saved_peer_type, saved_peer_id, date2, ttl_period from messages where user_id = ? and user_message_box_id < ? and deleted = 0 and message != '' and message like ? order by user_message_box_id desc limit ?"
		values []Messages
	)
	err = m.db.QueryRowsPartial(ctx, &values, query, userId, userMessageBoxId, q2, limit)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SearchGlobal(_), error: %v", err)
		return
	}

	rList = values

	return
}

// SearchGlobalWithCB
// select user_id, user_message_box_id, dialog_id1, dialog_id2, dialog_message_id, sender_user_id, peer_type, peer_id, random_id, message_filter_type, message_data, message, mentioned, media_unread, pinned, has_reaction, reaction, reaction_date, reaction_unread, saved_peer_type, saved_peer_id, date2, ttl_period from messages where user_id = :user_id and user_message_box_id < :user_message_box_id and deleted = 0 and message != ” and message like :q2 order by user_message_box_id desc limit :limit
func (m *defaultMessagesModel) SearchGlobalWithCB(ctx context.Context, userId int64, userMessageBoxId int32, q2 string, limit int32, cb func(sz, i int, v *Messages)) (rList []Messages, err error) {
	var (
		query  = "select user_id, user_message_box_id, dialog_id1, dialog_id2, dialog_message_id, sender_user_id, peer_type, peer_id, random_id, message_filter_type, message_data, message, mentioned, media_unread, pinned, has_reaction, reaction, reaction_date, reaction_unread, saved_peer_type, saved_peer_id, date2, ttl_period from messages where user_id = ? and user_message_box_id < ? and deleted = 0 and message != '' and message like ? order by user_message_box_id desc limit ?"
		values []Messages
	)
	err = m.db.QueryRowsPartial(ctx, &values, query, userId, userMessageBoxId, q2, limit)

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
func (m *defaultMessagesModel) SelectBackwardUnreadMentionsByOffsetIdLimit(ctx context.Context, userId int64, dialogId1 int64, dialogId2 int64, userMessageBoxId int32, limit int32) (rList []Messages, err error) {
	var (
		query  = "select user_id, user_message_box_id, dialog_id1, dialog_id2, dialog_message_id, sender_user_id, peer_type, peer_id, random_id, message_filter_type, message_data, message, mentioned, media_unread, pinned, has_reaction, reaction, reaction_date, reaction_unread, saved_peer_type, saved_peer_id, date2, ttl_period from messages where user_id = ? and (dialog_id1 = ? and dialog_id2 = ?) and user_message_box_id < ? and mentioned = 1 and media_unread = 1 and deleted = 0 order by user_message_box_id desc limit ?"
		values []Messages
	)
	err = m.db.QueryRowsPartial(ctx, &values, query, userId, dialogId1, dialogId2, userMessageBoxId, limit)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectBackwardUnreadMentionsByOffsetIdLimit(_), error: %v", err)
		return
	}

	rList = values

	return
}

// SelectBackwardUnreadMentionsByOffsetIdLimitWithCB
// select user_id, user_message_box_id, dialog_id1, dialog_id2, dialog_message_id, sender_user_id, peer_type, peer_id, random_id, message_filter_type, message_data, message, mentioned, media_unread, pinned, has_reaction, reaction, reaction_date, reaction_unread, saved_peer_type, saved_peer_id, date2, ttl_period from messages where user_id = :user_id and (dialog_id1 = :dialog_id1 and dialog_id2 = :dialog_id2) and user_message_box_id < :user_message_box_id and mentioned = 1 and media_unread = 1 and deleted = 0 order by user_message_box_id desc limit :limit
func (m *defaultMessagesModel) SelectBackwardUnreadMentionsByOffsetIdLimitWithCB(ctx context.Context, userId int64, dialogId1 int64, dialogId2 int64, userMessageBoxId int32, limit int32, cb func(sz, i int, v *Messages)) (rList []Messages, err error) {
	var (
		query  = "select user_id, user_message_box_id, dialog_id1, dialog_id2, dialog_message_id, sender_user_id, peer_type, peer_id, random_id, message_filter_type, message_data, message, mentioned, media_unread, pinned, has_reaction, reaction, reaction_date, reaction_unread, saved_peer_type, saved_peer_id, date2, ttl_period from messages where user_id = ? and (dialog_id1 = ? and dialog_id2 = ?) and user_message_box_id < ? and mentioned = 1 and media_unread = 1 and deleted = 0 order by user_message_box_id desc limit ?"
		values []Messages
	)
	err = m.db.QueryRowsPartial(ctx, &values, query, userId, dialogId1, dialogId2, userMessageBoxId, limit)

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
func (m *defaultMessagesModel) SelectForwardUnreadMentionsByOffsetIdLimit(ctx context.Context, userId int64, dialogId1 int64, dialogId2 int64, userMessageBoxId int32, limit int32) (rList []Messages, err error) {
	var (
		query  = "select user_id, user_message_box_id, dialog_id1, dialog_id2, dialog_message_id, sender_user_id, peer_type, peer_id, random_id, message_filter_type, message_data, message, mentioned, media_unread, pinned, has_reaction, reaction, reaction_date, reaction_unread, saved_peer_type, saved_peer_id, date2, ttl_period from messages where user_id = ? and (dialog_id1 = ? and dialog_id2 = ?) and user_message_box_id >= ? and mentioned = 1 and media_unread = 1 and deleted = 0 order by user_message_box_id asc limit ?"
		values []Messages
	)
	err = m.db.QueryRowsPartial(ctx, &values, query, userId, dialogId1, dialogId2, userMessageBoxId, limit)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectForwardUnreadMentionsByOffsetIdLimit(_), error: %v", err)
		return
	}

	rList = values

	return
}

// SelectForwardUnreadMentionsByOffsetIdLimitWithCB
// select user_id, user_message_box_id, dialog_id1, dialog_id2, dialog_message_id, sender_user_id, peer_type, peer_id, random_id, message_filter_type, message_data, message, mentioned, media_unread, pinned, has_reaction, reaction, reaction_date, reaction_unread, saved_peer_type, saved_peer_id, date2, ttl_period from messages where user_id = :user_id and (dialog_id1 = :dialog_id1 and dialog_id2 = :dialog_id2) and user_message_box_id >= :user_message_box_id and mentioned = 1 and media_unread = 1 and deleted = 0 order by user_message_box_id asc limit :limit
func (m *defaultMessagesModel) SelectForwardUnreadMentionsByOffsetIdLimitWithCB(ctx context.Context, userId int64, dialogId1 int64, dialogId2 int64, userMessageBoxId int32, limit int32, cb func(sz, i int, v *Messages)) (rList []Messages, err error) {
	var (
		query  = "select user_id, user_message_box_id, dialog_id1, dialog_id2, dialog_message_id, sender_user_id, peer_type, peer_id, random_id, message_filter_type, message_data, message, mentioned, media_unread, pinned, has_reaction, reaction, reaction_date, reaction_unread, saved_peer_type, saved_peer_id, date2, ttl_period from messages where user_id = ? and (dialog_id1 = ? and dialog_id2 = ?) and user_message_box_id >= ? and mentioned = 1 and media_unread = 1 and deleted = 0 order by user_message_box_id asc limit ?"
		values []Messages
	)
	err = m.db.QueryRowsPartial(ctx, &values, query, userId, dialogId1, dialogId2, userMessageBoxId, limit)

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
func (m *defaultMessagesModel) SelectPinnedList(ctx context.Context, userId int64, dialogId1 int64, dialogId2 int64) (rList []Messages, err error) {
	var (
		query  = "select user_id, user_message_box_id, dialog_id1, dialog_id2, dialog_message_id, sender_user_id, peer_type, peer_id, random_id, message_filter_type, message_data, message, mentioned, media_unread, pinned, has_reaction, reaction, reaction_date, reaction_unread, saved_peer_type, saved_peer_id, date2, ttl_period from messages where user_id = ? and (dialog_id1 = ? and dialog_id2 = ?) and pinned = 1 and deleted = 0 order by user_message_box_id desc"
		values []Messages
	)
	err = m.db.QueryRowsPartial(ctx, &values, query, userId, dialogId1, dialogId2)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectPinnedList(_), error: %v", err)
		return
	}

	rList = values

	return
}

// SelectPinnedListWithCB
// select user_id, user_message_box_id, dialog_id1, dialog_id2, dialog_message_id, sender_user_id, peer_type, peer_id, random_id, message_filter_type, message_data, message, mentioned, media_unread, pinned, has_reaction, reaction, reaction_date, reaction_unread, saved_peer_type, saved_peer_id, date2, ttl_period from messages where user_id = :user_id and (dialog_id1 = :dialog_id1 and dialog_id2 = :dialog_id2) and pinned = 1 and deleted = 0 order by user_message_box_id desc
func (m *defaultMessagesModel) SelectPinnedListWithCB(ctx context.Context, userId int64, dialogId1 int64, dialogId2 int64, cb func(sz, i int, v *Messages)) (rList []Messages, err error) {
	var (
		query  = "select user_id, user_message_box_id, dialog_id1, dialog_id2, dialog_message_id, sender_user_id, peer_type, peer_id, random_id, message_filter_type, message_data, message, mentioned, media_unread, pinned, has_reaction, reaction, reaction_date, reaction_unread, saved_peer_type, saved_peer_id, date2, ttl_period from messages where user_id = ? and (dialog_id1 = ? and dialog_id2 = ?) and pinned = 1 and deleted = 0 order by user_message_box_id desc"
		values []Messages
	)
	err = m.db.QueryRowsPartial(ctx, &values, query, userId, dialogId1, dialogId2)

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
func (m *defaultMessagesModel) SelectLastTwoPinnedList(ctx context.Context, userId int64, dialogId1 int64, dialogId2 int64) (rList []int32, err error) {
	var query = "select user_message_box_id from messages where user_id = ? and (dialog_id1 = ? and dialog_id2 = ?) and pinned = 1 and deleted = 0 order by user_message_box_id desc limit 2"
	err = m.db.QueryRowsPartial(ctx, &rList, query, userId, dialogId1, dialogId2)

	if err != nil {
		logx.WithContext(ctx).Errorf("select in SelectLastTwoPinnedList(_), error: %v", err)
	}

	return
}

// SelectLastTwoPinnedListWithCB
// select user_message_box_id from messages where user_id = :user_id and (dialog_id1 = :dialog_id1 and dialog_id2 = :dialog_id2) and pinned = 1 and deleted = 0 order by user_message_box_id desc limit 2
func (m *defaultMessagesModel) SelectLastTwoPinnedListWithCB(ctx context.Context, userId int64, dialogId1 int64, dialogId2 int64, cb func(sz, i int, v int32)) (rList []int32, err error) {
	var query = "select user_message_box_id from messages where user_id = ? and (dialog_id1 = ? and dialog_id2 = ?) and pinned = 1 and deleted = 0 order by user_message_box_id desc limit 2"
	err = m.db.QueryRowsPartial(ctx, &rList, query, userId, dialogId1, dialogId2)

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
func (m *defaultMessagesModel) UpdatePinned(ctx context.Context, pinned bool, userId int64, userMessageBoxId int32) (rowsAffected int64, err error) {
	var (
		query   = "update messages set pinned = ? where user_id = ? and user_message_box_id = ?"
		rResult sql.Result
	)

	rResult, err = m.db.Exec(ctx, query, pinned, userId, userMessageBoxId)

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
func (m *defaultMessagesModel) UpdatePinnedTx(tx *sqlx.Tx, pinned bool, userId int64, userMessageBoxId int32) (rowsAffected int64, err error) {
	var (
		query   = "update messages set pinned = ? where user_id = ? and user_message_box_id = ?"
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
func (m *defaultMessagesModel) SelectPinnedMessageIdList(ctx context.Context, userId int64, dialogId1 int64, dialogId2 int64) (rList []int32, err error) {
	var query = "select user_message_box_id from messages where user_id = ? and (dialog_id1 = ? and dialog_id2 = ?) and pinned = 1 and deleted = 0 order by user_message_box_id desc"
	err = m.db.QueryRowsPartial(ctx, &rList, query, userId, dialogId1, dialogId2)

	if err != nil {
		logx.WithContext(ctx).Errorf("select in SelectPinnedMessageIdList(_), error: %v", err)
	}

	return
}

// SelectPinnedMessageIdListWithCB
// select user_message_box_id from messages where user_id = :user_id and (dialog_id1 = :dialog_id1 and dialog_id2 = :dialog_id2) and pinned = 1 and deleted = 0 order by user_message_box_id desc
func (m *defaultMessagesModel) SelectPinnedMessageIdListWithCB(ctx context.Context, userId int64, dialogId1 int64, dialogId2 int64, cb func(sz, i int, v int32)) (rList []int32, err error) {
	var query = "select user_message_box_id from messages where user_id = ? and (dialog_id1 = ? and dialog_id2 = ?) and pinned = 1 and deleted = 0 order by user_message_box_id desc"
	err = m.db.QueryRowsPartial(ctx, &rList, query, userId, dialogId1, dialogId2)

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
func (m *defaultMessagesModel) UpdateUnPinnedByIdList(ctx context.Context, userId int64, idList []int32) (rowsAffected int64, err error) {
	var (
		query   = fmt.Sprintf("update messages set pinned = 0 where user_id = ? and user_message_box_id in (%s)", sqlx.InInt32List(idList))
		rResult sql.Result
	)

	if len(idList) == 0 {
		return
	}

	rResult, err = m.db.Exec(ctx, query, userId)

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
func (m *defaultMessagesModel) UpdateUnPinnedByIdListTx(tx *sqlx.Tx, userId int64, idList []int32) (rowsAffected int64, err error) {
	var (
		query   = fmt.Sprintf("update messages set pinned = 0 where user_id = ? and user_message_box_id in (%s)", sqlx.InInt32List(idList))
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
func (m *defaultMessagesModel) UpdateEditMessage(ctx context.Context, messageData string, message string, userId int64, userMessageBoxId int32) (rowsAffected int64, err error) {
	var (
		query   = "update messages set message_data = ?, message = ? where user_id = ? and user_message_box_id = ?"
		rResult sql.Result
	)

	rResult, err = m.db.Exec(ctx, query, messageData, message, userId, userMessageBoxId)

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
func (m *defaultMessagesModel) UpdateEditMessageTx(tx *sqlx.Tx, messageData string, message string, userId int64, userMessageBoxId int32) (rowsAffected int64, err error) {
	var (
		query   = "update messages set message_data = ?, message = ? where user_id = ? and user_message_box_id = ?"
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
func (m *defaultMessagesModel) UpdateCustomMap(ctx context.Context, cMap map[string]interface{}, userId int64, userMessageBoxId int32) (rowsAffected int64, err error) {
	names := make([]string, 0, len(cMap))
	aValues := make([]interface{}, 0, len(cMap))
	for k, v := range cMap {
		names = append(names, k+" = ?")
		aValues = append(aValues, v)
	}

	var (
		query   = fmt.Sprintf("update messages set %s where user_id = ? and user_message_box_id = ?", strings.Join(names, ", "))
		rResult sql.Result
	)

	aValues = append(aValues, userId)
	aValues = append(aValues, userMessageBoxId)

	rResult, err = m.db.Exec(ctx, query, aValues...)

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
func (m *defaultMessagesModel) UpdateCustomMapTx(tx *sqlx.Tx, cMap map[string]interface{}, userId int64, userMessageBoxId int32) (rowsAffected int64, err error) {
	names := make([]string, 0, len(cMap))
	aValues := make([]interface{}, 0, len(cMap))
	for k, v := range cMap {
		names = append(names, k+" = ?")
		aValues = append(aValues, v)
	}

	var (
		query   = fmt.Sprintf("update messages set %s where user_id = ? and user_message_box_id = ?", strings.Join(names, ", "))
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
func (m *defaultMessagesModel) SelectBackwardBySendUserIdOffsetIdLimit(ctx context.Context, userId int64, dialogId1 int64, dialogId2 int64, senderUserId int64, userMessageBoxId int32, limit int32) (rList []Messages, err error) {
	var (
		query  = "select user_id, user_message_box_id, dialog_id1, dialog_id2, dialog_message_id, sender_user_id, peer_type, peer_id, random_id, message_filter_type, message_data, message, mentioned, media_unread, pinned, has_reaction, reaction, reaction_date, reaction_unread, saved_peer_type, saved_peer_id, date2, ttl_period from messages where user_id = ? and (dialog_id1 = ? and dialog_id2 = ?) and sender_user_id = ? and user_message_box_id < ? and deleted = 0 order by user_message_box_id desc limit ?"
		values []Messages
	)
	err = m.db.QueryRowsPartial(ctx, &values, query, userId, dialogId1, dialogId2, senderUserId, userMessageBoxId, limit)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectBackwardBySendUserIdOffsetIdLimit(_), error: %v", err)
		return
	}

	rList = values

	return
}

// SelectBackwardBySendUserIdOffsetIdLimitWithCB
// select user_id, user_message_box_id, dialog_id1, dialog_id2, dialog_message_id, sender_user_id, peer_type, peer_id, random_id, message_filter_type, message_data, message, mentioned, media_unread, pinned, has_reaction, reaction, reaction_date, reaction_unread, saved_peer_type, saved_peer_id, date2, ttl_period from messages where user_id = :user_id and (dialog_id1 = :dialog_id1 and dialog_id2 = :dialog_id2) and sender_user_id = :sender_user_id and user_message_box_id < :user_message_box_id and deleted = 0 order by user_message_box_id desc limit :limit
func (m *defaultMessagesModel) SelectBackwardBySendUserIdOffsetIdLimitWithCB(ctx context.Context, userId int64, dialogId1 int64, dialogId2 int64, senderUserId int64, userMessageBoxId int32, limit int32, cb func(sz, i int, v *Messages)) (rList []Messages, err error) {
	var (
		query  = "select user_id, user_message_box_id, dialog_id1, dialog_id2, dialog_message_id, sender_user_id, peer_type, peer_id, random_id, message_filter_type, message_data, message, mentioned, media_unread, pinned, has_reaction, reaction, reaction_date, reaction_unread, saved_peer_type, saved_peer_id, date2, ttl_period from messages where user_id = ? and (dialog_id1 = ? and dialog_id2 = ?) and sender_user_id = ? and user_message_box_id < ? and deleted = 0 order by user_message_box_id desc limit ?"
		values []Messages
	)
	err = m.db.QueryRowsPartial(ctx, &values, query, userId, dialogId1, dialogId2, senderUserId, userMessageBoxId, limit)

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
func (m *defaultMessagesModel) SelectBackwardSavedByOffsetIdLimit(ctx context.Context, userId int64, savedPeerType int32, savedPeerId int64, userMessageBoxId int32, limit int32) (rList []Messages, err error) {
	var (
		query  = "select user_id, user_message_box_id, dialog_id1, dialog_id2, dialog_message_id, sender_user_id, peer_type, peer_id, random_id, message_filter_type, message_data, message, mentioned, media_unread, pinned, has_reaction, reaction, reaction_date, reaction_unread, saved_peer_type, saved_peer_id, date2, ttl_period from messages where user_id = ? and (saved_peer_type = ? and saved_peer_id = ?) and user_message_box_id < ? and deleted = 0 order by user_message_box_id desc limit ?"
		values []Messages
	)
	err = m.db.QueryRowsPartial(ctx, &values, query, userId, savedPeerType, savedPeerId, userMessageBoxId, limit)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectBackwardSavedByOffsetIdLimit(_), error: %v", err)
		return
	}

	rList = values

	return
}

// SelectBackwardSavedByOffsetIdLimitWithCB
// select user_id, user_message_box_id, dialog_id1, dialog_id2, dialog_message_id, sender_user_id, peer_type, peer_id, random_id, message_filter_type, message_data, message, mentioned, media_unread, pinned, has_reaction, reaction, reaction_date, reaction_unread, saved_peer_type, saved_peer_id, date2, ttl_period from messages where user_id = :user_id and (saved_peer_type = :saved_peer_type and saved_peer_id = :saved_peer_id) and user_message_box_id < :user_message_box_id and deleted = 0 order by user_message_box_id desc limit :limit
func (m *defaultMessagesModel) SelectBackwardSavedByOffsetIdLimitWithCB(ctx context.Context, userId int64, savedPeerType int32, savedPeerId int64, userMessageBoxId int32, limit int32, cb func(sz, i int, v *Messages)) (rList []Messages, err error) {
	var (
		query  = "select user_id, user_message_box_id, dialog_id1, dialog_id2, dialog_message_id, sender_user_id, peer_type, peer_id, random_id, message_filter_type, message_data, message, mentioned, media_unread, pinned, has_reaction, reaction, reaction_date, reaction_unread, saved_peer_type, saved_peer_id, date2, ttl_period from messages where user_id = ? and (saved_peer_type = ? and saved_peer_id = ?) and user_message_box_id < ? and deleted = 0 order by user_message_box_id desc limit ?"
		values []Messages
	)
	err = m.db.QueryRowsPartial(ctx, &values, query, userId, savedPeerType, savedPeerId, userMessageBoxId, limit)

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
func (m *defaultMessagesModel) SelectForwardSavedByOffsetIdLimit(ctx context.Context, userId int64, savedPeerType int32, savedPeerId int64, userMessageBoxId int32, limit int32) (rList []Messages, err error) {
	var (
		query  = "select user_id, user_message_box_id, dialog_id1, dialog_id2, dialog_message_id, sender_user_id, peer_type, peer_id, random_id, message_filter_type, message_data, message, mentioned, media_unread, pinned, has_reaction, reaction, reaction_date, reaction_unread, saved_peer_type, saved_peer_id, date2, ttl_period from messages where user_id = ? and (saved_peer_type = ? and saved_peer_id = ?) and user_message_box_id >= ? and deleted = 0 order by user_message_box_id asc limit ?"
		values []Messages
	)
	err = m.db.QueryRowsPartial(ctx, &values, query, userId, savedPeerType, savedPeerId, userMessageBoxId, limit)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectForwardSavedByOffsetIdLimit(_), error: %v", err)
		return
	}

	rList = values

	return
}

// SelectForwardSavedByOffsetIdLimitWithCB
// select user_id, user_message_box_id, dialog_id1, dialog_id2, dialog_message_id, sender_user_id, peer_type, peer_id, random_id, message_filter_type, message_data, message, mentioned, media_unread, pinned, has_reaction, reaction, reaction_date, reaction_unread, saved_peer_type, saved_peer_id, date2, ttl_period from messages where user_id = :user_id and (saved_peer_type = :saved_peer_type and saved_peer_id = :saved_peer_id) and user_message_box_id >= :user_message_box_id and deleted = 0 order by user_message_box_id asc limit :limit
func (m *defaultMessagesModel) SelectForwardSavedByOffsetIdLimitWithCB(ctx context.Context, userId int64, savedPeerType int32, savedPeerId int64, userMessageBoxId int32, limit int32, cb func(sz, i int, v *Messages)) (rList []Messages, err error) {
	var (
		query  = "select user_id, user_message_box_id, dialog_id1, dialog_id2, dialog_message_id, sender_user_id, peer_type, peer_id, random_id, message_filter_type, message_data, message, mentioned, media_unread, pinned, has_reaction, reaction, reaction_date, reaction_unread, saved_peer_type, saved_peer_id, date2, ttl_period from messages where user_id = ? and (saved_peer_type = ? and saved_peer_id = ?) and user_message_box_id >= ? and deleted = 0 order by user_message_box_id asc limit ?"
		values []Messages
	)
	err = m.db.QueryRowsPartial(ctx, &values, query, userId, savedPeerType, savedPeerId, userMessageBoxId, limit)

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
func (m *defaultMessagesModel) SelectBackwardSavedByOffsetDateLimit(ctx context.Context, userId int64, savedPeerType int32, savedPeerId int64, date2 int64, limit int32) (rList []Messages, err error) {
	var (
		query  = "select user_id, user_message_box_id, dialog_id1, dialog_id2, dialog_message_id, sender_user_id, peer_type, peer_id, random_id, message_filter_type, message_data, message, mentioned, media_unread, pinned, has_reaction, reaction, reaction_date, reaction_unread, saved_peer_type, saved_peer_id, date2, ttl_period from messages where user_id = ? and (saved_peer_type = ? and saved_peer_id = ?) and date2 < ? and deleted = 0 order by user_message_box_id desc limit ?"
		values []Messages
	)
	err = m.db.QueryRowsPartial(ctx, &values, query, userId, savedPeerType, savedPeerId, date2, limit)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectBackwardSavedByOffsetDateLimit(_), error: %v", err)
		return
	}

	rList = values

	return
}

// SelectBackwardSavedByOffsetDateLimitWithCB
// select user_id, user_message_box_id, dialog_id1, dialog_id2, dialog_message_id, sender_user_id, peer_type, peer_id, random_id, message_filter_type, message_data, message, mentioned, media_unread, pinned, has_reaction, reaction, reaction_date, reaction_unread, saved_peer_type, saved_peer_id, date2, ttl_period from messages where user_id = :user_id and (saved_peer_type = :saved_peer_type and saved_peer_id = :saved_peer_id) and date2 < :date2 and deleted = 0 order by user_message_box_id desc limit :limit
func (m *defaultMessagesModel) SelectBackwardSavedByOffsetDateLimitWithCB(ctx context.Context, userId int64, savedPeerType int32, savedPeerId int64, date2 int64, limit int32, cb func(sz, i int, v *Messages)) (rList []Messages, err error) {
	var (
		query  = "select user_id, user_message_box_id, dialog_id1, dialog_id2, dialog_message_id, sender_user_id, peer_type, peer_id, random_id, message_filter_type, message_data, message, mentioned, media_unread, pinned, has_reaction, reaction, reaction_date, reaction_unread, saved_peer_type, saved_peer_id, date2, ttl_period from messages where user_id = ? and (saved_peer_type = ? and saved_peer_id = ?) and date2 < ? and deleted = 0 order by user_message_box_id desc limit ?"
		values []Messages
	)
	err = m.db.QueryRowsPartial(ctx, &values, query, userId, savedPeerType, savedPeerId, date2, limit)

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
func (m *defaultMessagesModel) SelectForwardSavedByOffsetDateLimit(ctx context.Context, userId int64, savedPeerType int32, savedPeerId int64, date2 int64, limit int32) (rList []Messages, err error) {
	var (
		query  = "select user_id, user_message_box_id, dialog_id1, dialog_id2, dialog_message_id, sender_user_id, peer_type, peer_id, random_id, message_filter_type, message_data, message, mentioned, media_unread, pinned, has_reaction, reaction, reaction_date, reaction_unread, saved_peer_type, saved_peer_id, date2, ttl_period from messages where user_id = ? and (saved_peer_type = ? and saved_peer_id = ?) and date2 >= ? and deleted = 0 order by user_message_box_id asc limit ?"
		values []Messages
	)
	err = m.db.QueryRowsPartial(ctx, &values, query, userId, savedPeerType, savedPeerId, date2, limit)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectForwardSavedByOffsetDateLimit(_), error: %v", err)
		return
	}

	rList = values

	return
}

// SelectForwardSavedByOffsetDateLimitWithCB
// select user_id, user_message_box_id, dialog_id1, dialog_id2, dialog_message_id, sender_user_id, peer_type, peer_id, random_id, message_filter_type, message_data, message, mentioned, media_unread, pinned, has_reaction, reaction, reaction_date, reaction_unread, saved_peer_type, saved_peer_id, date2, ttl_period from messages where user_id = :user_id and (saved_peer_type = :saved_peer_type and saved_peer_id = :saved_peer_id) and date2 >= :date2 and deleted = 0 order by user_message_box_id asc limit :limit
func (m *defaultMessagesModel) SelectForwardSavedByOffsetDateLimitWithCB(ctx context.Context, userId int64, savedPeerType int32, savedPeerId int64, date2 int64, limit int32, cb func(sz, i int, v *Messages)) (rList []Messages, err error) {
	var (
		query  = "select user_id, user_message_box_id, dialog_id1, dialog_id2, dialog_message_id, sender_user_id, peer_type, peer_id, random_id, message_filter_type, message_data, message, mentioned, media_unread, pinned, has_reaction, reaction, reaction_date, reaction_unread, saved_peer_type, saved_peer_id, date2, ttl_period from messages where user_id = ? and (saved_peer_type = ? and saved_peer_id = ?) and date2 >= ? and deleted = 0 order by user_message_box_id asc limit ?"
		values []Messages
	)
	err = m.db.QueryRowsPartial(ctx, &values, query, userId, savedPeerType, savedPeerId, date2, limit)

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
