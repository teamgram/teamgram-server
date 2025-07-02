/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2025-present,  Teamgram Authors.
 *  All rights reserved.
 *
 * Author: Benqi (wubenqi@gmail.com)
 */

package dfs

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

// TLDfsWriteFilePartData <--
type TLDfsWriteFilePartData struct {
	ClazzID        uint32 `json:"_id"`
	Creator        int64  `json:"creator"`
	FileId         int64  `json:"file_id"`
	FilePart       int32  `json:"file_part"`
	Bytes          []byte `json:"bytes"`
	Big            bool   `json:"big"`
	FileTotalParts *int32 `json:"file_total_parts"`
}

func (m *TLDfsWriteFilePartData) String() string {
	wrapper := iface.WithNameWrapper{"", m}
	return wrapper.String()
}

// Encode <--
func (m *TLDfsWriteFilePartData) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0x1a484107: func() error {
			x.PutClazzID(0x1a484107)

			// set flags
			var getFlags = func() uint32 {
				var flags uint32 = 0

				if m.Big == true {
					flags |= 1 << 0
				}
				if m.FileTotalParts != nil {
					flags |= 1 << 1
				}

				return flags
			}

			// set flags
			var flags = getFlags()
			x.PutUint32(flags)
			x.PutInt64(m.Creator)
			x.PutInt64(m.FileId)
			x.PutInt32(m.FilePart)
			x.PutBytes(m.Bytes)
			if m.FileTotalParts != nil {
				x.PutInt32(*m.FileTotalParts)
			}

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_dfs_writeFilePartData, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_dfs_writeFilePartData, layer)
	}
}

// Decode <--
func (m *TLDfsWriteFilePartData) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0x1a484107: func() (err error) {
			flags, _ := d.Uint32()
			_ = flags
			m.Creator, err = d.Int64()
			m.FileId, err = d.Int64()
			m.FilePart, err = d.Int32()
			m.Bytes, err = d.Bytes()
			if (flags & (1 << 0)) != 0 {
				m.Big = true
			}
			if (flags & (1 << 1)) != 0 {
				m.FileTotalParts = new(int32)
				*m.FileTotalParts, err = d.Int32()
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

// TLDfsUploadPhotoFileV2 <--
type TLDfsUploadPhotoFileV2 struct {
	ClazzID uint32        `json:"_id"`
	Creator int64         `json:"creator"`
	File    *tg.InputFile `json:"file"`
}

func (m *TLDfsUploadPhotoFileV2) String() string {
	wrapper := iface.WithNameWrapper{"", m}
	return wrapper.String()
}

// Encode <--
func (m *TLDfsUploadPhotoFileV2) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0x2410d1a2: func() error {
			x.PutClazzID(0x2410d1a2)

			x.PutInt64(m.Creator)
			_ = m.File.Encode(x, layer)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_dfs_uploadPhotoFileV2, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_dfs_uploadPhotoFileV2, layer)
	}
}

// Decode <--
func (m *TLDfsUploadPhotoFileV2) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0x2410d1a2: func() (err error) {
			m.Creator, err = d.Int64()

			m2 := &tg.InputFile{}
			_ = m2.Decode(d)
			m.File = m2

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

// TLDfsUploadProfilePhotoFileV2 <--
type TLDfsUploadProfilePhotoFileV2 struct {
	ClazzID      uint32        `json:"_id"`
	Creator      int64         `json:"creator"`
	File         *tg.InputFile `json:"file"`
	Video        *tg.InputFile `json:"video"`
	VideoStartTs *float64      `json:"video_start_ts"`
}

func (m *TLDfsUploadProfilePhotoFileV2) String() string {
	wrapper := iface.WithNameWrapper{"", m}
	return wrapper.String()
}

// Encode <--
func (m *TLDfsUploadProfilePhotoFileV2) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0xcc1da2b2: func() error {
			x.PutClazzID(0xcc1da2b2)

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
			x.PutInt64(m.Creator)
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

	clazzId := iface.GetClazzIDByName(ClazzName_dfs_uploadProfilePhotoFileV2, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_dfs_uploadProfilePhotoFileV2, layer)
	}
}

