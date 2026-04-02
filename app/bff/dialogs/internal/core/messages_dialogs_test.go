package core

import (
	"context"
	"testing"

	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

func TestMessagesGetDialogsReturnsEmptyDialogsSlicePlaceholder(t *testing.T) {
	c := New(context.Background(), nil)

	result, err := c.MessagesGetDialogs(&tg.TLMessagesGetDialogs{
		OffsetPeer: tg.MakeTLInputPeerEmpty(&tg.TLInputPeerEmpty{}),
		Limit:      20,
	})
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
	if result == nil {
		t.Fatal("expected dialogs result, got nil")
	}

	dialogsSlice, ok := result.ToMessagesDialogsSlice()
	if !ok {
		t.Fatalf("expected messages.dialogsSlice placeholder, got %T", result.Clazz)
	}
	if dialogsSlice.Count != 0 {
		t.Fatalf("expected count=0, got %d", dialogsSlice.Count)
	}
	if len(dialogsSlice.Dialogs) != 0 || len(dialogsSlice.Messages) != 0 || len(dialogsSlice.Chats) != 0 || len(dialogsSlice.Users) != 0 {
		t.Fatalf("expected empty placeholder lists, got dialogs=%d messages=%d chats=%d users=%d",
			len(dialogsSlice.Dialogs), len(dialogsSlice.Messages), len(dialogsSlice.Chats), len(dialogsSlice.Users))
	}
}

func TestMessagesGetPeerDialogsReturnsEmptyPeerDialogsPlaceholder(t *testing.T) {
	c := New(context.Background(), nil)

	result, err := c.MessagesGetPeerDialogs(&tg.TLMessagesGetPeerDialogs{})
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
	if result == nil {
		t.Fatal("expected peer dialogs result, got nil")
	}
	if len(result.Dialogs) != 0 || len(result.Messages) != 0 || len(result.Chats) != 0 || len(result.Users) != 0 {
		t.Fatalf("expected empty placeholder lists, got dialogs=%d messages=%d chats=%d users=%d",
			len(result.Dialogs), len(result.Messages), len(result.Chats), len(result.Users))
	}
	if result.State == nil {
		t.Fatal("expected placeholder updates state, got nil")
	}
}
