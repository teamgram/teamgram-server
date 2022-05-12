/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright 2022 Teamgram Authors.
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package service

import (
	"context"

	"github.com/teamgram/teamgram-server/app/service/biz/message/message"
	"github.com/teamgram/teamgram-server/app/service/biz/message/internal/core"
	"github.com/teamgram/proto/mtproto"
)


// MessageGetUserMessage
// message.getUserMessage user_id:long id:int = MessageBox;
func (s *Service) MessageGetUserMessage(ctx context.Context, request *message.TLMessageGetUserMessage) (*mtproto.MessageBox, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("message.getUserMessage - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.MessageGetUserMessage(request)
	if err != nil {
		return nil, err
	}

	c.Infof("message.getUserMessage - reply: %s", r.DebugString())
	return r, err
}

// MessageGetUserMessageList
// message.getUserMessageList user_id:long id_list:Vector<int> = Vector<MessageBox>;
func (s *Service) MessageGetUserMessageList(ctx context.Context, request *message.TLMessageGetUserMessageList) (*message.Vector_MessageBox, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("message.getUserMessageList - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.MessageGetUserMessageList(request)
	if err != nil {
		return nil, err
	}

	c.Infof("message.getUserMessageList - reply: %s", r.DebugString())
	return r, err
}

// MessageGetUserMessageListByDataIdList
// message.getUserMessageListByDataIdList user_id:long id_list:Vector<long> = Vector<MessageBox>;
func (s *Service) MessageGetUserMessageListByDataIdList(ctx context.Context, request *message.TLMessageGetUserMessageListByDataIdList) (*message.Vector_MessageBox, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("message.getUserMessageListByDataIdList - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.MessageGetUserMessageListByDataIdList(request)
	if err != nil {
		return nil, err
	}

	c.Infof("message.getUserMessageListByDataIdList - reply: %s", r.DebugString())
	return r, err
}

// MessageGetHistoryMessages
// message.getHistoryMessages user_id:long peer_type:int peer_id:long offset_id:int offset_date:int add_offset:int limit:int max_id:int min_id:int hash:long = Vector<MessageBox>;
func (s *Service) MessageGetHistoryMessages(ctx context.Context, request *message.TLMessageGetHistoryMessages) (*message.Vector_MessageBox, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("message.getHistoryMessages - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.MessageGetHistoryMessages(request)
	if err != nil {
		return nil, err
	}

	c.Infof("message.getHistoryMessages - reply: %s", r.DebugString())
	return r, err
}

// MessageGetHistoryMessagesCount
// message.getHistoryMessagesCount user_id:long peer_type:int peer_id:long = Int32;
func (s *Service) MessageGetHistoryMessagesCount(ctx context.Context, request *message.TLMessageGetHistoryMessagesCount) (*mtproto.Int32, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("message.getHistoryMessagesCount - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.MessageGetHistoryMessagesCount(request)
	if err != nil {
		return nil, err
	}

	c.Infof("message.getHistoryMessagesCount - reply: %s", r.DebugString())
	return r, err
}

// MessageGetPeerUserMessageId
// message.getPeerUserMessageId user_id:long peer_user_id:long msg_id:int = Int32;
func (s *Service) MessageGetPeerUserMessageId(ctx context.Context, request *message.TLMessageGetPeerUserMessageId) (*mtproto.Int32, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("message.getPeerUserMessageId - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.MessageGetPeerUserMessageId(request)
	if err != nil {
		return nil, err
	}

	c.Infof("message.getPeerUserMessageId - reply: %s", r.DebugString())
	return r, err
}

// MessageGetPeerUserMessage
// message.getPeerUserMessage user_id:long peer_user_id:long msg_id:int = MessageBox;
func (s *Service) MessageGetPeerUserMessage(ctx context.Context, request *message.TLMessageGetPeerUserMessage) (*mtproto.MessageBox, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("message.getPeerUserMessage - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.MessageGetPeerUserMessage(request)
	if err != nil {
		return nil, err
	}

	c.Infof("message.getPeerUserMessage - reply: %s", r.DebugString())
	return r, err
}

// MessageGetPeerChatMessageIdList
// message.getPeerChatMessageIdList user_id:long peer_chat_id:long msg_id:int = Vector<PeerMessageId>;
func (s *Service) MessageGetPeerChatMessageIdList(ctx context.Context, request *message.TLMessageGetPeerChatMessageIdList) (*message.Vector_PeerMessageId, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("message.getPeerChatMessageIdList - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.MessageGetPeerChatMessageIdList(request)
	if err != nil {
		return nil, err
	}

	c.Infof("message.getPeerChatMessageIdList - reply: %s", r.DebugString())
	return r, err
}

// MessageGetPeerChatMessageList
// message.getPeerChatMessageList user_id:long peer_chat_id:long msg_id:int = Vector<MessageBox>;
func (s *Service) MessageGetPeerChatMessageList(ctx context.Context, request *message.TLMessageGetPeerChatMessageList) (*message.Vector_MessageBox, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("message.getPeerChatMessageList - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.MessageGetPeerChatMessageList(request)
	if err != nil {
		return nil, err
	}

	c.Infof("message.getPeerChatMessageList - reply: %s", r.DebugString())
	return r, err
}

// MessageSearchByMediaType
// message.searchByMediaType user_id:long peer_type:int peer_id:long media_type:int offset:int limit:int = Vector<MessageBox>;
func (s *Service) MessageSearchByMediaType(ctx context.Context, request *message.TLMessageSearchByMediaType) (*message.Vector_MessageBox, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("message.searchByMediaType - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.MessageSearchByMediaType(request)
	if err != nil {
		return nil, err
	}

	c.Infof("message.searchByMediaType - reply: %s", r.DebugString())
	return r, err
}

// MessageSearch
// message.search user_id:long peer_type:int peer_id:long q:string offset:int limit:int = Vector<MessageBox>;
func (s *Service) MessageSearch(ctx context.Context, request *message.TLMessageSearch) (*message.Vector_MessageBox, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("message.search - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.MessageSearch(request)
	if err != nil {
		return nil, err
	}

	c.Infof("message.search - reply: %s", r.DebugString())
	return r, err
}

// MessageSearchGlobal
// message.searchGlobal user_id:long q:string offset:int limit:int = Vector<MessageBox>;
func (s *Service) MessageSearchGlobal(ctx context.Context, request *message.TLMessageSearchGlobal) (*message.Vector_MessageBox, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("message.searchGlobal - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.MessageSearchGlobal(request)
	if err != nil {
		return nil, err
	}

	c.Infof("message.searchGlobal - reply: %s", r.DebugString())
	return r, err
}

// MessageSearchByPinned
// message.searchByPinned user_id:long peer_type:int peer_id:long = Vector<MessageBox>;
func (s *Service) MessageSearchByPinned(ctx context.Context, request *message.TLMessageSearchByPinned) (*message.Vector_MessageBox, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("message.searchByPinned - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.MessageSearchByPinned(request)
	if err != nil {
		return nil, err
	}

	c.Infof("message.searchByPinned - reply: %s", r.DebugString())
	return r, err
}

// MessageGetSearchCounter
// message.getSearchCounter user_id:long peer_type:int peer_id:long media_type:int = Int32;
func (s *Service) MessageGetSearchCounter(ctx context.Context, request *message.TLMessageGetSearchCounter) (*mtproto.Int32, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("message.getSearchCounter - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.MessageGetSearchCounter(request)
	if err != nil {
		return nil, err
	}

	c.Infof("message.getSearchCounter - reply: %s", r.DebugString())
	return r, err
}

// MessageSearchV2
// message.searchV2 user_id:long peer_type:int peer_id:long q:string from_id:long min_date:int max_date:int offset_id:int add_offset:int limit:int max_id:int min_id:int hash:long = Vector<MessageBox>;
func (s *Service) MessageSearchV2(ctx context.Context, request *message.TLMessageSearchV2) (*message.Vector_MessageBox, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("message.searchV2 - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.MessageSearchV2(request)
	if err != nil {
		return nil, err
	}

	c.Infof("message.searchV2 - reply: %s", r.DebugString())
	return r, err
}

// MessageGetLastTwoPinnedMessageId
// message.getLastTwoPinnedMessageId user_id:long peer_type:int peer_id:long = Vector<int>;
func (s *Service) MessageGetLastTwoPinnedMessageId(ctx context.Context, request *message.TLMessageGetLastTwoPinnedMessageId) (*message.Vector_Int, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("message.getLastTwoPinnedMessageId - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.MessageGetLastTwoPinnedMessageId(request)
	if err != nil {
		return nil, err
	}

	c.Infof("message.getLastTwoPinnedMessageId - reply: %s", r.DebugString())
	return r, err
}

// MessageUpdatePinnedMessageId
// message.updatePinnedMessageId user_id:long peer_type:int peer_id:long id:int pinned:Bool = Bool;
func (s *Service) MessageUpdatePinnedMessageId(ctx context.Context, request *message.TLMessageUpdatePinnedMessageId) (*mtproto.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("message.updatePinnedMessageId - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.MessageUpdatePinnedMessageId(request)
	if err != nil {
		return nil, err
	}

	c.Infof("message.updatePinnedMessageId - reply: %s", r.DebugString())
	return r, err
}

// MessageGetPinnedMessageIdList
// message.getPinnedMessageIdList user_id:long peer_type:int peer_id:long = Vector<int>;
func (s *Service) MessageGetPinnedMessageIdList(ctx context.Context, request *message.TLMessageGetPinnedMessageIdList) (*message.Vector_Int, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("message.getPinnedMessageIdList - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.MessageGetPinnedMessageIdList(request)
	if err != nil {
		return nil, err
	}

	c.Infof("message.getPinnedMessageIdList - reply: %s", r.DebugString())
	return r, err
}

// MessageUnPinAllMessages
// message.unPinAllMessages user_id:long peer_type:int peer_id:long = Vector<int>;
func (s *Service) MessageUnPinAllMessages(ctx context.Context, request *message.TLMessageUnPinAllMessages) (*message.Vector_Int, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("message.unPinAllMessages - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.MessageUnPinAllMessages(request)
	if err != nil {
		return nil, err
	}

	c.Infof("message.unPinAllMessages - reply: %s", r.DebugString())
	return r, err
}

// MessageGetUnreadMentions
// message.getUnreadMentions user_id:long peer_type:int peer_id:long offset_id:int add_offset:int limit:int min_id:int max_int:int = Vector<MessageBox>;
func (s *Service) MessageGetUnreadMentions(ctx context.Context, request *message.TLMessageGetUnreadMentions) (*message.Vector_MessageBox, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("message.getUnreadMentions - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.MessageGetUnreadMentions(request)
	if err != nil {
		return nil, err
	}

	c.Infof("message.getUnreadMentions - reply: %s", r.DebugString())
	return r, err
}

// MessageGetUnreadMentionsCount
// message.getUnreadMentionsCount user_id:long peer_type:int peer_id:long = Int32;
func (s *Service) MessageGetUnreadMentionsCount(ctx context.Context, request *message.TLMessageGetUnreadMentionsCount) (*mtproto.Int32, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("message.getUnreadMentionsCount - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.MessageGetUnreadMentionsCount(request)
	if err != nil {
		return nil, err
	}

	c.Infof("message.getUnreadMentionsCount - reply: %s", r.DebugString())
	return r, err
}

