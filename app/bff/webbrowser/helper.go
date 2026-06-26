/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2026 The Teamgram Authors (https://teamgram.net).
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package webbrowserhelper

import (
	"github.com/teamgram/teamgram-server/app/bff/webbrowser/internal/config"
	"github.com/teamgram/teamgram-server/app/bff/webbrowser/internal/server/grpc/service"
	"github.com/teamgram/teamgram-server/app/bff/webbrowser/internal/svc"
)

type (
	Config = config.Config
)

func New(c Config) *service.Service {
	return service.New(svc.NewServiceContext(c))
}
