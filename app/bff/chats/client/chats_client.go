/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright 2022 Teamgram Authors.
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package chats_client

import (
	"context"

	"github.com/teamgram/proto/mtproto"

	"github.com/zeromicro/go-zero/zrpc"
)

var _ *mtproto.Bool

type ChatsClient interface {
	MessagesGetChats(ctx context.Context, in *mtproto.TLMessagesGetChats) (*mtproto.Messages_Chats, error)
	MessagesGetFullChat(ctx context.Context, in *mtproto.TLMessagesGetFullChat) (*mtproto.Messages_ChatFull, error)
	MessagesEditChatTitle(ctx context.Context, in *mtproto.TLMessagesEditChatTitle) (*mtproto.Updates, error)
	MessagesEditChatPhoto(ctx context.Context, in *mtproto.TLMessagesEditChatPhoto) (*mtproto.Updates, error)
	MessagesAddChatUser(ctx context.Context, in *mtproto.TLMessagesAddChatUser) (*mtproto.Updates, error)
	MessagesDeleteChatUser(ctx context.Context, in *mtproto.TLMessagesDeleteChatUser) (*mtproto.Updates, error)
	MessagesCreateChat(ctx context.Context, in *mtproto.TLMessagesCreateChat) (*mtproto.Updates, error)
	MessagesExportChatInvite(ctx context.Context, in *mtproto.TLMessagesExportChatInvite) (*mtproto.ExportedChatInvite, error)
	MessagesCheckChatInvite(ctx context.Context, in *mtproto.TLMessagesCheckChatInvite) (*mtproto.ChatInvite, error)
	MessagesImportChatInvite(ctx context.Context, in *mtproto.TLMessagesImportChatInvite) (*mtproto.Updates, error)
	MessagesEditChatAdmin(ctx context.Context, in *mtproto.TLMessagesEditChatAdmin) (*mtproto.Bool, error)
	MessagesMigrateChat(ctx context.Context, in *mtproto.TLMessagesMigrateChat) (*mtproto.Updates, error)
	MessagesGetCommonChats(ctx context.Context, in *mtproto.TLMessagesGetCommonChats) (*mtproto.Messages_Chats, error)
	MessagesGetAllChats(ctx context.Context, in *mtproto.TLMessagesGetAllChats) (*mtproto.Messages_Chats, error)
	MessagesEditChatAbout(ctx context.Context, in *mtproto.TLMessagesEditChatAbout) (*mtproto.Bool, error)
	MessagesEditChatDefaultBannedRights(ctx context.Context, in *mtproto.TLMessagesEditChatDefaultBannedRights) (*mtproto.Updates, error)
	MessagesDeleteChat(ctx context.Context, in *mtproto.TLMessagesDeleteChat) (*mtproto.Bool, error)
	MessagesGetExportedChatInvites(ctx context.Context, in *mtproto.TLMessagesGetExportedChatInvites) (*mtproto.Messages_ExportedChatInvites, error)
	MessagesGetExportedChatInvite(ctx context.Context, in *mtproto.TLMessagesGetExportedChatInvite) (*mtproto.Messages_ExportedChatInvite, error)
	MessagesEditExportedChatInvite(ctx context.Context, in *mtproto.TLMessagesEditExportedChatInvite) (*mtproto.Messages_ExportedChatInvite, error)
	MessagesDeleteRevokedExportedChatInvites(ctx context.Context, in *mtproto.TLMessagesDeleteRevokedExportedChatInvites) (*mtproto.Bool, error)
	MessagesDeleteExportedChatInvite(ctx context.Context, in *mtproto.TLMessagesDeleteExportedChatInvite) (*mtproto.Bool, error)
	MessagesGetAdminsWithInvites(ctx context.Context, in *mtproto.TLMessagesGetAdminsWithInvites) (*mtproto.Messages_ChatAdminsWithInvites, error)
	MessagesGetChatInviteImporters(ctx context.Context, in *mtproto.TLMessagesGetChatInviteImporters) (*mtproto.Messages_ChatInviteImporters, error)
	MessagesGetMessageReadParticipants(ctx context.Context, in *mtproto.TLMessagesGetMessageReadParticipants) (*mtproto.Vector_Long, error)
	MessagesHideChatJoinRequest(ctx context.Context, in *mtproto.TLMessagesHideChatJoinRequest) (*mtproto.Updates, error)
	MessagesHideAllChatJoinRequests(ctx context.Context, in *mtproto.TLMessagesHideAllChatJoinRequests) (*mtproto.Updates, error)
	ChannelsConvertToGigagroup(ctx context.Context, in *mtproto.TLChannelsConvertToGigagroup) (*mtproto.Updates, error)
}

