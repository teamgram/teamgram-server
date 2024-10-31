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

	"github.com/teamgram/proto/mtproto"

	"github.com/zeromicro/go-zero/zrpc"
)

var _ *mtproto.Bool

type SponsoredMessagesClient interface {
	AccountToggleSponsoredMessages(ctx context.Context, in *mtproto.TLAccountToggleSponsoredMessages) (*mtproto.Bool, error)
	MessagesViewSponsoredMessage(ctx context.Context, in *mtproto.TLMessagesViewSponsoredMessage) (*mtproto.Bool, error)
	MessagesClickSponsoredMessage(ctx context.Context, in *mtproto.TLMessagesClickSponsoredMessage) (*mtproto.Bool, error)
	MessagesReportSponsoredMessage(ctx context.Context, in *mtproto.TLMessagesReportSponsoredMessage) (*mtproto.Channels_SponsoredMessageReportResult, error)
	MessagesGetSponsoredMessages(ctx context.Context, in *mtproto.TLMessagesGetSponsoredMessages) (*mtproto.Messages_SponsoredMessages, error)
	ChannelsRestrictSponsoredMessages(ctx context.Context, in *mtproto.TLChannelsRestrictSponsoredMessages) (*mtproto.Updates, error)
	ChannelsViewSponsoredMessage(ctx context.Context, in *mtproto.TLChannelsViewSponsoredMessage) (*mtproto.Bool, error)
	ChannelsGetSponsoredMessages(ctx context.Context, in *mtproto.TLChannelsGetSponsoredMessages) (*mtproto.Messages_SponsoredMessages, error)
	ChannelsClickSponsoredMessage(ctx context.Context, in *mtproto.TLChannelsClickSponsoredMessage) (*mtproto.Bool, error)
	ChannelsReportSponsoredMessage(ctx context.Context, in *mtproto.TLChannelsReportSponsoredMessage) (*mtproto.Channels_SponsoredMessageReportResult, error)
}

type defaultSponsoredMessagesClient struct {
	cli zrpc.Client
}

func NewSponsoredMessagesClient(cli zrpc.Client) SponsoredMessagesClient {
	return &defaultSponsoredMessagesClient{
		cli: cli,
	}
}

// AccountToggleSponsoredMessages
// account.toggleSponsoredMessages#b9d9a38d enabled:Bool = Bool;
func (m *defaultSponsoredMessagesClient) AccountToggleSponsoredMessages(ctx context.Context, in *mtproto.TLAccountToggleSponsoredMessages) (*mtproto.Bool, error) {
	client := mtproto.NewRPCSponsoredMessagesClient(m.cli.Conn())
	return client.AccountToggleSponsoredMessages(ctx, in)
}

// MessagesViewSponsoredMessage
// messages.viewSponsoredMessage#673ad8f1 peer:InputPeer random_id:bytes = Bool;
func (m *defaultSponsoredMessagesClient) MessagesViewSponsoredMessage(ctx context.Context, in *mtproto.TLMessagesViewSponsoredMessage) (*mtproto.Bool, error) {
	client := mtproto.NewRPCSponsoredMessagesClient(m.cli.Conn())
	return client.MessagesViewSponsoredMessage(ctx, in)
}

// MessagesClickSponsoredMessage
// messages.clickSponsoredMessage#f093465 flags:# media:flags.0?true fullscreen:flags.1?true peer:InputPeer random_id:bytes = Bool;
func (m *defaultSponsoredMessagesClient) MessagesClickSponsoredMessage(ctx context.Context, in *mtproto.TLMessagesClickSponsoredMessage) (*mtproto.Bool, error) {
	client := mtproto.NewRPCSponsoredMessagesClient(m.cli.Conn())
	return client.MessagesClickSponsoredMessage(ctx, in)
}

// MessagesReportSponsoredMessage
// messages.reportSponsoredMessage#1af3dbb8 peer:InputPeer random_id:bytes option:bytes = channels.SponsoredMessageReportResult;
func (m *defaultSponsoredMessagesClient) MessagesReportSponsoredMessage(ctx context.Context, in *mtproto.TLMessagesReportSponsoredMessage) (*mtproto.Channels_SponsoredMessageReportResult, error) {
	client := mtproto.NewRPCSponsoredMessagesClient(m.cli.Conn())
	return client.MessagesReportSponsoredMessage(ctx, in)
}

// MessagesGetSponsoredMessages
// messages.getSponsoredMessages#9bd2f439 peer:InputPeer = messages.SponsoredMessages;
func (m *defaultSponsoredMessagesClient) MessagesGetSponsoredMessages(ctx context.Context, in *mtproto.TLMessagesGetSponsoredMessages) (*mtproto.Messages_SponsoredMessages, error) {
	client := mtproto.NewRPCSponsoredMessagesClient(m.cli.Conn())
	return client.MessagesGetSponsoredMessages(ctx, in)
}

// ChannelsRestrictSponsoredMessages
// channels.restrictSponsoredMessages#9ae91519 channel:InputChannel restricted:Bool = Updates;
func (m *defaultSponsoredMessagesClient) ChannelsRestrictSponsoredMessages(ctx context.Context, in *mtproto.TLChannelsRestrictSponsoredMessages) (*mtproto.Updates, error) {
	client := mtproto.NewRPCSponsoredMessagesClient(m.cli.Conn())
	return client.ChannelsRestrictSponsoredMessages(ctx, in)
}

// ChannelsViewSponsoredMessage
// channels.viewSponsoredMessage#beaedb94 channel:InputChannel random_id:bytes = Bool;
func (m *defaultSponsoredMessagesClient) ChannelsViewSponsoredMessage(ctx context.Context, in *mtproto.TLChannelsViewSponsoredMessage) (*mtproto.Bool, error) {
	client := mtproto.NewRPCSponsoredMessagesClient(m.cli.Conn())
	return client.ChannelsViewSponsoredMessage(ctx, in)
}

// ChannelsGetSponsoredMessages
// channels.getSponsoredMessages#ec210fbf channel:InputChannel = messages.SponsoredMessages;
func (m *defaultSponsoredMessagesClient) ChannelsGetSponsoredMessages(ctx context.Context, in *mtproto.TLChannelsGetSponsoredMessages) (*mtproto.Messages_SponsoredMessages, error) {
	client := mtproto.NewRPCSponsoredMessagesClient(m.cli.Conn())
	return client.ChannelsGetSponsoredMessages(ctx, in)
}

// ChannelsClickSponsoredMessage
// channels.clickSponsoredMessage#1445d75 flags:# media:flags.0?true fullscreen:flags.1?true channel:InputChannel random_id:bytes = Bool;
func (m *defaultSponsoredMessagesClient) ChannelsClickSponsoredMessage(ctx context.Context, in *mtproto.TLChannelsClickSponsoredMessage) (*mtproto.Bool, error) {
	client := mtproto.NewRPCSponsoredMessagesClient(m.cli.Conn())
	return client.ChannelsClickSponsoredMessage(ctx, in)
}

// ChannelsReportSponsoredMessage
// channels.reportSponsoredMessage#af8ff6b9 channel:InputChannel random_id:bytes option:bytes = channels.SponsoredMessageReportResult;
func (m *defaultSponsoredMessagesClient) ChannelsReportSponsoredMessage(ctx context.Context, in *mtproto.TLChannelsReportSponsoredMessage) (*mtproto.Channels_SponsoredMessageReportResult, error) {
	client := mtproto.NewRPCSponsoredMessagesClient(m.cli.Conn())
	return client.ChannelsReportSponsoredMessage(ctx, in)
}
