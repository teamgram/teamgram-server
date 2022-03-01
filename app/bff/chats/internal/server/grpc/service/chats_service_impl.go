/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright 2022 Teamgram Authors.
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
	c.Infof("messages.getChats - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.MessagesGetChats(request)
	if err != nil {
		return nil, err
	}

	c.Infof("messages.getChats - reply: %s", r.DebugString())
	return r, err
}

// MessagesGetFullChat
// messages.getFullChat#aeb00b34 chat_id:long = messages.ChatFull;
func (s *Service) MessagesGetFullChat(ctx context.Context, request *mtproto.TLMessagesGetFullChat) (*mtproto.Messages_ChatFull, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("messages.getFullChat - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.MessagesGetFullChat(request)
	if err != nil {
		return nil, err
	}

	c.Infof("messages.getFullChat - reply: %s", r.DebugString())
	return r, err
}

// MessagesEditChatTitle
// messages.editChatTitle#73783ffd chat_id:long title:string = Updates;
func (s *Service) MessagesEditChatTitle(ctx context.Context, request *mtproto.TLMessagesEditChatTitle) (*mtproto.Updates, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("messages.editChatTitle - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.MessagesEditChatTitle(request)
	if err != nil {
		return nil, err
	}

	c.Infof("messages.editChatTitle - reply: %s", r.DebugString())
	return r, err
}

// MessagesEditChatPhoto
// messages.editChatPhoto#35ddd674 chat_id:long photo:InputChatPhoto = Updates;
func (s *Service) MessagesEditChatPhoto(ctx context.Context, request *mtproto.TLMessagesEditChatPhoto) (*mtproto.Updates, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("messages.editChatPhoto - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.MessagesEditChatPhoto(request)
	if err != nil {
		return nil, err
	}

	c.Infof("messages.editChatPhoto - reply: %s", r.DebugString())
	return r, err
}

// MessagesAddChatUser
// messages.addChatUser#f24753e3 chat_id:long user_id:InputUser fwd_limit:int = Updates;
func (s *Service) MessagesAddChatUser(ctx context.Context, request *mtproto.TLMessagesAddChatUser) (*mtproto.Updates, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("messages.addChatUser - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.MessagesAddChatUser(request)
	if err != nil {
		return nil, err
	}

	c.Infof("messages.addChatUser - reply: %s", r.DebugString())
	return r, err
}

// MessagesDeleteChatUser
// messages.deleteChatUser#a2185cab flags:# revoke_history:flags.0?true chat_id:long user_id:InputUser = Updates;
func (s *Service) MessagesDeleteChatUser(ctx context.Context, request *mtproto.TLMessagesDeleteChatUser) (*mtproto.Updates, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("messages.deleteChatUser - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.MessagesDeleteChatUser(request)
	if err != nil {
		return nil, err
	}

	c.Infof("messages.deleteChatUser - reply: %s", r.DebugString())
	return r, err
}

// MessagesCreateChat
// messages.createChat#9cb126e users:Vector<InputUser> title:string = Updates;
func (s *Service) MessagesCreateChat(ctx context.Context, request *mtproto.TLMessagesCreateChat) (*mtproto.Updates, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("messages.createChat - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.MessagesCreateChat(request)
	if err != nil {
		return nil, err
	}

	c.Infof("messages.createChat - reply: %s", r.DebugString())
	return r, err
}

// MessagesExportChatInvite
// messages.exportChatInvite#a02ce5d5 flags:# legacy_revoke_permanent:flags.2?true request_needed:flags.3?true peer:InputPeer expire_date:flags.0?int usage_limit:flags.1?int title:flags.4?string = ExportedChatInvite;
func (s *Service) MessagesExportChatInvite(ctx context.Context, request *mtproto.TLMessagesExportChatInvite) (*mtproto.ExportedChatInvite, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("messages.exportChatInvite - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.MessagesExportChatInvite(request)
	if err != nil {
		return nil, err
	}

	c.Infof("messages.exportChatInvite - reply: %s", r.DebugString())
	return r, err
}

// MessagesCheckChatInvite
// messages.checkChatInvite#3eadb1bb hash:string = ChatInvite;
func (s *Service) MessagesCheckChatInvite(ctx context.Context, request *mtproto.TLMessagesCheckChatInvite) (*mtproto.ChatInvite, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("messages.checkChatInvite - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.MessagesCheckChatInvite(request)
	if err != nil {
		return nil, err
	}

	c.Infof("messages.checkChatInvite - reply: %s", r.DebugString())
	return r, err
}

// MessagesImportChatInvite
// messages.importChatInvite#6c50051c hash:string = Updates;
func (s *Service) MessagesImportChatInvite(ctx context.Context, request *mtproto.TLMessagesImportChatInvite) (*mtproto.Updates, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("messages.importChatInvite - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.MessagesImportChatInvite(request)
	if err != nil {
		return nil, err
	}

	c.Infof("messages.importChatInvite - reply: %s", r.DebugString())
	return r, err
}

// MessagesEditChatAdmin
// messages.editChatAdmin#a85bd1c2 chat_id:long user_id:InputUser is_admin:Bool = Bool;
func (s *Service) MessagesEditChatAdmin(ctx context.Context, request *mtproto.TLMessagesEditChatAdmin) (*mtproto.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("messages.editChatAdmin - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.MessagesEditChatAdmin(request)
	if err != nil {
		return nil, err
	}

	c.Infof("messages.editChatAdmin - reply: %s", r.DebugString())
	return r, err
}

// MessagesMigrateChat
// messages.migrateChat#a2875319 chat_id:long = Updates;
func (s *Service) MessagesMigrateChat(ctx context.Context, request *mtproto.TLMessagesMigrateChat) (*mtproto.Updates, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("messages.migrateChat - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.MessagesMigrateChat(request)
	if err != nil {
		return nil, err
	}

	c.Infof("messages.migrateChat - reply: %s", r.DebugString())
	return r, err
}

// MessagesGetCommonChats
// messages.getCommonChats#e40ca104 user_id:InputUser max_id:long limit:int = messages.Chats;
func (s *Service) MessagesGetCommonChats(ctx context.Context, request *mtproto.TLMessagesGetCommonChats) (*mtproto.Messages_Chats, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("messages.getCommonChats - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.MessagesGetCommonChats(request)
	if err != nil {
		return nil, err
	}

	c.Infof("messages.getCommonChats - reply: %s", r.DebugString())
	return r, err
}

// MessagesGetAllChats
// messages.getAllChats#875f74be except_ids:Vector<long> = messages.Chats;
func (s *Service) MessagesGetAllChats(ctx context.Context, request *mtproto.TLMessagesGetAllChats) (*mtproto.Messages_Chats, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("messages.getAllChats - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.MessagesGetAllChats(request)
	if err != nil {
		return nil, err
	}

	c.Infof("messages.getAllChats - reply: %s", r.DebugString())
	return r, err
}

// MessagesEditChatAbout
// messages.editChatAbout#def60797 peer:InputPeer about:string = Bool;
func (s *Service) MessagesEditChatAbout(ctx context.Context, request *mtproto.TLMessagesEditChatAbout) (*mtproto.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("messages.editChatAbout - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.MessagesEditChatAbout(request)
	if err != nil {
		return nil, err
	}

	c.Infof("messages.editChatAbout - reply: %s", r.DebugString())
	return r, err
}

// MessagesEditChatDefaultBannedRights
// messages.editChatDefaultBannedRights#a5866b41 peer:InputPeer banned_rights:ChatBannedRights = Updates;
func (s *Service) MessagesEditChatDefaultBannedRights(ctx context.Context, request *mtproto.TLMessagesEditChatDefaultBannedRights) (*mtproto.Updates, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("messages.editChatDefaultBannedRights - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.MessagesEditChatDefaultBannedRights(request)
	if err != nil {
		return nil, err
	}

	c.Infof("messages.editChatDefaultBannedRights - reply: %s", r.DebugString())
	return r, err
}

// MessagesDeleteChat
// messages.deleteChat#5bd0ee50 chat_id:long = Bool;
func (s *Service) MessagesDeleteChat(ctx context.Context, request *mtproto.TLMessagesDeleteChat) (*mtproto.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("messages.deleteChat - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.MessagesDeleteChat(request)
	if err != nil {
		return nil, err
	}

	c.Infof("messages.deleteChat - reply: %s", r.DebugString())
	return r, err
}

// MessagesGetExportedChatInvites
// messages.getExportedChatInvites#a2b5a3f6 flags:# revoked:flags.3?true peer:InputPeer admin_id:InputUser offset_date:flags.2?int offset_link:flags.2?string limit:int = messages.ExportedChatInvites;
func (s *Service) MessagesGetExportedChatInvites(ctx context.Context, request *mtproto.TLMessagesGetExportedChatInvites) (*mtproto.Messages_ExportedChatInvites, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("messages.getExportedChatInvites - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.MessagesGetExportedChatInvites(request)
	if err != nil {
		return nil, err
	}

	c.Infof("messages.getExportedChatInvites - reply: %s", r.DebugString())
	return r, err
}

// MessagesGetExportedChatInvite
// messages.getExportedChatInvite#73746f5c peer:InputPeer link:string = messages.ExportedChatInvite;
func (s *Service) MessagesGetExportedChatInvite(ctx context.Context, request *mtproto.TLMessagesGetExportedChatInvite) (*mtproto.Messages_ExportedChatInvite, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("messages.getExportedChatInvite - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.MessagesGetExportedChatInvite(request)
	if err != nil {
		return nil, err
	}

	c.Infof("messages.getExportedChatInvite - reply: %s", r.DebugString())
	return r, err
}

// MessagesEditExportedChatInvite
// messages.editExportedChatInvite#bdca2f75 flags:# revoked:flags.2?true peer:InputPeer link:string expire_date:flags.0?int usage_limit:flags.1?int request_needed:flags.3?Bool title:flags.4?string = messages.ExportedChatInvite;
func (s *Service) MessagesEditExportedChatInvite(ctx context.Context, request *mtproto.TLMessagesEditExportedChatInvite) (*mtproto.Messages_ExportedChatInvite, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("messages.editExportedChatInvite - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.MessagesEditExportedChatInvite(request)
	if err != nil {
		return nil, err
	}

	c.Infof("messages.editExportedChatInvite - reply: %s", r.DebugString())
	return r, err
}

// MessagesDeleteRevokedExportedChatInvites
// messages.deleteRevokedExportedChatInvites#56987bd5 peer:InputPeer admin_id:InputUser = Bool;
func (s *Service) MessagesDeleteRevokedExportedChatInvites(ctx context.Context, request *mtproto.TLMessagesDeleteRevokedExportedChatInvites) (*mtproto.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("messages.deleteRevokedExportedChatInvites - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.MessagesDeleteRevokedExportedChatInvites(request)
	if err != nil {
		return nil, err
	}

	c.Infof("messages.deleteRevokedExportedChatInvites - reply: %s", r.DebugString())
	return r, err
}

// MessagesDeleteExportedChatInvite
// messages.deleteExportedChatInvite#d464a42b peer:InputPeer link:string = Bool;
func (s *Service) MessagesDeleteExportedChatInvite(ctx context.Context, request *mtproto.TLMessagesDeleteExportedChatInvite) (*mtproto.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("messages.deleteExportedChatInvite - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.MessagesDeleteExportedChatInvite(request)
	if err != nil {
		return nil, err
	}

	c.Infof("messages.deleteExportedChatInvite - reply: %s", r.DebugString())
	return r, err
}

// MessagesGetAdminsWithInvites
// messages.getAdminsWithInvites#3920e6ef peer:InputPeer = messages.ChatAdminsWithInvites;
func (s *Service) MessagesGetAdminsWithInvites(ctx context.Context, request *mtproto.TLMessagesGetAdminsWithInvites) (*mtproto.Messages_ChatAdminsWithInvites, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("messages.getAdminsWithInvites - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.MessagesGetAdminsWithInvites(request)
	if err != nil {
		return nil, err
	}

	c.Infof("messages.getAdminsWithInvites - reply: %s", r.DebugString())
	return r, err
}

// MessagesGetChatInviteImporters
// messages.getChatInviteImporters#df04dd4e flags:# requested:flags.0?true peer:InputPeer link:flags.1?string q:flags.2?string offset_date:int offset_user:InputUser limit:int = messages.ChatInviteImporters;
func (s *Service) MessagesGetChatInviteImporters(ctx context.Context, request *mtproto.TLMessagesGetChatInviteImporters) (*mtproto.Messages_ChatInviteImporters, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("messages.getChatInviteImporters - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.MessagesGetChatInviteImporters(request)
	if err != nil {
		return nil, err
	}

	c.Infof("messages.getChatInviteImporters - reply: %s", r.DebugString())
	return r, err
}

// MessagesGetMessageReadParticipants
// messages.getMessageReadParticipants#2c6f97b7 peer:InputPeer msg_id:int = Vector<long>;
func (s *Service) MessagesGetMessageReadParticipants(ctx context.Context, request *mtproto.TLMessagesGetMessageReadParticipants) (*mtproto.Vector_Long, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("messages.getMessageReadParticipants - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.MessagesGetMessageReadParticipants(request)
	if err != nil {
		return nil, err
	}

	c.Infof("messages.getMessageReadParticipants - reply: %s", r.DebugString())
	return r, err
}

// MessagesHideChatJoinRequest
// messages.hideChatJoinRequest#7fe7e815 flags:# approved:flags.0?true peer:InputPeer user_id:InputUser = Updates;
func (s *Service) MessagesHideChatJoinRequest(ctx context.Context, request *mtproto.TLMessagesHideChatJoinRequest) (*mtproto.Updates, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("messages.hideChatJoinRequest - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.MessagesHideChatJoinRequest(request)
	if err != nil {
		return nil, err
	}

	c.Infof("messages.hideChatJoinRequest - reply: %s", r.DebugString())
	return r, err
}

// MessagesHideAllChatJoinRequests
// messages.hideAllChatJoinRequests#e085f4ea flags:# approved:flags.0?true peer:InputPeer link:flags.1?string = Updates;
func (s *Service) MessagesHideAllChatJoinRequests(ctx context.Context, request *mtproto.TLMessagesHideAllChatJoinRequests) (*mtproto.Updates, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("messages.hideAllChatJoinRequests - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.MessagesHideAllChatJoinRequests(request)
	if err != nil {
		return nil, err
	}

	c.Infof("messages.hideAllChatJoinRequests - reply: %s", r.DebugString())
	return r, err
}

// ChannelsConvertToGigagroup
// channels.convertToGigagroup#b290c69 channel:InputChannel = Updates;
func (s *Service) ChannelsConvertToGigagroup(ctx context.Context, request *mtproto.TLChannelsConvertToGigagroup) (*mtproto.Updates, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("channels.convertToGigagroup - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.ChannelsConvertToGigagroup(request)
	if err != nil {
		return nil, err
	}

	c.Infof("channels.convertToGigagroup - reply: %s", r.DebugString())
	return r, err
}
