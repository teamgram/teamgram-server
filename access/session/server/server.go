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
	"github.com/gogo/protobuf/proto"
	"github.com/golang/glog"
	"github.com/nebula-chat/chatengine/mtproto"
	"github.com/nebula-chat/chatengine/mtproto/rpc"
	"github.com/nebula-chat/chatengine/pkg/grpc_util"
	"github.com/nebula-chat/chatengine/pkg/net2"
	"github.com/nebula-chat/chatengine/pkg/redis_client"
	"github.com/nebula-chat/chatengine/pkg/util"
	"github.com/nebula-chat/chatengine/service/idgen/client"
	"github.com/nebula-chat/chatengine/service/status/client"
	"sync"
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
	idgen      idgen.UUIDGen
	status     status_client.StatusClient
	server     *zrpc.ZRpcServer
	rpcClients map[string]*grpc_util.RPCClient

	bizRpcClient  *grpc_util.RPCClient
	nbfsRpcClient *grpc_util.RPCClient
	// syncRpcClient        mtproto.RPCSyncClient
	authSessionRpcClient mtproto.RPCSessionClient
	sessionManager       sync.Map // map[int64]*sessionClientList
	// sessionManager       *sessionManager
}

func NewSessionServer() *SessionServer {
	return &SessionServer{
		rpcClients: map[string]*grpc_util.RPCClient{},
	}
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

	// s.sessionManager = newSessionManager()
	s.server = zrpc.NewZRpcServer(Conf.Server, s)

	return nil
}

func (s *SessionServer) RunLoop() {
	// TODO(@benqi): check error
	// timingWheel.Start()

	s.bizRpcClient, _ = grpc_util.NewRPCClient(&Conf.BizRpcClient)
	s.nbfsRpcClient, _ = grpc_util.NewRPCClient(&Conf.NbfsRpcClient)

	for i := 0; i < len(Conf.RpcClients); i++ {
		s.rpcClients[Conf.RpcClients[i].ServiceName], _ = grpc_util.NewRPCClient(&Conf.RpcClients[i])
	}

	InstallRouter(s.rpcClients, Conf.RouterTables)
	// sync
	// c, _ := grpc_util.NewRPCClient(&Conf.SyncRpcClient)
	// s.syncRpcClient = mtproto.NewRPCSyncClient(c.GetClientConn())

	// auth_session
	c, _ := grpc_util.NewRPCClient(&Conf.AuthSessionRpcClient)
	s.authSessionRpcClient = mtproto.NewRPCSessionClient(c.GetClientConn())

	s.server.Serve()
}

func (s *SessionServer) Destroy() {
	glog.Infof("sessionServer - destroy...")
	s.server.Stop()
	// s.client.Stop()
}

////////////////////////////////////////////////////////////////////////////////////////////////////
// TcpConnectionCallback
func (s *SessionServer) OnServerNewConnection(conn *net2.TcpConnection) {
	glog.Infof("onNewConnection %v", conn.RemoteAddr())
}

