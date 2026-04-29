package repository

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/teamgram/marmota/pkg/stores/sqlx"
	"github.com/teamgram/teamgram-server/v2/app/messenger/userupdates/payload"
	"github.com/teamgram/teamgram-server/v2/app/messenger/userupdates/userupdates"
)

type partitionFenceRow struct {
	PartitionID     int32  `db:"partition_id"`
	OwnerEpoch      int64  `db:"owner_epoch"`
	OwnerInstanceID string `db:"owner_instance_id"`
}

type userPTSStateRow struct {
	UserID      int64 `db:"user_id"`
	Pts         int64 `db:"pts"`
	PartitionID int32 `db:"partition_id"`
	OwnerEpoch  int64 `db:"owner_epoch"`
	RowVersion  int64 `db:"row_version"`
}

func (r *Repository) ClaimPartitionOwner(ctx context.Context, partitionID int32) (int64, error) {
	db, err := r.requireDB()
	if err != nil {
		return 0, err
	}
	if _, err := db.Exec(ctx, `
INSERT IGNORE INTO userupdates_partition_fences
  (partition_id, owner_epoch, owner_instance_id, lease_id, lease_expires_at, created_at, updated_at)
VALUES
  (?, 0, 'unassigned', '', NULL, NOW(6), NOW(6))
`, partitionID); err != nil {
		return 0, storageError("insert partition fence", err)
	}

	var fence partitionFenceRow
	if err := db.QueryRowPartial(ctx, &fence, `
SELECT partition_id, owner_epoch, owner_instance_id
FROM userupdates_partition_fences
WHERE partition_id = ?
LIMIT 1
`, partitionID); err != nil {
		return 0, storageError("select partition fence", err)
	}

	rows, err := db.Exec(ctx, `
UPDATE userupdates_partition_fences
SET owner_epoch = owner_epoch + 1,
    owner_instance_id = ?,
    lease_id = '',
    lease_expires_at = NULL,
    updated_at = NOW(6)
WHERE partition_id = ?
  AND owner_epoch = ?
`, r.OwnerInstance(), partitionID, fence.OwnerEpoch)
	if err != nil {
		return 0, storageError("claim partition owner", err)
	}
	affected, err := rows.RowsAffected()
	if err != nil {
		return 0, storageError("claim partition owner rows affected", err)
	}
	if affected == 0 {
		return 0, userupdates.ErrOwnerFenceFailed
	}
	return fence.OwnerEpoch + 1, nil
}

