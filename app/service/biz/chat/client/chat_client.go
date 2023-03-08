/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright 2022 Teamgram Authors.
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package chat_client

import (
	"context"

	"github.com/teamgram/proto/mtproto"
	"github.com/teamgram/teamgram-server/app/service/biz/chat/chat"

	"github.com/zeromicro/go-zero/zrpc"
)

var _ *mtproto.Bool

type ChatClient interface {
	ChatGetMutableChat(ctx context.Context, in *chat.TLChatGetMutableChat) (*mtproto.MutableChat, error)
	ChatGetChatListByIdList(ctx context.Context, in *chat.TLChatGetChatListByIdList) (*chat.Vector_MutableChat, error)
	ChatGetChatBySelfId(ctx context.Context, in *chat.TLChatGetChatBySelfId) (*mtproto.MutableChat, error)
	ChatCreateChat2(ctx context.Context, in *chat.TLChatCreateChat2) (*mtproto.MutableChat, error)
	ChatDeleteChat(ctx context.Context, in *chat.TLChatDeleteChat) (*mtproto.MutableChat, error)
	ChatDeleteChatUser(ctx context.Context, in *chat.TLChatDeleteChatUser) (*mtproto.MutableChat, error)
	ChatEditChatTitle(ctx context.Context, in *chat.TLChatEditChatTitle) (*mtproto.MutableChat, error)
	ChatEditChatAbout(ctx context.Context, in *chat.TLChatEditChatAbout) (*mtproto.MutableChat, error)
	ChatEditChatPhoto(ctx context.Context, in *chat.TLChatEditChatPhoto) (*mtproto.MutableChat, error)
	ChatEditChatAdmin(ctx context.Context, in *chat.TLChatEditChatAdmin) (*mtproto.MutableChat, error)
	ChatEditChatDefaultBannedRights(ctx context.Context, in *chat.TLChatEditChatDefaultBannedRights) (*mtproto.MutableChat, error)
	ChatAddChatUser(ctx context.Context, in *chat.TLChatAddChatUser) (*mtproto.MutableChat, error)
	ChatGetMutableChatByLink(ctx context.Context, in *chat.TLChatGetMutableChatByLink) (*mtproto.MutableChat, error)
	ChatToggleNoForwards(ctx context.Context, in *chat.TLChatToggleNoForwards) (*mtproto.MutableChat, error)
	ChatMigratedToChannel(ctx context.Context, in *chat.TLChatMigratedToChannel) (*mtproto.Bool, error)
	ChatGetChatParticipantIdList(ctx context.Context, in *chat.TLChatGetChatParticipantIdList) (*chat.Vector_Long, error)
	ChatGetUsersChatIdList(ctx context.Context, in *chat.TLChatGetUsersChatIdList) (*chat.Vector_UserChatIdList, error)
	ChatGetMyChatList(ctx context.Context, in *chat.TLChatGetMyChatList) (*chat.Vector_MutableChat, error)
	ChatExportChatInvite(ctx context.Context, in *chat.TLChatExportChatInvite) (*mtproto.ExportedChatInvite, error)
	ChatGetAdminsWithInvites(ctx context.Context, in *chat.TLChatGetAdminsWithInvites) (*chat.Vector_ChatAdminWithInvites, error)
	ChatGetExportedChatInvite(ctx context.Context, in *chat.TLChatGetExportedChatInvite) (*mtproto.ExportedChatInvite, error)
	ChatGetExportedChatInvites(ctx context.Context, in *chat.TLChatGetExportedChatInvites) (*chat.Vector_ExportedChatInvite, error)
	ChatCheckChatInvite(ctx context.Context, in *chat.TLChatCheckChatInvite) (*chat.ChatInviteExt, error)
	ChatImportChatInvite(ctx context.Context, in *chat.TLChatImportChatInvite) (*mtproto.MutableChat, error)
	ChatGetChatInviteImporters(ctx context.Context, in *chat.TLChatGetChatInviteImporters) (*chat.Vector_ChatInviteImporter, error)
	ChatDeleteExportedChatInvite(ctx context.Context, in *chat.TLChatDeleteExportedChatInvite) (*mtproto.Bool, error)
	ChatDeleteRevokedExportedChatInvites(ctx context.Context, in *chat.TLChatDeleteRevokedExportedChatInvites) (*mtproto.Bool, error)
	ChatEditExportedChatInvite(ctx context.Context, in *chat.TLChatEditExportedChatInvite) (*chat.Vector_ExportedChatInvite, error)
	ChatSetChatAvailableReactions(ctx context.Context, in *chat.TLChatSetChatAvailableReactions) (*mtproto.MutableChat, error)
	ChatSetHistoryTTL(ctx context.Context, in *chat.TLChatSetHistoryTTL) (*mtproto.MutableChat, error)
	ChatSearch(ctx context.Context, in *chat.TLChatSearch) (*chat.Vector_MutableChat, error)
	ChatGetRecentChatInviteRequesters(ctx context.Context, in *chat.TLChatGetRecentChatInviteRequesters) (*chat.RecentChatInviteRequesters, error)
	ChatHideChatJoinRequests(ctx context.Context, in *chat.TLChatHideChatJoinRequests) (*chat.RecentChatInviteRequesters, error)
	ChatImportChatInvite2(ctx context.Context, in *chat.TLChatImportChatInvite2) (*chat.ChatInviteImported, error)
}

