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

// TLMediaUploadPhotoFile <--
type TLMediaUploadPhotoFile struct {
	ClazzID    uint32                  `json:"_id"`
	OwnerId    int64                   `json:"owner_id"`
	File       tg.InputFileClazz       `json:"file"`
	Stickers   []tg.InputDocumentClazz `json:"stickers"`
	TtlSeconds *int32                  `json:"ttl_seconds"`
}

func (m *TLMediaUploadPhotoFile) String() string {
	wrapper := iface.WithNameWrapper{ClazzName: ClazzName_media_uploadPhotoFile, TLObject: m}
	return wrapper.String()
}

// Encode <--
func (m *TLMediaUploadPhotoFile) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_media_uploadPhotoFile, int(layer)); clazzId {
	case 0x3c2b0b17:
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
			if err := iface.EncodeObjectList(x, m.Stickers, layer); err != nil {
				return err
			}
		}
		if m.TtlSeconds != nil {
			x.PutInt32(*m.TtlSeconds)
		}

		return nil
	default:
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_media_uploadPhotoFile, layer)
	}
}

// Decode <--
func (m *TLMediaUploadPhotoFile) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return err
		}
	}
	switch m.ClazzID {
	case 0x3c2b0b17:
		flags, err := d.Uint32()
		if err != nil {
			return err
		}
		_ = flags
		m.OwnerId, err = d.Int64()
		if err != nil {
			return err
		}

		m.File, err = tg.DecodeInputFileClazz(d)
		if err != nil {
			return err
		}

		if (flags & (1 << 0)) != 0 {
			c4, err2 := d.ClazzID()
			if err2 != nil {
				return err2
			}
			if c4 != iface.ClazzID_vector {
				return fmt.Errorf("invalid ClazzID_vector, c%d: %d", 4, c4)
			}
			l4, err3 := d.Int()
			if err3 != nil {
				return err3
			}
			v4 := make([]tg.InputDocumentClazz, l4)
			for i := 0; i < l4; i++ {
				v4[i], err3 = tg.DecodeInputDocumentClazz(d)
				if err3 != nil {
					return err3
				}
			}
			m.Stickers = v4
		}
		if (flags & (1 << 1)) != 0 {
			m.TtlSeconds = new(int32)
			*m.TtlSeconds, err = d.Int32()
			if err != nil {
				return err
			}
		}

		return nil
	default:
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// TLMediaUploadProfilePhotoFile <--
type TLMediaUploadProfilePhotoFile struct {
	ClazzID          uint32            `json:"_id"`
	OwnerId          int64             `json:"owner_id"`
	File             tg.InputFileClazz `json:"file"`
	Video            tg.InputFileClazz `json:"video"`
	VideoStartTs     *float64          `json:"video_start_ts"`
	VideoEmojiMarkup tg.VideoSizeClazz `json:"video_emoji_markup"`
}

func (m *TLMediaUploadProfilePhotoFile) String() string {
	wrapper := iface.WithNameWrapper{ClazzName: ClazzName_media_uploadProfilePhotoFile, TLObject: m}
	return wrapper.String()
}

// Encode <--
func (m *TLMediaUploadProfilePhotoFile) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_media_uploadProfilePhotoFile, int(layer)); clazzId {
	case 0xb6a04cc4:
		x.PutClazzID(0xb6a04cc4)

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
			if m.VideoEmojiMarkup != nil {
				flags |= 1 << 4
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

		if m.VideoEmojiMarkup != nil {
			_ = m.VideoEmojiMarkup.Encode(x, layer)
		}

		return nil
	default:
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_media_uploadProfilePhotoFile, layer)
	}
}

// Decode <--
func (m *TLMediaUploadProfilePhotoFile) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return err
		}
	}
	switch m.ClazzID {
	case 0xb6a04cc4:
		flags, err := d.Uint32()
		if err != nil {
			return err
		}
		_ = flags
		m.OwnerId, err = d.Int64()
		if err != nil {
			return err
		}
		if (flags & (1 << 0)) != 0 {
			m.File, err = tg.DecodeInputFileClazz(d)
			if err != nil {
				return err
			}
		}
		if (flags & (1 << 1)) != 0 {
			m.Video, err = tg.DecodeInputFileClazz(d)
			if err != nil {
				return err
			}
		}
		if (flags & (1 << 2)) != 0 {
			m.VideoStartTs = new(float64)
			*m.VideoStartTs, err = d.Double()
			if err != nil {
				return err
			}
		}

		if (flags & (1 << 4)) != 0 {
			m.VideoEmojiMarkup, err = tg.DecodeVideoSizeClazz(d)
			if err != nil {
				return err
			}
		}

		return nil
	default:
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// TLMediaGetPhoto <--
type TLMediaGetPhoto struct {
	ClazzID uint32 `json:"_id"`
	PhotoId int64  `json:"photo_id"`
}

func (m *TLMediaGetPhoto) String() string {
	wrapper := iface.WithNameWrapper{ClazzName: ClazzName_media_getPhoto, TLObject: m}
	return wrapper.String()
}

// Encode <--
func (m *TLMediaGetPhoto) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_media_getPhoto, int(layer)); clazzId {
	case 0x657eb86b:
		x.PutClazzID(0x657eb86b)

		x.PutInt64(m.PhotoId)

		return nil
	default:
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_media_getPhoto, layer)
	}
}

