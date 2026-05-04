package presence

import (
	"context"
	"errors"
	"sync"
	"testing"
	"time"

	presencepb "github.com/teamgram/teamgram-server/v2/app/service/presence/presence"
)

type fakePresenceClient struct {
	mu              sync.Mutex
	onlineErr       error
	onlineCh        chan *presencepb.OnlineSession
	onlineStartedCh chan struct{}
	blockCh         chan struct{}
	online          []*presencepb.OnlineSession
	offline         []offlineCall
	events          []string
}

type offlineCall struct {
	userID    int64
	authKeyID int64
	sessionID int64
}

func (f *fakePresenceClient) SetSessionOnline(ctx context.Context, session *presencepb.OnlineSession) error {
	if f.onlineStartedCh != nil {
		select {
		case f.onlineStartedCh <- struct{}{}:
		default:
		}
	}
	if f.blockCh != nil {
		<-f.blockCh
	}
	f.mu.Lock()
	f.online = append(f.online, session)
	f.events = append(f.events, "online")
	err := f.onlineErr
	f.mu.Unlock()
	if f.onlineCh != nil {
		select {
		case f.onlineCh <- session:
		default:
		}
	}
	return err
}

func (f *fakePresenceClient) SetSessionOffline(ctx context.Context, userID, authKeyID, sessionID int64) error {
	f.mu.Lock()
	defer f.mu.Unlock()
	f.offline = append(f.offline, offlineCall{userID: userID, authKeyID: authKeyID, sessionID: sessionID})
	f.events = append(f.events, "offline")
	return nil
}

func (f *fakePresenceClient) onlineCount() int {
	f.mu.Lock()
	defer f.mu.Unlock()
	return len(f.online)
}

func (f *fakePresenceClient) offlineCount() int {
	f.mu.Lock()
	defer f.mu.Unlock()
	return len(f.offline)
}

func (f *fakePresenceClient) setOnlineErr(err error) {
	f.mu.Lock()
	defer f.mu.Unlock()
	f.onlineErr = err
}

func (f *fakePresenceClient) eventSnapshot() []string {
	f.mu.Lock()
	defer f.mu.Unlock()
	return append([]string(nil), f.events...)
}

func (f *fakePresenceClient) waitOnline(t *testing.T, want int) {
	t.Helper()
	deadline := time.After(time.Second)
	for {
		if f.onlineCount() >= want {
			return
		}
		select {
		case <-deadline:
			t.Fatalf("online calls = %d, want at least %d", f.onlineCount(), want)
		case <-time.After(time.Millisecond):
		}
	}
}

func TestRegistrarRefreshesIdleSessionBeforeTTL(t *testing.T) {
	now := time.Unix(100, 0)
	client := &fakePresenceClient{}
	registrar := NewRegistrar(testConfig(), client, func() time.Time { return now })
	session := testSession()

	registrar.Register(context.Background(), session)
	client.waitOnline(t, 1)
	now = now.Add(21 * time.Second)
	registrar.RefreshDue(context.Background())

	if client.onlineCount() != 2 {
		t.Fatalf("online calls = %d, want 2", client.onlineCount())
	}
}

func TestRegistrarRouteChangeBypassesThrottle(t *testing.T) {
	now := time.Unix(100, 0)
	client := &fakePresenceClient{}
	registrar := NewRegistrar(testConfig(), client, func() time.Time { return now })
	session := testSession()

	registrar.Register(context.Background(), session)
	client.waitOnline(t, 1)
	session.Layer = 225
	registrar.Register(context.Background(), session)
	client.waitOnline(t, 2)

	if client.onlineCount() != 2 {
		t.Fatalf("online calls = %d, want 2", client.onlineCount())
	}
	client.mu.Lock()
	if got := client.online[1].Layer; got != 225 {
		client.mu.Unlock()
		t.Fatalf("second online layer = %d, want 225", got)
	}
	client.mu.Unlock()
}

func TestRegistrarShutdownStopsAtMaxSessions(t *testing.T) {
	now := time.Unix(100, 0)
	client := &fakePresenceClient{}
	cfg := testConfig()
	cfg.ShutdownOfflineMaxSessions = 2
	registrar := NewRegistrar(cfg, client, func() time.Time { return now })

	for i := int64(1); i <= 3; i++ {
		session := testSession()
		session.AuthKeyID = 1000 + i
		session.SessionID = 2000 + i
		registrar.Register(context.Background(), session)
	}
	registrar.OfflineAll(context.Background())

	if client.offlineCount() != 2 {
		t.Fatalf("offline calls = %d, want 2", client.offlineCount())
	}
	if len(registrar.sessions) != 1 {
		t.Fatalf("remaining sessions = %d, want 1", len(registrar.sessions))
	}
}

