/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2022-present,  Teamgram Authors.
 *  All rights reserved.
 *
 * Author: teagramio (teagram.io@gmail.com)
 */

// ConstructorList
// RequestList

package media

import (
	"fmt"

	"github.com/teamgram/proto/mtproto"

	"github.com/gogo/protobuf/jsonpb"
	"github.com/gogo/protobuf/types"
)

//////////////////////////////////////////////////////////////////////////////////////////

var _ *types.Int32Value
var _ *mtproto.Bool
var _ fmt.GoStringer

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

func (m *PhotoSizeList) Encode(layer int32) []byte {
	predicateName := m.PredicateName
	if predicateName == "" {
		if n, ok := clazzIdNameRegisters2[int32(m.Constructor)]; ok {
			predicateName = n
		}
	}

	var (
		xBuf []byte
	)

	switch predicateName {
	case Predicate_photoSizeList:
		t := m.To_PhotoSizeList()
		xBuf = t.Encode(layer)

	default:
		// logx.Errorf("invalid predicate error: %s",  m.PredicateName)
		return []byte{}
	}

	return xBuf
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

func (m *PhotoSizeList) DebugString() string {
	switch m.PredicateName {
	case Predicate_photoSizeList:
		t := m.To_PhotoSizeList()
		return t.DebugString()

	default:
		return "{}"
	}
}

// To_PhotoSizeList
// photoSizeList size_id:long sizes:Vector<PhotoSize> dc_id:int = PhotoSizeList;
func (m *PhotoSizeList) To_PhotoSizeList() *TLPhotoSizeList {
	m.PredicateName = Predicate_photoSizeList
	return &TLPhotoSizeList{
		Data2: m,
	}
}

// MakeTLPhotoSizeList
// photoSizeList size_id:long sizes:Vector<PhotoSize> dc_id:int = PhotoSizeList;
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

func (m *TLPhotoSizeList) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)

	var encodeF = map[uint32]func() []byte{
		0x67139b3: func() []byte {
			// photoSizeList size_id:long sizes:Vector<PhotoSize> dc_id:int = PhotoSizeList;
			x.UInt(0x67139b3)

			x.Long(m.GetSizeId())

			x.Int(int32(mtproto.CRC32_vector))
			x.Int(int32(len(m.GetSizes())))
			for _, v := range m.GetSizes() {
				x.Bytes((*v).Encode(layer))
			}

			x.Int(m.GetDcId())
			return x.GetBuf()
		},
	}

	clazzId := GetClazzID(Predicate_photoSizeList, int(layer))
	if f, ok := encodeF[uint32(clazzId)]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		// log.Errorf("not found clazzId by (%s, %d)", Predicate_photoSizeList, layer)
		return x.GetBuf()
	}

	return x.GetBuf()
}

