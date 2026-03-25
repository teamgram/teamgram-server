// Copyright 2024 Teamgooo Authors
//  All rights reserved.
//
// Author: Benqi (wubenqi@gmail.com)
//

package server

import (
	"flag"
	"log"

	"github.com/teamgram/teamgram-server/v2/pkg/proto/rpc/codec"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/rpc/examples/echo/echo/echo"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/rpc/examples/echo/internal/config"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/rpc/examples/echo/internal/server/tg/service"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/rpc/examples/echo/internal/svc"

	"github.com/cloudwego/kitex/server"
)

var configFile = flag.String("f", "etc/echo.yaml", "the config file")

type Server struct {
	server.Server
}

func New() *Server {
	return new(Server)
}

func (s *Server) Initialize() error {
	var c config.Config
	ctx := svc.NewServiceContext(c)
	_ = ctx

	cCodec := codec.NewZRpcCodec(true)
	s.Server = echo.NewServer(service.New(ctx), server.WithCodec(cCodec))
	return nil
}

func (s *Server) RunLoop() {
	if err := s.Server.Run(); err != nil {
		log.Println("server stopped with error:", err)
	} else {
		log.Println("server stopped")
	}
}

func (s *Server) Destroy() {
}
