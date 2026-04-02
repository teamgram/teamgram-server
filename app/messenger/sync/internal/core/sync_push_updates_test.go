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

func TestSyncUpdatesMeReturnsVoidPlaceholder(t *testing.T) {
	c := New(context.Background(), nil)

	result, err := c.SyncUpdatesMe(&sync.TLSyncUpdatesMe{
		UserId:        1,
		PermAuthKeyId: 2,
		Updates:       &tg.Updates{},
	})
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
	if result == nil {
		t.Fatal("expected void result, got nil")
	}
}

func TestSyncUpdatesNotMeReturnsVoidPlaceholder(t *testing.T) {
	c := New(context.Background(), nil)

	result, err := c.SyncUpdatesNotMe(&sync.TLSyncUpdatesNotMe{
		UserId:        1,
		PermAuthKeyId: 2,
		Updates:       &tg.Updates{},
	})
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
	if result == nil {
		t.Fatal("expected void result, got nil")
	}
}

func TestSyncPushRpcResultReturnsVoidPlaceholder(t *testing.T) {
	c := New(context.Background(), nil)

	result, err := c.SyncPushRpcResult(&sync.TLSyncPushRpcResult{
		UserId:         1,
		AuthKeyId:      2,
		PermAuthKeyId:  3,
		ServerId:       "srv-1",
		SessionId:      4,
		ClientReqMsgId: 5,
		RpcResult:      []byte{1, 2, 3},
	})
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
	if result == nil {
		t.Fatal("expected void result, got nil")
	}
}

func TestSyncBroadcastUpdatesReturnsVoidPlaceholder(t *testing.T) {
	c := New(context.Background(), nil)

	result, err := c.SyncBroadcastUpdates(&sync.TLSyncBroadcastUpdates{
		BroadcastType: 1,
		ChatId:        2,
		ExcludeIdList: []int64{3, 4},
		Updates:       &tg.Updates{},
	})
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
	if result == nil {
		t.Fatal("expected void result, got nil")
	}
}

func TestSyncPushBotUpdatesReturnsVoidPlaceholder(t *testing.T) {
	c := New(context.Background(), nil)

	result, err := c.SyncPushBotUpdates(&sync.TLSyncPushBotUpdates{
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
