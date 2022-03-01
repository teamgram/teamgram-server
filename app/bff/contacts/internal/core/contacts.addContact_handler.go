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
	userpb "github.com/teamgram/teamgram-server/app/service/biz/user/user"
)

// ContactsAddContact
// contacts.addContact#e8f463d0 flags:# add_phone_privacy_exception:flags.0?true id:InputUser first_name:string last_name:string phone:string = Updates;
func (c *ContactsCore) ContactsAddContact(in *mtproto.TLContactsAddContact) (*mtproto.Updates, error) {
	// 400	CONTACT_NAME_EMPTY	Contact name empty.
	if in.FirstName == "" && in.LastName == "" {
		err := mtproto.ErrContactNameEmpty
		c.Logger.Errorf("contacts.addContact - error: %v", err)
		return nil, err
	}

	// 400	CONTACT_ID_INVALID	The provided contact ID is invalid
	id := mtproto.FromInputUser(c.MD.UserId, in.Id)
	switch id.PeerType {
	case mtproto.PEER_USER:
		//
	default:
		// TODO:
		/*
			Possible errors
			Code	Type	Description
			400	CHANNEL_PRIVATE	You haven't joined this channel/supergroup.
			400	CONTACT_ID_INVALID	The provided contact ID is invalid.
			400	CONTACT_NAME_EMPTY	Contact name empty.
			400	MSG_ID_INVALID	Invalid message ID provided.
		*/
		err := mtproto.ErrContactIdInvalid
		c.Logger.Errorf("contacts.addContact - error: %v", err)
		return nil, err
	}

	users, _ := c.svcCtx.Dao.UserClient.UserGetMutableUsers(c.ctx, &userpb.TLUserGetMutableUsers{
		Id: []int64{c.MD.UserId, id.PeerId},
	})

	if !users.CheckExistUser(id.PeerId) {
		err := mtproto.ErrContactIdInvalid
		c.Logger.Errorf("contacts.addContact - error: %v", err)
		return nil, err
	}

	c.svcCtx.Dao.UserClient.UserAddContact(c.ctx, &userpb.TLUserAddContact{
		UserId:                   c.MD.UserId,
		AddPhonePrivacyException: mtproto.ToBool(in.AddPhonePrivacyException),
		Id:                       id.PeerId,
		FirstName:                in.FirstName,
		LastName:                 in.LastName,
		Phone:                    in.Phone,
	})

	// TODO(@benqi): 性能优化，复用users
	rUpdates := mtproto.MakeUpdatesByUpdatesUsers(
		users.GetUserListByIdList(c.MD.UserId, c.MD.UserId, id.PeerId),
		mtproto.MakeTLUpdatePeerSettings(&mtproto.Update{
			Peer_PEER: id.ToPeer(),
			Settings: mtproto.MakeTLPeerSettings(&mtproto.PeerSettings{
				ReportSpam:            false,
				AddContact:            false,
				BlockContact:          false,
				ShareContact:          false,
				NeedContactsException: false,
				ReportGeo:             false,
			}).To_PeerSettings(),
		}).To_Update())

	return rUpdates, nil
}
