package core

import (
	"errors"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/teamgram/teamgram-server/v2/app/service/biz/chat/chat"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

func TestChatEditChatTitleRejectsEmptyTitle(t *testing.T) {
	write := &fakeWriteRepo{}
	core := newWriteTestCore(&fakeReadRepo{}, write)

	_, err := core.ChatEditChatTitle(&chat.TLChatEditChatTitle{ChatId: 10, EditUserId: 1})
	if !errors.Is(err, chat.ErrChatTitleEmpty) {
		t.Fatalf("ChatEditChatTitle error = %v, want ErrChatTitleEmpty", err)
	}
	if write.titleCalls != 0 {
		t.Fatalf("UpdateChatTitle calls = %d, want 0", write.titleCalls)
	}
}

func TestChatEditChatTitleRejectsSameTitle(t *testing.T) {
	m := mutableChatForMemberTests(10, 1,
		participantForMemberTests(10, 1, chat.ChatMemberCreator, chat.ChatMemberStateNormal, nil))
	m.Chat.Title = "team"
	write := &fakeWriteRepo{}
	core := newWriteTestCore(&fakeReadRepo{mutableChat: m}, write)

	_, err := core.ChatEditChatTitle(&chat.TLChatEditChatTitle{ChatId: 10, EditUserId: 1, Title: "team"})
	if !errors.Is(err, chat.ErrChatNotModified) {
		t.Fatalf("ChatEditChatTitle error = %v, want ErrChatNotModified", err)
	}
	if write.titleCalls != 0 {
		t.Fatalf("UpdateChatTitle calls = %d, want 0", write.titleCalls)
	}
}

func TestChatEditChatAboutRejectsSameAbout(t *testing.T) {
	m := mutableChatForMemberTests(10, 1,
		participantForMemberTests(10, 1, chat.ChatMemberCreator, chat.ChatMemberStateNormal, nil))
	m.Chat.About = "about"
	write := &fakeWriteRepo{}
	core := newWriteTestCore(&fakeReadRepo{mutableChat: m}, write)

	_, err := core.ChatEditChatAbout(&chat.TLChatEditChatAbout{ChatId: 10, EditUserId: 1, About: "about"})
	if !errors.Is(err, chat.ErrChatNotModified) {
		t.Fatalf("ChatEditChatAbout error = %v, want ErrChatNotModified", err)
	}
	if write.aboutCalls != 0 {
		t.Fatalf("UpdateChatAbout calls = %d, want 0", write.aboutCalls)
	}
}

func TestChatEditAttributeMutationRejectsNonAdmin(t *testing.T) {
	m := mutableChatForMemberTests(10, 1,
		participantForMemberTests(10, 2, chat.ChatMemberNormal, chat.ChatMemberStateNormal, nil))
	write := &fakeWriteRepo{}
	core := newWriteTestCore(&fakeReadRepo{mutableChat: m}, write)

	_, err := core.ChatEditChatTitle(&chat.TLChatEditChatTitle{ChatId: 10, EditUserId: 2, Title: "new"})
	if !errors.Is(err, chat.ErrChatAdminRequired) {
		t.Fatalf("ChatEditChatTitle error = %v, want ErrChatAdminRequired", err)
	}
	if write.titleCalls != 0 {
		t.Fatalf("UpdateChatTitle calls = %d, want 0", write.titleCalls)
	}
}

func TestChatEditChatAdminRejectsCreatorDemotion(t *testing.T) {
	m := mutableChatForMemberTests(10, 1,
		participantForMemberTests(10, 1, chat.ChatMemberCreator, chat.ChatMemberStateNormal, nil))
	write := &fakeWriteRepo{}
	core := newWriteTestCore(&fakeReadRepo{mutableChat: m}, write)

	_, err := core.ChatEditChatAdmin(&chat.TLChatEditChatAdmin{
		ChatId:          10,
		OperatorId:      1,
		EditChatAdminId: 1,
		IsAdmin:         tg.BoolFalseClazz,
	})
	if !errors.Is(err, chat.ErrParticipantInvalid) {
		t.Fatalf("ChatEditChatAdmin error = %v, want ErrParticipantInvalid", err)
	}
	if write.adminCalls != 0 {
		t.Fatalf("UpdateChatAdmin calls = %d, want 0", write.adminCalls)
	}
}

func TestChatEditChatAdminRejectsMissingTargetAsUserNotParticipant(t *testing.T) {
	m := mutableChatForMemberTests(10, 1,
		participantForMemberTests(10, 1, chat.ChatMemberCreator, chat.ChatMemberStateNormal, nil))
	write := &fakeWriteRepo{}
	core := newWriteTestCore(&fakeReadRepo{mutableChat: m}, write)

	_, err := core.ChatEditChatAdmin(&chat.TLChatEditChatAdmin{
		ChatId:          10,
		OperatorId:      1,
		EditChatAdminId: 2,
		IsAdmin:         tg.BoolTrueClazz,
	})
	if !errors.Is(err, chat.ErrUserNotParticipant) {
		t.Fatalf("ChatEditChatAdmin error = %v, want ErrUserNotParticipant", err)
	}
	if write.adminCalls != 0 {
		t.Fatalf("UpdateChatAdmin calls = %d, want 0", write.adminCalls)
	}
}

func TestChatEditChatDefaultBannedRightsRejectsAdminWithoutBanUsers(t *testing.T) {
	adminRights := tg.MakeTLChatAdminRights(&tg.TLChatAdminRights{AddAdmins: true}).ToChatAdminRights()
	m := mutableChatForMemberTests(10, 1,
		participantForMemberTests(10, 2, chat.ChatMemberAdmin, chat.ChatMemberStateNormal, adminRights))
	write := &fakeWriteRepo{}
	core := newWriteTestCore(&fakeReadRepo{mutableChat: m}, write)

	_, err := core.ChatEditChatDefaultBannedRights(&chat.TLChatEditChatDefaultBannedRights{
		ChatId:     10,
		OperatorId: 2,
		BannedRights: tg.MakeTLChatBannedRights(&tg.TLChatBannedRights{
			SendMessages: true,
		}).ToChatBannedRights(),
	})
	if !errors.Is(err, chat.ErrChatAdminRequired) {
		t.Fatalf("ChatEditChatDefaultBannedRights error = %v, want ErrChatAdminRequired", err)
	}
	if write.bannedCalls != 0 {
		t.Fatalf("UpdateChatDefaultBannedRights calls = %d, want 0", write.bannedCalls)
	}
}

func TestChatSetChatAvailableReactionsPersistsThroughRepository(t *testing.T) {
	adminRights := tg.MakeTLChatAdminRights(&tg.TLChatAdminRights{AddAdmins: true}).ToChatAdminRights()
	m := mutableChatForMemberTests(10, 1,
		participantForMemberTests(10, 2, chat.ChatMemberAdmin, chat.ChatMemberStateNormal, adminRights))
	m.Chat.Date = 77
	write := &fakeWriteRepo{updateDate: 123}
	core := newWriteTestCore(&fakeReadRepo{mutableChat: m}, write)

	got, err := core.ChatSetChatAvailableReactions(&chat.TLChatSetChatAvailableReactions{
		SelfId:                 2,
		ChatId:                 10,
		AvailableReactionsType: 4,
		AvailableReactions:     []string{"like", "fire"},
	})
	if err != nil {
		t.Fatalf("ChatSetChatAvailableReactions error: %v", err)
	}
	if write.reactionsCalls != 1 || write.reactionsType != 4 || len(write.reactions) != 2 {
		t.Fatalf("repository reactions call = calls:%d type:%d reactions:%#v", write.reactionsCalls, write.reactionsType, write.reactions)
	}
	if got.Chat.Version != 1 || got.Chat.Date != 77 || strings.Join(got.Chat.AvailableReactions, ",") != "like,fire" {
		t.Fatalf("chat reaction update = version:%d date:%d reactions:%#v", got.Chat.Version, got.Chat.Date, got.Chat.AvailableReactions)
	}
}

func TestChatToggleNoForwardsRejectsNonCreator(t *testing.T) {
	adminRights := tg.MakeTLChatAdminRights(&tg.TLChatAdminRights{AddAdmins: true}).ToChatAdminRights()
	m := mutableChatForMemberTests(10, 1,
		participantForMemberTests(10, 2, chat.ChatMemberAdmin, chat.ChatMemberStateNormal, adminRights))
	write := &fakeWriteRepo{}
	core := newWriteTestCore(&fakeReadRepo{mutableChat: m}, write)

	_, err := core.ChatToggleNoForwards(&chat.TLChatToggleNoForwards{
		ChatId:     10,
		OperatorId: 2,
		Enabled:    tg.BoolTrueClazz,
	})
	if !errors.Is(err, chat.ErrChatAdminRequired) {
		t.Fatalf("ChatToggleNoForwards error = %v, want ErrChatAdminRequired", err)
	}
	if write.noForwardsCalls != 0 {
		t.Fatalf("UpdateChatNoForwards calls = %d, want 0", write.noForwardsCalls)
	}
}

func TestChatSetHistoryTTLRejectsInactiveCreatorParticipant(t *testing.T) {
	for _, state := range []int32{chat.ChatMemberStateLeft, chat.ChatMemberStateKicked} {
		t.Run("state", func(t *testing.T) {
			m := mutableChatForMemberTests(10, 1,
				participantForMemberTests(10, 1, chat.ChatMemberCreator, state, nil))
			write := &fakeWriteRepo{}
			core := newWriteTestCore(&fakeReadRepo{mutableChat: m}, write)

			_, err := core.ChatSetHistoryTTL(&chat.TLChatSetHistoryTTL{SelfId: 1, ChatId: 10, TtlPeriod: 86400})
			if !errors.Is(err, chat.ErrInputUserDeactivated) {
				t.Fatalf("ChatSetHistoryTTL error = %v, want ErrInputUserDeactivated", err)
			}
			if write.ttlCalls != 0 {
				t.Fatalf("UpdateChatTTLPeriod calls = %d, want 0", write.ttlCalls)
			}
		})
	}
}

func TestChatSetHistoryTTLRejectsMissingCreatorParticipant(t *testing.T) {
	m := mutableChatForMemberTests(10, 1)
	write := &fakeWriteRepo{}
	core := newWriteTestCore(&fakeReadRepo{mutableChat: m}, write)

	_, err := core.ChatSetHistoryTTL(&chat.TLChatSetHistoryTTL{SelfId: 1, ChatId: 10, TtlPeriod: 86400})
	if !errors.Is(err, chat.ErrInputUserDeactivated) {
		t.Fatalf("ChatSetHistoryTTL error = %v, want ErrInputUserDeactivated", err)
	}
	if write.ttlCalls != 0 {
		t.Fatalf("UpdateChatTTLPeriod calls = %d, want 0", write.ttlCalls)
	}
}

func TestChatSetHistoryTTLPersistsCreatorMutation(t *testing.T) {
	m := mutableChatForMemberTests(10, 1,
		participantForMemberTests(10, 1, chat.ChatMemberCreator, chat.ChatMemberStateNormal, nil))
	m.Chat.Date = 88
	write := &fakeWriteRepo{updateDate: 321}
	core := newWriteTestCore(&fakeReadRepo{mutableChat: m}, write)

	got, err := core.ChatSetHistoryTTL(&chat.TLChatSetHistoryTTL{SelfId: 1, ChatId: 10, TtlPeriod: 86400})
	if err != nil {
		t.Fatalf("ChatSetHistoryTTL error: %v", err)
	}
	if write.ttlCalls != 1 || write.ttlPeriod != 86400 {
		t.Fatalf("repository ttl call = calls:%d ttl:%d", write.ttlCalls, write.ttlPeriod)
	}
	if got.Chat.TtlPeriod != 86400 || got.Chat.Version != 1 || got.Chat.Date != 88 {
		t.Fatalf("chat ttl update = ttl:%d version:%d date:%d", got.Chat.TtlPeriod, got.Chat.Version, got.Chat.Date)
	}
}

func TestTask6HandlersDoNotReturnMethodNotImplOrTGErrors(t *testing.T) {
	files := []string{
		"chat.editChatTitle_handler.go",
		"chat.editChatAbout_handler.go",
		"chat.editChatPhoto_handler.go",
		"chat.editChatAdmin_handler.go",
		"chat.editChatDefaultBannedRights_handler.go",
		"chat.toggleNoForwards_handler.go",
		"chat.setHistoryTTL_handler.go",
		"chat.setChatAvailableReactions_handler.go",
	}

	for _, name := range files {
		data, err := os.ReadFile(filepath.Join(".", name))
		if err != nil {
			t.Fatalf("read %s: %v", name, err)
		}
		s := string(data)
		for _, bad := range []string{"ErrMethodNotImpl", "method Chat", "not impl", "tg.Err"} {
			if strings.Contains(s, bad) {
				t.Fatalf("%s still contains %q", name, bad)
			}
		}
	}
}
