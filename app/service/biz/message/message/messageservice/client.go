/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2024-present,  Teamgram Authors.
 *  All rights reserved.
 *
 * Author: Benqi (wubenqi@gmail.com)
 */

package messageservice

import (
	"context"

	"github.com/teamgram/proto/v2/tg"
	"github.com/teamgram/teamgram-server/v2/app/service/biz/message/message"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/callopt"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	MessageGetUserMessage(ctx context.Context, req *message.TLMessageGetUserMessage, callOptions ...callopt.Option) (r *tg.MessageBox, err error)
	MessageGetUserMessageList(ctx context.Context, req *message.TLMessageGetUserMessageList, callOptions ...callopt.Option) (r *message.VectorMessageBox, err error)
	MessageGetUserMessageListByDataIdList(ctx context.Context, req *message.TLMessageGetUserMessageListByDataIdList, callOptions ...callopt.Option) (r *message.VectorMessageBox, err error)
	MessageGetUserMessageListByDataIdUserIdList(ctx context.Context, req *message.TLMessageGetUserMessageListByDataIdUserIdList, callOptions ...callopt.Option) (r *message.VectorMessageBox, err error)
	MessageGetHistoryMessages(ctx context.Context, req *message.TLMessageGetHistoryMessages, callOptions ...callopt.Option) (r *message.VectorMessageBox, err error)
	MessageGetHistoryMessagesCount(ctx context.Context, req *message.TLMessageGetHistoryMessagesCount, callOptions ...callopt.Option) (r *tg.Int32, err error)
	MessageGetPeerUserMessageId(ctx context.Context, req *message.TLMessageGetPeerUserMessageId, callOptions ...callopt.Option) (r *tg.Int32, err error)
	MessageGetPeerUserMessage(ctx context.Context, req *message.TLMessageGetPeerUserMessage, callOptions ...callopt.Option) (r *tg.MessageBox, err error)
	MessageSearchByMediaType(ctx context.Context, req *message.TLMessageSearchByMediaType, callOptions ...callopt.Option) (r *message.VectorMessageBox, err error)
	MessageSearch(ctx context.Context, req *message.TLMessageSearch, callOptions ...callopt.Option) (r *message.VectorMessageBox, err error)
	MessageSearchGlobal(ctx context.Context, req *message.TLMessageSearchGlobal, callOptions ...callopt.Option) (r *message.VectorMessageBox, err error)
	MessageSearchByPinned(ctx context.Context, req *message.TLMessageSearchByPinned, callOptions ...callopt.Option) (r *message.VectorMessageBox, err error)
	MessageGetSearchCounter(ctx context.Context, req *message.TLMessageGetSearchCounter, callOptions ...callopt.Option) (r *tg.Int32, err error)
	MessageSearchV2(ctx context.Context, req *message.TLMessageSearchV2, callOptions ...callopt.Option) (r *message.VectorMessageBox, err error)
	MessageGetLastTwoPinnedMessageId(ctx context.Context, req *message.TLMessageGetLastTwoPinnedMessageId, callOptions ...callopt.Option) (r *message.VectorInt, err error)
	MessageUpdatePinnedMessageId(ctx context.Context, req *message.TLMessageUpdatePinnedMessageId, callOptions ...callopt.Option) (r *tg.Bool, err error)
	MessageGetPinnedMessageIdList(ctx context.Context, req *message.TLMessageGetPinnedMessageIdList, callOptions ...callopt.Option) (r *message.VectorInt, err error)
	MessageUnPinAllMessages(ctx context.Context, req *message.TLMessageUnPinAllMessages, callOptions ...callopt.Option) (r *message.VectorInt, err error)
	MessageGetUnreadMentions(ctx context.Context, req *message.TLMessageGetUnreadMentions, callOptions ...callopt.Option) (r *message.VectorMessageBox, err error)
	MessageGetUnreadMentionsCount(ctx context.Context, req *message.TLMessageGetUnreadMentionsCount, callOptions ...callopt.Option) (r *tg.Int32, err error)
}

// NewClient creates a client for the service defined in IDL.
func NewClient(destService string, opts ...client.Option) (Client, error) {
	var options []client.Option
	options = append(options, client.WithDestService(destService))

	options = append(options, opts...)

	kc, err := client.NewClient(serviceInfoForClient(), options...)
	if err != nil {
		return nil, err
	}
	return &kMessageClient{
		kClient: newServiceClient(kc),
	}, nil
}

// MustNewClient creates a client for the service defined in IDL. It panics if any error occurs.
func MustNewClient(destService string, opts ...client.Option) Client {
	kc, err := NewClient(destService, opts...)
	if err != nil {
		panic(err)
	}
	return kc
}

type kMessageClient struct {
	*kClient
}

func NewRPCMessageClient(cli client.Client) Client {
	return &kMessageClient{
		kClient: newServiceClient(cli),
	}
}

func (p *kMessageClient) MessageGetUserMessage(ctx context.Context, req *message.TLMessageGetUserMessage, callOptions ...callopt.Option) (r *tg.MessageBox, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.MessageGetUserMessage(ctx, req)
}

