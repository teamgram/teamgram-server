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
	dialogDraftsFieldNames          = builder.RawFieldNames(&DialogDrafts{})
	dialogDraftsRows                = strings.Join(dialogDraftsFieldNames, ",")
	dialogDraftsRowsExpectAutoSet   = strings.Join(stringx.Remove(dialogDraftsFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	dialogDraftsRowsWithPlaceHolder = strings.Join(stringx.Remove(dialogDraftsFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"
)

type (
	dialogDraftsModel interface {
		Insert2(ctx context.Context, data *DialogDrafts) (sql.Result, error)

		FindOneByUserIdPeerTypePeerId(ctx context.Context, userId int64, peerType int32, peerId int64) (*DialogDrafts, error)
	}

	defaultDialogDraftsModel struct {
		db *sqlx.DB
	}

	DialogDrafts struct {
		UserId                    int64  `db:"user_id" json:"user_id"`
		PeerType                  int32  `db:"peer_type" json:"peer_type"`
		PeerId                    int64  `db:"peer_id" json:"peer_id"`
		PeerDialogId              int64  `db:"peer_dialog_id" json:"peer_dialog_id"`
		DraftKind                 int32  `db:"draft_kind" json:"draft_kind"`
		Message                   string `db:"message" json:"message"`
		EntitiesPayload           []byte `db:"entities_payload" json:"entities_payload"`
		ReplyToPeerSeq            int64  `db:"reply_to_peer_seq" json:"reply_to_peer_seq"`
		DraftPayloadSchemaVersion int32  `db:"draft_payload_schema_version" json:"draft_payload_schema_version"`
		DraftPayload              []byte `db:"draft_payload" json:"draft_payload"`
		Date                      string `db:"date" json:"date"`
	}
)

func newDialogDraftsModel(db *sqlx.DB) *defaultDialogDraftsModel {
	return &defaultDialogDraftsModel{
		db: db,
	}
}

func (m *defaultDialogDraftsModel) Insert2(ctx context.Context, data *DialogDrafts) (sql.Result, error) {
	tableName := "dialog_drafts"
	query := fmt.Sprintf("insert into `%s` (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", tableName, dialogDraftsRowsExpectAutoSet)

	r, err := m.db.Exec(ctx, query, data.UserId, data.PeerType, data.PeerId, data.PeerDialogId, data.DraftKind, data.Message, data.EntitiesPayload, data.ReplyToPeerSeq, data.DraftPayloadSchemaVersion, data.DraftPayload, data.Date)
	if err != nil {
		return nil, fmt.Errorf("dialog_drafts.Insert2 exec: %w", err)
	}

	return r, nil
}

func (m *defaultDialogDraftsModel) FindOneByUserIdPeerTypePeerId(ctx context.Context, userId int64, peerType int32, peerId int64) (*DialogDrafts, error) {
	tableName := "dialog_drafts"
	query := fmt.Sprintf("select %s from %s where user_id = ? AND peer_type = ? AND peer_id = ? limit 1", dialogDraftsRows, tableName)
	var resp DialogDrafts

	err := m.db.QueryRowPartial(ctx, &resp, query, userId, peerType, peerId)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return nil, &NotFoundError{
				Resource: "dialog_drafts",
				Key:      fmt.Sprintf("user_id=%v,peer_type=%v,peer_id=%v", userId, peerType, peerId),
				Cause:    err,
			}
		}
		return nil, fmt.Errorf("dialog_drafts.FindOneByUserIdPeerTypePeerId: %w", err)
	}

	return &resp, nil
}
