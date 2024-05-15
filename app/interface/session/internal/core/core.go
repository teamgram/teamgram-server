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
	"github.com/teamgram/teamgram-server/app/interface/session/internal/sess"
	"github.com/teamgram/teamgram-server/app/interface/session/internal/svc"
	"github.com/teamgram/teamgram-server/app/service/authsession/authsession"

	"github.com/zeromicro/go-zero/core/logx"
)

type SessionCore struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	MD *metadata.RpcMetadata
}

func New(ctx context.Context, svcCtx *svc.ServiceContext) *SessionCore {
	return &SessionCore{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
		MD:     metadata.RpcMetadataFromIncoming(ctx),
	}
}

func (c *SessionCore) getOrFetchMainAuthWrapper(mainAuthId int64) (*sess.MainAuthWrapper, error) {
	mainAuth := c.svcCtx.MainAuthMgr.GetMainAuthWrapper(mainAuthId)
	if mainAuth != nil {
		return mainAuth, nil
	}

	kData, err := c.svcCtx.Dao.AuthsessionClient.AuthsessionGetAuthStateData(c.ctx, &authsession.TLAuthsessionGetAuthStateData{
		AuthKeyId: mainAuthId,
	})
	if err != nil {
		c.Logger.Errorf("session.createSession - error: %v", err)
		return nil, err
	}

	mainAuth = c.svcCtx.MainAuthMgr.AllocMainAuthWrapper(sess.NewMainAuthWrapper(
		mainAuthId,
		kData.UserId,
		int(kData.KeyState),
		kData.Client,
		kData.GetAndroidPushSessionId().GetValue(),
		c.svcCtx.MainAuthMgr))

	return mainAuth, nil
}
