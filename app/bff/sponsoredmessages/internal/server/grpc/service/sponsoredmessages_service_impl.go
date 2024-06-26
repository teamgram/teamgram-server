/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright 2024 Teamgram Authors.
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

// AccountToggleSponsoredMessages
// account.toggleSponsoredMessages#b9d9a38d enabled:Bool = Bool;
func (s *Service) AccountToggleSponsoredMessages(ctx context.Context, request *mtproto.TLAccountToggleSponsoredMessages) (*mtproto.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("account.toggleSponsoredMessages - metadata: %s, request: %s", c.MD, request)

	r, err := c.AccountToggleSponsoredMessages(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("account.toggleSponsoredMessages - reply: %s", r)
	return r, err
}

// ChannelsViewSponsoredMessage
// channels.viewSponsoredMessage#beaedb94 channel:InputChannel random_id:bytes = Bool;
func (s *Service) ChannelsViewSponsoredMessage(ctx context.Context, request *mtproto.TLChannelsViewSponsoredMessage) (*mtproto.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("channels.viewSponsoredMessage - metadata: %s, request: %s", c.MD, request)

	r, err := c.ChannelsViewSponsoredMessage(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("channels.viewSponsoredMessage - reply: %s", r)
	return r, err
}

// ChannelsGetSponsoredMessages
// channels.getSponsoredMessages#ec210fbf channel:InputChannel = messages.SponsoredMessages;
func (s *Service) ChannelsGetSponsoredMessages(ctx context.Context, request *mtproto.TLChannelsGetSponsoredMessages) (*mtproto.Messages_SponsoredMessages, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("channels.getSponsoredMessages - metadata: %s, request: %s", c.MD, request)

	r, err := c.ChannelsGetSponsoredMessages(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("channels.getSponsoredMessages - reply: %s", r)
	return r, err
}

// ChannelsClickSponsoredMessage
// channels.clickSponsoredMessage#18afbc93 channel:InputChannel random_id:bytes = Bool;
func (s *Service) ChannelsClickSponsoredMessage(ctx context.Context, request *mtproto.TLChannelsClickSponsoredMessage) (*mtproto.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("channels.clickSponsoredMessage - metadata: %s, request: %s", c.MD, request)

	r, err := c.ChannelsClickSponsoredMessage(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("channels.clickSponsoredMessage - reply: %s", r)
	return r, err
}

// ChannelsReportSponsoredMessage
// channels.reportSponsoredMessage#af8ff6b9 channel:InputChannel random_id:bytes option:bytes = channels.SponsoredMessageReportResult;
func (s *Service) ChannelsReportSponsoredMessage(ctx context.Context, request *mtproto.TLChannelsReportSponsoredMessage) (*mtproto.Channels_SponsoredMessageReportResult, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("channels.reportSponsoredMessage - metadata: %s, request: %s", c.MD, request)

	r, err := c.ChannelsReportSponsoredMessage(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("channels.reportSponsoredMessage - reply: %s", r)
	return r, err
}

// ChannelsRestrictSponsoredMessages
// channels.restrictSponsoredMessages#9ae91519 channel:InputChannel restricted:Bool = Updates;
func (s *Service) ChannelsRestrictSponsoredMessages(ctx context.Context, request *mtproto.TLChannelsRestrictSponsoredMessages) (*mtproto.Updates, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("channels.restrictSponsoredMessages - metadata: %s, request: %s", c.MD, request)

	r, err := c.ChannelsRestrictSponsoredMessages(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("channels.restrictSponsoredMessages - reply: %s", r)
	return r, err
}
