/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright 2025 Teamgram Authors.
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package passkeyhelper

import (
	"github.com/teamgram/teamgram-server/app/bff/passkey/internal/config"
	"github.com/teamgram/teamgram-server/app/bff/passkey/internal/server/grpc/service"
	"github.com/teamgram/teamgram-server/app/bff/passkey/internal/svc"
)

type (
	Config = config.Config
)

func New(c Config) *service.Service {
	return service.New(svc.NewServiceContext(c))
}
