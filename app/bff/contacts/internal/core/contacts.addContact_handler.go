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

	userprojection "github.com/teamgram/teamgram-server/v2/app/bff/internal/userprojection"
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

	if _, err := c.svcCtx.Repo.UserClient.UserAddContact(c.ctx, &userpb.TLUserAddContact{
		UserId:                   c.MD.UserId,
		AddPhonePrivacyException: tg.ToBoolClazz(in.AddPhonePrivacyException),
		Id:                       id.PeerId,
		FirstName:                in.FirstName,
		LastName:                 in.LastName,
		Phone:                    in.Phone,
	}); err != nil {
		return nil, tg.ErrContactIdInvalid
	}

	users, err := c.projectUsers([]int64{c.MD.UserId, id.PeerId}, userprojection.MissingExplicitInput)
	if err != nil {
		if errors.Is(err, tg.ErrUserIdInvalid) {
			return nil, tg.ErrContactIdInvalid
		}
		return nil, err
	}

	return tg.MakeTLUpdates(&tg.TLUpdates{
		Updates: []tg.UpdateClazz{
			tg.MakeTLUpdatePeerSettings(&tg.TLUpdatePeerSettings{
				Peer:     tg.ToPeerByTypeAndID(id.PeerType, id.PeerId),
				Settings: makePeerSettings(),
			}),
		},
		Users: users,
		Chats: []tg.ChatClazz{},
	}).ToUpdates(), nil
}
