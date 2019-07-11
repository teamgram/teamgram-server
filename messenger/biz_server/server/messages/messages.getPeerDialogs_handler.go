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
)

// messages.getPeerDialogs#e470bcfd peers:Vector<InputDialogPeer> = messages.PeerDialogs;
func (s *MessagesServiceImpl) MessagesGetPeerDialogs(ctx context.Context, request *mtproto.TLMessagesGetPeerDialogs) (*mtproto.Messages_PeerDialogs, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	glog.Infof("messages.getPeerDialogs#e470bcfd - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

	peers := make([]*base.PeerUtil, 0, len(request.GetPeers()))
	for _, p := range request.GetPeers() {
		peer := base.FromInputPeer2(md.UserId, p.To_InputDialogPeer().GetPeer())
		peers = append(peers, peer)
	}
	dialogs := s.DialogModel.GetPeersDialogs(md.UserId, peers)
	dialogItems := s.DialogModel.PickAllIDListByDialogs(dialogs)
	glog.Info(dialogItems)
	messages := s.MessageModel.GetUserMessagesByMessageIdList(md.UserId, dialogItems.MessageIdList)

	//// TODO(@benqi): add channel message.
	//for k, v := range dialogItems.ChannelMessageIdMap {
	//	m := s.MessageModel.GetChannelMessage(md.UserId, k, v)
	//	if m != nil {
	//		messages = append(messages, m)
	//	}
	//}

	users := s.UserModel.GetUserListByIdList(md.UserId, dialogItems.UserIdList)
	chats := s.ChatModel.GetChatListBySelfAndIDList(md.UserId, dialogItems.ChatIdList)
	// chats = append(chats, s.ChannelModel.GetChannelListBySelfAndIDList(md.UserId, dialogItems.ChannelIdList)...)
	state, _ := sync_client.GetSyncClient().SyncGetState(md.AuthId, md.UserId)

	if len(dialogs) == 0 {
		notifySettings := &mtproto.TLPeerNotifySettings{Data2: &mtproto.PeerNotifySettings_Data{
			ShowPreviews: mtproto.ToBool(true),
			Silent:       mtproto.ToBool(false),
			MuteUntil:    1,
			Sound:        "default",
		}}

		dialog := &mtproto.TLDialog{Data2: &mtproto.Dialog_Data{
			Peer:           peers[0].ToPeer(),
			TopMessage:     0,
			NotifySettings: notifySettings.To_PeerNotifySettings(),
		}}

		dialogs = append(dialogs, dialog.To_Dialog())
	}

	peerDialogs := &mtproto.TLMessagesPeerDialogs{Data2: &mtproto.Messages_PeerDialogs_Data{
		Dialogs:  dialogs,
		Messages: messages,
		Users:    users,
		Chats:    chats,
		State:    state,
	}}

	glog.Infof("messages.getPeerDialogs#e470bcfd - reply: %s", logger.JsonDebugData(peerDialogs))
	return peerDialogs.To_Messages_PeerDialogs(), nil
}
