/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2026-present,  Teamgram Authors.
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package dfs

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
	wrapper := iface.WithNameWrapper{ClazzName: ClazzName_dfs_writeFilePartData, TLObject: m}
	return wrapper.String()
}

// Encode <--
func (m *TLDfsWriteFilePartData) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_dfs_writeFilePartData, int(layer)); clazzId {
	case 0x1a484107:
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
	default:
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_dfs_writeFilePartData, layer)
	}
}

// Decode <--
func (m *TLDfsWriteFilePartData) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return err
		}
	}
	switch m.ClazzID {
	case 0x1a484107:
		flags, err := d.Uint32()
		if err != nil {
			return err
		}
		_ = flags
		m.Creator, err = d.Int64()
		if err != nil {
			return err
		}
		m.FileId, err = d.Int64()
		if err != nil {
			return err
		}
		m.FilePart, err = d.Int32()
		if err != nil {
			return err
		}
		m.Bytes, err = d.Bytes()
		if err != nil {
			return err
		}
		if (flags & (1 << 0)) != 0 {
			m.Big = true
		}
		if (flags & (1 << 1)) != 0 {
			m.FileTotalParts = new(int32)
			*m.FileTotalParts, err = d.Int32()
			if err != nil {
				return err
			}
		}

		return nil
	default:
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// TLDfsUploadPhotoFileV2 <--
type TLDfsUploadPhotoFileV2 struct {
	ClazzID uint32            `json:"_id"`
	Creator int64             `json:"creator"`
	File    tg.InputFileClazz `json:"file"`
}

func (m *TLDfsUploadPhotoFileV2) String() string {
	wrapper := iface.WithNameWrapper{ClazzName: ClazzName_dfs_uploadPhotoFileV2, TLObject: m}
	return wrapper.String()
}

// Encode <--
func (m *TLDfsUploadPhotoFileV2) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_dfs_uploadPhotoFileV2, int(layer)); clazzId {
	case 0x2410d1a2:
		x.PutClazzID(0x2410d1a2)

		x.PutInt64(m.Creator)
		_ = m.File.Encode(x, layer)

		return nil
	default:
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_dfs_uploadPhotoFileV2, layer)
	}
}

// Decode <--
func (m *TLDfsUploadPhotoFileV2) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return err
		}
	}
	switch m.ClazzID {
	case 0x2410d1a2:
		m.Creator, err = d.Int64()
		if err != nil {
			return err
		}

		m.File, err = tg.DecodeInputFileClazz(d)
		if err != nil {
			return err
		}

		return nil
	default:
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// TLDfsUploadProfilePhotoFileV2 <--
type TLDfsUploadProfilePhotoFileV2 struct {
	ClazzID          uint32            `json:"_id"`
	Creator          int64             `json:"creator"`
	File             tg.InputFileClazz `json:"file"`
	Video            tg.InputFileClazz `json:"video"`
	VideoStartTs     *float64          `json:"video_start_ts"`
	VideoEmojiMarkup tg.VideoSizeClazz `json:"video_emoji_markup"`
}

func (m *TLDfsUploadProfilePhotoFileV2) String() string {
	wrapper := iface.WithNameWrapper{ClazzName: ClazzName_dfs_uploadProfilePhotoFileV2, TLObject: m}
	return wrapper.String()
}

// Encode <--
func (m *TLDfsUploadProfilePhotoFileV2) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_dfs_uploadProfilePhotoFileV2, int(layer)); clazzId {
	case 0x872313d8:
		x.PutClazzID(0x872313d8)

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

		if m.VideoEmojiMarkup != nil {
			_ = m.VideoEmojiMarkup.Encode(x, layer)
		}

		return nil
	default:
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_dfs_uploadProfilePhotoFileV2, layer)
	}
}

// Decode <--
func (m *TLDfsUploadProfilePhotoFileV2) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return err
		}
	}
	switch m.ClazzID {
	case 0x872313d8:
		flags, err := d.Uint32()
		if err != nil {
			return err
		}
		_ = flags
		m.Creator, err = d.Int64()
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

// TLDfsUploadEncryptedFileV2 <--
type TLDfsUploadEncryptedFileV2 struct {
	ClazzID uint32                     `json:"_id"`
	Creator int64                      `json:"creator"`
	File    tg.InputEncryptedFileClazz `json:"file"`
}

func (m *TLDfsUploadEncryptedFileV2) String() string {
	wrapper := iface.WithNameWrapper{ClazzName: ClazzName_dfs_uploadEncryptedFileV2, TLObject: m}
	return wrapper.String()
}

// Encode <--
func (m *TLDfsUploadEncryptedFileV2) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_dfs_uploadEncryptedFileV2, int(layer)); clazzId {
	case 0x79d3c523:
		x.PutClazzID(0x79d3c523)

		x.PutInt64(m.Creator)
		_ = m.File.Encode(x, layer)

		return nil
	default:
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_dfs_uploadEncryptedFileV2, layer)
	}
}