type defaultChatsClient struct {
	cli zrpc.Client
}

func NewChatsClient(cli zrpc.Client) ChatsClient {
	return &defaultChatsClient{
		cli: cli,
	}
}

// MessagesGetChats
// messages.getChats#49e9528f id:Vector<long> = messages.Chats;
func (m *defaultChatsClient) MessagesGetChats(ctx context.Context, in *mtproto.TLMessagesGetChats) (*mtproto.Messages_Chats, error) {
	client := mtproto.NewRPCChatsClient(m.cli.Conn())
	return client.MessagesGetChats(ctx, in)
}

// MessagesGetFullChat
// messages.getFullChat#aeb00b34 chat_id:long = messages.ChatFull;
func (m *defaultChatsClient) MessagesGetFullChat(ctx context.Context, in *mtproto.TLMessagesGetFullChat) (*mtproto.Messages_ChatFull, error) {
	client := mtproto.NewRPCChatsClient(m.cli.Conn())
	return client.MessagesGetFullChat(ctx, in)
}

// MessagesEditChatTitle
// messages.editChatTitle#73783ffd chat_id:long title:string = Updates;
func (m *defaultChatsClient) MessagesEditChatTitle(ctx context.Context, in *mtproto.TLMessagesEditChatTitle) (*mtproto.Updates, error) {
	client := mtproto.NewRPCChatsClient(m.cli.Conn())
	return client.MessagesEditChatTitle(ctx, in)
}

// MessagesEditChatPhoto
// messages.editChatPhoto#35ddd674 chat_id:long photo:InputChatPhoto = Updates;
func (m *defaultChatsClient) MessagesEditChatPhoto(ctx context.Context, in *mtproto.TLMessagesEditChatPhoto) (*mtproto.Updates, error) {
	client := mtproto.NewRPCChatsClient(m.cli.Conn())
	return client.MessagesEditChatPhoto(ctx, in)
}

// MessagesAddChatUser
// messages.addChatUser#f24753e3 chat_id:long user_id:InputUser fwd_limit:int = Updates;
func (m *defaultChatsClient) MessagesAddChatUser(ctx context.Context, in *mtproto.TLMessagesAddChatUser) (*mtproto.Updates, error) {
	client := mtproto.NewRPCChatsClient(m.cli.Conn())
	return client.MessagesAddChatUser(ctx, in)
}

// MessagesDeleteChatUser
// messages.deleteChatUser#a2185cab flags:# revoke_history:flags.0?true chat_id:long user_id:InputUser = Updates;
func (m *defaultChatsClient) MessagesDeleteChatUser(ctx context.Context, in *mtproto.TLMessagesDeleteChatUser) (*mtproto.Updates, error) {
	client := mtproto.NewRPCChatsClient(m.cli.Conn())
	return client.MessagesDeleteChatUser(ctx, in)
}

// MessagesCreateChat
// messages.createChat#9cb126e users:Vector<InputUser> title:string = Updates;
func (m *defaultChatsClient) MessagesCreateChat(ctx context.Context, in *mtproto.TLMessagesCreateChat) (*mtproto.Updates, error) {
	client := mtproto.NewRPCChatsClient(m.cli.Conn())
	return client.MessagesCreateChat(ctx, in)
}

// MessagesExportChatInvite
// messages.exportChatInvite#a02ce5d5 flags:# legacy_revoke_permanent:flags.2?true request_needed:flags.3?true peer:InputPeer expire_date:flags.0?int usage_limit:flags.1?int title:flags.4?string = ExportedChatInvite;
func (m *defaultChatsClient) MessagesExportChatInvite(ctx context.Context, in *mtproto.TLMessagesExportChatInvite) (*mtproto.ExportedChatInvite, error) {
	client := mtproto.NewRPCChatsClient(m.cli.Conn())
	return client.MessagesExportChatInvite(ctx, in)
}

// MessagesCheckChatInvite
// messages.checkChatInvite#3eadb1bb hash:string = ChatInvite;
func (m *defaultChatsClient) MessagesCheckChatInvite(ctx context.Context, in *mtproto.TLMessagesCheckChatInvite) (*mtproto.ChatInvite, error) {
	client := mtproto.NewRPCChatsClient(m.cli.Conn())
	return client.MessagesCheckChatInvite(ctx, in)
}

// MessagesImportChatInvite
// messages.importChatInvite#6c50051c hash:string = Updates;
func (m *defaultChatsClient) MessagesImportChatInvite(ctx context.Context, in *mtproto.TLMessagesImportChatInvite) (*mtproto.Updates, error) {
	client := mtproto.NewRPCChatsClient(m.cli.Conn())
	return client.MessagesImportChatInvite(ctx, in)
}

