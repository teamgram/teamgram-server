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

// UserSetColor
// user.setColor flags:# user_id:long for_profile:flags.1?true color:int background_emoji_id:long = Bool;
func (c *UserCore) UserSetColor(in *user.TLUserSetColor) (*mtproto.Bool, error) {
	var (
		rB bool
	)

	if in.ForProfile {
		rB = c.svcCtx.Dao.UpdateColor(
			c.ctx,
			in.UserId,
			true,
			in.Color,
			in.BackgroundEmojiId)
	} else {
		rB = c.svcCtx.Dao.UpdateColor(
			c.ctx,
			in.UserId,
			false,
			in.Color,
			in.BackgroundEmojiId)
	}

	return mtproto.ToBool(rB), nil
}