func TestRegistrarFailedOnlineRetriesAfterRetryInterval(t *testing.T) {
	now := time.Unix(100, 0)
	client := &fakePresenceClient{onlineErr: errors.New("presence unavailable")}
	registrar := NewRegistrar(testConfig(), client, func() time.Time { return now })
	session := testSession()

	registrar.Register(context.Background(), session)
	registrar.Register(context.Background(), session)
	now = now.Add(5 * time.Second)
	registrar.Register(context.Background(), session)
	client.waitOnline(t, 2)

	if client.onlineCount() != 2 {
		t.Fatalf("online calls = %d, want 2", client.onlineCount())
	}
}

func TestRegistrarRegisterDoesNotWaitForSlowOnline(t *testing.T) {
	now := time.Unix(100, 0)
	client := &fakePresenceClient{blockCh: make(chan struct{})}
	registrar := NewRegistrar(testConfig(), client, func() time.Time { return now })
	done := make(chan struct{})

	go func() {
		registrar.Register(context.Background(), testSession())
		close(done)
	}()

	select {
	case <-done:
	case <-time.After(50 * time.Millisecond):
		close(client.blockCh)
		t.Fatal("Register blocked on SetSessionOnline")
	}
	close(client.blockCh)
	client.waitOnline(t, 1)
}

func TestRegistrarUnregisterFencesInflightAsyncOnline(t *testing.T) {
	now := time.Unix(100, 0)
	client := &fakePresenceClient{
		onlineStartedCh: make(chan struct{}, 1),
		blockCh:         make(chan struct{}),
	}
	registrar := NewRegistrar(testConfig(), client, func() time.Time { return now })
	session := testSession()

	registrar.Register(context.Background(), session)
	select {
	case <-client.onlineStartedCh:
	case <-time.After(time.Second):
		t.Fatal("online attempt did not start")
	}

	unregistered := make(chan struct{})
	go func() {
		registrar.Unregister(context.Background(), session.AuthKeyID, session.SessionID)
		close(unregistered)
	}()

	select {
	case <-unregistered:
		close(client.blockCh)
		client.waitOnline(t, 1)
		t.Fatal("Unregister returned before in-flight online attempt completed")
	case <-time.After(20 * time.Millisecond):
	}

	close(client.blockCh)
	select {
	case <-unregistered:
	case <-time.After(time.Second):
		t.Fatal("Unregister did not finish after online attempt completed")
	}

	events := client.eventSnapshot()
	if len(events) != 2 {
		t.Fatalf("presence events = %v, want [online offline]", events)
	}
	if events[0] != "online" || events[1] != "offline" {
		t.Fatalf("presence events = %v, want online before offline so offline wins", events)
	}
}

func TestRegistrarRefreshRetryWaitsAfterFailedRefreshOfSuccessfulSession(t *testing.T) {
	now := time.Unix(100, 0)
	client := &fakePresenceClient{}
	registrar := NewRegistrar(testConfig(), client, func() time.Time { return now })
	session := testSession()

	registrar.Register(context.Background(), session)
	client.waitOnline(t, 1)
	now = now.Add(21 * time.Second)
	client.setOnlineErr(errors.New("presence unavailable"))
	registrar.RefreshDue(context.Background())
	if client.onlineCount() != 2 {
		t.Fatalf("online calls after first refresh = %d, want 2", client.onlineCount())
	}
	now = now.Add(4 * time.Second)
	registrar.RefreshDue(context.Background())
	if client.onlineCount() != 2 {
		t.Fatalf("online calls before retry interval = %d, want 2", client.onlineCount())
	}
	now = now.Add(time.Second)
	registrar.RefreshDue(context.Background())
	if client.onlineCount() != 3 {
		t.Fatalf("online calls after retry interval = %d, want 3", client.onlineCount())
	}
}

func testConfig() Config {
	return Config{
		GatewayID:                  "gateway-test",
		GatewayGeneration:          "generation-test",
		GatewayRPCAddr:             "127.0.0.1:20110",
		RefreshMinInterval:         20 * time.Second,
		RefreshRetryMinInterval:    5 * time.Second,
		RefreshScanInterval:        10 * time.Second,
		ShutdownOfflineDeadline:    time.Second,
		ShutdownOfflineMaxSessions: 10000,
	}
}

func testSession() ActiveSession {
	return ActiveSession{
		UserID:        123,
		PermAuthKeyID: 456,
		AuthKeyID:     789,
		AuthKeyType:   1,
		SessionID:     1001,
		Layer:         224,
		Client:        "tdesktop",
	}
}
