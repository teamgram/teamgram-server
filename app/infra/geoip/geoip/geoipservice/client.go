/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2026-present, The Teamgram Authors (https://teamgram.net).
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package geoipservice

import (
	"context"

	"github.com/teamgram/teamgram-server/v2/app/infra/geoip/geoip"
	"github.com/teamgram/teamgram-server/v2/pkg/net/kitex/codec"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/callopt"
)

var _ *tg.Bool

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	GeoipGetCountryAndRegionByIp(ctx context.Context, req *geoip.TLGeoipGetCountryAndRegionByIp, callOptions ...callopt.Option) (r *geoip.Region, err error)
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
	return &kGeoipClient{
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

type kGeoipClient struct {
	*kClient
}

func NewRPCGeoipClient(cli client.Client) Client {
	return &kGeoipClient{
		kClient: newServiceClient(cli),
	}
}

func (p *kGeoipClient) GeoipGetCountryAndRegionByIp(ctx context.Context, req *geoip.TLGeoipGetCountryAndRegionByIp, callOptions ...callopt.Option) (r *geoip.Region, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GeoipGetCountryAndRegionByIp(ctx, req)
}
