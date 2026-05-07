package repository

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"math"
	"strings"
	"time"
	"unicode"
	"unicode/utf8"

	"github.com/teamgram/marmota/pkg/stores/sqlx"
	"github.com/teamgram/teamgram-server/v2/app/messenger/userupdates/internal/repository/model"
	"github.com/teamgram/teamgram-server/v2/app/messenger/userupdates/payload"
	"github.com/teamgram/teamgram-server/v2/app/messenger/userupdates/userupdates"
)

func (r *Repository) ClaimPartitionOwner(ctx context.Context, partitionID int32) (int64, error) {
	_, err := r.requireDB()
	if err != nil {
		return 0, err
	}
	if _, _, err := r.models.UserupdatesPartitionFencesModel.InsertIgnore(ctx, &model.UserupdatesPartitionFences{
		PartitionId:     partitionID,
		OwnerEpoch:      0,
		OwnerInstanceId: "unassigned",
		LeaseId:         "",
	}); err != nil {
		return 0, storageError("insert partition fence", err)
	}

	fence, err := r.models.UserupdatesPartitionFencesModel.SelectByPartitionId(ctx, partitionID)
	if err != nil {
		return 0, storageError("select partition fence", err)
	}

	affected, err := r.models.UserupdatesPartitionFencesModel.CasAcquireOwner(ctx, r.OwnerInstance(), "", partitionID, fence.OwnerEpoch)
	if err != nil {
		return 0, storageError("claim partition owner", err)
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
	if !bytes.Equal(in.PayloadHash, payload.HashBytes(in.Payload)) {
		return nil, userupdates.ErrOperationPayloadConflict
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
	txModels := r.models.WithTx(tx)
	fence, err := r.lockPartitionFence(txModels, in.PartitionID)
	if err != nil {
		return nil, err
	}
	if fence.OwnerInstanceId != r.OwnerInstance() {
		return nil, userupdates.ErrNotOwner
	}

	if _, _, err := txModels.UserPtsStateModel.InsertIgnore(&model.UserPtsState{
		UserId:       in.UserID,
		Pts:          0,
		PtsUpdatedAt: mysqlNow(),
		PartitionId:  in.PartitionID,
		OwnerEpoch:   fence.OwnerEpoch,
		RowVersion:   0,
	}); err != nil {
		return nil, storageError("init user pts state", err)
	}

	state, err := r.lockUserPTSState(txModels, in.UserID)
	if err != nil {
		return nil, err
	}
	if state.PartitionId != in.PartitionID {
		return nil, fmt.Errorf("%w: user %d partition %d != operation partition %d", userupdates.ErrNotOwner, in.UserID, state.PartitionId, in.PartitionID)
	}

	existing, found, err := selectOperationResult(txModels, in.UserID, in.OperationID)
	if err != nil {
		return nil, err
	}
	if found {
		if !bytes.Equal(existing.PayloadHash, in.PayloadHash) {
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
	if op.SchemaVersion != payload.MessageOperationSchemaVersion {
		return nil, fmt.Errorf("%w: unsupported operation schema=%d kind=%s", userupdates.ErrOperationTerminal, op.SchemaVersion, op.OperationKind)
	}
	if len(in.DependencyPts) != 0 || len(op.DependencyPts) != 0 {
		return nil, userupdates.ErrOperationTerminal
	}
	if op.ClearDraft && op.SourcePermAuthKeyID == 0 {
		return nil, fmt.Errorf("%w: clear draft side effect requires source permanent auth key", userupdates.ErrOperationTerminal)
	}
	if op.ReplyToCanonicalMessageID != 0 {
		row, err := txModels.UserMessageViewsModel.SelectByUserCanonical(in.UserID, op.ReplyToCanonicalMessageID)
		if err != nil {
			if errors.Is(err, model.ErrNotFound) {
				return nil, fmt.Errorf("%w: reply target canonical_message_id=%d not visible to user_id=%d", userupdates.ErrOperationTerminal, op.ReplyToCanonicalMessageID, in.UserID)
			}
			return nil, storageError("select reply target view", err)
		}
		op.ReplyToPeerSeq = row.PeerSeq
	}

	nextPTS := state.Pts + 1
	ptsCount := int32(1)
	eventPayload, eventPayloadHash, responsePayload, responsePayloadHash, err := buildEventAndResponse(in, op, nextPTS, ptsCount)
	if err != nil {
		return nil, err
	}
	switch op.OperationKind {
	case payload.OperationKindSendMessage:
		if err := insertUserMessageView(txModels, in, op, eventPayload); err != nil {
			return nil, err
		}
		if err := upsertUserDialog(txModels, in, op, nextPTS, eventPayload); err != nil {
			return nil, err
		}
		if err := insertDialogSideEffects(ctx, txModels, r, in, op); err != nil {
			return nil, err
		}
	case payload.OperationKindReadHistory:
		if err := applyReadHistory(txModels, in, op, nextPTS); err != nil {
			return nil, err
		}
	case payload.OperationKindDeleteMessages:
		if err := applyDeleteMessages(tx, txModels, in, op, nextPTS); err != nil {
			return nil, err
		}
	case payload.OperationKindDeleteHistory:
		if err := applyDeleteHistory(tx, txModels, in, op, nextPTS); err != nil {
			return nil, err
		}
	case payload.OperationKindEditMessage:
		if err := applyEditMessage(tx, txModels, in, op, eventPayload, nextPTS); err != nil {
			return nil, err
		}
	case payload.OperationKindUpdatePinnedMessage:
		if err := applyUpdatePinnedMessage(txModels, in, op, nextPTS); err != nil {
			return nil, err
		}
	case payload.OperationKindMarkDialogUnread:
		if err := applyMarkDialogUnread(txModels, in, op, nextPTS); err != nil {
			return nil, err
		}
	case payload.OperationKindScheduledMarker:
		if err := applyScheduledMarker(txModels, in, op, nextPTS); err != nil {
			return nil, err
		}
	default:
		return nil, fmt.Errorf("%w: unsupported operation kind=%s", userupdates.ErrOperationTerminal, op.OperationKind)
	}
	if err := insertPTSEvent(txModels, in, op, nextPTS, ptsCount, eventPayload, eventPayloadHash); err != nil {
		return nil, err
	}
	if err := insertPushTask(ctx, txModels, r, in, op, nextPTS, eventPayload); err != nil {
		return nil, err
	}
	if err := insertOperationResult(txModels, in, nextPTS, ptsCount, responsePayload, responsePayloadHash); err != nil {
		return nil, err
	}
	affected, err := txModels.UserPtsStateModel.UpdatePts(nextPTS, mysqlNow(), in.PartitionID, fence.OwnerEpoch, in.UserID)
	if err != nil {
		return nil, storageError("update user pts state", err)
	}
	if affected == 0 {
		return nil, userupdates.ErrPtsContinuityViolation
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

func (r *Repository) lockPartitionFence(txModels *model.TxModels, partitionID int32) (*model.UserupdatesPartitionFences, error) {
	fence, err := txModels.UserupdatesPartitionFencesModel.SelectByPartitionId(partitionID)
	if err != nil {
		if errors.Is(err, model.ErrNotFound) {
			return nil, userupdates.ErrNotOwner
		}
		return nil, storageError("lock partition fence", err)
	}
	return fence, nil
}

func (r *Repository) lockUserPTSState(txModels *model.TxModels, userID int64) (*model.UserPtsState, error) {
	state, err := txModels.UserPtsStateModel.SelectForUpdate(userID)
	if err != nil {
		return nil, storageError("lock user pts state", err)
	}
	return state, nil
}

func buildEventAndResponse(in ApplyUserOperationInput, op payload.MessageOperationV1, pts int64, ptsCount int32) ([]byte, []byte, []byte, []byte, error) {
	eventKind := payload.EventKindNewMessage
	if op.OperationKind != payload.OperationKindSendMessage {
		eventKind = op.OperationKind
	}
	event := payload.MessageEventV1{
		SchemaVersion:      payload.MessageEventSchemaVersion,
		EventKind:          eventKind,
		CanonicalMessageID: op.CanonicalMessageID,
		MessageID:          op.PeerSeq,
		PeerType:           op.PeerType,
		PeerID:             op.PeerID,
		FromUserID:         op.FromUserID,
		ToUserID:           op.ToUserID,
		Date:               op.Date,
		EditDate:           op.EditDate,
		EditVersion:        op.EditVersion,
		Out:                op.Out,
		MessageText:        op.MessageText,
		Entities:           op.Entities,
		ReplyToPeerSeq:     op.ReplyToPeerSeq,
		AuthKeyIdExclude:   in.AuthKeyIDExclude,
	}
	eventPayload, err := json.Marshal(event)
	if err != nil {
		return nil, nil, nil, nil, storageError("marshal message event", err)
	}
	response := payload.OperationResponseV1{
		SchemaVersion: payload.OperationResponseSchemaVersion,
		OperationID:   in.OperationID,
		Pts:           pts,
		PtsCount:      ptsCount,
		EventType:     eventKind,
	}
	responsePayload, err := json.Marshal(response)
	if err != nil {
		return nil, nil, nil, nil, storageError("marshal operation response", err)
	}
	return eventPayload, payload.HashBytes(eventPayload), responsePayload, payload.HashBytes(responsePayload), nil
}

func insertUserMessageView(txModels *model.TxModels, in ApplyUserOperationInput, op payload.MessageOperationV1, viewPayload []byte) error {
	_, _, err := txModels.UserMessageViewsModel.InsertOrUpdate(&model.UserMessageViews{
		UserId:             in.UserID,
		PeerType:           op.PeerType,
		PeerId:             op.PeerID,
		PeerSeq:            op.PeerSeq,
		CanonicalMessageId: op.CanonicalMessageID,
		FromUserId:         op.FromUserID,
		Outgoing:           op.Out,
		MessageKind:        MessageKindText,
		MessageStatus:      MessageStatusLive,
		EditVersion:        0,
		Date:               mysqlDate(op.Date),
		EditDate:           mysqlNullInvalid(),
		DeletedAt:          mysqlNullInvalid(),
		ViewSchemaVersion:  1,
		ViewPayload:        viewPayload,
	})
	if err != nil {
		return storageError("insert user message view", err)
	}
	if err := insertHashTagsTx(txModels, in.UserID, op); err != nil {
		return err
	}
	return nil
}

func upsertUserDialog(txModels *model.TxModels, in ApplyUserOperationInput, op payload.MessageOperationV1, nextPTS int64, dialogPayload []byte) error {
	unread := int32(0)
	if !op.Out {
		unread = 1
	}
	now := mysqlNow()
	_, _, err := txModels.UserDialogsModel.InsertOrUpdateMessageEvent(&model.UserDialogs{
		UserId:                   in.UserID,
		PeerType:                 op.PeerType,
		PeerId:                   op.PeerID,
		TopPeerSeq:               op.PeerSeq,
		TopCanonicalMessageId:    op.CanonicalMessageID,
		TopMessageDate:           mysqlNullDate(op.Date),
		TopMessageStatus:         MessageStatusLive,
		ReadInboxMaxPeerSeq:      0,
		ReadOutboxMaxPeerSeq:     0,
		UnreadCount:              unread,
		UnreadMentionsCount:      0,
		UnreadReactionsCount:     0,
		UnreadMark:               false,
		PinnedPeerSeq:            0,
		PinnedCanonicalMessageId: 0,
		HasScheduled:             false,
		AvailableMinPeerSeq:      0,
		Hidden:                   false,
		DeletedAt:                mysqlNullInvalid(),
		LastPts:                  nextPTS,
		LastPtsAt:                mysqlNullTime(now),
		DialogSchemaVersion:      1,
		DialogPayload:            dialogPayload,
	})
	if err != nil {
		return storageError("upsert user dialog", err)
	}
	return nil
}

func applyReadHistory(txModels *model.TxModels, in ApplyUserOperationInput, op payload.MessageOperationV1, nextPTS int64) error {
	readInbox := op.ReadInboxMaxPeerSeq
	readOutbox := op.ReadOutboxMaxPeerSeq
	if readInbox == 0 || readOutbox == 0 {
		if row, err := txModels.UserDialogsModel.SelectByUserPeer(in.UserID, op.PeerType, op.PeerID); err == nil {
			if readInbox == 0 {
				readInbox = row.ReadInboxMaxPeerSeq
			}
			if readOutbox == 0 {
				readOutbox = row.ReadOutboxMaxPeerSeq
			}
		} else if !errors.Is(err, model.ErrNotFound) {
			return storageError("select dialog before read history", err)
		}
	}
	if readInbox == 0 && readOutbox == 0 {
		if op.Out {
			readOutbox = op.PeerSeq
		} else {
			readInbox = op.PeerSeq
		}
	}
	_, err := txModels.UserDialogsModel.UpdateReadState(0, 0, 0, false, readInbox, readOutbox, nextPTS, mysqlNullNow(), in.UserID, op.PeerType, op.PeerID)
	if err != nil {
		return storageError("apply read history", err)
	}
	if op.Out && readOutbox > 0 {
		if err := insertOutboxReadDateTx(txModels, in.UserID, op.PeerType, op.PeerID, op.PeerID, readOutbox, op.Date); err != nil {
			return err
		}
	}
	return nil
}

func insertHashTagsTx(txModels *model.TxModels, userID int64, op payload.MessageOperationV1) error {
	tags := extractHashTags(op.MessageText)
	if len(tags) == 0 || op.PeerSeq <= 0 {
		return nil
	}
	if op.PeerSeq > math.MaxInt32 {
		return storageError("index hashtag", fmt.Errorf("peer seq out of int32 range: %d", op.PeerSeq))
	}
	for _, tag := range tags {
		if _, _, err := txModels.HashTagsModel.InsertOrUpdate(&model.HashTags{
			UserId:           userID,
			PeerType:         op.PeerType,
			PeerId:           op.PeerID,
			HashTag:          tag,
			HashTagMessageId: int32(op.PeerSeq),
		}); err != nil {
			return storageError("index hashtag", err)
		}
	}
	return nil
}

func insertOutboxReadDateTx(txModels *model.TxModels, userID int64, peerType int32, peerID int64, readUserID int64, maxPeerSeq int64, date int32) error {
	if maxPeerSeq > math.MaxInt32 {
		return storageError("insert outbox read date", fmt.Errorf("peer seq out of int32 range: %d", maxPeerSeq))
	}
	if date <= 0 {
		date = int32(time.Now().Unix())
	}
	if _, _, err := txModels.MessageReadOutboxModel.InsertOrUpdate(&model.MessageReadOutbox{
		UserId:            userID,
		PeerType:          peerType,
		PeerId:            peerID,
		ReadUserId:        readUserID,
		ReadOutboxMaxId:   int32(maxPeerSeq),
		ReadOutboxMaxDate: mysqlDate(date),
	}); err != nil {
		return storageError("insert outbox read date", err)
	}
	return nil
}

func extractHashTags(text string) []string {
	seen := map[string]struct{}{}
	var out []string
	for i := 0; i < len(text); {
		r, size := utf8.DecodeRuneInString(text[i:])
		if r != '#' {
			i += size
			continue
		}
		start := i + size
		end := start
		for end < len(text) {
			next, nextSize := utf8.DecodeRuneInString(text[end:])
			if !isHashTagRune(next) {
				break
			}
			end += nextSize
		}
		if end > start {
			tag := strings.ToLower(text[start:end])
			if _, ok := seen[tag]; !ok {
				seen[tag] = struct{}{}
				out = append(out, tag)
			}
		}
		i = end
	}
	return out
}

func isHashTagRune(r rune) bool {
	return r == '_' || unicode.IsLetter(r) || unicode.IsDigit(r)
}

func applyDeleteMessages(tx *sqlx.Tx, txModels *model.TxModels, in ApplyUserOperationInput, op payload.MessageOperationV1, nextPTS int64) error {
	deletedAt := mysqlNow()
	for _, peerSeq := range op.DeletePeerSeqs {
		row, err := txModels.UserMessageViewsModel.SelectByUserPeerSeq(in.UserID, op.PeerType, op.PeerID, peerSeq)
		if err != nil {
			if errors.Is(err, model.ErrNotFound) {
				continue
			}
			return storageError("select message before delete", err)
		}
		row.MessageStatus = MessageStatusDeleted
		row.DeletedAt = mysqlNullTime(deletedAt)
		row.ViewPayload = []byte(`{"schema_version":1,"deleted":true}`)
		if _, _, err := txModels.UserMessageViewsModel.InsertOrUpdate(row); err != nil {
			return storageError("mark message deleted", err)
		}
	}
	return recomputeDialogTop(tx, in.UserID, op.PeerType, op.PeerID, nextPTS)
}

func applyDeleteHistory(tx *sqlx.Tx, txModels *model.TxModels, in ApplyUserOperationInput, op payload.MessageOperationV1, nextPTS int64) error {
	maxPeerSeq := op.DeleteMaxPeerSeq
	if maxPeerSeq == 0 {
		maxPeerSeq = op.PeerSeq
	}
	if maxPeerSeq <= 0 {
		if row, err := txModels.UserDialogsModel.SelectByUserPeer(in.UserID, op.PeerType, op.PeerID); err == nil {
			maxPeerSeq = row.TopPeerSeq
		} else if !errors.Is(err, model.ErrNotFound) {
			return storageError("select dialog before delete history", err)
		}
	}
	if maxPeerSeq > 0 {
		query := "update user_message_views set message_status = ?, view_payload = ? where user_id = ? and peer_type = ? and peer_id = ? and peer_seq <= ?"
		if _, err := tx.Exec(query, MessageStatusDeleted, []byte(`{"schema_version":1,"deleted":true}`), in.UserID, op.PeerType, op.PeerID, maxPeerSeq); err != nil {
			return storageError("mark history deleted", err)
		}
	}
	if err := recomputeDialogTop(tx, in.UserID, op.PeerType, op.PeerID, nextPTS); err != nil {
		return err
	}
	if op.JustClear || maxPeerSeq > 0 {
		query := "update user_dialogs set available_min_peer_seq = ?, last_pts = ?, last_pts_at = ? where user_id = ? and peer_type = ? and peer_id = ?"
		if _, err := tx.Exec(query, maxPeerSeq, nextPTS, mysqlNow(), in.UserID, op.PeerType, op.PeerID); err != nil {
			return storageError("update dialog available min peer seq", err)
		}
	}
	return nil
}

func applyEditMessage(tx *sqlx.Tx, txModels *model.TxModels, in ApplyUserOperationInput, op payload.MessageOperationV1, viewPayload []byte, nextPTS int64) error {
	row, err := txModels.UserMessageViewsModel.SelectByUserCanonical(in.UserID, op.CanonicalMessageID)
	if err != nil {
		if errors.Is(err, model.ErrNotFound) {
			return userupdates.ErrOperationTerminal
		}
		return storageError("select message before edit", err)
	}
	if row.MessageStatus != MessageStatusLive {
		return userupdates.ErrOperationTerminal
	}
	editVersion := op.EditVersion
	if editVersion == 0 {
		editVersion = row.EditVersion + 1
	}
	row.EditVersion = editVersion
	row.ViewPayload = viewPayload
	if _, _, err := txModels.UserMessageViewsModel.InsertOrUpdate(row); err != nil {
		return storageError("update message view after edit", err)
	}
	if op.EditDate > 0 {
		query := "update user_message_views set edit_date = ? where user_id = ? and canonical_message_id = ?"
		if _, err := tx.Exec(query, mysqlDate(op.EditDate), in.UserID, op.CanonicalMessageID); err != nil {
			return storageError("update message view edit date", err)
		}
	}
	if err := insertHashTagsTx(txModels, in.UserID, op); err != nil {
		return err
	}
	dialog, err := txModels.UserDialogsModel.SelectByUserPeer(in.UserID, op.PeerType, op.PeerID)
	if err != nil {
		if errors.Is(err, model.ErrNotFound) {
			return nil
		}
		return storageError("select dialog before edit", err)
	}
	if dialog.TopCanonicalMessageId != op.CanonicalMessageID {
		return nil
	}
	dialog.DialogPayload = viewPayload
	dialog.LastPts = nextPTS
	dialog.LastPtsAt = mysqlNullNow()
	if _, _, err := txModels.UserDialogsModel.InsertOrUpdateMessageEvent(dialog); err != nil {
		return storageError("update dialog payload after edit", err)
	}
	return nil
}

func recomputeDialogTop(tx *sqlx.Tx, userID int64, peerType int32, peerID int64, nextPTS int64) error {
	var top model.UserMessageViews
	query := "select user_id, peer_type, peer_id, peer_seq, canonical_message_id, from_user_id, outgoing, message_kind, message_status, edit_version, `date`, view_schema_version, view_payload from user_message_views where user_id = ? and peer_type = ? and peer_id = ? and message_status = ? order by peer_seq desc limit 1"
	if err := tx.QueryRowPartial(&top, query, userID, peerType, peerID, MessageStatusLive); err != nil {
		if !errors.Is(err, sqlx.ErrNotFound) && !errors.Is(err, model.ErrNotFound) {
			return storageError("select top message after delete", err)
		}
		update := "update user_dialogs set top_peer_seq = 0, top_canonical_message_id = 0, top_message_status = ?, hidden = 1, deleted_at = ?, last_pts = ?, last_pts_at = ? where user_id = ? and peer_type = ? and peer_id = ?"
		if _, execErr := tx.Exec(update, MessageStatusDeleted, mysqlNow(), nextPTS, mysqlNow(), userID, peerType, peerID); execErr != nil {
			return storageError("clear dialog top after delete", execErr)
		}
		return nil
	}
	update := "update user_dialogs set top_peer_seq = ?, top_canonical_message_id = ?, top_message_date = ?, top_message_status = ?, hidden = 0, deleted_at = ?, last_pts = ?, last_pts_at = ? where user_id = ? and peer_type = ? and peer_id = ?"
	if _, err := tx.Exec(update, top.PeerSeq, top.CanonicalMessageId, top.Date, top.MessageStatus, mysqlNullInvalid(), nextPTS, mysqlNow(), userID, peerType, peerID); err != nil {
		return storageError("update dialog top after delete", err)
	}
	return nil
}

func applyUpdatePinnedMessage(txModels *model.TxModels, in ApplyUserOperationInput, op payload.MessageOperationV1, nextPTS int64) error {
	pinnedPeerSeq := op.PinnedPeerSeq
	if pinnedPeerSeq == 0 {
		pinnedPeerSeq = op.PeerSeq
	}
	pinnedCanonicalID := op.PinnedCanonicalMessageID
	if pinnedCanonicalID == 0 {
		pinnedCanonicalID = op.CanonicalMessageID
	}
	_, err := txModels.UserDialogsModel.UpdatePinnedMessage(pinnedPeerSeq, pinnedCanonicalID, nextPTS, mysqlNullNow(), in.UserID, op.PeerType, op.PeerID)
	if err != nil {
		return storageError("apply update pinned message", err)
	}
	return nil
}

func applyMarkDialogUnread(txModels *model.TxModels, in ApplyUserOperationInput, op payload.MessageOperationV1, nextPTS int64) error {
	unreadMark := true
	if op.UnreadMark != nil {
		unreadMark = *op.UnreadMark
	}
	row, err := txModels.UserDialogsModel.SelectByUserPeer(in.UserID, op.PeerType, op.PeerID)
	if err != nil {
		if errors.Is(err, model.ErrNotFound) {
			return nil
		}
		return storageError("select dialog before mark unread", err)
	}
	_, err = txModels.UserDialogsModel.UpdateReadState(
		row.UnreadCount,
		row.UnreadMentionsCount,
		row.UnreadReactionsCount,
		unreadMark,
		row.ReadInboxMaxPeerSeq,
		row.ReadOutboxMaxPeerSeq,
		nextPTS,
		mysqlNullNow(),
		in.UserID,
		op.PeerType,
		op.PeerID,
	)
	if err != nil {
		return storageError("apply mark dialog unread", err)
	}
	return nil
}

func applyScheduledMarker(txModels *model.TxModels, in ApplyUserOperationInput, op payload.MessageOperationV1, nextPTS int64) error {
	hasScheduled := false
	if op.HasScheduled != nil {
		hasScheduled = *op.HasScheduled
	}
	_, _, err := txModels.UserDialogsModel.InsertOrUpdateMessageEvent(&model.UserDialogs{
		UserId:                   in.UserID,
		PeerType:                 op.PeerType,
		PeerId:                   op.PeerID,
		TopPeerSeq:               op.PeerSeq,
		TopCanonicalMessageId:    op.CanonicalMessageID,
		TopMessageDate:           mysqlNullDate(op.Date),
		TopMessageStatus:         MessageStatusLive,
		ReadInboxMaxPeerSeq:      0,
		ReadOutboxMaxPeerSeq:     0,
		UnreadCount:              0,
		UnreadMentionsCount:      0,
		UnreadReactionsCount:     0,
		UnreadMark:               false,
		PinnedPeerSeq:            0,
		PinnedCanonicalMessageId: 0,
		HasScheduled:             hasScheduled,
		AvailableMinPeerSeq:      0,
		Hidden:                   false,
		DeletedAt:                mysqlNullInvalid(),
		LastPts:                  nextPTS,
		LastPtsAt:                mysqlNullNow(),
		DialogSchemaVersion:      1,
		DialogPayload:            []byte(`{"schema_version":1}`),
	})
	if err != nil {
		return storageError("apply scheduled marker", err)
	}
	return nil
}

func insertPTSEvent(txModels *model.TxModels, in ApplyUserOperationInput, op payload.MessageOperationV1, pts int64, ptsCount int32, eventPayload []byte, eventPayloadHash []byte) error {
	eventType := EventTypeNewMessage
	switch op.OperationKind {
	case payload.OperationKindReadHistory:
		eventType = EventTypeReadHistory
	case payload.OperationKindDeleteMessages:
		eventType = EventTypeDeleteMessages
	case payload.OperationKindDeleteHistory:
		eventType = EventTypeDeleteHistory
	case payload.OperationKindEditMessage:
		eventType = EventTypeEditMessage
	case payload.OperationKindUpdatePinnedMessage:
		eventType = EventTypeUpdatePinnedMessage
	case payload.OperationKindMarkDialogUnread:
		eventType = EventTypeMarkDialogUnread
	case payload.OperationKindScheduledMarker:
		eventType = EventTypeScheduledMarker
	}
	_, _, err := txModels.UserPtsEventsModel.Insert(&model.UserPtsEvents{
		UserId:             in.UserID,
		Pts:                pts,
		PtsCount:           ptsCount,
		OperationId:        in.OperationID,
		OpType:             in.OpType,
		EventType:          eventType,
		PeerType:           op.PeerType,
		PeerId:             op.PeerID,
		CanonicalMessageId: op.CanonicalMessageID,
		PeerSeq:            op.PeerSeq,
		ActorUserId:        op.FromUserID,
		EventSchemaVersion: payload.MessageEventSchemaVersion,
		EventCodec:         PayloadCodecJSON,
		EventPayload:       eventPayload,
		EventPayloadHash:   eventPayloadHash,
	})
	if err != nil {
		return fmt.Errorf("%w: insert pts event: %w", userupdates.ErrPtsContinuityViolation, err)
	}
	return nil
}

func insertPushTask(ctx context.Context, txModels *model.TxModels, r *Repository, in ApplyUserOperationInput, op payload.MessageOperationV1, pts int64, taskPayload []byte) error {
	taskID, err := r.idgen.NextID(ctx)
	if err != nil {
		return storageError("next push task id", err)
	}
	route := payload.RouteUser(in.UserID)
	_, _, err = txModels.PushTaskOutboxModel.Insert(&model.PushTaskOutbox{
		TaskId:             taskID,
		UserId:             in.UserID,
		Pts:                pts,
		PushType:           PushTypeUserUpdate,
		PeerType:           op.PeerType,
		PeerId:             op.PeerID,
		OperationId:        in.OperationID,
		PushPartitionId:    int32(route.PushPartitionID),
		TaskSchemaVersion:  1,
		TaskCodec:          PayloadCodecJSON,
		TaskPayload:        taskPayload,
		Status:             PushTaskStatusPending,
		PublishAttempts:    0,
		AvailableAt:        mysqlNow(),
		PublishedTopic:     "",
		PublishedPartition: 0,
		PublishedOffset:    0,
		LastErrorCode:      "",
	})
	if err != nil {
		return storageError("insert push task", err)
	}
	return nil
}

func insertDialogSideEffects(ctx context.Context, txModels *model.TxModels, r *Repository, in ApplyUserOperationInput, op payload.MessageOperationV1) error {
	if op.ClearDraft {
		clearBeforeDate := op.ClearDraftBeforeDate
		if clearBeforeDate == 0 {
			clearBeforeDate = op.Date
		}
		body, err := json.Marshal(clearDraftSideEffectPayloadV1{
			SchemaVersion:      1,
			ClearBeforeDate:    clearBeforeDate,
			SourceMessageDate:  op.Date,
			SourcePeerSeq:      op.PeerSeq,
			CanonicalMessageID: op.CanonicalMessageID,
		})
		if err != nil {
			return storageError("marshal clear draft side effect", err)
		}
		sideEffectID, err := r.idgen.NextID(ctx)
		if err != nil {
			return storageError("next dialog side effect id", err)
		}
		if err := r.InsertDialogSideEffectTx(txModels, DialogSideEffect{
			SideEffectID:             sideEffectID,
			Kind:                     DialogSideEffectKindClearDraftAfterSend,
			UserID:                   in.UserID,
			PeerType:                 op.PeerType,
			PeerID:                   op.PeerID,
			SourcePermAuthKeyID:      op.SourcePermAuthKeyID,
			SourceOperationID:        in.OperationID,
			SourceMessageDate:        time.Unix(int64(op.Date), 0).UTC(),
			SourcePeerSeq:            op.PeerSeq,
			SourceCanonicalMessageID: op.CanonicalMessageID,
			ClearBeforeDate:          time.Unix(int64(clearBeforeDate), 0).UTC(),
			PayloadSchemaVersion:     1,
			Payload:                  body,
			PayloadHash:              payload.HashBytes(body),
			Status:                   DialogSideEffectStatusPending,
			AttemptCount:             0,
			NextRetryAt:              time.Now().UTC(),
		}); err != nil {
			return err
		}
	}
	if shouldWriteSavedDialogSideEffect(in, op) {
		body, err := json.Marshal(savedDialogSideEffectPayloadV1{
			SchemaVersion:         1,
			SavedPeerType:         op.PeerType,
			SavedPeerID:           op.PeerID,
			TopPeerSeq:            op.PeerSeq,
			TopCanonicalMessageID: op.CanonicalMessageID,
			MessageDate:           op.Date,
			Deleted:               false,
			Top:                   true,
		})
		if err != nil {
			return storageError("marshal saved dialog side effect", err)
		}
		sideEffectID, err := r.idgen.NextID(ctx)
		if err != nil {
			return storageError("next dialog side effect id", err)
		}
		if err := r.InsertDialogSideEffectTx(txModels, DialogSideEffect{
			SideEffectID:             sideEffectID,
			Kind:                     DialogSideEffectKindUpsertSavedDialogFromMessage,
			UserID:                   in.UserID,
			PeerType:                 op.PeerType,
			PeerID:                   op.PeerID,
			SourceOperationID:        in.OperationID,
			SourceMessageDate:        time.Unix(int64(op.Date), 0).UTC(),
			SourcePeerSeq:            op.PeerSeq,
			SourceCanonicalMessageID: op.CanonicalMessageID,
			PayloadSchemaVersion:     1,
			Payload:                  body,
			PayloadHash:              payload.HashBytes(body),
			Status:                   DialogSideEffectStatusPending,
			AttemptCount:             0,
			NextRetryAt:              time.Now().UTC(),
		}); err != nil {
			return err
		}
	}
	return nil
}

func shouldWriteSavedDialogSideEffect(in ApplyUserOperationInput, op payload.MessageOperationV1) bool {
	if op.SavedDialogSideEffect {
		return true
	}
	return op.PeerType == payload.PeerTypeUser && op.PeerID == in.UserID
}

func insertOperationResult(txModels *model.TxModels, in ApplyUserOperationInput, pts int64, ptsCount int32, responsePayload []byte, responseHash []byte) error {
	_, _, err := txModels.UserOperationResultsModel.Insert(&model.UserOperationResults{
		UserId:                in.UserID,
		OperationId:           in.OperationID,
		OpType:                in.OpType,
		Status:                OperationResultStatusCompleted,
		Pts:                   pts,
		PtsCount:              ptsCount,
		PayloadHash:           in.PayloadHash,
		ResponseSchemaVersion: payload.OperationResponseSchemaVersion,
		ResponseCodec:         PayloadCodecJSON,
		ResponsePayload:       responsePayload,
		ResponsePayloadHash:   responseHash,
		TerminalErrorCode:     "",
		CompletedAt:           mysqlNow(),
	})
	if err != nil {
		return storageError("insert operation result", err)
	}
	return nil
}
