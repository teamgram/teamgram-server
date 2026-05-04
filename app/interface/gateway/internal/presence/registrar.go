package presence

import (
	"context"
	"sync"
	"time"

	presencepb "github.com/teamgram/teamgram-server/v2/app/service/presence/presence"
	"github.com/teamgram/teamgram-server/v2/pkg/net/kitex/identity"
	"github.com/zeromicro/go-zero/core/logx"
)

type ActiveSession struct {
	UserID        int64
	PermAuthKeyID int64
	AuthKeyID     int64
	AuthKeyType   int32
	SessionID     int64
	Layer         int32
	Client        string
}

type Config struct {
	GatewayID                  string
	GatewayGeneration          string
	GatewayRPCAddr             string
	RefreshMinInterval         time.Duration
	RefreshRetryMinInterval    time.Duration
	RefreshScanInterval        time.Duration
	ShutdownOfflineDeadline    time.Duration
	ShutdownOfflineMaxSessions int
}

type PresenceClient interface {
	SetSessionOnline(ctx context.Context, session *presencepb.OnlineSession) error
	SetSessionOffline(ctx context.Context, userID, authKeyID, sessionID int64) error
}

type Registrar struct {
	cfg    Config
	client PresenceClient
	now    func() time.Time

	mu       sync.Mutex
	sessions map[sessionKey]*sessionState
	// ops serializes presence RPC side effects per session key.
	ops      map[sessionKey]*sessionOps
	started  bool
	stopOnce sync.Once
	stopCh   chan struct{}
	doneCh   chan struct{}
}

type sessionKey struct {
	authKeyID int64
	sessionID int64
}

type sessionState struct {
	session       ActiveSession
	lastAttemptAt time.Time
	lastSuccessAt time.Time
}

type sessionOps struct {
	mu sync.Mutex
}

type onlineAttempt struct {
	key       sessionKey
	session   ActiveSession
	attemptAt time.Time
	ops       *sessionOps
}

func NewRegistrar(cfg Config, client PresenceClient, now func() time.Time) *Registrar {
	if now == nil {
		now = time.Now
	}
	return &Registrar{
		cfg:      cfg,
		client:   client,
		now:      now,
		sessions: make(map[sessionKey]*sessionState),
		ops:      make(map[sessionKey]*sessionOps),
		stopCh:   make(chan struct{}),
		doneCh:   make(chan struct{}),
	}
}

func (r *Registrar) Register(ctx context.Context, session ActiveSession) {
	if !r.canUse(session) {
		return
	}
	key := sessionKey{authKeyID: session.AuthKeyID, sessionID: session.SessionID}
	now := r.now()

	r.mu.Lock()
	state := r.sessions[key]
	if state == nil {
		state = &sessionState{}
		r.sessions[key] = state
	}
	routeChanged := activeRouteChanged(state.session, session)
	state.session = session
	if !state.lastSuccessAt.IsZero() && !routeChanged && now.Sub(state.lastSuccessAt) < r.cfg.RefreshMinInterval {
		r.mu.Unlock()
		return
	}
	if state.lastSuccessAt.IsZero() && !routeChanged && !state.lastAttemptAt.IsZero() && now.Sub(state.lastAttemptAt) < r.cfg.RefreshRetryMinInterval {
		r.mu.Unlock()
		return
	}
	state.lastAttemptAt = now
	attempt := r.lockOnlineAttemptLocked(key, session, now)
	r.mu.Unlock()

	go r.completeOnlineAttempt(context.Background(), attempt, "online")
}

func (r *Registrar) Unregister(ctx context.Context, authKeyID, sessionID int64) {
	if r == nil || r.client == nil || authKeyID == 0 || sessionID == 0 {
		return
	}
	key := sessionKey{authKeyID: authKeyID, sessionID: sessionID}
	r.mu.Lock()
	state := r.sessions[key]
	delete(r.sessions, key)
	ops := r.sessionOpsLocked(key)
	ops.mu.Lock()
	r.mu.Unlock()
	defer ops.mu.Unlock()
	if state == nil || state.session.UserID <= 0 {
		return
	}
	if err := r.setOffline(ctx, state.session); err != nil {
		logx.WithContext(ctx).Errorf("gateway presence offline failed: user_id=%d auth_key_id=%d session_id=%d err=%v", state.session.UserID, authKeyID, sessionID, err)
	}
}

