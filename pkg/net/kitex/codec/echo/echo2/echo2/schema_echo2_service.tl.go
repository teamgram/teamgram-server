/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2025-present,  Teamgram Authors.
 *  All rights reserved.
 *
 * Author: Benqi (wubenqi@gmail.com)
 */

package echo2

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

// TLEcho2Echo <--
type TLEcho2Echo struct {
	ClazzID uint32 `json:"_id"`
	Message string `json:"message"`
}

func (m *TLEcho2Echo) String() string {
	wrapper := iface.WithNameWrapper{"", m}
	return wrapper.String()
}

// Encode <--
func (m *TLEcho2Echo) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0x9ddb01c5: func() error {
			x.PutClazzID(0x9ddb01c5)

			x.PutString(m.Message)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_echo2_echo, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_echo2_echo, layer)
	}
}

// Decode <--
func (m *TLEcho2Echo) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0x9ddb01c5: func() (err error) {
			m.Message, err = d.String()

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

type RPCEcho2 interface {
	Echo2Echo(ctx context.Context, in *TLEcho2Echo) (*Echo, error)
}
