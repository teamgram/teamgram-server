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
	"github.com/golang/glog"
	"github.com/nebula-chat/chatengine/pkg/util"
	"github.com/nebula-chat/chatengine/pkg/net2"
	"github.com/nebula-chat/chatengine/pkg/net2/codec"
	"net"
)

func init() {
	net2.RegisterProtocol("echo", codec.NewLengthBasedFrame(1024))
}

type EchoServer struct {
	server *net2.TcpServer
}

func NewEchoServer(listener net.Listener, protoName string) *EchoServer {
	//listener, err := net.Listen("tcp", "0.0.0.0:12345")
	//if err != nil {
	//	glog.Fatalf("listen error: %v", err)
	//	// return
	//}
	s := &EchoServer{}
	s.server = net2.NewTcpServer(net2.TcpServerArgs{Listener: listener, ServerName: "echo", ProtoName: protoName, SendChanSize: 1, ConnectionCallback: s, MaxConcurrentConnection: 2})
	return s
}

func (s *EchoServer) Serve() {
	s.server.Serve()
}

func (s *EchoServer) OnNewConnection(conn *net2.TcpConnection) {
	glog.Infof("OnNewConnection %v", conn.RemoteAddr())
}

func (s *EchoServer) OnConnectionDataArrived(conn *net2.TcpConnection, msg interface{}) error {
	glog.Infof("echo_server recv peer(%v) data: %v", conn.RemoteAddr(), msg)
	conn.Send(msg)
	return nil
}

func (s *EchoServer) OnConnectionClosed(conn *net2.TcpConnection) {
	glog.Infof("OnConnectionClosed - %v", conn.RemoteAddr())
}

type EchoClient struct {
	client *net2.TcpClientGroupManager
}

func NewEchoClient(protoName string, clients map[string][]string) *EchoClient {
	//listener, err := net.Listen("tcp", "0.0.0.0:12345")
	//if err != nil {
	//	glog.Fatalf("listen error: %v", err)
	//	// return
	//}
	c := &EchoClient{}
	c.client = net2.NewTcpClientGroupManager(protoName, clients, c)
	return c
}

func (c *EchoClient) Serve() {
	c.client.Serve()
}

func (c *EchoClient) OnNewClient(client *net2.TcpClient) {
	glog.Infof("OnNewConnection" + client.GetRemoteName())
	client.Send("ping\n")
}

func (c *EchoClient) OnClientDataArrived(client *net2.TcpClient, msg interface{}) error {
	glog.Infof("OnDataArrived - recv data: %v client: %s", msg, client.GetRemoteName())
	return client.Send("ping\n")
}

func (c *EchoClient) OnClientClosed(client *net2.TcpClient) {
	glog.Infof("OnConnectionClosed" + client.GetRemoteName())
	if client.AutoReconnect() {
		client.Reconnect()
	}
}

func (c *EchoClient) OnClientTimer(client *net2.TcpClient) {
	glog.Infof("OnTimer")
}

type EchoInsance struct {
	server *EchoServer
	client *EchoClient
}

func (this *EchoInsance) Initialize() error {
	//listener, err := net.Listen("tcp", "0.0.0.0:22345")
	//if err != nil {
	//	glog.Errorf("listen error: %v", err)
	//	return err
	//}
	//
	// this.server = NewEchoServer(listener, "echo")

	clients := map[string][]string{
		"echo1": []string{"127.0.0.1:22345", "192.168.1.101:22345"},
		"echo2": []string{"127.0.0.1:22345", "192.168.1.101:22345"},
		"echo3": []string{"127.0.0.1:22345", "192.168.1.101:22345"},
	}
	this.client = NewEchoClient("echo", clients)
	return nil
}

func (this *EchoInsance) RunLoop() {
	// go this.server.Serve()
	this.client.Serve()
}

func (this *EchoInsance) Destroy() {
	this.client.client.Stop()
	// this.server.server.Stop()
}

func main() {
	instance := &EchoInsance{}
	// app.AppInstance(instance)
	util.DoMainAppInstance(instance)
}
