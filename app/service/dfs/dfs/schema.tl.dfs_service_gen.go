/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2026-present, The Teamgram Authors (https://teamgram.net).
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

// TLDfsCommitUpload <--
type TLDfsCommitUpload struct {
	ClazzID         uint32            `json:"_id"`
	UploadSessionId string            `json:"upload_session_id"`
	OwnerId         int64             `json:"owner_id"`
	File            tg.InputFileClazz `json:"file"`
	Purpose         string            `json:"purpose"`
}

func (m *TLDfsCommitUpload) String() string {
	return iface.DebugStringWithName(ClazzName_dfs_commitUpload, m)
}

// Encode <--
func (m *TLDfsCommitUpload) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_dfs_commitUpload, int(layer)); clazzId {
	case 0xdddb9d2c:
		x.PutClazzID(0xdddb9d2c)

		x.PutString(m.UploadSessionId)
		x.PutInt64(m.OwnerId)
		if m.File == nil {
			return fmt.Errorf("unable to encode dfs_commitUpload#0xdddb9d2c: field file is nil")
		}
		if err := m.File.Encode(x, layer); err != nil {
			return fmt.Errorf("unable to encode dfs_commitUpload#0xdddb9d2c: field file: %w", err)
		}
		x.PutString(m.Purpose)

		return nil
	default:
		return fmt.Errorf("unable to encode dfs_commitUpload: unsupported layer %d", layer)
	}
}

// Decode <--
func (m *TLDfsCommitUpload) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return fmt.Errorf("unable to decode dfs_commitUpload: constructor: %w", err)
		}
	}
	switch m.ClazzID {
	case 0xdddb9d2c:
		m.UploadSessionId, err = d.String()
		if err != nil {
			return fmt.Errorf("unable to decode dfs_commitUpload#0xdddb9d2c: field upload_session_id: %w", err)
		}
		m.OwnerId, err = d.Int64()
		if err != nil {
			return fmt.Errorf("unable to decode dfs_commitUpload#0xdddb9d2c: field owner_id: %w", err)
		}

		m.File, err = tg.DecodeInputFileClazz(d)
		if err != nil {
			return fmt.Errorf("unable to decode dfs_commitUpload#0xdddb9d2c: field file: %w", err)
		}

		m.Purpose, err = d.String()
		if err != nil {
			return fmt.Errorf("unable to decode dfs_commitUpload#0xdddb9d2c: field purpose: %w", err)
		}

		return nil
	default:
		return fmt.Errorf("unable to decode dfs_commitUpload: invalid constructor %x", m.ClazzID)
	}
}

// TLDfsPutFile <--
type TLDfsPutFile struct {
	ClazzID  uint32 `json:"_id"`
	OwnerId  int64  `json:"owner_id"`
	Purpose  string `json:"purpose"`
	FileName string `json:"file_name"`
	MimeType string `json:"mime_type"`
	Bytes    []byte `json:"bytes"`
}

func (m *TLDfsPutFile) String() string {
	return iface.DebugStringWithName(ClazzName_dfs_putFile, m)
}

// Encode <--
func (m *TLDfsPutFile) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_dfs_putFile, int(layer)); clazzId {
	case 0x6e20c3e7:
		x.PutClazzID(0x6e20c3e7)

		x.PutInt64(m.OwnerId)
		x.PutString(m.Purpose)
		x.PutString(m.FileName)
		x.PutString(m.MimeType)
		x.PutBytes(m.Bytes)

		return nil
	default:
		return fmt.Errorf("unable to encode dfs_putFile: unsupported layer %d", layer)
	}
}

// Decode <--
func (m *TLDfsPutFile) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return fmt.Errorf("unable to decode dfs_putFile: constructor: %w", err)
		}
	}
	switch m.ClazzID {
	case 0x6e20c3e7:
		m.OwnerId, err = d.Int64()
		if err != nil {
			return fmt.Errorf("unable to decode dfs_putFile#0x6e20c3e7: field owner_id: %w", err)
		}
		m.Purpose, err = d.String()
		if err != nil {
			return fmt.Errorf("unable to decode dfs_putFile#0x6e20c3e7: field purpose: %w", err)
		}
		m.FileName, err = d.String()
		if err != nil {
			return fmt.Errorf("unable to decode dfs_putFile#0x6e20c3e7: field file_name: %w", err)
		}
		m.MimeType, err = d.String()
		if err != nil {
			return fmt.Errorf("unable to decode dfs_putFile#0x6e20c3e7: field mime_type: %w", err)
		}
		m.Bytes, err = d.Bytes()
		if err != nil {
			return fmt.Errorf("unable to decode dfs_putFile#0x6e20c3e7: field bytes: %w", err)
		}

		return nil
	default:
		return fmt.Errorf("unable to decode dfs_putFile: invalid constructor %x", m.ClazzID)
	}
}

