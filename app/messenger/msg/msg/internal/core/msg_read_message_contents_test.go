package core

import (
	"context"
	"testing"

	"github.com/teamgram/teamgram-server/v2/app/messenger/msg/msg/msg"
)

func TestMsgReadMessageContentsReturnsAffectedMessagesPlaceholder(t *testing.T) {
	c := New(context.Background(), nil)

	result, err := c.MsgReadMessageContents(&msg.TLMsgReadMessageContents{
		UserId:    1,
		AuthKeyId: 2,
		PeerType:  2,
		PeerId:    3,
		Id: []*msg.ContentMessage{
			msg.MakeContentMessage(&msg.TLContentMessage{Id: 7}),
			msg.MakeContentMessage(&msg.TLContentMessage{Id: 9}),
		},
	})
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
	if result == nil {
		t.Fatal("expected affectedMessages result, got nil")
	}
	if result.Pts != 9 || result.PtsCount != 2 {
		t.Fatalf("expected pts/pts_count=9/2, got %d/%d", result.Pts, result.PtsCount)
	}
}

func TestMsgReadMessageContentsFallsBackForEmptyIDs(t *testing.T) {
	c := New(context.Background(), nil)

	result, err := c.MsgReadMessageContents(&msg.TLMsgReadMessageContents{
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
