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
	"encoding/hex"
	"errors"
	"fmt"

	"github.com/teamgram/proto/mtproto/crypto"

	"github.com/panjf2000/gnet"
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

//// Transport类型，不支持UDP
//const (
//	TRANSPORT_TCP  = 1 // TCP
//	TRANSPORT_HTTP = 2 // HTTP
//	TRANSPORT_UDP  = 3 // UDP, TODO(@benqi): 未发现有支持UDP的客户端
//)

const (
	// Tcp Transport
	ABRIDGED_FLAG            = 0xef
	ABRIDGED_INT32_FLAG      = 0xefefefef
	INTERMEDIATE_FLAG        = 0xeeeeeeee
	PADDED_INTERMEDIATE_FLAG = 0xdddddddd
	UNKNOWN_FLAG             = 0x02010316
	PVRG_FLAG                = 0x47725650 // PVrG
	FULL_FLAG                = 0x00000000

	// Http Transport
	HTTP_HEAD_FLAG   = 0x44414548 // HEAD
	HTTP_POST_FLAG   = 0x54534f50 // POST
	HTTP_GET_FLAG    = 0x20544547 // GET
	HTTP_OPTION_FLAG = 0x4954504f // OPTION

	// 3d9ff4f1
)

/////////////////////////////////////////////////////////////////////////////////////////////////////
type MTProtoCodec struct {
	codec       gnet.ICodec
	firstPacket bool
	isHttp      bool // codec type
}

func NewMTProtoCodec() *MTProtoCodec {
	return &MTProtoCodec{
		codec:       nil,
		firstPacket: true,
		isHttp:      false,
	}
}

func (c *MTProtoCodec) peekCodec(conn gnet.Conn) (gnet.ICodec, error) {
	var (
		firstByte uint8
	)

	if size, bytes := conn.ReadN(1); size == 1 {
		firstByte = bytes[0]
	} else {
		return nil, errUnexpectedEOF
	}

	if firstByte == ABRIDGED_FLAG {
		logx.Infof("conn(%s) mtproto abridged version.", conn.DebugString())
		return new(AbridgedCodec), nil
	}

	var (
		firstInt uint32
	)
	// not abridged version, we'll lookup codec!
	if size, bytes := conn.ReadN(4); size == 4 {
		firstInt = binary.LittleEndian.Uint32(bytes)
	} else {
		return nil, errUnexpectedEOF
	}

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
		logx.Infof("conn(%s) mtproto http.", conn.DebugString())
		return newMTProtoHttpProxyCodec(), nil
	}

	// check intermediate version
	if firstInt == INTERMEDIATE_FLAG {
		logx.Infof("conn(%s) intermediate version.", conn.DebugString())
		conn.ShiftN(4)
		return new(IntermediateCodec), nil
	}

	// check intermediate version
	if firstInt == PADDED_INTERMEDIATE_FLAG {
		logx.Infof("conn(%s) padded intermediate version.", conn.DebugString())
		conn.ShiftN(4)
		return new(PaddedIntermediateCodec), nil
	}

	// check PVrG
	if firstInt == PVRG_FLAG {
		logx.Infof("conn(%s) PVrG version.", conn.DebugString())
		return nil, errors.New("PVrG transport")
	}

	// check 0x02010316
	if firstInt == UNKNOWN_FLAG {
		logx.Infof("conn(%s) firstInt is 0x02010316.", conn.DebugString())
		return nil, errors.New("0x02010316 transport")
	}

	var (
		checkFullBuf []byte
	)

	if size, bytes := conn.ReadN(12); size == 12 {
		checkFullBuf = bytes
	} else {
		return nil, errUnexpectedEOF
	}

	secondInt := binary.BigEndian.Uint32(checkFullBuf[:8])
	if secondInt == FULL_FLAG {
		logx.Infof("conn(%s) mtproto full version.", conn.DebugString())
		conn.ShiftN(12)
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
	if size, bytes := conn.ReadN(64); size == 64 {
		obfuscatedBuf = bytes
	} else {
		return nil, errUnexpectedEOF
	}

	var tmp [64]byte
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
			conn.DebugString(),
			hex.EncodeToString(obfuscatedBuf[56:60]))
	}

	dcId := int16(binary.BigEndian.Uint16(obfuscatedBuf[60:]))
	// TODO: check dcId

	conn.ShiftN(64)

	logx.Infof("conn(%s) mtproto obfuscated version, {protocol_type: %d, dc_id: %d}", conn.DebugString(), protocolType, dcId)
	return newMTProtoObfuscatedCodec(d, e, protocolType, dcId), nil
}

