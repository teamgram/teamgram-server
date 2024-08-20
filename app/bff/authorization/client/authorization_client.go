/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright 2024 Teamgram Authors.
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package authorizationclient

import (
	"context"

	"github.com/teamgram/proto/mtproto"

	"github.com/zeromicro/go-zero/zrpc"
)

var _ *mtproto.Bool

type AuthorizationClient interface {
	AuthSendCode(ctx context.Context, in *mtproto.TLAuthSendCode) (*mtproto.Auth_SentCode, error)
	AuthSignUp(ctx context.Context, in *mtproto.TLAuthSignUp) (*mtproto.Auth_Authorization, error)
	AuthSignIn(ctx context.Context, in *mtproto.TLAuthSignIn) (*mtproto.Auth_Authorization, error)
	AuthLogOut(ctx context.Context, in *mtproto.TLAuthLogOut) (*mtproto.Auth_LoggedOut, error)
	AuthResetAuthorizations(ctx context.Context, in *mtproto.TLAuthResetAuthorizations) (*mtproto.Bool, error)
	AuthExportAuthorization(ctx context.Context, in *mtproto.TLAuthExportAuthorization) (*mtproto.Auth_ExportedAuthorization, error)
	AuthImportAuthorization(ctx context.Context, in *mtproto.TLAuthImportAuthorization) (*mtproto.Auth_Authorization, error)
	AuthBindTempAuthKey(ctx context.Context, in *mtproto.TLAuthBindTempAuthKey) (*mtproto.Bool, error)
	AuthImportBotAuthorization(ctx context.Context, in *mtproto.TLAuthImportBotAuthorization) (*mtproto.Auth_Authorization, error)
	AuthCheckPassword(ctx context.Context, in *mtproto.TLAuthCheckPassword) (*mtproto.Auth_Authorization, error)
	AuthRequestPasswordRecovery(ctx context.Context, in *mtproto.TLAuthRequestPasswordRecovery) (*mtproto.Auth_PasswordRecovery, error)
	AuthRecoverPassword(ctx context.Context, in *mtproto.TLAuthRecoverPassword) (*mtproto.Auth_Authorization, error)
	AuthResendCode(ctx context.Context, in *mtproto.TLAuthResendCode) (*mtproto.Auth_SentCode, error)
	AuthCancelCode(ctx context.Context, in *mtproto.TLAuthCancelCode) (*mtproto.Bool, error)
	AuthDropTempAuthKeys(ctx context.Context, in *mtproto.TLAuthDropTempAuthKeys) (*mtproto.Bool, error)
	AuthCheckRecoveryPassword(ctx context.Context, in *mtproto.TLAuthCheckRecoveryPassword) (*mtproto.Bool, error)
	AuthImportWebTokenAuthorization(ctx context.Context, in *mtproto.TLAuthImportWebTokenAuthorization) (*mtproto.Auth_Authorization, error)
	AuthRequestFirebaseSms(ctx context.Context, in *mtproto.TLAuthRequestFirebaseSms) (*mtproto.Bool, error)
	AuthResetLoginEmail(ctx context.Context, in *mtproto.TLAuthResetLoginEmail) (*mtproto.Auth_SentCode, error)
	AuthReportMissingCode(ctx context.Context, in *mtproto.TLAuthReportMissingCode) (*mtproto.Bool, error)
	AccountSendVerifyEmailCode(ctx context.Context, in *mtproto.TLAccountSendVerifyEmailCode) (*mtproto.Account_SentEmailCode, error)
	AccountVerifyEmail32DA4CF(ctx context.Context, in *mtproto.TLAccountVerifyEmail32DA4CF) (*mtproto.Account_EmailVerified, error)
	AccountResetPassword(ctx context.Context, in *mtproto.TLAccountResetPassword) (*mtproto.Account_ResetPasswordResult, error)
	AccountSetAuthorizationTTL(ctx context.Context, in *mtproto.TLAccountSetAuthorizationTTL) (*mtproto.Bool, error)
	AccountChangeAuthorizationSettings(ctx context.Context, in *mtproto.TLAccountChangeAuthorizationSettings) (*mtproto.Bool, error)
	AccountInvalidateSignInCodes(ctx context.Context, in *mtproto.TLAccountInvalidateSignInCodes) (*mtproto.Bool, error)
	AccountVerifyEmailECBA39DB(ctx context.Context, in *mtproto.TLAccountVerifyEmailECBA39DB) (*mtproto.Bool, error)
	AuthToggleBan(ctx context.Context, in *mtproto.TLAuthToggleBan) (*mtproto.PredefinedUser, error)
}

