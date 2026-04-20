/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2026-present,  Teamgram Authors.
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package sync

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

// TLSyncUpdatesMe <--
type TLSyncUpdatesMe struct {
	ClazzID       uint32          `json:"_id"`
	UserId        int64           `json:"user_id"`
	PermAuthKeyId int64           `json:"perm_auth_key_id"`
	ServerId      *string         `json:"server_id"`
	AuthKeyId     *int64          `json:"auth_key_id"`
	SessionId     *int64          `json:"session_id"`
	Updates       tg.UpdatesClazz `json:"updates"`
}

func (m *TLSyncUpdatesMe) String() string {
	wrapper := iface.WithNameWrapper{ClazzName: ClazzName_sync_updatesMe, TLObject: m}
	return wrapper.String()
}

// Encode <--
func (m *TLSyncUpdatesMe) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_sync_updatesMe, int(layer)); clazzId {
	case 0xe57d411f:
		x.PutClazzID(0xe57d411f)

		// set flags
		var getFlags = func() uint32 {
			var flags uint32 = 0

			if m.ServerId != nil {
				flags |= 1 << 0
			}
			if m.AuthKeyId != nil {
				flags |= 1 << 1
			}
			if m.SessionId != nil {
				flags |= 1 << 1
			}

			return flags
		}

		// set flags
		var flags = getFlags()
		x.PutUint32(flags)
		x.PutInt64(m.UserId)
		x.PutInt64(m.PermAuthKeyId)
		if m.ServerId != nil {
			x.PutString(*m.ServerId)
		}

		if m.AuthKeyId != nil {
			x.PutInt64(*m.AuthKeyId)
		}

		if m.SessionId != nil {
			x.PutInt64(*m.SessionId)
		}

		_ = m.Updates.Encode(x, layer)

		return nil
	default:
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_sync_updatesMe, layer)
	}
}

// Decode <--
func (m *TLSyncUpdatesMe) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return err
		}
	}
	switch m.ClazzID {
	case 0xe57d411f:
		flags, err := d.Uint32()
		if err != nil {
			return err
		}
		_ = flags
		m.UserId, err = d.Int64()
		if err != nil {
			return err
		}
		m.PermAuthKeyId, err = d.Int64()
		if err != nil {
			return err
		}
		if (flags & (1 << 0)) != 0 {
			m.ServerId = new(string)
			*m.ServerId, err = d.String()
			if err != nil {
				return err
			}
		}

		if (flags & (1 << 1)) != 0 {
			m.AuthKeyId = new(int64)
			*m.AuthKeyId, err = d.Int64()
			if err != nil {
				return err
			}
		}

		if (flags & (1 << 1)) != 0 {
			m.SessionId = new(int64)
			*m.SessionId, err = d.Int64()
			if err != nil {
				return err
			}
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

// TLSyncUpdatesNotMe <--
type TLSyncUpdatesNotMe struct {
	ClazzID       uint32          `json:"_id"`
	UserId        int64           `json:"user_id"`
	PermAuthKeyId int64           `json:"perm_auth_key_id"`
	Updates       tg.UpdatesClazz `json:"updates"`
}

func (m *TLSyncUpdatesNotMe) String() string {
	wrapper := iface.WithNameWrapper{ClazzName: ClazzName_sync_updatesNotMe, TLObject: m}
	return wrapper.String()
}

// Encode <--
func (m *TLSyncUpdatesNotMe) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_sync_updatesNotMe, int(layer)); clazzId {
	case 0x97ac5031:
		x.PutClazzID(0x97ac5031)

		x.PutInt64(m.UserId)
		x.PutInt64(m.PermAuthKeyId)
		_ = m.Updates.Encode(x, layer)

		return nil
	default:
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_sync_updatesNotMe, layer)
	}
}

