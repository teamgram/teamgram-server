// Copyright 2024 Teamgooo Authors
//  All rights reserved.
//
// Author: Benqi (wubenqi@gmail.com)
//

package iface

import (
	"encoding/json"
	"fmt"

	"github.com/teamgram/teamgram-server/v2/pkg/proto/bin"
)

//const (
//	MTPROTO_VERSION = 2
//)

// Basic TL types.
const (
	ClazzID_int32     = 0x8ccffa3f // 0x8ccffa3f
	ClazzID_long      = 0x4ab29f6d // 0x4ab29f6d
	ClazzID_int64     = 0xa2813660 // 0xa2813660
	ClazzID_double    = 0x554d59c8 // 0x554d59c8
	ClazzID_string    = 0xb973445  // 0xb973445
	ClazzID_void      = 0x1c084438 // 0x1c084438
	ClazzID_boolFalse = 0xbc799737 // bc799737
	ClazzID_boolTrue  = 0x997275b5 // 997275b5
	CClazzID_true     = 0x3fedd339 // 3fedd339
	ClazzID_vector    = 0x1cb5c415
)

// Size32 represents 4-byte sequence.
// Values in TL are generally aligned to Word.
const Size32 = 4

type TLObject interface {
	Encode(x *bin.Encoder, layer int32) error
	Decode(d *bin.Decoder) error
	// EncodeBare(x *bin.Encoder, layer int32) error
	// DecodeBare(d *bin.Decoder) error
	// ClazzID() uint32
	// ClazzName() string
	// String() string
}

type WithNameWrapper struct {
	ClazzName string `json:"_name,omitempty"`
	TLObject  `json:"_object"`
}

func (m WithNameWrapper) String() string {
	data, _ := json.Marshal(m)
	return string(data)
}

func DecodeObject(d *bin.Decoder) (TLObject, error) {
	clazzID, err := d.ClazzID()
	if err != nil {
		return nil, err
	}

	r := NewTLObjectByClazzID(clazzID)
	if r == nil {
		err = fmt.Errorf("can't find registered classId: %x", clazzID)
		return nil, err
	}

	fmt.Printf("newTLObjectByClassID, classID: %x\n", uint32(clazzID))
	err = r.Decode(d)
	if err != nil {
		err = fmt.Errorf("object(%x) decode error: %v", uint32(clazzID), err)
		return nil, err
	}

	return r, nil
}

func EncodeObject(obj TLObject, layer int32) ([]byte, error) {
	x := bin.NewEncoder()
	if s, ok := obj.(TLObjectSizer); ok {
		x = bin.AcquireEncoderCap(s.CalcSize(layer))
	}

	err := obj.Encode(x, layer)
	if err != nil {
		return nil, err
	}

	return x.Bytes(), nil
}

func EncodeObjectList[T TLObject](x *bin.Encoder, vList []T, layer int32) error {
	x.PutClazzID(ClazzID_vector)
	x.PutInt(len(vList))
	for _, obj := range vList {
		_ = obj.Encode(x, layer)
	}

	return nil
}

func DecodeObjectList[T TLObject](d *bin.Decoder) ([]T, error) {
	if err := d.ConsumeClazzID(ClazzID_vector); err != nil {
		return nil, err
	}
	n, err := d.Int()
	if err != nil {
		return nil, err
	}
	if n < 0 {
		return nil, &bin.InvalidLengthError{
			Type:   "vector",
			Length: n,
		}
	}

	vList := make([]T, n)
	for i := 0; i < n; i++ {
		var (
			obj TLObject
		)

		obj, err = DecodeObject(d)
		if err != nil {
			return nil, err
		}
		vList[i] = obj.(T)
	}

	return vList, nil
}

// EncodeBool serializes bare boolean.
func EncodeBool(x *bin.Encoder, v bool) {
	var (
		c uint32
	)

	switch v {
	case true:
		c = ClazzID_boolTrue
		x.PutClazzID(c)
	case false:
		c = ClazzID_boolFalse
		x.PutClazzID(c)
	}
}

