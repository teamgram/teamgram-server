/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2025-present,  Teamgram Authors.
 *  All rights reserved.
 *
 * Author: Benqi (wubenqi@gmail.com)
 */

package media

import (
	"encoding/json"
	"fmt"

	"github.com/teamgram/proto/v2/bin"
	"github.com/teamgram/proto/v2/iface"
	"github.com/teamgram/proto/v2/tg"
)

var _ iface.TLObject
var _ fmt.Stringer
var _ *tg.Bool
var _ bin.Fields

// PhotoSizeListClazz <--
//   - TL_PhotoSizeList
type PhotoSizeListClazz interface {
	iface.TLObject
	PhotoSizeListClazzName() string
}

func DecodePhotoSizeListClazz(d *bin.Decoder) (PhotoSizeListClazz, error) {
	// id, err := d.PeekClazzID()
	id, err := d.ClazzID()
	if err != nil {
		return nil, err
	}

	clazzName := iface.GetClazzNameByID(id)
	switch clazzName {
	case ClazzName_photoSizeList:
		x := &TLPhotoSizeList{ClazzID: id, ClazzName2: ClazzName_photoSizeList}
		_ = x.Decode(d)
		return x, nil
	default:
		return nil, fmt.Errorf("DecodePhotoSizeList - unexpected clazzId: %d", id)
	}
}

// TLPhotoSizeList <--
type TLPhotoSizeList struct {
	ClazzID    uint32              `json:"_id"`
	ClazzName2 string              `json:"_name"`
	SizeId     int64               `json:"size_id"`
	Sizes      []tg.PhotoSizeClazz `json:"sizes"`
	DcId       int32               `json:"dc_id"`
}

func MakeTLPhotoSizeList(m *TLPhotoSizeList) *TLPhotoSizeList {
	if m == nil {
		return nil
	}
	m.ClazzName2 = ClazzName_photoSizeList

	return m
}

func (m *TLPhotoSizeList) String() string {
	data, _ := json.Marshal(m)
	return string(data)
}

// PhotoSizeListClazzName <--
func (m *TLPhotoSizeList) PhotoSizeListClazzName() string {
	return ClazzName_photoSizeList
}

// ClazzName <--
func (m *TLPhotoSizeList) ClazzName() string {
	return m.ClazzName2
}

// ToPhotoSizeList <--
func (m *TLPhotoSizeList) ToPhotoSizeList() *PhotoSizeList {
	if m == nil {
		return nil
	}

	return &PhotoSizeList{Clazz: m}
}

// Encode <--
func (m *TLPhotoSizeList) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0x67139b3: func() error {
			x.PutClazzID(0x67139b3)

			x.PutInt64(m.SizeId)

			_ = iface.EncodeObjectList(x, m.Sizes, layer)

			x.PutInt32(m.DcId)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_photoSizeList, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_photoSizeList, layer)
	}
}

