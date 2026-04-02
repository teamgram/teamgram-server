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
	"errors"

	"github.com/teamgram/teamgram-server/v2/app/bff/authorization/internal/logic"
	"github.com/teamgram/teamgram-server/v2/app/bff/authorization/model"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

// AuthSignIn
// auth.signIn#8d52a951 flags:# phone_number:string phone_code_hash:string phone_code:flags.0?string email_verification:flags.1?EmailVerification = auth.Authorization;
func (c *AuthorizationCore) AuthSignIn(in *tg.TLAuthSignIn) (*tg.AuthAuthorization, error) {
	var phoneCode string
	if in.PhoneCode != nil {
		phoneCode = *in.PhoneCode
	}

	if phoneCode == "" || in.PhoneCodeHash == "" {
		err := tg.ErrPhoneCodeEmpty
		c.Logger.Errorf("auth.signIn - error: %v", err)
		return nil, err
	}

	_, phoneNumber, err := checkPhoneNumberInvalid(in.PhoneNumber)
	if err != nil {
		c.Logger.Errorf("auth.signIn - invalid phone_number(%s): %v", in.PhoneNumber, err)
		return nil, tg.Err406PhoneNumberInvalid
	}

	if c.svcCtx == nil || c.svcCtx.Dao == nil || c.svcCtx.AuthLogic == nil || c.MD == nil {
		err := tg.ErrInternalServerError
		c.Logger.Errorf("auth.signIn - missing service context or metadata: %v", err)
		return nil, err
	}

	actionType := logic.GetActionType(in)
	if err = c.svcCtx.Dao.CheckCanDoAction(c.ctx, c.MD.PermAuthKeyId, phoneNumber, actionType); err != nil {
		c.Logger.Errorf("auth.signIn - check can do action failed, phone_number(%s): %v", phoneNumber, err)
		return nil, err
	}

	codeData, err := c.svcCtx.AuthLogic.DoAuthSignIn(
		c.ctx,
		c.MD.PermAuthKeyId,
		phoneNumber,
		phoneCode,
		in.PhoneCodeHash,
		func(codeData *model.PhoneCodeTransaction) error {
			if c.svcCtx.AuthLogic.VerifyCodeInterface != nil {
				return c.svcCtx.AuthLogic.VerifyCodeInterface.VerifySmsCode(
					c.ctx,
					codeData.PhoneCodeHash,
					phoneCode,
					codeData.PhoneCodeExtraData)
			}
			if codeData.PhoneCode != phoneCode {
				return errors.New("phone code mismatch")
			}
			return nil
		})
	if err != nil {
		c.Logger.Errorf("auth.signIn - sign in failed, phone_number(%s): %v", phoneNumber, err)
		return nil, err
	}

	if c.svcCtx.Plugin != nil {
		c.svcCtx.Plugin.OnAuthAction(c.ctx,
			c.MD.PermAuthKeyId,
			c.MD.ClientMsgId,
			c.MD.ClientAddr,
			in.PhoneNumber,
			actionType,
			"auth.signIn")
	}

	if !codeData.PhoneNumberRegistered {
		if c.MD.Layer >= 104 {
			return tg.MakeTLAuthAuthorizationSignUpRequired(&tg.TLAuthAuthorizationSignUpRequired{}).ToAuthAuthorization(), nil
		}
		return nil, tg.ErrPhoneNumberUnoccupied
	}

	// TODO: continue implementing the registered-user sign-in path.
	return nil, tg.ErrInternalServerError
}
