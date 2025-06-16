/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2024-present,  Teamgram Authors.
 *  All rights reserved.
 *
 * Author: Benqi (wubenqi@gmail.com)
 */

package idgen

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

// TLIdgenNextId <--
type TLIdgenNextId struct {
	ClazzID uint32 `json:"_id"`
}

// Encode <--
func (m *TLIdgenNextId) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0xbe711020: func() error {
			x.PutClazzID(0xbe711020)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_idgen_nextId, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_idgen_nextId, layer)
	}
}

// Decode <--
func (m *TLIdgenNextId) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0xbe711020: func() (err error) {

			return nil
		},
	}

	if f, ok := decodeF[m.ClazzID]; ok {
		return f()
	} else {
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// TLIdgenNextIds <--
type TLIdgenNextIds struct {
	ClazzID uint32 `json:"_id"`
	Num     int32  `json:"num"`
}

// Encode <--
func (m *TLIdgenNextIds) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0x47c56fae: func() error {
			x.PutClazzID(0x47c56fae)

			x.PutInt32(m.Num)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_idgen_nextIds, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_idgen_nextIds, layer)
	}
}

// Decode <--
func (m *TLIdgenNextIds) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0x47c56fae: func() (err error) {
			m.Num, err = d.Int32()

			return nil
		},
	}

	if f, ok := decodeF[m.ClazzID]; ok {
		return f()
	} else {
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// TLIdgenGetCurrentSeqId <--
type TLIdgenGetCurrentSeqId struct {
	ClazzID uint32 `json:"_id"`
	Key     string `json:"key"`
}

// Encode <--
func (m *TLIdgenGetCurrentSeqId) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0x9d5bab80: func() error {
			x.PutClazzID(0x9d5bab80)

			x.PutString(m.Key)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_idgen_getCurrentSeqId, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_idgen_getCurrentSeqId, layer)
	}
}

// Decode <--
func (m *TLIdgenGetCurrentSeqId) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0x9d5bab80: func() (err error) {
			m.Key, err = d.String()

			return nil
		},
	}

	if f, ok := decodeF[m.ClazzID]; ok {
		return f()
	} else {
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// TLIdgenSetCurrentSeqId <--
type TLIdgenSetCurrentSeqId struct {
	ClazzID uint32 `json:"_id"`
	Key     string `json:"key"`
	Id      int64  `json:"id"`
}

// Encode <--
func (m *TLIdgenSetCurrentSeqId) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0xcd2c196d: func() error {
			x.PutClazzID(0xcd2c196d)

			x.PutString(m.Key)
			x.PutInt64(m.Id)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_idgen_setCurrentSeqId, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_idgen_setCurrentSeqId, layer)
	}
}

// Decode <--
func (m *TLIdgenSetCurrentSeqId) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0xcd2c196d: func() (err error) {
			m.Key, err = d.String()
			m.Id, err = d.Int64()

			return nil
		},
	}

	if f, ok := decodeF[m.ClazzID]; ok {
		return f()
	} else {
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// TLIdgenGetNextSeqId <--
type TLIdgenGetNextSeqId struct {
	ClazzID uint32 `json:"_id"`
	Key     string `json:"key"`
}

// Encode <--
func (m *TLIdgenGetNextSeqId) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0xf6716968: func() error {
			x.PutClazzID(0xf6716968)

			x.PutString(m.Key)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_idgen_getNextSeqId, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_idgen_getNextSeqId, layer)
	}
}

// Decode <--
func (m *TLIdgenGetNextSeqId) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0xf6716968: func() (err error) {
			m.Key, err = d.String()

			return nil
		},
	}

	if f, ok := decodeF[m.ClazzID]; ok {
		return f()
	} else {
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// TLIdgenGetNextNSeqId <--
type TLIdgenGetNextNSeqId struct {
	ClazzID uint32 `json:"_id"`
	Key     string `json:"key"`
	N       int32  `json:"n"`
}

// Encode <--
func (m *TLIdgenGetNextNSeqId) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0xa7d4cc6e: func() error {
			x.PutClazzID(0xa7d4cc6e)

			x.PutString(m.Key)
			x.PutInt32(m.N)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_idgen_getNextNSeqId, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_idgen_getNextNSeqId, layer)
	}
}

// Decode <--
func (m *TLIdgenGetNextNSeqId) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0xa7d4cc6e: func() (err error) {
			m.Key, err = d.String()
			m.N, err = d.Int32()

			return nil
		},
	}

	if f, ok := decodeF[m.ClazzID]; ok {
		return f()
	} else {
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// TLIdgenGetNextIdValList <--
type TLIdgenGetNextIdValList struct {
	ClazzID uint32     `json:"_id"`
	Id      []*InputId `json:"id"`
}

// Encode <--
func (m *TLIdgenGetNextIdValList) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0xaa85f137: func() error {
			x.PutClazzID(0xaa85f137)

			_ = iface.EncodeObjectList(x, m.Id, layer)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_idgen_getNextIdValList, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_idgen_getNextIdValList, layer)
	}
}

// Decode <--
func (m *TLIdgenGetNextIdValList) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0xaa85f137: func() (err error) {
			c1, err2 := d.ClazzID()
			if c1 != iface.ClazzID_vector {
				// dBuf.err = fmt.Errorf("invalid ClazzID_vector, c%d: %d", 1, c1)
				return err2
			}
			l1, err3 := d.Int()
			v1 := make([]*InputId, l1)
			for i := 0; i < l1; i++ {
				vv := new(InputId)
				err3 = vv.Decode(d)
				_ = err3
				v1[i] = vv
			}
			m.Id = v1

			return nil
		},
	}

	if f, ok := decodeF[m.ClazzID]; ok {
		return f()
	} else {
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// TLIdgenGetCurrentSeqIdList <--
type TLIdgenGetCurrentSeqIdList struct {
	ClazzID uint32     `json:"_id"`
	Id      []*InputId `json:"id"`
}

// Encode <--
func (m *TLIdgenGetCurrentSeqIdList) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0xd229ae43: func() error {
			x.PutClazzID(0xd229ae43)

			_ = iface.EncodeObjectList(x, m.Id, layer)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_idgen_getCurrentSeqIdList, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_idgen_getCurrentSeqIdList, layer)
	}
}

// Decode <--
func (m *TLIdgenGetCurrentSeqIdList) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0xd229ae43: func() (err error) {
			c1, err2 := d.ClazzID()
			if c1 != iface.ClazzID_vector {
				// dBuf.err = fmt.Errorf("invalid ClazzID_vector, c%d: %d", 1, c1)
				return err2
			}
			l1, err3 := d.Int()
			v1 := make([]*InputId, l1)
			for i := 0; i < l1; i++ {
				vv := new(InputId)
				err3 = vv.Decode(d)
				_ = err3
				v1[i] = vv
			}
			m.Id = v1

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

// VectorLong <--
type VectorLong struct {
	Datas []int64 `json:"datas"`
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
	Datas []*IdVal `json:"datas"`
}

// Encode <--
func (m *VectorIdVal) Encode(x *bin.Encoder, layer int32) error {
	_ = iface.EncodeObjectList(x, m.Datas, layer)

	return nil
}

// Decode <--
func (m *VectorIdVal) Decode(d *bin.Decoder) (err error) {
	m.Datas, err = iface.DecodeObjectList[*IdVal](d)

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
