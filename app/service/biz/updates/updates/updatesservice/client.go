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
	"github.com/teamgram/teamgram-server/v2/app/service/biz/updates/updates"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/callopt"
)

var _ *tg.Bool

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	UpdatesGetStateV2(ctx context.Context, req *updates.TLUpdatesGetStateV2, callOptions ...callopt.Option) (r *tg.UpdatesState, err error)
	UpdatesGetDifferenceV2(ctx context.Context, req *updates.TLUpdatesGetDifferenceV2, callOptions ...callopt.Option) (r *updates.Difference, err error)
	UpdatesGetChannelDifferenceV2(ctx context.Context, req *updates.TLUpdatesGetChannelDifferenceV2, callOptions ...callopt.Option) (r *updates.ChannelDifference, err error)
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

func (p *kUpdatesClient) UpdatesGetStateV2(ctx context.Context, req *updates.TLUpdatesGetStateV2, callOptions ...callopt.Option) (r *tg.UpdatesState, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UpdatesGetStateV2(ctx, req)
}

func (p *kUpdatesClient) UpdatesGetDifferenceV2(ctx context.Context, req *updates.TLUpdatesGetDifferenceV2, callOptions ...callopt.Option) (r *updates.Difference, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UpdatesGetDifferenceV2(ctx, req)
}

func (p *kUpdatesClient) UpdatesGetChannelDifferenceV2(ctx context.Context, req *updates.TLUpdatesGetChannelDifferenceV2, callOptions ...callopt.Option) (r *updates.ChannelDifference, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UpdatesGetChannelDifferenceV2(ctx, req)
}
