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

package grpc_util

import (
	"github.com/golang/glog"
	"github.com/nebula-chat/chatengine/pkg/grpc_util/service_discovery"
	"github.com/nebula-chat/chatengine/pkg/etcd_util"
	"github.com/nebula-chat/chatengine/pkg/grpc_util/middleware/recovery2"
	"github.com/nebula-chat/chatengine/pkg/grpc_util/service_discovery/etcd3"
	"google.golang.org/grpc"
	"net"
	"os"
	"os/signal"
	"syscall"
)

type RPCServer struct {
	addr     string
	registry *etcd3.EtcdReigistry
	s        *grpc.Server
}

func NewRpcServer(addr string, discovery *service_discovery.ServiceDiscoveryServerConfig) *RPCServer {
	s := &RPCServer{
		addr: addr,
	}

	var err error
	s.registry, err = etcd_util.NewEtcdRegistry(*discovery)
	if err != nil {
		glog.Fatal(err)
	}
	s.s = grpc_recovery2.NewRecoveryServer2(BizUnaryRecoveryHandler, BizUnaryRecoveryHandler2, BizStreamRecoveryHandler)

	return s
}

// type func RegisterRPCServerHandler(s *grpc.Server)
type RegisterRPCServerFunc func(s *grpc.Server)

func (s *RPCServer) Serve(regFunc RegisterRPCServerFunc) {
	// defer s.GracefulStop()
	listener, err := net.Listen("tcp", s.addr)

	if err != nil {
		glog.Error("failed to listen: %v", err)
		return
	}
	glog.Infof("rpc listening on:%s", s.addr)

	if regFunc != nil {
		regFunc(s.s)
	}

	defer s.s.GracefulStop()
	go s.registry.Register()

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGTERM, syscall.SIGINT, syscall.SIGKILL, syscall.SIGHUP, syscall.SIGQUIT)
	go func() {
		s2 := <-ch
		glog.Infof("exit...")
		s.registry.Deregister()
		if i, ok := s2.(syscall.Signal); ok {
			os.Exit(int(i))
		} else {
			os.Exit(0)
		}

	}()

	if err := s.s.Serve(listener); err != nil {
		glog.Fatalf("failed to serve: %s", err)
	}
}

func (s *RPCServer) Stop() {
	s.s.GracefulStop()
}
