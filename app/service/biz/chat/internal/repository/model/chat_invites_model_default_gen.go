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
	"fmt"
	"strings"

	"github.com/teamgram/marmota/pkg/stores/sqlx"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	chatInvitesFieldNames          = builder.RawFieldNames(&ChatInvites{})
	chatInvitesRows                = strings.Join(chatInvitesFieldNames, ",")
	chatInvitesRowsExpectAutoSet   = strings.Join(stringx.Remove(chatInvitesFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	chatInvitesRowsWithPlaceHolder = strings.Join(stringx.Remove(chatInvitesFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"

	cacheTChatInvitesIdPrefix = "cache:t:chat_invites:id:"

	cacheChatInvitesIdPrefix = "cache#ChatInvites#id"

	cacheChatInvitesLinkPrefix = "cache#Link"
)

type (
	chatInvitesModel interface {
		Insert2(ctx context.Context, data *ChatInvites) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*ChatInvites, error)
		FindListByIdList(ctx context.Context, id ...int64) ([]ChatInvites, error)
		Update2(ctx context.Context, data *ChatInvites) error
		Delete2(ctx context.Context, id int64) error

		FindOneByLink(ctx context.Context, link string) (*ChatInvites, error)
		FindListByLinkList(ctx context.Context, link ...string) ([]ChatInvites, error)
	}

	defaultChatInvitesModel struct {
		db *sqlx.DB
	}

	ChatInvites struct {
		Id            int64  `db:"id" json:"id"`
		ChatId        int64  `db:"chat_id" json:"chat_id"`
		AdminId       int64  `db:"admin_id" json:"admin_id"`
		MigratedToId  int64  `db:"migrated_to_id" json:"migrated_to_id"`
		Link          string `db:"link" json:"link"`
		Permanent     bool   `db:"permanent" json:"permanent"`
		Revoked       bool   `db:"revoked" json:"revoked"`
		RequestNeeded bool   `db:"request_needed" json:"request_needed"`
		StartDate     int64  `db:"start_date" json:"start_date"`
		ExpireDate    int64  `db:"expire_date" json:"expire_date"`
		UsageLimit    int32  `db:"usage_limit" json:"usage_limit"`
		Usage2        int32  `db:"usage2" json:"usage2"`
		Requested     int32  `db:"requested" json:"requested"`
		Title         string `db:"title" json:"title"`
		Date2         int64  `db:"date2" json:"date2"`
		State         int32  `db:"state" json:"state"`
	}
)

func newChatInvitesModel(db *sqlx.DB) *defaultChatInvitesModel {
	return &defaultChatInvitesModel{
		db: db,
	}
}

func (m *defaultChatInvitesModel) Insert2(ctx context.Context, data *ChatInvites) (sql.Result, error) {
	query := fmt.Sprintf("insert into `chat_invites` (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", chatInvitesRowsExpectAutoSet)
	return m.db.Exec(ctx, query, data.ChatId, data.AdminId, data.MigratedToId, data.Link, data.Permanent, data.Revoked, data.RequestNeeded, data.StartDate, data.ExpireDate, data.UsageLimit, data.Usage2, data.Requested, data.Title, data.Date2, data.State)
}

func (m *defaultChatInvitesModel) Delete2(ctx context.Context, id int64) error {
	query := "delete from `chat_invites` where `id` = ?"
	_, err := m.db.Exec(ctx, query, id)
	return err
}

func (m *defaultChatInvitesModel) FindOne(ctx context.Context, id int64) (*ChatInvites, error) {
	query := fmt.Sprintf("select %s from chat_invites where id = ? limit 1", chatInvitesRows)
	var resp ChatInvites
	err := m.db.QueryRowPartial(ctx, &resp, query, id)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (m *defaultChatInvitesModel) FindListByIdList(ctx context.Context, id ...int64) ([]ChatInvites, error) {
	if len(id) == 0 {
		return []ChatInvites{}, nil
	}

	query := fmt.Sprintf("select %s from chat_invites where id in (%s)", chatInvitesRows, sqlx.InInt64List(id))

	var resp []ChatInvites
	err := m.db.QueryRowsPartial(ctx, &resp, query)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (m *defaultChatInvitesModel) Update2(ctx context.Context, data *ChatInvites) error {
	query := fmt.Sprintf("update `chat_invites` set %s where `id` = ?", chatInvitesRowsWithPlaceHolder)
	_, err := m.db.Exec(ctx, query, data.ChatId, data.AdminId, data.MigratedToId, data.Link, data.Permanent, data.Revoked, data.RequestNeeded, data.StartDate, data.ExpireDate, data.UsageLimit, data.Usage2, data.Requested, data.Title, data.Date2, data.State, data.Id)
	return err
}

func (m *defaultChatInvitesModel) FindOneByLink(ctx context.Context, link string) (*ChatInvites, error) {
	query := fmt.Sprintf("select %s from chat_invites where link = ? limit 1", chatInvitesRows)
	var resp ChatInvites
	err := m.db.QueryRowPartial(ctx, &resp, query, link)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (m *defaultChatInvitesModel) FindListByLinkList(ctx context.Context, link ...string) ([]ChatInvites, error) {
	if len(link) == 0 {
		return []ChatInvites{}, nil
	}

	query := fmt.Sprintf("select %s from chat_invites where link in (%s)", chatInvitesRows, sqlx.InStringList(link))
	var resp []ChatInvites
	err := m.db.QueryRowsPartial(ctx, &resp, query)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (m *defaultChatInvitesModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s#%v", cacheChatInvitesIdPrefix, primary)
}

func (m *defaultChatInvitesModel) queryPrimary(ctx context.Context, v interface{}, primary interface{}) error {
	query := fmt.Sprintf("select %s from chat_invites where id = ? limit 1", chatInvitesRows)
	return m.db.QueryRowPartial(ctx, v, query, primary)
}
