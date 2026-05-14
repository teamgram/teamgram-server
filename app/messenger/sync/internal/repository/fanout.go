package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/teamgram/teamgram-server/v2/app/messenger/sync/internal/metrics"
	syncpb "github.com/teamgram/teamgram-server/v2/app/messenger/sync/sync"
	"github.com/teamgram/teamgram-server/v2/pkg/net/kitex/identity"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

const (
	MaxIncludePermKeys = 200
	MaxExcludePermKeys = 200
)

type SessionRoute struct {
	UserID            int64
	PermAuthKeyID     int64
	AuthKeyID         int64
	AuthKeyType       int32
	SessionID         int64
	GatewayID         string
	GatewayGeneration string
	GatewayRPCAddr    string
}

type RpcResultRoute struct {
	UserID            int64
	PermAuthKeyID     int64
	AuthKeyID         int64
	SessionID         int64
	ClientReqMsgID    int64
	GatewayID         string
	GatewayGeneration string
	GatewayRPCAddr    string
	RPCResult         []byte
}

type TargetFilter struct {
	PermAuthKeyID        int64
	ExcludePermAuthKeyID int64
	IncludesSet          bool
	Includes             []int64
	ExcludesSet          bool
	Excludes             []int64
	PreciseSession       bool
	AuthKeyID            int64
	SessionID            int64
}

type PresenceClient interface {
	GetUserOnlineSessions(ctx context.Context, userID int64) ([]SessionRoute, error)
}

type GatewayPusher interface {
	PushSessionUpdates(ctx context.Context, route SessionRoute, updates tg.UpdatesClazz) error
	PushRpcResult(ctx context.Context, route RpcResultRoute) error
}

func FilterTargets(targets []SessionRoute, filter TargetFilter) []SessionRoute {
	if filter.IncludesSet && len(filter.Includes) == 0 {
		return nil
	}
	includes := int64Set(filter.Includes)
	excludes := int64Set(filter.Excludes)
	out := make([]SessionRoute, 0, len(targets))
	for _, target := range targets {
		if filter.PermAuthKeyID > 0 && target.PermAuthKeyID != filter.PermAuthKeyID {
			continue
		}
		if filter.ExcludePermAuthKeyID > 0 && target.PermAuthKeyID == filter.ExcludePermAuthKeyID {
			continue
		}
		if filter.PreciseSession && (target.AuthKeyID != filter.AuthKeyID || target.SessionID != filter.SessionID) {
			continue
		}
		if !filter.PreciseSession && !isNormalAuthKeyType(target.AuthKeyType) {
			continue
		}
		if filter.IncludesSet && !includes[target.PermAuthKeyID] {
			continue
		}
		if filter.ExcludesSet && excludes[target.PermAuthKeyID] {
			continue
		}
		out = append(out, target)
	}
	return out
}

func isNormalAuthKeyType(authKeyType int32) bool {
	return authKeyType == int32(tg.AuthKeyTypePerm) || authKeyType == int32(tg.AuthKeyTypeTemp)
}

func int64Set(values []int64) map[int64]bool {
	set := make(map[int64]bool, len(values))
	for _, v := range values {
		set[v] = true
	}
	return set
}

func (r *Repository) PushUpdates(ctx context.Context, userID int64, updates tg.UpdatesClazz) error {
	return r.pushToPresenceTargets(ctx, "pushUpdates", userID, TargetFilter{}, updates)
}

func (r *Repository) UpdatesNotMe(ctx context.Context, userID, permAuthKeyID int64, updates tg.UpdatesClazz) error {
	return r.pushToPresenceTargets(ctx, "updatesNotMe", userID, TargetFilter{ExcludePermAuthKeyID: permAuthKeyID}, updates)
}

func (r *Repository) PushUpdatesIfNot(ctx context.Context, userID int64, includesSet bool, includes []int64, excludesSet bool, excludes []int64, updates tg.UpdatesClazz) error {
	return r.pushToPresenceTargets(ctx, "pushUpdatesIfNot", userID, TargetFilter{
		IncludesSet: includesSet,
		Includes:    includes,
		ExcludesSet: excludesSet,
		Excludes:    excludes,
	}, updates)
}

func (r *Repository) UpdatesMe(ctx context.Context, userID, permAuthKeyID int64, authKeyID, sessionID int64, precise bool, updates tg.UpdatesClazz) error {
	return r.pushToPresenceTargets(ctx, "updatesMe", userID, TargetFilter{
		PermAuthKeyID:  permAuthKeyID,
		AuthKeyID:      authKeyID,
		SessionID:      sessionID,
		PreciseSession: precise,
	}, updates)
}

func (r *Repository) PushRpcResult(ctx context.Context, route RpcResultRoute) error {
	if _, err := r.addressPolicy.Validate(ctx, route.GatewayRPCAddr); err != nil {
		return err
	}
	if r.gateway == nil {
		return fmt.Errorf("%w: gateway pusher is nil", syncpb.ErrSyncGatewayFailure)
	}
	if err := r.gateway.PushRpcResult(identity.WithCallerService(ctx, "sync"), route); err != nil {
		if errors.Is(err, syncpb.ErrSyncTargetSessionMissing) {
			metrics.RpcResultTargetMissing(gatewayIDLabel(route.GatewayID), "missing")
		}
		metrics.FanoutFailure("pushRpcResult", gatewayIDLabel(route.GatewayID), "gateway")
		return fmt.Errorf("%w: %w", syncpb.ErrSyncRpcResultDeliveryFailed, err)
	}
	return nil
}

func (r *Repository) pushToPresenceTargets(ctx context.Context, method string, userID int64, filter TargetFilter, updates tg.UpdatesClazz) error {
	if r.presence == nil {
		return fmt.Errorf("%w: presence client is nil", syncpb.ErrSyncPresenceFailure)
	}
	routes, err := r.presence.GetUserOnlineSessions(identity.WithCallerService(ctx, "sync"), userID)
	if err != nil {
		return fmt.Errorf("%w: %w", syncpb.ErrSyncPresenceFailure, err)
	}
	targets := FilterTargets(routes, filter)
	if len(targets) > 0 && r.gateway == nil {
		return fmt.Errorf("%w: gateway pusher is nil", syncpb.ErrSyncGatewayFailure)
	}
	for _, target := range targets {
		if _, err := r.addressPolicy.Validate(ctx, target.GatewayRPCAddr); err != nil {
			return err
		}
		if err := r.gateway.PushSessionUpdates(identity.WithCallerService(ctx, "sync"), target, updates); err != nil {
			metrics.FanoutFailure(method, gatewayIDLabel(target.GatewayID), "gateway")
			return fmt.Errorf("%w: %w", syncpb.ErrSyncGatewayFailure, err)
		}
	}
	return nil
}

func gatewayIDLabel(gatewayID string) string {
	if gatewayID == "" {
		return "unknown"
	}
	return gatewayID
}
