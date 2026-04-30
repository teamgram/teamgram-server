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
	messageClientRandomsFieldNames          = builder.RawFieldNames(&MessageClientRandoms{})
	messageClientRandomsRows                = strings.Join(messageClientRandomsFieldNames, ",")
	messageClientRandomsRowsExpectAutoSet   = strings.Join(stringx.Remove(messageClientRandomsFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	messageClientRandomsRowsWithPlaceHolder = strings.Join(stringx.Remove(messageClientRandomsFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"
)

type (
	messageClientRandomsModel interface {
		Insert2(ctx context.Context, data *MessageClientRandoms) (sql.Result, error)

		FindOneByCanonicalMessageId(ctx context.Context, canonicalMessageId int64) (*MessageClientRandoms, error)
		FindListByCanonicalMessageIdList(ctx context.Context, canonicalMessageId ...int64) ([]MessageClientRandoms, error)
	}

	defaultMessageClientRandomsModel struct {
		db *sqlx.DB
	}

	MessageClientRandoms struct {
		SenderUserId       int64  `db:"sender_user_id" json:"sender_user_id"`
		PeerType           int32  `db:"peer_type" json:"peer_type"`
		PeerId             int64  `db:"peer_id" json:"peer_id"`
		ClientRandomId     int64  `db:"client_random_id" json:"client_random_id"`
		CanonicalMessageId int64  `db:"canonical_message_id" json:"canonical_message_id"`
		SendStateId        int64  `db:"send_state_id" json:"send_state_id"`
		RequestPayloadHash []byte `db:"request_payload_hash" json:"request_payload_hash"`
	}
)

func newMessageClientRandomsModel(db *sqlx.DB) *defaultMessageClientRandomsModel {
	return &defaultMessageClientRandomsModel{
		db: db,
	}
}

func (m *defaultMessageClientRandomsModel) Insert2(ctx context.Context, data *MessageClientRandoms) (sql.Result, error) {
	tableName := "message_client_randoms"
	query := fmt.Sprintf("insert into `%s` (%s) values (?, ?, ?, ?, ?, ?, ?)", tableName, messageClientRandomsRowsExpectAutoSet)

	r, err := m.db.Exec(ctx, query, data.SenderUserId, data.PeerType, data.PeerId, data.ClientRandomId, data.CanonicalMessageId, data.SendStateId, data.RequestPayloadHash)
	if err != nil {
		return nil, fmt.Errorf("message_client_randoms.Insert2 exec: %w", err)
	}

	return r, nil
}

func (m *defaultMessageClientRandomsModel) FindOneByCanonicalMessageId(ctx context.Context, canonicalMessageId int64) (*MessageClientRandoms, error) {
	tableName := "message_client_randoms"
	query := fmt.Sprintf("select %s from %s where canonical_message_id = ? limit 1", messageClientRandomsRows, tableName)
	var resp MessageClientRandoms

	err := m.db.QueryRowPartial(ctx, &resp, query, canonicalMessageId)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return nil, &NotFoundError{
				Resource: "message_client_randoms",
				Key:      fmt.Sprintf("canonical_message_id=%v", canonicalMessageId),
				Cause:    err,
			}
		}
		return nil, fmt.Errorf("message_client_randoms.FindOneByCanonicalMessageId: %w", err)
	}

	return &resp, nil
}

func (m *defaultMessageClientRandomsModel) FindListByCanonicalMessageIdList(ctx context.Context, canonicalMessageId ...int64) ([]MessageClientRandoms, error) {
	if len(canonicalMessageId) == 0 {
		return []MessageClientRandoms{}, nil
	}
	tableName := "message_client_randoms"

	query := fmt.Sprintf("select %s from %s where canonical_message_id in (%s)", messageClientRandomsRows, tableName, sqlx.InInt64List(canonicalMessageId))

	var resp []MessageClientRandoms
	err := m.db.QueryRowsPartial(ctx, &resp, query)
	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return []MessageClientRandoms{}, nil
		}
		return nil, fmt.Errorf("message_client_randoms.FindListByCanonicalMessageIdList: %w", err)
	}

	return resp, nil
}
