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

	kafka "github.com/teamgram/marmota/pkg/mq"
	inbox_helper "github.com/teamgram/teamgram-server/app/messenger/msg/inbox"
	"github.com/teamgram/teamgram-server/app/messenger/msg/internal/config"
	msg_helper "github.com/teamgram/teamgram-server/app/messenger/msg/msg"
	"github.com/teamgram/teamgram-server/app/messenger/msg/msg/msg"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

var configFile = flag.String("f", "etc/msg.yaml", "the config file")

type Server struct {
	grpcSrv *zrpc.RpcServer
	mq      *kafka.ConsumerGroup
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
		// msg_helper
		msg.RegisterRPCMsgServer(
			grpcServer,
			msg_helper.New(
				msg_helper.Config{
					RpcServerConf:   c.RpcServerConf,
					Mysql:           c.Mysql,
					KV:              c.KV,
					IdgenClient:     c.IdgenClient,
					UserClient:      c.BizServiceClient,
					ChatClient:      c.BizServiceClient,
					SyncClient:      c.SyncClient,
					InboxClient:     c.InboxClient,
					DialogClient:    c.BizServiceClient,
					MessageSharding: c.MessageSharding,
					Redis2:          c.Redis2,
					UsernameClient:  c.BizServiceClient,
				}, nil))
	})

	go func() {
		s.grpcSrv.Start()
	}()

	s.mq = inbox_helper.New(inbox_helper.Config{
		RpcServerConf:   c.RpcServerConf,
		InboxConsumer:   c.InboxConsumer,
		Mysql:           c.Mysql,
		KV:              c.KV,
		IdgenClient:     c.IdgenClient,
		UserClient:      c.BizServiceClient,
		ChatClient:      c.BizServiceClient,
		SyncClient:      c.SyncClient,
		BotSyncClient:   c.BotSyncClient,
		DialogClient:    c.BizServiceClient,
		MessageSharding: c.MessageSharding,
	})

	go func() {
		s.mq.Start()
	}()

	return nil
}

func (s *Server) RunLoop() {
}

func (s *Server) Destroy() {
	s.grpcSrv.Stop()
}