func (r *Registrar) RefreshDue(ctx context.Context) {
	if r == nil || r.client == nil {
		return
	}
	now := r.now()
	due := r.dueSessions(now)
	for _, attempt := range due {
		err := r.setOnline(ctx, attempt.session)
		attempt.ops.mu.Unlock()
		if err != nil {
			logx.WithContext(ctx).Errorf("gateway presence refresh failed: user_id=%d auth_key_id=%d session_id=%d err=%v", attempt.session.UserID, attempt.session.AuthKeyID, attempt.session.SessionID, err)
			continue
		}
		r.mu.Lock()
		if state := r.sessions[attempt.key]; state != nil && state.session == attempt.session && state.lastAttemptAt.Equal(attempt.attemptAt) {
			state.lastSuccessAt = now
		}
		r.mu.Unlock()
	}
}

func (r *Registrar) completeOnlineAttempt(ctx context.Context, attempt onlineAttempt, operation string) {
	err := r.setOnline(ctx, attempt.session)
	attempt.ops.mu.Unlock()
	if err != nil {
		logx.WithContext(ctx).Errorf("gateway presence %s failed: user_id=%d auth_key_id=%d session_id=%d err=%v", operation, attempt.session.UserID, attempt.session.AuthKeyID, attempt.session.SessionID, err)
		return
	}

	r.mu.Lock()
	if current := r.sessions[attempt.key]; current != nil && current.session == attempt.session && current.lastAttemptAt.Equal(attempt.attemptAt) {
		current.lastSuccessAt = attempt.attemptAt
	}
	r.mu.Unlock()
}

func (r *Registrar) Start(ctx context.Context) {
	if r == nil || r.cfg.RefreshScanInterval <= 0 {
		return
	}
	r.mu.Lock()
	if r.started {
		r.mu.Unlock()
		return
	}
	r.started = true
	r.mu.Unlock()
	go r.loop(ctx)
}

func (r *Registrar) Stop(ctx context.Context) error {
	if r == nil {
		return nil
	}
	r.stopOnce.Do(func() {
		close(r.stopCh)
		r.mu.Lock()
		started := r.started
		r.mu.Unlock()
		if started {
			<-r.doneCh
		}
		r.OfflineAll(ctx)
	})
	return nil
}

func (r *Registrar) OfflineAll(ctx context.Context) {
	if r == nil || r.client == nil {
		return
	}
	snapshot := r.offlineSnapshot()
	if len(snapshot) == 0 {
		return
	}
	deadline := r.cfg.ShutdownOfflineDeadline
	if deadline <= 0 {
		deadline = time.Second
	}
	ctx, cancel := context.WithTimeout(ctx, deadline)
	defer cancel()
	for _, attempt := range snapshot {
		if err := ctx.Err(); err != nil {
			attempt.ops.mu.Unlock()
			continue
		}
		if err := r.setOffline(ctx, attempt.session); err != nil {
			logx.WithContext(ctx).Errorf("gateway presence shutdown offline failed: user_id=%d auth_key_id=%d session_id=%d err=%v", attempt.session.UserID, attempt.session.AuthKeyID, attempt.session.SessionID, err)
		}
		attempt.ops.mu.Unlock()
	}
}

func (r *Registrar) canUse(session ActiveSession) bool {
	return r != nil &&
		r.client != nil &&
		r.cfg.GatewayID != "" &&
		r.cfg.GatewayGeneration != "" &&
		r.cfg.GatewayRPCAddr != "" &&
		session.UserID > 0 &&
		session.PermAuthKeyID != 0 &&
		session.AuthKeyID != 0 &&
		session.SessionID != 0
}

