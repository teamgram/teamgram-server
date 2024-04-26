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

package dao

import (
	"context"
	"time"

	"github.com/teamgram/marmota/pkg/random2"
	"github.com/teamgram/proto/mtproto"
	"github.com/teamgram/proto/mtproto/crypto"
	"github.com/teamgram/teamgram-server/app/bff/authorization/internal/model"

	"github.com/zeromicro/go-zero/core/logx"
)

// CreatePhoneCode
// for sendCode
func (d *Dao) CreatePhoneCode(ctx context.Context,
	authKeyId int64,
	sessionId int64,
	phoneNumber string,
	phoneNumberRegistered bool,
	sendCodeType, nextCodeType, state int) (codeData *model.PhoneCodeTransaction, err error) {
	_ = state
	newCodeData := func() *model.PhoneCodeTransaction {
		return &model.PhoneCodeTransaction{
			AuthKeyId:             authKeyId,
			PhoneNumber:           phoneNumber,
			SessionId:             sessionId,
			PhoneNumberRegistered: phoneNumberRegistered,
			PhoneCode:             random2.RandomNumeric(5),
			PhoneCodeHash:         crypto.GenerateStringNonce(16),
			PhoneCodeExpired:      int32(time.Now().Unix() + 3*60),
			SentCodeType:          sendCodeType,
			FlashCallPattern:      "*",
			NextCodeType:          nextCodeType,
			State:                 model.CodeStateSend, // model.CodeStateSent
		}
	}

	if codeData, err = d.GetCachePhoneCode(ctx, authKeyId, phoneNumber); err != nil {
		logx.WithContext(ctx).Errorf("getCachePhoneCode - error: %v", err)
		err = mtproto.ErrInternalServerError
		return
	}
	if codeData == nil {
		codeData = newCodeData()
	} else if sessionId != codeData.SessionId {
		codeData.State = model.CodeStateSend
		codeData.SessionId = sessionId
	}

	return
}

func (d *Dao) GetPhoneCode(ctx context.Context,
	authKeyId int64,
	phoneNumber, phoneCodeHash string) (codeData *model.PhoneCodeTransaction, err error) {

	if codeData, err = d.GetCachePhoneCode(ctx, authKeyId, phoneNumber); err != nil {
		logx.WithContext(ctx).Errorf("getPhoneCode by {authKeyId: %d, phoneNumber: %s} error - %v", authKeyId, phoneNumber, err)
		err = mtproto.ErrPhoneCodeExpired
		return
	} else if codeData == nil {
		logx.WithContext(ctx).Errorf("getPhoneCode by {authKeyId: %d, phoneNumber: %s} error - %v", authKeyId, phoneNumber, err)
		err = mtproto.ErrPhoneCodeExpired
		return
	} else if codeData.PhoneCodeHash != phoneCodeHash {
		logx.WithContext(ctx).Errorf("getPhoneCode by {authKeyId: %d, phoneNumber: %s, phoneCodeHash: %s} error - invalid phoneCodeHash",
			authKeyId,
			phoneNumber,
			phoneCodeHash)
		err = mtproto.ErrPhoneCodeInvalid
	}
	return
}

func (d *Dao) DeletePhoneCode(ctx context.Context, authKeyId int64, phoneNumber, phoneCodeHash string) error {
	return d.DeleteCachePhoneCode(ctx, authKeyId, phoneNumber)
}

func (d *Dao) UpdatePhoneCodeData(ctx context.Context,
	authKeyId int64,
	phoneNumber, phoneCodeHash string,
	codeData *model.PhoneCodeTransaction) error {
	// TODO(@benqi): check state??
	return d.PutCachePhoneCode(ctx, authKeyId, phoneNumber, codeData)
}

func (d *Dao) CheckCanDoAction(ctx context.Context,
	authKeyId int64,
	phoneNumber string,
	actionType int) error {
	// TODO(@benqi): check can do action

	_ = authKeyId
	_ = phoneNumber
	_ = actionType
	return nil
}

//func (d *Dao) LogAuthAction(ctx context.Context,
//	authKeyId, msgId int64,
//	clientIp string,
//	phoneNumber string,
//	actionType int, log string) {
//	_ = phoneNumber
//	do := &dataobject.AuthOpLogsDO{
//		AuthKeyId: authKeyId,
//		Ip:        clientIp,
//		OpType:    int32(actionType),
//		LogText:   log,
//	}
//	d.AuthOpLogsDAO.Insert(ctx, do)
//}
