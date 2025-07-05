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

	"github.com/teamgram/proto/v2/tg"
	"github.com/teamgram/teamgram-server/v2/app/bff/chats/internal/core"
)

// MessagesGetChats
// messages.getChats#49e9528f id:Vector<long> = messages.Chats;
func (s *Service) MessagesGetChats(ctx context.Context, request *tg.TLMessagesGetChats) (*tg.MessagesChats, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("messages.getChats - metadata: %s, request: %s", c.MD, request)

	r, err := c.MessagesGetChats(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("messages.getChats - reply: %s", r)
	return r, err
}

// MessagesGetFullChat
// messages.getFullChat#aeb00b34 chat_id:long = messages.ChatFull;
func (s *Service) MessagesGetFullChat(ctx context.Context, request *tg.TLMessagesGetFullChat) (*tg.MessagesChatFull, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("messages.getFullChat - metadata: %s, request: %s", c.MD, request)

	r, err := c.MessagesGetFullChat(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("messages.getFullChat - reply: %s", r)
	return r, err
}

// MessagesEditChatTitle
// messages.editChatTitle#73783ffd chat_id:long title:string = Updates;
func (s *Service) MessagesEditChatTitle(ctx context.Context, request *tg.TLMessagesEditChatTitle) (*tg.Updates, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("messages.editChatTitle - metadata: %s, request: %s", c.MD, request)

	r, err := c.MessagesEditChatTitle(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("messages.editChatTitle - reply: %s", r)
	return r, err
}

// MessagesEditChatPhoto
// messages.editChatPhoto#35ddd674 chat_id:long photo:InputChatPhoto = Updates;
func (s *Service) MessagesEditChatPhoto(ctx context.Context, request *tg.TLMessagesEditChatPhoto) (*tg.Updates, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("messages.editChatPhoto - metadata: %s, request: %s", c.MD, request)

	r, err := c.MessagesEditChatPhoto(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("messages.editChatPhoto - reply: %s", r)
	return r, err
}

// MessagesAddChatUser
// messages.addChatUser#cbc6d107 chat_id:long user_id:InputUser fwd_limit:int = messages.InvitedUsers;
func (s *Service) MessagesAddChatUser(ctx context.Context, request *tg.TLMessagesAddChatUser) (*tg.MessagesInvitedUsers, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("messages.addChatUser - metadata: %s, request: %s", c.MD, request)

	r, err := c.MessagesAddChatUser(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("messages.addChatUser - reply: %s", r)
	return r, err
}

// MessagesDeleteChatUser
// messages.deleteChatUser#a2185cab flags:# revoke_history:flags.0?true chat_id:long user_id:InputUser = Updates;
func (s *Service) MessagesDeleteChatUser(ctx context.Context, request *tg.TLMessagesDeleteChatUser) (*tg.Updates, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("messages.deleteChatUser - metadata: %s, request: %s", c.MD, request)

	r, err := c.MessagesDeleteChatUser(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("messages.deleteChatUser - reply: %s", r)
	return r, err
}

// MessagesCreateChat
// messages.createChat#92ceddd4 flags:# users:Vector<InputUser> title:string ttl_period:flags.0?int = messages.InvitedUsers;
func (s *Service) MessagesCreateChat(ctx context.Context, request *tg.TLMessagesCreateChat) (*tg.MessagesInvitedUsers, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("messages.createChat - metadata: %s, request: %s", c.MD, request)

	r, err := c.MessagesCreateChat(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("messages.createChat - reply: %s", r)
	return r, err
}

// MessagesEditChatAdmin
// messages.editChatAdmin#a85bd1c2 chat_id:long user_id:InputUser is_admin:Bool = Bool;
func (s *Service) MessagesEditChatAdmin(ctx context.Context, request *tg.TLMessagesEditChatAdmin) (*tg.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("messages.editChatAdmin - metadata: %s, request: %s", c.MD, request)

	r, err := c.MessagesEditChatAdmin(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("messages.editChatAdmin - reply: %s", r)
	return r, err
}

// MessagesMigrateChat
// messages.migrateChat#a2875319 chat_id:long = Updates;
func (s *Service) MessagesMigrateChat(ctx context.Context, request *tg.TLMessagesMigrateChat) (*tg.Updates, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("messages.migrateChat - metadata: %s, request: %s", c.MD, request)

	r, err := c.MessagesMigrateChat(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("messages.migrateChat - reply: %s", r)
	return r, err
}

// MessagesGetCommonChats
// messages.getCommonChats#e40ca104 user_id:InputUser max_id:long limit:int = messages.Chats;
func (s *Service) MessagesGetCommonChats(ctx context.Context, request *tg.TLMessagesGetCommonChats) (*tg.MessagesChats, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("messages.getCommonChats - metadata: %s, request: %s", c.MD, request)

	r, err := c.MessagesGetCommonChats(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("messages.getCommonChats - reply: %s", r)
	return r, err
}

// MessagesEditChatAbout
// messages.editChatAbout#def60797 peer:InputPeer about:string = Bool;
func (s *Service) MessagesEditChatAbout(ctx context.Context, request *tg.TLMessagesEditChatAbout) (*tg.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("messages.editChatAbout - metadata: %s, request: %s", c.MD, request)

	r, err := c.MessagesEditChatAbout(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("messages.editChatAbout - reply: %s", r)
	return r, err
}

// MessagesEditChatDefaultBannedRights
// messages.editChatDefaultBannedRights#a5866b41 peer:InputPeer banned_rights:ChatBannedRights = Updates;
func (s *Service) MessagesEditChatDefaultBannedRights(ctx context.Context, request *tg.TLMessagesEditChatDefaultBannedRights) (*tg.Updates, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("messages.editChatDefaultBannedRights - metadata: %s, request: %s", c.MD, request)

	r, err := c.MessagesEditChatDefaultBannedRights(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("messages.editChatDefaultBannedRights - reply: %s", r)
	return r, err
}

// MessagesDeleteChat
// messages.deleteChat#5bd0ee50 chat_id:long = Bool;
func (s *Service) MessagesDeleteChat(ctx context.Context, request *tg.TLMessagesDeleteChat) (*tg.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("messages.deleteChat - metadata: %s, request: %s", c.MD, request)

	r, err := c.MessagesDeleteChat(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("messages.deleteChat - reply: %s", r)
	return r, err
}

// MessagesGetMessageReadParticipants
// messages.getMessageReadParticipants#31c1c44f peer:InputPeer msg_id:int = Vector<ReadParticipantDate>;
func (s *Service) MessagesGetMessageReadParticipants(ctx context.Context, request *tg.TLMessagesGetMessageReadParticipants) (*tg.VectorReadParticipantDate, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("messages.getMessageReadParticipants - metadata: %s, request: %s", c.MD, request)

	r, err := c.MessagesGetMessageReadParticipants(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("messages.getMessageReadParticipants - reply: %s", r)
	return r, err
}

// ChannelsConvertToGigagroup
// channels.convertToGigagroup#b290c69 channel:InputChannel = Updates;
func (s *Service) ChannelsConvertToGigagroup(ctx context.Context, request *tg.TLChannelsConvertToGigagroup) (*tg.Updates, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("channels.convertToGigagroup - metadata: %s, request: %s", c.MD, request)

	r, err := c.ChannelsConvertToGigagroup(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("channels.convertToGigagroup - reply: %s", r)
	return r, err
}

// ChannelsSetEmojiStickers
// channels.setEmojiStickers#3cd930b7 channel:InputChannel stickerset:InputStickerSet = Bool;
func (s *Service) ChannelsSetEmojiStickers(ctx context.Context, request *tg.TLChannelsSetEmojiStickers) (*tg.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("channels.setEmojiStickers - metadata: %s, request: %s", c.MD, request)

	r, err := c.ChannelsSetEmojiStickers(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("channels.setEmojiStickers - reply: %s", r)
	return r, err
}
