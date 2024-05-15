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

// ContactsDeleteContacts
// contacts.deleteContacts#96a0e00 id:Vector<InputUser> = Updates;
func (c *ContactsCore) ContactsDeleteContacts(in *mtproto.TLContactsDeleteContacts) (*mtproto.Updates, error) {
	var (
		rUpdates = mtproto.MakeEmptyUpdates()
		idHelper = mtproto.NewIDListHelper(c.MD.UserId)
	)

	for _, id := range in.GetId() {
		switch id.GetPredicateName() {
		case mtproto.Predicate_inputUser:
		default:
			// TODO:
			c.Logger.Errorf("contacts.deleteContacts - error: invalid id %v", id)
			continue
		}
		idHelper.AppendUsers(id.GetUserId())
	}

	deleteUsers, err := c.svcCtx.Dao.UserGetMutableUsers(c.ctx, &userpb.TLUserGetMutableUsers{
		Id: idHelper.UserIdList,
	})
	if err != nil {
		c.Logger.Errorf("contacts.deleteContacts - error: %v", err)
		return nil, err
	}

	deleteUsers.VisitByMe(c.MD.UserId, func(me, it *mtproto.ImmutableUser) {
		if me.Id() != it.Id() {
			// TODO: mutual
			c.svcCtx.Dao.UserClient.UserDeleteContact(c.ctx, &userpb.TLUserDeleteContact{
				UserId: c.MD.UserId,
				Id:     it.Id(),
			})

			rUpdates.PushBackUpdate(
				mtproto.MakeTLUpdatePeerSettings(&mtproto.Update{
					Peer_PEER: mtproto.MakePeerUser(it.Id()),
					Settings: mtproto.MakeTLPeerSettings(&mtproto.PeerSettings{
						ReportSpam:            false,
						AddContact:            false,
						BlockContact:          false,
						ShareContact:          false,
						NeedContactsException: false,
						ReportGeo:             false,
						Autoarchived:          false,
						GeoDistance:           nil,
					}).To_PeerSettings(),
				}).To_Update())

			cUser := it.ToUnsafeUser(me)
			cUser.Contact = false
			cUser.MutualContact = false
			rUpdates.PushUser(cUser)
		}
	})

	c.svcCtx.Dao.SyncClient.SyncUpdatesNotMe(c.ctx, &sync.TLSyncUpdatesNotMe{
		UserId:        c.MD.UserId,
		PermAuthKeyId: c.MD.PermAuthKeyId,
		Updates:       rUpdates,
	})

	return rUpdates, nil
}
