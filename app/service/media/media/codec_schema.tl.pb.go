/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2024-present,  Teamgram Authors.
 *  All rights reserved.
 *
 * Author: Benqi (wubenqi@gmail.com)
 */

// ConstructorList
// RequestList

package media

import (
	"fmt"

	"github.com/teamgram/proto/mtproto"

	"google.golang.org/protobuf/types/known/wrapperspb"
)

//////////////////////////////////////////////////////////////////////////////////////////

var _ *wrapperspb.Int32Value
var _ *mtproto.Bool
var _ fmt.Stringer

var clazzIdRegisters2 = map[int32]func() mtproto.TLObject{
	// Constructor
	108083635: func() mtproto.TLObject { // 0x67139b3
		o := MakeTLPhotoSizeList(nil)
		o.Data2.Constructor = 108083635
		return o
	},
	953261042: func() mtproto.TLObject { // 0x38d19bf2
		o := MakeTLVideoSizeList(nil)
		o.Data2.Constructor = 953261042
		return o
	},

	// Method
	1009453847: func() mtproto.TLObject { // 0x3c2b0b17
		return &TLMediaUploadPhotoFile{
			Constructor: 1009453847,
		}
	},
	-1757466844: func() mtproto.TLObject { // 0x973f2f24
		return &TLMediaUploadProfilePhotoFile{
			Constructor: -1757466844,
		}
	},
	1702803563: func() mtproto.TLObject { // 0x657eb86b
		return &TLMediaGetPhoto{
			Constructor: 1702803563,
		}
	},
	-1578401979: func() mtproto.TLObject { // 0xa1eb7f45
		return &TLMediaGetPhotoSizeList{
			Constructor: -1578401979,
		}
	},
	-77823776: func() mtproto.TLObject { // 0xfb5c80e0
		return &TLMediaGetPhotoSizeListList{
			Constructor: -77823776,
		}
	},
	-998862102: func() mtproto.TLObject { // 0xc47692ea
		return &TLMediaGetVideoSizeList{
			Constructor: -998862102,
		}
	},
	1331671148: func() mtproto.TLObject { // 0x4f5fb06c
		return &TLMediaUploadedDocumentMedia{
			Constructor: 1331671148,
		}
	},
	1072011085: func() mtproto.TLObject { // 0x3fe5974d
		return &TLMediaGetDocument{
			Constructor: 1072011085,
		}
	},
	-986721681: func() mtproto.TLObject { // 0xc52fd26f
		return &TLMediaGetDocumentList{
			Constructor: -986721681,
		}
	},
	-1426012517: func() mtproto.TLObject { // 0xab00c69b
		return &TLMediaUploadEncryptedFile{
			Constructor: -1426012517,
		}
	},
	-60784431: func() mtproto.TLObject { // 0xfc6080d1
		return &TLMediaGetEncryptedFile{
			Constructor: -60784431,
		}
	},
	-1661293058: func() mtproto.TLObject { // 0x9cfaadfe
		return &TLMediaUploadWallPaperFile{
			Constructor: -1661293058,
		}
	},
	1122416736: func() mtproto.TLObject { // 0x42e6b860
		return &TLMediaUploadThemeFile{
			Constructor: 1122416736,
		}
	},
	-1397349139: func() mtproto.TLObject { // 0xacb624ed
		return &TLMediaUploadStickerFile{
			Constructor: -1397349139,
		}
	},
	1035645449: func() mtproto.TLObject { // 0x3dbab209
		return &TLMediaUploadRingtoneFile{
			Constructor: 1035645449,
		}
	},
}

func NewTLObjectByClassID(classId int32) mtproto.TLObject {
	f, ok := clazzIdRegisters2[classId]
	if !ok {
		return nil
	}
	return f()
}

func CheckClassID(classId int32) (ok bool) {
	_, ok = clazzIdRegisters2[classId]
	return
}

//----------------------------------------------------------------------------------------------------------------

///////////////////////////////////////////////////////////////////////////////
// PhotoSizeList <--
//  + TL_PhotoSizeList
//

func (m *PhotoSizeList) Encode(x *mtproto.EncodeBuf, layer int32) []byte {
	predicateName := m.PredicateName
	if predicateName == "" {
		if n, ok := clazzIdNameRegisters2[int32(m.Constructor)]; ok {
			predicateName = n
		}
	}

	switch predicateName {
	case Predicate_photoSizeList:
		t := m.To_PhotoSizeList()
		t.Encode(x, layer)

	default:
		// logx.Errorf("invalid predicate error: %s",  m.PredicateName)
		return nil
	}

	return nil
}

