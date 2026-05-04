package core

import (
	"context"
	"errors"
	"testing"

	"github.com/teamgram/teamgram-server/v2/app/service/presence/internal/config"
	"github.com/teamgram/teamgram-server/v2/app/service/presence/internal/repository"
	"github.com/teamgram/teamgram-server/v2/app/service/presence/internal/svc"
	presencepb "github.com/teamgram/teamgram-server/v2/app/service/presence/presence"
	"github.com/teamgram/teamgram-server/v2/pkg/net/kitex/identity"
)

type testRepo struct {
	onlineSession     *presencepb.OnlineSession
	onlineNow         int64
	onlineExpire      int
	onlineHashTTL     int
	onlineCleanup     int
	getUsersWasCalled bool
}

func (r *testRepo) SetSessionOnline(ctx context.Context, session *presencepb.OnlineSession, now int64, expireSeconds int, hashTTLSeconds int, cleanupIntervalSeconds int) error {
	r.onlineSession = session
	r.onlineNow = now
	r.onlineExpire = expireSeconds
	r.onlineHashTTL = hashTTLSeconds
	r.onlineCleanup = cleanupIntervalSeconds
	return nil
}

func (r *testRepo) SetSessionOffline(ctx context.Context, userID, authKeyID, sessionID int64) error {
	return nil
}

func (r *testRepo) GetUserOnlineSessions(ctx context.Context, userID int64, now int64) (*presencepb.TLUserOnlineSessions, error) {
	return presencepb.MakeTLUserOnlineSessions(&presencepb.TLUserOnlineSessions{UserId: userID, Sessions: []*presencepb.TLOnlineSession{}}), nil
}

func (r *testRepo) GetUsersOnlineSessions(ctx context.Context, userIDs []int64, now int64) (*presencepb.VectorUserOnlineSessions, error) {
	r.getUsersWasCalled = true
	return &presencepb.VectorUserOnlineSessions{}, nil
}

func (r *testRepo) CleanupExpiredForUser(ctx context.Context, userID int64, now int64) error {
	return nil
}

func newTestPresenceCore(repo svc.Repository, cfg config.Config) *PresenceCore {
	return New(context.Background(), &svc.ServiceContext{
		Config:  cfg,
		Repo:    repo,
		Limiter: svc.NewCallerLimiter(),
	})
}

func validSetSessionOnlineRequest() *presencepb.TLPresenceSetSessionOnline {
	return &presencepb.TLPresenceSetSessionOnline{
		Session: presencepb.MakeTLOnlineSession(&presencepb.TLOnlineSession{
			UserId:            42,
			PermAuthKeyId:     9001,
			AuthKeyId:         1001,
			AuthKeyType:       1,
			SessionId:         2002,
			GatewayId:         "gw1",
			GatewayGeneration: "gen1",
			GatewayRpcAddr:    "127.0.0.1:20110",
			Layer:             224,
			Client:            "tdesktop",
			UpdatedAt:         1,
			ExpiresAt:         2,
		}),
	}
}

func TestPresenceSetSessionOnlineRejectsInvalidSession(t *testing.T) {
	c := newTestPresenceCore(&testRepo{}, config.Config{})
	_, err := c.PresenceSetSessionOnline(&presencepb.TLPresenceSetSessionOnline{})
	if !errors.Is(err, presencepb.ErrPresenceInvalidArgument) {
		t.Fatalf("error = %v, want ErrPresenceInvalidArgument", err)
	}
}

func TestPresenceSetSessionOnlineRequiresGatewayCallerWhenEnabled(t *testing.T) {
	cfg := config.Config{RequireCallerIdentity: true, GatewayCallers: []string{"gateway"}}
	c := newTestPresenceCore(&testRepo{}, cfg)
	c.ctx = identity.WithCallerService(context.Background(), "sync")
	_, err := c.PresenceSetSessionOnline(validSetSessionOnlineRequest())
	if !errors.Is(err, presencepb.ErrPresencePermissionDenied) {
		t.Fatalf("error = %v, want ErrPresencePermissionDenied", err)
	}
}

func TestPermissionDeniedCallerLabelIsBounded(t *testing.T) {
	if got := permissionDeniedCallerLabel(""); got != "missing" {
		t.Fatalf("permissionDeniedCallerLabel(\"\") = %q, want missing", got)
	}
	if got := permissionDeniedCallerLabel("arbitrary-runtime-caller"); got != "unauthorized" {
		t.Fatalf("permissionDeniedCallerLabel(arbitrary) = %q, want unauthorized", got)
	}
}

func TestPresenceGetUsersOnlineSessionsRejectsOversizedBatch(t *testing.T) {
	users := make([]int64, repository.MaxBatchUsers+1)
	for i := range users {
		users[i] = int64(i + 1)
	}
	c := newTestPresenceCore(&testRepo{}, config.Config{})
	_, err := c.PresenceGetUsersOnlineSessions(&presencepb.TLPresenceGetUsersOnlineSessions{Users: users})
	if !errors.Is(err, presencepb.ErrPresenceInvalidArgument) {
		t.Fatalf("error = %v, want ErrPresenceInvalidArgument", err)
	}
}

