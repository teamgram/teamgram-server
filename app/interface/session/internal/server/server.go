/*
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2021-present,  Teamgram Studio (https://teamgram.io).
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package server

import (
	"flag"
	"time"

	"github.com/teamgram/teamgram-server/app/interface/session/internal/config"
	"github.com/teamgram/teamgram-server/app/interface/session/internal/server/grpc"
	"github.com/teamgram/teamgram-server/app/interface/session/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/zrpc"
)

var configFile = flag.String("f", "etc/session.yaml", "the config file")

type Server struct {
	grpcSrv *zrpc.RpcServer
	svcCtx  *svc.ServiceContext
}

func New() *Server {
	return new(Server)
}

func (s *Server) Initialize() error {
	var c config.Config
	conf.MustLoad(*configFile, &c)

	logx.Infov(c)

	s.svcCtx = svc.NewServiceContext(c)
	s.grpcSrv = grpc.New(s.svcCtx, c.RpcServerConf)

	go func() {
		s.grpcSrv.Start()
	}()
	return nil
}

func (s *Server) RunLoop() {
}

func (s *Server) Destroy() {
	// 优雅排空：先等待进行中的 RPC 请求处理完毕（最多 30s），再停止 gRPC 服务
	logx.Infof("session server destroying, draining auth wrappers...")
	s.svcCtx.MainAuthMgr.Drain(30 * time.Second)

	s.grpcSrv.Stop()
}
