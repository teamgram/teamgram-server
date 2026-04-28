package chat

import "github.com/teamgram/teamgram-server/v2/pkg/proto/tg"

func MutableChatData(m *tg.MutableChat) *tg.ImmutableChat {
	if m == nil {
		return nil
	}
	return m.Chat
}

func GetImmutableChatParticipant(m *tg.MutableChat, userID int64) (*tg.ImmutableChatParticipant, bool) {
	if m == nil {
		return nil, false
	}
	for _, p := range m.ChatParticipants {
		if p != nil && p.UserId == userID {
			return p, true
		}
	}
	return nil, false
}

func WalkChatParticipants(m *tg.MutableChat, cb func(*tg.ImmutableChatParticipant) bool) {
	if m == nil || cb == nil {
		return
	}
	for _, p := range m.ChatParticipants {
		if p == nil {
			continue
		}
		if !cb(p) {
			return
		}
	}
}

func ChatParticipantIDList(m *tg.MutableChat) []int64 {
	if m == nil {
		return nil
	}
	idList := make([]int64, 0, len(m.ChatParticipants))
	WalkChatParticipants(m, func(p *tg.ImmutableChatParticipant) bool {
		idList = append(idList, p.UserId)
		return true
	})
	return idList
}

func ChatCreator(m *tg.MutableChat) int64 {
	if data := MutableChatData(m); data != nil {
		return data.Creator
	}
	return 0
}

func ChatParticipantsCount(m *tg.MutableChat) int32 {
	if data := MutableChatData(m); data != nil {
		return data.ParticipantsCount
	}
	return 0
}

func ChatDeactivated(m *tg.MutableChat) bool {
	if data := MutableChatData(m); data != nil {
		return data.Deactivated
	}
	return false
}

func ChatTitle(m *tg.MutableChat) string {
	if data := MutableChatData(m); data != nil {
		return data.Title
	}
	return ""
}

func ChatAbout(m *tg.MutableChat) string {
	if data := MutableChatData(m); data != nil {
		return data.About
	}
	return ""
}

func ChatPhoto(m *tg.MutableChat) tg.PhotoClazz {
	if data := MutableChatData(m); data != nil {
		return data.Photo
	}
	return nil
}

func ChatDefaultBannedRights(m *tg.MutableChat) tg.ChatBannedRightsClazz {
	if data := MutableChatData(m); data != nil {
		return data.DefaultBannedRights
	}
	return nil
}

func IsChatMemberStateNormal(p *tg.ImmutableChatParticipant) bool {
	return p != nil && p.State == ChatMemberStateNormal
}

func IsChatMemberCreator(p *tg.ImmutableChatParticipant) bool {
	return IsChatMemberStateNormal(p) && p.ParticipantType == ChatMemberCreator
}

func IsChatMemberAdmin(p *tg.ImmutableChatParticipant) bool {
	return IsChatMemberStateNormal(p) && p.ParticipantType == ChatMemberAdmin
}

func CanInviteUsers(p *tg.ImmutableChatParticipant) bool {
	return isCreatorOrHasRight(p, func(rights tg.ChatAdminRightsClazz) bool {
		return rights.InviteUsers
	})
}

func CanChangeInfo(p *tg.ImmutableChatParticipant) bool {
	return isCreatorOrHasRight(p, func(rights tg.ChatAdminRightsClazz) bool {
		return rights.ChangeInfo
	})
}

func CanAdminBanUsers(p *tg.ImmutableChatParticipant) bool {
	return isCreatorOrHasRight(p, func(rights tg.ChatAdminRightsClazz) bool {
		return rights.BanUsers
	})
}

func CanAdminAddAdmins(p *tg.ImmutableChatParticipant) bool {
	return isCreatorOrHasRight(p, func(rights tg.ChatAdminRightsClazz) bool {
		return rights.AddAdmins
	})
}

func isCreatorOrHasRight(p *tg.ImmutableChatParticipant, hasRight func(tg.ChatAdminRightsClazz) bool) bool {
	if IsChatMemberCreator(p) {
		return true
	}
	if !IsChatMemberAdmin(p) || p.AdminRights == nil {
		return false
	}
	return hasRight(p.AdminRights)
}
