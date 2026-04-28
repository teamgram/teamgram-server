package repository

import (
	"reflect"
	"testing"

	"github.com/teamgram/teamgram-server/v2/app/service/biz/chat/chat"
	"github.com/teamgram/teamgram-server/v2/app/service/biz/chat/internal/repository/model"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

func TestChatBannedRightsStorageRoundTrip(t *testing.T) {
	rights := tg.MakeTLChatBannedRights(&tg.TLChatBannedRights{
		ViewMessages:    true,
		SendMessages:    true,
		SendMedia:       true,
		SendStickers:    true,
		SendGifs:        true,
		SendGames:       true,
		SendInline:      true,
		EmbedLinks:      true,
		SendPolls:       true,
		ChangeInfo:      true,
		InviteUsers:     true,
		PinMessages:     true,
		ManageTopics:    true,
		SendPhotos:      true,
		SendVideos:      true,
		SendRoundvideos: true,
		SendAudios:      true,
		SendVoices:      true,
		SendDocs:        true,
		SendPlain:       true,
		EditRank:        true,
		UntilDate:       12345,
	}).ToChatBannedRights()

	mask := chatBannedRightsToStorage(rights)
	got := chatBannedRightsFromStorage(mask)

	if !got.ViewMessages || !got.SendMessages || !got.SendMedia || !got.SendStickers ||
		!got.SendGifs || !got.SendGames || !got.SendInline || !got.EmbedLinks ||
		!got.SendPolls || !got.ChangeInfo || !got.InviteUsers || !got.PinMessages ||
		!got.ManageTopics || !got.SendPhotos || !got.SendVideos || !got.SendRoundvideos ||
		!got.SendAudios || !got.SendVoices || !got.SendDocs || !got.SendPlain || !got.EditRank {
		t.Fatalf("round trip rights lost: %#v", got)
	}
	if got.UntilDate != 0 {
		t.Fatalf("until_date must not be stored, got %d", got.UntilDate)
	}
}

func TestChatBannedRightsGoldenBits(t *testing.T) {
	tests := []struct {
		name  string
		mask  int64
		make  func() tg.ChatBannedRightsClazz
		check func(*tg.TLChatBannedRights) bool
	}{
		{"ViewMessages", 1 << 0, func() tg.ChatBannedRightsClazz {
			return tg.MakeTLChatBannedRights(&tg.TLChatBannedRights{ViewMessages: true}).ToChatBannedRights()
		}, func(r *tg.TLChatBannedRights) bool { return r.ViewMessages }},
		{"SendMessages", 1 << 1, func() tg.ChatBannedRightsClazz {
			return tg.MakeTLChatBannedRights(&tg.TLChatBannedRights{SendMessages: true}).ToChatBannedRights()
		}, func(r *tg.TLChatBannedRights) bool { return r.SendMessages }},
		{"SendMedia", 1 << 2, func() tg.ChatBannedRightsClazz {
			return tg.MakeTLChatBannedRights(&tg.TLChatBannedRights{SendMedia: true}).ToChatBannedRights()
		}, func(r *tg.TLChatBannedRights) bool { return r.SendMedia }},
		{"SendStickers", 1 << 3, func() tg.ChatBannedRightsClazz {
			return tg.MakeTLChatBannedRights(&tg.TLChatBannedRights{SendStickers: true}).ToChatBannedRights()
		}, func(r *tg.TLChatBannedRights) bool { return r.SendStickers }},
		{"SendGifs", 1 << 4, func() tg.ChatBannedRightsClazz {
			return tg.MakeTLChatBannedRights(&tg.TLChatBannedRights{SendGifs: true}).ToChatBannedRights()
		}, func(r *tg.TLChatBannedRights) bool { return r.SendGifs }},
		{"SendGames", 1 << 5, func() tg.ChatBannedRightsClazz {
			return tg.MakeTLChatBannedRights(&tg.TLChatBannedRights{SendGames: true}).ToChatBannedRights()
		}, func(r *tg.TLChatBannedRights) bool { return r.SendGames }},
		{"SendInline", 1 << 6, func() tg.ChatBannedRightsClazz {
			return tg.MakeTLChatBannedRights(&tg.TLChatBannedRights{SendInline: true}).ToChatBannedRights()
		}, func(r *tg.TLChatBannedRights) bool { return r.SendInline }},
		{"EmbedLinks", 1 << 7, func() tg.ChatBannedRightsClazz {
			return tg.MakeTLChatBannedRights(&tg.TLChatBannedRights{EmbedLinks: true}).ToChatBannedRights()
		}, func(r *tg.TLChatBannedRights) bool { return r.EmbedLinks }},
		{"SendPolls", 1 << 8, func() tg.ChatBannedRightsClazz {
			return tg.MakeTLChatBannedRights(&tg.TLChatBannedRights{SendPolls: true}).ToChatBannedRights()
		}, func(r *tg.TLChatBannedRights) bool { return r.SendPolls }},
		{"ChangeInfo", 1 << 10, func() tg.ChatBannedRightsClazz {
			return tg.MakeTLChatBannedRights(&tg.TLChatBannedRights{ChangeInfo: true}).ToChatBannedRights()
		}, func(r *tg.TLChatBannedRights) bool { return r.ChangeInfo }},
		{"InviteUsers", 1 << 15, func() tg.ChatBannedRightsClazz {
			return tg.MakeTLChatBannedRights(&tg.TLChatBannedRights{InviteUsers: true}).ToChatBannedRights()
		}, func(r *tg.TLChatBannedRights) bool { return r.InviteUsers }},
		{"PinMessages", 1 << 17, func() tg.ChatBannedRightsClazz {
			return tg.MakeTLChatBannedRights(&tg.TLChatBannedRights{PinMessages: true}).ToChatBannedRights()
		}, func(r *tg.TLChatBannedRights) bool { return r.PinMessages }},
		{"ManageTopics", 1 << 18, func() tg.ChatBannedRightsClazz {
			return tg.MakeTLChatBannedRights(&tg.TLChatBannedRights{ManageTopics: true}).ToChatBannedRights()
		}, func(r *tg.TLChatBannedRights) bool { return r.ManageTopics }},
		{"SendPhotos", 1 << 19, func() tg.ChatBannedRightsClazz {
			return tg.MakeTLChatBannedRights(&tg.TLChatBannedRights{SendPhotos: true}).ToChatBannedRights()
		}, func(r *tg.TLChatBannedRights) bool { return r.SendPhotos }},
		{"SendVideos", 1 << 20, func() tg.ChatBannedRightsClazz {
			return tg.MakeTLChatBannedRights(&tg.TLChatBannedRights{SendVideos: true}).ToChatBannedRights()
		}, func(r *tg.TLChatBannedRights) bool { return r.SendVideos }},
		{"SendRoundvideos", 1 << 21, func() tg.ChatBannedRightsClazz {
			return tg.MakeTLChatBannedRights(&tg.TLChatBannedRights{SendRoundvideos: true}).ToChatBannedRights()
		}, func(r *tg.TLChatBannedRights) bool { return r.SendRoundvideos }},
		{"SendAudios", 1 << 22, func() tg.ChatBannedRightsClazz {
			return tg.MakeTLChatBannedRights(&tg.TLChatBannedRights{SendAudios: true}).ToChatBannedRights()
		}, func(r *tg.TLChatBannedRights) bool { return r.SendAudios }},
		{"SendVoices", 1 << 23, func() tg.ChatBannedRightsClazz {
			return tg.MakeTLChatBannedRights(&tg.TLChatBannedRights{SendVoices: true}).ToChatBannedRights()
		}, func(r *tg.TLChatBannedRights) bool { return r.SendVoices }},
		{"SendDocs", 1 << 24, func() tg.ChatBannedRightsClazz {
			return tg.MakeTLChatBannedRights(&tg.TLChatBannedRights{SendDocs: true}).ToChatBannedRights()
		}, func(r *tg.TLChatBannedRights) bool { return r.SendDocs }},
		{"SendPlain", 1 << 25, func() tg.ChatBannedRightsClazz {
			return tg.MakeTLChatBannedRights(&tg.TLChatBannedRights{SendPlain: true}).ToChatBannedRights()
		}, func(r *tg.TLChatBannedRights) bool { return r.SendPlain }},
		{"EditRank", 1 << 26, func() tg.ChatBannedRightsClazz {
			return tg.MakeTLChatBannedRights(&tg.TLChatBannedRights{EditRank: true}).ToChatBannedRights()
		}, func(r *tg.TLChatBannedRights) bool { return r.EditRank }},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := chatBannedRightsToStorage(tt.make()); got != tt.mask {
				t.Fatalf("chatBannedRightsToStorage = %d, want %d", got, tt.mask)
			}
			got := chatBannedRightsFromStorage(tt.mask)
			if !tt.check(got) {
				t.Fatalf("chatBannedRightsFromStorage(%d) did not set %s: %#v", tt.mask, tt.name, got)
			}
			if roundTrip := chatBannedRightsToStorage(got); roundTrip != tt.mask {
				t.Fatalf("single-bit round trip = %d, want %d", roundTrip, tt.mask)
			}
		})
	}
}

