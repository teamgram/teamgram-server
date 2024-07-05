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
	"strconv"
	"strings"

	"github.com/teamgram/proto/mtproto"
	"github.com/teamgram/teamgram-server/app/messenger/msg/inbox/inbox"
	"github.com/teamgram/teamgram-server/app/messenger/msg/msg/msg"
	"github.com/teamgram/teamgram-server/app/messenger/msg/msg/plugin"
	chatpb "github.com/teamgram/teamgram-server/app/service/biz/chat/chat"
	userpb "github.com/teamgram/teamgram-server/app/service/biz/user/user"

	"github.com/zeromicro/go-zero/core/mr"
)

// MsgSendMessageV2
// msg.sendMessageV2 user_id:long auth_key_id:long peer_type:int peer_id:long message:Vector<OutboxMessage> = Updates;
func (c *MsgCore) MsgSendMessageV2(in *msg.TLMsgSendMessageV2) (*mtproto.Updates, error) {
	var (
		rUpdates   *mtproto.Updates
		err        error
		outBoxList = in.GetMessage()
		peer       = mtproto.MakePeerUtil(in.PeerType, in.PeerId)
	)

	// TODO: check request valid

	if len(in.Message) == 0 {
		err = mtproto.ErrMessageIdsEmpty
		c.Logger.Errorf("msg.sendMessageV2 - error: %v", err)
		return nil, err
	}

	for _, outBox := range outBoxList {
		if outBox.GetScheduleDate().GetValue() != 0 {
			c.Logger.Errorf("msg.sendMessageV2 blocked, License key from https://teamgram.net required to unlock enterprise features.")
			return nil, mtproto.ErrEnterpriseIsBlocked
		}
	}

	switch peer.PeerType {
	case mtproto.PEER_USER:
		if len(outBoxList) == 1 {
			rUpdates, err = c.sendUserOutgoingMessageV2(in.UserId, in.AuthKeyId, in.PeerId, outBoxList[0])
			if err != nil {
				c.Logger.Errorf("msg.sendMessageV2 - error: %v", err)
				return nil, err
			}
		} else {
			rUpdates, err = c.sendUserOutgoingMessageList(in.UserId, in.AuthKeyId, in.PeerId, outBoxList)
			if err != nil {
				c.Logger.Errorf("msg.sendMessageV2 - error: %v", err)
				return nil, err
			}
		}
	case mtproto.PEER_CHAT:
		if len(outBoxList) == 1 {
			rUpdates, err = c.sendChatOutgoingMessageV2(in.UserId, in.AuthKeyId, in.PeerId, outBoxList[0])
			if err != nil {
				c.Logger.Errorf("msg.sendMessageV2 - error: %v", err)
				return nil, err
			}
		} else {
			rUpdates, err = c.sendChatOutgoingMessageList(in.UserId, in.AuthKeyId, in.PeerId, outBoxList)
			if err != nil {
				c.Logger.Errorf("msg.sendMessageV2 - error: %v", err)
				return nil, err
			}
		}
	case mtproto.PEER_CHANNEL:
		c.Logger.Errorf("msg.sendMessageV2 blocked, License key from https://teamgram.net required to unlock enterprise features.")
		return nil, mtproto.ErrEnterpriseIsBlocked
	default:
		c.Logger.Errorf("msg.sendMessageV2 - error: invalid peer(%v)", peer)
		err = mtproto.ErrPeerIdInvalid
		return nil, err
	}

	return rUpdates, nil
}

