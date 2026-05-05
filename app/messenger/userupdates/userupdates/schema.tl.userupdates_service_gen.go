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

// TLUserupdatesListDialogs <--
type TLUserupdatesListDialogs struct {
	ClazzID        uint32 `json:"_id"`
	UserId         int64  `json:"user_id"`
	TopMessageDate int64  `json:"top_message_date"`
	TopPeerSeq     int64  `json:"top_peer_seq"`
	PeerType       int32  `json:"peer_type"`
	PeerId         int64  `json:"peer_id"`
	Limit          int32  `json:"limit"`
}

func (m *TLUserupdatesListDialogs) String() string {
	return iface.DebugStringWithName(ClazzName_userupdates_listDialogs, m)
}

// Encode <--
func (m *TLUserupdatesListDialogs) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_userupdates_listDialogs, int(layer)); clazzId {
	case 0x53638fcc:
		x.PutClazzID(0x53638fcc)

		x.PutInt64(m.UserId)
		x.PutInt64(m.TopMessageDate)
		x.PutInt64(m.TopPeerSeq)
		x.PutInt32(m.PeerType)
		x.PutInt64(m.PeerId)
		x.PutInt32(m.Limit)

		return nil
	default:
		return fmt.Errorf("unable to encode userupdates_listDialogs: unsupported layer %d", layer)
	}
}

// Decode <--
func (m *TLUserupdatesListDialogs) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return fmt.Errorf("unable to decode userupdates_listDialogs: constructor: %w", err)
		}
	}
	switch m.ClazzID {
	case 0x53638fcc:
		m.UserId, err = d.Int64()
		if err != nil {
			return fmt.Errorf("unable to decode userupdates_listDialogs#0x53638fcc: field user_id: %w", err)
		}
		m.TopMessageDate, err = d.Int64()
		if err != nil {
			return fmt.Errorf("unable to decode userupdates_listDialogs#0x53638fcc: field top_message_date: %w", err)
		}
		m.TopPeerSeq, err = d.Int64()
		if err != nil {
			return fmt.Errorf("unable to decode userupdates_listDialogs#0x53638fcc: field top_peer_seq: %w", err)
		}
		m.PeerType, err = d.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode userupdates_listDialogs#0x53638fcc: field peer_type: %w", err)
		}
		m.PeerId, err = d.Int64()
		if err != nil {
			return fmt.Errorf("unable to decode userupdates_listDialogs#0x53638fcc: field peer_id: %w", err)
		}
		m.Limit, err = d.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode userupdates_listDialogs#0x53638fcc: field limit: %w", err)
		}

		return nil
	default:
		return fmt.Errorf("unable to decode userupdates_listDialogs: invalid constructor %x", m.ClazzID)
	}
}

// TLUserupdatesGetDialogsByPeers <--
type TLUserupdatesGetDialogsByPeers struct {
	ClazzID uint32                      `json:"_id"`
	UserId  int64                       `json:"user_id"`
	Peers   []DialogProjectionPeerClazz `json:"peers"`
}

func (m *TLUserupdatesGetDialogsByPeers) String() string {
	return iface.DebugStringWithName(ClazzName_userupdates_getDialogsByPeers, m)
}

// Encode <--
func (m *TLUserupdatesGetDialogsByPeers) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_userupdates_getDialogsByPeers, int(layer)); clazzId {
	case 0xc6a9626f:
		x.PutClazzID(0xc6a9626f)

		x.PutInt64(m.UserId)

		if err := iface.EncodeObjectList(x, m.Peers, layer); err != nil {
			return fmt.Errorf("unable to encode userupdates_getDialogsByPeers#0xc6a9626f: field peers: %w", err)
		}

		return nil
	default:
		return fmt.Errorf("unable to encode userupdates_getDialogsByPeers: unsupported layer %d", layer)
	}
}

