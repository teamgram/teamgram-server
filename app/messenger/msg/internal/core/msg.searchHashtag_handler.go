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
	"context"
	"strings"

	"github.com/teamgram/teamgram-server/v2/app/messenger/msg/internal/repository"
	"github.com/teamgram/teamgram-server/v2/app/messenger/msg/msg"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

// MsgSearchHashtag
// msg.searchHashtag user_id:long auth_key_id:long peer_type:int peer_id:long hash_tag:string offset_id:int limit:int = messages.Messages;
func (c *MsgCore) MsgSearchHashtag(in *msg.TLMsgSearchHashtag) (*tg.MessagesMessages, error) {
	if in == nil {
		return nil, tg.ErrInputRequestInvalid
	}
	tag := strings.TrimPrefix(strings.TrimSpace(in.HashTag), "#")
	if tag == "" {
		return emptyMsgMessages(), nil
	}

	searchRepo, ok := c.svcCtx.Repo.(interface {
		SearchHashTagMessages(context.Context, repository.SearchHashTagMessagesInput) ([]repository.HistoryMessage, error)
	})
	if !ok {
		return nil, msg.ErrMsgStorage
	}
	history, err := searchRepo.SearchHashTagMessages(c.ctx, repository.SearchHashTagMessagesInput{
		UserID:   in.UserId,
		PeerType: in.PeerType,
		PeerID:   in.PeerId,
		HashTag:  tag,
		OffsetID: in.OffsetId,
		Limit:    in.Limit,
	})
	if err != nil {
		return nil, err
	}
	return messagesFromHistory(history)
}