func (m *PhotoSizeList) CalcByteSize(layer int32) int {
	return 0
}

func (m *PhotoSizeList) Decode(dBuf *mtproto.DecodeBuf) error {
	m.Constructor = TLConstructor(dBuf.Int())
	switch uint32(m.Constructor) {
	case 0x67139b3:
		m2 := MakeTLPhotoSizeList(m)
		m2.Decode(dBuf)

	default:
		return fmt.Errorf("invalid constructorId: 0x%x", uint32(m.Constructor))
	}
	return dBuf.GetError()
}

// To_PhotoSizeList
func (m *PhotoSizeList) To_PhotoSizeList() *TLPhotoSizeList {
	m.PredicateName = Predicate_photoSizeList
	return &TLPhotoSizeList{
		Data2: m,
	}
}

// MakeTLPhotoSizeList
func MakeTLPhotoSizeList(data2 *PhotoSizeList) *TLPhotoSizeList {
	if data2 == nil {
		return &TLPhotoSizeList{Data2: &PhotoSizeList{
			PredicateName: Predicate_photoSizeList,
		}}
	} else {
		data2.PredicateName = Predicate_photoSizeList
		return &TLPhotoSizeList{Data2: data2}
	}
}

func (m *TLPhotoSizeList) To_PhotoSizeList() *PhotoSizeList {
	m.Data2.PredicateName = Predicate_photoSizeList
	return m.Data2
}

func (m *TLPhotoSizeList) SetSizeId(v int64) { m.Data2.SizeId = v }
func (m *TLPhotoSizeList) GetSizeId() int64  { return m.Data2.SizeId }

func (m *TLPhotoSizeList) SetSizes(v []*mtproto.PhotoSize) { m.Data2.Sizes = v }
func (m *TLPhotoSizeList) GetSizes() []*mtproto.PhotoSize  { return m.Data2.Sizes }

func (m *TLPhotoSizeList) SetDcId(v int32) { m.Data2.DcId = v }
func (m *TLPhotoSizeList) GetDcId() int32  { return m.Data2.DcId }

func (m *TLPhotoSizeList) GetPredicateName() string {
	return Predicate_photoSizeList
}

func (m *TLPhotoSizeList) Encode(x *mtproto.EncodeBuf, layer int32) error {
	var encodeF = map[uint32]func() error{
		0x67139b3: func() error {
			x.UInt(0x67139b3)

			x.Long(m.GetSizeId())

			x.Int(int32(mtproto.CRC32_vector))
			x.Int(int32(len(m.GetSizes())))
			for _, v := range m.GetSizes() {
				v.Encode(x, layer)
			}

			x.Int(m.GetDcId())
			return nil
		},
	}

	clazzId := GetClazzID(Predicate_photoSizeList, int(layer))
	if f, ok := encodeF[uint32(clazzId)]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		// log.Errorf("not found clazzId by (%s, %d)", Predicate_photoSizeList, layer)
		return nil
	}

	return nil
}

func (m *TLPhotoSizeList) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLPhotoSizeList) Decode(dBuf *mtproto.DecodeBuf) error {
	var decodeF = map[uint32]func() error{
		0x67139b3: func() error {
			m.SetSizeId(dBuf.Long())
			c1 := dBuf.Int()
			if c1 != int32(mtproto.CRC32_vector) {
				// dBuf.err = fmt.Errorf("invalid mtproto.CRC32_vector, c%d: %d", 1, c1)
				return fmt.Errorf("invalid mtproto.CRC32_vector, c%d: %d", 1, c1)
			}
			l1 := dBuf.Int()
			v1 := make([]*mtproto.PhotoSize, l1)
			for i := int32(0); i < l1; i++ {
				v1[i] = &mtproto.PhotoSize{}
				v1[i].Decode(dBuf)
			}
			m.SetSizes(v1)

			m.SetDcId(dBuf.Int())
			return dBuf.GetError()
		},
	}

	if f, ok := decodeF[uint32(m.Data2.Constructor)]; ok {
		return f()
	} else {
		return fmt.Errorf("invalid constructor: %x", uint32(m.Data2.Constructor))
	}
}

