// Copyright (c) 2018-present,  NebulaChat Studio (https://nebula.chat).
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

// Author: Benqi (wubenqi@gmail.com)

package mtproto

import (
	"encoding/hex"
	"errors"
	"github.com/golang/glog"
	"github.com/nebula-chat/chatengine/pkg/crypto"
	"github.com/nebula-chat/chatengine/pkg/net2"
	"io"
	// "net/http"
	"net"
	// "time"
	"fmt"
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

// 客户端or服务端
const (
	CODEC_TYPE_CLIENT = 1 // Client端
	CODEC_TYPE_SERVER = 2 // Server端
)

// 协议版本
const (
	MTPROTO_ABRIDGED_VERSION     = 1 // 删节版本
	MTPROTO_INTERMEDIATE_VERSION = 2 // 中间版本
	MTPROTO_FULL_VERSION         = 3 // 完整版本
	MTPROTO_APP_VERSION          = 4 // Androd等当前客户端使用版本
)

// Transport类型，不支持UDP
const (
	TRANSPORT_TCP  = 1 // TCP
	TRANSPORT_HTTP = 2 // HTTP
	TRANSPORT_UDP  = 3 // UDP, @benqi: 未发现有支持UDP的客户端
)

const (
	STATE_CONNECTED = iota
	// STATE_FIRST_BYTE		//
	// STATE_FIRST_INT32		//
	// STATE_FIRST_64BYTES		//
	STATE_DATA //
)

const (
	// Tcp Transport
	MTPROTO_ABRIDGED_FLAG     = 0xef
	MTPROTO_INTERMEDIATE_FLAG = 0xeeeeeeee

	// Http Transport
	HTTP_HEAD_FLAG   = 0x44414548
	HTTP_POST_FLAG   = 0x54534f50
	HTTP_GET_FLAG    = 0x20544547
	HTTP_OPTION_FLAG = 0x4954504f

	VAL2_FLAG = 0x00000000
)

func init() {
	net2.RegisterProtocol("mtproto", NewMTProtoProxy())
}

// https://core.telegram.org/mtproto#tcp-transport
// 服务端MTPProto代理
// 服务端需要兼容各种协议
type MTProtoProxy struct {
}

func NewMTProtoProxy() *MTProtoProxy {
	return &MTProtoProxy{}
}

func (m *MTProtoProxy) NewCodec(rw io.ReadWriter) (net2.Codec, error) {
	codec := &MTProtoProxyCodec{
		codecType: TRANSPORT_TCP,
		conn:      rw.(net.Conn),
		State:     STATE_CONNECTED,
		proto:     m,
	}
	return codec, nil
}

type MTProtoProxyCodec struct {
	codecType int // codec type
	conn      net.Conn
	State     int
	codec     net2.Codec
	proto     *MTProtoProxy
}

/////////////////////////////////////////////////////////////////////////////////////
/**
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
func (c *MTProtoProxyCodec) peekCodec() error {
	//if c.State == STATE_DATA {
	//	return nil
	//}
	conn, _ := c.conn.(*net2.BufferedConn)
	// var b_0_1 = make([]byte, 1)
	b_0_1, err := conn.Peek(1)

	// _, err := io.ReadFull(c.conn, b_0_1)
	if err != nil {
		glog.Errorf("MTProtoProxyCodec - read b_0_1 error: %v", err)
		return err
	}

	if b_0_1[0] == MTPROTO_ABRIDGED_FLAG {
		glog.Info("mtproto abridged version.")
		c.codec = NewMTProtoAbridgedCodec(conn)
		conn.Discard(1)
		return nil
	}

	// not abridged version, we'll lookup codec!
	// b_1_3 = make([]byte, 3)
	b_1_3, err := conn.Peek(4)
	if err != nil {
		glog.Errorf("MTProtoProxyCodec - read b_1_3 error: %v", err)
		return err
	}

	b_1_3 = b_1_3[1:4]
	// first uint32
	val := (uint32(b_1_3[2]) << 24) | (uint32(b_1_3[1]) << 16) | (uint32(b_1_3[0]) << 8) | (uint32(b_0_1[0]))
	if val == HTTP_HEAD_FLAG || val == HTTP_POST_FLAG || val == HTTP_GET_FLAG || val == HTTP_OPTION_FLAG {
		// http 协议
		glog.Info("mtproto http.")

		// conn2 := NewMTProtoHttpProxyConn(conn)
		// c.conn = conn2
		c.codecType = TRANSPORT_HTTP
		c.codec = NewMTProtoHttpProxyCodec(c.conn)

		// c.proto.httpListener.acceptChan <- conn2
		return nil
	}

	// an intermediate version
	if val == MTPROTO_INTERMEDIATE_FLAG {
		//glog.Warning("MTProtoProxyCodec - mtproto intermediate version, impl in the future!!")
		//return nil, errors.New("mtproto intermediate version not impl!!")
		glog.Info("mtproto intermediate version.")
		c.codec = NewMTProtoIntermediateCodec(conn)
		conn.Discard(4)
		return nil
	}

	// recv 4~64 bytes
	// var b_4_60 = make([]byte, 60)
	b_4_60, err := conn.Peek(64)
	// io.ReadFull(c.conn, b_4_60)
	if err != nil {
		glog.Errorf("MTProtoProxyCodec - read b_4_60 error: %v", err)
		return err
	}
	b_4_60 = b_4_60[4:64]
	val2 := (uint32(b_4_60[3]) << 24) | (uint32(b_4_60[2]) << 16) | (uint32(b_4_60[1]) << 8) | (uint32(b_4_60[0]))
	if val2 == VAL2_FLAG {
		glog.Info("mtproto full version.")
		c.codec = NewMTProtoFullCodec(conn)
		return nil
	}

	var tmp [64]byte
	// 生成decrypt_key
	for i := 0; i < 48; i++ {
		tmp[i] = b_4_60[51-i]
	}

	e, err := crypto.NewAesCTR128Encrypt(tmp[:32], tmp[32:48])
	if err != nil {
		// glog.Error("NewAesCTR128Encrypt error: %s", err)
		return err
	}

	d, err := crypto.NewAesCTR128Encrypt(b_4_60[4:36], b_4_60[36:52])
	if err != nil {
		glog.Errorf("NewAesCTR128Encrypt error: %s", err)
		return err
	}

	d.Encrypt(b_0_1)
	d.Encrypt(b_1_3)
	d.Encrypt(b_4_60)

	if b_4_60[52] != 0xef && b_4_60[53] != 0xef && b_4_60[54] != 0xef && b_4_60[55] != 0xef {
		glog.Errorf("MTProtoProxyCodec - first 56~59 byte != 0xef")
		return errors.New("mtproto buf[56:60]'s byte != 0xef!!")
	}

	glog.Info("first_bytes_64: ", hex.EncodeToString(b_0_1), hex.EncodeToString(b_1_3), hex.EncodeToString(b_4_60))
	c.codec = NewMTProtoAppCodec(conn, d, e)
	conn.Discard(64)

	return nil
}

func (c *MTProtoProxyCodec) Receive() (interface{}, error) {
	if c.codec == nil {
		err := c.peekCodec()
		if err != nil {
			return nil, err
		}
	}
	return c.codec.Receive()
}

func (c *MTProtoProxyCodec) Send(msg interface{}) error {
	if c.codec != nil {
		return c.codec.Send(msg)
	}
	return fmt.Errorf("codec is nil")
}

func (c *MTProtoProxyCodec) Close() error {
	if c.codec != nil {
		return c.codec.Close()
	} else {
		return nil
	}
}
