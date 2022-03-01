// Copyright (c) 2021-present,  Teamgram Studio (https://teamgram.io).
//  All rights reserved.
//
// Author: teamgramio (teamgram.io@gmail.com)
//

package chat

import (
	"github.com/teamgram/proto/mtproto"
)

func (m *ChatInviteExt) ToChatInvite(selfId int64, cb func(idList []int64) []*mtproto.User) *mtproto.ChatInvite {
	switch m.GetPredicateName() {
	case Predicate_chatInviteAlready:
		return mtproto.MakeTLChatInviteAlready(&mtproto.ChatInvite{
			Chat: m.Chat.ToUnsafeChat(selfId),
		}).To_ChatInvite()
	case Predicate_chatInvite:
		chatInvite := mtproto.MakeTLChatInvite(&mtproto.ChatInvite{
			Channel:           false,
			Broadcast:         false,
			Public:            false,
			Megagroup:         false,
			RequestNeeded:     m.RequestNeeded,
			Title:             m.Title,
			About:             m.About,
			Photo:             m.Photo,
			ParticipantsCount: m.ParticipantsCount,
			Participants:      []*mtproto.User{},
		}).To_ChatInvite()
		if cb != nil {
			chatInvite.Participants = cb(m.Participants)
		}
		return chatInvite
	case Predicate_chatInvitePeek:
		return mtproto.MakeTLChatInvitePeek(&mtproto.ChatInvite{
			Chat:    m.Chat.ToUnsafeChat(selfId),
			Expires: m.Expires,
		}).To_ChatInvite()
	default:
		return nil
	}
}
