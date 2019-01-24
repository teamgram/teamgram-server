// Copyright (c) 2018-present,  NebulaChat Studio (https://nebula.chat).
//  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Author: Benqi (wubenqi@gmail.com)

package mtproto

import (
	"bytes"
	"compress/gzip"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"github.com/golang/glog"
	"io"
)

//const (
//	TLConstructor_CRC32_message2  		= 0x5bb8e511
//	TLConstructor_CRC32_msg_container  	= 0x73f1f8dc
//	TLConstructor_CRC32_msg_copy  		= 0xe06046b2
//	TLConstructor_CRC32_gzip_packed 	= 0x3072cfa1
//)

//message2 msg_id:long seqno:int bytes:int body:Object = Message2; // parsed manually
type TLMessageRawData struct {
	MsgId   int64
	Seqno   int32
	Bytes   int32
	ClassId int32
	Body    []byte
}

func (m *TLMessageRawData) String() string {
	return fmt.Sprintf("{message2#5bb8e511 - msg_id: %d, seq_no: %d, bytes: %d, class_id: %d}",
		m.MsgId,
		m.Seqno,
		m.Bytes,
		m.ClassId)
}

func (m *TLMessageRawData) Encode() []byte {
	x := NewEncodeBuf(512)
	// x.Int(int32(TLConstructor_CRC32_message2))
	x.Long(m.MsgId)
	x.Int(m.Seqno)
	x.Int(m.Bytes)
	x.Bytes(m.Body)
	return x.buf
}

func (m *TLMessageRawData) EncodeToLayer(layer int) []byte {
	x := NewEncodeBuf(512)
	// x.Int(int32(TLConstructor_CRC32_message2))
	x.Long(m.MsgId)
	x.Int(m.Seqno)
	x.Int(m.Bytes)
	x.Bytes(m.Body)
	return x.buf
}

func (m *TLMessageRawData) Decode(dbuf *DecodeBuf) error {
	m.MsgId = dbuf.Long()
	m.Seqno = dbuf.Int()
	m.Bytes = dbuf.Int()
	b := dbuf.Bytes(int(m.Bytes))
	if dbuf.err != nil {
		return dbuf.err
	}
	if len(b) < 4 {
		return fmt.Errorf("decode error")
	}
	m.ClassId = int32(binary.LittleEndian.Uint32(b))
	if !CheckClassID(m.ClassId) {
		return fmt.Errorf("not register classId: %d", m.ClassId)
	}
	m.Body = b
	return nil
}

///////////////////////////////////////////////////////////////////////////////
//message2 msg_id:long seqno:int bytes:int body:Object = Message2; // parsed manually
type TLMessage2 struct {
	MsgId  int64
	Seqno  int32
	Bytes  int32
	Object TLObject
}

func (m *TLMessage2) String() string {
	return fmt.Sprintf("{message2#5bb8e511 - msg_id: %d, seq_no: %d, object: {%v}}",
		m.MsgId,
		m.Seqno,
		m.Object)
}

func (m *TLMessage2) Encode() []byte {
	x := NewEncodeBuf(512)
	// x.Int(int32(TLConstructor_CRC32_message2))
	x.Long(m.MsgId)
	x.Int(m.Seqno)
	x.Int(m.Bytes)
	x.Bytes(m.Object.Encode())
	return x.buf
}

func (m *TLMessage2) EncodeToLayer(layer int) []byte {
	x := NewEncodeBuf(512)
	// x.Int(int32(TLConstructor_CRC32_message2))
	x.Long(m.MsgId)
	x.Int(m.Seqno)
	x.Int(m.Bytes)
	x.Bytes(m.Object.EncodeToLayer(layer))
	return x.buf
}

func (m *TLMessage2) Decode(dbuf *DecodeBuf) error {
	m.MsgId = dbuf.Long()
	m.Seqno = dbuf.Int()
	m.Bytes = dbuf.Int()
	// glog.Infof("message2: {msg_id: %d, seqno: %d, bytes: %d}", m.MsgId, )
	b := dbuf.Bytes(int(m.Bytes))

	dbuf2 := NewDecodeBuf(b)
	m.Object = dbuf2.Object()
	if m.Object == nil {
		err := fmt.Errorf("decode core_message error: %s", hex.EncodeToString(b))
		glog.Error(err)
		return err
	}

	// glog.Info("Sucess decoded core_message: ", m.Object.String())
	return dbuf2.err
}