// Decode <--
func (m *TLDfsUploadProfilePhotoFileV2) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0xcc1da2b2: func() (err error) {
			flags, _ := d.Uint32()
			_ = flags
			m.Creator, err = d.Int64()
			if (flags & (1 << 0)) != 0 {
				m3 := &tg.InputFile{}
				_ = m3.Decode(d)
				m.File = m3
			}
			if (flags & (1 << 1)) != 0 {
				m4 := &tg.InputFile{}
				_ = m4.Decode(d)
				m.Video = m4
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

// TLDfsUploadEncryptedFileV2 <--
type TLDfsUploadEncryptedFileV2 struct {
	ClazzID uint32                 `json:"_id"`
	Creator int64                  `json:"creator"`
	File    *tg.InputEncryptedFile `json:"file"`
}

func (m *TLDfsUploadEncryptedFileV2) String() string {
	wrapper := iface.WithNameWrapper{"", m}
	return wrapper.String()
}

// Encode <--
func (m *TLDfsUploadEncryptedFileV2) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0x79d3c523: func() error {
			x.PutClazzID(0x79d3c523)

			x.PutInt64(m.Creator)
			_ = m.File.Encode(x, layer)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_dfs_uploadEncryptedFileV2, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_dfs_uploadEncryptedFileV2, layer)
	}
}

// Decode <--
func (m *TLDfsUploadEncryptedFileV2) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0x79d3c523: func() (err error) {
			m.Creator, err = d.Int64()

			m2 := &tg.InputEncryptedFile{}
			_ = m2.Decode(d)
			m.File = m2

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

// TLDfsDownloadFile <--
type TLDfsDownloadFile struct {
	ClazzID  uint32                `json:"_id"`
	Location *tg.InputFileLocation `json:"location"`
	Offset   int64                 `json:"offset"`
	Limit    int32                 `json:"limit"`
}

func (m *TLDfsDownloadFile) String() string {
	wrapper := iface.WithNameWrapper{"", m}
	return wrapper.String()
}

// Encode <--
func (m *TLDfsDownloadFile) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0xd6bfee3e: func() error {
			x.PutClazzID(0xd6bfee3e)

			_ = m.Location.Encode(x, layer)
			x.PutInt64(m.Offset)
			x.PutInt32(m.Limit)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_dfs_downloadFile, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_dfs_downloadFile, layer)
	}
}

// Decode <--
func (m *TLDfsDownloadFile) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0xd6bfee3e: func() (err error) {

			m1 := &tg.InputFileLocation{}
			_ = m1.Decode(d)
			m.Location = m1

			m.Offset, err = d.Int64()
			m.Limit, err = d.Int32()

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

// TLDfsUploadDocumentFileV2 <--
type TLDfsUploadDocumentFileV2 struct {
	ClazzID uint32         `json:"_id"`
	Creator int64          `json:"creator"`
	Media   *tg.InputMedia `json:"media"`
}

func (m *TLDfsUploadDocumentFileV2) String() string {
	wrapper := iface.WithNameWrapper{"", m}
	return wrapper.String()
}

// Encode <--
func (m *TLDfsUploadDocumentFileV2) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0x76336db7: func() error {
			x.PutClazzID(0x76336db7)

			x.PutInt64(m.Creator)
			_ = m.Media.Encode(x, layer)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_dfs_uploadDocumentFileV2, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_dfs_uploadDocumentFileV2, layer)
	}
}

// Decode <--
func (m *TLDfsUploadDocumentFileV2) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0x76336db7: func() (err error) {
			m.Creator, err = d.Int64()

			m2 := &tg.InputMedia{}
			_ = m2.Decode(d)
			m.Media = m2

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

// TLDfsUploadGifDocumentMedia <--
type TLDfsUploadGifDocumentMedia struct {
	ClazzID uint32         `json:"_id"`
	Creator int64          `json:"creator"`
	Media   *tg.InputMedia `json:"media"`
}

func (m *TLDfsUploadGifDocumentMedia) String() string {
	wrapper := iface.WithNameWrapper{"", m}
	return wrapper.String()
}

// Encode <--
func (m *TLDfsUploadGifDocumentMedia) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0x41c4cd00: func() error {
			x.PutClazzID(0x41c4cd00)

			x.PutInt64(m.Creator)
			_ = m.Media.Encode(x, layer)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_dfs_uploadGifDocumentMedia, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_dfs_uploadGifDocumentMedia, layer)
	}
}

