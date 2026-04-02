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
