/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright 2024 Teamgooo Authors.
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package savedmessagedialogsclient

import (
	"context"

	"github.com/teamgooo/teamgooo-server/app/bff/savedmessagedialogs/savedmessagedialogs/savedmessagedialogsservice"
	"github.com/teamgooo/teamgooo-server/pkg/proto/tg"

	"github.com/cloudwego/kitex/client"
)

type SavedMessageDialogsClient interface {
	MessagesGetSavedDialogs(ctx context.Context, in *tg.TLMessagesGetSavedDialogs) (*tg.MessagesSavedDialogs, error)
	MessagesGetSavedHistory(ctx context.Context, in *tg.TLMessagesGetSavedHistory) (*tg.MessagesMessages, error)
	MessagesDeleteSavedHistory(ctx context.Context, in *tg.TLMessagesDeleteSavedHistory) (*tg.MessagesAffectedHistory, error)
	MessagesGetPinnedSavedDialogs(ctx context.Context, in *tg.TLMessagesGetPinnedSavedDialogs) (*tg.MessagesSavedDialogs, error)
	MessagesToggleSavedDialogPin(ctx context.Context, in *tg.TLMessagesToggleSavedDialogPin) (*tg.Bool, error)
	MessagesReorderPinnedSavedDialogs(ctx context.Context, in *tg.TLMessagesReorderPinnedSavedDialogs) (*tg.Bool, error)
	MessagesGetSavedDialogsByID(ctx context.Context, in *tg.TLMessagesGetSavedDialogsByID) (*tg.MessagesSavedDialogs, error)
	MessagesReadSavedHistory(ctx context.Context, in *tg.TLMessagesReadSavedHistory) (*tg.Bool, error)
	ChannelsGetMessageAuthor(ctx context.Context, in *tg.TLChannelsGetMessageAuthor) (*tg.User, error)
}

type defaultSavedMessageDialogsClient struct {
	cli client.Client
}

func NewSavedMessageDialogsClient(cli client.Client) SavedMessageDialogsClient {
	return &defaultSavedMessageDialogsClient{
		cli: cli,
	}
}

// MessagesGetSavedDialogs
// messages.getSavedDialogs#1e91fc99 flags:# exclude_pinned:flags.0?true parent_peer:flags.1?InputPeer offset_date:int offset_id:int offset_peer:InputPeer limit:int hash:long = messages.SavedDialogs;
func (m *defaultSavedMessageDialogsClient) MessagesGetSavedDialogs(ctx context.Context, in *tg.TLMessagesGetSavedDialogs) (*tg.MessagesSavedDialogs, error) {
	cli := savedmessagedialogsservice.NewRPCSavedMessageDialogsClient(m.cli)
	return cli.MessagesGetSavedDialogs(ctx, in)
}

// MessagesGetSavedHistory
// messages.getSavedHistory#998ab009 flags:# parent_peer:flags.0?InputPeer peer:InputPeer offset_id:int offset_date:int add_offset:int limit:int max_id:int min_id:int hash:long = messages.Messages;
func (m *defaultSavedMessageDialogsClient) MessagesGetSavedHistory(ctx context.Context, in *tg.TLMessagesGetSavedHistory) (*tg.MessagesMessages, error) {
	cli := savedmessagedialogsservice.NewRPCSavedMessageDialogsClient(m.cli)
	return cli.MessagesGetSavedHistory(ctx, in)
}

// MessagesDeleteSavedHistory
// messages.deleteSavedHistory#4dc5085f flags:# parent_peer:flags.0?InputPeer peer:InputPeer max_id:int min_date:flags.2?int max_date:flags.3?int = messages.AffectedHistory;
func (m *defaultSavedMessageDialogsClient) MessagesDeleteSavedHistory(ctx context.Context, in *tg.TLMessagesDeleteSavedHistory) (*tg.MessagesAffectedHistory, error) {
	cli := savedmessagedialogsservice.NewRPCSavedMessageDialogsClient(m.cli)
	return cli.MessagesDeleteSavedHistory(ctx, in)
}

// MessagesGetPinnedSavedDialogs
// messages.getPinnedSavedDialogs#d63d94e0 = messages.SavedDialogs;
func (m *defaultSavedMessageDialogsClient) MessagesGetPinnedSavedDialogs(ctx context.Context, in *tg.TLMessagesGetPinnedSavedDialogs) (*tg.MessagesSavedDialogs, error) {
	cli := savedmessagedialogsservice.NewRPCSavedMessageDialogsClient(m.cli)
	return cli.MessagesGetPinnedSavedDialogs(ctx, in)
}

// MessagesToggleSavedDialogPin
// messages.toggleSavedDialogPin#ac81bbde flags:# pinned:flags.0?true peer:InputDialogPeer = Bool;
func (m *defaultSavedMessageDialogsClient) MessagesToggleSavedDialogPin(ctx context.Context, in *tg.TLMessagesToggleSavedDialogPin) (*tg.Bool, error) {
	cli := savedmessagedialogsservice.NewRPCSavedMessageDialogsClient(m.cli)
	return cli.MessagesToggleSavedDialogPin(ctx, in)
}

// MessagesReorderPinnedSavedDialogs
// messages.reorderPinnedSavedDialogs#8b716587 flags:# force:flags.0?true order:Vector<InputDialogPeer> = Bool;
func (m *defaultSavedMessageDialogsClient) MessagesReorderPinnedSavedDialogs(ctx context.Context, in *tg.TLMessagesReorderPinnedSavedDialogs) (*tg.Bool, error) {
	cli := savedmessagedialogsservice.NewRPCSavedMessageDialogsClient(m.cli)
	return cli.MessagesReorderPinnedSavedDialogs(ctx, in)
}

// MessagesGetSavedDialogsByID
// messages.getSavedDialogsByID#6f6f9c96 flags:# parent_peer:flags.1?InputPeer ids:Vector<InputPeer> = messages.SavedDialogs;
func (m *defaultSavedMessageDialogsClient) MessagesGetSavedDialogsByID(ctx context.Context, in *tg.TLMessagesGetSavedDialogsByID) (*tg.MessagesSavedDialogs, error) {
	cli := savedmessagedialogsservice.NewRPCSavedMessageDialogsClient(m.cli)
	return cli.MessagesGetSavedDialogsByID(ctx, in)
}

// MessagesReadSavedHistory
// messages.readSavedHistory#ba4a3b5b parent_peer:InputPeer peer:InputPeer max_id:int = Bool;
func (m *defaultSavedMessageDialogsClient) MessagesReadSavedHistory(ctx context.Context, in *tg.TLMessagesReadSavedHistory) (*tg.Bool, error) {
	cli := savedmessagedialogsservice.NewRPCSavedMessageDialogsClient(m.cli)
	return cli.MessagesReadSavedHistory(ctx, in)
}

// ChannelsGetMessageAuthor
// channels.getMessageAuthor#ece2a0e6 channel:InputChannel id:int = User;
func (m *defaultSavedMessageDialogsClient) ChannelsGetMessageAuthor(ctx context.Context, in *tg.TLChannelsGetMessageAuthor) (*tg.User, error) {
	cli := savedmessagedialogsservice.NewRPCSavedMessageDialogsClient(m.cli)
	return cli.ChannelsGetMessageAuthor(ctx, in)
}
