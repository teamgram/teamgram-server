/*
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2021-present,  Teamgram Studio (https://teamgram.io).
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package dao

import (
	"github.com/teamgram/teamgram-server/app/service/dfs/internal/config"
	"github.com/teamgram/teamgram-server/app/service/dfs/internal/minio_util"
	idgen_client "github.com/teamgram/teamgram-server/app/service/idgen/client"

	"github.com/minio/minio-go/v7"
	"github.com/zeromicro/go-zero/core/stores/kv"
	"github.com/zeromicro/go-zero/zrpc"
)

type Dao struct {
	minio *minio.Core
	idgen_client.IDGenClient2
	ssdb kv.Store
}

func New(c config.Config) *Dao {
	return &Dao{
		minio:        minio_util.MustNewMinioClient(&c.Minio),
		IDGenClient2: idgen_client.NewIDGenClient2(zrpc.MustNewClient(c.IdGen)),
		ssdb:         kv.NewStore(c.SSDB),
	}
}

func NewDFSHelper(minio *minio_util.MinioConfig, idgen zrpc.RpcClientConf, ssdb kv.KvConf) *Dao {
	return &Dao{
		minio:        minio_util.MustNewMinioClient(minio),
		IDGenClient2: idgen_client.NewIDGenClient2(zrpc.MustNewClient(idgen)),
		ssdb:         kv.NewStore(ssdb),
	}
}

func NewMinioHelper(minio *minio_util.MinioConfig) minio_util.MinioHelper {
	return &Dao{
		minio: minio_util.MustNewMinioClient(minio),
	}
}
