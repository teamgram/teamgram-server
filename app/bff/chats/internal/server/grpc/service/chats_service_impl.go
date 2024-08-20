/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright 2024 Teamgram Authors.
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package service

import (
	"context"

	"github.com/teamgram/proto/mtproto"
	"github.com/teamgram/teamgram-server/app/bff/chats/internal/core"
)

// MessagesGetChats
// messages.getChats#49e9528f id:Vector<long> = messages.Chats;
func (s *Service) MessagesGetChats(ctx context.Context, request *mtproto.TLMessagesGetChats) (*mtproto.Messages_Chats, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("messages.getChats - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.MessagesGetChats(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("messages.getChats - reply: {%s}", r)
	return r, err
}

// MessagesGetFullChat
// messages.getFullChat#aeb00b34 chat_id:long = messages.ChatFull;
func (s *Service) MessagesGetFullChat(ctx context.Context, request *mtproto.TLMessagesGetFullChat) (*mtproto.Messages_ChatFull, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("messages.getFullChat - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.MessagesGetFullChat(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("messages.getFullChat - reply: {%s}", r)
	return r, err
}

// MessagesEditChatTitle
// messages.editChatTitle#73783ffd chat_id:long title:string = Updates;
func (s *Service) MessagesEditChatTitle(ctx context.Context, request *mtproto.TLMessagesEditChatTitle) (*mtproto.Updates, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("messages.editChatTitle - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.MessagesEditChatTitle(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("messages.editChatTitle - reply: {%s}", r)
	return r, err
}

// MessagesEditChatPhoto
// messages.editChatPhoto#35ddd674 chat_id:long photo:InputChatPhoto = Updates;
func (s *Service) MessagesEditChatPhoto(ctx context.Context, request *mtproto.TLMessagesEditChatPhoto) (*mtproto.Updates, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("messages.editChatPhoto - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.MessagesEditChatPhoto(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("messages.editChatPhoto - reply: {%s}", r)
	return r, err
}

// MessagesAddChatUserCBC6D107
// messages.addChatUser#cbc6d107 chat_id:long user_id:InputUser fwd_limit:int = messages.InvitedUsers;
func (s *Service) MessagesAddChatUserCBC6D107(ctx context.Context, request *mtproto.TLMessagesAddChatUserCBC6D107) (*mtproto.Messages_InvitedUsers, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("messages.addChatUserCBC6D107 - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.MessagesAddChatUserCBC6D107(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("messages.addChatUserCBC6D107 - reply: {%s}", r)
	return r, err
}

// MessagesDeleteChatUser
// messages.deleteChatUser#a2185cab flags:# revoke_history:flags.0?true chat_id:long user_id:InputUser = Updates;
func (s *Service) MessagesDeleteChatUser(ctx context.Context, request *mtproto.TLMessagesDeleteChatUser) (*mtproto.Updates, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("messages.deleteChatUser - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.MessagesDeleteChatUser(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("messages.deleteChatUser - reply: {%s}", r)
	return r, err
}

// MessagesCreateChat92CEDDD4
// messages.createChat#92ceddd4 flags:# users:Vector<InputUser> title:string ttl_period:flags.0?int = messages.InvitedUsers;
func (s *Service) MessagesCreateChat92CEDDD4(ctx context.Context, request *mtproto.TLMessagesCreateChat92CEDDD4) (*mtproto.Messages_InvitedUsers, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("messages.createChat92CEDDD4 - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.MessagesCreateChat92CEDDD4(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("messages.createChat92CEDDD4 - reply: {%s}", r)
	return r, err
}

// MessagesEditChatAdmin
// messages.editChatAdmin#a85bd1c2 chat_id:long user_id:InputUser is_admin:Bool = Bool;
func (s *Service) MessagesEditChatAdmin(ctx context.Context, request *mtproto.TLMessagesEditChatAdmin) (*mtproto.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("messages.editChatAdmin - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.MessagesEditChatAdmin(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("messages.editChatAdmin - reply: {%s}", r)
	return r, err
}

// MessagesMigrateChat
// messages.migrateChat#a2875319 chat_id:long = Updates;
func (s *Service) MessagesMigrateChat(ctx context.Context, request *mtproto.TLMessagesMigrateChat) (*mtproto.Updates, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("messages.migrateChat - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.MessagesMigrateChat(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("messages.migrateChat - reply: {%s}", r)
	return r, err
}

// MessagesGetCommonChats
// messages.getCommonChats#e40ca104 user_id:InputUser max_id:long limit:int = messages.Chats;
func (s *Service) MessagesGetCommonChats(ctx context.Context, request *mtproto.TLMessagesGetCommonChats) (*mtproto.Messages_Chats, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("messages.getCommonChats - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.MessagesGetCommonChats(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("messages.getCommonChats - reply: {%s}", r)
	return r, err
}

// MessagesEditChatAbout
// messages.editChatAbout#def60797 peer:InputPeer about:string = Bool;
func (s *Service) MessagesEditChatAbout(ctx context.Context, request *mtproto.TLMessagesEditChatAbout) (*mtproto.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("messages.editChatAbout - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.MessagesEditChatAbout(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("messages.editChatAbout - reply: {%s}", r)
	return r, err
}

// MessagesEditChatDefaultBannedRights
// messages.editChatDefaultBannedRights#a5866b41 peer:InputPeer banned_rights:ChatBannedRights = Updates;
func (s *Service) MessagesEditChatDefaultBannedRights(ctx context.Context, request *mtproto.TLMessagesEditChatDefaultBannedRights) (*mtproto.Updates, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("messages.editChatDefaultBannedRights - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.MessagesEditChatDefaultBannedRights(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("messages.editChatDefaultBannedRights - reply: {%s}", r)
	return r, err
}

// MessagesDeleteChat
// messages.deleteChat#5bd0ee50 chat_id:long = Bool;
func (s *Service) MessagesDeleteChat(ctx context.Context, request *mtproto.TLMessagesDeleteChat) (*mtproto.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("messages.deleteChat - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.MessagesDeleteChat(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("messages.deleteChat - reply: {%s}", r)
	return r, err
}

// MessagesGetMessageReadParticipants31C1C44F
// messages.getMessageReadParticipants#31c1c44f peer:InputPeer msg_id:int = Vector<ReadParticipantDate>;
func (s *Service) MessagesGetMessageReadParticipants31C1C44F(ctx context.Context, request *mtproto.TLMessagesGetMessageReadParticipants31C1C44F) (*mtproto.Vector_ReadParticipantDate, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("messages.getMessageReadParticipants31C1C44F - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.MessagesGetMessageReadParticipants31C1C44F(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("messages.getMessageReadParticipants31C1C44F - reply: {%s}", r)
	return r, err
}

// ChannelsConvertToGigagroup
// channels.convertToGigagroup#b290c69 channel:InputChannel = Updates;
func (s *Service) ChannelsConvertToGigagroup(ctx context.Context, request *mtproto.TLChannelsConvertToGigagroup) (*mtproto.Updates, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("channels.convertToGigagroup - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.ChannelsConvertToGigagroup(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("channels.convertToGigagroup - reply: {%s}", r)
	return r, err
}

// ChannelsSetEmojiStickers
// channels.setEmojiStickers#3cd930b7 channel:InputChannel stickerset:InputStickerSet = Bool;
func (s *Service) ChannelsSetEmojiStickers(ctx context.Context, request *mtproto.TLChannelsSetEmojiStickers) (*mtproto.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("channels.setEmojiStickers - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.ChannelsSetEmojiStickers(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("channels.setEmojiStickers - reply: {%s}", r)
	return r, err
}

// MessagesAddChatUserF24753E3
// messages.addChatUser#f24753e3 chat_id:long user_id:InputUser fwd_limit:int = Updates;
func (s *Service) MessagesAddChatUserF24753E3(ctx context.Context, request *mtproto.TLMessagesAddChatUserF24753E3) (*mtproto.Updates, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("messages.addChatUserF24753E3 - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.MessagesAddChatUserF24753E3(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("messages.addChatUserF24753E3 - reply: {%s}", r)
	return r, err
}

// MessagesCreateChat34A818
// messages.createChat#34a818 flags:# users:Vector<InputUser> title:string ttl_period:flags.0?int = Updates;
func (s *Service) MessagesCreateChat34A818(ctx context.Context, request *mtproto.TLMessagesCreateChat34A818) (*mtproto.Updates, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("messages.createChat34A818 - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.MessagesCreateChat34A818(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("messages.createChat34A818 - reply: {%s}", r)
	return r, err
}

// MessagesGetAllChats
// messages.getAllChats#875f74be except_ids:Vector<long> = messages.Chats;
func (s *Service) MessagesGetAllChats(ctx context.Context, request *mtproto.TLMessagesGetAllChats) (*mtproto.Messages_Chats, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("messages.getAllChats - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.MessagesGetAllChats(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("messages.getAllChats - reply: {%s}", r)
	return r, err
}

// MessagesGetMessageReadParticipants2C6F97B7
// messages.getMessageReadParticipants#2c6f97b7 peer:InputPeer msg_id:int = Vector<long>;
func (s *Service) MessagesGetMessageReadParticipants2C6F97B7(ctx context.Context, request *mtproto.TLMessagesGetMessageReadParticipants2C6F97B7) (*mtproto.Vector_Long, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("messages.getMessageReadParticipants2C6F97B7 - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.MessagesGetMessageReadParticipants2C6F97B7(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("messages.getMessageReadParticipants2C6F97B7 - reply: {%s}", r)
	return r, err
}

// MessagesCreateChat9CB126E
// messages.createChat#9cb126e users:Vector<InputUser> title:string = Updates;
func (s *Service) MessagesCreateChat9CB126E(ctx context.Context, request *mtproto.TLMessagesCreateChat9CB126E) (*mtproto.Updates, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("messages.createChat9CB126E - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.MessagesCreateChat9CB126E(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("messages.createChat9CB126E - reply: {%s}", r)
	return r, err
}
