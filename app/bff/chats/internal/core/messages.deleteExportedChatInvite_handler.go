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
	chatpb "github.com/teamgram/teamgram-server/app/service/biz/chat/chat"
)

// MessagesDeleteExportedChatInvite
// messages.deleteExportedChatInvite#d464a42b peer:InputPeer link:string = Bool;
func (c *ChatsCore) MessagesDeleteExportedChatInvite(in *mtproto.TLMessagesDeleteExportedChatInvite) (*mtproto.Bool, error) {
	var (
		peer = mtproto.FromInputPeer2(c.MD.UserId, in.Peer)
		err  error
	)

	switch peer.PeerType {
	case mtproto.PEER_CHAT:
		_, err = c.svcCtx.Dao.ChatClient.ChatDeleteExportedChatInvite(c.ctx, &chatpb.TLChatDeleteExportedChatInvite{
			SelfId: c.MD.UserId,
			ChatId: peer.PeerId,
			Link:   in.GetLink(),
		})
		if err != nil {
			c.Logger.Errorf("messages.deleteExportedChatInvite - error: %v", err)
			return nil, err
		}
	case mtproto.PEER_CHANNEL:
		c.Logger.Errorf("messages.deleteExportedChatInvite blocked, License key from https://teamgram.net required to unlock enterprise features.")

		return nil, mtproto.ErrEnterpriseIsBlocked
	default:
		err = mtproto.ErrPeerIdInvalid
		c.Logger.Errorf("messages.deleteExportedChatInvite - error: ", err)
		return nil, err
	}

	return mtproto.BoolTrue, nil
}
