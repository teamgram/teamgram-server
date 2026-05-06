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

// MessagesReorderPinnedSavedDialogs
// messages.reorderPinnedSavedDialogs#8b716587 flags:# force:flags.0?true order:Vector<InputDialogPeer> = Bool;
func (c *SavedMessageDialogsCore) MessagesReorderPinnedSavedDialogs(in *tg.TLMessagesReorderPinnedSavedDialogs) (*tg.Bool, error) {
	if c.MD == nil || c.MD.UserId <= 0 {
		return nil, tg.ErrUserIdInvalid
	}
	if c.MD.PermAuthKeyId == 0 {
		return nil, tg.ErrAuthKeyPermEmpty
	}
	if in == nil {
		return nil, tg.ErrInputRequestInvalid
	}
	order := make([]tg.PeerUtilClazz, 0, len(in.Order))
	for _, input := range in.Order {
		peer, err := c.resolveInputDialogPeer(input)
		if err != nil {
			return nil, err
		}
		facadePeer, err := savedDialogFacadePeerUtil(peer)
		if err != nil {
			return nil, err
		}
		order = append(order, facadePeer)
	}
	if _, err := c.svcCtx.Repo.DialogClient.DialogReorderPinnedSavedDialogs(c.ctx, &dialogpb.TLDialogReorderPinnedSavedDialogs{
		UserId: c.MD.UserId,
		Force:  tg.ToBoolClazz(in.Force),
		Order:  order,
	}); err != nil {
		c.Logger.Errorf("messages.reorderPinnedSavedDialogs - dialog.reorderPinnedSavedDialogs failed: user_id: %d err: %v", c.MD.UserId, err)
		return nil, tg.ErrInternalServerError
	}
	return tg.BoolTrue, nil
}
