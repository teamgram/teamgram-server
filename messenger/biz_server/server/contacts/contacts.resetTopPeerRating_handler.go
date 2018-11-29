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

// contacts.resetTopPeerRating#1ae373ac category:TopPeerCategory peer:InputPeer = Bool;
func (s *ContactsServiceImpl) ContactsResetTopPeerRating(ctx context.Context, request *mtproto.TLContactsResetTopPeerRating) (*mtproto.Bool, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	glog.Infof("ContactsResetTopPeerRating - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

	//// TODO(@benqi): Impl ContactsResetTopPeerRating logic
	//_ = helper.FromInputPeer(request.Peer)
	//
	//// TODO(@benqi): 看看客户端代码，什么情况会调用
	//switch request.GetCategory().GetPayload().(type) {
	//case *mtproto.TopPeerCategory_TopPeerCategoryBotsPM:
	//case *mtproto.TopPeerCategory_TopPeerCategoryBotsInline:
	//case *mtproto.TopPeerCategory_TopPeerCategoryCorrespondents:
	//case *mtproto.TopPeerCategory_TopPeerCategoryGroups:
	//case *mtproto.TopPeerCategory_TopPeerCategoryChannels:
	//case *mtproto.TopPeerCategory_TopPeerCategoryPhoneCalls:
	//}

	glog.Infof("ContactsResetTopPeerRating - reply: {true}")
	return mtproto.ToBool(true), nil
}
