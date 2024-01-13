/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright 2024 Teamgram Authors.
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package message_client

import (
	"context"

	"github.com/teamgram/proto/mtproto"
	"github.com/teamgram/teamgram-server/app/service/biz/message/message"

	"github.com/zeromicro/go-zero/zrpc"
)

var _ *mtproto.Bool

type MessageClient interface {
	MessageGetUserMessage(ctx context.Context, in *message.TLMessageGetUserMessage) (*mtproto.MessageBox, error)
	MessageGetUserMessageList(ctx context.Context, in *message.TLMessageGetUserMessageList) (*message.Vector_MessageBox, error)
	MessageGetUserMessageListByDataIdList(ctx context.Context, in *message.TLMessageGetUserMessageListByDataIdList) (*message.Vector_MessageBox, error)
	MessageGetUserMessageListByDataIdUserIdList(ctx context.Context, in *message.TLMessageGetUserMessageListByDataIdUserIdList) (*message.Vector_MessageBox, error)
	MessageGetHistoryMessages(ctx context.Context, in *message.TLMessageGetHistoryMessages) (*message.Vector_MessageBox, error)
	MessageGetHistoryMessagesCount(ctx context.Context, in *message.TLMessageGetHistoryMessagesCount) (*mtproto.Int32, error)
	MessageGetPeerUserMessageId(ctx context.Context, in *message.TLMessageGetPeerUserMessageId) (*mtproto.Int32, error)
	MessageGetPeerUserMessage(ctx context.Context, in *message.TLMessageGetPeerUserMessage) (*mtproto.MessageBox, error)
	MessageSearchByMediaType(ctx context.Context, in *message.TLMessageSearchByMediaType) (*mtproto.MessageBoxList, error)
	MessageSearch(ctx context.Context, in *message.TLMessageSearch) (*mtproto.MessageBoxList, error)
	MessageSearchGlobal(ctx context.Context, in *message.TLMessageSearchGlobal) (*mtproto.MessageBoxList, error)
	MessageSearchByPinned(ctx context.Context, in *message.TLMessageSearchByPinned) (*mtproto.MessageBoxList, error)
	MessageGetSearchCounter(ctx context.Context, in *message.TLMessageGetSearchCounter) (*mtproto.Int32, error)
	MessageSearchV2(ctx context.Context, in *message.TLMessageSearchV2) (*mtproto.MessageBoxList, error)
	MessageGetLastTwoPinnedMessageId(ctx context.Context, in *message.TLMessageGetLastTwoPinnedMessageId) (*message.Vector_Int, error)
	MessageUpdatePinnedMessageId(ctx context.Context, in *message.TLMessageUpdatePinnedMessageId) (*mtproto.Bool, error)
	MessageGetPinnedMessageIdList(ctx context.Context, in *message.TLMessageGetPinnedMessageIdList) (*message.Vector_Int, error)
	MessageUnPinAllMessages(ctx context.Context, in *message.TLMessageUnPinAllMessages) (*message.Vector_Int, error)
	MessageGetUnreadMentions(ctx context.Context, in *message.TLMessageGetUnreadMentions) (*message.Vector_MessageBox, error)
	MessageGetUnreadMentionsCount(ctx context.Context, in *message.TLMessageGetUnreadMentionsCount) (*mtproto.Int32, error)
	MessageGetSavedHistoryMessages(ctx context.Context, in *message.TLMessageGetSavedHistoryMessages) (*mtproto.MessageBoxList, error)
}

type defaultMessageClient struct {
	cli zrpc.Client
}

func NewMessageClient(cli zrpc.Client) MessageClient {
	return &defaultMessageClient{
		cli: cli,
	}
}

// MessageGetUserMessage
// message.getUserMessage user_id:long id:int = MessageBox;
func (m *defaultMessageClient) MessageGetUserMessage(ctx context.Context, in *message.TLMessageGetUserMessage) (*mtproto.MessageBox, error) {
	client := message.NewRPCMessageClient(m.cli.Conn())
	return client.MessageGetUserMessage(ctx, in)
}

// MessageGetUserMessageList
// message.getUserMessageList user_id:long id_list:Vector<int> = Vector<MessageBox>;
func (m *defaultMessageClient) MessageGetUserMessageList(ctx context.Context, in *message.TLMessageGetUserMessageList) (*message.Vector_MessageBox, error) {
	client := message.NewRPCMessageClient(m.cli.Conn())
	return client.MessageGetUserMessageList(ctx, in)
}