func TestChatAdminRightsStorageRoundTrip(t *testing.T) {
	rights := tg.MakeTLChatAdminRights(&tg.TLChatAdminRights{
		ChangeInfo:           true,
		BanUsers:             true,
		InviteUsers:          true,
		PinMessages:          true,
		AddAdmins:            true,
		ManageCall:           true,
		ManageTopics:         true,
		PostMessages:         true,
		EditMessages:         true,
		DeleteMessages:       true,
		Anonymous:            true,
		Other:                true,
		PostStories:          true,
		EditStories:          true,
		DeleteStories:        true,
		ManageDirectMessages: true,
		ManageRanks:          true,
	}).ToChatAdminRights()

	mask := chatAdminRightsToStorage(rights)
	got := chatAdminRightsFromStorage(mask)

	if !reflect.DeepEqual(got, rights) {
		t.Fatalf("round trip admin rights mismatch: got %#v want %#v", got, rights)
	}
}

func TestChatAdminRightsGoldenBits(t *testing.T) {
	tests := []struct {
		name  string
		mask  int32
		make  func() tg.ChatAdminRightsClazz
		check func(*tg.TLChatAdminRights) bool
	}{
		{"ChangeInfo", 1 << 0, func() tg.ChatAdminRightsClazz {
			return tg.MakeTLChatAdminRights(&tg.TLChatAdminRights{ChangeInfo: true}).ToChatAdminRights()
		}, func(r *tg.TLChatAdminRights) bool { return r.ChangeInfo }},
		{"PostMessages", 1 << 1, func() tg.ChatAdminRightsClazz {
			return tg.MakeTLChatAdminRights(&tg.TLChatAdminRights{PostMessages: true}).ToChatAdminRights()
		}, func(r *tg.TLChatAdminRights) bool { return r.PostMessages }},
		{"EditMessages", 1 << 2, func() tg.ChatAdminRightsClazz {
			return tg.MakeTLChatAdminRights(&tg.TLChatAdminRights{EditMessages: true}).ToChatAdminRights()
		}, func(r *tg.TLChatAdminRights) bool { return r.EditMessages }},
		{"DeleteMessages", 1 << 3, func() tg.ChatAdminRightsClazz {
			return tg.MakeTLChatAdminRights(&tg.TLChatAdminRights{DeleteMessages: true}).ToChatAdminRights()
		}, func(r *tg.TLChatAdminRights) bool { return r.DeleteMessages }},
		{"BanUsers", 1 << 4, func() tg.ChatAdminRightsClazz {
			return tg.MakeTLChatAdminRights(&tg.TLChatAdminRights{BanUsers: true}).ToChatAdminRights()
		}, func(r *tg.TLChatAdminRights) bool { return r.BanUsers }},
		{"InviteUsers", 1 << 5, func() tg.ChatAdminRightsClazz {
			return tg.MakeTLChatAdminRights(&tg.TLChatAdminRights{InviteUsers: true}).ToChatAdminRights()
		}, func(r *tg.TLChatAdminRights) bool { return r.InviteUsers }},
		{"PinMessages", 1 << 7, func() tg.ChatAdminRightsClazz {
			return tg.MakeTLChatAdminRights(&tg.TLChatAdminRights{PinMessages: true}).ToChatAdminRights()
		}, func(r *tg.TLChatAdminRights) bool { return r.PinMessages }},
		{"AddAdmins", 1 << 9, func() tg.ChatAdminRightsClazz {
			return tg.MakeTLChatAdminRights(&tg.TLChatAdminRights{AddAdmins: true}).ToChatAdminRights()
		}, func(r *tg.TLChatAdminRights) bool { return r.AddAdmins }},
		{"Anonymous", 1 << 10, func() tg.ChatAdminRightsClazz {
			return tg.MakeTLChatAdminRights(&tg.TLChatAdminRights{Anonymous: true}).ToChatAdminRights()
		}, func(r *tg.TLChatAdminRights) bool { return r.Anonymous }},
		{"ManageCall", 1 << 11, func() tg.ChatAdminRightsClazz {
			return tg.MakeTLChatAdminRights(&tg.TLChatAdminRights{ManageCall: true}).ToChatAdminRights()
		}, func(r *tg.TLChatAdminRights) bool { return r.ManageCall }},
		{"Other", 1 << 12, func() tg.ChatAdminRightsClazz {
			return tg.MakeTLChatAdminRights(&tg.TLChatAdminRights{Other: true}).ToChatAdminRights()
		}, func(r *tg.TLChatAdminRights) bool { return r.Other }},
		{"ManageTopics", 1 << 13, func() tg.ChatAdminRightsClazz {
			return tg.MakeTLChatAdminRights(&tg.TLChatAdminRights{ManageTopics: true}).ToChatAdminRights()
		}, func(r *tg.TLChatAdminRights) bool { return r.ManageTopics }},
		{"PostStories", 1 << 14, func() tg.ChatAdminRightsClazz {
			return tg.MakeTLChatAdminRights(&tg.TLChatAdminRights{PostStories: true}).ToChatAdminRights()
		}, func(r *tg.TLChatAdminRights) bool { return r.PostStories }},
		{"EditStories", 1 << 15, func() tg.ChatAdminRightsClazz {
			return tg.MakeTLChatAdminRights(&tg.TLChatAdminRights{EditStories: true}).ToChatAdminRights()
		}, func(r *tg.TLChatAdminRights) bool { return r.EditStories }},
		{"DeleteStories", 1 << 16, func() tg.ChatAdminRightsClazz {
			return tg.MakeTLChatAdminRights(&tg.TLChatAdminRights{DeleteStories: true}).ToChatAdminRights()
		}, func(r *tg.TLChatAdminRights) bool { return r.DeleteStories }},
		{"ManageDirectMessages", 1 << 17, func() tg.ChatAdminRightsClazz {
			return tg.MakeTLChatAdminRights(&tg.TLChatAdminRights{ManageDirectMessages: true}).ToChatAdminRights()
		}, func(r *tg.TLChatAdminRights) bool { return r.ManageDirectMessages }},
		{"ManageRanks", 1 << 18, func() tg.ChatAdminRightsClazz {
			return tg.MakeTLChatAdminRights(&tg.TLChatAdminRights{ManageRanks: true}).ToChatAdminRights()
		}, func(r *tg.TLChatAdminRights) bool { return r.ManageRanks }},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := chatAdminRightsToStorage(tt.make()); got != tt.mask {
				t.Fatalf("chatAdminRightsToStorage = %d, want %d", got, tt.mask)
			}
			got := chatAdminRightsFromStorage(tt.mask)
			if !tt.check(got) {
				t.Fatalf("chatAdminRightsFromStorage(%d) did not set %s: %#v", tt.mask, tt.name, got)
			}
			if roundTrip := chatAdminRightsToStorage(got); roundTrip != tt.mask {
				t.Fatalf("single-bit round trip = %d, want %d", roundTrip, tt.mask)
			}
		})
	}
}

