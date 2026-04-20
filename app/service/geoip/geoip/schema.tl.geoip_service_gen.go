/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2026-present,  Teamgram Authors.
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package geoip

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/teamgram/teamgram-server/v2/pkg/proto/bin"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/iface"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

var (
	_ iface.TLObject
	_ fmt.Stringer
	_ *tg.Bool
	_ bin.Fields
	_ json.Marshaler
)

// TLGeoipGetCountryAndRegionByIp <--
type TLGeoipGetCountryAndRegionByIp struct {
	ClazzID uint32 `json:"_id"`
	Ip      string `json:"ip"`
}

func (m *TLGeoipGetCountryAndRegionByIp) String() string {
	wrapper := iface.WithNameWrapper{ClazzName: ClazzName_geoip_getCountryAndRegionByIp, TLObject: m}
	return wrapper.String()
}

// Encode <--
func (m *TLGeoipGetCountryAndRegionByIp) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_geoip_getCountryAndRegionByIp, int(layer)); clazzId {
	case 0x676b04a4:
		x.PutClazzID(0x676b04a4)

		x.PutString(m.Ip)

		return nil
	default:
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_geoip_getCountryAndRegionByIp, layer)
	}
}

// Decode <--
func (m *TLGeoipGetCountryAndRegionByIp) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return err
		}
	}
	switch m.ClazzID {
	case 0x676b04a4:
		m.Ip, err = d.String()
		if err != nil {
			return err
		}

		return nil
	default:
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// Vector api result type
// ----------------------------------------------------------------------------
// VectorResList <--

// ----------------------------------------------------------------------------
// rpc

type RPCGeoip interface {
	GeoipGetCountryAndRegionByIp(ctx context.Context, in *TLGeoipGetCountryAndRegionByIp) (*Region, error)
}
