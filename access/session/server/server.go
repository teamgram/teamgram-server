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
	"time"
	"github.com/golang/glog"
	"github.com/nebula-chat/chatengine/pkg/util"
	"github.com/nebula-chat/chatengine/pkg/grpc_util"
	"github.com/nebula-chat/chatengine/pkg/net2"
	"github.com/nebula-chat/chatengine/pkg/redis_client"
	"github.com/nebula-chat/chatengine/mtproto"
	"github.com/nebula-chat/chatengine/service/idgen/client"
	"github.com/nebula-chat/chatengine/service/status/client"
	"github.com/nebula-chat/chatengine/mtproto/rpc"
	"github.com/gogo/protobuf/proto"
)

func init() {
	//
	proto.RegisterType((*mtproto.TLSessionClientCreated)(nil), "mtproto.TLSessionClientCreated")
	proto.RegisterType((*mtproto.TLSessionClientClosed)(nil), "mtproto.TLSessionClientClosed")
	proto.RegisterType((*mtproto.SessionClientEvent)(nil), "mtproto.SessionClientEvent")

	proto.RegisterType((*mtproto.TLSessionMessageData)(nil), "mtproto.TLSessionMessageData")
	proto.RegisterType((*mtproto.RawMessageData)(nil), "mtproto.RawMessageData")

	// sync
	proto.RegisterType((*mtproto.TLPushConnectToSessionServer)(nil), "mtproto.TLPushConnectToSessionServer")
	proto.RegisterType((*mtproto.ServerConnected)(nil), "mtproto.ServerConnected")

	proto.RegisterType((*mtproto.TLPushPushRpcResultData)(nil), "mtproto.TLPushPushRpcResultData")
	proto.RegisterType((*mtproto.TLPushPushUpdatesData)(nil), "mtproto.TLPushPushUpdatesData")
	proto.RegisterType((*mtproto.Bool)(nil), "mtproto.Bool")
}

type SessionServer struct {
	idgen                idgen.UUIDGen
	status               status_client.StatusClient
	server               *zrpc.ZRpcServer
	bizRpcClient         *grpc_util.RPCClient
	nbfsRpcClient        *grpc_util.RPCClient
	syncRpcClient        mtproto.RPCSyncClient
	authSessionRpcClient mtproto.RPCSessionClient
	sessionManager       *sessionManager
}

func NewSessionServer() *SessionServer {
	return &SessionServer{}
}

////////////////////////////////////////////////////////////////////////////////////////////////////
// AppInstance interface
func (s *SessionServer) Initialize() error {
	err := InitializeConfig()
	if err != nil {
		glog.Fatal(err)
		return err
	}
	glog.Info("load conf: ", Conf)

	// idgen
	s.idgen, _ = idgen.NewUUIDGen("snowflake", util.Int32ToString(Conf.ServerId))
	// 初始化mysql_client、redis_client
	redis_client.InstallRedisClientManager(Conf.Redis)

	s.status, _ = status_client.NewStatusClient("redis", "cache")

	// 初始化redis_dao、mysql_dao
	//dao.InstallRedisDAOManager(redis_client.GetRedisClientManager())
	// TODO(@benqi): config cap
	InitCacheAuthManager(1024*1024, &Conf.AuthSessionRpcClient)

	s.sessionManager = newSessionManager()
	s.server = zrpc.NewZRpcServer(Conf.Server, s)

	return nil
}

func (s *SessionServer) RunLoop() {
	// TODO(@benqi): check error
	// timingWheel.Start()

	s.bizRpcClient, _ = grpc_util.NewRPCClient(&Conf.BizRpcClient)
	s.nbfsRpcClient, _ = grpc_util.NewRPCClient(&Conf.NbfsRpcClient)

	// sync
	c, _ := grpc_util.NewRPCClient(&Conf.SyncRpcClient)
	s.syncRpcClient = mtproto.NewRPCSyncClient(c.GetClientConn())

	// auth_session
	c, _ = grpc_util.NewRPCClient(&Conf.AuthSessionRpcClient)
	s.authSessionRpcClient = mtproto.NewRPCSessionClient(c.GetClientConn())

	s.server.Serve()
}

