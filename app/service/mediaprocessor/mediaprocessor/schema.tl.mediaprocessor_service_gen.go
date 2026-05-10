/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2026-present, The Teamgram Authors (https://teamgram.net).
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package mediaprocessor

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

// TLMediaProcessorProcessPhoto <--
type TLMediaProcessorProcessPhoto struct {
	ClazzID   uint32       `json:"_id"`
	OwnerId   int64        `json:"owner_id"`
	ObjectId  string       `json:"object_id"`
	ReadLease []byte       `json:"read_lease"`
	FileName  string       `json:"file_name"`
	Profile   tg.BoolClazz `json:"profile"`
}

func (m *TLMediaProcessorProcessPhoto) String() string {
	return iface.DebugStringWithName(ClazzName_mediaProcessor_processPhoto, m)
}

// Encode <--
func (m *TLMediaProcessorProcessPhoto) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_mediaProcessor_processPhoto, int(layer)); clazzId {
	case 0x23289b04:
		x.PutClazzID(0x23289b04)

		x.PutInt64(m.OwnerId)
		x.PutString(m.ObjectId)
		x.PutBytes(m.ReadLease)
		x.PutString(m.FileName)
		if m.Profile == nil {
			return fmt.Errorf("unable to encode mediaProcessor_processPhoto#0x23289b04: field profile is nil")
		}
		if err := m.Profile.Encode(x, layer); err != nil {
			return fmt.Errorf("unable to encode mediaProcessor_processPhoto#0x23289b04: field profile: %w", err)
		}

		return nil
	default:
		return fmt.Errorf("unable to encode mediaProcessor_processPhoto: unsupported layer %d", layer)
	}
}

// Decode <--
func (m *TLMediaProcessorProcessPhoto) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return fmt.Errorf("unable to decode mediaProcessor_processPhoto: constructor: %w", err)
		}
	}
	switch m.ClazzID {
	case 0x23289b04:
		m.OwnerId, err = d.Int64()
		if err != nil {
			return fmt.Errorf("unable to decode mediaProcessor_processPhoto#0x23289b04: field owner_id: %w", err)
		}
		m.ObjectId, err = d.String()
		if err != nil {
			return fmt.Errorf("unable to decode mediaProcessor_processPhoto#0x23289b04: field object_id: %w", err)
		}
		m.ReadLease, err = d.Bytes()
		if err != nil {
			return fmt.Errorf("unable to decode mediaProcessor_processPhoto#0x23289b04: field read_lease: %w", err)
		}
		m.FileName, err = d.String()
		if err != nil {
			return fmt.Errorf("unable to decode mediaProcessor_processPhoto#0x23289b04: field file_name: %w", err)
		}

		m.Profile, err = tg.DecodeBoolClazz(d)
		if err != nil {
			return fmt.Errorf("unable to decode mediaProcessor_processPhoto#0x23289b04: field profile: %w", err)
		}

		return nil
	default:
		return fmt.Errorf("unable to decode mediaProcessor_processPhoto: invalid constructor %x", m.ClazzID)
	}
}

// TLMediaProcessorProcessGif <--
type TLMediaProcessorProcessGif struct {
	ClazzID        uint32 `json:"_id"`
	OwnerId        int64  `json:"owner_id"`
	ObjectId       string `json:"object_id"`
	ReadLease      []byte `json:"read_lease"`
	FileName       string `json:"file_name"`
	ThumbObjectId  string `json:"thumb_object_id"`
	ThumbReadLease []byte `json:"thumb_read_lease"`
}

func (m *TLMediaProcessorProcessGif) String() string {
	return iface.DebugStringWithName(ClazzName_mediaProcessor_processGif, m)
}

// Encode <--
func (m *TLMediaProcessorProcessGif) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_mediaProcessor_processGif, int(layer)); clazzId {
	case 0xcaa60c8c:
		x.PutClazzID(0xcaa60c8c)

		x.PutInt64(m.OwnerId)
		x.PutString(m.ObjectId)
		x.PutBytes(m.ReadLease)
		x.PutString(m.FileName)
		x.PutString(m.ThumbObjectId)
		x.PutBytes(m.ThumbReadLease)

		return nil
	default:
		return fmt.Errorf("unable to encode mediaProcessor_processGif: unsupported layer %d", layer)
	}
}

