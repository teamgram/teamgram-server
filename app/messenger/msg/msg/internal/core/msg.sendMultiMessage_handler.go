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
	chatpb "github.com/teamgram/teamgram-server/app/service/biz/chat/chat"
	userpb "github.com/teamgram/teamgram-server/app/service/biz/user/user"
)

// MsgSendMultiMessage
// msg.sendMultiMessage user_id:long auth_key_id:long peer_type:int peer_id:long message:Vector<OutboxMessage> = Updates;
func (c *MsgCore) MsgSendMultiMessage(in *msg.TLMsgSendMultiMessage) (*mtproto.Updates, error) {
	var (
		err      error
		rUpdates *mtproto.Updates
		peer     = mtproto.MakePeerUtil(in.PeerType, in.PeerId)
	)

	if peer.IsChannel() {
		// c.Logger.Errorf("msg.sendMultiMessage blocked, License key from https://teamgram.net required to unlock enterprise features.")
		return nil, mtproto.ErrEnterpriseIsBlocked
	}

	if len(in.Message) == 0 {
		err = mtproto.ErrGroupedMediaInvalid
		c.Logger.Errorf("msg.sendMultiMessage - error: %v", err)
		return nil, err
	}
	if in.Message[0].GetScheduleDate().GetValue() != 0 {
		// c.Logger.Errorf("msg.sendMultiMessage blocked, License key from https://teamgram.net required to unlock enterprise features.")
		return nil, mtproto.ErrEnterpriseIsBlocked
	}

	if !peer.IsChatOrUser() {
		err = mtproto.ErrPeerIdInvalid
		c.Logger.Errorf("msg.sendMultiMessage - error: %v", err)
		return nil, err
	}

	if peer.IsUser() {
		// private
		rUpdates, err = c.sendUserOutgoingMultiMessage(in)
		if err != nil {
			c.Logger.Errorf("msg.sendMultiMessage - error: %v", err)
			return nil, err
		}
	} else {
		// chat
		rUpdates, err = c.sendChatOutgoingMultiMessage(in)
		if err != nil {
			c.Logger.Errorf("msg.sendMultiMessage - error: %v", err)
			return nil, err
		}
	}

	return rUpdates, nil
}

func (c *MsgCore) sendUserOutgoingMultiMessage(in *msg.TLMsgSendMultiMessage) (*mtproto.Updates, error) {
	var (
		err      error
		rUpdates *mtproto.Updates
		users    *userpb.Vector_ImmutableUser
		// users model.MutableUsers
	)
	users, err = c.svcCtx.Dao.UserClient.UserGetMutableUsers(c.ctx, &userpb.TLUserGetMutableUsers{
		Id: []int64{in.UserId, in.PeerId},
	})
	if err != nil {
		c.Logger.Errorf("msg.sendMultiMessage - error: %v")
	}

	// .UserFacade.GetMutableUsers(ctx, r.From.Id, r.PeerId)
	sender, _ := users.GetImmutableUser(in.UserId)
	if sender == nil || sender.Deleted() {
		err = mtproto.ErrInputUserDeactivated
		c.Logger.Errorf("msg.sendMultiMessage - error: %v")
		return nil, err
	}
	// TODO(@benqi): check
	// if sender.Restricted() {
	//	err = mtproto.ErrUserRestricted
	//	return
	// }

	peerUser, _ := users.GetImmutableUser(in.PeerId)
	if peerUser == nil || peerUser.Deleted() {
		err = mtproto.ErrPeerIdInvalid
		c.Logger.Errorf("msg.sendMultiMessage - error: %v")
		return nil, err
	}

	// iOS
	//if s.UserFacade.IsBlockedByUser(ctx, r.PeerId, r.From.Id) {
	//	err = mtproto.ErrYouBlockedUser
	//	log.Errorf("sendUserOutgoingMultiMessage - error: %v", err)
	//	return
	//}

	sendMe := in.UserId == in.PeerId
	if !sendMe {
		// TODO(@benqi)
		// 1. check blocked
		// 2. span
	}

	rUpdates, err = c.sendUserMultiMessage(
		in.UserId,
		in.AuthKeyId,
		in.PeerId,
		in.Message,
		func(inboxMsgList []*mtproto.MessageBox) error {
			var (
				inboxMsgDataList = make([]*inbox.InboxMessageData, 0, len(inboxMsgList))
			)

			for _, box := range inboxMsgList {
				inboxMsgDataList = append(inboxMsgDataList, &inbox.InboxMessageData{
					RandomId:        box.RandomId,
					DialogMessageId: box.DialogMessageId,
					Message:         box.Message,
				})
			}
			c.svcCtx.Dao.InboxClient.InboxSendUserMultiMessageToInbox(c.ctx, &inbox.TLInboxSendUserMultiMessageToInbox{
				FromId:     in.UserId,
				PeerUserId: in.PeerId,
				Message:    inboxMsgDataList,
			})

			return nil
		})

	return rUpdates, nil
}

