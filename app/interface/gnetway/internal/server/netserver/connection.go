// Copyright (c) 2024 The Teamgooo Authors. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package netserver

import (
	"bufio"
	"math/rand"
	"net"
	"strings"
	"sync"
	"time"

	"github.com/lxzan/gws"
	"github.com/teamgooo/teamgooo-server/app/interface/gnetway/internal/server/gnet/codec"
	"github.com/teamgooo/teamgooo-server/pkg/proto/bin"
	"github.com/zeromicro/go-zero/core/jsonx"
	"github.com/zeromicro/go-zero/core/logx"
)

type HandshakeStateCtx struct {
	State         int32      `json:"state,omitempty"`
	ResState      int32      `json:"res_state,omitempty"`
	Nonce         bin.Int128 `json:"nonce,omitempty"`
	ServerNonce   bin.Int128 `json:"server_nonce,omitempty"`
	NewNonce      bin.Int256 `json:"new_nonce,omitempty"`
	A             []byte     `json:"a,omitempty"`
	P             []byte     `json:"p,omitempty"`
	HandshakeType int        `json:"handshake_type"`
	ExpiresIn     int32      `json:"expires_in,omitempty"`
}

func (m *HandshakeStateCtx) DebugString() string {
	s, _ := jsonx.MarshalToString(m)
	return s
}

type connection struct {
	id         int64
	conn       net.Conn
	reader     *bufio.Reader
	writer     *bufio.Writer
	writeMu    sync.Mutex
	codec      codec.Codec
	authKey    *authKeyUtil
	sessionId  int64
	handshakes []*HandshakeStateCtx
	clientIp   string
	tcp        bool
	websocket  bool
	gwsConn    *gws.Conn
	newSession bool
	nextSeqNo  int32
	closeDate  int64
	closed     bool
	closeMu    sync.Mutex
	logx.Logger
}

func newConnection(id int64, conn net.Conn, isTcp, isWebsocket bool) *connection {
	clientIp := strings.Split(conn.RemoteAddr().String(), ":")[0]

	c := &connection{
		id:         id,
		conn:       conn,
		reader:     bufio.NewReaderSize(conn, 65536),
		writer:     bufio.NewWriterSize(conn, 65536),
		tcp:        isTcp,
		websocket:  isWebsocket,
		clientIp:   clientIp,
		closeDate:  time.Now().Unix() + 30, // 30 seconds initial timeout
		codec:      nil,
	}

	return c
}

func (c *connection) ID() int64 {
	return c.id
}

func (c *connection) RemoteAddr() string {
	return c.conn.RemoteAddr().String()
}

func (c *connection) LocalAddr() string {
	return c.conn.LocalAddr().String()
}

func (c *connection) Write(data []byte) (int, error) {
	c.writeMu.Lock()
	defer c.writeMu.Unlock()

	if c.closed {
		return 0, net.ErrClosed
	}

	n, err := c.writer.Write(data)
	if err != nil {
		return n, err
	}

	return n, c.writer.Flush()
}

func (c *connection) Close() error {
	c.closeMu.Lock()
	defer c.closeMu.Unlock()

	if c.closed {
		return nil
	}

	c.closed = true

	if c.gwsConn != nil {
		_ = c.gwsConn.WriteClose(1000, []byte("connection closed"))
	}

	return c.conn.Close()
}

func (c *connection) IsClosed() bool {
	c.closeMu.Lock()
	defer c.closeMu.Unlock()
	return c.closed
}

func (c *connection) ResetTimeout() {
	c.closeDate = time.Now().Unix() + 300 + rand.Int63()%10
}

func (c *connection) generateMessageSeqNo(increment bool) int32 {
	value := c.nextSeqNo
	if increment {
		c.nextSeqNo++
		return value*2 + 1
	} else {
		return value * 2
	}
}

func (c *connection) getAuthKey() *authKeyUtil {
	return c.authKey
}

func (c *connection) putAuthKey(k *authKeyUtil) {
	c.authKey = k
}

func (c *connection) getHandshakeStateCtx(nonce bin.Int128) *HandshakeStateCtx {
	for _, state := c.handshakes {
		if nonce == state.Nonce {
			return state
		}
	}
	return nil
}

func (c *connection) putHandshakeStateCtx(state *HandshakeStateCtx) {
	c.handshakes = append(c.handshakes, state)
}

type connectionManager struct {
	connections sync.Map // map[int64]*connection
}

func newConnectionManager() *connectionManager {
	return &connectionManager{}
}

func (cm *connectionManager) newConnection(id int64, conn net.Conn, isTcp, isWebsocket bool) *connection {
	c := newConnection(id, conn, isTcp, isWebsocket)
	cm.connections.Store(id, c)
	return c
}

func (cm *connectionManager) get(id int64) (*connection, bool) {
	v, ok := cm.connections.Load(id)
	if !ok {
		return nil, false
	}
	return v.(*connection), true
}

func (cm *connectionManager) remove(id int64) {
	cm.connections.Delete(id)
}

func (cm *connectionManager) count() int {
	count := 0
	cm.connections.Range(func(key, value interface{}) bool {
		count++
		return true
	})
	return count
}

func (cm *connectionManager) iterate(fn func(c *connection)) {
	cm.connections.Range(func(key, value interface{}) bool {
		fn(value.(*connection))
		return true
	})
}

func (cm *connectionManager) withConnection(id int64, fn func(c *connection)) {
	if c, ok := cm.get(id); ok {
		fn(c)
	}
}

func (cm *connectionManager) closeAll() {
	cm.connections.Range(func(key, value interface{}) bool {
		c := value.(*connection)
		_ = c.Close()
		return true
	})
}