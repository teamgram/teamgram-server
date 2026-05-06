// Copyright (c) 2026-present, The Teamgram Authors (https://teamgram.net).
//  All rights reserved.
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
	"github.com/teamgram/teamgram-server/v2/app/messenger/msg/internal/repository"
	"github.com/teamgram/teamgram-server/v2/app/messenger/msg/msg"
	"github.com/teamgram/teamgram-server/v2/pkg/pagination"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

// MsgGetHistory
// msg.getHistory user_id:long auth_key_id:long peer_type:int peer_id:long offset_id:int offset_date:int add_offset:int limit:int max_id:int min_id:int hash:long = messages.Messages;
func (c *MsgCore) MsgGetHistory(in *msg.TLMsgGetHistory) (*tg.MessagesMessages, error) {
	if in == nil {
		return nil, tg.ErrInputRequestInvalid
	}

	history, err := c.svcCtx.Repo.ListHistoryMessages(c.ctx, repository.ListHistoryMessagesInput{
		UserID:    in.UserId,
		PeerType:  in.PeerType,
		PeerID:    in.PeerId,
		OffsetID:  in.OffsetId,
		AddOffset: in.AddOffset,
		MaxID:     in.MaxId,
		MinID:     in.MinId,
		Limit:     in.Limit,
	})
	if err != nil {
		return nil, err
	}
	if in.Hash != 0 && pagination.HashInt64IDs(historyMessageIDs(history)) == in.Hash {
		return tg.MakeTLMessagesMessagesNotModified(&tg.TLMessagesMessagesNotModified{
			Count: int32(len(history)),
		}).ToMessagesMessages(), nil
	}

	messages := make([]tg.MessageClazz, 0, len(history))
	for _, item := range history {
		if item.MessageKind != repository.MessageKindText {
			continue
		}
		messages = append(messages, tg.MakeTLMessage(&tg.TLMessage{
			Out:     item.Outgoing,
			Id:      int32(item.PeerSeq),
			FromId:  tg.MakePeerUser(item.FromUserID),
			PeerId:  tg.MakePeerUser(item.PeerID),
			Date:    item.MessageDate,
			Message: item.MessageText,
		}))
	}

	return tg.MakeTLMessagesMessages(&tg.TLMessagesMessages{
		Messages: messages,
		Chats:    []tg.ChatClazz{},
		Users:    []tg.UserClazz{},
	}).ToMessagesMessages(), nil
}

func historyMessageIDs(history []repository.HistoryMessage) []int64 {
	ids := make([]int64, 0, len(history))
	for _, item := range history {
		ids = append(ids, item.PeerSeq)
	}
	return ids
}
