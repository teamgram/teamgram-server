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

package report

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
	1976979630: func() mtproto.TLObject { // 0x75d650ae
		return &TLReportAccountReportPeer{
			Constructor: 1976979630,
		}
	},
	-1206920954: func() mtproto.TLObject { // 0xb80fd906
		return &TLReportAccountReportProfilePhoto{
			Constructor: -1206920954,
		}
	},
	-2120170998: func() mtproto.TLObject { // 0x81a0c20a
		return &TLReportMessagesReportSpam{
			Constructor: -2120170998,
		}
	},
	-1299590501: func() mtproto.TLObject { // 0xb289d29b
		return &TLReportMessagesReport{
			Constructor: -1299590501,
		}
	},
	762034535: func() mtproto.TLObject { // 0x2d6bb967
		return &TLReportMessagesReportEncryptedSpam{
			Constructor: 762034535,
		}
	},
	2010319160: func() mtproto.TLObject { // 0x77d30938
		return &TLReportChannelsReportSpam{
			Constructor: 2010319160,
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
// TLReportAccountReportPeer
///////////////////////////////////////////////////////////////////////////////
func (m *TLReportAccountReportPeer) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	// x.Int(int32(CRC32_report_accountReportPeer))

	switch uint32(m.Constructor) {
	case 0x75d650ae:
		// report.accountReportPeer reporter:long peer_type:int peer_id:long reason:ReportReason message:string = Bool;
		x.UInt(0x75d650ae)

		// no flags

		x.Long(m.GetReporter())
		x.Int(m.GetPeerType())
		x.Long(m.GetPeerId())
		x.Bytes(m.GetReason().Encode(layer))
		x.String(m.GetMessage())

	default:
		// log.Errorf("")
	}

	return x.GetBuf()
}

func (m *TLReportAccountReportPeer) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLReportAccountReportPeer) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0x75d650ae:
		// report.accountReportPeer reporter:long peer_type:int peer_id:long reason:ReportReason message:string = Bool;

		// not has flags

		m.Reporter = dBuf.Long()
		m.PeerType = dBuf.Int()
		m.PeerId = dBuf.Long()

		m4 := &mtproto.ReportReason{}
		m4.Decode(dBuf)
		m.Reason = m4

		m.Message = dBuf.String()
		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

func (m *TLReportAccountReportPeer) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

// TLReportAccountReportProfilePhoto
///////////////////////////////////////////////////////////////////////////////
func (m *TLReportAccountReportProfilePhoto) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	// x.Int(int32(CRC32_report_accountReportProfilePhoto))

	switch uint32(m.Constructor) {
	case 0xb80fd906:
		// report.accountReportProfilePhoto reporter:long peer_type:int peer_id:long photo_id:long reason:ReportReason message:string = Bool;
		x.UInt(0xb80fd906)

		// no flags

		x.Long(m.GetReporter())
		x.Int(m.GetPeerType())
		x.Long(m.GetPeerId())
		x.Long(m.GetPhotoId())
		x.Bytes(m.GetReason().Encode(layer))
		x.String(m.GetMessage())

	default:
		// log.Errorf("")
	}

	return x.GetBuf()
}

func (m *TLReportAccountReportProfilePhoto) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLReportAccountReportProfilePhoto) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0xb80fd906:
		// report.accountReportProfilePhoto reporter:long peer_type:int peer_id:long photo_id:long reason:ReportReason message:string = Bool;

		// not has flags

		m.Reporter = dBuf.Long()
		m.PeerType = dBuf.Int()
		m.PeerId = dBuf.Long()
		m.PhotoId = dBuf.Long()

		m5 := &mtproto.ReportReason{}
		m5.Decode(dBuf)
		m.Reason = m5

		m.Message = dBuf.String()
		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

func (m *TLReportAccountReportProfilePhoto) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

// TLReportMessagesReportSpam
///////////////////////////////////////////////////////////////////////////////
func (m *TLReportMessagesReportSpam) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	// x.Int(int32(CRC32_report_messagesReportSpam))

	switch uint32(m.Constructor) {
	case 0x81a0c20a:
		// report.messagesReportSpam reporter:long peer_type:int peer_id:long = Bool;
		x.UInt(0x81a0c20a)

		// no flags

		x.Long(m.GetReporter())
		x.Int(m.GetPeerType())
		x.Long(m.GetPeerId())

	default:
		// log.Errorf("")
	}

	return x.GetBuf()
}

