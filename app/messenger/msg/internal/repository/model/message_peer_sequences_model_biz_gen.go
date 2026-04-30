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

type (
	bizMessagePeerSequencesModel interface {
		InsertIgnore(ctx context.Context, data *MessagePeerSequences) (lastInsertId, rowsAffected int64, err error)
		InsertIgnoreTx(tx *sqlx.Tx, data *MessagePeerSequences) (lastInsertId, rowsAffected int64, err error)

		SelectForUpdate(ctx context.Context, peerType int32, peerId int64) (*MessagePeerSequences, error)
		SelectForUpdateTx(tx *sqlx.Tx, peerType int32, peerId int64) (*MessagePeerSequences, error)

		UpdateNextPeerSeq(ctx context.Context, nextPeerSeq int64, peerType int32, peerId int64) (rowsAffected int64, err error)
		UpdateNextPeerSeqTx(tx *sqlx.Tx, nextPeerSeq int64, peerType int32, peerId int64) (rowsAffected int64, err error)
	}
)

// InsertIgnore
// insert ignore into message_peer_sequences(peer_type, peer_id, next_peer_seq) values (:peer_type, :peer_id, :next_peer_seq)
func (m *defaultMessagePeerSequencesModel) InsertIgnore(ctx context.Context, data *MessagePeerSequences) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert ignore into message_peer_sequences(peer_type, peer_id, next_peer_seq) values (:peer_type, :peer_id, :next_peer_seq)"
		r     sql.Result
	)

	r, err = m.db.NamedExec(ctx, query, data)
	if err != nil {
		err = fmt.Errorf("message_peer_sequences.InsertIgnore named exec: %w", err)
		return
	}

	lastInsertId, err = r.LastInsertId()
	if err != nil {
		err = fmt.Errorf("message_peer_sequences.InsertIgnore last insert id: %w", err)
		return
	}
	rowsAffected, err = r.RowsAffected()
	if err != nil {
		err = fmt.Errorf("message_peer_sequences.InsertIgnore rows affected: %w", err)
	}

	return

}

// InsertIgnoreTx
// insert ignore into message_peer_sequences(peer_type, peer_id, next_peer_seq) values (:peer_type, :peer_id, :next_peer_seq)
func (m *defaultMessagePeerSequencesModel) InsertIgnoreTx(tx *sqlx.Tx, data *MessagePeerSequences) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert ignore into message_peer_sequences(peer_type, peer_id, next_peer_seq) values (:peer_type, :peer_id, :next_peer_seq)"
		r     sql.Result
	)

	r, err = tx.NamedExec(query, data)
	if err != nil {
		err = fmt.Errorf("message_peer_sequences.InsertIgnoreTx named exec: %w", err)
		return
	}

	lastInsertId, err = r.LastInsertId()
	if err != nil {
		err = fmt.Errorf("message_peer_sequences.InsertIgnoreTx last insert id: %w", err)
		return
	}
	rowsAffected, err = r.RowsAffected()
	if err != nil {
		err = fmt.Errorf("message_peer_sequences.InsertIgnoreTx rows affected: %w", err)
	}

	return
}

// SelectForUpdate
// select peer_type, peer_id, next_peer_seq from message_peer_sequences where peer_type = :peer_type and peer_id = :peer_id limit 1 for update
func (m *defaultMessagePeerSequencesModel) SelectForUpdate(ctx context.Context, peerType int32, peerId int64) (rValue *MessagePeerSequences, err error) {

	var (
		query = "select peer_type, peer_id, next_peer_seq from message_peer_sequences where peer_type = ? and peer_id = ? limit 1 for update"
		do    = &MessagePeerSequences{}
	)
	err = m.db.QueryRowPartial(ctx, do, query, peerType, peerId)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return nil, &NotFoundError{
				Resource: "message_peer_sequences",
				Key:      fmt.Sprintf("peer_type=%v,peer_id=%v", peerType, peerId),
				Cause:    err,
			}
		}
		err = fmt.Errorf("message_peer_sequences.SelectForUpdate: %w", err)
		return
	} else {
		rValue = do
	}

	return
}

// SelectForUpdateTx
// select peer_type, peer_id, next_peer_seq from message_peer_sequences where peer_type = :peer_type and peer_id = :peer_id limit 1 for update
func (m *defaultMessagePeerSequencesModel) SelectForUpdateTx(tx *sqlx.Tx, peerType int32, peerId int64) (rValue *MessagePeerSequences, err error) {
	var (
		query = "select peer_type, peer_id, next_peer_seq from message_peer_sequences where peer_type = ? and peer_id = ? limit 1 for update"
		do    = &MessagePeerSequences{}
	)
	err = tx.QueryRowPartial(do, query, peerType, peerId)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return nil, &NotFoundError{
				Resource: "message_peer_sequences",
				Key:      fmt.Sprintf("peer_type=%v,peer_id=%v", peerType, peerId),
				Cause:    err,
			}
		}
		err = fmt.Errorf("message_peer_sequences.SelectForUpdateTx: %w", err)
		return
	}
	rValue = do

	return
}

// UpdateNextPeerSeq
// update message_peer_sequences set next_peer_seq = :next_peer_seq where peer_type = :peer_type and peer_id = :peer_id
func (m *defaultMessagePeerSequencesModel) UpdateNextPeerSeq(ctx context.Context, nextPeerSeq int64, peerType int32, peerId int64) (rowsAffected int64, err error) {

	var (
		query   = "update message_peer_sequences set next_peer_seq = ? where peer_type = ? and peer_id = ?"
		rResult sql.Result
	)

	rResult, err = m.db.Exec(ctx, query, nextPeerSeq, peerType, peerId)

	if err != nil {
		err = fmt.Errorf("message_peer_sequences.UpdateNextPeerSeq exec: %w", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		err = fmt.Errorf("message_peer_sequences.UpdateNextPeerSeq rows affected: %w", err)
		return
	}

	return
}

// UpdateNextPeerSeqTx
// update message_peer_sequences set next_peer_seq = :next_peer_seq where peer_type = :peer_type and peer_id = :peer_id
func (m *defaultMessagePeerSequencesModel) UpdateNextPeerSeqTx(tx *sqlx.Tx, nextPeerSeq int64, peerType int32, peerId int64) (rowsAffected int64, err error) {
	var (
		query   = "update message_peer_sequences set next_peer_seq = ? where peer_type = ? and peer_id = ?"
		rResult sql.Result
	)
	rResult, err = tx.Exec(query, nextPeerSeq, peerType, peerId)

	if err != nil {
		err = fmt.Errorf("message_peer_sequences.UpdateNextPeerSeqTx exec: %w", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		err = fmt.Errorf("message_peer_sequences.UpdateNextPeerSeqTx rows affected: %w", err)
		return
	}

	return
}
