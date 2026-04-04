package core

import (
	"context"
	"testing"

	"github.com/teamgram/teamgram-server/v2/app/messenger/msg/msg/internal/svc"
	"github.com/teamgram/teamgram-server/v2/app/messenger/msg/msg/msg"
	synctypes "github.com/teamgram/teamgram-server/v2/app/messenger/sync/sync"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

// --- fake sync client for send tests ---

type fakeSyncClient struct {
	updatesMeReq    *synctypes.TLSyncUpdatesMe
	updatesNotMeReq *synctypes.TLSyncUpdatesNotMe
	pushUpdatesReq  *synctypes.TLSyncPushUpdates
}

func (f *fakeSyncClient) SyncUpdatesMe(ctx context.Context, in *synctypes.TLSyncUpdatesMe) (*tg.Void, error) {
	f.updatesMeReq = in
	return tg.MakeTLVoid(&tg.TLVoid{}).ToVoid(), nil
}

func (f *fakeSyncClient) SyncUpdatesNotMe(ctx context.Context, in *synctypes.TLSyncUpdatesNotMe) (*tg.Void, error) {
	f.updatesNotMeReq = in
	return tg.MakeTLVoid(&tg.TLVoid{}).ToVoid(), nil
}

func (f *fakeSyncClient) SyncPushUpdates(ctx context.Context, in *synctypes.TLSyncPushUpdates) (*tg.Void, error) {
	f.pushUpdatesReq = in
	return tg.MakeTLVoid(&tg.TLVoid{}).ToVoid(), nil
}

var _ svc.SyncPushClient = (*fakeSyncClient)(nil)

// --- basic tests ---

func TestMsgSendMessageV2RejectsEmptyOutboxList(t *testing.T) {
	c := New(context.Background(), nil)

	result, err := c.MsgSendMessageV2(&msg.TLMsgSendMessageV2{
		UserId:    1,
		AuthKeyId: 2,
		PeerType:  tg.PEER_USER,
		PeerId:    3,
	})
	if err != tg.ErrInputRequestInvalid {
		t.Fatalf("expected ErrInputRequestInvalid, got %v", err)
	}
	if result != nil {
		t.Fatalf("expected nil result, got %v", result)
	}
}

func TestMsgSendMessageV2ReturnsShortSentMessagePlaceholderForUserPeer(t *testing.T) {
	c := New(context.Background(), nil)

	result, err := c.MsgSendMessageV2(&msg.TLMsgSendMessageV2{
		UserId:    1,
		AuthKeyId: 2,
		PeerType:  tg.PEER_USER,
		PeerId:    3,
		Message: []*msg.OutboxMessage{
			msg.MakeTLOutboxMessage(&msg.TLOutboxMessage{
				RandomId: 1001,
				Message:  tg.MakeTLMessage(&tg.TLMessage{Date: 12345}),
			}),
		},
	})
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
	if result == nil {
		t.Fatal("expected updates result, got nil")
	}

	updates, ok := result.ToUpdateShortSentMessage()
	if !ok {
		t.Fatalf("expected updateShortSentMessage placeholder, got %T", result.Clazz)
	}
	if !updates.Out {
		t.Fatal("expected out=true")
	}
	if updates.Id != 1001 {
		t.Fatalf("expected placeholder id=1001, got %d", updates.Id)
	}
	if updates.Date != 12345 {
		t.Fatalf("expected date=12345 from outbox message, got %d", updates.Date)
	}
}

func TestMsgSendMessageV2RejectsInvalidPeerType(t *testing.T) {
	c := New(context.Background(), nil)

	result, err := c.MsgSendMessageV2(&msg.TLMsgSendMessageV2{
		UserId:    1,
		AuthKeyId: 2,
		PeerType:  99,
		PeerId:    3,
		Message: []*msg.OutboxMessage{
			msg.MakeTLOutboxMessage(&msg.TLOutboxMessage{
				RandomId: 1002,
				Message:  tg.MakeTLMessage(&tg.TLMessage{Message: "x"}),
			}),
		},
	})
	if err != tg.ErrPeerIdInvalid {
		t.Fatalf("expected ErrPeerIdInvalid, got %v", err)
	}
	if result != nil {
		t.Fatalf("expected nil result, got %v", result)
	}
}

func TestMsgSendMessageV2RejectsChannelPeerPlaceholder(t *testing.T) {
	c := New(context.Background(), nil)

	result, err := c.MsgSendMessageV2(&msg.TLMsgSendMessageV2{
		UserId:    1,
		AuthKeyId: 2,
		PeerType:  tg.PEER_CHANNEL,
		PeerId:    3,
		Message: []*msg.OutboxMessage{
			msg.MakeTLOutboxMessage(&msg.TLOutboxMessage{
				RandomId: 1003,
				Message:  tg.MakeTLMessage(&tg.TLMessage{Message: "x"}),
			}),
		},
	})
	if err != tg.ErrEnterpriseIsBlocked {
		t.Fatalf("expected ErrEnterpriseIsBlocked, got %v", err)
	}
	if result != nil {
		t.Fatalf("expected nil result, got %v", result)
	}
}

func TestMsgSendMessageV2ReusesPlaceholderIDForSameRandomID(t *testing.T) {
	c := New(context.Background(), nil)

	req := &msg.TLMsgSendMessageV2{
		UserId:    1,
		AuthKeyId: 2,
		PeerType:  tg.PEER_USER,
		PeerId:    3,
		Message: []*msg.OutboxMessage{
			msg.MakeTLOutboxMessage(&msg.TLOutboxMessage{
				RandomId: 2007,
				Message:  tg.MakeTLMessage(&tg.TLMessage{Message: "x"}),
			}),
		},
	}

	first, err := c.MsgSendMessageV2(req)
	if err != nil {
		t.Fatalf("first send: expected nil error, got %v", err)
	}
	second, err := c.MsgSendMessageV2(req)
	if err != nil {
		t.Fatalf("second send: expected nil error, got %v", err)
	}

	firstShort, ok := first.ToUpdateShortSentMessage()
	if !ok {
		t.Fatalf("expected first result to be updateShortSentMessage, got %T", first.Clazz)
	}
	secondShort, ok := second.ToUpdateShortSentMessage()
	if !ok {
		t.Fatalf("expected second result to be updateShortSentMessage, got %T", second.Clazz)
	}
	if firstShort.Id != secondShort.Id {
		t.Fatalf("expected same placeholder id for repeated random_id, got %d vs %d", firstShort.Id, secondShort.Id)
	}
}

// --- fanout side-effect tests ---

func TestMsgSendMessageV2PushesToInboxAndSync(t *testing.T) {
	fakeInbox := &fakeInboxClient{}
	fakeSync := &fakeSyncClient{}
	c := New(context.Background(), &svc.ServiceContext{
		InboxClient: fakeInbox,
		SyncClient:  fakeSync,
	})

	result, err := c.MsgSendMessageV2(&msg.TLMsgSendMessageV2{
		UserId:    10,
		AuthKeyId: 20,
		PeerType:  tg.PEER_USER,
		PeerId:    30,
		Message: []*msg.OutboxMessage{
			msg.MakeTLOutboxMessage(&msg.TLOutboxMessage{
				RandomId: 5001,
				Message:  tg.MakeTLMessage(&tg.TLMessage{Date: 99, Message: "hello"}),
			}),
		},
	})
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
	if result == nil {
		t.Fatal("expected non-nil result")
	}

	// Inbox should receive the message for recipient.
	if fakeInbox.req == nil {
		t.Fatal("expected inbox push for recipient")
	}
	if fakeInbox.req.UserId != 30 {
		t.Errorf("inbox userId = %d, want 30", fakeInbox.req.UserId)
	}
	if fakeInbox.req.FromId != 10 {
		t.Errorf("inbox fromId = %d, want 10", fakeInbox.req.FromId)
	}

	// Sync should push to sender's other sessions.
	if fakeSync.updatesNotMeReq == nil {
		t.Fatal("expected sync updatesNotMe for sender")
	}
	if fakeSync.updatesNotMeReq.UserId != 10 {
		t.Errorf("sync updatesNotMe userId = %d, want 10", fakeSync.updatesNotMeReq.UserId)
	}

	// Sync should push updates to recipient.
	if fakeSync.pushUpdatesReq == nil {
		t.Fatal("expected sync pushUpdates for recipient")
	}
	if fakeSync.pushUpdatesReq.UserId != 30 {
		t.Errorf("sync pushUpdates userId = %d, want 30", fakeSync.pushUpdatesReq.UserId)
	}
}

func TestMsgSendMessageV2SelfPeerSkipsInboxAndRecipientSync(t *testing.T) {
	fakeInbox := &fakeInboxClient{}
	fakeSync := &fakeSyncClient{}
	c := New(context.Background(), &svc.ServiceContext{
		InboxClient: fakeInbox,
		SyncClient:  fakeSync,
	})

	result, err := c.MsgSendMessageV2(&msg.TLMsgSendMessageV2{
		UserId:    10,
		AuthKeyId: 20,
		PeerType:  tg.PEER_USER,
		PeerId:    10, // self
		Message: []*msg.OutboxMessage{
			msg.MakeTLOutboxMessage(&msg.TLOutboxMessage{
				RandomId: 5002,
				Message:  tg.MakeTLMessage(&tg.TLMessage{Date: 100, Message: "self"}),
			}),
		},
	})
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
	if result == nil {
		t.Fatal("expected non-nil result")
	}

	// Inbox should NOT be called for self peer.
	if fakeInbox.req != nil {
		t.Fatal("expected no inbox push for self peer")
	}

	// Sync updatesNotMe should still fire for sender's other sessions.
	if fakeSync.updatesNotMeReq == nil {
		t.Fatal("expected sync updatesNotMe for self send")
	}

	// Sync pushUpdates to recipient should NOT fire for self.
	if fakeSync.pushUpdatesReq != nil {
		t.Fatal("expected no recipient sync push for self peer")
	}
}