func (m *TLPhotoSizeList) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLPhotoSizeList) Decode(dBuf *mtproto.DecodeBuf) error {
	var decodeF = map[uint32]func() error{
		0x67139b3: func() error {
			// photoSizeList size_id:long sizes:Vector<PhotoSize> dc_id:int = PhotoSizeList;
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

func (m *TLPhotoSizeList) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

///////////////////////////////////////////////////////////////////////////////
// VideoSizeList <--
//  + TL_VideoSizeList
//

func (m *VideoSizeList) Encode(layer int32) []byte {
	predicateName := m.PredicateName
	if predicateName == "" {
		if n, ok := clazzIdNameRegisters2[int32(m.Constructor)]; ok {
			predicateName = n
		}
	}

	var (
		xBuf []byte
	)

	switch predicateName {
	case Predicate_videoSizeList:
		t := m.To_VideoSizeList()
		xBuf = t.Encode(layer)

	default:
		// logx.Errorf("invalid predicate error: %s",  m.PredicateName)
		return []byte{}
	}

	return xBuf
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

func (m *VideoSizeList) DebugString() string {
	switch m.PredicateName {
	case Predicate_videoSizeList:
		t := m.To_VideoSizeList()
		return t.DebugString()

	default:
		return "{}"
	}
}

// To_VideoSizeList
// videoSizeList size_id:long sizes:Vector<VideoSize> dc_id:int = VideoSizeList;
func (m *VideoSizeList) To_VideoSizeList() *TLVideoSizeList {
	m.PredicateName = Predicate_videoSizeList
	return &TLVideoSizeList{
		Data2: m,
	}
}

// MakeTLVideoSizeList
// videoSizeList size_id:long sizes:Vector<VideoSize> dc_id:int = VideoSizeList;
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

func (m *TLVideoSizeList) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)

	var encodeF = map[uint32]func() []byte{
		0x38d19bf2: func() []byte {
			// videoSizeList size_id:long sizes:Vector<VideoSize> dc_id:int = VideoSizeList;
			x.UInt(0x38d19bf2)

			x.Long(m.GetSizeId())

			x.Int(int32(mtproto.CRC32_vector))
			x.Int(int32(len(m.GetSizes())))
			for _, v := range m.GetSizes() {
				x.Bytes((*v).Encode(layer))
			}

			x.Int(m.GetDcId())
			return x.GetBuf()
		},
	}

	clazzId := GetClazzID(Predicate_videoSizeList, int(layer))
	if f, ok := encodeF[uint32(clazzId)]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		// log.Errorf("not found clazzId by (%s, %d)", Predicate_videoSizeList, layer)
		return x.GetBuf()
	}

	return x.GetBuf()
}

func (m *TLVideoSizeList) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLVideoSizeList) Decode(dBuf *mtproto.DecodeBuf) error {
	var decodeF = map[uint32]func() error{
		0x38d19bf2: func() error {
			// videoSizeList size_id:long sizes:Vector<VideoSize> dc_id:int = VideoSizeList;
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

func (m *TLVideoSizeList) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

//----------------------------------------------------------------------------------------------------------------
// TLMediaUploadPhotoFile
///////////////////////////////////////////////////////////////////////////////

func (m *TLMediaUploadPhotoFile) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	// x.Int(int32(CRC32_media_uploadPhotoFile))

	switch uint32(m.Constructor) {
	case 0x3c2b0b17:
		// media.uploadPhotoFile flags:# owner_id:long file:InputFile stickers:flags.0?Vector<InputDocument> ttl_seconds:flags.1?int = Photo;
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
		x.Bytes(m.GetFile().Encode(layer))
		if m.GetStickers() != nil {
			x.Int(int32(mtproto.CRC32_vector))
			x.Int(int32(len(m.GetStickers())))
			for _, v := range m.GetStickers() {
				x.Bytes((*v).Encode(layer))
			}
		}
		if m.GetTtlSeconds() != nil {
			x.Int(m.GetTtlSeconds().Value)
		}

	default:
		// log.Errorf("")
	}

	return x.GetBuf()
}

func (m *TLMediaUploadPhotoFile) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLMediaUploadPhotoFile) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0x3c2b0b17:
		// media.uploadPhotoFile flags:# owner_id:long file:InputFile stickers:flags.0?Vector<InputDocument> ttl_seconds:flags.1?int = Photo;

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
			m.TtlSeconds = &types.Int32Value{Value: dBuf.Int()}
		}

		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

func (m *TLMediaUploadPhotoFile) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

// TLMediaUploadProfilePhotoFile
///////////////////////////////////////////////////////////////////////////////

func (m *TLMediaUploadProfilePhotoFile) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	// x.Int(int32(CRC32_media_uploadProfilePhotoFile))

	switch uint32(m.Constructor) {
	case 0x973f2f24:
		// media.uploadProfilePhotoFile flags:# owner_id:long file:flags.0?InputFile video:flags.1?InputFile video_start_ts:flags.2?double = Photo;
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
			x.Bytes(m.GetFile().Encode(layer))
		}

		if m.GetVideo() != nil {
			x.Bytes(m.GetVideo().Encode(layer))
		}

		if m.GetVideoStartTs() != nil {
			x.Double(m.GetVideoStartTs().Value)
		}

	default:
		// log.Errorf("")
	}

	return x.GetBuf()
}

func (m *TLMediaUploadProfilePhotoFile) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLMediaUploadProfilePhotoFile) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0x973f2f24:
		// media.uploadProfilePhotoFile flags:# owner_id:long file:flags.0?InputFile video:flags.1?InputFile video_start_ts:flags.2?double = Photo;

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
			m.VideoStartTs = &types.DoubleValue{Value: dBuf.Double()}
		}

		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

