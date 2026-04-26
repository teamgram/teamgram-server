/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2026-present, The Teamgram Authors (https://teamgram.net).
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package geoipclient

import (
	"context"

	"github.com/teamgram/teamgram-server/v2/app/infra/geoip/geoip"
	"github.com/teamgram/teamgram-server/v2/app/infra/geoip/geoip/geoipservice"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"

	"github.com/cloudwego/kitex/client"
)

var _ *tg.Bool

type GeoipClient interface {
	GeoipGetCountryAndRegionByIp(ctx context.Context, in *geoip.TLGeoipGetCountryAndRegionByIp) (*geoip.Region, error)
}

type defaultGeoipClient struct {
	cli client.Client
	rpc geoipservice.Client
}

func NewGeoipClient(cli client.Client) GeoipClient {
	return &defaultGeoipClient{
		cli: cli,
		rpc: geoipservice.NewRPCGeoipClient(cli),
	}
}

// GeoipGetCountryAndRegionByIp
// geoip.getCountryAndRegionByIp ip:string = Region;
func (m *defaultGeoipClient) GeoipGetCountryAndRegionByIp(ctx context.Context, in *geoip.TLGeoipGetCountryAndRegionByIp) (*geoip.Region, error) {
	return m.rpc.GeoipGetCountryAndRegionByIp(ctx, in)
}
