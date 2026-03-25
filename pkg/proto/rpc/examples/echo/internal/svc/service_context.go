// Copyright 2024 Teamgooo Authors
//  All rights reserved.
//
// Author: Benqi (wubenqi@gmail.com)
//

package svc

import (
	"github.com/teamgram/teamgram-server/v2/pkg/proto/rpc/examples/echo/internal/config"
)

type ServiceContext struct {
	Config config.Config
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
	}
}
