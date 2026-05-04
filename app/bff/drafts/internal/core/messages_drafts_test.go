package core

import (
	"context"
	"testing"

	"github.com/teamgram/teamgram-server/v2/app/bff/drafts/internal/repository"
	"github.com/teamgram/teamgram-server/v2/app/bff/drafts/internal/svc"
	"github.com/teamgram/teamgram-server/v2/pkg/net/kitex/metadata"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

func newDraftsCoreForTest(repo *repository.Repository, selfID int64) *DraftsCore {
	c := New(context.Background(), &svc.ServiceContext{Repo: repo})
	c.MD = &metadata.RpcMetadata{UserId: selfID}
	return c
}

func TestMessagesSaveDraftNoopsWhenDialogClientIsNotConfigured(t *testing.T) {
	c := newDraftsCoreForTest(&repository.Repository{}, 100)

	r, err := c.MessagesSaveDraft(&tg.TLMessagesSaveDraft{
		Peer:    tg.MakeTLInputPeerUser(&tg.TLInputPeerUser{UserId: 200}),
		Message: "draft",
	})
	if err != nil {
		t.Fatalf("MessagesSaveDraft error = %v", err)
	}
	if r != tg.BoolTrue {
		t.Fatalf("MessagesSaveDraft = %v, want boolTrue", r)
	}
}

func TestMessagesGetAllDraftsNoopsWhenDialogClientIsNotConfigured(t *testing.T) {
	c := newDraftsCoreForTest(&repository.Repository{}, 100)

	r, err := c.MessagesGetAllDrafts(&tg.TLMessagesGetAllDrafts{})
	if err != nil {
		t.Fatalf("MessagesGetAllDrafts error = %v", err)
	}
	if r == nil {
		t.Fatal("MessagesGetAllDrafts returned nil")
	}
	updates, ok := r.ToUpdates()
	if !ok {
		t.Fatalf("MessagesGetAllDrafts returned %s, want updates", r.ClazzName())
	}
	if len(updates.Updates) != 0 || len(updates.Users) != 0 || len(updates.Chats) != 0 {
		t.Fatalf("MessagesGetAllDrafts reply = %+v, want empty updates", updates)
	}
}

func TestMessagesClearAllDraftsNoopsWhenDialogClientIsNotConfigured(t *testing.T) {
	c := newDraftsCoreForTest(&repository.Repository{}, 100)

	r, err := c.MessagesClearAllDrafts(&tg.TLMessagesClearAllDrafts{})
	if err != nil {
		t.Fatalf("MessagesClearAllDrafts error = %v", err)
	}
	if r != tg.BoolTrue {
		t.Fatalf("MessagesClearAllDrafts = %v, want boolTrue", r)
	}
}
