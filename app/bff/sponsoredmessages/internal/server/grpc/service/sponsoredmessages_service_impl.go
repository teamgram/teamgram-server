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
	"github.com/teamgram/teamgram-server/app/bff/sponsoredmessages/internal/core"
)

// ChannelsViewSponsoredMessage
// channels.viewSponsoredMessage#beaedb94 channel:InputChannel random_id:bytes = Bool;
func (s *Service) ChannelsViewSponsoredMessage(ctx context.Context, request *mtproto.TLChannelsViewSponsoredMessage) (*mtproto.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("channels.viewSponsoredMessage - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.ChannelsViewSponsoredMessage(request)
	if err != nil {
		return nil, err
	}

	c.Infof("channels.viewSponsoredMessage - reply: %s", r.DebugString())
	return r, err
}

// ChannelsGetSponsoredMessages
// channels.getSponsoredMessages#ec210fbf channel:InputChannel = messages.SponsoredMessages;
func (s *Service) ChannelsGetSponsoredMessages(ctx context.Context, request *mtproto.TLChannelsGetSponsoredMessages) (*mtproto.Messages_SponsoredMessages, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("channels.getSponsoredMessages - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.ChannelsGetSponsoredMessages(request)
	if err != nil {
		return nil, err
	}

	c.Infof("channels.getSponsoredMessages - reply: %s", r.DebugString())
	return r, err
}