func (s *SessionServer) OnServerMessageDataArrived(conn *net2.TcpConnection, cntl *zrpc.ZRpcController, msg proto.Message) error {
	// glog.Infof("onServerMessageDataArrived - receive data: {peer: %s, cntl: %s, msg: %s}", conn, cntl.RpcMeta, msg)
	switch msg.(type) {
	case *mtproto.TLSessionClientCreated:
		// glog.Info("onSessionClientNew - sessionClientNew: ", conn)
		return s.onSessionClientNew(conn.GetConnID(), cntl, msg.(*mtproto.TLSessionClientCreated))
	case *mtproto.TLSessionMessageData:
		return s.onSessionData(conn.GetConnID(), cntl, msg.(*mtproto.TLSessionMessageData))
	case *mtproto.TLSessionClientClosed:
		// glog.Info("onSessionClientClosed - sessionClientClosed: ", conn)
		return s.onSessionClientClosed(conn.GetConnID(), cntl, msg.(*mtproto.TLSessionClientClosed))
	case *mtproto.TLPushConnectToSessionServer:
		// glog.Infof("onPushConnectToSessionServer - request(ConnectToSessionServerReq): {%v}", msg)
		pushSessionServerConnected := &mtproto.TLPushSessionServerConnected{Data2: &mtproto.ServerConnected_Data{
			SessionServerId: getServerID(),
			ServerName:      "session",
		}}
		serverConnected := pushSessionServerConnected.To_ServerConnected()

		cntl2 := cntl.Clone()
		cntl2.MoveAttachment()
		cntl2.SetMethodName(proto.MessageName(serverConnected))
		zrpc.SendMessageByConn(conn, cntl2, serverConnected)
	case *mtproto.TLPushPushRpcResultData:
		pushData, _ := msg.(*mtproto.TLPushPushRpcResultData)

		err := s.onSyncRpcResultData(pushData.GetClientReqMsgId(), pushData.GetAuthKeyId(), cntl)
		var mBool *mtproto.Bool
		if err != nil {
			mBool = mtproto.ToBool(false)
		} else {
			mBool = mtproto.ToBool(true)
		}

		cntl2 := cntl.Clone()
		cntl2.MoveAttachment()
		cntl2.SetMethodName(proto.MessageName(mBool))
		zrpc.SendMessageByConn(conn, cntl2, mBool)
	case *mtproto.TLPushPushUpdatesData:
		pushData, _ := msg.(*mtproto.TLPushPushUpdatesData)
		// glog.Info("pushData - ", pushData)
		// isPush := pushData.GetIsPush() == 1
		err := s.onSyncData(pushData.GetAuthKeyId(), pushData.Pts, pushData.PtsCount, cntl)
		var mBool *mtproto.Bool
		if err != nil {
			mBool = mtproto.ToBool(false)
		} else {
			mBool = mtproto.ToBool(true)
		}

		cntl2 := cntl.Clone()
		cntl2.MoveAttachment()
		cntl2.SetMethodName(proto.MessageName(mBool))
		zrpc.SendMessageByConn(conn, cntl2, mBool)
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

////////////////////////////////////////////////////////////////////////////////////////////////////////
func (s *SessionServer) onSessionClientNew(connID uint64, cntl *zrpc.ZRpcController, sessData *mtproto.TLSessionClientCreated) error {
	// glog.Infof("onSessionClientNew - receive data: {client_conn_id: %s, md: %s, sess_data: %s}", connID, cntl.RpcMeta, sessData)

	//authKeyId := sessData.GetAuthKeyId()
	//var sessList *authSessions
	//
	//if vv, ok := s.sessionManager.Load(authKeyId); !ok {
	//	sessList = makeAuthSessions(authKeyId)
	//	s.sessionManager.Store(authKeyId, sessList)
	//	s.onNewSessionClientManager(sessList)
	//} else {
	//	sessList, _ = vv.(*authSessions)
	//}
	//
	clientConnID := makeClientConnID(int(sessData.GetConnType()), connID, uint64(sessData.GetClientConnId()))
	glog.Infof("onSessionClientNew - ID: {conn_id: %s, auth_key_id: %d}", clientConnID, sessData.GetAuthKeyId())
	//return sessList.onSessionClientNew(clientConnID)
	return nil
}

////////////////////////////////////////////////////////////////////////////////////////////////////////
func (s *SessionServer) onSessionData(connID uint64, cntl *zrpc.ZRpcController, sessData *mtproto.TLSessionMessageData) error {
	//glog.Infof("onSessionData - receive data: {conn_id: %d, md: %s, sess_data: %s}",
	//	connID,
	//	cntl.RpcMeta,
	//	sessData)
	clientConnID := makeClientConnID(int(sessData.GetConnType()), connID, uint64(sessData.GetClientConnId()))

	authKeyId := sessData.GetAuthKeyId()
	var sessList *authSessions
	if vv, ok := s.sessionManager.Load(authKeyId); !ok {
		glog.Infof("onSessionDataNew - ID: {conn_id: %s, auth_key_id: %d}", clientConnID, sessData.GetAuthKeyId())

		sessList = makeAuthSessions(authKeyId)
		s.sessionManager.Store(authKeyId, sessList)
		s.onNewSessionClientManager(sessList)
	} else {
		sessList, _ = vv.(*authSessions)
	}

	return sessList.onSessionDataArrived(clientConnID, cntl, cntl.MoveAttachment())
}

func (s *SessionServer) onSessionClientClosed(connID uint64, cntl *zrpc.ZRpcController, sessData *mtproto.TLSessionClientClosed) error {
	clientConnID := makeClientConnID(int(sessData.GetConnType()), connID, uint64(sessData.GetClientConnId()))
	glog.Infof("onSessionClientClosed - ID: {conn_id: %s, auth_key_id: %d}", clientConnID, sessData.GetAuthKeyId())

	var sessList *authSessions

	if vv, ok := s.sessionManager.Load(sessData.GetAuthKeyId()); !ok {
		err := fmt.Errorf("onSessionClientClosed - not find sessionList by ID: {conn_id: %s, auth_key_id: %d}", clientConnID, sessData.GetAuthKeyId())
		glog.Warning(err)
		return err
	} else {
		sessList, _ = vv.(*authSessions)
	}

	return sessList.onSessionClientClosed(clientConnID)
}

func (s *SessionServer) onSyncRpcResultData(authKeyId, clientMsgId int64, cntl *zrpc.ZRpcController) error {
	glog.Infof("onSyncRpcResultData - receive data: {auth_key_id: %d, client_msg_id: %d, md: %s}",
		authKeyId,
		clientMsgId,
		cntl)

	rawData := cntl.MoveAttachment()

	var sessList *authSessions
	if vv, ok := s.sessionManager.Load(authKeyId); !ok {
		err := fmt.Errorf("pushToSessionData - not find sessionList by authKeyId: {%d}", authKeyId)
		glog.Warning(err)
		return err
	} else {
		sessList, _ = vv.(*authSessions)
	}

	return sessList.onSyncRpcResultDataArrived(clientMsgId, cntl, rawData)
}

func (s *SessionServer) onSyncData(authKeyId int64, pts, ptsCount int32, cntl *zrpc.ZRpcController) error {
	glog.Infof("onSyncData - receive data: {auth_key_id: %d, md: %s}",
		authKeyId,
		cntl)

	dbuf := mtproto.NewDecodeBuf(cntl.MoveAttachment())
	obj := dbuf.Object()
	if obj == nil {
		return dbuf.GetError()
	}

	var sessList *authSessions
	if vv, ok := s.sessionManager.Load(authKeyId); !ok {
		err := fmt.Errorf("pushToSessionData - not find sessionList by authKeyId: {%d}", authKeyId)
		glog.Warning(err)
		return err
	} else {
		sessList, _ = vv.(*authSessions)
	}

	return sessList.onSyncDataArrived(cntl, pts, ptsCount, &messageData{obj: obj})
}

////////////////////////////////////////////////////////////////////////////////////////////////////////
// session event
func (s *SessionServer) onNewSessionClientManager(sess *authSessions) {
	sess.Start()
}

func (s *SessionServer) onCloseSessionClientManager(authKeyId int64) {
	if vv, ok := s.sessionManager.Load(authKeyId); ok {
		vv.(*authSessions).Stop()
		s.sessionManager.Delete(authKeyId)
	}
}
