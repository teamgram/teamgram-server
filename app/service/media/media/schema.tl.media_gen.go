/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2026-present, The Teamgram Authors (https://teamgram.net).
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

// MediaResolvedFileObjectClazz <--
//   - TL_MediaResolvedFileObject
type MediaResolvedFileObjectClazz = *TLMediaResolvedFileObject

func DecodeMediaResolvedFileObjectClazz(d *bin.Decoder) (MediaResolvedFileObjectClazz, error) {
	// id, err := d.PeekClazzID()
	id, err := d.ClazzID()
	if err != nil {
		return nil, fmt.Errorf("unable to decode MediaResolvedFileObject: constructor: %w", err)
	}

	switch id {
	case 0x986d9e66:
		x := &TLMediaResolvedFileObject{ClazzID: id, ClazzName2: ClazzName_mediaResolvedFileObject}
		if err := x.Decode(d); err != nil {
			return nil, err
		}
		return x, nil
	default:
		return nil, fmt.Errorf("unable to decode MediaResolvedFileObject: invalid constructor %x", id)
	}

}

// TLMediaResolvedFileObject <--
type TLMediaResolvedFileObject struct {
	ClazzID         uint32 `json:"_id"`
	ClazzName2      string `json:"_name"`
	ObjectId        string `json:"object_id"`
	ReadLease       []byte `json:"read_lease"`
	Size2           int64  `json:"size2"`
	MimeType        string `json:"mime_type"`
	DcId            int32  `json:"dc_id"`
	StorageFileType int32  `json:"storage_file_type"`
}

func MakeTLMediaResolvedFileObject(m *TLMediaResolvedFileObject) *TLMediaResolvedFileObject {
	if m == nil {
		return nil
	}
	m.ClazzName2 = ClazzName_mediaResolvedFileObject

	return m
}

func (m *TLMediaResolvedFileObject) String() string {
	return iface.DebugStringWithName("mediaResolvedFileObject", m)
}

func (m *TLMediaResolvedFileObject) MarshalJSON() ([]byte, error) {
	if m == nil {
		return []byte("null"), nil
	}
	return iface.MarshalWithName("mediaResolvedFileObject", m)
}

// MediaResolvedFileObjectClazzName <--
func (m *TLMediaResolvedFileObject) MediaResolvedFileObjectClazzName() string {
	return ClazzName_mediaResolvedFileObject
}

// ClazzName <--
func (m *TLMediaResolvedFileObject) ClazzName() string {
	return m.ClazzName2
}

// ToMediaResolvedFileObject <--
func (m *TLMediaResolvedFileObject) ToMediaResolvedFileObject() *MediaResolvedFileObject {
	if m == nil {
		return nil
	}

	return m

}

func (m *TLMediaResolvedFileObject) CalcSize(layer int32) int {
	switch clazzId := iface.GetClazzIDByName(ClazzName_mediaResolvedFileObject, int(layer)); clazzId {
	case 0x986d9e66:
		size := 4
		size += iface.CalcStringSize(m.ObjectId)
		size += iface.CalcBytesSize(m.ReadLease)
		size += 8
		size += iface.CalcStringSize(m.MimeType)
		size += 4
		size += 4

		return size
	default:
		return 0
	}
}

func (m *TLMediaResolvedFileObject) Validate(layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_mediaResolvedFileObject, int(layer)); clazzId {
	case 0x986d9e66:
		if err := iface.ValidateRequiredString("object_id", m.ObjectId); err != nil {
			return err
		}

		if err := iface.ValidateRequiredBytes("read_lease", m.ReadLease); err != nil {
			return err
		}

		if err := iface.ValidateRequiredString("mime_type", m.MimeType); err != nil {
			return err
		}

		return nil
	default:
		return fmt.Errorf("unable to encode mediaResolvedFileObject: unsupported layer %d", layer)
	}
}

