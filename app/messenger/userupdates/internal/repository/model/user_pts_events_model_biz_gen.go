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
)

var _ *sql.Result
var _ = fmt.Sprintf
var _ = strings.Join
var _ = errors.Is
var _ *sqlx.DB
var _ *sqlx.Tx

type bizUserPtsEventsModel interface {
	Insert(ctx context.Context, data *UserPtsEvents) (lastInsertId, rowsAffected int64, err error)
	SelectByOperation(ctx context.Context, userId int64, operationId string) (*UserPtsEvents, error)
	SelectLatestPts(ctx context.Context, userId int64) (*UserPtsEvents, error)
	SelectByGtPts(ctx context.Context, userId int64, pts int64, limit int32) ([]UserPtsEvents, error)
	SelectByGtPtsWithCB(ctx context.Context, userId int64, pts int64, limit int32, cb func(sz, i int, v *UserPtsEvents)) ([]UserPtsEvents, error)
}

type UserPtsEventsTxModel interface {
	Insert(data *UserPtsEvents) (lastInsertId, rowsAffected int64, err error)
	SelectByOperation(userId int64, operationId string) (*UserPtsEvents, error)
	SelectLatestPts(userId int64) (*UserPtsEvents, error)
	SelectByGtPts(userId int64, pts int64, limit int32) ([]UserPtsEvents, error)
}

type defaultUserPtsEventsTxModel struct {
	tx *sqlx.Tx
}

func NewUserPtsEventsTxModel(tx *sqlx.Tx) UserPtsEventsTxModel {
	return &defaultUserPtsEventsTxModel{tx: tx}
}

// Insert
// insert into user_pts_events(user_id, pts, pts_count, operation_id, op_type, event_type, peer_type, peer_id, canonical_message_id, peer_seq, actor_user_id, event_schema_version, event_codec, event_payload, event_payload_hash) values (:user_id, :pts, :pts_count, :operation_id, :op_type, :event_type, :peer_type, :peer_id, :canonical_message_id, :peer_seq, :actor_user_id, :event_schema_version, :event_codec, :event_payload, :event_payload_hash)
func (m *defaultUserPtsEventsModel) Insert(ctx context.Context, data *UserPtsEvents) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into user_pts_events(user_id, pts, pts_count, operation_id, op_type, event_type, peer_type, peer_id, canonical_message_id, peer_seq, actor_user_id, event_schema_version, event_codec, event_payload, event_payload_hash) values (:user_id, :pts, :pts_count, :operation_id, :op_type, :event_type, :peer_type, :peer_id, :canonical_message_id, :peer_seq, :actor_user_id, :event_schema_version, :event_codec, :event_payload, :event_payload_hash)"
		r     sql.Result
	)

	r, err = m.db.NamedExec(ctx, query, data)
	if err != nil {
		err = fmt.Errorf("user_pts_events.Insert named exec: %w", err)
		return
	}

	lastInsertId, err = r.LastInsertId()
	if err != nil {
		err = fmt.Errorf("user_pts_events.Insert last insert id: %w", err)
		return
	}
	rowsAffected, err = r.RowsAffected()
	if err != nil {
		err = fmt.Errorf("user_pts_events.Insert rows affected: %w", err)
	}

	return

}

// Insert
// insert into user_pts_events(user_id, pts, pts_count, operation_id, op_type, event_type, peer_type, peer_id, canonical_message_id, peer_seq, actor_user_id, event_schema_version, event_codec, event_payload, event_payload_hash) values (:user_id, :pts, :pts_count, :operation_id, :op_type, :event_type, :peer_type, :peer_id, :canonical_message_id, :peer_seq, :actor_user_id, :event_schema_version, :event_codec, :event_payload, :event_payload_hash)
func (m *defaultUserPtsEventsTxModel) Insert(data *UserPtsEvents) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into user_pts_events(user_id, pts, pts_count, operation_id, op_type, event_type, peer_type, peer_id, canonical_message_id, peer_seq, actor_user_id, event_schema_version, event_codec, event_payload, event_payload_hash) values (:user_id, :pts, :pts_count, :operation_id, :op_type, :event_type, :peer_type, :peer_id, :canonical_message_id, :peer_seq, :actor_user_id, :event_schema_version, :event_codec, :event_payload, :event_payload_hash)"
		r     sql.Result
	)

	r, err = m.tx.NamedExec(query, data)
	if err != nil {
		err = fmt.Errorf("user_pts_events.Insert named exec: %w", err)
		return
	}

	lastInsertId, err = r.LastInsertId()
	if err != nil {
		err = fmt.Errorf("user_pts_events.Insert last insert id: %w", err)
		return
	}
	rowsAffected, err = r.RowsAffected()
	if err != nil {
		err = fmt.Errorf("user_pts_events.Insert rows affected: %w", err)
	}

	return
}

