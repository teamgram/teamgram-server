/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2024-present,  Teamgram Authors.
 *  All rights reserved.
 *
 * Author: Benqi (wubenqi@gmail.com)
 */

package chatinvitesservice

import (
	"context"

	"github.com/teamgram/proto/v2/tg"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/callopt"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	MessagesExportChatInvite(ctx context.Context, req *tg.TLMessagesExportChatInvite, callOptions ...callopt.Option) (r *tg.ExportedChatInvite, err error)
	MessagesCheckChatInvite(ctx context.Context, req *tg.TLMessagesCheckChatInvite, callOptions ...callopt.Option) (r *tg.ChatInvite, err error)
	MessagesImportChatInvite(ctx context.Context, req *tg.TLMessagesImportChatInvite, callOptions ...callopt.Option) (r *tg.Updates, err error)
	MessagesGetExportedChatInvites(ctx context.Context, req *tg.TLMessagesGetExportedChatInvites, callOptions ...callopt.Option) (r *tg.MessagesExportedChatInvites, err error)
	MessagesGetExportedChatInvite(ctx context.Context, req *tg.TLMessagesGetExportedChatInvite, callOptions ...callopt.Option) (r *tg.MessagesExportedChatInvite, err error)
	MessagesEditExportedChatInvite(ctx context.Context, req *tg.TLMessagesEditExportedChatInvite, callOptions ...callopt.Option) (r *tg.MessagesExportedChatInvite, err error)
	MessagesDeleteRevokedExportedChatInvites(ctx context.Context, req *tg.TLMessagesDeleteRevokedExportedChatInvites, callOptions ...callopt.Option) (r *tg.Bool, err error)
	MessagesDeleteExportedChatInvite(ctx context.Context, req *tg.TLMessagesDeleteExportedChatInvite, callOptions ...callopt.Option) (r *tg.Bool, err error)
	MessagesGetAdminsWithInvites(ctx context.Context, req *tg.TLMessagesGetAdminsWithInvites, callOptions ...callopt.Option) (r *tg.MessagesChatAdminsWithInvites, err error)
	MessagesGetChatInviteImporters(ctx context.Context, req *tg.TLMessagesGetChatInviteImporters, callOptions ...callopt.Option) (r *tg.MessagesChatInviteImporters, err error)
	MessagesHideChatJoinRequest(ctx context.Context, req *tg.TLMessagesHideChatJoinRequest, callOptions ...callopt.Option) (r *tg.Updates, err error)
	MessagesHideAllChatJoinRequests(ctx context.Context, req *tg.TLMessagesHideAllChatJoinRequests, callOptions ...callopt.Option) (r *tg.Updates, err error)
	ChannelsToggleJoinToSend(ctx context.Context, req *tg.TLChannelsToggleJoinToSend, callOptions ...callopt.Option) (r *tg.Updates, err error)
	ChannelsToggleJoinRequest(ctx context.Context, req *tg.TLChannelsToggleJoinRequest, callOptions ...callopt.Option) (r *tg.Updates, err error)
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
	return &kChatInvitesClient{
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

type kChatInvitesClient struct {
	*kClient
}

func NewRPCChatInvitesClient(cli client.Client) Client {
	return &kChatInvitesClient{
		kClient: newServiceClient(cli),
	}
}

func (p *kChatInvitesClient) MessagesExportChatInvite(ctx context.Context, req *tg.TLMessagesExportChatInvite, callOptions ...callopt.Option) (r *tg.ExportedChatInvite, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.MessagesExportChatInvite(ctx, req)
}

func (p *kChatInvitesClient) MessagesCheckChatInvite(ctx context.Context, req *tg.TLMessagesCheckChatInvite, callOptions ...callopt.Option) (r *tg.ChatInvite, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.MessagesCheckChatInvite(ctx, req)
}

func (p *kChatInvitesClient) MessagesImportChatInvite(ctx context.Context, req *tg.TLMessagesImportChatInvite, callOptions ...callopt.Option) (r *tg.Updates, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.MessagesImportChatInvite(ctx, req)
}

func (p *kChatInvitesClient) MessagesGetExportedChatInvites(ctx context.Context, req *tg.TLMessagesGetExportedChatInvites, callOptions ...callopt.Option) (r *tg.MessagesExportedChatInvites, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.MessagesGetExportedChatInvites(ctx, req)
}

func (p *kChatInvitesClient) MessagesGetExportedChatInvite(ctx context.Context, req *tg.TLMessagesGetExportedChatInvite, callOptions ...callopt.Option) (r *tg.MessagesExportedChatInvite, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.MessagesGetExportedChatInvite(ctx, req)
}

func (p *kChatInvitesClient) MessagesEditExportedChatInvite(ctx context.Context, req *tg.TLMessagesEditExportedChatInvite, callOptions ...callopt.Option) (r *tg.MessagesExportedChatInvite, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.MessagesEditExportedChatInvite(ctx, req)
}

func (p *kChatInvitesClient) MessagesDeleteRevokedExportedChatInvites(ctx context.Context, req *tg.TLMessagesDeleteRevokedExportedChatInvites, callOptions ...callopt.Option) (r *tg.Bool, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.MessagesDeleteRevokedExportedChatInvites(ctx, req)
}

func (p *kChatInvitesClient) MessagesDeleteExportedChatInvite(ctx context.Context, req *tg.TLMessagesDeleteExportedChatInvite, callOptions ...callopt.Option) (r *tg.Bool, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.MessagesDeleteExportedChatInvite(ctx, req)
}

func (p *kChatInvitesClient) MessagesGetAdminsWithInvites(ctx context.Context, req *tg.TLMessagesGetAdminsWithInvites, callOptions ...callopt.Option) (r *tg.MessagesChatAdminsWithInvites, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.MessagesGetAdminsWithInvites(ctx, req)
}

func (p *kChatInvitesClient) MessagesGetChatInviteImporters(ctx context.Context, req *tg.TLMessagesGetChatInviteImporters, callOptions ...callopt.Option) (r *tg.MessagesChatInviteImporters, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.MessagesGetChatInviteImporters(ctx, req)
}

func (p *kChatInvitesClient) MessagesHideChatJoinRequest(ctx context.Context, req *tg.TLMessagesHideChatJoinRequest, callOptions ...callopt.Option) (r *tg.Updates, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.MessagesHideChatJoinRequest(ctx, req)
}

func (p *kChatInvitesClient) MessagesHideAllChatJoinRequests(ctx context.Context, req *tg.TLMessagesHideAllChatJoinRequests, callOptions ...callopt.Option) (r *tg.Updates, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.MessagesHideAllChatJoinRequests(ctx, req)
}

func (p *kChatInvitesClient) ChannelsToggleJoinToSend(ctx context.Context, req *tg.TLChannelsToggleJoinToSend, callOptions ...callopt.Option) (r *tg.Updates, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.ChannelsToggleJoinToSend(ctx, req)
}

func (p *kChatInvitesClient) ChannelsToggleJoinRequest(ctx context.Context, req *tg.TLChannelsToggleJoinRequest, callOptions ...callopt.Option) (r *tg.Updates, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.ChannelsToggleJoinRequest(ctx, req)
}
