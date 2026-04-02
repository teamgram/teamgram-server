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

func TestMessagesSendMessageReturnsShortSentMessagePlaceholderForUserPeer(t *testing.T) {
	c := New(context.Background(), nil)

	result, err := c.MessagesSendMessage(&tg.TLMessagesSendMessage{
		Peer:     tg.MakeTLInputPeerUser(&tg.TLInputPeerUser{UserId: 2, AccessHash: 0}),
		Message:  "hello",
		RandomId: 22,
	})
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
	if result == nil {
		t.Fatal("expected updates result, got nil")
	}

	updates, ok := result.ToUpdateShortSentMessage()
	if !ok {
		t.Fatalf("expected updateShortSentMessage placeholder, got %T", result.Clazz)
	}
	if !updates.Out {
		t.Fatal("expected out=true")
	}
	if updates.Id != 22 {
		t.Fatalf("expected placeholder id=22, got %d", updates.Id)
	}
	if updates.Pts != 1 || updates.PtsCount != 1 {
		t.Fatalf("expected pts/pts_count to be 1/1, got %d/%d", updates.Pts, updates.PtsCount)
	}
	if updates.Date == 0 {
		t.Fatal("expected non-zero date")
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

func TestMessagesSendMessageReusesPlaceholderIDForSameRandomID(t *testing.T) {
	c := New(context.Background(), nil)

	req := &tg.TLMessagesSendMessage{
		Peer:     tg.MakeTLInputPeerUser(&tg.TLInputPeerUser{UserId: 2, AccessHash: 0}),
		Message:  "hello",
		RandomId: 77,
	}

	first, err := c.MessagesSendMessage(req)
	if err != nil {
		t.Fatalf("first send: expected nil error, got %v", err)
	}
	second, err := c.MessagesSendMessage(req)
	if err != nil {
		t.Fatalf("second send: expected nil error, got %v", err)
	}

	firstShort, ok := first.ToUpdateShortSentMessage()
	if !ok {
		t.Fatalf("expected first result to be updateShortSentMessage, got %T", first.Clazz)
	}
	secondShort, ok := second.ToUpdateShortSentMessage()
	if !ok {
		t.Fatalf("expected second result to be updateShortSentMessage, got %T", second.Clazz)
	}
	if firstShort.Id != secondShort.Id {
		t.Fatalf("expected same placeholder id for repeated random_id, got %d vs %d", firstShort.Id, secondShort.Id)
	}
}