// Decode <--
func (m *TLDfsUploadEncryptedFileV2) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return err
		}
	}
	switch m.ClazzID {
	case 0x79d3c523:
		m.Creator, err = d.Int64()
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

// TLDfsDownloadFile <--
type TLDfsDownloadFile struct {
	ClazzID  uint32                    `json:"_id"`
	Location tg.InputFileLocationClazz `json:"location"`
	Offset   int64                     `json:"offset"`
	Limit    int32                     `json:"limit"`
}

func (m *TLDfsDownloadFile) String() string {
	wrapper := iface.WithNameWrapper{ClazzName: ClazzName_dfs_downloadFile, TLObject: m}
	return wrapper.String()
}

// Encode <--
func (m *TLDfsDownloadFile) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_dfs_downloadFile, int(layer)); clazzId {
	case 0xd6bfee3e:
		x.PutClazzID(0xd6bfee3e)

		_ = m.Location.Encode(x, layer)
		x.PutInt64(m.Offset)
		x.PutInt32(m.Limit)

		return nil
	default:
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_dfs_downloadFile, layer)
	}
}

// Decode <--
func (m *TLDfsDownloadFile) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return err
		}
	}
	switch m.ClazzID {
	case 0xd6bfee3e:

		m.Location, err = tg.DecodeInputFileLocationClazz(d)
		if err != nil {
			return err
		}

		m.Offset, err = d.Int64()
		if err != nil {
			return err
		}
		m.Limit, err = d.Int32()
		if err != nil {
			return err
		}

		return nil
	default:
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// TLDfsUploadDocumentFileV2 <--
type TLDfsUploadDocumentFileV2 struct {
	ClazzID uint32             `json:"_id"`
	Creator int64              `json:"creator"`
	Media   tg.InputMediaClazz `json:"media"`
}

func (m *TLDfsUploadDocumentFileV2) String() string {
	wrapper := iface.WithNameWrapper{ClazzName: ClazzName_dfs_uploadDocumentFileV2, TLObject: m}
	return wrapper.String()
}

// Encode <--
func (m *TLDfsUploadDocumentFileV2) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_dfs_uploadDocumentFileV2, int(layer)); clazzId {
	case 0x76336db7:
		x.PutClazzID(0x76336db7)

		x.PutInt64(m.Creator)
		_ = m.Media.Encode(x, layer)

		return nil
	default:
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_dfs_uploadDocumentFileV2, layer)
	}
}

// Decode <--
func (m *TLDfsUploadDocumentFileV2) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return err
		}
	}
	switch m.ClazzID {
	case 0x76336db7:
		m.Creator, err = d.Int64()
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

// TLDfsUploadGifDocumentMedia <--
type TLDfsUploadGifDocumentMedia struct {
	ClazzID uint32             `json:"_id"`
	Creator int64              `json:"creator"`
	Media   tg.InputMediaClazz `json:"media"`
}

func (m *TLDfsUploadGifDocumentMedia) String() string {
	wrapper := iface.WithNameWrapper{ClazzName: ClazzName_dfs_uploadGifDocumentMedia, TLObject: m}
	return wrapper.String()
}

// Encode <--
func (m *TLDfsUploadGifDocumentMedia) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_dfs_uploadGifDocumentMedia, int(layer)); clazzId {
	case 0x41c4cd00:
		x.PutClazzID(0x41c4cd00)

		x.PutInt64(m.Creator)
		_ = m.Media.Encode(x, layer)

		return nil
	default:
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_dfs_uploadGifDocumentMedia, layer)
	}
}

// Decode <--
func (m *TLDfsUploadGifDocumentMedia) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return err
		}
	}
	switch m.ClazzID {
	case 0x41c4cd00:
		m.Creator, err = d.Int64()
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

// TLDfsUploadMp4DocumentMedia <--
type TLDfsUploadMp4DocumentMedia struct {
	ClazzID uint32             `json:"_id"`
	Creator int64              `json:"creator"`
	Media   tg.InputMediaClazz `json:"media"`
}

func (m *TLDfsUploadMp4DocumentMedia) String() string {
	wrapper := iface.WithNameWrapper{ClazzName: ClazzName_dfs_uploadMp4DocumentMedia, TLObject: m}
	return wrapper.String()
}

// Encode <--
func (m *TLDfsUploadMp4DocumentMedia) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_dfs_uploadMp4DocumentMedia, int(layer)); clazzId {
	case 0xa2a4f818:
		x.PutClazzID(0xa2a4f818)

		x.PutInt64(m.Creator)
		_ = m.Media.Encode(x, layer)

		return nil
	default:
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_dfs_uploadMp4DocumentMedia, layer)
	}
}

