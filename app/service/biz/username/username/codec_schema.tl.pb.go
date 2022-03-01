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

package username

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
	-885195923: func() mtproto.TLObject { // 0xcb3cfb6d
		o := MakeTLUsernameNotExisted(nil)
		o.Data2.Constructor = -885195923
		return o
	},
	-1394084659: func() mtproto.TLObject { // 0xace7f4cd
		o := MakeTLUsernameExisted(nil)
		o.Data2.Constructor = -1394084659
		return o
	},
	-803256399: func() mtproto.TLObject { // 0xd01f47b1
		o := MakeTLUsernameExistedNotMe(nil)
		o.Data2.Constructor = -803256399
		return o
	},
	-2024900751: func() mtproto.TLObject { // 0x874e7771
		o := MakeTLUsernameExistedIsMe(nil)
		o.Data2.Constructor = -2024900751
		return o
	},
	-1438646081: func() mtproto.TLObject { // 0xaa4000bf
		o := MakeTLUsernameData(nil)
		o.Data2.Constructor = -1438646081
		return o
	},

	// Method
	154073301: func() mtproto.TLObject { // 0x92ef8d5
		return &TLUsernameGetAccountUsername{
			Constructor: 154073301,
		}
	},
	1240985861: func() mtproto.TLObject { // 0x49f7f105
		return &TLUsernameCheckAccountUsername{
			Constructor: 1240985861,
		}
	},
	-2038134827: func() mtproto.TLObject { // 0x868487d5
		return &TLUsernameGetChannelUsername{
			Constructor: -2038134827,
		}
	},
	651476637: func() mtproto.TLObject { // 0x26d4be9d
		return &TLUsernameCheckChannelUsername{
			Constructor: 651476637,
		}
	},
	1718205916: func() mtproto.TLObject { // 0x6669bddc
		return &TLUsernameUpdateUsernameByPeer{
			Constructor: 1718205916,
		}
	},
	684369621: func() mtproto.TLObject { // 0x28caa6d5
		return &TLUsernameCheckUsername{
			Constructor: 684369621,
		}
	},
	1389777971: func() mtproto.TLObject { // 0x52d65433
		return &TLUsernameUpdateUsername{
			Constructor: 1389777971,
		}
	},
	-1065913464: func() mtproto.TLObject { // 0xc0777388
		return &TLUsernameDeleteUsername{
			Constructor: -1065913464,
		}
	},
	2008689862: func() mtproto.TLObject { // 0x77ba2cc6
		return &TLUsernameResolveUsername{
			Constructor: 2008689862,
		}
	},
	1218942797: func() mtproto.TLObject { // 0x48a7974d
		return &TLUsernameGetListByUsernameList{
			Constructor: 1218942797,
		}
	},
	507822189: func() mtproto.TLObject { // 0x1e44c06d
		return &TLUsernameDeleteUsernameByPeer{
			Constructor: 507822189,
		}
	},
	-391798010: func() mtproto.TLObject { // 0xe8a5a306
		return &TLUsernameSearch{
			Constructor: -391798010,
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
// UsernameExist <--
//  + TL_UsernameNotExisted
//  + TL_UsernameExisted
//  + TL_UsernameExistedNotMe
//  + TL_UsernameExistedIsMe
//

func (m *UsernameExist) Encode(layer int32) []byte {
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
	case Predicate_usernameNotExisted:
		t := m.To_UsernameNotExisted()
		xBuf = t.Encode(layer)
	case Predicate_usernameExisted:
		t := m.To_UsernameExisted()
		xBuf = t.Encode(layer)
	case Predicate_usernameExistedNotMe:
		t := m.To_UsernameExistedNotMe()
		xBuf = t.Encode(layer)
	case Predicate_usernameExistedIsMe:
		t := m.To_UsernameExistedIsMe()
		xBuf = t.Encode(layer)

	default:
		// logx.Errorf("invalid predicate error: %s",  m.PredicateName)
		return []byte{}
	}

	return xBuf
}

func (m *UsernameExist) CalcByteSize(layer int32) int {
	return 0
}

func (m *UsernameExist) Decode(dBuf *mtproto.DecodeBuf) error {
	m.Constructor = TLConstructor(dBuf.Int())
	switch uint32(m.Constructor) {
	case 0xcb3cfb6d:
		m2 := MakeTLUsernameNotExisted(m)
		m2.Decode(dBuf)
	case 0xace7f4cd:
		m2 := MakeTLUsernameExisted(m)
		m2.Decode(dBuf)
	case 0xd01f47b1:
		m2 := MakeTLUsernameExistedNotMe(m)
		m2.Decode(dBuf)
	case 0x874e7771:
		m2 := MakeTLUsernameExistedIsMe(m)
		m2.Decode(dBuf)

	default:
		return fmt.Errorf("invalid constructorId: 0x%x", uint32(m.Constructor))
	}
	return dBuf.GetError()
}

func (m *UsernameExist) DebugString() string {
	switch m.PredicateName {
	case Predicate_usernameNotExisted:
		t := m.To_UsernameNotExisted()
		return t.DebugString()
	case Predicate_usernameExisted:
		t := m.To_UsernameExisted()
		return t.DebugString()
	case Predicate_usernameExistedNotMe:
		t := m.To_UsernameExistedNotMe()
		return t.DebugString()
	case Predicate_usernameExistedIsMe:
		t := m.To_UsernameExistedIsMe()
		return t.DebugString()

	default:
		return "{}"
	}
}

// To_UsernameNotExisted
// usernameNotExisted = UsernameExist;
func (m *UsernameExist) To_UsernameNotExisted() *TLUsernameNotExisted {
	m.PredicateName = Predicate_usernameNotExisted
	return &TLUsernameNotExisted{
		Data2: m,
	}
}

// To_UsernameExisted
// usernameExisted = UsernameExist;
func (m *UsernameExist) To_UsernameExisted() *TLUsernameExisted {
	m.PredicateName = Predicate_usernameExisted
	return &TLUsernameExisted{
		Data2: m,
	}
}

// To_UsernameExistedNotMe
// usernameExistedNotMe = UsernameExist;
func (m *UsernameExist) To_UsernameExistedNotMe() *TLUsernameExistedNotMe {
	m.PredicateName = Predicate_usernameExistedNotMe
	return &TLUsernameExistedNotMe{
		Data2: m,
	}
}

// To_UsernameExistedIsMe
// usernameExistedIsMe = UsernameExist;
func (m *UsernameExist) To_UsernameExistedIsMe() *TLUsernameExistedIsMe {
	m.PredicateName = Predicate_usernameExistedIsMe
	return &TLUsernameExistedIsMe{
		Data2: m,
	}
}

// MakeTLUsernameNotExisted
// usernameNotExisted = UsernameExist;
func MakeTLUsernameNotExisted(data2 *UsernameExist) *TLUsernameNotExisted {
	if data2 == nil {
		return &TLUsernameNotExisted{Data2: &UsernameExist{
			PredicateName: Predicate_usernameNotExisted,
		}}
	} else {
		data2.PredicateName = Predicate_usernameNotExisted
		return &TLUsernameNotExisted{Data2: data2}
	}
}

func (m *TLUsernameNotExisted) To_UsernameExist() *UsernameExist {
	m.Data2.PredicateName = Predicate_usernameNotExisted
	return m.Data2
}

func (m *TLUsernameNotExisted) GetPredicateName() string {
	return Predicate_usernameNotExisted
}

func (m *TLUsernameNotExisted) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)

	var encodeF = map[uint32]func() []byte{
		0xcb3cfb6d: func() []byte {
			// usernameNotExisted = UsernameExist;
			x.UInt(0xcb3cfb6d)

			return x.GetBuf()
		},
	}

	clazzId := GetClazzID(Predicate_usernameNotExisted, int(layer))
	if f, ok := encodeF[uint32(clazzId)]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		// log.Errorf("not found clazzId by (%s, %d)", Predicate_usernameNotExisted, layer)
		return x.GetBuf()
	}

	return x.GetBuf()
}

