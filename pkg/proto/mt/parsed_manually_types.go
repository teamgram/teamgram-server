// Copyright (c) 2026-present, The Teamgram Authors (https://teamgram.net).
//  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package mt

import (
	"bytes"
	"compress/gzip"
	"compress/zlib"
	"encoding/binary"
	"fmt"
	"io"

	"github.com/teamgram/teamgram-server/v2/pkg/proto/bin"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/iface"
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
		return fmt.Errorf("unable to decode message2: field msg_id: %w", err)
	}

	m.Seqno, err = d.Int32()
	if err != nil {
		return fmt.Errorf("unable to decode message2: field seqno: %w", err)
	}

	m.Bytes, err = d.Int32()
	if err != nil {
		return fmt.Errorf("unable to decode message2: field bytes: %w", err)
	}

	m.ClazzID, err = d.ClazzID()
	if err != nil {
		return fmt.Errorf("unable to decode message2: field clazz_id: %w", err)
	}

	m.Body = d.Raw()

	return nil
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
	for i, v := range m.Messages {
		if v == nil {
			return fmt.Errorf("unable to encode msg_container: field messages element %d is nil", i)
		}
		if err := v.Encode(x, layer); err != nil {
			return fmt.Errorf("unable to encode msg_container: field messages element %d: %w", i, err)
		}
	}

	return nil
}

func (m *TLMsgRawDataContainer) Decode(d *bin.Decoder) error {
	len2, err := d.Int()
	if err != nil {
		return fmt.Errorf("unable to decode msg_container: field messages length: %w", err)
	}

	m.Messages = make([]*TLMessageRawData, 0, len2)
	for i := 0; i < len2; i++ {
		message2 := new(TLMessageRawData)
		err = message2.Decode(d)
		if err != nil {
			return fmt.Errorf("unable to decode msg_container: field messages element %d: %w", i, err)
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
	if m.Object == nil {
		return fmt.Errorf("unable to encode message2: field body is nil")
	}
	if err := m.Object.Encode(x, layer); err != nil {
		return fmt.Errorf("unable to encode message2: field body: %w", err)
	}
	b := x.Bytes()

	binary.LittleEndian.PutUint32(b[offset:], uint32(x.Len()-offset-4))

	return nil
}

func (m *TLMessage2) Decode(d *bin.Decoder) (err error) {
	m.MsgId, err = d.Int64()
	if err != nil {
		return fmt.Errorf("unable to decode message2: field msg_id: %w", err)
	}

	m.Seqno, err = d.Int32()
	if err != nil {
		return fmt.Errorf("unable to decode message2: field seqno: %w", err)
	}

	m.Bytes, err = d.Int32()
	if err != nil {
		return fmt.Errorf("unable to decode message2: field bytes: %w", err)
	}

	b := make([]byte, m.Bytes)
	err = d.ConsumeN(b, int(m.Bytes))
	if err != nil {
		return fmt.Errorf("unable to decode message2: field body bytes: %w", err)
	}

	m.Object, err = iface.DecodeObject(bin.NewDecoder(b))
	if err != nil {
		return fmt.Errorf("unable to decode message2: field body: %w", err)
	}

	return nil
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
	for i, v := range m.Messages {
		if v == nil {
			return fmt.Errorf("unable to encode msg_container: field messages element %d is nil", i)
		}
		if err := v.Encode(x, layer); err != nil {
			return fmt.Errorf("unable to encode msg_container: field messages element %d: %w", i, err)
		}
	}

	return nil
}

func (m *TLMsgContainer) Decode(d *bin.Decoder) error {
	len2, err := d.Int()
	if err != nil {
		return fmt.Errorf("unable to decode msg_container: field messages length: %w", err)
	}

	m.Messages = make([]*TLMessage2, 0, len2)

	for i := 0; i < len2; i++ {
		message2 := new(TLMessage2)
		err = message2.Decode(d)
		if err != nil {
			return fmt.Errorf("unable to decode msg_container: field messages element %d: %w", i, err)
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
	if m.OrigMessage == nil {
		return fmt.Errorf("unable to encode msg_copy: field orig_message is nil")
	}
	if err := m.OrigMessage.Encode(x, layer); err != nil {
		return fmt.Errorf("unable to encode msg_copy: field orig_message: %w", err)
	}

	return nil
}

func (m *TLMsgCopy) Decode(d *bin.Decoder) error {
	message2 := new(TLMessage2)
	err := message2.Decode(d)
	if err != nil {
		return fmt.Errorf("unable to decode msg_copy: field orig_message: %w", err)
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
		return fmt.Errorf("unable to encode gzip_packed: field packed_data is empty")
	}

	var (
		err error
		b   = new(bytes.Buffer)
	)
	gz := gzip.NewWriter(b)
	_, err = gz.Write(m.PackedData)
	if err == nil {
		err = gz.Flush()
	}
	clErr := gz.Close()

	if err != nil {
		return fmt.Errorf("unable to encode gzip_packed: compress packed_data: %w", err)
	}
	if clErr != nil {
		return fmt.Errorf("unable to encode gzip_packed: close gzip writer: %w", clErr)
	}

	x.PutClazzID(ClazzID_gzip_packed)
	x.PutBytes(b.Bytes())

	return nil
}

func (m *TLGzipPacked) Decode(d *bin.Decoder) error {
	data, err := d.Bytes()
	if err != nil {
		return fmt.Errorf("unable to decode gzip_packed: field packed_data: %w", err)
	}

	var (
		gz io.ReadCloser
		// err error
	)

	gz, err = gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		gz, err = zlib.NewReader(bytes.NewBuffer(data))
		if err != nil {
			return fmt.Errorf("unable to decode gzip_packed: create decompressor: %w", err)
		}
	}

	var (
		buf bytes.Buffer
	)

	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return fmt.Errorf("unable to decode gzip_packed: decompress packed_data: %w", err)
	}
	if clErr != nil {
		return fmt.Errorf("unable to decode gzip_packed: close decompressor: %w", clErr)
	}

	m.PackedData = buf.Bytes()

	d2 := bin.NewDecoder(m.PackedData)
	m.Obj, err = iface.DecodeObject(d2)
	if err != nil {
		return fmt.Errorf("unable to decode gzip_packed: field object: %w", err)
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

	if m.Result == nil {
		return fmt.Errorf("unable to encode rpc_result: field result is nil")
	}
	if err := m.Result.Encode(x2, layer); err != nil {
		return fmt.Errorf("unable to encode rpc_result: field result: %w", err)
	}
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
			if err := gzipPacked.Encode(x, layer); err != nil {
				return fmt.Errorf("unable to encode rpc_result: gzip packed result: %w", err)
			}
		}
	} else {
		x.Put(x2.Bytes())
	}

	return nil
}

func (m *TLRpcResult) Decode(d *bin.Decoder) (err error) {
	m.ReqMsgId, err = d.Int64()
	if err != nil {
		return fmt.Errorf("unable to decode rpc_result: field req_msg_id: %w", err)
	}

	m.Result, err = iface.DecodeObject(d)
	if err != nil {
		return fmt.Errorf("unable to decode rpc_result: field result: %w", err)
	}

	return nil
}