type defaultAuthorizationClient struct {
	cli zrpc.Client
}

func NewAuthorizationClient(cli zrpc.Client) AuthorizationClient {
	return &defaultAuthorizationClient{
		cli: cli,
	}
}

// AuthSendCode
// auth.sendCode#a677244f phone_number:string api_id:int api_hash:string settings:CodeSettings = auth.SentCode;
func (m *defaultAuthorizationClient) AuthSendCode(ctx context.Context, in *mtproto.TLAuthSendCode) (*mtproto.Auth_SentCode, error) {
	client := mtproto.NewRPCAuthorizationClient(m.cli.Conn())
	return client.AuthSendCode(ctx, in)
}

// AuthSignUp
// auth.signUp#aac7b717 flags:# no_joined_notifications:flags.0?true phone_number:string phone_code_hash:string first_name:string last_name:string = auth.Authorization;
func (m *defaultAuthorizationClient) AuthSignUp(ctx context.Context, in *mtproto.TLAuthSignUp) (*mtproto.Auth_Authorization, error) {
	client := mtproto.NewRPCAuthorizationClient(m.cli.Conn())
	return client.AuthSignUp(ctx, in)
}

// AuthSignIn
// auth.signIn#8d52a951 flags:# phone_number:string phone_code_hash:string phone_code:flags.0?string email_verification:flags.1?EmailVerification = auth.Authorization;
func (m *defaultAuthorizationClient) AuthSignIn(ctx context.Context, in *mtproto.TLAuthSignIn) (*mtproto.Auth_Authorization, error) {
	client := mtproto.NewRPCAuthorizationClient(m.cli.Conn())
	return client.AuthSignIn(ctx, in)
}

// AuthLogOut
// auth.logOut#3e72ba19 = auth.LoggedOut;
func (m *defaultAuthorizationClient) AuthLogOut(ctx context.Context, in *mtproto.TLAuthLogOut) (*mtproto.Auth_LoggedOut, error) {
	client := mtproto.NewRPCAuthorizationClient(m.cli.Conn())
	return client.AuthLogOut(ctx, in)
}

// AuthResetAuthorizations
// auth.resetAuthorizations#9fab0d1a = Bool;
func (m *defaultAuthorizationClient) AuthResetAuthorizations(ctx context.Context, in *mtproto.TLAuthResetAuthorizations) (*mtproto.Bool, error) {
	client := mtproto.NewRPCAuthorizationClient(m.cli.Conn())
	return client.AuthResetAuthorizations(ctx, in)
}

// AuthExportAuthorization
// auth.exportAuthorization#e5bfffcd dc_id:int = auth.ExportedAuthorization;
func (m *defaultAuthorizationClient) AuthExportAuthorization(ctx context.Context, in *mtproto.TLAuthExportAuthorization) (*mtproto.Auth_ExportedAuthorization, error) {
	client := mtproto.NewRPCAuthorizationClient(m.cli.Conn())
	return client.AuthExportAuthorization(ctx, in)
}

// AuthImportAuthorization
// auth.importAuthorization#a57a7dad id:long bytes:bytes = auth.Authorization;
func (m *defaultAuthorizationClient) AuthImportAuthorization(ctx context.Context, in *mtproto.TLAuthImportAuthorization) (*mtproto.Auth_Authorization, error) {
	client := mtproto.NewRPCAuthorizationClient(m.cli.Conn())
	return client.AuthImportAuthorization(ctx, in)
}

// AuthBindTempAuthKey
// auth.bindTempAuthKey#cdd42a05 perm_auth_key_id:long nonce:long expires_at:int encrypted_message:bytes = Bool;
func (m *defaultAuthorizationClient) AuthBindTempAuthKey(ctx context.Context, in *mtproto.TLAuthBindTempAuthKey) (*mtproto.Bool, error) {
	client := mtproto.NewRPCAuthorizationClient(m.cli.Conn())
	return client.AuthBindTempAuthKey(ctx, in)
}

// AuthImportBotAuthorization
// auth.importBotAuthorization#67a3ff2c flags:int api_id:int api_hash:string bot_auth_token:string = auth.Authorization;
func (m *defaultAuthorizationClient) AuthImportBotAuthorization(ctx context.Context, in *mtproto.TLAuthImportBotAuthorization) (*mtproto.Auth_Authorization, error) {
	client := mtproto.NewRPCAuthorizationClient(m.cli.Conn())
	return client.AuthImportBotAuthorization(ctx, in)
}