func (c *MsgCore) sendUserOutgoingMessageV2(fromUserId, fromAuthKeyId, toUserId int64, outBox *msg.OutboxMessage) (*mtproto.Updates, error) {
	var (
		idHelper = mtproto.NewIDListHelper(fromUserId, toUserId)
	)

	idHelper.PickByMessage(outBox.GetMessage())

	users, err := c.svcCtx.Dao.UserClient.UserGetMutableUsers(c.ctx, &userpb.TLUserGetMutableUsers{
		Id: idHelper.UserIdList,
		To: []int64{fromUserId, toUserId},
	})
	if err != nil {
		c.Logger.Errorf("msg.sendUserOutgoingMessageV2 - error: %v", err)
		return nil, err
	}

	sender, _ := users.GetImmutableUser(fromUserId)
	if sender == nil || sender.Deleted() {
		err = mtproto.ErrInputUserDeactivated
		c.Logger.Errorf("msg.sendUserOutgoingMessageV2 - error: %v", err)
		return nil, err
	}

	// TODO(@benqi): check
	// if sender.Restricted() {
	//	err = mtproto.ErrUserRestricted
	//	return
	// }

	peerUser, _ := users.GetImmutableUser(toUserId)
	if peerUser == nil || peerUser.Deleted() {
		err = mtproto.ErrInputUserDeactivated
		c.Logger.Errorf("msg.sendUserOutgoingMessage - error: %v", err)
		return nil, err
	}

	sendMe := fromUserId == toUserId
	if !sendMe {
		// TODO(@benqi)
		// 1. check blocked
		// 2. span
	}

	outBox.Message = plugin.RemakeMessage(
		c.ctx,
		c.svcCtx.MsgPlugin,
		outBox.Message,
		fromUserId,
		outBox.NoWebpage,
		func() bool {
			hasBot := false
			users.Visit(func(it *mtproto.ImmutableUser) {
				if it.IsBot() {
					hasBot = true
				}
			})

			return hasBot
		})

	var (
		// updateNewMessage *mtproto.Update
		rUpdates *mtproto.Updates
	)

	_, err = c.svcCtx.Dao.DoIdempotent(
		c.ctx,
		fromUserId,
		strconv.FormatInt(outBox.RandomId, 10),
		&rUpdates,
		func(ctx context.Context, v any) error {
			box, err := c.svcCtx.Dao.SendUserMessageV2(ctx, fromUserId, toUserId, outBox, true)
			if err != nil {
				c.Logger.Error(err.Error())
				return err
			}

			_, err2 := c.svcCtx.Dao.InboxClient.InboxSendUserMessageToInboxV2(
				c.ctx,
				&inbox.TLInboxSendUserMessageToInboxV2{
					UserId:        fromUserId,
					Out:           true,
					FromId:        fromUserId,
					FromAuthKeyId: fromAuthKeyId,
					PeerType:      mtproto.PEER_USER,
					PeerId:        toUserId,
					BoxList:       []*mtproto.MessageBox{box},
					Users:         users.GetUserListByIdList(fromUserId, idHelper.UserIdList...),
					Chats:         nil,
				})
			if err2 != nil {
				return err2
			}

			if fromUserId != toUserId {
				blocked, _ := c.svcCtx.Dao.UserClient.UserBlockedByUser(c.ctx, &userpb.TLUserBlockedByUser{
					UserId:     toUserId,
					PeerUserId: fromUserId,
				})

				if !mtproto.FromBool(blocked) {
					_, err2 = c.svcCtx.Dao.InboxClient.InboxSendUserMessageToInboxV2(
						c.ctx,
						&inbox.TLInboxSendUserMessageToInboxV2{
							UserId:        toUserId,
							Out:           false,
							FromId:        fromUserId,
							FromAuthKeyId: fromAuthKeyId,
							PeerType:      mtproto.PEER_USER,
							PeerId:        toUserId,
							BoxList:       []*mtproto.MessageBox{box},
							Users:         users.GetUserListByIdList(toUserId, idHelper.UserIdList...),
							Chats:         nil,
						})
					if err2 != nil {
						return err2
					}
				}
			}

			*v.(**mtproto.Updates) = mtproto.MakeReplyUpdates(
				func(idList []int64) []*mtproto.User {
					// TODO: check
					//users, _ := c.svcCtx.Dao.UserClient.UserGetMutableUsers(ctx,
					//	&userpb.TLUserGetMutableUsers{
					//		Id: idList,
					//	})
					return users.GetUserListByIdList(fromUserId, idList...)
				},
				func(idList []int64) []*mtproto.Chat {
					return []*mtproto.Chat{}
				},
				func(idList []int64) []*mtproto.Chat {
					// TODO
					return []*mtproto.Chat{}
				},
				mtproto.MakeTLUpdateNewMessage(&mtproto.Update{
					Pts_INT32:       box.Pts,
					PtsCount:        box.PtsCount,
					RandomId:        box.RandomId,
					Message_MESSAGE: box.Message,
				}).To_Update())

			return nil
		})
	if err != nil {
		c.Logger.Errorf("msg.sendUserOutgoingMessageV2 - error: %v", err)
		return nil, err
	}

	return rUpdates, nil
}

