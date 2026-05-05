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
	dialogpb "github.com/teamgram/teamgram-server/v2/app/service/biz/dialog/dialog"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

// MessagesGetSavedDialogs
// messages.getSavedDialogs#1e91fc99 flags:# exclude_pinned:flags.0?true parent_peer:flags.1?InputPeer offset_date:int offset_id:int offset_peer:InputPeer limit:int hash:long = messages.SavedDialogs;
func (c *SavedMessageDialogsCore) MessagesGetSavedDialogs(in *tg.TLMessagesGetSavedDialogs) (*tg.MessagesSavedDialogs, error) {
	if c.MD == nil || c.MD.UserId <= 0 {
		return nil, tg.ErrUserIdInvalid
	}
	if in == nil {
		return nil, tg.ErrInputRequestInvalid
	}
	got, err := c.svcCtx.Repo.DialogClient.DialogGetSavedDialogs(c.ctx, &dialogpb.TLDialogGetSavedDialogs{
		UserId:        c.MD.UserId,
		ExcludePinned: tg.ToBoolClazz(in.ExcludePinned),
		OffsetDate:    in.OffsetDate,
		OffsetId:      in.OffsetId,
		OffsetPeer:    tg.MakeTLPeerUtil(&tg.TLPeerUtil{}),
		Limit:         in.Limit,
	})
	if err != nil {
		c.Logger.Errorf("messages.getSavedDialogs - dialog.getSavedDialogs failed: user_id: %d err: %v", c.MD.UserId, err)
		return nil, tg.ErrInternalServerError
	}
	return makeMessagesSavedDialogs(got), nil
}
