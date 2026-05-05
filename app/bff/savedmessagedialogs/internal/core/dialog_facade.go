package core

import (
	dialogpb "github.com/teamgram/teamgram-server/v2/app/service/biz/dialog/dialog"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

const (
	dialogFacadePeerTypeUser    int32 = 1
	dialogFacadePeerTypeChat    int32 = 2
	dialogFacadePeerTypeChannel int32 = 3
)

func (c *SavedMessageDialogsCore) resolveInputDialogPeer(peer tg.InputDialogPeerClazz) (*tg.TLPeerUtil, error) {
	inputDialogPeer, ok := (&tg.InputDialogPeer{Clazz: peer}).ToInputDialogPeer()
	if !ok {
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

func savedDialogFacadePeer(peerType int32) (int32, error) {
	switch peerType {
	case tg.PEER_SELF, tg.PEER_USER:
		return dialogFacadePeerTypeUser, nil
	case tg.PEER_CHAT:
		return dialogFacadePeerTypeChat, nil
	case tg.PEER_CHANNEL:
		return dialogFacadePeerTypeChannel, nil
	default:
		return 0, tg.Err400PeerIdInvalid
	}
}

func savedDialogFacadePeerUtil(peer *tg.TLPeerUtil) (*tg.TLPeerUtil, error) {
	peerType, err := savedDialogFacadePeer(peer.PeerType)
	if err != nil {
		return nil, err
	}
	return tg.MakeTLPeerUtil(&tg.TLPeerUtil{PeerType: peerType, PeerId: peer.PeerId}), nil
}

func makeMessagesSavedDialogs(in *dialogpb.SavedDialogList) *tg.MessagesSavedDialogs {
	dialogs := []tg.SavedDialogClazz{}
	if in != nil {
		dialogs = append(dialogs, in.Dialogs...)
	}
	return tg.MakeTLMessagesSavedDialogs(&tg.TLMessagesSavedDialogs{
		Dialogs:  dialogs,
		Messages: []tg.MessageClazz{},
		Chats:    []tg.ChatClazz{},
		Users:    []tg.UserClazz{},
	}).ToMessagesSavedDialogs()
}
