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
		err = fmt.Errorf("messages.InsertOrReturnId named exec: %w", err)
		return
	}

	lastInsertId, err = r.LastInsertId()
	if err != nil {
		err = fmt.Errorf("messages.InsertOrReturnId last insert id: %w", err)
		return
	}
	rowsAffected, err = r.RowsAffected()
	if err != nil {
		err = fmt.Errorf("messages.InsertOrReturnId rows affected: %w", err)
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
		err = fmt.Errorf("messages.InsertOrReturnIdTx named exec: %w", err)
		return
	}

	lastInsertId, err = r.LastInsertId()
	if err != nil {
		err = fmt.Errorf("messages.InsertOrReturnIdTx last insert id: %w", err)
		return
	}
	rowsAffected, err = r.RowsAffected()
	if err != nil {
		err = fmt.Errorf("messages.InsertOrReturnIdTx rows affected: %w", err)
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
		if errors.Is(err, sqlx.ErrNotFound) {
			return nil, &NotFoundError{
				Resource: "messages",
				Key:      fmt.Sprintf("sender_user_id=%v,random_id=%v", senderUserId, randomId),
				Cause:    err,
			}
		}
		err = fmt.Errorf("messages.SelectByRandomId: %w", err)
		return
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
		if errors.Is(err, sqlx.ErrNotFound) {
			rList = []Messages{}
			err = nil
			return
		}
		err = fmt.Errorf("messages.SelectByMessageIdList: %w", err)
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
		if errors.Is(err, sqlx.ErrNotFound) {
			rList = []Messages{}
			err = nil
			return
		}
		err = fmt.Errorf("messages.SelectByMessageIdListWithCB: %w", err)
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
		if errors.Is(err, sqlx.ErrNotFound) {
			return nil, &NotFoundError{
				Resource: "messages",
				Key:      fmt.Sprintf("user_id=%v,user_message_box_id=%v", userId, userMessageBoxId),
				Cause:    err,
			}
		}
		err = fmt.Errorf("messages.SelectByMessageId: %w", err)
		return
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
		if errors.Is(err, sqlx.ErrNotFound) {
			rList = []Messages{}
			err = nil
			return
		}
		err = fmt.Errorf("messages.SelectByMessageDataIdList: %w", err)
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
		if errors.Is(err, sqlx.ErrNotFound) {
			rList = []Messages{}
			err = nil
			return
		}
		err = fmt.Errorf("messages.SelectByMessageDataIdListWithCB: %w", err)
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
		if errors.Is(err, sqlx.ErrNotFound) {
			return nil, &NotFoundError{
				Resource: "messages",
				Key:      fmt.Sprintf("user_id=%v,dialog_message_id=%v", userId, dialogMessageId),
				Cause:    err,
			}
		}
		err = fmt.Errorf("messages.SelectByMessageDataId: %w", err)
		return
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
		if errors.Is(err, sqlx.ErrNotFound) {
			rList = []Messages{}
			err = nil
			return
		}
		err = fmt.Errorf("messages.SelectByMessageDataIdUserIdList: %w", err)
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
		if errors.Is(err, sqlx.ErrNotFound) {
			rList = []Messages{}
			err = nil
			return
		}
		err = fmt.Errorf("messages.SelectByMessageDataIdUserIdListWithCB: %w", err)
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
		if errors.Is(err, sqlx.ErrNotFound) {
			rList = []Messages{}
			err = nil
			return
		}
		err = fmt.Errorf("messages.SelectBackwardByOffsetIdLimit: %w", err)
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
		if errors.Is(err, sqlx.ErrNotFound) {
			rList = []Messages{}
			err = nil
			return
		}
		err = fmt.Errorf("messages.SelectBackwardByOffsetIdLimitWithCB: %w", err)
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
		if errors.Is(err, sqlx.ErrNotFound) {
			rList = []Messages{}
			err = nil
			return
		}
		err = fmt.Errorf("messages.SelectForwardByOffsetIdLimit: %w", err)
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
		if errors.Is(err, sqlx.ErrNotFound) {
			rList = []Messages{}
			err = nil
			return
		}
		err = fmt.Errorf("messages.SelectForwardByOffsetIdLimitWithCB: %w", err)
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
		if errors.Is(err, sqlx.ErrNotFound) {
			rList = []Messages{}
			err = nil
			return
		}
		err = fmt.Errorf("messages.SelectBackwardByOffsetDateLimit: %w", err)
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
		if errors.Is(err, sqlx.ErrNotFound) {
			rList = []Messages{}
			err = nil
			return
		}
		err = fmt.Errorf("messages.SelectBackwardByOffsetDateLimitWithCB: %w", err)
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
		if errors.Is(err, sqlx.ErrNotFound) {
			rList = []Messages{}
			err = nil
			return
		}
		err = fmt.Errorf("messages.SelectForwardByOffsetDateLimit: %w", err)
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
		if errors.Is(err, sqlx.ErrNotFound) {
			rList = []Messages{}
			err = nil
			return
		}
		err = fmt.Errorf("messages.SelectForwardByOffsetDateLimitWithCB: %w", err)
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
		if errors.Is(err, sqlx.ErrNotFound) {
			return nil, &NotFoundError{
				Resource: "messages",
				Key:      fmt.Sprintf("peerId=%v,user_id=%v,user_message_box_id=%v", peerId, userId, userMessageBoxId),
				Cause:    err,
			}
		}
		err = fmt.Errorf("messages.SelectPeerUserMessageId: %w", err)
		return
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
		if errors.Is(err, sqlx.ErrNotFound) {
			return nil, &NotFoundError{
				Resource: "messages",
				Key:      fmt.Sprintf("peerId=%v,user_id=%v,user_message_box_id=%v", peerId, userId, userMessageBoxId),
				Cause:    err,
			}
		}
		err = fmt.Errorf("messages.SelectPeerUserMessage: %w", err)
		return
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
		if errors.Is(err, sqlx.ErrNotFound) {
			err = &NotFoundError{
				Resource: "messages",
				Key:      fmt.Sprintf("user_id=%v,dialog_id1=%v,dialog_id2=%v", userId, dialogId1, dialogId2),
				Cause:    err,
			}
			return
		}
		err = fmt.Errorf("messages.SelectDialogLastMessageId: %w", err)
		return
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
		if errors.Is(err, sqlx.ErrNotFound) {
			err = &NotFoundError{
				Resource: "messages",
				Key:      fmt.Sprintf("user_id=%v,dialog_id1=%v,dialog_id2=%v,idList=%v", userId, dialogId1, dialogId2, idList),
				Cause:    err,
			}
			return
		}
		err = fmt.Errorf("messages.SelectDialogLastMessageIdNotIdList: %w", err)
		return
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
		if errors.Is(err, sqlx.ErrNotFound) {
			rList = []Messages{}
			err = nil
			return
		}
		err = fmt.Errorf("messages.SelectDialogsByMessageIdList: %w", err)
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
		if errors.Is(err, sqlx.ErrNotFound) {
			rList = []Messages{}
			err = nil
			return
		}
		err = fmt.Errorf("messages.SelectDialogsByMessageIdListWithCB: %w", err)
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
		if errors.Is(err, sqlx.ErrNotFound) {
			rList = []Messages{}
			err = nil
			return
		}
		err = fmt.Errorf("messages.SelectDialogLastMessageList: %w", err)
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
		if errors.Is(err, sqlx.ErrNotFound) {
			rList = []Messages{}
			err = nil
			return
		}
		err = fmt.Errorf("messages.SelectDialogLastMessageListWithCB: %w", err)
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
		err = fmt.Errorf("messages.DeleteMessagesByMessageIdList exec: %w", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		err = fmt.Errorf("messages.DeleteMessagesByMessageIdList rows affected: %w", err)
		return
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
		err = fmt.Errorf("messages.DeleteMessagesByMessageIdListTx exec: %w", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		err = fmt.Errorf("messages.DeleteMessagesByMessageIdListTx rows affected: %w", err)
		return
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
		if errors.Is(err, sqlx.ErrNotFound) {
			rList = []Messages{}
			err = nil
			return
		}
		err = fmt.Errorf("messages.SelectDialogMessageIdList: %w", err)
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
		if errors.Is(err, sqlx.ErrNotFound) {
			rList = []Messages{}
			err = nil
			return
		}
		err = fmt.Errorf("messages.SelectDialogMessageIdListWithCB: %w", err)
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
		err = fmt.Errorf("messages.UpdateMediaUnread exec: %w", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		err = fmt.Errorf("messages.UpdateMediaUnread rows affected: %w", err)
		return
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
		err = fmt.Errorf("messages.UpdateMediaUnreadTx exec: %w", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		err = fmt.Errorf("messages.UpdateMediaUnreadTx rows affected: %w", err)
		return
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
		err = fmt.Errorf("messages.UpdateMentionedAndMediaUnread exec: %w", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		err = fmt.Errorf("messages.UpdateMentionedAndMediaUnread rows affected: %w", err)
		return
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
		err = fmt.Errorf("messages.UpdateMentionedAndMediaUnreadTx exec: %w", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		err = fmt.Errorf("messages.UpdateMentionedAndMediaUnreadTx rows affected: %w", err)
		return
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
		if errors.Is(err, sqlx.ErrNotFound) {
			rList = []Messages{}
			err = nil
			return
		}
		err = fmt.Errorf("messages.SelectByMediaType: %w", err)
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
		if errors.Is(err, sqlx.ErrNotFound) {
			rList = []Messages{}
			err = nil
			return
		}
		err = fmt.Errorf("messages.SelectByMediaTypeWithCB: %w", err)
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
		if errors.Is(err, sqlx.ErrNotFound) {
			rList = []Messages{}
			err = nil
			return
		}
		err = fmt.Errorf("messages.SelectPhoneCallList: %w", err)
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
		if errors.Is(err, sqlx.ErrNotFound) {
			rList = []Messages{}
			err = nil
			return
		}
		err = fmt.Errorf("messages.SelectPhoneCallListWithCB: %w", err)
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
		if errors.Is(err, sqlx.ErrNotFound) {
			rList = []Messages{}
			err = nil
			return
		}
		err = fmt.Errorf("messages.Search: %w", err)
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
		if errors.Is(err, sqlx.ErrNotFound) {
			rList = []Messages{}
			err = nil
			return
		}
		err = fmt.Errorf("messages.SearchWithCB: %w", err)
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
		if errors.Is(err, sqlx.ErrNotFound) {
			rList = []Messages{}
			err = nil
			return
		}
		err = fmt.Errorf("messages.SearchGlobal: %w", err)
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
		if errors.Is(err, sqlx.ErrNotFound) {
			rList = []Messages{}
			err = nil
			return
		}
		err = fmt.Errorf("messages.SearchGlobalWithCB: %w", err)
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
		if errors.Is(err, sqlx.ErrNotFound) {
			rList = []Messages{}
			err = nil
			return
		}
		err = fmt.Errorf("messages.SelectBackwardUnreadMentionsByOffsetIdLimit: %w", err)
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
		if errors.Is(err, sqlx.ErrNotFound) {
			rList = []Messages{}
			err = nil
			return
		}
		err = fmt.Errorf("messages.SelectBackwardUnreadMentionsByOffsetIdLimitWithCB: %w", err)
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
		if errors.Is(err, sqlx.ErrNotFound) {
			rList = []Messages{}
			err = nil
			return
		}
		err = fmt.Errorf("messages.SelectForwardUnreadMentionsByOffsetIdLimit: %w", err)
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
		if errors.Is(err, sqlx.ErrNotFound) {
			rList = []Messages{}
			err = nil
			return
		}
		err = fmt.Errorf("messages.SelectForwardUnreadMentionsByOffsetIdLimitWithCB: %w", err)
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
		if errors.Is(err, sqlx.ErrNotFound) {
			rList = []Messages{}
			err = nil
			return
		}
		err = fmt.Errorf("messages.SelectPinnedList: %w", err)
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
		if errors.Is(err, sqlx.ErrNotFound) {
			rList = []Messages{}
			err = nil
			return
		}
		err = fmt.Errorf("messages.SelectPinnedListWithCB: %w", err)
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
		if errors.Is(err, sqlx.ErrNotFound) {
			rList = []int32{}
			err = nil
			return
		}
		err = fmt.Errorf("messages.SelectLastTwoPinnedList: %w", err)
	}

	return
}

// SelectLastTwoPinnedListWithCB
// select user_message_box_id from messages where user_id = :user_id and (dialog_id1 = :dialog_id1 and dialog_id2 = :dialog_id2) and pinned = 1 and deleted = 0 order by user_message_box_id desc limit 2
func (m *defaultMessagesModel) SelectLastTwoPinnedListWithCB(ctx context.Context, userId int64, dialogId1 int64, dialogId2 int64, cb func(sz, i int, v int32)) (rList []int32, err error) {
	var query = "select user_message_box_id from messages where user_id = ? and (dialog_id1 = ? and dialog_id2 = ?) and pinned = 1 and deleted = 0 order by user_message_box_id desc limit 2"
	err = m.db.QueryRowsPartial(ctx, &rList, query, userId, dialogId1, dialogId2)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			rList = []int32{}
			err = nil
			return
		}
		err = fmt.Errorf("messages.SelectLastTwoPinnedListWithCB: %w", err)
		return
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
		err = fmt.Errorf("messages.UpdatePinned exec: %w", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		err = fmt.Errorf("messages.UpdatePinned rows affected: %w", err)
		return
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
		err = fmt.Errorf("messages.UpdatePinnedTx exec: %w", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		err = fmt.Errorf("messages.UpdatePinnedTx rows affected: %w", err)
		return
	}

	return
}

// SelectPinnedMessageIdList
// select user_message_box_id from messages where user_id = :user_id and (dialog_id1 = :dialog_id1 and dialog_id2 = :dialog_id2) and pinned = 1 and deleted = 0 order by user_message_box_id desc
func (m *defaultMessagesModel) SelectPinnedMessageIdList(ctx context.Context, userId int64, dialogId1 int64, dialogId2 int64) (rList []int32, err error) {
	var query = "select user_message_box_id from messages where user_id = ? and (dialog_id1 = ? and dialog_id2 = ?) and pinned = 1 and deleted = 0 order by user_message_box_id desc"
	err = m.db.QueryRowsPartial(ctx, &rList, query, userId, dialogId1, dialogId2)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			rList = []int32{}
			err = nil
			return
		}
		err = fmt.Errorf("messages.SelectPinnedMessageIdList: %w", err)
	}

	return
}

// SelectPinnedMessageIdListWithCB
// select user_message_box_id from messages where user_id = :user_id and (dialog_id1 = :dialog_id1 and dialog_id2 = :dialog_id2) and pinned = 1 and deleted = 0 order by user_message_box_id desc
func (m *defaultMessagesModel) SelectPinnedMessageIdListWithCB(ctx context.Context, userId int64, dialogId1 int64, dialogId2 int64, cb func(sz, i int, v int32)) (rList []int32, err error) {
	var query = "select user_message_box_id from messages where user_id = ? and (dialog_id1 = ? and dialog_id2 = ?) and pinned = 1 and deleted = 0 order by user_message_box_id desc"
	err = m.db.QueryRowsPartial(ctx, &rList, query, userId, dialogId1, dialogId2)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			rList = []int32{}
			err = nil
			return
		}
		err = fmt.Errorf("messages.SelectPinnedMessageIdListWithCB: %w", err)
		return
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
		err = fmt.Errorf("messages.UpdateUnPinnedByIdList exec: %w", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		err = fmt.Errorf("messages.UpdateUnPinnedByIdList rows affected: %w", err)
		return
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
		err = fmt.Errorf("messages.UpdateUnPinnedByIdListTx exec: %w", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		err = fmt.Errorf("messages.UpdateUnPinnedByIdListTx rows affected: %w", err)
		return
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
		err = fmt.Errorf("messages.UpdateEditMessage exec: %w", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		err = fmt.Errorf("messages.UpdateEditMessage rows affected: %w", err)
		return
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
		err = fmt.Errorf("messages.UpdateEditMessageTx exec: %w", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		err = fmt.Errorf("messages.UpdateEditMessageTx rows affected: %w", err)
		return
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
		err = fmt.Errorf("messages.UpdateCustomMap exec: %w", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		err = fmt.Errorf("messages.UpdateCustomMap rows affected: %w", err)
		return
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
		err = fmt.Errorf("messages.UpdateCustomMapTx exec: %w", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		err = fmt.Errorf("messages.UpdateCustomMapTx rows affected: %w", err)
		return
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
		if errors.Is(err, sqlx.ErrNotFound) {
			rList = []Messages{}
			err = nil
			return
		}
		err = fmt.Errorf("messages.SelectBackwardBySendUserIdOffsetIdLimit: %w", err)
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
		if errors.Is(err, sqlx.ErrNotFound) {
			rList = []Messages{}
			err = nil
			return
		}
		err = fmt.Errorf("messages.SelectBackwardBySendUserIdOffsetIdLimitWithCB: %w", err)
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
		if errors.Is(err, sqlx.ErrNotFound) {
			rList = []Messages{}
			err = nil
			return
		}
		err = fmt.Errorf("messages.SelectBackwardSavedByOffsetIdLimit: %w", err)
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
		if errors.Is(err, sqlx.ErrNotFound) {
			rList = []Messages{}
			err = nil
			return
		}
		err = fmt.Errorf("messages.SelectBackwardSavedByOffsetIdLimitWithCB: %w", err)
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
		if errors.Is(err, sqlx.ErrNotFound) {
			rList = []Messages{}
			err = nil
			return
		}
		err = fmt.Errorf("messages.SelectForwardSavedByOffsetIdLimit: %w", err)
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
		if errors.Is(err, sqlx.ErrNotFound) {
			rList = []Messages{}
			err = nil
			return
		}
		err = fmt.Errorf("messages.SelectForwardSavedByOffsetIdLimitWithCB: %w", err)
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
		if errors.Is(err, sqlx.ErrNotFound) {
			rList = []Messages{}
			err = nil
			return
		}
		err = fmt.Errorf("messages.SelectBackwardSavedByOffsetDateLimit: %w", err)
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
		if errors.Is(err, sqlx.ErrNotFound) {
			rList = []Messages{}
			err = nil
			return
		}
		err = fmt.Errorf("messages.SelectBackwardSavedByOffsetDateLimitWithCB: %w", err)
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
		if errors.Is(err, sqlx.ErrNotFound) {
			rList = []Messages{}
			err = nil
			return
		}
		err = fmt.Errorf("messages.SelectForwardSavedByOffsetDateLimit: %w", err)
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
		if errors.Is(err, sqlx.ErrNotFound) {
			rList = []Messages{}
			err = nil
			return
		}
		err = fmt.Errorf("messages.SelectForwardSavedByOffsetDateLimitWithCB: %w", err)
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