func (m *TLUsernameNotExisted) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLUsernameNotExisted) Decode(dBuf *mtproto.DecodeBuf) error {
	var decodeF = map[uint32]func() error{
		0xcb3cfb6d: func() error {
			// usernameNotExisted = UsernameExist;
			return dBuf.GetError()
		},
	}

	if f, ok := decodeF[uint32(m.Data2.Constructor)]; ok {
		return f()
	} else {
		return fmt.Errorf("invalid constructor: %x", uint32(m.Data2.Constructor))
	}
}

func (m *TLUsernameNotExisted) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

// MakeTLUsernameExisted
// usernameExisted = UsernameExist;
func MakeTLUsernameExisted(data2 *UsernameExist) *TLUsernameExisted {
	if data2 == nil {
		return &TLUsernameExisted{Data2: &UsernameExist{
			PredicateName: Predicate_usernameExisted,
		}}
	} else {
		data2.PredicateName = Predicate_usernameExisted
		return &TLUsernameExisted{Data2: data2}
	}
}

func (m *TLUsernameExisted) To_UsernameExist() *UsernameExist {
	m.Data2.PredicateName = Predicate_usernameExisted
	return m.Data2
}

func (m *TLUsernameExisted) GetPredicateName() string {
	return Predicate_usernameExisted
}

