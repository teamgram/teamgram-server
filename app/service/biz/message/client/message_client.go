/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright 2024 Teamgram Authors.
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package messageclient

import (
	"context"

	"github.com/teamgram/proto/v2/tg"
	"github.com/teamgram/teamgram-server/v2/app/service/biz/message/message"
	"github.com/teamgram/teamgram-server/v2/app/service/biz/message/message/messageservice"

	"github.com/cloudwego/kitex/client"
)

type MessageClient interface {
	MessageGetUserMessage(ctx context.Context, in *message.TLMessageGetUserMessage) (*tg.MessageBox, error)
	MessageGetUserMessageList(ctx context.Context, in *message.TLMessageGetUserMessageList) (*message.VectorMessageBox, error)
	MessageGetUserMessageListByDataIdList(ctx context.Context, in *message.TLMessageGetUserMessageListByDataIdList) (*message.VectorMessageBox, error)
	MessageGetUserMessageListByDataIdUserIdList(ctx context.Context, in *message.TLMessageGetUserMessageListByDataIdUserIdList) (*message.VectorMessageBox, error)
	MessageGetHistoryMessages(ctx context.Context, in *message.TLMessageGetHistoryMessages) (*message.VectorMessageBox, error)
	MessageGetHistoryMessagesCount(ctx context.Context, in *message.TLMessageGetHistoryMessagesCount) (*tg.Int32, error)
	MessageGetPeerUserMessageId(ctx context.Context, in *message.TLMessageGetPeerUserMessageId) (*tg.Int32, error)
	MessageGetPeerUserMessage(ctx context.Context, in *message.TLMessageGetPeerUserMessage) (*tg.MessageBox, error)
	MessageSearchByMediaType(ctx context.Context, in *message.TLMessageSearchByMediaType) (*message.VectorMessageBox, error)
	MessageSearch(ctx context.Context, in *message.TLMessageSearch) (*message.VectorMessageBox, error)
	MessageSearchGlobal(ctx context.Context, in *message.TLMessageSearchGlobal) (*message.VectorMessageBox, error)
	MessageSearchByPinned(ctx context.Context, in *message.TLMessageSearchByPinned) (*message.VectorMessageBox, error)
	MessageGetSearchCounter(ctx context.Context, in *message.TLMessageGetSearchCounter) (*tg.Int32, error)
	MessageSearchV2(ctx context.Context, in *message.TLMessageSearchV2) (*message.VectorMessageBox, error)
	MessageGetLastTwoPinnedMessageId(ctx context.Context, in *message.TLMessageGetLastTwoPinnedMessageId) (*message.VectorInt, error)
	MessageUpdatePinnedMessageId(ctx context.Context, in *message.TLMessageUpdatePinnedMessageId) (*tg.Bool, error)
	MessageGetPinnedMessageIdList(ctx context.Context, in *message.TLMessageGetPinnedMessageIdList) (*message.VectorInt, error)
	MessageUnPinAllMessages(ctx context.Context, in *message.TLMessageUnPinAllMessages) (*message.VectorInt, error)
	MessageGetUnreadMentions(ctx context.Context, in *message.TLMessageGetUnreadMentions) (*message.VectorMessageBox, error)
	MessageGetUnreadMentionsCount(ctx context.Context, in *message.TLMessageGetUnreadMentionsCount) (*tg.Int32, error)
}

type defaultMessageClient struct {
	cli client.Client
}

func NewMessageClient(cli client.Client) MessageClient {
	return &defaultMessageClient{
		cli: cli,
	}
}

// MessageGetUserMessage
// message.getUserMessage user_id:long id:int = MessageBox;
func (m *defaultMessageClient) MessageGetUserMessage(ctx context.Context, in *message.TLMessageGetUserMessage) (*tg.MessageBox, error) {
	cli := messageservice.NewRPCMessageClient(m.cli)
	return cli.MessageGetUserMessage(ctx, in)
}

// MessageGetUserMessageList
// message.getUserMessageList user_id:long id_list:Vector<int> = Vector<MessageBox>;
func (m *defaultMessageClient) MessageGetUserMessageList(ctx context.Context, in *message.TLMessageGetUserMessageList) (*message.VectorMessageBox, error) {
	cli := messageservice.NewRPCMessageClient(m.cli)
	return cli.MessageGetUserMessageList(ctx, in)
}

