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
	"errors"
	"fmt"

	userpb "github.com/teamgram/teamgram-server/v2/app/service/biz/user/user"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

// ContactsResolvePhone
// contacts.resolvePhone#8af94344 phone:string = contacts.ResolvedPeer;
func (c *UsersCore) ContactsResolvePhone(in *tg.TLContactsResolvePhone) (*tg.ContactsResolvedPeer, error) {
	selfID, err := requireSelfID(c)
	if err != nil {
		return nil, err
	}
	if in == nil || in.Phone == "" {
		return nil, tg.ErrInputRequestInvalid
	}
	if c.svcCtx == nil || c.svcCtx.Repo == nil || c.svcCtx.Repo.UserClient == nil {
		return nil, fmt.Errorf("contacts.resolvePhone: user client is nil")
	}

	resolvedID, err := c.svcCtx.Repo.UserClient.UserGetUserIdByPhone(c.ctx, &userpb.TLUserGetUserIdByPhone{
		Phone: in.Phone,
	})
	if err != nil {
		if errors.Is(err, userpb.ErrUserNotFound) {
			return nil, tg.ErrPhoneNotOccupied
		}
		return nil, err
	}
	if resolvedID == nil || resolvedID.V <= 0 {
		return nil, tg.ErrPhoneNotOccupied
	}

	mutableUsers, err := c.svcCtx.Repo.UserClient.UserGetMutableUsersV2(c.ctx, &userpb.TLUserGetMutableUsersV2{
		Id:      []int64{resolvedID.V, selfID},
		Privacy: true,
		HasTo:   true,
		To:      []int64{selfID},
	})
	if err != nil {
		return nil, err
	}

	var resolvedUser tg.UserClazz
	if mutableUsers != nil {
		for _, immutableUser := range mutableUsers.Users {
			user := projectImmutableUser(immutableUser)
			if id, ok := userID(user); ok && id == resolvedID.V {
				resolvedUser = user
				break
			}
		}
	}
	if resolvedUser == nil {
		return nil, tg.ErrPhoneNotOccupied
	}

	return tg.MakeTLContactsResolvedPeer(&tg.TLContactsResolvedPeer{
		Peer:  tg.MakeTLPeerUser(&tg.TLPeerUser{UserId: resolvedID.V}),
		Chats: []tg.ChatClazz{},
		Users: []tg.UserClazz{resolvedUser},
	}).ToContactsResolvedPeer(), nil
}
