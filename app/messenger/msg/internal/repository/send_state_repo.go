package repository

import (
	"bytes"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"math"

	"github.com/teamgram/teamgram-server/v2/app/messenger/msg/internal/repository/model"
	"github.com/teamgram/teamgram-server/v2/app/messenger/msg/msg"
)

func (r *Repository) CreateOrLoadSendState(ctx context.Context, in CreateSendStateInput) (*SendState, error) {
	if _, err := r.requireDB(); err != nil {
		return nil, err
	}
	existing, found, err := r.selectSendStateByRandom(ctx, in.SenderUserID, in.PeerType, in.PeerID, in.ClientRandomID)
	if err != nil {
		return nil, err
	}
	if found {
		if !bytes.Equal(existing.RequestPayloadHash, in.RequestPayloadHash) {
			return nil, msg.ErrRandomIdConflict
		}
		return existing, nil
	}

	sendStateID, err := r.nextID(ctx, "next send state id")
	if err != nil {
		return nil, err
	}
	_, _, err = r.models.MessageSendStatesModel.Insert(ctx, &model.MessageSendStates{
		SendStateId:                 sendStateID,
		SenderUserId:                in.SenderUserID,
		PeerType:                    in.PeerType,
		PeerId:                      in.PeerID,
		ClientRandomId:              in.ClientRandomID,
		CanonicalMessageId:          0,
		PeerSeq:                     0,
		Status:                      SendStateStatusInitialized,
		RequestPayloadSchemaVersion: in.RequestPayloadSchemaVersion,
		RequestPayloadHash:          in.RequestPayloadHash,
		SenderOperationId:           "",
		SenderPts:                   0,
		SenderPtsCount:              0,
		SenderUpdateSchemaVersion:   0,
		SenderUpdatePayload:         nil,
		SenderUpdatePayloadHash:     nil,
		ReceiverManifestId:          0,
		LastErrorCategory:           0,
		LastErrorCode:               "",
		LastErrorMessage:            "",
		RetryCount:                  0,
		CompletedAt:                 sql.NullTime{},
	})
	if err != nil {
		again, found, selectErr := r.selectSendStateByRandom(ctx, in.SenderUserID, in.PeerType, in.PeerID, in.ClientRandomID)
		if selectErr == nil && found {
			if !bytes.Equal(again.RequestPayloadHash, in.RequestPayloadHash) {
				return nil, msg.ErrRandomIdConflict
			}
			return again, nil
		}
		return nil, storageError("insert send state", err)
	}
	return r.selectSendStateByID(ctx, sendStateID)
}

func (r *Repository) MarkCanonicalCreated(ctx context.Context, sendStateID int64, canonicalMessageID int64, peerSeq int64) error {
	if _, err := r.requireDB(); err != nil {
		return err
	}
	affected, err := r.models.MessageSendStatesModel.MarkCanonicalCreated(ctx, canonicalMessageID, peerSeq, SendStateStatusCanonical, sendStateID)
	if err != nil {
		return storageError("mark canonical created", err)
	}
	return requireAffectedRows(affected, "mark canonical created")
}

func (r *Repository) MarkSenderCommitted(ctx context.Context, in MarkSenderCommittedInput) error {
	if in.SenderPTS < math.MinInt32 || in.SenderPTS > math.MaxInt32 {
		return fmt.Errorf("%w: sender pts out of int32 range", msg.ErrSenderSyncFailed)
	}
	if _, err := r.requireDB(); err != nil {
		return err
	}
	state, err := r.selectSendStateByID(ctx, in.SendStateID)
	if err != nil {
		return err
	}
	if state.PeerSeq < math.MinInt32 || state.PeerSeq > math.MaxInt32 {
		return fmt.Errorf("%w: peer seq out of int32 range", msg.ErrSenderSyncFailed)
	}
	affected, err := r.models.MessageSendStatesModel.MarkSenderCommitted(ctx, in.SenderOperationID, in.SenderPTS, in.SenderPTSCount, in.SenderUpdateSchemaVersion, in.SenderUpdatePayload, in.SenderUpdatePayloadHash, SendStateStatusSenderCommitted, in.SendStateID)
	if err != nil {
		return storageError("mark sender committed", err)
	}
	return requireAffectedRows(affected, "mark sender committed")
}

