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
	"github.com/teamgram/teamgram-server/app/bff/passport/internal/core"
)

// AccountGetAuthorizations
// account.getAuthorizations#e320c158 = account.Authorizations;
func (s *Service) AccountGetAuthorizations(ctx context.Context, request *mtproto.TLAccountGetAuthorizations) (*mtproto.Account_Authorizations, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("account.getAuthorizations - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.AccountGetAuthorizations(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("account.getAuthorizations - reply: {%s}", r)
	return r, err
}

// AccountGetAllSecureValues
// account.getAllSecureValues#b288bc7d = Vector<SecureValue>;
func (s *Service) AccountGetAllSecureValues(ctx context.Context, request *mtproto.TLAccountGetAllSecureValues) (*mtproto.Vector_SecureValue, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("account.getAllSecureValues - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.AccountGetAllSecureValues(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("account.getAllSecureValues - reply: {%s}", r)
	return r, err
}

// AccountGetSecureValue
// account.getSecureValue#73665bc2 types:Vector<SecureValueType> = Vector<SecureValue>;
func (s *Service) AccountGetSecureValue(ctx context.Context, request *mtproto.TLAccountGetSecureValue) (*mtproto.Vector_SecureValue, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("account.getSecureValue - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.AccountGetSecureValue(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("account.getSecureValue - reply: {%s}", r)
	return r, err
}

// AccountSaveSecureValue
// account.saveSecureValue#899fe31d value:InputSecureValue secure_secret_id:long = SecureValue;
func (s *Service) AccountSaveSecureValue(ctx context.Context, request *mtproto.TLAccountSaveSecureValue) (*mtproto.SecureValue, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("account.saveSecureValue - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.AccountSaveSecureValue(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("account.saveSecureValue - reply: {%s}", r)
	return r, err
}

// AccountDeleteSecureValue
// account.deleteSecureValue#b880bc4b types:Vector<SecureValueType> = Bool;
func (s *Service) AccountDeleteSecureValue(ctx context.Context, request *mtproto.TLAccountDeleteSecureValue) (*mtproto.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("account.deleteSecureValue - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.AccountDeleteSecureValue(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("account.deleteSecureValue - reply: {%s}", r)
	return r, err
}

// AccountGetAuthorizationForm
// account.getAuthorizationForm#a929597a bot_id:long scope:string public_key:string = account.AuthorizationForm;
func (s *Service) AccountGetAuthorizationForm(ctx context.Context, request *mtproto.TLAccountGetAuthorizationForm) (*mtproto.Account_AuthorizationForm, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("account.getAuthorizationForm - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.AccountGetAuthorizationForm(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("account.getAuthorizationForm - reply: {%s}", r)
	return r, err
}

// AccountAcceptAuthorization
// account.acceptAuthorization#f3ed4c73 bot_id:long scope:string public_key:string value_hashes:Vector<SecureValueHash> credentials:SecureCredentialsEncrypted = Bool;
func (s *Service) AccountAcceptAuthorization(ctx context.Context, request *mtproto.TLAccountAcceptAuthorization) (*mtproto.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("account.acceptAuthorization - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.AccountAcceptAuthorization(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("account.acceptAuthorization - reply: {%s}", r)
	return r, err
}

// AccountSendVerifyPhoneCode
// account.sendVerifyPhoneCode#a5a356f9 phone_number:string settings:CodeSettings = auth.SentCode;
func (s *Service) AccountSendVerifyPhoneCode(ctx context.Context, request *mtproto.TLAccountSendVerifyPhoneCode) (*mtproto.Auth_SentCode, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("account.sendVerifyPhoneCode - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.AccountSendVerifyPhoneCode(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("account.sendVerifyPhoneCode - reply: {%s}", r)
	return r, err
}

// AccountVerifyPhone
// account.verifyPhone#4dd3a7f6 phone_number:string phone_code_hash:string phone_code:string = Bool;
func (s *Service) AccountVerifyPhone(ctx context.Context, request *mtproto.TLAccountVerifyPhone) (*mtproto.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("account.verifyPhone - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.AccountVerifyPhone(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("account.verifyPhone - reply: {%s}", r)
	return r, err
}

// UsersSetSecureValueErrors
// users.setSecureValueErrors#90c894b5 id:InputUser errors:Vector<SecureValueError> = Bool;
func (s *Service) UsersSetSecureValueErrors(ctx context.Context, request *mtproto.TLUsersSetSecureValueErrors) (*mtproto.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("users.setSecureValueErrors - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.UsersSetSecureValueErrors(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("users.setSecureValueErrors - reply: {%s}", r)
	return r, err
}

// HelpGetPassportConfig
// help.getPassportConfig#c661ad08 hash:int = help.PassportConfig;
func (s *Service) HelpGetPassportConfig(ctx context.Context, request *mtproto.TLHelpGetPassportConfig) (*mtproto.Help_PassportConfig, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("help.getPassportConfig - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.HelpGetPassportConfig(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("help.getPassportConfig - reply: {%s}", r)
	return r, err
}
