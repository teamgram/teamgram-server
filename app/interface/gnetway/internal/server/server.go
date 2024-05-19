// Copyright 2022 Teamgram Authors
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
//
// Author: teamgramio (teamgram.io@gmail.com)
//

package server

import (
	"flag"

	"github.com/teamgram/teamgram-server/app/interface/gnetway/internal/config"
	"github.com/teamgram/teamgram-server/app/interface/gnetway/internal/server/gnet"
	"github.com/teamgram/teamgram-server/app/interface/gnetway/internal/server/grpc"
	"github.com/teamgram/teamgram-server/app/interface/gnetway/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/zrpc"
)

var (
	configFile = flag.String("f", "etc/gateway.yaml", "the config file")
)

type Server struct {
	grpcSrv *zrpc.RpcServer
	server  *gnet.Server
}

func (s *Server) Initialize() error {
	var c config.Config
	conf.MustLoad(*configFile, &c)

	logx.Infov(c)

	ctx := svc.NewServiceContext(c)
	s.server = gnet.New(ctx, c)
	s.grpcSrv = grpc.New(ctx, c.RpcServerConf, s.server)

	go func() {
		s.grpcSrv.Start()
	}()

	return nil
}

func (s *Server) RunLoop() {
	// s.server.Serve()
	//if err := s.server.Serve(); err != nil {
	//	logx.Errorf("run server error: %v, quit...", err)
	//	commands.GSignal <- syscall.SIGQUIT
	//}
}

func (s *Server) Destroy() {
	s.grpcSrv.Stop()
	s.server.Close()
}
