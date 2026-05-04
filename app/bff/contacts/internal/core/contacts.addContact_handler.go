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

// ContactsAddContact
// contacts.addContact#d9ba2e54 flags:# add_phone_privacy_exception:flags.0?true id:InputUser first_name:string last_name:string phone:string note:flags.1?TextWithEntities = Updates;
func (c *ContactsCore) ContactsAddContact(in *tg.TLContactsAddContact) (*tg.Updates, error) {
	if in.FirstName == "" && in.LastName == "" {
		return nil, tg.ErrContactNameEmpty
	}

	id := tg.FromInputUser(c.MD.UserId, in.Id)
	if id.PeerType != tg.PEER_USER || id.PeerId == c.MD.UserId {
		return nil, tg.ErrContactIdInvalid
	}

	users, err := c.svcCtx.Repo.UserClient.UserGetMutableUsersV2(c.ctx, &userpb.TLUserGetMutableUsersV2{
		Id:      []int64{c.MD.UserId, id.PeerId},
		Privacy: true,
		HasTo:   true,
		To:      []int64{id.PeerId},
	})
	if err != nil {
		return nil, tg.ErrContactIdInvalid
	}

	var immutableUsers []tg.ImmutableUserClazz
	if users != nil {
		immutableUsers = users.Users
	}

	contactUser := immutableUserByID(immutableUsers, id.PeerId)
	if contactUser == nil {
		return nil, tg.ErrContactIdInvalid
	}

	changeMutual, err := c.svcCtx.Repo.UserClient.UserAddContact(c.ctx, &userpb.TLUserAddContact{
		UserId:                   c.MD.UserId,
		AddPhonePrivacyException: tg.ToBoolClazz(in.AddPhonePrivacyException),
		Id:                       id.PeerId,
		FirstName:                in.FirstName,
		LastName:                 in.LastName,
		Phone:                    in.Phone,
	})
	if err != nil {
		return nil, tg.ErrContactIdInvalid
	}

	selfUser := projectImmutableUser(immutableUserByID(immutableUsers, c.MD.UserId))
	contact := projectImmutableUser(contactUser)
	if user, ok := contact.(*tg.TLUser); ok {
		user.Contact = true
		user.MutualContact = tg.FromBool(changeMutual)
		user.FirstName = nonEmptyStringPtr(in.FirstName)
		user.LastName = nonEmptyStringPtr(in.LastName)
		user.Phone = nonEmptyStringPtr(in.Phone)
	}

	return tg.MakeTLUpdates(&tg.TLUpdates{
		Updates: []tg.UpdateClazz{
			tg.MakeTLUpdatePeerSettings(&tg.TLUpdatePeerSettings{
				Peer:     tg.ToPeerByTypeAndID(id.PeerType, id.PeerId),
				Settings: makePeerSettings(),
			}),
		},
		Users: []tg.UserClazz{selfUser, contact},
		Chats: []tg.ChatClazz{},
	}).ToUpdates(), nil
}
