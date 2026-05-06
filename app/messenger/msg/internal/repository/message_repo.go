package repository

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"math"
	"strings"
	"time"

	"github.com/teamgram/marmota/pkg/stores/sqlx"
	"github.com/teamgram/teamgram-server/v2/app/messenger/msg/internal/repository/model"
	"github.com/teamgram/teamgram-server/v2/app/messenger/msg/msg"
	"github.com/teamgram/teamgram-server/v2/app/messenger/userupdates/payload"
	"github.com/teamgram/teamgram-server/v2/pkg/pagination"
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

		sequencePeerID, sequenceScoped := messageSequencePeerID(in.PeerType, in.SenderUserID, in.PeerID)
		var minNextPeerSeq int64
		if sequenceScoped {
			minNextPeerSeq, err = nextConversationViewPeerSeqFloorTx(tx, in.PeerType, in.SenderUserID, in.PeerID)
			if err != nil {
				return err
			}
		}
		peerSeq, err := nextPeerSeqTx(txModels, in.PeerType, sequencePeerID, minNextPeerSeq)
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
		if err := insertCanonicalMessageTx(txModels, canonicalID, sequencePeerID, peerSeq, messageDate, in); err != nil {
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
	if _, err := r.requireDB(); err != nil {
		return nil, err
	}

	limit := pagination.NormalizeLimit(in.Limit)
	offset, err := r.historySliceOffset(ctx, in)
	if err != nil {
		return nil, err
	}

	rows, err := r.selectHistoryMessagesSlice(ctx, in, offset, limit)
	if err != nil {
		return nil, storageError("list history messages", err)
	}
	out := make([]HistoryMessage, 0, len(rows))
	for _, row := range rows {
		if in.MaxID > 0 && row.PeerSeq >= int64(in.MaxID) {
			continue
		}
		if in.MinID > 0 && row.PeerSeq <= int64(in.MinID) {
			continue
		}
		item, err := historyMessageRowToMessage(row)
		if err != nil {
			return nil, err
		}
		out = append(out, item)
	}
	return out, nil
}

func (r *Repository) SearchHashTagMessages(ctx context.Context, in SearchHashTagMessagesInput) ([]HistoryMessage, error) {
	if _, err := r.requireDB(); err != nil {
		return nil, err
	}

	tag := strings.ToLower(strings.TrimPrefix(strings.TrimSpace(in.HashTag), "#"))
	if tag == "" {
		return []HistoryMessage{}, nil
	}
	limit := pagination.NormalizeLimit(in.Limit)
	query := `
SELECT DISTINCT
	v.canonical_message_id,
	v.peer_seq,
	c.from_user_id,
	v.outgoing,
	v.peer_type,
	v.peer_id,
	v.message_kind,
	c.message_text,
	v.date AS message_date,
	v.view_payload
FROM
	user_message_views v
JOIN
	canonical_messages c
ON
	c.canonical_message_id = v.canonical_message_id
LEFT JOIN
	hash_tags h
ON
	h.user_id = v.user_id
	AND h.peer_type = v.peer_type
	AND h.peer_id = v.peer_id
	AND h.hash_tag_message_id = v.peer_seq
	AND h.hash_tag = ?
	AND h.deleted = 0
WHERE
	v.user_id = ?
	AND v.peer_type = ?
	AND v.peer_id = ?
	AND v.message_status = ?
	AND (? <= 0 OR v.peer_seq < ?)
	AND (
		h.hash_tag_message_id IS NOT NULL
		OR c.message_text LIKE ?
	)
ORDER BY
	v.peer_seq DESC
LIMIT ?`
	var rows []model.HistoryMessageRow
	likeTag := "%#" + tag + "%"
	if err := r.db.QueryRowsPartial(ctx, &rows, query, tag, in.UserID, in.PeerType, in.PeerID, MessageStatusLive, in.OffsetID, in.OffsetID, likeTag, limit); err != nil {
		return nil, storageError("search hashtag messages", err)
	}
	out := make([]HistoryMessage, 0, len(rows))
	for _, row := range rows {
		item, err := historyMessageRowToMessage(row)
		if err != nil {
			return nil, err
		}
		out = append(out, item)
	}
	return out, nil
}

