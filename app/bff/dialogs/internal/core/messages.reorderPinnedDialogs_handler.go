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

// MessagesReorderPinnedDialogs
// messages.reorderPinnedDialogs#3b1adf37 flags:# force:flags.0?true folder_id:int order:Vector<InputDialogPeer> = Bool;
func (c *DialogsCore) MessagesReorderPinnedDialogs(in *tg.TLMessagesReorderPinnedDialogs) (*tg.Bool, error) {
	if c.MD == nil || c.MD.UserId <= 0 {
		return nil, tg.ErrUserIdInvalid
	}
	if c.MD.PermAuthKeyId == 0 {
		return nil, tg.ErrAuthKeyPermEmpty
	}
	if in == nil {
		return nil, tg.ErrInputRequestInvalid
	}
	if in.FolderId < 0 {
		return nil, tg.ErrFolderIdInvalid
	}

	ids := make([]int64, 0, len(in.Order))
	for _, input := range in.Order {
		peer, err := c.resolveInputDialogPeer(input)
		if err != nil {
			return nil, err
		}
		id, err := dialogFacadePeerDialogID(peer.PeerType, peer.PeerId)
		if err != nil {
			return nil, err
		}
		ids = append(ids, id)
	}
	token := dialogOperationToken()
	operationID := dialogOperationID("reorder_pinned", c.MD.UserId, token)
	_, err := c.svcCtx.Repo.DialogClient.DialogReorderPinnedDialogs(c.ctx, &dialogpb.TLDialogReorderPinnedDialogs{
		UserId:              c.MD.UserId,
		Force:               tg.ToBoolClazz(in.Force),
		FolderId:            in.FolderId,
		IdList:              ids,
		SourcePermAuthKeyId: c.MD.PermAuthKeyId,
		OperationId:         operationID,
		OutboxId:            dialogOutboxID(operationID),
	})
	if err != nil {
		c.Logger.Errorf("messages.reorderPinnedDialogs - dialog.reorderPinnedDialogs failed: user_id: %d, folder_id: %d, err: %v",
			c.MD.UserId, in.FolderId, err)
		return nil, tg.ErrInternalServerError
	}
	return tg.BoolTrue, nil
}
