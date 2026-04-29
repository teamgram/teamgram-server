package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"math"

	"github.com/teamgram/marmota/pkg/stores/sqlx"
	"github.com/teamgram/teamgram-server/v2/app/messenger/msg/msg"
)

func (r *Repository) CreateOrLoadSendState(ctx context.Context, in CreateSendStateInput) (*SendState, error) {
	db, err := r.requireDB()
	if err != nil {
		return nil, err
	}
	existing, found, err := selectSendStateByRandom(ctx, db, in.SenderUserID, in.PeerType, in.PeerID, in.ClientRandomID)
	if err != nil {
		return nil, err
	}
	if found {
		if existing.RequestPayloadHash != in.RequestPayloadHash {
			return nil, msg.ErrRandomIdConflict
		}
		return existing, nil
	}

	sendStateID, err := r.nextID(ctx, "next send state id")
	if err != nil {
		return nil, err
	}
	_, err = db.Exec(ctx, `
INSERT INTO message_send_states
  (send_state_id, sender_user_id, peer_type, peer_id, client_random_id,
   canonical_message_id, peer_seq, status, request_payload_schema_version, request_payload_hash,
   sender_operation_id, sender_pts, sender_pts_count, sender_update_schema_version,
   sender_update_payload, sender_update_payload_hash, receiver_manifest_id,
   last_error_category, last_error_code, last_error_message, retry_count, created_at, updated_at, completed_at)
VALUES
  (?, ?, ?, ?, ?, NULL, NULL, ?, ?, UNHEX(?), NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, 0, NOW(6), NOW(6), NULL)
`, sendStateID, in.SenderUserID, in.PeerType, in.PeerID, in.ClientRandomID, SendStateStatusInitialized, in.RequestPayloadSchemaVersion, in.RequestPayloadHash)
	if err != nil {
		again, found, selectErr := selectSendStateByRandom(ctx, db, in.SenderUserID, in.PeerType, in.PeerID, in.ClientRandomID)
		if selectErr == nil && found {
			if again.RequestPayloadHash != in.RequestPayloadHash {
				return nil, msg.ErrRandomIdConflict
			}
			return again, nil
		}
		return nil, storageError("insert send state", err)
	}
	return selectSendStateByID(ctx, db, sendStateID)
}

func (r *Repository) MarkCanonicalCreated(ctx context.Context, sendStateID int64, canonicalMessageID int64, peerSeq int64) error {
	db, err := r.requireDB()
	if err != nil {
		return err
	}
	rows, err := db.Exec(ctx, `
UPDATE message_send_states
SET canonical_message_id = ?,
    peer_seq = ?,
    status = GREATEST(status, ?),
    updated_at = NOW(6)
WHERE send_state_id = ?
  AND (canonical_message_id IS NULL OR canonical_message_id = ?)
`, canonicalMessageID, peerSeq, SendStateStatusCanonical, sendStateID, canonicalMessageID)
	if err != nil {
		return storageError("mark canonical created", err)
	}
	return requireAffected(rows, "mark canonical created")
}

func (r *Repository) MarkSenderCommitted(ctx context.Context, in MarkSenderCommittedInput) error {
	if in.SenderPTS < math.MinInt32 || in.SenderPTS > math.MaxInt32 {
		return fmt.Errorf("%w: sender pts out of int32 range", msg.ErrSenderSyncFailed)
	}
	db, err := r.requireDB()
	if err != nil {
		return err
	}
	state, err := selectSendStateByID(ctx, db, in.SendStateID)
	if err != nil {
		return err
	}
	if state.PeerSeq < math.MinInt32 || state.PeerSeq > math.MaxInt32 {
		return fmt.Errorf("%w: peer seq out of int32 range", msg.ErrSenderSyncFailed)
	}
	if state.SenderOperationID != "" {
		if state.SenderOperationID != in.SenderOperationID || state.SenderUpdatePayloadHash != in.SenderUpdatePayloadHash {
			return msg.ErrSendStateConflict
		}
		return nil
	}
	rows, err := db.Exec(ctx, `
UPDATE message_send_states
SET sender_operation_id = ?,
    sender_pts = ?,
    sender_pts_count = ?,
    sender_update_schema_version = ?,
    sender_update_payload = ?,
    sender_update_payload_hash = UNHEX(?),
    status = GREATEST(status, ?),
    updated_at = NOW(6)
WHERE send_state_id = ?
`, in.SenderOperationID, in.SenderPTS, in.SenderPTSCount, in.SenderUpdateSchemaVersion, in.SenderUpdatePayload, in.SenderUpdatePayloadHash, SendStateStatusSenderCommitted, in.SendStateID)
	if err != nil {
		return storageError("mark sender committed", err)
	}
	return requireAffected(rows, "mark sender committed")
}