// MessageGetUserMessageListByDataIdList
// message.getUserMessageListByDataIdList user_id:long id_list:Vector<long> = Vector<MessageBox>;
func (m *defaultMessageClient) MessageGetUserMessageListByDataIdList(ctx context.Context, in *message.TLMessageGetUserMessageListByDataIdList) (*message.VectorMessageBox, error) {
	cli := messageservice.NewRPCMessageClient(m.cli)
	return cli.MessageGetUserMessageListByDataIdList(ctx, in)
}

// MessageGetUserMessageListByDataIdUserIdList
// message.getUserMessageListByDataIdUserIdList id:long user_id_list:Vector<long> = Vector<MessageBox>;
func (m *defaultMessageClient) MessageGetUserMessageListByDataIdUserIdList(ctx context.Context, in *message.TLMessageGetUserMessageListByDataIdUserIdList) (*message.VectorMessageBox, error) {
	cli := messageservice.NewRPCMessageClient(m.cli)
	return cli.MessageGetUserMessageListByDataIdUserIdList(ctx, in)
}

// MessageGetHistoryMessages
// message.getHistoryMessages user_id:long peer_type:int peer_id:long offset_id:int offset_date:int add_offset:int limit:int max_id:int min_id:int hash:long = Vector<MessageBox>;
func (m *defaultMessageClient) MessageGetHistoryMessages(ctx context.Context, in *message.TLMessageGetHistoryMessages) (*message.VectorMessageBox, error) {
	cli := messageservice.NewRPCMessageClient(m.cli)
	return cli.MessageGetHistoryMessages(ctx, in)
}

// MessageGetHistoryMessagesCount
// message.getHistoryMessagesCount user_id:long peer_type:int peer_id:long = Int32;
func (m *defaultMessageClient) MessageGetHistoryMessagesCount(ctx context.Context, in *message.TLMessageGetHistoryMessagesCount) (*tg.Int32, error) {
	cli := messageservice.NewRPCMessageClient(m.cli)
	return cli.MessageGetHistoryMessagesCount(ctx, in)
}

// MessageGetPeerUserMessageId
// message.getPeerUserMessageId user_id:long peer_user_id:long msg_id:int = Int32;
func (m *defaultMessageClient) MessageGetPeerUserMessageId(ctx context.Context, in *message.TLMessageGetPeerUserMessageId) (*tg.Int32, error) {
	cli := messageservice.NewRPCMessageClient(m.cli)
	return cli.MessageGetPeerUserMessageId(ctx, in)
}

// MessageGetPeerUserMessage
// message.getPeerUserMessage user_id:long peer_user_id:long msg_id:int = MessageBox;
func (m *defaultMessageClient) MessageGetPeerUserMessage(ctx context.Context, in *message.TLMessageGetPeerUserMessage) (*tg.MessageBox, error) {
	cli := messageservice.NewRPCMessageClient(m.cli)
	return cli.MessageGetPeerUserMessage(ctx, in)
}

// MessageSearchByMediaType
// message.searchByMediaType user_id:long peer_type:int peer_id:long media_type:int offset:int limit:int = Vector<MessageBox>;
func (m *defaultMessageClient) MessageSearchByMediaType(ctx context.Context, in *message.TLMessageSearchByMediaType) (*message.VectorMessageBox, error) {
	cli := messageservice.NewRPCMessageClient(m.cli)
	return cli.MessageSearchByMediaType(ctx, in)
}

// MessageSearch
// message.search user_id:long peer_type:int peer_id:long q:string offset:int limit:int = Vector<MessageBox>;
func (m *defaultMessageClient) MessageSearch(ctx context.Context, in *message.TLMessageSearch) (*message.VectorMessageBox, error) {
	cli := messageservice.NewRPCMessageClient(m.cli)
	return cli.MessageSearch(ctx, in)
}

// MessageSearchGlobal
// message.searchGlobal user_id:long q:string offset:int limit:int = Vector<MessageBox>;
func (m *defaultMessageClient) MessageSearchGlobal(ctx context.Context, in *message.TLMessageSearchGlobal) (*message.VectorMessageBox, error) {
	cli := messageservice.NewRPCMessageClient(m.cli)
	return cli.MessageSearchGlobal(ctx, in)
}