// TLDfsGetFileByReadLease <--
type TLDfsGetFileByReadLease struct {
	ClazzID   uint32 `json:"_id"`
	ReadLease []byte `json:"read_lease"`
	Offset    int64  `json:"offset"`
	Limit     int32  `json:"limit"`
}

func (m *TLDfsGetFileByReadLease) String() string {
	return iface.DebugStringWithName(ClazzName_dfs_getFileByReadLease, m)
}

// Encode <--
func (m *TLDfsGetFileByReadLease) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_dfs_getFileByReadLease, int(layer)); clazzId {
	case 0x86c7c115:
		x.PutClazzID(0x86c7c115)

		x.PutBytes(m.ReadLease)
		x.PutInt64(m.Offset)
		x.PutInt32(m.Limit)

		return nil
	default:
		return fmt.Errorf("unable to encode dfs_getFileByReadLease: unsupported layer %d", layer)
	}
}

// Decode <--
func (m *TLDfsGetFileByReadLease) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return fmt.Errorf("unable to decode dfs_getFileByReadLease: constructor: %w", err)
		}
	}
	switch m.ClazzID {
	case 0x86c7c115:
		m.ReadLease, err = d.Bytes()
		if err != nil {
			return fmt.Errorf("unable to decode dfs_getFileByReadLease#0x86c7c115: field read_lease: %w", err)
		}
		m.Offset, err = d.Int64()
		if err != nil {
			return fmt.Errorf("unable to decode dfs_getFileByReadLease#0x86c7c115: field offset: %w", err)
		}
		m.Limit, err = d.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode dfs_getFileByReadLease#0x86c7c115: field limit: %w", err)
		}

		return nil
	default:
		return fmt.Errorf("unable to decode dfs_getFileByReadLease: invalid constructor %x", m.ClazzID)
	}
}

// TLDfsGetFileHashesByReadLease <--
type TLDfsGetFileHashesByReadLease struct {
	ClazzID   uint32 `json:"_id"`
	ReadLease []byte `json:"read_lease"`
	Offset    int64  `json:"offset"`
	Limit     int32  `json:"limit"`
}

func (m *TLDfsGetFileHashesByReadLease) String() string {
	return iface.DebugStringWithName(ClazzName_dfs_getFileHashesByReadLease, m)
}

// Encode <--
func (m *TLDfsGetFileHashesByReadLease) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_dfs_getFileHashesByReadLease, int(layer)); clazzId {
	case 0xff974b78:
		x.PutClazzID(0xff974b78)

		x.PutBytes(m.ReadLease)
		x.PutInt64(m.Offset)
		x.PutInt32(m.Limit)

		return nil
	default:
		return fmt.Errorf("unable to encode dfs_getFileHashesByReadLease: unsupported layer %d", layer)
	}
}

// Decode <--
func (m *TLDfsGetFileHashesByReadLease) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return fmt.Errorf("unable to decode dfs_getFileHashesByReadLease: constructor: %w", err)
		}
	}
	switch m.ClazzID {
	case 0xff974b78:
		m.ReadLease, err = d.Bytes()
		if err != nil {
			return fmt.Errorf("unable to decode dfs_getFileHashesByReadLease#0xff974b78: field read_lease: %w", err)
		}
		m.Offset, err = d.Int64()
		if err != nil {
			return fmt.Errorf("unable to decode dfs_getFileHashesByReadLease#0xff974b78: field offset: %w", err)
		}
		m.Limit, err = d.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode dfs_getFileHashesByReadLease#0xff974b78: field limit: %w", err)
		}

		return nil
	default:
		return fmt.Errorf("unable to decode dfs_getFileHashesByReadLease: invalid constructor %x", m.ClazzID)
	}
}

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
	return iface.DebugStringWithName(ClazzName_dfs_writeFilePartData, m)
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
		return fmt.Errorf("unable to encode dfs_writeFilePartData: unsupported layer %d", layer)
	}
}