// SelectByOperation
// select user_id, pts, pts_count, operation_id, op_type, event_type, peer_type, peer_id, canonical_message_id, peer_seq, actor_user_id, event_schema_version, event_codec, event_payload, event_payload_hash from user_pts_events where user_id = :user_id and operation_id = :operation_id limit 1
func (m *defaultUserPtsEventsModel) SelectByOperation(ctx context.Context, userId int64, operationId string) (rValue *UserPtsEvents, err error) {

	var (
		query = "select user_id, pts, pts_count, operation_id, op_type, event_type, peer_type, peer_id, canonical_message_id, peer_seq, actor_user_id, event_schema_version, event_codec, event_payload, event_payload_hash from user_pts_events where user_id = ? and operation_id = ? limit 1"
		do    = &UserPtsEvents{}
	)
	err = m.db.QueryRowPartial(ctx, do, query, userId, operationId)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return nil, &NotFoundError{
				Resource: "user_pts_events",
				Key:      fmt.Sprintf("user_id=%v,operation_id=%v", userId, operationId),
				Cause:    err,
			}
		}
		err = fmt.Errorf("user_pts_events.SelectByOperation: %w", err)
		return
	} else {
		rValue = do
	}

	return
}

// SelectByOperation
// select user_id, pts, pts_count, operation_id, op_type, event_type, peer_type, peer_id, canonical_message_id, peer_seq, actor_user_id, event_schema_version, event_codec, event_payload, event_payload_hash from user_pts_events where user_id = :user_id and operation_id = :operation_id limit 1
func (m *defaultUserPtsEventsTxModel) SelectByOperation(userId int64, operationId string) (rValue *UserPtsEvents, err error) {
	var (
		query = "select user_id, pts, pts_count, operation_id, op_type, event_type, peer_type, peer_id, canonical_message_id, peer_seq, actor_user_id, event_schema_version, event_codec, event_payload, event_payload_hash from user_pts_events where user_id = ? and operation_id = ? limit 1"
		do    = &UserPtsEvents{}
	)
	err = m.tx.QueryRowPartial(do, query, userId, operationId)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return nil, &NotFoundError{
				Resource: "user_pts_events",
				Key:      fmt.Sprintf("user_id=%v,operation_id=%v", userId, operationId),
				Cause:    err,
			}
		}
		err = fmt.Errorf("user_pts_events.SelectByOperation: %w", err)
		return
	}
	rValue = do

	return
}

// SelectLatestPts
// select user_id, pts, pts_count, operation_id, op_type, event_type, peer_type, peer_id, canonical_message_id, peer_seq, actor_user_id, event_schema_version, event_codec, event_payload, event_payload_hash from user_pts_events where user_id = :user_id order by pts desc limit 1
func (m *defaultUserPtsEventsModel) SelectLatestPts(ctx context.Context, userId int64) (rValue *UserPtsEvents, err error) {

	var (
		query = "select user_id, pts, pts_count, operation_id, op_type, event_type, peer_type, peer_id, canonical_message_id, peer_seq, actor_user_id, event_schema_version, event_codec, event_payload, event_payload_hash from user_pts_events where user_id = ? order by pts desc limit 1"
		do    = &UserPtsEvents{}
	)
	err = m.db.QueryRowPartial(ctx, do, query, userId)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return nil, &NotFoundError{
				Resource: "user_pts_events",
				Key:      fmt.Sprintf("user_id=%v", userId),
				Cause:    err,
			}
		}
		err = fmt.Errorf("user_pts_events.SelectLatestPts: %w", err)
		return
	} else {
		rValue = do
	}

	return
}