// MessageGetUserMessageListByDataIdList
// message.getUserMessageListByDataIdList user_id:long id_list:Vector<long> = Vector<MessageBox>;
func (m *defaultMessageClient) MessageGetUserMessageListByDataIdList(ctx context.Context, in *message.TLMessageGetUserMessageListByDataIdList) (*message.Vector_MessageBox, error) {
	client := message.NewRPCMessageClient(m.cli.Conn())
	return client.MessageGetUserMessageListByDataIdList(ctx, in)
}

// MessageGetUserMessageListByDataIdUserIdList
// message.getUserMessageListByDataIdUserIdList id:long user_id_list:Vector<long> = Vector<MessageBox>;
func (m *defaultMessageClient) MessageGetUserMessageListByDataIdUserIdList(ctx context.Context, in *message.TLMessageGetUserMessageListByDataIdUserIdList) (*message.Vector_MessageBox, error) {
	client := message.NewRPCMessageClient(m.cli.Conn())
	return client.MessageGetUserMessageListByDataIdUserIdList(ctx, in)
}

// MessageGetHistoryMessages
// message.getHistoryMessages user_id:long peer_type:int peer_id:long offset_id:int offset_date:int add_offset:int limit:int max_id:int min_id:int hash:long = Vector<MessageBox>;
func (m *defaultMessageClient) MessageGetHistoryMessages(ctx context.Context, in *message.TLMessageGetHistoryMessages) (*message.Vector_MessageBox, error) {
	client := message.NewRPCMessageClient(m.cli.Conn())
	return client.MessageGetHistoryMessages(ctx, in)
}

// MessageGetHistoryMessagesCount
// message.getHistoryMessagesCount user_id:long peer_type:int peer_id:long = Int32;
func (m *defaultMessageClient) MessageGetHistoryMessagesCount(ctx context.Context, in *message.TLMessageGetHistoryMessagesCount) (*mtproto.Int32, error) {
	client := message.NewRPCMessageClient(m.cli.Conn())
	return client.MessageGetHistoryMessagesCount(ctx, in)
}

// MessageGetPeerUserMessageId
// message.getPeerUserMessageId user_id:long peer_user_id:long msg_id:int = Int32;
func (m *defaultMessageClient) MessageGetPeerUserMessageId(ctx context.Context, in *message.TLMessageGetPeerUserMessageId) (*mtproto.Int32, error) {
	client := message.NewRPCMessageClient(m.cli.Conn())
	return client.MessageGetPeerUserMessageId(ctx, in)
}

// MessageGetPeerUserMessage
// message.getPeerUserMessage user_id:long peer_user_id:long msg_id:int = MessageBox;
func (m *defaultMessageClient) MessageGetPeerUserMessage(ctx context.Context, in *message.TLMessageGetPeerUserMessage) (*mtproto.MessageBox, error) {
	client := message.NewRPCMessageClient(m.cli.Conn())
	return client.MessageGetPeerUserMessage(ctx, in)
}

// MessageSearchByMediaType
// message.searchByMediaType user_id:long peer_type:int peer_id:long media_type:int offset:int limit:int = MessageBoxList;
func (m *defaultMessageClient) MessageSearchByMediaType(ctx context.Context, in *message.TLMessageSearchByMediaType) (*mtproto.MessageBoxList, error) {
	client := message.NewRPCMessageClient(m.cli.Conn())
	return client.MessageSearchByMediaType(ctx, in)
}

// MessageSearch
// message.search user_id:long peer_type:int peer_id:long q:string offset:int limit:int = MessageBoxList;
func (m *defaultMessageClient) MessageSearch(ctx context.Context, in *message.TLMessageSearch) (*mtproto.MessageBoxList, error) {
	client := message.NewRPCMessageClient(m.cli.Conn())
	return client.MessageSearch(ctx, in)
}

// MessageSearchGlobal
// message.searchGlobal user_id:long q:string offset:int limit:int = MessageBoxList;
func (m *defaultMessageClient) MessageSearchGlobal(ctx context.Context, in *message.TLMessageSearchGlobal) (*mtproto.MessageBoxList, error) {
	client := message.NewRPCMessageClient(m.cli.Conn())
	return client.MessageSearchGlobal(ctx, in)
}

// MessageSearchByPinned
// message.searchByPinned user_id:long peer_type:int peer_id:long = MessageBoxList;
func (m *defaultMessageClient) MessageSearchByPinned(ctx context.Context, in *message.TLMessageSearchByPinned) (*mtproto.MessageBoxList, error) {
	client := message.NewRPCMessageClient(m.cli.Conn())
	return client.MessageSearchByPinned(ctx, in)
}

