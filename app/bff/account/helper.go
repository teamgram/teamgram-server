/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright 2022 Teamgram Authors.
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package account_helper

import (
	"github.com/teamgram/teamgram-server/app/bff/account/internal/config"
	"github.com/teamgram/teamgram-server/app/bff/account/internal/server/grpc/service"
	"github.com/teamgram/teamgram-server/app/bff/account/internal/svc"
	"github.com/teamgram/teamgram-server/app/bff/account/plugin"
	"github.com/teamgram/teamgram-server/pkg/code"
)

type (
	Config = config.Config
)

func New(c Config, code2 code.VerifyCodeInterface, plugin plugin.AuthorizationPlugin) *service.Service {
	return service.New(svc.NewServiceContext(c, code2, plugin))
}
