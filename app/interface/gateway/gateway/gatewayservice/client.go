/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2026-present, The Teamgram Authors (https://teamgram.net).
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package gatewayservice

import (
	"context"

	"github.com/teamgram/teamgram-server/v2/app/interface/gateway/gateway"
	"github.com/teamgram/teamgram-server/v2/pkg/net/kitex/codec"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/callopt"
)

var _ *tg.Bool

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	GatewayPushUpdatesData(ctx context.Context, req *gateway.TLGatewayPushUpdatesData, callOptions ...callopt.Option) (r *tg.Bool, err error)
	GatewayPushSessionUpdatesData(ctx context.Context, req *gateway.TLGatewayPushSessionUpdatesData, callOptions ...callopt.Option) (r *tg.Bool, err error)
	GatewayPushRpcResultData(ctx context.Context, req *gateway.TLGatewayPushRpcResultData, callOptions ...callopt.Option) (r *tg.Bool, err error)
}

// Deprecated: prefer the generated app client helper or pkg/net/kitex.NewClient for TL-aware transport setup.
// NewClient creates a client for the service defined in IDL.
func NewClient(destService string, opts ...client.Option) (Client, error) {
	var options []client.Option
	options = append(options, client.WithDestService(destService))
	options = append(options, client.WithCodec(codec.NewZRpcCodec(false)))

	options = append(options, opts...)

	kc, err := client.NewClient(serviceInfoForClient(), options...)
	if err != nil {
		return nil, err
	}
	return &kGatewayClient{
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

type kGatewayClient struct {
	*kClient
}

func NewRPCGatewayClient(cli client.Client) Client {
	return &kGatewayClient{
		kClient: newServiceClient(cli),
	}
}

func (p *kGatewayClient) GatewayPushUpdatesData(ctx context.Context, req *gateway.TLGatewayPushUpdatesData, callOptions ...callopt.Option) (r *tg.Bool, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GatewayPushUpdatesData(ctx, req)
}

func (p *kGatewayClient) GatewayPushSessionUpdatesData(ctx context.Context, req *gateway.TLGatewayPushSessionUpdatesData, callOptions ...callopt.Option) (r *tg.Bool, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GatewayPushSessionUpdatesData(ctx, req)
}

func (p *kGatewayClient) GatewayPushRpcResultData(ctx context.Context, req *gateway.TLGatewayPushRpcResultData, callOptions ...callopt.Option) (r *tg.Bool, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GatewayPushRpcResultData(ctx, req)
}