// Decode <--
func (m *TLDfsUploadMp4DocumentMedia) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return err
		}
	}
	switch m.ClazzID {
	case 0xa2a4f818:
		m.Creator, err = d.Int64()
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

// TLDfsUploadWallPaperFile <--
type TLDfsUploadWallPaperFile struct {
	ClazzID  uint32            `json:"_id"`
	Creator  int64             `json:"creator"`
	File     tg.InputFileClazz `json:"file"`
	MimeType string            `json:"mime_type"`
	Admin    tg.BoolClazz      `json:"admin"`
}

func (m *TLDfsUploadWallPaperFile) String() string {
	wrapper := iface.WithNameWrapper{ClazzName: ClazzName_dfs_uploadWallPaperFile, TLObject: m}
	return wrapper.String()
}

// Encode <--
func (m *TLDfsUploadWallPaperFile) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_dfs_uploadWallPaperFile, int(layer)); clazzId {
	case 0xc1a61056:
		x.PutClazzID(0xc1a61056)

		x.PutInt64(m.Creator)
		_ = m.File.Encode(x, layer)
		x.PutString(m.MimeType)
		_ = m.Admin.Encode(x, layer)

		return nil
	default:
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_dfs_uploadWallPaperFile, layer)
	}
}

// Decode <--
func (m *TLDfsUploadWallPaperFile) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return err
		}
	}
	switch m.ClazzID {
	case 0xc1a61056:
		m.Creator, err = d.Int64()
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

// TLDfsUploadThemeFile <--
type TLDfsUploadThemeFile struct {
	ClazzID  uint32            `json:"_id"`
	Creator  int64             `json:"creator"`
	File     tg.InputFileClazz `json:"file"`
	Thumb    tg.InputFileClazz `json:"thumb"`
	MimeType string            `json:"mime_type"`
	FileName string            `json:"file_name"`
}

func (m *TLDfsUploadThemeFile) String() string {
	wrapper := iface.WithNameWrapper{ClazzName: ClazzName_dfs_uploadThemeFile, TLObject: m}
	return wrapper.String()
}

// Encode <--
func (m *TLDfsUploadThemeFile) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_dfs_uploadThemeFile, int(layer)); clazzId {
	case 0xdea64f97:
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
	default:
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_dfs_uploadThemeFile, layer)
	}
}

// Decode <--
func (m *TLDfsUploadThemeFile) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return err
		}
	}
	switch m.ClazzID {
	case 0xdea64f97:
		flags, err := d.Uint32()
		if err != nil {
			return err
		}
		_ = flags
		m.Creator, err = d.Int64()
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

// TLDfsUploadRingtoneFile <--
type TLDfsUploadRingtoneFile struct {
	ClazzID  uint32            `json:"_id"`
	Creator  int64             `json:"creator"`
	File     tg.InputFileClazz `json:"file"`
	MimeType string            `json:"mime_type"`
	FileName string            `json:"file_name"`
}

func (m *TLDfsUploadRingtoneFile) String() string {
	wrapper := iface.WithNameWrapper{ClazzName: ClazzName_dfs_uploadRingtoneFile, TLObject: m}
	return wrapper.String()
}

// Encode <--
func (m *TLDfsUploadRingtoneFile) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_dfs_uploadRingtoneFile, int(layer)); clazzId {
	case 0x2b3c5b1:
		x.PutClazzID(0x2b3c5b1)

		x.PutInt64(m.Creator)
		_ = m.File.Encode(x, layer)
		x.PutString(m.MimeType)
		x.PutString(m.FileName)

		return nil
	default:
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_dfs_uploadRingtoneFile, layer)
	}
}

// Decode <--
func (m *TLDfsUploadRingtoneFile) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return err
		}
	}
	switch m.ClazzID {
	case 0x2b3c5b1:
		m.Creator, err = d.Int64()
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

// TLDfsUploadedProfilePhoto <--
type TLDfsUploadedProfilePhoto struct {
	ClazzID uint32 `json:"_id"`
	Creator int64  `json:"creator"`
	PhotoId int64  `json:"photo_id"`
}

func (m *TLDfsUploadedProfilePhoto) String() string {
	wrapper := iface.WithNameWrapper{ClazzName: ClazzName_dfs_uploadedProfilePhoto, TLObject: m}
	return wrapper.String()
}

// Encode <--
func (m *TLDfsUploadedProfilePhoto) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_dfs_uploadedProfilePhoto, int(layer)); clazzId {
	case 0xa3aa2874:
		x.PutClazzID(0xa3aa2874)

		x.PutInt64(m.Creator)
		x.PutInt64(m.PhotoId)

		return nil
	default:
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_dfs_uploadedProfilePhoto, layer)
	}
}

// Decode <--
func (m *TLDfsUploadedProfilePhoto) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return err
		}
	}
	switch m.ClazzID {
	case 0xa3aa2874:
		m.Creator, err = d.Int64()
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
