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

// ContactsGetBlocked
// contacts.getBlocked#9a868f80 flags:# my_stories_from:flags.0?true offset:int limit:int = contacts.Blocked;
func (c *ContactsCore) ContactsGetBlocked(in *tg.TLContactsGetBlocked) (*tg.ContactsBlocked, error) {
	limit := in.Limit
	if limit > 50 {
		limit = 50
	}

	blockedList, err := c.svcCtx.Repo.UserClient.UserGetBlockedList(c.ctx, &userpb.TLUserGetBlockedList{
		UserId: c.MD.UserId,
		Offset: in.Offset,
		Limit:  limit,
	})
	if err != nil {
		return nil, err
	}
	if blockedList == nil || len(blockedList.Datas) == 0 {
		return tg.MakeTLContactsBlocked(&tg.TLContactsBlocked{
			Blocked: []tg.PeerBlockedClazz{},
			Chats:   []tg.ChatClazz{},
			Users:   []tg.UserClazz{},
		}).ToContactsBlocked(), nil
	}

	userIDs := make([]int64, 0, len(blockedList.Datas))
	for _, blocked := range blockedList.Datas {
		if blocked == nil {
			continue
		}
		peer := tg.FromPeer(blocked.PeerId)
		if peer.PeerType == tg.PEER_USER {
			userIDs = append(userIDs, peer.PeerId)
		}
	}

	users, err := c.svcCtx.Repo.UserClient.UserGetMutableUsers(c.ctx, &userpb.TLUserGetMutableUsers{Id: userIDs})
	if err != nil {
		return nil, err
	}

	var immutableUsers []tg.ImmutableUserClazz
	if users != nil {
		immutableUsers = users.Datas
	}
	return tg.MakeTLContactsBlocked(&tg.TLContactsBlocked{
		Blocked: blockedList.Datas,
		Chats:   []tg.ChatClazz{},
		Users:   projectUsersByIDs(immutableUsers, userIDs),
	}).ToContactsBlocked(), nil
}
