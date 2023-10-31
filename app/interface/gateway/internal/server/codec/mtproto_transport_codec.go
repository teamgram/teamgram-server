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
	"errors"
	"flag"
	"fmt"
	"io"
	"net"

	"github.com/teamgram/marmota/pkg/net/ip"
	"github.com/teamgram/marmota/pkg/net2"
	"github.com/teamgram/proto/mtproto/crypto"

	log "github.com/zeromicro/go-zero/core/logx"
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
	// Tcp Transport

	// ABRIDGED_FLAG = 0xef
	ABRIDGED_FLAG = 0xef
	// ABRIDGED_INT32_FLAG = 0xfefefefe
	ABRIDGED_INT32_FLAG = 0xefefefef
	// INTERMEDIATE_FLAG = 0xcccccccc
	INTERMEDIATE_FLAG = 0xeeeeeeee
	// PADDED_INTERMEDIATE_FLAG = 0xbbbbbbbb
	PADDED_INTERMEDIATE_FLAG = 0xdddddddd
	UNKNOWN_FLAG             = 0x02010316
	PVRG_FLAG                = 0x47725650 // PVrG
	FULL_FLAG                = 0x00000000

	// PROXY_FLAG websocket flag
	PROXY_FLAG = 0xaaaaaaaa

	// HTTP_HEAD_FLAG Http Transport
	HTTP_HEAD_FLAG   = 0x44414548 // HEAD
	HTTP_POST_FLAG   = 0x54534f50 // POST
	HTTP_GET_FLAG    = 0x20544547 // GET
	HTTP_OPTION_FLAG = 0x4954504f // OPTION

	// 3d9ff4f1
)

var (
	isClientType bool
	isMTProto    bool // 是否使用MTProto - true为官方mtproto协议，false为定制协议（当前实现为ntproto）
	// transportType uint32
)

func init() {
	net2.RegisterProtocol("mtproto", NewMTProtoTransport())
	flag.BoolVar(&isClientType, "client", false, "client conn")
	flag.BoolVar(&isMTProto, "mtproto", true, "mtproto")
}

// MTProtoTransport
// //////////////////////////////////////////////////////////////////////////////////////////////////////////
type MTProtoTransport struct {
}

func NewMTProtoTransport() *MTProtoTransport {
	return &MTProtoTransport{}
}

func (m *MTProtoTransport) NewCodec(rw io.ReadWriter) (net2.Codec, error) {
	codec := &TransportCodec{
		codecType: TRANSPORT_TCP,
		conn:      rw.(net.Conn),
		proto:     m,
	}
	return codec, nil
}

type TransportCodec struct {
	codecType int // codec type
	conn      net.Conn
	codec     net2.Codec
	proto     *MTProtoTransport
	remoteIp  string
}

func (c *TransportCodec) peekCodec() error {
	if isMTProto {
		return c.peekMTProtoCodec()
	} else {
		return c.peekNTProtoCodec()
	}
}

