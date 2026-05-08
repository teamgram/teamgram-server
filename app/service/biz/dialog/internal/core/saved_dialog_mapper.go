package core

import (
	"encoding/json"
	"math"

	"github.com/teamgram/teamgram-server/v2/app/service/biz/dialog/dialog"
	"github.com/teamgram/teamgram-server/v2/app/service/biz/dialog/internal/repository"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

func makeSavedDialogList(records []repository.SavedDialogRecord) *dialog.SavedDialogList {
	dialogs := make([]tg.SavedDialogClazz, 0, len(records))
	for _, record := range records {
		dialogs = append(dialogs, tg.MakeTLSavedDialog(&tg.TLSavedDialog{
			Pinned:     record.Pinned,
			Peer:       makeSavedDialogPeer(record.PeerType, record.PeerID),
			TopMessage: savedDialogTopMessageID(record),
		}))
	}
	return dialog.MakeTLSavedDialogList(&dialog.TLSavedDialogList{
		Count:   int32(len(records)),
		Dialogs: dialogs,
	})
}

type savedDialogPayloadV1 struct {
	TopUserMessageID int64 `json:"top_user_message_id,omitempty"`
}

func savedDialogTopMessageID(record repository.SavedDialogRecord) int32 {
	if len(record.SavedPayload) == 0 {
		return 0
	}
	var payload savedDialogPayloadV1
	if err := json.Unmarshal(record.SavedPayload, &payload); err != nil {
		return 0
	}
	if payload.TopUserMessageID <= 0 || payload.TopUserMessageID > math.MaxInt32 {
		return 0
	}
	return int32(payload.TopUserMessageID)
}

func unixOffsetDate(offsetDate int32) int64 {
	if offsetDate <= 0 {
		return 0
	}
	return int64(offsetDate)
}

func dialogDateInt32FromUnixSeconds(seconds int64, field string) (int32, error) {
	date, err := tg.DateInt32FromUnixSeconds(seconds)
	if err != nil {
		return 0, dialog.WrapDialogStorage("convert "+field, err)
	}
	return date, nil
}

func peerRefsFromPeerUtils(peers []tg.PeerUtilClazz) []repository.PeerRef {
	out := make([]repository.PeerRef, 0, len(peers))
	for _, peer := range peers {
		if peer == nil {
			continue
		}
		out = append(out, repository.PeerRef{PeerType: peer.PeerType, PeerID: peer.PeerId})
	}
	return out
}

func makeSavedDialogPeer(peerType int32, peerID int64) tg.PeerClazz {
	switch peerType {
	case repository.PeerTypeChat:
		return tg.MakePeerChat(peerID)
	case repository.PeerTypeChannel:
		return tg.MakePeerChannel(peerID)
	default:
		return tg.MakePeerUser(peerID)
	}
}
