/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2026-present,  Teamgram Authors.
 *  All rights reserved.
 *
 * Author: Benqi (wubenqi@gmail.com)
 */

package echo

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

// TLEchoEcho <--
type TLEchoEcho struct {
	ClazzID uint32 `json:"_id"`
	Message string `json:"message"`
}

func (m *TLEchoEcho) String() string {
	wrapper := iface.WithNameWrapper{ClazzName: ClazzName_echo_echo, TLObject: m}
	return wrapper.String()
}

// Encode <--
func (m *TLEchoEcho) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_echo_echo, int(layer)); clazzId {
	case 0xf653b67d:
		x.PutClazzID(0xf653b67d)

		x.PutString(m.Message)

		return nil
	default:
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_echo_echo, layer)
	}
}

// Decode <--
func (m *TLEchoEcho) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return err
		}
	}
	switch m.ClazzID {
	case 0xf653b67d:
		m.Message, err = d.String()
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

type RPCEcho interface {
	EchoEcho(ctx context.Context, in *TLEchoEcho) (*Echo, error)
}
