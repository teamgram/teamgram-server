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
	"encoding/binary"
	"fmt"
	"github.com/golang/glog"
	"github.com/nebula-chat/chatengine/pkg/net2"
	"io"
)

// https://core.telegram.org/mtproto#tcp-transport
//
// In case 4-byte data alignment is needed,
// an intermediate version of the original protocol may be used:
// if the client sends 0xeeeeeeee as the first int (four bytes),
// then packet length is encoded always by four bytes as in the original version,
// but the sequence number and CRC32 are omitted,
// thus decreasing total packet size by 8 bytes.
//
type MTProtoIntermediateCodec struct {
	conn *net2.BufferedConn
}

func NewMTProtoIntermediateCodec(conn *net2.BufferedConn) *MTProtoIntermediateCodec {
	return &MTProtoIntermediateCodec{
		conn: conn,
	}
}

func (c *MTProtoIntermediateCodec) Receive() (interface{}, error) {
	var size int
	var n int
	var err error

	b := make([]byte, 4)
	n, err = io.ReadFull(c.conn, b)
	if err != nil {
		return nil, err
	}

	size = int(binary.LittleEndian.Uint32(b) << 2)

	// glog.Info("first_byte: ", hex.EncodeToString(b[:1]))
	// needAck := bool(b[0] >> 7 == 1)
	// _ = needAck

	//b[0] = b[0] & 0x7f
	//// glog.Info("first_byte2: ", hex.EncodeToString(b[:1]))
	//
	//if b[0] < 0x7f {
	//	size = int(b[0]) << 2
	//	glog.Info("size1: ", size)
	//	if size == 0 {
	//		return nil, nil
	//	}
	//} else {
	//	glog.Info("first_byte2: ", hex.EncodeToString(b[:1]))
	//	b2 := make([]byte, 3)
	//	n, err = io.ReadFull(c.conn, b2)
	//	if err != nil {
	//		return nil, err
	//	}
	//	size = (int(b2[0]) | int(b2[1])<<8 | int(b2[2])<<16) << 2
	//	glog.Info("size2: ", size)
	//}

	left := size
	buf := make([]byte, size)
	for left > 0 {
		n, err = io.ReadFull(c.conn, buf[size-left:])
		if err != nil {
			glog.Error("ReadFull2 error: ", err)
			return nil, err
		}
		left -= n
	}
	//if size > 10240 {
	//	glog.Info("ReadFull2: ", hex.EncodeToString(buf[:256]))
	//}

	// TODO(@benqi): process report ack and quickack
	// 截断QuickAck消息，客户端有问题
	if size == 4 {
		glog.Errorf("Server response error: ", int32(binary.LittleEndian.Uint32(buf)))
		// return nil, fmt.Errorf("Recv QuickAckMessage, ignore!!!!") //  connId: ", c.stream, ", by client ", m.RemoteAddr())
		return nil, nil
	}

	authKeyId := int64(binary.LittleEndian.Uint64(buf))
	message := NewMTPRawMessage(authKeyId, 0, TRANSPORT_TCP)
	message.Decode(buf)
	return message, nil
}

func (c *MTProtoIntermediateCodec) Send(msg interface{}) error {
	message, ok := msg.(*MTPRawMessage)
	if !ok {
		err := fmt.Errorf("msg type error, only MTPRawMessage, msg: {%v}", msg)
		glog.Error(err)
		return err
	}

	b := message.Encode()

	sb := make([]byte, 4)
	// minus padding
	size := len(b) / 4

	//if size < 127 {
	//	sb = []byte{byte(size)}
	//} else {
	binary.LittleEndian.PutUint32(sb, uint32(size))
	//}

	b = append(sb, b...)
	_, err := c.conn.Write(b)

	if err != nil {
		glog.Errorf("Send msg error: %s", err)
	}

	return err
}

func (c *MTProtoIntermediateCodec) Close() error {
	return c.conn.Close()
}
