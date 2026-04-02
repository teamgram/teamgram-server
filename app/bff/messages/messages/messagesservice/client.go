/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2026-present,  Teamgram Authors.
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package messagesservice

import (
	"context"

	"github.com/teamgram/teamgram-server/v2/pkg/net/kitex/codec"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/callopt"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	MessagesComposeMessageWithAI(ctx context.Context, req *tg.TLMessagesComposeMessageWithAI, callOptions ...callopt.Option) (r *tg.MessagesComposedMessageWithAI, err error)
	MessagesReportReadMetrics(ctx context.Context, req *tg.TLMessagesReportReadMetrics, callOptions ...callopt.Option) (r *tg.Bool, err error)
	MessagesReportMusicListen(ctx context.Context, req *tg.TLMessagesReportMusicListen, callOptions ...callopt.Option) (r *tg.Bool, err error)
	MessagesAddPollAnswer(ctx context.Context, req *tg.TLMessagesAddPollAnswer, callOptions ...callopt.Option) (r *tg.Updates, err error)
	MessagesDeletePollAnswer(ctx context.Context, req *tg.TLMessagesDeletePollAnswer, callOptions ...callopt.Option) (r *tg.Updates, err error)
	MessagesGetUnreadPollVotes(ctx context.Context, req *tg.TLMessagesGetUnreadPollVotes, callOptions ...callopt.Option) (r *tg.MessagesMessages, err error)
	MessagesReadPollVotes(ctx context.Context, req *tg.TLMessagesReadPollVotes, callOptions ...callopt.Option) (r *tg.MessagesAffectedHistory, err error)
}

// Deprecated: prefer the generated app client helper or pkg/net/kitex.NewClient for TL-aware transport setup.
// NewClient creates a client for the service defined in IDL.
func NewClient(destService string, opts ...client.Option) (Client, error) {
	var options []client.Option
	options = append(options, client.WithDestService(destService))
	options = append(options, client.WithCodec(codec.NewZRpcCodec(false)))

	options = append(options, opts...)

	kc, err := client.NewClient(serviceInfoForClient(), options...)
	if err != nil {
		return nil, err
	}
	return &kMessagesClient{
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

type kMessagesClient struct {
	*kClient
}

func NewRPCMessagesClient(cli client.Client) Client {
	return &kMessagesClient{
		kClient: newServiceClient(cli),
	}
}

func (p *kMessagesClient) MessagesComposeMessageWithAI(ctx context.Context, req *tg.TLMessagesComposeMessageWithAI, callOptions ...callopt.Option) (r *tg.MessagesComposedMessageWithAI, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.MessagesComposeMessageWithAI(ctx, req)
}

func (p *kMessagesClient) MessagesReportReadMetrics(ctx context.Context, req *tg.TLMessagesReportReadMetrics, callOptions ...callopt.Option) (r *tg.Bool, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.MessagesReportReadMetrics(ctx, req)
}

func (p *kMessagesClient) MessagesReportMusicListen(ctx context.Context, req *tg.TLMessagesReportMusicListen, callOptions ...callopt.Option) (r *tg.Bool, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.MessagesReportMusicListen(ctx, req)
}

func (p *kMessagesClient) MessagesAddPollAnswer(ctx context.Context, req *tg.TLMessagesAddPollAnswer, callOptions ...callopt.Option) (r *tg.Updates, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.MessagesAddPollAnswer(ctx, req)
}

func (p *kMessagesClient) MessagesDeletePollAnswer(ctx context.Context, req *tg.TLMessagesDeletePollAnswer, callOptions ...callopt.Option) (r *tg.Updates, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.MessagesDeletePollAnswer(ctx, req)
}

func (p *kMessagesClient) MessagesGetUnreadPollVotes(ctx context.Context, req *tg.TLMessagesGetUnreadPollVotes, callOptions ...callopt.Option) (r *tg.MessagesMessages, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.MessagesGetUnreadPollVotes(ctx, req)
}

func (p *kMessagesClient) MessagesReadPollVotes(ctx context.Context, req *tg.TLMessagesReadPollVotes, callOptions ...callopt.Option) (r *tg.MessagesAffectedHistory, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.MessagesReadPollVotes(ctx, req)
}
