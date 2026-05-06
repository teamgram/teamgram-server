package core

import "github.com/teamgram/teamgram-server/v2/pkg/proto/tg"

const (
	dialogPeerTypeUser    int32 = 1
	dialogPeerTypeChat    int32 = 2
	dialogPeerTypeChannel int32 = 3
)

type draftDialogPeer struct {
	PeerType int32
	PeerID   int64
}

func resolveDraftDialogPeer(selfUserID int64, peer tg.InputPeerClazz) (draftDialogPeer, error) {
	resolved := tg.FromInputPeer2(selfUserID, peer)
	switch resolved.PeerType {
	case tg.PEER_SELF:
		resolved.PeerType = tg.PEER_USER
		resolved.PeerId = selfUserID
	case tg.PEER_USER, tg.PEER_CHAT, tg.PEER_CHANNEL:
	default:
		return draftDialogPeer{}, tg.Err400PeerIdInvalid
	}
	if resolved.PeerId <= 0 {
		return draftDialogPeer{}, tg.Err400PeerIdInvalid
	}

	switch resolved.PeerType {
	case tg.PEER_USER:
		return draftDialogPeer{PeerType: dialogPeerTypeUser, PeerID: resolved.PeerId}, nil
	case tg.PEER_CHAT:
		return draftDialogPeer{PeerType: dialogPeerTypeChat, PeerID: resolved.PeerId}, nil
	case tg.PEER_CHANNEL:
		return draftDialogPeer{PeerType: dialogPeerTypeChannel, PeerID: resolved.PeerId}, nil
	default:
		return draftDialogPeer{}, tg.Err400PeerIdInvalid
	}
}
