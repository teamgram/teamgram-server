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
	"github.com/nebula-chat/chatengine/messenger/biz_server/biz/base"
)

// contacts.resolveUsername#f93ccba3 username:string = contacts.ResolvedPeer;
func (s *ContactsServiceImpl) ContactsResolveUsername(ctx context.Context, request *mtproto.TLContactsResolveUsername) (*mtproto.Contacts_ResolvedPeer, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	glog.Infof("contacts.resolveUsername#f93ccba3 - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

	peer, err := s.UsernameModel.ResolveUsername(request.GetUsername())
	if err != nil {
		glog.Errorf("contacts.resolveUsername#f93ccba3 - reply: {%v}", err)
		return nil, err
	}

	resolvedPeer := &mtproto.TLContactsResolvedPeer{Data2: &mtproto.Contacts_ResolvedPeer_Data{
		Peer: peer.ToPeer(),
		Chats: []*mtproto.Chat{},
		Users: []*mtproto.User{},
	}}
	// peer.ToPeer()
	if peer.PeerType == base.PEER_USER {
		resolvedPeer.SetUsers([]*mtproto.User{s.UserModel.GetUserById(md.UserId, peer.PeerId)})
	} else if peer.PeerType == base.PEER_CHAT {
		resolvedPeer.SetChats([]*mtproto.Chat{s.ChatModel.GetChatBySelfID(md.UserId, peer.PeerId)})
	} else {
		glog.Warning("blocked, License key from https://nebula.chat required to unlock enterprise features.")
		// resolvedPeer.SetChats([]*mtproto.Chat{s.ChannelModel.GetChannelBySelfID(md.UserId, peer.PeerId)})
	}

	glog.Infof("contacts.resolveUsername#f93ccba3 - reply: {%v}", resolvedPeer)
	return resolvedPeer.To_Contacts_ResolvedPeer(), nil
}
