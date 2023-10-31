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
	"math/rand"

	"github.com/teamgram/proto/mtproto"
	"github.com/teamgram/proto/mtproto/crypto"

	log "github.com/zeromicro/go-zero/core/logx"
)

// PaddedIntermediateCodec
// https://core.telegram.org/mtproto#tcp-transport
//
// In case 4-byte data alignment is needed,
// an intermediate version of the original protocol may be used:
// if the client sends 0xeeeeeeee as the first int (four bytes),
// then packet length is encoded always by four bytes as in the original version,
// but the sequence number and CRC32 are omitted,
// thus decreasing total packet size by 8 bytes.
type PaddedIntermediateCodec struct {
	conn io.ReadWriteCloser
}

func NewMTProtoPaddedIntermediateCodec(conn io.ReadWriteCloser) *PaddedIntermediateCodec {
	return &PaddedIntermediateCodec{
		conn: conn,
	}
}

func (c *PaddedIntermediateCodec) Receive() (interface{}, error) {
	var size int
	var n int
	var err error

	b := make([]byte, 4)
	n, err = io.ReadFull(c.conn, b)
	if err != nil {
		return nil, err
	}

	size = int(binary.LittleEndian.Uint32(b))
	log.Infof("size1: %d", size)

	left := size
	buf := make([]byte, size)
	for left > 0 {
		n, err = io.ReadFull(c.conn, buf[size-left:])
		if err != nil {
			log.Errorf("readFull2 error: %v", err)
			return nil, err
		}
		left -= n
	}

	if size > 4096 {
		log.Infof("ReadFull2: %s", hex.EncodeToString(buf[:256]))
	}

	// TODO(@benqi): process report ack and quickack
	// 截断QuickAck消息，客户端有问题
	if size == 4 {
		log.Errorf("Server response error: ", int32(binary.LittleEndian.Uint32(buf)))
		// return nil, fmt.Errorf("Recv QuickAckMessage, ignore!!!!") //  connId: ", c.stream, ", by client ", m.RemoteAddr())
		return nil, nil
	}

	authKeyId := int64(binary.LittleEndian.Uint64(buf))
	message := mtproto.NewMTPRawMessage(authKeyId, 0, TRANSPORT_TCP)
	message.Decode(buf)
	return message, nil
}

func (c *PaddedIntermediateCodec) Send(msg interface{}) error {
	message, ok := msg.(*mtproto.MTPRawMessage)
	if !ok {
		err := fmt.Errorf("msg type error, only MTPRawMessage, msg: {%v}", msg)
		log.Error(err.Error())
		return err
	}

	b := message.Encode()

	sb := make([]byte, 4)
	size := len(b)

	binary.LittleEndian.PutUint32(sb, uint32(size))
	b = append(sb, b...)
	b = append(b, crypto.GenerateNonce(int(rand.Uint32()%16))...)

	_, err := c.conn.Write(b)

	if err != nil {
		log.Errorf("Send msg error: %s", err)
	}

	return err
}

func (c *PaddedIntermediateCodec) Close() error {
	return c.conn.Close()
}

func (c *PaddedIntermediateCodec) Context() interface{} {
	return ""
}
