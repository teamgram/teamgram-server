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
	"github.com/teamgram/teamgram-server/app/bff/gdpr/internal/core"
)

// AccountInitTakeoutSession
// account.initTakeoutSession#f05b4804 flags:# contacts:flags.0?true message_users:flags.1?true message_chats:flags.2?true message_megagroups:flags.3?true message_channels:flags.4?true files:flags.5?true file_max_size:flags.5?int = account.Takeout;
func (s *Service) AccountInitTakeoutSession(ctx context.Context, request *mtproto.TLAccountInitTakeoutSession) (*mtproto.Account_Takeout, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("account.initTakeoutSession - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.AccountInitTakeoutSession(request)
	if err != nil {
		return nil, err
	}

	c.Infof("account.initTakeoutSession - reply: %s", r.DebugString())
	return r, err
}

// AccountFinishTakeoutSession
// account.finishTakeoutSession#1d2652ee flags:# success:flags.0?true = Bool;
func (s *Service) AccountFinishTakeoutSession(ctx context.Context, request *mtproto.TLAccountFinishTakeoutSession) (*mtproto.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("account.finishTakeoutSession - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.AccountFinishTakeoutSession(request)
	if err != nil {
		return nil, err
	}

	c.Infof("account.finishTakeoutSession - reply: %s", r.DebugString())
	return r, err
}

// MessagesGetSplitRanges
// messages.getSplitRanges#1cff7e08 = Vector<MessageRange>;
func (s *Service) MessagesGetSplitRanges(ctx context.Context, request *mtproto.TLMessagesGetSplitRanges) (*mtproto.Vector_MessageRange, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("messages.getSplitRanges - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.MessagesGetSplitRanges(request)
	if err != nil {
		return nil, err
	}

	c.Infof("messages.getSplitRanges - reply: %s", r.DebugString())
	return r, err
}

// ChannelsGetLeftChannels
// channels.getLeftChannels#8341ecc0 offset:int = messages.Chats;
func (s *Service) ChannelsGetLeftChannels(ctx context.Context, request *mtproto.TLChannelsGetLeftChannels) (*mtproto.Messages_Chats, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("channels.getLeftChannels - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.ChannelsGetLeftChannels(request)
	if err != nil {
		return nil, err
	}

	c.Infof("channels.getLeftChannels - reply: %s", r.DebugString())
	return r, err
}
