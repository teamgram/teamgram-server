/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright 2026 Teamgram Authors.
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package geoipclient

import (
	"context"

	"github.com/teamgram/teamgram-server/v2/app/service/geoip/geoip"
	"github.com/teamgram/teamgram-server/v2/app/service/geoip/geoip/geoipservice"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"

	"github.com/cloudwego/kitex/client"
)

var _ *tg.Bool

type GeoipClient interface {
	GeoipGetCountryAndRegionByIp(ctx context.Context, in *geoip.TLGeoipGetCountryAndRegionByIp) (*geoip.Region, error)
	Close() error
}

type defaultGeoipClient struct {
	cli client.Client
}

func NewGeoipClient(cli client.Client) GeoipClient {
	return &defaultGeoipClient{
		cli: cli,
	}
}

func (m *defaultGeoipClient) Close() error {
	if closer, ok := any(m.cli).(interface{ Close() error }); ok {
		return closer.Close()
	}
	return nil
}

// GeoipGetCountryAndRegionByIp
// geoip.getCountryAndRegionByIp ip:string = Region;
func (m *defaultGeoipClient) GeoipGetCountryAndRegionByIp(ctx context.Context, in *geoip.TLGeoipGetCountryAndRegionByIp) (*geoip.Region, error) {
	cli := geoipservice.NewRPCGeoipClient(m.cli)
	return cli.GeoipGetCountryAndRegionByIp(ctx, in)
}
