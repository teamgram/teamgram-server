// Copyright 2022 Teamgram Authors
//  All rights reserved.
//
// Author: Benqi (wubenqi@gmail.com)
//

package codec

import (
	"encoding/binary"
	"encoding/hex"
	"errors"
	"fmt"

	"github.com/teamgram/proto/mtproto/crypto"

	"github.com/zeromicro/go-zero/core/logx"
)

// TODO(@benqi): Quick ack (https://core.telegram.org/mtproto#tcp-transport)
//
// The full, the intermediate and the abridged versions of the protocol have support for quick acknowledgment.
// In this case, the client sets the highest-order length bit in the query packet,
// and the server responds with a special 4 bytes as a separate packet.
// They are the 32 higher-order bits of SHA256 of the encrypted
// portion of the packet prepended by 32 bytes from the authorization key
// (the same hash as computed for verifying the message key),
// with the most significant bit set to make clear that this is not the length of a regular server response packet;
// if the abridged version is used, bswap is applied to these four bytes.
//

// Transport类型，不支持UDP
const (
	TRANSPORT_TCP  = 1 // TCP
	TRANSPORT_HTTP = 2 // HTTP
	TRANSPORT_UDP  = 3 // UDP, TODO(@benqi): 未发现有支持UDP的客户端
)

const (
	// ABRIDGED_FLAG -- FULL_FLAG
	// Tcp Transport
	ABRIDGED_FLAG            = 0xef
	ABRIDGED_INT32_FLAG      = 0xefefefef
	INTERMEDIATE_FLAG        = 0xeeeeeeee
	PADDED_INTERMEDIATE_FLAG = 0xdddddddd
	UNKNOWN_FLAG             = 0x02010316
	PVRG_FLAG                = 0x47725650 // PVrG
	FULL_FLAG                = 0x00000000

	// HTTP_HEAD_FLAG -- HTTP_OPTION_FLAG
	// Http Transport
	HTTP_HEAD_FLAG   = 0x44414548 // HEAD
	HTTP_POST_FLAG   = 0x54534f50 // POST
	HTTP_GET_FLAG    = 0x20544547 // GET
	HTTP_OPTION_FLAG = 0x4954504f // OPTION

	// 3d9ff4f1
)

var (
	// ErrIncompletePacket occurs when there is an incomplete packet under TCP protocol.
	ErrIncompletePacket = errors.New("incomplete packet")
	// ErrInvalidFixedLength occurs when the output data have invalid fixed length.
	ErrInvalidFixedLength = errors.New("invalid fixed length of bytes")
	// ErrUnexpectedEOF occurs when no enough data to read by codec.
	ErrUnexpectedEOF = errors.New("there is no enough data")
	// ErrUnsupportedLength occurs when unsupported lengthFieldLength is from input data.
	ErrUnsupportedLength = errors.New("unsupported lengthFieldLength. (expected: 1, 2, 3, 4, or 8)")
	// ErrTooLessLength occurs when adjusted frame length is less than zero.
	ErrTooLessLength = errors.New("adjusted frame length is less than zero")
	// ErrHttpTransport occurs when there is http transport protocol.
	ErrHttpTransport = errors.New("there is http transport protocol")
	// ErrPvrgNotSupport occurs when there is PVrG transport protocol.
	ErrPvrgNotSupport = errors.New("PVrG transport not support")
	// Err0x02010316NotSupport occurs when there is 0x02010316 transport protocol.
	Err0x02010316NotSupport = errors.New("0x02010316 transport not support")
)

// var ErrShortBuffer = io.ErrShortBuffer

type CodecReader interface {
	// Peek returns the next n bytes without advancing the reader. The bytes stop
	// being valid at the next read call. If Peek returns fewer than n bytes, it
	// also returns an error explaining why the read is short. The error is
	// ErrBufferFull if n is larger than b's buffer size.
	//
	// Note that the []byte buf returned by Peek() is not allowed to be passed to a new goroutine,
	// as this []byte will be reused within event-loop.
	// If you have to use buf in a new goroutine, then you need to make a copy of buf and pass this copy
	// to that new goroutine.
	Peek(n int) (buf []byte, err error)

	// Discard skips the next n bytes, returning the number of bytes discarded.
	//
	// If Discard skips fewer than n bytes, it also returns an error.
	// If 0 <= n <= b.Buffered(), Discard is guaranteed to succeed without
	// reading from the underlying io.Reader.
	Discard(n int) (discarded int, err error)
}

type CodecWriter interface {
}

type Codec interface {
	Encode(conn CodecWriter, msg interface{}) ([]byte, error)
	Decode(conn CodecReader) (interface{}, error)
	// FirstBytes() int
}

