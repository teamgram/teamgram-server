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

type bizMessageClientRandomsModel interface {
	Insert(ctx context.Context, data *MessageClientRandoms) (lastInsertId, rowsAffected int64, err error)
	SelectByRandom(ctx context.Context, senderUserId int64, peerType int32, peerId int64, clientRandomId int64) (*MessageClientRandoms, error)
	SelectByCanonicalMessageId(ctx context.Context, canonicalMessageId int64) (*MessageClientRandoms, error)
}

type MessageClientRandomsTxModel interface {
	Insert(data *MessageClientRandoms) (lastInsertId, rowsAffected int64, err error)
	SelectByRandom(senderUserId int64, peerType int32, peerId int64, clientRandomId int64) (*MessageClientRandoms, error)
	SelectByCanonicalMessageId(canonicalMessageId int64) (*MessageClientRandoms, error)
}

type defaultMessageClientRandomsTxModel struct {
	tx *sqlx.Tx
}

func NewMessageClientRandomsTxModel(tx *sqlx.Tx) MessageClientRandomsTxModel {
	return &defaultMessageClientRandomsTxModel{tx: tx}
}

// Insert
// insert into message_client_randoms(sender_user_id, peer_type, peer_id, client_random_id, canonical_message_id, send_state_id, request_payload_hash) values (:sender_user_id, :peer_type, :peer_id, :client_random_id, :canonical_message_id, :send_state_id, :request_payload_hash)
func (m *defaultMessageClientRandomsModel) Insert(ctx context.Context, data *MessageClientRandoms) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into message_client_randoms(sender_user_id, peer_type, peer_id, client_random_id, canonical_message_id, send_state_id, request_payload_hash) values (:sender_user_id, :peer_type, :peer_id, :client_random_id, :canonical_message_id, :send_state_id, :request_payload_hash)"
		r     sql.Result
	)

	r, err = m.db.NamedExec(ctx, query, data)
	if err != nil {
		err = fmt.Errorf("message_client_randoms.Insert named exec: %w", err)
		return
	}

	lastInsertId, err = r.LastInsertId()
	if err != nil {
		err = fmt.Errorf("message_client_randoms.Insert last insert id: %w", err)
		return
	}
	rowsAffected, err = r.RowsAffected()
	if err != nil {
		err = fmt.Errorf("message_client_randoms.Insert rows affected: %w", err)
	}

	return

}

// Insert
// insert into message_client_randoms(sender_user_id, peer_type, peer_id, client_random_id, canonical_message_id, send_state_id, request_payload_hash) values (:sender_user_id, :peer_type, :peer_id, :client_random_id, :canonical_message_id, :send_state_id, :request_payload_hash)
func (m *defaultMessageClientRandomsTxModel) Insert(data *MessageClientRandoms) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into message_client_randoms(sender_user_id, peer_type, peer_id, client_random_id, canonical_message_id, send_state_id, request_payload_hash) values (:sender_user_id, :peer_type, :peer_id, :client_random_id, :canonical_message_id, :send_state_id, :request_payload_hash)"
		r     sql.Result
	)

	r, err = m.tx.NamedExec(query, data)
	if err != nil {
		err = fmt.Errorf("message_client_randoms.Insert named exec: %w", err)
		return
	}

	lastInsertId, err = r.LastInsertId()
	if err != nil {
		err = fmt.Errorf("message_client_randoms.Insert last insert id: %w", err)
		return
	}
	rowsAffected, err = r.RowsAffected()
	if err != nil {
		err = fmt.Errorf("message_client_randoms.Insert rows affected: %w", err)
	}

	return
}

