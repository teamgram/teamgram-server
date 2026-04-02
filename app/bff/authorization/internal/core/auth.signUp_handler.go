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

	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

// AuthSignUp
// auth.signUp#aac7b717 flags:# no_joined_notifications:flags.0?true phone_number:string phone_code_hash:string first_name:string last_name:string = auth.Authorization;
func (c *AuthorizationCore) AuthSignUp(in *tg.TLAuthSignUp) (*tg.AuthAuthorization, error) {
	_, phoneNumber, err := checkPhoneNumberInvalid(in.PhoneNumber)
	if err != nil {
		c.Logger.Errorf("auth.signUp - invalid phone_number(%s): %v", in.PhoneNumber, err)
		return nil, tg.Err406PhoneNumberInvalid
	}

	if in.PhoneCodeHash == "" {
		c.Logger.Errorf("auth.signUp - phone_code_hash is empty")
		return nil, tg.ErrPhoneCodeHashEmpty
	}

	if in.FirstName == "" {
		c.Logger.Errorf("auth.signUp - first_name is empty")
		return nil, tg.ErrFirstnameInvalid
	}

	if c.svcCtx == nil || c.svcCtx.AuthLogic == nil || c.MD == nil {
		err = tg.ErrInternalServerError
		c.Logger.Errorf("auth.signUp - missing service context or metadata: %v", err)
		return nil, err
	}

	if c.svcCtx.Plugin != nil {
		c.svcCtx.Plugin.OnAuthAction(c.ctx,
			c.MD.PermAuthKeyId,
			c.MD.ClientMsgId,
			c.MD.ClientAddr,
			in.PhoneNumber,
			logic.GetActionType(in),
			"auth.signUp")
	}

	if _, err = c.svcCtx.AuthLogic.DoAuthSignUp(c.ctx, c.MD.PermAuthKeyId, phoneNumber, nil, in.PhoneCodeHash); err != nil {
		c.Logger.Errorf("auth.signUp - sign up failed, phone_number(%s): %v", phoneNumber, err)
		return nil, err
	}

	// TODO: continue implementing user creation and authorization binding.
	return nil, tg.ErrInternalServerError
}
