/*
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2021-present,  Teamgram Studio (https://teamgram.io).
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package grpc

import (
	"github.com/teamgram/teamgram-server/app/service/biz/webpage/internal/server/grpc/service"
	"github.com/teamgram/teamgram-server/app/service/biz/webpage/internal/svc"
	"github.com/teamgram/teamgram-server/app/service/biz/webpage/webpage"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

// New new a grpc server.
func New(ctx *svc.ServiceContext, c zrpc.RpcServerConf) *zrpc.RpcServer {
	s, err := zrpc.NewServer(c, func(grpcServer *grpc.Server) {
		webpage.RegisterRPCWebpageServer(grpcServer, service.New(ctx))
	})
	logx.Must(err)
	return s
}
