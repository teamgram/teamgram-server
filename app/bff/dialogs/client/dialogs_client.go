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

	"github.com/teamgram/proto/v2/tg"
	"github.com/teamgram/teamgram-server/v2/app/bff/dialogs/dialogs/dialogsservice"

	"github.com/cloudwego/kitex/client"
)

type DialogsClient interface {
	MessagesGetDialogs(ctx context.Context, in *tg.TLMessagesGetDialogs) (*tg.MessagesDialogs, error)
	MessagesSetTyping(ctx context.Context, in *tg.TLMessagesSetTyping) (*tg.Bool, error)
	MessagesGetPeerSettings(ctx context.Context, in *tg.TLMessagesGetPeerSettings) (*tg.MessagesPeerSettings, error)
	MessagesGetPeerDialogs(ctx context.Context, in *tg.TLMessagesGetPeerDialogs) (*tg.MessagesPeerDialogs, error)
	MessagesToggleDialogPin(ctx context.Context, in *tg.TLMessagesToggleDialogPin) (*tg.Bool, error)
	MessagesReorderPinnedDialogs(ctx context.Context, in *tg.TLMessagesReorderPinnedDialogs) (*tg.Bool, error)
	MessagesGetPinnedDialogs(ctx context.Context, in *tg.TLMessagesGetPinnedDialogs) (*tg.MessagesPeerDialogs, error)
	MessagesSendScreenshotNotification(ctx context.Context, in *tg.TLMessagesSendScreenshotNotification) (*tg.Updates, error)
	MessagesMarkDialogUnread(ctx context.Context, in *tg.TLMessagesMarkDialogUnread) (*tg.Bool, error)
	MessagesGetDialogUnreadMarks(ctx context.Context, in *tg.TLMessagesGetDialogUnreadMarks) (*tg.VectorDialogPeer, error)
	MessagesGetOnlines(ctx context.Context, in *tg.TLMessagesGetOnlines) (*tg.ChatOnlines, error)
	MessagesHidePeerSettingsBar(ctx context.Context, in *tg.TLMessagesHidePeerSettingsBar) (*tg.Bool, error)
	MessagesSetHistoryTTL(ctx context.Context, in *tg.TLMessagesSetHistoryTTL) (*tg.Updates, error)
}

type defaultDialogsClient struct {
	cli client.Client
}

func NewDialogsClient(cli client.Client) DialogsClient {
	return &defaultDialogsClient{
		cli: cli,
	}
}

// MessagesGetDialogs
// messages.getDialogs#a0f4cb4f flags:# exclude_pinned:flags.0?true folder_id:flags.1?int offset_date:int offset_id:int offset_peer:InputPeer limit:int hash:long = messages.Dialogs;
func (m *defaultDialogsClient) MessagesGetDialogs(ctx context.Context, in *tg.TLMessagesGetDialogs) (*tg.MessagesDialogs, error) {
	cli := dialogsservice.NewRPCDialogsClient(m.cli)
	return cli.MessagesGetDialogs(ctx, in)
}

// MessagesSetTyping
// messages.setTyping#58943ee2 flags:# peer:InputPeer top_msg_id:flags.0?int action:SendMessageAction = Bool;
func (m *defaultDialogsClient) MessagesSetTyping(ctx context.Context, in *tg.TLMessagesSetTyping) (*tg.Bool, error) {
	cli := dialogsservice.NewRPCDialogsClient(m.cli)
	return cli.MessagesSetTyping(ctx, in)
}

// MessagesGetPeerSettings
// messages.getPeerSettings#efd9a6a2 peer:InputPeer = messages.PeerSettings;
func (m *defaultDialogsClient) MessagesGetPeerSettings(ctx context.Context, in *tg.TLMessagesGetPeerSettings) (*tg.MessagesPeerSettings, error) {
	cli := dialogsservice.NewRPCDialogsClient(m.cli)
	return cli.MessagesGetPeerSettings(ctx, in)
}

// MessagesGetPeerDialogs
// messages.getPeerDialogs#e470bcfd peers:Vector<InputDialogPeer> = messages.PeerDialogs;
func (m *defaultDialogsClient) MessagesGetPeerDialogs(ctx context.Context, in *tg.TLMessagesGetPeerDialogs) (*tg.MessagesPeerDialogs, error) {
	cli := dialogsservice.NewRPCDialogsClient(m.cli)
	return cli.MessagesGetPeerDialogs(ctx, in)
}