type defaultChatClient struct {
	cli zrpc.Client
}

func NewChatClient(cli zrpc.Client) ChatClient {
	return &defaultChatClient{
		cli: cli,
	}
}

// ChatGetMutableChat
// chat.getMutableChat chat_id:long = MutableChat;
func (m *defaultChatClient) ChatGetMutableChat(ctx context.Context, in *chat.TLChatGetMutableChat) (*mtproto.MutableChat, error) {
	client := chat.NewRPCChatClient(m.cli.Conn())
	return client.ChatGetMutableChat(ctx, in)
}

// ChatGetChatListByIdList
// chat.getChatListByIdList self_id:long id_list:Vector<long> = Vector<MutableChat>;
func (m *defaultChatClient) ChatGetChatListByIdList(ctx context.Context, in *chat.TLChatGetChatListByIdList) (*chat.Vector_MutableChat, error) {
	client := chat.NewRPCChatClient(m.cli.Conn())
	return client.ChatGetChatListByIdList(ctx, in)
}

// ChatGetChatBySelfId
// chat.getChatBySelfId self_id:long chat_id:long = MutableChat;
func (m *defaultChatClient) ChatGetChatBySelfId(ctx context.Context, in *chat.TLChatGetChatBySelfId) (*mtproto.MutableChat, error) {
	client := chat.NewRPCChatClient(m.cli.Conn())
	return client.ChatGetChatBySelfId(ctx, in)
}

// ChatCreateChat2
// chat.createChat2 creator_id:long user_id_list:Vector<long> title:string = MutableChat;
func (m *defaultChatClient) ChatCreateChat2(ctx context.Context, in *chat.TLChatCreateChat2) (*mtproto.MutableChat, error) {
	client := chat.NewRPCChatClient(m.cli.Conn())
	return client.ChatCreateChat2(ctx, in)
}

// ChatDeleteChat
// chat.deleteChat chat_id:long operator_id:long = MutableChat;
func (m *defaultChatClient) ChatDeleteChat(ctx context.Context, in *chat.TLChatDeleteChat) (*mtproto.MutableChat, error) {
	client := chat.NewRPCChatClient(m.cli.Conn())
	return client.ChatDeleteChat(ctx, in)
}

// ChatDeleteChatUser
// chat.deleteChatUser chat_id:long operator_id:long delete_user_id:long = MutableChat;
func (m *defaultChatClient) ChatDeleteChatUser(ctx context.Context, in *chat.TLChatDeleteChatUser) (*mtproto.MutableChat, error) {
	client := chat.NewRPCChatClient(m.cli.Conn())
	return client.ChatDeleteChatUser(ctx, in)
}

// ChatEditChatTitle
// chat.editChatTitle chat_id:long edit_user_id:long title:string = MutableChat;
func (m *defaultChatClient) ChatEditChatTitle(ctx context.Context, in *chat.TLChatEditChatTitle) (*mtproto.MutableChat, error) {
	client := chat.NewRPCChatClient(m.cli.Conn())
	return client.ChatEditChatTitle(ctx, in)
}

