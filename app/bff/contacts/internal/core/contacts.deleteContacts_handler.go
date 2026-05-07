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

// ContactsDeleteContacts
// contacts.deleteContacts#96a0e00 id:Vector<InputUser> = Updates;
func (c *ContactsCore) ContactsDeleteContacts(in *tg.TLContactsDeleteContacts) (*tg.Updates, error) {
	userIDs := make([]int64, 0, len(in.Id))
	for _, id := range in.Id {
		peerID := tg.FromInputUser(c.MD.UserId, id)
		if peerID.PeerType != tg.PEER_USER || peerID.PeerId == c.MD.UserId {
			c.Logger.Errorf("contacts.deleteContacts - error: invalid id %v", id)
			continue
		}
		userIDs = append(userIDs, peerID.PeerId)
	}

	updates := make([]tg.UpdateClazz, 0, len(userIDs))
	for _, id := range userIDs {
		if _, err := c.svcCtx.Repo.UserClient.UserDeleteContact(c.ctx, &userpb.TLUserDeleteContact{
			UserId: c.MD.UserId,
			Id:     id,
		}); err != nil {
			return nil, err
		}

		updates = append(updates, tg.MakeTLUpdatePeerSettings(&tg.TLUpdatePeerSettings{
			Peer:     tg.MakePeerUser(id),
			Settings: makePeerSettings(),
		}))
	}
	usersOut, err := c.projectUsers(userIDs, userprojection.MissingStoredReference)
	if err != nil {
		return nil, err
	}

	rUpdates := tg.MakeTLUpdates(&tg.TLUpdates{
		Updates: updates,
		Users:   usersOut,
		Chats:   []tg.ChatClazz{},
		Date:    0,
		Seq:     0,
	}).ToUpdates()
	// TODO: master sends these updates through sync.SyncUpdatesNotMe here.
	return rUpdates, nil
}
