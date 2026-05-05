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
	userAuthSeqEventsFieldNames          = builder.RawFieldNames(&UserAuthSeqEvents{})
	userAuthSeqEventsRows                = strings.Join(userAuthSeqEventsFieldNames, ",")
	userAuthSeqEventsRowsExpectAutoSet   = strings.Join(stringx.Remove(userAuthSeqEventsFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	userAuthSeqEventsRowsWithPlaceHolder = strings.Join(stringx.Remove(userAuthSeqEventsFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"
)

type (
	userAuthSeqEventsModel interface {
		Insert2(ctx context.Context, data *UserAuthSeqEvents) (sql.Result, error)

		FindOneByUserIdOperationId(ctx context.Context, userId int64, operationId string) (*UserAuthSeqEvents, error)
	}

	defaultUserAuthSeqEventsModel struct {
		db *sqlx.DB
	}

	UserAuthSeqEvents struct {
		UserId              int64  `db:"user_id" json:"user_id"`
		Seq                 int64  `db:"seq" json:"seq"`
		Date                int32  `db:"date" json:"date"`
		OperationId         string `db:"operation_id" json:"operation_id"`
		SourcePermAuthKeyId int64  `db:"source_perm_auth_key_id" json:"source_perm_auth_key_id"`
		TargetAuthPolicy    string `db:"target_auth_policy" json:"target_auth_policy"`
		PublicUpdateType    string `db:"public_update_type" json:"public_update_type"`
		PeerType            int32  `db:"peer_type" json:"peer_type"`
		PeerId              int64  `db:"peer_id" json:"peer_id"`
		EventSchemaVersion  int32  `db:"event_schema_version" json:"event_schema_version"`
		EventCodec          int32  `db:"event_codec" json:"event_codec"`
		EventPayload        []byte `db:"event_payload" json:"event_payload"`
		EventPayloadHash    []byte `db:"event_payload_hash" json:"event_payload_hash"`
	}
)

func newUserAuthSeqEventsModel(db *sqlx.DB) *defaultUserAuthSeqEventsModel {
	return &defaultUserAuthSeqEventsModel{
		db: db,
	}
}

func (m *defaultUserAuthSeqEventsModel) Insert2(ctx context.Context, data *UserAuthSeqEvents) (sql.Result, error) {
	tableName := "user_auth_seq_events"
	query := fmt.Sprintf("insert into `%s` (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", tableName, userAuthSeqEventsRowsExpectAutoSet)

	r, err := m.db.Exec(ctx, query, data.UserId, data.Seq, data.Date, data.OperationId, data.SourcePermAuthKeyId, data.TargetAuthPolicy, data.PublicUpdateType, data.PeerType, data.PeerId, data.EventSchemaVersion, data.EventCodec, data.EventPayload, data.EventPayloadHash)
	if err != nil {
		return nil, fmt.Errorf("user_auth_seq_events.Insert2 exec: %w", err)
	}

	return r, nil
}

func (m *defaultUserAuthSeqEventsModel) FindOneByUserIdOperationId(ctx context.Context, userId int64, operationId string) (*UserAuthSeqEvents, error) {
	tableName := "user_auth_seq_events"
	query := fmt.Sprintf("select %s from %s where user_id = ? AND operation_id = ? limit 1", userAuthSeqEventsRows, tableName)
	var resp UserAuthSeqEvents

	err := m.db.QueryRowPartial(ctx, &resp, query, userId, operationId)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return nil, &NotFoundError{
				Resource: "user_auth_seq_events",
				Key:      fmt.Sprintf("user_id=%v,operation_id=%v", userId, operationId),
				Cause:    err,
			}
		}
		return nil, fmt.Errorf("user_auth_seq_events.FindOneByUserIdOperationId: %w", err)
	}

	return &resp, nil
}