func (s *SessionServer) Destroy() {
	glog.Infof("sessionServer - destroy...")
	s.server.Stop()
	time.Sleep(1 * time.Second)
	// s.client.Stop()
}

////////////////////////////////////////////////////////////////////////////////////////////////////
// TcpConnectionCallback
func (s *SessionServer) OnServerNewConnection(conn *net2.TcpConnection) {
	glog.Infof("OnNewConnection %v", conn.RemoteAddr())
}

func (s *SessionServer) OnServerMessageDataArrived(conn *net2.TcpConnection, cntl *zrpc.ZRpcController, msg proto.Message) error {
	glog.Infof("OnServerMessageDataArrived - receive data: {peer: %s, cntl: %s, msg: %s}", conn, cntl.RpcMeta, msg)
	switch msg.(type) {
	case *mtproto.TLSessionClientCreated:
		// glog.Info("onSessionClientNew - sessionClientNew: ", conn)
		// return s.sessionManager.onSessionClientNew(conn.GetConnID(), md, msg.(*zproto.ZProtoSessionClientNew))
	case *mtproto.TLSessionMessageData:
		return s.sessionManager.onSessionData(conn.GetConnID(), cntl, msg.(*mtproto.TLSessionMessageData))
	case *mtproto.TLSessionClientClosed:
		// glog.Info("onSessionClientClosed - sessionClientClosed: ", conn)
		// return s.sessionManager.onSessionClientClosed(conn.GetConnID(), md, msg.(*zproto.ZProtoSessionClientClosed))
	case *mtproto.TLPushConnectToSessionServer:
		glog.Infof("onSyncData - request(ConnectToSessionServerReq): {%v}", msg)
		pushSessionServerConnected := &mtproto.TLPushSessionServerConnected{Data2: &mtproto.ServerConnected_Data{
			SessionServerId: getServerID(),
			ServerName:      "session",
		}}
		serverConnected := pushSessionServerConnected.To_ServerConnected()
		cntl.SetMethodName(proto.MessageName(serverConnected))
		zrpc.SendMessageByConn(conn, cntl, serverConnected)
	case *mtproto.TLPushPushRpcResultData:
		pushData, _ := msg.(*mtproto.TLPushPushRpcResultData)

		err := s.sessionManager.onSyncRpcResultData(pushData.GetClientReqMsgId(), pushData.GetAuthKeyId(), cntl)
		var mBool *mtproto.Bool
		if err != nil {
			mBool = mtproto.ToBool(false)
		} else {
			mBool = mtproto.ToBool(true)
		}
		cntl.SetMethodName(proto.MessageName(mBool))
		zrpc.SendMessageByConn(conn, cntl, mBool)
	case *mtproto.TLPushPushUpdatesData:
		pushData, _ := msg.(*mtproto.TLPushPushUpdatesData)

		err := s.sessionManager.onSyncData(pushData.GetAuthKeyId(), cntl)
		var mBool *mtproto.Bool
		if err != nil {
			mBool = mtproto.ToBool(false)
		} else {
			mBool = mtproto.ToBool(true)
		}
		cntl.SetMethodName(proto.MessageName(mBool))
		zrpc.SendMessageByConn(conn, cntl, mBool)
	default:
		err := fmt.Errorf("invalid payload type: %v", msg)
		glog.Error(err)
		return err
	}

	return nil
}

func (s *SessionServer) OnServerConnectionClosed(conn *net2.TcpConnection) {
	glog.Infof("onConnectionClosed - %v", conn.RemoteAddr())
}

//func (s *SessionServer) SendToClientData(connID, sessionID uint64, md *zproto.ZProtoMetadata, buf []byte) error {
//	glog.Infof("sendToClientData - {%d, %d}", connID, sessionID)
//	//conn := s.server.GetConnection(connID)
//	//if conn != nil {
//	//	return sendDataByConnection(conn, sessionID, md, buf)
//	//} else {
//	//	err := fmt.Errorf("send data error, conn offline, connID: %d", connID)
//	//	glog.Error(err)
//	//	return err
//	//}
//	return nil
//}