///////////////////////////////////////////////////////////////////////////////
// VideoSizeList <--
//  + TL_VideoSizeList
//

func (m *VideoSizeList) Encode(x *mtproto.EncodeBuf, layer int32) []byte {
	predicateName := m.PredicateName
	if predicateName == "" {
		if n, ok := clazzIdNameRegisters2[int32(m.Constructor)]; ok {
			predicateName = n
		}
	}

	switch predicateName {
	case Predicate_videoSizeList:
		t := m.To_VideoSizeList()
		t.Encode(x, layer)

	default:
		// logx.Errorf("invalid predicate error: %s",  m.PredicateName)
		return nil
	}

	return nil
}

func (m *VideoSizeList) CalcByteSize(layer int32) int {
	return 0
}

func (m *VideoSizeList) Decode(dBuf *mtproto.DecodeBuf) error {
	m.Constructor = TLConstructor(dBuf.Int())
	switch uint32(m.Constructor) {
	case 0x38d19bf2:
		m2 := MakeTLVideoSizeList(m)
		m2.Decode(dBuf)

	default:
		return fmt.Errorf("invalid constructorId: 0x%x", uint32(m.Constructor))
	}
	return dBuf.GetError()
}

// To_VideoSizeList
func (m *VideoSizeList) To_VideoSizeList() *TLVideoSizeList {
	m.PredicateName = Predicate_videoSizeList
	return &TLVideoSizeList{
		Data2: m,
	}
}

// MakeTLVideoSizeList
func MakeTLVideoSizeList(data2 *VideoSizeList) *TLVideoSizeList {
	if data2 == nil {
		return &TLVideoSizeList{Data2: &VideoSizeList{
			PredicateName: Predicate_videoSizeList,
		}}
	} else {
		data2.PredicateName = Predicate_videoSizeList
		return &TLVideoSizeList{Data2: data2}
	}
}

func (m *TLVideoSizeList) To_VideoSizeList() *VideoSizeList {
	m.Data2.PredicateName = Predicate_videoSizeList
	return m.Data2
}

func (m *TLVideoSizeList) SetSizeId(v int64) { m.Data2.SizeId = v }
func (m *TLVideoSizeList) GetSizeId() int64  { return m.Data2.SizeId }

func (m *TLVideoSizeList) SetSizes(v []*mtproto.VideoSize) { m.Data2.Sizes = v }
func (m *TLVideoSizeList) GetSizes() []*mtproto.VideoSize  { return m.Data2.Sizes }

func (m *TLVideoSizeList) SetDcId(v int32) { m.Data2.DcId = v }
func (m *TLVideoSizeList) GetDcId() int32  { return m.Data2.DcId }

func (m *TLVideoSizeList) GetPredicateName() string {
	return Predicate_videoSizeList
}

func (m *TLVideoSizeList) Encode(x *mtproto.EncodeBuf, layer int32) error {
	var encodeF = map[uint32]func() error{
		0x38d19bf2: func() error {
			x.UInt(0x38d19bf2)

			x.Long(m.GetSizeId())

			x.Int(int32(mtproto.CRC32_vector))
			x.Int(int32(len(m.GetSizes())))
			for _, v := range m.GetSizes() {
				v.Encode(x, layer)
			}

			x.Int(m.GetDcId())
			return nil
		},
	}

	clazzId := GetClazzID(Predicate_videoSizeList, int(layer))
	if f, ok := encodeF[uint32(clazzId)]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		// log.Errorf("not found clazzId by (%s, %d)", Predicate_videoSizeList, layer)
		return nil
	}

	return nil
}

func (m *TLVideoSizeList) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLVideoSizeList) Decode(dBuf *mtproto.DecodeBuf) error {
	var decodeF = map[uint32]func() error{
		0x38d19bf2: func() error {
			m.SetSizeId(dBuf.Long())
			c1 := dBuf.Int()
			if c1 != int32(mtproto.CRC32_vector) {
				// dBuf.err = fmt.Errorf("invalid mtproto.CRC32_vector, c%d: %d", 1, c1)
				return fmt.Errorf("invalid mtproto.CRC32_vector, c%d: %d", 1, c1)
			}
			l1 := dBuf.Int()
			v1 := make([]*mtproto.VideoSize, l1)
			for i := int32(0); i < l1; i++ {
				v1[i] = &mtproto.VideoSize{}
				v1[i].Decode(dBuf)
			}
			m.SetSizes(v1)

			m.SetDcId(dBuf.Int())
			return dBuf.GetError()
		},
	}

	if f, ok := decodeF[uint32(m.Data2.Constructor)]; ok {
		return f()
	} else {
		return fmt.Errorf("invalid constructor: %x", uint32(m.Data2.Constructor))
	}
}

