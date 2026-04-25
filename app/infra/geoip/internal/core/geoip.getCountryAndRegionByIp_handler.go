// Copyright (c) 2026 The Teamgram Authors. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Author: teamgramio (teamgram.io@gmail.com)

package core

import (
	"net"

	"github.com/teamgram/teamgram-server/v2/app/infra/geoip/geoip"
)

var (
	unknownRegion = geoip.MakeTLRegion(&geoip.TLRegion{
		Region:  "UNKNOWN",
		IsoCode: "",
	})
)

// GeoipGetCountryAndRegionByIp
// geoip.getCountryAndRegionByIp = Region;
func (c *GeoipCore) GeoipGetCountryAndRegionByIp(in *geoip.TLGeoipGetCountryAndRegionByIp) (*geoip.Region, error) {
	if c.svcCtx.Repo.MMDB == nil {
		return unknownRegion, nil
	}

	r, err := c.svcCtx.Repo.MMDB.City(net.ParseIP(in.Ip))
	if err != nil {
		c.Logger.Errorf("getCountryAndRegionByIp - error: %v", err)

		return unknownRegion, nil
	}

	return geoip.MakeTLRegion(&geoip.TLRegion{
		Region:  r.City.Names["en"] + ", " + r.Country.Names["en"],
		IsoCode: r.Country.IsoCode,
	}), nil
}