// Decode <--
func (m *TLPhotoSizeList) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0x67139b3: func() (err error) {
			m.SizeId, err = d.Int64()
			c1, err2 := d.ClazzID()
			if c1 != iface.ClazzID_vector {
				// dBuf.err = fmt.Errorf("invalid ClazzID_vector, c%d: %d", 1, c1)
				return err2
			}
			l1, err3 := d.Int()
			v1 := make([]tg.PhotoSizeClazz, l1)
			for i := 0; i < l1; i++ {
				// vv := new(PhotoSize)
				// err3 = vv.Decode(d)
				// _ = err3
				// v1[i] = vv
				v1[i], err3 = tg.DecodePhotoSizeClazz(d)
				_ = err3
			}
			m.Sizes = v1

			m.DcId, err = d.Int32()

			return nil
		},
	}

	if f, ok := decodeF[m.ClazzID]; ok {
		return f()
	} else {
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// PhotoSizeList <--
type PhotoSizeList struct {
	// ClazzID   uint32 `json:"_id"`
	// ClazzName string `json:"_name"`
	Clazz PhotoSizeListClazz `json:"_clazz"`
}

func (m *PhotoSizeList) String() string {
	data, _ := json.Marshal(m)
	return string(data)
}

func (m *PhotoSizeList) ClazzName() string {
	if m.Clazz == nil {
		return ""
	} else {
		return m.Clazz.PhotoSizeListClazzName()
	}
}

// Encode <--
func (m *PhotoSizeList) Encode(x *bin.Encoder, layer int32) error {
	if m.Clazz != nil {
		return m.Clazz.Encode(x, layer)
	}

	return fmt.Errorf("PhotoSizeList - invalid Clazz")
}

// Decode <--
func (m *PhotoSizeList) Decode(d *bin.Decoder) (err error) {
	m.Clazz, err = DecodePhotoSizeListClazz(d)
	return
}

// Match <--
func (m *PhotoSizeList) Match(f ...interface{}) {
	if m.Clazz == nil {
		return
	}
	switch c := m.Clazz.(type) {
	case *TLPhotoSizeList:
		for _, v := range f {
			if f1, ok := v.(func(c *TLPhotoSizeList) interface{}); ok {
				f1(c)
			}
		}
	default:
		//
	}
}

// ToPhotoSizeList <--
func (m *PhotoSizeList) ToPhotoSizeList() (*TLPhotoSizeList, bool) {
	if m == nil {
		return nil, false
	}

	if m.Clazz == nil {
		return nil, false
	}

	if x, ok := m.Clazz.(*TLPhotoSizeList); ok {
		return x, true
	}

	return nil, false
}

// VideoSizeListClazz <--
//   - TL_VideoSizeList
type VideoSizeListClazz interface {
	iface.TLObject
	VideoSizeListClazzName() string
}

func DecodeVideoSizeListClazz(d *bin.Decoder) (VideoSizeListClazz, error) {
	// id, err := d.PeekClazzID()
	id, err := d.ClazzID()
	if err != nil {
		return nil, err
	}

	clazzName := iface.GetClazzNameByID(id)
	switch clazzName {
	case ClazzName_videoSizeList:
		x := &TLVideoSizeList{ClazzID: id, ClazzName2: ClazzName_videoSizeList}
		_ = x.Decode(d)
		return x, nil
	default:
		return nil, fmt.Errorf("DecodeVideoSizeList - unexpected clazzId: %d", id)
	}
}

// TLVideoSizeList <--
type TLVideoSizeList struct {
	ClazzID    uint32              `json:"_id"`
	ClazzName2 string              `json:"_name"`
	SizeId     int64               `json:"size_id"`
	Sizes      []tg.VideoSizeClazz `json:"sizes"`
	DcId       int32               `json:"dc_id"`
}

func MakeTLVideoSizeList(m *TLVideoSizeList) *TLVideoSizeList {
	if m == nil {
		return nil
	}
	m.ClazzName2 = ClazzName_videoSizeList

	return m
}

func (m *TLVideoSizeList) String() string {
	data, _ := json.Marshal(m)
	return string(data)
}

// VideoSizeListClazzName <--
func (m *TLVideoSizeList) VideoSizeListClazzName() string {
	return ClazzName_videoSizeList
}

// ClazzName <--
func (m *TLVideoSizeList) ClazzName() string {
	return m.ClazzName2
}

// ToVideoSizeList <--
func (m *TLVideoSizeList) ToVideoSizeList() *VideoSizeList {
	if m == nil {
		return nil
	}

	return &VideoSizeList{Clazz: m}
}

// Encode <--
func (m *TLVideoSizeList) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0x38d19bf2: func() error {
			x.PutClazzID(0x38d19bf2)

			x.PutInt64(m.SizeId)

			_ = iface.EncodeObjectList(x, m.Sizes, layer)

			x.PutInt32(m.DcId)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_videoSizeList, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_videoSizeList, layer)
	}
}

// Decode <--
func (m *TLVideoSizeList) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0x38d19bf2: func() (err error) {
			m.SizeId, err = d.Int64()
			c1, err2 := d.ClazzID()
			if c1 != iface.ClazzID_vector {
				// dBuf.err = fmt.Errorf("invalid ClazzID_vector, c%d: %d", 1, c1)
				return err2
			}
			l1, err3 := d.Int()
			v1 := make([]tg.VideoSizeClazz, l1)
			for i := 0; i < l1; i++ {
				// vv := new(VideoSize)
				// err3 = vv.Decode(d)
				// _ = err3
				// v1[i] = vv
				v1[i], err3 = tg.DecodeVideoSizeClazz(d)
				_ = err3
			}
			m.Sizes = v1

			m.DcId, err = d.Int32()

			return nil
		},
	}

	if f, ok := decodeF[m.ClazzID]; ok {
		return f()
	} else {
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// VideoSizeList <--
type VideoSizeList struct {
	// ClazzID   uint32 `json:"_id"`
	// ClazzName string `json:"_name"`
	Clazz VideoSizeListClazz `json:"_clazz"`
}

func (m *VideoSizeList) String() string {
	data, _ := json.Marshal(m)
	return string(data)
}

func (m *VideoSizeList) ClazzName() string {
	if m.Clazz == nil {
		return ""
	} else {
		return m.Clazz.VideoSizeListClazzName()
	}
}

// Encode <--
func (m *VideoSizeList) Encode(x *bin.Encoder, layer int32) error {
	if m.Clazz != nil {
		return m.Clazz.Encode(x, layer)
	}

	return fmt.Errorf("VideoSizeList - invalid Clazz")
}

// Decode <--
func (m *VideoSizeList) Decode(d *bin.Decoder) (err error) {
	m.Clazz, err = DecodeVideoSizeListClazz(d)
	return
}

// Match <--
func (m *VideoSizeList) Match(f ...interface{}) {
	if m.Clazz == nil {
		return
	}
	switch c := m.Clazz.(type) {
	case *TLVideoSizeList:
		for _, v := range f {
			if f1, ok := v.(func(c *TLVideoSizeList) interface{}); ok {
				f1(c)
			}
		}
	default:
		//
	}
}

// ToVideoSizeList <--
func (m *VideoSizeList) ToVideoSizeList() (*TLVideoSizeList, bool) {
	if m == nil {
		return nil, false
	}

	if m.Clazz == nil {
		return nil, false
	}

	if x, ok := m.Clazz.(*TLVideoSizeList); ok {
		return x, true
	}

	return nil, false
}
