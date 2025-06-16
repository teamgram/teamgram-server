/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright 2024 Teamgram Authors.
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package chatsclient

import (
	"context"

	"github.com/teamgram/proto/v2/tg"
	"github.com/teamgram/teamgram-server/v2/app/bff/chats/chats/chatsservice"

	"github.com/cloudwego/kitex/client"
)

type ChatsClient interface {
	MessagesGetChats(ctx context.Context, in *tg.TLMessagesGetChats) (*tg.MessagesChats, error)
	MessagesGetFullChat(ctx context.Context, in *tg.TLMessagesGetFullChat) (*tg.MessagesChatFull, error)
	MessagesEditChatTitle(ctx context.Context, in *tg.TLMessagesEditChatTitle) (*tg.Updates, error)
	MessagesEditChatPhoto(ctx context.Context, in *tg.TLMessagesEditChatPhoto) (*tg.Updates, error)
	MessagesAddChatUser(ctx context.Context, in *tg.TLMessagesAddChatUser) (*tg.MessagesInvitedUsers, error)
	MessagesDeleteChatUser(ctx context.Context, in *tg.TLMessagesDeleteChatUser) (*tg.Updates, error)
	MessagesCreateChat(ctx context.Context, in *tg.TLMessagesCreateChat) (*tg.MessagesInvitedUsers, error)
	MessagesEditChatAdmin(ctx context.Context, in *tg.TLMessagesEditChatAdmin) (*tg.Bool, error)
	MessagesMigrateChat(ctx context.Context, in *tg.TLMessagesMigrateChat) (*tg.Updates, error)
	MessagesGetCommonChats(ctx context.Context, in *tg.TLMessagesGetCommonChats) (*tg.MessagesChats, error)
	MessagesEditChatAbout(ctx context.Context, in *tg.TLMessagesEditChatAbout) (*tg.Bool, error)
	MessagesEditChatDefaultBannedRights(ctx context.Context, in *tg.TLMessagesEditChatDefaultBannedRights) (*tg.Updates, error)
	MessagesDeleteChat(ctx context.Context, in *tg.TLMessagesDeleteChat) (*tg.Bool, error)
	MessagesGetMessageReadParticipants(ctx context.Context, in *tg.TLMessagesGetMessageReadParticipants) (*tg.VectorReadParticipantDate, error)
	ChannelsConvertToGigagroup(ctx context.Context, in *tg.TLChannelsConvertToGigagroup) (*tg.Updates, error)
	ChannelsSetEmojiStickers(ctx context.Context, in *tg.TLChannelsSetEmojiStickers) (*tg.Bool, error)
}

type defaultChatsClient struct {
	cli client.Client
}

func NewChatsClient(cli client.Client) ChatsClient {
	return &defaultChatsClient{
		cli: cli,
	}
}

// MessagesGetChats
// messages.getChats#49e9528f id:Vector<long> = messages.Chats;
func (m *defaultChatsClient) MessagesGetChats(ctx context.Context, in *tg.TLMessagesGetChats) (*tg.MessagesChats, error) {
	cli := chatsservice.NewRPCChatsClient(m.cli)
	return cli.MessagesGetChats(ctx, in)
}

// MessagesGetFullChat
// messages.getFullChat#aeb00b34 chat_id:long = messages.ChatFull;
func (m *defaultChatsClient) MessagesGetFullChat(ctx context.Context, in *tg.TLMessagesGetFullChat) (*tg.MessagesChatFull, error) {
	cli := chatsservice.NewRPCChatsClient(m.cli)
	return cli.MessagesGetFullChat(ctx, in)
}

// MessagesEditChatTitle
// messages.editChatTitle#73783ffd chat_id:long title:string = Updates;
func (m *defaultChatsClient) MessagesEditChatTitle(ctx context.Context, in *tg.TLMessagesEditChatTitle) (*tg.Updates, error) {
	cli := chatsservice.NewRPCChatsClient(m.cli)
	return cli.MessagesEditChatTitle(ctx, in)
}

// MessagesEditChatPhoto
// messages.editChatPhoto#35ddd674 chat_id:long photo:InputChatPhoto = Updates;
func (m *defaultChatsClient) MessagesEditChatPhoto(ctx context.Context, in *tg.TLMessagesEditChatPhoto) (*tg.Updates, error) {
	cli := chatsservice.NewRPCChatsClient(m.cli)
	return cli.MessagesEditChatPhoto(ctx, in)
}

// MessagesAddChatUser
// messages.addChatUser#cbc6d107 chat_id:long user_id:InputUser fwd_limit:int = messages.InvitedUsers;
func (m *defaultChatsClient) MessagesAddChatUser(ctx context.Context, in *tg.TLMessagesAddChatUser) (*tg.MessagesInvitedUsers, error) {
	cli := chatsservice.NewRPCChatsClient(m.cli)
	return cli.MessagesAddChatUser(ctx, in)
}

