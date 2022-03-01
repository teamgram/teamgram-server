/*
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2021-present,  Teamgram Studio (https://teamgram.io).
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package core

import (
	"context"

	"github.com/teamgram/proto/mtproto/rpc/metadata"
	"github.com/teamgram/teamgram-server/app/service/biz/banned/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type BannedCore struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	MD     *metadata.RpcMetadata
}

func New(ctx context.Context, svcCtx *svc.ServiceContext) *BannedCore {
	return &BannedCore{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
		MD:     metadata.RpcMetadataFromIncoming(ctx),
	}
}