//----------------------------------------------------------------------------------------------------------------
// TLMediaUploadPhotoFile
///////////////////////////////////////////////////////////////////////////////

func (m *TLMediaUploadPhotoFile) Encode(x *mtproto.EncodeBuf, layer int32) error {
	switch uint32(m.Constructor) {
	case 0x3c2b0b17:
		x.UInt(0x3c2b0b17)

		// set flags
		var flags uint32 = 0

		if m.GetStickers() != nil {
			flags |= 1 << 0
		}
		if m.GetTtlSeconds() != nil {
			flags |= 1 << 1
		}

		x.UInt(flags)

		// flags Debug by @benqi
		x.Long(m.GetOwnerId())
		m.GetFile().Encode(x, layer)
		if m.GetStickers() != nil {
			x.Int(int32(mtproto.CRC32_vector))
			x.Int(int32(len(m.GetStickers())))
			for _, v := range m.GetStickers() {
				v.Encode(x, layer)
			}
		}
		if m.GetTtlSeconds() != nil {
			x.Int(m.GetTtlSeconds().Value)
		}

	default:
		// log.Errorf("")
	}

	return nil
}

func (m *TLMediaUploadPhotoFile) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLMediaUploadPhotoFile) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0x3c2b0b17:

		flags := dBuf.UInt()
		_ = flags

		// flags Debug by @benqi
		m.OwnerId = dBuf.Long()

		m3 := &mtproto.InputFile{}
		m3.Decode(dBuf)
		m.File = m3

		if (flags & (1 << 0)) != 0 {
			c4 := dBuf.Int()
			if c4 != int32(mtproto.CRC32_vector) {
				// dBuf.err = fmt.Errorf("invalid mtproto.CRC32_vector, c%d: %d", 4, c4)
				return fmt.Errorf("invalid mtproto.CRC32_vector, c%d: %d", 4, c4)
			}
			l4 := dBuf.Int()
			v4 := make([]*mtproto.InputDocument, l4)
			for i := int32(0); i < l4; i++ {
				v4[i] = &mtproto.InputDocument{}
				v4[i].Decode(dBuf)
			}
			m.Stickers = v4
		}
		if (flags & (1 << 1)) != 0 {
			m.TtlSeconds = &wrapperspb.Int32Value{Value: dBuf.Int()}
		}

		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

// TLMediaUploadProfilePhotoFile
///////////////////////////////////////////////////////////////////////////////

func (m *TLMediaUploadProfilePhotoFile) Encode(x *mtproto.EncodeBuf, layer int32) error {
	switch uint32(m.Constructor) {
	case 0x973f2f24:
		x.UInt(0x973f2f24)

		// set flags
		var flags uint32 = 0

		if m.GetFile() != nil {
			flags |= 1 << 0
		}
		if m.GetVideo() != nil {
			flags |= 1 << 1
		}
		if m.GetVideoStartTs() != nil {
			flags |= 1 << 2
		}

		x.UInt(flags)

		// flags Debug by @benqi
		x.Long(m.GetOwnerId())
		if m.GetFile() != nil {
			m.GetFile().Encode(x, layer)
		}

		if m.GetVideo() != nil {
			m.GetVideo().Encode(x, layer)
		}

		if m.GetVideoStartTs() != nil {
			x.Double(m.GetVideoStartTs().Value)
		}

	default:
		// log.Errorf("")
	}

	return nil
}

func (m *TLMediaUploadProfilePhotoFile) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLMediaUploadProfilePhotoFile) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0x973f2f24:

		flags := dBuf.UInt()
		_ = flags

		// flags Debug by @benqi
		m.OwnerId = dBuf.Long()
		if (flags & (1 << 0)) != 0 {
			m3 := &mtproto.InputFile{}
			m3.Decode(dBuf)
			m.File = m3
		}
		if (flags & (1 << 1)) != 0 {
			m4 := &mtproto.InputFile{}
			m4.Decode(dBuf)
			m.Video = m4
		}
		if (flags & (1 << 2)) != 0 {
			m.VideoStartTs = &wrapperspb.DoubleValue{Value: dBuf.Double()}
		}

		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

