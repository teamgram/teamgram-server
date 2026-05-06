package repository

import (
	"context"
	"errors"
	"time"

	"github.com/teamgram/marmota/pkg/stores/sqlx"
	"github.com/teamgram/teamgram-server/v2/app/messenger/userupdates/internal/repository/model"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
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

func (r *Repository) GetOutboxReadDate(ctx context.Context, in OutboxReadDateInput) (int32, error) {
	if _, err := r.requireDB(); err != nil {
		return 0, err
	}
	var exists int32
	existsQuery := `
SELECT
	1
FROM
	user_message_views
WHERE
	user_id = ?
	AND peer_type = ?
	AND peer_id = ?
	AND peer_seq = ?
	AND outgoing = 1
	AND message_status = ?
LIMIT 1`
	if err := r.db.QueryRow(ctx, &exists, existsQuery, in.UserID, in.PeerType, in.PeerID, in.MsgID, MessageStatusLive); err != nil {
		if errors.Is(err, sqlx.ErrNotFound) || errors.Is(err, model.ErrNotFound) {
			return 0, tg.ErrMessageIdInvalid
		}
		return 0, storageError("validate outbox read message", err)
	}

	query := `
SELECT
	read_outbox_max_date
FROM
	message_read_outbox
WHERE
	user_id = ?
	AND peer_type = ?
	AND peer_id = ?
	AND read_user_id = ?
	AND read_outbox_max_id >= ?
ORDER BY
	read_outbox_max_id ASC
LIMIT 1`
	var readDate time.Time
	if err := r.db.QueryRow(ctx, &readDate, query, in.UserID, in.PeerType, in.PeerID, in.PeerID, in.MsgID); err != nil {
		if errors.Is(err, sqlx.ErrNotFound) || errors.Is(err, model.ErrNotFound) {
			return 0, tg.ErrMessageNotReadYet
		}
		return 0, storageError("get outbox read date", err)
	}
	return int32(readDate.Unix()), nil
}
