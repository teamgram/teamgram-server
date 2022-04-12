// Copyright (c) 2021-present,  Teamgram Studio (https://teamgram.io).
//  All rights reserved.
//
// Author: teamgramio (teamgram.io@gmail.com)
//

package chat

import (
	"github.com/teamgram/proto/mtproto"
	"time"
)

func (m *MutableChat) Id() int64 {
	return m.Chat.GetId()
}

func (m *MutableChat) Creator() int64 {
	return m.GetChat().GetCreator()
}

func (m *MutableChat) Title() string {
	return m.GetChat().GetTitle()
}

func (m *MutableChat) About() string {
	return m.GetChat().GetAbout()
}

func (m *MutableChat) ChatPhoto() *mtproto.ChatPhoto {
	return mtproto.MakeChatPhotoByPhoto(m.GetChat().GetPhoto())
}

func (m *MutableChat) Photo() *mtproto.Photo {
	return m.GetChat().GetPhoto()
}

func (m *MutableChat) ParticipantsCount() int32 {
	return m.GetChat().GetParticipantsCount()
}

func (m *MutableChat) Date() int32 {
	return int32(m.GetChat().GetDate())
}

func (m *MutableChat) Version() int32 {
	return m.GetChat().GetVersion()
}

func (m *MutableChat) MigratedTo() *mtproto.InputChannel {
	return m.GetChat().GetMigratedTo()
}

func (m *MutableChat) DefaultBannedRights() *mtproto.ChatBannedRights {
	return m.GetChat().GetDefaultBannedRights()
}

func (m *MutableChat) Deactivated() bool {
	return m.GetChat().GetDeactivated()
}

func (m *MutableChat) CallActive() bool {
	return m.GetChat().GetCallActive()
}

func (m *MutableChat) CallNotEmpty() bool {
	return m.GetChat().GetCallNotEmpty()
}

func (m *MutableChat) Noforwards() bool {
	return m.GetChat().GetNoforwards()
}

func (m *MutableChat) Call() *mtproto.InputGroupCall {
	return m.GetChat().GetCall()
}

func (m *MutableChat) AvailableReactions() []string {
	reactions := m.GetChat().GetAvailableReactions()
	if len(reactions) == 0 {
		return nil
	} else {
		return reactions
	}
}

func (m *MutableChat) ParticipantIdList() []int64 {
	var (
		idList []int64
	)

	m.Walk(func(userId int64, participant *ImmutableChatParticipant) error {
		if participant.IsChatMemberStateNormal() {
			idList = append(idList, userId)
		}
		return nil
	})

	return idList
}

func (m *MutableChat) MakeMessageService(fromId int64, action *mtproto.MessageAction) *mtproto.Message {
	// messageService#2b085862 flags:#
	//	out:flags.1?true
	//	mentioned:flags.4?true
	//	media_unread:flags.5?true
	//	silent:flags.13?true
	//	post:flags.14?true
	//	legacy:flags.19?true
	//	id:int
	//	from_id:flags.8?Peer
	//	peer_id:Peer
	//	reply_to:flags.3?MessageReplyHeader
	//	date:int
	//	action:MessageAction
	//	ttl_period:flags.25?int = Message;
	//
	message := mtproto.MakeTLMessageService(&mtproto.Message{
		// TODO(@benqi): fill it
		Out:         true,
		Mentioned:   false,
		MediaUnread: false,
		Silent:      false,
		Post:        false,
		Legacy:      false,
		Id:          0,
		FromId:      mtproto.MakePeerUser(fromId),
		PeerId:      mtproto.MakePeerChat(m.Chat.Id),
		ReplyTo:     nil,
		Date:        int32(time.Now().Unix()),
		Action:      action,
		TtlPeriod:   nil,
	})
	return message.To_Message()
}

func (m *MutableChat) ToUnsafeChat(id int64) *mtproto.Chat {
	var (
		ok bool
		me *ImmutableChatParticipant
	)

	chat := &mtproto.Chat{
		Creator:                 false,
		Kicked:                  false,
		Left:                    false,
		Deactivated:             false,
		CallActive:              m.CallActive(),
		CallNotEmpty:            m.CallNotEmpty(),
		Noforwards:              m.Noforwards(),
		Id:                      m.Id(),
		Title:                   m.Title(),
		Photo:                   mtproto.MakeChatPhotoByPhoto(m.Chat.Photo),
		ParticipantsCount_INT32: m.ParticipantsCount(),
		Date:                    m.Date(),
		Version:                 m.Version(),
		MigratedTo:              nil,
		AdminRights:             nil,
		DefaultBannedRights:     nil,
	}

	// chat deactivated
	if m.Deactivated() {
		chat.Deactivated = true
		chat.MigratedTo = m.MigratedTo()
		return mtproto.MakeTLChat(chat).To_Chat()
	}

	if me, ok = m.GetImmutableChatParticipant(id); !ok {
		// TODO(@benqi): ???
		return mtproto.MakeTLChatEmpty(chat).To_Chat()
	}

	if me.IsChatMemberStateKicked() {
		return mtproto.MakeTLChatForbidden(chat).To_Chat()
	}

	chat.Creator = me.IsChatMemberCreator()
	chat.Left = me.IsChatMemberStateLeft()
	chat.AdminRights = me.GetAdminRights()
	chat.DefaultBannedRights = m.DefaultBannedRights()

	return mtproto.MakeTLChat(chat).To_Chat()
}

func (m *MutableChat) GetImmutableChatParticipant(id int64) (u *ImmutableChatParticipant, ok bool) {
	for _, v := range m.ChatParticipants {
		if v.UserId == id {
			return v, true
		}
	}

	return nil, false
}

func (m *MutableChat) ToChatParticipants(selfId int64) (participants *mtproto.ChatParticipants) {
	if selfId != 0 {
		if me, ok := m.GetImmutableChatParticipant(selfId); ok && !me.IsChatMemberStateNormal() {
			participants = mtproto.MakeTLChatParticipantsForbidden(&mtproto.ChatParticipants{
				ChatId:          m.Chat.Id,
				SelfParticipant: me.ToChatParticipant(),
			}).To_ChatParticipants()
			return
		}
	}

	participants = mtproto.MakeTLChatParticipants(&mtproto.ChatParticipants{
		ChatId:       m.Chat.Id,
		Participants: make([]*mtproto.ChatParticipant, 0, len(m.ChatParticipants)),
		Version:      m.Chat.Version,
	}).To_ChatParticipants()

	for _, cp := range m.ChatParticipants {
		if cp.IsChatMemberStateNormal() {
			participants.Participants = append(participants.Participants, cp.ToChatParticipant())
		}
	}

	return
}

func (m *MutableChat) Walk(visit func(userId int64, participant *ImmutableChatParticipant) error) {
	if visit == nil {
		return
	}
	for _, v := range m.ChatParticipants {
		visit(v.UserId, v)
	}
}

func (m *MutableChat) ToChatForbidden() (chat *mtproto.Chat) {
	chat = mtproto.MakeTLChatForbidden(&mtproto.Chat{
		Id:    m.Chat.Id,
		Title: m.Chat.Title,
	}).To_Chat()
	return
}

func (m *Vector_MutableChat) GetChatListByIdList(selfId int64, id ...int64) []*mtproto.Chat {
	chatList := make([]*mtproto.Chat, 0, len(m.Datas))
	for _, chat2 := range m.Datas {
		chatList = append(chatList, chat2.ToUnsafeChat(selfId))
	}

	return chatList
}
