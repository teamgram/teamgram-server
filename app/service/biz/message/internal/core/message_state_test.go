package core

import (
	"context"
	"testing"

	"github.com/teamgram/teamgram-server/v2/app/service/biz/message/message"
)

func TestMessageGetUserMessageListReturnsEmptyPlaceholder(t *testing.T) {
	c := New(context.Background(), nil)

	result, err := c.MessageGetUserMessageList(&message.TLMessageGetUserMessageList{
		UserId: 1,
		IdList: []int32{10, 11},
	})
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
	if result == nil {
		t.Fatal("expected message box vector, got nil")
	}
	if len(result.Datas) != 0 {
		t.Fatalf("expected empty placeholder list, got %d items", len(result.Datas))
	}
}

func TestMessageGetUserMessageReturnsPlaceholderMessageBox(t *testing.T) {
	c := New(context.Background(), nil)

	result, err := c.MessageGetUserMessage(&message.TLMessageGetUserMessage{
		UserId: 1,
		Id:     10,
	})
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
	if result == nil {
		t.Fatal("expected message box, got nil")
	}
}
