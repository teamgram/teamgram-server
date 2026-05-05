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
	"github.com/teamgram/teamgram-server/v2/app/messenger/msg/msg"
	"github.com/teamgram/teamgram-server/v2/app/messenger/userupdates/payload"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

// MessagesDeleteMessages
// messages.deleteMessages#e58e95d2 flags:# revoke:flags.0?true id:Vector<int> = messages.AffectedMessages;
func (c *MessagesCore) MessagesDeleteMessages(in *tg.TLMessagesDeleteMessages) (*tg.MessagesAffectedMessages, error) {
	md := c.MD
	if md == nil || md.UserId <= 0 {
		return nil, tg.ErrUserIdInvalid
	}
	if in == nil || len(in.Id) == 0 {
		return nil, tg.ErrInputRequestInvalid
	}
	var deleteClient deleteMessagesClient = c.svcCtx.Repo.MsgClient
	r, err := deleteClient.MsgDeleteMessages(c.ctx, &msg.TLMsgDeleteMessages{
		UserId:    md.UserId,
		AuthKeyId: md.PermAuthKeyId,
		PeerType:  payload.PeerTypeUser,
		PeerId:    md.UserId,
		Revoke:    in.Revoke,
		Id:        in.Id,
	})
	if err != nil {
		c.Logger.Errorf("messages.deleteMessages - msg error: self_user_id: %d, ids: %v, revoke: %t, err: %v", md.UserId, in.Id, in.Revoke, err)
		return nil, mapMsgSendError(err)
	}
	return r, nil
}
