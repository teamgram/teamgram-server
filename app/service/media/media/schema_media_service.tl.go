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
	"context"
	"encoding/json"
	"fmt"

	"github.com/teamgram/proto/v2/bin"
	"github.com/teamgram/proto/v2/iface"
	"github.com/teamgram/proto/v2/tg"
)

var (
	_ iface.TLObject
	_ fmt.Stringer
	_ *tg.Bool
	_ bin.Fields
	_ json.Marshaler
)

// TLMediaUploadPhotoFile <--
type TLMediaUploadPhotoFile struct {
	ClazzID    uint32                  `json:"_id"`
	OwnerId    int64                   `json:"owner_id"`
	File       tg.InputFileClazz       `json:"file"`
	Stickers   []tg.InputDocumentClazz `json:"stickers"`
	TtlSeconds *int32                  `json:"ttl_seconds"`
}

func (m *TLMediaUploadPhotoFile) String() string {
	wrapper := iface.WithNameWrapper{"", m}
	return wrapper.String()
}

// Encode <--
func (m *TLMediaUploadPhotoFile) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0x3c2b0b17: func() error {
			x.PutClazzID(0x3c2b0b17)

			// set flags
			var getFlags = func() uint32 {
				var flags uint32 = 0

				if m.Stickers != nil {
					flags |= 1 << 0
				}
				if m.TtlSeconds != nil {
					flags |= 1 << 1
				}

				return flags
			}

			// set flags
			var flags = getFlags()
			x.PutUint32(flags)
			x.PutInt64(m.OwnerId)
			_ = m.File.Encode(x, layer)
			if m.Stickers != nil {
				_ = iface.EncodeObjectList(x, m.Stickers, layer)
			}
			if m.TtlSeconds != nil {
				x.PutInt32(*m.TtlSeconds)
			}

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_media_uploadPhotoFile, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_media_uploadPhotoFile, layer)
	}
}

// Decode <--
func (m *TLMediaUploadPhotoFile) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0x3c2b0b17: func() (err error) {
			flags, _ := d.Uint32()
			_ = flags
			m.OwnerId, err = d.Int64()

			// m3 := &tg.InputFile{}
			// _ = m3.Decode(d)
			// m.File = m3
			m.File, _ = tg.DecodeInputFileClazz(d)

			if (flags & (1 << 0)) != 0 {
				c4, err2 := d.ClazzID()
				if c4 != iface.ClazzID_vector {
					// dBuf.err = fmt.Errorf("invalid ClazzID_vector, c%d: %d", 4, c4)
					return err2
				}
				l4, err3 := d.Int()
				v4 := make([]tg.InputDocumentClazz, l4)
				for i := 0; i < l4; i++ {
					// vv := new(InputDocument)
					// err3 = vv.Decode(d)
					// _ = err3
					// v4[i] = vv
					v4[i], err3 = tg.DecodeInputDocumentClazz(d)
					_ = err3
				}
				m.Stickers = v4
			}
			if (flags & (1 << 1)) != 0 {
				m.TtlSeconds = new(int32)
				*m.TtlSeconds, err = d.Int32()
			}

			return nil
		},
	}

	if m.ClazzID == 0 {
		m.ClazzID, _ = d.ClazzID()
	}
	if f, ok := decodeF[m.ClazzID]; ok {
		return f()
	} else {
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// TLMediaUploadProfilePhotoFile <--
type TLMediaUploadProfilePhotoFile struct {
	ClazzID      uint32            `json:"_id"`
	OwnerId      int64             `json:"owner_id"`
	File         tg.InputFileClazz `json:"file"`
	Video        tg.InputFileClazz `json:"video"`
	VideoStartTs *float64          `json:"video_start_ts"`
}

func (m *TLMediaUploadProfilePhotoFile) String() string {
	wrapper := iface.WithNameWrapper{"", m}
	return wrapper.String()
}

// Encode <--
func (m *TLMediaUploadProfilePhotoFile) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0x973f2f24: func() error {
			x.PutClazzID(0x973f2f24)

			// set flags
			var getFlags = func() uint32 {
				var flags uint32 = 0

				if m.File != nil {
					flags |= 1 << 0
				}
				if m.Video != nil {
					flags |= 1 << 1
				}
				if m.VideoStartTs != nil {
					flags |= 1 << 2
				}

				return flags
			}

			// set flags
			var flags = getFlags()
			x.PutUint32(flags)
			x.PutInt64(m.OwnerId)
			if m.File != nil {
				_ = m.File.Encode(x, layer)
			}

			if m.Video != nil {
				_ = m.Video.Encode(x, layer)
			}

			if m.VideoStartTs != nil {
				x.PutDouble(*m.VideoStartTs)
			}

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_media_uploadProfilePhotoFile, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_media_uploadProfilePhotoFile, layer)
	}
}