func (c *MsgCore) sendUserMultiMessage(
	fromUserId int64,
	fromAuthKeyId int64,
	toUserId int64,
	outBoxList []*msg.OutboxMessage,
	cb func(inboxMsgList []*mtproto.MessageBox) error) (*mtproto.Updates, error) {

	sendMe := fromUserId == toUserId
	if !sendMe {
		// TODO(@benqi)
		// 1. check blocked
		// 2. span
	}

	boxList, err := c.svcCtx.Dao.SendUserMultiMessage(c.ctx, fromUserId, toUserId, outBoxList)
	if err != nil {
		c.Logger.Errorf("msg.sendMultiMessage - error: %v")
		return nil, err
	}

	if cb != nil {
		err = cb(boxList)
		if err != nil {
			c.Logger.Errorf("msg.sendMultiMessage - error: %v")
			return nil, err
		}
	}

	var (
		updateNewMessageList = make([]*mtproto.Update, 0, len(boxList))
	)

	for _, outBoxMsg := range boxList {
		updateNewMessageList = append(updateNewMessageList, mtproto.MakeTLUpdateNewMessage(&mtproto.Update{
			Pts_INT32:       outBoxMsg.Pts,
			PtsCount:        outBoxMsg.PtsCount,
			RandomId:        outBoxMsg.RandomId,
			Message_MESSAGE: outBoxMsg.Message,
		}).To_Update())
	}

	rUpdates := mtproto.MakeReplyUpdates(
		func(idList []int64) []*mtproto.User {
			users, _ := c.svcCtx.Dao.UserClient.UserGetMutableUsers(c.ctx,
				&userpb.TLUserGetMutableUsers{
					Id: idList,
				})
			return users.GetUserListByIdList(fromUserId, idList...)
		},
		func(idList []int64) []*mtproto.Chat {
			chats, _ := c.svcCtx.Dao.ChatClient.ChatGetChatListByIdList(c.ctx,
				&chatpb.TLChatGetChatListByIdList{
					IdList: idList,
				})
			return chats.GetChatListByIdList(fromUserId, idList...)
		},
		func(idList []int64) []*mtproto.Chat {
			// TODO
			return nil
		},
		updateNewMessageList...)

	c.svcCtx.Dao.SyncClient.SyncUpdatesNotMe(c.ctx, &sync.TLSyncUpdatesNotMe{
		UserId:    fromUserId,
		AuthKeyId: fromAuthKeyId,
		Updates: mtproto.MakeSyncNotMeUpdates(
			func(idList []int64) []*mtproto.User {
				return rUpdates.Users
			},
			func(idList []int64) []*mtproto.Chat {
				return rUpdates.Chats
			},
			func(idList []int64) []*mtproto.Chat {
				// rUpdates.Chats include chats
				return nil
			},
			updateNewMessageList...),
	})

	return rUpdates, nil
}