// Decode <--
func (m *TLSyncUpdatesNotMe) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return err
		}
	}
	switch m.ClazzID {
	case 0x97ac5031:
		m.UserId, err = d.Int64()
		if err != nil {
			return err
		}
		m.PermAuthKeyId, err = d.Int64()
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

// TLSyncPushUpdates <--
type TLSyncPushUpdates struct {
	ClazzID uint32          `json:"_id"`
	UserId  int64           `json:"user_id"`
	Updates tg.UpdatesClazz `json:"updates"`
}

func (m *TLSyncPushUpdates) String() string {
	wrapper := iface.WithNameWrapper{ClazzName: ClazzName_sync_pushUpdates, TLObject: m}
	return wrapper.String()
}

// Encode <--
func (m *TLSyncPushUpdates) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_sync_pushUpdates, int(layer)); clazzId {
	case 0x8f0ad9be:
		x.PutClazzID(0x8f0ad9be)

		x.PutInt64(m.UserId)
		_ = m.Updates.Encode(x, layer)

		return nil
	default:
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_sync_pushUpdates, layer)
	}
}

// Decode <--
func (m *TLSyncPushUpdates) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return err
		}
	}
	switch m.ClazzID {
	case 0x8f0ad9be:
		m.UserId, err = d.Int64()
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

// TLSyncPushUpdatesIfNot <--
type TLSyncPushUpdatesIfNot struct {
	ClazzID  uint32          `json:"_id"`
	UserId   int64           `json:"user_id"`
	Includes []int64         `json:"includes"`
	Excludes []int64         `json:"excludes"`
	Updates  tg.UpdatesClazz `json:"updates"`
}

func (m *TLSyncPushUpdatesIfNot) String() string {
	wrapper := iface.WithNameWrapper{ClazzName: ClazzName_sync_pushUpdatesIfNot, TLObject: m}
	return wrapper.String()
}

// Encode <--
func (m *TLSyncPushUpdatesIfNot) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_sync_pushUpdatesIfNot, int(layer)); clazzId {
	case 0x2d3778bc:
		x.PutClazzID(0x2d3778bc)

		// set flags
		var getFlags = func() uint32 {
			var flags uint32 = 0

			if m.Includes != nil {
				flags |= 1 << 0
			}
			if m.Excludes != nil {
				flags |= 1 << 1
			}

			return flags
		}

		// set flags
		var flags = getFlags()
		x.PutUint32(flags)
		x.PutInt64(m.UserId)
		if m.Includes != nil {
			iface.EncodeInt64List(x, m.Includes)
		}
		if m.Excludes != nil {
			iface.EncodeInt64List(x, m.Excludes)
		}
		_ = m.Updates.Encode(x, layer)

		return nil
	default:
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_sync_pushUpdatesIfNot, layer)
	}
}

// Decode <--
func (m *TLSyncPushUpdatesIfNot) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return err
		}
	}
	switch m.ClazzID {
	case 0x2d3778bc:
		flags, err := d.Uint32()
		if err != nil {
			return err
		}
		_ = flags
		m.UserId, err = d.Int64()
		if err != nil {
			return err
		}
		if (flags & (1 << 0)) != 0 {
			m.Includes, err = iface.DecodeInt64List(d)
			if err != nil {
				return err
			}
		}
		if (flags & (1 << 1)) != 0 {
			m.Excludes, err = iface.DecodeInt64List(d)
			if err != nil {
				return err
			}
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

// TLSyncPushBotUpdates <--
type TLSyncPushBotUpdates struct {
	ClazzID uint32          `json:"_id"`
	UserId  int64           `json:"user_id"`
	Updates tg.UpdatesClazz `json:"updates"`
}

func (m *TLSyncPushBotUpdates) String() string {
	wrapper := iface.WithNameWrapper{ClazzName: ClazzName_sync_pushBotUpdates, TLObject: m}
	return wrapper.String()
}

// Encode <--
func (m *TLSyncPushBotUpdates) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_sync_pushBotUpdates, int(layer)); clazzId {
	case 0xadc3f000:
		x.PutClazzID(0xadc3f000)

		x.PutInt64(m.UserId)
		_ = m.Updates.Encode(x, layer)

		return nil
	default:
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_sync_pushBotUpdates, layer)
	}
}

