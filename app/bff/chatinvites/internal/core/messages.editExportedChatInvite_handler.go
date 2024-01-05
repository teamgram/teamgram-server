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

// MessagesEditExportedChatInvite
// messages.editExportedChatInvite#bdca2f75 flags:# revoked:flags.2?true peer:InputPeer link:string expire_date:flags.0?int usage_limit:flags.1?int request_needed:flags.3?Bool title:flags.4?string = messages.ExportedChatInvite;
func (c *ChatInvitesCore) MessagesEditExportedChatInvite(in *mtproto.TLMessagesEditExportedChatInvite) (*mtproto.Messages_ExportedChatInvite, error) {
	var (
		peer    = mtproto.FromInputPeer2(c.MD.UserId, in.Peer)
		invites []*mtproto.ExportedChatInvite
		rValue  *mtproto.Messages_ExportedChatInvite
	)

	if !peer.IsChat() {
		err := mtproto.ErrPeerIdInvalid
		c.Logger.Errorf("messages.editExportedChatInvite - error: ", err)
		return nil, err
	}

	chatInvites, err := c.svcCtx.Dao.ChatClient.ChatEditExportedChatInvite(c.ctx, &chatpb.TLChatEditExportedChatInvite{
		SelfId:        c.MD.UserId,
		ChatId:        peer.PeerId,
		Revoked:       in.Revoked,
		Link:          in.Link,
		ExpireDate:    in.ExpireDate,
		UsageLimit:    in.UsageLimit,
		RequestNeeded: in.RequestNeeded,
		Title:         in.Title,
	})
	if err != nil {
		c.Logger.Errorf("messages.editExportedChatInvite - error: %v", err)
		return nil, err
	}
	if len(chatInvites.Datas) == 0 || len(chatInvites.Datas) > 2 {
		err = mtproto.ErrInternalServerError
		c.Logger.Errorf("messages.editExportedChatInvite - error: %v", err)
		return nil, err
	}

	invites = chatInvites.Datas

	users, err2 := c.svcCtx.Dao.UserClient.UserGetMutableUsers(c.ctx, &userpb.TLUserGetMutableUsers{
		Id: []int64{c.MD.UserId, invites[0].AdminId},
	})
	if err2 != nil {
		err := mtproto.ErrPeerIdInvalid
		c.Logger.Errorf("messages.editExportedChatInvite - error: ", err)
		return nil, err
	}
	inviteUsers := users.GetUserListByIdList(c.MD.UserId, invites[0].AdminId)

	if len(invites) == 1 {
		rValue = mtproto.MakeTLMessagesExportedChatInvite(&mtproto.Messages_ExportedChatInvite{
			Invite: invites[0],
			Users:  inviteUsers,
		}).To_Messages_ExportedChatInvite()
	} else {
		rValue = mtproto.MakeTLMessagesExportedChatInviteReplaced(&mtproto.Messages_ExportedChatInvite{
			Invite:    invites[0],
			NewInvite: invites[1],
			Users:     inviteUsers,
		}).To_Messages_ExportedChatInvite()
	}

	return rValue, nil
}
