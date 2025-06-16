/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2024-present,  Teamgram Authors.
 *  All rights reserved.
 *
 * Author: Benqi (wubenqi@gmail.com)
 */

package session

import (
	"context"
	"fmt"

	"github.com/teamgram/proto/v2/bin"
	"github.com/teamgram/proto/v2/iface"
	"github.com/teamgram/proto/v2/mt"
	"github.com/teamgram/proto/v2/tg"
)

var _ iface.TLObject
var _ fmt.Stringer
var _ *tg.Bool
var _ bin.Fields

// TLSessionQueryAuthKey <--
type TLSessionQueryAuthKey struct {
	ClazzID   uint32 `json:"_id"`
	AuthKeyId int64  `json:"auth_key_id"`
}

// Encode <--
func (m *TLSessionQueryAuthKey) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0x6b2df851: func() error {
			x.PutClazzID(0x6b2df851)

			x.PutInt64(m.AuthKeyId)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_session_queryAuthKey, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_session_queryAuthKey, layer)
	}
}

// Decode <--
func (m *TLSessionQueryAuthKey) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0x6b2df851: func() (err error) {
			m.AuthKeyId, err = d.Int64()

			return nil
		},
	}

	if f, ok := decodeF[m.ClazzID]; ok {
		return f()
	} else {
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// TLSessionSetAuthKey <--
type TLSessionSetAuthKey struct {
	ClazzID    uint32          `json:"_id"`
	AuthKey    *tg.AuthKeyInfo `json:"auth_key"`
	FutureSalt *mt.FutureSalt  `json:"future_salt"`
	ExpiresIn  int32           `json:"expires_in"`
}

// Encode <--
func (m *TLSessionSetAuthKey) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0x1d11490b: func() error {
			x.PutClazzID(0x1d11490b)

			_ = m.AuthKey.Encode(x, layer)
			_ = m.FutureSalt.Encode(x, layer)
			x.PutInt32(m.ExpiresIn)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_session_setAuthKey, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_session_setAuthKey, layer)
	}
}

// Decode <--
func (m *TLSessionSetAuthKey) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0x1d11490b: func() (err error) {

			m1 := &tg.AuthKeyInfo{}
			_ = m1.Decode(d)
			m.AuthKey = m1

			m2 := &mt.FutureSalt{}
			_ = m2.Decode(d)
			m.FutureSalt = m2

			m.ExpiresIn, err = d.Int32()

			return nil
		},
	}

	if f, ok := decodeF[m.ClazzID]; ok {
		return f()
	} else {
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// TLSessionCreateSession <--
type TLSessionCreateSession struct {
	ClazzID uint32              `json:"_id"`
	Client  *SessionClientEvent `json:"client"`
}

// Encode <--
func (m *TLSessionCreateSession) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0x410cb20d: func() error {
			x.PutClazzID(0x410cb20d)

			_ = m.Client.Encode(x, layer)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_session_createSession, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_session_createSession, layer)
	}
}

// Decode <--
func (m *TLSessionCreateSession) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0x410cb20d: func() (err error) {

			m1 := &SessionClientEvent{}
			_ = m1.Decode(d)
			m.Client = m1

			return nil
		},
	}

	if f, ok := decodeF[m.ClazzID]; ok {
		return f()
	} else {
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// TLSessionSendDataToSession <--
type TLSessionSendDataToSession struct {
	ClazzID uint32             `json:"_id"`
	Data    *SessionClientData `json:"data"`
}

// Encode <--
func (m *TLSessionSendDataToSession) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0x876b2dec: func() error {
			x.PutClazzID(0x876b2dec)

			_ = m.Data.Encode(x, layer)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_session_sendDataToSession, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_session_sendDataToSession, layer)
	}
}

// Decode <--
func (m *TLSessionSendDataToSession) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0x876b2dec: func() (err error) {

			m1 := &SessionClientData{}
			_ = m1.Decode(d)
			m.Data = m1

			return nil
		},
	}

	if f, ok := decodeF[m.ClazzID]; ok {
		return f()
	} else {
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// TLSessionSendHttpDataToSession <--
type TLSessionSendHttpDataToSession struct {
	ClazzID uint32             `json:"_id"`
	Client  *SessionClientData `json:"client"`
}

// Encode <--
func (m *TLSessionSendHttpDataToSession) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0xbbec23ae: func() error {
			x.PutClazzID(0xbbec23ae)

			_ = m.Client.Encode(x, layer)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_session_sendHttpDataToSession, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_session_sendHttpDataToSession, layer)
	}
}

// Decode <--
func (m *TLSessionSendHttpDataToSession) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0xbbec23ae: func() (err error) {

			m1 := &SessionClientData{}
			_ = m1.Decode(d)
			m.Client = m1

			return nil
		},
	}

	if f, ok := decodeF[m.ClazzID]; ok {
		return f()
	} else {
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// TLSessionCloseSession <--
type TLSessionCloseSession struct {
	ClazzID uint32              `json:"_id"`
	Client  *SessionClientEvent `json:"client"`
}

// Encode <--
func (m *TLSessionCloseSession) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0x176fc253: func() error {
			x.PutClazzID(0x176fc253)

			_ = m.Client.Encode(x, layer)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_session_closeSession, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_session_closeSession, layer)
	}
}

