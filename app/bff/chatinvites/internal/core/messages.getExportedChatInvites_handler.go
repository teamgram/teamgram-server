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

// MessagesGetExportedChatInvites
// messages.getExportedChatInvites#a2b5a3f6 flags:# revoked:flags.3?true peer:InputPeer admin_id:InputUser offset_date:flags.2?int offset_link:flags.2?string limit:int = messages.ExportedChatInvites;
func (c *ChatInvitesCore) MessagesGetExportedChatInvites(in *mtproto.TLMessagesGetExportedChatInvites) (*mtproto.Messages_ExportedChatInvites, error) {
	var (
		peer    = mtproto.FromInputPeer2(c.MD.UserId, in.Peer)
		adminId = mtproto.FromInputUser(c.MD.UserId, in.AdminId)
		// limit   = in.GetLimit()
	)

	if !peer.IsChat() {
		err := mtproto.ErrPeerIdInvalid
		c.Logger.Errorf("messages.getExportedChatInvites - error: ", err)
		return nil, err
	}

	// TODO: check adminId
	rInvites, err := c.svcCtx.Dao.ChatClient.ChatGetExportedChatInvites(c.ctx, &chatpb.TLChatGetExportedChatInvites{
		ChatId:     peer.PeerId,
		AdminId:    adminId.PeerId,
		Revoked:    in.Revoked,
		OffsetDate: in.OffsetDate,
		OffsetLink: in.OffsetLink,
		Limit:      in.GetLimit(),
	})
	if err != nil {
		c.Logger.Errorf("messages.getExportedChatInvites - error: ", err)
		return nil, err
	}

	rValues := mtproto.MakeTLMessagesExportedChatInvites(&mtproto.Messages_ExportedChatInvites{
		Count:   0,
		Invites: []*mtproto.ExportedChatInvite{},
		Users:   []*mtproto.User{},
	}).To_Messages_ExportedChatInvites()

	rValues.Count = int32(len(rInvites.Datas))
	rValues.Invites = rInvites.Datas

	if len(rValues.Invites) == 0 {
		return rValues, nil
	}

	idHelper := mtproto.NewIDListHelper(c.MD.UserId)
	for _, a := range rValues.Invites {
		idHelper.AppendUsers(a.AdminId)
	}

	idHelper.Visit(func(userIdList []int64) {
		users, _ := c.svcCtx.Dao.UserClient.UserGetMutableUsers(c.ctx, &userpb.TLUserGetMutableUsers{
			Id: userIdList,
		})
		rValues.Users = users.GetUserListByIdList(c.MD.UserId, userIdList...)
	}, nil, nil)

	return rValues, nil
}
