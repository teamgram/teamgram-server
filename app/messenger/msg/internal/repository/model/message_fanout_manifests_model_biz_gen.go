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

type bizMessageFanoutManifestsModel interface {
	Insert(ctx context.Context, data *MessageFanoutManifests) (lastInsertId, rowsAffected int64, err error)
	SelectByManifestId(ctx context.Context, manifestId int64) (*MessageFanoutManifests, error)
	SelectByCanonicalMessageId(ctx context.Context, canonicalMessageId int64) (*MessageFanoutManifests, error)
	MarkCompleted(ctx context.Context, status int32, completedAt sql.NullTime, manifestId int64) (rowsAffected int64, err error)
}

type MessageFanoutManifestsTxModel interface {
	Insert(data *MessageFanoutManifests) (lastInsertId, rowsAffected int64, err error)
	SelectByManifestId(manifestId int64) (*MessageFanoutManifests, error)
	SelectByCanonicalMessageId(canonicalMessageId int64) (*MessageFanoutManifests, error)
	MarkCompleted(status int32, completedAt sql.NullTime, manifestId int64) (rowsAffected int64, err error)
}

type defaultMessageFanoutManifestsTxModel struct {
	tx *sqlx.Tx
}

func NewMessageFanoutManifestsTxModel(tx *sqlx.Tx) MessageFanoutManifestsTxModel {
	return &defaultMessageFanoutManifestsTxModel{tx: tx}
}

// Insert
// insert into message_fanout_manifests(manifest_id, canonical_message_id, peer_type, peer_id, peer_seq, actor_user_id, affected_user_count, `status`, completed_at) values (:manifest_id, :canonical_message_id, :peer_type, :peer_id, :peer_seq, :actor_user_id, :affected_user_count, :status, :completed_at)
func (m *defaultMessageFanoutManifestsModel) Insert(ctx context.Context, data *MessageFanoutManifests) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into message_fanout_manifests(manifest_id, canonical_message_id, peer_type, peer_id, peer_seq, actor_user_id, affected_user_count, `status`, completed_at) values (:manifest_id, :canonical_message_id, :peer_type, :peer_id, :peer_seq, :actor_user_id, :affected_user_count, :status, :completed_at)"
		r     sql.Result
	)

	r, err = m.db.NamedExec(ctx, query, data)
	if err != nil {
		err = fmt.Errorf("message_fanout_manifests.Insert named exec: %w", err)
		return
	}

	lastInsertId, err = r.LastInsertId()
	if err != nil {
		err = fmt.Errorf("message_fanout_manifests.Insert last insert id: %w", err)
		return
	}
	rowsAffected, err = r.RowsAffected()
	if err != nil {
		err = fmt.Errorf("message_fanout_manifests.Insert rows affected: %w", err)
	}

	return

}

// Insert
// insert into message_fanout_manifests(manifest_id, canonical_message_id, peer_type, peer_id, peer_seq, actor_user_id, affected_user_count, `status`, completed_at) values (:manifest_id, :canonical_message_id, :peer_type, :peer_id, :peer_seq, :actor_user_id, :affected_user_count, :status, :completed_at)
func (m *defaultMessageFanoutManifestsTxModel) Insert(data *MessageFanoutManifests) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into message_fanout_manifests(manifest_id, canonical_message_id, peer_type, peer_id, peer_seq, actor_user_id, affected_user_count, `status`, completed_at) values (:manifest_id, :canonical_message_id, :peer_type, :peer_id, :peer_seq, :actor_user_id, :affected_user_count, :status, :completed_at)"
		r     sql.Result
	)

	r, err = m.tx.NamedExec(query, data)
	if err != nil {
		err = fmt.Errorf("message_fanout_manifests.Insert named exec: %w", err)
		return
	}

	lastInsertId, err = r.LastInsertId()
	if err != nil {
		err = fmt.Errorf("message_fanout_manifests.Insert last insert id: %w", err)
		return
	}
	rowsAffected, err = r.RowsAffected()
	if err != nil {
		err = fmt.Errorf("message_fanout_manifests.Insert rows affected: %w", err)
	}

	return
}

// SelectByManifestId
// select manifest_id, canonical_message_id, peer_type, peer_id, peer_seq, actor_user_id, affected_user_count, `status`, completed_at from message_fanout_manifests where manifest_id = :manifest_id limit 1
func (m *defaultMessageFanoutManifestsModel) SelectByManifestId(ctx context.Context, manifestId int64) (rValue *MessageFanoutManifests, err error) {

	var (
		query = "select manifest_id, canonical_message_id, peer_type, peer_id, peer_seq, actor_user_id, affected_user_count, `status`, completed_at from message_fanout_manifests where manifest_id = ? limit 1"
		do    = &MessageFanoutManifests{}
	)
	err = m.db.QueryRowPartial(ctx, do, query, manifestId)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return nil, &NotFoundError{
				Resource: "message_fanout_manifests",
				Key:      fmt.Sprintf("manifest_id=%v", manifestId),
				Cause:    err,
			}
		}
		err = fmt.Errorf("message_fanout_manifests.SelectByManifestId: %w", err)
		return
	} else {
		rValue = do
	}

	return
}