func (m *TLUsernameExisted) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)

	var encodeF = map[uint32]func() []byte{
		0xace7f4cd: func() []byte {
			// usernameExisted = UsernameExist;
			x.UInt(0xace7f4cd)

			return x.GetBuf()
		},
	}

	clazzId := GetClazzID(Predicate_usernameExisted, int(layer))
	if f, ok := encodeF[uint32(clazzId)]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		// log.Errorf("not found clazzId by (%s, %d)", Predicate_usernameExisted, layer)
		return x.GetBuf()
	}

	return x.GetBuf()
}

func (m *TLUsernameExisted) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLUsernameExisted) Decode(dBuf *mtproto.DecodeBuf) error {
	var decodeF = map[uint32]func() error{
		0xace7f4cd: func() error {
			// usernameExisted = UsernameExist;
			return dBuf.GetError()
		},
	}

	if f, ok := decodeF[uint32(m.Data2.Constructor)]; ok {
		return f()
	} else {
		return fmt.Errorf("invalid constructor: %x", uint32(m.Data2.Constructor))
	}
}

func (m *TLUsernameExisted) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

// MakeTLUsernameExistedNotMe
// usernameExistedNotMe = UsernameExist;
func MakeTLUsernameExistedNotMe(data2 *UsernameExist) *TLUsernameExistedNotMe {
	if data2 == nil {
		return &TLUsernameExistedNotMe{Data2: &UsernameExist{
			PredicateName: Predicate_usernameExistedNotMe,
		}}
	} else {
		data2.PredicateName = Predicate_usernameExistedNotMe
		return &TLUsernameExistedNotMe{Data2: data2}
	}
}

func (m *TLUsernameExistedNotMe) To_UsernameExist() *UsernameExist {
	m.Data2.PredicateName = Predicate_usernameExistedNotMe
	return m.Data2
}

func (m *TLUsernameExistedNotMe) GetPredicateName() string {
	return Predicate_usernameExistedNotMe
}

func (m *TLUsernameExistedNotMe) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)

	var encodeF = map[uint32]func() []byte{
		0xd01f47b1: func() []byte {
			// usernameExistedNotMe = UsernameExist;
			x.UInt(0xd01f47b1)

			return x.GetBuf()
		},
	}

	clazzId := GetClazzID(Predicate_usernameExistedNotMe, int(layer))
	if f, ok := encodeF[uint32(clazzId)]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		// log.Errorf("not found clazzId by (%s, %d)", Predicate_usernameExistedNotMe, layer)
		return x.GetBuf()
	}

	return x.GetBuf()
}

func (m *TLUsernameExistedNotMe) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLUsernameExistedNotMe) Decode(dBuf *mtproto.DecodeBuf) error {
	var decodeF = map[uint32]func() error{
		0xd01f47b1: func() error {
			// usernameExistedNotMe = UsernameExist;
			return dBuf.GetError()
		},
	}

	if f, ok := decodeF[uint32(m.Data2.Constructor)]; ok {
		return f()
	} else {
		return fmt.Errorf("invalid constructor: %x", uint32(m.Data2.Constructor))
	}
}

func (m *TLUsernameExistedNotMe) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

// MakeTLUsernameExistedIsMe
// usernameExistedIsMe = UsernameExist;
func MakeTLUsernameExistedIsMe(data2 *UsernameExist) *TLUsernameExistedIsMe {
	if data2 == nil {
		return &TLUsernameExistedIsMe{Data2: &UsernameExist{
			PredicateName: Predicate_usernameExistedIsMe,
		}}
	} else {
		data2.PredicateName = Predicate_usernameExistedIsMe
		return &TLUsernameExistedIsMe{Data2: data2}
	}
}