// ChatEditChatAbout
// chat.editChatAbout chat_id:long edit_user_id:long about:string = MutableChat;
func (m *defaultChatClient) ChatEditChatAbout(ctx context.Context, in *chat.TLChatEditChatAbout) (*mtproto.MutableChat, error) {
	client := chat.NewRPCChatClient(m.cli.Conn())
	return client.ChatEditChatAbout(ctx, in)
}

// ChatEditChatPhoto
// chat.editChatPhoto chat_id:long edit_user_id:long chat_photo:Photo = MutableChat;
func (m *defaultChatClient) ChatEditChatPhoto(ctx context.Context, in *chat.TLChatEditChatPhoto) (*mtproto.MutableChat, error) {
	client := chat.NewRPCChatClient(m.cli.Conn())
	return client.ChatEditChatPhoto(ctx, in)
}

// ChatEditChatAdmin
// chat.editChatAdmin chat_id:long operator_id:long edit_chat_admin_id:long is_admin:Bool = MutableChat;
func (m *defaultChatClient) ChatEditChatAdmin(ctx context.Context, in *chat.TLChatEditChatAdmin) (*mtproto.MutableChat, error) {
	client := chat.NewRPCChatClient(m.cli.Conn())
	return client.ChatEditChatAdmin(ctx, in)
}

// ChatEditChatDefaultBannedRights
// chat.editChatDefaultBannedRights chat_id:long operator_id:long banned_rights:ChatBannedRights = MutableChat;
func (m *defaultChatClient) ChatEditChatDefaultBannedRights(ctx context.Context, in *chat.TLChatEditChatDefaultBannedRights) (*mtproto.MutableChat, error) {
	client := chat.NewRPCChatClient(m.cli.Conn())
	return client.ChatEditChatDefaultBannedRights(ctx, in)
}

// ChatAddChatUser
// chat.addChatUser chat_id:long inviter_id:long user_id:long = MutableChat;
func (m *defaultChatClient) ChatAddChatUser(ctx context.Context, in *chat.TLChatAddChatUser) (*mtproto.MutableChat, error) {
	client := chat.NewRPCChatClient(m.cli.Conn())
	return client.ChatAddChatUser(ctx, in)
}

// ChatGetMutableChatByLink
// chat.getMutableChatByLink link:string = MutableChat;
func (m *defaultChatClient) ChatGetMutableChatByLink(ctx context.Context, in *chat.TLChatGetMutableChatByLink) (*mtproto.MutableChat, error) {
	client := chat.NewRPCChatClient(m.cli.Conn())
	return client.ChatGetMutableChatByLink(ctx, in)
}

// ChatToggleNoForwards
// chat.toggleNoForwards chat_id:long operator_id:long enabled:Bool = MutableChat;
func (m *defaultChatClient) ChatToggleNoForwards(ctx context.Context, in *chat.TLChatToggleNoForwards) (*mtproto.MutableChat, error) {
	client := chat.NewRPCChatClient(m.cli.Conn())
	return client.ChatToggleNoForwards(ctx, in)
}

// ChatMigratedToChannel
// chat.migratedToChannel chat:MutableChat id:long access_hash:long = Bool;
func (m *defaultChatClient) ChatMigratedToChannel(ctx context.Context, in *chat.TLChatMigratedToChannel) (*mtproto.Bool, error) {
	client := chat.NewRPCChatClient(m.cli.Conn())
	return client.ChatMigratedToChannel(ctx, in)
}

// ChatGetChatParticipantIdList
// chat.getChatParticipantIdList chat_id:long = Vector<long>;
func (m *defaultChatClient) ChatGetChatParticipantIdList(ctx context.Context, in *chat.TLChatGetChatParticipantIdList) (*chat.Vector_Long, error) {
	client := chat.NewRPCChatClient(m.cli.Conn())
	return client.ChatGetChatParticipantIdList(ctx, in)
}

// ChatGetUsersChatIdList
// chat.getUsersChatIdList id:Vector<long> = Vector<UserChatIdList>;
func (m *defaultChatClient) ChatGetUsersChatIdList(ctx context.Context, in *chat.TLChatGetUsersChatIdList) (*chat.Vector_UserChatIdList, error) {
	client := chat.NewRPCChatClient(m.cli.Conn())
	return client.ChatGetUsersChatIdList(ctx, in)
}

