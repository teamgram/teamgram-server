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

// contacts.blocked#1c138d15 blocked:Vector<ContactBlocked> users:Vector<User> = contacts.Blocked;
// contacts.blockedSlice#900802a1 count:int blocked:Vector<ContactBlocked> users:Vector<User> = contacts.Blocked;
//
// contacts.getBlocked#f57c350f offset:int limit:int = contacts.Blocked;
func (s *ContactsServiceImpl) ContactsGetBlocked(ctx context.Context, request *mtproto.TLContactsGetBlocked) (*mtproto.Contacts_Blocked, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	glog.Infof("contacts.getBlocked#f57c350f - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

	contactLogic := s.ContactModel.MakeContactLogic(md.UserId)
	blockedList := contactLogic.GetBlockedList(request.Offset, request.Limit)

	// TODO(@benqi): impl blockedSlice

	contactsBlocked := &mtproto.TLContactsBlocked{Data2: &mtproto.Contacts_Blocked_Data{
		Blocked: blockedList,
	}}
	// .NewTLContactsBlocked()
	if len(blockedList) > 0 {
		blockedIdList := make([]int32, 0, len(blockedList))
		userIdList := make([]int32, 0, len(blockedList))
		for _, c := range blockedList {
			userIdList = append(userIdList, c.GetData2().GetUserId())
		}

		users := s.UserModel.GetUserListByIdList(md.UserId, blockedIdList)
		contactsBlocked.SetUsers(users)
	}

	glog.Infof("contacts.getBlocked#f57c350f - reply: %s\n", logger.JsonDebugData(contactsBlocked))
	return contactsBlocked.To_Contacts_Blocked(), nil
}