// MessageGetSearchCounter
// message.getSearchCounter user_id:long peer_type:int peer_id:long media_type:int = Int32;
func (m *defaultMessageClient) MessageGetSearchCounter(ctx context.Context, in *message.TLMessageGetSearchCounter) (*mtproto.Int32, error) {
	client := message.NewRPCMessageClient(m.cli.Conn())
	return client.MessageGetSearchCounter(ctx, in)
}

// MessageSearchV2
// message.searchV2 user_id:long peer_type:int peer_id:long q:string from_id:long min_date:int max_date:int offset_id:int add_offset:int limit:int max_id:int min_id:int hash:long = MessageBoxList;
func (m *defaultMessageClient) MessageSearchV2(ctx context.Context, in *message.TLMessageSearchV2) (*mtproto.MessageBoxList, error) {
	client := message.NewRPCMessageClient(m.cli.Conn())
	return client.MessageSearchV2(ctx, in)
}

// MessageGetLastTwoPinnedMessageId
// message.getLastTwoPinnedMessageId user_id:long peer_type:int peer_id:long = Vector<int>;
func (m *defaultMessageClient) MessageGetLastTwoPinnedMessageId(ctx context.Context, in *message.TLMessageGetLastTwoPinnedMessageId) (*message.Vector_Int, error) {
	client := message.NewRPCMessageClient(m.cli.Conn())
	return client.MessageGetLastTwoPinnedMessageId(ctx, in)
}

// MessageUpdatePinnedMessageId
// message.updatePinnedMessageId user_id:long peer_type:int peer_id:long id:int pinned:Bool = Bool;
func (m *defaultMessageClient) MessageUpdatePinnedMessageId(ctx context.Context, in *message.TLMessageUpdatePinnedMessageId) (*mtproto.Bool, error) {
	client := message.NewRPCMessageClient(m.cli.Conn())
	return client.MessageUpdatePinnedMessageId(ctx, in)
}

// MessageGetPinnedMessageIdList
// message.getPinnedMessageIdList user_id:long peer_type:int peer_id:long = Vector<int>;
func (m *defaultMessageClient) MessageGetPinnedMessageIdList(ctx context.Context, in *message.TLMessageGetPinnedMessageIdList) (*message.Vector_Int, error) {
	client := message.NewRPCMessageClient(m.cli.Conn())
	return client.MessageGetPinnedMessageIdList(ctx, in)
}

// MessageUnPinAllMessages
// message.unPinAllMessages user_id:long peer_type:int peer_id:long = Vector<int>;
func (m *defaultMessageClient) MessageUnPinAllMessages(ctx context.Context, in *message.TLMessageUnPinAllMessages) (*message.Vector_Int, error) {
	client := message.NewRPCMessageClient(m.cli.Conn())
	return client.MessageUnPinAllMessages(ctx, in)
}

// MessageGetUnreadMentions
// message.getUnreadMentions user_id:long peer_type:int peer_id:long offset_id:int add_offset:int limit:int min_id:int max_int:int = Vector<MessageBox>;
func (m *defaultMessageClient) MessageGetUnreadMentions(ctx context.Context, in *message.TLMessageGetUnreadMentions) (*message.Vector_MessageBox, error) {
	client := message.NewRPCMessageClient(m.cli.Conn())
	return client.MessageGetUnreadMentions(ctx, in)
}

// MessageGetUnreadMentionsCount
// message.getUnreadMentionsCount user_id:long peer_type:int peer_id:long = Int32;
func (m *defaultMessageClient) MessageGetUnreadMentionsCount(ctx context.Context, in *message.TLMessageGetUnreadMentionsCount) (*mtproto.Int32, error) {
	client := message.NewRPCMessageClient(m.cli.Conn())
	return client.MessageGetUnreadMentionsCount(ctx, in)
}

// MessageGetSavedHistoryMessages
// message.getSavedHistoryMessages user_id:long peer_type:int peer_id:long offset_id:int offset_date:int add_offset:int limit:int max_id:int min_id:int hash:long = MessageBoxList;
func (m *defaultMessageClient) MessageGetSavedHistoryMessages(ctx context.Context, in *message.TLMessageGetSavedHistoryMessages) (*mtproto.MessageBoxList, error) {
	client := message.NewRPCMessageClient(m.cli.Conn())
	return client.MessageGetSavedHistoryMessages(ctx, in)
}
