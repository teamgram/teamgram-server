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

func TestMsgSendMessageV2ReturnsShortSentMessagePlaceholderForUserPeer(t *testing.T) {
	c := New(context.Background(), nil)

	result, err := c.MsgSendMessageV2(&msg.TLMsgSendMessageV2{
		UserId:    1,
		AuthKeyId: 2,
		PeerType:  tg.PEER_USER,
		PeerId:    3,
		Message: []*msg.OutboxMessage{
			msg.MakeTLOutboxMessage(&msg.TLOutboxMessage{
				RandomId: 1001,
				Message:  tg.MakeTLMessage(&tg.TLMessage{Date: 12345}),
			}),
		},
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
	if updates.Id != 1001 {
		t.Fatalf("expected placeholder id=1001, got %d", updates.Id)
	}
	if updates.Date != 12345 {
		t.Fatalf("expected date=12345 from outbox message, got %d", updates.Date)
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
			msg.MakeTLOutboxMessage(&msg.TLOutboxMessage{
				RandomId: 1002,
				Message:  tg.MakeTLMessage(&tg.TLMessage{Message: "x"}),
			}),
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
			msg.MakeTLOutboxMessage(&msg.TLOutboxMessage{
				RandomId: 1003,
				Message:  tg.MakeTLMessage(&tg.TLMessage{Message: "x"}),
			}),
		},
	})
	if err != tg.ErrEnterpriseIsBlocked {
		t.Fatalf("expected ErrEnterpriseIsBlocked, got %v", err)
	}
	if result != nil {
		t.Fatalf("expected nil result, got %v", result)
	}
}

func TestMsgSendMessageV2ReusesPlaceholderIDForSameRandomID(t *testing.T) {
	c := New(context.Background(), nil)

	req := &msg.TLMsgSendMessageV2{
		UserId:    1,
		AuthKeyId: 2,
		PeerType:  tg.PEER_USER,
		PeerId:    3,
		Message: []*msg.OutboxMessage{
			msg.MakeTLOutboxMessage(&msg.TLOutboxMessage{
				RandomId: 2007,
				Message:  tg.MakeTLMessage(&tg.TLMessage{Message: "x"}),
			}),
		},
	}

	first, err := c.MsgSendMessageV2(req)
	if err != nil {
		t.Fatalf("first send: expected nil error, got %v", err)
	}
	second, err := c.MsgSendMessageV2(req)
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
