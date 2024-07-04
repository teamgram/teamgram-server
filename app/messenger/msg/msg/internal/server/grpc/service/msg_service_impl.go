/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright 2024 Teamgram Authors.
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package service

import (
	"context"

	"github.com/teamgram/proto/mtproto"
	"github.com/teamgram/teamgram-server/app/messenger/msg/msg/internal/core"
	"github.com/teamgram/teamgram-server/app/messenger/msg/msg/msg"
)

// MsgPushUserMessage
// msg.pushUserMessage user_id:long auth_key_id:long peer_type:int peer_id:long push_type:int message:OutboxMessage = Bool;
func (s *Service) MsgPushUserMessage(ctx context.Context, request *msg.TLMsgPushUserMessage) (*mtproto.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("msg.pushUserMessage - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.MsgPushUserMessage(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("msg.pushUserMessage - reply: {%s}", r)
	return r, err
}

// MsgReadMessageContents
// msg.readMessageContents user_id:long auth_key_id:long peer_type:int peer_id:long id:Vector<ContentMessage> = messages.AffectedMessages;
func (s *Service) MsgReadMessageContents(ctx context.Context, request *msg.TLMsgReadMessageContents) (*mtproto.Messages_AffectedMessages, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("msg.readMessageContents - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.MsgReadMessageContents(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("msg.readMessageContents - reply: {%s}", r)
	return r, err
}

// MsgSendMessageV2
// msg.sendMessageV2 user_id:long auth_key_id:long peer_type:int peer_id:long message:Vector<OutboxMessage> = Updates;
func (s *Service) MsgSendMessageV2(ctx context.Context, request *msg.TLMsgSendMessageV2) (*mtproto.Updates, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("msg.sendMessageV2 - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.MsgSendMessageV2(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("msg.sendMessageV2 - reply: {%s}", r)
	return r, err
}

// MsgEditMessage
// msg.editMessage user_id:long auth_key_id:long peer_type:int peer_id:long edit_type:int message:OutboxMessage = Updates;
func (s *Service) MsgEditMessage(ctx context.Context, request *msg.TLMsgEditMessage) (*mtproto.Updates, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("msg.editMessage - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.MsgEditMessage(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("msg.editMessage - reply: {%s}", r)
	return r, err
}

// MsgEditMessageV2
// msg.editMessageV2 user_id:long auth_key_id:long peer_type:int peer_id:long edit_type:int new_message:OutboxMessage dst_message:MessageBox = Updates;
func (s *Service) MsgEditMessageV2(ctx context.Context, request *msg.TLMsgEditMessageV2) (*mtproto.Updates, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("msg.editMessageV2 - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.MsgEditMessageV2(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("msg.editMessageV2 - reply: {%s}", r)
	return r, err
}

// MsgDeleteMessages
// msg.deleteMessages flags:# user_id:long auth_key_id:long peer_type:int peer_id:long revoke:flags.1?true id:Vector<int> = messages.AffectedMessages;
func (s *Service) MsgDeleteMessages(ctx context.Context, request *msg.TLMsgDeleteMessages) (*mtproto.Messages_AffectedMessages, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("msg.deleteMessages - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.MsgDeleteMessages(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("msg.deleteMessages - reply: {%s}", r)
	return r, err
}

// MsgDeleteHistory
// msg.deleteHistory flags:# user_id:long auth_key_id:long peer_type:int peer_id:long just_clear:flags.0?true revoke:flags.1?true max_id:int = messages.AffectedHistory;
func (s *Service) MsgDeleteHistory(ctx context.Context, request *msg.TLMsgDeleteHistory) (*mtproto.Messages_AffectedHistory, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("msg.deleteHistory - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.MsgDeleteHistory(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("msg.deleteHistory - reply: {%s}", r)
	return r, err
}

// MsgDeletePhoneCallHistory
// msg.deletePhoneCallHistory flags:# user_id:long auth_key_id:long revoke:flags.1?true = messages.AffectedFoundMessages;
func (s *Service) MsgDeletePhoneCallHistory(ctx context.Context, request *msg.TLMsgDeletePhoneCallHistory) (*mtproto.Messages_AffectedFoundMessages, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("msg.deletePhoneCallHistory - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.MsgDeletePhoneCallHistory(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("msg.deletePhoneCallHistory - reply: {%s}", r)
	return r, err
}

// MsgDeleteChatHistory
// msg.deleteChatHistory chat_id:long delete_user_id:long = Bool;
func (s *Service) MsgDeleteChatHistory(ctx context.Context, request *msg.TLMsgDeleteChatHistory) (*mtproto.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("msg.deleteChatHistory - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.MsgDeleteChatHistory(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("msg.deleteChatHistory - reply: {%s}", r)
	return r, err
}

// MsgReadHistory
// msg.readHistory user_id:long auth_key_id:long peer_type:int peer_id:long max_id:int = messages.AffectedMessages;
func (s *Service) MsgReadHistory(ctx context.Context, request *msg.TLMsgReadHistory) (*mtproto.Messages_AffectedMessages, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("msg.readHistory - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.MsgReadHistory(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("msg.readHistory - reply: {%s}", r)
	return r, err
}

// MsgReadHistoryV2
// msg.readHistoryV2 user_id:long auth_key_id:long peer_type:int peer_id:long max_id:int = messages.AffectedMessages;
func (s *Service) MsgReadHistoryV2(ctx context.Context, request *msg.TLMsgReadHistoryV2) (*mtproto.Messages_AffectedMessages, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("msg.readHistoryV2 - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.MsgReadHistoryV2(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("msg.readHistoryV2 - reply: {%s}", r)
	return r, err
}

// MsgUpdatePinnedMessage
// msg.updatePinnedMessage flags:# user_id:long auth_key_id:long silent:flags.0?true unpin:flags.1?true pm_oneside:flags.2?true peer_type:int peer_id:long id:int = Updates;
func (s *Service) MsgUpdatePinnedMessage(ctx context.Context, request *msg.TLMsgUpdatePinnedMessage) (*mtproto.Updates, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("msg.updatePinnedMessage - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.MsgUpdatePinnedMessage(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("msg.updatePinnedMessage - reply: {%s}", r)
	return r, err
}

// MsgUnpinAllMessages
// msg.unpinAllMessages user_id:long auth_key_id:long peer_type:int peer_id:long = messages.AffectedHistory;
func (s *Service) MsgUnpinAllMessages(ctx context.Context, request *msg.TLMsgUnpinAllMessages) (*mtproto.Messages_AffectedHistory, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("msg.unpinAllMessages - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.MsgUnpinAllMessages(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("msg.unpinAllMessages - reply: {%s}", r)
	return r, err
}
