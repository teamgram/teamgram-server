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
	"github.com/teamgram/proto/mtproto"
	"github.com/teamgram/teamgram-server/app/service/biz/dialog/dialog"
)

// MessagesReorderPinnedSavedDialogs
// messages.reorderPinnedSavedDialogs#8b716587 flags:# force:flags.0?true order:Vector<InputDialogPeer> = Bool;
func (c *SavedMessageDialogsCore) MessagesReorderPinnedSavedDialogs(in *mtproto.TLMessagesReorderPinnedSavedDialogs) (*mtproto.Bool, error) {
	if len(in.GetOrder()) == 0 {
		c.Logger.Errorf("messages.reorderPinnedDialogs - len(order) == 0")
		return mtproto.BoolTrue, nil
	}

	var (
		order []*mtproto.PeerUtil
	)
	for _, peer := range in.GetOrder() {
		switch peer.PredicateName {
		case mtproto.Predicate_inputDialogPeer:
			p := mtproto.FromInputPeer2(c.MD.UserId, peer.Peer)
			switch p.PeerType {
			case mtproto.PEER_SELF,
				mtproto.PEER_USER,
				mtproto.PEER_CHAT,
				mtproto.PEER_CHANNEL:
				order = append(order, p)
			default:
				err := mtproto.ErrPeerIdInvalid
				c.Logger.Errorf("messages.reorderPinnedSavedDialogs - error: %v", err)
				return nil, err
			}
		default:
			err := mtproto.ErrPeerIdInvalid
			c.Logger.Errorf("messages.reorderPinnedSavedDialogs - error: %v", err)
			return nil, err
		}
	}

	_, err := c.svcCtx.Dao.DialogClient.DialogReorderPinnedSavedDialogs(c.ctx, &dialog.TLDialogReorderPinnedSavedDialogs{
		UserId: c.MD.UserId,
		Force:  mtproto.ToBool(in.Force),
		Order:  order,
	})
	if err != nil {
		c.Logger.Errorf("messages.reorderPinnedSavedDialogs - error: %v", err)
		return nil, err
	}

	return mtproto.BoolTrue, nil
}
