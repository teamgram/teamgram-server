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
	savedDialogsFieldNames          = builder.RawFieldNames(&SavedDialogs{})
	savedDialogsRows                = strings.Join(savedDialogsFieldNames, ",")
	savedDialogsRowsExpectAutoSet   = strings.Join(stringx.Remove(savedDialogsFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	savedDialogsRowsWithPlaceHolder = strings.Join(stringx.Remove(savedDialogsFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"
)

type (
	savedDialogsModel interface {
		Insert2(ctx context.Context, data *SavedDialogs) (sql.Result, error)
	}

	defaultSavedDialogsModel struct {
		db *sqlx.DB
	}

	SavedDialogs struct {
		UserId                int64  `db:"user_id" json:"user_id"`
		PeerType              int32  `db:"peer_type" json:"peer_type"`
		PeerId                int64  `db:"peer_id" json:"peer_id"`
		TopPeerSeq            int64  `db:"top_peer_seq" json:"top_peer_seq"`
		TopCanonicalMessageId int64  `db:"top_canonical_message_id" json:"top_canonical_message_id"`
		TopMessageDate        int64  `db:"top_message_date" json:"top_message_date"`
		Pinned                bool   `db:"pinned" json:"pinned"`
		PinOrder              int64  `db:"pin_order" json:"pin_order"`
		Deleted               bool   `db:"deleted" json:"deleted"`
		SavedSchemaVersion    int32  `db:"saved_schema_version" json:"saved_schema_version"`
		SavedPayload          []byte `db:"saved_payload" json:"saved_payload"`
	}
)

func newSavedDialogsModel(db *sqlx.DB) *defaultSavedDialogsModel {
	return &defaultSavedDialogsModel{
		db: db,
	}
}

func (m *defaultSavedDialogsModel) Insert2(ctx context.Context, data *SavedDialogs) (sql.Result, error) {
	tableName := "saved_dialogs"
	query := fmt.Sprintf("insert into `%s` (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", tableName, savedDialogsRowsExpectAutoSet)

	r, err := m.db.Exec(ctx, query, data.UserId, data.PeerType, data.PeerId, data.TopPeerSeq, data.TopCanonicalMessageId, data.TopMessageDate, data.Pinned, data.PinOrder, data.Deleted, data.SavedSchemaVersion, data.SavedPayload)
	if err != nil {
		return nil, fmt.Errorf("saved_dialogs.Insert2 exec: %w", err)
	}

	return r, nil
}
