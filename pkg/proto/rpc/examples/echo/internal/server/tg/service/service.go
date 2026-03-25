// Copyright 2024 Teamgooo Authors
//  All rights reserved.
//
// Author: Benqi (wubenqi@gmail.com)
//

package service

import (
	"github.com/teamgram/teamgram-server/v2/pkg/proto/rpc/examples/echo/internal/svc"
)

type Service struct {
	svcCtx *svc.ServiceContext
}

func (s *Service) GetServiceContext() *svc.ServiceContext {
	return s.svcCtx
}

func New(ctx *svc.ServiceContext) *Service {
	return &Service{
		svcCtx: ctx,
	}
}