func (m *TLReportMessagesReportSpam) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLReportMessagesReportSpam) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0x81a0c20a:
		// report.messagesReportSpam reporter:long peer_type:int peer_id:long = Bool;

		// not has flags

		m.Reporter = dBuf.Long()
		m.PeerType = dBuf.Int()
		m.PeerId = dBuf.Long()
		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

func (m *TLReportMessagesReportSpam) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

// TLReportMessagesReport
///////////////////////////////////////////////////////////////////////////////
func (m *TLReportMessagesReport) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	// x.Int(int32(CRC32_report_messagesReport))

	switch uint32(m.Constructor) {
	case 0xb289d29b:
		// report.messagesReport reporter:long peer_type:int peer_id:long id:Vector<int> reason:ReportReason message:string = Bool;
		x.UInt(0xb289d29b)

		// no flags

		x.Long(m.GetReporter())
		x.Int(m.GetPeerType())
		x.Long(m.GetPeerId())

		x.VectorInt(m.GetId())

		x.Bytes(m.GetReason().Encode(layer))
		x.String(m.GetMessage())

	default:
		// log.Errorf("")
	}

	return x.GetBuf()
}

func (m *TLReportMessagesReport) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLReportMessagesReport) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0xb289d29b:
		// report.messagesReport reporter:long peer_type:int peer_id:long id:Vector<int> reason:ReportReason message:string = Bool;

		// not has flags

		m.Reporter = dBuf.Long()
		m.PeerType = dBuf.Int()
		m.PeerId = dBuf.Long()

		m.Id = dBuf.VectorInt()

		m5 := &mtproto.ReportReason{}
		m5.Decode(dBuf)
		m.Reason = m5

		m.Message = dBuf.String()
		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

func (m *TLReportMessagesReport) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

// TLReportMessagesReportEncryptedSpam
///////////////////////////////////////////////////////////////////////////////
func (m *TLReportMessagesReportEncryptedSpam) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	// x.Int(int32(CRC32_report_messagesReportEncryptedSpam))

	switch uint32(m.Constructor) {
	case 0x2d6bb967:
		// report.messagesReportEncryptedSpam reporter:long chat_id:int = Bool;
		x.UInt(0x2d6bb967)

		// no flags

		x.Long(m.GetReporter())
		x.Int(m.GetChatId())

	default:
		// log.Errorf("")
	}

	return x.GetBuf()
}

func (m *TLReportMessagesReportEncryptedSpam) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLReportMessagesReportEncryptedSpam) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0x2d6bb967:
		// report.messagesReportEncryptedSpam reporter:long chat_id:int = Bool;

		// not has flags

		m.Reporter = dBuf.Long()
		m.ChatId = dBuf.Int()
		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

func (m *TLReportMessagesReportEncryptedSpam) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

// TLReportChannelsReportSpam
///////////////////////////////////////////////////////////////////////////////
func (m *TLReportChannelsReportSpam) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	// x.Int(int32(CRC32_report_channelsReportSpam))

	switch uint32(m.Constructor) {
	case 0x77d30938:
		// report.channelsReportSpam reporter:long channel_id:long user_id:long id:Vector<int> = Bool;
		x.UInt(0x77d30938)

		// no flags

		x.Long(m.GetReporter())
		x.Long(m.GetChannelId())
		x.Long(m.GetUserId())

		x.VectorInt(m.GetId())

	default:
		// log.Errorf("")
	}

	return x.GetBuf()
}

func (m *TLReportChannelsReportSpam) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLReportChannelsReportSpam) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0x77d30938:
		// report.channelsReportSpam reporter:long channel_id:long user_id:long id:Vector<int> = Bool;

		// not has flags

		m.Reporter = dBuf.Long()
		m.ChannelId = dBuf.Long()
		m.UserId = dBuf.Long()

		m.Id = dBuf.VectorInt()

		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

func (m *TLReportChannelsReportSpam) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

//----------------------------------------------------------------------------------------------------------------