// Decode <--
func (m *TLMediaGetPhoto) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return err
		}
	}
	switch m.ClazzID {
	case 0x657eb86b:
		m.PhotoId, err = d.Int64()
		if err != nil {
			return err
		}

		return nil
	default:
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// TLMediaGetPhotoSizeList <--
type TLMediaGetPhotoSizeList struct {
	ClazzID uint32 `json:"_id"`
	SizeId  int64  `json:"size_id"`
}

func (m *TLMediaGetPhotoSizeList) String() string {
	wrapper := iface.WithNameWrapper{ClazzName: ClazzName_media_getPhotoSizeList, TLObject: m}
	return wrapper.String()
}

// Encode <--
func (m *TLMediaGetPhotoSizeList) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_media_getPhotoSizeList, int(layer)); clazzId {
	case 0xa1eb7f45:
		x.PutClazzID(0xa1eb7f45)

		x.PutInt64(m.SizeId)

		return nil
	default:
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_media_getPhotoSizeList, layer)
	}
}

// Decode <--
func (m *TLMediaGetPhotoSizeList) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return err
		}
	}
	switch m.ClazzID {
	case 0xa1eb7f45:
		m.SizeId, err = d.Int64()
		if err != nil {
			return err
		}

		return nil
	default:
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// TLMediaGetPhotoSizeListList <--
type TLMediaGetPhotoSizeListList struct {
	ClazzID uint32  `json:"_id"`
	IdList  []int64 `json:"id_list"`
}

func (m *TLMediaGetPhotoSizeListList) String() string {
	wrapper := iface.WithNameWrapper{ClazzName: ClazzName_media_getPhotoSizeListList, TLObject: m}
	return wrapper.String()
}

// Encode <--
func (m *TLMediaGetPhotoSizeListList) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_media_getPhotoSizeListList, int(layer)); clazzId {
	case 0xfb5c80e0:
		x.PutClazzID(0xfb5c80e0)

		iface.EncodeInt64List(x, m.IdList)

		return nil
	default:
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_media_getPhotoSizeListList, layer)
	}
}

// Decode <--
func (m *TLMediaGetPhotoSizeListList) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return err
		}
	}
	switch m.ClazzID {
	case 0xfb5c80e0:

		m.IdList, err = iface.DecodeInt64List(d)
		if err != nil {
			return err
		}

		return nil
	default:
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// TLMediaGetVideoSizeList <--
type TLMediaGetVideoSizeList struct {
	ClazzID uint32 `json:"_id"`
	SizeId  int64  `json:"size_id"`
}

func (m *TLMediaGetVideoSizeList) String() string {
	wrapper := iface.WithNameWrapper{ClazzName: ClazzName_media_getVideoSizeList, TLObject: m}
	return wrapper.String()
}

// Encode <--
func (m *TLMediaGetVideoSizeList) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_media_getVideoSizeList, int(layer)); clazzId {
	case 0xc47692ea:
		x.PutClazzID(0xc47692ea)

		x.PutInt64(m.SizeId)

		return nil
	default:
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_media_getVideoSizeList, layer)
	}
}

// Decode <--
func (m *TLMediaGetVideoSizeList) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return err
		}
	}
	switch m.ClazzID {
	case 0xc47692ea:
		m.SizeId, err = d.Int64()
		if err != nil {
			return err
		}

		return nil
	default:
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
	wrapper := iface.WithNameWrapper{ClazzName: ClazzName_media_uploadedDocumentMedia, TLObject: m}
	return wrapper.String()
}

// Encode <--
func (m *TLMediaUploadedDocumentMedia) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_media_uploadedDocumentMedia, int(layer)); clazzId {
	case 0x4f5fb06c:
		x.PutClazzID(0x4f5fb06c)

		x.PutInt64(m.OwnerId)
		_ = m.Media.Encode(x, layer)

		return nil
	default:
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_media_uploadedDocumentMedia, layer)
	}
}

