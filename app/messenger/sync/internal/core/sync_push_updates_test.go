package core

import (
	"context"
	"testing"

	sessionclient "github.com/teamgram/teamgram-server/v2/app/interface/session/client"
	"github.com/teamgram/teamgram-server/v2/app/interface/session/session"
	"github.com/teamgram/teamgram-server/v2/app/messenger/sync/internal/svc"
	"github.com/teamgram/teamgram-server/v2/app/messenger/sync/sync"
	"github.com/teamgram/teamgram-server/v2/app/service/status/status"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

// --- fake session client ---

type fakeSessionClient struct {
	pushReq            *session.TLSessionPushRpcResultData
	pushUpdatesReq     *session.TLSessionPushUpdatesData
	pushSessionUpdates *session.TLSessionPushSessionUpdatesData
	// track all pushUpdatesData calls for fanout tests
	allPushUpdates []*session.TLSessionPushUpdatesData
}

func (f *fakeSessionClient) SessionQueryAuthKey(ctx context.Context, in *session.TLSessionQueryAuthKey) (*tg.AuthKeyInfo, error) {
	panic("unexpected call")
}

func (f *fakeSessionClient) SessionSetAuthKey(ctx context.Context, in *session.TLSessionSetAuthKey) (*tg.Bool, error) {
	panic("unexpected call")
}

func (f *fakeSessionClient) SessionCreateSession(ctx context.Context, in *session.TLSessionCreateSession) (*tg.Bool, error) {
	panic("unexpected call")
}

func (f *fakeSessionClient) SessionSendDataToSession(ctx context.Context, in *session.TLSessionSendDataToSession) (*tg.Bool, error) {
	panic("unexpected call")
}

func (f *fakeSessionClient) SessionSendHttpDataToSession(ctx context.Context, in *session.TLSessionSendHttpDataToSession) (*session.HttpSessionData, error) {
	panic("unexpected call")
}

func (f *fakeSessionClient) SessionCloseSession(ctx context.Context, in *session.TLSessionCloseSession) (*tg.Bool, error) {
	panic("unexpected call")
}

func (f *fakeSessionClient) SessionPushUpdatesData(ctx context.Context, in *session.TLSessionPushUpdatesData) (*tg.Bool, error) {
	f.pushUpdatesReq = in
	f.allPushUpdates = append(f.allPushUpdates, in)
	return tg.BoolTrue, nil
}

func (f *fakeSessionClient) SessionPushSessionUpdatesData(ctx context.Context, in *session.TLSessionPushSessionUpdatesData) (*tg.Bool, error) {
	f.pushSessionUpdates = in
	return tg.BoolTrue, nil
}

func (f *fakeSessionClient) SessionPushRpcResultData(ctx context.Context, in *session.TLSessionPushRpcResultData) (*tg.Bool, error) {
	f.pushReq = in
	return tg.BoolTrue, nil
}

func (f *fakeSessionClient) Close() error { return nil }

var _ sessionclient.SessionClient = (*fakeSessionClient)(nil)

// --- fake status client ---

type fakeStatusClient struct {
	sessions []*status.TLSessionEntry
}

func (f *fakeStatusClient) StatusGetUserOnlineSessions(ctx context.Context, in *status.TLStatusGetUserOnlineSessions) (*status.UserSessionEntryList, error) {
	return &status.UserSessionEntryList{
		UserSessions: f.sessions,
	}, nil
}

var _ svc.StatusQueryClient = (*fakeStatusClient)(nil)

// --- helpers ---

func makeSyncPlaceholderUpdates() *tg.TLUpdates {
	return tg.MakeTLUpdates(&tg.TLUpdates{
		Updates: []tg.UpdateClazz{},
		Users:   []tg.UserClazz{},
		Chats:   []tg.ChatClazz{},
		Date:    1,
		Seq:     1,
	})
}

func makeSessionEntry(authKeyId, permAuthKeyId int64, gateway string) *status.TLSessionEntry {
	return &status.TLSessionEntry{
		AuthKeyId:     authKeyId,
		PermAuthKeyId: permAuthKeyId,
		Gateway:       gateway,
	}
}

// --- pushUpdates tests ---

func TestSyncPushUpdatesReturnsVoidWithNilContext(t *testing.T) {
	c := New(context.Background(), nil)

	result, err := c.SyncPushUpdates(&sync.TLSyncPushUpdates{
		UserId:  1,
		Updates: makeSyncPlaceholderUpdates(),
	})
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
	if result == nil {
		t.Fatal("expected void result, got nil")
	}
}

func TestSyncPushUpdatesFansOutToAllOnlineSessions(t *testing.T) {
	fakeCli := &fakeSessionClient{}
	fakeStatus := &fakeStatusClient{
		sessions: []*status.TLSessionEntry{
			makeSessionEntry(100, 10, "gw-1"),
			makeSessionEntry(200, 20, "gw-2"),
			makeSessionEntry(300, 30, "gw-1"),
		},
	}
	c := New(context.Background(), &svc.ServiceContext{
		SessionClient: fakeCli,
		StatusClient:  fakeStatus,
	})

	updates := makeSyncPlaceholderUpdates()
	result, err := c.SyncPushUpdates(&sync.TLSyncPushUpdates{
		UserId:  1,
		Updates: updates,
	})
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
	if result == nil {
		t.Fatal("expected void result, got nil")
	}
	if len(fakeCli.allPushUpdates) != 3 {
		t.Fatalf("expected 3 push calls, got %d", len(fakeCli.allPushUpdates))
	}

	wantKeys := []int64{10, 20, 30}
	for i, want := range wantKeys {
		if fakeCli.allPushUpdates[i].PermAuthKeyId != want {
			t.Errorf("push[%d] permAuthKeyId = %d, want %d", i, fakeCli.allPushUpdates[i].PermAuthKeyId, want)
		}
		if fakeCli.allPushUpdates[i].Updates != updates {
			t.Errorf("push[%d] updates object mismatch", i)
		}
	}
}

func TestSyncPushUpdatesNoopWithoutStatusClient(t *testing.T) {
	fakeCli := &fakeSessionClient{}
	c := New(context.Background(), &svc.ServiceContext{
		SessionClient: fakeCli,
	})

	result, err := c.SyncPushUpdates(&sync.TLSyncPushUpdates{
		UserId:  1,
		Updates: makeSyncPlaceholderUpdates(),
	})
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
	if result == nil {
		t.Fatal("expected void result, got nil")
	}
	if len(fakeCli.allPushUpdates) != 0 {
		t.Fatalf("expected no push calls without StatusClient, got %d", len(fakeCli.allPushUpdates))
	}
}

// --- pushUpdatesIfNot tests ---

func TestSyncPushUpdatesIfNotReturnsVoidWithNilContext(t *testing.T) {
	c := New(context.Background(), nil)

	result, err := c.SyncPushUpdatesIfNot(&sync.TLSyncPushUpdatesIfNot{
		UserId:   1,
		Excludes: []int64{2, 3},
		Updates:  makeSyncPlaceholderUpdates(),
	})
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
	if result == nil {
		t.Fatal("expected void result, got nil")
	}
}

func TestSyncPushUpdatesIfNotExcludesSessions(t *testing.T) {
	fakeCli := &fakeSessionClient{}
	fakeStatus := &fakeStatusClient{
		sessions: []*status.TLSessionEntry{
			makeSessionEntry(100, 10, "gw-1"),
			makeSessionEntry(200, 20, "gw-2"),
			makeSessionEntry(300, 30, "gw-1"),
		},
	}
	c := New(context.Background(), &svc.ServiceContext{
		SessionClient: fakeCli,
		StatusClient:  fakeStatus,
	})

	updates := makeSyncPlaceholderUpdates()
	result, err := c.SyncPushUpdatesIfNot(&sync.TLSyncPushUpdatesIfNot{
		UserId:   1,
		Excludes: []int64{20}, // exclude permAuthKeyId 20
		Updates:  updates,
	})
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
	if result == nil {
		t.Fatal("expected void result, got nil")
	}
	if len(fakeCli.allPushUpdates) != 2 {
		t.Fatalf("expected 2 push calls (excluded 1), got %d", len(fakeCli.allPushUpdates))
	}
	if fakeCli.allPushUpdates[0].PermAuthKeyId != 10 {
		t.Errorf("push[0] permAuthKeyId = %d, want 10", fakeCli.allPushUpdates[0].PermAuthKeyId)
	}
	if fakeCli.allPushUpdates[1].PermAuthKeyId != 30 {
		t.Errorf("push[1] permAuthKeyId = %d, want 30", fakeCli.allPushUpdates[1].PermAuthKeyId)
	}
}

// --- updatesMe tests ---

func TestSyncUpdatesMeReturnsVoidWithNilContext(t *testing.T) {
	c := New(context.Background(), nil)

	result, err := c.SyncUpdatesMe(&sync.TLSyncUpdatesMe{
		UserId:        1,
		PermAuthKeyId: 2,
		Updates:       makeSyncPlaceholderUpdates(),
	})
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
	if result == nil {
		t.Fatal("expected void result, got nil")
	}
}

func TestSyncUpdatesMeForwardsSessionScopedUpdatesWhenSessionProvided(t *testing.T) {
	fakeCli := &fakeSessionClient{}
	c := New(context.Background(), &svc.ServiceContext{SessionClient: fakeCli})

	updates := makeSyncPlaceholderUpdates()
	authKeyID := int64(3)
	sessionID := int64(4)
	result, err := c.SyncUpdatesMe(&sync.TLSyncUpdatesMe{
		PermAuthKeyId: 2,
		AuthKeyId:     &authKeyID,
		SessionId:     &sessionID,
		Updates:       updates,
	})
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
	if result == nil {
		t.Fatal("expected void result, got nil")
	}
	if fakeCli.pushSessionUpdates == nil {
		t.Fatal("expected session-scoped updates to be forwarded")
	}
	if fakeCli.pushSessionUpdates.PermAuthKeyId != 2 || fakeCli.pushSessionUpdates.AuthKeyId != 3 || fakeCli.pushSessionUpdates.SessionId != 4 {
		t.Fatalf("unexpected session update mapping: %+v", fakeCli.pushSessionUpdates)
	}
	if fakeCli.pushSessionUpdates.Updates != updates {
		t.Fatalf("expected updates object to be forwarded unchanged")
	}
}

func TestSyncUpdatesMeFallsBackToPermAuthPushWhenSessionMissing(t *testing.T) {
	fakeCli := &fakeSessionClient{}
	c := New(context.Background(), &svc.ServiceContext{SessionClient: fakeCli})

	updates := makeSyncPlaceholderUpdates()
	result, err := c.SyncUpdatesMe(&sync.TLSyncUpdatesMe{
		PermAuthKeyId: 2,
		Updates:       updates,
	})
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
	if result == nil {
		t.Fatal("expected void result, got nil")
	}
	if fakeCli.pushUpdatesReq == nil {
		t.Fatal("expected perm-auth updates to be forwarded")
	}
	if fakeCli.pushUpdatesReq.PermAuthKeyId != 2 {
		t.Fatalf("unexpected perm auth key mapping: %+v", fakeCli.pushUpdatesReq)
	}
	if fakeCli.pushUpdatesReq.Updates != updates {
		t.Fatalf("expected updates object to be forwarded unchanged")
	}
}

// --- updatesNotMe tests ---

func TestSyncUpdatesNotMeReturnsVoidWithNilContext(t *testing.T) {
	c := New(context.Background(), nil)

	result, err := c.SyncUpdatesNotMe(&sync.TLSyncUpdatesNotMe{
		UserId:        1,
		PermAuthKeyId: 2,
		Updates:       makeSyncPlaceholderUpdates(),
	})
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
	if result == nil {
		t.Fatal("expected void result, got nil")
	}
}

func TestSyncUpdatesNotMeFansOutExcludingSender(t *testing.T) {
	fakeCli := &fakeSessionClient{}
	fakeStatus := &fakeStatusClient{
		sessions: []*status.TLSessionEntry{
			makeSessionEntry(100, 10, "gw-1"),
			makeSessionEntry(200, 20, "gw-2"),
			makeSessionEntry(300, 30, "gw-1"),
		},
	}
	c := New(context.Background(), &svc.ServiceContext{
		SessionClient: fakeCli,
		StatusClient:  fakeStatus,
	})

	updates := makeSyncPlaceholderUpdates()
	result, err := c.SyncUpdatesNotMe(&sync.TLSyncUpdatesNotMe{
		UserId:        1,
		PermAuthKeyId: 200, // sender's authKeyId — session with AuthKeyId=200 should be skipped
		Updates:       updates,
	})
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
	if result == nil {
		t.Fatal("expected void result, got nil")
	}
	if len(fakeCli.allPushUpdates) != 2 {
		t.Fatalf("expected 2 push calls (excluded sender), got %d", len(fakeCli.allPushUpdates))
	}
	if fakeCli.allPushUpdates[0].PermAuthKeyId != 10 {
		t.Errorf("push[0] permAuthKeyId = %d, want 10", fakeCli.allPushUpdates[0].PermAuthKeyId)
	}
	if fakeCli.allPushUpdates[1].PermAuthKeyId != 30 {
		t.Errorf("push[1] permAuthKeyId = %d, want 30", fakeCli.allPushUpdates[1].PermAuthKeyId)
	}
}

func TestSyncUpdatesNotMeFallsBackWithoutStatusClient(t *testing.T) {
	fakeCli := &fakeSessionClient{}
	c := New(context.Background(), &svc.ServiceContext{SessionClient: fakeCli})

	updates := makeSyncPlaceholderUpdates()
	result, err := c.SyncUpdatesNotMe(&sync.TLSyncUpdatesNotMe{
		UserId:        1,
		PermAuthKeyId: 2,
		Updates:       updates,
	})
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
	if result == nil {
		t.Fatal("expected void result, got nil")
	}
	if fakeCli.pushUpdatesReq == nil {
		t.Fatal("expected updatesNotMe to fallback to direct push")
	}
	if fakeCli.pushUpdatesReq.PermAuthKeyId != 2 {
		t.Fatalf("unexpected perm auth key mapping: %+v", fakeCli.pushUpdatesReq)
	}
}

// --- pushRpcResult tests ---

func TestSyncPushRpcResultReturnsVoidWithNilContext(t *testing.T) {
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

func TestSyncPushRpcResultForwardsToSessionClientWhenAvailable(t *testing.T) {
	fakeCli := &fakeSessionClient{}
	c := New(context.Background(), &svc.ServiceContext{SessionClient: fakeCli})

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
	if fakeCli.pushReq == nil {
		t.Fatal("expected rpc result to be forwarded to session client")
	}
	if fakeCli.pushReq.PermAuthKeyId != 3 || fakeCli.pushReq.AuthKeyId != 2 {
		t.Fatalf("unexpected auth key mapping: %+v", fakeCli.pushReq)
	}
	if fakeCli.pushReq.SessionId != 4 || fakeCli.pushReq.ClientReqMsgId != 5 {
		t.Fatalf("unexpected session mapping: %+v", fakeCli.pushReq)
	}
	if string(fakeCli.pushReq.RpcResultData) != string([]byte{1, 2, 3}) {
		t.Fatalf("unexpected rpc result payload: %v", fakeCli.pushReq.RpcResultData)
	}
}

// --- broadcastUpdates tests ---

func TestSyncBroadcastUpdatesReturnsVoidPlaceholder(t *testing.T) {
	c := New(context.Background(), nil)

	result, err := c.SyncBroadcastUpdates(&sync.TLSyncBroadcastUpdates{
		BroadcastType: 1,
		ChatId:        2,
		ExcludeIdList: []int64{3, 4},
		Updates:       makeSyncPlaceholderUpdates(),
	})
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
	if result == nil {
		t.Fatal("expected void result, got nil")
	}
}

// --- pushBotUpdates tests ---

func TestSyncPushBotUpdatesReturnsVoidPlaceholder(t *testing.T) {
	c := New(context.Background(), nil)

	result, err := c.SyncPushBotUpdates(&sync.TLSyncPushBotUpdates{
		UserId:  1,
		Updates: makeSyncPlaceholderUpdates(),
	})
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
	if result == nil {
		t.Fatal("expected void result, got nil")
	}
}
