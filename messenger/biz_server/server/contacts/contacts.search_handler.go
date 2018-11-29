// Copyright (c) 2018-present,  NebulaChat Studio (https://nebula.chat).
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

// Author: Benqi (wubenqi@gmail.com)

package contacts

import (
	"github.com/golang/glog"
	"github.com/nebula-chat/chatengine/pkg/grpc_util"
	"github.com/nebula-chat/chatengine/pkg/logger"
	"github.com/nebula-chat/chatengine/mtproto"
	"golang.org/x/net/context"
)

// contacts.search#11f812d8 q:string limit:int = contacts.Found;
func (s *ContactsServiceImpl) ContactsSearch(ctx context.Context, request *mtproto.TLContactsSearch) (*mtproto.Contacts_Found, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	glog.Infof("contacts.search#11f812d8 - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

	found := &mtproto.TLContactsFound{Data2: &mtproto.Contacts_Found_Data{
		MyResults: []*mtproto.Peer{},
		Results:   []*mtproto.Peer{},
		// Chats:     []*mtproto.Chat{},
	}}

	q := request.Q
	if len(q) >= 5 {
		if q[0] == '@' {
			q = q[1:]
		}
	}

	// Check query string and limit
	if len(q) >= 5 && request.Limit > 0 {
		contactLogic := s.ContactModel.MakeContactLogic(md.UserId)
		userIdList, channelIdList := contactLogic.SearchContacts(q, request.Limit)

		// TODO(@benqi): impl channelIdList
		_ = channelIdList

		userList := s.UserModel.GetUserListByIdList(md.UserId, userIdList)
		found.Data2.Users = userList
		for _, u := range userList {
			peer := &mtproto.TLPeerUser{Data2: &mtproto.Peer_Data{
				UserId: u.GetData2().GetId(),
			}}
			if u.GetData2().GetContact() {
				found.Data2.MyResults = append(found.Data2.MyResults, peer.To_Peer())
			} else {
				found.Data2.Results = append(found.Data2.Results, peer.To_Peer())
			}
		}
	} else {
		found.Data2.MyResults = []*mtproto.Peer{}
		found.Data2.Results = []*mtproto.Peer{}
		found.Data2.Users = []*mtproto.User{}
		found.Data2.Chats = []*mtproto.Chat{}
	}

	glog.Infof("contacts.search#11f812d8 - reply: %s", logger.JsonDebugData(found))
	return found.To_Contacts_Found(), nil
}
