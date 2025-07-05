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
	"github.com/teamgram/teamgram-server/v2/app/bff/usernames/internal/core"
)

// AccountCheckUsername
// account.checkUsername#2714d86c username:string = Bool;
func (s *Service) AccountCheckUsername(ctx context.Context, request *tg.TLAccountCheckUsername) (*tg.Bool, error) {
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
func (s *Service) AccountUpdateUsername(ctx context.Context, request *tg.TLAccountUpdateUsername) (*tg.User, error) {
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
// contacts.resolveUsername#725afbbc flags:# username:string referer:flags.0?string = contacts.ResolvedPeer;
func (s *Service) ContactsResolveUsername(ctx context.Context, request *tg.TLContactsResolveUsername) (*tg.ContactsResolvedPeer, error) {
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
func (s *Service) ChannelsCheckUsername(ctx context.Context, request *tg.TLChannelsCheckUsername) (*tg.Bool, error) {
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
func (s *Service) ChannelsUpdateUsername(ctx context.Context, request *tg.TLChannelsUpdateUsername) (*tg.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("channels.updateUsername - metadata: %s, request: %s", c.MD, request)

	r, err := c.ChannelsUpdateUsername(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("channels.updateUsername - reply: %s", r)
	return r, err
}
