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

package gif

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
	556825867: func() mtproto.TLObject { // 0x21307d0b
		return &TLGifSaveGif{
			Constructor: 556825867,
		}
	},
	926787430: func() mtproto.TLObject { // 0x373da766
		return &TLGifGetSavedGifs{
			Constructor: 926787430,
		}
	},
	523645139: func() mtproto.TLObject { // 0x1f3630d3
		return &TLGifDeleteSavedGif{
			Constructor: 523645139,
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
// TLGifSaveGif
///////////////////////////////////////////////////////////////////////////////

func (m *TLGifSaveGif) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	// x.Int(int32(CRC32_gif_saveGif))

	switch uint32(m.Constructor) {
	case 0x21307d0b:
		// gif.saveGif user_id:long gif_id:long = Bool;
		x.UInt(0x21307d0b)

		// no flags

		x.Long(m.GetUserId())
		x.Long(m.GetGifId())

	default:
		// log.Errorf("")
	}

	return x.GetBuf()
}

func (m *TLGifSaveGif) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLGifSaveGif) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0x21307d0b:
		// gif.saveGif user_id:long gif_id:long = Bool;

		// not has flags

		m.UserId = dBuf.Long()
		m.GifId = dBuf.Long()
		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

func (m *TLGifSaveGif) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

// TLGifGetSavedGifs
///////////////////////////////////////////////////////////////////////////////

func (m *TLGifGetSavedGifs) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	// x.Int(int32(CRC32_gif_getSavedGifs))

	switch uint32(m.Constructor) {
	case 0x373da766:
		// gif.getSavedGifs user_id:long = Vector<long>;
		x.UInt(0x373da766)

		// no flags

		x.Long(m.GetUserId())

	default:
		// log.Errorf("")
	}

	return x.GetBuf()
}

func (m *TLGifGetSavedGifs) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLGifGetSavedGifs) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0x373da766:
		// gif.getSavedGifs user_id:long = Vector<long>;

		// not has flags

		m.UserId = dBuf.Long()
		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

func (m *TLGifGetSavedGifs) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

// TLGifDeleteSavedGif
///////////////////////////////////////////////////////////////////////////////

func (m *TLGifDeleteSavedGif) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	// x.Int(int32(CRC32_gif_deleteSavedGif))

	switch uint32(m.Constructor) {
	case 0x1f3630d3:
		// gif.deleteSavedGif user_id:long gif_id:long = Bool;
		x.UInt(0x1f3630d3)

		// no flags

		x.Long(m.GetUserId())
		x.Long(m.GetGifId())

	default:
		// log.Errorf("")
	}

	return x.GetBuf()
}

func (m *TLGifDeleteSavedGif) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLGifDeleteSavedGif) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0x1f3630d3:
		// gif.deleteSavedGif user_id:long gif_id:long = Bool;

		// not has flags

		m.UserId = dBuf.Long()
		m.GifId = dBuf.Long()
		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

func (m *TLGifDeleteSavedGif) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

//----------------------------------------------------------------------------------------------------------------
// Vector_Long
///////////////////////////////////////////////////////////////////////////////
func (m *Vector_Long) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	x.VectorLong(m.Datas)

	return x.GetBuf()
}

func (m *Vector_Long) Decode(dBuf *mtproto.DecodeBuf) error {
	m.Datas = dBuf.VectorLong()

	return dBuf.GetError()
}

func (m *Vector_Long) CalcByteSize(layer int32) int {
	return 0
}

func (m *Vector_Long) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}
