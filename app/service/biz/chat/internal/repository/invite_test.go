package repository

import (
	"testing"

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