/*
*

	  Android client code:

		RAND_bytes(bytes, 64);
		uint32_t val = (bytes[3] << 24) | (bytes[2] << 16) | (bytes[1] << 8) | (bytes[0]);
		uint32_t val2 = (bytes[7] << 24) | (bytes[6] << 16) | (bytes[5] << 8) | (bytes[4]);
		if (bytes[0] != 0xef &&
			val != 0x44414548 &&
			val != 0x54534f50 &&
			val != 0x20544547 &&
			val != 0x4954504f &&
			val != 0xeeeeeeee &&
			val2 != 0x00000000) {
			bytes[56] = bytes[57] = bytes[58] = bytes[59] = 0xef;
			break;
		}
*/
func (c *TransportCodec) peekMTProtoCodec() error {
	peek, _ := c.conn.(net2.PeekAble)

	// check abridged
	firstByte, err := peek.PeekByte()
	if err != nil {
		log.Errorf("transportCodec - read firstByte error: %v", err)
		return err
	}
	// log.Debugf("firstByte: %s", hex.EncodeToString([]byte{firstByte}))

	if firstByte == ABRIDGED_FLAG {
		log.Infof("mtproto abridged version.")
		c.codec = NewMTProtoAbridgedCodec(c.conn)
		peek.Discard(1)
		return nil
	}

	// not abridged version, we'll lookup codec!
	fB, err := peek.Peek(4)
	if err != nil {
		log.Errorf("read firstInt error: %v", err)
		return err
	}

	firstInt := binary.BigEndian.Uint32(fB)

	// check http
	if firstInt == HTTP_HEAD_FLAG ||
		firstInt == HTTP_POST_FLAG ||
		firstInt == HTTP_GET_FLAG ||
		firstInt == HTTP_OPTION_FLAG {
		// http 协议
		log.Infof("mtproto http.")

		// conn2 := NewMTProtoHttpProxyConn(conn)
		// c.conn = conn2
		c.codecType = TRANSPORT_HTTP
		c.codec = NewMTProtoHttpProxyCodec(c.conn)

		// c.proto.httpListener.acceptChan <- conn2
		return nil
	}

	// check intermediate version
	if firstInt == INTERMEDIATE_FLAG {
		//log.Warn("MTProtoProxyCodec - mtproto intermediate version, impl in the future!!")
		//return nil, errors.New("mtproto intermediate version not impl!!")
		log.Infof("mtproto intermediate version.")
		c.codec = NewMTProtoIntermediateCodec(c.conn)
		peek.Discard(4)
		return nil
	}

	// check intermediate version
	if firstInt == PADDED_INTERMEDIATE_FLAG {
		//log.Warn("MTProtoProxyCodec - mtproto intermediate version, impl in the future!!")
		//return nil, errors.New("mtproto intermediate version not impl!!")
		log.Infof("mtproto padded intermediate version.")
		c.codec = NewMTProtoPaddedIntermediateCodec(c.conn)
		peek.Discard(4)
		return nil
	}

	// check PVrG
	if firstInt == PVRG_FLAG {
		log.Infof("PVrG version")
		return errors.New("PVrG transport")
	}

	// check 0x02010316
	if firstInt == UNKNOWN_FLAG {
		log.Errorf("PVrG version")
		return errors.New("0x02010316 transport")
	}

	checkFullBuf, err := peek.Peek(12)
	if err != nil {
		log.Errorf("transportCodec - read b_4_60 error: %v", err)
		return err
	}
	secondInt := binary.BigEndian.Uint32(checkFullBuf[4:8])
	if secondInt == FULL_FLAG {
		log.Infof("mtproto full version.")
		c.codec = NewMTProtoFullCodec(c.conn)
		peek.Discard(12)
		return nil
	}

	// check obfuscated version
	obfuscatedBuf, err := peek.Peek(64)
	if err != nil {
		log.Errorf("peek error: %v", err)
		return err
	}

	log.Infof("obfuscatedBuf: %s", hex.EncodeToString(obfuscatedBuf))

	var tmp [64]byte
	// 生成decrypt_key
	for i := 0; i < 48; i++ {
		tmp[i] = obfuscatedBuf[55-i]
	}

	e, err := crypto.NewAesCTR128Encrypt(tmp[:32], tmp[32:48])
	if err != nil {
		// log.Errorf("NewAesCTR128Encrypt error: %s", err)
		return err
	}

	d, err := crypto.NewAesCTR128Encrypt(obfuscatedBuf[8:40], obfuscatedBuf[40:56])
	if err != nil {
		log.Errorf("NewAesCTR128Encrypt error: %s", err)
		return err
	}

	d.Encrypt(obfuscatedBuf)

	protocolType := binary.BigEndian.Uint32(obfuscatedBuf[56:])
	if protocolType != ABRIDGED_INT32_FLAG &&
		protocolType != INTERMEDIATE_FLAG &&
		protocolType != PADDED_INTERMEDIATE_FLAG {
		log.Errorf("transportCodec - invalid obfuscated protocol type - %s",
			hex.EncodeToString(obfuscatedBuf))
		return errors.New("mtproto buf[56:60]'s byte != 0xef")
	}

	dcId := int16(binary.BigEndian.Uint16(obfuscatedBuf[60:]))
	// TODO: check dcId

	if secondInt == PROXY_FLAG {
		c.remoteIp = ip.IntToIP(firstInt)
	}

	log.Infof("mtproto obfuscated version, protocol_type: %s", hex.EncodeToString(obfuscatedBuf[56:60]))
	c.codec = NewMTProtoObfuscatedCodec(c.conn, d, e, protocolType, dcId)

	peek.Discard(64)
	return nil
}