func (m *TLUsernameExistedIsMe) To_UsernameExist() *UsernameExist {
	m.Data2.PredicateName = Predicate_usernameExistedIsMe
	return m.Data2
}

func (m *TLUsernameExistedIsMe) GetPredicateName() string {
	return Predicate_usernameExistedIsMe
}

func (m *TLUsernameExistedIsMe) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)

	var encodeF = map[uint32]func() []byte{
		0x874e7771: func() []byte {
			// usernameExistedIsMe = UsernameExist;
			x.UInt(0x874e7771)

			return x.GetBuf()
		},
	}

	clazzId := GetClazzID(Predicate_usernameExistedIsMe, int(layer))
	if f, ok := encodeF[uint32(clazzId)]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		// log.Errorf("not found clazzId by (%s, %d)", Predicate_usernameExistedIsMe, layer)
		return x.GetBuf()
	}

	return x.GetBuf()
}

func (m *TLUsernameExistedIsMe) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLUsernameExistedIsMe) Decode(dBuf *mtproto.DecodeBuf) error {
	var decodeF = map[uint32]func() error{
		0x874e7771: func() error {
			// usernameExistedIsMe = UsernameExist;
			return dBuf.GetError()
		},
	}

	if f, ok := decodeF[uint32(m.Data2.Constructor)]; ok {
		return f()
	} else {
		return fmt.Errorf("invalid constructor: %x", uint32(m.Data2.Constructor))
	}
}

func (m *TLUsernameExistedIsMe) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

///////////////////////////////////////////////////////////////////////////////
// UsernameData <--
//  + TL_UsernameData
//

func (m *UsernameData) Encode(layer int32) []byte {
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
	case Predicate_usernameData:
		t := m.To_UsernameData()
		xBuf = t.Encode(layer)

	default:
		// logx.Errorf("invalid predicate error: %s",  m.PredicateName)
		return []byte{}
	}

	return xBuf
}

func (m *UsernameData) CalcByteSize(layer int32) int {
	return 0
}

func (m *UsernameData) Decode(dBuf *mtproto.DecodeBuf) error {
	m.Constructor = TLConstructor(dBuf.Int())
	switch uint32(m.Constructor) {
	case 0xaa4000bf:
		m2 := MakeTLUsernameData(m)
		m2.Decode(dBuf)

	default:
		return fmt.Errorf("invalid constructorId: 0x%x", uint32(m.Constructor))
	}
	return dBuf.GetError()
}

func (m *UsernameData) DebugString() string {
	switch m.PredicateName {
	case Predicate_usernameData:
		t := m.To_UsernameData()
		return t.DebugString()

	default:
		return "{}"
	}
}

// To_UsernameData
// usernameData flags:# username:string peer:flags.0?Peer = UsernameData;
func (m *UsernameData) To_UsernameData() *TLUsernameData {
	m.PredicateName = Predicate_usernameData
	return &TLUsernameData{
		Data2: m,
	}
}

// MakeTLUsernameData
// usernameData flags:# username:string peer:flags.0?Peer = UsernameData;
func MakeTLUsernameData(data2 *UsernameData) *TLUsernameData {
	if data2 == nil {
		return &TLUsernameData{Data2: &UsernameData{
			PredicateName: Predicate_usernameData,
		}}
	} else {
		data2.PredicateName = Predicate_usernameData
		return &TLUsernameData{Data2: data2}
	}
}

func (m *TLUsernameData) To_UsernameData() *UsernameData {
	m.Data2.PredicateName = Predicate_usernameData
	return m.Data2
}

//// flags
func (m *TLUsernameData) SetUsername(v string) { m.Data2.Username = v }
func (m *TLUsernameData) GetUsername() string  { return m.Data2.Username }

func (m *TLUsernameData) SetPeer(v *mtproto.Peer) { m.Data2.Peer = v }
func (m *TLUsernameData) GetPeer() *mtproto.Peer  { return m.Data2.Peer }

func (m *TLUsernameData) GetPredicateName() string {
	return Predicate_usernameData
}

