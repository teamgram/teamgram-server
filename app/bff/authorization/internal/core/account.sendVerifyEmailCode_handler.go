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
	"net/mail"

	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

// AccountSendVerifyEmailCode
// account.sendVerifyEmailCode#98e037bb purpose:EmailVerifyPurpose email:string = account.SentEmailCode;
func (c *AuthorizationCore) AccountSendVerifyEmailCode(in *tg.TLAccountSendVerifyEmailCode) (*tg.AccountSentEmailCode, error) {
	if _, err := mail.ParseAddress(in.Email); err != nil {
		return nil, tg.ErrEmailInvalid
	}

	return tg.MakeTLAccountSentEmailCode(&tg.TLAccountSentEmailCode{
		EmailPattern: "t***@example.com",
		Length:       6,
	}).ToAccountSentEmailCode(), nil
}
