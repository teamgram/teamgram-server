package core

import (
	"context"
	"testing"

	sessionclient "github.com/teamgram/teamgram-server/v2/app/interface/session/client"
	"github.com/teamgram/teamgram-server/v2/app/interface/session/session"
	"github.com/teamgram/teamgram-server/v2/app/messenger/sync/internal/svc"
	"github.com/teamgram/teamgram-server/v2/app/messenger/sync/sync"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

type fakeSessionClient struct {
	pushReq            *session.TLSessionPushRpcResultData
	pushUpdatesReq     *session.TLSessionPushUpdatesData
	pushSessionUpdates *session.TLSessionPushSessionUpdatesData
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

func makeSyncPlaceholderUpdates() *tg.TLUpdates {
	return tg.MakeTLUpdates(&tg.TLUpdates{
		Updates: []tg.UpdateClazz{},
		Users:   []tg.UserClazz{},
		Chats:   []tg.ChatClazz{},
		Date:    1,
		Seq:     1,
	})
}

func TestSyncPushUpdatesReturnsVoidPlaceholder(t *testing.T) {
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

func TestSyncPushUpdatesIfNotReturnsVoidPlaceholder(t *testing.T) {
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

func TestSyncUpdatesMeReturnsVoidPlaceholder(t *testing.T) {
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

func TestSyncUpdatesNotMeReturnsVoidPlaceholder(t *testing.T) {
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

func TestSyncUpdatesNotMeForwardsToPermAuthPush(t *testing.T) {
	fakeCli := &fakeSessionClient{}
	c := New(context.Background(), &svc.ServiceContext{SessionClient: fakeCli})

	updates := makeSyncPlaceholderUpdates()
	result, err := c.SyncUpdatesNotMe(&sync.TLSyncUpdatesNotMe{
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
		t.Fatal("expected updatesNotMe to be forwarded to session client")
	}
	if fakeCli.pushUpdatesReq.PermAuthKeyId != 2 {
		t.Fatalf("unexpected perm auth key mapping: %+v", fakeCli.pushUpdatesReq)
	}
	if fakeCli.pushUpdatesReq.Updates != updates {
		t.Fatalf("expected updates object to be forwarded unchanged")
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
