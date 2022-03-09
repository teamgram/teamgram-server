// Copyright (c) 2021-present,  Teamgram Studio (https://teamgram.io).
//  All rights reserved.
//
// Author: teamgramio (teamgram.io@gmail.com)
//

package dao

import (
	"net"

	"github.com/zeromicro/go-zero/core/logx"
)

func (d *Dao) CheckApiIdAndHash(apiId int32, apiHash string) error {
	// TODO(@benqi): check api_id and api_hash
	// 400	API_ID_INVALID	API ID无效
	// 400	API_ID_PUBLISHED_FLOOD	这个API ID已发布在某个地方，您现在不能使用

	_ = apiId
	_ = apiHash

	return nil
}

func (d *Dao) GetCountryAndRegionByIp(ip string) (string, string) {
	if d.MMDB == nil {
		return "UNKNOWN", ""
	} else {
		r, err := d.MMDB.City(net.ParseIP(ip))
		if err != nil {
			logx.Errorf("getCountryAndRegionByIp - error: %v", err)
			return "UNKNOWN", ""
		}

		return r.City.Names["en"] + ", " + r.Country.Names["en"], r.Country.IsoCode
	}
}
