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
	"fmt"
	"github.com/golang/glog"
	"github.com/nebula-chat/chatengine/pkg/etcd_util"
	"github.com/nebula-chat/chatengine/pkg/grpc_util/service_discovery"
	"github.com/nebula-chat/chatengine/pkg/grpc_util/service_discovery/etcd3"
	"github.com/nebula-chat/chatengine/pkg/net2"
	"net"
)

type MTProtoServerCallback interface {
	OnServerNewConnection(conn *net2.TcpConnection)
	OnServerMessageDataArrived(c *net2.TcpConnection, msg *MTPRawMessage) error
	OnServerConnectionClosed(c *net2.TcpConnection)
}

type MTProtoServerConfig struct {
	Server    net2.ServerConfig
	Discovery service_discovery.ServiceDiscoveryServerConfig
}

type MTProtoServer struct {
	server   *net2.TcpServer
	registry *etcd3.EtcdReigistry
	callback MTProtoServerCallback
}

func NewMTProtoServer(conf *MTProtoServerConfig, cb MTProtoServerCallback) *MTProtoServer {
	lsn, err := net.Listen("tcp", conf.Server.Addr)
	if err != nil {
		glog.Fatal("listen error: %v", err)
	}

	server := &MTProtoServer{
		callback: cb,
	}
	server.server = net2.NewTcpServer(net2.TcpServerArgs{
		Listener:           lsn,
		ServerName:         conf.Server.Name,
		ProtoName:          conf.Server.ProtoName,
		SendChanSize:       1024,
		ConnectionCallback: server,
	}) // todo (yumcoder): set max connection

	server.registry, err = etcd_util.NewEtcdRegistry(conf.Discovery)
	if err != nil {
		glog.Fatal(err)
	}

	return server
}

///////////////////////////////////////////////////////////////////////////////////////////////
func (s *MTProtoServer) GetConnection(connID uint64) *net2.TcpConnection {
	return s.server.GetConnection(connID)
}

///////////////////////////////////////////////////////////////////////////////////////////////
func (s *MTProtoServer) Serve() {
	go s.server.Serve2()
	go s.registry.Register()
}

func (s *MTProtoServer) Stop() {
	s.registry.Deregister()
	s.server.Stop()
}

func (s *MTProtoServer) Pause() {
	s.server.Pause()
}

///////////////////////////////////////////////////////////////////////////////////////////////
func (s *MTProtoServer) OnNewConnection(conn *net2.TcpConnection) {
	glog.Infof("onNewConnection %v", conn.RemoteAddr())

	if s.callback != nil {
		s.callback.OnServerNewConnection(conn)
	}
}

func (s *MTProtoServer) OnConnectionDataArrived(conn *net2.TcpConnection, msg interface{}) error {
	glog.Infof("onConnectionDataArrived %v", conn.RemoteAddr())

	message, ok := msg.(*MTPRawMessage)
	if !ok {
		err := fmt.Errorf("recv invalid MTPRawMessage: {%v}", msg)
		glog.Error(err)
		return err
	}

	if s.callback != nil {
		s.callback.OnServerMessageDataArrived(conn, message)
	}

	return nil
}

func (s *MTProtoServer) OnConnectionClosed(conn *net2.TcpConnection) {
	glog.Infof("onConnectionClosed - %v", conn.RemoteAddr())

	if s.callback != nil {
		s.callback.OnServerConnectionClosed(conn)
	}
}
