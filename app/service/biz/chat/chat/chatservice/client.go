/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2024-present,  Teamgram Authors.
 *  All rights reserved.
 *
 * Author: Benqi (wubenqi@gmail.com)
 */

package chatservice

import (
	"context"

	"github.com/teamgram/proto/v2/tg"
	"github.com/teamgram/teamgram-server/v2/app/service/biz/chat/chat"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/callopt"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	ChatGetMutableChat(ctx context.Context, req *chat.TLChatGetMutableChat, callOptions ...callopt.Option) (r *tg.MutableChat, err error)
	ChatGetChatListByIdList(ctx context.Context, req *chat.TLChatGetChatListByIdList, callOptions ...callopt.Option) (r *chat.VectorMutableChat, err error)
	ChatGetChatBySelfId(ctx context.Context, req *chat.TLChatGetChatBySelfId, callOptions ...callopt.Option) (r *tg.MutableChat, err error)
	ChatCreateChat2(ctx context.Context, req *chat.TLChatCreateChat2, callOptions ...callopt.Option) (r *tg.MutableChat, err error)
	ChatDeleteChat(ctx context.Context, req *chat.TLChatDeleteChat, callOptions ...callopt.Option) (r *tg.MutableChat, err error)
	ChatDeleteChatUser(ctx context.Context, req *chat.TLChatDeleteChatUser, callOptions ...callopt.Option) (r *tg.MutableChat, err error)
	ChatEditChatTitle(ctx context.Context, req *chat.TLChatEditChatTitle, callOptions ...callopt.Option) (r *tg.MutableChat, err error)
	ChatEditChatAbout(ctx context.Context, req *chat.TLChatEditChatAbout, callOptions ...callopt.Option) (r *tg.MutableChat, err error)
	ChatEditChatPhoto(ctx context.Context, req *chat.TLChatEditChatPhoto, callOptions ...callopt.Option) (r *tg.MutableChat, err error)
	ChatEditChatAdmin(ctx context.Context, req *chat.TLChatEditChatAdmin, callOptions ...callopt.Option) (r *tg.MutableChat, err error)
	ChatEditChatDefaultBannedRights(ctx context.Context, req *chat.TLChatEditChatDefaultBannedRights, callOptions ...callopt.Option) (r *tg.MutableChat, err error)
	ChatAddChatUser(ctx context.Context, req *chat.TLChatAddChatUser, callOptions ...callopt.Option) (r *tg.MutableChat, err error)
	ChatGetMutableChatByLink(ctx context.Context, req *chat.TLChatGetMutableChatByLink, callOptions ...callopt.Option) (r *tg.MutableChat, err error)
	ChatToggleNoForwards(ctx context.Context, req *chat.TLChatToggleNoForwards, callOptions ...callopt.Option) (r *tg.MutableChat, err error)
	ChatMigratedToChannel(ctx context.Context, req *chat.TLChatMigratedToChannel, callOptions ...callopt.Option) (r *tg.Bool, err error)
	ChatGetChatParticipantIdList(ctx context.Context, req *chat.TLChatGetChatParticipantIdList, callOptions ...callopt.Option) (r *chat.VectorLong, err error)
	ChatGetUsersChatIdList(ctx context.Context, req *chat.TLChatGetUsersChatIdList, callOptions ...callopt.Option) (r *chat.VectorUserChatIdList, err error)
	ChatGetMyChatList(ctx context.Context, req *chat.TLChatGetMyChatList, callOptions ...callopt.Option) (r *chat.VectorMutableChat, err error)
	ChatExportChatInvite(ctx context.Context, req *chat.TLChatExportChatInvite, callOptions ...callopt.Option) (r *tg.ExportedChatInvite, err error)
	ChatGetAdminsWithInvites(ctx context.Context, req *chat.TLChatGetAdminsWithInvites, callOptions ...callopt.Option) (r *chat.VectorChatAdminWithInvites, err error)
	ChatGetExportedChatInvite(ctx context.Context, req *chat.TLChatGetExportedChatInvite, callOptions ...callopt.Option) (r *tg.ExportedChatInvite, err error)
	ChatGetExportedChatInvites(ctx context.Context, req *chat.TLChatGetExportedChatInvites, callOptions ...callopt.Option) (r *chat.VectorExportedChatInvite, err error)
	ChatCheckChatInvite(ctx context.Context, req *chat.TLChatCheckChatInvite, callOptions ...callopt.Option) (r *chat.ChatInviteExt, err error)
	ChatImportChatInvite(ctx context.Context, req *chat.TLChatImportChatInvite, callOptions ...callopt.Option) (r *tg.MutableChat, err error)
	ChatGetChatInviteImporters(ctx context.Context, req *chat.TLChatGetChatInviteImporters, callOptions ...callopt.Option) (r *chat.VectorChatInviteImporter, err error)
	ChatDeleteExportedChatInvite(ctx context.Context, req *chat.TLChatDeleteExportedChatInvite, callOptions ...callopt.Option) (r *tg.Bool, err error)
	ChatDeleteRevokedExportedChatInvites(ctx context.Context, req *chat.TLChatDeleteRevokedExportedChatInvites, callOptions ...callopt.Option) (r *tg.Bool, err error)
	ChatEditExportedChatInvite(ctx context.Context, req *chat.TLChatEditExportedChatInvite, callOptions ...callopt.Option) (r *chat.VectorExportedChatInvite, err error)
	ChatSetChatAvailableReactions(ctx context.Context, req *chat.TLChatSetChatAvailableReactions, callOptions ...callopt.Option) (r *tg.MutableChat, err error)
	ChatSetHistoryTTL(ctx context.Context, req *chat.TLChatSetHistoryTTL, callOptions ...callopt.Option) (r *tg.MutableChat, err error)
	ChatSearch(ctx context.Context, req *chat.TLChatSearch, callOptions ...callopt.Option) (r *chat.VectorMutableChat, err error)
	ChatGetRecentChatInviteRequesters(ctx context.Context, req *chat.TLChatGetRecentChatInviteRequesters, callOptions ...callopt.Option) (r *chat.RecentChatInviteRequesters, err error)
	ChatHideChatJoinRequests(ctx context.Context, req *chat.TLChatHideChatJoinRequests, callOptions ...callopt.Option) (r *chat.RecentChatInviteRequesters, err error)
	ChatImportChatInvite2(ctx context.Context, req *chat.TLChatImportChatInvite2, callOptions ...callopt.Option) (r *chat.ChatInviteImported, err error)
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
	return &kChatClient{
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

type kChatClient struct {
	*kClient
}

func NewRPCChatClient(cli client.Client) Client {
	return &kChatClient{
		kClient: newServiceClient(cli),
	}
}

func (p *kChatClient) ChatGetMutableChat(ctx context.Context, req *chat.TLChatGetMutableChat, callOptions ...callopt.Option) (r *tg.MutableChat, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.ChatGetMutableChat(ctx, req)
}

func (p *kChatClient) ChatGetChatListByIdList(ctx context.Context, req *chat.TLChatGetChatListByIdList, callOptions ...callopt.Option) (r *chat.VectorMutableChat, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.ChatGetChatListByIdList(ctx, req)
}

func (p *kChatClient) ChatGetChatBySelfId(ctx context.Context, req *chat.TLChatGetChatBySelfId, callOptions ...callopt.Option) (r *tg.MutableChat, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.ChatGetChatBySelfId(ctx, req)
}

func (p *kChatClient) ChatCreateChat2(ctx context.Context, req *chat.TLChatCreateChat2, callOptions ...callopt.Option) (r *tg.MutableChat, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.ChatCreateChat2(ctx, req)
}

func (p *kChatClient) ChatDeleteChat(ctx context.Context, req *chat.TLChatDeleteChat, callOptions ...callopt.Option) (r *tg.MutableChat, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.ChatDeleteChat(ctx, req)
}

func (p *kChatClient) ChatDeleteChatUser(ctx context.Context, req *chat.TLChatDeleteChatUser, callOptions ...callopt.Option) (r *tg.MutableChat, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.ChatDeleteChatUser(ctx, req)
}

func (p *kChatClient) ChatEditChatTitle(ctx context.Context, req *chat.TLChatEditChatTitle, callOptions ...callopt.Option) (r *tg.MutableChat, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.ChatEditChatTitle(ctx, req)
}

func (p *kChatClient) ChatEditChatAbout(ctx context.Context, req *chat.TLChatEditChatAbout, callOptions ...callopt.Option) (r *tg.MutableChat, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.ChatEditChatAbout(ctx, req)
}

func (p *kChatClient) ChatEditChatPhoto(ctx context.Context, req *chat.TLChatEditChatPhoto, callOptions ...callopt.Option) (r *tg.MutableChat, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.ChatEditChatPhoto(ctx, req)
}

func (p *kChatClient) ChatEditChatAdmin(ctx context.Context, req *chat.TLChatEditChatAdmin, callOptions ...callopt.Option) (r *tg.MutableChat, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.ChatEditChatAdmin(ctx, req)
}

func (p *kChatClient) ChatEditChatDefaultBannedRights(ctx context.Context, req *chat.TLChatEditChatDefaultBannedRights, callOptions ...callopt.Option) (r *tg.MutableChat, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.ChatEditChatDefaultBannedRights(ctx, req)
}

func (p *kChatClient) ChatAddChatUser(ctx context.Context, req *chat.TLChatAddChatUser, callOptions ...callopt.Option) (r *tg.MutableChat, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.ChatAddChatUser(ctx, req)
}

func (p *kChatClient) ChatGetMutableChatByLink(ctx context.Context, req *chat.TLChatGetMutableChatByLink, callOptions ...callopt.Option) (r *tg.MutableChat, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.ChatGetMutableChatByLink(ctx, req)
}

func (p *kChatClient) ChatToggleNoForwards(ctx context.Context, req *chat.TLChatToggleNoForwards, callOptions ...callopt.Option) (r *tg.MutableChat, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.ChatToggleNoForwards(ctx, req)
}

func (p *kChatClient) ChatMigratedToChannel(ctx context.Context, req *chat.TLChatMigratedToChannel, callOptions ...callopt.Option) (r *tg.Bool, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.ChatMigratedToChannel(ctx, req)
}

func (p *kChatClient) ChatGetChatParticipantIdList(ctx context.Context, req *chat.TLChatGetChatParticipantIdList, callOptions ...callopt.Option) (r *chat.VectorLong, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.ChatGetChatParticipantIdList(ctx, req)
}

func (p *kChatClient) ChatGetUsersChatIdList(ctx context.Context, req *chat.TLChatGetUsersChatIdList, callOptions ...callopt.Option) (r *chat.VectorUserChatIdList, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.ChatGetUsersChatIdList(ctx, req)
}

func (p *kChatClient) ChatGetMyChatList(ctx context.Context, req *chat.TLChatGetMyChatList, callOptions ...callopt.Option) (r *chat.VectorMutableChat, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.ChatGetMyChatList(ctx, req)
}

func (p *kChatClient) ChatExportChatInvite(ctx context.Context, req *chat.TLChatExportChatInvite, callOptions ...callopt.Option) (r *tg.ExportedChatInvite, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.ChatExportChatInvite(ctx, req)
}

func (p *kChatClient) ChatGetAdminsWithInvites(ctx context.Context, req *chat.TLChatGetAdminsWithInvites, callOptions ...callopt.Option) (r *chat.VectorChatAdminWithInvites, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.ChatGetAdminsWithInvites(ctx, req)
}

func (p *kChatClient) ChatGetExportedChatInvite(ctx context.Context, req *chat.TLChatGetExportedChatInvite, callOptions ...callopt.Option) (r *tg.ExportedChatInvite, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.ChatGetExportedChatInvite(ctx, req)
}

func (p *kChatClient) ChatGetExportedChatInvites(ctx context.Context, req *chat.TLChatGetExportedChatInvites, callOptions ...callopt.Option) (r *chat.VectorExportedChatInvite, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.ChatGetExportedChatInvites(ctx, req)
}

func (p *kChatClient) ChatCheckChatInvite(ctx context.Context, req *chat.TLChatCheckChatInvite, callOptions ...callopt.Option) (r *chat.ChatInviteExt, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.ChatCheckChatInvite(ctx, req)
}

func (p *kChatClient) ChatImportChatInvite(ctx context.Context, req *chat.TLChatImportChatInvite, callOptions ...callopt.Option) (r *tg.MutableChat, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.ChatImportChatInvite(ctx, req)
}

func (p *kChatClient) ChatGetChatInviteImporters(ctx context.Context, req *chat.TLChatGetChatInviteImporters, callOptions ...callopt.Option) (r *chat.VectorChatInviteImporter, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.ChatGetChatInviteImporters(ctx, req)
}

func (p *kChatClient) ChatDeleteExportedChatInvite(ctx context.Context, req *chat.TLChatDeleteExportedChatInvite, callOptions ...callopt.Option) (r *tg.Bool, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.ChatDeleteExportedChatInvite(ctx, req)
}

func (p *kChatClient) ChatDeleteRevokedExportedChatInvites(ctx context.Context, req *chat.TLChatDeleteRevokedExportedChatInvites, callOptions ...callopt.Option) (r *tg.Bool, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.ChatDeleteRevokedExportedChatInvites(ctx, req)
}

func (p *kChatClient) ChatEditExportedChatInvite(ctx context.Context, req *chat.TLChatEditExportedChatInvite, callOptions ...callopt.Option) (r *chat.VectorExportedChatInvite, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.ChatEditExportedChatInvite(ctx, req)
}

func (p *kChatClient) ChatSetChatAvailableReactions(ctx context.Context, req *chat.TLChatSetChatAvailableReactions, callOptions ...callopt.Option) (r *tg.MutableChat, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.ChatSetChatAvailableReactions(ctx, req)
}

func (p *kChatClient) ChatSetHistoryTTL(ctx context.Context, req *chat.TLChatSetHistoryTTL, callOptions ...callopt.Option) (r *tg.MutableChat, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.ChatSetHistoryTTL(ctx, req)
}

func (p *kChatClient) ChatSearch(ctx context.Context, req *chat.TLChatSearch, callOptions ...callopt.Option) (r *chat.VectorMutableChat, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.ChatSearch(ctx, req)
}

func (p *kChatClient) ChatGetRecentChatInviteRequesters(ctx context.Context, req *chat.TLChatGetRecentChatInviteRequesters, callOptions ...callopt.Option) (r *chat.RecentChatInviteRequesters, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.ChatGetRecentChatInviteRequesters(ctx, req)
}

func (p *kChatClient) ChatHideChatJoinRequests(ctx context.Context, req *chat.TLChatHideChatJoinRequests, callOptions ...callopt.Option) (r *chat.RecentChatInviteRequesters, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.ChatHideChatJoinRequests(ctx, req)
}

func (p *kChatClient) ChatImportChatInvite2(ctx context.Context, req *chat.TLChatImportChatInvite2, callOptions ...callopt.Option) (r *chat.ChatInviteImported, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.ChatImportChatInvite2(ctx, req)
}