// Decode <--
func (m *TLMediaUploadProfilePhotoFile) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0x973f2f24: func() (err error) {
			flags, _ := d.Uint32()
			_ = flags
			m.OwnerId, err = d.Int64()
			if (flags & (1 << 0)) != 0 {
				// m3 := &tg.InputFile{}
				// _ = m3.Decode(d)
				// m.File = m3
				m.File, _ = tg.DecodeInputFileClazz(d)
			}
			if (flags & (1 << 1)) != 0 {
				// m4 := &tg.InputFile{}
				// _ = m4.Decode(d)
				// m.Video = m4
				m.Video, _ = tg.DecodeInputFileClazz(d)
			}
			if (flags & (1 << 2)) != 0 {
				m.VideoStartTs = new(float64)
				*m.VideoStartTs, err = d.Double()
			}

			return nil
		},
	}

	if m.ClazzID == 0 {
		m.ClazzID, _ = d.ClazzID()
	}
	if f, ok := decodeF[m.ClazzID]; ok {
		return f()
	} else {
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// TLMediaGetPhoto <--
type TLMediaGetPhoto struct {
	ClazzID uint32 `json:"_id"`
	PhotoId int64  `json:"photo_id"`
}

func (m *TLMediaGetPhoto) String() string {
	wrapper := iface.WithNameWrapper{"", m}
	return wrapper.String()
}

// Encode <--
func (m *TLMediaGetPhoto) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0x657eb86b: func() error {
			x.PutClazzID(0x657eb86b)

			x.PutInt64(m.PhotoId)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_media_getPhoto, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_media_getPhoto, layer)
	}
}

// Decode <--
func (m *TLMediaGetPhoto) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0x657eb86b: func() (err error) {
			m.PhotoId, err = d.Int64()

			return nil
		},
	}

	if m.ClazzID == 0 {
		m.ClazzID, _ = d.ClazzID()
	}
	if f, ok := decodeF[m.ClazzID]; ok {
		return f()
	} else {
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// TLMediaGetPhotoSizeList <--
type TLMediaGetPhotoSizeList struct {
	ClazzID uint32 `json:"_id"`
	SizeId  int64  `json:"size_id"`
}

func (m *TLMediaGetPhotoSizeList) String() string {
	wrapper := iface.WithNameWrapper{"", m}
	return wrapper.String()
}

// Encode <--
func (m *TLMediaGetPhotoSizeList) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0xa1eb7f45: func() error {
			x.PutClazzID(0xa1eb7f45)

			x.PutInt64(m.SizeId)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_media_getPhotoSizeList, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_media_getPhotoSizeList, layer)
	}
}

// Decode <--
func (m *TLMediaGetPhotoSizeList) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0xa1eb7f45: func() (err error) {
			m.SizeId, err = d.Int64()

			return nil
		},
	}

	if m.ClazzID == 0 {
		m.ClazzID, _ = d.ClazzID()
	}
	if f, ok := decodeF[m.ClazzID]; ok {
		return f()
	} else {
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// TLMediaGetPhotoSizeListList <--
type TLMediaGetPhotoSizeListList struct {
	ClazzID uint32  `json:"_id"`
	IdList  []int64 `json:"id_list"`
}

func (m *TLMediaGetPhotoSizeListList) String() string {
	wrapper := iface.WithNameWrapper{"", m}
	return wrapper.String()
}

// Encode <--
func (m *TLMediaGetPhotoSizeListList) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0xfb5c80e0: func() error {
			x.PutClazzID(0xfb5c80e0)

			iface.EncodeInt64List(x, m.IdList)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_media_getPhotoSizeListList, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_media_getPhotoSizeListList, layer)
	}
}

