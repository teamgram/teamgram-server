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
	userprojection "github.com/teamgram/teamgram-server/v2/app/bff/internal/userprojection"
	userpb "github.com/teamgram/teamgram-server/v2/app/service/biz/user/user"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

// ContactsSearch
// contacts.search#11f812d8 q:string limit:int = contacts.Found;
func (c *ContactsCore) ContactsSearch(in *tg.TLContactsSearch) (*tg.ContactsFound, error) {
	limit := in.Limit
	if limit > 50 || limit == 0 {
		limit = 50
	}

	q := in.Q
	if q == "" {
		return nil, tg.ErrSearchQueryEmpty
	}
	if q[0] == '@' {
		q = q[1:]
	}
	if len(q) < 3 {
		return nil, tg.ErrQueryTooShort
	}

	found := tg.MakeTLContactsFound(&tg.TLContactsFound{
		MyResults: []tg.PeerClazz{},
		Results:   []tg.PeerClazz{},
		Users:     []tg.UserClazz{},
		Chats:     []tg.ChatClazz{},
	}).ToContactsFound()

	contacts, err := c.svcCtx.Repo.UserClient.UserGetContactIdList(c.ctx, &userpb.TLUserGetContactIdList{
		UserId: c.MD.UserId,
	})
	if err != nil {
		c.Logger.Errorf("contacts.search - user.getContactIdList error: %v", err)
		return found, nil
	}

	var contactIDs []int64
	if contacts != nil {
		contactIDs = contacts.Datas
	}

	excluded := append([]int64{}, contactIDs...)
	excluded = append(excluded, c.MD.UserId)
	userIDs := make([]int64, 0, limit)
	seen := make(map[int64]struct{})

	usernameMatches, err := c.svcCtx.Repo.UserClient.UserSearchUsername(c.ctx, &userpb.TLUserSearchUsername{
		Q:                q,
		ExcludedContacts: excluded,
		Limit:            limit,
	})
	if err != nil {
		c.Logger.Errorf("contacts.search - user.searchUsername error: q: %q, err: %v", q, err)
		return found, nil
	}
	var usernameDatas []userpb.UsernameDataClazz
	if usernameMatches != nil {
		usernameDatas = usernameMatches.Datas
	}
	for _, username := range usernameDatas {
		if username == nil {
			continue
		}
		peer := tg.FromPeer(username.Peer)
		if peer.PeerType != tg.PEER_USER {
			continue
		}
		if _, ok := seen[peer.PeerId]; ok {
			continue
		}
		seen[peer.PeerId] = struct{}{}
		userIDs = append(userIDs, peer.PeerId)
	}

	searchMatches, err := c.svcCtx.Repo.UserClient.UserSearch(c.ctx, &userpb.TLUserSearch{
		Q:                in.Q,
		ExcludedContacts: excluded,
		Offset:           0,
		Limit:            limit,
	})
	if err == nil && searchMatches != nil {
		if idFound, ok := searchMatches.ToUsersIdFound(); ok {
			for _, id := range idFound.IdList {
				if _, ok := seen[id]; ok {
					continue
				}
				seen[id] = struct{}{}
				userIDs = append(userIDs, id)
			}
		}
	} else {
		c.Logger.Errorf("contacts.search - user.search error: q: %q, err: %v", in.Q, err)
	}

	users, err := c.projectUsers(userIDs, userprojection.MissingStoredReference)
	if err != nil {
		c.Logger.Errorf("contacts.search - user.getUserProjectionBundle error: ids: %v, err: %v", userIDs, err)
		return found, nil
	}

	found.Users = users
	contactSet := make(map[int64]struct{}, len(contactIDs))
	for _, id := range contactIDs {
		contactSet[id] = struct{}{}
	}
	for _, id := range userIDs {
		peer := tg.MakePeerUser(id)
		if _, ok := contactSet[id]; ok {
			found.MyResults = append(found.MyResults, peer)
		} else {
			found.Results = append(found.Results, peer)
		}
	}

	return found, nil
}
