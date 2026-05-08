package repository

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"

	"github.com/teamgram/teamgram-server/v2/app/messenger/userupdates/internal/repository/model"
	"github.com/teamgram/teamgram-server/v2/app/messenger/userupdates/payload"
	"github.com/teamgram/teamgram-server/v2/app/messenger/userupdates/userupdates"
)

func nextUserMessageID(txModels *model.TxModels, userID int64) (int64, error) {
	if _, _, err := txModels.UserMessageSequencesModel.InsertIgnore(&model.UserMessageSequences{
		UserId:            userID,
		NextUserMessageId: 1,
	}); err != nil {
		return 0, storageError("init user message sequence", err)
	}
	seq, err := txModels.UserMessageSequencesModel.SelectForUpdate(userID)
	if err != nil {
		return 0, storageError("lock user message sequence", err)
	}
	if seq.NextUserMessageId <= 0 {
		return 0, fmt.Errorf("%w: invalid next user message id %d for user %d", userupdates.ErrUserupdatesStorage, seq.NextUserMessageId, userID)
	}
	id := seq.NextUserMessageId
	if _, err := txModels.UserMessageSequencesModel.UpdateNext(id+1, userID); err != nil {
		return 0, storageError("advance user message sequence", err)
	}
	return id, nil
}

func existingUserMessageID(txModels *model.TxModels, userID, canonicalMessageID int64) (int64, bool, error) {
	row, found, err := existingUserMessageView(txModels, userID, canonicalMessageID)
	if err != nil || !found {
		return 0, found, err
	}
	if row.UserMessageId <= 0 {
		return 0, true, nil
	}
	return row.UserMessageId, true, nil
}

func existingUserMessageView(txModels *model.TxModels, userID, canonicalMessageID int64) (*model.UserMessageViews, bool, error) {
	if canonicalMessageID == 0 {
		return nil, false, nil
	}
	row, err := txModels.UserMessageViewsModel.SelectByUserCanonical(userID, canonicalMessageID)
	if err != nil {
		if errors.Is(err, model.ErrNotFound) {
			return nil, false, nil
		}
		return nil, false, storageError("select existing user message view", err)
	}
	return row, true, nil
}

func ensureExistingMessageViewMatchesOperation(row *model.UserMessageViews, op payload.MessageOperationV1) error {
	if row == nil {
		return fmt.Errorf("%w: missing existing canonical message view", userupdates.ErrUserupdatesStorage)
	}
	if row.UserMessageId <= 0 {
		return fmt.Errorf("%w: existing canonical message view missing user_message_id", userupdates.ErrUserupdatesStorage)
	}
	if row.MessageStatus != MessageStatusLive ||
		row.CanonicalMessageId != op.CanonicalMessageID ||
		row.PeerType != op.PeerType ||
		row.PeerId != op.PeerID ||
		row.PeerSeq != op.PeerSeq ||
		row.FromUserId != op.FromUserID ||
		row.Outgoing != op.Out ||
		row.Date != int64(op.Date) {
		return userupdates.ErrOperationPayloadConflict
	}
	switch row.ViewSchemaVersion {
	case payload.MessageEventSchemaVersionV1:
		var event payload.MessageEventV1
		if err := json.Unmarshal(row.ViewPayload, &event); err != nil {
			return fmt.Errorf("%w: decode existing legacy message view: %v", userupdates.ErrUserupdatesStorage, err)
		}
		if event.SchemaVersion != payload.MessageEventSchemaVersionV1 ||
			event.EventKind != payload.EventKindNewMessage ||
			event.CanonicalMessageID != op.CanonicalMessageID ||
			event.MessageID != op.PeerSeq ||
			event.PeerType != op.PeerType ||
			event.PeerID != op.PeerID ||
			event.FromUserID != op.FromUserID ||
			event.ToUserID != op.ToUserID ||
			event.Date != op.Date ||
			event.Out != op.Out ||
			event.MessageText != op.MessageText ||
			!reflect.DeepEqual(event.Entities, op.Entities) {
			return userupdates.ErrOperationPayloadConflict
		}
		if op.ReplyToPeerSeq != 0 && event.ReplyToPeerSeq != op.ReplyToPeerSeq {
			return userupdates.ErrOperationPayloadConflict
		}
	case payload.MessageEventSchemaVersion:
		var event payload.MessageEventV2
		if err := json.Unmarshal(row.ViewPayload, &event); err != nil {
			return fmt.Errorf("%w: decode existing message view: %v", userupdates.ErrUserupdatesStorage, err)
		}
		if event.SchemaVersion != payload.MessageEventSchemaVersion ||
			event.EventKind != payload.EventKindNewMessage ||
			event.CanonicalMessageID != op.CanonicalMessageID ||
			event.PeerSeq != op.PeerSeq ||
			event.MessageID != row.UserMessageId ||
			event.PeerType != op.PeerType ||
			event.PeerID != op.PeerID ||
			event.FromUserID != op.FromUserID ||
			event.ToUserID != op.ToUserID ||
			event.Date != op.Date ||
			event.Out != op.Out ||
			event.MessageText != op.MessageText ||
			!reflect.DeepEqual(event.Entities, op.Entities) {
			return userupdates.ErrOperationPayloadConflict
		}
		if op.UserMessageID != 0 && event.MessageID != op.UserMessageID {
			return userupdates.ErrOperationPayloadConflict
		}
		if op.ReplyToUserMessageID != 0 && event.ReplyToUserMessageID != op.ReplyToUserMessageID {
			return userupdates.ErrOperationPayloadConflict
		}
	default:
		return fmt.Errorf("%w: unsupported existing message view schema=%d", userupdates.ErrUserupdatesStorage, row.ViewSchemaVersion)
	}
	return nil
}

func resolveUserMessageIDByPeerSeq(txModels *model.TxModels, userID int64, peerType int32, peerID, peerSeq int64) (int64, error) {
	if peerSeq <= 0 {
		return 0, nil
	}
	row, err := txModels.UserMessageViewsModel.SelectNearestLiveByUserPeerSeq(userID, peerType, peerID, peerSeq, MessageStatusLive)
	if err != nil {
		if errors.Is(err, model.ErrNotFound) {
			return 0, nil
		}
		return 0, storageError("resolve user message id by peer seq", err)
	}
	return row.UserMessageId, nil
}

func createsUserMessageView(kind string) bool {
	return kind == payload.OperationKindSendMessage
}
