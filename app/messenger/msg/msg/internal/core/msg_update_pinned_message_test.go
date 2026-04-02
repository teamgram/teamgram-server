package core

import (
	"context"
	"testing"

	"github.com/teamgram/teamgram-server/v2/app/messenger/msg/msg/msg"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

func TestMsgUpdatePinnedMessageReturnsPinnedUpdateShort(t *testing.T) {
	c := New(context.Background(), nil)

	result, err := c.MsgUpdatePinnedMessage(&msg.TLMsgUpdatePinnedMessage{
		UserId:    1,
		AuthKeyId: 2,
		PeerType:  tg.PEER_USER,
		PeerId:    3,
		Id:        11,
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
	if !update.Pinned {
		t.Fatal("expected pinned=true")
	}
	if len(update.Messages) != 1 || update.Messages[0] != 11 {
		t.Fatalf("expected pinned message id 11, got %#v", update.Messages)
	}
}

func TestMsgUpdatePinnedMessageReturnsUnpinnedUpdateShort(t *testing.T) {
	c := New(context.Background(), nil)

	result, err := c.MsgUpdatePinnedMessage(&msg.TLMsgUpdatePinnedMessage{
		UserId:    1,
		AuthKeyId: 2,
		PeerType:  tg.PEER_CHAT,
		PeerId:    9,
		Id:        22,
		Unpin:     true,
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
	if update.Pinned {
		t.Fatal("expected pinned=false for unpin flow")
	}
	peer, ok := update.Peer.(*tg.TLPeerChat)
	if !ok {
		t.Fatalf("expected peerChat, got %T", update.Peer)
	}
	if peer.ChatId != 9 {
		t.Fatalf("expected chat id 9, got %d", peer.ChatId)
	}
}

func TestMsgUpdatePinnedMessageRejectsInvalidPeerType(t *testing.T) {
	c := New(context.Background(), nil)

	result, err := c.MsgUpdatePinnedMessage(&msg.TLMsgUpdatePinnedMessage{
		UserId:    1,
		AuthKeyId: 2,
		PeerType:  99,
		PeerId:    3,
		Id:        11,
	})
	if err != tg.ErrPeerIdInvalid {
		t.Fatalf("expected ErrPeerIdInvalid, got %v", err)
	}
	if result != nil {
		t.Fatalf("expected nil result, got %v", result)
	}
}
