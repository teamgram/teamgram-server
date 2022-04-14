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

// MessagesGetAdminsWithInvites
// messages.getAdminsWithInvites#3920e6ef peer:InputPeer = messages.ChatAdminsWithInvites;
func (c *ChatInvitesCore) MessagesGetAdminsWithInvites(in *mtproto.TLMessagesGetAdminsWithInvites) (*mtproto.Messages_ChatAdminsWithInvites, error) {
	var (
		peer = mtproto.FromInputPeer2(c.MD.UserId, in.Peer)
	)

	if !peer.IsChat() {
		err := mtproto.ErrPeerIdInvalid
		c.Logger.Errorf("messages.getAdminsWithInvites - error: ", err)
		return nil, err
	}

	rAdmins, err := c.svcCtx.Dao.ChatClient.ChatGetAdminsWithInvites(c.ctx, &chatpb.TLChatGetAdminsWithInvites{
		SelfId: c.MD.UserId,
		ChatId: peer.PeerId,
	})
	if err != nil {
		c.Logger.Errorf("messages.getAdminsWithInvites - error: %v", err)
		return nil, err
	}

	rValues := mtproto.MakeTLMessagesChatAdminsWithInvites(&mtproto.Messages_ChatAdminsWithInvites{
		Admins: rAdmins.GetDatas(),
		Users:  []*mtproto.User{},
	}).To_Messages_ChatAdminsWithInvites()

	if len(rValues.Admins) == 0 {
		return rValues, nil
	}

	idHelper := mtproto.NewIDListHelper(c.MD.UserId)
	for _, a := range rValues.Admins {
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
