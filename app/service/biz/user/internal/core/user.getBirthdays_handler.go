// Copyright 2024 Teamgram Authors
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
	"github.com/teamgram/teamgram-server/app/service/biz/user/user"
)

// UserGetBirthdays
// user.getBirthdays user_id:long = Vector<ContactBirthday>;
func (c *UserCore) UserGetBirthdays(in *user.TLUserGetBirthdays) (*user.Vector_ContactBirthday, error) {
	// TODO: not impl
	c.Logger.Errorf("user.getBirthdays blocked, License key from https://teamgram.net required to unlock enterprise features.")

	rV := &user.Vector_ContactBirthday{
		Datas: make([]*mtproto.ContactBirthday, 0),
	}

	return rV, nil
}
