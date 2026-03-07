package core

import (
	"context"
	"errors"
	"testing"

	"github.com/teamgram/marmota/pkg/stores/kv"
	"github.com/teamgram/teamgram-server/app/service/status/internal/config"
	"github.com/teamgram/teamgram-server/app/service/status/internal/dao"
	"github.com/teamgram/teamgram-server/app/service/status/internal/svc"
	"github.com/teamgram/teamgram-server/app/service/status/status"
)

// mockKVStore embeds kv.Store interface; override only the methods we need.
// Unimplemented methods will panic if called (signals a test gap).
type mockKVStore struct {
	kv.Store
	evalCtxFn    func(ctx context.Context, script, key string, args ...any) (any, error)
	hdelCtxFn    func(ctx context.Context, key, field string) (bool, error)
	hgetallCtxFn func(ctx context.Context, key string) (map[string]string, error)
}

func (m *mockKVStore) EvalCtx(ctx context.Context, script, key string, args ...any) (any, error) {
	return m.evalCtxFn(ctx, script, key, args...)
}

func (m *mockKVStore) HdelCtx(ctx context.Context, key, field string) (bool, error) {
	return m.hdelCtxFn(ctx, key, field)
}

func (m *mockKVStore) HgetallCtx(ctx context.Context, key string) (map[string]string, error) {
	return m.hgetallCtxFn(ctx, key)
}

func newTestCore(store kv.Store, expire int) *StatusCore {
	svcCtx := &svc.ServiceContext{
		Config: config.Config{StatusExpire: expire},
		Dao:    &dao.Dao{KV: store},
	}
	return New(context.Background(), svcCtx)
}

// ---- setSessionOnline tests ----

func TestStatusSetSessionOnline_Success(t *testing.T) {
	store := &mockKVStore{
		evalCtxFn: func(ctx context.Context, script, key string, args ...any) (any, error) {
			return int64(1), nil
		},
	}
	c := newTestCore(store, 300)

	in := &status.TLStatusSetSessionOnline{
		UserId: 100,
		Session: &status.SessionEntry{
			AuthKeyId: 12345,
			Gateway:   "gw1",
		},
	}

	r, err := c.StatusSetSessionOnline(in)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if r == nil {
		t.Fatal("expected non-nil result")
	}
}

func TestStatusSetSessionOnline_InvalidParams(t *testing.T) {
	store := &mockKVStore{}
	c := newTestCore(store, 300)

	tests := []struct {
		name string
		in   *status.TLStatusSetSessionOnline
	}{
		{"zero userId", &status.TLStatusSetSessionOnline{UserId: 0, Session: &status.SessionEntry{AuthKeyId: 1}}},
		{"negative userId", &status.TLStatusSetSessionOnline{UserId: -1, Session: &status.SessionEntry{AuthKeyId: 1}}},
		{"nil session", &status.TLStatusSetSessionOnline{UserId: 1, Session: nil}},
		{"zero authKeyId", &status.TLStatusSetSessionOnline{UserId: 1, Session: &status.SessionEntry{AuthKeyId: 0}}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := c.StatusSetSessionOnline(tt.in)
			if err == nil {
				t.Error("expected error for invalid params")
			}
		})
	}
}

func TestStatusSetSessionOnline_EvalError(t *testing.T) {
	evalErr := errors.New("redis eval failed")
	store := &mockKVStore{
		evalCtxFn: func(ctx context.Context, script, key string, args ...any) (any, error) {
			return nil, evalErr
		},
	}
	c := newTestCore(store, 300)

	in := &status.TLStatusSetSessionOnline{
		UserId:  100,
		Session: &status.SessionEntry{AuthKeyId: 12345},
	}

	_, err := c.StatusSetSessionOnline(in)
	if err == nil {
		t.Fatal("expected error")
	}
	if !errors.Is(err, evalErr) {
		t.Errorf("expected wrapped evalErr, got: %v", err)
	}
}

// ---- setSessionOffline tests ----

func TestStatusSetSessionOffline_Success(t *testing.T) {
	store := &mockKVStore{
		hdelCtxFn: func(ctx context.Context, key, field string) (bool, error) {
			return true, nil
		},
	}
	c := newTestCore(store, 300)

	in := &status.TLStatusSetSessionOffline{UserId: 100, AuthKeyId: 12345}

	r, err := c.StatusSetSessionOffline(in)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if r == nil {
		t.Fatal("expected non-nil result")
	}
}

