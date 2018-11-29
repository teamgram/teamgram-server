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
	"github.com/golang/glog"
	"net"
	"runtime/debug"
	"time"
)

type TcpClientCallBack interface {
	OnNewClient(c *TcpClient)
	OnClientDataArrived(c *TcpClient, msg interface{}) error
	OnClientClosed(c *TcpClient)
	OnClientTimer(c *TcpClient)
}

type TcpClient struct {
	remoteName    string
	conn          *TcpConnection
	autoReConnect bool
	chanSize      int
	callback      TcpClientCallBack
	remoteAddress string
	protoName     string
	timeInterval  time.Duration
}

func NewTcpClient(name string, chanSize int, protoName, address string, cb TcpClientCallBack) *TcpClient {
	client := &TcpClient{
		remoteName:    name,
		chanSize:      chanSize,
		autoReConnect: true,
		callback:      cb,
		remoteAddress: address,
		protoName:     protoName,
		timeInterval:  30 * time.Second,
	}

	// log.Info("NewTcpClient complete ", client)
	return client
}

func (c *TcpClient) GetRemoteName() string {
	return c.remoteName
}

func (c *TcpClient) GetRemoteAddress() string {
	return c.remoteAddress
}

func (c *TcpClient) Serve() bool {
	// log.Info("Start Connect to ", c.remoteName, " address ", c.remoteAddress)
	tcpConn, err := net.DialTimeout("tcp", c.remoteAddress, 5*time.Second)
	if err != nil {
		// log.Error(err.Error())
		c.Reconnect()
		return false
	}

	codec, err := NewCodecByName(c.protoName, tcpConn.(*net.TCPConn))
	c.conn = NewTcpConnection(c.remoteName, tcpConn.(*net.TCPConn), c.chanSize, codec, c)
	go c.establishTcpConnection(c.conn)
	return true
}

func (c *TcpClient) establishTcpConnection(conn *TcpConnection) {
	// glog.Info("establishTcpConnection...")
	defer func() {
		if err := recover(); err != nil {
			glog.Errorf("tcp_client handle panic: %v\n%s", err, debug.Stack())
		}
	}()

	// c.conn = conn
	c.onNewConnection(conn)

	for {
		msg, err := conn.Receive()
		if err != nil {
			glog.Errorf("recv error: %v", err)
			return
		}

		if msg == nil {
			// glog.Errorf("recv a nil msg: %v", conn)
			// 是否需要关闭？
			continue
		}

		if c.callback != nil {
			if err := c.callback.OnClientDataArrived(c, msg); err != nil {
				// TODO: 是否需要关闭?
			}
		}
	}
}

func (c *TcpClient) Stop() {
	c.autoReConnect = false
	if c.conn != nil {
		c.conn.Close()
	}
}

func (c *TcpClient) GetConnection() *TcpConnection {
	return c.conn
}

func (c *TcpClient) Send(msg interface{}) error {
	if c.conn != nil && !c.conn.IsClosed() {
		return c.conn.Send(msg)
	}
	return errors.New("tcpClient is not running")
}

func (c *TcpClient) AutoReconnect() bool {
	return c.autoReConnect
}

func (c *TcpClient) Reconnect() {
	time.AfterFunc(5*time.Second, func() {
		glog.Error("auto Reconnect server ", c.remoteName, " address ", c.remoteAddress)
		c.Serve()
	})
}

func (c *TcpClient) OnConnectionClosed(conn Connection) {
	c.onConnectionClosed(conn.(*TcpConnection))
}

func (c *TcpClient) onNewConnection(conn *TcpConnection) {
	if c.callback != nil {
		c.callback.OnNewClient(c)
	}
}

func (c *TcpClient) onConnectionClosed(conn *TcpConnection) {
	if c.callback != nil {
		c.callback.OnClientClosed(c)
	}
}

func (c *TcpClient) GetTimer() time.Duration {
	return c.timeInterval
}

func (c *TcpClient) SetTimer(d time.Duration) {
	c.timeInterval = d
}

func (c *TcpClient) StartTimer() {
	if c.conn != nil && !c.conn.IsClosed() {
		time.AfterFunc(c.timeInterval, func() {
			if c.callback != nil {
				c.callback.OnClientTimer(c)
			}

			c.StartTimer()
		})
	}
}