func TestAvailableReactionsFromStorage(t *testing.T) {
	got := availableReactionsFromStorage(4, `["👍","🔥"]`)
	want := []string{"👍", "🔥"}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("availableReactionsFromStorage valid = %#v, want %#v", got, want)
	}

	if got := availableReactionsFromStorage(4, `{bad json`); len(got) != 0 {
		t.Fatalf("malformed available reactions = %#v, want empty", got)
	}

	if got := availableReactionsFromStorage(1, `["👍"]`); len(got) != 0 {
		t.Fatalf("non-some available reactions = %#v, want empty", got)
	}
}

func TestAvailableReactionsToStorage(t *testing.T) {
	if got := availableReactionsToStorage(nil); got != "[]" {
		t.Fatalf("nil available reactions storage = %q, want []", got)
	}
	if got := availableReactionsToStorage([]string{}); got != "[]" {
		t.Fatalf("empty available reactions storage = %q, want []", got)
	}
	if got := availableReactionsToStorage([]string{"👍", "🔥"}); got != `["👍","🔥"]` {
		t.Fatalf("available reactions storage = %q", got)
	}
}

func TestImmutableChatMapsStorageFields(t *testing.T) {
	photo := tg.MakeTLPhotoEmpty(&tg.TLPhotoEmpty{Id: 10})
	row := &model.Chats{
		Id:                     100,
		CreatorUserId:          200,
		Title:                  "chat title",
		About:                  "about",
		ParticipantCount:       3,
		DefaultBannedRights:    chatBannedRightsToStorage(tg.MakeTLChatBannedRights(&tg.TLChatBannedRights{SendMessages: true}).ToChatBannedRights()),
		AvailableReactionsType: 4,
		AvailableReactions:     `["👍","🔥"]`,
		Deactivated:            true,
		Noforwards:             true,
		TtlPeriod:              3600,
		Version:                7,
		Date:                   123456,
		MigratedToId:           300,
		MigratedToAccessHash:   400,
	}

	got := makeImmutableChat(row, photo)
	if got == nil {
		t.Fatal("makeImmutableChat returned nil")
	}
	if got.Id != row.Id || got.Creator != row.CreatorUserId || got.Title != row.Title ||
		got.About != row.About || got.ParticipantsCount != row.ParticipantCount ||
		got.Deactivated != row.Deactivated || got.Noforwards != row.Noforwards ||
		got.TtlPeriod != row.TtlPeriod || got.Version != row.Version || got.Date != row.Date {
		t.Fatalf("immutable chat fields mismatch: %#v", got)
	}
	if got.Photo != photo {
		t.Fatalf("immutable chat photo = %p, want %p", got.Photo, photo)
	}
	if got.DefaultBannedRights == nil || !got.DefaultBannedRights.SendMessages || got.DefaultBannedRights.UntilDate != 0 {
		t.Fatalf("default banned rights mismatch: %#v", got.DefaultBannedRights)
	}
	if !reflect.DeepEqual(got.AvailableReactions, []string{"👍", "🔥"}) {
		t.Fatalf("available reactions = %#v", got.AvailableReactions)
	}
	migrated, ok := got.MigratedTo.(*tg.TLInputChannel)
	if !ok || migrated.ChannelId != row.MigratedToId || migrated.AccessHash != row.MigratedToAccessHash {
		t.Fatalf("migrated_to mismatch: %#v", got.MigratedTo)
	}
}

