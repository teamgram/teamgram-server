// Copyright (c) 2021-present,  Teamgram Studio (https://teamgram.io).
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

package codec

import (
	"encoding/binary"
	"fmt"

	"github.com/teamgram/proto/mtproto"
)

// https://core.telegram.org/mtproto#tcp-transport
//
// If a payload (packet) needs to be transmitted from server to client or from client to server,
// it is encapsulated as follows:
// 4 length bytes are added at the front
// (to include the length, the sequence number, and CRC32; always divisible by 4)
// and 4 bytes with the packet sequence number within this TCP connection
// (the first packet sent is numbered 0, the next one 1, etc.),
// and 4 CRC32 bytes at the end (length, sequence number, and payload together).
//

// FullCodec FullCodec
type FullCodec struct {
}

func newMTProtoFullCodec() *FullCodec {
	return new(FullCodec)
}

// Encode encodes frames upon server responses into TCP stream.
func (c *FullCodec) Encode(conn CodecWriter, msg interface{}) ([]byte, error) {
	rawMsg, ok := msg.(*mtproto.MTPRawMessage)
	if !ok {
		err := fmt.Errorf("msg type error, only MTPRawMessage, msg: {%v}", msg)
		// log.Error(err.Error())
		return nil, err
	}

	b := rawMsg.Payload

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

	return b, nil
}

// Decode decodes frames from TCP stream via specific implementation.
func (c *FullCodec) Decode(conn CodecReader) (interface{}, error) {
	var (
		size int
		buf  []byte
		n    int
		in   innerBuffer
		err  error
	)

	in, _ = conn.Peek(-1)

	if buf, err = in.readN(4); err != nil {
		return nil, ErrUnexpectedEOF
	}
	size += 4

	n = int(binary.LittleEndian.Uint32(buf))
	// Check bufLen
	if n < 12 {
		err = fmt.Errorf("invalid len: %d", size)
		return nil, err
	}

	if buf, err = in.readN(n); err != nil {
		return nil, ErrUnexpectedEOF
	}
	size += n
	conn.Discard(size)

	seq := binary.LittleEndian.Uint32(buf[:4])
	// TODO(@benqi): check seqNum, save last seq_num
	_ = seq

	crc32 := binary.LittleEndian.Uint32(buf[len(buf)-4:])
	// TODO(@benqi): check crc32
	_ = crc32

	message := mtproto.NewMTPRawMessage(int64(binary.LittleEndian.Uint64(buf[4:])), 0, TRANSPORT_TCP)
	message.Decode(buf)

	return message, nil
}