// Decode <--
func (m *TLDfsUploadGifDocumentMedia) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0x41c4cd00: func() (err error) {
			m.Creator, err = d.Int64()

			m2 := &tg.InputMedia{}
			_ = m2.Decode(d)
			m.Media = m2

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

// TLDfsUploadMp4DocumentMedia <--
type TLDfsUploadMp4DocumentMedia struct {
	ClazzID uint32         `json:"_id"`
	Creator int64          `json:"creator"`
	Media   *tg.InputMedia `json:"media"`
}

func (m *TLDfsUploadMp4DocumentMedia) String() string {
	wrapper := iface.WithNameWrapper{"", m}
	return wrapper.String()
}

// Encode <--
func (m *TLDfsUploadMp4DocumentMedia) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0xa2a4f818: func() error {
			x.PutClazzID(0xa2a4f818)

			x.PutInt64(m.Creator)
			_ = m.Media.Encode(x, layer)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_dfs_uploadMp4DocumentMedia, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_dfs_uploadMp4DocumentMedia, layer)
	}
}

// Decode <--
func (m *TLDfsUploadMp4DocumentMedia) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0xa2a4f818: func() (err error) {
			m.Creator, err = d.Int64()

			m2 := &tg.InputMedia{}
			_ = m2.Decode(d)
			m.Media = m2

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

// TLDfsUploadWallPaperFile <--
type TLDfsUploadWallPaperFile struct {
	ClazzID  uint32        `json:"_id"`
	Creator  int64         `json:"creator"`
	File     *tg.InputFile `json:"file"`
	MimeType string        `json:"mime_type"`
	Admin    *tg.Bool      `json:"admin"`
}

func (m *TLDfsUploadWallPaperFile) String() string {
	wrapper := iface.WithNameWrapper{"", m}
	return wrapper.String()
}

// Encode <--
func (m *TLDfsUploadWallPaperFile) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0xc1a61056: func() error {
			x.PutClazzID(0xc1a61056)

			x.PutInt64(m.Creator)
			_ = m.File.Encode(x, layer)
			x.PutString(m.MimeType)
			_ = m.Admin.Encode(x, layer)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_dfs_uploadWallPaperFile, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_dfs_uploadWallPaperFile, layer)
	}
}

// Decode <--
func (m *TLDfsUploadWallPaperFile) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0xc1a61056: func() (err error) {
			m.Creator, err = d.Int64()

			m2 := &tg.InputFile{}
			_ = m2.Decode(d)
			m.File = m2

			m.MimeType, err = d.String()

			m4 := &tg.Bool{}
			_ = m4.Decode(d)
			m.Admin = m4

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

// TLDfsUploadThemeFile <--
type TLDfsUploadThemeFile struct {
	ClazzID  uint32        `json:"_id"`
	Creator  int64         `json:"creator"`
	File     *tg.InputFile `json:"file"`
	Thumb    *tg.InputFile `json:"thumb"`
	MimeType string        `json:"mime_type"`
	FileName string        `json:"file_name"`
}

func (m *TLDfsUploadThemeFile) String() string {
	wrapper := iface.WithNameWrapper{"", m}
	return wrapper.String()
}

// Encode <--
func (m *TLDfsUploadThemeFile) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0xdea64f97: func() error {
			x.PutClazzID(0xdea64f97)

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
			x.PutInt64(m.Creator)
			_ = m.File.Encode(x, layer)
			if m.Thumb != nil {
				_ = m.Thumb.Encode(x, layer)
			}

			x.PutString(m.MimeType)
			x.PutString(m.FileName)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_dfs_uploadThemeFile, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_dfs_uploadThemeFile, layer)
	}
}

