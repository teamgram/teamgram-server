/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright 2024 Teamgram Authors.
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package sponsoredmessagesclient

import (
	"context"

	"github.com/teamgram/proto/v2/tg"
	"github.com/teamgram/teamgram-server/v2/app/bff/sponsoredmessages/sponsoredmessages/sponsoredmessagesservice"

	"github.com/cloudwego/kitex/client"
)

type SponsoredMessagesClient interface {
	AccountToggleSponsoredMessages(ctx context.Context, in *tg.TLAccountToggleSponsoredMessages) (*tg.Bool, error)
	MessagesViewSponsoredMessage(ctx context.Context, in *tg.TLMessagesViewSponsoredMessage) (*tg.Bool, error)
	MessagesClickSponsoredMessage(ctx context.Context, in *tg.TLMessagesClickSponsoredMessage) (*tg.Bool, error)
	MessagesReportSponsoredMessage(ctx context.Context, in *tg.TLMessagesReportSponsoredMessage) (*tg.ChannelsSponsoredMessageReportResult, error)
	MessagesGetSponsoredMessages(ctx context.Context, in *tg.TLMessagesGetSponsoredMessages) (*tg.MessagesSponsoredMessages, error)
	ChannelsRestrictSponsoredMessages(ctx context.Context, in *tg.TLChannelsRestrictSponsoredMessages) (*tg.Updates, error)
	ChannelsViewSponsoredMessage(ctx context.Context, in *tg.TLChannelsViewSponsoredMessage) (*tg.Bool, error)
	ChannelsGetSponsoredMessages(ctx context.Context, in *tg.TLChannelsGetSponsoredMessages) (*tg.MessagesSponsoredMessages, error)
	ChannelsClickSponsoredMessage(ctx context.Context, in *tg.TLChannelsClickSponsoredMessage) (*tg.Bool, error)
	ChannelsReportSponsoredMessage(ctx context.Context, in *tg.TLChannelsReportSponsoredMessage) (*tg.ChannelsSponsoredMessageReportResult, error)
}

type defaultSponsoredMessagesClient struct {
	cli client.Client
}

func NewSponsoredMessagesClient(cli client.Client) SponsoredMessagesClient {
	return &defaultSponsoredMessagesClient{
		cli: cli,
	}
}

// AccountToggleSponsoredMessages
// account.toggleSponsoredMessages#b9d9a38d enabled:Bool = Bool;
func (m *defaultSponsoredMessagesClient) AccountToggleSponsoredMessages(ctx context.Context, in *tg.TLAccountToggleSponsoredMessages) (*tg.Bool, error) {
	cli := sponsoredmessagesservice.NewRPCSponsoredMessagesClient(m.cli)
	return cli.AccountToggleSponsoredMessages(ctx, in)
}

// MessagesViewSponsoredMessage
// messages.viewSponsoredMessage#673ad8f1 peer:InputPeer random_id:bytes = Bool;
func (m *defaultSponsoredMessagesClient) MessagesViewSponsoredMessage(ctx context.Context, in *tg.TLMessagesViewSponsoredMessage) (*tg.Bool, error) {
	cli := sponsoredmessagesservice.NewRPCSponsoredMessagesClient(m.cli)
	return cli.MessagesViewSponsoredMessage(ctx, in)
}

// MessagesClickSponsoredMessage
// messages.clickSponsoredMessage#f093465 flags:# media:flags.0?true fullscreen:flags.1?true peer:InputPeer random_id:bytes = Bool;
func (m *defaultSponsoredMessagesClient) MessagesClickSponsoredMessage(ctx context.Context, in *tg.TLMessagesClickSponsoredMessage) (*tg.Bool, error) {
	cli := sponsoredmessagesservice.NewRPCSponsoredMessagesClient(m.cli)
	return cli.MessagesClickSponsoredMessage(ctx, in)
}

