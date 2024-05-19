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
	"github.com/teamgram/teamgram-server/app/interface/gateway/gateway"
	"github.com/teamgram/teamgram-server/app/interface/gnetway/internal/server/grpc/service"
	"github.com/teamgram/teamgram-server/app/interface/gnetway/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

// New new a grpc server.
func New(svcCtx *svc.ServiceContext, c zrpc.RpcServerConf, srv gateway.RPCGatewayServer) *zrpc.RpcServer {
	s, err := zrpc.NewServer(c, func(grpcServer *grpc.Server) {
		gateway.RegisterRPCGatewayServer(grpcServer, service.New(svcCtx, srv))
	})
	logx.Must(err)
	return s
}