// MessagesDeleteChatUser
// messages.deleteChatUser#a2185cab flags:# revoke_history:flags.0?true chat_id:long user_id:InputUser = Updates;
func (m *defaultChatsClient) MessagesDeleteChatUser(ctx context.Context, in *tg.TLMessagesDeleteChatUser) (*tg.Updates, error) {
	cli := chatsservice.NewRPCChatsClient(m.cli)
	return cli.MessagesDeleteChatUser(ctx, in)
}

// MessagesCreateChat
// messages.createChat#92ceddd4 flags:# users:Vector<InputUser> title:string ttl_period:flags.0?int = messages.InvitedUsers;
func (m *defaultChatsClient) MessagesCreateChat(ctx context.Context, in *tg.TLMessagesCreateChat) (*tg.MessagesInvitedUsers, error) {
	cli := chatsservice.NewRPCChatsClient(m.cli)
	return cli.MessagesCreateChat(ctx, in)
}

// MessagesEditChatAdmin
// messages.editChatAdmin#a85bd1c2 chat_id:long user_id:InputUser is_admin:Bool = Bool;
func (m *defaultChatsClient) MessagesEditChatAdmin(ctx context.Context, in *tg.TLMessagesEditChatAdmin) (*tg.Bool, error) {
	cli := chatsservice.NewRPCChatsClient(m.cli)
	return cli.MessagesEditChatAdmin(ctx, in)
}

// MessagesMigrateChat
// messages.migrateChat#a2875319 chat_id:long = Updates;
func (m *defaultChatsClient) MessagesMigrateChat(ctx context.Context, in *tg.TLMessagesMigrateChat) (*tg.Updates, error) {
	cli := chatsservice.NewRPCChatsClient(m.cli)
	return cli.MessagesMigrateChat(ctx, in)
}

// MessagesGetCommonChats
// messages.getCommonChats#e40ca104 user_id:InputUser max_id:long limit:int = messages.Chats;
func (m *defaultChatsClient) MessagesGetCommonChats(ctx context.Context, in *tg.TLMessagesGetCommonChats) (*tg.MessagesChats, error) {
	cli := chatsservice.NewRPCChatsClient(m.cli)
	return cli.MessagesGetCommonChats(ctx, in)
}

// MessagesEditChatAbout
// messages.editChatAbout#def60797 peer:InputPeer about:string = Bool;
func (m *defaultChatsClient) MessagesEditChatAbout(ctx context.Context, in *tg.TLMessagesEditChatAbout) (*tg.Bool, error) {
	cli := chatsservice.NewRPCChatsClient(m.cli)
	return cli.MessagesEditChatAbout(ctx, in)
}

// MessagesEditChatDefaultBannedRights
// messages.editChatDefaultBannedRights#a5866b41 peer:InputPeer banned_rights:ChatBannedRights = Updates;
func (m *defaultChatsClient) MessagesEditChatDefaultBannedRights(ctx context.Context, in *tg.TLMessagesEditChatDefaultBannedRights) (*tg.Updates, error) {
	cli := chatsservice.NewRPCChatsClient(m.cli)
	return cli.MessagesEditChatDefaultBannedRights(ctx, in)
}

// MessagesDeleteChat
// messages.deleteChat#5bd0ee50 chat_id:long = Bool;
func (m *defaultChatsClient) MessagesDeleteChat(ctx context.Context, in *tg.TLMessagesDeleteChat) (*tg.Bool, error) {
	cli := chatsservice.NewRPCChatsClient(m.cli)
	return cli.MessagesDeleteChat(ctx, in)
}

// MessagesGetMessageReadParticipants
// messages.getMessageReadParticipants#31c1c44f peer:InputPeer msg_id:int = Vector<ReadParticipantDate>;
func (m *defaultChatsClient) MessagesGetMessageReadParticipants(ctx context.Context, in *tg.TLMessagesGetMessageReadParticipants) (*tg.VectorReadParticipantDate, error) {
	cli := chatsservice.NewRPCChatsClient(m.cli)
	return cli.MessagesGetMessageReadParticipants(ctx, in)
}

// ChannelsConvertToGigagroup
// channels.convertToGigagroup#b290c69 channel:InputChannel = Updates;
func (m *defaultChatsClient) ChannelsConvertToGigagroup(ctx context.Context, in *tg.TLChannelsConvertToGigagroup) (*tg.Updates, error) {
	cli := chatsservice.NewRPCChatsClient(m.cli)
	return cli.ChannelsConvertToGigagroup(ctx, in)
}

// ChannelsSetEmojiStickers
// channels.setEmojiStickers#3cd930b7 channel:InputChannel stickerset:InputStickerSet = Bool;
func (m *defaultChatsClient) ChannelsSetEmojiStickers(ctx context.Context, in *tg.TLChannelsSetEmojiStickers) (*tg.Bool, error) {
	cli := chatsservice.NewRPCChatsClient(m.cli)
	return cli.ChannelsSetEmojiStickers(ctx, in)
}
