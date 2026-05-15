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
			SendStateID:                  in.SendStateID,
			CanonicalMessageID:           canonicalID,
			PeerSeq:                      peerSeq,
			MessageDate:                  messageDate,
			RequestPayloadHash:           in.RequestPayloadHash,
			EntitiesPayloadSchemaVersion: canonicalEntitiesSchemaVersion(in.EntitiesPayloadSchemaVersion),
			EntitiesPayload:              in.EntitiesPayload,
			MediaRefSchemaVersion:        canonicalMediaSchemaVersion(in.MediaRefSchemaVersion),
			MediaRefPayload:              in.MediaRefPayload,
			MessageAttrsSchemaVersion:    in.MessageAttrsSchemaVersion,
			MessageAttrsPayload:          in.MessageAttrsPayload,
			ForwardRefSchemaVersion:      in.ForwardRefSchemaVersion,
			ForwardRefPayload:            in.ForwardRefPayload,
			ServiceActionSchemaVersion:   in.ServiceActionSchemaVersion,
			ServiceActionPayload:         in.ServiceActionPayload,
			CreatedNew:                   true,
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (r *Repository) CreateOrGetCanonicalBatchByClientRandom(ctx context.Context, in CreateCanonicalBatchInput) (*CanonicalBatchResult, error) {
	db, err := r.requireDB()
	if err != nil {
		return nil, err
	}
	if len(in.Items) == 0 {
		return &CanonicalBatchResult{}, nil
	}

	var out *CanonicalBatchResult
	err = db.Transact(ctx, func(tx *sqlx.Tx) error {
		txModels := r.models.WithTx(tx)
		results := make([]CanonicalMessageResult, len(in.Items))
		pending := make([]int, 0, len(in.Items))
		seenRandoms := make(map[int64]struct{}, len(in.Items))

		for i, item := range in.Items {
			if _, ok := seenRandoms[item.ClientRandomID]; ok {
				return msg.ErrRandomIdConflict
			}
			seenRandoms[item.ClientRandomID] = struct{}{}

			canonicalInput := batchItemToCanonicalInput(in, item, 0)
			existing, found, err := selectCanonicalByRandomTx(txModels, in.SenderUserID, in.PeerType, in.PeerID, item.ClientRandomID)
			if err != nil {
				return err
			}
			if found {
				if !bytes.Equal(existing.RequestPayloadHash, item.RequestPayloadHash) {
					return msg.ErrRandomIdConflict
				}
				matches, err := existingCanonicalMatchesRetryTx(txModels, existing.CanonicalMessageID, canonicalInput)
				if err != nil {
					return err
				}
				if !matches {
					return msg.ErrRandomIdConflict
				}
				state, found, err := selectSendStateByRandomTx(txModels, in.SenderUserID, in.PeerType, in.PeerID, item.ClientRandomID)
				if err != nil {
					return err
				}
				if found {
					attachSendStateToCanonicalResult(existing, state)
				}
				results[i] = *existing
				continue
			}

			state, found, err := selectSendStateByRandomTx(txModels, in.SenderUserID, in.PeerType, in.PeerID, item.ClientRandomID)
			if err != nil {
				return err
			}
			if !found {
				state, err = r.insertSendStateTx(ctx, txModels, CreateSendStateInput{
					SenderUserID:                in.SenderUserID,
					PeerType:                    in.PeerType,
					PeerID:                      in.PeerID,
					ClientRandomID:              item.ClientRandomID,
					RequestPayloadSchemaVersion: item.RequestPayloadSchemaVersion,
					RequestPayloadHash:          item.RequestPayloadHash,
					MessageText:                 item.MessageText,
				})
				if err != nil {
					return err
				}
			} else if state.SenderUserID != in.SenderUserID ||
				state.PeerType != in.PeerType ||
				state.PeerID != in.PeerID ||
				state.ClientRandomID != item.ClientRandomID ||
				!bytes.Equal(state.RequestPayloadHash, item.RequestPayloadHash) {
				return msg.ErrRandomIdConflict
			}

			canonicalInput.SendStateID = state.SendStateID
			if state.CanonicalMessageID != 0 {
				existing, err := selectCanonicalByIDTx(txModels, state.CanonicalMessageID, item.RequestPayloadHash, state.SendStateID)
				if err != nil {
					return err
				}
				matches, err := existingCanonicalMatchesRetryTx(txModels, existing.CanonicalMessageID, canonicalInput)
				if err != nil {
					return err
				}
				if !matches {
					return msg.ErrRandomIdConflict
				}
				attachSendStateToCanonicalResult(existing, state)
				results[i] = *existing
				continue
			}

			results[i] = CanonicalMessageResult{
				SendStateID:        state.SendStateID,
				RequestPayloadHash: item.RequestPayloadHash,
				SendStateStatus:    state.Status,
			}
			pending = append(pending, i)
		}

		if len(pending) > 0 {
			sequencePeerID, sequenceScoped := messageSequencePeerID(in.PeerType, in.SenderUserID, in.PeerID)
			var minNextPeerSeq int64
			if sequenceScoped {
				minNextPeerSeq, err = nextConversationViewPeerSeqFloorTx(txModels, in.PeerType, in.SenderUserID, in.PeerID)
				if err != nil {
					return err
				}
			}
			firstPeerSeq, err := nextPeerSeqBlockTx(txModels, in.PeerType, sequencePeerID, minNextPeerSeq, int64(len(pending)))
			if err != nil {
				return err
			}
			for offset, index := range pending {
				item := in.Items[index]
				messageDate := unixOrNow(item.MessageDate)
				peerSeq := firstPeerSeq + int64(offset)
				canonicalID, err := r.nextID(ctx, "next canonical message id")
				if err != nil {
					return err
				}
				canonicalInput := batchItemToCanonicalInput(in, item, results[index].SendStateID)
				if err := insertCanonicalMessageTx(txModels, canonicalID, sequencePeerID, peerSeq, messageDate, canonicalInput); err != nil {
					return err
				}
				if err := insertClientRandomTx(txModels, canonicalID, peerSeq, messageDate, canonicalInput); err != nil {
					return err
				}
				if affected, err := txModels.MessageSendStatesModel.MarkCanonicalCreated(canonicalID, peerSeq, SendStateStatusCanonical, results[index].SendStateID); err != nil {
					return storageError("mark batch canonical created", err)
				} else if affected == 0 {
					return msg.ErrSendStateConflict
				}
				results[index] = CanonicalMessageResult{
					SendStateID:                  results[index].SendStateID,
					CanonicalMessageID:           canonicalID,
					PeerSeq:                      peerSeq,
					MessageDate:                  messageDate,
					RequestPayloadHash:           item.RequestPayloadHash,
					EntitiesPayloadSchemaVersion: canonicalEntitiesSchemaVersion(item.EntitiesPayloadSchemaVersion),
					EntitiesPayload:              item.EntitiesPayload,
					MediaRefSchemaVersion:        canonicalMediaSchemaVersion(item.MediaRefSchemaVersion),
					MediaRefPayload:              item.MediaRefPayload,
					MessageAttrsSchemaVersion:    item.MessageAttrsSchemaVersion,
					MessageAttrsPayload:          item.MessageAttrsPayload,
					ForwardRefSchemaVersion:      item.ForwardRefSchemaVersion,
					ForwardRefPayload:            item.ForwardRefPayload,
					ServiceActionSchemaVersion:   item.ServiceActionSchemaVersion,
					ServiceActionPayload:         item.ServiceActionPayload,
					SendStateStatus:              SendStateStatusCanonical,
					CreatedNew:                   true,
				}
			}
		}

		out = &CanonicalBatchResult{Items: results}
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
	bounds := in.ResolvedCursorBounds
	if !in.CursorsResolved {
		var err error
		bounds, err = r.ResolveHistoryCursorIDs(ctx, in.UserID, in.PeerType, in.PeerID, 0, in.MaxID, in.MinID)
		if err != nil {
			return nil, err
		}
	}
	if bounds.NoMatch {
		return []HistoryMessage{}, nil
	}

	limit := pagination.NormalizeLimit(in.Limit)
	rows, err := r.selectHistoryMessagesByUserMessageID(ctx, in, limit)
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

func (r *Repository) selectHistoryMessagesByUserMessageID(ctx context.Context, in ListHistoryMessagesInput, limit int32) ([]model.HistoryMessageRow, error) {
	offsetID := int64(in.OffsetID)
	if offsetID <= 0 {
		offsetID = math.MaxInt32
	}

	switch historyLoadType(in.AddOffset, limit) {
	case loadTypeHistoryBackward:
		return r.selectHistoryMessagesBackwardByUserMessageID(ctx, in, offsetID, limit+in.AddOffset)
	case loadTypeHistoryAround:
		forwardRows, err := r.selectHistoryMessagesForwardByUserMessageID(ctx, in, offsetID, -in.AddOffset)
		if err != nil {
			return nil, err
		}
		reverseHistoryRows(forwardRows)
		backwardRows, err := r.selectHistoryMessagesBackwardByUserMessageID(ctx, in, offsetID, limit+in.AddOffset)
		if err != nil {
			return nil, err
		}
		return append(forwardRows, backwardRows...), nil
	default:
		rows, err := r.selectHistoryMessagesForwardByUserMessageID(ctx, in, offsetID, -in.AddOffset)
		if err != nil {
			return nil, err
		}
		reverseHistoryRows(rows)
		return rows, nil
	}
}

func (r *Repository) ListRecentCanonicalMessagesBeforePeerSeq(ctx context.Context, peerType int32, peerID int64, beforePeerSeq int64, limit int32) ([]CanonicalMessage, error) {
	if _, err := r.requireDB(); err != nil {
		return nil, err
	}
	if peerType <= 0 || peerID <= 0 || beforePeerSeq <= 1 || limit <= 0 {
		return []CanonicalMessage{}, nil
	}
	rows, err := r.models.CanonicalMessagesModel.SelectRecentBeforePeerSeq(ctx, peerType, peerID, beforePeerSeq, MessageStatusLive, limit)
	if err != nil {
		return nil, storageError("list recent canonical messages before peer seq", err)
	}
	out := make([]CanonicalMessage, 0, len(rows))
	for i := len(rows) - 1; i >= 0; i-- {
		item := canonicalMessageModelToDTO(&rows[i])
		if item != nil {
			out = append(out, *item)
		}
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
	var resolvedOffset *ResolvedMessageID
	if in.OffsetID > 0 {
		var err error
		resolvedOffset, err = r.ResolveMessageID(ctx, in.UserID, in.PeerType, in.PeerID, int64(in.OffsetID))
		if err != nil {
			return nil, err
		}
	}
	offsetPeerSeq, noMatch := searchHashTagOffsetPeerSeq(in.OffsetID, resolvedOffset)
	if noMatch {
		return []HistoryMessage{}, nil
	}
	rows, err := r.models.CanonicalQueries.SearchHashTagMessages(ctx, tag, in.UserID, in.PeerType, in.PeerID, MessageStatusLive, offsetPeerSeq, offsetPeerSeq, likeTag, limit)
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

func searchHashTagOffsetPeerSeq(offsetID int32, resolved *ResolvedMessageID) (int64, bool) {
	if offsetID <= 0 {
		return 0, false
	}
	if resolved == nil {
		return 0, true
	}
	return resolved.PeerSeq, false
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
	return canonicalMessageModelToDTO(row), nil
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

const (
	loadTypeHistoryBackward = iota
	loadTypeHistoryAround
	loadTypeHistoryForward
)

func historyLoadType(addOffset int32, limit int32) int {
	if addOffset >= 0 {
		return loadTypeHistoryBackward
	}
	if addOffset+limit > 0 {
		return loadTypeHistoryAround
	}
	return loadTypeHistoryForward
}

func (r *Repository) selectHistoryMessagesBackwardByUserMessageID(ctx context.Context, in ListHistoryMessagesInput, offsetID int64, limit int32) ([]model.HistoryMessageRow, error) {
	if limit <= 0 {
		return []model.HistoryMessageRow{}, nil
	}
	return r.models.CanonicalQueries.SelectHistoryMessagesBackwardByUserMessageID(ctx, in.UserID, in.PeerType, in.PeerID, MessageStatusLive, offsetID, limit)
}

func (r *Repository) selectHistoryMessagesForwardByUserMessageID(ctx context.Context, in ListHistoryMessagesInput, offsetID int64, limit int32) ([]model.HistoryMessageRow, error) {
	if limit <= 0 {
		return []model.HistoryMessageRow{}, nil
	}
	return r.models.CanonicalQueries.SelectHistoryMessagesForwardByUserMessageID(ctx, in.UserID, in.PeerType, in.PeerID, MessageStatusLive, offsetID, limit)
}

func reverseHistoryRows(rows []model.HistoryMessageRow) {
	for i, j := 0, len(rows)-1; i < j; i, j = i+1, j-1 {
		rows[i], rows[j] = rows[j], rows[i]
	}
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
		case payload.MessageEventSchemaVersionV3:
			var event payload.MessageEventV3
			if err := json.Unmarshal(r.ViewPayload, &event); err != nil {
				return HistoryMessage{}, storageError("decode history view payload v3", err)
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
		ViewPayload:          append([]byte(nil), r.ViewPayload...),
		ReplyToPeerSeq:       replyToPeerSeq,
		ReplyToUserMessageID: replyToUserMessageID,
	}, nil
}

func historyMessageWithinBounds(peerSeq int64, bounds HistoryCursorBounds) bool {
	if bounds.NoMatch {
		return false
	}
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
		SendStateID:                  sendStateID,
		CanonicalMessageID:           row.CanonicalMessageId,
		PeerSeq:                      row.PeerSeq,
		MessageDate:                  row.Date,
		EntitiesPayloadSchemaVersion: row.EntitiesPayloadSchemaVersion,
		EntitiesPayload:              row.EntitiesPayload,
		MediaRefSchemaVersion:        row.MediaRefSchemaVersion,
		MediaRefPayload:              row.MediaRefPayload,
		MessageAttrsSchemaVersion:    row.MessageAttrsSchemaVersion,
		MessageAttrsPayload:          row.MessageAttrsPayload,
		ForwardRefSchemaVersion:      row.ForwardRefSchemaVersion,
		ForwardRefPayload:            row.ForwardRefPayload,
		ServiceActionSchemaVersion:   row.ServiceActionSchemaVersion,
		ServiceActionPayload:         row.ServiceActionPayload,
		CreatedNew:                   false,
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
	return canonicalMessageMatchesRetryInput(row, in), nil
}

func canonicalMessageMatchesRetryInput(row *model.CanonicalMessages, in CreateCanonicalMessageInput) bool {
	if row.MessageText != in.MessageText ||
		row.FromUserId != in.SenderUserID ||
		row.MessageKind != canonicalMessageKind(in) {
		return false
	}
	if canonicalRetryHasNoRichPayload(row, in) {
		return true
	}
	return row.EntitiesPayloadSchemaVersion == canonicalEntitiesSchemaVersion(in.EntitiesPayloadSchemaVersion) &&
		payloadBytesEqual(row.EntitiesPayload, in.EntitiesPayload) &&
		row.MediaRefSchemaVersion == canonicalMediaSchemaVersion(in.MediaRefSchemaVersion) &&
		payloadBytesEqual(row.MediaRefPayload, in.MediaRefPayload) &&
		row.MessageAttrsSchemaVersion == in.MessageAttrsSchemaVersion &&
		payloadBytesEqual(row.MessageAttrsPayload, in.MessageAttrsPayload) &&
		row.ForwardRefSchemaVersion == in.ForwardRefSchemaVersion &&
		payloadBytesEqual(row.ForwardRefPayload, in.ForwardRefPayload) &&
		row.ServiceActionSchemaVersion == in.ServiceActionSchemaVersion &&
		payloadBytesEqual(row.ServiceActionPayload, in.ServiceActionPayload)
}

func canonicalRetryHasNoRichPayload(row *model.CanonicalMessages, in CreateCanonicalMessageInput) bool {
	return len(row.EntitiesPayload) == 0 &&
		len(in.EntitiesPayload) == 0 &&
		len(row.MediaRefPayload) == 0 &&
		len(in.MediaRefPayload) == 0 &&
		len(row.MessageAttrsPayload) == 0 &&
		len(in.MessageAttrsPayload) == 0 &&
		len(row.ForwardRefPayload) == 0 &&
		len(in.ForwardRefPayload) == 0 &&
		len(row.ServiceActionPayload) == 0 &&
		len(in.ServiceActionPayload) == 0
}

func payloadBytesEqual(a []byte, b []byte) bool {
	return bytes.Equal(a, b) || len(a) == 0 && len(b) == 0
}

func nextPeerSeqTx(txModels *model.TxModels, peerType int32, peerID int64, minNextPeerSeq int64) (int64, error) {
	return nextPeerSeqBlockTx(txModels, peerType, peerID, minNextPeerSeq, 1)
}

func nextPeerSeqBlockTx(txModels *model.TxModels, peerType int32, peerID int64, minNextPeerSeq int64, count int64) (int64, error) {
	if count <= 0 {
		return 0, msg.ErrSendStateConflict
	}
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
	if affected, err := txModels.MessagePeerSequencesModel.UpdateNextPeerSeq(peerSeq+count, peerType, peerID); err != nil {
		return 0, storageError("advance peer sequence", err)
	} else if affected == 0 {
		return 0, msg.ErrSendStateConflict
	}
	return peerSeq, nil
}

func batchItemToCanonicalInput(batch CreateCanonicalBatchInput, item CreateCanonicalBatchItem, sendStateID int64) CreateCanonicalMessageInput {
	return CreateCanonicalMessageInput{
		SendStateID:                  sendStateID,
		SenderUserID:                 batch.SenderUserID,
		PeerType:                     batch.PeerType,
		PeerID:                       batch.PeerID,
		ClientRandomID:               item.ClientRandomID,
		RequestPayloadHash:           item.RequestPayloadHash,
		MessageText:                  item.MessageText,
		MessageDate:                  item.MessageDate,
		EntitiesPayloadSchemaVersion: item.EntitiesPayloadSchemaVersion,
		EntitiesPayload:              item.EntitiesPayload,
		MediaRefSchemaVersion:        item.MediaRefSchemaVersion,
		MediaRefPayload:              item.MediaRefPayload,
		MessageAttrsSchemaVersion:    item.MessageAttrsSchemaVersion,
		MessageAttrsPayload:          item.MessageAttrsPayload,
		ForwardRefSchemaVersion:      item.ForwardRefSchemaVersion,
		ForwardRefPayload:            item.ForwardRefPayload,
		ServiceActionSchemaVersion:   item.ServiceActionSchemaVersion,
		ServiceActionPayload:         item.ServiceActionPayload,
	}
}

func insertCanonicalMessageTx(txModels *model.TxModels, canonicalID int64, canonicalPeerID int64, peerSeq int64, messageDate int64, in CreateCanonicalMessageInput) error {
	_, _, err := txModels.CanonicalMessagesModel.Insert(&model.CanonicalMessages{
		CanonicalMessageId:           canonicalID,
		PeerType:                     in.PeerType,
		PeerId:                       canonicalPeerID,
		PeerSeq:                      peerSeq,
		FromUserId:                   in.SenderUserID,
		MessageKind:                  canonicalMessageKind(in),
		MessageText:                  in.MessageText,
		EntitiesPayloadSchemaVersion: canonicalEntitiesSchemaVersion(in.EntitiesPayloadSchemaVersion),
		EntitiesPayload:              in.EntitiesPayload,
		MediaRefSchemaVersion:        canonicalMediaSchemaVersion(in.MediaRefSchemaVersion),
		MediaRefPayload:              in.MediaRefPayload,
		MessageAttrsSchemaVersion:    in.MessageAttrsSchemaVersion,
		MessageAttrsPayload:          in.MessageAttrsPayload,
		ForwardRefSchemaVersion:      in.ForwardRefSchemaVersion,
		ForwardRefPayload:            in.ForwardRefPayload,
		ServiceActionSchemaVersion:   in.ServiceActionSchemaVersion,
		ServiceActionPayload:         in.ServiceActionPayload,
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

func canonicalEntitiesSchemaVersion(version int32) int32 {
	if version > 0 {
		return version
	}
	return 1
}

func canonicalMediaSchemaVersion(version int32) int32 {
	if version > 0 {
		return version
	}
	return 1
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
		SendStateID:                  r.SendStateID,
		CanonicalMessageID:           r.CanonicalMessageID,
		PeerSeq:                      r.PeerSeq,
		MessageDate:                  r.MessageDate,
		RequestPayloadHash:           r.RequestPayloadHash,
		EntitiesPayloadSchemaVersion: r.EntitiesPayloadSchemaVersion,
		EntitiesPayload:              r.EntitiesPayload,
		MediaRefSchemaVersion:        r.MediaRefSchemaVersion,
		MediaRefPayload:              r.MediaRefPayload,
		MessageAttrsSchemaVersion:    r.MessageAttrsSchemaVersion,
		MessageAttrsPayload:          r.MessageAttrsPayload,
		ForwardRefSchemaVersion:      r.ForwardRefSchemaVersion,
		ForwardRefPayload:            r.ForwardRefPayload,
		ServiceActionSchemaVersion:   r.ServiceActionSchemaVersion,
		ServiceActionPayload:         r.ServiceActionPayload,
		CreatedNew:                   created,
	}
}

func attachSendStateToCanonicalResult(result *CanonicalMessageResult, state *SendState) {
	if result == nil || state == nil {
		return
	}
	result.SendStateID = state.SendStateID
	result.SendStateStatus = state.Status
	result.SenderOperationID = state.SenderOperationID
	result.SenderPTS = state.SenderPTS
	result.SenderPTSCount = state.SenderPTSCount
	result.SenderUpdateSchemaVersion = state.SenderUpdateSchemaVersion
	result.SenderUpdatePayload = state.SenderUpdatePayload
	result.SenderUpdatePayloadHash = state.SenderUpdatePayloadHash
}

func canonicalMessageModelToDTO(row *model.CanonicalMessages) *CanonicalMessage {
	if row == nil {
		return nil
	}
	return &CanonicalMessage{
		CanonicalMessageID:           row.CanonicalMessageId,
		PeerSeq:                      row.PeerSeq,
		FromUserID:                   row.FromUserId,
		PeerType:                     row.PeerType,
		PeerID:                       row.PeerId,
		MessageKind:                  row.MessageKind,
		MessageText:                  row.MessageText,
		MessageDate:                  row.Date,
		EntitiesPayloadSchemaVersion: row.EntitiesPayloadSchemaVersion,
		EntitiesPayload:              row.EntitiesPayload,
		MediaRefSchemaVersion:        row.MediaRefSchemaVersion,
		MediaRefPayload:              row.MediaRefPayload,
		MessageAttrsSchemaVersion:    row.MessageAttrsSchemaVersion,
		MessageAttrsPayload:          row.MessageAttrsPayload,
		ForwardRefSchemaVersion:      row.ForwardRefSchemaVersion,
		ForwardRefPayload:            row.ForwardRefPayload,
		ServiceActionSchemaVersion:   row.ServiceActionSchemaVersion,
		ServiceActionPayload:         row.ServiceActionPayload,
	}
}

func canonicalMessageKind(in CreateCanonicalMessageInput) int32 {
	if len(in.ServiceActionPayload) > 0 {
		return MessageKindService
	}
	if len(in.MediaRefPayload) > 0 {
		return MessageKindMedia
	}
	return MessageKindText
}