func (m *TLMediaUploadProfilePhotoFile) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

// TLMediaGetPhoto
///////////////////////////////////////////////////////////////////////////////

func (m *TLMediaGetPhoto) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	// x.Int(int32(CRC32_media_getPhoto))

	switch uint32(m.Constructor) {
	case 0x657eb86b:
		// media.getPhoto photo_id:long = Photo;
		x.UInt(0x657eb86b)

		// no flags

		x.Long(m.GetPhotoId())

	default:
		// log.Errorf("")
	}

	return x.GetBuf()
}

func (m *TLMediaGetPhoto) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLMediaGetPhoto) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0x657eb86b:
		// media.getPhoto photo_id:long = Photo;

		// not has flags

		m.PhotoId = dBuf.Long()
		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

func (m *TLMediaGetPhoto) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

// TLMediaGetPhotoSizeList
///////////////////////////////////////////////////////////////////////////////

func (m *TLMediaGetPhotoSizeList) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	// x.Int(int32(CRC32_media_getPhotoSizeList))

	switch uint32(m.Constructor) {
	case 0xa1eb7f45:
		// media.getPhotoSizeList size_id:long = PhotoSizeList;
		x.UInt(0xa1eb7f45)

		// no flags

		x.Long(m.GetSizeId())

	default:
		// log.Errorf("")
	}

	return x.GetBuf()
}

func (m *TLMediaGetPhotoSizeList) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLMediaGetPhotoSizeList) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0xa1eb7f45:
		// media.getPhotoSizeList size_id:long = PhotoSizeList;

		// not has flags

		m.SizeId = dBuf.Long()
		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

func (m *TLMediaGetPhotoSizeList) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

// TLMediaGetPhotoSizeListList
///////////////////////////////////////////////////////////////////////////////

func (m *TLMediaGetPhotoSizeListList) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	// x.Int(int32(CRC32_media_getPhotoSizeListList))

	switch uint32(m.Constructor) {
	case 0xfb5c80e0:
		// media.getPhotoSizeListList id_list:Vector<long> = Vector<PhotoSizeList>;
		x.UInt(0xfb5c80e0)

		// no flags

		x.VectorLong(m.GetIdList())

	default:
		// log.Errorf("")
	}

	return x.GetBuf()
}

func (m *TLMediaGetPhotoSizeListList) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLMediaGetPhotoSizeListList) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0xfb5c80e0:
		// media.getPhotoSizeListList id_list:Vector<long> = Vector<PhotoSizeList>;

		// not has flags

		m.IdList = dBuf.VectorLong()

		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

func (m *TLMediaGetPhotoSizeListList) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

// TLMediaGetVideoSizeList
///////////////////////////////////////////////////////////////////////////////

func (m *TLMediaGetVideoSizeList) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	// x.Int(int32(CRC32_media_getVideoSizeList))

	switch uint32(m.Constructor) {
	case 0xc47692ea:
		// media.getVideoSizeList size_id:long = VideoSizeList;
		x.UInt(0xc47692ea)

		// no flags

		x.Long(m.GetSizeId())

	default:
		// log.Errorf("")
	}

	return x.GetBuf()
}

func (m *TLMediaGetVideoSizeList) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLMediaGetVideoSizeList) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0xc47692ea:
		// media.getVideoSizeList size_id:long = VideoSizeList;

		// not has flags

		m.SizeId = dBuf.Long()
		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

func (m *TLMediaGetVideoSizeList) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

// TLMediaUploadedDocumentMedia
///////////////////////////////////////////////////////////////////////////////

func (m *TLMediaUploadedDocumentMedia) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	// x.Int(int32(CRC32_media_uploadedDocumentMedia))

	switch uint32(m.Constructor) {
	case 0x4f5fb06c:
		// media.uploadedDocumentMedia owner_id:long media:InputMedia = MessageMedia;
		x.UInt(0x4f5fb06c)

		// no flags

		x.Long(m.GetOwnerId())
		x.Bytes(m.GetMedia().Encode(layer))

	default:
		// log.Errorf("")
	}

	return x.GetBuf()
}