// Decode <--
func (m *TLUserupdatesGetDialogsByPeers) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return fmt.Errorf("unable to decode userupdates_getDialogsByPeers: constructor: %w", err)
		}
	}
	switch m.ClazzID {
	case 0xc6a9626f:
		m.UserId, err = d.Int64()
		if err != nil {
			return fmt.Errorf("unable to decode userupdates_getDialogsByPeers#0xc6a9626f: field user_id: %w", err)
		}
		l2, err3 := d.VectorHeader()
		if err3 != nil {
			return fmt.Errorf("unable to decode userupdates_getDialogsByPeers#0xc6a9626f: field peers: %w", err3)
		}
		if l2 > bin.MaxVectorLen {
			return fmt.Errorf("unable to decode userupdates_getDialogsByPeers#0xc6a9626f: field peers: %w", &bin.InvalidLengthError{Type: "vector", Length: int(l2)})
		}
		prealloc2 := int(l2)
		if prealloc2 > bin.PreallocateLimit {
			prealloc2 = bin.PreallocateLimit
		}
		v2 := make([]DialogProjectionPeerClazz, 0, prealloc2)
		for i := int32(0); i < l2; i++ {
			vv2, err3 := DecodeDialogProjectionPeerClazz(d)
			if err3 != nil {
				return fmt.Errorf("unable to decode userupdates_getDialogsByPeers#0xc6a9626f: field peers: %w", err3)
			}
			v2 = append(v2, vv2)
		}
		m.Peers = v2

		return nil
	default:
		return fmt.Errorf("unable to decode userupdates_getDialogsByPeers: invalid constructor %x", m.ClazzID)
	}
}

// TLUserupdatesGetDialogCount <--
type TLUserupdatesGetDialogCount struct {
	ClazzID uint32 `json:"_id"`
	UserId  int64  `json:"user_id"`
}

func (m *TLUserupdatesGetDialogCount) String() string {
	return iface.DebugStringWithName(ClazzName_userupdates_getDialogCount, m)
}

// Encode <--
func (m *TLUserupdatesGetDialogCount) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_userupdates_getDialogCount, int(layer)); clazzId {
	case 0x12060b16:
		x.PutClazzID(0x12060b16)

		x.PutInt64(m.UserId)

		return nil
	default:
		return fmt.Errorf("unable to encode userupdates_getDialogCount: unsupported layer %d", layer)
	}
}

// Decode <--
func (m *TLUserupdatesGetDialogCount) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return fmt.Errorf("unable to decode userupdates_getDialogCount: constructor: %w", err)
		}
	}
	switch m.ClazzID {
	case 0x12060b16:
		m.UserId, err = d.Int64()
		if err != nil {
			return fmt.Errorf("unable to decode userupdates_getDialogCount#0x12060b16: field user_id: %w", err)
		}

		return nil
	default:
		return fmt.Errorf("unable to decode userupdates_getDialogCount: invalid constructor %x", m.ClazzID)
	}
}

// TLUserupdatesAppendDialogAuthSeqSideEffect <--
type TLUserupdatesAppendDialogAuthSeqSideEffect struct {
	ClazzID              uint32 `json:"_id"`
	UserId               int64  `json:"user_id"`
	SourcePermAuthKeyId  int64  `json:"source_perm_auth_key_id"`
	OperationId          string `json:"operation_id"`
	TargetAuthPolicy     string `json:"target_auth_policy"`
	PublicUpdateType     string `json:"public_update_type"`
	PeerType             int32  `json:"peer_type"`
	PeerId               int64  `json:"peer_id"`
	PayloadSchemaVersion int32  `json:"payload_schema_version"`
	Payload              []byte `json:"payload"`
	PayloadHash          []byte `json:"payload_hash"`
}

func (m *TLUserupdatesAppendDialogAuthSeqSideEffect) String() string {
	return iface.DebugStringWithName(ClazzName_userupdates_appendDialogAuthSeqSideEffect, m)
}

// Encode <--
func (m *TLUserupdatesAppendDialogAuthSeqSideEffect) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_userupdates_appendDialogAuthSeqSideEffect, int(layer)); clazzId {
	case 0x170844e5:
		x.PutClazzID(0x170844e5)

		// set flags
		var getFlags = func() uint32 {
			var flags uint32 = 0

			return flags
		}

		// set flags
		var flags = getFlags()
		x.PutUint32(flags)
		x.PutInt64(m.UserId)
		x.PutInt64(m.SourcePermAuthKeyId)
		x.PutString(m.OperationId)
		x.PutString(m.TargetAuthPolicy)
		x.PutString(m.PublicUpdateType)
		x.PutInt32(m.PeerType)
		x.PutInt64(m.PeerId)
		x.PutInt32(m.PayloadSchemaVersion)
		x.PutBytes(m.Payload)
		x.PutBytes(m.PayloadHash)

		return nil
	default:
		return fmt.Errorf("unable to encode userupdates_appendDialogAuthSeqSideEffect: unsupported layer %d", layer)
	}
}

