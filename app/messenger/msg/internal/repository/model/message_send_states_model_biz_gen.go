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
	bizMessageSendStatesModel interface {
		Insert(ctx context.Context, data *MessageSendStates) (lastInsertId, rowsAffected int64, err error)
		InsertTx(tx *sqlx.Tx, data *MessageSendStates) (lastInsertId, rowsAffected int64, err error)

		SelectBySendStateId(ctx context.Context, sendStateId int64) (*MessageSendStates, error)
		SelectBySendStateIdTx(tx *sqlx.Tx, sendStateId int64) (*MessageSendStates, error)

		SelectByRandom(ctx context.Context, senderUserId int64, peerType int32, peerId int64, clientRandomId int64) (*MessageSendStates, error)
		SelectByRandomTx(tx *sqlx.Tx, senderUserId int64, peerType int32, peerId int64, clientRandomId int64) (*MessageSendStates, error)

		MarkCanonicalCreated(ctx context.Context, canonicalMessageId int64, peerSeq int64, status int32, sendStateId int64) (rowsAffected int64, err error)
		MarkCanonicalCreatedTx(tx *sqlx.Tx, canonicalMessageId int64, peerSeq int64, status int32, sendStateId int64) (rowsAffected int64, err error)

		MarkSenderCommitted(ctx context.Context, senderOperationId string, senderPts int64, senderPtsCount int32, senderUpdateSchemaVersion int32, senderUpdatePayload []byte, senderUpdatePayloadHash []byte, status int32, sendStateId int64) (rowsAffected int64, err error)
		MarkSenderCommittedTx(tx *sqlx.Tx, senderOperationId string, senderPts int64, senderPtsCount int32, senderUpdateSchemaVersion int32, senderUpdatePayload []byte, senderUpdatePayloadHash []byte, status int32, sendStateId int64) (rowsAffected int64, err error)

		MarkReceiverOpsAcked(ctx context.Context, receiverManifestId int64, status int32, sendStateId int64) (rowsAffected int64, err error)
		MarkReceiverOpsAckedTx(tx *sqlx.Tx, receiverManifestId int64, status int32, sendStateId int64) (rowsAffected int64, err error)

		MarkCompleted(ctx context.Context, status int32, completedAt string, sendStateId int64) (rowsAffected int64, err error)
		MarkCompletedTx(tx *sqlx.Tx, status int32, completedAt string, sendStateId int64) (rowsAffected int64, err error)

		MarkRetryableFailure(ctx context.Context, status int32, lastErrorCategory int32, lastErrorCode string, lastErrorMessage string, sendStateId int64) (rowsAffected int64, err error)
		MarkRetryableFailureTx(tx *sqlx.Tx, status int32, lastErrorCategory int32, lastErrorCode string, lastErrorMessage string, sendStateId int64) (rowsAffected int64, err error)
	}
)

// Insert
// insert into message_send_states(send_state_id, sender_user_id, peer_type, peer_id, client_random_id, canonical_message_id, peer_seq, `status`, request_payload_schema_version, request_payload_hash, sender_pts, sender_pts_count, sender_update_schema_version, sender_update_payload, sender_update_payload_hash, receiver_manifest_id, last_error_category, last_error_code, last_error_message, retry_count) values (:send_state_id, :sender_user_id, :peer_type, :peer_id, :client_random_id, :canonical_message_id, :peer_seq, :status, :request_payload_schema_version, :request_payload_hash, :sender_pts, :sender_pts_count, :sender_update_schema_version, :sender_update_payload, :sender_update_payload_hash, :receiver_manifest_id, :last_error_category, :last_error_code, :last_error_message, :retry_count)
func (m *defaultMessageSendStatesModel) Insert(ctx context.Context, data *MessageSendStates) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into message_send_states(send_state_id, sender_user_id, peer_type, peer_id, client_random_id, canonical_message_id, peer_seq, `status`, request_payload_schema_version, request_payload_hash, sender_pts, sender_pts_count, sender_update_schema_version, sender_update_payload, sender_update_payload_hash, receiver_manifest_id, last_error_category, last_error_code, last_error_message, retry_count) values (:send_state_id, :sender_user_id, :peer_type, :peer_id, :client_random_id, :canonical_message_id, :peer_seq, :status, :request_payload_schema_version, :request_payload_hash, :sender_pts, :sender_pts_count, :sender_update_schema_version, :sender_update_payload, :sender_update_payload_hash, :receiver_manifest_id, :last_error_category, :last_error_code, :last_error_message, :retry_count)"
		r     sql.Result
	)

	r, err = m.db.NamedExec(ctx, query, data)
	if err != nil {
		err = fmt.Errorf("message_send_states.Insert named exec: %w", err)
		return
	}

	lastInsertId, err = r.LastInsertId()
	if err != nil {
		err = fmt.Errorf("message_send_states.Insert last insert id: %w", err)
		return
	}
	rowsAffected, err = r.RowsAffected()
	if err != nil {
		err = fmt.Errorf("message_send_states.Insert rows affected: %w", err)
	}

	return

}