// DecodeBool decodes bare boolean from Buffer.
func DecodeBool(d *bin.Decoder) (bool, error) {
	v, err := d.PeekClazzID()
	if err != nil {
		return false, err
	}
	switch v {
	case ClazzID_boolTrue:
		_ = d.ConsumeClazzID(ClazzID_boolTrue)
		return true, nil
	case ClazzID_boolFalse:
		_ = d.ConsumeClazzID(ClazzID_boolTrue)
		return false, nil
	default:
		return false, bin.NewUnexpectedClazzID(0, v, d.Offset())
	}
}

func EncodeInt32List(x *bin.Encoder, vList []int32) {
	x.PutClazzID(ClazzID_vector)
	x.PutInt(len(vList))
	for _, v := range vList {
		x.PutInt32(v)
	}
}

func DecodeInt32List(d *bin.Decoder) ([]int32, error) {
	if err := d.ConsumeClazzID(ClazzID_vector); err != nil {
		return nil, err
	}
	n, err := d.Int()
	if err != nil {
		return nil, err
	}
	if n < 0 {
		return nil, &bin.InvalidLengthError{
			Type:   "vector",
			Length: n,
		}
	}

	vList := make([]int32, n)
	for i := 0; i < n; i++ {
		vList[i], err = d.Int32()
		if err != nil {
			return nil, err
		}
	}

	return vList, nil
}

func EncodeInt64List(x *bin.Encoder, vList []int64) {
	x.PutClazzID(ClazzID_vector)
	x.PutInt(len(vList))
	for _, v := range vList {
		x.PutInt64(v)
	}
}

func DecodeInt64List(d *bin.Decoder) ([]int64, error) {
	if err := d.ConsumeClazzID(ClazzID_vector); err != nil {
		return nil, err
	}
	n, err := d.Int()
	if err != nil {
		return nil, err
	}
	if n < 0 {
		return nil, &bin.InvalidLengthError{
			Type:   "vector",
			Length: n,
		}
	}

	vList := make([]int64, n)
	for i := 0; i < n; i++ {
		vList[i], err = d.Int64()
		if err != nil {
			return nil, err
		}
	}

	return vList, nil
}

func EncodeStringList(x *bin.Encoder, vList []string) {
	x.PutClazzID(ClazzID_vector)
	x.PutInt(len(vList))
	for _, v := range vList {
		x.PutString(v)
	}
}

func DecodeStringList(d *bin.Decoder) ([]string, error) {
	if err := d.ConsumeClazzID(ClazzID_vector); err != nil {
		return nil, err
	}
	n, err := d.Int()
	if err != nil {
		return nil, err
	}
	if n < 0 {
		return nil, &bin.InvalidLengthError{
			Type:   "vector",
			Length: n,
		}
	}

	vList := make([]string, n)
	for i := 0; i < n; i++ {
		vList[i], err = d.String()
		if err != nil {
			return nil, err
		}
	}

	return vList, nil
}

func EncodeBytesList(x *bin.Encoder, vList [][]byte) {
	x.PutClazzID(ClazzID_vector)
	x.PutInt(len(vList))
	for _, v := range vList {
		x.PutBytes(v)
	}
}

func DecodeBytesList(d *bin.Decoder) ([][]byte, error) {
	if err := d.ConsumeClazzID(ClazzID_vector); err != nil {
		return nil, err
	}
	n, err := d.Int()
	if err != nil {
		return nil, err
	}
	if n < 0 {
		return nil, &bin.InvalidLengthError{
			Type:   "vector",
			Length: n,
		}
	}

	vList := make([][]byte, n)
	for i := 0; i < n; i++ {
		vList[i], err = d.Bytes()
		if err != nil {
			return nil, err
		}
	}

	return vList, nil
}
