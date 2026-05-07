package repository

import (
	"context"
	"errors"

	"github.com/teamgram/teamgram-server/v2/app/messenger/userupdates/internal/repository/model"
	"github.com/teamgram/teamgram-server/v2/app/messenger/userupdates/userupdates"
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
	Date               int64
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
		Date:               row.Date,
		ViewSchemaVersion:  row.ViewSchemaVersion,
		ViewPayload:        row.ViewPayload,
	}
}

func (r *Repository) GetOutboxReadDate(ctx context.Context, in OutboxReadDateInput) (int64, error) {
	if _, err := r.requireDB(); err != nil {
		return 0, err
	}

	view, err := r.models.UserMessageViewsModel.SelectByUserPeerSeq(ctx, in.UserID, in.PeerType, in.PeerID, int64(in.MsgID))
	if err != nil {
		if errors.Is(err, model.ErrNotFound) {
			return 0, userupdates.ErrOutboxReadMessageInvalid
		}
		return 0, storageError("validate outbox read message", err)
	}
	if !view.Outgoing || view.MessageStatus != MessageStatusLive {
		return 0, userupdates.ErrOutboxReadMessageInvalid
	}

	rows, err := r.models.MessageReadOutboxModel.SelectFirstReadDate(ctx, in.UserID, in.PeerType, in.PeerID, in.PeerID, int64(in.MsgID), 1)
	if err != nil {
		return 0, storageError("get outbox read date", err)
	}
	if len(rows) == 0 {
		return 0, userupdates.ErrOutboxReadDateNotFound
	}
	return rows[0].ReadOutboxMaxDate, nil
}
