package core

import (
	"time"

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
			TopMessage: int32(record.TopPeerSeq),
		}))
	}
	return dialog.MakeTLSavedDialogList(&dialog.TLSavedDialogList{
		Count:   int32(len(records)),
		Dialogs: dialogs,
	})
}

func unixOffsetDate(offsetDate int32) time.Time {
	if offsetDate <= 0 {
		return time.Time{}
	}
	return time.Unix(int64(offsetDate), 0).UTC()
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
