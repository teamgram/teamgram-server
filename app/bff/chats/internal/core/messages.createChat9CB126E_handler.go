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

// MessagesCreateChat9CB126E
// messages.createChat#9cb126e users:Vector<InputUser> title:string = Updates;
func (c *ChatsCore) MessagesCreateChat9CB126E(in *mtproto.TLMessagesCreateChat9CB126E) (*mtproto.Updates, error) {
	return c.createChat(in.Users, in.Title, 0)
}
