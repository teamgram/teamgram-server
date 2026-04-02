package core

import (
	"context"
	"testing"

	"github.com/teamgram/teamgram-server/v2/app/messenger/msg/msg/msg"
)

func TestMsgReadHistoryReturnsAffectedMessagesPlaceholder(t *testing.T) {
	c := New(context.Background(), nil)

	result, err := c.MsgReadHistory(&msg.TLMsgReadHistory{
		UserId:    1,
		AuthKeyId: 2,
		PeerType:  2,
		PeerId:    3,
		MaxId:     11,
	})
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
	if result == nil {
		t.Fatal("expected affectedMessages result, got nil")
	}
	if result.Pts != 11 || result.PtsCount != 1 {
		t.Fatalf("expected pts/pts_count=11/1, got %d/%d", result.Pts, result.PtsCount)
	}
}

func TestMsgReadHistoryV2ReturnsAffectedMessagesPlaceholder(t *testing.T) {
	c := New(context.Background(), nil)

	result, err := c.MsgReadHistoryV2(&msg.TLMsgReadHistoryV2{
		UserId:    1,
		AuthKeyId: 2,
		PeerType:  2,
		PeerId:    3,
		MaxId:     0,
	})
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
	if result == nil {
		t.Fatal("expected affectedMessages result, got nil")
	}
	if result.Pts != 1 || result.PtsCount != 1 {
		t.Fatalf("expected fallback pts/pts_count=1/1, got %d/%d", result.Pts, result.PtsCount)
	}
}
