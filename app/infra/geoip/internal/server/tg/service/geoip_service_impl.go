/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2026-present, The Teamgram Authors (https://teamgram.net).
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package service

import (
	"context"

	"github.com/teamgram/teamgram-server/v2/app/infra/geoip/geoip"
	"github.com/teamgram/teamgram-server/v2/app/infra/geoip/internal/core"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

var _ *tg.Bool

// GeoipGetCountryAndRegionByIp
// geoip.getCountryAndRegionByIp ip:string = Region;
func (s *Service) GeoipGetCountryAndRegionByIp(ctx context.Context, request *geoip.TLGeoipGetCountryAndRegionByIp) (*geoip.Region, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("geoip.getCountryAndRegionByIp - request: %s", request)

	r, err := c.GeoipGetCountryAndRegionByIp(request)
	if err != nil {
		c.Logger.Errorf("geoip.getCountryAndRegionByIp - error: request: %s, err: %v", request, err)
		return nil, err
	}

	c.Logger.Debugf("geoip.getCountryAndRegionByIp - reply: %s", r)
	return r, err
}
