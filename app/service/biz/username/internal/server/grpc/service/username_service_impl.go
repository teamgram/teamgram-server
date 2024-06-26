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
	"github.com/teamgram/teamgram-server/app/service/biz/username/internal/core"
	"github.com/teamgram/teamgram-server/app/service/biz/username/username"
)

// UsernameGetAccountUsername
// username.getAccountUsername user_id:long = UsernameData;
func (s *Service) UsernameGetAccountUsername(ctx context.Context, request *username.TLUsernameGetAccountUsername) (*username.UsernameData, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("username.getAccountUsername - metadata: %s, request: %s", c.MD, request)

	r, err := c.UsernameGetAccountUsername(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("username.getAccountUsername - reply: %s", r)
	return r, err
}

// UsernameCheckAccountUsername
// username.checkAccountUsername user_id:long username:string = UsernameExist;
func (s *Service) UsernameCheckAccountUsername(ctx context.Context, request *username.TLUsernameCheckAccountUsername) (*username.UsernameExist, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("username.checkAccountUsername - metadata: %s, request: %s", c.MD, request)

	r, err := c.UsernameCheckAccountUsername(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("username.checkAccountUsername - reply: %s", r)
	return r, err
}

// UsernameGetChannelUsername
// username.getChannelUsername channel_id:long = UsernameData;
func (s *Service) UsernameGetChannelUsername(ctx context.Context, request *username.TLUsernameGetChannelUsername) (*username.UsernameData, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("username.getChannelUsername - metadata: %s, request: %s", c.MD, request)

	r, err := c.UsernameGetChannelUsername(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("username.getChannelUsername - reply: %s", r)
	return r, err
}

// UsernameCheckChannelUsername
// username.checkChannelUsername channel_id:long username:string = UsernameExist;
func (s *Service) UsernameCheckChannelUsername(ctx context.Context, request *username.TLUsernameCheckChannelUsername) (*username.UsernameExist, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("username.checkChannelUsername - metadata: %s, request: %s", c.MD, request)

	r, err := c.UsernameCheckChannelUsername(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("username.checkChannelUsername - reply: %s", r)
	return r, err
}

// UsernameUpdateUsernameByPeer
// username.updateUsernameByPeer peer_type:int peer_id:long username:string = Bool;
func (s *Service) UsernameUpdateUsernameByPeer(ctx context.Context, request *username.TLUsernameUpdateUsernameByPeer) (*mtproto.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("username.updateUsernameByPeer - metadata: %s, request: %s", c.MD, request)

	r, err := c.UsernameUpdateUsernameByPeer(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("username.updateUsernameByPeer - reply: %s", r)
	return r, err
}

// UsernameCheckUsername
// username.checkUsername username:string = UsernameExist;
func (s *Service) UsernameCheckUsername(ctx context.Context, request *username.TLUsernameCheckUsername) (*username.UsernameExist, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("username.checkUsername - metadata: %s, request: %s", c.MD, request)

	r, err := c.UsernameCheckUsername(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("username.checkUsername - reply: %s", r)
	return r, err
}

// UsernameUpdateUsername
// username.updateUsername peer_type:int peer_id:long username:string = Bool;
func (s *Service) UsernameUpdateUsername(ctx context.Context, request *username.TLUsernameUpdateUsername) (*mtproto.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("username.updateUsername - metadata: %s, request: %s", c.MD, request)

	r, err := c.UsernameUpdateUsername(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("username.updateUsername - reply: %s", r)
	return r, err
}

// UsernameDeleteUsername
// username.deleteUsername username:string = Bool;
func (s *Service) UsernameDeleteUsername(ctx context.Context, request *username.TLUsernameDeleteUsername) (*mtproto.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("username.deleteUsername - metadata: %s, request: %s", c.MD, request)

	r, err := c.UsernameDeleteUsername(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("username.deleteUsername - reply: %s", r)
	return r, err
}

// UsernameResolveUsername
// username.resolveUsername username:string = Peer;
func (s *Service) UsernameResolveUsername(ctx context.Context, request *username.TLUsernameResolveUsername) (*mtproto.Peer, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("username.resolveUsername - metadata: %s, request: %s", c.MD, request)

	r, err := c.UsernameResolveUsername(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("username.resolveUsername - reply: %s", r)
	return r, err
}

// UsernameGetListByUsernameList
// username.getListByUsernameList names:Vector<string> = Vector<UsernameData>;
func (s *Service) UsernameGetListByUsernameList(ctx context.Context, request *username.TLUsernameGetListByUsernameList) (*username.Vector_UsernameData, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("username.getListByUsernameList - metadata: %s, request: %s", c.MD, request)

	r, err := c.UsernameGetListByUsernameList(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("username.getListByUsernameList - reply: %s", r)
	return r, err
}

// UsernameDeleteUsernameByPeer
// username.deleteUsernameByPeer peer_type:int peer_id:long = Bool;
func (s *Service) UsernameDeleteUsernameByPeer(ctx context.Context, request *username.TLUsernameDeleteUsernameByPeer) (*mtproto.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("username.deleteUsernameByPeer - metadata: %s, request: %s", c.MD, request)

	r, err := c.UsernameDeleteUsernameByPeer(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("username.deleteUsernameByPeer - reply: %s", r)
	return r, err
}

// UsernameSearch
// username.search q:string excluded_contacts:Vector<long> limit:int = Vector<UsernameData>;
func (s *Service) UsernameSearch(ctx context.Context, request *username.TLUsernameSearch) (*username.Vector_UsernameData, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("username.search - metadata: %s, request: %s", c.MD, request)

	r, err := c.UsernameSearch(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("username.search - reply: %s", r)
	return r, err
}
