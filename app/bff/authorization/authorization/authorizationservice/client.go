/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2024-present,  Teamgram Authors.
 *  All rights reserved.
 *
 * Author: Benqi (wubenqi@gmail.com)
 */

package authorizationservice

import (
	"context"

	"github.com/teamgram/proto/v2/tg"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/callopt"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	AuthSendCode(ctx context.Context, req *tg.TLAuthSendCode, callOptions ...callopt.Option) (r *tg.AuthSentCode, err error)
	AuthSignUp(ctx context.Context, req *tg.TLAuthSignUp, callOptions ...callopt.Option) (r *tg.AuthAuthorization, err error)
	AuthSignIn(ctx context.Context, req *tg.TLAuthSignIn, callOptions ...callopt.Option) (r *tg.AuthAuthorization, err error)
	AuthLogOut(ctx context.Context, req *tg.TLAuthLogOut, callOptions ...callopt.Option) (r *tg.AuthLoggedOut, err error)
	AuthResetAuthorizations(ctx context.Context, req *tg.TLAuthResetAuthorizations, callOptions ...callopt.Option) (r *tg.Bool, err error)
	AuthExportAuthorization(ctx context.Context, req *tg.TLAuthExportAuthorization, callOptions ...callopt.Option) (r *tg.AuthExportedAuthorization, err error)
	AuthImportAuthorization(ctx context.Context, req *tg.TLAuthImportAuthorization, callOptions ...callopt.Option) (r *tg.AuthAuthorization, err error)
	AuthBindTempAuthKey(ctx context.Context, req *tg.TLAuthBindTempAuthKey, callOptions ...callopt.Option) (r *tg.Bool, err error)
	AuthImportBotAuthorization(ctx context.Context, req *tg.TLAuthImportBotAuthorization, callOptions ...callopt.Option) (r *tg.AuthAuthorization, err error)
	AuthCheckPassword(ctx context.Context, req *tg.TLAuthCheckPassword, callOptions ...callopt.Option) (r *tg.AuthAuthorization, err error)
	AuthRequestPasswordRecovery(ctx context.Context, req *tg.TLAuthRequestPasswordRecovery, callOptions ...callopt.Option) (r *tg.AuthPasswordRecovery, err error)
	AuthRecoverPassword(ctx context.Context, req *tg.TLAuthRecoverPassword, callOptions ...callopt.Option) (r *tg.AuthAuthorization, err error)
	AuthResendCode(ctx context.Context, req *tg.TLAuthResendCode, callOptions ...callopt.Option) (r *tg.AuthSentCode, err error)
	AuthCancelCode(ctx context.Context, req *tg.TLAuthCancelCode, callOptions ...callopt.Option) (r *tg.Bool, err error)
	AuthDropTempAuthKeys(ctx context.Context, req *tg.TLAuthDropTempAuthKeys, callOptions ...callopt.Option) (r *tg.Bool, err error)
	AuthCheckRecoveryPassword(ctx context.Context, req *tg.TLAuthCheckRecoveryPassword, callOptions ...callopt.Option) (r *tg.Bool, err error)
	AuthImportWebTokenAuthorization(ctx context.Context, req *tg.TLAuthImportWebTokenAuthorization, callOptions ...callopt.Option) (r *tg.AuthAuthorization, err error)
	AuthRequestFirebaseSms(ctx context.Context, req *tg.TLAuthRequestFirebaseSms, callOptions ...callopt.Option) (r *tg.Bool, err error)
	AuthResetLoginEmail(ctx context.Context, req *tg.TLAuthResetLoginEmail, callOptions ...callopt.Option) (r *tg.AuthSentCode, err error)
	AuthReportMissingCode(ctx context.Context, req *tg.TLAuthReportMissingCode, callOptions ...callopt.Option) (r *tg.Bool, err error)
	AccountSendVerifyEmailCode(ctx context.Context, req *tg.TLAccountSendVerifyEmailCode, callOptions ...callopt.Option) (r *tg.AccountSentEmailCode, err error)
	AccountVerifyEmail(ctx context.Context, req *tg.TLAccountVerifyEmail, callOptions ...callopt.Option) (r *tg.AccountEmailVerified, err error)
	AccountResetPassword(ctx context.Context, req *tg.TLAccountResetPassword, callOptions ...callopt.Option) (r *tg.AccountResetPasswordResult, err error)
	AccountSetAuthorizationTTL(ctx context.Context, req *tg.TLAccountSetAuthorizationTTL, callOptions ...callopt.Option) (r *tg.Bool, err error)
	AccountChangeAuthorizationSettings(ctx context.Context, req *tg.TLAccountChangeAuthorizationSettings, callOptions ...callopt.Option) (r *tg.Bool, err error)
	AccountInvalidateSignInCodes(ctx context.Context, req *tg.TLAccountInvalidateSignInCodes, callOptions ...callopt.Option) (r *tg.Bool, err error)
	AuthToggleBan(ctx context.Context, req *tg.TLAuthToggleBan, callOptions ...callopt.Option) (r *tg.PredefinedUser, err error)
}

