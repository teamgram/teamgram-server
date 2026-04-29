/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2026-present, The Teamgram Authors (https://teamgram.net).
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package userupdates

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

// TLUserupdatesProcessUserOperation <--
type TLUserupdatesProcessUserOperation struct {
	ClazzID   uint32             `json:"_id"`
	Operation UserOperationClazz `json:"operation"`
}

func (m *TLUserupdatesProcessUserOperation) String() string {
	return iface.DebugStringWithName(ClazzName_userupdates_processUserOperation, m)
}

// Encode <--
func (m *TLUserupdatesProcessUserOperation) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_userupdates_processUserOperation, int(layer)); clazzId {
	case 0xc200ea59:
		x.PutClazzID(0xc200ea59)

		if m.Operation == nil {
			return fmt.Errorf("unable to encode userupdates_processUserOperation#0xc200ea59: field operation is nil")
		}
		if err := m.Operation.Encode(x, layer); err != nil {
			return fmt.Errorf("unable to encode userupdates_processUserOperation#0xc200ea59: field operation: %w", err)
		}

		return nil
	default:
		return fmt.Errorf("unable to encode userupdates_processUserOperation: unsupported layer %d", layer)
	}
}

// Decode <--
func (m *TLUserupdatesProcessUserOperation) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return fmt.Errorf("unable to decode userupdates_processUserOperation: constructor: %w", err)
		}
	}
	switch m.ClazzID {
	case 0xc200ea59:

		m.Operation, err = DecodeUserOperationClazz(d)
		if err != nil {
			return fmt.Errorf("unable to decode userupdates_processUserOperation#0xc200ea59: field operation: %w", err)
		}

		return nil
	default:
		return fmt.Errorf("unable to decode userupdates_processUserOperation: invalid constructor %x", m.ClazzID)
	}
}

// TLUserupdatesGetOperationResult <--
type TLUserupdatesGetOperationResult struct {
	ClazzID     uint32 `json:"_id"`
	UserId      int64  `json:"user_id"`
	OperationId string `json:"operation_id"`
	PayloadHash []byte `json:"payload_hash"`
}

func (m *TLUserupdatesGetOperationResult) String() string {
	return iface.DebugStringWithName(ClazzName_userupdates_getOperationResult, m)
}

// Encode <--
func (m *TLUserupdatesGetOperationResult) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_userupdates_getOperationResult, int(layer)); clazzId {
	case 0x47a995d1:
		x.PutClazzID(0x47a995d1)

		x.PutInt64(m.UserId)
		x.PutString(m.OperationId)
		x.PutBytes(m.PayloadHash)

		return nil
	default:
		return fmt.Errorf("unable to encode userupdates_getOperationResult: unsupported layer %d", layer)
	}
}

// Decode <--
func (m *TLUserupdatesGetOperationResult) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return fmt.Errorf("unable to decode userupdates_getOperationResult: constructor: %w", err)
		}
	}
	switch m.ClazzID {
	case 0x47a995d1:
		m.UserId, err = d.Int64()
		if err != nil {
			return fmt.Errorf("unable to decode userupdates_getOperationResult#0x47a995d1: field user_id: %w", err)
		}
		m.OperationId, err = d.String()
		if err != nil {
			return fmt.Errorf("unable to decode userupdates_getOperationResult#0x47a995d1: field operation_id: %w", err)
		}
		m.PayloadHash, err = d.Bytes()
		if err != nil {
			return fmt.Errorf("unable to decode userupdates_getOperationResult#0x47a995d1: field payload_hash: %w", err)
		}

		return nil
	default:
		return fmt.Errorf("unable to decode userupdates_getOperationResult: invalid constructor %x", m.ClazzID)
	}
}

// TLUserupdatesGetState <--
type TLUserupdatesGetState struct {
	ClazzID   uint32 `json:"_id"`
	UserId    int64  `json:"user_id"`
	AuthKeyId int64  `json:"auth_key_id"`
}

func (m *TLUserupdatesGetState) String() string {
	return iface.DebugStringWithName(ClazzName_userupdates_getState, m)
}

// Encode <--
func (m *TLUserupdatesGetState) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_userupdates_getState, int(layer)); clazzId {
	case 0x3bbbad80:
		x.PutClazzID(0x3bbbad80)

		x.PutInt64(m.UserId)
		x.PutInt64(m.AuthKeyId)

		return nil
	default:
		return fmt.Errorf("unable to encode userupdates_getState: unsupported layer %d", layer)
	}
}

