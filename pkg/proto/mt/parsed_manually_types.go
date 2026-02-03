// Copyright 2024 Teamgooo Authors
//  All rights reserved.
//
// Author: Benqi (wubenqi@gmail.com)
//

package mt

import (
	"bytes"
	"compress/gzip"
	"compress/zlib"
	"encoding/binary"
	"fmt"
	"io"

	"github.com/teamgooo/teamgooo-server/pkg/proto/bin"
	"github.com/teamgooo/teamgooo-server/pkg/proto/iface"
)

const (
	// ClazzID_message2      = 1538843921 // parsed_manually_types

	ClazzID_msg_container = 0x73f1f8dc // parsed_manually_types
	ClazzID_msg_copy      = 0xe06046b2 // parsed_manually_types
	ClazzID_gzip_packed   = 0x3072cfa1 // parsed_manually_types
	ClazzID_rpc_result    = 0xf35c6d01 // parsed_manually_types
)

func init() {
	// iface.RegisterClazzID(ClazzID_message2, func() iface.TLObject { return &TLMessage2{} })

	iface.RegisterClazzID(ClazzID_msg_container, func() iface.TLObject { return &TLMsgContainer{} })
	iface.RegisterClazzID(ClazzID_msg_copy, func() iface.TLObject { return &TLMsgCopy{} })
	iface.RegisterClazzID(ClazzID_gzip_packed, func() iface.TLObject { return &TLGzipPacked{} })
	iface.RegisterClazzID(ClazzID_rpc_result, func() iface.TLObject { return &TLRpcResult{} })
}

// TLMessageRawData
// message2 msg_id:long seqno:int bytes:int body:Object = Message2; // parsed manually
type TLMessageRawData struct {
	MsgId   int64  `json:"msg_id"`
	Seqno   int32  `json:"seqno"`
	Bytes   int32  `json:"bytes"`
	ClazzID uint32 `json:"clazz_id"`
	Body    []byte `json:"body"`
}

func (m *TLMessageRawData) ClazzName() string {
	return "message2"
}

func (m *TLMessageRawData) Encode(x *bin.Encoder, layer int32) error {
	_ = layer

	x.PutInt64(m.MsgId)
	x.PutInt32(m.Seqno)
	x.PutInt32(m.Bytes)
	x.Put(m.Body)

	return nil
}

func (m *TLMessageRawData) Decode(d *bin.Decoder) (err error) {
	m.MsgId, err = d.Int64()
	if err != nil {
		return
	}

	m.Seqno, err = d.Int32()
	if err != nil {
		return
	}

	m.Bytes, err = d.Int32()
	if err != nil {
		return
	}

	m.ClazzID, err = d.ClazzID()
	if err != nil {
		return
	}

	m.Body = d.Raw()

	return
}

// TLMsgRawDataContainer
// msg_container#73f1f8dc messages:vector<message2> = MessageContainer; // parsed manually
type TLMsgRawDataContainer struct {
	Messages []*TLMessageRawData
}

func (m *TLMsgRawDataContainer) ClazzName() string {
	return "msg_container"
}

func (m *TLMsgRawDataContainer) Encode(x *bin.Encoder, layer int32) error {
	x.PutClazzID(ClazzID_msg_container)
	x.PutInt(len(m.Messages))
	for _, v := range m.Messages {
		_ = v.Encode(x, layer)
	}

	return nil
}

func (m *TLMsgRawDataContainer) Decode(d *bin.Decoder) error {
	len2, err := d.Int()
	if err != nil {
		return err
	}

	m.Messages = make([]*TLMessageRawData, 0, len2)
	for i := 0; i < len2; i++ {
		message2 := new(TLMessageRawData)
		err = message2.Decode(d)
		if err != nil {
			return err
		}
		m.Messages = append(m.Messages, message2)
	}

	return nil
}

// TLMessage2
// message2 msg_id:long seqno:int bytes:int body:Object = Message2; // parsed manually
type TLMessage2 struct {
	MsgId  int64
	Seqno  int32
	Bytes  int32
	Object iface.TLObject
}

func (m *TLMessage2) ClazzName() string {
	return "message2"
}

func (m *TLMessage2) Encode(x *bin.Encoder, layer int32) error {
	x.PutInt64(m.MsgId)
	x.PutInt32(m.Seqno)

	offset := x.Len()

	x.PutInt32(m.Bytes)
	_ = m.Object.Encode(x, layer)
	b := x.Bytes()

	binary.LittleEndian.PutUint32(b[offset:], uint32(x.Len()-offset-4))

	return nil
}

func (m *TLMessage2) Decode(d *bin.Decoder) (err error) {
	m.MsgId, err = d.Int64()
	if err != nil {
		return
	}

	m.Seqno, err = d.Int32()
	if err != nil {
		return
	}

	m.Bytes, err = d.Int32()
	if err != nil {
		return
	}

	b := make([]byte, m.Bytes)
	err = d.ConsumeN(b, int(m.Bytes))
	if err != nil {
		return
	}

	m.Object, err = iface.DecodeObject(bin.NewDecoder(b))
	if err != nil {
		return
	}

	return
}