func (r *Registrar) setOnline(ctx context.Context, session ActiveSession) error {
	ctx = identity.WithCallerService(ctx, "gateway")
	return r.client.SetSessionOnline(ctx, presencepb.MakeTLOnlineSession(&presencepb.TLOnlineSession{
		UserId:            session.UserID,
		PermAuthKeyId:     session.PermAuthKeyID,
		AuthKeyId:         session.AuthKeyID,
		AuthKeyType:       session.AuthKeyType,
		SessionId:         session.SessionID,
		GatewayId:         r.cfg.GatewayID,
		GatewayGeneration: r.cfg.GatewayGeneration,
		GatewayRpcAddr:    r.cfg.GatewayRPCAddr,
		Layer:             session.Layer,
		Client:            session.Client,
	}))
}

func (r *Registrar) setOffline(ctx context.Context, session ActiveSession) error {
	ctx = identity.WithCallerService(ctx, "gateway")
	return r.client.SetSessionOffline(ctx, session.UserID, session.AuthKeyID, session.SessionID)
}

func (r *Registrar) loop(ctx context.Context) {
	defer close(r.doneCh)
	ticker := time.NewTicker(r.cfg.RefreshScanInterval)
	defer ticker.Stop()
	for {
		select {
		case <-ctx.Done():
			return
		case <-r.stopCh:
			return
		case <-ticker.C:
			r.RefreshDue(ctx)
		}
	}
}

func (r *Registrar) dueSessions(now time.Time) []onlineAttempt {
	r.mu.Lock()
	defer r.mu.Unlock()
	due := make([]onlineAttempt, 0)
	for key, state := range r.sessions {
		if state.lastSuccessAt.IsZero() {
			if state.lastAttemptAt.IsZero() || now.Sub(state.lastAttemptAt) >= r.cfg.RefreshRetryMinInterval {
				state.lastAttemptAt = now
				due = append(due, r.lockOnlineAttemptLocked(key, state.session, now))
			}
			continue
		}
		if now.Sub(state.lastSuccessAt) >= r.cfg.RefreshMinInterval {
			if state.lastAttemptAt.After(state.lastSuccessAt) && now.Sub(state.lastAttemptAt) < r.cfg.RefreshRetryMinInterval {
				continue
			}
			state.lastAttemptAt = now
			due = append(due, r.lockOnlineAttemptLocked(key, state.session, now))
		}
	}
	return due
}

func (r *Registrar) offlineSnapshot() []onlineAttempt {
	r.mu.Lock()
	defer r.mu.Unlock()
	maxSessions := r.cfg.ShutdownOfflineMaxSessions
	if maxSessions <= 0 || maxSessions > len(r.sessions) {
		maxSessions = len(r.sessions)
	}
	snapshot := make([]onlineAttempt, 0, maxSessions)
	for key, state := range r.sessions {
		if len(snapshot) >= maxSessions {
			break
		}
		ops := r.sessionOpsLocked(key)
		ops.mu.Lock()
		delete(r.sessions, key)
		snapshot = append(snapshot, onlineAttempt{
			key:     key,
			session: state.session,
			ops:     ops,
		})
	}
	return snapshot
}

func (r *Registrar) sessionOpsLocked(key sessionKey) *sessionOps {
	ops := r.ops[key]
	if ops == nil {
		ops = &sessionOps{}
		r.ops[key] = ops
	}
	return ops
}

func (r *Registrar) lockOnlineAttemptLocked(key sessionKey, session ActiveSession, attemptAt time.Time) onlineAttempt {
	ops := r.sessionOpsLocked(key)
	ops.mu.Lock()
	return onlineAttempt{
		key:       key,
		session:   session,
		attemptAt: attemptAt,
		ops:       ops,
	}
}

func activeRouteChanged(old, next ActiveSession) bool {
	return old.UserID != next.UserID ||
		old.PermAuthKeyID != next.PermAuthKeyID ||
		old.AuthKeyID != next.AuthKeyID ||
		old.AuthKeyType != next.AuthKeyType ||
		old.SessionID != next.SessionID ||
		old.Layer != next.Layer ||
		old.Client != next.Client
}
