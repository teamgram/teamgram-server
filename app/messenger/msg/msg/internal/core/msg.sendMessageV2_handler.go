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
	"github.com/teamgram/teamgram-server/app/messenger/msg/msg/msg"
)

// MsgSendMessageV2
// msg.sendMessageV2 user_id:long auth_key_id:long peer_type:int peer_id:long message:Vector<OutboxMessage> = Updates;
func (c *MsgCore) MsgSendMessageV2(in *msg.TLMsgSendMessageV2) (*mtproto.Updates, error) {
	var (
		rUpdates *mtproto.Updates
		err      error
		// outBox   = in.GetMessage()
		peer = mtproto.MakePeerUtil(in.PeerType, in.PeerId)
	)

	if peer.IsChannel() {
		// c.Logger.Errorf("msg.sendMultiMessage blocked, License key from https://teamgram.net required to unlock enterprise features.")
		return nil, mtproto.ErrEnterpriseIsBlocked
	}

	for _, outBox := range in.GetMessage() {
		if outBox.GetScheduleDate().GetValue() != 0 {
			// c.Logger.Errorf("msg.sendMessage blocked, License key from https://teamgram.net required to unlock enterprise features.")
			return nil, mtproto.ErrEnterpriseIsBlocked
		}
	}

	if !peer.IsChatOrUser() {
		err = mtproto.ErrPeerIdInvalid
		c.Logger.Errorf("msg.sendMessage - error: %v", err)
		return nil, err
	}

	// TODO: check request valid

	if len(in.Message) == 0 {
		err = mtproto.ErrMessageIdsEmpty
		c.Logger.Errorf("msg.sendMessageV2 - error: %v", err)
		return nil, err
	}

	if peer.IsUser() {
		//rUpdates, err = c.sendUserOutgoingMessageV2(in.UserId, in.AuthKeyId, in.PeerId, outBox)
		//if err != nil {
		//	c.Logger.Errorf("msg.sendMessage - error: %v", err)
		//	return nil, err
		//}
	} else {
		//rUpdates, err = c.sendChatOutgoingMessageV2(in.UserId, in.AuthKeyId, in.PeerId, outBox)
		//if err != nil {
		//	c.Logger.Errorf("msg.sendMessage - error: %v", err)
		//	return nil, err
		//}
	}

	return rUpdates, nil
}

func (c *MsgCore) sendOutgoingMessageV2(in *msg.TLMsgSendMessageV2) (*mtproto.UpdateList, error) {
	var (
		rUpdates = &mtproto.UpdateList{Updates: nil}
	)

	// c.svcCtx.Dao.SendUserMessage()
	return rUpdates, nil
}
