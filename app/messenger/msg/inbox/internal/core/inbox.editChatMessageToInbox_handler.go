/*
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2021-present,  Teamgram Studio (https://teamgram.io).
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package core

import (
	"google.golang.org/protobuf/proto"

	"github.com/teamgram/proto/mtproto"
	"github.com/teamgram/teamgram-server/app/messenger/msg/inbox/inbox"
	"github.com/teamgram/teamgram-server/app/messenger/sync/sync"
	chatpb "github.com/teamgram/teamgram-server/app/service/biz/chat/chat"
	userpb "github.com/teamgram/teamgram-server/app/service/biz/user/user"
)

// InboxEditChatMessageToInbox
// inbox.editChatMessageToInbox from_id:long peer_chat_id:long message:Message = Void;
func (c *InboxCore) InboxEditChatMessageToInbox(in *inbox.TLInboxEditChatMessageToInbox) (*mtproto.Void, error) {
	chatUserIdList, err := c.svcCtx.ChatClient.ChatGetChatParticipantIdList(c.ctx, &chatpb.TLChatGetChatParticipantIdList{
		ChatId: in.PeerChatId,
	})
	if err != nil {
		c.Logger.Errorf("inbox.editChatMessageToInbox - error: %v", err)
		return nil, err
	} else if len(chatUserIdList.Datas) == 0 {
		return mtproto.EmptyVoid, nil
	}

	for _, toId := range chatUserIdList.Datas {
		message := proto.Clone(in.Message).(*mtproto.Message)
		if toId == in.FromId {
			continue
		}

		inBox, err := c.svcCtx.Dao.EditUserInboxMessage(c.ctx, in.FromId, toId, message)
		if err != nil {
			c.Logger.Errorf("inbox.editChatMessageToInbox - error: %v", err)
			// return err
			continue
		} else if inBox == nil {
			c.Logger.Errorf("inbox.editChatMessageToInbox - error: {inBox is nil}")
			continue
		}

		pushUpdates := mtproto.MakePushUpdates(
			func(idList []int64) []*mtproto.User {
				users, _ := c.svcCtx.Dao.UserClient.UserGetMutableUsers(c.ctx,
					&userpb.TLUserGetMutableUsers{
						Id: idList,
					})
				return users.GetUserListByIdList(toId, idList...)
			},
			func(idList []int64) []*mtproto.Chat {
				chats, _ := c.svcCtx.Dao.ChatClient.ChatGetChatListByIdList(c.ctx,
					&chatpb.TLChatGetChatListByIdList{
						IdList: idList,
					})
				return chats.GetChatListByIdList(toId, idList...)
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
			UserId:  toId,
			Updates: pushUpdates,
		})
	}

	return mtproto.EmptyVoid, nil
}