func (r *Repository) ApplyUserOperation(ctx context.Context, in ApplyUserOperationInput) (*ApplyUserOperationResult, error) {
	db, err := r.requireDB()
	if err != nil {
		return nil, err
	}
	var out *ApplyUserOperationResult
	err = db.Transact(ctx, func(tx *sqlx.Tx) error {
		result, err := r.applyUserOperationTx(ctx, tx, in)
		if err != nil {
			return err
		}
		out = result
		return nil
	})
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (r *Repository) applyUserOperationTx(ctx context.Context, tx *sqlx.Tx, in ApplyUserOperationInput) (*ApplyUserOperationResult, error) {
	fence, err := lockPartitionFence(tx, in.PartitionID)
	if err != nil {
		return nil, err
	}
	if fence.OwnerInstanceID != r.OwnerInstance() {
		return nil, userupdates.ErrNotOwner
	}

	if _, err := tx.Exec(`
INSERT INTO user_pts_state
  (user_id, pts, pts_updated_at, partition_id, owner_epoch, row_version, created_at, updated_at)
VALUES
  (?, 0, NOW(6), ?, ?, 0, NOW(6), NOW(6))
ON DUPLICATE KEY UPDATE user_id = user_id
`, in.UserID, in.PartitionID, fence.OwnerEpoch); err != nil {
		return nil, storageError("init user pts state", err)
	}

	state, err := lockUserPTSState(tx, in.UserID)
	if err != nil {
		return nil, err
	}
	if state.PartitionID != in.PartitionID {
		return nil, fmt.Errorf("%w: user %d partition %d != operation partition %d", userupdates.ErrNotOwner, in.UserID, state.PartitionID, in.PartitionID)
	}

	existing, found, err := selectOperationResult(tx, in.UserID, in.OperationID)
	if err != nil {
		return nil, err
	}
	if found {
		if existing.PayloadHash != in.PayloadHash {
			return nil, userupdates.ErrOperationPayloadConflict
		}
		return &ApplyUserOperationResult{
			UserID:          in.UserID,
			OperationID:     in.OperationID,
			Pts:             existing.Pts,
			PtsCount:        existing.PtsCount,
			ResponsePayload: existing.ResponsePayload,
			ResponseHash:    existing.ResponseHash,
			AlreadyApplied:  true,
		}, nil
	}

	var op payload.MessageOperationV1
	if err := json.Unmarshal(in.Payload, &op); err != nil {
		return nil, fmt.Errorf("%w: decode message operation: %v", userupdates.ErrOperationTerminal, err)
	}
	if op.SchemaVersion != payload.MessageOperationSchemaVersion || op.OperationKind != payload.OperationKindSendMessage {
		return nil, fmt.Errorf("%w: unsupported operation schema=%d kind=%s", userupdates.ErrOperationTerminal, op.SchemaVersion, op.OperationKind)
	}
	if len(in.DependencyPts) != 0 || len(op.DependencyPts) != 0 {
		return nil, userupdates.ErrOperationTerminal
	}

	nextPTS := state.Pts + 1
	ptsCount := int32(1)
	eventPayload, eventPayloadHash, responsePayload, responsePayloadHash, err := buildEventAndResponse(in, op, nextPTS, ptsCount)
	if err != nil {
		return nil, err
	}
	if err := insertUserMessageView(tx, in, op, eventPayload); err != nil {
		return nil, err
	}
	if err := upsertUserDialog(tx, in, op, eventPayload); err != nil {
		return nil, err
	}
	if err := insertPTSEvent(tx, in, op, nextPTS, ptsCount, eventPayload, eventPayloadHash); err != nil {
		return nil, err
	}
	if err := insertPushTask(ctx, tx, r, in, op, nextPTS, eventPayload); err != nil {
		return nil, err
	}
	if err := insertOperationResult(tx, in, nextPTS, ptsCount, responsePayload, responsePayloadHash); err != nil {
		return nil, err
	}
	if _, err := tx.Exec(`
UPDATE user_pts_state
SET pts = ?,
    pts_updated_at = NOW(6),
    partition_id = ?,
    owner_epoch = ?,
    row_version = row_version + 1,
    updated_at = NOW(6)
WHERE user_id = ?
`, nextPTS, in.PartitionID, fence.OwnerEpoch, in.UserID); err != nil {
		return nil, storageError("update user pts state", err)
	}
	return &ApplyUserOperationResult{
		UserID:          in.UserID,
		OperationID:     in.OperationID,
		Pts:             nextPTS,
		PtsCount:        ptsCount,
		ResponsePayload: responsePayload,
		ResponseHash:    responsePayloadHash,
	}, nil
}

func lockPartitionFence(tx *sqlx.Tx, partitionID int32) (*partitionFenceRow, error) {
	var fence partitionFenceRow
	err := tx.QueryRowPartial(&fence, `
SELECT partition_id, owner_epoch, owner_instance_id
FROM userupdates_partition_fences
WHERE partition_id = ?
LIMIT 1 FOR UPDATE
`, partitionID)
	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return nil, userupdates.ErrNotOwner
		}
		return nil, storageError("lock partition fence", err)
	}
	return &fence, nil
}