// MessagesEditChatAdmin
// messages.editChatAdmin#a85bd1c2 chat_id:long user_id:InputUser is_admin:Bool = Bool;
func (m *defaultChatsClient) MessagesEditChatAdmin(ctx context.Context, in *mtproto.TLMessagesEditChatAdmin) (*mtproto.Bool, error) {
	client := mtproto.NewRPCChatsClient(m.cli.Conn())
	return client.MessagesEditChatAdmin(ctx, in)
}

// MessagesMigrateChat
// messages.migrateChat#a2875319 chat_id:long = Updates;
func (m *defaultChatsClient) MessagesMigrateChat(ctx context.Context, in *mtproto.TLMessagesMigrateChat) (*mtproto.Updates, error) {
	client := mtproto.NewRPCChatsClient(m.cli.Conn())
	return client.MessagesMigrateChat(ctx, in)
}

// MessagesGetCommonChats
// messages.getCommonChats#e40ca104 user_id:InputUser max_id:long limit:int = messages.Chats;
func (m *defaultChatsClient) MessagesGetCommonChats(ctx context.Context, in *mtproto.TLMessagesGetCommonChats) (*mtproto.Messages_Chats, error) {
	client := mtproto.NewRPCChatsClient(m.cli.Conn())
	return client.MessagesGetCommonChats(ctx, in)
}

// MessagesGetAllChats
// messages.getAllChats#875f74be except_ids:Vector<long> = messages.Chats;
func (m *defaultChatsClient) MessagesGetAllChats(ctx context.Context, in *mtproto.TLMessagesGetAllChats) (*mtproto.Messages_Chats, error) {
	client := mtproto.NewRPCChatsClient(m.cli.Conn())
	return client.MessagesGetAllChats(ctx, in)
}

// MessagesEditChatAbout
// messages.editChatAbout#def60797 peer:InputPeer about:string = Bool;
func (m *defaultChatsClient) MessagesEditChatAbout(ctx context.Context, in *mtproto.TLMessagesEditChatAbout) (*mtproto.Bool, error) {
	client := mtproto.NewRPCChatsClient(m.cli.Conn())
	return client.MessagesEditChatAbout(ctx, in)
}

// MessagesEditChatDefaultBannedRights
// messages.editChatDefaultBannedRights#a5866b41 peer:InputPeer banned_rights:ChatBannedRights = Updates;
func (m *defaultChatsClient) MessagesEditChatDefaultBannedRights(ctx context.Context, in *mtproto.TLMessagesEditChatDefaultBannedRights) (*mtproto.Updates, error) {
	client := mtproto.NewRPCChatsClient(m.cli.Conn())
	return client.MessagesEditChatDefaultBannedRights(ctx, in)
}

// MessagesDeleteChat
// messages.deleteChat#5bd0ee50 chat_id:long = Bool;
func (m *defaultChatsClient) MessagesDeleteChat(ctx context.Context, in *mtproto.TLMessagesDeleteChat) (*mtproto.Bool, error) {
	client := mtproto.NewRPCChatsClient(m.cli.Conn())
	return client.MessagesDeleteChat(ctx, in)
}

// MessagesGetExportedChatInvites
// messages.getExportedChatInvites#a2b5a3f6 flags:# revoked:flags.3?true peer:InputPeer admin_id:InputUser offset_date:flags.2?int offset_link:flags.2?string limit:int = messages.ExportedChatInvites;
func (m *defaultChatsClient) MessagesGetExportedChatInvites(ctx context.Context, in *mtproto.TLMessagesGetExportedChatInvites) (*mtproto.Messages_ExportedChatInvites, error) {
	client := mtproto.NewRPCChatsClient(m.cli.Conn())
	return client.MessagesGetExportedChatInvites(ctx, in)
}

// MessagesGetExportedChatInvite
// messages.getExportedChatInvite#73746f5c peer:InputPeer link:string = messages.ExportedChatInvite;
func (m *defaultChatsClient) MessagesGetExportedChatInvite(ctx context.Context, in *mtproto.TLMessagesGetExportedChatInvite) (*mtproto.Messages_ExportedChatInvite, error) {
	client := mtproto.NewRPCChatsClient(m.cli.Conn())
	return client.MessagesGetExportedChatInvite(ctx, in)
}

