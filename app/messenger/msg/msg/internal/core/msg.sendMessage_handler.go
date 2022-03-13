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
	"context"

	"github.com/teamgram/proto/mtproto"
	"github.com/teamgram/teamgram-server/app/messenger/msg/inbox/inbox"
	"github.com/teamgram/teamgram-server/app/messenger/msg/msg/msg"
	"github.com/teamgram/teamgram-server/app/messenger/sync/sync"
	chatpb "github.com/teamgram/teamgram-server/app/service/biz/chat/chat"
	userpb "github.com/teamgram/teamgram-server/app/service/biz/user/user"
)

// MsgSendMessage
// msg.sendMessage user_id:long auth_key_id:long peer_type:int peer_id:long message:OutboxMessage = Updates;
func (c *MsgCore) MsgSendMessage(in *msg.TLMsgSendMessage) (*mtproto.Updates, error) {
	var (
		rUpdates *mtproto.Updates
		err      error
		outBox   = in.GetMessage()
		peer     = mtproto.MakePeerUtil(in.PeerType, in.PeerId)
	)

	if peer.IsChannel() {
		// c.Logger.Errorf("msg.sendMultiMessage blocked, License key from https://teamgram.net required to unlock enterprise features.")
		return nil, mtproto.ErrEnterpriseIsBlocked
	}

	if outBox.GetScheduleDate().GetValue() != 0 {
		// c.Logger.Errorf("msg.sendMessage blocked, License key from https://teamgram.net required to unlock enterprise features.")
		return nil, mtproto.ErrEnterpriseIsBlocked
	}

	if !peer.IsChatOrUser() {
		err = mtproto.ErrPeerIdInvalid
		c.Logger.Errorf("msg.sendMessage - error: %v", err)
		return nil, err
	}

	if peer.IsUser() {
		rUpdates, err = c.sendUserOutgoingMessage(in.UserId, in.AuthKeyId, in.PeerId, outBox)
		if err != nil {
			c.Logger.Errorf("msg.sendMessage - error: %v", err)
			return nil, err
		}
	} else {
		rUpdates, err = c.sendChatOutgoingMessage(in.UserId, in.AuthKeyId, in.PeerId, outBox)
		if err != nil {
			c.Logger.Errorf("msg.sendMessage - error: %v", err)
			return nil, err
		}
	}

	return rUpdates, nil
}

func (c *MsgCore) sendUserOutgoingMessage(userId, authKeyId, peerUserId int64, outBox *msg.OutboxMessage) (*mtproto.Updates, error) {
	users, err := c.svcCtx.Dao.UserClient.UserGetMutableUsers(c.ctx, &userpb.TLUserGetMutableUsers{
		Id: []int64{userId, peerUserId},
	})
	if err != nil {
		c.Logger.Errorf("msg.sendUserOutgoingMessage - error: %v", err)
		return nil, err
	}

	sender, _ := users.GetImmutableUser(userId)
	if sender == nil || sender.Deleted() {
		err = mtproto.ErrInputUserDeactivated
		c.Logger.Errorf("msg.sendUserOutgoingMessage - error: %v", err)
		return nil, err
	}

	// TODO(@benqi): check
	// if sender.Restricted() {
	//	err = mtproto.ErrUserRestricted
	//	return
	// }

	peerUser, _ := users.GetImmutableUser(peerUserId)
	if peerUser == nil || peerUser.Deleted() {
		err = mtproto.ErrPeerIdInvalid
		c.Logger.Errorf("msg.sendUserOutgoingMessage - error: %v", err)
		return nil, err
	}

	sendMe := userId == peerUserId
	if !sendMe {
		// TODO(@benqi)
		// 1. check blocked
		// 2. span
	}

	var (
		rUpdates *mtproto.Updates
	)
	rUpdates, err = c.sendUserMessage(
		userId,
		authKeyId,
		peerUserId,
		outBox,
		func(did int64, inboxMsg *mtproto.Message) error {
			//inBox := &msgpb.InboxUserMessage{
			//	From:            makeSender(r.From),
			//	PeerUserId:      r.PeerId,
			//	RandomId:        r.Message.RandomId,
			//	DialogMessageId: did,
			//	MessageDataId:   mid,
			//	Message:         inboxMsg,
			//}
			//if model.IsBotFather(r.PeerUserId) {
			//	return s.botsClient.SendUserMessageToInbox(ctx, inBox)
			//} else {
			// toUser, _ := s.UserFacade.GetUserById(ctx, r.From.Id, r.PeerId)
			// log.Debug(toUser.DebugString())
			blocked, _ := c.svcCtx.Dao.UserClient.UserBlockedByUser(c.ctx, &userpb.TLUserBlockedByUser{
				UserId:     peerUserId,
				PeerUserId: userId,
			})
			// UserIsBlockedByUser(ctx, r.PeerId, r.From.Id)
			if !mtproto.FromBool(blocked) {
				_, err = c.svcCtx.Dao.InboxClient.InboxSendUserMessageToInbox(c.ctx, &inbox.TLInboxSendUserMessageToInbox{
					FromId:     userId,
					PeerUserId: peerUserId,
					Message: inbox.MakeTLInboxMessageData(&inbox.InboxMessageData{
						RandomId:        outBox.RandomId,
						DialogMessageId: did,
						// MessageDataId:   mid,
						Message: inboxMsg,
					}).To_InboxMessageData(),
				})
			}
			//}
			return nil
		})
	if err != nil {
		c.Logger.Errorf("msg.sendUserOutgoingMessage - error: %v", err)
		return nil, err
	}

	return rUpdates, nil
}

