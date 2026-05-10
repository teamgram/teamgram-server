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

// ProcessedDocumentClazz <--
//   - TL_ProcessedDocument
type ProcessedDocumentClazz = *TLProcessedDocument

func DecodeProcessedDocumentClazz(d *bin.Decoder) (ProcessedDocumentClazz, error) {
	// id, err := d.PeekClazzID()
	id, err := d.ClazzID()
	if err != nil {
		return nil, fmt.Errorf("unable to decode ProcessedDocument: constructor: %w", err)
	}

	switch id {
	case 0xfb5d44f8:
		x := &TLProcessedDocument{ClazzID: id, ClazzName2: ClazzName_processedDocument}
		if err := x.Decode(d); err != nil {
			return nil, err
		}
		return x, nil
	default:
		return nil, fmt.Errorf("unable to decode ProcessedDocument: invalid constructor %x", id)
	}

}

// TLProcessedDocument <--
type TLProcessedDocument struct {
	ClazzID    uint32                     `json:"_id"`
	ClazzName2 string                     `json:"_name"`
	ObjectId   string                     `json:"object_id"`
	MimeType   string                     `json:"mime_type"`
	Size2      int64                      `json:"size2"`
	Attributes []byte                     `json:"attributes"`
	Thumbs     []ProcessorDerivativeClazz `json:"thumbs"`
}

func MakeTLProcessedDocument(m *TLProcessedDocument) *TLProcessedDocument {
	if m == nil {
		return nil
	}
	m.ClazzName2 = ClazzName_processedDocument

	return m
}

func (m *TLProcessedDocument) String() string {
	return iface.DebugStringWithName("processedDocument", m)
}

func (m *TLProcessedDocument) MarshalJSON() ([]byte, error) {
	if m == nil {
		return []byte("null"), nil
	}
	return iface.MarshalWithName("processedDocument", m)
}

// ProcessedDocumentClazzName <--
func (m *TLProcessedDocument) ProcessedDocumentClazzName() string {
	return ClazzName_processedDocument
}

// ClazzName <--
func (m *TLProcessedDocument) ClazzName() string {
	return m.ClazzName2
}

// ToProcessedDocument <--
func (m *TLProcessedDocument) ToProcessedDocument() *ProcessedDocument {
	if m == nil {
		return nil
	}

	return m

}

func (m *TLProcessedDocument) CalcSize(layer int32) int {
	switch clazzId := iface.GetClazzIDByName(ClazzName_processedDocument, int(layer)); clazzId {
	case 0xfb5d44f8:
		size := 4
		size += iface.CalcStringSize(m.ObjectId)
		size += iface.CalcStringSize(m.MimeType)
		size += 8
		size += iface.CalcBytesSize(m.Attributes)
		size += iface.CalcObjectListSize(m.Thumbs, layer)

		return size
	default:
		return 0
	}
}

func (m *TLProcessedDocument) Validate(layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_processedDocument, int(layer)); clazzId {
	case 0xfb5d44f8:
		if err := iface.ValidateRequiredString("object_id", m.ObjectId); err != nil {
			return err
		}

		if err := iface.ValidateRequiredString("mime_type", m.MimeType); err != nil {
			return err
		}

		if err := iface.ValidateRequiredBytes("attributes", m.Attributes); err != nil {
			return err
		}

		if err := iface.ValidateRequiredSlice("thumbs", m.Thumbs); err != nil {
			return err
		}

		return nil
	default:
		return fmt.Errorf("unable to encode processedDocument: unsupported layer %d", layer)
	}
}

// Encode <--
func (m *TLProcessedDocument) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_processedDocument, int(layer)); clazzId {
	case 0xfb5d44f8:
		x.PutClazzID(0xfb5d44f8)

		x.PutString(m.ObjectId)
		x.PutString(m.MimeType)
		x.PutInt64(m.Size2)
		x.PutBytes(m.Attributes)

		if err := iface.EncodeObjectList(x, m.Thumbs, layer); err != nil {
			return fmt.Errorf("unable to encode processedDocument#0xfb5d44f8: field thumbs: %w", err)
		}

		return nil
	default:
		return fmt.Errorf("unable to encode processedDocument: unsupported layer %d", layer)
	}
}

