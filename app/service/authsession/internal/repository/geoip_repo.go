package repository

import (
	"context"

	"github.com/teamgram/teamgram-server/v2/app/infra/geoip/geoip"

	"github.com/zeromicro/go-zero/core/logx"
)

func (r *Repository) getCountryAndRegionByIP(ctx context.Context, ip string) (country string, region string) {
	if r.geoipClient == nil {
		return "", ""
	}

	rValue, err := r.geoipClient.GeoipGetCountryAndRegionByIp(ctx, &geoip.TLGeoipGetCountryAndRegionByIp{Ip: ip})
	if err != nil {
		logx.WithContext(ctx).Errorf("getCountryAndRegionByIP(%s) error: %v", ip, err)
		return "", ""
	}
	if rValue == nil {
		return "", ""
	}

	return rValue.IsoCode, rValue.Region
}
