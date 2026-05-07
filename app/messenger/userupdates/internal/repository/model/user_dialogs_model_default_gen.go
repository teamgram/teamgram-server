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
	userDialogsFieldNames          = builder.RawFieldNames(&UserDialogs{})
	userDialogsRows                = strings.Join(userDialogsFieldNames, ",")
	userDialogsRowsExpectAutoSet   = strings.Join(stringx.Remove(userDialogsFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	userDialogsRowsWithPlaceHolder = strings.Join(stringx.Remove(userDialogsFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"
)

type (
	userDialogsModel interface {
		Insert2(ctx context.Context, data *UserDialogs) (sql.Result, error)
	}

	defaultUserDialogsModel struct {
		db *sqlx.DB
	}

	UserDialogs struct {
		UserId                   int64  `db:"user_id" json:"user_id"`
		PeerType                 int32  `db:"peer_type" json:"peer_type"`
		PeerId                   int64  `db:"peer_id" json:"peer_id"`
		TopPeerSeq               int64  `db:"top_peer_seq" json:"top_peer_seq"`
		TopCanonicalMessageId    int64  `db:"top_canonical_message_id" json:"top_canonical_message_id"`
		TopMessageDate           int64  `db:"top_message_date" json:"top_message_date"`
		TopMessageStatus         int32  `db:"top_message_status" json:"top_message_status"`
		UnreadCount              int32  `db:"unread_count" json:"unread_count"`
		UnreadMentionsCount      int32  `db:"unread_mentions_count" json:"unread_mentions_count"`
		UnreadReactionsCount     int32  `db:"unread_reactions_count" json:"unread_reactions_count"`
		UnreadMark               bool   `db:"unread_mark" json:"unread_mark"`
		PinnedPeerSeq            int64  `db:"pinned_peer_seq" json:"pinned_peer_seq"`
		PinnedCanonicalMessageId int64  `db:"pinned_canonical_message_id" json:"pinned_canonical_message_id"`
		HasScheduled             bool   `db:"has_scheduled" json:"has_scheduled"`
		AvailableMinPeerSeq      int64  `db:"available_min_peer_seq" json:"available_min_peer_seq"`
		Hidden                   bool   `db:"hidden" json:"hidden"`
		DeletedAt                int64  `db:"deleted_at" json:"deleted_at"`
		LastPts                  int64  `db:"last_pts" json:"last_pts"`
		LastPtsAt                int64  `db:"last_pts_at" json:"last_pts_at"`
		ReadInboxMaxPeerSeq      int64  `db:"read_inbox_max_peer_seq" json:"read_inbox_max_peer_seq"`
		ReadOutboxMaxPeerSeq     int64  `db:"read_outbox_max_peer_seq" json:"read_outbox_max_peer_seq"`
		DialogSchemaVersion      int32  `db:"dialog_schema_version" json:"dialog_schema_version"`
		DialogPayload            []byte `db:"dialog_payload" json:"dialog_payload"`
	}
)

func newUserDialogsModel(db *sqlx.DB) *defaultUserDialogsModel {
	return &defaultUserDialogsModel{
		db: db,
	}
}

func (m *defaultUserDialogsModel) Insert2(ctx context.Context, data *UserDialogs) (sql.Result, error) {
	tableName := "user_dialogs"
	query := fmt.Sprintf("insert into `%s` (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", tableName, userDialogsRowsExpectAutoSet)

	r, err := m.db.Exec(ctx, query, data.UserId, data.PeerType, data.PeerId, data.TopPeerSeq, data.TopCanonicalMessageId, data.TopMessageDate, data.TopMessageStatus, data.UnreadCount, data.UnreadMentionsCount, data.UnreadReactionsCount, data.UnreadMark, data.PinnedPeerSeq, data.PinnedCanonicalMessageId, data.HasScheduled, data.AvailableMinPeerSeq, data.Hidden, data.DeletedAt, data.LastPts, data.LastPtsAt, data.ReadInboxMaxPeerSeq, data.ReadOutboxMaxPeerSeq, data.DialogSchemaVersion, data.DialogPayload)
	if err != nil {
		return nil, fmt.Errorf("user_dialogs.Insert2 exec: %w", err)
	}

	return r, nil
}
