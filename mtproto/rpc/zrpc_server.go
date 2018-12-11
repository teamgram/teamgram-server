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

package zrpc

import (
	"fmt"
	"github.com/golang/glog"
	"github.com/nebula-chat/chatengine/pkg/etcd_util"
	"github.com/nebula-chat/chatengine/pkg/grpc_util/service_discovery"
	"github.com/nebula-chat/chatengine/pkg/grpc_util/service_discovery/etcd3"
	"github.com/nebula-chat/chatengine/pkg/net2"
	"net"
	"github.com/gogo/protobuf/proto"
	"github.com/nebula-chat/chatengine/mtproto/rpc/brpc"
	"github.com/nebula-chat/chatengine/mtproto"
)

type ZRpcServerCallback interface {
	OnServerNewConnection(conn *net2.TcpConnection)
	OnServerMessageDataArrived(conn *net2.TcpConnection, cntl *ZRpcController, msg proto.Message) error
	OnServerConnectionClosed(conn *net2.TcpConnection)
}

type ZRpcServerConfig struct {
	Server    net2.ServerConfig
	Discovery service_discovery.ServiceDiscoveryServerConfig
}

type ZRpcServer struct {
	server   *net2.TcpServer
	registry *etcd3.EtcdReigistry
	callback ZRpcServerCallback
}

func NewZRpcServer(conf *ZRpcServerConfig, cb ZRpcServerCallback) *ZRpcServer {
	lsn, err := net.Listen("tcp", conf.Server.Addr)
	if err != nil {
		glog.Fatal("listen error: ", err)
	}

	server := &ZRpcServer{
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
func (s *ZRpcServer) Serve() {
	go s.server.Serve()
	go s.registry.Register()
}

func (s *ZRpcServer) Stop() {
	s.registry.Deregister()
	s.server.Stop()
}

func (s *ZRpcServer) Pause() {
	s.server.Pause()
}

///////////////////////////////////////////////////////////////////////////////////////////////
func (s *ZRpcServer) SendMessageByConnID(connID uint64, cntl *ZRpcController, msg proto.Message) error {
	conn := s.server.GetConnection(connID)
	if conn != nil {
		return SendMessageByConn(conn, cntl, msg)
	} else {
		err := fmt.Errorf("send data error, conn offline, connID: %d", connID)
		glog.Error(err)
		return err
	}
}

///////////////////////////////////////////////////////////////////////////////////////////////
func (s *ZRpcServer) OnNewConnection(conn *net2.TcpConnection) {
	// glog.Info("onNewConnection - ", conn)

	if s.callback != nil {
		s.callback.OnServerNewConnection(conn)
	}
}

func (s *ZRpcServer) OnConnectionDataArrived(conn *net2.TcpConnection, msg interface{}) error {
	// glog.Info("onConnectionDataArrived - ", conn)
	bmsg, ok := msg.(*brpc.BaiduRpcMessage)
	if !ok {
		return fmt.Errorf("recv invalid msg - {%v}", bmsg)
	}

	cntl, zmsg, err := SplitBaiduRpcMessage(bmsg)
	if err != nil {
		return err
	}

	switch zmsg.(type) {
	case *mtproto.TLPing:
		pong := &mtproto.TLPong{Data2: &mtproto.Pong_Data{
			MsgId:  cntl.GetCorrelationId(),
			PingId: zmsg.(*mtproto.TLPing).PingId,
		}}
		pong2 := pong.To_Pong()
		cntl.SetMethodName(proto.MessageName(pong2))
		SendMessageByConn(conn, cntl, pong2)
		return nil
	default:
		if s.callback != nil {
			return s.callback.OnServerMessageDataArrived(conn, cntl, zmsg)
		} else {
			err = fmt.Errorf("callback is nil")
			return err
		}
	}
}

func (s *ZRpcServer) OnConnectionClosed(conn *net2.TcpConnection) {
	// glog.Info("onConnectionClosed - ", conn)

	if s.callback != nil {
		s.callback.OnServerConnectionClosed(conn)
	}
}

///////////////////////////////////////////////////////////////////////////////////////////////
func SendMessageByConn(conn *net2.TcpConnection, cntl *ZRpcController, msg proto.Message) error {
	bmsg := MakeBaiduRpcMessage(cntl, msg)
	return conn.Send(bmsg)
}
