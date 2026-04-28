package chat

import (
	"reflect"
	"testing"

	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

func TestMutableChatHelpersAreNilSafe(t *testing.T) {
	if MutableChatData(nil) != nil {
		t.Fatal("MutableChatData(nil) != nil")
	}
	if p, ok := GetImmutableChatParticipant(nil, 1); p != nil || ok {
		t.Fatalf("GetImmutableChatParticipant(nil) = %v, %v; want nil, false", p, ok)
	}
	WalkChatParticipants(nil, func(*tg.ImmutableChatParticipant) bool {
		t.Fatal("WalkChatParticipants called callback for nil chat")
		return true
	})
	if got := ChatParticipantIDList(nil); got != nil {
		t.Fatalf("ChatParticipantIDList(nil) = %#v, want nil", got)
	}
	if ChatCreator(nil) != 0 || ChatParticipantsCount(nil) != 0 || ChatDeactivated(nil) {
		t.Fatal("nil aggregate scalar helpers returned non-zero values")
	}
	if ChatTitle(nil) != "" || ChatAbout(nil) != "" || ChatPhoto(nil) != nil || ChatDefaultBannedRights(nil) != nil {
		t.Fatal("nil aggregate object/string helpers returned non-zero values")
	}
	if IsChatMemberStateNormal(nil) || IsChatMemberCreator(nil) || IsChatMemberAdmin(nil) {
		t.Fatal("nil participant role helpers returned true")
	}
	if CanInviteUsers(nil) || CanChangeInfo(nil) || CanAdminBanUsers(nil) || CanAdminAddAdmins(nil) {
		t.Fatal("nil participant permission helpers returned true")
	}
}

func TestMutableChatAggregateAccessors(t *testing.T) {
	photo := tg.MakeTLPhotoEmpty(&tg.TLPhotoEmpty{})
	bannedRights := tg.MakeTLChatBannedRights(&tg.TLChatBannedRights{SendMessages: true})
	data := tg.MakeTLImmutableChat(&tg.TLImmutableChat{
		Id:                  10,
		Creator:             100,
		Title:               "chat title",
		About:               "about",
		Photo:               photo,
		Deactivated:         true,
		ParticipantsCount:   3,
		DefaultBannedRights: bannedRights,
	})
	participant := participantWithRights(200, ChatMemberAdmin, ChatMemberStateNormal, nil)
	mutable := tg.MakeTLMutableChat(&tg.TLMutableChat{
		Chat:             data,
		ChatParticipants: []tg.ImmutableChatParticipantClazz{participant, nil},
	})

	if got := MutableChatData(mutable); got != data {
		t.Fatalf("MutableChatData = %p, want %p", got, data)
	}
	if got := ChatCreator(mutable); got != 100 {
		t.Fatalf("ChatCreator = %d, want 100", got)
	}
	if got := ChatParticipantsCount(mutable); got != 3 {
		t.Fatalf("ChatParticipantsCount = %d, want 3", got)
	}
	if !ChatDeactivated(mutable) {
		t.Fatal("ChatDeactivated = false, want true")
	}
	if got := ChatTitle(mutable); got != "chat title" {
		t.Fatalf("ChatTitle = %q, want chat title", got)
	}
	if got := ChatAbout(mutable); got != "about" {
		t.Fatalf("ChatAbout = %q, want about", got)
	}
	if got := ChatPhoto(mutable); got != photo {
		t.Fatalf("ChatPhoto = %p, want %p", got, photo)
	}
	if got := ChatDefaultBannedRights(mutable); got != bannedRights {
		t.Fatalf("ChatDefaultBannedRights = %p, want %p", got, bannedRights)
	}
	if got, ok := GetImmutableChatParticipant(mutable, 200); got != participant || !ok {
		t.Fatalf("GetImmutableChatParticipant = %v, %v; want participant, true", got, ok)
	}
	if got := ChatParticipantIDList(mutable); !reflect.DeepEqual(got, []int64{200}) {
		t.Fatalf("ChatParticipantIDList = %#v, want []int64{200}", got)
	}
}

func TestWalkChatParticipantsStopsWhenCallbackReturnsFalse(t *testing.T) {
	mutable := tg.MakeTLMutableChat(&tg.TLMutableChat{
		ChatParticipants: []tg.ImmutableChatParticipantClazz{
			participantWithRights(1, ChatMemberNormal, ChatMemberStateNormal, nil),
			participantWithRights(2, ChatMemberNormal, ChatMemberStateNormal, nil),
		},
	})

	var seen []int64
	WalkChatParticipants(mutable, func(p *tg.ImmutableChatParticipant) bool {
		seen = append(seen, p.UserId)
		return false
	})

	if !reflect.DeepEqual(seen, []int64{1}) {
		t.Fatalf("WalkChatParticipants seen = %#v, want []int64{1}", seen)
	}
}

func TestParticipantStateRoleAndPermissions(t *testing.T) {
	creator := participantWithRights(1, ChatMemberCreator, ChatMemberStateNormal, nil)
	admin := participantWithRights(2, ChatMemberAdmin, ChatMemberStateNormal, &tg.TLChatAdminRights{
		InviteUsers: true,
		ChangeInfo:  true,
		BanUsers:    true,
		AddAdmins:   true,
	})
	normal := participantWithRights(3, ChatMemberNormal, ChatMemberStateLeft, &tg.TLChatAdminRights{
		InviteUsers: true,
		ChangeInfo:  true,
		BanUsers:    true,
		AddAdmins:   true,
	})
	adminWithoutRights := participantWithRights(4, ChatMemberAdmin, ChatMemberStateNormal, nil)

	if !IsChatMemberStateNormal(creator) || IsChatMemberStateNormal(normal) {
		t.Fatal("IsChatMemberStateNormal returned unexpected result")
	}
	if !IsChatMemberCreator(creator) || IsChatMemberAdmin(creator) {
		t.Fatal("creator role helpers returned unexpected result")
	}
	if !IsChatMemberAdmin(admin) || IsChatMemberCreator(admin) {
		t.Fatal("admin role helpers returned unexpected result")
	}
	for name, fn := range map[string]func(*tg.ImmutableChatParticipant) bool{
		"CanInviteUsers":    CanInviteUsers,
		"CanChangeInfo":     CanChangeInfo,
		"CanAdminBanUsers":  CanAdminBanUsers,
		"CanAdminAddAdmins": CanAdminAddAdmins,
	} {
		if !fn(creator) {
			t.Fatalf("%s(creator) = false, want true", name)
		}
		if !fn(admin) {
			t.Fatalf("%s(admin) = false, want true", name)
		}
		if !fn(normal) {
			t.Fatalf("%s(normal with rights) = false, want true", name)
		}
		if fn(adminWithoutRights) {
			t.Fatalf("%s(admin without rights) = true, want false", name)
		}
	}
}

func participantWithRights(userID int64, participantType int32, state int32, rights *tg.TLChatAdminRights) *tg.ImmutableChatParticipant {
	return tg.MakeTLImmutableChatParticipant(&tg.TLImmutableChatParticipant{
		UserId:          userID,
		ParticipantType: participantType,
		State:           state,
		AdminRights:     tg.MakeTLChatAdminRights(rights),
	})
}