// InsertTx
// insert into message_send_states(send_state_id, sender_user_id, peer_type, peer_id, client_random_id, canonical_message_id, peer_seq, `status`, request_payload_schema_version, request_payload_hash, sender_pts, sender_pts_count, sender_update_schema_version, sender_update_payload, sender_update_payload_hash, receiver_manifest_id, last_error_category, last_error_code, last_error_message, retry_count) values (:send_state_id, :sender_user_id, :peer_type, :peer_id, :client_random_id, :canonical_message_id, :peer_seq, :status, :request_payload_schema_version, :request_payload_hash, :sender_pts, :sender_pts_count, :sender_update_schema_version, :sender_update_payload, :sender_update_payload_hash, :receiver_manifest_id, :last_error_category, :last_error_code, :last_error_message, :retry_count)
func (m *defaultMessageSendStatesModel) InsertTx(tx *sqlx.Tx, data *MessageSendStates) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into message_send_states(send_state_id, sender_user_id, peer_type, peer_id, client_random_id, canonical_message_id, peer_seq, `status`, request_payload_schema_version, request_payload_hash, sender_pts, sender_pts_count, sender_update_schema_version, sender_update_payload, sender_update_payload_hash, receiver_manifest_id, last_error_category, last_error_code, last_error_message, retry_count) values (:send_state_id, :sender_user_id, :peer_type, :peer_id, :client_random_id, :canonical_message_id, :peer_seq, :status, :request_payload_schema_version, :request_payload_hash, :sender_pts, :sender_pts_count, :sender_update_schema_version, :sender_update_payload, :sender_update_payload_hash, :receiver_manifest_id, :last_error_category, :last_error_code, :last_error_message, :retry_count)"
		r     sql.Result
	)

	r, err = tx.NamedExec(query, data)
	if err != nil {
		err = fmt.Errorf("message_send_states.InsertTx named exec: %w", err)
		return
	}

	lastInsertId, err = r.LastInsertId()
	if err != nil {
		err = fmt.Errorf("message_send_states.InsertTx last insert id: %w", err)
		return
	}
	rowsAffected, err = r.RowsAffected()
	if err != nil {
		err = fmt.Errorf("message_send_states.InsertTx rows affected: %w", err)
	}

	return
}

// SelectBySendStateId
// select send_state_id, sender_user_id, peer_type, peer_id, client_random_id, canonical_message_id, peer_seq, `status`, request_payload_schema_version, request_payload_hash, sender_pts, sender_pts_count, sender_update_schema_version, sender_update_payload, sender_update_payload_hash, receiver_manifest_id, last_error_category, last_error_code, last_error_message, retry_count from message_send_states where send_state_id = :send_state_id limit 1
func (m *defaultMessageSendStatesModel) SelectBySendStateId(ctx context.Context, sendStateId int64) (rValue *MessageSendStates, err error) {

	var (
		query = "select send_state_id, sender_user_id, peer_type, peer_id, client_random_id, canonical_message_id, peer_seq, `status`, request_payload_schema_version, request_payload_hash, sender_pts, sender_pts_count, sender_update_schema_version, sender_update_payload, sender_update_payload_hash, receiver_manifest_id, last_error_category, last_error_code, last_error_message, retry_count from message_send_states where send_state_id = ? limit 1"
		do    = &MessageSendStates{}
	)
	err = m.db.QueryRowPartial(ctx, do, query, sendStateId)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return nil, &NotFoundError{
				Resource: "message_send_states",
				Key:      fmt.Sprintf("send_state_id=%v", sendStateId),
				Cause:    err,
			}
		}
		err = fmt.Errorf("message_send_states.SelectBySendStateId: %w", err)
		return
	} else {
		rValue = do
	}

	return
}