func (m *TLMediaUploadedDocumentMedia) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLMediaUploadedDocumentMedia) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0x4f5fb06c:
		// media.uploadedDocumentMedia owner_id:long media:InputMedia = MessageMedia;

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

func (m *TLMediaUploadedDocumentMedia) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

// TLMediaGetDocument
///////////////////////////////////////////////////////////////////////////////

func (m *TLMediaGetDocument) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	// x.Int(int32(CRC32_media_getDocument))

	switch uint32(m.Constructor) {
	case 0x3fe5974d:
		// media.getDocument id:long = Document;
		x.UInt(0x3fe5974d)

		// no flags

		x.Long(m.GetId())

	default:
		// log.Errorf("")
	}

	return x.GetBuf()
}

func (m *TLMediaGetDocument) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLMediaGetDocument) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0x3fe5974d:
		// media.getDocument id:long = Document;

		// not has flags

		m.Id = dBuf.Long()
		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

func (m *TLMediaGetDocument) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

// TLMediaGetDocumentList
///////////////////////////////////////////////////////////////////////////////

func (m *TLMediaGetDocumentList) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	// x.Int(int32(CRC32_media_getDocumentList))

	switch uint32(m.Constructor) {
	case 0xc52fd26f:
		// media.getDocumentList id_list:Vector<long> = Vector<Document>;
		x.UInt(0xc52fd26f)

		// no flags

		x.VectorLong(m.GetIdList())

	default:
		// log.Errorf("")
	}

	return x.GetBuf()
}

func (m *TLMediaGetDocumentList) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLMediaGetDocumentList) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0xc52fd26f:
		// media.getDocumentList id_list:Vector<long> = Vector<Document>;

		// not has flags

		m.IdList = dBuf.VectorLong()

		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

func (m *TLMediaGetDocumentList) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

// TLMediaUploadEncryptedFile
///////////////////////////////////////////////////////////////////////////////

func (m *TLMediaUploadEncryptedFile) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	// x.Int(int32(CRC32_media_uploadEncryptedFile))

	switch uint32(m.Constructor) {
	case 0xab00c69b:
		// media.uploadEncryptedFile owner_id:long file:InputEncryptedFile = EncryptedFile;
		x.UInt(0xab00c69b)

		// no flags

		x.Long(m.GetOwnerId())
		x.Bytes(m.GetFile().Encode(layer))

	default:
		// log.Errorf("")
	}

	return x.GetBuf()
}

func (m *TLMediaUploadEncryptedFile) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLMediaUploadEncryptedFile) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0xab00c69b:
		// media.uploadEncryptedFile owner_id:long file:InputEncryptedFile = EncryptedFile;

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

func (m *TLMediaUploadEncryptedFile) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

// TLMediaGetEncryptedFile
///////////////////////////////////////////////////////////////////////////////

func (m *TLMediaGetEncryptedFile) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	// x.Int(int32(CRC32_media_getEncryptedFile))

	switch uint32(m.Constructor) {
	case 0xfc6080d1:
		// media.getEncryptedFile id:long access_hash:long = EncryptedFile;
		x.UInt(0xfc6080d1)

		// no flags

		x.Long(m.GetId())
		x.Long(m.GetAccessHash())

	default:
		// log.Errorf("")
	}

	return x.GetBuf()
}

func (m *TLMediaGetEncryptedFile) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLMediaGetEncryptedFile) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0xfc6080d1:
		// media.getEncryptedFile id:long access_hash:long = EncryptedFile;

		// not has flags

		m.Id = dBuf.Long()
		m.AccessHash = dBuf.Long()
		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

func (m *TLMediaGetEncryptedFile) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

// TLMediaUploadWallPaperFile
///////////////////////////////////////////////////////////////////////////////

func (m *TLMediaUploadWallPaperFile) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	// x.Int(int32(CRC32_media_uploadWallPaperFile))

	switch uint32(m.Constructor) {
	case 0x9cfaadfe:
		// media.uploadWallPaperFile owner_id:long file:InputFile mime_type:string admin:Bool = Document;
		x.UInt(0x9cfaadfe)

		// no flags

		x.Long(m.GetOwnerId())
		x.Bytes(m.GetFile().Encode(layer))
		x.String(m.GetMimeType())
		x.Bytes(m.GetAdmin().Encode(layer))

	default:
		// log.Errorf("")
	}

	return x.GetBuf()
}