// ChatGetMyChatList
// chat.getMyChatList user_id:long is_creator:Bool = Vector<MutableChat>;
func (m *defaultChatClient) ChatGetMyChatList(ctx context.Context, in *chat.TLChatGetMyChatList) (*chat.Vector_MutableChat, error) {
	client := chat.NewRPCChatClient(m.cli.Conn())
	return client.ChatGetMyChatList(ctx, in)
}

// ChatExportChatInvite
// chat.exportChatInvite flags:# chat_id:long admin_id:long legacy_revoke_permanent:flags.2?true request_needed:flags.3?true expire_date:flags.0?int usage_limit:flags.1?int title:flags.4?string = ExportedChatInvite;
func (m *defaultChatClient) ChatExportChatInvite(ctx context.Context, in *chat.TLChatExportChatInvite) (*mtproto.ExportedChatInvite, error) {
	client := chat.NewRPCChatClient(m.cli.Conn())
	return client.ChatExportChatInvite(ctx, in)
}

// ChatGetAdminsWithInvites
// chat.getAdminsWithInvites self_id:long chat_id:long = Vector<ChatAdminWithInvites>;
func (m *defaultChatClient) ChatGetAdminsWithInvites(ctx context.Context, in *chat.TLChatGetAdminsWithInvites) (*chat.Vector_ChatAdminWithInvites, error) {
	client := chat.NewRPCChatClient(m.cli.Conn())
	return client.ChatGetAdminsWithInvites(ctx, in)
}

// ChatGetExportedChatInvite
// chat.getExportedChatInvite chat_id:long link:string = ExportedChatInvite;
func (m *defaultChatClient) ChatGetExportedChatInvite(ctx context.Context, in *chat.TLChatGetExportedChatInvite) (*mtproto.ExportedChatInvite, error) {
	client := chat.NewRPCChatClient(m.cli.Conn())
	return client.ChatGetExportedChatInvite(ctx, in)
}

// ChatGetExportedChatInvites
// chat.getExportedChatInvites flags:# chat_id:long admin_id:long revoked:flags.3?true offset_date:flags.2?int offset_link:flags.2?string limit:int = Vector<ExportedChatInvite>;
func (m *defaultChatClient) ChatGetExportedChatInvites(ctx context.Context, in *chat.TLChatGetExportedChatInvites) (*chat.Vector_ExportedChatInvite, error) {
	client := chat.NewRPCChatClient(m.cli.Conn())
	return client.ChatGetExportedChatInvites(ctx, in)
}

// ChatCheckChatInvite
// chat.checkChatInvite self_id:long hash:string = ChatInviteExt;
func (m *defaultChatClient) ChatCheckChatInvite(ctx context.Context, in *chat.TLChatCheckChatInvite) (*chat.ChatInviteExt, error) {
	client := chat.NewRPCChatClient(m.cli.Conn())
	return client.ChatCheckChatInvite(ctx, in)
}

// ChatImportChatInvite
// chat.importChatInvite self_id:long hash:string = MutableChat;
func (m *defaultChatClient) ChatImportChatInvite(ctx context.Context, in *chat.TLChatImportChatInvite) (*mtproto.MutableChat, error) {
	client := chat.NewRPCChatClient(m.cli.Conn())
	return client.ChatImportChatInvite(ctx, in)
}

// ChatGetChatInviteImporters
// chat.getChatInviteImporters flags:# self_id:long chat_id:long requested:flags.0?true link:flags.1?string q:flags.2?string offset_date:int offset_user:long limit:int = Vector<ChatInviteImporter>;
func (m *defaultChatClient) ChatGetChatInviteImporters(ctx context.Context, in *chat.TLChatGetChatInviteImporters) (*chat.Vector_ChatInviteImporter, error) {
	client := chat.NewRPCChatClient(m.cli.Conn())
	return client.ChatGetChatInviteImporters(ctx, in)
}