// SelectByRandom
// select sender_user_id, peer_type, peer_id, client_random_id, canonical_message_id, send_state_id, request_payload_hash from message_client_randoms where sender_user_id = :sender_user_id and peer_type = :peer_type and peer_id = :peer_id and client_random_id = :client_random_id limit 1
func (m *defaultMessageClientRandomsModel) SelectByRandom(ctx context.Context, senderUserId int64, peerType int32, peerId int64, clientRandomId int64) (rValue *MessageClientRandoms, err error) {

	var (
		query = "select sender_user_id, peer_type, peer_id, client_random_id, canonical_message_id, send_state_id, request_payload_hash from message_client_randoms where sender_user_id = ? and peer_type = ? and peer_id = ? and client_random_id = ? limit 1"
		do    = &MessageClientRandoms{}
	)
	err = m.db.QueryRowPartial(ctx, do, query, senderUserId, peerType, peerId, clientRandomId)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return nil, &NotFoundError{
				Resource: "message_client_randoms",
				Key:      fmt.Sprintf("sender_user_id=%v,peer_type=%v,peer_id=%v,client_random_id=%v", senderUserId, peerType, peerId, clientRandomId),
				Cause:    err,
			}
		}
		err = fmt.Errorf("message_client_randoms.SelectByRandom: %w", err)
		return
	} else {
		rValue = do
	}

	return
}

// SelectByRandom
// select sender_user_id, peer_type, peer_id, client_random_id, canonical_message_id, send_state_id, request_payload_hash from message_client_randoms where sender_user_id = :sender_user_id and peer_type = :peer_type and peer_id = :peer_id and client_random_id = :client_random_id limit 1
func (m *defaultMessageClientRandomsTxModel) SelectByRandom(senderUserId int64, peerType int32, peerId int64, clientRandomId int64) (rValue *MessageClientRandoms, err error) {
	var (
		query = "select sender_user_id, peer_type, peer_id, client_random_id, canonical_message_id, send_state_id, request_payload_hash from message_client_randoms where sender_user_id = ? and peer_type = ? and peer_id = ? and client_random_id = ? limit 1"
		do    = &MessageClientRandoms{}
	)
	err = m.tx.QueryRowPartial(do, query, senderUserId, peerType, peerId, clientRandomId)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return nil, &NotFoundError{
				Resource: "message_client_randoms",
				Key:      fmt.Sprintf("sender_user_id=%v,peer_type=%v,peer_id=%v,client_random_id=%v", senderUserId, peerType, peerId, clientRandomId),
				Cause:    err,
			}
		}
		err = fmt.Errorf("message_client_randoms.SelectByRandom: %w", err)
		return
	}
	rValue = do

	return
}

// SelectByCanonicalMessageId
// select sender_user_id, peer_type, peer_id, client_random_id, canonical_message_id, send_state_id, request_payload_hash from message_client_randoms where canonical_message_id = :canonical_message_id limit 1
func (m *defaultMessageClientRandomsModel) SelectByCanonicalMessageId(ctx context.Context, canonicalMessageId int64) (rValue *MessageClientRandoms, err error) {

	var (
		query = "select sender_user_id, peer_type, peer_id, client_random_id, canonical_message_id, send_state_id, request_payload_hash from message_client_randoms where canonical_message_id = ? limit 1"
		do    = &MessageClientRandoms{}
	)
	err = m.db.QueryRowPartial(ctx, do, query, canonicalMessageId)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return nil, &NotFoundError{
				Resource: "message_client_randoms",
				Key:      fmt.Sprintf("canonical_message_id=%v", canonicalMessageId),
				Cause:    err,
			}
		}
		err = fmt.Errorf("message_client_randoms.SelectByCanonicalMessageId: %w", err)
		return
	} else {
		rValue = do
	}

	return
}

// SelectByCanonicalMessageId
// select sender_user_id, peer_type, peer_id, client_random_id, canonical_message_id, send_state_id, request_payload_hash from message_client_randoms where canonical_message_id = :canonical_message_id limit 1
func (m *defaultMessageClientRandomsTxModel) SelectByCanonicalMessageId(canonicalMessageId int64) (rValue *MessageClientRandoms, err error) {
	var (
		query = "select sender_user_id, peer_type, peer_id, client_random_id, canonical_message_id, send_state_id, request_payload_hash from message_client_randoms where canonical_message_id = ? limit 1"
		do    = &MessageClientRandoms{}
	)
	err = m.tx.QueryRowPartial(do, query, canonicalMessageId)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return nil, &NotFoundError{
				Resource: "message_client_randoms",
				Key:      fmt.Sprintf("canonical_message_id=%v", canonicalMessageId),
				Cause:    err,
			}
		}
		err = fmt.Errorf("message_client_randoms.SelectByCanonicalMessageId: %w", err)
		return
	}
	rValue = do

	return
}
