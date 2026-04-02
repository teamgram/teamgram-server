package core

import (
	"context"
	"testing"

	"github.com/teamgram/teamgram-server/v2/app/messenger/msg/msg/msg"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

func TestMsgDeleteChatHistoryReturnsBoolTruePlaceholder(t *testing.T) {
	c := New(context.Background(), nil)

	result, err := c.MsgDeleteChatHistory(&msg.TLMsgDeleteChatHistory{
		ChatId:       10,
		DeleteUserId: 1,
	})
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
	if result != tg.BoolTrue {
		t.Fatalf("expected BoolTrue, got %v", result)
	}
}

func TestMsgDeletePhoneCallHistoryReturnsAffectedFoundMessagesPlaceholder(t *testing.T) {
	c := New(context.Background(), nil)

	result, err := c.MsgDeletePhoneCallHistory(&msg.TLMsgDeletePhoneCallHistory{
		UserId:    1,
		AuthKeyId: 2,
	})
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
	if result == nil {
		t.Fatal("expected affectedFoundMessages result, got nil")
	}
	if result.Pts != 1 || result.PtsCount != 1 || result.Offset != 0 {
		t.Fatalf("expected pts/pts_count/offset=1/1/0, got %d/%d/%d", result.Pts, result.PtsCount, result.Offset)
	}
	if len(result.Messages) != 1 || result.Messages[0] != 1 {
		t.Fatalf("expected placeholder messages [1], got %#v", result.Messages)
	}
}
