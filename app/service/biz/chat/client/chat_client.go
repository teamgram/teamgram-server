/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2026-present, The Teamgram Authors (https://teamgram.net).
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package chatclient

import (
	"context"

	"github.com/teamgram/teamgram-server/v2/app/service/biz/chat/chat"
	"github.com/teamgram/teamgram-server/v2/app/service/biz/chat/chat/chatservice"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"

	"github.com/cloudwego/kitex/client"
)

var _ *tg.Bool

type ChatClient interface {
	ChatGetMutableChat(ctx context.Context, in *chat.TLChatGetMutableChat) (*tg.MutableChat, error)
	ChatGetChatListByIdList(ctx context.Context, in *chat.TLChatGetChatListByIdList) (*chat.VectorMutableChat, error)
	ChatGetChatProjectionBundle(ctx context.Context, in *chat.TLChatGetChatProjectionBundle) (*chat.ChatProjectionBundle, error)
	ChatGetChatBySelfId(ctx context.Context, in *chat.TLChatGetChatBySelfId) (*tg.MutableChat, error)
	ChatCheckChatAccess(ctx context.Context, in *chat.TLChatCheckChatAccess) (*chat.ChatAccessCheckResult, error)
	ChatCheckMessageAction(ctx context.Context, in *chat.TLChatCheckMessageAction) (*chat.MessageActionCheckResult, error)
	ChatCreateChat2(ctx context.Context, in *chat.TLChatCreateChat2) (*tg.MutableChat, error)
	ChatDeleteChat(ctx context.Context, in *chat.TLChatDeleteChat) (*tg.MutableChat, error)
	ChatDeleteChatUser(ctx context.Context, in *chat.TLChatDeleteChatUser) (*tg.MutableChat, error)
	ChatEditChatTitle(ctx context.Context, in *chat.TLChatEditChatTitle) (*tg.MutableChat, error)
	ChatEditChatAbout(ctx context.Context, in *chat.TLChatEditChatAbout) (*tg.MutableChat, error)
	ChatEditChatPhoto(ctx context.Context, in *chat.TLChatEditChatPhoto) (*tg.MutableChat, error)
	ChatEditChatAdmin(ctx context.Context, in *chat.TLChatEditChatAdmin) (*tg.MutableChat, error)
	ChatEditChatDefaultBannedRights(ctx context.Context, in *chat.TLChatEditChatDefaultBannedRights) (*tg.MutableChat, error)
	ChatAddChatUser(ctx context.Context, in *chat.TLChatAddChatUser) (*tg.MutableChat, error)
	ChatGetMutableChatByLink(ctx context.Context, in *chat.TLChatGetMutableChatByLink) (*tg.MutableChat, error)
	ChatToggleNoForwards(ctx context.Context, in *chat.TLChatToggleNoForwards) (*tg.MutableChat, error)
	ChatMigratedToChannel(ctx context.Context, in *chat.TLChatMigratedToChannel) (*tg.Bool, error)
	ChatGetChatParticipantIdList(ctx context.Context, in *chat.TLChatGetChatParticipantIdList) (*chat.VectorLong, error)
	ChatGetUsersChatIdList(ctx context.Context, in *chat.TLChatGetUsersChatIdList) (*chat.VectorUserChatIdList, error)
	ChatGetMyChatList(ctx context.Context, in *chat.TLChatGetMyChatList) (*chat.VectorMutableChat, error)
	ChatExportChatInvite(ctx context.Context, in *chat.TLChatExportChatInvite) (*tg.ExportedChatInvite, error)
	ChatGetAdminsWithInvites(ctx context.Context, in *chat.TLChatGetAdminsWithInvites) (*chat.VectorChatAdminWithInvites, error)
	ChatGetExportedChatInvite(ctx context.Context, in *chat.TLChatGetExportedChatInvite) (*tg.ExportedChatInvite, error)
	ChatGetExportedChatInvites(ctx context.Context, in *chat.TLChatGetExportedChatInvites) (*chat.VectorExportedChatInvite, error)
	ChatCheckChatInvite(ctx context.Context, in *chat.TLChatCheckChatInvite) (*chat.ChatInviteExt, error)
	ChatImportChatInvite(ctx context.Context, in *chat.TLChatImportChatInvite) (*tg.MutableChat, error)
	ChatGetChatInviteImporters(ctx context.Context, in *chat.TLChatGetChatInviteImporters) (*chat.VectorChatInviteImporter, error)
	ChatDeleteExportedChatInvite(ctx context.Context, in *chat.TLChatDeleteExportedChatInvite) (*tg.Bool, error)
	ChatDeleteRevokedExportedChatInvites(ctx context.Context, in *chat.TLChatDeleteRevokedExportedChatInvites) (*tg.Bool, error)
	ChatEditExportedChatInvite(ctx context.Context, in *chat.TLChatEditExportedChatInvite) (*chat.VectorExportedChatInvite, error)
	ChatSetChatAvailableReactions(ctx context.Context, in *chat.TLChatSetChatAvailableReactions) (*tg.MutableChat, error)
	ChatSetHistoryTTL(ctx context.Context, in *chat.TLChatSetHistoryTTL) (*tg.MutableChat, error)
	ChatSearch(ctx context.Context, in *chat.TLChatSearch) (*chat.VectorMutableChat, error)
	ChatGetRecentChatInviteRequesters(ctx context.Context, in *chat.TLChatGetRecentChatInviteRequesters) (*chat.RecentChatInviteRequesters, error)
	ChatHideChatJoinRequests(ctx context.Context, in *chat.TLChatHideChatJoinRequests) (*chat.RecentChatInviteRequesters, error)
	ChatImportChatInvite2(ctx context.Context, in *chat.TLChatImportChatInvite2) (*chat.ChatInviteImported, error)
}

type defaultChatClient struct {
	cli client.Client
	rpc chatservice.Client
}

func NewChatClient(cli client.Client) ChatClient {
	return &defaultChatClient{
		cli: cli,
		rpc: chatservice.NewRPCChatClient(cli),
	}
}

// ChatGetMutableChat
// chat.getMutableChat chat_id:long = MutableChat;
func (m *defaultChatClient) ChatGetMutableChat(ctx context.Context, in *chat.TLChatGetMutableChat) (*tg.MutableChat, error) {
	return m.rpc.ChatGetMutableChat(ctx, in)
}

// ChatGetChatListByIdList
// chat.getChatListByIdList self_id:long id_list:Vector<long> = Vector<MutableChat>;
func (m *defaultChatClient) ChatGetChatListByIdList(ctx context.Context, in *chat.TLChatGetChatListByIdList) (*chat.VectorMutableChat, error) {
	return m.rpc.ChatGetChatListByIdList(ctx, in)
}

// ChatGetChatProjectionBundle
// chat.getChatProjectionBundle viewer_user_ids:Vector<long> target_chat_ids:Vector<long> = ChatProjectionBundle;
func (m *defaultChatClient) ChatGetChatProjectionBundle(ctx context.Context, in *chat.TLChatGetChatProjectionBundle) (*chat.ChatProjectionBundle, error) {
	return m.rpc.ChatGetChatProjectionBundle(ctx, in)
}

// ChatGetChatBySelfId
// chat.getChatBySelfId self_id:long chat_id:long = MutableChat;
func (m *defaultChatClient) ChatGetChatBySelfId(ctx context.Context, in *chat.TLChatGetChatBySelfId) (*tg.MutableChat, error) {
	return m.rpc.ChatGetChatBySelfId(ctx, in)
}

// ChatCheckChatAccess
// chat.checkChatAccess self_id:long chat_id:long access_kind:string = ChatAccessCheckResult;
func (m *defaultChatClient) ChatCheckChatAccess(ctx context.Context, in *chat.TLChatCheckChatAccess) (*chat.ChatAccessCheckResult, error) {
	return m.rpc.ChatCheckChatAccess(ctx, in)
}

// ChatCheckMessageAction
// chat.checkMessageAction self_id:long chat_id:long action:string media_kind:string = MessageActionCheckResult;
func (m *defaultChatClient) ChatCheckMessageAction(ctx context.Context, in *chat.TLChatCheckMessageAction) (*chat.MessageActionCheckResult, error) {
	return m.rpc.ChatCheckMessageAction(ctx, in)
}

// ChatCreateChat2
// chat.createChat2 flags:# creator_id:long user_id_list:Vector<long> title:string bots:flags.0?Vector<long> ttl_period:flags.1?int client_msg_id:flags.2?long operation_id:flags.3?string = MutableChat;
func (m *defaultChatClient) ChatCreateChat2(ctx context.Context, in *chat.TLChatCreateChat2) (*tg.MutableChat, error) {
	return m.rpc.ChatCreateChat2(ctx, in)
}

// ChatDeleteChat
// chat.deleteChat chat_id:long operator_id:long = MutableChat;
func (m *defaultChatClient) ChatDeleteChat(ctx context.Context, in *chat.TLChatDeleteChat) (*tg.MutableChat, error) {
	return m.rpc.ChatDeleteChat(ctx, in)
}

// ChatDeleteChatUser
// chat.deleteChatUser chat_id:long operator_id:long delete_user_id:long = MutableChat;
func (m *defaultChatClient) ChatDeleteChatUser(ctx context.Context, in *chat.TLChatDeleteChatUser) (*tg.MutableChat, error) {
	return m.rpc.ChatDeleteChatUser(ctx, in)
}

// ChatEditChatTitle
// chat.editChatTitle chat_id:long edit_user_id:long title:string = MutableChat;
func (m *defaultChatClient) ChatEditChatTitle(ctx context.Context, in *chat.TLChatEditChatTitle) (*tg.MutableChat, error) {
	return m.rpc.ChatEditChatTitle(ctx, in)
}

// ChatEditChatAbout
// chat.editChatAbout chat_id:long edit_user_id:long about:string = MutableChat;
func (m *defaultChatClient) ChatEditChatAbout(ctx context.Context, in *chat.TLChatEditChatAbout) (*tg.MutableChat, error) {
	return m.rpc.ChatEditChatAbout(ctx, in)
}

// ChatEditChatPhoto
// chat.editChatPhoto chat_id:long edit_user_id:long chat_photo:Photo = MutableChat;
func (m *defaultChatClient) ChatEditChatPhoto(ctx context.Context, in *chat.TLChatEditChatPhoto) (*tg.MutableChat, error) {
	return m.rpc.ChatEditChatPhoto(ctx, in)
}

// ChatEditChatAdmin
// chat.editChatAdmin chat_id:long operator_id:long edit_chat_admin_id:long is_admin:Bool = MutableChat;
func (m *defaultChatClient) ChatEditChatAdmin(ctx context.Context, in *chat.TLChatEditChatAdmin) (*tg.MutableChat, error) {
	return m.rpc.ChatEditChatAdmin(ctx, in)
}

// ChatEditChatDefaultBannedRights
// chat.editChatDefaultBannedRights chat_id:long operator_id:long banned_rights:ChatBannedRights = MutableChat;
func (m *defaultChatClient) ChatEditChatDefaultBannedRights(ctx context.Context, in *chat.TLChatEditChatDefaultBannedRights) (*tg.MutableChat, error) {
	return m.rpc.ChatEditChatDefaultBannedRights(ctx, in)
}

// ChatAddChatUser
// chat.addChatUser flags:# chat_id:long inviter_id:long user_id:long is_bot:flags.0?true = MutableChat;
func (m *defaultChatClient) ChatAddChatUser(ctx context.Context, in *chat.TLChatAddChatUser) (*tg.MutableChat, error) {
	return m.rpc.ChatAddChatUser(ctx, in)
}

// ChatGetMutableChatByLink
// chat.getMutableChatByLink link:string = MutableChat;
func (m *defaultChatClient) ChatGetMutableChatByLink(ctx context.Context, in *chat.TLChatGetMutableChatByLink) (*tg.MutableChat, error) {
	return m.rpc.ChatGetMutableChatByLink(ctx, in)
}

// ChatToggleNoForwards
// chat.toggleNoForwards chat_id:long operator_id:long enabled:Bool = MutableChat;
func (m *defaultChatClient) ChatToggleNoForwards(ctx context.Context, in *chat.TLChatToggleNoForwards) (*tg.MutableChat, error) {
	return m.rpc.ChatToggleNoForwards(ctx, in)
}

// ChatMigratedToChannel
// chat.migratedToChannel chat:MutableChat id:long access_hash:long = Bool;
func (m *defaultChatClient) ChatMigratedToChannel(ctx context.Context, in *chat.TLChatMigratedToChannel) (*tg.Bool, error) {
	return m.rpc.ChatMigratedToChannel(ctx, in)
}

// ChatGetChatParticipantIdList
// chat.getChatParticipantIdList chat_id:long = Vector<long>;
func (m *defaultChatClient) ChatGetChatParticipantIdList(ctx context.Context, in *chat.TLChatGetChatParticipantIdList) (*chat.VectorLong, error) {
	return m.rpc.ChatGetChatParticipantIdList(ctx, in)
}

// ChatGetUsersChatIdList
// chat.getUsersChatIdList id:Vector<long> = Vector<UserChatIdList>;
func (m *defaultChatClient) ChatGetUsersChatIdList(ctx context.Context, in *chat.TLChatGetUsersChatIdList) (*chat.VectorUserChatIdList, error) {
	return m.rpc.ChatGetUsersChatIdList(ctx, in)
}

// ChatGetMyChatList
// chat.getMyChatList user_id:long is_creator:Bool = Vector<MutableChat>;
func (m *defaultChatClient) ChatGetMyChatList(ctx context.Context, in *chat.TLChatGetMyChatList) (*chat.VectorMutableChat, error) {
	return m.rpc.ChatGetMyChatList(ctx, in)
}

// ChatExportChatInvite
// chat.exportChatInvite flags:# chat_id:long admin_id:long legacy_revoke_permanent:flags.2?true request_needed:flags.3?true expire_date:flags.0?int usage_limit:flags.1?int title:flags.4?string = ExportedChatInvite;
func (m *defaultChatClient) ChatExportChatInvite(ctx context.Context, in *chat.TLChatExportChatInvite) (*tg.ExportedChatInvite, error) {
	return m.rpc.ChatExportChatInvite(ctx, in)
}

// ChatGetAdminsWithInvites
// chat.getAdminsWithInvites self_id:long chat_id:long = Vector<ChatAdminWithInvites>;
func (m *defaultChatClient) ChatGetAdminsWithInvites(ctx context.Context, in *chat.TLChatGetAdminsWithInvites) (*chat.VectorChatAdminWithInvites, error) {
	return m.rpc.ChatGetAdminsWithInvites(ctx, in)
}

// ChatGetExportedChatInvite
// chat.getExportedChatInvite chat_id:long link:string = ExportedChatInvite;
func (m *defaultChatClient) ChatGetExportedChatInvite(ctx context.Context, in *chat.TLChatGetExportedChatInvite) (*tg.ExportedChatInvite, error) {
	return m.rpc.ChatGetExportedChatInvite(ctx, in)
}

// ChatGetExportedChatInvites
// chat.getExportedChatInvites flags:# chat_id:long admin_id:long revoked:flags.3?true offset_date:flags.2?int offset_link:flags.2?string limit:int = Vector<ExportedChatInvite>;
func (m *defaultChatClient) ChatGetExportedChatInvites(ctx context.Context, in *chat.TLChatGetExportedChatInvites) (*chat.VectorExportedChatInvite, error) {
	return m.rpc.ChatGetExportedChatInvites(ctx, in)
}

// ChatCheckChatInvite
// chat.checkChatInvite self_id:long hash:string = ChatInviteExt;
func (m *defaultChatClient) ChatCheckChatInvite(ctx context.Context, in *chat.TLChatCheckChatInvite) (*chat.ChatInviteExt, error) {
	return m.rpc.ChatCheckChatInvite(ctx, in)
}

// ChatImportChatInvite
// chat.importChatInvite self_id:long hash:string = MutableChat;
func (m *defaultChatClient) ChatImportChatInvite(ctx context.Context, in *chat.TLChatImportChatInvite) (*tg.MutableChat, error) {
	return m.rpc.ChatImportChatInvite(ctx, in)
}

// ChatGetChatInviteImporters
// chat.getChatInviteImporters flags:# self_id:long chat_id:long requested:flags.0?true link:flags.1?string q:flags.2?string offset_date:int offset_user:long limit:int = Vector<ChatInviteImporter>;
func (m *defaultChatClient) ChatGetChatInviteImporters(ctx context.Context, in *chat.TLChatGetChatInviteImporters) (*chat.VectorChatInviteImporter, error) {
	return m.rpc.ChatGetChatInviteImporters(ctx, in)
}

// ChatDeleteExportedChatInvite
// chat.deleteExportedChatInvite self_id:long chat_id:long link:string = Bool;
func (m *defaultChatClient) ChatDeleteExportedChatInvite(ctx context.Context, in *chat.TLChatDeleteExportedChatInvite) (*tg.Bool, error) {
	return m.rpc.ChatDeleteExportedChatInvite(ctx, in)
}

// ChatDeleteRevokedExportedChatInvites
// chat.deleteRevokedExportedChatInvites self_id:long chat_id:long admin_id:long = Bool;
func (m *defaultChatClient) ChatDeleteRevokedExportedChatInvites(ctx context.Context, in *chat.TLChatDeleteRevokedExportedChatInvites) (*tg.Bool, error) {
	return m.rpc.ChatDeleteRevokedExportedChatInvites(ctx, in)
}

// ChatEditExportedChatInvite
// chat.editExportedChatInvite flags:# self_id:long chat_id:long revoked:flags.2?true link:string expire_date:flags.0?int usage_limit:flags.1?int request_needed:flags.3?Bool title:flags.4?string = Vector<ExportedChatInvite>;
func (m *defaultChatClient) ChatEditExportedChatInvite(ctx context.Context, in *chat.TLChatEditExportedChatInvite) (*chat.VectorExportedChatInvite, error) {
	return m.rpc.ChatEditExportedChatInvite(ctx, in)
}

// ChatSetChatAvailableReactions
// chat.setChatAvailableReactions self_id:long chat_id:long available_reactions_type:int available_reactions:Vector<string> = MutableChat;
func (m *defaultChatClient) ChatSetChatAvailableReactions(ctx context.Context, in *chat.TLChatSetChatAvailableReactions) (*tg.MutableChat, error) {
	return m.rpc.ChatSetChatAvailableReactions(ctx, in)
}

// ChatSetHistoryTTL
// chat.setHistoryTTL self_id:long chat_id:long ttl_period:int = MutableChat;
func (m *defaultChatClient) ChatSetHistoryTTL(ctx context.Context, in *chat.TLChatSetHistoryTTL) (*tg.MutableChat, error) {
	return m.rpc.ChatSetHistoryTTL(ctx, in)
}

// ChatSearch
// chat.search self_id:long q:string offset:long limit:int = Vector<MutableChat>;
func (m *defaultChatClient) ChatSearch(ctx context.Context, in *chat.TLChatSearch) (*chat.VectorMutableChat, error) {
	return m.rpc.ChatSearch(ctx, in)
}

// ChatGetRecentChatInviteRequesters
// chat.getRecentChatInviteRequesters self_id:long chat_id:long = RecentChatInviteRequesters;
func (m *defaultChatClient) ChatGetRecentChatInviteRequesters(ctx context.Context, in *chat.TLChatGetRecentChatInviteRequesters) (*chat.RecentChatInviteRequesters, error) {
	return m.rpc.ChatGetRecentChatInviteRequesters(ctx, in)
}

// ChatHideChatJoinRequests
// chat.hideChatJoinRequests flags:# self_id:long chat_id:long approved:flags.0?true link:flags.1?string user_id:flags.2?long = RecentChatInviteRequesters;
func (m *defaultChatClient) ChatHideChatJoinRequests(ctx context.Context, in *chat.TLChatHideChatJoinRequests) (*chat.RecentChatInviteRequesters, error) {
	return m.rpc.ChatHideChatJoinRequests(ctx, in)
}

// ChatImportChatInvite2
// chat.importChatInvite2 self_id:long hash:string = ChatInviteImported;
func (m *defaultChatClient) ChatImportChatInvite2(ctx context.Context, in *chat.TLChatImportChatInvite2) (*chat.ChatInviteImported, error) {
	return m.rpc.ChatImportChatInvite2(ctx, in)
}
