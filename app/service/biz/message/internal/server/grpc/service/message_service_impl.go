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
	"github.com/teamgram/teamgram-server/app/service/biz/message/internal/core"
	"github.com/teamgram/teamgram-server/app/service/biz/message/message"
)

// MessageGetUserMessage
// message.getUserMessage user_id:long id:int = MessageBox;
func (s *Service) MessageGetUserMessage(ctx context.Context, request *message.TLMessageGetUserMessage) (*mtproto.MessageBox, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("message.getUserMessage - metadata: %s, request: %s", c.MD, request)

	r, err := c.MessageGetUserMessage(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("message.getUserMessage - reply: %s", r)
	return r, err
}

// MessageGetUserMessageList
// message.getUserMessageList user_id:long id_list:Vector<int> = Vector<MessageBox>;
func (s *Service) MessageGetUserMessageList(ctx context.Context, request *message.TLMessageGetUserMessageList) (*message.Vector_MessageBox, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("message.getUserMessageList - metadata: %s, request: %s", c.MD, request)

	r, err := c.MessageGetUserMessageList(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("message.getUserMessageList - reply: %s", r)
	return r, err
}

// MessageGetUserMessageListByDataIdList
// message.getUserMessageListByDataIdList user_id:long id_list:Vector<long> = Vector<MessageBox>;
func (s *Service) MessageGetUserMessageListByDataIdList(ctx context.Context, request *message.TLMessageGetUserMessageListByDataIdList) (*message.Vector_MessageBox, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("message.getUserMessageListByDataIdList - metadata: %s, request: %s", c.MD, request)

	r, err := c.MessageGetUserMessageListByDataIdList(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("message.getUserMessageListByDataIdList - reply: %s", r)
	return r, err
}

// MessageGetUserMessageListByDataIdUserIdList
// message.getUserMessageListByDataIdUserIdList id:long user_id_list:Vector<long> = Vector<MessageBox>;
func (s *Service) MessageGetUserMessageListByDataIdUserIdList(ctx context.Context, request *message.TLMessageGetUserMessageListByDataIdUserIdList) (*message.Vector_MessageBox, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("message.getUserMessageListByDataIdUserIdList - metadata: %s, request: %s", c.MD, request)

	r, err := c.MessageGetUserMessageListByDataIdUserIdList(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("message.getUserMessageListByDataIdUserIdList - reply: %s", r)
	return r, err
}

// MessageGetHistoryMessages
// message.getHistoryMessages user_id:long peer_type:int peer_id:long offset_id:int offset_date:int add_offset:int limit:int max_id:int min_id:int hash:long = Vector<MessageBox>;
func (s *Service) MessageGetHistoryMessages(ctx context.Context, request *message.TLMessageGetHistoryMessages) (*message.Vector_MessageBox, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("message.getHistoryMessages - metadata: %s, request: %s", c.MD, request)

	r, err := c.MessageGetHistoryMessages(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("message.getHistoryMessages - reply: %s", r)
	return r, err
}

// MessageGetHistoryMessagesCount
// message.getHistoryMessagesCount user_id:long peer_type:int peer_id:long = Int32;
func (s *Service) MessageGetHistoryMessagesCount(ctx context.Context, request *message.TLMessageGetHistoryMessagesCount) (*mtproto.Int32, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("message.getHistoryMessagesCount - metadata: %s, request: %s", c.MD, request)

	r, err := c.MessageGetHistoryMessagesCount(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("message.getHistoryMessagesCount - reply: %s", r)
	return r, err
}

// MessageGetPeerUserMessageId
// message.getPeerUserMessageId user_id:long peer_user_id:long msg_id:int = Int32;
func (s *Service) MessageGetPeerUserMessageId(ctx context.Context, request *message.TLMessageGetPeerUserMessageId) (*mtproto.Int32, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("message.getPeerUserMessageId - metadata: %s, request: %s", c.MD, request)

	r, err := c.MessageGetPeerUserMessageId(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("message.getPeerUserMessageId - reply: %s", r)
	return r, err
}

// MessageGetPeerUserMessage
// message.getPeerUserMessage user_id:long peer_user_id:long msg_id:int = MessageBox;
func (s *Service) MessageGetPeerUserMessage(ctx context.Context, request *message.TLMessageGetPeerUserMessage) (*mtproto.MessageBox, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("message.getPeerUserMessage - metadata: %s, request: %s", c.MD, request)

	r, err := c.MessageGetPeerUserMessage(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("message.getPeerUserMessage - reply: %s", r)
	return r, err
}

// MessageSearchByMediaType
// message.searchByMediaType user_id:long peer_type:int peer_id:long media_type:int offset:int limit:int = MessageBoxList;
func (s *Service) MessageSearchByMediaType(ctx context.Context, request *message.TLMessageSearchByMediaType) (*mtproto.MessageBoxList, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("message.searchByMediaType - metadata: %s, request: %s", c.MD, request)

	r, err := c.MessageSearchByMediaType(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("message.searchByMediaType - reply: %s", r)
	return r, err
}

// MessageSearch
// message.search user_id:long peer_type:int peer_id:long q:string offset:int limit:int = MessageBoxList;
func (s *Service) MessageSearch(ctx context.Context, request *message.TLMessageSearch) (*mtproto.MessageBoxList, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("message.search - metadata: %s, request: %s", c.MD, request)

	r, err := c.MessageSearch(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("message.search - reply: %s", r)
	return r, err
}

// MessageSearchGlobal
// message.searchGlobal user_id:long q:string offset:int limit:int = MessageBoxList;
func (s *Service) MessageSearchGlobal(ctx context.Context, request *message.TLMessageSearchGlobal) (*mtproto.MessageBoxList, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("message.searchGlobal - metadata: %s, request: %s", c.MD, request)

	r, err := c.MessageSearchGlobal(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("message.searchGlobal - reply: %s", r)
	return r, err
}

// MessageSearchByPinned
// message.searchByPinned user_id:long peer_type:int peer_id:long = MessageBoxList;
func (s *Service) MessageSearchByPinned(ctx context.Context, request *message.TLMessageSearchByPinned) (*mtproto.MessageBoxList, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("message.searchByPinned - metadata: %s, request: %s", c.MD, request)

	r, err := c.MessageSearchByPinned(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("message.searchByPinned - reply: %s", r)
	return r, err
}

// MessageGetSearchCounter
// message.getSearchCounter user_id:long peer_type:int peer_id:long media_type:int = Int32;
func (s *Service) MessageGetSearchCounter(ctx context.Context, request *message.TLMessageGetSearchCounter) (*mtproto.Int32, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("message.getSearchCounter - metadata: %s, request: %s", c.MD, request)

	r, err := c.MessageGetSearchCounter(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("message.getSearchCounter - reply: %s", r)
	return r, err
}

// MessageSearchV2
// message.searchV2 user_id:long peer_type:int peer_id:long q:string from_id:long min_date:int max_date:int offset_id:int add_offset:int limit:int max_id:int min_id:int hash:long = MessageBoxList;
func (s *Service) MessageSearchV2(ctx context.Context, request *message.TLMessageSearchV2) (*mtproto.MessageBoxList, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("message.searchV2 - metadata: %s, request: %s", c.MD, request)

	r, err := c.MessageSearchV2(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("message.searchV2 - reply: %s", r)
	return r, err
}

// MessageGetLastTwoPinnedMessageId
// message.getLastTwoPinnedMessageId user_id:long peer_type:int peer_id:long = Vector<int>;
func (s *Service) MessageGetLastTwoPinnedMessageId(ctx context.Context, request *message.TLMessageGetLastTwoPinnedMessageId) (*message.Vector_Int, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("message.getLastTwoPinnedMessageId - metadata: %s, request: %s", c.MD, request)

	r, err := c.MessageGetLastTwoPinnedMessageId(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("message.getLastTwoPinnedMessageId - reply: %s", r)
	return r, err
}

// MessageUpdatePinnedMessageId
// message.updatePinnedMessageId user_id:long peer_type:int peer_id:long id:int pinned:Bool = Bool;
func (s *Service) MessageUpdatePinnedMessageId(ctx context.Context, request *message.TLMessageUpdatePinnedMessageId) (*mtproto.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("message.updatePinnedMessageId - metadata: %s, request: %s", c.MD, request)

	r, err := c.MessageUpdatePinnedMessageId(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("message.updatePinnedMessageId - reply: %s", r)
	return r, err
}

// MessageGetPinnedMessageIdList
// message.getPinnedMessageIdList user_id:long peer_type:int peer_id:long = Vector<int>;
func (s *Service) MessageGetPinnedMessageIdList(ctx context.Context, request *message.TLMessageGetPinnedMessageIdList) (*message.Vector_Int, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("message.getPinnedMessageIdList - metadata: %s, request: %s", c.MD, request)

	r, err := c.MessageGetPinnedMessageIdList(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("message.getPinnedMessageIdList - reply: %s", r)
	return r, err
}

// MessageUnPinAllMessages
// message.unPinAllMessages user_id:long peer_type:int peer_id:long = Vector<int>;
func (s *Service) MessageUnPinAllMessages(ctx context.Context, request *message.TLMessageUnPinAllMessages) (*message.Vector_Int, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("message.unPinAllMessages - metadata: %s, request: %s", c.MD, request)

	r, err := c.MessageUnPinAllMessages(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("message.unPinAllMessages - reply: %s", r)
	return r, err
}

// MessageGetUnreadMentions
// message.getUnreadMentions user_id:long peer_type:int peer_id:long offset_id:int add_offset:int limit:int min_id:int max_int:int = Vector<MessageBox>;
func (s *Service) MessageGetUnreadMentions(ctx context.Context, request *message.TLMessageGetUnreadMentions) (*message.Vector_MessageBox, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("message.getUnreadMentions - metadata: %s, request: %s", c.MD, request)

	r, err := c.MessageGetUnreadMentions(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("message.getUnreadMentions - reply: %s", r)
	return r, err
}

// MessageGetUnreadMentionsCount
// message.getUnreadMentionsCount user_id:long peer_type:int peer_id:long = Int32;
func (s *Service) MessageGetUnreadMentionsCount(ctx context.Context, request *message.TLMessageGetUnreadMentionsCount) (*mtproto.Int32, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("message.getUnreadMentionsCount - metadata: %s, request: %s", c.MD, request)

	r, err := c.MessageGetUnreadMentionsCount(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("message.getUnreadMentionsCount - reply: %s", r)
	return r, err
}

// MessageGetSavedHistoryMessages
// message.getSavedHistoryMessages user_id:long peer_type:int peer_id:long offset_id:int offset_date:int add_offset:int limit:int max_id:int min_id:int hash:long = MessageBoxList;
func (s *Service) MessageGetSavedHistoryMessages(ctx context.Context, request *message.TLMessageGetSavedHistoryMessages) (*mtproto.MessageBoxList, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("message.getSavedHistoryMessages - metadata: %s, request: %s", c.MD, request)

	r, err := c.MessageGetSavedHistoryMessages(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("message.getSavedHistoryMessages - reply: %s", r)
	return r, err
}
