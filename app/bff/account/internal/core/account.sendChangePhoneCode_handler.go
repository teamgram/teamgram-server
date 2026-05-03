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
	codepb "github.com/teamgram/teamgram-server/v2/app/service/biz/code/code"
	userpb "github.com/teamgram/teamgram-server/v2/app/service/biz/user/user"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

// AccountSendChangePhoneCode
// account.sendChangePhoneCode#82574ae5 phone_number:string settings:CodeSettings = auth.SentCode;
func (c *AccountCore) AccountSendChangePhoneCode(in *tg.TLAccountSendChangePhoneCode) (*tg.AuthSentCode, error) {
	if _, err := requireSelfID(c); err != nil {
		return nil, err
	}
	if in == nil {
		return nil, tg.Err406PhoneNumberInvalid
	}
	if err := requireUserClient(c); err != nil {
		return nil, err
	}
	if err := requireCodeClient(c); err != nil {
		return nil, err
	}

	phone, err := normalizeAccountPhone(in.PhoneNumber)
	if err != nil {
		return nil, err
	}

	user, err := c.svcCtx.Repo.UserClient.UserGetImmutableUserByPhone(c.ctx, &userpb.TLUserGetImmutableUserByPhone{
		Phone: phone,
	})
	if err != nil {
		if !isUserNotFound(err) {
			return nil, err
		}
	} else if user != nil {
		return nil, tg.ErrPhoneNumberOccupied
	}

	codeData, err := c.svcCtx.Repo.CodeClient.CodeCreatePhoneCode(c.ctx, &codepb.TLCodeCreatePhoneCode{
		AuthKeyId:             accountAuthKeyID(c),
		SessionId:             accountSessionID(c),
		Phone:                 phone,
		PhoneNumberRegistered: false,
		SentCodeType:          1,
		NextCodeType:          1,
		State:                 1,
	})
	if err != nil {
		return nil, err
	}

	// TODO(v2 account): real SMS/app delivery is intentionally not migrated from master; route through the v2 verification delivery contract when it is defined.
	return makeAccountSentCode(codeData)
}
