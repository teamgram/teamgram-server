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

// MessagesDeleteRevokedExportedChatInvites
// messages.deleteRevokedExportedChatInvites#56987bd5 peer:InputPeer admin_id:InputUser = Bool;
func (c *ChatInvitesCore) MessagesDeleteRevokedExportedChatInvites(in *mtproto.TLMessagesDeleteRevokedExportedChatInvites) (*mtproto.Bool, error) {
	var (
		err     error
		peer    = mtproto.FromInputPeer2(c.MD.UserId, in.Peer)
		adminId = mtproto.FromInputUser(c.MD.UserId, in.AdminId)
	)

	if !peer.IsChat() {
		err = mtproto.ErrPeerIdInvalid
		c.Logger.Errorf("messages.deleteRevokedExportedChatInvites - error: ", err)
		return nil, err
	}

	_, err = c.svcCtx.Dao.ChatClient.ChatDeleteRevokedExportedChatInvites(c.ctx, &chatpb.TLChatDeleteRevokedExportedChatInvites{
		SelfId:  c.MD.UserId,
		ChatId:  peer.PeerId,
		AdminId: adminId.PeerId,
	})
	if err != nil {
		c.Logger.Errorf("messages.deleteRevokedExportedChatInvites - error: %v", err)
		return nil, err
	}

	return mtproto.BoolTrue, nil
}
