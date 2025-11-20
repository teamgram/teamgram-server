/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright 2025 Teamgram Authors.
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package service

import (
	"context"

	"github.com/teamgram/proto/mtproto"
	"github.com/teamgram/teamgram-server/app/messenger/msg/msgtransfer/internal/core"
	"github.com/teamgram/teamgram-server/app/messenger/msg/msgtransfer/msgtransfer"
)

// MsgtransferSendMessageToOutbox
// msgtransfer.sendMessageToOutbox flags:# user_id:long auth_key_id:long peer_type:int peer_id:long message:Vector<OutboxMessage> users:flags.1?Vector<User> chats:flags.2?Vector<Chat> = MessageBoxList;
func (s *Service) MsgtransferSendMessageToOutbox(ctx context.Context, request *msgtransfer.TLMsgtransferSendMessageToOutbox) (*mtproto.MessageBoxList, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("msgtransfer.sendMessageToOutbox - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.MsgtransferSendMessageToOutbox(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("msgtransfer.sendMessageToOutbox - reply: {%s}", r)
	return r, err
}

// MsgtransferSendMessageToInbox
// msgtransfer.sendMessageToInbox flags:# user_id:long from_id:long from_auth_keyId:long peer_type:int peer_id:long box_list:Vector<MessageBox> users:flags.1?Vector<User> chats:flags.2?Vector<Chat> = Void;
func (s *Service) MsgtransferSendMessageToInbox(ctx context.Context, request *msgtransfer.TLMsgtransferSendMessageToInbox) (*mtproto.Void, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("msgtransfer.sendMessageToInbox - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.MsgtransferSendMessageToInbox(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("msgtransfer.sendMessageToInbox - reply: {%s}", r)
	return r, err
}
