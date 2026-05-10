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
	"fmt"
	"math"

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

	bounds, err := c.svcCtx.Repo.ResolveHistoryCursorIDs(c.ctx, in.UserId, in.PeerType, in.PeerId, in.OffsetId, in.MaxId, in.MinId)
	if err != nil {
		return nil, err
	}
	history, err := c.svcCtx.Repo.ListHistoryMessages(c.ctx, repository.ListHistoryMessagesInput{
		UserID:               in.UserId,
		PeerType:             in.PeerType,
		PeerID:               in.PeerId,
		AddOffset:            in.AddOffset,
		Limit:                in.Limit,
		CursorsResolved:      true,
		ResolvedCursorBounds: bounds,
	})
	if err != nil {
		return nil, err
	}
	if in.Hash != 0 && pagination.HashInt64IDs(historyMessageIDs(history)) == in.Hash {
		return tg.MakeTLMessagesMessagesNotModified(&tg.TLMessagesMessagesNotModified{
			Count: int32(len(history)),
		}).ToMessagesMessages(), nil
	}

	return messagesFromHistory(history)
}

func messagesFromHistory(history []repository.HistoryMessage) (*tg.MessagesMessages, error) {
	messages := make([]tg.MessageClazz, 0, len(history))
	for _, item := range history {
		message, err := messageFromHistoryItem(item)
		if err != nil {
			return nil, err
		}
		if message != nil {
			messages = append(messages, message)
		}
	}
	return tg.MakeTLMessagesMessages(&tg.TLMessagesMessages{
		Messages: messages,
		Chats:    []tg.ChatClazz{},
		Users:    []tg.UserClazz{},
	}).ToMessagesMessages(), nil
}

func messageFromHistoryItem(item repository.HistoryMessage) (tg.MessageClazz, error) {
	messageID, err := historyIDInt32(item.UserMessageID, "history message id")
	if err != nil {
		return nil, err
	}
	date, err := msgDateInt32FromUnixSeconds(item.MessageDate, "history message date")
	if err != nil {
		return nil, err
	}
	if len(item.ViewPayload) > 0 {
		return userMessageBoxTLMessage(&repository.UserMessageBox{
			UserMessageID:      item.UserMessageID,
			CanonicalMessageID: item.CanonicalMessageID,
			PeerType:           item.PeerType,
			PeerID:             item.PeerID,
			PeerSeq:            item.PeerSeq,
			FromUserID:         item.FromUserID,
			Outgoing:           item.Outgoing,
			MessageText:        item.MessageText,
			MessageDate:        item.MessageDate,
			ViewPayload:        item.ViewPayload,
		}, messageID, date)
	}
	if item.MessageKind != repository.MessageKindText {
		return nil, nil
	}
	replyTo, err := historyReplyHeader(item.ReplyToUserMessageID)
	if err != nil {
		return nil, err
	}
	return tg.MakeTLMessage(&tg.TLMessage{
		Out:     item.Outgoing,
		Id:      messageID,
		FromId:  userMessageFromPeer(item.Outgoing, item.PeerType, item.FromUserID),
		PeerId:  tg.MakePeerUser(item.PeerID),
		ReplyTo: replyTo,
		Date:    date,
		Message: item.MessageText,
	}), nil
}

func historyReplyHeader(userMessageID int64) (tg.MessageReplyHeaderClazz, error) {
	if userMessageID <= 0 {
		return nil, nil
	}
	replyToMsgID, err := historyIDInt32(userMessageID, "history reply message id")
	if err != nil {
		return nil, err
	}
	return tg.MakeTLMessageReplyHeader(&tg.TLMessageReplyHeader{ReplyToMsgId: &replyToMsgID}), nil
}

func historyIDInt32(v int64, field string) (int32, error) {
	if v < math.MinInt32 || v > math.MaxInt32 {
		return 0, fmt.Errorf("%w: %s out of int32 range", msg.ErrMsgStorage, field)
	}
	return int32(v), nil
}

func emptyMsgMessages() *tg.MessagesMessages {
	messages, _ := messagesFromHistory(nil)
	return messages
}

func historyMessageIDs(history []repository.HistoryMessage) []int64 {
	ids := make([]int64, 0, len(history))
	for _, item := range history {
		ids = append(ids, item.UserMessageID)
	}
	return ids
}
