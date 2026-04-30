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
		UserId                int64  `db:"user_id" json:"user_id"`
		PeerType              int32  `db:"peer_type" json:"peer_type"`
		PeerId                int64  `db:"peer_id" json:"peer_id"`
		TopPeerSeq            int64  `db:"top_peer_seq" json:"top_peer_seq"`
		TopCanonicalMessageId int64  `db:"top_canonical_message_id" json:"top_canonical_message_id"`
		TopMessageDate        string `db:"top_message_date" json:"top_message_date"`
		UnreadCount           int32  `db:"unread_count" json:"unread_count"`
		UnreadMentionsCount   int32  `db:"unread_mentions_count" json:"unread_mentions_count"`
		ReadInboxMaxPeerSeq   int64  `db:"read_inbox_max_peer_seq" json:"read_inbox_max_peer_seq"`
		ReadOutboxMaxPeerSeq  int64  `db:"read_outbox_max_peer_seq" json:"read_outbox_max_peer_seq"`
		Pinned                bool   `db:"pinned" json:"pinned"`
		FolderId              int32  `db:"folder_id" json:"folder_id"`
		DialogSchemaVersion   int32  `db:"dialog_schema_version" json:"dialog_schema_version"`
		DialogPayload         []byte `db:"dialog_payload" json:"dialog_payload"`
	}
)

func newUserDialogsModel(db *sqlx.DB) *defaultUserDialogsModel {
	return &defaultUserDialogsModel{
		db: db,
	}
}

func (m *defaultUserDialogsModel) Insert2(ctx context.Context, data *UserDialogs) (sql.Result, error) {
	tableName := "user_dialogs"
	query := fmt.Sprintf("insert into `%s` (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", tableName, userDialogsRowsExpectAutoSet)

	r, err := m.db.Exec(ctx, query, data.UserId, data.PeerType, data.PeerId, data.TopPeerSeq, data.TopCanonicalMessageId, data.TopMessageDate, data.UnreadCount, data.UnreadMentionsCount, data.ReadInboxMaxPeerSeq, data.ReadOutboxMaxPeerSeq, data.Pinned, data.FolderId, data.DialogSchemaVersion, data.DialogPayload)
	if err != nil {
		return nil, fmt.Errorf("user_dialogs.Insert2 exec: %w", err)
	}

	return r, nil
}