// TLMediaGetPhoto
///////////////////////////////////////////////////////////////////////////////

func (m *TLMediaGetPhoto) Encode(x *mtproto.EncodeBuf, layer int32) error {
	switch uint32(m.Constructor) {
	case 0x657eb86b:
		x.UInt(0x657eb86b)

		// no flags

		x.Long(m.GetPhotoId())

	default:
		// log.Errorf("")
	}

	return nil
}

func (m *TLMediaGetPhoto) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLMediaGetPhoto) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0x657eb86b:

		// not has flags

		m.PhotoId = dBuf.Long()
		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

// TLMediaGetPhotoSizeList
///////////////////////////////////////////////////////////////////////////////

func (m *TLMediaGetPhotoSizeList) Encode(x *mtproto.EncodeBuf, layer int32) error {
	switch uint32(m.Constructor) {
	case 0xa1eb7f45:
		x.UInt(0xa1eb7f45)

		// no flags

		x.Long(m.GetSizeId())

	default:
		// log.Errorf("")
	}

	return nil
}

func (m *TLMediaGetPhotoSizeList) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLMediaGetPhotoSizeList) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0xa1eb7f45:

		// not has flags

		m.SizeId = dBuf.Long()
		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

// TLMediaGetPhotoSizeListList
///////////////////////////////////////////////////////////////////////////////

func (m *TLMediaGetPhotoSizeListList) Encode(x *mtproto.EncodeBuf, layer int32) error {
	switch uint32(m.Constructor) {
	case 0xfb5c80e0:
		x.UInt(0xfb5c80e0)

		// no flags

		x.VectorLong(m.GetIdList())

	default:
		// log.Errorf("")
	}

	return nil
}

func (m *TLMediaGetPhotoSizeListList) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLMediaGetPhotoSizeListList) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0xfb5c80e0:

		// not has flags

		m.IdList = dBuf.VectorLong()

		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

// TLMediaGetVideoSizeList
///////////////////////////////////////////////////////////////////////////////

func (m *TLMediaGetVideoSizeList) Encode(x *mtproto.EncodeBuf, layer int32) error {
	switch uint32(m.Constructor) {
	case 0xc47692ea:
		x.UInt(0xc47692ea)

		// no flags

		x.Long(m.GetSizeId())

	default:
		// log.Errorf("")
	}

	return nil
}

func (m *TLMediaGetVideoSizeList) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLMediaGetVideoSizeList) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0xc47692ea:

		// not has flags

		m.SizeId = dBuf.Long()
		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

// TLMediaUploadedDocumentMedia
///////////////////////////////////////////////////////////////////////////////

func (m *TLMediaUploadedDocumentMedia) Encode(x *mtproto.EncodeBuf, layer int32) error {
	switch uint32(m.Constructor) {
	case 0x4f5fb06c:
		x.UInt(0x4f5fb06c)

		// no flags

		x.Long(m.GetOwnerId())
		m.GetMedia().Encode(x, layer)

	default:
		// log.Errorf("")
	}

	return nil
}

func (m *TLMediaUploadedDocumentMedia) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLMediaUploadedDocumentMedia) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0x4f5fb06c:

		// not has flags

		m.OwnerId = dBuf.Long()

		m2 := &mtproto.InputMedia{}
		m2.Decode(dBuf)
		m.Media = m2

		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

// TLMediaGetDocument
///////////////////////////////////////////////////////////////////////////////

func (m *TLMediaGetDocument) Encode(x *mtproto.EncodeBuf, layer int32) error {
	switch uint32(m.Constructor) {
	case 0x3fe5974d:
		x.UInt(0x3fe5974d)

		// no flags

		x.Long(m.GetId())

	default:
		// log.Errorf("")
	}

	return nil
}

func (m *TLMediaGetDocument) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLMediaGetDocument) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0x3fe5974d:

		// not has flags

		m.Id = dBuf.Long()
		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

// TLMediaGetDocumentList
///////////////////////////////////////////////////////////////////////////////

func (m *TLMediaGetDocumentList) Encode(x *mtproto.EncodeBuf, layer int32) error {
	switch uint32(m.Constructor) {
	case 0xc52fd26f:
		x.UInt(0xc52fd26f)

		// no flags

		x.VectorLong(m.GetIdList())

	default:
		// log.Errorf("")
	}

	return nil
}