// Decode <--
func (m *TLMediaUploadedDocumentMedia) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return err
		}
	}
	switch m.ClazzID {
	case 0x4f5fb06c:
		m.OwnerId, err = d.Int64()
		if err != nil {
			return err
		}

		m.Media, err = tg.DecodeInputMediaClazz(d)
		if err != nil {
			return err
		}

		return nil
	default:
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// TLMediaGetDocument <--
type TLMediaGetDocument struct {
	ClazzID uint32 `json:"_id"`
	Id      int64  `json:"id"`
}

func (m *TLMediaGetDocument) String() string {
	wrapper := iface.WithNameWrapper{ClazzName: ClazzName_media_getDocument, TLObject: m}
	return wrapper.String()
}

// Encode <--
func (m *TLMediaGetDocument) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_media_getDocument, int(layer)); clazzId {
	case 0x3fe5974d:
		x.PutClazzID(0x3fe5974d)

		x.PutInt64(m.Id)

		return nil
	default:
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_media_getDocument, layer)
	}
}

// Decode <--
func (m *TLMediaGetDocument) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return err
		}
	}
	switch m.ClazzID {
	case 0x3fe5974d:
		m.Id, err = d.Int64()
		if err != nil {
			return err
		}

		return nil
	default:
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// TLMediaGetDocumentList <--
type TLMediaGetDocumentList struct {
	ClazzID uint32  `json:"_id"`
	IdList  []int64 `json:"id_list"`
}

func (m *TLMediaGetDocumentList) String() string {
	wrapper := iface.WithNameWrapper{ClazzName: ClazzName_media_getDocumentList, TLObject: m}
	return wrapper.String()
}

// Encode <--
func (m *TLMediaGetDocumentList) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_media_getDocumentList, int(layer)); clazzId {
	case 0xc52fd26f:
		x.PutClazzID(0xc52fd26f)

		iface.EncodeInt64List(x, m.IdList)

		return nil
	default:
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_media_getDocumentList, layer)
	}
}

// Decode <--
func (m *TLMediaGetDocumentList) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return err
		}
	}
	switch m.ClazzID {
	case 0xc52fd26f:

		m.IdList, err = iface.DecodeInt64List(d)
		if err != nil {
			return err
		}

		return nil
	default:
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
	wrapper := iface.WithNameWrapper{ClazzName: ClazzName_media_uploadEncryptedFile, TLObject: m}
	return wrapper.String()
}

// Encode <--
func (m *TLMediaUploadEncryptedFile) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_media_uploadEncryptedFile, int(layer)); clazzId {
	case 0xab00c69b:
		x.PutClazzID(0xab00c69b)

		x.PutInt64(m.OwnerId)
		_ = m.File.Encode(x, layer)

		return nil
	default:
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_media_uploadEncryptedFile, layer)
	}
}

// Decode <--
func (m *TLMediaUploadEncryptedFile) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return err
		}
	}
	switch m.ClazzID {
	case 0xab00c69b:
		m.OwnerId, err = d.Int64()
		if err != nil {
			return err
		}

		m.File, err = tg.DecodeInputEncryptedFileClazz(d)
		if err != nil {
			return err
		}

		return nil
	default:
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
	wrapper := iface.WithNameWrapper{ClazzName: ClazzName_media_getEncryptedFile, TLObject: m}
	return wrapper.String()
}

// Encode <--
func (m *TLMediaGetEncryptedFile) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_media_getEncryptedFile, int(layer)); clazzId {
	case 0xfc6080d1:
		x.PutClazzID(0xfc6080d1)

		x.PutInt64(m.Id)
		x.PutInt64(m.AccessHash)

		return nil
	default:
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_media_getEncryptedFile, layer)
	}
}

// Decode <--
func (m *TLMediaGetEncryptedFile) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return err
		}
	}
	switch m.ClazzID {
	case 0xfc6080d1:
		m.Id, err = d.Int64()
		if err != nil {
			return err
		}
		m.AccessHash, err = d.Int64()
		if err != nil {
			return err
		}

		return nil
	default:
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
	wrapper := iface.WithNameWrapper{ClazzName: ClazzName_media_uploadWallPaperFile, TLObject: m}
	return wrapper.String()
}

// Encode <--
func (m *TLMediaUploadWallPaperFile) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_media_uploadWallPaperFile, int(layer)); clazzId {
	case 0x9cfaadfe:
		x.PutClazzID(0x9cfaadfe)

		x.PutInt64(m.OwnerId)
		_ = m.File.Encode(x, layer)
		x.PutString(m.MimeType)
		_ = m.Admin.Encode(x, layer)

		return nil
	default:
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_media_uploadWallPaperFile, layer)
	}
}

