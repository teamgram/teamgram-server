/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2026-present,  Teamgram Authors.
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package media

import (
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

// PhotoSizeListClazz <--
//   - TL_PhotoSizeList
type PhotoSizeListClazz = *TLPhotoSizeList

func DecodePhotoSizeListClazz(d *bin.Decoder) (PhotoSizeListClazz, error) {
	// id, err := d.PeekClazzID()
	id, err := d.ClazzID()
	if err != nil {
		return nil, err
	}

	switch id {
	case 0x67139b3:
		x := &TLPhotoSizeList{ClazzID: id, ClazzName2: ClazzName_photoSizeList}
		if err := x.Decode(d); err != nil {
			return nil, err
		}
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

func (m *TLPhotoSizeList) MarshalJSON() ([]byte, error) {
	if m == nil {
		return []byte("null"), nil
	}
	return iface.MarshalWithName("photoSizeList", m)
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

	return m

}

func (m *TLPhotoSizeList) CalcSize(layer int32) int {
	switch clazzId := iface.GetClazzIDByName(ClazzName_photoSizeList, int(layer)); clazzId {
	case 0x67139b3:
		size := 4
		size += 8
		size += iface.CalcObjectListSize(m.Sizes, layer)
		size += 4

		return size
	default:
		return 0
	}
}

func (m *TLPhotoSizeList) Validate(layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_photoSizeList, int(layer)); clazzId {
	case 0x67139b3:
		if err := iface.ValidateRequiredSlice("sizes", m.Sizes); err != nil {
			return err
		}

		return nil
	default:
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_photoSizeList, layer)
	}
}

// Encode <--
func (m *TLPhotoSizeList) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_photoSizeList, int(layer)); clazzId {
	case 0x67139b3:
		x.PutClazzID(0x67139b3)

		x.PutInt64(m.SizeId)

		if err := iface.EncodeObjectList(x, m.Sizes, layer); err != nil {
			return err
		}

		x.PutInt32(m.DcId)

		return nil
	default:
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_photoSizeList, layer)
	}
}

// Decode <--
func (m *TLPhotoSizeList) Decode(d *bin.Decoder) (err error) {
	switch m.ClazzID {
	case 0x67139b3:
		m.SizeId, err = d.Int64()
		if err != nil {
			return err
		}
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
		v1 := make([]tg.PhotoSizeClazz, l1)
		for i := 0; i < l1; i++ {
			v1[i], err3 = tg.DecodePhotoSizeClazz(d)
			if err3 != nil {
				return err3
			}
		}
		m.Sizes = v1

		m.DcId, err = d.Int32()
		if err != nil {
			return err
		}

		return nil
	default:
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// PhotoSizeList <--
type PhotoSizeList = TLPhotoSizeList

// VideoSizeListClazz <--
//   - TL_VideoSizeList
type VideoSizeListClazz = *TLVideoSizeList

func DecodeVideoSizeListClazz(d *bin.Decoder) (VideoSizeListClazz, error) {
	// id, err := d.PeekClazzID()
	id, err := d.ClazzID()
	if err != nil {
		return nil, err
	}

	switch id {
	case 0x38d19bf2:
		x := &TLVideoSizeList{ClazzID: id, ClazzName2: ClazzName_videoSizeList}
		if err := x.Decode(d); err != nil {
			return nil, err
		}
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

func (m *TLVideoSizeList) MarshalJSON() ([]byte, error) {
	if m == nil {
		return []byte("null"), nil
	}
	return iface.MarshalWithName("videoSizeList", m)
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

	return m

}

func (m *TLVideoSizeList) CalcSize(layer int32) int {
	switch clazzId := iface.GetClazzIDByName(ClazzName_videoSizeList, int(layer)); clazzId {
	case 0x38d19bf2:
		size := 4
		size += 8
		size += iface.CalcObjectListSize(m.Sizes, layer)
		size += 4

		return size
	default:
		return 0
	}
}

func (m *TLVideoSizeList) Validate(layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_videoSizeList, int(layer)); clazzId {
	case 0x38d19bf2:
		if err := iface.ValidateRequiredSlice("sizes", m.Sizes); err != nil {
			return err
		}

		return nil
	default:
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_videoSizeList, layer)
	}
}

// Encode <--
func (m *TLVideoSizeList) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_videoSizeList, int(layer)); clazzId {
	case 0x38d19bf2:
		x.PutClazzID(0x38d19bf2)

		x.PutInt64(m.SizeId)

		if err := iface.EncodeObjectList(x, m.Sizes, layer); err != nil {
			return err
		}

		x.PutInt32(m.DcId)

		return nil
	default:
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_videoSizeList, layer)
	}
}

// Decode <--
func (m *TLVideoSizeList) Decode(d *bin.Decoder) (err error) {
	switch m.ClazzID {
	case 0x38d19bf2:
		m.SizeId, err = d.Int64()
		if err != nil {
			return err
		}
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
		v1 := make([]tg.VideoSizeClazz, l1)
		for i := 0; i < l1; i++ {
			v1[i], err3 = tg.DecodeVideoSizeClazz(d)
			if err3 != nil {
				return err3
			}
		}
		m.Sizes = v1

		m.DcId, err = d.Int32()
		if err != nil {
			return err
		}

		return nil
	default:
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// VideoSizeList <--
type VideoSizeList = TLVideoSizeList
