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

// MessagesGetSearchResultsPositions
// messages.getSearchResultsPositions#9c7f2f10 flags:# peer:InputPeer saved_peer_id:flags.2?InputPeer filter:MessagesFilter offset_id:int limit:int = messages.SearchResultsPositions;
func (c *MessagesCore) MessagesGetSearchResultsPositions(in *tg.TLMessagesGetSearchResultsPositions) (*tg.MessagesSearchResultsPositions, error) {
	startID := historyPlaceholderStartID(in.OffsetId, 0, 0)
	return tg.MakeTLMessagesSearchResultsPositions(&tg.TLMessagesSearchResultsPositions{
		Count: 1,
		Positions: []tg.SearchResultsPositionClazz{
			tg.MakeTLSearchResultPosition(&tg.TLSearchResultPosition{
				MsgId:  startID,
				Date:   int32(time.Now().Unix()),
				Offset: 0,
			}),
		},
	}).ToMessagesSearchResultsPositions(), nil
}
