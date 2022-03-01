/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2021-present,  Teamgram Studio (https://teamgram.io).
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package service

import (
	"context"

	"github.com/teamgram/proto/mtproto"
	"github.com/teamgram/teamgram-server/app/service/status/internal/core"
	"github.com/teamgram/teamgram-server/app/service/status/status"
)

// StatusSetSessionOnline
// status.setSessionOnline user_id:long auth_key_id:long gateway:string expired:long layer:int = Bool;
func (s *Service) StatusSetSessionOnline(ctx context.Context, request *status.TLStatusSetSessionOnline) (*mtproto.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("status.setSessionOnline - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.StatusSetSessionOnline(request)
	if err != nil {
		return nil, err
	}

	c.Infof("status.setSessionOnline - reply: %s", r.DebugString())
	return r, err
}

// StatusSetSessionOffline
// status.setSessionOffline user_id:long auth_key_id:long = Bool;
func (s *Service) StatusSetSessionOffline(ctx context.Context, request *status.TLStatusSetSessionOffline) (*mtproto.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("status.setSessionOffline - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.StatusSetSessionOffline(request)
	if err != nil {
		return nil, err
	}

	c.Infof("status.setSessionOffline - reply: %s", r.DebugString())
	return r, err
}

// StatusGetUserOnlineSessions
// status.getUserOnlineSessions user_id:long = UserSessionEntryList;
func (s *Service) StatusGetUserOnlineSessions(ctx context.Context, request *status.TLStatusGetUserOnlineSessions) (*status.UserSessionEntryList, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("status.getUserOnlineSessions - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.StatusGetUserOnlineSessions(request)
	if err != nil {
		return nil, err
	}

	c.Infof("status.getUserOnlineSessions - reply: %s", r.DebugString())
	return r, err
}

// StatusGetUsersOnlineSessionsList
// status.getUsersOnlineSessionsList Vector<long> = Vector<UserSessionEntryList>;
func (s *Service) StatusGetUsersOnlineSessionsList(ctx context.Context, request *status.TLStatusGetUsersOnlineSessionsList) (*status.Vector_UserSessionEntryList, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("status.getUsersOnlineSessionsList - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.StatusGetUsersOnlineSessionsList(request)
	if err != nil {
		return nil, err
	}

	c.Infof("status.getUsersOnlineSessionsList - reply: %s", r.DebugString())
	return r, err
}