///////////////////////////////////////////////////////////////////////////////
//msg_container#73f1f8dc messages:vector<message2> = MessageContainer; // parsed manually
type TLMsgContainer struct {
	Messages []TLMessage2
	// RawMessages []TLMessageRawData

}

func (m *TLMsgContainer) String() string {
	return "{msg_container#73f1f8dc}"
}

func (m *TLMsgContainer) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_msg_container))

	//if len(m.RawMessages) > 0 {
	//	x.Int(int32(len(m.RawMessages)))
	//	for _, m := range m.RawMessages {
	//		x.Bytes(m.Encode())
	//	}
	//	return x.buf
	//} else {
	//	x.Int(int32(len(m.Messages)))
	//	for _, m := range m.Messages {
	//		x.Bytes(m.Encode())
	//	}
	//	return x.buf
	//}

	x.Int(int32(len(m.Messages)))
	for _, m := range m.Messages {
		x.Bytes(m.Encode())
	}
	return x.buf
}

func (m *TLMsgContainer) EncodeToLayer(layer int) []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_msg_container))
	x.Int(int32(len(m.Messages)))
	for _, m := range m.Messages {
		x.Bytes(m.EncodeToLayer(layer))
	}
	return x.buf
}

func (m *TLMsgContainer) Decode(dbuf *DecodeBuf) error {
	len := dbuf.Int()
	glog.Info("msg_container: messages size: ", len)
	for i := 0; i < int(len); i++ {
		glog.Infof("msg_container: decode messages[%d]: ", i)
		message2 := &TLMessage2{}
		err := message2.Decode(dbuf)
		if err != nil {
			glog.Error("Decode message2 error: ", err)
			return err
		}
		m.Messages = append(m.Messages, *message2)
	}
	return dbuf.err
}

///////////////////////////////////////////////////////////////////////////////
//msg_copy#e06046b2 orig_message:Message2 = MessageCopy; // parsed manually, not used - use msg_container
type TLMsgCopy struct {
	OrigMessage TLMessage2
}

func (m *TLMsgCopy) String() string {
	return "{msg_copy#e06046b2}"
}

func (m *TLMsgCopy) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_msg_copy))
	x.Bytes(m.OrigMessage.Encode())
	return x.buf
}

func (m *TLMsgCopy) EncodeToLayer(layer int) []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_msg_copy))
	x.Bytes(m.OrigMessage.EncodeToLayer(layer))
	return x.buf
}

func (m *TLMsgCopy) Decode(dbuf *DecodeBuf) error {
	o := dbuf.Object()
	message2, _ := o.(*TLMessage2)
	m.OrigMessage = *message2
	return dbuf.err
}

///////////////////////////////////////////////////////////////////////////////
//gzip_packed#3072cfa1 packed_data:string = Object; // parsed manually
type TLGzipPacked struct {
	PackedData []byte
	Obj        TLObject
}

func (m *TLGzipPacked) String() string {
	return "{gzip_packed#3072cfa1}"
}

func (m *TLGzipPacked) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_gzip_packed))
	x.Bytes(m.PackedData)
	return x.buf
}

func (m *TLGzipPacked) EncodeToLayer(int) []byte {
	return m.Encode()
}

func (m *TLGzipPacked) Decode(dbuf *DecodeBuf) error {
	data := dbuf.StringBytes()
	if dbuf.err != nil {
		return dbuf.err
	}
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		dbuf.err = err
		return fmt.Errorf("gzip read: %v", err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		dbuf.err = err
		return fmt.Errorf("gzip read: %v", err)
	}
	if clErr != nil {
		dbuf.err = clErr
		return clErr
	}

	m.PackedData = buf.Bytes()

	dbuf2 := NewDecodeBuf(m.PackedData)
	m.Obj = dbuf2.Object()
	dbuf.err = dbuf2.err

	return dbuf.err
}

