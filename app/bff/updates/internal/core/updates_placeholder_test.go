package core

import (
	"context"
	"testing"

	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

func TestUpdatesGetStateReturnsPlaceholderState(t *testing.T) {
	c := New(context.Background(), nil)

	result, err := c.UpdatesGetState(&tg.TLUpdatesGetState{})
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
	if result == nil || result.Pts != 1 || result.Date != 10 {
		t.Fatalf("expected placeholder state pts=1 date=10, got %#v", result)
	}
}

func TestUpdatesGetDifferenceReturnsCatchUpPayloadForBehindClient(t *testing.T) {
	c := New(context.Background(), nil)

	result, err := c.UpdatesGetDifference(&tg.TLUpdatesGetDifference{
		Pts:  0,
		Date: 0,
	})
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
	diff, ok := result.ToUpdatesDifference()
	if !ok {
		t.Fatalf("expected updates.difference, got %T", result.Clazz)
	}
	if diff.State == nil || diff.State.Pts != 1 || diff.State.Date != 10 {
		t.Fatalf("expected placeholder state pts=1 date=10, got %#v", diff.State)
	}
	if len(diff.NewMessages) != 1 || len(diff.OtherUpdates) != 1 {
		t.Fatalf("expected single catch-up payload, got messages=%d updates=%d", len(diff.NewMessages), len(diff.OtherUpdates))
	}
}

func TestUpdatesGetDifferenceReturnsDifferenceEmptyForForwardClient(t *testing.T) {
	c := New(context.Background(), nil)

	result, err := c.UpdatesGetDifference(&tg.TLUpdatesGetDifference{
		Pts:  7,
		Date: 99,
	})
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
	diff, ok := result.ToUpdatesDifferenceEmpty()
	if !ok {
		t.Fatalf("expected updates.differenceEmpty, got %T", result.Clazz)
	}
	if diff.Date != 99 || diff.Seq != 0 {
		t.Fatalf("expected placeholder differenceEmpty date=99 seq=0, got %#v", diff)
	}
}

func TestUpdatesGetChannelDifferenceReturnsPlaceholder(t *testing.T) {
	c := New(context.Background(), nil)

	result, err := c.UpdatesGetChannelDifference(&tg.TLUpdatesGetChannelDifference{
		Pts: 4,
	})
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
	diff, ok := result.ToUpdatesChannelDifferenceEmpty()
	if !ok {
		t.Fatalf("expected updates.channelDifferenceEmpty, got %T", result.Clazz)
	}
	if !diff.Final || diff.Pts != 4 {
		t.Fatalf("expected final placeholder pts=4, got %#v", diff)
	}
}
