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
	"github.com/teamgram/teamgram-server/app/bff/folders/internal/core"
)

// MessagesGetDialogFilters
// messages.getDialogFilters#f19ed96d = Vector<DialogFilter>;
func (s *Service) MessagesGetDialogFilters(ctx context.Context, request *mtproto.TLMessagesGetDialogFilters) (*mtproto.Vector_DialogFilter, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("messages.getDialogFilters - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.MessagesGetDialogFilters(request)
	if err != nil {
		return nil, err
	}

	c.Infof("messages.getDialogFilters - reply: %s", r.DebugString())
	return r, err
}

// MessagesGetSuggestedDialogFilters
// messages.getSuggestedDialogFilters#a29cd42c = Vector<DialogFilterSuggested>;
func (s *Service) MessagesGetSuggestedDialogFilters(ctx context.Context, request *mtproto.TLMessagesGetSuggestedDialogFilters) (*mtproto.Vector_DialogFilterSuggested, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("messages.getSuggestedDialogFilters - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.MessagesGetSuggestedDialogFilters(request)
	if err != nil {
		return nil, err
	}

	c.Infof("messages.getSuggestedDialogFilters - reply: %s", r.DebugString())
	return r, err
}

// MessagesUpdateDialogFilter
// messages.updateDialogFilter#1ad4a04a flags:# id:int filter:flags.0?DialogFilter = Bool;
func (s *Service) MessagesUpdateDialogFilter(ctx context.Context, request *mtproto.TLMessagesUpdateDialogFilter) (*mtproto.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("messages.updateDialogFilter - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.MessagesUpdateDialogFilter(request)
	if err != nil {
		return nil, err
	}

	c.Infof("messages.updateDialogFilter - reply: %s", r.DebugString())
	return r, err
}

// MessagesUpdateDialogFiltersOrder
// messages.updateDialogFiltersOrder#c563c1e4 order:Vector<int> = Bool;
func (s *Service) MessagesUpdateDialogFiltersOrder(ctx context.Context, request *mtproto.TLMessagesUpdateDialogFiltersOrder) (*mtproto.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("messages.updateDialogFiltersOrder - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.MessagesUpdateDialogFiltersOrder(request)
	if err != nil {
		return nil, err
	}

	c.Infof("messages.updateDialogFiltersOrder - reply: %s", r.DebugString())
	return r, err
}

// FoldersEditPeerFolders
// folders.editPeerFolders#6847d0ab folder_peers:Vector<InputFolderPeer> = Updates;
func (s *Service) FoldersEditPeerFolders(ctx context.Context, request *mtproto.TLFoldersEditPeerFolders) (*mtproto.Updates, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("folders.editPeerFolders - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.FoldersEditPeerFolders(request)
	if err != nil {
		return nil, err
	}

	c.Infof("folders.editPeerFolders - reply: %s", r.DebugString())
	return r, err
}

// FoldersDeleteFolder
// folders.deleteFolder#1c295881 folder_id:int = Updates;
func (s *Service) FoldersDeleteFolder(ctx context.Context, request *mtproto.TLFoldersDeleteFolder) (*mtproto.Updates, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("folders.deleteFolder - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.FoldersDeleteFolder(request)
	if err != nil {
		return nil, err
	}

	c.Infof("folders.deleteFolder - reply: %s", r.DebugString())
	return r, err
}
