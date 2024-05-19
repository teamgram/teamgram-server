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
// There is an abridged version of the same protocol:
// if the client sends 0xef as the first byte (**important:** only prior to the very first data packet),
// then packet length is encoded by a single byte (0x01..0x7e = data length divided by 4;
// or 0x7f followed by 3 length bytes (little endian) divided by 4) followed
// by the data themselves (sequence number and CRC32 not added).
// In this case, server responses look the same (the server does not send 0xefas the first byte).
//

type AbridgedCodec struct {
	*AesCTR128Crypto
	state     int
	packetLen [4]byte
}

func newMTProtoAbridgedCodec(crypto *AesCTR128Crypto) *AbridgedCodec {
	return &AbridgedCodec{
		AesCTR128Crypto: crypto,
		state:           WAIT_PACKET_LENGTH_1,
	}
}

// Encode encodes frames upon server responses into TCP stream.
func (c *AbridgedCodec) Encode(conn CodecWriter, msg interface{}) ([]byte, error) {
	if msg == nil {
		//logx.Error("conn(%s) msg is nil", conn)
		return nil, nil
	}

	rm, ok := msg.(*mtproto.MTPRawMessage)
	if !ok {
		err := fmt.Errorf("conn(%s) msg type error, only MTPRawMessage, msg: %s", conn, rm)
		return nil, err
	} else if rm == nil {
		// logx.Errorf("conn(%s) msg is nil, msg: %#v", conn, msg)
		return nil, nil
	}

	out := rm.Payload

	// b := message.Encode() d
	sb := make([]byte, 4)
	// minus padding
	size := len(out) / 4

	if size < 127 {
		sb = []byte{byte(size)}
	} else {
		binary.LittleEndian.PutUint32(sb, uint32(size<<8|127))
	}

	buf := append(sb, out...)
	return c.Encrypt(buf), nil
}

// Decode decodes frames from TCP stream via specific implementation.
func (c *AbridgedCodec) Decode(conn CodecReader) (interface{}, error) {
	var (
		in  innerBuffer
		buf []byte
		n   int
		err error
	)

	in, _ = conn.Peek(-1)
	// log.Debugf("connId: %d, n = %d", conn.ConnID(), len(in))
	if len(in) == 0 {
		return nil, nil
	}

	switch c.state {
	case WAIT_PACKET_LENGTH_1:
		if buf, err = in.readN(1); err != nil {
			return nil, ErrUnexpectedEOF
		}
		buf = c.Decrypt(buf)
		c.packetLen[0] = buf[0]
		conn.Discard(1)

		needAck := c.packetLen[0]>>7 == 1
		_ = needAck

		n = int(c.packetLen[0] & 0x7f)
		if n < 0x7f {
			c.state = WAIT_PACKET_LENGTH_1_PACKET
			n = n << 2
			// log.Debugf("n = %d", n)
		} else {
			c.state = WAIT_PACKET_LENGTH_3
			if buf, err = in.readN(3); err != nil {
				return nil, ErrUnexpectedEOF
			}
			buf = c.Decrypt(buf)
			c.packetLen[1] = buf[0]
			c.packetLen[2] = buf[1]
			c.packetLen[3] = buf[2]
			conn.Discard(3)

			c.state = WAIT_PACKET_LENGTH_3_PACKET
			n = (int(c.packetLen[1]) | int(c.packetLen[2])<<8 | int(c.packetLen[3])<<16) << 2
			// log.Debugf("n = %d", n)
			if n > MAX_MTPRORO_FRAME_SIZE {
				// TODO(@benqi): close conn
				return nil, fmt.Errorf("too large data(%d)", n)
			}
		}
		if buf, err = in.readN(n); err != nil {
			return nil, ErrUnexpectedEOF
		} else if len(buf) <= 4 {
			// TODO: fix
			return nil, ErrUnexpectedEOF
		}

		buf = c.Decrypt(buf)
		conn.Discard(n)
		c.state = WAIT_PACKET_LENGTH_1

		message := mtproto.NewMTPRawMessage(int64(binary.LittleEndian.Uint64(buf)), 0, TRANSPORT_TCP)
		message.Decode(buf)

		return message, nil
	case WAIT_PACKET_LENGTH_1_PACKET:
		n = int(c.packetLen[0]&0x7f) << 2
		if buf, err = in.readN(n); err != nil {
			return nil, ErrUnexpectedEOF
		}
		// log.Debugf("n = %d", n)

		buf = c.Decrypt(buf)
		conn.Discard(n)
		c.state = WAIT_PACKET_LENGTH_1

		message := mtproto.NewMTPRawMessage(int64(binary.LittleEndian.Uint64(buf)), 0, TRANSPORT_TCP)
		message.Decode(buf)

		return message, nil
	case WAIT_PACKET_LENGTH_3:
		if buf, err = in.readN(3); err != nil {
			return nil, ErrUnexpectedEOF
		}
		buf = c.Decrypt(buf)
		c.packetLen[1] = buf[0]
		c.packetLen[2] = buf[1]
		c.packetLen[3] = buf[2]
		conn.Discard(3)

		c.state = WAIT_PACKET_LENGTH_3_PACKET
		n = (int(c.packetLen[1]) | int(c.packetLen[2])<<8 | int(c.packetLen[3])<<16) << 2
		// log.Debugf("n = %d", n)
		if n > MAX_MTPRORO_FRAME_SIZE {
			// TODO(@benqi): close conn
			return nil, fmt.Errorf("too large data(%d)", n)
		}
		if buf, err = in.readN(n); err != nil {
			return nil, ErrUnexpectedEOF
		}

		buf = c.Decrypt(buf)
		conn.Discard(n)
		c.state = WAIT_PACKET_LENGTH_1

		message := mtproto.NewMTPRawMessage(int64(binary.LittleEndian.Uint64(buf)), 0, TRANSPORT_TCP)
		message.Decode(buf)

		return message, nil
	case WAIT_PACKET_LENGTH_3_PACKET:
		n = (int(c.packetLen[1]) | int(c.packetLen[2])<<8 | int(c.packetLen[3])<<16) << 2
		// log.Debugf("n = %d", n)
		if n > MAX_MTPRORO_FRAME_SIZE {
			// TODO(@benqi): close conn
			return nil, fmt.Errorf("too large data(%d)", n)
		}
		if buf, err = in.readN(n); err != nil {
			return nil, ErrUnexpectedEOF
		}

		buf = c.Decrypt(buf)
		conn.Discard(n)
		c.state = WAIT_PACKET_LENGTH_1

		message := mtproto.NewMTPRawMessage(int64(binary.LittleEndian.Uint64(buf)), 0, TRANSPORT_TCP)
		message.Decode(buf)

		return message, nil
	}

	// TODO(@benqi): close conn
	return nil, fmt.Errorf("unknown error")
}
