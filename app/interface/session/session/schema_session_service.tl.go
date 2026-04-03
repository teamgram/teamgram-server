/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2026-present,  Teamgram Authors.
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package session

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

// TLSessionQueryAuthKey <--
type TLSessionQueryAuthKey struct {
	ClazzID   uint32 `json:"_id"`
	AuthKeyId int64  `json:"auth_key_id"`
}

func (m *TLSessionQueryAuthKey) String() string {
	wrapper := iface.WithNameWrapper{ClazzName: ClazzName_session_queryAuthKey, TLObject: m}
	return wrapper.String()
}

// Encode <--
func (m *TLSessionQueryAuthKey) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_session_queryAuthKey, int(layer)); clazzId {
	case 0x6b2df851:
		x.PutClazzID(0x6b2df851)

		x.PutInt64(m.AuthKeyId)

		return nil
	default:
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_session_queryAuthKey, layer)
	}
}

// Decode <--
func (m *TLSessionQueryAuthKey) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return err
		}
	}
	switch m.ClazzID {
	case 0x6b2df851:
		m.AuthKeyId, err = d.Int64()
		if err != nil {
			return err
		}

		return nil
	default:
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// TLSessionSetAuthKey <--
type TLSessionSetAuthKey struct {
	ClazzID    uint32              `json:"_id"`
	AuthKey    tg.AuthKeyInfoClazz `json:"auth_key"`
	FutureSalt tg.FutureSaltClazz  `json:"future_salt"`
	ExpiresIn  int32               `json:"expires_in"`
}

func (m *TLSessionSetAuthKey) String() string {
	wrapper := iface.WithNameWrapper{ClazzName: ClazzName_session_setAuthKey, TLObject: m}
	return wrapper.String()
}

// Encode <--
func (m *TLSessionSetAuthKey) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_session_setAuthKey, int(layer)); clazzId {
	case 0x1d11490b:
		x.PutClazzID(0x1d11490b)

		_ = m.AuthKey.Encode(x, layer)
		_ = m.FutureSalt.Encode(x, layer)
		x.PutInt32(m.ExpiresIn)

		return nil
	default:
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_session_setAuthKey, layer)
	}
}

// Decode <--
func (m *TLSessionSetAuthKey) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return err
		}
	}
	switch m.ClazzID {
	case 0x1d11490b:

		m.AuthKey, err = tg.DecodeAuthKeyInfoClazz(d)
		if err != nil {
			return err
		}

		m.FutureSalt, err = tg.DecodeFutureSaltClazz(d)
		if err != nil {
			return err
		}

		m.ExpiresIn, err = d.Int32()
		if err != nil {
			return err
		}

		return nil
	default:
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// TLSessionCreateSession <--
type TLSessionCreateSession struct {
	ClazzID uint32                  `json:"_id"`
	Client  SessionClientEventClazz `json:"client"`
}

func (m *TLSessionCreateSession) String() string {
	wrapper := iface.WithNameWrapper{ClazzName: ClazzName_session_createSession, TLObject: m}
	return wrapper.String()
}

// Encode <--
func (m *TLSessionCreateSession) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_session_createSession, int(layer)); clazzId {
	case 0x410cb20d:
		x.PutClazzID(0x410cb20d)

		_ = m.Client.Encode(x, layer)

		return nil
	default:
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_session_createSession, layer)
	}
}

// Decode <--
func (m *TLSessionCreateSession) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return err
		}
	}
	switch m.ClazzID {
	case 0x410cb20d:

		m.Client, err = DecodeSessionClientEventClazz(d)
		if err != nil {
			return err
		}

		return nil
	default:
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// TLSessionSendDataToSession <--
type TLSessionSendDataToSession struct {
	ClazzID uint32                 `json:"_id"`
	Data    SessionClientDataClazz `json:"data"`
}

func (m *TLSessionSendDataToSession) String() string {
	wrapper := iface.WithNameWrapper{ClazzName: ClazzName_session_sendDataToSession, TLObject: m}
	return wrapper.String()
}

// Encode <--
func (m *TLSessionSendDataToSession) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_session_sendDataToSession, int(layer)); clazzId {
	case 0x876b2dec:
		x.PutClazzID(0x876b2dec)

		_ = m.Data.Encode(x, layer)

		return nil
	default:
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_session_sendDataToSession, layer)
	}
}

// Decode <--
func (m *TLSessionSendDataToSession) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return err
		}
	}
	switch m.ClazzID {
	case 0x876b2dec:

		m.Data, err = DecodeSessionClientDataClazz(d)
		if err != nil {
			return err
		}

		return nil
	default:
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// TLSessionSendHttpDataToSession <--
type TLSessionSendHttpDataToSession struct {
	ClazzID uint32                 `json:"_id"`
	Client  SessionClientDataClazz `json:"client"`
}

func (m *TLSessionSendHttpDataToSession) String() string {
	wrapper := iface.WithNameWrapper{ClazzName: ClazzName_session_sendHttpDataToSession, TLObject: m}
	return wrapper.String()
}

// Encode <--
func (m *TLSessionSendHttpDataToSession) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_session_sendHttpDataToSession, int(layer)); clazzId {
	case 0xbbec23ae:
		x.PutClazzID(0xbbec23ae)

		_ = m.Client.Encode(x, layer)

		return nil
	default:
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_session_sendHttpDataToSession, layer)
	}
}

