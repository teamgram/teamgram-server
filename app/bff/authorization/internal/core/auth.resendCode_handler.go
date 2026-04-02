// Copyright (c) 2024 The Teamgooo Authors. All rights reserved.
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
	"github.com/teamgram/teamgram-server/v2/app/bff/authorization/internal/logic"
	"github.com/teamgram/teamgram-server/v2/app/bff/authorization/model"

	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

// AuthResendCode
// auth.resendCode#cae47523 flags:# phone_number:string phone_code_hash:string reason:flags.0?string = auth.SentCode;
func (c *AuthorizationCore) AuthResendCode(in *tg.TLAuthResendCode) (*tg.AuthSentCode, error) {
	if in.PhoneCodeHash == "" {
		err := tg.ErrPhoneCodeHashEmpty
		c.Logger.Errorf("auth.resendCode - error: %v", err)
		return nil, err
	}

	_, phoneNumber, err := checkPhoneNumberInvalid(in.PhoneNumber)
	if err != nil {
		c.Logger.Errorf("auth.resendCode - invalid phone_number(%s): %v", in.PhoneNumber, err)
		return nil, tg.Err406PhoneNumberInvalid
	}

	if c.svcCtx == nil || c.svcCtx.Dao == nil || c.svcCtx.AuthLogic == nil || c.MD == nil {
		err = tg.ErrInternalServerError
		c.Logger.Errorf("auth.resendCode - missing service context or metadata: %v", err)
		return nil, err
	}

	actionType := logic.GetActionType(in)
	if err = c.svcCtx.Dao.CheckCanDoAction(c.ctx, c.MD.PermAuthKeyId, phoneNumber, actionType); err != nil {
		c.Logger.Errorf("auth.resendCode - check can do action failed, phone_number(%s): %v", phoneNumber, err)
		return nil, err
	}

	codeData, err := c.svcCtx.AuthLogic.DoAuthReSendCode(
		c.ctx,
		c.MD.PermAuthKeyId,
		phoneNumber,
		in.PhoneCodeHash,
		func(codeData *model.PhoneCodeTransaction) error {
			if c.svcCtx.AuthLogic.VerifyCodeInterface != nil {
				extraData, err := c.svcCtx.AuthLogic.VerifyCodeInterface.SendSmsVerifyCode(
					c.ctx,
					phoneNumber,
					codeData.PhoneCode,
					codeData.PhoneCodeHash)
				if err != nil {
					return err
				}
				codeData.PhoneCodeExtraData = extraData
			}

			codeData.SentCodeType = model.SentCodeTypeSms
			codeData.NextCodeType = model.CodeTypeSms
			codeData.State = model.CodeStateSent
			return nil
		})
	if err != nil {
		c.Logger.Errorf("auth.resendCode - resend failed, phone_number(%s): %v", phoneNumber, err)
		return nil, err
	}

	if c.svcCtx.Plugin != nil {
		c.svcCtx.Plugin.OnAuthAction(c.ctx,
			c.MD.PermAuthKeyId,
			c.MD.ClientMsgId,
			c.MD.ClientAddr,
			in.PhoneNumber,
			actionType,
			"auth.resendCode")
	}

	return codeData.ToAuthSentCode(), nil
}
