/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright 2022 Teamgram Authors.
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package service

import (
	"context"

	"github.com/teamgram/proto/mtproto"
	"github.com/teamgram/teamgram-server/app/bff/usernames/internal/core"
)

// AccountCheckUsername
// account.checkUsername#2714d86c username:string = Bool;
func (s *Service) AccountCheckUsername(ctx context.Context, request *mtproto.TLAccountCheckUsername) (*mtproto.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("account.checkUsername - metadata: %s, request: %s", c.MD, request)

	r, err := c.AccountCheckUsername(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("account.checkUsername - reply: %s", r)
	return r, err
}

// AccountUpdateUsername
// account.updateUsername#3e0bdd7c username:string = User;
func (s *Service) AccountUpdateUsername(ctx context.Context, request *mtproto.TLAccountUpdateUsername) (*mtproto.User, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("account.updateUsername - metadata: %s, request: %s", c.MD, request)

	r, err := c.AccountUpdateUsername(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("account.updateUsername - reply: %s", r)
	return r, err
}

// ContactsResolveUsername
// contacts.resolveUsername#f93ccba3 username:string = contacts.ResolvedPeer;
func (s *Service) ContactsResolveUsername(ctx context.Context, request *mtproto.TLContactsResolveUsername) (*mtproto.Contacts_ResolvedPeer, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("contacts.resolveUsername - metadata: %s, request: %s", c.MD, request)

	r, err := c.ContactsResolveUsername(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("contacts.resolveUsername - reply: %s", r)
	return r, err
}

// ChannelsCheckUsername
// channels.checkUsername#10e6bd2c channel:InputChannel username:string = Bool;
func (s *Service) ChannelsCheckUsername(ctx context.Context, request *mtproto.TLChannelsCheckUsername) (*mtproto.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("channels.checkUsername - metadata: %s, request: %s", c.MD, request)

	r, err := c.ChannelsCheckUsername(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("channels.checkUsername - reply: %s", r)
	return r, err
}

// ChannelsUpdateUsername
// channels.updateUsername#3514b3de channel:InputChannel username:string = Bool;
func (s *Service) ChannelsUpdateUsername(ctx context.Context, request *mtproto.TLChannelsUpdateUsername) (*mtproto.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("channels.updateUsername - metadata: %s, request: %s", c.MD, request)

	r, err := c.ChannelsUpdateUsername(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("channels.updateUsername - reply: %s", r)
	return r, err
}
