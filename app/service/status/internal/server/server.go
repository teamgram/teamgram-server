// Copyright 2024 Teamgram Authors
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
	"github.com/teamgram/teamgram-server/v2/app/service/status/status/statusservice"

	"github.com/teamgram/teamgram-server/v2/app/service/status/internal/config"
	"github.com/teamgram/teamgram-server/v2/app/service/status/internal/server/tg/service"
	"github.com/teamgram/teamgram-server/v2/app/service/status/internal/svc"
	"github.com/teamgram/teamgram-server/v2/pkg/net/kitex"

	"github.com/cloudwego/kitex/server"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
)

var configFile = flag.String("f", "etc/status.yaml", "the config file")

type Server struct {
	kitexSrv *kitex.RpcServer
}

func New() *Server {
	return new(Server)
}

func (s *Server) Initialize() error {
	var c config.Config
	conf.MustLoad(*configFile, &c)

	logx.Infov(c)

	ctx := svc.NewServiceContext(c)
	_ = ctx

	s.kitexSrv = kitex.MustNewServer(
		c.RpcServerConf,
		func(s server.Server) error {
			return statusservice.RegisterService(s, service.New(ctx))
		})

	return nil
}

func (s *Server) RunLoop() {
	if err := s.kitexSrv.Run(); err != nil {
		// log.Println("server stopped with error:", err)
	} else {
		// log.Println("server stopped")
	}
}

func (s *Server) Destroy() {
	if err := s.kitexSrv.Stop(); err != nil {
		// log.Println("server stopped with error:", err)
	} else {
		// log.Println("server stopped")
	}
}