func CreateMTProtoCodec(conn CodecReader) (Codec, error) {
	var (
		firstByte uint8
		err       error
	)

	bytes, _ := conn.Peek(1)
	firstByte = bytes[0]

	if firstByte == ABRIDGED_FLAG {
		logx.Debugf("conn(%s) mtproto abridged version.", conn)
		conn.Discard(1)
		return newMTProtoAbridgedCodec(nil), nil
	}

	var (
		firstInt uint32
	)
	// not abridged version, we'll lookup codec!
	bytes, err = conn.Peek(4)
	if err != nil {
		return nil, ErrUnexpectedEOF
	}

	firstInt = binary.LittleEndian.Uint32(bytes)

	// check http
	if firstInt == HTTP_HEAD_FLAG ||
		firstInt == HTTP_POST_FLAG ||
		firstInt == HTTP_GET_FLAG ||
		firstInt == HTTP_OPTION_FLAG {
		// http 协议
		// log.Debugf("mtproto http.")

		// conn2 := NewMTProtoHttpProxyConn(conn)
		// c.conn = conn2
		// c.codecType = TRANSPORT_HTTP
		logx.Debugf("conn(%s) mtproto http.", conn)
		return nil, ErrHttpTransport
	}

	// check intermediate version
	if firstInt == INTERMEDIATE_FLAG {
		logx.Debugf("conn(%s) intermediate version.", conn)
		conn.Discard(4)
		return newMTProtoIntermediateCodec(nil), nil
	}

	// check intermediate version
	if firstInt == PADDED_INTERMEDIATE_FLAG {
		logx.Debugf("conn(%s) padded intermediate version.", conn)
		conn.Discard(4)
		return newMTProtoPaddedIntermediateCodec(nil), nil
	}

	// check PVrG
	if firstInt == PVRG_FLAG {
		logx.Errorf("conn(%s) PVrG version.", conn)
		return nil, ErrPvrgNotSupport
	}

	// check 0x02010316
	if firstInt == UNKNOWN_FLAG {
		logx.Errorf("conn(%s) firstInt is 0x02010316.", conn)
		return nil, Err0x02010316NotSupport
	}

	var (
		checkFullBuf []byte
	)

	if bytes, err = conn.Peek(12); err != nil {
		return nil, ErrUnexpectedEOF
	} else {
		checkFullBuf = bytes
	}

	secondInt := binary.BigEndian.Uint32(checkFullBuf[4:])
	if secondInt == FULL_FLAG {
		logx.Infof("conn(%s) mtproto full version.", conn)
		// conn.Discard(12)
		return newMTProtoFullCodec(), nil
	}

	// 5. app version.

	// bytes
	// |  0-3  |  4-7   |    8-55    |     56-59    | 60-63 |
	// |  val  |  val2  |            | 0xefefefefef |       |
	//
	// temp
	// |    0 ~ 47       |
	// | 55 ~ 8 (bytes)  |
	//
	// encrypt_key_: 8  ~ 39 (btes)
	// encrypt_iv_ : 40 ~ 55 (bytes)
	// decrypt_key_: 0  ~ 31 (temp)
	// decrypt_iv_ : 32 ~ 47 (temp)
	//

	var (
		obfuscatedBuf []byte
	)

	if bytes, err = conn.Peek(64); err != nil {
		return nil, ErrUnexpectedEOF
	} else {
		obfuscatedBuf = bytes
	}

	var (
		tmp [64]byte
	)

	// 生成decrypt_key
	for i := 0; i < 48; i++ {
		tmp[i] = obfuscatedBuf[55-i]
	}

	e, err := crypto.NewAesCTR128Encrypt(tmp[:32], tmp[32:48])
	if err != nil {
		return nil, err
	}

	d, err := crypto.NewAesCTR128Encrypt(obfuscatedBuf[8:40], obfuscatedBuf[40:56])
	if err != nil {
		return nil, err
	}

	d.Encrypt(obfuscatedBuf)

	protocolType := binary.BigEndian.Uint32(obfuscatedBuf[56:])
	if protocolType != ABRIDGED_INT32_FLAG &&
		protocolType != INTERMEDIATE_FLAG &&
		protocolType != PADDED_INTERMEDIATE_FLAG {
		return nil, fmt.Errorf("conn(%s) mtproto buf[56:60]'s byte != 0xef, received: %s",
			conn,
			hex.EncodeToString(obfuscatedBuf[56:60]))
	}

	dcId := int16(binary.BigEndian.Uint16(obfuscatedBuf[60:]))
	// TODO: check dcId

	conn.Discard(64)

	logx.Infof("conn(%s) mtproto obfuscated version, {protocol_type: %d, dc_id: %d}", conn, protocolType, dcId)
	return newMTProtoObfuscatedCodec(d, e, protocolType, dcId), nil
}