// MessagesEditExportedChatInvite
// messages.editExportedChatInvite#bdca2f75 flags:# revoked:flags.2?true peer:InputPeer link:string expire_date:flags.0?int usage_limit:flags.1?int request_needed:flags.3?Bool title:flags.4?string = messages.ExportedChatInvite;
func (m *defaultChatsClient) MessagesEditExportedChatInvite(ctx context.Context, in *mtproto.TLMessagesEditExportedChatInvite) (*mtproto.Messages_ExportedChatInvite, error) {
	client := mtproto.NewRPCChatsClient(m.cli.Conn())
	return client.MessagesEditExportedChatInvite(ctx, in)
}

// MessagesDeleteRevokedExportedChatInvites
// messages.deleteRevokedExportedChatInvites#56987bd5 peer:InputPeer admin_id:InputUser = Bool;
func (m *defaultChatsClient) MessagesDeleteRevokedExportedChatInvites(ctx context.Context, in *mtproto.TLMessagesDeleteRevokedExportedChatInvites) (*mtproto.Bool, error) {
	client := mtproto.NewRPCChatsClient(m.cli.Conn())
	return client.MessagesDeleteRevokedExportedChatInvites(ctx, in)
}

// MessagesDeleteExportedChatInvite
// messages.deleteExportedChatInvite#d464a42b peer:InputPeer link:string = Bool;
func (m *defaultChatsClient) MessagesDeleteExportedChatInvite(ctx context.Context, in *mtproto.TLMessagesDeleteExportedChatInvite) (*mtproto.Bool, error) {
	client := mtproto.NewRPCChatsClient(m.cli.Conn())
	return client.MessagesDeleteExportedChatInvite(ctx, in)
}

// MessagesGetAdminsWithInvites
// messages.getAdminsWithInvites#3920e6ef peer:InputPeer = messages.ChatAdminsWithInvites;
func (m *defaultChatsClient) MessagesGetAdminsWithInvites(ctx context.Context, in *mtproto.TLMessagesGetAdminsWithInvites) (*mtproto.Messages_ChatAdminsWithInvites, error) {
	client := mtproto.NewRPCChatsClient(m.cli.Conn())
	return client.MessagesGetAdminsWithInvites(ctx, in)
}

// MessagesGetChatInviteImporters
// messages.getChatInviteImporters#df04dd4e flags:# requested:flags.0?true peer:InputPeer link:flags.1?string q:flags.2?string offset_date:int offset_user:InputUser limit:int = messages.ChatInviteImporters;
func (m *defaultChatsClient) MessagesGetChatInviteImporters(ctx context.Context, in *mtproto.TLMessagesGetChatInviteImporters) (*mtproto.Messages_ChatInviteImporters, error) {
	client := mtproto.NewRPCChatsClient(m.cli.Conn())
	return client.MessagesGetChatInviteImporters(ctx, in)
}

// MessagesGetMessageReadParticipants
// messages.getMessageReadParticipants#2c6f97b7 peer:InputPeer msg_id:int = Vector<long>;
func (m *defaultChatsClient) MessagesGetMessageReadParticipants(ctx context.Context, in *mtproto.TLMessagesGetMessageReadParticipants) (*mtproto.Vector_Long, error) {
	client := mtproto.NewRPCChatsClient(m.cli.Conn())
	return client.MessagesGetMessageReadParticipants(ctx, in)
}

// MessagesHideChatJoinRequest
// messages.hideChatJoinRequest#7fe7e815 flags:# approved:flags.0?true peer:InputPeer user_id:InputUser = Updates;
func (m *defaultChatsClient) MessagesHideChatJoinRequest(ctx context.Context, in *mtproto.TLMessagesHideChatJoinRequest) (*mtproto.Updates, error) {
	client := mtproto.NewRPCChatsClient(m.cli.Conn())
	return client.MessagesHideChatJoinRequest(ctx, in)
}

// MessagesHideAllChatJoinRequests
// messages.hideAllChatJoinRequests#e085f4ea flags:# approved:flags.0?true peer:InputPeer link:flags.1?string = Updates;
func (m *defaultChatsClient) MessagesHideAllChatJoinRequests(ctx context.Context, in *mtproto.TLMessagesHideAllChatJoinRequests) (*mtproto.Updates, error) {
	client := mtproto.NewRPCChatsClient(m.cli.Conn())
	return client.MessagesHideAllChatJoinRequests(ctx, in)
}

// ChannelsConvertToGigagroup
// channels.convertToGigagroup#b290c69 channel:InputChannel = Updates;
func (m *defaultChatsClient) ChannelsConvertToGigagroup(ctx context.Context, in *mtproto.TLChannelsConvertToGigagroup) (*mtproto.Updates, error) {
	client := mtproto.NewRPCChatsClient(m.cli.Conn())
	return client.ChannelsConvertToGigagroup(ctx, in)
}