///////////////////////////////////////////////////////////////////////////////
//rpc_result#f35c6d01 req_msg_id:long result:Object = RpcResult; // parsed manually
type TLRpcResult struct {
	ReqMsgId int64
	Result   TLObject
}

func (m *TLRpcResult) String() string {
	return "{rpc_result#f35c6d01: req_msg_id:" + string(m.ReqMsgId) + "}"
}

func (m *TLRpcResult) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_rpc_result))
	x.Long(m.ReqMsgId)
	x.Bytes(m.Result.Encode())
	return x.buf
}

func (m *TLRpcResult) EncodeToLayer(layer int) []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_rpc_result))
	x.Long(m.ReqMsgId)
	x.Bytes(m.Result.EncodeToLayer(layer))
	return x.buf
}

func (m *TLRpcResult) Decode(dbuf *DecodeBuf) error {
	m.ReqMsgId = dbuf.Long()
	m.Result = dbuf.Object()
	return dbuf.err
}

///////////////////////////////////////////////////////////////////////////////
// contacts.getContactsLayer70#22c6aa08 hash:string = contacts.Contacts;
//func NewTLContactsGetContactsLayer70() *TLContactsGetContactsLayer70 {
//	return &TLContactsGetContactsLayer70{}
//}
//
//func (m *TLContactsGetContactsLayer70) Encode() []byte {
//	x := NewEncodeBuf(512)
//	x.Int(int32(TLConstructor_CRC32_contacts_getContactsLayer70))
//
//	x.String(m.Hash)
//
//	return x.buf
//}
//
//func (m *TLContactsGetContactsLayer70) Decode(dbuf *DecodeBuf) error {
//	m.Hash = dbuf.String()
//
//	return dbuf.err
//}

////////////////////////////////////////////////////////////////////////////////
// Vector

////////////////////////////////////////////////////////////////////////////////
//// Vector api result type
//message Vector_WallPaper {
//    repeated WallPaper datas = 1;
//}
func NewVector_WallPaper() *Vector_WallPaper {
	return &Vector_WallPaper{}
}

func (m *Vector_WallPaper) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_vector))
	x.Int(int32(len(m.Datas)))
	for _, v := range m.Datas {
		x.buf = append(x.buf, (*v).Encode()...)
	}
	return x.buf
}

func (m *Vector_WallPaper) EncodeToLayer(layer int) []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_vector))
	x.Int(int32(len(m.Datas)))
	for _, v := range m.Datas {
		x.buf = append(x.buf, (*v).EncodeToLayer(layer)...)
	}
	return x.buf
}

func (m *Vector_WallPaper) Decode(dbuf *DecodeBuf) error {
	dbuf.Int() // TODO(@benqi): Check crc32 invalid
	l1 := dbuf.Int()
	m.Datas = make([]*WallPaper, l1)
	for i := int32(0); i < l1; i++ {
		m.Datas[i] = &WallPaper{}
		(*m.Datas[i]).Decode(dbuf)
	}

	return dbuf.err
}

////////////////////////////////////////////////////////////////////////////////
//message Vector_User {
//    repeated User datas = 1;
//}
func NewVector_User() *Vector_User {
	return &Vector_User{}
}
func (m *Vector_User) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_vector))
	x.Int(int32(len(m.Datas)))
	for _, v := range m.Datas {
		x.buf = append(x.buf, (*v).Encode()...)
	}
	return x.buf
}

func (m *Vector_User) EncodeToLayer(layer int) []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_vector))
	x.Int(int32(len(m.Datas)))
	for _, v := range m.Datas {
		x.buf = append(x.buf, (*v).EncodeToLayer(layer)...)
	}
	return x.buf
}

func (m *Vector_User) Decode(dbuf *DecodeBuf) error {
	dbuf.Int() // TODO(@benqi): Check crc32 invalid
	l1 := dbuf.Int()
	m.Datas = make([]*User, l1)
	for i := int32(0); i < l1; i++ {
		m.Datas[i] = &User{}
		(*m.Datas[i]).Decode(dbuf)
	}

	return dbuf.err
}

