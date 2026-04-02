// Copyright (c) 2024 The Teamgooo Authors. All rights reserved.
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
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

// MessagesGetMessagesViews
// messages.getMessagesViews#5784d3e1 peer:InputPeer id:Vector<int> increment:Bool = messages.MessageViews;
func (c *MessagesCore) MessagesGetMessagesViews(in *tg.TLMessagesGetMessagesViews) (*tg.MessagesMessageViews, error) {
	views := make([]tg.MessageViewsClazz, 0, len(in.Id))
	for _, id := range in.Id {
		v := id
		views = append(views, tg.MakeTLMessageViews(&tg.TLMessageViews{
			Views: &v,
		}))
	}

	return tg.MakeTLMessagesMessageViews(&tg.TLMessagesMessageViews{
		Views: views,
		Chats: []tg.ChatClazz{},
		Users: []tg.UserClazz{},
	}).ToMessagesMessageViews(), nil
}