// Decode <--
func (m *TLUserupdatesAppendDialogAuthSeqSideEffect) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return fmt.Errorf("unable to decode userupdates_appendDialogAuthSeqSideEffect: constructor: %w", err)
		}
	}
	switch m.ClazzID {
	case 0x170844e5:
		flags, err := d.Uint32()
		if err != nil {
			return fmt.Errorf("unable to decode userupdates_appendDialogAuthSeqSideEffect: field flags: %w", err)
		}
		_ = flags
		m.UserId, err = d.Int64()
		if err != nil {
			return fmt.Errorf("unable to decode userupdates_appendDialogAuthSeqSideEffect#0x170844e5: field user_id: %w", err)
		}
		m.SourcePermAuthKeyId, err = d.Int64()
		if err != nil {
			return fmt.Errorf("unable to decode userupdates_appendDialogAuthSeqSideEffect#0x170844e5: field source_perm_auth_key_id: %w", err)
		}
		m.OperationId, err = d.String()
		if err != nil {
			return fmt.Errorf("unable to decode userupdates_appendDialogAuthSeqSideEffect#0x170844e5: field operation_id: %w", err)
		}
		m.TargetAuthPolicy, err = d.String()
		if err != nil {
			return fmt.Errorf("unable to decode userupdates_appendDialogAuthSeqSideEffect#0x170844e5: field target_auth_policy: %w", err)
		}
		m.PublicUpdateType, err = d.String()
		if err != nil {
			return fmt.Errorf("unable to decode userupdates_appendDialogAuthSeqSideEffect#0x170844e5: field public_update_type: %w", err)
		}
		m.PeerType, err = d.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode userupdates_appendDialogAuthSeqSideEffect#0x170844e5: field peer_type: %w", err)
		}
		m.PeerId, err = d.Int64()
		if err != nil {
			return fmt.Errorf("unable to decode userupdates_appendDialogAuthSeqSideEffect#0x170844e5: field peer_id: %w", err)
		}
		m.PayloadSchemaVersion, err = d.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode userupdates_appendDialogAuthSeqSideEffect#0x170844e5: field payload_schema_version: %w", err)
		}
		m.Payload, err = d.Bytes()
		if err != nil {
			return fmt.Errorf("unable to decode userupdates_appendDialogAuthSeqSideEffect#0x170844e5: field payload: %w", err)
		}
		m.PayloadHash, err = d.Bytes()
		if err != nil {
			return fmt.Errorf("unable to decode userupdates_appendDialogAuthSeqSideEffect#0x170844e5: field payload_hash: %w", err)
		}

		return nil
	default:
		return fmt.Errorf("unable to decode userupdates_appendDialogAuthSeqSideEffect: invalid constructor %x", m.ClazzID)
	}
}

// TLUserupdatesAppendDialogPtsSideEffect <--
type TLUserupdatesAppendDialogPtsSideEffect struct {
	ClazzID              uint32 `json:"_id"`
	UserId               int64  `json:"user_id"`
	SourcePermAuthKeyId  int64  `json:"source_perm_auth_key_id"`
	OperationId          string `json:"operation_id"`
	TargetAuthPolicy     string `json:"target_auth_policy"`
	PublicUpdateType     string `json:"public_update_type"`
	PeerType             int32  `json:"peer_type"`
	PeerId               int64  `json:"peer_id"`
	PayloadSchemaVersion int32  `json:"payload_schema_version"`
	Payload              []byte `json:"payload"`
	PayloadHash          []byte `json:"payload_hash"`
}

func (m *TLUserupdatesAppendDialogPtsSideEffect) String() string {
	return iface.DebugStringWithName(ClazzName_userupdates_appendDialogPtsSideEffect, m)
}

// Encode <--
func (m *TLUserupdatesAppendDialogPtsSideEffect) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_userupdates_appendDialogPtsSideEffect, int(layer)); clazzId {
	case 0xe93427fd:
		x.PutClazzID(0xe93427fd)

		// set flags
		var getFlags = func() uint32 {
			var flags uint32 = 0

			return flags
		}

		// set flags
		var flags = getFlags()
		x.PutUint32(flags)
		x.PutInt64(m.UserId)
		x.PutInt64(m.SourcePermAuthKeyId)
		x.PutString(m.OperationId)
		x.PutString(m.TargetAuthPolicy)
		x.PutString(m.PublicUpdateType)
		x.PutInt32(m.PeerType)
		x.PutInt64(m.PeerId)
		x.PutInt32(m.PayloadSchemaVersion)
		x.PutBytes(m.Payload)
		x.PutBytes(m.PayloadHash)

		return nil
	default:
		return fmt.Errorf("unable to encode userupdates_appendDialogPtsSideEffect: unsupported layer %d", layer)
	}
}

