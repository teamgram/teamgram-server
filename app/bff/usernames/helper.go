/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright 2022 Teamgram Authors.
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package usernames_helper

import (
	"github.com/teamgram/teamgram-server/app/bff/usernames/internal/config"
	"github.com/teamgram/teamgram-server/app/bff/usernames/internal/server/grpc/service"
	"github.com/teamgram/teamgram-server/app/bff/usernames/internal/svc"
	"github.com/teamgram/teamgram-server/app/bff/usernames/plugin"
)

type (
	Config = config.Config
)

func New(c Config, plugin plugin.UsernamesPlugin) *service.Service {
	return service.New(svc.NewServiceContext(c, plugin))
}
