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
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

// MessagesGetPinnedDialogs
// messages.getPinnedDialogs#d6b94df2 folder_id:int = messages.PeerDialogs;
func (c *DialogsCore) MessagesGetPinnedDialogs(in *tg.TLMessagesGetPinnedDialogs) (*tg.MessagesPeerDialogs, error) {
	if c.MD == nil || c.MD.UserId <= 0 {
		return nil, tg.ErrUserIdInvalid
	}
	if in == nil {
		return nil, tg.ErrInputRequestInvalid
	}
	if in.FolderId != 0 && in.FolderId != 1 {
		return nil, tg.ErrFolderIdInvalid
	}

	return tg.MakeTLMessagesPeerDialogs(&tg.TLMessagesPeerDialogs{
		Dialogs:  []tg.DialogClazz{},
		Messages: []tg.MessageClazz{},
		Chats:    []tg.ChatClazz{},
		Users:    []tg.UserClazz{},
		State:    emptyUpdatesState(),
	}).ToMessagesPeerDialogs(), nil
}

func emptyUpdatesState() tg.UpdatesStateClazz {
	return tg.MakeTLUpdatesState(&tg.TLUpdatesState{
		Pts:         0,
		Qts:         0,
		Date:        0,
		Seq:         0,
		UnreadCount: 0,
	})
}
