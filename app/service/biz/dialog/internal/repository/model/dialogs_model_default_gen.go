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
	dialogsFieldNames          = builder.RawFieldNames(&Dialogs{})
	dialogsRows                = strings.Join(dialogsFieldNames, ",")
	dialogsRowsExpectAutoSet   = strings.Join(stringx.Remove(dialogsFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	dialogsRowsWithPlaceHolder = strings.Join(stringx.Remove(dialogsFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"
)

type (
	dialogsModel interface {
		Insert2(ctx context.Context, data *Dialogs) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*Dialogs, error)
		FindListByIdList(ctx context.Context, id ...int64) ([]Dialogs, error)
		Update2(ctx context.Context, data *Dialogs) error
		Delete2(ctx context.Context, id int64) error

		FindOneByUserIdPeerTypePeerId(ctx context.Context, userId int64, peerType int32, peerId int64) (*Dialogs, error)

		FindOneByUserIdPeerDialogId(ctx context.Context, userId int64, peerDialogId int64) (*Dialogs, error)
	}

	defaultDialogsModel struct {
		db *sqlx.DB
	}

	Dialogs struct {
		Id                   int64  `db:"id" json:"id"`
		UserId               int64  `db:"user_id" json:"user_id"`
		PeerType             int32  `db:"peer_type" json:"peer_type"`
		PeerId               int64  `db:"peer_id" json:"peer_id"`
		PeerDialogId         int64  `db:"peer_dialog_id" json:"peer_dialog_id"`
		Pinned               int64  `db:"pinned" json:"pinned"`
		TopMessage           int32  `db:"top_message" json:"top_message"`
		PinnedMsgId          int32  `db:"pinned_msg_id" json:"pinned_msg_id"`
		ReadInboxMaxId       int32  `db:"read_inbox_max_id" json:"read_inbox_max_id"`
		ReadOutboxMaxId      int32  `db:"read_outbox_max_id" json:"read_outbox_max_id"`
		UnreadCount          int32  `db:"unread_count" json:"unread_count"`
		UnreadMentionsCount  int32  `db:"unread_mentions_count" json:"unread_mentions_count"`
		UnreadReactionsCount int32  `db:"unread_reactions_count" json:"unread_reactions_count"`
		UnreadMark           bool   `db:"unread_mark" json:"unread_mark"`
		DraftType            int32  `db:"draft_type" json:"draft_type"`
		DraftMessageData     string `db:"draft_message_data" json:"draft_message_data"`
		FolderId             int32  `db:"folder_id" json:"folder_id"`
		FolderPinned         int64  `db:"folder_pinned" json:"folder_pinned"`
		HasScheduled         bool   `db:"has_scheduled" json:"has_scheduled"`
		TtlPeriod            int32  `db:"ttl_period" json:"ttl_period"`
		ThemeEmoticon        string `db:"theme_emoticon" json:"theme_emoticon"`
		WallpaperId          int64  `db:"wallpaper_id" json:"wallpaper_id"`
		WallpaperOverridden  bool   `db:"wallpaper_overridden" json:"wallpaper_overridden"`
		Date2                int64  `db:"date2" json:"date2"`
		Deleted              bool   `db:"deleted" json:"deleted"`
	}
)

func newDialogsModel(db *sqlx.DB) *defaultDialogsModel {
	return &defaultDialogsModel{
		db: db,
	}
}

func (m *defaultDialogsModel) Insert2(ctx context.Context, data *Dialogs) (sql.Result, error) {
	tableName := "dialogs"
	query := fmt.Sprintf("insert into `%s` (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", tableName, dialogsRowsExpectAutoSet)

	r, err := m.db.Exec(ctx, query, data.UserId, data.PeerType, data.PeerId, data.PeerDialogId, data.Pinned, data.TopMessage, data.PinnedMsgId, data.ReadInboxMaxId, data.ReadOutboxMaxId, data.UnreadCount, data.UnreadMentionsCount, data.UnreadReactionsCount, data.UnreadMark, data.DraftType, data.DraftMessageData, data.FolderId, data.FolderPinned, data.HasScheduled, data.TtlPeriod, data.ThemeEmoticon, data.WallpaperId, data.WallpaperOverridden, data.Date2, data.Deleted)
	if err != nil {
		return nil, fmt.Errorf("dialogs.Insert2 exec: %w", err)
	}

	return r, nil
}

func (m *defaultDialogsModel) Delete2(ctx context.Context, id int64) error {
	tableName := "dialogs"
	query := fmt.Sprintf("delete from `%s` where `id` = ?", tableName)

	_, err := m.db.Exec(ctx, query, id)
	if err != nil {
		return fmt.Errorf("dialogs.Delete2 exec: %w", err)
	}

	return nil
}

func (m *defaultDialogsModel) FindOne(ctx context.Context, id int64) (*Dialogs, error) {
	tableName := "dialogs"
	query := fmt.Sprintf("select %s from %s where id = ? limit 1", dialogsRows, tableName)
	var resp Dialogs

	err := m.db.QueryRowPartial(ctx, &resp, query, id)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return nil, &NotFoundError{
				Resource: "dialogs",
				Key:      fmt.Sprintf("id=%v", id),
				Cause:    err,
			}
		}
		return nil, fmt.Errorf("dialogs.FindOne: %w", err)
	}

	return &resp, nil
}

