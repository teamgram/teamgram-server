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
	"github.com/teamgram/teamgram-server/app/bff/authorization/internal/core"
)

// AuthSendCode
// auth.sendCode#a677244f phone_number:string api_id:int api_hash:string settings:CodeSettings = auth.SentCode;
func (s *Service) AuthSendCode(ctx context.Context, request *mtproto.TLAuthSendCode) (*mtproto.Auth_SentCode, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("auth.sendCode - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.AuthSendCode(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("auth.sendCode - reply: %s", r.DebugString())
	return r, err
}

// AuthSignUp
// auth.signUp#80eee427 phone_number:string phone_code_hash:string first_name:string last_name:string = auth.Authorization;
func (s *Service) AuthSignUp(ctx context.Context, request *mtproto.TLAuthSignUp) (*mtproto.Auth_Authorization, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("auth.signUp - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.AuthSignUp(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("auth.signUp - reply: %s", r.DebugString())
	return r, err
}

// AuthSignIn
// auth.signIn#bcd51581 phone_number:string phone_code_hash:string phone_code:string = auth.Authorization;
func (s *Service) AuthSignIn(ctx context.Context, request *mtproto.TLAuthSignIn) (*mtproto.Auth_Authorization, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("auth.signIn - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.AuthSignIn(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("auth.signIn - reply: %s", r.DebugString())
	return r, err
}

// AuthLogOut
// auth.logOut#3e72ba19 = auth.LoggedOut;
func (s *Service) AuthLogOut(ctx context.Context, request *mtproto.TLAuthLogOut) (*mtproto.Auth_LoggedOut, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("auth.logOut - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.AuthLogOut(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("auth.logOut - reply: %s", r.DebugString())
	return r, err
}

// AuthResetAuthorizations
// auth.resetAuthorizations#9fab0d1a = Bool;
func (s *Service) AuthResetAuthorizations(ctx context.Context, request *mtproto.TLAuthResetAuthorizations) (*mtproto.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("auth.resetAuthorizations - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.AuthResetAuthorizations(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("auth.resetAuthorizations - reply: %s", r.DebugString())
	return r, err
}

// AuthExportAuthorization
// auth.exportAuthorization#e5bfffcd dc_id:int = auth.ExportedAuthorization;
func (s *Service) AuthExportAuthorization(ctx context.Context, request *mtproto.TLAuthExportAuthorization) (*mtproto.Auth_ExportedAuthorization, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("auth.exportAuthorization - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.AuthExportAuthorization(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("auth.exportAuthorization - reply: %s", r.DebugString())
	return r, err
}

// AuthImportAuthorization
// auth.importAuthorization#a57a7dad id:long bytes:bytes = auth.Authorization;
func (s *Service) AuthImportAuthorization(ctx context.Context, request *mtproto.TLAuthImportAuthorization) (*mtproto.Auth_Authorization, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("auth.importAuthorization - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.AuthImportAuthorization(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("auth.importAuthorization - reply: %s", r.DebugString())
	return r, err
}

// AuthBindTempAuthKey
// auth.bindTempAuthKey#cdd42a05 perm_auth_key_id:long nonce:long expires_at:int encrypted_message:bytes = Bool;
func (s *Service) AuthBindTempAuthKey(ctx context.Context, request *mtproto.TLAuthBindTempAuthKey) (*mtproto.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("auth.bindTempAuthKey - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.AuthBindTempAuthKey(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("auth.bindTempAuthKey - reply: %s", r.DebugString())
	return r, err
}

// AuthImportBotAuthorization
// auth.importBotAuthorization#67a3ff2c flags:int api_id:int api_hash:string bot_auth_token:string = auth.Authorization;
func (s *Service) AuthImportBotAuthorization(ctx context.Context, request *mtproto.TLAuthImportBotAuthorization) (*mtproto.Auth_Authorization, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("auth.importBotAuthorization - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.AuthImportBotAuthorization(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("auth.importBotAuthorization - reply: %s", r.DebugString())
	return r, err
}

// AuthCheckPassword
// auth.checkPassword#d18b4d16 password:InputCheckPasswordSRP = auth.Authorization;
func (s *Service) AuthCheckPassword(ctx context.Context, request *mtproto.TLAuthCheckPassword) (*mtproto.Auth_Authorization, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("auth.checkPassword - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.AuthCheckPassword(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("auth.checkPassword - reply: %s", r.DebugString())
	return r, err
}

// AuthRequestPasswordRecovery
// auth.requestPasswordRecovery#d897bc66 = auth.PasswordRecovery;
func (s *Service) AuthRequestPasswordRecovery(ctx context.Context, request *mtproto.TLAuthRequestPasswordRecovery) (*mtproto.Auth_PasswordRecovery, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("auth.requestPasswordRecovery - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.AuthRequestPasswordRecovery(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("auth.requestPasswordRecovery - reply: %s", r.DebugString())
	return r, err
}

// AuthRecoverPassword
// auth.recoverPassword#37096c70 flags:# code:string new_settings:flags.0?account.PasswordInputSettings = auth.Authorization;
func (s *Service) AuthRecoverPassword(ctx context.Context, request *mtproto.TLAuthRecoverPassword) (*mtproto.Auth_Authorization, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("auth.recoverPassword - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.AuthRecoverPassword(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("auth.recoverPassword - reply: %s", r.DebugString())
	return r, err
}

// AuthResendCode
// auth.resendCode#3ef1a9bf phone_number:string phone_code_hash:string = auth.SentCode;
func (s *Service) AuthResendCode(ctx context.Context, request *mtproto.TLAuthResendCode) (*mtproto.Auth_SentCode, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("auth.resendCode - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.AuthResendCode(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("auth.resendCode - reply: %s", r.DebugString())
	return r, err
}

// AuthCancelCode
// auth.cancelCode#1f040578 phone_number:string phone_code_hash:string = Bool;
func (s *Service) AuthCancelCode(ctx context.Context, request *mtproto.TLAuthCancelCode) (*mtproto.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("auth.cancelCode - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.AuthCancelCode(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("auth.cancelCode - reply: %s", r.DebugString())
	return r, err
}

// AuthDropTempAuthKeys
// auth.dropTempAuthKeys#8e48a188 except_auth_keys:Vector<long> = Bool;
func (s *Service) AuthDropTempAuthKeys(ctx context.Context, request *mtproto.TLAuthDropTempAuthKeys) (*mtproto.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("auth.dropTempAuthKeys - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.AuthDropTempAuthKeys(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("auth.dropTempAuthKeys - reply: %s", r.DebugString())
	return r, err
}

// AuthCheckRecoveryPassword
// auth.checkRecoveryPassword#d36bf79 code:string = Bool;
func (s *Service) AuthCheckRecoveryPassword(ctx context.Context, request *mtproto.TLAuthCheckRecoveryPassword) (*mtproto.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("auth.checkRecoveryPassword - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.AuthCheckRecoveryPassword(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("auth.checkRecoveryPassword - reply: %s", r.DebugString())
	return r, err
}

// AccountResetPassword
// account.resetPassword#9308ce1b = account.ResetPasswordResult;
func (s *Service) AccountResetPassword(ctx context.Context, request *mtproto.TLAccountResetPassword) (*mtproto.Account_ResetPasswordResult, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("account.resetPassword - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.AccountResetPassword(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("account.resetPassword - reply: %s", r.DebugString())
	return r, err
}

// AuthToggleBan
// auth.toggleBan flags:# phone:string predefined:flags.0?true expires:flags.1?int reason:flags.1?string = PredefinedUser;
func (s *Service) AuthToggleBan(ctx context.Context, request *mtproto.TLAuthToggleBan) (*mtproto.PredefinedUser, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("auth.toggleBan - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.AuthToggleBan(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("auth.toggleBan - reply: %s", r.DebugString())
	return r, err
}

// AccountSetAuthorizationTTL
// account.setAuthorizationTTL#bf899aa0 authorization_ttl_days:int = Bool;
func (s *Service) AccountSetAuthorizationTTL(ctx context.Context, request *mtproto.TLAccountSetAuthorizationTTL) (*mtproto.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("account.setAuthorizationTTL - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.AccountSetAuthorizationTTL(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("account.setAuthorizationTTL - reply: %s", r.DebugString())
	return r, err
}

// AccountChangeAuthorizationSettings
// account.changeAuthorizationSettings#40f48462 flags:# hash:long encrypted_requests_disabled:flags.0?Bool call_requests_disabled:flags.1?Bool = Bool;
func (s *Service) AccountChangeAuthorizationSettings(ctx context.Context, request *mtproto.TLAccountChangeAuthorizationSettings) (*mtproto.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("account.changeAuthorizationSettings - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.AccountChangeAuthorizationSettings(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("account.changeAuthorizationSettings - reply: %s", r.DebugString())
	return r, err
}