func (c *TransportCodec) peekNTProtoCodec() error {
	peek, _ := c.conn.(net2.PeekAble)

	// check obfuscated version
	obfuscatedBuf, err := peek.Peek(64)
	if err != nil {
		log.Errorf("peek error: %v", err)
		return err
	}

	log.Infof("obfuscatedBuf: %s", hex.EncodeToString(obfuscatedBuf))

	var tmp [64]byte
	// 生成decrypt_key
	for i := 0; i < 48; i++ {
		// tmp[i] = obfuscatedBuf[55-i]
		tmp[i] = obfuscatedBuf[63-i]
	}
	log.Infof("e: %s", hex.EncodeToString(tmp[:48]))
	log.Infof("d: %s", hex.EncodeToString(obfuscatedBuf[16:]))

	e, err := crypto.NewAesCTR128Encrypt(tmp[:32], tmp[32:48])
	if err != nil {
		// log.Errorf("NewAesCTR128Encrypt error: %s", err)
		return err
	}

	// d, err := crypto.NewAesCTR128Encrypt(obfuscatedBuf[16:48], obfuscatedBuf[40:56])
	d, err := crypto.NewAesCTR128Encrypt(obfuscatedBuf[16:48], obfuscatedBuf[48:64])
	if err != nil {
		log.Errorf("NewAesCTR128Encrypt error: %s", err)
		return err
	}

	d.Encrypt(obfuscatedBuf)

	// protocolType := binary.BigEndian.Uint32(obfuscatedBuf[56:])
	protocolType := binary.BigEndian.Uint32(obfuscatedBuf[8:])
	if protocolType != ABRIDGED_INT32_FLAG &&
		protocolType != INTERMEDIATE_FLAG &&
		protocolType != PADDED_INTERMEDIATE_FLAG {
		log.Errorf("transportCodec - invalid obfuscated protocol type - %s, buf[8:12] = %s",
			hex.EncodeToString(obfuscatedBuf), hex.EncodeToString(obfuscatedBuf[8:12]))

		return errors.New("mtproto buf[8:12]'s byte != 0xfefefefe")
	}

	// dcId := int16(binary.BigEndian.Uint16(obfuscatedBuf[60:]))
	dcId := int16(binary.BigEndian.Uint16(obfuscatedBuf[12:]))
	// TODO: check dcId

	// log.Debugf("mtproto obfuscated version, protocol_type: %s", hex.EncodeToString(obfuscatedBuf[56:60]))
	log.Infof("mtproto obfuscated version, protocol_type: %s", hex.EncodeToString(obfuscatedBuf[8:12]))
	c.codec = NewMTProtoObfuscatedCodec(c.conn, d, e, protocolType, dcId)

	peek.Discard(64)
	return nil
}

func (c *TransportCodec) selectCodec() (net2.Codec, error) {
	var (
		temp  [64]byte
		bytes []byte

		dcId = int16(1)
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

func (c *TransportCodec) Receive() (interface{}, error) {
	if isClientType {
		if c.codec == nil {
			return nil, fmt.Errorf("codec is nil")
		}
	} else {
		if c.codec == nil {
			err := c.peekCodec()
			if err != nil {
				return nil, err
			}
		}
	}
	return c.codec.Receive()
}

func (c *TransportCodec) Send(msg interface{}) error {
	if isClientType {
		//err := c.peekCodec()
		//if err != nil {
		//	return nil, err
		//}
	} else {
		if c.codec != nil {
			return c.codec.Send(msg)
		}
	}
	return fmt.Errorf("codec is nil")
}

func (c *TransportCodec) Close() error {
	if c.codec != nil {
		return c.codec.Close()
	} else {
		return nil
	}
}

func (c *TransportCodec) Context() interface{} {
	return c.remoteIp
}
