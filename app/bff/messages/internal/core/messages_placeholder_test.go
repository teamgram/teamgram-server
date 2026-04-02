package core

import (
	"context"
	"testing"

	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

func TestMessagesReadAndDeletePlaceholders(t *testing.T) {
	c := New(context.Background(), nil)

	readHistory, err := c.MessagesReadHistory(&tg.TLMessagesReadHistory{
		Peer:  tg.MakeTLInputPeerUser(&tg.TLInputPeerUser{UserId: 2}),
		MaxId: 9,
	})
	if err != nil || readHistory == nil || readHistory.Pts != 9 || readHistory.PtsCount != 1 {
		t.Fatalf("expected readHistory affectedMessages placeholder, got result=%#v err=%v", readHistory, err)
	}

	readContents, err := c.MessagesReadMessageContents(&tg.TLMessagesReadMessageContents{Id: []int32{4, 12}})
	if err != nil || readContents == nil || readContents.Pts != 12 || readContents.PtsCount != 2 {
		t.Fatalf("expected readMessageContents placeholder, got result=%#v err=%v", readContents, err)
	}

	deleteMessages, err := c.MessagesDeleteMessages(&tg.TLMessagesDeleteMessages{Id: []int32{3, 8}})
	if err != nil || deleteMessages == nil || deleteMessages.Pts != 8 || deleteMessages.PtsCount != 2 {
		t.Fatalf("expected deleteMessages placeholder, got result=%#v err=%v", deleteMessages, err)
	}

	deleteHistory, err := c.MessagesDeleteHistory(&tg.TLMessagesDeleteHistory{
		Peer:  tg.MakeTLInputPeerUser(&tg.TLInputPeerUser{UserId: 2}),
		MaxId: 7,
	})
	if err != nil || deleteHistory == nil || deleteHistory.Pts != 7 || deleteHistory.PtsCount != 1 {
		t.Fatalf("expected deleteHistory placeholder, got result=%#v err=%v", deleteHistory, err)
	}

	unpinAll, err := c.MessagesUnpinAllMessages(&tg.TLMessagesUnpinAllMessages{
		Peer: tg.MakeTLInputPeerUser(&tg.TLInputPeerUser{UserId: 2}),
	})
	if err != nil || unpinAll == nil || unpinAll.Pts != 1 || unpinAll.Offset != 0 {
		t.Fatalf("expected unpinAllMessages affectedHistory placeholder, got result=%#v err=%v", unpinAll, err)
	}
}

func TestMessagesUpdatePinnedMessagePlaceholder(t *testing.T) {
	c := New(context.Background(), nil)

	result, err := c.MessagesUpdatePinnedMessage(&tg.TLMessagesUpdatePinnedMessage{
		Peer: tg.MakeTLInputPeerUser(&tg.TLInputPeerUser{UserId: 2}),
		Id:   11,
	})
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}

	short, ok := result.ToUpdateShort()
	if !ok {
		t.Fatalf("expected updateShort, got %T", result.Clazz)
	}
	update, ok := short.Update.(*tg.TLUpdatePinnedMessages)
	if !ok {
		t.Fatalf("expected updatePinnedMessages, got %T", short.Update)
	}
	if !update.Pinned || len(update.Messages) != 1 || update.Messages[0] != 11 {
		t.Fatalf("unexpected pinned placeholder: %#v", update)
	}
}

func TestMessagesEditMessagePlaceholder(t *testing.T) {
	c := New(context.Background(), nil)

	text := "edited"
	result, err := c.MessagesEditMessage(&tg.TLMessagesEditMessage{
		Peer:    tg.MakeTLInputPeerUser(&tg.TLInputPeerUser{UserId: 2}),
		Id:      15,
		Message: &text,
	})
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}

	short, ok := result.ToUpdateShortSentMessage()
	if !ok {
		t.Fatalf("expected updateShortSentMessage, got %T", result.Clazz)
	}
	if short.Id != 15 || !short.Out {
		t.Fatalf("unexpected edit placeholder: %#v", short)
	}
}

func TestMessagesPlaceholderRejectsChannelPeer(t *testing.T) {
	c := New(context.Background(), nil)

	_, err := c.MessagesUpdatePinnedMessage(&tg.TLMessagesUpdatePinnedMessage{
		Peer: tg.MakeTLInputPeerChannel(&tg.TLInputPeerChannel{ChannelId: 3}),
		Id:   11,
	})
	if err != tg.ErrEnterpriseIsBlocked {
		t.Fatalf("expected ErrEnterpriseIsBlocked, got %v", err)
	}
}

func TestMessagesQueryPlaceholders(t *testing.T) {
	c := New(context.Background(), nil)

	history, err := c.MessagesGetHistory(&tg.TLMessagesGetHistory{
		Peer:     tg.MakeTLInputPeerUser(&tg.TLInputPeerUser{UserId: 2}),
		OffsetId: 20,
		Limit:    2,
	})
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
	historyMsgs, ok := history.ToMessagesMessages()
	if !ok || len(historyMsgs.Messages) != 2 {
		t.Fatalf("expected 2 history placeholders, got %#v", history)
	}

	first, ok := historyMsgs.Messages[0].(*tg.TLMessage)
	if !ok || first.Id != 20 {
		t.Fatalf("expected first placeholder id=20, got %#v", historyMsgs.Messages[0])
	}

	getMessages, err := c.MessagesGetMessages(&tg.TLMessagesGetMessages{
		Id_VECTORINPUTMESSAGE: []tg.InputMessageClazz{
			tg.MakeTLInputMessageID(&tg.TLInputMessageID{Id: 7}),
			tg.MakeTLInputMessageReplyTo(&tg.TLInputMessageReplyTo{Id: 9}),
		},
	})
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
	getMsgs, ok := getMessages.ToMessagesMessages()
	if !ok || len(getMsgs.Messages) != 2 {
		t.Fatalf("expected 2 getMessages placeholders, got %#v", getMessages)
	}

	second, ok := getMsgs.Messages[1].(*tg.TLMessage)
	if !ok || second.Id != 9 {
		t.Fatalf("expected second placeholder id=9, got %#v", getMsgs.Messages[1])
	}

	unread, err := c.MessagesGetUnreadMentions(&tg.TLMessagesGetUnreadMentions{
		Peer:     tg.MakeTLInputPeerUser(&tg.TLInputPeerUser{UserId: 2}),
		OffsetId: 30,
		Limit:    1,
	})
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
	unreadMsgs, ok := unread.ToMessagesMessages()
	if !ok || len(unreadMsgs.Messages) != 1 {
		t.Fatalf("expected 1 unread mention placeholder, got %#v", unread)
	}

	unreadMsg, ok := unreadMsgs.Messages[0].(*tg.TLMessage)
	if !ok || !unreadMsg.Mentioned {
		t.Fatalf("expected mentioned placeholder message, got %#v", unreadMsgs.Messages[0])
	}

	search, err := c.MessagesSearch(&tg.TLMessagesSearch{
		Peer:     tg.MakeTLInputPeerUser(&tg.TLInputPeerUser{UserId: 2}),
		Q:        "hi",
		OffsetId: 40,
		Limit:    1,
	})
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
	searchMsgs, ok := search.ToMessagesMessages()
	if !ok || len(searchMsgs.Messages) != 1 {
		t.Fatalf("expected 1 search placeholder, got %#v", search)
	}
}
