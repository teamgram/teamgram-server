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
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"io"

	"github.com/panjf2000/gnet/v2"
	"google.golang.org/protobuf/proto"
)

var (
	brpcMagicNumber = []byte{'P', 'R', 'B', 'C'}
)

// PARSE_ERROR_TOO_BIG_DATA

var (
	ErrIncompletePacket = errors.New("incomplete packet")
)

// BaiduRpcMessage is the message type for brpc.
// //////////////////////////////////////////////////////////////////////////////
type BaiduRpcMessage struct {
	Meta       *RpcMeta
	Payload    []byte
	Attachment []byte
}

func (m *BaiduRpcMessage) String() string {
	return fmt.Sprintf("{meta={%v}, payload_size=%d, attachment_size=%d}", m.Meta, len(m.Payload), len(m.Attachment))
}

func (m *BaiduRpcMessage) Encode() ([][]byte, error) {
	var (
		headBuf = make([]byte, 12)
	)

	if len(m.Attachment) > 0 {
		m.Meta.AttachmentSize = int32(len(m.Attachment))
	}

	metaData, err := proto.Marshal(m.Meta)
	if err != nil {
		return nil, err
	}

	// pSize := len(msg.Payload) + len(msg.Attachment) + proto.Size(msg.Meta)
	copy(headBuf[:4], brpcMagicNumber)
	binary.LittleEndian.PutUint32(headBuf[4:], uint32(12+len(m.Payload)+len(m.Attachment)+len(metaData)))
	binary.LittleEndian.PutUint32(headBuf[8:], uint32(len(metaData)))

	bufList := make([][]byte, 0, 4)

	bufList = append(bufList, headBuf)
	bufList = append(bufList, m.Payload)
	bufList = append(bufList, m.Attachment)
	bufList = append(bufList, metaData)

	return bufList, nil
}

type BaiduRpcCodec struct {
	// headBuf [12]byte
}

func NewBaiduRpcCodec() *BaiduRpcCodec {
	return new(BaiduRpcCodec)
}

func (codec *BaiduRpcCodec) Encode(msg *BaiduRpcMessage) ([][]byte, error) {
	return msg.Encode()
}

func (codec *BaiduRpcCodec) Decode(c gnet.Conn) (*BaiduRpcMessage, error) {
	headBuf, err := c.Peek(12)
	if errors.Is(err, io.ErrShortBuffer) {
		return nil, ErrIncompletePacket
	} else if len(headBuf) < 12 {
		return nil, ErrIncompletePacket
	}

	if !bytes.Equal(headBuf[:4], brpcMagicNumber) {
		err = fmt.Errorf("invalid magic number: %s", string(headBuf[:4]))
		return nil, err
	}

	bodySize := int(binary.LittleEndian.Uint32(headBuf[4:8]))
	metaSize := int(binary.LittleEndian.Uint32(headBuf[8:]))

	if bodySize > 64*1024*1024 {
		// We need this log to report the body_size to give users some clues
		// which is not printed in InputMessenger.
		err = fmt.Errorf("body_size=%d, from %s is too large", bodySize, c.String())
		return nil, err
	}

	if metaSize > bodySize {
		err = fmt.Errorf("meta_size=%d is bigger than body_size=%d", metaSize, bodySize)
		return nil, err
	}

	if c.InboundBuffered() < bodySize {
		return nil, ErrIncompletePacket
	}

	c.Discard(12)
	bodySize = bodySize - 12 - metaSize
	bodyData, _ := c.Next(bodySize)
	metaData, _ := c.Next(metaSize)

	meta := new(RpcMeta)
	proto.Unmarshal(metaData, meta)

	attachmentSize := int(meta.GetAttachmentSize())

	message := &BaiduRpcMessage{
		Meta:       meta,
		Payload:    bodyData[:bodySize-attachmentSize],
		Attachment: bodyData[bodySize-attachmentSize:],
	}

	return message, nil
}