// Decode <--
func (m *TLDfsWriteFilePartData) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return fmt.Errorf("unable to decode dfs_writeFilePartData: constructor: %w", err)
		}
	}
	switch m.ClazzID {
	case 0x1a484107:
		flags, err := d.Uint32()
		if err != nil {
			return fmt.Errorf("unable to decode dfs_writeFilePartData: field flags: %w", err)
		}
		_ = flags
		m.Creator, err = d.Int64()
		if err != nil {
			return fmt.Errorf("unable to decode dfs_writeFilePartData#0x1a484107: field creator: %w", err)
		}
		m.FileId, err = d.Int64()
		if err != nil {
			return fmt.Errorf("unable to decode dfs_writeFilePartData#0x1a484107: field file_id: %w", err)
		}
		m.FilePart, err = d.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode dfs_writeFilePartData#0x1a484107: field file_part: %w", err)
		}
		m.Bytes, err = d.Bytes()
		if err != nil {
			return fmt.Errorf("unable to decode dfs_writeFilePartData#0x1a484107: field bytes: %w", err)
		}
		if (flags & (1 << 0)) != 0 {
			m.Big = true
		}
		if (flags & (1 << 1)) != 0 {
			m.FileTotalParts = new(int32)
			*m.FileTotalParts, err = d.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode dfs_writeFilePartData#0x1a484107: field file_total_parts: %w", err)
			}
		}

		return nil
	default:
		return fmt.Errorf("unable to decode dfs_writeFilePartData: invalid constructor %x", m.ClazzID)
	}
}

// TLDfsUploadPhotoFileV2 <--
type TLDfsUploadPhotoFileV2 struct {
	ClazzID uint32            `json:"_id"`
	Creator int64             `json:"creator"`
	File    tg.InputFileClazz `json:"file"`
}

func (m *TLDfsUploadPhotoFileV2) String() string {
	return iface.DebugStringWithName(ClazzName_dfs_uploadPhotoFileV2, m)
}

// Encode <--
func (m *TLDfsUploadPhotoFileV2) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_dfs_uploadPhotoFileV2, int(layer)); clazzId {
	case 0x2410d1a2:
		x.PutClazzID(0x2410d1a2)

		x.PutInt64(m.Creator)
		if m.File == nil {
			return fmt.Errorf("unable to encode dfs_uploadPhotoFileV2#0x2410d1a2: field file is nil")
		}
		if err := m.File.Encode(x, layer); err != nil {
			return fmt.Errorf("unable to encode dfs_uploadPhotoFileV2#0x2410d1a2: field file: %w", err)
		}

		return nil
	default:
		return fmt.Errorf("unable to encode dfs_uploadPhotoFileV2: unsupported layer %d", layer)
	}
}

// Decode <--
func (m *TLDfsUploadPhotoFileV2) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return fmt.Errorf("unable to decode dfs_uploadPhotoFileV2: constructor: %w", err)
		}
	}
	switch m.ClazzID {
	case 0x2410d1a2:
		m.Creator, err = d.Int64()
		if err != nil {
			return fmt.Errorf("unable to decode dfs_uploadPhotoFileV2#0x2410d1a2: field creator: %w", err)
		}

		m.File, err = tg.DecodeInputFileClazz(d)
		if err != nil {
			return fmt.Errorf("unable to decode dfs_uploadPhotoFileV2#0x2410d1a2: field file: %w", err)
		}

		return nil
	default:
		return fmt.Errorf("unable to decode dfs_uploadPhotoFileV2: invalid constructor %x", m.ClazzID)
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
	return iface.DebugStringWithName(ClazzName_dfs_uploadProfilePhotoFileV2, m)
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
			if err := m.File.Encode(x, layer); err != nil {
				return fmt.Errorf("unable to encode dfs_uploadProfilePhotoFileV2#0x872313d8: field file: %w", err)
			}
		}

		if m.Video != nil {
			if err := m.Video.Encode(x, layer); err != nil {
				return fmt.Errorf("unable to encode dfs_uploadProfilePhotoFileV2#0x872313d8: field video: %w", err)
			}
		}

		if m.VideoStartTs != nil {
			x.PutDouble(*m.VideoStartTs)
		}

		if m.VideoEmojiMarkup != nil {
			if err := m.VideoEmojiMarkup.Encode(x, layer); err != nil {
				return fmt.Errorf("unable to encode dfs_uploadProfilePhotoFileV2#0x872313d8: field video_emoji_markup: %w", err)
			}
		}

		return nil
	default:
		return fmt.Errorf("unable to encode dfs_uploadProfilePhotoFileV2: unsupported layer %d", layer)
	}
}

