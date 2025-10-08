// Copyright 2022 Teamgram Authors
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
	"context"

	"github.com/teamgram/marmota/pkg/threading2"
	"github.com/teamgram/proto/mtproto"
	"github.com/teamgram/teamgram-server/app/messenger/sync/sync"
	chatpb "github.com/teamgram/teamgram-server/app/service/biz/chat/chat"
	"github.com/teamgram/teamgram-server/app/service/biz/dialog/dialog"
	userpb "github.com/teamgram/teamgram-server/app/service/biz/user/user"
)

// MessagesReorderPinnedDialogs
// messages.reorderPinnedDialogs#3b1adf37 flags:# force:flags.0?true folder_id:int order:Vector<InputDialogPeer> = Bool;
func (c *DialogsCore) MessagesReorderPinnedDialogs(in *mtproto.TLMessagesReorderPinnedDialogs) (*mtproto.Bool, error) {
	var (
		peerDialogIdList []int64
		peerDialogList   = make([]*mtproto.DialogPeer, 0, len(in.GetOrder())+1)
		idHelper         = mtproto.NewIDListHelper(c.MD.UserId)
	)

	for _, peer := range in.GetOrder() {
		switch peer.PredicateName {
		case mtproto.Predicate_inputDialogPeer:
			p := mtproto.FromInputPeer2(c.MD.UserId, peer.Peer)
			peerDialogIdList = append(peerDialogIdList, mtproto.MakePeerDialogId(p.PeerType, p.PeerId))
			peerDialogList = append(peerDialogList, mtproto.MakeTLDialogPeer(&mtproto.DialogPeer{
				Peer: p.ToPeer(),
			}).To_DialogPeer())
			idHelper.PickByPeerUtil(p.PeerType, p.PeerId)
		case mtproto.Predicate_inputDialogPeerFolder:
			c.Logger.Info("messages.reorderPinnedDialogs - inputDialogPeerFolder %s", peer)
		default:
			err := mtproto.ErrPeerIdInvalid
			c.Logger.Errorf("messages.reorderPinnedDialogs - error: %v", err)
			return nil, err
		}
	}

	_, err := c.svcCtx.Dao.DialogClient.DialogReorderPinnedDialogs(c.ctx, &dialog.TLDialogReorderPinnedDialogs{
		UserId:   c.MD.UserId,
		Force:    mtproto.ToBool(in.Force),
		FolderId: in.FolderId,
		IdList:   peerDialogIdList,
	})
	if err != nil {
		c.Logger.Errorf("messages.reorderPinnedDialogs - error: %v", err)
		return nil, err
	}

	return threading2.WrapperGoFunc(
		c.ctx,
		mtproto.BoolTrue,
		func(ctx context.Context) {
			syncUpdates := mtproto.MakeUpdatesByUpdates(mtproto.MakeTLUpdatePinnedDialogs(&mtproto.Update{
				FolderId:                   mtproto.MakeFlagsInt32(in.FolderId),
				Order_FLAGVECTORDIALOGPEER: peerDialogList,
			}).To_Update())

			if len(peerDialogList) > 0 {
				idHelper.Visit(
					func(userIdList []int64) {
						users, _ := c.svcCtx.Dao.UserClient.UserGetMutableUsers(ctx,
							&userpb.TLUserGetMutableUsers{
								Id: userIdList,
							})
						syncUpdates.PushUser(users.GetUserListByIdList(c.MD.UserId, userIdList...)...)
					},
					func(chatIdList []int64) {
						chats, _ := c.svcCtx.Dao.ChatClient.ChatGetChatListByIdList(ctx,
							&chatpb.TLChatGetChatListByIdList{
								IdList: chatIdList,
							})
						syncUpdates.PushChat(chats.GetChatListByIdList(c.MD.UserId, chatIdList...)...)
					},
					func(channelIdList []int64) {
						// TODO
					})
			}

			_, _ = c.svcCtx.Dao.SyncClient.SyncUpdatesNotMe(ctx, &sync.TLSyncUpdatesNotMe{
				UserId:        c.MD.UserId,
				PermAuthKeyId: c.MD.PermAuthKeyId,
				Updates:       syncUpdates,
			})
		}).(*mtproto.Bool), nil
}
