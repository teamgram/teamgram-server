/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright 2022 Teamgram Authors.
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package chats_helper

import (
	"github.com/teamgram/teamgram-server/app/bff/chats/internal/config"
	"github.com/teamgram/teamgram-server/app/bff/chats/internal/server/grpc/service"
	"github.com/teamgram/teamgram-server/app/bff/chats/internal/svc"
)

type (
	Config = config.Config
)

func New(c Config) *service.Service {
	return service.New(svc.NewServiceContext(c))
}