// Decode <--
func (m *TLProcessedDocument) Decode(d *bin.Decoder) (err error) {
	switch m.ClazzID {
	case 0xfb5d44f8:
		m.ObjectId, err = d.String()
		if err != nil {
			return fmt.Errorf("unable to decode processedDocument#0xfb5d44f8: field object_id: %w", err)
		}
		m.MimeType, err = d.String()
		if err != nil {
			return fmt.Errorf("unable to decode processedDocument#0xfb5d44f8: field mime_type: %w", err)
		}
		m.Size2, err = d.Int64()
		if err != nil {
			return fmt.Errorf("unable to decode processedDocument#0xfb5d44f8: field size2: %w", err)
		}
		m.Attributes, err = d.Bytes()
		if err != nil {
			return fmt.Errorf("unable to decode processedDocument#0xfb5d44f8: field attributes: %w", err)
		}
		l4, err3 := d.VectorHeader()
		if err3 != nil {
			return fmt.Errorf("unable to decode processedDocument#0xfb5d44f8: field thumbs: %w", err3)
		}
		if l4 > bin.MaxVectorLen {
			return fmt.Errorf("unable to decode processedDocument#0xfb5d44f8: field thumbs: %w", &bin.InvalidLengthError{Type: "vector", Length: int(l4)})
		}
		prealloc4 := int(l4)
		if prealloc4 > bin.PreallocateLimit {
			prealloc4 = bin.PreallocateLimit
		}
		v4 := make([]ProcessorDerivativeClazz, 0, prealloc4)
		for i := int32(0); i < l4; i++ {
			vv4, err3 := DecodeProcessorDerivativeClazz(d)
			if err3 != nil {
				return fmt.Errorf("unable to decode processedDocument#0xfb5d44f8: field thumbs: %w", err3)
			}
			v4 = append(v4, vv4)
		}
		m.Thumbs = v4

		return nil
	default:
		return fmt.Errorf("unable to decode processedDocument: invalid constructor %x", m.ClazzID)
	}
}

// ProcessedDocument <--
type ProcessedDocument = TLProcessedDocument

// ProcessedPhotoClazz <--
//   - TL_ProcessedPhoto
type ProcessedPhotoClazz = *TLProcessedPhoto

func DecodeProcessedPhotoClazz(d *bin.Decoder) (ProcessedPhotoClazz, error) {
	// id, err := d.PeekClazzID()
	id, err := d.ClazzID()
	if err != nil {
		return nil, fmt.Errorf("unable to decode ProcessedPhoto: constructor: %w", err)
	}

	switch id {
	case 0x606d445:
		x := &TLProcessedPhoto{ClazzID: id, ClazzName2: ClazzName_processedPhoto}
		if err := x.Decode(d); err != nil {
			return nil, err
		}
		return x, nil
	default:
		return nil, fmt.Errorf("unable to decode ProcessedPhoto: invalid constructor %x", id)
	}

}

// TLProcessedPhoto <--
type TLProcessedPhoto struct {
	ClazzID          uint32                     `json:"_id"`
	ClazzName2       string                     `json:"_name"`
	OriginalObjectId string                     `json:"original_object_id"`
	Sizes            []ProcessorDerivativeClazz `json:"sizes"`
}

func MakeTLProcessedPhoto(m *TLProcessedPhoto) *TLProcessedPhoto {
	if m == nil {
		return nil
	}
	m.ClazzName2 = ClazzName_processedPhoto

	return m
}

func (m *TLProcessedPhoto) String() string {
	return iface.DebugStringWithName("processedPhoto", m)
}

func (m *TLProcessedPhoto) MarshalJSON() ([]byte, error) {
	if m == nil {
		return []byte("null"), nil
	}
	return iface.MarshalWithName("processedPhoto", m)
}

// ProcessedPhotoClazzName <--
func (m *TLProcessedPhoto) ProcessedPhotoClazzName() string {
	return ClazzName_processedPhoto
}

// ClazzName <--
func (m *TLProcessedPhoto) ClazzName() string {
	return m.ClazzName2
}

// ToProcessedPhoto <--
func (m *TLProcessedPhoto) ToProcessedPhoto() *ProcessedPhoto {
	if m == nil {
		return nil
	}

	return m

}

func (m *TLProcessedPhoto) CalcSize(layer int32) int {
	switch clazzId := iface.GetClazzIDByName(ClazzName_processedPhoto, int(layer)); clazzId {
	case 0x606d445:
		size := 4
		size += iface.CalcStringSize(m.OriginalObjectId)
		size += iface.CalcObjectListSize(m.Sizes, layer)

		return size
	default:
		return 0
	}
}