/*
 * for client

func (c *MTProtoCodec) selectCodec() (net2.Codec, error) {
	var (
		temp  [64]byte
		bytes []byte

		dcId = int16(2)
	)

	for {
		bytes = crypto.RandomBytes(64)
		val := binary.BigEndian.Uint32(bytes)
		val2 := binary.BigEndian.Uint32(bytes[4:])

		if bytes[0] != 0xef &&
			val != 0x44414548 &&
			val != 0x54534f50 &&
			val != 0x20544547 &&
			val != 0x4954504f &&
			val != 0xeeeeeeee &&
			val != 0xdddddddd &&
			val != 0x02010316 &&
			val2 != 0x00000000 {

			bytes[56] = 0xef
			bytes[57] = 0xef
			bytes[58] = 0xef
			bytes[59] = 0xef
		}
		break
	}

	for a := 0; a < 48; a++ {
		temp[a] = bytes[a+8]
	}

	e, _ := crypto.NewAesCTR128Encrypt(temp[:32], temp[16:])

	for a := 0; a < 48; a++ {
		temp[a] = bytes[55-a]
	}
	d, _ := crypto.NewAesCTR128Encrypt(temp[:32], temp[16:])

	binary.BigEndian.PutUint16(bytes[60:], uint16(dcId))
	copy(temp[:], bytes)
	e.Encrypt(bytes)
	copy(bytes[56:], temp[56:])

	_, err := c.conn.Write(bytes)
	if err != nil {
		return nil, err
	}

	return NewMTProtoObfuscatedCodec(c.conn, d, e, ABRIDGED_INT32_FLAG, dcId), nil
}
*/

// Encode encodes frames upon server responses into TCP stream.
func (c *MTProtoCodec) Encode(conn gnet.Conn, msg interface{}) ([]byte, error) {
	if msg == nil {
		logx.Infof("conn(%s) msg is nil", conn.DebugString())
		return nil, nil
	}

	// log.Debugf("conn(%s), msg: %#v", conn.DebugString(), msg)
	if isClientType {
		//err := c.peekCodec()
		//if err != nil {
		//	return nil, err
		//}
		return nil, fmt.Errorf("conn(%s) clientType not impl", conn.DebugString())
	} else {
		if c.codec != nil {
			b, err := c.codec.Encode(conn, msg)
			if err != nil {
				logx.Errorf("conn(%s) encode msg error: %v", conn.DebugString(), err)
				return nil, err
			} else {
				return b, nil
			}
		} else {
			return nil, fmt.Errorf("conn(%s) codec is nil", conn.DebugString())
		}
	}
}

// Decode decodes frames from TCP stream via specific implementation.
func (c *MTProtoCodec) Decode(conn gnet.Conn) (interface{}, error) {
	if isClientType {
		return nil, fmt.Errorf("conn(%s) clientType not impl", conn.DebugString())
	} else {
		if c.firstPacket {
			if codec, err := c.peekCodec(conn); err != nil {
				logx.Errorf("connId(%s) peekCodec error: %v", conn.DebugString(), err)
				if err != errUnexpectedEOF {
					conn.Close()
				}
				return nil, err
			} else {
				c.codec = codec
				c.firstPacket = false
			}
		}
	}
	if msg, err := c.codec.Decode(conn); err != nil {
		logx.Errorf("conn(%s) decode error: %v", conn.DebugString(), err)
		if err != errUnexpectedEOF {
			conn.Close()
		}
		return nil, err
	} else {
		return msg, nil
	}
}

// Clone ...
func (c *MTProtoCodec) Clone() gnet.ICodec {
	return NewMTProtoCodec()
}

// Release ...
func (c *MTProtoCodec) Release() {
}
