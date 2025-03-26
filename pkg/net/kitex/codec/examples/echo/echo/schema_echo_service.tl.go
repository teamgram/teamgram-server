/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2024-present,  Teamgram Authors.
 *  All rights reserved.
 *
 * Author: Benqi (wubenqi@gmail.com)
 */

package echo

import (
	"context"
	"fmt"

	"github.com/teamgram/proto/v2/bin"
	"github.com/teamgram/proto/v2/iface"
	"github.com/teamgram/proto/v2/tg"
)

var _ iface.TLObject
var _ fmt.Stringer
var _ *tg.Bool
var _ bin.Fields

// TLEchoEcho <--
type TLEchoEcho struct {
	ClazzID uint32 `json:"_id"`
	Message string `json:"message"`
}

func NewTLEchoEchoArg() interface{} {
	return &TLEchoEcho{}
}

// Encode <--
func (m *TLEchoEcho) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0xf653b67d: func() error {
			x.PutClazzID(0xf653b67d)

			x.PutString(m.Message)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_echo_echo, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_echo_echo, layer)
	}
}

// Decode <--
func (m *TLEchoEcho) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0xf653b67d: func() (err error) {
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

type RPCEcho interface {
	EchoEcho(ctx context.Context, in *TLEchoEcho) (*Echo, error)
}
