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
	"github.com/nebula-chat/chatengine/service/idgen/proto"
	"google.golang.org/grpc"
	"log"
)

type idgenServer struct {
	rpcServer *grpc_util.RPCServer
}

func NewIDGenServer() *idgenServer {
	return &idgenServer{}
}

func (s *idgenServer) Initialize() error {
	var err error

	if err = InitializeConfig(); err != nil {
		glog.Error("decode config file error: ", err)
		return err
	}

	glog.Infof("config loaded: %v", Conf)

	s.rpcServer = grpc_util.NewRpcServer(Conf.Server.Addr, &Conf.Discovery)

	return err
}

func (s *idgenServer) RunLoop() {
	go s.rpcServer.Serve(func(s2 *grpc.Server) {
		rpcServer, err := newIDGenServiceImpl(Conf.Etcd)
		if err != nil {
			log.Fatal(err)
		}
		rpcServer.init()
		seqsvr.RegisterRPCIDGenServer(s2, rpcServer)
	})
}

func (s *idgenServer) Destroy() {
	s.rpcServer.Stop()
}
