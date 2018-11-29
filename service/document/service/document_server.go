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

package service

import (
	"github.com/golang/glog"
	"github.com/nebula-chat/chatengine/pkg/grpc_util"
	"github.com/nebula-chat/chatengine/pkg/mysql_client"
	"github.com/nebula-chat/chatengine/pkg/redis_client"
	"github.com/nebula-chat/chatengine/mtproto"
	"google.golang.org/grpc"
)

type documentServer struct {
	rpcServer *grpc_util.RPCServer
}

func NewDocumentServer() *documentServer {
	return &documentServer{}
}

// AppInstance interface
func (s *documentServer) Initialize() error {
	glog.Infof("documentServer - initialize...")

	err := InitializeConfig()
	if err != nil {
		glog.Fatal(err)
		return err
	}
	glog.Info("documentServer - load conf: ", Conf)

	// 初始化mysql_client、redis_client
	redis_client.InstallRedisClientManager(Conf.Redis)
	mysql_client.InstallMysqlClientManager(Conf.Mysql)

	s.rpcServer = grpc_util.NewRpcServer(Conf.RpcServer.Addr, &Conf.RpcServer.RpcDiscovery)
	return nil
}

func (s *documentServer) RunLoop() {
	glog.Infof("documentServer - runLoop...")

	// TODO(@benqi): check error
	s.rpcServer.Serve(func(s2 *grpc.Server) {
		impl := NewDocumentServiceImpl(Conf.ServerId, Conf.DataPath, "immaster", "cache")
		mtproto.RegisterRPCNbfsServer(s2, impl)
	})
}

func (s *documentServer) Destroy() {
	glog.Infof("documentServer - destroy...")
	s.rpcServer.Stop()
	//time.Sleep(1*time.Second)
}
