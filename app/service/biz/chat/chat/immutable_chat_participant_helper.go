// Copyright (c) 2021-present,  Teamgram Studio (https://teamgram.io).
//  All rights reserved.
//
// Author: teamgramio (teamgram.io@gmail.com)
//

package chat

import (
	"github.com/teamgram/proto/mtproto"
)

func (m *ImmutableChatParticipant) IsChatMemberNormal() bool {
	return m.ParticipantType == mtproto.ChatMemberNormal
}

func (m *ImmutableChatParticipant) IsChatMemberAdmin() bool {
	return m.ParticipantType == mtproto.ChatMemberAdmin
}

func (m *ImmutableChatParticipant) IsChatMemberCreator() bool {
	return m.ParticipantType == mtproto.ChatMemberCreator
}

func (m *ImmutableChatParticipant) IsChatMemberStateNormal() bool {
	return m.State == mtproto.ChatMemberStateNormal
}

func (m *ImmutableChatParticipant) IsChatMemberStateLeft() bool {
	return m.State == mtproto.ChatMemberStateLeft
}

func (m *ImmutableChatParticipant) IsChatMemberStateKicked() bool {
	return m.State == mtproto.ChatMemberStateKicked
}

func (m *ImmutableChatParticipant) IsChatMemberStateMigrated() bool {
	return m.State == mtproto.ChatMemberStateMigrated
}

func (m *ImmutableChatParticipant) CanViewMessages() bool {
	if m.IsChatMemberCreator() || m.IsChatMemberAdmin() {
		return true
	}
	return false // m.Chat.CanViewMessages()
}

func (m *ImmutableChatParticipant) CanSendMessages() bool {
	if m.IsChatMemberCreator() || m.IsChatMemberAdmin() {
		return true
	}
	return false // m.Chat.CanSendMessages()
}

func (m *ImmutableChatParticipant) CanSendMedia() bool {
	if m.IsChatMemberCreator() || m.IsChatMemberAdmin() {
		return true
	}
	return false // m.Chat.CanSendMedia()
}

func (m *ImmutableChatParticipant) CanSendStickers() bool {
	if m.IsChatMemberCreator() || m.IsChatMemberAdmin() {
		return true
	}
	return false // m.Chat.CanSendStickers()
}

func (m *ImmutableChatParticipant) CanSendGifs() bool {
	if m.IsChatMemberCreator() || m.IsChatMemberAdmin() {
		return true
	}
	return false // m.Chat.CanSendGifs()
}

func (m *ImmutableChatParticipant) CanSendGames() bool {
	if m.IsChatMemberCreator() || m.IsChatMemberAdmin() {
		return true
	}
	return false // m.Chat.CanSendGames()
}

func (m *ImmutableChatParticipant) CanSendInline() bool {
	if m.IsChatMemberCreator() || m.IsChatMemberAdmin() {
		return true
	}
	return false // m.Chat.CanSendInline()
}

func (m *ImmutableChatParticipant) CanEmbedLinks() bool {
	if m.IsChatMemberCreator() || m.IsChatMemberAdmin() {
		return true
	}
	return false // m.Chat.CanEmbedLinks()
}

func (m *ImmutableChatParticipant) CanSendPolls() bool {
	if m.IsChatMemberCreator() || m.IsChatMemberAdmin() {
		return true
	}
	return false // m.Chat.CanSendPolls()
}

// CanChangeInfo - merge ChatAdminRights
func (m *ImmutableChatParticipant) CanChangeInfo() bool {
	if m.IsChatMemberCreator() || m.IsChatMemberAdmin() {
		return true
	}
	return false // m.Chat.CanChangeInfo()
}

func (m *ImmutableChatParticipant) CanInviteUsers() bool {
	if m.IsChatMemberCreator() || m.IsChatMemberAdmin() {
		return true
	}
	return false // m.Chat.CanInviteUsers()
}

func (m *ImmutableChatParticipant) CanPinMessages() bool {
	if m.IsChatMemberCreator() || m.IsChatMemberAdmin() {
		return true
	}
	return false // m.Chat.CanPinMessages()
}

func (m *ImmutableChatParticipant) CanAdminChangeInfo() bool {
	if m.IsChatMemberCreator() {
		return true
	}
	return m.IsChatMemberAdmin() && m.AdminRights.GetChangeInfo()
}

func (m *ImmutableChatParticipant) CanAdminPostMessages() bool {
	if m.IsChatMemberCreator() {
		return true
	}

	return m.IsChatMemberAdmin() && m.AdminRights.GetPinMessages()
}

func (m *ImmutableChatParticipant) CanAdminEditMessages() bool {
	if m.IsChatMemberCreator() {
		return true
	}

	return m.IsChatMemberAdmin() && m.AdminRights.GetEditMessages()
}

func (m *ImmutableChatParticipant) CanAdminDeleteMessages() bool {
	if m.IsChatMemberCreator() {
		return true
	}

	return m.IsChatMemberAdmin() && m.AdminRights.GetDeleteMessages()
}

func (m *ImmutableChatParticipant) CanAdminBanUsers() bool {
	if m.IsChatMemberCreator() {
		return true
	}

	return m.IsChatMemberAdmin() && m.AdminRights.GetBanUsers()
}

func (m *ImmutableChatParticipant) CanAdminInviteUsers() bool {
	if m.IsChatMemberCreator() {
		return true
	}

	return m.IsChatMemberAdmin() && m.AdminRights.GetInviteUsers()
}

func (m *ImmutableChatParticipant) CanAdminPinMessages() bool {
	if m.IsChatMemberCreator() {
		return true
	}

	return m.IsChatMemberAdmin() && m.AdminRights.GetPinMessages()
}

func (m *ImmutableChatParticipant) CanAdminAddAdmins() bool {
	if m.IsChatMemberCreator() {
		return true
	}

	return m.IsChatMemberAdmin() && m.AdminRights.GetAddAdmins()
}

func (m *ImmutableChatParticipant) ToChatParticipant() *mtproto.ChatParticipant {
	switch m.ParticipantType {
	case mtproto.ChatMemberCreator:
		return mtproto.MakeTLChatParticipantCreator(&mtproto.ChatParticipant{
			UserId: m.UserId,
		}).To_ChatParticipant()
	case mtproto.ChatMemberAdmin:
		return mtproto.MakeTLChatParticipantAdmin(&mtproto.ChatParticipant{
			UserId:    m.UserId,
			InviterId: m.InviterUserId,
			Date:      int32(m.InvitedAt),
		}).To_ChatParticipant()
	default:
		return mtproto.MakeTLChatParticipant(&mtproto.ChatParticipant{
			UserId:    m.UserId,
			InviterId: m.InviterUserId,
			Date:      int32(m.InvitedAt),
		}).To_ChatParticipant()
	}
}
