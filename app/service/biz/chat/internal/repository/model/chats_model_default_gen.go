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
	chatsFieldNames          = builder.RawFieldNames(&Chats{})
	chatsRows                = strings.Join(chatsFieldNames, ",")
	chatsRowsExpectAutoSet   = strings.Join(stringx.Remove(chatsFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	chatsRowsWithPlaceHolder = strings.Join(stringx.Remove(chatsFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"
)

type (
	chatsModel interface {
		Insert2(ctx context.Context, data *Chats) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*Chats, error)
		FindListByIdList(ctx context.Context, id ...int64) ([]Chats, error)
		Update2(ctx context.Context, data *Chats) error
		Delete2(ctx context.Context, id int64) error
	}

	defaultChatsModel struct {
		db *sqlx.DB
	}

	Chats struct {
		Id                     int64  `db:"id" json:"id"`
		CreatorUserId          int64  `db:"creator_user_id" json:"creator_user_id"`
		AccessHash             int64  `db:"access_hash" json:"access_hash"`
		RandomId               int64  `db:"random_id" json:"random_id"`
		ParticipantCount       int32  `db:"participant_count" json:"participant_count"`
		Title                  string `db:"title" json:"title"`
		About                  string `db:"about" json:"about"`
		PhotoId                int64  `db:"photo_id" json:"photo_id"`
		DefaultBannedRights    int64  `db:"default_banned_rights" json:"default_banned_rights"`
		MigratedToId           int64  `db:"migrated_to_id" json:"migrated_to_id"`
		MigratedToAccessHash   int64  `db:"migrated_to_access_hash" json:"migrated_to_access_hash"`
		AvailableReactionsType int32  `db:"available_reactions_type" json:"available_reactions_type"`
		AvailableReactions     string `db:"available_reactions" json:"available_reactions"`
		Deactivated            bool   `db:"deactivated" json:"deactivated"`
		Noforwards             bool   `db:"noforwards" json:"noforwards"`
		TtlPeriod              int32  `db:"ttl_period" json:"ttl_period"`
		Version                int32  `db:"version" json:"version"`
		Date                   int64  `db:"date" json:"date"`
	}
)

func newChatsModel(db *sqlx.DB) *defaultChatsModel {
	return &defaultChatsModel{
		db: db,
	}
}

func (m *defaultChatsModel) Insert2(ctx context.Context, data *Chats) (sql.Result, error) {
	query := fmt.Sprintf("insert into `chats` (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", chatsRowsExpectAutoSet)

	r, err := m.db.Exec(ctx, query, data.CreatorUserId, data.AccessHash, data.RandomId, data.ParticipantCount, data.Title, data.About, data.PhotoId, data.DefaultBannedRights, data.MigratedToId, data.MigratedToAccessHash, data.AvailableReactionsType, data.AvailableReactions, data.Deactivated, data.Noforwards, data.TtlPeriod, data.Version, data.Date)
	if err != nil {
		return nil, fmt.Errorf("chats.Insert2 exec: %w", err)
	}

	return r, nil
}

func (m *defaultChatsModel) Delete2(ctx context.Context, id int64) error {
	query := "delete from `chats` where `id` = ?"

	_, err := m.db.Exec(ctx, query, id)
	if err != nil {
		return fmt.Errorf("chats.Delete2 exec: %w", err)
	}

	return nil
}

func (m *defaultChatsModel) FindOne(ctx context.Context, id int64) (*Chats, error) {
	query := fmt.Sprintf("select %s from chats where id = ? limit 1", chatsRows)
	var resp Chats

	err := m.db.QueryRowPartial(ctx, &resp, query, id)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return nil, &NotFoundError{
				Resource: "chats",
				Key:      fmt.Sprintf("id=%v", id),
				Cause:    err,
			}
		}
		return nil, fmt.Errorf("chats.FindOne: %w", err)
	}

	return &resp, nil
}

func (m *defaultChatsModel) FindListByIdList(ctx context.Context, id ...int64) ([]Chats, error) {
	if len(id) == 0 {
		return []Chats{}, nil
	}

	query := fmt.Sprintf("select %s from chats where id in (%s)", chatsRows, sqlx.InInt64List(id))

	var resp []Chats
	err := m.db.QueryRowsPartial(ctx, &resp, query)
	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return []Chats{}, nil
		}
		return nil, fmt.Errorf("chats.FindListByIdList: %w", err)
	}

	return resp, nil
}

func (m *defaultChatsModel) Update2(ctx context.Context, data *Chats) error {
	query := fmt.Sprintf("update `chats` set %s where `id` = ?", chatsRowsWithPlaceHolder)

	_, err := m.db.Exec(ctx, query, data.CreatorUserId, data.AccessHash, data.RandomId, data.ParticipantCount, data.Title, data.About, data.PhotoId, data.DefaultBannedRights, data.MigratedToId, data.MigratedToAccessHash, data.AvailableReactionsType, data.AvailableReactions, data.Deactivated, data.Noforwards, data.TtlPeriod, data.Version, data.Date, data.Id)
	if err != nil {
		return fmt.Errorf("chats.Update2 exec: %w", err)
	}

	return nil
}