// Decode <--
func (m *TLMediaGetPhotoSizeListList) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0xfb5c80e0: func() (err error) {

			m.IdList, err = iface.DecodeInt64List(d)

			return nil
		},
	}

	if m.ClazzID == 0 {
		m.ClazzID, _ = d.ClazzID()
	}
	if f, ok := decodeF[m.ClazzID]; ok {
		return f()
	} else {
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// TLMediaGetVideoSizeList <--
type TLMediaGetVideoSizeList struct {
	ClazzID uint32 `json:"_id"`
	SizeId  int64  `json:"size_id"`
}

func (m *TLMediaGetVideoSizeList) String() string {
	wrapper := iface.WithNameWrapper{"", m}
	return wrapper.String()
}

// Encode <--
func (m *TLMediaGetVideoSizeList) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0xc47692ea: func() error {
			x.PutClazzID(0xc47692ea)

			x.PutInt64(m.SizeId)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_media_getVideoSizeList, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_media_getVideoSizeList, layer)
	}
}

// Decode <--
func (m *TLMediaGetVideoSizeList) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0xc47692ea: func() (err error) {
			m.SizeId, err = d.Int64()

			return nil
		},
	}

	if m.ClazzID == 0 {
		m.ClazzID, _ = d.ClazzID()
	}
	if f, ok := decodeF[m.ClazzID]; ok {
		return f()
	} else {
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// TLMediaUploadedDocumentMedia <--
type TLMediaUploadedDocumentMedia struct {
	ClazzID uint32             `json:"_id"`
	OwnerId int64              `json:"owner_id"`
	Media   tg.InputMediaClazz `json:"media"`
}

func (m *TLMediaUploadedDocumentMedia) String() string {
	wrapper := iface.WithNameWrapper{"", m}
	return wrapper.String()
}

// Encode <--
func (m *TLMediaUploadedDocumentMedia) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0x4f5fb06c: func() error {
			x.PutClazzID(0x4f5fb06c)

			x.PutInt64(m.OwnerId)
			_ = m.Media.Encode(x, layer)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_media_uploadedDocumentMedia, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_media_uploadedDocumentMedia, layer)
	}
}

// Decode <--
func (m *TLMediaUploadedDocumentMedia) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0x4f5fb06c: func() (err error) {
			m.OwnerId, err = d.Int64()

			// m2 := &tg.InputMedia{}
			// _ = m2.Decode(d)
			// m.Media = m2
			m.Media, _ = tg.DecodeInputMediaClazz(d)

			return nil
		},
	}

	if m.ClazzID == 0 {
		m.ClazzID, _ = d.ClazzID()
	}
	if f, ok := decodeF[m.ClazzID]; ok {
		return f()
	} else {
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// TLMediaGetDocument <--
type TLMediaGetDocument struct {
	ClazzID uint32 `json:"_id"`
	Id      int64  `json:"id"`
}

func (m *TLMediaGetDocument) String() string {
	wrapper := iface.WithNameWrapper{"", m}
	return wrapper.String()
}

// Encode <--
func (m *TLMediaGetDocument) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0x3fe5974d: func() error {
			x.PutClazzID(0x3fe5974d)

			x.PutInt64(m.Id)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_media_getDocument, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_media_getDocument, layer)
	}
}

