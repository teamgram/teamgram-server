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
	"github.com/teamgram/teamgram-server/app/service/biz/dialog/dialog"
)

// DialogUpdateUnreadCount
// dialog.updateUnreadCount flags:# user_id:long peer_type:int peer_id:long unread_count:flags.0?int unread_mentions_count:flags.1?int unread_reactions_count:flags.2?int = Bool;
func (c *DialogCore) DialogUpdateUnreadCount(in *dialog.TLDialogUpdateUnreadCount) (*mtproto.Bool, error) {
	var (
		unreadCount          = in.GetUnreadCount().GetValue()
		unreadMentionsCount  = in.GetUnreadMentionsCount().GetValue()
		unreadReactionsCount = in.GetUnreadReactionsCount().GetValue()
	)

	c.svcCtx.Dao.DialogsDAO.UpdateUnreadCount(
		c.ctx,
		unreadCount,
		unreadMentionsCount,
		unreadReactionsCount,
		in.UserId,
		in.PeerType,
		in.PeerId)

	return mtproto.BoolTrue, nil
}
