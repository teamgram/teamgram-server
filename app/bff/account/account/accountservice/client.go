/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2024-present,  Teamgram Authors.
 *  All rights reserved.
 *
 * Author: Benqi (wubenqi@gmail.com)
 */

package accountservice

import (
	"context"

	"github.com/teamgram/proto/v2/tg"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/callopt"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	AccountDeleteAccount(ctx context.Context, req *tg.TLAccountDeleteAccount, callOptions ...callopt.Option) (r *tg.Bool, err error)
	AccountGetAccountTTL(ctx context.Context, req *tg.TLAccountGetAccountTTL, callOptions ...callopt.Option) (r *tg.AccountDaysTTL, err error)
	AccountSetAccountTTL(ctx context.Context, req *tg.TLAccountSetAccountTTL, callOptions ...callopt.Option) (r *tg.Bool, err error)
	AccountSendChangePhoneCode(ctx context.Context, req *tg.TLAccountSendChangePhoneCode, callOptions ...callopt.Option) (r *tg.AuthSentCode, err error)
	AccountChangePhone(ctx context.Context, req *tg.TLAccountChangePhone, callOptions ...callopt.Option) (r *tg.User, err error)
	AccountResetAuthorization(ctx context.Context, req *tg.TLAccountResetAuthorization, callOptions ...callopt.Option) (r *tg.Bool, err error)
	AccountSendConfirmPhoneCode(ctx context.Context, req *tg.TLAccountSendConfirmPhoneCode, callOptions ...callopt.Option) (r *tg.AuthSentCode, err error)
	AccountConfirmPhone(ctx context.Context, req *tg.TLAccountConfirmPhone, callOptions ...callopt.Option) (r *tg.Bool, err error)
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
	return &kAccountClient{
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

type kAccountClient struct {
	*kClient
}

func NewRPCAccountClient(cli client.Client) Client {
	return &kAccountClient{
		kClient: newServiceClient(cli),
	}
}

func (p *kAccountClient) AccountDeleteAccount(ctx context.Context, req *tg.TLAccountDeleteAccount, callOptions ...callopt.Option) (r *tg.Bool, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.AccountDeleteAccount(ctx, req)
}

func (p *kAccountClient) AccountGetAccountTTL(ctx context.Context, req *tg.TLAccountGetAccountTTL, callOptions ...callopt.Option) (r *tg.AccountDaysTTL, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.AccountGetAccountTTL(ctx, req)
}

func (p *kAccountClient) AccountSetAccountTTL(ctx context.Context, req *tg.TLAccountSetAccountTTL, callOptions ...callopt.Option) (r *tg.Bool, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.AccountSetAccountTTL(ctx, req)
}

func (p *kAccountClient) AccountSendChangePhoneCode(ctx context.Context, req *tg.TLAccountSendChangePhoneCode, callOptions ...callopt.Option) (r *tg.AuthSentCode, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.AccountSendChangePhoneCode(ctx, req)
}

func (p *kAccountClient) AccountChangePhone(ctx context.Context, req *tg.TLAccountChangePhone, callOptions ...callopt.Option) (r *tg.User, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.AccountChangePhone(ctx, req)
}

func (p *kAccountClient) AccountResetAuthorization(ctx context.Context, req *tg.TLAccountResetAuthorization, callOptions ...callopt.Option) (r *tg.Bool, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.AccountResetAuthorization(ctx, req)
}

func (p *kAccountClient) AccountSendConfirmPhoneCode(ctx context.Context, req *tg.TLAccountSendConfirmPhoneCode, callOptions ...callopt.Option) (r *tg.AuthSentCode, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.AccountSendConfirmPhoneCode(ctx, req)
}

func (p *kAccountClient) AccountConfirmPhone(ctx context.Context, req *tg.TLAccountConfirmPhone, callOptions ...callopt.Option) (r *tg.Bool, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.AccountConfirmPhone(ctx, req)
}
