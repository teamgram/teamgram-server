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
	"time"

	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

// MessagesGetSearchResultsCalendar
// messages.getSearchResultsCalendar#6aa3f6bd flags:# peer:InputPeer saved_peer_id:flags.2?InputPeer filter:MessagesFilter offset_id:int offset_date:int = messages.SearchResultsCalendar;
func (c *MessagesCore) MessagesGetSearchResultsCalendar(in *tg.TLMessagesGetSearchResultsCalendar) (*tg.MessagesSearchResultsCalendar, error) {
	peer, err := bffPeerFromInput(c, in.Peer)
	if err != nil {
		return nil, err
	}
	startID := historyPlaceholderStartID(in.OffsetId, 0, 0)
	now := int32(time.Now().Unix())

	return tg.MakeTLMessagesSearchResultsCalendar(&tg.TLMessagesSearchResultsCalendar{
		Count:    1,
		MinDate:  now,
		MinMsgId: startID,
		Periods: []tg.SearchResultsCalendarPeriodClazz{
			tg.MakeTLSearchResultsCalendarPeriod(&tg.TLSearchResultsCalendarPeriod{
				Date:     now,
				MinMsgId: startID,
				MaxMsgId: startID,
				Count:    1,
			}),
		},
		Messages: []tg.MessageClazz{
			tg.MakeTLMessage(&tg.TLMessage{
				Id:      startID,
				Out:     true,
				Date:    now,
				Message: "placeholder",
				PeerId:  peer,
			}),
		},
		Chats: []tg.ChatClazz{},
		Users: []tg.UserClazz{},
	}).ToMessagesSearchResultsCalendar(), nil
}