// Decode <--
func (m *TLDfsUploadProfilePhotoFileV2) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return fmt.Errorf("unable to decode dfs_uploadProfilePhotoFileV2: constructor: %w", err)
		}
	}
	switch m.ClazzID {
	case 0x872313d8:
		flags, err := d.Uint32()
		if err != nil {
			return fmt.Errorf("unable to decode dfs_uploadProfilePhotoFileV2: field flags: %w", err)
		}
		_ = flags
		m.Creator, err = d.Int64()
		if err != nil {
			return fmt.Errorf("unable to decode dfs_uploadProfilePhotoFileV2#0x872313d8: field creator: %w", err)
		}
		if (flags & (1 << 0)) != 0 {
			m.File, err = tg.DecodeInputFileClazz(d)
			if err != nil {
				return fmt.Errorf("unable to decode dfs_uploadProfilePhotoFileV2#0x872313d8: field file: %w", err)
			}
		}
		if (flags & (1 << 1)) != 0 {
			m.Video, err = tg.DecodeInputFileClazz(d)
			if err != nil {
				return fmt.Errorf("unable to decode dfs_uploadProfilePhotoFileV2#0x872313d8: field video: %w", err)
			}
		}
		if (flags & (1 << 2)) != 0 {
			m.VideoStartTs = new(float64)
			*m.VideoStartTs, err = d.Double()
			if err != nil {
				return fmt.Errorf("unable to decode dfs_uploadProfilePhotoFileV2#0x872313d8: field video_start_ts: %w", err)
			}
		}

		if (flags & (1 << 4)) != 0 {
			m.VideoEmojiMarkup, err = tg.DecodeVideoSizeClazz(d)
			if err != nil {
				return fmt.Errorf("unable to decode dfs_uploadProfilePhotoFileV2#0x872313d8: field video_emoji_markup: %w", err)
			}
		}

		return nil
	default:
		return fmt.Errorf("unable to decode dfs_uploadProfilePhotoFileV2: invalid constructor %x", m.ClazzID)
	}
}

// TLDfsUploadEncryptedFileV2 <--
type TLDfsUploadEncryptedFileV2 struct {
	ClazzID uint32                     `json:"_id"`
	Creator int64                      `json:"creator"`
	File    tg.InputEncryptedFileClazz `json:"file"`
}

func (m *TLDfsUploadEncryptedFileV2) String() string {
	return iface.DebugStringWithName(ClazzName_dfs_uploadEncryptedFileV2, m)
}

// Encode <--
func (m *TLDfsUploadEncryptedFileV2) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_dfs_uploadEncryptedFileV2, int(layer)); clazzId {
	case 0x79d3c523:
		x.PutClazzID(0x79d3c523)

		x.PutInt64(m.Creator)
		if m.File == nil {
			return fmt.Errorf("unable to encode dfs_uploadEncryptedFileV2#0x79d3c523: field file is nil")
		}
		if err := m.File.Encode(x, layer); err != nil {
			return fmt.Errorf("unable to encode dfs_uploadEncryptedFileV2#0x79d3c523: field file: %w", err)
		}

		return nil
	default:
		return fmt.Errorf("unable to encode dfs_uploadEncryptedFileV2: unsupported layer %d", layer)
	}
}

// Decode <--
func (m *TLDfsUploadEncryptedFileV2) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return fmt.Errorf("unable to decode dfs_uploadEncryptedFileV2: constructor: %w", err)
		}
	}
	switch m.ClazzID {
	case 0x79d3c523:
		m.Creator, err = d.Int64()
		if err != nil {
			return fmt.Errorf("unable to decode dfs_uploadEncryptedFileV2#0x79d3c523: field creator: %w", err)
		}

		m.File, err = tg.DecodeInputEncryptedFileClazz(d)
		if err != nil {
			return fmt.Errorf("unable to decode dfs_uploadEncryptedFileV2#0x79d3c523: field file: %w", err)
		}

		return nil
	default:
		return fmt.Errorf("unable to decode dfs_uploadEncryptedFileV2: invalid constructor %x", m.ClazzID)
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
	return iface.DebugStringWithName(ClazzName_dfs_downloadFile, m)
}