// MessagesReportSponsoredMessage
// messages.reportSponsoredMessage#1af3dbb8 peer:InputPeer random_id:bytes option:bytes = channels.SponsoredMessageReportResult;
func (m *defaultSponsoredMessagesClient) MessagesReportSponsoredMessage(ctx context.Context, in *tg.TLMessagesReportSponsoredMessage) (*tg.ChannelsSponsoredMessageReportResult, error) {
	cli := sponsoredmessagesservice.NewRPCSponsoredMessagesClient(m.cli)
	return cli.MessagesReportSponsoredMessage(ctx, in)
}

// MessagesGetSponsoredMessages
// messages.getSponsoredMessages#9bd2f439 peer:InputPeer = messages.SponsoredMessages;
func (m *defaultSponsoredMessagesClient) MessagesGetSponsoredMessages(ctx context.Context, in *tg.TLMessagesGetSponsoredMessages) (*tg.MessagesSponsoredMessages, error) {
	cli := sponsoredmessagesservice.NewRPCSponsoredMessagesClient(m.cli)
	return cli.MessagesGetSponsoredMessages(ctx, in)
}

// ChannelsRestrictSponsoredMessages
// channels.restrictSponsoredMessages#9ae91519 channel:InputChannel restricted:Bool = Updates;
func (m *defaultSponsoredMessagesClient) ChannelsRestrictSponsoredMessages(ctx context.Context, in *tg.TLChannelsRestrictSponsoredMessages) (*tg.Updates, error) {
	cli := sponsoredmessagesservice.NewRPCSponsoredMessagesClient(m.cli)
	return cli.ChannelsRestrictSponsoredMessages(ctx, in)
}

// ChannelsViewSponsoredMessage
// channels.viewSponsoredMessage#beaedb94 channel:InputChannel random_id:bytes = Bool;
func (m *defaultSponsoredMessagesClient) ChannelsViewSponsoredMessage(ctx context.Context, in *tg.TLChannelsViewSponsoredMessage) (*tg.Bool, error) {
	cli := sponsoredmessagesservice.NewRPCSponsoredMessagesClient(m.cli)
	return cli.ChannelsViewSponsoredMessage(ctx, in)
}

// ChannelsGetSponsoredMessages
// channels.getSponsoredMessages#ec210fbf channel:InputChannel = messages.SponsoredMessages;
func (m *defaultSponsoredMessagesClient) ChannelsGetSponsoredMessages(ctx context.Context, in *tg.TLChannelsGetSponsoredMessages) (*tg.MessagesSponsoredMessages, error) {
	cli := sponsoredmessagesservice.NewRPCSponsoredMessagesClient(m.cli)
	return cli.ChannelsGetSponsoredMessages(ctx, in)
}

// ChannelsClickSponsoredMessage
// channels.clickSponsoredMessage#1445d75 flags:# media:flags.0?true fullscreen:flags.1?true channel:InputChannel random_id:bytes = Bool;
func (m *defaultSponsoredMessagesClient) ChannelsClickSponsoredMessage(ctx context.Context, in *tg.TLChannelsClickSponsoredMessage) (*tg.Bool, error) {
	cli := sponsoredmessagesservice.NewRPCSponsoredMessagesClient(m.cli)
	return cli.ChannelsClickSponsoredMessage(ctx, in)
}

// ChannelsReportSponsoredMessage
// channels.reportSponsoredMessage#af8ff6b9 channel:InputChannel random_id:bytes option:bytes = channels.SponsoredMessageReportResult;
func (m *defaultSponsoredMessagesClient) ChannelsReportSponsoredMessage(ctx context.Context, in *tg.TLChannelsReportSponsoredMessage) (*tg.ChannelsSponsoredMessageReportResult, error) {
	cli := sponsoredmessagesservice.NewRPCSponsoredMessagesClient(m.cli)
	return cli.ChannelsReportSponsoredMessage(ctx, in)
}
