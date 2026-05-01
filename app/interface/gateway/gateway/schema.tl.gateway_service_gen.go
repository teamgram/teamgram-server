/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2026-present, The Teamgram Authors (https://teamgram.net).
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package gateway

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

// TLGatewayPushUpdatesData <--
type TLGatewayPushUpdatesData struct {
	ClazzID       uint32          `json:"_id"`
	PermAuthKeyId int64           `json:"perm_auth_key_id"`
	Notification  bool            `json:"notification"`
	Updates       tg.UpdatesClazz `json:"updates"`
}

func (m *TLGatewayPushUpdatesData) String() string {
	return iface.DebugStringWithName(ClazzName_gateway_pushUpdatesData, m)
}

// Encode <--
func (m *TLGatewayPushUpdatesData) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_gateway_pushUpdatesData, int(layer)); clazzId {
	case 0x10dcca87:
		x.PutClazzID(0x10dcca87)

		// set flags
		var getFlags = func() uint32 {
			var flags uint32 = 0

			if m.Notification == true {
				flags |= 1 << 0
			}

			return flags
		}

		// set flags
		var flags = getFlags()
		x.PutUint32(flags)
		x.PutInt64(m.PermAuthKeyId)
		if m.Updates == nil {
			return fmt.Errorf("unable to encode gateway_pushUpdatesData#0x10dcca87: field updates is nil")
		}
		if err := m.Updates.Encode(x, layer); err != nil {
			return fmt.Errorf("unable to encode gateway_pushUpdatesData#0x10dcca87: field updates: %w", err)
		}

		return nil
	default:
		return fmt.Errorf("unable to encode gateway_pushUpdatesData: unsupported layer %d", layer)
	}
}

// Decode <--
func (m *TLGatewayPushUpdatesData) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return fmt.Errorf("unable to decode gateway_pushUpdatesData: constructor: %w", err)
		}
	}
	switch m.ClazzID {
	case 0x10dcca87:
		flags, err := d.Uint32()
		if err != nil {
			return fmt.Errorf("unable to decode gateway_pushUpdatesData: field flags: %w", err)
		}
		_ = flags
		m.PermAuthKeyId, err = d.Int64()
		if err != nil {
			return fmt.Errorf("unable to decode gateway_pushUpdatesData#0x10dcca87: field perm_auth_key_id: %w", err)
		}
		if (flags & (1 << 0)) != 0 {
			m.Notification = true
		}

		m.Updates, err = tg.DecodeUpdatesClazz(d)
		if err != nil {
			return fmt.Errorf("unable to decode gateway_pushUpdatesData#0x10dcca87: field updates: %w", err)
		}

		return nil
	default:
		return fmt.Errorf("unable to decode gateway_pushUpdatesData: invalid constructor %x", m.ClazzID)
	}
}

// TLGatewayPushSessionUpdatesData <--
type TLGatewayPushSessionUpdatesData struct {
	ClazzID       uint32          `json:"_id"`
	PermAuthKeyId int64           `json:"perm_auth_key_id"`
	AuthKeyId     int64           `json:"auth_key_id"`
	SessionId     int64           `json:"session_id"`
	Updates       tg.UpdatesClazz `json:"updates"`
}

func (m *TLGatewayPushSessionUpdatesData) String() string {
	return iface.DebugStringWithName(ClazzName_gateway_pushSessionUpdatesData, m)
}

// Encode <--
func (m *TLGatewayPushSessionUpdatesData) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_gateway_pushSessionUpdatesData, int(layer)); clazzId {
	case 0x794c7ded:
		x.PutClazzID(0x794c7ded)

		// set flags
		var getFlags = func() uint32 {
			var flags uint32 = 0

			return flags
		}

		// set flags
		var flags = getFlags()
		x.PutUint32(flags)
		x.PutInt64(m.PermAuthKeyId)
		x.PutInt64(m.AuthKeyId)
		x.PutInt64(m.SessionId)
		if m.Updates == nil {
			return fmt.Errorf("unable to encode gateway_pushSessionUpdatesData#0x794c7ded: field updates is nil")
		}
		if err := m.Updates.Encode(x, layer); err != nil {
			return fmt.Errorf("unable to encode gateway_pushSessionUpdatesData#0x794c7ded: field updates: %w", err)
		}

		return nil
	default:
		return fmt.Errorf("unable to encode gateway_pushSessionUpdatesData: unsupported layer %d", layer)
	}
}

