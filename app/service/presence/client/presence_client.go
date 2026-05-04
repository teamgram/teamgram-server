/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2026-present, The Teamgram Authors (https://teamgram.net).
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package presenceclient

import (
	"context"

	"github.com/teamgram/teamgram-server/v2/app/service/presence/presence"
	"github.com/teamgram/teamgram-server/v2/app/service/presence/presence/presenceservice"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"

	"github.com/cloudwego/kitex/client"
)

var _ *tg.Bool

type PresenceClient interface {
	PresenceSetSessionOnline(ctx context.Context, in *presence.TLPresenceSetSessionOnline) (*tg.Bool, error)
	PresenceSetSessionOffline(ctx context.Context, in *presence.TLPresenceSetSessionOffline) (*tg.Bool, error)
	PresenceGetUserOnlineSessions(ctx context.Context, in *presence.TLPresenceGetUserOnlineSessions) (*presence.UserOnlineSessions, error)
	PresenceGetUsersOnlineSessions(ctx context.Context, in *presence.TLPresenceGetUsersOnlineSessions) (*presence.VectorUserOnlineSessions, error)
	PresenceGetGatewaySessions(ctx context.Context, in *presence.TLPresenceGetGatewaySessions) (*presence.VectorOnlineSession, error)
}

type defaultPresenceClient struct {
	cli client.Client
	rpc presenceservice.Client
}

func NewPresenceClient(cli client.Client) PresenceClient {
	return &defaultPresenceClient{
		cli: cli,
		rpc: presenceservice.NewRPCPresenceClient(cli),
	}
}

// PresenceSetSessionOnline
// presence.setSessionOnline session:OnlineSession = Bool;
func (m *defaultPresenceClient) PresenceSetSessionOnline(ctx context.Context, in *presence.TLPresenceSetSessionOnline) (*tg.Bool, error) {
	return m.rpc.PresenceSetSessionOnline(ctx, in)
}

// PresenceSetSessionOffline
// presence.setSessionOffline user_id:long auth_key_id:long session_id:long = Bool;
func (m *defaultPresenceClient) PresenceSetSessionOffline(ctx context.Context, in *presence.TLPresenceSetSessionOffline) (*tg.Bool, error) {
	return m.rpc.PresenceSetSessionOffline(ctx, in)
}

// PresenceGetUserOnlineSessions
// presence.getUserOnlineSessions user_id:long = UserOnlineSessions;
func (m *defaultPresenceClient) PresenceGetUserOnlineSessions(ctx context.Context, in *presence.TLPresenceGetUserOnlineSessions) (*presence.UserOnlineSessions, error) {
	return m.rpc.PresenceGetUserOnlineSessions(ctx, in)
}

// PresenceGetUsersOnlineSessions
// presence.getUsersOnlineSessions users:Vector<long> = Vector<UserOnlineSessions>;
func (m *defaultPresenceClient) PresenceGetUsersOnlineSessions(ctx context.Context, in *presence.TLPresenceGetUsersOnlineSessions) (*presence.VectorUserOnlineSessions, error) {
	return m.rpc.PresenceGetUsersOnlineSessions(ctx, in)
}

// PresenceGetGatewaySessions
// presence.getGatewaySessions gateway_id:string = Vector<OnlineSession>;
func (m *defaultPresenceClient) PresenceGetGatewaySessions(ctx context.Context, in *presence.TLPresenceGetGatewaySessions) (*presence.VectorOnlineSession, error) {
	return m.rpc.PresenceGetGatewaySessions(ctx, in)
}
