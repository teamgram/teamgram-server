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
	"flag"
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
	isClientType = false
)

const (
	ERROR                       = -1
	INVALID                     = 0
	WAIT_FIRST_PACKET           = 1
	WAIT_PACKET_LENGTH_1        = 2
	WAIT_PACKET_LENGTH_1_PACKET = 3
	WAIT_PACKET_LENGTH_3        = 4
	WAIT_PACKET_LENGTH_3_PACKET = 5
	WAIT_PACKET_LENGTH          = 6
	WAIT_PACKET                 = 7
)

const (
	MAX_MTPRORO_FRAME_SIZE = 16777216
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

// Protocol-level errors that help callers distinguish between
// client-side protocol violations and internal server failures.
var (
	// ErrProtoBadMagic 表示首包 magic / obfuscated 协议头与预期不符。
	ErrProtoBadMagic = errors.New("mtproto: bad transport magic")
	// ErrProtoBadLength 表示长度字段非法（<=0、超出限制或不满足对齐要求）。
	ErrProtoBadLength = errors.New("mtproto: bad frame length")
	// ErrProtoBadCRC 表示 full transport 下 CRC32 校验失败。
	ErrProtoBadCRC = errors.New("mtproto: bad crc32")
	// ErrProtoBadSeq 表示 full transport 下 seqno 不连续。
	ErrProtoBadSeq = errors.New("mtproto: bad sequence number")
	// ErrProtoDecrypt 表示 AES 解密或 obfuscated 握手阶段出错。
	ErrProtoDecrypt = errors.New("mtproto: decrypt failed")
	// ErrTransportNotSupported 表示客户端选择了当前服务端未实现的 transport 方式。
	ErrTransportNotSupported = errors.New("mtproto: transport not supported")
)

var (
	isMTProto    bool // 是否使用MTProto - true为官方mtproto协议，false为定制协议（当前实现为ntproto）
	isObfuscated bool
)

func init() {
	flag.BoolVar(&isMTProto, "mtproto", true, "mtproto")
	flag.BoolVar(&isObfuscated, "obfuscated", true, "obfuscated")
}

// var ErrShortBuffer = io.ErrShortBuffer

// CodecReader exposes a read-only view of the underlying connection.
// Implementations are expected to reuse the underlying buffer between calls,
// so callers must not retain the returned slice beyond the next read.
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

// CodecWriter is a placeholder for future write-side helpers.
// Currently Encode returns a complete frame and the caller writes it to the connection.
type CodecWriter interface {
}

// Codec represents one MTProto TCP transport mode (abridged, intermediate,
// padded intermediate, full, or an obfuscated wrapper around them).
// A single Codec instance is created per connection during the handshake phase.
type Codec interface {
	Encode(conn CodecWriter, msg interface{}) ([]byte, error)
	Decode(conn CodecReader) (bool, []byte, error)
	// EncodeQuickAck encodes a 4-byte Quick ACK token for the transport.
	// For abridged transport, the token bytes are swapped (big-endian on the wire)
	// per the MTProto spec. For intermediate/padded-intermediate, the token is sent
	// as little-endian. In both cases the result is passed through the CTR cipher
	// (if the connection uses an obfuscated transport).
	// Returns nil if Quick ACK is not supported by the transport (e.g. full codec).
	EncodeQuickAck(token uint32) []byte
}

// CreateCodec chooses either the official MTProto transports or the custom
// ntproto variant based on global flags, and performs the initial handshake
// and transport detection on the underlying connection.
func CreateCodec(conn CodecReader) (Codec, error) {
	if isMTProto {
		return CreateMTProtoCodec(conn)
	} else {
		return CreateMyProtoCodec(conn)
	}
}

// CreateMTProtoCodec inspects the first bytes from the connection and selects
// the appropriate MTProto TCP transport (HTTP, full, abridged, intermediate,
// padded intermediate, or obfuscated). For obfuscated transports it also
// performs the 64‑byte handshake and derives per-connection AES‑CTR keys.
func CreateMTProtoCodec(conn CodecReader) (Codec, error) {
	rData, _ := conn.Peek(-1)
	if len(rData) == 0 {
		logx.Errorf("conn(%s) peek fail", conn)
		return nil, ErrUnexpectedEOF
	}

	if !isObfuscated {
		var (
			firstByte uint8
		)
		firstByte = rData[0]

		if firstByte == ABRIDGED_FLAG {
			logx.Debugf("conn(%s) mtproto abridged version, data: %s", conn, hex.EncodeToString(rData))
			_, _ = conn.Discard(1)
			return newMTProtoAbridgedCodec(nil), nil
		}
	}

	if len(rData) < 4 {
		logx.Errorf("conn(%s) peek bytes length < 4, data: %s", conn, hex.EncodeToString(rData))
		return nil, ErrUnexpectedEOF
	}

	// not abridged version, we'll lookup codec!
	// bytes = rData[:4]
	firstInt := binary.LittleEndian.Uint32(rData)

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
		logx.Errorf("conn(%s) mtproto http. data: %s", conn, hex.EncodeToString(rData))
		return nil, ErrHttpTransport
	}

	// check intermediate version
	if firstInt == INTERMEDIATE_FLAG {
		logx.Debugf("conn(%s) intermediate version. data: %s", conn, hex.EncodeToString(rData))
		_, _ = conn.Discard(4)
		return newMTProtoIntermediateCodec(nil), nil
	}

	// check intermediate version
	if firstInt == PADDED_INTERMEDIATE_FLAG {
		logx.Debugf("conn(%s) padded intermediate version. data: %s", conn, hex.EncodeToString(rData))
		_, _ = conn.Discard(4)
		return newMTProtoPaddedIntermediateCodec(nil), nil
	}

	// check PVrG (non‑standard transport, explicitly marked as unsupported)
	if firstInt == PVRG_FLAG {
		logx.Errorf("conn(%s) PVrG version. data: %s", conn, hex.EncodeToString(rData))
		return nil, fmt.Errorf("%w: PVrG version", ErrTransportNotSupported)
	}

	// check 0x02010316
	if firstInt == UNKNOWN_FLAG {
		logx.Errorf("conn(%s) firstInt is 0x02010316. data: %s", conn, hex.EncodeToString(rData))
		return nil, fmt.Errorf("%w: 0x02010316 version", ErrTransportNotSupported)
	}

	if len(rData) < 12 {
		logx.Errorf("conn(%s) peek bytes length < 12, data: %s", conn, hex.EncodeToString(rData))
		return nil, ErrUnexpectedEOF
	}

	checkFullBuf := rData[:12]

	secondInt := binary.BigEndian.Uint32(checkFullBuf[4:])
	if secondInt == FULL_FLAG {
		logx.Infof("conn(%s) mtproto full version. data: %s", conn, hex.EncodeToString(rData))
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

	if len(rData) < 64 {
		logx.Errorf("conn(%s) peek bytes length < 64, data: %s", conn, hex.EncodeToString(rData))
		return nil, ErrUnexpectedEOF
	}

	obfuscatedBuf := rData[:64]

	var (
		tmp [64]byte
	)

	// 生成decrypt_key
	for i := 0; i < 48; i++ {
		tmp[i] = obfuscatedBuf[55-i]
	}

	e, err := crypto.NewAesCTR128Encrypt(tmp[:32], tmp[32:48])
	if err != nil {
		return nil, fmt.Errorf("%w: create decrypt crypto failed", ErrProtoDecrypt)
	}

	d, err := crypto.NewAesCTR128Encrypt(obfuscatedBuf[8:40], obfuscatedBuf[40:56])
	if err != nil {
		return nil, fmt.Errorf("%w: create encrypt crypto failed", ErrProtoDecrypt)
	}

	d.Encrypt(obfuscatedBuf)

	protocolType := binary.BigEndian.Uint32(obfuscatedBuf[56:])
	if protocolType != ABRIDGED_INT32_FLAG &&
		protocolType != INTERMEDIATE_FLAG &&
		protocolType != PADDED_INTERMEDIATE_FLAG {
		return nil, fmt.Errorf("%w: conn(%s) mtproto buf[56:60] invalid, received: %s",
			ErrProtoBadMagic,
			conn,
			hex.EncodeToString(rData))
	}

	dcId := int16(binary.BigEndian.Uint16(obfuscatedBuf[60:]))
	if dcId == 0 {
		return nil, fmt.Errorf("%w: invalid dc id: %d", ErrProtoBadMagic, dcId)
	}

	//if secondInt == PROXY_FLAG {
	//	c.remoteIp = ip.IntToIP(firstInt)
	//}

	_, _ = conn.Discard(64)

	logx.Infof("conn(%s) mtproto obfuscated version, {protocol_type: %d, dc_id: %d}", conn, protocolType, dcId)
	return newMTProtoObfuscatedCodec(d, e, protocolType, dcId), nil
}

// CreateMyProtoCodec parses the custom ntproto obfuscated header.
// The layout is similar to MTProto obfuscated transports, but uses different
// byte offsets for key/iv derivation and for protocolType/dcId.
func CreateMyProtoCodec(conn CodecReader) (Codec, error) {
	// 5. app version.

	// bytes
	// |  0-3  |  4-7   |     8-11     | 12-15 |    16-63    |
	// |  val  |  val2  | 0xefefefefef |       |             |
	//
	// temp
	// |    0 ~ 47        |
	// | 63 ~ 16 (bytes)  |
	//
	// encrypt_key_: 16 ~ 47 (bytes)
	// encrypt_iv_ : 48 ~ 63 (bytes)
	// decrypt_key_: 0  ~ 31 (temp)
	// decrypt_iv_ : 32 ~ 47 (temp)
	//

	var (
		obfuscatedBuf []byte
	)

	bytes, err := conn.Peek(64)
	if err != nil {
		return nil, ErrUnexpectedEOF
	} else {
		obfuscatedBuf = bytes
	}

	var (
		tmp [64]byte
	)

	// 生成decrypt_key
	for i := 0; i < 48; i++ {
		// tmp[i] = obfuscatedBuf[55-i]
		tmp[i] = obfuscatedBuf[63-i]
	}

	e, err := crypto.NewAesCTR128Encrypt(tmp[:32], tmp[32:48])
	if err != nil {
		return nil, fmt.Errorf("%w: create decrypt crypto failed", ErrProtoDecrypt)
	}

	// d, err := crypto.NewAesCTR128Encrypt(obfuscatedBuf[8:40], obfuscatedBuf[40:56])
	d, err := crypto.NewAesCTR128Encrypt(obfuscatedBuf[16:48], obfuscatedBuf[48:64])
	if err != nil {
		return nil, fmt.Errorf("%w: create encrypt crypto failed", ErrProtoDecrypt)
	}

	d.Encrypt(obfuscatedBuf)

	// protocolType := binary.BigEndian.Uint32(obfuscatedBuf[56:])
	protocolType := binary.BigEndian.Uint32(obfuscatedBuf[8:])
	if protocolType != ABRIDGED_INT32_FLAG &&
		protocolType != INTERMEDIATE_FLAG &&
		protocolType != PADDED_INTERMEDIATE_FLAG {
		return nil, fmt.Errorf("%w: conn(%s) mtproto buf[8:12] invalid, received: %s",
			ErrProtoBadMagic,
			conn,
			hex.EncodeToString(obfuscatedBuf[8:12]))
	}

	// dcId := int16(binary.BigEndian.Uint16(obfuscatedBuf[60:]))
	dcId := int16(binary.BigEndian.Uint16(obfuscatedBuf[12:]))
	if dcId == 0 {
		return nil, fmt.Errorf("%w: invalid dc id: %d", ErrProtoBadMagic, dcId)
	}

	_, _ = conn.Discard(64)

	logx.Infof("conn(%s) mtproto obfuscated version, {protocol_type: %d, dc_id: %d}", conn, protocolType, dcId)
	return newMTProtoObfuscatedCodec(d, e, protocolType, dcId), nil
}
