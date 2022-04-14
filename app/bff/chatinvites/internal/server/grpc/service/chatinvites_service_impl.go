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
	"github.com/teamgram/teamgram-server/app/bff/chatinvites/internal/core"
)

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
