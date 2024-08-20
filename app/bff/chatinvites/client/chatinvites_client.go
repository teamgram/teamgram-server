/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright 2024 Teamgram Authors.
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package chatinvitesclient

import (
	"context"

	"github.com/teamgram/proto/mtproto"

	"github.com/zeromicro/go-zero/zrpc"
)

var _ *mtproto.Bool

type ChatInvitesClient interface {
	MessagesExportChatInvite(ctx context.Context, in *mtproto.TLMessagesExportChatInvite) (*mtproto.ExportedChatInvite, error)
	MessagesCheckChatInvite(ctx context.Context, in *mtproto.TLMessagesCheckChatInvite) (*mtproto.ChatInvite, error)
	MessagesImportChatInvite(ctx context.Context, in *mtproto.TLMessagesImportChatInvite) (*mtproto.Updates, error)
	MessagesGetExportedChatInvites(ctx context.Context, in *mtproto.TLMessagesGetExportedChatInvites) (*mtproto.Messages_ExportedChatInvites, error)
	MessagesGetExportedChatInvite(ctx context.Context, in *mtproto.TLMessagesGetExportedChatInvite) (*mtproto.Messages_ExportedChatInvite, error)
	MessagesEditExportedChatInvite(ctx context.Context, in *mtproto.TLMessagesEditExportedChatInvite) (*mtproto.Messages_ExportedChatInvite, error)
	MessagesDeleteRevokedExportedChatInvites(ctx context.Context, in *mtproto.TLMessagesDeleteRevokedExportedChatInvites) (*mtproto.Bool, error)
	MessagesDeleteExportedChatInvite(ctx context.Context, in *mtproto.TLMessagesDeleteExportedChatInvite) (*mtproto.Bool, error)
	MessagesGetAdminsWithInvites(ctx context.Context, in *mtproto.TLMessagesGetAdminsWithInvites) (*mtproto.Messages_ChatAdminsWithInvites, error)
	MessagesGetChatInviteImporters(ctx context.Context, in *mtproto.TLMessagesGetChatInviteImporters) (*mtproto.Messages_ChatInviteImporters, error)
	MessagesHideChatJoinRequest(ctx context.Context, in *mtproto.TLMessagesHideChatJoinRequest) (*mtproto.Updates, error)
	MessagesHideAllChatJoinRequests(ctx context.Context, in *mtproto.TLMessagesHideAllChatJoinRequests) (*mtproto.Updates, error)
	ChannelsToggleJoinToSend(ctx context.Context, in *mtproto.TLChannelsToggleJoinToSend) (*mtproto.Updates, error)
	ChannelsToggleJoinRequest(ctx context.Context, in *mtproto.TLChannelsToggleJoinRequest) (*mtproto.Updates, error)
}

type defaultChatInvitesClient struct {
	cli zrpc.Client
}

func NewChatInvitesClient(cli zrpc.Client) ChatInvitesClient {
	return &defaultChatInvitesClient{
		cli: cli,
	}
}

// MessagesExportChatInvite
// messages.exportChatInvite#a455de90 flags:# legacy_revoke_permanent:flags.2?true request_needed:flags.3?true peer:InputPeer expire_date:flags.0?int usage_limit:flags.1?int title:flags.4?string subscription_pricing:flags.5?StarsSubscriptionPricing = ExportedChatInvite;
func (m *defaultChatInvitesClient) MessagesExportChatInvite(ctx context.Context, in *mtproto.TLMessagesExportChatInvite) (*mtproto.ExportedChatInvite, error) {
	client := mtproto.NewRPCChatInvitesClient(m.cli.Conn())
	return client.MessagesExportChatInvite(ctx, in)
}

// MessagesCheckChatInvite
// messages.checkChatInvite#3eadb1bb hash:string = ChatInvite;
func (m *defaultChatInvitesClient) MessagesCheckChatInvite(ctx context.Context, in *mtproto.TLMessagesCheckChatInvite) (*mtproto.ChatInvite, error) {
	client := mtproto.NewRPCChatInvitesClient(m.cli.Conn())
	return client.MessagesCheckChatInvite(ctx, in)
}

// MessagesImportChatInvite
// messages.importChatInvite#6c50051c hash:string = Updates;
func (m *defaultChatInvitesClient) MessagesImportChatInvite(ctx context.Context, in *mtproto.TLMessagesImportChatInvite) (*mtproto.Updates, error) {
	client := mtproto.NewRPCChatInvitesClient(m.cli.Conn())
	return client.MessagesImportChatInvite(ctx, in)
}

// MessagesGetExportedChatInvites
// messages.getExportedChatInvites#a2b5a3f6 flags:# revoked:flags.3?true peer:InputPeer admin_id:InputUser offset_date:flags.2?int offset_link:flags.2?string limit:int = messages.ExportedChatInvites;
func (m *defaultChatInvitesClient) MessagesGetExportedChatInvites(ctx context.Context, in *mtproto.TLMessagesGetExportedChatInvites) (*mtproto.Messages_ExportedChatInvites, error) {
	client := mtproto.NewRPCChatInvitesClient(m.cli.Conn())
	return client.MessagesGetExportedChatInvites(ctx, in)
}