// SelectByManifestId
// select manifest_id, canonical_message_id, peer_type, peer_id, peer_seq, actor_user_id, affected_user_count, `status`, completed_at from message_fanout_manifests where manifest_id = :manifest_id limit 1
func (m *defaultMessageFanoutManifestsTxModel) SelectByManifestId(manifestId int64) (rValue *MessageFanoutManifests, err error) {
	var (
		query = "select manifest_id, canonical_message_id, peer_type, peer_id, peer_seq, actor_user_id, affected_user_count, `status`, completed_at from message_fanout_manifests where manifest_id = ? limit 1"
		do    = &MessageFanoutManifests{}
	)
	err = m.tx.QueryRowPartial(do, query, manifestId)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return nil, &NotFoundError{
				Resource: "message_fanout_manifests",
				Key:      fmt.Sprintf("manifest_id=%v", manifestId),
				Cause:    err,
			}
		}
		err = fmt.Errorf("message_fanout_manifests.SelectByManifestId: %w", err)
		return
	}
	rValue = do

	return
}

// SelectByCanonicalMessageId
// select manifest_id, canonical_message_id, peer_type, peer_id, peer_seq, actor_user_id, affected_user_count, `status`, completed_at from message_fanout_manifests where canonical_message_id = :canonical_message_id limit 1
func (m *defaultMessageFanoutManifestsModel) SelectByCanonicalMessageId(ctx context.Context, canonicalMessageId int64) (rValue *MessageFanoutManifests, err error) {

	var (
		query = "select manifest_id, canonical_message_id, peer_type, peer_id, peer_seq, actor_user_id, affected_user_count, `status`, completed_at from message_fanout_manifests where canonical_message_id = ? limit 1"
		do    = &MessageFanoutManifests{}
	)
	err = m.db.QueryRowPartial(ctx, do, query, canonicalMessageId)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return nil, &NotFoundError{
				Resource: "message_fanout_manifests",
				Key:      fmt.Sprintf("canonical_message_id=%v", canonicalMessageId),
				Cause:    err,
			}
		}
		err = fmt.Errorf("message_fanout_manifests.SelectByCanonicalMessageId: %w", err)
		return
	} else {
		rValue = do
	}

	return
}

// SelectByCanonicalMessageId
// select manifest_id, canonical_message_id, peer_type, peer_id, peer_seq, actor_user_id, affected_user_count, `status`, completed_at from message_fanout_manifests where canonical_message_id = :canonical_message_id limit 1
func (m *defaultMessageFanoutManifestsTxModel) SelectByCanonicalMessageId(canonicalMessageId int64) (rValue *MessageFanoutManifests, err error) {
	var (
		query = "select manifest_id, canonical_message_id, peer_type, peer_id, peer_seq, actor_user_id, affected_user_count, `status`, completed_at from message_fanout_manifests where canonical_message_id = ? limit 1"
		do    = &MessageFanoutManifests{}
	)
	err = m.tx.QueryRowPartial(do, query, canonicalMessageId)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return nil, &NotFoundError{
				Resource: "message_fanout_manifests",
				Key:      fmt.Sprintf("canonical_message_id=%v", canonicalMessageId),
				Cause:    err,
			}
		}
		err = fmt.Errorf("message_fanout_manifests.SelectByCanonicalMessageId: %w", err)
		return
	}
	rValue = do

	return
}

// MarkCompleted
// update message_fanout_manifests set `status` = :status, completed_at = :completed_at where manifest_id = :manifest_id
func (m *defaultMessageFanoutManifestsModel) MarkCompleted(ctx context.Context, status int32, completedAt sql.NullTime, manifestId int64) (rowsAffected int64, err error) {

	var (
		query   = "update message_fanout_manifests set `status` = ?, completed_at = ? where manifest_id = ?"
		rResult sql.Result
	)

	rResult, err = m.db.Exec(ctx, query, status, completedAt, manifestId)

	if err != nil {
		err = fmt.Errorf("message_fanout_manifests.MarkCompleted exec: %w", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		err = fmt.Errorf("message_fanout_manifests.MarkCompleted rows affected: %w", err)
		return
	}

	return
}

// MarkCompleted
// update message_fanout_manifests set `status` = :status, completed_at = :completed_at where manifest_id = :manifest_id
func (m *defaultMessageFanoutManifestsTxModel) MarkCompleted(status int32, completedAt sql.NullTime, manifestId int64) (rowsAffected int64, err error) {
	var (
		query   = "update message_fanout_manifests set `status` = ?, completed_at = ? where manifest_id = ?"
		rResult sql.Result
	)
	rResult, err = m.tx.Exec(query, status, completedAt, manifestId)

	if err != nil {
		err = fmt.Errorf("message_fanout_manifests.MarkCompleted exec: %w", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		err = fmt.Errorf("message_fanout_manifests.MarkCompleted rows affected: %w", err)
		return
	}

	return
}
