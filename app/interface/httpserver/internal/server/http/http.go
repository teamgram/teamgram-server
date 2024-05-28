// Copyright (c) 2024-present, Teamgram Studio (https://teamgram.io).
//  All rights reserved.
//
// Author: teamgramio (teamgram.io@gmail.com)
//

package http

import (
	"github.com/teamgram/teamgram-server/app/interface/httpserver/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

// New new a grpc server.
func New(ctx *svc.ServiceContext, c rest.RestConf) *rest.Server {
	srv := rest.MustNewServer(c)

	go func(s *rest.Server, c *svc.ServiceContext) {
		defer s.Stop()

		RegisterHandlers(s, c)

		s.Start()
	}(srv, ctx)

	return srv
}
