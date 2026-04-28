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

import (
	"github.com/teamgram/teamgram-server/v2/app/service/biz/chat/chat"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

// ChatSearch
// chat.search self_id:long q:string offset:long limit:int = Vector<MutableChat>;
func (c *ChatCore) ChatSearch(in *chat.TLChatSearch) (*chat.VectorMutableChat, error) {
	r := &chat.VectorMutableChat{Datas: []tg.MutableChatClazz{}}
	if len(in.Q) < 3 || in.Limit <= 0 {
		return r, nil
	}
	limit := in.Limit
	if limit > 50 {
		limit = 50
	}

	chats, err := c.repo().Search(c.ctx, in.SelfId, "%"+in.Q+"%", in.Offset, limit)
	if err != nil {
		return nil, err
	}
	for _, item := range chats {
		if item != nil {
			r.Datas = append(r.Datas, item)
		}
	}
	return r, nil
}
