/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright 2024 Teamgram Authors.
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package dialogsclient

import (
	"context"

	"github.com/teamgram/proto/mtproto"

	"github.com/zeromicro/go-zero/zrpc"
)

var _ *mtproto.Bool

type DialogsClient interface {
	MessagesGetDialogs(ctx context.Context, in *mtproto.TLMessagesGetDialogs) (*mtproto.Messages_Dialogs, error)
	MessagesSetTyping(ctx context.Context, in *mtproto.TLMessagesSetTyping) (*mtproto.Bool, error)
	MessagesGetPeerSettings(ctx context.Context, in *mtproto.TLMessagesGetPeerSettings) (*mtproto.Messages_PeerSettings, error)
	MessagesGetPeerDialogs(ctx context.Context, in *mtproto.TLMessagesGetPeerDialogs) (*mtproto.Messages_PeerDialogs, error)
	MessagesToggleDialogPin(ctx context.Context, in *mtproto.TLMessagesToggleDialogPin) (*mtproto.Bool, error)
	MessagesReorderPinnedDialogs(ctx context.Context, in *mtproto.TLMessagesReorderPinnedDialogs) (*mtproto.Bool, error)
	MessagesGetPinnedDialogs(ctx context.Context, in *mtproto.TLMessagesGetPinnedDialogs) (*mtproto.Messages_PeerDialogs, error)
	MessagesSendScreenshotNotification(ctx context.Context, in *mtproto.TLMessagesSendScreenshotNotification) (*mtproto.Updates, error)
	MessagesMarkDialogUnread(ctx context.Context, in *mtproto.TLMessagesMarkDialogUnread) (*mtproto.Bool, error)
	MessagesGetDialogUnreadMarks(ctx context.Context, in *mtproto.TLMessagesGetDialogUnreadMarks) (*mtproto.Vector_DialogPeer, error)
	MessagesGetOnlines(ctx context.Context, in *mtproto.TLMessagesGetOnlines) (*mtproto.ChatOnlines, error)
	MessagesHidePeerSettingsBar(ctx context.Context, in *mtproto.TLMessagesHidePeerSettingsBar) (*mtproto.Bool, error)
	MessagesSetHistoryTTL(ctx context.Context, in *mtproto.TLMessagesSetHistoryTTL) (*mtproto.Updates, error)
}

type defaultDialogsClient struct {
	cli zrpc.Client
}

func NewDialogsClient(cli zrpc.Client) DialogsClient {
	return &defaultDialogsClient{
		cli: cli,
	}
}

// MessagesGetDialogs
// messages.getDialogs#a0f4cb4f flags:# exclude_pinned:flags.0?true folder_id:flags.1?int offset_date:int offset_id:int offset_peer:InputPeer limit:int hash:long = messages.Dialogs;
func (m *defaultDialogsClient) MessagesGetDialogs(ctx context.Context, in *mtproto.TLMessagesGetDialogs) (*mtproto.Messages_Dialogs, error) {
	client := mtproto.NewRPCDialogsClient(m.cli.Conn())
	return client.MessagesGetDialogs(ctx, in)
}

// MessagesSetTyping
// messages.setTyping#58943ee2 flags:# peer:InputPeer top_msg_id:flags.0?int action:SendMessageAction = Bool;
func (m *defaultDialogsClient) MessagesSetTyping(ctx context.Context, in *mtproto.TLMessagesSetTyping) (*mtproto.Bool, error) {
	client := mtproto.NewRPCDialogsClient(m.cli.Conn())
	return client.MessagesSetTyping(ctx, in)
}

// MessagesGetPeerSettings
// messages.getPeerSettings#efd9a6a2 peer:InputPeer = messages.PeerSettings;
func (m *defaultDialogsClient) MessagesGetPeerSettings(ctx context.Context, in *mtproto.TLMessagesGetPeerSettings) (*mtproto.Messages_PeerSettings, error) {
	client := mtproto.NewRPCDialogsClient(m.cli.Conn())
	return client.MessagesGetPeerSettings(ctx, in)
}

// MessagesGetPeerDialogs
// messages.getPeerDialogs#e470bcfd peers:Vector<InputDialogPeer> = messages.PeerDialogs;
func (m *defaultDialogsClient) MessagesGetPeerDialogs(ctx context.Context, in *mtproto.TLMessagesGetPeerDialogs) (*mtproto.Messages_PeerDialogs, error) {
	client := mtproto.NewRPCDialogsClient(m.cli.Conn())
	return client.MessagesGetPeerDialogs(ctx, in)
}

