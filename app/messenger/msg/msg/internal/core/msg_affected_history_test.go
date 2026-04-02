package core

import (
	"context"
	"testing"

	"github.com/teamgram/teamgram-server/v2/app/messenger/msg/msg/msg"
)

func TestMsgDeleteHistoryReturnsAffectedHistoryPlaceholder(t *testing.T) {
	c := New(context.Background(), nil)

	result, err := c.MsgDeleteHistory(&msg.TLMsgDeleteHistory{
		UserId:    1,
		AuthKeyId: 2,
		PeerType:  2,
		PeerId:    3,
		MaxId:     15,
	})
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
	if result == nil {
		t.Fatal("expected affectedHistory result, got nil")
	}
	if result.Pts != 15 || result.PtsCount != 1 || result.Offset != 1 {
		t.Fatalf("expected pts/pts_count/offset=15/1/1, got %d/%d/%d", result.Pts, result.PtsCount, result.Offset)
	}
}

func TestMsgUnpinAllMessagesReturnsAffectedHistoryPlaceholder(t *testing.T) {
	c := New(context.Background(), nil)

	result, err := c.MsgUnpinAllMessages(&msg.TLMsgUnpinAllMessages{
		UserId:    1,
		AuthKeyId: 2,
		PeerType:  2,
		PeerId:    3,
	})
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
	if result == nil {
		t.Fatal("expected affectedHistory result, got nil")
	}
	if result.Pts != 1 || result.PtsCount != 1 || result.Offset != 0 {
		t.Fatalf("expected pts/pts_count/offset=1/1/0, got %d/%d/%d", result.Pts, result.PtsCount, result.Offset)
	}
}
