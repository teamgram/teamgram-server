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

// contacts.getTopPeers#d4982db5 flags:# correspondents:flags.0?true bots_pm:flags.1?true bots_inline:flags.2?true phone_calls:flags.3?true groups:flags.10?true channels:flags.15?true offset:int limit:int hash:int = contacts.TopPeers;
func (s *ContactsServiceImpl) ContactsGetTopPeers(ctx context.Context, request *mtproto.TLContactsGetTopPeers) (*mtproto.Contacts_TopPeers, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	glog.Infof("contacts.getTopPeers#d4982db5 - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

	// TODO(@benqi): Impl ContactsGetTopPeers logic
	topPeers := &mtproto.TLContactsTopPeers{Data2: &mtproto.Contacts_TopPeers_Data{
		Categories: []*mtproto.TopPeerCategoryPeers{},
		Chats:      []*mtproto.Chat{},
		Users:      []*mtproto.User{},
	}}

	glog.Infof("contacts.getTopPeers#d4982db5 - reply: %s", logger.JsonDebugData(topPeers))
	return topPeers.To_Contacts_TopPeers(), nil
}
