package core

import (
	"context"
	"testing"

	"github.com/teamgram/teamgram-server/v2/app/messenger/msg/msg/msg"
)

func TestMsgDeleteMessagesReturnsAffectedMessagesPlaceholder(t *testing.T) {
	c := New(context.Background(), nil)

	result, err := c.MsgDeleteMessages(&msg.TLMsgDeleteMessages{
		UserId:    1,
		AuthKeyId: 2,
		PeerType:  2,
		PeerId:    3,
		Id:        []int32{4, 8},
	})
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
	if result == nil {
		t.Fatal("expected affectedMessages result, got nil")
	}
	if result.Pts != 8 || result.PtsCount != 2 {
		t.Fatalf("expected pts/pts_count=8/2, got %d/%d", result.Pts, result.PtsCount)
	}
}

func TestMsgDeleteMessagesFallsBackForEmptyIDs(t *testing.T) {
	c := New(context.Background(), nil)

	result, err := c.MsgDeleteMessages(&msg.TLMsgDeleteMessages{
		UserId:    1,
		AuthKeyId: 2,
		PeerType:  2,
		PeerId:    3,
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
