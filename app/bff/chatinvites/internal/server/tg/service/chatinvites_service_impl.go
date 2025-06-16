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
	"github.com/teamgram/teamgram-server/v2/app/bff/chatinvites/internal/core"
)

// MessagesExportChatInvite
// messages.exportChatInvite#a455de90 flags:# legacy_revoke_permanent:flags.2?true request_needed:flags.3?true peer:InputPeer expire_date:flags.0?int usage_limit:flags.1?int title:flags.4?string subscription_pricing:flags.5?StarsSubscriptionPricing = ExportedChatInvite;
func (s *Service) MessagesExportChatInvite(ctx context.Context, request *tg.TLMessagesExportChatInvite) (*tg.ExportedChatInvite, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("messages.exportChatInvite - metadata: {}, request: {%v}", request)

	r, err := c.MessagesExportChatInvite(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("messages.exportChatInvite - reply: {%v}", r)
	return r, err
}

// MessagesCheckChatInvite
// messages.checkChatInvite#3eadb1bb hash:string = ChatInvite;
func (s *Service) MessagesCheckChatInvite(ctx context.Context, request *tg.TLMessagesCheckChatInvite) (*tg.ChatInvite, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("messages.checkChatInvite - metadata: {}, request: {%v}", request)

	r, err := c.MessagesCheckChatInvite(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("messages.checkChatInvite - reply: {%v}", r)
	return r, err
}

// MessagesImportChatInvite
// messages.importChatInvite#6c50051c hash:string = Updates;
func (s *Service) MessagesImportChatInvite(ctx context.Context, request *tg.TLMessagesImportChatInvite) (*tg.Updates, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("messages.importChatInvite - metadata: {}, request: {%v}", request)

	r, err := c.MessagesImportChatInvite(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("messages.importChatInvite - reply: {%v}", r)
	return r, err
}

// MessagesGetExportedChatInvites
// messages.getExportedChatInvites#a2b5a3f6 flags:# revoked:flags.3?true peer:InputPeer admin_id:InputUser offset_date:flags.2?int offset_link:flags.2?string limit:int = messages.ExportedChatInvites;
func (s *Service) MessagesGetExportedChatInvites(ctx context.Context, request *tg.TLMessagesGetExportedChatInvites) (*tg.MessagesExportedChatInvites, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("messages.getExportedChatInvites - metadata: {}, request: {%v}", request)

	r, err := c.MessagesGetExportedChatInvites(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("messages.getExportedChatInvites - reply: {%v}", r)
	return r, err
}

// MessagesGetExportedChatInvite
// messages.getExportedChatInvite#73746f5c peer:InputPeer link:string = messages.ExportedChatInvite;
func (s *Service) MessagesGetExportedChatInvite(ctx context.Context, request *tg.TLMessagesGetExportedChatInvite) (*tg.MessagesExportedChatInvite, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("messages.getExportedChatInvite - metadata: {}, request: {%v}", request)

	r, err := c.MessagesGetExportedChatInvite(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("messages.getExportedChatInvite - reply: {%v}", r)
	return r, err
}

// MessagesEditExportedChatInvite
// messages.editExportedChatInvite#bdca2f75 flags:# revoked:flags.2?true peer:InputPeer link:string expire_date:flags.0?int usage_limit:flags.1?int request_needed:flags.3?Bool title:flags.4?string = messages.ExportedChatInvite;
func (s *Service) MessagesEditExportedChatInvite(ctx context.Context, request *tg.TLMessagesEditExportedChatInvite) (*tg.MessagesExportedChatInvite, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("messages.editExportedChatInvite - metadata: {}, request: {%v}", request)

	r, err := c.MessagesEditExportedChatInvite(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("messages.editExportedChatInvite - reply: {%v}", r)
	return r, err
}

// MessagesDeleteRevokedExportedChatInvites
// messages.deleteRevokedExportedChatInvites#56987bd5 peer:InputPeer admin_id:InputUser = Bool;
func (s *Service) MessagesDeleteRevokedExportedChatInvites(ctx context.Context, request *tg.TLMessagesDeleteRevokedExportedChatInvites) (*tg.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("messages.deleteRevokedExportedChatInvites - metadata: {}, request: {%v}", request)

	r, err := c.MessagesDeleteRevokedExportedChatInvites(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("messages.deleteRevokedExportedChatInvites - reply: {%v}", r)
	return r, err
}

// MessagesDeleteExportedChatInvite
// messages.deleteExportedChatInvite#d464a42b peer:InputPeer link:string = Bool;
func (s *Service) MessagesDeleteExportedChatInvite(ctx context.Context, request *tg.TLMessagesDeleteExportedChatInvite) (*tg.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("messages.deleteExportedChatInvite - metadata: {}, request: {%v}", request)

	r, err := c.MessagesDeleteExportedChatInvite(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("messages.deleteExportedChatInvite - reply: {%v}", r)
	return r, err
}

// MessagesGetAdminsWithInvites
// messages.getAdminsWithInvites#3920e6ef peer:InputPeer = messages.ChatAdminsWithInvites;
func (s *Service) MessagesGetAdminsWithInvites(ctx context.Context, request *tg.TLMessagesGetAdminsWithInvites) (*tg.MessagesChatAdminsWithInvites, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("messages.getAdminsWithInvites - metadata: {}, request: {%v}", request)

	r, err := c.MessagesGetAdminsWithInvites(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("messages.getAdminsWithInvites - reply: {%v}", r)
	return r, err
}

// MessagesGetChatInviteImporters
// messages.getChatInviteImporters#df04dd4e flags:# requested:flags.0?true subscription_expired:flags.3?true peer:InputPeer link:flags.1?string q:flags.2?string offset_date:int offset_user:InputUser limit:int = messages.ChatInviteImporters;
func (s *Service) MessagesGetChatInviteImporters(ctx context.Context, request *tg.TLMessagesGetChatInviteImporters) (*tg.MessagesChatInviteImporters, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("messages.getChatInviteImporters - metadata: {}, request: {%v}", request)

	r, err := c.MessagesGetChatInviteImporters(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("messages.getChatInviteImporters - reply: {%v}", r)
	return r, err
}

// MessagesHideChatJoinRequest
// messages.hideChatJoinRequest#7fe7e815 flags:# approved:flags.0?true peer:InputPeer user_id:InputUser = Updates;
func (s *Service) MessagesHideChatJoinRequest(ctx context.Context, request *tg.TLMessagesHideChatJoinRequest) (*tg.Updates, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("messages.hideChatJoinRequest - metadata: {}, request: {%v}", request)

	r, err := c.MessagesHideChatJoinRequest(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("messages.hideChatJoinRequest - reply: {%v}", r)
	return r, err
}

// MessagesHideAllChatJoinRequests
// messages.hideAllChatJoinRequests#e085f4ea flags:# approved:flags.0?true peer:InputPeer link:flags.1?string = Updates;
func (s *Service) MessagesHideAllChatJoinRequests(ctx context.Context, request *tg.TLMessagesHideAllChatJoinRequests) (*tg.Updates, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("messages.hideAllChatJoinRequests - metadata: {}, request: {%v}", request)

	r, err := c.MessagesHideAllChatJoinRequests(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("messages.hideAllChatJoinRequests - reply: {%v}", r)
	return r, err
}

// ChannelsToggleJoinToSend
// channels.toggleJoinToSend#e4cb9580 channel:InputChannel enabled:Bool = Updates;
func (s *Service) ChannelsToggleJoinToSend(ctx context.Context, request *tg.TLChannelsToggleJoinToSend) (*tg.Updates, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("channels.toggleJoinToSend - metadata: {}, request: {%v}", request)

	r, err := c.ChannelsToggleJoinToSend(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("channels.toggleJoinToSend - reply: {%v}", r)
	return r, err
}

// ChannelsToggleJoinRequest
// channels.toggleJoinRequest#4c2985b6 channel:InputChannel enabled:Bool = Updates;
func (s *Service) ChannelsToggleJoinRequest(ctx context.Context, request *tg.TLChannelsToggleJoinRequest) (*tg.Updates, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("channels.toggleJoinRequest - metadata: {}, request: {%v}", request)

	r, err := c.ChannelsToggleJoinRequest(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("channels.toggleJoinRequest - reply: {%v}", r)
	return r, err
}