// NewClient creates a client for the service defined in IDL.
func NewClient(destService string, opts ...client.Option) (Client, error) {
	var options []client.Option
	options = append(options, client.WithDestService(destService))

	options = append(options, opts...)

	kc, err := client.NewClient(serviceInfoForClient(), options...)
	if err != nil {
		return nil, err
	}
	return &kAuthorizationClient{
		kClient: newServiceClient(kc),
	}, nil
}

// MustNewClient creates a client for the service defined in IDL. It panics if any error occurs.
func MustNewClient(destService string, opts ...client.Option) Client {
	kc, err := NewClient(destService, opts...)
	if err != nil {
		panic(err)
	}
	return kc
}

type kAuthorizationClient struct {
	*kClient
}

func NewRPCAuthorizationClient(cli client.Client) Client {
	return &kAuthorizationClient{
		kClient: newServiceClient(cli),
	}
}

func (p *kAuthorizationClient) AuthSendCode(ctx context.Context, req *tg.TLAuthSendCode, callOptions ...callopt.Option) (r *tg.AuthSentCode, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.AuthSendCode(ctx, req)
}

func (p *kAuthorizationClient) AuthSignUp(ctx context.Context, req *tg.TLAuthSignUp, callOptions ...callopt.Option) (r *tg.AuthAuthorization, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.AuthSignUp(ctx, req)
}

func (p *kAuthorizationClient) AuthSignIn(ctx context.Context, req *tg.TLAuthSignIn, callOptions ...callopt.Option) (r *tg.AuthAuthorization, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.AuthSignIn(ctx, req)
}

func (p *kAuthorizationClient) AuthLogOut(ctx context.Context, req *tg.TLAuthLogOut, callOptions ...callopt.Option) (r *tg.AuthLoggedOut, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.AuthLogOut(ctx, req)
}

func (p *kAuthorizationClient) AuthResetAuthorizations(ctx context.Context, req *tg.TLAuthResetAuthorizations, callOptions ...callopt.Option) (r *tg.Bool, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.AuthResetAuthorizations(ctx, req)
}

func (p *kAuthorizationClient) AuthExportAuthorization(ctx context.Context, req *tg.TLAuthExportAuthorization, callOptions ...callopt.Option) (r *tg.AuthExportedAuthorization, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.AuthExportAuthorization(ctx, req)
}

func (p *kAuthorizationClient) AuthImportAuthorization(ctx context.Context, req *tg.TLAuthImportAuthorization, callOptions ...callopt.Option) (r *tg.AuthAuthorization, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.AuthImportAuthorization(ctx, req)
}

func (p *kAuthorizationClient) AuthBindTempAuthKey(ctx context.Context, req *tg.TLAuthBindTempAuthKey, callOptions ...callopt.Option) (r *tg.Bool, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.AuthBindTempAuthKey(ctx, req)
}

func (p *kAuthorizationClient) AuthImportBotAuthorization(ctx context.Context, req *tg.TLAuthImportBotAuthorization, callOptions ...callopt.Option) (r *tg.AuthAuthorization, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.AuthImportBotAuthorization(ctx, req)
}

func (p *kAuthorizationClient) AuthCheckPassword(ctx context.Context, req *tg.TLAuthCheckPassword, callOptions ...callopt.Option) (r *tg.AuthAuthorization, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.AuthCheckPassword(ctx, req)
}

func (p *kAuthorizationClient) AuthRequestPasswordRecovery(ctx context.Context, req *tg.TLAuthRequestPasswordRecovery, callOptions ...callopt.Option) (r *tg.AuthPasswordRecovery, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.AuthRequestPasswordRecovery(ctx, req)
}

