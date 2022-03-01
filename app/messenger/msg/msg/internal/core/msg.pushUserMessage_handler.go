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
	"github.com/teamgram/proto/mtproto"
	"github.com/teamgram/teamgram-server/app/messenger/msg/inbox/inbox"
	"github.com/teamgram/teamgram-server/app/messenger/msg/msg/msg"
	"github.com/teamgram/teamgram-server/app/messenger/sync/sync"
	userpb "github.com/teamgram/teamgram-server/app/service/biz/user/user"
)

// MsgPushUserMessage
// msg.pushUserMessage user_id:long auth_key_id:long (UserMessage) = Bool;
func (c *MsgCore) MsgPushUserMessage(in *msg.TLMsgPushUserMessage) (*mtproto.Bool, error) {
	var (
		peer   = mtproto.MakePeerUtil(in.PeerType, in.PeerId)
		boxMsg = in.Message
	)

	if !peer.IsUser() {
		c.Logger.Errorf("peer must is user")
		return mtproto.BoolFalse, nil
	}

	sendMe := in.UserId == in.PeerId
	if !sendMe {
		// TODO(@benqi)
		// 1. check blocked
		// 2. span
	}

	// TODO(@benqi): r.From.Type
	switch in.PushType {
	case 0:
		err := c.pushUserMessage(
			in.UserId,
			peer.PeerId,
			boxMsg,
			func(did int64, inboxMsg *mtproto.Message) error {
				c.svcCtx.Dao.InboxClient.InboxSendUserMessageToInbox(
					c.ctx,
					&inbox.TLInboxSendUserMessageToInbox{
						FromId:     in.UserId,
						PeerUserId: peer.PeerId,
						Message: inbox.MakeTLInboxMessageData(&inbox.InboxMessageData{
							RandomId:        boxMsg.RandomId,
							DialogMessageId: did,
							Message:         inboxMsg,
						}).To_InboxMessageData(),
					})
				//
				return nil
			})
		if err != nil {
			return nil, err
		}
	default:
		c.svcCtx.Dao.InboxClient.InboxSendUserMessageToInbox(
			c.ctx,
			&inbox.TLInboxSendUserMessageToInbox{
				FromId:     in.UserId,
				PeerUserId: peer.PeerId,
				Message: inbox.MakeTLInboxMessageData(&inbox.InboxMessageData{
					RandomId:        boxMsg.RandomId,
					DialogMessageId: 0,
					Message:         in.Message.Message,
				}).To_InboxMessageData(),
			})
	}

	return mtproto.BoolTrue, nil
}

func (c *MsgCore) pushUserMessage(
	fromUserId int64,
	toUserId int64,
	outBox *msg.OutboxMessage,
	cb func(did int64, inboxMsg *mtproto.Message) error) error {

	sendMe := fromUserId == toUserId
	if !sendMe {
		// TODO(@benqi)
		// 1. check blocked
		// 2. span
	}

	// handle duplicateMessage
	hasDuplicateMessage, err := c.svcCtx.Dao.HasDuplicateMessage(c.ctx, fromUserId, outBox.RandomId)
	if err != nil {
		c.Logger.Errorf("checkDuplicateMessage error - %v", err)
		return err
	} else if hasDuplicateMessage {
		upd, err := c.svcCtx.Dao.GetDuplicateMessage(c.ctx, fromUserId, outBox.RandomId)
		if err != nil {
			c.Logger.Errorf("checkDuplicateMessage error - %v", err)
			return err
		} else if upd != nil {
			return nil
		}
	}

	box, err := c.svcCtx.Dao.SendUserMessage(c.ctx, fromUserId, toUserId, outBox)
	if err != nil {
		c.Logger.Error(err.Error())
		return err
	}

	if !hasDuplicateMessage && cb != nil {
		err = cb(box.DialogMessageId, box.ToMessage(fromUserId))
		if err != nil {
			c.Logger.Error(err.Error())
			return err
		}
	}

	updateNewMessage := mtproto.MakeTLUpdateNewMessage(&mtproto.Update{
		Pts_INT32:       box.Pts,
		PtsCount:        box.PtsCount,
		RandomId:        box.RandomId,
		Message_MESSAGE: box.Message,
	}).To_Update()

	c.svcCtx.Dao.SyncClient.SyncPushUpdates(c.ctx, &sync.TLSyncPushUpdates{
		UserId: fromUserId,
		Updates: mtproto.MakePushUpdates(
			func(idList []int64) []*mtproto.User {
				users, _ := c.svcCtx.Dao.UserClient.UserGetMutableUsers(c.ctx,
					&userpb.TLUserGetMutableUsers{
						Id: idList,
					})
				return users.GetUserListByIdList(fromUserId, idList...)
			},
			func(idList []int64) []*mtproto.Chat {
				//chats, _ := c.svcCtx.Dao.ChatClient.ChatGetChatListByIdList(c.ctx,
				//	&chatpb.TLChatGetChatListByIdList{
				//		IdList: idList,
				//	})
				//return chats.GetChatListByIdList(fromUserId, idList...)
				return nil
			},
			func(idList []int64) []*mtproto.Chat {
				// TODO
				return nil
			},
			updateNewMessage),
	})

	return err
}