// Decode <--
func (m *TLSyncPushBotUpdates) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return err
		}
	}
	switch m.ClazzID {
	case 0xadc3f000:
		m.UserId, err = d.Int64()
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

// TLSyncPushRpcResult <--
type TLSyncPushRpcResult struct {
	ClazzID        uint32 `json:"_id"`
	UserId         int64  `json:"user_id"`
	AuthKeyId      int64  `json:"auth_key_id"`
	PermAuthKeyId  int64  `json:"perm_auth_key_id"`
	ServerId       string `json:"server_id"`
	SessionId      int64  `json:"session_id"`
	ClientReqMsgId int64  `json:"client_req_msg_id"`
	RpcResult      []byte `json:"rpc_result"`
}

func (m *TLSyncPushRpcResult) String() string {
	wrapper := iface.WithNameWrapper{ClazzName: ClazzName_sync_pushRpcResult, TLObject: m}
	return wrapper.String()
}

// Encode <--
func (m *TLSyncPushRpcResult) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_sync_pushRpcResult, int(layer)); clazzId {
	case 0x1a9d4b2:
		x.PutClazzID(0x1a9d4b2)

		x.PutInt64(m.UserId)
		x.PutInt64(m.AuthKeyId)
		x.PutInt64(m.PermAuthKeyId)
		x.PutString(m.ServerId)
		x.PutInt64(m.SessionId)
		x.PutInt64(m.ClientReqMsgId)
		x.PutBytes(m.RpcResult)

		return nil
	default:
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_sync_pushRpcResult, layer)
	}
}

// Decode <--
func (m *TLSyncPushRpcResult) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return err
		}
	}
	switch m.ClazzID {
	case 0x1a9d4b2:
		m.UserId, err = d.Int64()
		if err != nil {
			return err
		}
		m.AuthKeyId, err = d.Int64()
		if err != nil {
			return err
		}
		m.PermAuthKeyId, err = d.Int64()
		if err != nil {
			return err
		}
		m.ServerId, err = d.String()
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
		m.RpcResult, err = d.Bytes()
		if err != nil {
			return err
		}

		return nil
	default:
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// TLSyncBroadcastUpdates <--
type TLSyncBroadcastUpdates struct {
	ClazzID       uint32          `json:"_id"`
	BroadcastType int32           `json:"broadcast_type"`
	ChatId        int64           `json:"chat_id"`
	ExcludeIdList []int64         `json:"exclude_id_list"`
	Updates       tg.UpdatesClazz `json:"updates"`
}

func (m *TLSyncBroadcastUpdates) String() string {
	wrapper := iface.WithNameWrapper{ClazzName: ClazzName_sync_broadcastUpdates, TLObject: m}
	return wrapper.String()
}

// Encode <--
func (m *TLSyncBroadcastUpdates) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_sync_broadcastUpdates, int(layer)); clazzId {
	case 0xf5e35cb6:
		x.PutClazzID(0xf5e35cb6)

		x.PutInt32(m.BroadcastType)
		x.PutInt64(m.ChatId)

		iface.EncodeInt64List(x, m.ExcludeIdList)

		_ = m.Updates.Encode(x, layer)

		return nil
	default:
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_sync_broadcastUpdates, layer)
	}
}

// Decode <--
func (m *TLSyncBroadcastUpdates) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return err
		}
	}
	switch m.ClazzID {
	case 0xf5e35cb6:
		m.BroadcastType, err = d.Int32()
		if err != nil {
			return err
		}
		m.ChatId, err = d.Int64()
		if err != nil {
			return err
		}

		m.ExcludeIdList, err = iface.DecodeInt64List(d)
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

// Vector api result type
// ----------------------------------------------------------------------------
// VectorResList <--

// ----------------------------------------------------------------------------
// rpc

type RPCSync interface {
	SyncUpdatesMe(ctx context.Context, in *TLSyncUpdatesMe) (*tg.Void, error)
	SyncUpdatesNotMe(ctx context.Context, in *TLSyncUpdatesNotMe) (*tg.Void, error)
	SyncPushUpdates(ctx context.Context, in *TLSyncPushUpdates) (*tg.Void, error)
	SyncPushUpdatesIfNot(ctx context.Context, in *TLSyncPushUpdatesIfNot) (*tg.Void, error)
	SyncPushBotUpdates(ctx context.Context, in *TLSyncPushBotUpdates) (*tg.Void, error)
	SyncPushRpcResult(ctx context.Context, in *TLSyncPushRpcResult) (*tg.Void, error)
	SyncBroadcastUpdates(ctx context.Context, in *TLSyncBroadcastUpdates) (*tg.Void, error)
}
