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
)

// MessagesAddChatUserCBC6D107
// messages.addChatUser#cbc6d107 chat_id:long user_id:InputUser fwd_limit:int = messages.InvitedUsers;
func (c *ChatsCore) MessagesAddChatUserCBC6D107(in *mtproto.TLMessagesAddChatUserCBC6D107) (*mtproto.Messages_InvitedUsers, error) {
	rV, err := c.addChatUser(in.ChatId, in.UserId, in.FwdLimit)
	if err != nil {
		return nil, err
	}

	return mtproto.MakeTLMessagesInvitedUsers(&mtproto.Messages_InvitedUsers{
		Updates:         rV,
		MissingInvitees: []*mtproto.MissingInvitee{},
	}).To_Messages_InvitedUsers(), nil
}
