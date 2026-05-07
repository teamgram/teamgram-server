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
	"errors"

	"github.com/teamgram/teamgram-server/v2/app/bff/authorization/internal/repository"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

// AuthSignIn
// auth.signIn#8d52a951 flags:# phone_number:string phone_code_hash:string phone_code:flags.0?string email_verification:flags.1?EmailVerification = auth.Authorization;
func (c *AuthorizationCore) AuthSignIn(in *tg.TLAuthSignIn) (*tg.AuthAuthorization, error) {
	if in.PhoneCode == nil {
		return nil, tg.ErrPhoneCodeEmpty
	}
	_, phone, err := normalizeStartupPhone(in.PhoneNumber)
	if err != nil {
		return nil, err
	}
	if _, err = verifyStartupPhoneCode(phone, in.PhoneCodeHash, *in.PhoneCode); err != nil {
		return nil, err
	}

	user, err := c.svcCtx.Repo.GetUserByPhone(c.ctx, phone)
	if err != nil {
		if errors.Is(err, repository.ErrUserNotFound) {
			return makeSignupRequired(), nil
		}
		return nil, err
	}

	authKeyID := startupAuthKeyID(c)
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
