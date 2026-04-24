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
	savedDialogsFieldNames          = builder.RawFieldNames(&SavedDialogs{})
	savedDialogsRows                = strings.Join(savedDialogsFieldNames, ",")
	savedDialogsRowsExpectAutoSet   = strings.Join(stringx.Remove(savedDialogsFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	savedDialogsRowsWithPlaceHolder = strings.Join(stringx.Remove(savedDialogsFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"
)

type (
	savedDialogsModel interface {
		Insert2(ctx context.Context, data *SavedDialogs) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*SavedDialogs, error)
		FindListByIdList(ctx context.Context, id ...int64) ([]SavedDialogs, error)
		Update2(ctx context.Context, data *SavedDialogs) error
		Delete2(ctx context.Context, id int64) error

		FindOneByUserIdPeerTypePeerId(ctx context.Context, userId int64, peerType int32, peerId int64) (*SavedDialogs, error)
	}

	defaultSavedDialogsModel struct {
		db *sqlx.DB
	}

	SavedDialogs struct {
		Id         int64 `db:"id" json:"id"`
		UserId     int64 `db:"user_id" json:"user_id"`
		PeerType   int32 `db:"peer_type" json:"peer_type"`
		PeerId     int64 `db:"peer_id" json:"peer_id"`
		Pinned     int64 `db:"pinned" json:"pinned"`
		TopMessage int32 `db:"top_message" json:"top_message"`
		Deleted    bool  `db:"deleted" json:"deleted"`
	}
)

func newSavedDialogsModel(db *sqlx.DB) *defaultSavedDialogsModel {
	return &defaultSavedDialogsModel{
		db: db,
	}
}

func (m *defaultSavedDialogsModel) Insert2(ctx context.Context, data *SavedDialogs) (sql.Result, error) {
	query := fmt.Sprintf("insert into `saved_dialogs` (%s) values (?, ?, ?, ?, ?, ?)", savedDialogsRowsExpectAutoSet)

	r, err := m.db.Exec(ctx, query, data.UserId, data.PeerType, data.PeerId, data.Pinned, data.TopMessage, data.Deleted)
	if err != nil {
		return nil, fmt.Errorf("saved_dialogs.Insert2 exec: %w", err)
	}

	return r, nil
}

func (m *defaultSavedDialogsModel) Delete2(ctx context.Context, id int64) error {
	query := "delete from `saved_dialogs` where `id` = ?"

	_, err := m.db.Exec(ctx, query, id)
	if err != nil {
		return fmt.Errorf("saved_dialogs.Delete2 exec: %w", err)
	}

	return nil
}

func (m *defaultSavedDialogsModel) FindOne(ctx context.Context, id int64) (*SavedDialogs, error) {
	query := fmt.Sprintf("select %s from saved_dialogs where id = ? limit 1", savedDialogsRows)
	var resp SavedDialogs

	err := m.db.QueryRowPartial(ctx, &resp, query, id)
	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return nil, nil
		}
		return nil, fmt.Errorf("saved_dialogs.FindOne: %w", err)
	}

	return &resp, nil
}

func (m *defaultSavedDialogsModel) FindListByIdList(ctx context.Context, id ...int64) ([]SavedDialogs, error) {
	if len(id) == 0 {
		return []SavedDialogs{}, nil
	}

	query := fmt.Sprintf("select %s from saved_dialogs where id in (%s)", savedDialogsRows, sqlx.InInt64List(id))

	var resp []SavedDialogs
	err := m.db.QueryRowsPartial(ctx, &resp, query)
	if err != nil {
		return nil, fmt.Errorf("saved_dialogs.FindListByIdList: %w", err)
	}

	return resp, nil
}

func (m *defaultSavedDialogsModel) Update2(ctx context.Context, data *SavedDialogs) error {
	query := fmt.Sprintf("update `saved_dialogs` set %s where `id` = ?", savedDialogsRowsWithPlaceHolder)

	_, err := m.db.Exec(ctx, query, data.UserId, data.PeerType, data.PeerId, data.Pinned, data.TopMessage, data.Deleted, data.Id)
	if err != nil {
		return fmt.Errorf("saved_dialogs.Update2 exec: %w", err)
	}

	return nil
}

func (m *defaultSavedDialogsModel) FindOneByUserIdPeerTypePeerId(ctx context.Context, userId int64, peerType int32, peerId int64) (*SavedDialogs, error) {
	query := fmt.Sprintf("select %s from saved_dialogs where user_id = ? AND peer_type = ? AND peer_id = ? limit 1", savedDialogsRows)
	var resp SavedDialogs

	err := m.db.QueryRowPartial(ctx, &resp, query, userId, peerType, peerId)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return nil, nil
		}
		return nil, fmt.Errorf("saved_dialogs.FindOneByUserIdPeerTypePeerId: %w", err)
	}

	return &resp, nil
}
