/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright 2022 Teamgram Authors.
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package media_helper

import (
	"github.com/teamgram/teamgram-server/app/service/media/internal/config"
	"github.com/teamgram/teamgram-server/app/service/media/internal/dal/dataobject"
	"github.com/teamgram/teamgram-server/app/service/media/internal/dao"
	"github.com/teamgram/teamgram-server/app/service/media/internal/server/grpc/service"
	"github.com/teamgram/teamgram-server/app/service/media/internal/svc"
)

type (
	Dao = dao.Dao
)

type (
	DocumentsDO = dataobject.DocumentsDO
)

var (
	GenCacheDocumentKey   = dao.GenCacheDocumentKey
	ParseCacheDocumentKey = dao.ParseCacheDocumentKey
	GetPhotoSize          = dao.GetPhotoSize
)

type (
	Config  = config.Config
	Service = service.Service
)

var (
	NewServiceContext = svc.NewServiceContext
)

func New(c Config) *service.Service {
	return service.New(svc.NewServiceContext(c))
}
