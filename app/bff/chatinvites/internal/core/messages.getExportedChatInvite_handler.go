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
	userpb "github.com/teamgram/teamgram-server/app/service/biz/user/user"
)

// MessagesGetExportedChatInvite
// messages.getExportedChatInvite#73746f5c peer:InputPeer link:string = messages.ExportedChatInvite;
func (c *ChatInvitesCore) MessagesGetExportedChatInvite(in *mtproto.TLMessagesGetExportedChatInvite) (*mtproto.Messages_ExportedChatInvite, error) {
	var (
		peer               = mtproto.FromInputPeer2(c.MD.UserId, in.Peer)
		exportedChatInvite *mtproto.ExportedChatInvite
		err                error
	)

	if !peer.IsChat() {
		err := mtproto.ErrPeerIdInvalid
		c.Logger.Errorf("messages.getExportedChatInvite - error: ", err)
		return nil, err
	}

	exportedChatInvite, err = c.svcCtx.Dao.ChatClient.ChatGetExportedChatInvite(c.ctx, &chatpb.TLChatGetExportedChatInvite{
		ChatId: peer.PeerId,
		Link:   in.GetLink(),
	})
	if err != nil {
		c.Logger.Errorf("messages.getExportedChatInvite - error: ", err)
		return nil, err
	}

	users, err2 := c.svcCtx.Dao.UserClient.UserGetMutableUsers(c.ctx, &userpb.TLUserGetMutableUsers{
		Id: []int64{c.MD.UserId, exportedChatInvite.AdminId},
	})
	if err2 != nil {
		err = mtproto.ErrPeerIdInvalid
		c.Logger.Errorf("messages.getExportedChatInvite - error: ", err)
		return nil, err
	}

	return mtproto.MakeTLMessagesExportedChatInvite(&mtproto.Messages_ExportedChatInvite{
		Invite: exportedChatInvite,
		Users:  users.GetUserListByIdList(c.MD.UserId, exportedChatInvite.AdminId),
	}).To_Messages_ExportedChatInvite(), nil
}
