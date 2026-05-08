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

	"github.com/teamgram/teamgram-server/v2/app/messenger/msg/msg"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

// MsgResolveDialogCursorTopMessage
// msg.resolveDialogCursorTopMessage user_id:long top_message_id:int = ResolvedDialogCursor;
func (c *MsgCore) MsgResolveDialogCursorTopMessage(in *msg.TLMsgResolveDialogCursorTopMessage) (*msg.ResolvedDialogCursor, error) {
	if in == nil || in.UserId <= 0 || in.TopMessageId <= 0 {
		return nil, fmt.Errorf("%w: invalid dialog cursor resolver request", msg.ErrSendStateConflict)
	}
	resolved, err := c.svcCtx.Repo.ResolveMessageIDs(c.ctx, in.UserId, []int64{int64(in.TopMessageId)})
	if err != nil {
		return nil, err
	}
	if len(resolved) == 0 {
		return msg.MakeTLResolvedDialogCursor(&msg.TLResolvedDialogCursor{Found: tg.BoolFalseClazz}).ToResolvedDialogCursor(), nil
	}
	row := resolved[0]
	return msg.MakeTLResolvedDialogCursor(&msg.TLResolvedDialogCursor{
		Found:       tg.BoolTrueClazz,
		PeerType:    row.PeerType,
		PeerId:      row.PeerID,
		PeerSeq:     row.PeerSeq,
		MessageDate: row.MessageDate,
	}).ToResolvedDialogCursor(), nil
}
