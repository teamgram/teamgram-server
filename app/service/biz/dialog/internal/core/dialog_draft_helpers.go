package core

import "github.com/teamgram/teamgram-server/v2/pkg/proto/tg"

func tgPeer(peerType int32, peerID int64) tg.PeerClazz {
	return tg.MakePeerHelper(peerType, peerID)
}

func tgDraftEmpty() tg.DraftMessageClazz {
	return tg.MakeTLDraftMessageEmpty(&tg.TLDraftMessageEmpty{})
}
