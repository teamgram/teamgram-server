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
	"github.com/teamgram/teamgram-server/app/service/biz/chat/chat"
)

// ChatSearch
// chat.search self_id:long q:string offset:long limit:int = Vector<UserChatIdList>;
func (c *ChatCore) ChatSearch(in *chat.TLChatSearch) (*chat.Vector_MutableChat, error) {
	var (
		chatList = &chat.Vector_MutableChat{
			Datas: []*mtproto.MutableChat{},
		}
	)

	// Check query string and limit
	if len(in.Q) < 3 || in.Limit <= 0 {
		return chatList, nil
	}

	if in.Limit > 50 {
		in.Limit = 50
	}

	// 构造模糊查询字符串
	q := "%" + in.Q + "%"

	c.svcCtx.Dao.ChatsDAO.SearchByQueryStringWithCB(
		c.ctx,
		q,
		in.Limit,
		func(sz, i int, v int64) {
			chat, err := c.svcCtx.Dao.GetExcludeParticipantsMutableChat(c.ctx, v)
			if err != nil {
				c.Logger.Errorf("chat.search - error: %v", err)
			} else if chat != nil {
				chatList.Datas = append(chatList.Datas, chat)
			}
		})

	return chatList, nil
}
