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
	"github.com/teamgram/proto/mtproto"
	"github.com/teamgram/teamgram-server/app/messenger/msg/inbox/inbox"
	"github.com/teamgram/teamgram-server/app/messenger/sync/sync"
	chatpb "github.com/teamgram/teamgram-server/app/service/biz/chat/chat"
	userpb "github.com/teamgram/teamgram-server/app/service/biz/user/user"
)

// InboxEditUserMessageToInbox
// inbox.editUserMessageToInbox from_id:long peer_user_id:long message:Message = Void;
func (c *InboxCore) InboxEditUserMessageToInbox(in *inbox.TLInboxEditUserMessageToInbox) (*mtproto.Void, error) {
	inBox, err := c.svcCtx.Dao.EditUserInboxMessage(c.ctx, in.FromId, in.PeerUserId, in.Message)
	if err != nil {
		c.Logger.Errorf("editUserInboxMessage - error: %v", err)
		return nil, err
	} else if inBox == nil {
		err = mtproto.ErrMsgIdInvalid
		c.Logger.Errorf("editUserInboxMessage - error: %v", err)
		return nil, err
	}

	pushUpdates := mtproto.MakePushUpdates(
		func(idList []int64) []*mtproto.User {
			users, _ := c.svcCtx.Dao.UserClient.UserGetMutableUsers(c.ctx,
				&userpb.TLUserGetMutableUsers{
					Id: idList,
				})
			return users.GetUserListByIdList(in.PeerUserId, idList...)
		},
		func(idList []int64) []*mtproto.Chat {
			chats, _ := c.svcCtx.Dao.ChatClient.ChatGetChatListByIdList(c.ctx,
				&chatpb.TLChatGetChatListByIdList{
					IdList: idList,
				})
			return chats.GetChatListByIdList(in.PeerUserId, idList...)
		},
		func(idList []int64) []*mtproto.Chat {
			// TODO
			return nil
		},
		mtproto.MakeTLUpdateEditMessage(&mtproto.Update{
			Pts_INT32:       inBox.Pts,
			PtsCount:        inBox.PtsCount,
			Message_MESSAGE: inBox.Message,
		}).To_Update())

	c.svcCtx.Dao.SyncClient.SyncPushUpdates(c.ctx, &sync.TLSyncPushUpdates{
		UserId:  in.PeerUserId,
		Updates: pushUpdates,
	})

	return mtproto.EmptyVoid, nil
}