// Decode <--
func (m *TLSessionSendHttpDataToSession) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return err
		}
	}
	switch m.ClazzID {
	case 0xbbec23ae:

		m.Client, err = DecodeSessionClientDataClazz(d)
		if err != nil {
			return err
		}

		return nil
	default:
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// TLSessionCloseSession <--
type TLSessionCloseSession struct {
	ClazzID uint32                  `json:"_id"`
	Client  SessionClientEventClazz `json:"client"`
}

func (m *TLSessionCloseSession) String() string {
	wrapper := iface.WithNameWrapper{ClazzName: ClazzName_session_closeSession, TLObject: m}
	return wrapper.String()
}

// Encode <--
func (m *TLSessionCloseSession) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_session_closeSession, int(layer)); clazzId {
	case 0x176fc253:
		x.PutClazzID(0x176fc253)

		_ = m.Client.Encode(x, layer)

		return nil
	default:
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_session_closeSession, layer)
	}
}

// Decode <--
func (m *TLSessionCloseSession) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return err
		}
	}
	switch m.ClazzID {
	case 0x176fc253:

		m.Client, err = DecodeSessionClientEventClazz(d)
		if err != nil {
			return err
		}

		return nil
	default:
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// TLSessionPushUpdatesData <--
type TLSessionPushUpdatesData struct {
	ClazzID       uint32          `json:"_id"`
	PermAuthKeyId int64           `json:"perm_auth_key_id"`
	Notification  bool            `json:"notification"`
	Updates       tg.UpdatesClazz `json:"updates"`
}

func (m *TLSessionPushUpdatesData) String() string {
	wrapper := iface.WithNameWrapper{ClazzName: ClazzName_session_pushUpdatesData, TLObject: m}
	return wrapper.String()
}

// Encode <--
func (m *TLSessionPushUpdatesData) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_session_pushUpdatesData, int(layer)); clazzId {
	case 0xa574d829:
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
	default:
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_session_pushUpdatesData, layer)
	}
}

// Decode <--
func (m *TLSessionPushUpdatesData) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return err
		}
	}
	switch m.ClazzID {
	case 0xa574d829:
		flags, err := d.Uint32()
		if err != nil {
			return err
		}
		_ = flags
		m.PermAuthKeyId, err = d.Int64()
		if err != nil {
			return err
		}
		if (flags & (1 << 0)) != 0 {
			m.Notification = true
		}

		m.Updates, err = tg.DecodeUpdatesClazz(d)
		if err != nil {
			return err
		}

		return nil
	default:
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// TLSessionPushSessionUpdatesData <--
type TLSessionPushSessionUpdatesData struct {
	ClazzID       uint32          `json:"_id"`
	PermAuthKeyId int64           `json:"perm_auth_key_id"`
	AuthKeyId     int64           `json:"auth_key_id"`
	SessionId     int64           `json:"session_id"`
	Updates       tg.UpdatesClazz `json:"updates"`
}

func (m *TLSessionPushSessionUpdatesData) String() string {
	wrapper := iface.WithNameWrapper{ClazzName: ClazzName_session_pushSessionUpdatesData, TLObject: m}
	return wrapper.String()
}

// Encode <--
func (m *TLSessionPushSessionUpdatesData) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_session_pushSessionUpdatesData, int(layer)); clazzId {
	case 0x45f3fda0:
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
	default:
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_session_pushSessionUpdatesData, layer)
	}
}

// Decode <--
func (m *TLSessionPushSessionUpdatesData) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return err
		}
	}
	switch m.ClazzID {
	case 0x45f3fda0:
		flags, err := d.Uint32()
		if err != nil {
			return err
		}
		_ = flags
		m.PermAuthKeyId, err = d.Int64()
		if err != nil {
			return err
		}
		m.AuthKeyId, err = d.Int64()
		if err != nil {
			return err
		}
		m.SessionId, err = d.Int64()
		if err != nil {
			return err
		}

		m.Updates, err = tg.DecodeUpdatesClazz(d)
		if err != nil {
			return err
		}

		return nil
	default:
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

func (m *TLSessionPushRpcResultData) String() string {
	wrapper := iface.WithNameWrapper{ClazzName: ClazzName_session_pushRpcResultData, TLObject: m}
	return wrapper.String()
}

// Encode <--
func (m *TLSessionPushRpcResultData) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_session_pushRpcResultData, int(layer)); clazzId {
	case 0x4b470c89:
		x.PutClazzID(0x4b470c89)

		x.PutInt64(m.PermAuthKeyId)
		x.PutInt64(m.AuthKeyId)
		x.PutInt64(m.SessionId)
		x.PutInt64(m.ClientReqMsgId)
		x.PutBytes(m.RpcResultData)

		return nil
	default:
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_session_pushRpcResultData, layer)
	}
}

// Decode <--
func (m *TLSessionPushRpcResultData) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return err
		}
	}
	switch m.ClazzID {
	case 0x4b470c89:
		m.PermAuthKeyId, err = d.Int64()
		if err != nil {
			return err
		}
		m.AuthKeyId, err = d.Int64()
		if err != nil {
			return err
		}
		m.SessionId, err = d.Int64()
		if err != nil {
			return err
		}
		m.ClientReqMsgId, err = d.Int64()
		if err != nil {
			return err
		}
		m.RpcResultData, err = d.Bytes()
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
