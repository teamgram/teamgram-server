package repository

import (
	"context"
	"errors"

	"github.com/teamgram/marmota/pkg/stores/sqlx"
	"github.com/teamgram/teamgram-server/v2/app/messenger/userupdates/userupdates"
)

func (r *Repository) GetOperationResult(ctx context.Context, userID int64, operationID string) (*OperationResult, error) {
	db, err := r.requireDB()
	if err != nil {
		return nil, err
	}
	var row operationResultRow
	err = db.QueryRowPartial(ctx, &row, `
SELECT user_id, operation_id, op_type, status, IFNULL(pts, 0) AS pts, IFNULL(pts_count, 0) AS pts_count,
       LOWER(HEX(payload_hash)) AS payload_hash,
       response_payload,
       LOWER(HEX(response_payload_hash)) AS response_hash,
       IFNULL(terminal_error_code, '') AS terminal_error_code
FROM user_operation_results
WHERE user_id = ? AND operation_id = ?
LIMIT 1
`, userID, operationID)
	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return nil, userupdates.ErrOperationTerminal
		}
		return nil, storageError("get operation result", err)
	}
	return row.toOperationResult(), nil
}

func (r *Repository) GetState(ctx context.Context, userID int64) (*UserState, error) {
	db, err := r.requireDB()
	if err != nil {
		return nil, err
	}
	var row userPTSStateRow
	err = db.QueryRowPartial(ctx, &row, `
SELECT user_id, pts, partition_id, owner_epoch, row_version
FROM user_pts_state
WHERE user_id = ?
LIMIT 1
`, userID)
	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return &UserState{UserID: userID}, nil
		}
		return nil, storageError("get state", err)
	}
	return &UserState{
		UserID:      row.UserID,
		Pts:         row.Pts,
		PartitionID: row.PartitionID,
		OwnerEpoch:  row.OwnerEpoch,
		RowVersion:  row.RowVersion,
	}, nil
}

func (r *Repository) GetDifference(ctx context.Context, in GetDifferenceInput) (*GetDifferenceResult, error) {
	db, err := r.requireDB()
	if err != nil {
		return nil, err
	}
	limit := in.Limit
	if limit <= 0 {
		limit = 100
	}
	var rows []userEventRow
	err = db.QueryRowsPartial(ctx, &rows, `
SELECT user_id, pts, pts_count, operation_id, op_type, event_type, peer_type, peer_id,
       IFNULL(canonical_message_id, 0) AS canonical_message_id,
       IFNULL(peer_seq, 0) AS peer_seq,
       IFNULL(actor_user_id, 0) AS actor_user_id,
       event_schema_version,
       event_codec,
       event_payload,
       LOWER(HEX(event_payload_hash)) AS event_payload_hash
FROM user_pts_events
WHERE user_id = ? AND pts > ?
ORDER BY pts
LIMIT ?
`, in.UserID, in.Pts, limit)
	if err != nil && !errors.Is(err, sqlx.ErrNotFound) {
		return nil, storageError("get difference events", err)
	}
	state, err := r.GetState(ctx, in.UserID)
	if err != nil {
		return nil, err
	}
	events := make([]UserEvent, 0, len(rows))
	for _, row := range rows {
		events = append(events, row.toUserEvent())
	}
	return &GetDifferenceResult{State: *state, Events: events}, nil
}

type operationResultRow struct {
	UserID            int64  `db:"user_id"`
	OperationID       string `db:"operation_id"`
	OpType            int32  `db:"op_type"`
	Status            int32  `db:"status"`
	Pts               int64  `db:"pts"`
	PtsCount          int32  `db:"pts_count"`
	PayloadHash       string `db:"payload_hash"`
	ResponsePayload   []byte `db:"response_payload"`
	ResponseHash      string `db:"response_hash"`
	TerminalErrorCode string `db:"terminal_error_code"`
}

func (r operationResultRow) toOperationResult() *OperationResult {
	return &OperationResult{
		UserID:            r.UserID,
		OperationID:       r.OperationID,
		OpType:            r.OpType,
		Status:            r.Status,
		Pts:               r.Pts,
		PtsCount:          r.PtsCount,
		PayloadHash:       r.PayloadHash,
		ResponsePayload:   r.ResponsePayload,
		ResponseHash:      r.ResponseHash,
		TerminalErrorCode: r.TerminalErrorCode,
	}
}

type userEventRow struct {
	UserID             int64  `db:"user_id"`
	Pts                int64  `db:"pts"`
	PtsCount           int32  `db:"pts_count"`
	OperationID        string `db:"operation_id"`
	OpType             int32  `db:"op_type"`
	EventType          int32  `db:"event_type"`
	PeerType           int32  `db:"peer_type"`
	PeerID             int64  `db:"peer_id"`
	CanonicalMessageID int64  `db:"canonical_message_id"`
	PeerSeq            int64  `db:"peer_seq"`
	ActorUserID        int64  `db:"actor_user_id"`
	EventSchemaVersion int32  `db:"event_schema_version"`
	EventCodec         int32  `db:"event_codec"`
	EventPayload       []byte `db:"event_payload"`
	EventPayloadHash   string `db:"event_payload_hash"`
}

func (r userEventRow) toUserEvent() UserEvent {
	return UserEvent{
		UserID:             r.UserID,
		Pts:                r.Pts,
		PtsCount:           r.PtsCount,
		OperationID:        r.OperationID,
		OpType:             r.OpType,
		EventType:          r.EventType,
		PeerType:           r.PeerType,
		PeerID:             r.PeerID,
		CanonicalMessageID: r.CanonicalMessageID,
		PeerSeq:            r.PeerSeq,
		ActorUserID:        r.ActorUserID,
		EventSchemaVersion: r.EventSchemaVersion,
		EventCodec:         r.EventCodec,
		EventPayload:       r.EventPayload,
		EventPayloadHash:   r.EventPayloadHash,
	}
}

func selectOperationResult(tx *sqlx.Tx, userID int64, operationID string) (*OperationResult, bool, error) {
	var row operationResultRow
	err := tx.QueryRowPartial(&row, `
SELECT user_id, operation_id, op_type, status, IFNULL(pts, 0) AS pts, IFNULL(pts_count, 0) AS pts_count,
       LOWER(HEX(payload_hash)) AS payload_hash,
       response_payload,
       LOWER(HEX(response_payload_hash)) AS response_hash,
       IFNULL(terminal_error_code, '') AS terminal_error_code
FROM user_operation_results
WHERE user_id = ? AND operation_id = ?
LIMIT 1
`, userID, operationID)
	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return nil, false, nil
		}
		return nil, false, storageError("select operation result", err)
	}
	return row.toOperationResult(), true, nil
}