func (c *MsgCore) sendUserMessage(
	fromUserId int64,
	fromAuthKeyId int64,
	toUserId int64,
	outBox *msg.OutboxMessage,
	cb func(did int64, inboxMsg *mtproto.Message) error) (*mtproto.Updates, error) {

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
		return nil, err
	} else if hasDuplicateMessage {
		upd, err := c.svcCtx.Dao.GetDuplicateMessage(c.ctx, fromUserId, outBox.RandomId)
		if err != nil {
			c.Logger.Errorf("checkDuplicateMessage error - %v", err)
			return nil, err
		} else if upd != nil {
			return upd, nil
		}
	}

	box, err := c.svcCtx.Dao.SendUserMessage(c.ctx, fromUserId, toUserId, outBox)
	if err != nil {
		c.Logger.Error(err.Error())
		return nil, err
	}

	if !hasDuplicateMessage && cb != nil {
		err = cb(box.DialogMessageId, box.ToMessage(fromUserId))
		if err != nil {
			c.Logger.Error(err.Error())
			return nil, err
		}
	}

	updateNewMessage := mtproto.MakeTLUpdateNewMessage(&mtproto.Update{
		Pts_INT32:       box.Pts,
		PtsCount:        box.PtsCount,
		RandomId:        box.RandomId,
		Message_MESSAGE: box.Message,
	}).To_Update()

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
		updateNewMessage)

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
			updateNewMessage),
	})

	return rUpdates, nil
}

func (c *MsgCore) sendChatOutgoingMessage(userId, authKeyId, peerChatId int64, outBox *msg.OutboxMessage) (*mtproto.Updates, error) {
	rUpdates, err := c.sendChatMessage(c.ctx,
		userId,
		authKeyId,
		peerChatId,
		outBox,
		func(did int64, inboxMsg *mtproto.Message) error {
			_, err := c.svcCtx.Dao.InboxClient.InboxSendChatMessageToInbox(
				c.ctx,
				&inbox.TLInboxSendChatMessageToInbox{
					FromId:     userId,
					PeerChatId: peerChatId,
					Message: inbox.MakeTLInboxMessageData(&inbox.InboxMessageData{
						RandomId:        outBox.RandomId,
						DialogMessageId: did,
						// MessageDataId:   mid,
						Message: inboxMsg,
					}).To_InboxMessageData(),
				})
			if err != nil {
				c.Logger.Errorf("checkDuplicateMessage error - %v", err)
				return err
			}

			return err
		})
	if err != nil {
		c.Logger.Errorf("checkDuplicateMessage error - %v", err)
		return nil, err
	}

	return rUpdates, nil
}

func (c *MsgCore) sendChatMessage(
	ctx context.Context,
	fromUserId int64,
	fromAuthKeyId int64,
	chatId int64,
	outBox *msg.OutboxMessage,
	cb func(did int64, inboxMsg *mtproto.Message) error) (*mtproto.Updates, error) {

	hasDuplicateMessage, err := c.svcCtx.Dao.HasDuplicateMessage(ctx, fromUserId, outBox.RandomId)
	if err != nil {
		c.Logger.Errorf("checkDuplicateMessage error - %v", err)
		return nil, err
	} else if hasDuplicateMessage {
		upd, err := c.svcCtx.Dao.GetDuplicateMessage(ctx, fromUserId, outBox.RandomId)
		if err != nil {
			c.Logger.Errorf("checkDuplicateMessage error - %v", err)
			return nil, err
		} else if upd != nil {
			return upd, nil
		}
	}

	box, err := c.svcCtx.Dao.SendChatMessage(ctx, fromUserId, chatId, outBox)
	if err != nil {
		c.Logger.Error(err.Error())
		return nil, err
	}

	if !hasDuplicateMessage && cb != nil {
		err = cb(box.DialogMessageId, box.ToMessage(fromUserId))
		if err != nil {
			c.Logger.Error(err.Error())
			return nil, err
		}
	}

	updateNewMessage := mtproto.MakeTLUpdateNewMessage(&mtproto.Update{
		Pts_INT32:       box.Pts,
		PtsCount:        box.PtsCount,
		RandomId:        box.RandomId,
		Message_MESSAGE: box.Message,
	}).To_Update()

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
		updateNewMessage)

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
			updateNewMessage),
	})

	c.svcCtx.Dao.PutDuplicateMessage(ctx, fromUserId, outBox.RandomId, rUpdates)

	return rUpdates, nil
}
