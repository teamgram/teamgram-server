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

type bizUserAuthSeqEventsModel interface {
	Insert(ctx context.Context, data *UserAuthSeqEvents) (lastInsertId, rowsAffected int64, err error)
	SelectByOperation(ctx context.Context, userId int64, operationId string) (*UserAuthSeqEvents, error)
	SelectAfterDate(ctx context.Context, userId int64, date int32, limit int32) ([]UserAuthSeqEvents, error)
	SelectAfterDateWithCB(ctx context.Context, userId int64, date int32, limit int32, cb func(sz, i int, v *UserAuthSeqEvents)) ([]UserAuthSeqEvents, error)
}

type UserAuthSeqEventsTxModel interface {
	Insert(data *UserAuthSeqEvents) (lastInsertId, rowsAffected int64, err error)
	SelectByOperation(userId int64, operationId string) (*UserAuthSeqEvents, error)
	SelectAfterDate(userId int64, date int32, limit int32) ([]UserAuthSeqEvents, error)
}

type defaultUserAuthSeqEventsTxModel struct {
	tx *sqlx.Tx
}

func NewUserAuthSeqEventsTxModel(tx *sqlx.Tx) UserAuthSeqEventsTxModel {
	return &defaultUserAuthSeqEventsTxModel{tx: tx}
}

// Insert
// insert into user_auth_seq_events(user_id, seq, `date`, operation_id, source_perm_auth_key_id, target_auth_policy, public_update_type, peer_type, peer_id, event_schema_version, event_codec, event_payload, event_payload_hash) values (:user_id, :seq, :date, :operation_id, :source_perm_auth_key_id, :target_auth_policy, :public_update_type, :peer_type, :peer_id, :event_schema_version, :event_codec, :event_payload, :event_payload_hash)
func (m *defaultUserAuthSeqEventsModel) Insert(ctx context.Context, data *UserAuthSeqEvents) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into user_auth_seq_events(user_id, seq, `date`, operation_id, source_perm_auth_key_id, target_auth_policy, public_update_type, peer_type, peer_id, event_schema_version, event_codec, event_payload, event_payload_hash) values (:user_id, :seq, :date, :operation_id, :source_perm_auth_key_id, :target_auth_policy, :public_update_type, :peer_type, :peer_id, :event_schema_version, :event_codec, :event_payload, :event_payload_hash)"
		r     sql.Result
	)

	r, err = m.db.NamedExec(ctx, query, data)
	if err != nil {
		err = fmt.Errorf("user_auth_seq_events.Insert named exec: %w", err)
		return
	}

	lastInsertId, err = r.LastInsertId()
	if err != nil {
		err = fmt.Errorf("user_auth_seq_events.Insert last insert id: %w", err)
		return
	}
	rowsAffected, err = r.RowsAffected()
	if err != nil {
		err = fmt.Errorf("user_auth_seq_events.Insert rows affected: %w", err)
	}

	return

}

// Insert
// insert into user_auth_seq_events(user_id, seq, `date`, operation_id, source_perm_auth_key_id, target_auth_policy, public_update_type, peer_type, peer_id, event_schema_version, event_codec, event_payload, event_payload_hash) values (:user_id, :seq, :date, :operation_id, :source_perm_auth_key_id, :target_auth_policy, :public_update_type, :peer_type, :peer_id, :event_schema_version, :event_codec, :event_payload, :event_payload_hash)
func (m *defaultUserAuthSeqEventsTxModel) Insert(data *UserAuthSeqEvents) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into user_auth_seq_events(user_id, seq, `date`, operation_id, source_perm_auth_key_id, target_auth_policy, public_update_type, peer_type, peer_id, event_schema_version, event_codec, event_payload, event_payload_hash) values (:user_id, :seq, :date, :operation_id, :source_perm_auth_key_id, :target_auth_policy, :public_update_type, :peer_type, :peer_id, :event_schema_version, :event_codec, :event_payload, :event_payload_hash)"
		r     sql.Result
	)

	r, err = m.tx.NamedExec(query, data)
	if err != nil {
		err = fmt.Errorf("user_auth_seq_events.Insert named exec: %w", err)
		return
	}

	lastInsertId, err = r.LastInsertId()
	if err != nil {
		err = fmt.Errorf("user_auth_seq_events.Insert last insert id: %w", err)
		return
	}
	rowsAffected, err = r.RowsAffected()
	if err != nil {
		err = fmt.Errorf("user_auth_seq_events.Insert rows affected: %w", err)
	}

	return
}

// SelectByOperation
// select user_id, seq, `date`, operation_id, source_perm_auth_key_id, target_auth_policy, public_update_type, peer_type, peer_id, event_schema_version, event_codec, event_payload, event_payload_hash from user_auth_seq_events where user_id = :user_id and operation_id = :operation_id limit 1
func (m *defaultUserAuthSeqEventsModel) SelectByOperation(ctx context.Context, userId int64, operationId string) (rValue *UserAuthSeqEvents, err error) {

	var (
		query = "select user_id, seq, `date`, operation_id, source_perm_auth_key_id, target_auth_policy, public_update_type, peer_type, peer_id, event_schema_version, event_codec, event_payload, event_payload_hash from user_auth_seq_events where user_id = ? and operation_id = ? limit 1"
		do    = &UserAuthSeqEvents{}
	)
	err = m.db.QueryRowPartial(ctx, do, query, userId, operationId)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return nil, &NotFoundError{
				Resource: "user_auth_seq_events",
				Key:      fmt.Sprintf("user_id=%v,operation_id=%v", userId, operationId),
				Cause:    err,
			}
		}
		err = fmt.Errorf("user_auth_seq_events.SelectByOperation: %w", err)
		return
	} else {
		rValue = do
	}

	return
}

