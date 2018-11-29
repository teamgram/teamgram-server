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
	"github.com/gogo/protobuf/proto"
	"github.com/golang/glog"
	"github.com/nebula-chat/chatengine/pkg/util"
	"github.com/nebula-chat/chatengine/pkg/grpc_util"
	"github.com/nebula-chat/chatengine/pkg/mysql_client"
	"github.com/nebula-chat/chatengine/pkg/net2"
	"github.com/nebula-chat/chatengine/pkg/redis_client"
	"github.com/nebula-chat/chatengine/messenger/sync/biz/dal/dao"
	"github.com/nebula-chat/chatengine/mtproto"
	"github.com/nebula-chat/chatengine/messenger/sync/biz/core/update"
	"github.com/nebula-chat/chatengine/service/idgen/client"
	"github.com/nebula-chat/chatengine/service/status/client"
	"google.golang.org/grpc"
	"sync"
	"github.com/nebula-chat/chatengine/messenger/sync/server/rpc"
	"github.com/nebula-chat/chatengine/mtproto/rpc"
)

func init() {
	proto.RegisterType((*mtproto.TLPushConnectToSessionServer)(nil), "mtproto.TLPushConnectToSessionServer")
	proto.RegisterType((*mtproto.ServerConnected)(nil), "mtproto.ServerConnected")
	proto.RegisterType((*mtproto.TLPushPushRpcResultData)(nil), "mtproto.TLPushPushRpcResultData")
	proto.RegisterType((*mtproto.TLPushPushUpdatesData)(nil), "mtproto.TLPushPushUpdatesData")
	proto.RegisterType((*mtproto.Bool)(nil), "mtproto.Bool")
}

type connContext struct {
	serverId  int32
	sessionId uint64
}

type syncServer struct {
	idgen idgen.UUIDGen
	// update     *update.UpdateModel
	status     status_client.StatusClient
	client     *zrpc.ZRpcClient
	server     *grpc_util.RPCServer
	impl       *rpc.SyncServiceImpl
	sessionMap sync.Map
}

func NewSyncServer() *syncServer {
	return &syncServer{}
}

////////////////////////////////////////////////////////////////////////////////////////////////////
// AppInstance interface
func (s *syncServer) Initialize() error {
	var err error

	err = InitializeConfig()
	if err != nil {
		glog.Fatal(err)
		return err
	}
	glog.Info("config loaded: ", Conf)

	// idgen
	s.idgen, _ = idgen.NewUUIDGen("snowflake", util.Int32ToString(Conf.ServerId))

	// 初始化mysql_client、redis_client
	mysql_client.InstallMysqlClientManager(Conf.Mysql)
	redis_client.InstallRedisClientManager(Conf.Redis)

	// 初始化redis_dao、mysql_dao
	dao.InstallMysqlDAOManager(mysql_client.GetMysqlClientManager())
	// dao.InstallRedisDAOManager(redis_client.GetRedisClientManager())

	s.status, _ = status_client.NewStatusClient("redis", "cache")

	s.server = grpc_util.NewRpcServer(Conf.Server.Addr, &Conf.Server.RpcDiscovery)
	s.client = zrpc.NewZRpcClient("brpc", Conf.SessionClient, s)

	return nil
}

func (s *syncServer) RunLoop() {
	go s.server.Serve(func(s2 *grpc.Server) {
		updateModel := update.NewUpdateModel(Conf.ServerId, "immaster", "cache")
		s.impl = rpc.NewSyncService(s, s.status, updateModel)
		mtproto.RegisterRPCSyncServer(s2, s.impl)
	})
	s.client.Serve()
}

func (s *syncServer) Destroy() {
	if s.impl != nil {
		s.impl.Destroy()
	}

	s.server.Stop()
	s.client.Stop()
}

///////////////////////////////////////////////////////////////////////////////////////
// Impl ZProtoClientCallBack
func (s *syncServer) OnNewClient(client *net2.TcpClient) {
	glog.Infof("onNewClient")
	connectToSession := &mtproto.TLPushConnectToSessionServer{
		SyncServerId: 1,
	}
	cntl := zrpc.NewController()
	cntl.SetServiceName("session")
	cntl.SetMethodName(proto.MessageName(connectToSession))
	// ...
	zrpc.SendMessageByClient(client, cntl, connectToSession)
}

func (s *syncServer) OnClientMessageArrived(client *net2.TcpClient, cntl *zrpc.ZRpcController, msg proto.Message) error {
	switch msg.(type) {
	case *mtproto.ServerConnected:
		glog.Infof("onSyncData - request(SessionServerConnectedRsp): {%v}", msg)
		// TODO(@benqi): bind server_id, server_name
		serverConnected, _ := msg.(*mtproto.ServerConnected)
		// res.GetServerId()
		ctx := &connContext{
			serverId:  serverConnected.To_PushSessionServerConnected().GetSessionServerId(),
			sessionId: client.GetConnection().GetConnID(),
		}
		client.GetConnection().Context = ctx
		glog.Info("store serverId: ", ctx)
		s.sessionMap.Store(ctx.serverId, client)
		// glog.Info("store serverId: ", ctx)
	case *mtproto.Bool:
		glog.Infof("onSyncData - request(PushUpdatesData): {%v}", msg)
	default:
		glog.Errorf("invalid register proto type: {%v}", msg)
	}
	return nil
}

func (s *syncServer) OnClientClosed(client *net2.TcpClient) {
	glog.Infof("OnConnectionClosed")

	ctx := client.GetConnection().Context
	if ctx != nil {
		if connCtx, ok := ctx.(*connContext); ok {
			s.sessionMap.Delete(connCtx.serverId)
		}
	}
}

func (s *syncServer) OnClientTimer(client *net2.TcpClient) {
	// glog.Infof("OnTimer")
}

func (s *syncServer) SendToSessionServer(serverId int, cntl *zrpc.ZRpcController, msg proto.Message) {
	if c, ok := s.sessionMap.Load(int32(serverId)); ok {
		client := c.(*net2.TcpClient)
		if client != nil {
			zrpc.SendMessageByClient(client, cntl, msg)
		} else {
			glog.Error("client type invalid")
		}
	} else {
		glog.Error("not found server id: ", serverId)
	}
}
