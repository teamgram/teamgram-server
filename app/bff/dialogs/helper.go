/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright 2024 Teamgram Authors.
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package dialogs_helper

import (
	"github.com/teamgram/teamgram-server/app/bff/dialogs/internal/config"
	"github.com/teamgram/teamgram-server/app/bff/dialogs/internal/server/grpc/service"
	"github.com/teamgram/teamgram-server/app/bff/dialogs/internal/svc"
	"github.com/teamgram/teamgram-server/app/bff/dialogs/plugin"
)

type (
	Config = config.Config
)

func New(c Config, plugin plugin.DialogsPlugin) *service.Service {
	return service.New(svc.NewServiceContext(c, plugin))
}