// SelectByOperation
// select user_id, seq, `date`, operation_id, source_perm_auth_key_id, target_auth_policy, public_update_type, peer_type, peer_id, event_schema_version, event_codec, event_payload, event_payload_hash from user_auth_seq_events where user_id = :user_id and operation_id = :operation_id limit 1
func (m *defaultUserAuthSeqEventsTxModel) SelectByOperation(userId int64, operationId string) (rValue *UserAuthSeqEvents, err error) {
	var (
		query = "select user_id, seq, `date`, operation_id, source_perm_auth_key_id, target_auth_policy, public_update_type, peer_type, peer_id, event_schema_version, event_codec, event_payload, event_payload_hash from user_auth_seq_events where user_id = ? and operation_id = ? limit 1"
		do    = &UserAuthSeqEvents{}
	)
	err = m.tx.QueryRowPartial(do, query, userId, operationId)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return nil, &NotFoundError{
				Resource: "user_auth_seq_events",
				Key:      fmt.Sprintf("user_id=%v,operation_id=%v", userId, operationId),
				Cause:    err,
			}
		}
		err = fmt.Errorf("user_auth_seq_events.SelectByOperation: %w", err)
		return
	}
	rValue = do

	return
}

// SelectAfterDate
// select user_id, seq, `date`, operation_id, source_perm_auth_key_id, target_auth_policy, public_update_type, peer_type, peer_id, event_schema_version, event_codec, event_payload, event_payload_hash from user_auth_seq_events where user_id = :user_id and `date` > :date order by seq asc limit :limit
func (m *defaultUserAuthSeqEventsModel) SelectAfterDate(ctx context.Context, userId int64, date int32, limit int32) (rList []UserAuthSeqEvents, err error) {
	var (
		query  = "select user_id, seq, `date`, operation_id, source_perm_auth_key_id, target_auth_policy, public_update_type, peer_type, peer_id, event_schema_version, event_codec, event_payload, event_payload_hash from user_auth_seq_events where user_id = ? and `date` > ? order by seq asc limit ?"
		values []UserAuthSeqEvents
	)
	err = m.db.QueryRowsPartial(ctx, &values, query, userId, date, limit)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			rList = []UserAuthSeqEvents{}
			err = nil
			return
		}
		err = fmt.Errorf("user_auth_seq_events.SelectAfterDate: %w", err)
		return
	}

	rList = values

	return
}

// SelectAfterDate
// select user_id, seq, `date`, operation_id, source_perm_auth_key_id, target_auth_policy, public_update_type, peer_type, peer_id, event_schema_version, event_codec, event_payload, event_payload_hash from user_auth_seq_events where user_id = :user_id and `date` > :date order by seq asc limit :limit
func (m *defaultUserAuthSeqEventsTxModel) SelectAfterDate(userId int64, date int32, limit int32) (rList []UserAuthSeqEvents, err error) {
	var (
		query  = "select user_id, seq, `date`, operation_id, source_perm_auth_key_id, target_auth_policy, public_update_type, peer_type, peer_id, event_schema_version, event_codec, event_payload, event_payload_hash from user_auth_seq_events where user_id = ? and `date` > ? order by seq asc limit ?"
		values []UserAuthSeqEvents
	)
	err = m.tx.QueryRowsPartial(&values, query, userId, date, limit)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			rList = []UserAuthSeqEvents{}
			err = nil
			return
		}
		err = fmt.Errorf("user_auth_seq_events.SelectAfterDate: %w", err)
		return
	}

	rList = values

	return
}

// SelectAfterDateWithCB
// select user_id, seq, `date`, operation_id, source_perm_auth_key_id, target_auth_policy, public_update_type, peer_type, peer_id, event_schema_version, event_codec, event_payload, event_payload_hash from user_auth_seq_events where user_id = :user_id and `date` > :date order by seq asc limit :limit
func (m *defaultUserAuthSeqEventsModel) SelectAfterDateWithCB(ctx context.Context, userId int64, date int32, limit int32, cb func(sz, i int, v *UserAuthSeqEvents)) (rList []UserAuthSeqEvents, err error) {
	var (
		query  = "select user_id, seq, `date`, operation_id, source_perm_auth_key_id, target_auth_policy, public_update_type, peer_type, peer_id, event_schema_version, event_codec, event_payload, event_payload_hash from user_auth_seq_events where user_id = ? and `date` > ? order by seq asc limit ?"
		values []UserAuthSeqEvents
	)
	err = m.db.QueryRowsPartial(ctx, &values, query, userId, date, limit)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			rList = []UserAuthSeqEvents{}
			err = nil
			return
		}
		err = fmt.Errorf("user_auth_seq_events.SelectAfterDateWithCB: %w", err)
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