// Decode <--
func (m *TLMediaGetDocument) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0x3fe5974d: func() (err error) {
			m.Id, err = d.Int64()

			return nil
		},
	}

	if m.ClazzID == 0 {
		m.ClazzID, _ = d.ClazzID()
	}
	if f, ok := decodeF[m.ClazzID]; ok {
		return f()
	} else {
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// TLMediaGetDocumentList <--
type TLMediaGetDocumentList struct {
	ClazzID uint32  `json:"_id"`
	IdList  []int64 `json:"id_list"`
}

func (m *TLMediaGetDocumentList) String() string {
	wrapper := iface.WithNameWrapper{"", m}
	return wrapper.String()
}

// Encode <--
func (m *TLMediaGetDocumentList) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0xc52fd26f: func() error {
			x.PutClazzID(0xc52fd26f)

			iface.EncodeInt64List(x, m.IdList)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_media_getDocumentList, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_media_getDocumentList, layer)
	}
}

// Decode <--
func (m *TLMediaGetDocumentList) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0xc52fd26f: func() (err error) {

			m.IdList, err = iface.DecodeInt64List(d)

			return nil
		},
	}

	if m.ClazzID == 0 {
		m.ClazzID, _ = d.ClazzID()
	}
	if f, ok := decodeF[m.ClazzID]; ok {
		return f()
	} else {
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// TLMediaUploadEncryptedFile <--
type TLMediaUploadEncryptedFile struct {
	ClazzID uint32                     `json:"_id"`
	OwnerId int64                      `json:"owner_id"`
	File    tg.InputEncryptedFileClazz `json:"file"`
}

func (m *TLMediaUploadEncryptedFile) String() string {
	wrapper := iface.WithNameWrapper{"", m}
	return wrapper.String()
}

// Encode <--
func (m *TLMediaUploadEncryptedFile) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0xab00c69b: func() error {
			x.PutClazzID(0xab00c69b)

			x.PutInt64(m.OwnerId)
			_ = m.File.Encode(x, layer)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_media_uploadEncryptedFile, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_media_uploadEncryptedFile, layer)
	}
}

// Decode <--
func (m *TLMediaUploadEncryptedFile) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0xab00c69b: func() (err error) {
			m.OwnerId, err = d.Int64()

			// m2 := &tg.InputEncryptedFile{}
			// _ = m2.Decode(d)
			// m.File = m2
			m.File, _ = tg.DecodeInputEncryptedFileClazz(d)

			return nil
		},
	}

	if m.ClazzID == 0 {
		m.ClazzID, _ = d.ClazzID()
	}
	if f, ok := decodeF[m.ClazzID]; ok {
		return f()
	} else {
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// TLMediaGetEncryptedFile <--
type TLMediaGetEncryptedFile struct {
	ClazzID    uint32 `json:"_id"`
	Id         int64  `json:"id"`
	AccessHash int64  `json:"access_hash"`
}

func (m *TLMediaGetEncryptedFile) String() string {
	wrapper := iface.WithNameWrapper{"", m}
	return wrapper.String()
}

// Encode <--
func (m *TLMediaGetEncryptedFile) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0xfc6080d1: func() error {
			x.PutClazzID(0xfc6080d1)

			x.PutInt64(m.Id)
			x.PutInt64(m.AccessHash)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_media_getEncryptedFile, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_media_getEncryptedFile, layer)
	}
}

// Decode <--
func (m *TLMediaGetEncryptedFile) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0xfc6080d1: func() (err error) {
			m.Id, err = d.Int64()
			m.AccessHash, err = d.Int64()

			return nil
		},
	}

	if m.ClazzID == 0 {
		m.ClazzID, _ = d.ClazzID()
	}
	if f, ok := decodeF[m.ClazzID]; ok {
		return f()
	} else {
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// TLMediaUploadWallPaperFile <--
type TLMediaUploadWallPaperFile struct {
	ClazzID  uint32            `json:"_id"`
	OwnerId  int64             `json:"owner_id"`
	File     tg.InputFileClazz `json:"file"`
	MimeType string            `json:"mime_type"`
	Admin    tg.BoolClazz      `json:"admin"`
}

func (m *TLMediaUploadWallPaperFile) String() string {
	wrapper := iface.WithNameWrapper{"", m}
	return wrapper.String()
}

// Encode <--
func (m *TLMediaUploadWallPaperFile) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0x9cfaadfe: func() error {
			x.PutClazzID(0x9cfaadfe)

			x.PutInt64(m.OwnerId)
			_ = m.File.Encode(x, layer)
			x.PutString(m.MimeType)
			_ = m.Admin.Encode(x, layer)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_media_uploadWallPaperFile, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_media_uploadWallPaperFile, layer)
	}
}