// SelectLatestPts
// select user_id, pts, pts_count, operation_id, op_type, event_type, peer_type, peer_id, canonical_message_id, peer_seq, actor_user_id, event_schema_version, event_codec, event_payload, event_payload_hash from user_pts_events where user_id = :user_id order by pts desc limit 1
func (m *defaultUserPtsEventsTxModel) SelectLatestPts(userId int64) (rValue *UserPtsEvents, err error) {
	var (
		query = "select user_id, pts, pts_count, operation_id, op_type, event_type, peer_type, peer_id, canonical_message_id, peer_seq, actor_user_id, event_schema_version, event_codec, event_payload, event_payload_hash from user_pts_events where user_id = ? order by pts desc limit 1"
		do    = &UserPtsEvents{}
	)
	err = m.tx.QueryRowPartial(do, query, userId)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return nil, &NotFoundError{
				Resource: "user_pts_events",
				Key:      fmt.Sprintf("user_id=%v", userId),
				Cause:    err,
			}
		}
		err = fmt.Errorf("user_pts_events.SelectLatestPts: %w", err)
		return
	}
	rValue = do

	return
}

// SelectByGtPts
// select user_id, pts, pts_count, operation_id, op_type, event_type, peer_type, peer_id, canonical_message_id, peer_seq, actor_user_id, event_schema_version, event_codec, event_payload, event_payload_hash from user_pts_events where user_id = :user_id and pts > :pts order by pts asc limit :limit
func (m *defaultUserPtsEventsModel) SelectByGtPts(ctx context.Context, userId int64, pts int64, limit int32) (rList []UserPtsEvents, err error) {
	var (
		query  = "select user_id, pts, pts_count, operation_id, op_type, event_type, peer_type, peer_id, canonical_message_id, peer_seq, actor_user_id, event_schema_version, event_codec, event_payload, event_payload_hash from user_pts_events where user_id = ? and pts > ? order by pts asc limit ?"
		values []UserPtsEvents
	)
	err = m.db.QueryRowsPartial(ctx, &values, query, userId, pts, limit)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			rList = []UserPtsEvents{}
			err = nil
			return
		}
		err = fmt.Errorf("user_pts_events.SelectByGtPts: %w", err)
		return
	}

	rList = values

	return
}

// SelectByGtPts
// select user_id, pts, pts_count, operation_id, op_type, event_type, peer_type, peer_id, canonical_message_id, peer_seq, actor_user_id, event_schema_version, event_codec, event_payload, event_payload_hash from user_pts_events where user_id = :user_id and pts > :pts order by pts asc limit :limit
func (m *defaultUserPtsEventsTxModel) SelectByGtPts(userId int64, pts int64, limit int32) (rList []UserPtsEvents, err error) {
	var (
		query  = "select user_id, pts, pts_count, operation_id, op_type, event_type, peer_type, peer_id, canonical_message_id, peer_seq, actor_user_id, event_schema_version, event_codec, event_payload, event_payload_hash from user_pts_events where user_id = ? and pts > ? order by pts asc limit ?"
		values []UserPtsEvents
	)
	err = m.tx.QueryRowsPartial(&values, query, userId, pts, limit)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			rList = []UserPtsEvents{}
			err = nil
			return
		}
		err = fmt.Errorf("user_pts_events.SelectByGtPts: %w", err)
		return
	}

	rList = values

	return
}

// SelectByGtPtsWithCB
// select user_id, pts, pts_count, operation_id, op_type, event_type, peer_type, peer_id, canonical_message_id, peer_seq, actor_user_id, event_schema_version, event_codec, event_payload, event_payload_hash from user_pts_events where user_id = :user_id and pts > :pts order by pts asc limit :limit
func (m *defaultUserPtsEventsModel) SelectByGtPtsWithCB(ctx context.Context, userId int64, pts int64, limit int32, cb func(sz, i int, v *UserPtsEvents)) (rList []UserPtsEvents, err error) {
	var (
		query  = "select user_id, pts, pts_count, operation_id, op_type, event_type, peer_type, peer_id, canonical_message_id, peer_seq, actor_user_id, event_schema_version, event_codec, event_payload, event_payload_hash from user_pts_events where user_id = ? and pts > ? order by pts asc limit ?"
		values []UserPtsEvents
	)
	err = m.db.QueryRowsPartial(ctx, &values, query, userId, pts, limit)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			rList = []UserPtsEvents{}
			err = nil
			return
		}
		err = fmt.Errorf("user_pts_events.SelectByGtPtsWithCB: %w", err)
		return
	}

	rList = values

	if cb != nil {
		sz := len(rList)
		for i := 0; i < sz; i++ {
			cb(sz, i, &rList[i])
		}
	}

	return
}
