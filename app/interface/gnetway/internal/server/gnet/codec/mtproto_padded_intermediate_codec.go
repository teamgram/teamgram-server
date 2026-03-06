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
	"crypto/rand"
	"encoding/binary"
	"fmt"
	rand2 "math/rand"

	"github.com/teamgram/proto/mtproto"
)

// PaddedIntermediateCodec implements the MTProto padded intermediate transport.
// Compared to intermediate, it allows random padding bytes after the payload
// to make traffic analysis more difficult.
type PaddedIntermediateCodec struct {
	*AesCTR128Crypto
	state     int
	packetLen uint32
}

func newMTProtoPaddedIntermediateCodec(crypto *AesCTR128Crypto) *PaddedIntermediateCodec {
	return &PaddedIntermediateCodec{
		AesCTR128Crypto: crypto,
		state:           WAIT_PACKET_LENGTH,
	}
}

func generatePadding(size int) []byte {
	b := make([]byte, size)
	_, _ = rand.Read(b)
	return b
}

// Encode encodes frames upon server responses into TCP stream.
func (c *PaddedIntermediateCodec) Encode(conn CodecWriter, msg interface{}) ([]byte, error) {
	rawMsg, ok := msg.(*mtproto.MTPRawMessage)
	if !ok {
		err := fmt.Errorf("msg type error, only MTPRawMessage, msg: {%v}", msg)
		return nil, err
	}

	paddingLen := rand2.Int() % 16
	buf := make([]byte, 4+len(rawMsg.Payload)+paddingLen)
	binary.LittleEndian.PutUint32(buf, uint32(len(rawMsg.Payload)))
	copy(buf[4:], rawMsg.Payload)
	if paddingLen > 0 {
		_, _ = rand.Read(buf[4+len(rawMsg.Payload):])
	}

	return c.Encrypt(buf), nil
}

// EncodeQuickAck encodes the Quick ACK token for the padded intermediate transport.
func (c *PaddedIntermediateCodec) EncodeQuickAck(token uint32) []byte {
	var buf [4]byte
	binary.LittleEndian.PutUint32(buf[:], token)
	return c.Encrypt(buf[:])
}

// Decode decodes frames from TCP stream via specific implementation.
func (c *PaddedIntermediateCodec) Decode(conn CodecReader) (bool, []byte, error) {
	var (
		buf     []byte
		n       int
		in      innerBuffer
		err     error
		needAck bool
	)

	in, _ = conn.Peek(-1)

	if c.state == WAIT_PACKET_LENGTH {
		if buf, err = in.readN(4); err != nil {
			return false, nil, ErrUnexpectedEOF
		}
		_, _ = conn.Discard(4)
		buf = c.Decrypt(buf)
		c.packetLen = binary.LittleEndian.Uint32(buf)
		c.state = WAIT_PACKET
	}

	needAck = c.packetLen>>31 == 1
	n = int(c.packetLen & 0x7fffffff)
	if n <= 0 || n%4 != 0 {
		return false, nil, ErrProtoBadLength
	}
	if n > MAX_MTPRORO_FRAME_SIZE {
		return false, nil, fmt.Errorf("%w: too large data(%d)", ErrProtoBadLength, n)
	}

	if buf, err = in.readN(n); err != nil {
		return false, nil, ErrUnexpectedEOF
	}
	buf = c.Decrypt(buf)
	_, _ = conn.Discard(n)
	c.state = WAIT_PACKET_LENGTH

	return needAck, buf, nil
}
