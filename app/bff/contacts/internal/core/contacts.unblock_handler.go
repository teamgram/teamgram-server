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
	userpb "github.com/teamgram/teamgram-server/app/service/biz/user/user"
)

// ContactsUnblock
// contacts.unblock#bea65d50 id:InputPeer = Bool;
func (c *ContactsCore) ContactsUnblock(in *mtproto.TLContactsUnblock) (*mtproto.Bool, error) {
	var (
		mUsers    *userpb.Vector_ImmutableUser
		err       error
		unblockId = mtproto.FromInputPeer2(c.MD.UserId, in.GetId())
		idHelper  = mtproto.NewIDListHelper(c.MD.UserId)
	)

	if !unblockId.IsUser() || unblockId.IsSelf() {
		err = mtproto.ErrPeerIdInvalid
		c.Logger.Errorf("contacts.block - error: %v", err)
		return nil, err
	}

	if unblockId.PeerId == c.MD.UserId {
		err = mtproto.ErrContactIdInvalid
		c.Logger.Errorf("contacts.block - error: %v", err)
		return nil, err
	}

	// TODO
	/*
		auto BlockPeerBoxController::createRow(not_null<History*> history)
		-> std::unique_ptr<BlockPeerBoxController::Row> {
			if (!history->peer->isUser()
				|| history->peer->isServiceUser()
				|| history->peer->isSelf()
				|| history->peer->isRepliesChat()) {
				return nullptr;
			}
			auto row = std::make_unique<Row>(history);
			updateIsBlocked(row.get(), history->peer);
			return row;
		}
	*/

	mUsers, err = c.svcCtx.Dao.UserClient.UserGetMutableUsers(c.ctx, &userpb.TLUserGetMutableUsers{
		Id: []int64{c.MD.UserId, unblockId.PeerId},
	})

	// me, _ := users.GetImmutableUser(c.MD.UserId)
	blocked, _ := mUsers.GetImmutableUser(unblockId.PeerId)

	if blocked == nil {
		err = mtproto.ErrContactIdInvalid
		c.Logger.Errorf("contacts.unblock - error: %v", err)
		return nil, err
	} else if blocked.GetUser().GetDeleted() {
		err = mtproto.ErrInputUserDeactivated
		c.Logger.Errorf("contacts.unblock - error: %v", err)
		return nil, err
	}
	c.svcCtx.Dao.UserClient.UserUnBlockPeer(c.ctx, &userpb.TLUserUnBlockPeer{
		UserId:   c.MD.UserId,
		PeerType: unblockId.PeerType,
		PeerId:   unblockId.PeerId,
	})
	idHelper.AppendUsers(unblockId.PeerId)

	syncUpdates := mtproto.MakeUpdatesByUpdates(mtproto.MakeTLUpdatePeerBlocked(&mtproto.Update{
		Blocked_BOOL:         mtproto.BoolFalse,
		Blocked_FLAGBOOLEAN:  false,
		BlockedMyStoriesFrom: false,
		PeerId:               unblockId.ToPeer(),
	}).To_Update())

	idHelper.Visit(
		func(userIdList []int64) {
			syncUpdates.PushUser(mUsers.GetUserListByIdList(c.MD.UserId, userIdList...)...)
		},
		func(chatIdList []int64) {
		},
		func(channelIdList []int64) {
		})

	c.svcCtx.Dao.SyncClient.SyncUpdatesNotMe(c.ctx, &sync.TLSyncUpdatesNotMe{
		UserId:        c.MD.UserId,
		PermAuthKeyId: c.MD.PermAuthKeyId,
		Updates:       syncUpdates,
	})

	return mtproto.BoolTrue, nil
}
