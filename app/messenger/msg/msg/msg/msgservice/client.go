/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2024-present,  Teamgram Authors.
 *  All rights reserved.
 *
 * Author: Benqi (wubenqi@gmail.com)
 */

package msgservice

import (
	"context"

	"github.com/teamgram/proto/v2/tg"
	"github.com/teamgram/teamgram-server/v2/app/messenger/msg/msg/msg"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/callopt"
)

var _ *tg.Bool

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	MsgPushUserMessage(ctx context.Context, req *msg.TLMsgPushUserMessage, callOptions ...callopt.Option) (r *tg.Bool, err error)
	MsgReadMessageContents(ctx context.Context, req *msg.TLMsgReadMessageContents, callOptions ...callopt.Option) (r *tg.MessagesAffectedMessages, err error)
	MsgSendMessageV2(ctx context.Context, req *msg.TLMsgSendMessageV2, callOptions ...callopt.Option) (r *tg.Updates, err error)
	MsgEditMessageV2(ctx context.Context, req *msg.TLMsgEditMessageV2, callOptions ...callopt.Option) (r *tg.Updates, err error)
	MsgDeleteMessages(ctx context.Context, req *msg.TLMsgDeleteMessages, callOptions ...callopt.Option) (r *tg.MessagesAffectedMessages, err error)
	MsgDeleteHistory(ctx context.Context, req *msg.TLMsgDeleteHistory, callOptions ...callopt.Option) (r *tg.MessagesAffectedHistory, err error)
	MsgDeletePhoneCallHistory(ctx context.Context, req *msg.TLMsgDeletePhoneCallHistory, callOptions ...callopt.Option) (r *tg.MessagesAffectedFoundMessages, err error)
	MsgDeleteChatHistory(ctx context.Context, req *msg.TLMsgDeleteChatHistory, callOptions ...callopt.Option) (r *tg.Bool, err error)
	MsgReadHistory(ctx context.Context, req *msg.TLMsgReadHistory, callOptions ...callopt.Option) (r *tg.MessagesAffectedMessages, err error)
	MsgReadHistoryV2(ctx context.Context, req *msg.TLMsgReadHistoryV2, callOptions ...callopt.Option) (r *tg.MessagesAffectedMessages, err error)
	MsgUpdatePinnedMessage(ctx context.Context, req *msg.TLMsgUpdatePinnedMessage, callOptions ...callopt.Option) (r *tg.Updates, err error)
	MsgUnpinAllMessages(ctx context.Context, req *msg.TLMsgUnpinAllMessages, callOptions ...callopt.Option) (r *tg.MessagesAffectedHistory, err error)
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
	return &kMsgClient{
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

type kMsgClient struct {
	*kClient
}

func NewRPCMsgClient(cli client.Client) Client {
	return &kMsgClient{
		kClient: newServiceClient(cli),
	}
}

func (p *kMsgClient) MsgPushUserMessage(ctx context.Context, req *msg.TLMsgPushUserMessage, callOptions ...callopt.Option) (r *tg.Bool, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.MsgPushUserMessage(ctx, req)
}

func (p *kMsgClient) MsgReadMessageContents(ctx context.Context, req *msg.TLMsgReadMessageContents, callOptions ...callopt.Option) (r *tg.MessagesAffectedMessages, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.MsgReadMessageContents(ctx, req)
}

func (p *kMsgClient) MsgSendMessageV2(ctx context.Context, req *msg.TLMsgSendMessageV2, callOptions ...callopt.Option) (r *tg.Updates, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.MsgSendMessageV2(ctx, req)
}

func (p *kMsgClient) MsgEditMessageV2(ctx context.Context, req *msg.TLMsgEditMessageV2, callOptions ...callopt.Option) (r *tg.Updates, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.MsgEditMessageV2(ctx, req)
}

func (p *kMsgClient) MsgDeleteMessages(ctx context.Context, req *msg.TLMsgDeleteMessages, callOptions ...callopt.Option) (r *tg.MessagesAffectedMessages, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.MsgDeleteMessages(ctx, req)
}

func (p *kMsgClient) MsgDeleteHistory(ctx context.Context, req *msg.TLMsgDeleteHistory, callOptions ...callopt.Option) (r *tg.MessagesAffectedHistory, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.MsgDeleteHistory(ctx, req)
}

func (p *kMsgClient) MsgDeletePhoneCallHistory(ctx context.Context, req *msg.TLMsgDeletePhoneCallHistory, callOptions ...callopt.Option) (r *tg.MessagesAffectedFoundMessages, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.MsgDeletePhoneCallHistory(ctx, req)
}

func (p *kMsgClient) MsgDeleteChatHistory(ctx context.Context, req *msg.TLMsgDeleteChatHistory, callOptions ...callopt.Option) (r *tg.Bool, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.MsgDeleteChatHistory(ctx, req)
}

func (p *kMsgClient) MsgReadHistory(ctx context.Context, req *msg.TLMsgReadHistory, callOptions ...callopt.Option) (r *tg.MessagesAffectedMessages, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.MsgReadHistory(ctx, req)
}

func (p *kMsgClient) MsgReadHistoryV2(ctx context.Context, req *msg.TLMsgReadHistoryV2, callOptions ...callopt.Option) (r *tg.MessagesAffectedMessages, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.MsgReadHistoryV2(ctx, req)
}

func (p *kMsgClient) MsgUpdatePinnedMessage(ctx context.Context, req *msg.TLMsgUpdatePinnedMessage, callOptions ...callopt.Option) (r *tg.Updates, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.MsgUpdatePinnedMessage(ctx, req)
}

func (p *kMsgClient) MsgUnpinAllMessages(ctx context.Context, req *msg.TLMsgUnpinAllMessages, callOptions ...callopt.Option) (r *tg.MessagesAffectedHistory, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.MsgUnpinAllMessages(ctx, req)
}