////////////////////////////////////////////////////////////////////////////////
//message Vector_ContactStatus {
//    repeated ContactStatus datas = 1;
//}
func NewVector_ContactStatus() *Vector_ContactStatus {
	return &Vector_ContactStatus{}
}
func (m *Vector_ContactStatus) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_vector))
	x.Int(int32(len(m.Datas)))
	for _, v := range m.Datas {
		x.buf = append(x.buf, (*v).Encode()...)
	}
	return x.buf
}

func (m *Vector_ContactStatus) EncodeToLayer(layer int) []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_vector))
	x.Int(int32(len(m.Datas)))
	for _, v := range m.Datas {
		x.buf = append(x.buf, (*v).EncodeToLayer(layer)...)
	}
	return x.buf
}

func (m *Vector_ContactStatus) Decode(dbuf *DecodeBuf) error {
	dbuf.Int() // TODO(@benqi): Check crc32 invalid
	l1 := dbuf.Int()
	m.Datas = make([]*ContactStatus, l1)
	for i := int32(0); i < l1; i++ {
		m.Datas[i] = &ContactStatus{}
		(*m.Datas[i]).Decode(dbuf)
	}

	return dbuf.err
}

////////////////////////////////////////////////////////////////////////////////
//message Vector_int {
//    repeated int32 datas = 1;
//}
func NewVectorInt() *VectorInt {
	return &VectorInt{}
}

func (m *VectorInt) Encode() []byte {
	x := NewEncodeBuf(512)
	x.VectorInt(m.Datas)
	return x.buf
}

func (m *VectorInt) EncodeToLayer(int) []byte {
	return m.Encode()
}

func (m *VectorInt) Decode(dbuf *DecodeBuf) error {
	// dbuf.Int() // TODO(@benqi): Check crc32 invalid
	m.Datas = dbuf.VectorInt()
	return dbuf.err
}

//message Vector_ReceivedNotifyMessage {
//    repeated ReceivedNotifyMessage datas = 1;
//}
func NewVector_ReceivedNotifyMessage() *Vector_ReceivedNotifyMessage {
	return &Vector_ReceivedNotifyMessage{}
}
func (m *Vector_ReceivedNotifyMessage) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_vector))
	x.Int(int32(len(m.Datas)))
	for _, v := range m.Datas {
		x.buf = append(x.buf, (*v).Encode()...)
	}
	return x.buf
}

func (m *Vector_ReceivedNotifyMessage) EncodeToLayer(layer int) []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_vector))
	x.Int(int32(len(m.Datas)))
	for _, v := range m.Datas {
		x.buf = append(x.buf, (*v).EncodeToLayer(layer)...)
	}
	return x.buf
}

func (m *Vector_ReceivedNotifyMessage) Decode(dbuf *DecodeBuf) error {
	dbuf.Int() // TODO(@benqi): Check crc32 invalid
	l1 := dbuf.Int()
	m.Datas = make([]*ReceivedNotifyMessage, l1)
	for i := int32(0); i < l1; i++ {
		m.Datas[i] = &ReceivedNotifyMessage{}
		(*m.Datas[i]).Decode(dbuf)
	}

	return dbuf.err
}

////////////////////////////////////////////////////////////////////////////////
//message Vector_long {
//    repeated int64 datas = 1;
//}
func NewVectorLong() *VectorLong {
	return &VectorLong{}
}
func (m *VectorLong) Encode() []byte {
	x := NewEncodeBuf(512)
	x.VectorLong(m.Datas)
	return x.buf
}

func (m *VectorLong) EncodeToLayer(int) []byte {
	return m.Encode()
}

func (m *VectorLong) Decode(dbuf *DecodeBuf) error {
	m.Datas = dbuf.VectorLong()

	return dbuf.err
}

////////////////////////////////////////////////////////////////////////////////
//message Vector_StickerSetCovered {
//    repeated StickerSetCovered datas = 1;
//}
func NewVector_StickerSetCovered() *Vector_StickerSetCovered {
	return &Vector_StickerSetCovered{}
}
func (m *Vector_StickerSetCovered) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_vector))
	x.Int(int32(len(m.Datas)))
	for _, v := range m.Datas {
		x.buf = append(x.buf, (*v).Encode()...)
	}
	return x.buf
}

