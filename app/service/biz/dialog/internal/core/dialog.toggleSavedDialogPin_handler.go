// Copyright 2024 Teamgram Authors
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
	"time"

	"github.com/teamgram/proto/mtproto"
	"github.com/teamgram/teamgram-server/app/service/biz/dialog/dialog"
)

// DialogToggleSavedDialogPin
// dialog.toggleSavedDialogPin user_id:long peer:PeerUtil pinned:Bool = Bool;
func (c *DialogCore) DialogToggleSavedDialogPin(in *dialog.TLDialogToggleSavedDialogPin) (*mtproto.Bool, error) {
	var (
		peer   = in.Peer
		pinned int64
	)

	if mtproto.FromBool(in.Pinned) {
		pinned = time.Now().Unix() << 32
	} else {
		pinned = 0
	}

	_, err := c.svcCtx.Dao.SavedDialogsDAO.UpdateUserPeerPinned(
		c.ctx,
		pinned,
		in.GetUserId(),
		peer.PeerType,
		peer.PeerId)
	if err != nil {
		c.Logger.Error("toggleSavedDialogPin - query saved_dialogs error:", err)
		return nil, err
	}

	return mtproto.BoolTrue, nil
}