func (m *TLMediaGetDocumentList) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLMediaGetDocumentList) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0xc52fd26f:

		// not has flags

		m.IdList = dBuf.VectorLong()

		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

// TLMediaUploadEncryptedFile
///////////////////////////////////////////////////////////////////////////////

func (m *TLMediaUploadEncryptedFile) Encode(x *mtproto.EncodeBuf, layer int32) error {
	switch uint32(m.Constructor) {
	case 0xab00c69b:
		x.UInt(0xab00c69b)

		// no flags

		x.Long(m.GetOwnerId())
		m.GetFile().Encode(x, layer)

	default:
		// log.Errorf("")
	}

	return nil
}

func (m *TLMediaUploadEncryptedFile) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLMediaUploadEncryptedFile) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0xab00c69b:

		// not has flags

		m.OwnerId = dBuf.Long()

		m2 := &mtproto.InputEncryptedFile{}
		m2.Decode(dBuf)
		m.File = m2

		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

// TLMediaGetEncryptedFile
///////////////////////////////////////////////////////////////////////////////

func (m *TLMediaGetEncryptedFile) Encode(x *mtproto.EncodeBuf, layer int32) error {
	switch uint32(m.Constructor) {
	case 0xfc6080d1:
		x.UInt(0xfc6080d1)

		// no flags

		x.Long(m.GetId())
		x.Long(m.GetAccessHash())

	default:
		// log.Errorf("")
	}

	return nil
}

func (m *TLMediaGetEncryptedFile) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLMediaGetEncryptedFile) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0xfc6080d1:

		// not has flags

		m.Id = dBuf.Long()
		m.AccessHash = dBuf.Long()
		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

// TLMediaUploadWallPaperFile
///////////////////////////////////////////////////////////////////////////////

func (m *TLMediaUploadWallPaperFile) Encode(x *mtproto.EncodeBuf, layer int32) error {
	switch uint32(m.Constructor) {
	case 0x9cfaadfe:
		x.UInt(0x9cfaadfe)

		// no flags

		x.Long(m.GetOwnerId())
		m.GetFile().Encode(x, layer)
		x.String(m.GetMimeType())
		m.GetAdmin().Encode(x, layer)

	default:
		// log.Errorf("")
	}

	return nil
}

func (m *TLMediaUploadWallPaperFile) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLMediaUploadWallPaperFile) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0x9cfaadfe:

		// not has flags

		m.OwnerId = dBuf.Long()

		m2 := &mtproto.InputFile{}
		m2.Decode(dBuf)
		m.File = m2

		m.MimeType = dBuf.String()

		m4 := &mtproto.Bool{}
		m4.Decode(dBuf)
		m.Admin = m4

		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

// TLMediaUploadThemeFile
///////////////////////////////////////////////////////////////////////////////

func (m *TLMediaUploadThemeFile) Encode(x *mtproto.EncodeBuf, layer int32) error {
	switch uint32(m.Constructor) {
	case 0x42e6b860:
		x.UInt(0x42e6b860)

		// set flags
		var flags uint32 = 0

		if m.GetThumb() != nil {
			flags |= 1 << 0
		}

		x.UInt(flags)

		// flags Debug by @benqi
		x.Long(m.GetOwnerId())
		m.GetFile().Encode(x, layer)
		if m.GetThumb() != nil {
			m.GetThumb().Encode(x, layer)
		}

		x.String(m.GetMimeType())
		x.String(m.GetFileName())

	default:
		// log.Errorf("")
	}

	return nil
}

func (m *TLMediaUploadThemeFile) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLMediaUploadThemeFile) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0x42e6b860:

		flags := dBuf.UInt()
		_ = flags

		// flags Debug by @benqi
		m.OwnerId = dBuf.Long()

		m3 := &mtproto.InputFile{}
		m3.Decode(dBuf)
		m.File = m3

		if (flags & (1 << 0)) != 0 {
			m4 := &mtproto.InputFile{}
			m4.Decode(dBuf)
			m.Thumb = m4
		}
		m.MimeType = dBuf.String()
		m.FileName = dBuf.String()
		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

// TLMediaUploadStickerFile
///////////////////////////////////////////////////////////////////////////////

