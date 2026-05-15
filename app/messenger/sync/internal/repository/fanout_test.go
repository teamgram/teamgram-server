package repository

import (
	"context"
	"testing"

	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

func TestFilterPushUpdatesIfNotDistinguishesEmptyIncludes(t *testing.T) {
	targets := []SessionRoute{
		{PermAuthKeyID: 1, AuthKeyID: 10, SessionID: 100, AuthKeyType: int32(tg.AuthKeyTypePerm)},
		{PermAuthKeyID: 2, AuthKeyID: 20, SessionID: 200, AuthKeyType: int32(tg.AuthKeyTypePerm)},
	}
	got := FilterTargets(targets, TargetFilter{IncludesSet: true, Includes: []int64{}})
	if len(got) != 0 {
		t.Fatalf("len(got) = %d, want 0", len(got))
	}
}

func TestFilterUpdatesNotMeExcludesSourcePermKey(t *testing.T) {
	targets := []SessionRoute{
		{PermAuthKeyID: 1, AuthKeyID: 10, SessionID: 100, AuthKeyType: int32(tg.AuthKeyTypePerm)},
		{PermAuthKeyID: 2, AuthKeyID: 20, SessionID: 200, AuthKeyType: int32(tg.AuthKeyTypePerm)},
	}
	got := FilterTargets(targets, TargetFilter{ExcludePermAuthKeyID: 1})
	if len(got) != 1 || got[0].PermAuthKeyID != 2 {
		t.Fatalf("got = %+v, want only perm key 2", got)
	}
}

func TestFilterUpdatesNotMeExcludesSignedSourcePermKey(t *testing.T) {
	targets := []SessionRoute{
		{PermAuthKeyID: -1, AuthKeyID: 10, SessionID: 100, AuthKeyType: int32(tg.AuthKeyTypePerm)},
		{PermAuthKeyID: -2, AuthKeyID: 20, SessionID: 200, AuthKeyType: int32(tg.AuthKeyTypePerm)},
	}
	got := FilterTargets(targets, TargetFilter{ExcludePermAuthKeyID: -1})
	if len(got) != 1 || got[0].PermAuthKeyID != -2 {
		t.Fatalf("got = %+v, want only perm key -2", got)
	}
}

func TestFilterUpdatesMeCanTargetSingleSession(t *testing.T) {
	targets := []SessionRoute{
		{PermAuthKeyID: 1, AuthKeyID: 10, SessionID: 100, AuthKeyType: int32(tg.AuthKeyTypePerm)},
		{PermAuthKeyID: 1, AuthKeyID: 10, SessionID: 101, AuthKeyType: int32(tg.AuthKeyTypePerm)},
	}
	got := FilterTargets(targets, TargetFilter{PermAuthKeyID: 1, AuthKeyID: 10, SessionID: 101, PreciseSession: true})
	if len(got) != 1 || got[0].SessionID != 101 {
		t.Fatalf("got = %+v, want precise session", got)
	}
}

func TestFilterUpdatesMeCanTargetSignedPermKey(t *testing.T) {
	targets := []SessionRoute{
		{PermAuthKeyID: -1, AuthKeyID: 10, SessionID: 100, AuthKeyType: int32(tg.AuthKeyTypePerm)},
		{PermAuthKeyID: -2, AuthKeyID: 20, SessionID: 200, AuthKeyType: int32(tg.AuthKeyTypePerm)},
	}
	got := FilterTargets(targets, TargetFilter{PermAuthKeyID: -2})
	if len(got) != 1 || got[0].PermAuthKeyID != -2 {
		t.Fatalf("got = %+v, want only perm key -2", got)
	}
}

func TestFilterTargetsOnlyNormalAuthSessions(t *testing.T) {
	targets := []SessionRoute{
		{PermAuthKeyID: 1, AuthKeyID: 10, SessionID: 100, AuthKeyType: int32(tg.AuthKeyTypePerm)},
		{PermAuthKeyID: 2, AuthKeyID: 20, SessionID: 200, AuthKeyType: int32(tg.AuthKeyTypeTemp)},
		{PermAuthKeyID: 3, AuthKeyID: 30, SessionID: 300, AuthKeyType: int32(tg.AuthKeyTypeMediaTemp)},
	}
	got := FilterTargets(targets, TargetFilter{})
	if len(got) != 2 {
		t.Fatalf("len(got) = %d, want 2", len(got))
	}
}

func TestPushUpdatesExcludesMediaTempPresenceRoutes(t *testing.T) {
	presence := &fakePresenceClient{routes: []SessionRoute{
		{
			UserID:         42,
			PermAuthKeyID:  1,
			AuthKeyID:      10,
			SessionID:      100,
			AuthKeyType:    int32(tg.AuthKeyTypePerm),
			GatewayID:      "gw-perm",
			GatewayRPCAddr: "127.0.0.1:20110",
		},
		{
			UserID:         42,
			PermAuthKeyID:  1,
			AuthKeyID:      20,
			SessionID:      200,
			AuthKeyType:    int32(tg.AuthKeyTypeMediaTemp),
			GatewayID:      "gw-media",
			GatewayRPCAddr: "127.0.0.1:20110",
		},
	}}
	gateway := &fakeGatewayPusher{}
	repo := NewTestRepository(presence, gateway, allowLocalhostPolicy())

	err := repo.PushUpdates(context.Background(), 42, tg.MakeTLUpdatesTooLong(&tg.TLUpdatesTooLong{}))
	if err != nil {
		t.Fatalf("PushUpdates() error = %v", err)
	}
	if gateway.pushSessionUpdatesCount != 1 {
		t.Fatalf("pushSessionUpdatesCount = %d, want 1", gateway.pushSessionUpdatesCount)
	}
	if gateway.lastSessionRoute.GatewayID != "gw-perm" {
		t.Fatalf("lastSessionRoute.GatewayID = %q, want gw-perm", gateway.lastSessionRoute.GatewayID)
	}
}

func TestUpdatesMePreciseSessionCanRouteMediaTemp(t *testing.T) {
	presence := &fakePresenceClient{routes: []SessionRoute{
		{
			UserID:         42,
			PermAuthKeyID:  1,
			AuthKeyID:      10,
			SessionID:      100,
			AuthKeyType:    int32(tg.AuthKeyTypePerm),
			GatewayID:      "gw-perm",
			GatewayRPCAddr: "127.0.0.1:20110",
		},
		{
			UserID:         42,
			PermAuthKeyID:  1,
			AuthKeyID:      20,
			SessionID:      200,
			AuthKeyType:    int32(tg.AuthKeyTypeMediaTemp),
			GatewayID:      "gw-media",
			GatewayRPCAddr: "127.0.0.1:20110",
		},
	}}
	gateway := &fakeGatewayPusher{}
	repo := NewTestRepository(presence, gateway, allowLocalhostPolicy())

	err := repo.UpdatesMe(context.Background(), 42, 1, 20, 200, true, tg.MakeTLUpdatesTooLong(&tg.TLUpdatesTooLong{}))
	if err != nil {
		t.Fatalf("UpdatesMe() error = %v", err)
	}
	if gateway.pushSessionUpdatesCount != 1 {
		t.Fatalf("pushSessionUpdatesCount = %d, want 1", gateway.pushSessionUpdatesCount)
	}
	if gateway.lastSessionRoute.GatewayID != "gw-media" {
		t.Fatalf("lastSessionRoute.GatewayID = %q, want gw-media", gateway.lastSessionRoute.GatewayID)
	}
}

func TestPushRpcResultUsesExplicitRouteWithoutPresence(t *testing.T) {
	presence := &fakePresenceClient{}
	gateway := &fakeGatewayPusher{}
	repo := NewTestRepository(presence, gateway, allowLocalhostPolicy())
	err := repo.PushRpcResult(context.Background(), RpcResultRoute{
		UserID:            42,
		PermAuthKeyID:     9001,
		AuthKeyID:         1001,
		SessionID:         2002,
		ClientReqMsgID:    3003,
		GatewayID:         "gw1",
		GatewayGeneration: "gen1",
		GatewayRPCAddr:    "127.0.0.1:20110",
		RPCResult:         []byte{1, 2, 3},
	})
	if err != nil {
		t.Fatalf("PushRpcResult() error = %v", err)
	}
	if gateway.pushRpcResultCount != 1 {
		t.Fatalf("pushRpcResultCount = %d, want 1", gateway.pushRpcResultCount)
	}
	if presence.getUserOnlineSessionsCount != 0 {
		t.Fatalf("getUserOnlineSessionsCount = %d, want 0", presence.getUserOnlineSessionsCount)
	}
}

type fakePresenceClient struct {
	routes                     []SessionRoute
	getUserOnlineSessionsCount int
}

func (f *fakePresenceClient) GetUserOnlineSessions(ctx context.Context, userID int64) ([]SessionRoute, error) {
	f.getUserOnlineSessionsCount++
	return f.routes, nil
}

type fakeGatewayPusher struct {
	pushSessionUpdatesCount int
	pushRpcResultCount      int
	lastSessionRoute        SessionRoute
}

func (f *fakeGatewayPusher) PushSessionUpdates(ctx context.Context, route SessionRoute, updates tg.UpdatesClazz) error {
	f.pushSessionUpdatesCount++
	f.lastSessionRoute = route
	return nil
}

func (f *fakeGatewayPusher) PushRpcResult(ctx context.Context, route RpcResultRoute) error {
	f.pushRpcResultCount++
	return nil
}

func allowLocalhostPolicy() AddressPolicy {
	return AddressPolicy{
		AllowedCIDRs:  []string{"127.0.0.0/8"},
		AllowedPorts:  []int{20110},
		AllowLoopback: true,
	}
}
