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
	"github.com/teamgram/teamgram-server/app/bff/notification/internal/core"
)

// AccountRegisterDevice
// account.registerDevice#ec86017a flags:# no_muted:flags.0?true token_type:int token:string app_sandbox:Bool secret:bytes other_uids:Vector<long> = Bool;
func (s *Service) AccountRegisterDevice(ctx context.Context, request *mtproto.TLAccountRegisterDevice) (*mtproto.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("account.registerDevice - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.AccountRegisterDevice(request)
	if err != nil {
		return nil, err
	}

	c.Infof("account.registerDevice - reply: %s", r.DebugString())
	return r, err
}

// AccountUnregisterDevice
// account.unregisterDevice#6a0d3206 token_type:int token:string other_uids:Vector<long> = Bool;
func (s *Service) AccountUnregisterDevice(ctx context.Context, request *mtproto.TLAccountUnregisterDevice) (*mtproto.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("account.unregisterDevice - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.AccountUnregisterDevice(request)
	if err != nil {
		return nil, err
	}

	c.Infof("account.unregisterDevice - reply: %s", r.DebugString())
	return r, err
}

// AccountUpdateNotifySettings
// account.updateNotifySettings#84be5b93 peer:InputNotifyPeer settings:InputPeerNotifySettings = Bool;
func (s *Service) AccountUpdateNotifySettings(ctx context.Context, request *mtproto.TLAccountUpdateNotifySettings) (*mtproto.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("account.updateNotifySettings - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.AccountUpdateNotifySettings(request)
	if err != nil {
		return nil, err
	}

	c.Infof("account.updateNotifySettings - reply: %s", r.DebugString())
	return r, err
}

// AccountGetNotifySettings
// account.getNotifySettings#12b3ad31 peer:InputNotifyPeer = PeerNotifySettings;
func (s *Service) AccountGetNotifySettings(ctx context.Context, request *mtproto.TLAccountGetNotifySettings) (*mtproto.PeerNotifySettings, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("account.getNotifySettings - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.AccountGetNotifySettings(request)
	if err != nil {
		return nil, err
	}

	c.Infof("account.getNotifySettings - reply: %s", r.DebugString())
	return r, err
}

// AccountResetNotifySettings
// account.resetNotifySettings#db7e1747 = Bool;
func (s *Service) AccountResetNotifySettings(ctx context.Context, request *mtproto.TLAccountResetNotifySettings) (*mtproto.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("account.resetNotifySettings - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.AccountResetNotifySettings(request)
	if err != nil {
		return nil, err
	}

	c.Infof("account.resetNotifySettings - reply: %s", r.DebugString())
	return r, err
}

// AccountUpdateDeviceLocked
// account.updateDeviceLocked#38df3532 period:int = Bool;
func (s *Service) AccountUpdateDeviceLocked(ctx context.Context, request *mtproto.TLAccountUpdateDeviceLocked) (*mtproto.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("account.updateDeviceLocked - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.AccountUpdateDeviceLocked(request)
	if err != nil {
		return nil, err
	}

	c.Infof("account.updateDeviceLocked - reply: %s", r.DebugString())
	return r, err
}

// AccountGetNotifyExceptions
// account.getNotifyExceptions#53577479 flags:# compare_sound:flags.1?true peer:flags.0?InputNotifyPeer = Updates;
func (s *Service) AccountGetNotifyExceptions(ctx context.Context, request *mtproto.TLAccountGetNotifyExceptions) (*mtproto.Updates, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("account.getNotifyExceptions - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.AccountGetNotifyExceptions(request)
	if err != nil {
		return nil, err
	}

	c.Infof("account.getNotifyExceptions - reply: %s", r.DebugString())
	return r, err
}
