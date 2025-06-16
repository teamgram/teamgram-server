/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2024-present,  Teamgram Authors.
 *  All rights reserved.
 *
 * Author: Benqi (wubenqi@gmail.com)
 */

package updatesservice

import (
	"context"

	"github.com/teamgram/proto/v2/tg"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/callopt"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	UpdatesGetState(ctx context.Context, req *tg.TLUpdatesGetState, callOptions ...callopt.Option) (r *tg.UpdatesState, err error)
	UpdatesGetDifference(ctx context.Context, req *tg.TLUpdatesGetDifference, callOptions ...callopt.Option) (r *tg.UpdatesDifference, err error)
	UpdatesGetChannelDifference(ctx context.Context, req *tg.TLUpdatesGetChannelDifference, callOptions ...callopt.Option) (r *tg.UpdatesChannelDifference, err error)
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
	return &kUpdatesClient{
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

type kUpdatesClient struct {
	*kClient
}

func NewRPCUpdatesClient(cli client.Client) Client {
	return &kUpdatesClient{
		kClient: newServiceClient(cli),
	}
}

func (p *kUpdatesClient) UpdatesGetState(ctx context.Context, req *tg.TLUpdatesGetState, callOptions ...callopt.Option) (r *tg.UpdatesState, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UpdatesGetState(ctx, req)
}

func (p *kUpdatesClient) UpdatesGetDifference(ctx context.Context, req *tg.TLUpdatesGetDifference, callOptions ...callopt.Option) (r *tg.UpdatesDifference, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UpdatesGetDifference(ctx, req)
}

func (p *kUpdatesClient) UpdatesGetChannelDifference(ctx context.Context, req *tg.TLUpdatesGetChannelDifference, callOptions ...callopt.Option) (r *tg.UpdatesChannelDifference, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UpdatesGetChannelDifference(ctx, req)
}
