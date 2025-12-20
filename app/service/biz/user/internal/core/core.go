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
	"github.com/teamgram/teamgram-server/app/service/biz/user/internal/svc"
	"github.com/teamgram/teamgram-server/app/service/biz/user/user"

	"github.com/zeromicro/go-zero/core/logx"
)

var (
	usernameNotExisted   = user.MakeTLUsernameNotExisted(nil).To_UsernameExist()
	usernameExisted      = user.MakeTLUsernameExisted(nil).To_UsernameExist()
	usernameExistedNotMe = user.MakeTLUsernameExistedNotMe(nil).To_UsernameExist()
	usernameExistedIsMe  = user.MakeTLUsernameExistedIsMe(nil).To_UsernameExist()
)

type UserCore struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	MD *metadata.RpcMetadata
}

func New(ctx context.Context, svcCtx *svc.ServiceContext) *UserCore {
	return &UserCore{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
		MD:     metadata.RpcMetadataFromIncoming(ctx),
	}
}
