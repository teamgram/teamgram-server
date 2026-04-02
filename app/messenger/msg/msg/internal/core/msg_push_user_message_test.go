package core

import (
	"context"
	"testing"

	"github.com/teamgram/teamgram-server/v2/app/messenger/msg/msg/msg"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

func TestMsgPushUserMessageRejectsNonUserPeer(t *testing.T) {
	c := New(context.Background(), nil)

	result, err := c.MsgPushUserMessage(&msg.TLMsgPushUserMessage{
		UserId:    1,
		AuthKeyId: 2,
		PeerType:  tg.PEER_CHAT,
		PeerId:    3,
		PushType:  1,
		Message:   msg.MakeOutboxMessage(&msg.TLOutboxMessage{}),
	})
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
	if result != tg.BoolFalse {
		t.Fatalf("expected BoolFalse, got %v", result)
	}
}

func TestMsgPushUserMessageAcceptsUserPeerAsPlaceholder(t *testing.T) {
	c := New(context.Background(), nil)

	result, err := c.MsgPushUserMessage(&msg.TLMsgPushUserMessage{
		UserId:    1,
		AuthKeyId: 2,
		PeerType:  tg.PEER_USER,
		PeerId:    3,
		PushType:  1,
		Message:   msg.MakeOutboxMessage(&msg.TLOutboxMessage{}),
	})
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
	if result != tg.BoolTrue {
		t.Fatalf("expected BoolTrue, got %v", result)
	}
}