func (m *defaultDialogsModel) FindListByIdList(ctx context.Context, id ...int64) ([]Dialogs, error) {
	if len(id) == 0 {
		return []Dialogs{}, nil
	}
	tableName := "dialogs"

	query := fmt.Sprintf("select %s from %s where id in (%s)", dialogsRows, tableName, sqlx.InInt64List(id))

	var resp []Dialogs
	err := m.db.QueryRowsPartial(ctx, &resp, query)
	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return []Dialogs{}, nil
		}
		return nil, fmt.Errorf("dialogs.FindListByIdList: %w", err)
	}

	return resp, nil
}

func (m *defaultDialogsModel) Update2(ctx context.Context, data *Dialogs) error {
	tableName := "dialogs"
	query := fmt.Sprintf("update `%s` set %s where `id` = ?", tableName, dialogsRowsWithPlaceHolder)

	_, err := m.db.Exec(ctx, query, data.UserId, data.PeerType, data.PeerId, data.PeerDialogId, data.Pinned, data.TopMessage, data.PinnedMsgId, data.ReadInboxMaxId, data.ReadOutboxMaxId, data.UnreadCount, data.UnreadMentionsCount, data.UnreadReactionsCount, data.UnreadMark, data.DraftType, data.DraftMessageData, data.FolderId, data.FolderPinned, data.HasScheduled, data.TtlPeriod, data.ThemeEmoticon, data.WallpaperId, data.WallpaperOverridden, data.Date2, data.Deleted, data.Id)
	if err != nil {
		return fmt.Errorf("dialogs.Update2 exec: %w", err)
	}

	return nil
}

func (m *defaultDialogsModel) FindOneByUserIdPeerTypePeerId(ctx context.Context, userId int64, peerType int32, peerId int64) (*Dialogs, error) {
	tableName := "dialogs"
	query := fmt.Sprintf("select %s from %s where user_id = ? AND peer_type = ? AND peer_id = ? limit 1", dialogsRows, tableName)
	var resp Dialogs

	err := m.db.QueryRowPartial(ctx, &resp, query, userId, peerType, peerId)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return nil, &NotFoundError{
				Resource: "dialogs",
				Key:      fmt.Sprintf("user_id=%v,peer_type=%v,peer_id=%v", userId, peerType, peerId),
				Cause:    err,
			}
		}
		return nil, fmt.Errorf("dialogs.FindOneByUserIdPeerTypePeerId: %w", err)
	}

	return &resp, nil
}

func (m *defaultDialogsModel) FindOneByUserIdPeerDialogId(ctx context.Context, userId int64, peerDialogId int64) (*Dialogs, error) {
	tableName := "dialogs"
	query := fmt.Sprintf("select %s from %s where user_id = ? AND peer_dialog_id = ? limit 1", dialogsRows, tableName)
	var resp Dialogs

	err := m.db.QueryRowPartial(ctx, &resp, query, userId, peerDialogId)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return nil, &NotFoundError{
				Resource: "dialogs",
				Key:      fmt.Sprintf("user_id=%v,peer_dialog_id=%v", userId, peerDialogId),
				Cause:    err,
			}
		}
		return nil, fmt.Errorf("dialogs.FindOneByUserIdPeerDialogId: %w", err)
	}

	return &resp, nil
}
