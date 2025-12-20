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
	"github.com/teamgram/teamgram-server/app/service/biz/user/user"
)

// UserCheckChannelUsername
// user.checkChannelUsername channel_id:long username:string = UsernameExist;
func (c *UserCore) UserCheckChannelUsername(in *user.TLUserCheckChannelUsername) (*user.UsernameExist, error) {
	var (
		checked = usernameNotExisted
	)

	// TODO(@benqi): check len(username) >= 5
	usernameDO, _ := c.svcCtx.Dao.UsernameDAO.SelectByUsername(c.ctx, in.Username)
	if usernameDO != nil {
		if usernameDO.PeerType == mtproto.PEER_CHANNEL && usernameDO.PeerId == in.ChannelId {
			checked = usernameExistedIsMe
		} else {
			checked = usernameExistedNotMe
		}
	}

	return checked, nil
}
