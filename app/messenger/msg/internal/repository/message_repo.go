package repository

import (
	"context"
	"errors"
	"time"

	"github.com/teamgram/marmota/pkg/stores/sqlx"
	"github.com/teamgram/teamgram-server/v2/app/messenger/msg/msg"
)

func (r *Repository) CreateOrGetByClientRandom(ctx context.Context, in CreateCanonicalMessageInput) (*CanonicalMessageResult, error) {
	db, err := r.requireDB()
	if err != nil {
		return nil, err
	}
	var out *CanonicalMessageResult
	err = db.Transact(ctx, func(tx *sqlx.Tx) error {
		existing, found, err := selectCanonicalByRandomTx(tx, in.SenderUserID, in.PeerType, in.PeerID, in.ClientRandomID)
		if err != nil {
			return err
		}
		if found {
			if existing.RequestPayloadHash != in.RequestPayloadHash {
				return msg.ErrRandomIdConflict
			}
			out = existing
			return nil
		}

		state, err := selectSendStateByIDTx(tx, in.SendStateID)
		if err != nil {
			return err
		}
		if state.RequestPayloadHash != in.RequestPayloadHash ||
			state.SenderUserID != in.SenderUserID ||
			state.PeerType != in.PeerType ||
			state.PeerID != in.PeerID ||
			state.ClientRandomID != in.ClientRandomID {
			return msg.ErrRandomIdConflict
		}
		if state.CanonicalMessageID != 0 {
			existing, err := selectCanonicalByIDTx(tx, state.CanonicalMessageID, state.RequestPayloadHash, state.SendStateID)
			if err != nil {
				return err
			}
			out = existing
			return nil
		}

		peerSeq, err := nextPeerSeqTx(tx, in.PeerType, in.PeerID)
		if err != nil {
			return err
		}
		canonicalID, err := r.nextID(ctx, "next canonical message id")
		if err != nil {
			return err
		}
		messageDate := in.MessageDate
		if messageDate == 0 {
			messageDate = int32(time.Now().Unix())
		}
		if err := insertCanonicalMessageTx(tx, canonicalID, peerSeq, messageDate, in); err != nil {
			return err
		}
		if err := insertClientRandomTx(tx, canonicalID, peerSeq, messageDate, in); err != nil {
			return err
		}
		out = &CanonicalMessageResult{
			SendStateID:        in.SendStateID,
			CanonicalMessageID: canonicalID,
			PeerSeq:            peerSeq,
			MessageDate:        messageDate,
			RequestPayloadHash: in.RequestPayloadHash,
			CreatedNew:         true,
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return out, nil
}

func selectCanonicalByRandomTx(tx *sqlx.Tx, senderUserID int64, peerType int32, peerID int64, clientRandomID int64) (*CanonicalMessageResult, bool, error) {
	var row canonicalResultRow
	err := tx.QueryRowPartial(&row, `
SELECT r.send_state_id,
       r.canonical_message_id,
       c.peer_seq,
       TIMESTAMPDIFF(SECOND, '1970-01-01 00:00:00', c.date) AS message_date,
       LOWER(HEX(r.request_payload_hash)) AS request_payload_hash
FROM message_client_randoms r
JOIN canonical_messages c ON c.canonical_message_id = r.canonical_message_id
WHERE r.sender_user_id = ?
  AND r.peer_type = ?
  AND r.peer_id = ?
  AND r.client_random_id = ?
LIMIT 1
`, senderUserID, peerType, peerID, clientRandomID)
	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return nil, false, nil
		}
		return nil, false, storageError("select canonical by random", err)
	}
	return row.toResult(false), true, nil
}

func selectCanonicalByIDTx(tx *sqlx.Tx, canonicalMessageID int64, requestPayloadHash string, sendStateID int64) (*CanonicalMessageResult, error) {
	var row canonicalResultRow
	err := tx.QueryRowPartial(&row, `
SELECT ? AS send_state_id,
       canonical_message_id,
       peer_seq,
       TIMESTAMPDIFF(SECOND, '1970-01-01 00:00:00', date) AS message_date,
       ? AS request_payload_hash
FROM canonical_messages
WHERE canonical_message_id = ?
LIMIT 1
`, sendStateID, requestPayloadHash, canonicalMessageID)
	if err != nil {
		return nil, storageError("select canonical by id", err)
	}
	return row.toResult(false), nil
}

func nextPeerSeqTx(tx *sqlx.Tx, peerType int32, peerID int64) (int64, error) {
	if _, err := tx.Exec(`
INSERT INTO message_peer_sequences (peer_type, peer_id, next_peer_seq, created_at, updated_at)
VALUES (?, ?, 1, NOW(6), NOW(6))
ON DUPLICATE KEY UPDATE peer_type = peer_type
`, peerType, peerID); err != nil {
		return 0, storageError("init peer sequence", err)
	}
	var row struct {
		NextPeerSeq int64 `db:"next_peer_seq"`
	}
	if err := tx.QueryRowPartial(&row, `
SELECT next_peer_seq
FROM message_peer_sequences
WHERE peer_type = ? AND peer_id = ?
LIMIT 1 FOR UPDATE
`, peerType, peerID); err != nil {
		return 0, storageError("lock peer sequence", err)
	}
	if _, err := tx.Exec(`
UPDATE message_peer_sequences
SET next_peer_seq = next_peer_seq + 1,
    updated_at = NOW(6)
WHERE peer_type = ? AND peer_id = ?
`, peerType, peerID); err != nil {
		return 0, storageError("advance peer sequence", err)
	}
	return row.NextPeerSeq, nil
}

func insertCanonicalMessageTx(tx *sqlx.Tx, canonicalID int64, peerSeq int64, messageDate int32, in CreateCanonicalMessageInput) error {
	_, err := tx.Exec(`
INSERT INTO canonical_messages
  (canonical_message_id, peer_type, peer_id, peer_seq, from_user_id, message_kind, message_text,
   entities_payload_schema_version, entities_payload, media_ref_schema_version, media_ref_payload,
   service_action_schema_version, service_action_payload, message_status, edit_version, date,
   edit_date, deleted_at, storage_schema_version, created_at, updated_at)
VALUES
  (?, ?, ?, ?, ?, ?, ?, 1, NULL, 1, NULL, 1, NULL, ?, 0, ?, NULL, NULL, 1, NOW(6), NOW(6))
`, canonicalID, in.PeerType, in.PeerID, peerSeq, in.SenderUserID, MessageKindText, in.MessageText, MessageStatusLive, mysqlDate(messageDate))
	if err != nil {
		return storageError("insert canonical message", err)
	}
	return nil
}

func insertClientRandomTx(tx *sqlx.Tx, canonicalID int64, _ int64, _ int32, in CreateCanonicalMessageInput) error {
	_, err := tx.Exec(`
INSERT INTO message_client_randoms
  (sender_user_id, peer_type, peer_id, client_random_id, canonical_message_id, send_state_id, request_payload_hash, created_at, updated_at)
VALUES
  (?, ?, ?, ?, ?, ?, UNHEX(?), NOW(6), NOW(6))
`, in.SenderUserID, in.PeerType, in.PeerID, in.ClientRandomID, canonicalID, in.SendStateID, in.RequestPayloadHash)
	if err != nil {
		return storageError("insert client random", err)
	}
	return nil
}

type canonicalResultRow struct {
	SendStateID        int64  `db:"send_state_id"`
	CanonicalMessageID int64  `db:"canonical_message_id"`
	PeerSeq            int64  `db:"peer_seq"`
	MessageDate        int64  `db:"message_date"`
	RequestPayloadHash string `db:"request_payload_hash"`
}

func (r canonicalResultRow) toResult(created bool) *CanonicalMessageResult {
	return &CanonicalMessageResult{
		SendStateID:        r.SendStateID,
		CanonicalMessageID: r.CanonicalMessageID,
		PeerSeq:            r.PeerSeq,
		MessageDate:        int32(r.MessageDate),
		RequestPayloadHash: r.RequestPayloadHash,
		CreatedNew:         created,
	}
}

func mysqlDate(unix int32) string {
	return time.Unix(int64(unix), 0).UTC().Format("2006-01-02 15:04:05.000000")
}
