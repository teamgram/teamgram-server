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

package net2

import (
	"errors"
	"fmt"
	"net"
	"sync"
	"sync/atomic"
)

var ConnectionClosedError = errors.New("Connection Closed")
var ConnectionBlockedError = errors.New("Connection Blocked")

var globalConnectionId uint64

type TcpConnection struct {
	name          string
	conn          net.Conn
	id            uint64
	codec         Codec
	sendChan      chan interface{}
	recvMutex     sync.Mutex
	sendMutex     sync.RWMutex
	closeFlag     int32
	closeChan     chan int
	closeMutex    sync.Mutex
	closeCallback closeCallback
	Context       interface{}
}

func NewTcpConnection(name string, conn net.Conn, sendChanSize int, codec Codec, cb closeCallback) *TcpConnection {
	// TODO(@benqi): globalConnectionId use

	if globalConnectionId >= (1 << 60) {
		atomic.StoreUint64(&globalConnectionId, 0)
	}

	conn2 := &TcpConnection{
		name:          name,
		conn:          conn,
		codec:         codec,
		closeChan:     make(chan int),
		id:            atomic.AddUint64(&globalConnectionId, 1),
		closeCallback: cb,
	}

	if sendChanSize > 0 {
		conn2.sendChan = make(chan interface{}, sendChanSize)
		go conn2.sendLoop()
	}
	return conn2
}

func (c *TcpConnection) String() string {
	return fmt.Sprintf("{connID: %d@%s-(%s->%s)}", c.id, c.name, c.conn.LocalAddr(), c.conn.RemoteAddr())
}

func (c *TcpConnection) LoadAddr() net.Addr {
	return c.conn.LocalAddr()
}

func (c *TcpConnection) RemoteAddr() net.Addr {
	return c.conn.RemoteAddr()
}

func (c *TcpConnection) Name() string {
	return c.name
}

func (c *TcpConnection) GetConnID() uint64 {
	return c.id
}

func (c *TcpConnection) GetNetConn() net.Conn {
	return c.conn
}

func (c *TcpConnection) IsClosed() bool {
	return atomic.LoadInt32(&c.closeFlag) == 1
}

func (c *TcpConnection) Close() error {
	if atomic.CompareAndSwapInt32(&c.closeFlag, 0, 1) {
		if c.closeCallback != nil {
			c.closeCallback.OnConnectionClosed(c)
		}

		close(c.closeChan)

		if c.sendChan != nil {
			c.sendMutex.Lock()
			close(c.sendChan)
			if clear, ok := c.codec.(ClearSendChan); ok {
				clear.ClearSendChan(c.sendChan)
			}
			c.sendMutex.Unlock()
		}

		err := c.codec.Close()
		return err
	}
	return ConnectionClosedError
}

func (c *TcpConnection) Codec() Codec {
	return c.codec
}

func (c *TcpConnection) Receive() (interface{}, error) {
	c.recvMutex.Lock()
	defer c.recvMutex.Unlock()

	msg, err := c.codec.Receive()
	if err != nil {
		c.Close()
	}
	return msg, err
}

func (c *TcpConnection) sendLoop() {
	defer c.Close()
	for {
		select {
		case msg, ok := <-c.sendChan:
			if !ok || c.codec.Send(msg) != nil {
				return
			}
		case <-c.closeChan:
			return
		}
	}
}

func (c *TcpConnection) Send(msg interface{}) error {
	if c.sendChan == nil {
		if c.IsClosed() {
			return ConnectionClosedError
		}

		c.sendMutex.Lock()
		defer c.sendMutex.Unlock()

		err := c.codec.Send(msg)
		if err != nil {
			c.Close()
		}
		return err
	}

	c.sendMutex.RLock()
	if c.IsClosed() {
		c.sendMutex.RUnlock()
		return ConnectionClosedError
	}

	select {
	case c.sendChan <- msg:
		c.sendMutex.RUnlock()
		return nil
	default:
		c.sendMutex.RUnlock()
		c.Close()
		return ConnectionBlockedError
	}
}
