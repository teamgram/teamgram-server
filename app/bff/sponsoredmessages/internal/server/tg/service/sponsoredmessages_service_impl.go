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

// ContactsGetSponsoredPeers
// contacts.getSponsoredPeers#b6c8c393 q:string = contacts.SponsoredPeers;
func (s *Service) ContactsGetSponsoredPeers(ctx context.Context, request *tg.TLContactsGetSponsoredPeers) (*tg.ContactsSponsoredPeers, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("contacts.getSponsoredPeers - metadata: {}, request: {%v}", request)

	r, err := c.ContactsGetSponsoredPeers(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("contacts.getSponsoredPeers - reply: {%v}", r)
	return r, err
}

// MessagesViewSponsoredMessage
// messages.viewSponsoredMessage#269e3643 random_id:bytes = Bool;
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
// messages.clickSponsoredMessage#8235057e flags:# media:flags.0?true fullscreen:flags.1?true random_id:bytes = Bool;
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
// messages.reportSponsoredMessage#12cbf0c4 random_id:bytes option:bytes = channels.SponsoredMessageReportResult;
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