func (c *MsgCore) sendChatOutgoingMessageV2(fromUserId, fromAuthKeyId, peerChatId int64, outBox *msg.OutboxMessage) (*mtproto.Updates, error) {
	var (
		chat      *mtproto.MutableChat
		sUserList *mtproto.MutableUsers
		idHelper  = mtproto.NewIDListHelper(fromUserId)
	)
	idHelper.PickByMessage(outBox.GetMessage())

	err := mr.Finish(
		func() error {
			var (
				err error
			)
			chat, err = c.svcCtx.Dao.ChatClient.ChatGetMutableChat(
				c.ctx,
				&chatpb.TLChatGetMutableChat{
					ChatId: peerChatId,
				})
			if err != nil {
				c.Logger.Errorf("inbox.sendChatMessageToInbox - error: %v", err)
			}
			return err
		},
		func() error {
			var (
				err error
			)
			sUserList, err = c.svcCtx.Dao.UserClient.UserGetMutableUsersV2(
				c.ctx,
				&userpb.TLUserGetMutableUsersV2{
					Id:      idHelper.UserIdList,
					Privacy: true,
					HasTo:   true,
					To:      nil,
				})
			if err != nil {
				c.Logger.Errorf("inbox.sendChatMessageToInbox - error: %v", err)
			}

			return err
		})
	if err != nil {
		// c.Logger.Errorf("inbox.sendChatMessageToInbox - error: %v", err)
		return nil, err
	}

	if _, ok := chat.GetImmutableChatParticipant(fromUserId); !ok {
		c.Logger.Errorf("msg.sendChatOutgoingMessageV2 - error: ErrChatParticipantNotExists")
		err = mtproto.ErrChatWriteForbidden
		return nil, err
	}

	outBox.Message = plugin.RemakeMessage(
		c.ctx,
		c.svcCtx.MsgPlugin,
		outBox.Message,
		fromUserId,
		outBox.NoWebpage,
		func() bool {
			hasBot := false
			chat.Walk(func(userId int64, participant *mtproto.ImmutableChatParticipant) error {
				if participant.IsBot {
					hasBot = true
				}
				return nil
			})

			return hasBot
		})

	var (
		rUpdates *mtproto.Updates
	)

	_, err = c.svcCtx.Dao.DoIdempotent(
		c.ctx,
		fromUserId,
		strconv.FormatInt(outBox.RandomId, 10),
		&rUpdates,
		func(ctx context.Context, v any) error {
			box, err2 := c.svcCtx.Dao.SendChatMessageV2(ctx, fromUserId, peerChatId, outBox)
			if err2 != nil {
				c.Logger.Error(err2.Error())
				return err
			}

			chat.Walk(func(userId int64, participant *mtproto.ImmutableChatParticipant) error {
				if !participant.IsChatMemberStateNormal() {
					return nil
				}
				if err2 != nil {
					return nil
				}

				toUsers := make([]*mtproto.User, 0, sUserList.Length())
				sUserList.Visit(func(it *mtproto.ImmutableUser) {
					toUsers = append(toUsers, it.ToUser(participant.UserId))
				})
				_, err2 = c.svcCtx.Dao.InboxClient.InboxSendUserMessageToInboxV2(ctx, &inbox.TLInboxSendUserMessageToInboxV2{
					UserId:        participant.UserId,
					Out:           participant.UserId == fromUserId,
					FromId:        fromUserId,
					FromAuthKeyId: fromAuthKeyId,
					PeerType:      mtproto.PEER_CHAT,
					PeerId:        peerChatId,
					BoxList:       []*mtproto.MessageBox{box},
					Users:         toUsers,
					Chats:         []*mtproto.Chat{chat.ToUnsafeChat(participant.UserId)},
				})
				return nil
			})

			if err2 != nil {
				c.Logger.Error(err2.Error())
				return err
			}

			*v.(**mtproto.Updates) = mtproto.MakeReplyUpdates(
				func(idList []int64) []*mtproto.User {
					return sUserList.GetUserListByIdList(fromUserId, idList...)
				},
				func(idList []int64) []*mtproto.Chat {
					return []*mtproto.Chat{chat.ToUnsafeChat(fromUserId)}
				},
				func(idList []int64) []*mtproto.Chat {
					// TODO
					return nil
				},
				mtproto.MakeTLUpdateNewMessage(&mtproto.Update{
					Pts_INT32:       box.Pts,
					PtsCount:        box.PtsCount,
					RandomId:        box.RandomId,
					Message_MESSAGE: box.Message,
				}).To_Update())

			return nil
		})
	if err != nil {
		c.Logger.Errorf("msg.sendChatOutgoingMessageV2 - error: %v", err)
		return nil, err
	}

	return rUpdates, nil
}

