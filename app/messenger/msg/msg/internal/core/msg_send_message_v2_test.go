package core

import (
	"context"
	"testing"

	"github.com/teamgram/teamgram-server/v2/app/messenger/msg/msg/msg"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

func TestMsgSendMessageV2RejectsEmptyOutboxList(t *testing.T) {
	c := New(context.Background(), nil)

	result, err := c.MsgSendMessageV2(&msg.TLMsgSendMessageV2{
		UserId:    1,
		AuthKeyId: 2,
		PeerType:  tg.PEER_USER,
		PeerId:    3,
	})
	if err != tg.ErrInputRequestInvalid {
		t.Fatalf("expected ErrInputRequestInvalid, got %v", err)
	}
	if result != nil {
		t.Fatalf("expected nil result, got %v", result)
	}
}

func TestMsgSendMessageV2ReturnsEmptyUpdatesPlaceholderForUserPeer(t *testing.T) {
	c := New(context.Background(), nil)

	result, err := c.MsgSendMessageV2(&msg.TLMsgSendMessageV2{
		UserId:    1,
		AuthKeyId: 2,
		PeerType:  tg.PEER_USER,
		PeerId:    3,
		Message: []*msg.OutboxMessage{
			msg.MakeOutboxMessage(&msg.TLOutboxMessage{RandomId: 1001}),
		},
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

func TestMsgSendMessageV2RejectsInvalidPeerType(t *testing.T) {
	c := New(context.Background(), nil)

	result, err := c.MsgSendMessageV2(&msg.TLMsgSendMessageV2{
		UserId:    1,
		AuthKeyId: 2,
		PeerType:  99,
		PeerId:    3,
		Message: []*msg.OutboxMessage{
			msg.MakeOutboxMessage(&msg.TLOutboxMessage{RandomId: 1002}),
		},
	})
	if err != tg.ErrPeerIdInvalid {
		t.Fatalf("expected ErrPeerIdInvalid, got %v", err)
	}
	if result != nil {
		t.Fatalf("expected nil result, got %v", result)
	}
}

func TestMsgSendMessageV2RejectsChannelPeerPlaceholder(t *testing.T) {
	c := New(context.Background(), nil)

	result, err := c.MsgSendMessageV2(&msg.TLMsgSendMessageV2{
		UserId:    1,
		AuthKeyId: 2,
		PeerType:  tg.PEER_CHANNEL,
		PeerId:    3,
		Message: []*msg.OutboxMessage{
			msg.MakeOutboxMessage(&msg.TLOutboxMessage{RandomId: 1003}),
		},
	})
	if err != tg.ErrEnterpriseIsBlocked {
		t.Fatalf("expected ErrEnterpriseIsBlocked, got %v", err)
	}
	if result != nil {
		t.Fatalf("expected nil result, got %v", result)
	}
}