func (r *Repository) GetCanonicalMessageByPeerSeq(ctx context.Context, userID int64, peerType int32, peerID int64, peerSeq int64) (*CanonicalMessage, error) {
	if _, err := r.requireDB(); err != nil {
		return nil, err
	}
	viewRow, err := r.selectCanonicalMessageByUserView(ctx, userID, peerType, peerID, peerSeq)
	if err == nil {
		return viewRow, nil
	}
	if !errors.Is(err, sqlx.ErrNotFound) {
		return nil, storageError("select canonical by user view", err)
	}

	sequencePeerID, _ := messageSequencePeerID(peerType, userID, peerID)
	row, err := r.models.CanonicalMessagesModel.SelectByPeerSeq(ctx, peerType, peerID, peerSeq)
	if err != nil && sequencePeerID != peerID {
		row, err = r.models.CanonicalMessagesModel.SelectByPeerSeq(ctx, peerType, sequencePeerID, peerSeq)
	}
	if err != nil {
		if errors.Is(err, model.ErrNotFound) {
			return nil, msg.ErrSendStateConflict
		}
		return nil, storageError("select canonical by peer seq", err)
	}
	parsedDate, err := time.ParseInLocation("2006-01-02 15:04:05.000000", row.Date, time.UTC)
	if err != nil {
		return nil, storageError("parse canonical message date", err)
	}
	return &CanonicalMessage{
		CanonicalMessageID: row.CanonicalMessageId,
		PeerSeq:            row.PeerSeq,
		FromUserID:         row.FromUserId,
		PeerType:           row.PeerType,
		PeerID:             row.PeerId,
		MessageKind:        row.MessageKind,
		MessageText:        row.MessageText,
		MessageDate:        int32(parsedDate.Unix()),
	}, nil
}

func (r *Repository) selectCanonicalMessageByUserView(ctx context.Context, userID int64, peerType int32, peerID int64, peerSeq int64) (*CanonicalMessage, error) {
	query := `
SELECT
	c.canonical_message_id,
	v.peer_seq,
	c.from_user_id,
	v.peer_type,
	v.peer_id,
	v.message_kind,
	c.message_text,
	v.date AS message_date,
	v.view_payload
FROM
	user_message_views v
JOIN
	canonical_messages c
ON
	c.canonical_message_id = v.canonical_message_id
WHERE
	v.user_id = ?
	AND v.peer_type = ?
	AND v.peer_id = ?
	AND v.peer_seq = ?
	AND v.message_status = ?
LIMIT 1`
	var row model.HistoryMessageRow
	if err := r.db.QueryRowPartial(ctx, &row, query, userID, peerType, peerID, peerSeq, MessageStatusLive); err != nil {
		return nil, err
	}
	return &CanonicalMessage{
		CanonicalMessageID: row.CanonicalMessageID,
		PeerSeq:            row.PeerSeq,
		FromUserID:         row.FromUserID,
		PeerType:           row.PeerType,
		PeerID:             row.PeerID,
		MessageKind:        row.MessageKind,
		MessageText:        row.MessageText,
		MessageDate:        int32(time.Date(row.MessageDate.Year(), row.MessageDate.Month(), row.MessageDate.Day(), row.MessageDate.Hour(), row.MessageDate.Minute(), row.MessageDate.Second(), row.MessageDate.Nanosecond(), time.UTC).Unix()),
	}, nil
}

func (r *Repository) historySliceOffset(ctx context.Context, in ListHistoryMessagesInput) (int64, error) {
	if in.OffsetID <= 0 {
		return pagination.SliceOffset(0, pagination.IDOffsetInput{
			OffsetID:  in.OffsetID,
			AddOffset: in.AddOffset,
		}), nil
	}

	query := `
SELECT
	COUNT(*)
FROM
	user_message_views v
JOIN
	user_message_views cur
ON
	cur.user_id = ?
	AND cur.peer_type = ?
	AND cur.peer_id = ?
	AND cur.peer_seq = ?
WHERE
	v.user_id = ?
	AND v.peer_type = ?
	AND v.peer_id = ?
	AND v.message_status = ?
	AND (
		v.date > cur.date
		OR (v.date = cur.date AND v.peer_seq >= cur.peer_seq)
	)`
	var offsetFromID int64
	if err := r.db.QueryRow(ctx, &offsetFromID, query, in.UserID, in.PeerType, in.PeerID, in.OffsetID, in.UserID, in.PeerType, in.PeerID, MessageStatusLive); err != nil {
		return 0, storageError("count history offset", err)
	}

	return pagination.SliceOffset(offsetFromID, pagination.IDOffsetInput{
		OffsetID:  in.OffsetID,
		AddOffset: in.AddOffset,
	}), nil
}

