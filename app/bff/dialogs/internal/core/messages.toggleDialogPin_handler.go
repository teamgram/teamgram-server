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

// MessagesToggleDialogPin
// messages.toggleDialogPin#a731e257 flags:# pinned:flags.0?true peer:InputDialogPeer = Bool;
func (c *DialogsCore) MessagesToggleDialogPin(in *tg.TLMessagesToggleDialogPin) (*tg.Bool, error) {
	if c.MD == nil || c.MD.UserId <= 0 {
		return nil, tg.ErrUserIdInvalid
	}
	if c.MD.PermAuthKeyId == 0 {
		return nil, tg.ErrAuthKeyPermEmpty
	}
	if in == nil {
		return nil, tg.ErrInputRequestInvalid
	}
	peer, err := c.resolveInputDialogPeer(in.Peer)
	if err != nil {
		return nil, err
	}
	peerType, err := dialogFacadePeerType(peer.PeerType)
	if err != nil {
		return nil, err
	}
	token := dialogOperationToken()
	operationID := dialogOperationID("toggle_pin", c.MD.UserId, token)
	_, err = c.svcCtx.Repo.DialogClient.DialogToggleDialogPin(c.ctx, &dialogpb.TLDialogToggleDialogPin{
		UserId:              c.MD.UserId,
		PeerType:            peerType,
		PeerId:              peer.PeerId,
		Pinned:              tg.ToBoolClazz(in.Pinned),
		SourcePermAuthKeyId: c.MD.PermAuthKeyId,
		OperationId:         operationID,
		OutboxId:            dialogOutboxID(operationID),
	})
	if err != nil {
		c.Logger.Errorf("messages.toggleDialogPin - dialog.toggleDialogPin failed: user_id: %d, peer_type: %d, peer_id: %d, err: %v",
			c.MD.UserId, peerType, peer.PeerId, err)
		return nil, tg.ErrInternalServerError
	}
	return tg.BoolTrue, nil
}