// Decode <--
func (m *TLMediaUploadWallPaperFile) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0x9cfaadfe: func() (err error) {
			m.OwnerId, err = d.Int64()

			// m2 := &tg.InputFile{}
			// _ = m2.Decode(d)
			// m.File = m2
			m.File, _ = tg.DecodeInputFileClazz(d)

			m.MimeType, err = d.String()

			// m4 := &tg.Bool{}
			// _ = m4.Decode(d)
			// m.Admin = m4
			m.Admin, _ = tg.DecodeBoolClazz(d)

			return nil
		},
	}

	if m.ClazzID == 0 {
		m.ClazzID, _ = d.ClazzID()
	}
	if f, ok := decodeF[m.ClazzID]; ok {
		return f()
	} else {
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// TLMediaUploadThemeFile <--
type TLMediaUploadThemeFile struct {
	ClazzID  uint32            `json:"_id"`
	OwnerId  int64             `json:"owner_id"`
	File     tg.InputFileClazz `json:"file"`
	Thumb    tg.InputFileClazz `json:"thumb"`
	MimeType string            `json:"mime_type"`
	FileName string            `json:"file_name"`
}

func (m *TLMediaUploadThemeFile) String() string {
	wrapper := iface.WithNameWrapper{"", m}
	return wrapper.String()
}

// Encode <--
func (m *TLMediaUploadThemeFile) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0x42e6b860: func() error {
			x.PutClazzID(0x42e6b860)

			// set flags
			var getFlags = func() uint32 {
				var flags uint32 = 0

				if m.Thumb != nil {
					flags |= 1 << 0
				}

				return flags
			}

			// set flags
			var flags = getFlags()
			x.PutUint32(flags)
			x.PutInt64(m.OwnerId)
			_ = m.File.Encode(x, layer)
			if m.Thumb != nil {
				_ = m.Thumb.Encode(x, layer)
			}

			x.PutString(m.MimeType)
			x.PutString(m.FileName)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_media_uploadThemeFile, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_media_uploadThemeFile, layer)
	}
}

// Decode <--
func (m *TLMediaUploadThemeFile) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0x42e6b860: func() (err error) {
			flags, _ := d.Uint32()
			_ = flags
			m.OwnerId, err = d.Int64()

			// m3 := &tg.InputFile{}
			// _ = m3.Decode(d)
			// m.File = m3
			m.File, _ = tg.DecodeInputFileClazz(d)

			if (flags & (1 << 0)) != 0 {
				// m4 := &tg.InputFile{}
				// _ = m4.Decode(d)
				// m.Thumb = m4
				m.Thumb, _ = tg.DecodeInputFileClazz(d)
			}
			m.MimeType, err = d.String()
			m.FileName, err = d.String()

			return nil
		},
	}

	if m.ClazzID == 0 {
		m.ClazzID, _ = d.ClazzID()
	}
	if f, ok := decodeF[m.ClazzID]; ok {
		return f()
	} else {
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// TLMediaUploadStickerFile <--
type TLMediaUploadStickerFile struct {
	ClazzID                  uint32                    `json:"_id"`
	OwnerId                  int64                     `json:"owner_id"`
	File                     tg.InputFileClazz         `json:"file"`
	Thumb                    tg.InputFileClazz         `json:"thumb"`
	MimeType                 string                    `json:"mime_type"`
	FileName                 string                    `json:"file_name"`
	DocumentAttributeSticker tg.DocumentAttributeClazz `json:"document_attribute_sticker"`
}

func (m *TLMediaUploadStickerFile) String() string {
	wrapper := iface.WithNameWrapper{"", m}
	return wrapper.String()
}

// Encode <--
func (m *TLMediaUploadStickerFile) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0xacb624ed: func() error {
			x.PutClazzID(0xacb624ed)

			// set flags
			var getFlags = func() uint32 {
				var flags uint32 = 0

				if m.Thumb != nil {
					flags |= 1 << 0
				}

				return flags
			}

			// set flags
			var flags = getFlags()
			x.PutUint32(flags)
			x.PutInt64(m.OwnerId)
			_ = m.File.Encode(x, layer)
			if m.Thumb != nil {
				_ = m.Thumb.Encode(x, layer)
			}

			x.PutString(m.MimeType)
			x.PutString(m.FileName)
			_ = m.DocumentAttributeSticker.Encode(x, layer)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_media_uploadStickerFile, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_media_uploadStickerFile, layer)
	}
}