// Decode <--
func (m *TLMediaUploadWallPaperFile) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return err
		}
	}
	switch m.ClazzID {
	case 0x9cfaadfe:
		m.OwnerId, err = d.Int64()
		if err != nil {
			return err
		}

		m.File, err = tg.DecodeInputFileClazz(d)
		if err != nil {
			return err
		}

		m.MimeType, err = d.String()
		if err != nil {
			return err
		}

		m.Admin, err = tg.DecodeBoolClazz(d)
		if err != nil {
			return err
		}

		return nil
	default:
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
	wrapper := iface.WithNameWrapper{ClazzName: ClazzName_media_uploadThemeFile, TLObject: m}
	return wrapper.String()
}

// Encode <--
func (m *TLMediaUploadThemeFile) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_media_uploadThemeFile, int(layer)); clazzId {
	case 0x42e6b860:
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
	default:
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_media_uploadThemeFile, layer)
	}
}

// Decode <--
func (m *TLMediaUploadThemeFile) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return err
		}
	}
	switch m.ClazzID {
	case 0x42e6b860:
		flags, err := d.Uint32()
		if err != nil {
			return err
		}
		_ = flags
		m.OwnerId, err = d.Int64()
		if err != nil {
			return err
		}

		m.File, err = tg.DecodeInputFileClazz(d)
		if err != nil {
			return err
		}

		if (flags & (1 << 0)) != 0 {
			m.Thumb, err = tg.DecodeInputFileClazz(d)
			if err != nil {
				return err
			}
		}
		m.MimeType, err = d.String()
		if err != nil {
			return err
		}
		m.FileName, err = d.String()
		if err != nil {
			return err
		}

		return nil
	default:
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
	wrapper := iface.WithNameWrapper{ClazzName: ClazzName_media_uploadStickerFile, TLObject: m}
	return wrapper.String()
}

// Encode <--
func (m *TLMediaUploadStickerFile) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_media_uploadStickerFile, int(layer)); clazzId {
	case 0xacb624ed:
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
	default:
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_media_uploadStickerFile, layer)
	}
}

// Decode <--
func (m *TLMediaUploadStickerFile) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return err
		}
	}
	switch m.ClazzID {
	case 0xacb624ed:
		flags, err := d.Uint32()
		if err != nil {
			return err
		}
		_ = flags
		m.OwnerId, err = d.Int64()
		if err != nil {
			return err
		}

		m.File, err = tg.DecodeInputFileClazz(d)
		if err != nil {
			return err
		}

		if (flags & (1 << 0)) != 0 {
			m.Thumb, err = tg.DecodeInputFileClazz(d)
			if err != nil {
				return err
			}
		}
		m.MimeType, err = d.String()
		if err != nil {
			return err
		}
		m.FileName, err = d.String()
		if err != nil {
			return err
		}

		m.DocumentAttributeSticker, err = tg.DecodeDocumentAttributeClazz(d)
		if err != nil {
			return err
		}

		return nil
	default:
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
	wrapper := iface.WithNameWrapper{ClazzName: ClazzName_media_uploadRingtoneFile, TLObject: m}
	return wrapper.String()
}

// Encode <--
func (m *TLMediaUploadRingtoneFile) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_media_uploadRingtoneFile, int(layer)); clazzId {
	case 0x3dbab209:
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
	default:
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_media_uploadRingtoneFile, layer)
	}
}

// Decode <--
func (m *TLMediaUploadRingtoneFile) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return err
		}
	}
	switch m.ClazzID {
	case 0x3dbab209:
		flags, err := d.Uint32()
		if err != nil {
			return err
		}
		_ = flags
		m.OwnerId, err = d.Int64()
		if err != nil {
			return err
		}

		m.File, err = tg.DecodeInputFileClazz(d)
		if err != nil {
			return err
		}

		m.MimeType, err = d.String()
		if err != nil {
			return err
		}
		m.FileName, err = d.String()
		if err != nil {
			return err
		}

		return nil
	default:
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
	wrapper := iface.WithNameWrapper{ClazzName: ClazzName_media_uploadedProfilePhoto, TLObject: m}
	return wrapper.String()
}

// Encode <--
func (m *TLMediaUploadedProfilePhoto) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_media_uploadedProfilePhoto, int(layer)); clazzId {
	case 0x89d159d2:
		x.PutClazzID(0x89d159d2)

		x.PutInt64(m.OwnerId)
		x.PutInt64(m.PhotoId)

		return nil
	default:
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_media_uploadedProfilePhoto, layer)
	}
}

// Decode <--
func (m *TLMediaUploadedProfilePhoto) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return err
		}
	}
	switch m.ClazzID {
	case 0x89d159d2:
		m.OwnerId, err = d.Int64()
		if err != nil {
			return err
		}
		m.PhotoId, err = d.Int64()
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