func (m *TLProcessedPhoto) Validate(layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_processedPhoto, int(layer)); clazzId {
	case 0x606d445:
		if err := iface.ValidateRequiredString("original_object_id", m.OriginalObjectId); err != nil {
			return err
		}

		if err := iface.ValidateRequiredSlice("sizes", m.Sizes); err != nil {
			return err
		}

		return nil
	default:
		return fmt.Errorf("unable to encode processedPhoto: unsupported layer %d", layer)
	}
}

// Encode <--
func (m *TLProcessedPhoto) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_processedPhoto, int(layer)); clazzId {
	case 0x606d445:
		x.PutClazzID(0x606d445)

		x.PutString(m.OriginalObjectId)

		if err := iface.EncodeObjectList(x, m.Sizes, layer); err != nil {
			return fmt.Errorf("unable to encode processedPhoto#0x606d445: field sizes: %w", err)
		}

		return nil
	default:
		return fmt.Errorf("unable to encode processedPhoto: unsupported layer %d", layer)
	}
}

// Decode <--
func (m *TLProcessedPhoto) Decode(d *bin.Decoder) (err error) {
	switch m.ClazzID {
	case 0x606d445:
		m.OriginalObjectId, err = d.String()
		if err != nil {
			return fmt.Errorf("unable to decode processedPhoto#0x606d445: field original_object_id: %w", err)
		}
		l1, err3 := d.VectorHeader()
		if err3 != nil {
			return fmt.Errorf("unable to decode processedPhoto#0x606d445: field sizes: %w", err3)
		}
		if l1 > bin.MaxVectorLen {
			return fmt.Errorf("unable to decode processedPhoto#0x606d445: field sizes: %w", &bin.InvalidLengthError{Type: "vector", Length: int(l1)})
		}
		prealloc1 := int(l1)
		if prealloc1 > bin.PreallocateLimit {
			prealloc1 = bin.PreallocateLimit
		}
		v1 := make([]ProcessorDerivativeClazz, 0, prealloc1)
		for i := int32(0); i < l1; i++ {
			vv1, err3 := DecodeProcessorDerivativeClazz(d)
			if err3 != nil {
				return fmt.Errorf("unable to decode processedPhoto#0x606d445: field sizes: %w", err3)
			}
			v1 = append(v1, vv1)
		}
		m.Sizes = v1

		return nil
	default:
		return fmt.Errorf("unable to decode processedPhoto: invalid constructor %x", m.ClazzID)
	}
}

// ProcessedPhoto <--
type ProcessedPhoto = TLProcessedPhoto

// ProcessorDerivativeClazz <--
//   - TL_ProcessorDerivative
type ProcessorDerivativeClazz = *TLProcessorDerivative

func DecodeProcessorDerivativeClazz(d *bin.Decoder) (ProcessorDerivativeClazz, error) {
	// id, err := d.PeekClazzID()
	id, err := d.ClazzID()
	if err != nil {
		return nil, fmt.Errorf("unable to decode ProcessorDerivative: constructor: %w", err)
	}

	switch id {
	case 0x2af751de:
		x := &TLProcessorDerivative{ClazzID: id, ClazzName2: ClazzName_processorDerivative}
		if err := x.Decode(d); err != nil {
			return nil, err
		}
		return x, nil
	default:
		return nil, fmt.Errorf("unable to decode ProcessorDerivative: invalid constructor %x", id)
	}

}

// TLProcessorDerivative <--
type TLProcessorDerivative struct {
	ClazzID    uint32 `json:"_id"`
	ClazzName2 string `json:"_name"`
	Kind       string `json:"kind"`
	ObjectId   string `json:"object_id"`
	FileName   string `json:"file_name"`
	MimeType   string `json:"mime_type"`
	Size2      int64  `json:"size2"`
	Width      int32  `json:"width"`
	Height     int32  `json:"height"`
	Bytes      []byte `json:"bytes"`
}

func MakeTLProcessorDerivative(m *TLProcessorDerivative) *TLProcessorDerivative {
	if m == nil {
		return nil
	}
	m.ClazzName2 = ClazzName_processorDerivative

	return m
}

