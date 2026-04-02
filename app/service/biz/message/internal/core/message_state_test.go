package core

import (
	"context"
	"testing"

	"github.com/teamgram/teamgram-server/v2/app/service/biz/message/message"
)

func TestMessageGetUserMessageListReturnsStablePlaceholderBoxes(t *testing.T) {
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
	if len(result.Datas) != 2 {
		t.Fatalf("expected 2 placeholder boxes, got %d items", len(result.Datas))
	}

	first := result.Datas[0]
	if first.MessageId != 10 {
		t.Fatalf("expected first message_id=10, got %d", first.MessageId)
	}
	second := result.Datas[1]
	if second.MessageId != 11 {
		t.Fatalf("expected second message_id=11, got %d", second.MessageId)
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
	if result.MessageId != 10 {
		t.Fatalf("expected message_id=10, got %d", result.MessageId)
	}
	if result.UserId != 1 {
		t.Fatalf("expected user_id=1, got %d", result.UserId)
	}
}
