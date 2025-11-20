/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright 2025 Teamgram Authors.
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package service

import (
	"context"

	"github.com/teamgram/proto/mtproto"
	"github.com/teamgram/teamgram-server/app/bff/savedmessagedialogs/internal/core"
)

// MessagesGetSavedDialogs
// messages.getSavedDialogs#1e91fc99 flags:# exclude_pinned:flags.0?true parent_peer:flags.1?InputPeer offset_date:int offset_id:int offset_peer:InputPeer limit:int hash:long = messages.SavedDialogs;
func (s *Service) MessagesGetSavedDialogs(ctx context.Context, request *mtproto.TLMessagesGetSavedDialogs) (*mtproto.Messages_SavedDialogs, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("messages.getSavedDialogs - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.MessagesGetSavedDialogs(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("messages.getSavedDialogs - reply: {%s}", r)
	return r, err
}

// MessagesGetSavedHistory
// messages.getSavedHistory#998ab009 flags:# parent_peer:flags.0?InputPeer peer:InputPeer offset_id:int offset_date:int add_offset:int limit:int max_id:int min_id:int hash:long = messages.Messages;
func (s *Service) MessagesGetSavedHistory(ctx context.Context, request *mtproto.TLMessagesGetSavedHistory) (*mtproto.Messages_Messages, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("messages.getSavedHistory - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.MessagesGetSavedHistory(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("messages.getSavedHistory - reply: {%s}", r)
	return r, err
}

// MessagesDeleteSavedHistory
// messages.deleteSavedHistory#4dc5085f flags:# parent_peer:flags.0?InputPeer peer:InputPeer max_id:int min_date:flags.2?int max_date:flags.3?int = messages.AffectedHistory;
func (s *Service) MessagesDeleteSavedHistory(ctx context.Context, request *mtproto.TLMessagesDeleteSavedHistory) (*mtproto.Messages_AffectedHistory, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("messages.deleteSavedHistory - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.MessagesDeleteSavedHistory(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("messages.deleteSavedHistory - reply: {%s}", r)
	return r, err
}

// MessagesGetPinnedSavedDialogs
// messages.getPinnedSavedDialogs#d63d94e0 = messages.SavedDialogs;
func (s *Service) MessagesGetPinnedSavedDialogs(ctx context.Context, request *mtproto.TLMessagesGetPinnedSavedDialogs) (*mtproto.Messages_SavedDialogs, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("messages.getPinnedSavedDialogs - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.MessagesGetPinnedSavedDialogs(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("messages.getPinnedSavedDialogs - reply: {%s}", r)
	return r, err
}

// MessagesToggleSavedDialogPin
// messages.toggleSavedDialogPin#ac81bbde flags:# pinned:flags.0?true peer:InputDialogPeer = Bool;
func (s *Service) MessagesToggleSavedDialogPin(ctx context.Context, request *mtproto.TLMessagesToggleSavedDialogPin) (*mtproto.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("messages.toggleSavedDialogPin - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.MessagesToggleSavedDialogPin(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("messages.toggleSavedDialogPin - reply: {%s}", r)
	return r, err
}

// MessagesReorderPinnedSavedDialogs
// messages.reorderPinnedSavedDialogs#8b716587 flags:# force:flags.0?true order:Vector<InputDialogPeer> = Bool;
func (s *Service) MessagesReorderPinnedSavedDialogs(ctx context.Context, request *mtproto.TLMessagesReorderPinnedSavedDialogs) (*mtproto.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("messages.reorderPinnedSavedDialogs - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.MessagesReorderPinnedSavedDialogs(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("messages.reorderPinnedSavedDialogs - reply: {%s}", r)
	return r, err
}

// MessagesGetSavedDialogsByID
// messages.getSavedDialogsByID#6f6f9c96 flags:# parent_peer:flags.1?InputPeer ids:Vector<InputPeer> = messages.SavedDialogs;
func (s *Service) MessagesGetSavedDialogsByID(ctx context.Context, request *mtproto.TLMessagesGetSavedDialogsByID) (*mtproto.Messages_SavedDialogs, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("messages.getSavedDialogsByID - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.MessagesGetSavedDialogsByID(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("messages.getSavedDialogsByID - reply: {%s}", r)
	return r, err
}

// MessagesReadSavedHistory
// messages.readSavedHistory#ba4a3b5b parent_peer:InputPeer peer:InputPeer max_id:int = Bool;
func (s *Service) MessagesReadSavedHistory(ctx context.Context, request *mtproto.TLMessagesReadSavedHistory) (*mtproto.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("messages.readSavedHistory - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.MessagesReadSavedHistory(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("messages.readSavedHistory - reply: {%s}", r)
	return r, err
}

// ChannelsGetMessageAuthor
// channels.getMessageAuthor#ece2a0e6 channel:InputChannel id:int = User;
func (s *Service) ChannelsGetMessageAuthor(ctx context.Context, request *mtproto.TLChannelsGetMessageAuthor) (*mtproto.User, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("channels.getMessageAuthor - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.ChannelsGetMessageAuthor(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("channels.getMessageAuthor - reply: {%s}", r)
	return r, err
}
