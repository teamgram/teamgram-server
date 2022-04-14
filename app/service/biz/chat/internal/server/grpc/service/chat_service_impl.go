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
	"github.com/teamgram/teamgram-server/app/service/biz/chat/chat"
	"github.com/teamgram/teamgram-server/app/service/biz/chat/internal/core"
)

// ChatGetMutableChat
// chat.getMutableChat chat_id:long = MutableChat;
func (s *Service) ChatGetMutableChat(ctx context.Context, request *chat.TLChatGetMutableChat) (*chat.MutableChat, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("chat.getMutableChat - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.ChatGetMutableChat(request)
	if err != nil {
		return nil, err
	}

	c.Infof("chat.getMutableChat - reply: %s", r.DebugString())
	return r, err
}

// ChatGetChatListByIdList
// chat.getChatListByIdList self_id:long id_list:Vector<long> = Vector<MutableChat>;
func (s *Service) ChatGetChatListByIdList(ctx context.Context, request *chat.TLChatGetChatListByIdList) (*chat.Vector_MutableChat, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("chat.getChatListByIdList - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.ChatGetChatListByIdList(request)
	if err != nil {
		return nil, err
	}

	c.Infof("chat.getChatListByIdList - reply: %s", r.DebugString())
	return r, err
}

// ChatGetChatBySelfId
// chat.getChatBySelfId self_id:long chat_id:long = MutableChat;
func (s *Service) ChatGetChatBySelfId(ctx context.Context, request *chat.TLChatGetChatBySelfId) (*chat.MutableChat, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("chat.getChatBySelfId - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.ChatGetChatBySelfId(request)
	if err != nil {
		return nil, err
	}

	c.Infof("chat.getChatBySelfId - reply: %s", r.DebugString())
	return r, err
}

// ChatCreateChat2
// chat.createChat2 creator_id:long user_id_list:Vector<long> title:string = MutableChat;
func (s *Service) ChatCreateChat2(ctx context.Context, request *chat.TLChatCreateChat2) (*chat.MutableChat, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("chat.createChat2 - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.ChatCreateChat2(request)
	if err != nil {
		return nil, err
	}

	c.Infof("chat.createChat2 - reply: %s", r.DebugString())
	return r, err
}

// ChatDeleteChat
// chat.deleteChat chat_id:long operator_id:long = MutableChat;
func (s *Service) ChatDeleteChat(ctx context.Context, request *chat.TLChatDeleteChat) (*chat.MutableChat, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("chat.deleteChat - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.ChatDeleteChat(request)
	if err != nil {
		return nil, err
	}

	c.Infof("chat.deleteChat - reply: %s", r.DebugString())
	return r, err
}

// ChatDeleteChatUser
// chat.deleteChatUser chat_id:long operator_id:long delete_user_id:long = MutableChat;
func (s *Service) ChatDeleteChatUser(ctx context.Context, request *chat.TLChatDeleteChatUser) (*chat.MutableChat, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("chat.deleteChatUser - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.ChatDeleteChatUser(request)
	if err != nil {
		return nil, err
	}

	c.Infof("chat.deleteChatUser - reply: %s", r.DebugString())
	return r, err
}

// ChatEditChatTitle
// chat.editChatTitle chat_id:long edit_user_id:long title:string = MutableChat;
func (s *Service) ChatEditChatTitle(ctx context.Context, request *chat.TLChatEditChatTitle) (*chat.MutableChat, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("chat.editChatTitle - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.ChatEditChatTitle(request)
	if err != nil {
		return nil, err
	}

	c.Infof("chat.editChatTitle - reply: %s", r.DebugString())
	return r, err
}

// ChatEditChatAbout
// chat.editChatAbout chat_id:long edit_user_id:long about:string = MutableChat;
func (s *Service) ChatEditChatAbout(ctx context.Context, request *chat.TLChatEditChatAbout) (*chat.MutableChat, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("chat.editChatAbout - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.ChatEditChatAbout(request)
	if err != nil {
		return nil, err
	}

	c.Infof("chat.editChatAbout - reply: %s", r.DebugString())
	return r, err
}

// ChatEditChatPhoto
// chat.editChatPhoto chat_id:long edit_user_id:long chat_photo:Photo = MutableChat;
func (s *Service) ChatEditChatPhoto(ctx context.Context, request *chat.TLChatEditChatPhoto) (*chat.MutableChat, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("chat.editChatPhoto - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.ChatEditChatPhoto(request)
	if err != nil {
		return nil, err
	}

	c.Infof("chat.editChatPhoto - reply: %s", r.DebugString())
	return r, err
}

// ChatEditChatAdmin
// chat.editChatAdmin chat_id:long operator_id:long edit_chat_admin_id:long is_admin:Bool = MutableChat;
func (s *Service) ChatEditChatAdmin(ctx context.Context, request *chat.TLChatEditChatAdmin) (*chat.MutableChat, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("chat.editChatAdmin - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.ChatEditChatAdmin(request)
	if err != nil {
		return nil, err
	}

	c.Infof("chat.editChatAdmin - reply: %s", r.DebugString())
	return r, err
}

// ChatEditChatDefaultBannedRights
// chat.editChatDefaultBannedRights chat_id:long operator_id:long banned_rights:ChatBannedRights = MutableChat;
func (s *Service) ChatEditChatDefaultBannedRights(ctx context.Context, request *chat.TLChatEditChatDefaultBannedRights) (*chat.MutableChat, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("chat.editChatDefaultBannedRights - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.ChatEditChatDefaultBannedRights(request)
	if err != nil {
		return nil, err
	}

	c.Infof("chat.editChatDefaultBannedRights - reply: %s", r.DebugString())
	return r, err
}

// ChatAddChatUser
// chat.addChatUser chat_id:long inviter_id:long user_id:long = MutableChat;
func (s *Service) ChatAddChatUser(ctx context.Context, request *chat.TLChatAddChatUser) (*chat.MutableChat, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("chat.addChatUser - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.ChatAddChatUser(request)
	if err != nil {
		return nil, err
	}

	c.Infof("chat.addChatUser - reply: %s", r.DebugString())
	return r, err
}

// ChatGetMutableChatByLink
// chat.getMutableChatByLink link:string = MutableChat;
func (s *Service) ChatGetMutableChatByLink(ctx context.Context, request *chat.TLChatGetMutableChatByLink) (*chat.MutableChat, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("chat.getMutableChatByLink - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.ChatGetMutableChatByLink(request)
	if err != nil {
		return nil, err
	}

	c.Infof("chat.getMutableChatByLink - reply: %s", r.DebugString())
	return r, err
}

// ChatToggleNoForwards
// chat.toggleNoForwards chat_id:long operator_id:long enabled:Bool = MutableChat;
func (s *Service) ChatToggleNoForwards(ctx context.Context, request *chat.TLChatToggleNoForwards) (*chat.MutableChat, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("chat.toggleNoForwards - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.ChatToggleNoForwards(request)
	if err != nil {
		return nil, err
	}

	c.Infof("chat.toggleNoForwards - reply: %s", r.DebugString())
	return r, err
}

// ChatMigratedToChannel
// chat.migratedToChannel chat:MutableChat id:long access_hash:long = Bool;
func (s *Service) ChatMigratedToChannel(ctx context.Context, request *chat.TLChatMigratedToChannel) (*mtproto.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("chat.migratedToChannel - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.ChatMigratedToChannel(request)
	if err != nil {
		return nil, err
	}

	c.Infof("chat.migratedToChannel - reply: %s", r.DebugString())
	return r, err
}

// ChatGetChatParticipantIdList
// chat.getChatParticipantIdList chat_id:long = Vector<long>;
func (s *Service) ChatGetChatParticipantIdList(ctx context.Context, request *chat.TLChatGetChatParticipantIdList) (*chat.Vector_Long, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("chat.getChatParticipantIdList - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.ChatGetChatParticipantIdList(request)
	if err != nil {
		return nil, err
	}

	c.Infof("chat.getChatParticipantIdList - reply: %s", r.DebugString())
	return r, err
}

// ChatGetUsersChatIdList
// chat.getUsersChatIdList id:Vector<long> = Vector<UserChatIdList>;
func (s *Service) ChatGetUsersChatIdList(ctx context.Context, request *chat.TLChatGetUsersChatIdList) (*chat.Vector_UserChatIdList, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("chat.getUsersChatIdList - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.ChatGetUsersChatIdList(request)
	if err != nil {
		return nil, err
	}

	c.Infof("chat.getUsersChatIdList - reply: %s", r.DebugString())
	return r, err
}

// ChatGetMyChatList
// chat.getMyChatList user_id:long is_creator:Bool = Vector<MutableChat>;
func (s *Service) ChatGetMyChatList(ctx context.Context, request *chat.TLChatGetMyChatList) (*chat.Vector_MutableChat, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("chat.getMyChatList - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.ChatGetMyChatList(request)
	if err != nil {
		return nil, err
	}

	c.Infof("chat.getMyChatList - reply: %s", r.DebugString())
	return r, err
}

// ChatExportChatInvite
// chat.exportChatInvite flags:# chat_id:long admin_id:long legacy_revoke_permanent:flags.2?true request_needed:flags.3?true expire_date:flags.0?int usage_limit:flags.1?int title:flags.4?string = ExportedChatInvite;
func (s *Service) ChatExportChatInvite(ctx context.Context, request *chat.TLChatExportChatInvite) (*mtproto.ExportedChatInvite, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("chat.exportChatInvite - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.ChatExportChatInvite(request)
	if err != nil {
		return nil, err
	}

	c.Infof("chat.exportChatInvite - reply: %s", r.DebugString())
	return r, err
}

// ChatGetAdminsWithInvites
// chat.getAdminsWithInvites self_id:long chat_id:long = Vector<ChatAdminWithInvites>;
func (s *Service) ChatGetAdminsWithInvites(ctx context.Context, request *chat.TLChatGetAdminsWithInvites) (*chat.Vector_ChatAdminWithInvites, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("chat.getAdminsWithInvites - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.ChatGetAdminsWithInvites(request)
	if err != nil {
		return nil, err
	}

	c.Infof("chat.getAdminsWithInvites - reply: %s", r.DebugString())
	return r, err
}

// ChatGetExportedChatInvite
// chat.getExportedChatInvite chat_id:long link:string = ExportedChatInvite;
func (s *Service) ChatGetExportedChatInvite(ctx context.Context, request *chat.TLChatGetExportedChatInvite) (*mtproto.ExportedChatInvite, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("chat.getExportedChatInvite - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.ChatGetExportedChatInvite(request)
	if err != nil {
		return nil, err
	}

	c.Infof("chat.getExportedChatInvite - reply: %s", r.DebugString())
	return r, err
}

// ChatGetExportedChatInvites
// chat.getExportedChatInvites flags:# chat_id:long admin_id:long revoked:flags.3?true offset_date:flags.2?int offset_link:flags.2?string limit:int = Vector<ExportedChatInvite>;
func (s *Service) ChatGetExportedChatInvites(ctx context.Context, request *chat.TLChatGetExportedChatInvites) (*chat.Vector_ExportedChatInvite, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("chat.getExportedChatInvites - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.ChatGetExportedChatInvites(request)
	if err != nil {
		return nil, err
	}

	c.Infof("chat.getExportedChatInvites - reply: %s", r.DebugString())
	return r, err
}

// ChatCheckChatInvite
// chat.checkChatInvite self_id:long hash:string = ChatInviteExt;
func (s *Service) ChatCheckChatInvite(ctx context.Context, request *chat.TLChatCheckChatInvite) (*chat.ChatInviteExt, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("chat.checkChatInvite - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.ChatCheckChatInvite(request)
	if err != nil {
		return nil, err
	}

	c.Infof("chat.checkChatInvite - reply: %s", r.DebugString())
	return r, err
}

// ChatImportChatInvite
// chat.importChatInvite self_id:long hash:string = MutableChat;
func (s *Service) ChatImportChatInvite(ctx context.Context, request *chat.TLChatImportChatInvite) (*chat.MutableChat, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("chat.importChatInvite - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.ChatImportChatInvite(request)
	if err != nil {
		return nil, err
	}

	c.Infof("chat.importChatInvite - reply: %s", r.DebugString())
	return r, err
}

// ChatGetChatInviteImporters
// chat.getChatInviteImporters flags:# self_id:long chat_id:long requested:flags.0?true link:flags.1?string q:flags.2?string offset_date:int offset_user:long limit:int = Vector<ChatInviteImporter>;
func (s *Service) ChatGetChatInviteImporters(ctx context.Context, request *chat.TLChatGetChatInviteImporters) (*chat.Vector_ChatInviteImporter, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("chat.getChatInviteImporters - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.ChatGetChatInviteImporters(request)
	if err != nil {
		return nil, err
	}

	c.Infof("chat.getChatInviteImporters - reply: %s", r.DebugString())
	return r, err
}

// ChatDeleteExportedChatInvite
// chat.deleteExportedChatInvite self_id:long chat_id:long link:string = Bool;
func (s *Service) ChatDeleteExportedChatInvite(ctx context.Context, request *chat.TLChatDeleteExportedChatInvite) (*mtproto.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("chat.deleteExportedChatInvite - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.ChatDeleteExportedChatInvite(request)
	if err != nil {
		return nil, err
	}

	c.Infof("chat.deleteExportedChatInvite - reply: %s", r.DebugString())
	return r, err
}

// ChatDeleteRevokedExportedChatInvites
// chat.deleteRevokedExportedChatInvites self_id:long chat_id:long admin_id:long = Bool;
func (s *Service) ChatDeleteRevokedExportedChatInvites(ctx context.Context, request *chat.TLChatDeleteRevokedExportedChatInvites) (*mtproto.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("chat.deleteRevokedExportedChatInvites - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.ChatDeleteRevokedExportedChatInvites(request)
	if err != nil {
		return nil, err
	}

	c.Infof("chat.deleteRevokedExportedChatInvites - reply: %s", r.DebugString())
	return r, err
}

// ChatEditExportedChatInvite
// chat.editExportedChatInvite flags:# self_id:long chat_id:long revoked:flags.2?true link:string expire_date:flags.0?int usage_limit:flags.1?int request_needed:flags.3?Bool title:flags.4?string = Vector<ExportedChatInvite>;
func (s *Service) ChatEditExportedChatInvite(ctx context.Context, request *chat.TLChatEditExportedChatInvite) (*chat.Vector_ExportedChatInvite, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("chat.editExportedChatInvite - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.ChatEditExportedChatInvite(request)
	if err != nil {
		return nil, err
	}

	c.Infof("chat.editExportedChatInvite - reply: %s", r.DebugString())
	return r, err
}

// ChatSetChatAvailableReactions
// chat.setChatAvailableReactions self_id:long chat_id:long available_reactions:Vector<string> = MutableChat;
func (s *Service) ChatSetChatAvailableReactions(ctx context.Context, request *chat.TLChatSetChatAvailableReactions) (*chat.MutableChat, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("chat.setChatAvailableReactions - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.ChatSetChatAvailableReactions(request)
	if err != nil {
		return nil, err
	}

	c.Infof("chat.setChatAvailableReactions - reply: %s", r.DebugString())
	return r, err
}
