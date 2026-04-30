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
	userPtsEventsFieldNames          = builder.RawFieldNames(&UserPtsEvents{})
	userPtsEventsRows                = strings.Join(userPtsEventsFieldNames, ",")
	userPtsEventsRowsExpectAutoSet   = strings.Join(stringx.Remove(userPtsEventsFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	userPtsEventsRowsWithPlaceHolder = strings.Join(stringx.Remove(userPtsEventsFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"
)

type (
	userPtsEventsModel interface {
		Insert2(ctx context.Context, data *UserPtsEvents) (sql.Result, error)

		FindOneByUserIdOperationId(ctx context.Context, userId int64, operationId string) (*UserPtsEvents, error)
	}

	defaultUserPtsEventsModel struct {
		db *sqlx.DB
	}

	UserPtsEvents struct {
		UserId             int64  `db:"user_id" json:"user_id"`
		Pts                int64  `db:"pts" json:"pts"`
		PtsCount           int32  `db:"pts_count" json:"pts_count"`
		OperationId        string `db:"operation_id" json:"operation_id"`
		OpType             int32  `db:"op_type" json:"op_type"`
		EventType          int32  `db:"event_type" json:"event_type"`
		PeerType           int32  `db:"peer_type" json:"peer_type"`
		PeerId             int64  `db:"peer_id" json:"peer_id"`
		CanonicalMessageId int64  `db:"canonical_message_id" json:"canonical_message_id"`
		PeerSeq            int64  `db:"peer_seq" json:"peer_seq"`
		ActorUserId        int64  `db:"actor_user_id" json:"actor_user_id"`
		EventSchemaVersion int32  `db:"event_schema_version" json:"event_schema_version"`
		EventCodec         int32  `db:"event_codec" json:"event_codec"`
		EventPayload       []byte `db:"event_payload" json:"event_payload"`
		EventPayloadHash   []byte `db:"event_payload_hash" json:"event_payload_hash"`
	}
)

func newUserPtsEventsModel(db *sqlx.DB) *defaultUserPtsEventsModel {
	return &defaultUserPtsEventsModel{
		db: db,
	}
}

func (m *defaultUserPtsEventsModel) Insert2(ctx context.Context, data *UserPtsEvents) (sql.Result, error) {
	tableName := "user_pts_events"
	query := fmt.Sprintf("insert into `%s` (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", tableName, userPtsEventsRowsExpectAutoSet)

	r, err := m.db.Exec(ctx, query, data.UserId, data.Pts, data.PtsCount, data.OperationId, data.OpType, data.EventType, data.PeerType, data.PeerId, data.CanonicalMessageId, data.PeerSeq, data.ActorUserId, data.EventSchemaVersion, data.EventCodec, data.EventPayload, data.EventPayloadHash)
	if err != nil {
		return nil, fmt.Errorf("user_pts_events.Insert2 exec: %w", err)
	}

	return r, nil
}

func (m *defaultUserPtsEventsModel) FindOneByUserIdOperationId(ctx context.Context, userId int64, operationId string) (*UserPtsEvents, error) {
	tableName := "user_pts_events"
	query := fmt.Sprintf("select %s from %s where user_id = ? AND operation_id = ? limit 1", userPtsEventsRows, tableName)
	var resp UserPtsEvents

	err := m.db.QueryRowPartial(ctx, &resp, query, userId, operationId)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return nil, &NotFoundError{
				Resource: "user_pts_events",
				Key:      fmt.Sprintf("user_id=%v,operation_id=%v", userId, operationId),
				Cause:    err,
			}
		}
		return nil, fmt.Errorf("user_pts_events.FindOneByUserIdOperationId: %w", err)
	}

	return &resp, nil
}
