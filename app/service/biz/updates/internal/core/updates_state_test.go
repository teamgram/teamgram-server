package core

import (
	"context"
	"testing"

	"github.com/teamgram/teamgram-server/v2/app/service/biz/updates/updates"
)

func TestUpdatesGetStateV2ReturnsPlaceholderState(t *testing.T) {
	c := New(context.Background(), nil)

	result, err := c.UpdatesGetStateV2(&updates.TLUpdatesGetStateV2{
		AuthKeyId: 1,
		UserId:    2,
	})
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
	if result == nil {
		t.Fatal("expected updates state, got nil")
	}
}

func TestUpdatesGetDifferenceV2ReturnsEmptyDifference(t *testing.T) {
	c := New(context.Background(), nil)

	result, err := c.UpdatesGetDifferenceV2(&updates.TLUpdatesGetDifferenceV2{
		AuthKeyId: 1,
		UserId:    2,
		Pts:       0,
		Date:      0,
	})
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
	if result == nil {
		t.Fatal("expected difference, got nil")
	}
	if _, ok := result.ToDifferenceEmpty(); !ok {
		t.Fatalf("expected differenceEmpty, got %T", result.Clazz)
	}
}

func TestUpdatesGetChannelDifferenceV2ReturnsPlaceholderChannelDifference(t *testing.T) {
	c := New(context.Background(), nil)

	result, err := c.UpdatesGetChannelDifferenceV2(&updates.TLUpdatesGetChannelDifferenceV2{
		AuthKeyId: 1,
		UserId:    2,
		ChannelId: 3,
		Pts:       4,
		Limit:     100,
	})
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
	if result == nil {
		t.Fatal("expected channel difference, got nil")
	}
	channelDiff, ok := result.ToChannelDifference()
	if !ok {
		t.Fatalf("expected channelDifference, got %T", result.Clazz)
	}
	if channelDiff.Pts != 4 {
		t.Fatalf("expected pts=4, got %d", channelDiff.Pts)
	}
	if len(channelDiff.NewMessages) != 0 || len(channelDiff.OtherUpdates) != 0 {
		t.Fatalf("expected empty channel difference payload, got new_messages=%d other_updates=%d",
			len(channelDiff.NewMessages), len(channelDiff.OtherUpdates))
	}
}
