package repository

import (
	"context"
	"errors"

	"github.com/teamgram/teamgram-server/v2/app/messenger/userupdates/internal/repository/model"
	"github.com/teamgram/teamgram-server/v2/pkg/paging"
)

type DialogProjectionPeer struct {
	PeerType int32
	PeerID   int64
}

type DialogProjectionCursor struct {
	TopMessageDate int64
	TopPeerSeq     int64
	PeerType       int32
	PeerID         int64
}

type DialogProjection struct {
	UserID                     int64
	PeerType                   int32
	PeerID                     int64
	TopPeerSeq                 int64
	TopUserMessageID           int64
	TopCanonicalMessageID      int64
	TopMessageDate             int64
	TopMessageStatus           int32
	ReadInboxMaxPeerSeq        int64
	ReadInboxMaxUserMessageID  int64
	ReadOutboxMaxPeerSeq       int64
	ReadOutboxMaxUserMessageID int64
	UnreadCount                int32
	UnreadMentionsCount        int32
	UnreadReactionsCount       int32
	UnreadMark                 bool
	PinnedPeerSeq              int64
	PinnedUserMessageID        int64
	PinnedCanonicalMessageID   int64
	HasScheduled               bool
	AvailableMinPeerSeq        int64
	AvailableMinUserMessageID  int64
	LastPTS                    int64
	LastPTSAt                  int64
	DialogSchemaVersion        int32
	DialogPayload              []byte
}

func (r *Repository) ListDialogProjections(ctx context.Context, userID int64, cursor DialogProjectionCursor, limit int32) ([]DialogProjection, error) {
	if _, err := r.requireDB(); err != nil {
		return nil, err
	}
	limit = paging.NormalizeDialogLimit(limit)
	if limit == 0 {
		return []DialogProjection{}, nil
	}
	rows, err := r.models.UserDialogsModel.SelectByUserCursor(ctx, userID, cursor.TopMessageDate, cursor.TopPeerSeq, cursor.PeerType, cursor.PeerID, limit)
	if err != nil {
		if errors.Is(err, model.ErrNotFound) {
			return []DialogProjection{}, nil
		}
		return nil, storageError("list dialog projections", err)
	}
	return mapDialogProjectionRows(rows), nil
}

func (r *Repository) GetDialogProjectionsByPeers(ctx context.Context, userID int64, peers []DialogProjectionPeer) (map[DialogProjectionPeer]DialogProjection, error) {
	if _, err := r.requireDB(); err != nil {
		return nil, err
	}
	out := make(map[DialogProjectionPeer]DialogProjection, len(peers))
	if len(peers) == 0 {
		return out, nil
	}

	requested := make(map[DialogProjectionPeer]struct{}, len(peers))
	peerIDSet := make(map[int64]struct{}, len(peers))
	peerIDs := make([]int64, 0, len(peers))
	for _, peer := range peers {
		requested[peer] = struct{}{}
		if _, ok := peerIDSet[peer.PeerID]; ok {
			continue
		}
		peerIDSet[peer.PeerID] = struct{}{}
		peerIDs = append(peerIDs, peer.PeerID)
	}

	rows, err := r.models.UserDialogsModel.SelectByUserPeers(ctx, userID, peerIDs)
	if err != nil {
		if errors.Is(err, model.ErrNotFound) {
			return out, nil
		}
		return nil, storageError("get dialog projections by peers", err)
	}
	for _, row := range rows {
		peer := DialogProjectionPeer{PeerType: row.PeerType, PeerID: row.PeerId}
		if _, ok := requested[peer]; !ok {
			continue
		}
		out[peer] = mapDialogProjectionRow(row)
	}
	return out, nil
}

func (r *Repository) CountVisibleDialogs(ctx context.Context, userID int64) (int32, error) {
	if _, err := r.requireDB(); err != nil {
		return 0, err
	}
	row, err := r.models.UserupdatesQueries.CountVisibleDialogs(ctx, userID)
	if err != nil {
		return 0, storageError("count visible dialogs", err)
	}
	return row.Count, nil
}

func mapDialogProjectionRows(rows []model.UserDialogs) []DialogProjection {
	out := make([]DialogProjection, 0, len(rows))
	for _, row := range rows {
		out = append(out, mapDialogProjectionRow(row))
	}
	return out
}

func mapDialogProjectionRow(row model.UserDialogs) DialogProjection {
	return DialogProjection{
		UserID:                     row.UserId,
		PeerType:                   row.PeerType,
		PeerID:                     row.PeerId,
		TopPeerSeq:                 row.TopPeerSeq,
		TopUserMessageID:           row.TopUserMessageId,
		TopCanonicalMessageID:      row.TopCanonicalMessageId,
		TopMessageDate:             row.TopMessageDate,
		TopMessageStatus:           row.TopMessageStatus,
		ReadInboxMaxPeerSeq:        row.ReadInboxMaxPeerSeq,
		ReadInboxMaxUserMessageID:  row.ReadInboxMaxUserMessageId,
		ReadOutboxMaxPeerSeq:       row.ReadOutboxMaxPeerSeq,
		ReadOutboxMaxUserMessageID: row.ReadOutboxMaxUserMessageId,
		UnreadCount:                row.UnreadCount,
		UnreadMentionsCount:        row.UnreadMentionsCount,
		UnreadReactionsCount:       row.UnreadReactionsCount,
		UnreadMark:                 row.UnreadMark,
		PinnedPeerSeq:              row.PinnedPeerSeq,
		PinnedUserMessageID:        row.PinnedUserMessageId,
		PinnedCanonicalMessageID:   row.PinnedCanonicalMessageId,
		HasScheduled:               row.HasScheduled,
		AvailableMinPeerSeq:        row.AvailableMinPeerSeq,
		AvailableMinUserMessageID:  row.AvailableMinUserMessageId,
		LastPTS:                    row.LastPts,
		LastPTSAt:                  row.LastPtsAt,
		DialogSchemaVersion:        row.DialogSchemaVersion,
		DialogPayload:              row.DialogPayload,
	}
}