// Decode <--
func (m *TLUserupdatesAppendDialogPtsSideEffect) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return fmt.Errorf("unable to decode userupdates_appendDialogPtsSideEffect: constructor: %w", err)
		}
	}
	switch m.ClazzID {
	case 0xe93427fd:
		flags, err := d.Uint32()
		if err != nil {
			return fmt.Errorf("unable to decode userupdates_appendDialogPtsSideEffect: field flags: %w", err)
		}
		_ = flags
		m.UserId, err = d.Int64()
		if err != nil {
			return fmt.Errorf("unable to decode userupdates_appendDialogPtsSideEffect#0xe93427fd: field user_id: %w", err)
		}
		m.SourcePermAuthKeyId, err = d.Int64()
		if err != nil {
			return fmt.Errorf("unable to decode userupdates_appendDialogPtsSideEffect#0xe93427fd: field source_perm_auth_key_id: %w", err)
		}
		m.OperationId, err = d.String()
		if err != nil {
			return fmt.Errorf("unable to decode userupdates_appendDialogPtsSideEffect#0xe93427fd: field operation_id: %w", err)
		}
		m.TargetAuthPolicy, err = d.String()
		if err != nil {
			return fmt.Errorf("unable to decode userupdates_appendDialogPtsSideEffect#0xe93427fd: field target_auth_policy: %w", err)
		}
		m.PublicUpdateType, err = d.String()
		if err != nil {
			return fmt.Errorf("unable to decode userupdates_appendDialogPtsSideEffect#0xe93427fd: field public_update_type: %w", err)
		}
		m.PeerType, err = d.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode userupdates_appendDialogPtsSideEffect#0xe93427fd: field peer_type: %w", err)
		}
		m.PeerId, err = d.Int64()
		if err != nil {
			return fmt.Errorf("unable to decode userupdates_appendDialogPtsSideEffect#0xe93427fd: field peer_id: %w", err)
		}
		m.PayloadSchemaVersion, err = d.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode userupdates_appendDialogPtsSideEffect#0xe93427fd: field payload_schema_version: %w", err)
		}
		m.Payload, err = d.Bytes()
		if err != nil {
			return fmt.Errorf("unable to decode userupdates_appendDialogPtsSideEffect#0xe93427fd: field payload: %w", err)
		}
		m.PayloadHash, err = d.Bytes()
		if err != nil {
			return fmt.Errorf("unable to decode userupdates_appendDialogPtsSideEffect#0xe93427fd: field payload_hash: %w", err)
		}

		return nil
	default:
		return fmt.Errorf("unable to decode userupdates_appendDialogPtsSideEffect: invalid constructor %x", m.ClazzID)
	}
}

// Vector api result type
// ----------------------------------------------------------------------------
// VectorResList <--

// VectorDialogProjection <--
type VectorDialogProjection struct {
	Datas []DialogProjectionClazz `json:"_datas"`
}

func (m *VectorDialogProjection) String() string {
	data, _ := json.Marshal(m)
	return string(data)
}

// Encode <--
func (m *VectorDialogProjection) Encode(x *bin.Encoder, layer int32) error {
	_ = iface.EncodeObjectList(x, m.Datas, layer)

	return nil
}

// Decode <--
func (m *VectorDialogProjection) Decode(d *bin.Decoder) (err error) {
	m.Datas, err = iface.DecodeObjectList[DialogProjectionClazz](d)

	return err
}

// ----------------------------------------------------------------------------
// rpc

type RPCUserupdates interface {
	UserupdatesProcessUserOperation(ctx context.Context, in *TLUserupdatesProcessUserOperation) (*UserOperationResult, error)
	UserupdatesGetOperationResult(ctx context.Context, in *TLUserupdatesGetOperationResult) (*UserOperationResult, error)
	UserupdatesGetState(ctx context.Context, in *TLUserupdatesGetState) (*UserState, error)
	UserupdatesGetDifference(ctx context.Context, in *TLUserupdatesGetDifference) (*UserDifference, error)
	UserupdatesListDialogs(ctx context.Context, in *TLUserupdatesListDialogs) (*DialogProjectionList, error)
	UserupdatesGetDialogsByPeers(ctx context.Context, in *TLUserupdatesGetDialogsByPeers) (*VectorDialogProjection, error)
	UserupdatesGetDialogCount(ctx context.Context, in *TLUserupdatesGetDialogCount) (*tg.Int32, error)
	UserupdatesAppendDialogAuthSeqSideEffect(ctx context.Context, in *TLUserupdatesAppendDialogAuthSeqSideEffect) (*UserAuthSeqAppendResult, error)
	UserupdatesAppendDialogPtsSideEffect(ctx context.Context, in *TLUserupdatesAppendDialogPtsSideEffect) (*UserPtsAppendResult, error)
}