func (m *Vector_StickerSetCovered) EncodeToLayer(layer int) []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_vector))
	x.Int(int32(len(m.Datas)))
	for _, v := range m.Datas {
		x.buf = append(x.buf, (*v).EncodeToLayer(layer)...)
	}
	return x.buf
}

func (m *Vector_StickerSetCovered) Decode(dbuf *DecodeBuf) error {
	dbuf.Int() // TODO(@benqi): Check crc32 invalid
	l1 := dbuf.Int()
	m.Datas = make([]*StickerSetCovered, l1)
	for i := int32(0); i < l1; i++ {
		m.Datas[i] = &StickerSetCovered{}
		(*m.Datas[i]).Decode(dbuf)
	}

	return dbuf.err
}

////////////////////////////////////////////////////////////////////////////////
//message Vector_FileHash {
//    repeated FileHash datas = 1;
//}
func NewVector_FileHash() *Vector_FileHash {
	return &Vector_FileHash{}
}
func (m *Vector_FileHash) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_vector))
	x.Int(int32(len(m.Datas)))
	for _, v := range m.Datas {
		x.buf = append(x.buf, (*v).Encode()...)
	}
	return x.buf
}

func (m *Vector_FileHash) EncodeToLayer(layer int) []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_vector))
	x.Int(int32(len(m.Datas)))
	for _, v := range m.Datas {
		x.buf = append(x.buf, (*v).EncodeToLayer(layer)...)
	}
	return x.buf
}

func (m *Vector_FileHash) Decode(dbuf *DecodeBuf) error {
	dbuf.Int() // TODO(@benqi): Check crc32 invalid
	l1 := dbuf.Int()
	m.Datas = make([]*FileHash, l1)
	for i := int32(0); i < l1; i++ {
		m.Datas[i] = &FileHash{}
		(*m.Datas[i]).Decode(dbuf)
	}

	return dbuf.err
}

////////////////////////////////////////////////////////////////////////////////
//message Vector_LangPackString {
//    repeated LangPackString datas = 1;
//}
func NewVector_LangPackString() *Vector_LangPackString {
	return &Vector_LangPackString{}
}
func (m *Vector_LangPackString) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_vector))
	x.Int(int32(len(m.Datas)))
	for _, v := range m.Datas {
		x.buf = append(x.buf, (*v).Encode()...)
	}
	return x.buf
}

func (m *Vector_LangPackString) EncodeToLayer(layer int) []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_vector))
	x.Int(int32(len(m.Datas)))
	for _, v := range m.Datas {
		x.buf = append(x.buf, (*v).EncodeToLayer(layer)...)
	}
	return x.buf
}

func (m *Vector_LangPackString) Decode(dbuf *DecodeBuf) error {
	dbuf.Int() // TODO(@benqi): Check crc32 invalid
	l1 := dbuf.Int()
	m.Datas = make([]*LangPackString, l1)
	for i := int32(0); i < l1; i++ {
		m.Datas[i] = &LangPackString{}
		(*m.Datas[i]).Decode(dbuf)
	}

	return dbuf.err
}

////////////////////////////////////////////////////////////////////////////////
//message Vector_LangPackLanguage {
//    repeated LangPackLanguage datas = 1;
//}
func NewVector_LangPackLanguage() *Vector_LangPackLanguage {
	return &Vector_LangPackLanguage{}
}
func (m *Vector_LangPackLanguage) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_vector))
	x.Int(int32(len(m.Datas)))
	for _, v := range m.Datas {
		x.buf = append(x.buf, (*v).Encode()...)
	}
	return x.buf
}

func (m *Vector_LangPackLanguage) EncodeToLayer(layer int) []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_vector))
	x.Int(int32(len(m.Datas)))
	for _, v := range m.Datas {
		x.buf = append(x.buf, (*v).EncodeToLayer(layer)...)
	}
	return x.buf
}

func (m *Vector_LangPackLanguage) Decode(dbuf *DecodeBuf) error {
	dbuf.Int() // TODO(@benqi): Check crc32 invalid
	l1 := dbuf.Int()
	m.Datas = make([]*LangPackLanguage, l1)
	for i := int32(0); i < l1; i++ {
		m.Datas[i] = &LangPackLanguage{}
		(*m.Datas[i]).Decode(dbuf)
	}

	return dbuf.err
}
