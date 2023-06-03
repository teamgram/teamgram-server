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
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/teamgram/marmota/pkg/stores/sqlx"
	"github.com/teamgram/proto/mtproto"
	"github.com/teamgram/teamgram-server/app/service/biz/user/internal/dal/dataobject"
	"github.com/teamgram/teamgram-server/app/service/biz/user/user"
)

const (
	minTestUserId = 10000000 - 1
	maxTestUserId = 20000000
)

// UserCreateNewTestUser
// user.createNewTestUser secret_key_id:long min_id:long max_id:long = ImmutableUser;
func (c *UserCore) UserCreateNewTestUser(in *user.TLUserCreateNewTestUser) (*mtproto.ImmutableUser, error) {
	var (
		userDO *dataobject.UsersDO
		now    = time.Now().Unix()
		minId  = in.GetMinId()
		maxId  = in.GetMaxId()
	)

	if minId == 0 {
		minId = minTestUserId
	}
	if maxId == 0 {
		maxId = maxTestUserId
	}

	var (
		id       = minId + 1
		hasRetry = 0
	)

retry:
	userDO = &dataobject.UsersDO{
		Id:             id,
		UserType:       user.UserTypeTest,
		AccessHash:     rand.Int63(),
		Phone:          "-" + strconv.FormatInt(id, 10),
		SecretKeyId:    in.GetSecretKeyId(),
		FirstName:      "t" + strconv.FormatInt(id, 10),
		LastName:       "",
		CountryCode:    "CN",
		AccountDaysTtl: 180,
	}

	_, _, err2 := c.svcCtx.Dao.UsersDAO.InsertTestUser(
		c.ctx,
		userDO)
	if err2 != nil {
		if sqlx.IsDuplicate(err2) {
			do2, err := c.svcCtx.Dao.UsersDAO.SelectNextTestUserId(c.ctx, maxId)
			if err != nil {
				return nil, err
			}
			if do2 == nil {
				return nil, fmt.Errorf("not found next test user id")
			}
			if do2.Id+1 == maxId {
				return nil, fmt.Errorf("not found next test user id")
			}
			if hasRetry < 10 {
				id = do2.Id + 1
				hasRetry++
				goto retry
			}
		}
		return nil, err2
	}

	return mtproto.MakeTLImmutableUser(&mtproto.ImmutableUser{
		User:             c.svcCtx.Dao.MakeUserDataByDO(userDO),
		LastSeenAt:       now,
		Contacts:         nil,
		KeysPrivacyRules: nil,
	}).To_ImmutableUser(), nil
}
