package core

import (
	"context"
	"testing"

	"github.com/teamgram/teamgram-server/v2/app/messenger/msg/msg/msg"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

func TestMsgEditMessageV2ReturnsShortSentPlaceholderForUserPeer(t *testing.T) {
	c := New(context.Background(), nil)

	result, err := c.MsgEditMessageV2(&msg.TLMsgEditMessageV2{
		UserId:    1,
		AuthKeyId: 2,
		PeerType:  tg.PEER_USER,
		PeerId:    3,
		NewMessage: msg.MakeOutboxMessage(&msg.TLOutboxMessage{
			RandomId: 1005,
			Message:  tg.MakeTLMessage(&tg.TLMessage{Date: 123, Message: "edited"}).ToMessage(),
		}),
		DstMessage: tg.MakeTLMessageBox(&tg.TLMessageBox{
			UserId:    1,
			MessageId: 77,
			PeerType:  tg.PEER_USER,
			PeerId:    3,
		}).ToMessageBox(),
	})
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
	short, ok := result.ToUpdateShortSentMessage()
	if !ok {
		t.Fatalf("expected updateShortSentMessage, got %T", result.Clazz)
	}
	if short.Id != 77 {
		t.Fatalf("expected edited message id=77, got %d", short.Id)
	}
	if short.Date != 123 {
		t.Fatalf("expected date=123, got %d", short.Date)
	}
}

func TestMsgEditMessageV2RejectsMissingTargetMessage(t *testing.T) {
	c := New(context.Background(), nil)

	result, err := c.MsgEditMessageV2(&msg.TLMsgEditMessageV2{
		UserId:     1,
		AuthKeyId:  2,
		PeerType:   tg.PEER_USER,
		PeerId:     3,
		NewMessage: msg.MakeOutboxMessage(&msg.TLOutboxMessage{RandomId: 1005}),
	})
	if err != tg.ErrInputRequestInvalid {
		t.Fatalf("expected ErrInputRequestInvalid, got %v", err)
	}
	if result != nil {
		t.Fatalf("expected nil result, got %v", result)
	}
}

func TestMsgEditMessageV2RejectsChannelPeerPlaceholder(t *testing.T) {
	c := New(context.Background(), nil)

	result, err := c.MsgEditMessageV2(&msg.TLMsgEditMessageV2{
		UserId:    1,
		AuthKeyId: 2,
		PeerType:  tg.PEER_CHANNEL,
		PeerId:    3,
		NewMessage: msg.MakeOutboxMessage(&msg.TLOutboxMessage{
			RandomId: 1006,
		}),
		DstMessage: tg.MakeTLMessageBox(&tg.TLMessageBox{
			UserId:    1,
			MessageId: 88,
			PeerType:  tg.PEER_CHANNEL,
			PeerId:    3,
		}).ToMessageBox(),
	})
	if err != tg.ErrEnterpriseIsBlocked {
		t.Fatalf("expected ErrEnterpriseIsBlocked, got %v", err)
	}
	if result != nil {
		t.Fatalf("expected nil result, got %v", result)
	}
}
