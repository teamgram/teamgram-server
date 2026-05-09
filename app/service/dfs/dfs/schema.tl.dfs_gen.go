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

// FileFinalizedObjectClazz <--
//   - TL_FileFinalizedObject
type FileFinalizedObjectClazz = *TLFileFinalizedObject

func DecodeFileFinalizedObjectClazz(d *bin.Decoder) (FileFinalizedObjectClazz, error) {
	// id, err := d.PeekClazzID()
	id, err := d.ClazzID()
	if err != nil {
		return nil, fmt.Errorf("unable to decode FileFinalizedObject: constructor: %w", err)
	}

	switch id {
	case 0xe83380f0:
		x := &TLFileFinalizedObject{ClazzID: id, ClazzName2: ClazzName_fileFinalizedObject}
		if err := x.Decode(d); err != nil {
			return nil, err
		}
		return x, nil
	default:
		return nil, fmt.Errorf("unable to decode FileFinalizedObject: invalid constructor %x", id)
	}

}

// TLFileFinalizedObject <--
type TLFileFinalizedObject struct {
	ClazzID         uint32 `json:"_id"`
	ClazzName2      string `json:"_name"`
	ObjectId        string `json:"object_id"`
	UploadSessionId string `json:"upload_session_id"`
	Bucket          string `json:"bucket"`
	Key             string `json:"key"`
	Size2           int64  `json:"size2"`
	MimeType        string `json:"mime_type"`
	Sha256          []byte `json:"sha256"`
	ReadLease       []byte `json:"read_lease"`
	DcId            int32  `json:"dc_id"`
}

func MakeTLFileFinalizedObject(m *TLFileFinalizedObject) *TLFileFinalizedObject {
	if m == nil {
		return nil
	}
	m.ClazzName2 = ClazzName_fileFinalizedObject

	return m
}

func (m *TLFileFinalizedObject) String() string {
	return iface.DebugStringWithName("fileFinalizedObject", m)
}

func (m *TLFileFinalizedObject) MarshalJSON() ([]byte, error) {
	if m == nil {
		return []byte("null"), nil
	}
	return iface.MarshalWithName("fileFinalizedObject", m)
}

// FileFinalizedObjectClazzName <--
func (m *TLFileFinalizedObject) FileFinalizedObjectClazzName() string {
	return ClazzName_fileFinalizedObject
}

// ClazzName <--
func (m *TLFileFinalizedObject) ClazzName() string {
	return m.ClazzName2
}

// ToFileFinalizedObject <--
func (m *TLFileFinalizedObject) ToFileFinalizedObject() *FileFinalizedObject {
	if m == nil {
		return nil
	}

	return m

}

func (m *TLFileFinalizedObject) CalcSize(layer int32) int {
	switch clazzId := iface.GetClazzIDByName(ClazzName_fileFinalizedObject, int(layer)); clazzId {
	case 0xe83380f0:
		size := 4
		size += iface.CalcStringSize(m.ObjectId)
		size += iface.CalcStringSize(m.UploadSessionId)
		size += iface.CalcStringSize(m.Bucket)
		size += iface.CalcStringSize(m.Key)
		size += 8
		size += iface.CalcStringSize(m.MimeType)
		size += iface.CalcBytesSize(m.Sha256)
		size += iface.CalcBytesSize(m.ReadLease)
		size += 4

		return size
	default:
		return 0
	}
}

func (m *TLFileFinalizedObject) Validate(layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_fileFinalizedObject, int(layer)); clazzId {
	case 0xe83380f0:
		if err := iface.ValidateRequiredString("object_id", m.ObjectId); err != nil {
			return err
		}

		if err := iface.ValidateRequiredString("upload_session_id", m.UploadSessionId); err != nil {
			return err
		}

		if err := iface.ValidateRequiredString("bucket", m.Bucket); err != nil {
			return err
		}

		if err := iface.ValidateRequiredString("key", m.Key); err != nil {
			return err
		}

		if err := iface.ValidateRequiredString("mime_type", m.MimeType); err != nil {
			return err
		}

		if err := iface.ValidateRequiredBytes("sha256", m.Sha256); err != nil {
			return err
		}

		if err := iface.ValidateRequiredBytes("read_lease", m.ReadLease); err != nil {
			return err
		}

		return nil
	default:
		return fmt.Errorf("unable to encode fileFinalizedObject: unsupported layer %d", layer)
	}
}

// Encode <--
func (m *TLFileFinalizedObject) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_fileFinalizedObject, int(layer)); clazzId {
	case 0xe83380f0:
		x.PutClazzID(0xe83380f0)

		x.PutString(m.ObjectId)
		x.PutString(m.UploadSessionId)
		x.PutString(m.Bucket)
		x.PutString(m.Key)
		x.PutInt64(m.Size2)
		x.PutString(m.MimeType)
		x.PutBytes(m.Sha256)
		x.PutBytes(m.ReadLease)
		x.PutInt32(m.DcId)

		return nil
	default:
		return fmt.Errorf("unable to encode fileFinalizedObject: unsupported layer %d", layer)
	}
}

