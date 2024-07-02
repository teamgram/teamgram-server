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
	"github.com/teamgram/teamgram-server/app/messenger/msg/msg/plugin"
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

	sendMe := in.UserId == peer.PeerId
	if !sendMe {
		// TODO(@benqi)
		// 1. check blocked
		// 2. span
	}

	// TODO(@benqi): r.From.Type
	switch in.PushType {
	case 0:
		_, err := c.sendUserOutgoingMessageV2(in.UserId, 0, peer.PeerId, boxMsg)
		if err != nil {
			return nil, err
		}
	default:
		var (
			idHelper = mtproto.NewIDListHelper(in.UserId, peer.PeerId)
		)

		idHelper.PickByMessage(in.GetMessage().GetMessage())

		users, err := c.svcCtx.Dao.UserClient.UserGetMutableUsers(c.ctx, &userpb.TLUserGetMutableUsers{
			Id: idHelper.UserIdList,
			To: []int64{peer.PeerId},
		})
		if err != nil {
			c.Logger.Errorf("msg.sendUserOutgoingMessageV2 - error: %v", err)
			return nil, err
		}

		sender, _ := users.GetImmutableUser(in.UserId)
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

		peerUser, _ := users.GetImmutableUser(peer.PeerId)
		if peerUser == nil || peerUser.Deleted() {
			err = mtproto.ErrInputUserDeactivated
			c.Logger.Errorf("msg.sendUserOutgoingMessage - error: %v", err)
			return nil, err
		}

		//sendMe := fromUserId == toUserId
		//if !sendMe {
		//	// TODO(@benqi)
		//	// 1. check blocked
		//	// 2. span
		//}

		in.Message.Message = plugin.RemakeMessage(
			c.ctx,
			c.svcCtx.MsgPlugin,
			in.Message.Message,
			in.UserId,
			false,
			func() bool {
				hasBot := false
				users.Visit(func(it *mtproto.ImmutableUser) {
					if it.IsBot() {
						hasBot = true
					}
				})

				return hasBot
			})

		box, err := c.svcCtx.Dao.SendUserMessageV2(c.ctx, in.UserId, peer.PeerId, in.Message, false)
		if err != nil {
			c.Logger.Error(err.Error())
			return nil, err
		}

		_, err = c.svcCtx.Dao.InboxClient.InboxSendUserMessageToInboxV2(
			c.ctx,
			&inbox.TLInboxSendUserMessageToInboxV2{
				UserId:        peer.PeerId,
				Out:           false,
				FromId:        in.UserId,
				FromAuthKeyId: 0,
				PeerType:      mtproto.PEER_USER,
				PeerId:        peer.PeerId,
				BoxList:       []*mtproto.MessageBox{box},
				Users:         users.GetUserListByIdList(peer.PeerId, idHelper.UserIdList...),
				Chats:         nil,
			})
		if err != nil {
			return nil, err
		}
	}

	return mtproto.BoolTrue, nil
}
