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

// MessagesCheckChatInvite
// messages.checkChatInvite#3eadb1bb hash:string = ChatInvite;
func (c *ChatInvitesCore) MessagesCheckChatInvite(in *mtproto.TLMessagesCheckChatInvite) (*mtproto.ChatInvite, error) {
	// Code	Type	Description
	// 400	INVITE_HASH_EMPTY	The invite hash is empty.
	// 400	INVITE_HASH_EXPIRED	The invite link has expired.
	// 400	INVITE_HASH_INVALID	The invite hash is invalid.

	if len(in.Hash) == 0 {
		err := mtproto.ErrInviteHashEmpty
		c.Logger.Errorf("messages.checkChatInvite - error: %v", err)
		return nil, err
	}
	if len(in.Hash) != 20 {
		err := mtproto.ErrInviteHashInvalid
		c.Logger.Errorf("messages.checkChatInvite - error: %v", err)
		return nil, err
	}
	if !chatpb.IsChatInviteHash(in.Hash) {
		err := mtproto.ErrInviteHashInvalid
		c.Logger.Errorf("messages.checkChatInvite - error: %v", err)
		return nil, err
	}

	getUserListF := func(idList []int64) []*mtproto.User {
		users, _ := c.svcCtx.Dao.UserClient.UserGetMutableUsers(c.ctx, &userpb.TLUserGetMutableUsers{
			Id: append(idList, c.MD.UserId),
		})
		return users.GetUserListByIdList(c.MD.UserId, idList...)
	}

	chatInviteExt, err := c.svcCtx.Dao.ChatClient.ChatCheckChatInvite(c.ctx, &chatpb.TLChatCheckChatInvite{
		SelfId: c.MD.UserId,
		Hash:   in.Hash,
	})
	if err != nil {
		c.Logger.Errorf("messages.checkChatInvite - error: %v", err)
		return nil, err
	}

	rValue := chatInviteExt.ToChatInvite(c.MD.UserId, func(idList []int64) []*mtproto.User {
		return getUserListF(idList)
	})
	if rValue == nil {
		err = mtproto.ErrInternalServerError
		c.Logger.Errorf("messages.checkChatInvite - error: ", err)
		return nil, err
	}

	return rValue, nil
}
