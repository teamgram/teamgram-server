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

	"github.com/teamgram/proto/v2/tg"
	"github.com/teamgram/teamgram-server/v2/app/bff/authorization/authorization/authorizationservice"

	"github.com/cloudwego/kitex/client"
)

type AuthorizationClient interface {
	AuthSendCode(ctx context.Context, in *tg.TLAuthSendCode) (*tg.AuthSentCode, error)
	AuthSignUp(ctx context.Context, in *tg.TLAuthSignUp) (*tg.AuthAuthorization, error)
	AuthSignIn(ctx context.Context, in *tg.TLAuthSignIn) (*tg.AuthAuthorization, error)
	AuthLogOut(ctx context.Context, in *tg.TLAuthLogOut) (*tg.AuthLoggedOut, error)
	AuthResetAuthorizations(ctx context.Context, in *tg.TLAuthResetAuthorizations) (*tg.Bool, error)
	AuthExportAuthorization(ctx context.Context, in *tg.TLAuthExportAuthorization) (*tg.AuthExportedAuthorization, error)
	AuthImportAuthorization(ctx context.Context, in *tg.TLAuthImportAuthorization) (*tg.AuthAuthorization, error)
	AuthBindTempAuthKey(ctx context.Context, in *tg.TLAuthBindTempAuthKey) (*tg.Bool, error)
	AuthImportBotAuthorization(ctx context.Context, in *tg.TLAuthImportBotAuthorization) (*tg.AuthAuthorization, error)
	AuthCheckPassword(ctx context.Context, in *tg.TLAuthCheckPassword) (*tg.AuthAuthorization, error)
	AuthRequestPasswordRecovery(ctx context.Context, in *tg.TLAuthRequestPasswordRecovery) (*tg.AuthPasswordRecovery, error)
	AuthRecoverPassword(ctx context.Context, in *tg.TLAuthRecoverPassword) (*tg.AuthAuthorization, error)
	AuthResendCode(ctx context.Context, in *tg.TLAuthResendCode) (*tg.AuthSentCode, error)
	AuthCancelCode(ctx context.Context, in *tg.TLAuthCancelCode) (*tg.Bool, error)
	AuthDropTempAuthKeys(ctx context.Context, in *tg.TLAuthDropTempAuthKeys) (*tg.Bool, error)
	AuthCheckRecoveryPassword(ctx context.Context, in *tg.TLAuthCheckRecoveryPassword) (*tg.Bool, error)
	AuthImportWebTokenAuthorization(ctx context.Context, in *tg.TLAuthImportWebTokenAuthorization) (*tg.AuthAuthorization, error)
	AuthRequestFirebaseSms(ctx context.Context, in *tg.TLAuthRequestFirebaseSms) (*tg.Bool, error)
	AuthResetLoginEmail(ctx context.Context, in *tg.TLAuthResetLoginEmail) (*tg.AuthSentCode, error)
	AuthReportMissingCode(ctx context.Context, in *tg.TLAuthReportMissingCode) (*tg.Bool, error)
	AccountSendVerifyEmailCode(ctx context.Context, in *tg.TLAccountSendVerifyEmailCode) (*tg.AccountSentEmailCode, error)
	AccountVerifyEmail(ctx context.Context, in *tg.TLAccountVerifyEmail) (*tg.AccountEmailVerified, error)
	AccountResetPassword(ctx context.Context, in *tg.TLAccountResetPassword) (*tg.AccountResetPasswordResult, error)
	AccountSetAuthorizationTTL(ctx context.Context, in *tg.TLAccountSetAuthorizationTTL) (*tg.Bool, error)
	AccountChangeAuthorizationSettings(ctx context.Context, in *tg.TLAccountChangeAuthorizationSettings) (*tg.Bool, error)
	AccountInvalidateSignInCodes(ctx context.Context, in *tg.TLAccountInvalidateSignInCodes) (*tg.Bool, error)
	AuthToggleBan(ctx context.Context, in *tg.TLAuthToggleBan) (*tg.PredefinedUser, error)
}

type defaultAuthorizationClient struct {
	cli client.Client
}

func NewAuthorizationClient(cli client.Client) AuthorizationClient {
	return &defaultAuthorizationClient{
		cli: cli,
	}
}