// SelectBySendStateIdTx
// select send_state_id, sender_user_id, peer_type, peer_id, client_random_id, canonical_message_id, peer_seq, `status`, request_payload_schema_version, request_payload_hash, sender_pts, sender_pts_count, sender_update_schema_version, sender_update_payload, sender_update_payload_hash, receiver_manifest_id, last_error_category, last_error_code, last_error_message, retry_count from message_send_states where send_state_id = :send_state_id limit 1
func (m *defaultMessageSendStatesModel) SelectBySendStateIdTx(tx *sqlx.Tx, sendStateId int64) (rValue *MessageSendStates, err error) {
	var (
		query = "select send_state_id, sender_user_id, peer_type, peer_id, client_random_id, canonical_message_id, peer_seq, `status`, request_payload_schema_version, request_payload_hash, sender_pts, sender_pts_count, sender_update_schema_version, sender_update_payload, sender_update_payload_hash, receiver_manifest_id, last_error_category, last_error_code, last_error_message, retry_count from message_send_states where send_state_id = ? limit 1"
		do    = &MessageSendStates{}
	)
	err = tx.QueryRowPartial(do, query, sendStateId)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return nil, &NotFoundError{
				Resource: "message_send_states",
				Key:      fmt.Sprintf("send_state_id=%v", sendStateId),
				Cause:    err,
			}
		}
		err = fmt.Errorf("message_send_states.SelectBySendStateIdTx: %w", err)
		return
	}
	rValue = do

	return
}

// SelectByRandom
// select send_state_id, sender_user_id, peer_type, peer_id, client_random_id, canonical_message_id, peer_seq, `status`, request_payload_schema_version, request_payload_hash, sender_pts, sender_pts_count, sender_update_schema_version, sender_update_payload, sender_update_payload_hash, receiver_manifest_id, last_error_category, last_error_code, last_error_message, retry_count from message_send_states where sender_user_id = :sender_user_id and peer_type = :peer_type and peer_id = :peer_id and client_random_id = :client_random_id limit 1
func (m *defaultMessageSendStatesModel) SelectByRandom(ctx context.Context, senderUserId int64, peerType int32, peerId int64, clientRandomId int64) (rValue *MessageSendStates, err error) {

	var (
		query = "select send_state_id, sender_user_id, peer_type, peer_id, client_random_id, canonical_message_id, peer_seq, `status`, request_payload_schema_version, request_payload_hash, sender_pts, sender_pts_count, sender_update_schema_version, sender_update_payload, sender_update_payload_hash, receiver_manifest_id, last_error_category, last_error_code, last_error_message, retry_count from message_send_states where sender_user_id = ? and peer_type = ? and peer_id = ? and client_random_id = ? limit 1"
		do    = &MessageSendStates{}
	)
	err = m.db.QueryRowPartial(ctx, do, query, senderUserId, peerType, peerId, clientRandomId)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return nil, &NotFoundError{
				Resource: "message_send_states",
				Key:      fmt.Sprintf("sender_user_id=%v,peer_type=%v,peer_id=%v,client_random_id=%v", senderUserId, peerType, peerId, clientRandomId),
				Cause:    err,
			}
		}
		err = fmt.Errorf("message_send_states.SelectByRandom: %w", err)
		return
	} else {
		rValue = do
	}

	return
}