// Decode <--
func (m *TLUserupdatesGetState) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return fmt.Errorf("unable to decode userupdates_getState: constructor: %w", err)
		}
	}
	switch m.ClazzID {
	case 0x3bbbad80:
		m.UserId, err = d.Int64()
		if err != nil {
			return fmt.Errorf("unable to decode userupdates_getState#0x3bbbad80: field user_id: %w", err)
		}
		m.AuthKeyId, err = d.Int64()
		if err != nil {
			return fmt.Errorf("unable to decode userupdates_getState#0x3bbbad80: field auth_key_id: %w", err)
		}

		return nil
	default:
		return fmt.Errorf("unable to decode userupdates_getState: invalid constructor %x", m.ClazzID)
	}
}

// TLUserupdatesGetDifference <--
type TLUserupdatesGetDifference struct {
	ClazzID       uint32 `json:"_id"`
	UserId        int64  `json:"user_id"`
	AuthKeyId     int64  `json:"auth_key_id"`
	Pts           int64  `json:"pts"`
	PtsTotalLimit *int32 `json:"pts_total_limit"`
	Date          *int64 `json:"date"`
}

func (m *TLUserupdatesGetDifference) String() string {
	return iface.DebugStringWithName(ClazzName_userupdates_getDifference, m)
}

// Encode <--
func (m *TLUserupdatesGetDifference) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_userupdates_getDifference, int(layer)); clazzId {
	case 0x38cdd9fc:
		x.PutClazzID(0x38cdd9fc)

		// set flags
		var getFlags = func() uint32 {
			var flags uint32 = 0

			if m.PtsTotalLimit != nil {
				flags |= 1 << 0
			}
			if m.Date != nil {
				flags |= 1 << 1
			}

			return flags
		}

		// set flags
		var flags = getFlags()
		x.PutUint32(flags)
		x.PutInt64(m.UserId)
		x.PutInt64(m.AuthKeyId)
		x.PutInt64(m.Pts)
		if m.PtsTotalLimit != nil {
			x.PutInt32(*m.PtsTotalLimit)
		}

		if m.Date != nil {
			x.PutInt64(*m.Date)
		}

		return nil
	default:
		return fmt.Errorf("unable to encode userupdates_getDifference: unsupported layer %d", layer)
	}
}

// Decode <--
func (m *TLUserupdatesGetDifference) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return fmt.Errorf("unable to decode userupdates_getDifference: constructor: %w", err)
		}
	}
	switch m.ClazzID {
	case 0x38cdd9fc:
		flags, err := d.Uint32()
		if err != nil {
			return fmt.Errorf("unable to decode userupdates_getDifference: field flags: %w", err)
		}
		_ = flags
		m.UserId, err = d.Int64()
		if err != nil {
			return fmt.Errorf("unable to decode userupdates_getDifference#0x38cdd9fc: field user_id: %w", err)
		}
		m.AuthKeyId, err = d.Int64()
		if err != nil {
			return fmt.Errorf("unable to decode userupdates_getDifference#0x38cdd9fc: field auth_key_id: %w", err)
		}
		m.Pts, err = d.Int64()
		if err != nil {
			return fmt.Errorf("unable to decode userupdates_getDifference#0x38cdd9fc: field pts: %w", err)
		}
		if (flags & (1 << 0)) != 0 {
			m.PtsTotalLimit = new(int32)
			*m.PtsTotalLimit, err = d.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode userupdates_getDifference#0x38cdd9fc: field pts_total_limit: %w", err)
			}
		}
		if (flags & (1 << 1)) != 0 {
			m.Date = new(int64)
			*m.Date, err = d.Int64()
			if err != nil {
				return fmt.Errorf("unable to decode userupdates_getDifference#0x38cdd9fc: field date: %w", err)
			}
		}

		return nil
	default:
		return fmt.Errorf("unable to decode userupdates_getDifference: invalid constructor %x", m.ClazzID)
	}
}

// Vector api result type
// ----------------------------------------------------------------------------
// VectorResList <--

// ----------------------------------------------------------------------------
// rpc

type RPCUserupdates interface {
	UserupdatesProcessUserOperation(ctx context.Context, in *TLUserupdatesProcessUserOperation) (*UserOperationResult, error)
	UserupdatesGetOperationResult(ctx context.Context, in *TLUserupdatesGetOperationResult) (*UserOperationResult, error)
	UserupdatesGetState(ctx context.Context, in *TLUserupdatesGetState) (*UserState, error)
	UserupdatesGetDifference(ctx context.Context, in *TLUserupdatesGetDifference) (*UserDifference, error)
}
