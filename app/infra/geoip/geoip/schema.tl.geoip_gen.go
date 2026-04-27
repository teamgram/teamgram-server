/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2026-present, The Teamgram Authors (https://teamgram.net).
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package geoip

import (
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

// RegionClazz <--
//   - TL_Region
type RegionClazz = *TLRegion

func DecodeRegionClazz(d *bin.Decoder) (RegionClazz, error) {
	// id, err := d.PeekClazzID()
	id, err := d.ClazzID()
	if err != nil {
		return nil, fmt.Errorf("unable to decode Region: constructor: %w", err)
	}

	switch id {
	case 0xca964a2f:
		x := &TLRegion{ClazzID: id, ClazzName2: ClazzName_region}
		if err := x.Decode(d); err != nil {
			return nil, err
		}
		return x, nil
	default:
		return nil, fmt.Errorf("unable to decode Region: invalid constructor %x", id)
	}

}

// TLRegion <--
type TLRegion struct {
	ClazzID    uint32 `json:"_id"`
	ClazzName2 string `json:"_name"`
	Region     string `json:"region"`
	IsoCode    string `json:"iso_code"`
}

func MakeTLRegion(m *TLRegion) *TLRegion {
	if m == nil {
		return nil
	}
	m.ClazzName2 = ClazzName_region

	return m
}

func (m *TLRegion) String() string {
	return iface.DebugStringWithName("region", m)
}

func (m *TLRegion) MarshalJSON() ([]byte, error) {
	if m == nil {
		return []byte("null"), nil
	}
	return iface.MarshalWithName("region", m)
}

// RegionClazzName <--
func (m *TLRegion) RegionClazzName() string {
	return ClazzName_region
}

// ClazzName <--
func (m *TLRegion) ClazzName() string {
	return m.ClazzName2
}

// ToRegion <--
func (m *TLRegion) ToRegion() *Region {
	if m == nil {
		return nil
	}

	return m

}

func (m *TLRegion) CalcSize(layer int32) int {
	switch clazzId := iface.GetClazzIDByName(ClazzName_region, int(layer)); clazzId {
	case 0xca964a2f:
		size := 4
		size += iface.CalcStringSize(m.Region)
		size += iface.CalcStringSize(m.IsoCode)

		return size
	default:
		return 0
	}
}

func (m *TLRegion) Validate(layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_region, int(layer)); clazzId {
	case 0xca964a2f:
		if err := iface.ValidateRequiredString("region", m.Region); err != nil {
			return err
		}

		if err := iface.ValidateRequiredString("iso_code", m.IsoCode); err != nil {
			return err
		}

		return nil
	default:
		return fmt.Errorf("unable to encode region: unsupported layer %d", layer)
	}
}

// Encode <--
func (m *TLRegion) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_region, int(layer)); clazzId {
	case 0xca964a2f:
		x.PutClazzID(0xca964a2f)

		x.PutString(m.Region)
		x.PutString(m.IsoCode)

		return nil
	default:
		return fmt.Errorf("unable to encode region: unsupported layer %d", layer)
	}
}

// Decode <--
func (m *TLRegion) Decode(d *bin.Decoder) (err error) {
	switch m.ClazzID {
	case 0xca964a2f:
		m.Region, err = d.String()
		if err != nil {
			return fmt.Errorf("unable to decode region#0xca964a2f: field region: %w", err)
		}
		m.IsoCode, err = d.String()
		if err != nil {
			return fmt.Errorf("unable to decode region#0xca964a2f: field iso_code: %w", err)
		}

		return nil
	default:
		return fmt.Errorf("unable to decode region: invalid constructor %x", m.ClazzID)
	}
}

// Region <--
type Region = TLRegion
