/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2024-present,  Teamgram Authors.
 *  All rights reserved.
 *
 * Author: Benqi (wubenqi@gmail.com)
 */

package passportservice

import (
	"context"

	"github.com/teamgram/proto/v2/tg"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/callopt"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	AccountGetAuthorizations(ctx context.Context, req *tg.TLAccountGetAuthorizations, callOptions ...callopt.Option) (r *tg.AccountAuthorizations, err error)
	AccountGetAllSecureValues(ctx context.Context, req *tg.TLAccountGetAllSecureValues, callOptions ...callopt.Option) (r *tg.VectorSecureValue, err error)
	AccountGetSecureValue(ctx context.Context, req *tg.TLAccountGetSecureValue, callOptions ...callopt.Option) (r *tg.VectorSecureValue, err error)
	AccountSaveSecureValue(ctx context.Context, req *tg.TLAccountSaveSecureValue, callOptions ...callopt.Option) (r *tg.SecureValue, err error)
	AccountDeleteSecureValue(ctx context.Context, req *tg.TLAccountDeleteSecureValue, callOptions ...callopt.Option) (r *tg.Bool, err error)
	AccountGetAuthorizationForm(ctx context.Context, req *tg.TLAccountGetAuthorizationForm, callOptions ...callopt.Option) (r *tg.AccountAuthorizationForm, err error)
	AccountAcceptAuthorization(ctx context.Context, req *tg.TLAccountAcceptAuthorization, callOptions ...callopt.Option) (r *tg.Bool, err error)
	AccountSendVerifyPhoneCode(ctx context.Context, req *tg.TLAccountSendVerifyPhoneCode, callOptions ...callopt.Option) (r *tg.AuthSentCode, err error)
	AccountVerifyPhone(ctx context.Context, req *tg.TLAccountVerifyPhone, callOptions ...callopt.Option) (r *tg.Bool, err error)
	UsersSetSecureValueErrors(ctx context.Context, req *tg.TLUsersSetSecureValueErrors, callOptions ...callopt.Option) (r *tg.Bool, err error)
	HelpGetPassportConfig(ctx context.Context, req *tg.TLHelpGetPassportConfig, callOptions ...callopt.Option) (r *tg.HelpPassportConfig, err error)
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
	return &kPassportClient{
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

type kPassportClient struct {
	*kClient
}

func NewRPCPassportClient(cli client.Client) Client {
	return &kPassportClient{
		kClient: newServiceClient(cli),
	}
}

func (p *kPassportClient) AccountGetAuthorizations(ctx context.Context, req *tg.TLAccountGetAuthorizations, callOptions ...callopt.Option) (r *tg.AccountAuthorizations, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.AccountGetAuthorizations(ctx, req)
}

func (p *kPassportClient) AccountGetAllSecureValues(ctx context.Context, req *tg.TLAccountGetAllSecureValues, callOptions ...callopt.Option) (r *tg.VectorSecureValue, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.AccountGetAllSecureValues(ctx, req)
}

func (p *kPassportClient) AccountGetSecureValue(ctx context.Context, req *tg.TLAccountGetSecureValue, callOptions ...callopt.Option) (r *tg.VectorSecureValue, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.AccountGetSecureValue(ctx, req)
}

func (p *kPassportClient) AccountSaveSecureValue(ctx context.Context, req *tg.TLAccountSaveSecureValue, callOptions ...callopt.Option) (r *tg.SecureValue, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.AccountSaveSecureValue(ctx, req)
}

func (p *kPassportClient) AccountDeleteSecureValue(ctx context.Context, req *tg.TLAccountDeleteSecureValue, callOptions ...callopt.Option) (r *tg.Bool, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.AccountDeleteSecureValue(ctx, req)
}

func (p *kPassportClient) AccountGetAuthorizationForm(ctx context.Context, req *tg.TLAccountGetAuthorizationForm, callOptions ...callopt.Option) (r *tg.AccountAuthorizationForm, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.AccountGetAuthorizationForm(ctx, req)
}

func (p *kPassportClient) AccountAcceptAuthorization(ctx context.Context, req *tg.TLAccountAcceptAuthorization, callOptions ...callopt.Option) (r *tg.Bool, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.AccountAcceptAuthorization(ctx, req)
}

func (p *kPassportClient) AccountSendVerifyPhoneCode(ctx context.Context, req *tg.TLAccountSendVerifyPhoneCode, callOptions ...callopt.Option) (r *tg.AuthSentCode, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.AccountSendVerifyPhoneCode(ctx, req)
}

func (p *kPassportClient) AccountVerifyPhone(ctx context.Context, req *tg.TLAccountVerifyPhone, callOptions ...callopt.Option) (r *tg.Bool, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.AccountVerifyPhone(ctx, req)
}

func (p *kPassportClient) UsersSetSecureValueErrors(ctx context.Context, req *tg.TLUsersSetSecureValueErrors, callOptions ...callopt.Option) (r *tg.Bool, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UsersSetSecureValueErrors(ctx, req)
}

func (p *kPassportClient) HelpGetPassportConfig(ctx context.Context, req *tg.TLHelpGetPassportConfig, callOptions ...callopt.Option) (r *tg.HelpPassportConfig, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.HelpGetPassportConfig(ctx, req)
}
