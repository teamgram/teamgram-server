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
	"github.com/teamgram/teamgram-server/v2/app/messenger/msg/msg/msg"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

// MessagesReadMessageContents
// messages.readMessageContents#36a73f77 id:Vector<int> = messages.AffectedMessages;
func (c *MessagesCore) MessagesReadMessageContents(in *tg.TLMessagesReadMessageContents) (*tg.MessagesAffectedMessages, error) {
	var userId int64
	if c.MD != nil {
		userId = c.MD.UserId
	}

	// MessagesReadMessageContents marks messages as read for content (media),
	// routed to the first message's peer as a default.
	peerId := int64(0)
	if len(in.Id) > 0 {
		peerId = int64(in.Id[0])
	}

	// When MsgClient is wired, delegate to msg service.
	if c.svcCtx != nil && c.svcCtx.MsgClient != nil {
		var authKeyId int64
		if c.MD != nil {
			authKeyId = c.MD.AuthId
		}

		contentMsgs := make([]msg.ContentMessageClazz, 0, len(in.Id))
		for _, id := range in.Id {
			contentMsgs = append(contentMsgs, msg.MakeTLContentMessage(&msg.TLContentMessage{
				Id: id,
			}))
		}

		return c.svcCtx.MsgClient.MsgReadMessageContents(c.ctx, &msg.TLMsgReadMessageContents{
			UserId:    userId,
			AuthKeyId: authKeyId,
			PeerType:  tg.PEER_USER, // default to USER peer
			PeerId:    peerId,
			Id:        contentMsgs,
		})
	}

	// Fallback placeholder when MsgClient is not available.
	pts := int32(1)
	ptsCount := int32(1)
	if len(in.Id) > 0 {
		ptsCount = int32(len(in.Id))
		for _, id := range in.Id {
			if id > pts {
				pts = id
			}
		}
	}

	return makeBffAffectedMessagesPlaceholder(pts, ptsCount), nil
}
