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
	chatParticipantsFieldNames          = builder.RawFieldNames(&ChatParticipants{})
	chatParticipantsRows                = strings.Join(chatParticipantsFieldNames, ",")
	chatParticipantsRowsExpectAutoSet   = strings.Join(stringx.Remove(chatParticipantsFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	chatParticipantsRowsWithPlaceHolder = strings.Join(stringx.Remove(chatParticipantsFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"
)

type (
	chatParticipantsModel interface {
		Insert2(ctx context.Context, data *ChatParticipants) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*ChatParticipants, error)
		FindListByIdList(ctx context.Context, id ...int64) ([]ChatParticipants, error)
		Update2(ctx context.Context, data *ChatParticipants) error
		Delete2(ctx context.Context, id int64) error

		FindOneByChatIdUserId(ctx context.Context, chatId int64, userId int64) (*ChatParticipants, error)
	}

	defaultChatParticipantsModel struct {
		db *sqlx.DB
	}

	ChatParticipants struct {
		Id                             int64  `db:"id" json:"id"`
		ChatId                         int64  `db:"chat_id" json:"chat_id"`
		UserId                         int64  `db:"user_id" json:"user_id"`
		ParticipantType                int32  `db:"participant_type" json:"participant_type"`
		Link                           string `db:"link" json:"link"`
		Usage2                         int32  `db:"usage2" json:"usage2"`
		AdminRights                    int32  `db:"admin_rights" json:"admin_rights"`
		InviterUserId                  int64  `db:"inviter_user_id" json:"inviter_user_id"`
		InvitedAt                      int64  `db:"invited_at" json:"invited_at"`
		KickedAt                       int64  `db:"kicked_at" json:"kicked_at"`
		LeftAt                         int64  `db:"left_at" json:"left_at"`
		GroupcallDefaultJoinAsPeerType int32  `db:"groupcall_default_join_as_peer_type" json:"groupcall_default_join_as_peer_type"`
		GroupcallDefaultJoinAsPeerId   int64  `db:"groupcall_default_join_as_peer_id" json:"groupcall_default_join_as_peer_id"`
		IsBot                          bool   `db:"is_bot" json:"is_bot"`
		State                          int32  `db:"state" json:"state"`
		Date2                          int64  `db:"date2" json:"date2"`
	}
)

func newChatParticipantsModel(db *sqlx.DB) *defaultChatParticipantsModel {
	return &defaultChatParticipantsModel{
		db: db,
	}
}

func (m *defaultChatParticipantsModel) Insert2(ctx context.Context, data *ChatParticipants) (sql.Result, error) {
	query := fmt.Sprintf("insert into `chat_participants` (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", chatParticipantsRowsExpectAutoSet)

	r, err := m.db.Exec(ctx, query, data.ChatId, data.UserId, data.ParticipantType, data.Link, data.Usage2, data.AdminRights, data.InviterUserId, data.InvitedAt, data.KickedAt, data.LeftAt, data.GroupcallDefaultJoinAsPeerType, data.GroupcallDefaultJoinAsPeerId, data.IsBot, data.State, data.Date2)
	if err != nil {
		return nil, fmt.Errorf("chat_participants.Insert2 exec: %w", err)
	}

	return r, nil
}

func (m *defaultChatParticipantsModel) Delete2(ctx context.Context, id int64) error {
	query := "delete from `chat_participants` where `id` = ?"

	_, err := m.db.Exec(ctx, query, id)
	if err != nil {
		return fmt.Errorf("chat_participants.Delete2 exec: %w", err)
	}

	return nil
}

func (m *defaultChatParticipantsModel) FindOne(ctx context.Context, id int64) (*ChatParticipants, error) {
	query := fmt.Sprintf("select %s from chat_participants where id = ? limit 1", chatParticipantsRows)
	var resp ChatParticipants

	err := m.db.QueryRowPartial(ctx, &resp, query, id)
	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return nil, nil
		}
		return nil, fmt.Errorf("chat_participants.FindOne: %w", err)
	}

	return &resp, nil
}

func (m *defaultChatParticipantsModel) FindListByIdList(ctx context.Context, id ...int64) ([]ChatParticipants, error) {
	if len(id) == 0 {
		return []ChatParticipants{}, nil
	}

	query := fmt.Sprintf("select %s from chat_participants where id in (%s)", chatParticipantsRows, sqlx.InInt64List(id))

	var resp []ChatParticipants
	err := m.db.QueryRowsPartial(ctx, &resp, query)
	if err != nil {
		return nil, fmt.Errorf("chat_participants.FindListByIdList: %w", err)
	}

	return resp, nil
}

func (m *defaultChatParticipantsModel) Update2(ctx context.Context, data *ChatParticipants) error {
	query := fmt.Sprintf("update `chat_participants` set %s where `id` = ?", chatParticipantsRowsWithPlaceHolder)

	_, err := m.db.Exec(ctx, query, data.ChatId, data.UserId, data.ParticipantType, data.Link, data.Usage2, data.AdminRights, data.InviterUserId, data.InvitedAt, data.KickedAt, data.LeftAt, data.GroupcallDefaultJoinAsPeerType, data.GroupcallDefaultJoinAsPeerId, data.IsBot, data.State, data.Date2, data.Id)
	if err != nil {
		return fmt.Errorf("chat_participants.Update2 exec: %w", err)
	}

	return nil
}

func (m *defaultChatParticipantsModel) FindOneByChatIdUserId(ctx context.Context, chatId int64, userId int64) (*ChatParticipants, error) {
	query := fmt.Sprintf("select %s from chat_participants where chat_id = ? AND user_id = ? limit 1", chatParticipantsRows)
	var resp ChatParticipants

	err := m.db.QueryRowPartial(ctx, &resp, query, chatId, userId)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return nil, nil
		}
		return nil, fmt.Errorf("chat_participants.FindOneByChatIdUserId: %w", err)
	}

	return &resp, nil
}