func (p *kAuthorizationClient) AuthRecoverPassword(ctx context.Context, req *tg.TLAuthRecoverPassword, callOptions ...callopt.Option) (r *tg.AuthAuthorization, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.AuthRecoverPassword(ctx, req)
}

func (p *kAuthorizationClient) AuthResendCode(ctx context.Context, req *tg.TLAuthResendCode, callOptions ...callopt.Option) (r *tg.AuthSentCode, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.AuthResendCode(ctx, req)
}

func (p *kAuthorizationClient) AuthCancelCode(ctx context.Context, req *tg.TLAuthCancelCode, callOptions ...callopt.Option) (r *tg.Bool, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.AuthCancelCode(ctx, req)
}

func (p *kAuthorizationClient) AuthDropTempAuthKeys(ctx context.Context, req *tg.TLAuthDropTempAuthKeys, callOptions ...callopt.Option) (r *tg.Bool, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.AuthDropTempAuthKeys(ctx, req)
}

func (p *kAuthorizationClient) AuthCheckRecoveryPassword(ctx context.Context, req *tg.TLAuthCheckRecoveryPassword, callOptions ...callopt.Option) (r *tg.Bool, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.AuthCheckRecoveryPassword(ctx, req)
}

func (p *kAuthorizationClient) AuthImportWebTokenAuthorization(ctx context.Context, req *tg.TLAuthImportWebTokenAuthorization, callOptions ...callopt.Option) (r *tg.AuthAuthorization, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.AuthImportWebTokenAuthorization(ctx, req)
}

func (p *kAuthorizationClient) AuthRequestFirebaseSms(ctx context.Context, req *tg.TLAuthRequestFirebaseSms, callOptions ...callopt.Option) (r *tg.Bool, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.AuthRequestFirebaseSms(ctx, req)
}

func (p *kAuthorizationClient) AuthResetLoginEmail(ctx context.Context, req *tg.TLAuthResetLoginEmail, callOptions ...callopt.Option) (r *tg.AuthSentCode, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.AuthResetLoginEmail(ctx, req)
}

func (p *kAuthorizationClient) AuthReportMissingCode(ctx context.Context, req *tg.TLAuthReportMissingCode, callOptions ...callopt.Option) (r *tg.Bool, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.AuthReportMissingCode(ctx, req)
}

func (p *kAuthorizationClient) AccountSendVerifyEmailCode(ctx context.Context, req *tg.TLAccountSendVerifyEmailCode, callOptions ...callopt.Option) (r *tg.AccountSentEmailCode, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.AccountSendVerifyEmailCode(ctx, req)
}

func (p *kAuthorizationClient) AccountVerifyEmail(ctx context.Context, req *tg.TLAccountVerifyEmail, callOptions ...callopt.Option) (r *tg.AccountEmailVerified, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.AccountVerifyEmail(ctx, req)
}

func (p *kAuthorizationClient) AccountResetPassword(ctx context.Context, req *tg.TLAccountResetPassword, callOptions ...callopt.Option) (r *tg.AccountResetPasswordResult, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.AccountResetPassword(ctx, req)
}

func (p *kAuthorizationClient) AccountSetAuthorizationTTL(ctx context.Context, req *tg.TLAccountSetAuthorizationTTL, callOptions ...callopt.Option) (r *tg.Bool, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.AccountSetAuthorizationTTL(ctx, req)
}

func (p *kAuthorizationClient) AccountChangeAuthorizationSettings(ctx context.Context, req *tg.TLAccountChangeAuthorizationSettings, callOptions ...callopt.Option) (r *tg.Bool, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.AccountChangeAuthorizationSettings(ctx, req)
}

func (p *kAuthorizationClient) AccountInvalidateSignInCodes(ctx context.Context, req *tg.TLAccountInvalidateSignInCodes, callOptions ...callopt.Option) (r *tg.Bool, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.AccountInvalidateSignInCodes(ctx, req)
}

func (p *kAuthorizationClient) AuthToggleBan(ctx context.Context, req *tg.TLAuthToggleBan, callOptions ...callopt.Option) (r *tg.PredefinedUser, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.AuthToggleBan(ctx, req)
}
