package core

import "github.com/teamgram/teamgram-server/v2/pkg/proto/tg"

const (
	dialogPeerTypeUser    int32 = 1
	dialogPeerTypeChat    int32 = 2
	dialogPeerTypeChannel int32 = 3
)

func (c *DialogsCore) resolveInputDialogPeer(peer tg.InputDialogPeerClazz) (tg.PeerUtilClazz, error) {
	inputDialogPeer, ok := (&tg.InputDialogPeer{Clazz: peer}).ToInputDialogPeer()
	if !ok {
		if _, isFolder := (&tg.InputDialogPeer{Clazz: peer}).ToInputDialogPeerFolder(); isFolder {
			return nil, tg.ErrFolderIdInvalid
		}
		return nil, tg.ErrInputConstructorInvalid
	}
	resolved := tg.FromInputPeer2(c.MD.UserId, inputDialogPeer.Peer)
	switch resolved.PeerType {
	case tg.PEER_SELF:
		resolved.PeerType = tg.PEER_USER
		resolved.PeerId = c.MD.UserId
	case tg.PEER_USER, tg.PEER_CHAT, tg.PEER_CHANNEL:
	default:
		return nil, tg.Err400PeerIdInvalid
	}
	if resolved.PeerId <= 0 {
		return nil, tg.Err400PeerIdInvalid
	}
	return resolved, nil
}

func dialogFacadePeerType(peerType int32) (int32, error) {
	switch peerType {
	case tg.PEER_SELF, tg.PEER_USER:
		return dialogPeerTypeUser, nil
	case tg.PEER_CHAT:
		return dialogPeerTypeChat, nil
	case tg.PEER_CHANNEL:
		return dialogPeerTypeChannel, nil
	default:
		return 0, tg.Err400PeerIdInvalid
	}
}

func dialogFacadePeerDialogID(peerType int32, peerID int64) (int64, error) {
	facadePeerType, err := dialogFacadePeerType(peerType)
	if err != nil {
		return 0, err
	}
	if peerID <= 0 {
		return 0, tg.Err400PeerIdInvalid
	}
	return peerID*16 + int64(facadePeerType), nil
}

func makePublicPeerFromDialogFacade(peerType int32, peerID int64) tg.PeerClazz {
	switch peerType {
	case dialogPeerTypeChat:
		return tg.MakeTLPeerChat(&tg.TLPeerChat{ChatId: peerID})
	case dialogPeerTypeChannel:
		return tg.MakeTLPeerChannel(&tg.TLPeerChannel{ChannelId: peerID})
	default:
		return tg.MakeTLPeerUser(&tg.TLPeerUser{UserId: peerID})
	}
}
