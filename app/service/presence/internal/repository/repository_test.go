package repository

import (
	"context"
	"errors"
	"testing"

	presencepb "github.com/teamgram/teamgram-server/v2/app/service/presence/presence"
)

func TestSessionFieldUsesAuthKeyAndSession(t *testing.T) {
	got := sessionField(1001, 2002)
	if got != "1001:2002" {
		t.Fatalf("sessionField() = %q, want 1001:2002", got)
	}
}

func TestBuildOnlineSessionsFiltersExpiredAndCorruptEntries(t *testing.T) {
	raw := map[string]string{
		"1001:2002": `{"user_id":42,"perm_auth_key_id":9001,"auth_key_id":1001,"auth_key_type":1,"session_id":2002,"gateway_id":"gw1","gateway_generation":"gen1","gateway_rpc_addr":"127.0.0.1:20110","layer":224,"client":"tdesktop","updated_at":100,"expires_at":160}`,
		"1002:2003": `{"user_id":42,"perm_auth_key_id":9002,"auth_key_id":1002,"auth_key_type":1,"session_id":2003,"gateway_id":"gw1","gateway_generation":"gen1","gateway_rpc_addr":"127.0.0.1:20110","layer":224,"client":"tdesktop","updated_at":100,"expires_at":120}`,
		"bad":       `{`,
		"mismatch":  `{"user_id":43,"perm_auth_key_id":9003,"auth_key_id":1003,"auth_key_type":1,"session_id":2004,"gateway_id":"gw1","gateway_generation":"gen1","gateway_rpc_addr":"127.0.0.1:20110","layer":224,"client":"tdesktop","updated_at":100,"expires_at":160}`,
		"wrong:1":   `{"user_id":42,"perm_auth_key_id":9004,"auth_key_id":1004,"auth_key_type":1,"session_id":2005,"gateway_id":"gw1","gateway_generation":"gen1","gateway_rpc_addr":"127.0.0.1:20110","layer":224,"client":"tdesktop","updated_at":100,"expires_at":160}`,
	}
	sessions, cleanup := buildOnlineSessions(context.Background(), 42, raw, 150)
	if len(sessions) != 1 {
		t.Fatalf("len(sessions) = %d, want 1", len(sessions))
	}
	if sessions[0].AuthKeyId != 1001 || sessions[0].SessionId != 2002 {
		t.Fatalf("unexpected session: %+v", sessions[0])
	}
	if len(cleanup) != 4 {
		t.Fatalf("len(cleanup) = %d, want 4", len(cleanup))
	}
}

func TestOnlineSessionCacheDataUsesServiceTime(t *testing.T) {
	in := presencepb.MakeTLOnlineSession(&presencepb.TLOnlineSession{
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
	})
	got := sessionCacheDataFromTL(in, 100, 60)
	if got.UpdatedAt != 100 {
		t.Fatalf("UpdatedAt = %d, want 100", got.UpdatedAt)
	}
	if got.ExpiresAt != 160 {
		t.Fatalf("ExpiresAt = %d, want 160", got.ExpiresAt)
	}
}

func TestSetSessionOnlineNilSessionReturnsInvalidArgument(t *testing.T) {
	repo := &Repository{}
	err := repo.SetSessionOnline(context.Background(), nil, 100, 60, 600, 60)
	if !errors.Is(err, presencepb.ErrPresenceInvalidArgument) {
		t.Fatalf("error = %v, want ErrPresenceInvalidArgument", err)
	}
}

func TestNilRepositoryGetUserOnlineSessionsReturnsStorageError(t *testing.T) {
	var repo *Repository
	_, err := repo.GetUserOnlineSessions(context.Background(), 42, 100)
	if !errors.Is(err, presencepb.ErrPresenceStorage) {
		t.Fatalf("error = %v, want ErrPresenceStorage", err)
	}
}

func TestShouldCleanupOnWriteRequiresPositiveInterval(t *testing.T) {
	if shouldCleanupOnWrite(0) {
		t.Fatal("shouldCleanupOnWrite(0) = true, want false")
	}
	if !shouldCleanupOnWrite(60) {
		t.Fatal("shouldCleanupOnWrite(60) = false, want true")
	}
}