// MessagesToggleDialogPin
// messages.toggleDialogPin#a731e257 flags:# pinned:flags.0?true peer:InputDialogPeer = Bool;
func (m *defaultDialogsClient) MessagesToggleDialogPin(ctx context.Context, in *mtproto.TLMessagesToggleDialogPin) (*mtproto.Bool, error) {
	client := mtproto.NewRPCDialogsClient(m.cli.Conn())
	return client.MessagesToggleDialogPin(ctx, in)
}

// MessagesReorderPinnedDialogs
// messages.reorderPinnedDialogs#3b1adf37 flags:# force:flags.0?true folder_id:int order:Vector<InputDialogPeer> = Bool;
func (m *defaultDialogsClient) MessagesReorderPinnedDialogs(ctx context.Context, in *mtproto.TLMessagesReorderPinnedDialogs) (*mtproto.Bool, error) {
	client := mtproto.NewRPCDialogsClient(m.cli.Conn())
	return client.MessagesReorderPinnedDialogs(ctx, in)
}

// MessagesGetPinnedDialogs
// messages.getPinnedDialogs#d6b94df2 folder_id:int = messages.PeerDialogs;
func (m *defaultDialogsClient) MessagesGetPinnedDialogs(ctx context.Context, in *mtproto.TLMessagesGetPinnedDialogs) (*mtproto.Messages_PeerDialogs, error) {
	client := mtproto.NewRPCDialogsClient(m.cli.Conn())
	return client.MessagesGetPinnedDialogs(ctx, in)
}

// MessagesSendScreenshotNotification
// messages.sendScreenshotNotification#a1405817 peer:InputPeer reply_to:InputReplyTo random_id:long = Updates;
func (m *defaultDialogsClient) MessagesSendScreenshotNotification(ctx context.Context, in *mtproto.TLMessagesSendScreenshotNotification) (*mtproto.Updates, error) {
	client := mtproto.NewRPCDialogsClient(m.cli.Conn())
	return client.MessagesSendScreenshotNotification(ctx, in)
}

// MessagesMarkDialogUnread
// messages.markDialogUnread#c286d98f flags:# unread:flags.0?true peer:InputDialogPeer = Bool;
func (m *defaultDialogsClient) MessagesMarkDialogUnread(ctx context.Context, in *mtproto.TLMessagesMarkDialogUnread) (*mtproto.Bool, error) {
	client := mtproto.NewRPCDialogsClient(m.cli.Conn())
	return client.MessagesMarkDialogUnread(ctx, in)
}

// MessagesGetDialogUnreadMarks
// messages.getDialogUnreadMarks#22e24e22 = Vector<DialogPeer>;
func (m *defaultDialogsClient) MessagesGetDialogUnreadMarks(ctx context.Context, in *mtproto.TLMessagesGetDialogUnreadMarks) (*mtproto.Vector_DialogPeer, error) {
	client := mtproto.NewRPCDialogsClient(m.cli.Conn())
	return client.MessagesGetDialogUnreadMarks(ctx, in)
}

// MessagesGetOnlines
// messages.getOnlines#6e2be050 peer:InputPeer = ChatOnlines;
func (m *defaultDialogsClient) MessagesGetOnlines(ctx context.Context, in *mtproto.TLMessagesGetOnlines) (*mtproto.ChatOnlines, error) {
	client := mtproto.NewRPCDialogsClient(m.cli.Conn())
	return client.MessagesGetOnlines(ctx, in)
}

// MessagesHidePeerSettingsBar
// messages.hidePeerSettingsBar#4facb138 peer:InputPeer = Bool;
func (m *defaultDialogsClient) MessagesHidePeerSettingsBar(ctx context.Context, in *mtproto.TLMessagesHidePeerSettingsBar) (*mtproto.Bool, error) {
	client := mtproto.NewRPCDialogsClient(m.cli.Conn())
	return client.MessagesHidePeerSettingsBar(ctx, in)
}

// MessagesSetHistoryTTL
// messages.setHistoryTTL#b80e5fe4 peer:InputPeer period:int = Updates;
func (m *defaultDialogsClient) MessagesSetHistoryTTL(ctx context.Context, in *mtproto.TLMessagesSetHistoryTTL) (*mtproto.Updates, error) {
	client := mtproto.NewRPCDialogsClient(m.cli.Conn())
	return client.MessagesSetHistoryTTL(ctx, in)
}
