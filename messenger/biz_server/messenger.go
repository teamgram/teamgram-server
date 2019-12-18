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
	"flag"
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/golang/glog"
	"github.com/nebula-chat/chatengine/messenger/biz_server/biz/core"
	"github.com/nebula-chat/chatengine/messenger/biz_server/biz/dal/dao"
	"github.com/nebula-chat/chatengine/messenger/biz_server/server/account"
	"github.com/nebula-chat/chatengine/messenger/biz_server/server/auth"
	"github.com/nebula-chat/chatengine/messenger/biz_server/server/bots"
	"github.com/nebula-chat/chatengine/messenger/biz_server/server/channels"
	"github.com/nebula-chat/chatengine/messenger/biz_server/server/contacts"
	"github.com/nebula-chat/chatengine/messenger/biz_server/server/help"
	"github.com/nebula-chat/chatengine/messenger/biz_server/server/langpack"
	"github.com/nebula-chat/chatengine/messenger/biz_server/server/messages"
	"github.com/nebula-chat/chatengine/messenger/biz_server/server/payments"
	"github.com/nebula-chat/chatengine/messenger/biz_server/server/phone"
	"github.com/nebula-chat/chatengine/messenger/biz_server/server/photos"
	"github.com/nebula-chat/chatengine/messenger/biz_server/server/stickers"
	"github.com/nebula-chat/chatengine/messenger/biz_server/server/updates"
	"github.com/nebula-chat/chatengine/messenger/biz_server/server/users"
	"github.com/nebula-chat/chatengine/messenger/sync/sync_client"
	"github.com/nebula-chat/chatengine/mtproto"
	"github.com/nebula-chat/chatengine/pkg/grpc_util"
	"github.com/nebula-chat/chatengine/pkg/grpc_util/service_discovery"
	"github.com/nebula-chat/chatengine/pkg/mysql_client"
	"github.com/nebula-chat/chatengine/pkg/redis_client"
	"github.com/nebula-chat/chatengine/pkg/util"
	"github.com/nebula-chat/chatengine/service/auth_session/client"
	"github.com/nebula-chat/chatengine/service/document/client"
	"google.golang.org/grpc"
	// "github.com/nebula-chat/chatengine/service/bots/client"
)

////////////////////////////////////////////////////////////////////////////////////////////////////
// Conf.go
var (
	confPath string
	Conf     *messengerConfig
)

type messengerConfig struct {
	ServerId             int32 // 服务器ID
	RelayIp              string
	RpcServer            *grpc_util.RPCServerConfig
	Mysql                []mysql_client.MySQLConfig
	Redis                []redis_client.RedisConfig
	NbfsRpcClient        *service_discovery.ServiceDiscoveryClientConfig
	SyncRpcClient1       *service_discovery.ServiceDiscoveryClientConfig
	SyncRpcClient2       *service_discovery.ServiceDiscoveryClientConfig
	AuthSessionRpcClient *service_discovery.ServiceDiscoveryClientConfig
	// BotsRpcClient        *service_discovery.ServiceDiscoveryClientConfig
}

func init() {
	tomlPath := util.GetWorkingDirectory() + "/biz_server.toml"
	flag.StringVar(&confPath, "conf", tomlPath, "config path")
}

func InitializeConfig() (err error) {
	_, err = toml.DecodeFile(confPath, &Conf)
	if err != nil {
		err = fmt.Errorf("decode file %s error: %v", confPath, err)
	}
	return
}

////////////////////////////////////////////////////////////////////////////////////////////////////
// messenger_server.go
type messengerServer struct {
	rpcServer *grpc_util.RPCServer
	models    []core.CoreModel
}

func newMessengerServer() *messengerServer {
	return &messengerServer{}
}

// AppInstance interface
func (s *messengerServer) Initialize() error {
	glog.Infof("messengerServer - initialize...")

	err := InitializeConfig()
	if err != nil {
		glog.Fatal(err)
		return err
	}
	glog.Info("messengerServer - load conf: ", Conf)

	s.models = core.InstallCoreModels(Conf.ServerId, func() {
		// 初始化mysql_client、redis_client
		redis_client.InstallRedisClientManager(Conf.Redis)
		mysql_client.InstallMysqlClientManager(Conf.Mysql)

		// 初始化redis_dao、mysql_dao
		dao.InstallMysqlDAOManager(mysql_client.GetMysqlClientManager())
		dao.InstallRedisDAOManager(redis_client.GetRedisClientManager())

		document_client.InstallNbfsClient(Conf.NbfsRpcClient)
		sync_client.InstallSyncClient(Conf.SyncRpcClient2)
		auth_session_client.InstallAuthSessionClient(Conf.AuthSessionRpcClient)
		// bots_client.InstallBotsClient(Conf.BotsRpcClient)
	})

	s.rpcServer = grpc_util.NewRpcServer(Conf.RpcServer.Addr, &Conf.RpcServer.RpcDiscovery)

	return nil
}

func (s *messengerServer) RunLoop() {
	glog.Infof("messengerServer - runLoop...")

	// TODO(@benqi): check error
	s.rpcServer.Serve(func(s2 *grpc.Server) {
		mtproto.RegisterRPCAccountServer(s2, account.NewAccountServiceImpl(s.models))
		mtproto.RegisterRPCAuthServer(s2, auth.NewAuthServiceImpl(s.models))
		mtproto.RegisterRPCBotsServer(s2, bots.NewBotsServiceImpl(s.models))
		mtproto.RegisterRPCChannelsServer(s2, channels.NewChannelsServiceImpl(s.models))
		mtproto.RegisterRPCContactsServer(s2, contacts.NewContactsServiceImpl(s.models))
		mtproto.RegisterRPCHelpServer(s2, help.NewHelpServiceImpl(s.models))
		mtproto.RegisterRPCLangpackServer(s2, langpack.NewLangpackServiceImpl(s.models))
		mtproto.RegisterRPCMessagesServer(s2, messages.NewMessagesServiceImpl(s.models))
		mtproto.RegisterRPCPaymentsServer(s2, payments.NewPaymentsServiceImpl(s.models))
		mtproto.RegisterRPCPhoneServer(s2, phone.NewPhoneServiceImpl(s.models, Conf.RelayIp))
		mtproto.RegisterRPCPhotosServer(s2, photos.NewPhotosServiceImpl(s.models))
		mtproto.RegisterRPCStickersServer(s2, stickers.NewStickersServiceImpl(s.models))
		mtproto.RegisterRPCUpdatesServer(s2, updates.NewUpdatesServiceImpl(s.models))
		mtproto.RegisterRPCUsersServer(s2, users.NewUsersServiceImpl(s.models))
	})
}

func (s *messengerServer) Destroy() {
	glog.Infof("messengerServer - destroy...")
	//s.server.Stop()
	s.rpcServer.Stop()
	//time.Sleep(1*time.Second)
}

////////////////////////////////////////////////////////////////////////////////////////////////////
// main
func main() {
	flag.Parse()

	instance := newMessengerServer()
	util.DoMainAppInstance(instance)
}