func (c *MsgCore) sendChatOutgoingMultiMessage(in *msg.TLMsgSendMultiMessage) (*mtproto.Updates, error) {
	users, _ := c.svcCtx.Dao.UserClient.UserGetMutableUsers(c.ctx, &userpb.TLUserGetMutableUsers{
		Id: []int64{in.UserId},
	})

	sender, _ := users.GetImmutableUser(in.UserId)
	if sender == nil || sender.Deleted() {
		err := mtproto.ErrInputUserDeactivated
		c.Logger.Errorf("msg.sendMultiMessage - error: %v")
		return nil, err
	}

	rUpdates, _ := c.sendChatMultiMessage(
		in.UserId,
		in.AuthKeyId,
		in.PeerId,
		in.Message,
		func(inboxMsgList []*mtproto.MessageBox) error {
			var (
				inboxMsgDataList = make([]*inbox.InboxMessageData, 0, len(inboxMsgList))
			)

			for _, box := range inboxMsgList {
				inboxMsgDataList = append(inboxMsgDataList, &inbox.InboxMessageData{
					RandomId:        box.RandomId,
					DialogMessageId: box.DialogMessageId,
					Message:         box.Message,
				})
			}
			c.svcCtx.Dao.InboxClient.InboxSendChatMultiMessageToInbox(c.ctx, &inbox.TLInboxSendChatMultiMessageToInbox{
				FromId:     in.UserId,
				PeerChatId: in.PeerId,
				Message:    inboxMsgDataList,
			})

			return nil
		})

	return rUpdates, nil
}

func (c *MsgCore) sendChatMultiMessage(
	fromUserId int64,
	fromAuthKeyId int64,
	chatId int64,
	outBoxList []*msg.OutboxMessage,
	cb func(inboxMsgList []*mtproto.MessageBox) error) (*mtproto.Updates, error) {

	boxList, err := c.svcCtx.Dao.SendChatMultiMessage(c.ctx, fromUserId, chatId, outBoxList)
	if err != nil {
		c.Logger.Errorf("msg.sendMultiMessage - error: %v")
		return nil, err
	}

	if cb != nil {
		err = cb(boxList)
		if err != nil {
			c.Logger.Errorf("msg.sendMultiMessage - error: %v")
			return nil, err
		}
	}

	var (
		updateNewMessageList = make([]*mtproto.Update, 0, len(boxList))
	)

	for _, outBoxMsg := range boxList {
		updateNewMessageList = append(updateNewMessageList, mtproto.MakeTLUpdateNewMessage(&mtproto.Update{
			Pts_INT32:       outBoxMsg.Pts,
			PtsCount:        outBoxMsg.PtsCount,
			RandomId:        outBoxMsg.RandomId,
			Message_MESSAGE: outBoxMsg.Message,
		}).To_Update())
	}

	rUpdates := mtproto.MakeReplyUpdates(
		func(idList []int64) []*mtproto.User {
			users, _ := c.svcCtx.Dao.UserClient.UserGetMutableUsers(c.ctx,
				&userpb.TLUserGetMutableUsers{
					Id: idList,
				})
			return users.GetUserListByIdList(fromUserId, idList...)
		},
		func(idList []int64) []*mtproto.Chat {
			chats, _ := c.svcCtx.Dao.ChatClient.ChatGetChatListByIdList(c.ctx,
				&chatpb.TLChatGetChatListByIdList{
					IdList: idList,
				})
			return chats.GetChatListByIdList(fromUserId, idList...)
		},
		func(idList []int64) []*mtproto.Chat {
			// TODO
			return nil
		},
		updateNewMessageList...)

	c.svcCtx.Dao.SyncClient.SyncUpdatesNotMe(c.ctx, &sync.TLSyncUpdatesNotMe{
		UserId:    fromUserId,
		AuthKeyId: fromAuthKeyId,
		Updates: mtproto.MakeSyncNotMeUpdates(
			func(idList []int64) []*mtproto.User {
				return rUpdates.Users
			},
			func(idList []int64) []*mtproto.Chat {
				return rUpdates.Chats
			},
			func(idList []int64) []*mtproto.Chat {
				// rUpdates.Chats include chats
				return nil
			},
			updateNewMessageList...),
	})

	return rUpdates, nil
}
