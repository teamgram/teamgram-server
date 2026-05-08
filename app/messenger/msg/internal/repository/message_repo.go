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
				retry, err := existingCanonicalMatchesRetryTx(txModels, existing.CanonicalMessageID, in)
				if err != nil {
					return err
				}
				if !retry {
					return msg.ErrRandomIdConflict
				}
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
			retry := false
			if state.SenderUserID == in.SenderUserID &&
				state.PeerType == in.PeerType &&
				state.PeerID == in.PeerID &&
				state.ClientRandomID == in.ClientRandomID {
				retry, err = existingCanonicalMatchesRetryTx(txModels, state.CanonicalMessageID, in)
				if err != nil {
					return err
				}
			}
			if !retry {
				return msg.ErrRandomIdConflict
			}
		}
		if state.CanonicalMessageID != 0 {
			existing, err := selectCanonicalByIDTx(txModels, state.CanonicalMessageID, state.RequestPayloadHash, state.SendStateID)
			if err != nil && !bytes.Equal(state.RequestPayloadHash, in.RequestPayloadHash) {
				existing, err = selectCanonicalByIDOnlyTx(txModels, state.CanonicalMessageID, state.SendStateID)
			}
			if err != nil {
				return err
			}
			out = existing
			return nil
		}

		sequencePeerID, sequenceScoped := messageSequencePeerID(in.PeerType, in.SenderUserID, in.PeerID)
		var minNextPeerSeq int64
		if sequenceScoped {
			minNextPeerSeq, err = nextConversationViewPeerSeqFloorTx(txModels, in.PeerType, in.SenderUserID, in.PeerID)
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
		messageDate := unixOrNow(in.MessageDate)
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
	bounds, err := r.ResolveHistoryCursorIDs(ctx, in.UserID, in.PeerType, in.PeerID, in.OffsetID, in.MaxID, in.MinID)
	if err != nil {
		return nil, err
	}

	limit := pagination.NormalizeLimit(in.Limit)
	offset, err := r.historySliceOffset(ctx, in, bounds.OffsetPeerSeq)
	if err != nil {
		return nil, err
	}

	rows, err := r.selectHistoryMessagesSlice(ctx, in, offset, limit)
	if err != nil {
		return nil, storageError("list history messages", err)
	}
	out := make([]HistoryMessage, 0, len(rows))
	for _, row := range rows {
		if !historyMessageWithinBounds(row.PeerSeq, bounds) {
			continue
		}
		item, err := r.historyMessageRowToMessage(ctx, in.UserID, in.PeerType, in.PeerID, row)
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
	likeTag := "%#" + tag + "%"
	rows, err := r.models.CanonicalQueries.SearchHashTagMessages(ctx, tag, in.UserID, in.PeerType, in.PeerID, MessageStatusLive, int64(in.OffsetID), int64(in.OffsetID), likeTag, limit)
	if err != nil {
		return nil, storageError("search hashtag messages", err)
	}
	out := make([]HistoryMessage, 0, len(rows))
	for _, row := range rows {
		item, err := r.historyMessageRowToMessage(ctx, in.UserID, in.PeerType, in.PeerID, row)
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
	return &CanonicalMessage{
		CanonicalMessageID: row.CanonicalMessageId,
		PeerSeq:            row.PeerSeq,
		FromUserID:         row.FromUserId,
		PeerType:           row.PeerType,
		PeerID:             row.PeerId,
		MessageKind:        row.MessageKind,
		MessageText:        row.MessageText,
		MessageDate:        row.Date,
	}, nil
}

func (r *Repository) EditCanonicalMessage(ctx context.Context, in EditCanonicalMessageInput) (*EditMessageResult, error) {
	db, err := r.requireDB()
	if err != nil {
		return nil, err
	}
	if in.ActorUserID <= 0 || in.PeerSeq <= 0 || in.NewMessageText == "" {
		return nil, msg.ErrSendStateConflict
	}

	var out *EditMessageResult
	err = db.Transact(ctx, func(tx *sqlx.Tx) error {
		txModels := r.models.WithTx(tx)
		row, err := txModels.CanonicalQueries.SelectEditableMessageForUpdate(in.ActorUserID, in.PeerType, in.PeerID, in.PeerSeq, MessageStatusLive)
		if err != nil {
			if errors.Is(err, sqlx.ErrNotFound) {
				return msg.ErrSendStateConflict
			}
			return storageError("select editable message", err)
		}
		if row.FromUserID != in.ActorUserID {
			return msg.ErrMessageAuthorRequired
		}
		if row.MessageText == in.NewMessageText {
			if row.EditVersion <= 0 {
				return msg.ErrMessageNotModified
			}
			out = &EditMessageResult{
				CanonicalMessageID: row.CanonicalMessageID,
				PeerSeq:            row.PeerSeq,
				FromUserID:         row.FromUserID,
				PeerType:           row.PeerType,
				PeerID:             row.PeerID,
				MessageKind:        row.MessageKind,
				MessageText:        row.MessageText,
				MessageDate:        row.MessageDate,
				EditDate:           row.EditDate,
				EditVersion:        row.EditVersion,
			}
			return nil
		}

		editDate := unixOrNow(in.RequestEditDate)
		editVersion := row.EditVersion + 1
		affected, err := txModels.CanonicalMessagesModel.UpdateMessageEdit(in.NewMessageText, editVersion, editDate, row.CanonicalMessageID, row.EditVersion)
		if err != nil {
			return storageError("update canonical message edit", err)
		}
		if affected == 0 {
			return msg.ErrSendStateConflict
		}
		out = &EditMessageResult{
			CanonicalMessageID: row.CanonicalMessageID,
			PeerSeq:            row.PeerSeq,
			FromUserID:         row.FromUserID,
			PeerType:           row.PeerType,
			PeerID:             row.PeerID,
			MessageKind:        row.MessageKind,
			MessageText:        in.NewMessageText,
			MessageDate:        row.MessageDate,
			EditDate:           editDate,
			EditVersion:        editVersion,
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (r *Repository) selectCanonicalMessageByUserView(ctx context.Context, userID int64, peerType int32, peerID int64, peerSeq int64) (*CanonicalMessage, error) {
	row, err := r.models.CanonicalQueries.SelectCanonicalByUserView(ctx, userID, peerType, peerID, peerSeq, MessageStatusLive)
	if err != nil {
		return nil, err
	}
	return historyRowToCanonicalMessage(*row), nil
}

func historyRowToCanonicalMessage(row model.HistoryMessageRow) *CanonicalMessage {
	return &CanonicalMessage{
		CanonicalMessageID: row.CanonicalMessageID,
		PeerSeq:            row.PeerSeq,
		FromUserID:         row.FromUserID,
		PeerType:           row.PeerType,
		PeerID:             row.PeerID,
		MessageKind:        row.MessageKind,
		MessageText:        row.MessageText,
		MessageDate:        row.MessageDate,
	}
}

func unixNow() int64 {
	return time.Now().UTC().Unix()
}

func unixOrNow(seconds int64) int64 {
	if seconds > 0 {
		return seconds
	}
	return unixNow()
}

func (r *Repository) historySliceOffset(ctx context.Context, in ListHistoryMessagesInput, offsetPeerSeq int64) (int64, error) {
	if offsetPeerSeq <= 0 {
		return pagination.SliceOffset(0, pagination.IDOffsetInput{
			OffsetID:  historyOffsetMarker(offsetPeerSeq).OffsetID,
			AddOffset: in.AddOffset,
		}), nil
	}

	row, err := r.models.CanonicalQueries.CountHistoryOffset(ctx, in.UserID, int64(in.PeerType), in.PeerID, offsetPeerSeq, in.UserID, in.PeerType, in.PeerID, MessageStatusLive)
	if err != nil {
		return 0, storageError("count history offset", err)
	}

	return pagination.SliceOffset(row.OffsetFromID, pagination.IDOffsetInput{
		OffsetID:  historyOffsetMarker(offsetPeerSeq).OffsetID,
		AddOffset: in.AddOffset,
	}), nil
}

func (r *Repository) selectHistoryMessagesSlice(ctx context.Context, in ListHistoryMessagesInput, offset int64, limit int32) ([]model.HistoryMessageRow, error) {
	return r.models.CanonicalQueries.SelectHistoryMessagesPage(ctx, in.UserID, in.PeerType, in.PeerID, MessageStatusLive, offset, limit)
}

func (r *Repository) historyMessageRowToMessage(ctx context.Context, userID int64, peerType int32, peerID int64, row model.HistoryMessageRow) (HistoryMessage, error) {
	return historyMessageRowToMessage(row, func(replyPeerSeq int64) (int64, error) {
		return r.ResolvePeerSeqToUserMessageID(ctx, userID, peerType, peerID, replyPeerSeq)
	})
}

type replyPublicIDResolver func(peerSeq int64) (int64, error)

func historyMessageRowToMessage(r model.HistoryMessageRow, resolveReply replyPublicIDResolver) (HistoryMessage, error) {
	var replyToPeerSeq int64
	var replyToUserMessageID int64
	if len(r.ViewPayload) > 0 {
		var envelope struct {
			SchemaVersion int `json:"schema_version"`
		}
		if err := json.Unmarshal(r.ViewPayload, &envelope); err != nil {
			return HistoryMessage{}, storageError("decode history view payload", err)
		}
		switch envelope.SchemaVersion {
		case payload.MessageEventSchemaVersion:
			var event payload.MessageEventV2
			if err := json.Unmarshal(r.ViewPayload, &event); err != nil {
				return HistoryMessage{}, storageError("decode history view payload v2", err)
			}
			replyToUserMessageID = event.ReplyToUserMessageID
		default:
			var event payload.MessageEventV1
			if err := json.Unmarshal(r.ViewPayload, &event); err != nil {
				return HistoryMessage{}, storageError("decode history view payload v1", err)
			}
			replyToPeerSeq = event.ReplyToPeerSeq
			if replyToPeerSeq > 0 && resolveReply != nil {
				var err error
				replyToUserMessageID, err = resolveReply(replyToPeerSeq)
				if err != nil {
					return HistoryMessage{}, err
				}
			}
		}
	}
	return HistoryMessage{
		CanonicalMessageID:   r.CanonicalMessageID,
		PeerSeq:              r.PeerSeq,
		UserMessageID:        r.UserMessageID,
		FromUserID:           r.FromUserID,
		Outgoing:             r.Outgoing,
		PeerType:             r.PeerType,
		PeerID:               r.PeerID,
		MessageKind:          r.MessageKind,
		MessageText:          r.MessageText,
		MessageDate:          r.MessageDate,
		ReplyToPeerSeq:       replyToPeerSeq,
		ReplyToUserMessageID: replyToUserMessageID,
	}, nil
}

func historyMessageWithinBounds(peerSeq int64, bounds HistoryCursorBounds) bool {
	if bounds.MaxPeerSeq > 0 && peerSeq >= bounds.MaxPeerSeq {
		return false
	}
	if bounds.MinPeerSeq > 0 && peerSeq <= bounds.MinPeerSeq {
		return false
	}
	return true
}

func historyOffsetMarker(offsetPeerSeq int64) pagination.IDOffsetInput {
	if offsetPeerSeq <= 0 {
		return pagination.IDOffsetInput{}
	}
	return pagination.IDOffsetInput{OffsetID: 1}
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

func selectCanonicalByIDOnlyTx(txModels *model.TxModels, canonicalMessageID int64, sendStateID int64) (*CanonicalMessageResult, error) {
	row, err := txModels.CanonicalMessagesModel.SelectByCanonicalMessageId(canonicalMessageID)
	if err != nil {
		return nil, storageError("select canonical by id", err)
	}
	return &CanonicalMessageResult{
		SendStateID:        sendStateID,
		CanonicalMessageID: row.CanonicalMessageId,
		PeerSeq:            row.PeerSeq,
		MessageDate:        row.Date,
		CreatedNew:         false,
	}, nil
}

func existingCanonicalMatchesRetryTx(txModels *model.TxModels, canonicalMessageID int64, in CreateCanonicalMessageInput) (bool, error) {
	if canonicalMessageID == 0 {
		return false, nil
	}
	row, err := txModels.CanonicalMessagesModel.SelectByCanonicalMessageId(canonicalMessageID)
	if err != nil {
		if errors.Is(err, model.ErrNotFound) {
			return false, nil
		}
		return false, storageError("select canonical for random retry", err)
	}
	return row.MessageText == in.MessageText &&
		row.FromUserId == in.SenderUserID &&
		row.MessageKind == MessageKindText, nil
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

func insertCanonicalMessageTx(txModels *model.TxModels, canonicalID int64, canonicalPeerID int64, peerSeq int64, messageDate int64, in CreateCanonicalMessageInput) error {
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
		Date:                         messageDate,
		EditDate:                     0,
		DeletedAt:                    0,
		StorageSchemaVersion:         1,
	})
	if err != nil {
		return storageError("insert canonical message", err)
	}
	return nil
}

func nextConversationViewPeerSeqFloorTx(txModels *model.TxModels, peerType int32, userID int64, peerID int64) (int64, error) {
	row, err := txModels.CanonicalQueries.SelectConversationViewPeerSeqFloor(peerType, userID, peerID, peerID, userID)
	if err != nil {
		return 0, storageError("select conversation view peer seq floor", err)
	}
	return row.PeerSeqFloor, nil
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

func insertClientRandomTx(txModels *model.TxModels, canonicalID int64, _ int64, _ int64, in CreateCanonicalMessageInput) error {
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
	return &CanonicalMessageResult{
		SendStateID:        r.SendStateID,
		CanonicalMessageID: r.CanonicalMessageID,
		PeerSeq:            r.PeerSeq,
		MessageDate:        r.MessageDate,
		RequestPayloadHash: r.RequestPayloadHash,
		CreatedNew:         created,
	}
}