func TestImmutableChatParticipantMapsStorageFields(t *testing.T) {
	row := &model.ChatParticipants{
		Id:              1,
		ChatId:          2,
		UserId:          3,
		ParticipantType: chat.ChatMemberAdmin,
		Link:            "invite-link",
		Usage2:          4,
		AdminRights:     chatAdminRightsToStorage(tg.MakeTLChatAdminRights(&tg.TLChatAdminRights{BanUsers: true, InviteUsers: true}).ToChatAdminRights()),
		InviterUserId:   5,
		InvitedAt:       6,
		KickedAt:        7,
		LeftAt:          8,
		State:           chat.ChatMemberStateNormal,
		Date2:           9,
		IsBot:           true,
	}

	got := makeImmutableChatParticipant(row)
	if got == nil {
		t.Fatal("makeImmutableChatParticipant returned nil")
	}
	if got.Id != row.Id || got.ChatId != row.ChatId || got.UserId != row.UserId ||
		got.ParticipantType != row.ParticipantType || got.Link != row.Link || got.Useage != row.Usage2 ||
		got.InviterUserId != row.InviterUserId || got.InvitedAt != row.InvitedAt ||
		got.KickedAt != row.KickedAt || got.LeftAt != row.LeftAt || got.State != row.State ||
		got.Date != row.Date2 || got.IsBot != row.IsBot {
		t.Fatalf("immutable participant fields mismatch: %#v", got)
	}
	if got.AdminRights == nil || !got.AdminRights.BanUsers || !got.AdminRights.InviteUsers {
		t.Fatalf("admin rights mismatch: %#v", got.AdminRights)
	}
}

func TestImmutableChatParticipantAdminZeroMaskDoesNotOvergrant(t *testing.T) {
	got := makeImmutableChatParticipant(&model.ChatParticipants{
		ParticipantType: chat.ChatMemberAdmin,
		State:           chat.ChatMemberStateNormal,
	})
	if got == nil || got.AdminRights == nil {
		t.Fatalf("admin participant should have non-nil rights: %#v", got)
	}
	if got.AdminRights.ChangeInfo || got.AdminRights.BanUsers || got.AdminRights.InviteUsers ||
		got.AdminRights.PinMessages || got.AdminRights.AddAdmins || got.AdminRights.ManageCall ||
		got.AdminRights.ManageTopics {
		t.Fatalf("zero admin rights mask overgranted rights: %#v", got.AdminRights)
	}
}