// AuthCheckPassword
// auth.checkPassword#d18b4d16 password:InputCheckPasswordSRP = auth.Authorization;
func (m *defaultAuthorizationClient) AuthCheckPassword(ctx context.Context, in *mtproto.TLAuthCheckPassword) (*mtproto.Auth_Authorization, error) {
	client := mtproto.NewRPCAuthorizationClient(m.cli.Conn())
	return client.AuthCheckPassword(ctx, in)
}

// AuthRequestPasswordRecovery
// auth.requestPasswordRecovery#d897bc66 = auth.PasswordRecovery;
func (m *defaultAuthorizationClient) AuthRequestPasswordRecovery(ctx context.Context, in *mtproto.TLAuthRequestPasswordRecovery) (*mtproto.Auth_PasswordRecovery, error) {
	client := mtproto.NewRPCAuthorizationClient(m.cli.Conn())
	return client.AuthRequestPasswordRecovery(ctx, in)
}

// AuthRecoverPassword
// auth.recoverPassword#37096c70 flags:# code:string new_settings:flags.0?account.PasswordInputSettings = auth.Authorization;
func (m *defaultAuthorizationClient) AuthRecoverPassword(ctx context.Context, in *mtproto.TLAuthRecoverPassword) (*mtproto.Auth_Authorization, error) {
	client := mtproto.NewRPCAuthorizationClient(m.cli.Conn())
	return client.AuthRecoverPassword(ctx, in)
}

// AuthResendCode
// auth.resendCode#cae47523 flags:# phone_number:string phone_code_hash:string reason:flags.0?string = auth.SentCode;
func (m *defaultAuthorizationClient) AuthResendCode(ctx context.Context, in *mtproto.TLAuthResendCode) (*mtproto.Auth_SentCode, error) {
	client := mtproto.NewRPCAuthorizationClient(m.cli.Conn())
	return client.AuthResendCode(ctx, in)
}

// AuthCancelCode
// auth.cancelCode#1f040578 phone_number:string phone_code_hash:string = Bool;
func (m *defaultAuthorizationClient) AuthCancelCode(ctx context.Context, in *mtproto.TLAuthCancelCode) (*mtproto.Bool, error) {
	client := mtproto.NewRPCAuthorizationClient(m.cli.Conn())
	return client.AuthCancelCode(ctx, in)
}

// AuthDropTempAuthKeys
// auth.dropTempAuthKeys#8e48a188 except_auth_keys:Vector<long> = Bool;
func (m *defaultAuthorizationClient) AuthDropTempAuthKeys(ctx context.Context, in *mtproto.TLAuthDropTempAuthKeys) (*mtproto.Bool, error) {
	client := mtproto.NewRPCAuthorizationClient(m.cli.Conn())
	return client.AuthDropTempAuthKeys(ctx, in)
}

// AuthCheckRecoveryPassword
// auth.checkRecoveryPassword#d36bf79 code:string = Bool;
func (m *defaultAuthorizationClient) AuthCheckRecoveryPassword(ctx context.Context, in *mtproto.TLAuthCheckRecoveryPassword) (*mtproto.Bool, error) {
	client := mtproto.NewRPCAuthorizationClient(m.cli.Conn())
	return client.AuthCheckRecoveryPassword(ctx, in)
}

// AuthImportWebTokenAuthorization
// auth.importWebTokenAuthorization#2db873a9 api_id:int api_hash:string web_auth_token:string = auth.Authorization;
func (m *defaultAuthorizationClient) AuthImportWebTokenAuthorization(ctx context.Context, in *mtproto.TLAuthImportWebTokenAuthorization) (*mtproto.Auth_Authorization, error) {
	client := mtproto.NewRPCAuthorizationClient(m.cli.Conn())
	return client.AuthImportWebTokenAuthorization(ctx, in)
}

// AuthRequestFirebaseSms
// auth.requestFirebaseSms#8e39261e flags:# phone_number:string phone_code_hash:string safety_net_token:flags.0?string play_integrity_token:flags.2?string ios_push_secret:flags.1?string = Bool;
func (m *defaultAuthorizationClient) AuthRequestFirebaseSms(ctx context.Context, in *mtproto.TLAuthRequestFirebaseSms) (*mtproto.Bool, error) {
	client := mtproto.NewRPCAuthorizationClient(m.cli.Conn())
	return client.AuthRequestFirebaseSms(ctx, in)
}

// AuthResetLoginEmail
// auth.resetLoginEmail#7e960193 phone_number:string phone_code_hash:string = auth.SentCode;
func (m *defaultAuthorizationClient) AuthResetLoginEmail(ctx context.Context, in *mtproto.TLAuthResetLoginEmail) (*mtproto.Auth_SentCode, error) {
	client := mtproto.NewRPCAuthorizationClient(m.cli.Conn())
	return client.AuthResetLoginEmail(ctx, in)
}

