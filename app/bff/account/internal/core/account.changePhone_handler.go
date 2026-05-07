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
	"github.com/teamgram/teamgram-server/v2/app/bff/internal/userprojection"
	codepb "github.com/teamgram/teamgram-server/v2/app/service/biz/code/code"
	userpb "github.com/teamgram/teamgram-server/v2/app/service/biz/user/user"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

// AccountChangePhone
// account.changePhone#70c32edb phone_number:string phone_code_hash:string phone_code:string = User;
func (c *AccountCore) AccountChangePhone(in *tg.TLAccountChangePhone) (*tg.User, error) {
	selfID, err := requireSelfID(c)
	if err != nil {
		return nil, err
	}
	if in == nil {
		return nil, tg.Err406PhoneNumberInvalid
	}
	if in.PhoneCodeHash == "" || in.PhoneCode == "" {
		return nil, tg.ErrPhoneCodeEmpty
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

	existing, err := c.svcCtx.Repo.UserClient.UserGetImmutableUserByPhone(c.ctx, &userpb.TLUserGetImmutableUserByPhone{
		Phone: phone,
	})
	if err != nil {
		if !isUserNotFound(err) {
			return nil, err
		}
	} else if existing != nil {
		return nil, tg.ErrPhoneNumberOccupied
	}

	codeData, err := c.svcCtx.Repo.CodeClient.CodeGetPhoneCode(c.ctx, &codepb.TLCodeGetPhoneCode{
		AuthKeyId:     accountAuthKeyID(c),
		Phone:         phone,
		PhoneCodeHash: in.PhoneCodeHash,
	})
	if err != nil {
		return nil, mapPhoneCodeError(err)
	}
	if codeData == nil || codeData.PhoneCode != in.PhoneCode {
		return nil, tg.ErrPhoneCodeInvalid
	}

	me, err := c.svcCtx.Repo.UserClient.UserGetImmutableUser(c.ctx, &userpb.TLUserGetImmutableUser{
		Id: selfID,
	})
	if err != nil {
		return nil, err
	}
	if me == nil || me.User == nil {
		return nil, tg.ErrUserIdInvalid
	}

	if _, err = c.svcCtx.Repo.UserClient.UserChangePhone(c.ctx, &userpb.TLUserChangePhone{
		UserId: selfID,
		Phone:  phone,
	}); err != nil {
		return nil, err
	}
	if _, err = c.svcCtx.Repo.CodeClient.CodeDeletePhoneCode(c.ctx, &codepb.TLCodeDeletePhoneCode{
		AuthKeyId:     accountAuthKeyID(c),
		Phone:         phone,
		PhoneCodeHash: in.PhoneCodeHash,
	}); err != nil {
		c.Logger.Errorf("account.changePhone - delete phone code failed: auth_key_id: %d, phone: %s, phone_code_hash: %s, err: %v",
			accountAuthKeyID(c), phone, in.PhoneCodeHash, err)
	}

	// TODO(v2 account): user/session update delivery is intentionally not migrated from master; route phone-change updates through userupdates/gateway when the v2 delivery contract is defined.
	users, err := userprojection.ProjectUsers(c.ctx, c.svcCtx.Repo.UserClient, selfID, []int64{selfID}, userprojection.MissingExplicitInput)
	if err != nil {
		return nil, err
	}
	if len(users) == 0 {
		return nil, tg.ErrUserIdInvalid
	}

	return &tg.User{Clazz: users[0]}, nil
}