// Decode <--
func (m *TLMediaUploadStickerFile) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0xacb624ed: func() (err error) {
			flags, _ := d.Uint32()
			_ = flags
			m.OwnerId, err = d.Int64()

			// m3 := &tg.InputFile{}
			// _ = m3.Decode(d)
			// m.File = m3
			m.File, _ = tg.DecodeInputFileClazz(d)

			if (flags & (1 << 0)) != 0 {
				// m4 := &tg.InputFile{}
				// _ = m4.Decode(d)
				// m.Thumb = m4
				m.Thumb, _ = tg.DecodeInputFileClazz(d)
			}
			m.MimeType, err = d.String()
			m.FileName, err = d.String()

			// m7 := &tg.DocumentAttribute{}
			// _ = m7.Decode(d)
			// m.DocumentAttributeSticker = m7
			m.DocumentAttributeSticker, _ = tg.DecodeDocumentAttributeClazz(d)

			return nil
		},
	}

	if m.ClazzID == 0 {
		m.ClazzID, _ = d.ClazzID()
	}
	if f, ok := decodeF[m.ClazzID]; ok {
		return f()
	} else {
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// TLMediaUploadRingtoneFile <--
type TLMediaUploadRingtoneFile struct {
	ClazzID  uint32            `json:"_id"`
	OwnerId  int64             `json:"owner_id"`
	File     tg.InputFileClazz `json:"file"`
	MimeType string            `json:"mime_type"`
	FileName string            `json:"file_name"`
}

func (m *TLMediaUploadRingtoneFile) String() string {
	wrapper := iface.WithNameWrapper{"", m}
	return wrapper.String()
}

// Encode <--
func (m *TLMediaUploadRingtoneFile) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0x3dbab209: func() error {
			x.PutClazzID(0x3dbab209)

			// set flags
			var getFlags = func() uint32 {
				var flags uint32 = 0

				return flags
			}

			// set flags
			var flags = getFlags()
			x.PutUint32(flags)
			x.PutInt64(m.OwnerId)
			_ = m.File.Encode(x, layer)
			x.PutString(m.MimeType)
			x.PutString(m.FileName)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_media_uploadRingtoneFile, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_media_uploadRingtoneFile, layer)
	}
}

// Decode <--
func (m *TLMediaUploadRingtoneFile) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0x3dbab209: func() (err error) {
			flags, _ := d.Uint32()
			_ = flags
			m.OwnerId, err = d.Int64()

			// m3 := &tg.InputFile{}
			// _ = m3.Decode(d)
			// m.File = m3
			m.File, _ = tg.DecodeInputFileClazz(d)

			m.MimeType, err = d.String()
			m.FileName, err = d.String()

			return nil
		},
	}

	if m.ClazzID == 0 {
		m.ClazzID, _ = d.ClazzID()
	}
	if f, ok := decodeF[m.ClazzID]; ok {
		return f()
	} else {
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// TLMediaUploadedProfilePhoto <--
type TLMediaUploadedProfilePhoto struct {
	ClazzID uint32 `json:"_id"`
	OwnerId int64  `json:"owner_id"`
	PhotoId int64  `json:"photo_id"`
}

func (m *TLMediaUploadedProfilePhoto) String() string {
	wrapper := iface.WithNameWrapper{"", m}
	return wrapper.String()
}

// Encode <--
func (m *TLMediaUploadedProfilePhoto) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0x89d159d2: func() error {
			x.PutClazzID(0x89d159d2)

			x.PutInt64(m.OwnerId)
			x.PutInt64(m.PhotoId)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_media_uploadedProfilePhoto, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_media_uploadedProfilePhoto, layer)
	}
}