// AuthReportMissingCode
// auth.reportMissingCode#cb9deff6 phone_number:string phone_code_hash:string mnc:string = Bool;
func (m *defaultAuthorizationClient) AuthReportMissingCode(ctx context.Context, in *mtproto.TLAuthReportMissingCode) (*mtproto.Bool, error) {
	client := mtproto.NewRPCAuthorizationClient(m.cli.Conn())
	return client.AuthReportMissingCode(ctx, in)
}

// AccountSendVerifyEmailCode
// account.sendVerifyEmailCode#98e037bb purpose:EmailVerifyPurpose email:string = account.SentEmailCode;
func (m *defaultAuthorizationClient) AccountSendVerifyEmailCode(ctx context.Context, in *mtproto.TLAccountSendVerifyEmailCode) (*mtproto.Account_SentEmailCode, error) {
	client := mtproto.NewRPCAuthorizationClient(m.cli.Conn())
	return client.AccountSendVerifyEmailCode(ctx, in)
}

// AccountVerifyEmail32DA4CF
// account.verifyEmail#32da4cf purpose:EmailVerifyPurpose verification:EmailVerification = account.EmailVerified;
func (m *defaultAuthorizationClient) AccountVerifyEmail32DA4CF(ctx context.Context, in *mtproto.TLAccountVerifyEmail32DA4CF) (*mtproto.Account_EmailVerified, error) {
	client := mtproto.NewRPCAuthorizationClient(m.cli.Conn())
	return client.AccountVerifyEmail32DA4CF(ctx, in)
}

// AccountResetPassword
// account.resetPassword#9308ce1b = account.ResetPasswordResult;
func (m *defaultAuthorizationClient) AccountResetPassword(ctx context.Context, in *mtproto.TLAccountResetPassword) (*mtproto.Account_ResetPasswordResult, error) {
	client := mtproto.NewRPCAuthorizationClient(m.cli.Conn())
	return client.AccountResetPassword(ctx, in)
}

// AccountSetAuthorizationTTL
// account.setAuthorizationTTL#bf899aa0 authorization_ttl_days:int = Bool;
func (m *defaultAuthorizationClient) AccountSetAuthorizationTTL(ctx context.Context, in *mtproto.TLAccountSetAuthorizationTTL) (*mtproto.Bool, error) {
	client := mtproto.NewRPCAuthorizationClient(m.cli.Conn())
	return client.AccountSetAuthorizationTTL(ctx, in)
}

// AccountChangeAuthorizationSettings
// account.changeAuthorizationSettings#40f48462 flags:# confirmed:flags.3?true hash:long encrypted_requests_disabled:flags.0?Bool call_requests_disabled:flags.1?Bool = Bool;
func (m *defaultAuthorizationClient) AccountChangeAuthorizationSettings(ctx context.Context, in *mtproto.TLAccountChangeAuthorizationSettings) (*mtproto.Bool, error) {
	client := mtproto.NewRPCAuthorizationClient(m.cli.Conn())
	return client.AccountChangeAuthorizationSettings(ctx, in)
}

// AccountInvalidateSignInCodes
// account.invalidateSignInCodes#ca8ae8ba codes:Vector<string> = Bool;
func (m *defaultAuthorizationClient) AccountInvalidateSignInCodes(ctx context.Context, in *mtproto.TLAccountInvalidateSignInCodes) (*mtproto.Bool, error) {
	client := mtproto.NewRPCAuthorizationClient(m.cli.Conn())
	return client.AccountInvalidateSignInCodes(ctx, in)
}

// AccountVerifyEmailECBA39DB
// account.verifyEmail#ecba39db email:string code:string = Bool;
func (m *defaultAuthorizationClient) AccountVerifyEmailECBA39DB(ctx context.Context, in *mtproto.TLAccountVerifyEmailECBA39DB) (*mtproto.Bool, error) {
	client := mtproto.NewRPCAuthorizationClient(m.cli.Conn())
	return client.AccountVerifyEmailECBA39DB(ctx, in)
}

// AuthToggleBan
// auth.toggleBan flags:# phone:string predefined:flags.0?true expires:flags.1?int reason:flags.1?string = PredefinedUser;
func (m *defaultAuthorizationClient) AuthToggleBan(ctx context.Context, in *mtproto.TLAuthToggleBan) (*mtproto.PredefinedUser, error) {
	client := mtproto.NewRPCAuthorizationClient(m.cli.Conn())
	return client.AuthToggleBan(ctx, in)
}
