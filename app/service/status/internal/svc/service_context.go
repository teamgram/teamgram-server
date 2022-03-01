/*
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2021-present,  Teamgram Studio (https://teamgram.io).
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package svc

import (
	"github.com/teamgram/teamgram-server/app/service/status/internal/config"
	"github.com/zeromicro/go-zero/core/stores/kv"
)

type ServiceContext struct {
	Config config.Config
	KV     kv.Store
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		KV:     kv.NewStore(c.Status),
	}
}
