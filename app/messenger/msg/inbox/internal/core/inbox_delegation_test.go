package core

import (
	"context"
	"testing"

	"github.com/teamgram/teamgram-server/v2/app/messenger/msg/inbox/inbox"
	"github.com/teamgram/teamgram-server/v2/app/messenger/msg/inbox/internal/svc"
	synctypes "github.com/teamgram/teamgram-server/v2/app/messenger/sync/sync"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

// --- fake sync push client ---

type fakeInboxSyncClient struct {
	pushReq *synctypes.TLSyncPushUpdates
}

func (f *fakeInboxSyncClient) SyncPushUpdates(ctx context.Context, in *synctypes.TLSyncPushUpdates) (*tg.Void, error) {
	f.pushReq = in
	return tg.MakeTLVoid(&tg.TLVoid{}).ToVoid(), nil
}

var _ svc.SyncPushClient = (*fakeInboxSyncClient)(nil)

func TestInboxSendUserMessageToInboxV2PushesToSync(t *testing.T) {
	fake := &fakeInboxSyncClient{}

	c := New(context.Background(), &svc.ServiceContext{
		SyncClient: fake,
	})

	msg := tg.MakeTLMessage(&tg.TLMessage{
		Id:      123,
		Date:    1000,
		Message: "hello",
		PeerId:  tg.MakeTLPeerUser(&tg.TLPeerUser{UserId: 1}),
	})
	box := &tg.TLMessageBox{
		MessageId: 123,
		Pts:       5,
		PtsCount:  2,
		Message:    msg,
	}

	result, err := c.InboxSendUserMessageToInboxV2(&inbox.TLInboxSendUserMessageToInboxV2{
		UserId:   42,
		Out:      false,
		FromId:   100,
		PeerType: tg.PEER_USER,
		PeerId:   42,
		BoxList:  []tg.MessageBoxClazz{box},
		Users:    []tg.UserClazz{},
		Chats:    []tg.ChatClazz{},
	})
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
	if result == nil {
		t.Fatal("expected non-nil void")
	}

	if fake.pushReq == nil {
		t.Fatal("expected SyncPushUpdates to be called")
	}
	if fake.pushReq.UserId != 42 {
		t.Errorf("expected userId=42, got %d", fake.pushReq.UserId)
	}
	updates, ok := fake.pushReq.Updates.(*tg.TLUpdates)
	if !ok {
		t.Fatalf("expected *tg.TLUpdates, got %T", fake.pushReq.Updates)
	}
	if len(updates.Updates) != 1 {
		t.Fatalf("expected 1 update, got %d", len(updates.Updates))
	}
}

func TestInboxSendUserMessageToInboxV2NoopWithoutSyncClient(t *testing.T) {
	c := New(context.Background(), &svc.ServiceContext{
		SyncClient: nil,
	})

	result, err := c.InboxSendUserMessageToInboxV2(&inbox.TLInboxSendUserMessageToInboxV2{
		UserId:   42,
		BoxList:  []tg.MessageBoxClazz{},
	})
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
	if result == nil {
		t.Fatal("expected non-nil void")
	}
}

func TestInboxSendUserMessageToInboxV2NoopWithNilSvcCtx(t *testing.T) {
	c := New(context.Background(), nil)

	result, err := c.InboxSendUserMessageToInboxV2(&inbox.TLInboxSendUserMessageToInboxV2{
		UserId:   42,
		BoxList:  []tg.MessageBoxClazz{},
	})
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
	if result == nil {
		t.Fatal("expected non-nil void")
	}
}

func TestBuildUpdateNewMessages(t *testing.T) {
	msg := tg.MakeTLMessage(&tg.TLMessage{
		Id:      10,
		Date:    500,
		Message: "test",
	})
	box := &tg.TLMessageBox{
		MessageId: 10,
		Pts:       3,
		PtsCount:  1,
		Message:   msg,
	}
	updates := buildUpdateNewMessages([]tg.MessageBoxClazz{box, nil, box})
	if len(updates) != 2 {
		t.Fatalf("expected 2 updates (nil skipped), got %d", len(updates))
	}
	update0, ok := updates[0].(*tg.TLUpdateNewMessage)
	if !ok {
		t.Fatalf("expected *tg.TLUpdateNewMessage, got %T", updates[0])
	}
	if update0.Pts != 3 || update0.PtsCount != 1 {
		t.Errorf("expected pts=3 ptsCount=1, got pts=%d ptsCount=%d", update0.Pts, update0.PtsCount)
	}
}
