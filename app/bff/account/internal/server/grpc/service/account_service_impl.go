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

	"github.com/teamgram/proto/mtproto"
	"github.com/teamgram/teamgram-server/app/bff/account/internal/core"
)

// AccountDeleteAccount
// account.deleteAccount#a2c0cf74 flags:# reason:string password:flags.0?InputCheckPasswordSRP = Bool;
func (s *Service) AccountDeleteAccount(ctx context.Context, request *mtproto.TLAccountDeleteAccount) (*mtproto.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("account.deleteAccount - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.AccountDeleteAccount(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("account.deleteAccount - reply: {%s}", r)
	return r, err
}

// AccountGetAccountTTL
// account.getAccountTTL#8fc711d = AccountDaysTTL;
func (s *Service) AccountGetAccountTTL(ctx context.Context, request *mtproto.TLAccountGetAccountTTL) (*mtproto.AccountDaysTTL, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("account.getAccountTTL - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.AccountGetAccountTTL(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("account.getAccountTTL - reply: {%s}", r)
	return r, err
}

// AccountSetAccountTTL
// account.setAccountTTL#2442485e ttl:AccountDaysTTL = Bool;
func (s *Service) AccountSetAccountTTL(ctx context.Context, request *mtproto.TLAccountSetAccountTTL) (*mtproto.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("account.setAccountTTL - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.AccountSetAccountTTL(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("account.setAccountTTL - reply: {%s}", r)
	return r, err
}

// AccountSendChangePhoneCode
// account.sendChangePhoneCode#82574ae5 phone_number:string settings:CodeSettings = auth.SentCode;
func (s *Service) AccountSendChangePhoneCode(ctx context.Context, request *mtproto.TLAccountSendChangePhoneCode) (*mtproto.Auth_SentCode, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("account.sendChangePhoneCode - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.AccountSendChangePhoneCode(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("account.sendChangePhoneCode - reply: {%s}", r)
	return r, err
}

// AccountChangePhone
// account.changePhone#70c32edb phone_number:string phone_code_hash:string phone_code:string = User;
func (s *Service) AccountChangePhone(ctx context.Context, request *mtproto.TLAccountChangePhone) (*mtproto.User, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("account.changePhone - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.AccountChangePhone(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("account.changePhone - reply: {%s}", r)
	return r, err
}

// AccountResetAuthorization
// account.resetAuthorization#df77f3bc hash:long = Bool;
func (s *Service) AccountResetAuthorization(ctx context.Context, request *mtproto.TLAccountResetAuthorization) (*mtproto.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("account.resetAuthorization - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.AccountResetAuthorization(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("account.resetAuthorization - reply: {%s}", r)
	return r, err
}

// AccountSendConfirmPhoneCode
// account.sendConfirmPhoneCode#1b3faa88 hash:string settings:CodeSettings = auth.SentCode;
func (s *Service) AccountSendConfirmPhoneCode(ctx context.Context, request *mtproto.TLAccountSendConfirmPhoneCode) (*mtproto.Auth_SentCode, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("account.sendConfirmPhoneCode - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.AccountSendConfirmPhoneCode(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("account.sendConfirmPhoneCode - reply: {%s}", r)
	return r, err
}

// AccountConfirmPhone
// account.confirmPhone#5f2178c3 phone_code_hash:string phone_code:string = Bool;
func (s *Service) AccountConfirmPhone(ctx context.Context, request *mtproto.TLAccountConfirmPhone) (*mtproto.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("account.confirmPhone - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.AccountConfirmPhone(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("account.confirmPhone - reply: {%s}", r)
	return r, err
}