// Encode <--
func (m *TLDfsDownloadFile) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_dfs_downloadFile, int(layer)); clazzId {
	case 0xd6bfee3e:
		x.PutClazzID(0xd6bfee3e)

		if m.Location == nil {
			return fmt.Errorf("unable to encode dfs_downloadFile#0xd6bfee3e: field location is nil")
		}
		if err := m.Location.Encode(x, layer); err != nil {
			return fmt.Errorf("unable to encode dfs_downloadFile#0xd6bfee3e: field location: %w", err)
		}
		x.PutInt64(m.Offset)
		x.PutInt32(m.Limit)

		return nil
	default:
		return fmt.Errorf("unable to encode dfs_downloadFile: unsupported layer %d", layer)
	}
}

// Decode <--
func (m *TLDfsDownloadFile) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return fmt.Errorf("unable to decode dfs_downloadFile: constructor: %w", err)
		}
	}
	switch m.ClazzID {
	case 0xd6bfee3e:

		m.Location, err = tg.DecodeInputFileLocationClazz(d)
		if err != nil {
			return fmt.Errorf("unable to decode dfs_downloadFile#0xd6bfee3e: field location: %w", err)
		}

		m.Offset, err = d.Int64()
		if err != nil {
			return fmt.Errorf("unable to decode dfs_downloadFile#0xd6bfee3e: field offset: %w", err)
		}
		m.Limit, err = d.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode dfs_downloadFile#0xd6bfee3e: field limit: %w", err)
		}

		return nil
	default:
		return fmt.Errorf("unable to decode dfs_downloadFile: invalid constructor %x", m.ClazzID)
	}
}

// TLDfsUploadDocumentFileV2 <--
type TLDfsUploadDocumentFileV2 struct {
	ClazzID uint32             `json:"_id"`
	Creator int64              `json:"creator"`
	Media   tg.InputMediaClazz `json:"media"`
}

func (m *TLDfsUploadDocumentFileV2) String() string {
	return iface.DebugStringWithName(ClazzName_dfs_uploadDocumentFileV2, m)
}

// Encode <--
func (m *TLDfsUploadDocumentFileV2) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_dfs_uploadDocumentFileV2, int(layer)); clazzId {
	case 0x76336db7:
		x.PutClazzID(0x76336db7)

		x.PutInt64(m.Creator)
		if m.Media == nil {
			return fmt.Errorf("unable to encode dfs_uploadDocumentFileV2#0x76336db7: field media is nil")
		}
		if err := m.Media.Encode(x, layer); err != nil {
			return fmt.Errorf("unable to encode dfs_uploadDocumentFileV2#0x76336db7: field media: %w", err)
		}

		return nil
	default:
		return fmt.Errorf("unable to encode dfs_uploadDocumentFileV2: unsupported layer %d", layer)
	}
}

// Decode <--
func (m *TLDfsUploadDocumentFileV2) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return fmt.Errorf("unable to decode dfs_uploadDocumentFileV2: constructor: %w", err)
		}
	}
	switch m.ClazzID {
	case 0x76336db7:
		m.Creator, err = d.Int64()
		if err != nil {
			return fmt.Errorf("unable to decode dfs_uploadDocumentFileV2#0x76336db7: field creator: %w", err)
		}

		m.Media, err = tg.DecodeInputMediaClazz(d)
		if err != nil {
			return fmt.Errorf("unable to decode dfs_uploadDocumentFileV2#0x76336db7: field media: %w", err)
		}

		return nil
	default:
		return fmt.Errorf("unable to decode dfs_uploadDocumentFileV2: invalid constructor %x", m.ClazzID)
	}
}

// TLDfsUploadGifDocumentMedia <--
type TLDfsUploadGifDocumentMedia struct {
	ClazzID uint32             `json:"_id"`
	Creator int64              `json:"creator"`
	Media   tg.InputMediaClazz `json:"media"`
}

func (m *TLDfsUploadGifDocumentMedia) String() string {
	return iface.DebugStringWithName(ClazzName_dfs_uploadGifDocumentMedia, m)
}

