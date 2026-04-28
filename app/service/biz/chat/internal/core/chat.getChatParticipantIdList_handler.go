// Copyright (c) 2026 The Teamgram Authors. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Author: teamgramio (teamgram.io@gmail.com)

package core

import "github.com/teamgram/teamgram-server/v2/app/service/biz/chat/chat"

// ChatGetChatParticipantIdList
// chat.getChatParticipantIdList chat_id:long = Vector<long>;
func (c *ChatCore) ChatGetChatParticipantIdList(in *chat.TLChatGetChatParticipantIdList) (*chat.VectorLong, error) {
	ids, err := c.repo().GetChatParticipantIDList(c.ctx, in.ChatId)
	if err != nil {
		return nil, err
	}
	return &chat.VectorLong{Datas: ids}, nil
}