// Decode <--
func (m *TLMediaProcessorProcessGif) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return fmt.Errorf("unable to decode mediaProcessor_processGif: constructor: %w", err)
		}
	}
	switch m.ClazzID {
	case 0xcaa60c8c:
		m.OwnerId, err = d.Int64()
		if err != nil {
			return fmt.Errorf("unable to decode mediaProcessor_processGif#0xcaa60c8c: field owner_id: %w", err)
		}
		m.ObjectId, err = d.String()
		if err != nil {
			return fmt.Errorf("unable to decode mediaProcessor_processGif#0xcaa60c8c: field object_id: %w", err)
		}
		m.ReadLease, err = d.Bytes()
		if err != nil {
			return fmt.Errorf("unable to decode mediaProcessor_processGif#0xcaa60c8c: field read_lease: %w", err)
		}
		m.FileName, err = d.String()
		if err != nil {
			return fmt.Errorf("unable to decode mediaProcessor_processGif#0xcaa60c8c: field file_name: %w", err)
		}
		m.ThumbObjectId, err = d.String()
		if err != nil {
			return fmt.Errorf("unable to decode mediaProcessor_processGif#0xcaa60c8c: field thumb_object_id: %w", err)
		}
		m.ThumbReadLease, err = d.Bytes()
		if err != nil {
			return fmt.Errorf("unable to decode mediaProcessor_processGif#0xcaa60c8c: field thumb_read_lease: %w", err)
		}

		return nil
	default:
		return fmt.Errorf("unable to decode mediaProcessor_processGif: invalid constructor %x", m.ClazzID)
	}
}

// TLMediaProcessorProcessMp4 <--
type TLMediaProcessorProcessMp4 struct {
	ClazzID    uint32 `json:"_id"`
	OwnerId    int64  `json:"owner_id"`
	ObjectId   string `json:"object_id"`
	ReadLease  []byte `json:"read_lease"`
	FileName   string `json:"file_name"`
	Attributes []byte `json:"attributes"`
}

func (m *TLMediaProcessorProcessMp4) String() string {
	return iface.DebugStringWithName(ClazzName_mediaProcessor_processMp4, m)
}

// Encode <--
func (m *TLMediaProcessorProcessMp4) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_mediaProcessor_processMp4, int(layer)); clazzId {
	case 0xac180ca1:
		x.PutClazzID(0xac180ca1)

		x.PutInt64(m.OwnerId)
		x.PutString(m.ObjectId)
		x.PutBytes(m.ReadLease)
		x.PutString(m.FileName)
		x.PutBytes(m.Attributes)

		return nil
	default:
		return fmt.Errorf("unable to encode mediaProcessor_processMp4: unsupported layer %d", layer)
	}
}

// Decode <--
func (m *TLMediaProcessorProcessMp4) Decode(d *bin.Decoder) (err error) {
	if m.ClazzID == 0 {
		m.ClazzID, err = d.ClazzID()
		if err != nil {
			return fmt.Errorf("unable to decode mediaProcessor_processMp4: constructor: %w", err)
		}
	}
	switch m.ClazzID {
	case 0xac180ca1:
		m.OwnerId, err = d.Int64()
		if err != nil {
			return fmt.Errorf("unable to decode mediaProcessor_processMp4#0xac180ca1: field owner_id: %w", err)
		}
		m.ObjectId, err = d.String()
		if err != nil {
			return fmt.Errorf("unable to decode mediaProcessor_processMp4#0xac180ca1: field object_id: %w", err)
		}
		m.ReadLease, err = d.Bytes()
		if err != nil {
			return fmt.Errorf("unable to decode mediaProcessor_processMp4#0xac180ca1: field read_lease: %w", err)
		}
		m.FileName, err = d.String()
		if err != nil {
			return fmt.Errorf("unable to decode mediaProcessor_processMp4#0xac180ca1: field file_name: %w", err)
		}
		m.Attributes, err = d.Bytes()
		if err != nil {
			return fmt.Errorf("unable to decode mediaProcessor_processMp4#0xac180ca1: field attributes: %w", err)
		}

		return nil
	default:
		return fmt.Errorf("unable to decode mediaProcessor_processMp4: invalid constructor %x", m.ClazzID)
	}
}

// Vector api result type
// ----------------------------------------------------------------------------
// VectorResList <--

// ----------------------------------------------------------------------------
// rpc

type RPCMediaProcessor interface {
	MediaProcessorProcessPhoto(ctx context.Context, in *TLMediaProcessorProcessPhoto) (*ProcessedPhoto, error)
	MediaProcessorProcessGif(ctx context.Context, in *TLMediaProcessorProcessGif) (*ProcessedDocument, error)
	MediaProcessorProcessMp4(ctx context.Context, in *TLMediaProcessorProcessMp4) (*ProcessedDocument, error)
}