func (r *Repository) selectHistoryMessagesSlice(ctx context.Context, in ListHistoryMessagesInput, offset int64, limit int32) ([]model.HistoryMessageRow, error) {
	query := `
SELECT
	v.canonical_message_id,
	v.peer_seq,
	c.from_user_id,
	v.outgoing,
	v.peer_type,
	v.peer_id,
	v.message_kind,
	c.message_text,
	v.date AS message_date,
	v.view_payload
FROM
	user_message_views v
JOIN
	canonical_messages c
ON
	c.canonical_message_id = v.canonical_message_id
WHERE
	v.user_id = ?
	AND v.peer_type = ?
	AND v.peer_id = ?
	AND v.message_status = ?
ORDER BY
	v.date DESC,
	v.peer_seq DESC
LIMIT ?, ?`
	var rows []model.HistoryMessageRow
	if err := r.db.QueryRowsPartial(ctx, &rows, query, in.UserID, in.PeerType, in.PeerID, MessageStatusLive, offset, limit); err != nil {
		return nil, err
	}
	return rows, nil
}

func historyMessageRowToMessage(r model.HistoryMessageRow) (HistoryMessage, error) {
	var replyToPeerSeq int64
	if len(r.ViewPayload) > 0 {
		var event payload.MessageEventV1
		if err := json.Unmarshal(r.ViewPayload, &event); err != nil {
			return HistoryMessage{}, storageError("decode history view payload", err)
		}
		replyToPeerSeq = event.ReplyToPeerSeq
	}
	messageDate := time.Date(r.MessageDate.Year(), r.MessageDate.Month(), r.MessageDate.Day(), r.MessageDate.Hour(), r.MessageDate.Minute(), r.MessageDate.Second(), r.MessageDate.Nanosecond(), time.UTC)
	return HistoryMessage{
		CanonicalMessageID: r.CanonicalMessageID,
		PeerSeq:            r.PeerSeq,
		FromUserID:         r.FromUserID,
		Outgoing:           r.Outgoing,
		PeerType:           r.PeerType,
		PeerID:             r.PeerID,
		MessageKind:        r.MessageKind,
		MessageText:        r.MessageText,
		MessageDate:        int32(messageDate.Unix()),
		ReplyToPeerSeq:     replyToPeerSeq,
	}, nil
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

func nextPeerSeqTx(txModels *model.TxModels, peerType int32, peerID int64, minNextPeerSeq int64) (int64, error) {
	if minNextPeerSeq < 1 {
		minNextPeerSeq = 1
	}
	if _, _, err := txModels.MessagePeerSequencesModel.InsertIgnore(&model.MessagePeerSequences{
		PeerType:    peerType,
		PeerId:      peerID,
		NextPeerSeq: minNextPeerSeq,
	}); err != nil {
		return 0, storageError("init peer sequence", err)
	}
	row, err := txModels.MessagePeerSequencesModel.SelectForUpdate(peerType, peerID)
	if err != nil {
		return 0, storageError("lock peer sequence", err)
	}
	peerSeq := row.NextPeerSeq
	if peerSeq < minNextPeerSeq {
		peerSeq = minNextPeerSeq
	}
	if affected, err := txModels.MessagePeerSequencesModel.UpdateNextPeerSeq(peerSeq+1, peerType, peerID); err != nil {
		return 0, storageError("advance peer sequence", err)
	} else if affected == 0 {
		return 0, msg.ErrSendStateConflict
	}
	return peerSeq, nil
}

func insertCanonicalMessageTx(txModels *model.TxModels, canonicalID int64, canonicalPeerID int64, peerSeq int64, messageDate int32, in CreateCanonicalMessageInput) error {
	_, _, err := txModels.CanonicalMessagesModel.Insert(&model.CanonicalMessages{
		CanonicalMessageId:           canonicalID,
		PeerType:                     in.PeerType,
		PeerId:                       canonicalPeerID,
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

func nextConversationViewPeerSeqFloorTx(tx *sqlx.Tx, peerType int32, userID int64, peerID int64) (int64, error) {
	var maxPeerSeq int64
	query := `
SELECT
	COALESCE(MAX(peer_seq), 0)
FROM
	user_message_views
WHERE
	peer_type = ?
	AND (
		(user_id = ? AND peer_id = ?)
		OR (user_id = ? AND peer_id = ?)
	)`
	if err := tx.QueryRow(&maxPeerSeq, query, peerType, userID, peerID, peerID, userID); err != nil {
		return 0, storageError("select conversation view peer seq floor", err)
	}
	return maxPeerSeq + 1, nil
}

func messageSequencePeerID(peerType int32, userID int64, peerID int64) (int64, bool) {
	if peerType != payload.PeerTypeUser || userID <= 0 || peerID <= 0 || userID == peerID {
		return peerID, false
	}
	if userID > math.MaxUint32 || peerID > math.MaxUint32 {
		return peerID, false
	}
	lo, hi := userID, peerID
	if lo > hi {
		lo, hi = hi, lo
	}
	return int64((uint64(lo) << 32) | uint64(hi)), true
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
