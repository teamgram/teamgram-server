package repository

import (
	"errors"
	"testing"

	chatpb "github.com/teamgram/teamgram-server/v2/app/service/biz/chat/chat"
	"github.com/teamgram/teamgram-server/v2/app/service/biz/chat/internal/repository/model"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

func TestInviteExpiredBranches(t *testing.T) {
	if !IsInviteExpired(&model.ChatInvites{Revoked: true}, 0, 10) {
		t.Fatal("revoked invite should be expired")
	}
	if !IsInviteExpired(&model.ChatInvites{ExpireDate: 9}, 0, 10) {
		t.Fatal("past expire_date should be expired")
	}
	if !IsInviteExpired(&model.ChatInvites{UsageLimit: 2}, 2, 10) {
		t.Fatal("usage limit should be expired")
	}
	if IsInviteExpired(&model.ChatInvites{UsageLimit: 3}, 2, 10) {
		t.Fatal("invite below usage limit should be usable")
	}
}

func TestExportedInviteProjectionBuildsFullLinkAndCounts(t *testing.T) {
	title := "primary"
	got := makeChatInviteExported(&model.ChatInvites{
		ChatId:        10,
		AdminId:       1,
		Link:          "hash",
		Permanent:     true,
		RequestNeeded: true,
		ExpireDate:    99,
		UsageLimit:    3,
		Title:         title,
		Date2:         7,
	}, 2, 1)
	if got == nil {
		t.Fatal("projection returned nil")
	}
	invite, ok := got.Clazz.(*tg.TLChatInviteExported)
	if !ok {
		t.Fatalf("projection clazz = %T, want TLChatInviteExported", got.Clazz)
	}
	if invite.Link != "https://t.me/+hash" {
		t.Fatalf("link = %q, want full t.me invite link", invite.Link)
	}
	if invite.Usage == nil || *invite.Usage != 2 || invite.Requested == nil || *invite.Requested != 1 {
		t.Fatalf("counts usage=%v requested=%v, want 2/1", invite.Usage, invite.Requested)
	}
	if invite.UsageLimit == nil || *invite.UsageLimit != 3 || invite.ExpireDate == nil || *invite.ExpireDate != 99 {
		t.Fatalf("limits expire=%v usage_limit=%v, want 99/3", invite.ExpireDate, invite.UsageLimit)
	}
	if invite.Title == nil || *invite.Title != title {
		t.Fatalf("title = %v, want %q", invite.Title, title)
	}
}

func TestInviteRowForWrongChatReturnsLinkExists(t *testing.T) {
	err := requireInviteRowForChat(&model.ChatInvites{ChatId: 20, Link: "hash"}, 10)
	if !errors.Is(err, chatpb.ErrChatLinkExists) {
		t.Fatalf("requireInviteRowForChat error = %v, want ErrChatLinkExists", err)
	}
}

func TestInviteImporterRequestedLinkUsesLinkSpecificQuery(t *testing.T) {
	requested, recent := inviteImporterLinkQuery(ChatInviteImporterQuery{
		ChatID:    10,
		Requested: true,
		Link:      "hash",
	})
	if recent {
		t.Fatal("requested query with link should not use chat-wide recent requested list")
	}
	if requested != 1 {
		t.Fatalf("requested flag = %d, want 1", requested)
	}
}

func TestInviteImporterRequestedWithoutLinkUsesRecentQuery(t *testing.T) {
	requested, recent := inviteImporterLinkQuery(ChatInviteImporterQuery{ChatID: 10, Requested: true})
	if !recent {
		t.Fatal("requested query without link should use chat-wide recent requested list")
	}
	if requested != 1 {
		t.Fatalf("requested flag = %d, want 1", requested)
	}
}

func TestRequireRowsAffectedMapsMissingJoinRequest(t *testing.T) {
	if err := requireRowsAffected(1); err != nil {
		t.Fatalf("requireRowsAffected(1) error = %v", err)
	}
	err := requireRowsAffected(0)
	if !errors.Is(err, chatpb.ErrUserNotParticipant) {
		t.Fatalf("requireRowsAffected(0) error = %v, want ErrUserNotParticipant", err)
	}
}

func TestExportedInviteOffsetSkipsCursorAndNormalizesLink(t *testing.T) {
	invites := []tg.ExportedChatInviteClazz{
		tg.MakeTLChatInviteExported(&tg.TLChatInviteExported{Link: "https://t.me/+first", Date: 10}).ToExportedChatInvite().Clazz,
		tg.MakeTLChatInviteExported(&tg.TLChatInviteExported{Link: "https://t.me/+second", Date: 20}).ToExportedChatInvite().Clazz,
		tg.MakeTLChatInviteExported(&tg.TLChatInviteExported{Link: "https://t.me/+third", Date: 30}).ToExportedChatInvite().Clazz,
	}
	offset := exportedInviteOffset(invites, 20, "second")
	if offset != 2 {
		t.Fatalf("offset = %d, want 2 to skip cursor", offset)
	}
	if got := exportedInviteOffset(invites, 20, "https://t.me/+second"); got != 2 {
		t.Fatalf("full-link offset = %d, want 2", got)
	}
	if got := exportedInviteOffset(invites, 99, "missing"); got != -1 {
		t.Fatalf("missing offset = %d, want -1", got)
	}
}