func (c *MsgCore) sendUserOutgoingMessageList(fromUserId, fromAuthKeyId, toUserId int64, outBoxList []*msg.OutboxMessage) (*mtproto.Updates, error) {
	var (
		idHelper     = mtproto.NewIDListHelper(fromUserId, toUserId)
		rUpdates     *mtproto.Updates
		randomIdList []string
	)

	for _, outBox := range outBoxList {
		idHelper.PickByMessage(outBox.GetMessage())
		randomIdList = append(randomIdList, strconv.FormatInt(outBox.RandomId, 10))
	}

	users, err := c.svcCtx.Dao.UserClient.UserGetMutableUsers(c.ctx, &userpb.TLUserGetMutableUsers{
		Id: idHelper.UserIdList,
		To: []int64{fromUserId, toUserId},
	})
	if err != nil {
		c.Logger.Errorf("msg.sendUserOutgoingMessageV2 - error: %v", err)
		return nil, err
	}

	sender, _ := users.GetImmutableUser(fromUserId)
	if sender == nil || sender.Deleted() {
		err = mtproto.ErrInputUserDeactivated
		c.Logger.Errorf("msg.sendUserOutgoingMessageV2 - error: %v", err)
		return nil, err
	}

	// TODO(@benqi): check
	// if sender.Restricted() {
	//	err = mtproto.ErrUserRestricted
	//	return
	// }

	peerUser, _ := users.GetImmutableUser(toUserId)
	if peerUser == nil || peerUser.Deleted() {
		err = mtproto.ErrInputUserDeactivated
		c.Logger.Errorf("msg.sendUserOutgoingMessage - error: %v", err)
		return nil, err
	}

	sendMe := fromUserId == toUserId
	if !sendMe {
		// TODO(@benqi)
		// 1. check blocked
		// 2. span
	}

	for _, outBox := range outBoxList {
		outBox.Message = plugin.RemakeMessage(
			c.ctx,
			c.svcCtx.MsgPlugin,
			outBox.Message,
			fromUserId,
			outBox.NoWebpage,
			func() bool {
				hasBot := false
				users.Visit(func(it *mtproto.ImmutableUser) {
					if it.IsBot() {
						hasBot = true
					}
				})

				return hasBot
			})
	}

	_, err = c.svcCtx.Dao.DoIdempotent(
		c.ctx,
		fromUserId,
		strings.Join(randomIdList, "-"),
		&rUpdates,
		func(ctx context.Context, v any) error {
			var (
				boxList    []*mtproto.MessageBox
				updateList []*mtproto.Update
			)

			for _, outBox := range outBoxList {
				box, err := c.svcCtx.Dao.SendUserMessageV2(ctx, fromUserId, toUserId, outBox, true)
				if err != nil {
					c.Logger.Error(err.Error())
					return err
				}
				boxList = append(boxList, box)
				updateList = append(updateList, mtproto.MakeTLUpdateNewMessage(&mtproto.Update{
					Pts_INT32:       box.Pts,
					PtsCount:        box.PtsCount,
					RandomId:        box.RandomId,
					Message_MESSAGE: box.Message,
				}).To_Update())
			}

			_, err2 := c.svcCtx.Dao.InboxClient.InboxSendUserMessageToInboxV2(
				c.ctx,
				&inbox.TLInboxSendUserMessageToInboxV2{
					UserId:        fromUserId,
					Out:           true,
					FromId:        fromUserId,
					FromAuthKeyId: fromAuthKeyId,
					PeerType:      mtproto.PEER_USER,
					PeerId:        toUserId,
					BoxList:       boxList,
					Users:         users.GetUserListByIdList(fromUserId, idHelper.UserIdList...),
					Chats:         nil,
				})
			if err2 != nil {
				return err2
			}

			if fromUserId != toUserId {
				blocked, _ := c.svcCtx.Dao.UserClient.UserBlockedByUser(c.ctx, &userpb.TLUserBlockedByUser{
					UserId:     toUserId,
					PeerUserId: fromUserId,
				})

				if !mtproto.FromBool(blocked) {
					_, err2 = c.svcCtx.Dao.InboxClient.InboxSendUserMessageToInboxV2(
						c.ctx,
						&inbox.TLInboxSendUserMessageToInboxV2{
							UserId:        toUserId,
							Out:           false,
							FromId:        fromUserId,
							FromAuthKeyId: fromAuthKeyId,
							PeerType:      mtproto.PEER_USER,
							PeerId:        toUserId,
							BoxList:       boxList,
							Users:         users.GetUserListByIdList(toUserId, idHelper.UserIdList...),
							Chats:         nil,
						})
					if err2 != nil {
						return err2
					}
				}
			}

			*v.(**mtproto.Updates) = mtproto.MakeReplyUpdates(
				func(idList []int64) []*mtproto.User {
					// TODO: check
					//users, _ := c.svcCtx.Dao.UserClient.UserGetMutableUsers(ctx,
					//	&userpb.TLUserGetMutableUsers{
					//		Id: idList,
					//	})
					return users.GetUserListByIdList(fromUserId, idList...)
				},
				func(idList []int64) []*mtproto.Chat {
					return []*mtproto.Chat{}
				},
				func(idList []int64) []*mtproto.Chat {
					// TODO
					return []*mtproto.Chat{}
				},
				updateList...)

			return nil
		})
	if err != nil {
		c.Logger.Errorf("msg.sendUserOutgoingMessageV2 - error: %v", err)
		return nil, err
	}
	//}

	return rUpdates, nil
}