// Encode <--
func (m *TLMediaResolvedFileObject) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_mediaResolvedFileObject, int(layer)); clazzId {
	case 0x986d9e66:
		x.PutClazzID(0x986d9e66)

		x.PutString(m.ObjectId)
		x.PutBytes(m.ReadLease)
		x.PutInt64(m.Size2)
		x.PutString(m.MimeType)
		x.PutInt32(m.DcId)
		x.PutInt32(m.StorageFileType)

		return nil
	default:
		return fmt.Errorf("unable to encode mediaResolvedFileObject: unsupported layer %d", layer)
	}
}

// Decode <--
func (m *TLMediaResolvedFileObject) Decode(d *bin.Decoder) (err error) {
	switch m.ClazzID {
	case 0x986d9e66:
		m.ObjectId, err = d.String()
		if err != nil {
			return fmt.Errorf("unable to decode mediaResolvedFileObject#0x986d9e66: field object_id: %w", err)
		}
		m.ReadLease, err = d.Bytes()
		if err != nil {
			return fmt.Errorf("unable to decode mediaResolvedFileObject#0x986d9e66: field read_lease: %w", err)
		}
		m.Size2, err = d.Int64()
		if err != nil {
			return fmt.Errorf("unable to decode mediaResolvedFileObject#0x986d9e66: field size2: %w", err)
		}
		m.MimeType, err = d.String()
		if err != nil {
			return fmt.Errorf("unable to decode mediaResolvedFileObject#0x986d9e66: field mime_type: %w", err)
		}
		m.DcId, err = d.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode mediaResolvedFileObject#0x986d9e66: field dc_id: %w", err)
		}
		m.StorageFileType, err = d.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode mediaResolvedFileObject#0x986d9e66: field storage_file_type: %w", err)
		}

		return nil
	default:
		return fmt.Errorf("unable to decode mediaResolvedFileObject: invalid constructor %x", m.ClazzID)
	}
}

// MediaResolvedFileObject <--
type MediaResolvedFileObject = TLMediaResolvedFileObject

// PhotoSizeListClazz <--
//   - TL_PhotoSizeList
type PhotoSizeListClazz = *TLPhotoSizeList

func DecodePhotoSizeListClazz(d *bin.Decoder) (PhotoSizeListClazz, error) {
	// id, err := d.PeekClazzID()
	id, err := d.ClazzID()
	if err != nil {
		return nil, fmt.Errorf("unable to decode PhotoSizeList: constructor: %w", err)
	}

	switch id {
	case 0x67139b3:
		x := &TLPhotoSizeList{ClazzID: id, ClazzName2: ClazzName_photoSizeList}
		if err := x.Decode(d); err != nil {
			return nil, err
		}
		return x, nil
	default:
		return nil, fmt.Errorf("unable to decode PhotoSizeList: invalid constructor %x", id)
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
	return iface.DebugStringWithName("photoSizeList", m)
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
		return fmt.Errorf("unable to encode photoSizeList: unsupported layer %d", layer)
	}
}

// Encode <--
func (m *TLPhotoSizeList) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_photoSizeList, int(layer)); clazzId {
	case 0x67139b3:
		x.PutClazzID(0x67139b3)

		x.PutInt64(m.SizeId)

		if err := iface.EncodeObjectList(x, m.Sizes, layer); err != nil {
			return fmt.Errorf("unable to encode photoSizeList#0x67139b3: field sizes: %w", err)
		}

		x.PutInt32(m.DcId)

		return nil
	default:
		return fmt.Errorf("unable to encode photoSizeList: unsupported layer %d", layer)
	}
}

