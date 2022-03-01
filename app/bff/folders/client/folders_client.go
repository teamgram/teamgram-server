/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright 2022 Teamgram Authors.
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package folders_client

import (
	"context"

	"github.com/teamgram/proto/mtproto"

	"github.com/zeromicro/go-zero/zrpc"
)

var _ *mtproto.Bool

type FoldersClient interface {
	MessagesGetDialogFilters(ctx context.Context, in *mtproto.TLMessagesGetDialogFilters) (*mtproto.Vector_DialogFilter, error)
	MessagesGetSuggestedDialogFilters(ctx context.Context, in *mtproto.TLMessagesGetSuggestedDialogFilters) (*mtproto.Vector_DialogFilterSuggested, error)
	MessagesUpdateDialogFilter(ctx context.Context, in *mtproto.TLMessagesUpdateDialogFilter) (*mtproto.Bool, error)
	MessagesUpdateDialogFiltersOrder(ctx context.Context, in *mtproto.TLMessagesUpdateDialogFiltersOrder) (*mtproto.Bool, error)
	FoldersEditPeerFolders(ctx context.Context, in *mtproto.TLFoldersEditPeerFolders) (*mtproto.Updates, error)
	FoldersDeleteFolder(ctx context.Context, in *mtproto.TLFoldersDeleteFolder) (*mtproto.Updates, error)
}

type defaultFoldersClient struct {
	cli zrpc.Client
}

func NewFoldersClient(cli zrpc.Client) FoldersClient {
	return &defaultFoldersClient{
		cli: cli,
	}
}

// MessagesGetDialogFilters
// messages.getDialogFilters#f19ed96d = Vector<DialogFilter>;
func (m *defaultFoldersClient) MessagesGetDialogFilters(ctx context.Context, in *mtproto.TLMessagesGetDialogFilters) (*mtproto.Vector_DialogFilter, error) {
	client := mtproto.NewRPCFoldersClient(m.cli.Conn())
	return client.MessagesGetDialogFilters(ctx, in)
}

// MessagesGetSuggestedDialogFilters
// messages.getSuggestedDialogFilters#a29cd42c = Vector<DialogFilterSuggested>;
func (m *defaultFoldersClient) MessagesGetSuggestedDialogFilters(ctx context.Context, in *mtproto.TLMessagesGetSuggestedDialogFilters) (*mtproto.Vector_DialogFilterSuggested, error) {
	client := mtproto.NewRPCFoldersClient(m.cli.Conn())
	return client.MessagesGetSuggestedDialogFilters(ctx, in)
}

// MessagesUpdateDialogFilter
// messages.updateDialogFilter#1ad4a04a flags:# id:int filter:flags.0?DialogFilter = Bool;
func (m *defaultFoldersClient) MessagesUpdateDialogFilter(ctx context.Context, in *mtproto.TLMessagesUpdateDialogFilter) (*mtproto.Bool, error) {
	client := mtproto.NewRPCFoldersClient(m.cli.Conn())
	return client.MessagesUpdateDialogFilter(ctx, in)
}

// MessagesUpdateDialogFiltersOrder
// messages.updateDialogFiltersOrder#c563c1e4 order:Vector<int> = Bool;
func (m *defaultFoldersClient) MessagesUpdateDialogFiltersOrder(ctx context.Context, in *mtproto.TLMessagesUpdateDialogFiltersOrder) (*mtproto.Bool, error) {
	client := mtproto.NewRPCFoldersClient(m.cli.Conn())
	return client.MessagesUpdateDialogFiltersOrder(ctx, in)
}

// FoldersEditPeerFolders
// folders.editPeerFolders#6847d0ab folder_peers:Vector<InputFolderPeer> = Updates;
func (m *defaultFoldersClient) FoldersEditPeerFolders(ctx context.Context, in *mtproto.TLFoldersEditPeerFolders) (*mtproto.Updates, error) {
	client := mtproto.NewRPCFoldersClient(m.cli.Conn())
	return client.FoldersEditPeerFolders(ctx, in)
}

// FoldersDeleteFolder
// folders.deleteFolder#1c295881 folder_id:int = Updates;
func (m *defaultFoldersClient) FoldersDeleteFolder(ctx context.Context, in *mtproto.TLFoldersDeleteFolder) (*mtproto.Updates, error) {
	client := mtproto.NewRPCFoldersClient(m.cli.Conn())
	return client.FoldersDeleteFolder(ctx, in)
}
