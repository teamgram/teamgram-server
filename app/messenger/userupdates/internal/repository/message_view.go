package repository

import (
	"context"
	"errors"

	"github.com/teamgram/teamgram-server/v2/app/messenger/userupdates/internal/repository/model"
)

type MessageViewPeerSeq struct {
	PeerType int32
	PeerID   int64
	PeerSeq  int64
}

type MessageView struct {
	UserID             int64
	PeerType           int32
	PeerID             int64
	PeerSeq            int64
	CanonicalMessageID int64
	FromUserID         int64
	Outgoing           bool
	MessageKind        int32
	MessageStatus      int32
	EditVersion        int32
	Date               int32
	ViewSchemaVersion  int32
	ViewPayload        []byte
}

func (r *Repository) GetMessageViewsByPeerSeqs(ctx context.Context, userID int64, peers []MessageViewPeerSeq) (map[MessageViewPeerSeq]MessageView, error) {
	if _, err := r.requireDB(); err != nil {
		return nil, err
	}
	out := make(map[MessageViewPeerSeq]MessageView, len(peers))
	for _, peer := range peers {
		if peer.PeerType == 0 || peer.PeerID == 0 || peer.PeerSeq == 0 {
			continue
		}
		if _, ok := out[peer]; ok {
			continue
		}
		row, err := r.models.UserMessageViewsModel.SelectByUserPeerSeq(ctx, userID, peer.PeerType, peer.PeerID, peer.PeerSeq)
		if err != nil {
			if errors.Is(err, model.ErrNotFound) {
				continue
			}
			return nil, storageError("get message view by peer seq", err)
		}
		out[peer] = mapMessageViewRow(row)
	}
	return out, nil
}

func mapMessageViewRow(row *model.UserMessageViews) MessageView {
	if row == nil {
		return MessageView{}
	}
	return MessageView{
		UserID:             row.UserId,
		PeerType:           row.PeerType,
		PeerID:             row.PeerId,
		PeerSeq:            row.PeerSeq,
		CanonicalMessageID: row.CanonicalMessageId,
		FromUserID:         row.FromUserId,
		Outgoing:           row.Outgoing,
		MessageKind:        row.MessageKind,
		MessageStatus:      row.MessageStatus,
		EditVersion:        row.EditVersion,
		Date:               int32(row.Date.Unix()),
		ViewSchemaVersion:  row.ViewSchemaVersion,
		ViewPayload:        row.ViewPayload,
	}
}
