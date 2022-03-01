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
	"github.com/teamgram/teamgram-server/app/messenger/sync/sync"
	chatpb "github.com/teamgram/teamgram-server/app/service/biz/chat/chat"
	userpb "github.com/teamgram/teamgram-server/app/service/biz/user/user"
)

// ContactsBlock
// contacts.block#68cc1411 id:InputPeer = Bool;
func (c *ContactsCore) ContactsBlock(in *mtproto.TLContactsBlock) (*mtproto.Bool, error) {
	var (
		mUsers   *userpb.Vector_ImmutableUser
		mChat    *chatpb.MutableChat
		err      error
		blockId  = mtproto.FromInputPeer2(c.MD.UserId, in.GetId())
		idHelper = mtproto.NewIDListHelper(c.MD.UserId)
	)

	switch blockId.PeerType {
	case mtproto.PEER_SELF:
		err = mtproto.ErrContactIdInvalid
		c.Logger.Errorf("contacts.block - error: %v", err)
		return nil, err
	case mtproto.PEER_USER:
		mUsers, err = c.svcCtx.Dao.UserClient.UserGetMutableUsers(c.ctx, &userpb.TLUserGetMutableUsers{
			Id: []int64{c.MD.UserId, blockId.PeerId},
		})

		// me, _ := users.GetImmutableUser(c.MD.UserId)
		blocked, _ := mUsers.GetImmutableUser(blockId.PeerId)

		if blocked == nil {
			err = mtproto.ErrContactIdInvalid
			c.Logger.Errorf("contacts.block - error: %v", err)
			return nil, err
		} else if blocked.GetUser().GetDeleted() {
			err = mtproto.ErrInputUserDeactivated
			c.Logger.Errorf("contacts.block - error: %v", err)
			return nil, err
		}
		c.svcCtx.Dao.UserClient.UserBlockPeer(c.ctx, &userpb.TLUserBlockPeer{
			UserId:   c.MD.UserId,
			PeerType: blockId.PeerType,
			PeerId:   blockId.PeerId,
		})
		idHelper.AppendUsers(blockId.PeerId)
	case mtproto.PEER_CHAT:
		mChat, _ = c.svcCtx.Dao.ChatClient.ChatGetMutableChat(c.ctx, &chatpb.TLChatGetMutableChat{
			ChatId: blockId.PeerId,
		})
		if mChat == nil {
			err = mtproto.ErrPeerIdInvalid
			c.Logger.Errorf("contacts.block - error: %v", err)
			return nil, err
		}
		c.svcCtx.Dao.UserClient.UserBlockPeer(c.ctx, &userpb.TLUserBlockPeer{
			UserId:   c.MD.UserId,
			PeerType: blockId.PeerType,
			PeerId:   blockId.PeerId,
		})
		idHelper.AppendChats(blockId.PeerId)
	case mtproto.PEER_CHANNEL:
		c.Logger.Errorf("contacts.block blocked, License key from https://teamgram.net required to unlock enterprise features.")

		return nil, mtproto.ErrEnterpriseIsBlocked
	default:
		err := mtproto.ErrInputRequestInvalid
		c.Logger.Errorf("contacts.block - error: %v", err)
		return nil, err
	}

	syncUpdates := mtproto.MakeUpdatesByUpdates(mtproto.MakeTLUpdatePeerBlocked(&mtproto.Update{
		PeerId:  blockId.ToPeer(),
		Blocked: mtproto.BoolTrue,
	}).To_Update())

	idHelper.Visit(
		func(userIdList []int64) {
			syncUpdates.PushUser(mUsers.GetUserListByIdList(c.MD.UserId, userIdList...)...)
		},
		func(chatIdList []int64) {
			syncUpdates.PushChat(mChat.ToUnsafeChat(c.MD.UserId))
		},
		func(channelIdList []int64) {
			// TODO
		})

	c.svcCtx.Dao.SyncClient.SyncUpdatesNotMe(c.ctx, &sync.TLSyncUpdatesNotMe{
		UserId:    c.MD.UserId,
		AuthKeyId: c.MD.AuthId,
		Updates:   syncUpdates,
	})

	return mtproto.BoolTrue, nil
}