func (m *TLUsernameData) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)

	var encodeF = map[uint32]func() []byte{
		0xaa4000bf: func() []byte {
			// usernameData flags:# username:string peer:flags.0?Peer = UsernameData;
			x.UInt(0xaa4000bf)

			// set flags
			var getFlags = func() uint32 {
				var flags uint32 = 0

				if m.GetPeer() != nil {
					flags |= 1 << 0
				}

				return flags
			}

			// set flags
			var flags = getFlags()
			x.UInt(flags)
			x.String(m.GetUsername())
			if m.GetPeer() != nil {
				x.Bytes(m.GetPeer().Encode(layer))
			}

			return x.GetBuf()
		},
	}

	clazzId := GetClazzID(Predicate_usernameData, int(layer))
	if f, ok := encodeF[uint32(clazzId)]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		// log.Errorf("not found clazzId by (%s, %d)", Predicate_usernameData, layer)
		return x.GetBuf()
	}

	return x.GetBuf()
}

func (m *TLUsernameData) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLUsernameData) Decode(dBuf *mtproto.DecodeBuf) error {
	var decodeF = map[uint32]func() error{
		0xaa4000bf: func() error {
			// usernameData flags:# username:string peer:flags.0?Peer = UsernameData;
			var flags = dBuf.UInt()
			_ = flags
			m.SetUsername(dBuf.String())
			if (flags & (1 << 0)) != 0 {
				m2 := &mtproto.Peer{}
				m2.Decode(dBuf)
				m.SetPeer(m2)
			}
			return dBuf.GetError()
		},
	}

	if f, ok := decodeF[uint32(m.Data2.Constructor)]; ok {
		return f()
	} else {
		return fmt.Errorf("invalid constructor: %x", uint32(m.Data2.Constructor))
	}
}

func (m *TLUsernameData) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

//----------------------------------------------------------------------------------------------------------------
// TLUsernameGetAccountUsername
///////////////////////////////////////////////////////////////////////////////

func (m *TLUsernameGetAccountUsername) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	// x.Int(int32(CRC32_username_getAccountUsername))

	switch uint32(m.Constructor) {
	case 0x92ef8d5:
		// username.getAccountUsername user_id:long = UsernameData;
		x.UInt(0x92ef8d5)

		// no flags

		x.Long(m.GetUserId())

	default:
		// log.Errorf("")
	}

	return x.GetBuf()
}

func (m *TLUsernameGetAccountUsername) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLUsernameGetAccountUsername) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0x92ef8d5:
		// username.getAccountUsername user_id:long = UsernameData;

		// not has flags

		m.UserId = dBuf.Long()
		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

func (m *TLUsernameGetAccountUsername) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

// TLUsernameCheckAccountUsername
///////////////////////////////////////////////////////////////////////////////

func (m *TLUsernameCheckAccountUsername) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	// x.Int(int32(CRC32_username_checkAccountUsername))

	switch uint32(m.Constructor) {
	case 0x49f7f105:
		// username.checkAccountUsername user_id:long username:string = UsernameExist;
		x.UInt(0x49f7f105)

		// no flags

		x.Long(m.GetUserId())
		x.String(m.GetUsername())

	default:
		// log.Errorf("")
	}

	return x.GetBuf()
}

func (m *TLUsernameCheckAccountUsername) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLUsernameCheckAccountUsername) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0x49f7f105:
		// username.checkAccountUsername user_id:long username:string = UsernameExist;

		// not has flags

		m.UserId = dBuf.Long()
		m.Username = dBuf.String()
		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

func (m *TLUsernameCheckAccountUsername) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

// TLUsernameGetChannelUsername
///////////////////////////////////////////////////////////////////////////////

func (m *TLUsernameGetChannelUsername) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	// x.Int(int32(CRC32_username_getChannelUsername))

	switch uint32(m.Constructor) {
	case 0x868487d5:
		// username.getChannelUsername channel_id:long = UsernameData;
		x.UInt(0x868487d5)

		// no flags

		x.Long(m.GetChannelId())

	default:
		// log.Errorf("")
	}

	return x.GetBuf()
}

func (m *TLUsernameGetChannelUsername) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLUsernameGetChannelUsername) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0x868487d5:
		// username.getChannelUsername channel_id:long = UsernameData;

		// not has flags

		m.ChannelId = dBuf.Long()
		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

func (m *TLUsernameGetChannelUsername) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

// TLUsernameCheckChannelUsername
///////////////////////////////////////////////////////////////////////////////

