// Copyright 2025 Teamgram Authors
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

// AccountSaveMusic
// account.saveMusic#b26732a9 flags:# unsave:flags.0?true id:InputDocument after_id:flags.1?InputDocument = Bool;
func (c *UserChannelProfilesCore) AccountSaveMusic(in *mtproto.TLAccountSaveMusic) (*mtproto.Bool, error) {
	_, err := c.svcCtx.Dao.UserClient.UserSaveMusic(c.ctx, &userpb.TLUserSaveMusic{
		Unsave:  in.GetUnsave(),
		UserId:  c.MD.UserId,
		Id:      in.GetId().GetId(),
		AfterId: mtproto.MakeFlagsInt64(in.GetAfterId().GetId()),
	})
	if err != nil {
		c.Logger.Errorf("account.saveMusic - error: %v", err)
		return mtproto.BoolFalse, nil
	}

	return mtproto.BoolTrue, nil
}
