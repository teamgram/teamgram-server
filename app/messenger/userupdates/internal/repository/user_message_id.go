package repository

import (
	"errors"
	"fmt"

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
	if canonicalMessageID == 0 {
		return 0, false, nil
	}
	row, err := txModels.UserMessageViewsModel.SelectByUserCanonical(userID, canonicalMessageID)
	if err != nil {
		if errors.Is(err, model.ErrNotFound) {
			return 0, false, nil
		}
		return 0, false, storageError("select existing user message view", err)
	}
	if row.UserMessageId <= 0 {
		return 0, true, nil
	}
	return row.UserMessageId, true, nil
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