// SelectByRandomTx
// select send_state_id, sender_user_id, peer_type, peer_id, client_random_id, canonical_message_id, peer_seq, `status`, request_payload_schema_version, request_payload_hash, sender_pts, sender_pts_count, sender_update_schema_version, sender_update_payload, sender_update_payload_hash, receiver_manifest_id, last_error_category, last_error_code, last_error_message, retry_count from message_send_states where sender_user_id = :sender_user_id and peer_type = :peer_type and peer_id = :peer_id and client_random_id = :client_random_id limit 1
func (m *defaultMessageSendStatesModel) SelectByRandomTx(tx *sqlx.Tx, senderUserId int64, peerType int32, peerId int64, clientRandomId int64) (rValue *MessageSendStates, err error) {
	var (
		query = "select send_state_id, sender_user_id, peer_type, peer_id, client_random_id, canonical_message_id, peer_seq, `status`, request_payload_schema_version, request_payload_hash, sender_pts, sender_pts_count, sender_update_schema_version, sender_update_payload, sender_update_payload_hash, receiver_manifest_id, last_error_category, last_error_code, last_error_message, retry_count from message_send_states where sender_user_id = ? and peer_type = ? and peer_id = ? and client_random_id = ? limit 1"
		do    = &MessageSendStates{}
	)
	err = tx.QueryRowPartial(do, query, senderUserId, peerType, peerId, clientRandomId)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return nil, &NotFoundError{
				Resource: "message_send_states",
				Key:      fmt.Sprintf("sender_user_id=%v,peer_type=%v,peer_id=%v,client_random_id=%v", senderUserId, peerType, peerId, clientRandomId),
				Cause:    err,
			}
		}
		err = fmt.Errorf("message_send_states.SelectByRandomTx: %w", err)
		return
	}
	rValue = do

	return
}

