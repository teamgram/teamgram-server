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

package brpc

import (
	"github.com/nebula-chat/chatengine/pkg/net2"
	"io"
	"encoding/binary"
	"net"
	"bytes"
	"github.com/golang/glog"
	"fmt"
	"github.com/gogo/protobuf/proto"
	"github.com/nebula-chat/chatengine/pkg/util"
)

func init() {
	net2.RegisterProtocol("brpc", &BaiduRpcProtocol{})
}

// PARSE_ERROR_TOO_BIG_DATA
var brpcMagicNumber = []byte{'P', 'R', 'B', 'C'}

type BaiduRpcProtocol struct {
}

func (m *BaiduRpcProtocol) NewCodec(rw io.ReadWriter) (cc net2.Codec, err error) {
	codec := &BaiduRpcCodec{rw.(*net.TCPConn)}
	return codec, nil
}

func NewBaiduRpcProtocol() net2.Protocol {
	return new(BaiduRpcProtocol)
}

////////////////////////////////////////////////////////////////////////////////
type BaiduRpcMessage struct {
	Meta       *RpcMeta
	Payload    []byte
	Attachment []byte
}

func (m *BaiduRpcMessage) encode() [] byte {
	x := util.NewBufferOutput(512)
	var (
		payloadSize = len(m.Payload)
		metaSize int
		attachmentSize = len(m.Attachment)
	)

	if attachmentSize > 0 {
		m.Meta.AttachmentSize = new(int32)
		*m.Meta.AttachmentSize = int32(attachmentSize)
	}
	if m.Meta != nil {
		metaSize = proto.Size(m.Meta)
	}

	x.Bytes(brpcMagicNumber)
	x.Int32(int32(payloadSize + attachmentSize + metaSize))
	x.Int32(int32(metaSize))

	x.Bytes(m.Payload)
	if attachmentSize > 0 {
		x.Bytes(m.Attachment)
	}
	mData, _ := proto.Marshal(m.Meta)
	x.Bytes(mData)

	return x.Buf()
}

type BaiduRpcCodec struct {
	*net.TCPConn
}

func (c *BaiduRpcCodec) Send(msg interface{}) error {
	var (
		err error
		message *BaiduRpcMessage
	)

	// glog.Info("send: ", msg)

	switch msg.(type) {
	case *BaiduRpcMessage:
		message, _ = msg.(*BaiduRpcMessage)
		b := message.encode()
		_, err = c.TCPConn.Write(b)
	default:
		err = fmt.Errorf("invalid zproto message: %v", msg)
	}

	return err
}

func (c *BaiduRpcCodec) Receive() (interface{}, error) {
	var (
		headBuf = make([]byte, 12)
		bodySize uint32
		metaSize uint32
	)

	_, err := io.ReadFull(c.TCPConn, headBuf)
	if err != nil {
		glog.Error(err)
		return nil, err
	}

	if !bytes.Equal(headBuf[:4], brpcMagicNumber) {
		err = fmt.Errorf("invalid magic number: %s", string(headBuf[:4]))
		glog.Error(err)
		return nil, err
	}

	bodySize = binary.LittleEndian.Uint32(headBuf[4:8])
	metaSize = binary.LittleEndian.Uint32(headBuf[8:])

	if bodySize > 64*1024*1024 {
		// We need this log to report the body_size to give users some clues
		// which is not printed in InputMessenger.
		err = fmt.Errorf("body_size=%d, from %s is too large", bodySize, c.TCPConn.RemoteAddr())
		return nil, err
	}

	if metaSize > bodySize {
		err = fmt.Errorf("meta_size=%d is bigger than body_size=%d", metaSize, bodySize)
		return nil, err
	}
	bodySize -= metaSize
	bodyData := make([]byte, bodySize)
	_, err = io.ReadFull(c.TCPConn, bodyData)
	if err != nil {
		glog.Error(err)
		return nil, err
	}

	metaData := make([]byte, metaSize)
	_, err = io.ReadFull(c.TCPConn, metaData)
	if err != nil {
		glog.Error(err)
		return nil, err
	}

	meta := new(RpcMeta)
	err = proto.Unmarshal(metaData, meta)
	if err != nil {
		glog.Error(err)
		return nil, err
	}

	attachedSize := uint32(meta.GetAttachmentSize())
	if attachedSize > bodySize {
		err = fmt.Errorf("attached_size=%d is bigger than body_size=%d", attachedSize, bodySize)
		return nil, err
	}

	message := &BaiduRpcMessage{
		Meta:       meta,
		Payload:    bodyData[:bodySize-attachedSize],
		Attachment: bodyData[bodySize-attachedSize:],
	}

	return message, nil
}

func (c *BaiduRpcCodec) Close() error {
	return c.TCPConn.Close()
}
