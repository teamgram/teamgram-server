// Copyright 2022 Teamgram Authors
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
	"fmt"
	"time"

	"github.com/teamgram/proto/mtproto"
	"github.com/teamgram/proto/mtproto/crypto"
	"github.com/teamgram/teamgram-server/app/bff/qrcode/internal/model"
	userpb "github.com/teamgram/teamgram-server/app/service/biz/user/user"

	"google.golang.org/grpc/status"
)

const (
	qrCodeTimeout = 60 // salt timeout
)

// AuthExportLoginToken
// auth.exportLoginToken#b7e085fe api_id:int api_hash:string except_ids:Vector<long> = auth.LoginToken;
func (c *QrCodeCore) AuthExportLoginToken(in *mtproto.TLAuthExportLoginToken) (*mtproto.Auth_LoginToken, error) {
	// TODO:
	// 1. check api_id, api_hash
	// 2. check except_ids

	qrCode, err := c.svcCtx.Dao.GetCacheQRLoginCode(c.ctx, c.MD.PermAuthKeyId)
	if err != nil {
		c.Logger.Errorf("getQRCode - error: %v", err)
		return nil, err
	} else if qrCode == nil {
		qrCode = &model.QRCodeTransaction{
			PermAuthKeyId: c.MD.PermAuthKeyId,
			AuthKeyId:     c.MD.AuthId,
			SessionId:     c.MD.SessionId,
			ServerId:      c.MD.ServerId,
			ApiId:         in.ApiId,
			ApiHash:       in.ApiHash,
			CodeHash:      crypto.GenerateStringNonce(16),
			ExpireAt:      time.Now().Unix() + qrCodeTimeout,
			UserId:        0,
			State:         model.QRCodeStateNew,
		}
		c.Logger.Infof("putQRCode - %#v", qrCode)
		if err = c.svcCtx.Dao.PutCacheQRLoginCode(c.ctx, c.MD.PermAuthKeyId, qrCode, qrCodeTimeout+2); err != nil {
			c.Logger.Errorf("putQRCode - error: %v", err)
			return nil, err
		}
	} else {
		c.Logger.Infof("putQRCode - %#v", qrCode)
	}

	var (
		rQRLoginToken *mtproto.Auth_LoginToken
	)

	switch qrCode.State {
	case model.QRCodeStateAccepted, model.QRCodeStateSuccess:
		//// Check SESSION_PASSWORD_NEEDED
		//if sessionPasswordNeeded, _ := c.svcCtx.Dao.TwofaClient.TwofaCheckSessionPasswordNeeded(c.ctx, &twofa.TLTwofaCheckSessionPasswordNeeded{
		//	UserId: qrCode.UserId,
		//}); mtproto.FromBool(sessionPasswordNeeded) {
		//	err = mtproto.ErrSessionPasswordNeeded
		//	c.Logger.Infof("auth.exportLoginToken - registered, next step auth.checkPassword: %v", err)
		//	return nil, err
		//}

		user, err := c.svcCtx.Dao.UserClient.UserGetImmutableUser(c.ctx, &userpb.TLUserGetImmutableUser{
			Id: qrCode.UserId,
		})
		if err != nil {
			c.Logger.Errorf("auth.exportLoginToken - error: %v", err)
			return nil, err
		}
		if c.svcCtx.Plugin != nil {
			if c.svcCtx.Plugin.CheckSessionPasswordNeeded(c.ctx, user.User.Id) {
				// hack
				// err = mtproto.ErrSessionPasswordNeeded
				err = status.Error(mtproto.ErrUnauthorized, fmt.Sprintf("SESSION_PASSWORD_NEEDED_%d", user.Id()))
				c.Logger.Infof("auth.exportLoginToken - registered, next step auth.checkPassword: %v", err)
				return nil, err
			}
		}

		rQRLoginToken = mtproto.MakeTLAuthLoginTokenSuccess(&mtproto.Auth_LoginToken{
			Authorization: mtproto.MakeTLAuthAuthorization(&mtproto.Auth_Authorization{
				SetupPasswordRequired: false,
				OtherwiseReloginDays:  nil,
				TmpSessions:           nil,
				FutureAuthToken:       nil,
				User:                  user.ToSelfUser(),
			}).To_Auth_Authorization(),
		}).To_Auth_LoginToken()
	default:
		rQRLoginToken = mtproto.MakeTLAuthLoginToken(&mtproto.Auth_LoginToken{
			Expires: int32(qrCode.ExpireAt),
			Token:   qrCode.Token(),
		}).To_Auth_LoginToken()
	}

	return rQRLoginToken, nil
}
