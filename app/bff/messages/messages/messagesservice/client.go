/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2024-present,  Teamgram Authors.
 *  All rights reserved.
 *
 * Author: Benqi (wubenqi@gmail.com)
 */

package messagesservice

import (
	"context"

	"github.com/teamgram/proto/v2/tg"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/callopt"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	MessagesGetMessages(ctx context.Context, req *tg.TLMessagesGetMessages, callOptions ...callopt.Option) (r *tg.MessagesMessages, err error)
	MessagesGetHistory(ctx context.Context, req *tg.TLMessagesGetHistory, callOptions ...callopt.Option) (r *tg.MessagesMessages, err error)
	MessagesSearch(ctx context.Context, req *tg.TLMessagesSearch, callOptions ...callopt.Option) (r *tg.MessagesMessages, err error)
	MessagesReadHistory(ctx context.Context, req *tg.TLMessagesReadHistory, callOptions ...callopt.Option) (r *tg.MessagesAffectedMessages, err error)
	MessagesDeleteHistory(ctx context.Context, req *tg.TLMessagesDeleteHistory, callOptions ...callopt.Option) (r *tg.MessagesAffectedHistory, err error)
	MessagesDeleteMessages(ctx context.Context, req *tg.TLMessagesDeleteMessages, callOptions ...callopt.Option) (r *tg.MessagesAffectedMessages, err error)
	MessagesReceivedMessages(ctx context.Context, req *tg.TLMessagesReceivedMessages, callOptions ...callopt.Option) (r *tg.VectorReceivedNotifyMessage, err error)
	MessagesSendMessage(ctx context.Context, req *tg.TLMessagesSendMessage, callOptions ...callopt.Option) (r *tg.Updates, err error)
	MessagesSendMedia(ctx context.Context, req *tg.TLMessagesSendMedia, callOptions ...callopt.Option) (r *tg.Updates, err error)
	MessagesForwardMessages(ctx context.Context, req *tg.TLMessagesForwardMessages, callOptions ...callopt.Option) (r *tg.Updates, err error)
	MessagesReadMessageContents(ctx context.Context, req *tg.TLMessagesReadMessageContents, callOptions ...callopt.Option) (r *tg.MessagesAffectedMessages, err error)
	MessagesGetMessagesViews(ctx context.Context, req *tg.TLMessagesGetMessagesViews, callOptions ...callopt.Option) (r *tg.MessagesMessageViews, err error)
	MessagesSearchGlobal(ctx context.Context, req *tg.TLMessagesSearchGlobal, callOptions ...callopt.Option) (r *tg.MessagesMessages, err error)
	MessagesGetMessageEditData(ctx context.Context, req *tg.TLMessagesGetMessageEditData, callOptions ...callopt.Option) (r *tg.MessagesMessageEditData, err error)
	MessagesEditMessage(ctx context.Context, req *tg.TLMessagesEditMessage, callOptions ...callopt.Option) (r *tg.Updates, err error)
	MessagesGetUnreadMentions(ctx context.Context, req *tg.TLMessagesGetUnreadMentions, callOptions ...callopt.Option) (r *tg.MessagesMessages, err error)
	MessagesReadMentions(ctx context.Context, req *tg.TLMessagesReadMentions, callOptions ...callopt.Option) (r *tg.MessagesAffectedHistory, err error)
	MessagesGetRecentLocations(ctx context.Context, req *tg.TLMessagesGetRecentLocations, callOptions ...callopt.Option) (r *tg.MessagesMessages, err error)
	MessagesSendMultiMedia(ctx context.Context, req *tg.TLMessagesSendMultiMedia, callOptions ...callopt.Option) (r *tg.Updates, err error)
	MessagesUpdatePinnedMessage(ctx context.Context, req *tg.TLMessagesUpdatePinnedMessage, callOptions ...callopt.Option) (r *tg.Updates, err error)
	MessagesGetSearchCounters(ctx context.Context, req *tg.TLMessagesGetSearchCounters, callOptions ...callopt.Option) (r *tg.VectorMessagesSearchCounter, err error)
	MessagesUnpinAllMessages(ctx context.Context, req *tg.TLMessagesUnpinAllMessages, callOptions ...callopt.Option) (r *tg.MessagesAffectedHistory, err error)
	MessagesGetSearchResultsCalendar(ctx context.Context, req *tg.TLMessagesGetSearchResultsCalendar, callOptions ...callopt.Option) (r *tg.MessagesSearchResultsCalendar, err error)
	MessagesGetSearchResultsPositions(ctx context.Context, req *tg.TLMessagesGetSearchResultsPositions, callOptions ...callopt.Option) (r *tg.MessagesSearchResultsPositions, err error)
	MessagesToggleNoForwards(ctx context.Context, req *tg.TLMessagesToggleNoForwards, callOptions ...callopt.Option) (r *tg.Updates, err error)
	MessagesSaveDefaultSendAs(ctx context.Context, req *tg.TLMessagesSaveDefaultSendAs, callOptions ...callopt.Option) (r *tg.Bool, err error)
	MessagesSearchSentMedia(ctx context.Context, req *tg.TLMessagesSearchSentMedia, callOptions ...callopt.Option) (r *tg.MessagesMessages, err error)
	MessagesGetOutboxReadDate(ctx context.Context, req *tg.TLMessagesGetOutboxReadDate, callOptions ...callopt.Option) (r *tg.OutboxReadDate, err error)
	ChannelsGetSendAs(ctx context.Context, req *tg.TLChannelsGetSendAs, callOptions ...callopt.Option) (r *tg.ChannelsSendAsPeers, err error)
	ChannelsSearchPosts(ctx context.Context, req *tg.TLChannelsSearchPosts, callOptions ...callopt.Option) (r *tg.MessagesMessages, err error)
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

func (p *kMessagesClient) MessagesGetMessages(ctx context.Context, req *tg.TLMessagesGetMessages, callOptions ...callopt.Option) (r *tg.MessagesMessages, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.MessagesGetMessages(ctx, req)
}

func (p *kMessagesClient) MessagesGetHistory(ctx context.Context, req *tg.TLMessagesGetHistory, callOptions ...callopt.Option) (r *tg.MessagesMessages, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.MessagesGetHistory(ctx, req)
}

func (p *kMessagesClient) MessagesSearch(ctx context.Context, req *tg.TLMessagesSearch, callOptions ...callopt.Option) (r *tg.MessagesMessages, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.MessagesSearch(ctx, req)
}

func (p *kMessagesClient) MessagesReadHistory(ctx context.Context, req *tg.TLMessagesReadHistory, callOptions ...callopt.Option) (r *tg.MessagesAffectedMessages, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.MessagesReadHistory(ctx, req)
}

func (p *kMessagesClient) MessagesDeleteHistory(ctx context.Context, req *tg.TLMessagesDeleteHistory, callOptions ...callopt.Option) (r *tg.MessagesAffectedHistory, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.MessagesDeleteHistory(ctx, req)
}

func (p *kMessagesClient) MessagesDeleteMessages(ctx context.Context, req *tg.TLMessagesDeleteMessages, callOptions ...callopt.Option) (r *tg.MessagesAffectedMessages, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.MessagesDeleteMessages(ctx, req)
}

func (p *kMessagesClient) MessagesReceivedMessages(ctx context.Context, req *tg.TLMessagesReceivedMessages, callOptions ...callopt.Option) (r *tg.VectorReceivedNotifyMessage, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.MessagesReceivedMessages(ctx, req)
}

func (p *kMessagesClient) MessagesSendMessage(ctx context.Context, req *tg.TLMessagesSendMessage, callOptions ...callopt.Option) (r *tg.Updates, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.MessagesSendMessage(ctx, req)
}

func (p *kMessagesClient) MessagesSendMedia(ctx context.Context, req *tg.TLMessagesSendMedia, callOptions ...callopt.Option) (r *tg.Updates, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.MessagesSendMedia(ctx, req)
}

func (p *kMessagesClient) MessagesForwardMessages(ctx context.Context, req *tg.TLMessagesForwardMessages, callOptions ...callopt.Option) (r *tg.Updates, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.MessagesForwardMessages(ctx, req)
}

func (p *kMessagesClient) MessagesReadMessageContents(ctx context.Context, req *tg.TLMessagesReadMessageContents, callOptions ...callopt.Option) (r *tg.MessagesAffectedMessages, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.MessagesReadMessageContents(ctx, req)
}

func (p *kMessagesClient) MessagesGetMessagesViews(ctx context.Context, req *tg.TLMessagesGetMessagesViews, callOptions ...callopt.Option) (r *tg.MessagesMessageViews, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.MessagesGetMessagesViews(ctx, req)
}

func (p *kMessagesClient) MessagesSearchGlobal(ctx context.Context, req *tg.TLMessagesSearchGlobal, callOptions ...callopt.Option) (r *tg.MessagesMessages, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.MessagesSearchGlobal(ctx, req)
}

func (p *kMessagesClient) MessagesGetMessageEditData(ctx context.Context, req *tg.TLMessagesGetMessageEditData, callOptions ...callopt.Option) (r *tg.MessagesMessageEditData, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.MessagesGetMessageEditData(ctx, req)
}

func (p *kMessagesClient) MessagesEditMessage(ctx context.Context, req *tg.TLMessagesEditMessage, callOptions ...callopt.Option) (r *tg.Updates, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.MessagesEditMessage(ctx, req)
}

func (p *kMessagesClient) MessagesGetUnreadMentions(ctx context.Context, req *tg.TLMessagesGetUnreadMentions, callOptions ...callopt.Option) (r *tg.MessagesMessages, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.MessagesGetUnreadMentions(ctx, req)
}

func (p *kMessagesClient) MessagesReadMentions(ctx context.Context, req *tg.TLMessagesReadMentions, callOptions ...callopt.Option) (r *tg.MessagesAffectedHistory, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.MessagesReadMentions(ctx, req)
}

func (p *kMessagesClient) MessagesGetRecentLocations(ctx context.Context, req *tg.TLMessagesGetRecentLocations, callOptions ...callopt.Option) (r *tg.MessagesMessages, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.MessagesGetRecentLocations(ctx, req)
}

func (p *kMessagesClient) MessagesSendMultiMedia(ctx context.Context, req *tg.TLMessagesSendMultiMedia, callOptions ...callopt.Option) (r *tg.Updates, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.MessagesSendMultiMedia(ctx, req)
}

func (p *kMessagesClient) MessagesUpdatePinnedMessage(ctx context.Context, req *tg.TLMessagesUpdatePinnedMessage, callOptions ...callopt.Option) (r *tg.Updates, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.MessagesUpdatePinnedMessage(ctx, req)
}

func (p *kMessagesClient) MessagesGetSearchCounters(ctx context.Context, req *tg.TLMessagesGetSearchCounters, callOptions ...callopt.Option) (r *tg.VectorMessagesSearchCounter, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.MessagesGetSearchCounters(ctx, req)
}

func (p *kMessagesClient) MessagesUnpinAllMessages(ctx context.Context, req *tg.TLMessagesUnpinAllMessages, callOptions ...callopt.Option) (r *tg.MessagesAffectedHistory, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.MessagesUnpinAllMessages(ctx, req)
}

func (p *kMessagesClient) MessagesGetSearchResultsCalendar(ctx context.Context, req *tg.TLMessagesGetSearchResultsCalendar, callOptions ...callopt.Option) (r *tg.MessagesSearchResultsCalendar, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.MessagesGetSearchResultsCalendar(ctx, req)
}

func (p *kMessagesClient) MessagesGetSearchResultsPositions(ctx context.Context, req *tg.TLMessagesGetSearchResultsPositions, callOptions ...callopt.Option) (r *tg.MessagesSearchResultsPositions, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.MessagesGetSearchResultsPositions(ctx, req)
}

func (p *kMessagesClient) MessagesToggleNoForwards(ctx context.Context, req *tg.TLMessagesToggleNoForwards, callOptions ...callopt.Option) (r *tg.Updates, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.MessagesToggleNoForwards(ctx, req)
}

func (p *kMessagesClient) MessagesSaveDefaultSendAs(ctx context.Context, req *tg.TLMessagesSaveDefaultSendAs, callOptions ...callopt.Option) (r *tg.Bool, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.MessagesSaveDefaultSendAs(ctx, req)
}

func (p *kMessagesClient) MessagesSearchSentMedia(ctx context.Context, req *tg.TLMessagesSearchSentMedia, callOptions ...callopt.Option) (r *tg.MessagesMessages, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.MessagesSearchSentMedia(ctx, req)
}

func (p *kMessagesClient) MessagesGetOutboxReadDate(ctx context.Context, req *tg.TLMessagesGetOutboxReadDate, callOptions ...callopt.Option) (r *tg.OutboxReadDate, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.MessagesGetOutboxReadDate(ctx, req)
}

func (p *kMessagesClient) ChannelsGetSendAs(ctx context.Context, req *tg.TLChannelsGetSendAs, callOptions ...callopt.Option) (r *tg.ChannelsSendAsPeers, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.ChannelsGetSendAs(ctx, req)
}

func (p *kMessagesClient) ChannelsSearchPosts(ctx context.Context, req *tg.TLChannelsSearchPosts, callOptions ...callopt.Option) (r *tg.MessagesMessages, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.ChannelsSearchPosts(ctx, req)
}