// MessagesToggleDialogPin
// messages.toggleDialogPin#a731e257 flags:# pinned:flags.0?true peer:InputDialogPeer = Bool;
func (m *defaultDialogsClient) MessagesToggleDialogPin(ctx context.Context, in *tg.TLMessagesToggleDialogPin) (*tg.Bool, error) {
	cli := dialogsservice.NewRPCDialogsClient(m.cli)
	return cli.MessagesToggleDialogPin(ctx, in)
}

// MessagesReorderPinnedDialogs
// messages.reorderPinnedDialogs#3b1adf37 flags:# force:flags.0?true folder_id:int order:Vector<InputDialogPeer> = Bool;
func (m *defaultDialogsClient) MessagesReorderPinnedDialogs(ctx context.Context, in *tg.TLMessagesReorderPinnedDialogs) (*tg.Bool, error) {
	cli := dialogsservice.NewRPCDialogsClient(m.cli)
	return cli.MessagesReorderPinnedDialogs(ctx, in)
}

// MessagesGetPinnedDialogs
// messages.getPinnedDialogs#d6b94df2 folder_id:int = messages.PeerDialogs;
func (m *defaultDialogsClient) MessagesGetPinnedDialogs(ctx context.Context, in *tg.TLMessagesGetPinnedDialogs) (*tg.MessagesPeerDialogs, error) {
	cli := dialogsservice.NewRPCDialogsClient(m.cli)
	return cli.MessagesGetPinnedDialogs(ctx, in)
}

// MessagesSendScreenshotNotification
// messages.sendScreenshotNotification#a1405817 peer:InputPeer reply_to:InputReplyTo random_id:long = Updates;
func (m *defaultDialogsClient) MessagesSendScreenshotNotification(ctx context.Context, in *tg.TLMessagesSendScreenshotNotification) (*tg.Updates, error) {
	cli := dialogsservice.NewRPCDialogsClient(m.cli)
	return cli.MessagesSendScreenshotNotification(ctx, in)
}

// MessagesMarkDialogUnread
// messages.markDialogUnread#c286d98f flags:# unread:flags.0?true peer:InputDialogPeer = Bool;
func (m *defaultDialogsClient) MessagesMarkDialogUnread(ctx context.Context, in *tg.TLMessagesMarkDialogUnread) (*tg.Bool, error) {
	cli := dialogsservice.NewRPCDialogsClient(m.cli)
	return cli.MessagesMarkDialogUnread(ctx, in)
}

// MessagesGetDialogUnreadMarks
// messages.getDialogUnreadMarks#22e24e22 = Vector<DialogPeer>;
func (m *defaultDialogsClient) MessagesGetDialogUnreadMarks(ctx context.Context, in *tg.TLMessagesGetDialogUnreadMarks) (*tg.VectorDialogPeer, error) {
	cli := dialogsservice.NewRPCDialogsClient(m.cli)
	return cli.MessagesGetDialogUnreadMarks(ctx, in)
}

// MessagesGetOnlines
// messages.getOnlines#6e2be050 peer:InputPeer = ChatOnlines;
func (m *defaultDialogsClient) MessagesGetOnlines(ctx context.Context, in *tg.TLMessagesGetOnlines) (*tg.ChatOnlines, error) {
	cli := dialogsservice.NewRPCDialogsClient(m.cli)
	return cli.MessagesGetOnlines(ctx, in)
}

// MessagesHidePeerSettingsBar
// messages.hidePeerSettingsBar#4facb138 peer:InputPeer = Bool;
func (m *defaultDialogsClient) MessagesHidePeerSettingsBar(ctx context.Context, in *tg.TLMessagesHidePeerSettingsBar) (*tg.Bool, error) {
	cli := dialogsservice.NewRPCDialogsClient(m.cli)
	return cli.MessagesHidePeerSettingsBar(ctx, in)
}

// MessagesSetHistoryTTL
// messages.setHistoryTTL#b80e5fe4 peer:InputPeer period:int = Updates;
func (m *defaultDialogsClient) MessagesSetHistoryTTL(ctx context.Context, in *tg.TLMessagesSetHistoryTTL) (*tg.Updates, error) {
	cli := dialogsservice.NewRPCDialogsClient(m.cli)
	return cli.MessagesSetHistoryTTL(ctx, in)
}
