// Copyright (c) 2026 The Teamgram Authors. All rights reserved.
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
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

// AuthSignUp
// auth.signUp#aac7b717 flags:# no_joined_notifications:flags.0?true phone_number:string phone_code_hash:string first_name:string last_name:string = auth.Authorization;
func (c *AuthorizationCore) AuthSignUp(in *tg.TLAuthSignUp) (*tg.AuthAuthorization, error) {
	_, phone, err := normalizeStartupPhone(in.PhoneNumber)
	if err != nil {
		return nil, err
	}
	if in.FirstName == "" {
		return nil, tg.ErrFirstnameInvalid
	}
	tx, err := lookupStartupPhoneCode(in.PhoneCodeHash)
	if err != nil {
		return nil, err
	}
	if tx.Phone != phone {
		return nil, tg.ErrPhoneCodeExpired
	}

	authKeyID := startupAuthKeyID(c)
	user, err := c.svcCtx.Repo.CreateUser(c.ctx, startupSecretKeyID(authKeyID), phone, tx.CountryCode, in.FirstName, in.LastName)
	if err != nil {
		return nil, err
	}
	userID := immutableUserID(user)
	if authKeyID != 0 && userID != 0 {
		if err = c.svcCtx.Repo.BindAuthKeyUser(c.ctx, authKeyID, userID); err != nil {
			return nil, err
		}
	}

	selfUser, err := c.svcCtx.Repo.ProjectSelfUser(c.ctx, userID)
	if err != nil {
		return nil, err
	}
	deleteStartupPhoneCode(in.PhoneCodeHash)
	return makeAuthAuthorization(selfUser), nil
}
