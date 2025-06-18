// Copyright 2024 Teamgram Authors
//  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Author: teamgramio (teamgram.io@gmail.com)
//

package core

import (
	"context"

	"github.com/teamgram/proto/mtproto"
	"github.com/teamgram/teamgram-server/v2/app/interface/session/internal/sess"
	"github.com/teamgram/teamgram-server/v2/app/interface/session/internal/svc"
	"github.com/teamgram/teamgram-server/v2/app/service/authsession/authsession"

	"github.com/zeromicro/go-zero/core/logx"
	//"github.com/teamgram/proto/mtproto/rpc/metadata"
)

type SessionCore struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	// MD *metadata.RpcMetadata
}

func New(ctx context.Context, svcCtx *svc.ServiceContext) *SessionCore {
	return &SessionCore{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
		// MD:     metadata.RpcMetadataFromIncoming(ctx),
	}
}

func (c *SessionCore) getOrFetchMainAuthWrapper(mainAuthId int64) (*sess.MainAuthWrapper, error) {
	mainAuth := c.svcCtx.MainAuthMgr.GetMainAuthWrapper(mainAuthId)
	if mainAuth != nil {
		return mainAuth, nil
	}

	var (
		kData *authsession.TLAuthKeyStateData
		// err   error
	)

	if mainAuthId == 0 {
		kData = &authsession.TLAuthKeyStateData{
			AuthKeyId:            0,
			KeyState:             mtproto.AuthStateNew,
			UserId:               0,
			AccessHash:           0,
			Client:               nil,
			AndroidPushSessionId: nil,
		}
	} else {
		kData2, err := c.svcCtx.Dao.AuthsessionClient.AuthsessionGetAuthStateData(c.ctx, &authsession.TLAuthsessionGetAuthStateData{
			AuthKeyId: mainAuthId,
		})
		if err != nil {
			c.Logger.Errorf("getOrFetchMainAuthWrapper - error: %v", err)
			return nil, err
		}
		kData, _ = kData2.ToAuthKeyStateData()
	}

	mainAuth = c.svcCtx.MainAuthMgr.AllocMainAuthWrapper(
		mainAuthId,
		func(authKeyId int64) *sess.MainAuthWrapper {
			androidPushSessionId := int64(0)
			if kData.AndroidPushSessionId != nil {
				androidPushSessionId = *kData.AndroidPushSessionId
			}
			return sess.NewMainAuthWrapper(
				mainAuthId,
				kData.UserId,
				int(kData.KeyState),
				kData.Client,
				androidPushSessionId,
				c.svcCtx.MainAuthMgr)
		})

	return mainAuth, nil
}
