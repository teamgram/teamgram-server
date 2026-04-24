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
	chatInviteParticipantsFieldNames          = builder.RawFieldNames(&ChatInviteParticipants{})
	chatInviteParticipantsRows                = strings.Join(chatInviteParticipantsFieldNames, ",")
	chatInviteParticipantsRowsExpectAutoSet   = strings.Join(stringx.Remove(chatInviteParticipantsFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	chatInviteParticipantsRowsWithPlaceHolder = strings.Join(stringx.Remove(chatInviteParticipantsFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"
)

type (
	chatInviteParticipantsModel interface {
		Insert2(ctx context.Context, data *ChatInviteParticipants) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*ChatInviteParticipants, error)
		FindListByIdList(ctx context.Context, id ...int64) ([]ChatInviteParticipants, error)
		Update2(ctx context.Context, data *ChatInviteParticipants) error
		Delete2(ctx context.Context, id int64) error

		FindOneByLinkUserId(ctx context.Context, link string, userId int64) (*ChatInviteParticipants, error)
	}

	defaultChatInviteParticipantsModel struct {
		db *sqlx.DB
	}

	ChatInviteParticipants struct {
		Id         int64  `db:"id" json:"id"`
		ChatId     int64  `db:"chat_id" json:"chat_id"`
		Link       string `db:"link" json:"link"`
		UserId     int64  `db:"user_id" json:"user_id"`
		Requested  bool   `db:"requested" json:"requested"`
		ApprovedBy int64  `db:"approved_by" json:"approved_by"`
		Date2      int64  `db:"date2" json:"date2"`
		Deleted    bool   `db:"deleted" json:"deleted"`
	}
)

func newChatInviteParticipantsModel(db *sqlx.DB) *defaultChatInviteParticipantsModel {
	return &defaultChatInviteParticipantsModel{
		db: db,
	}
}

func (m *defaultChatInviteParticipantsModel) Insert2(ctx context.Context, data *ChatInviteParticipants) (sql.Result, error) {
	query := fmt.Sprintf("insert into `chat_invite_participants` (%s) values (?, ?, ?, ?, ?, ?, ?)", chatInviteParticipantsRowsExpectAutoSet)

	r, err := m.db.Exec(ctx, query, data.ChatId, data.Link, data.UserId, data.Requested, data.ApprovedBy, data.Date2, data.Deleted)
	if err != nil {
		return nil, fmt.Errorf("chat_invite_participants.Insert2 exec: %w", err)
	}

	return r, nil
}

func (m *defaultChatInviteParticipantsModel) Delete2(ctx context.Context, id int64) error {
	query := "delete from `chat_invite_participants` where `id` = ?"

	_, err := m.db.Exec(ctx, query, id)
	if err != nil {
		return fmt.Errorf("chat_invite_participants.Delete2 exec: %w", err)
	}

	return nil
}

func (m *defaultChatInviteParticipantsModel) FindOne(ctx context.Context, id int64) (*ChatInviteParticipants, error) {
	query := fmt.Sprintf("select %s from chat_invite_participants where id = ? limit 1", chatInviteParticipantsRows)
	var resp ChatInviteParticipants

	err := m.db.QueryRowPartial(ctx, &resp, query, id)
	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return nil, nil
		}
		return nil, fmt.Errorf("chat_invite_participants.FindOne: %w", err)
	}

	return &resp, nil
}

func (m *defaultChatInviteParticipantsModel) FindListByIdList(ctx context.Context, id ...int64) ([]ChatInviteParticipants, error) {
	if len(id) == 0 {
		return []ChatInviteParticipants{}, nil
	}

	query := fmt.Sprintf("select %s from chat_invite_participants where id in (%s)", chatInviteParticipantsRows, sqlx.InInt64List(id))

	var resp []ChatInviteParticipants
	err := m.db.QueryRowsPartial(ctx, &resp, query)
	if err != nil {
		return nil, fmt.Errorf("chat_invite_participants.FindListByIdList: %w", err)
	}

	return resp, nil
}

func (m *defaultChatInviteParticipantsModel) Update2(ctx context.Context, data *ChatInviteParticipants) error {
	query := fmt.Sprintf("update `chat_invite_participants` set %s where `id` = ?", chatInviteParticipantsRowsWithPlaceHolder)

	_, err := m.db.Exec(ctx, query, data.ChatId, data.Link, data.UserId, data.Requested, data.ApprovedBy, data.Date2, data.Deleted, data.Id)
	if err != nil {
		return fmt.Errorf("chat_invite_participants.Update2 exec: %w", err)
	}

	return nil
}

func (m *defaultChatInviteParticipantsModel) FindOneByLinkUserId(ctx context.Context, link string, userId int64) (*ChatInviteParticipants, error) {
	query := fmt.Sprintf("select %s from chat_invite_participants where link = ? AND user_id = ? limit 1", chatInviteParticipantsRows)
	var resp ChatInviteParticipants

	err := m.db.QueryRowPartial(ctx, &resp, query, link, userId)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return nil, nil
		}
		return nil, fmt.Errorf("chat_invite_participants.FindOneByLinkUserId: %w", err)
	}

	return &resp, nil
}