// Decode <--
func (m *TLGatewayPushSessionUpdatesData) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return fmt.Errorf("unable to decode gateway_pushSessionUpdatesData: constructor: %w", err)
		}
	}
	switch m.ClazzID {
	case 0x794c7ded:
		flags, err := d.Uint32()
		if err != nil {
			return fmt.Errorf("unable to decode gateway_pushSessionUpdatesData: field flags: %w", err)
		}
		_ = flags
		m.PermAuthKeyId, err = d.Int64()
		if err != nil {
			return fmt.Errorf("unable to decode gateway_pushSessionUpdatesData#0x794c7ded: field perm_auth_key_id: %w", err)
		}
		m.AuthKeyId, err = d.Int64()
		if err != nil {
			return fmt.Errorf("unable to decode gateway_pushSessionUpdatesData#0x794c7ded: field auth_key_id: %w", err)
		}
		m.SessionId, err = d.Int64()
		if err != nil {
			return fmt.Errorf("unable to decode gateway_pushSessionUpdatesData#0x794c7ded: field session_id: %w", err)
		}

		m.Updates, err = tg.DecodeUpdatesClazz(d)
		if err != nil {
			return fmt.Errorf("unable to decode gateway_pushSessionUpdatesData#0x794c7ded: field updates: %w", err)
		}

		return nil
	default:
		return fmt.Errorf("unable to decode gateway_pushSessionUpdatesData: invalid constructor %x", m.ClazzID)
	}
}

// TLGatewayPushRpcResultData <--
type TLGatewayPushRpcResultData struct {
	ClazzID        uint32 `json:"_id"`
	PermAuthKeyId  int64  `json:"perm_auth_key_id"`
	AuthKeyId      int64  `json:"auth_key_id"`
	SessionId      int64  `json:"session_id"`
	ClientReqMsgId int64  `json:"client_req_msg_id"`
	RpcResultData  []byte `json:"rpc_result_data"`
}

func (m *TLGatewayPushRpcResultData) String() string {
	return iface.DebugStringWithName(ClazzName_gateway_pushRpcResultData, m)
}

// Encode <--
func (m *TLGatewayPushRpcResultData) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_gateway_pushRpcResultData, int(layer)); clazzId {
	case 0xfc960f5:
		x.PutClazzID(0xfc960f5)

		x.PutInt64(m.PermAuthKeyId)
		x.PutInt64(m.AuthKeyId)
		x.PutInt64(m.SessionId)
		x.PutInt64(m.ClientReqMsgId)
		x.PutBytes(m.RpcResultData)

		return nil
	default:
		return fmt.Errorf("unable to encode gateway_pushRpcResultData: unsupported layer %d", layer)
	}
}

// Decode <--
func (m *TLGatewayPushRpcResultData) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return fmt.Errorf("unable to decode gateway_pushRpcResultData: constructor: %w", err)
		}
	}
	switch m.ClazzID {
	case 0xfc960f5:
		m.PermAuthKeyId, err = d.Int64()
		if err != nil {
			return fmt.Errorf("unable to decode gateway_pushRpcResultData#0xfc960f5: field perm_auth_key_id: %w", err)
		}
		m.AuthKeyId, err = d.Int64()
		if err != nil {
			return fmt.Errorf("unable to decode gateway_pushRpcResultData#0xfc960f5: field auth_key_id: %w", err)
		}
		m.SessionId, err = d.Int64()
		if err != nil {
			return fmt.Errorf("unable to decode gateway_pushRpcResultData#0xfc960f5: field session_id: %w", err)
		}
		m.ClientReqMsgId, err = d.Int64()
		if err != nil {
			return fmt.Errorf("unable to decode gateway_pushRpcResultData#0xfc960f5: field client_req_msg_id: %w", err)
		}
		m.RpcResultData, err = d.Bytes()
		if err != nil {
			return fmt.Errorf("unable to decode gateway_pushRpcResultData#0xfc960f5: field rpc_result_data: %w", err)
		}

		return nil
	default:
		return fmt.Errorf("unable to decode gateway_pushRpcResultData: invalid constructor %x", m.ClazzID)
	}
}

// Vector api result type
// ----------------------------------------------------------------------------
// VectorResList <--

// ----------------------------------------------------------------------------
// rpc

type RPCGateway interface {
	GatewayPushUpdatesData(ctx context.Context, in *TLGatewayPushUpdatesData) (*tg.Bool, error)
	GatewayPushSessionUpdatesData(ctx context.Context, in *TLGatewayPushSessionUpdatesData) (*tg.Bool, error)
	GatewayPushRpcResultData(ctx context.Context, in *TLGatewayPushRpcResultData) (*tg.Bool, error)
}
