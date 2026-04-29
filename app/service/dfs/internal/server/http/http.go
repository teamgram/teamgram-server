package http

import (
	"context"
	"net/http"

	"github.com/teamgram/teamgram-server/v2/app/service/dfs/internal/core"
	"github.com/teamgram/teamgram-server/v2/app/service/dfs/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

type Server struct {
	srv *rest.Server
}

func New(ctx context.Context, svcCtx *svc.ServiceContext, c rest.RestConf) (*Server, error) {
	srv, err := rest.NewServer(c)
	if err != nil {
		return nil, err
	}
	opener := core.New(ctx, svcCtx)
	srv.AddRoutes([]rest.Route{
		{
			Method:  http.MethodGet,
			Path:    "/dfs/file/:file",
			Handler: GetDfsFile(opener),
		},
	})
	return &Server{srv: srv}, nil
}

func (s *Server) Start() error {
	if s == nil || s.srv == nil {
		return nil
	}
	s.srv.Start()
	return nil
}

func (s *Server) Stop() {
	if s == nil || s.srv == nil {
		return
	}
	s.srv.Stop()
}