// Decode <--
func (m *TLPhotoSizeList) Decode(d *bin.Decoder) (err error) {
	switch m.ClazzID {
	case 0x67139b3:
		m.SizeId, err = d.Int64()
		if err != nil {
			return fmt.Errorf("unable to decode photoSizeList#0x67139b3: field size_id: %w", err)
		}
		l1, err3 := d.VectorHeader()
		if err3 != nil {
			return fmt.Errorf("unable to decode photoSizeList#0x67139b3: field sizes: %w", err3)
		}
		if l1 > bin.MaxVectorLen {
			return fmt.Errorf("unable to decode photoSizeList#0x67139b3: field sizes: %w", &bin.InvalidLengthError{Type: "vector", Length: int(l1)})
		}
		prealloc1 := int(l1)
		if prealloc1 > bin.PreallocateLimit {
			prealloc1 = bin.PreallocateLimit
		}
		v1 := make([]tg.PhotoSizeClazz, 0, prealloc1)
		for i := int32(0); i < l1; i++ {
			vv1, err3 := tg.DecodePhotoSizeClazz(d)
			if err3 != nil {
				return fmt.Errorf("unable to decode photoSizeList#0x67139b3: field sizes: %w", err3)
			}
			v1 = append(v1, vv1)
		}
		m.Sizes = v1

		m.DcId, err = d.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode photoSizeList#0x67139b3: field dc_id: %w", err)
		}

		return nil
	default:
		return fmt.Errorf("unable to decode photoSizeList: invalid constructor %x", m.ClazzID)
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
		return nil, fmt.Errorf("unable to decode VideoSizeList: constructor: %w", err)
	}

	switch id {
	case 0x38d19bf2:
		x := &TLVideoSizeList{ClazzID: id, ClazzName2: ClazzName_videoSizeList}
		if err := x.Decode(d); err != nil {
			return nil, err
		}
		return x, nil
	default:
		return nil, fmt.Errorf("unable to decode VideoSizeList: invalid constructor %x", id)
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
	return iface.DebugStringWithName("videoSizeList", m)
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
		return fmt.Errorf("unable to encode videoSizeList: unsupported layer %d", layer)
	}
}

// Encode <--
func (m *TLVideoSizeList) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_videoSizeList, int(layer)); clazzId {
	case 0x38d19bf2:
		x.PutClazzID(0x38d19bf2)

		x.PutInt64(m.SizeId)

		if err := iface.EncodeObjectList(x, m.Sizes, layer); err != nil {
			return fmt.Errorf("unable to encode videoSizeList#0x38d19bf2: field sizes: %w", err)
		}

		x.PutInt32(m.DcId)

		return nil
	default:
		return fmt.Errorf("unable to encode videoSizeList: unsupported layer %d", layer)
	}
}

// Decode <--
func (m *TLVideoSizeList) Decode(d *bin.Decoder) (err error) {
	switch m.ClazzID {
	case 0x38d19bf2:
		m.SizeId, err = d.Int64()
		if err != nil {
			return fmt.Errorf("unable to decode videoSizeList#0x38d19bf2: field size_id: %w", err)
		}
		l1, err3 := d.VectorHeader()
		if err3 != nil {
			return fmt.Errorf("unable to decode videoSizeList#0x38d19bf2: field sizes: %w", err3)
		}
		if l1 > bin.MaxVectorLen {
			return fmt.Errorf("unable to decode videoSizeList#0x38d19bf2: field sizes: %w", &bin.InvalidLengthError{Type: "vector", Length: int(l1)})
		}
		prealloc1 := int(l1)
		if prealloc1 > bin.PreallocateLimit {
			prealloc1 = bin.PreallocateLimit
		}
		v1 := make([]tg.VideoSizeClazz, 0, prealloc1)
		for i := int32(0); i < l1; i++ {
			vv1, err3 := tg.DecodeVideoSizeClazz(d)
			if err3 != nil {
				return fmt.Errorf("unable to decode videoSizeList#0x38d19bf2: field sizes: %w", err3)
			}
			v1 = append(v1, vv1)
		}
		m.Sizes = v1

		m.DcId, err = d.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode videoSizeList#0x38d19bf2: field dc_id: %w", err)
		}

		return nil
	default:
		return fmt.Errorf("unable to decode videoSizeList: invalid constructor %x", m.ClazzID)
	}
}

// VideoSizeList <--
type VideoSizeList = TLVideoSizeList
