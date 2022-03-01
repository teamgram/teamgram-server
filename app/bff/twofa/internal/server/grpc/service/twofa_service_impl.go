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
	"github.com/teamgram/teamgram-server/app/bff/twofa/internal/core"
)

// AccountGetPassword
// account.getPassword#548a30f5 = account.Password;
func (s *Service) AccountGetPassword(ctx context.Context, request *mtproto.TLAccountGetPassword) (*mtproto.Account_Password, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("account.getPassword - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.AccountGetPassword(request)
	if err != nil {
		return nil, err
	}

	c.Infof("account.getPassword - reply: %s", r.DebugString())
	return r, err
}

// AccountGetPasswordSettings
// account.getPasswordSettings#9cd4eaf9 password:InputCheckPasswordSRP = account.PasswordSettings;
func (s *Service) AccountGetPasswordSettings(ctx context.Context, request *mtproto.TLAccountGetPasswordSettings) (*mtproto.Account_PasswordSettings, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("account.getPasswordSettings - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.AccountGetPasswordSettings(request)
	if err != nil {
		return nil, err
	}

	c.Infof("account.getPasswordSettings - reply: %s", r.DebugString())
	return r, err
}

// AccountUpdatePasswordSettings
// account.updatePasswordSettings#a59b102f password:InputCheckPasswordSRP new_settings:account.PasswordInputSettings = Bool;
func (s *Service) AccountUpdatePasswordSettings(ctx context.Context, request *mtproto.TLAccountUpdatePasswordSettings) (*mtproto.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("account.updatePasswordSettings - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.AccountUpdatePasswordSettings(request)
	if err != nil {
		return nil, err
	}

	c.Infof("account.updatePasswordSettings - reply: %s", r.DebugString())
	return r, err
}

// AccountConfirmPasswordEmail
// account.confirmPasswordEmail#8fdf1920 code:string = Bool;
func (s *Service) AccountConfirmPasswordEmail(ctx context.Context, request *mtproto.TLAccountConfirmPasswordEmail) (*mtproto.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("account.confirmPasswordEmail - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.AccountConfirmPasswordEmail(request)
	if err != nil {
		return nil, err
	}

	c.Infof("account.confirmPasswordEmail - reply: %s", r.DebugString())
	return r, err
}

// AccountResendPasswordEmail
// account.resendPasswordEmail#7a7f2a15 = Bool;
func (s *Service) AccountResendPasswordEmail(ctx context.Context, request *mtproto.TLAccountResendPasswordEmail) (*mtproto.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("account.resendPasswordEmail - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.AccountResendPasswordEmail(request)
	if err != nil {
		return nil, err
	}

	c.Infof("account.resendPasswordEmail - reply: %s", r.DebugString())
	return r, err
}

// AccountCancelPasswordEmail
// account.cancelPasswordEmail#c1cbd5b6 = Bool;
func (s *Service) AccountCancelPasswordEmail(ctx context.Context, request *mtproto.TLAccountCancelPasswordEmail) (*mtproto.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("account.cancelPasswordEmail - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.AccountCancelPasswordEmail(request)
	if err != nil {
		return nil, err
	}

	c.Infof("account.cancelPasswordEmail - reply: %s", r.DebugString())
	return r, err
}

// AccountDeclinePasswordReset
// account.declinePasswordReset#4c9409f6 = Bool;
func (s *Service) AccountDeclinePasswordReset(ctx context.Context, request *mtproto.TLAccountDeclinePasswordReset) (*mtproto.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("account.declinePasswordReset - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.AccountDeclinePasswordReset(request)
	if err != nil {
		return nil, err
	}

	c.Infof("account.declinePasswordReset - reply: %s", r.DebugString())
	return r, err
}
