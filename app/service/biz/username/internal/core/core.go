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

	"github.com/zeromicro/go-zero/core/logx"

	"github.com/teamgram/proto/mtproto/rpc/metadata"
	"github.com/teamgram/teamgram-server/app/service/biz/username/internal/svc"
	"github.com/teamgram/teamgram-server/app/service/biz/username/username"
)

var (
	usernameNotExisted   = username.MakeTLUsernameNotExisted(nil).To_UsernameExist()
	usernameExisted      = username.MakeTLUsernameExisted(nil).To_UsernameExist()
	usernameExistedNotMe = username.MakeTLUsernameExistedNotMe(nil).To_UsernameExist()
	usernameExistedIsMe  = username.MakeTLUsernameExistedIsMe(nil).To_UsernameExist()
)

type UsernameCore struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	MD *metadata.RpcMetadata
}

func New(ctx context.Context, svcCtx *svc.ServiceContext) *UsernameCore {
	return &UsernameCore{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
		MD:     metadata.RpcMetadataFromIncoming(ctx),
	}
}