func (r *Repository) MarkReceiverOpsAcked(ctx context.Context, sendStateID int64, receiverManifestID int64) error {
	db, err := r.requireDB()
	if err != nil {
		return err
	}
	var manifest any
	if receiverManifestID != 0 {
		manifest = receiverManifestID
	}
	rows, err := db.Exec(ctx, `
UPDATE message_send_states
SET receiver_manifest_id = ?,
    status = GREATEST(status, ?),
    updated_at = NOW(6)
WHERE send_state_id = ?
`, manifest, SendStateStatusReceiverAcked, sendStateID)
	if err != nil {
		return storageError("mark receiver ops acked", err)
	}
	return requireAffected(rows, "mark receiver ops acked")
}

func (r *Repository) MarkCompleted(ctx context.Context, sendStateID int64) error {
	db, err := r.requireDB()
	if err != nil {
		return err
	}
	rows, err := db.Exec(ctx, `
UPDATE message_send_states
SET status = ?,
    completed_at = NOW(6),
    updated_at = NOW(6)
WHERE send_state_id = ?
`, SendStateStatusCompleted, sendStateID)
	if err != nil {
		return storageError("mark completed", err)
	}
	return requireAffected(rows, "mark completed")
}

func (r *Repository) MarkRetryableFailure(ctx context.Context, in MarkRetryableFailureInput) error {
	db, err := r.requireDB()
	if err != nil {
		return err
	}
	rows, err := db.Exec(ctx, `
UPDATE message_send_states
SET status = ?,
    last_error_category = ?,
    last_error_code = ?,
    last_error_message = ?,
    retry_count = retry_count + 1,
    updated_at = NOW(6)
WHERE send_state_id = ?
`, SendStateStatusFailedRetryable, in.LastErrorCategory, in.LastErrorCode, in.LastErrorMessage, in.SendStateID)
	if err != nil {
		return storageError("mark retryable failure", err)
	}
	return requireAffected(rows, "mark retryable failure")
}

func selectSendStateByRandom(ctx context.Context, db *sqlx.DB, senderUserID int64, peerType int32, peerID int64, clientRandomID int64) (*SendState, bool, error) {
	var row sendStateRow
	err := db.QueryRowPartial(ctx, &row, sendStateSelectSQL+`
WHERE sender_user_id = ?
  AND peer_type = ?
  AND peer_id = ?
  AND client_random_id = ?
LIMIT 1
`, senderUserID, peerType, peerID, clientRandomID)
	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return nil, false, nil
		}
		return nil, false, storageError("select send state by random", err)
	}
	return row.toSendState(), true, nil
}

func selectSendStateByID(ctx context.Context, db *sqlx.DB, sendStateID int64) (*SendState, error) {
	var row sendStateRow
	err := db.QueryRowPartial(ctx, &row, sendStateSelectSQL+`
WHERE send_state_id = ?
LIMIT 1
`, sendStateID)
	if err != nil {
		return nil, storageError("select send state by id", err)
	}
	return row.toSendState(), nil
}

