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

package server

import (
	"github.com/golang/glog"
	"github.com/nebula-chat/chatengine/pkg/net2"
	"net"
)

type MultiProtoServer struct {
	server *net2.TcpServer
}

func NewMultiProtoServer(listener net.Listener, protoName string) *MultiProtoServer {
	s := &MultiProtoServer{}
	s.server = net2.NewTcpServer(net2.TcpServerArgs{Listener: listener, ServerName: "multi_proto", ProtoName: protoName, SendChanSize: 1, ConnectionCallback: s, MaxConcurrentConnection: 2})
	return s
}

func (s *MultiProtoServer) Serve() {
	s.server.Serve()
}

func (s *MultiProtoServer) OnNewConnection(conn *net2.TcpConnection) {
	glog.Infof("OnNewConnection %v", conn.RemoteAddr())

}

func (s *MultiProtoServer) OnConnectionDataArrived(conn *net2.TcpConnection, msg interface{}) error {
	glog.Infof("echo_server recv peer(%v) data: %v", conn.RemoteAddr(), msg)
	conn.Send(msg)
	return nil
}

func (s *MultiProtoServer) OnConnectionClosed(conn *net2.TcpConnection) {
	glog.Infof("OnConnectionClosed - %v", conn.RemoteAddr())
}

type MultiProtoInsance struct {
	server *MultiProtoServer
	// client       *EchoClient
}

func (this *MultiProtoInsance) Initialize() error {
	net2.RegisterProtocol("multi_proto", NewMultiProtoProxy())

	listener, err := net.Listen("tcp", "0.0.0.0:22345")
	if err != nil {
		glog.Errorf("listen error: %v", err)
		return err
	}

	this.server = NewMultiProtoServer(listener, "multi_proto")
	return nil
}

func (this *MultiProtoInsance) RunLoop() {
	// go this.server.httpServer.Serve(this.server.httpListener)
	this.server.Serve()
	// this.client.Serve()
}

func (this *MultiProtoInsance) Destroy() {
	// this.client.client.Stop()
	this.server.server.Stop()
}
