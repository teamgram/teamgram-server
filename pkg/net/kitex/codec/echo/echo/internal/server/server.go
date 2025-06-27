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

	"github.com/teamgram/teamgram-server/v2/pkg/net/kitex"
	"github.com/teamgram/teamgram-server/v2/pkg/net/kitex/codec/echo/echo/internal/config"
	echo1helper "github.com/teamgram/teamgram-server/v2/pkg/net/kitex/codec/echo/echo1"
	"github.com/teamgram/teamgram-server/v2/pkg/net/kitex/codec/echo/echo1/echo1/echo1service"
	echo2helper "github.com/teamgram/teamgram-server/v2/pkg/net/kitex/codec/echo/echo2"
	"github.com/teamgram/teamgram-server/v2/pkg/net/kitex/codec/echo/echo2/echo2/echo2service"

	"github.com/cloudwego/kitex/server"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
)

var configFile = flag.String("f", "etc/echo.yaml", "the config file")

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
	// ctx := svc.NewServiceContext(c)
	// s.grpcSrv = grpc.New(ctx, c.RpcServerConf)

	s.kitexSrv = kitex.MustNewServer(
		c.RpcServerConf,
		func(s server.Server) error {
			// chathelper
			_ = echo1service.RegisterService(
				s,
				echo1helper.New(
					echo1helper.Config{
						RpcServerConf: c.RpcServerConf,
						//Mysql:         c.Mysql,
						//Cache:         c.Cache,
						//MediaClient:   c.MediaClient,
					}))

			// codehelper
			_ = echo2service.RegisterService(
				s,
				echo2helper.New(echo2helper.Config{
					RpcServerConf: c.RpcServerConf,
					//Mysql:         c.Mysql,
					//Cache:         c.Cache,
					//KV:            c.KV,
				}))

			return nil
		})

	// logx.Must(err)
	go func() {
		_ = s.kitexSrv.Run()
	}()
	return nil
}

func (s *Server) RunLoop() {
}

func (s *Server) Destroy() {
	_ = s.kitexSrv.Stop()
}