func (m *TLUsernameCheckChannelUsername) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	// x.Int(int32(CRC32_username_checkChannelUsername))

	switch uint32(m.Constructor) {
	case 0x26d4be9d:
		// username.checkChannelUsername channel_id:long username:string = UsernameExist;
		x.UInt(0x26d4be9d)

		// no flags

		x.Long(m.GetChannelId())
		x.String(m.GetUsername())

	default:
		// log.Errorf("")
	}

	return x.GetBuf()
}

func (m *TLUsernameCheckChannelUsername) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLUsernameCheckChannelUsername) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0x26d4be9d:
		// username.checkChannelUsername channel_id:long username:string = UsernameExist;

		// not has flags

		m.ChannelId = dBuf.Long()
		m.Username = dBuf.String()
		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

func (m *TLUsernameCheckChannelUsername) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

// TLUsernameUpdateUsernameByPeer
///////////////////////////////////////////////////////////////////////////////

func (m *TLUsernameUpdateUsernameByPeer) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	// x.Int(int32(CRC32_username_updateUsernameByPeer))

	switch uint32(m.Constructor) {
	case 0x6669bddc:
		// username.updateUsernameByPeer peer_type:int peer_id:long username:string = Bool;
		x.UInt(0x6669bddc)

		// no flags

		x.Int(m.GetPeerType())
		x.Long(m.GetPeerId())
		x.String(m.GetUsername())

	default:
		// log.Errorf("")
	}

	return x.GetBuf()
}

func (m *TLUsernameUpdateUsernameByPeer) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLUsernameUpdateUsernameByPeer) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0x6669bddc:
		// username.updateUsernameByPeer peer_type:int peer_id:long username:string = Bool;

		// not has flags

		m.PeerType = dBuf.Int()
		m.PeerId = dBuf.Long()
		m.Username = dBuf.String()
		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

func (m *TLUsernameUpdateUsernameByPeer) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

// TLUsernameCheckUsername
///////////////////////////////////////////////////////////////////////////////

func (m *TLUsernameCheckUsername) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	// x.Int(int32(CRC32_username_checkUsername))

	switch uint32(m.Constructor) {
	case 0x28caa6d5:
		// username.checkUsername username:string = UsernameExist;
		x.UInt(0x28caa6d5)

		// no flags

		x.String(m.GetUsername())

	default:
		// log.Errorf("")
	}

	return x.GetBuf()
}

func (m *TLUsernameCheckUsername) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLUsernameCheckUsername) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0x28caa6d5:
		// username.checkUsername username:string = UsernameExist;

		// not has flags

		m.Username = dBuf.String()
		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

func (m *TLUsernameCheckUsername) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

// TLUsernameUpdateUsername
///////////////////////////////////////////////////////////////////////////////

func (m *TLUsernameUpdateUsername) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	// x.Int(int32(CRC32_username_updateUsername))

	switch uint32(m.Constructor) {
	case 0x52d65433:
		// username.updateUsername peer_type:int peer_id:long username:string = Bool;
		x.UInt(0x52d65433)

		// no flags

		x.Int(m.GetPeerType())
		x.Long(m.GetPeerId())
		x.String(m.GetUsername())

	default:
		// log.Errorf("")
	}

	return x.GetBuf()
}

func (m *TLUsernameUpdateUsername) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLUsernameUpdateUsername) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0x52d65433:
		// username.updateUsername peer_type:int peer_id:long username:string = Bool;

		// not has flags

		m.PeerType = dBuf.Int()
		m.PeerId = dBuf.Long()
		m.Username = dBuf.String()
		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

func (m *TLUsernameUpdateUsername) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

// TLUsernameDeleteUsername
///////////////////////////////////////////////////////////////////////////////

func (m *TLUsernameDeleteUsername) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	// x.Int(int32(CRC32_username_deleteUsername))

	switch uint32(m.Constructor) {
	case 0xc0777388:
		// username.deleteUsername username:string = Bool;
		x.UInt(0xc0777388)

		// no flags

		x.String(m.GetUsername())

	default:
		// log.Errorf("")
	}

	return x.GetBuf()
}

func (m *TLUsernameDeleteUsername) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLUsernameDeleteUsername) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0xc0777388:
		// username.deleteUsername username:string = Bool;

		// not has flags

		m.Username = dBuf.String()
		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

func (m *TLUsernameDeleteUsername) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

// TLUsernameResolveUsername
///////////////////////////////////////////////////////////////////////////////

