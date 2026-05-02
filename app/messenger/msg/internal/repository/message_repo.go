package repository

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"strings"
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
		txModels := r.models.WithTx(tx)
		existing, found, err := selectCanonicalByRandomTx(txModels, in.SenderUserID, in.PeerType, in.PeerID, in.ClientRandomID)
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

		state, err := selectSendStateByIDTx(txModels, in.SendStateID)
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
			existing, err := selectCanonicalByIDTx(txModels, state.CanonicalMessageID, state.RequestPayloadHash, state.SendStateID)
			if err != nil {
				return err
			}
			out = existing
			return nil
		}

		peerSeq, err := nextPeerSeqTx(txModels, in.PeerType, in.PeerID)
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
		if err := insertCanonicalMessageTx(txModels, canonicalID, peerSeq, messageDate, in); err != nil {
			return err
		}
		if err := insertClientRandomTx(txModels, canonicalID, peerSeq, messageDate, in); err != nil {
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

func (r *Repository) ListHistoryMessages(ctx context.Context, in ListHistoryMessagesInput) ([]HistoryMessage, error) {
	db, err := r.requireDB()
	if err != nil {
		return nil, err
	}

	limit := in.Limit
	if limit <= 0 {
		limit = 50
	} else if limit > 100 {
		limit = 100
	}

	conditions := []string{
		"`peer_type` = ?",
		"`peer_id` = ?",
		"`message_status` = ?",
	}
	args := []any{in.PeerType, in.PeerID, MessageStatusLive}
	if in.OffsetID > 0 {
		conditions = append(conditions, "`peer_seq` < ?")
		args = append(args, int64(in.OffsetID))
	}
	if in.MaxID > 0 {
		conditions = append(conditions, "`peer_seq` < ?")
		args = append(args, int64(in.MaxID))
	}
	if in.MinID > 0 {
		conditions = append(conditions, "`peer_seq` > ?")
		args = append(args, int64(in.MinID))
	}
	args = append(args, limit)

	query := fmt.Sprintf(`
SELECT
	`+"`canonical_message_id`"+`,
	`+"`peer_seq`"+`,
	`+"`from_user_id`"+`,
	`+"`peer_type`"+`,
	`+"`peer_id`"+`,
	`+"`message_kind`"+`,
	`+"`message_text`"+`,
	`+"`date`"+`
FROM `+"`canonical_messages`"+`
WHERE %s
ORDER BY `+"`peer_seq`"+` DESC
LIMIT ?`, strings.Join(conditions, " AND "))

	var rows []historyMessageRow
	if err := db.QueryRowsPartial(ctx, &rows, query, args...); err != nil {
		return nil, storageError("list history messages", err)
	}
	out := make([]HistoryMessage, 0, len(rows))
	for _, row := range rows {
		out = append(out, historyMessageRowToMessage(row))
	}
	return out, nil
}

type historyMessageRow struct {
	CanonicalMessageID int64     `db:"canonical_message_id"`
	PeerSeq            int64     `db:"peer_seq"`
	FromUserID         int64     `db:"from_user_id"`
	PeerType           int32     `db:"peer_type"`
	PeerID             int64     `db:"peer_id"`
	MessageKind        int32     `db:"message_kind"`
	MessageText        string    `db:"message_text"`
	Date               time.Time `db:"date"`
}

func historyMessageRowToMessage(r historyMessageRow) HistoryMessage {
	messageDate := time.Date(r.Date.Year(), r.Date.Month(), r.Date.Day(), r.Date.Hour(), r.Date.Minute(), r.Date.Second(), r.Date.Nanosecond(), time.UTC)
	return HistoryMessage{
		CanonicalMessageID: r.CanonicalMessageID,
		PeerSeq:            r.PeerSeq,
		FromUserID:         r.FromUserID,
		PeerType:           r.PeerType,
		PeerID:             r.PeerID,
		MessageKind:        r.MessageKind,
		MessageText:        r.MessageText,
		MessageDate:        int32(messageDate.Unix()),
	}
}

func selectCanonicalByRandomTx(txModels *model.TxModels, senderUserID int64, peerType int32, peerID int64, clientRandomID int64) (*CanonicalMessageResult, bool, error) {
	row, err := txModels.CanonicalQueries.SelectCanonicalByRandom(senderUserID, peerType, peerID, clientRandomID)
	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return nil, false, nil
		}
		return nil, false, storageError("select canonical by random", err)
	}
	return canonicalMessageRowToResult(row, false), true, nil
}

func selectCanonicalByIDTx(txModels *model.TxModels, canonicalMessageID int64, requestPayloadHash []byte, sendStateID int64) (*CanonicalMessageResult, error) {
	row, err := txModels.CanonicalQueries.SelectCanonicalByID(sendStateID, requestPayloadHash, canonicalMessageID)
	if err != nil {
		return nil, storageError("select canonical by id", err)
	}
	return canonicalMessageRowToResult(row, false), nil
}

func nextPeerSeqTx(txModels *model.TxModels, peerType int32, peerID int64) (int64, error) {
	if _, _, err := txModels.MessagePeerSequencesModel.InsertIgnore(&model.MessagePeerSequences{
		PeerType:    peerType,
		PeerId:      peerID,
		NextPeerSeq: 1,
	}); err != nil {
		return 0, storageError("init peer sequence", err)
	}
	row, err := txModels.MessagePeerSequencesModel.SelectForUpdate(peerType, peerID)
	if err != nil {
		return 0, storageError("lock peer sequence", err)
	}
	if affected, err := txModels.MessagePeerSequencesModel.UpdateNextPeerSeq(row.NextPeerSeq+1, peerType, peerID); err != nil {
		return 0, storageError("advance peer sequence", err)
	} else if affected == 0 {
		return 0, msg.ErrSendStateConflict
	}
	return row.NextPeerSeq, nil
}

func insertCanonicalMessageTx(txModels *model.TxModels, canonicalID int64, peerSeq int64, messageDate int32, in CreateCanonicalMessageInput) error {
	_, _, err := txModels.CanonicalMessagesModel.Insert(&model.CanonicalMessages{
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

func insertClientRandomTx(txModels *model.TxModels, canonicalID int64, _ int64, _ int32, in CreateCanonicalMessageInput) error {
	_, _, err := txModels.MessageClientRandomsModel.Insert(&model.MessageClientRandoms{
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

func canonicalMessageRowToResult(r *model.CanonicalMessageRow, created bool) *CanonicalMessageResult {
	if r == nil {
		return nil
	}
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