// Encode <--
func (m *TLDfsUploadGifDocumentMedia) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_dfs_uploadGifDocumentMedia, int(layer)); clazzId {
	case 0x41c4cd00:
		x.PutClazzID(0x41c4cd00)

		x.PutInt64(m.Creator)
		if m.Media == nil {
			return fmt.Errorf("unable to encode dfs_uploadGifDocumentMedia#0x41c4cd00: field media is nil")
		}
		if err := m.Media.Encode(x, layer); err != nil {
			return fmt.Errorf("unable to encode dfs_uploadGifDocumentMedia#0x41c4cd00: field media: %w", err)
		}

		return nil
	default:
		return fmt.Errorf("unable to encode dfs_uploadGifDocumentMedia: unsupported layer %d", layer)
	}
}

// Decode <--
func (m *TLDfsUploadGifDocumentMedia) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return fmt.Errorf("unable to decode dfs_uploadGifDocumentMedia: constructor: %w", err)
		}
	}
	switch m.ClazzID {
	case 0x41c4cd00:
		m.Creator, err = d.Int64()
		if err != nil {
			return fmt.Errorf("unable to decode dfs_uploadGifDocumentMedia#0x41c4cd00: field creator: %w", err)
		}

		m.Media, err = tg.DecodeInputMediaClazz(d)
		if err != nil {
			return fmt.Errorf("unable to decode dfs_uploadGifDocumentMedia#0x41c4cd00: field media: %w", err)
		}

		return nil
	default:
		return fmt.Errorf("unable to decode dfs_uploadGifDocumentMedia: invalid constructor %x", m.ClazzID)
	}
}

// TLDfsUploadMp4DocumentMedia <--
type TLDfsUploadMp4DocumentMedia struct {
	ClazzID uint32             `json:"_id"`
	Creator int64              `json:"creator"`
	Media   tg.InputMediaClazz `json:"media"`
}

func (m *TLDfsUploadMp4DocumentMedia) String() string {
	return iface.DebugStringWithName(ClazzName_dfs_uploadMp4DocumentMedia, m)
}

// Encode <--
func (m *TLDfsUploadMp4DocumentMedia) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_dfs_uploadMp4DocumentMedia, int(layer)); clazzId {
	case 0xa2a4f818:
		x.PutClazzID(0xa2a4f818)

		x.PutInt64(m.Creator)
		if m.Media == nil {
			return fmt.Errorf("unable to encode dfs_uploadMp4DocumentMedia#0xa2a4f818: field media is nil")
		}
		if err := m.Media.Encode(x, layer); err != nil {
			return fmt.Errorf("unable to encode dfs_uploadMp4DocumentMedia#0xa2a4f818: field media: %w", err)
		}

		return nil
	default:
		return fmt.Errorf("unable to encode dfs_uploadMp4DocumentMedia: unsupported layer %d", layer)
	}
}

// Decode <--
func (m *TLDfsUploadMp4DocumentMedia) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return fmt.Errorf("unable to decode dfs_uploadMp4DocumentMedia: constructor: %w", err)
		}
	}
	switch m.ClazzID {
	case 0xa2a4f818:
		m.Creator, err = d.Int64()
		if err != nil {
			return fmt.Errorf("unable to decode dfs_uploadMp4DocumentMedia#0xa2a4f818: field creator: %w", err)
		}

		m.Media, err = tg.DecodeInputMediaClazz(d)
		if err != nil {
			return fmt.Errorf("unable to decode dfs_uploadMp4DocumentMedia#0xa2a4f818: field media: %w", err)
		}

		return nil
	default:
		return fmt.Errorf("unable to decode dfs_uploadMp4DocumentMedia: invalid constructor %x", m.ClazzID)
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
	return iface.DebugStringWithName(ClazzName_dfs_uploadWallPaperFile, m)
}

// Encode <--
func (m *TLDfsUploadWallPaperFile) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_dfs_uploadWallPaperFile, int(layer)); clazzId {
	case 0xc1a61056:
		x.PutClazzID(0xc1a61056)

		x.PutInt64(m.Creator)
		if m.File == nil {
			return fmt.Errorf("unable to encode dfs_uploadWallPaperFile#0xc1a61056: field file is nil")
		}
		if err := m.File.Encode(x, layer); err != nil {
			return fmt.Errorf("unable to encode dfs_uploadWallPaperFile#0xc1a61056: field file: %w", err)
		}
		x.PutString(m.MimeType)
		if m.Admin == nil {
			return fmt.Errorf("unable to encode dfs_uploadWallPaperFile#0xc1a61056: field admin is nil")
		}
		if err := m.Admin.Encode(x, layer); err != nil {
			return fmt.Errorf("unable to encode dfs_uploadWallPaperFile#0xc1a61056: field admin: %w", err)
		}

		return nil
	default:
		return fmt.Errorf("unable to encode dfs_uploadWallPaperFile: unsupported layer %d", layer)
	}
}

