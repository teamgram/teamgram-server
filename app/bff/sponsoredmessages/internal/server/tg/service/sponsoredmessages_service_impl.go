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

	"github.com/teamgram/proto/v2/tg"
	"github.com/teamgram/teamgram-server/v2/app/bff/sponsoredmessages/internal/core"
)

// AccountToggleSponsoredMessages
// account.toggleSponsoredMessages#b9d9a38d enabled:Bool = Bool;
func (s *Service) AccountToggleSponsoredMessages(ctx context.Context, request *tg.TLAccountToggleSponsoredMessages) (*tg.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("account.toggleSponsoredMessages - metadata: {}, request: {%v}", request)

	r, err := c.AccountToggleSponsoredMessages(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("account.toggleSponsoredMessages - reply: {%v}", r)
	return r, err
}

// MessagesViewSponsoredMessage
// messages.viewSponsoredMessage#673ad8f1 peer:InputPeer random_id:bytes = Bool;
func (s *Service) MessagesViewSponsoredMessage(ctx context.Context, request *tg.TLMessagesViewSponsoredMessage) (*tg.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("messages.viewSponsoredMessage - metadata: {}, request: {%v}", request)

	r, err := c.MessagesViewSponsoredMessage(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("messages.viewSponsoredMessage - reply: {%v}", r)
	return r, err
}

// MessagesClickSponsoredMessage
// messages.clickSponsoredMessage#f093465 flags:# media:flags.0?true fullscreen:flags.1?true peer:InputPeer random_id:bytes = Bool;
func (s *Service) MessagesClickSponsoredMessage(ctx context.Context, request *tg.TLMessagesClickSponsoredMessage) (*tg.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("messages.clickSponsoredMessage - metadata: {}, request: {%v}", request)

	r, err := c.MessagesClickSponsoredMessage(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("messages.clickSponsoredMessage - reply: {%v}", r)
	return r, err
}

// MessagesReportSponsoredMessage
// messages.reportSponsoredMessage#1af3dbb8 peer:InputPeer random_id:bytes option:bytes = channels.SponsoredMessageReportResult;
func (s *Service) MessagesReportSponsoredMessage(ctx context.Context, request *tg.TLMessagesReportSponsoredMessage) (*tg.ChannelsSponsoredMessageReportResult, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("messages.reportSponsoredMessage - metadata: {}, request: {%v}", request)

	r, err := c.MessagesReportSponsoredMessage(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("messages.reportSponsoredMessage - reply: {%v}", r)
	return r, err
}

// MessagesGetSponsoredMessages
// messages.getSponsoredMessages#9bd2f439 peer:InputPeer = messages.SponsoredMessages;
func (s *Service) MessagesGetSponsoredMessages(ctx context.Context, request *tg.TLMessagesGetSponsoredMessages) (*tg.MessagesSponsoredMessages, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("messages.getSponsoredMessages - metadata: {}, request: {%v}", request)

	r, err := c.MessagesGetSponsoredMessages(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("messages.getSponsoredMessages - reply: {%v}", r)
	return r, err
}

// ChannelsRestrictSponsoredMessages
// channels.restrictSponsoredMessages#9ae91519 channel:InputChannel restricted:Bool = Updates;
func (s *Service) ChannelsRestrictSponsoredMessages(ctx context.Context, request *tg.TLChannelsRestrictSponsoredMessages) (*tg.Updates, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("channels.restrictSponsoredMessages - metadata: {}, request: {%v}", request)

	r, err := c.ChannelsRestrictSponsoredMessages(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("channels.restrictSponsoredMessages - reply: {%v}", r)
	return r, err
}

// ChannelsViewSponsoredMessage
// channels.viewSponsoredMessage#beaedb94 channel:InputChannel random_id:bytes = Bool;
func (s *Service) ChannelsViewSponsoredMessage(ctx context.Context, request *tg.TLChannelsViewSponsoredMessage) (*tg.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("channels.viewSponsoredMessage - metadata: {}, request: {%v}", request)

	r, err := c.ChannelsViewSponsoredMessage(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("channels.viewSponsoredMessage - reply: {%v}", r)
	return r, err
}

// ChannelsGetSponsoredMessages
// channels.getSponsoredMessages#ec210fbf channel:InputChannel = messages.SponsoredMessages;
func (s *Service) ChannelsGetSponsoredMessages(ctx context.Context, request *tg.TLChannelsGetSponsoredMessages) (*tg.MessagesSponsoredMessages, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("channels.getSponsoredMessages - metadata: {}, request: {%v}", request)

	r, err := c.ChannelsGetSponsoredMessages(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("channels.getSponsoredMessages - reply: {%v}", r)
	return r, err
}

// ChannelsClickSponsoredMessage
// channels.clickSponsoredMessage#1445d75 flags:# media:flags.0?true fullscreen:flags.1?true channel:InputChannel random_id:bytes = Bool;
func (s *Service) ChannelsClickSponsoredMessage(ctx context.Context, request *tg.TLChannelsClickSponsoredMessage) (*tg.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("channels.clickSponsoredMessage - metadata: {}, request: {%v}", request)

	r, err := c.ChannelsClickSponsoredMessage(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("channels.clickSponsoredMessage - reply: {%v}", r)
	return r, err
}

// ChannelsReportSponsoredMessage
// channels.reportSponsoredMessage#af8ff6b9 channel:InputChannel random_id:bytes option:bytes = channels.SponsoredMessageReportResult;
func (s *Service) ChannelsReportSponsoredMessage(ctx context.Context, request *tg.TLChannelsReportSponsoredMessage) (*tg.ChannelsSponsoredMessageReportResult, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("channels.reportSponsoredMessage - metadata: {}, request: {%v}", request)

	r, err := c.ChannelsReportSponsoredMessage(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("channels.reportSponsoredMessage - reply: {%v}", r)
	return r, err
}
