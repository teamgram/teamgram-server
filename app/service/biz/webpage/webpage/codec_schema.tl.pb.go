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

package webpage

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

	// Method
	1074946247: func() mtproto.TLObject { // 0x401260c7
		return &TLWebpageGetPendingWebPagePreview{
			Constructor: 1074946247,
		}
	},
	-2059356164: func() mtproto.TLObject { // 0x8540b7fc
		return &TLWebpageGetWebPagePreview{
			Constructor: -2059356164,
		}
	},
	-142855528: func() mtproto.TLObject { // 0xf77c3298
		return &TLWebpageGetWebPage{
			Constructor: -142855528,
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

//----------------------------------------------------------------------------------------------------------------
// TLWebpageGetPendingWebPagePreview
///////////////////////////////////////////////////////////////////////////////

func (m *TLWebpageGetPendingWebPagePreview) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	// x.Int(int32(CRC32_webpage_getPendingWebPagePreview))

	switch uint32(m.Constructor) {
	case 0x401260c7:
		// webpage.getPendingWebPagePreview flags:# message:string entities:flags.3?Vector<MessageEntity> = WebPage;
		x.UInt(0x401260c7)

		// set flags
		var flags uint32 = 0

		if m.GetEntities() != nil {
			flags |= 1 << 3
		}

		x.UInt(flags)

		// flags Debug by @benqi
		x.String(m.GetMessage())
		if m.GetEntities() != nil {
			x.Int(int32(mtproto.CRC32_vector))
			x.Int(int32(len(m.GetEntities())))
			for _, v := range m.GetEntities() {
				x.Bytes((*v).Encode(layer))
			}
		}

	default:
		// log.Errorf("")
	}

	return x.GetBuf()
}

func (m *TLWebpageGetPendingWebPagePreview) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLWebpageGetPendingWebPagePreview) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0x401260c7:
		// webpage.getPendingWebPagePreview flags:# message:string entities:flags.3?Vector<MessageEntity> = WebPage;

		flags := dBuf.UInt()
		_ = flags

		// flags Debug by @benqi
		m.Message = dBuf.String()
		if (flags & (1 << 3)) != 0 {
			c3 := dBuf.Int()
			if c3 != int32(mtproto.CRC32_vector) {
				// dBuf.err = fmt.Errorf("invalid mtproto.CRC32_vector, c%d: %d", 3, c3)
				return fmt.Errorf("invalid mtproto.CRC32_vector, c%d: %d", 3, c3)
			}
			l3 := dBuf.Int()
			v3 := make([]*mtproto.MessageEntity, l3)
			for i := int32(0); i < l3; i++ {
				v3[i] = &mtproto.MessageEntity{}
				v3[i].Decode(dBuf)
			}
			m.Entities = v3
		}
		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

func (m *TLWebpageGetPendingWebPagePreview) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

// TLWebpageGetWebPagePreview
///////////////////////////////////////////////////////////////////////////////

func (m *TLWebpageGetWebPagePreview) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	// x.Int(int32(CRC32_webpage_getWebPagePreview))

	switch uint32(m.Constructor) {
	case 0x8540b7fc:
		// webpage.getWebPagePreview flags:# message:string entities:flags.3?Vector<MessageEntity> = WebPage;
		x.UInt(0x8540b7fc)

		// set flags
		var flags uint32 = 0

		if m.GetEntities() != nil {
			flags |= 1 << 3
		}

		x.UInt(flags)

		// flags Debug by @benqi
		x.String(m.GetMessage())
		if m.GetEntities() != nil {
			x.Int(int32(mtproto.CRC32_vector))
			x.Int(int32(len(m.GetEntities())))
			for _, v := range m.GetEntities() {
				x.Bytes((*v).Encode(layer))
			}
		}

	default:
		// log.Errorf("")
	}

	return x.GetBuf()
}

func (m *TLWebpageGetWebPagePreview) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLWebpageGetWebPagePreview) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0x8540b7fc:
		// webpage.getWebPagePreview flags:# message:string entities:flags.3?Vector<MessageEntity> = WebPage;

		flags := dBuf.UInt()
		_ = flags

		// flags Debug by @benqi
		m.Message = dBuf.String()
		if (flags & (1 << 3)) != 0 {
			c3 := dBuf.Int()
			if c3 != int32(mtproto.CRC32_vector) {
				// dBuf.err = fmt.Errorf("invalid mtproto.CRC32_vector, c%d: %d", 3, c3)
				return fmt.Errorf("invalid mtproto.CRC32_vector, c%d: %d", 3, c3)
			}
			l3 := dBuf.Int()
			v3 := make([]*mtproto.MessageEntity, l3)
			for i := int32(0); i < l3; i++ {
				v3[i] = &mtproto.MessageEntity{}
				v3[i].Decode(dBuf)
			}
			m.Entities = v3
		}
		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

func (m *TLWebpageGetWebPagePreview) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

// TLWebpageGetWebPage
///////////////////////////////////////////////////////////////////////////////

func (m *TLWebpageGetWebPage) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	// x.Int(int32(CRC32_webpage_getWebPage))

	switch uint32(m.Constructor) {
	case 0xf77c3298:
		// webpage.getWebPage url:string hash:int = WebPage;
		x.UInt(0xf77c3298)

		// no flags

		x.String(m.GetUrl())
		x.Int(m.GetHash())

	default:
		// log.Errorf("")
	}

	return x.GetBuf()
}

func (m *TLWebpageGetWebPage) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLWebpageGetWebPage) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0xf77c3298:
		// webpage.getWebPage url:string hash:int = WebPage;

		// not has flags

		m.Url = dBuf.String()
		m.Hash = dBuf.Int()
		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

func (m *TLWebpageGetWebPage) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

//----------------------------------------------------------------------------------------------------------------
