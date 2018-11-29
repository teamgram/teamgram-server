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
	"github.com/nebula-chat/chatengine/pkg/grpc_util"
	"github.com/nebula-chat/chatengine/mtproto"
	"google.golang.org/grpc"
)

type uploadServer struct {
	rpcServer *grpc_util.RPCServer
}

func NewUploadServer() *uploadServer {
	return &uploadServer{}
}

// AppInstance interface
func (s *uploadServer) Initialize() error {
	glog.Infof("uploadServer - initialize...")

	err := InitializeConfig()
	if err != nil {
		glog.Fatal(err)
		return err
	}
	glog.Info("uploadServer - load conf: ", Conf)

	// cachefs.InitCacheFS(Conf.DataPath)
	// 初始化mysql_client、redis_client
	// redis_client.InstallRedisClientManager(Conf.Redis)
	// mysql_client.InstallMysqlClientManager(Conf.Mysql)
	s.rpcServer = grpc_util.NewRpcServer(Conf.RpcServer.Addr, &Conf.RpcServer.RpcDiscovery)
	return nil
}

func (s *uploadServer) RunLoop() {
	glog.Infof("uploadServer - runLoop...")

	// TODO(@benqi): check error
	s.rpcServer.Serve(func(s2 *grpc.Server) {
		impl := NewUploadServiceImpl(Conf.ServerId, Conf.DataPath)
		mtproto.RegisterRPCUploadServer(s2, impl)
	})
}

func (s *uploadServer) Destroy() {
	glog.Infof("uploadServer - destroy...")
	s.rpcServer.Stop()
	//time.Sleep(1*time.Second)
}
