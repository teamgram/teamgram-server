// Copyright 2022 Teamgram Authors
//  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Author: teamgramio (teamgram.io@gmail.com)
//

package core

import (
	"github.com/teamgram/proto/mtproto"
	userpb "github.com/teamgram/teamgram-server/app/service/biz/user/user"
)

// AccountSetAccountTTL
// account.setAccountTTL#2442485e ttl:AccountDaysTTL = Bool;
func (c *AccountCore) AccountSetAccountTTL(in *mtproto.TLAccountSetAccountTTL) (*mtproto.Bool, error) {
	// TODO(@benqi): Check ttl
	ttl := in.GetTtl().GetDays()
	switch ttl {
	case 30:
	case 90:
	case 182:
	case 365:
	default:
		err := mtproto.ErrTtlDaysInvalid
		c.Logger.Errorf("account.setAccountTTL - error: %v", err)
		return nil, err
	}

	if _, err := c.svcCtx.Dao.UserClient.UserSetAccountDaysTTL(c.ctx, &userpb.TLUserSetAccountDaysTTL{
		UserId: c.MD.UserId,
		Ttl:    ttl,
	}); err != nil {
		c.Logger.Errorf("account.setAccountTTL - error: %v", err)
	}

	return mtproto.BoolTrue, nil
}