// Decode <--
func (m *TLMediaUploadedProfilePhoto) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0x89d159d2: func() (err error) {
			m.OwnerId, err = d.Int64()
			m.PhotoId, err = d.Int64()

			return nil
		},
	}

	if m.ClazzID == 0 {
		m.ClazzID, _ = d.ClazzID()
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

// VectorPhotoSizeList <--
type VectorPhotoSizeList struct {
	Datas []PhotoSizeListClazz `json:"_datas"`
}

func (m *VectorPhotoSizeList) String() string {
	data, _ := json.Marshal(m)
	return string(data)
}

// Encode <--
func (m *VectorPhotoSizeList) Encode(x *bin.Encoder, layer int32) error {
	_ = iface.EncodeObjectList(x, m.Datas, layer)

	return nil
}

// Decode <--
func (m *VectorPhotoSizeList) Decode(d *bin.Decoder) (err error) {
	m.Datas, err = iface.DecodeObjectList[PhotoSizeListClazz](d)

	return err
}

// VectorDocument <--
type VectorDocument struct {
	Datas []tg.DocumentClazz `json:"_datas"`
}

func (m *VectorDocument) String() string {
	data, _ := json.Marshal(m)
	return string(data)
}

// Encode <--
func (m *VectorDocument) Encode(x *bin.Encoder, layer int32) error {
	_ = iface.EncodeObjectList(x, m.Datas, layer)

	return nil
}

// Decode <--
func (m *VectorDocument) Decode(d *bin.Decoder) (err error) {
	m.Datas, err = iface.DecodeObjectList[tg.DocumentClazz](d)

	return err
}

// ----------------------------------------------------------------------------
// rpc

type RPCMedia interface {
	MediaUploadPhotoFile(ctx context.Context, in *TLMediaUploadPhotoFile) (*tg.Photo, error)
	MediaUploadProfilePhotoFile(ctx context.Context, in *TLMediaUploadProfilePhotoFile) (*tg.Photo, error)
	MediaGetPhoto(ctx context.Context, in *TLMediaGetPhoto) (*tg.Photo, error)
	MediaGetPhotoSizeList(ctx context.Context, in *TLMediaGetPhotoSizeList) (*PhotoSizeList, error)
	MediaGetPhotoSizeListList(ctx context.Context, in *TLMediaGetPhotoSizeListList) (*VectorPhotoSizeList, error)
	MediaGetVideoSizeList(ctx context.Context, in *TLMediaGetVideoSizeList) (*VideoSizeList, error)
	MediaUploadedDocumentMedia(ctx context.Context, in *TLMediaUploadedDocumentMedia) (*tg.MessageMedia, error)
	MediaGetDocument(ctx context.Context, in *TLMediaGetDocument) (*tg.Document, error)
	MediaGetDocumentList(ctx context.Context, in *TLMediaGetDocumentList) (*VectorDocument, error)
	MediaUploadEncryptedFile(ctx context.Context, in *TLMediaUploadEncryptedFile) (*tg.EncryptedFile, error)
	MediaGetEncryptedFile(ctx context.Context, in *TLMediaGetEncryptedFile) (*tg.EncryptedFile, error)
	MediaUploadWallPaperFile(ctx context.Context, in *TLMediaUploadWallPaperFile) (*tg.Document, error)
	MediaUploadThemeFile(ctx context.Context, in *TLMediaUploadThemeFile) (*tg.Document, error)
	MediaUploadStickerFile(ctx context.Context, in *TLMediaUploadStickerFile) (*tg.Document, error)
	MediaUploadRingtoneFile(ctx context.Context, in *TLMediaUploadRingtoneFile) (*tg.Document, error)
	MediaUploadedProfilePhoto(ctx context.Context, in *TLMediaUploadedProfilePhoto) (*tg.Photo, error)
}
