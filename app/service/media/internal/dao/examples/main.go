// Copyright 2022 Teamgram Authors
//  All rights reserved.
//
// Author: Benqi (wubenqi@gmail.com)
//

package main

import (
	"context"
	"flag"
	"fmt"

	"github.com/teamgram/marmota/pkg/stores/sqlx"
	"github.com/teamgram/proto/mtproto"
	"github.com/teamgram/teamgram-server/app/service/media/internal/config"
	"github.com/teamgram/teamgram-server/app/service/media/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
)

var (
	configFile = flag.String("f", "etc/media.yaml", "the config file")
)

var (
	cacheId    = int64(1401151020526600192)
	cacheIdKey = "document#1401151020526600192"
)

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	svcCtx := svc.NewServiceContext(c)
	GetCacheDocument(svcCtx)
}

func GetCacheDocument(svcCtx *svc.ServiceContext) {
	ctx := context.Background()

	// fmt.Println(document)

	var (
		document2 *mtproto.Document
	)

	svcCtx.Dao.CachedConn.QueryRow(
		ctx,
		&document2,
		cacheIdKey,
		func(ctx context.Context, conn *sqlx.DB, v interface{}) error {
			document := svcCtx.Dao.GetDocumentById(ctx, cacheId)
			*v.(**mtproto.Document) = document
			return nil
		})

	fmt.Println(document2)
}