func (m *TLMediaUploadStickerFile) Encode(x *mtproto.EncodeBuf, layer int32) error {
	switch uint32(m.Constructor) {
	case 0xacb624ed:
		x.UInt(0xacb624ed)

		// set flags
		var flags uint32 = 0

		if m.GetThumb() != nil {
			flags |= 1 << 0
		}

		x.UInt(flags)

		// flags Debug by @benqi
		x.Long(m.GetOwnerId())
		m.GetFile().Encode(x, layer)
		if m.GetThumb() != nil {
			m.GetThumb().Encode(x, layer)
		}

		x.String(m.GetMimeType())
		x.String(m.GetFileName())
		m.GetDocumentAttributeSticker().Encode(x, layer)

	default:
		// log.Errorf("")
	}

	return nil
}

func (m *TLMediaUploadStickerFile) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLMediaUploadStickerFile) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0xacb624ed:

		flags := dBuf.UInt()
		_ = flags

		// flags Debug by @benqi
		m.OwnerId = dBuf.Long()

		m3 := &mtproto.InputFile{}
		m3.Decode(dBuf)
		m.File = m3

		if (flags & (1 << 0)) != 0 {
			m4 := &mtproto.InputFile{}
			m4.Decode(dBuf)
			m.Thumb = m4
		}
		m.MimeType = dBuf.String()
		m.FileName = dBuf.String()

		m7 := &mtproto.DocumentAttribute{}
		m7.Decode(dBuf)
		m.DocumentAttributeSticker = m7

		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

// TLMediaUploadRingtoneFile
///////////////////////////////////////////////////////////////////////////////

func (m *TLMediaUploadRingtoneFile) Encode(x *mtproto.EncodeBuf, layer int32) error {
	switch uint32(m.Constructor) {
	case 0x3dbab209:
		x.UInt(0x3dbab209)

		// set flags
		var flags uint32 = 0

		x.UInt(flags)

		// flags Debug by @benqi
		x.Long(m.GetOwnerId())
		m.GetFile().Encode(x, layer)
		x.String(m.GetMimeType())
		x.String(m.GetFileName())

	default:
		// log.Errorf("")
	}

	return nil
}

func (m *TLMediaUploadRingtoneFile) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLMediaUploadRingtoneFile) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0x3dbab209:

		flags := dBuf.UInt()
		_ = flags

		// flags Debug by @benqi
		m.OwnerId = dBuf.Long()

		m3 := &mtproto.InputFile{}
		m3.Decode(dBuf)
		m.File = m3

		m.MimeType = dBuf.String()
		m.FileName = dBuf.String()
		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

// Vector_PhotoSizeList
// /////////////////////////////////////////////////////////////////////////////
func (m *Vector_PhotoSizeList) Encode(x *mtproto.EncodeBuf, layer int32) error {
	x.Int(int32(mtproto.CRC32_vector))
	x.Int(int32(len(m.Datas)))
	for _, v := range m.Datas {
		v.Encode(x, layer)
	}

	return nil
}

func (m *Vector_PhotoSizeList) Decode(dBuf *mtproto.DecodeBuf) error {
	dBuf.Int() // TODO(@benqi): Check crc32 invalid
	l1 := dBuf.Int()
	m.Datas = make([]*PhotoSizeList, l1)
	for i := int32(0); i < l1; i++ {
		m.Datas[i] = new(PhotoSizeList)
		(*m.Datas[i]).Decode(dBuf)
	}

	return dBuf.GetError()
}

func (m *Vector_PhotoSizeList) CalcByteSize(layer int32) int {
	return 0
}

// Vector_Document
// /////////////////////////////////////////////////////////////////////////////
func (m *Vector_Document) Encode(x *mtproto.EncodeBuf, layer int32) error {
	x.Int(int32(mtproto.CRC32_vector))
	x.Int(int32(len(m.Datas)))
	for _, v := range m.Datas {
		v.Encode(x, layer)
	}

	return nil
}

func (m *Vector_Document) Decode(dBuf *mtproto.DecodeBuf) error {
	dBuf.Int() // TODO(@benqi): Check crc32 invalid
	l1 := dBuf.Int()
	m.Datas = make([]*mtproto.Document, l1)
	for i := int32(0); i < l1; i++ {
		m.Datas[i] = new(mtproto.Document)
		(*m.Datas[i]).Decode(dBuf)
	}

	return dBuf.GetError()
}

func (m *Vector_Document) CalcByteSize(layer int32) int {
	return 0
}