func (p *kMessageClient) MessageGetUserMessageList(ctx context.Context, req *message.TLMessageGetUserMessageList, callOptions ...callopt.Option) (r *message.VectorMessageBox, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.MessageGetUserMessageList(ctx, req)
}

func (p *kMessageClient) MessageGetUserMessageListByDataIdList(ctx context.Context, req *message.TLMessageGetUserMessageListByDataIdList, callOptions ...callopt.Option) (r *message.VectorMessageBox, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.MessageGetUserMessageListByDataIdList(ctx, req)
}

func (p *kMessageClient) MessageGetUserMessageListByDataIdUserIdList(ctx context.Context, req *message.TLMessageGetUserMessageListByDataIdUserIdList, callOptions ...callopt.Option) (r *message.VectorMessageBox, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.MessageGetUserMessageListByDataIdUserIdList(ctx, req)
}

func (p *kMessageClient) MessageGetHistoryMessages(ctx context.Context, req *message.TLMessageGetHistoryMessages, callOptions ...callopt.Option) (r *message.VectorMessageBox, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.MessageGetHistoryMessages(ctx, req)
}

func (p *kMessageClient) MessageGetHistoryMessagesCount(ctx context.Context, req *message.TLMessageGetHistoryMessagesCount, callOptions ...callopt.Option) (r *tg.Int32, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.MessageGetHistoryMessagesCount(ctx, req)
}

func (p *kMessageClient) MessageGetPeerUserMessageId(ctx context.Context, req *message.TLMessageGetPeerUserMessageId, callOptions ...callopt.Option) (r *tg.Int32, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.MessageGetPeerUserMessageId(ctx, req)
}

func (p *kMessageClient) MessageGetPeerUserMessage(ctx context.Context, req *message.TLMessageGetPeerUserMessage, callOptions ...callopt.Option) (r *tg.MessageBox, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.MessageGetPeerUserMessage(ctx, req)
}

func (p *kMessageClient) MessageSearchByMediaType(ctx context.Context, req *message.TLMessageSearchByMediaType, callOptions ...callopt.Option) (r *message.VectorMessageBox, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.MessageSearchByMediaType(ctx, req)
}

func (p *kMessageClient) MessageSearch(ctx context.Context, req *message.TLMessageSearch, callOptions ...callopt.Option) (r *message.VectorMessageBox, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.MessageSearch(ctx, req)
}

func (p *kMessageClient) MessageSearchGlobal(ctx context.Context, req *message.TLMessageSearchGlobal, callOptions ...callopt.Option) (r *message.VectorMessageBox, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.MessageSearchGlobal(ctx, req)
}

func (p *kMessageClient) MessageSearchByPinned(ctx context.Context, req *message.TLMessageSearchByPinned, callOptions ...callopt.Option) (r *message.VectorMessageBox, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.MessageSearchByPinned(ctx, req)
}

func (p *kMessageClient) MessageGetSearchCounter(ctx context.Context, req *message.TLMessageGetSearchCounter, callOptions ...callopt.Option) (r *tg.Int32, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.MessageGetSearchCounter(ctx, req)
}

func (p *kMessageClient) MessageSearchV2(ctx context.Context, req *message.TLMessageSearchV2, callOptions ...callopt.Option) (r *message.VectorMessageBox, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.MessageSearchV2(ctx, req)
}

func (p *kMessageClient) MessageGetLastTwoPinnedMessageId(ctx context.Context, req *message.TLMessageGetLastTwoPinnedMessageId, callOptions ...callopt.Option) (r *message.VectorInt, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.MessageGetLastTwoPinnedMessageId(ctx, req)
}

func (p *kMessageClient) MessageUpdatePinnedMessageId(ctx context.Context, req *message.TLMessageUpdatePinnedMessageId, callOptions ...callopt.Option) (r *tg.Bool, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.MessageUpdatePinnedMessageId(ctx, req)
}

func (p *kMessageClient) MessageGetPinnedMessageIdList(ctx context.Context, req *message.TLMessageGetPinnedMessageIdList, callOptions ...callopt.Option) (r *message.VectorInt, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.MessageGetPinnedMessageIdList(ctx, req)
}

func (p *kMessageClient) MessageUnPinAllMessages(ctx context.Context, req *message.TLMessageUnPinAllMessages, callOptions ...callopt.Option) (r *message.VectorInt, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.MessageUnPinAllMessages(ctx, req)
}

func (p *kMessageClient) MessageGetUnreadMentions(ctx context.Context, req *message.TLMessageGetUnreadMentions, callOptions ...callopt.Option) (r *message.VectorMessageBox, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.MessageGetUnreadMentions(ctx, req)
}

func (p *kMessageClient) MessageGetUnreadMentionsCount(ctx context.Context, req *message.TLMessageGetUnreadMentionsCount, callOptions ...callopt.Option) (r *tg.Int32, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.MessageGetUnreadMentionsCount(ctx, req)
}
