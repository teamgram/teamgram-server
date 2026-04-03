/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2026-present,  Teamgram Authors.
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package idgen

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

// TLIdgenNextId <--
type TLIdgenNextId struct {
	ClazzID uint32 `json:"_id"`
}

func (m *TLIdgenNextId) String() string {
	wrapper := iface.WithNameWrapper{ClazzName: ClazzName_idgen_nextId, TLObject: m}
	return wrapper.String()
}

// Encode <--
func (m *TLIdgenNextId) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_idgen_nextId, int(layer)); clazzId {
	case 0xbe711020:
		x.PutClazzID(0xbe711020)

		return nil
	default:
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_idgen_nextId, layer)
	}
}

// Decode <--
func (m *TLIdgenNextId) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return err
		}
	}
	switch m.ClazzID {
	case 0xbe711020:

		return nil
	default:
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// TLIdgenNextIds <--
type TLIdgenNextIds struct {
	ClazzID uint32 `json:"_id"`
	Num     int32  `json:"num"`
}

func (m *TLIdgenNextIds) String() string {
	wrapper := iface.WithNameWrapper{ClazzName: ClazzName_idgen_nextIds, TLObject: m}
	return wrapper.String()
}

// Encode <--
func (m *TLIdgenNextIds) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_idgen_nextIds, int(layer)); clazzId {
	case 0x47c56fae:
		x.PutClazzID(0x47c56fae)

		x.PutInt32(m.Num)

		return nil
	default:
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_idgen_nextIds, layer)
	}
}

// Decode <--
func (m *TLIdgenNextIds) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return err
		}
	}
	switch m.ClazzID {
	case 0x47c56fae:
		m.Num, err = d.Int32()
		if err != nil {
			return err
		}

		return nil
	default:
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// TLIdgenGetCurrentSeqId <--
type TLIdgenGetCurrentSeqId struct {
	ClazzID uint32 `json:"_id"`
	Key     string `json:"key"`
}

func (m *TLIdgenGetCurrentSeqId) String() string {
	wrapper := iface.WithNameWrapper{ClazzName: ClazzName_idgen_getCurrentSeqId, TLObject: m}
	return wrapper.String()
}

// Encode <--
func (m *TLIdgenGetCurrentSeqId) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_idgen_getCurrentSeqId, int(layer)); clazzId {
	case 0x9d5bab80:
		x.PutClazzID(0x9d5bab80)

		x.PutString(m.Key)

		return nil
	default:
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_idgen_getCurrentSeqId, layer)
	}
}

// Decode <--
func (m *TLIdgenGetCurrentSeqId) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return err
		}
	}
	switch m.ClazzID {
	case 0x9d5bab80:
		m.Key, err = d.String()
		if err != nil {
			return err
		}

		return nil
	default:
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// TLIdgenSetCurrentSeqId <--
type TLIdgenSetCurrentSeqId struct {
	ClazzID uint32 `json:"_id"`
	Key     string `json:"key"`
	Id      int64  `json:"id"`
}

func (m *TLIdgenSetCurrentSeqId) String() string {
	wrapper := iface.WithNameWrapper{ClazzName: ClazzName_idgen_setCurrentSeqId, TLObject: m}
	return wrapper.String()
}

// Encode <--
func (m *TLIdgenSetCurrentSeqId) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_idgen_setCurrentSeqId, int(layer)); clazzId {
	case 0xcd2c196d:
		x.PutClazzID(0xcd2c196d)

		x.PutString(m.Key)
		x.PutInt64(m.Id)

		return nil
	default:
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_idgen_setCurrentSeqId, layer)
	}
}

// Decode <--
func (m *TLIdgenSetCurrentSeqId) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return err
		}
	}
	switch m.ClazzID {
	case 0xcd2c196d:
		m.Key, err = d.String()
		if err != nil {
			return err
		}
		m.Id, err = d.Int64()
		if err != nil {
			return err
		}

		return nil
	default:
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// TLIdgenGetNextSeqId <--
type TLIdgenGetNextSeqId struct {
	ClazzID uint32 `json:"_id"`
	Key     string `json:"key"`
}

func (m *TLIdgenGetNextSeqId) String() string {
	wrapper := iface.WithNameWrapper{ClazzName: ClazzName_idgen_getNextSeqId, TLObject: m}
	return wrapper.String()
}

// Encode <--
func (m *TLIdgenGetNextSeqId) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_idgen_getNextSeqId, int(layer)); clazzId {
	case 0xf6716968:
		x.PutClazzID(0xf6716968)

		x.PutString(m.Key)

		return nil
	default:
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_idgen_getNextSeqId, layer)
	}
}

// Decode <--
func (m *TLIdgenGetNextSeqId) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return err
		}
	}
	switch m.ClazzID {
	case 0xf6716968:
		m.Key, err = d.String()
		if err != nil {
			return err
		}

		return nil
	default:
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// TLIdgenGetNextNSeqId <--
type TLIdgenGetNextNSeqId struct {
	ClazzID uint32 `json:"_id"`
	Key     string `json:"key"`
	N       int32  `json:"n"`
}

func (m *TLIdgenGetNextNSeqId) String() string {
	wrapper := iface.WithNameWrapper{ClazzName: ClazzName_idgen_getNextNSeqId, TLObject: m}
	return wrapper.String()
}

// Encode <--
func (m *TLIdgenGetNextNSeqId) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_idgen_getNextNSeqId, int(layer)); clazzId {
	case 0xa7d4cc6e:
		x.PutClazzID(0xa7d4cc6e)

		x.PutString(m.Key)
		x.PutInt32(m.N)

		return nil
	default:
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_idgen_getNextNSeqId, layer)
	}
}

