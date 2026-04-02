package core

import (
	"context"
	"testing"

	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

func TestMessagesGetDialogsReturnsSinglePlaceholderForUserPeer(t *testing.T) {
	c := New(context.Background(), nil)

	result, err := c.MessagesGetDialogs(&tg.TLMessagesGetDialogs{
		OffsetPeer: tg.MakeTLInputPeerUser(&tg.TLInputPeerUser{UserId: 42, AccessHash: 0}),
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
	if dialogsSlice.Count != 1 {
		t.Fatalf("expected count=1, got %d", dialogsSlice.Count)
	}
	if len(dialogsSlice.Dialogs) != 1 || len(dialogsSlice.Messages) != 1 || len(dialogsSlice.Chats) != 0 || len(dialogsSlice.Users) != 1 {
		t.Fatalf("expected single placeholder dialog/message/user, got dialogs=%d messages=%d chats=%d users=%d",
			len(dialogsSlice.Dialogs), len(dialogsSlice.Messages), len(dialogsSlice.Chats), len(dialogsSlice.Users))
	}
	dialog, ok := dialogsSlice.Dialogs[0].(*tg.TLDialog)
	if !ok {
		t.Fatalf("expected dialog placeholder, got %T", dialogsSlice.Dialogs[0])
	}
	if dialog.TopMessage != 10 {
		t.Fatalf("expected top_message=10, got %d", dialog.TopMessage)
	}
	user, ok := dialogsSlice.Users[0].(*tg.TLUserEmpty)
	if !ok {
		t.Fatalf("expected userEmpty placeholder, got %T", dialogsSlice.Users[0])
	}
	if user.Id != 42 {
		t.Fatalf("expected placeholder user id=42, got %d", user.Id)
	}
}

func TestMessagesGetDialogsReturnsEmptyDialogsSlicePlaceholderForEmptyPeer(t *testing.T) {
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

func TestMessagesGetPeerDialogsReturnsSinglePlaceholderForUserPeer(t *testing.T) {
	c := New(context.Background(), nil)

	result, err := c.MessagesGetPeerDialogs(&tg.TLMessagesGetPeerDialogs{
		Peers: []tg.InputDialogPeerClazz{
			tg.MakeTLInputDialogPeer(&tg.TLInputDialogPeer{
				Peer: tg.MakeTLInputPeerUser(&tg.TLInputPeerUser{UserId: 42, AccessHash: 0}),
			}),
		},
	})
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
	if result == nil {
		t.Fatal("expected peer dialogs result, got nil")
	}
	if len(result.Dialogs) != 1 {
		t.Fatalf("expected 1 dialog placeholder, got %d", len(result.Dialogs))
	}
	if len(result.Messages) != 1 {
		t.Fatalf("expected 1 message placeholder, got %d", len(result.Messages))
	}
	if len(result.Users) != 1 {
		t.Fatalf("expected 1 user placeholder, got %d", len(result.Users))
	}
	dialog, ok := result.Dialogs[0].(*tg.TLDialog)
	if !ok {
		t.Fatalf("expected dialog placeholder, got %T", result.Dialogs[0])
	}
	if dialog.TopMessage != 10 {
		t.Fatalf("expected top_message=10, got %d", dialog.TopMessage)
	}
	user, ok := result.Users[0].(*tg.TLUserEmpty)
	if !ok {
		t.Fatalf("expected userEmpty placeholder, got %T", result.Users[0])
	}
	if user.Id != 42 {
		t.Fatalf("expected placeholder user id=42, got %d", user.Id)
	}
}