func (c *MsgCore) sendChatOutgoingMessageList(fromUserId, fromAuthKeyId, peerChatId int64, outBoxList []*msg.OutboxMessage) (*mtproto.Updates, error) {
	var (
		chat         *mtproto.MutableChat
		sUserList    *mtproto.MutableUsers
		idHelper     = mtproto.NewIDListHelper(fromUserId)
		rUpdates     *mtproto.Updates
		randomIdList []string
	)

	for _, outBox := range outBoxList {
		idHelper.PickByMessage(outBox.GetMessage())
		randomIdList = append(randomIdList, strconv.FormatInt(outBox.RandomId, 10))
	}

	err := mr.Finish(
		func() error {
			var (
				err error
			)
			chat, err = c.svcCtx.Dao.ChatClient.ChatGetMutableChat(
				c.ctx,
				&chatpb.TLChatGetMutableChat{
					ChatId: peerChatId,
				})
			if err != nil {
				c.Logger.Errorf("inbox.sendChatMessageToInbox - error: %v", err)
			}
			return err
		},
		func() error {
			var (
				err error
			)
			sUserList, err = c.svcCtx.Dao.UserClient.UserGetMutableUsersV2(
				c.ctx,
				&userpb.TLUserGetMutableUsersV2{
					Id:      idHelper.UserIdList,
					Privacy: true,
					HasTo:   true,
					To:      nil,
				})
			if err != nil {
				c.Logger.Errorf("inbox.sendChatMessageToInbox - error: %v", err)
			}

			return err
		})
	if err != nil {
		// c.Logger.Errorf("inbox.sendChatMessageToInbox - error: %v", err)
		return nil, err
	}

	if _, ok := chat.GetImmutableChatParticipant(fromUserId); !ok {
		c.Logger.Errorf("msg.sendChatOutgoingMessageV2 - error: ErrChatParticipantNotExists")
		err = mtproto.ErrChatWriteForbidden
		return nil, err
	}

	for _, outBox := range outBoxList {
		outBox.Message = plugin.RemakeMessage(
			c.ctx,
			c.svcCtx.MsgPlugin,
			outBox.Message,
			fromUserId,
			outBox.NoWebpage,
			func() bool {
				hasBot := false
				chat.Walk(func(userId int64, participant *mtproto.ImmutableChatParticipant) error {
					if participant.IsBot {
						hasBot = true
					}
					return nil
				})

				return hasBot
			})
	}

	//var (
	//	rUpdates *mtproto.Updates
	//)

	_, err = c.svcCtx.Dao.DoIdempotent(
		c.ctx,
		fromUserId,
		strings.Join(randomIdList, "-"),
		&rUpdates,
		func(ctx context.Context, v any) error {
			var (
				boxList    []*mtproto.MessageBox
				updateList []*mtproto.Update
			)

			for _, outBox := range outBoxList {
				box, err2 := c.svcCtx.Dao.SendChatMessageV2(ctx, fromUserId, peerChatId, outBox)
				if err2 != nil {
					c.Logger.Error(err2.Error())
					return err
				}
				boxList = append(boxList, box)
				updateList = append(updateList, mtproto.MakeTLUpdateNewMessage(&mtproto.Update{
					Pts_INT32:       box.Pts,
					PtsCount:        box.PtsCount,
					RandomId:        box.RandomId,
					Message_MESSAGE: box.Message,
				}).To_Update())
			}

			var (
				err2 error
			)

			chat.Walk(func(userId int64, participant *mtproto.ImmutableChatParticipant) error {
				if !participant.IsChatMemberStateNormal() {
					return nil
				}
				if err2 != nil {
					return nil
				}

				toUsers := make([]*mtproto.User, 0, sUserList.Length())
				sUserList.Visit(func(it *mtproto.ImmutableUser) {
					toUsers = append(toUsers, it.ToUser(participant.UserId))
				})
				_, err2 = c.svcCtx.Dao.InboxClient.InboxSendUserMessageToInboxV2(ctx, &inbox.TLInboxSendUserMessageToInboxV2{
					UserId:        participant.UserId,
					Out:           participant.UserId == fromUserId,
					FromId:        fromUserId,
					FromAuthKeyId: fromAuthKeyId,
					PeerType:      mtproto.PEER_CHAT,
					PeerId:        peerChatId,
					BoxList:       boxList,
					Users:         toUsers,
					Chats:         []*mtproto.Chat{chat.ToUnsafeChat(participant.UserId)},
				})
				return nil
			})

			if err2 != nil {
				c.Logger.Error(err2.Error())
				return err
			}

			*v.(**mtproto.Updates) = mtproto.MakeReplyUpdates(
				func(idList []int64) []*mtproto.User {
					return sUserList.GetUserListByIdList(fromUserId, idList...)
				},
				func(idList []int64) []*mtproto.Chat {
					return []*mtproto.Chat{chat.ToUnsafeChat(fromUserId)}
				},
				func(idList []int64) []*mtproto.Chat {
					// TODO
					return nil
				},
				updateList...)

			return nil
		})
	if err != nil {
		c.Logger.Errorf("msg.sendChatOutgoingMessageV2 - error: %v", err)
		return nil, err
	}

	return rUpdates, nil
}