// MarkCanonicalCreated
// update message_send_states set canonical_message_id = :canonical_message_id, peer_seq = :peer_seq, `status` = :status where send_state_id = :send_state_id
func (m *defaultMessageSendStatesModel) MarkCanonicalCreated(ctx context.Context, canonicalMessageId int64, peerSeq int64, status int32, sendStateId int64) (rowsAffected int64, err error) {

	var (
		query   = "update message_send_states set canonical_message_id = ?, peer_seq = ?, `status` = ? where send_state_id = ?"
		rResult sql.Result
	)

	rResult, err = m.db.Exec(ctx, query, canonicalMessageId, peerSeq, status, sendStateId)

	if err != nil {
		err = fmt.Errorf("message_send_states.MarkCanonicalCreated exec: %w", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		err = fmt.Errorf("message_send_states.MarkCanonicalCreated rows affected: %w", err)
		return
	}

	return
}

// MarkCanonicalCreatedTx
// update message_send_states set canonical_message_id = :canonical_message_id, peer_seq = :peer_seq, `status` = :status where send_state_id = :send_state_id
func (m *defaultMessageSendStatesModel) MarkCanonicalCreatedTx(tx *sqlx.Tx, canonicalMessageId int64, peerSeq int64, status int32, sendStateId int64) (rowsAffected int64, err error) {
	var (
		query   = "update message_send_states set canonical_message_id = ?, peer_seq = ?, `status` = ? where send_state_id = ?"
		rResult sql.Result
	)
	rResult, err = tx.Exec(query, canonicalMessageId, peerSeq, status, sendStateId)

	if err != nil {
		err = fmt.Errorf("message_send_states.MarkCanonicalCreatedTx exec: %w", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		err = fmt.Errorf("message_send_states.MarkCanonicalCreatedTx rows affected: %w", err)
		return
	}

	return
}

// MarkSenderCommitted
// update message_send_states set sender_operation_id = :sender_operation_id, sender_pts = :sender_pts, sender_pts_count = :sender_pts_count, sender_update_schema_version = :sender_update_schema_version, sender_update_payload = :sender_update_payload, sender_update_payload_hash = :sender_update_payload_hash, `status` = :status where send_state_id = :send_state_id and (sender_operation_id is null or sender_operation_id = :sender_operation_id)
func (m *defaultMessageSendStatesModel) MarkSenderCommitted(ctx context.Context, senderOperationId string, senderPts int64, senderPtsCount int32, senderUpdateSchemaVersion int32, senderUpdatePayload []byte, senderUpdatePayloadHash []byte, status int32, sendStateId int64) (rowsAffected int64, err error) {

	var (
		query   = "update message_send_states set sender_operation_id = ?, sender_pts = ?, sender_pts_count = ?, sender_update_schema_version = ?, sender_update_payload = ?, sender_update_payload_hash = ?, `status` = ? where send_state_id = ? and (sender_operation_id is null or sender_operation_id = ?)"
		rResult sql.Result
	)

	rResult, err = m.db.Exec(ctx, query, senderOperationId, senderPts, senderPtsCount, senderUpdateSchemaVersion, senderUpdatePayload, senderUpdatePayloadHash, status, sendStateId, senderOperationId)

	if err != nil {
		err = fmt.Errorf("message_send_states.MarkSenderCommitted exec: %w", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		err = fmt.Errorf("message_send_states.MarkSenderCommitted rows affected: %w", err)
		return
	}

	return
}

// MarkSenderCommittedTx
// update message_send_states set sender_operation_id = :sender_operation_id, sender_pts = :sender_pts, sender_pts_count = :sender_pts_count, sender_update_schema_version = :sender_update_schema_version, sender_update_payload = :sender_update_payload, sender_update_payload_hash = :sender_update_payload_hash, `status` = :status where send_state_id = :send_state_id and (sender_operation_id is null or sender_operation_id = :sender_operation_id)
func (m *defaultMessageSendStatesModel) MarkSenderCommittedTx(tx *sqlx.Tx, senderOperationId string, senderPts int64, senderPtsCount int32, senderUpdateSchemaVersion int32, senderUpdatePayload []byte, senderUpdatePayloadHash []byte, status int32, sendStateId int64) (rowsAffected int64, err error) {
	var (
		query   = "update message_send_states set sender_operation_id = ?, sender_pts = ?, sender_pts_count = ?, sender_update_schema_version = ?, sender_update_payload = ?, sender_update_payload_hash = ?, `status` = ? where send_state_id = ? and (sender_operation_id is null or sender_operation_id = ?)"
		rResult sql.Result
	)
	rResult, err = tx.Exec(query, senderOperationId, senderPts, senderPtsCount, senderUpdateSchemaVersion, senderUpdatePayload, senderUpdatePayloadHash, status, sendStateId, senderOperationId)

	if err != nil {
		err = fmt.Errorf("message_send_states.MarkSenderCommittedTx exec: %w", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		err = fmt.Errorf("message_send_states.MarkSenderCommittedTx rows affected: %w", err)
		return
	}

	return
}

// MarkReceiverOpsAcked
// update message_send_states set receiver_manifest_id = :receiver_manifest_id, `status` = :status where send_state_id = :send_state_id
func (m *defaultMessageSendStatesModel) MarkReceiverOpsAcked(ctx context.Context, receiverManifestId int64, status int32, sendStateId int64) (rowsAffected int64, err error) {

	var (
		query   = "update message_send_states set receiver_manifest_id = ?, `status` = ? where send_state_id = ?"
		rResult sql.Result
	)

	rResult, err = m.db.Exec(ctx, query, receiverManifestId, status, sendStateId)

	if err != nil {
		err = fmt.Errorf("message_send_states.MarkReceiverOpsAcked exec: %w", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		err = fmt.Errorf("message_send_states.MarkReceiverOpsAcked rows affected: %w", err)
		return
	}

	return
}

// MarkReceiverOpsAckedTx
// update message_send_states set receiver_manifest_id = :receiver_manifest_id, `status` = :status where send_state_id = :send_state_id
func (m *defaultMessageSendStatesModel) MarkReceiverOpsAckedTx(tx *sqlx.Tx, receiverManifestId int64, status int32, sendStateId int64) (rowsAffected int64, err error) {
	var (
		query   = "update message_send_states set receiver_manifest_id = ?, `status` = ? where send_state_id = ?"
		rResult sql.Result
	)
	rResult, err = tx.Exec(query, receiverManifestId, status, sendStateId)

	if err != nil {
		err = fmt.Errorf("message_send_states.MarkReceiverOpsAckedTx exec: %w", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		err = fmt.Errorf("message_send_states.MarkReceiverOpsAckedTx rows affected: %w", err)
		return
	}

	return
}

// MarkCompleted
// update message_send_states set `status` = :status, completed_at = :completed_at where send_state_id = :send_state_id
func (m *defaultMessageSendStatesModel) MarkCompleted(ctx context.Context, status int32, completedAt string, sendStateId int64) (rowsAffected int64, err error) {

	var (
		query   = "update message_send_states set `status` = ?, completed_at = ? where send_state_id = ?"
		rResult sql.Result
	)

	rResult, err = m.db.Exec(ctx, query, status, completedAt, sendStateId)

	if err != nil {
		err = fmt.Errorf("message_send_states.MarkCompleted exec: %w", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		err = fmt.Errorf("message_send_states.MarkCompleted rows affected: %w", err)
		return
	}

	return
}

// MarkCompletedTx
// update message_send_states set `status` = :status, completed_at = :completed_at where send_state_id = :send_state_id
func (m *defaultMessageSendStatesModel) MarkCompletedTx(tx *sqlx.Tx, status int32, completedAt string, sendStateId int64) (rowsAffected int64, err error) {
	var (
		query   = "update message_send_states set `status` = ?, completed_at = ? where send_state_id = ?"
		rResult sql.Result
	)
	rResult, err = tx.Exec(query, status, completedAt, sendStateId)

	if err != nil {
		err = fmt.Errorf("message_send_states.MarkCompletedTx exec: %w", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		err = fmt.Errorf("message_send_states.MarkCompletedTx rows affected: %w", err)
		return
	}

	return
}

// MarkRetryableFailure
// update message_send_states set `status` = :status, last_error_category = :last_error_category, last_error_code = :last_error_code, last_error_message = :last_error_message, retry_count = retry_count + 1 where send_state_id = :send_state_id
func (m *defaultMessageSendStatesModel) MarkRetryableFailure(ctx context.Context, status int32, lastErrorCategory int32, lastErrorCode string, lastErrorMessage string, sendStateId int64) (rowsAffected int64, err error) {

	var (
		query   = "update message_send_states set `status` = ?, last_error_category = ?, last_error_code = ?, last_error_message = ?, retry_count = retry_count + 1 where send_state_id = ?"
		rResult sql.Result
	)

	rResult, err = m.db.Exec(ctx, query, status, lastErrorCategory, lastErrorCode, lastErrorMessage, sendStateId)

	if err != nil {
		err = fmt.Errorf("message_send_states.MarkRetryableFailure exec: %w", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		err = fmt.Errorf("message_send_states.MarkRetryableFailure rows affected: %w", err)
		return
	}

	return
}

// MarkRetryableFailureTx
// update message_send_states set `status` = :status, last_error_category = :last_error_category, last_error_code = :last_error_code, last_error_message = :last_error_message, retry_count = retry_count + 1 where send_state_id = :send_state_id
func (m *defaultMessageSendStatesModel) MarkRetryableFailureTx(tx *sqlx.Tx, status int32, lastErrorCategory int32, lastErrorCode string, lastErrorMessage string, sendStateId int64) (rowsAffected int64, err error) {
	var (
		query   = "update message_send_states set `status` = ?, last_error_category = ?, last_error_code = ?, last_error_message = ?, retry_count = retry_count + 1 where send_state_id = ?"
		rResult sql.Result
	)
	rResult, err = tx.Exec(query, status, lastErrorCategory, lastErrorCode, lastErrorMessage, sendStateId)

	if err != nil {
		err = fmt.Errorf("message_send_states.MarkRetryableFailureTx exec: %w", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		err = fmt.Errorf("message_send_states.MarkRetryableFailureTx rows affected: %w", err)
		return
	}

	return
}
