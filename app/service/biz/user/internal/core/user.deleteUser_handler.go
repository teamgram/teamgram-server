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
	"github.com/teamgram/teamgram-server/app/service/biz/user/user"
)

// UserDeleteUser
// user.deleteUser user_id:long reason:string = Bool;
func (c *UserCore) UserDeleteUser(in *user.TLUserDeleteUser) (*mtproto.Bool, error) {
	rB := c.svcCtx.Dao.DeleteUser(
		c.ctx,
		in.GetUserId(),
		in.GetReason())

	return mtproto.ToBool(rB), nil
}
