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
	"github.com/teamgram/teamgram-server/v2/app/service/biz/code/code"
)

// CodeGetPhoneCode
// code.getPhoneCode auth_key_id:long phone:string phone_code_hash:string = PhoneCodeTransaction;
func (c *CodeCore) CodeGetPhoneCode(in *code.TLCodeGetPhoneCode) (*code.PhoneCodeTransaction, error) {
	codeData, err := c.repo.GetCachePhoneCode(c.ctx, in.AuthKeyId, in.Phone)
	if err != nil {
		c.Logger.Errorf("code.getPhoneCode - get cache failed: auth_key_id: %d, phone: %s, err: %v",
			in.AuthKeyId, in.Phone, err)
		return nil, code.ErrPhoneCodeExpired
	}
	if codeData == nil {
		c.Logger.Errorf("code.getPhoneCode - not found: auth_key_id: %d, phone: %s",
			in.AuthKeyId, in.Phone)
		return nil, code.ErrPhoneCodeExpired
	}
	if codeData.PhoneCodeHash != in.PhoneCodeHash {
		c.Logger.Errorf("code.getPhoneCode - hash mismatch: auth_key_id: %d, phone: %s",
			in.AuthKeyId, in.Phone)
		return nil, code.ErrPhoneCodeInvalid
	}

	return codeData, nil
}
