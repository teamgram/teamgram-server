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
)

// MessagesReorderPinnedDialogs
// messages.reorderPinnedDialogs#3b1adf37 flags:# force:flags.0?true folder_id:int order:Vector<InputDialogPeer> = Bool;
func (c *DialogsCore) MessagesReorderPinnedDialogs(in *mtproto.TLMessagesReorderPinnedDialogs) (*mtproto.Bool, error) {
	if len(in.GetOrder()) == 0 {
		c.Logger.Errorf("messages.reorderPinnedDialogs - len(order) == 0")
		return mtproto.BoolTrue, nil
	}

	var (
		peerDialogIdList []int64
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
				peerDialogIdList = append(peerDialogIdList, p.PeerId)
			default:
				err := mtproto.ErrPeerIdInvalid
				c.Logger.Errorf("messages.reorderPinnedDialogs - error: %v", err)
				return nil, err
			}
		default:
			err := mtproto.ErrPeerIdInvalid
			c.Logger.Errorf("messages.reorderPinnedDialogs - error: %v", err)
			return nil, err
		}
	}

	_, err := c.svcCtx.Dao.DialogClient.DialogReorderPinnedDialogs(c.ctx, &dialog.TLDialogReorderPinnedDialogs{
		UserId:   c.MD.UserId,
		Force:    mtproto.ToBool(in.Force),
		FolderId: in.FolderId,
		IdList:   peerDialogIdList,
	})
	if err != nil {
		c.Logger.Errorf("messages.reorderPinnedDialogs - error: %v", err)
		return nil, err
	}

	return mtproto.BoolTrue, nil
}