func TestStatusSetSessionOffline_InvalidParams(t *testing.T) {
	store := &mockKVStore{}
	c := newTestCore(store, 300)

	tests := []struct {
		name string
		in   *status.TLStatusSetSessionOffline
	}{
		{"zero userId", &status.TLStatusSetSessionOffline{UserId: 0, AuthKeyId: 1}},
		{"zero authKeyId", &status.TLStatusSetSessionOffline{UserId: 1, AuthKeyId: 0}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := c.StatusSetSessionOffline(tt.in)
			if err == nil {
				t.Error("expected error for invalid params")
			}
		})
	}
}

func TestStatusSetSessionOffline_HdelError(t *testing.T) {
	hdelErr := errors.New("redis hdel failed")
	store := &mockKVStore{
		hdelCtxFn: func(ctx context.Context, key, field string) (bool, error) {
			return false, hdelErr
		},
	}
	c := newTestCore(store, 300)

	in := &status.TLStatusSetSessionOffline{UserId: 100, AuthKeyId: 12345}

	_, err := c.StatusSetSessionOffline(in)
	if err == nil {
		t.Fatal("expected error")
	}
	if !errors.Is(err, hdelErr) {
		t.Errorf("expected wrapped hdelErr, got: %v", err)
	}
}

// ---- getUserOnlineSessions tests ----

func TestStatusGetUserOnlineSessions_Success(t *testing.T) {
	store := &mockKVStore{
		hgetallCtxFn: func(ctx context.Context, key string) (map[string]string, error) {
			return map[string]string{
				"111": `{"auth_key_id":111,"gateway":"gw1"}`,
				"222": `{"auth_key_id":222,"gateway":"gw2"}`,
			}, nil
		},
	}
	c := newTestCore(store, 300)

	in := &status.TLStatusGetUserOnlineSessions{UserId: 100}

	r, err := c.StatusGetUserOnlineSessions(in)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(r.UserSessions) != 2 {
		t.Errorf("expected 2 sessions, got %d", len(r.UserSessions))
	}
}

func TestStatusGetUserOnlineSessions_Empty(t *testing.T) {
	store := &mockKVStore{
		hgetallCtxFn: func(ctx context.Context, key string) (map[string]string, error) {
			return map[string]string{}, nil
		},
	}
	c := newTestCore(store, 300)

	in := &status.TLStatusGetUserOnlineSessions{UserId: 100}

	r, err := c.StatusGetUserOnlineSessions(in)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(r.UserSessions) != 0 {
		t.Errorf("expected 0 sessions, got %d", len(r.UserSessions))
	}
}

func TestStatusGetUserOnlineSessions_SkipsBadJSON(t *testing.T) {
	store := &mockKVStore{
		hgetallCtxFn: func(ctx context.Context, key string) (map[string]string, error) {
			return map[string]string{
				"111": `{"auth_key_id":111}`,
				"222": `bad json`,
			}, nil
		},
	}
	c := newTestCore(store, 300)

	in := &status.TLStatusGetUserOnlineSessions{UserId: 100}

	r, err := c.StatusGetUserOnlineSessions(in)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(r.UserSessions) != 1 {
		t.Errorf("expected 1 valid session (bad json skipped), got %d", len(r.UserSessions))
	}
}

func TestStatusGetUserOnlineSessions_HgetallError(t *testing.T) {
	store := &mockKVStore{
		hgetallCtxFn: func(ctx context.Context, key string) (map[string]string, error) {
			return nil, errors.New("redis error")
		},
	}
	c := newTestCore(store, 300)

	in := &status.TLStatusGetUserOnlineSessions{UserId: 100}

	_, err := c.StatusGetUserOnlineSessions(in)
	if err == nil {
		t.Fatal("expected error")
	}
}

// ---- getUsersOnlineSessionsList tests ----

func TestStatusGetUsersOnlineSessionsList_TooManyUsers(t *testing.T) {
	store := &mockKVStore{}
	c := newTestCore(store, 300)

	users := make([]int64, maxBatchUsers+1)
	for i := range users {
		users[i] = int64(i + 1)
	}

	in := &status.TLStatusGetUsersOnlineSessionsList{Users: users}

	_, err := c.StatusGetUsersOnlineSessionsList(in)
	if err == nil {
		t.Fatal("expected error for too many users")
	}
}
