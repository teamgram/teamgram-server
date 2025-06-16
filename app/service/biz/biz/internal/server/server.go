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
	"github.com/cloudwego/kitex/server"
	"github.com/teamgram/teamgram-server/v2/app/service/biz/chat/chat/chatservice"
	"github.com/teamgram/teamgram-server/v2/app/service/biz/code/code/codeservice"
	"github.com/teamgram/teamgram-server/v2/app/service/biz/dialog/dialog/dialogservice"
	"github.com/teamgram/teamgram-server/v2/app/service/biz/message/message/messageservice"
	"github.com/teamgram/teamgram-server/v2/app/service/biz/updates/updates/updatesservice"
	"github.com/teamgram/teamgram-server/v2/app/service/biz/user/user/userservice"
	"github.com/teamgram/teamgram-server/v2/app/service/biz/username/username/usernameservice"
	"github.com/teamgram/teamgram-server/v2/pkg/net/kitex"

	"github.com/teamgram/teamgram-server/v2/app/service/biz/biz/internal/config"
	chathelper "github.com/teamgram/teamgram-server/v2/app/service/biz/chat"
	codehelper "github.com/teamgram/teamgram-server/v2/app/service/biz/code"
	dialoghelper "github.com/teamgram/teamgram-server/v2/app/service/biz/dialog"
	messagehelper "github.com/teamgram/teamgram-server/v2/app/service/biz/message"
	updateshelper "github.com/teamgram/teamgram-server/v2/app/service/biz/updates"
	userhelper "github.com/teamgram/teamgram-server/v2/app/service/biz/user"
	usernamehelper "github.com/teamgram/teamgram-server/v2/app/service/biz/username"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
)

var configFile = flag.String("f", "etc/biz.yaml", "the config file")

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
			_ = chatservice.RegisterService(
				s,
				chathelper.New(
					chathelper.Config{
						RpcServerConf: c.RpcServerConf,
						//Mysql:         c.Mysql,
						//Cache:         c.Cache,
						//MediaClient:   c.MediaClient,
					}))

			// codehelper
			_ = codeservice.RegisterService(
				s,
				codehelper.New(codehelper.Config{
					RpcServerConf: c.RpcServerConf,
					//Mysql:         c.Mysql,
					//Cache:         c.Cache,
					//KV:            c.KV,
				}))

			// dialoghelper
			_ = dialogservice.RegisterService(
				s,
				dialoghelper.New(dialoghelper.Config{
					RpcServerConf: c.RpcServerConf,
					//Mysql:         c.Mysql,
					//Cache:         c.Cache,
				}))

			// messagehelper
			_ = messageservice.RegisterService(
				s,
				messagehelper.New(
					messagehelper.Config{
						RpcServerConf: c.RpcServerConf,
						//Mysql:           c.Mysql,
						//Cache:           c.Cache,
						//MessageSharding: c.MessageSharding,
					}))

			// updateshelper
			_ = updatesservice.RegisterService(
				s,
				updateshelper.New(updateshelper.Config{
					RpcServerConf: c.RpcServerConf,
					//Mysql:         c.Mysql,
					//KV:            c.KV,
					//IdgenClient:   c.IdgenClient,
				}))

			// userhelper
			_ = userservice.RegisterService(
				s,
				userhelper.New(userhelper.Config{
					RpcServerConf: c.RpcServerConf,
					//Mysql:         c.Mysql,
					//Cache:         c.Cache,
					//MediaClient:   c.MediaClient,
				}))

			// usernamehelper
			_ = usernameservice.RegisterService(
				s,
				usernamehelper.New(usernamehelper.Config{
					RpcServerConf: c.RpcServerConf,
					//Mysql:         c.Mysql,
					//Cache:         c.Cache,
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
