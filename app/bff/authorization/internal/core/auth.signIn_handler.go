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

	if _, _, err := checkPhoneNumberInvalid(in.PhoneNumber); err != nil {
		c.Logger.Errorf("auth.signIn - invalid phone_number(%s): %v", in.PhoneNumber, err)
		return nil, tg.Err406PhoneNumberInvalid
	}

	// TODO: continue implementing the full sign-in path.
	return nil, tg.ErrInternalServerError
}
