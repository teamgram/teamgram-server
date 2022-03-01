// Copyright 2022 Teamgram Authors
//  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Author: teamgramio (teamgram.io@gmail.com)
//

package core

import (
	"github.com/teamgram/proto/mtproto"
	"github.com/teamgram/teamgram-server/app/service/biz/dialog/dialog"
	"github.com/teamgram/teamgram-server/app/service/biz/updates/updates"
)

// MessagesGetPinnedDialogs
// messages.getPinnedDialogs#d6b94df2 folder_id:int = messages.PeerDialogs;
func (c *DialogsCore) MessagesGetPinnedDialogs(in *mtproto.TLMessagesGetPinnedDialogs) (*mtproto.Messages_PeerDialogs, error) {
	var (
		dialogList dialog.DialogExtList
		folderId   = in.GetFolderId()
		state      *mtproto.Updates_State
	)

	if folderId != 0 && folderId != 1 {
		err := mtproto.ErrFolderIdInvalid
		c.Logger.Errorf("messages.getPinnedDialogs - error: %v", err)
		return nil, err
	}

	state, err := c.svcCtx.Dao.UpdatesClient.UpdatesGetState(c.ctx, &updates.TLUpdatesGetState{
		AuthKeyId: c.MD.AuthId,
		UserId:    c.MD.UserId,
	})
	if err != nil {
		c.Logger.Errorf("messages.getPinnedDialogs - error: %v", err)
		return nil, mtproto.ErrInternelServerError
	}

	if folderId == 0 {
		if dialogFolder, err := c.svcCtx.Dao.DialogClient.DialogGetDialogFolder(c.ctx, &dialog.TLDialogGetDialogFolder{
			UserId:   0,
			FolderId: 0,
		}); err != nil {
			c.Logger.Errorf("messages.getPinnedDialogs - error: %v", err)
			return nil, err
		} else {
			dialogList = append(dialogList, dialogFolder.GetDatas()...)
		}
	}

	if pinnedDialogList, err := c.svcCtx.Dao.DialogClient.DialogGetPinnedDialogs(c.ctx, &dialog.TLDialogGetPinnedDialogs{
		UserId:   c.MD.UserId,
		FolderId: folderId,
	}); err != nil {
		c.Logger.Errorf("messages.getPinnedDialogs - error: %v", err)
		return nil, err
	} else {
		dialogList = append(dialogList, pinnedDialogList.GetDatas()...)
	}

	return c.makeMessagesDialogs(dialogList).ToMessagesPeerDialogs(state), nil
}
