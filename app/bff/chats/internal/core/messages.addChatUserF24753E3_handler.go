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

// MessagesAddChatUserF24753E3
// messages.addChatUser#f24753e3 chat_id:long user_id:InputUser fwd_limit:int = Updates;
func (c *ChatsCore) MessagesAddChatUserF24753E3(in *mtproto.TLMessagesAddChatUserF24753E3) (*mtproto.Updates, error) {
	return c.addChatUser(in.ChatId, in.UserId, in.FwdLimit)
}
