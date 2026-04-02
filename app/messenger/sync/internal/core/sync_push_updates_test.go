package core

import (
	"context"
	"testing"

	"github.com/teamgram/teamgram-server/v2/app/messenger/sync/sync"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

func TestSyncPushUpdatesReturnsVoidPlaceholder(t *testing.T) {
	c := New(context.Background(), nil)

	result, err := c.SyncPushUpdates(&sync.TLSyncPushUpdates{
		UserId:  1,
		Updates: &tg.Updates{},
	})
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
	if result == nil {
		t.Fatal("expected void result, got nil")
	}
}

func TestSyncPushUpdatesIfNotReturnsVoidPlaceholder(t *testing.T) {
	c := New(context.Background(), nil)

	result, err := c.SyncPushUpdatesIfNot(&sync.TLSyncPushUpdatesIfNot{
		UserId:   1,
		Excludes: []int64{2, 3},
		Updates:  &tg.Updates{},
	})
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
	if result == nil {
		t.Fatal("expected void result, got nil")
	}
}
