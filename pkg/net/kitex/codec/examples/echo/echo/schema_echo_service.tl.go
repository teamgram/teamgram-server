/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2026-present,  Teamgram Authors.
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
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

// TLEcho1Echo <--
type TLEcho1Echo struct {
	ClazzID uint32 `json:"_id"`
	Message string `json:"message"`
}

func (m *TLEcho1Echo) String() string {
	wrapper := iface.WithNameWrapper{ClazzName: ClazzName_echo1_echo, TLObject: m}
	return wrapper.String()
}

// Encode <--
func (m *TLEcho1Echo) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_echo1_echo, int(layer)); clazzId {
	case 0x9f0506e2:
		x.PutClazzID(0x9f0506e2)

		x.PutString(m.Message)

		return nil
	default:
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_echo1_echo, layer)
	}
}

// Decode <--
func (m *TLEcho1Echo) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return err
		}
	}
	switch m.ClazzID {
	case 0x9f0506e2:
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

type RPCEcho1 interface {
	Echo1Echo(ctx context.Context, in *TLEcho1Echo) (*Echo, error)
}