func lockUserPTSState(tx *sqlx.Tx, userID int64) (*userPTSStateRow, error) {
	var state userPTSStateRow
	err := tx.QueryRowPartial(&state, `
SELECT user_id, pts, partition_id, owner_epoch, row_version
FROM user_pts_state
WHERE user_id = ?
LIMIT 1 FOR UPDATE
`, userID)
	if err != nil {
		return nil, storageError("lock user pts state", err)
	}
	return &state, nil
}

func buildEventAndResponse(in ApplyUserOperationInput, op payload.MessageOperationV1, pts int64, ptsCount int32) ([]byte, string, []byte, string, error) {
	event := payload.MessageEventV1{
		SchemaVersion:      payload.MessageEventSchemaVersion,
		EventKind:          payload.EventKindNewMessage,
		CanonicalMessageID: op.CanonicalMessageID,
		MessageID:          op.PeerSeq,
		PeerType:           op.PeerType,
		PeerID:             op.PeerID,
		FromUserID:         op.FromUserID,
		ToUserID:           op.ToUserID,
		Date:               op.Date,
		Out:                op.Out,
		MessageText:        op.MessageText,
		Entities:           op.Entities,
	}
	eventPayload, err := json.Marshal(event)
	if err != nil {
		return nil, "", nil, "", storageError("marshal message event", err)
	}
	response := payload.OperationResponseV1{
		SchemaVersion: payload.OperationResponseSchemaVersion,
		OperationID:   in.OperationID,
		Pts:           pts,
		PtsCount:      ptsCount,
		EventType:     payload.EventKindNewMessage,
	}
	responsePayload, err := json.Marshal(response)
	if err != nil {
		return nil, "", nil, "", storageError("marshal operation response", err)
	}
	return eventPayload, payload.HashBytes(eventPayload), responsePayload, payload.HashBytes(responsePayload), nil
}

func insertUserMessageView(tx *sqlx.Tx, in ApplyUserOperationInput, op payload.MessageOperationV1, viewPayload []byte) error {
	_, err := tx.Exec(`
INSERT INTO user_message_views
  (user_id, peer_type, peer_id, peer_seq, canonical_message_id, from_user_id, outgoing, message_kind, message_status, edit_version, date, edit_date, deleted_at, view_schema_version, view_payload, created_at, updated_at)
VALUES
  (?, ?, ?, ?, ?, ?, ?, ?, ?, 0, ?, NULL, NULL, 1, ?, NOW(6), NOW(6))
ON DUPLICATE KEY UPDATE
  message_status = VALUES(message_status),
  edit_version = VALUES(edit_version),
  edit_date = VALUES(edit_date),
  deleted_at = VALUES(deleted_at),
  view_schema_version = VALUES(view_schema_version),
  view_payload = VALUES(view_payload),
  updated_at = NOW(6)
`, in.UserID, op.PeerType, op.PeerID, op.PeerSeq, op.CanonicalMessageID, op.FromUserID, op.Out, MessageKindText, MessageStatusLive, mysqlDate(op.Date), viewPayload)
	if err != nil {
		return storageError("insert user message view", err)
	}
	return nil
}

func upsertUserDialog(tx *sqlx.Tx, in ApplyUserOperationInput, op payload.MessageOperationV1, dialogPayload []byte) error {
	unread := int32(0)
	if !op.Out {
		unread = 1
	}
	_, err := tx.Exec(`
INSERT INTO user_dialogs
  (user_id, peer_type, peer_id, top_peer_seq, top_canonical_message_id, top_message_date, unread_count, unread_mentions_count, read_inbox_max_peer_seq, read_outbox_max_peer_seq, pinned, folder_id, dialog_schema_version, dialog_payload, created_at, updated_at)
VALUES
  (?, ?, ?, ?, ?, ?, ?, 0, 0, 0, false, NULL, 1, ?, NOW(6), NOW(6))
ON DUPLICATE KEY UPDATE
  top_peer_seq = VALUES(top_peer_seq),
  top_canonical_message_id = VALUES(top_canonical_message_id),
  top_message_date = VALUES(top_message_date),
  unread_count = unread_count + VALUES(unread_count),
  unread_mentions_count = unread_mentions_count + VALUES(unread_mentions_count),
  dialog_schema_version = VALUES(dialog_schema_version),
  dialog_payload = VALUES(dialog_payload),
  updated_at = NOW(6)
`, in.UserID, op.PeerType, op.PeerID, op.PeerSeq, op.CanonicalMessageID, mysqlDate(op.Date), unread, dialogPayload)
	if err != nil {
		return storageError("upsert user dialog", err)
	}
	return nil
}

