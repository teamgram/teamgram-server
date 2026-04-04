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

// MessagesDeleteMessages
// messages.deleteMessages#e58e95d2 flags:# revoke:flags.0?true id:Vector<int> = messages.AffectedMessages;
func (c *MessagesCore) MessagesDeleteMessages(in *tg.TLMessagesDeleteMessages) (*tg.MessagesAffectedMessages, error) {
	var userId int64
	if c.MD != nil {
		userId = c.MD.UserId
	}

	// MessagesDeleteMessages is broadcast to all relevant peers, so we need
	// at least a valid peer to route the request. Use the first message ID if available.
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

		return c.svcCtx.MsgClient.MsgDeleteMessages(c.ctx, &msg.TLMsgDeleteMessages{
			UserId:    userId,
			AuthKeyId: authKeyId,
			PeerType:  tg.PEER_USER, // deleteMessages is broadcast, use USER as default
			PeerId:    peerId,
			Revoke:    in.Revoke,
			Id:        in.Id,
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
