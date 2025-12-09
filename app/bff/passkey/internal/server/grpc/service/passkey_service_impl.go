/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright 2025 Teamgram Authors.
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package service

import (
	"context"

	"github.com/teamgram/proto/mtproto"
	"github.com/teamgram/teamgram-server/app/bff/passkey/internal/core"
)

// AuthInitPasskeyLogin
// auth.initPasskeyLogin#518ad0b7 api_id:int api_hash:string = auth.PasskeyLoginOptions;
func (s *Service) AuthInitPasskeyLogin(ctx context.Context, request *mtproto.TLAuthInitPasskeyLogin) (*mtproto.Auth_PasskeyLoginOptions, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("auth.initPasskeyLogin - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.AuthInitPasskeyLogin(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("auth.initPasskeyLogin - reply: {%s}", r)
	return r, err
}

// AuthFinishPasskeyLogin
// auth.finishPasskeyLogin#9857ad07 flags:# credential:InputPasskeyCredential from_dc_id:flags.0?int from_auth_key_id:flags.0?long = auth.Authorization;
func (s *Service) AuthFinishPasskeyLogin(ctx context.Context, request *mtproto.TLAuthFinishPasskeyLogin) (*mtproto.Auth_Authorization, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("auth.finishPasskeyLogin - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.AuthFinishPasskeyLogin(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("auth.finishPasskeyLogin - reply: {%s}", r)
	return r, err
}

// AccountInitPasskeyRegistration
// account.initPasskeyRegistration#429547e8 = account.PasskeyRegistrationOptions;
func (s *Service) AccountInitPasskeyRegistration(ctx context.Context, request *mtproto.TLAccountInitPasskeyRegistration) (*mtproto.Account_PasskeyRegistrationOptions, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("account.initPasskeyRegistration - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.AccountInitPasskeyRegistration(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("account.initPasskeyRegistration - reply: {%s}", r)
	return r, err
}

// AccountRegisterPasskey
// account.registerPasskey#55b41fd6 credential:InputPasskeyCredential = Passkey;
func (s *Service) AccountRegisterPasskey(ctx context.Context, request *mtproto.TLAccountRegisterPasskey) (*mtproto.Passkey, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("account.registerPasskey - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.AccountRegisterPasskey(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("account.registerPasskey - reply: {%s}", r)
	return r, err
}

// AccountGetPasskeys
// account.getPasskeys#ea1f0c52 = account.Passkeys;
func (s *Service) AccountGetPasskeys(ctx context.Context, request *mtproto.TLAccountGetPasskeys) (*mtproto.Account_Passkeys, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("account.getPasskeys - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.AccountGetPasskeys(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("account.getPasskeys - reply: {%s}", r)
	return r, err
}

// AccountDeletePasskey
// account.deletePasskey#f5b5563f id:string = Bool;
func (s *Service) AccountDeletePasskey(ctx context.Context, request *mtproto.TLAccountDeletePasskey) (*mtproto.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("account.deletePasskey - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.AccountDeletePasskey(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("account.deletePasskey - reply: {%s}", r)
	return r, err
}