// Decode <--
func (m *TLDfsUploadWallPaperFile) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return fmt.Errorf("unable to decode dfs_uploadWallPaperFile: constructor: %w", err)
		}
	}
	switch m.ClazzID {
	case 0xc1a61056:
		m.Creator, err = d.Int64()
		if err != nil {
			return fmt.Errorf("unable to decode dfs_uploadWallPaperFile#0xc1a61056: field creator: %w", err)
		}

		m.File, err = tg.DecodeInputFileClazz(d)
		if err != nil {
			return fmt.Errorf("unable to decode dfs_uploadWallPaperFile#0xc1a61056: field file: %w", err)
		}

		m.MimeType, err = d.String()
		if err != nil {
			return fmt.Errorf("unable to decode dfs_uploadWallPaperFile#0xc1a61056: field mime_type: %w", err)
		}

		m.Admin, err = tg.DecodeBoolClazz(d)
		if err != nil {
			return fmt.Errorf("unable to decode dfs_uploadWallPaperFile#0xc1a61056: field admin: %w", err)
		}

		return nil
	default:
		return fmt.Errorf("unable to decode dfs_uploadWallPaperFile: invalid constructor %x", m.ClazzID)
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
	return iface.DebugStringWithName(ClazzName_dfs_uploadThemeFile, m)
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
		if m.File == nil {
			return fmt.Errorf("unable to encode dfs_uploadThemeFile#0xdea64f97: field file is nil")
		}
		if err := m.File.Encode(x, layer); err != nil {
			return fmt.Errorf("unable to encode dfs_uploadThemeFile#0xdea64f97: field file: %w", err)
		}
		if m.Thumb != nil {
			if err := m.Thumb.Encode(x, layer); err != nil {
				return fmt.Errorf("unable to encode dfs_uploadThemeFile#0xdea64f97: field thumb: %w", err)
			}
		}

		x.PutString(m.MimeType)
		x.PutString(m.FileName)

		return nil
	default:
		return fmt.Errorf("unable to encode dfs_uploadThemeFile: unsupported layer %d", layer)
	}
}

// Decode <--
func (m *TLDfsUploadThemeFile) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return fmt.Errorf("unable to decode dfs_uploadThemeFile: constructor: %w", err)
		}
	}
	switch m.ClazzID {
	case 0xdea64f97:
		flags, err := d.Uint32()
		if err != nil {
			return fmt.Errorf("unable to decode dfs_uploadThemeFile: field flags: %w", err)
		}
		_ = flags
		m.Creator, err = d.Int64()
		if err != nil {
			return fmt.Errorf("unable to decode dfs_uploadThemeFile#0xdea64f97: field creator: %w", err)
		}

		m.File, err = tg.DecodeInputFileClazz(d)
		if err != nil {
			return fmt.Errorf("unable to decode dfs_uploadThemeFile#0xdea64f97: field file: %w", err)
		}

		if (flags & (1 << 0)) != 0 {
			m.Thumb, err = tg.DecodeInputFileClazz(d)
			if err != nil {
				return fmt.Errorf("unable to decode dfs_uploadThemeFile#0xdea64f97: field thumb: %w", err)
			}
		}
		m.MimeType, err = d.String()
		if err != nil {
			return fmt.Errorf("unable to decode dfs_uploadThemeFile#0xdea64f97: field mime_type: %w", err)
		}
		m.FileName, err = d.String()
		if err != nil {
			return fmt.Errorf("unable to decode dfs_uploadThemeFile#0xdea64f97: field file_name: %w", err)
		}

		return nil
	default:
		return fmt.Errorf("unable to decode dfs_uploadThemeFile: invalid constructor %x", m.ClazzID)
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
	return iface.DebugStringWithName(ClazzName_dfs_uploadRingtoneFile, m)
}

