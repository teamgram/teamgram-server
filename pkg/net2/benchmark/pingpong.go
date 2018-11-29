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

package main

import (
	"io"
	"net"
	"github.com/nebula-chat/chatengine/pkg/util"
	"github.com/golang/glog"
	"github.com/nebula-chat/chatengine/pkg/net2"
)

func init() {
	net2.RegisterProtocol("pingpong", NewPingpong(4096))
}

type Pingpong struct {
	readBuf int
}

func NewPingpong(readBuf int) net2.Protocol {
	if readBuf <= 0 {
		readBuf = 4096
	}

	return &Pingpong{
		readBuf: readBuf,
	}
}

func (b *Pingpong) NewCodec(rw io.ReadWriter) (cc net2.Codec, err error) {
	codec := new(PingpongCodec)

	codec.readBuf = b.readBuf
	codec.rw = rw
	codec.c = rw.(io.Closer)

	codec.buf = make([]byte, b.readBuf)
	return codec, nil
}

type PingpongCodec struct {
	readBuf int
	rw      io.ReadWriter
	c       io.Closer
	buf     []byte
}

func (c *PingpongCodec) Send(msg interface{}) error {
	if _, err := c.rw.Write(msg.([]byte)); err != nil {
		return err
	}
	return nil
}

func (c *PingpongCodec) Receive() (interface{}, error) {
	buf := make([]byte, c.readBuf)
	n, err := c.rw.Read(buf)
	if err == nil {
		return buf[:n], nil
	}
	return nil, err
}

func (c *PingpongCodec) Close() error {
	return c.c.Close()
}

/////////////////////////////////////////////////////////////////////////////////////////
type PingpongServer struct {
	server *net2.TcpServer
}

func NewPingpongServer(listener net.Listener, protoName string) *PingpongServer {
	s := &PingpongServer{}
	s.server = net2.NewTcpServer(net2.TcpServerArgs{Listener: listener, ServerName: "pingpong", ProtoName: protoName, SendChanSize: 0, ConnectionCallback: s, MaxConcurrentConnection: 2})
	return s
}

func (s *PingpongServer) Serve() {
	s.server.Serve()
}

func (s *PingpongServer) OnNewConnection(conn *net2.TcpConnection) {
	glog.Infof("OnNewConnection %v", conn.RemoteAddr())
}

func (s *PingpongServer) OnConnectionDataArrived(conn *net2.TcpConnection, msg interface{}) error {
	// glog.Infof("echo_server recv peer(%v) data: %v", conn.RemoteAddr(), msg)
	conn.Send(msg)
	return nil
}

func (s *PingpongServer) OnConnectionClosed(conn *net2.TcpConnection) {
	glog.Infof("OnConnectionClosed - %v", conn.RemoteAddr())
}

/////////////////////////////////////////////////////////////////////////////////////////
type PingpongInsance struct {
	server *PingpongServer
}

func (this *PingpongInsance) Initialize() error {
	listener, err := net.Listen("tcp", "0.0.0.0:33333")
	if err != nil {
		glog.Errorf("listen error: %v", err)
		return err
	}

	this.server = NewPingpongServer(listener, "pingpong")
	return nil
}

func (this *PingpongInsance) RunLoop() {
	this.server.Serve()
}

func (this *PingpongInsance) Destroy() {
	this.server.server.Stop()
}

func main() {
	instance := &PingpongInsance{}
	util.DoMainAppInstance(instance)
}
