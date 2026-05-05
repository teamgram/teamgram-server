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

// MessagesGetPeerDialogs
// messages.getPeerDialogs#e470bcfd peers:Vector<InputDialogPeer> = messages.PeerDialogs;
func (c *DialogsCore) MessagesGetPeerDialogs(in *tg.TLMessagesGetPeerDialogs) (*tg.MessagesPeerDialogs, error) {
	if c.MD == nil || c.MD.UserId <= 0 {
		return nil, tg.ErrUserIdInvalid
	}
	if in == nil {
		return nil, tg.ErrInputRequestInvalid
	}

	peers := make([]dialogpb.DialogPeerClazz, 0, len(in.Peers))
	for _, peer := range in.Peers {
		resolved, err := c.resolveInputDialogPeer(peer)
		if err != nil {
			return nil, err
		}
		peerType, err := dialogFacadePeerType(resolved.PeerType)
		if err != nil {
			return nil, err
		}
		peers = append(peers, dialogpb.MakeTLDialogPeer(&dialogpb.TLDialogPeer{
			PeerType: peerType,
			PeerId:   resolved.PeerId,
		}))
	}
	if len(peers) == 0 {
		return emptyPeerDialogs(), nil
	}

	result, err := c.svcCtx.Repo.DialogClient.DialogGetPeerDialogsV2(c.ctx, &dialogpb.TLDialogGetPeerDialogsV2{
		UserId: c.MD.UserId,
		Peers:  peers,
	})
	if err != nil {
		c.Logger.Errorf("messages.getPeerDialogs - dialog.getPeerDialogsV2 failed: user_id: %d, peer_count: %d, err: %v",
			c.MD.UserId, len(peers), err)
		return nil, tg.ErrInternalServerError
	}
	hydrated, err := c.hydrateDialogExtV2s("messages.getPeerDialogs", vectorDialogExtV2Datas(result))
	if err != nil {
		return nil, err
	}
	return tg.MakeTLMessagesPeerDialogs(&tg.TLMessagesPeerDialogs{
		Dialogs:  hydrated.Dialogs,
		Messages: hydrated.Messages,
		Chats:    hydrated.Chats,
		Users:    hydrated.Users,
		State:    emptyUpdatesState(),
	}).ToMessagesPeerDialogs(), nil
}
