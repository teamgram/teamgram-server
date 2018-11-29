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
	"encoding/hex"
	"fmt"
	"github.com/golang/glog"
	"github.com/nebula-chat/chatengine/pkg/crypto"
	"github.com/nebula-chat/chatengine/pkg/net2"
	"io"
)

type MTProtoAppCodec struct {
	// conn *net.TCPConn
	stream *AesCTR128Stream
}

func NewMTProtoAppCodec(conn *net2.BufferedConn, d *crypto.AesCTR128Encrypt, e *crypto.AesCTR128Encrypt) *MTProtoAppCodec {
	return &MTProtoAppCodec{
		// conn:   conn,
		stream: NewAesCTR128Stream(conn, d, e),
	}
}

func (c *MTProtoAppCodec) Receive() (interface{}, error) {
	var size int
	var n int
	var err error

	b := make([]byte, 1)
	n, err = io.ReadFull(c.stream, b)
	if err != nil {
		return nil, err
	}

	// TODO(@benqi): dispatch to session
	// glog.Info("first_byte: ", hex.EncodeToString(b[:1]))
	needAck := bool(b[0]>>7 == 1)
	_ = needAck

	b[0] = b[0] & 0x7f
	// glog.Info("first_byte2: ", hex.EncodeToString(b[:1]))

	if b[0] < 0x7f {
		size = int(b[0]) << 2
		glog.Info("size1: ", size)
		if size == 0 {
			return nil, nil
		}
	} else {
		glog.Info("first_byte2: ", hex.EncodeToString(b[:1]))
		b2 := make([]byte, 3)
		n, err = io.ReadFull(c.stream, b2)
		if err != nil {
			return nil, err
		}
		size = (int(b2[0]) | int(b2[1])<<8 | int(b2[2])<<16) << 2
		glog.Info("size2: ", size)
	}

	left := size
	buf := make([]byte, size)
	for left > 0 {
		n, err = io.ReadFull(c.stream, buf[size-left:])
		if err != nil {
			glog.Error("ReadFull2 error: ", err)
			return nil, err
		}
		left -= n
	}
	if size > 10240 {
		glog.Info("ReadFull2: ", hex.EncodeToString(buf[:256]))
	}

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

	//var message MessageBase
	//if authKeyId == 0 {
	//	message = NewUnencryptedRawMessage()
	//	// message.Decode(buf[8:])
	//} else {
	//	message = NewEncryptedRawMessage(authKeyId)
	//}
	//
	//err = message.Decode(buf[8:])
	//if err != nil {
	//	glog.Errorf("decode message error: {%v}", err)
	//	return nil, err
	//}
	//
	// return message, nil
}

func (c *MTProtoAppCodec) Send(msg interface{}) error {
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

	if size < 127 {
		sb = []byte{byte(size)}
	} else {
		binary.LittleEndian.PutUint32(sb, uint32(size<<8|127))
	}

	b = append(sb, b...)
	_, err := c.stream.Write(b)

	if err != nil {
		glog.Errorf("Send msg error: %s", err)
	}

	return err
}

func (c *MTProtoAppCodec) Close() error {
	return c.stream.conn.Close()
}

type AesCTR128Stream struct {
	conn    *net2.BufferedConn
	encrypt *crypto.AesCTR128Encrypt
	decrypt *crypto.AesCTR128Encrypt
}

func NewAesCTR128Stream(conn *net2.BufferedConn, d *crypto.AesCTR128Encrypt, e *crypto.AesCTR128Encrypt) *AesCTR128Stream {
	return &AesCTR128Stream{
		conn:    conn,
		decrypt: d,
		encrypt: e,
	}
}

func (this *AesCTR128Stream) Read(p []byte) (int, error) {
	n, err := this.conn.Read(p)
	if err == nil {
		this.decrypt.Encrypt(p[:n])
		return n, nil
	}
	return n, err
}

func (this *AesCTR128Stream) Write(p []byte) (int, error) {
	this.encrypt.Encrypt(p[:])
	return this.conn.Write(p)
}