// ChatDeleteExportedChatInvite
// chat.deleteExportedChatInvite self_id:long chat_id:long link:string = Bool;
func (m *defaultChatClient) ChatDeleteExportedChatInvite(ctx context.Context, in *chat.TLChatDeleteExportedChatInvite) (*mtproto.Bool, error) {
	client := chat.NewRPCChatClient(m.cli.Conn())
	return client.ChatDeleteExportedChatInvite(ctx, in)
}

// ChatDeleteRevokedExportedChatInvites
// chat.deleteRevokedExportedChatInvites self_id:long chat_id:long admin_id:long = Bool;
func (m *defaultChatClient) ChatDeleteRevokedExportedChatInvites(ctx context.Context, in *chat.TLChatDeleteRevokedExportedChatInvites) (*mtproto.Bool, error) {
	client := chat.NewRPCChatClient(m.cli.Conn())
	return client.ChatDeleteRevokedExportedChatInvites(ctx, in)
}

// ChatEditExportedChatInvite
// chat.editExportedChatInvite flags:# self_id:long chat_id:long revoked:flags.2?true link:string expire_date:flags.0?int usage_limit:flags.1?int request_needed:flags.3?Bool title:flags.4?string = Vector<ExportedChatInvite>;
func (m *defaultChatClient) ChatEditExportedChatInvite(ctx context.Context, in *chat.TLChatEditExportedChatInvite) (*chat.Vector_ExportedChatInvite, error) {
	client := chat.NewRPCChatClient(m.cli.Conn())
	return client.ChatEditExportedChatInvite(ctx, in)
}

// ChatSetChatAvailableReactions
// chat.setChatAvailableReactions self_id:long chat_id:long available_reactions_type:int available_reactions:Vector<string> = MutableChat;
func (m *defaultChatClient) ChatSetChatAvailableReactions(ctx context.Context, in *chat.TLChatSetChatAvailableReactions) (*mtproto.MutableChat, error) {
	client := chat.NewRPCChatClient(m.cli.Conn())
	return client.ChatSetChatAvailableReactions(ctx, in)
}

// ChatSetHistoryTTL
// chat.setHistoryTTL self_id:long chat_id:long ttl_period:int = MutableChat;
func (m *defaultChatClient) ChatSetHistoryTTL(ctx context.Context, in *chat.TLChatSetHistoryTTL) (*mtproto.MutableChat, error) {
	client := chat.NewRPCChatClient(m.cli.Conn())
	return client.ChatSetHistoryTTL(ctx, in)
}

// ChatSearch
// chat.search self_id:long q:string offset:long limit:int = Vector<MutableChat>;
func (m *defaultChatClient) ChatSearch(ctx context.Context, in *chat.TLChatSearch) (*chat.Vector_MutableChat, error) {
	client := chat.NewRPCChatClient(m.cli.Conn())
	return client.ChatSearch(ctx, in)
}

// ChatGetRecentChatInviteRequesters
// chat.getRecentChatInviteRequesters self_id:long chat_id:long = RecentChatInviteRequesters;
func (m *defaultChatClient) ChatGetRecentChatInviteRequesters(ctx context.Context, in *chat.TLChatGetRecentChatInviteRequesters) (*chat.RecentChatInviteRequesters, error) {
	client := chat.NewRPCChatClient(m.cli.Conn())
	return client.ChatGetRecentChatInviteRequesters(ctx, in)
}

// ChatHideChatJoinRequests
// chat.hideChatJoinRequests flags:# self_id:long chat_id:long approved:flags.0?true link:flags.1?string user_id:flags.2?long = RecentChatInviteRequesters;
func (m *defaultChatClient) ChatHideChatJoinRequests(ctx context.Context, in *chat.TLChatHideChatJoinRequests) (*chat.RecentChatInviteRequesters, error) {
	client := chat.NewRPCChatClient(m.cli.Conn())
	return client.ChatHideChatJoinRequests(ctx, in)
}

// ChatImportChatInvite2
// chat.importChatInvite2 self_id:long hash:string = ChatInviteImported;
func (m *defaultChatClient) ChatImportChatInvite2(ctx context.Context, in *chat.TLChatImportChatInvite2) (*chat.ChatInviteImported, error) {
	client := chat.NewRPCChatClient(m.cli.Conn())
	return client.ChatImportChatInvite2(ctx, in)
}