// Decode <--
func (m *TLDfsUploadThemeFile) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0xdea64f97: func() (err error) {
			flags, _ := d.Uint32()
			_ = flags
			m.Creator, err = d.Int64()

			m3 := &tg.InputFile{}
			_ = m3.Decode(d)
			m.File = m3

			if (flags & (1 << 0)) != 0 {
				m4 := &tg.InputFile{}
				_ = m4.Decode(d)
				m.Thumb = m4
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

// TLDfsUploadRingtoneFile <--
type TLDfsUploadRingtoneFile struct {
	ClazzID  uint32        `json:"_id"`
	Creator  int64         `json:"creator"`
	File     *tg.InputFile `json:"file"`
	MimeType string        `json:"mime_type"`
	FileName string        `json:"file_name"`
}

func (m *TLDfsUploadRingtoneFile) String() string {
	wrapper := iface.WithNameWrapper{"", m}
	return wrapper.String()
}

// Encode <--
func (m *TLDfsUploadRingtoneFile) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0x2b3c5b1: func() error {
			x.PutClazzID(0x2b3c5b1)

			x.PutInt64(m.Creator)
			_ = m.File.Encode(x, layer)
			x.PutString(m.MimeType)
			x.PutString(m.FileName)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_dfs_uploadRingtoneFile, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_dfs_uploadRingtoneFile, layer)
	}
}

// Decode <--
func (m *TLDfsUploadRingtoneFile) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0x2b3c5b1: func() (err error) {
			m.Creator, err = d.Int64()

			m2 := &tg.InputFile{}
			_ = m2.Decode(d)
			m.File = m2

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

// TLDfsUploadedProfilePhoto <--
type TLDfsUploadedProfilePhoto struct {
	ClazzID uint32 `json:"_id"`
	Creator int64  `json:"creator"`
	PhotoId int64  `json:"photo_id"`
}

func (m *TLDfsUploadedProfilePhoto) String() string {
	wrapper := iface.WithNameWrapper{"", m}
	return wrapper.String()
}

// Encode <--
func (m *TLDfsUploadedProfilePhoto) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0xa3aa2874: func() error {
			x.PutClazzID(0xa3aa2874)

			x.PutInt64(m.Creator)
			x.PutInt64(m.PhotoId)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_dfs_uploadedProfilePhoto, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_dfs_uploadedProfilePhoto, layer)
	}
}

// Decode <--
func (m *TLDfsUploadedProfilePhoto) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0xa3aa2874: func() (err error) {
			m.Creator, err = d.Int64()
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

// ----------------------------------------------------------------------------
// rpc

type RPCDfs interface {
	DfsWriteFilePartData(ctx context.Context, in *TLDfsWriteFilePartData) (*tg.Bool, error)
	DfsUploadPhotoFileV2(ctx context.Context, in *TLDfsUploadPhotoFileV2) (*tg.Photo, error)
	DfsUploadProfilePhotoFileV2(ctx context.Context, in *TLDfsUploadProfilePhotoFileV2) (*tg.Photo, error)
	DfsUploadEncryptedFileV2(ctx context.Context, in *TLDfsUploadEncryptedFileV2) (*tg.EncryptedFile, error)
	DfsDownloadFile(ctx context.Context, in *TLDfsDownloadFile) (*tg.UploadFile, error)
	DfsUploadDocumentFileV2(ctx context.Context, in *TLDfsUploadDocumentFileV2) (*tg.Document, error)
	DfsUploadGifDocumentMedia(ctx context.Context, in *TLDfsUploadGifDocumentMedia) (*tg.Document, error)
	DfsUploadMp4DocumentMedia(ctx context.Context, in *TLDfsUploadMp4DocumentMedia) (*tg.Document, error)
	DfsUploadWallPaperFile(ctx context.Context, in *TLDfsUploadWallPaperFile) (*tg.Document, error)
	DfsUploadThemeFile(ctx context.Context, in *TLDfsUploadThemeFile) (*tg.Document, error)
	DfsUploadRingtoneFile(ctx context.Context, in *TLDfsUploadRingtoneFile) (*tg.Document, error)
	DfsUploadedProfilePhoto(ctx context.Context, in *TLDfsUploadedProfilePhoto) (*tg.Photo, error)
}
