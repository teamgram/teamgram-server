package core

import "github.com/teamgram/teamgram-server/v2/pkg/proto/tg"

func makePlaceholderUpdatesState(pts int32, date int32) *tg.UpdatesState {
	return tg.MakeTLUpdatesState(&tg.TLUpdatesState{
		Pts:         pts,
		Qts:         0,
		Date:        date,
		Seq:         0,
		UnreadCount: 0,
	}).ToUpdatesState()
}

func makePlaceholderBFFDifferenceMessage(messageID int32, date int32) tg.MessageClazz {
	return tg.MakeTLMessage(&tg.TLMessage{
		Out:     true,
		Id:      messageID,
		FromId:  tg.MakeTLPeerUser(&tg.TLPeerUser{UserId: 0}),
		PeerId:  tg.MakeTLPeerUser(&tg.TLPeerUser{UserId: 0}),
		Date:    date,
		Message: "placeholder",
	})
}
