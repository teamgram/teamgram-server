package repository

import (
	"context"
	"errors"

	"github.com/teamgram/marmota/pkg/stores/sqlx"
	"github.com/teamgram/teamgram-server/v2/app/messenger/msg/internal/repository/model"
)

type ResolvedMessageID struct {
	UserID             int64
	PeerType           int32
	PeerID             int64
	UserMessageID      int64
	PeerSeq            int64
	CanonicalMessageID int64
	MessageDate        int64
	Outgoing           bool
}

type HistoryCursorBounds struct {
	OffsetPeerSeq int64
	MaxPeerSeq    int64
	MinPeerSeq    int64
	NoMatch       bool
}

func (r *Repository) ResolveMessageID(ctx context.Context, userID int64, peerType int32, peerID int64, userMessageID int64) (*ResolvedMessageID, error) {
	if userMessageID <= 0 {
		return nil, nil
	}
	if _, err := r.requireDB(); err != nil {
		return nil, err
	}
	row, err := r.models.CanonicalQueries.SelectUserMessageByID(ctx, userID, peerType, peerID, userMessageID, MessageStatusLive)
	if err != nil {
		if isModelNotFound(err) {
			return nil, nil
		}
		return nil, storageError("resolve message id", err)
	}
	return resolvedMessageIDFromRow(row), nil
}

func (r *Repository) ResolvePeerSeqToUserMessageID(ctx context.Context, userID int64, peerType int32, peerID int64, peerSeq int64) (int64, error) {
	if peerSeq <= 0 {
		return 0, nil
	}
	if _, err := r.requireDB(); err != nil {
		return 0, err
	}
	row, err := r.models.CanonicalQueries.SelectNearestLiveUserMessageByPeerSeq(ctx, userID, peerType, peerID, peerSeq, MessageStatusLive)
	if err != nil {
		if isModelNotFound(err) {
			return 0, nil
		}
		return 0, storageError("resolve peer seq to user message id", err)
	}
	return row.UserMessageID, nil
}

func (r *Repository) ResolveHistoryCursorIDs(ctx context.Context, userID int64, peerType int32, peerID int64, offsetID int32, maxID int32, minID int32) (HistoryCursorBounds, error) {
	var out HistoryCursorBounds
	if offsetID > 0 {
		row, err := r.ResolveMessageID(ctx, userID, peerType, peerID, int64(offsetID))
		if err != nil {
			return out, err
		}
		if row == nil {
			out.NoMatch = true
			return out, nil
		}
		out.OffsetPeerSeq = row.PeerSeq
	}
	if maxID > 0 {
		row, err := r.ResolveMessageID(ctx, userID, peerType, peerID, int64(maxID))
		if err != nil {
			return out, err
		}
		if row == nil {
			out.NoMatch = true
			return out, nil
		}
		out.MaxPeerSeq = row.PeerSeq
	}
	if minID > 0 {
		row, err := r.ResolveMessageID(ctx, userID, peerType, peerID, int64(minID))
		if err != nil {
			return out, err
		}
		if row == nil {
			out.NoMatch = true
			return out, nil
		}
		out.MinPeerSeq = row.PeerSeq
	}
	return out, nil
}

func resolvedMessageIDFromRow(row *model.ResolvedMessageIDRow) *ResolvedMessageID {
	if row == nil {
		return nil
	}
	return &ResolvedMessageID{
		UserID:             row.UserID,
		PeerType:           row.PeerType,
		PeerID:             row.PeerID,
		UserMessageID:      row.UserMessageID,
		PeerSeq:            row.PeerSeq,
		CanonicalMessageID: row.CanonicalMessageID,
		MessageDate:        row.MessageDate,
		Outgoing:           row.Outgoing,
	}
}

func isModelNotFound(err error) bool {
	return errors.Is(err, model.ErrNotFound) || errors.Is(err, sqlx.ErrNotFound)
}
