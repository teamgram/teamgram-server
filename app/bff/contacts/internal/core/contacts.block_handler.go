// Copyright (c) 2026 The Teamgram Authors. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Author: teamgramio (teamgram.io@gmail.com)

package core

import (
	userpb "github.com/teamgram/teamgram-server/v2/app/service/biz/user/user"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

// ContactsBlock
// contacts.block#2e2e8734 flags:# my_stories_from:flags.0?true id:InputPeer = Bool;
func (c *ContactsCore) ContactsBlock(in *tg.TLContactsBlock) (*tg.Bool, error) {
	blockID := tg.FromInputPeer2(c.MD.UserId, in.Id)
	if blockID.PeerType != tg.PEER_USER {
		return nil, tg.Err400PeerIdInvalid
	}
	if blockID.PeerId == c.MD.UserId {
		return nil, tg.ErrContactIdInvalid
	}

	users, err := c.svcCtx.Repo.UserClient.UserGetMutableUsers(c.ctx, &userpb.TLUserGetMutableUsers{
		Id: []int64{c.MD.UserId, blockID.PeerId},
		To: []int64{c.MD.UserId},
	})
	if err != nil {
		return nil, err
	}

	var immutableUsers []tg.ImmutableUserClazz
	if users != nil {
		immutableUsers = users.Datas
	}
	blocked := immutableUserByID(immutableUsers, blockID.PeerId)
	if blocked == nil {
		return nil, tg.ErrContactIdInvalid
	}
	if blocked.User != nil && blocked.User.Deleted {
		return nil, tg.ErrInputUserDeactivated
	}

	if _, err = c.svcCtx.Repo.UserClient.UserBlockPeer(c.ctx, &userpb.TLUserBlockPeer{
		UserId:   c.MD.UserId,
		PeerType: blockID.PeerType,
		PeerId:   blockID.PeerId,
	}); err != nil {
		return nil, err
	}

	// TODO: master sends updatePeerBlocked through sync.SyncUpdatesNotMe here.
	return tg.BoolTrue, nil
}