// Decode <--
func (m *TLIdgenGetNextNSeqId) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return err
		}
	}
	switch m.ClazzID {
	case 0xa7d4cc6e:
		m.Key, err = d.String()
		if err != nil {
			return err
		}
		m.N, err = d.Int32()
		if err != nil {
			return err
		}

		return nil
	default:
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// TLIdgenGetNextIdValList <--
type TLIdgenGetNextIdValList struct {
	ClazzID uint32         `json:"_id"`
	Id      []InputIdClazz `json:"id"`
}

func (m *TLIdgenGetNextIdValList) String() string {
	wrapper := iface.WithNameWrapper{ClazzName: ClazzName_idgen_getNextIdValList, TLObject: m}
	return wrapper.String()
}

// Encode <--
func (m *TLIdgenGetNextIdValList) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_idgen_getNextIdValList, int(layer)); clazzId {
	case 0xaa85f137:
		x.PutClazzID(0xaa85f137)

		if err := iface.EncodeObjectList(x, m.Id, layer); err != nil {
			return err
		}

		return nil
	default:
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_idgen_getNextIdValList, layer)
	}
}

// Decode <--
func (m *TLIdgenGetNextIdValList) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return err
		}
	}
	switch m.ClazzID {
	case 0xaa85f137:
		c1, err2 := d.ClazzID()
		if err2 != nil {
			return err2
		}
		if c1 != iface.ClazzID_vector {
			return fmt.Errorf("invalid ClazzID_vector, c%d: %d", 1, c1)
		}
		l1, err3 := d.Int()
		if err3 != nil {
			return err3
		}
		v1 := make([]InputIdClazz, l1)
		for i := 0; i < l1; i++ {
			v1[i], err3 = DecodeInputIdClazz(d)
			if err3 != nil {
				return err3
			}
		}
		m.Id = v1

		return nil
	default:
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// TLIdgenGetCurrentSeqIdList <--
type TLIdgenGetCurrentSeqIdList struct {
	ClazzID uint32         `json:"_id"`
	Id      []InputIdClazz `json:"id"`
}

func (m *TLIdgenGetCurrentSeqIdList) String() string {
	wrapper := iface.WithNameWrapper{ClazzName: ClazzName_idgen_getCurrentSeqIdList, TLObject: m}
	return wrapper.String()
}

// Encode <--
func (m *TLIdgenGetCurrentSeqIdList) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_idgen_getCurrentSeqIdList, int(layer)); clazzId {
	case 0xd229ae43:
		x.PutClazzID(0xd229ae43)

		if err := iface.EncodeObjectList(x, m.Id, layer); err != nil {
			return err
		}

		return nil
	default:
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_idgen_getCurrentSeqIdList, layer)
	}
}

// Decode <--
func (m *TLIdgenGetCurrentSeqIdList) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return err
		}
	}
	switch m.ClazzID {
	case 0xd229ae43:
		c1, err2 := d.ClazzID()
		if err2 != nil {
			return err2
		}
		if c1 != iface.ClazzID_vector {
			return fmt.Errorf("invalid ClazzID_vector, c%d: %d", 1, c1)
		}
		l1, err3 := d.Int()
		if err3 != nil {
			return err3
		}
		v1 := make([]InputIdClazz, l1)
		for i := 0; i < l1; i++ {
			v1[i], err3 = DecodeInputIdClazz(d)
			if err3 != nil {
				return err3
			}
		}
		m.Id = v1

		return nil
	default:
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// Vector api result type
// ----------------------------------------------------------------------------
// VectorResList <--

// VectorLong <--
type VectorLong struct {
	Datas []int64 `json:"_datas"`
}

func (m *VectorLong) String() string {
	data, _ := json.Marshal(m)
	return string(data)
}

// Encode <--
func (m *VectorLong) Encode(x *bin.Encoder, layer int32) error {
	iface.EncodeInt64List(x, m.Datas)

	return nil
}

// Decode <--
func (m *VectorLong) Decode(d *bin.Decoder) (err error) {
	m.Datas, err = iface.DecodeInt64List(d)

	return err
}

// VectorIdVal <--
type VectorIdVal struct {
	Datas []IdValClazz `json:"_datas"`
}

func (m *VectorIdVal) String() string {
	data, _ := json.Marshal(m)
	return string(data)
}

// Encode <--
func (m *VectorIdVal) Encode(x *bin.Encoder, layer int32) error {
	_ = iface.EncodeObjectList(x, m.Datas, layer)

	return nil
}

// Decode <--
func (m *VectorIdVal) Decode(d *bin.Decoder) (err error) {
	m.Datas, err = iface.DecodeObjectList[IdValClazz](d)

	return err
}

// ----------------------------------------------------------------------------
// rpc

type RPCIdgen interface {
	IdgenNextId(ctx context.Context, in *TLIdgenNextId) (*tg.Int64, error)
	IdgenNextIds(ctx context.Context, in *TLIdgenNextIds) (*VectorLong, error)
	IdgenGetCurrentSeqId(ctx context.Context, in *TLIdgenGetCurrentSeqId) (*tg.Int64, error)
	IdgenSetCurrentSeqId(ctx context.Context, in *TLIdgenSetCurrentSeqId) (*tg.Bool, error)
	IdgenGetNextSeqId(ctx context.Context, in *TLIdgenGetNextSeqId) (*tg.Int64, error)
	IdgenGetNextNSeqId(ctx context.Context, in *TLIdgenGetNextNSeqId) (*tg.Int64, error)
	IdgenGetNextIdValList(ctx context.Context, in *TLIdgenGetNextIdValList) (*VectorIdVal, error)
	IdgenGetCurrentSeqIdList(ctx context.Context, in *TLIdgenGetCurrentSeqIdList) (*VectorIdVal, error)
}
