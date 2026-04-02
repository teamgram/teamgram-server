package core

import (
	"context"
	"testing"

	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

func TestMessagesSendMessageRejectsEmptyMessage(t *testing.T) {
	c := New(context.Background(), nil)

	result, err := c.MessagesSendMessage(&tg.TLMessagesSendMessage{
		Peer:     tg.MakeTLInputPeerUser(&tg.TLInputPeerUser{UserId: 2, AccessHash: 0}),
		Message:  "",
		RandomId: 1,
	})
	if err != tg.ErrMessageEmpty {
		t.Fatalf("expected ErrMessageEmpty, got %v", err)
	}
	if result != nil {
		t.Fatalf("expected nil result, got %v", result)
	}
}

func TestMessagesSendMessageReturnsEmptyUpdatesPlaceholderForUserPeer(t *testing.T) {
	c := New(context.Background(), nil)

	result, err := c.MessagesSendMessage(&tg.TLMessagesSendMessage{
		Peer:     tg.MakeTLInputPeerUser(&tg.TLInputPeerUser{UserId: 2, AccessHash: 0}),
		Message:  "hello",
		RandomId: 2,
	})
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
	if result == nil {
		t.Fatal("expected updates result, got nil")
	}

	updates, ok := result.Clazz.(*tg.TLUpdates)
	if !ok {
		t.Fatalf("expected updates placeholder, got %T", result.Clazz)
	}
	if len(updates.Updates) != 0 || len(updates.Users) != 0 || len(updates.Chats) != 0 {
		t.Fatalf("expected empty updates payload, got updates=%d users=%d chats=%d",
			len(updates.Updates), len(updates.Users), len(updates.Chats))
	}
}

func TestMessagesSendMessageRejectsChannelPeerPlaceholder(t *testing.T) {
	c := New(context.Background(), nil)

	result, err := c.MessagesSendMessage(&tg.TLMessagesSendMessage{
		Peer:     tg.MakeTLInputPeerChannel(&tg.TLInputPeerChannel{ChannelId: 3, AccessHash: 0}),
		Message:  "hello",
		RandomId: 3,
	})
	if err != tg.ErrEnterpriseIsBlocked {
		t.Fatalf("expected ErrEnterpriseIsBlocked, got %v", err)
	}
	if result != nil {
		t.Fatalf("expected nil result, got %v", result)
	}
}

func TestMessagesSendMessageRejectsEmptyPeerPlaceholder(t *testing.T) {
	c := New(context.Background(), nil)

	result, err := c.MessagesSendMessage(&tg.TLMessagesSendMessage{
		Peer:     tg.MakeTLInputPeerEmpty(&tg.TLInputPeerEmpty{}),
		Message:  "hello",
		RandomId: 4,
	})
	if err != tg.ErrPeerIdInvalid {
		t.Fatalf("expected ErrPeerIdInvalid, got %v", err)
	}
	if result != nil {
		t.Fatalf("expected nil result, got %v", result)
	}
}