// MessagesGetExportedChatInvite
// messages.getExportedChatInvite#73746f5c peer:InputPeer link:string = messages.ExportedChatInvite;
func (m *defaultChatInvitesClient) MessagesGetExportedChatInvite(ctx context.Context, in *mtproto.TLMessagesGetExportedChatInvite) (*mtproto.Messages_ExportedChatInvite, error) {
	client := mtproto.NewRPCChatInvitesClient(m.cli.Conn())
	return client.MessagesGetExportedChatInvite(ctx, in)
}

// MessagesEditExportedChatInvite
// messages.editExportedChatInvite#bdca2f75 flags:# revoked:flags.2?true peer:InputPeer link:string expire_date:flags.0?int usage_limit:flags.1?int request_needed:flags.3?Bool title:flags.4?string = messages.ExportedChatInvite;
func (m *defaultChatInvitesClient) MessagesEditExportedChatInvite(ctx context.Context, in *mtproto.TLMessagesEditExportedChatInvite) (*mtproto.Messages_ExportedChatInvite, error) {
	client := mtproto.NewRPCChatInvitesClient(m.cli.Conn())
	return client.MessagesEditExportedChatInvite(ctx, in)
}

// MessagesDeleteRevokedExportedChatInvites
// messages.deleteRevokedExportedChatInvites#56987bd5 peer:InputPeer admin_id:InputUser = Bool;
func (m *defaultChatInvitesClient) MessagesDeleteRevokedExportedChatInvites(ctx context.Context, in *mtproto.TLMessagesDeleteRevokedExportedChatInvites) (*mtproto.Bool, error) {
	client := mtproto.NewRPCChatInvitesClient(m.cli.Conn())
	return client.MessagesDeleteRevokedExportedChatInvites(ctx, in)
}

// MessagesDeleteExportedChatInvite
// messages.deleteExportedChatInvite#d464a42b peer:InputPeer link:string = Bool;
func (m *defaultChatInvitesClient) MessagesDeleteExportedChatInvite(ctx context.Context, in *mtproto.TLMessagesDeleteExportedChatInvite) (*mtproto.Bool, error) {
	client := mtproto.NewRPCChatInvitesClient(m.cli.Conn())
	return client.MessagesDeleteExportedChatInvite(ctx, in)
}

// MessagesGetAdminsWithInvites
// messages.getAdminsWithInvites#3920e6ef peer:InputPeer = messages.ChatAdminsWithInvites;
func (m *defaultChatInvitesClient) MessagesGetAdminsWithInvites(ctx context.Context, in *mtproto.TLMessagesGetAdminsWithInvites) (*mtproto.Messages_ChatAdminsWithInvites, error) {
	client := mtproto.NewRPCChatInvitesClient(m.cli.Conn())
	return client.MessagesGetAdminsWithInvites(ctx, in)
}

// MessagesGetChatInviteImporters
// messages.getChatInviteImporters#df04dd4e flags:# requested:flags.0?true subscription_expired:flags.3?true peer:InputPeer link:flags.1?string q:flags.2?string offset_date:int offset_user:InputUser limit:int = messages.ChatInviteImporters;
func (m *defaultChatInvitesClient) MessagesGetChatInviteImporters(ctx context.Context, in *mtproto.TLMessagesGetChatInviteImporters) (*mtproto.Messages_ChatInviteImporters, error) {
	client := mtproto.NewRPCChatInvitesClient(m.cli.Conn())
	return client.MessagesGetChatInviteImporters(ctx, in)
}

// MessagesHideChatJoinRequest
// messages.hideChatJoinRequest#7fe7e815 flags:# approved:flags.0?true peer:InputPeer user_id:InputUser = Updates;
func (m *defaultChatInvitesClient) MessagesHideChatJoinRequest(ctx context.Context, in *mtproto.TLMessagesHideChatJoinRequest) (*mtproto.Updates, error) {
	client := mtproto.NewRPCChatInvitesClient(m.cli.Conn())
	return client.MessagesHideChatJoinRequest(ctx, in)
}

// MessagesHideAllChatJoinRequests
// messages.hideAllChatJoinRequests#e085f4ea flags:# approved:flags.0?true peer:InputPeer link:flags.1?string = Updates;
func (m *defaultChatInvitesClient) MessagesHideAllChatJoinRequests(ctx context.Context, in *mtproto.TLMessagesHideAllChatJoinRequests) (*mtproto.Updates, error) {
	client := mtproto.NewRPCChatInvitesClient(m.cli.Conn())
	return client.MessagesHideAllChatJoinRequests(ctx, in)
}

// ChannelsToggleJoinToSend
// channels.toggleJoinToSend#e4cb9580 channel:InputChannel enabled:Bool = Updates;
func (m *defaultChatInvitesClient) ChannelsToggleJoinToSend(ctx context.Context, in *mtproto.TLChannelsToggleJoinToSend) (*mtproto.Updates, error) {
	client := mtproto.NewRPCChatInvitesClient(m.cli.Conn())
	return client.ChannelsToggleJoinToSend(ctx, in)
}

// ChannelsToggleJoinRequest
// channels.toggleJoinRequest#4c2985b6 channel:InputChannel enabled:Bool = Updates;
func (m *defaultChatInvitesClient) ChannelsToggleJoinRequest(ctx context.Context, in *mtproto.TLChannelsToggleJoinRequest) (*mtproto.Updates, error) {
	client := mtproto.NewRPCChatInvitesClient(m.cli.Conn())
	return client.ChannelsToggleJoinRequest(ctx, in)
}
