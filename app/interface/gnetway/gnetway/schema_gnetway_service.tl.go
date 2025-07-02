/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2025-present,  Teamgram Authors.
 *  All rights reserved.
 *
 * Author: Benqi (wubenqi@gmail.com)
 */

package gnetway

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/teamgram/proto/v2/bin"
	"github.com/teamgram/proto/v2/iface"
	"github.com/teamgram/proto/v2/tg"
)

var (
	_ iface.TLObject
	_ fmt.Stringer
	_ *tg.Bool
	_ bin.Fields
	_ json.Marshaler
)

// TLGnetwaySendDataToGateway <--
type TLGnetwaySendDataToGateway struct {
	ClazzID   uint32 `json:"_id"`
	AuthKeyId int64  `json:"auth_key_id"`
	SessionId int64  `json:"session_id"`
	Payload   []byte `json:"payload"`
}

func (m *TLGnetwaySendDataToGateway) String() string {
	wrapper := iface.WithNameWrapper{"", m}
	return wrapper.String()
}

// Encode <--
func (m *TLGnetwaySendDataToGateway) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0x722d5ce0: func() error {
			x.PutClazzID(0x722d5ce0)

			x.PutInt64(m.AuthKeyId)
			x.PutInt64(m.SessionId)
			x.PutBytes(m.Payload)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_gnetway_sendDataToGateway, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_gnetway_sendDataToGateway, layer)
	}
}

// Decode <--
func (m *TLGnetwaySendDataToGateway) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0x722d5ce0: func() (err error) {
			m.AuthKeyId, err = d.Int64()
			m.SessionId, err = d.Int64()
			m.Payload, err = d.Bytes()

			return nil
		},
	}

	if m.ClazzID == 0 {
		m.ClazzID, _ = d.ClazzID()
	}
	if f, ok := decodeF[m.ClazzID]; ok {
		return f()
	} else {
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// Vector api result type
// ----------------------------------------------------------------------------
// VectorResList <--

// ----------------------------------------------------------------------------
// rpc

type RPCGnetway interface {
	GnetwaySendDataToGateway(ctx context.Context, in *TLGnetwaySendDataToGateway) (*tg.Bool, error)
}
