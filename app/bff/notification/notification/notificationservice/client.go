/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2024-present,  Teamgram Authors.
 *  All rights reserved.
 *
 * Author: Benqi (wubenqi@gmail.com)
 */

package notificationservice

import (
	"context"

	"github.com/teamgram/proto/v2/tg"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/callopt"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	AccountRegisterDevice(ctx context.Context, req *tg.TLAccountRegisterDevice, callOptions ...callopt.Option) (r *tg.Bool, err error)
	AccountUnregisterDevice(ctx context.Context, req *tg.TLAccountUnregisterDevice, callOptions ...callopt.Option) (r *tg.Bool, err error)
	AccountUpdateNotifySettings(ctx context.Context, req *tg.TLAccountUpdateNotifySettings, callOptions ...callopt.Option) (r *tg.Bool, err error)
	AccountGetNotifySettings(ctx context.Context, req *tg.TLAccountGetNotifySettings, callOptions ...callopt.Option) (r *tg.PeerNotifySettings, err error)
	AccountResetNotifySettings(ctx context.Context, req *tg.TLAccountResetNotifySettings, callOptions ...callopt.Option) (r *tg.Bool, err error)
	AccountUpdateDeviceLocked(ctx context.Context, req *tg.TLAccountUpdateDeviceLocked, callOptions ...callopt.Option) (r *tg.Bool, err error)
	AccountGetNotifyExceptions(ctx context.Context, req *tg.TLAccountGetNotifyExceptions, callOptions ...callopt.Option) (r *tg.Updates, err error)
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
	return &kNotificationClient{
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

type kNotificationClient struct {
	*kClient
}

func NewRPCNotificationClient(cli client.Client) Client {
	return &kNotificationClient{
		kClient: newServiceClient(cli),
	}
}

func (p *kNotificationClient) AccountRegisterDevice(ctx context.Context, req *tg.TLAccountRegisterDevice, callOptions ...callopt.Option) (r *tg.Bool, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.AccountRegisterDevice(ctx, req)
}

func (p *kNotificationClient) AccountUnregisterDevice(ctx context.Context, req *tg.TLAccountUnregisterDevice, callOptions ...callopt.Option) (r *tg.Bool, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.AccountUnregisterDevice(ctx, req)
}

func (p *kNotificationClient) AccountUpdateNotifySettings(ctx context.Context, req *tg.TLAccountUpdateNotifySettings, callOptions ...callopt.Option) (r *tg.Bool, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.AccountUpdateNotifySettings(ctx, req)
}

func (p *kNotificationClient) AccountGetNotifySettings(ctx context.Context, req *tg.TLAccountGetNotifySettings, callOptions ...callopt.Option) (r *tg.PeerNotifySettings, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.AccountGetNotifySettings(ctx, req)
}

func (p *kNotificationClient) AccountResetNotifySettings(ctx context.Context, req *tg.TLAccountResetNotifySettings, callOptions ...callopt.Option) (r *tg.Bool, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.AccountResetNotifySettings(ctx, req)
}

func (p *kNotificationClient) AccountUpdateDeviceLocked(ctx context.Context, req *tg.TLAccountUpdateDeviceLocked, callOptions ...callopt.Option) (r *tg.Bool, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.AccountUpdateDeviceLocked(ctx, req)
}

func (p *kNotificationClient) AccountGetNotifyExceptions(ctx context.Context, req *tg.TLAccountGetNotifyExceptions, callOptions ...callopt.Option) (r *tg.Updates, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.AccountGetNotifyExceptions(ctx, req)
}