func (m *TLProcessorDerivative) String() string {
	return iface.DebugStringWithName("processorDerivative", m)
}

func (m *TLProcessorDerivative) MarshalJSON() ([]byte, error) {
	if m == nil {
		return []byte("null"), nil
	}
	return iface.MarshalWithName("processorDerivative", m)
}

// ProcessorDerivativeClazzName <--
func (m *TLProcessorDerivative) ProcessorDerivativeClazzName() string {
	return ClazzName_processorDerivative
}

// ClazzName <--
func (m *TLProcessorDerivative) ClazzName() string {
	return m.ClazzName2
}

// ToProcessorDerivative <--
func (m *TLProcessorDerivative) ToProcessorDerivative() *ProcessorDerivative {
	if m == nil {
		return nil
	}

	return m

}

func (m *TLProcessorDerivative) CalcSize(layer int32) int {
	switch clazzId := iface.GetClazzIDByName(ClazzName_processorDerivative, int(layer)); clazzId {
	case 0x2af751de:
		size := 4
		size += iface.CalcStringSize(m.Kind)
		size += iface.CalcStringSize(m.ObjectId)
		size += iface.CalcStringSize(m.FileName)
		size += iface.CalcStringSize(m.MimeType)
		size += 8
		size += 4
		size += 4
		size += iface.CalcBytesSize(m.Bytes)

		return size
	default:
		return 0
	}
}

func (m *TLProcessorDerivative) Validate(layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_processorDerivative, int(layer)); clazzId {
	case 0x2af751de:
		if err := iface.ValidateRequiredString("kind", m.Kind); err != nil {
			return err
		}

		if err := iface.ValidateRequiredString("object_id", m.ObjectId); err != nil {
			return err
		}

		if err := iface.ValidateRequiredString("file_name", m.FileName); err != nil {
			return err
		}

		if err := iface.ValidateRequiredString("mime_type", m.MimeType); err != nil {
			return err
		}

		if err := iface.ValidateRequiredBytes("bytes", m.Bytes); err != nil {
			return err
		}

		return nil
	default:
		return fmt.Errorf("unable to encode processorDerivative: unsupported layer %d", layer)
	}
}

// Encode <--
func (m *TLProcessorDerivative) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_processorDerivative, int(layer)); clazzId {
	case 0x2af751de:
		x.PutClazzID(0x2af751de)

		x.PutString(m.Kind)
		x.PutString(m.ObjectId)
		x.PutString(m.FileName)
		x.PutString(m.MimeType)
		x.PutInt64(m.Size2)
		x.PutInt32(m.Width)
		x.PutInt32(m.Height)
		x.PutBytes(m.Bytes)

		return nil
	default:
		return fmt.Errorf("unable to encode processorDerivative: unsupported layer %d", layer)
	}
}

// Decode <--
func (m *TLProcessorDerivative) Decode(d *bin.Decoder) (err error) {
	switch m.ClazzID {
	case 0x2af751de:
		m.Kind, err = d.String()
		if err != nil {
			return fmt.Errorf("unable to decode processorDerivative#0x2af751de: field kind: %w", err)
		}
		m.ObjectId, err = d.String()
		if err != nil {
			return fmt.Errorf("unable to decode processorDerivative#0x2af751de: field object_id: %w", err)
		}
		m.FileName, err = d.String()
		if err != nil {
			return fmt.Errorf("unable to decode processorDerivative#0x2af751de: field file_name: %w", err)
		}
		m.MimeType, err = d.String()
		if err != nil {
			return fmt.Errorf("unable to decode processorDerivative#0x2af751de: field mime_type: %w", err)
		}
		m.Size2, err = d.Int64()
		if err != nil {
			return fmt.Errorf("unable to decode processorDerivative#0x2af751de: field size2: %w", err)
		}
		m.Width, err = d.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode processorDerivative#0x2af751de: field width: %w", err)
		}
		m.Height, err = d.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode processorDerivative#0x2af751de: field height: %w", err)
		}
		m.Bytes, err = d.Bytes()
		if err != nil {
			return fmt.Errorf("unable to decode processorDerivative#0x2af751de: field bytes: %w", err)
		}

		return nil
	default:
		return fmt.Errorf("unable to decode processorDerivative: invalid constructor %x", m.ClazzID)
	}
}

// ProcessorDerivative <--
type ProcessorDerivative = TLProcessorDerivative
