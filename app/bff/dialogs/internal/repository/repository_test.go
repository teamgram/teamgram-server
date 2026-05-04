package repository

import (
	"context"
	"testing"

	syncclient "github.com/teamgram/teamgram-server/v2/app/messenger/sync/client"
	syncpb "github.com/teamgram/teamgram-server/v2/app/messenger/sync/sync"
	"github.com/teamgram/teamgram-server/v2/pkg/net/kitex"
	"github.com/teamgram/teamgram-server/v2/pkg/net/kitex/identity"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

type fakeSyncClient struct {
	syncclient.SyncClient
	callerService string
	userID        int64
	updates       tg.UpdatesClazz
}

func (f *fakeSyncClient) SyncPushUpdates(ctx context.Context, in *syncpb.TLSyncPushUpdates) (*tg.Void, error) {
	f.callerService, _ = identity.CallerService(ctx)
	f.userID = in.UserId
	f.updates = in.Updates
	return &tg.Void{}, nil
}

func TestPushTypingUpdatesSetsCallerService(t *testing.T) {
	syncClient := &fakeSyncClient{}
	updates := tg.MakeTLUpdateShort(&tg.TLUpdateShort{
		Update: tg.MakeTLUpdateUserTyping(&tg.TLUpdateUserTyping{
			UserId: 1001,
			Action: tg.MakeTLSendMessageTypingAction(&tg.TLSendMessageTypingAction{}),
		}),
		Date: 123,
	})

	err := (&Repository{SyncClient: syncClient}).PushTypingUpdates(context.Background(), 1002, updates)
	if err != nil {
		t.Fatalf("PushTypingUpdates() error = %v", err)
	}
	if syncClient.callerService != "bff.dialogs" {
		t.Fatalf("callerService = %q, want bff.dialogs", syncClient.callerService)
	}
	if syncClient.userID != 1002 {
		t.Fatalf("userID = %d, want 1002", syncClient.userID)
	}
	if syncClient.updates != updates {
		t.Fatal("updates was not forwarded")
	}
}

func TestHasSyncClientConfigRequiresDestAndServiceName(t *testing.T) {
	if hasSyncClientConfig(kitex.RpcClientConf{}) {
		t.Fatal("empty sync config should be disabled")
	}
	if hasSyncClientConfig(kitex.RpcClientConf{DestService: "messenger.sync"}) {
		t.Fatal("sync config without service name should be disabled")
	}
	if !hasSyncClientConfig(kitex.RpcClientConf{DestService: "messenger.sync", ServiceName: "RPCSync"}) {
		t.Fatal("sync config with dest and service name should be enabled")
	}
}