func insertPTSEvent(tx *sqlx.Tx, in ApplyUserOperationInput, op payload.MessageOperationV1, pts int64, ptsCount int32, eventPayload []byte, eventPayloadHash string) error {
	_, err := tx.Exec(`
INSERT INTO user_pts_events
  (user_id, pts, pts_count, operation_id, op_type, event_type, peer_type, peer_id, canonical_message_id, peer_seq, actor_user_id, event_schema_version, event_codec, event_payload, event_payload_hash, created_at)
VALUES
  (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, UNHEX(?), NOW(6))
`, in.UserID, pts, ptsCount, in.OperationID, in.OpType, EventTypeNewMessage, op.PeerType, op.PeerID, op.CanonicalMessageID, op.PeerSeq, op.FromUserID, payload.MessageEventSchemaVersion, PayloadCodecJSON, eventPayload, eventPayloadHash)
	if err != nil {
		return fmt.Errorf("%w: insert pts event: %w", userupdates.ErrPtsContinuityViolation, err)
	}
	return nil
}

func insertPushTask(ctx context.Context, tx *sqlx.Tx, r *Repository, in ApplyUserOperationInput, op payload.MessageOperationV1, pts int64, taskPayload []byte) error {
	taskID, err := r.idgen.NextID(ctx)
	if err != nil {
		return storageError("next push task id", err)
	}
	route := payload.RouteUser(in.UserID)
	_, err = tx.Exec(`
INSERT INTO push_task_outbox
  (task_id, user_id, pts, push_type, peer_type, peer_id, operation_id, push_partition_id, task_schema_version, task_codec, task_payload, status, publish_attempts, next_retry_at, published_topic, published_partition, published_offset, last_error_code, created_at, updated_at, published_at)
VALUES
  (?, ?, ?, ?, ?, ?, ?, ?, 1, ?, ?, ?, 0, NULL, '', 0, 0, '', NOW(6), NOW(6), NULL)
`, taskID, in.UserID, pts, PushTypeUserUpdate, op.PeerType, op.PeerID, in.OperationID, int32(route.PushPartitionID), PayloadCodecJSON, taskPayload, PushTaskStatusPending)
	if err != nil {
		return storageError("insert push task", err)
	}
	return nil
}

func insertOperationResult(tx *sqlx.Tx, in ApplyUserOperationInput, pts int64, ptsCount int32, responsePayload []byte, responseHash string) error {
	_, err := tx.Exec(`
INSERT INTO user_operation_results
  (user_id, operation_id, op_type, status, pts, pts_count, payload_hash, response_schema_version, response_codec, response_payload, response_payload_hash, terminal_error_code, created_at, updated_at)
VALUES
  (?, ?, ?, ?, ?, ?, UNHEX(?), ?, ?, ?, UNHEX(?), NULL, NOW(6), NOW(6))
`, in.UserID, in.OperationID, in.OpType, OperationResultStatusCompleted, pts, ptsCount, in.PayloadHash, payload.OperationResponseSchemaVersion, PayloadCodecJSON, responsePayload, responseHash)
	if err != nil {
		return storageError("insert operation result", err)
	}
	return nil
}

func mysqlDate(unix int32) string {
	return time.Unix(int64(unix), 0).UTC().Format("2006-01-02 15:04:05.000000")
}