func selectSendStateByIDTx(tx *sqlx.Tx, sendStateID int64) (*SendState, error) {
	var row sendStateRow
	err := tx.QueryRowPartial(&row, sendStateSelectSQL+`
WHERE send_state_id = ?
LIMIT 1 FOR UPDATE
`, sendStateID)
	if err != nil {
		return nil, storageError("select send state by id for update", err)
	}
	return row.toSendState(), nil
}

func requireAffected(result sql.Result, op string) error {
	affected, err := result.RowsAffected()
	if err != nil {
		return storageError(op+" rows affected", err)
	}
	if affected == 0 {
		return msg.ErrSendStateConflict
	}
	return nil
}

const sendStateSelectSQL = `
SELECT send_state_id,
       sender_user_id,
       peer_type,
       peer_id,
       client_random_id,
       IFNULL(canonical_message_id, 0) AS canonical_message_id,
       IFNULL(peer_seq, 0) AS peer_seq,
       status,
       request_payload_schema_version,
       LOWER(HEX(request_payload_hash)) AS request_payload_hash,
       IFNULL(sender_operation_id, '') AS sender_operation_id,
       IFNULL(sender_pts, 0) AS sender_pts,
       IFNULL(sender_pts_count, 0) AS sender_pts_count,
       IFNULL(sender_update_schema_version, 0) AS sender_update_schema_version,
       sender_update_payload,
       IFNULL(LOWER(HEX(sender_update_payload_hash)), '') AS sender_update_payload_hash,
       IFNULL(receiver_manifest_id, 0) AS receiver_manifest_id,
       retry_count
FROM message_send_states
`

type sendStateRow struct {
	SendStateID                 int64  `db:"send_state_id"`
	SenderUserID                int64  `db:"sender_user_id"`
	PeerType                    int32  `db:"peer_type"`
	PeerID                      int64  `db:"peer_id"`
	ClientRandomID              int64  `db:"client_random_id"`
	CanonicalMessageID          int64  `db:"canonical_message_id"`
	PeerSeq                     int64  `db:"peer_seq"`
	Status                      int32  `db:"status"`
	RequestPayloadSchemaVersion int32  `db:"request_payload_schema_version"`
	RequestPayloadHash          string `db:"request_payload_hash"`
	SenderOperationID           string `db:"sender_operation_id"`
	SenderPTS                   int64  `db:"sender_pts"`
	SenderPTSCount              int32  `db:"sender_pts_count"`
	SenderUpdateSchemaVersion   int32  `db:"sender_update_schema_version"`
	SenderUpdatePayload         []byte `db:"sender_update_payload"`
	SenderUpdatePayloadHash     string `db:"sender_update_payload_hash"`
	ReceiverManifestID          int64  `db:"receiver_manifest_id"`
	RetryCount                  int32  `db:"retry_count"`
}

func (r sendStateRow) toSendState() *SendState {
	return &SendState{
		SendStateID:                 r.SendStateID,
		SenderUserID:                r.SenderUserID,
		PeerType:                    r.PeerType,
		PeerID:                      r.PeerID,
		ClientRandomID:              r.ClientRandomID,
		CanonicalMessageID:          r.CanonicalMessageID,
		PeerSeq:                     r.PeerSeq,
		Status:                      r.Status,
		RequestPayloadSchemaVersion: r.RequestPayloadSchemaVersion,
		RequestPayloadHash:          r.RequestPayloadHash,
		SenderOperationID:           r.SenderOperationID,
		SenderPTS:                   r.SenderPTS,
		SenderPTSCount:              r.SenderPTSCount,
		SenderUpdateSchemaVersion:   r.SenderUpdateSchemaVersion,
		SenderUpdatePayload:         r.SenderUpdatePayload,
		SenderUpdatePayloadHash:     r.SenderUpdatePayloadHash,
		ReceiverManifestID:          r.ReceiverManifestID,
		RetryCount:                  r.RetryCount,
	}
}
