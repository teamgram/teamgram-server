// Copyright 2022 Teamgram Authors
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
//
// Author: teamgramio (teamgram.io@gmail.com)
//

package codec

import (
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"io"

	"github.com/teamgram/proto/mtproto"

	log "github.com/zeromicro/go-zero/core/logx"
)

// FullCodec
// https://core.telegram.org/mtproto#tcp-transport
//
// If a payload (packet) needs to be transmitted from server to client or from client to server,
// it is encapsulated as follows:
// 4 length bytes are added at the front
// (to include the length, the sequence number, and CRC32; always divisible by 4)
// and 4 bytes with the packet sequence number within this TCP connection
// (the first packet sent is numbered 0, the next one 1, etc.),
// and 4 CRC32 bytes at the end (length, sequence number, and payload together).
type FullCodec struct {
	conn io.ReadWriteCloser
}

func NewMTProtoFullCodec(conn io.ReadWriteCloser) *FullCodec {
	return &FullCodec{
		conn: conn,
	}
}

func (c *FullCodec) Receive() (interface{}, error) {
	var size int
	var n int
	var err error

	b := make([]byte, 4)
	n, err = io.ReadFull(c.conn, b)
	if err != nil {
		return nil, err
	}

	size = int(binary.LittleEndian.Uint32(b))
	// Check bufLen
	if size < 12 {
		err = fmt.Errorf("invalid len: %d", size)
		return nil, err
	}

	log.Infof("size1: %d", size)

	left := size
	buf := make([]byte, size-4)
	for left > 0 {
		n, err = io.ReadFull(c.conn, buf[size-left:])
		if err != nil {
			log.Errorf("ReadFull2 error: %v", err)
			return nil, err
		}
		left -= n
	}

	if size > 4096 {
		log.Infof("readFull2: %s", hex.EncodeToString(buf[:256]))
	}

	seqNum := binary.LittleEndian.Uint32(buf[:4])
	// TODO(@benqi): check seqNum, save last seq_num
	_ = seqNum

	crc32 := binary.LittleEndian.Uint32(buf[len(buf)-4:])
	// TODO(@benqi): check crc32
	_ = crc32

	authKeyId := int64(binary.LittleEndian.Uint64(buf[4:]))
	message := mtproto.NewMTPRawMessage(authKeyId, 0, TRANSPORT_TCP)
	message.Decode(buf)
	return message, nil
}

func (c *FullCodec) Send(msg interface{}) error {
	message, ok := msg.(*mtproto.MTPRawMessage)
	if !ok {
		err := fmt.Errorf("msg type error, only MTPRawMessage, msg: {%v}", msg)
		log.Error(err.Error())
		return err
	}

	b := message.Encode()

	sb := make([]byte, 8)
	// minus padding
	size := len(b) / 4

	binary.LittleEndian.PutUint32(sb, uint32(size))
	// TODO(@benqi): gen seq_num
	var seqNum uint32 = 0
	binary.LittleEndian.PutUint32(sb[4:], seqNum)
	b = append(sb, b...)
	var crc32Buf = make([]byte, 4)
	var crc32 uint32 = 0
	binary.LittleEndian.PutUint32(crc32Buf, crc32)
	b = append(sb, crc32Buf...)

	_, err := c.conn.Write(b)
	if err != nil {
		log.Errorf("Send msg error: %s", err)
	}

	return err
}

func (c *FullCodec) Close() error {
	return c.conn.Close()
}

func (c *FullCodec) Context() interface{} {
	return ""
}