// AuthSendCode
// auth.sendCode#a677244f phone_number:string api_id:int api_hash:string settings:CodeSettings = auth.SentCode;
func (m *defaultAuthorizationClient) AuthSendCode(ctx context.Context, in *tg.TLAuthSendCode) (*tg.AuthSentCode, error) {
	cli := authorizationservice.NewRPCAuthorizationClient(m.cli)
	return cli.AuthSendCode(ctx, in)
}

// AuthSignUp
// auth.signUp#aac7b717 flags:# no_joined_notifications:flags.0?true phone_number:string phone_code_hash:string first_name:string last_name:string = auth.Authorization;
func (m *defaultAuthorizationClient) AuthSignUp(ctx context.Context, in *tg.TLAuthSignUp) (*tg.AuthAuthorization, error) {
	cli := authorizationservice.NewRPCAuthorizationClient(m.cli)
	return cli.AuthSignUp(ctx, in)
}

// AuthSignIn
// auth.signIn#8d52a951 flags:# phone_number:string phone_code_hash:string phone_code:flags.0?string email_verification:flags.1?EmailVerification = auth.Authorization;
func (m *defaultAuthorizationClient) AuthSignIn(ctx context.Context, in *tg.TLAuthSignIn) (*tg.AuthAuthorization, error) {
	cli := authorizationservice.NewRPCAuthorizationClient(m.cli)
	return cli.AuthSignIn(ctx, in)
}

// AuthLogOut
// auth.logOut#3e72ba19 = auth.LoggedOut;
func (m *defaultAuthorizationClient) AuthLogOut(ctx context.Context, in *tg.TLAuthLogOut) (*tg.AuthLoggedOut, error) {
	cli := authorizationservice.NewRPCAuthorizationClient(m.cli)
	return cli.AuthLogOut(ctx, in)
}

// AuthResetAuthorizations
// auth.resetAuthorizations#9fab0d1a = Bool;
func (m *defaultAuthorizationClient) AuthResetAuthorizations(ctx context.Context, in *tg.TLAuthResetAuthorizations) (*tg.Bool, error) {
	cli := authorizationservice.NewRPCAuthorizationClient(m.cli)
	return cli.AuthResetAuthorizations(ctx, in)
}

// AuthExportAuthorization
// auth.exportAuthorization#e5bfffcd dc_id:int = auth.ExportedAuthorization;
func (m *defaultAuthorizationClient) AuthExportAuthorization(ctx context.Context, in *tg.TLAuthExportAuthorization) (*tg.AuthExportedAuthorization, error) {
	cli := authorizationservice.NewRPCAuthorizationClient(m.cli)
	return cli.AuthExportAuthorization(ctx, in)
}

// AuthImportAuthorization
// auth.importAuthorization#a57a7dad id:long bytes:bytes = auth.Authorization;
func (m *defaultAuthorizationClient) AuthImportAuthorization(ctx context.Context, in *tg.TLAuthImportAuthorization) (*tg.AuthAuthorization, error) {
	cli := authorizationservice.NewRPCAuthorizationClient(m.cli)
	return cli.AuthImportAuthorization(ctx, in)
}

// AuthBindTempAuthKey
// auth.bindTempAuthKey#cdd42a05 perm_auth_key_id:long nonce:long expires_at:int encrypted_message:bytes = Bool;
func (m *defaultAuthorizationClient) AuthBindTempAuthKey(ctx context.Context, in *tg.TLAuthBindTempAuthKey) (*tg.Bool, error) {
	cli := authorizationservice.NewRPCAuthorizationClient(m.cli)
	return cli.AuthBindTempAuthKey(ctx, in)
}

// AuthImportBotAuthorization
// auth.importBotAuthorization#67a3ff2c flags:int api_id:int api_hash:string bot_auth_token:string = auth.Authorization;
func (m *defaultAuthorizationClient) AuthImportBotAuthorization(ctx context.Context, in *tg.TLAuthImportBotAuthorization) (*tg.AuthAuthorization, error) {
	cli := authorizationservice.NewRPCAuthorizationClient(m.cli)
	return cli.AuthImportBotAuthorization(ctx, in)
}

