package core

import (
	"time"

	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

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

func historyPlaceholderStartID(offsetID, maxID, minID int32) int32 {
	switch {
	case offsetID > 0:
		return offsetID
	case maxID > 0:
		return maxID
	case minID > 0:
		return minID
	default:
		return 1
	}
}

func historyPlaceholderCount(limit int32) int {
	switch {
	case limit <= 0:
		return 0
	case limit > 3:
		return 3
	default:
		return int(limit)
	}
}

func makeBffMessagesMessagesPlaceholder(peer tg.PeerClazz, startID int32, count int, mentioned bool) *tg.MessagesMessages {
	ids := make([]int32, 0, count)
	for i := 0; i < count; i++ {
		id := startID + int32(i)
		if id <= 0 {
			id = int32(i + 1)
		}
		ids = append(ids, id)
	}
	return makeBffMessagesMessagesByIDs(peer, ids, mentioned)
}

func makeBffMessagesMessagesByIDs(peer tg.PeerClazz, ids []int32, mentioned bool) *tg.MessagesMessages {
	messages := make([]tg.MessageClazz, 0, len(ids))
	now := int32(time.Now().Unix())
	for _, id := range ids {
		if id <= 0 {
			id = 1
		}
		messages = append(messages, tg.MakeTLMessage(&tg.TLMessage{
			Id:        id,
			Out:       true,
			Mentioned: mentioned,
			Date:      now,
			Message:   "placeholder",
			PeerId:    peer,
		}))
	}

	return tg.MakeTLMessagesMessages(&tg.TLMessagesMessages{
		Messages: messages,
	}).ToMessagesMessages()
}

func makeBffMessagesMessagesFromBoxes(boxes []tg.MessageBoxClazz) *tg.MessagesMessages {
	messages := make([]tg.MessageClazz, 0, len(boxes))
	for _, box := range boxes {
		if box != nil && box.Message != nil {
			messages = append(messages, box.Message)
		}
	}
	return tg.MakeTLMessagesMessages(&tg.TLMessagesMessages{
		Messages: messages,
	}).ToMessagesMessages()
}

func bffPeerTypeAndID(peer tg.PeerClazz) (int32, int64, bool) {
	switch p := peer.(type) {
	case *tg.TLPeerUser:
		return tg.PEER_USER, p.UserId, true
	case *tg.TLPeerChat:
		return tg.PEER_CHAT, p.ChatId, true
	case *tg.TLPeerChannel:
		return tg.PEER_CHANNEL, p.ChannelId, true
	default:
		return 0, 0, false
	}
}
