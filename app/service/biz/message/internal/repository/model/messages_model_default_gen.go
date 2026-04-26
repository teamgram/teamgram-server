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

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	messagesFieldNames          = builder.RawFieldNames(&Messages{})
	messagesRows                = strings.Join(messagesFieldNames, ",")
	messagesRowsExpectAutoSet   = strings.Join(stringx.Remove(messagesFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	messagesRowsWithPlaceHolder = strings.Join(stringx.Remove(messagesFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"
)

type (
	messagesModel interface {
		Insert2(ctx context.Context, data *Messages) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*Messages, error)
		FindListByIdList(ctx context.Context, id ...int64) ([]Messages, error)
		Update2(ctx context.Context, data *Messages) error
		Delete2(ctx context.Context, id int64) error

		FindOneByUserIdUserMessageBoxId(ctx context.Context, userId int64, userMessageBoxId int32) (*Messages, error)
	}

	defaultMessagesModel struct {
		db *sqlx.DB
	}

	Messages struct {
		Id                int64  `db:"id" json:"id"`
		UserId            int64  `db:"user_id" json:"user_id"`
		UserMessageBoxId  int32  `db:"user_message_box_id" json:"user_message_box_id"`
		DialogId1         int64  `db:"dialog_id1" json:"dialog_id1"`
		DialogId2         int64  `db:"dialog_id2" json:"dialog_id2"`
		DialogMessageId   int64  `db:"dialog_message_id" json:"dialog_message_id"`
		SenderUserId      int64  `db:"sender_user_id" json:"sender_user_id"`
		PeerType          int32  `db:"peer_type" json:"peer_type"`
		PeerId            int64  `db:"peer_id" json:"peer_id"`
		RandomId          int64  `db:"random_id" json:"random_id"`
		MessageFilterType int32  `db:"message_filter_type" json:"message_filter_type"`
		MessageData       string `db:"message_data" json:"message_data"`
		Message           string `db:"message" json:"message"`
		Mentioned         bool   `db:"mentioned" json:"mentioned"`
		MediaUnread       bool   `db:"media_unread" json:"media_unread"`
		Pinned            bool   `db:"pinned" json:"pinned"`
		HasReaction       bool   `db:"has_reaction" json:"has_reaction"`
		Reaction          string `db:"reaction" json:"reaction"`
		ReactionDate      int64  `db:"reaction_date" json:"reaction_date"`
		ReactionUnread    bool   `db:"reaction_unread" json:"reaction_unread"`
		Date2             int64  `db:"date2" json:"date2"`
		TtlPeriod         int32  `db:"ttl_period" json:"ttl_period"`
		SavedPeerType     int32  `db:"saved_peer_type" json:"saved_peer_type"`
		SavedPeerId       int64  `db:"saved_peer_id" json:"saved_peer_id"`
		OutboxReadDate    int64  `db:"outbox_read_date" json:"outbox_read_date"`
		Deleted           bool   `db:"deleted" json:"deleted"`
	}
)

func newMessagesModel(db *sqlx.DB) *defaultMessagesModel {
	return &defaultMessagesModel{
		db: db,
	}
}

func (m *defaultMessagesModel) Insert2(ctx context.Context, data *Messages) (sql.Result, error) {
	query := fmt.Sprintf("insert into `messages` (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", messagesRowsExpectAutoSet)

	r, err := m.db.Exec(ctx, query, data.UserId, data.UserMessageBoxId, data.DialogId1, data.DialogId2, data.DialogMessageId, data.SenderUserId, data.PeerType, data.PeerId, data.RandomId, data.MessageFilterType, data.MessageData, data.Message, data.Mentioned, data.MediaUnread, data.Pinned, data.HasReaction, data.Reaction, data.ReactionDate, data.ReactionUnread, data.Date2, data.TtlPeriod, data.SavedPeerType, data.SavedPeerId, data.OutboxReadDate, data.Deleted)
	if err != nil {
		return nil, fmt.Errorf("messages.Insert2 exec: %w", err)
	}

	return r, nil
}

func (m *defaultMessagesModel) Delete2(ctx context.Context, id int64) error {
	query := "delete from `messages` where `id` = ?"

	_, err := m.db.Exec(ctx, query, id)
	if err != nil {
		return fmt.Errorf("messages.Delete2 exec: %w", err)
	}

	return nil
}

func (m *defaultMessagesModel) FindOne(ctx context.Context, id int64) (*Messages, error) {
	query := fmt.Sprintf("select %s from messages where id = ? limit 1", messagesRows)
	var resp Messages

	err := m.db.QueryRowPartial(ctx, &resp, query, id)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return nil, &NotFoundError{
				Resource: "messages",
				Key:      fmt.Sprintf("id=%v", id),
				Cause:    err,
			}
		}
		return nil, fmt.Errorf("messages.FindOne: %w", err)
	}

	return &resp, nil
}

func (m *defaultMessagesModel) FindListByIdList(ctx context.Context, id ...int64) ([]Messages, error) {
	if len(id) == 0 {
		return []Messages{}, nil
	}

	query := fmt.Sprintf("select %s from messages where id in (%s)", messagesRows, sqlx.InInt64List(id))

	var resp []Messages
	err := m.db.QueryRowsPartial(ctx, &resp, query)
	if err != nil {
		return nil, fmt.Errorf("messages.FindListByIdList: %w", err)
	}

	return resp, nil
}

func (m *defaultMessagesModel) Update2(ctx context.Context, data *Messages) error {
	query := fmt.Sprintf("update `messages` set %s where `id` = ?", messagesRowsWithPlaceHolder)

	_, err := m.db.Exec(ctx, query, data.UserId, data.UserMessageBoxId, data.DialogId1, data.DialogId2, data.DialogMessageId, data.SenderUserId, data.PeerType, data.PeerId, data.RandomId, data.MessageFilterType, data.MessageData, data.Message, data.Mentioned, data.MediaUnread, data.Pinned, data.HasReaction, data.Reaction, data.ReactionDate, data.ReactionUnread, data.Date2, data.TtlPeriod, data.SavedPeerType, data.SavedPeerId, data.OutboxReadDate, data.Deleted, data.Id)
	if err != nil {
		return fmt.Errorf("messages.Update2 exec: %w", err)
	}

	return nil
}

func (m *defaultMessagesModel) FindOneByUserIdUserMessageBoxId(ctx context.Context, userId int64, userMessageBoxId int32) (*Messages, error) {
	query := fmt.Sprintf("select %s from messages where user_id = ? AND user_message_box_id = ? limit 1", messagesRows)
	var resp Messages

	err := m.db.QueryRowPartial(ctx, &resp, query, userId, userMessageBoxId)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return nil, &NotFoundError{
				Resource: "messages",
				Key:      fmt.Sprintf("user_id=%v,user_message_box_id=%v", userId, userMessageBoxId),
				Cause:    err,
			}
		}
		return nil, fmt.Errorf("messages.FindOneByUserIdUserMessageBoxId: %w", err)
	}

	return &resp, nil
}
