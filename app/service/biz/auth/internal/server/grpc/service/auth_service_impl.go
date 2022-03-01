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
	"github.com/teamgram/teamgram-server/app/service/biz/auth/auth"
	"github.com/teamgram/teamgram-server/app/service/biz/auth/internal/core"
)

// AuthExportLoginToken
// auth.exportLoginToken api_id:int api_hash:string except_ids:Vector<long> = auth.LoginToken;
func (s *Service) AuthExportLoginToken(ctx context.Context, request *auth.TLAuthExportLoginToken) (*mtproto.Auth_LoginToken, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("auth.exportLoginToken - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.AuthExportLoginToken(request)
	if err != nil {
		return nil, err
	}

	c.Infof("auth.exportLoginToken - reply: %s", r.DebugString())
	return r, err
}

// AuthImportLoginToken
// auth.importLoginToken token:bytes = auth.LoginToken;
func (s *Service) AuthImportLoginToken(ctx context.Context, request *auth.TLAuthImportLoginToken) (*mtproto.Auth_LoginToken, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("auth.importLoginToken - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.AuthImportLoginToken(request)
	if err != nil {
		return nil, err
	}

	c.Infof("auth.importLoginToken - reply: %s", r.DebugString())
	return r, err
}

// AuthAcceptLoginToken
// auth.acceptLoginToken token:bytes = Authorization;
func (s *Service) AuthAcceptLoginToken(ctx context.Context, request *auth.TLAuthAcceptLoginToken) (*mtproto.Authorization, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("auth.acceptLoginToken - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.AuthAcceptLoginToken(request)
	if err != nil {
		return nil, err
	}

	c.Infof("auth.acceptLoginToken - reply: %s", r.DebugString())
	return r, err
}
