/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright 2024 Teamgram Authors.
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package savedmessagedialogsclient

import (
	"context"

	"github.com/teamgram/proto/mtproto"

	"github.com/zeromicro/go-zero/zrpc"
)

var _ *mtproto.Bool

type SavedMessageDialogsClient interface {
	MessagesGetSavedDialogs(ctx context.Context, in *mtproto.TLMessagesGetSavedDialogs) (*mtproto.Messages_SavedDialogs, error)
	MessagesGetSavedHistory(ctx context.Context, in *mtproto.TLMessagesGetSavedHistory) (*mtproto.Messages_Messages, error)
	MessagesDeleteSavedHistory(ctx context.Context, in *mtproto.TLMessagesDeleteSavedHistory) (*mtproto.Messages_AffectedHistory, error)
	MessagesGetPinnedSavedDialogs(ctx context.Context, in *mtproto.TLMessagesGetPinnedSavedDialogs) (*mtproto.Messages_SavedDialogs, error)
	MessagesToggleSavedDialogPin(ctx context.Context, in *mtproto.TLMessagesToggleSavedDialogPin) (*mtproto.Bool, error)
	MessagesReorderPinnedSavedDialogs(ctx context.Context, in *mtproto.TLMessagesReorderPinnedSavedDialogs) (*mtproto.Bool, error)
}

type defaultSavedMessageDialogsClient struct {
	cli zrpc.Client
}

func NewSavedMessageDialogsClient(cli zrpc.Client) SavedMessageDialogsClient {
	return &defaultSavedMessageDialogsClient{
		cli: cli,
	}
}

// MessagesGetSavedDialogs
// messages.getSavedDialogs#5381d21a flags:# exclude_pinned:flags.0?true offset_date:int offset_id:int offset_peer:InputPeer limit:int hash:long = messages.SavedDialogs;
func (m *defaultSavedMessageDialogsClient) MessagesGetSavedDialogs(ctx context.Context, in *mtproto.TLMessagesGetSavedDialogs) (*mtproto.Messages_SavedDialogs, error) {
	client := mtproto.NewRPCSavedMessageDialogsClient(m.cli.Conn())
	return client.MessagesGetSavedDialogs(ctx, in)
}

// MessagesGetSavedHistory
// messages.getSavedHistory#3d9a414d peer:InputPeer offset_id:int offset_date:int add_offset:int limit:int max_id:int min_id:int hash:long = messages.Messages;
func (m *defaultSavedMessageDialogsClient) MessagesGetSavedHistory(ctx context.Context, in *mtproto.TLMessagesGetSavedHistory) (*mtproto.Messages_Messages, error) {
	client := mtproto.NewRPCSavedMessageDialogsClient(m.cli.Conn())
	return client.MessagesGetSavedHistory(ctx, in)
}

// MessagesDeleteSavedHistory
// messages.deleteSavedHistory#6e98102b flags:# peer:InputPeer max_id:int min_date:flags.2?int max_date:flags.3?int = messages.AffectedHistory;
func (m *defaultSavedMessageDialogsClient) MessagesDeleteSavedHistory(ctx context.Context, in *mtproto.TLMessagesDeleteSavedHistory) (*mtproto.Messages_AffectedHistory, error) {
	client := mtproto.NewRPCSavedMessageDialogsClient(m.cli.Conn())
	return client.MessagesDeleteSavedHistory(ctx, in)
}

// MessagesGetPinnedSavedDialogs
// messages.getPinnedSavedDialogs#d63d94e0 = messages.SavedDialogs;
func (m *defaultSavedMessageDialogsClient) MessagesGetPinnedSavedDialogs(ctx context.Context, in *mtproto.TLMessagesGetPinnedSavedDialogs) (*mtproto.Messages_SavedDialogs, error) {
	client := mtproto.NewRPCSavedMessageDialogsClient(m.cli.Conn())
	return client.MessagesGetPinnedSavedDialogs(ctx, in)
}

// MessagesToggleSavedDialogPin
// messages.toggleSavedDialogPin#ac81bbde flags:# pinned:flags.0?true peer:InputDialogPeer = Bool;
func (m *defaultSavedMessageDialogsClient) MessagesToggleSavedDialogPin(ctx context.Context, in *mtproto.TLMessagesToggleSavedDialogPin) (*mtproto.Bool, error) {
	client := mtproto.NewRPCSavedMessageDialogsClient(m.cli.Conn())
	return client.MessagesToggleSavedDialogPin(ctx, in)
}

// MessagesReorderPinnedSavedDialogs
// messages.reorderPinnedSavedDialogs#8b716587 flags:# force:flags.0?true order:Vector<InputDialogPeer> = Bool;
func (m *defaultSavedMessageDialogsClient) MessagesReorderPinnedSavedDialogs(ctx context.Context, in *mtproto.TLMessagesReorderPinnedSavedDialogs) (*mtproto.Bool, error) {
	client := mtproto.NewRPCSavedMessageDialogsClient(m.cli.Conn())
	return client.MessagesReorderPinnedSavedDialogs(ctx, in)
}