// Encode <--
func (m *TLDfsUploadRingtoneFile) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_dfs_uploadRingtoneFile, int(layer)); clazzId {
	case 0x2b3c5b1:
		x.PutClazzID(0x2b3c5b1)

		x.PutInt64(m.Creator)
		if m.File == nil {
			return fmt.Errorf("unable to encode dfs_uploadRingtoneFile#0x2b3c5b1: field file is nil")
		}
		if err := m.File.Encode(x, layer); err != nil {
			return fmt.Errorf("unable to encode dfs_uploadRingtoneFile#0x2b3c5b1: field file: %w", err)
		}
		x.PutString(m.MimeType)
		x.PutString(m.FileName)

		return nil
	default:
		return fmt.Errorf("unable to encode dfs_uploadRingtoneFile: unsupported layer %d", layer)
	}
}

// Decode <--
func (m *TLDfsUploadRingtoneFile) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return fmt.Errorf("unable to decode dfs_uploadRingtoneFile: constructor: %w", err)
		}
	}
	switch m.ClazzID {
	case 0x2b3c5b1:
		m.Creator, err = d.Int64()
		if err != nil {
			return fmt.Errorf("unable to decode dfs_uploadRingtoneFile#0x2b3c5b1: field creator: %w", err)
		}

		m.File, err = tg.DecodeInputFileClazz(d)
		if err != nil {
			return fmt.Errorf("unable to decode dfs_uploadRingtoneFile#0x2b3c5b1: field file: %w", err)
		}

		m.MimeType, err = d.String()
		if err != nil {
			return fmt.Errorf("unable to decode dfs_uploadRingtoneFile#0x2b3c5b1: field mime_type: %w", err)
		}
		m.FileName, err = d.String()
		if err != nil {
			return fmt.Errorf("unable to decode dfs_uploadRingtoneFile#0x2b3c5b1: field file_name: %w", err)
		}

		return nil
	default:
		return fmt.Errorf("unable to decode dfs_uploadRingtoneFile: invalid constructor %x", m.ClazzID)
	}
}

// TLDfsUploadedProfilePhoto <--
type TLDfsUploadedProfilePhoto struct {
	ClazzID uint32 `json:"_id"`
	Creator int64  `json:"creator"`
	PhotoId int64  `json:"photo_id"`
}

func (m *TLDfsUploadedProfilePhoto) String() string {
	return iface.DebugStringWithName(ClazzName_dfs_uploadedProfilePhoto, m)
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
		return fmt.Errorf("unable to encode dfs_uploadedProfilePhoto: unsupported layer %d", layer)
	}
}

// Decode <--
func (m *TLDfsUploadedProfilePhoto) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return fmt.Errorf("unable to decode dfs_uploadedProfilePhoto: constructor: %w", err)
		}
	}
	switch m.ClazzID {
	case 0xa3aa2874:
		m.Creator, err = d.Int64()
		if err != nil {
			return fmt.Errorf("unable to decode dfs_uploadedProfilePhoto#0xa3aa2874: field creator: %w", err)
		}
		m.PhotoId, err = d.Int64()
		if err != nil {
			return fmt.Errorf("unable to decode dfs_uploadedProfilePhoto#0xa3aa2874: field photo_id: %w", err)
		}

		return nil
	default:
		return fmt.Errorf("unable to decode dfs_uploadedProfilePhoto: invalid constructor %x", m.ClazzID)
	}
}

// Vector api result type
// ----------------------------------------------------------------------------
// VectorResList <--

// VectorFileHash <--
type VectorFileHash struct {
	Datas []tg.FileHashClazz `json:"_datas"`
}

func (m *VectorFileHash) String() string {
	data, _ := json.Marshal(m)
	return string(data)
}

// Encode <--
func (m *VectorFileHash) Encode(x *bin.Encoder, layer int32) error {
	_ = iface.EncodeObjectList(x, m.Datas, layer)

	return nil
}

// Decode <--
func (m *VectorFileHash) Decode(d *bin.Decoder) (err error) {
	m.Datas, err = iface.DecodeObjectList[tg.FileHashClazz](d)

	return err
}

// ----------------------------------------------------------------------------
// rpc

type RPCDfs interface {
	DfsCommitUpload(ctx context.Context, in *TLDfsCommitUpload) (*FileFinalizedObject, error)
	DfsPutFile(ctx context.Context, in *TLDfsPutFile) (*FileFinalizedObject, error)
	DfsGetFileByReadLease(ctx context.Context, in *TLDfsGetFileByReadLease) (*tg.UploadFile, error)
	DfsGetFileHashesByReadLease(ctx context.Context, in *TLDfsGetFileHashesByReadLease) (*VectorFileHash, error)
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