// MessageSearchByPinned
// message.searchByPinned user_id:long peer_type:int peer_id:long = Vector<MessageBox>;
func (m *defaultMessageClient) MessageSearchByPinned(ctx context.Context, in *message.TLMessageSearchByPinned) (*message.VectorMessageBox, error) {
	cli := messageservice.NewRPCMessageClient(m.cli)
	return cli.MessageSearchByPinned(ctx, in)
}

// MessageGetSearchCounter
// message.getSearchCounter user_id:long peer_type:int peer_id:long media_type:int = Int32;
func (m *defaultMessageClient) MessageGetSearchCounter(ctx context.Context, in *message.TLMessageGetSearchCounter) (*tg.Int32, error) {
	cli := messageservice.NewRPCMessageClient(m.cli)
	return cli.MessageGetSearchCounter(ctx, in)
}

// MessageSearchV2
// message.searchV2 user_id:long peer_type:int peer_id:long q:string from_id:long min_date:int max_date:int offset_id:int add_offset:int limit:int max_id:int min_id:int hash:long = Vector<MessageBox>;
func (m *defaultMessageClient) MessageSearchV2(ctx context.Context, in *message.TLMessageSearchV2) (*message.VectorMessageBox, error) {
	cli := messageservice.NewRPCMessageClient(m.cli)
	return cli.MessageSearchV2(ctx, in)
}

// MessageGetLastTwoPinnedMessageId
// message.getLastTwoPinnedMessageId user_id:long peer_type:int peer_id:long = Vector<int>;
func (m *defaultMessageClient) MessageGetLastTwoPinnedMessageId(ctx context.Context, in *message.TLMessageGetLastTwoPinnedMessageId) (*message.VectorInt, error) {
	cli := messageservice.NewRPCMessageClient(m.cli)
	return cli.MessageGetLastTwoPinnedMessageId(ctx, in)
}

// MessageUpdatePinnedMessageId
// message.updatePinnedMessageId user_id:long peer_type:int peer_id:long id:int pinned:Bool = Bool;
func (m *defaultMessageClient) MessageUpdatePinnedMessageId(ctx context.Context, in *message.TLMessageUpdatePinnedMessageId) (*tg.Bool, error) {
	cli := messageservice.NewRPCMessageClient(m.cli)
	return cli.MessageUpdatePinnedMessageId(ctx, in)
}

// MessageGetPinnedMessageIdList
// message.getPinnedMessageIdList user_id:long peer_type:int peer_id:long = Vector<int>;
func (m *defaultMessageClient) MessageGetPinnedMessageIdList(ctx context.Context, in *message.TLMessageGetPinnedMessageIdList) (*message.VectorInt, error) {
	cli := messageservice.NewRPCMessageClient(m.cli)
	return cli.MessageGetPinnedMessageIdList(ctx, in)
}

// MessageUnPinAllMessages
// message.unPinAllMessages user_id:long peer_type:int peer_id:long = Vector<int>;
func (m *defaultMessageClient) MessageUnPinAllMessages(ctx context.Context, in *message.TLMessageUnPinAllMessages) (*message.VectorInt, error) {
	cli := messageservice.NewRPCMessageClient(m.cli)
	return cli.MessageUnPinAllMessages(ctx, in)
}

// MessageGetUnreadMentions
// message.getUnreadMentions user_id:long peer_type:int peer_id:long offset_id:int add_offset:int limit:int min_id:int max_int:int = Vector<MessageBox>;
func (m *defaultMessageClient) MessageGetUnreadMentions(ctx context.Context, in *message.TLMessageGetUnreadMentions) (*message.VectorMessageBox, error) {
	cli := messageservice.NewRPCMessageClient(m.cli)
	return cli.MessageGetUnreadMentions(ctx, in)
}

// MessageGetUnreadMentionsCount
// message.getUnreadMentionsCount user_id:long peer_type:int peer_id:long = Int32;
func (m *defaultMessageClient) MessageGetUnreadMentionsCount(ctx context.Context, in *message.TLMessageGetUnreadMentionsCount) (*tg.Int32, error) {
	cli := messageservice.NewRPCMessageClient(m.cli)
	return cli.MessageGetUnreadMentionsCount(ctx, in)
}
