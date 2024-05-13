/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright 2022 Teamgram Authors.
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package files_helper

import (
	"github.com/teamgram/teamgram-server/app/bff/files/internal/config"
	"github.com/teamgram/teamgram-server/app/bff/files/internal/server/grpc/service"
	"github.com/teamgram/teamgram-server/app/bff/files/internal/svc"
	"github.com/teamgram/teamgram-server/app/bff/files/plugin"
	"github.com/zeromicro/go-zero/zrpc"
)

type (
	Config = config.Config
)

func init() {
	zrpc.DontLogContentForMethod("/mtproto.RPCFiles/upload_saveFilePart")
	zrpc.DontLogContentForMethod("/mtproto.RPCFiles/upload_saveBigFilePart")

	zrpc.DontLogClientContentForMethod("/dfs.RPCDfs/dfs_writeFilePartData")
	zrpc.DontLogClientContentForMethod("/dfs.RPCDfs/dfs_downloadFile")
}

func New(c Config, plugin1 plugin.ThumbFilesPlugin, plugin2 plugin.StreamFilesPlugin) *service.Service {
	return service.New(svc.NewServiceContext(c, plugin1, plugin2))
}