func (r *Repository) MarkReceiverOpsAcked(ctx context.Context, sendStateID int64, receiverManifestID int64) error {
	if _, err := r.requireDB(); err != nil {
		return err
	}
	affected, err := r.models.MessageSendStatesModel.MarkReceiverOpsAcked(ctx, receiverManifestID, SendStateStatusReceiverAcked, sendStateID)
	if err != nil {
		return storageError("mark receiver ops acked", err)
	}
	return requireAffectedRows(affected, "mark receiver ops acked")
}

func (r *Repository) MarkCompleted(ctx context.Context, sendStateID int64) error {
	if _, err := r.requireDB(); err != nil {
		return err
	}
	affected, err := r.models.MessageSendStatesModel.MarkCompleted(ctx, SendStateStatusCompleted, mysqlNow(), sendStateID)
	if err != nil {
		return storageError("mark completed", err)
	}
	return requireAffectedRows(affected, "mark completed")
}

func (r *Repository) MarkRetryableFailure(ctx context.Context, in MarkRetryableFailureInput) error {
	if _, err := r.requireDB(); err != nil {
		return err
	}
	affected, err := r.models.MessageSendStatesModel.MarkRetryableFailure(ctx, SendStateStatusFailedRetryable, in.LastErrorCategory, in.LastErrorCode, in.LastErrorMessage, in.SendStateID)
	if err != nil {
		return storageError("mark retryable failure", err)
	}
	return requireAffectedRows(affected, "mark retryable failure")
}

func (r *Repository) selectSendStateByRandom(ctx context.Context, senderUserID int64, peerType int32, peerID int64, clientRandomID int64) (*SendState, bool, error) {
	row, err := r.models.MessageSendStatesModel.SelectByRandom(ctx, senderUserID, peerType, peerID, clientRandomID)
	if err != nil {
		if errors.Is(err, model.ErrNotFound) {
			return nil, false, nil
		}
		return nil, false, storageError("select send state by random", err)
	}
	return sendStateFromModel(row), true, nil
}

func (r *Repository) selectSendStateByID(ctx context.Context, sendStateID int64) (*SendState, error) {
	row, err := r.models.MessageSendStatesModel.SelectBySendStateId(ctx, sendStateID)
	if err != nil {
		return nil, storageError("select send state by id", err)
	}
	return sendStateFromModel(row), nil
}

func selectSendStateByIDTx(txModels *model.TxModels, sendStateID int64) (*SendState, error) {
	row, err := txModels.MessageSendStatesModel.SelectBySendStateId(sendStateID)
	if err != nil {
		return nil, storageError("select send state by id for update", err)
	}
	return sendStateFromModel(row), nil
}

func requireAffectedRows(affected int64, _ string) error {
	if affected == 0 {
		return msg.ErrSendStateConflict
	}
	return nil
}

func sendStateFromModel(r *model.MessageSendStates) *SendState {
	return &SendState{
		SendStateID:                 r.SendStateId,
		SenderUserID:                r.SenderUserId,
		PeerType:                    r.PeerType,
		PeerID:                      r.PeerId,
		ClientRandomID:              r.ClientRandomId,
		CanonicalMessageID:          r.CanonicalMessageId,
		PeerSeq:                     r.PeerSeq,
		Status:                      r.Status,
		RequestPayloadSchemaVersion: r.RequestPayloadSchemaVersion,
		RequestPayloadHash:          r.RequestPayloadHash,
		SenderOperationID:           r.SenderOperationId,
		SenderPTS:                   r.SenderPts,
		SenderPTSCount:              r.SenderPtsCount,
		SenderUpdateSchemaVersion:   r.SenderUpdateSchemaVersion,
		SenderUpdatePayload:         r.SenderUpdatePayload,
		SenderUpdatePayloadHash:     r.SenderUpdatePayloadHash,
		ReceiverManifestID:          r.ReceiverManifestId,
		RetryCount:                  r.RetryCount,
	}
}