// AuthCheckPassword
// auth.checkPassword#d18b4d16 password:InputCheckPasswordSRP = auth.Authorization;
func (m *defaultAuthorizationClient) AuthCheckPassword(ctx context.Context, in *tg.TLAuthCheckPassword) (*tg.AuthAuthorization, error) {
	cli := authorizationservice.NewRPCAuthorizationClient(m.cli)
	return cli.AuthCheckPassword(ctx, in)
}

// AuthRequestPasswordRecovery
// auth.requestPasswordRecovery#d897bc66 = auth.PasswordRecovery;
func (m *defaultAuthorizationClient) AuthRequestPasswordRecovery(ctx context.Context, in *tg.TLAuthRequestPasswordRecovery) (*tg.AuthPasswordRecovery, error) {
	cli := authorizationservice.NewRPCAuthorizationClient(m.cli)
	return cli.AuthRequestPasswordRecovery(ctx, in)
}

// AuthRecoverPassword
// auth.recoverPassword#37096c70 flags:# code:string new_settings:flags.0?account.PasswordInputSettings = auth.Authorization;
func (m *defaultAuthorizationClient) AuthRecoverPassword(ctx context.Context, in *tg.TLAuthRecoverPassword) (*tg.AuthAuthorization, error) {
	cli := authorizationservice.NewRPCAuthorizationClient(m.cli)
	return cli.AuthRecoverPassword(ctx, in)
}

// AuthResendCode
// auth.resendCode#cae47523 flags:# phone_number:string phone_code_hash:string reason:flags.0?string = auth.SentCode;
func (m *defaultAuthorizationClient) AuthResendCode(ctx context.Context, in *tg.TLAuthResendCode) (*tg.AuthSentCode, error) {
	cli := authorizationservice.NewRPCAuthorizationClient(m.cli)
	return cli.AuthResendCode(ctx, in)
}

// AuthCancelCode
// auth.cancelCode#1f040578 phone_number:string phone_code_hash:string = Bool;
func (m *defaultAuthorizationClient) AuthCancelCode(ctx context.Context, in *tg.TLAuthCancelCode) (*tg.Bool, error) {
	cli := authorizationservice.NewRPCAuthorizationClient(m.cli)
	return cli.AuthCancelCode(ctx, in)
}

// AuthDropTempAuthKeys
// auth.dropTempAuthKeys#8e48a188 except_auth_keys:Vector<long> = Bool;
func (m *defaultAuthorizationClient) AuthDropTempAuthKeys(ctx context.Context, in *tg.TLAuthDropTempAuthKeys) (*tg.Bool, error) {
	cli := authorizationservice.NewRPCAuthorizationClient(m.cli)
	return cli.AuthDropTempAuthKeys(ctx, in)
}

// AuthCheckRecoveryPassword
// auth.checkRecoveryPassword#d36bf79 code:string = Bool;
func (m *defaultAuthorizationClient) AuthCheckRecoveryPassword(ctx context.Context, in *tg.TLAuthCheckRecoveryPassword) (*tg.Bool, error) {
	cli := authorizationservice.NewRPCAuthorizationClient(m.cli)
	return cli.AuthCheckRecoveryPassword(ctx, in)
}

// AuthImportWebTokenAuthorization
// auth.importWebTokenAuthorization#2db873a9 api_id:int api_hash:string web_auth_token:string = auth.Authorization;
func (m *defaultAuthorizationClient) AuthImportWebTokenAuthorization(ctx context.Context, in *tg.TLAuthImportWebTokenAuthorization) (*tg.AuthAuthorization, error) {
	cli := authorizationservice.NewRPCAuthorizationClient(m.cli)
	return cli.AuthImportWebTokenAuthorization(ctx, in)
}

// AuthRequestFirebaseSms
// auth.requestFirebaseSms#8e39261e flags:# phone_number:string phone_code_hash:string safety_net_token:flags.0?string play_integrity_token:flags.2?string ios_push_secret:flags.1?string = Bool;
func (m *defaultAuthorizationClient) AuthRequestFirebaseSms(ctx context.Context, in *tg.TLAuthRequestFirebaseSms) (*tg.Bool, error) {
	cli := authorizationservice.NewRPCAuthorizationClient(m.cli)
	return cli.AuthRequestFirebaseSms(ctx, in)
}