// TLMsgContainer
// msg_container#73f1f8dc messages:vector<message2> = MessageContainer; // parsed manually
type TLMsgContainer struct {
	Messages []*TLMessage2
}

func (m *TLMsgContainer) ClazzName() string {
	return "msg_container"
}

func (m *TLMsgContainer) Encode(x *bin.Encoder, layer int32) error {
	x.PutClazzID(ClazzID_msg_container)

	x.PutInt(len(m.Messages))
	for _, v := range m.Messages {
		_ = v.Encode(x, layer)
	}

	return nil
}

func (m *TLMsgContainer) Decode(d *bin.Decoder) error {
	len2, err := d.Int()
	if err != nil {
		return err
	}

	m.Messages = make([]*TLMessage2, 0, len2)

	for i := 0; i < len2; i++ {
		message2 := new(TLMessage2)
		err = message2.Decode(d)
		if err != nil {
			fmt.Printf("decode message2 error: %v\n", err)
			return err
		}
		m.Messages = append(m.Messages, message2)
	}

	return nil
}

// TLMsgCopy
// msg_copy#e06046b2 orig_message:Message2 = MessageCopy; // parsed manually, not used - use msg_container
type TLMsgCopy struct {
	OrigMessage *TLMessage2
}

func (m *TLMsgCopy) ClazzName() string {
	return "msg_copy"
}

func (m *TLMsgCopy) Encode(x *bin.Encoder, layer int32) error {
	x.PutClazzID(ClazzID_msg_copy)
	_ = m.OrigMessage.Encode(x, layer)

	return nil
}

func (m *TLMsgCopy) Decode(d *bin.Decoder) error {
	message2 := new(TLMessage2)
	err := message2.Decode(d)
	if err != nil {
		return err
	}
	m.OrigMessage = message2

	return nil
}

// TLGzipPacked
// gzip_packed#3072cfa1 packed_data:string = Object; // parsed manually
type TLGzipPacked struct {
	PackedData []byte
	Obj        iface.TLObject
}

func (m *TLGzipPacked) ClazzName() string {
	return "gzip_packed"
}

func (m *TLGzipPacked) Encode(x *bin.Encoder, layer int32) error {
	_ = layer

	if len(m.PackedData) == 0 {
		return nil
	}

	var (
		err error
		b   = new(bytes.Buffer)
	)
	gz := gzip.NewWriter(b)
	// _, err = io.Copy(gz, bytes.NewBuffer(m.PackedData))
	_, err = gz.Write(m.PackedData)
	_ = gz.Flush()
	clErr := gz.Close()

	if err != nil {
		// log.Errorf("gzip write: %v", err)
		x.Put(m.PackedData)
		return nil
	}
	if clErr != nil {
		// log.Errorf("gzip write: %v", err)
		x.Put(m.PackedData)
		return nil
	}

	x.PutClazzID(ClazzID_gzip_packed)
	x.PutBytes(b.Bytes())

	return nil
}

func (m *TLGzipPacked) Decode(d *bin.Decoder) error {
	data, err := d.Bytes()
	if err != nil {
		return err
	}

	var (
		gz io.ReadCloser
		// err error
	)

	gz, err = gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		gz, err = zlib.NewReader(bytes.NewBuffer(data))
		if err != nil {
			return fmt.Errorf("gzip read1: %v", err)
		}
	}

	var (
		buf bytes.Buffer
	)

	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return fmt.Errorf("gzip read2: %v", err)
	}
	if clErr != nil {
		return clErr
	}

	m.PackedData = buf.Bytes()

	d2 := bin.NewDecoder(m.PackedData)
	m.Obj, err = iface.DecodeObject(d2)
	if err != nil {
		return err
	}

	return nil
}

// TLRpcResult
// rpc_result#f35c6d01 req_msg_id:long result:Object = RpcResult; // parsed manually
type TLRpcResult struct {
	ReqMsgId int64
	Result   iface.TLObject
}

func (m *TLRpcResult) ClazzName() string {
	return "rpc_result"
}

func (m *TLRpcResult) Encode(x *bin.Encoder, layer int32) error {
	var (
		c = ClazzID_rpc_result
	)

	x.PutClazzID(uint32(c))
	x.PutInt64(m.ReqMsgId)

	x2 := bin.NewEncoder()
	defer x2.End()

	_ = m.Result.Encode(x2, layer)
	// rawBody := x2.GetBuf()

	if x2.Len() > 256 {
		switch m.Result.(type) {
		//case *Upload_WebFile:
		//	x.Bytes(rawBody)
		//case *Upload_CdnFile:
		//	x.Bytes(rawBody)
		//case *Upload_File:
		//	x.Bytes(rawBody)
		default:
			gzipPacked := &TLGzipPacked{
				PackedData: x2.Bytes(),
			}
			_ = gzipPacked.Encode(x, layer)
		}
	} else {
		x.Put(x2.Bytes())
	}

	return nil
}

func (m *TLRpcResult) Decode(d *bin.Decoder) (err error) {
	m.ReqMsgId, err = d.Int64()
	if err != nil {
		return
	}

	m.Result, err = iface.DecodeObject(d)

	return
}
