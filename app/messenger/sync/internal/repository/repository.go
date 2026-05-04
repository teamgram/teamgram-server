// Copyright (c) 2026 The Teamgram Authors. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Author: teamgramio (teamgram.io@gmail.com)

package repository

import (
	"context"
	"fmt"
	"time"

	gatewayclient "github.com/teamgram/teamgram-server/v2/app/interface/gateway/client"
	gatewaypb "github.com/teamgram/teamgram-server/v2/app/interface/gateway/gateway"
	"github.com/teamgram/teamgram-server/v2/app/messenger/sync/internal/config"
	syncpb "github.com/teamgram/teamgram-server/v2/app/messenger/sync/sync"
	presenceclient "github.com/teamgram/teamgram-server/v2/app/service/presence/client"
	presencepb "github.com/teamgram/teamgram-server/v2/app/service/presence/presence"
	"github.com/teamgram/teamgram-server/v2/pkg/net/kitex"
	"github.com/teamgram/teamgram-server/v2/pkg/net/kitex/identity"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

// Repository is the dependency container for repository instances.
type Repository struct {
	presence      PresenceClient
	gateway       GatewayPusher
	addressPolicy AddressPolicy
}

// NewRepository creates a new Repository.
func NewRepository(c config.Config) *Repository {
	policy := AddressPolicy{
		AllowedCIDRs:     c.GatewayAllowedCIDRs,
		AllowedIPv6CIDRs: c.GatewayAllowedIPv6CIDRs,
		AllowedPorts:     c.GatewayAllowedPorts,
		AllowLoopback:    c.AllowGatewayLoopback,
	}
	return &Repository{
		presence:      newPresenceAdapter(c.PresenceClient),
		gateway:       NewGatewayClientCache(c.GatewayClientBase, policy, c.GatewayClientCacheMaxEntries, time.Duration(c.GatewayClientCacheTTLSeconds)*time.Second, time.Duration(c.GatewayClientIdleTimeoutSeconds)*time.Second),
		addressPolicy: policy,
	}
}

func NewTestRepository(presence PresenceClient, gateway GatewayPusher, policy AddressPolicy) *Repository {
	return &Repository{
		presence:      presence,
		gateway:       gateway,
		addressPolicy: policy,
	}
}

type presenceAdapter struct {
	client presenceclient.PresenceClient
}

func newPresenceAdapter(c kitex.RpcClientConf) PresenceClient {
	if c.DestService == "" && len(c.Endpoints) == 0 && !c.HasEtcd() {
		return nil
	}
	return presenceAdapter{client: presenceclient.NewPresenceClient(presenceclient.MustNewKitexClient(c))}
}

func (p presenceAdapter) GetUserOnlineSessions(ctx context.Context, userID int64) ([]SessionRoute, error) {
	resp, err := p.client.PresenceGetUserOnlineSessions(identity.WithCallerService(ctx, "sync"), &presencepb.TLPresenceGetUserOnlineSessions{UserId: userID})
	if err != nil {
		return nil, fmt.Errorf("%w: get user online sessions: %w", syncpb.ErrSyncPresenceFailure, err)
	}
	routes := make([]SessionRoute, 0, len(resp.Sessions))
	for _, session := range resp.Sessions {
		if session == nil {
			continue
		}
		routes = append(routes, SessionRoute{
			UserID:            session.UserId,
			PermAuthKeyID:     session.PermAuthKeyId,
			AuthKeyID:         session.AuthKeyId,
			AuthKeyType:       session.AuthKeyType,
			SessionID:         session.SessionId,
			GatewayID:         session.GatewayId,
			GatewayGeneration: session.GatewayGeneration,
			GatewayRPCAddr:    session.GatewayRpcAddr,
		})
	}
	return routes, nil
}

type gatewayClientPusher struct {
	client gatewayclient.GatewayClient
}

func (p gatewayClientPusher) PushSessionUpdates(ctx context.Context, route SessionRoute, updates tg.UpdatesClazz) error {
	ok, err := p.client.GatewayPushSessionUpdatesData(identity.WithCallerService(ctx, "sync"), &gatewaypb.TLGatewayPushSessionUpdatesData{
		PermAuthKeyId: route.PermAuthKeyID,
		AuthKeyId:     route.AuthKeyID,
		SessionId:     route.SessionID,
		Updates:       updates,
	})
	if err != nil {
		return err
	}
	if ok == nil {
		return fmt.Errorf("%w: gateway returned nil", syncpb.ErrSyncGatewayFailure)
	}
	if _, ok := ok.ToBoolTrue(); !ok {
		return fmt.Errorf("%w: gateway returned false", syncpb.ErrSyncGatewayFailure)
	}
	return nil
}

func (p gatewayClientPusher) PushRpcResult(ctx context.Context, route RpcResultRoute) error {
	ok, err := p.client.GatewayPushRpcResultData(identity.WithCallerService(ctx, "sync"), &gatewaypb.TLGatewayPushRpcResultData{
		PermAuthKeyId:  route.PermAuthKeyID,
		AuthKeyId:      route.AuthKeyID,
		SessionId:      route.SessionID,
		ClientReqMsgId: route.ClientReqMsgID,
		RpcResultData:  route.RPCResult,
	})
	if err != nil {
		return err
	}
	if ok == nil {
		return fmt.Errorf("%w: gateway returned nil", syncpb.ErrSyncTargetSessionMissing)
	}
	if _, ok := ok.ToBoolTrue(); !ok {
		return fmt.Errorf("%w: gateway returned false", syncpb.ErrSyncTargetSessionMissing)
	}
	return nil
}
