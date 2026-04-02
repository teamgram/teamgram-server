package core

import "github.com/teamgram/teamgram-server/v2/pkg/proto/tg"

func makeBffAffectedMessagesPlaceholder(pts int32, ptsCount int32) *tg.MessagesAffectedMessages {
	if pts <= 0 {
		pts = 1
	}
	if ptsCount <= 0 {
		ptsCount = 1
	}

	return tg.MakeTLMessagesAffectedMessages(&tg.TLMessagesAffectedMessages{
		Pts:      pts,
		PtsCount: ptsCount,
	}).ToMessagesAffectedMessages()
}

func makeBffAffectedHistoryPlaceholder(pts int32, offset int32) *tg.MessagesAffectedHistory {
	if pts <= 0 {
		pts = 1
	}

	return tg.MakeTLMessagesAffectedHistory(&tg.TLMessagesAffectedHistory{
		Pts:      pts,
		PtsCount: 1,
		Offset:   offset,
	}).ToMessagesAffectedHistory()
}

func bffPeerFromInput(c *MessagesCore, input tg.InputPeerClazz) (tg.PeerClazz, error) {
	userID := int64(0)
	if c != nil && c.MD != nil {
		userID = c.MD.UserId
	}

	peer := tg.FromInputPeer2(userID, input)
	switch peer.PeerType {
	case tg.PEER_SELF:
		return tg.MakeTLPeerUser(&tg.TLPeerUser{UserId: userID}), nil
	case tg.PEER_USER:
		return tg.MakeTLPeerUser(&tg.TLPeerUser{UserId: peer.PeerId}), nil
	case tg.PEER_CHAT:
		return tg.MakeTLPeerChat(&tg.TLPeerChat{ChatId: peer.PeerId}), nil
	case tg.PEER_CHANNEL:
		return nil, tg.ErrEnterpriseIsBlocked
	default:
		return nil, tg.ErrPeerIdInvalid
	}
}
