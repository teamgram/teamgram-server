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

package account

import (
	"github.com/golang/glog"
	"github.com/nebula-chat/chatengine/pkg/grpc_util"
	"github.com/nebula-chat/chatengine/pkg/logger"
	"github.com/nebula-chat/chatengine/mtproto"
	"github.com/nebula-chat/chatengine/messenger/biz_server/biz/base"
	updates2 "github.com/nebula-chat/chatengine/messenger/biz_server/biz/core/update"
	"golang.org/x/net/context"
	"github.com/nebula-chat/chatengine/messenger/sync/sync_client"
)

/*
 Android client's account.updateNotifySettings#84be5b93 source code:

	TLRPC.TL_account_updateNotifySettings req = new TLRPC.TL_account_updateNotifySettings();
	req.settings = new TLRPC.TL_inputPeerNotifySettings();
	req.settings.sound = "default";
	int mute_type = preferences.getInt("notify2_" + dialog_id, 0);
	if (mute_type == 3) {
		req.settings.mute_until = preferences.getInt("notifyuntil_" + dialog_id, 0);
	} else {
		req.settings.mute_until = mute_type != 2 ? 0 : Integer.MAX_VALUE;
	}
	req.settings.show_previews = preferences.getBoolean("preview_" + dialog_id, true);
	req.settings.silent = preferences.getBoolean("silent_" + dialog_id, false);
	req.peer = new TLRPC.TL_inputNotifyPeer();
	((TLRPC.TL_inputNotifyPeer) req.peer).peer = MessagesController.getInputPeer((int) dialog_id);
*/

// account.updateNotifySettings#84be5b93 peer:InputNotifyPeer settings:InputPeerNotifySettings = Bool;
func (s *AccountServiceImpl) AccountUpdateNotifySettings(ctx context.Context, request *mtproto.TLAccountUpdateNotifySettings) (*mtproto.Bool, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	glog.Infof("account.updateNotifySettings#84be5b93 - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

	// TODO(@benqi): by android client source code, we only impl inputNotifyPeer - (inputNotifyPeer#b8bc5b0c peer:InputPeer = InputNotifyPeer)

	peer := base.FromInputNotifyPeer(md.UserId, request.GetPeer())
	if peer.PeerType == base.PEER_EMPTY {
		glog.Error("peer is empty")
		return mtproto.ToBool(false), nil
	}

	settings := request.GetSettings().To_InputPeerNotifySettings()
	s.AccountModel.SetNotifySettings(md.UserId, peer, settings)

	// sync
	updateNotifySettings := &mtproto.TLUpdateNotifySettings{Data2: &mtproto.Update_Data{
		Peer_28: peer.ToNotifyPeer(),
		NotifySettings: &mtproto.PeerNotifySettings{
			Constructor: mtproto.TLConstructor_CRC32_peerNotifySettings,
			Data2: &mtproto.PeerNotifySettings_Data{
				ShowPreviews: settings.GetShowPreviews(),
				Silent:       settings.GetSilent(),
				MuteUntil:    settings.GetMuteUntil(),
				Sound:        settings.GetSound(),
			},
		},
	}}
	notifySettingUpdates := updates2.NewUpdatesLogic(md.UserId)
	notifySettingUpdates.AddUpdate(updateNotifySettings.To_Update())

	switch peer.PeerType {
	case base.PEER_SELF:
		user := s.UserModel.GetUserById(md.UserId, md.UserId)
		notifySettingUpdates.AddUser(user)
	case base.PEER_USER:
		user := s.UserModel.GetUserById(md.UserId, peer.PeerId)
		notifySettingUpdates.AddUser(user)
	case base.PEER_CHAT:
		chat := s.ChatModel.GetChatBySelfID(md.UserId, peer.PeerId)
		notifySettingUpdates.AddChat(chat)
	case base.PEER_CHANNEL:
		// TODO(@benqi): impl
		glog.Warning("channels.createChannel blocked, License key from https://nebula.chat required to unlock enterprise features.")
	case base.PEER_USERS:
	case base.PEER_CHATS:
	}

	sync_client.GetSyncClient().SyncUpdatesNotMe(md.UserId, md.AuthId, notifySettingUpdates.ToUpdates())

	glog.Infof("account.updateNotifySettings#84be5b93 - reply: {true}")
	return mtproto.ToBool(true), nil
}
