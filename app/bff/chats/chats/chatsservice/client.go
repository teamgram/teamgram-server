/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2024-present,  Teamgram Authors.
 *  All rights reserved.
 *
 * Author: Benqi (wubenqi@gmail.com)
 */

package chatsservice

import (
	"context"

	"github.com/teamgram/proto/v2/tg"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/callopt"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	MessagesGetChats(ctx context.Context, req *tg.TLMessagesGetChats, callOptions ...callopt.Option) (r *tg.MessagesChats, err error)
	MessagesGetFullChat(ctx context.Context, req *tg.TLMessagesGetFullChat, callOptions ...callopt.Option) (r *tg.MessagesChatFull, err error)
	MessagesEditChatTitle(ctx context.Context, req *tg.TLMessagesEditChatTitle, callOptions ...callopt.Option) (r *tg.Updates, err error)
	MessagesEditChatPhoto(ctx context.Context, req *tg.TLMessagesEditChatPhoto, callOptions ...callopt.Option) (r *tg.Updates, err error)
	MessagesAddChatUser(ctx context.Context, req *tg.TLMessagesAddChatUser, callOptions ...callopt.Option) (r *tg.MessagesInvitedUsers, err error)
	MessagesDeleteChatUser(ctx context.Context, req *tg.TLMessagesDeleteChatUser, callOptions ...callopt.Option) (r *tg.Updates, err error)
	MessagesCreateChat(ctx context.Context, req *tg.TLMessagesCreateChat, callOptions ...callopt.Option) (r *tg.MessagesInvitedUsers, err error)
	MessagesEditChatAdmin(ctx context.Context, req *tg.TLMessagesEditChatAdmin, callOptions ...callopt.Option) (r *tg.Bool, err error)
	MessagesMigrateChat(ctx context.Context, req *tg.TLMessagesMigrateChat, callOptions ...callopt.Option) (r *tg.Updates, err error)
	MessagesGetCommonChats(ctx context.Context, req *tg.TLMessagesGetCommonChats, callOptions ...callopt.Option) (r *tg.MessagesChats, err error)
	MessagesEditChatAbout(ctx context.Context, req *tg.TLMessagesEditChatAbout, callOptions ...callopt.Option) (r *tg.Bool, err error)
	MessagesEditChatDefaultBannedRights(ctx context.Context, req *tg.TLMessagesEditChatDefaultBannedRights, callOptions ...callopt.Option) (r *tg.Updates, err error)
	MessagesDeleteChat(ctx context.Context, req *tg.TLMessagesDeleteChat, callOptions ...callopt.Option) (r *tg.Bool, err error)
	MessagesGetMessageReadParticipants(ctx context.Context, req *tg.TLMessagesGetMessageReadParticipants, callOptions ...callopt.Option) (r *tg.VectorReadParticipantDate, err error)
	ChannelsConvertToGigagroup(ctx context.Context, req *tg.TLChannelsConvertToGigagroup, callOptions ...callopt.Option) (r *tg.Updates, err error)
	ChannelsSetEmojiStickers(ctx context.Context, req *tg.TLChannelsSetEmojiStickers, callOptions ...callopt.Option) (r *tg.Bool, err error)
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
	return &kChatsClient{
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

type kChatsClient struct {
	*kClient
}

func NewRPCChatsClient(cli client.Client) Client {
	return &kChatsClient{
		kClient: newServiceClient(cli),
	}
}

func (p *kChatsClient) MessagesGetChats(ctx context.Context, req *tg.TLMessagesGetChats, callOptions ...callopt.Option) (r *tg.MessagesChats, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.MessagesGetChats(ctx, req)
}

func (p *kChatsClient) MessagesGetFullChat(ctx context.Context, req *tg.TLMessagesGetFullChat, callOptions ...callopt.Option) (r *tg.MessagesChatFull, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.MessagesGetFullChat(ctx, req)
}

func (p *kChatsClient) MessagesEditChatTitle(ctx context.Context, req *tg.TLMessagesEditChatTitle, callOptions ...callopt.Option) (r *tg.Updates, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.MessagesEditChatTitle(ctx, req)
}

func (p *kChatsClient) MessagesEditChatPhoto(ctx context.Context, req *tg.TLMessagesEditChatPhoto, callOptions ...callopt.Option) (r *tg.Updates, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.MessagesEditChatPhoto(ctx, req)
}

func (p *kChatsClient) MessagesAddChatUser(ctx context.Context, req *tg.TLMessagesAddChatUser, callOptions ...callopt.Option) (r *tg.MessagesInvitedUsers, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.MessagesAddChatUser(ctx, req)
}

func (p *kChatsClient) MessagesDeleteChatUser(ctx context.Context, req *tg.TLMessagesDeleteChatUser, callOptions ...callopt.Option) (r *tg.Updates, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.MessagesDeleteChatUser(ctx, req)
}

func (p *kChatsClient) MessagesCreateChat(ctx context.Context, req *tg.TLMessagesCreateChat, callOptions ...callopt.Option) (r *tg.MessagesInvitedUsers, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.MessagesCreateChat(ctx, req)
}

func (p *kChatsClient) MessagesEditChatAdmin(ctx context.Context, req *tg.TLMessagesEditChatAdmin, callOptions ...callopt.Option) (r *tg.Bool, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.MessagesEditChatAdmin(ctx, req)
}

func (p *kChatsClient) MessagesMigrateChat(ctx context.Context, req *tg.TLMessagesMigrateChat, callOptions ...callopt.Option) (r *tg.Updates, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.MessagesMigrateChat(ctx, req)
}

func (p *kChatsClient) MessagesGetCommonChats(ctx context.Context, req *tg.TLMessagesGetCommonChats, callOptions ...callopt.Option) (r *tg.MessagesChats, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.MessagesGetCommonChats(ctx, req)
}

func (p *kChatsClient) MessagesEditChatAbout(ctx context.Context, req *tg.TLMessagesEditChatAbout, callOptions ...callopt.Option) (r *tg.Bool, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.MessagesEditChatAbout(ctx, req)
}

func (p *kChatsClient) MessagesEditChatDefaultBannedRights(ctx context.Context, req *tg.TLMessagesEditChatDefaultBannedRights, callOptions ...callopt.Option) (r *tg.Updates, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.MessagesEditChatDefaultBannedRights(ctx, req)
}

func (p *kChatsClient) MessagesDeleteChat(ctx context.Context, req *tg.TLMessagesDeleteChat, callOptions ...callopt.Option) (r *tg.Bool, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.MessagesDeleteChat(ctx, req)
}

func (p *kChatsClient) MessagesGetMessageReadParticipants(ctx context.Context, req *tg.TLMessagesGetMessageReadParticipants, callOptions ...callopt.Option) (r *tg.VectorReadParticipantDate, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.MessagesGetMessageReadParticipants(ctx, req)
}

func (p *kChatsClient) ChannelsConvertToGigagroup(ctx context.Context, req *tg.TLChannelsConvertToGigagroup, callOptions ...callopt.Option) (r *tg.Updates, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.ChannelsConvertToGigagroup(ctx, req)
}

func (p *kChatsClient) ChannelsSetEmojiStickers(ctx context.Context, req *tg.TLChannelsSetEmojiStickers, callOptions ...callopt.Option) (r *tg.Bool, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.ChannelsSetEmojiStickers(ctx, req)
}