// Decode <--
func (m *TLFileFinalizedObject) Decode(d *bin.Decoder) (err error) {
	switch m.ClazzID {
	case 0xe83380f0:
		m.ObjectId, err = d.String()
		if err != nil {
			return fmt.Errorf("unable to decode fileFinalizedObject#0xe83380f0: field object_id: %w", err)
		}
		m.UploadSessionId, err = d.String()
		if err != nil {
			return fmt.Errorf("unable to decode fileFinalizedObject#0xe83380f0: field upload_session_id: %w", err)
		}
		m.Bucket, err = d.String()
		if err != nil {
			return fmt.Errorf("unable to decode fileFinalizedObject#0xe83380f0: field bucket: %w", err)
		}
		m.Key, err = d.String()
		if err != nil {
			return fmt.Errorf("unable to decode fileFinalizedObject#0xe83380f0: field key: %w", err)
		}
		m.Size2, err = d.Int64()
		if err != nil {
			return fmt.Errorf("unable to decode fileFinalizedObject#0xe83380f0: field size2: %w", err)
		}
		m.MimeType, err = d.String()
		if err != nil {
			return fmt.Errorf("unable to decode fileFinalizedObject#0xe83380f0: field mime_type: %w", err)
		}
		m.Sha256, err = d.Bytes()
		if err != nil {
			return fmt.Errorf("unable to decode fileFinalizedObject#0xe83380f0: field sha256: %w", err)
		}
		m.ReadLease, err = d.Bytes()
		if err != nil {
			return fmt.Errorf("unable to decode fileFinalizedObject#0xe83380f0: field read_lease: %w", err)
		}
		m.DcId, err = d.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode fileFinalizedObject#0xe83380f0: field dc_id: %w", err)
		}

		return nil
	default:
		return fmt.Errorf("unable to decode fileFinalizedObject: invalid constructor %x", m.ClazzID)
	}
}

// FileFinalizedObject <--
type FileFinalizedObject = TLFileFinalizedObject

// FileHashChunkClazz <--
//   - TL_FileHashChunk
type FileHashChunkClazz = *TLFileHashChunk

func DecodeFileHashChunkClazz(d *bin.Decoder) (FileHashChunkClazz, error) {
	// id, err := d.PeekClazzID()
	id, err := d.ClazzID()
	if err != nil {
		return nil, fmt.Errorf("unable to decode FileHashChunk: constructor: %w", err)
	}

	switch id {
	case 0x146aad14:
		x := &TLFileHashChunk{ClazzID: id, ClazzName2: ClazzName_fileHashChunk}
		if err := x.Decode(d); err != nil {
			return nil, err
		}
		return x, nil
	default:
		return nil, fmt.Errorf("unable to decode FileHashChunk: invalid constructor %x", id)
	}

}

// TLFileHashChunk <--
type TLFileHashChunk struct {
	ClazzID    uint32 `json:"_id"`
	ClazzName2 string `json:"_name"`
	Offset     int64  `json:"offset"`
	Limit      int32  `json:"limit"`
	Hash       []byte `json:"hash"`
}

func MakeTLFileHashChunk(m *TLFileHashChunk) *TLFileHashChunk {
	if m == nil {
		return nil
	}
	m.ClazzName2 = ClazzName_fileHashChunk

	return m
}

func (m *TLFileHashChunk) String() string {
	return iface.DebugStringWithName("fileHashChunk", m)
}

func (m *TLFileHashChunk) MarshalJSON() ([]byte, error) {
	if m == nil {
		return []byte("null"), nil
	}
	return iface.MarshalWithName("fileHashChunk", m)
}

// FileHashChunkClazzName <--
func (m *TLFileHashChunk) FileHashChunkClazzName() string {
	return ClazzName_fileHashChunk
}

// ClazzName <--
func (m *TLFileHashChunk) ClazzName() string {
	return m.ClazzName2
}

// ToFileHashChunk <--
func (m *TLFileHashChunk) ToFileHashChunk() *FileHashChunk {
	if m == nil {
		return nil
	}

	return m

}

func (m *TLFileHashChunk) CalcSize(layer int32) int {
	switch clazzId := iface.GetClazzIDByName(ClazzName_fileHashChunk, int(layer)); clazzId {
	case 0x146aad14:
		size := 4
		size += 8
		size += 4
		size += iface.CalcBytesSize(m.Hash)

		return size
	default:
		return 0
	}
}

func (m *TLFileHashChunk) Validate(layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_fileHashChunk, int(layer)); clazzId {
	case 0x146aad14:
		if err := iface.ValidateRequiredBytes("hash", m.Hash); err != nil {
			return err
		}

		return nil
	default:
		return fmt.Errorf("unable to encode fileHashChunk: unsupported layer %d", layer)
	}
}

// Encode <--
func (m *TLFileHashChunk) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_fileHashChunk, int(layer)); clazzId {
	case 0x146aad14:
		x.PutClazzID(0x146aad14)

		x.PutInt64(m.Offset)
		x.PutInt32(m.Limit)
		x.PutBytes(m.Hash)

		return nil
	default:
		return fmt.Errorf("unable to encode fileHashChunk: unsupported layer %d", layer)
	}
}

// Decode <--
func (m *TLFileHashChunk) Decode(d *bin.Decoder) (err error) {
	switch m.ClazzID {
	case 0x146aad14:
		m.Offset, err = d.Int64()
		if err != nil {
			return fmt.Errorf("unable to decode fileHashChunk#0x146aad14: field offset: %w", err)
		}
		m.Limit, err = d.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode fileHashChunk#0x146aad14: field limit: %w", err)
		}
		m.Hash, err = d.Bytes()
		if err != nil {
			return fmt.Errorf("unable to decode fileHashChunk#0x146aad14: field hash: %w", err)
		}

		return nil
	default:
		return fmt.Errorf("unable to decode fileHashChunk: invalid constructor %x", m.ClazzID)
	}
}

// FileHashChunk <--
type FileHashChunk = TLFileHashChunk
