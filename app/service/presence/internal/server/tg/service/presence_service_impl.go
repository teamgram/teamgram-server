/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2026-present, The Teamgram Authors (https://teamgram.net).
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package service

import (
	"context"

	"github.com/teamgram/teamgram-server/v2/app/service/presence/internal/core"
	"github.com/teamgram/teamgram-server/v2/app/service/presence/presence"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

var _ *tg.Bool

// PresenceSetSessionOnline
// presence.setSessionOnline session:OnlineSession = Bool;
func (s *Service) PresenceSetSessionOnline(ctx context.Context, request *presence.TLPresenceSetSessionOnline) (*tg.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("presence.setSessionOnline - request: %s", request)

	r, err := c.PresenceSetSessionOnline(request)
	if err != nil {
		c.Logger.Errorf("presence.setSessionOnline - error: request: %s, err: %v", request, err)
		return nil, err
	}

	c.Logger.Debugf("presence.setSessionOnline - reply: %s", r)
	return r, err
}

// PresenceSetSessionOffline
// presence.setSessionOffline user_id:long auth_key_id:long session_id:long = Bool;
func (s *Service) PresenceSetSessionOffline(ctx context.Context, request *presence.TLPresenceSetSessionOffline) (*tg.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("presence.setSessionOffline - request: %s", request)

	r, err := c.PresenceSetSessionOffline(request)
	if err != nil {
		c.Logger.Errorf("presence.setSessionOffline - error: request: %s, err: %v", request, err)
		return nil, err
	}

	c.Logger.Debugf("presence.setSessionOffline - reply: %s", r)
	return r, err
}

// PresenceGetUserOnlineSessions
// presence.getUserOnlineSessions user_id:long = UserOnlineSessions;
func (s *Service) PresenceGetUserOnlineSessions(ctx context.Context, request *presence.TLPresenceGetUserOnlineSessions) (*presence.UserOnlineSessions, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("presence.getUserOnlineSessions - request: %s", request)

	r, err := c.PresenceGetUserOnlineSessions(request)
	if err != nil {
		c.Logger.Errorf("presence.getUserOnlineSessions - error: request: %s, err: %v", request, err)
		return nil, err
	}

	c.Logger.Debugf("presence.getUserOnlineSessions - reply: %s", r)
	return r, err
}

// PresenceGetUsersOnlineSessions
// presence.getUsersOnlineSessions users:Vector<long> = Vector<UserOnlineSessions>;
func (s *Service) PresenceGetUsersOnlineSessions(ctx context.Context, request *presence.TLPresenceGetUsersOnlineSessions) (*presence.VectorUserOnlineSessions, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("presence.getUsersOnlineSessions - request: %s", request)

	r, err := c.PresenceGetUsersOnlineSessions(request)
	if err != nil {
		c.Logger.Errorf("presence.getUsersOnlineSessions - error: request: %s, err: %v", request, err)
		return nil, err
	}

	c.Logger.Debugf("presence.getUsersOnlineSessions - reply: %s", r)
	return r, err
}

// PresenceGetGatewaySessions
// presence.getGatewaySessions gateway_id:string = Vector<OnlineSession>;
func (s *Service) PresenceGetGatewaySessions(ctx context.Context, request *presence.TLPresenceGetGatewaySessions) (*presence.VectorOnlineSession, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("presence.getGatewaySessions - request: %s", request)

	r, err := c.PresenceGetGatewaySessions(request)
	if err != nil {
		c.Logger.Errorf("presence.getGatewaySessions - error: request: %s, err: %v", request, err)
		return nil, err
	}

	c.Logger.Debugf("presence.getGatewaySessions - reply: %s", r)
	return r, err
}
