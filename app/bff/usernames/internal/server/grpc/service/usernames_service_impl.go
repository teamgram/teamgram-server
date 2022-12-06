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
	c.Logger.Debugf("account.checkUsername - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.AccountCheckUsername(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("account.checkUsername - reply: %s", r.DebugString())
	return r, err
}

// AccountUpdateUsername
// account.updateUsername#3e0bdd7c username:string = User;
func (s *Service) AccountUpdateUsername(ctx context.Context, request *mtproto.TLAccountUpdateUsername) (*mtproto.User, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("account.updateUsername - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.AccountUpdateUsername(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("account.updateUsername - reply: %s", r.DebugString())
	return r, err
}

// AccountReorderUsernames
// account.reorderUsernames#ef500eab order:Vector<string> = Bool;
func (s *Service) AccountReorderUsernames(ctx context.Context, request *mtproto.TLAccountReorderUsernames) (*mtproto.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("account.reorderUsernames - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.AccountReorderUsernames(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("account.reorderUsernames - reply: %s", r.DebugString())
	return r, err
}

// AccountToggleUsername
// account.toggleUsername#58d6b376 username:string active:Bool = Bool;
func (s *Service) AccountToggleUsername(ctx context.Context, request *mtproto.TLAccountToggleUsername) (*mtproto.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("account.toggleUsername - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.AccountToggleUsername(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("account.toggleUsername - reply: %s", r.DebugString())
	return r, err
}

// ContactsResolveUsername
// contacts.resolveUsername#f93ccba3 username:string = contacts.ResolvedPeer;
func (s *Service) ContactsResolveUsername(ctx context.Context, request *mtproto.TLContactsResolveUsername) (*mtproto.Contacts_ResolvedPeer, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("contacts.resolveUsername - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.ContactsResolveUsername(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("contacts.resolveUsername - reply: %s", r.DebugString())
	return r, err
}

// ChannelsCheckUsername
// channels.checkUsername#10e6bd2c channel:InputChannel username:string = Bool;
func (s *Service) ChannelsCheckUsername(ctx context.Context, request *mtproto.TLChannelsCheckUsername) (*mtproto.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("channels.checkUsername - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.ChannelsCheckUsername(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("channels.checkUsername - reply: %s", r.DebugString())
	return r, err
}

// ChannelsUpdateUsername
// channels.updateUsername#3514b3de channel:InputChannel username:string = Bool;
func (s *Service) ChannelsUpdateUsername(ctx context.Context, request *mtproto.TLChannelsUpdateUsername) (*mtproto.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("channels.updateUsername - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.ChannelsUpdateUsername(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("channels.updateUsername - reply: %s", r.DebugString())
	return r, err
}

// ChannelsReorderUsernames
// channels.reorderUsernames#b45ced1d channel:InputChannel order:Vector<string> = Bool;
func (s *Service) ChannelsReorderUsernames(ctx context.Context, request *mtproto.TLChannelsReorderUsernames) (*mtproto.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("channels.reorderUsernames - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.ChannelsReorderUsernames(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("channels.reorderUsernames - reply: %s", r.DebugString())
	return r, err
}

// ChannelsToggleUsername
// channels.toggleUsername#50f24105 channel:InputChannel username:string active:Bool = Bool;
func (s *Service) ChannelsToggleUsername(ctx context.Context, request *mtproto.TLChannelsToggleUsername) (*mtproto.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("channels.toggleUsername - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.ChannelsToggleUsername(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("channels.toggleUsername - reply: %s", r.DebugString())
	return r, err
}

// ChannelsDeactivateAllUsernames
// channels.deactivateAllUsernames#a245dd3 channel:InputChannel = Bool;
func (s *Service) ChannelsDeactivateAllUsernames(ctx context.Context, request *mtproto.TLChannelsDeactivateAllUsernames) (*mtproto.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("channels.deactivateAllUsernames - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.ChannelsDeactivateAllUsernames(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("channels.deactivateAllUsernames - reply: %s", r.DebugString())
	return r, err
}
