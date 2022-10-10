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

	"github.com/teamgram/teamgram-server/app/service/biz/biz/internal/config"
	chat_helper "github.com/teamgram/teamgram-server/app/service/biz/chat"
	"github.com/teamgram/teamgram-server/app/service/biz/chat/chat"
	code_helper "github.com/teamgram/teamgram-server/app/service/biz/code"
	"github.com/teamgram/teamgram-server/app/service/biz/code/code"
	dialog_helper "github.com/teamgram/teamgram-server/app/service/biz/dialog"
	"github.com/teamgram/teamgram-server/app/service/biz/dialog/dialog"
	message_helper "github.com/teamgram/teamgram-server/app/service/biz/message"
	"github.com/teamgram/teamgram-server/app/service/biz/message/message"
	updates_helper "github.com/teamgram/teamgram-server/app/service/biz/updates"
	"github.com/teamgram/teamgram-server/app/service/biz/updates/updates"
	user_helper "github.com/teamgram/teamgram-server/app/service/biz/user"
	"github.com/teamgram/teamgram-server/app/service/biz/user/user"
	username_helper "github.com/teamgram/teamgram-server/app/service/biz/username"
	"github.com/teamgram/teamgram-server/app/service/biz/username/username"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

var configFile = flag.String("f", "etc/biz.yaml", "the config file")

type Server struct {
	grpcSrv *zrpc.RpcServer
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

	s.grpcSrv = zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		// chat_helper
		chat.RegisterRPCChatServer(
			grpcServer,
			chat_helper.New(
				chat_helper.Config{
					RpcServerConf: c.RpcServerConf,
					Mysql:         c.Mysql,
					Cache:         c.Cache,
					MediaClient:   c.MediaClient,
				},
				nil))

		// code_helper
		code.RegisterRPCCodeServer(
			grpcServer,
			code_helper.New(code_helper.Config{
				RpcServerConf: c.RpcServerConf,
				Mysql:         c.Mysql,
				Cache:         c.Cache,
				KV:            c.KV,
			}))

		// dialog_helper
		dialog.RegisterRPCDialogServer(
			grpcServer,
			dialog_helper.New(dialog_helper.Config{
				RpcServerConf: c.RpcServerConf,
				Mysql:         c.Mysql,
				Cache:         c.Cache,
			}))

		// message_helper
		message.RegisterRPCMessageServer(
			grpcServer,
			message_helper.New(
				message_helper.Config{
					RpcServerConf:   c.RpcServerConf,
					Mysql:           c.Mysql,
					Cache:           c.Cache,
					MessageSharding: c.MessageSharding,
				},
				nil))

		// updates_helper
		updates.RegisterRPCUpdatesServer(
			grpcServer,
			updates_helper.New(updates_helper.Config{
				RpcServerConf: c.RpcServerConf,
				Mysql:         c.Mysql,
				KV:            c.KV,
				IdgenClient:   c.IdgenClient,
			}))

		// user_helper
		user.RegisterRPCUserServer(
			grpcServer,
			user_helper.New(user_helper.Config{
				RpcServerConf: c.RpcServerConf,
				Mysql:         c.Mysql,
				Cache:         c.Cache,
				MediaClient:   c.MediaClient,
			}))

		// username_helper
		username.RegisterRPCUsernameServer(
			grpcServer,
			username_helper.New(username_helper.Config{
				RpcServerConf: c.RpcServerConf,
				Mysql:         c.Mysql,
				Cache:         c.Cache,
			}))
	})

	// logx.Must(err)
	go func() {
		s.grpcSrv.Start()
	}()
	return nil
}

func (s *Server) RunLoop() {
}

func (s *Server) Destroy() {
	s.grpcSrv.Stop()
}
