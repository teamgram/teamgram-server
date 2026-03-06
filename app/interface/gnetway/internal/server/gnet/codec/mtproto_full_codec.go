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
	"hash/crc32"

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

// FullCodec implements the MTProto full TCP transport.
// As per https://core.telegram.org/mtproto#tcp-transport, each packet has
// length(4) + seqno(4) + payload + crc32(4), where length covers the whole
// packet, seqno is a per‑connection sequence number, and crc32 validates
// length+seqno+payload.
//
// In this implementation FullCodec runs over plain TCP and is mainly used
// for compatibility with clients that only support the full transport.
type FullCodec struct {
	recvSeqNo int32
	sendSeqNo int32
}

func newMTProtoFullCodec() *FullCodec {
	return new(FullCodec)
}

// Encode encodes frames upon server responses into TCP stream.
func (c *FullCodec) Encode(conn CodecWriter, msg interface{}) ([]byte, error) {
	rawMsg, ok := msg.(*mtproto.MTPRawMessage)
	if !ok {
		err := fmt.Errorf("msg type error, only MTPRawMessage, msg: {%v}", msg)
		return nil, err
	}

	payload := rawMsg.Payload
	totalLen := 4 + 4 + len(payload) + 4 // length + seqno + payload + crc32

	buf := make([]byte, totalLen)
	binary.LittleEndian.PutUint32(buf, uint32(totalLen))
	binary.LittleEndian.PutUint32(buf[4:], uint32(c.sendSeqNo))
	c.sendSeqNo++
	copy(buf[8:], payload)
	// CRC32 of (length + seqno + payload)
	checksum := crc32.ChecksumIEEE(buf[:8+len(payload)])
	binary.LittleEndian.PutUint32(buf[8+len(payload):], checksum)

	return buf, nil
}

// EncodeQuickAck returns nil: the full transport does not support Quick ACK.
func (c *FullCodec) EncodeQuickAck(_ uint32) []byte { return nil }

// Decode decodes frames from TCP stream via specific implementation.
func (c *FullCodec) Decode(conn CodecReader) (bool, []byte, error) {
	var (
		in  innerBuffer
		buf []byte
		err error
	)

	in, _ = conn.Peek(-1)

	// Read 4-byte length header
	if buf, err = in.readN(4); err != nil {
		return false, nil, ErrUnexpectedEOF
	}

	totalLen := int(binary.LittleEndian.Uint32(buf))
	// Minimum: 4 (length) + 4 (seqno) + 4 (crc32) = 12
	if totalLen < 12 {
		return false, nil, fmt.Errorf("%w: full codec: invalid total length: %d", ErrProtoBadLength, totalLen)
	}
	if totalLen > MAX_MTPRORO_FRAME_SIZE {
		return false, nil, fmt.Errorf("%w: full codec: too large data(%d)", ErrProtoBadLength, totalLen)
	}

	// Read remaining bytes: seqno + payload + crc32
	remainLen := totalLen - 4
	if buf, err = in.readN(remainLen); err != nil {
		return false, nil, ErrUnexpectedEOF
	}
	_, _ = conn.Discard(totalLen)

	// Validate CRC32: checksum covers length(4) + seqno(4) + payload
	payloadEnd := len(buf) - 4
	recvCrc := binary.LittleEndian.Uint32(buf[payloadEnd:])
	// Build the data to checksum: length header + seqno + payload
	var lenBuf [4]byte
	binary.LittleEndian.PutUint32(lenBuf[:], uint32(totalLen))
	h := crc32.NewIEEE()
	_, _ = h.Write(lenBuf[:])
	_, _ = h.Write(buf[:payloadEnd])
	calcCrc := h.Sum32()
	if recvCrc != calcCrc {
		return false, nil, fmt.Errorf("%w: full codec: crc32 mismatch: received 0x%08x, calculated 0x%08x", ErrProtoBadCRC, recvCrc, calcCrc)
	}

	// Validate sequence number
	seq := int32(binary.LittleEndian.Uint32(buf[:4]))
	if seq != c.recvSeqNo {
		return false, nil, fmt.Errorf("%w: full codec: seq mismatch: received %d, expected %d", ErrProtoBadSeq, seq, c.recvSeqNo)
	}
	c.recvSeqNo++

	// Payload is between seqno and crc32
	return false, buf[4:payloadEnd], nil
}
