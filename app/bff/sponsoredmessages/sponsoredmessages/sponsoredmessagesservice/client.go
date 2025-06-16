/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2024-present,  Teamgram Authors.
 *  All rights reserved.
 *
 * Author: Benqi (wubenqi@gmail.com)
 */

package sponsoredmessagesservice

import (
	"context"

	"github.com/teamgram/proto/v2/tg"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/callopt"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	AccountToggleSponsoredMessages(ctx context.Context, req *tg.TLAccountToggleSponsoredMessages, callOptions ...callopt.Option) (r *tg.Bool, err error)
	MessagesViewSponsoredMessage(ctx context.Context, req *tg.TLMessagesViewSponsoredMessage, callOptions ...callopt.Option) (r *tg.Bool, err error)
	MessagesClickSponsoredMessage(ctx context.Context, req *tg.TLMessagesClickSponsoredMessage, callOptions ...callopt.Option) (r *tg.Bool, err error)
	MessagesReportSponsoredMessage(ctx context.Context, req *tg.TLMessagesReportSponsoredMessage, callOptions ...callopt.Option) (r *tg.ChannelsSponsoredMessageReportResult, err error)
	MessagesGetSponsoredMessages(ctx context.Context, req *tg.TLMessagesGetSponsoredMessages, callOptions ...callopt.Option) (r *tg.MessagesSponsoredMessages, err error)
	ChannelsRestrictSponsoredMessages(ctx context.Context, req *tg.TLChannelsRestrictSponsoredMessages, callOptions ...callopt.Option) (r *tg.Updates, err error)
	ChannelsViewSponsoredMessage(ctx context.Context, req *tg.TLChannelsViewSponsoredMessage, callOptions ...callopt.Option) (r *tg.Bool, err error)
	ChannelsGetSponsoredMessages(ctx context.Context, req *tg.TLChannelsGetSponsoredMessages, callOptions ...callopt.Option) (r *tg.MessagesSponsoredMessages, err error)
	ChannelsClickSponsoredMessage(ctx context.Context, req *tg.TLChannelsClickSponsoredMessage, callOptions ...callopt.Option) (r *tg.Bool, err error)
	ChannelsReportSponsoredMessage(ctx context.Context, req *tg.TLChannelsReportSponsoredMessage, callOptions ...callopt.Option) (r *tg.ChannelsSponsoredMessageReportResult, err error)
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
	return &kSponsoredMessagesClient{
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

type kSponsoredMessagesClient struct {
	*kClient
}

func NewRPCSponsoredMessagesClient(cli client.Client) Client {
	return &kSponsoredMessagesClient{
		kClient: newServiceClient(cli),
	}
}

func (p *kSponsoredMessagesClient) AccountToggleSponsoredMessages(ctx context.Context, req *tg.TLAccountToggleSponsoredMessages, callOptions ...callopt.Option) (r *tg.Bool, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.AccountToggleSponsoredMessages(ctx, req)
}

func (p *kSponsoredMessagesClient) MessagesViewSponsoredMessage(ctx context.Context, req *tg.TLMessagesViewSponsoredMessage, callOptions ...callopt.Option) (r *tg.Bool, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.MessagesViewSponsoredMessage(ctx, req)
}

func (p *kSponsoredMessagesClient) MessagesClickSponsoredMessage(ctx context.Context, req *tg.TLMessagesClickSponsoredMessage, callOptions ...callopt.Option) (r *tg.Bool, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.MessagesClickSponsoredMessage(ctx, req)
}

func (p *kSponsoredMessagesClient) MessagesReportSponsoredMessage(ctx context.Context, req *tg.TLMessagesReportSponsoredMessage, callOptions ...callopt.Option) (r *tg.ChannelsSponsoredMessageReportResult, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.MessagesReportSponsoredMessage(ctx, req)
}

func (p *kSponsoredMessagesClient) MessagesGetSponsoredMessages(ctx context.Context, req *tg.TLMessagesGetSponsoredMessages, callOptions ...callopt.Option) (r *tg.MessagesSponsoredMessages, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.MessagesGetSponsoredMessages(ctx, req)
}

func (p *kSponsoredMessagesClient) ChannelsRestrictSponsoredMessages(ctx context.Context, req *tg.TLChannelsRestrictSponsoredMessages, callOptions ...callopt.Option) (r *tg.Updates, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.ChannelsRestrictSponsoredMessages(ctx, req)
}

func (p *kSponsoredMessagesClient) ChannelsViewSponsoredMessage(ctx context.Context, req *tg.TLChannelsViewSponsoredMessage, callOptions ...callopt.Option) (r *tg.Bool, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.ChannelsViewSponsoredMessage(ctx, req)
}

func (p *kSponsoredMessagesClient) ChannelsGetSponsoredMessages(ctx context.Context, req *tg.TLChannelsGetSponsoredMessages, callOptions ...callopt.Option) (r *tg.MessagesSponsoredMessages, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.ChannelsGetSponsoredMessages(ctx, req)
}

func (p *kSponsoredMessagesClient) ChannelsClickSponsoredMessage(ctx context.Context, req *tg.TLChannelsClickSponsoredMessage, callOptions ...callopt.Option) (r *tg.Bool, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.ChannelsClickSponsoredMessage(ctx, req)
}

func (p *kSponsoredMessagesClient) ChannelsReportSponsoredMessage(ctx context.Context, req *tg.TLChannelsReportSponsoredMessage, callOptions ...callopt.Option) (r *tg.ChannelsSponsoredMessageReportResult, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.ChannelsReportSponsoredMessage(ctx, req)
}
