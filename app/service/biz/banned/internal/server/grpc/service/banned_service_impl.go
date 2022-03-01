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
	"github.com/teamgram/teamgram-server/app/service/biz/banned/banned"
	"github.com/teamgram/teamgram-server/app/service/biz/banned/internal/core"
)

// BannedCheckPhoneNumberBanned
// banned.checkPhoneNumberBanned phone:string = Bool;
func (s *Service) BannedCheckPhoneNumberBanned(ctx context.Context, request *banned.TLBannedCheckPhoneNumberBanned) (*mtproto.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("banned.checkPhoneNumberBanned - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.BannedCheckPhoneNumberBanned(request)
	if err != nil {
		return nil, err
	}

	c.Infof("banned.checkPhoneNumberBanned - reply: %s", r.DebugString())
	return r, err
}

// BannedGetBannedByPhoneList
// banned.getBannedByPhoneList phones:Vector<string> = Vector<string>;
func (s *Service) BannedGetBannedByPhoneList(ctx context.Context, request *banned.TLBannedGetBannedByPhoneList) (*banned.Vector_String, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("banned.getBannedByPhoneList - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.BannedGetBannedByPhoneList(request)
	if err != nil {
		return nil, err
	}

	c.Infof("banned.getBannedByPhoneList - reply: %s", r.DebugString())
	return r, err
}

// BannedBan
// banned.ban phone:string expires:int reason:string = Bool;
func (s *Service) BannedBan(ctx context.Context, request *banned.TLBannedBan) (*mtproto.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("banned.ban - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.BannedBan(request)
	if err != nil {
		return nil, err
	}

	c.Infof("banned.ban - reply: %s", r.DebugString())
	return r, err
}

// BannedUnBan
// banned.unBan phone:string = Bool;
func (s *Service) BannedUnBan(ctx context.Context, request *banned.TLBannedUnBan) (*mtproto.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("banned.unBan - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.BannedUnBan(request)
	if err != nil {
		return nil, err
	}

	c.Infof("banned.unBan - reply: %s", r.DebugString())
	return r, err
}
