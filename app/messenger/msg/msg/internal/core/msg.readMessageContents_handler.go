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

var _ *tg.Bool

// MsgReadMessageContents
// msg.readMessageContents user_id:long auth_key_id:long peer_type:int peer_id:long id:Vector<ContentMessage> = messages.AffectedMessages;
func (c *MsgCore) MsgReadMessageContents(in *msg.TLMsgReadMessageContents) (*tg.MessagesAffectedMessages, error) {
	pts := int32(1)
	ptsCount := int32(1)
	if len(in.Id) > 0 {
		ptsCount = int32(len(in.Id))
		for _, content := range in.Id {
			if content == nil {
				continue
			}
			if x := content.ToContentMessage(); x != nil && x.Id > pts {
				pts = x.Id
			}
		}
	}

	return tg.MakeTLMessagesAffectedMessages(&tg.TLMessagesAffectedMessages{
		Pts:      pts,
		PtsCount: ptsCount,
	}).ToMessagesAffectedMessages(), nil
}