// Decode <--
func (m *TLSessionCloseSession) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0x176fc253: func() (err error) {

			m1 := &SessionClientEvent{}
			_ = m1.Decode(d)
			m.Client = m1

			return nil
		},
	}

	if f, ok := decodeF[m.ClazzID]; ok {
		return f()
	} else {
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// TLSessionPushUpdatesData <--
type TLSessionPushUpdatesData struct {
	ClazzID       uint32      `json:"_id"`
	PermAuthKeyId int64       `json:"perm_auth_key_id"`
	Notification  bool        `json:"notification"`
	Updates       *tg.Updates `json:"updates"`
}

// Encode <--
func (m *TLSessionPushUpdatesData) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0xa574d829: func() error {
			x.PutClazzID(0xa574d829)

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
			_ = m.Updates.Encode(x, layer)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_session_pushUpdatesData, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_session_pushUpdatesData, layer)
	}
}

// Decode <--
func (m *TLSessionPushUpdatesData) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0xa574d829: func() (err error) {
			flags, _ := d.Uint32()
			_ = flags
			m.PermAuthKeyId, err = d.Int64()
			if (flags & (1 << 0)) != 0 {
				m.Notification = true
			}

			m4 := &tg.Updates{}
			_ = m4.Decode(d)
			m.Updates = m4

			return nil
		},
	}

	if f, ok := decodeF[m.ClazzID]; ok {
		return f()
	} else {
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// TLSessionPushSessionUpdatesData <--
type TLSessionPushSessionUpdatesData struct {
	ClazzID       uint32      `json:"_id"`
	PermAuthKeyId int64       `json:"perm_auth_key_id"`
	AuthKeyId     int64       `json:"auth_key_id"`
	SessionId     int64       `json:"session_id"`
	Updates       *tg.Updates `json:"updates"`
}

// Encode <--
func (m *TLSessionPushSessionUpdatesData) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0x45f3fda0: func() error {
			x.PutClazzID(0x45f3fda0)

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
			_ = m.Updates.Encode(x, layer)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_session_pushSessionUpdatesData, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_session_pushSessionUpdatesData, layer)
	}
}

// Decode <--
func (m *TLSessionPushSessionUpdatesData) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0x45f3fda0: func() (err error) {
			flags, _ := d.Uint32()
			_ = flags
			m.PermAuthKeyId, err = d.Int64()
			m.AuthKeyId, err = d.Int64()
			m.SessionId, err = d.Int64()

			m5 := &tg.Updates{}
			_ = m5.Decode(d)
			m.Updates = m5

			return nil
		},
	}

	if f, ok := decodeF[m.ClazzID]; ok {
		return f()
	} else {
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// TLSessionPushRpcResultData <--
type TLSessionPushRpcResultData struct {
	ClazzID        uint32 `json:"_id"`
	PermAuthKeyId  int64  `json:"perm_auth_key_id"`
	AuthKeyId      int64  `json:"auth_key_id"`
	SessionId      int64  `json:"session_id"`
	ClientReqMsgId int64  `json:"client_req_msg_id"`
	RpcResultData  []byte `json:"rpc_result_data"`
}

// Encode <--
func (m *TLSessionPushRpcResultData) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0x4b470c89: func() error {
			x.PutClazzID(0x4b470c89)

			x.PutInt64(m.PermAuthKeyId)
			x.PutInt64(m.AuthKeyId)
			x.PutInt64(m.SessionId)
			x.PutInt64(m.ClientReqMsgId)
			x.PutBytes(m.RpcResultData)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_session_pushRpcResultData, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_session_pushRpcResultData, layer)
	}
}

// Decode <--
func (m *TLSessionPushRpcResultData) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0x4b470c89: func() (err error) {
			m.PermAuthKeyId, err = d.Int64()
			m.AuthKeyId, err = d.Int64()
			m.SessionId, err = d.Int64()
			m.ClientReqMsgId, err = d.Int64()
			m.RpcResultData, err = d.Bytes()

			return nil
		},
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

type RPCSession interface {
	SessionQueryAuthKey(ctx context.Context, in *TLSessionQueryAuthKey) (*tg.AuthKeyInfo, error)
	SessionSetAuthKey(ctx context.Context, in *TLSessionSetAuthKey) (*tg.Bool, error)
	SessionCreateSession(ctx context.Context, in *TLSessionCreateSession) (*tg.Bool, error)
	SessionSendDataToSession(ctx context.Context, in *TLSessionSendDataToSession) (*tg.Bool, error)
	SessionSendHttpDataToSession(ctx context.Context, in *TLSessionSendHttpDataToSession) (*HttpSessionData, error)
	SessionCloseSession(ctx context.Context, in *TLSessionCloseSession) (*tg.Bool, error)
	SessionPushUpdatesData(ctx context.Context, in *TLSessionPushUpdatesData) (*tg.Bool, error)
	SessionPushSessionUpdatesData(ctx context.Context, in *TLSessionPushSessionUpdatesData) (*tg.Bool, error)
	SessionPushRpcResultData(ctx context.Context, in *TLSessionPushRpcResultData) (*tg.Bool, error)
}
