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
	messagePeerSequencesFieldNames          = builder.RawFieldNames(&MessagePeerSequences{})
	messagePeerSequencesRows                = strings.Join(messagePeerSequencesFieldNames, ",")
	messagePeerSequencesRowsExpectAutoSet   = strings.Join(stringx.Remove(messagePeerSequencesFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	messagePeerSequencesRowsWithPlaceHolder = strings.Join(stringx.Remove(messagePeerSequencesFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"
)

type (
	messagePeerSequencesModel interface {
		Insert2(ctx context.Context, data *MessagePeerSequences) (sql.Result, error)
	}

	defaultMessagePeerSequencesModel struct {
		db *sqlx.DB
	}

	MessagePeerSequences struct {
		PeerType    int32 `db:"peer_type" json:"peer_type"`
		PeerId      int64 `db:"peer_id" json:"peer_id"`
		NextPeerSeq int64 `db:"next_peer_seq" json:"next_peer_seq"`
	}
)

func newMessagePeerSequencesModel(db *sqlx.DB) *defaultMessagePeerSequencesModel {
	return &defaultMessagePeerSequencesModel{
		db: db,
	}
}

func (m *defaultMessagePeerSequencesModel) Insert2(ctx context.Context, data *MessagePeerSequences) (sql.Result, error) {
	tableName := "message_peer_sequences"
	query := fmt.Sprintf("insert into `%s` (%s) values (?, ?, ?)", tableName, messagePeerSequencesRowsExpectAutoSet)

	r, err := m.db.Exec(ctx, query, data.PeerType, data.PeerId, data.NextPeerSeq)
	if err != nil {
		return nil, fmt.Errorf("message_peer_sequences.Insert2 exec: %w", err)
	}

	return r, nil
}
