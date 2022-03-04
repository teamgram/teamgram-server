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
	"github.com/teamgram/teamgram-server/app/service/biz/chat/chat"
	userpb "github.com/teamgram/teamgram-server/app/service/biz/user/user"
	"github.com/teamgram/teamgram-server/app/service/biz/username/username"
)

// ContactsResolveUsername
// contacts.resolveUsername#f93ccba3 username:string = contacts.ResolvedPeer;
func (c *UsernamesCore) ContactsResolveUsername(in *mtproto.TLContactsResolveUsername) (*mtproto.Contacts_ResolvedPeer, error) {
	// TODO(@benqi):
	// 401	AUTH_KEY_PERM_EMPTY	The temporary auth key must be binded to the permanent auth key to use these methods.
	// 401	SESSION_PASSWORD_NEEDED	2FA is enabled, use a password to login
	// 400	USERNAME_INVALID	The provided username is not valid
	// 400	USERNAME_NOT_OCCUPIED	The provided username is not occupied
	//
	var (
		peer *mtproto.PeerUtil
	)

	id := userpb.GetBotIdByName(in.GetUsername())
	if id > 0 {
		peer = mtproto.MakeUserPeerUtil(id)
	} else {
		rName, err := c.svcCtx.Dao.UsernameClient.UsernameResolveUsername(c.ctx, &username.TLUsernameResolveUsername{
			Username: in.GetUsername(),
		})
		if err != nil {
			c.Logger.Errorf("contacts.resolveUsername - reply: {%v}", err)
			return nil, err
		}

		peer = mtproto.FromPeer(rName)
	}

	resolvedPeer := mtproto.MakeTLContactsResolvedPeer(&mtproto.Contacts_ResolvedPeer{
		Peer:  peer.ToPeer(),
		Chats: []*mtproto.Chat{},
		Users: []*mtproto.User{},
	}).To_Contacts_ResolvedPeer()

	switch peer.PeerType {
	case mtproto.PEER_USER:
		mUsers, _ := c.svcCtx.Dao.UserClient.UserGetMutableUsers(c.ctx, &userpb.TLUserGetMutableUsers{
			Id: []int64{c.MD.UserId, peer.PeerId},
		})
		// .UserFacade.GetUserById(ctx, md.UserId, peer.PeerId)
		if mUsers != nil {
			resolvedPeer.Users = mUsers.GetUserListByIdList(c.MD.UserId, peer.PeerId)
		}
	case mtproto.PEER_CHAT:
		chat, _ := c.svcCtx.Dao.ChatClient.ChatGetChatBySelfId(c.ctx, &chat.TLChatGetChatBySelfId{
			SelfId: c.MD.UserId,
			ChatId: peer.PeerId,
		})
		if chat != nil {
			resolvedPeer.Chats = []*mtproto.Chat{chat.ToUnsafeChat(c.MD.UserId)}
		}
	case mtproto.PEER_CHANNEL:
		if c.svcCtx.Plugin != nil {
			resolvedPeer.Chats = c.svcCtx.Plugin.GetChannelListByIdList(c.ctx, c.MD.UserId, peer.PeerId)
		} else {
			c.Logger.Errorf("contacts.resolveUsername blocked, License key from https://teamgram.net required to unlock enterprise features.")
		}
	}

	return resolvedPeer, nil
}