func (m *TLUsernameResolveUsername) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	// x.Int(int32(CRC32_username_resolveUsername))

	switch uint32(m.Constructor) {
	case 0x77ba2cc6:
		// username.resolveUsername username:string = Peer;
		x.UInt(0x77ba2cc6)

		// no flags

		x.String(m.GetUsername())

	default:
		// log.Errorf("")
	}

	return x.GetBuf()
}

func (m *TLUsernameResolveUsername) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLUsernameResolveUsername) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0x77ba2cc6:
		// username.resolveUsername username:string = Peer;

		// not has flags

		m.Username = dBuf.String()
		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

func (m *TLUsernameResolveUsername) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

// TLUsernameGetListByUsernameList
///////////////////////////////////////////////////////////////////////////////

func (m *TLUsernameGetListByUsernameList) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	// x.Int(int32(CRC32_username_getListByUsernameList))

	switch uint32(m.Constructor) {
	case 0x48a7974d:
		// username.getListByUsernameList names:Vector<string> = Vector<UsernameData>;
		x.UInt(0x48a7974d)

		// no flags

		x.VectorString(m.GetNames())

	default:
		// log.Errorf("")
	}

	return x.GetBuf()
}

func (m *TLUsernameGetListByUsernameList) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLUsernameGetListByUsernameList) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0x48a7974d:
		// username.getListByUsernameList names:Vector<string> = Vector<UsernameData>;

		// not has flags

		m.Names = dBuf.VectorString()

		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

func (m *TLUsernameGetListByUsernameList) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

// TLUsernameDeleteUsernameByPeer
///////////////////////////////////////////////////////////////////////////////

func (m *TLUsernameDeleteUsernameByPeer) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	// x.Int(int32(CRC32_username_deleteUsernameByPeer))

	switch uint32(m.Constructor) {
	case 0x1e44c06d:
		// username.deleteUsernameByPeer peer_type:int peer_id:long = Bool;
		x.UInt(0x1e44c06d)

		// no flags

		x.Int(m.GetPeerType())
		x.Long(m.GetPeerId())

	default:
		// log.Errorf("")
	}

	return x.GetBuf()
}

func (m *TLUsernameDeleteUsernameByPeer) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLUsernameDeleteUsernameByPeer) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0x1e44c06d:
		// username.deleteUsernameByPeer peer_type:int peer_id:long = Bool;

		// not has flags

		m.PeerType = dBuf.Int()
		m.PeerId = dBuf.Long()
		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

func (m *TLUsernameDeleteUsernameByPeer) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

// TLUsernameSearch
///////////////////////////////////////////////////////////////////////////////

func (m *TLUsernameSearch) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	// x.Int(int32(CRC32_username_search))

	switch uint32(m.Constructor) {
	case 0xe8a5a306:
		// username.search q:string excluded_contacts:Vector<long> limit:int = Vector<UsernameData>;
		x.UInt(0xe8a5a306)

		// no flags

		x.String(m.GetQ())

		x.VectorLong(m.GetExcludedContacts())

		x.Int(m.GetLimit())

	default:
		// log.Errorf("")
	}

	return x.GetBuf()
}

func (m *TLUsernameSearch) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLUsernameSearch) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0xe8a5a306:
		// username.search q:string excluded_contacts:Vector<long> limit:int = Vector<UsernameData>;

		// not has flags

		m.Q = dBuf.String()

		m.ExcludedContacts = dBuf.VectorLong()

		m.Limit = dBuf.Int()
		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

func (m *TLUsernameSearch) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

//----------------------------------------------------------------------------------------------------------------
// Vector_UsernameData
///////////////////////////////////////////////////////////////////////////////
func (m *Vector_UsernameData) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	x.Int(int32(mtproto.CRC32_vector))
	x.Int(int32(len(m.Datas)))
	for _, v := range m.Datas {
		x.Bytes((*v).Encode(layer))
	}

	return x.GetBuf()
}

func (m *Vector_UsernameData) Decode(dBuf *mtproto.DecodeBuf) error {
	dBuf.Int() // TODO(@benqi): Check crc32 invalid
	l1 := dBuf.Int()
	m.Datas = make([]*UsernameData, l1)
	for i := int32(0); i < l1; i++ {
		m.Datas[i] = new(UsernameData)
		(*m.Datas[i]).Decode(dBuf)
	}

	return dBuf.GetError()
}

func (m *Vector_UsernameData) CalcByteSize(layer int32) int {
	return 0
}

func (m *Vector_UsernameData) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}
