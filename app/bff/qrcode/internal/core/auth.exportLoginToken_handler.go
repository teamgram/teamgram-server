// Copyright (c) 2024 The Teamgram Authors. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Author: teamgramio (teamgram.io@gmail.com)

package core

import (
	"time"

	"github.com/teamgram/proto/v2/crypto"
	"github.com/teamgram/proto/v2/tg"
	"github.com/teamgram/teamgram-server/v2/app/bff/qrcode/model"
)

const (
	qrCodeTimeout = 60 // salt timeout
)

// AuthExportLoginToken
// auth.exportLoginToken#b7e085fe api_id:int api_hash:string except_ids:Vector<long> = auth.LoginToken;
func (c *QrCodeCore) AuthExportLoginToken(in *tg.TLAuthExportLoginToken) (*tg.AuthLoginToken, error) {
	qrCode := &model.QRCodeTransaction{
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

	rQRLoginToken := tg.MakeAuthLoginToken(&tg.TLAuthLoginToken{
		Expires: int32(qrCode.ExpireAt),
		Token:   qrCode.Token(),
	})

	return rQRLoginToken, nil
}
