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
	"time"

	"github.com/teamgram/marmota/pkg/stores/sqlx"
)

var _ *sqlx.DB
var _ *sqlx.Tx
var _ time.Time

type CanonicalMessageRow struct {
	SendStateID        int64     `db:"send_state_id"`
	CanonicalMessageID int64     `db:"canonical_message_id"`
	PeerSeq            int64     `db:"peer_seq"`
	MessageDate        time.Time `db:"message_date"`
	RequestPayloadHash []byte    `db:"request_payload_hash"`
}

type CanonicalQueriesModel interface {
	SelectCanonicalByRandom(ctx context.Context, senderUserId int64, peerType int32, peerId int64, clientRandomId int64) (*CanonicalMessageRow, error)
	SelectCanonicalByID(ctx context.Context, sendStateId int64, requestPayloadHash []byte, canonicalMessageId int64) (*CanonicalMessageRow, error)
}

type CanonicalQueriesTxModel interface {
	SelectCanonicalByRandom(senderUserId int64, peerType int32, peerId int64, clientRandomId int64) (*CanonicalMessageRow, error)
	SelectCanonicalByID(sendStateId int64, requestPayloadHash []byte, canonicalMessageId int64) (*CanonicalMessageRow, error)
}

type defaultCanonicalQueriesModel struct {
	db *sqlx.DB
}

func NewCanonicalQueriesModel(db *sqlx.DB) CanonicalQueriesModel {
	return &defaultCanonicalQueriesModel{db: db}
}

type defaultCanonicalQueriesTxModel struct {
	tx *sqlx.Tx
}

func NewCanonicalQueriesTxModel(tx *sqlx.Tx) CanonicalQueriesTxModel {
	return &defaultCanonicalQueriesTxModel{tx: tx}
}

func (m *defaultCanonicalQueriesModel) SelectCanonicalByRandom(ctx context.Context, senderUserId int64, peerType int32, peerId int64, clientRandomId int64) (*CanonicalMessageRow, error) {
	var rValue CanonicalMessageRow
	query := "select r.send_state_id, r.canonical_message_id, c.peer_seq, c.`date` as message_date, r.request_payload_hash from message_client_randoms as r join canonical_messages as c on c.canonical_message_id = r.canonical_message_id where r.sender_user_id = ? and r.peer_type = ? and r.peer_id = ? and r.client_random_id = ? limit 1"

	err := m.db.QueryRowPartial(ctx, &rValue, query, senderUserId, peerType, peerId, clientRandomId)
	if err != nil {
		return nil, err
	}
	return &rValue, nil
}

func (m *defaultCanonicalQueriesTxModel) SelectCanonicalByRandom(senderUserId int64, peerType int32, peerId int64, clientRandomId int64) (*CanonicalMessageRow, error) {
	var rValue CanonicalMessageRow
	query := "select r.send_state_id, r.canonical_message_id, c.peer_seq, c.`date` as message_date, r.request_payload_hash from message_client_randoms as r join canonical_messages as c on c.canonical_message_id = r.canonical_message_id where r.sender_user_id = ? and r.peer_type = ? and r.peer_id = ? and r.client_random_id = ? limit 1"

	err := m.tx.QueryRowPartial(&rValue, query, senderUserId, peerType, peerId, clientRandomId)
	if err != nil {
		return nil, err
	}
	return &rValue, nil
}

func (m *defaultCanonicalQueriesModel) SelectCanonicalByID(ctx context.Context, sendStateId int64, requestPayloadHash []byte, canonicalMessageId int64) (*CanonicalMessageRow, error) {
	var rValue CanonicalMessageRow
	query := "select ? as send_state_id, canonical_message_id, peer_seq, `date` as message_date, ? as request_payload_hash from canonical_messages where canonical_message_id = ? limit 1"

	err := m.db.QueryRowPartial(ctx, &rValue, query, sendStateId, requestPayloadHash, canonicalMessageId)
	if err != nil {
		return nil, err
	}
	return &rValue, nil
}

func (m *defaultCanonicalQueriesTxModel) SelectCanonicalByID(sendStateId int64, requestPayloadHash []byte, canonicalMessageId int64) (*CanonicalMessageRow, error) {
	var rValue CanonicalMessageRow
	query := "select ? as send_state_id, canonical_message_id, peer_seq, `date` as message_date, ? as request_payload_hash from canonical_messages where canonical_message_id = ? limit 1"

	err := m.tx.QueryRowPartial(&rValue, query, sendStateId, requestPayloadHash, canonicalMessageId)
	if err != nil {
		return nil, err
	}
	return &rValue, nil
}