func TestPresenceSetSessionOnlinePassesServiceTimeAndTTLs(t *testing.T) {
	repo := &testRepo{}
	cfg := config.Config{SessionExpiresSeconds: 60, HashKeyTTLSeconds: 600, CleanupOnWriteIntervalSeconds: 30}
	c := newTestPresenceCore(repo, cfg)
	_, err := c.PresenceSetSessionOnline(validSetSessionOnlineRequest())
	if err != nil {
		t.Fatalf("PresenceSetSessionOnline() error = %v", err)
	}
	if repo.onlineSession == nil {
		t.Fatal("SetSessionOnline was not called")
	}
	if repo.onlineNow <= 0 {
		t.Fatalf("now = %d, want service unix time", repo.onlineNow)
	}
	if repo.onlineExpire != 60 {
		t.Fatalf("expireSeconds = %d, want 60", repo.onlineExpire)
	}
	if repo.onlineHashTTL != 600 {
		t.Fatalf("hashTTLSeconds = %d, want 600", repo.onlineHashTTL)
	}
	if repo.onlineCleanup != 30 {
		t.Fatalf("cleanupIntervalSeconds = %d, want 30", repo.onlineCleanup)
	}
	if repo.onlineSession.UpdatedAt != 1 || repo.onlineSession.ExpiresAt != 2 {
		t.Fatalf("handler should pass original session to repository, got updated_at=%d expires_at=%d", repo.onlineSession.UpdatedAt, repo.onlineSession.ExpiresAt)
	}
}

func TestPresenceSetSessionOnlineAllowsSignedAuthKeyID(t *testing.T) {
	repo := &testRepo{}
	c := newTestPresenceCore(repo, config.Config{SessionExpiresSeconds: 60, HashKeyTTLSeconds: 600})
	req := validSetSessionOnlineRequest()
	req.Session.PermAuthKeyId = -3922385800037876977
	req.Session.AuthKeyId = -3213093451295049619

	_, err := c.PresenceSetSessionOnline(req)
	if err != nil {
		t.Fatalf("PresenceSetSessionOnline() error = %v", err)
	}
	if repo.onlineSession == nil || repo.onlineSession.AuthKeyId != -3213093451295049619 {
		t.Fatalf("online auth_key_id = %v, want signed auth key", repo.onlineSession)
	}
	if repo.onlineSession.PermAuthKeyId != -3922385800037876977 {
		t.Fatalf("online perm_auth_key_id = %d, want signed perm auth key", repo.onlineSession.PermAuthKeyId)
	}
}

func TestPresenceGetUserOnlineSessionsRejectsQueryQuotaExceeded(t *testing.T) {
	cfg := config.Config{
		RequireCallerIdentity:            true,
		SyncCallers:                      []string{"sync"},
		PresenceQueryDefaultQPSPerCaller: 1,
	}
	c := newTestPresenceCore(&testRepo{}, cfg)
	c.ctx = identity.WithCallerService(context.Background(), "sync")
	if _, err := c.PresenceGetUserOnlineSessions(&presencepb.TLPresenceGetUserOnlineSessions{UserId: 42}); err != nil {
		t.Fatalf("first PresenceGetUserOnlineSessions() error = %v", err)
	}
	_, err := c.PresenceGetUserOnlineSessions(&presencepb.TLPresenceGetUserOnlineSessions{UserId: 42})
	if !errors.Is(err, presencepb.ErrPresenceQuotaExceeded) {
		t.Fatalf("second error = %v, want ErrPresenceQuotaExceeded", err)
	}
}

func TestPresenceGetGatewaySessionsRejectsDiagnosticsQuotaExceeded(t *testing.T) {
	cfg := config.Config{
		RequireCallerIdentity:                  true,
		AdminCallers:                           []string{"admin"},
		PresenceGatewayDiagnosticsQPSPerCaller: 1,
	}
	c := newTestPresenceCore(&testRepo{}, cfg)
	c.ctx = identity.WithCallerService(context.Background(), "admin")
	if _, err := c.PresenceGetGatewaySessions(&presencepb.TLPresenceGetGatewaySessions{GatewayId: "gw1"}); err != nil {
		t.Fatalf("first PresenceGetGatewaySessions() error = %v", err)
	}
	_, err := c.PresenceGetGatewaySessions(&presencepb.TLPresenceGetGatewaySessions{GatewayId: "gw1"})
	if !errors.Is(err, presencepb.ErrPresenceQuotaExceeded) {
		t.Fatalf("second error = %v, want ErrPresenceQuotaExceeded", err)
	}
}