func (m *TLMediaUploadWallPaperFile) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLMediaUploadWallPaperFile) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0x9cfaadfe:
		// media.uploadWallPaperFile owner_id:long file:InputFile mime_type:string admin:Bool = Document;

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

func (m *TLMediaUploadWallPaperFile) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

// TLMediaUploadThemeFile
///////////////////////////////////////////////////////////////////////////////

func (m *TLMediaUploadThemeFile) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	// x.Int(int32(CRC32_media_uploadThemeFile))

	switch uint32(m.Constructor) {
	case 0x42e6b860:
		// media.uploadThemeFile flags:# owner_id:long file:InputFile thumb:flags.0?InputFile mime_type:string file_name:string = Document;
		x.UInt(0x42e6b860)

		// set flags
		var flags uint32 = 0

		if m.GetThumb() != nil {
			flags |= 1 << 0
		}

		x.UInt(flags)

		// flags Debug by @benqi
		x.Long(m.GetOwnerId())
		x.Bytes(m.GetFile().Encode(layer))
		if m.GetThumb() != nil {
			x.Bytes(m.GetThumb().Encode(layer))
		}

		x.String(m.GetMimeType())
		x.String(m.GetFileName())

	default:
		// log.Errorf("")
	}

	return x.GetBuf()
}

func (m *TLMediaUploadThemeFile) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLMediaUploadThemeFile) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0x42e6b860:
		// media.uploadThemeFile flags:# owner_id:long file:InputFile thumb:flags.0?InputFile mime_type:string file_name:string = Document;

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

func (m *TLMediaUploadThemeFile) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

// TLMediaUploadStickerFile
///////////////////////////////////////////////////////////////////////////////

func (m *TLMediaUploadStickerFile) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	// x.Int(int32(CRC32_media_uploadStickerFile))

	switch uint32(m.Constructor) {
	case 0xacb624ed:
		// media.uploadStickerFile flags:# owner_id:long file:InputFile thumb:flags.0?InputFile mime_type:string file_name:string document_attribute_sticker:DocumentAttribute = Document;
		x.UInt(0xacb624ed)

		// set flags
		var flags uint32 = 0

		if m.GetThumb() != nil {
			flags |= 1 << 0
		}

		x.UInt(flags)

		// flags Debug by @benqi
		x.Long(m.GetOwnerId())
		x.Bytes(m.GetFile().Encode(layer))
		if m.GetThumb() != nil {
			x.Bytes(m.GetThumb().Encode(layer))
		}

		x.String(m.GetMimeType())
		x.String(m.GetFileName())
		x.Bytes(m.GetDocumentAttributeSticker().Encode(layer))

	default:
		// log.Errorf("")
	}

	return x.GetBuf()
}

func (m *TLMediaUploadStickerFile) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLMediaUploadStickerFile) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0xacb624ed:
		// media.uploadStickerFile flags:# owner_id:long file:InputFile thumb:flags.0?InputFile mime_type:string file_name:string document_attribute_sticker:DocumentAttribute = Document;

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

func (m *TLMediaUploadStickerFile) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

//----------------------------------------------------------------------------------------------------------------
// Vector_PhotoSizeList
///////////////////////////////////////////////////////////////////////////////
func (m *Vector_PhotoSizeList) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	x.Int(int32(mtproto.CRC32_vector))
	x.Int(int32(len(m.Datas)))
	for _, v := range m.Datas {
		x.Bytes((*v).Encode(layer))
	}

	return x.GetBuf()
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

func (m *Vector_PhotoSizeList) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

// Vector_Document
///////////////////////////////////////////////////////////////////////////////
func (m *Vector_Document) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	x.Int(int32(mtproto.CRC32_vector))
	x.Int(int32(len(m.Datas)))
	for _, v := range m.Datas {
		x.Bytes((*v).Encode(layer))
	}

	return x.GetBuf()
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

func (m *Vector_Document) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}
