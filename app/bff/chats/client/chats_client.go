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
	MessagesEditChatAdmin(ctx context.Context, in *mtproto.TLMessagesEditChatAdmin) (*mtproto.Bool, error)
	MessagesMigrateChat(ctx context.Context, in *mtproto.TLMessagesMigrateChat) (*mtproto.Updates, error)
	MessagesGetCommonChats(ctx context.Context, in *mtproto.TLMessagesGetCommonChats) (*mtproto.Messages_Chats, error)
	MessagesGetAllChats(ctx context.Context, in *mtproto.TLMessagesGetAllChats) (*mtproto.Messages_Chats, error)
	MessagesEditChatAbout(ctx context.Context, in *mtproto.TLMessagesEditChatAbout) (*mtproto.Bool, error)
	MessagesEditChatDefaultBannedRights(ctx context.Context, in *mtproto.TLMessagesEditChatDefaultBannedRights) (*mtproto.Updates, error)
	MessagesDeleteChat(ctx context.Context, in *mtproto.TLMessagesDeleteChat) (*mtproto.Bool, error)
	MessagesGetMessageReadParticipants(ctx context.Context, in *mtproto.TLMessagesGetMessageReadParticipants) (*mtproto.Vector_Long, error)
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

// MessagesGetMessageReadParticipants
// messages.getMessageReadParticipants#2c6f97b7 peer:InputPeer msg_id:int = Vector<long>;
func (m *defaultChatsClient) MessagesGetMessageReadParticipants(ctx context.Context, in *mtproto.TLMessagesGetMessageReadParticipants) (*mtproto.Vector_Long, error) {
	client := mtproto.NewRPCChatsClient(m.cli.Conn())
	return client.MessagesGetMessageReadParticipants(ctx, in)
}

// ChannelsConvertToGigagroup
// channels.convertToGigagroup#b290c69 channel:InputChannel = Updates;
func (m *defaultChatsClient) ChannelsConvertToGigagroup(ctx context.Context, in *mtproto.TLChannelsConvertToGigagroup) (*mtproto.Updates, error) {
	client := mtproto.NewRPCChatsClient(m.cli.Conn())
	return client.ChannelsConvertToGigagroup(ctx, in)
}
