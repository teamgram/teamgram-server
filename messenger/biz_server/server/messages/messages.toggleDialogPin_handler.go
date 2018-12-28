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

package messages

import (
	"github.com/golang/glog"
	"github.com/nebula-chat/chatengine/messenger/biz_server/biz/base"
	"github.com/nebula-chat/chatengine/messenger/sync/sync_client"
	"github.com/nebula-chat/chatengine/mtproto"
	"github.com/nebula-chat/chatengine/pkg/grpc_util"
	"github.com/nebula-chat/chatengine/pkg/logger"
	"golang.org/x/net/context"
	"time"
)

// messages.toggleDialogPin#a731e257 flags:# pinned:flags.0?true peer:InputDialogPeer = Bool;
func (s *MessagesServiceImpl) MessagesToggleDialogPin(ctx context.Context, request *mtproto.TLMessagesToggleDialogPin) (*mtproto.Bool, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	glog.Infof("messages.toggleDialogPin#3289be6a - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

	peer := base.FromInputPeer2(md.UserId, request.GetPeer().GetData2().GetPeer())

	if peer.PeerType == base.PEER_EMPTY {
		glog.Error("empty peer")
		return mtproto.ToBool(false), nil
	}

	// TODO(@benqi): check access_hash
	dialogLogic := s.DialogModel.MakeDialogLogic(md.UserId, peer.PeerType, peer.PeerId)
	dialogLogic.ToggleDialogPin(request.GetPinned())

	// peer:DialogPeer

	dialogPeer := &mtproto.TLDialogPeer{Data2: &mtproto.DialogPeer_Data{
		Peer: peer.ToPeer(),
	}}
	// sync other sessions
	updateDialogPinned := &mtproto.TLUpdateDialogPinned{Data2: &mtproto.Update_Data{
		Pinned:  request.GetPinned(),
		Peer_61: dialogPeer.To_DialogPeer(),
	}}
	updates := &mtproto.TLUpdates{Data2: &mtproto.Updates_Data{
		Updates: []*mtproto.Update{updateDialogPinned.To_Update()},
		Users:   []*mtproto.User{},
		Chats:   []*mtproto.Chat{},
		Seq:     0,
		Date:    int32(time.Now().Unix()),
	}}

	switch peer.PeerType {
	case base.PEER_USER:
		updates.Data2.Users = []*mtproto.User{s.UserModel.GetUserById(md.UserId, peer.PeerId)}
	case base.PEER_CHAT:
		updates.Data2.Chats = []*mtproto.Chat{s.ChatModel.GetChatBySelfID(md.UserId, peer.PeerId)}
	case base.PEER_CHANNEL:
	default:
		// TODO(@benqi): log
	}

	sync_client.GetSyncClient().SyncUpdatesNotMe(md.UserId, md.AuthId, updates.To_Updates())

	glog.Info("messages.toggleDialogPin#a731e257 - reply {true}")
	return mtproto.ToBool(true), nil
}
