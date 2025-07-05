/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright 2024 Teamgram Authors.
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package service

import (
	"context"

	"github.com/teamgram/proto/v2/tg"
	"github.com/teamgram/teamgram-server/v2/app/service/status/internal/core"
	"github.com/teamgram/teamgram-server/v2/app/service/status/status"
)

var _ *tg.Bool

// StatusSetSessionOnline
// status.setSessionOnline user_id:long session:SessionEntry = Bool;
func (s *Service) StatusSetSessionOnline(ctx context.Context, request *status.TLStatusSetSessionOnline) (*tg.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("status.setSessionOnline - metadata: %s, request: %s", c.MD, request)

	r, err := c.StatusSetSessionOnline(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("status.setSessionOnline - reply: %s", r)
	return r, err
}

// StatusSetSessionOffline
// status.setSessionOffline user_id:long auth_key_id:long = Bool;
func (s *Service) StatusSetSessionOffline(ctx context.Context, request *status.TLStatusSetSessionOffline) (*tg.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("status.setSessionOffline - metadata: %s, request: %s", c.MD, request)

	r, err := c.StatusSetSessionOffline(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("status.setSessionOffline - reply: %s", r)
	return r, err
}

// StatusGetUserOnlineSessions
// status.getUserOnlineSessions user_id:long = UserSessionEntryList;
func (s *Service) StatusGetUserOnlineSessions(ctx context.Context, request *status.TLStatusGetUserOnlineSessions) (*status.UserSessionEntryList, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("status.getUserOnlineSessions - metadata: %s, request: %s", c.MD, request)

	r, err := c.StatusGetUserOnlineSessions(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("status.getUserOnlineSessions - reply: %s", r)
	return r, err
}

// StatusGetUsersOnlineSessionsList
// status.getUsersOnlineSessionsList users:Vector<long> = Vector<UserSessionEntryList>;
func (s *Service) StatusGetUsersOnlineSessionsList(ctx context.Context, request *status.TLStatusGetUsersOnlineSessionsList) (*status.VectorUserSessionEntryList, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("status.getUsersOnlineSessionsList - metadata: %s, request: %s", c.MD, request)

	r, err := c.StatusGetUsersOnlineSessionsList(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("status.getUsersOnlineSessionsList - reply: %s", r)
	return r, err
}

// StatusGetChannelOnlineUsers
// status.getChannelOnlineUsers channel_id:long = Vector<long>;
func (s *Service) StatusGetChannelOnlineUsers(ctx context.Context, request *status.TLStatusGetChannelOnlineUsers) (*status.VectorLong, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("status.getChannelOnlineUsers - metadata: %s, request: %s", c.MD, request)

	r, err := c.StatusGetChannelOnlineUsers(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("status.getChannelOnlineUsers - reply: %s", r)
	return r, err
}

// StatusSetUserChannelsOnline
// status.setUserChannelsOnline user_id:long channels:Vector<long> = Bool;
func (s *Service) StatusSetUserChannelsOnline(ctx context.Context, request *status.TLStatusSetUserChannelsOnline) (*tg.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("status.setUserChannelsOnline - metadata: %s, request: %s", c.MD, request)

	r, err := c.StatusSetUserChannelsOnline(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("status.setUserChannelsOnline - reply: %s", r)
	return r, err
}

// StatusSetUserChannelsOffline
// status.setUserChannelsOffline user_id:long channels:Vector<long> = Bool;
func (s *Service) StatusSetUserChannelsOffline(ctx context.Context, request *status.TLStatusSetUserChannelsOffline) (*tg.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("status.setUserChannelsOffline - metadata: %s, request: %s", c.MD, request)

	r, err := c.StatusSetUserChannelsOffline(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("status.setUserChannelsOffline - reply: %s", r)
	return r, err
}

// StatusSetChannelUserOffline
// status.setChannelUserOffline channel_id:long user_id:long = Bool;
func (s *Service) StatusSetChannelUserOffline(ctx context.Context, request *status.TLStatusSetChannelUserOffline) (*tg.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("status.setChannelUserOffline - metadata: %s, request: %s", c.MD, request)

	r, err := c.StatusSetChannelUserOffline(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("status.setChannelUserOffline - reply: %s", r)
	return r, err
}

// StatusSetChannelUsersOnline
// status.setChannelUsersOnline channel_id:long id:Vector<long> = Bool;
func (s *Service) StatusSetChannelUsersOnline(ctx context.Context, request *status.TLStatusSetChannelUsersOnline) (*tg.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("status.setChannelUsersOnline - metadata: %s, request: %s", c.MD, request)

	r, err := c.StatusSetChannelUsersOnline(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("status.setChannelUsersOnline - reply: %s", r)
	return r, err
}

// StatusSetChannelOffline
// status.setChannelOffline channel_id:long = Bool;
func (s *Service) StatusSetChannelOffline(ctx context.Context, request *status.TLStatusSetChannelOffline) (*tg.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("status.setChannelOffline - metadata: %s, request: %s", c.MD, request)

	r, err := c.StatusSetChannelOffline(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("status.setChannelOffline - reply: %s", r)
	return r, err
}
