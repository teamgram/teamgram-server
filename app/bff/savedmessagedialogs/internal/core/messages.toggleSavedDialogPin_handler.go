// Copyright 2024 Teamgram Authors
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
	chatpb "github.com/teamgram/teamgram-server/app/service/biz/chat/chat"
	"github.com/teamgram/teamgram-server/app/service/biz/dialog/dialog"
	userpb "github.com/teamgram/teamgram-server/app/service/biz/user/user"
)

// MessagesToggleSavedDialogPin
// messages.toggleSavedDialogPin#ac81bbde flags:# pinned:flags.0?true peer:InputDialogPeer = Bool;
func (c *SavedMessageDialogsCore) MessagesToggleSavedDialogPin(in *mtproto.TLMessagesToggleSavedDialogPin) (*mtproto.Bool, error) {
	var (
		peer *mtproto.PeerUtil
	)

	switch in.GetPeer().GetPredicateName() {
	case mtproto.Predicate_inputDialogPeer:
		peer = mtproto.FromInputPeer2(c.MD.UserId, in.GetPeer().GetPeer())
	case mtproto.Predicate_inputDialogPeerFolder:
		// error
		c.Logger.Errorf("messages.toggleSavedDialogPin - error: client not send inputDialogPeerFolder: %v", in.GetPeer())
		return mtproto.BoolFalse, nil
	default:
		err := mtproto.ErrInputConstructorInvalid
		c.Logger.Errorf("messages.toggleSavedDialogPin - error: %v", err)
		return nil, err
	}

	_, err := c.svcCtx.Dao.DialogClient.DialogToggleSavedDialogPin(c.ctx, &dialog.TLDialogToggleSavedDialogPin{
		UserId: c.MD.UserId,
		Peer:   peer,
		Pinned: mtproto.ToBool(in.Pinned),
	})
	if err != nil {
		c.Logger.Errorf("messages.toggleSavedDialogPin - error: %v", err)
		return mtproto.BoolFalse, nil
	}

	var (
		idHelper    = mtproto.NewIDListHelper(c.MD.UserId)
		syncUpdates = mtproto.MakeUpdatesByUpdates(mtproto.MakeTLUpdateSavedDialogPinned(&mtproto.Update{
			Pinned: in.GetPinned(),
			Peer_DIALOGPEER: mtproto.MakeTLDialogPeer(&mtproto.DialogPeer{
				Peer: peer.ToPeer(),
			}).To_DialogPeer(),
		}).To_Update())
	)

	idHelper.PickByPeerUtil(peer.PeerType, peer.PeerId)
	idHelper.Visit(
		func(userIdList []int64) {
			users, _ := c.svcCtx.Dao.UserClient.UserGetMutableUsers(c.ctx,
				&userpb.TLUserGetMutableUsers{
					Id: userIdList,
				})
			syncUpdates.PushUser(users.GetUserListByIdList(c.MD.UserId, userIdList...)...)
		},
		func(chatIdList []int64) {
			chats, _ := c.svcCtx.Dao.ChatClient.ChatGetChatListByIdList(c.ctx,
				&chatpb.TLChatGetChatListByIdList{
					IdList: chatIdList,
				})
			syncUpdates.PushChat(chats.GetChatListByIdList(c.MD.UserId, chatIdList...)...)
		},
		func(channelIdList []int64) {
			// TODO
		})
	c.svcCtx.Dao.SyncClient.SyncUpdatesNotMe(c.ctx, &sync.TLSyncUpdatesNotMe{
		UserId:        c.MD.UserId,
		PermAuthKeyId: c.MD.PermAuthKeyId,
		Updates:       syncUpdates,
	})

	return mtproto.BoolTrue, nil
}
