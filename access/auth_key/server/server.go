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
	"fmt"
	"github.com/nebula-chat/chatengine/pkg/grpc_util"
	"github.com/nebula-chat/chatengine/pkg/net2"
	"github.com/nebula-chat/chatengine/mtproto"
	"github.com/nebula-chat/chatengine/mtproto/rpc"
	"time"
	"github.com/gogo/protobuf/proto"
	"github.com/golang/glog"
	"strconv"
)

func init() {
	proto.RegisterType((*mtproto.TLSessionMessageData)(nil), "mtproto.TLSessionMessageData")
	proto.RegisterType((*mtproto.TLHandshakeData)(nil), "mtproto.TLHandshakeData")
	proto.RegisterType((*mtproto.RawMessageData)(nil), "mtproto.RawMessageData")
}

type AuthKeyServer struct {
	handshake            *handshake
	server               *zrpc.ZRpcServer
	authSessionRpcClient mtproto.RPCSessionClient
}

func NewAuthKeyServer() *AuthKeyServer {
	return &AuthKeyServer{}
}

////////////////////////////////////////////////////////////////////////////////////////////////////
// AppInstance interface
func (s *AuthKeyServer) Initialize() error {
	err := InitializeConfig()
	if err != nil {
		glog.Fatal(err)
		return err
	}
	glog.Info("load conf: ", Conf)

	// 初始化mysql_client、redis_client
	// mysql_client.InstallMysqlClientManager(Conf.Mysql)
	// 初始化redis_dao、mysql_dao
	// dao.InstallMysqlDAOManager(mysql_client.GetMysqlClientManager())

	s.server = zrpc.NewZRpcServer(Conf.Server, s)
	// s.rpcServer = grpc_util.NewRpcServer(Conf.RpcServer.Addr, &Conf.RpcServer.RpcDiscovery)

	return nil
}

func (s *AuthKeyServer) RunLoop() {
	keyFingerprint, err := strconv.ParseUint(Conf.KeyFingerprint, 10, 64)
	if err != nil {
		glog.Fatal(err)
		return
	}
	c, _ := grpc_util.NewRPCClient(&Conf.AuthSessionRpcClient)
	s.authSessionRpcClient = mtproto.NewRPCSessionClient(c.GetClientConn())
	s.handshake = newHandshake(s.authSessionRpcClient, Conf.KeyFile, keyFingerprint)

	go s.server.Serve()
}

func (s *AuthKeyServer) Destroy() {
	glog.Infof("sessionServer - destroy...")
	s.server.Stop()
	// s.rpcServer.Stop()
	time.Sleep(1 * time.Second)
}

//////////////////////////////////////////////////////////////////////////////////////////////////////
func (s *AuthKeyServer) OnServerNewConnection(conn *net2.TcpConnection) {
	glog.Infof("onNewConnection %v", conn.RemoteAddr())
}

func (s *AuthKeyServer) OnServerMessageDataArrived(conn *net2.TcpConnection, cntl *zrpc.ZRpcController, msg proto.Message) error {
	glog.Infof("onServerMessageDataArrived - msg: %v", msg)

	hmsg, ok := msg.(*mtproto.TLHandshakeData)
	if !ok {
		err := fmt.Errorf("invalid handshakeMessage: {%v}", msg)
		glog.Error(err)
		return err
	}

	hrsp, err := s.handshake.onHandshake(conn, cntl, hmsg)
	if err != nil {
		glog.Error(err)
		return nil
	}

	// Fix onMsgAck return nil bug.
	if hrsp == nil {
		return nil
	}

	return zrpc.SendMessageByConn(conn, cntl, hrsp)
}

func (s *AuthKeyServer) OnServerConnectionClosed(conn *net2.TcpConnection) {
	glog.Infof("onConnectionClosed - %v", conn.RemoteAddr())
}
