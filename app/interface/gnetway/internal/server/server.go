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

	"github.com/teamgram/teamgram-server/v2/app/interface/gnetway/gnetway/gnetwayservice"
	"github.com/teamgram/teamgram-server/v2/app/interface/gnetway/internal/config"
	"github.com/teamgram/teamgram-server/v2/app/interface/gnetway/internal/server/gnet"
	"github.com/teamgram/teamgram-server/v2/app/interface/gnetway/internal/svc"
	"github.com/teamgram/teamgram-server/v2/pkg/net/kitex"

	"github.com/cloudwego/kitex/server"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
)

var (
	configFile = flag.String("f", "etc/gateway.yaml", "the config file")
)

func New() *Server {
	return new(Server)
}

type Server struct {
	kitexSrv *kitex.RpcServer
	server   *gnet.Server
}

func (s *Server) Initialize() error {
	var c config.Config
	conf.MustLoad(*configFile, &c)

	logx.Infov(c)

	ctx := svc.NewServiceContext(c)
	s.server = gnet.New(ctx, c)

	s.kitexSrv = kitex.MustNewServer(
		c.RpcServerConf,
		func(s2 server.Server) error {
			return gnetwayservice.RegisterService(s2, s.server)
		})

	go func() {
		s.kitexSrv.Run()
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
	s.kitexSrv.Stop()
	s.server.Close()
}
