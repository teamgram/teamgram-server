/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2026-present,  Teamgram Authors.
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package updates

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

// TLUpdatesGetStateV2 <--
type TLUpdatesGetStateV2 struct {
	ClazzID   uint32 `json:"_id"`
	AuthKeyId int64  `json:"auth_key_id"`
	UserId    int64  `json:"user_id"`
}

func (m *TLUpdatesGetStateV2) String() string {
	wrapper := iface.WithNameWrapper{ClazzName: ClazzName_updates_getStateV2, TLObject: m}
	return wrapper.String()
}

// Encode <--
func (m *TLUpdatesGetStateV2) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_updates_getStateV2, int(layer)); clazzId {
	case 0x45f4cd65:
		x.PutClazzID(0x45f4cd65)

		x.PutInt64(m.AuthKeyId)
		x.PutInt64(m.UserId)

		return nil
	default:
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_updates_getStateV2, layer)
	}
}

// Decode <--
func (m *TLUpdatesGetStateV2) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return err
		}
	}
	switch m.ClazzID {
	case 0x45f4cd65:
		m.AuthKeyId, err = d.Int64()
		if err != nil {
			return err
		}
		m.UserId, err = d.Int64()
		if err != nil {
			return err
		}

		return nil
	default:
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// TLUpdatesGetDifferenceV2 <--
type TLUpdatesGetDifferenceV2 struct {
	ClazzID       uint32 `json:"_id"`
	AuthKeyId     int64  `json:"auth_key_id"`
	UserId        int64  `json:"user_id"`
	Pts           int32  `json:"pts"`
	PtsTotalLimit *int32 `json:"pts_total_limit"`
	Date          int64  `json:"date"`
}

func (m *TLUpdatesGetDifferenceV2) String() string {
	wrapper := iface.WithNameWrapper{ClazzName: ClazzName_updates_getDifferenceV2, TLObject: m}
	return wrapper.String()
}

// Encode <--
func (m *TLUpdatesGetDifferenceV2) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_updates_getDifferenceV2, int(layer)); clazzId {
	case 0xb76b6699:
		x.PutClazzID(0xb76b6699)

		// set flags
		var getFlags = func() uint32 {
			var flags uint32 = 0

			if m.PtsTotalLimit != nil {
				flags |= 1 << 0
			}

			return flags
		}

		// set flags
		var flags = getFlags()
		x.PutUint32(flags)
		x.PutInt64(m.AuthKeyId)
		x.PutInt64(m.UserId)
		x.PutInt32(m.Pts)
		if m.PtsTotalLimit != nil {
			x.PutInt32(*m.PtsTotalLimit)
		}

		x.PutInt64(m.Date)

		return nil
	default:
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_updates_getDifferenceV2, layer)
	}
}

// Decode <--
func (m *TLUpdatesGetDifferenceV2) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return err
		}
	}
	switch m.ClazzID {
	case 0xb76b6699:
		flags, err := d.Uint32()
		if err != nil {
			return err
		}
		_ = flags
		m.AuthKeyId, err = d.Int64()
		if err != nil {
			return err
		}
		m.UserId, err = d.Int64()
		if err != nil {
			return err
		}
		m.Pts, err = d.Int32()
		if err != nil {
			return err
		}
		if (flags & (1 << 0)) != 0 {
			m.PtsTotalLimit = new(int32)
			*m.PtsTotalLimit, err = d.Int32()
			if err != nil {
				return err
			}
		}
		m.Date, err = d.Int64()
		if err != nil {
			return err
		}

		return nil
	default:
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// TLUpdatesGetChannelDifferenceV2 <--
type TLUpdatesGetChannelDifferenceV2 struct {
	ClazzID   uint32 `json:"_id"`
	AuthKeyId int64  `json:"auth_key_id"`
	UserId    int64  `json:"user_id"`
	ChannelId int64  `json:"channel_id"`
	Pts       int32  `json:"pts"`
	Limit     int32  `json:"limit"`
}

func (m *TLUpdatesGetChannelDifferenceV2) String() string {
	wrapper := iface.WithNameWrapper{ClazzName: ClazzName_updates_getChannelDifferenceV2, TLObject: m}
	return wrapper.String()
}

// Encode <--
func (m *TLUpdatesGetChannelDifferenceV2) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_updates_getChannelDifferenceV2, int(layer)); clazzId {
	case 0x4da3318a:
		x.PutClazzID(0x4da3318a)

		x.PutInt64(m.AuthKeyId)
		x.PutInt64(m.UserId)
		x.PutInt64(m.ChannelId)
		x.PutInt32(m.Pts)
		x.PutInt32(m.Limit)

		return nil
	default:
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_updates_getChannelDifferenceV2, layer)
	}
}

// Decode <--
func (m *TLUpdatesGetChannelDifferenceV2) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return err
		}
	}
	switch m.ClazzID {
	case 0x4da3318a:
		m.AuthKeyId, err = d.Int64()
		if err != nil {
			return err
		}
		m.UserId, err = d.Int64()
		if err != nil {
			return err
		}
		m.ChannelId, err = d.Int64()
		if err != nil {
			return err
		}
		m.Pts, err = d.Int32()
		if err != nil {
			return err
		}
		m.Limit, err = d.Int32()
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

type RPCUpdates interface {
	UpdatesGetStateV2(ctx context.Context, in *TLUpdatesGetStateV2) (*tg.UpdatesState, error)
	UpdatesGetDifferenceV2(ctx context.Context, in *TLUpdatesGetDifferenceV2) (*Difference, error)
	UpdatesGetChannelDifferenceV2(ctx context.Context, in *TLUpdatesGetChannelDifferenceV2) (*ChannelDifference, error)
}
