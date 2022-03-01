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
	"github.com/teamgram/teamgram-server/app/bff/reactions/internal/core"
)

// MessagesSendReaction
// messages.sendReaction#25690ce4 flags:# big:flags.1?true peer:InputPeer msg_id:int reaction:flags.0?string = Updates;
func (s *Service) MessagesSendReaction(ctx context.Context, request *mtproto.TLMessagesSendReaction) (*mtproto.Updates, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("messages.sendReaction - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.MessagesSendReaction(request)
	if err != nil {
		return nil, err
	}

	c.Infof("messages.sendReaction - reply: %s", r.DebugString())
	return r, err
}

// MessagesGetMessagesReactions
// messages.getMessagesReactions#8bba90e6 peer:InputPeer id:Vector<int> = Updates;
func (s *Service) MessagesGetMessagesReactions(ctx context.Context, request *mtproto.TLMessagesGetMessagesReactions) (*mtproto.Updates, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("messages.getMessagesReactions - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.MessagesGetMessagesReactions(request)
	if err != nil {
		return nil, err
	}

	c.Infof("messages.getMessagesReactions - reply: %s", r.DebugString())
	return r, err
}

// MessagesGetMessageReactionsList
// messages.getMessageReactionsList#e0ee6b77 flags:# peer:InputPeer id:int reaction:flags.0?string offset:flags.1?string limit:int = messages.MessageReactionsList;
func (s *Service) MessagesGetMessageReactionsList(ctx context.Context, request *mtproto.TLMessagesGetMessageReactionsList) (*mtproto.Messages_MessageReactionsList, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("messages.getMessageReactionsList - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.MessagesGetMessageReactionsList(request)
	if err != nil {
		return nil, err
	}

	c.Infof("messages.getMessageReactionsList - reply: %s", r.DebugString())
	return r, err
}

// MessagesSetChatAvailableReactions
// messages.setChatAvailableReactions#14050ea6 peer:InputPeer available_reactions:Vector<string> = Updates;
func (s *Service) MessagesSetChatAvailableReactions(ctx context.Context, request *mtproto.TLMessagesSetChatAvailableReactions) (*mtproto.Updates, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("messages.setChatAvailableReactions - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.MessagesSetChatAvailableReactions(request)
	if err != nil {
		return nil, err
	}

	c.Infof("messages.setChatAvailableReactions - reply: %s", r.DebugString())
	return r, err
}

// MessagesGetAvailableReactions
// messages.getAvailableReactions#18dea0ac hash:int = messages.AvailableReactions;
func (s *Service) MessagesGetAvailableReactions(ctx context.Context, request *mtproto.TLMessagesGetAvailableReactions) (*mtproto.Messages_AvailableReactions, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("messages.getAvailableReactions - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.MessagesGetAvailableReactions(request)
	if err != nil {
		return nil, err
	}

	c.Infof("messages.getAvailableReactions - reply: %s", r.DebugString())
	return r, err
}

// MessagesSetDefaultReaction
// messages.setDefaultReaction#d960c4d4 reaction:string = Bool;
func (s *Service) MessagesSetDefaultReaction(ctx context.Context, request *mtproto.TLMessagesSetDefaultReaction) (*mtproto.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("messages.setDefaultReaction - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.MessagesSetDefaultReaction(request)
	if err != nil {
		return nil, err
	}

	c.Infof("messages.setDefaultReaction - reply: %s", r.DebugString())
	return r, err
}

// MessagesGetUnreadReactions
// messages.getUnreadReactions#e85bae1a peer:InputPeer offset_id:int add_offset:int limit:int max_id:int min_id:int = messages.Messages;
func (s *Service) MessagesGetUnreadReactions(ctx context.Context, request *mtproto.TLMessagesGetUnreadReactions) (*mtproto.Messages_Messages, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("messages.getUnreadReactions - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.MessagesGetUnreadReactions(request)
	if err != nil {
		return nil, err
	}

	c.Infof("messages.getUnreadReactions - reply: %s", r.DebugString())
	return r, err
}

// MessagesReadReactions
// messages.readReactions#82e251d7 peer:InputPeer = messages.AffectedHistory;
func (s *Service) MessagesReadReactions(ctx context.Context, request *mtproto.TLMessagesReadReactions) (*mtproto.Messages_AffectedHistory, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("messages.readReactions - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.MessagesReadReactions(request)
	if err != nil {
		return nil, err
	}

	c.Infof("messages.readReactions - reply: %s", r.DebugString())
	return r, err
}
