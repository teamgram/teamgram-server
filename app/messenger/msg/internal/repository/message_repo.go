package repository

import (
	"bytes"
	"context"
	"errors"
	"time"

	"github.com/teamgram/marmota/pkg/stores/sqlx"
	"github.com/teamgram/teamgram-server/v2/app/messenger/msg/internal/repository/model"
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
			if !bytes.Equal(existing.RequestPayloadHash, in.RequestPayloadHash) {
				return msg.ErrRandomIdConflict
			}
			out = existing
			return nil
		}

		state, err := selectSendStateByIDTx(tx, in.SendStateID)
		if err != nil {
			return err
		}
		if !bytes.Equal(state.RequestPayloadHash, in.RequestPayloadHash) ||
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
       c.date AS message_date,
       r.request_payload_hash
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

func selectCanonicalByIDTx(tx *sqlx.Tx, canonicalMessageID int64, requestPayloadHash []byte, sendStateID int64) (*CanonicalMessageResult, error) {
	var row canonicalResultRow
	err := tx.QueryRowPartial(&row, `
SELECT ? AS send_state_id,
       canonical_message_id,
       peer_seq,
       date AS message_date,
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
	if _, _, err := model.NewMessagePeerSequencesModel(nil).InsertIgnoreTx(tx, &model.MessagePeerSequences{
		PeerType:    peerType,
		PeerId:      peerID,
		NextPeerSeq: 1,
	}); err != nil {
		return 0, storageError("init peer sequence", err)
	}
	row, err := model.NewMessagePeerSequencesModel(nil).SelectForUpdateTx(tx, peerType, peerID)
	if err != nil {
		return 0, storageError("lock peer sequence", err)
	}
	if affected, err := model.NewMessagePeerSequencesModel(nil).UpdateNextPeerSeqTx(tx, row.NextPeerSeq+1, peerType, peerID); err != nil {
		return 0, storageError("advance peer sequence", err)
	} else if affected == 0 {
		return 0, msg.ErrSendStateConflict
	}
	return row.NextPeerSeq, nil
}

func insertCanonicalMessageTx(tx *sqlx.Tx, canonicalID int64, peerSeq int64, messageDate int32, in CreateCanonicalMessageInput) error {
	_, _, err := model.NewCanonicalMessagesModel(nil).InsertTx(tx, &model.CanonicalMessages{
		CanonicalMessageId:           canonicalID,
		PeerType:                     in.PeerType,
		PeerId:                       in.PeerID,
		PeerSeq:                      peerSeq,
		FromUserId:                   in.SenderUserID,
		MessageKind:                  MessageKindText,
		MessageText:                  in.MessageText,
		EntitiesPayloadSchemaVersion: 1,
		EntitiesPayload:              nil,
		MediaRefSchemaVersion:        1,
		MediaRefPayload:              nil,
		ServiceActionSchemaVersion:   1,
		ServiceActionPayload:         nil,
		MessageStatus:                MessageStatusLive,
		EditVersion:                  0,
		Date:                         mysqlDate(messageDate),
		EditDate:                     "",
		DeletedAt:                    "",
		StorageSchemaVersion:         1,
	})
	if err != nil {
		return storageError("insert canonical message", err)
	}
	return nil
}

func insertClientRandomTx(tx *sqlx.Tx, canonicalID int64, _ int64, _ int32, in CreateCanonicalMessageInput) error {
	_, _, err := model.NewMessageClientRandomsModel(nil).InsertTx(tx, &model.MessageClientRandoms{
		SenderUserId:       in.SenderUserID,
		PeerType:           in.PeerType,
		PeerId:             in.PeerID,
		ClientRandomId:     in.ClientRandomID,
		CanonicalMessageId: canonicalID,
		SendStateId:        in.SendStateID,
		RequestPayloadHash: in.RequestPayloadHash,
	})
	if err != nil {
		return storageError("insert client random", err)
	}
	return nil
}

type canonicalResultRow struct {
	SendStateID        int64     `db:"send_state_id"`
	CanonicalMessageID int64     `db:"canonical_message_id"`
	PeerSeq            int64     `db:"peer_seq"`
	MessageDate        time.Time `db:"message_date"`
	RequestPayloadHash []byte    `db:"request_payload_hash"`
}

func (r canonicalResultRow) toResult(created bool) *CanonicalMessageResult {
	messageDate := time.Date(r.MessageDate.Year(), r.MessageDate.Month(), r.MessageDate.Day(), r.MessageDate.Hour(), r.MessageDate.Minute(), r.MessageDate.Second(), r.MessageDate.Nanosecond(), time.UTC)
	return &CanonicalMessageResult{
		SendStateID:        r.SendStateID,
		CanonicalMessageID: r.CanonicalMessageID,
		PeerSeq:            r.PeerSeq,
		MessageDate:        int32(messageDate.Unix()),
		RequestPayloadHash: r.RequestPayloadHash,
		CreatedNew:         created,
	}
}

func mysqlDate(unix int32) string {
	return time.Unix(int64(unix), 0).UTC().Format("2006-01-02 15:04:05.000000")
}

func mysqlNow() string {
	return time.Now().UTC().Format("2006-01-02 15:04:05.000000")
}