// AuthResetLoginEmail
// auth.resetLoginEmail#7e960193 phone_number:string phone_code_hash:string = auth.SentCode;
func (m *defaultAuthorizationClient) AuthResetLoginEmail(ctx context.Context, in *tg.TLAuthResetLoginEmail) (*tg.AuthSentCode, error) {
	cli := authorizationservice.NewRPCAuthorizationClient(m.cli)
	return cli.AuthResetLoginEmail(ctx, in)
}

// AuthReportMissingCode
// auth.reportMissingCode#cb9deff6 phone_number:string phone_code_hash:string mnc:string = Bool;
func (m *defaultAuthorizationClient) AuthReportMissingCode(ctx context.Context, in *tg.TLAuthReportMissingCode) (*tg.Bool, error) {
	cli := authorizationservice.NewRPCAuthorizationClient(m.cli)
	return cli.AuthReportMissingCode(ctx, in)
}

// AccountSendVerifyEmailCode
// account.sendVerifyEmailCode#98e037bb purpose:EmailVerifyPurpose email:string = account.SentEmailCode;
func (m *defaultAuthorizationClient) AccountSendVerifyEmailCode(ctx context.Context, in *tg.TLAccountSendVerifyEmailCode) (*tg.AccountSentEmailCode, error) {
	cli := authorizationservice.NewRPCAuthorizationClient(m.cli)
	return cli.AccountSendVerifyEmailCode(ctx, in)
}

// AccountVerifyEmail
// account.verifyEmail#32da4cf purpose:EmailVerifyPurpose verification:EmailVerification = account.EmailVerified;
func (m *defaultAuthorizationClient) AccountVerifyEmail(ctx context.Context, in *tg.TLAccountVerifyEmail) (*tg.AccountEmailVerified, error) {
	cli := authorizationservice.NewRPCAuthorizationClient(m.cli)
	return cli.AccountVerifyEmail(ctx, in)
}

// AccountResetPassword
// account.resetPassword#9308ce1b = account.ResetPasswordResult;
func (m *defaultAuthorizationClient) AccountResetPassword(ctx context.Context, in *tg.TLAccountResetPassword) (*tg.AccountResetPasswordResult, error) {
	cli := authorizationservice.NewRPCAuthorizationClient(m.cli)
	return cli.AccountResetPassword(ctx, in)
}

// AccountSetAuthorizationTTL
// account.setAuthorizationTTL#bf899aa0 authorization_ttl_days:int = Bool;
func (m *defaultAuthorizationClient) AccountSetAuthorizationTTL(ctx context.Context, in *tg.TLAccountSetAuthorizationTTL) (*tg.Bool, error) {
	cli := authorizationservice.NewRPCAuthorizationClient(m.cli)
	return cli.AccountSetAuthorizationTTL(ctx, in)
}

// AccountChangeAuthorizationSettings
// account.changeAuthorizationSettings#40f48462 flags:# confirmed:flags.3?true hash:long encrypted_requests_disabled:flags.0?Bool call_requests_disabled:flags.1?Bool = Bool;
func (m *defaultAuthorizationClient) AccountChangeAuthorizationSettings(ctx context.Context, in *tg.TLAccountChangeAuthorizationSettings) (*tg.Bool, error) {
	cli := authorizationservice.NewRPCAuthorizationClient(m.cli)
	return cli.AccountChangeAuthorizationSettings(ctx, in)
}

// AccountInvalidateSignInCodes
// account.invalidateSignInCodes#ca8ae8ba codes:Vector<string> = Bool;
func (m *defaultAuthorizationClient) AccountInvalidateSignInCodes(ctx context.Context, in *tg.TLAccountInvalidateSignInCodes) (*tg.Bool, error) {
	cli := authorizationservice.NewRPCAuthorizationClient(m.cli)
	return cli.AccountInvalidateSignInCodes(ctx, in)
}

// AuthToggleBan
// auth.toggleBan flags:# phone:string predefined:flags.0?true expires:flags.1?int reason:flags.1?string = PredefinedUser;
func (m *defaultAuthorizationClient) AuthToggleBan(ctx context.Context, in *tg.TLAuthToggleBan) (*tg.PredefinedUser, error) {
	cli := authorizationservice.NewRPCAuthorizationClient(m.cli)
	return cli.AuthToggleBan(ctx, in)
}
